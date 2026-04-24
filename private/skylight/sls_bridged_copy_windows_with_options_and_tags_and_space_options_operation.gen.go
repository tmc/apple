// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

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
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation
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
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation
type ISLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	GetClearedTags(tags objectivec.IObject)
	GetSetTags(tags objectivec.IObject)
	MakeResultWithNumbers(numbers objectivec.IObject) objectivec.IObject
	Options() uint32
	Owner() uint32
	SpaceOptions() uint32
	InitWithOwnerSpaceOptionsOptionsSetTagsClearedTags(owner uint32, options uint32, options2 uint32, tags objectivec.IObject, tags2 objectivec.IObject) SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation
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

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation/initWithCoder:
func NewSLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperationWithCoder(coder objectivec.IObject) SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation {
	instance := getSLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation/initWithOwner:spaceOptions:options:setTags:clearedTags:
func NewSLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperationWithOwnerSpaceOptionsOptionsSetTagsClearedTags(owner uint32, options uint32, options2 uint32, tags objectivec.IObject, tags2 objectivec.IObject) SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation {
	instance := getSLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithOwner:spaceOptions:options:setTags:clearedTags:"), owner, options, options2, tags, tags2)
	return SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation/getClearedTags:
func (s SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation) GetClearedTags(tags objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("getClearedTags:"), tags)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation/getSetTags:
func (s SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation) GetSetTags(tags objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("getSetTags:"), tags)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation/makeResultWithNumbers:
func (s SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation) MakeResultWithNumbers(numbers objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeResultWithNumbers:"), numbers)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation/initWithOwner:spaceOptions:options:setTags:clearedTags:
func (s SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation) InitWithOwnerSpaceOptionsOptionsSetTagsClearedTags(owner uint32, options uint32, options2 uint32, tags objectivec.IObject, tags2 objectivec.IObject) SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation {
	rv := objc.Send[SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation](s.ID, objc.Sel("initWithOwner:spaceOptions:options:setTags:clearedTags:"), owner, options, options2, tags, tags2)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation/options
func (s SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation) Options() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("options"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation/owner
func (s SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation) Owner() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("owner"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation/spaceOptions
func (s SLSBridgedCopyWindowsWithOptionsAndTagsAndSpaceOptionsOperation) SpaceOptions() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("spaceOptions"))
	return rv
}
