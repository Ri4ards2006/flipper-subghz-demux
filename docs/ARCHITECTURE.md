# Architecture — `flipper-subghz-demux`

## Tech Stack

- **Language:** Go 1.25+ (`CGO_ENABLED=0` static ARMv6 binary for Pi Zero W)
- **Hardware Input:** Flipper Zero (CC1101 Sub-GHz transceiver, 433.92 / 868.35 MHz) via 3.3V UART
- **Transport Output:** MQTT v3.1.1 / v5 broker (e.g., Eclipse Mosquitto on K3s)

## Data Flow

```mermaid
flowchart LR
    FZ["Flipper Zero (CC1101 Sub-GHz)"] -->|"UART 3.3V (115200 8N1)"| READER["internal/uart Reader"]
    READER -->|"Raw Frame Bytes"| DEMUX["pkg/subghz Demuxer"]
    DEMUX -->|"Normalized JSON/Telemetry"| PUB["internal/mqtt Publisher"]
    PUB -->|"rf/subghz/<freq>/<proto>"| BROKER["MQTT Broker (Mosquitto)"]
```

## Hardware Wiring (Flipper Zero ↔ Raspberry Pi Zero W)

> **Voltage Warning:** Both Flipper Zero GPIO and Raspberry Pi Zero W UART operate at **3.3V logic**. Do not connect 5V lines to RX/TX.

| Flipper Zero Pin | Signal | Pi Zero W Physical Pin | Pi BCM GPIO |
| :--- | :--- | :--- | :--- |
| Pin 13 (`TX`) | UART TX → RX | Pin 10 (`RXD0`) | GPIO 15 |
| Pin 14 (`RX`) | UART RX ← TX | Pin 8 (`TXD0`) | GPIO 14 |
| Pin 11 / 18 (`GND`) | Common Ground | Pin 6 (`GND`) | GND |

## Module Boundaries

- `cmd/flipper-subghz-demux/`: Entrypoint, signal handling (`SIGINT`/`SIGTERM`), lifecycle orchestration.
- `internal/config/`: Strict `os.Getenv` loader with fail-fast validation.
- `internal/uart/`: Non-blocking UART stream framing and sync-byte detection.
- `internal/mqtt/`: Broker connection management and topic routing.
- `pkg/subghz/`: Public baseband frame types and Protobuf/custom payload decoders.
