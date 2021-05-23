package config

import (
	"flag"
	"fmt"
)

const (
	defaultPort       = 7
	defaultHost       = "127.0.0.1"
	defaultBufferSize = 100
)

type Config struct {
	Host       string
	Port       int
	BufferSize int
}

func New(name string, args []string) (Config, error) {
	flags := flag.NewFlagSet(name, flag.ContinueOnError)
	var cfg Config

	flags.IntVar(&cfg.Port, "port", defaultPort, "Port for the echo server")
	flags.StringVar(&cfg.Host, "host", defaultHost, "Hostname or IP address where the port will be open")
	flags.IntVar(
		&cfg.BufferSize,
		"buffer-size",
		defaultBufferSize,
		"Size of the buffer, in bytes, that will be used to read from the opened connection",
	)

	err := flags.Parse(args)
	if err != nil {
		return Config{}, NewParseConfigError(err)
	}

	if cfg.BufferSize < 1 {
		return Config{}, NewInvalidConfigError("-buffer-size", fmt.Sprint(cfg.BufferSize))
	}

	if cfg.Port < 1 {
		return Config{}, NewInvalidConfigError("-port", fmt.Sprint(cfg.Port))
	}

	return cfg, nil
}
