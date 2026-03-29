// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NumBytes] class.
var (
	_NumBytesClass     NumBytesClass
	_NumBytesClassOnce sync.Once
)

func getNumBytesClass() NumBytesClass {
	_NumBytesClassOnce.Do(func() {
		_NumBytesClass = NumBytesClass{class: objc.GetClass("numBytes")}
	})
	return _NumBytesClass
}

// GetNumBytesClass returns the class object for numBytes.
func GetNumBytesClass() NumBytesClass {
	return getNumBytesClass()
}

type NumBytesClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NumBytesClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NumBytesClass) Alloc() NumBytes {
	rv := objc.Send[NumBytes](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSSimpleCString/numBytes
type NumBytes struct {
	objectivec.Object
}

// NumBytesFromID constructs a [NumBytes] from an objc.ID.
func NumBytesFromID(id objc.ID) NumBytes {
	return NumBytes{objectivec.Object{ID: id}}
}
// Ensure NumBytes implements INumBytes.
var _ INumBytes = NumBytes{}

// An interface definition for the [NumBytes] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSSimpleCString/numBytes
type INumBytes interface {
	objectivec.IObject
}

// Init initializes the instance.
func (n NumBytes) Init() NumBytes {
	rv := objc.Send[NumBytes](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NumBytes) Autorelease() NumBytes {
	rv := objc.Send[NumBytes](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNumBytes creates a new NumBytes instance.
func NewNumBytes() NumBytes {
	class := getNumBytesClass()
	rv := objc.Send[NumBytes](objc.ID(class.class), objc.Sel("new"))
	return rv
}

