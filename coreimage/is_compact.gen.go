// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IsCompact] class.
var (
	_IsCompactClass     IsCompactClass
	_IsCompactClassOnce sync.Once
)

func getIsCompactClass() IsCompactClass {
	_IsCompactClassOnce.Do(func() {
		_IsCompactClass = IsCompactClass{class: objc.GetClass("isCompact")}
	})
	return _IsCompactClass
}

// GetIsCompactClass returns the class object for isCompact.
func GetIsCompactClass() IsCompactClass {
	return getIsCompactClass()
}

type IsCompactClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IsCompactClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IsCompactClass) Alloc() IsCompact {
	rv := objc.Send[IsCompact](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIAztecCodeDescriptor/isCompact-c.ivar
type IsCompact struct {
	objectivec.Object
}

// IsCompactFromID constructs a [IsCompact] from an objc.ID.
func IsCompactFromID(id objc.ID) IsCompact {
	return IsCompact{objectivec.Object{ID: id}}
}

// Ensure IsCompact implements IIsCompact.
var _ IIsCompact = IsCompact{}

// An interface definition for the [IsCompact] class.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAztecCodeDescriptor/isCompact-c.ivar
type IIsCompact interface {
	objectivec.IObject
}

// Init initializes the instance.
func (i IsCompact) Init() IsCompact {
	rv := objc.Send[IsCompact](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IsCompact) Autorelease() IsCompact {
	rv := objc.Send[IsCompact](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIsCompact creates a new IsCompact instance.
func NewIsCompact() IsCompact {
	class := getIsCompactClass()
	rv := objc.Send[IsCompact](objc.ID(class.class), objc.Sel("new"))
	return rv
}
