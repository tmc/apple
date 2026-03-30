// Command menubar-clock shows the current time in the macOS menu bar.
package main

import (
	"runtime"
	"time"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/objectivec"
)

func main() {
	appkit.RunApp(func(app appkit.NSApplication, delegate appkit.NSApplicationDelegateObject) {
		app.SetActivationPolicy(appkit.NSApplicationActivationPolicyAccessory)

		statusBar := appkit.GetNSStatusBarClass().SystemStatusBar()
		item := statusBar.StatusItemWithLength(appkit.VariableStatusItemLength)
		runtime.KeepAlive(item)

		button := item.Button()
		button.SetTitle(time.Now().Format("15:04:05"))

		menu := appkit.GetNSMenuClass().Alloc().Init()
		quit := appkit.GetNSMenuItemClass().Alloc().Init()
		quit.SetTitle("Quit")
		quit.SetActionHandler(func() {
			app.Terminate(objectivec.Object{})
		})
		menu.AddItem(quit)
		item.(appkit.NSStatusItem).SetMenu(menu)

		mainQ := dispatch.MainQueue()
		go func() {
			ticker := time.NewTicker(time.Second)
			defer ticker.Stop()
			for range ticker.C {
				mainQ.Async(func() {
					button.SetTitle(time.Now().Format("15:04:05"))
				})
			}
		}()
	})
}
