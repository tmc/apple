// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVMetadataItemFilter] class.
var (
	_AVMetadataItemFilterClass     AVMetadataItemFilterClass
	_AVMetadataItemFilterClassOnce sync.Once
)

func getAVMetadataItemFilterClass() AVMetadataItemFilterClass {
	_AVMetadataItemFilterClassOnce.Do(func() {
		_AVMetadataItemFilterClass = AVMetadataItemFilterClass{class: objc.GetClass("AVMetadataItemFilter")}
	})
	return _AVMetadataItemFilterClass
}

// GetAVMetadataItemFilterClass returns the class object for AVMetadataItemFilter.
func GetAVMetadataItemFilterClass() AVMetadataItemFilterClass {
	return getAVMetadataItemFilterClass()
}

type AVMetadataItemFilterClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMetadataItemFilterClass) Alloc() AVMetadataItemFilter {
	rv := objc.Send[AVMetadataItemFilter](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that filters selected information from a metadata item.
//
// # Overview
// 
// Filter instances are opaque, unmodifiable objects, that you create with the
// [AVMetadataItemFilter.MetadataItemFilterForSharing] class method.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataItemFilter
type AVMetadataItemFilter struct {
	objectivec.Object
}

// AVMetadataItemFilterFromID constructs a [AVMetadataItemFilter] from an objc.ID.
//
// An object that filters selected information from a metadata item.
func AVMetadataItemFilterFromID(id objc.ID) AVMetadataItemFilter {
	return AVMetadataItemFilter{objectivec.Object{ID: id}}
}
// NOTE: AVMetadataItemFilter adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMetadataItemFilter] class.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataItemFilter
type IAVMetadataItemFilter interface {
	objectivec.IObject
}

// Init initializes the instance.
func (m AVMetadataItemFilter) Init() AVMetadataItemFilter {
	rv := objc.Send[AVMetadataItemFilter](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMetadataItemFilter) Autorelease() AVMetadataItemFilter {
	rv := objc.Send[AVMetadataItemFilter](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMetadataItemFilter creates a new AVMetadataItemFilter instance.
func NewAVMetadataItemFilter() AVMetadataItemFilter {
	class := getAVMetadataItemFilterClass()
	rv := objc.Send[AVMetadataItemFilter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a metadata filter to use for sharing assets.
//
// # Return Value
// 
// An instance of an [AVMetadataItemFilter].
//
// # Discussion
// 
// Removes user-identifying metadata items, such as location information, and
// leaves only metadata related to commerce or playback itself. For example,
// playback, copyright, and commercial-related metadata, such as a
// purchaser’s ID as set by a vendor of digital media, along with metadata
// either derivable from the media itself or necessary for its proper behavior
// are all left intact.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataItemFilter/forSharing()
func (_AVMetadataItemFilterClass AVMetadataItemFilterClass) MetadataItemFilterForSharing() AVMetadataItemFilter {
	rv := objc.Send[objc.ID](objc.ID(_AVMetadataItemFilterClass.class), objc.Sel("metadataItemFilterForSharing"))
	return AVMetadataItemFilterFromID(rv)
}

