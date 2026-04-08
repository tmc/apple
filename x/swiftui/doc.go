// Package swiftui loads app-owned Swift packages that expose SwiftUI content to
// AppKit.
//
// The intended workflow is:
//
//  1. Put your SwiftUI code in a Swift package that builds a dynamic library.
//  2. Export an Objective-C-visible factory class from that package.
//  3. Optionally prebuild a universal dylib with go generate.
//  4. Optionally embed that dylib in your Go binary.
//  5. Use this Go package to load the embedded dylib first, then
//     `Dir/lib<Product>.dylib`, then fall back to building the Swift package
//     on demand from source.
//
// The default contract expects the Swift factory class to provide
// `newRootViewController` and `newRootView` class methods. Those selectors use
// Cocoa's "new" ownership convention, so the Go caller owns the returned
// object and must call Release when finished with it.
//
// Example Swift package:
//
//	import SwiftUI
//	import AppKit
//
//	struct ContentView: View {
//	    var body: some View {
//	        Text("Hello from SwiftUI").padding(24)
//	    }
//	}
//
//	@objc(ExampleUIBridge)
//	public final class ExampleUIBridge: NSObject {
//	    @objc public static func newRootViewController() -> NSViewController {
//	        NSHostingController(rootView: ContentView())
//	    }
//
//	    @objc public static func newRootView() -> NSView {
//	        NSHostingView(rootView: ContentView())
//	    }
//	}
//
// Example Go usage:
//
//	bridge, err := swiftui.Open(swiftui.Config{
//		Dir:          "/path/to/ui",
//		Product:      "ExampleUI",
//		FactoryClass: "ExampleUIBridge",
//		EmbeddedLibrary: embeddedExampleUI(),
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//	controller, err := bridge.NewViewController()
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer controller.Release()
//	window.SetContentViewController(controller)
//
// Typical embedding setup:
//
//	//go:generate go run ./ui/generate.go
//
//	//go:build swiftui_embed
//	//go:embed ui/libExampleUI.dylib
//	var exampleUIDylib []byte
//
// Callers should lock the main goroutine to the main thread before creating or
// using AppKit objects returned by this package.
package swiftui
