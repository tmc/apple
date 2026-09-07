// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ECTestOnlyEventAuthenticationMessage] class.
var (
	_ECTestOnlyEventAuthenticationMessageClass     ECTestOnlyEventAuthenticationMessageClass
	_ECTestOnlyEventAuthenticationMessageClassOnce sync.Once
)

func getECTestOnlyEventAuthenticationMessageClass() ECTestOnlyEventAuthenticationMessageClass {
	_ECTestOnlyEventAuthenticationMessageClassOnce.Do(func() {
		_ECTestOnlyEventAuthenticationMessageClass = ECTestOnlyEventAuthenticationMessageClass{class: objc.GetClass("ECTestOnlyEventAuthenticationMessage")}
	})
	return _ECTestOnlyEventAuthenticationMessageClass
}

// GetECTestOnlyEventAuthenticationMessageClass returns the class object for ECTestOnlyEventAuthenticationMessage.
func GetECTestOnlyEventAuthenticationMessageClass() ECTestOnlyEventAuthenticationMessageClass {
	return getECTestOnlyEventAuthenticationMessageClass()
}

type ECTestOnlyEventAuthenticationMessageClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec ECTestOnlyEventAuthenticationMessageClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec ECTestOnlyEventAuthenticationMessageClass) Alloc() ECTestOnlyEventAuthenticationMessage {
	rv := objc.SendIfResponds[ECTestOnlyEventAuthenticationMessage](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ECTestOnlyEventAuthenticationMessage.Capabilities]
//   - [ECTestOnlyEventAuthenticationMessage.Context]
//   - [ECTestOnlyEventAuthenticationMessage.EncodeWithCoder]
//   - [ECTestOnlyEventAuthenticationMessage.EventType]
//   - [ECTestOnlyEventAuthenticationMessage.MatchesEvent]
//   - [ECTestOnlyEventAuthenticationMessage.OriginIdentifier]
//   - [ECTestOnlyEventAuthenticationMessage.ProxyTargetProcess]
//   - [ECTestOnlyEventAuthenticationMessage.TargetProcess]
//   - [ECTestOnlyEventAuthenticationMessage.Timestamp]
//   - [ECTestOnlyEventAuthenticationMessage.Valid]
//   - [ECTestOnlyEventAuthenticationMessage.ValidateWithOptionsAndResultBlock]
//   - [ECTestOnlyEventAuthenticationMessage.InitWithCoder]
//   - [ECTestOnlyEventAuthenticationMessage.InitWithValidity]
//   - [ECTestOnlyEventAuthenticationMessage.DebugDescription]
//   - [ECTestOnlyEventAuthenticationMessage.Description]
//   - [ECTestOnlyEventAuthenticationMessage.Hash]
//   - [ECTestOnlyEventAuthenticationMessage.Superclass]
type ECTestOnlyEventAuthenticationMessage struct {
	objectivec.Object
}

// ECTestOnlyEventAuthenticationMessageFromID constructs a [ECTestOnlyEventAuthenticationMessage] from an objc.ID.
func ECTestOnlyEventAuthenticationMessageFromID(id objc.ID) ECTestOnlyEventAuthenticationMessage {
	return ECTestOnlyEventAuthenticationMessage{objectivec.Object{ID: id}}
}

// Ensure ECTestOnlyEventAuthenticationMessage implements IECTestOnlyEventAuthenticationMessage.
var _ IECTestOnlyEventAuthenticationMessage = ECTestOnlyEventAuthenticationMessage{}

// An interface definition for the [ECTestOnlyEventAuthenticationMessage] class.
//
// # Methods
//
//   - [IECTestOnlyEventAuthenticationMessage.Capabilities]
//   - [IECTestOnlyEventAuthenticationMessage.Context]
//   - [IECTestOnlyEventAuthenticationMessage.EncodeWithCoder]
//   - [IECTestOnlyEventAuthenticationMessage.EventType]
//   - [IECTestOnlyEventAuthenticationMessage.MatchesEvent]
//   - [IECTestOnlyEventAuthenticationMessage.OriginIdentifier]
//   - [IECTestOnlyEventAuthenticationMessage.ProxyTargetProcess]
//   - [IECTestOnlyEventAuthenticationMessage.TargetProcess]
//   - [IECTestOnlyEventAuthenticationMessage.Timestamp]
//   - [IECTestOnlyEventAuthenticationMessage.Valid]
//   - [IECTestOnlyEventAuthenticationMessage.ValidateWithOptionsAndResultBlock]
//   - [IECTestOnlyEventAuthenticationMessage.InitWithCoder]
//   - [IECTestOnlyEventAuthenticationMessage.InitWithValidity]
//   - [IECTestOnlyEventAuthenticationMessage.DebugDescription]
//   - [IECTestOnlyEventAuthenticationMessage.Description]
//   - [IECTestOnlyEventAuthenticationMessage.Hash]
//   - [IECTestOnlyEventAuthenticationMessage.Superclass]
type IECTestOnlyEventAuthenticationMessage interface {
	objectivec.IObject

	// Topic: Methods

	Capabilities() uint64
	Context() uint64
	EncodeWithCoder(coder foundation.INSCoder)
	EventType() unsafe.Pointer
	MatchesEvent(event coregraphics.CGEvent) bool
	OriginIdentifier() uint64
	ProxyTargetProcess() unsafe.Pointer
	TargetProcess() unsafe.Pointer
	Timestamp() uint64
	Valid() bool
	ValidateWithOptionsAndResultBlock(options objectivec.IObject, block VoidHandler)
	InitWithCoder(coder foundation.INSCoder) ECTestOnlyEventAuthenticationMessage
	InitWithValidity(validity bool) ECTestOnlyEventAuthenticationMessage
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (e ECTestOnlyEventAuthenticationMessage) Init() ECTestOnlyEventAuthenticationMessage {
	rv := objc.SendIfResponds[ECTestOnlyEventAuthenticationMessage](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e ECTestOnlyEventAuthenticationMessage) Autorelease() ECTestOnlyEventAuthenticationMessage {
	rv := objc.SendIfResponds[ECTestOnlyEventAuthenticationMessage](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewECTestOnlyEventAuthenticationMessage creates a new ECTestOnlyEventAuthenticationMessage instance.
func NewECTestOnlyEventAuthenticationMessage() ECTestOnlyEventAuthenticationMessage {
	class := getECTestOnlyEventAuthenticationMessageClass()
	rv := objc.SendIfResponds[ECTestOnlyEventAuthenticationMessage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewECTestOnlyEventAuthenticationMessageWithCoder(coder objectivec.IObject) ECTestOnlyEventAuthenticationMessage {
	instance := getECTestOnlyEventAuthenticationMessageClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return ECTestOnlyEventAuthenticationMessageFromID(rv)
}

func NewECTestOnlyEventAuthenticationMessageWithValidity(validity bool) ECTestOnlyEventAuthenticationMessage {
	instance := getECTestOnlyEventAuthenticationMessageClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithValidity:"), validity)
	return ECTestOnlyEventAuthenticationMessageFromID(rv)
}

func (e ECTestOnlyEventAuthenticationMessage) EncodeWithCoder(coder foundation.INSCoder) {
	objc.SendIfResponds[objc.ID](e.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (e ECTestOnlyEventAuthenticationMessage) MatchesEvent(event coregraphics.CGEvent) bool {
	rv := objc.SendIfResponds[bool](e.ID, objc.Sel("matchesEvent:"), event)
	return rv
}
func (e ECTestOnlyEventAuthenticationMessage) ValidateWithOptionsAndResultBlock(options objectivec.IObject, block VoidHandler) {
	_block1, _ := NewVoidBlock(block)
	objc.SendIfResponds[objc.ID](e.ID, objc.Sel("validateWithOptions:andResultBlock:"), options, _block1)
}
func (e ECTestOnlyEventAuthenticationMessage) InitWithCoder(coder foundation.INSCoder) ECTestOnlyEventAuthenticationMessage {
	rv := objc.SendIfResponds[ECTestOnlyEventAuthenticationMessage](e.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (e ECTestOnlyEventAuthenticationMessage) InitWithValidity(validity bool) ECTestOnlyEventAuthenticationMessage {
	rv := objc.SendIfResponds[ECTestOnlyEventAuthenticationMessage](e.ID, objc.Sel("initWithValidity:"), validity)
	return rv
}

func (_ECTestOnlyEventAuthenticationMessageClass ECTestOnlyEventAuthenticationMessageClass) MessageWithValidity(validity bool) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ECTestOnlyEventAuthenticationMessageClass.class), objc.Sel("messageWithValidity:"), validity)
	return objectivec.Object{ID: rv}
}
func (_ECTestOnlyEventAuthenticationMessageClass ECTestOnlyEventAuthenticationMessageClass) SupportsSecureCoding() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_ECTestOnlyEventAuthenticationMessageClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (e ECTestOnlyEventAuthenticationMessage) Capabilities() uint64 {
	rv := objc.SendIfResponds[uint64](e.ID, objc.Sel("capabilities"))
	return rv
}
func (e ECTestOnlyEventAuthenticationMessage) Context() uint64 {
	rv := objc.SendIfResponds[uint64](e.ID, objc.Sel("context"))
	return rv
}
func (e ECTestOnlyEventAuthenticationMessage) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](e.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (e ECTestOnlyEventAuthenticationMessage) Description() string {
	rv := objc.SendIfResponds[objc.ID](e.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (e ECTestOnlyEventAuthenticationMessage) EventType() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](e.ID, objc.Sel("eventType"))
	return rv
}
func (e ECTestOnlyEventAuthenticationMessage) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](e.ID, objc.Sel("hash"))
	return rv
}
func (e ECTestOnlyEventAuthenticationMessage) OriginIdentifier() uint64 {
	rv := objc.SendIfResponds[uint64](e.ID, objc.Sel("originIdentifier"))
	return rv
}
func (e ECTestOnlyEventAuthenticationMessage) ProxyTargetProcess() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](e.ID, objc.Sel("proxyTargetProcess"))
	return rv
}
func (e ECTestOnlyEventAuthenticationMessage) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](e.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
func (e ECTestOnlyEventAuthenticationMessage) TargetProcess() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](e.ID, objc.Sel("targetProcess"))
	return rv
}
func (e ECTestOnlyEventAuthenticationMessage) Timestamp() uint64 {
	rv := objc.SendIfResponds[uint64](e.ID, objc.Sel("timestamp"))
	return rv
}
func (e ECTestOnlyEventAuthenticationMessage) Valid() bool {
	rv := objc.SendIfResponds[bool](e.ID, objc.Sel("valid"))
	return rv
}

// ValidateWithOptionsAndResultBlockSync is a synchronous wrapper around [ECTestOnlyEventAuthenticationMessage.ValidateWithOptionsAndResultBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (e ECTestOnlyEventAuthenticationMessage) ValidateWithOptionsAndResultBlockSync(ctx context.Context, options objectivec.IObject) error {
	done := make(chan struct{}, 1)
	e.ValidateWithOptionsAndResultBlock(options, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
