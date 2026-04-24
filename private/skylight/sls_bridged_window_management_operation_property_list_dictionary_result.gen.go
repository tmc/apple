// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedWindowManagementOperationPropertyListDictionaryResult] class.
var (
	_SLSBridgedWindowManagementOperationPropertyListDictionaryResultClass     SLSBridgedWindowManagementOperationPropertyListDictionaryResultClass
	_SLSBridgedWindowManagementOperationPropertyListDictionaryResultClassOnce sync.Once
)

func getSLSBridgedWindowManagementOperationPropertyListDictionaryResultClass() SLSBridgedWindowManagementOperationPropertyListDictionaryResultClass {
	_SLSBridgedWindowManagementOperationPropertyListDictionaryResultClassOnce.Do(func() {
		_SLSBridgedWindowManagementOperationPropertyListDictionaryResultClass = SLSBridgedWindowManagementOperationPropertyListDictionaryResultClass{class: objc.GetClass("SLSBridgedWindowManagementOperationPropertyListDictionaryResult")}
	})
	return _SLSBridgedWindowManagementOperationPropertyListDictionaryResultClass
}

// GetSLSBridgedWindowManagementOperationPropertyListDictionaryResultClass returns the class object for SLSBridgedWindowManagementOperationPropertyListDictionaryResult.
func GetSLSBridgedWindowManagementOperationPropertyListDictionaryResultClass() SLSBridgedWindowManagementOperationPropertyListDictionaryResultClass {
	return getSLSBridgedWindowManagementOperationPropertyListDictionaryResultClass()
}

type SLSBridgedWindowManagementOperationPropertyListDictionaryResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedWindowManagementOperationPropertyListDictionaryResultClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedWindowManagementOperationPropertyListDictionaryResultClass) Alloc() SLSBridgedWindowManagementOperationPropertyListDictionaryResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationPropertyListDictionaryResult](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedWindowManagementOperationPropertyListDictionaryResult.PropertyListDictionary]
//   - [SLSBridgedWindowManagementOperationPropertyListDictionaryResult.InitWithPropertyListDictionary]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationPropertyListDictionaryResult
type SLSBridgedWindowManagementOperationPropertyListDictionaryResult struct {
	SLSBridgedWindowManagementOperationResult
}

// SLSBridgedWindowManagementOperationPropertyListDictionaryResultFromID constructs a [SLSBridgedWindowManagementOperationPropertyListDictionaryResult] from an objc.ID.
func SLSBridgedWindowManagementOperationPropertyListDictionaryResultFromID(id objc.ID) SLSBridgedWindowManagementOperationPropertyListDictionaryResult {
	return SLSBridgedWindowManagementOperationPropertyListDictionaryResult{SLSBridgedWindowManagementOperationResult: SLSBridgedWindowManagementOperationResultFromID(id)}
}

// Ensure SLSBridgedWindowManagementOperationPropertyListDictionaryResult implements ISLSBridgedWindowManagementOperationPropertyListDictionaryResult.
var _ ISLSBridgedWindowManagementOperationPropertyListDictionaryResult = SLSBridgedWindowManagementOperationPropertyListDictionaryResult{}

// An interface definition for the [SLSBridgedWindowManagementOperationPropertyListDictionaryResult] class.
//
// # Methods
//
//   - [ISLSBridgedWindowManagementOperationPropertyListDictionaryResult.PropertyListDictionary]
//   - [ISLSBridgedWindowManagementOperationPropertyListDictionaryResult.InitWithPropertyListDictionary]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationPropertyListDictionaryResult
type ISLSBridgedWindowManagementOperationPropertyListDictionaryResult interface {
	ISLSBridgedWindowManagementOperationResult

	// Topic: Methods

	PropertyListDictionary() foundation.INSDictionary
	InitWithPropertyListDictionary(dictionary objectivec.IObject) SLSBridgedWindowManagementOperationPropertyListDictionaryResult
}

// Init initializes the instance.
func (s SLSBridgedWindowManagementOperationPropertyListDictionaryResult) Init() SLSBridgedWindowManagementOperationPropertyListDictionaryResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationPropertyListDictionaryResult](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedWindowManagementOperationPropertyListDictionaryResult) Autorelease() SLSBridgedWindowManagementOperationPropertyListDictionaryResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationPropertyListDictionaryResult](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedWindowManagementOperationPropertyListDictionaryResult creates a new SLSBridgedWindowManagementOperationPropertyListDictionaryResult instance.
func NewSLSBridgedWindowManagementOperationPropertyListDictionaryResult() SLSBridgedWindowManagementOperationPropertyListDictionaryResult {
	class := getSLSBridgedWindowManagementOperationPropertyListDictionaryResultClass()
	rv := objc.Send[SLSBridgedWindowManagementOperationPropertyListDictionaryResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationPropertyListDictionaryResult/initWithCoder:
func NewSLSBridgedWindowManagementOperationPropertyListDictionaryResultWithCoder(coder objectivec.IObject) SLSBridgedWindowManagementOperationPropertyListDictionaryResult {
	instance := getSLSBridgedWindowManagementOperationPropertyListDictionaryResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedWindowManagementOperationPropertyListDictionaryResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationPropertyListDictionaryResult/initWithPropertyListDictionary:
func NewSLSBridgedWindowManagementOperationPropertyListDictionaryResultWithPropertyListDictionary(dictionary objectivec.IObject) SLSBridgedWindowManagementOperationPropertyListDictionaryResult {
	instance := getSLSBridgedWindowManagementOperationPropertyListDictionaryResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPropertyListDictionary:"), dictionary)
	return SLSBridgedWindowManagementOperationPropertyListDictionaryResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationPropertyListDictionaryResult/initWithPropertyListDictionary:
func (s SLSBridgedWindowManagementOperationPropertyListDictionaryResult) InitWithPropertyListDictionary(dictionary objectivec.IObject) SLSBridgedWindowManagementOperationPropertyListDictionaryResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationPropertyListDictionaryResult](s.ID, objc.Sel("initWithPropertyListDictionary:"), dictionary)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationPropertyListDictionaryResult/propertyListDictionary
func (s SLSBridgedWindowManagementOperationPropertyListDictionaryResult) PropertyListDictionary() foundation.INSDictionary {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("propertyListDictionary"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
