// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLArchitecture] class.
var (
	_MTLArchitectureClass     MTLArchitectureClass
	_MTLArchitectureClassOnce sync.Once
)

func getMTLArchitectureClass() MTLArchitectureClass {
	_MTLArchitectureClassOnce.Do(func() {
		_MTLArchitectureClass = MTLArchitectureClass{class: objc.GetClass("MTLArchitecture")}
	})
	return _MTLArchitectureClass
}

// GetMTLArchitectureClass returns the class object for MTLArchitecture.
func GetMTLArchitectureClass() MTLArchitectureClass {
	return getMTLArchitectureClass()
}

type MTLArchitectureClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLArchitectureClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLArchitectureClass) Alloc() MTLArchitecture {
	rv := objc.Send[MTLArchitecture](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A class that contains the architectural details of a GPU device.
//
// # Inspecting a GPU device’s architecture details
//
//   - [MTLArchitecture.Name]: The name of a GPU device’s architecture.
//
// See: https://developer.apple.com/documentation/Metal/MTLArchitecture
type MTLArchitecture struct {
	objectivec.Object
}

// MTLArchitectureFromID constructs a [MTLArchitecture] from an objc.ID.
//
// A class that contains the architectural details of a GPU device.
func MTLArchitectureFromID(id objc.ID) MTLArchitecture {
	return MTLArchitecture{objectivec.Object{ID: id}}
}

// NOTE: MTLArchitecture adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLArchitecture] class.
//
// # Inspecting a GPU device’s architecture details
//
//   - [IMTLArchitecture.Name]: The name of a GPU device’s architecture.
//
// See: https://developer.apple.com/documentation/Metal/MTLArchitecture
type IMTLArchitecture interface {
	objectivec.IObject

	// Topic: Inspecting a GPU device’s architecture details

	// The name of a GPU device’s architecture.
	Name() string
}

// Init initializes the instance.
func (a MTLArchitecture) Init() MTLArchitecture {
	rv := objc.Send[MTLArchitecture](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MTLArchitecture) Autorelease() MTLArchitecture {
	rv := objc.Send[MTLArchitecture](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLArchitecture creates a new MTLArchitecture instance.
func NewMTLArchitecture() MTLArchitecture {
	class := getMTLArchitectureClass()
	rv := objc.Send[MTLArchitecture](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The name of a GPU device’s architecture.
//
// # Discussion
//
// The property’s value is equivalent to the output from the `metal-arch`
// command line tool on the same system.
//
// Apps can use this property’s value to make decisions at runtime. For
// example, an app could retrieve a GPU-specific file from its developer’s
// content delivery network (CDN), such as a shader library or binary archive.
// See [Shader libraries] and [Shader library and archive creation] for more
// information.
//
// See: https://developer.apple.com/documentation/Metal/MTLArchitecture/name
//
// [Shader libraries]: https://developer.apple.com/documentation/Metal/shader-libraries
// [Shader library and archive creation]: https://developer.apple.com/documentation/Metal/shader-library-and-archive-creation
func (a MTLArchitecture) Name() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
