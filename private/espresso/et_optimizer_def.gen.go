// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ETOptimizerDef] class.
var (
	_ETOptimizerDefClass     ETOptimizerDefClass
	_ETOptimizerDefClassOnce sync.Once
)

func getETOptimizerDefClass() ETOptimizerDefClass {
	_ETOptimizerDefClassOnce.Do(func() {
		_ETOptimizerDefClass = ETOptimizerDefClass{class: objc.GetClass("ETOptimizerDef")}
	})
	return _ETOptimizerDefClass
}

// GetETOptimizerDefClass returns the class object for ETOptimizerDef.
func GetETOptimizerDefClass() ETOptimizerDefClass {
	return getETOptimizerDefClass()
}

type ETOptimizerDefClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec ETOptimizerDefClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec ETOptimizerDefClass) Alloc() ETOptimizerDef {
	rv := objc.Send[ETOptimizerDef](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [ETOptimizerDef.Batch_size]
//   - [ETOptimizerDef.SetBatch_size]
// See: https://developer.apple.com/documentation/Espresso/ETOptimizerDef
type ETOptimizerDef struct {
	objectivec.Object
}

// ETOptimizerDefFromID constructs a [ETOptimizerDef] from an objc.ID.
func ETOptimizerDefFromID(id objc.ID) ETOptimizerDef {
	return ETOptimizerDef{objectivec.Object{ID: id}}
}
// Ensure ETOptimizerDef implements IETOptimizerDef.
var _ IETOptimizerDef = ETOptimizerDef{}

// An interface definition for the [ETOptimizerDef] class.
//
// # Methods
//
//   - [IETOptimizerDef.Batch_size]
//   - [IETOptimizerDef.SetBatch_size]
//
// See: https://developer.apple.com/documentation/Espresso/ETOptimizerDef
type IETOptimizerDef interface {
	objectivec.IObject

	// Topic: Methods

	Batch_size() uint32
	SetBatch_size(value uint32)
}

// Init initializes the instance.
func (e ETOptimizerDef) Init() ETOptimizerDef {
	rv := objc.Send[ETOptimizerDef](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e ETOptimizerDef) Autorelease() ETOptimizerDef {
	rv := objc.Send[ETOptimizerDef](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewETOptimizerDef creates a new ETOptimizerDef instance.
func NewETOptimizerDef() ETOptimizerDef {
	class := getETOptimizerDefClass()
	rv := objc.Send[ETOptimizerDef](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ETOptimizerDef/batch_size
func (e ETOptimizerDef) Batch_size() uint32 {
	rv := objc.Send[uint32](e.ID, objc.Sel("batch_size"))
	return rv
}
func (e ETOptimizerDef) SetBatch_size(value uint32) {
	objc.Send[struct{}](e.ID, objc.Sel("setBatch_size:"), value)
}

