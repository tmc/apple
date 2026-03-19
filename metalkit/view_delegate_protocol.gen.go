// Code generated from Apple documentation for MetalKit. DO NOT EDIT.

package metalkit

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// Methods for responding to a MetalKit view’s drawing and resizing events.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKViewDelegate
type MTKViewDelegate interface {
	objectivec.IObject

	// Updates the view’s contents upon receiving a change in layout, resolution, or size.
	//
	// See: https://developer.apple.com/documentation/MetalKit/MTKViewDelegate/mtkView(_:drawableSizeWillChange:)
	MtkViewDrawableSizeWillChange(view IMTKView, size corefoundation.CGSize)

	// Draws the view’s contents.
	//
	// See: https://developer.apple.com/documentation/MetalKit/MTKViewDelegate/draw(in:)
	DrawInMTKView(view IMTKView)
}

// MTKViewDelegateObject wraps an existing Objective-C object that conforms to the MTKViewDelegate protocol.
type MTKViewDelegateObject struct {
	objectivec.Object
}
func (o MTKViewDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTKViewDelegateObjectFromID constructs a [MTKViewDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTKViewDelegateObjectFromID(id objc.ID) MTKViewDelegateObject {
	return MTKViewDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Updates the view’s contents upon receiving a change in layout,
// resolution, or size.
//
// view: The view requesting that its contents be updated.
//
// size: The view’s new drawable size.
//
// # Discussion
// 
// Use this method to recompute any view or projection matrices, or to
// regenerate any buffers to be compatible with the view’s new size.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKViewDelegate/mtkView(_:drawableSizeWillChange:)
func (o MTKViewDelegateObject) MtkViewDrawableSizeWillChange(view IMTKView, size corefoundation.CGSize) {
	
	objc.Send[struct{}](o.ID, objc.Sel("mtkView:drawableSizeWillChange:"), view, size)
	}
// Draws the view’s contents.
//
// view: The view requesting that its contents be redrawn.
//
// # Discussion
// 
// This method is called on the delegate when it is asked to render into the
// view.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKViewDelegate/draw(in:)
func (o MTKViewDelegateObject) DrawInMTKView(view IMTKView) {
	
	objc.Send[struct{}](o.ID, objc.Sel("drawInMTKView:"), view)
	}

// MTKViewDelegateConfig holds optional typed callbacks for [MTKViewDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/metalkit/mtkviewdelegate
type MTKViewDelegateConfig struct {

	// Other Methods
	// DrawInMTKView — Draws the view’s contents.
	DrawInMTKView func(view MTKView)
}

// NewMTKViewDelegate creates an Objective-C object implementing the [MTKViewDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [MTKViewDelegateObject] satisfies the [MTKViewDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/metalkit/mtkviewdelegate
func NewMTKViewDelegate(config MTKViewDelegateConfig) MTKViewDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoMTKViewDelegate_%d", n)

	var methods []objc.MethodDef

	if config.DrawInMTKView != nil {
		fn := config.DrawInMTKView
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("drawInMTKView:"),
			Fn: func(self objc.ID, _cmd objc.SEL, viewID objc.ID) {
				view := MTKViewFromID(viewID)
				fn(view)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("MTKViewDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewMTKViewDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return MTKViewDelegateObjectFromID(instance)
}

