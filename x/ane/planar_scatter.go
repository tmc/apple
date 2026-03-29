//go:build darwin

package ane

import (
	_ "embed"
	"fmt"
	"unsafe"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objectivec"
)

//go:embed planar_scatter.metal
var planarScatterSource string

// ScatterParams matches the Metal-side ScatterParams layout.
type ScatterParams struct {
	Channels    uint32
	Width       uint32
	Height      uint32
	ElemSize    uint32
	RowStride   uint32
	PlaneStride uint32
	Count       uint32
}

// ScatterParamsForSurface returns scatter parameters for a layout-backed surface.
func ScatterParamsForSurface(surf *IOSurfaceFloat32) (ScatterParams, bool) {
	if surf == nil {
		return ScatterParams{}, false
	}
	layout, ok := surf.Layout()
	if !ok {
		return ScatterParams{}, false
	}
	return ScatterParams{
		Channels:    uint32(layout.Channels),
		Width:       uint32(layout.Width),
		Height:      uint32(layout.Height),
		ElemSize:    uint32(layout.ElemSize),
		RowStride:   uint32(layout.RowStride),
		PlaneStride: uint32(layout.PlaneStride),
		Count:       uint32(layout.LogicalElements()),
	}, true
}

// PlanarScatterPipeline scatters contiguous float32 data into planar ANE layout.
type PlanarScatterPipeline struct {
	device   metal.MTLDeviceObject
	pipeline metal.MTLComputePipelineState
	queue    metal.MTLCommandQueue
	library  metal.MTLLibrary
}

// NewPlanarScatterPipeline compiles the scatter shader on the given device.
func NewPlanarScatterPipeline(device *MetalDevice) (*PlanarScatterPipeline, error) {
	if device == nil || device.device.GetID() == 0 {
		return nil, fmt.Errorf("ane: planar scatter device is nil")
	}

	options := metal.GetMTLCompileOptionsClass().Alloc().Init()
	defer objectivec.ObjectFromID(options.GetID()).Release()

	library, err := device.device.NewLibraryWithSourceOptionsError(planarScatterSource, options)
	if err != nil {
		return nil, fmt.Errorf("ane: compile planar scatter shader: %w", err)
	}
	fn := library.NewFunctionWithName("scatter_contiguous_to_planar")
	if fn == nil || fn.GetID() == 0 {
		objectivec.ObjectFromID(library.GetID()).Release()
		return nil, fmt.Errorf("ane: planar scatter function not found")
	}
	defer objectivec.ObjectFromID(fn.GetID()).Release()

	pipeline, err := device.device.NewComputePipelineStateWithFunctionError(fn)
	if err != nil {
		objectivec.ObjectFromID(library.GetID()).Release()
		return nil, fmt.Errorf("ane: create planar scatter pipeline: %w", err)
	}
	queue := device.device.NewCommandQueue()
	if queue == nil || queue.GetID() == 0 {
		objectivec.ObjectFromID(pipeline.GetID()).Release()
		objectivec.ObjectFromID(library.GetID()).Release()
		return nil, fmt.Errorf("ane: create planar scatter command queue failed")
	}

	return &PlanarScatterPipeline{
		device:   device.device,
		pipeline: pipeline,
		queue:    queue,
		library:  library,
	}, nil
}

func (p *PlanarScatterPipeline) dispatch(srcBuffer, dstBuffer metal.MTLBuffer, params ScatterParams) (metal.MTLCommandBuffer, error) {
	if p == nil {
		return nil, fmt.Errorf("ane: planar scatter pipeline is nil")
	}
	if srcBuffer == nil || srcBuffer.GetID() == 0 {
		return nil, fmt.Errorf("ane: planar scatter source buffer is nil")
	}
	if dstBuffer == nil || dstBuffer.GetID() == 0 {
		return nil, fmt.Errorf("ane: planar scatter destination buffer is nil")
	}
	if params.Count == 0 {
		return nil, fmt.Errorf("ane: planar scatter count is zero")
	}

	cmdBuf := p.queue.CommandBuffer()
	if cmdBuf == nil || cmdBuf.GetID() == 0 {
		return nil, fmt.Errorf("ane: create planar scatter command buffer failed")
	}
	encoder := cmdBuf.ComputeCommandEncoder()
	if encoder == nil || encoder.GetID() == 0 {
		return nil, fmt.Errorf("ane: create planar scatter compute encoder failed")
	}
	defer objectivec.ObjectFromID(encoder.GetID()).Release()

	encoder.SetComputePipelineState(p.pipeline)
	encoder.SetBufferWithOffsetAtIndex(srcBuffer, 0, 0)
	encoder.SetBufferWithOffsetAtIndex(dstBuffer, 0, 1)

	paramBytes := unsafe.Slice((*byte)(unsafe.Pointer(&params)), unsafe.Sizeof(params))
	encoder.SetBytesLengthAtIndex(paramBytes, 2)

	threadgroupWidth := p.pipeline.MaxTotalThreadsPerThreadgroup()
	if threadgroupWidth > uint(params.Count) {
		threadgroupWidth = uint(params.Count)
	}
	grid := metal.MTLSize{Width: uint(params.Count), Height: 1, Depth: 1}
	threads := metal.MTLSize{Width: threadgroupWidth, Height: 1, Depth: 1}
	encoder.DispatchThreadsThreadsPerThreadgroup(grid, threads)
	encoder.EndEncoding()

	return cmdBuf, nil
}

// ScatterToBuffers performs a blocking scatter from srcBuffer into dstBuffer.
func (p *PlanarScatterPipeline) ScatterToBuffers(srcBuffer, dstBuffer metal.MTLBuffer, params ScatterParams) error {
	cmdBuf, err := p.dispatch(srcBuffer, dstBuffer, params)
	if err != nil {
		return err
	}
	cmdBuf.Commit()
	cmdBuf.WaitUntilCompleted()
	return nil
}

// ScatterToBuffersSignalEvent performs a scatter and signals a Metal shared event.
func (p *PlanarScatterPipeline) ScatterToBuffersSignalEvent(srcBuffer, dstBuffer metal.MTLBuffer, params ScatterParams, event metal.MTLSharedEvent, value uint64) error {
	if event == nil || event.GetID() == 0 {
		return fmt.Errorf("ane: planar scatter shared event is nil")
	}
	cmdBuf, err := p.dispatch(srcBuffer, dstBuffer, params)
	if err != nil {
		return err
	}
	cmdBuf.EncodeSignalEventValue(event, value)
	cmdBuf.Commit()
	return nil
}

// Close releases the pipeline state, queue, and library.
func (p *PlanarScatterPipeline) Close() error {
	if p == nil {
		return nil
	}
	if p.queue != nil {
		objectivec.ObjectFromID(p.queue.GetID()).Release()
		p.queue = nil
	}
	if p.pipeline != nil {
		objectivec.ObjectFromID(p.pipeline.GetID()).Release()
		p.pipeline = nil
	}
	if p.library != nil {
		objectivec.ObjectFromID(p.library.GetID()).Release()
		p.library = nil
	}
	return nil
}
