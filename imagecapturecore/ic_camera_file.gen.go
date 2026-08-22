// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [ICCameraFile] class.
var (
	_ICCameraFileClass     ICCameraFileClass
	_ICCameraFileClassOnce sync.Once
)

func getICCameraFileClass() ICCameraFileClass {
	_ICCameraFileClassOnce.Do(func() {
		_ICCameraFileClass = ICCameraFileClass{class: objc.GetClass("ICCameraFile")}
	})
	return _ICCameraFileClass
}

// GetICCameraFileClass returns the class object for ICCameraFile.
func GetICCameraFileClass() ICCameraFileClass {
	return getICCameraFileClass()
}

type ICCameraFileClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic ICCameraFileClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic ICCameraFileClass) Alloc() ICCameraFile {
	rv := objc.Send[ICCameraFile](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a file on a camera.
//
// # Requesting Thumbnails
//
//   - [ICCameraFile.RequestThumbnailDataWithOptionsCompletion]: Requests a thumbnail and executes the completion block in place of the delegate.
//
// # Requesting Downloads
//
//   - [ICCameraFile.RequestDownloadWithOptionsCompletion]: Requests a download and executes the completion block in place of the delegate.
//
// # Requesting Data
//
//   - [ICCameraFile.RequestReadDataAtOffsetLengthCompletion]: Requests to asynchronously read data of a specified length from a specified offset, then executes the completion block.
//
// # Inspecting a File’s Name
//
//   - [ICCameraFile.OriginalFilename]: The original name of the file on disk.
//   - [ICCameraFile.CreatedFilename]: The created name of the file.
//
// # Inspecting a File’s Identity
//
//   - [ICCameraFile.GroupUUID]: The group [UUID] of the file.
//   - [ICCameraFile.RelatedUUID]: A related UUID correlating several images from an Apple device.
//   - [ICCameraFile.OriginatingAssetID]: The originating asset ID of an [HEIF] or [HVEC] file.
//
// # Determining When a File Was Created or Modified
//
//   - [ICCameraFile.FileCreationDate]: The creation date of the file.
//   - [ICCameraFile.FileModificationDate]: The modification date of the file.
//
// # Inspecting a File’s Size
//
//   - [ICCameraFile.FileSize]: The size of the file, in bytes.
//
// # Inspecting a File’s Dimensions
//
//   - [ICCameraFile.Width]: The width of an image or movie frame.
//   - [ICCameraFile.Height]: The height of an image or movie frame.
//
// # Inspecting a File’s EXIF Data
//
//   - [ICCameraFile.Orientation]: The orientation to use when downloading the image.
//   - [ICCameraFile.SetOrientation]
//   - [ICCameraFile.ExifCreationDate]: The [EXIF] creation date of the file.
//   - [ICCameraFile.ExifModificationDate]: The [EXIF] modification date of the file.
//
// # Identifying a File’s Location
//
//   - [ICCameraFile.GpsString]: The GPS String of the file in standard format.
//
// # Inspecting a File in a Burst
//
//   - [ICCameraFile.FirstPicked]: A Boolean value that indicates whether a file is autopicked by Photos to represent the burst.
//   - [ICCameraFile.BurstUUID]: The burst UUID of the file if it is in a burst.
//   - [ICCameraFile.BurstFavorite]: A Boolean value that indicates this file is the burst favorite in a burst.
//   - [ICCameraFile.BurstPicked]: A Boolean value that indicates whether this file is user picked in a burst.
//
// # Inspecting Video Properties
//
//   - [ICCameraFile.Duration]: The duration, in seconds, of an audio or video file.
//   - [ICCameraFile.HighFramerate]: A Boolean value that indicates whether the file is a slow motion or high-frame-rate video file.
//   - [ICCameraFile.TimeLapse]: A Boolean value that indicates whether the file is a time-lapse video file.
//
// # Identifying Related Files
//
//   - [ICCameraFile.SidecarFiles]: An array of two camera files associated with this file.
//   - [ICCameraFile.PairedRawImage]: A sidecar file containing the logical [RAW] compliment of a [JPG] or other two-format image.
//
// # Instance Properties
//
//   - [ICCameraFile.Fingerprint]
//
// # Instance Methods
//
//   - [ICCameraFile.RequestFingerprintWithCompletion]
//   - [ICCameraFile.RequestSecurityScopedURLWithCompletion]
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile
type ICCameraFile struct {
	ICCameraItem
}

// ICCameraFileFromID constructs a [ICCameraFile] from an objc.ID.
//
// An object that represents a file on a camera.
func ICCameraFileFromID(id objc.ID) ICCameraFile {
	return ICCameraFile{ICCameraItem: ICCameraItemFromID(id)}
}

// NOTE: ICCameraFile adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [ICCameraFile] class.
//
// # Requesting Thumbnails
//
//   - [IICCameraFile.RequestThumbnailDataWithOptionsCompletion]: Requests a thumbnail and executes the completion block in place of the delegate.
//
// # Requesting Downloads
//
//   - [IICCameraFile.RequestDownloadWithOptionsCompletion]: Requests a download and executes the completion block in place of the delegate.
//
// # Requesting Data
//
//   - [IICCameraFile.RequestReadDataAtOffsetLengthCompletion]: Requests to asynchronously read data of a specified length from a specified offset, then executes the completion block.
//
// # Inspecting a File’s Name
//
//   - [IICCameraFile.OriginalFilename]: The original name of the file on disk.
//   - [IICCameraFile.CreatedFilename]: The created name of the file.
//
// # Inspecting a File’s Identity
//
//   - [IICCameraFile.GroupUUID]: The group [UUID] of the file.
//   - [IICCameraFile.RelatedUUID]: A related UUID correlating several images from an Apple device.
//   - [IICCameraFile.OriginatingAssetID]: The originating asset ID of an [HEIF] or [HVEC] file.
//
// # Determining When a File Was Created or Modified
//
//   - [IICCameraFile.FileCreationDate]: The creation date of the file.
//   - [IICCameraFile.FileModificationDate]: The modification date of the file.
//
// # Inspecting a File’s Size
//
//   - [IICCameraFile.FileSize]: The size of the file, in bytes.
//
// # Inspecting a File’s Dimensions
//
//   - [IICCameraFile.Width]: The width of an image or movie frame.
//   - [IICCameraFile.Height]: The height of an image or movie frame.
//
// # Inspecting a File’s EXIF Data
//
//   - [IICCameraFile.Orientation]: The orientation to use when downloading the image.
//   - [IICCameraFile.SetOrientation]
//   - [IICCameraFile.ExifCreationDate]: The [EXIF] creation date of the file.
//   - [IICCameraFile.ExifModificationDate]: The [EXIF] modification date of the file.
//
// # Identifying a File’s Location
//
//   - [IICCameraFile.GpsString]: The GPS String of the file in standard format.
//
// # Inspecting a File in a Burst
//
//   - [IICCameraFile.FirstPicked]: A Boolean value that indicates whether a file is autopicked by Photos to represent the burst.
//   - [IICCameraFile.BurstUUID]: The burst UUID of the file if it is in a burst.
//   - [IICCameraFile.BurstFavorite]: A Boolean value that indicates this file is the burst favorite in a burst.
//   - [IICCameraFile.BurstPicked]: A Boolean value that indicates whether this file is user picked in a burst.
//
// # Inspecting Video Properties
//
//   - [IICCameraFile.Duration]: The duration, in seconds, of an audio or video file.
//   - [IICCameraFile.HighFramerate]: A Boolean value that indicates whether the file is a slow motion or high-frame-rate video file.
//   - [IICCameraFile.TimeLapse]: A Boolean value that indicates whether the file is a time-lapse video file.
//
// # Identifying Related Files
//
//   - [IICCameraFile.SidecarFiles]: An array of two camera files associated with this file.
//   - [IICCameraFile.PairedRawImage]: A sidecar file containing the logical [RAW] compliment of a [JPG] or other two-format image.
//
// # Instance Properties
//
//   - [IICCameraFile.Fingerprint]
//
// # Instance Methods
//
//   - [IICCameraFile.RequestFingerprintWithCompletion]
//   - [IICCameraFile.RequestSecurityScopedURLWithCompletion]
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile
type IICCameraFile interface {
	IICCameraItem

	// Topic: Requesting Thumbnails

	// Requests a thumbnail and executes the completion block in place of the delegate.
	RequestThumbnailDataWithOptionsCompletion(options foundation.INSDictionary, completion DataErrorHandler)

	// Topic: Requesting Downloads

	// Requests a download and executes the completion block in place of the delegate.
	RequestDownloadWithOptionsCompletion(options foundation.INSDictionary, completion StringErrorHandler) foundation.Progress

	// Topic: Requesting Data

	// Requests to asynchronously read data of a specified length from a specified offset, then executes the completion block.
	RequestReadDataAtOffsetLengthCompletion(offset int64, length int64, completion DataErrorHandler)

	// Topic: Inspecting a File’s Name

	// The original name of the file on disk.
	OriginalFilename() string
	// The created name of the file.
	CreatedFilename() string

	// Topic: Inspecting a File’s Identity

	// The group [UUID] of the file.
	GroupUUID() string
	// A related UUID correlating several images from an Apple device.
	RelatedUUID() string
	// The originating asset ID of an [HEIF] or [HVEC] file.
	OriginatingAssetID() string

	// Topic: Determining When a File Was Created or Modified

	// The creation date of the file.
	FileCreationDate() foundation.NSDate
	// The modification date of the file.
	FileModificationDate() foundation.NSDate

	// Topic: Inspecting a File’s Size

	// The size of the file, in bytes.
	FileSize() int64

	// Topic: Inspecting a File’s Dimensions

	// The width of an image or movie frame.
	Width() int
	// The height of an image or movie frame.
	Height() int

	// Topic: Inspecting a File’s EXIF Data

	// The orientation to use when downloading the image.
	Orientation() ICEXIFOrientationType
	SetOrientation(value ICEXIFOrientationType)
	// The [EXIF] creation date of the file.
	ExifCreationDate() foundation.NSDate
	// The [EXIF] modification date of the file.
	ExifModificationDate() foundation.NSDate

	// Topic: Identifying a File’s Location

	// The GPS String of the file in standard format.
	GpsString() string

	// Topic: Inspecting a File in a Burst

	// A Boolean value that indicates whether a file is autopicked by Photos to represent the burst.
	FirstPicked() bool
	// The burst UUID of the file if it is in a burst.
	BurstUUID() string
	// A Boolean value that indicates this file is the burst favorite in a burst.
	BurstFavorite() bool
	// A Boolean value that indicates whether this file is user picked in a burst.
	BurstPicked() bool

	// Topic: Inspecting Video Properties

	// The duration, in seconds, of an audio or video file.
	Duration() float64
	// A Boolean value that indicates whether the file is a slow motion or high-frame-rate video file.
	HighFramerate() bool
	// A Boolean value that indicates whether the file is a time-lapse video file.
	TimeLapse() bool

	// Topic: Identifying Related Files

	// An array of two camera files associated with this file.
	SidecarFiles() []ICCameraItem
	// A sidecar file containing the logical [RAW] compliment of a [JPG] or other two-format image.
	PairedRawImage() IICCameraFile

	// Topic: Instance Properties

	Fingerprint() string

	// Topic: Instance Methods

	RequestFingerprintWithCompletion(completion StringErrorHandler)
	RequestSecurityScopedURLWithCompletion(completion URLErrorHandler)
}

// Init initializes the instance.
func (c ICCameraFile) Init() ICCameraFile {
	rv := objc.Send[ICCameraFile](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c ICCameraFile) Autorelease() ICCameraFile {
	rv := objc.Send[ICCameraFile](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewICCameraFile creates a new ICCameraFile instance.
func NewICCameraFile() ICCameraFile {
	class := getICCameraFileClass()
	rv := objc.Send[ICCameraFile](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Requests a thumbnail and executes the completion block in place of the
// delegate.
//
// # Discussion
//
// The completion block executes on an any available queue; often this is not
// the main queue.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/requestThumbnailData(options:completion:)
func (c ICCameraFile) RequestThumbnailDataWithOptionsCompletion(options foundation.INSDictionary, completion DataErrorHandler) {
	_block1, _ := NewDataErrorBlock(completion)
	objc.Send[objc.ID](c.ID, objc.Sel("requestThumbnailDataWithOptions:completion:"), options, _block1)
}

// Requests a download and executes the completion block in place of the
// delegate.
//
// # Discussion
//
// The completion block executes on an any available queue; often this is not
// the main queue.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/requestDownload(options:completion:)
func (c ICCameraFile) RequestDownloadWithOptionsCompletion(options foundation.INSDictionary, completion StringErrorHandler) foundation.Progress {
	_block1, _ := NewStringErrorBlock(completion)
	rv := objc.Send[objc.ID](c.ID, objc.Sel("requestDownloadWithOptions:completion:"), options, _block1)
	return foundation.NSProgressFromID(rv)
}

// Requests to asynchronously read data of a specified length from a specified
// offset, then executes the completion block.
//
// # Discussion
//
// The completion block executes on an any available queue; often this is not
// the main queue.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/requestReadData(atOffset:length:completion:)
func (c ICCameraFile) RequestReadDataAtOffsetLengthCompletion(offset int64, length int64, completion DataErrorHandler) {
	_block2, _ := NewDataErrorBlock(completion)
	objc.Send[objc.ID](c.ID, objc.Sel("requestReadDataAtOffset:length:completion:"), offset, length, _block2)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/requestFingerprint(completion:)
func (c ICCameraFile) RequestFingerprintWithCompletion(completion StringErrorHandler) {
	_block0, _ := NewStringErrorBlock(completion)
	objc.Send[objc.ID](c.ID, objc.Sel("requestFingerprintWithCompletion:"), _block0)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/requestSecurityScopedURL(completion:)
func (c ICCameraFile) RequestSecurityScopedURLWithCompletion(completion URLErrorHandler) {
	_block0, _ := NewURLErrorBlock(completion)
	objc.Send[objc.ID](c.ID, objc.Sel("requestSecurityScopedURLWithCompletion:"), _block0)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/fingerprintForFile(at:)
func (_ICCameraFileClass ICCameraFileClass) FingerprintForFileAtURL(url foundation.NSURL) string {
	rv := objc.Send[objc.ID](objc.ID(_ICCameraFileClass.class), objc.Sel("fingerprintForFileAtURL:"), url)
	return foundation.NSStringFromID(rv).String()
}

// The original name of the file on disk.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/originalFilename
func (c ICCameraFile) OriginalFilename() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("originalFilename"))
	return foundation.NSStringFromID(rv).String()
}

// The created name of the file.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/createdFilename
func (c ICCameraFile) CreatedFilename() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("createdFilename"))
	return foundation.NSStringFromID(rv).String()
}

// The group [UUID] of the file.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/groupUUID
func (c ICCameraFile) GroupUUID() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("groupUUID"))
	return foundation.NSStringFromID(rv).String()
}

// A related UUID correlating several images from an Apple device.
//
// # Discussion
//
// This value is the same for both the image and video of a LivePhoto.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/relatedUUID
func (c ICCameraFile) RelatedUUID() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("relatedUUID"))
	return foundation.NSStringFromID(rv).String()
}

// The originating asset ID of an [HEIF] or [HVEC] file.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/originatingAssetID
func (c ICCameraFile) OriginatingAssetID() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("originatingAssetID"))
	return foundation.NSStringFromID(rv).String()
}

// The creation date of the file.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/fileCreationDate
func (c ICCameraFile) FileCreationDate() foundation.NSDate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("fileCreationDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}

// The modification date of the file.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/fileModificationDate
func (c ICCameraFile) FileModificationDate() foundation.NSDate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("fileModificationDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}

// The size of the file, in bytes.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/fileSize
func (c ICCameraFile) FileSize() int64 {
	rv := objc.Send[int64](c.ID, objc.Sel("fileSize"))
	return rv
}

// The width of an image or movie frame.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/width
func (c ICCameraFile) Width() int {
	rv := objc.Send[int](c.ID, objc.Sel("width"))
	return rv
}

// The height of an image or movie frame.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/height
func (c ICCameraFile) Height() int {
	rv := objc.Send[int](c.ID, objc.Sel("height"))
	return rv
}

// The orientation to use when downloading the image.
//
// # Discussion
//
// This property is initially set to [ICEXIFOrientationType.orientation1] If
// the format of the file supports the [EXIF] orientation tag, then this
// property updates to match the value of that tag on receipt of the thumbnail
// or metadata for this file.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/orientation
//
// [ICEXIFOrientationType.orientation1]: https://developer.apple.com/documentation/ImageCaptureCore/ICEXIFOrientationType/orientation1
func (c ICCameraFile) Orientation() ICEXIFOrientationType {
	rv := objc.Send[ICEXIFOrientationType](c.ID, objc.Sel("orientation"))
	return ICEXIFOrientationType(rv)
}
func (c ICCameraFile) SetOrientation(value ICEXIFOrientationType) {
	objc.Send[struct{}](c.ID, objc.Sel("setOrientation:"), value)
}

// The [EXIF] creation date of the file.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/exifCreationDate
func (c ICCameraFile) ExifCreationDate() foundation.NSDate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("exifCreationDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}

// The [EXIF] modification date of the file.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/exifModificationDate
func (c ICCameraFile) ExifModificationDate() foundation.NSDate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("exifModificationDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}

// The GPS String of the file in standard format.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/gpsString
func (c ICCameraFile) GpsString() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("gpsString"))
	return foundation.NSStringFromID(rv).String()
}

// A Boolean value that indicates whether a file is autopicked by Photos to
// represent the burst.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/firstPicked
func (c ICCameraFile) FirstPicked() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("firstPicked"))
	return rv
}

// The burst UUID of the file if it is in a burst.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/burstUUID
func (c ICCameraFile) BurstUUID() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("burstUUID"))
	return foundation.NSStringFromID(rv).String()
}

// A Boolean value that indicates this file is the burst favorite in a burst.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/burstFavorite
func (c ICCameraFile) BurstFavorite() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("burstFavorite"))
	return rv
}

// A Boolean value that indicates whether this file is user picked in a burst.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/burstPicked
func (c ICCameraFile) BurstPicked() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("burstPicked"))
	return rv
}

// The duration, in seconds, of an audio or video file.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/duration
func (c ICCameraFile) Duration() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("duration"))
	return rv
}

// A Boolean value that indicates whether the file is a slow motion or
// high-frame-rate video file.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/highFramerate
func (c ICCameraFile) HighFramerate() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("highFramerate"))
	return rv
}

// A Boolean value that indicates whether the file is a time-lapse video file.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/timeLapse
func (c ICCameraFile) TimeLapse() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("timeLapse"))
	return rv
}

// An array of two camera files associated with this file.
//
// # Discussion
//
// An example of a sidecar file is a file with the same base 3 name as this
// file and an [XMP] extension.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/sidecarFiles
func (c ICCameraFile) SidecarFiles() []ICCameraItem {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("sidecarFiles"))
	return objc.ConvertSlice(rv, func(id objc.ID) ICCameraItem {
		return ICCameraItemFromID(id)
	})
}

// A sidecar file containing the logical [RAW] compliment of a [JPG] or other
// two-format image.
//
// # Discussion
//
// This value contains a single-item subset of the [ICCameraFile.SidecarFiles]
// array.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/pairedRawImage
func (c ICCameraFile) PairedRawImage() IICCameraFile {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("pairedRawImage"))
	return ICCameraFileFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFile/fingerprint
func (c ICCameraFile) Fingerprint() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("fingerprint"))
	return foundation.NSStringFromID(rv).String()
}

// RequestThumbnailDataWithOptionsCompletionSync is a synchronous wrapper around [ICCameraFile.RequestThumbnailDataWithOptionsCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (c ICCameraFile) RequestThumbnailDataWithOptionsCompletionSync(ctx context.Context, options foundation.INSDictionary) (*foundation.NSData, error) {
	type result struct {
		val *foundation.NSData
		err error
	}
	done := make(chan result, 1)
	c.RequestThumbnailDataWithOptionsCompletion(options, func(val *foundation.NSData, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// RequestDownloadWithOptionsCompletionSync is a synchronous wrapper around [ICCameraFile.RequestDownloadWithOptionsCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (c ICCameraFile) RequestDownloadWithOptionsCompletionSync(ctx context.Context, options foundation.INSDictionary) (*string, error) {
	type result struct {
		val *string
		err error
	}
	done := make(chan result, 1)
	c.RequestDownloadWithOptionsCompletion(options, func(val *string, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// RequestReadDataAtOffsetLengthCompletionSync is a synchronous wrapper around [ICCameraFile.RequestReadDataAtOffsetLengthCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (c ICCameraFile) RequestReadDataAtOffsetLengthCompletionSync(ctx context.Context, offset int64, length int64) (*foundation.NSData, error) {
	type result struct {
		val *foundation.NSData
		err error
	}
	done := make(chan result, 1)
	c.RequestReadDataAtOffsetLengthCompletion(offset, length, func(val *foundation.NSData, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// RequestFingerprint is a synchronous wrapper around [ICCameraFile.RequestFingerprintWithCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (c ICCameraFile) RequestFingerprint(ctx context.Context) (*string, error) {
	type result struct {
		val *string
		err error
	}
	done := make(chan result, 1)
	c.RequestFingerprintWithCompletion(func(val *string, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// RequestSecurityScopedURL is a synchronous wrapper around [ICCameraFile.RequestSecurityScopedURLWithCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (c ICCameraFile) RequestSecurityScopedURL(ctx context.Context) (*foundation.NSURL, error) {
	type result struct {
		val *foundation.NSURL
		err error
	}
	done := make(chan result, 1)
	c.RequestSecurityScopedURLWithCompletion(func(val *foundation.NSURL, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
