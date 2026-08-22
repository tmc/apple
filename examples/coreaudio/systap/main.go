// Command systap records macOS system audio output to a 48kHz WAV file using pure-Go Core Audio process taps.
//
// Usage:
//
//	systap -o out.wav [-pid N | -bundle com.brave.Browser] [-d 10s]
package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tmc/apple/foundation"
)

func main() {
	opts := Options{}
	flag.StringVar(&opts.OutputFile, "o", "out.wav", "Output WAV file path ('-' for stdout)")
	flag.IntVar(&opts.PID, "pid", 0, "Target process ID to capture (0 for system audio output)")
	flag.StringVar(&opts.BundleID, "bundle", "", "Target bundle ID to capture (e.g. 'com.brave.Browser')")
	flag.IntVar(&opts.SampleRate, "rate", 48000, "Audio sample rate in Hz")
	flag.IntVar(&opts.Channels, "channels", 2, "Number of audio channels")
	flag.BoolVar(&opts.Verbose, "v", false, "Enable verbose output")
	duration := flag.Duration("d", 0, "Recording duration (e.g. 10s, 1m). Default 0 records until SIGINT/SIGTERM")
	flag.Parse()

	if err := run(opts, *duration); err != nil {
		fmt.Fprintf(os.Stderr, "systap: %v\n", err)
		os.Exit(1)
	}
}

func run(opts Options, duration time.Duration) error {
	pool := foundation.NewNSAutoreleasePool()
	defer pool.Drain()

	rec, err := NewRecorder(opts)
	if err != nil {
		return fmt.Errorf("initialize recorder: %w", err)
	}

	if opts.Verbose {
		fmt.Fprintf(os.Stderr, "systap: starting audio capture to %s (rate=%d, channels=%d)...\n",
			opts.OutputFile, opts.SampleRate, opts.Channels)
	}

	if err := rec.Start(); err != nil {
		return fmt.Errorf("start capture: %w", err)
	}
	defer rec.Stop()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	if duration > 0 {
		timer := time.NewTimer(duration)
		select {
		case <-timer.C:
			if opts.Verbose {
				fmt.Fprintf(os.Stderr, "systap: duration %v elapsed, stopping...\n", duration)
			}
		case sig := <-sigCh:
			timer.Stop()
			if opts.Verbose {
				fmt.Fprintf(os.Stderr, "systap: received signal %v, stopping...\n", sig)
			}
		}
	} else {
		sig := <-sigCh
		if opts.Verbose {
			fmt.Fprintf(os.Stderr, "systap: received signal %v, stopping...\n", sig)
		}
	}

	if err := rec.Stop(); err != nil {
		return fmt.Errorf("stop capture: %w", err)
	}

	if opts.Verbose {
		fmt.Fprintf(os.Stderr, "systap: capture complete. Callbacks: %d, Bytes written: %d\n",
			rec.CallbacksCount(), rec.BytesRecorded())
	}

	return nil
}
