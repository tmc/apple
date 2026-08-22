// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSMatrixDescriptor] class.
var (
	_MPSMatrixDescriptorClass     MPSMatrixDescriptorClass
	_MPSMatrixDescriptorClassOnce sync.Once
)

func getMPSMatrixDescriptorClass() MPSMatrixDescriptorClass {
	_MPSMatrixDescriptorClassOnce.Do(func() {
		_MPSMatrixDescriptorClass = MPSMatrixDescriptorClass{class: objc.GetClass("MPSMatrixDescriptor")}
	})
	return _MPSMatrixDescriptorClass
}

// GetMPSMatrixDescriptorClass returns the class object for MPSMatrixDescriptor.
func GetMPSMatrixDescriptorClass() MPSMatrixDescriptorClass {
	return getMPSMatrixDescriptorClass()
}

type MPSMatrixDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSMatrixDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSMatrixDescriptorClass) Alloc() MPSMatrixDescriptor {
	rv := objc.Send[MPSMatrixDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A description of attributes used to create an MPS matrix.
//
// # Overview
//
// Matrix data is assumed to be stored in row-major order.
//
// # Properties
//
//   - [MPSMatrixDescriptor.Rows]: The number of rows in the matrix.
//   - [MPSMatrixDescriptor.SetRows]
//   - [MPSMatrixDescriptor.Columns]: The number of columns in the matrix.
//   - [MPSMatrixDescriptor.SetColumns]
//   - [MPSMatrixDescriptor.DataType]: The type of the values in the matrix.
//   - [MPSMatrixDescriptor.SetDataType]
//   - [MPSMatrixDescriptor.RowBytes]: The stride, in bytes, between corresponding elements of consecutive rows in the matrix.
//   - [MPSMatrixDescriptor.SetRowBytes]
//   - [MPSMatrixDescriptor.Matrices]
//   - [MPSMatrixDescriptor.MatrixBytes]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixDescriptor
type MPSMatrixDescriptor struct {
	objectivec.Object
}

// MPSMatrixDescriptorFromID constructs a [MPSMatrixDescriptor] from an objc.ID.
//
// A description of attributes used to create an MPS matrix.
func MPSMatrixDescriptorFromID(id objc.ID) MPSMatrixDescriptor {
	return MPSMatrixDescriptor{objectivec.Object{ID: id}}
}

// NOTE: MPSMatrixDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSMatrixDescriptor] class.
//
// # Properties
//
//   - [IMPSMatrixDescriptor.Rows]: The number of rows in the matrix.
//   - [IMPSMatrixDescriptor.SetRows]
//   - [IMPSMatrixDescriptor.Columns]: The number of columns in the matrix.
//   - [IMPSMatrixDescriptor.SetColumns]
//   - [IMPSMatrixDescriptor.DataType]: The type of the values in the matrix.
//   - [IMPSMatrixDescriptor.SetDataType]
//   - [IMPSMatrixDescriptor.RowBytes]: The stride, in bytes, between corresponding elements of consecutive rows in the matrix.
//   - [IMPSMatrixDescriptor.SetRowBytes]
//   - [IMPSMatrixDescriptor.Matrices]
//   - [IMPSMatrixDescriptor.MatrixBytes]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixDescriptor
type IMPSMatrixDescriptor interface {
	objectivec.IObject

	// Topic: Properties

	// The number of rows in the matrix.
	Rows() uint
	SetRows(value uint)
	// The number of columns in the matrix.
	Columns() uint
	SetColumns(value uint)
	// The type of the values in the matrix.
	DataType() MPSDataType
	SetDataType(value MPSDataType)
	// The stride, in bytes, between corresponding elements of consecutive rows in the matrix.
	RowBytes() uint
	SetRowBytes(value uint)
	Matrices() uint
	MatrixBytes() uint
}

// Init initializes the instance.
func (m MPSMatrixDescriptor) Init() MPSMatrixDescriptor {
	rv := objc.Send[MPSMatrixDescriptor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MPSMatrixDescriptor) Autorelease() MPSMatrixDescriptor {
	rv := objc.Send[MPSMatrixDescriptor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSMatrixDescriptor creates a new MPSMatrixDescriptor instance.
func NewMPSMatrixDescriptor() MPSMatrixDescriptor {
	class := getMPSMatrixDescriptorClass()
	rv := objc.Send[MPSMatrixDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a matrix descriptor with the specified dimensions and data type.
//
// rows: The number of rows in the matrix.
//
// columns: The number of columns in the matrix.
//
// rowBytes: The stride, in bytes, between corresponding elements of consecutive rows in
// the matrix.
//
// dataType: The type of the data to be stored in the matrix.
//
// # Return Value
//
// A valid [MPSMatrixDescriptor] object.
//
// # Discussion
//
// For performance considerations, the optimal row stride may not necessarily
// be equal to the number of columns in the matrix. The
// [MPSMatrixDescriptorClass.RowBytesFromColumnsDataType] method may be used
// to help you determine this value.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixDescriptor/init(dimensions:columns:rowBytes:dataType:)
func NewMatrixDescriptorWithDimensionsColumnsRowBytesDataType(rows uint, columns uint, rowBytes uint, dataType MPSDataType) MPSMatrixDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getMPSMatrixDescriptorClass().class), objc.Sel("matrixDescriptorWithDimensions:columns:rowBytes:dataType:"), rows, columns, rowBytes, dataType)
	return MPSMatrixDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixDescriptor/init(rows:columns:matrices:rowBytes:matrixBytes:dataType:)
func NewMatrixDescriptorWithRowsColumnsMatricesRowBytesMatrixBytesDataType(rows uint, columns uint, matrices uint, rowBytes uint, matrixBytes uint, dataType MPSDataType) MPSMatrixDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getMPSMatrixDescriptorClass().class), objc.Sel("matrixDescriptorWithRows:columns:matrices:rowBytes:matrixBytes:dataType:"), rows, columns, matrices, rowBytes, matrixBytes, dataType)
	return MPSMatrixDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixDescriptor/init(rows:columns:rowBytes:dataType:)
func NewMatrixDescriptorWithRowsColumnsRowBytesDataType(rows uint, columns uint, rowBytes uint, dataType MPSDataType) MPSMatrixDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getMPSMatrixDescriptorClass().class), objc.Sel("matrixDescriptorWithRows:columns:rowBytes:dataType:"), rows, columns, rowBytes, dataType)
	return MPSMatrixDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixDescriptor/rowBytes(forColumns:dataType:)
func (_MPSMatrixDescriptorClass MPSMatrixDescriptorClass) RowBytesForColumnsDataType(columns uint, dataType MPSDataType) uintptr {
	rv := objc.Send[uintptr](objc.ID(_MPSMatrixDescriptorClass.class), objc.Sel("rowBytesForColumns:dataType:"), columns, dataType)
	return rv
}

// The number of rows in the matrix.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixDescriptor/rows
func (m MPSMatrixDescriptor) Rows() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("rows"))
	return rv
}
func (m MPSMatrixDescriptor) SetRows(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setRows:"), value)
}

// The number of columns in the matrix.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixDescriptor/columns
func (m MPSMatrixDescriptor) Columns() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("columns"))
	return rv
}
func (m MPSMatrixDescriptor) SetColumns(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setColumns:"), value)
}

// The type of the values in the matrix.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixDescriptor/dataType
func (m MPSMatrixDescriptor) DataType() MPSDataType {
	rv := objc.Send[MPSDataType](m.ID, objc.Sel("dataType"))
	return MPSDataType(rv)
}
func (m MPSMatrixDescriptor) SetDataType(value MPSDataType) {
	objc.Send[struct{}](m.ID, objc.Sel("setDataType:"), value)
}

// The stride, in bytes, between corresponding elements of consecutive rows in
// the matrix.
//
// # Discussion
//
// This value must be a multiple of the element size.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixDescriptor/rowBytes
func (m MPSMatrixDescriptor) RowBytes() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("rowBytes"))
	return rv
}
func (m MPSMatrixDescriptor) SetRowBytes(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setRowBytes:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixDescriptor/matrices
func (m MPSMatrixDescriptor) Matrices() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("matrices"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixDescriptor/matrixBytes
func (m MPSMatrixDescriptor) MatrixBytes() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("matrixBytes"))
	return rv
}
