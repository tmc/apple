// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANEProgramProcedurePriv] class.
var (
	_ANEProgramProcedurePrivClass     ANEProgramProcedurePrivClass
	_ANEProgramProcedurePrivClassOnce sync.Once
)

func getANEProgramProcedurePrivClass() ANEProgramProcedurePrivClass {
	_ANEProgramProcedurePrivClassOnce.Do(func() {
		_ANEProgramProcedurePrivClass = ANEProgramProcedurePrivClass{class: objc.GetClass("_ANEProgramProcedurePriv")}
	})
	return _ANEProgramProcedurePrivClass
}

// GetANEProgramProcedurePrivClass returns the class object for _ANEProgramProcedurePriv.
func GetANEProgramProcedurePrivClass() ANEProgramProcedurePrivClass {
	return getANEProgramProcedurePrivClass()
}

type ANEProgramProcedurePrivClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac ANEProgramProcedurePrivClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEProgramProcedurePrivClass) Alloc() ANEProgramProcedurePriv {
	rv := objc.SendIfResponds[ANEProgramProcedurePriv](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ANEProgramProcedurePriv.MutableWeightsBuffer]
//   - [ANEProgramProcedurePriv.SetMutableWeightsBuffer]
//   - [ANEProgramProcedurePriv.MutableWeightsBufferPtr]
//   - [ANEProgramProcedurePriv.SymbolName]
//   - [ANEProgramProcedurePriv.SetSymbolName]
//   - [ANEProgramProcedurePriv.InitWithSymbolName]
type ANEProgramProcedurePriv struct {
	objectivec.Object
}

// ANEProgramProcedurePrivFromID constructs a [ANEProgramProcedurePriv] from an objc.ID.
func ANEProgramProcedurePrivFromID(id objc.ID) ANEProgramProcedurePriv {
	return ANEProgramProcedurePriv{objectivec.Object{ID: id}}
}

// Ensure ANEProgramProcedurePriv implements IANEProgramProcedurePriv.
var _ IANEProgramProcedurePriv = ANEProgramProcedurePriv{}

// An interface definition for the [ANEProgramProcedurePriv] class.
//
// # Methods
//
//   - [IANEProgramProcedurePriv.MutableWeightsBuffer]
//   - [IANEProgramProcedurePriv.SetMutableWeightsBuffer]
//   - [IANEProgramProcedurePriv.MutableWeightsBufferPtr]
//   - [IANEProgramProcedurePriv.SymbolName]
//   - [IANEProgramProcedurePriv.SetSymbolName]
//   - [IANEProgramProcedurePriv.InitWithSymbolName]
type IANEProgramProcedurePriv interface {
	objectivec.IObject

	// Topic: Methods

	MutableWeightsBuffer() ANEBufferMapping
	SetMutableWeightsBuffer(value ANEBufferMapping)
	MutableWeightsBufferPtr() *ANEBufferMapping
	SymbolName() string
	SetSymbolName(value string)
	InitWithSymbolName(name objectivec.IObject) ANEProgramProcedurePriv
}

// Init initializes the instance.
func (a ANEProgramProcedurePriv) Init() ANEProgramProcedurePriv {
	rv := objc.SendIfResponds[ANEProgramProcedurePriv](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEProgramProcedurePriv) Autorelease() ANEProgramProcedurePriv {
	rv := objc.SendIfResponds[ANEProgramProcedurePriv](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEProgramProcedurePriv creates a new ANEProgramProcedurePriv instance.
func NewANEProgramProcedurePriv() ANEProgramProcedurePriv {
	class := getANEProgramProcedurePrivClass()
	rv := objc.SendIfResponds[ANEProgramProcedurePriv](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewANEProgramProcedurePrivWithSymbolName(name objectivec.IObject) ANEProgramProcedurePriv {
	instance := getANEProgramProcedurePrivClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithSymbolName:"), name)
	return ANEProgramProcedurePrivFromID(rv)
}

func (a ANEProgramProcedurePriv) MutableWeightsBufferPtr() *ANEBufferMapping {
	rv := objc.SendIfResponds[unsafe.Pointer](a.ID, objc.Sel("mutableWeightsBufferPtr"))
	return (*ANEBufferMapping)(rv)
}
func (a ANEProgramProcedurePriv) InitWithSymbolName(name objectivec.IObject) ANEProgramProcedurePriv {
	rv := objc.SendIfResponds[ANEProgramProcedurePriv](a.ID, objc.Sel("initWithSymbolName:"), name)
	return rv
}

func (a ANEProgramProcedurePriv) MutableWeightsBuffer() ANEBufferMapping {
	rv := objc.SendIfResponds[ANEBufferMapping](a.ID, objc.Sel("mutableWeightsBuffer"))
	return ANEBufferMapping(rv)
}
func (a ANEProgramProcedurePriv) SetMutableWeightsBuffer(value ANEBufferMapping) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setMutableWeightsBuffer:"), value)
}
func (a ANEProgramProcedurePriv) SymbolName() string {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("symbolName"))
	return foundation.NSStringFromID(rv).String()
}
func (a ANEProgramProcedurePriv) SetSymbolName(value string) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setSymbolName:"), objc.String(value))
}
