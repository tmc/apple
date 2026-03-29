// Code generated from Apple documentation. DO NOT EDIT.

package diskarbitration

import (
	"unsafe"
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
)

var (
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionBusNameKey
	KDADiskDescriptionBusNameKey string
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionBusPathKey
	KDADiskDescriptionBusPathKey string
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionDeviceGUIDKey
	KDADiskDescriptionDeviceGUIDKey string
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionDeviceInternalKey
	KDADiskDescriptionDeviceInternalKey string
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionDeviceModelKey
	KDADiskDescriptionDeviceModelKey string
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionDevicePathKey
	KDADiskDescriptionDevicePathKey string
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionDeviceProtocolKey
	KDADiskDescriptionDeviceProtocolKey string
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionDeviceRevisionKey
	KDADiskDescriptionDeviceRevisionKey string
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionDeviceTDMLockedKey
	KDADiskDescriptionDeviceTDMLockedKey string
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionDeviceUnitKey
	KDADiskDescriptionDeviceUnitKey string
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionDeviceVendorKey
	KDADiskDescriptionDeviceVendorKey string
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionFSKitPrefix
	KDADiskDescriptionFSKitPrefix string
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionMediaBSDMajorKey
	KDADiskDescriptionMediaBSDMajorKey string
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionMediaBSDMinorKey
	KDADiskDescriptionMediaBSDMinorKey string
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionMediaBSDNameKey
	KDADiskDescriptionMediaBSDNameKey string
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionMediaBSDUnitKey
	KDADiskDescriptionMediaBSDUnitKey string
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionMediaBlockSizeKey
	KDADiskDescriptionMediaBlockSizeKey string
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionMediaContentKey
	KDADiskDescriptionMediaContentKey string
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionMediaEjectableKey
	KDADiskDescriptionMediaEjectableKey string
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionMediaEncryptedKey
	KDADiskDescriptionMediaEncryptedKey string
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionMediaEncryptionDetailKey
	KDADiskDescriptionMediaEncryptionDetailKey string
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionMediaIconKey
	KDADiskDescriptionMediaIconKey string
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionMediaKindKey
	KDADiskDescriptionMediaKindKey string
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionMediaLeafKey
	KDADiskDescriptionMediaLeafKey string
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionMediaNameKey
	KDADiskDescriptionMediaNameKey string
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionMediaPathKey
	KDADiskDescriptionMediaPathKey string
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionMediaRemovableKey
	KDADiskDescriptionMediaRemovableKey string
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionMediaSizeKey
	KDADiskDescriptionMediaSizeKey string
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionMediaTypeKey
	KDADiskDescriptionMediaTypeKey string
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionMediaUUIDKey
	KDADiskDescriptionMediaUUIDKey string
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionMediaWholeKey
	KDADiskDescriptionMediaWholeKey string
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionMediaWritableKey
	KDADiskDescriptionMediaWritableKey string
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionRepairRunningKey
	KDADiskDescriptionRepairRunningKey string
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionVolumeKindKey
	KDADiskDescriptionVolumeKindKey string
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionVolumeMountableKey
	KDADiskDescriptionVolumeMountableKey string
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionVolumeNameKey
	KDADiskDescriptionVolumeNameKey string
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionVolumeNetworkKey
	KDADiskDescriptionVolumeNetworkKey string
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionVolumePathKey
	KDADiskDescriptionVolumePathKey string
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionVolumeTypeKey
	KDADiskDescriptionVolumeTypeKey string
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionVolumeUUIDKey
	KDADiskDescriptionVolumeUUIDKey string
)

var (
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionMatchMediaUnformatted
	KDADiskDescriptionMatchMediaUnformatted corefoundation.CFDictionaryRef
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionMatchMediaWhole
	KDADiskDescriptionMatchMediaWhole corefoundation.CFDictionaryRef
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionMatchVolumeMountable
	KDADiskDescriptionMatchVolumeMountable corefoundation.CFDictionaryRef
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionMatchVolumeUnrecognized
	KDADiskDescriptionMatchVolumeUnrecognized corefoundation.CFDictionaryRef
)

var (
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionWatchVolumeName
	KDADiskDescriptionWatchVolumeName corefoundation.CFArrayRef
	// See: https://developer.apple.com/documentation/DiskArbitration/kDADiskDescriptionWatchVolumePath
	KDADiskDescriptionWatchVolumePath corefoundation.CFArrayRef
)
func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionBusNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionBusNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionBusPathKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionBusPathKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionDeviceGUIDKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionDeviceGUIDKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionDeviceInternalKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionDeviceInternalKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionDeviceModelKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionDeviceModelKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionDevicePathKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionDevicePathKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionDeviceProtocolKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionDeviceProtocolKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionDeviceRevisionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionDeviceRevisionKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionDeviceTDMLockedKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionDeviceTDMLockedKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionDeviceUnitKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionDeviceUnitKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionDeviceVendorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionDeviceVendorKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionFSKitPrefix"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionFSKitPrefix = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionMatchMediaUnformatted"); err == nil && ptr != 0 {
		KDADiskDescriptionMatchMediaUnformatted = *(*corefoundation.CFDictionaryRef)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionMatchMediaWhole"); err == nil && ptr != 0 {
		KDADiskDescriptionMatchMediaWhole = *(*corefoundation.CFDictionaryRef)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionMatchVolumeMountable"); err == nil && ptr != 0 {
		KDADiskDescriptionMatchVolumeMountable = *(*corefoundation.CFDictionaryRef)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionMatchVolumeUnrecognized"); err == nil && ptr != 0 {
		KDADiskDescriptionMatchVolumeUnrecognized = *(*corefoundation.CFDictionaryRef)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionMediaBSDMajorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionMediaBSDMajorKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionMediaBSDMinorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionMediaBSDMinorKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionMediaBSDNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionMediaBSDNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionMediaBSDUnitKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionMediaBSDUnitKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionMediaBlockSizeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionMediaBlockSizeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionMediaContentKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionMediaContentKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionMediaEjectableKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionMediaEjectableKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionMediaEncryptedKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionMediaEncryptedKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionMediaEncryptionDetailKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionMediaEncryptionDetailKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionMediaIconKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionMediaIconKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionMediaKindKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionMediaKindKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionMediaLeafKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionMediaLeafKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionMediaNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionMediaNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionMediaPathKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionMediaPathKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionMediaRemovableKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionMediaRemovableKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionMediaSizeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionMediaSizeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionMediaTypeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionMediaTypeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionMediaUUIDKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionMediaUUIDKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionMediaWholeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionMediaWholeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionMediaWritableKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionMediaWritableKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionRepairRunningKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionRepairRunningKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionVolumeKindKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionVolumeKindKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionVolumeMountableKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionVolumeMountableKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionVolumeNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionVolumeNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionVolumeNetworkKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionVolumeNetworkKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionVolumePathKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionVolumePathKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionVolumeTypeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionVolumeTypeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionVolumeUUIDKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KDADiskDescriptionVolumeUUIDKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionWatchVolumeName"); err == nil && ptr != 0 {
		KDADiskDescriptionWatchVolumeName = *(*corefoundation.CFArrayRef)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kDADiskDescriptionWatchVolumePath"); err == nil && ptr != 0 {
		KDADiskDescriptionWatchVolumePath = *(*corefoundation.CFArrayRef)(unsafe.Pointer(ptr))
	}

}

