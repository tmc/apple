// Code generated from Apple documentation. DO NOT EDIT.

package coremediaio

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

var (
	// CMIOExtensionInfoDictionaryKey is a key that specifies the extension information dictionary.
	//
	// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOExtensionInfoDictionaryKey
	CMIOExtensionInfoDictionaryKey string
	// CMIOExtensionMachServiceNameKey is a key that specifies the mach service name.
	//
	// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOExtensionMachServiceNameKey
	CMIOExtensionMachServiceNameKey string
)

var ()

var (
	// See: https://developer.apple.com/documentation/CoreMediaIO/kCMIOBlockBufferAttachmentKey_CVPixelBufferReference
	KCMIOBlockBufferAttachmentKey_CVPixelBufferReference string
	// See: https://developer.apple.com/documentation/CoreMediaIO/kCMIOSampleBufferAttachmentKey_CAAudioTimeStamp
	KCMIOSampleBufferAttachmentKey_CAAudioTimeStamp string
	// See: https://developer.apple.com/documentation/CoreMediaIO/kCMIOSampleBufferAttachmentKey_ClientSequenceID
	KCMIOSampleBufferAttachmentKey_ClientSequenceID string
	// See: https://developer.apple.com/documentation/CoreMediaIO/kCMIOSampleBufferAttachmentKey_ClosedCaptionSampleBuffer
	KCMIOSampleBufferAttachmentKey_ClosedCaptionSampleBuffer string
	// See: https://developer.apple.com/documentation/CoreMediaIO/kCMIOSampleBufferAttachmentKey_DiscontinuityFlags
	KCMIOSampleBufferAttachmentKey_DiscontinuityFlags string
	// See: https://developer.apple.com/documentation/CoreMediaIO/kCMIOSampleBufferAttachmentKey_HDV1_PackData
	KCMIOSampleBufferAttachmentKey_HDV1_PackData string
	// See: https://developer.apple.com/documentation/CoreMediaIO/kCMIOSampleBufferAttachmentKey_HDV2_VAUX
	KCMIOSampleBufferAttachmentKey_HDV2_VAUX string
	// See: https://developer.apple.com/documentation/CoreMediaIO/kCMIOSampleBufferAttachmentKey_HostTime
	KCMIOSampleBufferAttachmentKey_HostTime string
	// See: https://developer.apple.com/documentation/CoreMediaIO/kCMIOSampleBufferAttachmentKey_MouseAndKeyboardModifiers
	KCMIOSampleBufferAttachmentKey_MouseAndKeyboardModifiers string
	// See: https://developer.apple.com/documentation/CoreMediaIO/kCMIOSampleBufferAttachmentKey_MuxedSourcePresentationTimeStamp
	KCMIOSampleBufferAttachmentKey_MuxedSourcePresentationTimeStamp string
	// See: https://developer.apple.com/documentation/CoreMediaIO/kCMIOSampleBufferAttachmentKey_NativeSMPTEFrameCount
	KCMIOSampleBufferAttachmentKey_NativeSMPTEFrameCount string
	// See: https://developer.apple.com/documentation/CoreMediaIO/kCMIOSampleBufferAttachmentKey_NoDataMarker
	KCMIOSampleBufferAttachmentKey_NoDataMarker string
	// See: https://developer.apple.com/documentation/CoreMediaIO/kCMIOSampleBufferAttachmentKey_NumberOfVideoFramesInBuffer
	KCMIOSampleBufferAttachmentKey_NumberOfVideoFramesInBuffer string
	// See: https://developer.apple.com/documentation/CoreMediaIO/kCMIOSampleBufferAttachmentKey_NumberOfVideoFramesInGOP
	KCMIOSampleBufferAttachmentKey_NumberOfVideoFramesInGOP string
	// See: https://developer.apple.com/documentation/CoreMediaIO/kCMIOSampleBufferAttachmentKey_PixelBufferOverlaidByStaticImage
	KCMIOSampleBufferAttachmentKey_PixelBufferOverlaidByStaticImage string
	// See: https://developer.apple.com/documentation/CoreMediaIO/kCMIOSampleBufferAttachmentKey_PulldownCadenceInfo
	KCMIOSampleBufferAttachmentKey_PulldownCadenceInfo string
	// See: https://developer.apple.com/documentation/CoreMediaIO/kCMIOSampleBufferAttachmentKey_RepeatedBufferContents
	KCMIOSampleBufferAttachmentKey_RepeatedBufferContents string
	// See: https://developer.apple.com/documentation/CoreMediaIO/kCMIOSampleBufferAttachmentKey_SMPTETime
	KCMIOSampleBufferAttachmentKey_SMPTETime string
	// See: https://developer.apple.com/documentation/CoreMediaIO/kCMIOSampleBufferAttachmentKey_SequenceNumber
	KCMIOSampleBufferAttachmentKey_SequenceNumber string
	// See: https://developer.apple.com/documentation/CoreMediaIO/kCMIOSampleBufferAttachmentKey_SourceAudioFormatDescription
	KCMIOSampleBufferAttachmentKey_SourceAudioFormatDescription string
	// See: https://developer.apple.com/documentation/CoreMediaIO/kCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_CursorFrameRect
	KCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_CursorFrameRect string
	// See: https://developer.apple.com/documentation/CoreMediaIO/kCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_CursorIsDrawnInFramebuffer
	KCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_CursorIsDrawnInFramebuffer string
	// See: https://developer.apple.com/documentation/CoreMediaIO/kCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_CursorIsVisible
	KCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_CursorIsVisible string
	// See: https://developer.apple.com/documentation/CoreMediaIO/kCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_CursorPositionX
	KCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_CursorPositionX string
	// See: https://developer.apple.com/documentation/CoreMediaIO/kCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_CursorPositionY
	KCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_CursorPositionY string
	// See: https://developer.apple.com/documentation/CoreMediaIO/kCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_CursorReference
	KCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_CursorReference string
	// See: https://developer.apple.com/documentation/CoreMediaIO/kCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_CursorScale
	KCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_CursorScale string
	// See: https://developer.apple.com/documentation/CoreMediaIO/kCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_CursorSeed
	KCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_CursorSeed string
	// See: https://developer.apple.com/documentation/CoreMediaIO/kCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_KeyboardModifiers
	KCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_KeyboardModifiers string
	// See: https://developer.apple.com/documentation/CoreMediaIO/kCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_KeyboardModifiersEvent
	KCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_KeyboardModifiersEvent string
	// See: https://developer.apple.com/documentation/CoreMediaIO/kCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_MouseButtonState
	KCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_MouseButtonState string
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CMIOExtensionInfoDictionaryKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CMIOExtensionInfoDictionaryKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CMIOExtensionMachServiceNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CMIOExtensionMachServiceNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CMIOExtensionPropertyDeviceCanBeDefaultInputDevice"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CMIOExtensionPropertys.DeviceCanBeDefaultInputDevice = CMIOExtensionProperty(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CMIOExtensionPropertyDeviceCanBeDefaultOutputDevice"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CMIOExtensionPropertys.DeviceCanBeDefaultOutputDevice = CMIOExtensionProperty(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CMIOExtensionPropertyDeviceIsSuspended"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CMIOExtensionPropertys.DeviceIsSuspended = CMIOExtensionProperty(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CMIOExtensionPropertyDeviceLatency"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CMIOExtensionPropertys.DeviceLatency = CMIOExtensionProperty(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CMIOExtensionPropertyDeviceLinkedCoreAudioDeviceUID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CMIOExtensionPropertys.DeviceLinkedCoreAudioDeviceUID = CMIOExtensionProperty(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CMIOExtensionPropertyDeviceModel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CMIOExtensionPropertys.DeviceModel = CMIOExtensionProperty(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CMIOExtensionPropertyDeviceTransportType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CMIOExtensionPropertys.DeviceTransportType = CMIOExtensionProperty(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CMIOExtensionPropertyProviderManufacturer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CMIOExtensionPropertys.ProviderManufacturer = CMIOExtensionProperty(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CMIOExtensionPropertyProviderName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CMIOExtensionPropertys.ProviderName = CMIOExtensionProperty(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CMIOExtensionPropertyStreamActiveFormatIndex"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CMIOExtensionPropertys.StreamActiveFormatIndex = CMIOExtensionProperty(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CMIOExtensionPropertyStreamFrameDuration"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CMIOExtensionPropertys.StreamFrameDuration = CMIOExtensionProperty(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CMIOExtensionPropertyStreamLatency"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CMIOExtensionPropertys.StreamLatency = CMIOExtensionProperty(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CMIOExtensionPropertyStreamMaxFrameDuration"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CMIOExtensionPropertys.StreamMaxFrameDuration = CMIOExtensionProperty(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CMIOExtensionPropertyStreamSinkBufferQueueSize"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CMIOExtensionPropertys.StreamSinkBufferQueueSize = CMIOExtensionProperty(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CMIOExtensionPropertyStreamSinkBufferUnderrunCount"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CMIOExtensionPropertys.StreamSinkBufferUnderrunCount = CMIOExtensionProperty(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CMIOExtensionPropertyStreamSinkBuffersRequiredForStartup"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CMIOExtensionPropertys.StreamSinkBuffersRequiredForStartup = CMIOExtensionProperty(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CMIOExtensionPropertyStreamSinkEndOfData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CMIOExtensionPropertys.StreamSinkEndOfData = CMIOExtensionProperty(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMIOBlockBufferAttachmentKey_CVPixelBufferReference"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMIOBlockBufferAttachmentKey_CVPixelBufferReference = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMIOSampleBufferAttachmentKey_CAAudioTimeStamp"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMIOSampleBufferAttachmentKey_CAAudioTimeStamp = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMIOSampleBufferAttachmentKey_ClientSequenceID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMIOSampleBufferAttachmentKey_ClientSequenceID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMIOSampleBufferAttachmentKey_ClosedCaptionSampleBuffer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMIOSampleBufferAttachmentKey_ClosedCaptionSampleBuffer = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMIOSampleBufferAttachmentKey_DiscontinuityFlags"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMIOSampleBufferAttachmentKey_DiscontinuityFlags = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMIOSampleBufferAttachmentKey_HDV1_PackData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMIOSampleBufferAttachmentKey_HDV1_PackData = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMIOSampleBufferAttachmentKey_HDV2_VAUX"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMIOSampleBufferAttachmentKey_HDV2_VAUX = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMIOSampleBufferAttachmentKey_HostTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMIOSampleBufferAttachmentKey_HostTime = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMIOSampleBufferAttachmentKey_MouseAndKeyboardModifiers"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMIOSampleBufferAttachmentKey_MouseAndKeyboardModifiers = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMIOSampleBufferAttachmentKey_MuxedSourcePresentationTimeStamp"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMIOSampleBufferAttachmentKey_MuxedSourcePresentationTimeStamp = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMIOSampleBufferAttachmentKey_NativeSMPTEFrameCount"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMIOSampleBufferAttachmentKey_NativeSMPTEFrameCount = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMIOSampleBufferAttachmentKey_NoDataMarker"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMIOSampleBufferAttachmentKey_NoDataMarker = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMIOSampleBufferAttachmentKey_NumberOfVideoFramesInBuffer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMIOSampleBufferAttachmentKey_NumberOfVideoFramesInBuffer = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMIOSampleBufferAttachmentKey_NumberOfVideoFramesInGOP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMIOSampleBufferAttachmentKey_NumberOfVideoFramesInGOP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMIOSampleBufferAttachmentKey_PixelBufferOverlaidByStaticImage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMIOSampleBufferAttachmentKey_PixelBufferOverlaidByStaticImage = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMIOSampleBufferAttachmentKey_PulldownCadenceInfo"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMIOSampleBufferAttachmentKey_PulldownCadenceInfo = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMIOSampleBufferAttachmentKey_RepeatedBufferContents"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMIOSampleBufferAttachmentKey_RepeatedBufferContents = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMIOSampleBufferAttachmentKey_SMPTETime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMIOSampleBufferAttachmentKey_SMPTETime = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMIOSampleBufferAttachmentKey_SequenceNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMIOSampleBufferAttachmentKey_SequenceNumber = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMIOSampleBufferAttachmentKey_SourceAudioFormatDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMIOSampleBufferAttachmentKey_SourceAudioFormatDescription = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_CursorFrameRect"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_CursorFrameRect = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_CursorIsDrawnInFramebuffer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_CursorIsDrawnInFramebuffer = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_CursorIsVisible"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_CursorIsVisible = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_CursorPositionX"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_CursorPositionX = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_CursorPositionY"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_CursorPositionY = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_CursorReference"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_CursorReference = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_CursorScale"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_CursorScale = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_CursorSeed"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_CursorSeed = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_KeyboardModifiers"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_KeyboardModifiers = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_KeyboardModifiersEvent"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_KeyboardModifiersEvent = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_MouseButtonState"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMIOSampleBufferAttachment_MouseAndKeyboardModifiersKey_MouseButtonState = objc.GoString(cstr)
			}
		}
	}

}

// CMIOExtensionPropertys provides typed accessors for [CMIOExtensionProperty] constants.
var CMIOExtensionPropertys struct {
	// DeviceCanBeDefaultInputDevice: A property key for a Boolean value that indicates whether the device can be a default input device.
	DeviceCanBeDefaultInputDevice CMIOExtensionProperty
	// DeviceCanBeDefaultOutputDevice: A property key for a Boolean value that indicates whether the device can be a default output device.
	DeviceCanBeDefaultOutputDevice CMIOExtensionProperty
	// DeviceIsSuspended: A property key for a Boolean value that indicates whether the device is in a suspended state.
	DeviceIsSuspended CMIOExtensionProperty
	DeviceLatency     CMIOExtensionProperty
	// DeviceLinkedCoreAudioDeviceUID: A property key for the UID of the linked Core Audio device.
	DeviceLinkedCoreAudioDeviceUID CMIOExtensionProperty
	// DeviceModel: A property key for the device model.
	DeviceModel CMIOExtensionProperty
	// DeviceTransportType: A property key for the device transport type.
	DeviceTransportType CMIOExtensionProperty
	// ProviderManufacturer: A property key for the provider manufacturer.
	ProviderManufacturer CMIOExtensionProperty
	// ProviderName: A property key for the provider name.
	ProviderName CMIOExtensionProperty
	// StreamActiveFormatIndex: A property key for the index of the active stream format.
	StreamActiveFormatIndex CMIOExtensionProperty
	// StreamFrameDuration: A property key for the frame duration.
	StreamFrameDuration CMIOExtensionProperty
	StreamLatency       CMIOExtensionProperty
	// StreamMaxFrameDuration: A property key for the maximum frame duration.
	StreamMaxFrameDuration CMIOExtensionProperty
	// StreamSinkBufferQueueSize: A property key for the sink buffer queue size.
	StreamSinkBufferQueueSize CMIOExtensionProperty
	// StreamSinkBufferUnderrunCount: A property key for the buffer underrun count.
	StreamSinkBufferUnderrunCount CMIOExtensionProperty
	// StreamSinkBuffersRequiredForStartup: A property key for the number of buffers required for startup.
	StreamSinkBuffersRequiredForStartup CMIOExtensionProperty
	// StreamSinkEndOfData: A property key for a Boolean value that indicates whether the stream has more data.
	StreamSinkEndOfData CMIOExtensionProperty
}
