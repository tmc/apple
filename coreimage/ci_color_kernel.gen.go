// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CIColorKernel] class.
var (
	_CIColorKernelClass     CIColorKernelClass
	_CIColorKernelClassOnce sync.Once
)

func getCIColorKernelClass() CIColorKernelClass {
	_CIColorKernelClassOnce.Do(func() {
		_CIColorKernelClass = CIColorKernelClass{class: objc.GetClass("CIColorKernel")}
	})
	return _CIColorKernelClass
}

// GetCIColorKernelClass returns the class object for CIColorKernel.
func GetCIColorKernelClass() CIColorKernelClass {
	return getCIColorKernelClass()
}

type CIColorKernelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CIColorKernelClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CIColorKernelClass) Alloc() CIColorKernel {
	rv := objc.Send[CIColorKernel](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A GPU-based image-processing routine that processes only the color
// information in images, used to create custom Core Image filters.
//
// # Overview
// 
// The kernel language routine for a color kernel has the following
// characteristics:
// 
// - Its return type is `vec4` (Core Image Kernel Language) or `float4` (Metal
// Shading Language); that is, it returns a pixel color for the output image.
// - It may use zero or more input images. Each input image is represented by
// a parameter of type `__sample` (Core Image Kernel Language) or `sample_t`
// (Metal Shading Language), which can be treated as a single pixel color of
// type `vec4` (Core Image Kernel Language) or `float4` (Metal Shading
// Language);.
// 
// A color kernel routine receives as input single-pixel colors (one sampled
// from each input image) and computes a final pixel color (output using the
// `return` keyword). For example, the Metal Shading Language source below
// implements a filter that passes through its input image unchanged.
// 
// The equivalent code in Core Image Kernel Language is:
// 
// The Core Image Kernel Language is a dialect of the OpenGL Shading Language.
// See [Core Image Kernel Language Reference] and [Core Image Programming
// Guide] for more details.
//
// [Core Image Kernel Language Reference]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Reference/CIKernelLangRef/Introduction/Introduction.html#//apple_ref/doc/uid/TP40004397
// [Core Image Programming Guide]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Conceptual/CoreImaging/ci_intro/ci_intro.html#//apple_ref/doc/uid/TP30001185
//
// # Applying a Kernel to Filter an Image
//
//   - [CIColorKernel.ApplyWithExtentArguments]: Creates a new image using the kernel and specified arguments.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorKernel
type CIColorKernel struct {
	CIKernel
}

// CIColorKernelFromID constructs a [CIColorKernel] from an objc.ID.
//
// A GPU-based image-processing routine that processes only the color
// information in images, used to create custom Core Image filters.
func CIColorKernelFromID(id objc.ID) CIColorKernel {
	return CIColorKernel{CIKernel: CIKernelFromID(id)}
}
// NOTE: CIColorKernel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CIColorKernel] class.
//
// # Applying a Kernel to Filter an Image
//
//   - [ICIColorKernel.ApplyWithExtentArguments]: Creates a new image using the kernel and specified arguments.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorKernel
type ICIColorKernel interface {
	ICIKernel

	// Topic: Applying a Kernel to Filter an Image

	// Creates a new image using the kernel and specified arguments.
	ApplyWithExtentArguments(extent corefoundation.CGRect, args []objectivec.IObject) ICIImage
}

// Init initializes the instance.
func (c CIColorKernel) Init() CIColorKernel {
	rv := objc.Send[CIColorKernel](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CIColorKernel) Autorelease() CIColorKernel {
	rv := objc.Send[CIColorKernel](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCIColorKernel creates a new CIColorKernel instance.
func NewCIColorKernel() CIColorKernel {
	class := getCIColorKernelClass()
	rv := objc.Send[CIColorKernel](objc.ID(class.class), objc.Sel("new"))
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
// [Bundle]: https://developer.apple.com/library/archive/documentation/General/Conceptual/DevPedia-CocoaCore/Bundle.html#//apple_ref/doc/uid/TP40008195-CH4
//
// See: https://developer.apple.com/documentation/CoreImage/CIKernel/init(functionName:fromMetalLibraryData:)
func NewColorKernelWithFunctionNameFromMetalLibraryDataError(name string, data foundation.INSData) (CIColorKernel, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getCIColorKernelClass().class), objc.Sel("kernelWithFunctionName:fromMetalLibraryData:error:"), objc.String(name), data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return CIColorKernel{}, foundation.NSErrorFrom(errorPtr)
	}
	return CIColorKernelFromID(rv), nil
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
func NewColorKernelWithFunctionNameFromMetalLibraryDataOutputPixelFormatError(name string, data foundation.INSData, format int) (CIColorKernel, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getCIColorKernelClass().class), objc.Sel("kernelWithFunctionName:fromMetalLibraryData:outputPixelFormat:error:"), objc.String(name), data, format, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return CIColorKernel{}, foundation.NSErrorFrom(errorPtr)
	}
	return CIColorKernelFromID(rv), nil
}

// Creates a new image using the kernel and specified arguments.
//
// extent: The extent of the output image.
//
// args: An array of arguments to pass to the kernel routine. The type of each
// object in the array must be compatible with the corresponding parameter
// declared in the kernel routine source code. For details, see [Core Image
// Kernel Language Reference].
// //
// [Core Image Kernel Language Reference]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Reference/CIKernelLangRef/Introduction/Introduction.html#//apple_ref/doc/uid/TP40004397
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
// See: https://developer.apple.com/documentation/CoreImage/CIColorKernel/apply(extent:arguments:)
func (c CIColorKernel) ApplyWithExtentArguments(extent corefoundation.CGRect, args []objectivec.IObject) ICIImage {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("applyWithExtent:arguments:"), extent, objectivec.IObjectSliceToNSArray(args))
	return CIImageFromID(rv)
}

