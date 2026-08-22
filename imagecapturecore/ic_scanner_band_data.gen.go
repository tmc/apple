// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ICScannerBandData] class.
var (
	_ICScannerBandDataClass     ICScannerBandDataClass
	_ICScannerBandDataClassOnce sync.Once
)

func getICScannerBandDataClass() ICScannerBandDataClass {
	_ICScannerBandDataClassOnce.Do(func() {
		_ICScannerBandDataClass = ICScannerBandDataClass{class: objc.GetClass("ICScannerBandData")}
	})
	return _ICScannerBandDataClass
}

// GetICScannerBandDataClass returns the class object for ICScannerBandData.
func GetICScannerBandDataClass() ICScannerBandDataClass {
	return getICScannerBandDataClass()
}

type ICScannerBandDataClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic ICScannerBandDataClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic ICScannerBandDataClass) Alloc() ICScannerBandData {
	rv := objc.Send[ICScannerBandData](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// The options for each band of data that the scanner reads.
//
// # Instance Properties
//
//   - [ICScannerBandData.BitsPerComponent]
//   - [ICScannerBandData.BitsPerPixel]
//   - [ICScannerBandData.BytesPerRow]
//   - [ICScannerBandData.ColorSyncProfilePath]
//   - [ICScannerBandData.DataBuffer]
//   - [ICScannerBandData.DataNumRows]
//   - [ICScannerBandData.DataSize]
//   - [ICScannerBandData.DataStartRow]
//   - [ICScannerBandData.FullImageHeight]
//   - [ICScannerBandData.FullImageWidth]
//   - [ICScannerBandData.IsBigEndian]
//   - [ICScannerBandData.NumComponents]
//   - [ICScannerBandData.PixelDataType]
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerBandData
type ICScannerBandData struct {
	objectivec.Object
}

// ICScannerBandDataFromID constructs a [ICScannerBandData] from an objc.ID.
//
// The options for each band of data that the scanner reads.
func ICScannerBandDataFromID(id objc.ID) ICScannerBandData {
	return ICScannerBandData{objectivec.Object{ID: id}}
}

// NOTE: ICScannerBandData adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [ICScannerBandData] class.
//
// # Instance Properties
//
//   - [IICScannerBandData.BitsPerComponent]
//   - [IICScannerBandData.BitsPerPixel]
//   - [IICScannerBandData.BytesPerRow]
//   - [IICScannerBandData.ColorSyncProfilePath]
//   - [IICScannerBandData.DataBuffer]
//   - [IICScannerBandData.DataNumRows]
//   - [IICScannerBandData.DataSize]
//   - [IICScannerBandData.DataStartRow]
//   - [IICScannerBandData.FullImageHeight]
//   - [IICScannerBandData.FullImageWidth]
//   - [IICScannerBandData.IsBigEndian]
//   - [IICScannerBandData.NumComponents]
//   - [IICScannerBandData.PixelDataType]
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerBandData
type IICScannerBandData interface {
	objectivec.IObject

	// Topic: Instance Properties

	BitsPerComponent() uint
	BitsPerPixel() uint
	BytesPerRow() uint
	ColorSyncProfilePath() string
	DataBuffer() foundation.NSData
	DataNumRows() uint
	DataSize() uint
	DataStartRow() uint
	FullImageHeight() uint
	FullImageWidth() uint
	IsBigEndian() bool
	NumComponents() uint
	PixelDataType() ICScannerPixelDataType
}

// Init initializes the instance.
func (s ICScannerBandData) Init() ICScannerBandData {
	rv := objc.Send[ICScannerBandData](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s ICScannerBandData) Autorelease() ICScannerBandData {
	rv := objc.Send[ICScannerBandData](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewICScannerBandData creates a new ICScannerBandData instance.
func NewICScannerBandData() ICScannerBandData {
	class := getICScannerBandDataClass()
	rv := objc.Send[ICScannerBandData](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerBandData/bitsPerComponent
func (s ICScannerBandData) BitsPerComponent() uint {
	rv := objc.Send[uint](s.ID, objc.Sel("bitsPerComponent"))
	return rv
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerBandData/bitsPerPixel
func (s ICScannerBandData) BitsPerPixel() uint {
	rv := objc.Send[uint](s.ID, objc.Sel("bitsPerPixel"))
	return rv
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerBandData/bytesPerRow
func (s ICScannerBandData) BytesPerRow() uint {
	rv := objc.Send[uint](s.ID, objc.Sel("bytesPerRow"))
	return rv
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerBandData/colorSyncProfilePath
func (s ICScannerBandData) ColorSyncProfilePath() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("colorSyncProfilePath"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerBandData/dataBuffer
func (s ICScannerBandData) DataBuffer() foundation.NSData {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("dataBuffer"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerBandData/dataNumRows
func (s ICScannerBandData) DataNumRows() uint {
	rv := objc.Send[uint](s.ID, objc.Sel("dataNumRows"))
	return rv
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerBandData/dataSize
func (s ICScannerBandData) DataSize() uint {
	rv := objc.Send[uint](s.ID, objc.Sel("dataSize"))
	return rv
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerBandData/dataStartRow
func (s ICScannerBandData) DataStartRow() uint {
	rv := objc.Send[uint](s.ID, objc.Sel("dataStartRow"))
	return rv
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerBandData/fullImageHeight
func (s ICScannerBandData) FullImageHeight() uint {
	rv := objc.Send[uint](s.ID, objc.Sel("fullImageHeight"))
	return rv
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerBandData/fullImageWidth
func (s ICScannerBandData) FullImageWidth() uint {
	rv := objc.Send[uint](s.ID, objc.Sel("fullImageWidth"))
	return rv
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerBandData/isBigEndian
func (s ICScannerBandData) IsBigEndian() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isBigEndian"))
	return rv
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerBandData/numComponents
func (s ICScannerBandData) NumComponents() uint {
	rv := objc.Send[uint](s.ID, objc.Sel("numComponents"))
	return rv
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerBandData/pixelDataType
func (s ICScannerBandData) PixelDataType() ICScannerPixelDataType {
	rv := objc.Send[ICScannerPixelDataType](s.ID, objc.Sel("pixelDataType"))
	return ICScannerPixelDataType(rv)
}
