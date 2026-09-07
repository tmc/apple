// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DataFramer] class.
var (
	_DataFramerClass     DataFramerClass
	_DataFramerClassOnce sync.Once
)

func getDataFramerClass() DataFramerClass {
	_DataFramerClassOnce.Do(func() {
		_DataFramerClass = DataFramerClass{class: objc.GetClass("_TtC12TextToSpeechP33_69066F15C84B7E6FACC0750A66976E5010DataFramer")}
	})
	return _DataFramerClass
}

// GetDataFramerClass returns the class object for _TtC12TextToSpeechP33_69066F15C84B7E6FACC0750A66976E5010DataFramer.
func GetDataFramerClass() DataFramerClass {
	return getDataFramerClass()
}

type DataFramerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DataFramerClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DataFramerClass) Alloc() DataFramer {
	rv := objc.SendIfResponds[DataFramer](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

type DataFramer struct {
	objectivec.Object
}

// DataFramerFromID constructs a [DataFramer] from an objc.ID.
func DataFramerFromID(id objc.ID) DataFramer {
	return DataFramer{objectivec.Object{ID: id}}
}

// Ensure DataFramer implements IDataFramer.
var _ IDataFramer = DataFramer{}

// An interface definition for the [DataFramer] class.
type IDataFramer interface {
	objectivec.IObject
}

// Init initializes the instance.
func (d DataFramer) Init() DataFramer {
	rv := objc.SendIfResponds[DataFramer](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DataFramer) Autorelease() DataFramer {
	rv := objc.SendIfResponds[DataFramer](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDataFramer creates a new DataFramer instance.
func NewDataFramer() DataFramer {
	class := getDataFramerClass()
	rv := objc.SendIfResponds[DataFramer](objc.ID(class.class), objc.Sel("new"))
	return rv
}
