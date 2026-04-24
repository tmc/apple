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
	rv := objc.Send[CGXConnectionBox](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [CGXConnectionBox.Connection]
//   - [CGXConnectionBox.SetConnection]
//   - [CGXConnectionBox.InvalidateBackreference]
//   - [CGXConnectionBox.InitWithCGXConnection]
//
// See: https://developer.apple.com/documentation/SkyLight/CGXConnectionBox
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
//
// See: https://developer.apple.com/documentation/SkyLight/CGXConnectionBox
type ICGXConnectionBox interface {
	objectivec.IObject

	// Topic: Methods

	Connection() unsafe.Pointer
	SetConnection(value unsafe.Pointer)
	InvalidateBackreference()
	InitWithCGXConnection(cGXConnection unsafe.Pointer) CGXConnectionBox
}

// Init initializes the instance.
func (x CGXConnectionBox) Init() CGXConnectionBox {
	rv := objc.Send[CGXConnectionBox](x.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (x CGXConnectionBox) Autorelease() CGXConnectionBox {
	rv := objc.Send[CGXConnectionBox](x.ID, objc.Sel("autorelease"))
	return rv
}

// NewCGXConnectionBox creates a new CGXConnectionBox instance.
func NewCGXConnectionBox() CGXConnectionBox {
	class := getCGXConnectionBoxClass()
	rv := objc.Send[CGXConnectionBox](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CGXConnectionBox/initWithCGXConnection:
func NewXConnectionBoxWithCGXConnection(cGXConnection unsafe.Pointer) CGXConnectionBox {
	instance := getCGXConnectionBoxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCGXConnection:"), cGXConnection)
	return CGXConnectionBoxFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/CGXConnectionBox/invalidateBackreference
func (x CGXConnectionBox) InvalidateBackreference() {
	objc.Send[objc.ID](x.ID, objc.Sel("invalidateBackreference"))
}

// See: https://developer.apple.com/documentation/SkyLight/CGXConnectionBox/initWithCGXConnection:
func (x CGXConnectionBox) InitWithCGXConnection(cGXConnection unsafe.Pointer) CGXConnectionBox {
	rv := objc.Send[CGXConnectionBox](x.ID, objc.Sel("initWithCGXConnection:"), cGXConnection)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CGXConnectionBox/connection
func (x CGXConnectionBox) Connection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x.ID, objc.Sel("connection"))
	return rv
}
func (x CGXConnectionBox) SetConnection(value unsafe.Pointer) {
	objc.Send[struct{}](x.ID, objc.Sel("setConnection:"), value)
}
