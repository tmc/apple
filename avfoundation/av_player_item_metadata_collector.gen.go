// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVPlayerItemMetadataCollector] class.
var (
	_AVPlayerItemMetadataCollectorClass     AVPlayerItemMetadataCollectorClass
	_AVPlayerItemMetadataCollectorClassOnce sync.Once
)

func getAVPlayerItemMetadataCollectorClass() AVPlayerItemMetadataCollectorClass {
	_AVPlayerItemMetadataCollectorClassOnce.Do(func() {
		_AVPlayerItemMetadataCollectorClass = AVPlayerItemMetadataCollectorClass{class: objc.GetClass("AVPlayerItemMetadataCollector")}
	})
	return _AVPlayerItemMetadataCollectorClass
}

// GetAVPlayerItemMetadataCollectorClass returns the class object for AVPlayerItemMetadataCollector.
func GetAVPlayerItemMetadataCollectorClass() AVPlayerItemMetadataCollectorClass {
	return getAVPlayerItemMetadataCollectorClass()
}

type AVPlayerItemMetadataCollectorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVPlayerItemMetadataCollectorClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVPlayerItemMetadataCollectorClass) Alloc() AVPlayerItemMetadataCollector {
	rv := objc.Send[AVPlayerItemMetadataCollector](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object used to capture the date range metadata defined for an HTTP Live
// Streaming asset.
//
// # Overview
//
// You can use the HLS `#EXT-X-DATERANGE` tag to define date range metadata in
// a media playlist. This tag is useful for defining timed metadata for
// interstitial regions such as advertisements, but can be used to define any
// timed metadata needed by your stream. To access this metadata when the
// stream is played using an [AVPlayer], you create an instance of
// [AVPlayerItemMetadataCollector], configure its delegate object (see
// [AVPlayerItemMetadataCollectorPushDelegate]), and add it as a media data
// collector to the [AVPlayerItem] (see example).
//
// Creating an [AVPlayerItemMetadataCollector] as shown in the example, will
// capture all `#EXT-X-DATERANGE` metadata defined in your stream. If you
// would like to filter the output to only the metadata of interest, you can
// create an instance to filter by identifier and/or classifying labels using
// the [AVPlayerItemMetadataCollector.InitWithIdentifiersClassifyingLabels] initializer.
//
// # Creating a metadata collector
//
//   - [AVPlayerItemMetadataCollector.InitWithIdentifiersClassifyingLabels]: Creates a metadata collector to access a stream’s metadata groups matching the specified array of identifiers and classifying labels.
//
// # Accessing the delegate and callback queue
//
//   - [AVPlayerItemMetadataCollector.SetDelegateQueue]: Sets the delegate and a dispatch queue on which the delegate will be called.
//   - [AVPlayerItemMetadataCollector.Delegate]: Accesses the metadata collector’s delegate object.
//   - [AVPlayerItemMetadataCollector.DelegateQueue]: The dispatch queue on which the delegate’s methods are called.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemMetadataCollector
type AVPlayerItemMetadataCollector struct {
	AVPlayerItemMediaDataCollector
}

// AVPlayerItemMetadataCollectorFromID constructs a [AVPlayerItemMetadataCollector] from an objc.ID.
//
// An object used to capture the date range metadata defined for an HTTP Live
// Streaming asset.
func AVPlayerItemMetadataCollectorFromID(id objc.ID) AVPlayerItemMetadataCollector {
	return AVPlayerItemMetadataCollector{AVPlayerItemMediaDataCollector: AVPlayerItemMediaDataCollectorFromID(id)}
}

// NOTE: AVPlayerItemMetadataCollector adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVPlayerItemMetadataCollector] class.
//
// # Creating a metadata collector
//
//   - [IAVPlayerItemMetadataCollector.InitWithIdentifiersClassifyingLabels]: Creates a metadata collector to access a stream’s metadata groups matching the specified array of identifiers and classifying labels.
//
// # Accessing the delegate and callback queue
//
//   - [IAVPlayerItemMetadataCollector.SetDelegateQueue]: Sets the delegate and a dispatch queue on which the delegate will be called.
//   - [IAVPlayerItemMetadataCollector.Delegate]: Accesses the metadata collector’s delegate object.
//   - [IAVPlayerItemMetadataCollector.DelegateQueue]: The dispatch queue on which the delegate’s methods are called.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemMetadataCollector
type IAVPlayerItemMetadataCollector interface {
	IAVPlayerItemMediaDataCollector

	// Topic: Creating a metadata collector

	// Creates a metadata collector to access a stream’s metadata groups matching the specified array of identifiers and classifying labels.
	InitWithIdentifiersClassifyingLabels(identifiers []string, classifyingLabels []string) AVPlayerItemMetadataCollector

	// Topic: Accessing the delegate and callback queue

	// Sets the delegate and a dispatch queue on which the delegate will be called.
	SetDelegateQueue(delegate AVPlayerItemMetadataCollectorPushDelegate, delegateQueue dispatch.Queue)
	// Accesses the metadata collector’s delegate object.
	Delegate() AVPlayerItemMetadataCollectorPushDelegate
	// The dispatch queue on which the delegate’s methods are called.
	DelegateQueue() dispatch.Queue
}

// Init initializes the instance.
func (p AVPlayerItemMetadataCollector) Init() AVPlayerItemMetadataCollector {
	rv := objc.Send[AVPlayerItemMetadataCollector](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AVPlayerItemMetadataCollector) Autorelease() AVPlayerItemMetadataCollector {
	rv := objc.Send[AVPlayerItemMetadataCollector](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayerItemMetadataCollector creates a new AVPlayerItemMetadataCollector instance.
func NewAVPlayerItemMetadataCollector() AVPlayerItemMetadataCollector {
	class := getAVPlayerItemMetadataCollectorClass()
	rv := objc.Send[AVPlayerItemMetadataCollector](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a metadata collector to access a stream’s metadata groups
// matching the specified array of identifiers and classifying labels.
//
// identifiers: An optional array of metadata identifiers indicating the metadata items
// that the output should provide.
//
// classifyingLabels: A optional array of classifying labels indicating the metadata items that
// the output should provide.
//
// # Return Value
//
// An instance of [AVPlayerItemMetadataCollector].
//
// # Discussion
//
// You can use the `identifiers` and `classifyingLabels` arguments to
// configure the metadata collector to filter its output to only the metadata
// items matching the specified criteria.
//
// You use the `identifiers` argument to filter the output to a particular set
// of metadata identifiers. For instance, if the stream’s `#EXT-X-DATERANGE`
// tags define multiple metadata attributes, but you are only interested in
// the values for the `X-AD-ID` and `X-AD-URL` attributes, you could configure
// the collector as follows:
//
// The `classifyingLabels` argument is used to filter by the
// `#EXT-X-DATERANGE` tag’s [CLASS] attribute. The [CLASS] attribute can be
// used to define a set of attribute/value pairs for a particular purpose
// (such as describing an ad) and then mark each `#EXT-X-DATERANGE` instance
// as having those semantics. For instance, this might be used by a third
// party advertising SDK to filter the output to only the metadata relevant to
// its needs. It could define an “Advertiser-ad” [CLASS] with the
// following attributes:
//
// - `X-ADVERTISER-AD-GUID` (the unique identifier for the ad) -
// `X-ADVERTISER-AD-AGE` (the number of days in its inventory)
//
// The SDK could require clients to pass it any player items the app creates
// so it could configure them to output its needed data as shown below:
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemMetadataCollector/init(identifiers:classifyingLabels:)
func NewPlayerItemMetadataCollectorWithIdentifiersClassifyingLabels(identifiers []string, classifyingLabels []string) AVPlayerItemMetadataCollector {
	instance := getAVPlayerItemMetadataCollectorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIdentifiers:classifyingLabels:"), objectivec.StringSliceToNSArray(identifiers), objectivec.StringSliceToNSArray(classifyingLabels))
	return AVPlayerItemMetadataCollectorFromID(rv)
}

// Creates a metadata collector to access a stream’s metadata groups
// matching the specified array of identifiers and classifying labels.
//
// identifiers: An optional array of metadata identifiers indicating the metadata items
// that the output should provide.
//
// classifyingLabels: A optional array of classifying labels indicating the metadata items that
// the output should provide.
//
// # Return Value
//
// An instance of [AVPlayerItemMetadataCollector].
//
// # Discussion
//
// You can use the `identifiers` and `classifyingLabels` arguments to
// configure the metadata collector to filter its output to only the metadata
// items matching the specified criteria.
//
// You use the `identifiers` argument to filter the output to a particular set
// of metadata identifiers. For instance, if the stream’s `#EXT-X-DATERANGE`
// tags define multiple metadata attributes, but you are only interested in
// the values for the `X-AD-ID` and `X-AD-URL` attributes, you could configure
// the collector as follows:
//
// The `classifyingLabels` argument is used to filter by the
// `#EXT-X-DATERANGE` tag’s [CLASS] attribute. The [CLASS] attribute can be
// used to define a set of attribute/value pairs for a particular purpose
// (such as describing an ad) and then mark each `#EXT-X-DATERANGE` instance
// as having those semantics. For instance, this might be used by a third
// party advertising SDK to filter the output to only the metadata relevant to
// its needs. It could define an “Advertiser-ad” [CLASS] with the
// following attributes:
//
// - `X-ADVERTISER-AD-GUID` (the unique identifier for the ad) -
// `X-ADVERTISER-AD-AGE` (the number of days in its inventory)
//
// The SDK could require clients to pass it any player items the app creates
// so it could configure them to output its needed data as shown below:
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemMetadataCollector/init(identifiers:classifyingLabels:)
func (p AVPlayerItemMetadataCollector) InitWithIdentifiersClassifyingLabels(identifiers []string, classifyingLabels []string) AVPlayerItemMetadataCollector {
	rv := objc.Send[AVPlayerItemMetadataCollector](p.ID, objc.Sel("initWithIdentifiers:classifyingLabels:"), objectivec.StringSliceToNSArray(identifiers), objectivec.StringSliceToNSArray(classifyingLabels))
	return rv
}

// Sets the delegate and a dispatch queue on which the delegate will be
// called.
//
// delegate: An object conforming to [AVPlayerItemMetadataCollectorPushDelegate]
// protocol.
//
// delegateQueue: A dispatch queue on which all delegate methods will be called.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemMetadataCollector/setDelegate(_:queue:)
func (p AVPlayerItemMetadataCollector) SetDelegateQueue(delegate AVPlayerItemMetadataCollectorPushDelegate, delegateQueue dispatch.Queue) {
	objc.Send[objc.ID](p.ID, objc.Sel("setDelegate:queue:"), delegate, uintptr(delegateQueue.Handle()))
}

// Accesses the metadata collector’s delegate object.
//
// # Discussion
//
// The delegate is held using a zeroing-weak reference, so this property will
// have a value of `nil` after a delegate that was previously set has been
// deallocated.
//
// This property is not key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemMetadataCollector/delegate
func (p AVPlayerItemMetadataCollector) Delegate() AVPlayerItemMetadataCollectorPushDelegate {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("delegate"))
	return AVPlayerItemMetadataCollectorPushDelegateObjectFromID(rv)
}

// The dispatch queue on which the delegate’s methods are called.
//
// # Discussion
//
// This property is not key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemMetadataCollector/delegateQueue
func (p AVPlayerItemMetadataCollector) DelegateQueue() dispatch.Queue {
	rv := objc.Send[uintptr](p.ID, objc.Sel("delegateQueue"))
	return dispatch.QueueFromHandle(rv)
}
