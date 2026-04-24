// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WSSLSToHIDEventTranslator] class.
var (
	_WSSLSToHIDEventTranslatorClass     WSSLSToHIDEventTranslatorClass
	_WSSLSToHIDEventTranslatorClassOnce sync.Once
)

func getWSSLSToHIDEventTranslatorClass() WSSLSToHIDEventTranslatorClass {
	_WSSLSToHIDEventTranslatorClassOnce.Do(func() {
		_WSSLSToHIDEventTranslatorClass = WSSLSToHIDEventTranslatorClass{class: objc.GetClass("_WSSLSToHIDEventTranslator")}
	})
	return _WSSLSToHIDEventTranslatorClass
}

// GetWSSLSToHIDEventTranslatorClass returns the class object for _WSSLSToHIDEventTranslator.
func GetWSSLSToHIDEventTranslatorClass() WSSLSToHIDEventTranslatorClass {
	return getWSSLSToHIDEventTranslatorClass()
}

type WSSLSToHIDEventTranslatorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSSLSToHIDEventTranslatorClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSSLSToHIDEventTranslatorClass) Alloc() WSSLSToHIDEventTranslator {
	rv := objc.Send[WSSLSToHIDEventTranslator](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [WSSLSToHIDEventTranslator.HidEventForSLSEventOutSenderDescriptor]
//
// See: https://developer.apple.com/documentation/SkyLight/_WSSLSToHIDEventTranslator
type WSSLSToHIDEventTranslator struct {
	objectivec.Object
}

// WSSLSToHIDEventTranslatorFromID constructs a [WSSLSToHIDEventTranslator] from an objc.ID.
func WSSLSToHIDEventTranslatorFromID(id objc.ID) WSSLSToHIDEventTranslator {
	return WSSLSToHIDEventTranslator{objectivec.Object{ID: id}}
}

// Ensure WSSLSToHIDEventTranslator implements IWSSLSToHIDEventTranslator.
var _ IWSSLSToHIDEventTranslator = WSSLSToHIDEventTranslator{}

// An interface definition for the [WSSLSToHIDEventTranslator] class.
//
// # Methods
//
//   - [IWSSLSToHIDEventTranslator.HidEventForSLSEventOutSenderDescriptor]
//
// See: https://developer.apple.com/documentation/SkyLight/_WSSLSToHIDEventTranslator
type IWSSLSToHIDEventTranslator interface {
	objectivec.IObject

	// Topic: Methods

	HidEventForSLSEventOutSenderDescriptor(sLSEvent *SLSEventRecordRef, descriptor []objectivec.IObject) uintptr
}

// Init initializes the instance.
func (w WSSLSToHIDEventTranslator) Init() WSSLSToHIDEventTranslator {
	rv := objc.Send[WSSLSToHIDEventTranslator](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSSLSToHIDEventTranslator) Autorelease() WSSLSToHIDEventTranslator {
	rv := objc.Send[WSSLSToHIDEventTranslator](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSSLSToHIDEventTranslator creates a new WSSLSToHIDEventTranslator instance.
func NewWSSLSToHIDEventTranslator() WSSLSToHIDEventTranslator {
	class := getWSSLSToHIDEventTranslatorClass()
	rv := objc.Send[WSSLSToHIDEventTranslator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/_WSSLSToHIDEventTranslator/hidEventForSLSEvent:outSenderDescriptor:
func (w WSSLSToHIDEventTranslator) HidEventForSLSEventOutSenderDescriptor(sLSEvent *SLSEventRecordRef, descriptor []objectivec.IObject) uintptr {
	rv := objc.Send[uintptr](w.ID, objc.Sel("hidEventForSLSEvent:outSenderDescriptor:"), sLSEvent, objectivec.IObjectSliceToNSArray(descriptor))
	return rv
}
