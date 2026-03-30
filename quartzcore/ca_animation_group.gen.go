// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CAAnimationGroup] class.
var (
	_CAAnimationGroupClass     CAAnimationGroupClass
	_CAAnimationGroupClassOnce sync.Once
)

func getCAAnimationGroupClass() CAAnimationGroupClass {
	_CAAnimationGroupClassOnce.Do(func() {
		_CAAnimationGroupClass = CAAnimationGroupClass{class: objc.GetClass("CAAnimationGroup")}
	})
	return _CAAnimationGroupClass
}

// GetCAAnimationGroupClass returns the class object for CAAnimationGroup.
func GetCAAnimationGroupClass() CAAnimationGroupClass {
	return getCAAnimationGroupClass()
}

type CAAnimationGroupClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CAAnimationGroupClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CAAnimationGroupClass) Alloc() CAAnimationGroup {
	rv := objc.Send[CAAnimationGroup](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that allows multiple animations to be grouped and run
// concurrently.
//
// # Overview
//
// The grouped animations run in the time space specified by the
// [CAAnimationGroup] instance.
//
// The duration of the grouped animations are not scaled to the duration of
// their [CAAnimationGroup]. Instead, the animations are clipped to the
// duration of the animation group. For example, a 10 second animation grouped
// within an animation group with a duration of 5 seconds displays only the
// first 5 seconds of the animation.
//
// The following code shows how you can create a grouped animation containing
// opacity and scale animations to fade out a layer while expanding it. The
// animation starts with an opacity of `1` and a scale of `1` on all axes. As
// the animation’s scale increases to `(3, 3, 3)`, the opacity drops to `0`
// and the animated layer vanishes.
//
// # Grouped animations
//
//   - [CAAnimationGroup.Animations]: An array of [CAAnimation] objects to be evaluated in the time space of the receiver.
//   - [CAAnimationGroup.SetAnimations]
//
// See: https://developer.apple.com/documentation/QuartzCore/CAAnimationGroup
type CAAnimationGroup struct {
	CAAnimation
}

// CAAnimationGroupFromID constructs a [CAAnimationGroup] from an objc.ID.
//
// An object that allows multiple animations to be grouped and run
// concurrently.
func CAAnimationGroupFromID(id objc.ID) CAAnimationGroup {
	return CAAnimationGroup{CAAnimation: CAAnimationFromID(id)}
}

// NOTE: CAAnimationGroup adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CAAnimationGroup] class.
//
// # Grouped animations
//
//   - [ICAAnimationGroup.Animations]: An array of [CAAnimation] objects to be evaluated in the time space of the receiver.
//   - [ICAAnimationGroup.SetAnimations]
//
// See: https://developer.apple.com/documentation/QuartzCore/CAAnimationGroup
type ICAAnimationGroup interface {
	ICAAnimation

	// Topic: Grouped animations

	// An array of [CAAnimation] objects to be evaluated in the time space of the receiver.
	Animations() []CAAnimation
	SetAnimations(value []CAAnimation)

	// Determines if the animation is removed from the target layer’s animations upon completion.
	IsRemovedOnCompletion() bool
	SetIsRemovedOnCompletion(value bool)
}

// Init initializes the instance.
func (a CAAnimationGroup) Init() CAAnimationGroup {
	rv := objc.Send[CAAnimationGroup](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a CAAnimationGroup) Autorelease() CAAnimationGroup {
	rv := objc.Send[CAAnimationGroup](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewCAAnimationGroup creates a new CAAnimationGroup instance.
func NewCAAnimationGroup() CAAnimationGroup {
	class := getCAAnimationGroupClass()
	rv := objc.Send[CAAnimationGroup](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// An array of [CAAnimation] objects to be evaluated in the time space of the
// receiver.
//
// # Discussion
//
// The animations run concurrently in the receiver’s time space.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAAnimationGroup/animations
func (a CAAnimationGroup) Animations() []CAAnimation {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("animations"))
	return objc.ConvertSlice(rv, func(id objc.ID) CAAnimation {
		return CAAnimationFromID(id)
	})
}
func (a CAAnimationGroup) SetAnimations(value []CAAnimation) {
	objc.Send[struct{}](a.ID, objc.Sel("setAnimations:"), objectivec.IObjectSliceToNSArray(value))
}

// Determines if the animation is removed from the target layer’s animations
// upon completion.
//
// See: https://developer.apple.com/documentation/quartzcore/caanimation/isremovedoncompletion
func (a CAAnimationGroup) IsRemovedOnCompletion() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("removedOnCompletion"))
	return rv
}
func (a CAAnimationGroup) SetIsRemovedOnCompletion(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setRemovedOnCompletion:"), value)
}
