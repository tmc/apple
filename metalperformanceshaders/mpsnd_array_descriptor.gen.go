// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNDArrayDescriptor] class.
var (
	_MPSNDArrayDescriptorClass     MPSNDArrayDescriptorClass
	_MPSNDArrayDescriptorClassOnce sync.Once
)

func getMPSNDArrayDescriptorClass() MPSNDArrayDescriptorClass {
	_MPSNDArrayDescriptorClassOnce.Do(func() {
		_MPSNDArrayDescriptorClass = MPSNDArrayDescriptorClass{class: objc.GetClass("MPSNDArrayDescriptor")}
	})
	return _MPSNDArrayDescriptorClass
}

// GetMPSNDArrayDescriptorClass returns the class object for MPSNDArrayDescriptor.
func GetMPSNDArrayDescriptorClass() MPSNDArrayDescriptorClass {
	return getMPSNDArrayDescriptorClass()
}

type MPSNDArrayDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNDArrayDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNDArrayDescriptorClass) Alloc() MPSNDArrayDescriptor {
	rv := objc.Send[MPSNDArrayDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [MPSNDArrayDescriptor.DataType]
//   - [MPSNDArrayDescriptor.SetDataType]
//   - [MPSNDArrayDescriptor.NumberOfDimensions]
//   - [MPSNDArrayDescriptor.SetNumberOfDimensions]
//   - [MPSNDArrayDescriptor.PreferPackedRows]
//   - [MPSNDArrayDescriptor.SetPreferPackedRows]
//
// # Instance Methods
//
//   - [MPSNDArrayDescriptor.DimensionOrder]
//   - [MPSNDArrayDescriptor.GetShape]
//   - [MPSNDArrayDescriptor.LengthOfDimension]
//   - [MPSNDArrayDescriptor.PermuteWithDimensionOrder]
//   - [MPSNDArrayDescriptor.ReshapeWithDimensionCountDimensionSizes]
//   - [MPSNDArrayDescriptor.ReshapeWithShape]
//   - [MPSNDArrayDescriptor.SliceDimensionWithSubrange]
//   - [MPSNDArrayDescriptor.SliceRangeForDimension]
//   - [MPSNDArrayDescriptor.TransposeDimensionWithDimension]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayDescriptor
type MPSNDArrayDescriptor struct {
	objectivec.Object
}

// MPSNDArrayDescriptorFromID constructs a [MPSNDArrayDescriptor] from an objc.ID.
func MPSNDArrayDescriptorFromID(id objc.ID) MPSNDArrayDescriptor {
	return MPSNDArrayDescriptor{objectivec.Object{ID: id}}
}

// NOTE: MPSNDArrayDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNDArrayDescriptor] class.
//
// # Instance Properties
//
//   - [IMPSNDArrayDescriptor.DataType]
//   - [IMPSNDArrayDescriptor.SetDataType]
//   - [IMPSNDArrayDescriptor.NumberOfDimensions]
//   - [IMPSNDArrayDescriptor.SetNumberOfDimensions]
//   - [IMPSNDArrayDescriptor.PreferPackedRows]
//   - [IMPSNDArrayDescriptor.SetPreferPackedRows]
//
// # Instance Methods
//
//   - [IMPSNDArrayDescriptor.DimensionOrder]
//   - [IMPSNDArrayDescriptor.GetShape]
//   - [IMPSNDArrayDescriptor.LengthOfDimension]
//   - [IMPSNDArrayDescriptor.PermuteWithDimensionOrder]
//   - [IMPSNDArrayDescriptor.ReshapeWithDimensionCountDimensionSizes]
//   - [IMPSNDArrayDescriptor.ReshapeWithShape]
//   - [IMPSNDArrayDescriptor.SliceDimensionWithSubrange]
//   - [IMPSNDArrayDescriptor.SliceRangeForDimension]
//   - [IMPSNDArrayDescriptor.TransposeDimensionWithDimension]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayDescriptor
type IMPSNDArrayDescriptor interface {
	objectivec.IObject

	// Topic: Instance Properties

	DataType() MPSDataType
	SetDataType(value MPSDataType)
	NumberOfDimensions() uint
	SetNumberOfDimensions(value uint)
	PreferPackedRows() bool
	SetPreferPackedRows(value bool)

	// Topic: Instance Methods

	DimensionOrder() uint8
	GetShape() []foundation.NSNumber
	LengthOfDimension(dimensionIndex uint) uint
	PermuteWithDimensionOrder(dimensionOrder *uint)
	ReshapeWithDimensionCountDimensionSizes(numberOfDimensions uint, dimensionSizes *uint)
	ReshapeWithShape(shape []foundation.NSNumber)
	SliceDimensionWithSubrange(dimensionIndex uint, subRange MPSDimensionSlice)
	SliceRangeForDimension(dimensionIndex uint) MPSDimensionSlice
	TransposeDimensionWithDimension(dimensionIndex uint, dimensionIndex2 uint)
}

// Init initializes the instance.
func (n MPSNDArrayDescriptor) Init() MPSNDArrayDescriptor {
	rv := objc.Send[MPSNDArrayDescriptor](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MPSNDArrayDescriptor) Autorelease() MPSNDArrayDescriptor {
	rv := objc.Send[MPSNDArrayDescriptor](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNDArrayDescriptor creates a new MPSNDArrayDescriptor instance.
func NewMPSNDArrayDescriptor() MPSNDArrayDescriptor {
	class := getMPSNDArrayDescriptorClass()
	rv := objc.Send[MPSNDArrayDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayDescriptor/init(dataType:dimensionCount:dimensionSizes:)
func NewNDArrayDescriptorWithDataTypeDimensionCountDimensionSizes(dataType MPSDataType, numberOfDimensions uint, dimensionSizes *uint) MPSNDArrayDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getMPSNDArrayDescriptorClass().class), objc.Sel("descriptorWithDataType:dimensionCount:dimensionSizes:"), dataType, numberOfDimensions, unsafe.Pointer(dimensionSizes))
	return MPSNDArrayDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayDescriptor/init(dataType:shape:)
func NewNDArrayDescriptorWithDataTypeShape(dataType MPSDataType, shape []foundation.NSNumber) MPSNDArrayDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getMPSNDArrayDescriptorClass().class), objc.Sel("descriptorWithDataType:shape:"), dataType, objectivec.IObjectSliceToNSArray(shape))
	return MPSNDArrayDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayDescriptor/dimensionOrder()
func (n MPSNDArrayDescriptor) DimensionOrder() uint8 {
	rv := objc.Send[uint8](n.ID, objc.Sel("dimensionOrder"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayDescriptor/getShape()
func (n MPSNDArrayDescriptor) GetShape() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](n.ID, objc.Sel("getShape"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayDescriptor/length(ofDimension:)
func (n MPSNDArrayDescriptor) LengthOfDimension(dimensionIndex uint) uint {
	rv := objc.Send[uint](n.ID, objc.Sel("lengthOfDimension:"), dimensionIndex)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayDescriptor/permute(withDimensionOrder:)
func (n MPSNDArrayDescriptor) PermuteWithDimensionOrder(dimensionOrder *uint) {
	objc.Send[objc.ID](n.ID, objc.Sel("permuteWithDimensionOrder:"), unsafe.Pointer(dimensionOrder))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayDescriptor/reshape(withDimensionCount:dimensionSizes:)
func (n MPSNDArrayDescriptor) ReshapeWithDimensionCountDimensionSizes(numberOfDimensions uint, dimensionSizes *uint) {
	objc.Send[objc.ID](n.ID, objc.Sel("reshapeWithDimensionCount:dimensionSizes:"), numberOfDimensions, unsafe.Pointer(dimensionSizes))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayDescriptor/reshape(withShape:)
func (n MPSNDArrayDescriptor) ReshapeWithShape(shape []foundation.NSNumber) {
	objc.Send[objc.ID](n.ID, objc.Sel("reshapeWithShape:"), objectivec.IObjectSliceToNSArray(shape))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayDescriptor/sliceDimension(_:withSubrange:)
func (n MPSNDArrayDescriptor) SliceDimensionWithSubrange(dimensionIndex uint, subRange MPSDimensionSlice) {
	objc.Send[objc.ID](n.ID, objc.Sel("sliceDimension:withSubrange:"), dimensionIndex, subRange)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayDescriptor/sliceRange(forDimension:)
func (n MPSNDArrayDescriptor) SliceRangeForDimension(dimensionIndex uint) MPSDimensionSlice {
	rv := objc.Send[MPSDimensionSlice](n.ID, objc.Sel("sliceRangeForDimension:"), dimensionIndex)
	return MPSDimensionSlice(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayDescriptor/transposeDimension(_:withDimension:)
func (n MPSNDArrayDescriptor) TransposeDimensionWithDimension(dimensionIndex uint, dimensionIndex2 uint) {
	objc.Send[objc.ID](n.ID, objc.Sel("transposeDimension:withDimension:"), dimensionIndex, dimensionIndex2)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayDescriptor/descriptorWithDataType:dimensionSizes:
func (_MPSNDArrayDescriptorClass MPSNDArrayDescriptorClass) DescriptorWithDataTypeDimensionSizes(dataType MPSDataType, dimension0 uint) MPSNDArrayDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_MPSNDArrayDescriptorClass.class), objc.Sel("descriptorWithDataType:dimensionSizes:"), dataType, dimension0)
	return MPSNDArrayDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayDescriptor/dataType
func (n MPSNDArrayDescriptor) DataType() MPSDataType {
	rv := objc.Send[MPSDataType](n.ID, objc.Sel("dataType"))
	return MPSDataType(rv)
}
func (n MPSNDArrayDescriptor) SetDataType(value MPSDataType) {
	objc.Send[struct{}](n.ID, objc.Sel("setDataType:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayDescriptor/numberOfDimensions
func (n MPSNDArrayDescriptor) NumberOfDimensions() uint {
	rv := objc.Send[uint](n.ID, objc.Sel("numberOfDimensions"))
	return rv
}
func (n MPSNDArrayDescriptor) SetNumberOfDimensions(value uint) {
	objc.Send[struct{}](n.ID, objc.Sel("setNumberOfDimensions:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayDescriptor/preferPackedRows
func (n MPSNDArrayDescriptor) PreferPackedRows() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("preferPackedRows"))
	return rv
}
func (n MPSNDArrayDescriptor) SetPreferPackedRows(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setPreferPackedRows:"), value)
}
