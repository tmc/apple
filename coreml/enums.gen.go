// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/CoreML/MLComputeUnits
type MLComputeUnits int

const (
	// MLComputeUnitsAll: The option you choose to allow the model to use all compute units available, including the neural engine.
	MLComputeUnitsAll MLComputeUnits = 2
	// MLComputeUnitsCPUAndGPU: The option you choose to allow the model to use both the CPU and GPU, but not the neural engine.
	MLComputeUnitsCPUAndGPU MLComputeUnits = 1
	// MLComputeUnitsCPUAndNeuralEngine: The option you choose to allow the model to use both the CPU and neural engine, but not the GPU.
	MLComputeUnitsCPUAndNeuralEngine MLComputeUnits = 3
	// MLComputeUnitsCPUOnly: The option you choose to limit the model to only use the CPU.
	MLComputeUnitsCPUOnly MLComputeUnits = 0
)

func (e MLComputeUnits) String() string {
	switch e {
	case MLComputeUnitsAll:
		return "MLComputeUnitsAll"
	case MLComputeUnitsCPUAndGPU:
		return "MLComputeUnitsCPUAndGPU"
	case MLComputeUnitsCPUAndNeuralEngine:
		return "MLComputeUnitsCPUAndNeuralEngine"
	case MLComputeUnitsCPUOnly:
		return "MLComputeUnitsCPUOnly"
	default:
		return fmt.Sprintf("MLComputeUnits(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureType
type MLFeatureType int

const (
	// MLFeatureTypeDictionary: The type for dictionary features and feature values.
	MLFeatureTypeDictionary MLFeatureType = 6
	// MLFeatureTypeDouble: The type for double features and feature values.
	MLFeatureTypeDouble MLFeatureType = 2
	// MLFeatureTypeImage: The type for image features and feature values.
	MLFeatureTypeImage MLFeatureType = 4
	// MLFeatureTypeInt64: The type for integer features and feature values.
	MLFeatureTypeInt64 MLFeatureType = 1
	// MLFeatureTypeInvalid: The type for invalid feature values.
	MLFeatureTypeInvalid MLFeatureType = 0
	// MLFeatureTypeMultiArray: The type for multidimensional array features and feature values.
	MLFeatureTypeMultiArray MLFeatureType = 5
	// MLFeatureTypeSequence: The type for sequence features and feature values.
	MLFeatureTypeSequence MLFeatureType = 7
	// MLFeatureTypeState: MLState.
	MLFeatureTypeState MLFeatureType = 8
	// MLFeatureTypeString: The type for string features and feature values.
	MLFeatureTypeString MLFeatureType = 3
)

func (e MLFeatureType) String() string {
	switch e {
	case MLFeatureTypeDictionary:
		return "MLFeatureTypeDictionary"
	case MLFeatureTypeDouble:
		return "MLFeatureTypeDouble"
	case MLFeatureTypeImage:
		return "MLFeatureTypeImage"
	case MLFeatureTypeInt64:
		return "MLFeatureTypeInt64"
	case MLFeatureTypeInvalid:
		return "MLFeatureTypeInvalid"
	case MLFeatureTypeMultiArray:
		return "MLFeatureTypeMultiArray"
	case MLFeatureTypeSequence:
		return "MLFeatureTypeSequence"
	case MLFeatureTypeState:
		return "MLFeatureTypeState"
	case MLFeatureTypeString:
		return "MLFeatureTypeString"
	default:
		return fmt.Sprintf("MLFeatureType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraintType
type MLImageSizeConstraintType int

const (
	// MLImageSizeConstraintTypeEnumerated: The image feature accepts image sizes listed in an array.
	MLImageSizeConstraintTypeEnumerated MLImageSizeConstraintType = 2
	// MLImageSizeConstraintTypeRange: The image feature accepts image sizes defined by a range of widths and a range of heights.
	MLImageSizeConstraintTypeRange MLImageSizeConstraintType = 3
	// MLImageSizeConstraintTypeUnspecified: The image size constraint is not configured and should be ignored.
	MLImageSizeConstraintTypeUnspecified MLImageSizeConstraintType = 0
)

func (e MLImageSizeConstraintType) String() string {
	switch e {
	case MLImageSizeConstraintTypeEnumerated:
		return "MLImageSizeConstraintTypeEnumerated"
	case MLImageSizeConstraintTypeRange:
		return "MLImageSizeConstraintTypeRange"
	case MLImageSizeConstraintTypeUnspecified:
		return "MLImageSizeConstraintTypeUnspecified"
	default:
		return fmt.Sprintf("MLImageSizeConstraintType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreML/MLModelError-swift.struct/Code
type MLModelError int

const (
	// MLModelErrorCustomLayer: An error code for problems related to custom layers.
	MLModelErrorCustomLayer MLModelError = 4
	// MLModelErrorCustomModel: An error code for problems related to custom models.
	MLModelErrorCustomModel MLModelError = 5
	// MLModelErrorFeatureType: An error code for problems related to model features.
	MLModelErrorFeatureType MLModelError = 1
	// MLModelErrorGeneric: An error code for runtime issues that don’t apply to the other error codes.
	MLModelErrorGeneric MLModelError = 0
	// MLModelErrorIO: An error code for problems related to the system’s input or output.
	MLModelErrorIO MLModelError = 3
	// MLModelErrorModelCollection: An error code for problems related to retrieving a model collection from the deployment system.
	MLModelErrorModelCollection MLModelError = 10
	// MLModelErrorModelDecryption: An error code for problems related to decrypting models.
	MLModelErrorModelDecryption MLModelError = 9
	// MLModelErrorModelDecryptionKeyFetch: An error code for problems related to retrieving a model’s decryption key.
	MLModelErrorModelDecryptionKeyFetch MLModelError = 8
	// MLModelErrorParameters: An error code for problems related to model parameters.
	MLModelErrorParameters MLModelError = 7
	// MLModelErrorPredictionCancelled: An error code for problems related to canceling the prediction before it completes.
	MLModelErrorPredictionCancelled MLModelError = 11
	// MLModelErrorUpdate: An error code for problems related to on-device model updates.
	MLModelErrorUpdate MLModelError = 6
)

func (e MLModelError) String() string {
	switch e {
	case MLModelErrorCustomLayer:
		return "MLModelErrorCustomLayer"
	case MLModelErrorCustomModel:
		return "MLModelErrorCustomModel"
	case MLModelErrorFeatureType:
		return "MLModelErrorFeatureType"
	case MLModelErrorGeneric:
		return "MLModelErrorGeneric"
	case MLModelErrorIO:
		return "MLModelErrorIO"
	case MLModelErrorModelCollection:
		return "MLModelErrorModelCollection"
	case MLModelErrorModelDecryption:
		return "MLModelErrorModelDecryption"
	case MLModelErrorModelDecryptionKeyFetch:
		return "MLModelErrorModelDecryptionKeyFetch"
	case MLModelErrorParameters:
		return "MLModelErrorParameters"
	case MLModelErrorPredictionCancelled:
		return "MLModelErrorPredictionCancelled"
	case MLModelErrorUpdate:
		return "MLModelErrorUpdate"
	default:
		return fmt.Sprintf("MLModelError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayDataType
type MLMultiArrayDataType int

const (
	// MLMultiArrayDataTypeDouble: Designates the multiarray’s elements as doubles.
	MLMultiArrayDataTypeDouble MLMultiArrayDataType = 65536
	// MLMultiArrayDataTypeFloat16: Designates the multiarray’s elements as 16-bit floats.
	MLMultiArrayDataTypeFloat16 MLMultiArrayDataType = 65536
	// MLMultiArrayDataTypeFloat32: Designates the multiarray’s elements as 32-bit floats.
	MLMultiArrayDataTypeFloat32 MLMultiArrayDataType = 65536
	// MLMultiArrayDataTypeInt32: Designates the multiarray’s elements as 32-bit integers.
	MLMultiArrayDataTypeInt32 MLMultiArrayDataType = 131072
)

func (e MLMultiArrayDataType) String() string {
	switch e {
	case MLMultiArrayDataTypeDouble:
		return "MLMultiArrayDataTypeDouble"
	case MLMultiArrayDataTypeInt32:
		return "MLMultiArrayDataTypeInt32"
	default:
		return fmt.Sprintf("MLMultiArrayDataType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayShapeConstraintType
type MLMultiArrayShapeConstraintType int

const (
	// MLMultiArrayShapeConstraintTypeEnumerated: The constraint is an array of allowed shapes.
	MLMultiArrayShapeConstraintTypeEnumerated MLMultiArrayShapeConstraintType = 2
	// MLMultiArrayShapeConstraintTypeRange: The constraint is a set of ranges allowed for the array shape.
	MLMultiArrayShapeConstraintTypeRange MLMultiArrayShapeConstraintType = 3
	// MLMultiArrayShapeConstraintTypeUnspecified: The constraint type is undefined.
	MLMultiArrayShapeConstraintTypeUnspecified MLMultiArrayShapeConstraintType = 1
)

func (e MLMultiArrayShapeConstraintType) String() string {
	switch e {
	case MLMultiArrayShapeConstraintTypeEnumerated:
		return "MLMultiArrayShapeConstraintTypeEnumerated"
	case MLMultiArrayShapeConstraintTypeRange:
		return "MLMultiArrayShapeConstraintTypeRange"
	case MLMultiArrayShapeConstraintTypeUnspecified:
		return "MLMultiArrayShapeConstraintTypeUnspecified"
	default:
		return fmt.Sprintf("MLMultiArrayShapeConstraintType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreML/MLReshapeFrequencyHint
type MLReshapeFrequencyHint int

const (
	MLReshapeFrequencyHintFrequent MLReshapeFrequencyHint = 0
	MLReshapeFrequencyHintInfrequent MLReshapeFrequencyHint = 1
)

func (e MLReshapeFrequencyHint) String() string {
	switch e {
	case MLReshapeFrequencyHintFrequent:
		return "MLReshapeFrequencyHintFrequent"
	case MLReshapeFrequencyHintInfrequent:
		return "MLReshapeFrequencyHintInfrequent"
	default:
		return fmt.Sprintf("MLReshapeFrequencyHint(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreML/MLSpecializationStrategy
type MLSpecializationStrategy int

const (
	MLSpecializationStrategyDefault MLSpecializationStrategy = 0
	MLSpecializationStrategyFastPrediction MLSpecializationStrategy = 1
)

func (e MLSpecializationStrategy) String() string {
	switch e {
	case MLSpecializationStrategyDefault:
		return "MLSpecializationStrategyDefault"
	case MLSpecializationStrategyFastPrediction:
		return "MLSpecializationStrategyFastPrediction"
	default:
		return fmt.Sprintf("MLSpecializationStrategy(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreML/MLTaskState
type MLTaskState int

const (
	// MLTaskStateCancelling: The state of a machine learning task that’s in mid-termination, before it could finish successfully.
	MLTaskStateCancelling MLTaskState = 3
	// MLTaskStateCompleted: The state of a machine learning task that has finished successfully.
	MLTaskStateCompleted MLTaskState = 4
	// MLTaskStateFailed: The state of a machine learning task that has terminated due to an error.
	MLTaskStateFailed MLTaskState = 5
	// MLTaskStateRunning: The state of a machine learning task that’s executing.
	MLTaskStateRunning MLTaskState = 2
	// MLTaskStateSuspended: The state of a machine learning task that’s paused.
	MLTaskStateSuspended MLTaskState = 1
)

func (e MLTaskState) String() string {
	switch e {
	case MLTaskStateCancelling:
		return "MLTaskStateCancelling"
	case MLTaskStateCompleted:
		return "MLTaskStateCompleted"
	case MLTaskStateFailed:
		return "MLTaskStateFailed"
	case MLTaskStateRunning:
		return "MLTaskStateRunning"
	case MLTaskStateSuspended:
		return "MLTaskStateSuspended"
	default:
		return fmt.Sprintf("MLTaskState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreML/MLUpdateProgressEvent
type MLUpdateProgressEvent int

const (
	// MLUpdateProgressEventEpochEnd: An event that represents the end of training epoch.
	MLUpdateProgressEventEpochEnd MLUpdateProgressEvent = 2
	// MLUpdateProgressEventMiniBatchEnd: An event that represents the end of a mini-batch within a training epoch.
	MLUpdateProgressEventMiniBatchEnd MLUpdateProgressEvent = 4
	// MLUpdateProgressEventTrainingBegin: An event that represents the start of training.
	MLUpdateProgressEventTrainingBegin MLUpdateProgressEvent = 1
)

func (e MLUpdateProgressEvent) String() string {
	switch e {
	case MLUpdateProgressEventEpochEnd:
		return "MLUpdateProgressEventEpochEnd"
	case MLUpdateProgressEventMiniBatchEnd:
		return "MLUpdateProgressEventMiniBatchEnd"
	case MLUpdateProgressEventTrainingBegin:
		return "MLUpdateProgressEventTrainingBegin"
	default:
		return fmt.Sprintf("MLUpdateProgressEvent(%d)", e)
	}
}

