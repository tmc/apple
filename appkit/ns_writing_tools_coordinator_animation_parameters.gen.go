// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSWritingToolsCoordinatorAnimationParameters] class.
var (
	_NSWritingToolsCoordinatorAnimationParametersClass     NSWritingToolsCoordinatorAnimationParametersClass
	_NSWritingToolsCoordinatorAnimationParametersClassOnce sync.Once
)

func getNSWritingToolsCoordinatorAnimationParametersClass() NSWritingToolsCoordinatorAnimationParametersClass {
	_NSWritingToolsCoordinatorAnimationParametersClassOnce.Do(func() {
		_NSWritingToolsCoordinatorAnimationParametersClass = NSWritingToolsCoordinatorAnimationParametersClass{class: objc.GetClass("NSWritingToolsCoordinatorAnimationParameters")}
	})
	return _NSWritingToolsCoordinatorAnimationParametersClass
}

// GetNSWritingToolsCoordinatorAnimationParametersClass returns the class object for NSWritingToolsCoordinatorAnimationParameters.
func GetNSWritingToolsCoordinatorAnimationParametersClass() NSWritingToolsCoordinatorAnimationParametersClass {
	return getNSWritingToolsCoordinatorAnimationParametersClass()
}

type NSWritingToolsCoordinatorAnimationParametersClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSWritingToolsCoordinatorAnimationParametersClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSWritingToolsCoordinatorAnimationParametersClass) Alloc() NSWritingToolsCoordinatorAnimationParameters {
	rv := objc.Send[NSWritingToolsCoordinatorAnimationParameters](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object you use to configure additional tasks or animations to run
// alongside the Writing Tools animations.
//
// # Overview
//
// When Writing Tools replaces text in one of your context objects, it
// provides an `NSWritingToolsCoordinator.AnimationParameters` object for you
// to use to configure any additional animations. During a Writing Tools
// session, you hide the text under evaluation and provide a targeted preview
// of your content. Writing Tools animations changes to that preview, but you
// might need to provide additional animations for other parts of your
// view’s content. For example, you might need to animate any layout changes
// caused by the insertion or removal of text in other parts of your view. Use
// this object to configure those animations.
//
// You don’t create an `NSWritingToolsCoordinator.AnimationParameters`
// object directly. Instead, the system creates one and passes it to the
// [WritingToolsCoordinatorReplaceRangeInContextProposedTextReasonAnimationParametersCompletion]
// method of your [NSWritingToolsCoordinatorDelegate] object. Use that object
// to specify the blocks to run during and after the system animations.
//
// # Getting the animation values
//
//   - [NSWritingToolsCoordinatorAnimationParameters.Duration]: The number of seconds it takes the system animations to run.
//   - [NSWritingToolsCoordinatorAnimationParameters.Delay]: The number of seconds the system waits before starting its animations.
//
// # Creating custom animations
//
//   - [NSWritingToolsCoordinatorAnimationParameters.ProgressHandler]: A custom block that runs at the same time as the system animations.
//   - [NSWritingToolsCoordinatorAnimationParameters.SetProgressHandler]
//   - [NSWritingToolsCoordinatorAnimationParameters.CompletionHandler]: A custom block to run when the system animations finish.
//   - [NSWritingToolsCoordinatorAnimationParameters.SetCompletionHandler]
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/AnimationParameters
type NSWritingToolsCoordinatorAnimationParameters struct {
	objectivec.Object
}

// NSWritingToolsCoordinatorAnimationParametersFromID constructs a [NSWritingToolsCoordinatorAnimationParameters] from an objc.ID.
//
// An object you use to configure additional tasks or animations to run
// alongside the Writing Tools animations.
func NSWritingToolsCoordinatorAnimationParametersFromID(id objc.ID) NSWritingToolsCoordinatorAnimationParameters {
	return NSWritingToolsCoordinatorAnimationParameters{objectivec.Object{ID: id}}
}

// NOTE: NSWritingToolsCoordinatorAnimationParameters adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSWritingToolsCoordinatorAnimationParameters] class.
//
// # Getting the animation values
//
//   - [INSWritingToolsCoordinatorAnimationParameters.Duration]: The number of seconds it takes the system animations to run.
//   - [INSWritingToolsCoordinatorAnimationParameters.Delay]: The number of seconds the system waits before starting its animations.
//
// # Creating custom animations
//
//   - [INSWritingToolsCoordinatorAnimationParameters.ProgressHandler]: A custom block that runs at the same time as the system animations.
//   - [INSWritingToolsCoordinatorAnimationParameters.SetProgressHandler]
//   - [INSWritingToolsCoordinatorAnimationParameters.CompletionHandler]: A custom block to run when the system animations finish.
//   - [INSWritingToolsCoordinatorAnimationParameters.SetCompletionHandler]
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/AnimationParameters
type INSWritingToolsCoordinatorAnimationParameters interface {
	objectivec.IObject

	// Topic: Getting the animation values

	// The number of seconds it takes the system animations to run.
	Duration() float64
	// The number of seconds the system waits before starting its animations.
	Delay() float64

	// Topic: Creating custom animations

	// A custom block that runs at the same time as the system animations.
	ProgressHandler() Float32Handler
	SetProgressHandler(value Float32Handler)
	// A custom block to run when the system animations finish.
	CompletionHandler() VoidHandler
	SetCompletionHandler(value VoidHandler)
}

// Init initializes the instance.
func (w NSWritingToolsCoordinatorAnimationParameters) Init() NSWritingToolsCoordinatorAnimationParameters {
	rv := objc.Send[NSWritingToolsCoordinatorAnimationParameters](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w NSWritingToolsCoordinatorAnimationParameters) Autorelease() NSWritingToolsCoordinatorAnimationParameters {
	rv := objc.Send[NSWritingToolsCoordinatorAnimationParameters](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSWritingToolsCoordinatorAnimationParameters creates a new NSWritingToolsCoordinatorAnimationParameters instance.
func NewNSWritingToolsCoordinatorAnimationParameters() NSWritingToolsCoordinatorAnimationParameters {
	class := getNSWritingToolsCoordinatorAnimationParametersClass()
	rv := objc.Send[NSWritingToolsCoordinatorAnimationParameters](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The number of seconds it takes the system animations to run.
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/AnimationParameters/duration
func (w NSWritingToolsCoordinatorAnimationParameters) Duration() float64 {
	rv := objc.Send[float64](w.ID, objc.Sel("duration"))
	return rv
}

// The number of seconds the system waits before starting its animations.
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/AnimationParameters/delay
func (w NSWritingToolsCoordinatorAnimationParameters) Delay() float64 {
	rv := objc.Send[float64](w.ID, objc.Sel("delay"))
	return rv
}

// A custom block that runs at the same time as the system animations.
//
// # Discussion
//
// If you have animations you want to run at the same time as the system
// animations, assign a block to this property and use it to run your
// animations. The block you provide must have no return value and take a
// floating-point value as a parameter. The parameter indicates the current
// progress of the animations as a percentage value between `0.0` to `1.0`.
// The system executes your block multiple times during the course of the
// animations, providing an updated completion value each time.
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/AnimationParameters/progressHandler
func (w NSWritingToolsCoordinatorAnimationParameters) ProgressHandler() Float32Handler {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("progressHandler"))
	_ = rv
	return nil
}
func (w NSWritingToolsCoordinatorAnimationParameters) SetProgressHandler(value Float32Handler) {
	block, cleanup := NewFloat32Block(value)
	defer cleanup()
	objc.Send[struct{}](w.ID, objc.Sel("setProgressHandler:"), block)
}

// A custom block to run when the system animations finish.
//
// # Discussion
//
// Set this property to a block that you want the system to run when any
// animations finish. The block you provide must have no return value and no
// parameters. The system executes this block once when the current animation
// finish.
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/AnimationParameters/completionHandler
func (w NSWritingToolsCoordinatorAnimationParameters) CompletionHandler() VoidHandler {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("completionHandler"))
	_ = rv
	return nil
}
func (w NSWritingToolsCoordinatorAnimationParameters) SetCompletionHandler(value VoidHandler) {
	block, cleanup := NewVoidBlock(value)
	defer cleanup()
	objc.Send[struct{}](w.ID, objc.Sel("setCompletionHandler:"), block)
}
