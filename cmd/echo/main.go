package main

import (
	"fmt"
	"net"
	"os"

	"github.com/maitesin/echo/config"
	"github.com/maitesin/echo/internal/app"
	"github.com/maitesin/echo/internal/infra"
	log "github.com/sirupsen/logrus" //nolint: depguard
)

const (
	exitStatusFailedConfiguration int = iota + 1
	exitStatusFailedListen
	exitStatusFailedAccept
)

func main() {
	logger := log.New()
	logger.SetFormatter(&log.JSONFormatter{})

	cfg, err := config.New(os.Args[0], os.Args[1:])
	if err != nil {
		logger.Println(err)
		os.Exit(exitStatusFailedConfiguration)
	}

	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	if err != nil {
		logger.Println(err)
		os.Exit(exitStatusFailedListen)
	}

	logger.WithFields(log.Fields{
		"port":        cfg.Port,
		"host":        cfg.Host,
		"buffer-size": cfg.BufferSize,
	}).Println("Echo server started")

	err = infra.ConnectionsHandler(logger, ln, app.EchoHandler(cfg.BufferSize))
	if err != nil {
		logger.Println(err)
		os.Exit(exitStatusFailedAccept)
	}
}
