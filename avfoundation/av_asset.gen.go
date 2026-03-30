// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAsset] class.
var (
	_AVAssetClass     AVAssetClass
	_AVAssetClassOnce sync.Once
)

func getAVAssetClass() AVAssetClass {
	_AVAssetClassOnce.Do(func() {
		_AVAssetClass = AVAssetClass{class: objc.GetClass("AVAsset")}
	})
	return _AVAssetClass
}

// GetAVAssetClass returns the class object for AVAsset.
func GetAVAssetClass() AVAssetClass {
	return getAVAssetClass()
}

type AVAssetClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAssetClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetClass) Alloc() AVAsset {
	rv := objc.Send[AVAsset](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that models timed audiovisual media.
//
// # Overview
//
// An asset models file-based media like a QuickTime movie or an MP3 audio
// file, and also media streamed using HTTP Live Streaming (HLS). An asset is
// a container object for one or more instances of [AVAssetTrack] that model
// the uniformly typed tracks of media. The most commonly used track types are
// audio and video, but assets may also contain supplementary tracks, like
// closed captions, subtitles, and timed metadata.
//
// [media-3845943]
//
// You load the tracks for an asset by asynchronously loading its [tracks]
// property. In some cases, you may want to perform operations on a subset of
// an asset’s tracks rather than on its complete collection. For those
// situations, an asset provides methods to retrieve subsets of tracks
// according to particular criteria, such as identifier, media type, or
// characteristic.
//
// # Loading tracks
//
//   - [AVAsset.LoadTrackWithTrackIDCompletionHandler]: Loads a track that contains the specified identifier.
//   - [AVAsset.FindUnusedTrackIDWithCompletionHandler]: Loads an identifier that no other track in the asset uses.
//
// # Loading media selections
//
//   - [AVAsset.LoadMediaSelectionGroupForMediaCharacteristicCompletionHandler]: Loads a media selection group that contains one or more options with the specified media characteristic.
//
// # Canceling property loading
//
//   - [AVAsset.CancelLoading]: Cancels all pending requests to asynchronously load property values.
//
// # Retrieving reference restrictions
//
//   - [AVAsset.ReferenceRestrictions]: The restrictions that an asset places on how it resolves references to external media.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAsset
//
// [tracks]: https://developer.apple.com/documentation/AVFoundation/AVPartialAsyncProperty/tracks-48zyw
type AVAsset struct {
	objectivec.Object
}

// AVAssetFromID constructs a [AVAsset] from an objc.ID.
//
// An object that models timed audiovisual media.
func AVAssetFromID(id objc.ID) AVAsset {
	return AVAsset{objectivec.Object{ID: id}}
}

// NOTE: AVAsset adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAsset] class.
//
// # Loading tracks
//
//   - [IAVAsset.LoadTrackWithTrackIDCompletionHandler]: Loads a track that contains the specified identifier.
//   - [IAVAsset.FindUnusedTrackIDWithCompletionHandler]: Loads an identifier that no other track in the asset uses.
//
// # Loading media selections
//
//   - [IAVAsset.LoadMediaSelectionGroupForMediaCharacteristicCompletionHandler]: Loads a media selection group that contains one or more options with the specified media characteristic.
//
// # Canceling property loading
//
//   - [IAVAsset.CancelLoading]: Cancels all pending requests to asynchronously load property values.
//
// # Retrieving reference restrictions
//
//   - [IAVAsset.ReferenceRestrictions]: The restrictions that an asset places on how it resolves references to external media.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAsset
type IAVAsset interface {
	objectivec.IObject
	AVAsynchronousKeyValueLoading

	// Topic: Loading tracks

	// Loads a track that contains the specified identifier.
	LoadTrackWithTrackIDCompletionHandler(trackID int32, completionHandler AVAssetTrackErrorHandler)
	// Loads an identifier that no other track in the asset uses.
	FindUnusedTrackIDWithCompletionHandler(completionHandler CMPersistentTrackIDErrorHandler)

	// Topic: Loading media selections

	// Loads a media selection group that contains one or more options with the specified media characteristic.
	LoadMediaSelectionGroupForMediaCharacteristicCompletionHandler(mediaCharacteristic AVMediaCharacteristic, completionHandler AVMediaSelectionGroupErrorHandler)

	// Topic: Canceling property loading

	// Cancels all pending requests to asynchronously load property values.
	CancelLoading()

	// Topic: Retrieving reference restrictions

	// The restrictions that an asset places on how it resolves references to external media.
	ReferenceRestrictions() AVAssetReferenceRestrictions
}

// Init initializes the instance.
func (a AVAsset) Init() AVAsset {
	rv := objc.Send[AVAsset](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAsset) Autorelease() AVAsset {
	rv := objc.Send[AVAsset](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAsset creates a new AVAsset instance.
func NewAVAsset() AVAsset {
	class := getAVAssetClass()
	rv := objc.Send[AVAsset](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an asset that models the media at the specified URL.
//
// URL: A URL to a local, remote, or HTTP Live Streaming media resource.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAsset/init(url:)
func NewAssetWithURL(URL foundation.INSURL) AVAsset {
	rv := objc.Send[objc.ID](objc.ID(getAVAssetClass().class), objc.Sel("assetWithURL:"), URL)
	return AVAssetFromID(rv)
}

// Loads a track that contains the specified identifier.
//
// trackID: The identifier of the track to load.
//
// completionHandler: A callback that the system invokes after it finishes the loading request.
// It passes the completion handler the following parameters:
//
// track: The loaded track, or `nil` if no track with the specified identifier
// exists or if an error occurs. error: An error object if the request fails;
// otherwise, `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAsset/loadTrack(withTrackID:completionHandler:)
func (a AVAsset) LoadTrackWithTrackIDCompletionHandler(trackID int32, completionHandler AVAssetTrackErrorHandler) {
	_block1, _ := NewAVAssetTrackErrorBlock(completionHandler)
	objc.Send[objc.ID](a.ID, objc.Sel("loadTrackWithTrackID:completionHandler:"), trackID, _block1)
}

// Loads an identifier that no other track in the asset uses.
//
// completionHandler: A completion handler the system calls after it finishes the request.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAsset/findUnusedTrackID(completionHandler:)
func (a AVAsset) FindUnusedTrackIDWithCompletionHandler(completionHandler CMPersistentTrackIDErrorHandler) {
	_block0, _ := NewCMPersistentTrackIDErrorBlock(completionHandler)
	objc.Send[objc.ID](a.ID, objc.Sel("findUnusedTrackIDWithCompletionHandler:"), _block0)
}

// Loads a media selection group that contains one or more options with the
// specified media characteristic.
//
// mediaCharacteristic: A media characteristic to load the available media selection options for.
// The supported characterisics are:
//
// - [audible] to return the group of available options for audio media in
// various languages and for various purposes, such as descriptive audio -
// [legible] to return the group of available options for subtitles in various
// languages and for various purposes - [visual] to return the group of
// available options for video media
//
// completionHandler: A callback that the system invokes after it finishes the loading request.
// It passes the completion handler the following parameters:
//
// mediaSelectionGroup: The loaded media selection group, or `nil` if no group
// is available or if an error occurs. error: An error object if the request
// fails; otherwise, `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAsset/loadMediaSelectionGroup(for:completionHandler:)
//
// [audible]: https://developer.apple.com/documentation/AVFoundation/AVMediaCharacteristic/audible
// [legible]: https://developer.apple.com/documentation/AVFoundation/AVMediaCharacteristic/legible
// [visual]: https://developer.apple.com/documentation/AVFoundation/AVMediaCharacteristic/visual
func (a AVAsset) LoadMediaSelectionGroupForMediaCharacteristicCompletionHandler(mediaCharacteristic AVMediaCharacteristic, completionHandler AVMediaSelectionGroupErrorHandler) {
	_block1, _ := NewAVMediaSelectionGroupErrorBlock(completionHandler)
	objc.Send[objc.ID](a.ID, objc.Sel("loadMediaSelectionGroupForMediaCharacteristic:completionHandler:"), mediaCharacteristic, _block1)
}

// Cancels all pending requests to asynchronously load property values.
//
// # Discussion
//
// Calling this method cancels pending requests to load an asset’s property
// values. Call this method only when you’re done using an asset and you
// want to cancel any outstanding requests. Deallocating an asset implicitly
// calls this method if loading requests are still pending.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAsset/cancelLoading()
func (a AVAsset) CancelLoading() {
	objc.Send[objc.ID](a.ID, objc.Sel("cancelLoading"))
}

// The restrictions that an asset places on how it resolves references to
// external media.
//
// # Discussion
//
// For [AVURLAsset], this property reflects the value passed in for
// [AVURLAssetReferenceRestrictionsKey], if any.
//
// The default value for this property is
// [AVAssetReferenceRestrictionDefaultPolicy]. See
// [AVURLAssetReferenceRestrictionsKey] for more information about reference
// restrictions.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAsset/referenceRestrictions
//
// [AVURLAssetReferenceRestrictionsKey]: https://developer.apple.com/documentation/AVFoundation/AVURLAssetReferenceRestrictionsKey
func (a AVAsset) ReferenceRestrictions() AVAssetReferenceRestrictions {
	rv := objc.Send[AVAssetReferenceRestrictions](a.ID, objc.Sel("referenceRestrictions"))
	return AVAssetReferenceRestrictions(rv)
}

// Protocol methods for AVAsynchronousKeyValueLoading

// LoadTrackWithTrackID is a synchronous wrapper around [AVAsset.LoadTrackWithTrackIDCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (a AVAsset) LoadTrackWithTrackID(ctx context.Context, trackID int32) (*AVAssetTrack, error) {
	type result struct {
		val *AVAssetTrack
		err error
	}
	done := make(chan result, 1)
	a.LoadTrackWithTrackIDCompletionHandler(trackID, func(val *AVAssetTrack, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// FindUnusedTrackID is a synchronous wrapper around [AVAsset.FindUnusedTrackIDWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (a AVAsset) FindUnusedTrackID(ctx context.Context) (int32, error) {
	type result struct {
		val int32
		err error
	}
	done := make(chan result, 1)
	a.FindUnusedTrackIDWithCompletionHandler(func(val int32, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

// LoadMediaSelectionGroupForMediaCharacteristic is a synchronous wrapper around [AVAsset.LoadMediaSelectionGroupForMediaCharacteristicCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (a AVAsset) LoadMediaSelectionGroupForMediaCharacteristic(ctx context.Context, mediaCharacteristic AVMediaCharacteristic) (*AVMediaSelectionGroup, error) {
	type result struct {
		val *AVMediaSelectionGroup
		err error
	}
	done := make(chan result, 1)
	a.LoadMediaSelectionGroupForMediaCharacteristicCompletionHandler(mediaCharacteristic, func(val *AVMediaSelectionGroup, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
