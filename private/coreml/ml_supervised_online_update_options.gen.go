// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLSupervisedOnlineUpdateOptions] class.
var (
	_MLSupervisedOnlineUpdateOptionsClass     MLSupervisedOnlineUpdateOptionsClass
	_MLSupervisedOnlineUpdateOptionsClassOnce sync.Once
)

func getMLSupervisedOnlineUpdateOptionsClass() MLSupervisedOnlineUpdateOptionsClass {
	_MLSupervisedOnlineUpdateOptionsClassOnce.Do(func() {
		_MLSupervisedOnlineUpdateOptionsClass = MLSupervisedOnlineUpdateOptionsClass{class: objc.GetClass("MLSupervisedOnlineUpdateOptions")}
	})
	return _MLSupervisedOnlineUpdateOptionsClass
}

// GetMLSupervisedOnlineUpdateOptionsClass returns the class object for MLSupervisedOnlineUpdateOptions.
func GetMLSupervisedOnlineUpdateOptionsClass() MLSupervisedOnlineUpdateOptionsClass {
	return getMLSupervisedOnlineUpdateOptionsClass()
}

type MLSupervisedOnlineUpdateOptionsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLSupervisedOnlineUpdateOptionsClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLSupervisedOnlineUpdateOptionsClass) Alloc() MLSupervisedOnlineUpdateOptions {
	rv := objc.Send[MLSupervisedOnlineUpdateOptions](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSupervisedOnlineUpdateOptions
type MLSupervisedOnlineUpdateOptions struct {
	MLPredictionOptions
}

// MLSupervisedOnlineUpdateOptionsFromID constructs a [MLSupervisedOnlineUpdateOptions] from an objc.ID.
func MLSupervisedOnlineUpdateOptionsFromID(id objc.ID) MLSupervisedOnlineUpdateOptions {
	return MLSupervisedOnlineUpdateOptions{MLPredictionOptions: MLPredictionOptionsFromID(id)}
}

// Ensure MLSupervisedOnlineUpdateOptions implements IMLSupervisedOnlineUpdateOptions.
var _ IMLSupervisedOnlineUpdateOptions = MLSupervisedOnlineUpdateOptions{}

// An interface definition for the [MLSupervisedOnlineUpdateOptions] class.
//
// See: https://developer.apple.com/documentation/CoreML/MLSupervisedOnlineUpdateOptions
type IMLSupervisedOnlineUpdateOptions interface {
	IMLPredictionOptions
}

// Init initializes the instance.
func (s MLSupervisedOnlineUpdateOptions) Init() MLSupervisedOnlineUpdateOptions {
	rv := objc.Send[MLSupervisedOnlineUpdateOptions](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MLSupervisedOnlineUpdateOptions) Autorelease() MLSupervisedOnlineUpdateOptions {
	rv := objc.Send[MLSupervisedOnlineUpdateOptions](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLSupervisedOnlineUpdateOptions creates a new MLSupervisedOnlineUpdateOptions instance.
func NewMLSupervisedOnlineUpdateOptions() MLSupervisedOnlineUpdateOptions {
	class := getMLSupervisedOnlineUpdateOptionsClass()
	rv := objc.Send[MLSupervisedOnlineUpdateOptions](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionOptions/initWithCoder:
func NewSupervisedOnlineUpdateOptionsWithCoder(coder objectivec.IObject) MLSupervisedOnlineUpdateOptions {
	instance := getMLSupervisedOnlineUpdateOptionsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLSupervisedOnlineUpdateOptionsFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionOptions/initWithUsesCPUOnly:
func NewSupervisedOnlineUpdateOptionsWithUsesCPUOnly(cPUOnly bool) MLSupervisedOnlineUpdateOptions {
	instance := getMLSupervisedOnlineUpdateOptionsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUsesCPUOnly:"), cPUOnly)
	return MLSupervisedOnlineUpdateOptionsFromID(rv)
}
