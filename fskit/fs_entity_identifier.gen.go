// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [FSEntityIdentifier] class.
var (
	_FSEntityIdentifierClass     FSEntityIdentifierClass
	_FSEntityIdentifierClassOnce sync.Once
)

func getFSEntityIdentifierClass() FSEntityIdentifierClass {
	_FSEntityIdentifierClassOnce.Do(func() {
		_FSEntityIdentifierClass = FSEntityIdentifierClass{class: objc.GetClass("FSEntityIdentifier")}
	})
	return _FSEntityIdentifierClass
}

// GetFSEntityIdentifierClass returns the class object for FSEntityIdentifier.
func GetFSEntityIdentifierClass() FSEntityIdentifierClass {
	return getFSEntityIdentifierClass()
}

type FSEntityIdentifierClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (fc FSEntityIdentifierClass) Class() objc.Class {
	return fc.class
}

// Alloc allocates memory for a new instance of the class.
func (fc FSEntityIdentifierClass) Alloc() FSEntityIdentifier {
	rv := objc.Send[FSEntityIdentifier](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// A base type that identifies containers and volumes.
//
// # Overview
//
// An [FSEntityIdentifier] is a UUID to identify a container or volume,
// optionally with eight bytes of qualifying (differentiating) data. You use
// the qualifiers in cases in which a file server can receive multiple
// connections from the same client, which differ by user credentials. In this
// case, the identifier for each client is the server’s base UUID, and a
// unique qualifier that differs by client.
//
// # Creating an entity identifier
//
//   - [FSEntityIdentifier.InitWithUUID]: Creates an entity identifier with the given UUID.
//   - [FSEntityIdentifier.InitWithUUIDData]: Creates an entity identifier with the given UUID and qualifier data.
//   - [FSEntityIdentifier.InitWithUUIDQualifier]: Creates an entity identifier with the given UUID and qualifier data as a 64-bit unsigned integer.
//
// # Inspecting identifier properties
//
//   - [FSEntityIdentifier.Uuid]: A UUID to uniquely identify this entity.
//   - [FSEntityIdentifier.SetUuid]
//   - [FSEntityIdentifier.Qualifier]: An optional piece of data to distinguish entities that otherwise share the same UUID.
//   - [FSEntityIdentifier.SetQualifier]
//
// See: https://developer.apple.com/documentation/FSKit/FSEntityIdentifier
type FSEntityIdentifier struct {
	objectivec.Object
}

// FSEntityIdentifierFromID constructs a [FSEntityIdentifier] from an objc.ID.
//
// A base type that identifies containers and volumes.
func FSEntityIdentifierFromID(id objc.ID) FSEntityIdentifier {
	return FSEntityIdentifier{objectivec.Object{ID: id}}
}

// NOTE: FSEntityIdentifier adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [FSEntityIdentifier] class.
//
// # Creating an entity identifier
//
//   - [IFSEntityIdentifier.InitWithUUID]: Creates an entity identifier with the given UUID.
//   - [IFSEntityIdentifier.InitWithUUIDData]: Creates an entity identifier with the given UUID and qualifier data.
//   - [IFSEntityIdentifier.InitWithUUIDQualifier]: Creates an entity identifier with the given UUID and qualifier data as a 64-bit unsigned integer.
//
// # Inspecting identifier properties
//
//   - [IFSEntityIdentifier.Uuid]: A UUID to uniquely identify this entity.
//   - [IFSEntityIdentifier.SetUuid]
//   - [IFSEntityIdentifier.Qualifier]: An optional piece of data to distinguish entities that otherwise share the same UUID.
//   - [IFSEntityIdentifier.SetQualifier]
//
// See: https://developer.apple.com/documentation/FSKit/FSEntityIdentifier
type IFSEntityIdentifier interface {
	objectivec.IObject

	// Topic: Creating an entity identifier

	// Creates an entity identifier with the given UUID.
	InitWithUUID(uuid foundation.NSUUID) FSEntityIdentifier
	// Creates an entity identifier with the given UUID and qualifier data.
	InitWithUUIDData(uuid foundation.NSUUID, qualifierData foundation.NSData) FSEntityIdentifier
	// Creates an entity identifier with the given UUID and qualifier data as a 64-bit unsigned integer.
	InitWithUUIDQualifier(uuid foundation.NSUUID, qualifier uint64) FSEntityIdentifier

	// Topic: Inspecting identifier properties

	// A UUID to uniquely identify this entity.
	Uuid() foundation.NSUUID
	SetUuid(value foundation.NSUUID)
	// An optional piece of data to distinguish entities that otherwise share the same UUID.
	Qualifier() foundation.NSData
	SetQualifier(value foundation.NSData)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (e FSEntityIdentifier) Init() FSEntityIdentifier {
	rv := objc.Send[FSEntityIdentifier](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e FSEntityIdentifier) Autorelease() FSEntityIdentifier {
	rv := objc.Send[FSEntityIdentifier](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSEntityIdentifier creates a new FSEntityIdentifier instance.
func NewFSEntityIdentifier() FSEntityIdentifier {
	class := getFSEntityIdentifierClass()
	rv := objc.Send[FSEntityIdentifier](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an entity identifier with the given UUID.
//
// uuid: The UUID to use for this identifier.
//
// See: https://developer.apple.com/documentation/FSKit/FSEntityIdentifier/init(uuid:)-9e20k
func NewEntityIdentifierWithUUID(uuid foundation.NSUUID) FSEntityIdentifier {
	instance := getFSEntityIdentifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUUID:"), uuid)
	return FSEntityIdentifierFromID(rv)
}

// Creates an entity identifier with the given UUID and qualifier data.
//
// uuid: The UUID to use for this identifier.
//
// qualifierData: The data to distinguish entities that otherwise share the same UUID.
//
// See: https://developer.apple.com/documentation/FSKit/FSEntityIdentifier/init(uuid:data:)-8dixs
func NewEntityIdentifierWithUUIDData(uuid foundation.NSUUID, qualifierData foundation.NSData) FSEntityIdentifier {
	instance := getFSEntityIdentifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUUID:data:"), uuid, qualifierData)
	return FSEntityIdentifierFromID(rv)
}

// Creates an entity identifier with the given UUID and qualifier data as a
// 64-bit unsigned integer.
//
// uuid: The UUID to use for this identifier.
//
// qualifier: The data to distinguish entities that otherwise share the same UUID.
//
// See: https://developer.apple.com/documentation/FSKit/FSEntityIdentifier/init(uuid:qualifier:)-9ty70
func NewEntityIdentifierWithUUIDQualifier(uuid foundation.NSUUID, qualifier uint64) FSEntityIdentifier {
	instance := getFSEntityIdentifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUUID:qualifier:"), uuid, qualifier)
	return FSEntityIdentifierFromID(rv)
}

// Creates an entity identifier with the given UUID.
//
// uuid: The UUID to use for this identifier.
//
// See: https://developer.apple.com/documentation/FSKit/FSEntityIdentifier/init(uuid:)-9e20k
func (e FSEntityIdentifier) InitWithUUID(uuid foundation.NSUUID) FSEntityIdentifier {
	rv := objc.Send[FSEntityIdentifier](e.ID, objc.Sel("initWithUUID:"), uuid)
	return rv
}

// Creates an entity identifier with the given UUID and qualifier data.
//
// uuid: The UUID to use for this identifier.
//
// qualifierData: The data to distinguish entities that otherwise share the same UUID.
//
// See: https://developer.apple.com/documentation/FSKit/FSEntityIdentifier/init(uuid:data:)-8dixs
func (e FSEntityIdentifier) InitWithUUIDData(uuid foundation.NSUUID, qualifierData foundation.NSData) FSEntityIdentifier {
	rv := objc.Send[FSEntityIdentifier](e.ID, objc.Sel("initWithUUID:data:"), uuid, qualifierData)
	return rv
}

// Creates an entity identifier with the given UUID and qualifier data as a
// 64-bit unsigned integer.
//
// uuid: The UUID to use for this identifier.
//
// qualifier: The data to distinguish entities that otherwise share the same UUID.
//
// See: https://developer.apple.com/documentation/FSKit/FSEntityIdentifier/init(uuid:qualifier:)-9ty70
func (e FSEntityIdentifier) InitWithUUIDQualifier(uuid foundation.NSUUID, qualifier uint64) FSEntityIdentifier {
	rv := objc.Send[FSEntityIdentifier](e.ID, objc.Sel("initWithUUID:qualifier:"), uuid, qualifier)
	return rv
}
func (e FSEntityIdentifier) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](e.ID, objc.Sel("encodeWithCoder:"), coder)
}

// A UUID to uniquely identify this entity.
//
// See: https://developer.apple.com/documentation/FSKit/FSEntityIdentifier/uuid
func (e FSEntityIdentifier) Uuid() foundation.NSUUID {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("uuid"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}
func (e FSEntityIdentifier) SetUuid(value foundation.NSUUID) {
	objc.Send[struct{}](e.ID, objc.Sel("setUuid:"), value)
}

// An optional piece of data to distinguish entities that otherwise share the
// same UUID.
//
// See: https://developer.apple.com/documentation/FSKit/FSEntityIdentifier/qualifier
func (e FSEntityIdentifier) Qualifier() foundation.NSData {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("qualifier"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (e FSEntityIdentifier) SetQualifier(value foundation.NSData) {
	objc.Send[struct{}](e.ID, objc.Sel("setQualifier:"), value)
}
