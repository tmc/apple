// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SCScreenshotOutput] class.
var (
	_SCScreenshotOutputClass     SCScreenshotOutputClass
	_SCScreenshotOutputClassOnce sync.Once
)

func getSCScreenshotOutputClass() SCScreenshotOutputClass {
	_SCScreenshotOutputClassOnce.Do(func() {
		_SCScreenshotOutputClass = SCScreenshotOutputClass{class: objc.GetClass("SCScreenshotOutput")}
	})
	return _SCScreenshotOutputClass
}

// GetSCScreenshotOutputClass returns the class object for SCScreenshotOutput.
func GetSCScreenshotOutputClass() SCScreenshotOutputClass {
	return getSCScreenshotOutputClass()
}

type SCScreenshotOutputClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SCScreenshotOutputClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SCScreenshotOutputClass) Alloc() SCScreenshotOutput {
	rv := objc.Send[SCScreenshotOutput](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// An object that contains all images requested by the client.
//
// # Instance Properties
//
//   - [SCScreenshotOutput.FileURL]: A URL property that specifies the location of the saved image.
//   - [SCScreenshotOutput.SetFileURL]
//   - [SCScreenshotOutput.HdrImage]: An output property that specifies the high dynamic range version of the screenshot.
//   - [SCScreenshotOutput.SetHdrImage]
//   - [SCScreenshotOutput.SdrImage]: An output property that specifies the standard dynamic range version of the screenshot.
//   - [SCScreenshotOutput.SetSdrImage]
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotOutput
type SCScreenshotOutput struct {
	objectivec.Object
}

// SCScreenshotOutputFromID constructs a [SCScreenshotOutput] from an objc.ID.
//
// An object that contains all images requested by the client.
func SCScreenshotOutputFromID(id objc.ID) SCScreenshotOutput {
	return SCScreenshotOutput{objectivec.Object{ID: id}}
}
// NOTE: SCScreenshotOutput adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SCScreenshotOutput] class.
//
// # Instance Properties
//
//   - [ISCScreenshotOutput.FileURL]: A URL property that specifies the location of the saved image.
//   - [ISCScreenshotOutput.SetFileURL]
//   - [ISCScreenshotOutput.HdrImage]: An output property that specifies the high dynamic range version of the screenshot.
//   - [ISCScreenshotOutput.SetHdrImage]
//   - [ISCScreenshotOutput.SdrImage]: An output property that specifies the standard dynamic range version of the screenshot.
//   - [ISCScreenshotOutput.SetSdrImage]
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotOutput
type ISCScreenshotOutput interface {
	objectivec.IObject

	// Topic: Instance Properties

	// A URL property that specifies the location of the saved image.
	FileURL() foundation.INSURL
	SetFileURL(value foundation.INSURL)
	// An output property that specifies the high dynamic range version of the screenshot.
	HdrImage() coregraphics.CGImageRef
	SetHdrImage(value coregraphics.CGImageRef)
	// An output property that specifies the standard dynamic range version of the screenshot.
	SdrImage() coregraphics.CGImageRef
	SetSdrImage(value coregraphics.CGImageRef)
}

// Init initializes the instance.
func (s SCScreenshotOutput) Init() SCScreenshotOutput {
	rv := objc.Send[SCScreenshotOutput](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SCScreenshotOutput) Autorelease() SCScreenshotOutput {
	rv := objc.Send[SCScreenshotOutput](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSCScreenshotOutput creates a new SCScreenshotOutput instance.
func NewSCScreenshotOutput() SCScreenshotOutput {
	class := getSCScreenshotOutputClass()
	rv := objc.Send[SCScreenshotOutput](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A URL property that specifies the location of the saved image.
//
// # Discussion
// 
// If `fileURL` is `nil`, then the file isn’t saved.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotOutput/fileURL
func (s SCScreenshotOutput) FileURL() foundation.INSURL {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("fileURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (s SCScreenshotOutput) SetFileURL(value foundation.INSURL) {
	objc.Send[struct{}](s.ID, objc.Sel("setFileURL:"), value)
}
// An output property that specifies the high dynamic range version of the
// screenshot.
//
// # Discussion
// 
// The output [CGImage] uses the extended sRGB color space.
//
// [CGImage]: https://developer.apple.com/documentation/CoreGraphics/CGImage
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotOutput/hdrImage
func (s SCScreenshotOutput) HdrImage() coregraphics.CGImageRef {
	rv := objc.Send[coregraphics.CGImageRef](s.ID, objc.Sel("hdrImage"))
	return coregraphics.CGImageRef(rv)
}
func (s SCScreenshotOutput) SetHdrImage(value coregraphics.CGImageRef) {
	objc.Send[struct{}](s.ID, objc.Sel("setHdrImage:"), value)
}
// An output property that specifies the standard dynamic range version of the
// screenshot.
//
// # Discussion
// 
// The output [CGImage] uses the same color space as the content capture
// display.
//
// [CGImage]: https://developer.apple.com/documentation/CoreGraphics/CGImage
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotOutput/sdrImage
func (s SCScreenshotOutput) SdrImage() coregraphics.CGImageRef {
	rv := objc.Send[coregraphics.CGImageRef](s.ID, objc.Sel("sdrImage"))
	return coregraphics.CGImageRef(rv)
}
func (s SCScreenshotOutput) SetSdrImage(value coregraphics.CGImageRef) {
	objc.Send[struct{}](s.ID, objc.Sel("setSdrImage:"), value)
}

