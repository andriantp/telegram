package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"telegram/telegram"
)

func Run(repo telegram.RepositoryI) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	defer signal.Stop(sigChan)

	errChan := make(chan error, 1)
	go func() {
		if err := repo.Received(ctx); err != nil {
			errChan <- fmt.Errorf("[Received] failed: %w", err)
			cancel()
		}
	}()

	select {
	case sig := <-sigChan:
		var message string
		switch sig {
		case syscall.SIGHUP:
			message = "[hungup]"
		case syscall.SIGINT:
			message = "[interrupt]"
		case syscall.SIGTERM:
			message = "[force stop]"
		case syscall.SIGQUIT:
			message = "[stop and core dump]"
		default:
			message = "[unknown signal]"
		}
		cancel()
		return errors.New(message)

	case err := <-errChan:
		cancel()
		return err

	case <-ctx.Done():
		return ctx.Err()
	}

}
