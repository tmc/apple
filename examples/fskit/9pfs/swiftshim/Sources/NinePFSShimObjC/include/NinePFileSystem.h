#import <Foundation/Foundation.h>
#import <FSKit/FSKit.h>

NS_ASSUME_NONNULL_BEGIN

API_AVAILABLE(macos(15.4))
@interface NinePFileSystem : FSUnaryFileSystem <FSUnaryFileSystemOperations>
@end

API_AVAILABLE(macos(15.4))
@interface NinePFSItem : FSItem
@end

API_AVAILABLE(macos(15.4))
@interface NinePFSVolume : FSVolume <FSVolumeOperations, FSVolumeReadWriteOperations, FSVolumeOpenCloseOperations, FSVolumeAccessCheckOperations, FSVolumePathConfOperations, FSVolumeXattrOperations, FSVolumePreallocateOperations, FSVolumeRenameOperations>
@end

NS_ASSUME_NONNULL_END
