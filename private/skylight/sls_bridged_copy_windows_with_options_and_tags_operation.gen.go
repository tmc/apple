// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedCopyWindowsWithOptionsAndTagsOperation] class.
var (
	_SLSBridgedCopyWindowsWithOptionsAndTagsOperationClass     SLSBridgedCopyWindowsWithOptionsAndTagsOperationClass
	_SLSBridgedCopyWindowsWithOptionsAndTagsOperationClassOnce sync.Once
)

func getSLSBridgedCopyWindowsWithOptionsAndTagsOperationClass() SLSBridgedCopyWindowsWithOptionsAndTagsOperationClass {
	_SLSBridgedCopyWindowsWithOptionsAndTagsOperationClassOnce.Do(func() {
		_SLSBridgedCopyWindowsWithOptionsAndTagsOperationClass = SLSBridgedCopyWindowsWithOptionsAndTagsOperationClass{class: objc.GetClass("SLSBridgedCopyWindowsWithOptionsAndTagsOperation")}
	})
	return _SLSBridgedCopyWindowsWithOptionsAndTagsOperationClass
}

// GetSLSBridgedCopyWindowsWithOptionsAndTagsOperationClass returns the class object for SLSBridgedCopyWindowsWithOptionsAndTagsOperation.
func GetSLSBridgedCopyWindowsWithOptionsAndTagsOperationClass() SLSBridgedCopyWindowsWithOptionsAndTagsOperationClass {
	return getSLSBridgedCopyWindowsWithOptionsAndTagsOperationClass()
}

type SLSBridgedCopyWindowsWithOptionsAndTagsOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedCopyWindowsWithOptionsAndTagsOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedCopyWindowsWithOptionsAndTagsOperationClass) Alloc() SLSBridgedCopyWindowsWithOptionsAndTagsOperation {
	rv := objc.Send[SLSBridgedCopyWindowsWithOptionsAndTagsOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedCopyWindowsWithOptionsAndTagsOperation.GetClearedTags]
//   - [SLSBridgedCopyWindowsWithOptionsAndTagsOperation.GetSetTags]
//   - [SLSBridgedCopyWindowsWithOptionsAndTagsOperation.MakeResultWithNumbers]
//   - [SLSBridgedCopyWindowsWithOptionsAndTagsOperation.Options]
//   - [SLSBridgedCopyWindowsWithOptionsAndTagsOperation.Owner]
//   - [SLSBridgedCopyWindowsWithOptionsAndTagsOperation.Spaces]
//   - [SLSBridgedCopyWindowsWithOptionsAndTagsOperation.InitWithOwnerSpacesOptionsSetTagsClearedTags]
type SLSBridgedCopyWindowsWithOptionsAndTagsOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedCopyWindowsWithOptionsAndTagsOperationFromID constructs a [SLSBridgedCopyWindowsWithOptionsAndTagsOperation] from an objc.ID.
func SLSBridgedCopyWindowsWithOptionsAndTagsOperationFromID(id objc.ID) SLSBridgedCopyWindowsWithOptionsAndTagsOperation {
	return SLSBridgedCopyWindowsWithOptionsAndTagsOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedCopyWindowsWithOptionsAndTagsOperation implements ISLSBridgedCopyWindowsWithOptionsAndTagsOperation.
var _ ISLSBridgedCopyWindowsWithOptionsAndTagsOperation = SLSBridgedCopyWindowsWithOptionsAndTagsOperation{}

// An interface definition for the [SLSBridgedCopyWindowsWithOptionsAndTagsOperation] class.
//
// # Methods
//
//   - [ISLSBridgedCopyWindowsWithOptionsAndTagsOperation.GetClearedTags]
//   - [ISLSBridgedCopyWindowsWithOptionsAndTagsOperation.GetSetTags]
//   - [ISLSBridgedCopyWindowsWithOptionsAndTagsOperation.MakeResultWithNumbers]
//   - [ISLSBridgedCopyWindowsWithOptionsAndTagsOperation.Options]
//   - [ISLSBridgedCopyWindowsWithOptionsAndTagsOperation.Owner]
//   - [ISLSBridgedCopyWindowsWithOptionsAndTagsOperation.Spaces]
//   - [ISLSBridgedCopyWindowsWithOptionsAndTagsOperation.InitWithOwnerSpacesOptionsSetTagsClearedTags]
type ISLSBridgedCopyWindowsWithOptionsAndTagsOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	GetClearedTags(tags unsafe.Pointer)
	GetSetTags(tags unsafe.Pointer)
	MakeResultWithNumbers(numbers objectivec.IObject) objectivec.IObject
	Options() uint32
	Owner() uint32
	Spaces() foundation.INSArray
	InitWithOwnerSpacesOptionsSetTagsClearedTags(owner uint32, spaces objectivec.IObject, options uint32, tags unsafe.Pointer, tags2 unsafe.Pointer) SLSBridgedCopyWindowsWithOptionsAndTagsOperation
}

// Init initializes the instance.
func (s SLSBridgedCopyWindowsWithOptionsAndTagsOperation) Init() SLSBridgedCopyWindowsWithOptionsAndTagsOperation {
	rv := objc.Send[SLSBridgedCopyWindowsWithOptionsAndTagsOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedCopyWindowsWithOptionsAndTagsOperation) Autorelease() SLSBridgedCopyWindowsWithOptionsAndTagsOperation {
	rv := objc.Send[SLSBridgedCopyWindowsWithOptionsAndTagsOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedCopyWindowsWithOptionsAndTagsOperation creates a new SLSBridgedCopyWindowsWithOptionsAndTagsOperation instance.
func NewSLSBridgedCopyWindowsWithOptionsAndTagsOperation() SLSBridgedCopyWindowsWithOptionsAndTagsOperation {
	class := getSLSBridgedCopyWindowsWithOptionsAndTagsOperationClass()
	rv := objc.Send[SLSBridgedCopyWindowsWithOptionsAndTagsOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLSBridgedCopyWindowsWithOptionsAndTagsOperationWithCoder(coder objectivec.IObject) SLSBridgedCopyWindowsWithOptionsAndTagsOperation {
	instance := getSLSBridgedCopyWindowsWithOptionsAndTagsOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedCopyWindowsWithOptionsAndTagsOperationFromID(rv)
}

func NewSLSBridgedCopyWindowsWithOptionsAndTagsOperationWithOwnerSpacesOptionsSetTagsClearedTags(owner uint32, spaces objectivec.IObject, options uint32, tags unsafe.Pointer, tags2 unsafe.Pointer) SLSBridgedCopyWindowsWithOptionsAndTagsOperation {
	instance := getSLSBridgedCopyWindowsWithOptionsAndTagsOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithOwner:spaces:options:setTags:clearedTags:"), owner, spaces, options, tags, tags2)
	return SLSBridgedCopyWindowsWithOptionsAndTagsOperationFromID(rv)
}

func (s SLSBridgedCopyWindowsWithOptionsAndTagsOperation) GetClearedTags(tags unsafe.Pointer) {
	objc.Send[objc.ID](s.ID, objc.Sel("getClearedTags:"), tags)
}
func (s SLSBridgedCopyWindowsWithOptionsAndTagsOperation) GetSetTags(tags unsafe.Pointer) {
	objc.Send[objc.ID](s.ID, objc.Sel("getSetTags:"), tags)
}
func (s SLSBridgedCopyWindowsWithOptionsAndTagsOperation) MakeResultWithNumbers(numbers objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeResultWithNumbers:"), numbers)
	return objectivec.Object{ID: rv}
}
func (s SLSBridgedCopyWindowsWithOptionsAndTagsOperation) InitWithOwnerSpacesOptionsSetTagsClearedTags(owner uint32, spaces objectivec.IObject, options uint32, tags unsafe.Pointer, tags2 unsafe.Pointer) SLSBridgedCopyWindowsWithOptionsAndTagsOperation {
	rv := objc.Send[SLSBridgedCopyWindowsWithOptionsAndTagsOperation](s.ID, objc.Sel("initWithOwner:spaces:options:setTags:clearedTags:"), owner, spaces, options, tags, tags2)
	return rv
}

func (s SLSBridgedCopyWindowsWithOptionsAndTagsOperation) Options() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("options"))
	return rv
}
func (s SLSBridgedCopyWindowsWithOptionsAndTagsOperation) Owner() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("owner"))
	return rv
}
func (s SLSBridgedCopyWindowsWithOptionsAndTagsOperation) Spaces() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("spaces"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
