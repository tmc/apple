import ExtensionFoundation
import Foundation
import FSKit

final class TinyItem: FSItem {
    let itemID: FSItem.Identifier
    let parentID: FSItem.Identifier
    let type: FSItem.ItemType

    init(itemID: FSItem.Identifier, parentID: FSItem.Identifier, type: FSItem.ItemType) {
        self.itemID = itemID
        self.parentID = parentID
        self.type = type
        super.init()
    }
}

final class TinyVolume: FSVolume, FSVolume.Operations {
    private let rootItem = TinyItem(
        itemID: .rootDirectory,
        parentID: .parentOfRoot,
        type: .directory
    )

    init() {
        super.init(
            volumeID: FSVolume.Identifier(uuid: UUID()),
            volumeName: FSFileName(string: "TinyFS")
        )
    }

    var maximumLinkCount: Int { 1 }
    var maximumNameLength: Int { 255 }
    var restrictsOwnershipChanges: Bool { false }
    var truncatesLongNames: Bool { false }

    var supportedVolumeCapabilities: FSVolume.SupportedCapabilities {
        let capabilities = FSVolume.SupportedCapabilities()
        capabilities.supports64BitObjectIDs = true
        capabilities.supportsFastStatFS = true
        capabilities.supportsHiddenFiles = true
        capabilities.doesNotSupportSettingFilePermissions = true
        capabilities.caseFormat = .sensitive
        return capabilities
    }

    var volumeStatistics: FSStatFSResult {
        let result = FSStatFSResult(fileSystemTypeName: "tinyfs")
        result.blockSize = 4096
        result.ioSize = 4096
        result.totalBlocks = 1
        result.availableBlocks = 0
        result.freeBlocks = 0
        result.usedBlocks = 1
        result.totalFiles = 1
        result.freeFiles = 0
        return result
    }

    func mount(options: FSTaskOptions, replyHandler reply: @escaping (Error?) -> Void) {
        reply(nil)
    }

    func unmount(replyHandler reply: @escaping () -> Void) {
        reply()
    }

    func synchronize(flags: FSSyncFlags, replyHandler reply: @escaping (Error?) -> Void) {
        reply(nil)
    }

    func getAttributes(
        _ desiredAttributes: FSItem.GetAttributesRequest,
        of item: FSItem,
        replyHandler reply: @escaping (FSItem.Attributes?, Error?) -> Void
    ) {
        guard let tinyItem = item as? TinyItem else {
            return reply(nil, POSIXError(.EINVAL))
        }
        reply(attributes(for: tinyItem), nil)
    }

    func setAttributes(
        _ newAttributes: FSItem.SetAttributesRequest,
        on item: FSItem,
        replyHandler reply: @escaping (FSItem.Attributes?, Error?) -> Void
    ) {
        reply(nil, POSIXError(.EROFS))
    }

    func lookupItem(
        named name: FSFileName,
        inDirectory directory: FSItem,
        replyHandler reply: @escaping (FSItem?, FSFileName?, Error?) -> Void
    ) {
        reply(nil, nil, POSIXError(.ENOENT))
    }

    func reclaimItem(_ item: FSItem, replyHandler reply: @escaping (Error?) -> Void) {
        reply(nil)
    }

    func readSymbolicLink(_ item: FSItem, replyHandler reply: @escaping (FSFileName?, Error?) -> Void) {
        reply(nil, POSIXError(.EINVAL))
    }

    func createItem(
        named name: FSFileName,
        type: FSItem.ItemType,
        inDirectory directory: FSItem,
        attributes newAttributes: FSItem.SetAttributesRequest,
        replyHandler reply: @escaping (FSItem?, FSFileName?, Error?) -> Void
    ) {
        reply(nil, nil, POSIXError(.EROFS))
    }

    func createSymbolicLink(
        named name: FSFileName,
        inDirectory directory: FSItem,
        attributes newAttributes: FSItem.SetAttributesRequest,
        linkContents contents: FSFileName,
        replyHandler reply: @escaping (FSItem?, FSFileName?, Error?) -> Void
    ) {
        reply(nil, nil, POSIXError(.EROFS))
    }

    func createLink(
        to item: FSItem,
        named name: FSFileName,
        inDirectory directory: FSItem,
        replyHandler reply: @escaping (FSFileName?, Error?) -> Void
    ) {
        reply(nil, POSIXError(.EROFS))
    }

    func removeItem(
        _ item: FSItem,
        named name: FSFileName,
        fromDirectory directory: FSItem,
        replyHandler reply: @escaping (Error?) -> Void
    ) {
        reply(POSIXError(.EROFS))
    }

    func renameItem(
        _ item: FSItem,
        inDirectory sourceDirectory: FSItem,
        named sourceName: FSFileName,
        to destinationName: FSFileName,
        inDirectory destinationDirectory: FSItem,
        overItem: FSItem?,
        replyHandler reply: @escaping (FSFileName?, Error?) -> Void
    ) {
        reply(nil, POSIXError(.EROFS))
    }

    func enumerateDirectory(
        _ directory: FSItem,
        startingAt cookie: FSDirectoryCookie,
        verifier: FSDirectoryVerifier,
        attributes: FSItem.GetAttributesRequest?,
        packer: FSDirectoryEntryPacker,
        replyHandler reply: @escaping (FSDirectoryVerifier, Error?) -> Void
    ) {
        reply(verifier, nil)
    }

    func activate(options: FSTaskOptions, replyHandler reply: @escaping (FSItem?, Error?) -> Void) {
        reply(rootItem, nil)
    }

    func deactivate(options: FSDeactivateOptions, replyHandler reply: @escaping (Error?) -> Void) {
        reply(nil)
    }

    private func attributes(for item: TinyItem) -> FSItem.Attributes {
        let attributes = FSItem.Attributes()
        attributes.type = item.type
        attributes.mode = item.type == .directory ? 0o555 : 0o444
        attributes.linkCount = 1
        attributes.uid = getuid()
        attributes.gid = getgid()
        attributes.size = 0
        attributes.allocSize = 0
        attributes.fileID = item.itemID
        attributes.parentID = item.parentID
        return attributes
    }
}

final class TinyFileSystem: FSUnaryFileSystem, FSUnaryFileSystemOperations {
    private var volume: TinyVolume?

    func probeResource(resource: FSResource, replyHandler reply: @escaping (FSProbeResult?, Error?) -> Void) {
        reply(
            FSProbeResult.usable(
                name: "TinyFS",
                containerID: FSContainerIdentifier(uuid: UUID())
            ),
            nil
        )
    }

    func loadResource(
        resource: FSResource,
        options: FSTaskOptions,
        replyHandler reply: @escaping (FSVolume?, Error?) -> Void
    ) {
        containerStatus = .ready
        let volume = TinyVolume()
        self.volume = volume
        reply(volume, nil)
    }

    func unloadResource(
        resource: FSResource,
        options: FSTaskOptions,
        replyHandler reply: @escaping (Error?) -> Void
    ) {
        volume = nil
        reply(nil)
    }
}

@main
struct TinyFSExtension: UnaryFileSystemExtension {
    let fileSystem = TinyFileSystem()
}
