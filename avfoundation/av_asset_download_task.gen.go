// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVAssetDownloadTask] class.
var (
	_AVAssetDownloadTaskClass     AVAssetDownloadTaskClass
	_AVAssetDownloadTaskClassOnce sync.Once
)

func getAVAssetDownloadTaskClass() AVAssetDownloadTaskClass {
	_AVAssetDownloadTaskClassOnce.Do(func() {
		_AVAssetDownloadTaskClass = AVAssetDownloadTaskClass{class: objc.GetClass("AVAssetDownloadTask")}
	})
	return _AVAssetDownloadTaskClass
}

// GetAVAssetDownloadTaskClass returns the class object for AVAssetDownloadTask.
func GetAVAssetDownloadTaskClass() AVAssetDownloadTaskClass {
	return getAVAssetDownloadTaskClass()
}

type AVAssetDownloadTaskClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAssetDownloadTaskClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetDownloadTaskClass) Alloc() AVAssetDownloadTask {
	rv := objc.Send[AVAssetDownloadTask](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A session used to download HTTP Live Streaming assets.
//
// # Overview
//
// This class is a subclass of [URLSessionTask] that you use to download HTTP
// Live Streaming assets. You create instances of this class by calling
// [AssetDownloadTaskWithConfiguration] on the download session.
//
// # Accessing task information
//
//   - [AVAssetDownloadTask.URLAsset]: The asset that this task downloads.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadTask
//
// [URLSessionTask]: https://developer.apple.com/documentation/Foundation/URLSessionTask
type AVAssetDownloadTask struct {
	foundation.URLSessionTask
}

// AVAssetDownloadTaskFromID constructs a [AVAssetDownloadTask] from an objc.ID.
//
// A session used to download HTTP Live Streaming assets.
func AVAssetDownloadTaskFromID(id objc.ID) AVAssetDownloadTask {
	return AVAssetDownloadTask{URLSessionTask: foundation.URLSessionTaskFromID(id)}
}

// NOTE: AVAssetDownloadTask adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetDownloadTask] class.
//
// # Accessing task information
//
//   - [IAVAssetDownloadTask.URLAsset]: The asset that this task downloads.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadTask
type IAVAssetDownloadTask interface {
	foundation.IURLSessionTask

	// Topic: Accessing task information

	// The asset that this task downloads.
	URLAsset() IAVURLAsset
}

// Init initializes the instance.
func (a AVAssetDownloadTask) Init() AVAssetDownloadTask {
	rv := objc.Send[AVAssetDownloadTask](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetDownloadTask) Autorelease() AVAssetDownloadTask {
	rv := objc.Send[AVAssetDownloadTask](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetDownloadTask creates a new AVAssetDownloadTask instance.
func NewAVAssetDownloadTask() AVAssetDownloadTask {
	class := getAVAssetDownloadTaskClass()
	rv := objc.Send[AVAssetDownloadTask](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The asset that this task downloads.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadTask/urlAsset
func (a AVAssetDownloadTask) URLAsset() IAVURLAsset {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("URLAsset"))
	return AVURLAssetFromID(objc.ID(rv))
}
