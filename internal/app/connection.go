package app

import (
	"errors"
	"io"

	log "github.com/sirupsen/logrus" //nolint: depguard
)

//go:generate moq -out zmock_read_write_closer_test.go -pkg app_test . ReadWriteCloser

type ReadWriteCloser interface {
	Read([]byte) (int, error)
	Write([]byte) (int, error)
	Close() error
}

func EchoHandler(bufferSize int) func(*log.Entry, ReadWriteCloser) {
	return func(logger *log.Entry, conn ReadWriteCloser) {
		defer conn.Close()
		logger.Println("Connection opened")
		buffer := make([]byte, bufferSize)

		for {
			sizeRead, err := conn.Read(buffer)
			if err != nil {
				if errors.Is(err, io.EOF) {
					logger.Println("Connection closed")
				} else {
					logger.Error(err)
				}
				return
			}
			sizeWritten, err := conn.Write(buffer[:sizeRead])
			if err != nil {
				logger.Error(err)
				return
			}
			for sizeWritten != sizeRead {
				moreWritten, err := conn.Write(buffer[sizeWritten:sizeRead])
				if err != nil {
					logger.Error(err)
					return
				}
				sizeWritten += moreWritten
			}
		}
	}
}
