// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CIWarpKernel] class.
var (
	_CIWarpKernelClass     CIWarpKernelClass
	_CIWarpKernelClassOnce sync.Once
)

func getCIWarpKernelClass() CIWarpKernelClass {
	_CIWarpKernelClassOnce.Do(func() {
		_CIWarpKernelClass = CIWarpKernelClass{class: objc.GetClass("CIWarpKernel")}
	})
	return _CIWarpKernelClass
}

// GetCIWarpKernelClass returns the class object for CIWarpKernel.
func GetCIWarpKernelClass() CIWarpKernelClass {
	return getCIWarpKernelClass()
}

type CIWarpKernelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CIWarpKernelClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CIWarpKernelClass) Alloc() CIWarpKernel {
	rv := objc.Send[CIWarpKernel](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A GPU-based image-processing routine that processes only the geometry
// information in an image, used to create custom Core Image filters.
//
// # Overview
//
// The kernel language routine for a warp kernel has the following
// characteristics:
//
// - It uses exactly one input image. - Its return type is `vec2` (Core Image
// Kernel Language) or `float2` (Metal Shading Language), specifying a
// position in source image coordinates.
//
// A warp kernel routine requires no input parameters (but can use additional
// custom parameters you declare). Typically, a warp kernel uses the
// destination coordinate function to look up the coordinates of the
// destination pixel currently being rendered, then computes a corresponding
// position in source image coordinates (output using the `return` keyword).
// Core Image then samples from the source image at the returned coordinates
// to produce a pixel color for the output image. For example, the Metal
// Shading Language source below implements a filter that passes through its
// input image unchanged.
//
// The equivalent code in Core Image Kernel Language is:
//
// The Core Image Kernel Language is a dialect of the OpenGL Shading Language.
// See [Core Image Kernel Language Reference] and [Core Image Programming
// Guide] for more details.
//
// # Applying a Kernel to Filter an Image
//
//   - [CIWarpKernel.ApplyWithExtentRoiCallbackInputImageArguments]: Creates a new image using the kernel and the specified input image and arguments.
//
// See: https://developer.apple.com/documentation/CoreImage/CIWarpKernel
//
// [Core Image Kernel Language Reference]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Reference/CIKernelLangRef/Introduction/Introduction.html#//apple_ref/doc/uid/TP40004397
// [Core Image Programming Guide]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Conceptual/CoreImaging/ci_intro/ci_intro.html#//apple_ref/doc/uid/TP30001185
type CIWarpKernel struct {
	CIKernel
}

// CIWarpKernelFromID constructs a [CIWarpKernel] from an objc.ID.
//
// A GPU-based image-processing routine that processes only the geometry
// information in an image, used to create custom Core Image filters.
func CIWarpKernelFromID(id objc.ID) CIWarpKernel {
	return CIWarpKernel{CIKernel: CIKernelFromID(id)}
}

// NOTE: CIWarpKernel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CIWarpKernel] class.
//
// # Applying a Kernel to Filter an Image
//
//   - [ICIWarpKernel.ApplyWithExtentRoiCallbackInputImageArguments]: Creates a new image using the kernel and the specified input image and arguments.
//
// See: https://developer.apple.com/documentation/CoreImage/CIWarpKernel
type ICIWarpKernel interface {
	ICIKernel

	// Topic: Applying a Kernel to Filter an Image

	// Creates a new image using the kernel and the specified input image and arguments.
	ApplyWithExtentRoiCallbackInputImageArguments(extent corefoundation.CGRect, callback CIKernelROICallback, image ICIImage, args []objectivec.IObject) ICIImage
}

// Init initializes the instance.
func (w CIWarpKernel) Init() CIWarpKernel {
	rv := objc.Send[CIWarpKernel](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w CIWarpKernel) Autorelease() CIWarpKernel {
	rv := objc.Send[CIWarpKernel](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewCIWarpKernel creates a new CIWarpKernel instance.
func NewCIWarpKernel() CIWarpKernel {
	class := getCIWarpKernelClass()
	rv := objc.Send[CIWarpKernel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a single kernel object using a Metal Shading Language (MSL) kernel
// function.
//
// # Discussion
//
// - name: The name of the function in the Metal shading language. - data: A
// metallib file compiled with the Core Image Standard Library.
//
// # Discussion
//
// This method allows you to use MSL as the shader language for a Core Image
// kernel. Since MSL based kernels are precompiled, initializing them is
// faster than their than Core Image Kernel Language (CIKL) counterparts and
// Xcode can provide error diagnostics during development rather than at
// runtime. MSL is a more modern language than CIKL, and you can write shader
// code that uses arrays, structs and matrices.
//
// MSL based kernels still support concatenation and tiling and can work in
// the same filter graph as traditional CIKL kernels.
//
// # Specifying Compiler and Linker Options
//
// To use MSL as the shader language for a [CIKernel], you must specify some
// options in Xcode under the tab of your project’s target. The first option
// you need to specify is an `-fcikernel` flag in the Other Metal Compiler
// Flags option. The second is to add a user-defined setting with a key called
// `MTLLINKER_FLAGS` with a value of `-`
//
// [media-2929842]
//
// # Creating a General Kernel in Swift
//
// The following code shows how you can create a general kernel based on a
// Metal function named `myKernel`.
//
// The first step is to create a [Data] object that represents the default
// Metal library. If you have built this in Xcode, it will be called
// `default.Metallib()` and can be loaded using the [Bundle] type’s `url`
// method.
//
// Using the representation of the Metal library and the function name
// `myKernel`, you initialize a [CIKernel].
//
// The code to create general, color, warp and blend filters is identical.
//
// # Metal Shading Language Extensions
//
// Core Image provides a set of language extensions to MSL in
// `CIKernelMetalLib.H()`. These extensions include three new data types for
// working with images: `sampler` (for accessing the input image), `sample_t`
// (represents a single color sample from the input image), and `destination`
// (for accessing the output image). The extensions also include convenience
// functions such as color conversion and premultiply / unpremultiply.
//
// Whereas in CIKL, you typically called global functions when working with,
// for example, coordinates and samples, these functions are implemented as
// member functions on the new types.
//
// The following table shows a summary of CIKL global functions and their
// equivalent MSL functions.
//
// [Table data omitted]
//
// See: https://developer.apple.com/documentation/CoreImage/CIKernel/init(functionName:fromMetalLibraryData:)
//
// [Bundle]: https://developer.apple.com/library/archive/documentation/General/Conceptual/DevPedia-CocoaCore/Bundle.html#//apple_ref/doc/uid/TP40008195-CH4
func NewWarpKernelWithFunctionNameFromMetalLibraryDataError(name string, data foundation.INSData) (CIWarpKernel, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getCIWarpKernelClass().class), objc.Sel("kernelWithFunctionName:fromMetalLibraryData:error:"), objc.String(name), data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return CIWarpKernel{}, foundation.NSErrorFrom(errorPtr)
	}
	return CIWarpKernelFromID(rv), nil
}

// Creates a single kernel object using a Metal Shading Language kernel
// function with optional pixel format.
//
// name: The name of the function in the Metal shading language.
//
// data: A metallib file compiled with the Core Image Standard Library.
//
// format: The pixel format of the output kernel.
//
// # Discussion
//
// This method allows you to use MSL as the shader language for a Core Image
// kernel. Since MSL based kernels are precompiled, initializing them is
// faster than their than Core Image Kernel Language (CIKL) counterparts and
// Xcode can provide error diagnostics during development rather than at
// runtime. MSL is a more modern language than CIKL, and you can write shader
// code that uses arrays, structs and matrices.
//
// MSL based kernels still support concatenation and tiling and can work in
// the same filter graph as traditional CIKL kernels.
//
// See: https://developer.apple.com/documentation/CoreImage/CIKernel/init(functionName:fromMetalLibraryData:outputPixelFormat:)
func NewWarpKernelWithFunctionNameFromMetalLibraryDataOutputPixelFormatError(name string, data foundation.INSData, format int) (CIWarpKernel, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getCIWarpKernelClass().class), objc.Sel("kernelWithFunctionName:fromMetalLibraryData:outputPixelFormat:error:"), objc.String(name), data, format, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return CIWarpKernel{}, foundation.NSErrorFrom(errorPtr)
	}
	return CIWarpKernelFromID(rv), nil
}

// Creates a new image using the kernel and the specified input image and
// arguments.
//
// extent: The extent of the output image.
//
// callback: A block or closure that computes the region of interest for a given
// rectangle of destination image pixels. See [CIKernelROICallback].
//
// image: The input image to be processed by the warp kernel.
//
// args: An array of arguments to pass to the kernel routine. The type of each
// object in the array must be compatible with the corresponding parameter
// declared in the kernel routine source code. For details, see [Core Image
// Kernel Language Reference].
//
// # Return Value
//
// A new image object describing the result of applying the kernel.
//
// # Discussion
//
// This method is analogous to the [CIFilter] method [ApplyArgumentsOptions],
// but it does not require construction of a [CIFilter] object, and it allows
// you to specify a callback for determining the kernel’s region of interest
// as a block or closure. As with the similar [CIFilter] method, calling this
// method does not execute the kernel code—filters and their kernel code are
// evaluated only when rendering a final output image.
//
// When applying a filter kernel, the region of interest (ROI) is the area of
// source image pixels that must be processed to produce a given area of
// destination image pixels. For a more detailed definition, see [The Region
// of Interest]. Core Image calls your `callback` block or closure to
// determine the ROI when rendering the filter output. Core Image
// automatically splits large images into smaller tiles for rendering, so your
// callback may be called multiple times.
//
// See: https://developer.apple.com/documentation/CoreImage/CIWarpKernel/apply(extent:roiCallback:image:arguments:)
//
// [Core Image Kernel Language Reference]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Reference/CIKernelLangRef/Introduction/Introduction.html#//apple_ref/doc/uid/TP40004397
// [The Region of Interest]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Conceptual/CoreImaging/ci_advanced_concepts/ci.advanced_concepts.html#//apple_ref/doc/uid/TP30001185-CH9-SW12
func (w CIWarpKernel) ApplyWithExtentRoiCallbackInputImageArguments(extent corefoundation.CGRect, callback CIKernelROICallback, image ICIImage, args []objectivec.IObject) ICIImage {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("applyWithExtent:roiCallback:inputImage:arguments:"), extent, callback, image, objectivec.IObjectSliceToNSArray(args))
	return CIImageFromID(rv)
}
