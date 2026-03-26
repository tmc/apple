// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCaptionGrouper] class.
var (
	_AVCaptionGrouperClass     AVCaptionGrouperClass
	_AVCaptionGrouperClassOnce sync.Once
)

func getAVCaptionGrouperClass() AVCaptionGrouperClass {
	_AVCaptionGrouperClassOnce.Do(func() {
		_AVCaptionGrouperClass = AVCaptionGrouperClass{class: objc.GetClass("AVCaptionGrouper")}
	})
	return _AVCaptionGrouperClass
}

// GetAVCaptionGrouperClass returns the class object for AVCaptionGrouper.
func GetAVCaptionGrouperClass() AVCaptionGrouperClass {
	return getAVCaptionGrouperClass()
}

type AVCaptionGrouperClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptionGrouperClass) Alloc() AVCaptionGrouper {
	rv := objc.Send[AVCaptionGrouper](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that analyzes the temporal overlaps of caption objects to create
// caption groups for each span of concurrent captions.
//
// # Adding captions
//
//   - [AVCaptionGrouper.AddCaption]: Adds a caption to the pending group.
//
// # Generating captions groups
//
//   - [AVCaptionGrouper.FlushAddedCaptionsIntoGroupsUpToTime]: Creates caption groups for the captions you enqueue up to the time.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionGrouper
type AVCaptionGrouper struct {
	objectivec.Object
}

// AVCaptionGrouperFromID constructs a [AVCaptionGrouper] from an objc.ID.
//
// An object that analyzes the temporal overlaps of caption objects to create
// caption groups for each span of concurrent captions.
func AVCaptionGrouperFromID(id objc.ID) AVCaptionGrouper {
	return AVCaptionGrouper{objectivec.Object{ID: id}}
}
// NOTE: AVCaptionGrouper adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptionGrouper] class.
//
// # Adding captions
//
//   - [IAVCaptionGrouper.AddCaption]: Adds a caption to the pending group.
//
// # Generating captions groups
//
//   - [IAVCaptionGrouper.FlushAddedCaptionsIntoGroupsUpToTime]: Creates caption groups for the captions you enqueue up to the time.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionGrouper
type IAVCaptionGrouper interface {
	objectivec.IObject

	// Topic: Adding captions

	// Adds a caption to the pending group.
	AddCaption(input IAVCaption)

	// Topic: Generating captions groups

	// Creates caption groups for the captions you enqueue up to the time.
	FlushAddedCaptionsIntoGroupsUpToTime(upToTime objectivec.IObject) []AVCaptionGroup
}

// Init initializes the instance.
func (c AVCaptionGrouper) Init() AVCaptionGrouper {
	rv := objc.Send[AVCaptionGrouper](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptionGrouper) Autorelease() AVCaptionGrouper {
	rv := objc.Send[AVCaptionGrouper](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptionGrouper creates a new AVCaptionGrouper instance.
func NewAVCaptionGrouper() AVCaptionGrouper {
	class := getAVCaptionGrouperClass()
	rv := objc.Send[AVCaptionGrouper](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Adds a caption to the pending group.
//
// input: The caption to add.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionGrouper/add(_:)
func (c AVCaptionGrouper) AddCaption(input IAVCaption) {
	objc.Send[objc.ID](c.ID, objc.Sel("addCaption:"), input)
}
// Creates caption groups for the captions you enqueue up to the time.
//
// upToTime: The time up to which the system flushes the queue.
//
// upToTime is a [coremedia.CMTime].
//
// # Return Value
// 
// An array of zero or more caption groups.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionGrouper/flushAddedCaptions(upTo:)
func (c AVCaptionGrouper) FlushAddedCaptionsIntoGroupsUpToTime(upToTime objectivec.IObject) []AVCaptionGroup {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("flushAddedCaptionsIntoGroupsUpToTime:"), upToTime)
	return objc.ConvertSlice(rv, func(id objc.ID) AVCaptionGroup {
		return AVCaptionGroupFromID(id)
	})
}

