// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CIImageProcessorKernel] class.
var (
	_CIImageProcessorKernelClass     CIImageProcessorKernelClass
	_CIImageProcessorKernelClassOnce sync.Once
)

func getCIImageProcessorKernelClass() CIImageProcessorKernelClass {
	_CIImageProcessorKernelClassOnce.Do(func() {
		_CIImageProcessorKernelClass = CIImageProcessorKernelClass{class: objc.GetClass("CIImageProcessorKernel")}
	})
	return _CIImageProcessorKernelClass
}

// GetCIImageProcessorKernelClass returns the class object for CIImageProcessorKernel.
func GetCIImageProcessorKernelClass() CIImageProcessorKernelClass {
	return getCIImageProcessorKernelClass()
}

type CIImageProcessorKernelClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (cc CIImageProcessorKernelClass) Alloc() CIImageProcessorKernel {
	rv := objc.Send[CIImageProcessorKernel](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// The abstract class you extend to create custom image processors that can
// integrate with Core Image workflows.
//
// # Overview
// 
// Unlike the [CIKernel] class and its other subclasses that allow you to
// create new image-processing effects with the Core Image Kernel Language,
// the [CIImageProcessorKernel] class provides direct access to the underlying
// bitmap image data for a step in the Core Image processing pipeline. As
// such, you can create subclasses of this class to integrate other
// image-processing technologies—such as Metal compute shaders, [Metal
// Performance Shaders], [Accelerate] [vImage] operations, or your own
// CPU-based image-processing routines—with a Core Image filter chain.
// 
// Your custom image processing operation is invoked by your subclassed image
// processor kernel’s [CIImageProcessorKernel.ProcessWithInputsArgumentsOutputError] method. The
// method can accept zero, one or more inputs: kernels that generate imagery
// (such as a noise or pattern generator) need no inputs, while kernels that
// composite source images together require multiple inputs. The `arguments`
// dictionary allows the caller to pass in additional parameter values (such
// as the radius of a blur) and the `output` contains the destination for your
// image processing code to write to.
// 
// The following code shows how you can subclass [CIImageProcessorKernel] to
// apply the Metal Performance Shader [MPSImageThresholdBinary] kernel to a
// [CIImage]:
// 
// To apply to kernel to an image, the calling side invokes the image
// processor’s [CIImageProcessorKernel.ApplyWithExtentInputsArgumentsError] method. The following
// code generates a new [CIImage] object named `result` which contains a
// thresholded version of the source image, `inputImage`.
// 
// # Subclassing Notes
// 
// The [CIImageProcessorKernel] class is abstract; to create a custom image
// processor, you define a subclass of this class.
// 
// You do not directly create instances of a custom [CIImageProcessorKernel]
// subclass. Image processors must not carry or use state specific to any
// single invocation of the processor, so all methods (and accessors for
// readonly properties) of an image processor kernel class are class methods.
// 
// Your subclass should override at least the
// [CIImageProcessorKernel.ProcessWithInputsArgumentsOutputError] method to perform its image
// processing.
// 
// If your image processor needs to work with a larger or smaller region of
// interest in the input image than each corresponding region of the output
// image (for example, a blur filter, which samples several input pixels for
// each output pixel), you should also override the
// [CIImageProcessorKernel.RoiForInputArgumentsOutputRect] method.
// 
// You can also override the [CIImageProcessorKernel.FormatForInputAtIndex] method and [CIImageProcessorKernel.OutputFormat]
// property getter to customize the input and output pixel formats for your
// processor (for example, as part of a multi-step workflow where you extract
// a single channel from an RGBA image, apply an effect to that channel only,
// then recombine the channels).
// 
// # Using a Custom Image Processor
// 
// To apply your custom image processor class to filter one or more images,
// call the [CIImageProcessorKernel.ApplyWithExtentInputsArgumentsError] class method. (Do not
// override this method.)
//
// [Accelerate]: https://developer.apple.com/library/archive/documentation/MacOSX/Conceptual/OSX_Technology_Overview/CoreOSLayer/CoreOSLayer.html#//apple_ref/doc/uid/TP40001067-CH9-SW6
// [MPSImageThresholdBinary]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdBinary
// [Metal Performance Shaders]: https://developer.apple.com/library/archive/releasenotes/General/WhatsNewIniOS/Articles/iOS9.html#//apple_ref/doc/uid/TP40016198-SW7
// [vImage]: https://developer.apple.com/library/archive/releasenotes/Performance/RN-vecLib/index.html#//apple_ref/doc/uid/TP40001049-CH2-SW2
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorKernel
type CIImageProcessorKernel struct {
	objectivec.Object
}

// CIImageProcessorKernelFromID constructs a [CIImageProcessorKernel] from an objc.ID.
//
// The abstract class you extend to create custom image processors that can
// integrate with Core Image workflows.
func CIImageProcessorKernelFromID(id objc.ID) CIImageProcessorKernel {
	return CIImageProcessorKernel{objectivec.Object{ID: id}}
}
// NOTE: CIImageProcessorKernel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CIImageProcessorKernel] class.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorKernel
type ICIImageProcessorKernel interface {
	objectivec.IObject
}

// Init initializes the instance.
func (i CIImageProcessorKernel) Init() CIImageProcessorKernel {
	rv := objc.Send[CIImageProcessorKernel](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i CIImageProcessorKernel) Autorelease() CIImageProcessorKernel {
	rv := objc.Send[CIImageProcessorKernel](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewCIImageProcessorKernel creates a new CIImageProcessorKernel instance.
func NewCIImageProcessorKernel() CIImageProcessorKernel {
	class := getCIImageProcessorKernelClass()
	rv := objc.Send[CIImageProcessorKernel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Call this method on your Core Image Processor Kernel subclass to create a
// new image of the specified extent.
//
// extent: The bounding [CGRect] of pixels that the [CIImageProcessorKernel] can
// produce. This method will return `/CIImage/emptyImage` if extent is empty.
//
// inputs: An array of [CIImage] objects to use as input.
//
// arguments: This dictionary contains any additional parameters that the processor needs
// to produce its output. The argument objects can be of any type but in order
// for CoreImage to cache intermediates, they must be of the following
// immutable types: [NSArray], [NSDictionary], [NSNumber], [NSValue],
// [NSData], [NSString], [NSNull], [CIVector], [CIColor], [CGImage],
// [CGColorSpace], or [MLModel].
//
// # Return Value
// 
// An autoreleased [CIImage]
//
// # Discussion
// 
// The inputs and arguments will be retained so that your subclass can be
// called when the image is drawn.
// 
// This method will return `nil` and an error if:
// 
// - calling [OutputFormat] on your subclass returns an unsupported format. -
// calling [FormatForInputAtIndex] on your subclass returns an unsupported
// format. - your subclass does not implement
// [ProcessWithInputsArgumentsOutputError]
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorKernel/apply(withExtent:inputs:arguments:)
func (_CIImageProcessorKernelClass CIImageProcessorKernelClass) ApplyWithExtentInputsArgumentsError(extent corefoundation.CGRect, inputs []CIImage, arguments foundation.INSDictionary) (CIImage, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_CIImageProcessorKernelClass.class), objc.Sel("applyWithExtent:inputs:arguments:error:"), extent, objectivec.IObjectSliceToNSArray(inputs), arguments, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return CIImage{}, foundation.NSErrorFrom(errorPtr)
	}
	return CIImageFromID(rv), nil

}
// Override this class method if you want your any of the inputs to be in a
// specific pixel format.
//
// # Discussion
// 
// The format must be one of `kCIFormatBGRA8`, `kCIFormatRGBAh`,
// `kCIFormatRGBAf` or `kCIFormatR8`. On iOS 12 and macOS 10.14, the formats
// `kCIFormatRh` and `kCIFormatRf` are also supported.
// 
// If the requested inputFormat is `0`, then the input will be a supported
// format that best matches the rendering context’s
// `/CIContext/workingFormat`.
// 
// If a processor wants data in a colorspace other than the context’s
// working color space, then call `/CIImage/` on the processor input. If a
// processor wants it input as alpha-unpremultiplied RGBA data, then call
// `/CIImage/imageByUnpremultiplyingAlpha` on the processor input.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorKernel/formatForInput(at:)
func (_CIImageProcessorKernelClass CIImageProcessorKernelClass) FormatForInputAtIndex(inputIndex int) CIFormat {
	rv := objc.Send[CIFormat](objc.ID(_CIImageProcessorKernelClass.class), objc.Sel("formatForInputAtIndex:"), inputIndex)
	return CIFormat(rv)
}
// Override this class method to implement your Core Image Processor Kernel
// subclass.
//
// inputs: An array of `id` that the class consumes to produce its output. The
// `input.Region()` may be larger than the rect returned by
// [RoiForInputArgumentsOutputRect].
//
// arguments: The arguments dictionary that was passed to
// [ApplyWithExtentInputsArgumentsError].
//
// output: The `id` that the [CIImageProcessorKernel] must provide results to.
//
// # Discussion
// 
// When a [CIImage] containing your [CIImageProcessorKernel] class is
// rendered, your class’ implementation of this method will be called as
// needed for that render. The method may be called more than once if Core
// Image needs to tile to limit memory usage.
// 
// When your implementation of this class method is called, use the provided
// `inputs` and `arguments` objects to return processed pixel data to Core
// Image via `output`.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorKernel/process(with:arguments:output:)
func (_CIImageProcessorKernelClass CIImageProcessorKernelClass) ProcessWithInputsArgumentsOutputError(inputs []objectivec.IObject, arguments foundation.INSDictionary, output CIImageProcessorOutput) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_CIImageProcessorKernelClass.class), objc.Sel("processWithInputs:arguments:output:error:"), objectivec.IObjectSliceToNSArray(inputs), arguments, output, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("processWithInputs:arguments:output:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Override this class method to implement your processor’s ROI callback.
//
// inputIndex: The index that tells you which processor input for which to return the ROI
// rectangle.
//
// arguments: The arguments dictionary that was passed to
// [ApplyWithExtentInputsArgumentsError].
//
// outputRect: The output [CGRect] that processor will be asked to output.
//
// # Return Value
// 
// The [CGRect] of the `inputIndex`th input that is required for the above
// `outputRect`
//
// # Discussion
// 
// This will be called one or more times per render to determine what portion
// of the input images are needed to render a given ‘outputRect’ of the
// output. This will not be called if processor has no input images.
// 
// The default implementation would return outputRect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorKernel/roi(forInput:arguments:outputRect:)
func (_CIImageProcessorKernelClass CIImageProcessorKernelClass) RoiForInputArgumentsOutputRect(inputIndex int, arguments foundation.INSDictionary, outputRect corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](objc.ID(_CIImageProcessorKernelClass.class), objc.Sel("roiForInput:arguments:outputRect:"), inputIndex, arguments, outputRect)
	return corefoundation.CGRect(rv)
}
// Override this class method to implement your processor’s tiled ROI
// callback.
//
// inputIndex: The index that tells you which processor input for which to return the
// array of ROI rectangles
//
// arguments: The arguments dictionary that was passed to
// [ApplyWithExtentInputsArgumentsError].
//
// outputRect: The output [CGRect] that processor will be asked to output.
//
// # Return Value
// 
// An array of [CIVector] that specify tile regions of the `inputIndex`’th
// input that is required for the above `outputRect` Each region tile in the
// array is a created by calling `/CIVector//` The tiles may overlap but
// should fully cover the area of ‘input’ that is needed. If a processor
// has multiple inputs, then each input should return the same number of
// region tiles.
//
// # Discussion
// 
// This will be called one or more times per render to determine what tiles of
// the input images are needed to render a given `outputRect` of the output.
// 
// If the processor implements this method, then when rendered;
// 
// - as CoreImage prepares for a render, this method will be called for each
// input to return an ROI tile array. - as CoreImage performs the render, the
// method [ProcessWithInputsArgumentsOutputError] will be called once for each
// tile.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorKernel/roiTileArray(forInput:arguments:outputRect:)
func (_CIImageProcessorKernelClass CIImageProcessorKernelClass) RoiTileArrayForInputArgumentsOutputRect(inputIndex int, arguments foundation.INSDictionary, outputRect corefoundation.CGRect) []CIVector {
	rv := objc.Send[[]objc.ID](objc.ID(_CIImageProcessorKernelClass.class), objc.Sel("roiTileArrayForInput:arguments:outputRect:"), inputIndex, arguments, outputRect)
	return objc.ConvertSlice(rv, func(id objc.ID) CIVector {
		return CIVectorFromID(id)
	})
}
// Call this method on your multiple-output Core Image Processor Kernel
// subclass to create an array of new image objects given the specified array
// of extents.
//
// extents: The array of bounding rectangles that the [CIImageProcessorKernel] can
// produce. Each rectangle in the array is an object created using
// `/CIVector/` This method will return `CIImage.EmptyImage()` if a rectangle
// in the array is empty.
//
// inputs: An array of [CIImage] objects to use as input.
//
// arguments: This dictionary contains any additional parameters that the processor needs
// to produce its output. The argument objects can be of any type but in order
// for CoreImage to cache intermediates, they must be of the following
// immutable types: [NSArray], [NSDictionary], [NSNumber], [NSValue],
// [NSData], [NSString], [NSNull], [CIVector], [CIColor], [CGImage],
// [CGColorSpace], or [MLModel].
//
// # Return Value
// 
// An autoreleased [CIImage]
//
// # Discussion
// 
// The inputs and arguments will be retained so that your subclass can be
// called when the image is drawn.
// 
// This method will return `nil` and an error if:
// 
// - calling [OutputFormatAtIndexArguments] on your subclass returns an
// unsupported format. - calling [FormatForInputAtIndex] on your subclass
// returns an unsupported format. - your subclass does not implement
// [ProcessWithInputsArgumentsOutputError]
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorKernel/apply(withExtents:inputs:arguments:)
func (_CIImageProcessorKernelClass CIImageProcessorKernelClass) ApplyWithExtentsInputsArgumentsError(extents []CIVector, inputs []CIImage, arguments foundation.INSDictionary) ([]CIImage, error) {
	var errorPtr objc.ID
	rv := objc.Send[[]objc.ID](objc.ID(_CIImageProcessorKernelClass.class), objc.Sel("applyWithExtents:inputs:arguments:error:"), objectivec.IObjectSliceToNSArray(extents), objectivec.IObjectSliceToNSArray(inputs), arguments, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objc.ConvertSlice(rv, func(id objc.ID) CIImage {
		return CIImageFromID(id)
	}), nil

}
// Override this class method if your processor has more than one output and
// you want your processor’s output to be in a specific supported
// [CIPixelFormat].
//
// outputIndex: The index that tells you which processor output for which to return the
// desired [CIPixelFormat]
//
// arguments: The arguments dictionary that was passed to
// [ApplyWithExtentInputsArgumentsError].
//
// # Return Value
// 
// Return the desired [CIPixelFormat]
//
// # Discussion
// 
// The format must be one of `kCIFormatBGRA8`, `kCIFormatRGBAh`,
// `kCIFormatRGBAf` or `kCIFormatR8`. On iOS 12 and macOS 10.14, the formats
// `kCIFormatRh` and `kCIFormatRf` are also supported.
// 
// If the outputFormat is `0`, then the output will be a supported format that
// best matches the rendering context’s `/CIContext/workingFormat`.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorKernel/outputFormat(at:arguments:)
func (_CIImageProcessorKernelClass CIImageProcessorKernelClass) OutputFormatAtIndexArguments(outputIndex int, arguments foundation.INSDictionary) CIFormat {
	rv := objc.Send[CIFormat](objc.ID(_CIImageProcessorKernelClass.class), objc.Sel("outputFormatAtIndex:arguments:"), outputIndex, arguments)
	return CIFormat(rv)
}
// Override this class method of your Core Image Processor Kernel subclass if
// it needs to produce multiple outputs.
//
// inputs: An array of `id` that the class consumes to produce its output. The
// `input.Region()` may be larger than the rect returned by
// [RoiForInputArgumentsOutputRect].
//
// arguments: The arguments dictionary that was passed to
// [ApplyWithExtentInputsArgumentsError].
//
// outputs: An array `id` that the [CIImageProcessorKernel] must provide results to.
//
// # Discussion
// 
// This supports 0, 1, 2 or more input images and 2 or more output images.
// 
// When a [CIImage] containing your [CIImageProcessorKernel] class is
// rendered, your class’ implementation of this method will be called as
// needed for that render. The method may be called more than once if Core
// Image needs to tile to limit memory usage.
// 
// When your implementation of this class method is called, use the provided
// `inputs` and `arguments` objects to return processed pixel data to Core
// Image via multiple `outputs`.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorKernel/process(with:arguments:outputs:)
func (_CIImageProcessorKernelClass CIImageProcessorKernelClass) ProcessWithInputsArgumentsOutputsError(inputs []objectivec.IObject, arguments foundation.INSDictionary, outputs []objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_CIImageProcessorKernelClass.class), objc.Sel("processWithInputs:arguments:outputs:error:"), objectivec.IObjectSliceToNSArray(inputs), arguments, objectivec.IObjectSliceToNSArray(outputs), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("processWithInputs:arguments:outputs:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Override this class property if you want your processor’s output to be in
// a specific pixel format.
//
// # Discussion
// 
// The format must be one of `kCIFormatBGRA8`, `kCIFormatRGBAh`,
// `kCIFormatRGBAf` or `kCIFormatR8`. On iOS 12 and macOS 10.14, the formats
// `kCIFormatRh` and `kCIFormatRf` are also supported.
// 
// If the outputFormat is `0`, then the output will be a supported format that
// best matches the rendering context’s `/CIContext/workingFormat`.
// 
// If a processor returns data in a color space other than the context working
// color space, then call `/CIImage/` on the processor output. If a processor
// returns data as alpha-unpremultiplied RGBA data, then call,
// `/CIImage/imageByPremultiplyingAlpha` on the processor output.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorKernel/outputFormat
func (_CIImageProcessorKernelClass CIImageProcessorKernelClass) OutputFormat() CIFormat {
	rv := objc.Send[CIFormat](objc.ID(_CIImageProcessorKernelClass.class), objc.Sel("outputFormat"))
	return CIFormat(rv)
}
// Override this class property if your processor’s output stores 1.0 into
// the alpha channel of all pixels within the output extent.
//
// # Discussion
// 
// If not overridden, false is returned.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorKernel/outputIsOpaque
func (_CIImageProcessorKernelClass CIImageProcessorKernelClass) OutputIsOpaque() bool {
	rv := objc.Send[bool](objc.ID(_CIImageProcessorKernelClass.class), objc.Sel("outputIsOpaque"))
	return rv
}
// Override this class property to return false if you want your processor to
// be given input objects that have not been synchronized for CPU access.
//
// # Discussion
// 
// Generally, if your subclass uses the GPU your should override this method
// to return false. If not overridden, true is returned.
//
// See: https://developer.apple.com/documentation/CoreImage/CIImageProcessorKernel/synchronizeInputs
func (_CIImageProcessorKernelClass CIImageProcessorKernelClass) SynchronizeInputs() bool {
	rv := objc.Send[bool](objc.ID(_CIImageProcessorKernelClass.class), objc.Sel("synchronizeInputs"))
	return rv
}

