// This file is hand-written (not generated). The generated binding in
// vz_graphics_display.gen.go models -[VZGraphicsDisplay
// _takeScreenshotWithCompletionHandler:] as taking a void(^)(NSError*) block,
// which drops the captured image.
//
// Reversing Virtualization.framework (macOS 26/Tahoe, arm64e) shows the true
// completion block is void(^)(CGImageRef, NSError*): the implementation routes
// through
//
//	screenshot_issue_wrong_state_error(VZVirtualMachineState,
//	    Base::CompletionHandler<CGImage*, NSError* __strong>)
//	screenshot_handle_frame_update(std::optional<VzCore::Hardware::FrameUpdate>,
//	    Base::CfPtr<CGColorSpace*>,
//	    Base::CompletionHandler<CGImage*, NSError* __strong>)
//
// i.e. the guest FrameUpdate delivered over XPC is turned into a CGImage that is
// handed to the completion handler as its first argument.

package virtualization

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// CGImageErrorHandler receives the captured image and error from
// -[VZGraphicsDisplay _takeScreenshotWithCompletionHandler:]. The image is a
// raw CGImageRef (an opaque pointer); it is valid only for the duration of the
// call unless the handler retains it with CGImageRetain.
type CGImageErrorHandler = func(image uintptr, err error)

// Block type encoding for void(^)(CGImageRef, NSError*): void return, block,
// an opaque CGImage pointer, and an NSError object.
var cgImageErrorBlockSignature = []byte("v@?^{CGImage=}@\"NSError\"\x00")

// NewCGImageErrorBlock wraps handler as an Objective-C block matching the real
// screenshot completion ABI, void(^)(CGImageRef, NSError*). The returned
// cleanup must not run until the block has fired, because the framework invokes
// it asynchronously.
func NewCGImageErrorBlock(handler CGImageErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, image objc.ID, errID objc.ID) {
		handler(uintptr(image), foundation.SafeErrorFrom(errID))
	})
	objc.SetBlockSignature(block, cgImageErrorBlockSignature)
	return objc.ID(block), func() { block.Release() }
}

// TakeScreenshotWithImageCompletionHandler calls the private selector
// _takeScreenshotWithCompletionHandler: with a block that receives the captured
// CGImage. The completion is asynchronous, so the returned cleanup releases the
// block and must be called only after handler has run (or never, if the caller
// abandons the request). An error is returned if the receiver does not respond
// to the selector.
func (v VZGraphicsDisplay) TakeScreenshotWithImageCompletionHandler(handler CGImageErrorHandler) (func(), error) {
	sel := objc.Sel("_takeScreenshotWithCompletionHandler:")
	if !objc.RespondsToSelector(v.ID, sel) {
		return func() {}, &objc.UnrecognizedSelectorError{Selector: "_takeScreenshotWithCompletionHandler:"}
	}
	block, cleanup := NewCGImageErrorBlock(handler)
	objc.SendIfResponds[objc.ID](v.ID, sel, block)
	return cleanup, nil
}
