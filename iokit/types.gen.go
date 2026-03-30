// Code generated from Apple documentation for iokit. DO NOT EDIT.

package iokit

import (
	"unsafe"

	"github.com/tmc/apple/objectivec"
)

// C struct types

// IOAsyncCompletionContent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/ioasynccompletioncontent-ieg
type IOAsyncCompletionContent struct {
	Result int
	Args   unsafe.Pointer
}

// IOGraphicsEngineContext
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/iographicsenginecontext
type IOGraphicsEngineContext struct {
	Owner    unsafe.Pointer
	Reserved unsafe.Pointer
	State    unsafe.Pointer
	Version  unsafe.Pointer
}

// IONamedValue
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/ionamedvalue
type IONamedValue struct {
	Value unsafe.Pointer
	Name  unsafe.Pointer
}

// IOPMSystemCapabilityChangeParameters - A structure describing a system capability change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/iopmsystemcapabilitychangeparameters
type IOPMSystemCapabilityChangeParameters struct {
}

// IOPhysicalRange
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/iophysicalrange
type IOPhysicalRange struct {
	Address IOPhysicalAddress
	Length  uint
}

// IORPC
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/iorpc
type IORPC struct {
	Message   IORPCMessageMach
	Reply     IORPCMessageMach
	ReplySize uint32
	SendSize  uint32
}

// IORPCMessage
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/iorpcmessage
type IORPCMessage struct {
	Flags      uint64
	Msgid      uint64
	ObjectRefs uint64
	Objects    unsafe.Pointer
}

// IORPCMessageErrorReturnContent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/iorpcmessageerrorreturncontent
type IORPCMessageErrorReturnContent struct {
	Hdr    IORPCMessage
	Pad    uint32
	Result int32
}

// IORPCMessageMach
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/iorpcmessagemach
type IORPCMessageMach struct {
	Msgh      unsafe.Pointer
	Msgh_body unsafe.Pointer
	Objects   unsafe.Pointer
}

// IOServiceInterestContent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/ioserviceinterestcontent-ieh
type IOServiceInterestContent struct {
	MessageType     objectivec.IObject
	MessageArgument unsafe.Pointer
}

// IOStreamBufferQueue
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/iostreambufferqueue
type IOStreamBufferQueue struct {
	Queue    unsafe.Pointer // The array of queue entries.
	Reserved unsafe.Pointer // Reserved for future use.

}

// IOStreamBufferQueueEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/iostreambufferqueueentry
type IOStreamBufferQueueEntry struct {
	Reserved unsafe.Pointer // Reserved for future use.

}

// IOVirtualRange
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/iovirtualrange
type IOVirtualRange struct {
	Address IOVirtualAddress
	Length  uint
}

// OSNotificationHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/osnotificationheader-iei
type OSNotificationHeader struct {
	Size      uint32
	Type      objectivec.IObject
	Reference unsafe.Pointer
	Content   unsafe.Pointer
}

// Bm12Cursor - Cursor image for 1-bit cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/bm12cursor
type Bm12Cursor struct {
	Image unsafe.Pointer // This array contains the cursor images.
	Mask  unsafe.Pointer // This array contains the cursor mask.
	Save  unsafe.Pointer // This array stores the pixel values of the region underneath the cursor in its last drawn position.

}

// Bm18Cursor - Cursor image for 8-bit cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/bm18cursor
type Bm18Cursor struct {
	Image unsafe.Pointer // This array contains cursor color values, which are converted to displayed colors through the color table. The array is two dimensional and its first index is the cursor frame and the second index is the cursor pixel.
	Mask  unsafe.Pointer // This array contains the cursor alpha mask. The array is two dimensional with the same indexing as the image. If an alpha mask pixel is 0 and the corresponding image pixel is set to white for the display, then this cursor pixel will invert pixels on the display.
	Save  unsafe.Pointer // This array stores the color values of the region underneath the cursor in its last drawn position.

}

// Bm34Cursor - Cursor image for 15-bit cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/bm34cursor
type Bm34Cursor struct {
	Image unsafe.Pointer // This array defines the cursor color values and transparency. The array is two dimensional and its first index is the cursor frame and the second index is the cursor pixel. A value of 0 means the pixel is transparent.
	Save  unsafe.Pointer // This array stores the color values of the region underneath the cursor in its last drawn position.

}

// Bm38Cursor - Cursor image for 24-bit cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/bm38cursor
type Bm38Cursor struct {
	Image unsafe.Pointer // This array defines the cursor color values and transparency. The array is two dimensional and its first index is the cursor frame and the second index is the cursor pixel. The lower 24 bits of a pixel's value contain the RGB color, while the upper 8 bits contain the alpha value.
	Save  unsafe.Pointer // This array stores the color values of the region underneath the cursor in its last drawn position.

}
