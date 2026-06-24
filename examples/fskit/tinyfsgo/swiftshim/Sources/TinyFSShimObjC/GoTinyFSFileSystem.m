#import "GoTinyFSFileSystem.h"
#import <errno.h>
#import <stdio.h>
#import <unistd.h>

static TinyFSProbeResourceCallback tinyFSProbeResourceCallback;
static TinyFSLoadResourceCallback tinyFSLoadResourceCallback;
static TinyFSUnloadResourceCallback tinyFSUnloadResourceCallback;
static TinyFSDidFinishLoadingCallback tinyFSDidFinishLoadingCallback;
static TinyFSProbeResourceResultCallback tinyFSProbeResourceResultCallback;
static TinyFSLoadResourceResultCallback tinyFSLoadResourceResultCallback;
static TinyFSUnloadResourceResultCallback tinyFSUnloadResourceResultCallback;
static TinyFSVolumeCallback tinyFSVolumeActivateCallback;
static TinyFSVolumeCallback tinyFSVolumeMountCallback;
static TinyFSVolumeCallback tinyFSVolumeUnmountCallback;

static void
TinyFSShimMarker(const char *name)
{
	FILE *f = fopen(name, "w");
	if (f == NULL) {
		return;
	}
	fputs(name, f);
	fputc('\n', f);
	fclose(f);
}

static NSError *
TinyFSShimPOSIXError(NSInteger code)
{
	return [NSError errorWithDomain:NSPOSIXErrorDomain code:code userInfo:nil];
}

static FSItemAttributes *
TinyFSShimAttributesForItem(GoTinyFSItem *item)
{
	FSItemAttributes *attributes = [FSItemAttributes new];
	attributes.type = item.type;
	attributes.mode = item.type == FSItemTypeDirectory ? 0555 : 0444;
	attributes.linkCount = 1;
	attributes.uid = getuid();
	attributes.gid = getgid();
	attributes.size = 0;
	attributes.allocSize = 0;
	attributes.fileID = item.itemID;
	attributes.parentID = item.parentID;
	return attributes;
}

void
TinyFSShimSetProbeResourceCallback(TinyFSProbeResourceCallback callback)
{
	tinyFSProbeResourceCallback = callback;
}

void
TinyFSShimSetLoadResourceCallback(TinyFSLoadResourceCallback callback)
{
	tinyFSLoadResourceCallback = callback;
}

void
TinyFSShimSetUnloadResourceCallback(TinyFSUnloadResourceCallback callback)
{
	tinyFSUnloadResourceCallback = callback;
}

void
TinyFSShimSetDidFinishLoadingCallback(TinyFSDidFinishLoadingCallback callback)
{
	tinyFSDidFinishLoadingCallback = callback;
}

void
TinyFSShimSetProbeResourceResultCallback(TinyFSProbeResourceResultCallback callback)
{
	tinyFSProbeResourceResultCallback = callback;
}

void
TinyFSShimSetLoadResourceResultCallback(TinyFSLoadResourceResultCallback callback)
{
	tinyFSLoadResourceResultCallback = callback;
}

void
TinyFSShimSetUnloadResourceResultCallback(TinyFSUnloadResourceResultCallback callback)
{
	tinyFSUnloadResourceResultCallback = callback;
}

void
TinyFSShimSetVolumeActivateCallback(TinyFSVolumeCallback callback)
{
	tinyFSVolumeActivateCallback = callback;
}

void
TinyFSShimSetVolumeMountCallback(TinyFSVolumeCallback callback)
{
	tinyFSVolumeMountCallback = callback;
}

void
TinyFSShimSetVolumeUnmountCallback(TinyFSVolumeCallback callback)
{
	tinyFSVolumeUnmountCallback = callback;
}

void
TinyFSShimReplyProbe(id reply, id result, NSError *error)
{
	void (^block)(FSProbeResult *_Nullable, NSError *_Nullable) = reply;
	block(result, error);
}

void
TinyFSShimReplyVolume(id reply, id volume, NSError *error)
{
	void (^block)(FSVolume *_Nullable, NSError *_Nullable) = reply;
	block(volume, error);
}

void
TinyFSShimReplyError(id reply, NSError *error)
{
	void (^block)(NSError *_Nullable) = reply;
	block(error);
}

id
TinyFSShimNewVolume(void)
{
	FSVolumeIdentifier *volumeID = [[FSVolumeIdentifier alloc] initWithUUID:[NSUUID UUID]];
	FSFileName *volumeName = [[FSFileName alloc] initWithString:@"TinyFS"];
	return [[GoTinyFSVolume alloc] initWithVolumeID:volumeID volumeName:volumeName];
}

@implementation GoTinyFSItem
@end

@implementation GoTinyFSVolume {
	GoTinyFSItem *_rootItem;
}

- (instancetype)initWithVolumeID:(FSVolumeIdentifier *)volumeID volumeName:(FSFileName *)volumeName
{
	self = [super initWithVolumeID:volumeID volumeName:volumeName];
	if (self != nil) {
		_rootItem = [GoTinyFSItem new];
		_rootItem.itemID = FSItemIDRootDirectory;
		_rootItem.parentID = FSItemIDParentOfRoot;
		_rootItem.type = FSItemTypeDirectory;
	}
	return self;
}

- (NSInteger)maximumLinkCount { return 1; }
- (NSInteger)maximumNameLength { return 255; }
- (uint64_t)maximumFileSize { return 0xffffffff; }
- (NSInteger)maximumFileSizeInBits { return 33; }
- (BOOL)restrictsOwnershipChanges { return NO; }
- (BOOL)truncatesLongNames { return NO; }
- (BOOL)enableOpenUnlinkEmulation { return NO; }
- (void)setEnableOpenUnlinkEmulation:(BOOL)value {}

- (FSVolumeSupportedCapabilities *)supportedVolumeCapabilities
{
	FSVolumeSupportedCapabilities *capabilities = [FSVolumeSupportedCapabilities new];
	capabilities.supports64BitObjectIDs = YES;
	capabilities.supportsFastStatFS = YES;
	capabilities.supportsHiddenFiles = YES;
	capabilities.doesNotSupportSettingFilePermissions = YES;
	capabilities.caseFormat = FSVolumeCaseFormatSensitive;
	return capabilities;
}

- (FSStatFSResult *)volumeStatistics
{
	FSStatFSResult *result = [[FSStatFSResult alloc] initWithFileSystemTypeName:@"tinyfs"];
	result.blockSize = 4096;
	result.ioSize = 4096;
	result.totalBlocks = 1;
	result.availableBlocks = 0;
	result.freeBlocks = 0;
	result.usedBlocks = 1;
	result.totalFiles = 1;
	result.freeFiles = 0;
	return result;
}

- (void)activateWithOptions:(FSTaskOptions *)options replyHandler:(void (^)(FSItem *_Nullable, NSError *_Nullable))reply
{
	TinyFSShimMarker("tinyfsgo-objc-activate.marker");
	if (tinyFSVolumeActivateCallback != NULL) {
		tinyFSVolumeActivateCallback(self, _cmd);
	}
	reply(_rootItem, nil);
}

- (void)deactivateWithOptions:(FSDeactivateOptions)options replyHandler:(void (^)(NSError *_Nullable))reply
{
	reply(nil);
}

- (void)mountWithOptions:(FSTaskOptions *)options replyHandler:(void (^)(NSError *_Nullable))reply
{
	TinyFSShimMarker("tinyfsgo-objc-mount.marker");
	if (tinyFSVolumeMountCallback != NULL) {
		tinyFSVolumeMountCallback(self, _cmd);
	}
	reply(nil);
}

- (void)unmountWithReplyHandler:(void (^)(void))reply
{
	if (tinyFSVolumeUnmountCallback != NULL) {
		tinyFSVolumeUnmountCallback(self, _cmd);
	}
	reply();
}

- (void)synchronizeWithFlags:(FSSyncFlags)flags replyHandler:(void (^)(NSError *_Nullable))reply
{
	reply(nil);
}

- (void)getAttributes:(FSItemGetAttributesRequest *)desiredAttributes ofItem:(FSItem *)item replyHandler:(void (^)(FSItemAttributes *_Nullable, NSError *_Nullable))reply
{
	if (![item isKindOfClass:[GoTinyFSItem class]]) {
		reply(nil, TinyFSShimPOSIXError(EINVAL));
		return;
	}
	reply(TinyFSShimAttributesForItem((GoTinyFSItem *)item), nil);
}

- (void)setAttributes:(FSItemSetAttributesRequest *)newAttributes onItem:(FSItem *)item replyHandler:(void (^)(FSItemAttributes *_Nullable, NSError *_Nullable))reply
{
	reply(nil, TinyFSShimPOSIXError(EROFS));
}

- (void)lookupItemNamed:(FSFileName *)name inDirectory:(FSItem *)directory replyHandler:(void (^)(FSItem *_Nullable, FSFileName *_Nullable, NSError *_Nullable))reply
{
	reply(nil, nil, TinyFSShimPOSIXError(ENOENT));
}

- (void)reclaimItem:(FSItem *)item replyHandler:(void (^)(NSError *_Nullable))reply
{
	reply(nil);
}

- (void)readSymbolicLink:(FSItem *)item replyHandler:(void (^)(FSFileName *_Nullable, NSError *_Nullable))reply
{
	reply(nil, TinyFSShimPOSIXError(EINVAL));
}

- (void)createItemNamed:(FSFileName *)name type:(FSItemType)type inDirectory:(FSItem *)directory attributes:(FSItemSetAttributesRequest *)newAttributes replyHandler:(void (^)(FSItem *_Nullable, FSFileName *_Nullable, NSError *_Nullable))reply
{
	reply(nil, nil, TinyFSShimPOSIXError(EROFS));
}

- (void)createSymbolicLinkNamed:(FSFileName *)name inDirectory:(FSItem *)directory attributes:(FSItemSetAttributesRequest *)newAttributes linkContents:(FSFileName *)contents replyHandler:(void (^)(FSItem *_Nullable, FSFileName *_Nullable, NSError *_Nullable))reply
{
	reply(nil, nil, TinyFSShimPOSIXError(EROFS));
}

- (void)createLinkToItem:(FSItem *)item named:(FSFileName *)name inDirectory:(FSItem *)directory replyHandler:(void (^)(FSFileName *_Nullable, NSError *_Nullable))reply
{
	reply(nil, TinyFSShimPOSIXError(EROFS));
}

- (void)removeItem:(FSItem *)item named:(FSFileName *)name fromDirectory:(FSItem *)directory replyHandler:(void (^)(NSError *_Nullable))reply
{
	reply(TinyFSShimPOSIXError(EROFS));
}

- (void)renameItem:(FSItem *)item inDirectory:(FSItem *)sourceDirectory named:(FSFileName *)sourceName toNewName:(FSFileName *)destinationName inDirectory:(FSItem *)destinationDirectory overItem:(FSItem *_Nullable)overItem replyHandler:(void (^)(FSFileName *_Nullable, NSError *_Nullable))reply
{
	reply(nil, TinyFSShimPOSIXError(EROFS));
}

- (void)enumerateDirectory:(FSItem *)directory startingAtCookie:(FSDirectoryCookie)cookie verifier:(FSDirectoryVerifier)verifier providingAttributes:(FSItemGetAttributesRequest *_Nullable)attributes usingPacker:(FSDirectoryEntryPacker *)packer replyHandler:(void (^)(FSDirectoryVerifier, NSError *_Nullable))reply
{
	reply(verifier, nil);
}

@end

@implementation GoTinyFSFileSystem

- (void)didFinishLoading
{
	if (tinyFSDidFinishLoadingCallback != NULL) {
		tinyFSDidFinishLoadingCallback(self, _cmd);
	}
}

- (void)probeResource:(id)resource replyHandler:(void (^)(FSProbeResult *_Nullable, NSError *_Nullable))reply
{
	if (tinyFSProbeResourceResultCallback != NULL) {
		NSError *error = nil;
		id result = tinyFSProbeResourceResultCallback(self, _cmd, resource, &error);
		TinyFSShimMarker("tinyfsgo-objc-before-probe-reply.marker");
		reply(result, error);
		TinyFSShimMarker("tinyfsgo-objc-after-probe-reply.marker");
		return;
	}
	if (tinyFSProbeResourceCallback != NULL) {
		tinyFSProbeResourceCallback(self, _cmd, resource, reply);
	}
}

- (void)loadResource:(id)resource options:(id)options replyHandler:(void (^)(FSVolume *_Nullable, NSError *_Nullable))reply
{
	if (tinyFSLoadResourceResultCallback != NULL) {
		NSError *error = nil;
		id volume = tinyFSLoadResourceResultCallback(self, _cmd, resource, options, &error);
		TinyFSShimMarker("tinyfsgo-objc-before-load-reply.marker");
		reply(volume, error);
		TinyFSShimMarker("tinyfsgo-objc-after-load-reply.marker");
		return;
	}
	if (tinyFSLoadResourceCallback != NULL) {
		tinyFSLoadResourceCallback(self, _cmd, resource, options, reply);
	}
}

- (void)unloadResource:(id)resource options:(id)options replyHandler:(void (^)(NSError *_Nullable))reply
{
	if (tinyFSUnloadResourceResultCallback != NULL) {
		NSError *error = tinyFSUnloadResourceResultCallback(self, _cmd, resource, options);
		TinyFSShimMarker("tinyfsgo-objc-before-unload-reply.marker");
		reply(error);
		TinyFSShimMarker("tinyfsgo-objc-after-unload-reply.marker");
		return;
	}
	if (tinyFSUnloadResourceCallback != NULL) {
		tinyFSUnloadResourceCallback(self, _cmd, resource, options, reply);
	}
}

@end
