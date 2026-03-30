// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAssetVariantVideoLayoutAttributes] class.
var (
	_AVAssetVariantVideoLayoutAttributesClass     AVAssetVariantVideoLayoutAttributesClass
	_AVAssetVariantVideoLayoutAttributesClassOnce sync.Once
)

func getAVAssetVariantVideoLayoutAttributesClass() AVAssetVariantVideoLayoutAttributesClass {
	_AVAssetVariantVideoLayoutAttributesClassOnce.Do(func() {
		_AVAssetVariantVideoLayoutAttributesClass = AVAssetVariantVideoLayoutAttributesClass{class: objc.GetClass("AVAssetVariantVideoLayoutAttributes")}
	})
	return _AVAssetVariantVideoLayoutAttributesClass
}

// GetAVAssetVariantVideoLayoutAttributesClass returns the class object for AVAssetVariantVideoLayoutAttributes.
func GetAVAssetVariantVideoLayoutAttributesClass() AVAssetVariantVideoLayoutAttributesClass {
	return getAVAssetVariantVideoLayoutAttributesClass()
}

type AVAssetVariantVideoLayoutAttributesClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAssetVariantVideoLayoutAttributesClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetVariantVideoLayoutAttributesClass) Alloc() AVAssetVariantVideoLayoutAttributes {
	rv := objc.Send[AVAssetVariantVideoLayoutAttributes](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// Attributes that describe the layout of video content.
//
// # Accessing attributes
//
//   - [AVAssetVariantVideoLayoutAttributes.StereoViewComponents]: Attributes that describe the video’s stereo components.
//   - [AVAssetVariantVideoLayoutAttributes.ProjectionType]: Describes the video projection.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/VideoAttributes-swift.class/LayoutAttributes
type AVAssetVariantVideoLayoutAttributes struct {
	objectivec.Object
}

// AVAssetVariantVideoLayoutAttributesFromID constructs a [AVAssetVariantVideoLayoutAttributes] from an objc.ID.
//
// Attributes that describe the layout of video content.
func AVAssetVariantVideoLayoutAttributesFromID(id objc.ID) AVAssetVariantVideoLayoutAttributes {
	return AVAssetVariantVideoLayoutAttributes{objectivec.Object{ID: id}}
}

// NOTE: AVAssetVariantVideoLayoutAttributes adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetVariantVideoLayoutAttributes] class.
//
// # Accessing attributes
//
//   - [IAVAssetVariantVideoLayoutAttributes.StereoViewComponents]: Attributes that describe the video’s stereo components.
//   - [IAVAssetVariantVideoLayoutAttributes.ProjectionType]: Describes the video projection.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/VideoAttributes-swift.class/LayoutAttributes
type IAVAssetVariantVideoLayoutAttributes interface {
	objectivec.IObject

	// Topic: Accessing attributes

	// Attributes that describe the video’s stereo components.
	StereoViewComponents() uint64
	// Describes the video projection.
	ProjectionType() uint64

	// The video sample codec types present in the variant’s renditions.
	CodecTypes() uint32
	SetCodecTypes(value uint32)
	// The nominal frame rate of the variant’s renditions.
	NominalFrameRate() float64
	SetNominalFrameRate(value float64)
	// The presentation size of the variant’s renditions.
	PresentationSize() corefoundation.CGSize
	SetPresentationSize(value corefoundation.CGSize)
	// Attributes that describe the layout of the video content.
	VideoLayoutAttributes() IAVAssetVariantVideoLayoutAttributes
	SetVideoLayoutAttributes(value IAVAssetVariantVideoLayoutAttributes)
	// The video range of the variant.
	VideoRange() AVVideoRange
	SetVideoRange(value AVVideoRange)
}

// Init initializes the instance.
func (a AVAssetVariantVideoLayoutAttributes) Init() AVAssetVariantVideoLayoutAttributes {
	rv := objc.Send[AVAssetVariantVideoLayoutAttributes](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetVariantVideoLayoutAttributes) Autorelease() AVAssetVariantVideoLayoutAttributes {
	rv := objc.Send[AVAssetVariantVideoLayoutAttributes](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetVariantVideoLayoutAttributes creates a new AVAssetVariantVideoLayoutAttributes instance.
func NewAVAssetVariantVideoLayoutAttributes() AVAssetVariantVideoLayoutAttributes {
	class := getAVAssetVariantVideoLayoutAttributesClass()
	rv := objc.Send[AVAssetVariantVideoLayoutAttributes](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Attributes that describe the video’s stereo components.
//
// # Discussion
//
// In the case of 3D or stereoscopic content, the value contains [leftEye] and
// [rightEye] components. In the case of monoscopic content, this value is
// [kCMStereoView_None].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/VideoAttributes-swift.class/LayoutAttributes/stereoViewComponents
//
// [kCMStereoView_None]: https://developer.apple.com/documentation/CoreMedia/CMStereoViewComponents/kCMStereoView_None
// [leftEye]: https://developer.apple.com/documentation/CoreMedia/CMStereoViewComponents/leftEye
// [rightEye]: https://developer.apple.com/documentation/CoreMedia/CMStereoViewComponents/rightEye
func (a AVAssetVariantVideoLayoutAttributes) StereoViewComponents() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("stereoViewComponents"))
	return rv
}

// Describes the video projection.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/VideoAttributes-swift.class/LayoutAttributes/projectionType
func (a AVAssetVariantVideoLayoutAttributes) ProjectionType() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("projectionType"))
	return rv
}

// The video sample codec types present in the variant’s renditions.
//
// See: https://developer.apple.com/documentation/avfoundation/avassetvariant/videoattributes-swift.class/codectypes
func (a AVAssetVariantVideoLayoutAttributes) CodecTypes() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("codecTypes"))
	return rv
}
func (a AVAssetVariantVideoLayoutAttributes) SetCodecTypes(value uint32) {
	objc.Send[struct{}](a.ID, objc.Sel("setCodecTypes:"), value)
}

// The nominal frame rate of the variant’s renditions.
//
// See: https://developer.apple.com/documentation/avfoundation/avassetvariant/videoattributes-swift.class/nominalframerate
func (a AVAssetVariantVideoLayoutAttributes) NominalFrameRate() float64 {
	rv := objc.Send[float64](a.ID, objc.Sel("nominalFrameRate"))
	return rv
}
func (a AVAssetVariantVideoLayoutAttributes) SetNominalFrameRate(value float64) {
	objc.Send[struct{}](a.ID, objc.Sel("setNominalFrameRate:"), value)
}

// The presentation size of the variant’s renditions.
//
// See: https://developer.apple.com/documentation/avfoundation/avassetvariant/videoattributes-swift.class/presentationsize
func (a AVAssetVariantVideoLayoutAttributes) PresentationSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](a.ID, objc.Sel("presentationSize"))
	return corefoundation.CGSize(rv)
}
func (a AVAssetVariantVideoLayoutAttributes) SetPresentationSize(value corefoundation.CGSize) {
	objc.Send[struct{}](a.ID, objc.Sel("setPresentationSize:"), value)
}

// Attributes that describe the layout of the video content.
//
// See: https://developer.apple.com/documentation/avfoundation/avassetvariant/videoattributes-swift.class/videolayoutattributes
func (a AVAssetVariantVideoLayoutAttributes) VideoLayoutAttributes() IAVAssetVariantVideoLayoutAttributes {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("videoLayoutAttributes"))
	return AVAssetVariantVideoLayoutAttributesFromID(objc.ID(rv))
}
func (a AVAssetVariantVideoLayoutAttributes) SetVideoLayoutAttributes(value IAVAssetVariantVideoLayoutAttributes) {
	objc.Send[struct{}](a.ID, objc.Sel("setVideoLayoutAttributes:"), value)
}

// The video range of the variant.
//
// See: https://developer.apple.com/documentation/avfoundation/avassetvariant/videoattributes-swift.class/videorange
func (a AVAssetVariantVideoLayoutAttributes) VideoRange() AVVideoRange {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("videoRange"))
	return AVVideoRange(foundation.NSStringFromID(rv).String())
}
func (a AVAssetVariantVideoLayoutAttributes) SetVideoRange(value AVVideoRange) {
	objc.Send[struct{}](a.ID, objc.Sel("setVideoRange:"), objc.String(string(value)))
}
