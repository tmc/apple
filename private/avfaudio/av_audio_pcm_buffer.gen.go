// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
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
	rv := objc.Send[AVAudioPCMBuffer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
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
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPCMBuffer
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
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPCMBuffer
type IAVAudioPCMBuffer interface {
	IAVAudioBuffer

	// Topic: Methods

	_initChannelPtrs()
	AppendDataFromBuffer(buffer objectivec.IObject) bool
	AppendDataFromBufferChannel(buffer objectivec.IObject, channel int64) bool
	AveragePowerPerChannel() foundation.INSArray
	CalculatePower(power uint64) objectivec.IObject
	CalculatePowerForFloatDataStrideFrameLength(power uint64, data unsafe.Pointer, stride int64, length uint32) float32
	PeakPowerPerChannel() foundation.INSArray
	SplitIntoSingleChannelBuffers() objectivec.IObject
}

// Init initializes the instance.
func (a AVAudioPCMBuffer) Init() AVAudioPCMBuffer {
	rv := objc.Send[AVAudioPCMBuffer](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioPCMBuffer) Autorelease() AVAudioPCMBuffer {
	rv := objc.Send[AVAudioPCMBuffer](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioPCMBuffer creates a new AVAudioPCMBuffer instance.
func NewAVAudioPCMBuffer() AVAudioPCMBuffer {
	class := getAVAudioPCMBufferClass()
	rv := objc.Send[AVAudioPCMBuffer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioBuffer/initWithFormat:byteCapacity:
func NewAudioPCMBufferWithFormatByteCapacity(format objectivec.IObject, capacity uint32) AVAudioPCMBuffer {
	instance := getAVAudioPCMBufferClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFormat:byteCapacity:"), format, capacity)
	return AVAudioPCMBufferFromID(rv)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPCMBuffer/_initChannelPtrs
func (a AVAudioPCMBuffer) _initChannelPtrs() {
	objc.Send[objc.ID](a.ID, objc.Sel("_initChannelPtrs"))
}

// InitChannelPtrs is an exported wrapper for the private method _initChannelPtrs.
func (a AVAudioPCMBuffer) InitChannelPtrs() {
	a._initChannelPtrs()
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPCMBuffer/appendDataFromBuffer:
func (a AVAudioPCMBuffer) AppendDataFromBuffer(buffer objectivec.IObject) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("appendDataFromBuffer:"), buffer)
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPCMBuffer/appendDataFromBuffer:channel:
func (a AVAudioPCMBuffer) AppendDataFromBufferChannel(buffer objectivec.IObject, channel int64) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("appendDataFromBuffer:channel:"), buffer, channel)
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPCMBuffer/calculatePower:
func (a AVAudioPCMBuffer) CalculatePower(power uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("calculatePower:"), power)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPCMBuffer/calculatePower:forFloatData:stride:frameLength:
func (a AVAudioPCMBuffer) CalculatePowerForFloatDataStrideFrameLength(power uint64, data unsafe.Pointer, stride int64, length uint32) float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("calculatePower:forFloatData:stride:frameLength:"), power, data, stride, length)
	return rv
}
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPCMBuffer/splitIntoSingleChannelBuffers
func (a AVAudioPCMBuffer) SplitIntoSingleChannelBuffers() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("splitIntoSingleChannelBuffers"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPCMBuffer/averagePowerPerChannel
func (a AVAudioPCMBuffer) AveragePowerPerChannel() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("averagePowerPerChannel"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPCMBuffer/peakPowerPerChannel
func (a AVAudioPCMBuffer) PeakPowerPerChannel() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("peakPowerPerChannel"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

