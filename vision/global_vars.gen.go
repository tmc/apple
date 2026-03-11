// Code generated from Apple documentation. DO NOT EDIT.

package vision

import (
	"unsafe"
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
)

const (

	VNCalculateImageAestheticsScoresRequestRevision1 uint = 1

	VNClassifyImageRequestRevision1 uint = 1

	VNClassifyImageRequestRevision2 uint = 2

	VNCoreMLRequestRevision1 uint = 1

	VNDetectAnimalBodyPoseRequestRevision1 uint = 1

//
// Deprecated: Deprecated since macOS 14.0. Use [VNDetectBarcodesRequestRevision3](<doc://Vision/documentation/Vision/VNDetectBarcodesRequestRevision3>) instead.
	VNDetectBarcodesRequestRevision1 uint = 1

	VNDetectBarcodesRequestRevision2 uint = 2

	VNDetectBarcodesRequestRevision3 uint = 3

	VNDetectBarcodesRequestRevision4 uint = 4

	VNDetectContourRequestRevision1 uint = 1

	VNDetectDocumentSegmentationRequestRevision1 uint = 1

	VNDetectFaceCaptureQualityRequestRevision1 uint = 1

	VNDetectFaceCaptureQualityRequestRevision2 uint = 2

	VNDetectFaceCaptureQualityRequestRevision3 uint = 3

	VNDetectFaceLandmarksRequestRevision2 uint = 2

	VNDetectFaceLandmarksRequestRevision3 uint = 3

	VNDetectFaceRectanglesRequestRevision2 uint = 2

	VNDetectFaceRectanglesRequestRevision3 uint = 3

	VNDetectHorizonRequestRevision1 uint = 1

	VNDetectHumanBodyPose3DRequestRevision1 uint = 1

	VNDetectHumanBodyPoseRequestRevision1 uint = 1

	VNDetectHumanHandPoseRequestRevision1 uint = 1

	VNDetectHumanRectanglesRequestRevision1 uint = 1

	VNDetectHumanRectanglesRequestRevision2 uint = 2

	VNDetectRectanglesRequestRevision1 uint = 1

	VNDetectTextRectanglesRequestRevision1 uint = 1

	VNDetectTrajectoriesRequestRevision1 uint = 1

	VNGenerateAttentionBasedSaliencyImageRequestRevision1 uint = 1

	VNGenerateAttentionBasedSaliencyImageRequestRevision2 uint = 2

	VNGenerateForegroundInstanceMaskRequestRevision1 uint = 1

	VNGenerateImageFeaturePrintRequestRevision1 uint = 1

	VNGenerateImageFeaturePrintRequestRevision2 uint = 2

	VNGenerateObjectnessBasedSaliencyImageRequestRevision1 uint = 1

	VNGenerateObjectnessBasedSaliencyImageRequestRevision2 uint = 2

	VNGenerateOpticalFlowRequestRevision1 uint = 1

	VNGenerateOpticalFlowRequestRevision2 uint = 2

	VNGeneratePersonInstanceMaskRequestRevision1 uint = 1

	VNGeneratePersonSegmentationRequestRevision1 uint = 1

	VNHomographicImageRegistrationRequestRevision1 uint = 1

	VNRecognizeAnimalsRequestRevision1 uint = 1

	VNRecognizeAnimalsRequestRevision2 uint = 2

	VNRecognizeTextRequestRevision1 uint = 1

	VNRecognizeTextRequestRevision2 uint = 2

	VNRecognizeTextRequestRevision3 uint = 3

	VNRequestRevisionUnspecified uint = 0

	VNTrackHomographicImageRegistrationRequestRevision1 uint = 1

	VNTrackObjectRequestRevision1 uint = 1

	VNTrackObjectRequestRevision2 uint = 2

	VNTrackOpticalFlowRequestRevision1 uint = 1

	VNTrackRectangleRequestRevision1 uint = 1

	VNTrackTranslationalImageRegistrationRequestRevision1 uint = 1

	VNTranslationalImageRegistrationRequestRevision1 uint = 1
)























































































var VNErrorDomain string





























































































var VNNormalizedIdentityRect corefoundation.CGRect






var VNRecognizedPoint3DGroupKeyAll VNRecognizedPointGroupKey

var VNRecognizedPointGroupKeyAll VNRecognizedPointGroupKey









var VNVisionVersionNumber float64

func init() {
	if frameworkHandle == 0 {
		return
	}


	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameLeftBackElbow"); err == nil && ptr != 0 {
		VNAnimalBodyPoseObservationJointNames.LeftBackElbow = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameLeftBackKnee"); err == nil && ptr != 0 {
		VNAnimalBodyPoseObservationJointNames.LeftBackKnee = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameLeftBackPaw"); err == nil && ptr != 0 {
		VNAnimalBodyPoseObservationJointNames.LeftBackPaw = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameLeftEarBottom"); err == nil && ptr != 0 {
		VNAnimalBodyPoseObservationJointNames.LeftEarBottom = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameLeftEarMiddle"); err == nil && ptr != 0 {
		VNAnimalBodyPoseObservationJointNames.LeftEarMiddle = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameLeftEarTop"); err == nil && ptr != 0 {
		VNAnimalBodyPoseObservationJointNames.LeftEarTop = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameLeftEye"); err == nil && ptr != 0 {
		VNAnimalBodyPoseObservationJointNames.LeftEye = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameLeftFrontElbow"); err == nil && ptr != 0 {
		VNAnimalBodyPoseObservationJointNames.LeftFrontElbow = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameLeftFrontKnee"); err == nil && ptr != 0 {
		VNAnimalBodyPoseObservationJointNames.LeftFrontKnee = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameLeftFrontPaw"); err == nil && ptr != 0 {
		VNAnimalBodyPoseObservationJointNames.LeftFrontPaw = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameNeck"); err == nil && ptr != 0 {
		VNAnimalBodyPoseObservationJointNames.Neck = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameNose"); err == nil && ptr != 0 {
		VNAnimalBodyPoseObservationJointNames.Nose = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameRightBackElbow"); err == nil && ptr != 0 {
		VNAnimalBodyPoseObservationJointNames.RightBackElbow = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameRightBackKnee"); err == nil && ptr != 0 {
		VNAnimalBodyPoseObservationJointNames.RightBackKnee = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameRightBackPaw"); err == nil && ptr != 0 {
		VNAnimalBodyPoseObservationJointNames.RightBackPaw = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameRightEarBottom"); err == nil && ptr != 0 {
		VNAnimalBodyPoseObservationJointNames.RightEarBottom = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameRightEarMiddle"); err == nil && ptr != 0 {
		VNAnimalBodyPoseObservationJointNames.RightEarMiddle = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameRightEarTop"); err == nil && ptr != 0 {
		VNAnimalBodyPoseObservationJointNames.RightEarTop = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameRightEye"); err == nil && ptr != 0 {
		VNAnimalBodyPoseObservationJointNames.RightEye = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameRightFrontElbow"); err == nil && ptr != 0 {
		VNAnimalBodyPoseObservationJointNames.RightFrontElbow = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameRightFrontKnee"); err == nil && ptr != 0 {
		VNAnimalBodyPoseObservationJointNames.RightFrontKnee = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameRightFrontPaw"); err == nil && ptr != 0 {
		VNAnimalBodyPoseObservationJointNames.RightFrontPaw = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameTailBottom"); err == nil && ptr != 0 {
		VNAnimalBodyPoseObservationJointNames.TailBottom = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameTailMiddle"); err == nil && ptr != 0 {
		VNAnimalBodyPoseObservationJointNames.TailMiddle = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameTailTop"); err == nil && ptr != 0 {
		VNAnimalBodyPoseObservationJointNames.TailTop = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointsGroupNameAll"); err == nil && ptr != 0 {
		VNAnimalBodyPoseObservationJointsGroupNames.All = *(*VNAnimalBodyPoseObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointsGroupNameForelegs"); err == nil && ptr != 0 {
		VNAnimalBodyPoseObservationJointsGroupNames.Forelegs = *(*VNAnimalBodyPoseObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointsGroupNameHead"); err == nil && ptr != 0 {
		VNAnimalBodyPoseObservationJointsGroupNames.Head = *(*VNAnimalBodyPoseObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointsGroupNameHindlegs"); err == nil && ptr != 0 {
		VNAnimalBodyPoseObservationJointsGroupNames.Hindlegs = *(*VNAnimalBodyPoseObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointsGroupNameTail"); err == nil && ptr != 0 {
		VNAnimalBodyPoseObservationJointsGroupNames.Tail = *(*VNAnimalBodyPoseObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointsGroupNameTrunk"); err == nil && ptr != 0 {
		VNAnimalBodyPoseObservationJointsGroupNames.Trunk = *(*VNAnimalBodyPoseObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalIdentifierCat"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VNAnimalIdentifiers.Cat = VNAnimalIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalIdentifierDog"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VNAnimalIdentifiers.Dog = VNAnimalIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyAztec"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VNBarcodeSymbologys.Aztec = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyCodabar"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VNBarcodeSymbologys.Codabar = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyCode128"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VNBarcodeSymbologys.Code128 = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyCode39"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VNBarcodeSymbologys.Code39 = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyCode39Checksum"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VNBarcodeSymbologys.Code39Checksum = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyCode39FullASCII"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VNBarcodeSymbologys.Code39FullASCII = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyCode39FullASCIIChecksum"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VNBarcodeSymbologys.Code39FullASCIIChecksum = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyCode93"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VNBarcodeSymbologys.Code93 = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyCode93i"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VNBarcodeSymbologys.Code93i = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyDataMatrix"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VNBarcodeSymbologys.DataMatrix = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyEAN13"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VNBarcodeSymbologys.EAN13 = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyEAN8"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VNBarcodeSymbologys.EAN8 = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyGS1DataBar"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VNBarcodeSymbologys.GS1DataBar = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyGS1DataBarExpanded"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VNBarcodeSymbologys.GS1DataBarExpanded = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyGS1DataBarLimited"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VNBarcodeSymbologys.GS1DataBarLimited = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyI2of5"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VNBarcodeSymbologys.I2of5 = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyI2of5Checksum"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VNBarcodeSymbologys.I2of5Checksum = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyITF14"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VNBarcodeSymbologys.ITF14 = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyMSIPlessey"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VNBarcodeSymbologys.MSIPlessey = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyMicroPDF417"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VNBarcodeSymbologys.MicroPDF417 = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyMicroQR"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VNBarcodeSymbologys.MicroQR = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyPDF417"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VNBarcodeSymbologys.PDF417 = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyQR"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VNBarcodeSymbologys.QR = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyUPCE"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VNBarcodeSymbologys.UPCE = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}




	if ptr, err := purego.Dlsym(frameworkHandle, "VNComputeStageMain"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VNComputeStages.Main = VNComputeStage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNComputeStagePostProcessing"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VNComputeStages.PostProcessing = VNComputeStage(objc.GoString(cstr))
			}
		}
	}

























	if ptr, err := purego.Dlsym(frameworkHandle, "VNErrorDomain"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VNErrorDomain = objc.GoString(cstr)
			}
		}
	}













	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointNameCenterHead"); err == nil && ptr != 0 {
		VNHumanBodyPose3DObservationJointNames.CenterHead = *(*VNHumanBodyPose3DObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointNameCenterShoulder"); err == nil && ptr != 0 {
		VNHumanBodyPose3DObservationJointNames.CenterShoulder = *(*VNHumanBodyPose3DObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointNameLeftAnkle"); err == nil && ptr != 0 {
		VNHumanBodyPose3DObservationJointNames.LeftAnkle = *(*VNHumanBodyPose3DObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointNameLeftElbow"); err == nil && ptr != 0 {
		VNHumanBodyPose3DObservationJointNames.LeftElbow = *(*VNHumanBodyPose3DObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointNameLeftHip"); err == nil && ptr != 0 {
		VNHumanBodyPose3DObservationJointNames.LeftHip = *(*VNHumanBodyPose3DObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointNameLeftKnee"); err == nil && ptr != 0 {
		VNHumanBodyPose3DObservationJointNames.LeftKnee = *(*VNHumanBodyPose3DObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointNameLeftShoulder"); err == nil && ptr != 0 {
		VNHumanBodyPose3DObservationJointNames.LeftShoulder = *(*VNHumanBodyPose3DObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointNameLeftWrist"); err == nil && ptr != 0 {
		VNHumanBodyPose3DObservationJointNames.LeftWrist = *(*VNHumanBodyPose3DObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointNameRightAnkle"); err == nil && ptr != 0 {
		VNHumanBodyPose3DObservationJointNames.RightAnkle = *(*VNHumanBodyPose3DObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointNameRightElbow"); err == nil && ptr != 0 {
		VNHumanBodyPose3DObservationJointNames.RightElbow = *(*VNHumanBodyPose3DObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointNameRightHip"); err == nil && ptr != 0 {
		VNHumanBodyPose3DObservationJointNames.RightHip = *(*VNHumanBodyPose3DObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointNameRightKnee"); err == nil && ptr != 0 {
		VNHumanBodyPose3DObservationJointNames.RightKnee = *(*VNHumanBodyPose3DObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointNameRightShoulder"); err == nil && ptr != 0 {
		VNHumanBodyPose3DObservationJointNames.RightShoulder = *(*VNHumanBodyPose3DObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointNameRightWrist"); err == nil && ptr != 0 {
		VNHumanBodyPose3DObservationJointNames.RightWrist = *(*VNHumanBodyPose3DObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointNameRoot"); err == nil && ptr != 0 {
		VNHumanBodyPose3DObservationJointNames.Root = *(*VNHumanBodyPose3DObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointNameSpine"); err == nil && ptr != 0 {
		VNHumanBodyPose3DObservationJointNames.Spine = *(*VNHumanBodyPose3DObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointNameTopHead"); err == nil && ptr != 0 {
		VNHumanBodyPose3DObservationJointNames.TopHead = *(*VNHumanBodyPose3DObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointsGroupNameAll"); err == nil && ptr != 0 {
		VNHumanBodyPose3DObservationJointsGroupNames.All = *(*VNHumanBodyPose3DObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointsGroupNameHead"); err == nil && ptr != 0 {
		VNHumanBodyPose3DObservationJointsGroupNames.Head = *(*VNHumanBodyPose3DObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointsGroupNameLeftArm"); err == nil && ptr != 0 {
		VNHumanBodyPose3DObservationJointsGroupNames.LeftArm = *(*VNHumanBodyPose3DObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointsGroupNameLeftLeg"); err == nil && ptr != 0 {
		VNHumanBodyPose3DObservationJointsGroupNames.LeftLeg = *(*VNHumanBodyPose3DObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointsGroupNameRightArm"); err == nil && ptr != 0 {
		VNHumanBodyPose3DObservationJointsGroupNames.RightArm = *(*VNHumanBodyPose3DObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointsGroupNameRightLeg"); err == nil && ptr != 0 {
		VNHumanBodyPose3DObservationJointsGroupNames.RightLeg = *(*VNHumanBodyPose3DObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointsGroupNameTorso"); err == nil && ptr != 0 {
		VNHumanBodyPose3DObservationJointsGroupNames.Torso = *(*VNHumanBodyPose3DObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointNameLeftAnkle"); err == nil && ptr != 0 {
		VNHumanBodyPoseObservationJointNames.LeftAnkle = *(*VNHumanBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointNameLeftEar"); err == nil && ptr != 0 {
		VNHumanBodyPoseObservationJointNames.LeftEar = *(*VNHumanBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointNameLeftElbow"); err == nil && ptr != 0 {
		VNHumanBodyPoseObservationJointNames.LeftElbow = *(*VNHumanBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointNameLeftEye"); err == nil && ptr != 0 {
		VNHumanBodyPoseObservationJointNames.LeftEye = *(*VNHumanBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointNameLeftHip"); err == nil && ptr != 0 {
		VNHumanBodyPoseObservationJointNames.LeftHip = *(*VNHumanBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointNameLeftKnee"); err == nil && ptr != 0 {
		VNHumanBodyPoseObservationJointNames.LeftKnee = *(*VNHumanBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointNameLeftShoulder"); err == nil && ptr != 0 {
		VNHumanBodyPoseObservationJointNames.LeftShoulder = *(*VNHumanBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointNameLeftWrist"); err == nil && ptr != 0 {
		VNHumanBodyPoseObservationJointNames.LeftWrist = *(*VNHumanBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointNameNeck"); err == nil && ptr != 0 {
		VNHumanBodyPoseObservationJointNames.Neck = *(*VNHumanBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointNameNose"); err == nil && ptr != 0 {
		VNHumanBodyPoseObservationJointNames.Nose = *(*VNHumanBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointNameRightAnkle"); err == nil && ptr != 0 {
		VNHumanBodyPoseObservationJointNames.RightAnkle = *(*VNHumanBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointNameRightEar"); err == nil && ptr != 0 {
		VNHumanBodyPoseObservationJointNames.RightEar = *(*VNHumanBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointNameRightElbow"); err == nil && ptr != 0 {
		VNHumanBodyPoseObservationJointNames.RightElbow = *(*VNHumanBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointNameRightEye"); err == nil && ptr != 0 {
		VNHumanBodyPoseObservationJointNames.RightEye = *(*VNHumanBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointNameRightHip"); err == nil && ptr != 0 {
		VNHumanBodyPoseObservationJointNames.RightHip = *(*VNHumanBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointNameRightKnee"); err == nil && ptr != 0 {
		VNHumanBodyPoseObservationJointNames.RightKnee = *(*VNHumanBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointNameRightShoulder"); err == nil && ptr != 0 {
		VNHumanBodyPoseObservationJointNames.RightShoulder = *(*VNHumanBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointNameRightWrist"); err == nil && ptr != 0 {
		VNHumanBodyPoseObservationJointNames.RightWrist = *(*VNHumanBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointNameRoot"); err == nil && ptr != 0 {
		VNHumanBodyPoseObservationJointNames.Root = *(*VNHumanBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointsGroupNameAll"); err == nil && ptr != 0 {
		VNHumanBodyPoseObservationJointsGroupNames.All = *(*VNHumanBodyPoseObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointsGroupNameFace"); err == nil && ptr != 0 {
		VNHumanBodyPoseObservationJointsGroupNames.Face = *(*VNHumanBodyPoseObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointsGroupNameLeftArm"); err == nil && ptr != 0 {
		VNHumanBodyPoseObservationJointsGroupNames.LeftArm = *(*VNHumanBodyPoseObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointsGroupNameLeftLeg"); err == nil && ptr != 0 {
		VNHumanBodyPoseObservationJointsGroupNames.LeftLeg = *(*VNHumanBodyPoseObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointsGroupNameRightArm"); err == nil && ptr != 0 {
		VNHumanBodyPoseObservationJointsGroupNames.RightArm = *(*VNHumanBodyPoseObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointsGroupNameRightLeg"); err == nil && ptr != 0 {
		VNHumanBodyPoseObservationJointsGroupNames.RightLeg = *(*VNHumanBodyPoseObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointsGroupNameTorso"); err == nil && ptr != 0 {
		VNHumanBodyPoseObservationJointsGroupNames.Torso = *(*VNHumanBodyPoseObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameIndexDIP"); err == nil && ptr != 0 {
		VNHumanHandPoseObservationJointNames.IndexDIP = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameIndexMCP"); err == nil && ptr != 0 {
		VNHumanHandPoseObservationJointNames.IndexMCP = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameIndexPIP"); err == nil && ptr != 0 {
		VNHumanHandPoseObservationJointNames.IndexPIP = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameIndexTip"); err == nil && ptr != 0 {
		VNHumanHandPoseObservationJointNames.IndexTip = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameLittleDIP"); err == nil && ptr != 0 {
		VNHumanHandPoseObservationJointNames.LittleDIP = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameLittleMCP"); err == nil && ptr != 0 {
		VNHumanHandPoseObservationJointNames.LittleMCP = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameLittlePIP"); err == nil && ptr != 0 {
		VNHumanHandPoseObservationJointNames.LittlePIP = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameLittleTip"); err == nil && ptr != 0 {
		VNHumanHandPoseObservationJointNames.LittleTip = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameMiddleDIP"); err == nil && ptr != 0 {
		VNHumanHandPoseObservationJointNames.MiddleDIP = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameMiddleMCP"); err == nil && ptr != 0 {
		VNHumanHandPoseObservationJointNames.MiddleMCP = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameMiddlePIP"); err == nil && ptr != 0 {
		VNHumanHandPoseObservationJointNames.MiddlePIP = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameMiddleTip"); err == nil && ptr != 0 {
		VNHumanHandPoseObservationJointNames.MiddleTip = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameRingDIP"); err == nil && ptr != 0 {
		VNHumanHandPoseObservationJointNames.RingDIP = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameRingMCP"); err == nil && ptr != 0 {
		VNHumanHandPoseObservationJointNames.RingMCP = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameRingPIP"); err == nil && ptr != 0 {
		VNHumanHandPoseObservationJointNames.RingPIP = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameRingTip"); err == nil && ptr != 0 {
		VNHumanHandPoseObservationJointNames.RingTip = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameThumbCMC"); err == nil && ptr != 0 {
		VNHumanHandPoseObservationJointNames.ThumbCMC = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameThumbIP"); err == nil && ptr != 0 {
		VNHumanHandPoseObservationJointNames.ThumbIP = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameThumbMP"); err == nil && ptr != 0 {
		VNHumanHandPoseObservationJointNames.ThumbMP = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameThumbTip"); err == nil && ptr != 0 {
		VNHumanHandPoseObservationJointNames.ThumbTip = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameWrist"); err == nil && ptr != 0 {
		VNHumanHandPoseObservationJointNames.Wrist = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointsGroupNameAll"); err == nil && ptr != 0 {
		VNHumanHandPoseObservationJointsGroupNames.All = *(*VNHumanHandPoseObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointsGroupNameIndexFinger"); err == nil && ptr != 0 {
		VNHumanHandPoseObservationJointsGroupNames.IndexFinger = *(*VNHumanHandPoseObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointsGroupNameLittleFinger"); err == nil && ptr != 0 {
		VNHumanHandPoseObservationJointsGroupNames.LittleFinger = *(*VNHumanHandPoseObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointsGroupNameMiddleFinger"); err == nil && ptr != 0 {
		VNHumanHandPoseObservationJointsGroupNames.MiddleFinger = *(*VNHumanHandPoseObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointsGroupNameRingFinger"); err == nil && ptr != 0 {
		VNHumanHandPoseObservationJointsGroupNames.RingFinger = *(*VNHumanHandPoseObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointsGroupNameThumb"); err == nil && ptr != 0 {
		VNHumanHandPoseObservationJointsGroupNames.Thumb = *(*VNHumanHandPoseObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNImageOptionCIContext"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VNImageOptions.CIContext = VNImageOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNImageOptionCameraIntrinsics"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VNImageOptions.CameraIntrinsics = VNImageOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNImageOptionProperties"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VNImageOptions.Properties = VNImageOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNNormalizedIdentityRect"); err == nil && ptr != 0 {
		VNNormalizedIdentityRect = *(*corefoundation.CGRect)(unsafe.Pointer(ptr))
	}






	if ptr, err := purego.Dlsym(frameworkHandle, "VNRecognizedPoint3DGroupKeyAll"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VNRecognizedPoint3DGroupKeyAll = VNRecognizedPointGroupKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNRecognizedPointGroupKeyAll"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				VNRecognizedPointGroupKeyAll = VNRecognizedPointGroupKey(objc.GoString(cstr))
			}
		}
	}









	if ptr, err := purego.Dlsym(frameworkHandle, "VNVisionVersionNumber"); err == nil && ptr != 0 {
		VNVisionVersionNumber = *(*float64)(unsafe.Pointer(ptr))
	}

}

// VNAnimalBodyPoseObservationJointNames provides typed accessors for [VNAnimalBodyPoseObservationJointName] constants.
var VNAnimalBodyPoseObservationJointNames struct {
	// LeftBackElbow: A joint name that represents the back of the left elbow.
	LeftBackElbow VNAnimalBodyPoseObservationJointName
	// LeftBackKnee: A joint name that represents the back of the left knee.
	LeftBackKnee VNAnimalBodyPoseObservationJointName
	// LeftBackPaw: A joint name that represents the back of the left paw.
	LeftBackPaw VNAnimalBodyPoseObservationJointName
	// LeftEarBottom: A joint name that represents the bottom of the left ear.
	LeftEarBottom VNAnimalBodyPoseObservationJointName
	// LeftEarMiddle: A joint name that represents the middle of the left ear.
	LeftEarMiddle VNAnimalBodyPoseObservationJointName
	// LeftEarTop: A joint name that represents the top of the left ear.
	LeftEarTop VNAnimalBodyPoseObservationJointName
	// LeftEye: A joint name that represents the left eye.
	LeftEye VNAnimalBodyPoseObservationJointName
	// LeftFrontElbow: A joint name that represents the front of the left elbow.
	LeftFrontElbow VNAnimalBodyPoseObservationJointName
	// LeftFrontKnee: A joint name that represents the front of the left knee.
	LeftFrontKnee VNAnimalBodyPoseObservationJointName
	// LeftFrontPaw: A joint name that represents the front of the left paw.
	LeftFrontPaw VNAnimalBodyPoseObservationJointName
	// Neck: A joint name that represents the neck.
	Neck VNAnimalBodyPoseObservationJointName
	// Nose: A joint name that represents the nose.
	Nose VNAnimalBodyPoseObservationJointName
	// RightBackElbow: A joint name that represents the back of the right elbow.
	RightBackElbow VNAnimalBodyPoseObservationJointName
	// RightBackKnee: A joint name that represents the back of the right knee.
	RightBackKnee VNAnimalBodyPoseObservationJointName
	// RightBackPaw: A joint name that represents the back of the right paw.
	RightBackPaw VNAnimalBodyPoseObservationJointName
	// RightEarBottom: A joint name that represents the bottom of the right ear.
	RightEarBottom VNAnimalBodyPoseObservationJointName
	// RightEarMiddle: A joint name that represents the middle of the right ear.
	RightEarMiddle VNAnimalBodyPoseObservationJointName
	// RightEarTop: A joint name that represents the top of the right ear.
	RightEarTop VNAnimalBodyPoseObservationJointName
	// RightEye: A joint name that represents the right eye.
	RightEye VNAnimalBodyPoseObservationJointName
	// RightFrontElbow: A joint name that represents the front of the right elbow.
	RightFrontElbow VNAnimalBodyPoseObservationJointName
	// RightFrontKnee: A joint name that represents the front of the right knee.
	RightFrontKnee VNAnimalBodyPoseObservationJointName
	// RightFrontPaw: A joint name that represents the front of the right paw.
	RightFrontPaw VNAnimalBodyPoseObservationJointName
	// TailBottom: A joint name that represents the bottom of the tail.
	TailBottom VNAnimalBodyPoseObservationJointName
	// TailMiddle: A joint name that represents the middle of the tail.
	TailMiddle VNAnimalBodyPoseObservationJointName
	// TailTop: A joint name that represents the top of the tail.
	TailTop VNAnimalBodyPoseObservationJointName
}

// VNAnimalBodyPoseObservationJointsGroupNames provides typed accessors for [VNAnimalBodyPoseObservationJointsGroupName] constants.
var VNAnimalBodyPoseObservationJointsGroupNames struct {
	// All: A group name that represents all joints.
	All VNAnimalBodyPoseObservationJointsGroupName
	// Forelegs: A group name that represents the forelegs.
	Forelegs VNAnimalBodyPoseObservationJointsGroupName
	// Head: A group name that represents the head.
	Head VNAnimalBodyPoseObservationJointsGroupName
	// Hindlegs: A group name that represents the hindlegs.
	Hindlegs VNAnimalBodyPoseObservationJointsGroupName
	// Tail: A group name that represents the tail.
	Tail VNAnimalBodyPoseObservationJointsGroupName
	// Trunk: A group name that represents the trunk.
	Trunk VNAnimalBodyPoseObservationJointsGroupName
}

// VNAnimalIdentifiers provides typed accessors for [VNAnimalIdentifier] constants.
var VNAnimalIdentifiers struct {
	// Cat: An animal identifier for cats.
	Cat VNAnimalIdentifier
	// Dog: An animal identifier for dogs.
	Dog VNAnimalIdentifier
}

// VNBarcodeSymbologys provides typed accessors for [VNBarcodeSymbology] constants.
var VNBarcodeSymbologys struct {
	// Aztec: A constant that indicates Aztec symbology.
	Aztec VNBarcodeSymbology
	// Codabar: A constant that indicates Codabar symbology.
	Codabar VNBarcodeSymbology
	// Code128: A constant that indicates Code 128 symbology.
	Code128 VNBarcodeSymbology
	// Code39: A constant that indicates Code 39 symbology.
	Code39 VNBarcodeSymbology
	// Code39Checksum: A constant that indicates Code 39 symbology with a checksum.
	Code39Checksum VNBarcodeSymbology
	// Code39FullASCII: A constant that indicates Code 39 Full ASCII symbology.
	Code39FullASCII VNBarcodeSymbology
	// Code39FullASCIIChecksum: A constant that indicates Code 39 Full ASCII symbology with a checksum.
	Code39FullASCIIChecksum VNBarcodeSymbology
	// Code93: A constant that indicates Code 93 symbology.
	Code93 VNBarcodeSymbology
	// Code93i: A constant that indicates Code 93i symbology.
	Code93i VNBarcodeSymbology
	// DataMatrix: A constant that indicates Data Matrix symbology.
	DataMatrix VNBarcodeSymbology
	// EAN13: A constant that indicates EAN-13 symbology.
	EAN13 VNBarcodeSymbology
	// EAN8: A constant that indicates EAN-8 symbology.
	EAN8 VNBarcodeSymbology
	// GS1DataBar: A constant that indicates GS1 DataBar symbology.
	GS1DataBar VNBarcodeSymbology
	// GS1DataBarExpanded: A constant that indicates GS1 DataBar Expanded symbology.
	GS1DataBarExpanded VNBarcodeSymbology
	// GS1DataBarLimited: A constant that indicates GS1 DataBar Limited symbology.
	GS1DataBarLimited VNBarcodeSymbology
	// I2of5: A constant that indicates Interleaved 2 of 5 (ITF) symbology.
	I2of5 VNBarcodeSymbology
	// I2of5Checksum: A constant that indicates Interleaved 2 of 5 (ITF) symbology with a checksum.
	I2of5Checksum VNBarcodeSymbology
	// ITF14: A constant that indicates ITF-14 symbology.
	ITF14 VNBarcodeSymbology
	// MSIPlessey: A constant that indicates Modified Plessey symbology.
	MSIPlessey VNBarcodeSymbology
	// MicroPDF417: A constant that indicates MicroPDF417 symbology.
	MicroPDF417 VNBarcodeSymbology
	// MicroQR: A constant that indicates MicroQR symbology.
	MicroQR VNBarcodeSymbology
	// PDF417: A constant that indicates PDF417 symbology.
	PDF417 VNBarcodeSymbology
	// QR: A constant that indicates Quick Response (QR) symbology.
	QR VNBarcodeSymbology
	// UPCE: A constant that indicates UPC-E symbology.
	UPCE VNBarcodeSymbology
}

// VNComputeStages provides typed accessors for [VNComputeStage] constants.
var VNComputeStages struct {
	// Main: A stage that represents where the system performs the main functionality.
	Main VNComputeStage
	// PostProcessing: A stage that represents where the system performs additional analysis from the main compute stage.
	PostProcessing VNComputeStage
}

// VNHumanBodyPose3DObservationJointNames provides typed accessors for [VNHumanBodyPose3DObservationJointName] constants.
var VNHumanBodyPose3DObservationJointNames struct {
	// CenterHead: A joint name that represents the center of the head.
	CenterHead VNHumanBodyPose3DObservationJointName
	// CenterShoulder: A joint name that represents the point between the shoulders.
	CenterShoulder VNHumanBodyPose3DObservationJointName
	// LeftAnkle: A joint name that represents the left ankle.
	LeftAnkle VNHumanBodyPose3DObservationJointName
	// LeftElbow: A joint name that represents the left elbow.
	LeftElbow VNHumanBodyPose3DObservationJointName
	// LeftHip: A joint name that represents the left hip.
	LeftHip VNHumanBodyPose3DObservationJointName
	// LeftKnee: A joint name that represents the left knee.
	LeftKnee VNHumanBodyPose3DObservationJointName
	// LeftShoulder: A joint name that represents the left shoulder.
	LeftShoulder VNHumanBodyPose3DObservationJointName
	// LeftWrist: A joint name that represents the left wrist.
	LeftWrist VNHumanBodyPose3DObservationJointName
	// RightAnkle: A joint name that represents the right ankle.
	RightAnkle VNHumanBodyPose3DObservationJointName
	// RightElbow: A joint name that represents the right elbow.
	RightElbow VNHumanBodyPose3DObservationJointName
	// RightHip: A joint name that represents the right hip.
	RightHip VNHumanBodyPose3DObservationJointName
	// RightKnee: A joint name that represents the right knee.
	RightKnee VNHumanBodyPose3DObservationJointName
	// RightShoulder: A joint name that represents the right shoulder.
	RightShoulder VNHumanBodyPose3DObservationJointName
	// RightWrist: A joint name that represents the right wrist.
	RightWrist VNHumanBodyPose3DObservationJointName
	// Root: A joint name that represents the point between the left hip and right hip.
	Root VNHumanBodyPose3DObservationJointName
	// Spine: A joint name that represents the spine.
	Spine VNHumanBodyPose3DObservationJointName
	// TopHead: A joint name that represents the top of the head.
	TopHead VNHumanBodyPose3DObservationJointName
}

// VNHumanBodyPose3DObservationJointsGroupNames provides typed accessors for [VNHumanBodyPose3DObservationJointsGroupName] constants.
var VNHumanBodyPose3DObservationJointsGroupNames struct {
	// All: A group name that represents all joints.
	All VNHumanBodyPose3DObservationJointsGroupName
	// Head: A group name that represents the head joints.
	Head VNHumanBodyPose3DObservationJointsGroupName
	// LeftArm: A group name that represents the left arm joints.
	LeftArm VNHumanBodyPose3DObservationJointsGroupName
	// LeftLeg: A group name that represents the left leg joints.
	LeftLeg VNHumanBodyPose3DObservationJointsGroupName
	// RightArm: A group name that represents the right arm joints.
	RightArm VNHumanBodyPose3DObservationJointsGroupName
	// RightLeg: A group name that represents the right leg joints.
	RightLeg VNHumanBodyPose3DObservationJointsGroupName
	// Torso: A group name that represents the torso joints.
	Torso VNHumanBodyPose3DObservationJointsGroupName
}

// VNHumanBodyPoseObservationJointNames provides typed accessors for [VNHumanBodyPoseObservationJointName] constants.
var VNHumanBodyPoseObservationJointNames struct {
	// LeftAnkle: The left ankle.
	LeftAnkle VNHumanBodyPoseObservationJointName
	// LeftEar: The left ear.
	LeftEar VNHumanBodyPoseObservationJointName
	// LeftElbow: The left elbow.
	LeftElbow VNHumanBodyPoseObservationJointName
	// LeftEye: The left eye.
	LeftEye VNHumanBodyPoseObservationJointName
	// LeftHip: The left hip.
	LeftHip VNHumanBodyPoseObservationJointName
	// LeftKnee: The left knee.
	LeftKnee VNHumanBodyPoseObservationJointName
	// LeftShoulder: The left shoulder.
	LeftShoulder VNHumanBodyPoseObservationJointName
	// LeftWrist: The left wrist.
	LeftWrist VNHumanBodyPoseObservationJointName
	// Neck: The neck.
	Neck VNHumanBodyPoseObservationJointName
	// Nose: The nose.
	Nose VNHumanBodyPoseObservationJointName
	// RightAnkle: The right ankle.
	RightAnkle VNHumanBodyPoseObservationJointName
	// RightEar: The right ear.
	RightEar VNHumanBodyPoseObservationJointName
	// RightElbow: The right elbow.
	RightElbow VNHumanBodyPoseObservationJointName
	// RightEye: The right eye.
	RightEye VNHumanBodyPoseObservationJointName
	// RightHip: The right hip.
	RightHip VNHumanBodyPoseObservationJointName
	// RightKnee: The right knee.
	RightKnee VNHumanBodyPoseObservationJointName
	// RightShoulder: The right shoulder.
	RightShoulder VNHumanBodyPoseObservationJointName
	// RightWrist: The right wrist.
	RightWrist VNHumanBodyPoseObservationJointName
	// Root: The root (waist).
	Root VNHumanBodyPoseObservationJointName
}

// VNHumanBodyPoseObservationJointsGroupNames provides typed accessors for [VNHumanBodyPoseObservationJointsGroupName] constants.
var VNHumanBodyPoseObservationJointsGroupNames struct {
	// All: All body point groups.
	All VNHumanBodyPoseObservationJointsGroupName
	// Face: The face.
	Face VNHumanBodyPoseObservationJointsGroupName
	// LeftArm: The left arm.
	LeftArm VNHumanBodyPoseObservationJointsGroupName
	// LeftLeg: The left leg.
	LeftLeg VNHumanBodyPoseObservationJointsGroupName
	// RightArm: The right arm.
	RightArm VNHumanBodyPoseObservationJointsGroupName
	// RightLeg: The right leg.
	RightLeg VNHumanBodyPoseObservationJointsGroupName
	// Torso: The torso.
	Torso VNHumanBodyPoseObservationJointsGroupName
}

// VNHumanHandPoseObservationJointNames provides typed accessors for [VNHumanHandPoseObservationJointName] constants.
var VNHumanHandPoseObservationJointNames struct {
	// IndexDIP: The index finger’s distal interphalangeal (DIP) joint.
	IndexDIP VNHumanHandPoseObservationJointName
	// IndexMCP: The index finger’s metacarpophalangeal (MCP) joint.
	IndexMCP VNHumanHandPoseObservationJointName
	// IndexPIP: The index finger’s proximal interphalangeal (PIP) joint.
	IndexPIP VNHumanHandPoseObservationJointName
	// IndexTip: The tip of the index finger.
	IndexTip VNHumanHandPoseObservationJointName
	// LittleDIP: The little finger’s distal interphalangeal (DIP) joint.
	LittleDIP VNHumanHandPoseObservationJointName
	// LittleMCP: The little finger’s metacarpophalangeal (MCP) joint.
	LittleMCP VNHumanHandPoseObservationJointName
	// LittlePIP: The little finger’s proximal interphalangeal (PIP) joint.
	LittlePIP VNHumanHandPoseObservationJointName
	// LittleTip: The tip of the little finger.
	LittleTip VNHumanHandPoseObservationJointName
	// MiddleDIP: The middle finger’s distal interphalangeal (DIP) joint.
	MiddleDIP VNHumanHandPoseObservationJointName
	// MiddleMCP: The middle finger’s metacarpophalangeal (MCP) joint.
	MiddleMCP VNHumanHandPoseObservationJointName
	// MiddlePIP: The middle finger’s proximal interphalangeal (PIP) joint.
	MiddlePIP VNHumanHandPoseObservationJointName
	// MiddleTip: The tip of the middle finger.
	MiddleTip VNHumanHandPoseObservationJointName
	// RingDIP: The ring finger’s distal interphalangeal (DIP) joint.
	RingDIP VNHumanHandPoseObservationJointName
	// RingMCP: The ring finger’s metacarpophalangeal (MCP) joint.
	RingMCP VNHumanHandPoseObservationJointName
	// RingPIP: The ring finger’s proximal interphalangeal (PIP) joint.
	RingPIP VNHumanHandPoseObservationJointName
	// RingTip: The tip of the ring finger.
	RingTip VNHumanHandPoseObservationJointName
	// ThumbCMC: The thumb’s carpometacarpal (CMC) joint.
	ThumbCMC VNHumanHandPoseObservationJointName
	// ThumbIP: The thumb’s interphalangeal (IP) joint.
	ThumbIP VNHumanHandPoseObservationJointName
	// ThumbMP: The thumb’s metacarpophalangeal (MP) joint.
	ThumbMP VNHumanHandPoseObservationJointName
	// ThumbTip: The tip of the thumb.
	ThumbTip VNHumanHandPoseObservationJointName
	// Wrist: The wrist.
	Wrist VNHumanHandPoseObservationJointName
}

// VNHumanHandPoseObservationJointsGroupNames provides typed accessors for [VNHumanHandPoseObservationJointsGroupName] constants.
var VNHumanHandPoseObservationJointsGroupNames struct {
	// All: All hand group names.
	All VNHumanHandPoseObservationJointsGroupName
	// IndexFinger: The index finger.
	IndexFinger VNHumanHandPoseObservationJointsGroupName
	// LittleFinger: The little finger.
	LittleFinger VNHumanHandPoseObservationJointsGroupName
	// MiddleFinger: The middle finger.
	MiddleFinger VNHumanHandPoseObservationJointsGroupName
	// RingFinger: The ring finger.
	RingFinger VNHumanHandPoseObservationJointsGroupName
	// Thumb: The thumb.
	Thumb VNHumanHandPoseObservationJointsGroupName
}

// VNImageOptions provides typed accessors for [VNImageOption] constants.
var VNImageOptions struct {
	// CIContext: An option key to specify the context to use in the handler’s Core Image operations.
	CIContext VNImageOption
	// CameraIntrinsics: An option to specify the camera intrinstics.
	CameraIntrinsics VNImageOption
	// Properties: The dictionary from the image source that contains the metadata for algorithms like horizon detection.
	Properties VNImageOption
}

