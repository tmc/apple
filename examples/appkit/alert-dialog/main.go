// Alert-dialog shows a modal NSAlert with two buttons and prints which was clicked.
package main

import (
	"fmt"

	"github.com/tmc/apple/appkit"
)

func main() {
	appkit.RunApp(func(app appkit.NSApplication, delegate appkit.NSApplicationDelegateObject) {
		alert := appkit.GetNSAlertClass().Alloc().Init()
		alert.SetMessageText("Hello from Go")
		alert.SetInformativeText("This alert was created using Apple framework bindings for Go.")
		alert.AddButtonWithTitle("OK")
		alert.AddButtonWithTitle("Cancel")

		response := alert.RunModal()
		switch response {
		case appkit.AlertFirstButtonReturn:
			fmt.Println("Clicked: OK")
		case appkit.AlertSecondButtonReturn:
			fmt.Println("Clicked: Cancel")
		default:
			fmt.Printf("Clicked: button %d\n", response)
		}

		app.Terminate(app)
	})
}
