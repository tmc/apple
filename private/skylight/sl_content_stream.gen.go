// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"context"
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLContentStream] class.
var (
	_SLContentStreamClass     SLContentStreamClass
	_SLContentStreamClassOnce sync.Once
)

func getSLContentStreamClass() SLContentStreamClass {
	_SLContentStreamClassOnce.Do(func() {
		_SLContentStreamClass = SLContentStreamClass{class: objc.GetClass("SLContentStream")}
	})
	return _SLContentStreamClass
}

// GetSLContentStreamClass returns the class object for SLContentStream.
func GetSLContentStreamClass() SLContentStreamClass {
	return getSLContentStreamClass()
}

type SLContentStreamClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLContentStreamClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLContentStreamClass) Alloc() SLContentStream {
	rv := objc.Send[SLContentStream](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLContentStream.BridgingHandler]
//   - [SLContentStream.CreateStreamWithFilterError]
//   - [SLContentStream.CreateStreamWithSessionError]
//   - [SLContentStream.Filter]
//   - [SLContentStream.SetFilter]
//   - [SLContentStream.PopulateDisplayStreamPropertiesWith]
//   - [SLContentStream.Properties]
//   - [SLContentStream.SetProperties]
//   - [SLContentStream.Queue]
//   - [SLContentStream.SetQueue]
//   - [SLContentStream.Running]
//   - [SLContentStream.SetRunning]
//   - [SLContentStream.Session]
//   - [SLContentStream.SetSession]
//   - [SLContentStream.SetHandler]
//   - [SLContentStream.Start]
//   - [SLContentStream.Stop]
//   - [SLContentStream.Stream]
//   - [SLContentStream.SetStream]
//   - [SLContentStream.UpdateFilterError]
//   - [SLContentStream.UpdatePropertiesError]
//   - [SLContentStream.ZeroWeakSelf]
//   - [SLContentStream.InitWithFilterPropertiesQueueHandler]
//   - [SLContentStream.InitWithFilterPropertiesQueueHandlerError]
//   - [SLContentStream.DebugDescription]
//   - [SLContentStream.Description]
//   - [SLContentStream.Hash]
//   - [SLContentStream.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/SLContentStream
type SLContentStream struct {
	objectivec.Object
}

// SLContentStreamFromID constructs a [SLContentStream] from an objc.ID.
func SLContentStreamFromID(id objc.ID) SLContentStream {
	return SLContentStream{objectivec.Object{ID: id}}
}

// Ensure SLContentStream implements ISLContentStream.
var _ ISLContentStream = SLContentStream{}

// An interface definition for the [SLContentStream] class.
//
// # Methods
//
//   - [ISLContentStream.BridgingHandler]
//   - [ISLContentStream.CreateStreamWithFilterError]
//   - [ISLContentStream.CreateStreamWithSessionError]
//   - [ISLContentStream.Filter]
//   - [ISLContentStream.SetFilter]
//   - [ISLContentStream.PopulateDisplayStreamPropertiesWith]
//   - [ISLContentStream.Properties]
//   - [ISLContentStream.SetProperties]
//   - [ISLContentStream.Queue]
//   - [ISLContentStream.SetQueue]
//   - [ISLContentStream.Running]
//   - [ISLContentStream.SetRunning]
//   - [ISLContentStream.Session]
//   - [ISLContentStream.SetSession]
//   - [ISLContentStream.SetHandler]
//   - [ISLContentStream.Start]
//   - [ISLContentStream.Stop]
//   - [ISLContentStream.Stream]
//   - [ISLContentStream.SetStream]
//   - [ISLContentStream.UpdateFilterError]
//   - [ISLContentStream.UpdatePropertiesError]
//   - [ISLContentStream.ZeroWeakSelf]
//   - [ISLContentStream.InitWithFilterPropertiesQueueHandler]
//   - [ISLContentStream.InitWithFilterPropertiesQueueHandlerError]
//   - [ISLContentStream.DebugDescription]
//   - [ISLContentStream.Description]
//   - [ISLContentStream.Hash]
//   - [ISLContentStream.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/SLContentStream
type ISLContentStream interface {
	objectivec.IObject

	// Topic: Methods

	BridgingHandler() VoidHandler
	CreateStreamWithFilterError(filter objectivec.IObject) (coregraphics.CGDisplayStreamRef, error)
	CreateStreamWithSessionError(session objectivec.IObject) (coregraphics.CGDisplayStreamRef, error)
	Filter() ISLContentFilter
	SetFilter(value ISLContentFilter)
	PopulateDisplayStreamPropertiesWith(properties objectivec.IObject, with objectivec.IObject)
	Properties() foundation.INSDictionary
	SetProperties(value foundation.INSDictionary)
	Queue() objectivec.Object
	SetQueue(value objectivec.Object)
	Running() bool
	SetRunning(value bool)
	Session() ISLSharingSession
	SetSession(value ISLSharingSession)
	SetHandler(handler VoidHandler)
	Start(start []objectivec.IObject) bool
	Stop(stop []objectivec.IObject) bool
	Stream() coregraphics.CGDisplayStreamRef
	SetStream(value coregraphics.CGDisplayStreamRef)
	UpdateFilterError(filter objectivec.IObject) (bool, error)
	UpdatePropertiesError(properties objectivec.IObject) (bool, error)
	ZeroWeakSelf() VoidHandler
	InitWithFilterPropertiesQueueHandler(filter objectivec.IObject, properties objectivec.IObject, queue objectivec.IObject, handler VoidHandler) SLContentStream
	InitWithFilterPropertiesQueueHandlerError(filter objectivec.IObject, properties objectivec.IObject, queue objectivec.IObject, handler func()) (SLContentStream, error)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (s SLContentStream) Init() SLContentStream {
	rv := objc.Send[SLContentStream](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLContentStream) Autorelease() SLContentStream {
	rv := objc.Send[SLContentStream](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLContentStream creates a new SLContentStream instance.
func NewSLContentStream() SLContentStream {
	class := getSLContentStreamClass()
	rv := objc.Send[SLContentStream](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentStream/createStreamWithFilter:error:
func (s SLContentStream) CreateStreamWithFilterError(filter objectivec.IObject) (coregraphics.CGDisplayStreamRef, error) {
	var errorPtr objc.ID
	rv := objc.Send[coregraphics.CGDisplayStreamRef](s.ID, objc.Sel("createStreamWithFilter:error:"), filter, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/SkyLight/SLContentStream/createStreamWithSession:error:
func (s SLContentStream) CreateStreamWithSessionError(session objectivec.IObject) (coregraphics.CGDisplayStreamRef, error) {
	var errorPtr objc.ID
	rv := objc.Send[coregraphics.CGDisplayStreamRef](s.ID, objc.Sel("createStreamWithSession:error:"), session, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/SkyLight/SLContentStream/populateDisplayStreamProperties:with:
func (s SLContentStream) PopulateDisplayStreamPropertiesWith(properties objectivec.IObject, with objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("populateDisplayStreamProperties:with:"), properties, with)
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentStream/setHandler:
func (s SLContentStream) SetHandler(handler VoidHandler) {
	_block0, _ := NewVoidBlock(handler)
	objc.Send[objc.ID](s.ID, objc.Sel("setHandler:"), _block0)
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentStream/start:
func (s SLContentStream) Start(start []objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("start:"), objectivec.IObjectSliceToNSArray(start))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentStream/stop:
func (s SLContentStream) Stop(stop []objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("stop:"), objectivec.IObjectSliceToNSArray(stop))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentStream/updateFilter:error:
func (s SLContentStream) UpdateFilterError(filter objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](s.ID, objc.Sel("updateFilter:error:"), filter, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("updateFilter:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/SkyLight/SLContentStream/updateProperties:error:
func (s SLContentStream) UpdatePropertiesError(properties objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](s.ID, objc.Sel("updateProperties:error:"), properties, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("updateProperties:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/SkyLight/SLContentStream/initWithFilter:properties:queue:handler:
func (s SLContentStream) InitWithFilterPropertiesQueueHandler(filter objectivec.IObject, properties objectivec.IObject, queue objectivec.IObject, handler VoidHandler) SLContentStream {
	_block3, _ := NewVoidBlock(handler)
	rv := objc.Send[SLContentStream](s.ID, objc.Sel("initWithFilter:properties:queue:handler:"), filter, properties, queue, _block3)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentStream/initWithFilter:properties:queue:handler:error:
func (s SLContentStream) InitWithFilterPropertiesQueueHandlerError(filter objectivec.IObject, properties objectivec.IObject, queue objectivec.IObject, handler func()) (SLContentStream, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](s.ID, objc.Sel("initWithFilter:properties:queue:handler:error:"), filter, properties, queue, handler, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SLContentStream{}, foundation.NSErrorFrom(errorPtr)
	}
	return SLContentStreamFromID(rv), nil

}

// See: https://developer.apple.com/documentation/SkyLight/SLContentStream/createScreenshot:properties:queue:handler:error:
func (_SLContentStreamClass SLContentStreamClass) CreateScreenshotPropertiesQueueHandlerError(screenshot objectivec.IObject, properties objectivec.IObject, queue objectivec.IObject, handler func()) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_SLContentStreamClass.class), objc.Sel("createScreenshot:properties:queue:handler:error:"), screenshot, properties, queue, handler, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("createScreenshot:properties:queue:handler:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/SkyLight/SLContentStream/createSessionContentWithFilter:error:
func (_SLContentStreamClass SLContentStreamClass) CreateSessionContentWithFilterError(filter objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_SLContentStreamClass.class), objc.Sel("createSessionContentWithFilter:error:"), filter, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/SkyLight/SLContentStream/bridgingHandler
func (s SLContentStream) BridgingHandler() VoidHandler {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("bridgingHandler"))
	_ = rv
	return nil
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentStream/debugDescription
func (s SLContentStream) DebugDescription() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentStream/description
func (s SLContentStream) Description() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentStream/filter
func (s SLContentStream) Filter() ISLContentFilter {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("filter"))
	return SLContentFilterFromID(objc.ID(rv))
}
func (s SLContentStream) SetFilter(value ISLContentFilter) {
	objc.Send[struct{}](s.ID, objc.Sel("setFilter:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentStream/hash
func (s SLContentStream) Hash() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentStream/properties
func (s SLContentStream) Properties() foundation.INSDictionary {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("properties"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (s SLContentStream) SetProperties(value foundation.INSDictionary) {
	objc.Send[struct{}](s.ID, objc.Sel("setProperties:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentStream/queue
func (s SLContentStream) Queue() objectivec.Object {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("queue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (s SLContentStream) SetQueue(value objectivec.Object) {
	objc.Send[struct{}](s.ID, objc.Sel("setQueue:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentStream/running
func (s SLContentStream) Running() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("running"))
	return rv
}
func (s SLContentStream) SetRunning(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setRunning:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentStream/session
func (s SLContentStream) Session() ISLSharingSession {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("session"))
	return SLSharingSessionFromID(objc.ID(rv))
}
func (s SLContentStream) SetSession(value ISLSharingSession) {
	objc.Send[struct{}](s.ID, objc.Sel("setSession:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentStream/stream
func (s SLContentStream) Stream() coregraphics.CGDisplayStreamRef {
	rv := objc.Send[coregraphics.CGDisplayStreamRef](s.ID, objc.Sel("stream"))
	return coregraphics.CGDisplayStreamRef(rv)
}
func (s SLContentStream) SetStream(value coregraphics.CGDisplayStreamRef) {
	objc.Send[struct{}](s.ID, objc.Sel("setStream:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentStream/superclass
func (s SLContentStream) Superclass() objc.Class {
	rv := objc.Send[objc.Class](s.ID, objc.Sel("superclass"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentStream/zeroWeakSelf
func (s SLContentStream) ZeroWeakSelf() VoidHandler {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("zeroWeakSelf"))
	_ = rv
	return nil
}

// SetHandlerSync is a synchronous wrapper around [SLContentStream.SetHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (s SLContentStream) SetHandlerSync(ctx context.Context) error {
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

// InitWithFilterPropertiesQueueHandlerSync is a synchronous wrapper around [SLContentStream.InitWithFilterPropertiesQueueHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (s SLContentStream) InitWithFilterPropertiesQueueHandlerSync(ctx context.Context, filter objectivec.IObject, properties objectivec.IObject, queue objectivec.IObject) error {
	done := make(chan struct{}, 1)
	s.InitWithFilterPropertiesQueueHandler(filter, properties, queue, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
