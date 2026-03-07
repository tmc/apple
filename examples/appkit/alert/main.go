// Alert shows a native macOS alert dialog from the command line.
//
// Usage:
//
//	alert "Build complete"
//	alert -style warning "Delete files?"
//	alert -btn OK -btn Cancel "Proceed?"
//	alert -j -btn Yes -btn No "Continue?"
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/tmc/apple/appkit"
)

type stringSlice []string

func (s *stringSlice) String() string { return strings.Join(*s, ", ") }
func (s *stringSlice) Set(v string) error {
	*s = append(*s, v)
	return nil
}

var (
	buttons  stringSlice
	style    string
	msg      string
	jsonOut  bool
	message  string
	exitCode int
)

func main() {
	flag.Var(&buttons, "btn", "button title (repeatable)")
	flag.StringVar(&style, "style", "info", "alert style: info, warning, critical")
	flag.StringVar(&msg, "msg", "", "informative text")
	flag.BoolVar(&jsonOut, "j", false, "JSON output")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: alert [-style info|warning|critical] [-btn title]... [-msg text] [-j] <message>\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(2)
	}
	message = strings.Join(flag.Args(), " ")

	if len(buttons) == 0 {
		buttons = stringSlice{"OK"}
	}

	var alertStyle appkit.NSAlertStyle
	switch style {
	case "info":
		alertStyle = appkit.NSAlertStyleInformational
	case "warning":
		alertStyle = appkit.NSAlertStyleWarning
	case "critical":
		alertStyle = appkit.NSAlertStyleCritical
	default:
		fmt.Fprintf(os.Stderr, "alert: unknown style %q\n", style)
		os.Exit(2)
	}

	appkit.RunApp(func(app appkit.NSApplication, delegate appkit.NSApplicationDelegateObject) {
		alert := appkit.GetNSAlertClass().Alloc().Init()
		alert.SetMessageText(message)
		alert.SetAlertStyle(alertStyle)
		if msg != "" {
			alert.SetInformativeText(msg)
		}
		for _, b := range buttons {
			alert.AddButtonWithTitle(b)
		}

		response := alert.RunModal()
		buttonIndex := int(response - appkit.AlertFirstButtonReturn)

		if jsonOut {
			label := ""
			if buttonIndex >= 0 && buttonIndex < len(buttons) {
				label = buttons[buttonIndex]
			}
			json.NewEncoder(os.Stdout).Encode(map[string]any{
				"button": buttonIndex,
				"label":  label,
			})
		}

		exitCode = buttonIndex
		app.Terminate(app)
		os.Exit(exitCode)
	})
}
