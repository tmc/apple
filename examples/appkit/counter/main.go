// Counter displays a window with a count label and buttons to increment or reset.
package main

import (
	"fmt"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/corefoundation"
)

func main() {
	appkit.RunApp(func(app appkit.NSApplication, delegate appkit.NSApplicationDelegateObject) {
		window := appkit.GetNSWindowClass().Alloc().InitWithContentRectStyleMaskBackingDefer(
			corefoundation.CGRect{Origin: corefoundation.CGPoint{X: 200, Y: 200}, Size: corefoundation.CGSize{Width: 300, Height: 200}},
			appkit.NSWindowStyleMaskTitled|appkit.NSWindowStyleMaskClosable|appkit.NSWindowStyleMaskMiniaturizable,
			appkit.NSBackingStoreBuffered,
			false,
		)
		window.SetTitle("Counter")

		count := 0

		label := appkit.GetNSTextFieldClass().Alloc().Init()
		label.SetStringValue("0")
		label.SetEditable(false)
		label.SetBezeled(false)
		label.SetDrawsBackground(false)
		label.SetFont(appkit.GetNSFontClass().SystemFontOfSize(48))
		label.SetAlignment(appkit.NSTextAlignmentCenter)

		increment := appkit.GetNSButtonClass().Alloc().Init()
		increment.SetTitle("+1")
		increment.SetBezelStyle(appkit.NSBezelStyleAutomatic)
		increment.SetActionHandler(func() {
			count++
			label.SetStringValue(fmt.Sprintf("%d", count))
		})

		reset := appkit.GetNSButtonClass().Alloc().Init()
		reset.SetTitle("Reset")
		reset.SetBezelStyle(appkit.NSBezelStyleAutomatic)
		reset.SetActionHandler(func() {
			count = 0
			label.SetStringValue("0")
		})

		stack := appkit.GetNSStackViewClass().Alloc().Init()
		stack.SetOrientation(appkit.NSUserInterfaceLayoutOrientationVertical)
		stack.SetSpacing(10)
		stack.AddArrangedSubview(label)
		stack.AddArrangedSubview(increment)
		stack.AddArrangedSubview(reset)

		window.SetContentView(stack)
		window.Center()
		window.MakeKeyAndOrderFront(app)
	})
}
