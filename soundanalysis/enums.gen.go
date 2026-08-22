// Code generated from Apple documentation for SoundAnalysis. DO NOT EDIT.

package soundanalysis

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/SoundAnalysis/SNError/Code
type SNErrorCode int

const (
	// SNErrorCodeInvalidFile: An error that indicates an audio file is invalid.
	SNErrorCodeInvalidFile SNErrorCode = 5
	// SNErrorCodeInvalidFormat: An error that indicates the audio data’s format isn’t valid.
	SNErrorCodeInvalidFormat SNErrorCode = 3
	// SNErrorCodeInvalidModel: An error that indicates the sound classifier’s underlying Core ML model is an invalid model type.
	SNErrorCodeInvalidModel SNErrorCode = 4
	// SNErrorCodeOperationFailed: An error that occurs when the framework fails to analyze audio.
	SNErrorCodeOperationFailed SNErrorCode = 2
	// SNErrorCodeUnknownError: An error that represents a failure that no other error handles.
	SNErrorCodeUnknownError SNErrorCode = 1
)

func (e SNErrorCode) String() string {
	switch e {
	case SNErrorCodeInvalidFile:
		return "SNErrorCodeInvalidFile"
	case SNErrorCodeInvalidFormat:
		return "SNErrorCodeInvalidFormat"
	case SNErrorCodeInvalidModel:
		return "SNErrorCodeInvalidModel"
	case SNErrorCodeOperationFailed:
		return "SNErrorCodeOperationFailed"
	case SNErrorCodeUnknownError:
		return "SNErrorCodeUnknownError"
	default:
		return fmt.Sprintf("SNErrorCode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/SoundAnalysis/SNTimeDurationConstraintType
type SNTimeDurationConstraintType int

const (
	// SNTimeDurationConstraintTypeEnumerated: A constraint type that uses an array of time durations to define what a request’s underlying sound classifier accepts.
	SNTimeDurationConstraintTypeEnumerated SNTimeDurationConstraintType = 1
	// SNTimeDurationConstraintTypeRange: A constraint type that uses a time duration range to define what a request’s underlying sound classifier accepts.
	SNTimeDurationConstraintTypeRange SNTimeDurationConstraintType = 2
)

func (e SNTimeDurationConstraintType) String() string {
	switch e {
	case SNTimeDurationConstraintTypeEnumerated:
		return "SNTimeDurationConstraintTypeEnumerated"
	case SNTimeDurationConstraintTypeRange:
		return "SNTimeDurationConstraintTypeRange"
	default:
		return fmt.Sprintf("SNTimeDurationConstraintType(%d)", e)
	}
}
