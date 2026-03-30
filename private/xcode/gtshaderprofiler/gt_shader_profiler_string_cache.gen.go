// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTShaderProfilerStringCache] class.
var (
	_GTShaderProfilerStringCacheClass     GTShaderProfilerStringCacheClass
	_GTShaderProfilerStringCacheClassOnce sync.Once
)

func getGTShaderProfilerStringCacheClass() GTShaderProfilerStringCacheClass {
	_GTShaderProfilerStringCacheClassOnce.Do(func() {
		_GTShaderProfilerStringCacheClass = GTShaderProfilerStringCacheClass{class: objc.GetClass("GTShaderProfilerStringCache")}
	})
	return _GTShaderProfilerStringCacheClass
}

// GetGTShaderProfilerStringCacheClass returns the class object for GTShaderProfilerStringCache.
func GetGTShaderProfilerStringCacheClass() GTShaderProfilerStringCacheClass {
	return getGTShaderProfilerStringCacheClass()
}

type GTShaderProfilerStringCacheClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTShaderProfilerStringCacheClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTShaderProfilerStringCacheClass) Alloc() GTShaderProfilerStringCache {
	rv := objc.Send[GTShaderProfilerStringCache](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTShaderProfilerStringCache.AddString]
//   - [GTShaderProfilerStringCache.EncodeWithCoder]
//   - [GTShaderProfilerStringCache.StringFromIndex]
//   - [GTShaderProfilerStringCache.Strings]
//   - [GTShaderProfilerStringCache.InitWithCoder]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStringCache
type GTShaderProfilerStringCache struct {
	objectivec.Object
}

// GTShaderProfilerStringCacheFromID constructs a [GTShaderProfilerStringCache] from an objc.ID.
func GTShaderProfilerStringCacheFromID(id objc.ID) GTShaderProfilerStringCache {
	return GTShaderProfilerStringCache{objectivec.Object{ID: id}}
}

// Ensure GTShaderProfilerStringCache implements IGTShaderProfilerStringCache.
var _ IGTShaderProfilerStringCache = GTShaderProfilerStringCache{}

// An interface definition for the [GTShaderProfilerStringCache] class.
//
// # Methods
//
//   - [IGTShaderProfilerStringCache.AddString]
//   - [IGTShaderProfilerStringCache.EncodeWithCoder]
//   - [IGTShaderProfilerStringCache.StringFromIndex]
//   - [IGTShaderProfilerStringCache.Strings]
//   - [IGTShaderProfilerStringCache.InitWithCoder]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStringCache
type IGTShaderProfilerStringCache interface {
	objectivec.IObject

	// Topic: Methods

	AddString(string_ objectivec.IObject) uint64
	EncodeWithCoder(coder foundation.INSCoder)
	StringFromIndex(index uint64) objectivec.IObject
	Strings() foundation.INSArray
	InitWithCoder(coder foundation.INSCoder) GTShaderProfilerStringCache
}

// Init initializes the instance.
func (g GTShaderProfilerStringCache) Init() GTShaderProfilerStringCache {
	rv := objc.Send[GTShaderProfilerStringCache](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTShaderProfilerStringCache) Autorelease() GTShaderProfilerStringCache {
	rv := objc.Send[GTShaderProfilerStringCache](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTShaderProfilerStringCache creates a new GTShaderProfilerStringCache instance.
func NewGTShaderProfilerStringCache() GTShaderProfilerStringCache {
	class := getGTShaderProfilerStringCacheClass()
	rv := objc.Send[GTShaderProfilerStringCache](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStringCache/initWithCoder:
func NewGTShaderProfilerStringCacheWithCoder(coder objectivec.IObject) GTShaderProfilerStringCache {
	instance := getGTShaderProfilerStringCacheClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return GTShaderProfilerStringCacheFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStringCache/addString:
func (g GTShaderProfilerStringCache) AddString(string_ objectivec.IObject) uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("addString:"), string_)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStringCache/encodeWithCoder:
func (g GTShaderProfilerStringCache) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](g.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStringCache/stringFromIndex:
func (g GTShaderProfilerStringCache) StringFromIndex(index uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("stringFromIndex:"), index)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStringCache/initWithCoder:
func (g GTShaderProfilerStringCache) InitWithCoder(coder foundation.INSCoder) GTShaderProfilerStringCache {
	rv := objc.Send[GTShaderProfilerStringCache](g.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStringCache/supportsSecureCoding
func (_GTShaderProfilerStringCacheClass GTShaderProfilerStringCacheClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_GTShaderProfilerStringCacheClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStringCache/strings
func (g GTShaderProfilerStringCache) Strings() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("strings"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
