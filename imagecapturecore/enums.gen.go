// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceLocationType
type ICDeviceLocationType uint

const (
	// ICDeviceLocationTypeBluetooth: A paired Bluetooth device.
	ICDeviceLocationTypeBluetooth ICDeviceLocationType = 0x800
	// ICDeviceLocationTypeBonjour: A supported Bonjour services device.
	ICDeviceLocationTypeBonjour ICDeviceLocationType = 0x400
	// ICDeviceLocationTypeLocal: A device that’s directly attached to the Mac through its USB or FireWire port.
	ICDeviceLocationTypeLocal ICDeviceLocationType = 0x100
	// ICDeviceLocationTypeShared: A device that’s shared by other Mac hosts.
	ICDeviceLocationTypeShared ICDeviceLocationType = 0x200
)

func (e ICDeviceLocationType) String() string {
	switch e {
	case ICDeviceLocationTypeBluetooth:
		return "ICDeviceLocationTypeBluetooth"
	case ICDeviceLocationTypeBonjour:
		return "ICDeviceLocationTypeBonjour"
	case ICDeviceLocationTypeLocal:
		return "ICDeviceLocationTypeLocal"
	case ICDeviceLocationTypeShared:
		return "ICDeviceLocationTypeShared"
	default:
		return fmt.Sprintf("ICDeviceLocationType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceLocationTypeMask
type ICDeviceLocationTypeMask uint

const (
	// ICDeviceLocationTypeMaskBluetooth: A mask for detecting a paired Bluetooth device.
	ICDeviceLocationTypeMaskBluetooth ICDeviceLocationTypeMask = 0x800
	// ICDeviceLocationTypeMaskBonjour: A mask for detecting a network device that publishes a Bonjour service.
	ICDeviceLocationTypeMaskBonjour ICDeviceLocationTypeMask = 0x400
	// ICDeviceLocationTypeMaskLocal: A mask for detecting a local device, such as USB or FireWire.
	ICDeviceLocationTypeMaskLocal ICDeviceLocationTypeMask = 0x100
	// ICDeviceLocationTypeMaskRemote: A mask for detecting a remote device, such as a shared, Bonjour, or Bluetooth device.
	ICDeviceLocationTypeMaskRemote ICDeviceLocationTypeMask = 0xfe00
	// ICDeviceLocationTypeMaskShared: A mask for detecting a device shared by another Mac host.
	ICDeviceLocationTypeMaskShared ICDeviceLocationTypeMask = 0x200
)

func (e ICDeviceLocationTypeMask) String() string {
	switch e {
	case ICDeviceLocationTypeMaskBluetooth:
		return "ICDeviceLocationTypeMaskBluetooth"
	case ICDeviceLocationTypeMaskBonjour:
		return "ICDeviceLocationTypeMaskBonjour"
	case ICDeviceLocationTypeMaskLocal:
		return "ICDeviceLocationTypeMaskLocal"
	case ICDeviceLocationTypeMaskRemote:
		return "ICDeviceLocationTypeMaskRemote"
	case ICDeviceLocationTypeMaskShared:
		return "ICDeviceLocationTypeMaskShared"
	default:
		return fmt.Sprintf("ICDeviceLocationTypeMask(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceType
type ICDeviceType uint

const (
	// ICDeviceTypeCamera: The device is a camera.
	ICDeviceTypeCamera ICDeviceType = 0x1
	// ICDeviceTypeScanner: The device is a scanner.
	ICDeviceTypeScanner ICDeviceType = 0x2
)

func (e ICDeviceType) String() string {
	switch e {
	case ICDeviceTypeCamera:
		return "ICDeviceTypeCamera"
	case ICDeviceTypeScanner:
		return "ICDeviceTypeScanner"
	default:
		return fmt.Sprintf("ICDeviceType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceTypeMask
type ICDeviceTypeMask uint

const (
	// ICDeviceTypeMaskCamera: A mask for detecting a camera.
	ICDeviceTypeMaskCamera ICDeviceTypeMask = 0x1
	// ICDeviceTypeMaskScanner: A mask for detecting a scanner.
	ICDeviceTypeMaskScanner ICDeviceTypeMask = 0x2
)

func (e ICDeviceTypeMask) String() string {
	switch e {
	case ICDeviceTypeMaskCamera:
		return "ICDeviceTypeMaskCamera"
	case ICDeviceTypeMaskScanner:
		return "ICDeviceTypeMaskScanner"
	default:
		return fmt.Sprintf("ICDeviceTypeMask(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICEXIFOrientationType
type ICEXIFOrientationType uint

const (
	// ICEXIFOrientation1: Normal
	ICEXIFOrientation1 ICEXIFOrientationType = 1
	// ICEXIFOrientation2: Flipped horizontally
	ICEXIFOrientation2 ICEXIFOrientationType = 2
	// ICEXIFOrientation3: Rotated 180°
	ICEXIFOrientation3 ICEXIFOrientationType = 3
	// ICEXIFOrientation4: Flipped vertically
	ICEXIFOrientation4 ICEXIFOrientationType = 4
	// ICEXIFOrientation5: Rotated 90° CCW and flipped vertically
	ICEXIFOrientation5 ICEXIFOrientationType = 5
	// ICEXIFOrientation6: Rotated 90° CCW
	ICEXIFOrientation6 ICEXIFOrientationType = 6
	// ICEXIFOrientation7: Rotated 90° CW and flipped vertically
	ICEXIFOrientation7 ICEXIFOrientationType = 7
	// ICEXIFOrientation8: Rotated 90° CW
	ICEXIFOrientation8 ICEXIFOrientationType = 8
)

func (e ICEXIFOrientationType) String() string {
	switch e {
	case ICEXIFOrientation1:
		return "ICEXIFOrientation1"
	case ICEXIFOrientation2:
		return "ICEXIFOrientation2"
	case ICEXIFOrientation3:
		return "ICEXIFOrientation3"
	case ICEXIFOrientation4:
		return "ICEXIFOrientation4"
	case ICEXIFOrientation5:
		return "ICEXIFOrientation5"
	case ICEXIFOrientation6:
		return "ICEXIFOrientation6"
	case ICEXIFOrientation7:
		return "ICEXIFOrientation7"
	case ICEXIFOrientation8:
		return "ICEXIFOrientation8"
	default:
		return fmt.Sprintf("ICEXIFOrientationType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICMediaPresentation
type ICMediaPresentation uint

const (
	ICMediaPresentationConvertedAssets ICMediaPresentation = 1
	ICMediaPresentationOriginalAssets  ICMediaPresentation = 2
)

func (e ICMediaPresentation) String() string {
	switch e {
	case ICMediaPresentationConvertedAssets:
		return "ICMediaPresentationConvertedAssets"
	case ICMediaPresentationOriginalAssets:
		return "ICMediaPresentationOriginalAssets"
	default:
		return fmt.Sprintf("ICMediaPresentation(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnCodeOffset
type ICReturnCodeOffset int

const ()

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnConnectionError/Code
type ICReturnConnectionErrorCode int

const (
	// ICReturnConnectionClosedSessionSuddenly: Device closed session without request.
	ICReturnConnectionClosedSessionSuddenly ICReturnConnectionErrorCode = -21349
	// ICReturnConnectionDriverExited: Device driver exited without request.
	ICReturnConnectionDriverExited ICReturnConnectionErrorCode = -21350
	// ICReturnConnectionEjectFailed: Device reports eject has failed.
	ICReturnConnectionEjectFailed ICReturnConnectionErrorCode = -21346
	// ICReturnConnectionEjectedSuddenly: Device ejected without request.
	ICReturnConnectionEjectedSuddenly ICReturnConnectionErrorCode = -21348
	// ICReturnConnectionFailedToOpen: Failed to open a connection to the device.
	ICReturnConnectionFailedToOpen ICReturnConnectionErrorCode = -21345
	// ICReturnConnectionFailedToOpenDevice: Failed to open the device.
	ICReturnConnectionFailedToOpenDevice        ICReturnConnectionErrorCode = -21344
	ICReturnConnectionNotAuthorizedToOpenDevice ICReturnConnectionErrorCode = -21343
	// ICReturnConnectionSessionAlreadyOpen: Device reports session is already open.
	ICReturnConnectionSessionAlreadyOpen ICReturnConnectionErrorCode = -21347
)

func (e ICReturnConnectionErrorCode) String() string {
	switch e {
	case ICReturnConnectionClosedSessionSuddenly:
		return "ICReturnConnectionClosedSessionSuddenly"
	case ICReturnConnectionDriverExited:
		return "ICReturnConnectionDriverExited"
	case ICReturnConnectionEjectFailed:
		return "ICReturnConnectionEjectFailed"
	case ICReturnConnectionEjectedSuddenly:
		return "ICReturnConnectionEjectedSuddenly"
	case ICReturnConnectionFailedToOpen:
		return "ICReturnConnectionFailedToOpen"
	case ICReturnConnectionFailedToOpenDevice:
		return "ICReturnConnectionFailedToOpenDevice"
	case ICReturnConnectionNotAuthorizedToOpenDevice:
		return "ICReturnConnectionNotAuthorizedToOpenDevice"
	case ICReturnConnectionSessionAlreadyOpen:
		return "ICReturnConnectionSessionAlreadyOpen"
	default:
		return fmt.Sprintf("ICReturnConnectionErrorCode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnDownloadError/Code
type ICReturnDownloadErrorCode int

const (
	// ICReturnDownloadFileWritable: The destination file is not writable.
	ICReturnDownloadFileWritable ICReturnDownloadErrorCode = -21099
	// ICReturnDownloadPathInvalid: The destination path is invalid.
	ICReturnDownloadPathInvalid ICReturnDownloadErrorCode = -21100
)

func (e ICReturnDownloadErrorCode) String() string {
	switch e {
	case ICReturnDownloadFileWritable:
		return "ICReturnDownloadFileWritable"
	case ICReturnDownloadPathInvalid:
		return "ICReturnDownloadPathInvalid"
	default:
		return fmt.Sprintf("ICReturnDownloadErrorCode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnMetadataError/Code
type ICReturnMetadataErrorCode int

const (
	// ICReturnMetadataAlreadyFetching: Item metadata request is being serviced.
	ICReturnMetadataAlreadyFetching ICReturnMetadataErrorCode = -20149
	// ICReturnMetadataCanceled: Item metadata request has been canceled.
	ICReturnMetadataCanceled ICReturnMetadataErrorCode = -20148
	// ICReturnMetadataInvalid: Item metadata request completed with invalid result.
	ICReturnMetadataInvalid ICReturnMetadataErrorCode = -20147
	// ICReturnMetadataNotAvailable: Item does not have metadata available.
	ICReturnMetadataNotAvailable ICReturnMetadataErrorCode = -20150
)

func (e ICReturnMetadataErrorCode) String() string {
	switch e {
	case ICReturnMetadataAlreadyFetching:
		return "ICReturnMetadataAlreadyFetching"
	case ICReturnMetadataCanceled:
		return "ICReturnMetadataCanceled"
	case ICReturnMetadataInvalid:
		return "ICReturnMetadataInvalid"
	case ICReturnMetadataNotAvailable:
		return "ICReturnMetadataNotAvailable"
	default:
		return fmt.Sprintf("ICReturnMetadataErrorCode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnObjectError/Code
type ICReturnObjectErrorCode int

const (
	// ICReturnCodeObjectCouldNotBeRead: The object could not be read.
	ICReturnCodeObjectCouldNotBeRead ICReturnObjectErrorCode = -21448
	// ICReturnCodeObjectDataEmpty: The object data is empty.
	ICReturnCodeObjectDataEmpty ICReturnObjectErrorCode = -21447
	// ICReturnCodeObjectDataOffsetInvalid: The object data offset is invalid.
	ICReturnCodeObjectDataOffsetInvalid   ICReturnObjectErrorCode = -21449
	ICReturnCodeObjectDataRequestTooLarge ICReturnObjectErrorCode = -21446
	// ICReturnCodeObjectDoesNotExist: The object does not exist.
	ICReturnCodeObjectDoesNotExist ICReturnObjectErrorCode = -21450
)

func (e ICReturnObjectErrorCode) String() string {
	switch e {
	case ICReturnCodeObjectCouldNotBeRead:
		return "ICReturnCodeObjectCouldNotBeRead"
	case ICReturnCodeObjectDataEmpty:
		return "ICReturnCodeObjectDataEmpty"
	case ICReturnCodeObjectDataOffsetInvalid:
		return "ICReturnCodeObjectDataOffsetInvalid"
	case ICReturnCodeObjectDataRequestTooLarge:
		return "ICReturnCodeObjectDataRequestTooLarge"
	case ICReturnCodeObjectDoesNotExist:
		return "ICReturnCodeObjectDoesNotExist"
	default:
		return fmt.Sprintf("ICReturnObjectErrorCode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnPTPDeviceError/Code
type ICReturnPTPDeviceErrorCode int

const (
	// ICReturnPTPFailedToSendCommand: Sending a PTP command failed.
	ICReturnPTPFailedToSendCommand        ICReturnPTPDeviceErrorCode = -21250
	ICReturnPTPNotAuthorizedToSendCommand ICReturnPTPDeviceErrorCode = -21249
)

func (e ICReturnPTPDeviceErrorCode) String() string {
	switch e {
	case ICReturnPTPFailedToSendCommand:
		return "ICReturnPTPFailedToSendCommand"
	case ICReturnPTPNotAuthorizedToSendCommand:
		return "ICReturnPTPNotAuthorizedToSendCommand"
	default:
		return fmt.Sprintf("ICReturnPTPDeviceErrorCode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnThumbnailError/Code
type ICReturnThumbnailErrorCode int

const (
	// ICReturnThumbnailAlreadyFetching: Item thumbnail request is being serviced.
	ICReturnThumbnailAlreadyFetching ICReturnThumbnailErrorCode = -20999
	// ICReturnThumbnailCanceled: Item thumbnail request has been canceled.
	ICReturnThumbnailCanceled ICReturnThumbnailErrorCode = -20098
	// ICReturnThumbnailInvalid: Item thumbnail request completed with invalid result.
	ICReturnThumbnailInvalid ICReturnThumbnailErrorCode = -20097
	// ICReturnThumbnailNotAvailable: Item does not have thumbnail available.
	ICReturnThumbnailNotAvailable ICReturnThumbnailErrorCode = -21000
)

func (e ICReturnThumbnailErrorCode) String() string {
	switch e {
	case ICReturnThumbnailAlreadyFetching:
		return "ICReturnThumbnailAlreadyFetching"
	case ICReturnThumbnailCanceled:
		return "ICReturnThumbnailCanceled"
	case ICReturnThumbnailInvalid:
		return "ICReturnThumbnailInvalid"
	case ICReturnThumbnailNotAvailable:
		return "ICReturnThumbnailNotAvailable"
	default:
		return fmt.Sprintf("ICReturnThumbnailErrorCode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerBitDepth
type ICScannerBitDepth uint

const (
	ICScannerBitDepth16Bits ICScannerBitDepth = 16
	ICScannerBitDepth1Bit   ICScannerBitDepth = 1
	ICScannerBitDepth8Bits  ICScannerBitDepth = 8
)

func (e ICScannerBitDepth) String() string {
	switch e {
	case ICScannerBitDepth16Bits:
		return "ICScannerBitDepth16Bits"
	case ICScannerBitDepth1Bit:
		return "ICScannerBitDepth1Bit"
	case ICScannerBitDepth8Bits:
		return "ICScannerBitDepth8Bits"
	default:
		return fmt.Sprintf("ICScannerBitDepth(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerColorDataFormatType
type ICScannerColorDataFormatType uint

const (
	ICScannerColorDataFormatTypeChunky ICScannerColorDataFormatType = 0
	ICScannerColorDataFormatTypePlanar ICScannerColorDataFormatType = 1
)

func (e ICScannerColorDataFormatType) String() string {
	switch e {
	case ICScannerColorDataFormatTypeChunky:
		return "ICScannerColorDataFormatTypeChunky"
	case ICScannerColorDataFormatTypePlanar:
		return "ICScannerColorDataFormatTypePlanar"
	default:
		return fmt.Sprintf("ICScannerColorDataFormatType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerDocumentType
type ICScannerDocumentType uint

const (
	ICScannerDocumentType10           ICScannerDocumentType = 25
	ICScannerDocumentType10R          ICScannerDocumentType = 67
	ICScannerDocumentType110          ICScannerDocumentType = 72
	ICScannerDocumentType11R          ICScannerDocumentType = 69
	ICScannerDocumentType12R          ICScannerDocumentType = 70
	ICScannerDocumentType135          ICScannerDocumentType = 76
	ICScannerDocumentType2A0          ICScannerDocumentType = 18
	ICScannerDocumentType3R           ICScannerDocumentType = 61
	ICScannerDocumentType4A0          ICScannerDocumentType = 17
	ICScannerDocumentType4R           ICScannerDocumentType = 62
	ICScannerDocumentType5R           ICScannerDocumentType = 63
	ICScannerDocumentType6R           ICScannerDocumentType = 64
	ICScannerDocumentType8R           ICScannerDocumentType = 65
	ICScannerDocumentTypeA0           ICScannerDocumentType = 19
	ICScannerDocumentTypeA1           ICScannerDocumentType = 20
	ICScannerDocumentTypeA2           ICScannerDocumentType = 21
	ICScannerDocumentTypeA3           ICScannerDocumentType = 11
	ICScannerDocumentTypeA4           ICScannerDocumentType = 1
	ICScannerDocumentTypeA5           ICScannerDocumentType = 5
	ICScannerDocumentTypeA6           ICScannerDocumentType = 13
	ICScannerDocumentTypeA7           ICScannerDocumentType = 22
	ICScannerDocumentTypeA8           ICScannerDocumentType = 23
	ICScannerDocumentTypeA9           ICScannerDocumentType = 24
	ICScannerDocumentTypeAPSC         ICScannerDocumentType = 74
	ICScannerDocumentTypeAPSH         ICScannerDocumentType = 73
	ICScannerDocumentTypeAPSP         ICScannerDocumentType = 75
	ICScannerDocumentTypeB5           ICScannerDocumentType = 2
	ICScannerDocumentTypeBusinessCard ICScannerDocumentType = 53
	ICScannerDocumentTypeC0           ICScannerDocumentType = 44
	ICScannerDocumentTypeC1           ICScannerDocumentType = 45
	ICScannerDocumentTypeC10          ICScannerDocumentType = 51
	ICScannerDocumentTypeC2           ICScannerDocumentType = 46
	ICScannerDocumentTypeC3           ICScannerDocumentType = 47
	ICScannerDocumentTypeC4           ICScannerDocumentType = 14
	ICScannerDocumentTypeC5           ICScannerDocumentType = 15
	ICScannerDocumentTypeC6           ICScannerDocumentType = 16
	ICScannerDocumentTypeC7           ICScannerDocumentType = 48
	ICScannerDocumentTypeC8           ICScannerDocumentType = 49
	ICScannerDocumentTypeC9           ICScannerDocumentType = 50
	ICScannerDocumentTypeDefault      ICScannerDocumentType = 0
	ICScannerDocumentTypeE            ICScannerDocumentType = 60
	ICScannerDocumentTypeISOB0        ICScannerDocumentType = 26
	ICScannerDocumentTypeISOB1        ICScannerDocumentType = 27
	ICScannerDocumentTypeISOB10       ICScannerDocumentType = 33
	ICScannerDocumentTypeISOB2        ICScannerDocumentType = 28
	ICScannerDocumentTypeISOB3        ICScannerDocumentType = 12
	ICScannerDocumentTypeISOB4        ICScannerDocumentType = 6
	ICScannerDocumentTypeISOB5        ICScannerDocumentType = 29
	ICScannerDocumentTypeISOB6        ICScannerDocumentType = 7
	ICScannerDocumentTypeISOB7        ICScannerDocumentType = 30
	ICScannerDocumentTypeISOB8        ICScannerDocumentType = 31
	ICScannerDocumentTypeISOB9        ICScannerDocumentType = 32
	ICScannerDocumentTypeJISB0        ICScannerDocumentType = 34
	ICScannerDocumentTypeJISB1        ICScannerDocumentType = 35
	ICScannerDocumentTypeJISB10       ICScannerDocumentType = 43
	ICScannerDocumentTypeJISB2        ICScannerDocumentType = 36
	ICScannerDocumentTypeJISB3        ICScannerDocumentType = 37
	ICScannerDocumentTypeJISB4        ICScannerDocumentType = 38
	ICScannerDocumentTypeJISB6        ICScannerDocumentType = 39
	ICScannerDocumentTypeJISB7        ICScannerDocumentType = 40
	ICScannerDocumentTypeJISB8        ICScannerDocumentType = 41
	ICScannerDocumentTypeJISB9        ICScannerDocumentType = 42
	ICScannerDocumentTypeLF           ICScannerDocumentType = 78
	ICScannerDocumentTypeMF           ICScannerDocumentType = 77
	ICScannerDocumentTypeS10R         ICScannerDocumentType = 68
	ICScannerDocumentTypeS12R         ICScannerDocumentType = 71
	ICScannerDocumentTypeS8R          ICScannerDocumentType = 66
	ICScannerDocumentTypeUSExecutive  ICScannerDocumentType = 10
	ICScannerDocumentTypeUSLedger     ICScannerDocumentType = 9
	ICScannerDocumentTypeUSLegal      ICScannerDocumentType = 4
	ICScannerDocumentTypeUSLetter     ICScannerDocumentType = 3
	ICScannerDocumentTypeUSStatement  ICScannerDocumentType = 52
)

func (e ICScannerDocumentType) String() string {
	switch e {
	case ICScannerDocumentType10:
		return "ICScannerDocumentType10"
	case ICScannerDocumentType10R:
		return "ICScannerDocumentType10R"
	case ICScannerDocumentType110:
		return "ICScannerDocumentType110"
	case ICScannerDocumentType11R:
		return "ICScannerDocumentType11R"
	case ICScannerDocumentType12R:
		return "ICScannerDocumentType12R"
	case ICScannerDocumentType135:
		return "ICScannerDocumentType135"
	case ICScannerDocumentType2A0:
		return "ICScannerDocumentType2A0"
	case ICScannerDocumentType3R:
		return "ICScannerDocumentType3R"
	case ICScannerDocumentType4A0:
		return "ICScannerDocumentType4A0"
	case ICScannerDocumentType4R:
		return "ICScannerDocumentType4R"
	case ICScannerDocumentType5R:
		return "ICScannerDocumentType5R"
	case ICScannerDocumentType6R:
		return "ICScannerDocumentType6R"
	case ICScannerDocumentType8R:
		return "ICScannerDocumentType8R"
	case ICScannerDocumentTypeA0:
		return "ICScannerDocumentTypeA0"
	case ICScannerDocumentTypeA1:
		return "ICScannerDocumentTypeA1"
	case ICScannerDocumentTypeA2:
		return "ICScannerDocumentTypeA2"
	case ICScannerDocumentTypeA3:
		return "ICScannerDocumentTypeA3"
	case ICScannerDocumentTypeA4:
		return "ICScannerDocumentTypeA4"
	case ICScannerDocumentTypeA5:
		return "ICScannerDocumentTypeA5"
	case ICScannerDocumentTypeA6:
		return "ICScannerDocumentTypeA6"
	case ICScannerDocumentTypeA7:
		return "ICScannerDocumentTypeA7"
	case ICScannerDocumentTypeA8:
		return "ICScannerDocumentTypeA8"
	case ICScannerDocumentTypeA9:
		return "ICScannerDocumentTypeA9"
	case ICScannerDocumentTypeAPSC:
		return "ICScannerDocumentTypeAPSC"
	case ICScannerDocumentTypeAPSH:
		return "ICScannerDocumentTypeAPSH"
	case ICScannerDocumentTypeAPSP:
		return "ICScannerDocumentTypeAPSP"
	case ICScannerDocumentTypeB5:
		return "ICScannerDocumentTypeB5"
	case ICScannerDocumentTypeBusinessCard:
		return "ICScannerDocumentTypeBusinessCard"
	case ICScannerDocumentTypeC0:
		return "ICScannerDocumentTypeC0"
	case ICScannerDocumentTypeC1:
		return "ICScannerDocumentTypeC1"
	case ICScannerDocumentTypeC10:
		return "ICScannerDocumentTypeC10"
	case ICScannerDocumentTypeC2:
		return "ICScannerDocumentTypeC2"
	case ICScannerDocumentTypeC3:
		return "ICScannerDocumentTypeC3"
	case ICScannerDocumentTypeC4:
		return "ICScannerDocumentTypeC4"
	case ICScannerDocumentTypeC5:
		return "ICScannerDocumentTypeC5"
	case ICScannerDocumentTypeC6:
		return "ICScannerDocumentTypeC6"
	case ICScannerDocumentTypeC7:
		return "ICScannerDocumentTypeC7"
	case ICScannerDocumentTypeC8:
		return "ICScannerDocumentTypeC8"
	case ICScannerDocumentTypeC9:
		return "ICScannerDocumentTypeC9"
	case ICScannerDocumentTypeDefault:
		return "ICScannerDocumentTypeDefault"
	case ICScannerDocumentTypeE:
		return "ICScannerDocumentTypeE"
	case ICScannerDocumentTypeISOB0:
		return "ICScannerDocumentTypeISOB0"
	case ICScannerDocumentTypeISOB1:
		return "ICScannerDocumentTypeISOB1"
	case ICScannerDocumentTypeISOB10:
		return "ICScannerDocumentTypeISOB10"
	case ICScannerDocumentTypeISOB2:
		return "ICScannerDocumentTypeISOB2"
	case ICScannerDocumentTypeISOB3:
		return "ICScannerDocumentTypeISOB3"
	case ICScannerDocumentTypeISOB4:
		return "ICScannerDocumentTypeISOB4"
	case ICScannerDocumentTypeISOB5:
		return "ICScannerDocumentTypeISOB5"
	case ICScannerDocumentTypeISOB6:
		return "ICScannerDocumentTypeISOB6"
	case ICScannerDocumentTypeISOB7:
		return "ICScannerDocumentTypeISOB7"
	case ICScannerDocumentTypeISOB8:
		return "ICScannerDocumentTypeISOB8"
	case ICScannerDocumentTypeISOB9:
		return "ICScannerDocumentTypeISOB9"
	case ICScannerDocumentTypeJISB0:
		return "ICScannerDocumentTypeJISB0"
	case ICScannerDocumentTypeJISB1:
		return "ICScannerDocumentTypeJISB1"
	case ICScannerDocumentTypeJISB10:
		return "ICScannerDocumentTypeJISB10"
	case ICScannerDocumentTypeJISB2:
		return "ICScannerDocumentTypeJISB2"
	case ICScannerDocumentTypeJISB3:
		return "ICScannerDocumentTypeJISB3"
	case ICScannerDocumentTypeJISB4:
		return "ICScannerDocumentTypeJISB4"
	case ICScannerDocumentTypeJISB6:
		return "ICScannerDocumentTypeJISB6"
	case ICScannerDocumentTypeJISB7:
		return "ICScannerDocumentTypeJISB7"
	case ICScannerDocumentTypeJISB8:
		return "ICScannerDocumentTypeJISB8"
	case ICScannerDocumentTypeJISB9:
		return "ICScannerDocumentTypeJISB9"
	case ICScannerDocumentTypeLF:
		return "ICScannerDocumentTypeLF"
	case ICScannerDocumentTypeMF:
		return "ICScannerDocumentTypeMF"
	case ICScannerDocumentTypeS10R:
		return "ICScannerDocumentTypeS10R"
	case ICScannerDocumentTypeS12R:
		return "ICScannerDocumentTypeS12R"
	case ICScannerDocumentTypeS8R:
		return "ICScannerDocumentTypeS8R"
	case ICScannerDocumentTypeUSExecutive:
		return "ICScannerDocumentTypeUSExecutive"
	case ICScannerDocumentTypeUSLedger:
		return "ICScannerDocumentTypeUSLedger"
	case ICScannerDocumentTypeUSLegal:
		return "ICScannerDocumentTypeUSLegal"
	case ICScannerDocumentTypeUSLetter:
		return "ICScannerDocumentTypeUSLetter"
	case ICScannerDocumentTypeUSStatement:
		return "ICScannerDocumentTypeUSStatement"
	default:
		return fmt.Sprintf("ICScannerDocumentType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeatureType
type ICScannerFeatureType uint

const (
	// ICScannerFeatureTypeBoolean: A feature with a value of YES or NO.
	ICScannerFeatureTypeBoolean ICScannerFeatureType = 2
	// ICScannerFeatureTypeEnumeration: A feature that can have one of several discrete values, strings, or numbers.
	ICScannerFeatureTypeEnumeration ICScannerFeatureType = 0
	// ICScannerFeatureTypeRange: A feature with a value that lies within a range.
	ICScannerFeatureTypeRange ICScannerFeatureType = 1
	// ICScannerFeatureTypeTemplate: A group of one or more rectangular scan areas that can be used with a scanner functional unit.
	ICScannerFeatureTypeTemplate ICScannerFeatureType = 3
)

func (e ICScannerFeatureType) String() string {
	switch e {
	case ICScannerFeatureTypeBoolean:
		return "ICScannerFeatureTypeBoolean"
	case ICScannerFeatureTypeEnumeration:
		return "ICScannerFeatureTypeEnumeration"
	case ICScannerFeatureTypeRange:
		return "ICScannerFeatureTypeRange"
	case ICScannerFeatureTypeTemplate:
		return "ICScannerFeatureTypeTemplate"
	default:
		return fmt.Sprintf("ICScannerFeatureType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnitState
type ICScannerFunctionalUnitState uint

const (
	// ICScannerFunctionalUnitStateOverviewScanInProgress: A flag indicating that the functional unit is performing an overview scan.
	ICScannerFunctionalUnitStateOverviewScanInProgress ICScannerFunctionalUnitState = 4
	// ICScannerFunctionalUnitStateReady: A flag indicating that the functional unit is ready for operation.
	ICScannerFunctionalUnitStateReady ICScannerFunctionalUnitState = 1
	// ICScannerFunctionalUnitStateScanInProgress: A flag indicating that the functional unit is performing a scan.
	ICScannerFunctionalUnitStateScanInProgress ICScannerFunctionalUnitState = 2
)

func (e ICScannerFunctionalUnitState) String() string {
	switch e {
	case ICScannerFunctionalUnitStateOverviewScanInProgress:
		return "ICScannerFunctionalUnitStateOverviewScanInProgress"
	case ICScannerFunctionalUnitStateReady:
		return "ICScannerFunctionalUnitStateReady"
	case ICScannerFunctionalUnitStateScanInProgress:
		return "ICScannerFunctionalUnitStateScanInProgress"
	default:
		return fmt.Sprintf("ICScannerFunctionalUnitState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnitType
type ICScannerFunctionalUnitType uint

const (
	// ICScannerFunctionalUnitTypeDocumentFeeder: A document feeder functional unit.
	ICScannerFunctionalUnitTypeDocumentFeeder ICScannerFunctionalUnitType = 3
	// ICScannerFunctionalUnitTypeFlatbed: A flatbed functional unit.
	ICScannerFunctionalUnitTypeFlatbed ICScannerFunctionalUnitType = 0
	// ICScannerFunctionalUnitTypeNegativeTransparency: A transparency functional unit for scanning negatives.
	ICScannerFunctionalUnitTypeNegativeTransparency ICScannerFunctionalUnitType = 2
	// ICScannerFunctionalUnitTypePositiveTransparency: A transparency functional unit for scanning positives.
	ICScannerFunctionalUnitTypePositiveTransparency ICScannerFunctionalUnitType = 1
)

func (e ICScannerFunctionalUnitType) String() string {
	switch e {
	case ICScannerFunctionalUnitTypeDocumentFeeder:
		return "ICScannerFunctionalUnitTypeDocumentFeeder"
	case ICScannerFunctionalUnitTypeFlatbed:
		return "ICScannerFunctionalUnitTypeFlatbed"
	case ICScannerFunctionalUnitTypeNegativeTransparency:
		return "ICScannerFunctionalUnitTypeNegativeTransparency"
	case ICScannerFunctionalUnitTypePositiveTransparency:
		return "ICScannerFunctionalUnitTypePositiveTransparency"
	default:
		return fmt.Sprintf("ICScannerFunctionalUnitType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerMeasurementUnit
type ICScannerMeasurementUnit uint

const (
	ICScannerMeasurementUnitCentimeters ICScannerMeasurementUnit = 1
	ICScannerMeasurementUnitInches      ICScannerMeasurementUnit = 0
	ICScannerMeasurementUnitPicas       ICScannerMeasurementUnit = 2
	ICScannerMeasurementUnitPixels      ICScannerMeasurementUnit = 5
	ICScannerMeasurementUnitPoints      ICScannerMeasurementUnit = 3
	ICScannerMeasurementUnitTwips       ICScannerMeasurementUnit = 4
)

func (e ICScannerMeasurementUnit) String() string {
	switch e {
	case ICScannerMeasurementUnitCentimeters:
		return "ICScannerMeasurementUnitCentimeters"
	case ICScannerMeasurementUnitInches:
		return "ICScannerMeasurementUnitInches"
	case ICScannerMeasurementUnitPicas:
		return "ICScannerMeasurementUnitPicas"
	case ICScannerMeasurementUnitPixels:
		return "ICScannerMeasurementUnitPixels"
	case ICScannerMeasurementUnitPoints:
		return "ICScannerMeasurementUnitPoints"
	case ICScannerMeasurementUnitTwips:
		return "ICScannerMeasurementUnitTwips"
	default:
		return fmt.Sprintf("ICScannerMeasurementUnit(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerPixelDataType
type ICScannerPixelDataType uint

const (
	ICScannerPixelDataTypeBW      ICScannerPixelDataType = 0
	ICScannerPixelDataTypeCIEXYZ  ICScannerPixelDataType = 8
	ICScannerPixelDataTypeCMY     ICScannerPixelDataType = 4
	ICScannerPixelDataTypeCMYK    ICScannerPixelDataType = 5
	ICScannerPixelDataTypeGray    ICScannerPixelDataType = 1
	ICScannerPixelDataTypePalette ICScannerPixelDataType = 3
	ICScannerPixelDataTypeRGB     ICScannerPixelDataType = 2
	ICScannerPixelDataTypeYUV     ICScannerPixelDataType = 6
	ICScannerPixelDataTypeYUVK    ICScannerPixelDataType = 7
)

func (e ICScannerPixelDataType) String() string {
	switch e {
	case ICScannerPixelDataTypeBW:
		return "ICScannerPixelDataTypeBW"
	case ICScannerPixelDataTypeCIEXYZ:
		return "ICScannerPixelDataTypeCIEXYZ"
	case ICScannerPixelDataTypeCMY:
		return "ICScannerPixelDataTypeCMY"
	case ICScannerPixelDataTypeCMYK:
		return "ICScannerPixelDataTypeCMYK"
	case ICScannerPixelDataTypeGray:
		return "ICScannerPixelDataTypeGray"
	case ICScannerPixelDataTypePalette:
		return "ICScannerPixelDataTypePalette"
	case ICScannerPixelDataTypeRGB:
		return "ICScannerPixelDataTypeRGB"
	case ICScannerPixelDataTypeYUV:
		return "ICScannerPixelDataTypeYUV"
	case ICScannerPixelDataTypeYUVK:
		return "ICScannerPixelDataTypeYUVK"
	default:
		return fmt.Sprintf("ICScannerPixelDataType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerTransferMode
type ICScannerTransferMode uint

const (
	// ICScannerTransferModeFileBased: The mode for transferring the scan as a file.
	ICScannerTransferModeFileBased ICScannerTransferMode = 0
	// ICScannerTransferModeMemoryBased: The mode for transferring the scan as data.
	ICScannerTransferModeMemoryBased ICScannerTransferMode = 1
)

func (e ICScannerTransferMode) String() string {
	switch e {
	case ICScannerTransferModeFileBased:
		return "ICScannerTransferModeFileBased"
	case ICScannerTransferModeMemoryBased:
		return "ICScannerTransferModeMemoryBased"
	default:
		return fmt.Sprintf("ICScannerTransferMode(%d)", e)
	}
}
