package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ckinan/sysmon/internal/collector"
)

func main() {
	// context.WithCancel gives us a cancel function to stop the collector cleanly.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // ensure goroutine is always stopped when main() exits

	// start collector - returns a channel immediately, goroutine runs in background
	snapshotCh := collector.Start(ctx, 2*time.Second)

	// Read a few snapshots then exit (for now)
	for range 3 {
		snap, ok := <-snapshotCh
		if !ok {
			break // channel was closed - collector stopped
		}
		fmt.Printf("RAM used: %d kB, processes: %d\n", snap.Ram.MemUsed, len(snap.Processes))
	}
}
