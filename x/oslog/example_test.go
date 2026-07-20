package oslog_test

import (
	"errors"
	"log/slog"

	"github.com/tmc/apple/x/oslog"
)

// Log formatted messages to the unified logging system. View them with
// "log stream --predicate 'subsystem == \"com.example.app\"'" or Console.app.
func Example() {
	log := oslog.New("com.example.app", "network")

	host := "api.example.com"
	log.Info("connecting to %{public}s", host)
	log.Default("received %d bytes in %dms", 4096, 12)

	if err := errors.New("timeout"); err != nil {
		log.Error("request failed: %{public}s", err)
	}
}

// Correlate a group of log messages under a single activity so a reader can
// collapse all the work of one operation together, even across threads.
func ExampleActivity() {
	log := oslog.New("com.example.app", "worker")

	act := oslog.NewActivity("handle request")
	defer act.Enter()() // leave the scope when the function returns

	log.Info("started") // both messages are tagged with act.ID()
	log.Default("processing")
}

// Log a panic (with its Go stack) to the system log at Fault level, recovering
// so the goroutine survives. Defer it at any goroutine boundary.
func ExampleLogger_Recover() {
	log := oslog.New("com.example.app", "worker")

	func() {
		defer log.Recover() // logs panic + stack at Fault, then swallows it
		panic("unexpected state")
	}()

	log.Info("recovered, still running")
}

// Route Go's structured logger (log/slog) to the system log. Levels map to
// os_log types and attributes render into the message.
func ExampleHandler() {
	logger := oslog.New("com.example.app", "worker")
	slog.SetDefault(slog.New(oslog.NewHandler(logger, &oslog.HandlerOptions{
		Level: slog.LevelInfo,
	})))

	slog.Info("job started", "id", 42, "queue", "default")
	slog.With("request_id", "abc123").WithGroup("http").
		Info("handled", "method", "GET", "status", 200)
}
