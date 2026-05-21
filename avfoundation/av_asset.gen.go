// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"context"
	"sync"

	"github.com/tmc/apple/coremedia"
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
//   - [AVAsset.LoadTracksWithMediaTypeCompletionHandler]: Loads tracks that contain media of a specified type.
//   - [AVAsset.LoadTracksWithMediaCharacteristicCompletionHandler]: Loads tracks that contain media of a specified characteristic.
//   - [AVAsset.FindUnusedTrackIDWithCompletionHandler]: Loads an identifier that no other track in the asset uses.
//
// # Loading metadata
//
//   - [AVAsset.LoadMetadataForFormatCompletionHandler]: Loads an array of metadata items that the asset contains for the specified format.
//
// # Loading media selections
//
//   - [AVAsset.LoadMediaSelectionGroupForMediaCharacteristicCompletionHandler]: Loads a media selection group that contains one or more options with the specified media characteristic.
//
// # Loading chapter metadata
//
//   - [AVAsset.LoadChapterMetadataGroupsBestMatchingPreferredLanguagesCompletionHandler]: Loads chapter metadata with a locale that best matches the list of preferred languages.
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
//   - [IAVAsset.LoadTracksWithMediaTypeCompletionHandler]: Loads tracks that contain media of a specified type.
//   - [IAVAsset.LoadTracksWithMediaCharacteristicCompletionHandler]: Loads tracks that contain media of a specified characteristic.
//   - [IAVAsset.FindUnusedTrackIDWithCompletionHandler]: Loads an identifier that no other track in the asset uses.
//
// # Loading metadata
//
//   - [IAVAsset.LoadMetadataForFormatCompletionHandler]: Loads an array of metadata items that the asset contains for the specified format.
//
// # Loading media selections
//
//   - [IAVAsset.LoadMediaSelectionGroupForMediaCharacteristicCompletionHandler]: Loads a media selection group that contains one or more options with the specified media characteristic.
//
// # Loading chapter metadata
//
//   - [IAVAsset.LoadChapterMetadataGroupsBestMatchingPreferredLanguagesCompletionHandler]: Loads chapter metadata with a locale that best matches the list of preferred languages.
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
	LoadTrackWithTrackIDCompletionHandler(trackID coremedia.CMPersistentTrackID, completionHandler AVAssetTrackErrorHandler)
	// Loads tracks that contain media of a specified type.
	LoadTracksWithMediaTypeCompletionHandler(mediaType AVMediaType, completionHandler AVAssetTrackArrayErrorHandler)
	// Loads tracks that contain media of a specified characteristic.
	LoadTracksWithMediaCharacteristicCompletionHandler(mediaCharacteristic AVMediaCharacteristic, completionHandler AVAssetTrackArrayErrorHandler)
	// Loads an identifier that no other track in the asset uses.
	FindUnusedTrackIDWithCompletionHandler(completionHandler CMPersistentTrackIDErrorHandler)

	// Topic: Loading metadata

	// Loads an array of metadata items that the asset contains for the specified format.
	LoadMetadataForFormatCompletionHandler(format AVMetadataFormat, completionHandler AVMetadataItemArrayErrorHandler)

	// Topic: Loading media selections

	// Loads a media selection group that contains one or more options with the specified media characteristic.
	LoadMediaSelectionGroupForMediaCharacteristicCompletionHandler(mediaCharacteristic AVMediaCharacteristic, completionHandler AVMediaSelectionGroupErrorHandler)

	// Topic: Loading chapter metadata

	// Loads chapter metadata with a locale that best matches the list of preferred languages.
	LoadChapterMetadataGroupsBestMatchingPreferredLanguagesCompletionHandler(preferredLanguages []string, completionHandler AVTimedMetadataGroupArrayErrorHandler)

	// Topic: Canceling property loading

	// Cancels all pending requests to asynchronously load property values.
	CancelLoading()

	// Topic: Retrieving reference restrictions

	// The restrictions that an asset places on how it resolves references to external media.
	ReferenceRestrictions() AVAssetReferenceRestrictions

	// Loads chapter metadata that contains the specified title locale and common keys.
	LoadChapterMetadataGroupsWithTitleLocaleContainingItemsWithCommonKeysCompletionHandler(locale foundation.NSLocale, commonKeys []string, completionHandler AVTimedMetadataGroupArrayErrorHandler)
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
// See: https://developer.apple.com/documentation/AVFoundation/AVAsset/init(url:)-42gl8
func NewAssetWithURL(URL foundation.NSURL) AVAsset {
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
func (a AVAsset) LoadTrackWithTrackIDCompletionHandler(trackID coremedia.CMPersistentTrackID, completionHandler AVAssetTrackErrorHandler) {
	_block1, _ := NewAVAssetTrackErrorBlock(completionHandler)
	objc.Send[objc.ID](a.ID, objc.Sel("loadTrackWithTrackID:completionHandler:"), trackID, _block1)
}

// Loads tracks that contain media of a specified type.
//
// mediaType: The media type of the tracks to load.
//
// completionHandler: A callback that the system invokes after it finishes the loading operation.
// It passes the completion handler the following parameters:
//
// tracks: An array of tracks, which may be empty if no tracks with the
// specified media type exist. The value is `nil` if an error occurs. error:
// An error object if the request fails; otherwise, `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAsset/loadTracks(withMediaType:completionHandler:)
func (a AVAsset) LoadTracksWithMediaTypeCompletionHandler(mediaType AVMediaType, completionHandler AVAssetTrackArrayErrorHandler) {
	_block1, _ := NewAVAssetTrackArrayErrorBlock(completionHandler)
	objc.Send[objc.ID](a.ID, objc.Sel("loadTracksWithMediaType:completionHandler:"), mediaType, _block1)
}

// Loads tracks that contain media of a specified characteristic.
//
// mediaCharacteristic: The media characteristic of the tracks to load.
//
// completionHandler: A callback that the system invokes after it finishes the loading request.
// It passes the completion handler the following parameters:
//
// tracks: An array of tracks, which may be empty if no tracks with the
// specified media characteristic exist. The value is `nil` if an error
// occurs. error: An error object if the request fails; otherwise, `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAsset/loadTracks(withMediaCharacteristic:completionHandler:)
func (a AVAsset) LoadTracksWithMediaCharacteristicCompletionHandler(mediaCharacteristic AVMediaCharacteristic, completionHandler AVAssetTrackArrayErrorHandler) {
	_block1, _ := NewAVAssetTrackArrayErrorBlock(completionHandler)
	objc.Send[objc.ID](a.ID, objc.Sel("loadTracksWithMediaCharacteristic:completionHandler:"), mediaCharacteristic, _block1)
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

// Loads an array of metadata items that the asset contains for the specified
// format.
//
// format: The format of the metadata items to load.
//
// completionHandler: A callback that the system invokes after it finishes the loading request.
// It passes the completion handler the following parameters:
//
// metadata: An array of metadata items, which may be empty if there are no
// items of the specified format. The value is `nil` if an error occurs.
// error: An error object if the request fails; otherwise, `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAsset/loadMetadata(for:completionHandler:)
func (a AVAsset) LoadMetadataForFormatCompletionHandler(format AVMetadataFormat, completionHandler AVMetadataItemArrayErrorHandler) {
	_block1, _ := NewAVMetadataItemArrayErrorBlock(completionHandler)
	objc.Send[objc.ID](a.ID, objc.Sel("loadMetadataForFormat:completionHandler:"), format, _block1)
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

// Loads chapter metadata with a locale that best matches the list of
// preferred languages.
//
// preferredLanguages: An array of language identifiers in order of preference, each of which is
// an IETF BCP 47 (RFC 4646) language identifier. Call [preferredLanguages] to
// retrieve the list of languates the user prefers.
//
// completionHandler: A callback that the system invokes after it finishes the loading request.
// It passes the completion handler the following parameters:
//
// metadataGroups: An array of metadata groups, which may be empty if no
// groups exist for the specified languages. The value is `nil` if an error
// occurs. error: An error object if the request fails; otherwise, `nil`.
//
// # Discussion
//
// This method returns an array of [AVTimedMetadataGroup] objects
// asynchronously. Each object in the array contains an [AVMetadataItem] that
// represents the chapter’s title, and the metadata group’s
// [AVTimedMetadataGroup.TimeRange] value equals the time range of the chapter
// title item.
//
// The metadata group contains all chapter metadata, including items with the
// common key [commonKeyArtwork], if such items are present. The system adds
// an [AVMetadataItem] with the specified common key to an existing
// [AVTimedMetadataGroup] object if the time range (timestamp and duration) of
// the metadata item and the metadata group overlap. The locales of such items
// don’t need to match the locale of the chapter titles.
//
// You can use the
// [AVMetadataItemClass.MetadataItemsFromArrayFilteredAndSortedAccordingToPreferredLanguages]
// method to further filter the metadata items in each group. You can also
// filter the returned items based on locale using the
// [AVMetadataItemClass.MetadataItemsFromArrayWithLocale] method.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAsset/loadChapterMetadataGroups(bestMatchingPreferredLanguages:completionHandler:)
//
// [preferredLanguages]: https://developer.apple.com/documentation/Foundation/Locale/preferredLanguages
// [commonKeyArtwork]: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/commonKeyArtwork
func (a AVAsset) LoadChapterMetadataGroupsBestMatchingPreferredLanguagesCompletionHandler(preferredLanguages []string, completionHandler AVTimedMetadataGroupArrayErrorHandler) {
	_block1, _ := NewAVTimedMetadataGroupArrayErrorBlock(completionHandler)
	objc.Send[objc.ID](a.ID, objc.Sel("loadChapterMetadataGroupsBestMatchingPreferredLanguages:completionHandler:"), preferredLanguages, _block1)
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

// Loads chapter metadata that contains the specified title locale and common
// keys.
//
// locale: The locale of the chapter metadata to load.
//
// commonKeys: An array of common keys of [AVMetadataItem] to include in the returned
// array. The framework currently only supports the [commonKeyArtwork] key.
//
// completionHandler: A callback that the system invokes after it finishes the loading request.
// It passes the completion handler the following parameters:
//
// metadataGroups: An array of metadata groups, which may be empty if no
// groups exist for the locale. The value is `nil` if an error occurs. error:
// An error object if the request fails; otherwise, `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAsset/loadChapterMetadataGroupsWithTitleLocale:containingItemsWithCommonKeys:completionHandler:
//
// [commonKeyArtwork]: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/commonKeyArtwork
func (a AVAsset) LoadChapterMetadataGroupsWithTitleLocaleContainingItemsWithCommonKeysCompletionHandler(locale foundation.NSLocale, commonKeys []string, completionHandler AVTimedMetadataGroupArrayErrorHandler) {
	_block2, _ := NewAVTimedMetadataGroupArrayErrorBlock(completionHandler)
	objc.Send[objc.ID](a.ID, objc.Sel("loadChapterMetadataGroupsWithTitleLocale:containingItemsWithCommonKeys:completionHandler:"), locale, commonKeys, _block2)
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
func (a AVAsset) LoadTrackWithTrackID(ctx context.Context, trackID coremedia.CMPersistentTrackID) (*AVAssetTrack, error) {
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

// LoadTracksWithMediaType is a synchronous wrapper around [AVAsset.LoadTracksWithMediaTypeCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (a AVAsset) LoadTracksWithMediaType(ctx context.Context, mediaType AVMediaType) ([]AVAssetTrack, error) {
	type result struct {
		val []AVAssetTrack
		err error
	}
	done := make(chan result, 1)
	a.LoadTracksWithMediaTypeCompletionHandler(mediaType, func(val *[]AVAssetTrack, err error) {
		var out []AVAssetTrack
		if val != nil {
			out = append(out, (*val)...)
		}
		done <- result{out, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// LoadTracksWithMediaCharacteristic is a synchronous wrapper around [AVAsset.LoadTracksWithMediaCharacteristicCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (a AVAsset) LoadTracksWithMediaCharacteristic(ctx context.Context, mediaCharacteristic AVMediaCharacteristic) ([]AVAssetTrack, error) {
	type result struct {
		val []AVAssetTrack
		err error
	}
	done := make(chan result, 1)
	a.LoadTracksWithMediaCharacteristicCompletionHandler(mediaCharacteristic, func(val *[]AVAssetTrack, err error) {
		var out []AVAssetTrack
		if val != nil {
			out = append(out, (*val)...)
		}
		done <- result{out, err}
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
func (a AVAsset) FindUnusedTrackID(ctx context.Context) (coremedia.CMPersistentTrackID, error) {
	type result struct {
		val coremedia.CMPersistentTrackID
		err error
	}
	done := make(chan result, 1)
	a.FindUnusedTrackIDWithCompletionHandler(func(val coremedia.CMPersistentTrackID, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return *new(coremedia.CMPersistentTrackID), ctx.Err()
	}
}

// LoadMetadataForFormat is a synchronous wrapper around [AVAsset.LoadMetadataForFormatCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (a AVAsset) LoadMetadataForFormat(ctx context.Context, format AVMetadataFormat) ([]AVMetadataItem, error) {
	type result struct {
		val []AVMetadataItem
		err error
	}
	done := make(chan result, 1)
	a.LoadMetadataForFormatCompletionHandler(format, func(val *[]AVMetadataItem, err error) {
		var out []AVMetadataItem
		if val != nil {
			out = append(out, (*val)...)
		}
		done <- result{out, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
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
