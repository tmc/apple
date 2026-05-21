// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"context"
	"sync"

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
	Signature() foundation.NSData
	TargetProcess() ISLSEventAuthenticationMessageVersionedPID
	Timestamp() uint64
	ValidateWithOptionsAndResultBlock(options objectivec.IObject, block VoidHandler)
	InitWithBasisSignature(basis objectivec.IObject, signature objectivec.IObject) SLSEventAuthenticationMessage
	InitWithCoder(coder foundation.INSCoder) SLSEventAuthenticationMessage
	InitWithEventRecordPidVersion(record SLSEventRecord, pid int, version uint32) SLSEventAuthenticationMessage
	InitWithMessageInitData(data MessageInitData) SLSEventAuthenticationMessage
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
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

func NewSLSEventAuthenticationMessageWithBasisSignature(basis objectivec.IObject, signature objectivec.IObject) SLSEventAuthenticationMessage {
	instance := getSLSEventAuthenticationMessageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBasis:signature:"), basis, signature)
	return SLSEventAuthenticationMessageFromID(rv)
}

func NewSLSEventAuthenticationMessageWithCoder(coder objectivec.IObject) SLSEventAuthenticationMessage {
	instance := getSLSEventAuthenticationMessageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSEventAuthenticationMessageFromID(rv)
}

func NewSLSEventAuthenticationMessageWithEventRecordPidVersion(record SLSEventRecord, pid int, version uint32) SLSEventAuthenticationMessage {
	instance := getSLSEventAuthenticationMessageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEventRecord:pid:version:"), record, pid, version)
	return SLSEventAuthenticationMessageFromID(rv)
}

func NewSLSEventAuthenticationMessageWithMessageInitData(data MessageInitData) SLSEventAuthenticationMessage {
	instance := getSLSEventAuthenticationMessageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMessageInitData:"), data)
	return SLSEventAuthenticationMessageFromID(rv)
}

func (s SLSEventAuthenticationMessage) AddToSigningContext(context objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("addToSigningContext:"), context)
}
func (s SLSEventAuthenticationMessage) CopySignedByKey(key objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("copySignedByKey:"), key)
	return objectivec.Object{ID: rv}
}
func (s SLSEventAuthenticationMessage) CopyWithSignature(signature objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("copyWithSignature:"), signature)
	return objectivec.Object{ID: rv}
}
func (s SLSEventAuthenticationMessage) CopyWithZoneSignature(zone NSZoneRef, signature objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("copyWithZone:signature:"), zone, signature)
	return objectivec.Object{ID: rv}
}
func (s SLSEventAuthenticationMessage) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (s SLSEventAuthenticationMessage) HasSuperclassEquivalentUnsignedData(data objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("hasSuperclassEquivalentUnsignedData:"), data)
	return rv
}
func (s SLSEventAuthenticationMessage) IsSuperclassEquivalent(equivalent objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isSuperclassEquivalent:"), equivalent)
	return rv
}
func (s SLSEventAuthenticationMessage) MatchesMessageData(data objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("matchesMessageData:"), data)
	return rv
}
func (s SLSEventAuthenticationMessage) ValidateWithOptionsAndResultBlock(options objectivec.IObject, block VoidHandler) {
	_block1, _ := NewVoidBlock(block)
	objc.Send[objc.ID](s.ID, objc.Sel("validateWithOptions:andResultBlock:"), options, _block1)
}
func (s SLSEventAuthenticationMessage) InitWithBasisSignature(basis objectivec.IObject, signature objectivec.IObject) SLSEventAuthenticationMessage {
	rv := objc.Send[SLSEventAuthenticationMessage](s.ID, objc.Sel("initWithBasis:signature:"), basis, signature)
	return rv
}
func (s SLSEventAuthenticationMessage) InitWithCoder(coder foundation.INSCoder) SLSEventAuthenticationMessage {
	rv := objc.Send[SLSEventAuthenticationMessage](s.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (s SLSEventAuthenticationMessage) InitWithEventRecordPidVersion(record SLSEventRecord, pid int, version uint32) SLSEventAuthenticationMessage {
	rv := objc.Send[SLSEventAuthenticationMessage](s.ID, objc.Sel("initWithEventRecord:pid:version:"), record, pid, version)
	return rv
}
func (s SLSEventAuthenticationMessage) InitWithMessageInitData(data MessageInitData) SLSEventAuthenticationMessage {
	rv := objc.Send[SLSEventAuthenticationMessage](s.ID, objc.Sel("initWithMessageInitData:"), data)
	return rv
}

func (_SLSEventAuthenticationMessageClass SLSEventAuthenticationMessageClass) BasisMessageSignedByKey(message objectivec.IObject, key objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SLSEventAuthenticationMessageClass.class), objc.Sel("basisMessage:signedByKey:"), message, key)
	return objectivec.Object{ID: rv}
}
func (_SLSEventAuthenticationMessageClass SLSEventAuthenticationMessageClass) ClassForEventType(type_ uint32) objectivec.Class {
	rv := objc.Send[objectivec.Class](objc.ID(_SLSEventAuthenticationMessageClass.class), objc.Sel("classForEventType:"), type_)
	return objectivec.Class(rv)
}
func (_SLSEventAuthenticationMessageClass SLSEventAuthenticationMessageClass) MessageWithEventRecordPidVersion(record SLSEventRecord, pid int, version uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SLSEventAuthenticationMessageClass.class), objc.Sel("messageWithEventRecord:pid:version:"), record, pid, version)
	return objectivec.Object{ID: rv}
}
func (_SLSEventAuthenticationMessageClass SLSEventAuthenticationMessageClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_SLSEventAuthenticationMessageClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (s SLSEventAuthenticationMessage) Capabilities() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("capabilities"))
	return rv
}
func (s SLSEventAuthenticationMessage) Context() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("context"))
	return rv
}
func (s SLSEventAuthenticationMessage) DebugDescription() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (s SLSEventAuthenticationMessage) Description() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (s SLSEventAuthenticationMessage) EventType() ISLSEventAuthenticationMessageEventType {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("eventType"))
	return SLSEventAuthenticationMessageEventTypeFromID(objc.ID(rv))
}
func (s SLSEventAuthenticationMessage) Hash() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("hash"))
	return rv
}
func (s SLSEventAuthenticationMessage) OriginIdentifier() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("originIdentifier"))
	return rv
}
func (s SLSEventAuthenticationMessage) ProxyTargetProcess() ISLSEventAuthenticationMessageVersionedPID {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("proxyTargetProcess"))
	return SLSEventAuthenticationMessageVersionedPIDFromID(objc.ID(rv))
}
func (s SLSEventAuthenticationMessage) Signature() foundation.NSData {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("signature"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (s SLSEventAuthenticationMessage) Superclass() objectivec.Class {
	rv := objc.Send[objectivec.Class](s.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
func (s SLSEventAuthenticationMessage) TargetProcess() ISLSEventAuthenticationMessageVersionedPID {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("targetProcess"))
	return SLSEventAuthenticationMessageVersionedPIDFromID(objc.ID(rv))
}
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
