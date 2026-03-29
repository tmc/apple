// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ColumnCount] class.
var (
	_ColumnCountClass     ColumnCountClass
	_ColumnCountClassOnce sync.Once
)

func getColumnCountClass() ColumnCountClass {
	_ColumnCountClassOnce.Do(func() {
		_ColumnCountClass = ColumnCountClass{class: objc.GetClass("columnCount")}
	})
	return _ColumnCountClass
}

// GetColumnCountClass returns the class object for columnCount.
func GetColumnCountClass() ColumnCountClass {
	return getColumnCountClass()
}

type ColumnCountClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc ColumnCountClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc ColumnCountClass) Alloc() ColumnCount {
	rv := objc.Send[ColumnCount](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/columnCount-c.ivar
type ColumnCount struct {
	objectivec.Object
}

// ColumnCountFromID constructs a [ColumnCount] from an objc.ID.
func ColumnCountFromID(id objc.ID) ColumnCount {
	return ColumnCount{objectivec.Object{ID: id}}
}
// Ensure ColumnCount implements IColumnCount.
var _ IColumnCount = ColumnCount{}

// An interface definition for the [ColumnCount] class.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/columnCount-c.ivar
type IColumnCount interface {
	objectivec.IObject
}

// Init initializes the instance.
func (c ColumnCount) Init() ColumnCount {
	rv := objc.Send[ColumnCount](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c ColumnCount) Autorelease() ColumnCount {
	rv := objc.Send[ColumnCount](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewColumnCount creates a new ColumnCount instance.
func NewColumnCount() ColumnCount {
	class := getColumnCountClass()
	rv := objc.Send[ColumnCount](objc.ID(class.class), objc.Sel("new"))
	return rv
}

