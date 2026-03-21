
// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

// Package coreimage provides Go bindings for the CoreImage framework.
//
// Use built-in or custom filters to process still and video images.
//
// Core Image is an image processing and analysis technology that provides high-performance processing for still and video images. Use the many built-in image filters to process images and build complex effects by chaining filters. For a list of all the built-in filters see the [Filter Catalog](<doc://com.apple.coreimage/documentation/CoreImage#Filter-Catalog>).
//
// # Essentials
//
//   - Processing an Image Using Built-in Filters: Apply effects such as sepia tint, highlight strengthening, and scaling to images.
//   - CIContext: The Core Image context class provides an evaluation context for Core Image processing with Metal, OpenGL, or OpenCL. ([CIImageRepresentationOption], [CIContextOption])
//   - CIImage: A representation of an image to be processed or produced by Core Image filters. ([CIImageOption], [CIImageAutoAdjustmentOption])
//
// # Filters
//
//   - CIFilter: An image processor that produces an image by manipulating one or more input images or by generating new image data. ([CIFilterProtocol], [CIDynamicRangeOption])
//   - CIRAWFilter: A filter subclass that produces an image by manipulating RAW image sensor data from a digital camera or scanner. ([CIRAWDecoderVersion])
//   - CIColor: The Core Image class that defines a color object.
//   - CIVector: The Core Image class that defines a vector object.
//
// # Filter Catalog
//
//   - Blur Filters: Apply blurs, simulate motion and zoom effects, reduce noise, and erode and dilate image regions. ([CIBokehBlur], [CIBoxBlur], [CIDiscBlur], [CIGaussianBlur], [CIMaskedVariableBlur])
//   - Color Adjustment Filters: Apply color transformations, including exposure, hue, and tint adjustments. ([CIColorAbsoluteDifference], [CIColorClamp], [CIColorControls], [CIColorMatrix], [CIColorPolynomial])
//   - Color Effect Filters: Apply color effects, including photo effects, dithering, and color maps. ([CIColorCrossPolynomial], [CIColorCube], [CIColorCubeWithColorSpace], [CIColorCubesMixedWithMask], [CIColorCurves])
//   - Composite Operations: Composite images by using a range of blend modes and compositing operators. ([CICompositeOperation])
//   - Convolution Filters: Produce effects such as blurring, sharpening, edge detection, translation, and embossing. ([CIConvolution])
//   - Distortion Filters: Apply distortion to images. ([CIBumpDistortion], [CIBumpDistortionLinear], [CICircleSplashDistortion], [CICircularWrap], [CIDisplacementDistortion])
//   - Generator Filters: Generate barcode, geometric, and special-effect images. ([CICode128BarcodeGenerator], [CIAttributedTextImageGenerator], [CIAztecCodeGenerator], [CIBarcodeGenerator], [CIBlurredRectangleGenerator])
//   - Geometry Adjustment Filters: Translate, scale, and rotate images in 2D and 3D. ([CIBicubicScaleTransform], [CIEdgePreserveUpsample], [CIFourCoordinateGeometryFilter], [CIKeystoneCorrectionCombined], [CIKeystoneCorrectionHorizontal])
//   - Gradient Filters: Generate linear and radial gradients. ([CIGaussianGradient], [CIHueSaturationValueGradient], [CILinearGradient], [CIRadialGradient], [CISmoothLinearGradient])
//   - Halftone Effect Filters: Simulate monochrome and CMYK halftone screens. ([CICircularScreen], [CICMYKHalftone], [CIDotScreen], [CIHatchedScreen], [CILineScreen])
//   - Reduction Filters: Create statistical information about an image. ([CIAreaAverage], [CIAreaHistogram], [CIAreaLogarithmicHistogram], [CIAreaMaximum], [CIAreaMaximumAlpha])
//   - Sharpening Filters: Apply sharpening to images. ([CISharpenLuminance], [CIUnsharpMask])
//   - Stylizing Filters: Create stylized versions of images by applying effects including pixelation and line overlays. ([CIBlendWithMask], [CIBloom], [CICannyEdgeDetector], [CIComicEffect], [CICoreMLModel])
//   - Tile Effect Filters: Produce tiled images from source images. ([CIAffineClamp], [CIAffineTile], [CIEightfoldReflectedTile], [CIFourfoldReflectedTile], [CIFourfoldRotatedTile])
//   - Transition Filters: Transition between two images by using effects including page curl and swipe. ([CITransitionFilter], [CIBarsSwipeTransition], [CIAccordionFoldTransition], [CICopyMachineTransition], [CIDisintegrateWithMaskTransition])
//
// # Filter Recipes
//
//   - Applying a Chroma Key Effect: Replace a color in one image with the background from another.
//   - Selectively Focusing on an Image: Focus on a part of an image by applying Gaussian blur and gradient masks.
//   - Customizing Image Transitions: Transition between images in creative ways using Core Image filters.
//   - Simulating Scratchy Analog Film: Degrade the quality of an image to make it look like dated, analog film.
//
// # Custom Filters
//
//   - Writing Custom Kernels: Write your own custom kernels in either the Core Image Kernel Language or the Metal Shading Language.
//   - CIKernel: A GPU-based image-processing routine used to create custom Core Image filters. ([CIKernelROICallback])
//   - CIColorKernel: A GPU-based image-processing routine that processes only the color information in images, used to create custom Core Image filters.
//   - CIWarpKernel: A GPU-based image-processing routine that processes only the geometry information in an image, used to create custom Core Image filters.
//   - CIBlendKernel: A GPU-based image-processing routine that is optimized for blending two images.
//   - CISampler: An object that retrieves pixel samples for processing by a filter kernel.
//   - CIFilterShape: A description of the bounding shape of a filter and the domain of definition for a filter operation.
//   - CIFormat: Pixel data formats for image input, output, and processing.
//
// # Custom Image Processors
//
//   - CIImageProcessorKernel: The abstract class you extend to create custom image processors that can integrate with Core Image workflows.
//   - CIImageProcessorInput: A container of image data and information for use in a custom image processor.
//   - CIImageProcessorOutput: A container for writing image data and information produced by a custom image processor.
//
// # Custom Render Destination
//
//   - Generating an animation with a Core Image Render Destination: Animate a filtered image to a Metal view in a SwiftUI app using a Core Image Render Destination.
//   - CIRenderDestination: A specification for configuring all attributes of a render task’s destination and issuing asynchronous render tasks. ([CIRenderDestinationAlphaMode])
//   - CIRenderInfo: An encapsulation of a render task’s timing, passes, and pixels processed.
//   - CIRenderTask: A single render task.
//   - CIRenderDestinationAlphaMode: Different ways of representing alpha.
//
// # Feedback-Based Processing
//
//   - CIImageAccumulator: An object that manages feedback-based image processing for tasks such as painting or fluid simulation.
//
// # Barcode Descriptions
//
//   - CIBarcodeDescriptor: An abstract base class that represents a machine-readable code’s attributes.
//   - CIQRCodeDescriptor: A concrete subclass of the Core Image Barcode Descriptor that represents a square QR code symbol. ([ErrorCorrectedPayload], [SymbolVersion], [MaskPattern], [ErrorCorrectionLevel])
//   - CIAztecCodeDescriptor: A concrete subclass the Core Image Barcode Descriptor that represents an Aztec code symbol. ([ErrorCorrectedPayload], [IsCompact], [LayerCount], [DataCodewordCount])
//   - CIPDF417CodeDescriptor: A concrete subclass of Core Image Barcode Descriptor that represents a PDF417 symbol. ([ErrorCorrectedPayload], [IsCompact], [RowCount], [ColumnCount])
//   - CIDataMatrixCodeDescriptor: A concrete subclass the Core Image Barcode Descriptor that represents an Data Matrix code symbol. ([ErrorCorrectedPayload], [RowCount], [ColumnCount], [EccVersion])
//
// # Image Feature Detection
//
//   - CIDetector: An image processor that identifies notable features, such as faces and barcodes, in a still image or video.
//   - CIFeature: The abstract superclass for objects representing notable features detected in an image. ([Bounds])
//   - CIFaceFeature: Information about a face detected in a still or video image. ([Bounds], [HasFaceAngle], [FaceAngle], [HasLeftEyePosition], [HasRightEyePosition])
//   - CIRectangleFeature: Information about a rectangular region detected in a still or video image. ([Bounds], [BottomLeft], [BottomRight], [TopLeft], [TopRight])
//   - CITextFeature: Information about a text that was detected in a still or video image. ([Bounds], [BottomLeft], [BottomRight], [TopLeft], [TopRight])
//   - CIQRCodeFeature: Information about a Quick Response code detected in a still or video image. ([Bounds], [SymbolDescriptor], [BottomLeft], [BottomRight], [TopLeft])
//
// # Image Units
//
//   - CIPlugIn: The mechanism for loading image units in macOS.
//   - CIFilterGenerator: An object that creates and configures chains of individual image filters.
//   - CIPlugInRegistration: The interface for loading Core Image image units.
//   - CIFilterConstructor: A general interface for objects that produce filters.
//
// # Protocols
//
//   - CIAreaBoundsRed
//   - CIMaximumScaleTransform
//   - CIToneMapHeadroom
//   - CIAreaAverageMaximumRed: The protocol for the Area Average and Maximum Red filter.
//   - CIBlurredRoundedRectangleGenerator: The protocol for the Blurred Rounded Rectangle Generator filter.
//   - CIDistanceGradientFromRedMask: The protocol for the Distance Gradient From Red Mask filter.
//   - CIRoundedQRCodeGenerator: The protocol for the Rounded QR Code Generator filter.
//   - CISignedDistanceGradientFromRedMask: The protocol for the Signed Distance Gradient From Red Mask filter.
//
// # Variables
//
//   - kCIInputBacksideImageKey: A key to get or set the backside image for a transition Core Image filter.
//   - kCIInputBiasVectorKey: A key to get or set the vector bias value of a Core Image filter.
//   - kCIInputColor0Key: A key to get or set a color value of a Core Image filter.
//   - kCIInputColor1Key: A key to get or set a color value of a Core Image filter.
//   - kCIInputColorSpaceKey: A key to get or set a color space value of a Core Image filter.
//   - kCIInputCountKey: A key to get or set the scalar count value of a Core Image filter.
//   - kCIInputExtrapolateKey: A key to get or set the boolean behavior of a Core Image filter that specifies if the filter should extrapolate a table beyond the defined range.
//   - kCIInputPaletteImageKey: A key to get or set the palette image for a  Core Image filter.
//   - kCIInputPerceptualKey: A key to get or set the boolean behavior of a Core Image filter that specifies if the filter should operate in linear or perceptual colors.
//   - kCIInputPoint0Key: A key to get or set the coordinate value of a Core Image filter.
//   - kCIInputPoint1Key: A key to get or set a coordinate value of a Core Image filter.
//   - kCIInputRadius0Key: A key to get or set the geometric radius value of a Core Image filter.
//   - kCIInputRadius1Key: A key to get or set the geometric radius value of a Core Image filter.
//   - kCIInputThresholdKey: A key to get or set the scalar threshold value of a Core Image filter.
//
// # Key Types
//
//   - [CIFilter] - An image processor that produces an image by manipulating one or more input images or by generating new image data.
//   - [CIImage] - A representation of an image to be processed or produced by Core Image filters.
//   - [CIContext] - The Core Image context class provides an evaluation context for Core Image processing with Metal, OpenGL, or OpenCL.
//   - [CIRAWFilter] - A filter subclass that produces an image by manipulating RAW image sensor data from a digital camera or scanner.
//   - [CIBlendKernel] - A GPU-based image-processing routine that is optimized for blending two images.
//   - [CIColor] - The Core Image class that defines a color object.
//   - [CIVector] - The Core Image class that defines a vector object.
//   - [CIFaceFeature] - Information about a face detected in a still or video image.
//   - [CIRenderDestination] - A specification for configuring all attributes of a render task’s destination and issuing asynchronous render tasks.
//   - [CIFilterGenerator] - An object that creates and configures chains of individual image filters.
//
// [CoreImage Documentation]: https://developer.apple.com/documentation/CoreImage
package coreimage

import (
	"fmt"
	"os"
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/CoreImage.framework/CoreImage"
// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	var err error
	frameworkHandle, err = purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: CoreImage: failed to load %s: %v\n", frameworkPath, err)
		return 
	}
}

