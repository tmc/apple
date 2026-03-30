// Command file-picker opens an NSOpenPanel and prints the selected paths.
package main

import (
	"fmt"

	"github.com/tmc/apple/appkit"
)

func main() {
	appkit.RunApp(func(app appkit.NSApplication, delegate appkit.NSApplicationDelegateObject) {
		panel := appkit.NewNSOpenPanel()
		panel.SetCanChooseFiles(true)
		panel.SetCanChooseDirectories(false)
		panel.SetAllowsMultipleSelection(true)
		panel.SetMessage("Select one or more files")
		panel.SetPrompt("Open")

		response := panel.RunModal()
		if response == appkit.NSModalResponses.OK {
			urls := panel.URLs()
			fmt.Printf("Selected %d file(s):\n", len(urls))
			for _, u := range urls {
				fmt.Println(u.Path())
			}
		} else {
			fmt.Println("Cancelled")
		}

		app.Terminate(app)
	})
}
