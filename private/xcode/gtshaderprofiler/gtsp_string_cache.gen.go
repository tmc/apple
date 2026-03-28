// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTSPStringCache] class.
var (
	_GTSPStringCacheClass     GTSPStringCacheClass
	_GTSPStringCacheClassOnce sync.Once
)

func getGTSPStringCacheClass() GTSPStringCacheClass {
	_GTSPStringCacheClassOnce.Do(func() {
		_GTSPStringCacheClass = GTSPStringCacheClass{class: objc.GetClass("GTSPStringCache")}
	})
	return _GTSPStringCacheClass
}

// GetGTSPStringCacheClass returns the class object for GTSPStringCache.
func GetGTSPStringCacheClass() GTSPStringCacheClass {
	return getGTSPStringCacheClass()
}

type GTSPStringCacheClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTSPStringCacheClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTSPStringCacheClass) Alloc() GTSPStringCache {
	rv := objc.Send[GTSPStringCache](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTSPStringCache
type GTSPStringCache struct {
	GTShaderProfilerStringCache
}

// GTSPStringCacheFromID constructs a [GTSPStringCache] from an objc.ID.
func GTSPStringCacheFromID(id objc.ID) GTSPStringCache {
	return GTSPStringCache{GTShaderProfilerStringCache: GTShaderProfilerStringCacheFromID(id)}
}
// Ensure GTSPStringCache implements IGTSPStringCache.
var _ IGTSPStringCache = GTSPStringCache{}

// An interface definition for the [GTSPStringCache] class.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTSPStringCache
type IGTSPStringCache interface {
	IGTShaderProfilerStringCache
}

// Init initializes the instance.
func (g GTSPStringCache) Init() GTSPStringCache {
	rv := objc.Send[GTSPStringCache](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTSPStringCache) Autorelease() GTSPStringCache {
	rv := objc.Send[GTSPStringCache](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTSPStringCache creates a new GTSPStringCache instance.
func NewGTSPStringCache() GTSPStringCache {
	class := getGTSPStringCacheClass()
	rv := objc.Send[GTSPStringCache](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStringCache/initWithCoder:
func NewGTSPStringCacheWithCoder(coder objectivec.IObject) GTSPStringCache {
	instance := getGTSPStringCacheClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return GTSPStringCacheFromID(rv)
}

