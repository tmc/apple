// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTL4LibraryDescriptor] class.
var (
	_MTL4LibraryDescriptorClass     MTL4LibraryDescriptorClass
	_MTL4LibraryDescriptorClassOnce sync.Once
)

func getMTL4LibraryDescriptorClass() MTL4LibraryDescriptorClass {
	_MTL4LibraryDescriptorClassOnce.Do(func() {
		_MTL4LibraryDescriptorClass = MTL4LibraryDescriptorClass{class: objc.GetClass("MTL4LibraryDescriptor")}
	})
	return _MTL4LibraryDescriptorClass
}

// GetMTL4LibraryDescriptorClass returns the class object for MTL4LibraryDescriptor.
func GetMTL4LibraryDescriptorClass() MTL4LibraryDescriptorClass {
	return getMTL4LibraryDescriptorClass()
}

type MTL4LibraryDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTL4LibraryDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4LibraryDescriptorClass) Alloc() MTL4LibraryDescriptor {
	rv := objc.Send[MTL4LibraryDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Serves as the base descriptor for creating a Metal library.
//
// # Instance Properties
//
//   - [MTL4LibraryDescriptor.Name]: Assigns an optional name to the Metal library.
//   - [MTL4LibraryDescriptor.SetName]
//   - [MTL4LibraryDescriptor.Options]: Provides compile-time options for the Metal library.
//   - [MTL4LibraryDescriptor.SetOptions]
//   - [MTL4LibraryDescriptor.Source]: Assigns an optional string containing the source code of the shader language program to compile into a Metal library.
//   - [MTL4LibraryDescriptor.SetSource]
//
// See: https://developer.apple.com/documentation/Metal/MTL4LibraryDescriptor
type MTL4LibraryDescriptor struct {
	objectivec.Object
}

// MTL4LibraryDescriptorFromID constructs a [MTL4LibraryDescriptor] from an objc.ID.
//
// Serves as the base descriptor for creating a Metal library.
func MTL4LibraryDescriptorFromID(id objc.ID) MTL4LibraryDescriptor {
	return MTL4LibraryDescriptor{objectivec.Object{ID: id}}
}

// NOTE: MTL4LibraryDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTL4LibraryDescriptor] class.
//
// # Instance Properties
//
//   - [IMTL4LibraryDescriptor.Name]: Assigns an optional name to the Metal library.
//   - [IMTL4LibraryDescriptor.SetName]
//   - [IMTL4LibraryDescriptor.Options]: Provides compile-time options for the Metal library.
//   - [IMTL4LibraryDescriptor.SetOptions]
//   - [IMTL4LibraryDescriptor.Source]: Assigns an optional string containing the source code of the shader language program to compile into a Metal library.
//   - [IMTL4LibraryDescriptor.SetSource]
//
// See: https://developer.apple.com/documentation/Metal/MTL4LibraryDescriptor
type IMTL4LibraryDescriptor interface {
	objectivec.IObject

	// Topic: Instance Properties

	// Assigns an optional name to the Metal library.
	Name() string
	SetName(value string)
	// Provides compile-time options for the Metal library.
	Options() IMTLCompileOptions
	SetOptions(value IMTLCompileOptions)
	// Assigns an optional string containing the source code of the shader language program to compile into a Metal library.
	Source() string
	SetSource(value string)
}

// Init initializes the instance.
func (m MTL4LibraryDescriptor) Init() MTL4LibraryDescriptor {
	rv := objc.Send[MTL4LibraryDescriptor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4LibraryDescriptor) Autorelease() MTL4LibraryDescriptor {
	rv := objc.Send[MTL4LibraryDescriptor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4LibraryDescriptor creates a new MTL4LibraryDescriptor instance.
func NewMTL4LibraryDescriptor() MTL4LibraryDescriptor {
	class := getMTL4LibraryDescriptorClass()
	rv := objc.Send[MTL4LibraryDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Assigns an optional name to the Metal library.
//
// See: https://developer.apple.com/documentation/Metal/MTL4LibraryDescriptor/name
func (m MTL4LibraryDescriptor) Name() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (m MTL4LibraryDescriptor) SetName(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setName:"), objc.String(value))
}

// Provides compile-time options for the Metal library.
//
// See: https://developer.apple.com/documentation/Metal/MTL4LibraryDescriptor/options
func (m MTL4LibraryDescriptor) Options() IMTLCompileOptions {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("options"))
	return MTLCompileOptionsFromID(objc.ID(rv))
}
func (m MTL4LibraryDescriptor) SetOptions(value IMTLCompileOptions) {
	objc.Send[struct{}](m.ID, objc.Sel("setOptions:"), value)
}

// Assigns an optional string containing the source code of the shader
// language program to compile into a Metal library.
//
// See: https://developer.apple.com/documentation/Metal/MTL4LibraryDescriptor/source
func (m MTL4LibraryDescriptor) Source() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("source"))
	return foundation.NSStringFromID(rv).String()
}
func (m MTL4LibraryDescriptor) SetSource(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setSource:"), objc.String(value))
}
