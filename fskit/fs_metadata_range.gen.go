// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [FSMetadataRange] class.
var (
	_FSMetadataRangeClass     FSMetadataRangeClass
	_FSMetadataRangeClassOnce sync.Once
)

func getFSMetadataRangeClass() FSMetadataRangeClass {
	_FSMetadataRangeClassOnce.Do(func() {
		_FSMetadataRangeClass = FSMetadataRangeClass{class: objc.GetClass("FSMetadataRange")}
	})
	return _FSMetadataRangeClass
}

// GetFSMetadataRangeClass returns the class object for FSMetadataRange.
func GetFSMetadataRangeClass() FSMetadataRangeClass {
	return getFSMetadataRangeClass()
}

type FSMetadataRangeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (fc FSMetadataRangeClass) Class() objc.Class {
	return fc.class
}

// Alloc allocates memory for a new instance of the class.
func (fc FSMetadataRangeClass) Alloc() FSMetadataRange {
	rv := objc.Send[FSMetadataRange](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// A range that describes contiguous metadata segments on disk.
//
// # Overview
//
// This type represents a range that begins at `startOffset` and ends at
// `startOffset + segmentLength * segmentCount`. Each segment in the range
// represents a single block in the resource’s buffer cache.
//
// For example, given an [FSMetadataRange] with the following properties:
//
// - `startOffset = 0`
// - `segmentLength = 512`
// - `segmentCount = 8`
//
// The range represents eight segments: from 0 to 511, then from 512 to 1023,
// and so on until a final segment of 3584 to 4095.
//
// Ensure that each metadata segment represents a range that’s already
// present in the resource’s buffer cache. Similarly, ensure that each
// segment’s offset and length matches the offset and length of the
// corresponding block in the buffer cache.
//
// # Creating a metadata range
//
//   - [FSMetadataRange.InitWithOffsetSegmentLengthSegmentCount]: Initializes a metadata range with the given properties.
//
// # Accessing range properties
//
//   - [FSMetadataRange.StartOffset]: The start offset of the range in bytes.
//   - [FSMetadataRange.SegmentLength]: The segment length in bytes.
//   - [FSMetadataRange.SegmentCount]: The number of segments in the range.
//
// See: https://developer.apple.com/documentation/FSKit/FSMetadataRange
type FSMetadataRange struct {
	objectivec.Object
}

// FSMetadataRangeFromID constructs a [FSMetadataRange] from an objc.ID.
//
// A range that describes contiguous metadata segments on disk.
func FSMetadataRangeFromID(id objc.ID) FSMetadataRange {
	return FSMetadataRange{objectivec.Object{ID: id}}
}

// NOTE: FSMetadataRange adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [FSMetadataRange] class.
//
// # Creating a metadata range
//
//   - [IFSMetadataRange.InitWithOffsetSegmentLengthSegmentCount]: Initializes a metadata range with the given properties.
//
// # Accessing range properties
//
//   - [IFSMetadataRange.StartOffset]: The start offset of the range in bytes.
//   - [IFSMetadataRange.SegmentLength]: The segment length in bytes.
//   - [IFSMetadataRange.SegmentCount]: The number of segments in the range.
//
// See: https://developer.apple.com/documentation/FSKit/FSMetadataRange
type IFSMetadataRange interface {
	objectivec.IObject

	// Topic: Creating a metadata range

	// Initializes a metadata range with the given properties.
	InitWithOffsetSegmentLengthSegmentCount(startOffset int64, segmentLength uint64, segmentCount uint64) FSMetadataRange

	// Topic: Accessing range properties

	// The start offset of the range in bytes.
	StartOffset() int64
	// The segment length in bytes.
	SegmentLength() uint64
	// The number of segments in the range.
	SegmentCount() uint64
}

// Init initializes the instance.
func (m FSMetadataRange) Init() FSMetadataRange {
	rv := objc.Send[FSMetadataRange](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m FSMetadataRange) Autorelease() FSMetadataRange {
	rv := objc.Send[FSMetadataRange](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSMetadataRange creates a new FSMetadataRange instance.
func NewFSMetadataRange() FSMetadataRange {
	class := getFSMetadataRangeClass()
	rv := objc.Send[FSMetadataRange](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a metadata range with the given properties.
//
// startOffset: The start offset of the range in bytes. Ensure this value is a multiple of
// the corresponding resource’s [FSBlockDeviceResource.BlockSize].
//
// segmentLength: The segment length in bytes. Ensure this value is a multiple of the
// corresponding resource’s [FSBlockDeviceResource.BlockSize].
//
// segmentCount: The number of segments in the range.
//
// See: https://developer.apple.com/documentation/FSKit/FSMetadataRange/init(offset:segmentLength:segmentCount:)
func NewMetadataRangeWithOffsetSegmentLengthSegmentCount(startOffset int64, segmentLength uint64, segmentCount uint64) FSMetadataRange {
	instance := getFSMetadataRangeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithOffset:segmentLength:segmentCount:"), startOffset, segmentLength, segmentCount)
	return FSMetadataRangeFromID(rv)
}

// Initializes a metadata range with the given properties.
//
// startOffset: The start offset of the range in bytes. Ensure this value is a multiple of
// the corresponding resource’s [FSBlockDeviceResource.BlockSize].
//
// segmentLength: The segment length in bytes. Ensure this value is a multiple of the
// corresponding resource’s [FSBlockDeviceResource.BlockSize].
//
// segmentCount: The number of segments in the range.
//
// See: https://developer.apple.com/documentation/FSKit/FSMetadataRange/init(offset:segmentLength:segmentCount:)
func (m FSMetadataRange) InitWithOffsetSegmentLengthSegmentCount(startOffset int64, segmentLength uint64, segmentCount uint64) FSMetadataRange {
	rv := objc.Send[FSMetadataRange](m.ID, objc.Sel("initWithOffset:segmentLength:segmentCount:"), startOffset, segmentLength, segmentCount)
	return rv
}

// Creates a metadata range with the given properties.
//
// startOffset: The start offset of the range in bytes. Ensure this value is a multiple of
// the corresponding resource’s [FSBlockDeviceResource.BlockSize].
//
// segmentLength: The segment length in bytes. Ensure this value is a multiple of the
// corresponding resource’s [FSBlockDeviceResource.BlockSize].
//
// segmentCount: The number of segments in the range.
//
// See: https://developer.apple.com/documentation/FSKit/FSMetadataRange/rangeWithOffset:segmentLength:segmentCount:
func (_FSMetadataRangeClass FSMetadataRangeClass) RangeWithOffsetSegmentLengthSegmentCount(startOffset int64, segmentLength uint64, segmentCount uint64) FSMetadataRange {
	rv := objc.Send[objc.ID](objc.ID(_FSMetadataRangeClass.class), objc.Sel("rangeWithOffset:segmentLength:segmentCount:"), startOffset, segmentLength, segmentCount)
	return FSMetadataRangeFromID(rv)
}

// The start offset of the range in bytes.
//
// # Discussion
//
// Ensure this value is a multiple of the corresponding resource’s
// [FSBlockDeviceResource.BlockSize].
//
// See: https://developer.apple.com/documentation/FSKit/FSMetadataRange/startOffset
func (m FSMetadataRange) StartOffset() int64 {
	rv := objc.Send[int64](m.ID, objc.Sel("startOffset"))
	return rv
}

// The segment length in bytes.
//
// # Discussion
//
// Ensure this value is a multiple of the corresponding resource’s
// [FSBlockDeviceResource.BlockSize].
//
// See: https://developer.apple.com/documentation/FSKit/FSMetadataRange/segmentLength
func (m FSMetadataRange) SegmentLength() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("segmentLength"))
	return rv
}

// The number of segments in the range.
//
// See: https://developer.apple.com/documentation/FSKit/FSMetadataRange/segmentCount
func (m FSMetadataRange) SegmentCount() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("segmentCount"))
	return rv
}
