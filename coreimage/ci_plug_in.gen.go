// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CIPlugIn] class.
var (
	_CIPlugInClass     CIPlugInClass
	_CIPlugInClassOnce sync.Once
)

func getCIPlugInClass() CIPlugInClass {
	_CIPlugInClassOnce.Do(func() {
		_CIPlugInClass = CIPlugInClass{class: objc.GetClass("CIPlugIn")}
	})
	return _CIPlugInClass
}

// GetCIPlugInClass returns the class object for CIPlugIn.
func GetCIPlugInClass() CIPlugInClass {
	return getCIPlugInClass()
}

type CIPlugInClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (cc CIPlugInClass) Alloc() CIPlugIn {
	rv := objc.Send[CIPlugIn](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// The mechanism for loading image units in macOS.
//
// # Overview
// 
// An image unit is an image processing bundle that contains one or more Core
// Image filters. Th`e.Plugin()` extension indicates one or more filters
// packaged as an image unit.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPlugIn
type CIPlugIn struct {
	objectivec.Object
}

// CIPlugInFromID constructs a [CIPlugIn] from an objc.ID.
//
// The mechanism for loading image units in macOS.
func CIPlugInFromID(id objc.ID) CIPlugIn {
	return CIPlugIn{objectivec.Object{ID: id}}
}
// NOTE: CIPlugIn adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CIPlugIn] class.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPlugIn
type ICIPlugIn interface {
	objectivec.IObject
}

// Init initializes the instance.
func (p CIPlugIn) Init() CIPlugIn {
	rv := objc.Send[CIPlugIn](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p CIPlugIn) Autorelease() CIPlugIn {
	rv := objc.Send[CIPlugIn](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewCIPlugIn creates a new CIPlugIn instance.
func NewCIPlugIn() CIPlugIn {
	class := getCIPlugInClass()
	rv := objc.Send[CIPlugIn](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Scans directories for plugins.
//
// # Discussion
// 
// This call scans for plugins with the extension `XCUIElementTypePlugin` in
// the following directories:
// 
// - /Library/Graphics/Image Units
// - ~Library/Graphics/Image Units
// 
// This call adds new plug-ins. It doesn’t remove any plug-ins.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPlugIn/loadNonExecutablePlugIns()
func (_CIPlugInClass CIPlugInClass) LoadNonExecutablePlugIns() {
	objc.Send[objc.ID](objc.ID(_CIPlugInClass.class), objc.Sel("loadNonExecutablePlugIns"))
}
// Loads a non-executable plug-in specified by its URL.
//
// url: The location of the plugin to load.
//
// # Discussion
// 
// If the filters contain executable code the plugin isn’t loaded.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPlugIn/loadNonExecutablePlugIn(_:)
func (_CIPlugInClass CIPlugInClass) LoadNonExecutablePlugIn(url foundation.INSURL) {
	objc.Send[objc.ID](objc.ID(_CIPlugInClass.class), objc.Sel("loadNonExecutablePlugIn:"), url)
}

