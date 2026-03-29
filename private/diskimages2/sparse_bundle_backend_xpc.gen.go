// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SparseBundleBackendXPC] class.
var (
	_SparseBundleBackendXPCClass     SparseBundleBackendXPCClass
	_SparseBundleBackendXPCClassOnce sync.Once
)

func getSparseBundleBackendXPCClass() SparseBundleBackendXPCClass {
	_SparseBundleBackendXPCClassOnce.Do(func() {
		_SparseBundleBackendXPCClass = SparseBundleBackendXPCClass{class: objc.GetClass("SparseBundleBackendXPC")}
	})
	return _SparseBundleBackendXPCClass
}

// GetSparseBundleBackendXPCClass returns the class object for SparseBundleBackendXPC.
func GetSparseBundleBackendXPCClass() SparseBundleBackendXPCClass {
	return getSparseBundleBackendXPCClass()
}

type SparseBundleBackendXPCClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SparseBundleBackendXPCClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SparseBundleBackendXPCClass) Alloc() SparseBundleBackendXPC {
	rv := objc.Send[SparseBundleBackendXPC](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [SparseBundleBackendXPC.InitWithURLFileOpenFlags]
//   - [SparseBundleBackendXPC.InitWithURLFileOpenFlagsBandSize]
// See: https://developer.apple.com/documentation/DiskImages2/SparseBundleBackendXPC
type SparseBundleBackendXPC struct {
	BackendXPC
}

// SparseBundleBackendXPCFromID constructs a [SparseBundleBackendXPC] from an objc.ID.
func SparseBundleBackendXPCFromID(id objc.ID) SparseBundleBackendXPC {
	return SparseBundleBackendXPC{BackendXPC: BackendXPCFromID(id)}
}
// Ensure SparseBundleBackendXPC implements ISparseBundleBackendXPC.
var _ ISparseBundleBackendXPC = SparseBundleBackendXPC{}

// An interface definition for the [SparseBundleBackendXPC] class.
//
// # Methods
//
//   - [ISparseBundleBackendXPC.InitWithURLFileOpenFlags]
//   - [ISparseBundleBackendXPC.InitWithURLFileOpenFlagsBandSize]
//
// See: https://developer.apple.com/documentation/DiskImages2/SparseBundleBackendXPC
type ISparseBundleBackendXPC interface {
	IBackendXPC

	// Topic: Methods

	InitWithURLFileOpenFlags(url foundation.INSURL, flags int) SparseBundleBackendXPC
	InitWithURLFileOpenFlagsBandSize(url foundation.INSURL, flags int, size uint64) SparseBundleBackendXPC
}

// Init initializes the instance.
func (s SparseBundleBackendXPC) Init() SparseBundleBackendXPC {
	rv := objc.Send[SparseBundleBackendXPC](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SparseBundleBackendXPC) Autorelease() SparseBundleBackendXPC {
	rv := objc.Send[SparseBundleBackendXPC](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSparseBundleBackendXPC creates a new SparseBundleBackendXPC instance.
func NewSparseBundleBackendXPC() SparseBundleBackendXPC {
	class := getSparseBundleBackendXPCClass()
	rv := objc.Send[SparseBundleBackendXPC](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/DiskImages2/SparseBundleBackendXPC/initWithCoder:
func NewSparseBundleBackendXPCWithCoder(coder objectivec.IObject) SparseBundleBackendXPC {
	instance := getSparseBundleBackendXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SparseBundleBackendXPCFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/SparseBundleBackendXPC/initWithURL:fileOpenFlags:
func NewSparseBundleBackendXPCWithURLFileOpenFlags(url foundation.INSURL, flags int) SparseBundleBackendXPC {
	instance := getSparseBundleBackendXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:fileOpenFlags:"), url, flags)
	return SparseBundleBackendXPCFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/SparseBundleBackendXPC/initWithURL:fileOpenFlags:bandSize:
func NewSparseBundleBackendXPCWithURLFileOpenFlagsBandSize(url foundation.INSURL, flags int, size uint64) SparseBundleBackendXPC {
	instance := getSparseBundleBackendXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:fileOpenFlags:bandSize:"), url, flags, size)
	return SparseBundleBackendXPCFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/SparseBundleBackendXPC/initWithURL:fileOpenFlags:
func (s SparseBundleBackendXPC) InitWithURLFileOpenFlags(url foundation.INSURL, flags int) SparseBundleBackendXPC {
	rv := objc.Send[SparseBundleBackendXPC](s.ID, objc.Sel("initWithURL:fileOpenFlags:"), url, flags)
	return rv
}
//
// See: https://developer.apple.com/documentation/DiskImages2/SparseBundleBackendXPC/initWithURL:fileOpenFlags:bandSize:
func (s SparseBundleBackendXPC) InitWithURLFileOpenFlagsBandSize(url foundation.INSURL, flags int, size uint64) SparseBundleBackendXPC {
	rv := objc.Send[SparseBundleBackendXPC](s.ID, objc.Sel("initWithURL:fileOpenFlags:bandSize:"), url, flags, size)
	return rv
}

//
// See: https://developer.apple.com/documentation/DiskImages2/SparseBundleBackendXPC/isSparseBundleWithURL:
func (_SparseBundleBackendXPCClass SparseBundleBackendXPCClass) IsSparseBundleWithURL(url foundation.INSURL) bool {
	rv := objc.Send[bool](objc.ID(_SparseBundleBackendXPCClass.class), objc.Sel("isSparseBundleWithURL:"), url)
	return rv
}

