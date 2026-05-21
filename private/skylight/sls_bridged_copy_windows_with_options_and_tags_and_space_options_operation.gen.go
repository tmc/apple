// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation] class.
var (
	_SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperationClass     SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperationClass
	_SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperationClassOnce sync.Once
)

func getSLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperationClass() SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperationClass {
	_SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperationClassOnce.Do(func() {
		_SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperationClass = SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperationClass{class: objc.GetClass("SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation")}
	})
	return _SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperationClass
}

// GetSLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperationClass returns the class object for SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation.
func GetSLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperationClass() SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperationClass {
	return getSLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperationClass()
}

type SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperationClass) Alloc() SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation {
	rv := objc.Send[SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation.GetClearedTags]
//   - [SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation.GetSetTags]
//   - [SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation.MakeResultWithNumbers]
//   - [SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation.Options]
//   - [SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation.Owner]
//   - [SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation.SpaceOptions]
//   - [SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation.InitWithOwnerSpaceOptionsOptionsSetTagsClearedTags]
type SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperationFromID constructs a [SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation] from an objc.ID.
func SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperationFromID(id objc.ID) SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation {
	return SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation implements ISLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation.
var _ ISLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation = SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation{}

// An interface definition for the [SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation] class.
//
// # Methods
//
//   - [ISLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation.GetClearedTags]
//   - [ISLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation.GetSetTags]
//   - [ISLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation.MakeResultWithNumbers]
//   - [ISLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation.Options]
//   - [ISLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation.Owner]
//   - [ISLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation.SpaceOptions]
//   - [ISLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation.InitWithOwnerSpaceOptionsOptionsSetTagsClearedTags]
type ISLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	GetClearedTags(tags unsafe.Pointer)
	GetSetTags(tags unsafe.Pointer)
	MakeResultWithNumbers(numbers objectivec.IObject) objectivec.IObject
	Options() uint32
	Owner() uint32
	SpaceOptions() uint32
	InitWithOwnerSpaceOptionsOptionsSetTagsClearedTags(owner uint32, options uint32, options2 uint32, tags unsafe.Pointer, tags2 unsafe.Pointer) SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation
}

// Init initializes the instance.
func (s SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation) Init() SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation {
	rv := objc.Send[SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation) Autorelease() SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation {
	rv := objc.Send[SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation creates a new SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation instance.
func NewSLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation() SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation {
	class := getSLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperationClass()
	rv := objc.Send[SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperationWithCoder(coder objectivec.IObject) SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation {
	instance := getSLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperationFromID(rv)
}

func NewSLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperationWithOwnerSpaceOptionsOptionsSetTagsClearedTags(owner uint32, options uint32, options2 uint32, tags unsafe.Pointer, tags2 unsafe.Pointer) SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation {
	instance := getSLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithOwner:spaceOptions:options:setTags:clearedTags:"), owner, options, options2, tags, tags2)
	return SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperationFromID(rv)
}

func (s SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation) GetClearedTags(tags unsafe.Pointer) {
	objc.Send[objc.ID](s.ID, objc.Sel("getClearedTags:"), tags)
}
func (s SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation) GetSetTags(tags unsafe.Pointer) {
	objc.Send[objc.ID](s.ID, objc.Sel("getSetTags:"), tags)
}
func (s SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation) MakeResultWithNumbers(numbers objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeResultWithNumbers:"), numbers)
	return objectivec.Object{ID: rv}
}
func (s SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation) InitWithOwnerSpaceOptionsOptionsSetTagsClearedTags(owner uint32, options uint32, options2 uint32, tags unsafe.Pointer, tags2 unsafe.Pointer) SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation {
	rv := objc.Send[SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation](s.ID, objc.Sel("initWithOwner:spaceOptions:options:setTags:clearedTags:"), owner, options, options2, tags, tags2)
	return rv
}

func (s SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation) Options() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("options"))
	return rv
}
func (s SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation) Owner() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("owner"))
	return rv
}
func (s SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation) SpaceOptions() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("spaceOptions"))
	return rv
}
