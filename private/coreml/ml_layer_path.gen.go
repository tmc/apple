// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLLayerPath] class.
var (
	_MLLayerPathClass     MLLayerPathClass
	_MLLayerPathClassOnce sync.Once
)

func getMLLayerPathClass() MLLayerPathClass {
	_MLLayerPathClassOnce.Do(func() {
		_MLLayerPathClass = MLLayerPathClass{class: objc.GetClass("MLLayerPath")}
	})
	return _MLLayerPathClass
}

// GetMLLayerPathClass returns the class object for MLLayerPath.
func GetMLLayerPathClass() MLLayerPathClass {
	return getMLLayerPathClass()
}

type MLLayerPathClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLLayerPathClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLLayerPathClass) Alloc() MLLayerPath {
	rv := objc.Send[MLLayerPath](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLLayerPath.AppendPathComponent]
//   - [MLLayerPath.IsEqualToMLLayerPath]
//   - [MLLayerPath.LayerName]
//   - [MLLayerPath.SetLayerName]
//   - [MLLayerPath.ScopedModelNames]
//   - [MLLayerPath.SetScopedModelNames]
//   - [MLLayerPath.InitWithScopedModelAndLayerNameLayerName]
// See: https://developer.apple.com/documentation/CoreML/MLLayerPath
type MLLayerPath struct {
	objectivec.Object
}

// MLLayerPathFromID constructs a [MLLayerPath] from an objc.ID.
func MLLayerPathFromID(id objc.ID) MLLayerPath {
	return MLLayerPath{objectivec.Object{ID: id}}
}
// Ensure MLLayerPath implements IMLLayerPath.
var _ IMLLayerPath = MLLayerPath{}

// An interface definition for the [MLLayerPath] class.
//
// # Methods
//
//   - [IMLLayerPath.AppendPathComponent]
//   - [IMLLayerPath.IsEqualToMLLayerPath]
//   - [IMLLayerPath.LayerName]
//   - [IMLLayerPath.SetLayerName]
//   - [IMLLayerPath.ScopedModelNames]
//   - [IMLLayerPath.SetScopedModelNames]
//   - [IMLLayerPath.InitWithScopedModelAndLayerNameLayerName]
//
// See: https://developer.apple.com/documentation/CoreML/MLLayerPath
type IMLLayerPath interface {
	objectivec.IObject

	// Topic: Methods

	AppendPathComponent(component objectivec.IObject)
	IsEqualToMLLayerPath(path objectivec.IObject) bool
	LayerName() string
	SetLayerName(value string)
	ScopedModelNames() foundation.INSArray
	SetScopedModelNames(value foundation.INSArray)
	InitWithScopedModelAndLayerNameLayerName(name objectivec.IObject, name2 objectivec.IObject) MLLayerPath
}

// Init initializes the instance.
func (l MLLayerPath) Init() MLLayerPath {
	rv := objc.Send[MLLayerPath](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l MLLayerPath) Autorelease() MLLayerPath {
	rv := objc.Send[MLLayerPath](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLLayerPath creates a new MLLayerPath instance.
func NewMLLayerPath() MLLayerPath {
	class := getMLLayerPathClass()
	rv := objc.Send[MLLayerPath](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLLayerPath/initWithScopedModelAndLayerName:layerName:
func NewLayerPathWithScopedModelAndLayerNameLayerName(name objectivec.IObject, name2 objectivec.IObject) MLLayerPath {
	instance := getMLLayerPathClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithScopedModelAndLayerName:layerName:"), name, name2)
	return MLLayerPathFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLLayerPath/appendPathComponent:
func (l MLLayerPath) AppendPathComponent(component objectivec.IObject) {
	objc.Send[objc.ID](l.ID, objc.Sel("appendPathComponent:"), component)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLLayerPath/isEqualToMLLayerPath:
func (l MLLayerPath) IsEqualToMLLayerPath(path objectivec.IObject) bool {
	rv := objc.Send[bool](l.ID, objc.Sel("isEqualToMLLayerPath:"), path)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLLayerPath/initWithScopedModelAndLayerName:layerName:
func (l MLLayerPath) InitWithScopedModelAndLayerNameLayerName(name objectivec.IObject, name2 objectivec.IObject) MLLayerPath {
	rv := objc.Send[MLLayerPath](l.ID, objc.Sel("initWithScopedModelAndLayerName:layerName:"), name, name2)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLLayerPath/layerName
func (l MLLayerPath) LayerName() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("layerName"))
	return foundation.NSStringFromID(rv).String()
}
func (l MLLayerPath) SetLayerName(value string) {
	objc.Send[struct{}](l.ID, objc.Sel("setLayerName:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/CoreML/MLLayerPath/scopedModelNames
func (l MLLayerPath) ScopedModelNames() foundation.INSArray {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("scopedModelNames"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (l MLLayerPath) SetScopedModelNames(value foundation.INSArray) {
	objc.Send[struct{}](l.ID, objc.Sel("setScopedModelNames:"), value)
}

