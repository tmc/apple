// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVVideoOutputSpecification] class.
var (
	_AVVideoOutputSpecificationClass     AVVideoOutputSpecificationClass
	_AVVideoOutputSpecificationClassOnce sync.Once
)

func getAVVideoOutputSpecificationClass() AVVideoOutputSpecificationClass {
	_AVVideoOutputSpecificationClassOnce.Do(func() {
		_AVVideoOutputSpecificationClass = AVVideoOutputSpecificationClass{class: objc.GetClass("AVVideoOutputSpecification")}
	})
	return _AVVideoOutputSpecificationClass
}

// GetAVVideoOutputSpecificationClass returns the class object for AVVideoOutputSpecification.
func GetAVVideoOutputSpecificationClass() AVVideoOutputSpecificationClass {
	return getAVVideoOutputSpecificationClass()
}

type AVVideoOutputSpecificationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVVideoOutputSpecificationClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVVideoOutputSpecificationClass) Alloc() AVVideoOutputSpecification {
	rv := objc.Send[AVVideoOutputSpecification](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that specifies the pixel buffer attributes and tag collections
// handled by a player video output.
//
// # Configuring the specification
//
//   - [AVVideoOutputSpecification.DefaultOutputSettings]
//   - [AVVideoOutputSpecification.SetDefaultOutputSettings]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoOutputSpecification
type AVVideoOutputSpecification struct {
	objectivec.Object
}

// AVVideoOutputSpecificationFromID constructs a [AVVideoOutputSpecification] from an objc.ID.
//
// An object that specifies the pixel buffer attributes and tag collections
// handled by a player video output.
func AVVideoOutputSpecificationFromID(id objc.ID) AVVideoOutputSpecification {
	return AVVideoOutputSpecification{objectivec.Object{ID: id}}
}
// NOTE: AVVideoOutputSpecification adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVVideoOutputSpecification] class.
//
// # Configuring the specification
//
//   - [IAVVideoOutputSpecification.DefaultOutputSettings]
//   - [IAVVideoOutputSpecification.SetDefaultOutputSettings]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoOutputSpecification
type IAVVideoOutputSpecification interface {
	objectivec.IObject

	// Topic: Configuring the specification

	DefaultOutputSettings() foundation.INSDictionary
	SetDefaultOutputSettings(value foundation.INSDictionary)

	PreferredTagCollections() foundation.INSArray
	InitWithTagCollections(tagCollections foundation.INSArray) AVVideoOutputSpecification
	SetOutputSettingsForTagCollection(outputSettings foundation.INSDictionary, tagCollection coremedia.CMTagCollectionRef)
}

// Init initializes the instance.
func (v AVVideoOutputSpecification) Init() AVVideoOutputSpecification {
	rv := objc.Send[AVVideoOutputSpecification](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v AVVideoOutputSpecification) Autorelease() AVVideoOutputSpecification {
	rv := objc.Send[AVVideoOutputSpecification](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVideoOutputSpecification creates a new AVVideoOutputSpecification instance.
func NewAVVideoOutputSpecification() AVVideoOutputSpecification {
	class := getAVVideoOutputSpecificationClass()
	rv := objc.Send[AVVideoOutputSpecification](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoOutputSpecification/initWithTagCollections:
func NewVideoOutputSpecificationWithTagCollections(tagCollections foundation.INSArray) AVVideoOutputSpecification {
	instance := getAVVideoOutputSpecificationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTagCollections:"), tagCollections)
	return AVVideoOutputSpecificationFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoOutputSpecification/initWithTagCollections:
func (v AVVideoOutputSpecification) InitWithTagCollections(tagCollections foundation.INSArray) AVVideoOutputSpecification {
	rv := objc.Send[AVVideoOutputSpecification](v.ID, objc.Sel("initWithTagCollections:"), tagCollections)
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoOutputSpecification/setOutputSettings:forTagCollection:
func (v AVVideoOutputSpecification) SetOutputSettingsForTagCollection(outputSettings foundation.INSDictionary, tagCollection coremedia.CMTagCollectionRef) {
	objc.Send[objc.ID](v.ID, objc.Sel("setOutputSettings:forTagCollection:"), outputSettings, tagCollection)
}

// See: https://developer.apple.com/documentation/AVFoundation/AVVideoOutputSpecification/defaultOutputSettings
func (v AVVideoOutputSpecification) DefaultOutputSettings() foundation.INSDictionary {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("defaultOutputSettings"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (v AVVideoOutputSpecification) SetDefaultOutputSettings(value foundation.INSDictionary) {
	objc.Send[struct{}](v.ID, objc.Sel("setDefaultOutputSettings:"), value)
}
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoOutputSpecification/preferredTagCollections-2ikbd
func (v AVVideoOutputSpecification) PreferredTagCollections() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("preferredTagCollections"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

