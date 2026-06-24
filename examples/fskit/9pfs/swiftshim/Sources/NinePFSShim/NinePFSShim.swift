import ExtensionFoundation
import FSKit
import NinePFSShimObjC

@available(macOS 15.4, *)
struct NinePFSExtension: UnaryFileSystemExtension {
    let fileSystem = NinePFileSystem()
}

@available(macOS 15.4, *)
@_cdecl("NinePFSShimHasFileSystemClass")
public func NinePFSShimHasFileSystemClass() -> Bool {
    NSClassFromString("NinePFileSystem") != nil
}

@available(macOS 15.4, *)
@MainActor
@_cdecl("NinePFSRunExtensionMain")
public func NinePFSRunExtensionMain() {
    try! NinePFSExtension.main()
}
