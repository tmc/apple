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
	rv := objc.SendIfResponds[GTShaderProfilerMCABinary](objc.ID(gc.class), objc.Sel("alloc"))
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
	rv := objc.SendIfResponds[GTShaderProfilerMCABinary](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTShaderProfilerMCABinary) Autorelease() GTShaderProfilerMCABinary {
	rv := objc.SendIfResponds[GTShaderProfilerMCABinary](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTShaderProfilerMCABinary creates a new GTShaderProfilerMCABinary instance.
func NewGTShaderProfilerMCABinary() GTShaderProfilerMCABinary {
	class := getGTShaderProfilerMCABinaryClass()
	rv := objc.SendIfResponds[GTShaderProfilerMCABinary](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGTShaderProfilerMCABinaryWithAGX2BinaryProgramTypeUniqueIdentifier(aGX2Binary unsafe.Pointer, type_ uint32, identifier uint64) GTShaderProfilerMCABinary {
	instance := getGTShaderProfilerMCABinaryClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithAGX2Binary:programType:uniqueIdentifier:"), aGX2Binary, type_, identifier)
	return GTShaderProfilerMCABinaryFromID(rv)
}

func NewGTShaderProfilerMCABinaryWithAPSBinaryProgramTypeUniqueIdentifier(aPSBinary unsafe.Pointer, type_ uint32, identifier uint64) GTShaderProfilerMCABinary {
	instance := getGTShaderProfilerMCABinaryClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithAPSBinary:programType:uniqueIdentifier:"), aPSBinary, type_, identifier)
	return GTShaderProfilerMCABinaryFromID(rv)
}

func NewGTShaderProfilerMCABinaryWithMioBinaryProgramTypeUniqueIdentifier(binary objectivec.IObject, type_ uint32, identifier uint64) GTShaderProfilerMCABinary {
	instance := getGTShaderProfilerMCABinaryClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithMioBinary:programType:uniqueIdentifier:"), binary, type_, identifier)
	return GTShaderProfilerMCABinaryFromID(rv)
}

func (g GTShaderProfilerMCABinary) GenerateAGX2Assembly() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("generateAGX2Assembly"))
	return objectivec.Object{ID: rv}
}
func (g GTShaderProfilerMCABinary) GenerateAPSAssembly() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("generateAPSAssembly"))
	return objectivec.Object{ID: rv}
}
func (g GTShaderProfilerMCABinary) GenerateAssemblyContent() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("generateAssemblyContent"))
	return objectivec.Object{ID: rv}
}
func (g GTShaderProfilerMCABinary) InitWithAGX2BinaryProgramTypeUniqueIdentifier(aGX2Binary unsafe.Pointer, type_ uint32, identifier uint64) GTShaderProfilerMCABinary {
	rv := objc.SendIfResponds[GTShaderProfilerMCABinary](g.ID, objc.Sel("initWithAGX2Binary:programType:uniqueIdentifier:"), aGX2Binary, type_, identifier)
	return rv
}
func (g GTShaderProfilerMCABinary) InitWithAPSBinaryProgramTypeUniqueIdentifier(aPSBinary unsafe.Pointer, type_ uint32, identifier uint64) GTShaderProfilerMCABinary {
	rv := objc.SendIfResponds[GTShaderProfilerMCABinary](g.ID, objc.Sel("initWithAPSBinary:programType:uniqueIdentifier:"), aPSBinary, type_, identifier)
	return rv
}
func (g GTShaderProfilerMCABinary) InitWithMioBinaryProgramTypeUniqueIdentifier(binary objectivec.IObject, type_ uint32, identifier uint64) GTShaderProfilerMCABinary {
	rv := objc.SendIfResponds[GTShaderProfilerMCABinary](g.ID, objc.Sel("initWithMioBinary:programType:uniqueIdentifier:"), binary, type_, identifier)
	return rv
}

func (g GTShaderProfilerMCABinary) AllocatedGPRCount() int {
	rv := objc.SendIfResponds[int](g.ID, objc.Sel("allocatedGPRCount"))
	return rv
}
func (g GTShaderProfilerMCABinary) HighRegisterCount() int {
	rv := objc.SendIfResponds[int](g.ID, objc.Sel("highRegisterCount"))
	return rv
}
func (g GTShaderProfilerMCABinary) ProgramType() uint32 {
	rv := objc.SendIfResponds[uint32](g.ID, objc.Sel("programType"))
	return rv
}
func (g GTShaderProfilerMCABinary) UniqueIdentifier() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("uniqueIdentifier"))
	return rv
}
