// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SwiftNativeNSObject] class.
var (
	_SwiftNativeNSObjectClass     SwiftNativeNSObjectClass
	_SwiftNativeNSObjectClassOnce sync.Once
)

func getSwiftNativeNSObjectClass() SwiftNativeNSObjectClass {
	_SwiftNativeNSObjectClassOnce.Do(func() {
		_SwiftNativeNSObjectClass = SwiftNativeNSObjectClass{class: objc.GetClass("SwiftNativeNSObject")}
	})
	return _SwiftNativeNSObjectClass
}

// GetSwiftNativeNSObjectClass returns the class object for SwiftNativeNSObject.
func GetSwiftNativeNSObjectClass() SwiftNativeNSObjectClass {
	return getSwiftNativeNSObjectClass()
}

type SwiftNativeNSObjectClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (sc SwiftNativeNSObjectClass) Alloc() SwiftNativeNSObject {
	rv := objc.Send[SwiftNativeNSObject](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A parent class referenced by other texttospeech classes. [Full Topic]
type SwiftNativeNSObject struct {
	objectivec.Object
}

// SwiftNativeNSObjectFromID constructs a [SwiftNativeNSObject] from an objc.ID.
//
// A parent class referenced by other texttospeech classes.
func SwiftNativeNSObjectFromID(id objc.ID) SwiftNativeNSObject {
	return SwiftNativeNSObject{objectivec.Object{ID: id}}
}
// Ensure SwiftNativeNSObject implements ISwiftNativeNSObject.
var _ ISwiftNativeNSObject = SwiftNativeNSObject{}

// An interface definition for the [SwiftNativeNSObject] class.
type ISwiftNativeNSObject interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s SwiftNativeNSObject) Init() SwiftNativeNSObject {
	rv := objc.Send[SwiftNativeNSObject](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SwiftNativeNSObject) Autorelease() SwiftNativeNSObject {
	rv := objc.Send[SwiftNativeNSObject](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSwiftNativeNSObject creates a new SwiftNativeNSObject instance.
func NewSwiftNativeNSObject() SwiftNativeNSObject {
	class := getSwiftNativeNSObjectClass()
	rv := objc.Send[SwiftNativeNSObject](objc.ID(class.class), objc.Sel("new"))
	return rv
}

