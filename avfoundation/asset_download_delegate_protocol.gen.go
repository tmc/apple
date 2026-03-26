// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// A protocol that defines the methods to implement to respond to asset-download events.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadDelegate
type AVAssetDownloadDelegate interface {
	objectivec.IObject
}

// AVAssetDownloadDelegateObject wraps an existing Objective-C object that conforms to the AVAssetDownloadDelegate protocol.
type AVAssetDownloadDelegateObject struct {
	objectivec.Object
}
func (o AVAssetDownloadDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVAssetDownloadDelegateObjectFromID constructs a [AVAssetDownloadDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVAssetDownloadDelegateObjectFromID(id objc.ID) AVAssetDownloadDelegateObject {
	return AVAssetDownloadDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate that a download task resolved the media selection to
// download, including any automatic selections.
//
// session: The session the asset download task is on.
//
// assetDownloadTask: The task that resolved the media selection.
//
// resolvedMediaSelection: The media selection the task resolved.
//
// # Discussion
// 
// For the best chance of playing back downloaded content without further
// network I/O, set this selection on the associated [AVPlayerItem].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadDelegate/urlSession(_:assetDownloadTask:didResolve:)
func (o AVAssetDownloadDelegateObject) URLSessionAssetDownloadTaskDidResolveMediaSelection(session foundation.NSURLSession, assetDownloadTask IAVAssetDownloadTask, resolvedMediaSelection IAVMediaSelection) {
	objc.Send[struct{}](o.ID, objc.Sel("URLSession:assetDownloadTask:didResolveMediaSelection:"), session, assetDownloadTask, resolvedMediaSelection)
	}
// Tells the delegate that a download task completed variant selection.
//
// session: The session the asset download task is on.
//
// assetDownloadTask: The task that finished selecting variant selection.
//
// variants: The asset variants to download.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadDelegate/urlSession(_:assetDownloadTask:willDownloadVariants:)
func (o AVAssetDownloadDelegateObject) URLSessionAssetDownloadTaskWillDownloadVariants(session foundation.NSURLSession, assetDownloadTask IAVAssetDownloadTask, variants []AVAssetVariant) {
	objc.Send[struct{}](o.ID, objc.Sel("URLSession:assetDownloadTask:willDownloadVariants:"), session, assetDownloadTask, objectivec.IObjectSliceToNSArray(variants))
	}
// Tells the delegate when a download task determines its download location.
//
// session: The session the asset download task is on.
//
// assetDownloadTask: The download task.
//
// location: The URL the task downloads the asset to.
//
// # Discussion
// 
// Save the returned URL to instantiate the asset in the future.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadDelegate/urlSession(_:assetDownloadTask:willDownloadTo:)
func (o AVAssetDownloadDelegateObject) URLSessionAssetDownloadTaskWillDownloadToURL(session foundation.NSURLSession, assetDownloadTask IAVAssetDownloadTask, location foundation.INSURL) {
	objc.Send[struct{}](o.ID, objc.Sel("URLSession:assetDownloadTask:willDownloadToURL:"), session, assetDownloadTask, location)
	}
// Sent when a download task receives an AVMetricEvent.
//
// session: The NSURLSession corresponding to this AVAssetDownloadTask.
//
// assetDownloadTask: The asset download task.
//
// metricEvent: The metric event received.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadDelegate/urlSession(_:assetDownloadTask:didReceive:)
func (o AVAssetDownloadDelegateObject) URLSessionAssetDownloadTaskDidReceiveMetricEvent(session foundation.NSURLSession, assetDownloadTask IAVAssetDownloadTask, metricEvent IAVMetricEvent) {
	objc.Send[struct{}](o.ID, objc.Sel("URLSession:assetDownloadTask:didReceiveMetricEvent:"), session, assetDownloadTask, metricEvent)
	}

// AVAssetDownloadDelegateConfig holds optional typed callbacks for [AVAssetDownloadDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfoundation/avassetdownloaddelegate
type AVAssetDownloadDelegateConfig struct {

	// Other Methods
	// URLSessionAssetDownloadTaskDidResolveMediaSelection — Tells the delegate that a download task resolved the media selection to download, including any automatic selections.
	URLSessionAssetDownloadTaskDidResolveMediaSelection func(session foundation.NSURLSession, assetDownloadTask AVAssetDownloadTask, resolvedMediaSelection AVMediaSelection)
	// URLSessionAssetDownloadTaskWillDownloadToURL — Tells the delegate when a download task determines its download location.
	URLSessionAssetDownloadTaskWillDownloadToURL func(session foundation.NSURLSession, assetDownloadTask AVAssetDownloadTask, location foundation.NSURL)
	// URLSessionAssetDownloadTaskDidReceiveMetricEvent — Sent when a download task receives an AVMetricEvent.
	URLSessionAssetDownloadTaskDidReceiveMetricEvent func(session foundation.NSURLSession, assetDownloadTask AVAssetDownloadTask, metricEvent AVMetricEvent)
}

// NewAVAssetDownloadDelegate creates an Objective-C object implementing the [AVAssetDownloadDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [AVAssetDownloadDelegateObject] satisfies the [AVAssetDownloadDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfoundation/avassetdownloaddelegate
func NewAVAssetDownloadDelegate(config AVAssetDownloadDelegateConfig) AVAssetDownloadDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoAVAssetDownloadDelegate_%d", n)

	var methods []objc.MethodDef

	if config.URLSessionAssetDownloadTaskDidResolveMediaSelection != nil {
		fn := config.URLSessionAssetDownloadTaskDidResolveMediaSelection
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("URLSession:assetDownloadTask:didResolveMediaSelection:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sessionID objc.ID, assetDownloadTaskID objc.ID, resolvedMediaSelectionID objc.ID) {
				session := foundation.NSURLSessionFromID(sessionID)
				assetDownloadTask := AVAssetDownloadTaskFromID(assetDownloadTaskID)
				resolvedMediaSelection := AVMediaSelectionFromID(resolvedMediaSelectionID)
				fn(session, assetDownloadTask, resolvedMediaSelection)
			},
		})
	}

	if config.URLSessionAssetDownloadTaskWillDownloadToURL != nil {
		fn := config.URLSessionAssetDownloadTaskWillDownloadToURL
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("URLSession:assetDownloadTask:willDownloadToURL:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sessionID objc.ID, assetDownloadTaskID objc.ID, locationID objc.ID) {
				session := foundation.NSURLSessionFromID(sessionID)
				assetDownloadTask := AVAssetDownloadTaskFromID(assetDownloadTaskID)
				location := foundation.NSURLFromID(locationID)
				fn(session, assetDownloadTask, location)
			},
		})
	}

	if config.URLSessionAssetDownloadTaskDidReceiveMetricEvent != nil {
		fn := config.URLSessionAssetDownloadTaskDidReceiveMetricEvent
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("URLSession:assetDownloadTask:didReceiveMetricEvent:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sessionID objc.ID, assetDownloadTaskID objc.ID, metricEventID objc.ID) {
				session := foundation.NSURLSessionFromID(sessionID)
				assetDownloadTask := AVAssetDownloadTaskFromID(assetDownloadTaskID)
				metricEvent := AVMetricEventFromID(metricEventID)
				fn(session, assetDownloadTask, metricEvent)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("AVAssetDownloadDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewAVAssetDownloadDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return AVAssetDownloadDelegateObjectFromID(instance)
}

