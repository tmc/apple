// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ETModelDefLeNet] class.
var (
	_ETModelDefLeNetClass     ETModelDefLeNetClass
	_ETModelDefLeNetClassOnce sync.Once
)

func getETModelDefLeNetClass() ETModelDefLeNetClass {
	_ETModelDefLeNetClassOnce.Do(func() {
		_ETModelDefLeNetClass = ETModelDefLeNetClass{class: objc.GetClass("ETModelDefLeNet")}
	})
	return _ETModelDefLeNetClass
}

// GetETModelDefLeNetClass returns the class object for ETModelDefLeNet.
func GetETModelDefLeNetClass() ETModelDefLeNetClass {
	return getETModelDefLeNetClass()
}

type ETModelDefLeNetClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec ETModelDefLeNetClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec ETModelDefLeNetClass) Alloc() ETModelDefLeNet {
	rv := objc.Send[ETModelDefLeNet](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ETModelDefLeNet.Output_size]
//   - [ETModelDefLeNet.SetOutput_size]
//
// See: https://developer.apple.com/documentation/Espresso/ETModelDefLeNet
type ETModelDefLeNet struct {
	ETModelDef
}

// ETModelDefLeNetFromID constructs a [ETModelDefLeNet] from an objc.ID.
func ETModelDefLeNetFromID(id objc.ID) ETModelDefLeNet {
	return ETModelDefLeNet{ETModelDef: ETModelDefFromID(id)}
}

// Ensure ETModelDefLeNet implements IETModelDefLeNet.
var _ IETModelDefLeNet = ETModelDefLeNet{}

// An interface definition for the [ETModelDefLeNet] class.
//
// # Methods
//
//   - [IETModelDefLeNet.Output_size]
//   - [IETModelDefLeNet.SetOutput_size]
//
// See: https://developer.apple.com/documentation/Espresso/ETModelDefLeNet
type IETModelDefLeNet interface {
	IETModelDef

	// Topic: Methods

	Output_size() int
	SetOutput_size(value int)
}

// Init initializes the instance.
func (e ETModelDefLeNet) Init() ETModelDefLeNet {
	rv := objc.Send[ETModelDefLeNet](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e ETModelDefLeNet) Autorelease() ETModelDefLeNet {
	rv := objc.Send[ETModelDefLeNet](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewETModelDefLeNet creates a new ETModelDefLeNet instance.
func NewETModelDefLeNet() ETModelDefLeNet {
	class := getETModelDefLeNetClass()
	rv := objc.Send[ETModelDefLeNet](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ETModelDef/initWithNetwork:
func NewETModelDefLeNetWithNetwork(network objectivec.IObject) ETModelDefLeNet {
	instance := getETModelDefLeNetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNetwork:"), network)
	return ETModelDefLeNetFromID(rv)
}

// See: https://developer.apple.com/documentation/Espresso/ETModelDefLeNet/output_size
func (e ETModelDefLeNet) Output_size() int {
	rv := objc.Send[int](e.ID, objc.Sel("output_size"))
	return rv
}
func (e ETModelDefLeNet) SetOutput_size(value int) {
	objc.Send[struct{}](e.ID, objc.Sel("setOutput_size:"), value)
}
