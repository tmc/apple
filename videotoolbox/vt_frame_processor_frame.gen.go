// Code generated from Apple documentation for VideoToolbox. DO NOT EDIT.

package videotoolbox

import (
	"sync"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VTFrameProcessorFrame] class.
var (
	_VTFrameProcessorFrameClass     VTFrameProcessorFrameClass
	_VTFrameProcessorFrameClassOnce sync.Once
)

func getVTFrameProcessorFrameClass() VTFrameProcessorFrameClass {
	_VTFrameProcessorFrameClassOnce.Do(func() {
		_VTFrameProcessorFrameClass = VTFrameProcessorFrameClass{class: objc.GetClass("VTFrameProcessorFrame")}
	})
	return _VTFrameProcessorFrameClass
}

// GetVTFrameProcessorFrameClass returns the class object for VTFrameProcessorFrame.
func GetVTFrameProcessorFrameClass() VTFrameProcessorFrameClass {
	return getVTFrameProcessorFrameClass()
}

type VTFrameProcessorFrameClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VTFrameProcessorFrameClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VTFrameProcessorFrameClass) Alloc() VTFrameProcessorFrame {
	rv := objc.Send[VTFrameProcessorFrame](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object that wraps video frames to send to the processor, as source,
// reference, or output frames.
//
// # Overview
//
// Instances retain the buffer backing them.
//
// # Creating a frame object
//
//   - [VTFrameProcessorFrame.InitWithBufferPresentationTimeStamp]: Creates a frame object with a pixel buffer and presentation time.
//
// # Inspecting the frame
//
//   - [VTFrameProcessorFrame.Buffer]: The pixel buffer specified when the object was created.
//   - [VTFrameProcessorFrame.PresentationTimeStamp]: The presentation timestamp specified when the object was created.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorFrame
type VTFrameProcessorFrame struct {
	objectivec.Object
}

// VTFrameProcessorFrameFromID constructs a [VTFrameProcessorFrame] from an objc.ID.
//
// An object that wraps video frames to send to the processor, as source,
// reference, or output frames.
func VTFrameProcessorFrameFromID(id objc.ID) VTFrameProcessorFrame {
	return VTFrameProcessorFrame{objectivec.Object{ID: id}}
}

// NOTE: VTFrameProcessorFrame adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VTFrameProcessorFrame] class.
//
// # Creating a frame object
//
//   - [IVTFrameProcessorFrame.InitWithBufferPresentationTimeStamp]: Creates a frame object with a pixel buffer and presentation time.
//
// # Inspecting the frame
//
//   - [IVTFrameProcessorFrame.Buffer]: The pixel buffer specified when the object was created.
//   - [IVTFrameProcessorFrame.PresentationTimeStamp]: The presentation timestamp specified when the object was created.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorFrame
type IVTFrameProcessorFrame interface {
	objectivec.IObject

	// Topic: Creating a frame object

	// Creates a frame object with a pixel buffer and presentation time.
	InitWithBufferPresentationTimeStamp(buffer corevideo.CVImageBufferRef, presentationTimeStamp coremedia.CMTime) VTFrameProcessorFrame

	// Topic: Inspecting the frame

	// The pixel buffer specified when the object was created.
	Buffer() corevideo.CVImageBufferRef
	// The presentation timestamp specified when the object was created.
	PresentationTimeStamp() coremedia.CMTime
}

// Init initializes the instance.
func (v VTFrameProcessorFrame) Init() VTFrameProcessorFrame {
	rv := objc.Send[VTFrameProcessorFrame](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VTFrameProcessorFrame) Autorelease() VTFrameProcessorFrame {
	rv := objc.Send[VTFrameProcessorFrame](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVTFrameProcessorFrame creates a new VTFrameProcessorFrame instance.
func NewVTFrameProcessorFrame() VTFrameProcessorFrame {
	class := getVTFrameProcessorFrameClass()
	rv := objc.Send[VTFrameProcessorFrame](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a frame object with a pixel buffer and presentation time.
//
// buffer: A pixel buffer for the frame. This value must be non-NULL and IOSurface
// backed.
//
// presentationTimeStamp: The presentation timestamp of the buffer.
//
// # Discussion
//
// Initialization fails if you specify a [NULL] buffer or one that isn’t
// backed by an [IOSurface].
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorFrame/init(buffer:presentationTimeStamp:)
func NewVTFrameProcessorFrameWithBufferPresentationTimeStamp(buffer corevideo.CVImageBufferRef, presentationTimeStamp coremedia.CMTime) VTFrameProcessorFrame {
	instance := getVTFrameProcessorFrameClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBuffer:presentationTimeStamp:"), buffer, presentationTimeStamp)
	return VTFrameProcessorFrameFromID(rv)
}

// Creates a frame object with a pixel buffer and presentation time.
//
// buffer: A pixel buffer for the frame. This value must be non-NULL and IOSurface
// backed.
//
// presentationTimeStamp: The presentation timestamp of the buffer.
//
// # Discussion
//
// Initialization fails if you specify a [NULL] buffer or one that isn’t
// backed by an [IOSurface].
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorFrame/init(buffer:presentationTimeStamp:)
func (v VTFrameProcessorFrame) InitWithBufferPresentationTimeStamp(buffer corevideo.CVImageBufferRef, presentationTimeStamp coremedia.CMTime) VTFrameProcessorFrame {
	rv := objc.Send[VTFrameProcessorFrame](v.ID, objc.Sel("initWithBuffer:presentationTimeStamp:"), buffer, presentationTimeStamp)
	return rv
}

// The pixel buffer specified when the object was created.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorFrame/buffer
func (v VTFrameProcessorFrame) Buffer() corevideo.CVImageBufferRef {
	rv := objc.Send[corevideo.CVImageBufferRef](v.ID, objc.Sel("buffer"))
	return corevideo.CVImageBufferRef(rv)
}

// The presentation timestamp specified when the object was created.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorFrame/presentationTimeStamp
func (v VTFrameProcessorFrame) PresentationTimeStamp() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](v.ID, objc.Sel("presentationTimeStamp"))
	return coremedia.CMTime(rv)
}
