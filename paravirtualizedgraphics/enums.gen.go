// Code generated from Apple documentation for ParavirtualizedGraphics. DO NOT EDIT.

package paravirtualizedgraphics

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGResumeErrorCode
type PGResumeErrorCode uint

const (
	// PGResumeErrorCodeIncompatibleDevice: The resume device is missing capabilities that the suspended device provided.
	PGResumeErrorCodeIncompatibleDevice PGResumeErrorCode = 4
	// PGResumeErrorCodeInternalFault: An internal error occurred.
	PGResumeErrorCodeInternalFault PGResumeErrorCode = 0
	// PGResumeErrorCodeInvalidContent: The content of the suspend state or the guest memory isn’t valid.
	PGResumeErrorCodeInvalidContent          PGResumeErrorCode = 2
	PGResumeErrorCodeInvalidDisplayPortCount PGResumeErrorCode = 5
	// PGResumeErrorCodeInvalidGuestVersion: The guest version is incompatible with this framework version.
	PGResumeErrorCodeInvalidGuestVersion PGResumeErrorCode = 3
	// PGResumeErrorCodeInvalidSuspendStateVersion: The suspend state version is incompatible with this framework version.
	PGResumeErrorCodeInvalidSuspendStateVersion PGResumeErrorCode = 1
)

func (e PGResumeErrorCode) String() string {
	switch e {
	case PGResumeErrorCodeIncompatibleDevice:
		return "PGResumeErrorCodeIncompatibleDevice"
	case PGResumeErrorCodeInternalFault:
		return "PGResumeErrorCodeInternalFault"
	case PGResumeErrorCodeInvalidContent:
		return "PGResumeErrorCodeInvalidContent"
	case PGResumeErrorCodeInvalidDisplayPortCount:
		return "PGResumeErrorCodeInvalidDisplayPortCount"
	case PGResumeErrorCodeInvalidGuestVersion:
		return "PGResumeErrorCodeInvalidGuestVersion"
	case PGResumeErrorCodeInvalidSuspendStateVersion:
		return "PGResumeErrorCodeInvalidSuspendStateVersion"
	default:
		return fmt.Sprintf("PGResumeErrorCode(%d)", e)
	}
}
