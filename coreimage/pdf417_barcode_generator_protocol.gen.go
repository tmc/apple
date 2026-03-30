// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a PDF417 barcode generator filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator
type CIPDF417BarcodeGenerator interface {
	objectivec.IObject
	CIFilterProtocol

	// A Boolean value specifying whether to force compaction style.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/alwaysSpecifyCompaction
	AlwaysSpecifyCompaction() float32

	// A Boolean value specifying whether to force compact style Aztec code.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/compactStyle
	CompactStyle() float32

	// The compaction mode of the generated barcode.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/compactionMode
	CompactionMode() float32

	// The correction level ratio of the generated barcode.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/correctionLevel
	CorrectionLevel() float32

	// The number of data columns in the generated barcode.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/dataColumns
	DataColumns() float32

	// The maximum height, in pixels, of the generated barcode.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/maxHeight
	MaxHeight() float32

	// The maximum width, in pixels, of the generated barcode.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/maxWidth
	MaxWidth() float32

	// The message to encode in the PDF417 barcode.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/message
	Message() foundation.INSData

	// The minimum height, in pixels, of the generated barcode.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/minHeight
	MinHeight() float32

	// The minimum width, in pixels, of the generated barcode.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/minWidth
	MinWidth() float32

	// The preferred aspect ratio of the generated barcode.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/preferredAspectRatio
	PreferredAspectRatio() float32

	// The number of rows in the generated barcode.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/rows
	Rows() float32

	// A Boolean value specifying whether to force compaction style.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/alwaysSpecifyCompaction
	SetAlwaysSpecifyCompaction(value float32)

	// A Boolean value specifying whether to force compact style Aztec code.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/compactStyle
	SetCompactStyle(value float32)

	// The compaction mode of the generated barcode.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/compactionMode
	SetCompactionMode(value float32)

	// The correction level ratio of the generated barcode.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/correctionLevel
	SetCorrectionLevel(value float32)

	// The number of data columns in the generated barcode.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/dataColumns
	SetDataColumns(value float32)

	// The maximum height, in pixels, of the generated barcode.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/maxHeight
	SetMaxHeight(value float32)

	// The maximum width, in pixels, of the generated barcode.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/maxWidth
	SetMaxWidth(value float32)

	// The message to encode in the PDF417 barcode.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/message
	SetMessage(value foundation.INSData)

	// The minimum height, in pixels, of the generated barcode.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/minHeight
	SetMinHeight(value float32)

	// The minimum width, in pixels, of the generated barcode.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/minWidth
	SetMinWidth(value float32)

	// The preferred aspect ratio of the generated barcode.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/preferredAspectRatio
	SetPreferredAspectRatio(value float32)

	// The number of rows in the generated barcode.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/rows
	SetRows(value float32)
}

// CIPDF417BarcodeGeneratorObject wraps an existing Objective-C object that conforms to the CIPDF417BarcodeGenerator protocol.
type CIPDF417BarcodeGeneratorObject struct {
	objectivec.Object
}

func (o CIPDF417BarcodeGeneratorObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIPDF417BarcodeGeneratorObjectFromID constructs a [CIPDF417BarcodeGeneratorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIPDF417BarcodeGeneratorObjectFromID(id objc.ID) CIPDF417BarcodeGeneratorObject {
	return CIPDF417BarcodeGeneratorObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A Boolean value specifying whether to force compaction style.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/alwaysSpecifyCompaction
func (o CIPDF417BarcodeGeneratorObject) AlwaysSpecifyCompaction() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("alwaysSpecifyCompaction"))
	return rv
}

// A Boolean value specifying whether to force compact style Aztec code.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/compactStyle
func (o CIPDF417BarcodeGeneratorObject) CompactStyle() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("compactStyle"))
	return rv
}

// The compaction mode of the generated barcode.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/compactionMode
func (o CIPDF417BarcodeGeneratorObject) CompactionMode() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("compactionMode"))
	return rv
}

// The correction level ratio of the generated barcode.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/correctionLevel
func (o CIPDF417BarcodeGeneratorObject) CorrectionLevel() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("correctionLevel"))
	return rv
}

// The number of data columns in the generated barcode.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/dataColumns
func (o CIPDF417BarcodeGeneratorObject) DataColumns() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("dataColumns"))
	return rv
}

// The maximum height, in pixels, of the generated barcode.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/maxHeight
func (o CIPDF417BarcodeGeneratorObject) MaxHeight() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("maxHeight"))
	return rv
}

// The maximum width, in pixels, of the generated barcode.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/maxWidth
func (o CIPDF417BarcodeGeneratorObject) MaxWidth() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("maxWidth"))
	return rv
}

// The message to encode in the PDF417 barcode.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/message
func (o CIPDF417BarcodeGeneratorObject) Message() foundation.INSData {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("message"))
	return foundation.NSDataFromID(rv)
}

// The minimum height, in pixels, of the generated barcode.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/minHeight
func (o CIPDF417BarcodeGeneratorObject) MinHeight() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("minHeight"))
	return rv
}

// The minimum width, in pixels, of the generated barcode.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/minWidth
func (o CIPDF417BarcodeGeneratorObject) MinWidth() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("minWidth"))
	return rv
}

// The preferred aspect ratio of the generated barcode.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/preferredAspectRatio
func (o CIPDF417BarcodeGeneratorObject) PreferredAspectRatio() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("preferredAspectRatio"))
	return rv
}

// The number of rows in the generated barcode.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/rows
func (o CIPDF417BarcodeGeneratorObject) Rows() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("rows"))
	return rv
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIPDF417BarcodeGeneratorObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// A Boolean value specifying whether to force compaction style.
//
// # Discussion
//
// Set to nil for automatic.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/alwaysSpecifyCompaction
func (o CIPDF417BarcodeGeneratorObject) SetAlwaysSpecifyCompaction(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAlwaysSpecifyCompaction:"), value)
}

// A Boolean value specifying whether to force compact style Aztec code.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/compactStyle
func (o CIPDF417BarcodeGeneratorObject) SetCompactStyle(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setCompactStyle:"), value)
}

// The compaction mode of the generated barcode.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/compactionMode
func (o CIPDF417BarcodeGeneratorObject) SetCompactionMode(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setCompactionMode:"), value)
}

// The correction level ratio of the generated barcode.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/correctionLevel
func (o CIPDF417BarcodeGeneratorObject) SetCorrectionLevel(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setCorrectionLevel:"), value)
}

// The number of data columns in the generated barcode.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/dataColumns
func (o CIPDF417BarcodeGeneratorObject) SetDataColumns(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setDataColumns:"), value)
}

// The maximum height, in pixels, of the generated barcode.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/maxHeight
func (o CIPDF417BarcodeGeneratorObject) SetMaxHeight(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setMaxHeight:"), value)
}

// The maximum width, in pixels, of the generated barcode.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/maxWidth
func (o CIPDF417BarcodeGeneratorObject) SetMaxWidth(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setMaxWidth:"), value)
}

// The message to encode in the PDF417 barcode.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/message
func (o CIPDF417BarcodeGeneratorObject) SetMessage(value foundation.INSData) {
	objc.Send[struct{}](o.ID, objc.Sel("setMessage:"), value)
}

// The minimum height, in pixels, of the generated barcode.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/minHeight
func (o CIPDF417BarcodeGeneratorObject) SetMinHeight(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setMinHeight:"), value)
}

// The minimum width, in pixels, of the generated barcode.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/minWidth
func (o CIPDF417BarcodeGeneratorObject) SetMinWidth(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setMinWidth:"), value)
}

// The preferred aspect ratio of the generated barcode.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/preferredAspectRatio
func (o CIPDF417BarcodeGeneratorObject) SetPreferredAspectRatio(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setPreferredAspectRatio:"), value)
}

// The number of rows in the generated barcode.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPDF417BarcodeGenerator/rows
func (o CIPDF417BarcodeGeneratorObject) SetRows(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRows:"), value)
}
