// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Properties implemented by volumes that support providing the values of system limits or options.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/PathConfOperations
type FSVolumePathConfOperations interface {
	objectivec.IObject

	// A property that represents the maximum number of hard links to the object.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/PathConfOperations/maximumLinkCount
	MaximumLinkCount() int

	// A property that represents the maximum length of a component of a filename.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/PathConfOperations/maximumNameLength
	MaximumNameLength() int

	// A Boolean property that indicates whether the volume restricts ownership changes based on authorization.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/PathConfOperations/restrictsOwnershipChanges
	RestrictsOwnershipChanges() bool

	// A property that indicates whether the volume truncates files longer than its maximum supported length.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/PathConfOperations/truncatesLongNames
	TruncatesLongNames() bool

	// The maximum size of a regular file allowed in the volume.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/PathConfOperations/maximumFileSize
	MaximumFileSize() uint64

	// The minimum number of bits needed to represent, as a signed integer value, the maximum size of a regular file allowed in the volume.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/PathConfOperations/maximumFileSizeInBits
	MaximumFileSizeInBits() int

	// The maximum extended attribute size in bytes.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/PathConfOperations/maximumXattrSize
	MaximumXattrSize() int

	// The maximum extended attribute size in bits.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/PathConfOperations/maximumXattrSizeInBits
	MaximumXattrSizeInBits() int
}

// FSVolumePathConfOperationsObject wraps an existing Objective-C object that conforms to the FSVolumePathConfOperations protocol.
type FSVolumePathConfOperationsObject struct {
	objectivec.Object
}

func (o FSVolumePathConfOperationsObject) BaseObject() objectivec.Object {
	return o.Object
}

// FSVolumePathConfOperationsObjectFromID constructs a [FSVolumePathConfOperationsObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func FSVolumePathConfOperationsObjectFromID(id objc.ID) FSVolumePathConfOperationsObject {
	return FSVolumePathConfOperationsObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A property that represents the maximum number of hard links to the object.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/PathConfOperations/maximumLinkCount
func (o FSVolumePathConfOperationsObject) MaximumLinkCount() int {
	rv := objc.Send[int](o.ID, objc.Sel("maximumLinkCount"))
	return rv
}

// A property that represents the maximum length of a component of a filename.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/PathConfOperations/maximumNameLength
func (o FSVolumePathConfOperationsObject) MaximumNameLength() int {
	rv := objc.Send[int](o.ID, objc.Sel("maximumNameLength"))
	return rv
}

// A Boolean property that indicates whether the volume restricts ownership
// changes based on authorization.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/PathConfOperations/restrictsOwnershipChanges
func (o FSVolumePathConfOperationsObject) RestrictsOwnershipChanges() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("restrictsOwnershipChanges"))
	return rv
}

// A property that indicates whether the volume truncates files longer than
// its maximum supported length.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/PathConfOperations/truncatesLongNames
func (o FSVolumePathConfOperationsObject) TruncatesLongNames() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("truncatesLongNames"))
	return rv
}

// The maximum size of a regular file allowed in the volume.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/PathConfOperations/maximumFileSize
func (o FSVolumePathConfOperationsObject) MaximumFileSize() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("maximumFileSize"))
	return rv
}

// The minimum number of bits needed to represent, as a signed integer value,
// the maximum size of a regular file allowed in the volume.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/PathConfOperations/maximumFileSizeInBits
func (o FSVolumePathConfOperationsObject) MaximumFileSizeInBits() int {
	rv := objc.Send[int](o.ID, objc.Sel("maximumFileSizeInBits"))
	return rv
}

// The maximum extended attribute size in bytes.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/PathConfOperations/maximumXattrSize
func (o FSVolumePathConfOperationsObject) MaximumXattrSize() int {
	rv := objc.Send[int](o.ID, objc.Sel("maximumXattrSize"))
	return rv
}

// The maximum extended attribute size in bits.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/PathConfOperations/maximumXattrSizeInBits
func (o FSVolumePathConfOperationsObject) MaximumXattrSizeInBits() int {
	rv := objc.Send[int](o.ID, objc.Sel("maximumXattrSizeInBits"))
	return rv
}
