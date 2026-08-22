// Code generated from Apple documentation for ParavirtualizedGraphics. DO NOT EDIT.

package paravirtualizedgraphics

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An object that provides display functionality to the guest operating system in a way that the host-side virtual machine app can intercept.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplay
type PGDisplay interface {
	objectivec.IObject

	// Encodes Metal commands to process the current frame and write it to a texture.
	//
	// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplay/encodeCurrentFrame(to:texture:region:)
	EncodeCurrentFrameToCommandBufferTextureRegion(commandBuffer metal.MTLCommandBuffer, texture metal.MTLTexture, region metal.MTLRegion) bool

	// The list of display modes that the virtual display supports.
	//
	// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplay/modeList
	ModeList() []PGDisplayMode
	SetModeList(value []PGDisplayMode)

	// The current cursor location in the guest environment.
	//
	// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplay/cursorPosition
	CursorPosition() PGDisplayCoord_t

	// The number of frame presents that the guest has generated since object creation.
	//
	// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplay/guestPresentCount
	GuestPresentCount() uint

	// The number of unique frames that the host has encoded since object creation.
	//
	// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplay/hostPresentCount
	HostPresentCount() uint

	// The Metal texture usage flags necessary for any texture that can be a destination for frame data.
	//
	// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplay/minimumTextureUsage
	MinimumTextureUsage() metal.MTLTextureUsage

	// The display’s name that you specified at creation time.
	//
	// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplay/name
	Name() string

	// The display’s serial number that you specified at creation time.
	//
	// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplay/serialNum
	SerialNum() uint32

	// The display’s accelerator port that you specified at creation time.
	//
	// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplay/port
	Port() uint

	// The display’s virtual dimensions, in millimeters, that you specified at creation time.
	//
	// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplay/sizeInMillimeters
	SizeInMillimeters() corefoundation.CGSize

	// The queue that the framework uses when dispatching messages to any of the display’s registered handlers.
	//
	// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplay/queue
	Queue() dispatch.Queue

	// A handler that the framework calls to change the cursor’s appearance.
	//
	// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplay/cursorGlyphHandler
	CursorGlyphHandler() PGDisplayCursorGlyphHandler

	// A handler that the framework calls to change the cursor’s visibility.
	//
	// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplay/cursorShowHandler
	CursorShowHandler() PGDisplayCursorShowHandler

	// A handler that the framework calls to change the virtual display’s graphics mode.
	//
	// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplay/modeChangeHandler
	ModeChangeHandler() PGDisplayModeChangeHandler

	// A handler that the framework calls when the guest environment has a new frame to display.
	//
	// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplay/newFrameEventHandler
	NewFrameEventHandler() PGDisplayNewFrameEventHandler

	// cursorMoveHandler protocol.
	//
	// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplay/cursorMoveHandler
	CursorMoveHandler() PGDisplayCursorMoveHandler
}

// PGDisplayObject wraps an existing Objective-C object that conforms to the PGDisplay protocol.
type PGDisplayObject struct {
	objectivec.Object
}

func (o PGDisplayObject) BaseObject() objectivec.Object {
	return o.Object
}

// PGDisplayObjectFromID constructs a [PGDisplayObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func PGDisplayObjectFromID(id objc.ID) PGDisplayObject {
	return PGDisplayObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Encodes Metal commands to process the current frame and write it to a
// texture.
//
// commandBuffer: The Metal command buffer for encoding the processing of the new frame.
//
// texture: The destination texture.
//
// region: The region within the Metal texture for the frame data to land in.
//
// # Return Value
//
// Returns true if the method encoded the commands successfully; otherwise
// false.
//
// # Discussion
//
// This method encodes Metal commands to process the new frame, dealing with
// color matching, gamma, and other factors. The GPU writes data to a region
// inside the destination texture, linearly scaled as necessary to fit in the
// specified region. The encoded commands won’t change any texture data
// outside the specified region.
//
// The texture and command buffer must share the same Metal device object that
// you used to create the virtual graphics device. The destination texture
// must specify all of the usage bits that the display’s
// [MinimumTextureUsage] property provides, but can specify additional flags.
//
// Typically, you encode frame processing in response to a call to the frame
// event handler you specified when you created the display object. See
// [PGDisplayDescriptor.NewFrameEventHandler] for more information.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplay/encodeCurrentFrame(to:texture:region:)
func (o PGDisplayObject) EncodeCurrentFrameToCommandBufferTextureRegion(commandBuffer metal.MTLCommandBuffer, texture metal.MTLTexture, region metal.MTLRegion) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("encodeCurrentFrameToCommandBuffer:texture:region:"), commandBuffer, texture, region)
	return rv
}

// The list of display modes that the virtual display supports.
//
// # Discussion
//
// The maximum number of display modes is `128`. Setting this property updates
// the virtual graphics device’s supported mode list, and potentially forces
// it to change its current mode. The first time you set this property, the
// device simulates hot-plugging the display to the graphics device.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplay/modeList
func (o PGDisplayObject) ModeList() []PGDisplayMode {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("modeList"))
	result := make([]PGDisplayMode, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = PGDisplayModeFromID(id)
	}
	return result
}

func (o PGDisplayObject) SetModeList(value []PGDisplayMode) {
	objc.Send[struct{}](o.ID, objc.Sel("setModeList:"), objectivec.IObjectSliceToNSArray(value))
}

// The current cursor location in the guest environment.
//
// # Discussion
//
// If the cursor isn’t on the display, this property’s value is `(0xFFFF,
// 0xFFFF)`.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplay/cursorPosition
func (o PGDisplayObject) CursorPosition() PGDisplayCoord_t {
	rv := objc.Send[PGDisplayCoord_t](o.ID, objc.Sel("cursorPosition"))
	return PGDisplayCoord_t(rv)
}

// The number of frame presents that the guest has generated since object
// creation.
//
// # Discussion
//
// This value can exceed the number of times that the framework invokes the
// [PGDisplayDescriptor.NewFrameEventHandler] block if the host isn’t
// encoding frames fast enough to keep up.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplay/guestPresentCount
func (o PGDisplayObject) GuestPresentCount() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("guestPresentCount"))
	return uint(rv)
}

// The number of unique frames that the host has encoded since object
// creation.
//
// # Discussion
//
// The value of this property can be smaller than the number of times you’ve
// called the [EncodeCurrentFrameToCommandBufferTextureRegion] method if you
// encode the same frame multiple times.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplay/hostPresentCount
func (o PGDisplayObject) HostPresentCount() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("hostPresentCount"))
	return uint(rv)
}

// The Metal texture usage flags necessary for any texture that can be a
// destination for frame data.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplay/minimumTextureUsage
func (o PGDisplayObject) MinimumTextureUsage() metal.MTLTextureUsage {
	rv := objc.Send[metal.MTLTextureUsage](o.ID, objc.Sel("minimumTextureUsage"))
	return metal.MTLTextureUsage(rv)
}

// The display’s name that you specified at creation time.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplay/name
func (o PGDisplayObject) Name() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// The display’s serial number that you specified at creation time.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplay/serialNum
func (o PGDisplayObject) SerialNum() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("serialNum"))
	return uint32(rv)
}

// The display’s accelerator port that you specified at creation time.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplay/port
func (o PGDisplayObject) Port() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("port"))
	return uint(rv)
}

// The display’s virtual dimensions, in millimeters, that you specified at
// creation time.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplay/sizeInMillimeters
func (o PGDisplayObject) SizeInMillimeters() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](o.ID, objc.Sel("sizeInMillimeters"))
	return corefoundation.CGSize(rv)
}

// The queue that the framework uses when dispatching messages to any of the
// display’s registered handlers.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplay/queue
func (o PGDisplayObject) Queue() dispatch.Queue {
	rv := objc.Send[uintptr](o.ID, objc.Sel("queue"))
	return dispatch.QueueFromHandle(rv)
}

// A handler that the framework calls to change the cursor’s appearance.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplay/cursorGlyphHandler
func (o PGDisplayObject) CursorGlyphHandler() PGDisplayCursorGlyphHandler {
	rv := objc.Send[PGDisplayCursorGlyphHandler](o.ID, objc.Sel("cursorGlyphHandler"))
	return PGDisplayCursorGlyphHandler(rv)
}

// A handler that the framework calls to change the cursor’s visibility.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplay/cursorShowHandler
func (o PGDisplayObject) CursorShowHandler() PGDisplayCursorShowHandler {
	rv := objc.Send[PGDisplayCursorShowHandler](o.ID, objc.Sel("cursorShowHandler"))
	return PGDisplayCursorShowHandler(rv)
}

// A handler that the framework calls to change the virtual display’s
// graphics mode.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplay/modeChangeHandler
func (o PGDisplayObject) ModeChangeHandler() PGDisplayModeChangeHandler {
	rv := objc.Send[PGDisplayModeChangeHandler](o.ID, objc.Sel("modeChangeHandler"))
	return PGDisplayModeChangeHandler(rv)
}

// A handler that the framework calls when the guest environment has a new
// frame to display.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplay/newFrameEventHandler
func (o PGDisplayObject) NewFrameEventHandler() PGDisplayNewFrameEventHandler {
	rv := objc.Send[PGDisplayNewFrameEventHandler](o.ID, objc.Sel("newFrameEventHandler"))
	return PGDisplayNewFrameEventHandler(rv)
}

// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplay/cursorMoveHandler
func (o PGDisplayObject) CursorMoveHandler() PGDisplayCursorMoveHandler {
	rv := objc.Send[PGDisplayCursorMoveHandler](o.ID, objc.Sel("cursorMoveHandler"))
	return PGDisplayCursorMoveHandler(rv)
}
