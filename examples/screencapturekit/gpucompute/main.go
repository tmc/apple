// Command gpucompute applies a Metal luminance reduction to live
// ScreenCaptureKit frames. The shader reads each frame's IOSurface directly;
// no CPU pixel copy is made between capture and compute.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/iosurface"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/screencapturekit"
)

const shader = `
#include <metal_stdlib>
using namespace metal;
kernel void luminance(texture2d<float, access::read> image [[texture(0)]],
                      device atomic_uint *sum [[buffer(0)]],
                      uint2 p [[thread_position_in_grid]]) {
    if (p.x >= image.get_width() || p.y >= image.get_height()) return;
    float4 c = image.read(p);
    atomic_fetch_add_explicit(sum, uint(c.r * 255.0 + c.g * 255.0 + c.b * 255.0),
                              memory_order_relaxed);
}
`

type reducer struct {
	device   metal.MTLDeviceObject
	queue    metal.MTLCommandQueue
	pipeline metal.MTLComputePipelineState
	sum      metal.MTLBuffer
	mu       sync.Mutex
}

func newReducer() *reducer {
	d := metal.MTLCreateSystemDefaultDevice()
	if d.GetID() == 0 {
		log.Fatal("gpucompute: no Metal device")
	}
	lib, err := d.NewLibraryWithSourceOptionsError(shader, nil)
	if err != nil {
		log.Fatalf("gpucompute: compile shader: %v", err)
	}
	fn := lib.NewFunctionWithName("luminance")
	pipeline, err := d.NewComputePipelineStateWithFunctionError(fn)
	if err != nil {
		log.Fatalf("gpucompute: create pipeline: %v", err)
	}
	sum := d.NewBufferWithLengthOptions(4, metal.MTLResourceStorageModeShared)
	if sum.GetID() == 0 {
		log.Fatal("gpucompute: allocate result buffer")
	}
	return &reducer{device: d, queue: d.NewCommandQueue(), pipeline: pipeline, sum: sum}
}

func (r *reducer) reduce(buffer corevideo.CVImageBufferRef) (uint32, time.Duration, bool) {
	// CVImageBufferRef and CVPixelBufferRef have the same Core Video handle;
	// ScreenCaptureKit supplies a pixel buffer for screen frames.
	pixel := corevideo.CVPixelBufferRef(buffer)
	surface := iosurface.IOSurfaceRef(uintptr(corevideo.CVPixelBufferGetIOSurface(pixel)))
	if surface == 0 {
		return 0, 0, false
	}
	w, h := uint(iosurface.IOSurfaceGetWidth(surface)), uint(iosurface.IOSurfaceGetHeight(surface))
	if w == 0 || h == 0 {
		return 0, 0, false
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	desc := metal.GetMTLTextureDescriptorClass().Texture2DDescriptorWithPixelFormatWidthHeightMipmapped(
		metal.MTLPixelFormatBGRA8Unorm, w, h, false)
	desc.SetUsage(metal.MTLTextureUsageShaderRead)
	desc.SetStorageMode(metal.MTLStorageModeShared)
	tex := r.device.NewTextureWithDescriptorIosurfacePlane(desc, surface, 0)
	if tex.GetID() == 0 {
		return 0, 0, false
	}
	*(*uint32)(r.sum.Contents()) = 0
	cb := r.queue.CommandBuffer()
	enc := cb.ComputeCommandEncoder()
	enc.SetComputePipelineState(r.pipeline)
	enc.SetTextureAtIndex(tex, 0)
	enc.SetBufferWithOffsetAtIndex(r.sum, 0, 0)
	tg := r.pipeline.MaxTotalThreadsPerThreadgroup()
	if tg > 256 {
		tg = 256
	}
	enc.DispatchThreadsThreadsPerThreadgroup(metal.MTLSize{Width: w, Height: h, Depth: 1}, metal.MTLSize{Width: tg, Height: 1, Depth: 1})
	enc.EndEncoding()
	cb.Commit()
	cb.WaitUntilCompleted()
	d := time.Duration((cb.GPUEndTime() - cb.GPUStartTime()) * float64(time.Second))
	return *(*uint32)(r.sum.Contents()), d, true
}

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	duration := flag.Duration("duration", 5*time.Second, "maximum capture duration")
	fps := flag.Int("fps", 30, "capture and process rate")
	flag.Parse()
	if *duration <= 0 || *fps <= 0 {
		log.Fatal("gpucompute: duration and fps must be positive")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	content, err := screencapturekit.GetSCShareableContentClass().GetShareableContent(ctx)
	if err != nil || len(content.Displays()) == 0 {
		log.Fatal("gpucompute: screen recording permission or display unavailable")
	}
	display := content.Displays()[0]
	filter := screencapturekit.NewContentFilterWithDisplayExcludingWindows(display, nil)
	config := screencapturekit.NewSCStreamConfiguration()
	config.SetWidth(uintptr(display.Width()))
	config.SetHeight(uintptr(display.Height()))
	config.SetQueueDepth(3)
	config.SetCapturesAudio(false)
	config.SetMinimumFrameInterval(coremedia.CMTimeMake(1, int32(*fps)))

	r := newReducer()
	var frames atomic.Uint64
	var valid atomic.Uint64
	var totalGPU atomic.Int64
	output := screencapturekit.NewSCStreamOutput(screencapturekit.SCStreamOutputConfig{
		StreamDidOutputSampleBufferOfType: func(_ screencapturekit.SCStream, sample coremedia.CMSampleBufferRef, typ screencapturekit.SCStreamOutputType) {
			if typ != screencapturekit.SCStreamOutputTypeScreen {
				return
			}
			frames.Add(1)
			if sum, gpuTime, ok := r.reduce(coremedia.CMSampleBufferGetImageBuffer(sample)); ok {
				valid.Add(1)
				totalGPU.Add(gpuTime.Nanoseconds())
				if valid.Load() <= 5 {
					fmt.Printf("frame=%d luminance=%d gpu=%s\n", valid.Load(), sum, gpuTime.Round(time.Microsecond))
				}
			}
		},
	})
	delegate := screencapturekit.NewSCStreamDelegate(screencapturekit.SCStreamDelegateConfig{})
	stream := screencapturekit.NewStreamWithFilterConfigurationDelegate(filter, config, delegate)
	queue := dispatch.QueueCreate("github.com.tmc/apple.examples.screencapturekit.gpucompute")
	if _, err := stream.AddStreamOutputTypeSampleHandlerQueueError(output, screencapturekit.SCStreamOutputTypeScreen, queue); err != nil {
		log.Fatalf("gpucompute: add output: %v", err)
	}
	if err := stream.StartCapture(ctx); err != nil {
		log.Fatalf("gpucompute: start capture: %v", err)
	}
	timer := time.NewTimer(*duration)
	<-timer.C
	stopCtx, stopCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer stopCancel()
	if err := stream.StopCapture(stopCtx); err != nil {
		log.Fatalf("gpucompute: stop capture: %v", err)
	}
	if _, err := stream.RemoveStreamOutputTypeError(output, screencapturekit.SCStreamOutputTypeScreen); err != nil {
		log.Fatalf("gpucompute: remove output: %v", err)
	}
	n := valid.Load()
	avg := time.Duration(0)
	if n != 0 {
		avg = time.Duration(totalGPU.Load() / int64(n))
	}
	fmt.Printf("captured=%d computed=%d average_gpu=%s frame_budget=%s\n", frames.Load(), n, avg.Round(time.Microsecond), (time.Second / time.Duration(*fps)).Round(time.Microsecond))
	runtime.KeepAlive(output)
	runtime.KeepAlive(delegate)
	runtime.KeepAlive(stream)
	_ = os.Stdout
}
