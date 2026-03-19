// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSCollectionViewTransitionLayout] class.
var (
	_NSCollectionViewTransitionLayoutClass     NSCollectionViewTransitionLayoutClass
	_NSCollectionViewTransitionLayoutClassOnce sync.Once
)

func getNSCollectionViewTransitionLayoutClass() NSCollectionViewTransitionLayoutClass {
	_NSCollectionViewTransitionLayoutClassOnce.Do(func() {
		_NSCollectionViewTransitionLayoutClass = NSCollectionViewTransitionLayoutClass{class: objc.GetClass("NSCollectionViewTransitionLayout")}
	})
	return _NSCollectionViewTransitionLayoutClass
}

// GetNSCollectionViewTransitionLayoutClass returns the class object for NSCollectionViewTransitionLayout.
func GetNSCollectionViewTransitionLayoutClass() NSCollectionViewTransitionLayoutClass {
	return getNSCollectionViewTransitionLayoutClass()
}

type NSCollectionViewTransitionLayoutClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCollectionViewTransitionLayoutClass) Alloc() NSCollectionViewTransitionLayout {
	rv := objc.Send[NSCollectionViewTransitionLayout](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that implements custom behaviors when changing from one layout to
// another in a collection view.
//
// # Overview
// 
// Transition layout objects are commonly used to implement interactive
// transitions between layouts, where the transition itself is driven by a
// gesture recognizer.
//
// # Initializing the Transition Layout Object
//
//   - [NSCollectionViewTransitionLayout.InitWithCurrentLayoutNextLayout]: Initializes and returns the transition layout object.
//
// # Updating the Transition Information
//
//   - [NSCollectionViewTransitionLayout.TransitionProgress]: The completion percentage of the transition.
//   - [NSCollectionViewTransitionLayout.SetTransitionProgress]
//   - [NSCollectionViewTransitionLayout.UpdateValueForAnimatedKey]: Sets the value of a key whose value you use during the animation.
//   - [NSCollectionViewTransitionLayout.ValueForAnimatedKey]: Returns the most recently set value for the specified key.
//
// # Accessing the Layout Objects
//
//   - [NSCollectionViewTransitionLayout.CurrentLayout]: The collection view’s current layout object.
//   - [NSCollectionViewTransitionLayout.NextLayout]: The collection view’s new layout object.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewTransitionLayout
type NSCollectionViewTransitionLayout struct {
	NSCollectionViewLayout
}

// NSCollectionViewTransitionLayoutFromID constructs a [NSCollectionViewTransitionLayout] from an objc.ID.
//
// An object that implements custom behaviors when changing from one layout to
// another in a collection view.
func NSCollectionViewTransitionLayoutFromID(id objc.ID) NSCollectionViewTransitionLayout {
	return NSCollectionViewTransitionLayout{NSCollectionViewLayout: NSCollectionViewLayoutFromID(id)}
}
// NOTE: NSCollectionViewTransitionLayout adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSCollectionViewTransitionLayout] class.
//
// # Initializing the Transition Layout Object
//
//   - [INSCollectionViewTransitionLayout.InitWithCurrentLayoutNextLayout]: Initializes and returns the transition layout object.
//
// # Updating the Transition Information
//
//   - [INSCollectionViewTransitionLayout.TransitionProgress]: The completion percentage of the transition.
//   - [INSCollectionViewTransitionLayout.SetTransitionProgress]
//   - [INSCollectionViewTransitionLayout.UpdateValueForAnimatedKey]: Sets the value of a key whose value you use during the animation.
//   - [INSCollectionViewTransitionLayout.ValueForAnimatedKey]: Returns the most recently set value for the specified key.
//
// # Accessing the Layout Objects
//
//   - [INSCollectionViewTransitionLayout.CurrentLayout]: The collection view’s current layout object.
//   - [INSCollectionViewTransitionLayout.NextLayout]: The collection view’s new layout object.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewTransitionLayout
type INSCollectionViewTransitionLayout interface {
	INSCollectionViewLayout

	// Topic: Initializing the Transition Layout Object

	// Initializes and returns the transition layout object.
	InitWithCurrentLayoutNextLayout(currentLayout INSCollectionViewLayout, newLayout INSCollectionViewLayout) NSCollectionViewTransitionLayout

	// Topic: Updating the Transition Information

	// The completion percentage of the transition.
	TransitionProgress() float64
	SetTransitionProgress(value float64)
	// Sets the value of a key whose value you use during the animation.
	UpdateValueForAnimatedKey(value float64, key NSCollectionViewTransitionLayoutAnimatedKey)
	// Returns the most recently set value for the specified key.
	ValueForAnimatedKey(key NSCollectionViewTransitionLayoutAnimatedKey) float64

	// Topic: Accessing the Layout Objects

	// The collection view’s current layout object.
	CurrentLayout() INSCollectionViewLayout
	// The collection view’s new layout object.
	NextLayout() INSCollectionViewLayout
}

// Init initializes the instance.
func (c NSCollectionViewTransitionLayout) Init() NSCollectionViewTransitionLayout {
	rv := objc.Send[NSCollectionViewTransitionLayout](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCollectionViewTransitionLayout) Autorelease() NSCollectionViewTransitionLayout {
	rv := objc.Send[NSCollectionViewTransitionLayout](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCollectionViewTransitionLayout creates a new NSCollectionViewTransitionLayout instance.
func NewNSCollectionViewTransitionLayout() NSCollectionViewTransitionLayout {
	class := getNSCollectionViewTransitionLayoutClass()
	rv := objc.Send[NSCollectionViewTransitionLayout](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes and returns the transition layout object.
//
// currentLayout: The layout object currently in use by the collection view.
//
// newLayout: The new layout object that is about to be installed into the collection
// view.
//
// # Return Value
// 
// An initialized transition layout or `nil` if the object could not be
// initialized.
//
// # Discussion
// 
// This method initializes the transition layout object and saves references
// to the current and new layout objects. If you subclass and implement your
// own initialization method, you must call this method to initialize the
// superclass.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewTransitionLayout/init(currentLayout:nextLayout:)
func NewCollectionViewTransitionLayoutWithCurrentLayoutNextLayout(currentLayout INSCollectionViewLayout, newLayout INSCollectionViewLayout) NSCollectionViewTransitionLayout {
	instance := getNSCollectionViewTransitionLayoutClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCurrentLayout:nextLayout:"), currentLayout, newLayout)
	return NSCollectionViewTransitionLayoutFromID(rv)
}

// Initializes and returns the transition layout object.
//
// currentLayout: The layout object currently in use by the collection view.
//
// newLayout: The new layout object that is about to be installed into the collection
// view.
//
// # Return Value
// 
// An initialized transition layout or `nil` if the object could not be
// initialized.
//
// # Discussion
// 
// This method initializes the transition layout object and saves references
// to the current and new layout objects. If you subclass and implement your
// own initialization method, you must call this method to initialize the
// superclass.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewTransitionLayout/init(currentLayout:nextLayout:)
func (c NSCollectionViewTransitionLayout) InitWithCurrentLayoutNextLayout(currentLayout INSCollectionViewLayout, newLayout INSCollectionViewLayout) NSCollectionViewTransitionLayout {
	rv := objc.Send[NSCollectionViewTransitionLayout](c.ID, objc.Sel("initWithCurrentLayout:nextLayout:"), currentLayout, newLayout)
	return rv
}
// Sets the value of a key whose value you use during the animation.
//
// value: The value of the key.
//
// key: The key that you define for your custom transition layout.
//
// # Discussion
// 
// Use this method to update the value of a specific key that you use in your
// custom transition layout.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewTransitionLayout/updateValue(_:forAnimatedKey:)
func (c NSCollectionViewTransitionLayout) UpdateValueForAnimatedKey(value float64, key NSCollectionViewTransitionLayoutAnimatedKey) {
	objc.Send[objc.ID](c.ID, objc.Sel("updateValue:forAnimatedKey:"), value, objc.String(string(key)))
}
// Returns the most recently set value for the specified key.
//
// key: A key whose value you set previously using the [UpdateValueForAnimatedKey]
// method.
//
// # Return Value
// 
// The last value set for the key.
//
// # Discussion
// 
// Use this method to retrieve floating-point values that relate to laying out
// the contents of your collection view. The key you specify is a string that
// you define and that has some meaning to your layout’s implementation. At
// points during an interactive transition, you can assign new values to that
// key using the [UpdateValueForAnimatedKey] method.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewTransitionLayout/value(forAnimatedKey:)
func (c NSCollectionViewTransitionLayout) ValueForAnimatedKey(key NSCollectionViewTransitionLayoutAnimatedKey) float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("valueForAnimatedKey:"), objc.String(string(key)))
	return rv
}

// The completion percentage of the transition.
//
// # Discussion
// 
// During the transition, set the value of this property periodically and call
// the [InvalidateLayout] method to force the collection view to update item
// positions. For example, when driving a transition using a gesture
// recognizer, you can set this property from the handler method of your
// gesture recognizer.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewTransitionLayout/transitionProgress
func (c NSCollectionViewTransitionLayout) TransitionProgress() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("transitionProgress"))
	return rv
}
func (c NSCollectionViewTransitionLayout) SetTransitionProgress(value float64) {
	objc.Send[struct{}](c.ID, objc.Sel("setTransitionProgress:"), value)
}
// The collection view’s current layout object.
//
// # Discussion
// 
// Use this object to retrieve the initial layout attributes for elements of
// the collection view. If the transition is ultimately cancelled, the
// collection view animates its items back to the attributes provided by this
// object.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewTransitionLayout/currentLayout
func (c NSCollectionViewTransitionLayout) CurrentLayout() INSCollectionViewLayout {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("currentLayout"))
	return NSCollectionViewLayoutFromID(objc.ID(rv))
}
// The collection view’s new layout object.
//
// # Discussion
// 
// Use this object to retrieve the final layout attributes for elements of the
// collection view. If the transition completes as expected, the collection
// view animates its items to the attributes provided by this object.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewTransitionLayout/nextLayout
func (c NSCollectionViewTransitionLayout) NextLayout() INSCollectionViewLayout {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("nextLayout"))
	return NSCollectionViewLayoutFromID(objc.ID(rv))
}

