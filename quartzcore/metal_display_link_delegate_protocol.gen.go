// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// A protocol your app implements to respond to callbacks from Core Animation for a Metal display link.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMetalDisplayLinkDelegate
type CAMetalDisplayLinkDelegate interface {
	objectivec.IObject

	// A method the system calls to notify your app when it plans to update the display.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAMetalDisplayLinkDelegate/metalDisplayLink(_:needsUpdate:)
	MetalDisplayLinkNeedsUpdate(link ICAMetalDisplayLink, update ICAMetalDisplayLinkUpdate)
}

// CAMetalDisplayLinkDelegateObject wraps an existing Objective-C object that conforms to the CAMetalDisplayLinkDelegate protocol.
type CAMetalDisplayLinkDelegateObject struct {
	objectivec.Object
}

func (o CAMetalDisplayLinkDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// CAMetalDisplayLinkDelegateObjectFromID constructs a [CAMetalDisplayLinkDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CAMetalDisplayLinkDelegateObjectFromID(id objc.ID) CAMetalDisplayLinkDelegateObject {
	return CAMetalDisplayLinkDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A method the system calls to notify your app when it plans to update the
// display.
//
// link: A Metal display link instance the system notifies.
//
// update: An update instance that contains the time the system intends to update the
// display, a [CAMetalDrawable] instance, and a deadline to call its
// [present()] method.
//
// # Discussion
//
// In this method’s implementation, perform your app’s rendering on the
// [Layer] or [Texture] of the `update` instance’s [Drawable] property.
// Before calling [present()], encode all your Metal commands to the `link`
// parameter’s [MTLDevice]. The GPU has additional time to complete running
// your commands before the frame displays on screen, determined by the value
// of the `link` parameter’s [PreferredFrameLatency] property.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMetalDisplayLinkDelegate/metalDisplayLink(_:needsUpdate:)
//
// [present()]: https://developer.apple.com/documentation/Metal/MTLDrawable/present()
// [MTLDevice]: https://developer.apple.com/documentation/Metal/MTLDevice
//
// [present()]: https://developer.apple.com/documentation/Metal/MTLDrawable/present()
func (o CAMetalDisplayLinkDelegateObject) MetalDisplayLinkNeedsUpdate(link ICAMetalDisplayLink, update ICAMetalDisplayLinkUpdate) {
	objc.Send[struct{}](o.ID, objc.Sel("metalDisplayLink:needsUpdate:"), link, update)
}

// CAMetalDisplayLinkDelegateConfig holds optional typed callbacks for [CAMetalDisplayLinkDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/quartzcore/cametaldisplaylinkdelegate
type CAMetalDisplayLinkDelegateConfig struct {

	// Receiving Display Updates
	// MetalDisplayLinkNeedsUpdate — A method the system calls to notify your app when it plans to update the display.
	MetalDisplayLinkNeedsUpdate func(link CAMetalDisplayLink, update CAMetalDisplayLinkUpdate)
}

// NewCAMetalDisplayLinkDelegate creates an Objective-C object implementing the [CAMetalDisplayLinkDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [CAMetalDisplayLinkDelegateObject] satisfies the [CAMetalDisplayLinkDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/quartzcore/cametaldisplaylinkdelegate
func NewCAMetalDisplayLinkDelegate(config CAMetalDisplayLinkDelegateConfig) CAMetalDisplayLinkDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoCAMetalDisplayLinkDelegate_%d", n)

	var methods []objc.MethodDef

	if config.MetalDisplayLinkNeedsUpdate != nil {
		fn := config.MetalDisplayLinkNeedsUpdate
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("metalDisplayLink:needsUpdate:"),
			Fn: func(self objc.ID, _cmd objc.SEL, linkID objc.ID, updateID objc.ID) {
				link := CAMetalDisplayLinkFromID(linkID)
				update := CAMetalDisplayLinkUpdateFromID(updateID)
				fn(link, update)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("CAMetalDisplayLinkDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewCAMetalDisplayLinkDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return CAMetalDisplayLinkDelegateObjectFromID(instance)
}
