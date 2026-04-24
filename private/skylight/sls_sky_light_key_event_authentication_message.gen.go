// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSSkyLightKeyEventAuthenticationMessage] class.
var (
	_SLSSkyLightKeyEventAuthenticationMessageClass     SLSSkyLightKeyEventAuthenticationMessageClass
	_SLSSkyLightKeyEventAuthenticationMessageClassOnce sync.Once
)

func getSLSSkyLightKeyEventAuthenticationMessageClass() SLSSkyLightKeyEventAuthenticationMessageClass {
	_SLSSkyLightKeyEventAuthenticationMessageClassOnce.Do(func() {
		_SLSSkyLightKeyEventAuthenticationMessageClass = SLSSkyLightKeyEventAuthenticationMessageClass{class: objc.GetClass("SLSSkyLightKeyEventAuthenticationMessage")}
	})
	return _SLSSkyLightKeyEventAuthenticationMessageClass
}

// GetSLSSkyLightKeyEventAuthenticationMessageClass returns the class object for SLSSkyLightKeyEventAuthenticationMessage.
func GetSLSSkyLightKeyEventAuthenticationMessageClass() SLSSkyLightKeyEventAuthenticationMessageClass {
	return getSLSSkyLightKeyEventAuthenticationMessageClass()
}

type SLSSkyLightKeyEventAuthenticationMessageClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSSkyLightKeyEventAuthenticationMessageClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSSkyLightKeyEventAuthenticationMessageClass) Alloc() SLSSkyLightKeyEventAuthenticationMessage {
	rv := objc.Send[SLSSkyLightKeyEventAuthenticationMessage](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSSkyLightKeyEventAuthenticationMessage.CharCode]
//   - [SLSSkyLightKeyEventAuthenticationMessage.Repeat]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSSkyLightKeyEventAuthenticationMessage
type SLSSkyLightKeyEventAuthenticationMessage struct {
	SLSSkyLightEventAuthenticationMessage
}

// SLSSkyLightKeyEventAuthenticationMessageFromID constructs a [SLSSkyLightKeyEventAuthenticationMessage] from an objc.ID.
func SLSSkyLightKeyEventAuthenticationMessageFromID(id objc.ID) SLSSkyLightKeyEventAuthenticationMessage {
	return SLSSkyLightKeyEventAuthenticationMessage{SLSSkyLightEventAuthenticationMessage: SLSSkyLightEventAuthenticationMessageFromID(id)}
}

// Ensure SLSSkyLightKeyEventAuthenticationMessage implements ISLSSkyLightKeyEventAuthenticationMessage.
var _ ISLSSkyLightKeyEventAuthenticationMessage = SLSSkyLightKeyEventAuthenticationMessage{}

// An interface definition for the [SLSSkyLightKeyEventAuthenticationMessage] class.
//
// # Methods
//
//   - [ISLSSkyLightKeyEventAuthenticationMessage.CharCode]
//   - [ISLSSkyLightKeyEventAuthenticationMessage.Repeat]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSSkyLightKeyEventAuthenticationMessage
type ISLSSkyLightKeyEventAuthenticationMessage interface {
	ISLSSkyLightEventAuthenticationMessage

	// Topic: Methods

	CharCode() uint16
	Repeat() int16
}

// Init initializes the instance.
func (s SLSSkyLightKeyEventAuthenticationMessage) Init() SLSSkyLightKeyEventAuthenticationMessage {
	rv := objc.Send[SLSSkyLightKeyEventAuthenticationMessage](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSSkyLightKeyEventAuthenticationMessage) Autorelease() SLSSkyLightKeyEventAuthenticationMessage {
	rv := objc.Send[SLSSkyLightKeyEventAuthenticationMessage](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSSkyLightKeyEventAuthenticationMessage creates a new SLSSkyLightKeyEventAuthenticationMessage instance.
func NewSLSSkyLightKeyEventAuthenticationMessage() SLSSkyLightKeyEventAuthenticationMessage {
	class := getSLSSkyLightKeyEventAuthenticationMessageClass()
	rv := objc.Send[SLSSkyLightKeyEventAuthenticationMessage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/initWithBasis:signature:
func NewSLSSkyLightKeyEventAuthenticationMessageWithBasisSignature(basis objectivec.IObject, signature objectivec.IObject) SLSSkyLightKeyEventAuthenticationMessage {
	instance := getSLSSkyLightKeyEventAuthenticationMessageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBasis:signature:"), basis, signature)
	return SLSSkyLightKeyEventAuthenticationMessageFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/initWithCoder:
func NewSLSSkyLightKeyEventAuthenticationMessageWithCoder(coder objectivec.IObject) SLSSkyLightKeyEventAuthenticationMessage {
	instance := getSLSSkyLightKeyEventAuthenticationMessageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSSkyLightKeyEventAuthenticationMessageFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/initWithEventRecord:pid:version:
func NewSLSSkyLightKeyEventAuthenticationMessageWithEventRecordPidVersion(record *SLSEventRecordRef, pid int, version uint32) SLSSkyLightKeyEventAuthenticationMessage {
	instance := getSLSSkyLightKeyEventAuthenticationMessageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEventRecord:pid:version:"), record, pid, version)
	return SLSSkyLightKeyEventAuthenticationMessageFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSkyLightKeyEventAuthenticationMessage/initWithMessageInitData:
func NewSLSSkyLightKeyEventAuthenticationMessageWithMessageInitData(data unsafe.Pointer) SLSSkyLightKeyEventAuthenticationMessage {
	instance := getSLSSkyLightKeyEventAuthenticationMessageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMessageInitData:"), data)
	return SLSSkyLightKeyEventAuthenticationMessageFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSkyLightKeyEventAuthenticationMessage/charCode
func (s SLSSkyLightKeyEventAuthenticationMessage) CharCode() uint16 {
	rv := objc.Send[uint16](s.ID, objc.Sel("charCode"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSkyLightKeyEventAuthenticationMessage/repeat
func (s SLSSkyLightKeyEventAuthenticationMessage) Repeat() int16 {
	rv := objc.Send[int16](s.ID, objc.Sel("repeat"))
	return rv
}
