// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTAGX2InstructionPCStatInfoClass] class.
var (
	_GTAGX2InstructionPCStatInfoClassClass     GTAGX2InstructionPCStatInfoClassClass
	_GTAGX2InstructionPCStatInfoClassClassOnce sync.Once
)

func getGTAGX2InstructionPCStatInfoClassClass() GTAGX2InstructionPCStatInfoClassClass {
	_GTAGX2InstructionPCStatInfoClassClassOnce.Do(func() {
		_GTAGX2InstructionPCStatInfoClassClass = GTAGX2InstructionPCStatInfoClassClass{class: objc.GetClass("GTAGX2InstructionPCStatInfoClass")}
	})
	return _GTAGX2InstructionPCStatInfoClassClass
}

// GetGTAGX2InstructionPCStatInfoClassClass returns the class object for GTAGX2InstructionPCStatInfoClass.
func GetGTAGX2InstructionPCStatInfoClassClass() GTAGX2InstructionPCStatInfoClassClass {
	return getGTAGX2InstructionPCStatInfoClassClass()
}

type GTAGX2InstructionPCStatInfoClassClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTAGX2InstructionPCStatInfoClassClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTAGX2InstructionPCStatInfoClassClass) Alloc() GTAGX2InstructionPCStatInfoClass {
	rv := objc.Send[GTAGX2InstructionPCStatInfoClass](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [GTAGX2InstructionPCStatInfoClass.InstructionPCStatInfo]
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2InstructionPCStatInfoClass
type GTAGX2InstructionPCStatInfoClass struct {
	objectivec.Object
}

// GTAGX2InstructionPCStatInfoClassFromID constructs a [GTAGX2InstructionPCStatInfoClass] from an objc.ID.
func GTAGX2InstructionPCStatInfoClassFromID(id objc.ID) GTAGX2InstructionPCStatInfoClass {
	return GTAGX2InstructionPCStatInfoClass{objectivec.Object{ID: id}}
}
// Ensure GTAGX2InstructionPCStatInfoClass implements IGTAGX2InstructionPCStatInfoClass.
var _ IGTAGX2InstructionPCStatInfoClass = GTAGX2InstructionPCStatInfoClass{}

// An interface definition for the [GTAGX2InstructionPCStatInfoClass] class.
//
// # Methods
//
//   - [IGTAGX2InstructionPCStatInfoClass.InstructionPCStatInfo]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2InstructionPCStatInfoClass
type IGTAGX2InstructionPCStatInfoClass interface {
	objectivec.IObject

	// Topic: Methods

	InstructionPCStatInfo() objectivec.IObject
}

// Init initializes the instance.
func (g GTAGX2InstructionPCStatInfoClass) Init() GTAGX2InstructionPCStatInfoClass {
	rv := objc.Send[GTAGX2InstructionPCStatInfoClass](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTAGX2InstructionPCStatInfoClass) Autorelease() GTAGX2InstructionPCStatInfoClass {
	rv := objc.Send[GTAGX2InstructionPCStatInfoClass](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTAGX2InstructionPCStatInfoClass creates a new GTAGX2InstructionPCStatInfoClass instance.
func NewGTAGX2InstructionPCStatInfoClass() GTAGX2InstructionPCStatInfoClass {
	class := getGTAGX2InstructionPCStatInfoClassClass()
	rv := objc.Send[GTAGX2InstructionPCStatInfoClass](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2InstructionPCStatInfoClass/instructionPCStatInfo
func (g GTAGX2InstructionPCStatInfoClass) InstructionPCStatInfo() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("instructionPCStatInfo"))
	return objectivec.Object{ID: rv}
}

