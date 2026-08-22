// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A standard interface for objects that provide shader-based OpenGL rendering effects.
//
// See: https://developer.apple.com/documentation/GLKit/GLKNamedEffect
type GLKNamedEffect interface {
	objectivec.IObject

	// Prepares an effect for OpenGL ES rendering.
	//
	// See: https://developer.apple.com/documentation/GLKit/GLKNamedEffect/prepareToDraw()
	PrepareToDraw()
}

// GLKNamedEffectObject wraps an existing Objective-C object that conforms to the GLKNamedEffect protocol.
type GLKNamedEffectObject struct {
	objectivec.Object
}

func (o GLKNamedEffectObject) BaseObject() objectivec.Object {
	return o.Object
}

// GLKNamedEffectObjectFromID constructs a [GLKNamedEffectObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GLKNamedEffectObjectFromID(id objc.ID) GLKNamedEffectObject {
	return GLKNamedEffectObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Prepares an effect for OpenGL ES rendering.
//
// # Discussion
//
// An effect binds a compiled shader program to the context and returns. Many
// effects also bind data to other OpenGL state variables—see the
// appropriate reference for each effect class.
//
// See: https://developer.apple.com/documentation/GLKit/GLKNamedEffect/prepareToDraw()
func (o GLKNamedEffectObject) PrepareToDraw() {
	objc.Send[struct{}](o.ID, objc.Sel("prepareToDraw"))
}
