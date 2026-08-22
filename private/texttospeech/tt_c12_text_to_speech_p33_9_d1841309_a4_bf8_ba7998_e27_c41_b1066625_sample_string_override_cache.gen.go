// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SampleStringOverrideCache] class.
var (
	_SampleStringOverrideCacheClass     SampleStringOverrideCacheClass
	_SampleStringOverrideCacheClassOnce sync.Once
)

func getSampleStringOverrideCacheClass() SampleStringOverrideCacheClass {
	_SampleStringOverrideCacheClassOnce.Do(func() {
		_SampleStringOverrideCacheClass = SampleStringOverrideCacheClass{class: objc.GetClass("_TtC12TextToSpeechP33_9D1841309A4BF8BA7998E27C41B1066625SampleStringOverrideCache")}
	})
	return _SampleStringOverrideCacheClass
}

// GetSampleStringOverrideCacheClass returns the class object for _TtC12TextToSpeechP33_9D1841309A4BF8BA7998E27C41B1066625SampleStringOverrideCache.
func GetSampleStringOverrideCacheClass() SampleStringOverrideCacheClass {
	return getSampleStringOverrideCacheClass()
}

type SampleStringOverrideCacheClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SampleStringOverrideCacheClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SampleStringOverrideCacheClass) Alloc() SampleStringOverrideCache {
	rv := objc.SendIfResponds[SampleStringOverrideCache](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

type SampleStringOverrideCache struct {
	objectivec.Object
}

// SampleStringOverrideCacheFromID constructs a [SampleStringOverrideCache] from an objc.ID.
func SampleStringOverrideCacheFromID(id objc.ID) SampleStringOverrideCache {
	return SampleStringOverrideCache{objectivec.Object{ID: id}}
}

// Ensure SampleStringOverrideCache implements ISampleStringOverrideCache.
var _ ISampleStringOverrideCache = SampleStringOverrideCache{}

// An interface definition for the [SampleStringOverrideCache] class.
type ISampleStringOverrideCache interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s SampleStringOverrideCache) Init() SampleStringOverrideCache {
	rv := objc.SendIfResponds[SampleStringOverrideCache](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SampleStringOverrideCache) Autorelease() SampleStringOverrideCache {
	rv := objc.SendIfResponds[SampleStringOverrideCache](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSampleStringOverrideCache creates a new SampleStringOverrideCache instance.
func NewSampleStringOverrideCache() SampleStringOverrideCache {
	class := getSampleStringOverrideCacheClass()
	rv := objc.SendIfResponds[SampleStringOverrideCache](objc.ID(class.class), objc.Sel("new"))
	return rv
}
