// Code generated from Apple documentation for Accelerate. DO NOT EDIT.

// Package accelerate provides Go bindings for the Accelerate framework.
//
// Make large-scale mathematical computations and image calculations, optimized for high performance and low energy consumption.
//
// Accelerate provides high-performance, energy-efficient computation on the CPU by leveraging its vector-processing capability. The following Accelerate libraries abstract that capability so that code written for them executes appropriate instructions for the processor available at runtime:
//
// # Neural Networks
//
//   - Training a neural network to recognize digits: Build a simple neural network and train it to recognize randomly generated numbers.
//   - BNNS: Implement and run neural networks for training and inference. ([Bnns_graph_t], [Bnns_graph_compile_options_t], [BNNSGraphOptimizationPreference], [BNNSGraphMessageLevel], [Bnns_graph_compile_message_fn_t])
//
// # Directories, Files, and Data Archives
//
//   - Compressing single files: Compress a single file and store the result on the file system.
//   - Decompressing single files: Recreate a single file from a compressed file.
//   - Compressing file system directories: Compress the contents of an entire directory and store the result on the file system.
//   - Decompressing and extracting an archived directory: Recreate an entire file system directory from an archive file.
//   - Compressing and saving a string to the file system: Compress the contents of a Unicode string and store the result on the file system.
//   - Decompressing and parsing an archived string: Recreate a string from an archive file.
//
// # Compression
//
//   - Compressing and decompressing files with stream compression: Perform compression for all files and decompression for files with supported extension types.
//   - Compressing and decompressing data with buffer compression: Compress a string, write it to the file system, and decompress the same file using buffer compression.
//   - Compressing and decompressing data with input and output filters: Compress and decompress streamed or from-memory data, using input and output filters.
//
// # Image Processing Essentials
//
//   - Converting bitmap data between Core Graphics images and vImage buffers: Pass image data between Core Graphics and vImage to create and manipulate images.
//   - Creating and Populating Buffers from Core Graphics Images: Initialize vImage buffers from Core Graphics images.
//   - Creating a Core Graphics Image from a vImage Buffer: Create displayable representations of vImage buffers.
//   - Building a Basic Image-Processing Workflow: Resize an image with vImage.
//   - Applying geometric transforms to images: Reflect, shear, rotate, and scale image buffers using vImage.
//   - Compositing images with alpha blending: Combine two images by using alpha blending to create a single output.
//   - Compositing images with vImage blend modes: Combine two images by using blend modes to create a single output.
//   - Applying vImage operations to regions of interest: Limit the effect of vImage operations to rectangular regions of interest.
//   - Optimizing image-processing performance: Improve your app’s performance by converting image buffer formats from interleaved to planar.
//   - vImage: Manipulate large images using the CPU’s vector processor.
//
// # Signal Processing Essentials
//
//   - Controlling vDSP operations with stride: Operate selectively on the elements of a vector at regular intervals.
//   - Using linear interpolation to construct new data points: Fill the gaps in arrays of numerical data using linear interpolation.
//   - Using vDSP for vector-based arithmetic: Increase the performance of common mathematical tasks with vDSP vector-vector and vector-scalar operations.
//   - Resampling a signal with decimation: Reduce the sample rate of a signal by specifying a decimation factor and applying a custom antialiasing filter.
//   - vDSP: Perform basic arithmetic operations and common digital signal processing (DSP) routines on large vectors. ([VDSP_Length], [VDSP_Stride], [DSPComplex], [COMPLEX_SPLIT], [DSPDoubleComplex])
//
// # Fourier and Cosine Transforms
//
//   - Understanding data packing for Fourier transforms: Format source data for the vDSP Fourier functions, and interpret the results.
//   - Finding the component frequencies in a composite sine wave: Use 1D fast Fourier transform to compute the frequency components of a signal.
//   - Performing Fourier transforms on interleaved-complex data: Optimize discrete Fourier transform (DFT) performance with the vDSP interleaved DFT routines.
//   - Reducing spectral leakage with windowing: Multiply signal data by window sequence values when performing transforms with noninteger period signals.
//   - Signal extraction from noise: Use Accelerate’s discrete cosine transform to remove noise from a signal.
//   - Performing Fourier Transforms on Multiple Signals: Use Accelerate’s multiple-signal fast Fourier transform (FFT) functions to transform multiple signals with a single function call.
//   - Halftone descreening with 2D fast Fourier transform: Reduce or remove periodic artifacts from images.
//   - Fast Fourier transforms: Transform vectors and matrices of temporal and spatial domain complex values to the frequency domain, and vice versa. ([FFTSetup], [FFTSetupD], [FFTRadix], [FFTDirection])
//   - Discrete Fourier transforms: Transform vectors of temporal and spatial domain complex values to the frequency domain, and vice versa. ([VDSP_DFT_Interleaved_Setup], [VDSP_DFT_Interleaved_SetupD], [VDSP_DFT_Setup], [VDSP_DFT_SetupD], [VDSP_DFT_Direction])
//   - Discrete Cosine transforms: Transform vectors of temporal and spatial domain real values to the frequency domain, and vice versa. ([VDSP_DCT_Type])
//
// # Core Video Interoperation
//
//   - Using vImage pixel buffers to generate video effects: Render real-time video effects with the vImage Pixel Buffer.
//   - Integrating vImage pixel buffers into a Core Image workflow: Share image data between Core Video pixel buffers and vImage buffers to integrate vImage operations into a Core Image workflow.
//   - Applying vImage operations to video sample buffers: Use the vImage convert-any-to-any functionality to perform real-time image processing of video frames streamed from your device’s camera.
//   - Improving the quality of quantized images with dithering: Apply dithering to simulate colors that are unavailable in reduced bit depths.
//   - Core Video interoperability: Pass image data between Core Video and vImage.
//
// # Vectors, Matrices, and Quaternions
//
//   - Working with Vectors: Use vectors to calculate geometric values, calculate dot products and cross products, and interpolate between values.
//   - Working with Matrices: Solve simultaneous equations and transform points in space.
//   - Working with Quaternions: Rotate points around the surface of a sphere, and interpolate between them.
//   - Rotating a cube by transforming its vertices: Rotate a cube through a series of keyframes using quaternion interpolation to transition between them.
//   - simd: Perform computations on small vectors and matrices. ([Simd_bool], [Simd_quatf], [Simd_quatd])
//   - vForce: Perform transcendental and trigonometric functions on vectors of any length. ([COMPLEX], [DOUBLE_COMPLEX])
//
// # Audio Processing
//
//   - Visualizing sound as an audio spectrogram: Share image data between vDSP and vImage to visualize audio that a device microphone captures.
//   - Applying biquadratic filters to a music loop: Change the frequency response of an audio signal using a cascaded biquadratic filter.
//   - Equalizing audio with discrete cosine transforms (DCTs): Change the frequency response of an audio signal by manipulating frequency-domain data.
//   - Biquadratic IIR filters: Apply biquadratic filters to single-channel and multichannel data.
//   - Discrete Cosine transforms: Transform vectors of temporal and spatial domain real values to the frequency domain, and vice versa. ([VDSP_DCT_Type])
//
// # Conversion Between Image Formats
//
//   - Building a basic image conversion workflow: Learn the fundamentals of the convert-any-to-any function by converting a CMYK image to an RGB image.
//   - Converting color images to grayscale: Convert an RGB image to grayscale using matrix multiplication.
//   - Applying color transforms to images with a multidimensional lookup table: Precompute translation values to optimize color space conversion and other pointwise operations.
//   - Building a basic image conversion workflow: Learn the fundamentals of the convert-any-to-any function by converting a CMYK image to an RGB image.
//   - Converting luminance and chrominance planes to an ARGB image: Create a displayable ARGB image using the luminance and chrominance information from your device’s camera.
//   - Conversion: Convert an image to a different format.
//
// # Image Resampling
//
//   - Resampling in vImage: Learn how vImage resamples image data during geometric operations.
//   - Reducing artifacts with custom resampling filters: Implement custom linear interpolation to prevent the ringing effects associated with scaling an image with the default Lanczos algorithm.
//   - Image shearing: Shear images horizontally and vertically.
//
// # Convolution and Morphology
//
//   - Blurring an image: Filter an image by convolving it with custom and high-speed kernels.
//   - Adding a bokeh effect to images: Simulate a bokeh effect by applying dilation.
//   - Convolution: Apply a convolution kernel to an image.
//   - Morphology: Dilate and erode images.
//
// # Color and Tone Adjustment
//
//   - Adjusting the brightness and contrast of an image: Use a gamma function to apply a linear or exponential curve.
//   - Adjusting saturation and applying tone mapping: Convert an RGB image to discrete luminance and chrominance channels, and apply color and contrast treatments.
//   - Applying tone curve adjustments to images: Use the vImage library’s polynomial transform to apply tone curve adjustments to images.
//   - Adjusting the hue of an image: Convert an image to L*a*b* color space and apply hue adjustment.
//   - Specifying histograms with vImage: Calculate the histogram of one image, and apply it to a second image.
//   - Enhancing image contrast with histogram manipulation: Enhance and adjust the contrast of an image with histogram equalization and contrast stretching.
//   - Histogram: Calculate or manipulate an image’s histogram.
//
// # vImage / vDSP Interoperability
//
//   - Finding the sharpest image in a sequence of captured images: Share image data between vDSP and vImage to compute the sharpest image from a bracketed photo sequence.
//   - Visualizing sound as an audio spectrogram: Share image data between vDSP and vImage to visualize audio that a device microphone captures.
//
// # Sparse Matrices
//
//   - Creating sparse matrices: Create sparse matrices for factorization and solving systems.
//   - Solving systems using direct methods: Use direct methods to solve systems of equations where the coefficient matrix is sparse.
//   - Solving systems using iterative methods: Use iterative methods to solve systems of equations where the coefficient matrix is sparse.
//   - Creating a sparse matrix from coordinate format arrays: Use separate coordinate format arrays to create sparse matrices.
//   - Sparse Solvers: Solve systems of equations where the coefficient matrix is sparse. ([SparseMatrix_Double], [SparseMatrix_Float], [DenseMatrix_Double], [DenseMatrix_Float], [DenseVector_Double])
//
// # Arithmetic and Transcendental Functions
//
//   - vecLib: Perform computations on large vectors.
//
// # Linear Algebra
//
//   - Solving systems of linear equations with LAPACK: Select the optimal LAPACK routine to solve a system of linear equations.
//   - Finding an interpolating polynomial using the Vandermonde method: Use LAPACK to solve a linear system and find an interpolating polynomial to construct new points between a series of known data points.
//   - Compressing an image using linear algebra: Reduce the storage size of an image using singular value decomposition (SVD).
//   - BLAS: Perform common linear algebra operations with Apple’s implementation of the Basic Linear Algebra Subprograms (BLAS). ([BLAS_THREADING], [BLASParamErrorProc], [CBLAS_ORDER], [CBLAS_TRANSPOSE], [CBLAS_UPLO])
//
// # Definite Integration
//
//   - Quadrature: Approximate the definite integral of a function over a finite or infinite interval. ([Quadrature_function_array], [Quadrature_integrate_function], [Quadrature_integrate_options])
//
// # Macros
//
//   - Macros
//
// [Accelerate Documentation]: https://developer.apple.com/documentation/Accelerate
package accelerate

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the Accelerate library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/Accelerate.framework/Accelerate",
	"/usr/lib/libAccelerate.dylib",
}

// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	for _, path := range frameworkPaths {
		h, err := purego.Dlopen(path, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
		if err == nil {
			frameworkHandle = h
			return
		}
	}
	fmt.Fprintf(os.Stderr, "warning: Accelerate: failed to load framework from any known path\n")
}
