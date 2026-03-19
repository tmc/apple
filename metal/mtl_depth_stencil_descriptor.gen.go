// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLDepthStencilDescriptor] class.
var (
	_MTLDepthStencilDescriptorClass     MTLDepthStencilDescriptorClass
	_MTLDepthStencilDescriptorClassOnce sync.Once
)

func getMTLDepthStencilDescriptorClass() MTLDepthStencilDescriptorClass {
	_MTLDepthStencilDescriptorClassOnce.Do(func() {
		_MTLDepthStencilDescriptorClass = MTLDepthStencilDescriptorClass{class: objc.GetClass("MTLDepthStencilDescriptor")}
	})
	return _MTLDepthStencilDescriptorClass
}

// GetMTLDepthStencilDescriptorClass returns the class object for MTLDepthStencilDescriptor.
func GetMTLDepthStencilDescriptorClass() MTLDepthStencilDescriptorClass {
	return getMTLDepthStencilDescriptorClass()
}

type MTLDepthStencilDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLDepthStencilDescriptorClass) Alloc() MTLDepthStencilDescriptor {
	rv := objc.Send[MTLDepthStencilDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An instance that configures new [MTLDepthStencilState] instances.
//
// # Overview
// 
// An [MTLDepthStencilDescriptor] instance is used to define a specific
// configuration of the depth and stencil stages of a rendering pipeline. To
// create an [MTLDepthStencilDescriptor] instance, use standard allocation and
// initialization techniques.
// 
// To enable writing the depth value to a depth attachment, set the
// depthWriteEnabled property to [true].
// 
// The depthCompareFunction property specifies how the depth test is
// performed. If a fragment’s depth value fails the depth test, the fragment
// is discarded. [CompareFunctionLess] is a commonly used value for
// [MTLDepthStencilDescriptor.DepthCompareFunction], because fragment values that are farther away from
// the viewer than the pixel depth value (a previously written fragment) fail
// the depth test and are considered occluded by the earlier depth value.
// 
// The [MTLDepthStencilDescriptor.FrontFaceStencil] and [MTLDepthStencilDescriptor.BackFaceStencil] properties define two
// independent stencil descriptors: one for front-facing primitives and the
// other for back-facing primitives, respectively. Both properties can be set
// to the same MTLStencilDescriptor instance.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Specifying depth operations
//
//   - [MTLDepthStencilDescriptor.DepthCompareFunction]: The comparison that is performed between a fragment’s depth value and the depth value in the attachment, which determines whether to discard the fragment.
//   - [MTLDepthStencilDescriptor.SetDepthCompareFunction]
//   - [MTLDepthStencilDescriptor.DepthWriteEnabled]: A Boolean value that indicates whether depth values can be written to the depth attachment.
//   - [MTLDepthStencilDescriptor.SetDepthWriteEnabled]
//
// # Specifying stencil descriptors for primitives
//
//   - [MTLDepthStencilDescriptor.BackFaceStencil]: The stencil descriptor for back-facing primitives.
//   - [MTLDepthStencilDescriptor.SetBackFaceStencil]
//   - [MTLDepthStencilDescriptor.FrontFaceStencil]: The stencil descriptor for front-facing primitives.
//   - [MTLDepthStencilDescriptor.SetFrontFaceStencil]
//
// # Identifying properties
//
//   - [MTLDepthStencilDescriptor.Label]: A string that identifies this object.
//   - [MTLDepthStencilDescriptor.SetLabel]
//
// See: https://developer.apple.com/documentation/Metal/MTLDepthStencilDescriptor
type MTLDepthStencilDescriptor struct {
	objectivec.Object
}

// MTLDepthStencilDescriptorFromID constructs a [MTLDepthStencilDescriptor] from an objc.ID.
//
// An instance that configures new [MTLDepthStencilState] instances.
func MTLDepthStencilDescriptorFromID(id objc.ID) MTLDepthStencilDescriptor {
	return MTLDepthStencilDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTLDepthStencilDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLDepthStencilDescriptor] class.
//
// # Specifying depth operations
//
//   - [IMTLDepthStencilDescriptor.DepthCompareFunction]: The comparison that is performed between a fragment’s depth value and the depth value in the attachment, which determines whether to discard the fragment.
//   - [IMTLDepthStencilDescriptor.SetDepthCompareFunction]
//   - [IMTLDepthStencilDescriptor.DepthWriteEnabled]: A Boolean value that indicates whether depth values can be written to the depth attachment.
//   - [IMTLDepthStencilDescriptor.SetDepthWriteEnabled]
//
// # Specifying stencil descriptors for primitives
//
//   - [IMTLDepthStencilDescriptor.BackFaceStencil]: The stencil descriptor for back-facing primitives.
//   - [IMTLDepthStencilDescriptor.SetBackFaceStencil]
//   - [IMTLDepthStencilDescriptor.FrontFaceStencil]: The stencil descriptor for front-facing primitives.
//   - [IMTLDepthStencilDescriptor.SetFrontFaceStencil]
//
// # Identifying properties
//
//   - [IMTLDepthStencilDescriptor.Label]: A string that identifies this object.
//   - [IMTLDepthStencilDescriptor.SetLabel]
//
// See: https://developer.apple.com/documentation/Metal/MTLDepthStencilDescriptor
type IMTLDepthStencilDescriptor interface {
	objectivec.IObject

	// Topic: Specifying depth operations

	// The comparison that is performed between a fragment’s depth value and the depth value in the attachment, which determines whether to discard the fragment.
	DepthCompareFunction() MTLCompareFunction
	SetDepthCompareFunction(value MTLCompareFunction)
	// A Boolean value that indicates whether depth values can be written to the depth attachment.
	DepthWriteEnabled() bool
	SetDepthWriteEnabled(value bool)

	// Topic: Specifying stencil descriptors for primitives

	// The stencil descriptor for back-facing primitives.
	BackFaceStencil() IMTLStencilDescriptor
	SetBackFaceStencil(value IMTLStencilDescriptor)
	// The stencil descriptor for front-facing primitives.
	FrontFaceStencil() IMTLStencilDescriptor
	SetFrontFaceStencil(value IMTLStencilDescriptor)

	// Topic: Identifying properties

	// A string that identifies this object.
	Label() string
	SetLabel(value string)
}

// Init initializes the instance.
func (d MTLDepthStencilDescriptor) Init() MTLDepthStencilDescriptor {
	rv := objc.Send[MTLDepthStencilDescriptor](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d MTLDepthStencilDescriptor) Autorelease() MTLDepthStencilDescriptor {
	rv := objc.Send[MTLDepthStencilDescriptor](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLDepthStencilDescriptor creates a new MTLDepthStencilDescriptor instance.
func NewMTLDepthStencilDescriptor() MTLDepthStencilDescriptor {
	class := getMTLDepthStencilDescriptorClass()
	rv := objc.Send[MTLDepthStencilDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The comparison that is performed between a fragment’s depth value and the
// depth value in the attachment, which determines whether to discard the
// fragment.
//
// # Discussion
// 
// The default value is [CompareFunctionAlways], which indicates that the
// depth test always passes and the fragment remains a candidate to replace
// the data at the specified location. For more information on possible
// values, see [MTLCompareFunction].
//
// [MTLCompareFunction]: https://developer.apple.com/documentation/Metal/MTLCompareFunction
//
// See: https://developer.apple.com/documentation/Metal/MTLDepthStencilDescriptor/depthCompareFunction
func (d MTLDepthStencilDescriptor) DepthCompareFunction() MTLCompareFunction {
	rv := objc.Send[MTLCompareFunction](d.ID, objc.Sel("depthCompareFunction"))
	return MTLCompareFunction(rv)
}
func (d MTLDepthStencilDescriptor) SetDepthCompareFunction(value MTLCompareFunction) {
	objc.Send[struct{}](d.ID, objc.Sel("setDepthCompareFunction:"), value)
}
// A Boolean value that indicates whether depth values can be written to the
// depth attachment.
//
// # Discussion
// 
// The default value is [false], which indicates the depth attachment is
// read-only.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Metal/MTLDepthStencilDescriptor/isDepthWriteEnabled
func (d MTLDepthStencilDescriptor) DepthWriteEnabled() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isDepthWriteEnabled"))
	return rv
}
func (d MTLDepthStencilDescriptor) SetDepthWriteEnabled(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setDepthWriteEnabled:"), value)
}
// The stencil descriptor for back-facing primitives.
//
// # Discussion
// 
// The default value is `nil`, which indicates the stencil test is disabled
// for the back-facing primitives. For more information, see
// [MTLStencilDescriptor].
//
// See: https://developer.apple.com/documentation/Metal/MTLDepthStencilDescriptor/backFaceStencil
func (d MTLDepthStencilDescriptor) BackFaceStencil() IMTLStencilDescriptor {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("backFaceStencil"))
	return MTLStencilDescriptorFromID(objc.ID(rv))
}
func (d MTLDepthStencilDescriptor) SetBackFaceStencil(value IMTLStencilDescriptor) {
	objc.Send[struct{}](d.ID, objc.Sel("setBackFaceStencil:"), value)
}
// The stencil descriptor for front-facing primitives.
//
// # Discussion
// 
// The default value is `nil`, which indicates the stencil test is disabled
// for the front-facing primitives. For more information, see
// [MTLStencilDescriptor].
//
// See: https://developer.apple.com/documentation/Metal/MTLDepthStencilDescriptor/frontFaceStencil
func (d MTLDepthStencilDescriptor) FrontFaceStencil() IMTLStencilDescriptor {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("frontFaceStencil"))
	return MTLStencilDescriptorFromID(objc.ID(rv))
}
func (d MTLDepthStencilDescriptor) SetFrontFaceStencil(value IMTLStencilDescriptor) {
	objc.Send[struct{}](d.ID, objc.Sel("setFrontFaceStencil:"), value)
}
// A string that identifies this object.
//
// See: https://developer.apple.com/documentation/Metal/MTLDepthStencilDescriptor/label
func (d MTLDepthStencilDescriptor) Label() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}
func (d MTLDepthStencilDescriptor) SetLabel(value string) {
	objc.Send[struct{}](d.ID, objc.Sel("setLabel:"), objc.String(value))
}

