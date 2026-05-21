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
type IMLReporter interface {
	objectivec.IObject

	// Topic: Methods

	LogMetric(metric objectivec.IObject)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (m MLReporter) Init() MLReporter {
	rv := objc.Send[MLReporter](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLReporter) Autorelease() MLReporter {
	rv := objc.Send[MLReporter](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLReporter creates a new MLReporter instance.
func NewMLReporter() MLReporter {
	class := getMLReporterClass()
	rv := objc.Send[MLReporter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (m MLReporter) LogMetric(metric objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("logMetric:"), metric)
}

func (_MLReporterClass MLReporterClass) Reporter() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLReporterClass.class), objc.Sel("reporter"))
	return objectivec.Object{ID: rv}
}

func (m MLReporter) DebugDescription() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLReporter) Description() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLReporter) Hash() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("hash"))
	return rv
}
func (m MLReporter) Superclass() objectivec.Class {
	rv := objc.Send[objectivec.Class](m.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
