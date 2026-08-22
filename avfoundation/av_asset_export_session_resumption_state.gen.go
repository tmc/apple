// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAssetExportSessionResumptionState] class.
var (
	_AVAssetExportSessionResumptionStateClass     AVAssetExportSessionResumptionStateClass
	_AVAssetExportSessionResumptionStateClassOnce sync.Once
)

func getAVAssetExportSessionResumptionStateClass() AVAssetExportSessionResumptionStateClass {
	_AVAssetExportSessionResumptionStateClassOnce.Do(func() {
		_AVAssetExportSessionResumptionStateClass = AVAssetExportSessionResumptionStateClass{class: objc.GetClass("AVAssetExportSessionResumptionState")}
	})
	return _AVAssetExportSessionResumptionStateClass
}

// GetAVAssetExportSessionResumptionStateClass returns the class object for AVAssetExportSessionResumptionState.
func GetAVAssetExportSessionResumptionStateClass() AVAssetExportSessionResumptionStateClass {
	return getAVAssetExportSessionResumptionStateClass()
}

type AVAssetExportSessionResumptionStateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAssetExportSessionResumptionStateClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetExportSessionResumptionStateClass) Alloc() AVAssetExportSessionResumptionState {
	rv := objc.Send[AVAssetExportSessionResumptionState](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// The current resumption state of the export session.
//
// # Overview
//
// Configure a resumable export session with
// [configureForResumableExportWithCompletionHandler:].
//
// # Inspecting the resumption state
//
//   - [AVAssetExportSessionResumptionState.IsResumptionConfigured]: A Boolean value that indicates whether the export session is configured as resumable.
//   - [AVAssetExportSessionResumptionState.IsResumingFromPreviousState]: A Boolean value that indicates whether or not a resuming export is continuing from a previous state.
//   - [AVAssetExportSessionResumptionState.ConfigurationFailureReason]: The reason that the export session couldn’t be configured as resumable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSessionResumptionState
//
// [configureForResumableExportWithCompletionHandler:]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/configureForResumableExportWithCompletionHandler:
type AVAssetExportSessionResumptionState struct {
	objectivec.Object
}

// AVAssetExportSessionResumptionStateFromID constructs a [AVAssetExportSessionResumptionState] from an objc.ID.
//
// The current resumption state of the export session.
func AVAssetExportSessionResumptionStateFromID(id objc.ID) AVAssetExportSessionResumptionState {
	return AVAssetExportSessionResumptionState{objectivec.Object{ID: id}}
}

// Ensure AVAssetExportSessionResumptionState implements IAVAssetExportSessionResumptionState.
var _ IAVAssetExportSessionResumptionState = AVAssetExportSessionResumptionState{}

// An interface definition for the [AVAssetExportSessionResumptionState] class.
//
// # Inspecting the resumption state
//
//   - [IAVAssetExportSessionResumptionState.IsResumptionConfigured]: A Boolean value that indicates whether the export session is configured as resumable.
//   - [IAVAssetExportSessionResumptionState.IsResumingFromPreviousState]: A Boolean value that indicates whether or not a resuming export is continuing from a previous state.
//   - [IAVAssetExportSessionResumptionState.ConfigurationFailureReason]: The reason that the export session couldn’t be configured as resumable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSessionResumptionState
type IAVAssetExportSessionResumptionState interface {
	objectivec.IObject

	// Topic: Inspecting the resumption state

	// A Boolean value that indicates whether the export session is configured as resumable.
	IsResumptionConfigured() bool
	// A Boolean value that indicates whether or not a resuming export is continuing from a previous state.
	IsResumingFromPreviousState() bool
	// The reason that the export session couldn’t be configured as resumable.
	ConfigurationFailureReason() objectivec.IObject
}

// Init initializes the instance.
func (a AVAssetExportSessionResumptionState) Init() AVAssetExportSessionResumptionState {
	rv := objc.Send[AVAssetExportSessionResumptionState](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetExportSessionResumptionState) Autorelease() AVAssetExportSessionResumptionState {
	rv := objc.Send[AVAssetExportSessionResumptionState](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetExportSessionResumptionState creates a new AVAssetExportSessionResumptionState instance.
func NewAVAssetExportSessionResumptionState() AVAssetExportSessionResumptionState {
	class := getAVAssetExportSessionResumptionStateClass()
	rv := objc.Send[AVAssetExportSessionResumptionState](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Boolean value that indicates whether the export session is configured as
// resumable.
//
// # Discussion
//
// If `true`, the export session is configured as resumable. If `false`, the
// export session will remain as non-resumable (default). You can still call
// [AVAssetExportSession.ExportAsynchronouslyWithCompletionHandler] when this
// property is `false`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSessionResumptionState/resumptionConfigured
func (a AVAssetExportSessionResumptionState) IsResumptionConfigured() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isResumptionConfigured"))
	return rv
}

// A Boolean value that indicates whether or not a resuming export is
// continuing from a previous state.
//
// # Discussion
//
// A value of `true` means the export resumes from previous results; a value
// of `false` means it starts from the beginning. This value is valid only
// when [AVAssetExportSessionResumptionState.ResumptionConfigured] is `true`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSessionResumptionState/resumingFromPreviousState
func (a AVAssetExportSessionResumptionState) IsResumingFromPreviousState() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isResumingFromPreviousState"))
	return rv
}

// The reason that the export session couldn’t be configured as resumable.
//
// # Discussion
//
// This value is valid only when
// [AVAssetExportSessionResumptionState.ResumptionConfigured] is `false`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSessionResumptionState/configurationFailureReason
func (a AVAssetExportSessionResumptionState) ConfigurationFailureReason() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("configurationFailureReason"))
	return objectivec.Object{ID: rv}
}
