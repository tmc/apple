// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVVCAudioBuffer] class.
var (
	_AVVCAudioBufferClass     AVVCAudioBufferClass
	_AVVCAudioBufferClassOnce sync.Once
)

func getAVVCAudioBufferClass() AVVCAudioBufferClass {
	_AVVCAudioBufferClassOnce.Do(func() {
		_AVVCAudioBufferClass = AVVCAudioBufferClass{class: objc.GetClass("AVVCAudioBuffer")}
	})
	return _AVVCAudioBufferClass
}

// GetAVVCAudioBufferClass returns the class object for AVVCAudioBuffer.
func GetAVVCAudioBufferClass() AVVCAudioBufferClass {
	return getAVVCAudioBufferClass()
}

type AVVCAudioBufferClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVVCAudioBufferClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVVCAudioBufferClass) Alloc() AVVCAudioBuffer {
	rv := objc.Send[AVVCAudioBuffer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVVCAudioBuffer.BytesCapacity]
//   - [AVVCAudioBuffer.BytesDataSize]
//   - [AVVCAudioBuffer.SetBytesDataSize]
//   - [AVVCAudioBuffer.Channels]
//   - [AVVCAudioBuffer.Data]
//   - [AVVCAudioBuffer.PacketDescriptionCapacity]
//   - [AVVCAudioBuffer.PacketDescriptionCount]
//   - [AVVCAudioBuffer.PacketDescriptions]
//   - [AVVCAudioBuffer.RemoteVoiceActivityAvailable]
//   - [AVVCAudioBuffer.RemoteVoiceActivityRMS]
//   - [AVVCAudioBuffer.RemoteVoiceActivityVAD]
//   - [AVVCAudioBuffer.SetPacketDescriptionsCount]
//   - [AVVCAudioBuffer.StreamDescription]
//   - [AVVCAudioBuffer.TimeStamp]
//   - [AVVCAudioBuffer.SetTimeStamp]
//   - [AVVCAudioBuffer.UpsamplingSourceAudio]
//   - [AVVCAudioBuffer.InitWithAudioQueueBufferChannelsTimeStamp]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCAudioBuffer
type AVVCAudioBuffer struct {
	objectivec.Object
}

// AVVCAudioBufferFromID constructs a [AVVCAudioBuffer] from an objc.ID.
func AVVCAudioBufferFromID(id objc.ID) AVVCAudioBuffer {
	return AVVCAudioBuffer{objectivec.Object{ID: id}}
}

// Ensure AVVCAudioBuffer implements IAVVCAudioBuffer.
var _ IAVVCAudioBuffer = AVVCAudioBuffer{}

// An interface definition for the [AVVCAudioBuffer] class.
//
// # Methods
//
//   - [IAVVCAudioBuffer.BytesCapacity]
//   - [IAVVCAudioBuffer.BytesDataSize]
//   - [IAVVCAudioBuffer.SetBytesDataSize]
//   - [IAVVCAudioBuffer.Channels]
//   - [IAVVCAudioBuffer.Data]
//   - [IAVVCAudioBuffer.PacketDescriptionCapacity]
//   - [IAVVCAudioBuffer.PacketDescriptionCount]
//   - [IAVVCAudioBuffer.PacketDescriptions]
//   - [IAVVCAudioBuffer.RemoteVoiceActivityAvailable]
//   - [IAVVCAudioBuffer.RemoteVoiceActivityRMS]
//   - [IAVVCAudioBuffer.RemoteVoiceActivityVAD]
//   - [IAVVCAudioBuffer.SetPacketDescriptionsCount]
//   - [IAVVCAudioBuffer.StreamDescription]
//   - [IAVVCAudioBuffer.TimeStamp]
//   - [IAVVCAudioBuffer.SetTimeStamp]
//   - [IAVVCAudioBuffer.UpsamplingSourceAudio]
//   - [IAVVCAudioBuffer.InitWithAudioQueueBufferChannelsTimeStamp]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCAudioBuffer
type IAVVCAudioBuffer interface {
	objectivec.IObject

	// Topic: Methods

	BytesCapacity() int
	BytesDataSize() int
	SetBytesDataSize(value int)
	Channels() int
	Data() unsafe.Pointer
	PacketDescriptionCapacity() int
	PacketDescriptionCount() int
	PacketDescriptions() *AudioStreamPacketDescriptionRef
	RemoteVoiceActivityAvailable() bool
	RemoteVoiceActivityRMS() byte
	RemoteVoiceActivityVAD() byte
	SetPacketDescriptionsCount(descriptions []AudioStreamPacketDescriptionRef, count int)
	StreamDescription() unsafe.Pointer
	TimeStamp() uint64
	SetTimeStamp(value uint64)
	UpsamplingSourceAudio() bool
	InitWithAudioQueueBufferChannelsTimeStamp(buffer unsafe.Pointer, channels int, stamp uint64) AVVCAudioBuffer
}

// Init initializes the instance.
func (v AVVCAudioBuffer) Init() AVVCAudioBuffer {
	rv := objc.Send[AVVCAudioBuffer](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v AVVCAudioBuffer) Autorelease() AVVCAudioBuffer {
	rv := objc.Send[AVVCAudioBuffer](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVCAudioBuffer creates a new AVVCAudioBuffer instance.
func NewAVVCAudioBuffer() AVVCAudioBuffer {
	class := getAVVCAudioBufferClass()
	rv := objc.Send[AVVCAudioBuffer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCAudioBuffer/initWithAudioQueueBuffer:channels:timeStamp:
func NewVCAudioBufferWithAudioQueueBufferChannelsTimeStamp(buffer unsafe.Pointer, channels int, stamp uint64) AVVCAudioBuffer {
	instance := getAVVCAudioBufferClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAudioQueueBuffer:channels:timeStamp:"), buffer, channels, stamp)
	return AVVCAudioBufferFromID(rv)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCAudioBuffer/setPacketDescriptions:count:
func (v AVVCAudioBuffer) SetPacketDescriptionsCount(descriptions []AudioStreamPacketDescriptionRef, count int) {
	objc.Send[objc.ID](v.ID, objc.Sel("setPacketDescriptions:count:"), objc.CArray(descriptions), count)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCAudioBuffer/initWithAudioQueueBuffer:channels:timeStamp:
func (v AVVCAudioBuffer) InitWithAudioQueueBufferChannelsTimeStamp(buffer unsafe.Pointer, channels int, stamp uint64) AVVCAudioBuffer {
	rv := objc.Send[AVVCAudioBuffer](v.ID, objc.Sel("initWithAudioQueueBuffer:channels:timeStamp:"), buffer, channels, stamp)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCAudioBuffer/bytesCapacity
func (v AVVCAudioBuffer) BytesCapacity() int {
	rv := objc.Send[int](v.ID, objc.Sel("bytesCapacity"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCAudioBuffer/bytesDataSize
func (v AVVCAudioBuffer) BytesDataSize() int {
	rv := objc.Send[int](v.ID, objc.Sel("bytesDataSize"))
	return rv
}
func (v AVVCAudioBuffer) SetBytesDataSize(value int) {
	objc.Send[struct{}](v.ID, objc.Sel("setBytesDataSize:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCAudioBuffer/channels
func (v AVVCAudioBuffer) Channels() int {
	rv := objc.Send[int](v.ID, objc.Sel("channels"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCAudioBuffer/data
func (v AVVCAudioBuffer) Data() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v.ID, objc.Sel("data"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCAudioBuffer/packetDescriptionCapacity
func (v AVVCAudioBuffer) PacketDescriptionCapacity() int {
	rv := objc.Send[int](v.ID, objc.Sel("packetDescriptionCapacity"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCAudioBuffer/packetDescriptionCount
func (v AVVCAudioBuffer) PacketDescriptionCount() int {
	rv := objc.Send[int](v.ID, objc.Sel("packetDescriptionCount"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCAudioBuffer/packetDescriptions
func (v AVVCAudioBuffer) PacketDescriptions() *AudioStreamPacketDescriptionRef {
	rv := objc.Send[unsafe.Pointer](v.ID, objc.Sel("packetDescriptions"))
	return (*AudioStreamPacketDescriptionRef)(rv)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCAudioBuffer/remoteVoiceActivityAvailable
func (v AVVCAudioBuffer) RemoteVoiceActivityAvailable() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("remoteVoiceActivityAvailable"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCAudioBuffer/remoteVoiceActivityRMS
func (v AVVCAudioBuffer) RemoteVoiceActivityRMS() byte {
	rv := objc.Send[byte](v.ID, objc.Sel("remoteVoiceActivityRMS"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCAudioBuffer/remoteVoiceActivityVAD
func (v AVVCAudioBuffer) RemoteVoiceActivityVAD() byte {
	rv := objc.Send[byte](v.ID, objc.Sel("remoteVoiceActivityVAD"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCAudioBuffer/streamDescription
func (v AVVCAudioBuffer) StreamDescription() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v.ID, objc.Sel("streamDescription"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCAudioBuffer/timeStamp
func (v AVVCAudioBuffer) TimeStamp() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("timeStamp"))
	return rv
}
func (v AVVCAudioBuffer) SetTimeStamp(value uint64) {
	objc.Send[struct{}](v.ID, objc.Sel("setTimeStamp:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCAudioBuffer/upsamplingSourceAudio
func (v AVVCAudioBuffer) UpsamplingSourceAudio() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("upsamplingSourceAudio"))
	return rv
}
