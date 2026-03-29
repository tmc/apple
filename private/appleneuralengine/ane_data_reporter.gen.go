// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANEDataReporter] class.
var (
	_ANEDataReporterClass     ANEDataReporterClass
	_ANEDataReporterClassOnce sync.Once
)

func getANEDataReporterClass() ANEDataReporterClass {
	_ANEDataReporterClassOnce.Do(func() {
		_ANEDataReporterClass = ANEDataReporterClass{class: objc.GetClass("_ANEDataReporter")}
	})
	return _ANEDataReporterClass
}

// GetANEDataReporterClass returns the class object for _ANEDataReporter.
func GetANEDataReporterClass() ANEDataReporterClass {
	return getANEDataReporterClass()
}

type ANEDataReporterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac ANEDataReporterClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEDataReporterClass) Alloc() ANEDataReporter {
	rv := objc.Send[ANEDataReporter](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDataReporter
type ANEDataReporter struct {
	objectivec.Object
}

// ANEDataReporterFromID constructs a [ANEDataReporter] from an objc.ID.
func ANEDataReporterFromID(id objc.ID) ANEDataReporter {
	return ANEDataReporter{objectivec.Object{ID: id}}
}
// Ensure ANEDataReporter implements IANEDataReporter.
var _ IANEDataReporter = ANEDataReporter{}

// An interface definition for the [ANEDataReporter] class.
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDataReporter
type IANEDataReporter interface {
	objectivec.IObject
}

// Init initializes the instance.
func (a ANEDataReporter) Init() ANEDataReporter {
	rv := objc.Send[ANEDataReporter](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEDataReporter) Autorelease() ANEDataReporter {
	rv := objc.Send[ANEDataReporter](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEDataReporter creates a new ANEDataReporter instance.
func NewANEDataReporter() ANEDataReporter {
	class := getANEDataReporterClass()
	rv := objc.Send[ANEDataReporter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDataReporter/addValue:forScalarKey:
func (_ANEDataReporterClass ANEDataReporterClass) AddValueForScalarKey(value int64, key objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_ANEDataReporterClass.class), objc.Sel("addValue:forScalarKey:"), value, key)
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDataReporter/analyticsKey:
func (_ANEDataReporterClass ANEDataReporterClass) AnalyticsKey(key objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEDataReporterClass.class), objc.Sel("analyticsKey:"), key)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDataReporter/reportClient:modelName:
func (_ANEDataReporterClass ANEDataReporterClass) ReportClientModelName(client objectivec.IObject, name objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEDataReporterClass.class), objc.Sel("reportClient:modelName:"), client, name)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDataReporter/reportErrorMsg:status:
func (_ANEDataReporterClass ANEDataReporterClass) ReportErrorMsgStatus(msg uint32, status uint32) {
	objc.Send[objc.ID](objc.ID(_ANEDataReporterClass.class), objc.Sel("reportErrorMsg:status:"), msg, status)
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDataReporter/reportTelemetryToPPS:playload:
func (_ANEDataReporterClass ANEDataReporterClass) ReportTelemetryToPPSPlayload(pps objectivec.IObject, playload objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_ANEDataReporterClass.class), objc.Sel("reportTelemetryToPPS:playload:"), pps, playload)
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDataReporter/reportTelemetryToCoreAnalytics:payload:
func (_ANEDataReporterClass ANEDataReporterClass) ReportTelemetryToCoreAnalyticsPayload(analytics objectivec.IObject, payload objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_ANEDataReporterClass.class), objc.Sel("reportTelemetryToCoreAnalytics:payload:"), analytics, payload)
}

