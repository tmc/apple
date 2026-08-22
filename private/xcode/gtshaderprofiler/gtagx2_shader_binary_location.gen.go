// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTAGX2ShaderBinaryLocation] class.
var (
	_GTAGX2ShaderBinaryLocationClass     GTAGX2ShaderBinaryLocationClass
	_GTAGX2ShaderBinaryLocationClassOnce sync.Once
)

func getGTAGX2ShaderBinaryLocationClass() GTAGX2ShaderBinaryLocationClass {
	_GTAGX2ShaderBinaryLocationClassOnce.Do(func() {
		_GTAGX2ShaderBinaryLocationClass = GTAGX2ShaderBinaryLocationClass{class: objc.GetClass("GTAGX2ShaderBinaryLocation")}
	})
	return _GTAGX2ShaderBinaryLocationClass
}

// GetGTAGX2ShaderBinaryLocationClass returns the class object for GTAGX2ShaderBinaryLocation.
func GetGTAGX2ShaderBinaryLocationClass() GTAGX2ShaderBinaryLocationClass {
	return getGTAGX2ShaderBinaryLocationClass()
}

type GTAGX2ShaderBinaryLocationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTAGX2ShaderBinaryLocationClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTAGX2ShaderBinaryLocationClass) Alloc() GTAGX2ShaderBinaryLocation {
	rv := objc.SendIfResponds[GTAGX2ShaderBinaryLocation](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTAGX2ShaderBinaryLocation.Binary]
//   - [GTAGX2ShaderBinaryLocation.Column]
//   - [GTAGX2ShaderBinaryLocation.SetColumn]
//   - [GTAGX2ShaderBinaryLocation.EncodeWithCoder]
//   - [GTAGX2ShaderBinaryLocation.FileIndex]
//   - [GTAGX2ShaderBinaryLocation.FullPath]
//   - [GTAGX2ShaderBinaryLocation.FunctionName]
//   - [GTAGX2ShaderBinaryLocation.FunctionNameIndex]
//   - [GTAGX2ShaderBinaryLocation.Line]
//   - [GTAGX2ShaderBinaryLocation.SetLine]
//   - [GTAGX2ShaderBinaryLocation.SetParent]
//   - [GTAGX2ShaderBinaryLocation.InitWithCoder]
//   - [GTAGX2ShaderBinaryLocation.InitWithFunctionNameIndexFullPathIndexLineColumnBinary]
//   - [GTAGX2ShaderBinaryLocation.DebugDescription]
//   - [GTAGX2ShaderBinaryLocation.Description]
//   - [GTAGX2ShaderBinaryLocation.Hash]
//   - [GTAGX2ShaderBinaryLocation.Superclass]
type GTAGX2ShaderBinaryLocation struct {
	objectivec.Object
}

// GTAGX2ShaderBinaryLocationFromID constructs a [GTAGX2ShaderBinaryLocation] from an objc.ID.
func GTAGX2ShaderBinaryLocationFromID(id objc.ID) GTAGX2ShaderBinaryLocation {
	return GTAGX2ShaderBinaryLocation{objectivec.Object{ID: id}}
}

// Ensure GTAGX2ShaderBinaryLocation implements IGTAGX2ShaderBinaryLocation.
var _ IGTAGX2ShaderBinaryLocation = GTAGX2ShaderBinaryLocation{}

// An interface definition for the [GTAGX2ShaderBinaryLocation] class.
//
// # Methods
//
//   - [IGTAGX2ShaderBinaryLocation.Binary]
//   - [IGTAGX2ShaderBinaryLocation.Column]
//   - [IGTAGX2ShaderBinaryLocation.SetColumn]
//   - [IGTAGX2ShaderBinaryLocation.EncodeWithCoder]
//   - [IGTAGX2ShaderBinaryLocation.FileIndex]
//   - [IGTAGX2ShaderBinaryLocation.FullPath]
//   - [IGTAGX2ShaderBinaryLocation.FunctionName]
//   - [IGTAGX2ShaderBinaryLocation.FunctionNameIndex]
//   - [IGTAGX2ShaderBinaryLocation.Line]
//   - [IGTAGX2ShaderBinaryLocation.SetLine]
//   - [IGTAGX2ShaderBinaryLocation.SetParent]
//   - [IGTAGX2ShaderBinaryLocation.InitWithCoder]
//   - [IGTAGX2ShaderBinaryLocation.InitWithFunctionNameIndexFullPathIndexLineColumnBinary]
//   - [IGTAGX2ShaderBinaryLocation.DebugDescription]
//   - [IGTAGX2ShaderBinaryLocation.Description]
//   - [IGTAGX2ShaderBinaryLocation.Hash]
//   - [IGTAGX2ShaderBinaryLocation.Superclass]
type IGTAGX2ShaderBinaryLocation interface {
	objectivec.IObject

	// Topic: Methods

	Binary() unsafe.Pointer
	Column() int
	SetColumn(value int)
	EncodeWithCoder(coder foundation.INSCoder)
	FileIndex() uint64
	FullPath() string
	FunctionName() string
	FunctionNameIndex() uint64
	Line() int
	SetLine(value int)
	SetParent(parent objectivec.IObject)
	InitWithCoder(coder foundation.INSCoder) GTAGX2ShaderBinaryLocation
	InitWithFunctionNameIndexFullPathIndexLineColumnBinary(index uint64, index2 uint64, line int, column int, binary objectivec.IObject) GTAGX2ShaderBinaryLocation
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (g GTAGX2ShaderBinaryLocation) Init() GTAGX2ShaderBinaryLocation {
	rv := objc.SendIfResponds[GTAGX2ShaderBinaryLocation](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTAGX2ShaderBinaryLocation) Autorelease() GTAGX2ShaderBinaryLocation {
	rv := objc.SendIfResponds[GTAGX2ShaderBinaryLocation](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTAGX2ShaderBinaryLocation creates a new GTAGX2ShaderBinaryLocation instance.
func NewGTAGX2ShaderBinaryLocation() GTAGX2ShaderBinaryLocation {
	class := getGTAGX2ShaderBinaryLocationClass()
	rv := objc.SendIfResponds[GTAGX2ShaderBinaryLocation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGTAGX2ShaderBinaryLocationWithCoder(coder objectivec.IObject) GTAGX2ShaderBinaryLocation {
	instance := getGTAGX2ShaderBinaryLocationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return GTAGX2ShaderBinaryLocationFromID(rv)
}

func NewGTAGX2ShaderBinaryLocationWithFunctionNameIndexFullPathIndexLineColumnBinary(index uint64, index2 uint64, line int, column int, binary objectivec.IObject) GTAGX2ShaderBinaryLocation {
	instance := getGTAGX2ShaderBinaryLocationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithFunctionNameIndex:fullPathIndex:line:column:binary:"), index, index2, line, column, binary)
	return GTAGX2ShaderBinaryLocationFromID(rv)
}

func (g GTAGX2ShaderBinaryLocation) EncodeWithCoder(coder foundation.INSCoder) {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (g GTAGX2ShaderBinaryLocation) SetParent(parent objectivec.IObject) {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("setParent:"), parent)
}
func (g GTAGX2ShaderBinaryLocation) InitWithCoder(coder foundation.INSCoder) GTAGX2ShaderBinaryLocation {
	rv := objc.SendIfResponds[GTAGX2ShaderBinaryLocation](g.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (g GTAGX2ShaderBinaryLocation) InitWithFunctionNameIndexFullPathIndexLineColumnBinary(index uint64, index2 uint64, line int, column int, binary objectivec.IObject) GTAGX2ShaderBinaryLocation {
	rv := objc.SendIfResponds[GTAGX2ShaderBinaryLocation](g.ID, objc.Sel("initWithFunctionNameIndex:fullPathIndex:line:column:binary:"), index, index2, line, column, binary)
	return rv
}

func (_GTAGX2ShaderBinaryLocationClass GTAGX2ShaderBinaryLocationClass) SupportsSecureCoding() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_GTAGX2ShaderBinaryLocationClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (g GTAGX2ShaderBinaryLocation) Binary() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](g.ID, objc.Sel("binary"))
	return rv
}
func (g GTAGX2ShaderBinaryLocation) Column() int {
	rv := objc.SendIfResponds[int](g.ID, objc.Sel("column"))
	return rv
}
func (g GTAGX2ShaderBinaryLocation) SetColumn(value int) {
	objc.SendIfResponds[struct{}](g.ID, objc.Sel("setColumn:"), value)
}
func (g GTAGX2ShaderBinaryLocation) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTAGX2ShaderBinaryLocation) Description() string {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTAGX2ShaderBinaryLocation) FileIndex() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("fileIndex"))
	return rv
}
func (g GTAGX2ShaderBinaryLocation) FullPath() string {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("fullPath"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTAGX2ShaderBinaryLocation) FunctionName() string {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("functionName"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTAGX2ShaderBinaryLocation) FunctionNameIndex() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("functionNameIndex"))
	return rv
}
func (g GTAGX2ShaderBinaryLocation) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("hash"))
	return rv
}
func (g GTAGX2ShaderBinaryLocation) Line() int {
	rv := objc.SendIfResponds[int](g.ID, objc.Sel("line"))
	return rv
}
func (g GTAGX2ShaderBinaryLocation) SetLine(value int) {
	objc.SendIfResponds[struct{}](g.ID, objc.Sel("setLine:"), value)
}
func (g GTAGX2ShaderBinaryLocation) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](g.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
