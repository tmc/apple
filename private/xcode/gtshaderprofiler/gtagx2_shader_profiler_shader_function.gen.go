// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTAGX2ShaderProfilerShaderFunction] class.
var (
	_GTAGX2ShaderProfilerShaderFunctionClass     GTAGX2ShaderProfilerShaderFunctionClass
	_GTAGX2ShaderProfilerShaderFunctionClassOnce sync.Once
)

func getGTAGX2ShaderProfilerShaderFunctionClass() GTAGX2ShaderProfilerShaderFunctionClass {
	_GTAGX2ShaderProfilerShaderFunctionClassOnce.Do(func() {
		_GTAGX2ShaderProfilerShaderFunctionClass = GTAGX2ShaderProfilerShaderFunctionClass{class: objc.GetClass("GTAGX2ShaderProfilerShaderFunction")}
	})
	return _GTAGX2ShaderProfilerShaderFunctionClass
}

// GetGTAGX2ShaderProfilerShaderFunctionClass returns the class object for GTAGX2ShaderProfilerShaderFunction.
func GetGTAGX2ShaderProfilerShaderFunctionClass() GTAGX2ShaderProfilerShaderFunctionClass {
	return getGTAGX2ShaderProfilerShaderFunctionClass()
}

type GTAGX2ShaderProfilerShaderFunctionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTAGX2ShaderProfilerShaderFunctionClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTAGX2ShaderProfilerShaderFunctionClass) Alloc() GTAGX2ShaderProfilerShaderFunction {
	rv := objc.Send[GTAGX2ShaderProfilerShaderFunction](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTAGX2ShaderProfilerShaderFunction.EncodeWithCoder]
//   - [GTAGX2ShaderProfilerShaderFunction.FilePath]
//   - [GTAGX2ShaderProfilerShaderFunction.SetFilePath]
//   - [GTAGX2ShaderProfilerShaderFunction.Index]
//   - [GTAGX2ShaderProfilerShaderFunction.SetIndex]
//   - [GTAGX2ShaderProfilerShaderFunction.LibraryObjectId]
//   - [GTAGX2ShaderProfilerShaderFunction.SetLibraryObjectId]
//   - [GTAGX2ShaderProfilerShaderFunction.Name]
//   - [GTAGX2ShaderProfilerShaderFunction.SetName]
//   - [GTAGX2ShaderProfilerShaderFunction.ObjectId]
//   - [GTAGX2ShaderProfilerShaderFunction.SetObjectId]
//   - [GTAGX2ShaderProfilerShaderFunction.PointerId]
//   - [GTAGX2ShaderProfilerShaderFunction.SetPointerId]
//   - [GTAGX2ShaderProfilerShaderFunction.Type]
//   - [GTAGX2ShaderProfilerShaderFunction.SetType]
//   - [GTAGX2ShaderProfilerShaderFunction.InitWithCoder]
//   - [GTAGX2ShaderProfilerShaderFunction.DebugDescription]
//   - [GTAGX2ShaderProfilerShaderFunction.Description]
//   - [GTAGX2ShaderProfilerShaderFunction.Hash]
//   - [GTAGX2ShaderProfilerShaderFunction.Superclass]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerShaderFunction
type GTAGX2ShaderProfilerShaderFunction struct {
	objectivec.Object
}

// GTAGX2ShaderProfilerShaderFunctionFromID constructs a [GTAGX2ShaderProfilerShaderFunction] from an objc.ID.
func GTAGX2ShaderProfilerShaderFunctionFromID(id objc.ID) GTAGX2ShaderProfilerShaderFunction {
	return GTAGX2ShaderProfilerShaderFunction{objectivec.Object{ID: id}}
}

// Ensure GTAGX2ShaderProfilerShaderFunction implements IGTAGX2ShaderProfilerShaderFunction.
var _ IGTAGX2ShaderProfilerShaderFunction = GTAGX2ShaderProfilerShaderFunction{}

// An interface definition for the [GTAGX2ShaderProfilerShaderFunction] class.
//
// # Methods
//
//   - [IGTAGX2ShaderProfilerShaderFunction.EncodeWithCoder]
//   - [IGTAGX2ShaderProfilerShaderFunction.FilePath]
//   - [IGTAGX2ShaderProfilerShaderFunction.SetFilePath]
//   - [IGTAGX2ShaderProfilerShaderFunction.Index]
//   - [IGTAGX2ShaderProfilerShaderFunction.SetIndex]
//   - [IGTAGX2ShaderProfilerShaderFunction.LibraryObjectId]
//   - [IGTAGX2ShaderProfilerShaderFunction.SetLibraryObjectId]
//   - [IGTAGX2ShaderProfilerShaderFunction.Name]
//   - [IGTAGX2ShaderProfilerShaderFunction.SetName]
//   - [IGTAGX2ShaderProfilerShaderFunction.ObjectId]
//   - [IGTAGX2ShaderProfilerShaderFunction.SetObjectId]
//   - [IGTAGX2ShaderProfilerShaderFunction.PointerId]
//   - [IGTAGX2ShaderProfilerShaderFunction.SetPointerId]
//   - [IGTAGX2ShaderProfilerShaderFunction.Type]
//   - [IGTAGX2ShaderProfilerShaderFunction.SetType]
//   - [IGTAGX2ShaderProfilerShaderFunction.InitWithCoder]
//   - [IGTAGX2ShaderProfilerShaderFunction.DebugDescription]
//   - [IGTAGX2ShaderProfilerShaderFunction.Description]
//   - [IGTAGX2ShaderProfilerShaderFunction.Hash]
//   - [IGTAGX2ShaderProfilerShaderFunction.Superclass]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerShaderFunction
type IGTAGX2ShaderProfilerShaderFunction interface {
	objectivec.IObject

	// Topic: Methods

	EncodeWithCoder(coder foundation.INSCoder)
	FilePath() string
	SetFilePath(value string)
	Index() uint32
	SetIndex(value uint32)
	LibraryObjectId() uint64
	SetLibraryObjectId(value uint64)
	Name() string
	SetName(value string)
	ObjectId() uint64
	SetObjectId(value uint64)
	PointerId() uint64
	SetPointerId(value uint64)
	Type() uint32
	SetType(value uint32)
	InitWithCoder(coder foundation.INSCoder) GTAGX2ShaderProfilerShaderFunction
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (g GTAGX2ShaderProfilerShaderFunction) Init() GTAGX2ShaderProfilerShaderFunction {
	rv := objc.Send[GTAGX2ShaderProfilerShaderFunction](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTAGX2ShaderProfilerShaderFunction) Autorelease() GTAGX2ShaderProfilerShaderFunction {
	rv := objc.Send[GTAGX2ShaderProfilerShaderFunction](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTAGX2ShaderProfilerShaderFunction creates a new GTAGX2ShaderProfilerShaderFunction instance.
func NewGTAGX2ShaderProfilerShaderFunction() GTAGX2ShaderProfilerShaderFunction {
	class := getGTAGX2ShaderProfilerShaderFunctionClass()
	rv := objc.Send[GTAGX2ShaderProfilerShaderFunction](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerShaderFunction/initWithCoder:
func NewGTAGX2ShaderProfilerShaderFunctionWithCoder(coder objectivec.IObject) GTAGX2ShaderProfilerShaderFunction {
	instance := getGTAGX2ShaderProfilerShaderFunctionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return GTAGX2ShaderProfilerShaderFunctionFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerShaderFunction/encodeWithCoder:
func (g GTAGX2ShaderProfilerShaderFunction) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](g.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerShaderFunction/initWithCoder:
func (g GTAGX2ShaderProfilerShaderFunction) InitWithCoder(coder foundation.INSCoder) GTAGX2ShaderProfilerShaderFunction {
	rv := objc.Send[GTAGX2ShaderProfilerShaderFunction](g.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerShaderFunction/supportsSecureCoding
func (_GTAGX2ShaderProfilerShaderFunctionClass GTAGX2ShaderProfilerShaderFunctionClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_GTAGX2ShaderProfilerShaderFunctionClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerShaderFunction/debugDescription
func (g GTAGX2ShaderProfilerShaderFunction) DebugDescription() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerShaderFunction/description
func (g GTAGX2ShaderProfilerShaderFunction) Description() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerShaderFunction/filePath
func (g GTAGX2ShaderProfilerShaderFunction) FilePath() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("filePath"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTAGX2ShaderProfilerShaderFunction) SetFilePath(value string) {
	objc.Send[struct{}](g.ID, objc.Sel("setFilePath:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerShaderFunction/hash
func (g GTAGX2ShaderProfilerShaderFunction) Hash() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerShaderFunction/index
func (g GTAGX2ShaderProfilerShaderFunction) Index() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("index"))
	return rv
}
func (g GTAGX2ShaderProfilerShaderFunction) SetIndex(value uint32) {
	objc.Send[struct{}](g.ID, objc.Sel("setIndex:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerShaderFunction/libraryObjectId
func (g GTAGX2ShaderProfilerShaderFunction) LibraryObjectId() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("libraryObjectId"))
	return rv
}
func (g GTAGX2ShaderProfilerShaderFunction) SetLibraryObjectId(value uint64) {
	objc.Send[struct{}](g.ID, objc.Sel("setLibraryObjectId:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerShaderFunction/name
func (g GTAGX2ShaderProfilerShaderFunction) Name() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTAGX2ShaderProfilerShaderFunction) SetName(value string) {
	objc.Send[struct{}](g.ID, objc.Sel("setName:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerShaderFunction/objectId
func (g GTAGX2ShaderProfilerShaderFunction) ObjectId() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("objectId"))
	return rv
}
func (g GTAGX2ShaderProfilerShaderFunction) SetObjectId(value uint64) {
	objc.Send[struct{}](g.ID, objc.Sel("setObjectId:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerShaderFunction/pointerId
func (g GTAGX2ShaderProfilerShaderFunction) PointerId() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("pointerId"))
	return rv
}
func (g GTAGX2ShaderProfilerShaderFunction) SetPointerId(value uint64) {
	objc.Send[struct{}](g.ID, objc.Sel("setPointerId:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerShaderFunction/superclass
func (g GTAGX2ShaderProfilerShaderFunction) Superclass() objc.Class {
	rv := objc.Send[objc.Class](g.ID, objc.Sel("superclass"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerShaderFunction/type
func (g GTAGX2ShaderProfilerShaderFunction) Type() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("type"))
	return rv
}
func (g GTAGX2ShaderProfilerShaderFunction) SetType(value uint32) {
	objc.Send[struct{}](g.ID, objc.Sel("setType:"), value)
}
