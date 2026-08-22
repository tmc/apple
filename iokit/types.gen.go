// Code generated from Apple documentation for iokit. DO NOT EDIT.

package iokit

import (
	"encoding/binary"
	"unsafe"

	"github.com/tmc/apple/kernel"
	"github.com/tmc/apple/objc"
)

// C struct types

// IOAsyncCompletionContent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/ioasynccompletioncontent
type IOAsyncCompletionContent struct {
	Result kernel.IOReturn
}

// IOGraphicsEngineContext
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/iographicsenginecontext
type IOGraphicsEngineContext struct {
	ContextLock int32
	State       uint32
	Owner       unsafe.Pointer
	Version     uint32
	StructSize  uint
	Reserved    [8]uint32
}

// IONamedValue
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/ionamedvalue
type IONamedValue struct {
	Value int32
	Name  *byte
}

// IOPMSystemCapabilityChangeParameters - A structure describing a system capability change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/iopmsystemcapabilitychangeparameters
type IOPMSystemCapabilityChangeParameters struct {
	NotifyRef        uint32
	MaxWaitForReply  uint32
	ChangeFlags      uint32
	__reserved1      uint32
	FromCapabilities uint32
	ToCapabilities   uint32
	__reserved2      [4]uint32
}

// IOPhysicalRange
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/iophysicalrange
type IOPhysicalRange struct {
	Address kernel.IOPhysicalAddress
	Length  uint
}

// IORPC
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/iorpc
type IORPC struct {
	Message   *IORPCMessageMach
	Reply     *IORPCMessageMach
	SendSize  uint32
	ReplySize uint32
}

// IORPCMessage
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/iorpcmessage
type IORPCMessage struct {
	Msgid      uint64
	Flags      uint64
	ObjectRefs uint64
}

// IORPCMessageErrorReturnContent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/iorpcmessageerrorreturncontent
type IORPCMessageErrorReturnContent struct {
	Hdr    IORPCMessage
	Result int32
	Pad    uint32
}

// IORPCMessageMach
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/iorpcmessagemach
type IORPCMessageMach struct {
	Msgh      unsafe.Pointer
	Msgh_body kernel.Mach_msg_body_t
}

// IOServiceInterestContent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/ioserviceinterestcontent
type IOServiceInterestContent struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint32
	storage [12]byte
}

// MessageType returns the MessageType field from the record's packed storage.
func (s *IOServiceInterestContent) MessageType() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetMessageType updates the MessageType field in the record's packed storage.
func (s *IOServiceInterestContent) SetMessageType(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// MessageArgument returns the MessageArgument field from the record's packed storage.
func (s *IOServiceInterestContent) MessageArgument() *objc.ID {
	return *(**objc.ID)(unsafe.Pointer(&s.storage[4]))
}

// SetMessageArgument updates the MessageArgument field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *IOServiceInterestContent) SetMessageArgument(v *objc.ID) {
	*(**objc.ID)(unsafe.Pointer(&s.storage[4])) = v
}

// IOStreamBufferQueue
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/iostreambufferqueue
type IOStreamBufferQueue struct {
	EntryCount uint32
	HeadIndex  uint32
	TailIndex  uint32
	Reserved   uint32 // Reserved for future use.

}

// IOStreamBufferQueueEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/iostreambufferqueueentry
type IOStreamBufferQueueEntry struct {
	BufferID      IOStreamBufferID
	DataOffset    uint32
	DataLength    uint32
	ControlOffset uint32
	ControlLength uint32
	Reserved      [3]uint32 // Reserved for future use.

}

// IOVirtualRange
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/iovirtualrange
type IOVirtualRange struct {
	Address kernel.IOVirtualAddress
	Length  uint
}

// OSNotificationHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/osnotificationheader
type OSNotificationHeader struct {
	Size      uint32
	Type      uint32
	Reference [8]uint32
}

// Bm12Cursor - Cursor image for 1-bit cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/bm12cursor
type Bm12Cursor struct {
	Image [4][16]uint32 // This array contains the cursor images.
	Mask  [4][16]uint32 // This array contains the cursor mask.
	Save  [16]uint32    // This array stores the pixel values of the region underneath the cursor in its last drawn position.

}

// Bm18Cursor - Cursor image for 8-bit cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/bm18cursor
type Bm18Cursor struct {
	Image [4][256]uint8 // This array contains cursor color values, which are converted to displayed colors through the color table. The array is two dimensional and its first index is the cursor frame and the second index is the cursor pixel.
	Mask  [4][256]uint8 // This array contains the cursor alpha mask. The array is two dimensional with the same indexing as the image. If an alpha mask pixel is 0 and the corresponding image pixel is set to white for the display, then this cursor pixel will invert pixels on the display.
	Save  [256]uint8    // This array stores the color values of the region underneath the cursor in its last drawn position.

}

// Bm34Cursor - Cursor image for 15-bit cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/bm34cursor
type Bm34Cursor struct {
	Image [4][256]uint16 // This array defines the cursor color values and transparency. The array is two dimensional and its first index is the cursor frame and the second index is the cursor pixel. A value of 0 means the pixel is transparent.
	Save  [256]uint16    // This array stores the color values of the region underneath the cursor in its last drawn position.

}

// Bm38Cursor - Cursor image for 24-bit cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/bm38cursor
type Bm38Cursor struct {
	Image [4][256]uint32 // This array defines the cursor color values and transparency. The array is two dimensional and its first index is the cursor frame and the second index is the cursor pixel. The lower 24 bits of a pixel's value contain the RGB color, while the upper 8 bits contain the alpha value.
	Save  [256]uint32    // This array stores the color values of the region underneath the cursor in its last drawn position.

}

// Disk_blk0
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/disk_blk0
type Disk_blk0 struct {
	Bootcode  [446]uint8
	Parts     [4]Fdisk_part
	Signature uint16
}

// Fdisk_part
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/fdisk_part
type Fdisk_part struct {
	Bootid  uint8
	Beghead uint8
	Begsect uint8
	Begcyl  uint8
	Systid  uint8
	Endhead uint8
	Endsect uint8
	Endcyl  uint8
	Relsect uint32
	Numsect uint32
}

// Gpt_ent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/gpt_ent
type Gpt_ent struct {
	Ent_type      [16]uint8
	Ent_uuid      [16]uint8
	Ent_lba_start uint64
	Ent_lba_end   uint64
	Ent_attr      uint64
	Ent_name      [36]uint16
}

// Gpt_hdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iokit/gpt_hdr
type Gpt_hdr struct {
	Hdr_sig       [8]uint8
	Hdr_revision  uint32
	Hdr_size      uint32
	Hdr_crc_self  uint32
	__reserved    uint32
	Hdr_lba_self  uint64
	Hdr_lba_alt   uint64
	Hdr_lba_start uint64
	Hdr_lba_end   uint64
	Hdr_uuid      [16]uint8
	Hdr_lba_table uint64
	Hdr_entries   uint32
	Hdr_entsz     uint32
	Hdr_crc_table uint32
	Padding       uint32
}
