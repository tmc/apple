// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [PluginBackendXPC] class.
var (
	_PluginBackendXPCClass     PluginBackendXPCClass
	_PluginBackendXPCClassOnce sync.Once
)

func getPluginBackendXPCClass() PluginBackendXPCClass {
	_PluginBackendXPCClassOnce.Do(func() {
		_PluginBackendXPCClass = PluginBackendXPCClass{class: objc.GetClass("PluginBackendXPC")}
	})
	return _PluginBackendXPCClass
}

// GetPluginBackendXPCClass returns the class object for PluginBackendXPC.
func GetPluginBackendXPCClass() PluginBackendXPCClass {
	return getPluginBackendXPCClass()
}

type PluginBackendXPCClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PluginBackendXPCClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PluginBackendXPCClass) Alloc() PluginBackendXPC {
	rv := objc.SendIfResponds[PluginBackendXPC](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [PluginBackendXPC.URL]
//   - [PluginBackendXPC.PluginHeader]
//   - [PluginBackendXPC.InitWithURLOpenMode]
type PluginBackendXPC struct {
	BackendXPC
}

// PluginBackendXPCFromID constructs a [PluginBackendXPC] from an objc.ID.
func PluginBackendXPCFromID(id objc.ID) PluginBackendXPC {
	return PluginBackendXPC{BackendXPC: BackendXPCFromID(id)}
}

// Ensure PluginBackendXPC implements IPluginBackendXPC.
var _ IPluginBackendXPC = PluginBackendXPC{}

// An interface definition for the [PluginBackendXPC] class.
//
// # Methods
//
//   - [IPluginBackendXPC.URL]
//   - [IPluginBackendXPC.PluginHeader]
//   - [IPluginBackendXPC.InitWithURLOpenMode]
type IPluginBackendXPC interface {
	IBackendXPC

	// Topic: Methods

	URL() IDIURL
	PluginHeader() unsafe.Pointer
	InitWithURLOpenMode(url foundation.NSURL, mode int) PluginBackendXPC
}

// Init initializes the instance.
func (p PluginBackendXPC) Init() PluginBackendXPC {
	rv := objc.SendIfResponds[PluginBackendXPC](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PluginBackendXPC) Autorelease() PluginBackendXPC {
	rv := objc.SendIfResponds[PluginBackendXPC](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPluginBackendXPC creates a new PluginBackendXPC instance.
func NewPluginBackendXPC() PluginBackendXPC {
	class := getPluginBackendXPCClass()
	rv := objc.SendIfResponds[PluginBackendXPC](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewPluginBackendXPCWithCoder(coder objectivec.IObject) PluginBackendXPC {
	instance := getPluginBackendXPCClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return PluginBackendXPCFromID(rv)
}

func NewPluginBackendXPCWithURLOpenMode(url foundation.NSURL, mode int) PluginBackendXPC {
	instance := getPluginBackendXPCClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithURL:openMode:"), url, mode)
	return PluginBackendXPCFromID(rv)
}

func (p PluginBackendXPC) PluginHeader() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](p.ID, objc.Sel("pluginHeader"))
	return rv
}
func (p PluginBackendXPC) InitWithURLOpenMode(url foundation.NSURL, mode int) PluginBackendXPC {
	rv := objc.SendIfResponds[PluginBackendXPC](p.ID, objc.Sel("initWithURL:openMode:"), url, mode)
	return rv
}

func (p PluginBackendXPC) URL() IDIURL {
	rv := objc.SendIfResponds[objc.ID](p.ID, objc.Sel("URL"))
	return DIURLFromID(objc.ID(rv))
}
