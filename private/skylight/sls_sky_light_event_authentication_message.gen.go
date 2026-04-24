// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSSkyLightEventAuthenticationMessage] class.
var (
	_SLSSkyLightEventAuthenticationMessageClass     SLSSkyLightEventAuthenticationMessageClass
	_SLSSkyLightEventAuthenticationMessageClassOnce sync.Once
)

func getSLSSkyLightEventAuthenticationMessageClass() SLSSkyLightEventAuthenticationMessageClass {
	_SLSSkyLightEventAuthenticationMessageClassOnce.Do(func() {
		_SLSSkyLightEventAuthenticationMessageClass = SLSSkyLightEventAuthenticationMessageClass{class: objc.GetClass("SLSSkyLightEventAuthenticationMessage")}
	})
	return _SLSSkyLightEventAuthenticationMessageClass
}

// GetSLSSkyLightEventAuthenticationMessageClass returns the class object for SLSSkyLightEventAuthenticationMessage.
func GetSLSSkyLightEventAuthenticationMessageClass() SLSSkyLightEventAuthenticationMessageClass {
	return getSLSSkyLightEventAuthenticationMessageClass()
}

type SLSSkyLightEventAuthenticationMessageClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSSkyLightEventAuthenticationMessageClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSSkyLightEventAuthenticationMessageClass) Alloc() SLSSkyLightEventAuthenticationMessage {
	rv := objc.Send[SLSSkyLightEventAuthenticationMessage](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSSkyLightEventAuthenticationMessage.Attributes]
//   - [SLSSkyLightEventAuthenticationMessage.Connection]
//   - [SLSSkyLightEventAuthenticationMessage.Flags]
//   - [SLSSkyLightEventAuthenticationMessage.Gesture]
//   - [SLSSkyLightEventAuthenticationMessage.Key]
//   - [SLSSkyLightEventAuthenticationMessage.Location]
//   - [SLSSkyLightEventAuthenticationMessage.MatchesEvent]
//   - [SLSSkyLightEventAuthenticationMessage.Mouse]
//   - [SLSSkyLightEventAuthenticationMessage.Window]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSSkyLightEventAuthenticationMessage
type SLSSkyLightEventAuthenticationMessage struct {
	SLSEventAuthenticationMessage
}

// SLSSkyLightEventAuthenticationMessageFromID constructs a [SLSSkyLightEventAuthenticationMessage] from an objc.ID.
func SLSSkyLightEventAuthenticationMessageFromID(id objc.ID) SLSSkyLightEventAuthenticationMessage {
	return SLSSkyLightEventAuthenticationMessage{SLSEventAuthenticationMessage: SLSEventAuthenticationMessageFromID(id)}
}

// Ensure SLSSkyLightEventAuthenticationMessage implements ISLSSkyLightEventAuthenticationMessage.
var _ ISLSSkyLightEventAuthenticationMessage = SLSSkyLightEventAuthenticationMessage{}

// An interface definition for the [SLSSkyLightEventAuthenticationMessage] class.
//
// # Methods
//
//   - [ISLSSkyLightEventAuthenticationMessage.Attributes]
//   - [ISLSSkyLightEventAuthenticationMessage.Connection]
//   - [ISLSSkyLightEventAuthenticationMessage.Flags]
//   - [ISLSSkyLightEventAuthenticationMessage.Gesture]
//   - [ISLSSkyLightEventAuthenticationMessage.Key]
//   - [ISLSSkyLightEventAuthenticationMessage.Location]
//   - [ISLSSkyLightEventAuthenticationMessage.MatchesEvent]
//   - [ISLSSkyLightEventAuthenticationMessage.Mouse]
//   - [ISLSSkyLightEventAuthenticationMessage.Window]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSSkyLightEventAuthenticationMessage
type ISLSSkyLightEventAuthenticationMessage interface {
	ISLSEventAuthenticationMessage

	// Topic: Methods

	Attributes() unsafe.Pointer
	Connection() uint32
	Flags() uint32
	Gesture() objectivec.IObject
	Key() objectivec.IObject
	Location() corefoundation.CGPoint
	MatchesEvent(event objectivec.IObject) bool
	Mouse() objectivec.IObject
	Window() uint32
}

// Init initializes the instance.
func (s SLSSkyLightEventAuthenticationMessage) Init() SLSSkyLightEventAuthenticationMessage {
	rv := objc.Send[SLSSkyLightEventAuthenticationMessage](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSSkyLightEventAuthenticationMessage) Autorelease() SLSSkyLightEventAuthenticationMessage {
	rv := objc.Send[SLSSkyLightEventAuthenticationMessage](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSSkyLightEventAuthenticationMessage creates a new SLSSkyLightEventAuthenticationMessage instance.
func NewSLSSkyLightEventAuthenticationMessage() SLSSkyLightEventAuthenticationMessage {
	class := getSLSSkyLightEventAuthenticationMessageClass()
	rv := objc.Send[SLSSkyLightEventAuthenticationMessage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/initWithBasis:signature:
func NewSLSSkyLightEventAuthenticationMessageWithBasisSignature(basis objectivec.IObject, signature objectivec.IObject) SLSSkyLightEventAuthenticationMessage {
	instance := getSLSSkyLightEventAuthenticationMessageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBasis:signature:"), basis, signature)
	return SLSSkyLightEventAuthenticationMessageFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/initWithCoder:
func NewSLSSkyLightEventAuthenticationMessageWithCoder(coder objectivec.IObject) SLSSkyLightEventAuthenticationMessage {
	instance := getSLSSkyLightEventAuthenticationMessageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSSkyLightEventAuthenticationMessageFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/initWithEventRecord:pid:version:
func NewSLSSkyLightEventAuthenticationMessageWithEventRecordPidVersion(record *SLSEventRecordRef, pid int, version uint32) SLSSkyLightEventAuthenticationMessage {
	instance := getSLSSkyLightEventAuthenticationMessageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEventRecord:pid:version:"), record, pid, version)
	return SLSSkyLightEventAuthenticationMessageFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSkyLightEventAuthenticationMessage/initWithMessageInitData:
func NewSLSSkyLightEventAuthenticationMessageWithMessageInitData(data unsafe.Pointer) SLSSkyLightEventAuthenticationMessage {
	instance := getSLSSkyLightEventAuthenticationMessageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMessageInitData:"), data)
	return SLSSkyLightEventAuthenticationMessageFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSkyLightEventAuthenticationMessage/matchesEvent:
func (s SLSSkyLightEventAuthenticationMessage) MatchesEvent(event objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("matchesEvent:"), event)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSkyLightEventAuthenticationMessage/attributes
func (s SLSSkyLightEventAuthenticationMessage) Attributes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s.ID, objc.Sel("attributes"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSkyLightEventAuthenticationMessage/connection
func (s SLSSkyLightEventAuthenticationMessage) Connection() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("connection"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSkyLightEventAuthenticationMessage/flags
func (s SLSSkyLightEventAuthenticationMessage) Flags() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("flags"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSkyLightEventAuthenticationMessage/gesture
func (s SLSSkyLightEventAuthenticationMessage) Gesture() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("gesture"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSkyLightEventAuthenticationMessage/key
func (s SLSSkyLightEventAuthenticationMessage) Key() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("key"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSkyLightEventAuthenticationMessage/location
func (s SLSSkyLightEventAuthenticationMessage) Location() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](s.ID, objc.Sel("location"))
	return corefoundation.CGPoint(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSkyLightEventAuthenticationMessage/mouse
func (s SLSSkyLightEventAuthenticationMessage) Mouse() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("mouse"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSkyLightEventAuthenticationMessage/window
func (s SLSSkyLightEventAuthenticationMessage) Window() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("window"))
	return rv
}
