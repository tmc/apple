// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAssetResourceLoadingRequestor] class.
var (
	_AVAssetResourceLoadingRequestorClass     AVAssetResourceLoadingRequestorClass
	_AVAssetResourceLoadingRequestorClassOnce sync.Once
)

func getAVAssetResourceLoadingRequestorClass() AVAssetResourceLoadingRequestorClass {
	_AVAssetResourceLoadingRequestorClassOnce.Do(func() {
		_AVAssetResourceLoadingRequestorClass = AVAssetResourceLoadingRequestorClass{class: objc.GetClass("AVAssetResourceLoadingRequestor")}
	})
	return _AVAssetResourceLoadingRequestorClass
}

// GetAVAssetResourceLoadingRequestorClass returns the class object for AVAssetResourceLoadingRequestor.
func GetAVAssetResourceLoadingRequestorClass() AVAssetResourceLoadingRequestorClass {
	return getAVAssetResourceLoadingRequestorClass()
}

type AVAssetResourceLoadingRequestorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAssetResourceLoadingRequestorClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetResourceLoadingRequestorClass) Alloc() AVAssetResourceLoadingRequestor {
	rv := objc.Send[AVAssetResourceLoadingRequestor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that contains information about the originator of a
// resource-loading request.
//
// # Retrieving expired session reports
//
//   - [AVAssetResourceLoadingRequestor.ProvidesExpiredSessionReports]: A Boolean value that indicates whether the requestor provides expired session reports.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingRequestor
type AVAssetResourceLoadingRequestor struct {
	objectivec.Object
}

// AVAssetResourceLoadingRequestorFromID constructs a [AVAssetResourceLoadingRequestor] from an objc.ID.
//
// An object that contains information about the originator of a
// resource-loading request.
func AVAssetResourceLoadingRequestorFromID(id objc.ID) AVAssetResourceLoadingRequestor {
	return AVAssetResourceLoadingRequestor{objectivec.Object{ID: id}}
}
// NOTE: AVAssetResourceLoadingRequestor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetResourceLoadingRequestor] class.
//
// # Retrieving expired session reports
//
//   - [IAVAssetResourceLoadingRequestor.ProvidesExpiredSessionReports]: A Boolean value that indicates whether the requestor provides expired session reports.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingRequestor
type IAVAssetResourceLoadingRequestor interface {
	objectivec.IObject

	// Topic: Retrieving expired session reports

	// A Boolean value that indicates whether the requestor provides expired session reports.
	ProvidesExpiredSessionReports() bool
}

// Init initializes the instance.
func (a AVAssetResourceLoadingRequestor) Init() AVAssetResourceLoadingRequestor {
	rv := objc.Send[AVAssetResourceLoadingRequestor](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetResourceLoadingRequestor) Autorelease() AVAssetResourceLoadingRequestor {
	rv := objc.Send[AVAssetResourceLoadingRequestor](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetResourceLoadingRequestor creates a new AVAssetResourceLoadingRequestor instance.
func NewAVAssetResourceLoadingRequestor() AVAssetResourceLoadingRequestor {
	class := getAVAssetResourceLoadingRequestorClass()
	rv := objc.Send[AVAssetResourceLoadingRequestor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Boolean value that indicates whether the requestor provides expired
// session reports.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingRequestor/providesExpiredSessionReports
func (a AVAssetResourceLoadingRequestor) ProvidesExpiredSessionReports() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("providesExpiredSessionReports"))
	return rv
}

