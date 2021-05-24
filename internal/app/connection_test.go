package app_test

import (
	"io"
	"testing"

	"github.com/maitesin/echo/internal/app"
	log "github.com/sirupsen/logrus" //nolint: depguard
)

func validReadWriteCloser() *ReadWriteCloserMock {
	conn := &ReadWriteCloserMock{
		CloseFunc: func() error {
			return nil
		},
		WriteFunc: func(bytes []byte) (int, error) {
			return len(bytes), nil
		},
	}

	conn.ReadFunc = func(bytes []byte) (int, error) {
		if len(conn.ReadCalls()) == 1 {
			return len(bytes), nil
		}
		return 0, io.EOF
	}

	return conn
}

type readWriteCloserMutator func(app.ReadWriteCloser) app.ReadWriteCloser

func noopReadWriteCloserMutator(readWriteCloser app.ReadWriteCloser) app.ReadWriteCloser {
	return readWriteCloser
}

func TestEchoHandler(t *testing.T) {
	const bufferSize = 5
	testLogger := log.New().WithFields(log.Fields{})

	tests := []struct {
		name                   string
		readWriteCloserMutator readWriteCloserMutator
	}{
		{
			name: `Given a valid read write closer,
                   when the echo handler is called,
                   then echos the data provided and finishes when the read returns an EOF`,
			readWriteCloserMutator: noopReadWriteCloserMutator,
		},
		{
			name: `Given a valid read write closer,
                   when the echo handler is called,
                   then echos the data provided, with multiple writes, and finishes when the read returns an EOF`,
			readWriteCloserMutator: func(app.ReadWriteCloser) app.ReadWriteCloser {
				conn := validReadWriteCloser()

				conn.WriteFunc = func(bytes []byte) (int, error) {
					switch len(conn.WriteCalls()) {
					case 1:
						return 3, nil
					case 2:
						return 2, nil
					}

					t.Fatal()
					return 0, nil
				}

				return conn
			},
		},
		{
			name: `Given a failing read write closer that fails when reading from it,
                   when the echo handler is called,
                   then fails to read data and logs that the connection failed`,
			readWriteCloserMutator: func(app.ReadWriteCloser) app.ReadWriteCloser {
				conn := validReadWriteCloser()

				conn.ReadFunc = func([]byte) (int, error) {
					return 0, io.ErrUnexpectedEOF
				}

				return conn
			},
		},
		{
			name: `Given a failing read write closer that fails when writing the first time to it,
                   when the echo handler is called,
                   then fails to write data an logs that the connection failed`,
			readWriteCloserMutator: func(app.ReadWriteCloser) app.ReadWriteCloser {
				conn := validReadWriteCloser()

				conn.WriteFunc = func([]byte) (int, error) {
					return 0, io.ErrShortWrite
				}

				return conn
			},
		},
		{
			name: `Given a failing read write closer that fails when writing the second time to it,
                   when the echo handler is called,
                   then fails to write data an logs that the connection failed`,
			readWriteCloserMutator: func(app.ReadWriteCloser) app.ReadWriteCloser {
				conn := validReadWriteCloser()

				conn.WriteFunc = func([]byte) (int, error) {
					switch len(conn.WriteCalls()) {
					case 1:
						return 3, nil
					default:
						return 0, io.ErrShortWrite
					}
				}

				return conn
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			conn := tt.readWriteCloserMutator(validReadWriteCloser())
			handler := app.EchoHandler(bufferSize)

			handler(testLogger, conn)
		})
	}
}
