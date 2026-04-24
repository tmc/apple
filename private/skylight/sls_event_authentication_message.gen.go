// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSEventAuthenticationMessage] class.
var (
	_SLSEventAuthenticationMessageClass     SLSEventAuthenticationMessageClass
	_SLSEventAuthenticationMessageClassOnce sync.Once
)

func getSLSEventAuthenticationMessageClass() SLSEventAuthenticationMessageClass {
	_SLSEventAuthenticationMessageClassOnce.Do(func() {
		_SLSEventAuthenticationMessageClass = SLSEventAuthenticationMessageClass{class: objc.GetClass("SLSEventAuthenticationMessage")}
	})
	return _SLSEventAuthenticationMessageClass
}

// GetSLSEventAuthenticationMessageClass returns the class object for SLSEventAuthenticationMessage.
func GetSLSEventAuthenticationMessageClass() SLSEventAuthenticationMessageClass {
	return getSLSEventAuthenticationMessageClass()
}

type SLSEventAuthenticationMessageClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSEventAuthenticationMessageClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSEventAuthenticationMessageClass) Alloc() SLSEventAuthenticationMessage {
	rv := objc.Send[SLSEventAuthenticationMessage](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSEventAuthenticationMessage.AddToSigningContext]
//   - [SLSEventAuthenticationMessage.Capabilities]
//   - [SLSEventAuthenticationMessage.Context]
//   - [SLSEventAuthenticationMessage.CopySignedByKey]
//   - [SLSEventAuthenticationMessage.CopyWithSignature]
//   - [SLSEventAuthenticationMessage.CopyWithZoneSignature]
//   - [SLSEventAuthenticationMessage.EncodeWithCoder]
//   - [SLSEventAuthenticationMessage.EventType]
//   - [SLSEventAuthenticationMessage.HasSuperclassEquivalentUnsignedData]
//   - [SLSEventAuthenticationMessage.IsSuperclassEquivalent]
//   - [SLSEventAuthenticationMessage.MatchesMessageData]
//   - [SLSEventAuthenticationMessage.OriginIdentifier]
//   - [SLSEventAuthenticationMessage.ProxyTargetProcess]
//   - [SLSEventAuthenticationMessage.Signature]
//   - [SLSEventAuthenticationMessage.TargetProcess]
//   - [SLSEventAuthenticationMessage.Timestamp]
//   - [SLSEventAuthenticationMessage.ValidateWithOptionsAndResultBlock]
//   - [SLSEventAuthenticationMessage.InitWithBasisSignature]
//   - [SLSEventAuthenticationMessage.InitWithCoder]
//   - [SLSEventAuthenticationMessage.InitWithEventRecordPidVersion]
//   - [SLSEventAuthenticationMessage.InitWithMessageInitData]
//   - [SLSEventAuthenticationMessage.DebugDescription]
//   - [SLSEventAuthenticationMessage.Description]
//   - [SLSEventAuthenticationMessage.Hash]
//   - [SLSEventAuthenticationMessage.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage
type SLSEventAuthenticationMessage struct {
	objectivec.Object
}

// SLSEventAuthenticationMessageFromID constructs a [SLSEventAuthenticationMessage] from an objc.ID.
func SLSEventAuthenticationMessageFromID(id objc.ID) SLSEventAuthenticationMessage {
	return SLSEventAuthenticationMessage{objectivec.Object{ID: id}}
}

// Ensure SLSEventAuthenticationMessage implements ISLSEventAuthenticationMessage.
var _ ISLSEventAuthenticationMessage = SLSEventAuthenticationMessage{}

// An interface definition for the [SLSEventAuthenticationMessage] class.
//
// # Methods
//
//   - [ISLSEventAuthenticationMessage.AddToSigningContext]
//   - [ISLSEventAuthenticationMessage.Capabilities]
//   - [ISLSEventAuthenticationMessage.Context]
//   - [ISLSEventAuthenticationMessage.CopySignedByKey]
//   - [ISLSEventAuthenticationMessage.CopyWithSignature]
//   - [ISLSEventAuthenticationMessage.CopyWithZoneSignature]
//   - [ISLSEventAuthenticationMessage.EncodeWithCoder]
//   - [ISLSEventAuthenticationMessage.EventType]
//   - [ISLSEventAuthenticationMessage.HasSuperclassEquivalentUnsignedData]
//   - [ISLSEventAuthenticationMessage.IsSuperclassEquivalent]
//   - [ISLSEventAuthenticationMessage.MatchesMessageData]
//   - [ISLSEventAuthenticationMessage.OriginIdentifier]
//   - [ISLSEventAuthenticationMessage.ProxyTargetProcess]
//   - [ISLSEventAuthenticationMessage.Signature]
//   - [ISLSEventAuthenticationMessage.TargetProcess]
//   - [ISLSEventAuthenticationMessage.Timestamp]
//   - [ISLSEventAuthenticationMessage.ValidateWithOptionsAndResultBlock]
//   - [ISLSEventAuthenticationMessage.InitWithBasisSignature]
//   - [ISLSEventAuthenticationMessage.InitWithCoder]
//   - [ISLSEventAuthenticationMessage.InitWithEventRecordPidVersion]
//   - [ISLSEventAuthenticationMessage.InitWithMessageInitData]
//   - [ISLSEventAuthenticationMessage.DebugDescription]
//   - [ISLSEventAuthenticationMessage.Description]
//   - [ISLSEventAuthenticationMessage.Hash]
//   - [ISLSEventAuthenticationMessage.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage
type ISLSEventAuthenticationMessage interface {
	objectivec.IObject

	// Topic: Methods

	AddToSigningContext(context objectivec.IObject)
	Capabilities() uint64
	Context() uint64
	CopySignedByKey(key objectivec.IObject) objectivec.IObject
	CopyWithSignature(signature objectivec.IObject) objectivec.IObject
	CopyWithZoneSignature(zone NSZoneRef, signature objectivec.IObject) objectivec.IObject
	EncodeWithCoder(coder foundation.INSCoder)
	EventType() ISLSEventAuthenticationMessageEventType
	HasSuperclassEquivalentUnsignedData(data objectivec.IObject) bool
	IsSuperclassEquivalent(equivalent objectivec.IObject) bool
	MatchesMessageData(data objectivec.IObject) bool
	OriginIdentifier() uint64
	ProxyTargetProcess() ISLSEventAuthenticationMessageVersionedPID
	Signature() foundation.INSData
	TargetProcess() ISLSEventAuthenticationMessageVersionedPID
	Timestamp() uint64
	ValidateWithOptionsAndResultBlock(options objectivec.IObject, block VoidHandler)
	InitWithBasisSignature(basis objectivec.IObject, signature objectivec.IObject) SLSEventAuthenticationMessage
	InitWithCoder(coder foundation.INSCoder) SLSEventAuthenticationMessage
	InitWithEventRecordPidVersion(record *SLSEventRecordRef, pid int, version uint32) SLSEventAuthenticationMessage
	InitWithMessageInitData(data unsafe.Pointer) SLSEventAuthenticationMessage
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (s SLSEventAuthenticationMessage) Init() SLSEventAuthenticationMessage {
	rv := objc.Send[SLSEventAuthenticationMessage](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSEventAuthenticationMessage) Autorelease() SLSEventAuthenticationMessage {
	rv := objc.Send[SLSEventAuthenticationMessage](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSEventAuthenticationMessage creates a new SLSEventAuthenticationMessage instance.
func NewSLSEventAuthenticationMessage() SLSEventAuthenticationMessage {
	class := getSLSEventAuthenticationMessageClass()
	rv := objc.Send[SLSEventAuthenticationMessage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/initWithBasis:signature:
func NewSLSEventAuthenticationMessageWithBasisSignature(basis objectivec.IObject, signature objectivec.IObject) SLSEventAuthenticationMessage {
	instance := getSLSEventAuthenticationMessageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBasis:signature:"), basis, signature)
	return SLSEventAuthenticationMessageFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/initWithCoder:
func NewSLSEventAuthenticationMessageWithCoder(coder objectivec.IObject) SLSEventAuthenticationMessage {
	instance := getSLSEventAuthenticationMessageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSEventAuthenticationMessageFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/initWithEventRecord:pid:version:
func NewSLSEventAuthenticationMessageWithEventRecordPidVersion(record *SLSEventRecordRef, pid int, version uint32) SLSEventAuthenticationMessage {
	instance := getSLSEventAuthenticationMessageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEventRecord:pid:version:"), record, pid, version)
	return SLSEventAuthenticationMessageFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/initWithMessageInitData:
func NewSLSEventAuthenticationMessageWithMessageInitData(data unsafe.Pointer) SLSEventAuthenticationMessage {
	instance := getSLSEventAuthenticationMessageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMessageInitData:"), data)
	return SLSEventAuthenticationMessageFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/addToSigningContext:
func (s SLSEventAuthenticationMessage) AddToSigningContext(context objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("addToSigningContext:"), context)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/copySignedByKey:
func (s SLSEventAuthenticationMessage) CopySignedByKey(key objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("copySignedByKey:"), key)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/copyWithSignature:
func (s SLSEventAuthenticationMessage) CopyWithSignature(signature objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("copyWithSignature:"), signature)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/copyWithZone:signature:
func (s SLSEventAuthenticationMessage) CopyWithZoneSignature(zone NSZoneRef, signature objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("copyWithZone:signature:"), zone, signature)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/encodeWithCoder:
func (s SLSEventAuthenticationMessage) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/hasSuperclassEquivalentUnsignedData:
func (s SLSEventAuthenticationMessage) HasSuperclassEquivalentUnsignedData(data objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("hasSuperclassEquivalentUnsignedData:"), data)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/isSuperclassEquivalent:
func (s SLSEventAuthenticationMessage) IsSuperclassEquivalent(equivalent objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isSuperclassEquivalent:"), equivalent)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/matchesMessageData:
func (s SLSEventAuthenticationMessage) MatchesMessageData(data objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("matchesMessageData:"), data)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/validateWithOptions:andResultBlock:
func (s SLSEventAuthenticationMessage) ValidateWithOptionsAndResultBlock(options objectivec.IObject, block VoidHandler) {
	_block1, _ := NewVoidBlock(block)
	objc.Send[objc.ID](s.ID, objc.Sel("validateWithOptions:andResultBlock:"), options, _block1)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/initWithBasis:signature:
func (s SLSEventAuthenticationMessage) InitWithBasisSignature(basis objectivec.IObject, signature objectivec.IObject) SLSEventAuthenticationMessage {
	rv := objc.Send[SLSEventAuthenticationMessage](s.ID, objc.Sel("initWithBasis:signature:"), basis, signature)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/initWithCoder:
func (s SLSEventAuthenticationMessage) InitWithCoder(coder foundation.INSCoder) SLSEventAuthenticationMessage {
	rv := objc.Send[SLSEventAuthenticationMessage](s.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/initWithEventRecord:pid:version:
func (s SLSEventAuthenticationMessage) InitWithEventRecordPidVersion(record *SLSEventRecordRef, pid int, version uint32) SLSEventAuthenticationMessage {
	rv := objc.Send[SLSEventAuthenticationMessage](s.ID, objc.Sel("initWithEventRecord:pid:version:"), record, pid, version)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/initWithMessageInitData:
func (s SLSEventAuthenticationMessage) InitWithMessageInitData(data unsafe.Pointer) SLSEventAuthenticationMessage {
	rv := objc.Send[SLSEventAuthenticationMessage](s.ID, objc.Sel("initWithMessageInitData:"), data)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/basisMessage:signedByKey:
func (_SLSEventAuthenticationMessageClass SLSEventAuthenticationMessageClass) BasisMessageSignedByKey(message objectivec.IObject, key objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SLSEventAuthenticationMessageClass.class), objc.Sel("basisMessage:signedByKey:"), message, key)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/classForEventType:
func (_SLSEventAuthenticationMessageClass SLSEventAuthenticationMessageClass) ClassForEventType(type_ uint32) objc.Class {
	rv := objc.Send[objc.Class](objc.ID(_SLSEventAuthenticationMessageClass.class), objc.Sel("classForEventType:"), type_)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/messageWithEventRecord:pid:version:
func (_SLSEventAuthenticationMessageClass SLSEventAuthenticationMessageClass) MessageWithEventRecordPidVersion(record *SLSEventRecordRef, pid int, version uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SLSEventAuthenticationMessageClass.class), objc.Sel("messageWithEventRecord:pid:version:"), record, pid, version)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/supportsSecureCoding
func (_SLSEventAuthenticationMessageClass SLSEventAuthenticationMessageClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_SLSEventAuthenticationMessageClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/capabilities
func (s SLSEventAuthenticationMessage) Capabilities() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("capabilities"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/context
func (s SLSEventAuthenticationMessage) Context() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("context"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/debugDescription
func (s SLSEventAuthenticationMessage) DebugDescription() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/description
func (s SLSEventAuthenticationMessage) Description() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/eventType
func (s SLSEventAuthenticationMessage) EventType() ISLSEventAuthenticationMessageEventType {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("eventType"))
	return SLSEventAuthenticationMessageEventTypeFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/hash
func (s SLSEventAuthenticationMessage) Hash() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/originIdentifier
func (s SLSEventAuthenticationMessage) OriginIdentifier() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("originIdentifier"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/proxyTargetProcess
func (s SLSEventAuthenticationMessage) ProxyTargetProcess() ISLSEventAuthenticationMessageVersionedPID {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("proxyTargetProcess"))
	return SLSEventAuthenticationMessageVersionedPIDFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/signature
func (s SLSEventAuthenticationMessage) Signature() foundation.INSData {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("signature"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/superclass
func (s SLSEventAuthenticationMessage) Superclass() objc.Class {
	rv := objc.Send[objc.Class](s.ID, objc.Sel("superclass"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/targetProcess
func (s SLSEventAuthenticationMessage) TargetProcess() ISLSEventAuthenticationMessageVersionedPID {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("targetProcess"))
	return SLSEventAuthenticationMessageVersionedPIDFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/SkyLight/SLSEventAuthenticationMessage/timestamp
func (s SLSEventAuthenticationMessage) Timestamp() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("timestamp"))
	return rv
}

// ValidateWithOptionsAndResultBlockSync is a synchronous wrapper around [SLSEventAuthenticationMessage.ValidateWithOptionsAndResultBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (s SLSEventAuthenticationMessage) ValidateWithOptionsAndResultBlockSync(ctx context.Context, options objectivec.IObject) error {
	done := make(chan struct{}, 1)
	s.ValidateWithOptionsAndResultBlock(options, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
