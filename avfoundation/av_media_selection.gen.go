// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVMediaSelection] class.
var (
	_AVMediaSelectionClass     AVMediaSelectionClass
	_AVMediaSelectionClassOnce sync.Once
)

func getAVMediaSelectionClass() AVMediaSelectionClass {
	_AVMediaSelectionClassOnce.Do(func() {
		_AVMediaSelectionClass = AVMediaSelectionClass{class: objc.GetClass("AVMediaSelection")}
	})
	return _AVMediaSelectionClass
}

// GetAVMediaSelectionClass returns the class object for AVMediaSelection.
func GetAVMediaSelectionClass() AVMediaSelectionClass {
	return getAVMediaSelectionClass()
}

type AVMediaSelectionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMediaSelectionClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMediaSelectionClass) Alloc() AVMediaSelection {
	rv := objc.Send[AVMediaSelection](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a complete rendition of media selection options
// on an asset.
//
// # Inspecting the media selection
//
//   - [AVMediaSelection.SelectedMediaOptionInMediaSelectionGroup]: Returns the media selection option that’s currently selected in the specified group.
//   - [AVMediaSelection.MediaSelectionCriteriaCanBeAppliedAutomaticallyToMediaSelectionGroup]: Indicates whether the specified media selection group is subject to automatic media selection.
//
// # Accessing the asset
//
//   - [AVMediaSelection.Asset]: The asset associated with the media selection.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaSelection
type AVMediaSelection struct {
	objectivec.Object
}

// AVMediaSelectionFromID constructs a [AVMediaSelection] from an objc.ID.
//
// An object that represents a complete rendition of media selection options
// on an asset.
func AVMediaSelectionFromID(id objc.ID) AVMediaSelection {
	return AVMediaSelection{objectivec.Object{ID: id}}
}
// NOTE: AVMediaSelection adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMediaSelection] class.
//
// # Inspecting the media selection
//
//   - [IAVMediaSelection.SelectedMediaOptionInMediaSelectionGroup]: Returns the media selection option that’s currently selected in the specified group.
//   - [IAVMediaSelection.MediaSelectionCriteriaCanBeAppliedAutomaticallyToMediaSelectionGroup]: Indicates whether the specified media selection group is subject to automatic media selection.
//
// # Accessing the asset
//
//   - [IAVMediaSelection.Asset]: The asset associated with the media selection.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaSelection
type IAVMediaSelection interface {
	objectivec.IObject

	// Topic: Inspecting the media selection

	// Returns the media selection option that’s currently selected in the specified group.
	SelectedMediaOptionInMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) IAVMediaSelectionOption
	// Indicates whether the specified media selection group is subject to automatic media selection.
	MediaSelectionCriteriaCanBeAppliedAutomaticallyToMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) bool

	// Topic: Accessing the asset

	// The asset associated with the media selection.
	Asset() IAVAsset
}

// Init initializes the instance.
func (m AVMediaSelection) Init() AVMediaSelection {
	rv := objc.Send[AVMediaSelection](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMediaSelection) Autorelease() AVMediaSelection {
	rv := objc.Send[AVMediaSelection](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMediaSelection creates a new AVMediaSelection instance.
func NewAVMediaSelection() AVMediaSelection {
	class := getAVMediaSelectionClass()
	rv := objc.Send[AVMediaSelection](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the media selection option that’s currently selected in the
// specified group.
//
// mediaSelectionGroup: A media selection group obtained from the associated asset.
//
// # Return Value
// 
// The currently selected [AVMediaSelectionOption]. The return value may be
// `nil`.
//
// # Discussion
// 
// This method returns the currently selected [AVMediaSelectionOption] in the
// specified [AVMediaSelectionGroup], but may return `nil` if media selection
// group’s [AllowsEmptySelection] is set to [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaSelection/selectedMediaOption(in:)
func (m AVMediaSelection) SelectedMediaOptionInMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) IAVMediaSelectionOption {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("selectedMediaOptionInMediaSelectionGroup:"), mediaSelectionGroup)
	return AVMediaSelectionOptionFromID(rv)
}
// Indicates whether the specified media selection group is subject to
// automatic media selection.
//
// mediaSelectionGroup: A media selection group obtained from the associated asset.
//
// # Return Value
// 
// A Boolean value indicating whether the group is subject to automatic media
// selection.
//
// # Discussion
// 
// The automatic application of media selection criteria is suspended in any
// group in which a specific selection has been made by calling
// [SelectMediaOptionInMediaSelectionGroup] on the current [AVPlayerItem].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaSelection/mediaSelectionCriteriaCanBeAppliedAutomatically(to:)
func (m AVMediaSelection) MediaSelectionCriteriaCanBeAppliedAutomaticallyToMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("mediaSelectionCriteriaCanBeAppliedAutomaticallyToMediaSelectionGroup:"), mediaSelectionGroup)
	return rv
}

// The asset associated with the media selection.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaSelection/asset
func (m AVMediaSelection) Asset() IAVAsset {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("asset"))
	return AVAssetFromID(objc.ID(rv))
}

