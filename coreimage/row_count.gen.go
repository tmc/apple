// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [RowCount] class.
var (
	_RowCountClass     RowCountClass
	_RowCountClassOnce sync.Once
)

func getRowCountClass() RowCountClass {
	_RowCountClassOnce.Do(func() {
		_RowCountClass = RowCountClass{class: objc.GetClass("rowCount")}
	})
	return _RowCountClass
}

// GetRowCountClass returns the class object for rowCount.
func GetRowCountClass() RowCountClass {
	return getRowCountClass()
}

type RowCountClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (rc RowCountClass) Class() objc.Class {
	return rc.class
}

// Alloc allocates memory for a new instance of the class.
func (rc RowCountClass) Alloc() RowCount {
	rv := objc.Send[RowCount](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/rowCount-c.ivar
type RowCount struct {
	objectivec.Object
}

// RowCountFromID constructs a [RowCount] from an objc.ID.
func RowCountFromID(id objc.ID) RowCount {
	return RowCount{objectivec.Object{ID: id}}
}

// Ensure RowCount implements IRowCount.
var _ IRowCount = RowCount{}

// An interface definition for the [RowCount] class.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/rowCount-c.ivar
type IRowCount interface {
	objectivec.IObject
}

// Init initializes the instance.
func (r RowCount) Init() RowCount {
	rv := objc.Send[RowCount](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r RowCount) Autorelease() RowCount {
	rv := objc.Send[RowCount](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewRowCount creates a new RowCount instance.
func NewRowCount() RowCount {
	class := getRowCountClass()
	rv := objc.Send[RowCount](objc.ID(class.class), objc.Sel("new"))
	return rv
}
