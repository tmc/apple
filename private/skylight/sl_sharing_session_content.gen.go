// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSharingSessionContent] class.
var (
	_SLSharingSessionContentClass     SLSharingSessionContentClass
	_SLSharingSessionContentClassOnce sync.Once
)

func getSLSharingSessionContentClass() SLSharingSessionContentClass {
	_SLSharingSessionContentClassOnce.Do(func() {
		_SLSharingSessionContentClass = SLSharingSessionContentClass{class: objc.GetClass("SLSharingSessionContent")}
	})
	return _SLSharingSessionContentClass
}

// GetSLSharingSessionContentClass returns the class object for SLSharingSessionContent.
func GetSLSharingSessionContentClass() SLSharingSessionContentClass {
	return getSLSharingSessionContentClass()
}

type SLSharingSessionContentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSharingSessionContentClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSharingSessionContentClass) Alloc() SLSharingSessionContent {
	rv := objc.Send[SLSharingSessionContent](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSharingSessionContent.DisplayID]
//   - [SLSharingSessionContent.SetDisplayID]
//   - [SLSharingSessionContent.Filter]
//   - [SLSharingSessionContent.SetFilter]
//   - [SLSharingSessionContent.MetaData]
//   - [SLSharingSessionContent.SetMetaData]
//   - [SLSharingSessionContent.InitInternalDisplay]
//   - [SLSharingSessionContent.InitWithDisplayID]
//   - [SLSharingSessionContent.InitWithWindowID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSharingSessionContent
type SLSharingSessionContent struct {
	objectivec.Object
}

// SLSharingSessionContentFromID constructs a [SLSharingSessionContent] from an objc.ID.
func SLSharingSessionContentFromID(id objc.ID) SLSharingSessionContent {
	return SLSharingSessionContent{objectivec.Object{ID: id}}
}

// Ensure SLSharingSessionContent implements ISLSharingSessionContent.
var _ ISLSharingSessionContent = SLSharingSessionContent{}

// An interface definition for the [SLSharingSessionContent] class.
//
// # Methods
//
//   - [ISLSharingSessionContent.DisplayID]
//   - [ISLSharingSessionContent.SetDisplayID]
//   - [ISLSharingSessionContent.Filter]
//   - [ISLSharingSessionContent.SetFilter]
//   - [ISLSharingSessionContent.MetaData]
//   - [ISLSharingSessionContent.SetMetaData]
//   - [ISLSharingSessionContent.InitInternalDisplay]
//   - [ISLSharingSessionContent.InitWithDisplayID]
//   - [ISLSharingSessionContent.InitWithWindowID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSharingSessionContent
type ISLSharingSessionContent interface {
	objectivec.IObject

	// Topic: Methods

	DisplayID() foundation.NSNumber
	SetDisplayID(value foundation.NSNumber)
	Filter() ISLWindowFilter
	SetFilter(value ISLWindowFilter)
	MetaData() foundation.INSDictionary
	SetMetaData(value foundation.INSDictionary)
	InitInternalDisplay(internal objectivec.IObject, display objectivec.IObject) SLSharingSessionContent
	InitWithDisplayID(id objectivec.IObject) SLSharingSessionContent
	InitWithWindowID(id objectivec.IObject) SLSharingSessionContent
}

// Init initializes the instance.
func (s SLSharingSessionContent) Init() SLSharingSessionContent {
	rv := objc.Send[SLSharingSessionContent](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSharingSessionContent) Autorelease() SLSharingSessionContent {
	rv := objc.Send[SLSharingSessionContent](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSharingSessionContent creates a new SLSharingSessionContent instance.
func NewSLSharingSessionContent() SLSharingSessionContent {
	class := getSLSharingSessionContentClass()
	rv := objc.Send[SLSharingSessionContent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSessionContent/initInternal:display:
func NewSLSharingSessionContentInternalDisplay(internal objectivec.IObject, display objectivec.IObject) SLSharingSessionContent {
	instance := getSLSharingSessionContentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initInternal:display:"), internal, display)
	return SLSharingSessionContentFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSessionContent/initWithDisplayID:
func NewSLSharingSessionContentWithDisplayID(id objectivec.IObject) SLSharingSessionContent {
	instance := getSLSharingSessionContentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDisplayID:"), id)
	return SLSharingSessionContentFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSessionContent/initWithWindowID:
func NewSLSharingSessionContentWithWindowID(id objectivec.IObject) SLSharingSessionContent {
	instance := getSLSharingSessionContentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithWindowID:"), id)
	return SLSharingSessionContentFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSessionContent/initInternal:display:
func (s SLSharingSessionContent) InitInternalDisplay(internal objectivec.IObject, display objectivec.IObject) SLSharingSessionContent {
	rv := objc.Send[SLSharingSessionContent](s.ID, objc.Sel("initInternal:display:"), internal, display)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSessionContent/initWithDisplayID:
func (s SLSharingSessionContent) InitWithDisplayID(id objectivec.IObject) SLSharingSessionContent {
	rv := objc.Send[SLSharingSessionContent](s.ID, objc.Sel("initWithDisplayID:"), id)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSessionContent/initWithWindowID:
func (s SLSharingSessionContent) InitWithWindowID(id objectivec.IObject) SLSharingSessionContent {
	rv := objc.Send[SLSharingSessionContent](s.ID, objc.Sel("initWithWindowID:"), id)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSessionContent/displayID
func (s SLSharingSessionContent) DisplayID() foundation.NSNumber {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("displayID"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (s SLSharingSessionContent) SetDisplayID(value foundation.NSNumber) {
	objc.Send[struct{}](s.ID, objc.Sel("setDisplayID:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSessionContent/filter
func (s SLSharingSessionContent) Filter() ISLWindowFilter {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("filter"))
	return SLWindowFilterFromID(objc.ID(rv))
}
func (s SLSharingSessionContent) SetFilter(value ISLWindowFilter) {
	objc.Send[struct{}](s.ID, objc.Sel("setFilter:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSessionContent/metaData
func (s SLSharingSessionContent) MetaData() foundation.INSDictionary {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("metaData"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (s SLSharingSessionContent) SetMetaData(value foundation.INSDictionary) {
	objc.Send[struct{}](s.ID, objc.Sel("setMetaData:"), value)
}
