// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"context"
	"sync"

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
	rv := objc.Send[ECTestOnlyEventAuthenticationMessage](objc.ID(ec.class), objc.Sel("alloc"))
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
//
// See: https://developer.apple.com/documentation/SkyLight/ECTestOnlyEventAuthenticationMessage
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
//
// See: https://developer.apple.com/documentation/SkyLight/ECTestOnlyEventAuthenticationMessage
type IECTestOnlyEventAuthenticationMessage interface {
	objectivec.IObject

	// Topic: Methods

	Capabilities() uint64
	Context() uint64
	EncodeWithCoder(coder foundation.INSCoder)
	EventType() objectivec.IObject
	MatchesEvent(event coregraphics.CGEventRef) bool
	OriginIdentifier() uint64
	ProxyTargetProcess() objectivec.IObject
	TargetProcess() objectivec.IObject
	Timestamp() uint64
	Valid() bool
	ValidateWithOptionsAndResultBlock(options objectivec.IObject, block VoidHandler)
	InitWithCoder(coder foundation.INSCoder) ECTestOnlyEventAuthenticationMessage
	InitWithValidity(validity bool) ECTestOnlyEventAuthenticationMessage
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (e ECTestOnlyEventAuthenticationMessage) Init() ECTestOnlyEventAuthenticationMessage {
	rv := objc.Send[ECTestOnlyEventAuthenticationMessage](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e ECTestOnlyEventAuthenticationMessage) Autorelease() ECTestOnlyEventAuthenticationMessage {
	rv := objc.Send[ECTestOnlyEventAuthenticationMessage](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewECTestOnlyEventAuthenticationMessage creates a new ECTestOnlyEventAuthenticationMessage instance.
func NewECTestOnlyEventAuthenticationMessage() ECTestOnlyEventAuthenticationMessage {
	class := getECTestOnlyEventAuthenticationMessageClass()
	rv := objc.Send[ECTestOnlyEventAuthenticationMessage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/ECTestOnlyEventAuthenticationMessage/initWithCoder:
func NewECTestOnlyEventAuthenticationMessageWithCoder(coder objectivec.IObject) ECTestOnlyEventAuthenticationMessage {
	instance := getECTestOnlyEventAuthenticationMessageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return ECTestOnlyEventAuthenticationMessageFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/ECTestOnlyEventAuthenticationMessage/initWithValidity:
func NewECTestOnlyEventAuthenticationMessageWithValidity(validity bool) ECTestOnlyEventAuthenticationMessage {
	instance := getECTestOnlyEventAuthenticationMessageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithValidity:"), validity)
	return ECTestOnlyEventAuthenticationMessageFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/ECTestOnlyEventAuthenticationMessage/encodeWithCoder:
func (e ECTestOnlyEventAuthenticationMessage) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](e.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/SkyLight/ECTestOnlyEventAuthenticationMessage/matchesEvent:
func (e ECTestOnlyEventAuthenticationMessage) MatchesEvent(event coregraphics.CGEventRef) bool {
	rv := objc.Send[bool](e.ID, objc.Sel("matchesEvent:"), event)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/ECTestOnlyEventAuthenticationMessage/validateWithOptions:andResultBlock:
func (e ECTestOnlyEventAuthenticationMessage) ValidateWithOptionsAndResultBlock(options objectivec.IObject, block VoidHandler) {
	_block1, _ := NewVoidBlock(block)
	objc.Send[objc.ID](e.ID, objc.Sel("validateWithOptions:andResultBlock:"), options, _block1)
}

// See: https://developer.apple.com/documentation/SkyLight/ECTestOnlyEventAuthenticationMessage/initWithCoder:
func (e ECTestOnlyEventAuthenticationMessage) InitWithCoder(coder foundation.INSCoder) ECTestOnlyEventAuthenticationMessage {
	rv := objc.Send[ECTestOnlyEventAuthenticationMessage](e.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/ECTestOnlyEventAuthenticationMessage/initWithValidity:
func (e ECTestOnlyEventAuthenticationMessage) InitWithValidity(validity bool) ECTestOnlyEventAuthenticationMessage {
	rv := objc.Send[ECTestOnlyEventAuthenticationMessage](e.ID, objc.Sel("initWithValidity:"), validity)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/ECTestOnlyEventAuthenticationMessage/messageWithValidity:
func (_ECTestOnlyEventAuthenticationMessageClass ECTestOnlyEventAuthenticationMessageClass) MessageWithValidity(validity bool) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ECTestOnlyEventAuthenticationMessageClass.class), objc.Sel("messageWithValidity:"), validity)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/ECTestOnlyEventAuthenticationMessage/supportsSecureCoding
func (_ECTestOnlyEventAuthenticationMessageClass ECTestOnlyEventAuthenticationMessageClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_ECTestOnlyEventAuthenticationMessageClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/ECTestOnlyEventAuthenticationMessage/capabilities
func (e ECTestOnlyEventAuthenticationMessage) Capabilities() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("capabilities"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/ECTestOnlyEventAuthenticationMessage/context
func (e ECTestOnlyEventAuthenticationMessage) Context() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("context"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/ECTestOnlyEventAuthenticationMessage/debugDescription
func (e ECTestOnlyEventAuthenticationMessage) DebugDescription() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/ECTestOnlyEventAuthenticationMessage/description
func (e ECTestOnlyEventAuthenticationMessage) Description() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/ECTestOnlyEventAuthenticationMessage/eventType
func (e ECTestOnlyEventAuthenticationMessage) EventType() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("eventType"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/ECTestOnlyEventAuthenticationMessage/hash
func (e ECTestOnlyEventAuthenticationMessage) Hash() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/ECTestOnlyEventAuthenticationMessage/originIdentifier
func (e ECTestOnlyEventAuthenticationMessage) OriginIdentifier() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("originIdentifier"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/ECTestOnlyEventAuthenticationMessage/proxyTargetProcess
func (e ECTestOnlyEventAuthenticationMessage) ProxyTargetProcess() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("proxyTargetProcess"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/ECTestOnlyEventAuthenticationMessage/superclass
func (e ECTestOnlyEventAuthenticationMessage) Superclass() objc.Class {
	rv := objc.Send[objc.Class](e.ID, objc.Sel("superclass"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/ECTestOnlyEventAuthenticationMessage/targetProcess
func (e ECTestOnlyEventAuthenticationMessage) TargetProcess() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("targetProcess"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/ECTestOnlyEventAuthenticationMessage/timestamp
func (e ECTestOnlyEventAuthenticationMessage) Timestamp() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("timestamp"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/ECTestOnlyEventAuthenticationMessage/valid
func (e ECTestOnlyEventAuthenticationMessage) Valid() bool {
	rv := objc.Send[bool](e.ID, objc.Sel("valid"))
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
