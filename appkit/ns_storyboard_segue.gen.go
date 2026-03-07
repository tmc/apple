// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSStoryboardSegue] class.
var (
	_NSStoryboardSegueClass     NSStoryboardSegueClass
	_NSStoryboardSegueClassOnce sync.Once
)

func getNSStoryboardSegueClass() NSStoryboardSegueClass {
	_NSStoryboardSegueClassOnce.Do(func() {
		_NSStoryboardSegueClass = NSStoryboardSegueClass{class: objc.GetClass("NSStoryboardSegue")}
	})
	return _NSStoryboardSegueClass
}

// GetNSStoryboardSegueClass returns the class object for NSStoryboardSegue.
func GetNSStoryboardSegueClass() NSStoryboardSegueClass {
	return getNSStoryboardSegueClass()
}

type NSStoryboardSegueClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSStoryboardSegueClass) Alloc() NSStoryboardSegue {
	rv := objc.Send[NSStoryboardSegue](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A transition or containment relationship between two scenes in a
// storyboard.
//
// # Overview
// 
// In this context, a is a view controller or a window controller and a is an
// instance of the [NSStoryboard] class.
// 
// A storyboard segue has a procedural notion of being invoked, known in the
// API as being . You can take advantage of hooks into the segue performance
// process by way of the [NSSeguePerforming] protocol.
// 
// You do not create storyboard segue objects directly. Instead, the system
// creates them as needed as segues are invoked. To run code during
// initialization and performance of a segue, override the
// [NSStoryboardSegue.InitWithIdentifierSourceDestination] and [NSStoryboardSegue.Perform] methods.
// 
// You can initiate a segue programmatically with the
// [PerformSegueWithIdentifierSender] method of the [NSSeguePerforming]
// protocol. For example, you might do this to transition from a scene in one
// storyboard file to a scene in another.
//
// # Inspecting a Storyboard Segue
//
//   - [NSStoryboardSegue.SourceController]: The starting/containing view controller or window controller for the storyboard segue.
//   - [NSStoryboardSegue.DestinationController]: The ending/contained view controller or window controller for the storyboard segue.
//   - [NSStoryboardSegue.Identifier]: An optional, unique identifier for the storyboard segue that you can specify using the Identity inspector in Interface Builder.
//
// # Customizing Storyboard Segue Initialization and Invocation
//
//   - [NSStoryboardSegue.InitWithIdentifierSourceDestination]: The designated initializer for a storyboard segue.
//   - [NSStoryboardSegue.Perform]: Performs a visual transition from one controller to another.
//
// See: https://developer.apple.com/documentation/AppKit/NSStoryboardSegue
type NSStoryboardSegue struct {
	objectivec.Object
}

// NSStoryboardSegueFromID constructs a [NSStoryboardSegue] from an objc.ID.
//
// A transition or containment relationship between two scenes in a
// storyboard.
func NSStoryboardSegueFromID(id objc.ID) NSStoryboardSegue {
	return NSStoryboardSegue{objectivec.Object{id}}
}
// NOTE: NSStoryboardSegue adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSStoryboardSegue] class.
//
// # Inspecting a Storyboard Segue
//
//   - [INSStoryboardSegue.SourceController]: The starting/containing view controller or window controller for the storyboard segue.
//   - [INSStoryboardSegue.DestinationController]: The ending/contained view controller or window controller for the storyboard segue.
//   - [INSStoryboardSegue.Identifier]: An optional, unique identifier for the storyboard segue that you can specify using the Identity inspector in Interface Builder.
//
// # Customizing Storyboard Segue Initialization and Invocation
//
//   - [INSStoryboardSegue.InitWithIdentifierSourceDestination]: The designated initializer for a storyboard segue.
//   - [INSStoryboardSegue.Perform]: Performs a visual transition from one controller to another.
//
// See: https://developer.apple.com/documentation/AppKit/NSStoryboardSegue
type INSStoryboardSegue interface {
	objectivec.IObject

	// Topic: Inspecting a Storyboard Segue

	// The starting/containing view controller or window controller for the storyboard segue.
	SourceController() objectivec.IObject
	// The ending/contained view controller or window controller for the storyboard segue.
	DestinationController() objectivec.IObject
	// An optional, unique identifier for the storyboard segue that you can specify using the Identity inspector in Interface Builder.
	Identifier() NSStoryboardSegueIdentifier

	// Topic: Customizing Storyboard Segue Initialization and Invocation

	// The designated initializer for a storyboard segue.
	InitWithIdentifierSourceDestination(identifier NSStoryboardSegueIdentifier, sourceController objectivec.IObject, destinationController objectivec.IObject) NSStoryboardSegue
	// Performs a visual transition from one controller to another.
	Perform()
}





// Init initializes the instance.
func (s NSStoryboardSegue) Init() NSStoryboardSegue {
	rv := objc.Send[NSStoryboardSegue](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSStoryboardSegue) Autorelease() NSStoryboardSegue {
	rv := objc.Send[NSStoryboardSegue](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSStoryboardSegue creates a new NSStoryboardSegue instance.
func NewNSStoryboardSegue() NSStoryboardSegue {
	class := getNSStoryboardSegueClass()
	rv := objc.Send[NSStoryboardSegue](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// The designated initializer for a storyboard segue.
//
// identifier: The unique identifier for the storyboard segue. See the [Identifier]
// property.
//
// sourceController: The starting/containing view controller or window controller for the
// storyboard segue.
//
// destinationController: The ending/contained view controller or window controller for the
// storyboard segue.
//
// # Return Value
// 
// An initialized storyboard segue, ready to be performed.
//
// # Discussion
// 
// When a segue begins, the system calls this method. To run code during segue
// initialization, implement a storyboard segue subclass and override this
// method.
// 
// Whenever this method is called, the system then calls the [Perform] method.
//
// See: https://developer.apple.com/documentation/AppKit/NSStoryboardSegue/init(identifier:source:destination:)
func NewStoryboardSegueWithIdentifierSourceDestination(identifier NSStoryboardSegueIdentifier, sourceController objectivec.IObject, destinationController objectivec.IObject) NSStoryboardSegue {
	instance := getNSStoryboardSegueClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIdentifier:source:destination:"), objc.String(string(identifier)), sourceController, destinationController)
	return NSStoryboardSegueFromID(rv)
}







// The designated initializer for a storyboard segue.
//
// identifier: The unique identifier for the storyboard segue. See the [Identifier]
// property.
//
// sourceController: The starting/containing view controller or window controller for the
// storyboard segue.
//
// destinationController: The ending/contained view controller or window controller for the
// storyboard segue.
//
// # Return Value
// 
// An initialized storyboard segue, ready to be performed.
//
// # Discussion
// 
// When a segue begins, the system calls this method. To run code during segue
// initialization, implement a storyboard segue subclass and override this
// method.
// 
// Whenever this method is called, the system then calls the [Perform] method.
//
// See: https://developer.apple.com/documentation/AppKit/NSStoryboardSegue/init(identifier:source:destination:)
func (s NSStoryboardSegue) InitWithIdentifierSourceDestination(identifier NSStoryboardSegueIdentifier, sourceController objectivec.IObject, destinationController objectivec.IObject) NSStoryboardSegue {
	rv := objc.Send[NSStoryboardSegue](s.ID, objc.Sel("initWithIdentifier:source:destination:"), objc.String(string(identifier)), sourceController, destinationController)
	return rv
}

// Performs a visual transition from one controller to another.
//
// # Discussion
// 
// You can override this method in your [NSStoryboardSegue] subclass to
// perform custom animation between the starting/containing controller and the
// ending/contained controller for a storyboard segue. Typically, you would
// use Core Animation to set up an animation from one set of views to the
// next. For more complex animations, you might take a snapshot image of the
// two view hierarchies and manipulate the images instead of the view objects.
// 
// Regardless of how you perform the animation, you are responsible for
// installing the destination view controller o window controller (and its
// contained views) in the right place so that it can handle events.
// Typically, this entails calling one of the presentation methods in the
// [NSViewController] class.
//
// See: https://developer.apple.com/documentation/AppKit/NSStoryboardSegue/perform()
func (s NSStoryboardSegue) Perform() {
	objc.Send[objc.ID](s.ID, objc.Sel("perform"))
}





// Creates a storyboard segue and a block used when the segue is performed.
//
// identifier: The unique identifier for the storyboard segue. See the [Identifier]
// property.
//
// sourceController: The starting/containing view controller or window controller for the
// storyboard segue.
//
// destinationController: The ending/contained view controller or window controller for the
// storyboard segue.
//
// performHandler: A block of code that you provide, to be run each time the system calls the
// [Perform] method.
//
// # Return Value
// 
// An initialized storyboard segue and code block, ready to be performed.
//
// # Discussion
// 
// You can use this method to customize a storyboard segue in lieu of creating
// a subclass.
//
// See: https://developer.apple.com/documentation/AppKit/NSStoryboardSegue/init(identifier:source:destination:performHandler:)
func (_NSStoryboardSegueClass NSStoryboardSegueClass) SegueWithIdentifierSourceDestinationPerformHandler(identifier NSStoryboardSegueIdentifier, sourceController objectivec.IObject, destinationController objectivec.IObject, performHandler VoidHandler) NSStoryboardSegue {
		_block3, _cleanup3 := NewVoidBlock(performHandler)
	defer _cleanup3()
		rv := objc.Send[objc.ID](objc.ID(_NSStoryboardSegueClass.class), objc.Sel("segueWithIdentifier:source:destination:performHandler:"), identifier, sourceController, destinationController, _block3)
	return NSStoryboardSegueFromID(rv)
}








// The starting/containing view controller or window controller for the
// storyboard segue.
//
// # Discussion
// 
// In your storyboard segue subclass, you can read this property to get the
// starting/containing view controller or window controller for the segue.
//
// See: https://developer.apple.com/documentation/AppKit/NSStoryboardSegue/sourceController
func (s NSStoryboardSegue) SourceController() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("sourceController"))
	return objectivec.Object{ID: rv}
}



// The ending/contained view controller or window controller for the
// storyboard segue.
//
// # Discussion
// 
// In your storyboard segue subclass, you can read this property to get the
// ending/contained view controller or window controller for the segue. This
// property is essential if you override the [PrepareForSegueSender] method of
// the [NSSeguePerforming] protocol, to let you pass configuration data to the
// ending/contained controller.
//
// See: https://developer.apple.com/documentation/AppKit/NSStoryboardSegue/destinationController
func (s NSStoryboardSegue) DestinationController() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("destinationController"))
	return objectivec.Object{ID: rv}
}



// An optional, unique identifier for the storyboard segue that you can
// specify using the Identity inspector in Interface Builder.
//
// # Discussion
// 
// You use this property if you override the [PrepareForSegueSender] method of
// the [NSSeguePerforming] protocol.
//
// See: https://developer.apple.com/documentation/AppKit/NSStoryboardSegue/identifier-swift.property
func (s NSStoryboardSegue) Identifier() NSStoryboardSegueIdentifier {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("identifier"))
	return NSStoryboardSegueIdentifier(foundation.NSStringFromID(rv).String())
}


















// SegueWithIdentifierSourceDestinationPerformHandlerSync is a synchronous wrapper around [NSStoryboardSegue.SegueWithIdentifierSourceDestinationPerformHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (sc NSStoryboardSegueClass) SegueWithIdentifierSourceDestinationPerformHandlerSync(ctx context.Context, identifier NSStoryboardSegueIdentifier, sourceController objectivec.IObject, destinationController objectivec.IObject) error {
	done := make(chan struct{}, 1)
	sc.SegueWithIdentifierSourceDestinationPerformHandler(identifier, sourceController, destinationController, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}






