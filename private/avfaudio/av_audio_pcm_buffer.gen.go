// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioPCMBuffer] class.
var (
	_AVAudioPCMBufferClass     AVAudioPCMBufferClass
	_AVAudioPCMBufferClassOnce sync.Once
)

func getAVAudioPCMBufferClass() AVAudioPCMBufferClass {
	_AVAudioPCMBufferClassOnce.Do(func() {
		_AVAudioPCMBufferClass = AVAudioPCMBufferClass{class: objc.GetClass("AVAudioPCMBuffer")}
	})
	return _AVAudioPCMBufferClass
}

// GetAVAudioPCMBufferClass returns the class object for AVAudioPCMBuffer.
func GetAVAudioPCMBufferClass() AVAudioPCMBufferClass {
	return getAVAudioPCMBufferClass()
}

type AVAudioPCMBufferClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioPCMBufferClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioPCMBufferClass) Alloc() AVAudioPCMBuffer {
	rv := objc.SendIfResponds[AVAudioPCMBuffer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVAudioPCMBuffer._initChannelPtrs]
//   - [AVAudioPCMBuffer.AppendDataFromBuffer]
//   - [AVAudioPCMBuffer.AppendDataFromBufferChannel]
//   - [AVAudioPCMBuffer.AveragePowerPerChannel]
//   - [AVAudioPCMBuffer.CalculatePower]
//   - [AVAudioPCMBuffer.CalculatePowerForFloatDataStrideFrameLength]
//   - [AVAudioPCMBuffer.PeakPowerPerChannel]
//   - [AVAudioPCMBuffer.SplitIntoSingleChannelBuffers]
type AVAudioPCMBuffer struct {
	AVAudioBuffer
}

// AVAudioPCMBufferFromID constructs a [AVAudioPCMBuffer] from an objc.ID.
func AVAudioPCMBufferFromID(id objc.ID) AVAudioPCMBuffer {
	return AVAudioPCMBuffer{AVAudioBuffer: AVAudioBufferFromID(id)}
}

// Ensure AVAudioPCMBuffer implements IAVAudioPCMBuffer.
var _ IAVAudioPCMBuffer = AVAudioPCMBuffer{}

// An interface definition for the [AVAudioPCMBuffer] class.
//
// # Methods
//
//   - [IAVAudioPCMBuffer._initChannelPtrs]
//   - [IAVAudioPCMBuffer.AppendDataFromBuffer]
//   - [IAVAudioPCMBuffer.AppendDataFromBufferChannel]
//   - [IAVAudioPCMBuffer.AveragePowerPerChannel]
//   - [IAVAudioPCMBuffer.CalculatePower]
//   - [IAVAudioPCMBuffer.CalculatePowerForFloatDataStrideFrameLength]
//   - [IAVAudioPCMBuffer.PeakPowerPerChannel]
//   - [IAVAudioPCMBuffer.SplitIntoSingleChannelBuffers]
type IAVAudioPCMBuffer interface {
	IAVAudioBuffer

	// Topic: Methods

	_initChannelPtrs()
	AppendDataFromBuffer(buffer objectivec.IObject) bool
	AppendDataFromBufferChannel(buffer objectivec.IObject, channel int64) bool
	AveragePowerPerChannel() foundation.INSArray
	CalculatePower(power uint64) objectivec.IObject
	CalculatePowerForFloatDataStrideFrameLength(power uint64, data *float32, stride int64, length uint32) float32
	PeakPowerPerChannel() foundation.INSArray
	SplitIntoSingleChannelBuffers() objectivec.IObject
}

// Init initializes the instance.
func (a AVAudioPCMBuffer) Init() AVAudioPCMBuffer {
	rv := objc.SendIfResponds[AVAudioPCMBuffer](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioPCMBuffer) Autorelease() AVAudioPCMBuffer {
	rv := objc.SendIfResponds[AVAudioPCMBuffer](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioPCMBuffer creates a new AVAudioPCMBuffer instance.
func NewAVAudioPCMBuffer() AVAudioPCMBuffer {
	class := getAVAudioPCMBufferClass()
	rv := objc.SendIfResponds[AVAudioPCMBuffer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewAudioPCMBufferWithFormatByteCapacity(format objectivec.IObject, capacity uint32) AVAudioPCMBuffer {
	instance := getAVAudioPCMBufferClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithFormat:byteCapacity:"), format, capacity)
	return AVAudioPCMBufferFromID(rv)
}

func (a AVAudioPCMBuffer) _initChannelPtrs() {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("_initChannelPtrs"))
}

// InitChannelPtrs is an exported wrapper for the private method _initChannelPtrs.
func (a AVAudioPCMBuffer) InitChannelPtrs() error {
	if !objc.RespondsToSelector(a.ID, objc.Sel("_initChannelPtrs")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_initChannelPtrs"}
		return err
	}
	a._initChannelPtrs()
	return nil
}

// CanInitChannelPtrs reports whether the receiver responds to the private selector _initChannelPtrs.
func (a AVAudioPCMBuffer) CanInitChannelPtrs() bool {
	return objc.RespondsToSelector(a.ID, objc.Sel("_initChannelPtrs"))
}
func (a AVAudioPCMBuffer) AppendDataFromBuffer(buffer objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("appendDataFromBuffer:"), buffer)
	return rv
}
func (a AVAudioPCMBuffer) AppendDataFromBufferChannel(buffer objectivec.IObject, channel int64) bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("appendDataFromBuffer:channel:"), buffer, channel)
	return rv
}
func (a AVAudioPCMBuffer) CalculatePower(power uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("calculatePower:"), power)
	return objectivec.Object{ID: rv}
}
func (a AVAudioPCMBuffer) CalculatePowerForFloatDataStrideFrameLength(power uint64, data *float32, stride int64, length uint32) float32 {
	rv := objc.SendIfResponds[float32](a.ID, objc.Sel("calculatePower:forFloatData:stride:frameLength:"), power, data, stride, length)
	return rv
}
func (a AVAudioPCMBuffer) SplitIntoSingleChannelBuffers() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("splitIntoSingleChannelBuffers"))
	return objectivec.Object{ID: rv}
}

func (a AVAudioPCMBuffer) AveragePowerPerChannel() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("averagePowerPerChannel"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (a AVAudioPCMBuffer) PeakPowerPerChannel() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("peakPowerPerChannel"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
