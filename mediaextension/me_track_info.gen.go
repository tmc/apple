// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [METrackInfo] class.
var (
	_METrackInfoClass     METrackInfoClass
	_METrackInfoClassOnce sync.Once
)

func getMETrackInfoClass() METrackInfoClass {
	_METrackInfoClassOnce.Do(func() {
		_METrackInfoClass = METrackInfoClass{class: objc.GetClass("METrackInfo")}
	})
	return _METrackInfoClass
}

// GetMETrackInfoClass returns the class object for METrackInfo.
func GetMETrackInfoClass() METrackInfoClass {
	return getMETrackInfoClass()
}

type METrackInfoClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc METrackInfoClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc METrackInfoClass) Alloc() METrackInfo {
	rv := objc.Send[METrackInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that includes track properties parsed from the media asset.
//
// # Inspecting track information
//
//   - [METrackInfo.MediaType]: The media type of the track.
//   - [METrackInfo.TrackID]: An integer that identifies the track within the media asset.
//   - [METrackInfo.IsEnabled]: A Boolean value that indicates whether the track is enabled by default.
//   - [METrackInfo.SetEnabled]
//   - [METrackInfo.NaturalTimescale]: The natural timescale of the track.
//   - [METrackInfo.SetNaturalTimescale]
//   - [METrackInfo.ExtendedLanguageTag]: A string that indicates the language tag associated with the track, as an IETF BCP 47 (RFC 4646) language identifier.
//   - [METrackInfo.SetExtendedLanguageTag]
//   - [METrackInfo.NaturalSize]: Indicates the natural dimensions of the media data referenced by the track.
//   - [METrackInfo.SetNaturalSize]
//   - [METrackInfo.PreferredTransform]: Indicates the preferred affine display transform of the track media for visual display.
//   - [METrackInfo.SetPreferredTransform]
//   - [METrackInfo.NominalFrameRate]: The frame rate of the track in frames per second, as a 32-bit floating point number.
//   - [METrackInfo.SetNominalFrameRate]
//   - [METrackInfo.RequiresFrameReordering]: A Boolean value that indicates whether frame reordering occurs in the track.
//   - [METrackInfo.SetRequiresFrameReordering]
//
// See: https://developer.apple.com/documentation/MediaExtension/METrackInfo
type METrackInfo struct {
	objectivec.Object
}

// METrackInfoFromID constructs a [METrackInfo] from an objc.ID.
//
// An object that includes track properties parsed from the media asset.
func METrackInfoFromID(id objc.ID) METrackInfo {
	return METrackInfo{objectivec.Object{ID: id}}
}

// NOTE: METrackInfo adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [METrackInfo] class.
//
// # Inspecting track information
//
//   - [IMETrackInfo.MediaType]: The media type of the track.
//   - [IMETrackInfo.TrackID]: An integer that identifies the track within the media asset.
//   - [IMETrackInfo.IsEnabled]: A Boolean value that indicates whether the track is enabled by default.
//   - [IMETrackInfo.SetEnabled]
//   - [IMETrackInfo.NaturalTimescale]: The natural timescale of the track.
//   - [IMETrackInfo.SetNaturalTimescale]
//   - [IMETrackInfo.ExtendedLanguageTag]: A string that indicates the language tag associated with the track, as an IETF BCP 47 (RFC 4646) language identifier.
//   - [IMETrackInfo.SetExtendedLanguageTag]
//   - [IMETrackInfo.NaturalSize]: Indicates the natural dimensions of the media data referenced by the track.
//   - [IMETrackInfo.SetNaturalSize]
//   - [IMETrackInfo.PreferredTransform]: Indicates the preferred affine display transform of the track media for visual display.
//   - [IMETrackInfo.SetPreferredTransform]
//   - [IMETrackInfo.NominalFrameRate]: The frame rate of the track in frames per second, as a 32-bit floating point number.
//   - [IMETrackInfo.SetNominalFrameRate]
//   - [IMETrackInfo.RequiresFrameReordering]: A Boolean value that indicates whether frame reordering occurs in the track.
//   - [IMETrackInfo.SetRequiresFrameReordering]
//
// See: https://developer.apple.com/documentation/MediaExtension/METrackInfo
type IMETrackInfo interface {
	objectivec.IObject

	// Topic: Inspecting track information

	// The media type of the track.
	MediaType() coremedia.CMMediaType
	// An integer that identifies the track within the media asset.
	TrackID() coremedia.CMPersistentTrackID
	// A Boolean value that indicates whether the track is enabled by default.
	IsEnabled() bool
	SetEnabled(value bool)
	// The natural timescale of the track.
	NaturalTimescale() coremedia.CMTimeScale
	SetNaturalTimescale(value coremedia.CMTimeScale)
	// A string that indicates the language tag associated with the track, as an IETF BCP 47 (RFC 4646) language identifier.
	ExtendedLanguageTag() string
	SetExtendedLanguageTag(value string)
	// Indicates the natural dimensions of the media data referenced by the track.
	NaturalSize() corefoundation.CGSize
	SetNaturalSize(value corefoundation.CGSize)
	// Indicates the preferred affine display transform of the track media for visual display.
	PreferredTransform() corefoundation.CGAffineTransform
	SetPreferredTransform(value corefoundation.CGAffineTransform)
	// The frame rate of the track in frames per second, as a 32-bit floating point number.
	NominalFrameRate() float32
	SetNominalFrameRate(value float32)
	// A Boolean value that indicates whether frame reordering occurs in the track.
	RequiresFrameReordering() bool
	SetRequiresFrameReordering(value bool)
}

// Init initializes the instance.
func (m METrackInfo) Init() METrackInfo {
	rv := objc.Send[METrackInfo](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m METrackInfo) Autorelease() METrackInfo {
	rv := objc.Send[METrackInfo](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMETrackInfo creates a new METrackInfo instance.
func NewMETrackInfo() METrackInfo {
	class := getMETrackInfoClass()
	rv := objc.Send[METrackInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The media type of the track.
//
// See: https://developer.apple.com/documentation/MediaExtension/METrackInfo/mediaType
func (m METrackInfo) MediaType() coremedia.CMMediaType {
	rv := objc.Send[coremedia.CMMediaType](m.ID, objc.Sel("mediaType"))
	return coremedia.CMMediaType(rv)
}

// An integer that identifies the track within the media asset.
//
// # Discussion
//
// The track ID uniquely identifes the track within a [MEFormatReader] object.
// Track IDs must be unique within a media asset but don’t need to be unique
// across assets. If a media format doesn’t have a native concept of track
// IDs, track IDs can start from `1`. However, track ID `0` is a reserved
// value to indicate an invalid track ID.
//
// See: https://developer.apple.com/documentation/MediaExtension/METrackInfo/trackID
func (m METrackInfo) TrackID() coremedia.CMPersistentTrackID {
	rv := objc.Send[coremedia.CMPersistentTrackID](m.ID, objc.Sel("trackID"))
	return coremedia.CMPersistentTrackID(rv)
}

// A Boolean value that indicates whether the track is enabled by default.
//
// See: https://developer.apple.com/documentation/MediaExtension/METrackInfo/isEnabled
func (m METrackInfo) IsEnabled() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isEnabled"))
	return rv
}
func (m METrackInfo) SetEnabled(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setEnabled:"), value)
}

// The natural timescale of the track.
//
// See: https://developer.apple.com/documentation/MediaExtension/METrackInfo/naturalTimescale
func (m METrackInfo) NaturalTimescale() coremedia.CMTimeScale {
	rv := objc.Send[coremedia.CMTimeScale](m.ID, objc.Sel("naturalTimescale"))
	return coremedia.CMTimeScale(rv)
}
func (m METrackInfo) SetNaturalTimescale(value coremedia.CMTimeScale) {
	objc.Send[struct{}](m.ID, objc.Sel("setNaturalTimescale:"), value)
}

// A string that indicates the language tag associated with the track, as an
// IETF BCP 47 (RFC 4646) language identifier.
//
// # Discussion
//
// [MediaToolbox] uses this property to group similar language tracks together
// or to match audio and caption tracks. Set this value to `nil` if the track
// doesn’t have a language tag.
//
// See: https://developer.apple.com/documentation/MediaExtension/METrackInfo/extendedLanguageTag
func (m METrackInfo) ExtendedLanguageTag() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("extendedLanguageTag"))
	return foundation.NSStringFromID(rv).String()
}
func (m METrackInfo) SetExtendedLanguageTag(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setExtendedLanguageTag:"), objc.String(value))
}

// Indicates the natural dimensions of the media data referenced by the track.
//
// # Discussion
//
// This property is only valid for tracks with visual media types and is
// [zero] for other track types.
//
// See: https://developer.apple.com/documentation/MediaExtension/METrackInfo/naturalSize
//
// [zero]: https://developer.apple.com/documentation/CoreFoundation/CGSize/zero
func (m METrackInfo) NaturalSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](m.ID, objc.Sel("naturalSize"))
	return corefoundation.CGSize(rv)
}
func (m METrackInfo) SetNaturalSize(value corefoundation.CGSize) {
	objc.Send[struct{}](m.ID, objc.Sel("setNaturalSize:"), value)
}

// Indicates the preferred affine display transform of the track media for
// visual display.
//
// # Discussion
//
// This property is only valid for tracks with visual media types and is
// [CGAffineTransformIdentity] for other track types.
//
// See: https://developer.apple.com/documentation/MediaExtension/METrackInfo/preferredTransform
//
// [CGAffineTransformIdentity]: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformIdentity
func (m METrackInfo) PreferredTransform() corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](m.ID, objc.Sel("preferredTransform"))
	return corefoundation.CGAffineTransform(rv)
}
func (m METrackInfo) SetPreferredTransform(value corefoundation.CGAffineTransform) {
	objc.Send[struct{}](m.ID, objc.Sel("setPreferredTransform:"), value)
}

// The frame rate of the track in frames per second, as a 32-bit floating
// point number.
//
// # Discussion
//
// For field-based video tracks that carry one field per media sample, the
// value of this property is the field rate, not the frame rate.
// [MediaToolbox] uses this property to calculate the maximum playback speed.
//
// See: https://developer.apple.com/documentation/MediaExtension/METrackInfo/nominalFrameRate
func (m METrackInfo) NominalFrameRate() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("nominalFrameRate"))
	return rv
}
func (m METrackInfo) SetNominalFrameRate(value float32) {
	objc.Send[struct{}](m.ID, objc.Sel("setNominalFrameRate:"), value)
}

// A Boolean value that indicates whether frame reordering occurs in the
// track.
//
// # Discussion
//
// The value is true if frame reordering occurs, otherwise false. This
// property is only valid for tracks with video media type and is false for
// other track types.
//
// See: https://developer.apple.com/documentation/MediaExtension/METrackInfo/requiresFrameReordering
func (m METrackInfo) RequiresFrameReordering() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("requiresFrameReordering"))
	return rv
}
func (m METrackInfo) SetRequiresFrameReordering(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setRequiresFrameReordering:"), value)
}
