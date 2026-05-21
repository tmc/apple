// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTMioShaderProfilerShaderFunction] class.
var (
	_GTMioShaderProfilerShaderFunctionClass     GTMioShaderProfilerShaderFunctionClass
	_GTMioShaderProfilerShaderFunctionClassOnce sync.Once
)

func getGTMioShaderProfilerShaderFunctionClass() GTMioShaderProfilerShaderFunctionClass {
	_GTMioShaderProfilerShaderFunctionClassOnce.Do(func() {
		_GTMioShaderProfilerShaderFunctionClass = GTMioShaderProfilerShaderFunctionClass{class: objc.GetClass("GTMioShaderProfilerShaderFunction")}
	})
	return _GTMioShaderProfilerShaderFunctionClass
}

// GetGTMioShaderProfilerShaderFunctionClass returns the class object for GTMioShaderProfilerShaderFunction.
func GetGTMioShaderProfilerShaderFunctionClass() GTMioShaderProfilerShaderFunctionClass {
	return getGTMioShaderProfilerShaderFunctionClass()
}

type GTMioShaderProfilerShaderFunctionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioShaderProfilerShaderFunctionClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioShaderProfilerShaderFunctionClass) Alloc() GTMioShaderProfilerShaderFunction {
	rv := objc.Send[GTMioShaderProfilerShaderFunction](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTMioShaderProfilerShaderFunction.FilePath]
//   - [GTMioShaderProfilerShaderFunction.Index]
//   - [GTMioShaderProfilerShaderFunction.LibraryObjectId]
//   - [GTMioShaderProfilerShaderFunction.Name]
//   - [GTMioShaderProfilerShaderFunction.ObjectId]
//   - [GTMioShaderProfilerShaderFunction.PointerId]
//   - [GTMioShaderProfilerShaderFunction.Type]
//   - [GTMioShaderProfilerShaderFunction.InitWithInfoTraceData]
//   - [GTMioShaderProfilerShaderFunction.DebugDescription]
//   - [GTMioShaderProfilerShaderFunction.Description]
//   - [GTMioShaderProfilerShaderFunction.Hash]
//   - [GTMioShaderProfilerShaderFunction.Superclass]
type GTMioShaderProfilerShaderFunction struct {
	objectivec.Object
}

// GTMioShaderProfilerShaderFunctionFromID constructs a [GTMioShaderProfilerShaderFunction] from an objc.ID.
func GTMioShaderProfilerShaderFunctionFromID(id objc.ID) GTMioShaderProfilerShaderFunction {
	return GTMioShaderProfilerShaderFunction{objectivec.Object{ID: id}}
}

// Ensure GTMioShaderProfilerShaderFunction implements IGTMioShaderProfilerShaderFunction.
var _ IGTMioShaderProfilerShaderFunction = GTMioShaderProfilerShaderFunction{}

// An interface definition for the [GTMioShaderProfilerShaderFunction] class.
//
// # Methods
//
//   - [IGTMioShaderProfilerShaderFunction.FilePath]
//   - [IGTMioShaderProfilerShaderFunction.Index]
//   - [IGTMioShaderProfilerShaderFunction.LibraryObjectId]
//   - [IGTMioShaderProfilerShaderFunction.Name]
//   - [IGTMioShaderProfilerShaderFunction.ObjectId]
//   - [IGTMioShaderProfilerShaderFunction.PointerId]
//   - [IGTMioShaderProfilerShaderFunction.Type]
//   - [IGTMioShaderProfilerShaderFunction.InitWithInfoTraceData]
//   - [IGTMioShaderProfilerShaderFunction.DebugDescription]
//   - [IGTMioShaderProfilerShaderFunction.Description]
//   - [IGTMioShaderProfilerShaderFunction.Hash]
//   - [IGTMioShaderProfilerShaderFunction.Superclass]
type IGTMioShaderProfilerShaderFunction interface {
	objectivec.IObject

	// Topic: Methods

	FilePath() string
	Index() uint32
	LibraryObjectId() uint64
	Name() string
	ObjectId() uint64
	PointerId() uint64
	Type() uint32
	InitWithInfoTraceData(info unsafe.Pointer, data objectivec.IObject) GTMioShaderProfilerShaderFunction
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (g GTMioShaderProfilerShaderFunction) Init() GTMioShaderProfilerShaderFunction {
	rv := objc.Send[GTMioShaderProfilerShaderFunction](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioShaderProfilerShaderFunction) Autorelease() GTMioShaderProfilerShaderFunction {
	rv := objc.Send[GTMioShaderProfilerShaderFunction](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioShaderProfilerShaderFunction creates a new GTMioShaderProfilerShaderFunction instance.
func NewGTMioShaderProfilerShaderFunction() GTMioShaderProfilerShaderFunction {
	class := getGTMioShaderProfilerShaderFunctionClass()
	rv := objc.Send[GTMioShaderProfilerShaderFunction](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGTMioShaderProfilerShaderFunctionWithInfoTraceData(info unsafe.Pointer, data objectivec.IObject) GTMioShaderProfilerShaderFunction {
	instance := getGTMioShaderProfilerShaderFunctionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithInfo:traceData:"), info, data)
	return GTMioShaderProfilerShaderFunctionFromID(rv)
}

func (g GTMioShaderProfilerShaderFunction) InitWithInfoTraceData(info unsafe.Pointer, data objectivec.IObject) GTMioShaderProfilerShaderFunction {
	rv := objc.Send[GTMioShaderProfilerShaderFunction](g.ID, objc.Sel("initWithInfo:traceData:"), info, data)
	return rv
}

func (g GTMioShaderProfilerShaderFunction) DebugDescription() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTMioShaderProfilerShaderFunction) Description() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTMioShaderProfilerShaderFunction) FilePath() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("filePath"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTMioShaderProfilerShaderFunction) Hash() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("hash"))
	return rv
}
func (g GTMioShaderProfilerShaderFunction) Index() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("index"))
	return rv
}
func (g GTMioShaderProfilerShaderFunction) LibraryObjectId() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("libraryObjectId"))
	return rv
}
func (g GTMioShaderProfilerShaderFunction) Name() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTMioShaderProfilerShaderFunction) ObjectId() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("objectId"))
	return rv
}
func (g GTMioShaderProfilerShaderFunction) PointerId() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("pointerId"))
	return rv
}
func (g GTMioShaderProfilerShaderFunction) Superclass() objectivec.Class {
	rv := objc.Send[objectivec.Class](g.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
func (g GTMioShaderProfilerShaderFunction) Type() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("type"))
	return rv
}
