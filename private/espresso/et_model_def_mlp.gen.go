// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ETModelDefMLP] class.
var (
	_ETModelDefMLPClass     ETModelDefMLPClass
	_ETModelDefMLPClassOnce sync.Once
)

func getETModelDefMLPClass() ETModelDefMLPClass {
	_ETModelDefMLPClassOnce.Do(func() {
		_ETModelDefMLPClass = ETModelDefMLPClass{class: objc.GetClass("ETModelDefMLP")}
	})
	return _ETModelDefMLPClass
}

// GetETModelDefMLPClass returns the class object for ETModelDefMLP.
func GetETModelDefMLPClass() ETModelDefMLPClass {
	return getETModelDefMLPClass()
}

type ETModelDefMLPClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec ETModelDefMLPClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec ETModelDefMLPClass) Alloc() ETModelDefMLP {
	rv := objc.Send[ETModelDefMLP](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [ETModelDefMLP.BuildNetwork]
//   - [ETModelDefMLP.Hidden_size]
//   - [ETModelDefMLP.SetHidden_size]
//   - [ETModelDefMLP.Input_size]
//   - [ETModelDefMLP.SetInput_size]
//   - [ETModelDefMLP.Output_size]
//   - [ETModelDefMLP.SetOutput_size]
// See: https://developer.apple.com/documentation/Espresso/ETModelDefMLP
type ETModelDefMLP struct {
	ETModelDef
}

// ETModelDefMLPFromID constructs a [ETModelDefMLP] from an objc.ID.
func ETModelDefMLPFromID(id objc.ID) ETModelDefMLP {
	return ETModelDefMLP{ETModelDef: ETModelDefFromID(id)}
}
// Ensure ETModelDefMLP implements IETModelDefMLP.
var _ IETModelDefMLP = ETModelDefMLP{}

// An interface definition for the [ETModelDefMLP] class.
//
// # Methods
//
//   - [IETModelDefMLP.BuildNetwork]
//   - [IETModelDefMLP.Hidden_size]
//   - [IETModelDefMLP.SetHidden_size]
//   - [IETModelDefMLP.Input_size]
//   - [IETModelDefMLP.SetInput_size]
//   - [IETModelDefMLP.Output_size]
//   - [IETModelDefMLP.SetOutput_size]
//
// See: https://developer.apple.com/documentation/Espresso/ETModelDefMLP
type IETModelDefMLP interface {
	IETModelDef

	// Topic: Methods

	BuildNetwork()
	Hidden_size() int
	SetHidden_size(value int)
	Input_size() int
	SetInput_size(value int)
	Output_size() int
	SetOutput_size(value int)
}

// Init initializes the instance.
func (e ETModelDefMLP) Init() ETModelDefMLP {
	rv := objc.Send[ETModelDefMLP](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e ETModelDefMLP) Autorelease() ETModelDefMLP {
	rv := objc.Send[ETModelDefMLP](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewETModelDefMLP creates a new ETModelDefMLP instance.
func NewETModelDefMLP() ETModelDefMLP {
	class := getETModelDefMLPClass()
	rv := objc.Send[ETModelDefMLP](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/ETModelDef/initWithNetwork:
func NewETModelDefMLPWithNetwork(network objectivec.IObject) ETModelDefMLP {
	instance := getETModelDefMLPClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNetwork:"), network)
	return ETModelDefMLPFromID(rv)
}

// See: https://developer.apple.com/documentation/Espresso/ETModelDefMLP/buildNetwork
func (e ETModelDefMLP) BuildNetwork() {
	objc.Send[objc.ID](e.ID, objc.Sel("buildNetwork"))
}

// See: https://developer.apple.com/documentation/Espresso/ETModelDefMLP/hidden_size
func (e ETModelDefMLP) Hidden_size() int {
	rv := objc.Send[int](e.ID, objc.Sel("hidden_size"))
	return rv
}
func (e ETModelDefMLP) SetHidden_size(value int) {
	objc.Send[struct{}](e.ID, objc.Sel("setHidden_size:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/ETModelDefMLP/input_size
func (e ETModelDefMLP) Input_size() int {
	rv := objc.Send[int](e.ID, objc.Sel("input_size"))
	return rv
}
func (e ETModelDefMLP) SetInput_size(value int) {
	objc.Send[struct{}](e.ID, objc.Sel("setInput_size:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/ETModelDefMLP/output_size
func (e ETModelDefMLP) Output_size() int {
	rv := objc.Send[int](e.ID, objc.Sel("output_size"))
	return rv
}
func (e ETModelDefMLP) SetOutput_size(value int) {
	objc.Send[struct{}](e.ID, objc.Sel("setOutput_size:"), value)
}

