// Command hello-window opens a macOS window with one centered label.
//
// Start here if you want the smallest AppKit example in this tree.
package main

import (
	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/corefoundation"
)

func main() {
	appkit.RunApp(func(app appkit.NSApplication, delegate appkit.NSApplicationDelegateObject) {
		window := appkit.GetNSWindowClass().Alloc().InitWithContentRectStyleMaskBackingDefer(
			corefoundation.CGRect{Origin: corefoundation.CGPoint{X: 200, Y: 200}, Size: corefoundation.CGSize{Width: 400, Height: 200}},
			appkit.NSWindowStyleMaskTitled|appkit.NSWindowStyleMaskClosable|appkit.NSWindowStyleMaskMiniaturizable|appkit.NSWindowStyleMaskResizable,
			appkit.NSBackingStoreBuffered,
			false,
		)
		window.SetTitle("Hello")
		window.SetMinSize(corefoundation.CGSize{Width: 200, Height: 100})

		contentView := window.ContentView().(appkit.NSView)

		label := appkit.NewTextFieldLabelWithString("Hello from Go!")
		label.SetFont(appkit.GetNSFontClass().SystemFontOfSize(24))
		label.SetAlignment(appkit.NSTextAlignmentCenter)
		label.SetTranslatesAutoresizingMaskIntoConstraints(false)

		contentView.AddSubview(label)

		label.CenterXAnchor().ConstraintEqualToAnchor(contentView.CenterXAnchor()).SetActive(true)
		label.CenterYAnchor().ConstraintEqualToAnchor(contentView.CenterYAnchor()).SetActive(true)

		window.Center()
		window.MakeKeyAndOrderFront(nil)
		app.Activate()
	})
}
