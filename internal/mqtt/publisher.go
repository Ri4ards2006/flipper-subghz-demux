package mqtt

import (
	"context"
	"log/slog"

	"github.com/Ri4ards2006/flipper-subghz-demux/pkg/subghz"
)

// Publisher forwards decoded Sub-GHz frames to an MQTT broker.
type Publisher struct {
	Broker      string
	TopicPrefix string
	ClientID    string
	Logger      *slog.Logger
}

// Publish emits a single frame (wire-up point for paho.mqtt.golang).
func (p *Publisher) Publish(_ context.Context, f subghz.Frame) error {
	p.Logger.Debug("publishing subghz frame",
		"topic_prefix", p.TopicPrefix,
		"freq_hz", f.FreqHz,
		"protocol", f.Protocol,
	)
	return nil
}

// Close cleanly disconnects the publisher.
func (p *Publisher) Close() error {
	p.Logger.Info("mqtt publisher closed", "broker", p.Broker)
	return nil
}
