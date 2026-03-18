// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// Optional methods that delegates implement to respond to viewport layout changes.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewportLayoutControllerDelegate
type NSTextViewportLayoutControllerDelegate interface {
	objectivec.IObject

	// The method the framework calls when the layout controller lays out a text layout fragment in the UI.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextViewportLayoutControllerDelegate/textViewportLayoutController(_:configureRenderingSurfaceFor:)
	TextViewportLayoutControllerConfigureRenderingSurfaceForTextLayoutFragment(textViewportLayoutController INSTextViewportLayoutController, textLayoutFragment INSTextLayoutFragment)

	// Returns the current viewport, which is the view visible bounds plus the overdraw area.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextViewportLayoutControllerDelegate/viewportBounds(for:)
	ViewportBoundsForTextViewportLayoutController(textViewportLayoutController INSTextViewportLayoutController) corefoundation.CGRect
}

// NSTextViewportLayoutControllerDelegateObject wraps an existing Objective-C object that conforms to the NSTextViewportLayoutControllerDelegate protocol.
type NSTextViewportLayoutControllerDelegateObject struct {
	objectivec.Object
}
func (o NSTextViewportLayoutControllerDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSTextViewportLayoutControllerDelegateObjectFromID constructs a [NSTextViewportLayoutControllerDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSTextViewportLayoutControllerDelegateObjectFromID(id objc.ID) NSTextViewportLayoutControllerDelegateObject {
	return NSTextViewportLayoutControllerDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The method the framework calls when the layout controller lays out a text
// layout fragment in the UI.
//
// textViewportLayoutController: The [NSTextViewportLayoutController] associated with this text layout
// fragment.
//
// textLayoutFragment: An [NSTextLayoutFragment].
//
// # Discussion
// 
// The delegate presents the text layout fragment in the UI, for example, in a
// sublayer or a subview. Layout information such as `viewportBounds` on
// `textViewportLayoutController` isn’t up to date at the point of this
// call.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewportLayoutControllerDelegate/textViewportLayoutController(_:configureRenderingSurfaceFor:)

func (o NSTextViewportLayoutControllerDelegateObject) TextViewportLayoutControllerConfigureRenderingSurfaceForTextLayoutFragment(textViewportLayoutController INSTextViewportLayoutController, textLayoutFragment INSTextLayoutFragment) {
	
	objc.Send[struct{}](o.ID, objc.Sel("textViewportLayoutController:configureRenderingSurfaceForTextLayoutFragment:"), textViewportLayoutController, textLayoutFragment)
	}

// Returns the current viewport, which is the view visible bounds plus the
// overdraw area.
//
// textViewportLayoutController: The [NSTextViewportLayoutController].
//
// # Return Value
// 
// A [CGRect].
//
// [CGRect]: https://developer.apple.com/documentation/CoreFoundation/CGRect
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewportLayoutControllerDelegate/viewportBounds(for:)

func (o NSTextViewportLayoutControllerDelegateObject) ViewportBoundsForTextViewportLayoutController(textViewportLayoutController INSTextViewportLayoutController) corefoundation.CGRect {
	
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("viewportBoundsForTextViewportLayoutController:"), textViewportLayoutController)
	return rv
	}

// The method the framework calls when the text viewport layout controller
// finishes its layout process.
//
// textViewportLayoutController: The [NSTextViewportLayoutController].
//
// # Discussion
// 
// Layout information on `textViewportLayoutController` is up-to-date at the
// point of this call.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewportLayoutControllerDelegate/textViewportLayoutControllerDidLayout(_:)

func (o NSTextViewportLayoutControllerDelegateObject) TextViewportLayoutControllerDidLayout(textViewportLayoutController INSTextViewportLayoutController) {
	
	objc.Send[struct{}](o.ID, objc.Sel("textViewportLayoutControllerDidLayout:"), textViewportLayoutController)
	}

// The method the framework calls before the text viewport layout controller
// starts its layout process.
//
// textViewportLayoutController: The [NSTextViewportLayoutController].
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewportLayoutControllerDelegate/textViewportLayoutControllerWillLayout(_:)

func (o NSTextViewportLayoutControllerDelegateObject) TextViewportLayoutControllerWillLayout(textViewportLayoutController INSTextViewportLayoutController) {
	
	objc.Send[struct{}](o.ID, objc.Sel("textViewportLayoutControllerWillLayout:"), textViewportLayoutController)
	}

// NSTextViewportLayoutControllerDelegateConfig holds optional typed callbacks for [NSTextViewportLayoutControllerDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nstextviewportlayoutcontrollerdelegate
type NSTextViewportLayoutControllerDelegateConfig struct {

	// Responding to changes in the viewport
	// TextViewportLayoutControllerDidLayout — The method the framework calls when the text viewport layout controller finishes its layout process.
	TextViewportLayoutControllerDidLayout func(textViewportLayoutController NSTextViewportLayoutController)
	// TextViewportLayoutControllerWillLayout — The method the framework calls before the text viewport layout controller starts its layout process.
	TextViewportLayoutControllerWillLayout func(textViewportLayoutController NSTextViewportLayoutController)

	// Other Methods
	// TextViewportLayoutControllerConfigureRenderingSurfaceForTextLayoutFragment — The method the framework calls when the layout controller lays out a text layout fragment in the UI.
	TextViewportLayoutControllerConfigureRenderingSurfaceForTextLayoutFragment func(textViewportLayoutController NSTextViewportLayoutController, textLayoutFragment NSTextLayoutFragment)
}

// NewNSTextViewportLayoutControllerDelegate creates an Objective-C object implementing the [NSTextViewportLayoutControllerDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSTextViewportLayoutControllerDelegateObject] satisfies the [NSTextViewportLayoutControllerDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nstextviewportlayoutcontrollerdelegate
func NewNSTextViewportLayoutControllerDelegate(config NSTextViewportLayoutControllerDelegateConfig) NSTextViewportLayoutControllerDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSTextViewportLayoutControllerDelegate_%d", n)

	var methods []objc.MethodDef

	if config.TextViewportLayoutControllerConfigureRenderingSurfaceForTextLayoutFragment != nil {
		fn := config.TextViewportLayoutControllerConfigureRenderingSurfaceForTextLayoutFragment
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("textViewportLayoutController:configureRenderingSurfaceForTextLayoutFragment:"),
			Fn: func(self objc.ID, _cmd objc.SEL, textViewportLayoutControllerID objc.ID, textLayoutFragmentID objc.ID) {
				textViewportLayoutController := NSTextViewportLayoutControllerFromID(textViewportLayoutControllerID)
				textLayoutFragment := NSTextLayoutFragmentFromID(textLayoutFragmentID)
				fn(textViewportLayoutController, textLayoutFragment)
			},
		})
	}

	if config.TextViewportLayoutControllerDidLayout != nil {
		fn := config.TextViewportLayoutControllerDidLayout
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("textViewportLayoutControllerDidLayout:"),
			Fn: func(self objc.ID, _cmd objc.SEL, textViewportLayoutControllerID objc.ID) {
				textViewportLayoutController := NSTextViewportLayoutControllerFromID(textViewportLayoutControllerID)
				fn(textViewportLayoutController)
			},
		})
	}

	if config.TextViewportLayoutControllerWillLayout != nil {
		fn := config.TextViewportLayoutControllerWillLayout
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("textViewportLayoutControllerWillLayout:"),
			Fn: func(self objc.ID, _cmd objc.SEL, textViewportLayoutControllerID objc.ID) {
				textViewportLayoutController := NSTextViewportLayoutControllerFromID(textViewportLayoutControllerID)
				fn(textViewportLayoutController)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSTextViewportLayoutControllerDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSTextViewportLayoutControllerDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSTextViewportLayoutControllerDelegateObjectFromID(instance)
}

