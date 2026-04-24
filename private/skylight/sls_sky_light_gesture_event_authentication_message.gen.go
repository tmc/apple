// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSSkyLightGestureEventAuthenticationMessage] class.
var (
	_SLSSkyLightGestureEventAuthenticationMessageClass     SLSSkyLightGestureEventAuthenticationMessageClass
	_SLSSkyLightGestureEventAuthenticationMessageClassOnce sync.Once
)

func getSLSSkyLightGestureEventAuthenticationMessageClass() SLSSkyLightGestureEventAuthenticationMessageClass {
	_SLSSkyLightGestureEventAuthenticationMessageClassOnce.Do(func() {
		_SLSSkyLightGestureEventAuthenticationMessageClass = SLSSkyLightGestureEventAuthenticationMessageClass{class: objc.GetClass("SLSSkyLightGestureEventAuthenticationMessage")}
	})
	return _SLSSkyLightGestureEventAuthenticationMessageClass
}

// GetSLSSkyLightGestureEventAuthenticationMessageClass returns the class object for SLSSkyLightGestureEventAuthenticationMessage.
func GetSLSSkyLightGestureEventAuthenticationMessageClass() SLSSkyLightGestureEventAuthenticationMessageClass {
	return getSLSSkyLightGestureEventAuthenticationMessageClass()
}

type SLSSkyLightGestureEventAuthenticationMessageClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSSkyLightGestureEventAuthenticationMessageClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSSkyLightGestureEventAuthenticationMessageClass) Alloc() SLSSkyLightGestureEventAuthenticationMessage {
	rv := objc.Send[SLSSkyLightGestureEventAuthenticationMessage](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSSkyLightGestureEventAuthenticationMessage.GestureHidType]
//   - [SLSSkyLightGestureEventAuthenticationMessage.GesturePhase]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSSkyLightGestureEventAuthenticationMessage
type SLSSkyLightGestureEventAuthenticationMessage struct {
	SLSSkyLightEventAuthenticationMessage
}

// SLSSkyLightGestureEventAuthenticationMessageFromID constructs a [SLSSkyLightGestureEventAuthenticationMessage] from an objc.ID.
func SLSSkyLightGestureEventAuthenticationMessageFromID(id objc.ID) SLSSkyLightGestureEventAuthenticationMessage {
	return SLSSkyLightGestureEventAuthenticationMessage{SLSSkyLightEventAuthenticationMessage: SLSSkyLightEventAuthenticationMessageFromID(id)}
}

// Ensure SLSSkyLightGestureEventAuthenticationMessage implements ISLSSkyLightGestureEventAuthenticationMessage.
var _ ISLSSkyLightGestureEventAuthenticationMessage = SLSSkyLightGestureEventAuthenticationMessage{}

// An interface definition for the [SLSSkyLightGestureEventAuthenticationMessage] class.
//
// # Methods
//
//   - [ISLSSkyLightGestureEventAuthenticationMessage.GestureHidType]
//   - [ISLSSkyLightGestureEventAuthenticationMessage.GesturePhase]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSSkyLightGestureEventAuthenticationMessage
type ISLSSkyLightGestureEventAuthenticationMessage interface {
	ISLSSkyLightEventAuthenticationMessage

	// Topic: Methods

	GestureHidType() uint32
	GesturePhase() byte
}

// Init initializes the instance.
func (s SLSSkyLightGestureEventAuthenticationMessage) Init() SLSSkyLightGestureEventAuthenticationMessage {
	rv := objc.Send[SLSSkyLightGestureEventAuthenticationMessage](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSSkyLightGestureEventAuthenticationMessage) Autorelease() SLSSkyLightGestureEventAuthenticationMessage {
	rv := objc.Send[SLSSkyLightGestureEventAuthenticationMessage](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSSkyLightGestureEventAuthenticationMessage creates a new SLSSkyLightGestureEventAuthenticationMessage instance.
func NewSLSSkyLightGestureEventAuthenticationMessage() SLSSkyLightGestureEventAuthenticationMessage {
	class := getSLSSkyLightGestureEventAuthenticationMessageClass()
	rv := objc.Send[SLSSkyLightGestureEventAuthenticationMessage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/initWithBasis:signature:
func NewSLSSkyLightGestureEventAuthenticationMessageWithBasisSignature(basis objectivec.IObject, signature objectivec.IObject) SLSSkyLightGestureEventAuthenticationMessage {
	instance := getSLSSkyLightGestureEventAuthenticationMessageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBasis:signature:"), basis, signature)
	return SLSSkyLightGestureEventAuthenticationMessageFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/initWithCoder:
func NewSLSSkyLightGestureEventAuthenticationMessageWithCoder(coder objectivec.IObject) SLSSkyLightGestureEventAuthenticationMessage {
	instance := getSLSSkyLightGestureEventAuthenticationMessageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSSkyLightGestureEventAuthenticationMessageFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/initWithEventRecord:pid:version:
func NewSLSSkyLightGestureEventAuthenticationMessageWithEventRecordPidVersion(record *SLSEventRecordRef, pid int, version uint32) SLSSkyLightGestureEventAuthenticationMessage {
	instance := getSLSSkyLightGestureEventAuthenticationMessageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEventRecord:pid:version:"), record, pid, version)
	return SLSSkyLightGestureEventAuthenticationMessageFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSkyLightGestureEventAuthenticationMessage/initWithMessageInitData:
func NewSLSSkyLightGestureEventAuthenticationMessageWithMessageInitData(data unsafe.Pointer) SLSSkyLightGestureEventAuthenticationMessage {
	instance := getSLSSkyLightGestureEventAuthenticationMessageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMessageInitData:"), data)
	return SLSSkyLightGestureEventAuthenticationMessageFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSkyLightGestureEventAuthenticationMessage/gestureHidType
func (s SLSSkyLightGestureEventAuthenticationMessage) GestureHidType() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("gestureHidType"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSkyLightGestureEventAuthenticationMessage/gesturePhase
func (s SLSSkyLightGestureEventAuthenticationMessage) GesturePhase() byte {
	rv := objc.Send[byte](s.ID, objc.Sel("gesturePhase"))
	return rv
}
