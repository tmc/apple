// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DIEncryptionUnlocker] class.
var (
	_DIEncryptionUnlockerClass     DIEncryptionUnlockerClass
	_DIEncryptionUnlockerClassOnce sync.Once
)

func getDIEncryptionUnlockerClass() DIEncryptionUnlockerClass {
	_DIEncryptionUnlockerClassOnce.Do(func() {
		_DIEncryptionUnlockerClass = DIEncryptionUnlockerClass{class: objc.GetClass("DIEncryptionUnlocker")}
	})
	return _DIEncryptionUnlockerClass
}

// GetDIEncryptionUnlockerClass returns the class object for DIEncryptionUnlocker.
func GetDIEncryptionUnlockerClass() DIEncryptionUnlockerClass {
	return getDIEncryptionUnlockerClass()
}

type DIEncryptionUnlockerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIEncryptionUnlockerClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIEncryptionUnlockerClass) Alloc() DIEncryptionUnlocker {
	rv := objc.Send[DIEncryptionUnlocker](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionUnlocker
type DIEncryptionUnlocker struct {
	DIEncryptionFrontend
}

// DIEncryptionUnlockerFromID constructs a [DIEncryptionUnlocker] from an objc.ID.
func DIEncryptionUnlockerFromID(id objc.ID) DIEncryptionUnlocker {
	return DIEncryptionUnlocker{DIEncryptionFrontend: DIEncryptionFrontendFromID(id)}
}
// Ensure DIEncryptionUnlocker implements IDIEncryptionUnlocker.
var _ IDIEncryptionUnlocker = DIEncryptionUnlocker{}

// An interface definition for the [DIEncryptionUnlocker] class.
//
// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionUnlocker
type IDIEncryptionUnlocker interface {
	IDIEncryptionFrontend
}

// Init initializes the instance.
func (d DIEncryptionUnlocker) Init() DIEncryptionUnlocker {
	rv := objc.Send[DIEncryptionUnlocker](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIEncryptionUnlocker) Autorelease() DIEncryptionUnlocker {
	rv := objc.Send[DIEncryptionUnlocker](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIEncryptionUnlocker creates a new DIEncryptionUnlocker instance.
func NewDIEncryptionUnlocker() DIEncryptionUnlocker {
	class := getDIEncryptionUnlockerClass()
	rv := objc.Send[DIEncryptionUnlocker](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/initWithCoder:
func NewDIEncryptionUnlockerWithCoder(coder objectivec.IObject) DIEncryptionUnlocker {
	instance := getDIEncryptionUnlockerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DIEncryptionUnlockerFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/initWithParams:
func NewDIEncryptionUnlockerWithParams(params objectivec.IObject) DIEncryptionUnlocker {
	instance := getDIEncryptionUnlockerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParams:"), params)
	return DIEncryptionUnlockerFromID(rv)
}

