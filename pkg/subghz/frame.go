package subghz

import "time"

// Frame represents a decoded Sub-GHz baseband capture received over UART.
type Frame struct {
	Timestamp  time.Time `json:"timestamp"`
	FreqHz     uint64    `json:"freq_hz"`
	Modulation string    `json:"modulation"`
	Protocol   string    `json:"protocol"`
	RSSI       int       `json:"rssi"`
	RawHex     string    `json:"raw_hex"`
}
