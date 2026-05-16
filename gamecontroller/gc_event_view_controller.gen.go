// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/objc"
)

// The class instance for the [GCEventViewController] class.
var (
	_GCEventViewControllerClass     GCEventViewControllerClass
	_GCEventViewControllerClassOnce sync.Once
)

func getGCEventViewControllerClass() GCEventViewControllerClass {
	_GCEventViewControllerClassOnce.Do(func() {
		_GCEventViewControllerClass = GCEventViewControllerClass{class: objc.GetClass("GCEventViewController")}
	})
	return _GCEventViewControllerClass
}

// GetGCEventViewControllerClass returns the class object for GCEventViewController.
func GetGCEventViewControllerClass() GCEventViewControllerClass {
	return getGCEventViewControllerClass()
}

type GCEventViewControllerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GCEventViewControllerClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GCEventViewControllerClass) Alloc() GCEventViewController {
	rv := objc.Send[GCEventViewController](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// A view controller that delivers input either from the responder chain to
// views, or from game controllers to profiles.
//
// # Overview
//
// On systems, such as tvOS, where the player uses the game controller to both
// navigate the system interface and play your game, use a
// [GCEventViewController] object as the root view controller to selectively
// receive input directly from the game controller. You can’t simultaneously
// process input through the responder chain and Game Controller input
// elements.
//
// By default the system delivers input events to your app using the responder
// chain. To get the input values through the game controller objects, set a
// [GCEventViewController] object as the root view controller. The view
// controller delivers the input for its views and their subviews to the game
// controller’s profile. To switch back to the responder chain, set the view
// controller’s [GCEventViewController.ControllerUserInteractionEnabled] property to true.
//
// # Delivering game controller inputs
//
//   - [GCEventViewController.ControllerUserInteractionEnabled]: A Boolean value that indicates whether the system delivers game controller input to profile objects or to views using the responder chain.
//   - [GCEventViewController.SetControllerUserInteractionEnabled]
//
// See: https://developer.apple.com/documentation/GameController/GCEventViewController
type GCEventViewController struct {
	appkit.NSViewController
}

// GCEventViewControllerFromID constructs a [GCEventViewController] from an objc.ID.
//
// A view controller that delivers input either from the responder chain to
// views, or from game controllers to profiles.
func GCEventViewControllerFromID(id objc.ID) GCEventViewController {
	return GCEventViewController{NSViewController: appkit.NSViewControllerFromID(id)}
}

// NOTE: GCEventViewController adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [GCEventViewController] class.
//
// # Delivering game controller inputs
//
//   - [IGCEventViewController.ControllerUserInteractionEnabled]: A Boolean value that indicates whether the system delivers game controller input to profile objects or to views using the responder chain.
//   - [IGCEventViewController.SetControllerUserInteractionEnabled]
//
// See: https://developer.apple.com/documentation/GameController/GCEventViewController
type IGCEventViewController interface {
	appkit.INSViewController

	// Topic: Delivering game controller inputs

	// A Boolean value that indicates whether the system delivers game controller input to profile objects or to views using the responder chain.
	ControllerUserInteractionEnabled() bool
	SetControllerUserInteractionEnabled(value bool)
}

// Init initializes the instance.
func (g GCEventViewController) Init() GCEventViewController {
	rv := objc.Send[GCEventViewController](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GCEventViewController) Autorelease() GCEventViewController {
	rv := objc.Send[GCEventViewController](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCEventViewController creates a new GCEventViewController instance.
func NewGCEventViewController() GCEventViewController {
	class := getGCEventViewControllerClass()
	rv := objc.Send[GCEventViewController](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Boolean value that indicates whether the system delivers game controller
// input to profile objects or to views using the responder chain.
//
// # Discussion
//
// If this property is false, when the view controller’s view or its
// subviews are the first responder, the system delivers the game controller
// input to the profile objects. If this property is true, the system
// generates input events and delivers them through the responder chain.
//
// See: https://developer.apple.com/documentation/GameController/GCEventViewController/controllerUserInteractionEnabled
func (g GCEventViewController) ControllerUserInteractionEnabled() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("controllerUserInteractionEnabled"))
	return rv
}
func (g GCEventViewController) SetControllerUserInteractionEnabled(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setControllerUserInteractionEnabled:"), value)
}
