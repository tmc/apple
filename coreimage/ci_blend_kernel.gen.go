// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [CIBlendKernel] class.
var (
	_CIBlendKernelClass     CIBlendKernelClass
	_CIBlendKernelClassOnce sync.Once
)

func getCIBlendKernelClass() CIBlendKernelClass {
	_CIBlendKernelClassOnce.Do(func() {
		_CIBlendKernelClass = CIBlendKernelClass{class: objc.GetClass("CIBlendKernel")}
	})
	return _CIBlendKernelClass
}

// GetCIBlendKernelClass returns the class object for CIBlendKernel.
func GetCIBlendKernelClass() CIBlendKernelClass {
	return getCIBlendKernelClass()
}

type CIBlendKernelClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (cc CIBlendKernelClass) Alloc() CIBlendKernel {
	rv := objc.Send[CIBlendKernel](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A GPU-based image-processing routine that is optimized for blending two
// images.
//
// # Overview
// 
// The blend kernel function has the following characteristics:
// 
// - It has two arguments of type `__sample` (Core Image Kernel Language) or
// `sample_t` (Metal Shading Language), representing the foreground and
// background images. - Its return type is `vec4` (Core Image Kernel Language)
// or `float4` (Metal Shading Language); that is, it returns a pixel color for
// the output image.
// 
// A blend kernel routine receives as input single-pixel colors (one sampled
// from each input image) and computes a final pixel color (output using the
// return keyword). For example, the Metal Shading Language source below
// implements a filter that returns the average of its two input images.
// 
// Generally, the extent of the output image is the union of the extents of
// the foreground and background images.
//
// # Applying a Kernel to Filter an Image
//
//   - [CIBlendKernel.ApplyWithForegroundBackground]: Creates a new image using the blend kernel and specified foreground and background images.
//
// # Instance Methods
//
//   - [CIBlendKernel.ApplyWithForegroundBackgroundColorSpace]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel
type CIBlendKernel struct {
	CIColorKernel
}

// CIBlendKernelFromID constructs a [CIBlendKernel] from an objc.ID.
//
// A GPU-based image-processing routine that is optimized for blending two
// images.
func CIBlendKernelFromID(id objc.ID) CIBlendKernel {
	return CIBlendKernel{CIColorKernel: CIColorKernelFromID(id)}
}
// NOTE: CIBlendKernel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CIBlendKernel] class.
//
// # Applying a Kernel to Filter an Image
//
//   - [ICIBlendKernel.ApplyWithForegroundBackground]: Creates a new image using the blend kernel and specified foreground and background images.
//
// # Instance Methods
//
//   - [ICIBlendKernel.ApplyWithForegroundBackgroundColorSpace]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel
type ICIBlendKernel interface {
	ICIColorKernel

	// Topic: Applying a Kernel to Filter an Image

	// Creates a new image using the blend kernel and specified foreground and background images.
	ApplyWithForegroundBackground(foreground ICIImage, background ICIImage) ICIImage

	// Topic: Instance Methods

	ApplyWithForegroundBackgroundColorSpace(foreground ICIImage, background ICIImage, colorSpace coregraphics.CGColorSpaceRef) ICIImage
}

// Init initializes the instance.
func (b CIBlendKernel) Init() CIBlendKernel {
	rv := objc.Send[CIBlendKernel](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b CIBlendKernel) Autorelease() CIBlendKernel {
	rv := objc.Send[CIBlendKernel](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewCIBlendKernel creates a new CIBlendKernel instance.
func NewCIBlendKernel() CIBlendKernel {
	class := getCIBlendKernelClass()
	rv := objc.Send[CIBlendKernel](objc.ID(class.class), objc.Sel("new"))
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
func NewBlendKernelWithFunctionNameFromMetalLibraryDataError(name string, data foundation.INSData) (CIBlendKernel, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getCIBlendKernelClass().class), objc.Sel("kernelWithFunctionName:fromMetalLibraryData:error:"), objc.String(name), data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return CIBlendKernelFromID(rv), foundation.NSErrorFrom(errorPtr)
	}
	return CIBlendKernelFromID(rv), nil
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
func NewBlendKernelWithFunctionNameFromMetalLibraryDataOutputPixelFormatError(name string, data foundation.INSData, format int) (CIBlendKernel, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getCIBlendKernelClass().class), objc.Sel("kernelWithFunctionName:fromMetalLibraryData:outputPixelFormat:error:"), objc.String(name), data, format, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return CIBlendKernelFromID(rv), foundation.NSErrorFrom(errorPtr)
	}
	return CIBlendKernelFromID(rv), nil
}

// Creates a new image using the blend kernel and specified foreground and
// background images.
//
// foreground: The first input image to be blended
//
// background: The second input image to be blended
//
// # Return Value
// 
// A [CIImage] blending the foreground and background images. Its extent will
// be the union of the foreground and background image extents.
//
// # Discussion
// 
// The foreground and background images are not treated differently in the
// blending. You can think of them as equivalents A and B; the foreground is
// not given any precedence over the background.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/apply(foreground:background:)
func (b CIBlendKernel) ApplyWithForegroundBackground(foreground ICIImage, background ICIImage) ICIImage {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("applyWithForeground:background:"), foreground, background)
	return CIImageFromID(rv)
}
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/apply(foreground:background:colorSpace:)
func (b CIBlendKernel) ApplyWithForegroundBackgroundColorSpace(foreground ICIImage, background ICIImage, colorSpace coregraphics.CGColorSpaceRef) ICIImage {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("applyWithForeground:background:colorSpace:"), foreground, background, colorSpace)
	return CIImageFromID(rv)
}

// A blend kernel that returns a clear color.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/clear
func (_CIBlendKernelClass CIBlendKernelClass) Clear() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("clear"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that uses the luminance values of the background with the
// hue and saturation values of the foreground image.
//
// # Discussion
// 
// [media-2926849]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/color
func (_CIBlendKernelClass CIBlendKernelClass) Color() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("color"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that darkens the background image samples to reflect the
// foreground image samples.
//
// # Discussion
// 
// [media-2926850]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/colorBurn
func (_CIBlendKernelClass CIBlendKernelClass) ColorBurn() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("colorBurn"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that brightens the background image samples to reflect the
// foreground image samples.
//
// # Discussion
// 
// [media-2926851]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/colorDodge
func (_CIBlendKernelClass CIBlendKernelClass) ColorDodge() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("colorDodge"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that adds color components to achieve a brightening effect.
//
// # Discussion
// 
// [media-2926853]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/componentAdd
func (_CIBlendKernelClass CIBlendKernelClass) ComponentAdd() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("componentAdd"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that creates an image using the maximum values of two input
// images.
//
// # Discussion
// 
// [media-2926854]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/componentMax
func (_CIBlendKernelClass CIBlendKernelClass) ComponentMax() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("componentMax"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that creates an image using the minimum values of two input
// images.
//
// # Discussion
// 
// [media-2926852]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/componentMin
func (_CIBlendKernelClass CIBlendKernelClass) ComponentMin() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("componentMin"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that multiplies the color components of its input images.
//
// # Discussion
// 
// [media-2926856]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/componentMultiply
func (_CIBlendKernelClass CIBlendKernelClass) ComponentMultiply() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("componentMultiply"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that creates an image using the darker values of two input
// images.
//
// # Discussion
// 
// [media-2926855]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/darken
func (_CIBlendKernelClass CIBlendKernelClass) Darken() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("darken"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that creates an image using the darker color of two input
// images.
//
// # Discussion
// 
// [media-2926857]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/darkerColor
func (_CIBlendKernelClass CIBlendKernelClass) DarkerColor() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("darkerColor"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that returns the background input image.
//
// # Discussion
// 
// [media-2926859]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/destination
func (_CIBlendKernelClass CIBlendKernelClass) Destination() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("destination"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that places the background over the foreground and crops
// based on the visibility of the foreground.
//
// # Discussion
// 
// [media-2926858]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/destinationAtop
func (_CIBlendKernelClass CIBlendKernelClass) DestinationAtop() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("destinationAtop"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that places the background over the foreground and crops
// based on the visibility of both.
//
// # Discussion
// 
// [media-2926861]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/destinationIn
func (_CIBlendKernelClass CIBlendKernelClass) DestinationIn() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("destinationIn"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that uses the background image to define what to take out of
// the foreground image.
//
// # Discussion
// 
// [media-2926862]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/destinationOut
func (_CIBlendKernelClass CIBlendKernelClass) DestinationOut() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("destinationOut"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that places the background image over the input foreground
// image.
//
// # Discussion
// 
// [media-2926860]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/destinationOver
func (_CIBlendKernelClass CIBlendKernelClass) DestinationOver() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("destinationOver"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that creates an image using the difference between the
// background and foreground images.
//
// # Discussion
// 
// [media-2926863]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/difference
func (_CIBlendKernelClass CIBlendKernelClass) Difference() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("difference"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that divides the background image sample color with the
// foreground image sample color.
//
// # Discussion
// 
// [media-2926864]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/divide
func (_CIBlendKernelClass CIBlendKernelClass) Divide() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("divide"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that produces an effect similar to difference blending but
// with lower contrast.
//
// # Discussion
// 
// [media-2926865]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/exclusion
func (_CIBlendKernelClass CIBlendKernelClass) Exclusion() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("exclusion"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that returns either the foreground or background image if
// the other contains a clear color.
//
// # Discussion
// 
// [media-2926866]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/exclusiveOr
func (_CIBlendKernelClass CIBlendKernelClass) ExclusiveOr() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("exclusiveOr"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that either multiplies or screens colors, depending on the
// source image sample color.
//
// # Discussion
// 
// [media-2926867]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/hardLight
func (_CIBlendKernelClass CIBlendKernelClass) HardLight() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("hardLight"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that adds two images together, setting each color channel
// value to either 0 or 1.
//
// # Discussion
// 
// [media-2926868]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/hardMix
func (_CIBlendKernelClass CIBlendKernelClass) HardMix() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("hardMix"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that uses the luminance and saturation values of the
// background image with the hue of the foreground image.
//
// # Discussion
// 
// [media-2926871]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/hue
func (_CIBlendKernelClass CIBlendKernelClass) Hue() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("hue"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that creates an image using the lighter values of two input
// images.
//
// # Discussion
// 
// [media-2926869]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/lighten
func (_CIBlendKernelClass CIBlendKernelClass) Lighten() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("lighten"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that creates an image using the lighter color of two input
// images.
//
// # Discussion
// 
// [media-2926870]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/lighterColor
func (_CIBlendKernelClass CIBlendKernelClass) LighterColor() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("lighterColor"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that darkens the background image samples to reflect the
// foreground image samples while also increasing contrast.
//
// # Discussion
// 
// [media-2926872]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/linearBurn
func (_CIBlendKernelClass CIBlendKernelClass) LinearBurn() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("linearBurn"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that lightens the background image samples to reflect the
// foreground image samples while also increasing contrast.
//
// # Discussion
// 
// [media-2926874]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/linearDodge
func (_CIBlendKernelClass CIBlendKernelClass) LinearDodge() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("linearDodge"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that burns or dodges colors by changing brightness,
// depending on the blend color.
//
// # Discussion
// 
// [media-2926873]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/linearLight
func (_CIBlendKernelClass CIBlendKernelClass) LinearLight() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("linearLight"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that uses the hue and saturation of the background image
// with the luminance of the foreground image.
//
// # Discussion
// 
// [media-2926875]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/luminosity
func (_CIBlendKernelClass CIBlendKernelClass) Luminosity() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("luminosity"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that multiplies the background image sample color with the
// foreground image sample color.
//
// # Discussion
// 
// [media-2926879]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/multiply
func (_CIBlendKernelClass CIBlendKernelClass) Multiply() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("multiply"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that either multiplies or screens the foreground image
// samples with the background image samples, depending on the background
// color.
//
// # Discussion
// 
// [media-2926878]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/overlay
func (_CIBlendKernelClass CIBlendKernelClass) Overlay() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("overlay"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that conditionally replaces background image samples with
// source image samples depending on the brightness of the source image
// samples.
//
// # Discussion
// 
// [media-2926877]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/pinLight
func (_CIBlendKernelClass CIBlendKernelClass) PinLight() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("pinLight"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that uses the luminance and hue values of the background
// image with the saturation of the foreground image.
//
// # Discussion
// 
// [media-2926876]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/saturation
func (_CIBlendKernelClass CIBlendKernelClass) Saturation() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("saturation"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that multiplies the inverse of the foreground image samples
// with the inverse of the background image samples.
//
// # Discussion
// 
// [media-2926881]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/screen
func (_CIBlendKernelClass CIBlendKernelClass) Screen() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("screen"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that either darkens or lightens colors, depending on the
// foreground image sample color.
//
// # Discussion
// 
// [media-2926882]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/softLight
func (_CIBlendKernelClass CIBlendKernelClass) SoftLight() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("softLight"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that returns the foreground input image.
//
// # Discussion
// 
// [media-2926885]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/source
func (_CIBlendKernelClass CIBlendKernelClass) Source() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("source"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that places the foreground over the background and crops
// based on the visibility of the background.
//
// # Discussion
// 
// [media-2926887]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/sourceAtop
func (_CIBlendKernelClass CIBlendKernelClass) SourceAtop() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("sourceAtop"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that places the foreground over the background and crops
// based on the visibility of both.
//
// # Discussion
// 
// [media-2926883]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/sourceIn
func (_CIBlendKernelClass CIBlendKernelClass) SourceIn() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("sourceIn"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that uses the foreground image to define what to take out of
// the background image.
//
// # Discussion
// 
// [media-2926886]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/sourceOut
func (_CIBlendKernelClass CIBlendKernelClass) SourceOut() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("sourceOut"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that places the foreground image over the input background
// image.
//
// # Discussion
// 
// [media-2926884]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/sourceOver
func (_CIBlendKernelClass CIBlendKernelClass) SourceOver() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("sourceOver"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that subtracts the background image sample color from the
// foreground image sample color.
//
// # Discussion
// 
// [media-2926848]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/subtract
func (_CIBlendKernelClass CIBlendKernelClass) Subtract() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("subtract"))
	return CIBlendKernelFromID(objc.ID(rv))
}
// A blend kernel that burns or dodges colors by changing contrast, depending
// on the blend color.
//
// # Discussion
// 
// [media-2926888]
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendKernel/vividLight
func (_CIBlendKernelClass CIBlendKernelClass) VividLight() CIBlendKernel {
	rv := objc.Send[objc.ID](objc.ID(_CIBlendKernelClass.class), objc.Sel("vividLight"))
	return CIBlendKernelFromID(objc.ID(rv))
}

