// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLModelStructureProgramOperation] class.
var (
	_MLModelStructureProgramOperationClass     MLModelStructureProgramOperationClass
	_MLModelStructureProgramOperationClassOnce sync.Once
)

func getMLModelStructureProgramOperationClass() MLModelStructureProgramOperationClass {
	_MLModelStructureProgramOperationClassOnce.Do(func() {
		_MLModelStructureProgramOperationClass = MLModelStructureProgramOperationClass{class: objc.GetClass("MLModelStructureProgramOperation")}
	})
	return _MLModelStructureProgramOperationClass
}

// GetMLModelStructureProgramOperationClass returns the class object for MLModelStructureProgramOperation.
func GetMLModelStructureProgramOperationClass() MLModelStructureProgramOperationClass {
	return getMLModelStructureProgramOperationClass()
}

type MLModelStructureProgramOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLModelStructureProgramOperationClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelStructureProgramOperationClass) Alloc() MLModelStructureProgramOperation {
	rv := objc.SendIfResponds[MLModelStructureProgramOperation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLModelStructureProgramOperation.MilId]
//   - [MLModelStructureProgramOperation.MilTextLocation]
//   - [MLModelStructureProgramOperation.Path]
//   - [MLModelStructureProgramOperation.InitWithMILOperationPath]
//   - [MLModelStructureProgramOperation.InitWithOperatorNameInputsOutputsBlocksPathMilTextLocationMilId]
type MLModelStructureProgramOperation struct {
	objectivec.Object
}

// MLModelStructureProgramOperationFromID constructs a [MLModelStructureProgramOperation] from an objc.ID.
func MLModelStructureProgramOperationFromID(id objc.ID) MLModelStructureProgramOperation {
	return MLModelStructureProgramOperation{objectivec.Object{ID: id}}
}

// Ensure MLModelStructureProgramOperation implements IMLModelStructureProgramOperation.
var _ IMLModelStructureProgramOperation = MLModelStructureProgramOperation{}

// An interface definition for the [MLModelStructureProgramOperation] class.
//
// # Methods
//
//   - [IMLModelStructureProgramOperation.MilId]
//   - [IMLModelStructureProgramOperation.MilTextLocation]
//   - [IMLModelStructureProgramOperation.Path]
//   - [IMLModelStructureProgramOperation.InitWithMILOperationPath]
//   - [IMLModelStructureProgramOperation.InitWithOperatorNameInputsOutputsBlocksPathMilTextLocationMilId]
type IMLModelStructureProgramOperation interface {
	objectivec.IObject

	// Topic: Methods

	MilId() foundation.NSNumber
	MilTextLocation() string
	Path() IMLModelStructurePath
	InitWithMILOperationPath(mILOperation unsafe.Pointer, path unsafe.Pointer) MLModelStructureProgramOperation
	InitWithOperatorNameInputsOutputsBlocksPathMilTextLocationMilId(name objectivec.IObject, inputs objectivec.IObject, outputs objectivec.IObject, blocks objectivec.IObject, path objectivec.IObject, location objectivec.IObject, id objectivec.IObject) MLModelStructureProgramOperation
}

// Init initializes the instance.
func (m MLModelStructureProgramOperation) Init() MLModelStructureProgramOperation {
	rv := objc.SendIfResponds[MLModelStructureProgramOperation](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelStructureProgramOperation) Autorelease() MLModelStructureProgramOperation {
	rv := objc.SendIfResponds[MLModelStructureProgramOperation](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelStructureProgramOperation creates a new MLModelStructureProgramOperation instance.
func NewMLModelStructureProgramOperation() MLModelStructureProgramOperation {
	class := getMLModelStructureProgramOperationClass()
	rv := objc.SendIfResponds[MLModelStructureProgramOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewModelStructureProgramOperationWithMILOperationPath(mILOperation unsafe.Pointer, path unsafe.Pointer) MLModelStructureProgramOperation {
	instance := getMLModelStructureProgramOperationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithMILOperation:path:"), mILOperation, path)
	return MLModelStructureProgramOperationFromID(rv)
}

func NewModelStructureProgramOperationWithOperatorNameInputsOutputsBlocksPathMilTextLocationMilId(name objectivec.IObject, inputs objectivec.IObject, outputs objectivec.IObject, blocks objectivec.IObject, path objectivec.IObject, location objectivec.IObject, id objectivec.IObject) MLModelStructureProgramOperation {
	instance := getMLModelStructureProgramOperationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithOperatorName:inputs:outputs:blocks:path:milTextLocation:milId:"), name, inputs, outputs, blocks, path, location, id)
	return MLModelStructureProgramOperationFromID(rv)
}

func (m MLModelStructureProgramOperation) InitWithMILOperationPath(mILOperation unsafe.Pointer, path unsafe.Pointer) MLModelStructureProgramOperation {
	rv := objc.SendIfResponds[MLModelStructureProgramOperation](m.ID, objc.Sel("initWithMILOperation:path:"), mILOperation, path)
	return rv
}
func (m MLModelStructureProgramOperation) InitWithOperatorNameInputsOutputsBlocksPathMilTextLocationMilId(name objectivec.IObject, inputs objectivec.IObject, outputs objectivec.IObject, blocks objectivec.IObject, path objectivec.IObject, location objectivec.IObject, id objectivec.IObject) MLModelStructureProgramOperation {
	rv := objc.SendIfResponds[MLModelStructureProgramOperation](m.ID, objc.Sel("initWithOperatorName:inputs:outputs:blocks:path:milTextLocation:milId:"), name, inputs, outputs, blocks, path, location, id)
	return rv
}

func (m MLModelStructureProgramOperation) MilId() foundation.NSNumber {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("milId"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (m MLModelStructureProgramOperation) MilTextLocation() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("milTextLocation"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLModelStructureProgramOperation) Path() IMLModelStructurePath {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("path"))
	return MLModelStructurePathFromID(objc.ID(rv))
}
