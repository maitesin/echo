package infra

import (
	"net"

	"github.com/maitesin/echo/internal/app"
	log "github.com/sirupsen/logrus" //nolint: depguard
)

//go:generate moq -out zmock_accepter_test.go -pkg infra_test . Accepter

type Accepter interface {
	Accept() (net.Conn, error)
}

func ConnectionsHandler(logger *log.Logger, accepter Accepter, handler func(*log.Entry, app.ReadWriteCloser)) error {
	for {
		conn, err := accepter.Accept()
		if err != nil {
			return err
		}
		go handler(
			logger.WithField("remote", conn.RemoteAddr().String()),
			conn)
	}
}
