import ExtensionFoundation
import FSKit
import Foundation

@available(macOS 15.4, *)
@main
struct NinePFSExtension: UnaryFileSystemExtension {
    let fileSystem = NinePFileSystem()
}
