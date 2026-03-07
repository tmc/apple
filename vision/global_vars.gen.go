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

var vNAnimalBodyPoseObservationJointNameLeftBackElbow VNAnimalBodyPoseObservationJointName

var vNAnimalBodyPoseObservationJointNameLeftBackKnee VNAnimalBodyPoseObservationJointName

var vNAnimalBodyPoseObservationJointNameLeftBackPaw VNAnimalBodyPoseObservationJointName

var vNAnimalBodyPoseObservationJointNameLeftEarBottom VNAnimalBodyPoseObservationJointName

var vNAnimalBodyPoseObservationJointNameLeftEarMiddle VNAnimalBodyPoseObservationJointName

var vNAnimalBodyPoseObservationJointNameLeftEarTop VNAnimalBodyPoseObservationJointName

var vNAnimalBodyPoseObservationJointNameLeftEye VNAnimalBodyPoseObservationJointName

var vNAnimalBodyPoseObservationJointNameLeftFrontElbow VNAnimalBodyPoseObservationJointName

var vNAnimalBodyPoseObservationJointNameLeftFrontKnee VNAnimalBodyPoseObservationJointName

var vNAnimalBodyPoseObservationJointNameLeftFrontPaw VNAnimalBodyPoseObservationJointName

var vNAnimalBodyPoseObservationJointNameNeck VNAnimalBodyPoseObservationJointName

var vNAnimalBodyPoseObservationJointNameNose VNAnimalBodyPoseObservationJointName

var vNAnimalBodyPoseObservationJointNameRightBackElbow VNAnimalBodyPoseObservationJointName

var vNAnimalBodyPoseObservationJointNameRightBackKnee VNAnimalBodyPoseObservationJointName

var vNAnimalBodyPoseObservationJointNameRightBackPaw VNAnimalBodyPoseObservationJointName

var vNAnimalBodyPoseObservationJointNameRightEarBottom VNAnimalBodyPoseObservationJointName

var vNAnimalBodyPoseObservationJointNameRightEarMiddle VNAnimalBodyPoseObservationJointName

var vNAnimalBodyPoseObservationJointNameRightEarTop VNAnimalBodyPoseObservationJointName

var vNAnimalBodyPoseObservationJointNameRightEye VNAnimalBodyPoseObservationJointName

var vNAnimalBodyPoseObservationJointNameRightFrontElbow VNAnimalBodyPoseObservationJointName

var vNAnimalBodyPoseObservationJointNameRightFrontKnee VNAnimalBodyPoseObservationJointName

var vNAnimalBodyPoseObservationJointNameRightFrontPaw VNAnimalBodyPoseObservationJointName

var vNAnimalBodyPoseObservationJointNameTailBottom VNAnimalBodyPoseObservationJointName

var vNAnimalBodyPoseObservationJointNameTailMiddle VNAnimalBodyPoseObservationJointName

var vNAnimalBodyPoseObservationJointNameTailTop VNAnimalBodyPoseObservationJointName

var vNAnimalBodyPoseObservationJointsGroupNameAll VNAnimalBodyPoseObservationJointsGroupName

var vNAnimalBodyPoseObservationJointsGroupNameForelegs VNAnimalBodyPoseObservationJointsGroupName

var vNAnimalBodyPoseObservationJointsGroupNameHead VNAnimalBodyPoseObservationJointsGroupName

var vNAnimalBodyPoseObservationJointsGroupNameHindlegs VNAnimalBodyPoseObservationJointsGroupName

var vNAnimalBodyPoseObservationJointsGroupNameTail VNAnimalBodyPoseObservationJointsGroupName

var vNAnimalBodyPoseObservationJointsGroupNameTrunk VNAnimalBodyPoseObservationJointsGroupName

var vNAnimalIdentifierCat VNAnimalIdentifier

var vNAnimalIdentifierDog VNAnimalIdentifier

var vNBarcodeSymbologyAztec VNBarcodeSymbology

var vNBarcodeSymbologyCodabar VNBarcodeSymbology

var vNBarcodeSymbologyCode128 VNBarcodeSymbology

var vNBarcodeSymbologyCode39 VNBarcodeSymbology

var vNBarcodeSymbologyCode39Checksum VNBarcodeSymbology

var vNBarcodeSymbologyCode39FullASCII VNBarcodeSymbology

var vNBarcodeSymbologyCode39FullASCIIChecksum VNBarcodeSymbology

var vNBarcodeSymbologyCode93 VNBarcodeSymbology

var vNBarcodeSymbologyCode93i VNBarcodeSymbology

var vNBarcodeSymbologyDataMatrix VNBarcodeSymbology

var vNBarcodeSymbologyEAN13 VNBarcodeSymbology

var vNBarcodeSymbologyEAN8 VNBarcodeSymbology

var vNBarcodeSymbologyGS1DataBar VNBarcodeSymbology

var vNBarcodeSymbologyGS1DataBarExpanded VNBarcodeSymbology

var vNBarcodeSymbologyGS1DataBarLimited VNBarcodeSymbology

var vNBarcodeSymbologyI2of5 VNBarcodeSymbology

var vNBarcodeSymbologyI2of5Checksum VNBarcodeSymbology

var vNBarcodeSymbologyITF14 VNBarcodeSymbology

var vNBarcodeSymbologyMSIPlessey VNBarcodeSymbology

var vNBarcodeSymbologyMicroPDF417 VNBarcodeSymbology

var vNBarcodeSymbologyMicroQR VNBarcodeSymbology

var vNBarcodeSymbologyPDF417 VNBarcodeSymbology

var vNBarcodeSymbologyQR VNBarcodeSymbology

var vNBarcodeSymbologyUPCE VNBarcodeSymbology




var vNComputeStageMain VNComputeStage

var vNComputeStagePostProcessing VNComputeStage

























var VNErrorDomain string













var vNHumanBodyPose3DObservationJointNameCenterHead VNHumanBodyPose3DObservationJointName

var vNHumanBodyPose3DObservationJointNameCenterShoulder VNHumanBodyPose3DObservationJointName

var vNHumanBodyPose3DObservationJointNameLeftAnkle VNHumanBodyPose3DObservationJointName

var vNHumanBodyPose3DObservationJointNameLeftElbow VNHumanBodyPose3DObservationJointName

var vNHumanBodyPose3DObservationJointNameLeftHip VNHumanBodyPose3DObservationJointName

var vNHumanBodyPose3DObservationJointNameLeftKnee VNHumanBodyPose3DObservationJointName

var vNHumanBodyPose3DObservationJointNameLeftShoulder VNHumanBodyPose3DObservationJointName

var vNHumanBodyPose3DObservationJointNameLeftWrist VNHumanBodyPose3DObservationJointName

var vNHumanBodyPose3DObservationJointNameRightAnkle VNHumanBodyPose3DObservationJointName

var vNHumanBodyPose3DObservationJointNameRightElbow VNHumanBodyPose3DObservationJointName

var vNHumanBodyPose3DObservationJointNameRightHip VNHumanBodyPose3DObservationJointName

var vNHumanBodyPose3DObservationJointNameRightKnee VNHumanBodyPose3DObservationJointName

var vNHumanBodyPose3DObservationJointNameRightShoulder VNHumanBodyPose3DObservationJointName

var vNHumanBodyPose3DObservationJointNameRightWrist VNHumanBodyPose3DObservationJointName

var vNHumanBodyPose3DObservationJointNameRoot VNHumanBodyPose3DObservationJointName

var vNHumanBodyPose3DObservationJointNameSpine VNHumanBodyPose3DObservationJointName

var vNHumanBodyPose3DObservationJointNameTopHead VNHumanBodyPose3DObservationJointName

var vNHumanBodyPose3DObservationJointsGroupNameAll VNHumanBodyPose3DObservationJointsGroupName

var vNHumanBodyPose3DObservationJointsGroupNameHead VNHumanBodyPose3DObservationJointsGroupName

var vNHumanBodyPose3DObservationJointsGroupNameLeftArm VNHumanBodyPose3DObservationJointsGroupName

var vNHumanBodyPose3DObservationJointsGroupNameLeftLeg VNHumanBodyPose3DObservationJointsGroupName

var vNHumanBodyPose3DObservationJointsGroupNameRightArm VNHumanBodyPose3DObservationJointsGroupName

var vNHumanBodyPose3DObservationJointsGroupNameRightLeg VNHumanBodyPose3DObservationJointsGroupName

var vNHumanBodyPose3DObservationJointsGroupNameTorso VNHumanBodyPose3DObservationJointsGroupName

var vNHumanBodyPoseObservationJointNameLeftAnkle VNHumanBodyPoseObservationJointName

var vNHumanBodyPoseObservationJointNameLeftEar VNHumanBodyPoseObservationJointName

var vNHumanBodyPoseObservationJointNameLeftElbow VNHumanBodyPoseObservationJointName

var vNHumanBodyPoseObservationJointNameLeftEye VNHumanBodyPoseObservationJointName

var vNHumanBodyPoseObservationJointNameLeftHip VNHumanBodyPoseObservationJointName

var vNHumanBodyPoseObservationJointNameLeftKnee VNHumanBodyPoseObservationJointName

var vNHumanBodyPoseObservationJointNameLeftShoulder VNHumanBodyPoseObservationJointName

var vNHumanBodyPoseObservationJointNameLeftWrist VNHumanBodyPoseObservationJointName

var vNHumanBodyPoseObservationJointNameNeck VNHumanBodyPoseObservationJointName

var vNHumanBodyPoseObservationJointNameNose VNHumanBodyPoseObservationJointName

var vNHumanBodyPoseObservationJointNameRightAnkle VNHumanBodyPoseObservationJointName

var vNHumanBodyPoseObservationJointNameRightEar VNHumanBodyPoseObservationJointName

var vNHumanBodyPoseObservationJointNameRightElbow VNHumanBodyPoseObservationJointName

var vNHumanBodyPoseObservationJointNameRightEye VNHumanBodyPoseObservationJointName

var vNHumanBodyPoseObservationJointNameRightHip VNHumanBodyPoseObservationJointName

var vNHumanBodyPoseObservationJointNameRightKnee VNHumanBodyPoseObservationJointName

var vNHumanBodyPoseObservationJointNameRightShoulder VNHumanBodyPoseObservationJointName

var vNHumanBodyPoseObservationJointNameRightWrist VNHumanBodyPoseObservationJointName

var vNHumanBodyPoseObservationJointNameRoot VNHumanBodyPoseObservationJointName

var vNHumanBodyPoseObservationJointsGroupNameAll VNHumanBodyPoseObservationJointsGroupName

var vNHumanBodyPoseObservationJointsGroupNameFace VNHumanBodyPoseObservationJointsGroupName

var vNHumanBodyPoseObservationJointsGroupNameLeftArm VNHumanBodyPoseObservationJointsGroupName

var vNHumanBodyPoseObservationJointsGroupNameLeftLeg VNHumanBodyPoseObservationJointsGroupName

var vNHumanBodyPoseObservationJointsGroupNameRightArm VNHumanBodyPoseObservationJointsGroupName

var vNHumanBodyPoseObservationJointsGroupNameRightLeg VNHumanBodyPoseObservationJointsGroupName

var vNHumanBodyPoseObservationJointsGroupNameTorso VNHumanBodyPoseObservationJointsGroupName

var vNHumanHandPoseObservationJointNameIndexDIP VNHumanHandPoseObservationJointName

var vNHumanHandPoseObservationJointNameIndexMCP VNHumanHandPoseObservationJointName

var vNHumanHandPoseObservationJointNameIndexPIP VNHumanHandPoseObservationJointName

var vNHumanHandPoseObservationJointNameIndexTip VNHumanHandPoseObservationJointName

var vNHumanHandPoseObservationJointNameLittleDIP VNHumanHandPoseObservationJointName

var vNHumanHandPoseObservationJointNameLittleMCP VNHumanHandPoseObservationJointName

var vNHumanHandPoseObservationJointNameLittlePIP VNHumanHandPoseObservationJointName

var vNHumanHandPoseObservationJointNameLittleTip VNHumanHandPoseObservationJointName

var vNHumanHandPoseObservationJointNameMiddleDIP VNHumanHandPoseObservationJointName

var vNHumanHandPoseObservationJointNameMiddleMCP VNHumanHandPoseObservationJointName

var vNHumanHandPoseObservationJointNameMiddlePIP VNHumanHandPoseObservationJointName

var vNHumanHandPoseObservationJointNameMiddleTip VNHumanHandPoseObservationJointName

var vNHumanHandPoseObservationJointNameRingDIP VNHumanHandPoseObservationJointName

var vNHumanHandPoseObservationJointNameRingMCP VNHumanHandPoseObservationJointName

var vNHumanHandPoseObservationJointNameRingPIP VNHumanHandPoseObservationJointName

var vNHumanHandPoseObservationJointNameRingTip VNHumanHandPoseObservationJointName

var vNHumanHandPoseObservationJointNameThumbCMC VNHumanHandPoseObservationJointName

var vNHumanHandPoseObservationJointNameThumbIP VNHumanHandPoseObservationJointName

var vNHumanHandPoseObservationJointNameThumbMP VNHumanHandPoseObservationJointName

var vNHumanHandPoseObservationJointNameThumbTip VNHumanHandPoseObservationJointName

var vNHumanHandPoseObservationJointNameWrist VNHumanHandPoseObservationJointName

var vNHumanHandPoseObservationJointsGroupNameAll VNHumanHandPoseObservationJointsGroupName

var vNHumanHandPoseObservationJointsGroupNameIndexFinger VNHumanHandPoseObservationJointsGroupName

var vNHumanHandPoseObservationJointsGroupNameLittleFinger VNHumanHandPoseObservationJointsGroupName

var vNHumanHandPoseObservationJointsGroupNameMiddleFinger VNHumanHandPoseObservationJointsGroupName

var vNHumanHandPoseObservationJointsGroupNameRingFinger VNHumanHandPoseObservationJointsGroupName

var vNHumanHandPoseObservationJointsGroupNameThumb VNHumanHandPoseObservationJointsGroupName

var vNImageOptionCIContext VNImageOption

var vNImageOptionCameraIntrinsics VNImageOption

var vNImageOptionProperties VNImageOption

var VNNormalizedIdentityRect corefoundation.CGRect






var VNRecognizedPoint3DGroupKeyAll VNRecognizedPointGroupKey

var VNRecognizedPointGroupKeyAll VNRecognizedPointGroupKey









var VNVisionVersionNumber float64

func init() {
	if frameworkHandle == 0 {
		return
	}


	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameLeftBackElbow"); err == nil && ptr != 0 {
		vNAnimalBodyPoseObservationJointNameLeftBackElbow = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameLeftBackKnee"); err == nil && ptr != 0 {
		vNAnimalBodyPoseObservationJointNameLeftBackKnee = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameLeftBackPaw"); err == nil && ptr != 0 {
		vNAnimalBodyPoseObservationJointNameLeftBackPaw = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameLeftEarBottom"); err == nil && ptr != 0 {
		vNAnimalBodyPoseObservationJointNameLeftEarBottom = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameLeftEarMiddle"); err == nil && ptr != 0 {
		vNAnimalBodyPoseObservationJointNameLeftEarMiddle = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameLeftEarTop"); err == nil && ptr != 0 {
		vNAnimalBodyPoseObservationJointNameLeftEarTop = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameLeftEye"); err == nil && ptr != 0 {
		vNAnimalBodyPoseObservationJointNameLeftEye = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameLeftFrontElbow"); err == nil && ptr != 0 {
		vNAnimalBodyPoseObservationJointNameLeftFrontElbow = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameLeftFrontKnee"); err == nil && ptr != 0 {
		vNAnimalBodyPoseObservationJointNameLeftFrontKnee = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameLeftFrontPaw"); err == nil && ptr != 0 {
		vNAnimalBodyPoseObservationJointNameLeftFrontPaw = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameNeck"); err == nil && ptr != 0 {
		vNAnimalBodyPoseObservationJointNameNeck = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameNose"); err == nil && ptr != 0 {
		vNAnimalBodyPoseObservationJointNameNose = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameRightBackElbow"); err == nil && ptr != 0 {
		vNAnimalBodyPoseObservationJointNameRightBackElbow = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameRightBackKnee"); err == nil && ptr != 0 {
		vNAnimalBodyPoseObservationJointNameRightBackKnee = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameRightBackPaw"); err == nil && ptr != 0 {
		vNAnimalBodyPoseObservationJointNameRightBackPaw = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameRightEarBottom"); err == nil && ptr != 0 {
		vNAnimalBodyPoseObservationJointNameRightEarBottom = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameRightEarMiddle"); err == nil && ptr != 0 {
		vNAnimalBodyPoseObservationJointNameRightEarMiddle = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameRightEarTop"); err == nil && ptr != 0 {
		vNAnimalBodyPoseObservationJointNameRightEarTop = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameRightEye"); err == nil && ptr != 0 {
		vNAnimalBodyPoseObservationJointNameRightEye = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameRightFrontElbow"); err == nil && ptr != 0 {
		vNAnimalBodyPoseObservationJointNameRightFrontElbow = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameRightFrontKnee"); err == nil && ptr != 0 {
		vNAnimalBodyPoseObservationJointNameRightFrontKnee = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameRightFrontPaw"); err == nil && ptr != 0 {
		vNAnimalBodyPoseObservationJointNameRightFrontPaw = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameTailBottom"); err == nil && ptr != 0 {
		vNAnimalBodyPoseObservationJointNameTailBottom = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameTailMiddle"); err == nil && ptr != 0 {
		vNAnimalBodyPoseObservationJointNameTailMiddle = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointNameTailTop"); err == nil && ptr != 0 {
		vNAnimalBodyPoseObservationJointNameTailTop = *(*VNAnimalBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointsGroupNameAll"); err == nil && ptr != 0 {
		vNAnimalBodyPoseObservationJointsGroupNameAll = *(*VNAnimalBodyPoseObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointsGroupNameForelegs"); err == nil && ptr != 0 {
		vNAnimalBodyPoseObservationJointsGroupNameForelegs = *(*VNAnimalBodyPoseObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointsGroupNameHead"); err == nil && ptr != 0 {
		vNAnimalBodyPoseObservationJointsGroupNameHead = *(*VNAnimalBodyPoseObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointsGroupNameHindlegs"); err == nil && ptr != 0 {
		vNAnimalBodyPoseObservationJointsGroupNameHindlegs = *(*VNAnimalBodyPoseObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointsGroupNameTail"); err == nil && ptr != 0 {
		vNAnimalBodyPoseObservationJointsGroupNameTail = *(*VNAnimalBodyPoseObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalBodyPoseObservationJointsGroupNameTrunk"); err == nil && ptr != 0 {
		vNAnimalBodyPoseObservationJointsGroupNameTrunk = *(*VNAnimalBodyPoseObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalIdentifierCat"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				vNAnimalIdentifierCat = VNAnimalIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNAnimalIdentifierDog"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				vNAnimalIdentifierDog = VNAnimalIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyAztec"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				vNBarcodeSymbologyAztec = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyCodabar"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				vNBarcodeSymbologyCodabar = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyCode128"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				vNBarcodeSymbologyCode128 = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyCode39"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				vNBarcodeSymbologyCode39 = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyCode39Checksum"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				vNBarcodeSymbologyCode39Checksum = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyCode39FullASCII"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				vNBarcodeSymbologyCode39FullASCII = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyCode39FullASCIIChecksum"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				vNBarcodeSymbologyCode39FullASCIIChecksum = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyCode93"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				vNBarcodeSymbologyCode93 = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyCode93i"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				vNBarcodeSymbologyCode93i = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyDataMatrix"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				vNBarcodeSymbologyDataMatrix = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyEAN13"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				vNBarcodeSymbologyEAN13 = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyEAN8"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				vNBarcodeSymbologyEAN8 = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyGS1DataBar"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				vNBarcodeSymbologyGS1DataBar = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyGS1DataBarExpanded"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				vNBarcodeSymbologyGS1DataBarExpanded = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyGS1DataBarLimited"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				vNBarcodeSymbologyGS1DataBarLimited = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyI2of5"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				vNBarcodeSymbologyI2of5 = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyI2of5Checksum"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				vNBarcodeSymbologyI2of5Checksum = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyITF14"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				vNBarcodeSymbologyITF14 = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyMSIPlessey"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				vNBarcodeSymbologyMSIPlessey = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyMicroPDF417"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				vNBarcodeSymbologyMicroPDF417 = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyMicroQR"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				vNBarcodeSymbologyMicroQR = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyPDF417"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				vNBarcodeSymbologyPDF417 = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyQR"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				vNBarcodeSymbologyQR = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNBarcodeSymbologyUPCE"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				vNBarcodeSymbologyUPCE = VNBarcodeSymbology(objc.GoString(cstr))
			}
		}
	}




	if ptr, err := purego.Dlsym(frameworkHandle, "VNComputeStageMain"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				vNComputeStageMain = VNComputeStage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNComputeStagePostProcessing"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				vNComputeStagePostProcessing = VNComputeStage(objc.GoString(cstr))
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
		vNHumanBodyPose3DObservationJointNameCenterHead = *(*VNHumanBodyPose3DObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointNameCenterShoulder"); err == nil && ptr != 0 {
		vNHumanBodyPose3DObservationJointNameCenterShoulder = *(*VNHumanBodyPose3DObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointNameLeftAnkle"); err == nil && ptr != 0 {
		vNHumanBodyPose3DObservationJointNameLeftAnkle = *(*VNHumanBodyPose3DObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointNameLeftElbow"); err == nil && ptr != 0 {
		vNHumanBodyPose3DObservationJointNameLeftElbow = *(*VNHumanBodyPose3DObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointNameLeftHip"); err == nil && ptr != 0 {
		vNHumanBodyPose3DObservationJointNameLeftHip = *(*VNHumanBodyPose3DObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointNameLeftKnee"); err == nil && ptr != 0 {
		vNHumanBodyPose3DObservationJointNameLeftKnee = *(*VNHumanBodyPose3DObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointNameLeftShoulder"); err == nil && ptr != 0 {
		vNHumanBodyPose3DObservationJointNameLeftShoulder = *(*VNHumanBodyPose3DObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointNameLeftWrist"); err == nil && ptr != 0 {
		vNHumanBodyPose3DObservationJointNameLeftWrist = *(*VNHumanBodyPose3DObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointNameRightAnkle"); err == nil && ptr != 0 {
		vNHumanBodyPose3DObservationJointNameRightAnkle = *(*VNHumanBodyPose3DObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointNameRightElbow"); err == nil && ptr != 0 {
		vNHumanBodyPose3DObservationJointNameRightElbow = *(*VNHumanBodyPose3DObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointNameRightHip"); err == nil && ptr != 0 {
		vNHumanBodyPose3DObservationJointNameRightHip = *(*VNHumanBodyPose3DObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointNameRightKnee"); err == nil && ptr != 0 {
		vNHumanBodyPose3DObservationJointNameRightKnee = *(*VNHumanBodyPose3DObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointNameRightShoulder"); err == nil && ptr != 0 {
		vNHumanBodyPose3DObservationJointNameRightShoulder = *(*VNHumanBodyPose3DObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointNameRightWrist"); err == nil && ptr != 0 {
		vNHumanBodyPose3DObservationJointNameRightWrist = *(*VNHumanBodyPose3DObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointNameRoot"); err == nil && ptr != 0 {
		vNHumanBodyPose3DObservationJointNameRoot = *(*VNHumanBodyPose3DObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointNameSpine"); err == nil && ptr != 0 {
		vNHumanBodyPose3DObservationJointNameSpine = *(*VNHumanBodyPose3DObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointNameTopHead"); err == nil && ptr != 0 {
		vNHumanBodyPose3DObservationJointNameTopHead = *(*VNHumanBodyPose3DObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointsGroupNameAll"); err == nil && ptr != 0 {
		vNHumanBodyPose3DObservationJointsGroupNameAll = *(*VNHumanBodyPose3DObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointsGroupNameHead"); err == nil && ptr != 0 {
		vNHumanBodyPose3DObservationJointsGroupNameHead = *(*VNHumanBodyPose3DObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointsGroupNameLeftArm"); err == nil && ptr != 0 {
		vNHumanBodyPose3DObservationJointsGroupNameLeftArm = *(*VNHumanBodyPose3DObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointsGroupNameLeftLeg"); err == nil && ptr != 0 {
		vNHumanBodyPose3DObservationJointsGroupNameLeftLeg = *(*VNHumanBodyPose3DObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointsGroupNameRightArm"); err == nil && ptr != 0 {
		vNHumanBodyPose3DObservationJointsGroupNameRightArm = *(*VNHumanBodyPose3DObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointsGroupNameRightLeg"); err == nil && ptr != 0 {
		vNHumanBodyPose3DObservationJointsGroupNameRightLeg = *(*VNHumanBodyPose3DObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPose3DObservationJointsGroupNameTorso"); err == nil && ptr != 0 {
		vNHumanBodyPose3DObservationJointsGroupNameTorso = *(*VNHumanBodyPose3DObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointNameLeftAnkle"); err == nil && ptr != 0 {
		vNHumanBodyPoseObservationJointNameLeftAnkle = *(*VNHumanBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointNameLeftEar"); err == nil && ptr != 0 {
		vNHumanBodyPoseObservationJointNameLeftEar = *(*VNHumanBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointNameLeftElbow"); err == nil && ptr != 0 {
		vNHumanBodyPoseObservationJointNameLeftElbow = *(*VNHumanBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointNameLeftEye"); err == nil && ptr != 0 {
		vNHumanBodyPoseObservationJointNameLeftEye = *(*VNHumanBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointNameLeftHip"); err == nil && ptr != 0 {
		vNHumanBodyPoseObservationJointNameLeftHip = *(*VNHumanBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointNameLeftKnee"); err == nil && ptr != 0 {
		vNHumanBodyPoseObservationJointNameLeftKnee = *(*VNHumanBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointNameLeftShoulder"); err == nil && ptr != 0 {
		vNHumanBodyPoseObservationJointNameLeftShoulder = *(*VNHumanBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointNameLeftWrist"); err == nil && ptr != 0 {
		vNHumanBodyPoseObservationJointNameLeftWrist = *(*VNHumanBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointNameNeck"); err == nil && ptr != 0 {
		vNHumanBodyPoseObservationJointNameNeck = *(*VNHumanBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointNameNose"); err == nil && ptr != 0 {
		vNHumanBodyPoseObservationJointNameNose = *(*VNHumanBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointNameRightAnkle"); err == nil && ptr != 0 {
		vNHumanBodyPoseObservationJointNameRightAnkle = *(*VNHumanBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointNameRightEar"); err == nil && ptr != 0 {
		vNHumanBodyPoseObservationJointNameRightEar = *(*VNHumanBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointNameRightElbow"); err == nil && ptr != 0 {
		vNHumanBodyPoseObservationJointNameRightElbow = *(*VNHumanBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointNameRightEye"); err == nil && ptr != 0 {
		vNHumanBodyPoseObservationJointNameRightEye = *(*VNHumanBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointNameRightHip"); err == nil && ptr != 0 {
		vNHumanBodyPoseObservationJointNameRightHip = *(*VNHumanBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointNameRightKnee"); err == nil && ptr != 0 {
		vNHumanBodyPoseObservationJointNameRightKnee = *(*VNHumanBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointNameRightShoulder"); err == nil && ptr != 0 {
		vNHumanBodyPoseObservationJointNameRightShoulder = *(*VNHumanBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointNameRightWrist"); err == nil && ptr != 0 {
		vNHumanBodyPoseObservationJointNameRightWrist = *(*VNHumanBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointNameRoot"); err == nil && ptr != 0 {
		vNHumanBodyPoseObservationJointNameRoot = *(*VNHumanBodyPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointsGroupNameAll"); err == nil && ptr != 0 {
		vNHumanBodyPoseObservationJointsGroupNameAll = *(*VNHumanBodyPoseObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointsGroupNameFace"); err == nil && ptr != 0 {
		vNHumanBodyPoseObservationJointsGroupNameFace = *(*VNHumanBodyPoseObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointsGroupNameLeftArm"); err == nil && ptr != 0 {
		vNHumanBodyPoseObservationJointsGroupNameLeftArm = *(*VNHumanBodyPoseObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointsGroupNameLeftLeg"); err == nil && ptr != 0 {
		vNHumanBodyPoseObservationJointsGroupNameLeftLeg = *(*VNHumanBodyPoseObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointsGroupNameRightArm"); err == nil && ptr != 0 {
		vNHumanBodyPoseObservationJointsGroupNameRightArm = *(*VNHumanBodyPoseObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointsGroupNameRightLeg"); err == nil && ptr != 0 {
		vNHumanBodyPoseObservationJointsGroupNameRightLeg = *(*VNHumanBodyPoseObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanBodyPoseObservationJointsGroupNameTorso"); err == nil && ptr != 0 {
		vNHumanBodyPoseObservationJointsGroupNameTorso = *(*VNHumanBodyPoseObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameIndexDIP"); err == nil && ptr != 0 {
		vNHumanHandPoseObservationJointNameIndexDIP = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameIndexMCP"); err == nil && ptr != 0 {
		vNHumanHandPoseObservationJointNameIndexMCP = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameIndexPIP"); err == nil && ptr != 0 {
		vNHumanHandPoseObservationJointNameIndexPIP = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameIndexTip"); err == nil && ptr != 0 {
		vNHumanHandPoseObservationJointNameIndexTip = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameLittleDIP"); err == nil && ptr != 0 {
		vNHumanHandPoseObservationJointNameLittleDIP = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameLittleMCP"); err == nil && ptr != 0 {
		vNHumanHandPoseObservationJointNameLittleMCP = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameLittlePIP"); err == nil && ptr != 0 {
		vNHumanHandPoseObservationJointNameLittlePIP = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameLittleTip"); err == nil && ptr != 0 {
		vNHumanHandPoseObservationJointNameLittleTip = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameMiddleDIP"); err == nil && ptr != 0 {
		vNHumanHandPoseObservationJointNameMiddleDIP = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameMiddleMCP"); err == nil && ptr != 0 {
		vNHumanHandPoseObservationJointNameMiddleMCP = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameMiddlePIP"); err == nil && ptr != 0 {
		vNHumanHandPoseObservationJointNameMiddlePIP = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameMiddleTip"); err == nil && ptr != 0 {
		vNHumanHandPoseObservationJointNameMiddleTip = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameRingDIP"); err == nil && ptr != 0 {
		vNHumanHandPoseObservationJointNameRingDIP = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameRingMCP"); err == nil && ptr != 0 {
		vNHumanHandPoseObservationJointNameRingMCP = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameRingPIP"); err == nil && ptr != 0 {
		vNHumanHandPoseObservationJointNameRingPIP = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameRingTip"); err == nil && ptr != 0 {
		vNHumanHandPoseObservationJointNameRingTip = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameThumbCMC"); err == nil && ptr != 0 {
		vNHumanHandPoseObservationJointNameThumbCMC = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameThumbIP"); err == nil && ptr != 0 {
		vNHumanHandPoseObservationJointNameThumbIP = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameThumbMP"); err == nil && ptr != 0 {
		vNHumanHandPoseObservationJointNameThumbMP = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameThumbTip"); err == nil && ptr != 0 {
		vNHumanHandPoseObservationJointNameThumbTip = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointNameWrist"); err == nil && ptr != 0 {
		vNHumanHandPoseObservationJointNameWrist = *(*VNHumanHandPoseObservationJointName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointsGroupNameAll"); err == nil && ptr != 0 {
		vNHumanHandPoseObservationJointsGroupNameAll = *(*VNHumanHandPoseObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointsGroupNameIndexFinger"); err == nil && ptr != 0 {
		vNHumanHandPoseObservationJointsGroupNameIndexFinger = *(*VNHumanHandPoseObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointsGroupNameLittleFinger"); err == nil && ptr != 0 {
		vNHumanHandPoseObservationJointsGroupNameLittleFinger = *(*VNHumanHandPoseObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointsGroupNameMiddleFinger"); err == nil && ptr != 0 {
		vNHumanHandPoseObservationJointsGroupNameMiddleFinger = *(*VNHumanHandPoseObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointsGroupNameRingFinger"); err == nil && ptr != 0 {
		vNHumanHandPoseObservationJointsGroupNameRingFinger = *(*VNHumanHandPoseObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNHumanHandPoseObservationJointsGroupNameThumb"); err == nil && ptr != 0 {
		vNHumanHandPoseObservationJointsGroupNameThumb = *(*VNHumanHandPoseObservationJointsGroupName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNImageOptionCIContext"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				vNImageOptionCIContext = VNImageOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNImageOptionCameraIntrinsics"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				vNImageOptionCameraIntrinsics = VNImageOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "VNImageOptionProperties"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				vNImageOptionProperties = VNImageOption(objc.GoString(cstr))
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

type VNAnimalBodyPoseObservationJointNameValues struct{}

// VNAnimalBodyPoseObservationJointNames provides typed accessors for [VNAnimalBodyPoseObservationJointName] constants.
var VNAnimalBodyPoseObservationJointNames VNAnimalBodyPoseObservationJointNameValues

// LeftBackElbow returns A joint name that represents the back of the left elbow.
func (VNAnimalBodyPoseObservationJointNameValues) LeftBackElbow() VNAnimalBodyPoseObservationJointName { return vNAnimalBodyPoseObservationJointNameLeftBackElbow }

// LeftBackKnee returns A joint name that represents the back of the left knee.
func (VNAnimalBodyPoseObservationJointNameValues) LeftBackKnee() VNAnimalBodyPoseObservationJointName { return vNAnimalBodyPoseObservationJointNameLeftBackKnee }

// LeftBackPaw returns A joint name that represents the back of the left paw.
func (VNAnimalBodyPoseObservationJointNameValues) LeftBackPaw() VNAnimalBodyPoseObservationJointName { return vNAnimalBodyPoseObservationJointNameLeftBackPaw }

// LeftEarBottom returns A joint name that represents the bottom of the left ear.
func (VNAnimalBodyPoseObservationJointNameValues) LeftEarBottom() VNAnimalBodyPoseObservationJointName { return vNAnimalBodyPoseObservationJointNameLeftEarBottom }

// LeftEarMiddle returns A joint name that represents the middle of the left ear.
func (VNAnimalBodyPoseObservationJointNameValues) LeftEarMiddle() VNAnimalBodyPoseObservationJointName { return vNAnimalBodyPoseObservationJointNameLeftEarMiddle }

// LeftEarTop returns A joint name that represents the top of the left ear.
func (VNAnimalBodyPoseObservationJointNameValues) LeftEarTop() VNAnimalBodyPoseObservationJointName { return vNAnimalBodyPoseObservationJointNameLeftEarTop }

// LeftEye returns A joint name that represents the left eye.
func (VNAnimalBodyPoseObservationJointNameValues) LeftEye() VNAnimalBodyPoseObservationJointName { return vNAnimalBodyPoseObservationJointNameLeftEye }

// LeftFrontElbow returns A joint name that represents the front of the left elbow.
func (VNAnimalBodyPoseObservationJointNameValues) LeftFrontElbow() VNAnimalBodyPoseObservationJointName { return vNAnimalBodyPoseObservationJointNameLeftFrontElbow }

// LeftFrontKnee returns A joint name that represents the front of the left knee.
func (VNAnimalBodyPoseObservationJointNameValues) LeftFrontKnee() VNAnimalBodyPoseObservationJointName { return vNAnimalBodyPoseObservationJointNameLeftFrontKnee }

// LeftFrontPaw returns A joint name that represents the front of the left paw.
func (VNAnimalBodyPoseObservationJointNameValues) LeftFrontPaw() VNAnimalBodyPoseObservationJointName { return vNAnimalBodyPoseObservationJointNameLeftFrontPaw }

// Neck returns A joint name that represents the neck.
func (VNAnimalBodyPoseObservationJointNameValues) Neck() VNAnimalBodyPoseObservationJointName { return vNAnimalBodyPoseObservationJointNameNeck }

// Nose returns A joint name that represents the nose.
func (VNAnimalBodyPoseObservationJointNameValues) Nose() VNAnimalBodyPoseObservationJointName { return vNAnimalBodyPoseObservationJointNameNose }

// RightBackElbow returns A joint name that represents the back of the right elbow.
func (VNAnimalBodyPoseObservationJointNameValues) RightBackElbow() VNAnimalBodyPoseObservationJointName { return vNAnimalBodyPoseObservationJointNameRightBackElbow }

// RightBackKnee returns A joint name that represents the back of the right knee.
func (VNAnimalBodyPoseObservationJointNameValues) RightBackKnee() VNAnimalBodyPoseObservationJointName { return vNAnimalBodyPoseObservationJointNameRightBackKnee }

// RightBackPaw returns A joint name that represents the back of the right paw.
func (VNAnimalBodyPoseObservationJointNameValues) RightBackPaw() VNAnimalBodyPoseObservationJointName { return vNAnimalBodyPoseObservationJointNameRightBackPaw }

// RightEarBottom returns A joint name that represents the bottom of the right ear.
func (VNAnimalBodyPoseObservationJointNameValues) RightEarBottom() VNAnimalBodyPoseObservationJointName { return vNAnimalBodyPoseObservationJointNameRightEarBottom }

// RightEarMiddle returns A joint name that represents the middle of the right ear.
func (VNAnimalBodyPoseObservationJointNameValues) RightEarMiddle() VNAnimalBodyPoseObservationJointName { return vNAnimalBodyPoseObservationJointNameRightEarMiddle }

// RightEarTop returns A joint name that represents the top of the right ear.
func (VNAnimalBodyPoseObservationJointNameValues) RightEarTop() VNAnimalBodyPoseObservationJointName { return vNAnimalBodyPoseObservationJointNameRightEarTop }

// RightEye returns A joint name that represents the right eye.
func (VNAnimalBodyPoseObservationJointNameValues) RightEye() VNAnimalBodyPoseObservationJointName { return vNAnimalBodyPoseObservationJointNameRightEye }

// RightFrontElbow returns A joint name that represents the front of the right elbow.
func (VNAnimalBodyPoseObservationJointNameValues) RightFrontElbow() VNAnimalBodyPoseObservationJointName { return vNAnimalBodyPoseObservationJointNameRightFrontElbow }

// RightFrontKnee returns A joint name that represents the front of the right knee.
func (VNAnimalBodyPoseObservationJointNameValues) RightFrontKnee() VNAnimalBodyPoseObservationJointName { return vNAnimalBodyPoseObservationJointNameRightFrontKnee }

// RightFrontPaw returns A joint name that represents the front of the right paw.
func (VNAnimalBodyPoseObservationJointNameValues) RightFrontPaw() VNAnimalBodyPoseObservationJointName { return vNAnimalBodyPoseObservationJointNameRightFrontPaw }

// TailBottom returns A joint name that represents the bottom of the tail.
func (VNAnimalBodyPoseObservationJointNameValues) TailBottom() VNAnimalBodyPoseObservationJointName { return vNAnimalBodyPoseObservationJointNameTailBottom }

// TailMiddle returns A joint name that represents the middle of the tail.
func (VNAnimalBodyPoseObservationJointNameValues) TailMiddle() VNAnimalBodyPoseObservationJointName { return vNAnimalBodyPoseObservationJointNameTailMiddle }

// TailTop returns A joint name that represents the top of the tail.
func (VNAnimalBodyPoseObservationJointNameValues) TailTop() VNAnimalBodyPoseObservationJointName { return vNAnimalBodyPoseObservationJointNameTailTop }


type VNAnimalBodyPoseObservationJointsGroupNameValues struct{}

// VNAnimalBodyPoseObservationJointsGroupNames provides typed accessors for [VNAnimalBodyPoseObservationJointsGroupName] constants.
var VNAnimalBodyPoseObservationJointsGroupNames VNAnimalBodyPoseObservationJointsGroupNameValues

// All returns A group name that represents all joints.
func (VNAnimalBodyPoseObservationJointsGroupNameValues) All() VNAnimalBodyPoseObservationJointsGroupName { return vNAnimalBodyPoseObservationJointsGroupNameAll }

// Forelegs returns A group name that represents the forelegs.
func (VNAnimalBodyPoseObservationJointsGroupNameValues) Forelegs() VNAnimalBodyPoseObservationJointsGroupName { return vNAnimalBodyPoseObservationJointsGroupNameForelegs }

// Head returns A group name that represents the head.
func (VNAnimalBodyPoseObservationJointsGroupNameValues) Head() VNAnimalBodyPoseObservationJointsGroupName { return vNAnimalBodyPoseObservationJointsGroupNameHead }

// Hindlegs returns A group name that represents the hindlegs.
func (VNAnimalBodyPoseObservationJointsGroupNameValues) Hindlegs() VNAnimalBodyPoseObservationJointsGroupName { return vNAnimalBodyPoseObservationJointsGroupNameHindlegs }

// Tail returns A group name that represents the tail.
func (VNAnimalBodyPoseObservationJointsGroupNameValues) Tail() VNAnimalBodyPoseObservationJointsGroupName { return vNAnimalBodyPoseObservationJointsGroupNameTail }

// Trunk returns A group name that represents the trunk.
func (VNAnimalBodyPoseObservationJointsGroupNameValues) Trunk() VNAnimalBodyPoseObservationJointsGroupName { return vNAnimalBodyPoseObservationJointsGroupNameTrunk }


type VNAnimalIdentifierValues struct{}

// VNAnimalIdentifiers provides typed accessors for [VNAnimalIdentifier] constants.
var VNAnimalIdentifiers VNAnimalIdentifierValues

// Cat returns An animal identifier for cats.
func (VNAnimalIdentifierValues) Cat() VNAnimalIdentifier { return vNAnimalIdentifierCat }

// Dog returns An animal identifier for dogs.
func (VNAnimalIdentifierValues) Dog() VNAnimalIdentifier { return vNAnimalIdentifierDog }


type VNBarcodeSymbologyValues struct{}

// VNBarcodeSymbologys provides typed accessors for [VNBarcodeSymbology] constants.
var VNBarcodeSymbologys VNBarcodeSymbologyValues

// Aztec returns A constant that indicates Aztec symbology.
func (VNBarcodeSymbologyValues) Aztec() VNBarcodeSymbology { return vNBarcodeSymbologyAztec }

// Codabar returns A constant that indicates Codabar symbology.
func (VNBarcodeSymbologyValues) Codabar() VNBarcodeSymbology { return vNBarcodeSymbologyCodabar }

// Code128 returns A constant that indicates Code 128 symbology.
func (VNBarcodeSymbologyValues) Code128() VNBarcodeSymbology { return vNBarcodeSymbologyCode128 }

// Code39 returns A constant that indicates Code 39 symbology.
func (VNBarcodeSymbologyValues) Code39() VNBarcodeSymbology { return vNBarcodeSymbologyCode39 }

// Code39Checksum returns A constant that indicates Code 39 symbology with a checksum.
func (VNBarcodeSymbologyValues) Code39Checksum() VNBarcodeSymbology { return vNBarcodeSymbologyCode39Checksum }

// Code39FullASCII returns A constant that indicates Code 39 Full ASCII symbology.
func (VNBarcodeSymbologyValues) Code39FullASCII() VNBarcodeSymbology { return vNBarcodeSymbologyCode39FullASCII }

// Code39FullASCIIChecksum returns A constant that indicates Code 39 Full ASCII symbology with a checksum.
func (VNBarcodeSymbologyValues) Code39FullASCIIChecksum() VNBarcodeSymbology { return vNBarcodeSymbologyCode39FullASCIIChecksum }

// Code93 returns A constant that indicates Code 93 symbology.
func (VNBarcodeSymbologyValues) Code93() VNBarcodeSymbology { return vNBarcodeSymbologyCode93 }

// Code93i returns A constant that indicates Code 93i symbology.
func (VNBarcodeSymbologyValues) Code93i() VNBarcodeSymbology { return vNBarcodeSymbologyCode93i }

// DataMatrix returns A constant that indicates Data Matrix symbology.
func (VNBarcodeSymbologyValues) DataMatrix() VNBarcodeSymbology { return vNBarcodeSymbologyDataMatrix }

// EAN13 returns A constant that indicates EAN-13 symbology.
func (VNBarcodeSymbologyValues) EAN13() VNBarcodeSymbology { return vNBarcodeSymbologyEAN13 }

// EAN8 returns A constant that indicates EAN-8 symbology.
func (VNBarcodeSymbologyValues) EAN8() VNBarcodeSymbology { return vNBarcodeSymbologyEAN8 }

// GS1DataBar returns A constant that indicates GS1 DataBar symbology.
func (VNBarcodeSymbologyValues) GS1DataBar() VNBarcodeSymbology { return vNBarcodeSymbologyGS1DataBar }

// GS1DataBarExpanded returns A constant that indicates GS1 DataBar Expanded symbology.
func (VNBarcodeSymbologyValues) GS1DataBarExpanded() VNBarcodeSymbology { return vNBarcodeSymbologyGS1DataBarExpanded }

// GS1DataBarLimited returns A constant that indicates GS1 DataBar Limited symbology.
func (VNBarcodeSymbologyValues) GS1DataBarLimited() VNBarcodeSymbology { return vNBarcodeSymbologyGS1DataBarLimited }

// I2of5 returns A constant that indicates Interleaved 2 of 5 (ITF) symbology.
func (VNBarcodeSymbologyValues) I2of5() VNBarcodeSymbology { return vNBarcodeSymbologyI2of5 }

// I2of5Checksum returns A constant that indicates Interleaved 2 of 5 (ITF) symbology with a checksum.
func (VNBarcodeSymbologyValues) I2of5Checksum() VNBarcodeSymbology { return vNBarcodeSymbologyI2of5Checksum }

// ITF14 returns A constant that indicates ITF-14 symbology.
func (VNBarcodeSymbologyValues) ITF14() VNBarcodeSymbology { return vNBarcodeSymbologyITF14 }

// MSIPlessey returns A constant that indicates Modified Plessey symbology.
func (VNBarcodeSymbologyValues) MSIPlessey() VNBarcodeSymbology { return vNBarcodeSymbologyMSIPlessey }

// MicroPDF417 returns A constant that indicates MicroPDF417 symbology.
func (VNBarcodeSymbologyValues) MicroPDF417() VNBarcodeSymbology { return vNBarcodeSymbologyMicroPDF417 }

// MicroQR returns A constant that indicates MicroQR symbology.
func (VNBarcodeSymbologyValues) MicroQR() VNBarcodeSymbology { return vNBarcodeSymbologyMicroQR }

// PDF417 returns A constant that indicates PDF417 symbology.
func (VNBarcodeSymbologyValues) PDF417() VNBarcodeSymbology { return vNBarcodeSymbologyPDF417 }

// QR returns A constant that indicates Quick Response (QR) symbology.
func (VNBarcodeSymbologyValues) QR() VNBarcodeSymbology { return vNBarcodeSymbologyQR }

// UPCE returns A constant that indicates UPC-E symbology.
func (VNBarcodeSymbologyValues) UPCE() VNBarcodeSymbology { return vNBarcodeSymbologyUPCE }


type VNComputeStageValues struct{}

// VNComputeStages provides typed accessors for [VNComputeStage] constants.
var VNComputeStages VNComputeStageValues

// Main returns A stage that represents where the system performs the main functionality.
func (VNComputeStageValues) Main() VNComputeStage { return vNComputeStageMain }

// PostProcessing returns A stage that represents where the system performs additional analysis from the main compute stage.
func (VNComputeStageValues) PostProcessing() VNComputeStage { return vNComputeStagePostProcessing }


type VNHumanBodyPose3DObservationJointNameValues struct{}

// VNHumanBodyPose3DObservationJointNames provides typed accessors for [VNHumanBodyPose3DObservationJointName] constants.
var VNHumanBodyPose3DObservationJointNames VNHumanBodyPose3DObservationJointNameValues

// CenterHead returns A joint name that represents the center of the head.
func (VNHumanBodyPose3DObservationJointNameValues) CenterHead() VNHumanBodyPose3DObservationJointName { return vNHumanBodyPose3DObservationJointNameCenterHead }

// CenterShoulder returns A joint name that represents the point between the shoulders.
func (VNHumanBodyPose3DObservationJointNameValues) CenterShoulder() VNHumanBodyPose3DObservationJointName { return vNHumanBodyPose3DObservationJointNameCenterShoulder }

// LeftAnkle returns A joint name that represents the left ankle.
func (VNHumanBodyPose3DObservationJointNameValues) LeftAnkle() VNHumanBodyPose3DObservationJointName { return vNHumanBodyPose3DObservationJointNameLeftAnkle }

// LeftElbow returns A joint name that represents the left elbow.
func (VNHumanBodyPose3DObservationJointNameValues) LeftElbow() VNHumanBodyPose3DObservationJointName { return vNHumanBodyPose3DObservationJointNameLeftElbow }

// LeftHip returns A joint name that represents the left hip.
func (VNHumanBodyPose3DObservationJointNameValues) LeftHip() VNHumanBodyPose3DObservationJointName { return vNHumanBodyPose3DObservationJointNameLeftHip }

// LeftKnee returns A joint name that represents the left knee.
func (VNHumanBodyPose3DObservationJointNameValues) LeftKnee() VNHumanBodyPose3DObservationJointName { return vNHumanBodyPose3DObservationJointNameLeftKnee }

// LeftShoulder returns A joint name that represents the left shoulder.
func (VNHumanBodyPose3DObservationJointNameValues) LeftShoulder() VNHumanBodyPose3DObservationJointName { return vNHumanBodyPose3DObservationJointNameLeftShoulder }

// LeftWrist returns A joint name that represents the left wrist.
func (VNHumanBodyPose3DObservationJointNameValues) LeftWrist() VNHumanBodyPose3DObservationJointName { return vNHumanBodyPose3DObservationJointNameLeftWrist }

// RightAnkle returns A joint name that represents the right ankle.
func (VNHumanBodyPose3DObservationJointNameValues) RightAnkle() VNHumanBodyPose3DObservationJointName { return vNHumanBodyPose3DObservationJointNameRightAnkle }

// RightElbow returns A joint name that represents the right elbow.
func (VNHumanBodyPose3DObservationJointNameValues) RightElbow() VNHumanBodyPose3DObservationJointName { return vNHumanBodyPose3DObservationJointNameRightElbow }

// RightHip returns A joint name that represents the right hip.
func (VNHumanBodyPose3DObservationJointNameValues) RightHip() VNHumanBodyPose3DObservationJointName { return vNHumanBodyPose3DObservationJointNameRightHip }

// RightKnee returns A joint name that represents the right knee.
func (VNHumanBodyPose3DObservationJointNameValues) RightKnee() VNHumanBodyPose3DObservationJointName { return vNHumanBodyPose3DObservationJointNameRightKnee }

// RightShoulder returns A joint name that represents the right shoulder.
func (VNHumanBodyPose3DObservationJointNameValues) RightShoulder() VNHumanBodyPose3DObservationJointName { return vNHumanBodyPose3DObservationJointNameRightShoulder }

// RightWrist returns A joint name that represents the right wrist.
func (VNHumanBodyPose3DObservationJointNameValues) RightWrist() VNHumanBodyPose3DObservationJointName { return vNHumanBodyPose3DObservationJointNameRightWrist }

// Root returns A joint name that represents the point between the left hip and right hip.
func (VNHumanBodyPose3DObservationJointNameValues) Root() VNHumanBodyPose3DObservationJointName { return vNHumanBodyPose3DObservationJointNameRoot }

// Spine returns A joint name that represents the spine.
func (VNHumanBodyPose3DObservationJointNameValues) Spine() VNHumanBodyPose3DObservationJointName { return vNHumanBodyPose3DObservationJointNameSpine }

// TopHead returns A joint name that represents the top of the head.
func (VNHumanBodyPose3DObservationJointNameValues) TopHead() VNHumanBodyPose3DObservationJointName { return vNHumanBodyPose3DObservationJointNameTopHead }


type VNHumanBodyPose3DObservationJointsGroupNameValues struct{}

// VNHumanBodyPose3DObservationJointsGroupNames provides typed accessors for [VNHumanBodyPose3DObservationJointsGroupName] constants.
var VNHumanBodyPose3DObservationJointsGroupNames VNHumanBodyPose3DObservationJointsGroupNameValues

// All returns A group name that represents all joints.
func (VNHumanBodyPose3DObservationJointsGroupNameValues) All() VNHumanBodyPose3DObservationJointsGroupName { return vNHumanBodyPose3DObservationJointsGroupNameAll }

// Head returns A group name that represents the head joints.
func (VNHumanBodyPose3DObservationJointsGroupNameValues) Head() VNHumanBodyPose3DObservationJointsGroupName { return vNHumanBodyPose3DObservationJointsGroupNameHead }

// LeftArm returns A group name that represents the left arm joints.
func (VNHumanBodyPose3DObservationJointsGroupNameValues) LeftArm() VNHumanBodyPose3DObservationJointsGroupName { return vNHumanBodyPose3DObservationJointsGroupNameLeftArm }

// LeftLeg returns A group name that represents the left leg joints.
func (VNHumanBodyPose3DObservationJointsGroupNameValues) LeftLeg() VNHumanBodyPose3DObservationJointsGroupName { return vNHumanBodyPose3DObservationJointsGroupNameLeftLeg }

// RightArm returns A group name that represents the right arm joints.
func (VNHumanBodyPose3DObservationJointsGroupNameValues) RightArm() VNHumanBodyPose3DObservationJointsGroupName { return vNHumanBodyPose3DObservationJointsGroupNameRightArm }

// RightLeg returns A group name that represents the right leg joints.
func (VNHumanBodyPose3DObservationJointsGroupNameValues) RightLeg() VNHumanBodyPose3DObservationJointsGroupName { return vNHumanBodyPose3DObservationJointsGroupNameRightLeg }

// Torso returns A group name that represents the torso joints.
func (VNHumanBodyPose3DObservationJointsGroupNameValues) Torso() VNHumanBodyPose3DObservationJointsGroupName { return vNHumanBodyPose3DObservationJointsGroupNameTorso }


type VNHumanBodyPoseObservationJointNameValues struct{}

// VNHumanBodyPoseObservationJointNames provides typed accessors for [VNHumanBodyPoseObservationJointName] constants.
var VNHumanBodyPoseObservationJointNames VNHumanBodyPoseObservationJointNameValues

// LeftAnkle returns The left ankle.
func (VNHumanBodyPoseObservationJointNameValues) LeftAnkle() VNHumanBodyPoseObservationJointName { return vNHumanBodyPoseObservationJointNameLeftAnkle }

// LeftEar returns The left ear.
func (VNHumanBodyPoseObservationJointNameValues) LeftEar() VNHumanBodyPoseObservationJointName { return vNHumanBodyPoseObservationJointNameLeftEar }

// LeftElbow returns The left elbow.
func (VNHumanBodyPoseObservationJointNameValues) LeftElbow() VNHumanBodyPoseObservationJointName { return vNHumanBodyPoseObservationJointNameLeftElbow }

// LeftEye returns The left eye.
func (VNHumanBodyPoseObservationJointNameValues) LeftEye() VNHumanBodyPoseObservationJointName { return vNHumanBodyPoseObservationJointNameLeftEye }

// LeftHip returns The left hip.
func (VNHumanBodyPoseObservationJointNameValues) LeftHip() VNHumanBodyPoseObservationJointName { return vNHumanBodyPoseObservationJointNameLeftHip }

// LeftKnee returns The left knee.
func (VNHumanBodyPoseObservationJointNameValues) LeftKnee() VNHumanBodyPoseObservationJointName { return vNHumanBodyPoseObservationJointNameLeftKnee }

// LeftShoulder returns The left shoulder.
func (VNHumanBodyPoseObservationJointNameValues) LeftShoulder() VNHumanBodyPoseObservationJointName { return vNHumanBodyPoseObservationJointNameLeftShoulder }

// LeftWrist returns The left wrist.
func (VNHumanBodyPoseObservationJointNameValues) LeftWrist() VNHumanBodyPoseObservationJointName { return vNHumanBodyPoseObservationJointNameLeftWrist }

// Neck returns The neck.
func (VNHumanBodyPoseObservationJointNameValues) Neck() VNHumanBodyPoseObservationJointName { return vNHumanBodyPoseObservationJointNameNeck }

// Nose returns The nose.
func (VNHumanBodyPoseObservationJointNameValues) Nose() VNHumanBodyPoseObservationJointName { return vNHumanBodyPoseObservationJointNameNose }

// RightAnkle returns The right ankle.
func (VNHumanBodyPoseObservationJointNameValues) RightAnkle() VNHumanBodyPoseObservationJointName { return vNHumanBodyPoseObservationJointNameRightAnkle }

// RightEar returns The right ear.
func (VNHumanBodyPoseObservationJointNameValues) RightEar() VNHumanBodyPoseObservationJointName { return vNHumanBodyPoseObservationJointNameRightEar }

// RightElbow returns The right elbow.
func (VNHumanBodyPoseObservationJointNameValues) RightElbow() VNHumanBodyPoseObservationJointName { return vNHumanBodyPoseObservationJointNameRightElbow }

// RightEye returns The right eye.
func (VNHumanBodyPoseObservationJointNameValues) RightEye() VNHumanBodyPoseObservationJointName { return vNHumanBodyPoseObservationJointNameRightEye }

// RightHip returns The right hip.
func (VNHumanBodyPoseObservationJointNameValues) RightHip() VNHumanBodyPoseObservationJointName { return vNHumanBodyPoseObservationJointNameRightHip }

// RightKnee returns The right knee.
func (VNHumanBodyPoseObservationJointNameValues) RightKnee() VNHumanBodyPoseObservationJointName { return vNHumanBodyPoseObservationJointNameRightKnee }

// RightShoulder returns The right shoulder.
func (VNHumanBodyPoseObservationJointNameValues) RightShoulder() VNHumanBodyPoseObservationJointName { return vNHumanBodyPoseObservationJointNameRightShoulder }

// RightWrist returns The right wrist.
func (VNHumanBodyPoseObservationJointNameValues) RightWrist() VNHumanBodyPoseObservationJointName { return vNHumanBodyPoseObservationJointNameRightWrist }

// Root returns The root (waist).
func (VNHumanBodyPoseObservationJointNameValues) Root() VNHumanBodyPoseObservationJointName { return vNHumanBodyPoseObservationJointNameRoot }


type VNHumanBodyPoseObservationJointsGroupNameValues struct{}

// VNHumanBodyPoseObservationJointsGroupNames provides typed accessors for [VNHumanBodyPoseObservationJointsGroupName] constants.
var VNHumanBodyPoseObservationJointsGroupNames VNHumanBodyPoseObservationJointsGroupNameValues

// All returns All body point groups.
func (VNHumanBodyPoseObservationJointsGroupNameValues) All() VNHumanBodyPoseObservationJointsGroupName { return vNHumanBodyPoseObservationJointsGroupNameAll }

// Face returns The face.
func (VNHumanBodyPoseObservationJointsGroupNameValues) Face() VNHumanBodyPoseObservationJointsGroupName { return vNHumanBodyPoseObservationJointsGroupNameFace }

// LeftArm returns The left arm.
func (VNHumanBodyPoseObservationJointsGroupNameValues) LeftArm() VNHumanBodyPoseObservationJointsGroupName { return vNHumanBodyPoseObservationJointsGroupNameLeftArm }

// LeftLeg returns The left leg.
func (VNHumanBodyPoseObservationJointsGroupNameValues) LeftLeg() VNHumanBodyPoseObservationJointsGroupName { return vNHumanBodyPoseObservationJointsGroupNameLeftLeg }

// RightArm returns The right arm.
func (VNHumanBodyPoseObservationJointsGroupNameValues) RightArm() VNHumanBodyPoseObservationJointsGroupName { return vNHumanBodyPoseObservationJointsGroupNameRightArm }

// RightLeg returns The right leg.
func (VNHumanBodyPoseObservationJointsGroupNameValues) RightLeg() VNHumanBodyPoseObservationJointsGroupName { return vNHumanBodyPoseObservationJointsGroupNameRightLeg }

// Torso returns The torso.
func (VNHumanBodyPoseObservationJointsGroupNameValues) Torso() VNHumanBodyPoseObservationJointsGroupName { return vNHumanBodyPoseObservationJointsGroupNameTorso }


type VNHumanHandPoseObservationJointNameValues struct{}

// VNHumanHandPoseObservationJointNames provides typed accessors for [VNHumanHandPoseObservationJointName] constants.
var VNHumanHandPoseObservationJointNames VNHumanHandPoseObservationJointNameValues

// IndexDIP returns The index finger’s distal interphalangeal (DIP) joint.
func (VNHumanHandPoseObservationJointNameValues) IndexDIP() VNHumanHandPoseObservationJointName { return vNHumanHandPoseObservationJointNameIndexDIP }

// IndexMCP returns The index finger’s metacarpophalangeal (MCP) joint.
func (VNHumanHandPoseObservationJointNameValues) IndexMCP() VNHumanHandPoseObservationJointName { return vNHumanHandPoseObservationJointNameIndexMCP }

// IndexPIP returns The index finger’s proximal interphalangeal (PIP) joint.
func (VNHumanHandPoseObservationJointNameValues) IndexPIP() VNHumanHandPoseObservationJointName { return vNHumanHandPoseObservationJointNameIndexPIP }

// IndexTip returns The tip of the index finger.
func (VNHumanHandPoseObservationJointNameValues) IndexTip() VNHumanHandPoseObservationJointName { return vNHumanHandPoseObservationJointNameIndexTip }

// LittleDIP returns The little finger’s distal interphalangeal (DIP) joint.
func (VNHumanHandPoseObservationJointNameValues) LittleDIP() VNHumanHandPoseObservationJointName { return vNHumanHandPoseObservationJointNameLittleDIP }

// LittleMCP returns The little finger’s metacarpophalangeal (MCP) joint.
func (VNHumanHandPoseObservationJointNameValues) LittleMCP() VNHumanHandPoseObservationJointName { return vNHumanHandPoseObservationJointNameLittleMCP }

// LittlePIP returns The little finger’s proximal interphalangeal (PIP) joint.
func (VNHumanHandPoseObservationJointNameValues) LittlePIP() VNHumanHandPoseObservationJointName { return vNHumanHandPoseObservationJointNameLittlePIP }

// LittleTip returns The tip of the little finger.
func (VNHumanHandPoseObservationJointNameValues) LittleTip() VNHumanHandPoseObservationJointName { return vNHumanHandPoseObservationJointNameLittleTip }

// MiddleDIP returns The middle finger’s distal interphalangeal (DIP) joint.
func (VNHumanHandPoseObservationJointNameValues) MiddleDIP() VNHumanHandPoseObservationJointName { return vNHumanHandPoseObservationJointNameMiddleDIP }

// MiddleMCP returns The middle finger’s metacarpophalangeal (MCP) joint.
func (VNHumanHandPoseObservationJointNameValues) MiddleMCP() VNHumanHandPoseObservationJointName { return vNHumanHandPoseObservationJointNameMiddleMCP }

// MiddlePIP returns The middle finger’s proximal interphalangeal (PIP) joint.
func (VNHumanHandPoseObservationJointNameValues) MiddlePIP() VNHumanHandPoseObservationJointName { return vNHumanHandPoseObservationJointNameMiddlePIP }

// MiddleTip returns The tip of the middle finger.
func (VNHumanHandPoseObservationJointNameValues) MiddleTip() VNHumanHandPoseObservationJointName { return vNHumanHandPoseObservationJointNameMiddleTip }

// RingDIP returns The ring finger’s distal interphalangeal (DIP) joint.
func (VNHumanHandPoseObservationJointNameValues) RingDIP() VNHumanHandPoseObservationJointName { return vNHumanHandPoseObservationJointNameRingDIP }

// RingMCP returns The ring finger’s metacarpophalangeal (MCP) joint.
func (VNHumanHandPoseObservationJointNameValues) RingMCP() VNHumanHandPoseObservationJointName { return vNHumanHandPoseObservationJointNameRingMCP }

// RingPIP returns The ring finger’s proximal interphalangeal (PIP) joint.
func (VNHumanHandPoseObservationJointNameValues) RingPIP() VNHumanHandPoseObservationJointName { return vNHumanHandPoseObservationJointNameRingPIP }

// RingTip returns The tip of the ring finger.
func (VNHumanHandPoseObservationJointNameValues) RingTip() VNHumanHandPoseObservationJointName { return vNHumanHandPoseObservationJointNameRingTip }

// ThumbCMC returns The thumb’s carpometacarpal (CMC) joint.
func (VNHumanHandPoseObservationJointNameValues) ThumbCMC() VNHumanHandPoseObservationJointName { return vNHumanHandPoseObservationJointNameThumbCMC }

// ThumbIP returns The thumb’s interphalangeal (IP) joint.
func (VNHumanHandPoseObservationJointNameValues) ThumbIP() VNHumanHandPoseObservationJointName { return vNHumanHandPoseObservationJointNameThumbIP }

// ThumbMP returns The thumb’s metacarpophalangeal (MP) joint.
func (VNHumanHandPoseObservationJointNameValues) ThumbMP() VNHumanHandPoseObservationJointName { return vNHumanHandPoseObservationJointNameThumbMP }

// ThumbTip returns The tip of the thumb.
func (VNHumanHandPoseObservationJointNameValues) ThumbTip() VNHumanHandPoseObservationJointName { return vNHumanHandPoseObservationJointNameThumbTip }

// Wrist returns The wrist.
func (VNHumanHandPoseObservationJointNameValues) Wrist() VNHumanHandPoseObservationJointName { return vNHumanHandPoseObservationJointNameWrist }


type VNHumanHandPoseObservationJointsGroupNameValues struct{}

// VNHumanHandPoseObservationJointsGroupNames provides typed accessors for [VNHumanHandPoseObservationJointsGroupName] constants.
var VNHumanHandPoseObservationJointsGroupNames VNHumanHandPoseObservationJointsGroupNameValues

// All returns All hand group names.
func (VNHumanHandPoseObservationJointsGroupNameValues) All() VNHumanHandPoseObservationJointsGroupName { return vNHumanHandPoseObservationJointsGroupNameAll }

// IndexFinger returns The index finger.
func (VNHumanHandPoseObservationJointsGroupNameValues) IndexFinger() VNHumanHandPoseObservationJointsGroupName { return vNHumanHandPoseObservationJointsGroupNameIndexFinger }

// LittleFinger returns The little finger.
func (VNHumanHandPoseObservationJointsGroupNameValues) LittleFinger() VNHumanHandPoseObservationJointsGroupName { return vNHumanHandPoseObservationJointsGroupNameLittleFinger }

// MiddleFinger returns The middle finger.
func (VNHumanHandPoseObservationJointsGroupNameValues) MiddleFinger() VNHumanHandPoseObservationJointsGroupName { return vNHumanHandPoseObservationJointsGroupNameMiddleFinger }

// RingFinger returns The ring finger.
func (VNHumanHandPoseObservationJointsGroupNameValues) RingFinger() VNHumanHandPoseObservationJointsGroupName { return vNHumanHandPoseObservationJointsGroupNameRingFinger }

// Thumb returns The thumb.
func (VNHumanHandPoseObservationJointsGroupNameValues) Thumb() VNHumanHandPoseObservationJointsGroupName { return vNHumanHandPoseObservationJointsGroupNameThumb }


type VNImageOptionValues struct{}

// VNImageOptions provides typed accessors for [VNImageOption] constants.
var VNImageOptions VNImageOptionValues

// CIContext returns An option key to specify the context to use in the handler’s Core Image operations.
func (VNImageOptionValues) CIContext() VNImageOption { return vNImageOptionCIContext }

// CameraIntrinsics returns An option to specify the camera intrinstics.
func (VNImageOptionValues) CameraIntrinsics() VNImageOption { return vNImageOptionCameraIntrinsics }

// Properties returns The dictionary from the image source that contains the metadata for algorithms like horizon detection.
func (VNImageOptionValues) Properties() VNImageOption { return vNImageOptionProperties }


