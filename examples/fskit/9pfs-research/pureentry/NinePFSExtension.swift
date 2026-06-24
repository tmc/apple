import ExtensionFoundation
import FSKit

@_silgen_name("NinePFSDelayedLoadResource")
func NinePFSDelayedLoadResource(_ selfPtr: UnsafeMutableRawPointer?, _ resource: FSResource, _ options: FSTaskOptions, _ reply: @escaping @convention(block) (FSVolume?, NSError?) -> Void)

@available(macOS 15.4, *)
@objc(NinePFileSystem)
public class NinePFileSystem: FSUnaryFileSystem, FSUnaryFileSystemOperations {
    public func probeResource(resource: FSResource, replyHandler: @Sendable @escaping (FSProbeResult?, Error?) -> Void) {
        let result = FSProbeResult.usable(name: "9pfs", containerID: FSContainerIdentifier())
        replyHandler(result, nil)
    }

    public func loadResource(resource: FSResource, options: FSTaskOptions, replyHandler: @Sendable @escaping (FSVolume?, Error?) -> Void) {
        let replyBlock: @convention(block) (FSVolume?, NSError?) -> Void = { volume, error in
            replyHandler(volume, error)
        }
        NinePFSDelayedLoadResource(Unmanaged.passUnretained(self).toOpaque(), resource, options, replyBlock)
    }

    public func unloadResource(resource: FSResource, options: FSTaskOptions, replyHandler: @Sendable @escaping (Error?) -> Void) {
        replyHandler(nil)
    }
}

@available(macOS 15.4, *)
@main
public struct NinePFSExtension: UnaryFileSystemExtension {
    public init() {}
    public let fileSystem = NinePFileSystem()
}
