package signal

import (
	"os"
	"os/signal"
	"syscall"
)

type logger interface {
	Info(args ...interface{})
}

func SetupSignalHandler(log logger) chan struct{} {
	termC := make(chan os.Signal, 1)
	intC := make(chan os.Signal, 1)
	stop := make(chan struct{})

	signal.Notify(termC, syscall.SIGTERM)
	signal.Notify(intC, os.Interrupt)

	go func() {
		select {
		case <-termC:
			log.Info("Received SIGTERM. Terminating...")
		case <-intC:
			log.Info("Received Interrupt. Terminating...")
		}

		stop <- struct{}{}
	}()

	return stop
}
