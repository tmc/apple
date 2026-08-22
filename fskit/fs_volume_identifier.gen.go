// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [FSVolumeIdentifier] class.
var (
	_FSVolumeIdentifierClass     FSVolumeIdentifierClass
	_FSVolumeIdentifierClassOnce sync.Once
)

func getFSVolumeIdentifierClass() FSVolumeIdentifierClass {
	_FSVolumeIdentifierClassOnce.Do(func() {
		_FSVolumeIdentifierClass = FSVolumeIdentifierClass{class: objc.GetClass("FSVolumeIdentifier")}
	})
	return _FSVolumeIdentifierClass
}

// GetFSVolumeIdentifierClass returns the class object for FSVolumeIdentifier.
func GetFSVolumeIdentifierClass() FSVolumeIdentifierClass {
	return getFSVolumeIdentifierClass()
}

type FSVolumeIdentifierClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (fc FSVolumeIdentifierClass) Class() objc.Class {
	return fc.class
}

// Alloc allocates memory for a new instance of the class.
func (fc FSVolumeIdentifierClass) Alloc() FSVolumeIdentifier {
	rv := objc.Send[FSVolumeIdentifier](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// A type that identifies a volume.
//
// # Overview
//
// For most volumes, the volume identifier is the UUID identifying the volume.
//
// Network file systems may access the same underlying volume using different
// authentication credentials. To handle this situation, add qualifying data
// to identify the specific container, as discussed in the superclass,
// [FSEntityIdentifier].
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/Identifier
type FSVolumeIdentifier struct {
	FSEntityIdentifier
}

// FSVolumeIdentifierFromID constructs a [FSVolumeIdentifier] from an objc.ID.
//
// A type that identifies a volume.
func FSVolumeIdentifierFromID(id objc.ID) FSVolumeIdentifier {
	return FSVolumeIdentifier{FSEntityIdentifier: FSEntityIdentifierFromID(id)}
}

// NOTE: FSVolumeIdentifier adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [FSVolumeIdentifier] class.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/Identifier
type IFSVolumeIdentifier interface {
	IFSEntityIdentifier
}

// Init initializes the instance.
func (v FSVolumeIdentifier) Init() FSVolumeIdentifier {
	rv := objc.Send[FSVolumeIdentifier](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v FSVolumeIdentifier) Autorelease() FSVolumeIdentifier {
	rv := objc.Send[FSVolumeIdentifier](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSVolumeIdentifier creates a new FSVolumeIdentifier instance.
func NewFSVolumeIdentifier() FSVolumeIdentifier {
	class := getFSVolumeIdentifierClass()
	rv := objc.Send[FSVolumeIdentifier](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/FSKit/FSEntityIdentifier/init(coder:)
func NewVolumeIdentifierWithCoder(coder foundation.INSCoder) FSVolumeIdentifier {
	instance := getFSVolumeIdentifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return FSVolumeIdentifierFromID(rv)
}

// Creates an entity identifier with the given UUID.
//
// uuid: The UUID to use for this identifier.
//
// See: https://developer.apple.com/documentation/FSKit/FSEntityIdentifier/init(uuid:)
func NewVolumeIdentifierWithUUID(uuid foundation.NSUUID) FSVolumeIdentifier {
	instance := getFSVolumeIdentifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUUID:"), uuid)
	return FSVolumeIdentifierFromID(rv)
}

// Creates an entity identifier with the given UUID and qualifier data.
//
// uuid: The UUID to use for this identifier.
//
// qualifierData: The data to distinguish entities that otherwise share the same UUID.
//
// See: https://developer.apple.com/documentation/FSKit/FSEntityIdentifier/init(uuid:data:)
func NewVolumeIdentifierWithUUIDData(uuid foundation.NSUUID, qualifierData foundation.NSData) FSVolumeIdentifier {
	instance := getFSVolumeIdentifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUUID:data:"), uuid, qualifierData)
	return FSVolumeIdentifierFromID(rv)
}

// Creates an entity identifier with the given UUID and qualifier data as a
// 64-bit unsigned integer.
//
// uuid: The UUID to use for this identifier.
//
// qualifier: The data to distinguish entities that otherwise share the same UUID.
//
// See: https://developer.apple.com/documentation/FSKit/FSEntityIdentifier/init(uuid:qualifier:)
func NewVolumeIdentifierWithUUIDQualifier(uuid foundation.NSUUID, qualifier uint64) FSVolumeIdentifier {
	instance := getFSVolumeIdentifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUUID:qualifier:"), uuid, qualifier)
	return FSVolumeIdentifierFromID(rv)
}
