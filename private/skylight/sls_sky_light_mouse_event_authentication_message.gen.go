// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSSkyLightMouseEventAuthenticationMessage] class.
var (
	_SLSSkyLightMouseEventAuthenticationMessageClass     SLSSkyLightMouseEventAuthenticationMessageClass
	_SLSSkyLightMouseEventAuthenticationMessageClassOnce sync.Once
)

func getSLSSkyLightMouseEventAuthenticationMessageClass() SLSSkyLightMouseEventAuthenticationMessageClass {
	_SLSSkyLightMouseEventAuthenticationMessageClassOnce.Do(func() {
		_SLSSkyLightMouseEventAuthenticationMessageClass = SLSSkyLightMouseEventAuthenticationMessageClass{class: objc.GetClass("SLSSkyLightMouseEventAuthenticationMessage")}
	})
	return _SLSSkyLightMouseEventAuthenticationMessageClass
}

// GetSLSSkyLightMouseEventAuthenticationMessageClass returns the class object for SLSSkyLightMouseEventAuthenticationMessage.
func GetSLSSkyLightMouseEventAuthenticationMessageClass() SLSSkyLightMouseEventAuthenticationMessageClass {
	return getSLSSkyLightMouseEventAuthenticationMessageClass()
}

type SLSSkyLightMouseEventAuthenticationMessageClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSSkyLightMouseEventAuthenticationMessageClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSSkyLightMouseEventAuthenticationMessageClass) Alloc() SLSSkyLightMouseEventAuthenticationMessage {
	rv := objc.Send[SLSSkyLightMouseEventAuthenticationMessage](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSSkyLightMouseEventAuthenticationMessage.ButtonNumber]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSSkyLightMouseEventAuthenticationMessage
type SLSSkyLightMouseEventAuthenticationMessage struct {
	SLSSkyLightEventAuthenticationMessage
}

// SLSSkyLightMouseEventAuthenticationMessageFromID constructs a [SLSSkyLightMouseEventAuthenticationMessage] from an objc.ID.
func SLSSkyLightMouseEventAuthenticationMessageFromID(id objc.ID) SLSSkyLightMouseEventAuthenticationMessage {
	return SLSSkyLightMouseEventAuthenticationMessage{SLSSkyLightEventAuthenticationMessage: SLSSkyLightEventAuthenticationMessageFromID(id)}
}

// Ensure SLSSkyLightMouseEventAuthenticationMessage implements ISLSSkyLightMouseEventAuthenticationMessage.
var _ ISLSSkyLightMouseEventAuthenticationMessage = SLSSkyLightMouseEventAuthenticationMessage{}

// An interface definition for the [SLSSkyLightMouseEventAuthenticationMessage] class.
//
// # Methods
//
//   - [ISLSSkyLightMouseEventAuthenticationMessage.ButtonNumber]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSSkyLightMouseEventAuthenticationMessage
type ISLSSkyLightMouseEventAuthenticationMessage interface {
	ISLSSkyLightEventAuthenticationMessage

	// Topic: Methods

	ButtonNumber() int8
}

// Init initializes the instance.
func (s SLSSkyLightMouseEventAuthenticationMessage) Init() SLSSkyLightMouseEventAuthenticationMessage {
	rv := objc.Send[SLSSkyLightMouseEventAuthenticationMessage](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSSkyLightMouseEventAuthenticationMessage) Autorelease() SLSSkyLightMouseEventAuthenticationMessage {
	rv := objc.Send[SLSSkyLightMouseEventAuthenticationMessage](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSSkyLightMouseEventAuthenticationMessage creates a new SLSSkyLightMouseEventAuthenticationMessage instance.
func NewSLSSkyLightMouseEventAuthenticationMessage() SLSSkyLightMouseEventAuthenticationMessage {
	class := getSLSSkyLightMouseEventAuthenticationMessageClass()
	rv := objc.Send[SLSSkyLightMouseEventAuthenticationMessage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/initWithBasis:signature:
func NewSLSSkyLightMouseEventAuthenticationMessageWithBasisSignature(basis objectivec.IObject, signature objectivec.IObject) SLSSkyLightMouseEventAuthenticationMessage {
	instance := getSLSSkyLightMouseEventAuthenticationMessageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBasis:signature:"), basis, signature)
	return SLSSkyLightMouseEventAuthenticationMessageFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/initWithCoder:
func NewSLSSkyLightMouseEventAuthenticationMessageWithCoder(coder objectivec.IObject) SLSSkyLightMouseEventAuthenticationMessage {
	instance := getSLSSkyLightMouseEventAuthenticationMessageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSSkyLightMouseEventAuthenticationMessageFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/initWithEventRecord:pid:version:
func NewSLSSkyLightMouseEventAuthenticationMessageWithEventRecordPidVersion(record *SLSEventRecordRef, pid int, version uint32) SLSSkyLightMouseEventAuthenticationMessage {
	instance := getSLSSkyLightMouseEventAuthenticationMessageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEventRecord:pid:version:"), record, pid, version)
	return SLSSkyLightMouseEventAuthenticationMessageFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSkyLightMouseEventAuthenticationMessage/initWithMessageInitData:
func NewSLSSkyLightMouseEventAuthenticationMessageWithMessageInitData(data unsafe.Pointer) SLSSkyLightMouseEventAuthenticationMessage {
	instance := getSLSSkyLightMouseEventAuthenticationMessageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMessageInitData:"), data)
	return SLSSkyLightMouseEventAuthenticationMessageFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSkyLightMouseEventAuthenticationMessage/buttonNumber
func (s SLSSkyLightMouseEventAuthenticationMessage) ButtonNumber() int8 {
	rv := objc.Send[int8](s.ID, objc.Sel("buttonNumber"))
	return rv
}
