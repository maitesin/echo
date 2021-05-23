package infra_test

import (
	"errors"
	"net"
	"testing"

	"github.com/maitesin/echo/internal/app"
	"github.com/maitesin/echo/internal/infra"
	log "github.com/sirupsen/logrus" //nolint: depguard
	"github.com/stretchr/testify/require"
)

func TestMainLoop(t *testing.T) {
	testLogger := log.New()

	tests := []struct {
		name        string
		accepter    infra.Accepter
		expectedErr error
	}{
		{
			name: ``,
			accepter: &AccepterMock{AcceptFunc: func() (net.Conn, error) {
				return nil, errors.New("something went wrong")
			}},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := infra.ConnectionsHandler(testLogger, tt.accepter, func(*log.Entry, app.ReadWriteCloser) {})
			require.ErrorAs(t, err, &tt.expectedErr)
		})
	}
}
