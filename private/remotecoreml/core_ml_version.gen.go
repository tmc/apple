// Code generated from Apple documentation for remotecoreml. DO NOT EDIT.

package remotecoreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CoreMLVersion] class.
var (
	_CoreMLVersionClass     CoreMLVersionClass
	_CoreMLVersionClassOnce sync.Once
)

func getCoreMLVersionClass() CoreMLVersionClass {
	_CoreMLVersionClassOnce.Do(func() {
		_CoreMLVersionClass = CoreMLVersionClass{class: objc.GetClass("CoreMLVersion")}
	})
	return _CoreMLVersionClass
}

// GetCoreMLVersionClass returns the class object for CoreMLVersion.
func GetCoreMLVersionClass() CoreMLVersionClass {
	return getCoreMLVersionClass()
}

type CoreMLVersionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CoreMLVersionClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CoreMLVersionClass) Alloc() CoreMLVersion {
	rv := objc.Send[CoreMLVersion](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [CoreMLVersion.FrameworkVersionNumber]
//   - [CoreMLVersion.SetFrameworkVersionNumber]
//
// See: https://developer.apple.com/documentation/RemoteCoreML/CoreMLVersion
type CoreMLVersion struct {
	objectivec.Object
}

// CoreMLVersionFromID constructs a [CoreMLVersion] from an objc.ID.
func CoreMLVersionFromID(id objc.ID) CoreMLVersion {
	return CoreMLVersion{objectivec.Object{ID: id}}
}

// Ensure CoreMLVersion implements ICoreMLVersion.
var _ ICoreMLVersion = CoreMLVersion{}

// An interface definition for the [CoreMLVersion] class.
//
// # Methods
//
//   - [ICoreMLVersion.FrameworkVersionNumber]
//   - [ICoreMLVersion.SetFrameworkVersionNumber]
//
// See: https://developer.apple.com/documentation/RemoteCoreML/CoreMLVersion
type ICoreMLVersion interface {
	objectivec.IObject

	// Topic: Methods

	FrameworkVersionNumber() foundation.NSNumber
	SetFrameworkVersionNumber(value foundation.NSNumber)
}

// Init initializes the instance.
func (c CoreMLVersion) Init() CoreMLVersion {
	rv := objc.Send[CoreMLVersion](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CoreMLVersion) Autorelease() CoreMLVersion {
	rv := objc.Send[CoreMLVersion](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCoreMLVersion creates a new CoreMLVersion instance.
func NewCoreMLVersion() CoreMLVersion {
	class := getCoreMLVersionClass()
	rv := objc.Send[CoreMLVersion](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/RemoteCoreML/CoreMLVersion/getInternalFrameworkVersion
func (_CoreMLVersionClass CoreMLVersionClass) GetInternalFrameworkVersion() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_CoreMLVersionClass.class), objc.Sel("getInternalFrameworkVersion"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/RemoteCoreML/CoreMLVersion/frameworkVersionNumber
func (c CoreMLVersion) FrameworkVersionNumber() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("frameworkVersionNumber"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CoreMLVersion) SetFrameworkVersionNumber(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setFrameworkVersionNumber:"), value)
}
