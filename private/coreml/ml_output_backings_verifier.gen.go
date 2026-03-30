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
//
// See: https://developer.apple.com/documentation/CoreML/MLOutputBackingsVerifier
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
//
// See: https://developer.apple.com/documentation/CoreML/MLOutputBackingsVerifier
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
func (o MLOutputBackingsVerifier) Init() MLOutputBackingsVerifier {
	rv := objc.Send[MLOutputBackingsVerifier](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o MLOutputBackingsVerifier) Autorelease() MLOutputBackingsVerifier {
	rv := objc.Send[MLOutputBackingsVerifier](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLOutputBackingsVerifier creates a new MLOutputBackingsVerifier instance.
func NewMLOutputBackingsVerifier() MLOutputBackingsVerifier {
	class := getMLOutputBackingsVerifierClass()
	rv := objc.Send[MLOutputBackingsVerifier](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLOutputBackingsVerifier/initWithOutputDescriptions:
func NewOutputBackingsVerifierWithOutputDescriptions(descriptions objectivec.IObject) MLOutputBackingsVerifier {
	instance := getMLOutputBackingsVerifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithOutputDescriptions:"), descriptions)
	return MLOutputBackingsVerifierFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLOutputBackingsVerifier/_verifyMultiArrayOutputBacking:forFeature:error:
func (o MLOutputBackingsVerifier) _verifyMultiArrayOutputBackingForFeatureError(backing objectivec.IObject, feature objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](o.ID, objc.Sel("_verifyMultiArrayOutputBacking:forFeature:error:"), backing, feature, unsafe.Pointer(&errorPtr))
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
func (o MLOutputBackingsVerifier) VerifyMultiArrayOutputBackingForFeatureError(backing objectivec.IObject, feature objectivec.IObject) (bool, error) {
	return o._verifyMultiArrayOutputBackingForFeatureError(backing, feature)
}

// See: https://developer.apple.com/documentation/CoreML/MLOutputBackingsVerifier/_verifyOutputBacking:forFeature:error:
func (o MLOutputBackingsVerifier) _verifyOutputBackingForFeatureError(backing objectivec.IObject, feature objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](o.ID, objc.Sel("_verifyOutputBacking:forFeature:error:"), backing, feature, unsafe.Pointer(&errorPtr))
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
func (o MLOutputBackingsVerifier) VerifyOutputBackingForFeatureError(backing objectivec.IObject, feature objectivec.IObject) (bool, error) {
	return o._verifyOutputBackingForFeatureError(backing, feature)
}

// See: https://developer.apple.com/documentation/CoreML/MLOutputBackingsVerifier/_verifyPixelBufferOutputBacking:forFeature:error:
func (o MLOutputBackingsVerifier) _verifyPixelBufferOutputBackingForFeatureError(backing corevideo.CVImageBufferRef, feature objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](o.ID, objc.Sel("_verifyPixelBufferOutputBacking:forFeature:error:"), backing, feature, unsafe.Pointer(&errorPtr))
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
func (o MLOutputBackingsVerifier) VerifyPixelBufferOutputBackingForFeatureError(backing corevideo.CVImageBufferRef, feature objectivec.IObject) (bool, error) {
	return o._verifyPixelBufferOutputBackingForFeatureError(backing, feature)
}

// See: https://developer.apple.com/documentation/CoreML/MLOutputBackingsVerifier/verifyOutputBackings:predictionUsesBatch:error:
func (o MLOutputBackingsVerifier) VerifyOutputBackingsPredictionUsesBatchError(backings objectivec.IObject, batch bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](o.ID, objc.Sel("verifyOutputBackings:predictionUsesBatch:error:"), backings, batch, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("verifyOutputBackings:predictionUsesBatch:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLOutputBackingsVerifier/initWithOutputDescriptions:
func (o MLOutputBackingsVerifier) InitWithOutputDescriptions(descriptions objectivec.IObject) MLOutputBackingsVerifier {
	rv := objc.Send[MLOutputBackingsVerifier](o.ID, objc.Sel("initWithOutputDescriptions:"), descriptions)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLOutputBackingsVerifier/outputDescriptions
func (o MLOutputBackingsVerifier) OutputDescriptions() foundation.INSDictionary {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputDescriptions"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
