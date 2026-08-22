// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [RamXPC] class.
var (
	_RamXPCClass     RamXPCClass
	_RamXPCClassOnce sync.Once
)

func getRamXPCClass() RamXPCClass {
	_RamXPCClassOnce.Do(func() {
		_RamXPCClass = RamXPCClass{class: objc.GetClass("RamXPC")}
	})
	return _RamXPCClass
}

// GetRamXPCClass returns the class object for RamXPC.
func GetRamXPCClass() RamXPCClass {
	return getRamXPCClass()
}

type RamXPCClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (rc RamXPCClass) Class() objc.Class {
	return rc.class
}

// Alloc allocates memory for a new instance of the class.
func (rc RamXPCClass) Alloc() RamXPC {
	rv := objc.SendIfResponds[RamXPC](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [RamXPC.CreateRamBackend]
//   - [RamXPC.InitWithSize]
type RamXPC struct {
	BackendXPC
}

// RamXPCFromID constructs a [RamXPC] from an objc.ID.
func RamXPCFromID(id objc.ID) RamXPC {
	return RamXPC{BackendXPC: BackendXPCFromID(id)}
}

// Ensure RamXPC implements IRamXPC.
var _ IRamXPC = RamXPC{}

// An interface definition for the [RamXPC] class.
//
// # Methods
//
//   - [IRamXPC.CreateRamBackend]
//   - [IRamXPC.InitWithSize]
type IRamXPC interface {
	IBackendXPC

	// Topic: Methods

	CreateRamBackend()
	InitWithSize(size uint64) RamXPC
}

// Init initializes the instance.
func (r RamXPC) Init() RamXPC {
	rv := objc.SendIfResponds[RamXPC](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r RamXPC) Autorelease() RamXPC {
	rv := objc.SendIfResponds[RamXPC](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewRamXPC creates a new RamXPC instance.
func NewRamXPC() RamXPC {
	class := getRamXPCClass()
	rv := objc.SendIfResponds[RamXPC](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewRamXPCWithCoder(coder objectivec.IObject) RamXPC {
	instance := getRamXPCClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return RamXPCFromID(rv)
}

func NewRamXPCWithSize(size uint64) RamXPC {
	instance := getRamXPCClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithSize:"), size)
	return RamXPCFromID(rv)
}

func (r RamXPC) CreateRamBackend() {
	objc.SendIfResponds[objc.ID](r.ID, objc.Sel("createRamBackend"))
}
func (r RamXPC) InitWithSize(size uint64) RamXPC {
	rv := objc.SendIfResponds[RamXPC](r.ID, objc.Sel("initWithSize:"), size)
	return rv
}
