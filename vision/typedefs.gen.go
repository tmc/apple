// Code generated from Apple documentation. DO NOT EDIT.

package vision

import (
	"github.com/tmc/apple/foundation"
)

// VNAnimalBodyPoseObservationJointName is the joint names for an animal body pose.
//
// See: https://developer.apple.com/documentation/Vision/VNAnimalBodyPoseObservation/JointName
type VNAnimalBodyPoseObservationJointName = string

// VNAnimalBodyPoseObservationJointsGroupName is the joint group names for an animal body pose.
//
// See: https://developer.apple.com/documentation/Vision/VNAnimalBodyPoseObservation/JointsGroupName
type VNAnimalBodyPoseObservationJointsGroupName = string

// VNAnimalIdentifier is an animal identifier string.
//
// See: https://developer.apple.com/documentation/Vision/VNAnimalIdentifier
type VNAnimalIdentifier = string

// VNAspectRatio is a type alias for expressing rectangle aspect ratios in Vision.
//
// See: https://developer.apple.com/documentation/Vision/VNAspectRatio
type VNAspectRatio = float32

// VNBarcodeSymbology is the barcode symbologies that the framework detects.
//
// See: https://developer.apple.com/documentation/Vision/VNBarcodeSymbology
type VNBarcodeSymbology = string

// VNComputeStage is types that represent the compute stage.
//
// See: https://developer.apple.com/documentation/Vision/VNComputeStage
type VNComputeStage = string

// VNConfidence is a type alias for the confidence value of an observation.
//
// See: https://developer.apple.com/documentation/Vision/VNConfidence
type VNConfidence = float32

// VNDegrees is a typealias for expressing tolerance angles in Vision.
//
// See: https://developer.apple.com/documentation/Vision/VNDegrees
type VNDegrees = float32

// VNHumanBodyPose3DObservationJointName is the joint names for a 3D body pose.
//
// See: https://developer.apple.com/documentation/Vision/VNHumanBodyPose3DObservation/JointName
type VNHumanBodyPose3DObservationJointName = string

// VNHumanBodyPose3DObservationJointsGroupName is the joint group names for a 3D body pose.
//
// See: https://developer.apple.com/documentation/Vision/VNHumanBodyPose3DObservation/JointsGroupName
type VNHumanBodyPose3DObservationJointsGroupName = string

// VNHumanBodyPoseObservationJointName is the supported joint names for the body pose.
//
// See: https://developer.apple.com/documentation/Vision/VNHumanBodyPoseObservation/JointName
type VNHumanBodyPoseObservationJointName = string

// VNHumanBodyPoseObservationJointsGroupName is the supported joint group names for the body pose.
//
// See: https://developer.apple.com/documentation/Vision/VNHumanBodyPoseObservation/JointsGroupName
type VNHumanBodyPoseObservationJointsGroupName = string

// VNHumanHandPoseObservationJointName is the supported joint names for the hand pose.
//
// See: https://developer.apple.com/documentation/Vision/VNHumanHandPoseObservation/JointName
type VNHumanHandPoseObservationJointName = string

// VNHumanHandPoseObservationJointsGroupName is the supported joint group names for the hand pose.
//
// See: https://developer.apple.com/documentation/Vision/VNHumanHandPoseObservation/JointsGroupName
type VNHumanHandPoseObservationJointsGroupName = string

// VNImageOption is an option key passed into an image request handler that takes an auxiliary image.
//
// See: https://developer.apple.com/documentation/Vision/VNImageOption
type VNImageOption = string

// VNRecognizedPointGroupKey is the data type for all recognized-point group keys.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizedPointGroupKey
type VNRecognizedPointGroupKey = string

// VNRecognizedPointKey is the data type for all recognized point keys.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizedPointKey
type VNRecognizedPointKey = string

// VNRequestCompletionHandler is a type alias to encapsulate the syntax for the completion handler the system calls after the request finishes processing.
//
// See: https://developer.apple.com/documentation/Vision/VNRequestCompletionHandler
type VNRequestCompletionHandler = func(VNRequest, foundation.NSError)

// VNRequestProgressHandler is a block executed at intervals during the processing of a Vision request.
//
// See: https://developer.apple.com/documentation/Vision/VNRequestProgressHandler
type VNRequestProgressHandler = func(VNRequest, float64, foundation.NSError)

