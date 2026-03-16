// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [EspressoSharedKernelCacheEntry] class.
var (
	_EspressoSharedKernelCacheEntryClass     EspressoSharedKernelCacheEntryClass
	_EspressoSharedKernelCacheEntryClassOnce sync.Once
)

func getEspressoSharedKernelCacheEntryClass() EspressoSharedKernelCacheEntryClass {
	_EspressoSharedKernelCacheEntryClassOnce.Do(func() {
		_EspressoSharedKernelCacheEntryClass = EspressoSharedKernelCacheEntryClass{class: objc.GetClass("EspressoSharedKernelCacheEntry")}
	})
	return _EspressoSharedKernelCacheEntryClass
}

// GetEspressoSharedKernelCacheEntryClass returns the class object for EspressoSharedKernelCacheEntry.
func GetEspressoSharedKernelCacheEntryClass() EspressoSharedKernelCacheEntryClass {
	return getEspressoSharedKernelCacheEntryClass()
}

type EspressoSharedKernelCacheEntryClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoSharedKernelCacheEntryClass) Alloc() EspressoSharedKernelCacheEntry {
	rv := objc.Send[EspressoSharedKernelCacheEntry](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [EspressoSharedKernelCacheEntry.Cached]
//   - [EspressoSharedKernelCacheEntry.SetCached]
// See: https://developer.apple.com/documentation/Espresso/EspressoSharedKernelCacheEntry
type EspressoSharedKernelCacheEntry struct {
	objectivec.Object
}

// EspressoSharedKernelCacheEntryFromID constructs a [EspressoSharedKernelCacheEntry] from an objc.ID.
func EspressoSharedKernelCacheEntryFromID(id objc.ID) EspressoSharedKernelCacheEntry {
	return EspressoSharedKernelCacheEntry{objectivec.Object{id}}
}
// Ensure EspressoSharedKernelCacheEntry implements IEspressoSharedKernelCacheEntry.
var _ IEspressoSharedKernelCacheEntry = EspressoSharedKernelCacheEntry{}

// An interface definition for the [EspressoSharedKernelCacheEntry] class.
//
// # Methods
//
//   - [IEspressoSharedKernelCacheEntry.Cached]
//   - [IEspressoSharedKernelCacheEntry.SetCached]
//
// See: https://developer.apple.com/documentation/Espresso/EspressoSharedKernelCacheEntry
type IEspressoSharedKernelCacheEntry interface {
	objectivec.IObject

	// Topic: Methods

	Cached() IEspressoMetalKernelsCache
	SetCached(value IEspressoMetalKernelsCache)
}

// Init initializes the instance.
func (e EspressoSharedKernelCacheEntry) Init() EspressoSharedKernelCacheEntry {
	rv := objc.Send[EspressoSharedKernelCacheEntry](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoSharedKernelCacheEntry) Autorelease() EspressoSharedKernelCacheEntry {
	rv := objc.Send[EspressoSharedKernelCacheEntry](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoSharedKernelCacheEntry creates a new EspressoSharedKernelCacheEntry instance.
func NewEspressoSharedKernelCacheEntry() EspressoSharedKernelCacheEntry {
	class := getEspressoSharedKernelCacheEntryClass()
	rv := objc.Send[EspressoSharedKernelCacheEntry](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoSharedKernelCacheEntry/cached
func (e EspressoSharedKernelCacheEntry) Cached() IEspressoMetalKernelsCache {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("cached"))
	return EspressoMetalKernelsCacheFromID(objc.ID(rv))
}
func (e EspressoSharedKernelCacheEntry) SetCached(value IEspressoMetalKernelsCache) {
	objc.Send[struct{}](e.ID, objc.Sel("setCached:"), value)
}

