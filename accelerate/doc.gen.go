// Code generated from Apple documentation for Accelerate. DO NOT EDIT.

// Package accelerate provides Go bindings for the Accelerate framework.
//
// Make large-scale mathematical computations and image calculations,
// optimized for high performance and low energy consumption.
//
// Accelerate provides high-performance, energy-efficient computation on the
// CPU by leveraging its vector-processing capability. The following
// Accelerate libraries abstract that capability so that code written for them
// executes appropriate instructions for the processor available at runtime:
//
// # Neural Networks
//
//   - [Training a neural network to recognize digits]: Build a simple neural network and train it to recognize randomly generated numbers.
//   - BNNS: Implement and run neural networks for training and inference. ([Bnns_graph_t], [Bnns_graph_compile_options_t], [BNNSGraphOptimizationPreference], [BNNSGraphMessageLevel])
//
// # Directories, Files, and Data Archives
//
//   - [Compressing single files]: Compress a single file and store the result on the file system.
//   - [Decompressing single files]: Recreate a single file from a compressed file.
//   - [Compressing file system directories]: Compress the contents of an entire directory and store the result on the file system.
//   - [Decompressing and extracting an archived directory]: Recreate an entire file system directory from an archive file.
//   - [Compressing and saving a string to the file system]: Compress the contents of a Unicode string and store the result on the file system.
//   - [Decompressing and parsing an archived string]: Recreate a string from an archive file.
//
// # Compression
//
//   - [Compressing and decompressing files with stream compression]: Perform compression for all files and decompression for files with supported extension types.
//   - [Compressing and decompressing data with buffer compression]: Compress a string, write it to the file system, and decompress the same file using buffer compression.
//   - [Compressing and decompressing data with input and output filters]: Compress and decompress streamed or from-memory data, using input and output filters.
//
// # Image Processing Essentials
//
//   - [Converting bitmap data between Core Graphics images and vImage buffers]: Pass image data between Core Graphics and vImage to create and manipulate images.
//   - [Creating and Populating Buffers from Core Graphics Images]: Initialize vImage buffers from Core Graphics images.
//   - [Creating a Core Graphics Image from a vImage Buffer]: Create displayable representations of vImage buffers.
//   - [Building a Basic Image-Processing Workflow]: Resize an image with vImage.
//   - [Applying geometric transforms to images]: Reflect, shear, rotate, and scale image buffers using vImage.
//   - [Compositing images with alpha blending]: Combine two images by using alpha blending to create a single output.
//   - [Compositing images with vImage blend modes]: Combine two images by using blend modes to create a single output.
//   - [Applying vImage operations to regions of interest]: Limit the effect of vImage operations to rectangular regions of interest.
//   - [Optimizing image-processing performance]: Improve your app’s performance by converting image buffer formats from interleaved to planar.
//
// # Signal Processing Essentials
//
//   - [Controlling vDSP operations with stride]: Operate selectively on the elements of a vector at regular intervals.
//   - [Using linear interpolation to construct new data points]: Fill the gaps in arrays of numerical data using linear interpolation.
//   - [Using vDSP for vector-based arithmetic]: Increase the performance of common mathematical tasks with vDSP vector-vector and vector-scalar operations.
//   - [Resampling a signal with decimation]: Reduce the sample rate of a signal by specifying a decimation factor and applying a custom antialiasing filter.
//   - vDSP: Perform basic arithmetic operations and common digital signal processing (DSP) routines on large vectors. ([DSPComplex], [COMPLEX_SPLIT], [DSPDoubleComplex])
//
// # Fourier and Cosine Transforms
//
//   - [Understanding data packing for Fourier transforms]: Format source data for the vDSP Fourier functions, and interpret the results.
//   - [Finding the component frequencies in a composite sine wave]: Use 1D fast Fourier transform to compute the frequency components of a signal.
//   - [Performing Fourier transforms on interleaved-complex data]: Optimize discrete Fourier transform (DFT) performance with the vDSP interleaved DFT routines.
//   - [Reducing spectral leakage with windowing]: Multiply signal data by window sequence values when performing transforms with noninteger period signals.
//   - [Signal extraction from noise]: Use Accelerate’s discrete cosine transform to remove noise from a signal.
//   - [Performing Fourier Transforms on Multiple Signals]: Use Accelerate’s multiple-signal fast Fourier transform (FFT) functions to transform multiple signals with a single function call.
//   - [Halftone descreening with 2D fast Fourier transform]: Reduce or remove periodic artifacts from images.
//   - [Fast Fourier transforms]: Transform vectors and matrices of temporal and spatial domain complex values to the frequency domain, and vice versa. ([FFTSetup], [FFTSetupD], [FFTRadix], [FFTDirection])
//   - [Discrete Fourier transforms]: Transform vectors of temporal and spatial domain complex values to the frequency domain, and vice versa.
//   - [Discrete Cosine transforms]: Transform vectors of temporal and spatial domain real values to the frequency domain, and vice versa.
//
// # Core Video Interoperation
//
//   - [Using vImage pixel buffers to generate video effects]: Render real-time video effects with the vImage Pixel Buffer.
//   - [Integrating vImage pixel buffers into a Core Image workflow]: Share image data between Core Video pixel buffers and vImage buffers to integrate vImage operations into a Core Image workflow.
//   - [Applying vImage operations to video sample buffers]: Use the vImage convert-any-to-any functionality to perform real-time image processing of video frames streamed from your device’s camera.
//   - [Improving the quality of quantized images with dithering]: Apply dithering to simulate colors that are unavailable in reduced bit depths.
//   - [Core Video interoperability]: Pass image data between Core Video and vImage.
//
// # Vectors, Matrices, and Quaternions
//
//   - [Working with Vectors]: Use vectors to calculate geometric values, calculate dot products and cross products, and interpolate between values.
//   - [Working with Matrices]: Solve simultaneous equations and transform points in space.
//   - [Working with Quaternions]: Rotate points around the surface of a sphere, and interpolate between them.
//   - [Rotating a cube by transforming its vertices]: Rotate a cube through a series of keyframes using quaternion interpolation to transition between them.
//   - vForce: Perform transcendental and trigonometric functions on vectors of any length. ([COMPLEX], [DOUBLE_COMPLEX])
//
// # Audio Processing
//
//   - [Visualizing sound as an audio spectrogram]: Share image data between vDSP and vImage to visualize audio that a device microphone captures.
//   - [Applying biquadratic filters to a music loop]: Change the frequency response of an audio signal using a cascaded biquadratic filter.
//   - [Biquadratic IIR filters]: Apply biquadratic filters to single-channel and multichannel data.
//   - [Discrete Cosine transforms]: Transform vectors of temporal and spatial domain real values to the frequency domain, and vice versa.
//
// # Conversion Between Image Formats
//
//   - [Building a basic image conversion workflow]: Learn the fundamentals of the convert-any-to-any function by converting a CMYK image to an RGB image.
//   - [Converting color images to grayscale]: Convert an RGB image to grayscale using matrix multiplication.
//   - [Applying color transforms to images with a multidimensional lookup table]: Precompute translation values to optimize color space conversion and other pointwise operations.
//   - [Building a basic image conversion workflow]: Learn the fundamentals of the convert-any-to-any function by converting a CMYK image to an RGB image.
//   - [Converting luminance and chrominance planes to an ARGB image]: Create a displayable ARGB image using the luminance and chrominance information from your device’s camera.
//
// # Image Resampling
//
//   - [Resampling in vImage]: Learn how vImage resamples image data during geometric operations.
//   - [Reducing artifacts with custom resampling filters]: Implement custom linear interpolation to prevent the ringing effects associated with scaling an image with the default Lanczos algorithm.
//   - [Image shearing]: Shear images horizontally and vertically.
//
// # Convolution and Morphology
//
//   - [Blurring an image]: Filter an image by convolving it with custom and high-speed kernels.
//   - [Adding a bokeh effect to images]: Simulate a bokeh effect by applying dilation.
//
// # Color and Tone Adjustment
//
//   - [Adjusting the brightness and contrast of an image]: Use a gamma function to apply a linear or exponential curve.
//   - [Adjusting saturation and applying tone mapping]: Convert an RGB image to discrete luminance and chrominance channels, and apply color and contrast treatments.
//   - [Applying tone curve adjustments to images]: Use the vImage library’s polynomial transform to apply tone curve adjustments to images.
//   - [Adjusting the hue of an image]: Convert an image to L*a*b* color space and apply hue adjustment.
//   - [Specifying histograms with vImage]: Calculate the histogram of one image, and apply it to a second image.
//   - [Enhancing image contrast with histogram manipulation]: Enhance and adjust the contrast of an image with histogram equalization and contrast stretching.
//
// # vImage / vDSP Interoperability
//
//   - [Finding the sharpest image in a sequence of captured images]: Share image data between vDSP and vImage to compute the sharpest image from a bracketed photo sequence.
//   - [Visualizing sound as an audio spectrogram]: Share image data between vDSP and vImage to visualize audio that a device microphone captures.
//
// # Sparse Matrices
//
//   - [Creating sparse matrices]: Create sparse matrices for factorization and solving systems.
//   - [Solving systems using direct methods]: Use direct methods to solve systems of equations where the coefficient matrix is sparse.
//   - [Solving systems using iterative methods]: Use iterative methods to solve systems of equations where the coefficient matrix is sparse.
//   - [Creating a sparse matrix from coordinate format arrays]: Use separate coordinate format arrays to create sparse matrices.
//   - [Sparse Solvers]: Solve systems of equations where the coefficient matrix is sparse. ([SparseMatrix_Double], [SparseMatrix_Float], [DenseMatrix_Double], [DenseMatrix_Float], [DenseVector_Double])
//
// # Linear Algebra
//
//   - [Solving systems of linear equations with LAPACK]: Select the optimal LAPACK routine to solve a system of linear equations.
//   - [Finding an interpolating polynomial using the Vandermonde method]: Use LAPACK to solve a linear system and find an interpolating polynomial to construct new points between a series of known data points.
//   - [Compressing an image using linear algebra]: Reduce the storage size of an image using singular value decomposition (SVD).
//   - BLAS: Perform common linear algebra operations with Apple’s implementation of the Basic Linear Algebra Subprograms (BLAS). ([BLAS_THREADING], [BLASParamErrorProc], [CBLAS_ORDER], [CBLAS_TRANSPOSE], [CBLAS_UPLO])
//
// # Functions
//
//   - [Sparse_inner_product_dense_double_complex]
//   - [Sparse_inner_product_dense_float_complex]
//   - [Sparse_inner_product_sparse_double_complex]
//   - [Sparse_inner_product_sparse_float_complex]
//   - [Sparse_insert_entry_double_complex]
//   - [Sparse_insert_entry_float_complex]
//   - [Sparse_matrix_product_dense_double_complex]
//   - [Sparse_matrix_product_dense_float_complex]
//   - [Sparse_matrix_product_sparse_double_complex]
//   - [Sparse_matrix_product_sparse_float_complex]
//   - [Sparse_matrix_trace_double_complex]
//   - [Sparse_matrix_trace_float_complex]
//   - [Sparse_matrix_triangular_solve_dense_double_complex]
//   - [Sparse_matrix_triangular_solve_dense_float_complex]
//   - [Sparse_matrix_vector_product_dense_double_complex]
//   - [Sparse_matrix_vector_product_dense_float_complex]
//   - [Sparse_outer_product_dense_double_complex]
//   - [Sparse_outer_product_dense_float_complex]
//   - [Sparse_vector_add_with_scale_dense_double_complex]
//   - [Sparse_vector_add_with_scale_dense_float_complex]
//   - [Sparse_vector_triangular_solve_dense_double_complex]
//   - [Sparse_vector_triangular_solve_dense_float_complex]//
//
// [Adding a bokeh effect to images]: https://developer.apple.com/documentation/accelerate/adding-a-bokeh-effect-to-images
// [Adjusting saturation and applying tone mapping]: https://developer.apple.com/documentation/accelerate/adjusting-saturation-and-applying-tone-mapping
// [Adjusting the brightness and contrast of an image]: https://developer.apple.com/documentation/accelerate/adjusting-the-brightness-and-contrast-of-an-image
// [Adjusting the hue of an image]: https://developer.apple.com/documentation/accelerate/adjusting-the-hue-of-an-image
// [Applying biquadratic filters to a music loop]: https://developer.apple.com/documentation/accelerate/applying-biquadratic-filters-to-a-music-loop
// [Applying color transforms to images with a multidimensional lookup table]: https://developer.apple.com/documentation/accelerate/applying-color-transforms-to-images-with-a-multidimensional-lookup-table
// [Applying geometric transforms to images]: https://developer.apple.com/documentation/accelerate/applying-geometric-transforms-to-images
// [Applying tone curve adjustments to images]: https://developer.apple.com/documentation/accelerate/applying-tone-curve-adjustments-to-images
// [Applying vImage operations to regions of interest]: https://developer.apple.com/documentation/accelerate/applying-vimage-operations-to-regions-of-interest
// [Applying vImage operations to video sample buffers]: https://developer.apple.com/documentation/accelerate/applying-vimage-operations-to-video-sample-buffers
// [Biquadratic IIR filters]: https://developer.apple.com/documentation/accelerate/biquadratic-iir-filters
// [Blurring an image]: https://developer.apple.com/documentation/accelerate/blurring-an-image
// [Building a Basic Image-Processing Workflow]: https://developer.apple.com/documentation/accelerate/building-a-basic-image-processing-workflow
// [Building a basic image conversion workflow]: https://developer.apple.com/documentation/accelerate/building-a-basic-image-conversion-workflow
// [Compositing images with alpha blending]: https://developer.apple.com/documentation/accelerate/compositing-images-with-alpha-blending
// [Compositing images with vImage blend modes]: https://developer.apple.com/documentation/accelerate/compositing-images-with-vimage-blend-modes
// [Compressing an image using linear algebra]: https://developer.apple.com/documentation/accelerate/compressing-an-image-using-linear-algebra
// [Compressing and decompressing data with buffer compression]: https://developer.apple.com/documentation/accelerate/compressing-and-decompressing-data-with-buffer-compression
// [Compressing and decompressing data with input and output filters]: https://developer.apple.com/documentation/accelerate/compressing-and-decompressing-data-with-input-and-output-filters
// [Compressing and decompressing files with stream compression]: https://developer.apple.com/documentation/accelerate/compressing-and-decompressing-files-with-stream-compression
// [Compressing and saving a string to the file system]: https://developer.apple.com/documentation/accelerate/compressing-and-saving-a-string-to-the-file-system
// [Compressing file system directories]: https://developer.apple.com/documentation/accelerate/compressing-file-system-directories
// [Compressing single files]: https://developer.apple.com/documentation/accelerate/compressing-single-files
// [Controlling vDSP operations with stride]: https://developer.apple.com/documentation/accelerate/controlling-vdsp-operations-with-stride
// [Converting bitmap data between Core Graphics images and vImage buffers]: https://developer.apple.com/documentation/accelerate/converting-bitmap-data-between-core-graphics-images-and-vimage-buffers
// [Converting color images to grayscale]: https://developer.apple.com/documentation/accelerate/converting-color-images-to-grayscale
// [Converting luminance and chrominance planes to an ARGB image]: https://developer.apple.com/documentation/accelerate/converting-luminance-and-chrominance-planes-to-an-argb-image
// [Core Video interoperability]: https://developer.apple.com/documentation/accelerate/core-video-interoperability
// [Creating a Core Graphics Image from a vImage Buffer]: https://developer.apple.com/documentation/accelerate/creating-a-core-graphics-image-from-a-vimage-buffer
// [Creating a sparse matrix from coordinate format arrays]: https://developer.apple.com/documentation/accelerate/creating-a-sparse-matrix-from-coordinate-format-arrays
// [Creating and Populating Buffers from Core Graphics Images]: https://developer.apple.com/documentation/accelerate/creating-and-populating-buffers-from-core-graphics-images
// [Creating sparse matrices]: https://developer.apple.com/documentation/accelerate/creating-sparse-matrices
// [Decompressing and extracting an archived directory]: https://developer.apple.com/documentation/accelerate/decompressing-and-extracting-an-archived-directory
// [Decompressing and parsing an archived string]: https://developer.apple.com/documentation/accelerate/decompressing-and-parsing-an-archived-string
// [Decompressing single files]: https://developer.apple.com/documentation/accelerate/decompressing-single-files
// [Discrete Cosine transforms]: https://developer.apple.com/documentation/accelerate/discrete-cosine-transforms
// [Discrete Fourier transforms]: https://developer.apple.com/documentation/accelerate/discrete-fourier-transforms
// [Enhancing image contrast with histogram manipulation]: https://developer.apple.com/documentation/accelerate/enhancing-image-contrast-with-histogram-manipulation
// [Fast Fourier transforms]: https://developer.apple.com/documentation/accelerate/fast-fourier-transforms
// [Finding an interpolating polynomial using the Vandermonde method]: https://developer.apple.com/documentation/accelerate/finding-an-interpolating-polynomial-using-the-vandermonde-method
// [Finding the component frequencies in a composite sine wave]: https://developer.apple.com/documentation/accelerate/finding-the-component-frequencies-in-a-composite-sine-wave
// [Finding the sharpest image in a sequence of captured images]: https://developer.apple.com/documentation/accelerate/finding-the-sharpest-image-in-a-sequence-of-captured-images
// [Halftone descreening with 2D fast Fourier transform]: https://developer.apple.com/documentation/accelerate/halftone-descreening-with-2d-fast-fourier-transform
// [Image shearing]: https://developer.apple.com/documentation/accelerate/image-shearing
// [Improving the quality of quantized images with dithering]: https://developer.apple.com/documentation/accelerate/improving-the-quality-of-quantized-images-with-dithering
// [Integrating vImage pixel buffers into a Core Image workflow]: https://developer.apple.com/documentation/accelerate/integrating-vimage-pixel-buffers-into-a-core-image-workflow
// [Optimizing image-processing performance]: https://developer.apple.com/documentation/accelerate/optimizing-image-processing-performance
// [Performing Fourier Transforms on Multiple Signals]: https://developer.apple.com/documentation/accelerate/performing-fourier-transforms-on-multiple-signals
// [Performing Fourier transforms on interleaved-complex data]: https://developer.apple.com/documentation/accelerate/performing-fourier-transforms-on-interleaved-complex-data
// [Reducing artifacts with custom resampling filters]: https://developer.apple.com/documentation/accelerate/reducing-artifacts-with-custom-resampling-filters
// [Reducing spectral leakage with windowing]: https://developer.apple.com/documentation/accelerate/reducing-spectral-leakage-with-windowing
// [Resampling a signal with decimation]: https://developer.apple.com/documentation/accelerate/resampling-a-signal-with-decimation
// [Resampling in vImage]: https://developer.apple.com/documentation/accelerate/resampling-in-vimage
// [Rotating a cube by transforming its vertices]: https://developer.apple.com/documentation/accelerate/rotating-a-cube-by-transforming-its-vertices
// [Signal extraction from noise]: https://developer.apple.com/documentation/accelerate/signal-extraction-from-noise
// [Solving systems of linear equations with LAPACK]: https://developer.apple.com/documentation/accelerate/solving-systems-of-linear-equations-with-lapack
// [Solving systems using direct methods]: https://developer.apple.com/documentation/accelerate/solving-systems-using-direct-methods
// [Solving systems using iterative methods]: https://developer.apple.com/documentation/accelerate/solving-systems-using-iterative-methods
// [Sparse Solvers]: https://developer.apple.com/documentation/accelerate/sparse-solvers-library
// [Specifying histograms with vImage]: https://developer.apple.com/documentation/accelerate/specifying-histograms-with-vimage
// [Training a neural network to recognize digits]: https://developer.apple.com/documentation/accelerate/training-a-neural-network-to-recognize-digits
// [Understanding data packing for Fourier transforms]: https://developer.apple.com/documentation/accelerate/understanding-data-packing-for-fourier-transforms
// [Using linear interpolation to construct new data points]: https://developer.apple.com/documentation/accelerate/using-linear-interpolation-to-construct-new-data-points
// [Using vDSP for vector-based arithmetic]: https://developer.apple.com/documentation/accelerate/using-vdsp-for-vector-based-arithmetic
// [Using vImage pixel buffers to generate video effects]: https://developer.apple.com/documentation/accelerate/using-vimage-pixel-buffers-to-generate-video-effects
// [Visualizing sound as an audio spectrogram]: https://developer.apple.com/documentation/accelerate/visualizing-sound-as-an-audio-spectrogram
// [Working with Matrices]: https://developer.apple.com/documentation/accelerate/working-with-matrices
// [Working with Quaternions]: https://developer.apple.com/documentation/accelerate/working-with-quaternions
// [Working with Vectors]: https://developer.apple.com/documentation/accelerate/working-with-vectors
package accelerate

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the Accelerate library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
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
	// Loading is best-effort: the warning is silent by default because a missing
	// framework is harmless unless one of its symbols is actually called. Set
	// APPLE_FRAMEWORK_LOAD_DEBUG to surface load failures while diagnosing.
	if os.Getenv("APPLE_FRAMEWORK_LOAD_DEBUG") != "" {
		fmt.Fprintf(os.Stderr, "warning: Accelerate: failed to load framework from any known path\n")
	}
}
