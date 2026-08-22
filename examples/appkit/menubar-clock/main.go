// Command menubar-clock shows the current time in the macOS menu bar.
package main

import (
	"time"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/objectivec"
)

func main() {
	appkit.RunApp(func(app appkit.NSApplication, _ appkit.NSApplicationDelegateObject) {
		app.SetActivationPolicy(appkit.NSApplicationActivationPolicyAccessory)
		item := appkit.GetNSStatusBarClass().SystemStatusBar().StatusItemWithLength(appkit.VariableStatusItemLength)
		// NSStatusBar does not retain the item. Keep one native reference until Quit.
		itemObject := objectivec.Object{ID: item.GetID()}
		itemObject.Retain()
		button := item.Button()
		button.SetTitle(time.Now().Format("15:04:05"))

		menu := appkit.NewNSMenu()
		// The 0 action selector means none; SetActionHandler attaches a Go
		// closure immediately below.
		quit := appkit.NewMenuItemWithTitleActionKeyEquivalent("Quit", 0, "q")
		quit.SetActionHandler(func() {
			itemObject.Release()
			app.Terminate(nil)
		})
		menu.AddItem(quit)
		item.SetMenu(menu)

		mainQueue := dispatch.MainQueue()
		go func() {
			ticker := time.NewTicker(time.Second)
			defer ticker.Stop()
			for range ticker.C {
				mainQueue.Async(func() {
					button.SetTitle(time.Now().Format("15:04:05"))
				})
			}
		}()
	})
}
