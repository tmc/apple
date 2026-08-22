// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A displayable resource that can be rendered or written to.
//
// See: https://developer.apple.com/documentation/Metal/MTLDrawable
type MTLDrawable interface {
	objectivec.IObject

	// Presents the drawable onscreen as soon as possible.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDrawable/present()
	Present()

	// Presents the drawable onscreen as soon as possible after a previous drawable is visible for the specified duration.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDrawable/present(afterMinimumDuration:)
	PresentAfterMinimumDuration(duration corefoundation.CFTimeInterval)

	// Presents the drawable onscreen at a specific host time.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDrawable/present(at:)
	PresentAtTime(presentationTime corefoundation.CFTimeInterval)

	// Registers a block of code to be called immediately after the drawable is presented.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDrawable/addPresentedHandler(_:)
	AddPresentedHandler(block MTLDrawableHandler)

	// A positive integer that identifies the drawable.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDrawable/drawableID
	DrawableID() uint

	// The host time, in seconds, when the drawable was displayed onscreen.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDrawable/presentedTime
	PresentedTime() corefoundation.CFTimeInterval
}

// MTLDrawableObject wraps an existing Objective-C object that conforms to the MTLDrawable protocol.
type MTLDrawableObject struct {
	objectivec.Object
}

func (o MTLDrawableObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLDrawableObjectFromID constructs a [MTLDrawableObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLDrawableObjectFromID(id objc.ID) MTLDrawableObject {
	return MTLDrawableObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Presents the drawable onscreen as soon as possible.
//
// # Discussion
//
// When a command queue schedules a command buffer for execution, it tracks
// whether any commands in that command buffer need to render or write to the
// drawable object. When you call this method, the drawable presents its
// contents as soon as possible after all scheduled render or write requests
// for that drawable are complete.
//
// See: https://developer.apple.com/documentation/Metal/MTLDrawable/present()
func (o MTLDrawableObject) Present() {
	objc.Send[struct{}](o.ID, objc.Sel("present"))
}

// Presents the drawable onscreen as soon as possible after a previous
// drawable is visible for the specified duration.
//
// duration: The previous drawable’s minimum display time, in seconds.
//
// # Discussion
//
// When a command queue schedules a command buffer for execution, it tracks
// whether any commands in that command buffer need to render or write to the
// drawable object. When you call this method, the drawable presents its
// contents at a future time when all render and write requests for that
// drawable are complete and a previous drawable has been visible onscreen for
// the specified duration. Use this method to schedule drawables at a regular
// interval.
//
// See: https://developer.apple.com/documentation/Metal/MTLDrawable/present(afterMinimumDuration:)
func (o MTLDrawableObject) PresentAfterMinimumDuration(duration corefoundation.CFTimeInterval) {
	objc.Send[struct{}](o.ID, objc.Sel("presentAfterMinimumDuration:"), duration)
}

// Presents the drawable onscreen at a specific host time.
//
// presentationTime: The Mach absolute time at which the drawable should be presented, in
// seconds.
//
// # Discussion
//
// When a command queue schedules a command buffer for execution, it tracks
// whether any commands in that command buffer need to render or write to the
// drawable object. When you call this method, the drawable waits until all
// render and write requests for that drawable are complete. If they complete
// prior to the specified time, the drawable presents the content at that
// time. If the commands complete after the presentation time, the drawable
// presents its contents as soon as possible.
//
// See: https://developer.apple.com/documentation/Metal/MTLDrawable/present(at:)
func (o MTLDrawableObject) PresentAtTime(presentationTime corefoundation.CFTimeInterval) {
	objc.Send[struct{}](o.ID, objc.Sel("presentAtTime:"), presentationTime)
}

// Registers a block of code to be called immediately after the drawable is
// presented.
//
// block: A block of code to be invoked.
//
// # Discussion
//
// You can register multiple handlers for a single drawable object.
//
// The following example code schedules a presentation handler that reads the
// [PresentedTime] property and uses it to derive the interval between the
// last and current presentation times. From that information, it determines
// the app’s frame rate.
//
// See: https://developer.apple.com/documentation/Metal/MTLDrawable/addPresentedHandler(_:)
func (o MTLDrawableObject) AddPresentedHandler(block MTLDrawableHandler) {
	_block0, _cleanup0 := NewMTLDrawableBlock(block)
	defer _cleanup0()
	objc.Send[struct{}](o.ID, objc.Sel("addPresentedHandler:"), objc.ID(_block0))
}

// A positive integer that identifies the drawable.
//
// # Discussion
//
// Drawable objects are usually owned by some other object, such as a
// [CAMetalLayer]. The owning object gives the first drawable it creates an ID
// of `0`, and it increments the ID by `1` for each additional drawable it
// creates.
//
// See: https://developer.apple.com/documentation/Metal/MTLDrawable/drawableID
//
// [CAMetalLayer]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer
func (o MTLDrawableObject) DrawableID() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("drawableID"))
	return uint(rv)
}

// The host time, in seconds, when the drawable was displayed onscreen.
//
// # Discussion
//
// Typically, you query this property in a callback method. See
// [AddPresentedHandler].
//
// The property value is `0.0` if the drawable hasn’t been presented or if
// its associated frame was dropped.
//
// See: https://developer.apple.com/documentation/Metal/MTLDrawable/presentedTime
func (o MTLDrawableObject) PresentedTime() corefoundation.CFTimeInterval {
	rv := objc.Send[corefoundation.CFTimeInterval](o.ID, objc.Sel("presentedTime"))
	return corefoundation.CFTimeInterval(rv)
}
