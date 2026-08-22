// Code generated from Apple documentation for VideoToolbox. DO NOT EDIT.

package videotoolbox

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VTLowLatencyFrameInterpolationConfiguration] class.
var (
	_VTLowLatencyFrameInterpolationConfigurationClass     VTLowLatencyFrameInterpolationConfigurationClass
	_VTLowLatencyFrameInterpolationConfigurationClassOnce sync.Once
)

func getVTLowLatencyFrameInterpolationConfigurationClass() VTLowLatencyFrameInterpolationConfigurationClass {
	_VTLowLatencyFrameInterpolationConfigurationClassOnce.Do(func() {
		_VTLowLatencyFrameInterpolationConfigurationClass = VTLowLatencyFrameInterpolationConfigurationClass{class: objc.GetClass("VTLowLatencyFrameInterpolationConfiguration")}
	})
	return _VTLowLatencyFrameInterpolationConfigurationClass
}

// GetVTLowLatencyFrameInterpolationConfigurationClass returns the class object for VTLowLatencyFrameInterpolationConfiguration.
func GetVTLowLatencyFrameInterpolationConfigurationClass() VTLowLatencyFrameInterpolationConfigurationClass {
	return getVTLowLatencyFrameInterpolationConfigurationClass()
}

type VTLowLatencyFrameInterpolationConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VTLowLatencyFrameInterpolationConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VTLowLatencyFrameInterpolationConfigurationClass) Alloc() VTLowLatencyFrameInterpolationConfiguration {
	rv := objc.Send[VTLowLatencyFrameInterpolationConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// Configuration that you use to program Video Toolbox frame processor for
// low-latency frame interpolation.
//
// # Overview
//
// This configuration can do either purely temporal interpolation (frame-rate
// conversion) or temporal and spatial interpolation (scaling and frame-rate
// conversion). This processor requires a source frame and a previous frame.
// It does temporal scaling, which interpolates frames between the previous
// frame and the source frame. When performing both temporal and spatial
// interpolation, the processor can only perform 2x upscaling, and a single
// frame of temporal interpolation. When performing spatial scaling, the
// processor produces upscaled intermediate frames and an upscaled
// `sourceFrame`, but it does not upscale the previous reference frame you
// provided.
//
// # Creating a frame interpolation configuration
//
//   - [VTLowLatencyFrameInterpolationConfiguration.InitWithFrameWidthFrameHeightNumberOfInterpolatedFrames]: Creates a new low-latency frame interpolation configuration for frame-rate conversion.
//   - [VTLowLatencyFrameInterpolationConfiguration.InitWithFrameWidthFrameHeightSpatialScaleFactor]: Creates a new low-latency frame interpolation configuration for spatial scaling and temporal scaling.
//
// # Inspecting the configuration
//
//   - [VTLowLatencyFrameInterpolationConfiguration.FrameWidth]: Width of source frames in pixels.
//   - [VTLowLatencyFrameInterpolationConfiguration.FrameHeight]: Height of source frames in pixels.
//   - [VTLowLatencyFrameInterpolationConfiguration.NumberOfInterpolatedFrames]: Number of uniformly spaced frames for which you configured the processor.
//   - [VTLowLatencyFrameInterpolationConfiguration.SpatialScaleFactor]: Configured spatial scale factor as an integer.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencyFrameInterpolationConfiguration
type VTLowLatencyFrameInterpolationConfiguration struct {
	objectivec.Object
}

// VTLowLatencyFrameInterpolationConfigurationFromID constructs a [VTLowLatencyFrameInterpolationConfiguration] from an objc.ID.
//
// Configuration that you use to program Video Toolbox frame processor for
// low-latency frame interpolation.
func VTLowLatencyFrameInterpolationConfigurationFromID(id objc.ID) VTLowLatencyFrameInterpolationConfiguration {
	return VTLowLatencyFrameInterpolationConfiguration{objectivec.Object{ID: id}}
}

// NOTE: VTLowLatencyFrameInterpolationConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VTLowLatencyFrameInterpolationConfiguration] class.
//
// # Creating a frame interpolation configuration
//
//   - [IVTLowLatencyFrameInterpolationConfiguration.InitWithFrameWidthFrameHeightNumberOfInterpolatedFrames]: Creates a new low-latency frame interpolation configuration for frame-rate conversion.
//   - [IVTLowLatencyFrameInterpolationConfiguration.InitWithFrameWidthFrameHeightSpatialScaleFactor]: Creates a new low-latency frame interpolation configuration for spatial scaling and temporal scaling.
//
// # Inspecting the configuration
//
//   - [IVTLowLatencyFrameInterpolationConfiguration.FrameWidth]: Width of source frames in pixels.
//   - [IVTLowLatencyFrameInterpolationConfiguration.FrameHeight]: Height of source frames in pixels.
//   - [IVTLowLatencyFrameInterpolationConfiguration.NumberOfInterpolatedFrames]: Number of uniformly spaced frames for which you configured the processor.
//   - [IVTLowLatencyFrameInterpolationConfiguration.SpatialScaleFactor]: Configured spatial scale factor as an integer.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencyFrameInterpolationConfiguration
type IVTLowLatencyFrameInterpolationConfiguration interface {
	objectivec.IObject

	// Topic: Creating a frame interpolation configuration

	// Creates a new low-latency frame interpolation configuration for frame-rate conversion.
	InitWithFrameWidthFrameHeightNumberOfInterpolatedFrames(frameWidth int, frameHeight int, numberOfInterpolatedFrames int) VTLowLatencyFrameInterpolationConfiguration
	// Creates a new low-latency frame interpolation configuration for spatial scaling and temporal scaling.
	InitWithFrameWidthFrameHeightSpatialScaleFactor(frameWidth int, frameHeight int, spatialScaleFactor int) VTLowLatencyFrameInterpolationConfiguration

	// Topic: Inspecting the configuration

	// Width of source frames in pixels.
	FrameWidth() int
	// Height of source frames in pixels.
	FrameHeight() int
	// Number of uniformly spaced frames for which you configured the processor.
	NumberOfInterpolatedFrames() int
	// Configured spatial scale factor as an integer.
	SpatialScaleFactor() int
}

// Init initializes the instance.
func (v VTLowLatencyFrameInterpolationConfiguration) Init() VTLowLatencyFrameInterpolationConfiguration {
	rv := objc.Send[VTLowLatencyFrameInterpolationConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VTLowLatencyFrameInterpolationConfiguration) Autorelease() VTLowLatencyFrameInterpolationConfiguration {
	rv := objc.Send[VTLowLatencyFrameInterpolationConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVTLowLatencyFrameInterpolationConfiguration creates a new VTLowLatencyFrameInterpolationConfiguration instance.
func NewVTLowLatencyFrameInterpolationConfiguration() VTLowLatencyFrameInterpolationConfiguration {
	class := getVTLowLatencyFrameInterpolationConfigurationClass()
	rv := objc.Send[VTLowLatencyFrameInterpolationConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new low-latency frame interpolation configuration for frame-rate
// conversion.
//
// frameWidth: Width of source frame in pixels.
//
// frameHeight: Height of source frame in pixels.
//
// numberOfInterpolatedFrames: The number of uniformly spaced frames that you want to be used for
// interpolation.
//
// # Discussion
//
// The available interpolation points are the equal to the value of (2^x - 1),
// where x is equal to `numberOfInterpolatedFrames`. For example,
//
// - If you request 1 interpolated frame, 1 interpolation point at 0.5 is
// available. - If you request 2 interpolated frames, 3 interpolation points
// at 0.25, 0.5 and 0.75 are available. You don’t need to use all available
// interpolation points. Setting a higher `numberOfInterpolatedFrames`
// increases the resolution of interpolation in some cases, but also increases
// latency.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencyFrameInterpolationConfiguration/init(frameWidth:frameHeight:numberOfInterpolatedFrames:)
func NewVTLowLatencyFrameInterpolationConfigurationWithFrameWidthFrameHeightNumberOfInterpolatedFrames(frameWidth int, frameHeight int, numberOfInterpolatedFrames int) VTLowLatencyFrameInterpolationConfiguration {
	instance := getVTLowLatencyFrameInterpolationConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrameWidth:frameHeight:numberOfInterpolatedFrames:"), frameWidth, frameHeight, numberOfInterpolatedFrames)
	return VTLowLatencyFrameInterpolationConfigurationFromID(rv)
}

// Creates a new low-latency frame interpolation configuration for spatial
// scaling and temporal scaling.
//
// frameWidth: Width of source frame in pixels.
//
// frameHeight: Height of source frame in pixels.
//
// spatialScaleFactor: The requested spatial scale factor as an integer. Currently, the processor
// supports only 2x spatial scaling.
//
// # Discussion
//
// When you configure the processor for spatial scaling, the low-latency frame
// interpolation processor only supports 2x spatial upscaling and a single
// frame of temporal interpolation at a 0.5 interpolation phase.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencyFrameInterpolationConfiguration/init(frameWidth:frameHeight:spatialScaleFactor:)
func NewVTLowLatencyFrameInterpolationConfigurationWithFrameWidthFrameHeightSpatialScaleFactor(frameWidth int, frameHeight int, spatialScaleFactor int) VTLowLatencyFrameInterpolationConfiguration {
	instance := getVTLowLatencyFrameInterpolationConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrameWidth:frameHeight:spatialScaleFactor:"), frameWidth, frameHeight, spatialScaleFactor)
	return VTLowLatencyFrameInterpolationConfigurationFromID(rv)
}

// Creates a new low-latency frame interpolation configuration for frame-rate
// conversion.
//
// frameWidth: Width of source frame in pixels.
//
// frameHeight: Height of source frame in pixels.
//
// numberOfInterpolatedFrames: The number of uniformly spaced frames that you want to be used for
// interpolation.
//
// # Discussion
//
// The available interpolation points are the equal to the value of (2^x - 1),
// where x is equal to `numberOfInterpolatedFrames`. For example,
//
// - If you request 1 interpolated frame, 1 interpolation point at 0.5 is
// available. - If you request 2 interpolated frames, 3 interpolation points
// at 0.25, 0.5 and 0.75 are available. You don’t need to use all available
// interpolation points. Setting a higher `numberOfInterpolatedFrames`
// increases the resolution of interpolation in some cases, but also increases
// latency.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencyFrameInterpolationConfiguration/init(frameWidth:frameHeight:numberOfInterpolatedFrames:)
func (v VTLowLatencyFrameInterpolationConfiguration) InitWithFrameWidthFrameHeightNumberOfInterpolatedFrames(frameWidth int, frameHeight int, numberOfInterpolatedFrames int) VTLowLatencyFrameInterpolationConfiguration {
	rv := objc.Send[VTLowLatencyFrameInterpolationConfiguration](v.ID, objc.Sel("initWithFrameWidth:frameHeight:numberOfInterpolatedFrames:"), frameWidth, frameHeight, numberOfInterpolatedFrames)
	return rv
}

// Creates a new low-latency frame interpolation configuration for spatial
// scaling and temporal scaling.
//
// frameWidth: Width of source frame in pixels.
//
// frameHeight: Height of source frame in pixels.
//
// spatialScaleFactor: The requested spatial scale factor as an integer. Currently, the processor
// supports only 2x spatial scaling.
//
// # Discussion
//
// When you configure the processor for spatial scaling, the low-latency frame
// interpolation processor only supports 2x spatial upscaling and a single
// frame of temporal interpolation at a 0.5 interpolation phase.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencyFrameInterpolationConfiguration/init(frameWidth:frameHeight:spatialScaleFactor:)
func (v VTLowLatencyFrameInterpolationConfiguration) InitWithFrameWidthFrameHeightSpatialScaleFactor(frameWidth int, frameHeight int, spatialScaleFactor int) VTLowLatencyFrameInterpolationConfiguration {
	rv := objc.Send[VTLowLatencyFrameInterpolationConfiguration](v.ID, objc.Sel("initWithFrameWidth:frameHeight:spatialScaleFactor:"), frameWidth, frameHeight, spatialScaleFactor)
	return rv
}

// Width of source frames in pixels.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencyFrameInterpolationConfiguration/frameWidth
func (v VTLowLatencyFrameInterpolationConfiguration) FrameWidth() int {
	rv := objc.Send[int](v.ID, objc.Sel("frameWidth"))
	return rv
}

// Height of source frames in pixels.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencyFrameInterpolationConfiguration/frameHeight
func (v VTLowLatencyFrameInterpolationConfiguration) FrameHeight() int {
	rv := objc.Send[int](v.ID, objc.Sel("frameHeight"))
	return rv
}

// Number of uniformly spaced frames for which you configured the processor.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencyFrameInterpolationConfiguration/numberOfInterpolatedFrames
func (v VTLowLatencyFrameInterpolationConfiguration) NumberOfInterpolatedFrames() int {
	rv := objc.Send[int](v.ID, objc.Sel("numberOfInterpolatedFrames"))
	return rv
}

// Configured spatial scale factor as an integer.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencyFrameInterpolationConfiguration/spatialScaleFactor
func (v VTLowLatencyFrameInterpolationConfiguration) SpatialScaleFactor() int {
	rv := objc.Send[int](v.ID, objc.Sel("spatialScaleFactor"))
	return rv
}

// Pixel buffer attributes dictionary that describes requirements for pixel
// buffers which represent source frames and reference frames.
//
// # Discussion
//
// Use [CVPixelBufferCreateResolvedAttributesDictionary] to combine this
// dictionary with your pixel buffer attributes dictionary.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencyFrameInterpolationConfiguration/sourcePixelBufferAttributes
func (v VTLowLatencyFrameInterpolationConfiguration) SourcePixelBufferAttributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("sourcePixelBufferAttributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// Pixel buffer attributes dictionary that describes requirements for pixel
// buffers which represent destination frames.
//
// # Discussion
//
// Use [CVPixelBufferCreateResolvedAttributesDictionary] to combine this
// dictionary with your pixel buffer attributes dictionary.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencyFrameInterpolationConfiguration/destinationPixelBufferAttributes
func (v VTLowLatencyFrameInterpolationConfiguration) DestinationPixelBufferAttributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("destinationPixelBufferAttributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// Available supported pixel formats for current configuration.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencyFrameInterpolationConfiguration/frameSupportedPixelFormats
func (v VTLowLatencyFrameInterpolationConfiguration) FrameSupportedPixelFormats() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("frameSupportedPixelFormats"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

// Reports whether the system supports this processor.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencyFrameInterpolationConfiguration/isSupported
func (_VTLowLatencyFrameInterpolationConfigurationClass VTLowLatencyFrameInterpolationConfigurationClass) IsSupported() bool {
	rv := objc.Send[bool](objc.ID(_VTLowLatencyFrameInterpolationConfigurationClass.class), objc.Sel("isSupported"))
	return rv
}

// Protocol methods for VTFrameProcessorConfiguration

// Returns the number of “next” frames that this processor requires for
// processing.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorConfiguration/nextFrameCount-533br
func (o VTLowLatencyFrameInterpolationConfiguration) NextFrameCount() int {
	rv := objc.Send[int](o.ID, objc.Sel("nextFrameCount"))
	return int(rv)
}

// Returns the number of “previous” frames that this processor requires
// for processing.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorConfiguration/previousFrameCount-20ke2
func (o VTLowLatencyFrameInterpolationConfiguration) PreviousFrameCount() int {
	rv := objc.Send[int](o.ID, objc.Sel("previousFrameCount"))
	return int(rv)
}
