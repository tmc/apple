// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ETDataSourceWithExtractor] class.
var (
	_ETDataSourceWithExtractorClass     ETDataSourceWithExtractorClass
	_ETDataSourceWithExtractorClassOnce sync.Once
)

func getETDataSourceWithExtractorClass() ETDataSourceWithExtractorClass {
	_ETDataSourceWithExtractorClassOnce.Do(func() {
		_ETDataSourceWithExtractorClass = ETDataSourceWithExtractorClass{class: objc.GetClass("ETDataSourceWithExtractor")}
	})
	return _ETDataSourceWithExtractorClass
}

// GetETDataSourceWithExtractorClass returns the class object for ETDataSourceWithExtractor.
func GetETDataSourceWithExtractorClass() ETDataSourceWithExtractorClass {
	return getETDataSourceWithExtractorClass()
}

type ETDataSourceWithExtractorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec ETDataSourceWithExtractorClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec ETDataSourceWithExtractorClass) Alloc() ETDataSourceWithExtractor {
	rv := objc.Send[ETDataSourceWithExtractor](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [ETDataSourceWithExtractor.DataPointAtIndex]
//   - [ETDataSourceWithExtractor.NumberOfDataPoints]
//   - [ETDataSourceWithExtractor.InitWithDataSourceExtractor]
// See: https://developer.apple.com/documentation/Espresso/ETDataSourceWithExtractor
type ETDataSourceWithExtractor struct {
	objectivec.Object
}

// ETDataSourceWithExtractorFromID constructs a [ETDataSourceWithExtractor] from an objc.ID.
func ETDataSourceWithExtractorFromID(id objc.ID) ETDataSourceWithExtractor {
	return ETDataSourceWithExtractor{objectivec.Object{ID: id}}
}
// Ensure ETDataSourceWithExtractor implements IETDataSourceWithExtractor.
var _ IETDataSourceWithExtractor = ETDataSourceWithExtractor{}

// An interface definition for the [ETDataSourceWithExtractor] class.
//
// # Methods
//
//   - [IETDataSourceWithExtractor.DataPointAtIndex]
//   - [IETDataSourceWithExtractor.NumberOfDataPoints]
//   - [IETDataSourceWithExtractor.InitWithDataSourceExtractor]
//
// See: https://developer.apple.com/documentation/Espresso/ETDataSourceWithExtractor
type IETDataSourceWithExtractor interface {
	objectivec.IObject

	// Topic: Methods

	DataPointAtIndex(index int) objectivec.IObject
	NumberOfDataPoints() int
	InitWithDataSourceExtractor(source objectivec.IObject, extractor objectivec.IObject) ETDataSourceWithExtractor
}

// Init initializes the instance.
func (e ETDataSourceWithExtractor) Init() ETDataSourceWithExtractor {
	rv := objc.Send[ETDataSourceWithExtractor](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e ETDataSourceWithExtractor) Autorelease() ETDataSourceWithExtractor {
	rv := objc.Send[ETDataSourceWithExtractor](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewETDataSourceWithExtractor creates a new ETDataSourceWithExtractor instance.
func NewETDataSourceWithExtractor() ETDataSourceWithExtractor {
	class := getETDataSourceWithExtractorClass()
	rv := objc.Send[ETDataSourceWithExtractor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/ETDataSourceWithExtractor/initWithDataSource:extractor:
func NewETDataSourceWithExtractorWithDataSourceExtractor(source objectivec.IObject, extractor objectivec.IObject) ETDataSourceWithExtractor {
	instance := getETDataSourceWithExtractorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDataSource:extractor:"), source, extractor)
	return ETDataSourceWithExtractorFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Espresso/ETDataSourceWithExtractor/dataPointAtIndex:
func (e ETDataSourceWithExtractor) DataPointAtIndex(index int) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("dataPointAtIndex:"), index)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/Espresso/ETDataSourceWithExtractor/numberOfDataPoints
func (e ETDataSourceWithExtractor) NumberOfDataPoints() int {
	rv := objc.Send[int](e.ID, objc.Sel("numberOfDataPoints"))
	return rv
}
//
// See: https://developer.apple.com/documentation/Espresso/ETDataSourceWithExtractor/initWithDataSource:extractor:
func (e ETDataSourceWithExtractor) InitWithDataSourceExtractor(source objectivec.IObject, extractor objectivec.IObject) ETDataSourceWithExtractor {
	rv := objc.Send[ETDataSourceWithExtractor](e.ID, objc.Sel("initWithDataSource:extractor:"), source, extractor)
	return rv
}

