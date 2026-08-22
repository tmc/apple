// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"encoding/binary"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
)

// C struct types

// MIDI2DeviceManufacturer
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDI2DeviceManufacturer
type MIDI2DeviceManufacturer struct {
	SysExIDByte [3]uint8
}

// MIDI2DeviceRevisionLevel
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDI2DeviceRevisionLevel
type MIDI2DeviceRevisionLevel struct {
	RevisionLevel [4]uint8
}

// MIDICIDeviceIdentification - A structure that describes a MIDI-CI device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceIdentification
type MIDICIDeviceIdentification struct {
	Manufacturer  [3]uint8 // The MIDI System Exclusive (SysEx) ID of the device manufacturer.
	Family        [2]uint8 // The group of familes to which the device belongs.
	ModelNumber   [2]uint8 // The device model number.
	RevisionLevel [4]uint8 // The revision number of the device model number.
	Reserved      [5]uint8 // A reserved field whose value is always zero.

}

// MIDICIProfileID is a C union type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileID
type MIDICIProfileID [5]byte

// Standard returns the union interpreted as *MIDICIProfileIDStandard.
// The returned pointer aliases the receiver's memory.
func (u *MIDICIProfileID) Standard() *MIDICIProfileIDStandard {
	return (*MIDICIProfileIDStandard)(unsafe.Pointer(u))
}

// ManufacturerSpecific returns the union interpreted as *MIDICIProfileIDManufacturerSpecific.
// The returned pointer aliases the receiver's memory.
func (u *MIDICIProfileID) ManufacturerSpecific() *MIDICIProfileIDManufacturerSpecific {
	return (*MIDICIProfileIDManufacturerSpecific)(unsafe.Pointer(u))
}

// MIDICIProfileIDManufacturerSpecific
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileIDManufacturerSpecific
type MIDICIProfileIDManufacturerSpecific struct {
	SysExID1 MIDIUInteger7
	SysExID2 MIDIUInteger7
	SysExID3 MIDIUInteger7
	Info1    MIDIUInteger7
	Info2    MIDIUInteger7
}

// MIDICIProfileIDStandard
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileIDStandard
type MIDICIProfileIDStandard struct {
	ProfileIDByte1 MIDIUInteger7
	ProfileBank    MIDIUInteger7
	ProfileNumber  MIDIUInteger7
	ProfileVersion MIDIUInteger7
	ProfileLevel   MIDIUInteger7
}

// MIDIControlTransform - A structure that describes the transformation of MIDI control change events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIControlTransform
type MIDIControlTransform struct {
	ControlType         MIDITransformControlType // The type of control specified by the control number.
	RemappedControlType MIDITransformControlType // The remapped control type.
	ControlNumber       uint16                   // The control number to affect.
	Transform           MIDITransformType        // The type of transformation to apply to the event values.
	Param               int16                    // An argument to the transformation method.

}

// MIDIDriverInterface - The interface to a MIDI driver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIDriverInterface
type MIDIDriverInterface struct {
	_reserved      unsafe.Pointer
	QueryInterface func(unsafe.Pointer, corefoundation.CFUUIDBytes, unsafe.Pointer) int32
	AddRef         func(unsafe.Pointer) uint32
	Release        func(unsafe.Pointer) uint32
	FindDevices    func(uintptr, uint32) int32                                  // Finds the available devices.
	Start          func(uintptr, uint32) int32                                  // Starts MIDI I/O.
	Stop           func(uintptr) int32                                          // Stops MIDI I/O.
	Configure      func(uintptr, uint32) int32                                  // The system doesn’t currently use this method.
	Send           func(uintptr, uintptr, unsafe.Pointer, unsafe.Pointer) int32 // Sends a MIDI packet list to the specified destination endpoints.
	EnableSource   func(uintptr, uint32, uint8) int32                           // Tells the driver whether input from a particular source has listeners.
	Flush          func(uintptr, uint32, unsafe.Pointer, unsafe.Pointer) int32  // Unschedules all pending output to the specified destination.
	Monitor        func(uintptr, uint32, uintptr) int32                         // Enables monitoring of MIDI packet lists by the specified driver.
	SendPackets    func(uintptr, uintptr, unsafe.Pointer, unsafe.Pointer) int32 // Sends a MIDI event list to the specified destination endpoints.
	MonitorEvents  func(uintptr, uint32, uintptr) int32                         // Enables monitoring of MIDI event lists by the specified driver.

}

// MIDIEventList - A variable-length list of MIDI event packets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIEventList
type MIDIEventList struct {
	Protocol   MIDIProtocolID     // The MIDI protocol variant of the events in the list.
	NumPackets uint32             // The number of MIDI event packet structures in the list.
	Packet     [1]MIDIEventPacket // An array of variable-length MIDI event packet structures.

}

// MIDIEventPacket - A series of simultaneous MIDI events in Universal MIDI Packets (UMP) format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIEventPacket
type MIDIEventPacket struct {
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
	storage [268]byte
}

// TimeStamp returns the TimeStamp field from the record's packed storage.
func (s *MIDIEventPacket) TimeStamp() MIDITimeStamp {
	return MIDITimeStamp(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetTimeStamp updates the TimeStamp field in the record's packed storage.
func (s *MIDIEventPacket) SetTimeStamp(v MIDITimeStamp) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// WordCount returns the WordCount field from the record's packed storage.
func (s *MIDIEventPacket) WordCount() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetWordCount updates the WordCount field in the record's packed storage.
func (s *MIDIEventPacket) SetWordCount(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// Words returns the Words field from the record's packed storage.
func (s *MIDIEventPacket) Words() [64]uint32 {
	return *(*[64]uint32)(unsafe.Pointer(&s.storage[12]))
}

// SetWords updates the Words field in the record's packed storage.
func (s *MIDIEventPacket) SetWords(v [64]uint32) {
	*(*[64]uint32)(unsafe.Pointer(&s.storage[12])) = v
}

// MIDIIOErrorNotification - A general I/O error notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIIOErrorNotification
type MIDIIOErrorNotification struct {
	MessageID    MIDINotificationMessageID // The type of message.
	MessageSize  uint32                    // The size of the message.
	DriverDevice MIDIDeviceRef             // The device with an I/O error.
	ErrorCode    int32                     // The error code of the generated error.

}

// MIDIMessage_128 - A 128-bit MIDI message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIMessage_128
type MIDIMessage_128 struct {
	Word0 uint32
	Word1 uint32
	Word2 uint32
	Word3 uint32
}

// MIDIMessage_64 - A 64-bit MIDI message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIMessage_64
type MIDIMessage_64 struct {
	Word0 uint32
	Word1 uint32
}

// MIDIMessage_96 - A 96-bit MIDI message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIMessage_96
type MIDIMessage_96 struct {
	Word0 uint32
	Word1 uint32
	Word2 uint32
}

// MIDINotification - A message that describes a system state change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINotification
type MIDINotification struct {
	MessageID   MIDINotificationMessageID // An identifier that describes the type of state change.
	MessageSize uint32                    // The size of the message including its ID.

}

// MIDIObjectAddRemoveNotification - A message that describes the addition or removal of an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectAddRemoveNotification
type MIDIObjectAddRemoveNotification struct {
	MessageID   MIDINotificationMessageID // The message type.
	MessageSize uint32                    // The message size.
	Parent      MIDIObjectRef             // The parent object of the added or removed child.
	ParentType  MIDIObjectType            // The parent object type.
	Child       MIDIObjectRef             // The added or removed child object.
	ChildType   MIDIObjectType            // The child object type.

}

// MIDIObjectPropertyChangeNotification - A message that describes the change to an object property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectPropertyChangeNotification
type MIDIObjectPropertyChangeNotification struct {
	MessageID    MIDINotificationMessageID  // The message type.
	MessageSize  uint32                     // The message size.
	Object       MIDIObjectRef              // The object whose property changed.
	ObjectType   MIDIObjectType             // The object type.
	PropertyName corefoundation.CFStringRef // The name of the modified property.

}

// MIDIPacket - A collection of simultaneous MIDI events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIPacket
type MIDIPacket struct {
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
	storage [268]byte
}

// TimeStamp returns the TimeStamp field from the record's packed storage.
func (s *MIDIPacket) TimeStamp() MIDITimeStamp {
	return MIDITimeStamp(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetTimeStamp updates the TimeStamp field in the record's packed storage.
func (s *MIDIPacket) SetTimeStamp(v MIDITimeStamp) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// Length returns the Length field from the record's packed storage.
func (s *MIDIPacket) Length() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[8:10]))
}

// SetLength updates the Length field in the record's packed storage.
func (s *MIDIPacket) SetLength(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[8:10], uint16(v))
}

// Data returns the Data field from the record's packed storage.
func (s *MIDIPacket) Data() [256]uint8 {
	return *(*[256]uint8)(unsafe.Pointer(&s.storage[10]))
}

// SetData updates the Data field in the record's packed storage.
func (s *MIDIPacket) SetData(v [256]uint8) {
	*(*[256]uint8)(unsafe.Pointer(&s.storage[10])) = v
}

// MIDIPacketList - A list of MIDI events the system sends to or receives from an endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIPacketList
type MIDIPacketList struct {
	NumPackets uint32        // The number of MIDI packets in the list.
	Packet     [1]MIDIPacket // An open-ended array of variable-length MIDI packets.

}

// MIDISysexSendRequest - A request to asynchronously send a single system-exclusive (SysEx) event to a destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISysexSendRequest
type MIDISysexSendRequest struct {
	Destination      MIDIEndpointRef    // The endpoint to send the event to.
	Data             *uint8             // The request’s data.
	BytesToSend      uint32             // The number of bytes to send.
	Complete         bool               // A Boolean value that indicates whether the transmission is complete.
	Reserved         [3]uint8           // A field that’s reserved for future use.
	CompletionProc   MIDICompletionProc // A function that the system calls after it sends all bytes for the request, or after the client marks the request as complete.
	CompletionRefCon unsafe.Pointer     // Data to pass to the completion function.

}

// MIDISysexSendRequestUMP - A request to asynchronously send a single universal MIDI packet (UMP) system-exclusive (SysEx) event to a destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDISysexSendRequestUMP
type MIDISysexSendRequestUMP struct {
	Destination      MIDIEndpointRef       // The endpoint to send the event to.
	Words            *uint32               // A pointer to the event to send, which the system advances as it sends the data.
	WordsToSend      uint32                // A counter of the number of words to send.
	Complete         bool                  // A Boolean value that indicates whether the transmission is complete.
	CompletionProc   MIDICompletionProcUMP // A function that the system calls after it sends all data for the request, or after the client marks the request as complete.
	CompletionRefCon unsafe.Pointer        // Data to pass to the completion function.

}

// MIDIThruConnectionEndpoint - A source or destination in a MIDI thru connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIThruConnectionEndpoint
type MIDIThruConnectionEndpoint struct {
	EndpointRef MIDIEndpointRef // The endpoint reference.
	UniqueID    MIDIUniqueID    // The connection’s unique identifier.

}

// MIDIThruConnectionParams - A set of MIDI routings and transformations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIThruConnectionParams
type MIDIThruConnectionParams struct {
	Version              uint32                        // The version number.
	NumSources           uint32                        // The number of valid sources.
	Sources              [8]MIDIThruConnectionEndpoint // All MIDI sources for this connection.
	NumDestinations      uint32                        // The number of valid destinations.
	Destinations         [8]MIDIThruConnectionEndpoint // All MIDI destinations for this connection.
	ChannelMap           [16]uint8                     // A mapping of MIDI channels.
	LowVelocity          uint8                         // The velocity value below which the system filters out notes.
	HighVelocity         uint8                         // The velocity value above which the system filters out notes.
	LowNote              uint8                         // The note value below which the system filters out notes.
	HighNote             uint8                         // The note value above which the system filters out notes.
	NoteNumber           MIDITransform                 // The transformation of MIDI note numbers.
	Velocity             MIDITransform                 // A note velocity transformation.
	KeyPressure          MIDITransform                 // The transformation of polyphonic key pressure events.
	ChannelPressure      MIDITransform                 // The transformation of MIDI monophonic channel pressure events.
	ProgramChange        MIDITransform                 // A transformation of a MIDI program change event.
	PitchBend            MIDITransform                 // The transformation of a MIDI pitch bend event.
	FilterOutSysEx       uint8                         // A value that indicates wheter to filter out system-exclusive messages.
	FilterOutMTC         uint8                         // A value that indicates whether to filter out MIDI Time Code messages.
	FilterOutBeatClock   uint8                         // A value that indicates whether to filter out MIDI clock, play, stop, and resume messages.
	FilterOutTuneRequest uint8                         // A value that specifies whether to filter out MIDI tune request messages.
	Reserved2            [3]uint8                      // A reserved value that must be 0.
	FilterOutAllControls uint8                         // A value that indicates whether to filter out MIDI continuous control messages.
	NumControlTransforms uint16                        // The number of control transformations in the variable-length portion of the struct.
	NumMaps              uint16                        // The number of MIDI value maps in the variable-length portion of the struct.
	Reserved3            [4]uint16                     // A reserved value that must be 0.

}

// MIDITransform - The transformation of a single type of MIDI event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDITransform
type MIDITransform struct {
	Transform MIDITransformType // The type of transformation to apply to the event values.
	Param     int16             // An argument to the transformation method (see description of MIDITransformType).

}

// MIDIUniversalMessage
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUniversalMessage
type MIDIUniversalMessage struct {
	Type                     MIDIMessageType
	Group                    uint8
	Reserved                 [3]uint8
	Status                   MIDIUtilityStatus
	JitterReductionClock     uint16
	JitterReductionTimestamp uint16
	ChannelVoice1            unsafe.Pointer
	ChannelVoice2            unsafe.Pointer
	Data128                  unsafe.Pointer
	SysEx                    unsafe.Pointer
	System                   unsafe.Pointer
	Unknown                  unsafe.Pointer
	Utility                  unsafe.Pointer
}

// MIDIValueMap - A custom lookup table to transform MIDI 7-bit values, as contained in note numbers, velocities, control values, and so on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIValueMap
type MIDIValueMap struct {
	Value [128]uint8 // The array of unsigned 8-bit integers.

}
