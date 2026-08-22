// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"fmt"
)

type KMIDI int32

const (
	// KMIDIIDNotUnique: The identifier you’re trying to set isn’t unique.
	KMIDIIDNotUnique KMIDI = -10843
	// KMIDIInvalidClient: The client is invalid.
	KMIDIInvalidClient KMIDI = -10830
	// KMIDIInvalidPort: The port is invalid.
	KMIDIInvalidPort KMIDI = -10831
	// KMIDIMessageSendErr: The communication with the MIDI server failed.
	KMIDIMessageSendErr KMIDI = -10838
	// KMIDINoConnection: The connection you’re trying to close doesn’t exist.
	KMIDINoConnection KMIDI = -10833
	// KMIDINoCurrentSetup: A MIDI setup object doesn’t currently exist.
	KMIDINoCurrentSetup KMIDI = -10837
	// KMIDINotPermitted: The process doesn’t have privileges for the requested operation.
	KMIDINotPermitted KMIDI = -10844
	// KMIDIObjectNotFound: The requested object doesn’t exist.
	KMIDIObjectNotFound KMIDI = -10842
	// KMIDIServerStartErr: The system can’t start the MIDI server.
	KMIDIServerStartErr KMIDI = -10839
	// KMIDISetupFormatErr: The system can’t read the saved state.
	KMIDISetupFormatErr KMIDI = -10840
	// KMIDIUnknownEndpoint: The system doesn’t recognize the endpoint.
	KMIDIUnknownEndpoint KMIDI = -10834
	// KMIDIUnknownError: The system can’t perform the requested operation.
	KMIDIUnknownError KMIDI = -10845
	// KMIDIUnknownProperty: The property you’re trying to query isn’t set on the object.
	KMIDIUnknownProperty KMIDI = -10835
	// KMIDIWrongEndpointType: A function received a source endpoint when it required a destination endpoint, or vice versa.
	KMIDIWrongEndpointType KMIDI = -10832
	// KMIDIWrongPropertyType: The value you assigned to the property is the wrong type.
	KMIDIWrongPropertyType KMIDI = -10836
	// KMIDIWrongThread: A driver is calling a non-I/O function in the server from a thread other than the server’s main thread.
	KMIDIWrongThread KMIDI = -10841
)

func (e KMIDI) String() string {
	switch e {
	case KMIDIIDNotUnique:
		return "KMIDIIDNotUnique"
	case KMIDIInvalidClient:
		return "KMIDIInvalidClient"
	case KMIDIInvalidPort:
		return "KMIDIInvalidPort"
	case KMIDIMessageSendErr:
		return "KMIDIMessageSendErr"
	case KMIDINoConnection:
		return "KMIDINoConnection"
	case KMIDINoCurrentSetup:
		return "KMIDINoCurrentSetup"
	case KMIDINotPermitted:
		return "KMIDINotPermitted"
	case KMIDIObjectNotFound:
		return "KMIDIObjectNotFound"
	case KMIDIServerStartErr:
		return "KMIDIServerStartErr"
	case KMIDISetupFormatErr:
		return "KMIDISetupFormatErr"
	case KMIDIUnknownEndpoint:
		return "KMIDIUnknownEndpoint"
	case KMIDIUnknownError:
		return "KMIDIUnknownError"
	case KMIDIUnknownProperty:
		return "KMIDIUnknownProperty"
	case KMIDIWrongEndpointType:
		return "KMIDIWrongEndpointType"
	case KMIDIWrongPropertyType:
		return "KMIDIWrongPropertyType"
	case KMIDIWrongThread:
		return "KMIDIWrongThread"
	default:
		return fmt.Sprintf("KMIDI(%d)", e)
	}
}

const KMIDIInvalidUniqueID int32 = 0

type KmidithruconnectionMaxendpoints uint32

const (
	// KMIDIThruConnection_MaxEndpoints: The maximum number of endpoints for this connection.
	KMIDIThruConnection_MaxEndpoints KmidithruconnectionMaxendpoints = 8
)

func (e KmidithruconnectionMaxendpoints) String() string {
	switch e {
	case KMIDIThruConnection_MaxEndpoints:
		return "KMIDIThruConnection_MaxEndpoints"
	default:
		return fmt.Sprintf("KmidithruconnectionMaxendpoints(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDICICategoryOptions
type MIDICICategoryOptions uint8

const (
	KMIDICICategoryOptionsProcessInquirySupported       MIDICICategoryOptions = 16
	KMIDICICategoryOptionsProfileConfigurationSupported MIDICICategoryOptions = 4
	KMIDICICategoryOptionsPropertyExchangeSupported     MIDICICategoryOptions = 8
	KMIDICICategoryOptionsProtocolNegotiation           MIDICICategoryOptions = 2
)

func (e MIDICICategoryOptions) String() string {
	switch e {
	case KMIDICICategoryOptionsProcessInquirySupported:
		return "KMIDICICategoryOptionsProcessInquirySupported"
	case KMIDICICategoryOptionsProfileConfigurationSupported:
		return "KMIDICICategoryOptionsProfileConfigurationSupported"
	case KMIDICICategoryOptionsPropertyExchangeSupported:
		return "KMIDICICategoryOptionsPropertyExchangeSupported"
	case KMIDICICategoryOptionsProtocolNegotiation:
		return "KMIDICICategoryOptionsProtocolNegotiation"
	default:
		return fmt.Sprintf("MIDICICategoryOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceType
type MIDICIDeviceType uint8

const ()

// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIManagementMessageType
type MIDICIManagementMessageType uint8

const ()

// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIProcessInquiryMessageType
type MIDICIProcessInquiryMessageType uint8

const ()

// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileMessageType
type MIDICIProfileMessageType uint8

const ()

// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileType
type MIDICIProfileType uint8

const ()

// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIPropertyExchangeMessageType
type MIDICIPropertyExchangeMessageType uint8

const ()

// See: https://developer.apple.com/documentation/CoreMIDI/MIDICVStatus
type MIDICVStatus uint32

const ()

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIMessageType
type MIDIMessageType uint32

const ()

// See: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkConnectionPolicy
type MIDINetworkConnectionPolicy uint

const (
	MIDINetworkConnectionPolicy_Anyone             MIDINetworkConnectionPolicy = 2
	MIDINetworkConnectionPolicy_HostsInContactList MIDINetworkConnectionPolicy = 1
	MIDINetworkConnectionPolicy_NoOne              MIDINetworkConnectionPolicy = 0
)

func (e MIDINetworkConnectionPolicy) String() string {
	switch e {
	case MIDINetworkConnectionPolicy_Anyone:
		return "MIDINetworkConnectionPolicy_Anyone"
	case MIDINetworkConnectionPolicy_HostsInContactList:
		return "MIDINetworkConnectionPolicy_HostsInContactList"
	case MIDINetworkConnectionPolicy_NoOne:
		return "MIDINetworkConnectionPolicy_NoOne"
	default:
		return fmt.Sprintf("MIDINetworkConnectionPolicy(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDINoteAttribute
type MIDINoteAttribute uint8

const ()

// See: https://developer.apple.com/documentation/CoreMIDI/MIDINotificationMessageID
type MIDINotificationMessageID int32

const (
	// KMIDIMsgIOError: A driver I/O error occurred.
	KMIDIMsgIOError MIDINotificationMessageID = 7
	// KMIDIMsgObjectAdded: The system added a device, entity, or endpoint.
	KMIDIMsgObjectAdded MIDINotificationMessageID = 2
	// KMIDIMsgObjectRemoved: The system removed a device, entity, or endpoint.
	KMIDIMsgObjectRemoved MIDINotificationMessageID = 3
	// KMIDIMsgPropertyChanged: An object’s property value changed.
	KMIDIMsgPropertyChanged MIDINotificationMessageID = 4
	// KMIDIMsgSerialPortOwnerChanged: The system changed a serial port owner.
	KMIDIMsgSerialPortOwnerChanged MIDINotificationMessageID = 6
	// KMIDIMsgSetupChanged: Some aspect of the current MIDI setup changed.
	KMIDIMsgSetupChanged MIDINotificationMessageID = 1
	// KMIDIMsgThruConnectionsChanged: The system created or disposed of a persistent MIDI Thru connection.
	KMIDIMsgThruConnectionsChanged MIDINotificationMessageID = 5
)

func (e MIDINotificationMessageID) String() string {
	switch e {
	case KMIDIMsgIOError:
		return "KMIDIMsgIOError"
	case KMIDIMsgObjectAdded:
		return "KMIDIMsgObjectAdded"
	case KMIDIMsgObjectRemoved:
		return "KMIDIMsgObjectRemoved"
	case KMIDIMsgPropertyChanged:
		return "KMIDIMsgPropertyChanged"
	case KMIDIMsgSerialPortOwnerChanged:
		return "KMIDIMsgSerialPortOwnerChanged"
	case KMIDIMsgSetupChanged:
		return "KMIDIMsgSetupChanged"
	case KMIDIMsgThruConnectionsChanged:
		return "KMIDIMsgThruConnectionsChanged"
	default:
		return fmt.Sprintf("MIDINotificationMessageID(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectType
type MIDIObjectType int32

const (
	// KMIDIObjectType_Destination: A MIDI destination.
	KMIDIObjectType_Destination MIDIObjectType = 3
	// KMIDIObjectType_Device: A MIDI device.
	KMIDIObjectType_Device MIDIObjectType = 0
	// KMIDIObjectType_Entity: A MIDI entity.
	KMIDIObjectType_Entity MIDIObjectType = 1
	// KMIDIObjectType_ExternalDestination: An external destination.
	KMIDIObjectType_ExternalDestination MIDIObjectType = 19
	// KMIDIObjectType_ExternalDevice: An external device.
	KMIDIObjectType_ExternalDevice MIDIObjectType = 0x10
	// KMIDIObjectType_ExternalEntity: An external entity.
	KMIDIObjectType_ExternalEntity MIDIObjectType = 17
	// KMIDIObjectType_ExternalSource: An external source.
	KMIDIObjectType_ExternalSource MIDIObjectType = 18
	// KMIDIObjectType_Other: A MIDI object with an undefined type.
	KMIDIObjectType_Other MIDIObjectType = -1
	// KMIDIObjectType_Source: A MIDI source.
	KMIDIObjectType_Source MIDIObjectType = 2
)

func (e MIDIObjectType) String() string {
	switch e {
	case KMIDIObjectType_Destination:
		return "KMIDIObjectType_Destination"
	case KMIDIObjectType_Device:
		return "KMIDIObjectType_Device"
	case KMIDIObjectType_Entity:
		return "KMIDIObjectType_Entity"
	case KMIDIObjectType_ExternalDestination:
		return "KMIDIObjectType_ExternalDestination"
	case KMIDIObjectType_ExternalDevice:
		return "KMIDIObjectType_ExternalDevice"
	case KMIDIObjectType_ExternalEntity:
		return "KMIDIObjectType_ExternalEntity"
	case KMIDIObjectType_ExternalSource:
		return "KMIDIObjectType_ExternalSource"
	case KMIDIObjectType_Other:
		return "KMIDIObjectType_Other"
	case KMIDIObjectType_Source:
		return "KMIDIObjectType_Source"
	default:
		return fmt.Sprintf("MIDIObjectType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIPerNoteManagementOptions
type MIDIPerNoteManagementOptions uint8

const (
	KMIDIPerNoteManagementDetach MIDIPerNoteManagementOptions = 0x2
	KMIDIPerNoteManagementReset  MIDIPerNoteManagementOptions = 0x1
)

func (e MIDIPerNoteManagementOptions) String() string {
	switch e {
	case KMIDIPerNoteManagementDetach:
		return "KMIDIPerNoteManagementDetach"
	case KMIDIPerNoteManagementReset:
		return "KMIDIPerNoteManagementReset"
	default:
		return fmt.Sprintf("MIDIPerNoteManagementOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIProgramChangeOptions
type MIDIProgramChangeOptions uint8

const (
	KMIDIProgramChangeBankValid MIDIProgramChangeOptions = 0x1
)

func (e MIDIProgramChangeOptions) String() string {
	switch e {
	case KMIDIProgramChangeBankValid:
		return "KMIDIProgramChangeBankValid"
	default:
		return fmt.Sprintf("MIDIProgramChangeOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIProtocolID
type MIDIProtocolID int32

const ()

// See: https://developer.apple.com/documentation/CoreMIDI/MIDISysExStatus
type MIDISysExStatus uint32

const ()

// See: https://developer.apple.com/documentation/CoreMIDI/MIDISystemStatus
type MIDISystemStatus uint32

const ()

// See: https://developer.apple.com/documentation/CoreMIDI/MIDITransformControlType
type MIDITransformControlType uint8

const (
	// KMIDIControlType_14Bit: A 14-bit control type.
	KMIDIControlType_14Bit MIDITransformControlType = 1
	// KMIDIControlType_14BitNRPN: A 14-bit Nonregistered Parameter Number (RPN).
	KMIDIControlType_14BitNRPN MIDITransformControlType = 5
	// KMIDIControlType_14BitRPN: A 14-bit Registered Parameter Number (RPN).
	KMIDIControlType_14BitRPN MIDITransformControlType = 3
	// KMIDIControlType_7Bit: A 7-bit control type.
	KMIDIControlType_7Bit MIDITransformControlType = 0
	// KMIDIControlType_7BitNRPN: A 7-bit Nonregistered Parameter Number (RPN).
	KMIDIControlType_7BitNRPN MIDITransformControlType = 4
	// KMIDIControlType_7BitRPN: A 7-bit Registered Parameter Number (RPN).
	KMIDIControlType_7BitRPN MIDITransformControlType = 2
)

func (e MIDITransformControlType) String() string {
	switch e {
	case KMIDIControlType_14Bit:
		return "KMIDIControlType_14Bit"
	case KMIDIControlType_14BitNRPN:
		return "KMIDIControlType_14BitNRPN"
	case KMIDIControlType_14BitRPN:
		return "KMIDIControlType_14BitRPN"
	case KMIDIControlType_7Bit:
		return "KMIDIControlType_7Bit"
	case KMIDIControlType_7BitNRPN:
		return "KMIDIControlType_7BitNRPN"
	case KMIDIControlType_7BitRPN:
		return "KMIDIControlType_7BitRPN"
	default:
		return fmt.Sprintf("MIDITransformControlType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDITransformType
type MIDITransformType uint16

const (
	// KMIDITransform_Add: A transform that adds a parameter value.
	KMIDITransform_Add MIDITransformType = 8
	// KMIDITransform_FilterOut: A transformation that filters out an event type.
	KMIDITransform_FilterOut MIDITransformType = 1
	// KMIDITransform_MapControl: A transformation that changes a specified control number to a supplied parameter value.
	KMIDITransform_MapControl MIDITransformType = 2
	// KMIDITransform_MapValue: A transform that maps one value to another.
	KMIDITransform_MapValue MIDITransformType = 12
	// KMIDITransform_MaxValue: A transform that sets the maximum value to the specified parameter value.
	KMIDITransform_MaxValue MIDITransformType = 11
	// KMIDITransform_MinValue: A transform that sets the minimum value to the specified parameter value.
	KMIDITransform_MinValue MIDITransformType = 10
	// KMIDITransform_None: No transformation.
	KMIDITransform_None MIDITransformType = 0
	// KMIDITransform_Scale: A transform that multiplies by the specified parameter value.
	KMIDITransform_Scale MIDITransformType = 9
)

func (e MIDITransformType) String() string {
	switch e {
	case KMIDITransform_Add:
		return "KMIDITransform_Add"
	case KMIDITransform_FilterOut:
		return "KMIDITransform_FilterOut"
	case KMIDITransform_MapControl:
		return "KMIDITransform_MapControl"
	case KMIDITransform_MapValue:
		return "KMIDITransform_MapValue"
	case KMIDITransform_MaxValue:
		return "KMIDITransform_MaxValue"
	case KMIDITransform_MinValue:
		return "KMIDITransform_MinValue"
	case KMIDITransform_None:
		return "KMIDITransform_None"
	case KMIDITransform_Scale:
		return "KMIDITransform_Scale"
	default:
		return fmt.Sprintf("MIDITransformType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIObjectBackingType
type MIDIUMPCIObjectBackingType uint8

const ()

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlockDirection
type MIDIUMPFunctionBlockDirection int32

const ()

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlockMIDI1Info
type MIDIUMPFunctionBlockMIDI1Info int32

const ()

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlockUIHint
type MIDIUMPFunctionBlockUIHint int32

const ()

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPProtocolOptions
type MIDIUMPProtocolOptions uint8

const (
	KMIDIUMPProtocolOptionsMIDI1 MIDIUMPProtocolOptions = 1
	KMIDIUMPProtocolOptionsMIDI2 MIDIUMPProtocolOptions = 2
)

func (e MIDIUMPProtocolOptions) String() string {
	switch e {
	case KMIDIUMPProtocolOptionsMIDI1:
		return "KMIDIUMPProtocolOptionsMIDI1"
	case KMIDIUMPProtocolOptionsMIDI2:
		return "KMIDIUMPProtocolOptionsMIDI2"
	default:
		return fmt.Sprintf("MIDIUMPProtocolOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUtilityStatus
type MIDIUtilityStatus uint32

const ()

// See: https://developer.apple.com/documentation/CoreMIDI/UMPStreamMessageFormat
type UMPStreamMessageFormat uint8

const ()

// See: https://developer.apple.com/documentation/CoreMIDI/UMPStreamMessageStatus
type UMPStreamMessageStatus uint32

const ()
