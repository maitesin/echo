package config_test

import (
	"errors"
	"testing"

	"github.com/maitesin/echo/config"
	"github.com/stretchr/testify/require"
)

func TestInvalidConfigError(t *testing.T) {
	t.Parallel()

	err := config.NewInvalidConfigError("-wololo", "42")
	require.Equal(t, `invalid configuration in field "-wololo" with value "42"`, err.Error())
}

func TestParseConfigError(t *testing.T) {
	t.Parallel()

	err := config.NewParseConfigError(errors.New("something went wrong"))
	require.Equal(t, "error parsing parameters: something went wrong", err.Error())
}
