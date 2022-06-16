package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"eduid_captcha_service/internal/apiv1"
	"eduid_captcha_service/internal/httpserver"
	"eduid_captcha_service/pkg/configuration"
	"eduid_captcha_service/pkg/logger"
)

type service interface {
	Close(ctx context.Context) error
}

func main() {
	wg := &sync.WaitGroup{}
	ctx := context.Background()

	var (
		log      *logger.Logger
		mainLog  *logger.Logger
		services = map[string]service{}
	)

	cfg, err := configuration.Parse(logger.NewSimple("Configuration"))
	if err != nil {
		panic(err)
	}

	mainLog = logger.New("main", cfg.Production)
	log = logger.New("eduid_captcha_service", cfg.Production)

	apiv1, err := apiv1.New(ctx, cfg, log.New("apiv1"))
	if err != nil {
		panic(err)
	}
	httpserver, err := httpserver.New(ctx, cfg, apiv1, log.New("httpserver"))
	services["httpserver"] = httpserver
	if err != nil {
		panic(err)
	}

	// Handle sigterm and await termChan signal
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	<-termChan // Blocks here until interrupted

	mainLog.Info("HALTING SIGNAL!")

	for serviceName := range services {
		services[serviceName].Close(ctx)
		if err != nil {
			mainLog.Warn("Service", "Name", serviceName)
		}
	}

	wg.Wait() // Block here until are workers are done

	mainLog.Info("Stopped")
}
