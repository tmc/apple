// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
)

// The class instance for the [AVAudioCompressedBuffer] class.
var (
	_AVAudioCompressedBufferClass     AVAudioCompressedBufferClass
	_AVAudioCompressedBufferClassOnce sync.Once
)

func getAVAudioCompressedBufferClass() AVAudioCompressedBufferClass {
	_AVAudioCompressedBufferClassOnce.Do(func() {
		_AVAudioCompressedBufferClass = AVAudioCompressedBufferClass{class: objc.GetClass("AVAudioCompressedBuffer")}
	})
	return _AVAudioCompressedBufferClass
}

// GetAVAudioCompressedBufferClass returns the class object for AVAudioCompressedBuffer.
func GetAVAudioCompressedBufferClass() AVAudioCompressedBufferClass {
	return getAVAudioCompressedBufferClass()
}

type AVAudioCompressedBufferClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioCompressedBufferClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioCompressedBufferClass) Alloc() AVAudioCompressedBuffer {
	rv := objc.Send[AVAudioCompressedBuffer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents an audio buffer that you use for compressed audio
// formats.
//
// # Creating an Audio Buffer
//
//   - [AVAudioCompressedBuffer.InitWithFormatPacketCapacity]: Creates a buffer that contains constant bytes per packet of audio data in a compressed state.
//   - [AVAudioCompressedBuffer.InitWithFormatPacketCapacityMaximumPacketSize]: Creates a buffer that contains audio data in a compressed state.
//
// # Getting Audio Buffer Properties
//
//   - [AVAudioCompressedBuffer.ByteCapacity]: The number of packets the buffer contains.
//   - [AVAudioCompressedBuffer.ByteLength]: The number of valid bytes in the buffer.
//   - [AVAudioCompressedBuffer.SetByteLength]
//   - [AVAudioCompressedBuffer.Data]: The audio buffer’s data bytes.
//   - [AVAudioCompressedBuffer.MaximumPacketSize]: The maximum size of a packet, in bytes.
//   - [AVAudioCompressedBuffer.PacketCapacity]: The total number of packets that the buffer can contain.
//   - [AVAudioCompressedBuffer.PacketCount]: The number of packets currently in the buffer.
//   - [AVAudioCompressedBuffer.SetPacketCount]
//   - [AVAudioCompressedBuffer.PacketDescriptions]: The buffer’s array of packet descriptions.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioCompressedBuffer
type AVAudioCompressedBuffer struct {
	AVAudioBuffer
}

// AVAudioCompressedBufferFromID constructs a [AVAudioCompressedBuffer] from an objc.ID.
//
// An object that represents an audio buffer that you use for compressed audio
// formats.
func AVAudioCompressedBufferFromID(id objc.ID) AVAudioCompressedBuffer {
	return AVAudioCompressedBuffer{AVAudioBuffer: AVAudioBufferFromID(id)}
}

// NOTE: AVAudioCompressedBuffer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioCompressedBuffer] class.
//
// # Creating an Audio Buffer
//
//   - [IAVAudioCompressedBuffer.InitWithFormatPacketCapacity]: Creates a buffer that contains constant bytes per packet of audio data in a compressed state.
//   - [IAVAudioCompressedBuffer.InitWithFormatPacketCapacityMaximumPacketSize]: Creates a buffer that contains audio data in a compressed state.
//
// # Getting Audio Buffer Properties
//
//   - [IAVAudioCompressedBuffer.ByteCapacity]: The number of packets the buffer contains.
//   - [IAVAudioCompressedBuffer.ByteLength]: The number of valid bytes in the buffer.
//   - [IAVAudioCompressedBuffer.SetByteLength]
//   - [IAVAudioCompressedBuffer.Data]: The audio buffer’s data bytes.
//   - [IAVAudioCompressedBuffer.MaximumPacketSize]: The maximum size of a packet, in bytes.
//   - [IAVAudioCompressedBuffer.PacketCapacity]: The total number of packets that the buffer can contain.
//   - [IAVAudioCompressedBuffer.PacketCount]: The number of packets currently in the buffer.
//   - [IAVAudioCompressedBuffer.SetPacketCount]
//   - [IAVAudioCompressedBuffer.PacketDescriptions]: The buffer’s array of packet descriptions.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioCompressedBuffer
type IAVAudioCompressedBuffer interface {
	IAVAudioBuffer

	// Topic: Creating an Audio Buffer

	// Creates a buffer that contains constant bytes per packet of audio data in a compressed state.
	InitWithFormatPacketCapacity(format IAVAudioFormat, packetCapacity AVAudioPacketCount) AVAudioCompressedBuffer
	// Creates a buffer that contains audio data in a compressed state.
	InitWithFormatPacketCapacityMaximumPacketSize(format IAVAudioFormat, packetCapacity AVAudioPacketCount, maximumPacketSize int) AVAudioCompressedBuffer

	// Topic: Getting Audio Buffer Properties

	// The number of packets the buffer contains.
	ByteCapacity() uint32
	// The number of valid bytes in the buffer.
	ByteLength() uint32
	SetByteLength(value uint32)
	// The audio buffer’s data bytes.
	Data() unsafe.Pointer
	// The maximum size of a packet, in bytes.
	MaximumPacketSize() int
	// The total number of packets that the buffer can contain.
	PacketCapacity() AVAudioPacketCount
	// The number of packets currently in the buffer.
	PacketCount() AVAudioPacketCount
	SetPacketCount(value AVAudioPacketCount)
	// The buffer’s array of packet descriptions.
	PacketDescriptions() unsafe.Pointer

	// The buffer’s array of packet dependencies.
	PacketDependencies() unsafe.Pointer
}

// Init initializes the instance.
func (a AVAudioCompressedBuffer) Init() AVAudioCompressedBuffer {
	rv := objc.Send[AVAudioCompressedBuffer](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioCompressedBuffer) Autorelease() AVAudioCompressedBuffer {
	rv := objc.Send[AVAudioCompressedBuffer](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioCompressedBuffer creates a new AVAudioCompressedBuffer instance.
func NewAVAudioCompressedBuffer() AVAudioCompressedBuffer {
	class := getAVAudioCompressedBufferClass()
	rv := objc.Send[AVAudioCompressedBuffer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a buffer that contains constant bytes per packet of audio data in a
// compressed state.
//
// format: The format of the audio the buffer contains.
//
// packetCapacity: The capacity of the buffer, in packets.
//
// # Return Value
//
// A new [AVAudioCompressedBuffer] instance.
//
// # Discussion
//
// This fails if the format is PCM or if the format has variable bytes per
// packet (for example, `format.StreamDescription()->mBytesPerPacket == 0`).
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioCompressedBuffer/init(format:packetCapacity:)
func NewAudioCompressedBufferWithFormatPacketCapacity(format IAVAudioFormat, packetCapacity AVAudioPacketCount) AVAudioCompressedBuffer {
	instance := getAVAudioCompressedBufferClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFormat:packetCapacity:"), format, packetCapacity)
	return AVAudioCompressedBufferFromID(rv)
}

// Creates a buffer that contains audio data in a compressed state.
//
// format: The format of the audio the buffer contains.
//
// packetCapacity: The capacity of the buffer, in packets.
//
// maximumPacketSize: The maximum size in bytes of a packet in a compressed state.
//
// # Return Value
//
// A new [AVAudioCompressedBuffer] instance.
//
// # Discussion
//
// You can obtain the maximum packet size from the [MaximumOutputPacketSize]
// property of an [AVAudioConverter] you configure for encoding this format.
//
// The method raises an exception if the format is PCM.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioCompressedBuffer/init(format:packetCapacity:maximumPacketSize:)
func NewAudioCompressedBufferWithFormatPacketCapacityMaximumPacketSize(format IAVAudioFormat, packetCapacity AVAudioPacketCount, maximumPacketSize int) AVAudioCompressedBuffer {
	instance := getAVAudioCompressedBufferClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFormat:packetCapacity:maximumPacketSize:"), format, packetCapacity, maximumPacketSize)
	return AVAudioCompressedBufferFromID(rv)
}

// Creates a buffer that contains constant bytes per packet of audio data in a
// compressed state.
//
// format: The format of the audio the buffer contains.
//
// packetCapacity: The capacity of the buffer, in packets.
//
// # Return Value
//
// A new [AVAudioCompressedBuffer] instance.
//
// # Discussion
//
// This fails if the format is PCM or if the format has variable bytes per
// packet (for example, `format.StreamDescription()->mBytesPerPacket == 0`).
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioCompressedBuffer/init(format:packetCapacity:)
func (a AVAudioCompressedBuffer) InitWithFormatPacketCapacity(format IAVAudioFormat, packetCapacity AVAudioPacketCount) AVAudioCompressedBuffer {
	rv := objc.Send[AVAudioCompressedBuffer](a.ID, objc.Sel("initWithFormat:packetCapacity:"), format, packetCapacity)
	return rv
}

// Creates a buffer that contains audio data in a compressed state.
//
// format: The format of the audio the buffer contains.
//
// packetCapacity: The capacity of the buffer, in packets.
//
// maximumPacketSize: The maximum size in bytes of a packet in a compressed state.
//
// # Return Value
//
// A new [AVAudioCompressedBuffer] instance.
//
// # Discussion
//
// You can obtain the maximum packet size from the [MaximumOutputPacketSize]
// property of an [AVAudioConverter] you configure for encoding this format.
//
// The method raises an exception if the format is PCM.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioCompressedBuffer/init(format:packetCapacity:maximumPacketSize:)
func (a AVAudioCompressedBuffer) InitWithFormatPacketCapacityMaximumPacketSize(format IAVAudioFormat, packetCapacity AVAudioPacketCount, maximumPacketSize int) AVAudioCompressedBuffer {
	rv := objc.Send[AVAudioCompressedBuffer](a.ID, objc.Sel("initWithFormat:packetCapacity:maximumPacketSize:"), format, packetCapacity, maximumPacketSize)
	return rv
}

// The number of packets the buffer contains.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioCompressedBuffer/byteCapacity
func (a AVAudioCompressedBuffer) ByteCapacity() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("byteCapacity"))
	return rv
}

// The number of valid bytes in the buffer.
//
// # Discussion
//
// You can change this value as part of an operation that modifies the
// contents.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioCompressedBuffer/byteLength
func (a AVAudioCompressedBuffer) ByteLength() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("byteLength"))
	return rv
}
func (a AVAudioCompressedBuffer) SetByteLength(value uint32) {
	objc.Send[struct{}](a.ID, objc.Sel("setByteLength:"), value)
}

// The audio buffer’s data bytes.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioCompressedBuffer/data
func (a AVAudioCompressedBuffer) Data() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a.ID, objc.Sel("data"))
	return rv
}

// The maximum size of a packet, in bytes.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioCompressedBuffer/maximumPacketSize
func (a AVAudioCompressedBuffer) MaximumPacketSize() int {
	rv := objc.Send[int](a.ID, objc.Sel("maximumPacketSize"))
	return rv
}

// The total number of packets that the buffer can contain.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioCompressedBuffer/packetCapacity
func (a AVAudioCompressedBuffer) PacketCapacity() AVAudioPacketCount {
	rv := objc.Send[AVAudioPacketCount](a.ID, objc.Sel("packetCapacity"))
	return AVAudioPacketCount(rv)
}

// The number of packets currently in the buffer.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioCompressedBuffer/packetCount
func (a AVAudioCompressedBuffer) PacketCount() AVAudioPacketCount {
	rv := objc.Send[AVAudioPacketCount](a.ID, objc.Sel("packetCount"))
	return AVAudioPacketCount(rv)
}
func (a AVAudioCompressedBuffer) SetPacketCount(value AVAudioPacketCount) {
	objc.Send[struct{}](a.ID, objc.Sel("setPacketCount:"), value)
}

// The buffer’s array of packet descriptions.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioCompressedBuffer/packetDescriptions
func (a AVAudioCompressedBuffer) PacketDescriptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a.ID, objc.Sel("packetDescriptions"))
	return rv
}

// The buffer’s array of packet dependencies.
//
// # Discussion
//
// If the audio format doesn’t use packet dependencies, this value is `nil`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioCompressedBuffer/packetDependencies-5oae6
func (a AVAudioCompressedBuffer) PacketDependencies() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a.ID, objc.Sel("packetDependencies"))
	return rv
}
