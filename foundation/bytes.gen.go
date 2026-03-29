// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [Bytes] class.
var (
	_BytesClass     BytesClass
	_BytesClassOnce sync.Once
)

func getBytesClass() BytesClass {
	_BytesClassOnce.Do(func() {
		_BytesClass = BytesClass{class: objc.GetClass("bytes")}
	})
	return _BytesClass
}

// GetBytesClass returns the class object for bytes.
func GetBytesClass() BytesClass {
	return getBytesClass()
}

type BytesClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (bc BytesClass) Class() objc.Class {
	return bc.class
}

// Alloc allocates memory for a new instance of the class.
func (bc BytesClass) Alloc() Bytes {
	rv := objc.Send[Bytes](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSSimpleCString/bytes
type Bytes struct {
	objectivec.Object
}

// BytesFromID constructs a [Bytes] from an objc.ID.
func BytesFromID(id objc.ID) Bytes {
	return Bytes{objectivec.Object{ID: id}}
}
// Ensure Bytes implements IBytes.
var _ IBytes = Bytes{}

// An interface definition for the [Bytes] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSSimpleCString/bytes
type IBytes interface {
	objectivec.IObject
}

// Init initializes the instance.
func (b Bytes) Init() Bytes {
	rv := objc.Send[Bytes](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b Bytes) Autorelease() Bytes {
	rv := objc.Send[Bytes](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewBytes creates a new Bytes instance.
func NewBytes() Bytes {
	class := getBytesClass()
	rv := objc.Send[Bytes](objc.ID(class.class), objc.Sel("new"))
	return rv
}

