// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAssetDownloadConfiguration] class.
var (
	_AVAssetDownloadConfigurationClass     AVAssetDownloadConfigurationClass
	_AVAssetDownloadConfigurationClassOnce sync.Once
)

func getAVAssetDownloadConfigurationClass() AVAssetDownloadConfigurationClass {
	_AVAssetDownloadConfigurationClassOnce.Do(func() {
		_AVAssetDownloadConfigurationClass = AVAssetDownloadConfigurationClass{class: objc.GetClass("AVAssetDownloadConfiguration")}
	})
	return _AVAssetDownloadConfigurationClass
}

// GetAVAssetDownloadConfigurationClass returns the class object for AVAssetDownloadConfiguration.
func GetAVAssetDownloadConfigurationClass() AVAssetDownloadConfigurationClass {
	return getAVAssetDownloadConfigurationClass()
}

type AVAssetDownloadConfigurationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetDownloadConfigurationClass) Alloc() AVAssetDownloadConfiguration {
	rv := objc.Send[AVAssetDownloadConfiguration](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that provides the configuration for a download task.
//
// # Accessing configuration details
//
//   - [AVAssetDownloadConfiguration.ArtworkData]: A data value that represents the asset’s artwork.
//   - [AVAssetDownloadConfiguration.SetArtworkData]
//   - [AVAssetDownloadConfiguration.PrimaryContentConfiguration]: The configuration for the primary content that the task downloads.
//   - [AVAssetDownloadConfiguration.AuxiliaryContentConfigurations]: The configuration for the auxiliary content that the task downloads.
//   - [AVAssetDownloadConfiguration.SetAuxiliaryContentConfigurations]
//   - [AVAssetDownloadConfiguration.OptimizesAuxiliaryContentConfigurations]: A Boolean value that indicates whether the task optimizes auxiliary content selection.
//   - [AVAssetDownloadConfiguration.SetOptimizesAuxiliaryContentConfigurations]
//   - [AVAssetDownloadConfiguration.SetInterstitialMediaSelectionCriteriaForMediaCharacteristic]: Sets media selection on interstitials for this asset
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadConfiguration
type AVAssetDownloadConfiguration struct {
	objectivec.Object
}

// AVAssetDownloadConfigurationFromID constructs a [AVAssetDownloadConfiguration] from an objc.ID.
//
// An object that provides the configuration for a download task.
func AVAssetDownloadConfigurationFromID(id objc.ID) AVAssetDownloadConfiguration {
	return AVAssetDownloadConfiguration{objectivec.Object{ID: id}}
}
// NOTE: AVAssetDownloadConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetDownloadConfiguration] class.
//
// # Accessing configuration details
//
//   - [IAVAssetDownloadConfiguration.ArtworkData]: A data value that represents the asset’s artwork.
//   - [IAVAssetDownloadConfiguration.SetArtworkData]
//   - [IAVAssetDownloadConfiguration.PrimaryContentConfiguration]: The configuration for the primary content that the task downloads.
//   - [IAVAssetDownloadConfiguration.AuxiliaryContentConfigurations]: The configuration for the auxiliary content that the task downloads.
//   - [IAVAssetDownloadConfiguration.SetAuxiliaryContentConfigurations]
//   - [IAVAssetDownloadConfiguration.OptimizesAuxiliaryContentConfigurations]: A Boolean value that indicates whether the task optimizes auxiliary content selection.
//   - [IAVAssetDownloadConfiguration.SetOptimizesAuxiliaryContentConfigurations]
//   - [IAVAssetDownloadConfiguration.SetInterstitialMediaSelectionCriteriaForMediaCharacteristic]: Sets media selection on interstitials for this asset
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadConfiguration
type IAVAssetDownloadConfiguration interface {
	objectivec.IObject

	// Topic: Accessing configuration details

	// A data value that represents the asset’s artwork.
	ArtworkData() foundation.INSData
	SetArtworkData(value foundation.INSData)
	// The configuration for the primary content that the task downloads.
	PrimaryContentConfiguration() IAVAssetDownloadContentConfiguration
	// The configuration for the auxiliary content that the task downloads.
	AuxiliaryContentConfigurations() []AVAssetDownloadContentConfiguration
	SetAuxiliaryContentConfigurations(value []AVAssetDownloadContentConfiguration)
	// A Boolean value that indicates whether the task optimizes auxiliary content selection.
	OptimizesAuxiliaryContentConfigurations() bool
	SetOptimizesAuxiliaryContentConfigurations(value bool)
	// Sets media selection on interstitials for this asset
	SetInterstitialMediaSelectionCriteriaForMediaCharacteristic(criteria []AVPlayerMediaSelectionCriteria, mediaCharacteristic AVMediaCharacteristic)

	// Download interstitial assets as listed in the index file. False by default.
	DownloadsInterstitialAssets() bool
	SetDownloadsInterstitialAssets(value bool)
}

// Init initializes the instance.
func (a AVAssetDownloadConfiguration) Init() AVAssetDownloadConfiguration {
	rv := objc.Send[AVAssetDownloadConfiguration](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetDownloadConfiguration) Autorelease() AVAssetDownloadConfiguration {
	rv := objc.Send[AVAssetDownloadConfiguration](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetDownloadConfiguration creates a new AVAssetDownloadConfiguration instance.
func NewAVAssetDownloadConfiguration() AVAssetDownloadConfiguration {
	class := getAVAssetDownloadConfigurationClass()
	rv := objc.Send[AVAssetDownloadConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a download configuration for a media asset.
//
// asset: The asset the task downloads.
//
// title: A human-readable title for this asset. The system displays this value in
// the usage pane of the Settings app; choose a title suitable for display in
// the user’s preferred language.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadConfiguration/init(asset:title:)
func NewAssetDownloadConfigurationWithAssetTitle(asset IAVURLAsset, title string) AVAssetDownloadConfiguration {
	rv := objc.Send[objc.ID](objc.ID(getAVAssetDownloadConfigurationClass().class), objc.Sel("downloadConfigurationWithAsset:title:"), asset, objc.String(title))
	return AVAssetDownloadConfigurationFromID(rv)
}

// Sets media selection on interstitials for this asset
//
// criteria: The array of selection criteria to set
//
// mediaCharacteristic: The AVMediaCharacteristic to which the criteria will be applied
//
// # Discussion
// 
// Typically, interstitial assets have not been discovered when the main
// download is initiated. This method allows the user to specify
// AVMediaSelectionCriteria for all interstitials that are discovered. Each
// AVPlayerMediaSelectionCriteria in the array of criteria specfies a set of
// criteria for a variant to download.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadConfiguration/setInterstitialMediaSelectionCriteria(_:forMediaCharacteristic:)
func (a AVAssetDownloadConfiguration) SetInterstitialMediaSelectionCriteriaForMediaCharacteristic(criteria []AVPlayerMediaSelectionCriteria, mediaCharacteristic AVMediaCharacteristic) {
	objc.Send[objc.ID](a.ID, objc.Sel("setInterstitialMediaSelectionCriteria:forMediaCharacteristic:"), objectivec.IObjectSliceToNSArray(criteria), objc.String(string(mediaCharacteristic)))
}

// A data value that represents the asset’s artwork.
//
// # Discussion
// 
// The system displays this image in the usage pane of the Settings app.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadConfiguration/artworkData
func (a AVAssetDownloadConfiguration) ArtworkData() foundation.INSData {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("artworkData"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (a AVAssetDownloadConfiguration) SetArtworkData(value foundation.INSData) {
	objc.Send[struct{}](a.ID, objc.Sel("setArtworkData:"), value)
}
// The configuration for the primary content that the task downloads.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadConfiguration/primaryContentConfiguration
func (a AVAssetDownloadConfiguration) PrimaryContentConfiguration() IAVAssetDownloadContentConfiguration {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("primaryContentConfiguration"))
	return AVAssetDownloadContentConfigurationFromID(objc.ID(rv))
}
// The configuration for the auxiliary content that the task downloads.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadConfiguration/auxiliaryContentConfigurations
func (a AVAssetDownloadConfiguration) AuxiliaryContentConfigurations() []AVAssetDownloadContentConfiguration {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("auxiliaryContentConfigurations"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVAssetDownloadContentConfiguration {
		return AVAssetDownloadContentConfigurationFromID(id)
	})
}
func (a AVAssetDownloadConfiguration) SetAuxiliaryContentConfigurations(value []AVAssetDownloadContentConfiguration) {
	objc.Send[struct{}](a.ID, objc.Sel("setAuxiliaryContentConfigurations:"), objectivec.IObjectSliceToNSArray(value))
}
// A Boolean value that indicates whether the task optimizes auxiliary content
// selection.
//
// # Discussion
// 
// By default, a download task optimizes its selection of auxiliary content
// based on its primary download content. For example, if the primary content
// configuration represents stereo renditions, and auxiliary content
// configuration represents multichannel audio renditions, the task choses the
// auxiliary multichannel variant to avoid downloading duplicate video
// renditions.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadConfiguration/optimizesAuxiliaryContentConfigurations
func (a AVAssetDownloadConfiguration) OptimizesAuxiliaryContentConfigurations() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("optimizesAuxiliaryContentConfigurations"))
	return rv
}
func (a AVAssetDownloadConfiguration) SetOptimizesAuxiliaryContentConfigurations(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setOptimizesAuxiliaryContentConfigurations:"), value)
}
// Download interstitial assets as listed in the index file. False by default.
//
// # Discussion
// 
// Ordinarily, interstitial assets are skipped when downloading content for
// later playback. Setting this property to true will cause interstitial
// assets to be downloaded as well. Playback of the downloaded content can
// then match the experience of online streaming playback as closely as
// possible.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadConfiguration/downloadsInterstitialAssets
func (a AVAssetDownloadConfiguration) DownloadsInterstitialAssets() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("downloadsInterstitialAssets"))
	return rv
}
func (a AVAssetDownloadConfiguration) SetDownloadsInterstitialAssets(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setDownloadsInterstitialAssets:"), value)
}

