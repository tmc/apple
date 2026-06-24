@preconcurrency import FSKit
import SwiftUI

@MainActor
final class ModuleList: ObservableObject {
    struct Row: Identifiable, Sendable {
        let id: String
        let enabled: Bool
        let path: String
    }

    @Published var modules: [Row] = []
    @Published var error: String?

    func refresh() {
        FSClient.shared.fetchInstalledExtensions { modules, error in
            let rows = (modules ?? []).map { module in
                Row(id: module.bundleIdentifier, enabled: module.isEnabled, path: module.url.path)
            }
            DispatchQueue.main.async {
                if let error {
                    self.error = String(describing: error)
                    print("fskit: fetch installed extensions: \(error)")
                    return
                }
                self.modules = rows
                self.error = nil
                for module in rows {
                    print("fskit: \(module.id) enabled=\(module.enabled) url=\(module.path)")
                }
            }
        }
    }
}

@main
struct TinyFSHostApp: App {
    @StateObject private var moduleList = ModuleList()

    init() {
        if CommandLine.arguments.contains("--fskit-probe") {
            let semaphore = DispatchSemaphore(value: 0)
            FSClient.shared.fetchInstalledExtensions { modules, error in
                if let error {
                    print("fskit: fetch installed extensions: \(error)")
                }
                for module in modules ?? [] {
                    print("fskit: \(module.bundleIdentifier) enabled=\(module.isEnabled) url=\(module.url.path)")
                }
                semaphore.signal()
            }
            _ = semaphore.wait(timeout: .now() + 10)
            exit(0)
        }
    }

    var body: some Scene {
        WindowGroup {
            VStack(alignment: .leading, spacing: 12) {
                Text("TinyFS")
                    .font(.title)
                Text("The FSKit module lives in the bundled app extension.")
                if let error = moduleList.error {
                    Text(error)
                        .foregroundStyle(.red)
                }
                List(moduleList.modules) { module in
                    VStack(alignment: .leading) {
                        Text(module.id)
                        Text("\(module.enabled ? "enabled" : "disabled") - \(module.path)")
                            .font(.caption)
                            .foregroundStyle(.secondary)
                    }
                }
                .frame(minHeight: 180)
            }
            .padding(24)
            .frame(minWidth: 620, minHeight: 320)
            .task {
                moduleList.refresh()
            }
        }
    }
}
