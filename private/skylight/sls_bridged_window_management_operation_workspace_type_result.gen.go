// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedWindowManagementOperationWorkspaceTypeResult] class.
var (
	_SLSBridgedWindowManagementOperationWorkspaceTypeResultClass     SLSBridgedWindowManagementOperationWorkspaceTypeResultClass
	_SLSBridgedWindowManagementOperationWorkspaceTypeResultClassOnce sync.Once
)

func getSLSBridgedWindowManagementOperationWorkspaceTypeResultClass() SLSBridgedWindowManagementOperationWorkspaceTypeResultClass {
	_SLSBridgedWindowManagementOperationWorkspaceTypeResultClassOnce.Do(func() {
		_SLSBridgedWindowManagementOperationWorkspaceTypeResultClass = SLSBridgedWindowManagementOperationWorkspaceTypeResultClass{class: objc.GetClass("SLSBridgedWindowManagementOperationWorkspaceTypeResult")}
	})
	return _SLSBridgedWindowManagementOperationWorkspaceTypeResultClass
}

// GetSLSBridgedWindowManagementOperationWorkspaceTypeResultClass returns the class object for SLSBridgedWindowManagementOperationWorkspaceTypeResult.
func GetSLSBridgedWindowManagementOperationWorkspaceTypeResultClass() SLSBridgedWindowManagementOperationWorkspaceTypeResultClass {
	return getSLSBridgedWindowManagementOperationWorkspaceTypeResultClass()
}

type SLSBridgedWindowManagementOperationWorkspaceTypeResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedWindowManagementOperationWorkspaceTypeResultClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedWindowManagementOperationWorkspaceTypeResultClass) Alloc() SLSBridgedWindowManagementOperationWorkspaceTypeResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationWorkspaceTypeResult](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedWindowManagementOperationWorkspaceTypeResult.WorkspaceType]
//   - [SLSBridgedWindowManagementOperationWorkspaceTypeResult.InitWithWorkspaceType]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationWorkspaceTypeResult
type SLSBridgedWindowManagementOperationWorkspaceTypeResult struct {
	SLSBridgedWindowManagementOperationResult
}

// SLSBridgedWindowManagementOperationWorkspaceTypeResultFromID constructs a [SLSBridgedWindowManagementOperationWorkspaceTypeResult] from an objc.ID.
func SLSBridgedWindowManagementOperationWorkspaceTypeResultFromID(id objc.ID) SLSBridgedWindowManagementOperationWorkspaceTypeResult {
	return SLSBridgedWindowManagementOperationWorkspaceTypeResult{SLSBridgedWindowManagementOperationResult: SLSBridgedWindowManagementOperationResultFromID(id)}
}

// Ensure SLSBridgedWindowManagementOperationWorkspaceTypeResult implements ISLSBridgedWindowManagementOperationWorkspaceTypeResult.
var _ ISLSBridgedWindowManagementOperationWorkspaceTypeResult = SLSBridgedWindowManagementOperationWorkspaceTypeResult{}

// An interface definition for the [SLSBridgedWindowManagementOperationWorkspaceTypeResult] class.
//
// # Methods
//
//   - [ISLSBridgedWindowManagementOperationWorkspaceTypeResult.WorkspaceType]
//   - [ISLSBridgedWindowManagementOperationWorkspaceTypeResult.InitWithWorkspaceType]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationWorkspaceTypeResult
type ISLSBridgedWindowManagementOperationWorkspaceTypeResult interface {
	ISLSBridgedWindowManagementOperationResult

	// Topic: Methods

	WorkspaceType() int
	InitWithWorkspaceType(type_ int) SLSBridgedWindowManagementOperationWorkspaceTypeResult
}

// Init initializes the instance.
func (s SLSBridgedWindowManagementOperationWorkspaceTypeResult) Init() SLSBridgedWindowManagementOperationWorkspaceTypeResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationWorkspaceTypeResult](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedWindowManagementOperationWorkspaceTypeResult) Autorelease() SLSBridgedWindowManagementOperationWorkspaceTypeResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationWorkspaceTypeResult](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedWindowManagementOperationWorkspaceTypeResult creates a new SLSBridgedWindowManagementOperationWorkspaceTypeResult instance.
func NewSLSBridgedWindowManagementOperationWorkspaceTypeResult() SLSBridgedWindowManagementOperationWorkspaceTypeResult {
	class := getSLSBridgedWindowManagementOperationWorkspaceTypeResultClass()
	rv := objc.Send[SLSBridgedWindowManagementOperationWorkspaceTypeResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationWorkspaceTypeResult/initWithCoder:
func NewSLSBridgedWindowManagementOperationWorkspaceTypeResultWithCoder(coder objectivec.IObject) SLSBridgedWindowManagementOperationWorkspaceTypeResult {
	instance := getSLSBridgedWindowManagementOperationWorkspaceTypeResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedWindowManagementOperationWorkspaceTypeResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationWorkspaceTypeResult/initWithWorkspaceType:
func NewSLSBridgedWindowManagementOperationWorkspaceTypeResultWithWorkspaceType(type_ int) SLSBridgedWindowManagementOperationWorkspaceTypeResult {
	instance := getSLSBridgedWindowManagementOperationWorkspaceTypeResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithWorkspaceType:"), type_)
	return SLSBridgedWindowManagementOperationWorkspaceTypeResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationWorkspaceTypeResult/initWithWorkspaceType:
func (s SLSBridgedWindowManagementOperationWorkspaceTypeResult) InitWithWorkspaceType(type_ int) SLSBridgedWindowManagementOperationWorkspaceTypeResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationWorkspaceTypeResult](s.ID, objc.Sel("initWithWorkspaceType:"), type_)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationWorkspaceTypeResult/workspaceType
func (s SLSBridgedWindowManagementOperationWorkspaceTypeResult) WorkspaceType() int {
	rv := objc.Send[int](s.ID, objc.Sel("workspaceType"))
	return rv
}
