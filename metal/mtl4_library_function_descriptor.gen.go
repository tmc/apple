// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [MTL4LibraryFunctionDescriptor] class.
var (
	_MTL4LibraryFunctionDescriptorClass     MTL4LibraryFunctionDescriptorClass
	_MTL4LibraryFunctionDescriptorClassOnce sync.Once
)

func getMTL4LibraryFunctionDescriptorClass() MTL4LibraryFunctionDescriptorClass {
	_MTL4LibraryFunctionDescriptorClassOnce.Do(func() {
		_MTL4LibraryFunctionDescriptorClass = MTL4LibraryFunctionDescriptorClass{class: objc.GetClass("MTL4LibraryFunctionDescriptor")}
	})
	return _MTL4LibraryFunctionDescriptorClass
}

// GetMTL4LibraryFunctionDescriptorClass returns the class object for MTL4LibraryFunctionDescriptor.
func GetMTL4LibraryFunctionDescriptorClass() MTL4LibraryFunctionDescriptorClass {
	return getMTL4LibraryFunctionDescriptorClass()
}

type MTL4LibraryFunctionDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4LibraryFunctionDescriptorClass) Alloc() MTL4LibraryFunctionDescriptor {
	rv := objc.Send[MTL4LibraryFunctionDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// Describes a shader function from a Metal library.
//
// # Instance Properties
//
//   - [MTL4LibraryFunctionDescriptor.Library]: Returns a reference to the library containing the function.
//   - [MTL4LibraryFunctionDescriptor.SetLibrary]
//   - [MTL4LibraryFunctionDescriptor.Name]: Assigns a name to the function.
//   - [MTL4LibraryFunctionDescriptor.SetName]
//
// See: https://developer.apple.com/documentation/Metal/MTL4LibraryFunctionDescriptor
type MTL4LibraryFunctionDescriptor struct {
	MTL4FunctionDescriptor
}

// MTL4LibraryFunctionDescriptorFromID constructs a [MTL4LibraryFunctionDescriptor] from an objc.ID.
//
// Describes a shader function from a Metal library.
func MTL4LibraryFunctionDescriptorFromID(id objc.ID) MTL4LibraryFunctionDescriptor {
	return MTL4LibraryFunctionDescriptor{MTL4FunctionDescriptor: MTL4FunctionDescriptorFromID(id)}
}
// NOTE: MTL4LibraryFunctionDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MTL4LibraryFunctionDescriptor] class.
//
// # Instance Properties
//
//   - [IMTL4LibraryFunctionDescriptor.Library]: Returns a reference to the library containing the function.
//   - [IMTL4LibraryFunctionDescriptor.SetLibrary]
//   - [IMTL4LibraryFunctionDescriptor.Name]: Assigns a name to the function.
//   - [IMTL4LibraryFunctionDescriptor.SetName]
//
// See: https://developer.apple.com/documentation/Metal/MTL4LibraryFunctionDescriptor
type IMTL4LibraryFunctionDescriptor interface {
	IMTL4FunctionDescriptor

	// Topic: Instance Properties

	// Returns a reference to the library containing the function.
	Library() MTLLibrary
	SetLibrary(value MTLLibrary)
	// Assigns a name to the function.
	Name() string
	SetName(value string)
}





// Init initializes the instance.
func (m MTL4LibraryFunctionDescriptor) Init() MTL4LibraryFunctionDescriptor {
	rv := objc.Send[MTL4LibraryFunctionDescriptor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4LibraryFunctionDescriptor) Autorelease() MTL4LibraryFunctionDescriptor {
	rv := objc.Send[MTL4LibraryFunctionDescriptor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4LibraryFunctionDescriptor creates a new MTL4LibraryFunctionDescriptor instance.
func NewMTL4LibraryFunctionDescriptor() MTL4LibraryFunctionDescriptor {
	class := getMTL4LibraryFunctionDescriptorClass()
	rv := objc.Send[MTL4LibraryFunctionDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}





















// Returns a reference to the library containing the function.
//
// See: https://developer.apple.com/documentation/Metal/MTL4LibraryFunctionDescriptor/library
func (m MTL4LibraryFunctionDescriptor) Library() MTLLibrary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("library"))
	return MTLLibraryObjectFromID(rv)
}
func (m MTL4LibraryFunctionDescriptor) SetLibrary(value MTLLibrary) {
	objc.Send[struct{}](m.ID, objc.Sel("setLibrary:"), value)
}



// Assigns a name to the function.
//
// See: https://developer.apple.com/documentation/Metal/MTL4LibraryFunctionDescriptor/name
func (m MTL4LibraryFunctionDescriptor) Name() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (m MTL4LibraryFunctionDescriptor) SetName(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setName:"), objc.String(value))
}
























