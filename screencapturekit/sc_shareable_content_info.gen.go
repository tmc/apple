// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SCShareableContentInfo] class.
var (
	_SCShareableContentInfoClass     SCShareableContentInfoClass
	_SCShareableContentInfoClassOnce sync.Once
)

func getSCShareableContentInfoClass() SCShareableContentInfoClass {
	_SCShareableContentInfoClassOnce.Do(func() {
		_SCShareableContentInfoClass = SCShareableContentInfoClass{class: objc.GetClass("SCShareableContentInfo")}
	})
	return _SCShareableContentInfoClass
}

// GetSCShareableContentInfoClass returns the class object for SCShareableContentInfo.
func GetSCShareableContentInfoClass() SCShareableContentInfoClass {
	return getSCShareableContentInfoClass()
}

type SCShareableContentInfoClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (sc SCShareableContentInfoClass) Alloc() SCShareableContentInfo {
	rv := objc.Send[SCShareableContentInfo](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// An instance that provides information for the content in a given stream.
//
// # Shared content properties
//
//   - [SCShareableContentInfo.ContentRect]: The size and location of content for the stream.
//   - [SCShareableContentInfo.PointPixelScale]: The scaling from points to output pixel resolution for the stream.
//   - [SCShareableContentInfo.Style]: The current presentation style of the stream.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContentInfo
type SCShareableContentInfo struct {
	objectivec.Object
}

// SCShareableContentInfoFromID constructs a [SCShareableContentInfo] from an objc.ID.
//
// An instance that provides information for the content in a given stream.
func SCShareableContentInfoFromID(id objc.ID) SCShareableContentInfo {
	return SCShareableContentInfo{objectivec.Object{ID: id}}
}
// NOTE: SCShareableContentInfo adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SCShareableContentInfo] class.
//
// # Shared content properties
//
//   - [ISCShareableContentInfo.ContentRect]: The size and location of content for the stream.
//   - [ISCShareableContentInfo.PointPixelScale]: The scaling from points to output pixel resolution for the stream.
//   - [ISCShareableContentInfo.Style]: The current presentation style of the stream.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContentInfo
type ISCShareableContentInfo interface {
	objectivec.IObject

	// Topic: Shared content properties

	// The size and location of content for the stream.
	ContentRect() corefoundation.CGRect
	// The scaling from points to output pixel resolution for the stream.
	PointPixelScale() float32
	// The current presentation style of the stream.
	Style() SCShareableContentStyle
}

// Init initializes the instance.
func (s SCShareableContentInfo) Init() SCShareableContentInfo {
	rv := objc.Send[SCShareableContentInfo](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SCShareableContentInfo) Autorelease() SCShareableContentInfo {
	rv := objc.Send[SCShareableContentInfo](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSCShareableContentInfo creates a new SCShareableContentInfo instance.
func NewSCShareableContentInfo() SCShareableContentInfo {
	class := getSCShareableContentInfoClass()
	rv := objc.Send[SCShareableContentInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The size and location of content for the stream.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContentInfo/contentRect
func (s SCShareableContentInfo) ContentRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("contentRect"))
	return corefoundation.CGRect(rv)
}

// The scaling from points to output pixel resolution for the stream.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContentInfo/pointPixelScale
func (s SCShareableContentInfo) PointPixelScale() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("pointPixelScale"))
	return rv
}

// The current presentation style of the stream.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContentInfo/style
func (s SCShareableContentInfo) Style() SCShareableContentStyle {
	rv := objc.Send[SCShareableContentStyle](s.ID, objc.Sel("style"))
	return SCShareableContentStyle(rv)
}

