// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSStepperTouchBarItem] class.
var (
	_NSStepperTouchBarItemClass     NSStepperTouchBarItemClass
	_NSStepperTouchBarItemClassOnce sync.Once
)

func getNSStepperTouchBarItemClass() NSStepperTouchBarItemClass {
	_NSStepperTouchBarItemClassOnce.Do(func() {
		_NSStepperTouchBarItemClass = NSStepperTouchBarItemClass{class: objc.GetClass("NSStepperTouchBarItem")}
	})
	return _NSStepperTouchBarItemClass
}

// GetNSStepperTouchBarItemClass returns the class object for NSStepperTouchBarItem.
func GetNSStepperTouchBarItemClass() NSStepperTouchBarItemClass {
	return getNSStepperTouchBarItemClass()
}

type NSStepperTouchBarItemClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSStepperTouchBarItemClass) Alloc() NSStepperTouchBarItem {
	rv := objc.Send[NSStepperTouchBarItem](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A bar item that provides a stepper control for incrementing or decrementing
// a value.
//
// # Handling stepper interaction
//
//   - [NSStepperTouchBarItem.Target]
//   - [NSStepperTouchBarItem.SetTarget]
//   - [NSStepperTouchBarItem.Action]
//   - [NSStepperTouchBarItem.SetAction]
//
// # Managing the stepper’s value
//
//   - [NSStepperTouchBarItem.Value]
//   - [NSStepperTouchBarItem.SetValue]
//   - [NSStepperTouchBarItem.MaxValue]
//   - [NSStepperTouchBarItem.SetMaxValue]
//   - [NSStepperTouchBarItem.MinValue]
//   - [NSStepperTouchBarItem.SetMinValue]
//   - [NSStepperTouchBarItem.Increment]
//   - [NSStepperTouchBarItem.SetIncrement]
//
// See: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem
type NSStepperTouchBarItem struct {
	NSTouchBarItem
}

// NSStepperTouchBarItemFromID constructs a [NSStepperTouchBarItem] from an objc.ID.
//
// A bar item that provides a stepper control for incrementing or decrementing
// a value.
func NSStepperTouchBarItemFromID(id objc.ID) NSStepperTouchBarItem {
	return NSStepperTouchBarItem{NSTouchBarItem: NSTouchBarItemFromID(id)}
}
// NOTE: NSStepperTouchBarItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSStepperTouchBarItem] class.
//
// # Handling stepper interaction
//
//   - [INSStepperTouchBarItem.Target]
//   - [INSStepperTouchBarItem.SetTarget]
//   - [INSStepperTouchBarItem.Action]
//   - [INSStepperTouchBarItem.SetAction]
//
// # Managing the stepper’s value
//
//   - [INSStepperTouchBarItem.Value]
//   - [INSStepperTouchBarItem.SetValue]
//   - [INSStepperTouchBarItem.MaxValue]
//   - [INSStepperTouchBarItem.SetMaxValue]
//   - [INSStepperTouchBarItem.MinValue]
//   - [INSStepperTouchBarItem.SetMinValue]
//   - [INSStepperTouchBarItem.Increment]
//   - [INSStepperTouchBarItem.SetIncrement]
//
// See: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem
type INSStepperTouchBarItem interface {
	INSTouchBarItem

	// Topic: Handling stepper interaction

	Target() objectivec.IObject
	SetTarget(value objectivec.IObject)
	Action() objc.SEL
	SetAction(value objc.SEL)

	// Topic: Managing the stepper’s value

	Value() float64
	SetValue(value float64)
	MaxValue() float64
	SetMaxValue(value float64)
	MinValue() float64
	SetMinValue(value float64)
	Increment() float64
	SetIncrement(value float64)

	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (s NSStepperTouchBarItem) Init() NSStepperTouchBarItem {
	rv := objc.Send[NSStepperTouchBarItem](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSStepperTouchBarItem) Autorelease() NSStepperTouchBarItem {
	rv := objc.Send[NSStepperTouchBarItem](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSStepperTouchBarItem creates a new NSStepperTouchBarItem instance.
func NewNSStepperTouchBarItem() NSStepperTouchBarItem {
	class := getNSStepperTouchBarItemClass()
	rv := objc.Send[NSStepperTouchBarItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Initializes and returns a new item from a storyboard or nib file.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/init(coder:)
func NewStepperTouchBarItemWithCoder(coder foundation.INSCoder) NSStepperTouchBarItem {
	instance := getNSStepperTouchBarItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSStepperTouchBarItemFromID(rv)
}


// Creates a new item with the specified identifier.
//
// # Discussion
// 
// The designated initializer. The identifier must be globally unique for
// every item, except for space items.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/init(identifier:)
func NewStepperTouchBarItemWithIdentifier(identifier NSTouchBarItemIdentifier) NSStepperTouchBarItem {
	instance := getNSStepperTouchBarItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIdentifier:"), objc.String(string(identifier)))
	return NSStepperTouchBarItemFromID(rv)
}


//
// See: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/init(identifier:formatter:)
func NewStepperTouchBarItemWithIdentifierFormatter(identifier NSTouchBarItemIdentifier, formatter *foundation.NSFormatter) NSStepperTouchBarItem {
	rv := objc.Send[objc.ID](objc.ID(getNSStepperTouchBarItemClass().class), objc.Sel("stepperTouchBarItemWithIdentifier:formatter:"), objc.String(string(identifier)), formatter)
	return NSStepperTouchBarItemFromID(rv)
}






func (s NSStepperTouchBarItem) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}





//
// See: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/init(identifier:drawingHandler:)
func (_NSStepperTouchBarItemClass NSStepperTouchBarItemClass) StepperTouchBarItemWithIdentifierDrawingHandler(identifier NSTouchBarItemIdentifier, drawingHandler RectHandler) NSStepperTouchBarItem {
		_block1, _cleanup1 := NewRectBlock(drawingHandler)
	defer _cleanup1()
		rv := objc.Send[objc.ID](objc.ID(_NSStepperTouchBarItemClass.class), objc.Sel("stepperTouchBarItemWithIdentifier:drawingHandler:"), identifier, _block1)
	return NSStepperTouchBarItemFromID(rv)
}








// See: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/target
func (s NSStepperTouchBarItem) Target() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("target"))
	return objectivec.Object{ID: rv}
}
func (s NSStepperTouchBarItem) SetTarget(value objectivec.IObject) {
	objc.Send[struct{}](s.ID, objc.Sel("setTarget:"), value)
}



// See: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/action
func (s NSStepperTouchBarItem) Action() objc.SEL {
	rv := objc.Send[objc.SEL](s.ID, objc.Sel("action"))
	return rv
}
func (s NSStepperTouchBarItem) SetAction(value objc.SEL) {
	objc.Send[struct{}](s.ID, objc.Sel("setAction:"), value)
}



// See: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/value
func (s NSStepperTouchBarItem) Value() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("value"))
	return rv
}
func (s NSStepperTouchBarItem) SetValue(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setValue:"), value)
}



// See: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/maxValue
func (s NSStepperTouchBarItem) MaxValue() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("maxValue"))
	return rv
}
func (s NSStepperTouchBarItem) SetMaxValue(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setMaxValue:"), value)
}



// See: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/minValue
func (s NSStepperTouchBarItem) MinValue() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("minValue"))
	return rv
}
func (s NSStepperTouchBarItem) SetMinValue(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setMinValue:"), value)
}



// See: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/increment
func (s NSStepperTouchBarItem) Increment() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("increment"))
	return rv
}
func (s NSStepperTouchBarItem) SetIncrement(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setIncrement:"), value)
}





















// StepperTouchBarItemWithIdentifierDrawingHandlerSync is a synchronous wrapper around [NSStepperTouchBarItem.StepperTouchBarItemWithIdentifierDrawingHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (sc NSStepperTouchBarItemClass) StepperTouchBarItemWithIdentifierDrawingHandlerSync(ctx context.Context, identifier NSTouchBarItemIdentifier) (corefoundation.CGRect, error) {
	done := make(chan corefoundation.CGRect, 1)
	sc.StepperTouchBarItemWithIdentifierDrawingHandler(identifier, func(val corefoundation.CGRect) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return corefoundation.CGRect{}, ctx.Err()
	}
}






