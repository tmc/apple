// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [GTMioInstructionALUSubPipeCountCounter] class.
var (
	_GTMioInstructionALUSubPipeCountCounterClass     GTMioInstructionALUSubPipeCountCounterClass
	_GTMioInstructionALUSubPipeCountCounterClassOnce sync.Once
)

func getGTMioInstructionALUSubPipeCountCounterClass() GTMioInstructionALUSubPipeCountCounterClass {
	_GTMioInstructionALUSubPipeCountCounterClassOnce.Do(func() {
		_GTMioInstructionALUSubPipeCountCounterClass = GTMioInstructionALUSubPipeCountCounterClass{class: objc.GetClass("GTMioInstructionALUSubPipeCountCounter")}
	})
	return _GTMioInstructionALUSubPipeCountCounterClass
}

// GetGTMioInstructionALUSubPipeCountCounterClass returns the class object for GTMioInstructionALUSubPipeCountCounter.
func GetGTMioInstructionALUSubPipeCountCounterClass() GTMioInstructionALUSubPipeCountCounterClass {
	return getGTMioInstructionALUSubPipeCountCounterClass()
}

type GTMioInstructionALUSubPipeCountCounterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioInstructionALUSubPipeCountCounterClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioInstructionALUSubPipeCountCounterClass) Alloc() GTMioInstructionALUSubPipeCountCounter {
	rv := objc.Send[GTMioInstructionALUSubPipeCountCounter](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [GTMioInstructionALUSubPipeCountCounter.SubPipe]
//   - [GTMioInstructionALUSubPipeCountCounter.InitWithContainerInstructionCategoryScopeScopeIndex]
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioInstructionALUSubPipeCountCounter
type GTMioInstructionALUSubPipeCountCounter struct {
	GTMioCounterData
}

// GTMioInstructionALUSubPipeCountCounterFromID constructs a [GTMioInstructionALUSubPipeCountCounter] from an objc.ID.
func GTMioInstructionALUSubPipeCountCounterFromID(id objc.ID) GTMioInstructionALUSubPipeCountCounter {
	return GTMioInstructionALUSubPipeCountCounter{GTMioCounterData: GTMioCounterDataFromID(id)}
}
// Ensure GTMioInstructionALUSubPipeCountCounter implements IGTMioInstructionALUSubPipeCountCounter.
var _ IGTMioInstructionALUSubPipeCountCounter = GTMioInstructionALUSubPipeCountCounter{}

// An interface definition for the [GTMioInstructionALUSubPipeCountCounter] class.
//
// # Methods
//
//   - [IGTMioInstructionALUSubPipeCountCounter.SubPipe]
//   - [IGTMioInstructionALUSubPipeCountCounter.InitWithContainerInstructionCategoryScopeScopeIndex]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioInstructionALUSubPipeCountCounter
type IGTMioInstructionALUSubPipeCountCounter interface {
	IGTMioCounterData

	// Topic: Methods

	SubPipe() uint16
	InitWithContainerInstructionCategoryScopeScopeIndex(container unsafe.Pointer, category uint16, scope uint16, index uint64) GTMioInstructionALUSubPipeCountCounter
}

// Init initializes the instance.
func (g GTMioInstructionALUSubPipeCountCounter) Init() GTMioInstructionALUSubPipeCountCounter {
	rv := objc.Send[GTMioInstructionALUSubPipeCountCounter](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioInstructionALUSubPipeCountCounter) Autorelease() GTMioInstructionALUSubPipeCountCounter {
	rv := objc.Send[GTMioInstructionALUSubPipeCountCounter](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioInstructionALUSubPipeCountCounter creates a new GTMioInstructionALUSubPipeCountCounter instance.
func NewGTMioInstructionALUSubPipeCountCounter() GTMioInstructionALUSubPipeCountCounter {
	class := getGTMioInstructionALUSubPipeCountCounterClass()
	rv := objc.Send[GTMioInstructionALUSubPipeCountCounter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioCounterData/initWithContainer:index:scope:scopeIndex:
func NewGTMioInstructionALUSubPipeCountCounterWithContainerIndexScopeScopeIndex(container unsafe.Pointer, index uint64, scope uint16, index2 uint64) GTMioInstructionALUSubPipeCountCounter {
	instance := getGTMioInstructionALUSubPipeCountCounterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainer:index:scope:scopeIndex:"), container, index, scope, index2)
	return GTMioInstructionALUSubPipeCountCounterFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioInstructionALUSubPipeCountCounter/initWithContainer:instructionCategory:scope:scopeIndex:
func NewGTMioInstructionALUSubPipeCountCounterWithContainerInstructionCategoryScopeScopeIndex(container unsafe.Pointer, category uint16, scope uint16, index uint64) GTMioInstructionALUSubPipeCountCounter {
	instance := getGTMioInstructionALUSubPipeCountCounterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainer:instructionCategory:scope:scopeIndex:"), container, category, scope, index)
	return GTMioInstructionALUSubPipeCountCounterFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioInstructionALUSubPipeCountCounter/initWithContainer:instructionCategory:scope:scopeIndex:
func (g GTMioInstructionALUSubPipeCountCounter) InitWithContainerInstructionCategoryScopeScopeIndex(container unsafe.Pointer, category uint16, scope uint16, index uint64) GTMioInstructionALUSubPipeCountCounter {
	rv := objc.Send[GTMioInstructionALUSubPipeCountCounter](g.ID, objc.Sel("initWithContainer:instructionCategory:scope:scopeIndex:"), container, category, scope, index)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioInstructionALUSubPipeCountCounter/subPipe
func (g GTMioInstructionALUSubPipeCountCounter) SubPipe() uint16 {
	rv := objc.Send[uint16](g.ID, objc.Sel("subPipe"))
	return rv
}

