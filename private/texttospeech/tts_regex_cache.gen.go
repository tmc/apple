// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSRegexCache] class.
var (
	_TTSRegexCacheClass     TTSRegexCacheClass
	_TTSRegexCacheClassOnce sync.Once
)

func getTTSRegexCacheClass() TTSRegexCacheClass {
	_TTSRegexCacheClassOnce.Do(func() {
		_TTSRegexCacheClass = TTSRegexCacheClass{class: objc.GetClass("TTSRegexCache")}
	})
	return _TTSRegexCacheClass
}

// GetTTSRegexCacheClass returns the class object for TTSRegexCache.
func GetTTSRegexCacheClass() TTSRegexCacheClass {
	return getTTSRegexCacheClass()
}

type TTSRegexCacheClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSRegexCacheClass) Alloc() TTSRegexCache {
	rv := objc.Send[TTSRegexCache](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [TTSRegexCache.Cache]
//   - [TTSRegexCache.SetCache]
//   - [TTSRegexCache.RegexForString]
//   - [TTSRegexCache.RegexForStringAtStart]
// See: https://developer.apple.com/documentation/TextToSpeech/TTSRegexCache
type TTSRegexCache struct {
	objectivec.Object
}

// TTSRegexCacheFromID constructs a [TTSRegexCache] from an objc.ID.
func TTSRegexCacheFromID(id objc.ID) TTSRegexCache {
	return TTSRegexCache{objectivec.Object{ID: id}}
}
// Ensure TTSRegexCache implements ITTSRegexCache.
var _ ITTSRegexCache = TTSRegexCache{}

// An interface definition for the [TTSRegexCache] class.
//
// # Methods
//
//   - [ITTSRegexCache.Cache]
//   - [ITTSRegexCache.SetCache]
//   - [ITTSRegexCache.RegexForString]
//   - [ITTSRegexCache.RegexForStringAtStart]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSRegexCache
type ITTSRegexCache interface {
	objectivec.IObject

	// Topic: Methods

	Cache() foundation.INSDictionary
	SetCache(value foundation.INSDictionary)
	RegexForString(string_ objectivec.IObject) objectivec.IObject
	RegexForStringAtStart(string_ objectivec.IObject, start bool) objectivec.IObject
}

// Init initializes the instance.
func (t TTSRegexCache) Init() TTSRegexCache {
	rv := objc.Send[TTSRegexCache](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSRegexCache) Autorelease() TTSRegexCache {
	rv := objc.Send[TTSRegexCache](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSRegexCache creates a new TTSRegexCache instance.
func NewTTSRegexCache() TTSRegexCache {
	class := getTTSRegexCacheClass()
	rv := objc.Send[TTSRegexCache](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSRegexCache/regexForString:
func (t TTSRegexCache) RegexForString(string_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("regexForString:"), string_)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSRegexCache/regexForString:atStart:
func (t TTSRegexCache) RegexForStringAtStart(string_ objectivec.IObject, start bool) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("regexForString:atStart:"), string_, start)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRegexCache/sharedInstance
func (_TTSRegexCacheClass TTSRegexCacheClass) SharedInstance() TTSRegexCache {
	rv := objc.Send[objc.ID](objc.ID(_TTSRegexCacheClass.class), objc.Sel("sharedInstance"))
	return TTSRegexCacheFromID(rv)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRegexCache/cache
func (t TTSRegexCache) Cache() foundation.INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("cache"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (t TTSRegexCache) SetCache(value foundation.INSDictionary) {
	objc.Send[struct{}](t.ID, objc.Sel("setCache:"), value)
}

