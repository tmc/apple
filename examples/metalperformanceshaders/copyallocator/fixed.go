//go:build mpsfixedcopyallocator

// This file is the program that main.go's doc comment says cannot be written
// yet. It is the full copy-allocator test, including assertions 3 and 4, as
// it will read once the generator emits the corrected signatures:
//
//	type MTLTextureMPSKernelMTLCommandBufferMTLTextureHandler = func(*MPSKernel, metal.MTLCommandBuffer, metal.MTLTexture) metal.MTLTexture
//
//	func (u MPSUnaryImageKernel) EncodeToCommandBufferInPlaceTextureFallbackCopyAllocator(
//	    commandBuffer metal.MTLCommandBuffer,
//	    texture *metal.MTLTexture,
//	    copyAllocator MTLTextureMPSKernelMTLCommandBufferMTLTextureHandler) bool
//
// It does not compile against the shipped API — that is the point — and is
// therefore behind a build tag so the package still builds:
//
//	go run -tags mpsfixedcopyallocator ./examples/metalperformanceshaders/copyallocator
//
// The exact spelling of the fixed signature is a prediction. If the
// generator lands a different but equally correct shape (for example an
// objc.ID return, or an out-param wrapper type), adjust this file to match;
// what must not change is what it asserts.
//
// No child process is needed here: a correct binding does not crash, so the
// failure mode is an assertion, not a signal.
package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	mps "github.com/tmc/apple/metalperformanceshaders"
	"github.com/tmc/apple/objc"
)

func init() {
	runtime.LockOSThread()
}

const (
	width  = 64
	height = 64
	kernel = 7
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "copyallocator: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("PASS")
}

func run() error {
	device := metal.MTLCreateSystemDefaultDevice()
	if device.GetID() == 0 {
		return fmt.Errorf("no Metal device available")
	}
	queue := device.NewCommandQueue()
	if queue.GetID() == 0 {
		return fmt.Errorf("could not create a command queue")
	}
	filter := mps.NewImageAreaMaxWithDeviceKernelWidthKernelHeight(device, kernel, kernel)
	if filter.GetID() == 0 {
		return fmt.Errorf("could not create MPSImageAreaMax")
	}

	texture, err := newTexture(device)
	if err != nil {
		return err
	}
	original := texture.GetID()

	src := make([]float32, width*height)
	src[(height/2)*width+width/2] = 1
	region := metal.MTLRegion{Size: metal.MTLSize{Width: width, Height: height, Depth: 1}}
	texture.ReplaceRegionMipmapLevelWithBytesBytesPerRow(region, 0, unsafe.Pointer(&src[0]), width*4)
	runtime.KeepAlive(src)

	buf := queue.CommandBuffer()
	if buf.GetID() == 0 {
		return fmt.Errorf("could not create a command buffer")
	}

	var (
		invoked   int
		argsValid bool
		allocated objc.ID
		argErr    error
	)
	allocator := func(k *mps.MPSKernel, cb metal.MTLCommandBuffer, source metal.MTLTexture) metal.MTLTexture {
		// Assertion 1.
		invoked++

		// Assertion 2: message each argument rather than nil-checking it.
		var kernelID objc.ID
		if k != nil {
			kernelID = k.ID
		}
		argsValid = true
		for _, a := range []struct {
			name string
			id   objc.ID
		}{
			{"filter", kernelID},
			{"commandBuffer", cb.GetID()},
			{"sourceTexture", source.GetID()},
		} {
			if a.id == 0 {
				argErr = fmt.Errorf("allocator argument %s is nil", a.name)
				argsValid = false
				continue
			}
			fmt.Printf("allocator argument %s: class=%s description=%s\n",
				a.name, describe(objc.Send[objc.ID](a.id, objc.Sel("class"))), describe(a.id))
		}

		replacement, err := newTexture(device)
		if err != nil {
			argErr = err
			argsValid = false
			return metal.MTLTextureObjectFromID(0)
		}
		allocated = replacement.GetID()
		// NS_RETURNS_RETAINED: MPS takes ownership of a +1 texture and
		// releases the one it replaces.
		objc.Send[objc.ID](allocated, objc.Sel("retain"))
		return replacement
	}

	ok := filter.EncodeToCommandBufferInPlaceTextureFallbackCopyAllocator(buf, &texture, allocator)
	buf.Commit()
	buf.WaitUntilCompleted()

	if invoked == 0 {
		return fmt.Errorf("the fallback copy allocator was never invoked: MPSImageAreaMax %dx%d succeeded in place, so this run proves nothing; pick a kernel that cannot filter in place", kernel, kernel)
	}
	if !argsValid {
		return fmt.Errorf("allocator arguments were not valid: %v", argErr)
	}
	// Assertion 4a.
	if !ok {
		return fmt.Errorf("encodeToCommandBuffer:inPlaceTexture:fallbackCopyAllocator: returned false")
	}
	// Assertion 3: MPS must have written the allocator's texture back
	// through the out-parameter, replacing the caller's variable.
	if texture.GetID() != allocated {
		return fmt.Errorf("in-place texture was not replaced by the allocator's texture: have %#x, allocator returned %#x, original was %#x",
			texture.GetID(), allocated, original)
	}
	fmt.Printf("in-place texture replaced: %#x -> %#x (the allocator's texture)\n", original, allocated)

	// Assertion 4b: the replacement holds the filtered image. A 7x7 AreaMax
	// of an impulse is a 7x7 block of the peak value centred on the impulse.
	out := make([]float32, width*height)
	texture.GetBytesBytesPerRowFromRegionMipmapLevel(unsafe.Pointer(&out[0]), width*4, region, 0)
	if err := checkAreaMax(out); err != nil {
		return err
	}

	// Assertion 5 (NS_RETURNS_RETAINED refcount correctness) is NOT checked.
	// -retainCount is not a sound instrument under ARC and MPS may retain
	// the texture internally; the honest proof is a leaks run against a
	// fixed binding, which does not belong inside this program.
	fmt.Println("note: NS_RETURNS_RETAINED refcount balance is not checked here; verify with leaks(1)")
	return nil
}

// checkAreaMax verifies that a 7x7 AreaMax of a unit impulse produced a 7x7
// plateau of 1 centred on the impulse and zero elsewhere.
func checkAreaMax(out []float32) error {
	const r = kernel / 2
	cy, cx := height/2, width/2
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			want := float32(0)
			if abs(y-cy) <= r && abs(x-cx) <= r {
				want = 1
			}
			if math.Abs(float64(out[y*width+x]-want)) > 1e-6 {
				return fmt.Errorf("filtered output at (%d,%d) = %g, want %g", x, y, out[y*width+x], want)
			}
		}
	}
	fmt.Printf("filtered output: %dx%d plateau of 1 centred at (%d,%d), zero elsewhere\n", kernel, kernel, cx, cy)
	return nil
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func describe(id objc.ID) string {
	if id == 0 {
		return "<nil>"
	}
	s := objc.Send[objc.ID](id, objc.Sel("description"))
	if s == 0 {
		return "<no description>"
	}
	return foundation.NSStringFromID(s).UTF8String()
}

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
