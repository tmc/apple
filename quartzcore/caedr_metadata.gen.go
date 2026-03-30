// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CAEDRMetadata] class.
var (
	_CAEDRMetadataClass     CAEDRMetadataClass
	_CAEDRMetadataClassOnce sync.Once
)

func getCAEDRMetadataClass() CAEDRMetadataClass {
	_CAEDRMetadataClassOnce.Do(func() {
		_CAEDRMetadataClass = CAEDRMetadataClass{class: objc.GetClass("CAEDRMetadata")}
	})
	return _CAEDRMetadataClass
}

// GetCAEDRMetadataClass returns the class object for CAEDRMetadata.
func GetCAEDRMetadataClass() CAEDRMetadataClass {
	return getCAEDRMetadataClass()
}

type CAEDRMetadataClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CAEDRMetadataClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CAEDRMetadataClass) Alloc() CAEDRMetadata {
	rv := objc.Send[CAEDRMetadata](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// Metadata describing how extended dynamic range (EDR) values should be tone
// mapped.
//
// # Overview
//
// If you need specific tone-mapping behavior, set the [EDRMetadata] property
// of a [CAMetalLayer] to point to an instance of this class.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEDRMetadata
type CAEDRMetadata struct {
	objectivec.Object
}

// CAEDRMetadataFromID constructs a [CAEDRMetadata] from an objc.ID.
//
// Metadata describing how extended dynamic range (EDR) values should be tone
// mapped.
func CAEDRMetadataFromID(id objc.ID) CAEDRMetadata {
	return CAEDRMetadata{objectivec.Object{ID: id}}
}

// NOTE: CAEDRMetadata adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CAEDRMetadata] class.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEDRMetadata
type ICAEDRMetadata interface {
	objectivec.IObject

	// Metadata describing the tone mapping to apply to the extended dynamic range (EDR) values in the layer.
	EdrMetadata() ICAEDRMetadata
	SetEdrMetadata(value ICAEDRMetadata)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (e CAEDRMetadata) Init() CAEDRMetadata {
	rv := objc.Send[CAEDRMetadata](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e CAEDRMetadata) Autorelease() CAEDRMetadata {
	rv := objc.Send[CAEDRMetadata](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewCAEDRMetadata creates a new CAEDRMetadata instance.
func NewCAEDRMetadata() CAEDRMetadata {
	class := getCAEDRMetadataClass()
	rv := objc.Send[CAEDRMetadata](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (e CAEDRMetadata) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](e.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Creates EDR metadata for HDR10 content based on mastering display color
// information and content light levels.
//
// displayData: A value of 24 bytes that contains a big-endian structure, as defined in
// D.2.28 (Mastering Display Colour Volume SEI message).
//
// contentData: A value of 4 bytes that contains a big-endian structure, as defined in
// D.2.35 (Content Light Level Information SEI message).
//
// scale: A scale factor relating (display-referred linear) EDR values to the optical
// output of the reference display.
//
// # Discussion
//
// The MDCV and CLLI message formats are defined in ISO/IEC 23008-2:2017.
//
// The values in the drawable’s texture are assumed to be proportional to
// the optical output (in cd/m^2) of the reference display. For example, if
// the optical output scale is 100, then a value of 1.0 is assumed to be 100
// nits.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEDRMetadata/hdr10(displayInfo:contentInfo:opticalOutputScale:)
func (_CAEDRMetadataClass CAEDRMetadataClass) HDR10MetadataWithDisplayInfoContentInfoOpticalOutputScale(displayData foundation.INSData, contentData foundation.INSData, scale float32) CAEDRMetadata {
	rv := objc.Send[objc.ID](objc.ID(_CAEDRMetadataClass.class), objc.Sel("HDR10MetadataWithDisplayInfo:contentInfo:opticalOutputScale:"), displayData, contentData, scale)
	return CAEDRMetadataFromID(rv)
}

// Creates EDR metadata for HDR10 content based on the luminance
// characteristics of a mastering display.
//
// minNits: The minimum nits (cd/m^2) of the mastering display.
//
// maxNits: The maximum nits (cd/m^2) of the mastering display.
//
// scale: A scale factor relating (display-referred linear) extended range buffer
// values to the optical output of a reference display.
//
// # Return Value
//
// A new EDR metadata object.
//
// # Discussion
//
// Any content greater than the maximum luminance (`maxNits`) may be clamped
// when displayed.
//
// The values in the drawable’s texture are assumed to be proportional to
// the optical output (in cd/m^2) of the reference display. For example, if
// the optical output scale is 100, then a value of 1.0 is assumed to be 100
// nits.
//
// If the content is in a normalized pixel format, set `opticalOutputScale` to
// 10000.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEDRMetadata/hdr10(minLuminance:maxLuminance:opticalOutputScale:)
func (_CAEDRMetadataClass CAEDRMetadataClass) HDR10MetadataWithMinLuminanceMaxLuminanceOpticalOutputScale(minNits float32, maxNits float32, scale float32) CAEDRMetadata {
	rv := objc.Send[objc.ID](objc.ID(_CAEDRMetadataClass.class), objc.Sel("HDR10MetadataWithMinLuminance:maxLuminance:opticalOutputScale:"), minNits, maxNits, scale)
	return CAEDRMetadataFromID(rv)
}

// See: https://developer.apple.com/documentation/QuartzCore/CAEDRMetadata/hlg(ambientViewingEnvironment:)
func (_CAEDRMetadataClass CAEDRMetadataClass) HLGMetadataWithAmbientViewingEnvironment(data foundation.INSData) CAEDRMetadata {
	rv := objc.Send[objc.ID](objc.ID(_CAEDRMetadataClass.class), objc.Sel("HLGMetadataWithAmbientViewingEnvironment:"), data)
	return CAEDRMetadataFromID(rv)
}

// Metadata describing the tone mapping to apply to the extended dynamic range
// (EDR) values in the layer.
//
// See: https://developer.apple.com/documentation/quartzcore/cametallayer/edrmetadata
func (e CAEDRMetadata) EdrMetadata() ICAEDRMetadata {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("EDRMetadata"))
	return CAEDRMetadataFromID(objc.ID(rv))
}
func (e CAEDRMetadata) SetEdrMetadata(value ICAEDRMetadata) {
	objc.Send[struct{}](e.ID, objc.Sel("setEDRMetadata:"), value)
}

// Extended dynamic range (EDR) metadata for the Hybrid Log-Gamma (HLG)
// transfer function.
//
// # Discussion
//
// Your content should be scene referred and encoded with the ITU-R BT.2100-2
// Hybrid Log Gamma (HLG) opto-electrical transfer function (OETF). The system
// applies the opto-optical transfer function (OOTF) based on peak display
// brightness and ambient lighting. If you’re rendering to a [CAMetalLayer]
// with a linear colorspace (for floating point EDR layers), you must apply
// the HLG inverse OETF without normalization, to provide a nominal range of
// `[0, 12]`.
//
// For more information on HLG, see [https://www.itu.int/rec/R-REC-BT.2100].
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEDRMetadata/hlg
//
// [https://www.itu.int/rec/R-REC-BT.2100]: https://www.itu.int/rec/R-REC-BT.2100
func (_CAEDRMetadataClass CAEDRMetadataClass) HLGMetadata() CAEDRMetadata {
	rv := objc.Send[objc.ID](objc.ID(_CAEDRMetadataClass.class), objc.Sel("HLGMetadata"))
	return CAEDRMetadataFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/QuartzCore/CAEDRMetadata/isAvailable
func (_CAEDRMetadataClass CAEDRMetadataClass) Available() bool {
	rv := objc.Send[bool](objc.ID(_CAEDRMetadataClass.class), objc.Sel("isAvailable"))
	return rv
}
