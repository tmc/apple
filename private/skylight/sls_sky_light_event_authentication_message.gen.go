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
	rv := objc.SendIfResponds[SLSSkyLightEventAuthenticationMessage](objc.ID(sc.class), objc.Sel("alloc"))
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
type ISLSSkyLightEventAuthenticationMessage interface {
	ISLSEventAuthenticationMessage

	// Topic: Methods

	Attributes() unsafe.Pointer
	Connection() uint32
	Flags() uint32
	Gesture() unsafe.Pointer
	Key() unsafe.Pointer
	Location() corefoundation.CGPoint
	MatchesEvent(event *CGEvent) bool
	Mouse() unsafe.Pointer
	Window() uint32
}

// Init initializes the instance.
func (s SLSSkyLightEventAuthenticationMessage) Init() SLSSkyLightEventAuthenticationMessage {
	rv := objc.SendIfResponds[SLSSkyLightEventAuthenticationMessage](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSSkyLightEventAuthenticationMessage) Autorelease() SLSSkyLightEventAuthenticationMessage {
	rv := objc.SendIfResponds[SLSSkyLightEventAuthenticationMessage](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSSkyLightEventAuthenticationMessage creates a new SLSSkyLightEventAuthenticationMessage instance.
func NewSLSSkyLightEventAuthenticationMessage() SLSSkyLightEventAuthenticationMessage {
	class := getSLSSkyLightEventAuthenticationMessageClass()
	rv := objc.SendIfResponds[SLSSkyLightEventAuthenticationMessage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLSSkyLightEventAuthenticationMessageWithBasisSignature(basis objectivec.IObject, signature objectivec.IObject) SLSSkyLightEventAuthenticationMessage {
	instance := getSLSSkyLightEventAuthenticationMessageClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithBasis:signature:"), basis, signature)
	return SLSSkyLightEventAuthenticationMessageFromID(rv)
}

func NewSLSSkyLightEventAuthenticationMessageWithCoder(coder objectivec.IObject) SLSSkyLightEventAuthenticationMessage {
	instance := getSLSSkyLightEventAuthenticationMessageClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSSkyLightEventAuthenticationMessageFromID(rv)
}

func NewSLSSkyLightEventAuthenticationMessageWithEventRecordPidVersion(record *SLSEventRecord, pid int, version uint32) SLSSkyLightEventAuthenticationMessage {
	instance := getSLSSkyLightEventAuthenticationMessageClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithEventRecord:pid:version:"), unsafe.Pointer(record), pid, version)
	return SLSSkyLightEventAuthenticationMessageFromID(rv)
}

func NewSLSSkyLightEventAuthenticationMessageWithMessageInitData(data *MessageInitData) SLSSkyLightEventAuthenticationMessage {
	instance := getSLSSkyLightEventAuthenticationMessageClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithMessageInitData:"), unsafe.Pointer(data))
	return SLSSkyLightEventAuthenticationMessageFromID(rv)
}

func (s SLSSkyLightEventAuthenticationMessage) MatchesEvent(event *CGEvent) bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("matchesEvent:"), event)
	return rv
}

func (s SLSSkyLightEventAuthenticationMessage) Attributes() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](s.ID, objc.Sel("attributes"))
	return rv
}
func (s SLSSkyLightEventAuthenticationMessage) Connection() uint32 {
	rv := objc.SendIfResponds[uint32](s.ID, objc.Sel("connection"))
	return rv
}
func (s SLSSkyLightEventAuthenticationMessage) Flags() uint32 {
	rv := objc.SendIfResponds[uint32](s.ID, objc.Sel("flags"))
	return rv
}
func (s SLSSkyLightEventAuthenticationMessage) Gesture() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](s.ID, objc.Sel("gesture"))
	return rv
}
func (s SLSSkyLightEventAuthenticationMessage) Key() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](s.ID, objc.Sel("key"))
	return rv
}
func (s SLSSkyLightEventAuthenticationMessage) Location() corefoundation.CGPoint {
	rv := objc.SendIfResponds[corefoundation.CGPoint](s.ID, objc.Sel("location"))
	return corefoundation.CGPoint(rv)
}
func (s SLSSkyLightEventAuthenticationMessage) Mouse() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](s.ID, objc.Sel("mouse"))
	return rv
}
func (s SLSSkyLightEventAuthenticationMessage) Window() uint32 {
	rv := objc.SendIfResponds[uint32](s.ID, objc.Sel("window"))
	return rv
}
