// Command imagefilter runs MPSImageGaussianBlur over a Metal texture.
//
// MPS image kernels are MPSUnaryImageKernel subclasses: they read a source
// MTLTexture and write a destination MTLTexture, encoded into a command
// buffer. The interesting part is not that a picture comes out blurrier, it is
// that the filter is a known linear operator, so the result can be checked
// numerically.
//
// Two properties are checked, both on a single-channel R32Float texture:
//
//   - Blurring an impulse — one pixel set to 1, everything else 0 — makes the
//     output the blur kernel itself. A correct Gaussian is non-negative,
//     symmetric about the impulse, peaked at the impulse, and sums to 1.
//   - Blurring a constant field leaves it unchanged, since the kernel is
//     normalized. This is checked away from the border, where the default
//     MPSImageEdgeModeZero pulls values down.
//
// Usage:
//
//	imagefilter
package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"unsafe"

	"github.com/tmc/apple/metal"
	mps "github.com/tmc/apple/metalperformanceshaders"
)

func init() {
	runtime.LockOSThread()
}

const (
	width  = 64
	height = 64
	sigma  = 3
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "imagefilter: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	device := metal.MTLCreateSystemDefaultDevice()
	if device.GetID() == 0 {
		return fmt.Errorf("no Metal device available; this example needs a Metal-capable GPU")
	}
	queue := device.NewCommandQueue()
	if queue.GetID() == 0 {
		return fmt.Errorf("could not create a Metal command queue")
	}
	fmt.Printf("device: %s\n", device.Name())

	blur := mps.NewImageGaussianBlurWithDeviceSigma(device, sigma)
	if blur.GetID() == 0 {
		return fmt.Errorf("could not create MPSImageGaussianBlur")
	}
	fmt.Printf("MPSImageGaussianBlur sigma = %g\n", blur.Sigma())

	// An impulse at the centre: the blurred output is the kernel itself.
	impulse := make([]float32, width*height)
	impulse[centre*width+centre] = 1
	kernel, err := filter(device, queue, blur, impulse)
	if err != nil {
		return fmt.Errorf("blur impulse: %w", err)
	}
	if err := checkKernel(kernel); err != nil {
		return err
	}

	// A constant field: a normalized kernel leaves the interior alone.
	constant := make([]float32, width*height)
	for i := range constant {
		constant[i] = 0.25
	}
	flat, err := filter(device, queue, blur, constant)
	if err != nil {
		return fmt.Errorf("blur constant field: %w", err)
	}
	if err := checkConstant(flat, 0.25); err != nil {
		return err
	}

	fmt.Println("PASS")
	return nil
}

// centre is the pixel the impulse is placed at.
const centre = width / 2

// checkKernel verifies that the blurred impulse is a normalized, symmetric,
// non-negative kernel peaked at the impulse.
func checkKernel(kernel []float32) error {
	sum := 0.0
	min := math.Inf(1)
	for _, v := range kernel {
		sum += float64(v)
		if float64(v) < min {
			min = float64(v)
		}
	}
	peak := kernel[centre*width+centre]

	// Symmetry: the kernel must agree with itself reflected through the
	// impulse in both axes.
	maxAsym := 0.0
	for y := 1; y < height; y++ {
		for x := 1; x < width; x++ {
			my, mx := 2*centre-y, 2*centre-x
			if my < 0 || my >= height || mx < 0 || mx >= width {
				continue
			}
			for _, d := range []float64{
				math.Abs(float64(kernel[y*width+x] - kernel[my*width+x])),
				math.Abs(float64(kernel[y*width+x] - kernel[y*width+mx])),
			} {
				if d > maxAsym {
					maxAsym = d
				}
			}
		}
	}

	fmt.Printf("impulse response: sum = %.6f, peak = %.6f at (%d,%d), min = %g, max asymmetry = %.4f%% of peak\n",
		sum, peak, centre, centre, min, 100*maxAsym/float64(peak))

	// The kernel is truncated to a finite radius, so the sum is close to but
	// not exactly 1.
	if math.Abs(sum-1) > 1e-3 {
		return fmt.Errorf("blur kernel sums to %g, want 1 within 1e-3", sum)
	}
	if min < 0 {
		return fmt.Errorf("blur kernel has a negative weight %g", min)
	}
	// MPSImageGaussianBlur approximates the Gaussian rather than evaluating
	// it, so the kernel is symmetric only to a fraction of a percent. The
	// bound is relative to the peak because that is the scale the weights
	// live on.
	if maxAsym > 0.02*float64(peak) {
		return fmt.Errorf("blur kernel is asymmetric by %g, more than 2%% of the peak %g", maxAsym, peak)
	}
	for i, v := range kernel {
		if v > peak {
			return fmt.Errorf("blur kernel peaks at index %d (%g), not at the impulse (%g)", i, v, peak)
		}
	}
	// A sigma-3 Gaussian spreads the impulse out; if the peak is still near 1
	// nothing was filtered.
	if peak > 0.1 {
		return fmt.Errorf("blur kernel peak is %g, too sharp for sigma %g", peak, float32(sigma))
	}
	return nil
}

// checkConstant verifies that the interior of a constant field survives the
// blur unchanged. The border is skipped: MPSImageEdgeModeZero, the default,
// samples zero outside the texture and so darkens the edges.
func checkConstant(got []float32, want float32) error {
	const border = 4 * sigma
	maxDiff := 0.0
	for y := border; y < height-border; y++ {
		for x := border; x < width-border; x++ {
			if d := math.Abs(float64(got[y*width+x] - want)); d > maxDiff {
				maxDiff = d
			}
		}
	}
	fmt.Printf("constant field: max |out-in| in the interior = %g\n", maxDiff)
	if maxDiff > 1e-5 {
		return fmt.Errorf("blurring a constant field changed the interior by %g", maxDiff)
	}
	return nil
}

// filter runs kernel over src and returns the destination pixels. src holds
// width*height single-channel float32 pixels.
func filter(device metal.MTLDeviceObject, queue metal.MTLCommandQueue, kernel mps.MPSImageGaussianBlur, src []float32) ([]float32, error) {
	source, err := newTexture(device)
	if err != nil {
		return nil, err
	}
	dest, err := newTexture(device)
	if err != nil {
		return nil, err
	}

	region := metal.MTLRegion{Size: metal.MTLSize{Width: width, Height: height, Depth: 1}}
	source.ReplaceRegionMipmapLevelWithBytesBytesPerRow(region, 0, unsafe.Pointer(&src[0]), width*4)
	runtime.KeepAlive(src)

	buf := queue.CommandBuffer()
	if buf.GetID() == 0 {
		return nil, fmt.Errorf("could not create a command buffer")
	}
	kernel.EncodeToCommandBufferSourceTextureDestinationTexture(buf, source, dest)
	buf.Commit()
	buf.WaitUntilCompleted()

	out := make([]float32, width*height)
	dest.GetBytesBytesPerRowFromRegionMipmapLevel(unsafe.Pointer(&out[0]), width*4, region, 0)
	return out, nil
}

// newTexture returns a shared-storage single-channel float texture that MPS
// can both read and write.
func newTexture(device metal.MTLDeviceObject) (metal.MTLTexture, error) {
	desc := metal.GetMTLTextureDescriptorClass().Texture2DDescriptorWithPixelFormatWidthHeightMipmapped(
		metal.MTLPixelFormatR32Float, width, height, false)
	if desc.GetID() == 0 {
		return nil, fmt.Errorf("could not create a texture descriptor")
	}
	desc.SetUsage(metal.MTLTextureUsageShaderRead | metal.MTLTextureUsageShaderWrite)
	desc.SetStorageMode(metal.MTLStorageModeShared)

	tex := device.NewTextureWithDescriptor(desc)
	if tex.GetID() == 0 {
		return nil, fmt.Errorf("could not create a %dx%d texture", width, height)
	}
	return tex, nil
}
