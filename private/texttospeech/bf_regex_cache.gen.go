// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [BFRegexCache] class.
var (
	_BFRegexCacheClass     BFRegexCacheClass
	_BFRegexCacheClassOnce sync.Once
)

func getBFRegexCacheClass() BFRegexCacheClass {
	_BFRegexCacheClassOnce.Do(func() {
		_BFRegexCacheClass = BFRegexCacheClass{class: objc.GetClass("BFRegexCache")}
	})
	return _BFRegexCacheClass
}

// GetBFRegexCacheClass returns the class object for BFRegexCache.
func GetBFRegexCacheClass() BFRegexCacheClass {
	return getBFRegexCacheClass()
}

type BFRegexCacheClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (bc BFRegexCacheClass) Class() objc.Class {
	return bc.class
}

// Alloc allocates memory for a new instance of the class.
func (bc BFRegexCacheClass) Alloc() BFRegexCache {
	rv := objc.SendIfResponds[BFRegexCache](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [BFRegexCache.Cache]
//   - [BFRegexCache.SetCache]
//   - [BFRegexCache.RegexForString]
//   - [BFRegexCache.RegexForStringAtStart]
type BFRegexCache struct {
	objectivec.Object
}

// BFRegexCacheFromID constructs a [BFRegexCache] from an objc.ID.
func BFRegexCacheFromID(id objc.ID) BFRegexCache {
	return BFRegexCache{objectivec.Object{ID: id}}
}

// Ensure BFRegexCache implements IBFRegexCache.
var _ IBFRegexCache = BFRegexCache{}

// An interface definition for the [BFRegexCache] class.
//
// # Methods
//
//   - [IBFRegexCache.Cache]
//   - [IBFRegexCache.SetCache]
//   - [IBFRegexCache.RegexForString]
//   - [IBFRegexCache.RegexForStringAtStart]
type IBFRegexCache interface {
	objectivec.IObject

	// Topic: Methods

	Cache() foundation.INSDictionary
	SetCache(value foundation.INSDictionary)
	RegexForString(string_ objectivec.IObject) objectivec.IObject
	RegexForStringAtStart(string_ objectivec.IObject, start bool) objectivec.IObject
}

// Init initializes the instance.
func (b BFRegexCache) Init() BFRegexCache {
	rv := objc.SendIfResponds[BFRegexCache](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b BFRegexCache) Autorelease() BFRegexCache {
	rv := objc.SendIfResponds[BFRegexCache](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewBFRegexCache creates a new BFRegexCache instance.
func NewBFRegexCache() BFRegexCache {
	class := getBFRegexCacheClass()
	rv := objc.SendIfResponds[BFRegexCache](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (b BFRegexCache) RegexForString(string_ objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("regexForString:"), string_)
	return objectivec.Object{ID: rv}
}
func (b BFRegexCache) RegexForStringAtStart(string_ objectivec.IObject, start bool) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("regexForString:atStart:"), string_, start)
	return objectivec.Object{ID: rv}
}

func (_BFRegexCacheClass BFRegexCacheClass) SharedInstance() BFRegexCache {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_BFRegexCacheClass.class), objc.Sel("sharedInstance"))
	return BFRegexCacheFromID(rv)
}

func (b BFRegexCache) Cache() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("cache"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (b BFRegexCache) SetCache(value foundation.INSDictionary) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setCache:"), value)
}
