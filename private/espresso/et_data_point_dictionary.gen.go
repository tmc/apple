// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ETDataPointDictionary] class.
var (
	_ETDataPointDictionaryClass     ETDataPointDictionaryClass
	_ETDataPointDictionaryClassOnce sync.Once
)

func getETDataPointDictionaryClass() ETDataPointDictionaryClass {
	_ETDataPointDictionaryClassOnce.Do(func() {
		_ETDataPointDictionaryClass = ETDataPointDictionaryClass{class: objc.GetClass("ETDataPointDictionary")}
	})
	return _ETDataPointDictionaryClass
}

// GetETDataPointDictionaryClass returns the class object for ETDataPointDictionary.
func GetETDataPointDictionaryClass() ETDataPointDictionaryClass {
	return getETDataPointDictionaryClass()
}

type ETDataPointDictionaryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec ETDataPointDictionaryClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec ETDataPointDictionaryClass) Alloc() ETDataPointDictionary {
	rv := objc.Send[ETDataPointDictionary](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [ETDataPointDictionary.DataArrayForKeyError]
//   - [ETDataPointDictionary.DataForKeyError]
//   - [ETDataPointDictionary.Float_buffers]
//   - [ETDataPointDictionary.SetFloat_buffers]
//   - [ETDataPointDictionary.Image_buffers]
//   - [ETDataPointDictionary.SetImage_buffers]
//   - [ETDataPointDictionary.SetDataSizeForKeyFreeWhenDone]
//   - [ETDataPointDictionary.SetImageForKey]
// See: https://developer.apple.com/documentation/Espresso/ETDataPointDictionary
type ETDataPointDictionary struct {
	objectivec.Object
}

// ETDataPointDictionaryFromID constructs a [ETDataPointDictionary] from an objc.ID.
func ETDataPointDictionaryFromID(id objc.ID) ETDataPointDictionary {
	return ETDataPointDictionary{objectivec.Object{ID: id}}
}
// Ensure ETDataPointDictionary implements IETDataPointDictionary.
var _ IETDataPointDictionary = ETDataPointDictionary{}

// An interface definition for the [ETDataPointDictionary] class.
//
// # Methods
//
//   - [IETDataPointDictionary.DataArrayForKeyError]
//   - [IETDataPointDictionary.DataForKeyError]
//   - [IETDataPointDictionary.Float_buffers]
//   - [IETDataPointDictionary.SetFloat_buffers]
//   - [IETDataPointDictionary.Image_buffers]
//   - [IETDataPointDictionary.SetImage_buffers]
//   - [IETDataPointDictionary.SetDataSizeForKeyFreeWhenDone]
//   - [IETDataPointDictionary.SetImageForKey]
//
// See: https://developer.apple.com/documentation/Espresso/ETDataPointDictionary
type IETDataPointDictionary interface {
	objectivec.IObject

	// Topic: Methods

	DataArrayForKeyError(key objectivec.IObject) (objectivec.IObject, error)
	DataForKeyError(key objectivec.IObject) (unsafe.Pointer, error)
	Float_buffers() objectivec.IObject
	SetFloat_buffers(value objectivec.IObject)
	Image_buffers() objectivec.IObject
	SetImage_buffers(value objectivec.IObject)
	SetDataSizeForKeyFreeWhenDone(size uint64, key objectivec.IObject, done bool) (float32, bool)
	SetImageForKey(image unsafe.Pointer, key objectivec.IObject) bool
}

// Init initializes the instance.
func (e ETDataPointDictionary) Init() ETDataPointDictionary {
	rv := objc.Send[ETDataPointDictionary](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e ETDataPointDictionary) Autorelease() ETDataPointDictionary {
	rv := objc.Send[ETDataPointDictionary](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewETDataPointDictionary creates a new ETDataPointDictionary instance.
func NewETDataPointDictionary() ETDataPointDictionary {
	class := getETDataPointDictionaryClass()
	rv := objc.Send[ETDataPointDictionary](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/ETDataPointDictionary/dataArrayForKey:error:
func (e ETDataPointDictionary) DataArrayForKeyError(key objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("dataArrayForKey:error:"), key, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/Espresso/ETDataPointDictionary/dataForKey:error:
func (e ETDataPointDictionary) DataForKeyError(key objectivec.IObject) (unsafe.Pointer, error) {
	var errorPtr objc.ID
	rv := objc.Send[unsafe.Pointer](e.ID, objc.Sel("dataForKey:error:"), key, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/Espresso/ETDataPointDictionary/setData:size:forKey:freeWhenDone:
func (e ETDataPointDictionary) SetDataSizeForKeyFreeWhenDone(size uint64, key objectivec.IObject, done bool) (float32, bool) {
	var data float32
	rv := objc.Send[bool](e.ID, objc.Sel("setData:size:forKey:freeWhenDone:"), unsafe.Pointer(&data), size, key, done)
	return data, rv
}
//
// See: https://developer.apple.com/documentation/Espresso/ETDataPointDictionary/setImage:forKey:
func (e ETDataPointDictionary) SetImageForKey(image unsafe.Pointer, key objectivec.IObject) bool {
	rv := objc.Send[bool](e.ID, objc.Sel("setImage:forKey:"), image, key)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ETDataPointDictionary/float_buffers
func (e ETDataPointDictionary) Float_buffers() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("float_buffers"))
	return objectivec.Object{ID: rv}
}
func (e ETDataPointDictionary) SetFloat_buffers(value objectivec.IObject) {
	objc.Send[struct{}](e.ID, objc.Sel("setFloat_buffers:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/ETDataPointDictionary/image_buffers
func (e ETDataPointDictionary) Image_buffers() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("image_buffers"))
	return objectivec.Object{ID: rv}
}
func (e ETDataPointDictionary) SetImage_buffers(value objectivec.IObject) {
	objc.Send[struct{}](e.ID, objc.Sel("setImage_buffers:"), value)
}

