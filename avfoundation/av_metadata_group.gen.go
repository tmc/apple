// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVMetadataGroup] class.
var (
	_AVMetadataGroupClass     AVMetadataGroupClass
	_AVMetadataGroupClassOnce sync.Once
)

func getAVMetadataGroupClass() AVMetadataGroupClass {
	_AVMetadataGroupClassOnce.Do(func() {
		_AVMetadataGroupClass = AVMetadataGroupClass{class: objc.GetClass("AVMetadataGroup")}
	})
	return _AVMetadataGroupClass
}

// GetAVMetadataGroupClass returns the class object for AVMetadataGroup.
func GetAVMetadataGroupClass() AVMetadataGroupClass {
	return getAVMetadataGroupClass()
}

type AVMetadataGroupClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMetadataGroupClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMetadataGroupClass) Alloc() AVMetadataGroup {
	rv := objc.Send[AVMetadataGroup](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A collection of metadata items associated with a timeline segment.
//
// # Inspecting the metadata group
//
//   - [AVMetadataGroup.Items]: The array of metadata items associated with the metadata group.
//   - [AVMetadataGroup.UniqueID]: The unique identifier for the metadata group.
//   - [AVMetadataGroup.ClassifyingLabel]: The classifying label associated with the metadata group.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataGroup
type AVMetadataGroup struct {
	objectivec.Object
}

// AVMetadataGroupFromID constructs a [AVMetadataGroup] from an objc.ID.
//
// A collection of metadata items associated with a timeline segment.
func AVMetadataGroupFromID(id objc.ID) AVMetadataGroup {
	return AVMetadataGroup{objectivec.Object{ID: id}}
}
// NOTE: AVMetadataGroup adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMetadataGroup] class.
//
// # Inspecting the metadata group
//
//   - [IAVMetadataGroup.Items]: The array of metadata items associated with the metadata group.
//   - [IAVMetadataGroup.UniqueID]: The unique identifier for the metadata group.
//   - [IAVMetadataGroup.ClassifyingLabel]: The classifying label associated with the metadata group.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataGroup
type IAVMetadataGroup interface {
	objectivec.IObject

	// Topic: Inspecting the metadata group

	// The array of metadata items associated with the metadata group.
	Items() []AVMetadataItem
	// The unique identifier for the metadata group.
	UniqueID() string
	// The classifying label associated with the metadata group.
	ClassifyingLabel() string
}

// Init initializes the instance.
func (m AVMetadataGroup) Init() AVMetadataGroup {
	rv := objc.Send[AVMetadataGroup](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMetadataGroup) Autorelease() AVMetadataGroup {
	rv := objc.Send[AVMetadataGroup](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMetadataGroup creates a new AVMetadataGroup instance.
func NewAVMetadataGroup() AVMetadataGroup {
	class := getAVMetadataGroupClass()
	rv := objc.Send[AVMetadataGroup](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The array of metadata items associated with the metadata group.
//
// # Discussion
// 
// The `items` array may be empty if no metadata items are associated with
// this group.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataGroup/items
func (m AVMetadataGroup) Items() []AVMetadataItem {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("items"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVMetadataItem {
		return AVMetadataItemFromID(id)
	})
}
// The unique identifier for the metadata group.
//
// # Discussion
// 
// The value of this property may be `nil` if no unique identifier is defined
// for this group.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataGroup/uniqueID
func (m AVMetadataGroup) UniqueID() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("uniqueID"))
	return foundation.NSStringFromID(rv).String()
}
// The classifying label associated with the metadata group.
//
// # Discussion
// 
// The value of this property may be `nil` if no classifying label is defined
// for this group.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataGroup/classifyingLabel
func (m AVMetadataGroup) ClassifyingLabel() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("classifyingLabel"))
	return foundation.NSStringFromID(rv).String()
}

