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

}

// IOGraphicsEngineContext
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/iographicsenginecontext
type IOGraphicsEngineContext struct {

}

// IONDRVControlParameters
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/iondrvcontrolparameters
type IONDRVControlParameters struct {

}

// IONamedValue
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/ionamedvalue
type IONamedValue struct {
	Value unsafe.Pointer
	Name unsafe.Pointer

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
	Length uint

}

// IORPC
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/iorpc
type IORPC struct {
	Message IORPCMessageMach
	Reply IORPCMessageMach
	ReplySize uint32
	SendSize uint32

}

// IORPCMessage
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/iorpcmessage
type IORPCMessage struct {
	Flags uint64
	Msgid uint64
	ObjectRefs uint64
	Objects unsafe.Pointer

}

// IORPCMessageErrorReturn
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/iorpcmessageerrorreturn
type IORPCMessageErrorReturn struct {
	Content IORPCMessageErrorReturnContent
	Mach IORPCMessageMach

}

// IORPCMessageErrorReturnContent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/iorpcmessageerrorreturncontent
type IORPCMessageErrorReturnContent struct {
	Hdr IORPCMessage
	Pad uint32
	Result int32

}

// IORPCMessageMach
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/iorpcmessagemach
type IORPCMessageMach struct {
	Msgh unsafe.Pointer
	Msgh_body unsafe.Pointer
	Objects unsafe.Pointer

}

// IOServiceInterestContent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/ioserviceinterestcontent-ieh
type IOServiceInterestContent struct {
	MessageArgument unsafe.Pointer
	MessageType objectivec.IObject

}

// IOServiceInterestContent64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/ioserviceinterestcontent64
type IOServiceInterestContent64 struct {
	MessageType objectivec.IObject
	MessageArgument unsafe.Pointer

}

// IOStreamBufferQueue
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/iostreambufferqueue
type IOStreamBufferQueue struct {

}

// IOStreamBufferQueueEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/iostreambufferqueueentry
type IOStreamBufferQueueEntry struct {

}

// IOUSB30HubPortStatusExt
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/iousb30hubportstatusext
type IOUSB30HubPortStatusExt struct {

}

// IOVideoControlDictionary
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/iovideocontroldictionary
type IOVideoControlDictionary struct {

}

// IOVideoStreamDictionary
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/iovideostreamdictionary
type IOVideoStreamDictionary struct {

}

// IOVideoStreamFormatDictionary
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/iovideostreamformatdictionary
type IOVideoStreamFormatDictionary struct {

}

// IOVirtualRange
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/iovirtualrange
type IOVirtualRange struct {
	Length uint
	Address IOVirtualAddress

}

// OSClassDescription
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/osclassdescription
type OSClassDescription struct {
	DescriptionSize uint32
	DispatchNames unsafe.Pointer
	Flags uint64
	MetaMethodNames unsafe.Pointer
	MetaMethodNamesOffset uint32
	MetaMethodNamesSize uint32
	MetaMethodOptions unsafe.Pointer
	MetaMethodOptionsOffset uint32
	MetaMethodOptionsSize uint32
	MethodNames unsafe.Pointer
	MethodNamesOffset uint32
	MethodNamesSize uint32
	MethodOptions unsafe.Pointer
	MethodOptionsOffset uint32
	MethodOptionsSize uint32
	Name unsafe.Pointer
	QueueNamesOffset uint32
	QueueNamesSize uint32
	Resv1 uint64
	SuperName unsafe.Pointer

}

// OSNotificationHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/osnotificationheader-iei
type OSNotificationHeader struct {
	Reference unsafe.Pointer
	Size uint32
	Type objectivec.IObject

}

// OSNotificationHeader64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/osnotificationheader64
type OSNotificationHeader64 struct {
	Reference unsafe.Pointer
	Size uint32
	Type objectivec.IObject

}

// StdFBShmem_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/stdfbshmem_t
type StdFBShmem_t struct {

}

// Applelabel
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/applelabel
type Applelabel struct {

}

// Bm12Cursor - Cursor image for 1-bit cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/bm12cursor
type Bm12Cursor struct {

}

// Bm18Cursor - Cursor image for 8-bit cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/bm18cursor
type Bm18Cursor struct {

}

// Bm34Cursor - Cursor image for 15-bit cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/bm34cursor
type Bm34Cursor struct {

}

// Bm38Cursor - Cursor image for 24-bit cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/bm38cursor
type Bm38Cursor struct {

}

// Disk_blk0
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/disk_blk0
type Disk_blk0 struct {

}

// Fdisk_part
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/fdisk_part
type Fdisk_part struct {

}

// Gpt_ent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/gpt_ent
type Gpt_ent struct {

}

// Gpt_hdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/gpt_hdr
type Gpt_hdr struct {

}

