// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [FSContainerIdentifier] class.
var (
	_FSContainerIdentifierClass     FSContainerIdentifierClass
	_FSContainerIdentifierClassOnce sync.Once
)

func getFSContainerIdentifierClass() FSContainerIdentifierClass {
	_FSContainerIdentifierClassOnce.Do(func() {
		_FSContainerIdentifierClass = FSContainerIdentifierClass{class: objc.GetClass("FSContainerIdentifier")}
	})
	return _FSContainerIdentifierClass
}

// GetFSContainerIdentifierClass returns the class object for FSContainerIdentifier.
func GetFSContainerIdentifierClass() FSContainerIdentifierClass {
	return getFSContainerIdentifierClass()
}

type FSContainerIdentifierClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (fc FSContainerIdentifierClass) Class() objc.Class {
	return fc.class
}

// Alloc allocates memory for a new instance of the class.
func (fc FSContainerIdentifierClass) Alloc() FSContainerIdentifier {
	rv := objc.Send[FSContainerIdentifier](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// A type that identifies a container.
//
// # Overview
//
// The identifier is either a UUID or a UUID with additional differentiating
// bytes. Some network protocols evaluate access based on a user ID when
// connecting. In this situation, when a file server receives multiple client
// connections with different user IDs, the server provides different file
// hierarchies to each. For such systems, represent the container identifier
// as the UUID associated with the server, followed by four or eight bytes to
// differentiate connections.
//
// # Accessing identifier properties
//
//   - [FSContainerIdentifier.VolumeIdentifier]: The volume identifier associated with the container.
//
// See: https://developer.apple.com/documentation/FSKit/FSContainerIdentifier
type FSContainerIdentifier struct {
	FSEntityIdentifier
}

// FSContainerIdentifierFromID constructs a [FSContainerIdentifier] from an objc.ID.
//
// A type that identifies a container.
func FSContainerIdentifierFromID(id objc.ID) FSContainerIdentifier {
	return FSContainerIdentifier{FSEntityIdentifier: FSEntityIdentifierFromID(id)}
}

// NOTE: FSContainerIdentifier adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [FSContainerIdentifier] class.
//
// # Accessing identifier properties
//
//   - [IFSContainerIdentifier.VolumeIdentifier]: The volume identifier associated with the container.
//
// See: https://developer.apple.com/documentation/FSKit/FSContainerIdentifier
type IFSContainerIdentifier interface {
	IFSEntityIdentifier

	// Topic: Accessing identifier properties

	// The volume identifier associated with the container.
	VolumeIdentifier() IFSVolumeIdentifier
}

// Init initializes the instance.
func (c FSContainerIdentifier) Init() FSContainerIdentifier {
	rv := objc.Send[FSContainerIdentifier](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c FSContainerIdentifier) Autorelease() FSContainerIdentifier {
	rv := objc.Send[FSContainerIdentifier](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSContainerIdentifier creates a new FSContainerIdentifier instance.
func NewFSContainerIdentifier() FSContainerIdentifier {
	class := getFSContainerIdentifierClass()
	rv := objc.Send[FSContainerIdentifier](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/FSKit/FSEntityIdentifier/init(coder:)
func NewContainerIdentifierWithCoder(coder foundation.INSCoder) FSContainerIdentifier {
	instance := getFSContainerIdentifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return FSContainerIdentifierFromID(rv)
}

// Creates an entity identifier with the given UUID.
//
// uuid: The UUID to use for this identifier.
//
// See: https://developer.apple.com/documentation/FSKit/FSEntityIdentifier/init(uuid:)-9e20k
func NewContainerIdentifierWithUUID(uuid foundation.NSUUID) FSContainerIdentifier {
	instance := getFSContainerIdentifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUUID:"), uuid)
	return FSContainerIdentifierFromID(rv)
}

// Creates an entity identifier with the given UUID and qualifier data.
//
// uuid: The UUID to use for this identifier.
//
// qualifierData: The data to distinguish entities that otherwise share the same UUID.
//
// See: https://developer.apple.com/documentation/FSKit/FSEntityIdentifier/init(uuid:data:)-8dixs
func NewContainerIdentifierWithUUIDData(uuid foundation.NSUUID, qualifierData foundation.NSData) FSContainerIdentifier {
	instance := getFSContainerIdentifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUUID:data:"), uuid, qualifierData)
	return FSContainerIdentifierFromID(rv)
}

// Creates an entity identifier with the given UUID and qualifier data as a
// 64-bit unsigned integer.
//
// uuid: The UUID to use for this identifier.
//
// qualifier: The data to distinguish entities that otherwise share the same UUID.
//
// See: https://developer.apple.com/documentation/FSKit/FSEntityIdentifier/init(uuid:qualifier:)-9ty70
func NewContainerIdentifierWithUUIDQualifier(uuid foundation.NSUUID, qualifier uint64) FSContainerIdentifier {
	instance := getFSContainerIdentifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUUID:qualifier:"), uuid, qualifier)
	return FSContainerIdentifierFromID(rv)
}

// The volume identifier associated with the container.
//
// # Discussion
//
// For unary file systems, the volume identifier is the same as the container
// identifier.
//
// See: https://developer.apple.com/documentation/FSKit/FSContainerIdentifier/volumeIdentifier
func (c FSContainerIdentifier) VolumeIdentifier() IFSVolumeIdentifier {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("volumeIdentifier"))
	return FSVolumeIdentifierFromID(objc.ID(rv))
}
