// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ReadSynchronizerBackendXPC] class.
var (
	_ReadSynchronizerBackendXPCClass     ReadSynchronizerBackendXPCClass
	_ReadSynchronizerBackendXPCClassOnce sync.Once
)

func getReadSynchronizerBackendXPCClass() ReadSynchronizerBackendXPCClass {
	_ReadSynchronizerBackendXPCClassOnce.Do(func() {
		_ReadSynchronizerBackendXPCClass = ReadSynchronizerBackendXPCClass{class: objc.GetClass("ReadSynchronizerBackendXPC")}
	})
	return _ReadSynchronizerBackendXPCClass
}

// GetReadSynchronizerBackendXPCClass returns the class object for ReadSynchronizerBackendXPC.
func GetReadSynchronizerBackendXPCClass() ReadSynchronizerBackendXPCClass {
	return getReadSynchronizerBackendXPCClass()
}

type ReadSynchronizerBackendXPCClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (rc ReadSynchronizerBackendXPCClass) Class() objc.Class {
	return rc.class
}

// Alloc allocates memory for a new instance of the class.
func (rc ReadSynchronizerBackendXPCClass) Alloc() ReadSynchronizerBackendXPC {
	rv := objc.Send[ReadSynchronizerBackendXPC](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ReadSynchronizerBackendXPC.BaseBackendXPC]
//   - [ReadSynchronizerBackendXPC.InitWithBackend]
//
// See: https://developer.apple.com/documentation/DiskImages2/ReadSynchronizerBackendXPC
type ReadSynchronizerBackendXPC struct {
	BackendXPC
}

// ReadSynchronizerBackendXPCFromID constructs a [ReadSynchronizerBackendXPC] from an objc.ID.
func ReadSynchronizerBackendXPCFromID(id objc.ID) ReadSynchronizerBackendXPC {
	return ReadSynchronizerBackendXPC{BackendXPC: BackendXPCFromID(id)}
}

// Ensure ReadSynchronizerBackendXPC implements IReadSynchronizerBackendXPC.
var _ IReadSynchronizerBackendXPC = ReadSynchronizerBackendXPC{}

// An interface definition for the [ReadSynchronizerBackendXPC] class.
//
// # Methods
//
//   - [IReadSynchronizerBackendXPC.BaseBackendXPC]
//   - [IReadSynchronizerBackendXPC.InitWithBackend]
//
// See: https://developer.apple.com/documentation/DiskImages2/ReadSynchronizerBackendXPC
type IReadSynchronizerBackendXPC interface {
	IBackendXPC

	// Topic: Methods

	BaseBackendXPC() IBackendXPC
	InitWithBackend(backend objectivec.IObject) ReadSynchronizerBackendXPC
}

// Init initializes the instance.
func (r ReadSynchronizerBackendXPC) Init() ReadSynchronizerBackendXPC {
	rv := objc.Send[ReadSynchronizerBackendXPC](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r ReadSynchronizerBackendXPC) Autorelease() ReadSynchronizerBackendXPC {
	rv := objc.Send[ReadSynchronizerBackendXPC](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewReadSynchronizerBackendXPC creates a new ReadSynchronizerBackendXPC instance.
func NewReadSynchronizerBackendXPC() ReadSynchronizerBackendXPC {
	class := getReadSynchronizerBackendXPCClass()
	rv := objc.Send[ReadSynchronizerBackendXPC](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/ReadSynchronizerBackendXPC/initWithBackend:
func NewReadSynchronizerBackendXPCWithBackend(backend objectivec.IObject) ReadSynchronizerBackendXPC {
	instance := getReadSynchronizerBackendXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackend:"), backend)
	return ReadSynchronizerBackendXPCFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/ReadSynchronizerBackendXPC/initWithCoder:
func NewReadSynchronizerBackendXPCWithCoder(coder objectivec.IObject) ReadSynchronizerBackendXPC {
	instance := getReadSynchronizerBackendXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return ReadSynchronizerBackendXPCFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/ReadSynchronizerBackendXPC/initWithBackend:
func (r ReadSynchronizerBackendXPC) InitWithBackend(backend objectivec.IObject) ReadSynchronizerBackendXPC {
	rv := objc.Send[ReadSynchronizerBackendXPC](r.ID, objc.Sel("initWithBackend:"), backend)
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/ReadSynchronizerBackendXPC/baseBackendXPC
func (r ReadSynchronizerBackendXPC) BaseBackendXPC() IBackendXPC {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("baseBackendXPC"))
	return BackendXPCFromID(objc.ID(rv))
}
