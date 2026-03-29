// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTL4CompilerTaskOptions] class.
var (
	_MTL4CompilerTaskOptionsClass     MTL4CompilerTaskOptionsClass
	_MTL4CompilerTaskOptionsClassOnce sync.Once
)

func getMTL4CompilerTaskOptionsClass() MTL4CompilerTaskOptionsClass {
	_MTL4CompilerTaskOptionsClassOnce.Do(func() {
		_MTL4CompilerTaskOptionsClass = MTL4CompilerTaskOptionsClass{class: objc.GetClass("MTL4CompilerTaskOptions")}
	})
	return _MTL4CompilerTaskOptionsClass
}

// GetMTL4CompilerTaskOptionsClass returns the class object for MTL4CompilerTaskOptions.
func GetMTL4CompilerTaskOptionsClass() MTL4CompilerTaskOptionsClass {
	return getMTL4CompilerTaskOptionsClass()
}

type MTL4CompilerTaskOptionsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTL4CompilerTaskOptionsClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4CompilerTaskOptionsClass) Alloc() MTL4CompilerTaskOptions {
	rv := objc.Send[MTL4CompilerTaskOptions](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// The configuration options that control the behavior of a compilation task
// for a Metal 4 compiler instance.
//
// # Overview
// 
// You can configure task-specific settings that affect a compilation task by
// creating an instance of this class, setting its properties, and passing it
// to one of the applicable methods of an [MTL4Compiler] instance.
//
// # Instance Properties
//
//   - [MTL4CompilerTaskOptions.LookupArchives]: An array of archive instances that can potentially accelerate a compilation task.
//   - [MTL4CompilerTaskOptions.SetLookupArchives]
//
// See: https://developer.apple.com/documentation/Metal/MTL4CompilerTaskOptions
type MTL4CompilerTaskOptions struct {
	objectivec.Object
}

// MTL4CompilerTaskOptionsFromID constructs a [MTL4CompilerTaskOptions] from an objc.ID.
//
// The configuration options that control the behavior of a compilation task
// for a Metal 4 compiler instance.
func MTL4CompilerTaskOptionsFromID(id objc.ID) MTL4CompilerTaskOptions {
	return MTL4CompilerTaskOptions{objectivec.Object{ID: id}}
}
// NOTE: MTL4CompilerTaskOptions adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTL4CompilerTaskOptions] class.
//
// # Instance Properties
//
//   - [IMTL4CompilerTaskOptions.LookupArchives]: An array of archive instances that can potentially accelerate a compilation task.
//   - [IMTL4CompilerTaskOptions.SetLookupArchives]
//
// See: https://developer.apple.com/documentation/Metal/MTL4CompilerTaskOptions
type IMTL4CompilerTaskOptions interface {
	objectivec.IObject

	// Topic: Instance Properties

	// An array of archive instances that can potentially accelerate a compilation task.
	LookupArchives() []objectivec.IObject
	SetLookupArchives(value []objectivec.IObject)
}

// Init initializes the instance.
func (m MTL4CompilerTaskOptions) Init() MTL4CompilerTaskOptions {
	rv := objc.Send[MTL4CompilerTaskOptions](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4CompilerTaskOptions) Autorelease() MTL4CompilerTaskOptions {
	rv := objc.Send[MTL4CompilerTaskOptions](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4CompilerTaskOptions creates a new MTL4CompilerTaskOptions instance.
func NewMTL4CompilerTaskOptions() MTL4CompilerTaskOptions {
	class := getMTL4CompilerTaskOptionsClass()
	rv := objc.Send[MTL4CompilerTaskOptions](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// An array of archive instances that can potentially accelerate a compilation
// task.
//
// # Discussion
// 
// The compiler can reduce the runtime of a compilation task if it finds an
// entry that matches a function description within any of the archives in
// this array. The compiler searches the archives in the order of the
// array’s element.
// 
// Consider adding archives to the array in scenarios that can benefit from
// the runtime savings, such as repeat builds or when your app can share
// compilation results across multiple contexts.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CompilerTaskOptions/lookupArchives
func (m MTL4CompilerTaskOptions) LookupArchives() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("lookupArchives"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
func (m MTL4CompilerTaskOptions) SetLookupArchives(value []objectivec.IObject) {
	objc.Send[struct{}](m.ID, objc.Sel("setLookupArchives:"), objectivec.IObjectSliceToNSArray(value))
}

