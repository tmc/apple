import ExtensionFoundation
import FSKit
import TinyFSShimObjC

@available(macOS 15.4, *)
struct TinyFSExtension: UnaryFileSystemExtension {
    let fileSystem = GoTinyFSFileSystem()
}

@available(macOS 15.4, *)
@_cdecl("TinyFSShimHasFileSystemClass")
public func TinyFSShimHasFileSystemClass() -> Bool {
    NSClassFromString("GoTinyFSFileSystem") != nil
}

@available(macOS 15.4, *)
@MainActor
@_cdecl("TinyFSRunExtensionMain")
public func TinyFSRunExtensionMain() {
    try! TinyFSExtension.main()
}
