// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLFunctionStitchingAttributeAlwaysInline] class.
var (
	_MTLFunctionStitchingAttributeAlwaysInlineClass     MTLFunctionStitchingAttributeAlwaysInlineClass
	_MTLFunctionStitchingAttributeAlwaysInlineClassOnce sync.Once
)

func getMTLFunctionStitchingAttributeAlwaysInlineClass() MTLFunctionStitchingAttributeAlwaysInlineClass {
	_MTLFunctionStitchingAttributeAlwaysInlineClassOnce.Do(func() {
		_MTLFunctionStitchingAttributeAlwaysInlineClass = MTLFunctionStitchingAttributeAlwaysInlineClass{class: objc.GetClass("MTLFunctionStitchingAttributeAlwaysInline")}
	})
	return _MTLFunctionStitchingAttributeAlwaysInlineClass
}

// GetMTLFunctionStitchingAttributeAlwaysInlineClass returns the class object for MTLFunctionStitchingAttributeAlwaysInline.
func GetMTLFunctionStitchingAttributeAlwaysInlineClass() MTLFunctionStitchingAttributeAlwaysInlineClass {
	return getMTLFunctionStitchingAttributeAlwaysInlineClass()
}

type MTLFunctionStitchingAttributeAlwaysInlineClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLFunctionStitchingAttributeAlwaysInlineClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLFunctionStitchingAttributeAlwaysInlineClass) Alloc() MTLFunctionStitchingAttributeAlwaysInline {
	rv := objc.Send[MTLFunctionStitchingAttributeAlwaysInline](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An attribute to specify that Metal needs to inline all of the function
// calls when generating the stitched function.
//
// # Overview
//
// To inline functions in a call graph, instantiate an instance of this class
// and assign it as an attribute on the [MTLFunctionStitchingGraph].
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingAttributeAlwaysInline
type MTLFunctionStitchingAttributeAlwaysInline struct {
	objectivec.Object
}

// MTLFunctionStitchingAttributeAlwaysInlineFromID constructs a [MTLFunctionStitchingAttributeAlwaysInline] from an objc.ID.
//
// An attribute to specify that Metal needs to inline all of the function
// calls when generating the stitched function.
func MTLFunctionStitchingAttributeAlwaysInlineFromID(id objc.ID) MTLFunctionStitchingAttributeAlwaysInline {
	return MTLFunctionStitchingAttributeAlwaysInline{objectivec.Object{ID: id}}
}

// NOTE: MTLFunctionStitchingAttributeAlwaysInline adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLFunctionStitchingAttributeAlwaysInline] class.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingAttributeAlwaysInline
type IMTLFunctionStitchingAttributeAlwaysInline interface {
	objectivec.IObject
	MTLFunctionStitchingAttribute

	// A list of attributes to configure how the Metal device object generates the new stitched function.
	Attributes() MTLFunctionStitchingAttribute
	SetAttributes(value MTLFunctionStitchingAttribute)
}

// Init initializes the instance.
func (f MTLFunctionStitchingAttributeAlwaysInline) Init() MTLFunctionStitchingAttributeAlwaysInline {
	rv := objc.Send[MTLFunctionStitchingAttributeAlwaysInline](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f MTLFunctionStitchingAttributeAlwaysInline) Autorelease() MTLFunctionStitchingAttributeAlwaysInline {
	rv := objc.Send[MTLFunctionStitchingAttributeAlwaysInline](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLFunctionStitchingAttributeAlwaysInline creates a new MTLFunctionStitchingAttributeAlwaysInline instance.
func NewMTLFunctionStitchingAttributeAlwaysInline() MTLFunctionStitchingAttributeAlwaysInline {
	class := getMTLFunctionStitchingAttributeAlwaysInlineClass()
	rv := objc.Send[MTLFunctionStitchingAttributeAlwaysInline](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A list of attributes to configure how the Metal device object generates the
// new stitched function.
//
// See: https://developer.apple.com/documentation/metal/mtlfunctionstitchinggraph/attributes
func (f MTLFunctionStitchingAttributeAlwaysInline) Attributes() MTLFunctionStitchingAttribute {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("attributes"))
	return MTLFunctionStitchingAttributeObjectFromID(rv)
}
func (f MTLFunctionStitchingAttributeAlwaysInline) SetAttributes(value MTLFunctionStitchingAttribute) {
	objc.Send[struct{}](f.ID, objc.Sel("setAttributes:"), value)
}

// Protocol methods for MTLFunctionStitchingAttribute
