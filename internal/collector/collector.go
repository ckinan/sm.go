package collector

import (
	"context"
	"log/slog"
	"time"

	"github.com/ckinan/sysmon/internal"
)

type Snapshot struct {
	Ram       internal.Ram
	Processes []internal.Process
}

func Start(ctx context.Context, interval time.Duration) <-chan Snapshot {
	ch := make(chan Snapshot, 1)

	go func() {
		defer close(ch) // signal consumers when goroutine exits
		ticker := time.NewTicker(interval)
		defer ticker.Stop() // vmlinuz

		for {
			select {
			case <-ctx.Done():
				// Context was cancelled (e.g. user pressed q, or main() exited)
				// Return immediately, defer close(ch) will run
				return
			case <-ticker.C:
				// Ticker fired: collect metrics and send Snapshot
				snap, err := collect()
				if err != nil {
					// skip this tick on error (e.g. a /proc read failed)
					slog.Warn("error reading resources", "error", err)
					continue
				}
				select {
				case ch <- snap:
					// snapshot sent successfully
				default:
					// Consumer hasn't read the previous snapshot yet
					// Drop this tick rather than blocking the gorouting
					// This keeps the collector running even if the UI is slow
				}

			}
		}
	}()
	return ch // return immediately, goroutine runs in background
}

func collect() (Snapshot, error) {
	ram, err := internal.GetRam()
	if err != nil {
		return Snapshot{}, err
	}

	processes, err := internal.ListProcess()
	if err != nil {
		return Snapshot{}, err
	}

	return Snapshot{
		Ram:       ram,
		Processes: processes,
	}, nil
}
