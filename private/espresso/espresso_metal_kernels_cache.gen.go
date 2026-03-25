// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [EspressoMetalKernelsCache] class.
var (
	_EspressoMetalKernelsCacheClass     EspressoMetalKernelsCacheClass
	_EspressoMetalKernelsCacheClassOnce sync.Once
)

func getEspressoMetalKernelsCacheClass() EspressoMetalKernelsCacheClass {
	_EspressoMetalKernelsCacheClassOnce.Do(func() {
		_EspressoMetalKernelsCacheClass = EspressoMetalKernelsCacheClass{class: objc.GetClass("EspressoMetalKernelsCache")}
	})
	return _EspressoMetalKernelsCacheClass
}

// GetEspressoMetalKernelsCacheClass returns the class object for EspressoMetalKernelsCache.
func GetEspressoMetalKernelsCacheClass() EspressoMetalKernelsCacheClass {
	return getEspressoMetalKernelsCacheClass()
}

type EspressoMetalKernelsCacheClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoMetalKernelsCacheClass) Alloc() EspressoMetalKernelsCache {
	rv := objc.Send[EspressoMetalKernelsCache](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [EspressoMetalKernelsCache.AddLibraryAtPath]
//   - [EspressoMetalKernelsCache.KernelForFunction]
//   - [EspressoMetalKernelsCache.KernelForFunctionCacheStringWithConstants]
//   - [EspressoMetalKernelsCache.KernelPrefix]
//   - [EspressoMetalKernelsCache.SetKernelPrefix]
//   - [EspressoMetalKernelsCache.LazySetup]
//   - [EspressoMetalKernelsCache.LoadLibraryNamed]
//   - [EspressoMetalKernelsCache.M_kernelCache]
//   - [EspressoMetalKernelsCache.SetM_kernelCache]
//   - [EspressoMetalKernelsCache.ShouldUseTexArray]
//   - [EspressoMetalKernelsCache.WasSetup]
//   - [EspressoMetalKernelsCache.InitWithDevice]
// See: https://developer.apple.com/documentation/Espresso/EspressoMetalKernelsCache
type EspressoMetalKernelsCache struct {
	objectivec.Object
}

// EspressoMetalKernelsCacheFromID constructs a [EspressoMetalKernelsCache] from an objc.ID.
func EspressoMetalKernelsCacheFromID(id objc.ID) EspressoMetalKernelsCache {
	return EspressoMetalKernelsCache{objectivec.Object{ID: id}}
}
// Ensure EspressoMetalKernelsCache implements IEspressoMetalKernelsCache.
var _ IEspressoMetalKernelsCache = EspressoMetalKernelsCache{}

// An interface definition for the [EspressoMetalKernelsCache] class.
//
// # Methods
//
//   - [IEspressoMetalKernelsCache.AddLibraryAtPath]
//   - [IEspressoMetalKernelsCache.KernelForFunction]
//   - [IEspressoMetalKernelsCache.KernelForFunctionCacheStringWithConstants]
//   - [IEspressoMetalKernelsCache.KernelPrefix]
//   - [IEspressoMetalKernelsCache.SetKernelPrefix]
//   - [IEspressoMetalKernelsCache.LazySetup]
//   - [IEspressoMetalKernelsCache.LoadLibraryNamed]
//   - [IEspressoMetalKernelsCache.M_kernelCache]
//   - [IEspressoMetalKernelsCache.SetM_kernelCache]
//   - [IEspressoMetalKernelsCache.ShouldUseTexArray]
//   - [IEspressoMetalKernelsCache.WasSetup]
//   - [IEspressoMetalKernelsCache.InitWithDevice]
//
// See: https://developer.apple.com/documentation/Espresso/EspressoMetalKernelsCache
type IEspressoMetalKernelsCache interface {
	objectivec.IObject

	// Topic: Methods

	AddLibraryAtPath(path objectivec.IObject)
	KernelForFunction(function string) objectivec.IObject
	KernelForFunctionCacheStringWithConstants(function string, string_ string, constants objectivec.IObject) objectivec.IObject
	KernelPrefix() string
	SetKernelPrefix(value string)
	LazySetup()
	LoadLibraryNamed(named objectivec.IObject)
	M_kernelCache() foundation.INSDictionary
	SetM_kernelCache(value foundation.INSDictionary)
	ShouldUseTexArray() bool
	WasSetup() bool
	InitWithDevice(device objectivec.IObject) EspressoMetalKernelsCache
}

// Init initializes the instance.
func (e EspressoMetalKernelsCache) Init() EspressoMetalKernelsCache {
	rv := objc.Send[EspressoMetalKernelsCache](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoMetalKernelsCache) Autorelease() EspressoMetalKernelsCache {
	rv := objc.Send[EspressoMetalKernelsCache](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoMetalKernelsCache creates a new EspressoMetalKernelsCache instance.
func NewEspressoMetalKernelsCache() EspressoMetalKernelsCache {
	class := getEspressoMetalKernelsCacheClass()
	rv := objc.Send[EspressoMetalKernelsCache](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoMetalKernelsCache/initWithDevice:
func NewEspressoMetalKernelsCacheWithDevice(device objectivec.IObject) EspressoMetalKernelsCache {
	instance := getEspressoMetalKernelsCacheClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return EspressoMetalKernelsCacheFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoMetalKernelsCache/addLibraryAtPath:
func (e EspressoMetalKernelsCache) AddLibraryAtPath(path objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("addLibraryAtPath:"), path)
}
//
// See: https://developer.apple.com/documentation/Espresso/EspressoMetalKernelsCache/kernelForFunction:
func (e EspressoMetalKernelsCache) KernelForFunction(function string) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("kernelForFunction:"), unsafe.Pointer(unsafe.StringData(function + "\x00")))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Espresso/EspressoMetalKernelsCache/kernelForFunction:cacheString:withConstants:
func (e EspressoMetalKernelsCache) KernelForFunctionCacheStringWithConstants(function string, string_ string, constants objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("kernelForFunction:cacheString:withConstants:"), unsafe.Pointer(unsafe.StringData(function + "\x00")), unsafe.Pointer(unsafe.StringData(string_ + "\x00")), constants)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/Espresso/EspressoMetalKernelsCache/lazySetup
func (e EspressoMetalKernelsCache) LazySetup() {
	objc.Send[objc.ID](e.ID, objc.Sel("lazySetup"))
}
//
// See: https://developer.apple.com/documentation/Espresso/EspressoMetalKernelsCache/loadLibraryNamed:
func (e EspressoMetalKernelsCache) LoadLibraryNamed(named objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("loadLibraryNamed:"), named)
}
// See: https://developer.apple.com/documentation/Espresso/EspressoMetalKernelsCache/shouldUseTexArray
func (e EspressoMetalKernelsCache) ShouldUseTexArray() bool {
	rv := objc.Send[bool](e.ID, objc.Sel("shouldUseTexArray"))
	return rv
}
// See: https://developer.apple.com/documentation/Espresso/EspressoMetalKernelsCache/wasSetup
func (e EspressoMetalKernelsCache) WasSetup() bool {
	rv := objc.Send[bool](e.ID, objc.Sel("wasSetup"))
	return rv
}
//
// See: https://developer.apple.com/documentation/Espresso/EspressoMetalKernelsCache/initWithDevice:
func (e EspressoMetalKernelsCache) InitWithDevice(device objectivec.IObject) EspressoMetalKernelsCache {
	rv := objc.Send[EspressoMetalKernelsCache](e.ID, objc.Sel("initWithDevice:"), device)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoMetalKernelsCache/kernelPrefix
func (e EspressoMetalKernelsCache) KernelPrefix() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("kernelPrefix"))
	return foundation.NSStringFromID(rv).String()
}
func (e EspressoMetalKernelsCache) SetKernelPrefix(value string) {
	objc.Send[struct{}](e.ID, objc.Sel("setKernelPrefix:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/Espresso/EspressoMetalKernelsCache/m_kernelCache
func (e EspressoMetalKernelsCache) M_kernelCache() foundation.INSDictionary {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("m_kernelCache"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (e EspressoMetalKernelsCache) SetM_kernelCache(value foundation.INSDictionary) {
	objc.Send[struct{}](e.ID, objc.Sel("setM_kernelCache:"), value)
}

