# flipper-subghz-demux

[![Go](https://img.shields.io/badge/Go-1.25+-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![Platform](https://img.shields.io/badge/Platform-Raspberry%20Pi%20Zero%20W-C51A4A?logo=raspberrypi&logoColor=white)](https://www.raspberrypi.com/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

Lightweight Go baseband bridge running on a **Raspberry Pi Zero W** that reads raw **433 / 868 MHz Sub-GHz** frames from a **Flipper Zero** over UART, deserializes custom/Protobuf baseband payloads, and publishes structured telemetry to an **MQTT** broker.

## Features

- **Zero-Trust Configuration:** Strict environment-variable loading via `os.Getenv` with fail-fast startup validation.
- **ARMv6-Friendly Concurrency:** Low-memory Go worker design tailored for the single-core Raspberry Pi Zero W.
- **Graceful Shutdown:** Context-driven teardown on `SIGINT`/`SIGTERM` to flush in-flight UART frames cleanly.
- **Modular Demux Layer:** Clean separation between UART ingestion (`internal/uart`), baseband decoding (`pkg/subghz`), and MQTT publishing (`internal/mqtt`).

## Quickstart

```bash
# 1. Configure environment
cp .env.example .env
# Edit .env and export variables (or source it)

# 2. Build for Raspberry Pi Zero W (ARMv6)
GOOS=linux GOARCH=arm GOARM=6 CGO_ENABLED=0 \
  go build -trimpath -ldflags="-s -w" -o bin/flipper-subghz-demux ./cmd/flipper-subghz-demux

# 3. Run locally
UART_DEVICE=/dev/serial0 MQTT_BROKER=tcp://mqtt.example.local:1883 \
  ./bin/flipper-subghz-demux
```

## Architecture

See [docs/ARCHITECTURE.md](docs/ARCHITECTURE.md) for the data flow, tech stack, and Flipper Zero ↔ Pi Zero W UART wiring diagram.
