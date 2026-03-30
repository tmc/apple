// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLReporter] class.
var (
	_MLReporterClass     MLReporterClass
	_MLReporterClassOnce sync.Once
)

func getMLReporterClass() MLReporterClass {
	_MLReporterClassOnce.Do(func() {
		_MLReporterClass = MLReporterClass{class: objc.GetClass("MLReporter")}
	})
	return _MLReporterClass
}

// GetMLReporterClass returns the class object for MLReporter.
func GetMLReporterClass() MLReporterClass {
	return getMLReporterClass()
}

type MLReporterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLReporterClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLReporterClass) Alloc() MLReporter {
	rv := objc.Send[MLReporter](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLReporter.LogMetric]
//   - [MLReporter.DebugDescription]
//   - [MLReporter.Description]
//   - [MLReporter.Hash]
//   - [MLReporter.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLReporter
type MLReporter struct {
	objectivec.Object
}

// MLReporterFromID constructs a [MLReporter] from an objc.ID.
func MLReporterFromID(id objc.ID) MLReporter {
	return MLReporter{objectivec.Object{ID: id}}
}

// Ensure MLReporter implements IMLReporter.
var _ IMLReporter = MLReporter{}

// An interface definition for the [MLReporter] class.
//
// # Methods
//
//   - [IMLReporter.LogMetric]
//   - [IMLReporter.DebugDescription]
//   - [IMLReporter.Description]
//   - [IMLReporter.Hash]
//   - [IMLReporter.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLReporter
type IMLReporter interface {
	objectivec.IObject

	// Topic: Methods

	LogMetric(metric objectivec.IObject)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (r MLReporter) Init() MLReporter {
	rv := objc.Send[MLReporter](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MLReporter) Autorelease() MLReporter {
	rv := objc.Send[MLReporter](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLReporter creates a new MLReporter instance.
func NewMLReporter() MLReporter {
	class := getMLReporterClass()
	rv := objc.Send[MLReporter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLReporter/logMetric:
func (r MLReporter) LogMetric(metric objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("logMetric:"), metric)
}

// See: https://developer.apple.com/documentation/CoreML/MLReporter/reporter
func (_MLReporterClass MLReporterClass) Reporter() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLReporterClass.class), objc.Sel("reporter"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLReporter/debugDescription
func (r MLReporter) DebugDescription() string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLReporter/description
func (r MLReporter) Description() string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLReporter/hash
func (r MLReporter) Hash() uint64 {
	rv := objc.Send[uint64](r.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLReporter/superclass
func (r MLReporter) Superclass() objc.Class {
	rv := objc.Send[objc.Class](r.ID, objc.Sel("superclass"))
	return rv
}
