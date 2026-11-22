package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Config holds runtime settings loaded strictly from environment variables.
type Config struct {
	UARTDevice      string
	UARTBaud        int
	MQTTBroker      string
	MQTTTopicPrefix string
	MQTTClientID    string
	MQTTUsername    string
	MQTTPassword    string
	WorkerCount     int
	QueueSize       int
}

// LoadFromEnv reads configuration from the process environment and fails fast
// if any mandatory variable is unset.
func LoadFromEnv() (Config, error) {
	cfg := Config{
		UARTDevice:      strings.TrimSpace(os.Getenv("UART_DEVICE")),
		UARTBaud:        getEnvInt("UART_BAUD", 115200),
		MQTTBroker:      strings.TrimSpace(os.Getenv("MQTT_BROKER")),
		MQTTTopicPrefix: getEnvDefault("MQTT_TOPIC_PREFIX", "rf/subghz"),
		MQTTClientID:    getEnvDefault("MQTT_CLIENT_ID", "flipper-subghz-demux"),
		MQTTUsername:    os.Getenv("MQTT_USERNAME"),
		MQTTPassword:    os.Getenv("MQTT_PASSWORD"),
		WorkerCount:     getEnvInt("WORKER_COUNT", 2),
		QueueSize:       getEnvInt("QUEUE_SIZE", 512),
	}

	var missing []string
	if cfg.UARTDevice == "" {
		missing = append(missing, "UART_DEVICE")
	}
	if cfg.MQTTBroker == "" {
		missing = append(missing, "MQTT_BROKER")
	}
	if len(missing) > 0 {
		return Config{}, fmt.Errorf("missing required environment variables: %s", strings.Join(missing, ", "))
	}
	return cfg, nil
}

func getEnvDefault(key, fallback string) string {
	if v := strings.TrimSpace(os.Getenv(key)); v != "" {
		return v
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if v := strings.TrimSpace(os.Getenv(key)); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			return n
		}
	}
	return fallback
}
