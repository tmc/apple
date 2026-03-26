// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAssetDownloadContentConfiguration] class.
var (
	_AVAssetDownloadContentConfigurationClass     AVAssetDownloadContentConfigurationClass
	_AVAssetDownloadContentConfigurationClassOnce sync.Once
)

func getAVAssetDownloadContentConfigurationClass() AVAssetDownloadContentConfigurationClass {
	_AVAssetDownloadContentConfigurationClassOnce.Do(func() {
		_AVAssetDownloadContentConfigurationClass = AVAssetDownloadContentConfigurationClass{class: objc.GetClass("AVAssetDownloadContentConfiguration")}
	})
	return _AVAssetDownloadContentConfigurationClass
}

// GetAVAssetDownloadContentConfigurationClass returns the class object for AVAssetDownloadContentConfiguration.
func GetAVAssetDownloadContentConfigurationClass() AVAssetDownloadContentConfigurationClass {
	return getAVAssetDownloadContentConfigurationClass()
}

type AVAssetDownloadContentConfigurationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetDownloadContentConfigurationClass) Alloc() AVAssetDownloadContentConfiguration {
	rv := objc.Send[AVAssetDownloadContentConfiguration](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A configuration object that contains variant qualifiers and media options.
//
// # Accessing configuration details
//
//   - [AVAssetDownloadContentConfiguration.VariantQualifiers]: The variant qualifiers for this configuration.
//   - [AVAssetDownloadContentConfiguration.SetVariantQualifiers]
//   - [AVAssetDownloadContentConfiguration.MediaSelections]: The media selections of an asset that a task downloads.
//   - [AVAssetDownloadContentConfiguration.SetMediaSelections]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadContentConfiguration
type AVAssetDownloadContentConfiguration struct {
	objectivec.Object
}

// AVAssetDownloadContentConfigurationFromID constructs a [AVAssetDownloadContentConfiguration] from an objc.ID.
//
// A configuration object that contains variant qualifiers and media options.
func AVAssetDownloadContentConfigurationFromID(id objc.ID) AVAssetDownloadContentConfiguration {
	return AVAssetDownloadContentConfiguration{objectivec.Object{ID: id}}
}
// NOTE: AVAssetDownloadContentConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetDownloadContentConfiguration] class.
//
// # Accessing configuration details
//
//   - [IAVAssetDownloadContentConfiguration.VariantQualifiers]: The variant qualifiers for this configuration.
//   - [IAVAssetDownloadContentConfiguration.SetVariantQualifiers]
//   - [IAVAssetDownloadContentConfiguration.MediaSelections]: The media selections of an asset that a task downloads.
//   - [IAVAssetDownloadContentConfiguration.SetMediaSelections]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadContentConfiguration
type IAVAssetDownloadContentConfiguration interface {
	objectivec.IObject

	// Topic: Accessing configuration details

	// The variant qualifiers for this configuration.
	VariantQualifiers() []AVAssetVariantQualifier
	SetVariantQualifiers(value []AVAssetVariantQualifier)
	// The media selections of an asset that a task downloads.
	MediaSelections() []AVMediaSelection
	SetMediaSelections(value []AVMediaSelection)

	// A data value that represents the asset’s artwork.
	ArtworkData() foundation.INSData
	SetArtworkData(value foundation.INSData)
	// The configuration for the auxiliary content that the task downloads.
	AuxiliaryContentConfigurations() IAVAssetDownloadContentConfiguration
	SetAuxiliaryContentConfigurations(value IAVAssetDownloadContentConfiguration)
	// A Boolean value that indicates whether the task optimizes auxiliary content selection.
	OptimizesAuxiliaryContentConfigurations() bool
	SetOptimizesAuxiliaryContentConfigurations(value bool)
	// The configuration for the primary content that the task downloads.
	PrimaryContentConfiguration() IAVAssetDownloadContentConfiguration
	SetPrimaryContentConfiguration(value IAVAssetDownloadContentConfiguration)
}

// Init initializes the instance.
func (a AVAssetDownloadContentConfiguration) Init() AVAssetDownloadContentConfiguration {
	rv := objc.Send[AVAssetDownloadContentConfiguration](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetDownloadContentConfiguration) Autorelease() AVAssetDownloadContentConfiguration {
	rv := objc.Send[AVAssetDownloadContentConfiguration](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetDownloadContentConfiguration creates a new AVAssetDownloadContentConfiguration instance.
func NewAVAssetDownloadContentConfiguration() AVAssetDownloadContentConfiguration {
	class := getAVAssetDownloadContentConfigurationClass()
	rv := objc.Send[AVAssetDownloadContentConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The variant qualifiers for this configuration.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadContentConfiguration/variantQualifiers
func (a AVAssetDownloadContentConfiguration) VariantQualifiers() []AVAssetVariantQualifier {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("variantQualifiers"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVAssetVariantQualifier {
		return AVAssetVariantQualifierFromID(id)
	})
}
func (a AVAssetDownloadContentConfiguration) SetVariantQualifiers(value []AVAssetVariantQualifier) {
	objc.Send[struct{}](a.ID, objc.Sel("setVariantQualifiers:"), objectivec.IObjectSliceToNSArray(value))
}
// The media selections of an asset that a task downloads.
//
// # Discussion
// 
// If your configuration doesn’t indicate a media selection, the system uses
// the asset’s automatic media selection.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadContentConfiguration/mediaSelections
func (a AVAssetDownloadContentConfiguration) MediaSelections() []AVMediaSelection {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("mediaSelections"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVMediaSelection {
		return AVMediaSelectionFromID(id)
	})
}
func (a AVAssetDownloadContentConfiguration) SetMediaSelections(value []AVMediaSelection) {
	objc.Send[struct{}](a.ID, objc.Sel("setMediaSelections:"), objectivec.IObjectSliceToNSArray(value))
}
// A data value that represents the asset’s artwork.
//
// See: https://developer.apple.com/documentation/avfoundation/avassetdownloadconfiguration/artworkdata
func (a AVAssetDownloadContentConfiguration) ArtworkData() foundation.INSData {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("artworkData"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (a AVAssetDownloadContentConfiguration) SetArtworkData(value foundation.INSData) {
	objc.Send[struct{}](a.ID, objc.Sel("setArtworkData:"), value)
}
// The configuration for the auxiliary content that the task downloads.
//
// See: https://developer.apple.com/documentation/avfoundation/avassetdownloadconfiguration/auxiliarycontentconfigurations
func (a AVAssetDownloadContentConfiguration) AuxiliaryContentConfigurations() IAVAssetDownloadContentConfiguration {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("auxiliaryContentConfigurations"))
	return AVAssetDownloadContentConfigurationFromID(objc.ID(rv))
}
func (a AVAssetDownloadContentConfiguration) SetAuxiliaryContentConfigurations(value IAVAssetDownloadContentConfiguration) {
	objc.Send[struct{}](a.ID, objc.Sel("setAuxiliaryContentConfigurations:"), value)
}
// A Boolean value that indicates whether the task optimizes auxiliary content
// selection.
//
// See: https://developer.apple.com/documentation/avfoundation/avassetdownloadconfiguration/optimizesauxiliarycontentconfigurations
func (a AVAssetDownloadContentConfiguration) OptimizesAuxiliaryContentConfigurations() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("optimizesAuxiliaryContentConfigurations"))
	return rv
}
func (a AVAssetDownloadContentConfiguration) SetOptimizesAuxiliaryContentConfigurations(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setOptimizesAuxiliaryContentConfigurations:"), value)
}
// The configuration for the primary content that the task downloads.
//
// See: https://developer.apple.com/documentation/avfoundation/avassetdownloadconfiguration/primarycontentconfiguration
func (a AVAssetDownloadContentConfiguration) PrimaryContentConfiguration() IAVAssetDownloadContentConfiguration {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("primaryContentConfiguration"))
	return AVAssetDownloadContentConfigurationFromID(objc.ID(rv))
}
func (a AVAssetDownloadContentConfiguration) SetPrimaryContentConfiguration(value IAVAssetDownloadContentConfiguration) {
	objc.Send[struct{}](a.ID, objc.Sel("setPrimaryContentConfiguration:"), value)
}

