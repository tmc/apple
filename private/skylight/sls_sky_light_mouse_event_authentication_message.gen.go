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
	rv := objc.SendIfResponds[SLSSkyLightMouseEventAuthenticationMessage](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSSkyLightMouseEventAuthenticationMessage.ButtonNumber]
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
type ISLSSkyLightMouseEventAuthenticationMessage interface {
	ISLSSkyLightEventAuthenticationMessage

	// Topic: Methods

	ButtonNumber() int8
}

// Init initializes the instance.
func (s SLSSkyLightMouseEventAuthenticationMessage) Init() SLSSkyLightMouseEventAuthenticationMessage {
	rv := objc.SendIfResponds[SLSSkyLightMouseEventAuthenticationMessage](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSSkyLightMouseEventAuthenticationMessage) Autorelease() SLSSkyLightMouseEventAuthenticationMessage {
	rv := objc.SendIfResponds[SLSSkyLightMouseEventAuthenticationMessage](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSSkyLightMouseEventAuthenticationMessage creates a new SLSSkyLightMouseEventAuthenticationMessage instance.
func NewSLSSkyLightMouseEventAuthenticationMessage() SLSSkyLightMouseEventAuthenticationMessage {
	class := getSLSSkyLightMouseEventAuthenticationMessageClass()
	rv := objc.SendIfResponds[SLSSkyLightMouseEventAuthenticationMessage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLSSkyLightMouseEventAuthenticationMessageWithBasisSignature(basis objectivec.IObject, signature objectivec.IObject) SLSSkyLightMouseEventAuthenticationMessage {
	instance := getSLSSkyLightMouseEventAuthenticationMessageClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithBasis:signature:"), basis, signature)
	return SLSSkyLightMouseEventAuthenticationMessageFromID(rv)
}

func NewSLSSkyLightMouseEventAuthenticationMessageWithCoder(coder objectivec.IObject) SLSSkyLightMouseEventAuthenticationMessage {
	instance := getSLSSkyLightMouseEventAuthenticationMessageClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSSkyLightMouseEventAuthenticationMessageFromID(rv)
}

func NewSLSSkyLightMouseEventAuthenticationMessageWithEventRecordPidVersion(record *SLSEventRecord, pid int, version uint32) SLSSkyLightMouseEventAuthenticationMessage {
	instance := getSLSSkyLightMouseEventAuthenticationMessageClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithEventRecord:pid:version:"), unsafe.Pointer(record), pid, version)
	return SLSSkyLightMouseEventAuthenticationMessageFromID(rv)
}

func NewSLSSkyLightMouseEventAuthenticationMessageWithMessageInitData(data *MessageInitData) SLSSkyLightMouseEventAuthenticationMessage {
	instance := getSLSSkyLightMouseEventAuthenticationMessageClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithMessageInitData:"), unsafe.Pointer(data))
	return SLSSkyLightMouseEventAuthenticationMessageFromID(rv)
}

func (s SLSSkyLightMouseEventAuthenticationMessage) ButtonNumber() int8 {
	rv := objc.SendIfResponds[int8](s.ID, objc.Sel("buttonNumber"))
	return rv
}
