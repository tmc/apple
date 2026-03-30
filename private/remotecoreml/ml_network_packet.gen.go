// Code generated from Apple documentation for remotecoreml. DO NOT EDIT.

package remotecoreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLNetworkPacket] class.
var (
	_MLNetworkPacketClass     MLNetworkPacketClass
	_MLNetworkPacketClassOnce sync.Once
)

func getMLNetworkPacketClass() MLNetworkPacketClass {
	_MLNetworkPacketClassOnce.Do(func() {
		_MLNetworkPacketClass = MLNetworkPacketClass{class: objc.GetClass("_MLNetworkPacket")}
	})
	return _MLNetworkPacketClass
}

// GetMLNetworkPacketClass returns the class object for _MLNetworkPacket.
func GetMLNetworkPacketClass() MLNetworkPacketClass {
	return getMLNetworkPacketClass()
}

type MLNetworkPacketClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLNetworkPacketClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLNetworkPacketClass) Alloc() MLNetworkPacket {
	rv := objc.Send[MLNetworkPacket](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLNetworkPacket.Buffer]
//   - [MLNetworkPacket.SetBuffer]
//   - [MLNetworkPacket.CleanupDoubleBuffer]
//   - [MLNetworkPacket.Command]
//   - [MLNetworkPacket.SetCommand]
//   - [MLNetworkPacket.DoubleBuffer]
//   - [MLNetworkPacket.SetDoubleBuffer]
//   - [MLNetworkPacket.Reset]
//   - [MLNetworkPacket.ResetDoubleBuffer]
//   - [MLNetworkPacket.ResetMetadata]
//   - [MLNetworkPacket.SizeOfPacket]
//   - [MLNetworkPacket.SetSizeOfPacket]
//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkPacket
type MLNetworkPacket struct {
	objectivec.Object
}

// MLNetworkPacketFromID constructs a [MLNetworkPacket] from an objc.ID.
func MLNetworkPacketFromID(id objc.ID) MLNetworkPacket {
	return MLNetworkPacket{objectivec.Object{ID: id}}
}

// Ensure MLNetworkPacket implements IMLNetworkPacket.
var _ IMLNetworkPacket = MLNetworkPacket{}

// An interface definition for the [MLNetworkPacket] class.
//
// # Methods
//
//   - [IMLNetworkPacket.Buffer]
//   - [IMLNetworkPacket.SetBuffer]
//   - [IMLNetworkPacket.CleanupDoubleBuffer]
//   - [IMLNetworkPacket.Command]
//   - [IMLNetworkPacket.SetCommand]
//   - [IMLNetworkPacket.DoubleBuffer]
//   - [IMLNetworkPacket.SetDoubleBuffer]
//   - [IMLNetworkPacket.Reset]
//   - [IMLNetworkPacket.ResetDoubleBuffer]
//   - [IMLNetworkPacket.ResetMetadata]
//   - [IMLNetworkPacket.SizeOfPacket]
//   - [IMLNetworkPacket.SetSizeOfPacket]
//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkPacket
type IMLNetworkPacket interface {
	objectivec.IObject

	// Topic: Methods

	Buffer() foundation.NSMutableData
	SetBuffer(value foundation.NSMutableData)
	CleanupDoubleBuffer()
	Command() uint64
	SetCommand(value uint64)
	DoubleBuffer() foundation.NSMutableData
	SetDoubleBuffer(value foundation.NSMutableData)
	Reset()
	ResetDoubleBuffer()
	ResetMetadata()
	SizeOfPacket() uint64
	SetSizeOfPacket(value uint64)
}

// Init initializes the instance.
func (m MLNetworkPacket) Init() MLNetworkPacket {
	rv := objc.Send[MLNetworkPacket](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLNetworkPacket) Autorelease() MLNetworkPacket {
	rv := objc.Send[MLNetworkPacket](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLNetworkPacket creates a new MLNetworkPacket instance.
func NewMLNetworkPacket() MLNetworkPacket {
	class := getMLNetworkPacketClass()
	rv := objc.Send[MLNetworkPacket](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkPacket/cleanupDoubleBuffer
func (m MLNetworkPacket) CleanupDoubleBuffer() {
	objc.Send[objc.ID](m.ID, objc.Sel("cleanupDoubleBuffer"))
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkPacket/reset
func (m MLNetworkPacket) Reset() {
	objc.Send[objc.ID](m.ID, objc.Sel("reset"))
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkPacket/resetDoubleBuffer
func (m MLNetworkPacket) ResetDoubleBuffer() {
	objc.Send[objc.ID](m.ID, objc.Sel("resetDoubleBuffer"))
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkPacket/resetMetadata
func (m MLNetworkPacket) ResetMetadata() {
	objc.Send[objc.ID](m.ID, objc.Sel("resetMetadata"))
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkPacket/buffer
func (m MLNetworkPacket) Buffer() foundation.NSMutableData {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("buffer"))
	return foundation.NSMutableDataFromID(objc.ID(rv))
}
func (m MLNetworkPacket) SetBuffer(value foundation.NSMutableData) {
	objc.Send[struct{}](m.ID, objc.Sel("setBuffer:"), value)
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkPacket/command
func (m MLNetworkPacket) Command() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("command"))
	return rv
}
func (m MLNetworkPacket) SetCommand(value uint64) {
	objc.Send[struct{}](m.ID, objc.Sel("setCommand:"), value)
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkPacket/doubleBuffer
func (m MLNetworkPacket) DoubleBuffer() foundation.NSMutableData {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("doubleBuffer"))
	return foundation.NSMutableDataFromID(objc.ID(rv))
}
func (m MLNetworkPacket) SetDoubleBuffer(value foundation.NSMutableData) {
	objc.Send[struct{}](m.ID, objc.Sel("setDoubleBuffer:"), value)
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkPacket/sizeOfPacket
func (m MLNetworkPacket) SizeOfPacket() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("sizeOfPacket"))
	return rv
}
func (m MLNetworkPacket) SetSizeOfPacket(value uint64) {
	objc.Send[struct{}](m.ID, objc.Sel("setSizeOfPacket:"), value)
}
