// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [FSExtentPacker] class.
var (
	_FSExtentPackerClass     FSExtentPackerClass
	_FSExtentPackerClassOnce sync.Once
)

func getFSExtentPackerClass() FSExtentPackerClass {
	_FSExtentPackerClassOnce.Do(func() {
		_FSExtentPackerClass = FSExtentPackerClass{class: objc.GetClass("FSExtentPacker")}
	})
	return _FSExtentPackerClass
}

// GetFSExtentPackerClass returns the class object for FSExtentPacker.
func GetFSExtentPackerClass() FSExtentPackerClass {
	return getFSExtentPackerClass()
}

type FSExtentPackerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (fc FSExtentPackerClass) Class() objc.Class {
	return fc.class
}

// Alloc allocates memory for a new instance of the class.
func (fc FSExtentPackerClass) Alloc() FSExtentPacker {
	rv := objc.Send[FSExtentPacker](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// A type that directs the kernel to map space on disk to a specific file
// managed by this file system.
//
// # Overview
//
// provide the kernel the logical-to-physical mapping of a given file. An
// extent describes a physical offset on disk, and a length and a logical
// offset within the file. Rather than working with extents directly, you use
// this type’s methods to provide or “pack” extent information, which
// FSKit then passes to the kernel.
//
// # Packing extents
//
//   - [FSExtentPacker.PackExtentWithResourceTypeLogicalOffsetPhysicalOffsetLength]: Packs a single extent to send to the kernel.
//
// See: https://developer.apple.com/documentation/FSKit/FSExtentPacker
type FSExtentPacker struct {
	objectivec.Object
}

// FSExtentPackerFromID constructs a [FSExtentPacker] from an objc.ID.
//
// A type that directs the kernel to map space on disk to a specific file
// managed by this file system.
func FSExtentPackerFromID(id objc.ID) FSExtentPacker {
	return FSExtentPacker{objectivec.Object{ID: id}}
}

// NOTE: FSExtentPacker adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [FSExtentPacker] class.
//
// # Packing extents
//
//   - [IFSExtentPacker.PackExtentWithResourceTypeLogicalOffsetPhysicalOffsetLength]: Packs a single extent to send to the kernel.
//
// See: https://developer.apple.com/documentation/FSKit/FSExtentPacker
type IFSExtentPacker interface {
	objectivec.IObject

	// Topic: Packing extents

	// Packs a single extent to send to the kernel.
	PackExtentWithResourceTypeLogicalOffsetPhysicalOffsetLength(resource IFSBlockDeviceResource, type_ FSExtentType, logicalOffset int64, physicalOffset int64, length uintptr) bool
}

// Init initializes the instance.
func (e FSExtentPacker) Init() FSExtentPacker {
	rv := objc.Send[FSExtentPacker](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e FSExtentPacker) Autorelease() FSExtentPacker {
	rv := objc.Send[FSExtentPacker](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSExtentPacker creates a new FSExtentPacker instance.
func NewFSExtentPacker() FSExtentPacker {
	class := getFSExtentPackerClass()
	rv := objc.Send[FSExtentPacker](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Packs a single extent to send to the kernel.
//
// resource: The resource on which to perform I/O.
//
// type: The type of extent, indicating whether it contains valid data.
//
// logicalOffset: The extent offset within the file, in bytes.
//
// physicalOffset: The extent offset on disk, in bytes.
//
// length: The extent length, in bytes.
//
// # Return Value
//
// A Boolean value that indicates whether the packer can pack more extents.
//
// See: https://developer.apple.com/documentation/FSKit/FSExtentPacker/packExtent(resource:type:logicalOffset:physicalOffset:length:)
func (e FSExtentPacker) PackExtentWithResourceTypeLogicalOffsetPhysicalOffsetLength(resource IFSBlockDeviceResource, type_ FSExtentType, logicalOffset int64, physicalOffset int64, length uintptr) bool {
	rv := objc.Send[bool](e.ID, objc.Sel("packExtentWithResource:type:logicalOffset:physicalOffset:length:"), resource, type_, logicalOffset, physicalOffset, length)
	return rv
}
