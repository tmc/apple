// Hello-window opens a macOS window displaying "Hello from Go!" in large text.
package main

import (
	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/corefoundation"
)

func main() {
	appkit.RunApp(func(app appkit.NSApplication, delegate appkit.NSApplicationDelegateObject) {
		window := appkit.GetNSWindowClass().Alloc().InitWithContentRectStyleMaskBackingDefer(
			corefoundation.CGRect{Origin: corefoundation.CGPoint{X: 200, Y: 200}, Size: corefoundation.CGSize{Width: 400, Height: 200}},
			appkit.NSWindowStyleMaskTitled|appkit.NSWindowStyleMaskClosable|appkit.NSWindowStyleMaskMiniaturizable,
			appkit.NSBackingStoreBuffered,
			false,
		)
		window.SetTitle("Hello")

		label := appkit.GetNSTextFieldClass().Alloc().Init()
		label.SetStringValue("Hello from Go!")
		label.SetEditable(false)
		label.SetBezeled(false)
		label.SetDrawsBackground(false)
		label.SetFont(appkit.GetNSFontClass().SystemFontOfSize(24))
		label.SetAlignment(appkit.NSTextAlignmentCenter)

		window.SetContentView(label)
		window.Center()
		window.MakeKeyAndOrderFront(app)
	})
}
