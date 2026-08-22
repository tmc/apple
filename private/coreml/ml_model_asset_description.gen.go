// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLModelAssetDescription] class.
var (
	_MLModelAssetDescriptionClass     MLModelAssetDescriptionClass
	_MLModelAssetDescriptionClassOnce sync.Once
)

func getMLModelAssetDescriptionClass() MLModelAssetDescriptionClass {
	_MLModelAssetDescriptionClassOnce.Do(func() {
		_MLModelAssetDescriptionClass = MLModelAssetDescriptionClass{class: objc.GetClass("MLModelAssetDescription")}
	})
	return _MLModelAssetDescriptionClass
}

// GetMLModelAssetDescriptionClass returns the class object for MLModelAssetDescription.
func GetMLModelAssetDescriptionClass() MLModelAssetDescriptionClass {
	return getMLModelAssetDescriptionClass()
}

type MLModelAssetDescriptionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLModelAssetDescriptionClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelAssetDescriptionClass) Alloc() MLModelAssetDescription {
	rv := objc.SendIfResponds[MLModelAssetDescription](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLModelAssetDescription.AssetDescriptionBySettingClassLabels]
//   - [MLModelAssetDescription.AssetDescriptionBySettingClassLabelsFunctionName]
//   - [MLModelAssetDescription.DefaultFunctionName]
//   - [MLModelAssetDescription.DefaultFunctionNameOrEmptyString]
//   - [MLModelAssetDescription.DefaultModelDescription]
//   - [MLModelAssetDescription.FunctionNames]
//   - [MLModelAssetDescription.ModelDescriptionOfFunctionNamed]
//   - [MLModelAssetDescription.ModelDescriptionsByFunctionName]
//   - [MLModelAssetDescription.UsesMultiFunctionSyntax]
//   - [MLModelAssetDescription.InitFromModelDescriptionSpecification]
//   - [MLModelAssetDescription.InitFromModelSpecification]
//   - [MLModelAssetDescription.InitWithCompiledModelArchiveError]
//   - [MLModelAssetDescription.InitWithModelDescription]
//   - [MLModelAssetDescription.InitWithModelDescriptionsByFunctionNameFunctionNamesDefaultFunctionName]
//   - [MLModelAssetDescription.InitWithRawModelDescription]
type MLModelAssetDescription struct {
	objectivec.Object
}

// MLModelAssetDescriptionFromID constructs a [MLModelAssetDescription] from an objc.ID.
func MLModelAssetDescriptionFromID(id objc.ID) MLModelAssetDescription {
	return MLModelAssetDescription{objectivec.Object{ID: id}}
}

// Ensure MLModelAssetDescription implements IMLModelAssetDescription.
var _ IMLModelAssetDescription = MLModelAssetDescription{}

// An interface definition for the [MLModelAssetDescription] class.
//
// # Methods
//
//   - [IMLModelAssetDescription.AssetDescriptionBySettingClassLabels]
//   - [IMLModelAssetDescription.AssetDescriptionBySettingClassLabelsFunctionName]
//   - [IMLModelAssetDescription.DefaultFunctionName]
//   - [IMLModelAssetDescription.DefaultFunctionNameOrEmptyString]
//   - [IMLModelAssetDescription.DefaultModelDescription]
//   - [IMLModelAssetDescription.FunctionNames]
//   - [IMLModelAssetDescription.ModelDescriptionOfFunctionNamed]
//   - [IMLModelAssetDescription.ModelDescriptionsByFunctionName]
//   - [IMLModelAssetDescription.UsesMultiFunctionSyntax]
//   - [IMLModelAssetDescription.InitFromModelDescriptionSpecification]
//   - [IMLModelAssetDescription.InitFromModelSpecification]
//   - [IMLModelAssetDescription.InitWithCompiledModelArchiveError]
//   - [IMLModelAssetDescription.InitWithModelDescription]
//   - [IMLModelAssetDescription.InitWithModelDescriptionsByFunctionNameFunctionNamesDefaultFunctionName]
//   - [IMLModelAssetDescription.InitWithRawModelDescription]
type IMLModelAssetDescription interface {
	objectivec.IObject

	// Topic: Methods

	AssetDescriptionBySettingClassLabels(labels objectivec.IObject) objectivec.IObject
	AssetDescriptionBySettingClassLabelsFunctionName(labels objectivec.IObject, name objectivec.IObject) objectivec.IObject
	DefaultFunctionName() string
	DefaultFunctionNameOrEmptyString() string
	DefaultModelDescription() IMLModelDescription
	FunctionNames() foundation.INSArray
	ModelDescriptionOfFunctionNamed(named objectivec.IObject) objectivec.IObject
	ModelDescriptionsByFunctionName() foundation.INSDictionary
	UsesMultiFunctionSyntax() bool
	InitFromModelDescriptionSpecification(specification MLModelDescriptionSpecificationRef) MLModelAssetDescription
	InitFromModelSpecification(specification MLModelSpecificationRef) MLModelAssetDescription
	InitWithCompiledModelArchiveError(archive MLModelInputArchiverRef) (MLModelAssetDescription, error)
	InitWithModelDescription(description objectivec.IObject) MLModelAssetDescription
	InitWithModelDescriptionsByFunctionNameFunctionNamesDefaultFunctionName(name objectivec.IObject, names objectivec.IObject, name2 objectivec.IObject) MLModelAssetDescription
	InitWithRawModelDescription(description objectivec.IObject) MLModelAssetDescription
}

// Init initializes the instance.
func (m MLModelAssetDescription) Init() MLModelAssetDescription {
	rv := objc.SendIfResponds[MLModelAssetDescription](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelAssetDescription) Autorelease() MLModelAssetDescription {
	rv := objc.SendIfResponds[MLModelAssetDescription](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelAssetDescription creates a new MLModelAssetDescription instance.
func NewMLModelAssetDescription() MLModelAssetDescription {
	class := getMLModelAssetDescriptionClass()
	rv := objc.SendIfResponds[MLModelAssetDescription](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewModelAssetDescriptionFromModelDescriptionSpecification(specification MLModelDescriptionSpecificationRef) MLModelAssetDescription {
	instance := getMLModelAssetDescriptionClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initFromModelDescriptionSpecification:"), specification)
	return MLModelAssetDescriptionFromID(rv)
}

func NewModelAssetDescriptionFromModelSpecification(specification MLModelSpecificationRef) MLModelAssetDescription {
	instance := getMLModelAssetDescriptionClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initFromModelSpecification:"), specification)
	return MLModelAssetDescriptionFromID(rv)
}

func NewModelAssetDescriptionWithCompiledModelArchiveError(archive MLModelInputArchiverRef) (MLModelAssetDescription, error) {
	var errorPtr objc.ID
	instance := getMLModelAssetDescriptionClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCompiledModelArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModelAssetDescription{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLModelAssetDescription{}, objc.ErrInitFailed
	}
	return MLModelAssetDescriptionFromID(rv), nil
}

func NewModelAssetDescriptionWithModelDescription(description objectivec.IObject) MLModelAssetDescription {
	instance := getMLModelAssetDescriptionClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithModelDescription:"), description)
	return MLModelAssetDescriptionFromID(rv)
}

func NewModelAssetDescriptionWithModelDescriptionsByFunctionNameFunctionNamesDefaultFunctionName(name objectivec.IObject, names objectivec.IObject, name2 objectivec.IObject) MLModelAssetDescription {
	instance := getMLModelAssetDescriptionClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithModelDescriptionsByFunctionName:functionNames:defaultFunctionName:"), name, names, name2)
	return MLModelAssetDescriptionFromID(rv)
}

func NewModelAssetDescriptionWithRawModelDescription(description objectivec.IObject) MLModelAssetDescription {
	instance := getMLModelAssetDescriptionClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithRawModelDescription:"), description)
	return MLModelAssetDescriptionFromID(rv)
}

func (m MLModelAssetDescription) AssetDescriptionBySettingClassLabels(labels objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("assetDescriptionBySettingClassLabels:"), labels)
	return objectivec.Object{ID: rv}
}
func (m MLModelAssetDescription) AssetDescriptionBySettingClassLabelsFunctionName(labels objectivec.IObject, name objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("assetDescriptionBySettingClassLabels:functionName:"), labels, name)
	return objectivec.Object{ID: rv}
}
func (m MLModelAssetDescription) ModelDescriptionOfFunctionNamed(named objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("modelDescriptionOfFunctionNamed:"), named)
	return objectivec.Object{ID: rv}
}
func (m MLModelAssetDescription) InitFromModelDescriptionSpecification(specification MLModelDescriptionSpecificationRef) MLModelAssetDescription {
	rv := objc.SendIfResponds[MLModelAssetDescription](m.ID, objc.Sel("initFromModelDescriptionSpecification:"), specification)
	return rv
}
func (m MLModelAssetDescription) InitFromModelSpecification(specification MLModelSpecificationRef) MLModelAssetDescription {
	rv := objc.SendIfResponds[MLModelAssetDescription](m.ID, objc.Sel("initFromModelSpecification:"), specification)
	return rv
}
func (m MLModelAssetDescription) InitWithCompiledModelArchiveError(archive MLModelInputArchiverRef) (MLModelAssetDescription, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithCompiledModelArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModelAssetDescription{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLModelAssetDescriptionFromID(rv), nil

}
func (m MLModelAssetDescription) InitWithModelDescription(description objectivec.IObject) MLModelAssetDescription {
	rv := objc.SendIfResponds[MLModelAssetDescription](m.ID, objc.Sel("initWithModelDescription:"), description)
	return rv
}
func (m MLModelAssetDescription) InitWithModelDescriptionsByFunctionNameFunctionNamesDefaultFunctionName(name objectivec.IObject, names objectivec.IObject, name2 objectivec.IObject) MLModelAssetDescription {
	rv := objc.SendIfResponds[MLModelAssetDescription](m.ID, objc.Sel("initWithModelDescriptionsByFunctionName:functionNames:defaultFunctionName:"), name, names, name2)
	return rv
}
func (m MLModelAssetDescription) InitWithRawModelDescription(description objectivec.IObject) MLModelAssetDescription {
	rv := objc.SendIfResponds[MLModelAssetDescription](m.ID, objc.Sel("initWithRawModelDescription:"), description)
	return rv
}

func (m MLModelAssetDescription) DefaultFunctionName() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("defaultFunctionName"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLModelAssetDescription) DefaultFunctionNameOrEmptyString() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("defaultFunctionNameOrEmptyString"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLModelAssetDescription) DefaultModelDescription() IMLModelDescription {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("defaultModelDescription"))
	return MLModelDescriptionFromID(objc.ID(rv))
}
func (m MLModelAssetDescription) FunctionNames() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("functionNames"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLModelAssetDescription) ModelDescriptionsByFunctionName() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("modelDescriptionsByFunctionName"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLModelAssetDescription) UsesMultiFunctionSyntax() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("usesMultiFunctionSyntax"))
	return rv
}
