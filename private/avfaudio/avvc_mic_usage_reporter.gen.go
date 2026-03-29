// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVVCMicUsageReporter] class.
var (
	_AVVCMicUsageReporterClass     AVVCMicUsageReporterClass
	_AVVCMicUsageReporterClassOnce sync.Once
)

func getAVVCMicUsageReporterClass() AVVCMicUsageReporterClass {
	_AVVCMicUsageReporterClassOnce.Do(func() {
		_AVVCMicUsageReporterClass = AVVCMicUsageReporterClass{class: objc.GetClass("AVVCMicUsageReporter")}
	})
	return _AVVCMicUsageReporterClass
}

// GetAVVCMicUsageReporterClass returns the class object for AVVCMicUsageReporter.
func GetAVVCMicUsageReporterClass() AVVCMicUsageReporterClass {
	return getAVVCMicUsageReporterClass()
}

type AVVCMicUsageReporterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVVCMicUsageReporterClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVVCMicUsageReporterClass) Alloc() AVVCMicUsageReporter {
	rv := objc.Send[AVVCMicUsageReporter](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [AVVCMicUsageReporter._getAuditToken]
//   - [AVVCMicUsageReporter.ReportMicUsage]
// See: https://developer.apple.com/documentation/AVFAudio/AVVCMicUsageReporter
type AVVCMicUsageReporter struct {
	objectivec.Object
}

// AVVCMicUsageReporterFromID constructs a [AVVCMicUsageReporter] from an objc.ID.
func AVVCMicUsageReporterFromID(id objc.ID) AVVCMicUsageReporter {
	return AVVCMicUsageReporter{objectivec.Object{ID: id}}
}
// Ensure AVVCMicUsageReporter implements IAVVCMicUsageReporter.
var _ IAVVCMicUsageReporter = AVVCMicUsageReporter{}

// An interface definition for the [AVVCMicUsageReporter] class.
//
// # Methods
//
//   - [IAVVCMicUsageReporter._getAuditToken]
//   - [IAVVCMicUsageReporter.ReportMicUsage]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCMicUsageReporter
type IAVVCMicUsageReporter interface {
	objectivec.IObject

	// Topic: Methods

	_getAuditToken(token objectivec.IObject) bool
	ReportMicUsage(usage bool)
}

// Init initializes the instance.
func (v AVVCMicUsageReporter) Init() AVVCMicUsageReporter {
	rv := objc.Send[AVVCMicUsageReporter](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v AVVCMicUsageReporter) Autorelease() AVVCMicUsageReporter {
	rv := objc.Send[AVVCMicUsageReporter](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVCMicUsageReporter creates a new AVVCMicUsageReporter instance.
func NewAVVCMicUsageReporter() AVVCMicUsageReporter {
	class := getAVVCMicUsageReporterClass()
	rv := objc.Send[AVVCMicUsageReporter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCMicUsageReporter/_getAuditToken:
func (v AVVCMicUsageReporter) _getAuditToken(token objectivec.IObject) bool {
	rv := objc.Send[bool](v.ID, objc.Sel("_getAuditToken:"), token)
	return rv
}

// GetAuditToken is an exported wrapper for the private method _getAuditToken.
func (v AVVCMicUsageReporter) GetAuditToken(token objectivec.IObject) bool {
	return v._getAuditToken(token)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCMicUsageReporter/reportMicUsage:
func (v AVVCMicUsageReporter) ReportMicUsage(usage bool) {
	objc.Send[objc.ID](v.ID, objc.Sel("reportMicUsage:"), usage)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCMicUsageReporter/sharedInstance
func (_AVVCMicUsageReporterClass AVVCMicUsageReporterClass) SharedInstance() AVVCMicUsageReporter {
	rv := objc.Send[objc.ID](objc.ID(_AVVCMicUsageReporterClass.class), objc.Sel("sharedInstance"))
	return AVVCMicUsageReporterFromID(rv)
}

