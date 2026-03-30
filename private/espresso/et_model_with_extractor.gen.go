// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ETModelWithExtractor] class.
var (
	_ETModelWithExtractorClass     ETModelWithExtractorClass
	_ETModelWithExtractorClassOnce sync.Once
)

func getETModelWithExtractorClass() ETModelWithExtractorClass {
	_ETModelWithExtractorClassOnce.Do(func() {
		_ETModelWithExtractorClass = ETModelWithExtractorClass{class: objc.GetClass("ETModelWithExtractor")}
	})
	return _ETModelWithExtractorClass
}

// GetETModelWithExtractorClass returns the class object for ETModelWithExtractor.
func GetETModelWithExtractorClass() ETModelWithExtractorClass {
	return getETModelWithExtractorClass()
}

type ETModelWithExtractorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec ETModelWithExtractorClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec ETModelWithExtractorClass) Alloc() ETModelWithExtractor {
	rv := objc.Send[ETModelWithExtractor](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ETModelWithExtractor
type ETModelWithExtractor struct {
	ETModelDef
}

// ETModelWithExtractorFromID constructs a [ETModelWithExtractor] from an objc.ID.
func ETModelWithExtractorFromID(id objc.ID) ETModelWithExtractor {
	return ETModelWithExtractor{ETModelDef: ETModelDefFromID(id)}
}

// Ensure ETModelWithExtractor implements IETModelWithExtractor.
var _ IETModelWithExtractor = ETModelWithExtractor{}

// An interface definition for the [ETModelWithExtractor] class.
//
// See: https://developer.apple.com/documentation/Espresso/ETModelWithExtractor
type IETModelWithExtractor interface {
	IETModelDef
}

// Init initializes the instance.
func (e ETModelWithExtractor) Init() ETModelWithExtractor {
	rv := objc.Send[ETModelWithExtractor](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e ETModelWithExtractor) Autorelease() ETModelWithExtractor {
	rv := objc.Send[ETModelWithExtractor](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewETModelWithExtractor creates a new ETModelWithExtractor instance.
func NewETModelWithExtractor() ETModelWithExtractor {
	class := getETModelWithExtractorClass()
	rv := objc.Send[ETModelWithExtractor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ETModelDef/initWithNetwork:
func NewETModelWithExtractorWithNetwork(network objectivec.IObject) ETModelWithExtractor {
	instance := getETModelWithExtractorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNetwork:"), network)
	return ETModelWithExtractorFromID(rv)
}
