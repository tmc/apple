// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTShaderProfilerMCABinary] class.
var (
	_GTShaderProfilerMCABinaryClass     GTShaderProfilerMCABinaryClass
	_GTShaderProfilerMCABinaryClassOnce sync.Once
)

func getGTShaderProfilerMCABinaryClass() GTShaderProfilerMCABinaryClass {
	_GTShaderProfilerMCABinaryClassOnce.Do(func() {
		_GTShaderProfilerMCABinaryClass = GTShaderProfilerMCABinaryClass{class: objc.GetClass("GTShaderProfilerMCABinary")}
	})
	return _GTShaderProfilerMCABinaryClass
}

// GetGTShaderProfilerMCABinaryClass returns the class object for GTShaderProfilerMCABinary.
func GetGTShaderProfilerMCABinaryClass() GTShaderProfilerMCABinaryClass {
	return getGTShaderProfilerMCABinaryClass()
}

type GTShaderProfilerMCABinaryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTShaderProfilerMCABinaryClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTShaderProfilerMCABinaryClass) Alloc() GTShaderProfilerMCABinary {
	rv := objc.Send[GTShaderProfilerMCABinary](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTShaderProfilerMCABinary.AllocatedGPRCount]
//   - [GTShaderProfilerMCABinary.GenerateAGX2Assembly]
//   - [GTShaderProfilerMCABinary.GenerateAPSAssembly]
//   - [GTShaderProfilerMCABinary.GenerateAssemblyContent]
//   - [GTShaderProfilerMCABinary.HighRegisterCount]
//   - [GTShaderProfilerMCABinary.ProgramType]
//   - [GTShaderProfilerMCABinary.UniqueIdentifier]
//   - [GTShaderProfilerMCABinary.InitWithAGX2BinaryProgramTypeUniqueIdentifier]
//   - [GTShaderProfilerMCABinary.InitWithAPSBinaryProgramTypeUniqueIdentifier]
//   - [GTShaderProfilerMCABinary.InitWithMioBinaryProgramTypeUniqueIdentifier]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerMCABinary
type GTShaderProfilerMCABinary struct {
	objectivec.Object
}

// GTShaderProfilerMCABinaryFromID constructs a [GTShaderProfilerMCABinary] from an objc.ID.
func GTShaderProfilerMCABinaryFromID(id objc.ID) GTShaderProfilerMCABinary {
	return GTShaderProfilerMCABinary{objectivec.Object{ID: id}}
}

// Ensure GTShaderProfilerMCABinary implements IGTShaderProfilerMCABinary.
var _ IGTShaderProfilerMCABinary = GTShaderProfilerMCABinary{}

// An interface definition for the [GTShaderProfilerMCABinary] class.
//
// # Methods
//
//   - [IGTShaderProfilerMCABinary.AllocatedGPRCount]
//   - [IGTShaderProfilerMCABinary.GenerateAGX2Assembly]
//   - [IGTShaderProfilerMCABinary.GenerateAPSAssembly]
//   - [IGTShaderProfilerMCABinary.GenerateAssemblyContent]
//   - [IGTShaderProfilerMCABinary.HighRegisterCount]
//   - [IGTShaderProfilerMCABinary.ProgramType]
//   - [IGTShaderProfilerMCABinary.UniqueIdentifier]
//   - [IGTShaderProfilerMCABinary.InitWithAGX2BinaryProgramTypeUniqueIdentifier]
//   - [IGTShaderProfilerMCABinary.InitWithAPSBinaryProgramTypeUniqueIdentifier]
//   - [IGTShaderProfilerMCABinary.InitWithMioBinaryProgramTypeUniqueIdentifier]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerMCABinary
type IGTShaderProfilerMCABinary interface {
	objectivec.IObject

	// Topic: Methods

	AllocatedGPRCount() int
	GenerateAGX2Assembly() objectivec.IObject
	GenerateAPSAssembly() objectivec.IObject
	GenerateAssemblyContent() objectivec.IObject
	HighRegisterCount() int
	ProgramType() uint32
	UniqueIdentifier() uint64
	InitWithAGX2BinaryProgramTypeUniqueIdentifier(aGX2Binary unsafe.Pointer, type_ uint32, identifier uint64) GTShaderProfilerMCABinary
	InitWithAPSBinaryProgramTypeUniqueIdentifier(aPSBinary unsafe.Pointer, type_ uint32, identifier uint64) GTShaderProfilerMCABinary
	InitWithMioBinaryProgramTypeUniqueIdentifier(binary objectivec.IObject, type_ uint32, identifier uint64) GTShaderProfilerMCABinary
}

// Init initializes the instance.
func (g GTShaderProfilerMCABinary) Init() GTShaderProfilerMCABinary {
	rv := objc.Send[GTShaderProfilerMCABinary](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTShaderProfilerMCABinary) Autorelease() GTShaderProfilerMCABinary {
	rv := objc.Send[GTShaderProfilerMCABinary](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTShaderProfilerMCABinary creates a new GTShaderProfilerMCABinary instance.
func NewGTShaderProfilerMCABinary() GTShaderProfilerMCABinary {
	class := getGTShaderProfilerMCABinaryClass()
	rv := objc.Send[GTShaderProfilerMCABinary](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerMCABinary/initWithAGX2Binary:programType:uniqueIdentifier:
func NewGTShaderProfilerMCABinaryWithAGX2BinaryProgramTypeUniqueIdentifier(aGX2Binary unsafe.Pointer, type_ uint32, identifier uint64) GTShaderProfilerMCABinary {
	instance := getGTShaderProfilerMCABinaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAGX2Binary:programType:uniqueIdentifier:"), aGX2Binary, type_, identifier)
	return GTShaderProfilerMCABinaryFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerMCABinary/initWithAPSBinary:programType:uniqueIdentifier:
func NewGTShaderProfilerMCABinaryWithAPSBinaryProgramTypeUniqueIdentifier(aPSBinary unsafe.Pointer, type_ uint32, identifier uint64) GTShaderProfilerMCABinary {
	instance := getGTShaderProfilerMCABinaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAPSBinary:programType:uniqueIdentifier:"), aPSBinary, type_, identifier)
	return GTShaderProfilerMCABinaryFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerMCABinary/initWithMioBinary:programType:uniqueIdentifier:
func NewGTShaderProfilerMCABinaryWithMioBinaryProgramTypeUniqueIdentifier(binary objectivec.IObject, type_ uint32, identifier uint64) GTShaderProfilerMCABinary {
	instance := getGTShaderProfilerMCABinaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMioBinary:programType:uniqueIdentifier:"), binary, type_, identifier)
	return GTShaderProfilerMCABinaryFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerMCABinary/generateAGX2Assembly
func (g GTShaderProfilerMCABinary) GenerateAGX2Assembly() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("generateAGX2Assembly"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerMCABinary/generateAPSAssembly
func (g GTShaderProfilerMCABinary) GenerateAPSAssembly() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("generateAPSAssembly"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerMCABinary/generateAssemblyContent
func (g GTShaderProfilerMCABinary) GenerateAssemblyContent() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("generateAssemblyContent"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerMCABinary/initWithAGX2Binary:programType:uniqueIdentifier:
func (g GTShaderProfilerMCABinary) InitWithAGX2BinaryProgramTypeUniqueIdentifier(aGX2Binary unsafe.Pointer, type_ uint32, identifier uint64) GTShaderProfilerMCABinary {
	rv := objc.Send[GTShaderProfilerMCABinary](g.ID, objc.Sel("initWithAGX2Binary:programType:uniqueIdentifier:"), aGX2Binary, type_, identifier)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerMCABinary/initWithAPSBinary:programType:uniqueIdentifier:
func (g GTShaderProfilerMCABinary) InitWithAPSBinaryProgramTypeUniqueIdentifier(aPSBinary unsafe.Pointer, type_ uint32, identifier uint64) GTShaderProfilerMCABinary {
	rv := objc.Send[GTShaderProfilerMCABinary](g.ID, objc.Sel("initWithAPSBinary:programType:uniqueIdentifier:"), aPSBinary, type_, identifier)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerMCABinary/initWithMioBinary:programType:uniqueIdentifier:
func (g GTShaderProfilerMCABinary) InitWithMioBinaryProgramTypeUniqueIdentifier(binary objectivec.IObject, type_ uint32, identifier uint64) GTShaderProfilerMCABinary {
	rv := objc.Send[GTShaderProfilerMCABinary](g.ID, objc.Sel("initWithMioBinary:programType:uniqueIdentifier:"), binary, type_, identifier)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerMCABinary/allocatedGPRCount
func (g GTShaderProfilerMCABinary) AllocatedGPRCount() int {
	rv := objc.Send[int](g.ID, objc.Sel("allocatedGPRCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerMCABinary/highRegisterCount
func (g GTShaderProfilerMCABinary) HighRegisterCount() int {
	rv := objc.Send[int](g.ID, objc.Sel("highRegisterCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerMCABinary/programType
func (g GTShaderProfilerMCABinary) ProgramType() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("programType"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerMCABinary/uniqueIdentifier
func (g GTShaderProfilerMCABinary) UniqueIdentifier() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("uniqueIdentifier"))
	return rv
}
