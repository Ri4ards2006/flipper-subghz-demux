package uart

import (
	"context"
	"log/slog"

	"github.com/Ri4ards2006/flipper-subghz-demux/pkg/subghz"
)

// Reader streams Sub-GHz frames from a serial UART device.
type Reader struct {
	Device string
	Baud   int
	Logger *slog.Logger
}

// Run blocks until ctx is cancelled, emitting parsed frames on out.
func (r *Reader) Run(ctx context.Context, out chan<- subghz.Frame) error {
	defer close(out)
	r.Logger.Info("uart reader initialized", "device", r.Device, "baud", r.Baud)
	<-ctx.Done()
	return ctx.Err()
}
