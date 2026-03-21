// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DataCodewordCount] class.
var (
	_DataCodewordCountClass     DataCodewordCountClass
	_DataCodewordCountClassOnce sync.Once
)

func getDataCodewordCountClass() DataCodewordCountClass {
	_DataCodewordCountClassOnce.Do(func() {
		_DataCodewordCountClass = DataCodewordCountClass{class: objc.GetClass("dataCodewordCount")}
	})
	return _DataCodewordCountClass
}

// GetDataCodewordCountClass returns the class object for dataCodewordCount.
func GetDataCodewordCountClass() DataCodewordCountClass {
	return getDataCodewordCountClass()
}

type DataCodewordCountClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (dc DataCodewordCountClass) Alloc() DataCodewordCount {
	rv := objc.Send[DataCodewordCount](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIAztecCodeDescriptor/dataCodewordCount-c.ivar
type DataCodewordCount struct {
	objectivec.Object
}

// DataCodewordCountFromID constructs a [DataCodewordCount] from an objc.ID.
func DataCodewordCountFromID(id objc.ID) DataCodewordCount {
	return DataCodewordCount{objectivec.Object{ID: id}}
}
// Ensure DataCodewordCount implements IDataCodewordCount.
var _ IDataCodewordCount = DataCodewordCount{}

// An interface definition for the [DataCodewordCount] class.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAztecCodeDescriptor/dataCodewordCount-c.ivar
type IDataCodewordCount interface {
	objectivec.IObject
}

// Init initializes the instance.
func (d DataCodewordCount) Init() DataCodewordCount {
	rv := objc.Send[DataCodewordCount](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DataCodewordCount) Autorelease() DataCodewordCount {
	rv := objc.Send[DataCodewordCount](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDataCodewordCount creates a new DataCodewordCount instance.
func NewDataCodewordCount() DataCodewordCount {
	class := getDataCodewordCountClass()
	rv := objc.Send[DataCodewordCount](objc.ID(class.class), objc.Sel("new"))
	return rv
}

