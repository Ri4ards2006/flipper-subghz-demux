package main

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Ri4ards2006/flipper-subghz-demux/internal/config"
	"github.com/Ri4ards2006/flipper-subghz-demux/internal/mqtt"
	"github.com/Ri4ards2006/flipper-subghz-demux/internal/uart"
	"github.com/Ri4ards2006/flipper-subghz-demux/pkg/subghz"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	cfg, err := config.LoadFromEnv()
	if err != nil {
		logger.Error("fatal configuration error", "err", err)
		os.Exit(1)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	frames := make(chan subghz.Frame, cfg.QueueSize)
	reader := &uart.Reader{Device: cfg.UARTDevice, Baud: cfg.UARTBaud, Logger: logger}
	pub := &mqtt.Publisher{
		Broker:      cfg.MQTTBroker,
		TopicPrefix: cfg.MQTTTopicPrefix,
		ClientID:    cfg.MQTTClientID,
		Logger:      logger,
	}

	var wg sync.WaitGroup
	for i := 0; i < cfg.WorkerCount; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for f := range frames {
				if err := pub.Publish(ctx, f); err != nil {
					logger.Warn("publish failed", "worker", workerID, "err", err)
				}
			}
		}(i)
	}

	logger.Info("flipper-subghz-demux started",
		"uart_device", cfg.UARTDevice,
		"mqtt_broker", cfg.MQTTBroker,
		"workers", cfg.WorkerCount,
	)

	if err := reader.Run(ctx, frames); err != nil && !errors.Is(err, context.Canceled) {
		logger.Error("uart reader exited with error", "err", err)
	}

	wg.Wait()
	_ = pub.Close()
	logger.Info("flipper-subghz-demux stopped cleanly")
}
