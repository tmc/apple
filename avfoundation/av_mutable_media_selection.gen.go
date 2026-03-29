// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVMutableMediaSelection] class.
var (
	_AVMutableMediaSelectionClass     AVMutableMediaSelectionClass
	_AVMutableMediaSelectionClassOnce sync.Once
)

func getAVMutableMediaSelectionClass() AVMutableMediaSelectionClass {
	_AVMutableMediaSelectionClassOnce.Do(func() {
		_AVMutableMediaSelectionClass = AVMutableMediaSelectionClass{class: objc.GetClass("AVMutableMediaSelection")}
	})
	return _AVMutableMediaSelectionClass
}

// GetAVMutableMediaSelectionClass returns the class object for AVMutableMediaSelection.
func GetAVMutableMediaSelectionClass() AVMutableMediaSelectionClass {
	return getAVMutableMediaSelectionClass()
}

type AVMutableMediaSelectionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMutableMediaSelectionClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMutableMediaSelectionClass) Alloc() AVMutableMediaSelection {
	rv := objc.Send[AVMutableMediaSelection](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A mutable object that represents a complete rendition of media selection
// options on an asset.
//
// # Selecting media options
//
//   - [AVMutableMediaSelection.SelectMediaOptionInMediaSelectionGroup]: Selects the media option in the specified media selection group.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMediaSelection
type AVMutableMediaSelection struct {
	AVMediaSelection
}

// AVMutableMediaSelectionFromID constructs a [AVMutableMediaSelection] from an objc.ID.
//
// A mutable object that represents a complete rendition of media selection
// options on an asset.
func AVMutableMediaSelectionFromID(id objc.ID) AVMutableMediaSelection {
	return AVMutableMediaSelection{AVMediaSelection: AVMediaSelectionFromID(id)}
}
// NOTE: AVMutableMediaSelection adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMutableMediaSelection] class.
//
// # Selecting media options
//
//   - [IAVMutableMediaSelection.SelectMediaOptionInMediaSelectionGroup]: Selects the media option in the specified media selection group.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMediaSelection
type IAVMutableMediaSelection interface {
	IAVMediaSelection

	// Topic: Selecting media options

	// Selects the media option in the specified media selection group.
	SelectMediaOptionInMediaSelectionGroup(mediaSelectionOption IAVMediaSelectionOption, mediaSelectionGroup IAVMediaSelectionGroup)
}

// Init initializes the instance.
func (m AVMutableMediaSelection) Init() AVMutableMediaSelection {
	rv := objc.Send[AVMutableMediaSelection](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMutableMediaSelection) Autorelease() AVMutableMediaSelection {
	rv := objc.Send[AVMutableMediaSelection](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMutableMediaSelection creates a new AVMutableMediaSelection instance.
func NewAVMutableMediaSelection() AVMutableMediaSelection {
	class := getAVMutableMediaSelectionClass()
	rv := objc.Send[AVMutableMediaSelection](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Selects the media option in the specified media selection group.
//
// mediaSelectionOption: The media selection option to select.
//
// mediaSelectionGroup: The media selection group containing the specified media selection option.
//
// # Discussion
// 
// This method selects the [AVMediaSelectionOption] in the specified
// [AVMediaSelectionGroup] and deselects all other options in that group. If
// the specified media selection option isn’t a member of the specified
// media selection group, no change in state will be made. If the media
// selection group’s [AllowsEmptySelection] property is set to [true], you
// can pass `nil` for `mediaSelectionOption` argument to deselect all media
// selection options in the group.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMediaSelection/select(_:in:)
func (m AVMutableMediaSelection) SelectMediaOptionInMediaSelectionGroup(mediaSelectionOption IAVMediaSelectionOption, mediaSelectionGroup IAVMediaSelectionGroup) {
	objc.Send[objc.ID](m.ID, objc.Sel("selectMediaOption:inMediaSelectionGroup:"), mediaSelectionOption, mediaSelectionGroup)
}

