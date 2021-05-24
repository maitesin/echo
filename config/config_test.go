package config_test

import (
	"testing"

	"github.com/maitesin/echo/config"
	"github.com/stretchr/testify/require"
)

func validConfig() config.Config {
	return config.Config{
		Port:       7,
		Host:       "127.0.0.1",
		BufferSize: 100,
	}
}

type configMutator func(config.Config) config.Config

func noopConfigMutator(cfg config.Config) config.Config { return cfg }

func TestNew(t *testing.T) {
	tests := []struct {
		name          string
		args          []string
		configMutator configMutator
		expectedErr   error
	}{
		{
			name: `Given an empty list of arguments,
                   when the new configuration function is called,
                   then uses the default values for the application`,
			configMutator: noopConfigMutator,
		},
		{
			name: `Given a valid list of arguments with all arguments available present,
                   when the new configuration function is called,
                   then uses the values provided in the arguments for the application`,
			args: []string{"-port", "7000", "-host", "0.0.0.0", "-buffer-size", "1024"},
			configMutator: func(cfg config.Config) config.Config {
				cfg.Port = 7000
				cfg.Host = "0.0.0.0"
				cfg.BufferSize = 1024
				return cfg
			},
		},
		{
			name: `Given an invalid list of arguments,
                   when the new configuration function is called,
                   then returns a parsing configuration error`,
			args:          []string{"-wololo"},
			configMutator: noopConfigMutator,
			expectedErr:   config.ParseConfigError{},
		},
		{
			name: `Given a valid list of arguments, but with an invalid value used for the port argument,
                   when the new configuration function is called,
                   then returns an invalid configuration error`,
			args:          []string{"-port", "-10"},
			configMutator: noopConfigMutator,
			expectedErr:   config.InvalidConfigError{},
		},
		{
			name: `Given a valid list of arguments, but with an invalid value used for the buffer size argument,
                   when the new configuration function is called,
                   then returns an invalid configuration error`,
			args:          []string{"-buffer-size", "-10"},
			configMutator: noopConfigMutator,
			expectedErr:   config.InvalidConfigError{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			cfg, err := config.New("echo", tt.args)
			if tt.expectedErr != nil {
				require.ErrorAs(t, err, &tt.expectedErr)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.configMutator(validConfig()), cfg)
			}
		})
	}
}
