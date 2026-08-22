// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLOptimizationHints] class.
var (
	_MLOptimizationHintsClass     MLOptimizationHintsClass
	_MLOptimizationHintsClassOnce sync.Once
)

func getMLOptimizationHintsClass() MLOptimizationHintsClass {
	_MLOptimizationHintsClassOnce.Do(func() {
		_MLOptimizationHintsClass = MLOptimizationHintsClass{class: objc.GetClass("MLOptimizationHints")}
	})
	return _MLOptimizationHintsClass
}

// GetMLOptimizationHintsClass returns the class object for MLOptimizationHints.
func GetMLOptimizationHintsClass() MLOptimizationHintsClass {
	return getMLOptimizationHintsClass()
}

type MLOptimizationHintsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLOptimizationHintsClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLOptimizationHintsClass) Alloc() MLOptimizationHints {
	rv := objc.SendIfResponds[MLOptimizationHints](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLOptimizationHints.HotHandDuration]
//   - [MLOptimizationHints.SetHotHandDuration]
//   - [MLOptimizationHints.ReshapeFrequencyToString]
//   - [MLOptimizationHints.SpecializationStrategyToString]
//   - [MLOptimizationHints.InitWithCoder]
type MLOptimizationHints struct {
	objectivec.Object
}

// MLOptimizationHintsFromID constructs a [MLOptimizationHints] from an objc.ID.
func MLOptimizationHintsFromID(id objc.ID) MLOptimizationHints {
	return MLOptimizationHints{objectivec.Object{ID: id}}
}

// Ensure MLOptimizationHints implements IMLOptimizationHints.
var _ IMLOptimizationHints = MLOptimizationHints{}

// An interface definition for the [MLOptimizationHints] class.
//
// # Methods
//
//   - [IMLOptimizationHints.HotHandDuration]
//   - [IMLOptimizationHints.SetHotHandDuration]
//   - [IMLOptimizationHints.ReshapeFrequencyToString]
//   - [IMLOptimizationHints.SpecializationStrategyToString]
//   - [IMLOptimizationHints.InitWithCoder]
type IMLOptimizationHints interface {
	objectivec.IObject

	// Topic: Methods

	HotHandDuration() float64
	SetHotHandDuration(value float64)
	ReshapeFrequencyToString(string_ int64) objectivec.IObject
	SpecializationStrategyToString(string_ int64) objectivec.IObject
	InitWithCoder(coder foundation.INSCoder) MLOptimizationHints
}

// Init initializes the instance.
func (m MLOptimizationHints) Init() MLOptimizationHints {
	rv := objc.SendIfResponds[MLOptimizationHints](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLOptimizationHints) Autorelease() MLOptimizationHints {
	rv := objc.SendIfResponds[MLOptimizationHints](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLOptimizationHints creates a new MLOptimizationHints instance.
func NewMLOptimizationHints() MLOptimizationHints {
	class := getMLOptimizationHintsClass()
	rv := objc.SendIfResponds[MLOptimizationHints](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewOptimizationHintsWithCoder(coder objectivec.IObject) MLOptimizationHints {
	instance := getMLOptimizationHintsClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLOptimizationHintsFromID(rv)
}

func (m MLOptimizationHints) ReshapeFrequencyToString(string_ int64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("reshapeFrequencyToString:"), string_)
	return objectivec.Object{ID: rv}
}
func (m MLOptimizationHints) SpecializationStrategyToString(string_ int64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("specializationStrategyToString:"), string_)
	return objectivec.Object{ID: rv}
}
func (m MLOptimizationHints) InitWithCoder(coder foundation.INSCoder) MLOptimizationHints {
	rv := objc.SendIfResponds[MLOptimizationHints](m.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

func (_MLOptimizationHintsClass MLOptimizationHintsClass) SupportsSecureCoding() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_MLOptimizationHintsClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (m MLOptimizationHints) HotHandDuration() float64 {
	rv := objc.SendIfResponds[float64](m.ID, objc.Sel("hotHandDuration"))
	return rv
}
func (m MLOptimizationHints) SetHotHandDuration(value float64) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setHotHandDuration:"), value)
}
