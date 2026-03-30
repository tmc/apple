// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DIURL] class.
var (
	_DIURLClass     DIURLClass
	_DIURLClassOnce sync.Once
)

func getDIURLClass() DIURLClass {
	_DIURLClassOnce.Do(func() {
		_DIURLClass = DIURLClass{class: objc.GetClass("DIURL")}
	})
	return _DIURLClass
}

// GetDIURLClass returns the class object for DIURL.
func GetDIURLClass() DIURLClass {
	return getDIURLClass()
}

type DIURLClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIURLClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIURLClass) Alloc() DIURL {
	rv := objc.Send[DIURL](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [DIURL.IsPlugin]
//   - [DIURL.PluginName]
//   - [DIURL.PluginParams]
//   - [DIURL.InitWithPluginNameParams]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIURL
type DIURL struct {
	foundation.NSURL
}

// DIURLFromID constructs a [DIURL] from an objc.ID.
func DIURLFromID(id objc.ID) DIURL {
	return DIURL{NSURL: foundation.NSURLFromID(id)}
}

// Ensure DIURL implements IDIURL.
var _ IDIURL = DIURL{}

// An interface definition for the [DIURL] class.
//
// # Methods
//
//   - [IDIURL.IsPlugin]
//   - [IDIURL.PluginName]
//   - [IDIURL.PluginParams]
//   - [IDIURL.InitWithPluginNameParams]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIURL
type IDIURL interface {
	foundation.INSURL

	// Topic: Methods

	IsPlugin() bool
	PluginName() string
	PluginParams() foundation.INSDictionary
	InitWithPluginNameParams(name objectivec.IObject, params objectivec.IObject) DIURL
}

// Init initializes the instance.
func (d DIURL) Init() DIURL {
	rv := objc.Send[DIURL](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIURL) Autorelease() DIURL {
	rv := objc.Send[DIURL](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIURL creates a new DIURL instance.
func NewDIURL() DIURL {
	class := getDIURLClass()
	rv := objc.Send[DIURL](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIURL/initWithCoder:
func NewDIURLWithCoder(coder objectivec.IObject) DIURL {
	instance := getDIURLClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DIURLFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/DIURL/initWithPluginName:params:
func NewDIURLWithPluginNameParams(name objectivec.IObject, params objectivec.IObject) DIURL {
	instance := getDIURLClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPluginName:params:"), name, params)
	return DIURLFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/DIURL/initWithPluginName:params:
func (d DIURL) InitWithPluginNameParams(name objectivec.IObject, params objectivec.IObject) DIURL {
	rv := objc.Send[DIURL](d.ID, objc.Sel("initWithPluginName:params:"), name, params)
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIURL/newDIURLWithNSURL:
func (_DIURLClass DIURLClass) NewDIURLWithNSURL(nsurl foundation.INSURL) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_DIURLClass.class), objc.Sel("newDIURLWithNSURL:"), nsurl)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/DiskImages2/DIURL/newDIURLWithPluginName:params:
func (_DIURLClass DIURLClass) NewDIURLWithPluginNameParams(name objectivec.IObject, params objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_DIURLClass.class), objc.Sel("newDIURLWithPluginName:params:"), name, params)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/DiskImages2/DIURL/supportsSecureCoding
func (_DIURLClass DIURLClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_DIURLClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIURL/isPlugin
func (d DIURL) IsPlugin() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isPlugin"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIURL/pluginName
func (d DIURL) PluginName() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("pluginName"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/DiskImages2/DIURL/pluginParams
func (d DIURL) PluginParams() foundation.INSDictionary {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("pluginParams"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
