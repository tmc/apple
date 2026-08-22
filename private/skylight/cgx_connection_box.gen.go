// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CGXConnectionBox] class.
var (
	_CGXConnectionBoxClass     CGXConnectionBoxClass
	_CGXConnectionBoxClassOnce sync.Once
)

func getCGXConnectionBoxClass() CGXConnectionBoxClass {
	_CGXConnectionBoxClassOnce.Do(func() {
		_CGXConnectionBoxClass = CGXConnectionBoxClass{class: objc.GetClass("CGXConnectionBox")}
	})
	return _CGXConnectionBoxClass
}

// GetCGXConnectionBoxClass returns the class object for CGXConnectionBox.
func GetCGXConnectionBoxClass() CGXConnectionBoxClass {
	return getCGXConnectionBoxClass()
}

type CGXConnectionBoxClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CGXConnectionBoxClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CGXConnectionBoxClass) Alloc() CGXConnectionBox {
	rv := objc.SendIfResponds[CGXConnectionBox](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [CGXConnectionBox.Connection]
//   - [CGXConnectionBox.SetConnection]
//   - [CGXConnectionBox.InvalidateBackreference]
//   - [CGXConnectionBox.InitWithCGXConnection]
type CGXConnectionBox struct {
	objectivec.Object
}

// CGXConnectionBoxFromID constructs a [CGXConnectionBox] from an objc.ID.
func CGXConnectionBoxFromID(id objc.ID) CGXConnectionBox {
	return CGXConnectionBox{objectivec.Object{ID: id}}
}

// Ensure CGXConnectionBox implements ICGXConnectionBox.
var _ ICGXConnectionBox = CGXConnectionBox{}

// An interface definition for the [CGXConnectionBox] class.
//
// # Methods
//
//   - [ICGXConnectionBox.Connection]
//   - [ICGXConnectionBox.SetConnection]
//   - [ICGXConnectionBox.InvalidateBackreference]
//   - [ICGXConnectionBox.InitWithCGXConnection]
type ICGXConnectionBox interface {
	objectivec.IObject

	// Topic: Methods

	Connection() *CGXConnection
	SetConnection(value *CGXConnection)
	InvalidateBackreference()
	InitWithCGXConnection(cGXConnection *CGXConnection) CGXConnectionBox
}

// Init initializes the instance.
func (c CGXConnectionBox) Init() CGXConnectionBox {
	rv := objc.SendIfResponds[CGXConnectionBox](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CGXConnectionBox) Autorelease() CGXConnectionBox {
	rv := objc.SendIfResponds[CGXConnectionBox](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCGXConnectionBox creates a new CGXConnectionBox instance.
func NewCGXConnectionBox() CGXConnectionBox {
	class := getCGXConnectionBoxClass()
	rv := objc.SendIfResponds[CGXConnectionBox](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewXConnectionBoxWithCGXConnection(cGXConnection *CGXConnection) CGXConnectionBox {
	instance := getCGXConnectionBoxClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCGXConnection:"), unsafe.Pointer(cGXConnection))
	return CGXConnectionBoxFromID(rv)
}

func (c CGXConnectionBox) InvalidateBackreference() {
	objc.SendIfResponds[objc.ID](c.ID, objc.Sel("invalidateBackreference"))
}
func (c CGXConnectionBox) InitWithCGXConnection(cGXConnection *CGXConnection) CGXConnectionBox {
	rv := objc.SendIfResponds[CGXConnectionBox](c.ID, objc.Sel("initWithCGXConnection:"), unsafe.Pointer(cGXConnection))
	return rv
}

func (c CGXConnectionBox) Connection() *CGXConnection {
	rv := objc.SendIfResponds[unsafe.Pointer](c.ID, objc.Sel("connection"))
	return (*CGXConnection)(rv)
}
func (c CGXConnectionBox) SetConnection(value *CGXConnection) {
	objc.SendIfResponds[struct{}](c.ID, objc.Sel("setConnection:"), value)
}
