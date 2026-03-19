// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [EspressoDataFrameStorage] class.
var (
	_EspressoDataFrameStorageClass     EspressoDataFrameStorageClass
	_EspressoDataFrameStorageClassOnce sync.Once
)

func getEspressoDataFrameStorageClass() EspressoDataFrameStorageClass {
	_EspressoDataFrameStorageClassOnce.Do(func() {
		_EspressoDataFrameStorageClass = EspressoDataFrameStorageClass{class: objc.GetClass("EspressoDataFrameStorage")}
	})
	return _EspressoDataFrameStorageClass
}

// GetEspressoDataFrameStorageClass returns the class object for EspressoDataFrameStorage.
func GetEspressoDataFrameStorageClass() EspressoDataFrameStorageClass {
	return getEspressoDataFrameStorageClass()
}

type EspressoDataFrameStorageClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoDataFrameStorageClass) Alloc() EspressoDataFrameStorage {
	rv := objc.Send[EspressoDataFrameStorage](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [EspressoDataFrameStorage.BaseFilename]
//   - [EspressoDataFrameStorage.SetBaseFilename]
//   - [EspressoDataFrameStorage.DataFrameAtIndex]
//   - [EspressoDataFrameStorage.DataFrames]
//   - [EspressoDataFrameStorage.SetDataFrames]
//   - [EspressoDataFrameStorage.MappedFiles]
//   - [EspressoDataFrameStorage.SetMappedFiles]
//   - [EspressoDataFrameStorage.NumberOfDataFrames]
// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameStorage
type EspressoDataFrameStorage struct {
	objectivec.Object
}

// EspressoDataFrameStorageFromID constructs a [EspressoDataFrameStorage] from an objc.ID.
func EspressoDataFrameStorageFromID(id objc.ID) EspressoDataFrameStorage {
	return EspressoDataFrameStorage{objectivec.Object{ID: id}}
}
// Ensure EspressoDataFrameStorage implements IEspressoDataFrameStorage.
var _ IEspressoDataFrameStorage = EspressoDataFrameStorage{}

// An interface definition for the [EspressoDataFrameStorage] class.
//
// # Methods
//
//   - [IEspressoDataFrameStorage.BaseFilename]
//   - [IEspressoDataFrameStorage.SetBaseFilename]
//   - [IEspressoDataFrameStorage.DataFrameAtIndex]
//   - [IEspressoDataFrameStorage.DataFrames]
//   - [IEspressoDataFrameStorage.SetDataFrames]
//   - [IEspressoDataFrameStorage.MappedFiles]
//   - [IEspressoDataFrameStorage.SetMappedFiles]
//   - [IEspressoDataFrameStorage.NumberOfDataFrames]
//
// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameStorage
type IEspressoDataFrameStorage interface {
	objectivec.IObject

	// Topic: Methods

	BaseFilename() string
	SetBaseFilename(value string)
	DataFrameAtIndex(index uint64) objectivec.IObject
	DataFrames() foundation.INSArray
	SetDataFrames(value foundation.INSArray)
	MappedFiles() foundation.INSDictionary
	SetMappedFiles(value foundation.INSDictionary)
	NumberOfDataFrames() uint64
}

// Init initializes the instance.
func (e EspressoDataFrameStorage) Init() EspressoDataFrameStorage {
	rv := objc.Send[EspressoDataFrameStorage](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoDataFrameStorage) Autorelease() EspressoDataFrameStorage {
	rv := objc.Send[EspressoDataFrameStorage](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoDataFrameStorage creates a new EspressoDataFrameStorage instance.
func NewEspressoDataFrameStorage() EspressoDataFrameStorage {
	class := getEspressoDataFrameStorageClass()
	rv := objc.Send[EspressoDataFrameStorage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameStorage/dataFrameAtIndex:
func (e EspressoDataFrameStorage) DataFrameAtIndex(index uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("dataFrameAtIndex:"), index)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameStorage/numberOfDataFrames
func (e EspressoDataFrameStorage) NumberOfDataFrames() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("numberOfDataFrames"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameStorage/dataFrameStorageFromPath:error:
func (_EspressoDataFrameStorageClass EspressoDataFrameStorageClass) DataFrameStorageFromPathError(path objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_EspressoDataFrameStorageClass.class), objc.Sel("dataFrameStorageFromPath:error:"), path, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameStorage/baseFilename
func (e EspressoDataFrameStorage) BaseFilename() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("baseFilename"))
	return foundation.NSStringFromID(rv).String()
}
func (e EspressoDataFrameStorage) SetBaseFilename(value string) {
	objc.Send[struct{}](e.ID, objc.Sel("setBaseFilename:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameStorage/dataFrames
func (e EspressoDataFrameStorage) DataFrames() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("dataFrames"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (e EspressoDataFrameStorage) SetDataFrames(value foundation.INSArray) {
	objc.Send[struct{}](e.ID, objc.Sel("setDataFrames:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameStorage/mappedFiles
func (e EspressoDataFrameStorage) MappedFiles() foundation.INSDictionary {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("mappedFiles"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (e EspressoDataFrameStorage) SetMappedFiles(value foundation.INSDictionary) {
	objc.Send[struct{}](e.ID, objc.Sel("setMappedFiles:"), value)
}

