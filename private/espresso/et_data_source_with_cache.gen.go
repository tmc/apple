// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ETDataSourceWithCache] class.
var (
	_ETDataSourceWithCacheClass     ETDataSourceWithCacheClass
	_ETDataSourceWithCacheClassOnce sync.Once
)

func getETDataSourceWithCacheClass() ETDataSourceWithCacheClass {
	_ETDataSourceWithCacheClassOnce.Do(func() {
		_ETDataSourceWithCacheClass = ETDataSourceWithCacheClass{class: objc.GetClass("ETDataSourceWithCache")}
	})
	return _ETDataSourceWithCacheClass
}

// GetETDataSourceWithCacheClass returns the class object for ETDataSourceWithCache.
func GetETDataSourceWithCacheClass() ETDataSourceWithCacheClass {
	return getETDataSourceWithCacheClass()
}

type ETDataSourceWithCacheClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec ETDataSourceWithCacheClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec ETDataSourceWithCacheClass) Alloc() ETDataSourceWithCache {
	rv := objc.SendIfResponds[ETDataSourceWithCache](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ETDataSourceWithCache.DataPointAtIndex]
//   - [ETDataSourceWithCache.NumberOfDataPoints]
//   - [ETDataSourceWithCache.InitWithDataSource]
//   - [ETDataSourceWithCache.InitWithDataSourceDumpPath]
type ETDataSourceWithCache struct {
	objectivec.Object
}

// ETDataSourceWithCacheFromID constructs a [ETDataSourceWithCache] from an objc.ID.
func ETDataSourceWithCacheFromID(id objc.ID) ETDataSourceWithCache {
	return ETDataSourceWithCache{objectivec.Object{ID: id}}
}

// Ensure ETDataSourceWithCache implements IETDataSourceWithCache.
var _ IETDataSourceWithCache = ETDataSourceWithCache{}

// An interface definition for the [ETDataSourceWithCache] class.
//
// # Methods
//
//   - [IETDataSourceWithCache.DataPointAtIndex]
//   - [IETDataSourceWithCache.NumberOfDataPoints]
//   - [IETDataSourceWithCache.InitWithDataSource]
//   - [IETDataSourceWithCache.InitWithDataSourceDumpPath]
type IETDataSourceWithCache interface {
	objectivec.IObject

	// Topic: Methods

	DataPointAtIndex(index int) objectivec.IObject
	NumberOfDataPoints() int
	InitWithDataSource(source objectivec.IObject) ETDataSourceWithCache
	InitWithDataSourceDumpPath(source objectivec.IObject, path objectivec.IObject) ETDataSourceWithCache
}

// Init initializes the instance.
func (e ETDataSourceWithCache) Init() ETDataSourceWithCache {
	rv := objc.SendIfResponds[ETDataSourceWithCache](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e ETDataSourceWithCache) Autorelease() ETDataSourceWithCache {
	rv := objc.SendIfResponds[ETDataSourceWithCache](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewETDataSourceWithCache creates a new ETDataSourceWithCache instance.
func NewETDataSourceWithCache() ETDataSourceWithCache {
	class := getETDataSourceWithCacheClass()
	rv := objc.SendIfResponds[ETDataSourceWithCache](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewETDataSourceWithCacheWithDataSource(source objectivec.IObject) ETDataSourceWithCache {
	instance := getETDataSourceWithCacheClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDataSource:"), source)
	return ETDataSourceWithCacheFromID(rv)
}

func NewETDataSourceWithCacheWithDataSourceDumpPath(source objectivec.IObject, path objectivec.IObject) ETDataSourceWithCache {
	instance := getETDataSourceWithCacheClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDataSource:dumpPath:"), source, path)
	return ETDataSourceWithCacheFromID(rv)
}

func (e ETDataSourceWithCache) DataPointAtIndex(index int) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](e.ID, objc.Sel("dataPointAtIndex:"), index)
	return objectivec.Object{ID: rv}
}
func (e ETDataSourceWithCache) NumberOfDataPoints() int {
	rv := objc.SendIfResponds[int](e.ID, objc.Sel("numberOfDataPoints"))
	return rv
}
func (e ETDataSourceWithCache) InitWithDataSource(source objectivec.IObject) ETDataSourceWithCache {
	rv := objc.SendIfResponds[ETDataSourceWithCache](e.ID, objc.Sel("initWithDataSource:"), source)
	return rv
}
func (e ETDataSourceWithCache) InitWithDataSourceDumpPath(source objectivec.IObject, path objectivec.IObject) ETDataSourceWithCache {
	rv := objc.SendIfResponds[ETDataSourceWithCache](e.ID, objc.Sel("initWithDataSource:dumpPath:"), source, path)
	return rv
}
