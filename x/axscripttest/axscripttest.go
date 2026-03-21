package axscripttest

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/tmc/apple/x/axuiautomation"
	"github.com/tmc/macgo"
	"golang.org/x/tools/txtar"
	"rsc.io/script"
	"rsc.io/script/scripttest"
)

// Config configures the script engine for a target application.
type Config struct {
	App     string        // bundle id, pid, or name
	Timeout time.Duration // default wait timeout (default 10s)
}

// Engine returns a script engine and application configured for UI automation.
// The caller is responsible for closing the application when done.
func Engine(cfg Config) (*script.Engine, *axuiautomation.Application, error) {
	if cfg.Timeout == 0 {
		cfg.Timeout = 10 * time.Second
	}

	app, err := axuiautomation.NewApplication(cfg.App)
	if err != nil {
		return nil, nil, err
	}

	e := script.NewEngine()

	cmds := scripttest.DefaultCmds()
	cmds["app-activate"] = cmdAppActivate(app)
	cmds["window-wait"] = cmdWindowWait(app, cfg.Timeout)
	cmds["menu-click"] = cmdMenuClick(app)
	cmds["click"] = cmdClick(app)
	cmds["type"] = cmdType(app)
	cmds["select"] = cmdSelect(app)
	cmds["exists"] = cmdExists(app)
	cmds["value-wait"] = cmdValueWait(app, cfg.Timeout)
	cmds["screenshot"] = cmdScreenshot(app)
	e.Cmds = cmds

	conds := scripttest.DefaultConds()
	conds["trusted"] = script.OnceCondition("process has accessibility permissions", func() (bool, error) {
		if axuiautomation.IsProcessTrusted() {
			return true, nil
		}
		// If running inside a macgo bundle, prompt and wait for the
		// user to grant accessibility in System Settings.
		if os.Getenv("MACGO_BUNDLE_PATH") != "" {
			axuiautomation.PromptForAccessibility()
			waitForTrust(30 * time.Second)
			return axuiautomation.IsProcessTrusted(), nil
		}
		return false, nil
	})
	conds["app-running"] = script.Condition("target application is running", func(s *script.State) (bool, error) {
		return app.IsRunning(), nil
	})
	conds["screencapture"] = script.OnceCondition("process has screen recording permissions", func() (bool, error) {
		return axuiautomation.IsScreenRecordingTrusted(), nil
	})
	e.Conds = conds

	return e, app, nil
}

// EnsureTrusted sets up a TCC identity via macgo and triggers the system
// accessibility permission dialog if the process is not already trusted.
// After prompting, it polls for up to 30 seconds waiting for the user to
// approve accessibility in System Settings. If trust is not granted within
// that window the function returns nil and tests should use the [!trusted]
// condition to skip gracefully.
//
// Call this from TestMain before m.Run, since macgo may re-exec the process
// inside an app bundle:
//
//	func TestMain(m *testing.M) {
//	    runtime.LockOSThread()
//	    axscripttest.EnsureTrusted()
//	    os.Exit(m.Run())
//	}
func EnsureTrusted() error {
	cfg := macgo.NewConfig().
		WithAppName("axscripttest").
		WithPermissions(macgo.Accessibility).
		WithAdHocSign()
	cfg.BundleID = "dev.tmc.axscripttest"
	// Always call Start — when running as a macgo child (re-launched
	// via LaunchServices), Start restores the working directory and
	// sets up pipe redirection. Without this the child exits before
	// the parent reads the PID from the control FIFO.
	if err := macgo.Start(cfg); err != nil {
		return fmt.Errorf("macgo start: %w", err)
	}
	if axuiautomation.IsProcessTrusted() {
		return nil
	}
	axuiautomation.PromptForAccessibility()
	return waitForTrust(30 * time.Second)
}

// waitForTrust polls IsProcessTrusted at 500ms intervals for up to timeout.
// Returns nil regardless — callers use the [!trusted] script condition to
// skip tests when trust is not granted.
func waitForTrust(timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if axuiautomation.IsProcessTrusted() {
			return nil
		}
		time.Sleep(500 * time.Millisecond)
	}
	return nil
}

// RunFile runs a txtar script file against the configured application.
// Scripts should use the [trusted] condition to skip gracefully when
// accessibility permissions are not available. Use EnsureTrusted in
// TestMain to trigger the system permission dialog.
func RunFile(t *testing.T, cfg Config, file string) {
	t.Helper()

	data, err := os.ReadFile(file)
	if err != nil {
		t.Fatal(err)
	}

	ar := txtar.Parse(data)

	e, app, err := Engine(cfg)
	if err != nil {
		t.Fatal(err)
	}
	defer app.Close()

	workdir := t.TempDir()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	s, err := script.NewState(ctx, workdir, nil)
	if err != nil {
		t.Fatal(err)
	}
	if err := s.ExtractFiles(ar); err != nil {
		t.Fatal(err)
	}

	// Use scripttest.Run which properly handles the "skip" command
	// sentinel, calling t.Skip instead of t.Fatal.
	scripttest.Run(t, e, s, file, bytes.NewReader(ar.Comment))
}
