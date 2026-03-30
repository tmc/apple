// Command counter displays a small AppKit window with a label and two buttons.
//
// It is a compact follow-up to hello-window: same basic window setup, plus
// button handlers and a vertical stack layout.
package main

import (
	"runtime"
	"strconv"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

func init() { runtime.LockOSThread() }

func main() {
	app := appkit.GetNSApplicationClass().SharedApplication()
	app.SetActivationPolicy(appkit.NSApplicationActivationPolicyRegular)

	delegate := appkit.NewNSApplicationDelegate(appkit.NSApplicationDelegateConfig{
		ShouldTerminateAfterLastWindowClosed: func(_ appkit.NSApplication) bool { return true },
		DidFinishLaunching: func(_ foundation.NSNotification) {
			buildWindow(app)
		},
	})
	app.SetDelegate(delegate)
	app.Run()
}

func buildWindow(app appkit.NSApplication) {
	window := appkit.GetNSWindowClass().Alloc().InitWithContentRectStyleMaskBackingDefer(
		corefoundation.CGRect{Origin: corefoundation.CGPoint{X: 200, Y: 200}, Size: corefoundation.CGSize{Width: 300, Height: 200}},
		appkit.NSWindowStyleMaskTitled|appkit.NSWindowStyleMaskClosable|appkit.NSWindowStyleMaskMiniaturizable|appkit.NSWindowStyleMaskResizable,
		appkit.NSBackingStoreBuffered,
		false,
	)
	window.SetTitle("Counter")
	window.SetMinSize(corefoundation.CGSize{Width: 200, Height: 150})

	count := 0
	label := appkit.NewTextFieldLabelWithString("0")
	label.SetFont(appkit.GetNSFontClass().SystemFontOfSize(48))
	label.SetAlignment(appkit.NSTextAlignmentCenter)

	updateCount := func() {
		label.SetStringValue(strconv.Itoa(count))
	}

	increment := appkit.NewButtonWithTitleTargetAction("+1", nil, 0)
	increment.SetBezelStyle(appkit.NSBezelStyleAutomatic)
	increment.SetActionHandler(func() {
		count++
		updateCount()
	})

	reset := appkit.NewButtonWithTitleTargetAction("Reset", nil, 0)
	reset.SetBezelStyle(appkit.NSBezelStyleAutomatic)
	reset.SetActionHandler(func() {
		count = 0
		updateCount()
	})

	stack := appkit.GetNSStackViewClass().Alloc().Init()
	stack.SetOrientation(appkit.NSUserInterfaceLayoutOrientationVertical)
	stack.SetAlignment(appkit.NSLayoutAttributeCenterX)
	stack.SetSpacing(12)
	stack.SetEdgeInsets(foundation.NSEdgeInsets{Top: 20, Left: 20, Bottom: 20, Right: 20})
	stack.AddArrangedSubview(label)
	stack.AddArrangedSubview(increment)
	stack.AddArrangedSubview(reset)
	stack.SetTranslatesAutoresizingMaskIntoConstraints(false)

	contentView := window.ContentView().(appkit.NSView)
	contentView.AddSubview(stack)

	stack.LeadingAnchor().ConstraintEqualToAnchor(contentView.LeadingAnchor()).SetActive(true)
	stack.TrailingAnchor().ConstraintEqualToAnchor(contentView.TrailingAnchor()).SetActive(true)
	stack.TopAnchor().ConstraintEqualToAnchor(contentView.TopAnchor()).SetActive(true)
	stack.BottomAnchor().ConstraintEqualToAnchor(contentView.BottomAnchor()).SetActive(true)

	window.Center()
	window.MakeKeyAndOrderFront(nil)
	app.Activate()
}
