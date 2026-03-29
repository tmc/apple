// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [ETOptimizerDefSGD] class.
var (
	_ETOptimizerDefSGDClass     ETOptimizerDefSGDClass
	_ETOptimizerDefSGDClassOnce sync.Once
)

func getETOptimizerDefSGDClass() ETOptimizerDefSGDClass {
	_ETOptimizerDefSGDClassOnce.Do(func() {
		_ETOptimizerDefSGDClass = ETOptimizerDefSGDClass{class: objc.GetClass("ETOptimizerDefSGD")}
	})
	return _ETOptimizerDefSGDClass
}

// GetETOptimizerDefSGDClass returns the class object for ETOptimizerDefSGD.
func GetETOptimizerDefSGDClass() ETOptimizerDefSGDClass {
	return getETOptimizerDefSGDClass()
}

type ETOptimizerDefSGDClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec ETOptimizerDefSGDClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec ETOptimizerDefSGDClass) Alloc() ETOptimizerDefSGD {
	rv := objc.Send[ETOptimizerDefSGD](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [ETOptimizerDefSGD.Lr]
//   - [ETOptimizerDefSGD.SetLr]
//   - [ETOptimizerDefSGD.Lr_decay_epoch]
//   - [ETOptimizerDefSGD.SetLr_decay_epoch]
//   - [ETOptimizerDefSGD.Momentum]
//   - [ETOptimizerDefSGD.SetMomentum]
//   - [ETOptimizerDefSGD.Weight_decay]
//   - [ETOptimizerDefSGD.SetWeight_decay]
// See: https://developer.apple.com/documentation/Espresso/ETOptimizerDefSGD
type ETOptimizerDefSGD struct {
	ETOptimizerDef
}

// ETOptimizerDefSGDFromID constructs a [ETOptimizerDefSGD] from an objc.ID.
func ETOptimizerDefSGDFromID(id objc.ID) ETOptimizerDefSGD {
	return ETOptimizerDefSGD{ETOptimizerDef: ETOptimizerDefFromID(id)}
}
// Ensure ETOptimizerDefSGD implements IETOptimizerDefSGD.
var _ IETOptimizerDefSGD = ETOptimizerDefSGD{}

// An interface definition for the [ETOptimizerDefSGD] class.
//
// # Methods
//
//   - [IETOptimizerDefSGD.Lr]
//   - [IETOptimizerDefSGD.SetLr]
//   - [IETOptimizerDefSGD.Lr_decay_epoch]
//   - [IETOptimizerDefSGD.SetLr_decay_epoch]
//   - [IETOptimizerDefSGD.Momentum]
//   - [IETOptimizerDefSGD.SetMomentum]
//   - [IETOptimizerDefSGD.Weight_decay]
//   - [IETOptimizerDefSGD.SetWeight_decay]
//
// See: https://developer.apple.com/documentation/Espresso/ETOptimizerDefSGD
type IETOptimizerDefSGD interface {
	IETOptimizerDef

	// Topic: Methods

	Lr() float32
	SetLr(value float32)
	Lr_decay_epoch() float32
	SetLr_decay_epoch(value float32)
	Momentum() float32
	SetMomentum(value float32)
	Weight_decay() float32
	SetWeight_decay(value float32)
}

// Init initializes the instance.
func (e ETOptimizerDefSGD) Init() ETOptimizerDefSGD {
	rv := objc.Send[ETOptimizerDefSGD](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e ETOptimizerDefSGD) Autorelease() ETOptimizerDefSGD {
	rv := objc.Send[ETOptimizerDefSGD](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewETOptimizerDefSGD creates a new ETOptimizerDefSGD instance.
func NewETOptimizerDefSGD() ETOptimizerDefSGD {
	class := getETOptimizerDefSGDClass()
	rv := objc.Send[ETOptimizerDefSGD](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ETOptimizerDefSGD/lr
func (e ETOptimizerDefSGD) Lr() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("lr"))
	return rv
}
func (e ETOptimizerDefSGD) SetLr(value float32) {
	objc.Send[struct{}](e.ID, objc.Sel("setLr:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/ETOptimizerDefSGD/lr_decay_epoch
func (e ETOptimizerDefSGD) Lr_decay_epoch() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("lr_decay_epoch"))
	return rv
}
func (e ETOptimizerDefSGD) SetLr_decay_epoch(value float32) {
	objc.Send[struct{}](e.ID, objc.Sel("setLr_decay_epoch:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/ETOptimizerDefSGD/momentum
func (e ETOptimizerDefSGD) Momentum() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("momentum"))
	return rv
}
func (e ETOptimizerDefSGD) SetMomentum(value float32) {
	objc.Send[struct{}](e.ID, objc.Sel("setMomentum:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/ETOptimizerDefSGD/weight_decay
func (e ETOptimizerDefSGD) Weight_decay() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("weight_decay"))
	return rv
}
func (e ETOptimizerDefSGD) SetWeight_decay(value float32) {
	objc.Send[struct{}](e.ID, objc.Sel("setWeight_decay:"), value)
}

