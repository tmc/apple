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
	rv := objc.Send[PluginBackendXPC](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [PluginBackendXPC.URL]
//   - [PluginBackendXPC.PluginHeader]
//   - [PluginBackendXPC.InitWithURLOpenMode]
//
// See: https://developer.apple.com/documentation/DiskImages2/PluginBackendXPC
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
//
// See: https://developer.apple.com/documentation/DiskImages2/PluginBackendXPC
type IPluginBackendXPC interface {
	IBackendXPC

	// Topic: Methods

	URL() IDIURL
	PluginHeader() unsafe.Pointer
	InitWithURLOpenMode(url foundation.INSURL, mode int) PluginBackendXPC
}

// Init initializes the instance.
func (p PluginBackendXPC) Init() PluginBackendXPC {
	rv := objc.Send[PluginBackendXPC](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PluginBackendXPC) Autorelease() PluginBackendXPC {
	rv := objc.Send[PluginBackendXPC](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPluginBackendXPC creates a new PluginBackendXPC instance.
func NewPluginBackendXPC() PluginBackendXPC {
	class := getPluginBackendXPCClass()
	rv := objc.Send[PluginBackendXPC](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/PluginBackendXPC/initWithCoder:
func NewPluginBackendXPCWithCoder(coder objectivec.IObject) PluginBackendXPC {
	instance := getPluginBackendXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return PluginBackendXPCFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/PluginBackendXPC/initWithURL:openMode:
func NewPluginBackendXPCWithURLOpenMode(url foundation.INSURL, mode int) PluginBackendXPC {
	instance := getPluginBackendXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:openMode:"), url, mode)
	return PluginBackendXPCFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/PluginBackendXPC/pluginHeader
func (p PluginBackendXPC) PluginHeader() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p.ID, objc.Sel("pluginHeader"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/PluginBackendXPC/initWithURL:openMode:
func (p PluginBackendXPC) InitWithURLOpenMode(url foundation.INSURL, mode int) PluginBackendXPC {
	rv := objc.Send[PluginBackendXPC](p.ID, objc.Sel("initWithURL:openMode:"), url, mode)
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/PluginBackendXPC/URL
func (p PluginBackendXPC) URL() IDIURL {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("URL"))
	return DIURLFromID(objc.ID(rv))
}
