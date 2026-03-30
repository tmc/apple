// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WKNavigation] class.
var (
	_WKNavigationClass     WKNavigationClass
	_WKNavigationClassOnce sync.Once
)

func getWKNavigationClass() WKNavigationClass {
	_WKNavigationClassOnce.Do(func() {
		_WKNavigationClass = WKNavigationClass{class: objc.GetClass("WKNavigation")}
	})
	return _WKNavigationClass
}

// GetWKNavigationClass returns the class object for WKNavigation.
func GetWKNavigationClass() WKNavigationClass {
	return getWKNavigationClass()
}

type WKNavigationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WKNavigationClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WKNavigationClass) Alloc() WKNavigation {
	rv := objc.Send[WKNavigation](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// An object that tracks the loading progress of a webpage.
//
// # Overview
//
// A [WKNavigation] object uniquely identifies a load request for a webpage.
// When you ask a web view to load content or navigate to a page, the web view
// returns a [WKNavigation] object that identifies your request. As the load
// operation progresses, the web view reports progress of that operation to
// various methods of its navigation delegate, passing them the matching
// [WKNavigation] object.
//
// # Getting the Content Mode
//
//   - [WKNavigation.EffectiveContentMode]: The content mode WebKit uses to load the webpage.
//
// See: https://developer.apple.com/documentation/WebKit/WKNavigation
type WKNavigation struct {
	objectivec.Object
}

// WKNavigationFromID constructs a [WKNavigation] from an objc.ID.
//
// An object that tracks the loading progress of a webpage.
func WKNavigationFromID(id objc.ID) WKNavigation {
	return WKNavigation{objectivec.Object{ID: id}}
}

// NOTE: WKNavigation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [WKNavigation] class.
//
// # Getting the Content Mode
//
//   - [IWKNavigation.EffectiveContentMode]: The content mode WebKit uses to load the webpage.
//
// See: https://developer.apple.com/documentation/WebKit/WKNavigation
type IWKNavigation interface {
	objectivec.IObject

	// Topic: Getting the Content Mode

	// The content mode WebKit uses to load the webpage.
	EffectiveContentMode() WKContentMode
}

// Init initializes the instance.
func (n WKNavigation) Init() WKNavigation {
	rv := objc.Send[WKNavigation](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n WKNavigation) Autorelease() WKNavigation {
	rv := objc.Send[WKNavigation](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewWKNavigation creates a new WKNavigation instance.
func NewWKNavigation() WKNavigation {
	class := getWKNavigationClass()
	rv := objc.Send[WKNavigation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The content mode WebKit uses to load the webpage.
//
// See: https://developer.apple.com/documentation/WebKit/WKNavigation/effectiveContentMode
func (n WKNavigation) EffectiveContentMode() WKContentMode {
	rv := objc.Send[WKContentMode](n.ID, objc.Sel("effectiveContentMode"))
	return WKContentMode(rv)
}
