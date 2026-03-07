// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSSliderAccessoryBehavior] class.
var (
	_NSSliderAccessoryBehaviorClass     NSSliderAccessoryBehaviorClass
	_NSSliderAccessoryBehaviorClassOnce sync.Once
)

func getNSSliderAccessoryBehaviorClass() NSSliderAccessoryBehaviorClass {
	_NSSliderAccessoryBehaviorClassOnce.Do(func() {
		_NSSliderAccessoryBehaviorClass = NSSliderAccessoryBehaviorClass{class: objc.GetClass("NSSliderAccessoryBehavior")}
	})
	return _NSSliderAccessoryBehaviorClass
}

// GetNSSliderAccessoryBehaviorClass returns the class object for NSSliderAccessoryBehavior.
func GetNSSliderAccessoryBehaviorClass() NSSliderAccessoryBehaviorClass {
	return getNSSliderAccessoryBehaviorClass()
}

type NSSliderAccessoryBehaviorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSliderAccessoryBehaviorClass) Alloc() NSSliderAccessoryBehavior {
	rv := objc.Send[NSSliderAccessoryBehavior](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







//
// # Instance Methods
//
//   - [NSSliderAccessoryBehavior.HandleAction]
// See: https://developer.apple.com/documentation/AppKit/NSSliderAccessoryBehavior
type NSSliderAccessoryBehavior struct {
	objectivec.Object
}

// NSSliderAccessoryBehaviorFromID constructs a [NSSliderAccessoryBehavior] from an objc.ID.
func NSSliderAccessoryBehaviorFromID(id objc.ID) NSSliderAccessoryBehavior {
	return NSSliderAccessoryBehavior{objectivec.Object{id}}
}
// NOTE: NSSliderAccessoryBehavior adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSSliderAccessoryBehavior] class.
//
// # Instance Methods
//
//   - [INSSliderAccessoryBehavior.HandleAction]
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderAccessoryBehavior
type INSSliderAccessoryBehavior interface {
	objectivec.IObject

	// Topic: Instance Methods

	HandleAction(sender INSSliderAccessory)

	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (s NSSliderAccessoryBehavior) Init() NSSliderAccessoryBehavior {
	rv := objc.Send[NSSliderAccessoryBehavior](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSliderAccessoryBehavior) Autorelease() NSSliderAccessoryBehavior {
	rv := objc.Send[NSSliderAccessoryBehavior](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSliderAccessoryBehavior creates a new NSSliderAccessoryBehavior instance.
func NewNSSliderAccessoryBehavior() NSSliderAccessoryBehavior {
	class := getNSSliderAccessoryBehaviorClass()
	rv := objc.Send[NSSliderAccessoryBehavior](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/AppKit/NSSliderAccessoryBehavior/init(target:action:)
func NewSliderAccessoryBehaviorWithTargetAction(target objectivec.IObject, action objc.SEL) NSSliderAccessoryBehavior {
	rv := objc.Send[objc.ID](objc.ID(getNSSliderAccessoryBehaviorClass().class), objc.Sel("behaviorWithTarget:action:"), target, action)
	return NSSliderAccessoryBehaviorFromID(rv)
}







//
// See: https://developer.apple.com/documentation/AppKit/NSSliderAccessoryBehavior/handleAction(_:)
func (s NSSliderAccessoryBehavior) HandleAction(sender INSSliderAccessory) {
	objc.Send[objc.ID](s.ID, objc.Sel("handleAction:"), sender)
}
func (s NSSliderAccessoryBehavior) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}





//
// See: https://developer.apple.com/documentation/AppKit/NSSliderAccessoryBehavior/init(handler:)
func (_NSSliderAccessoryBehaviorClass NSSliderAccessoryBehaviorClass) BehaviorWithHandler(handler SliderAccessoryHandler) NSSliderAccessoryBehavior {
		_block0, _cleanup0 := NewSliderAccessoryBlock(handler)
	defer _cleanup0()
		rv := objc.Send[objc.ID](objc.ID(_NSSliderAccessoryBehaviorClass.class), objc.Sel("behaviorWithHandler:"), _block0)
	return NSSliderAccessoryBehaviorFromID(rv)
}












// See: https://developer.apple.com/documentation/AppKit/NSSliderAccessoryBehavior/automatic
func (_NSSliderAccessoryBehaviorClass NSSliderAccessoryBehaviorClass) AutomaticBehavior() NSSliderAccessoryBehavior {
	rv := objc.Send[objc.ID](objc.ID(_NSSliderAccessoryBehaviorClass.class), objc.Sel("automaticBehavior"))
	return NSSliderAccessoryBehaviorFromID(objc.ID(rv))
}



// See: https://developer.apple.com/documentation/AppKit/NSSliderAccessoryBehavior/valueReset
func (_NSSliderAccessoryBehaviorClass NSSliderAccessoryBehaviorClass) ValueResetBehavior() NSSliderAccessoryBehavior {
	rv := objc.Send[objc.ID](objc.ID(_NSSliderAccessoryBehaviorClass.class), objc.Sel("valueResetBehavior"))
	return NSSliderAccessoryBehaviorFromID(objc.ID(rv))
}



// See: https://developer.apple.com/documentation/AppKit/NSSliderAccessoryBehavior/valueStep
func (_NSSliderAccessoryBehaviorClass NSSliderAccessoryBehaviorClass) ValueStepBehavior() NSSliderAccessoryBehavior {
	rv := objc.Send[objc.ID](objc.ID(_NSSliderAccessoryBehaviorClass.class), objc.Sel("valueStepBehavior"))
	return NSSliderAccessoryBehaviorFromID(objc.ID(rv))
}

















// BehaviorWithHandlerSync is a synchronous wrapper around [NSSliderAccessoryBehavior.BehaviorWithHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (sc NSSliderAccessoryBehaviorClass) BehaviorWithHandlerSync(ctx context.Context) (*NSSliderAccessory, error) {
	done := make(chan *NSSliderAccessory, 1)
	sc.BehaviorWithHandler(func(val *NSSliderAccessory) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}






