// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/coreaudiotypes"
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
	rv := objc.SendIfResponds[AVVCAudioBuffer](objc.ID(ac.class), objc.Sel("alloc"))
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
	PacketDescriptions() coreaudiotypes.AudioStreamPacketDescription
	RemoteVoiceActivityAvailable() bool
	RemoteVoiceActivityRMS() byte
	RemoteVoiceActivityVAD() byte
	SetPacketDescriptionsCount(descriptions []coreaudiotypes.AudioStreamPacketDescription, count int)
	StreamDescription() coreaudiotypes.AudioStreamBasicDescription
	TimeStamp() uint64
	SetTimeStamp(value uint64)
	UpsamplingSourceAudio() bool
	InitWithAudioQueueBufferChannelsTimeStamp(buffer *MyAudioQueueBuffer, channels int, stamp uint64) AVVCAudioBuffer
}

// Init initializes the instance.
func (a AVVCAudioBuffer) Init() AVVCAudioBuffer {
	rv := objc.SendIfResponds[AVVCAudioBuffer](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVVCAudioBuffer) Autorelease() AVVCAudioBuffer {
	rv := objc.SendIfResponds[AVVCAudioBuffer](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVCAudioBuffer creates a new AVVCAudioBuffer instance.
func NewAVVCAudioBuffer() AVVCAudioBuffer {
	class := getAVVCAudioBufferClass()
	rv := objc.SendIfResponds[AVVCAudioBuffer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewVCAudioBufferWithAudioQueueBufferChannelsTimeStamp(buffer *MyAudioQueueBuffer, channels int, stamp uint64) AVVCAudioBuffer {
	instance := getAVVCAudioBufferClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithAudioQueueBuffer:channels:timeStamp:"), unsafe.Pointer(buffer), channels, stamp)
	return AVVCAudioBufferFromID(rv)
}

func (a AVVCAudioBuffer) SetPacketDescriptionsCount(descriptions []coreaudiotypes.AudioStreamPacketDescription, count int) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("setPacketDescriptions:count:"), objc.CArray(descriptions), count)
}
func (a AVVCAudioBuffer) InitWithAudioQueueBufferChannelsTimeStamp(buffer *MyAudioQueueBuffer, channels int, stamp uint64) AVVCAudioBuffer {
	rv := objc.SendIfResponds[AVVCAudioBuffer](a.ID, objc.Sel("initWithAudioQueueBuffer:channels:timeStamp:"), unsafe.Pointer(buffer), channels, stamp)
	return rv
}

func (a AVVCAudioBuffer) BytesCapacity() int {
	rv := objc.SendIfResponds[int](a.ID, objc.Sel("bytesCapacity"))
	return rv
}
func (a AVVCAudioBuffer) BytesDataSize() int {
	rv := objc.SendIfResponds[int](a.ID, objc.Sel("bytesDataSize"))
	return rv
}
func (a AVVCAudioBuffer) SetBytesDataSize(value int) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setBytesDataSize:"), value)
}
func (a AVVCAudioBuffer) Channels() int {
	rv := objc.SendIfResponds[int](a.ID, objc.Sel("channels"))
	return rv
}
func (a AVVCAudioBuffer) Data() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](a.ID, objc.Sel("data"))
	return rv
}
func (a AVVCAudioBuffer) PacketDescriptionCapacity() int {
	rv := objc.SendIfResponds[int](a.ID, objc.Sel("packetDescriptionCapacity"))
	return rv
}
func (a AVVCAudioBuffer) PacketDescriptionCount() int {
	rv := objc.SendIfResponds[int](a.ID, objc.Sel("packetDescriptionCount"))
	return rv
}
func (a AVVCAudioBuffer) PacketDescriptions() coreaudiotypes.AudioStreamPacketDescription {
	rv := objc.SendIfResponds[coreaudiotypes.AudioStreamPacketDescription](a.ID, objc.Sel("packetDescriptions"))
	return coreaudiotypes.AudioStreamPacketDescription(rv)
}
func (a AVVCAudioBuffer) RemoteVoiceActivityAvailable() bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("remoteVoiceActivityAvailable"))
	return rv
}
func (a AVVCAudioBuffer) RemoteVoiceActivityRMS() byte {
	rv := objc.SendIfResponds[byte](a.ID, objc.Sel("remoteVoiceActivityRMS"))
	return rv
}
func (a AVVCAudioBuffer) RemoteVoiceActivityVAD() byte {
	rv := objc.SendIfResponds[byte](a.ID, objc.Sel("remoteVoiceActivityVAD"))
	return rv
}
func (a AVVCAudioBuffer) StreamDescription() coreaudiotypes.AudioStreamBasicDescription {
	rv := objc.SendIfResponds[coreaudiotypes.AudioStreamBasicDescription](a.ID, objc.Sel("streamDescription"))
	return coreaudiotypes.AudioStreamBasicDescription(rv)
}
func (a AVVCAudioBuffer) TimeStamp() uint64 {
	rv := objc.SendIfResponds[uint64](a.ID, objc.Sel("timeStamp"))
	return rv
}
func (a AVVCAudioBuffer) SetTimeStamp(value uint64) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setTimeStamp:"), value)
}
func (a AVVCAudioBuffer) UpsamplingSourceAudio() bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("upsamplingSourceAudio"))
	return rv
}
