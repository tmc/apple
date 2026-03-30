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

// The class instance for the [CIKernel] class.
var (
	_CIKernelClass     CIKernelClass
	_CIKernelClassOnce sync.Once
)

func getCIKernelClass() CIKernelClass {
	_CIKernelClassOnce.Do(func() {
		_CIKernelClass = CIKernelClass{class: objc.GetClass("CIKernel")}
	})
	return _CIKernelClass
}

// GetCIKernelClass returns the class object for CIKernel.
func GetCIKernelClass() CIKernelClass {
	return getCIKernelClass()
}

type CIKernelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CIKernelClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CIKernelClass) Alloc() CIKernel {
	rv := objc.Send[CIKernel](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A GPU-based image-processing routine used to create custom Core Image
// filters.
//
// # Overview
//
// The kernel language routine for a general-purpose filter kernel has the
// following characteristics:
//
// - Its return type is `vec4` (Core Image Kernel Language) or `float4` (Metal
// Shading Language); that is, it returns a pixel color for the output image.
// - It may use zero or more input images. Each input image is represented by
// a parameter of type `sampler`.
//
// A kernel routine typically produces its output by calculating source image
// coordinates (using the `destCoord` and `samplerTransform` functions or the
// `samplerTransform` function), samples from the source images (using the
// `sample` function), and computes a final pixel color (output using the
// `return` keyword). For example, the Metal Shading Language source below
// implements a filter that passes through its input image unchanged.
//
// The equivalent code in Core Image Kernel Language is:
//
// The Core Image Kernel Language is a dialect of the OpenGL Shading Language.
// See [Core Image Kernel Language Reference] and [Core Image Programming
// Guide] for more details.
//
// # Getting a Kernel Name
//
//   - [CIKernel.Name]: The name of the kernel routine.
//
// # Identifying the Region of Interest for the Kernel
//
//   - [CIKernel.SetROISelector]: Sets the selector Core Image uses to query the region of interest for image processing with the kernel.
//
// # Applying a Kernel to Filter an Image
//
//   - [CIKernel.ApplyWithExtentRoiCallbackArguments]: Creates a new image using the kernel and specified arguments.
//
// See: https://developer.apple.com/documentation/CoreImage/CIKernel
//
// [Core Image Kernel Language Reference]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Reference/CIKernelLangRef/Introduction/Introduction.html#//apple_ref/doc/uid/TP40004397
// [Core Image Programming Guide]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Conceptual/CoreImaging/ci_intro/ci_intro.html#//apple_ref/doc/uid/TP30001185
type CIKernel struct {
	objectivec.Object
}

// CIKernelFromID constructs a [CIKernel] from an objc.ID.
//
// A GPU-based image-processing routine used to create custom Core Image
// filters.
func CIKernelFromID(id objc.ID) CIKernel {
	return CIKernel{objectivec.Object{ID: id}}
}

// NOTE: CIKernel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CIKernel] class.
//
// # Getting a Kernel Name
//
//   - [ICIKernel.Name]: The name of the kernel routine.
//
// # Identifying the Region of Interest for the Kernel
//
//   - [ICIKernel.SetROISelector]: Sets the selector Core Image uses to query the region of interest for image processing with the kernel.
//
// # Applying a Kernel to Filter an Image
//
//   - [ICIKernel.ApplyWithExtentRoiCallbackArguments]: Creates a new image using the kernel and specified arguments.
//
// See: https://developer.apple.com/documentation/CoreImage/CIKernel
type ICIKernel interface {
	objectivec.IObject

	// Topic: Getting a Kernel Name

	// The name of the kernel routine.
	Name() string

	// Topic: Identifying the Region of Interest for the Kernel

	// Sets the selector Core Image uses to query the region of interest for image processing with the kernel.
	SetROISelector(method objc.SEL)

	// Topic: Applying a Kernel to Filter an Image

	// Creates a new image using the kernel and specified arguments.
	ApplyWithExtentRoiCallbackArguments(extent corefoundation.CGRect, callback CIKernelROICallback, args []objectivec.IObject) ICIImage
}

// Init initializes the instance.
func (k CIKernel) Init() CIKernel {
	rv := objc.Send[CIKernel](k.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (k CIKernel) Autorelease() CIKernel {
	rv := objc.Send[CIKernel](k.ID, objc.Sel("autorelease"))
	return rv
}

// NewCIKernel creates a new CIKernel instance.
func NewCIKernel() CIKernel {
	class := getCIKernelClass()
	rv := objc.Send[CIKernel](objc.ID(class.class), objc.Sel("new"))
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
func NewKernelWithFunctionNameFromMetalLibraryDataError(name string, data foundation.INSData) (CIKernel, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getCIKernelClass().class), objc.Sel("kernelWithFunctionName:fromMetalLibraryData:error:"), objc.String(name), data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return CIKernel{}, foundation.NSErrorFrom(errorPtr)
	}
	return CIKernelFromID(rv), nil
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
func NewKernelWithFunctionNameFromMetalLibraryDataOutputPixelFormatError(name string, data foundation.INSData, format int) (CIKernel, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getCIKernelClass().class), objc.Sel("kernelWithFunctionName:fromMetalLibraryData:outputPixelFormat:error:"), objc.String(name), data, format, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return CIKernel{}, foundation.NSErrorFrom(errorPtr)
	}
	return CIKernelFromID(rv), nil
}

// Sets the selector Core Image uses to query the region of interest for image
// processing with the kernel.
//
// method: A selector name.
//
// # Discussion
//
// When applying a filter kernel, the region of interest (ROI) is the area of
// source image pixels that must be processed to produce a given area of
// destination image pixels. For a more detailed definition, see [The Region
// of Interest].
//
// The `aMethod` argument must use the signature that is defined for the “
// method, which is as follows:
//
// `- (CGRect) (int)samplerIndex (CGRect)r obj;`
//
// where:
//
// - `samplerIndex` defines the sampler to query - `destRect` is the extent of
// the region, in working space coordinates, to render. - `userInfo` is the
// object associated with the `kCIApplyOptionUserInfo` option when the kernel
// is applied to its arguments (with the [ApplyArgumentsOptions] method of a
// [CIFilter] object using the kernel). The `userInfo` is important because
// instance variables can’t be used by the defining class. Instance
// variables must be passed through the `userInfo` argument.
//
// The “ method of the CIFilter object is called by the framework. This
// method returns the rectangle that contains the region of the sampler that
// the kernel needs to render the specified destination rectangle.
//
// A sample “ method might look as follows:
//
// If your kernel does not need the image at `index` to produce output in the
// rectangle `rect`, your method should return [CGRectNull].
//
// In the filter code, you set the selector using the following:
//
// `[kernel @selector()`]
//
// Alternatively, use the [ApplyWithExtentRoiCallbackArguments] method to
// directly apply a kernel to create an output image, specifying the ROI
// callback as a block or closure.
//
// See: https://developer.apple.com/documentation/CoreImage/CIKernel/setROISelector(_:)
//
// [CGRectNull]: https://developer.apple.com/documentation/CoreGraphics/CGRectNull
// [The Region of Interest]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Conceptual/CoreImaging/ci_advanced_concepts/ci.advanced_concepts.html#//apple_ref/doc/uid/TP30001185-CH9-SW12
func (k CIKernel) SetROISelector(method objc.SEL) {
	objc.Send[objc.ID](k.ID, objc.Sel("setROISelector:"), method)
}

// Creates a new image using the kernel and specified arguments.
//
// extent: The extent of the output image.
//
// callback: A block or closure that computes the region of interest for a given
// rectangle of destination image pixels. See [CIKernelROICallback].
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
// See: https://developer.apple.com/documentation/CoreImage/CIKernel/apply(extent:roiCallback:arguments:)
//
// [Core Image Kernel Language Reference]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Reference/CIKernelLangRef/Introduction/Introduction.html#//apple_ref/doc/uid/TP40004397
// [The Region of Interest]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Conceptual/CoreImaging/ci_advanced_concepts/ci.advanced_concepts.html#//apple_ref/doc/uid/TP30001185-CH9-SW12
func (k CIKernel) ApplyWithExtentRoiCallbackArguments(extent corefoundation.CGRect, callback CIKernelROICallback, args []objectivec.IObject) ICIImage {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("applyWithExtent:roiCallback:arguments:"), extent, callback, objectivec.IObjectSliceToNSArray(args))
	return CIImageFromID(rv)
}

// Return an array of strings containing the names of all of the kernels
// contained in the Metal library.
//
// data: Contents of the Metal library.
//
// # Return Value
//
// An Array of strings containing the names of the kernels.
//
// See: https://developer.apple.com/documentation/CoreImage/CIKernel/kernelNames(fromMetalLibraryData:)
func (_CIKernelClass CIKernelClass) KernelNamesFromMetalLibraryData(data foundation.INSData) []string {
	rv := objc.Send[[]objc.ID](objc.ID(_CIKernelClass.class), objc.Sel("kernelNamesFromMetalLibraryData:"), data)
	return objc.ConvertSliceToStrings(rv)
}

// Load kernels from a Metal language string.
//
// source: A string containing the progam in Metal language.
//
// # Return Value
//
// An array of [CIKernel] objects.
//
// See: https://developer.apple.com/documentation/CoreImage/CIKernel/kernels(withMetalString:)
func (_CIKernelClass CIKernelClass) KernelsWithMetalStringError(source string) ([]CIKernel, error) {
	var errorPtr objc.ID
	rv := objc.Send[[]objc.ID](objc.ID(_CIKernelClass.class), objc.Sel("kernelsWithMetalString:error:"), objc.String(source), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objc.ConvertSlice(rv, func(id objc.ID) CIKernel {
		return CIKernelFromID(id)
	}), nil

}

// The name of the kernel routine.
//
// # Discussion
//
// The name of a kernel routine is the identifier used to declare it in the
// Core Image Kernel Language source code. For example, if you use the
// [init(source:)] method to create a kernel from the source code below, the
// name of the returned [CIKernel] object is “moveUpTwoPixels”.
//
// See: https://developer.apple.com/documentation/CoreImage/CIKernel/name
//
// [init(source:)]: https://developer.apple.com/documentation/CoreImage/CIKernel/init(source:)
func (k CIKernel) Name() string {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
