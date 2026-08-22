// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"encoding/binary"
	"math"
	"unsafe"

	"github.com/tmc/apple/coreaudiotypes"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coremidi"
	"github.com/tmc/apple/os"
)

// C struct types

// AUChannelInfo - The audio input and output channel capabilities for an audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUChannelInfo
type AUChannelInfo struct {
	InChannels  int16 // The number of input channels.
	OutChannels int16 // The number of output channels.

}

// AUDependentParameter - An audio unit parameter whose value can change in response to a change in its parent metaparameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUDependentParameter
type AUDependentParameter struct {
	MScope       AudioUnitScope
	MParameterID AudioUnitParameterID
}

// AUDistanceAttenuationData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUDistanceAttenuationData
type AUDistanceAttenuationData struct {
	InNumberOfPairs uint32
	InDistance      float32
	OutGain         float32
	Pairs           unsafe.Pointer
}

// AUHostIdentifier
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUHostIdentifier
type AUHostIdentifier struct {
	HostName    corefoundation.CFStringRef
	HostVersion AUNumVersion
}

// AUHostVersionIdentifier - The name and version of an audio unit’s host application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUHostVersionIdentifier
type AUHostVersionIdentifier struct {
	HostName    corefoundation.CFStringRef
	HostVersion uint32
}

// AUInputSamplesInOutputCallbackStruct - The callback function and custom data for providing input-to-output sample mapping for an audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUInputSamplesInOutputCallbackStruct
type AUInputSamplesInOutputCallbackStruct struct {
	InputToOutputCallback AUInputSamplesInOutputCallback // The callback function that provides input-to-output sample mapping for an audio unit.
	UserData              unsafe.Pointer                 // Custom data for input-to-output sample mapping for an audio unit.

}

// AUMIDIEvent - A structure that describes a scheduled MIDI event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUMIDIEvent
type AUMIDIEvent struct {
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
	storage [24]byte
}

// Next returns the Next field from the record's packed storage.
func (s *AUMIDIEvent) Next() *AURenderEvent {
	return *(**AURenderEvent)(unsafe.Pointer(&s.storage[0]))
}

// SetNext updates the Next field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *AUMIDIEvent) SetNext(v *AURenderEvent) {
	*(**AURenderEvent)(unsafe.Pointer(&s.storage[0])) = v
}

// EventSampleTime returns the EventSampleTime field from the record's packed storage.
func (s *AUMIDIEvent) EventSampleTime() AUEventSampleTime {
	return AUEventSampleTime(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetEventSampleTime updates the EventSampleTime field in the record's packed storage.
func (s *AUMIDIEvent) SetEventSampleTime(v AUEventSampleTime) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// EventType returns the EventType field from the record's packed storage.
func (s *AUMIDIEvent) EventType() AURenderEventType {
	return *(*AURenderEventType)(unsafe.Pointer(&s.storage[16]))
}

// SetEventType updates the EventType field in the record's packed storage.
func (s *AUMIDIEvent) SetEventType(v AURenderEventType) {
	*(*AURenderEventType)(unsafe.Pointer(&s.storage[16])) = v
}

// Reserved returns the Reserved field from the record's packed storage.
func (s *AUMIDIEvent) Reserved() uint8 {
	return uint8(s.storage[17])
}

// SetReserved updates the Reserved field in the record's packed storage.
func (s *AUMIDIEvent) SetReserved(v uint8) {
	s.storage[17] = uint8(v)
}

// Length returns the Length field from the record's packed storage.
func (s *AUMIDIEvent) Length() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[18:20]))
}

// SetLength updates the Length field in the record's packed storage.
func (s *AUMIDIEvent) SetLength(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[18:20], uint16(v))
}

// Cable returns the Cable field from the record's packed storage.
func (s *AUMIDIEvent) Cable() uint8 {
	return uint8(s.storage[20])
}

// SetCable updates the Cable field in the record's packed storage.
func (s *AUMIDIEvent) SetCable(v uint8) {
	s.storage[20] = uint8(v)
}

// Data returns the Data field from the record's packed storage.
func (s *AUMIDIEvent) Data() [3]uint8 {
	return *(*[3]uint8)(unsafe.Pointer(&s.storage[21]))
}

// SetData updates the Data field in the record's packed storage.
func (s *AUMIDIEvent) SetData(v [3]uint8) {
	*(*[3]uint8)(unsafe.Pointer(&s.storage[21])) = v
}

// AUMIDIEventList
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUMIDIEventList
type AUMIDIEventList struct {
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
	storage [296]byte
}

// Next returns the Next field from the record's packed storage.
func (s *AUMIDIEventList) Next() *AURenderEvent {
	return *(**AURenderEvent)(unsafe.Pointer(&s.storage[0]))
}

// SetNext updates the Next field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *AUMIDIEventList) SetNext(v *AURenderEvent) {
	*(**AURenderEvent)(unsafe.Pointer(&s.storage[0])) = v
}

// EventSampleTime returns the EventSampleTime field from the record's packed storage.
func (s *AUMIDIEventList) EventSampleTime() AUEventSampleTime {
	return AUEventSampleTime(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetEventSampleTime updates the EventSampleTime field in the record's packed storage.
func (s *AUMIDIEventList) SetEventSampleTime(v AUEventSampleTime) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// EventType returns the EventType field from the record's packed storage.
func (s *AUMIDIEventList) EventType() AURenderEventType {
	return *(*AURenderEventType)(unsafe.Pointer(&s.storage[16]))
}

// SetEventType updates the EventType field in the record's packed storage.
func (s *AUMIDIEventList) SetEventType(v AURenderEventType) {
	*(*AURenderEventType)(unsafe.Pointer(&s.storage[16])) = v
}

// Reserved returns the Reserved field from the record's packed storage.
func (s *AUMIDIEventList) Reserved() uint8 {
	return uint8(s.storage[17])
}

// SetReserved updates the Reserved field in the record's packed storage.
func (s *AUMIDIEventList) SetReserved(v uint8) {
	s.storage[17] = uint8(v)
}

// Cable returns the Cable field from the record's packed storage.
func (s *AUMIDIEventList) Cable() uint8 {
	return uint8(s.storage[18])
}

// SetCable updates the Cable field in the record's packed storage.
func (s *AUMIDIEventList) SetCable(v uint8) {
	s.storage[18] = uint8(v)
}

// EventList returns the EventList field from the record's packed storage.
func (s *AUMIDIEventList) EventList() coremidi.MIDIEventList {
	return *(*coremidi.MIDIEventList)(unsafe.Pointer(&s.storage[20]))
}

// SetEventList updates the EventList field in the record's packed storage.
func (s *AUMIDIEventList) SetEventList(v coremidi.MIDIEventList) {
	*(*coremidi.MIDIEventList)(unsafe.Pointer(&s.storage[20])) = v
}

// AUMIDIOutputCallbackStruct - The callback function and custom data for an audio unit that provides MIDI output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUMIDIOutputCallbackStruct
type AUMIDIOutputCallbackStruct struct {
	MidiOutputCallback AUMIDIOutputCallback // The callback function for an audio unit that provides MIDI output.
	UserData           unsafe.Pointer       // Custom data for an audio unit that provides MIDI output.

}

// AUNodeInteraction - Describes the interaction between two node objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUNodeInteraction
type AUNodeInteraction struct {
	NodeInteractionType uint32    // The interaction type.
	NodeInteraction     [3]uint64 // A union providing information about a node interaction.

}

// AUNodeRenderCallback - A callback used to provide input to an audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUNodeRenderCallback
type AUNodeRenderCallback struct {
	DestNode        AUNode
	DestInputNumber AudioUnitElement
	Cback           AURenderCallbackStruct
}

// AUNumVersion
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUNumVersion
type AUNumVersion struct {
	NonRelRev      uint8
	Stage          uint8
	MinorAndBugRev uint8
	MajorRev       uint8
}

// AUParameterAutomationEvent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterAutomationEvent
type AUParameterAutomationEvent struct {
	HostTime  uint64
	Address   AUParameterAddress
	Value     AUValue
	EventType AUParameterAutomationEventType
	Reserved  uint64
}

// AUParameterEvent - A structure that describes a scheduled parameter event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterEvent
type AUParameterEvent struct {
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
	storage [36]byte
}

// Next returns the Next field from the record's packed storage.
func (s *AUParameterEvent) Next() *AURenderEvent {
	return *(**AURenderEvent)(unsafe.Pointer(&s.storage[0]))
}

// SetNext updates the Next field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *AUParameterEvent) SetNext(v *AURenderEvent) {
	*(**AURenderEvent)(unsafe.Pointer(&s.storage[0])) = v
}

// EventSampleTime returns the EventSampleTime field from the record's packed storage.
func (s *AUParameterEvent) EventSampleTime() AUEventSampleTime {
	return AUEventSampleTime(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetEventSampleTime updates the EventSampleTime field in the record's packed storage.
func (s *AUParameterEvent) SetEventSampleTime(v AUEventSampleTime) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// EventType returns the EventType field from the record's packed storage.
func (s *AUParameterEvent) EventType() AURenderEventType {
	return *(*AURenderEventType)(unsafe.Pointer(&s.storage[16]))
}

// SetEventType updates the EventType field in the record's packed storage.
func (s *AUParameterEvent) SetEventType(v AURenderEventType) {
	*(*AURenderEventType)(unsafe.Pointer(&s.storage[16])) = v
}

// Reserved returns the Reserved field from the record's packed storage.
func (s *AUParameterEvent) Reserved() [3]uint8 {
	return *(*[3]uint8)(unsafe.Pointer(&s.storage[17]))
}

// SetReserved updates the Reserved field in the record's packed storage.
func (s *AUParameterEvent) SetReserved(v [3]uint8) {
	*(*[3]uint8)(unsafe.Pointer(&s.storage[17])) = v
}

// RampDurationSampleFrames returns the RampDurationSampleFrames field from the record's packed storage.
func (s *AUParameterEvent) RampDurationSampleFrames() AUAudioFrameCount {
	return AUAudioFrameCount(binary.NativeEndian.Uint32(s.storage[20:24]))
}

// SetRampDurationSampleFrames updates the RampDurationSampleFrames field in the record's packed storage.
func (s *AUParameterEvent) SetRampDurationSampleFrames(v AUAudioFrameCount) {
	binary.NativeEndian.PutUint32(s.storage[20:24], uint32(v))
}

// ParameterAddress returns the ParameterAddress field from the record's packed storage.
func (s *AUParameterEvent) ParameterAddress() AUParameterAddress {
	return AUParameterAddress(binary.NativeEndian.Uint64(s.storage[24:32]))
}

// SetParameterAddress updates the ParameterAddress field in the record's packed storage.
func (s *AUParameterEvent) SetParameterAddress(v AUParameterAddress) {
	binary.NativeEndian.PutUint64(s.storage[24:32], uint64(v))
}

// Value returns the Value field from the record's packed storage.
func (s *AUParameterEvent) Value() AUValue {
	return AUValue(math.Float32frombits(binary.NativeEndian.Uint32(s.storage[32:36])))
}

// SetValue updates the Value field in the record's packed storage.
func (s *AUParameterEvent) SetValue(v AUValue) {
	binary.NativeEndian.PutUint32(s.storage[32:36], math.Float32bits(float32(v)))
}

// AUParameterMIDIMapping
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterMIDIMapping
type AUParameterMIDIMapping struct {
	MScope       AudioUnitScope
	MElement     AudioUnitElement
	MParameterID AudioUnitParameterID
	MFlags       AUParameterMIDIMappingFlags
	MSubRangeMin AudioUnitParameterValue
	MSubRangeMax AudioUnitParameterValue
	MStatus      uint8
	MData1       uint8
	Reserved1    uint8
	Reserved2    uint8
	Reserved3    uint32
}

// AUPreset - Used to set factory presets for an audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUPreset
type AUPreset struct {
	PresetNumber int32                      // If less than `0`, then the preset is a user preset. If greater than or equal to `0`, then this field is used to select a factory preset.
	PresetName   corefoundation.CFStringRef // If a factory preset, the name of the specified factory preset.

}

// AUPresetEvent - Describes an audio unit preset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUPresetEvent
type AUPresetEvent struct {
	Scope   AudioUnitScope
	Element AudioUnitElement
	Preset  corefoundation.CFPropertyListRef
}

// AURecordedParameterEvent - An event recording the changing of a parameter at a particular host time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AURecordedParameterEvent
type AURecordedParameterEvent struct {
	HostTime uint64             // The host time at which the event occurred.
	Address  AUParameterAddress // The address of the parameter whose value changed.
	Value    AUValue            // The value of the parameter at the given time.

}

// AURenderCallbackStruct - Used for registering an input callback function with an audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AURenderCallbackStruct
type AURenderCallbackStruct struct {
	InputProc       AURenderCallback
	InputProcRefCon unsafe.Pointer
}

// AURenderEvent
type AURenderEvent struct {
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
	storage [296]byte
}

// Head returns the Head field from the record's packed storage.
func (s *AURenderEvent) Head() AURenderEventHeader {
	return *(*AURenderEventHeader)(unsafe.Pointer(&s.storage[0]))
}

// SetHead updates the Head field in the record's packed storage.
func (s *AURenderEvent) SetHead(v AURenderEventHeader) {
	*(*AURenderEventHeader)(unsafe.Pointer(&s.storage[0])) = v
}

// Parameter returns the Parameter field from the record's packed storage.
func (s *AURenderEvent) Parameter() AUParameterEvent {
	return *(*AUParameterEvent)(unsafe.Pointer(&s.storage[0]))
}

// SetParameter updates the Parameter field in the record's packed storage.
func (s *AURenderEvent) SetParameter(v AUParameterEvent) {
	*(*AUParameterEvent)(unsafe.Pointer(&s.storage[0])) = v
}

// MIDI returns the MIDI field from the record's packed storage.
func (s *AURenderEvent) MIDI() AUMIDIEvent {
	return *(*AUMIDIEvent)(unsafe.Pointer(&s.storage[0]))
}

// SetMIDI updates the MIDI field in the record's packed storage.
func (s *AURenderEvent) SetMIDI(v AUMIDIEvent) {
	*(*AUMIDIEvent)(unsafe.Pointer(&s.storage[0])) = v
}

// MIDIEventsList returns the MIDIEventsList field from the record's packed storage.
func (s *AURenderEvent) MIDIEventsList() AUMIDIEventList {
	return *(*AUMIDIEventList)(unsafe.Pointer(&s.storage[0]))
}

// SetMIDIEventsList updates the MIDIEventsList field in the record's packed storage.
func (s *AURenderEvent) SetMIDIEventsList(v AUMIDIEventList) {
	*(*AUMIDIEventList)(unsafe.Pointer(&s.storage[0])) = v
}

// AURenderEventHeader - The common header for a render event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AURenderEventHeader
type AURenderEventHeader struct {
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
	storage [20]byte
}

// Next returns the Next field from the record's packed storage.
func (s *AURenderEventHeader) Next() *AURenderEvent {
	return *(**AURenderEvent)(unsafe.Pointer(&s.storage[0]))
}

// SetNext updates the Next field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *AURenderEventHeader) SetNext(v *AURenderEvent) {
	*(**AURenderEvent)(unsafe.Pointer(&s.storage[0])) = v
}

// EventSampleTime returns the EventSampleTime field from the record's packed storage.
func (s *AURenderEventHeader) EventSampleTime() AUEventSampleTime {
	return AUEventSampleTime(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetEventSampleTime updates the EventSampleTime field in the record's packed storage.
func (s *AURenderEventHeader) SetEventSampleTime(v AUEventSampleTime) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// EventType returns the EventType field from the record's packed storage.
func (s *AURenderEventHeader) EventType() AURenderEventType {
	return *(*AURenderEventType)(unsafe.Pointer(&s.storage[16]))
}

// SetEventType updates the EventType field in the record's packed storage.
func (s *AURenderEventHeader) SetEventType(v AURenderEventType) {
	*(*AURenderEventType)(unsafe.Pointer(&s.storage[16])) = v
}

// Reserved returns the Reserved field from the record's packed storage.
func (s *AURenderEventHeader) Reserved() uint8 {
	return uint8(s.storage[17])
}

// SetReserved updates the Reserved field in the record's packed storage.
func (s *AURenderEventHeader) SetReserved(v uint8) {
	s.storage[17] = uint8(v)
}

// AUSamplerBankPresetData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUSamplerBankPresetData
type AUSamplerBankPresetData struct {
	BankURL  corefoundation.CFURLRef
	BankMSB  uint8
	BankLSB  uint8
	PresetID uint8
	Reserved uint8
}

// AUSamplerInstrumentData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUSamplerInstrumentData
type AUSamplerInstrumentData struct {
	FileURL        corefoundation.CFURLRef
	InstrumentType uint8
	BankMSB        uint8
	BankLSB        uint8
	PresetID       uint8
}

// AUVoiceIOOtherAudioDuckingConfiguration - A structure that you use to configure ducking of other non-voice audio in a voice chat.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUVoiceIOOtherAudioDuckingConfiguration
type AUVoiceIOOtherAudioDuckingConfiguration struct {
	MEnableAdvancedDucking bool                            // A Boolean value that specifies whether to enable advanced ducking.
	MDuckingLevel          AUVoiceIOOtherAudioDuckingLevel // The ducking level of other non-voice audio.

}

// AudioBalanceFade - Describes audio left/right balance and front/back fade values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioBalanceFade
type AudioBalanceFade struct {
	MLeftRightBalance float32                            // The audio left/right balance, where -1 represents full left, 0 represents center, and +1 represents full right.
	MBackFrontFade    float32                            // The audio front/back fade, where -1 represents full rear, 0 represents center, and +1 represents full front.
	MType             AudioBalanceFadeType               // An AudioBalanceFadeType constant. max unity gain, or equal power.
	MChannelLayout    *coreaudiotypes.AudioChannelLayout // The size, in bytes, of the `mMagicCookie` parameter.

}

// AudioBytePacketTranslation - A data structure used by the [kAudioFilePropertyByteToPacket](<https://developer.apple.com/documentation/AudioToolbox/kAudioFilePropertyByteToPacket>) and [kAudioFilePropertyPacketToByte](<https://developer.apple.com/documentation/AudioToolbox/kAudioFilePropertyPacketToByte>) properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioBytePacketTranslation
type AudioBytePacketTranslation struct {
	MByte               int64                           // A byte number.
	MPacket             int64                           // A packet number.
	MByteOffsetInPacket uint32                          // A byte offset in a packet.
	MFlags              AudioBytePacketTranslationFlags // A translation flag value.

}

// AudioCodecMagicCookieInfo - A structure holding magic cookie information needed by some codecs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecMagicCookieInfo
type AudioCodecMagicCookieInfo struct {
	MMagicCookieSize uint32         // The size of the magic cookie.
	MMagicCookie     unsafe.Pointer // Generic constant pointer to the magic cookie.

}

// AudioCodecPrimeInfo - A structure specifying the number of leading and trailing empty frames to be inserted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioCodecPrimeInfo
type AudioCodecPrimeInfo struct {
	LeadingFrames  uint32 // An unsigned integer specifying the number of leading empty frames.
	TrailingFrames uint32 // An unsigned integer specifying the number of trailing empty frames.

}

// AudioComponentDescription - Identifying information for an audio component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentDescription
type AudioComponentDescription struct {
	ComponentType         uint32 // A unique 4-byte code identifying the interface for the component.
	ComponentSubType      uint32 // A 4-byte code that you can use to indicate the purpose of a component. For example, you could use `lpas` or `lowp` as a mnemonic indication that an audio unit is a low-pass filter.
	ComponentManufacturer uint32 // The unique vendor identifier, registered with Apple, for the audio component.
	ComponentFlags        uint32 // Set this value to zero.
	ComponentFlagsMask    uint32 // Set this value to zero.

}

// AudioComponentPlugInInterface
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentPlugInInterface
type AudioComponentPlugInInterface struct {
	Open  func(unsafe.Pointer, AudioComponentInstance) int32
	Close func(unsafe.Pointer) int32
	// Lookup returns a raw C code address, not a Go func value. The address is
	// not callable directly from Go; invoke it through purego.
	Lookup   func(int16) uintptr
	Reserved unsafe.Pointer
}

// AudioConverterPrimeInfo - Specifies priming information for an audio converter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterPrimeInfo
type AudioConverterPrimeInfo struct {
	LeadingFrames  uint32 // The number of leading frames of input audio data required for the converter to perform high-quality conversion.
	TrailingFrames uint32 // The number of trailing frames of input audio data required by the converter to perform high-quality conversion. Trailing frames follow, in time, the expected final input frame. Your application should be prepared to provide this number of additional input frames except when using the `kConverterPrimeMethod_None` value for the `kAudioConverterPrimeMethod` property. If no additional frames are available in the input stream (because, for example, the desired end frame is at the end of an audio file), then the audio converter synthesizes a sufficient number of silent (`0`-valued) trailing frames.

}

// AudioFileFDFTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileFDFTable
type AudioFileFDFTable struct {
	MComponentStorage   unsafe.Pointer
	MReadBytesFDF       ReadBytesFDF
	MWriteBytesFDF      WriteBytesFDF
	MReadPacketsFDF     ReadPacketsFDF
	MWritePacketsFDF    WritePacketsFDF
	MGetPropertyInfoFDF GetPropertyInfoFDF
	MGetPropertyFDF     GetPropertyFDF
	MSetPropertyFDF     SetPropertyFDF
	MCountUserDataFDF   CountUserDataFDF
	MGetUserDataSizeFDF GetUserDataSizeFDF
	MGetUserDataFDF     GetUserDataFDF
	MSetUserDataFDF     SetUserDataFDF
}

// AudioFileFDFTableExtended
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileFDFTableExtended
type AudioFileFDFTableExtended struct {
	MComponentStorage   unsafe.Pointer
	MReadBytesFDF       ReadBytesFDF
	MWriteBytesFDF      WriteBytesFDF
	MReadPacketsFDF     ReadPacketsFDF
	MWritePacketsFDF    WritePacketsFDF
	MGetPropertyInfoFDF GetPropertyInfoFDF
	MGetPropertyFDF     GetPropertyFDF
	MSetPropertyFDF     SetPropertyFDF
	MCountUserDataFDF   CountUserDataFDF
	MGetUserDataSizeFDF GetUserDataSizeFDF
	MGetUserDataFDF     GetUserDataFDF
	MSetUserDataFDF     SetUserDataFDF
	MReadPacketDataFDF  ReadPacketDataFDF
}

// AudioFileMarker - Annotates a position in an audio file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileMarker
type AudioFileMarker struct {
	MFramePosition float64                    // The frame in the file, counting from the start of the audio data.
	MName          corefoundation.CFStringRef // The name of the marker.
	MMarkerID      int32                      // A unique ID for the marker.
	MSMPTETime     AudioFile_SMPTE_Time       // The SMPTE time for this marker.
	MType          uint32                     // The marker type.
	MReserved      uint16                     // A reserved field. Set to `0`.
	MChannel       uint16                     // The channel number referred to by the marker. Set to `0` if the marker applies to all channels.

}

// AudioFileMarkerList - A list of markers associated with an audio file, including their SMPTE time type, the number of markers, and the markers themselves.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileMarkerList
type AudioFileMarkerList struct {
	MSMPTE_TimeType uint32             // The SMPTE time type of the whole list of markers in an audio file.
	MNumberMarkers  uint32             // The number of markers in the list.
	MMarkers        [1]AudioFileMarker // An array of `mNumberMarkers` elements, each of which is an audio file marker.

}

// AudioFilePacketTableInfo - Contains information about the number of valid frames in a file and where they begin and end.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFilePacketTableInfo
type AudioFilePacketTableInfo struct {
	MNumberValidFrames int64 // The number of valid frames in the file.
	MPrimingFrames     int32 // The number of invalid frames at the beginning of the file.
	MRemainderFrames   int32 // The number of invalid frames at the end of the file.

}

// AudioFileRegion - An audio file region specifies a segment of audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileRegion
type AudioFileRegion struct {
	MRegionID      uint32                     // A unique ID associated with the audio file region.
	MName          corefoundation.CFStringRef // The name of the region.
	MFlags         AudioFileRegionFlags       // Audio File Services region flags.
	MNumberMarkers uint32                     // The number of markers in the array specified in the  `mMarkers` parameter.
	MMarkers       [1]AudioFileMarker         // An array of `mNumberMarkers` elements describing where the data in the region starts.

}

// AudioFileRegionList - A list of the audio file regions in a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileRegionList
type AudioFileRegionList struct {
	MSMPTE_TimeType uint32             // The SMPTE timing scheme used in the file. See Core Audio’s `CAFFile.H()` header file for the values used here. For more information, see Core Audio Overview.
	MNumberRegions  uint32             // The number of regions in the list specified in the  `mRegions` parameter.
	MRegions        [1]AudioFileRegion // A variable length array of audio file regions.

}

// AudioFileTypeAndFormatID - A specifier for the constant[kAudioFileGlobalInfo_AvailableStreamDescriptionsForFormat](<https://developer.apple.com/documentation/AudioToolbox/kAudioFileGlobalInfo_AvailableStreamDescriptionsForFormat>).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFileTypeAndFormatID
type AudioFileTypeAndFormatID struct {
	MFileType AudioFileTypeID // A four-character code for the file type.
	MFormatID uint32          // A four-character code for the format ID such as `kAudioFormatLinearPCM`, `kAudioFormatMPEG4AAC`, and so forth. (See the `AudioFormat.H()` header file for declarations.)

}

// AudioFile_SMPTE_Time - A data structure for describing SMPTE (Society of Motion Picture and Television Engineers) time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFile_SMPTE_Time
type AudioFile_SMPTE_Time struct {
	MHours                int8   // The hours.
	MMinutes              uint8  // The minutes.
	MSeconds              uint8  // The seconds.
	MFrames               uint8  // The frames.
	MSubFrameSampleOffset uint32 // The sample offset within a frame.

}

// AudioFormatInfo - A structure that specifies an audio format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFormatInfo
type AudioFormatInfo struct {
	MASBD            coreaudiotypes.AudioStreamBasicDescription // An [AudioStreamBasicDescription] structure.
	MMagicCookie     unsafe.Pointer                             // A pointer to the decompression information for the data described in the `mASBD` parameter.
	MMagicCookieSize uint32                                     // The size, in bytes, of the `mMagicCookie` parameter.

}

// AudioFormatListItem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioFormatListItem
type AudioFormatListItem struct {
	MASBD             coreaudiotypes.AudioStreamBasicDescription
	MChannelLayoutTag uint32
}

// AudioFramePacketTranslation - A structure that specifies frame and packet translations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioFramePacketTranslation
type AudioFramePacketTranslation struct {
	MFrame               int64  // A frame number.
	MPacket              int64  // A packet number.
	MFrameOffsetInPacket uint32 // A frame offset in a packet.

}

// AudioIndependentPacketTranslation
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioIndependentPacketTranslation
type AudioIndependentPacketTranslation struct {
	MPacket                       int64
	MIndependentlyDecodablePacket int64
}

// AudioOutputUnitMIDICallbacks
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioOutputUnitMIDICallbacks
type AudioOutputUnitMIDICallbacks struct {
	UserData      unsafe.Pointer
	MIDIEventProc func(unsafe.Pointer, uint32, uint32, uint32, uint32)
	MIDISysExProc func(unsafe.Pointer, *byte, uint32)
}

// AudioOutputUnitStartAtTimeParams - A timestamp for scheduled starting of an I/O audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioOutputUnitStartAtTimeParams
type AudioOutputUnitStartAtTimeParams struct {
	MTimestamp coreaudiotypes.AudioTimeStamp
	MFlags     uint32
}

// AudioPacketDependencyInfoTranslation
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioPacketDependencyInfoTranslation
type AudioPacketDependencyInfoTranslation struct {
	MPacket                   int64
	MIsIndependentlyDecodable uint32
	MNumberPrerollPackets     uint32
}

// AudioPacketRangeByteCountTranslation
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioPacketRangeByteCountTranslation
type AudioPacketRangeByteCountTranslation struct {
	MPacket              int64
	MPacketCount         int64
	MByteCountUpperBound int64
}

// AudioPacketRollDistanceTranslation
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioPacketRollDistanceTranslation
type AudioPacketRollDistanceTranslation struct {
	MPacket       int64
	MRollDistance int64
}

// AudioPanningInfo - Audio panning information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioPanningInfo
type AudioPanningInfo struct {
	MPanningMode      AudioPanningMode                   // The mode to use for panning.
	MCoordinateFlags  uint32                             // For the available coordinate flags, see Channel Coordinate Flags.
	MCoordinates      [3]float32                         // For the available coordinate index constants, see Channel Coordinate Index Constants.
	MGainScale        float32                            // A multiplier for audio panning values, typically representing a volume value in the range from 0 to 1. A value of 1 results in audio panning at unity gain. A value of 0 silences all channels.
	MOutputChannelMap *coreaudiotypes.AudioChannelLayout // The channel map used to determine channel volumes for the audio panning.

}

// AudioQueueBuffer - Defines an audio queue buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueBuffer
type AudioQueueBuffer struct {
	MAudioDataBytesCapacity    uint32                                       // The size of the audio queue buffer, in bytes. This size is set when a buffer is allocated and cannot be changed.
	MAudioData                 unsafe.Pointer                               // The audio data owned the audio queue buffer. The buffer address cannot be changed.
	MAudioDataByteSize         uint32                                       // The number of bytes of valid audio data in the audio queue buffer’s `mAudioData` field, initially set to `0`. Your callback must set this value for a playback audio queue; for recording, the recording audio queue sets the value.
	MUserData                  unsafe.Pointer                               // The custom data structure you specify, for use by your callback function, when creating a recording or playback audio queue.
	MPacketDescriptionCapacity uint32                                       // The maximum number of packet descriptions that can be stored in the `mPacketDescriptions` field.
	MPacketDescriptions        *coreaudiotypes.AudioStreamPacketDescription // An array of [AudioStreamPacketDescription] structures for the buffer.
	MPacketDescriptionCount    uint32                                       // The number of valid packet descriptions in the buffer. You set this value when providing buffers for playback. The audio queue sets this value when returning buffers from a recording queue.

}

// AudioQueueChannelAssignment
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueChannelAssignment
type AudioQueueChannelAssignment struct {
	MDeviceUID     corefoundation.CFStringRef
	MChannelNumber uint32
}

// AudioQueueLevelMeterState - Specifies the current level metering information for one channel of an audio queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueLevelMeterState
type AudioQueueLevelMeterState struct {
	MAveragePower float32 // The audio channel’s average RMS power.
	MPeakPower    float32 // The audio channel’s peak RMS power.

}

// AudioQueueParameterEvent - Specifies an audio queue parameter and associated value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueParameterEvent
type AudioQueueParameterEvent struct {
	MID    AudioQueueParameterID    // The parameter.
	MValue AudioQueueParameterValue // The value of the specified parameter.

}

// AudioUnitCocoaViewInfo - The name and number of custom Cocoa views for an audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitCocoaViewInfo
type AudioUnitCocoaViewInfo struct {
	MCocoaAUViewBundleLocation corefoundation.CFURLRef
	MCocoaAUViewClass          [1]corefoundation.CFStringRef
}

// AudioUnitConnection - An audio unit source-to-destination connection specification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitConnection
type AudioUnitConnection struct {
	SourceAudioUnit    AudioUnit // The audio unit that is serves as the source in the connection.
	SourceOutputNumber uint32    // The source audio unit’s output element to be used in the connection.
	DestInputNumber    uint32    // The destination audio unit’s input element to be used in the connection.

}

// AudioUnitEvent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitEvent
type AudioUnitEvent struct {
	MEventType AudioUnitEventType
	MArgument  [3]uint64
}

// AudioUnitExternalBuffer - Allows an audio unit host application to tell an audio unit to use a specified buffer for its input callback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitExternalBuffer
type AudioUnitExternalBuffer struct {
	Buffer *uint8
	Size   uint32
}

// AudioUnitFrequencyResponseBin - An audio unit’s audio level at a particular frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitFrequencyResponseBin
type AudioUnitFrequencyResponseBin struct {
	MFrequency float64
	MMagnitude float64
}

// AudioUnitMIDIControlMapping
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitMIDIControlMapping
type AudioUnitMIDIControlMapping struct {
	MidiNRPN    uint16
	MidiControl uint8
	Scope       uint8
	Element     AudioUnitElement
	Parameter   AudioUnitParameterID
}

// AudioUnitMeterClipping - Audio clipping that has occurred in a mixer unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitMeterClipping
type AudioUnitMeterClipping struct {
	PeakValueSinceLastCall float32 // The maximum value seen on the channel since the last time the property was retrieved.
	SawInfinity            bool    // [TRUE] if there was an infinite value on this channel since the last time the property was retrieved.
	SawNotANumber          bool    // [TRUE] if there was a floating point “not a number” value on this channel since the last time the property was retrieved.

}

// AudioUnitNodeConnection - A connection between two node objects in an audio processing graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitNodeConnection
type AudioUnitNodeConnection struct {
	SourceNode         AUNode
	SourceOutputNumber uint32
	DestNode           AUNode
	DestInputNumber    uint32
}

// AudioUnitOtherPluginDesc
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitOtherPluginDesc
type AudioUnitOtherPluginDesc struct {
	Format uint32
	Plugin coreaudiotypes.AudioClassDescription
}

// AudioUnitParameter - An adjustable audio unit attribute such as volume, pitch, or filter cutoff frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameter
type AudioUnitParameter struct {
	MAudioUnit   AudioUnit            // The audio unit instance that the parameter applies to.
	MParameterID AudioUnitParameterID // The audio unit parameter identifier.
	MScope       AudioUnitScope       // The audio unit scope for the parameter.
	MElement     AudioUnitElement     // The audio unit element for the parameter.

}

// AudioUnitParameterEvent - A scheduled change to an audio unit parameter’s value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterEvent
type AudioUnitParameterEvent struct {
	Scope       AudioUnitScope       // The scope for this parameter event.
	Element     AudioUnitElement     // The element for this parameter event.
	Parameter   AudioUnitParameterID // An identifier for this parameter event.
	EventType   AUParameterEventType // The type for this parameter event.
	EventValues [4]uint32            // The values for this parameter event.
	Immediate   unsafe.Pointer
	Ramp        unsafe.Pointer
}

// AudioUnitParameterHistoryInfo - The suggested update rate and history duration for parameters which have the [flag_PlotHistory](<https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterOptions/flag_PlotHistory>) flag set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterHistoryInfo
type AudioUnitParameterHistoryInfo struct {
	UpdatesPerSecond         float32
	HistoryDurationInSeconds float32
}

// AudioUnitParameterInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterInfo
type AudioUnitParameterInfo struct {
	Name         [52]int8                   // Must be set to `0`.
	UnitName     corefoundation.CFStringRef // If `kAudioUnitParameterUnit_CustomUnit` is set, this field must contain a valid [CFString] object. Only valid if `kAudioUnitParameterUnit_CustomUnit` is set.
	ClumpID      uint32                     // Only valid if `kAudioUnitParameterFlag_HasClump` is set.
	CfNameString corefoundation.CFStringRef // Only valid if `kAudioUnitParameterFlag_HasCFNameString` is set.
	Unit         AudioUnitParameterUnit     // If the `unit` field contains a value not in the [AudioUnitParameterUnit] enumeration, then assume the unit type is `kAudioUnitParameterUnit_Generic`.
	MinValue     AudioUnitParameterValue
	MaxValue     AudioUnitParameterValue
	DefaultValue AudioUnitParameterValue
	Flags        AudioUnitParameterOptions // The host should check for this flag and, if present, release the parameter name when it is finished with it.

}

// AudioUnitParameterNameInfo - A short version of the name for an audio unit parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterNameInfo
type AudioUnitParameterNameInfo struct {
	InID            AudioUnitParameterID       // The identifier for the audio unit parameter.
	InDesiredLength int32                      // When setting an audio unit property that uses this data structure for its value, the maximum length that you are specifying for the audio unit parameter name.
	OutName         corefoundation.CFStringRef // When getting an audio unit property that uses this data structure for its value, the short version of the parameter name provided by the audio unit. The host application then owns the string and is responsible for releasing it.

}

// AudioUnitParameterStringFromValue - A string representation of a parameter’s value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterStringFromValue
type AudioUnitParameterStringFromValue struct {
	InParamID AudioUnitParameterID
	InValue   *AudioUnitParameterValue
	OutString corefoundation.CFStringRef
}

// AudioUnitParameterValueFromString - A parameter’s value based on a string representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterValueFromString
type AudioUnitParameterValueFromString struct {
	InParamID AudioUnitParameterID
	InString  corefoundation.CFStringRef
	OutValue  AudioUnitParameterValue
}

// AudioUnitParameterValueName
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterValueName
type AudioUnitParameterValueName struct {
	InParamID AudioUnitParameterID
	InValue   *float32
	OutName   corefoundation.CFStringRef
}

// AudioUnitParameterValueTranslation
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterValueTranslation
type AudioUnitParameterValueTranslation struct {
	OtherDesc    AudioUnitOtherPluginDesc
	OtherParamID uint32
	OtherValue   float32
	AuParamID    AudioUnitParameterID
	AuValue      AudioUnitParameterValue
}

// AudioUnitPresetMAS_SettingData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitPresetMAS_SettingData
type AudioUnitPresetMAS_SettingData struct {
	IsStockSetting uint32
	SettingID      uint32
	DataLen        uint32
	Data           [1]uint8
}

// AudioUnitPresetMAS_Settings
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitPresetMAS_Settings
type AudioUnitPresetMAS_Settings struct {
	ManufacturerID   uint32
	EffectID         uint32
	VariantID        uint32
	SettingsVersion  uint32
	NumberOfSettings uint32
	Settings         [1]AudioUnitPresetMAS_SettingData
}

// AudioUnitProperty - A key-value pair that declares an attribute or behavior for an audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitProperty
type AudioUnitProperty struct {
	MAudioUnit  AudioUnit           // The audio unit instance that the parameter applies to.
	MPropertyID AudioUnitPropertyID // The audio unit property identifier.
	MScope      AudioUnitScope      // The audio unit scope for the property.
	MElement    AudioUnitElement    // The audio unit element for the property.

}

// AudioUnitRenderContext - A structure that contains thread context information for a real-time rendering operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitRenderContext
type AudioUnitRenderContext struct {
	Workgroup os.Os_workgroup_t // The workgroup that manages the rendering threads of the audio unit.
	Reserved  [6]uint32         // System-specific information.

}

// CABarBeatTime
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CABarBeatTime
type CABarBeatTime struct {
	Bar            int32
	Beat           uint16
	Subbeat        uint16
	SubbeatDivisor uint16
	Reserved       uint16
}

// CAClockTime
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAClockTime
type CAClockTime struct {
	Format   CAClockTimeFormat
	Reserved uint32
	Time     [3]uint64
}

// CAFAudioDescription
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFAudioDescription
type CAFAudioDescription struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [32]byte
}

// MSampleRate returns the MSampleRate field from the record's packed storage.
func (s *CAFAudioDescription) MSampleRate() float64 {
	return math.Float64frombits(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetMSampleRate updates the MSampleRate field in the record's packed storage.
func (s *CAFAudioDescription) SetMSampleRate(v float64) {
	binary.NativeEndian.PutUint64(s.storage[0:8], math.Float64bits(v))
}

// MFormatID returns the MFormatID field from the record's packed storage.
func (s *CAFAudioDescription) MFormatID() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetMFormatID updates the MFormatID field in the record's packed storage.
func (s *CAFAudioDescription) SetMFormatID(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// MFormatFlags returns the MFormatFlags field from the record's packed storage.
func (s *CAFAudioDescription) MFormatFlags() CAFFormatFlags {
	return *(*CAFFormatFlags)(unsafe.Pointer(&s.storage[12]))
}

// SetMFormatFlags updates the MFormatFlags field in the record's packed storage.
func (s *CAFAudioDescription) SetMFormatFlags(v CAFFormatFlags) {
	*(*CAFFormatFlags)(unsafe.Pointer(&s.storage[12])) = v
}

// MBytesPerPacket returns the MBytesPerPacket field from the record's packed storage.
func (s *CAFAudioDescription) MBytesPerPacket() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[16:20]))
}

// SetMBytesPerPacket updates the MBytesPerPacket field in the record's packed storage.
func (s *CAFAudioDescription) SetMBytesPerPacket(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[16:20], uint32(v))
}

// MFramesPerPacket returns the MFramesPerPacket field from the record's packed storage.
func (s *CAFAudioDescription) MFramesPerPacket() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[20:24]))
}

// SetMFramesPerPacket updates the MFramesPerPacket field in the record's packed storage.
func (s *CAFAudioDescription) SetMFramesPerPacket(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[20:24], uint32(v))
}

// MChannelsPerFrame returns the MChannelsPerFrame field from the record's packed storage.
func (s *CAFAudioDescription) MChannelsPerFrame() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[24:28]))
}

// SetMChannelsPerFrame updates the MChannelsPerFrame field in the record's packed storage.
func (s *CAFAudioDescription) SetMChannelsPerFrame(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[24:28], uint32(v))
}

// MBitsPerChannel returns the MBitsPerChannel field from the record's packed storage.
func (s *CAFAudioDescription) MBitsPerChannel() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[28:32]))
}

// SetMBitsPerChannel updates the MBitsPerChannel field in the record's packed storage.
func (s *CAFAudioDescription) SetMBitsPerChannel(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[28:32], uint32(v))
}

// CAFAudioFormatListItem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFAudioFormatListItem
type CAFAudioFormatListItem struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [36]byte
}

// MFormat returns the MFormat field from the record's packed storage.
func (s *CAFAudioFormatListItem) MFormat() CAFAudioDescription {
	return *(*CAFAudioDescription)(unsafe.Pointer(&s.storage[0]))
}

// SetMFormat updates the MFormat field in the record's packed storage.
func (s *CAFAudioFormatListItem) SetMFormat(v CAFAudioDescription) {
	*(*CAFAudioDescription)(unsafe.Pointer(&s.storage[0])) = v
}

// MChannelLayoutTag returns the MChannelLayoutTag field from the record's packed storage.
func (s *CAFAudioFormatListItem) MChannelLayoutTag() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[32:36]))
}

// SetMChannelLayoutTag updates the MChannelLayoutTag field in the record's packed storage.
func (s *CAFAudioFormatListItem) SetMChannelLayoutTag(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[32:36], uint32(v))
}

// CAFChunkHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFChunkHeader
type CAFChunkHeader struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [12]byte
}

// MChunkType returns the MChunkType field from the record's packed storage.
func (s *CAFChunkHeader) MChunkType() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetMChunkType updates the MChunkType field in the record's packed storage.
func (s *CAFChunkHeader) SetMChunkType(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// MChunkSize returns the MChunkSize field from the record's packed storage.
func (s *CAFChunkHeader) MChunkSize() int64 {
	return int64(binary.NativeEndian.Uint64(s.storage[4:12]))
}

// SetMChunkSize updates the MChunkSize field in the record's packed storage.
func (s *CAFChunkHeader) SetMChunkSize(v int64) {
	binary.NativeEndian.PutUint64(s.storage[4:12], uint64(v))
}

// CAFDataChunk
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFDataChunk
type CAFDataChunk struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [5]byte
}

// MEditCount returns the MEditCount field from the record's packed storage.
func (s *CAFDataChunk) MEditCount() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetMEditCount updates the MEditCount field in the record's packed storage.
func (s *CAFDataChunk) SetMEditCount(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// MData returns the MData field from the record's packed storage.
func (s *CAFDataChunk) MData() [1]uint8 {
	return *(*[1]uint8)(unsafe.Pointer(&s.storage[4]))
}

// SetMData updates the MData field in the record's packed storage.
func (s *CAFDataChunk) SetMData(v [1]uint8) {
	*(*[1]uint8)(unsafe.Pointer(&s.storage[4])) = v
}

// CAFFileHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFFileHeader
type CAFFileHeader struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [8]byte
}

// MFileType returns the MFileType field from the record's packed storage.
func (s *CAFFileHeader) MFileType() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetMFileType updates the MFileType field in the record's packed storage.
func (s *CAFFileHeader) SetMFileType(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// MFileVersion returns the MFileVersion field from the record's packed storage.
func (s *CAFFileHeader) MFileVersion() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[4:6]))
}

// SetMFileVersion updates the MFileVersion field in the record's packed storage.
func (s *CAFFileHeader) SetMFileVersion(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[4:6], uint16(v))
}

// MFileFlags returns the MFileFlags field from the record's packed storage.
func (s *CAFFileHeader) MFileFlags() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[6:8]))
}

// SetMFileFlags updates the MFileFlags field in the record's packed storage.
func (s *CAFFileHeader) SetMFileFlags(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[6:8], uint16(v))
}

// CAFInfoStrings
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFInfoStrings
type CAFInfoStrings struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [4]byte
}

// MNumEntries returns the MNumEntries field from the record's packed storage.
func (s *CAFInfoStrings) MNumEntries() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetMNumEntries updates the MNumEntries field in the record's packed storage.
func (s *CAFInfoStrings) SetMNumEntries(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// CAFInstrumentChunk
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFInstrumentChunk
type CAFInstrumentChunk struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [28]byte
}

// MBaseNote returns the MBaseNote field from the record's packed storage.
func (s *CAFInstrumentChunk) MBaseNote() float32 {
	return math.Float32frombits(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetMBaseNote updates the MBaseNote field in the record's packed storage.
func (s *CAFInstrumentChunk) SetMBaseNote(v float32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], math.Float32bits(v))
}

// MMIDILowNote returns the MMIDILowNote field from the record's packed storage.
func (s *CAFInstrumentChunk) MMIDILowNote() uint8 {
	return uint8(s.storage[4])
}

// SetMMIDILowNote updates the MMIDILowNote field in the record's packed storage.
func (s *CAFInstrumentChunk) SetMMIDILowNote(v uint8) {
	s.storage[4] = uint8(v)
}

// MMIDIHighNote returns the MMIDIHighNote field from the record's packed storage.
func (s *CAFInstrumentChunk) MMIDIHighNote() uint8 {
	return uint8(s.storage[5])
}

// SetMMIDIHighNote updates the MMIDIHighNote field in the record's packed storage.
func (s *CAFInstrumentChunk) SetMMIDIHighNote(v uint8) {
	s.storage[5] = uint8(v)
}

// MMIDILowVelocity returns the MMIDILowVelocity field from the record's packed storage.
func (s *CAFInstrumentChunk) MMIDILowVelocity() uint8 {
	return uint8(s.storage[6])
}

// SetMMIDILowVelocity updates the MMIDILowVelocity field in the record's packed storage.
func (s *CAFInstrumentChunk) SetMMIDILowVelocity(v uint8) {
	s.storage[6] = uint8(v)
}

// MMIDIHighVelocity returns the MMIDIHighVelocity field from the record's packed storage.
func (s *CAFInstrumentChunk) MMIDIHighVelocity() uint8 {
	return uint8(s.storage[7])
}

// SetMMIDIHighVelocity updates the MMIDIHighVelocity field in the record's packed storage.
func (s *CAFInstrumentChunk) SetMMIDIHighVelocity(v uint8) {
	s.storage[7] = uint8(v)
}

// MdBGain returns the MdBGain field from the record's packed storage.
func (s *CAFInstrumentChunk) MdBGain() float32 {
	return math.Float32frombits(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetMdBGain updates the MdBGain field in the record's packed storage.
func (s *CAFInstrumentChunk) SetMdBGain(v float32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], math.Float32bits(v))
}

// MStartRegionID returns the MStartRegionID field from the record's packed storage.
func (s *CAFInstrumentChunk) MStartRegionID() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetMStartRegionID updates the MStartRegionID field in the record's packed storage.
func (s *CAFInstrumentChunk) SetMStartRegionID(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// MSustainRegionID returns the MSustainRegionID field from the record's packed storage.
func (s *CAFInstrumentChunk) MSustainRegionID() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[16:20]))
}

// SetMSustainRegionID updates the MSustainRegionID field in the record's packed storage.
func (s *CAFInstrumentChunk) SetMSustainRegionID(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[16:20], uint32(v))
}

// MReleaseRegionID returns the MReleaseRegionID field from the record's packed storage.
func (s *CAFInstrumentChunk) MReleaseRegionID() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[20:24]))
}

// SetMReleaseRegionID updates the MReleaseRegionID field in the record's packed storage.
func (s *CAFInstrumentChunk) SetMReleaseRegionID(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[20:24], uint32(v))
}

// MInstrumentID returns the MInstrumentID field from the record's packed storage.
func (s *CAFInstrumentChunk) MInstrumentID() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[24:28]))
}

// SetMInstrumentID updates the MInstrumentID field in the record's packed storage.
func (s *CAFInstrumentChunk) SetMInstrumentID(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[24:28], uint32(v))
}

// CAFMarker
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFMarker
type CAFMarker struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [28]byte
}

// MType returns the MType field from the record's packed storage.
func (s *CAFMarker) MType() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetMType updates the MType field in the record's packed storage.
func (s *CAFMarker) SetMType(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// MFramePosition returns the MFramePosition field from the record's packed storage.
func (s *CAFMarker) MFramePosition() float64 {
	return math.Float64frombits(binary.NativeEndian.Uint64(s.storage[4:12]))
}

// SetMFramePosition updates the MFramePosition field in the record's packed storage.
func (s *CAFMarker) SetMFramePosition(v float64) {
	binary.NativeEndian.PutUint64(s.storage[4:12], math.Float64bits(v))
}

// MMarkerID returns the MMarkerID field from the record's packed storage.
func (s *CAFMarker) MMarkerID() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetMMarkerID updates the MMarkerID field in the record's packed storage.
func (s *CAFMarker) SetMMarkerID(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// MSMPTETime returns the MSMPTETime field from the record's packed storage.
func (s *CAFMarker) MSMPTETime() CAF_SMPTE_Time {
	return *(*CAF_SMPTE_Time)(unsafe.Pointer(&s.storage[16]))
}

// SetMSMPTETime updates the MSMPTETime field in the record's packed storage.
func (s *CAFMarker) SetMSMPTETime(v CAF_SMPTE_Time) {
	*(*CAF_SMPTE_Time)(unsafe.Pointer(&s.storage[16])) = v
}

// MChannel returns the MChannel field from the record's packed storage.
func (s *CAFMarker) MChannel() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[24:28]))
}

// SetMChannel updates the MChannel field in the record's packed storage.
func (s *CAFMarker) SetMChannel(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[24:28], uint32(v))
}

// CAFMarkerChunk
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFMarkerChunk
type CAFMarkerChunk struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [36]byte
}

// MSMPTE_TimeType returns the MSMPTE_TimeType field from the record's packed storage.
func (s *CAFMarkerChunk) MSMPTE_TimeType() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetMSMPTE_TimeType updates the MSMPTE_TimeType field in the record's packed storage.
func (s *CAFMarkerChunk) SetMSMPTE_TimeType(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// MNumberMarkers returns the MNumberMarkers field from the record's packed storage.
func (s *CAFMarkerChunk) MNumberMarkers() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetMNumberMarkers updates the MNumberMarkers field in the record's packed storage.
func (s *CAFMarkerChunk) SetMNumberMarkers(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// MMarkers returns the MMarkers field from the record's packed storage.
func (s *CAFMarkerChunk) MMarkers() [1]CAFMarker {
	return *(*[1]CAFMarker)(unsafe.Pointer(&s.storage[8]))
}

// SetMMarkers updates the MMarkers field in the record's packed storage.
func (s *CAFMarkerChunk) SetMMarkers(v [1]CAFMarker) {
	*(*[1]CAFMarker)(unsafe.Pointer(&s.storage[8])) = v
}

// CAFOverviewChunk
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFOverviewChunk
type CAFOverviewChunk struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [12]byte
}

// MEditCount returns the MEditCount field from the record's packed storage.
func (s *CAFOverviewChunk) MEditCount() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetMEditCount updates the MEditCount field in the record's packed storage.
func (s *CAFOverviewChunk) SetMEditCount(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// MNumFramesPerOVWSample returns the MNumFramesPerOVWSample field from the record's packed storage.
func (s *CAFOverviewChunk) MNumFramesPerOVWSample() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetMNumFramesPerOVWSample updates the MNumFramesPerOVWSample field in the record's packed storage.
func (s *CAFOverviewChunk) SetMNumFramesPerOVWSample(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// MData returns the MData field from the record's packed storage.
func (s *CAFOverviewChunk) MData() [1]CAFOverviewSample {
	return *(*[1]CAFOverviewSample)(unsafe.Pointer(&s.storage[8]))
}

// SetMData updates the MData field in the record's packed storage.
func (s *CAFOverviewChunk) SetMData(v [1]CAFOverviewSample) {
	*(*[1]CAFOverviewSample)(unsafe.Pointer(&s.storage[8])) = v
}

// CAFOverviewSample
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFOverviewSample
type CAFOverviewSample struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [4]byte
}

// MMinValue returns the MMinValue field from the record's packed storage.
func (s *CAFOverviewSample) MMinValue() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[0:2]))
}

// SetMMinValue updates the MMinValue field in the record's packed storage.
func (s *CAFOverviewSample) SetMMinValue(v int16) {
	binary.NativeEndian.PutUint16(s.storage[0:2], uint16(v))
}

// MMaxValue returns the MMaxValue field from the record's packed storage.
func (s *CAFOverviewSample) MMaxValue() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[2:4]))
}

// SetMMaxValue updates the MMaxValue field in the record's packed storage.
func (s *CAFOverviewSample) SetMMaxValue(v int16) {
	binary.NativeEndian.PutUint16(s.storage[2:4], uint16(v))
}

// CAFPacketTableHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFPacketTableHeader
type CAFPacketTableHeader struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [25]byte
}

// MNumberPackets returns the MNumberPackets field from the record's packed storage.
func (s *CAFPacketTableHeader) MNumberPackets() int64 {
	return int64(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetMNumberPackets updates the MNumberPackets field in the record's packed storage.
func (s *CAFPacketTableHeader) SetMNumberPackets(v int64) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// MNumberValidFrames returns the MNumberValidFrames field from the record's packed storage.
func (s *CAFPacketTableHeader) MNumberValidFrames() int64 {
	return int64(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetMNumberValidFrames updates the MNumberValidFrames field in the record's packed storage.
func (s *CAFPacketTableHeader) SetMNumberValidFrames(v int64) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// MPrimingFrames returns the MPrimingFrames field from the record's packed storage.
func (s *CAFPacketTableHeader) MPrimingFrames() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[16:20]))
}

// SetMPrimingFrames updates the MPrimingFrames field in the record's packed storage.
func (s *CAFPacketTableHeader) SetMPrimingFrames(v int32) {
	binary.NativeEndian.PutUint32(s.storage[16:20], uint32(v))
}

// MRemainderFrames returns the MRemainderFrames field from the record's packed storage.
func (s *CAFPacketTableHeader) MRemainderFrames() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[20:24]))
}

// SetMRemainderFrames updates the MRemainderFrames field in the record's packed storage.
func (s *CAFPacketTableHeader) SetMRemainderFrames(v int32) {
	binary.NativeEndian.PutUint32(s.storage[20:24], uint32(v))
}

// MPacketDescriptions returns the MPacketDescriptions field from the record's packed storage.
func (s *CAFPacketTableHeader) MPacketDescriptions() [1]uint8 {
	return *(*[1]uint8)(unsafe.Pointer(&s.storage[24]))
}

// SetMPacketDescriptions updates the MPacketDescriptions field in the record's packed storage.
func (s *CAFPacketTableHeader) SetMPacketDescriptions(v [1]uint8) {
	*(*[1]uint8)(unsafe.Pointer(&s.storage[24])) = v
}

// CAFPeakChunk
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFPeakChunk
type CAFPeakChunk struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [16]byte
}

// MEditCount returns the MEditCount field from the record's packed storage.
func (s *CAFPeakChunk) MEditCount() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetMEditCount updates the MEditCount field in the record's packed storage.
func (s *CAFPeakChunk) SetMEditCount(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// MPeaks returns the MPeaks field from the record's packed storage.
func (s *CAFPeakChunk) MPeaks() [1]CAFPositionPeak {
	return *(*[1]CAFPositionPeak)(unsafe.Pointer(&s.storage[4]))
}

// SetMPeaks updates the MPeaks field in the record's packed storage.
func (s *CAFPeakChunk) SetMPeaks(v [1]CAFPositionPeak) {
	*(*[1]CAFPositionPeak)(unsafe.Pointer(&s.storage[4])) = v
}

// CAFPositionPeak
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFPositionPeak
type CAFPositionPeak struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [12]byte
}

// MValue returns the MValue field from the record's packed storage.
func (s *CAFPositionPeak) MValue() float32 {
	return math.Float32frombits(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetMValue updates the MValue field in the record's packed storage.
func (s *CAFPositionPeak) SetMValue(v float32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], math.Float32bits(v))
}

// MFrameNumber returns the MFrameNumber field from the record's packed storage.
func (s *CAFPositionPeak) MFrameNumber() uint64 {
	return uint64(binary.NativeEndian.Uint64(s.storage[4:12]))
}

// SetMFrameNumber updates the MFrameNumber field in the record's packed storage.
func (s *CAFPositionPeak) SetMFrameNumber(v uint64) {
	binary.NativeEndian.PutUint64(s.storage[4:12], uint64(v))
}

// CAFRegion
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFRegion
type CAFRegion struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [40]byte
}

// MRegionID returns the MRegionID field from the record's packed storage.
func (s *CAFRegion) MRegionID() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetMRegionID updates the MRegionID field in the record's packed storage.
func (s *CAFRegion) SetMRegionID(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// MFlags returns the MFlags field from the record's packed storage.
func (s *CAFRegion) MFlags() CAFRegionFlags {
	return *(*CAFRegionFlags)(unsafe.Pointer(&s.storage[4]))
}

// SetMFlags updates the MFlags field in the record's packed storage.
func (s *CAFRegion) SetMFlags(v CAFRegionFlags) {
	*(*CAFRegionFlags)(unsafe.Pointer(&s.storage[4])) = v
}

// MNumberMarkers returns the MNumberMarkers field from the record's packed storage.
func (s *CAFRegion) MNumberMarkers() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetMNumberMarkers updates the MNumberMarkers field in the record's packed storage.
func (s *CAFRegion) SetMNumberMarkers(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// MMarkers returns the MMarkers field from the record's packed storage.
func (s *CAFRegion) MMarkers() [1]CAFMarker {
	return *(*[1]CAFMarker)(unsafe.Pointer(&s.storage[12]))
}

// SetMMarkers updates the MMarkers field in the record's packed storage.
func (s *CAFRegion) SetMMarkers(v [1]CAFMarker) {
	*(*[1]CAFMarker)(unsafe.Pointer(&s.storage[12])) = v
}

// CAFRegionChunk
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFRegionChunk
type CAFRegionChunk struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [48]byte
}

// MSMPTE_TimeType returns the MSMPTE_TimeType field from the record's packed storage.
func (s *CAFRegionChunk) MSMPTE_TimeType() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetMSMPTE_TimeType updates the MSMPTE_TimeType field in the record's packed storage.
func (s *CAFRegionChunk) SetMSMPTE_TimeType(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// MNumberRegions returns the MNumberRegions field from the record's packed storage.
func (s *CAFRegionChunk) MNumberRegions() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetMNumberRegions updates the MNumberRegions field in the record's packed storage.
func (s *CAFRegionChunk) SetMNumberRegions(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// MRegions returns the MRegions field from the record's packed storage.
func (s *CAFRegionChunk) MRegions() [1]CAFRegion {
	return *(*[1]CAFRegion)(unsafe.Pointer(&s.storage[8]))
}

// SetMRegions updates the MRegions field in the record's packed storage.
func (s *CAFRegionChunk) SetMRegions(v [1]CAFRegion) {
	*(*[1]CAFRegion)(unsafe.Pointer(&s.storage[8])) = v
}

// CAFStringID
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFStringID
type CAFStringID struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [12]byte
}

// MStringID returns the MStringID field from the record's packed storage.
func (s *CAFStringID) MStringID() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetMStringID updates the MStringID field in the record's packed storage.
func (s *CAFStringID) SetMStringID(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// MStringStartByteOffset returns the MStringStartByteOffset field from the record's packed storage.
func (s *CAFStringID) MStringStartByteOffset() int64 {
	return int64(binary.NativeEndian.Uint64(s.storage[4:12]))
}

// SetMStringStartByteOffset updates the MStringStartByteOffset field in the record's packed storage.
func (s *CAFStringID) SetMStringStartByteOffset(v int64) {
	binary.NativeEndian.PutUint64(s.storage[4:12], uint64(v))
}

// CAFStrings
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFStrings
type CAFStrings struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [16]byte
}

// MNumEntries returns the MNumEntries field from the record's packed storage.
func (s *CAFStrings) MNumEntries() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetMNumEntries updates the MNumEntries field in the record's packed storage.
func (s *CAFStrings) SetMNumEntries(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// MStringsIDs returns the MStringsIDs field from the record's packed storage.
func (s *CAFStrings) MStringsIDs() [1]CAFStringID {
	return *(*[1]CAFStringID)(unsafe.Pointer(&s.storage[4]))
}

// SetMStringsIDs updates the MStringsIDs field in the record's packed storage.
func (s *CAFStrings) SetMStringsIDs(v [1]CAFStringID) {
	*(*[1]CAFStringID)(unsafe.Pointer(&s.storage[4])) = v
}

// CAFUMIDChunk
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFUMIDChunk
type CAFUMIDChunk struct {
	MBytes [64]uint8
}

// CAF_SMPTE_Time
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAF_SMPTE_Time
type CAF_SMPTE_Time struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [8]byte
}

// MHours returns the MHours field from the record's packed storage.
func (s *CAF_SMPTE_Time) MHours() int8 {
	return int8(s.storage[0])
}

// SetMHours updates the MHours field in the record's packed storage.
func (s *CAF_SMPTE_Time) SetMHours(v int8) {
	s.storage[0] = uint8(v)
}

// MMinutes returns the MMinutes field from the record's packed storage.
func (s *CAF_SMPTE_Time) MMinutes() int8 {
	return int8(s.storage[1])
}

// SetMMinutes updates the MMinutes field in the record's packed storage.
func (s *CAF_SMPTE_Time) SetMMinutes(v int8) {
	s.storage[1] = uint8(v)
}

// MSeconds returns the MSeconds field from the record's packed storage.
func (s *CAF_SMPTE_Time) MSeconds() int8 {
	return int8(s.storage[2])
}

// SetMSeconds updates the MSeconds field in the record's packed storage.
func (s *CAF_SMPTE_Time) SetMSeconds(v int8) {
	s.storage[2] = uint8(v)
}

// MFrames returns the MFrames field from the record's packed storage.
func (s *CAF_SMPTE_Time) MFrames() int8 {
	return int8(s.storage[3])
}

// SetMFrames updates the MFrames field in the record's packed storage.
func (s *CAF_SMPTE_Time) SetMFrames(v int8) {
	s.storage[3] = uint8(v)
}

// MSubFrameSampleOffset returns the MSubFrameSampleOffset field from the record's packed storage.
func (s *CAF_SMPTE_Time) MSubFrameSampleOffset() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetMSubFrameSampleOffset updates the MSubFrameSampleOffset field in the record's packed storage.
func (s *CAF_SMPTE_Time) SetMSubFrameSampleOffset(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// CAF_UUID_ChunkHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAF_UUID_ChunkHeader
type CAF_UUID_ChunkHeader struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [28]byte
}

// MHeader returns the MHeader field from the record's packed storage.
func (s *CAF_UUID_ChunkHeader) MHeader() CAFChunkHeader {
	return *(*CAFChunkHeader)(unsafe.Pointer(&s.storage[0]))
}

// SetMHeader updates the MHeader field in the record's packed storage.
func (s *CAF_UUID_ChunkHeader) SetMHeader(v CAFChunkHeader) {
	*(*CAFChunkHeader)(unsafe.Pointer(&s.storage[0])) = v
}

// MUUID returns the MUUID field from the record's packed storage.
func (s *CAF_UUID_ChunkHeader) MUUID() [16]uint8 {
	return *(*[16]uint8)(unsafe.Pointer(&s.storage[12]))
}

// SetMUUID updates the MUUID field in the record's packed storage.
func (s *CAF_UUID_ChunkHeader) SetMUUID(v [16]uint8) {
	*(*[16]uint8)(unsafe.Pointer(&s.storage[12])) = v
}

// CAMeterTrackEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAMeterTrackEntry
type CAMeterTrackEntry struct {
	Beats      CAClockBeats
	MeterNumer uint16
	MeterDenom uint16
}

// CATempoMapEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CATempoMapEntry
type CATempoMapEntry struct {
	Beats    CAClockBeats
	TempoBPM CAClockTempo
}

// ExtendedAudioFormatInfo - A specifier for the [kAudioFormatProperty_FormatList](<https://developer.apple.com/documentation/AudioToolbox/kAudioFormatProperty_FormatList>) property, including the codec to use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtendedAudioFormatInfo
type ExtendedAudioFormatInfo struct {
	MASBD             coreaudiotypes.AudioStreamBasicDescription // A format specification for an audio stream.
	MMagicCookie      unsafe.Pointer                             // Decompression information for the audio data format specified in the `mASBD` field.
	MMagicCookieSize  uint32                                     // The size, in bytes, of the `mMagicCookie` field.
	MClassDescription coreaudiotypes.AudioClassDescription       // A structure that describes an audio codec.

}

// ExtendedControlEvent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtendedControlEvent
type ExtendedControlEvent struct {
	GroupID   MusicDeviceGroupID
	ControlID AudioUnitParameterID
	Value     AudioUnitParameterValue
}

// ExtendedNoteOnEvent - Describes a note-on event with extended parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtendedNoteOnEvent
type ExtendedNoteOnEvent struct {
	InstrumentID   MusicDeviceInstrumentID
	GroupID        MusicDeviceGroupID
	Duration       float32
	ExtendedParams MusicDeviceNoteParams
}

// ExtendedTempoEvent - Describes a music track tempo in beats-per-minute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ExtendedTempoEvent
type ExtendedTempoEvent struct {
	Bpm float64 // The number of beats-per-minute.

}

// HostCallbackInfo - The time- and transport-related callback functions for an audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/HostCallbackInfo
type HostCallbackInfo struct {
	HostUserData            unsafe.Pointer                      // Custom data specified by your application. May be [NULL].
	BeatAndTempoProc        HostCallback_GetBeatAndTempo        // Your callback function that provides beat and tempo information to an audio unit. May be [NULL].
	MusicalTimeLocationProc HostCallback_GetMusicalTimeLocation // Your callback function that provides musical timeline information to an audio unit. May be [NULL].
	TransportStateProc      HostCallback_GetTransportState      // Your callback function that provides audio transport state information (play, rewind, and so on) to an audio unit. May be [NULL].
	TransportStateProc2     HostCallback_GetTransportState2
}

// MIDIChannelMessage - Describes a MIDI channel message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MIDIChannelMessage
type MIDIChannelMessage struct {
	Status   uint8 // Data specific to the channel message.
	Data1    uint8
	Data2    uint8
	Reserved uint8
}

// MIDIMetaEvent - Describes a MIDI metaevent such as lyric text, time signature, and so on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MIDIMetaEvent
type MIDIMetaEvent struct {
	MetaEventType uint8 // An integer that designates one of the types of MIDI metaevents.
	Unused1       uint8
	Unused2       uint8
	Unused3       uint8
	DataLength    uint32
	Data          [1]uint8
}

// MIDINoteMessage - Describes a MIDI note.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MIDINoteMessage
type MIDINoteMessage struct {
	Channel         uint8   // The MIDI channel to play the note on.
	Note            uint8   // The note to play.
	Velocity        uint8   // The key-press velocity for the note.
	ReleaseVelocity uint8   // The key-release velocity for the note. Use 0 if you don’t want to specify a particular value.
	Duration        float32 // The duration for the note.

}

// MIDIRawData - Describes a MIDI system-exclusive (SysEx) message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MIDIRawData
type MIDIRawData struct {
	Length uint32
	Data   [1]uint8
}

// MixerDistanceParams
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MixerDistanceParams
type MixerDistanceParams struct {
	MReferenceDistance float32
	MMaxDistance       float32
	MMaxAttenuation    float32
}

// MusicDeviceNoteParams
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceNoteParams
type MusicDeviceNoteParams struct {
	ArgCount  uint32
	MPitch    float32
	MVelocity float32
	MControls [1]NoteParamsControlValue
}

// MusicDeviceStdNoteParams
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceStdNoteParams
type MusicDeviceStdNoteParams struct {
	ArgCount  uint32
	MPitch    float32
	MVelocity float32
}

// MusicEventUserData - Describes a user-defined event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicEventUserData
type MusicEventUserData struct {
	Length uint32   // The size, in bytes, of the user data.
	Data   [1]uint8 // User-defined data.

}

// MusicTrackLoopInfo - Supports control of the looping behavior of a music track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/MusicTrackLoopInfo
type MusicTrackLoopInfo struct {
	LoopDuration  MusicTimeStamp // The point in a music track, measured in beats from the end of the music track, at which to begin playback during looped playback.
	NumberOfLoops int32          // The number of times to play the designated portion of the music track.

}

// NoteParamsControlValue
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/NoteParamsControlValue
type NoteParamsControlValue struct {
	MID    AudioUnitParameterID
	MValue AudioUnitParameterValue
}

// ParameterEvent - Describes an audio unit parameter automation event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ParameterEvent
type ParameterEvent struct {
	ParameterID AudioUnitParameterID
	Scope       AudioUnitScope
	Element     AudioUnitElement
	Value       AudioUnitParameterValue
}

// ScheduledAudioFileRegion
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ScheduledAudioFileRegion
type ScheduledAudioFileRegion struct {
	MTimeStamp              coreaudiotypes.AudioTimeStamp
	MCompletionProc         ScheduledAudioFileRegionCompletionProc // may be [NULL]
	MCompletionProcUserData unsafe.Pointer
	MAudioFile              *string // Must be a valid and already-open audio file object (of type [AudioFileID]), as declared in `AudioToolbox/AudioFile.h`.
	MLoopCount              uint32  // `0` = do not loop
	MStartFrame             int64   // The frame offset into the file.
	MFramesToPlay           uint32  // The number of frames to play.

}

// ScheduledAudioSlice
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/ScheduledAudioSlice
type ScheduledAudioSlice struct {
	MTimeStamp              coreaudiotypes.AudioTimeStamp
	MCompletionProc         ScheduledAudioSliceCompletionProc
	MCompletionProcUserData unsafe.Pointer
	MFlags                  AUScheduledAudioSliceFlags
	MReserved               uint32
	MReserved2              unsafe.Pointer
	MNumberFrames           uint32
	MBufferList             *coreaudiotypes.AudioBufferList
}
