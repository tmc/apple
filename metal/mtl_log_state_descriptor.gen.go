// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLLogStateDescriptor] class.
var (
	_MTLLogStateDescriptorClass     MTLLogStateDescriptorClass
	_MTLLogStateDescriptorClassOnce sync.Once
)

func getMTLLogStateDescriptorClass() MTLLogStateDescriptorClass {
	_MTLLogStateDescriptorClassOnce.Do(func() {
		_MTLLogStateDescriptorClass = MTLLogStateDescriptorClass{class: objc.GetClass("MTLLogStateDescriptor")}
	})
	return _MTLLogStateDescriptorClass
}

// GetMTLLogStateDescriptorClass returns the class object for MTLLogStateDescriptor.
func GetMTLLogStateDescriptorClass() MTLLogStateDescriptorClass {
	return getMTLLogStateDescriptorClass()
}

type MTLLogStateDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLLogStateDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLLogStateDescriptorClass) Alloc() MTLLogStateDescriptor {
	rv := objc.Send[MTLLogStateDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An interface that represents a log state configuration.
//
// # Overview
//
// Configure the descriptor to create an [MTLLogState] by calling
// [NewLogStateWithDescriptorError].
//
// If you’ve set the environment variables `MTL_LOG_BUFFER_SIZE` or
// `MTL_LOG_LEVEL`, then the system automatically enables logging. If any
// command buffer or command queue has an attached log state, then the system
// uses the log state’s settings instead of the environment variable values.
//
// # Instance properties
//
//   - [MTLLogStateDescriptor.BufferSize]: The size of the internal buffer the log state uses, specified in bytes.
//   - [MTLLogStateDescriptor.SetBufferSize]
//   - [MTLLogStateDescriptor.Level]: The minimum level of messages that the shader can log.
//   - [MTLLogStateDescriptor.SetLevel]
//
// See: https://developer.apple.com/documentation/Metal/MTLLogStateDescriptor
type MTLLogStateDescriptor struct {
	objectivec.Object
}

// MTLLogStateDescriptorFromID constructs a [MTLLogStateDescriptor] from an objc.ID.
//
// An interface that represents a log state configuration.
func MTLLogStateDescriptorFromID(id objc.ID) MTLLogStateDescriptor {
	return MTLLogStateDescriptor{objectivec.Object{ID: id}}
}

// NOTE: MTLLogStateDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLLogStateDescriptor] class.
//
// # Instance properties
//
//   - [IMTLLogStateDescriptor.BufferSize]: The size of the internal buffer the log state uses, specified in bytes.
//   - [IMTLLogStateDescriptor.SetBufferSize]
//   - [IMTLLogStateDescriptor.Level]: The minimum level of messages that the shader can log.
//   - [IMTLLogStateDescriptor.SetLevel]
//
// See: https://developer.apple.com/documentation/Metal/MTLLogStateDescriptor
type IMTLLogStateDescriptor interface {
	objectivec.IObject

	// Topic: Instance properties

	// The size of the internal buffer the log state uses, specified in bytes.
	BufferSize() int
	SetBufferSize(value int)
	// The minimum level of messages that the shader can log.
	Level() MTLLogLevel
	SetLevel(value MTLLogLevel)
}

// Init initializes the instance.
func (l MTLLogStateDescriptor) Init() MTLLogStateDescriptor {
	rv := objc.Send[MTLLogStateDescriptor](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l MTLLogStateDescriptor) Autorelease() MTLLogStateDescriptor {
	rv := objc.Send[MTLLogStateDescriptor](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLLogStateDescriptor creates a new MTLLogStateDescriptor instance.
func NewMTLLogStateDescriptor() MTLLogStateDescriptor {
	class := getMTLLogStateDescriptorClass()
	rv := objc.Send[MTLLogStateDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The size of the internal buffer the log state uses, specified in bytes.
//
// # Discussion
//
// The default value is 1MB. The minimum size of log buffer is 1KB and the
// maximum size is 1GB.
//
// Carefully consider the size of this buffer based on how many messages you
// expect your shader to log and be useful to diagnose problems. A smaller
// size might lead to the shader dropping more messages while a larger size
// might result in a larger memory footprint and reduced performance due to
// excessive logging.
//
// See: https://developer.apple.com/documentation/Metal/MTLLogStateDescriptor/bufferSize
func (l MTLLogStateDescriptor) BufferSize() int {
	rv := objc.Send[int](l.ID, objc.Sel("bufferSize"))
	return rv
}
func (l MTLLogStateDescriptor) SetBufferSize(value int) {
	objc.Send[struct{}](l.ID, objc.Sel("setBufferSize:"), value)
}

// The minimum level of messages that the shader can log.
//
// # Discussion
//
// The default value is [MTLLogLevel.debug].
//
// Use this value to limit which logs from your shader the log state stores.
// The log state doesn’t store messages at a lower level. Increase the level
// to reduce verbosity of logging.
//
// See: https://developer.apple.com/documentation/Metal/MTLLogStateDescriptor/level
//
// [MTLLogLevel.debug]: https://developer.apple.com/documentation/Metal/MTLLogLevel/debug
func (l MTLLogStateDescriptor) Level() MTLLogLevel {
	rv := objc.Send[MTLLogLevel](l.ID, objc.Sel("level"))
	return MTLLogLevel(rv)
}
func (l MTLLogStateDescriptor) SetLevel(value MTLLogLevel) {
	objc.Send[struct{}](l.ID, objc.Sel("setLevel:"), value)
}
