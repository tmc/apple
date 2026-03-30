// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
)

// The class instance for the [GTMioInstructionTypeCountCounter] class.
var (
	_GTMioInstructionTypeCountCounterClass     GTMioInstructionTypeCountCounterClass
	_GTMioInstructionTypeCountCounterClassOnce sync.Once
)

func getGTMioInstructionTypeCountCounterClass() GTMioInstructionTypeCountCounterClass {
	_GTMioInstructionTypeCountCounterClassOnce.Do(func() {
		_GTMioInstructionTypeCountCounterClass = GTMioInstructionTypeCountCounterClass{class: objc.GetClass("GTMioInstructionTypeCountCounter")}
	})
	return _GTMioInstructionTypeCountCounterClass
}

// GetGTMioInstructionTypeCountCounterClass returns the class object for GTMioInstructionTypeCountCounter.
func GetGTMioInstructionTypeCountCounterClass() GTMioInstructionTypeCountCounterClass {
	return getGTMioInstructionTypeCountCounterClass()
}

type GTMioInstructionTypeCountCounterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioInstructionTypeCountCounterClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioInstructionTypeCountCounterClass) Alloc() GTMioInstructionTypeCountCounter {
	rv := objc.Send[GTMioInstructionTypeCountCounter](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTMioInstructionTypeCountCounter.InstructionType]
//   - [GTMioInstructionTypeCountCounter.InitWithContainerInstructionTypeScopeScopeIndex]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioInstructionTypeCountCounter
type GTMioInstructionTypeCountCounter struct {
	GTMioCounterData
}

// GTMioInstructionTypeCountCounterFromID constructs a [GTMioInstructionTypeCountCounter] from an objc.ID.
func GTMioInstructionTypeCountCounterFromID(id objc.ID) GTMioInstructionTypeCountCounter {
	return GTMioInstructionTypeCountCounter{GTMioCounterData: GTMioCounterDataFromID(id)}
}

// Ensure GTMioInstructionTypeCountCounter implements IGTMioInstructionTypeCountCounter.
var _ IGTMioInstructionTypeCountCounter = GTMioInstructionTypeCountCounter{}

// An interface definition for the [GTMioInstructionTypeCountCounter] class.
//
// # Methods
//
//   - [IGTMioInstructionTypeCountCounter.InstructionType]
//   - [IGTMioInstructionTypeCountCounter.InitWithContainerInstructionTypeScopeScopeIndex]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioInstructionTypeCountCounter
type IGTMioInstructionTypeCountCounter interface {
	IGTMioCounterData

	// Topic: Methods

	InstructionType() uint16
	InitWithContainerInstructionTypeScopeScopeIndex(container unsafe.Pointer, type_ uint16, scope uint16, index uint64) GTMioInstructionTypeCountCounter
}

// Init initializes the instance.
func (g GTMioInstructionTypeCountCounter) Init() GTMioInstructionTypeCountCounter {
	rv := objc.Send[GTMioInstructionTypeCountCounter](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioInstructionTypeCountCounter) Autorelease() GTMioInstructionTypeCountCounter {
	rv := objc.Send[GTMioInstructionTypeCountCounter](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioInstructionTypeCountCounter creates a new GTMioInstructionTypeCountCounter instance.
func NewGTMioInstructionTypeCountCounter() GTMioInstructionTypeCountCounter {
	class := getGTMioInstructionTypeCountCounterClass()
	rv := objc.Send[GTMioInstructionTypeCountCounter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioCounterData/initWithContainer:index:scope:scopeIndex:
func NewGTMioInstructionTypeCountCounterWithContainerIndexScopeScopeIndex(container unsafe.Pointer, index uint64, scope uint16, index2 uint64) GTMioInstructionTypeCountCounter {
	instance := getGTMioInstructionTypeCountCounterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainer:index:scope:scopeIndex:"), container, index, scope, index2)
	return GTMioInstructionTypeCountCounterFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioInstructionTypeCountCounter/initWithContainer:instructionType:scope:scopeIndex:
func NewGTMioInstructionTypeCountCounterWithContainerInstructionTypeScopeScopeIndex(container unsafe.Pointer, type_ uint16, scope uint16, index uint64) GTMioInstructionTypeCountCounter {
	instance := getGTMioInstructionTypeCountCounterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainer:instructionType:scope:scopeIndex:"), container, type_, scope, index)
	return GTMioInstructionTypeCountCounterFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioInstructionTypeCountCounter/initWithContainer:instructionType:scope:scopeIndex:
func (g GTMioInstructionTypeCountCounter) InitWithContainerInstructionTypeScopeScopeIndex(container unsafe.Pointer, type_ uint16, scope uint16, index uint64) GTMioInstructionTypeCountCounter {
	rv := objc.Send[GTMioInstructionTypeCountCounter](g.ID, objc.Sel("initWithContainer:instructionType:scope:scopeIndex:"), container, type_, scope, index)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioInstructionTypeCountCounter/instructionType
func (g GTMioInstructionTypeCountCounter) InstructionType() uint16 {
	rv := objc.Send[uint16](g.ID, objc.Sel("instructionType"))
	return rv
}
