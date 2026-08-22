// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVAssetDownloadURLSession] class.
var (
	_AVAssetDownloadURLSessionClass     AVAssetDownloadURLSessionClass
	_AVAssetDownloadURLSessionClassOnce sync.Once
)

func getAVAssetDownloadURLSessionClass() AVAssetDownloadURLSessionClass {
	_AVAssetDownloadURLSessionClassOnce.Do(func() {
		_AVAssetDownloadURLSessionClass = AVAssetDownloadURLSessionClass{class: objc.GetClass("AVAssetDownloadURLSession")}
	})
	return _AVAssetDownloadURLSessionClass
}

// GetAVAssetDownloadURLSessionClass returns the class object for AVAssetDownloadURLSession.
func GetAVAssetDownloadURLSessionClass() AVAssetDownloadURLSessionClass {
	return getAVAssetDownloadURLSessionClass()
}

type AVAssetDownloadURLSessionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAssetDownloadURLSessionClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetDownloadURLSessionClass) Alloc() AVAssetDownloadURLSession {
	rv := objc.Send[AVAssetDownloadURLSession](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A URL session that creates and executes asset download tasks.
//
// # Creating download tasks
//
//   - [AVAssetDownloadURLSession.AssetDownloadTaskWithConfiguration]: Creates a download task that uses the specified configuration.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadURLSession
type AVAssetDownloadURLSession struct {
	foundation.NSURLSession
}

// AVAssetDownloadURLSessionFromID constructs a [AVAssetDownloadURLSession] from an objc.ID.
//
// A URL session that creates and executes asset download tasks.
func AVAssetDownloadURLSessionFromID(id objc.ID) AVAssetDownloadURLSession {
	return AVAssetDownloadURLSession{NSURLSession: foundation.NSURLSessionFromID(id)}
}

// NOTE: AVAssetDownloadURLSession adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetDownloadURLSession] class.
//
// # Creating download tasks
//
//   - [IAVAssetDownloadURLSession.AssetDownloadTaskWithConfiguration]: Creates a download task that uses the specified configuration.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadURLSession
type IAVAssetDownloadURLSession interface {
	foundation.INSURLSession

	// Topic: Creating download tasks

	// Creates a download task that uses the specified configuration.
	AssetDownloadTaskWithConfiguration(downloadConfiguration IAVAssetDownloadConfiguration) IAVAssetDownloadTask
}

// Init initializes the instance.
func (a AVAssetDownloadURLSession) Init() AVAssetDownloadURLSession {
	rv := objc.Send[AVAssetDownloadURLSession](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetDownloadURLSession) Autorelease() AVAssetDownloadURLSession {
	rv := objc.Send[AVAssetDownloadURLSession](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetDownloadURLSession creates a new AVAssetDownloadURLSession instance.
func NewAVAssetDownloadURLSession() AVAssetDownloadURLSession {
	class := getAVAssetDownloadURLSessionClass()
	rv := objc.Send[AVAssetDownloadURLSession](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a URL session to download assets.
//
// configuration: The configuration for this download session. The configuration you provide
// must be a background configuration or the system raises an exception.
//
// delegate: The delegate object to handle asset download progress updates and other
// session related events.
//
// delegateQueue: The queue to receive delegate callbacks on. If you specify `nil`, the
// system provides a serial queue.
//
// # Return Value
//
// A new download session.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadURLSession/init(configuration:assetDownloadDelegate:delegateQueue:)
func NewAssetDownloadURLSessionWithConfigurationAssetDownloadDelegateDelegateQueue(configuration foundation.NSURLSessionConfiguration, delegate AVAssetDownloadDelegate, delegateQueue foundation.NSOperationQueue) AVAssetDownloadURLSession {
	rv := objc.Send[objc.ID](objc.ID(getAVAssetDownloadURLSessionClass().class), objc.Sel("sessionWithConfiguration:assetDownloadDelegate:delegateQueue:"), configuration, delegate, delegateQueue)
	return AVAssetDownloadURLSessionFromID(rv)
}

// Creates a download task that uses the specified configuration.
//
// downloadConfiguration: The configuration that the task uses.
//
// # Return Value
//
// A new download task.
//
// # Discussion
//
// This method raises an exception if you call it on an invalidated session.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadURLSession/makeAssetDownloadTask(downloadConfiguration:)
func (a AVAssetDownloadURLSession) AssetDownloadTaskWithConfiguration(downloadConfiguration IAVAssetDownloadConfiguration) IAVAssetDownloadTask {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("assetDownloadTaskWithConfiguration:"), downloadConfiguration)
	return AVAssetDownloadTaskFromID(rv)
}
