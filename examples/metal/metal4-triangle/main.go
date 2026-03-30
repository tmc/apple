// metal4-triangle renders a rotating triangle using the Metal 4 API.
//
// The renderer follows Apple's "Drawing a triangle with Metal 4" flow:
// - MTKView-driven frame callbacks
// - triple-buffered command allocators and vertex buffers
// - MTL4 argument table bindings for vertex and viewport buffers
// - queue submission via waitForDrawable -> commit:count: -> signalDrawable
package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"runtime"
	"sync"
	"time"
	"unsafe"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/metalkit"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/quartzcore"
)

const shaderSource = `
#include <metal_stdlib>
using namespace metal;

struct VertexData {
    float2 position;
    float4 color;
};

struct RasterizerData {
    float4 position [[position]];
    float4 color;
};

vertex RasterizerData vertexShader(
    uint vertexID [[vertex_id]],
    constant VertexData *vertexData [[buffer(0)]],
    constant uint2 *viewportSizePointer [[buffer(1)]]) {
    RasterizerData out;

    float2 pixelSpacePosition = vertexData[vertexID].position.xy;
    float2 viewportSize = float2(*viewportSizePointer);

    out.position.xy = pixelSpacePosition / (viewportSize / 2.0);
    out.position.z = 0.0;
    out.position.w = 1.0;
    out.color = vertexData[vertexID].color;

    return out;
}

fragment float4 fragmentShader(RasterizerData in [[stage_in]]) {
    return in.color;
}
`

const (
	mtlLoadActionClear       = metal.MTLLoadAction(2)
	mtlPixelFormatBGRA8Unorm = metal.MTLPixelFormat(80)
	mtlVertexFormatFloat2    = metal.MTLVertexFormat(29)
	mtlVertexFormatFloat4    = metal.MTLVertexFormat(31)
)

type vertexData struct {
	Position [2]float32
	_        [2]float32 // Pad to 16-byte alignment before float4 color.
	Color    [4]float32
}

type triangleData struct {
	Vertex0 vertexData
	Vertex1 vertexData
	Vertex2 vertexData
}

type viewportSize struct {
	Width  uint32
	Height uint32
}

const (
	windowWidth          = 800
	windowHeight         = 600
	maxFramesInFlight    = 3
	vertexBufferIndex    = 0
	viewportBufferIndex  = 1
	sharedEventTimeoutMS = 10
)

type triangleRenderer struct {
	device      metal.MTLDeviceObject
	queue       metal.MTL4CommandQueue
	commandBuff metal.MTL4CommandBuffer

	commandAllocators   [maxFramesInFlight]metal.MTL4CommandAllocator
	triangleVertexBufs  [maxFramesInFlight]metal.MTLBuffer
	viewportSizeBuffer  metal.MTLBuffer
	argumentTable       metal.MTL4ArgumentTable
	event               metal.MTLSharedEvent
	renderPipelineState metal.MTLRenderPipelineState

	frameNumber uint64
}

func main() {
	e2e := flag.Bool("e2e", false, "run non-interactive end-to-end check")
	flag.Parse()

	runtime.LockOSThread()

	device := metal.MTLCreateSystemDefaultDevice()
	if device.ID == 0 {
		fmt.Fprintln(os.Stderr, "metal4-triangle: no Metal device found")
		if *e2e {
			fmt.Println("e2e: skip no Metal device")
		}
		os.Exit(1)
	}

	fmt.Printf("device: %s\n", device.Name())

	if !device.SupportsFamily(metal.MTLGPUFamilyMetal4) {
		fmt.Fprintln(os.Stderr, "metal4-triangle: device does not support Metal 4")
		if *e2e {
			fmt.Println("e2e: skip no Metal 4 support")
		}
		os.Exit(1)
	}
	fmt.Println("Metal 4 supported")

	renderer := newTriangleRenderer(device)
	if *e2e {
		fmt.Println("e2e: pass")
		return
	}

	app := appkit.GetNSApplicationClass().SharedApplication()
	app.SetActivationPolicy(appkit.NSApplicationActivationPolicyRegular)
	app.SetMainMenu(buildMenuBar())

	window, view := createWindowAndView(device)
	viewDelegate := metalkit.NewMTKViewDelegate(metalkit.MTKViewDelegateConfig{
		DrawInMTKView: func(view metalkit.MTKView) {
			renderer.drawFrame(view)
		},
	})
	view.SetDelegate(viewDelegate)

	done := make(chan struct{})
	var stopOnce sync.Once
	stopRender := func() {
		stopOnce.Do(func() { close(done) })
	}

	ticker := time.NewTicker(time.Second / 60)
	defer ticker.Stop()
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				dispatch.MainQueue().Async(func() {
					view.Draw()
				})
			}
		}
	}()

	appDelegate := appkit.NewNSApplicationDelegate(appkit.NSApplicationDelegateConfig{
		ShouldTerminateAfterLastWindowClosed: func(_ appkit.NSApplication) bool {
			return true
		},
		WillTerminate: func(_ foundation.NSNotification) {
			stopRender()
		},
	})
	app.SetDelegate(appDelegate)

	window.Center()
	window.MakeKeyAndOrderFront(nil)
	app.Activate()

	fmt.Println("rendering rotating Metal 4 triangle (close window or press Cmd+Q to exit)")
	app.Run()

	stopRender()
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(viewDelegate)
	runtime.KeepAlive(appDelegate)
}

func newTriangleRenderer(device metal.MTLDeviceObject) *triangleRenderer {
	queue := device.NewMTL4CommandQueue()
	if queue == nil {
		fatal("failed to create MTL4 command queue")
	}

	commandBuffer := device.NewCommandBuffer()
	if commandBuffer == nil {
		fatal("failed to create MTL4 command buffer")
	}

	library := newLibraryWithSource(device, shaderSource)
	vertexDesc := createVertexDescriptor()

	compilerDesc := metal.NewMTL4CompilerDescriptor()
	compiler, err := device.NewCompilerWithDescriptorError(compilerDesc)
	if err != nil {
		fatal(fmt.Sprintf("failed to create compiler: %v", err))
	}

	pipelineState := createPipelineState(compiler, library, vertexDesc)

	argTableDesc := metal.NewMTL4ArgumentTableDescriptor()
	argTableDesc.SetMaxBufferBindCount(2)
	argumentTable, err := device.NewArgumentTableWithDescriptorError(argTableDesc)
	if err != nil {
		fatal(fmt.Sprintf("failed to create argument table: %v", err))
	}

	viewport := viewportSize{Width: windowWidth, Height: windowHeight}
	viewportBuffer := device.NewBufferWithBytesLengthOptions(
		unsafe.Pointer(&viewport),
		uint(unsafe.Sizeof(viewportSize{})),
		metal.MTLResourceStorageModeShared,
	)
	if viewportBuffer == nil {
		fatal("failed to create viewport size buffer")
	}

	var triangleBufs [maxFramesInFlight]metal.MTLBuffer
	for i := range triangleBufs {
		triangle := currentTriangleData(uint64(i))
		buf := device.NewBufferWithBytesLengthOptions(
			unsafe.Pointer(&triangle),
			uint(unsafe.Sizeof(triangle)),
			metal.MTLResourceStorageModeShared,
		)
		if buf == nil {
			fatal("failed to create triangle vertex buffer")
		}
		triangleBufs[i] = buf
	}

	var allocators [maxFramesInFlight]metal.MTL4CommandAllocator
	for i := range allocators {
		allocator := device.NewCommandAllocator()
		if allocator == nil {
			fatal("failed to create command allocator")
		}
		allocators[i] = allocator
	}

	event := device.NewSharedEvent()
	if event == nil {
		fatal("failed to create shared event")
	}
	event.SetSignaledValue(0)

	return &triangleRenderer{
		device:              device,
		queue:               queue,
		commandBuff:         commandBuffer,
		commandAllocators:   allocators,
		triangleVertexBufs:  triangleBufs,
		viewportSizeBuffer:  viewportBuffer,
		argumentTable:       argumentTable,
		event:               event,
		renderPipelineState: pipelineState,
	}
}

func (r *triangleRenderer) drawFrame(view metalkit.MTKView) {
	drawableObj := view.CurrentDrawable()
	if drawableObj == nil || drawableObj.GetID() == 0 {
		return
	}
	drawable := quartzcore.CAMetalDrawableObjectFromID(drawableObj.GetID())

	drawableTexture := drawable.Texture()
	if drawableTexture == nil || drawableTexture.GetID() == 0 {
		return
	}
	drawableWidth := drawableTexture.Width()
	drawableHeight := drawableTexture.Height()
	if drawableWidth == 0 || drawableHeight == 0 {
		return
	}

	renderPassDesc := view.CurrentMTL4RenderPassDescriptor()
	if renderPassDesc.ID == 0 {
		return
	}
	colorAttachment := renderPassDesc.ColorAttachments().ObjectAtIndexedSubscript(0)
	colorAttachment.SetLoadAction(mtlLoadActionClear)
	colorAttachment.SetStoreAction(metal.MTLStoreActionStore)
	colorAttachment.SetClearColor(metal.MTLClearColor{Red: 0, Green: 0, Blue: 0, Alpha: 1})

	r.frameNumber++
	frameIndex := int(r.frameNumber % maxFramesInFlight)
	if r.frameNumber > maxFramesInFlight {
		r.waitOnSharedEvent(r.frameNumber - maxFramesInFlight)
	}

	allocator := r.commandAllocators[frameIndex]
	allocator.Reset()

	r.commandBuff.BeginCommandBufferWithAllocator(allocator)
	label := fmt.Sprintf("Frame: %d", r.frameNumber)
	r.commandBuff.SetLabel(label)

	encoder := r.commandBuff.RenderCommandEncoderWithDescriptor(renderPassDesc)
	if encoder == nil {
		r.commandBuff.EndCommandBuffer()
		return
	}
	encoder.SetRenderPipelineState(r.renderPipelineState)

	r.updateViewportSize(uint32(drawableWidth), uint32(drawableHeight))
	r.updateTriangleVertexData(frameIndex)
	r.setRenderPassArguments(encoder, frameIndex)

	encoder.SetViewport(metal.MTLViewport{
		OriginX: 0,
		OriginY: 0,
		Width:   float64(drawableWidth),
		Height:  float64(drawableHeight),
		Znear:   0,
		Zfar:    1,
	})
	encoder.DrawPrimitivesVertexStartVertexCount(metal.MTLPrimitiveTypeTriangle, 0, 3)
	encoder.EndEncoding()

	r.commandBuff.EndCommandBuffer()
	r.submitCommandBuffer(drawable.MTLDrawableObject)
	r.queue.SignalEventValue(r.event, r.frameNumber)
	runtime.KeepAlive(drawable)
}

func (r *triangleRenderer) setRenderPassArguments(encoder metal.MTL4RenderCommandEncoder, frameIndex int) {
	vertexBuffer := r.triangleVertexBufs[frameIndex]
	r.argumentTable.SetAddressAtIndex(vertexBuffer.GpuAddress(), uint(vertexBufferIndex))
	r.argumentTable.SetAddressAtIndex(r.viewportSizeBuffer.GpuAddress(), uint(viewportBufferIndex))
	encoder.SetArgumentTableAtStages(r.argumentTable, metal.MTLRenderStageVertex)
}

func (r *triangleRenderer) submitCommandBuffer(drawable metal.MTLDrawableObject) {
	r.queue.WaitForDrawable(drawable)
	objc.Send[struct{}](r.queue.GetID(), objc.Sel("commit:count:"), objc.CArray([]metal.MTL4CommandBuffer{r.commandBuff}), uint(1))
	r.queue.SignalDrawable(drawable)
	drawable.Present()
}

func (r *triangleRenderer) waitOnSharedEvent(earlierFrame uint64) {
	if r.event.WaitUntilSignaledValueTimeoutMS(earlierFrame, sharedEventTimeoutMS) {
		return
	}
	fmt.Fprintf(os.Stderr, "metal4-triangle: no signal from frame %d after %dms\n", earlierFrame, sharedEventTimeoutMS)
}

func (r *triangleRenderer) updateViewportSize(width, height uint32) {
	viewport := viewportSize{Width: width, Height: height}
	dst := r.viewportSizeBuffer.Contents()
	if dst != nil {
		copyBytes(dst, unsafe.Pointer(&viewport), unsafe.Sizeof(viewport))
	}
}

func (r *triangleRenderer) updateTriangleVertexData(frameIndex int) {
	data := currentTriangleData(r.frameNumber)
	dst := r.triangleVertexBufs[frameIndex].Contents()
	if dst != nil {
		copyBytes(dst, unsafe.Pointer(&data), unsafe.Sizeof(data))
	}
}

// copyBytes copies n bytes from src to dst.
func copyBytes(dst, src unsafe.Pointer, n uintptr) {
	// Use a byte-slice overlay to let the runtime handle the copy.
	d := unsafe.Slice((*byte)(dst), n)
	s := unsafe.Slice((*byte)(src), n)
	copy(d, s)
}

func createWindowAndView(device metal.MTLDeviceObject) (appkit.NSWindow, metalkit.MTKView) {
	windowFrame := corefoundation.CGRect{
		Origin: corefoundation.CGPoint{X: 200, Y: 200},
		Size:   corefoundation.CGSize{Width: windowWidth, Height: windowHeight},
	}
	viewFrame := corefoundation.CGRect{
		Origin: corefoundation.CGPoint{X: 200, Y: 200},
		Size:   corefoundation.CGSize{Width: windowWidth, Height: windowHeight},
	}

	window := appkit.NewWindowWithContentRectStyleMaskBackingDefer(
		windowFrame,
		appkit.NSWindowStyleMaskTitled|
			appkit.NSWindowStyleMaskClosable|
			appkit.NSWindowStyleMaskMiniaturizable|
			appkit.NSWindowStyleMaskResizable,
		appkit.NSBackingStoreBuffered,
		false,
	)
	window.SetTitle("Metal 4 Triangle")

	view := metalkit.NewMTKView().InitWithFrameDevice(viewFrame, device)
	view.SetColorPixelFormat(mtlPixelFormatBGRA8Unorm)
	view.SetClearColor(metal.MTLClearColor{Red: 0, Green: 0, Blue: 0, Alpha: 1})
	view.SetPreferredFramesPerSecond(60)
	view.SetPaused(true)
	view.SetEnableSetNeedsDisplay(true)

	window.SetContentView(view)
	return window, view
}

func buildMenuBar() appkit.NSMenu {
	menuBar := appkit.NewNSMenu()

	appMenuItem := appkit.NewMenuItemWithTitleActionKeyEquivalent("", 0, "")
	appMenu := appkit.NewNSMenu()
	appMenu.AddItemWithTitleActionKeyEquivalent("Quit", objc.Sel("terminate:"), "q")
	appMenuItem.SetSubmenu(appMenu)
	menuBar.AddItem(appMenuItem)

	return menuBar
}

func newLibraryWithSource(device metal.MTLDeviceObject, source string) metal.MTLLibrary {
	lib, err := device.NewLibraryWithSourceOptionsError(source, nil)
	if err != nil {
		fatal(fmt.Sprintf("shader compilation failed: %v", err))
	}
	return lib
}

func createVertexDescriptor() metal.MTLVertexDescriptor {
	desc := metal.GetMTLVertexDescriptorClass().VertexDescriptor()

	attrs := desc.Attributes()
	attr0 := attrs.ObjectAtIndexedSubscript(0)
	attr0.SetFormat(mtlVertexFormatFloat2)
	attr0.SetOffset(0)
	attr0.SetBufferIndex(0)

	attr1 := attrs.ObjectAtIndexedSubscript(1)
	attr1.SetFormat(mtlVertexFormatFloat4)
	attr1.SetOffset(8)
	attr1.SetBufferIndex(0)

	layouts := desc.Layouts()
	layout0 := layouts.ObjectAtIndexedSubscript(0)
	layout0.SetStride(uint(unsafe.Sizeof(vertexData{})))
	layout0.SetStepFunction(metal.MTLVertexStepFunctionPerVertex)

	return desc
}

func createPipelineState(
	compiler metal.MTL4Compiler,
	library metal.MTLLibrary,
	vertexDesc metal.MTLVertexDescriptor,
) metal.MTLRenderPipelineState {
	vertFuncDesc := metal.NewMTL4LibraryFunctionDescriptor()
	vertFuncDesc.SetLibrary(library)
	vertFuncDesc.SetName("vertexShader")

	fragFuncDesc := metal.NewMTL4LibraryFunctionDescriptor()
	fragFuncDesc.SetLibrary(library)
	fragFuncDesc.SetName("fragmentShader")

	pipeDesc := metal.NewMTL4RenderPipelineDescriptor()
	pipeDesc.SetVertexFunctionDescriptor(vertFuncDesc)
	pipeDesc.SetFragmentFunctionDescriptor(fragFuncDesc)
	pipeDesc.SetVertexDescriptor(vertexDesc)
	pipeDesc.ColorAttachments().ObjectAtIndexedSubscript(0).SetPixelFormat(mtlPixelFormatBGRA8Unorm)

	pso, err := compiler.NewRenderPipelineStateWithDescriptorCompilerTaskOptionsError(pipeDesc, nil)
	if err != nil {
		fatal(fmt.Sprintf("pipeline state creation failed: %v", err))
	}
	return pso
}

func currentTriangleData(frameNumber uint64) triangleData {
	const radius = float32(180)
	angle0 := float32(frameNumber%360) * (float32(math.Pi) / 180)
	angle1 := angle0 + (2 * float32(math.Pi) / 3)
	angle2 := angle0 + (4 * float32(math.Pi) / 3)

	return triangleData{
		Vertex0: vertexData{
			Position: [2]float32{radius * float32(math.Cos(float64(angle0))), radius * float32(math.Sin(float64(angle0)))},
			Color:    [4]float32{1, 0, 0, 1},
		},
		Vertex1: vertexData{
			Position: [2]float32{radius * float32(math.Cos(float64(angle1))), radius * float32(math.Sin(float64(angle1)))},
			Color:    [4]float32{0, 1, 0, 1},
		},
		Vertex2: vertexData{
			Position: [2]float32{radius * float32(math.Cos(float64(angle2))), radius * float32(math.Sin(float64(angle2)))},
			Color:    [4]float32{0, 0, 1, 1},
		},
	}
}

func fatal(msg string) {
	fmt.Fprintf(os.Stderr, "metal4-triangle: %s\n", msg)
	os.Exit(1)
}
