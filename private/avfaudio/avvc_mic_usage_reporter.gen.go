// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

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
	rv := objc.SendIfResponds[AVVCMicUsageReporter](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVVCMicUsageReporter._getAuditToken]
//   - [AVVCMicUsageReporter.ReportMicUsage]
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
type IAVVCMicUsageReporter interface {
	objectivec.IObject

	// Topic: Methods

	_getAuditToken(token unsafe.Pointer) bool
	ReportMicUsage(usage bool)
}

// Init initializes the instance.
func (a AVVCMicUsageReporter) Init() AVVCMicUsageReporter {
	rv := objc.SendIfResponds[AVVCMicUsageReporter](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVVCMicUsageReporter) Autorelease() AVVCMicUsageReporter {
	rv := objc.SendIfResponds[AVVCMicUsageReporter](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVCMicUsageReporter creates a new AVVCMicUsageReporter instance.
func NewAVVCMicUsageReporter() AVVCMicUsageReporter {
	class := getAVVCMicUsageReporterClass()
	rv := objc.SendIfResponds[AVVCMicUsageReporter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (a AVVCMicUsageReporter) _getAuditToken(token unsafe.Pointer) bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("_getAuditToken:"), token)
	return rv
}

// GetAuditToken is an exported wrapper for the private method _getAuditToken.
func (a AVVCMicUsageReporter) GetAuditToken(token unsafe.Pointer) (bool, error) {
	if !objc.RespondsToSelector(a.ID, objc.Sel("_getAuditToken:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_getAuditToken:"}
		return false, err
	}
	return a._getAuditToken(token), nil
}

// CanGetAuditToken reports whether the receiver responds to the private selector _getAuditToken:.
func (a AVVCMicUsageReporter) CanGetAuditToken() bool {
	return objc.RespondsToSelector(a.ID, objc.Sel("_getAuditToken:"))
}
func (a AVVCMicUsageReporter) ReportMicUsage(usage bool) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("reportMicUsage:"), usage)
}

func (_AVVCMicUsageReporterClass AVVCMicUsageReporterClass) SharedInstance() AVVCMicUsageReporter {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_AVVCMicUsageReporterClass.class), objc.Sel("sharedInstance"))
	return AVVCMicUsageReporterFromID(rv)
}
