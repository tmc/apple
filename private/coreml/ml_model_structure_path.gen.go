// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLModelStructurePath] class.
var (
	_MLModelStructurePathClass     MLModelStructurePathClass
	_MLModelStructurePathClassOnce sync.Once
)

func getMLModelStructurePathClass() MLModelStructurePathClass {
	_MLModelStructurePathClassOnce.Do(func() {
		_MLModelStructurePathClass = MLModelStructurePathClass{class: objc.GetClass("MLModelStructurePath")}
	})
	return _MLModelStructurePathClass
}

// GetMLModelStructurePathClass returns the class object for MLModelStructurePath.
func GetMLModelStructurePathClass() MLModelStructurePathClass {
	return getMLModelStructurePathClass()
}

type MLModelStructurePathClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLModelStructurePathClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelStructurePathClass) Alloc() MLModelStructurePath {
	rv := objc.SendIfResponds[MLModelStructurePath](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLModelStructurePath.Components]
//   - [MLModelStructurePath.CppPath]
//   - [MLModelStructurePath.IsMLProgramOperationPath]
//   - [MLModelStructurePath.IsNeuralNetworkLayerPath]
//   - [MLModelStructurePath.InitWithCppPath]
//   - [MLModelStructurePath.InitWithMLProgramOperationPathComponentsScopedModelNamesError]
//   - [MLModelStructurePath.InitWithNeuralNetworkLayerNameScopedModelNames]
type MLModelStructurePath struct {
	objectivec.Object
}

// MLModelStructurePathFromID constructs a [MLModelStructurePath] from an objc.ID.
func MLModelStructurePathFromID(id objc.ID) MLModelStructurePath {
	return MLModelStructurePath{objectivec.Object{ID: id}}
}

// Ensure MLModelStructurePath implements IMLModelStructurePath.
var _ IMLModelStructurePath = MLModelStructurePath{}

// An interface definition for the [MLModelStructurePath] class.
//
// # Methods
//
//   - [IMLModelStructurePath.Components]
//   - [IMLModelStructurePath.CppPath]
//   - [IMLModelStructurePath.IsMLProgramOperationPath]
//   - [IMLModelStructurePath.IsNeuralNetworkLayerPath]
//   - [IMLModelStructurePath.InitWithCppPath]
//   - [IMLModelStructurePath.InitWithMLProgramOperationPathComponentsScopedModelNamesError]
//   - [IMLModelStructurePath.InitWithNeuralNetworkLayerNameScopedModelNames]
type IMLModelStructurePath interface {
	objectivec.IObject

	// Topic: Methods

	Components() foundation.INSArray
	CppPath() Path
	IsMLProgramOperationPath() bool
	IsNeuralNetworkLayerPath() bool
	InitWithCppPath(path Path) MLModelStructurePath
	InitWithMLProgramOperationPathComponentsScopedModelNamesError(components objectivec.IObject, names objectivec.IObject) (MLModelStructurePath, error)
	InitWithNeuralNetworkLayerNameScopedModelNames(name objectivec.IObject, names objectivec.IObject) MLModelStructurePath
}

// Init initializes the instance.
func (m MLModelStructurePath) Init() MLModelStructurePath {
	rv := objc.SendIfResponds[MLModelStructurePath](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelStructurePath) Autorelease() MLModelStructurePath {
	rv := objc.SendIfResponds[MLModelStructurePath](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelStructurePath creates a new MLModelStructurePath instance.
func NewMLModelStructurePath() MLModelStructurePath {
	class := getMLModelStructurePathClass()
	rv := objc.SendIfResponds[MLModelStructurePath](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewModelStructurePathWithCppPath(path Path) MLModelStructurePath {
	instance := getMLModelStructurePathClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCppPath:"), path)
	return MLModelStructurePathFromID(rv)
}

func NewModelStructurePathWithMLProgramOperationPathComponentsScopedModelNamesError(components objectivec.IObject, names objectivec.IObject) (MLModelStructurePath, error) {
	var errorPtr objc.ID
	instance := getMLModelStructurePathClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithMLProgramOperationPathComponents:scopedModelNames:error:"), components, names, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModelStructurePath{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLModelStructurePath{}, objc.ErrInitFailed
	}
	return MLModelStructurePathFromID(rv), nil
}

func NewModelStructurePathWithNeuralNetworkLayerNameScopedModelNames(name objectivec.IObject, names objectivec.IObject) MLModelStructurePath {
	instance := getMLModelStructurePathClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithNeuralNetworkLayerName:scopedModelNames:"), name, names)
	return MLModelStructurePathFromID(rv)
}

func (m MLModelStructurePath) InitWithCppPath(path Path) MLModelStructurePath {
	rv := objc.SendIfResponds[MLModelStructurePath](m.ID, objc.Sel("initWithCppPath:"), path)
	return rv
}
func (m MLModelStructurePath) InitWithMLProgramOperationPathComponentsScopedModelNamesError(components objectivec.IObject, names objectivec.IObject) (MLModelStructurePath, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithMLProgramOperationPathComponents:scopedModelNames:error:"), components, names, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModelStructurePath{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLModelStructurePathFromID(rv), nil

}
func (m MLModelStructurePath) InitWithNeuralNetworkLayerNameScopedModelNames(name objectivec.IObject, names objectivec.IObject) MLModelStructurePath {
	rv := objc.SendIfResponds[MLModelStructurePath](m.ID, objc.Sel("initWithNeuralNetworkLayerName:scopedModelNames:"), name, names)
	return rv
}

func (m MLModelStructurePath) Components() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("components"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLModelStructurePath) CppPath() Path {
	rv := objc.SendIfResponds[Path](m.ID, objc.Sel("cppPath"))
	return Path(rv)
}
func (m MLModelStructurePath) IsMLProgramOperationPath() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("isMLProgramOperationPath"))
	return rv
}
func (m MLModelStructurePath) IsNeuralNetworkLayerPath() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("isNeuralNetworkLayerPath"))
	return rv
}
