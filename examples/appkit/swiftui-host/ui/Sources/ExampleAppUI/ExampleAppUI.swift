import AppKit
import SwiftUI

struct ContentView: View {
    @State private var count = 0

    var body: some View {
        VStack(spacing: 16) {
            Text("SwiftUI from a Swift package")
                .font(.title)
            Text("This view tree lives in Swift. Go loads the package and hosts it in AppKit.")
                .multilineTextAlignment(.center)
                .foregroundStyle(.secondary)
            Text("\(count)")
                .font(.system(size: 42, weight: .semibold, design: .rounded))
            HStack(spacing: 12) {
                Button("Subtract") { count -= 1 }
                Button("Add") { count += 1 }
            }
        }
        .padding(24)
        .frame(maxWidth: .infinity, maxHeight: .infinity)
    }
}

@objc(ExampleAppUIBridge)
public final class ExampleAppUIBridge: NSObject {
    @objc public static func newRootViewController() -> NSViewController {
        NSHostingController(rootView: ContentView())
    }

    @objc public static func newRootView() -> NSView {
        NSHostingView(rootView: ContentView())
    }
}
