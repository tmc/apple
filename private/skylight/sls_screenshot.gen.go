// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"context"
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSScreenshot] class.
var (
	_SLSScreenshotClass     SLSScreenshotClass
	_SLSScreenshotClassOnce sync.Once
)

func getSLSScreenshotClass() SLSScreenshotClass {
	_SLSScreenshotClassOnce.Do(func() {
		_SLSScreenshotClass = SLSScreenshotClass{class: objc.GetClass("SLSScreenshot")}
	})
	return _SLSScreenshotClass
}

// GetSLSScreenshotClass returns the class object for SLSScreenshot.
func GetSLSScreenshotClass() SLSScreenshotClass {
	return getSLSScreenshotClass()
}

type SLSScreenshotClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSScreenshotClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSScreenshotClass) Alloc() SLSScreenshot {
	rv := objc.Send[SLSScreenshot](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSScreenshot.BridgingHandler]
//   - [SLSScreenshot.Filter]
//   - [SLSScreenshot.SetFilter]
//   - [SLSScreenshot.Properties]
//   - [SLSScreenshot.SetProperties]
//   - [SLSScreenshot.Queue]
//   - [SLSScreenshot.SetQueue]
//   - [SLSScreenshot.SetHandler]
//   - [SLSScreenshot.ZeroWeakSelf]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSScreenshot
type SLSScreenshot struct {
	objectivec.Object
}

// SLSScreenshotFromID constructs a [SLSScreenshot] from an objc.ID.
func SLSScreenshotFromID(id objc.ID) SLSScreenshot {
	return SLSScreenshot{objectivec.Object{ID: id}}
}

// Ensure SLSScreenshot implements ISLSScreenshot.
var _ ISLSScreenshot = SLSScreenshot{}

// An interface definition for the [SLSScreenshot] class.
//
// # Methods
//
//   - [ISLSScreenshot.BridgingHandler]
//   - [ISLSScreenshot.Filter]
//   - [ISLSScreenshot.SetFilter]
//   - [ISLSScreenshot.Properties]
//   - [ISLSScreenshot.SetProperties]
//   - [ISLSScreenshot.Queue]
//   - [ISLSScreenshot.SetQueue]
//   - [ISLSScreenshot.SetHandler]
//   - [ISLSScreenshot.ZeroWeakSelf]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSScreenshot
type ISLSScreenshot interface {
	objectivec.IObject

	// Topic: Methods

	BridgingHandler() VoidHandler
	Filter() ISLContentFilter
	SetFilter(value ISLContentFilter)
	Properties() foundation.INSDictionary
	SetProperties(value foundation.INSDictionary)
	Queue() objectivec.Object
	SetQueue(value objectivec.Object)
	SetHandler(handler VoidHandler)
	ZeroWeakSelf() VoidHandler
}

// Init initializes the instance.
func (s SLSScreenshot) Init() SLSScreenshot {
	rv := objc.Send[SLSScreenshot](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSScreenshot) Autorelease() SLSScreenshot {
	rv := objc.Send[SLSScreenshot](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSScreenshot creates a new SLSScreenshot instance.
func NewSLSScreenshot() SLSScreenshot {
	class := getSLSScreenshotClass()
	rv := objc.Send[SLSScreenshot](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenshot/setHandler:
func (s SLSScreenshot) SetHandler(handler VoidHandler) {
	_block0, _ := NewVoidBlock(handler)
	objc.Send[objc.ID](s.ID, objc.Sel("setHandler:"), _block0)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenshot/convertContentStreamPropertiesToScreenshot:
func (_SLSScreenshotClass SLSScreenshotClass) ConvertContentStreamPropertiesToScreenshot(screenshot objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SLSScreenshotClass.class), objc.Sel("convertContentStreamPropertiesToScreenshot:"), screenshot)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenshot/createScreenshot:properties:queue:handler:error:
func (_SLSScreenshotClass SLSScreenshotClass) CreateScreenshotPropertiesQueueHandlerError(screenshot objectivec.IObject, properties objectivec.IObject, queue objectivec.IObject, handler func()) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_SLSScreenshotClass.class), objc.Sel("createScreenshot:properties:queue:handler:error:"), screenshot, properties, queue, handler, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("createScreenshot:properties:queue:handler:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenshot/replaceColorSpaceInDictionaryWithProfileID:forKey:
func (_SLSScreenshotClass SLSScreenshotClass) ReplaceColorSpaceInDictionaryWithProfileIDForKey(id objectivec.IObject, key objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_SLSScreenshotClass.class), objc.Sel("replaceColorSpaceInDictionaryWithProfileID:forKey:"), id, key)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenshot/bridgingHandler
func (s SLSScreenshot) BridgingHandler() VoidHandler {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("bridgingHandler"))
	_ = rv
	return nil
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenshot/filter
func (s SLSScreenshot) Filter() ISLContentFilter {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("filter"))
	return SLContentFilterFromID(objc.ID(rv))
}
func (s SLSScreenshot) SetFilter(value ISLContentFilter) {
	objc.Send[struct{}](s.ID, objc.Sel("setFilter:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenshot/properties
func (s SLSScreenshot) Properties() foundation.INSDictionary {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("properties"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (s SLSScreenshot) SetProperties(value foundation.INSDictionary) {
	objc.Send[struct{}](s.ID, objc.Sel("setProperties:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenshot/queue
func (s SLSScreenshot) Queue() objectivec.Object {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("queue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (s SLSScreenshot) SetQueue(value objectivec.Object) {
	objc.Send[struct{}](s.ID, objc.Sel("setQueue:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenshot/zeroWeakSelf
func (s SLSScreenshot) ZeroWeakSelf() VoidHandler {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("zeroWeakSelf"))
	_ = rv
	return nil
}

// SetHandlerSync is a synchronous wrapper around [SLSScreenshot.SetHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (s SLSScreenshot) SetHandlerSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	s.SetHandler(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
