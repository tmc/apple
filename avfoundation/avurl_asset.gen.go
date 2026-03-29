// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/uniformtypeidentifiers"
)

// The class instance for the [AVURLAsset] class.
var (
	_AVURLAssetClass     AVURLAssetClass
	_AVURLAssetClassOnce sync.Once
)

func getAVURLAssetClass() AVURLAssetClass {
	_AVURLAssetClassOnce.Do(func() {
		_AVURLAssetClass = AVURLAssetClass{class: objc.GetClass("AVURLAsset")}
	})
	return _AVURLAssetClass
}

// GetAVURLAssetClass returns the class object for AVURLAsset.
func GetAVURLAssetClass() AVURLAssetClass {
	return getAVURLAssetClass()
}

type AVURLAssetClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVURLAssetClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVURLAssetClass) Alloc() AVURLAsset {
	rv := objc.Send[AVURLAsset](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An asset that represents media at a local or remote URL.
//
// # Overview
// 
// This class is a concrete subclass of [AVAsset]. When you create an asset as
// shown below, the system creates and returns an instance of [AVURLAsset].
// 
// In many cases, this is an appropriate way to create asset instances, but
// you can also directly instantiate an [AVURLAsset] when you need more
// fine-grained control over its initialization. The initializer for
// [AVURLAsset] accepts an options dictionary, which you use to customize the
// asset’s initialization for your particular purpose. For example, if
// you’re creating an asset for an HLS stream, you may want to prevent it
// from retrieving its media when it connects over a cellular network. You can
// do this by providing the initialization option and value as shown below.
//
// # Creating an asset
//
//   - [AVURLAsset.InitWithURLOptions]: Creates an asset that models the media resource at the specified URL.
//
// # Loading tracks
//
//   - [AVURLAsset.Tracks]: The tracks an asset contains.
//   - [AVURLAsset.SetTracks]
//   - [AVURLAsset.FindCompatibleTrackForCompositionTrackCompletionHandler]: Loads an asset track from which you can insert any time range into the composition track.
//
// # Assisting with resource loading
//
//   - [AVURLAsset.ResourceLoader]: The resource loader for the asset.
//
// # Working with offline assets
//
//   - [AVURLAsset.AssetCache]: The asset’s associated asset cache, if it exists.
//
// # Accessing the media URL
//
//   - [AVURLAsset.URL]: A URL to the asset’s media.
//
// # Accessing the session identifier
//
//   - [AVURLAsset.HttpSessionIdentifier]: A session identifier that the asset sends in HTTP requests that it makes.
//
// # Accessing Media Extension properties
//
//   - [AVURLAsset.MediaExtensionProperties]: The properties of the media extension format reader that decodes the asset.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVURLAsset
type AVURLAsset struct {
	AVAsset
}

// AVURLAssetFromID constructs a [AVURLAsset] from an objc.ID.
//
// An asset that represents media at a local or remote URL.
func AVURLAssetFromID(id objc.ID) AVURLAsset {
	return AVURLAsset{AVAsset: AVAssetFromID(id)}
}
// NOTE: AVURLAsset adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVURLAsset] class.
//
// # Creating an asset
//
//   - [IAVURLAsset.InitWithURLOptions]: Creates an asset that models the media resource at the specified URL.
//
// # Loading tracks
//
//   - [IAVURLAsset.Tracks]: The tracks an asset contains.
//   - [IAVURLAsset.SetTracks]
//   - [IAVURLAsset.FindCompatibleTrackForCompositionTrackCompletionHandler]: Loads an asset track from which you can insert any time range into the composition track.
//
// # Assisting with resource loading
//
//   - [IAVURLAsset.ResourceLoader]: The resource loader for the asset.
//
// # Working with offline assets
//
//   - [IAVURLAsset.AssetCache]: The asset’s associated asset cache, if it exists.
//
// # Accessing the media URL
//
//   - [IAVURLAsset.URL]: A URL to the asset’s media.
//
// # Accessing the session identifier
//
//   - [IAVURLAsset.HttpSessionIdentifier]: A session identifier that the asset sends in HTTP requests that it makes.
//
// # Accessing Media Extension properties
//
//   - [IAVURLAsset.MediaExtensionProperties]: The properties of the media extension format reader that decodes the asset.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVURLAsset
type IAVURLAsset interface {
	IAVAsset
	AVContentKeyRecipient

	// Topic: Creating an asset

	// Creates an asset that models the media resource at the specified URL.
	InitWithURLOptions(URL foundation.INSURL, options foundation.INSDictionary) AVURLAsset

	// Topic: Loading tracks

	// The tracks an asset contains.
	Tracks() IAVAssetTrack
	SetTracks(value IAVAssetTrack)
	// Loads an asset track from which you can insert any time range into the composition track.
	FindCompatibleTrackForCompositionTrackCompletionHandler(compositionTrack IAVCompositionTrack, completionHandler AVAssetTrackErrorHandler)

	// Topic: Assisting with resource loading

	// The resource loader for the asset.
	ResourceLoader() IAVAssetResourceLoader

	// Topic: Working with offline assets

	// The asset’s associated asset cache, if it exists.
	AssetCache() IAVAssetCache

	// Topic: Accessing the media URL

	// A URL to the asset’s media.
	URL() foundation.INSURL

	// Topic: Accessing the session identifier

	// A session identifier that the asset sends in HTTP requests that it makes.
	HttpSessionIdentifier() foundation.NSUUID

	// Topic: Accessing Media Extension properties

	// The properties of the media extension format reader that decodes the asset.
	MediaExtensionProperties() IAVMediaExtensionProperties

	// The sidecar URL used by the MediaExtension.
	SidecarURL() foundation.INSURL
}

// Init initializes the instance.
func (u AVURLAsset) Init() AVURLAsset {
	rv := objc.Send[AVURLAsset](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u AVURLAsset) Autorelease() AVURLAsset {
	rv := objc.Send[AVURLAsset](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVURLAsset creates a new AVURLAsset instance.
func NewAVURLAsset() AVURLAsset {
	class := getAVURLAssetClass()
	rv := objc.Send[AVURLAsset](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an asset that models the media at the specified URL.
//
// URL: A URL to a local, remote, or HTTP Live Streaming media resource.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAsset/init(url:)
func NewURLAssetWithURL(URL foundation.INSURL) AVURLAsset {
	rv := objc.Send[objc.ID](objc.ID(getAVURLAssetClass().class), objc.Sel("assetWithURL:"), URL)
	return AVURLAssetFromID(rv)
}

// Creates an asset that models the media resource at the specified URL.
//
// URL: A URL that references the media for the asset to model.
//
// options: A dictionary that contains options used to customize the initialization of
// the asset.
// 
// For supported keys and values, see [Initialization options].
// //
// [Initialization options]: https://developer.apple.com/documentation/AVFoundation/initialization-options
//
// # Return Value
// 
// An asset that models the media resource found at [URL].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/init(url:options:)
func NewURLAssetWithURLOptions(URL foundation.INSURL, options foundation.INSDictionary) AVURLAsset {
	instance := getAVURLAssetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:options:"), URL, options)
	return AVURLAssetFromID(rv)
}

// Creates an asset that models the media resource at the specified URL.
//
// URL: A URL that references the media for the asset to model.
//
// options: A dictionary that contains options used to customize the initialization of
// the asset.
// 
// For supported keys and values, see [Initialization options].
// //
// [Initialization options]: https://developer.apple.com/documentation/AVFoundation/initialization-options
//
// # Return Value
// 
// An asset that models the media resource found at [URL].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/init(url:options:)
func (u AVURLAsset) InitWithURLOptions(URL foundation.INSURL, options foundation.INSDictionary) AVURLAsset {
	rv := objc.Send[AVURLAsset](u.ID, objc.Sel("initWithURL:options:"), URL, options)
	return rv
}
// Loads an asset track from which you can insert any time range into the
// composition track.
//
// compositionTrack: A composition track to request an asset track for.
//
// completionHandler: A callback the system invokes after it finishes the request. The system
// calls the completion handler with the following arguments:
// 
// track: The compatible asset track, or `nil` if there isn’t one or an
// error occurs. error: An error object if the request fails; otherwise,
// `nil`.
//
// # Discussion
// 
// This method is the logical complement of [MutableTrackCompatibleWithTrack].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/findCompatibleTrack(for:completionHandler:)
func (u AVURLAsset) FindCompatibleTrackForCompositionTrackCompletionHandler(compositionTrack IAVCompositionTrack, completionHandler AVAssetTrackErrorHandler) {
_block1, _ := NewAVAssetTrackErrorBlock(completionHandler)
	objc.Send[objc.ID](u.ID, objc.Sel("findCompatibleTrackForCompositionTrack:completionHandler:"), compositionTrack, _block1)
}
// Tells the recipient that a content key is available.
//
// contentKeySession: The current content key session.
//
// contentKey: A content key to use with objects that support manual attachment of keys,
// such as [CMSampleBuffer].
// //
// [CMSampleBuffer]: https://developer.apple.com/documentation/CoreMedia/CMSampleBuffer
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRecipient/contentKeySession(_:didProvide:)
func (u AVURLAsset) ContentKeySessionDidProvideContentKey(contentKeySession IAVContentKeySession, contentKey IAVContentKey) {
	objc.Send[objc.ID](u.ID, objc.Sel("contentKeySession:didProvideContentKey:"), contentKeySession, contentKey)
}

// Returns an array of the MIME types the asset supports.
//
// # Return Value
// 
// An array of MIME type strings.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/audiovisualMIMETypes()
func (_AVURLAssetClass AVURLAssetClass) AudiovisualMIMETypes() []string {
	rv := objc.Send[[]objc.ID](objc.ID(_AVURLAssetClass.class), objc.Sel("audiovisualMIMETypes"))
	return objc.ConvertSliceToStrings(rv)
}
// Returns a Boolean value that indicates whether the asset is playable with
// the specified codecs and container type.
//
// extendedMIMEType: An extended MIME type string such as `video/3gpp2; codecs=“mp4v.20.9,
// mp4a.E1”` or `audio/aac; codecs=“mp4a.E1”`.
//
// # Return Value
// 
// [true] if the asset is playable with the specified codec and container
// type; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/isPlayableExtendedMIMEType(_:)
func (_AVURLAssetClass AVURLAssetClass) IsPlayableExtendedMIMEType(extendedMIMEType string) bool {
	rv := objc.Send[bool](objc.ID(_AVURLAssetClass.class), objc.Sel("isPlayableExtendedMIMEType:"), objc.String(extendedMIMEType))
	return rv
}
// Returns an asset that models the media resource found at the specified URL.
//
// URL: A URL that references the media for the asset to model.
//
// options: A dictionary that contains options used to customize the initialization of
// the asset.
// 
// For possible keys and values, see [Initialization options].
// //
// [Initialization options]: https://developer.apple.com/documentation/AVFoundation/initialization-options
//
// # Return Value
// 
// An asset that models the media resource found at [URL].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/URLAssetWithURL:options:
func (_AVURLAssetClass AVURLAssetClass) URLAssetWithURLOptions(URL foundation.INSURL, options foundation.INSDictionary) AVURLAsset {
	rv := objc.Send[objc.ID](objc.ID(_AVURLAssetClass.class), objc.Sel("URLAssetWithURL:options:"), URL, options)
	return AVURLAssetFromID(rv)
}

// The tracks an asset contains.
//
// See: https://developer.apple.com/documentation/avfoundation/avpartialasyncproperty/tracks-44ptx
func (u AVURLAsset) Tracks() IAVAssetTrack {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("tracks"))
	return AVAssetTrackFromID(objc.ID(rv))
}
func (u AVURLAsset) SetTracks(value IAVAssetTrack) {
	objc.Send[struct{}](u.ID, objc.Sel("setTracks:"), value)
}
// The resource loader for the asset.
//
// # Discussion
// 
// During loading, the system may ask the resource loader to assist loading
// the resource. For example, a resource that requires decryption may require
// the resource loader to provide the appropriate decryption keys. You can
// assign a delegate object to the resource loader object and use your
// delegate to intercept these requests and provide appropriate responses.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/resourceLoader
func (u AVURLAsset) ResourceLoader() IAVAssetResourceLoader {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("resourceLoader"))
	return AVAssetResourceLoaderFromID(objc.ID(rv))
}
// A Boolean value that indicates whether you can add this asset as a content
// key recipient to a content key session.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/mayRequireContentKeysForMediaDataProcessing
func (u AVURLAsset) MayRequireContentKeysForMediaDataProcessing() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("mayRequireContentKeysForMediaDataProcessing"))
	return rv
}
// The asset’s associated asset cache, if it exists.
//
// # Discussion
// 
// This property provides access to an instance of [AVAssetCache] to use for
// inspection of locally cached media data. The value of this property is
// `nil` if you haven’t configured the asset to store or access media data
// from disk.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/assetCache
func (u AVURLAsset) AssetCache() IAVAssetCache {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("assetCache"))
	return AVAssetCacheFromID(objc.ID(rv))
}
// A URL to the asset’s media.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/url
func (u AVURLAsset) URL() foundation.INSURL {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
// A session identifier that the asset sends in HTTP requests that it makes.
//
// # Discussion
// 
// The asset uses this value to set as the `X-Playback-Session-Id` header of
// HTTP requests that it creates.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/httpSessionIdentifier
func (u AVURLAsset) HttpSessionIdentifier() foundation.NSUUID {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("httpSessionIdentifier"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}
// The properties of the media extension format reader that decodes the asset.
//
// # Discussion
// 
// If the system decodes the asset using a MediaExtension format reader, the
// property value contains a valid object that describes the extension.
// Otherwise, this property value is `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/mediaExtensionProperties
func (u AVURLAsset) MediaExtensionProperties() IAVMediaExtensionProperties {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("mediaExtensionProperties"))
	return AVMediaExtensionPropertiesFromID(objc.ID(rv))
}
// The sidecar URL used by the MediaExtension.
//
// # Discussion
// 
// The sidecar URL is returned only if the MediaExtension format reader
// supports sidecar files, and implements this property [MEFileInfo
// setSidecarFilename:]. Will return nil otherwise.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/sidecarURL
func (u AVURLAsset) SidecarURL() foundation.INSURL {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("sidecarURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// Provides the content types the AVURLAsset class understands.
//
// # Return Value
// 
// An NSArray of UTTypes identifying the content types the AVURLAsset class
// understands.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/audiovisualContentTypes
func (_AVURLAssetClass AVURLAssetClass) AudiovisualContentTypes() []uniformtypeidentifiers.UTType {
	rv := objc.Send[[]objc.ID](objc.ID(_AVURLAssetClass.class), objc.Sel("audiovisualContentTypes"))
	return objc.ConvertSlice(rv, func(id objc.ID) uniformtypeidentifiers.UTType {
		return uniformtypeidentifiers.UTTypeFromID(id)
	})
}

			// Protocol methods for AVContentKeyRecipient
			

// FindCompatibleTrackForCompositionTrack is a synchronous wrapper around [AVURLAsset.FindCompatibleTrackForCompositionTrackCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (u AVURLAsset) FindCompatibleTrackForCompositionTrack(ctx context.Context, compositionTrack IAVCompositionTrack) (*AVAssetTrack, error) {
	type result struct {
		val *AVAssetTrack
		err error
	}
	done := make(chan result, 1)
	u.FindCompatibleTrackForCompositionTrackCompletionHandler(compositionTrack, func(val *AVAssetTrack, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

