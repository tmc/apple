// Code generated from Apple documentation. DO NOT EDIT.

package imagecapturecore

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

var (
	// ICButtonTypeCopy is a nonlocalized notification string to indicate that the Copy button on the device was pressed.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICButtonTypeCopy
	ICButtonTypeCopy string
	// ICButtonTypeMail is a nonlocalized notification string to indicate that the Mail button on the device was pressed.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICButtonTypeMail
	ICButtonTypeMail string
	// ICButtonTypePrint is a nonlocalized notification string to indicate that the Print button on the device was pressed.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICButtonTypePrint
	ICButtonTypePrint string
	// ICButtonTypeScan is a nonlocalized notification string to indicate that the Scan button on the device was pressed.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICButtonTypeScan
	ICButtonTypeScan string
	// ICButtonTypeTransfer is a nonlocalized notification string to indicate that the Transfer button on the device was pressed.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICButtonTypeTransfer
	ICButtonTypeTransfer string
	// ICButtonTypeWeb is a nonlocalized notification string to indicate that the Web button on the device was pressed.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICButtonTypeWeb
	ICButtonTypeWeb string
	// ICScannerStatusRequestsOverviewScan is a nonlocalized notification string to indicate that the scanner is requesting an overview scan.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerStatusRequestsOverviewScan
	ICScannerStatusRequestsOverviewScan string
	// ICScannerStatusWarmUpDone is a nonlocalized notification string to indicate that the scanner has warmed up.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerStatusWarmUpDone
	ICScannerStatusWarmUpDone string
	// ICScannerStatusWarmingUp is a nonlocalized notification string to indicate that the scanner is warming up.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerStatusWarmingUp
	ICScannerStatusWarmingUp string
)

var (
	// ICCameraDeviceCanAcceptPTPCommands is indicates that the camera can accept PTP commands.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceCapability/cameraDeviceCanAcceptPTPCommands
	ICCameraDeviceCanAcceptPTPCommands ICDeviceCapability
	// ICCameraDeviceCanDeleteAllFiles is indicates that the camera can delete all files in a single operation while it is connected.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceCapability/cameraDeviceCanDeleteAllFiles
	ICCameraDeviceCanDeleteAllFiles ICDeviceCapability
	// ICCameraDeviceCanDeleteOneFile is indicates that the camera can delete a file at a time while it is connected.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceCapability/cameraDeviceCanDeleteOneFile
	ICCameraDeviceCanDeleteOneFile ICDeviceCapability
	// ICCameraDeviceCanReceiveFile is indicates that the host can upload files to the camera.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceCapability/cameraDeviceCanReceiveFile
	ICCameraDeviceCanReceiveFile ICDeviceCapability
	// ICCameraDeviceCanSyncClock is indicates that the camera can synchronize its date and time with that of the host computer.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceCapability/cameraDeviceCanSyncClock
	ICCameraDeviceCanSyncClock ICDeviceCapability
	// ICCameraDeviceCanTakePicture is the capability for the client to request to take a picture while the camera is connected.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceCapability/cameraDeviceCanTakePicture
	ICCameraDeviceCanTakePicture ICDeviceCapability
	// ICCameraDeviceCanTakePictureUsingShutterReleaseOnCamera is the capability to capture a picture if the user presses the shutter release on the camera while the camera is connected.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceCapability/cameraDeviceCanTakePictureUsingShutterReleaseOnCamera
	ICCameraDeviceCanTakePictureUsingShutterReleaseOnCamera ICDeviceCapability
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceCapability/cameraDeviceSupportsHEIF
	ICCameraDeviceSupportsHEIF ICDeviceCapability
	// ICDeviceCanEjectOrDisconnect is indicates that the camera can eject or disconnect.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceCapability/canEjectOrDisconnect
	ICDeviceCanEjectOrDisconnect ICDeviceCapability
)

var (
	// ICDeleteAfterSuccessfulDownload is a Boolean value indicating whether to delete the file from the device after a successful download.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDownloadOption/deleteAfterSuccessfulDownload
	ICDeleteAfterSuccessfulDownload ICDownloadOption
	// ICDownloadSidecarFiles is a Boolean value indicating whether to download all sidecar files along with the media file.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDownloadOption/sidecarFiles
	ICDownloadSidecarFiles ICDownloadOption
	// ICDownloadsDirectoryURL is a writable directory where the downloaded files should be saved.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDownloadOption/downloadsDirectoryURL
	ICDownloadsDirectoryURL ICDownloadOption
	// ICOverwrite is a Boolean value indicating whether the downloaded file should overwrite an existing file with the same name and extension.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDownloadOption/overwrite
	ICOverwrite ICDownloadOption
	// ICSaveAsFilename is the name to be used for the downloaded file.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDownloadOption/saveAsFilename
	ICSaveAsFilename ICDownloadOption
	// ICSavedAncillaryFiles is an array of files associated with the file being downloaded.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDownloadOption/savedAncillaryFiles
	ICSavedAncillaryFiles ICDownloadOption
	// ICSavedFilename is the actual name of the saved file.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDownloadOption/savedFilename
	ICSavedFilename ICDownloadOption
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDownloadOption/truncateAfterSuccessfulDownload
	ICTruncateAfterSuccessfulDownload ICDownloadOption
)

var (
	// ICDeleteCanceled is the deletion was canceled.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeleteResult/canceled
	ICDeleteCanceled ICDeleteResult
	// ICDeleteFailed is the deletion failed.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeleteResult/failed
	ICDeleteFailed ICDeleteResult
	// ICDeleteSuccessful is the deletion succeeded.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeleteResult/successful
	ICDeleteSuccessful ICDeleteResult
)

var ()

var (
	// ICDeviceLocationDescriptionBluetooth is a paired Bluetooth device.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceLocationOptions/descriptionBluetooth
	ICDeviceLocationDescriptionBluetooth ICDeviceLocationOptions
	// ICDeviceLocationDescriptionFireWire is a device that’s directly attached to the Mac through its FireWire port.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceLocationOptions/descriptionFireWire
	ICDeviceLocationDescriptionFireWire ICDeviceLocationOptions
	// ICDeviceLocationDescriptionMassStorage is a mass storage device.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceLocationOptions/descriptionMassStorage
	ICDeviceLocationDescriptionMassStorage ICDeviceLocationOptions
	// ICDeviceLocationDescriptionUSB is a device that’s directly attached to the Mac through its USB port.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceLocationOptions/descriptionUSB
	ICDeviceLocationDescriptionUSB ICDeviceLocationOptions
)

var (
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICSessionOptions/enumerationChronologicalOrder
	ICEnumerationChronologicalOrder ICSessionOptions
)

var (
	// ICErrorDomain is an error returned by the ImageCaptureCore framework.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICErrorDomain
	ICErrorDomain foundation.NSErrorDomain
)

var (
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItemThumbnailOption/imageSourceShouldCache
	ICImageSourceShouldCache ICCameraItemThumbnailOption
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItemThumbnailOption/imageSourceThumbnailMaxPixelSize
	ICImageSourceThumbnailMaxPixelSize ICCameraItemThumbnailOption
)

var (
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceStatus/localizedStatusNotificationKey
	ICLocalizedStatusNotificationKey ICDeviceStatus
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceStatus/statusCodeKey
	ICStatusCodeKey ICDeviceStatus
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceStatus/statusNotificationKey
	ICStatusNotificationKey ICDeviceStatus
)

var (
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceTransport/transportTypeBluetooth
	ICTransportTypeBluetooth ICDeviceTransport
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceTransport/transportTypeFireWire
	ICTransportTypeFireWire ICDeviceTransport
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceTransport/transportTypeMassStorage
	ICTransportTypeMassStorage ICDeviceTransport
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceTransport/transportTypeProximity
	ICTransportTypeProximity ICDeviceTransport
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceTransport/transportTypeTCPIP
	ICTransportTypeTCPIP ICDeviceTransport
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceTransport/transportTypeUSB
	ICTransportTypeUSB ICDeviceTransport
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICButtonTypeCopy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICButtonTypeCopy = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICButtonTypeMail"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICButtonTypeMail = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICButtonTypePrint"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICButtonTypePrint = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICButtonTypeScan"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICButtonTypeScan = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICButtonTypeTransfer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICButtonTypeTransfer = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICButtonTypeWeb"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICButtonTypeWeb = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICCameraDeviceCanAcceptPTPCommands"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICCameraDeviceCanAcceptPTPCommands = ICDeviceCapability(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICCameraDeviceCanDeleteAllFiles"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICCameraDeviceCanDeleteAllFiles = ICDeviceCapability(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICCameraDeviceCanDeleteOneFile"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICCameraDeviceCanDeleteOneFile = ICDeviceCapability(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICCameraDeviceCanReceiveFile"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICCameraDeviceCanReceiveFile = ICDeviceCapability(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICCameraDeviceCanSyncClock"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICCameraDeviceCanSyncClock = ICDeviceCapability(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICCameraDeviceCanTakePicture"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICCameraDeviceCanTakePicture = ICDeviceCapability(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICCameraDeviceCanTakePictureUsingShutterReleaseOnCamera"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICCameraDeviceCanTakePictureUsingShutterReleaseOnCamera = ICDeviceCapability(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICCameraDeviceSupportsHEIF"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICCameraDeviceSupportsHEIF = ICDeviceCapability(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICDeleteAfterSuccessfulDownload"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICDeleteAfterSuccessfulDownload = ICDownloadOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICDeleteCanceled"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICDeleteCanceled = ICDeleteResult(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICDeleteErrorCanceled"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICDeleteErrors.Canceled = ICDeleteError(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICDeleteErrorDeviceMissing"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICDeleteErrors.DeviceMissing = ICDeleteError(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICDeleteErrorFileMissing"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICDeleteErrors.FileMissing = ICDeleteError(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICDeleteErrorReadOnly"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICDeleteErrors.ReadOnly = ICDeleteError(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICDeleteFailed"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICDeleteFailed = ICDeleteResult(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICDeleteSuccessful"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICDeleteSuccessful = ICDeleteResult(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICDeviceCanEjectOrDisconnect"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICDeviceCanEjectOrDisconnect = ICDeviceCapability(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICDeviceLocationDescriptionBluetooth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICDeviceLocationDescriptionBluetooth = ICDeviceLocationOptions(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICDeviceLocationDescriptionFireWire"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICDeviceLocationDescriptionFireWire = ICDeviceLocationOptions(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICDeviceLocationDescriptionMassStorage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICDeviceLocationDescriptionMassStorage = ICDeviceLocationOptions(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICDeviceLocationDescriptionUSB"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICDeviceLocationDescriptionUSB = ICDeviceLocationOptions(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICDownloadSidecarFiles"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICDownloadSidecarFiles = ICDownloadOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICDownloadsDirectoryURL"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICDownloadsDirectoryURL = ICDownloadOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICEnumerationChronologicalOrder"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICEnumerationChronologicalOrder = ICSessionOptions(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICErrorDomain = foundation.NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICImageSourceShouldCache"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICImageSourceShouldCache = ICCameraItemThumbnailOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICImageSourceThumbnailMaxPixelSize"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICImageSourceThumbnailMaxPixelSize = ICCameraItemThumbnailOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICLocalizedStatusNotificationKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICLocalizedStatusNotificationKey = ICDeviceStatus(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICOverwrite"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICOverwrite = ICDownloadOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICSaveAsFilename"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICSaveAsFilename = ICDownloadOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICSavedAncillaryFiles"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICSavedAncillaryFiles = ICDownloadOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICSavedFilename"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICSavedFilename = ICDownloadOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICScannerStatusRequestsOverviewScan"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICScannerStatusRequestsOverviewScan = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICScannerStatusWarmUpDone"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICScannerStatusWarmUpDone = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICScannerStatusWarmingUp"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICScannerStatusWarmingUp = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICStatusCodeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICStatusCodeKey = ICDeviceStatus(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICStatusNotificationKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICStatusNotificationKey = ICDeviceStatus(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICTransportTypeBluetooth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICTransportTypeBluetooth = ICDeviceTransport(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICTransportTypeFireWire"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICTransportTypeFireWire = ICDeviceTransport(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICTransportTypeMassStorage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICTransportTypeMassStorage = ICDeviceTransport(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICTransportTypeProximity"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICTransportTypeProximity = ICDeviceTransport(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICTransportTypeTCPIP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICTransportTypeTCPIP = ICDeviceTransport(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICTransportTypeUSB"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICTransportTypeUSB = ICDeviceTransport(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ICTruncateAfterSuccessfulDownload"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ICTruncateAfterSuccessfulDownload = ICDownloadOption(objc.GoString(cstr))
			}
		}
	}

}

// ICDeleteErrors provides typed accessors for [ICDeleteError] constants.
var ICDeleteErrors struct {
	// Canceled: The deletion was canceled.
	Canceled ICDeleteError
	// DeviceMissing: The deletion failed because the device could not be found.
	DeviceMissing ICDeleteError
	// FileMissing: The deletion failed because the file could not be found.
	FileMissing ICDeleteError
	// ReadOnly: The deletion failed because the file had read-only permissions.
	ReadOnly ICDeleteError
}
