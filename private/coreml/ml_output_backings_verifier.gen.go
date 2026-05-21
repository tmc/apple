// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLOutputBackingsVerifier] class.
var (
	_MLOutputBackingsVerifierClass     MLOutputBackingsVerifierClass
	_MLOutputBackingsVerifierClassOnce sync.Once
)

func getMLOutputBackingsVerifierClass() MLOutputBackingsVerifierClass {
	_MLOutputBackingsVerifierClassOnce.Do(func() {
		_MLOutputBackingsVerifierClass = MLOutputBackingsVerifierClass{class: objc.GetClass("MLOutputBackingsVerifier")}
	})
	return _MLOutputBackingsVerifierClass
}

// GetMLOutputBackingsVerifierClass returns the class object for MLOutputBackingsVerifier.
func GetMLOutputBackingsVerifierClass() MLOutputBackingsVerifierClass {
	return getMLOutputBackingsVerifierClass()
}

type MLOutputBackingsVerifierClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLOutputBackingsVerifierClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLOutputBackingsVerifierClass) Alloc() MLOutputBackingsVerifier {
	rv := objc.Send[MLOutputBackingsVerifier](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLOutputBackingsVerifier._verifyMultiArrayOutputBackingForFeatureError]
//   - [MLOutputBackingsVerifier._verifyOutputBackingForFeatureError]
//   - [MLOutputBackingsVerifier._verifyPixelBufferOutputBackingForFeatureError]
//   - [MLOutputBackingsVerifier.OutputDescriptions]
//   - [MLOutputBackingsVerifier.VerifyOutputBackingsPredictionUsesBatchError]
//   - [MLOutputBackingsVerifier.InitWithOutputDescriptions]
type MLOutputBackingsVerifier struct {
	objectivec.Object
}

// MLOutputBackingsVerifierFromID constructs a [MLOutputBackingsVerifier] from an objc.ID.
func MLOutputBackingsVerifierFromID(id objc.ID) MLOutputBackingsVerifier {
	return MLOutputBackingsVerifier{objectivec.Object{ID: id}}
}

// Ensure MLOutputBackingsVerifier implements IMLOutputBackingsVerifier.
var _ IMLOutputBackingsVerifier = MLOutputBackingsVerifier{}

// An interface definition for the [MLOutputBackingsVerifier] class.
//
// # Methods
//
//   - [IMLOutputBackingsVerifier._verifyMultiArrayOutputBackingForFeatureError]
//   - [IMLOutputBackingsVerifier._verifyOutputBackingForFeatureError]
//   - [IMLOutputBackingsVerifier._verifyPixelBufferOutputBackingForFeatureError]
//   - [IMLOutputBackingsVerifier.OutputDescriptions]
//   - [IMLOutputBackingsVerifier.VerifyOutputBackingsPredictionUsesBatchError]
//   - [IMLOutputBackingsVerifier.InitWithOutputDescriptions]
type IMLOutputBackingsVerifier interface {
	objectivec.IObject

	// Topic: Methods

	_verifyMultiArrayOutputBackingForFeatureError(backing objectivec.IObject, feature objectivec.IObject) (bool, error)
	_verifyOutputBackingForFeatureError(backing objectivec.IObject, feature objectivec.IObject) (bool, error)
	_verifyPixelBufferOutputBackingForFeatureError(backing corevideo.CVImageBufferRef, feature objectivec.IObject) (bool, error)
	OutputDescriptions() foundation.INSDictionary
	VerifyOutputBackingsPredictionUsesBatchError(backings objectivec.IObject, batch bool) (bool, error)
	InitWithOutputDescriptions(descriptions objectivec.IObject) MLOutputBackingsVerifier
}

// Init initializes the instance.
func (m MLOutputBackingsVerifier) Init() MLOutputBackingsVerifier {
	rv := objc.Send[MLOutputBackingsVerifier](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLOutputBackingsVerifier) Autorelease() MLOutputBackingsVerifier {
	rv := objc.Send[MLOutputBackingsVerifier](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLOutputBackingsVerifier creates a new MLOutputBackingsVerifier instance.
func NewMLOutputBackingsVerifier() MLOutputBackingsVerifier {
	class := getMLOutputBackingsVerifierClass()
	rv := objc.Send[MLOutputBackingsVerifier](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewOutputBackingsVerifierWithOutputDescriptions(descriptions objectivec.IObject) MLOutputBackingsVerifier {
	instance := getMLOutputBackingsVerifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithOutputDescriptions:"), descriptions)
	return MLOutputBackingsVerifierFromID(rv)
}

func (m MLOutputBackingsVerifier) _verifyMultiArrayOutputBackingForFeatureError(backing objectivec.IObject, feature objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("_verifyMultiArrayOutputBacking:forFeature:error:"), backing, feature, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_verifyMultiArrayOutputBacking:forFeature:error: returned NO with nil NSError")
	}
	return rv, nil

}

// VerifyMultiArrayOutputBackingForFeatureError is an exported wrapper for the private method _verifyMultiArrayOutputBackingForFeatureError.
func (m MLOutputBackingsVerifier) VerifyMultiArrayOutputBackingForFeatureError(backing objectivec.IObject, feature objectivec.IObject) (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_verifyMultiArrayOutputBacking:forFeature:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_verifyMultiArrayOutputBacking:forFeature:error:"}
		return false, err
	}
	return m._verifyMultiArrayOutputBackingForFeatureError(backing, feature)
}

// CanVerifyMultiArrayOutputBackingForFeatureError reports whether the receiver responds to the private selector _verifyMultiArrayOutputBacking:forFeature:error:.
func (m MLOutputBackingsVerifier) CanVerifyMultiArrayOutputBackingForFeatureError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_verifyMultiArrayOutputBacking:forFeature:error:"))
}
func (m MLOutputBackingsVerifier) _verifyOutputBackingForFeatureError(backing objectivec.IObject, feature objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("_verifyOutputBacking:forFeature:error:"), backing, feature, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_verifyOutputBacking:forFeature:error: returned NO with nil NSError")
	}
	return rv, nil

}

// VerifyOutputBackingForFeatureError is an exported wrapper for the private method _verifyOutputBackingForFeatureError.
func (m MLOutputBackingsVerifier) VerifyOutputBackingForFeatureError(backing objectivec.IObject, feature objectivec.IObject) (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_verifyOutputBacking:forFeature:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_verifyOutputBacking:forFeature:error:"}
		return false, err
	}
	return m._verifyOutputBackingForFeatureError(backing, feature)
}

// CanVerifyOutputBackingForFeatureError reports whether the receiver responds to the private selector _verifyOutputBacking:forFeature:error:.
func (m MLOutputBackingsVerifier) CanVerifyOutputBackingForFeatureError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_verifyOutputBacking:forFeature:error:"))
}
func (m MLOutputBackingsVerifier) _verifyPixelBufferOutputBackingForFeatureError(backing corevideo.CVImageBufferRef, feature objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("_verifyPixelBufferOutputBacking:forFeature:error:"), backing, feature, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_verifyPixelBufferOutputBacking:forFeature:error: returned NO with nil NSError")
	}
	return rv, nil

}

// VerifyPixelBufferOutputBackingForFeatureError is an exported wrapper for the private method _verifyPixelBufferOutputBackingForFeatureError.
func (m MLOutputBackingsVerifier) VerifyPixelBufferOutputBackingForFeatureError(backing corevideo.CVImageBufferRef, feature objectivec.IObject) (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_verifyPixelBufferOutputBacking:forFeature:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_verifyPixelBufferOutputBacking:forFeature:error:"}
		return false, err
	}
	return m._verifyPixelBufferOutputBackingForFeatureError(backing, feature)
}

// CanVerifyPixelBufferOutputBackingForFeatureError reports whether the receiver responds to the private selector _verifyPixelBufferOutputBacking:forFeature:error:.
func (m MLOutputBackingsVerifier) CanVerifyPixelBufferOutputBackingForFeatureError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_verifyPixelBufferOutputBacking:forFeature:error:"))
}
func (m MLOutputBackingsVerifier) VerifyOutputBackingsPredictionUsesBatchError(backings objectivec.IObject, batch bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("verifyOutputBackings:predictionUsesBatch:error:"), backings, batch, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("verifyOutputBackings:predictionUsesBatch:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLOutputBackingsVerifier) InitWithOutputDescriptions(descriptions objectivec.IObject) MLOutputBackingsVerifier {
	rv := objc.Send[MLOutputBackingsVerifier](m.ID, objc.Sel("initWithOutputDescriptions:"), descriptions)
	return rv
}

func (m MLOutputBackingsVerifier) OutputDescriptions() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("outputDescriptions"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
