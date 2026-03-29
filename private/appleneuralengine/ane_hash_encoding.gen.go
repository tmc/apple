// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANEHashEncoding] class.
var (
	_ANEHashEncodingClass     ANEHashEncodingClass
	_ANEHashEncodingClassOnce sync.Once
)

func getANEHashEncodingClass() ANEHashEncodingClass {
	_ANEHashEncodingClassOnce.Do(func() {
		_ANEHashEncodingClass = ANEHashEncodingClass{class: objc.GetClass("_ANEHashEncoding")}
	})
	return _ANEHashEncodingClass
}

// GetANEHashEncodingClass returns the class object for _ANEHashEncoding.
func GetANEHashEncodingClass() ANEHashEncodingClass {
	return getANEHashEncodingClass()
}

type ANEHashEncodingClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac ANEHashEncodingClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEHashEncodingClass) Alloc() ANEHashEncoding {
	rv := objc.Send[ANEHashEncoding](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEHashEncoding
type ANEHashEncoding struct {
	objectivec.Object
}

// ANEHashEncodingFromID constructs a [ANEHashEncoding] from an objc.ID.
func ANEHashEncodingFromID(id objc.ID) ANEHashEncoding {
	return ANEHashEncoding{objectivec.Object{ID: id}}
}
// Ensure ANEHashEncoding implements IANEHashEncoding.
var _ IANEHashEncoding = ANEHashEncoding{}

// An interface definition for the [ANEHashEncoding] class.
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEHashEncoding
type IANEHashEncoding interface {
	objectivec.IObject
}

// Init initializes the instance.
func (a ANEHashEncoding) Init() ANEHashEncoding {
	rv := objc.Send[ANEHashEncoding](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEHashEncoding) Autorelease() ANEHashEncoding {
	rv := objc.Send[ANEHashEncoding](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEHashEncoding creates a new ANEHashEncoding instance.
func NewANEHashEncoding() ANEHashEncoding {
	class := getANEHashEncodingClass()
	rv := objc.Send[ANEHashEncoding](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEHashEncoding/convertToHexString:length:
func (_ANEHashEncodingClass ANEHashEncodingClass) ConvertToHexStringLength(string_ string, length uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEHashEncodingClass.class), objc.Sel("convertToHexString:length:"), unsafe.Pointer(unsafe.StringData(string_ + "\x00")), length)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEHashEncoding/copySHA256For:toBuffer:
func (_ANEHashEncodingClass ANEHashEncodingClass) CopySHA256ForToBuffer(sHA256For objectivec.IObject, buffer string) {
	objc.Send[objc.ID](objc.ID(_ANEHashEncodingClass.class), objc.Sel("copySHA256For:toBuffer:"), sHA256For, unsafe.Pointer(unsafe.StringData(buffer + "\x00")))
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEHashEncoding/hashForString:seed:
func (_ANEHashEncodingClass ANEHashEncodingClass) HashForStringSeed(string_ objectivec.IObject, seed uint32) uint32 {
	rv := objc.Send[uint32](objc.ID(_ANEHashEncodingClass.class), objc.Sel("hashForString:seed:"), string_, seed)
	return rv
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEHashEncoding/hexStringForBytes:length:
func (_ANEHashEncodingClass ANEHashEncodingClass) HexStringForBytesLength(bytes string, length uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEHashEncodingClass.class), objc.Sel("hexStringForBytes:length:"), unsafe.Pointer(unsafe.StringData(bytes + "\x00")), length)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEHashEncoding/hexStringForData:
func (_ANEHashEncodingClass ANEHashEncodingClass) HexStringForData(data objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEHashEncodingClass.class), objc.Sel("hexStringForData:"), data)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEHashEncoding/hexStringForDataArray:
func (_ANEHashEncodingClass ANEHashEncodingClass) HexStringForDataArray(array objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEHashEncodingClass.class), objc.Sel("hexStringForDataArray:"), array)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEHashEncoding/hexStringForString:
func (_ANEHashEncodingClass ANEHashEncodingClass) HexStringForString(string_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEHashEncodingClass.class), objc.Sel("hexStringForString:"), string_)
	return objectivec.Object{ID: rv}
}

