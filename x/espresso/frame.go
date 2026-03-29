//go:build darwin

package espresso

import (
	"fmt"
	"unsafe"

	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/iosurface"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/private/espresso"
)

// Frame holds named input and output tensors for network execution.
type Frame struct {
	df espresso.EspressoDataFrame

	// inputs tracks pinned Go slices so they aren't GC'd during execution.
	inputs map[string][]byte
}

// NewFrame creates an empty data frame.
func NewFrame() *Frame {
	return &Frame{
		df:     espresso.NewEspressoDataFrame(),
		inputs: make(map[string][]byte),
	}
}

// SetInput attaches a named float32 input tensor.
func (f *Frame) SetInput(name string, data []float32, shape Shape) error {
	if shape.Elements() == 0 {
		return fmt.Errorf("espresso set input %q: empty shape", name)
	}
	if len(data) < shape.Elements() {
		return fmt.Errorf("espresso set input %q: data length %d < shape elements %d", name, len(data), shape.Elements())
	}

	att := espresso.NewEspressoDataFrameTensorAttachment()
	raw := unsafe.Pointer(unsafe.SliceData(data))
	att.SetRawPointer(raw)
	att.SetSize(uint64(len(data)) * 4) // float32 = 4 bytes

	// Pin the data to prevent GC.
	buf := unsafe.Slice((*byte)(raw), len(data)*4)
	f.inputs[name] = buf

	dict := f.inputDict()
	dict.SetObjectForKey(att, foundation.NSStringFromID(objc.String(name)))
	f.df.SetInputAttachments(dict)
	return nil
}

// SetInputBytes attaches a named raw byte input.
func (f *Frame) SetInputBytes(name string, data []byte) error {
	att := espresso.NewEspressoDataFrameTensorAttachment()
	att.SetRawPointer(unsafe.Pointer(unsafe.SliceData(data)))
	att.SetSize(uint64(len(data)))
	f.inputs[name] = data

	dict := f.inputDict()
	dict.SetObjectForKey(att, foundation.NSStringFromID(objc.String(name)))
	f.df.SetInputAttachments(dict)
	return nil
}

// SetInputIOSurface attaches a named input backed by an IOSurface.
// The IOSurface base address is used as the tensor's raw pointer.
func (f *Frame) SetInputIOSurface(name string, ref coregraphics.IOSurfaceRef) error {
	surfRef := iosurface.IOSurfaceRef(ref)
	base := iosurface.IOSurfaceGetBaseAddress(surfRef)
	if base == nil {
		return fmt.Errorf("espresso set input %q: IOSurface base address is nil", name)
	}
	size := iosurface.IOSurfaceGetAllocSize(surfRef)

	att := espresso.NewEspressoDataFrameTensorAttachment()
	att.SetRawPointer(base)
	att.SetSize(uint64(size))

	dict := f.inputDict()
	dict.SetObjectForKey(att, foundation.NSStringFromID(objc.String(name)))
	f.df.SetInputAttachments(dict)
	return nil
}

// SetInputANESurface attaches a named input from an Espresso ANE surface frame.
func (f *Frame) SetInputANESurface(name string, surface *ANESurface, frame uint64) error {
	if surface == nil {
		return fmt.Errorf("espresso set input %q: ANESurface is nil", name)
	}
	ref, err := surface.IOSurfaceForFrame(frame)
	if err != nil {
		return err
	}
	return f.SetInputIOSurface(name, ref)
}

// SetInputCVPixelBuffer attaches a named input from a CVPixelBuffer.
// Uses EspressoDataFrameTensorAttachment.CopyFromCVPixelBuffer.
func (f *Frame) SetInputCVPixelBuffer(name string, buf corevideo.CVImageBufferRef) error {
	cls := espresso.GetEspressoDataFrameTensorAttachmentClass()
	att := cls.CopyFromCVPixelBuffer(buf)
	if att == nil || att.GetID() == 0 {
		return fmt.Errorf("espresso set input %q: failed to copy from CVPixelBuffer", name)
	}

	dict := f.inputDict()
	dict.SetObjectForKey(att, foundation.NSStringFromID(objc.String(name)))
	f.df.SetInputAttachments(dict)
	return nil
}

// SetInputImage attaches a named image input with raw pixel data and channel count.
func (f *Frame) SetInputImage(name string, data []byte, nChannels int) error {
	att := espresso.NewEspressoDataFrameImageAttachment()
	att.SetNChannels(nChannels)
	raw := unsafe.Pointer(unsafe.SliceData(data))
	att.SetRawPointer(raw)
	att.SetSize(uint64(len(data)))
	f.inputs[name] = data

	dict := f.inputDict()
	dict.SetObjectForKey(att, foundation.NSStringFromID(objc.String(name)))
	f.df.SetInputAttachments(dict)
	return nil
}

// SetGroundTruth attaches a named ground truth tensor for training or evaluation.
func (f *Frame) SetGroundTruth(name string, data []float32, shape Shape) error {
	if shape.Elements() == 0 {
		return fmt.Errorf("espresso set ground truth %q: empty shape", name)
	}

	att := espresso.NewEspressoDataFrameTensorAttachment()
	raw := unsafe.Pointer(unsafe.SliceData(data))
	att.SetRawPointer(raw)
	att.SetSize(uint64(len(data)) * 4)

	dict := f.groundTruthDict()
	dict.SetObjectForKey(att, foundation.NSStringFromID(objc.String(name)))
	f.df.SetGroundTruthAttachments(dict)
	return nil
}

// SetOutputIOSurface attaches a named output backed by an IOSurface.
func (f *Frame) SetOutputIOSurface(name string, ref coregraphics.IOSurfaceRef) error {
	surfRef := iosurface.IOSurfaceRef(ref)
	base := iosurface.IOSurfaceGetBaseAddress(surfRef)
	if base == nil {
		return fmt.Errorf("espresso set output %q: IOSurface base address is nil", name)
	}
	size := iosurface.IOSurfaceGetAllocSize(surfRef)

	att := espresso.NewEspressoDataFrameTensorAttachment()
	att.SetRawPointer(base)
	att.SetSize(uint64(size))

	dict := f.outputDict()
	dict.SetObjectForKey(att, foundation.NSStringFromID(objc.String(name)))
	f.df.SetOutputAttachments(dict)
	return nil
}

// SetOutputANESurface attaches a named output from an Espresso ANE surface frame.
func (f *Frame) SetOutputANESurface(name string, surface *ANESurface, frame uint64) error {
	if surface == nil {
		return fmt.Errorf("espresso set output %q: ANESurface is nil", name)
	}
	ref, err := surface.IOSurfaceForFrame(frame)
	if err != nil {
		return err
	}
	return f.SetOutputIOSurface(name, ref)
}

// Output returns the named output tensor data after execution.
func (f *Frame) Output(name string) ([]float32, error) {
	key := objectivec.Object{ID: objc.String(name)}
	att := f.df.GetOutputAttachment(key)
	if att.GetID() == 0 {
		return nil, fmt.Errorf("%w: %q", ErrNoSuchOutput, name)
	}

	// Cast to EspressoDataFrameAttachment to access RawPointer/Size.
	dfa := espresso.EspressoDataFrameAttachmentFromID(att.GetID())
	ptr := dfa.RawPointer()
	size := dfa.Size()
	if ptr == nil || size == 0 {
		return nil, fmt.Errorf("%w: %q has no data", ErrNoSuchOutput, name)
	}

	n := int(size) / 4
	raw := unsafe.Slice((*float32)(ptr), n)
	out := make([]float32, n)
	copy(out, raw)
	return out, nil
}

// OutputInto copies the named output tensor data into the caller's buffer,
// avoiding allocation. Returns an error if dst is too small.
func (f *Frame) OutputInto(name string, dst []float32) error {
	key := objectivec.Object{ID: objc.String(name)}
	att := f.df.GetOutputAttachment(key)
	if att.GetID() == 0 {
		return fmt.Errorf("%w: %q", ErrNoSuchOutput, name)
	}

	dfa := espresso.EspressoDataFrameAttachmentFromID(att.GetID())
	ptr := dfa.RawPointer()
	size := dfa.Size()
	if ptr == nil || size == 0 {
		return fmt.Errorf("%w: %q has no data", ErrNoSuchOutput, name)
	}

	n := int(size) / 4
	if len(dst) < n {
		return fmt.Errorf("espresso output into %q: dst length %d < output elements %d", name, len(dst), n)
	}

	raw := unsafe.Slice((*float32)(ptr), n)
	copy(dst, raw)
	return nil
}

// OutputFloat64 returns the named output as float64 values.
func (f *Frame) OutputFloat64(name string) ([]float64, error) {
	f32, err := f.Output(name)
	if err != nil {
		return nil, err
	}
	out := make([]float64, len(f32))
	for i, v := range f32 {
		out[i] = float64(v)
	}
	return out, nil
}

// OutputRaw returns the raw bytes of a named output.
func (f *Frame) OutputRaw(name string) ([]byte, error) {
	key := objectivec.Object{ID: objc.String(name)}
	att := f.df.GetOutputAttachment(key)
	if att.GetID() == 0 {
		return nil, fmt.Errorf("%w: %q", ErrNoSuchOutput, name)
	}

	dfa := espresso.EspressoDataFrameAttachmentFromID(att.GetID())
	ptr := dfa.RawPointer()
	size := dfa.Size()
	if ptr == nil || size == 0 {
		return nil, fmt.Errorf("%w: %q has no data", ErrNoSuchOutput, name)
	}

	raw := unsafe.Slice((*byte)(ptr), int(size))
	out := make([]byte, int(size))
	copy(out, raw)
	return out, nil
}

// InputNames returns the names of all input attachments.
func (f *Frame) InputNames() []string {
	return nsArrayToStrings(f.df.InputAttachmentNames())
}

// OutputNames returns the names of all output attachments.
func (f *Frame) OutputNames() []string {
	return nsArrayToStrings(f.df.OutputAttachmentNames())
}

// FunctionName returns the function name set on this frame.
func (f *Frame) FunctionName() string {
	return f.df.Function_name()
}

// SetFunctionName sets the function name for multi-function models.
func (f *Frame) SetFunctionName(name string) {
	f.df.SetFunction_name(name)
}

func nsArrayToStrings(arr foundation.INSArray) []string {
	if arr == nil || arr.GetID() == 0 {
		return nil
	}
	a := foundation.NSArrayFromID(arr.GetID())
	n := a.Count()
	out := make([]string, n)
	for i := range n {
		obj := a.ObjectAtIndex(i)
		out[i] = foundation.NSStringFromID(obj.GetID()).String()
	}
	return out
}

func (f *Frame) inputDict() foundation.NSMutableDictionary {
	existing := f.df.InputAttachments()
	if existing != nil && existing.GetID() != 0 {
		return foundation.NewMutableDictionaryWithDictionary(existing)
	}
	return foundation.NewMutableDictionaryWithCapacity(4)
}

func (f *Frame) groundTruthDict() foundation.NSMutableDictionary {
	existing := f.df.GroundTruthAttachments()
	if existing != nil && existing.GetID() != 0 {
		return foundation.NewMutableDictionaryWithDictionary(existing)
	}
	return foundation.NewMutableDictionaryWithCapacity(4)
}

func (f *Frame) outputDict() foundation.NSMutableDictionary {
	existing := f.df.OutputAttachments()
	if existing != nil && existing.GetID() != 0 {
		return foundation.NewMutableDictionaryWithDictionary(existing)
	}
	return foundation.NewMutableDictionaryWithCapacity(4)
}
