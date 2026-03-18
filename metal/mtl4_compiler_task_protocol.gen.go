// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A reference to an asynchronous compilation task that you initiate from a compiler instance.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CompilerTask
type MTL4CompilerTask interface {
	objectivec.IObject

	// Returns the compiler instance that this asynchronous compiler task belongs to.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CompilerTask/compiler
	Compiler() MTL4Compiler

	// Returns the compiler task status.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CompilerTask/status
	Status() MTL4CompilerTaskStatus

	// Waits synchronously for this compile task to complete by blocking the calling thread.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CompilerTask/waitUntilCompleted
	WaitUntilCompleted()
}

// MTL4CompilerTaskObject wraps an existing Objective-C object that conforms to the MTL4CompilerTask protocol.
type MTL4CompilerTaskObject struct {
	objectivec.Object
}
func (o MTL4CompilerTaskObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTL4CompilerTaskObjectFromID constructs a [MTL4CompilerTaskObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTL4CompilerTaskObjectFromID(id objc.ID) MTL4CompilerTaskObject {
	return MTL4CompilerTaskObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns the compiler instance that this asynchronous compiler task belongs
// to.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CompilerTask/compiler

func (o MTL4CompilerTaskObject) Compiler() MTL4Compiler {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("compiler"))
	return MTL4CompilerObjectFromID(rv)
	}

// Returns the compiler task status.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CompilerTask/status

func (o MTL4CompilerTaskObject) Status() MTL4CompilerTaskStatus {
	
	rv := objc.Send[MTL4CompilerTaskStatus](o.ID, objc.Sel("status"))
	return rv
	}

// Waits synchronously for this compile task to complete by blocking the
// calling thread.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CompilerTask/waitUntilCompleted

func (o MTL4CompilerTaskObject) WaitUntilCompleted() {
	
	objc.Send[struct{}](o.ID, objc.Sel("waitUntilCompleted"))
	}

