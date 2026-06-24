#import <Foundation/Foundation.h>
#import <FSKit/FSKit.h>

NS_ASSUME_NONNULL_BEGIN

API_AVAILABLE(macos(15.4))
@interface GoTinyFSFileSystem : FSUnaryFileSystem <FSUnaryFileSystemOperations>
@end

API_AVAILABLE(macos(15.4))
@interface GoTinyFSItem : FSItem
@property(nonatomic) FSItemID itemID;
@property(nonatomic) FSItemID parentID;
@property(nonatomic) FSItemType type;
@end

API_AVAILABLE(macos(15.4))
@interface GoTinyFSVolume : FSVolume <FSVolumeOperations>
@end

typedef void (*TinyFSProbeResourceCallback)(id self, SEL _cmd, id resource, id reply);
typedef void (*TinyFSLoadResourceCallback)(id self, SEL _cmd, id resource, id options, id reply);
typedef void (*TinyFSUnloadResourceCallback)(id self, SEL _cmd, id resource, id options, id reply);
typedef void (*TinyFSDidFinishLoadingCallback)(id self, SEL _cmd);
typedef id _Nullable (*TinyFSProbeResourceResultCallback)(id self, SEL _cmd, id resource, NSError *_Nullable *_Nullable error);
typedef id _Nullable (*TinyFSLoadResourceResultCallback)(id self, SEL _cmd, id resource, id options, NSError *_Nullable *_Nullable error);
typedef NSError *_Nullable (*TinyFSUnloadResourceResultCallback)(id self, SEL _cmd, id resource, id options);
typedef void (*TinyFSVolumeCallback)(id self, SEL _cmd);

void TinyFSShimSetProbeResourceCallback(TinyFSProbeResourceCallback callback);
void TinyFSShimSetLoadResourceCallback(TinyFSLoadResourceCallback callback);
void TinyFSShimSetUnloadResourceCallback(TinyFSUnloadResourceCallback callback);
void TinyFSShimSetDidFinishLoadingCallback(TinyFSDidFinishLoadingCallback callback);
void TinyFSShimSetProbeResourceResultCallback(TinyFSProbeResourceResultCallback callback);
void TinyFSShimSetLoadResourceResultCallback(TinyFSLoadResourceResultCallback callback);
void TinyFSShimSetUnloadResourceResultCallback(TinyFSUnloadResourceResultCallback callback);
void TinyFSShimReplyProbe(id reply, id _Nullable result, NSError *_Nullable error);
void TinyFSShimReplyVolume(id reply, id _Nullable volume, NSError *_Nullable error);
void TinyFSShimReplyError(id reply, NSError *_Nullable error);
void TinyFSShimSetVolumeActivateCallback(TinyFSVolumeCallback callback);
void TinyFSShimSetVolumeMountCallback(TinyFSVolumeCallback callback);
void TinyFSShimSetVolumeUnmountCallback(TinyFSVolumeCallback callback);
id TinyFSShimNewVolume(void);

NS_ASSUME_NONNULL_END
