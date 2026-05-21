// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/coremedia"
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
	StereoViewComponents() coremedia.CMStereoViewComponents
	// Describes the video projection.
	ProjectionType() coremedia.CMProjectionType
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
func (a AVAssetVariantVideoLayoutAttributes) StereoViewComponents() coremedia.CMStereoViewComponents {
	rv := objc.Send[coremedia.CMStereoViewComponents](a.ID, objc.Sel("stereoViewComponents"))
	return coremedia.CMStereoViewComponents(rv)
}

// Describes the video projection.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/VideoAttributes-swift.class/LayoutAttributes/projectionType
func (a AVAssetVariantVideoLayoutAttributes) ProjectionType() coremedia.CMProjectionType {
	rv := objc.Send[coremedia.CMProjectionType](a.ID, objc.Sel("projectionType"))
	return coremedia.CMProjectionType(rv)
}
