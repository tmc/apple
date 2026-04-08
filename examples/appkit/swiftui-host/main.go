// Command swiftui-host loads a Swift package that defines SwiftUI content and
// hosts it inside an AppKit window.
package main

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/x/swiftui"
)

var (
	window appkit.NSWindow
	host   *swiftui.HostedViewController
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
		WillTerminate: func(_ foundation.NSNotification) {
			if host != nil {
				host.Release()
				host = nil
			}
		},
	})
	app.SetDelegate(delegate)
	app.Run()
}

func buildWindow(app appkit.NSApplication) {
	window = appkit.NewWindowWithContentRectStyleMaskBackingDefer(
		corefoundation.CGRect{
			Origin: corefoundation.CGPoint{X: 0, Y: 0},
			Size:   corefoundation.CGSize{Width: 720, Height: 420},
		},
		appkit.NSWindowStyleMaskTitled|
			appkit.NSWindowStyleMaskClosable|
			appkit.NSWindowStyleMaskMiniaturizable|
			appkit.NSWindowStyleMaskResizable,
		appkit.NSBackingStoreBuffered,
		false,
	)
	window.SetTitle("SwiftUI Host")
	window.SetMinSize(corefoundation.CGSize{Width: 560, Height: 320})
	window.Center()

	bridge, err := swiftui.Open(swiftui.Config{
		Dir:             uiPackageDir(),
		Product:         "ExampleAppUI",
		FactoryClass:    "ExampleAppUIBridge",
		EmbeddedLibrary: embeddedExampleAppUI(),
	})
	if err != nil {
		log.Fatalf("open swift package: %v", err)
	}
	host, err = bridge.NewViewController()
	if err != nil {
		log.Fatalf("create hosting controller: %v", err)
	}

	window.SetContentViewController(host)
	window.MakeKeyAndOrderFront(nil)
	app.Activate()
}

func uiPackageDir() string {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("resolve example directory")
	}
	return filepath.Join(filepath.Dir(file), "ui")
}
