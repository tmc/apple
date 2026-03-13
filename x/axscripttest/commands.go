package axscripttest

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/tmc/apple/x/axuiautomation"
	"rsc.io/script"
)

func cmdAppActivate(app *axuiautomation.Application) script.Cmd {
	return script.Command(
		script.CmdUsage{Summary: "bring application to foreground", Args: ""},
		func(s *script.State, args ...string) (script.WaitFunc, error) {
			return nil, app.Activate()
		},
	)
}

func cmdWindowWait(app *axuiautomation.Application, defaultTimeout time.Duration) script.Cmd {
	return script.Command(
		script.CmdUsage{Summary: "wait for a window with the given title", Args: "<title> [timeout]"},
		func(s *script.State, args ...string) (script.WaitFunc, error) {
			if len(args) < 1 {
				return nil, script.ErrUsage
			}
			title := args[0]
			timeout := defaultTimeout
			if len(args) >= 2 {
				d, err := time.ParseDuration(args[1])
				if err != nil {
					return nil, fmt.Errorf("parsing timeout: %w", err)
				}
				timeout = d
			}
			_, err := app.WaitForWindow(title, timeout)
			return nil, err
		},
	)
}

func cmdMenuClick(app *axuiautomation.Application) script.Cmd {
	return script.Command(
		script.CmdUsage{Summary: "click a menu item by path", Args: "<item>..."},
		func(s *script.State, args ...string) (script.WaitFunc, error) {
			if len(args) < 1 {
				return nil, script.ErrUsage
			}
			return nil, app.ClickMenuItem(args)
		},
	)
}

func cmdClick(app *axuiautomation.Application) script.Cmd {
	return script.Command(
		script.CmdUsage{Summary: "find and click an element", Args: "role=R [title=T] [id=I]"},
		func(s *script.State, args ...string) (script.WaitFunc, error) {
			sel, _, err := parseSelector(args)
			if err != nil {
				return nil, err
			}
			elem, err := sel.find(app)
			if err != nil {
				return nil, err
			}
			defer elem.Release()
			return nil, elem.Click()
		},
	)
}

func cmdType(app *axuiautomation.Application) script.Cmd {
	return script.Command(
		script.CmdUsage{Summary: "type text into the focused element", Args: "<text>"},
		func(s *script.State, args ...string) (script.WaitFunc, error) {
			if len(args) < 1 {
				return nil, script.ErrUsage
			}
			return nil, axuiautomation.TypeText(strings.Join(args, " "))
		},
	)
}

func cmdSelect(app *axuiautomation.Application) script.Cmd {
	return script.Command(
		script.CmdUsage{Summary: "select a value from a popup button", Args: "<popup-title> <value>"},
		func(s *script.State, args ...string) (script.WaitFunc, error) {
			if len(args) < 2 {
				return nil, script.ErrUsage
			}
			elem := app.Descendants().ByRole("AXPopUpButton").ByTitle(args[0]).First()
			if elem == nil {
				return nil, fmt.Errorf("popup button %q not found", args[0])
			}
			defer elem.Release()
			return nil, elem.SelectValue(args[1])
		},
	)
}

func cmdExists(app *axuiautomation.Application) script.Cmd {
	return script.Command(
		script.CmdUsage{Summary: "assert an element exists", Args: "role=R [title=T] [id=I]"},
		func(s *script.State, args ...string) (script.WaitFunc, error) {
			sel, _, err := parseSelector(args)
			if err != nil {
				return nil, err
			}
			elem, err := sel.find(app)
			if err != nil {
				return nil, err
			}
			elem.Release()
			return nil, nil
		},
	)
}

func cmdValueWait(app *axuiautomation.Application, defaultTimeout time.Duration) script.Cmd {
	return script.Command(
		script.CmdUsage{Summary: "wait for an element value to contain a string", Args: "role=R [title=T] [id=I] <contains> [timeout]"},
		func(s *script.State, args ...string) (script.WaitFunc, error) {
			sel, rest, err := parseSelector(args)
			if err != nil {
				return nil, err
			}
			if len(rest) < 1 {
				return nil, script.ErrUsage
			}
			contains := rest[0]
			timeout := defaultTimeout
			if len(rest) >= 2 {
				d, err := time.ParseDuration(rest[1])
				if err != nil {
					return nil, fmt.Errorf("parsing timeout: %w", err)
				}
				timeout = d
			}
			deadline := time.Now().Add(timeout)
			for time.Now().Before(deadline) {
				select {
				case <-s.Context().Done():
					return nil, s.Context().Err()
				default:
				}
				elem, err := sel.find(app)
				if err == nil {
					v := elem.Value()
					elem.Release()
					if strings.Contains(v, contains) {
						return nil, nil
					}
				}
				time.Sleep(100 * time.Millisecond)
			}
			return nil, fmt.Errorf("timed out waiting for value containing %q", contains)
		},
	)
}

func cmdScreenshot(app *axuiautomation.Application) script.Cmd {
	return script.Command(
		script.CmdUsage{Summary: "capture the main window to a PNG file", Args: "<file>"},
		func(s *script.State, args ...string) (script.WaitFunc, error) {
			if len(args) < 1 {
				return nil, script.ErrUsage
			}
			win := app.MainWindow()
			if win == nil {
				return nil, fmt.Errorf("no main window")
			}
			defer win.Release()
			png, err := win.Screenshot()
			if err != nil {
				return nil, fmt.Errorf("screenshot: %w", err)
			}
			return nil, os.WriteFile(s.Path(args[0]), png, 0o644)
		},
	)
}
