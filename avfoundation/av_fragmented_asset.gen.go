// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVFragmentedAsset] class.
var (
	_AVFragmentedAssetClass     AVFragmentedAssetClass
	_AVFragmentedAssetClassOnce sync.Once
)

func getAVFragmentedAssetClass() AVFragmentedAssetClass {
	_AVFragmentedAssetClassOnce.Do(func() {
		_AVFragmentedAssetClass = AVFragmentedAssetClass{class: objc.GetClass("AVFragmentedAsset")}
	})
	return _AVFragmentedAssetClass
}

// GetAVFragmentedAssetClass returns the class object for AVFragmentedAsset.
func GetAVFragmentedAssetClass() AVFragmentedAssetClass {
	return getAVFragmentedAssetClass()
}

type AVFragmentedAssetClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVFragmentedAssetClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVFragmentedAssetClass) Alloc() AVFragmentedAsset {
	rv := objc.Send[AVFragmentedAsset](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An asset with a duration that the system can extend without modifying its
// existing media data.
//
// # Overview
//
// By using an `mvex` box in their `moov` box, QuickTime movie files and
// MPEG-4 files can indicate that they accommodate additional fragments. To
// determine whether a fragmented asset can monitor the addition of fragments,
// check the value of its [AVFragmentedAsset.CanContainFragments] property.
//
// Associate a fragmented asset with an instance of [AVFragmentedAssetMinder]
// to know when the system appends new fragments. When it has an associated
// asset minder, [AVFragmentedAssetTrack] posts
// [AVAssetDurationDidChangeNotification] notifications whenever it detects
// new fragments. It may also post
// [AVAssetContainsFragmentsDidChangeNotification] and
// [AVAssetWasDefragmentedNotification], as the documentation of those
// notifications explains.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFragmentedAsset
//
// [AVAssetContainsFragmentsDidChangeNotification]: https://developer.apple.com/documentation/AVFoundation/AVAssetContainsFragmentsDidChangeNotification
// [AVAssetDurationDidChangeNotification]: https://developer.apple.com/documentation/AVFoundation/AVAssetDurationDidChangeNotification
// [AVAssetWasDefragmentedNotification]: https://developer.apple.com/documentation/AVFoundation/AVAssetWasDefragmentedNotification
type AVFragmentedAsset struct {
	AVURLAsset
}

// AVFragmentedAssetFromID constructs a [AVFragmentedAsset] from an objc.ID.
//
// An asset with a duration that the system can extend without modifying its
// existing media data.
func AVFragmentedAssetFromID(id objc.ID) AVFragmentedAsset {
	return AVFragmentedAsset{AVURLAsset: AVURLAssetFromID(id)}
}

// NOTE: AVFragmentedAsset adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVFragmentedAsset] class.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFragmentedAsset
type IAVFragmentedAsset interface {
	IAVURLAsset
	AVFragmentMinding

	// A Boolean value that indicates whether you can extend the asset by fragments.
	CanContainFragments() bool
	SetCanContainFragments(value bool)
}

// Init initializes the instance.
func (f AVFragmentedAsset) Init() AVFragmentedAsset {
	rv := objc.Send[AVFragmentedAsset](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f AVFragmentedAsset) Autorelease() AVFragmentedAsset {
	rv := objc.Send[AVFragmentedAsset](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVFragmentedAsset creates a new AVFragmentedAsset instance.
func NewAVFragmentedAsset() AVFragmentedAsset {
	class := getAVFragmentedAssetClass()
	rv := objc.Send[AVFragmentedAsset](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an asset that models the media at the specified URL.
//
// URL: A URL to a local, remote, or HTTP Live Streaming media resource.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAsset/init(url:)
func NewFragmentedAssetWithURL(URL foundation.INSURL) AVFragmentedAsset {
	rv := objc.Send[objc.ID](objc.ID(getAVFragmentedAssetClass().class), objc.Sel("assetWithURL:"), URL)
	return AVFragmentedAssetFromID(rv)
}

// Creates an asset that models the media resource at the specified URL.
//
// URL: A URL that references the media for the asset to model.
//
// options: A dictionary that contains options used to customize the initialization of
// the asset.
//
// For supported keys and values, see [Initialization options].
//
// # Return Value
//
// An asset that models the media resource found at [URL].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/init(url:options:)
//
// [Initialization options]: https://developer.apple.com/documentation/AVFoundation/initialization-options
func NewFragmentedAssetWithURLOptions(URL foundation.INSURL, options foundation.INSDictionary) AVFragmentedAsset {
	instance := getAVFragmentedAssetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:options:"), URL, options)
	return AVFragmentedAssetFromID(rv)
}

// A Boolean value that indicates whether an asset that supports fragment
// minding is currently associated with a fragment minder.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFragmentMinding/isAssociatedWithFragmentMinder
func (f AVFragmentedAsset) IsAssociatedWithFragmentMinder() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isAssociatedWithFragmentMinder"))
	return rv
}

// Creates a fragmented asset for the media at the specified URL.
//
// URL: A URL that points to the desired media resource.
//
// options: A dictionary of keys for specifying initialization options. Valid values
// are [AVURLAssetPreferPreciseDurationAndTimingKey] and
// [AVURLAssetReferenceRestrictionsKey].
//
// # Return Value
//
// A fragmented asset that models the media at the specified URL.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFragmentedAsset/fragmentedAssetWithURL:options:
//
// [AVURLAssetPreferPreciseDurationAndTimingKey]: https://developer.apple.com/documentation/AVFoundation/AVURLAssetPreferPreciseDurationAndTimingKey
// [AVURLAssetReferenceRestrictionsKey]: https://developer.apple.com/documentation/AVFoundation/AVURLAssetReferenceRestrictionsKey
func (_AVFragmentedAssetClass AVFragmentedAssetClass) FragmentedAssetWithURLOptions(URL foundation.INSURL, options foundation.INSDictionary) AVFragmentedAsset {
	rv := objc.Send[objc.ID](objc.ID(_AVFragmentedAssetClass.class), objc.Sel("fragmentedAssetWithURL:options:"), URL, options)
	return AVFragmentedAssetFromID(rv)
}

// A Boolean value that indicates whether you can extend the asset by
// fragments.
//
// See: https://developer.apple.com/documentation/avfoundation/avasset/cancontainfragments
func (f AVFragmentedAsset) CanContainFragments() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("canContainFragments"))
	return rv
}
func (f AVFragmentedAsset) SetCanContainFragments(value bool) {
	objc.Send[struct{}](f.ID, objc.Sel("setCanContainFragments:"), value)
}

// Protocol methods for AVFragmentMinding
