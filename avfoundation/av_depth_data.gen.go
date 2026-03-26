// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVDepthData] class.
var (
	_AVDepthDataClass     AVDepthDataClass
	_AVDepthDataClassOnce sync.Once
)

func getAVDepthDataClass() AVDepthDataClass {
	_AVDepthDataClassOnce.Do(func() {
		_AVDepthDataClass = AVDepthDataClass{class: objc.GetClass("AVDepthData")}
	})
	return _AVDepthDataClass
}

// GetAVDepthDataClass returns the class object for AVDepthData.
func GetAVDepthDataClass() AVDepthDataClass {
	return getAVDepthDataClass()
}

type AVDepthDataClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVDepthDataClass) Alloc() AVDepthData {
	rv := objc.Send[AVDepthData](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A container for per-pixel distance or disparity information captured by
// compatible camera devices.
//
// # Overview
// 
// is a generic term for a map of per-pixel data containing depth-related
// information. A depth data object wraps a disparity or depth map and
// provides conversion methods, focus information, and camera calibration data
// to aid in using the map for rendering or computer vision tasks.
// 
// A depth map describes at each pixel the distance to an object, in meters.
// 
// A disparity map describes normalized shift values for use in comparing two
// images. The value for each pixel in the map is in units of 1/meters:
// (`pixelShift / (pixelFocalLength * baselineInMeters)`).
// 
// The capture pipeline generates disparity or depth maps from camera images
// containing nonrectilinear data. Camera lenses have small imperfections that
// cause small distortions in their resultant images compared to an ideal
// pinhole camera model, so [AVDepthData] maps contain nonrectilinear
// (nondistortion-corrected) data as well. The maps’ values are warped to
// match the lens distortion characteristics present in the YUV image pixel
// buffers captured at the same time.
// 
// Because a depth data map is nonrectilinear, you can use an [AVDepthData]
// map as a proxy for depth when rendering effects to its accompanying image,
// but not to correlate points in 3D space. To use depth data for computer
// vision tasks, use the data in the [CameraCalibrationData] property to
// rectify the depth data.
// 
// There are two ways to capture depth data:
// 
// - The [AVCaptureDepthDataOutput] class captures and delivers depth data in
// a stream (similar to how the [AVCaptureVideoDataOutput] delivers video
// data). - The [AVCapturePhotoOutput] class delivers depth data as a property
// of an [AVCapturePhoto] object containing the captured image.
// 
// You can also create [AVDepthData] objects using information obtained from
// image files with the [Image I/O] framework.
// 
// When editing images containing depth information, use the methods listed in
// Transforming and Processing to generate derivative [AVDepthData] objects
// reflecting the edits that have been performed.
//
// [AVCaptureDepthDataOutput]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDepthDataOutput
// [Image I/O]: https://developer.apple.com/documentation/ImageIO
//
// # Creating depth data
//
//   - [AVDepthData.DictionaryRepresentationForAuxiliaryDataType]: Returns a dictionary representation of the depth data suitable for writing into an image file.
//
// # Reading pixel depth information
//
//   - [AVDepthData.DepthDataMap]: A pixel buffer containing the depth data’s per-pixel depth or disparity data map.
//   - [AVDepthData.DepthDataType]: The pixel format of the depth data map.
//
// # Evaluating depth data
//
//   - [AVDepthData.DepthDataFiltered]: A Boolean value indicating whether the depth map contains temporally smoothed data.
//   - [AVDepthData.DepthDataAccuracy]: The general accuracy of depth data map values.
//   - [AVDepthData.DepthDataQuality]: The overall quality of the depth map.
//
// # Transforming and processing
//
//   - [AVDepthData.DepthDataByApplyingExifOrientation]: Returns a derivative depth data object by mirroring or rotating it to the specified orientation.
//   - [AVDepthData.DepthDataByConvertingToDepthDataType]: Returns a derivative depth data object by converting the depth data map to the specified data type.
//   - [AVDepthData.DepthDataByReplacingDepthDataMapWithPixelBufferError]: Returns a derivative depth data object by replacing the depth data map.
//
// # Using calibration data
//
//   - [AVDepthData.CameraCalibrationData]: The imaging parameters with which this depth data was captured.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDepthData
type AVDepthData struct {
	objectivec.Object
}

// AVDepthDataFromID constructs a [AVDepthData] from an objc.ID.
//
// A container for per-pixel distance or disparity information captured by
// compatible camera devices.
func AVDepthDataFromID(id objc.ID) AVDepthData {
	return AVDepthData{objectivec.Object{ID: id}}
}
// NOTE: AVDepthData adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVDepthData] class.
//
// # Creating depth data
//
//   - [IAVDepthData.DictionaryRepresentationForAuxiliaryDataType]: Returns a dictionary representation of the depth data suitable for writing into an image file.
//
// # Reading pixel depth information
//
//   - [IAVDepthData.DepthDataMap]: A pixel buffer containing the depth data’s per-pixel depth or disparity data map.
//   - [IAVDepthData.DepthDataType]: The pixel format of the depth data map.
//
// # Evaluating depth data
//
//   - [IAVDepthData.DepthDataFiltered]: A Boolean value indicating whether the depth map contains temporally smoothed data.
//   - [IAVDepthData.DepthDataAccuracy]: The general accuracy of depth data map values.
//   - [IAVDepthData.DepthDataQuality]: The overall quality of the depth map.
//
// # Transforming and processing
//
//   - [IAVDepthData.DepthDataByApplyingExifOrientation]: Returns a derivative depth data object by mirroring or rotating it to the specified orientation.
//   - [IAVDepthData.DepthDataByConvertingToDepthDataType]: Returns a derivative depth data object by converting the depth data map to the specified data type.
//   - [IAVDepthData.DepthDataByReplacingDepthDataMapWithPixelBufferError]: Returns a derivative depth data object by replacing the depth data map.
//
// # Using calibration data
//
//   - [IAVDepthData.CameraCalibrationData]: The imaging parameters with which this depth data was captured.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDepthData
type IAVDepthData interface {
	objectivec.IObject

	// Topic: Creating depth data

	// Returns a dictionary representation of the depth data suitable for writing into an image file.
	DictionaryRepresentationForAuxiliaryDataType(outAuxDataType string) foundation.INSDictionary

	// Topic: Reading pixel depth information

	// A pixel buffer containing the depth data’s per-pixel depth or disparity data map.
	DepthDataMap() corevideo.CVImageBufferRef
	// The pixel format of the depth data map.
	DepthDataType() uint32

	// Topic: Evaluating depth data

	// A Boolean value indicating whether the depth map contains temporally smoothed data.
	DepthDataFiltered() bool
	// The general accuracy of depth data map values.
	DepthDataAccuracy() AVDepthDataAccuracy
	// The overall quality of the depth map.
	DepthDataQuality() AVDepthDataQuality

	// Topic: Transforming and processing

	// Returns a derivative depth data object by mirroring or rotating it to the specified orientation.
	DepthDataByApplyingExifOrientation(exifOrientation objectivec.IObject) IAVDepthData
	// Returns a derivative depth data object by converting the depth data map to the specified data type.
	DepthDataByConvertingToDepthDataType(depthDataType uint32) IAVDepthData
	// Returns a derivative depth data object by replacing the depth data map.
	DepthDataByReplacingDepthDataMapWithPixelBufferError(pixelBuffer corevideo.CVImageBufferRef) (IAVDepthData, error)

	// Topic: Using calibration data

	// The imaging parameters with which this depth data was captured.
	CameraCalibrationData() IAVCameraCalibrationData

	// The list of depth data formats to which you can convert this depth data.
	AvailableDepthDataTypes() []foundation.NSNumber
}

// Init initializes the instance.
func (d AVDepthData) Init() AVDepthData {
	rv := objc.Send[AVDepthData](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d AVDepthData) Autorelease() AVDepthData {
	rv := objc.Send[AVDepthData](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVDepthData creates a new AVDepthData instance.
func NewAVDepthData() AVDepthData {
	class := getAVDepthDataClass()
	rv := objc.Send[AVDepthData](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a depth data object from depth information such as that found in an
// image file.
//
// imageSourceAuxDataInfoDictionary: A dictionary of primitive depth-related information, in the format provided
// by the [CGImageSourceCopyAuxiliaryDataInfoAtIndex(_:_:_:)] function.
// //
// [CGImageSourceCopyAuxiliaryDataInfoAtIndex(_:_:_:)]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCopyAuxiliaryDataInfoAtIndex(_:_:_:)
//
// # Discussion
// 
// When using [CGImageSource] functions to read from a HEIF, JPEG, or DNG file
// containing depth data (as well as image data), you can use the
// [CGImageSourceCopyAuxiliaryDataInfoAtIndex(_:_:_:)] function to load
// primitive depth map information, then use this initializer to create an
// [AVDepthData] object, as shown below.
//
// [CGImageSourceCopyAuxiliaryDataInfoAtIndex(_:_:_:)]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCopyAuxiliaryDataInfoAtIndex(_:_:_:)
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDepthData/init(fromDictionaryRepresentation:)
func NewDepthDataFromDictionaryRepresentationError(imageSourceAuxDataInfoDictionary foundation.INSDictionary) (AVDepthData, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getAVDepthDataClass().class), objc.Sel("depthDataFromDictionaryRepresentation:error:"), imageSourceAuxDataInfoDictionary, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVDepthData{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVDepthDataFromID(rv), nil
}

// Returns a dictionary representation of the depth data suitable for writing
// into an image file.
//
// outAuxDataType: On output, either [kCGImageAuxiliaryDataTypeDisparity] or
// [kCGImageAuxiliaryDataTypeDepth], depending on the depth data’s type.
// //
// [kCGImageAuxiliaryDataTypeDepth]: https://developer.apple.com/documentation/ImageIO/kCGImageAuxiliaryDataTypeDepth
// [kCGImageAuxiliaryDataTypeDisparity]: https://developer.apple.com/documentation/ImageIO/kCGImageAuxiliaryDataTypeDisparity
//
// # Discussion
// 
// When using [CGImageDestination] functions to write depth data (along with
// image data) to a HEIF, JPEG, or DNG file, you can use this method to obtain
// a dictionary of primitive depth map information, then use the
// [CGImageDestinationAddAuxiliaryDataInfo(_:_:_:)] function to embed that
// data into the output file.
//
// [CGImageDestinationAddAuxiliaryDataInfo(_:_:_:)]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationAddAuxiliaryDataInfo(_:_:_:)
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDepthData/dictionaryRepresentation(forAuxiliaryDataType:)
func (d AVDepthData) DictionaryRepresentationForAuxiliaryDataType(outAuxDataType string) foundation.INSDictionary {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("dictionaryRepresentationForAuxiliaryDataType:"), objc.String(outAuxDataType))
	return foundation.NSDictionaryFromID(rv)
}
// Returns a derivative depth data object by mirroring or rotating it to the
// specified orientation.
//
// exifOrientation: The image orientation to apply to the depth data map.
//
// exifOrientation is a [imageio.CGImagePropertyOrientation].
//
// # Return Value
// 
// A new, transformed depth data object.
//
// # Discussion
// 
// When applying simple 90-degree rotation or mirroring edits to media
// containing depth data, you may use this method to create a derivative copy
// of the depth in which the specified orientation is applied to both the
// underlying pixel map data and the camera calibration data. This method
// throws an exception if you pass an unrecognized `exifOrientation` value.
// 
// A depth data object does not contain orientation metadata; this method
// assumes the data is in the default [CGImagePropertyOrientation.up]
// orientation and applies the transformation necessary to produce the
// orientation you specify.
//
// [CGImagePropertyOrientation.up]: https://developer.apple.com/documentation/ImageIO/CGImagePropertyOrientation/up
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDepthData/applyingExifOrientation(_:)
func (d AVDepthData) DepthDataByApplyingExifOrientation(exifOrientation objectivec.IObject) IAVDepthData {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("depthDataByApplyingExifOrientation:"), exifOrientation)
	return AVDepthDataFromID(rv)
}
// Returns a derivative depth data object by converting the depth data map to
// the specified data type.
//
// depthDataType: The data type to convert to. This value must be one of the formats present
// in the [AvailableDepthDataTypes] array.
//
// # Return Value
// 
// A new, converted depth data object.
//
// # Discussion
// 
// This method raises an exception if you pass an invalid `depthDataType`
// value.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDepthData/converting(toDepthDataType:)
func (d AVDepthData) DepthDataByConvertingToDepthDataType(depthDataType uint32) IAVDepthData {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("depthDataByConvertingToDepthDataType:"), depthDataType)
	return AVDepthDataFromID(rv)
}
// Returns a derivative depth data object by replacing the depth data map.
//
// pixelBuffer: A pixel buffer containing depth or disparity information in a compatible
// format.
//
// # Return Value
// 
// A new depth data object containing the pixel buffer.
//
// # Discussion
// 
// If you apply simple transforms to media containing depth data, you can use
// the [DepthDataByApplyingExifOrientation] method to apply parallel
// transforms to the corresponding depth data. More complex transforms and
// edits require creating a derivative depth map reflecting whatever edits you
// make to the corresponding image. In such cases, use this
// [DepthDataByReplacingDepthDataMapWithPixelBufferError] method to create a
// derivative depth data object.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDepthData/replacingDepthDataMap(with:)
func (d AVDepthData) DepthDataByReplacingDepthDataMapWithPixelBufferError(pixelBuffer corevideo.CVImageBufferRef) (IAVDepthData, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("depthDataByReplacingDepthDataMapWithPixelBuffer:error:"), pixelBuffer, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVDepthData{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVDepthDataFromID(rv), nil

}

// A pixel buffer containing the depth data’s per-pixel depth or disparity
// data map.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDepthData/depthDataMap
func (d AVDepthData) DepthDataMap() corevideo.CVImageBufferRef {
	rv := objc.Send[corevideo.CVImageBufferRef](d.ID, objc.Sel("depthDataMap"))
	return corevideo.CVImageBufferRef(rv)
}
// The pixel format of the depth data map.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDepthData/depthDataType
func (d AVDepthData) DepthDataType() uint32 {
	rv := objc.Send[uint32](d.ID, objc.Sel("depthDataType"))
	return rv
}
// A Boolean value indicating whether the depth map contains temporally
// smoothed data.
//
// # Discussion
// 
// The capture system can smooth noise and fill in missing values (caused by
// low light or lens occlusion) in depth data maps by temporally interpolating
// between previous and subsequent frames of captured depth data. Use the
// [AVCaptureDepthDataOutput] [isFilteringEnabled] property to control
// filtering for streaming depth capture, or the [AVCapturePhotoSettings]
// [DepthDataFiltered] property to control filtering for depth data captured
// alongside photo capture.
// 
// Filtering depth data makes it more useful for applying visual effects to a
// companion image, but alters the data such that it may no longer be suitable
// for computer vision tasks. (In an unfiltered depth map, missing values are
// represented as [NaN].)
//
// [AVCaptureDepthDataOutput]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDepthDataOutput
// [isFilteringEnabled]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDepthDataOutput/isFilteringEnabled
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDepthData/isDepthDataFiltered
func (d AVDepthData) DepthDataFiltered() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isDepthDataFiltered"))
	return rv
}
// The general accuracy of depth data map values.
//
// # Discussion
// 
// The accuracy of a depth data map is highly dependent on the camera
// calibration data used to generate it. If the camera’s focal length cannot
// be precisely determined at the time of capture, a scaling error in the z
// (depth) plane is introduced. If the camera’s optical center can’t be
// precisely determined at capture time, a principal point error is
// introduced, leading to an offset error in the disparity estimate.
// [AVDepthData.Accuracy] constants report the accuracy of a map’s values
// with respect to its reported units.
//
// [AVDepthData.Accuracy]: https://developer.apple.com/documentation/AVFoundation/AVDepthData/Accuracy
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDepthData/depthDataAccuracy
func (d AVDepthData) DepthDataAccuracy() AVDepthDataAccuracy {
	rv := objc.Send[AVDepthDataAccuracy](d.ID, objc.Sel("depthDataAccuracy"))
	return AVDepthDataAccuracy(rv)
}
// The overall quality of the depth map.
//
// # Discussion
// 
// A device typically generates depth data maps by comparing images and
// calculating disparity. If features are lacking in either input image, it
// may be difficult to find matching key points, resulting in a depth data map
// with substantial holes. These holes can be filled with depth data
// filtering, but still may produce a map of overall poor quality.
// 
// If a depth data map suffers from insufficient features, the capture system
// marks it as [DepthDataQualityLow] quality, indicating that the depth map is
// a poor candidate for rendering high-quality depth effects or reconstructing
// a 3D scene. A depth map with [DepthDataQualityHigh] quality is
// feature-rich, contains a high level of detail, making it a good candidate
// for rendering high-quality depth effects or reconstructing a 3D scene.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDepthData/depthDataQuality
func (d AVDepthData) DepthDataQuality() AVDepthDataQuality {
	rv := objc.Send[AVDepthDataQuality](d.ID, objc.Sel("depthDataQuality"))
	return AVDepthDataQuality(rv)
}
// The imaging parameters with which this depth data was captured.
//
// # Discussion
// 
// Using depth or disparity map data to render effects into a corresponding
// image or to perform computer vision tasks requires knowledge of the camera
// parameters that generated the depth data. Depth data captured by an
// [AVCaptureDevice] object contains camera calibration data that includes
// such information.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDepthData/cameraCalibrationData
func (d AVDepthData) CameraCalibrationData() IAVCameraCalibrationData {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("cameraCalibrationData"))
	return AVCameraCalibrationDataFromID(objc.ID(rv))
}
// The list of depth data formats to which you can convert this depth data.
//
// # Discussion
// 
// Use the [DepthDataByConvertingToDepthDataType] method to obtain a converted
// depth data object using one of the types in this list.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDepthData/availableDepthDataTypes-472g0
func (d AVDepthData) AvailableDepthDataTypes() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("availableDepthDataTypes"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

