// Code generated from Apple documentation. DO NOT EDIT.

package coremidi

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// MIDIChannelNumber values.
const (

	// MIDIChannelsWholePort is a constant value that indicates to use all channels of the port.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/MIDIChannelsWholePort
	MIDIChannelsWholePort MIDIChannelNumber = 0x7
)

// uint8 values.
const (

	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDI1UPMaxSysexSize
	KMIDI1UPMaxSysexSize uint8 = 6
)

// MIDICIDeviceID values.
const (

	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIDeviceIDFunctionBlock
	KMIDIDeviceIDFunctionBlock MIDICIDeviceID = 0x7

	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIDeviceIDUMPGroup
	KMIDIDeviceIDUMPGroup MIDICIDeviceID = 0x7e
)

// MIDIUInteger14 values.
const (

	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIUInteger14Max
	KMIDIUInteger14Max MIDIUInteger14 = 0x3FFF
)

// MIDIUInteger28 values.
const (

	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIUInteger28Max
	KMIDIUInteger28Max MIDIUInteger28 = 0xFFFFFFF
)

// MIDIUInteger2 values.
const (

	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIUInteger2Max
	KMIDIUInteger2Max MIDIUInteger2 = 0x3
)

// MIDIUInteger4 values.
const (

	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIUInteger4Max
	KMIDIUInteger4Max MIDIUInteger4 = 0xF
)

// MIDIUInteger7 values.
const (

	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIUInteger7Max
	KMIDIUInteger7Max MIDIUInteger7 = 0x7F
)

var (
	// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceManager/DictionaryKey/deviceObject
	MIDICIDeviceObjectKey MIDICIDeviceManagerDictionaryKey
	// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceManager/DictionaryKey/profileObject
	MIDICIProfileObjectKey MIDICIDeviceManagerDictionaryKey
)

var (
	// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceManager/deviceWasAddedNotification
	MIDICIDeviceWasAddedNotification foundation.NSNotificationName
	// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceManager/deviceWasRemovedNotification
	MIDICIDeviceWasRemovedNotification foundation.NSNotificationName
	// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceManager/profileWasRemovedNotification
	MIDICIProfileWasRemovedNotification foundation.NSNotificationName
	// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceManager/profileWasUpdatedNotification
	MIDICIProfileWasUpdatedNotification foundation.NSNotificationName
	// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpointManager/endpointWasAddedNotification
	MIDIUMPEndpointWasAddedNotification foundation.NSNotificationName
	// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpointManager/endpointWasRemovedNotification
	MIDIUMPEndpointWasRemovedNotification foundation.NSNotificationName
	// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpointManager/endpointWasUpdatedNotification
	MIDIUMPEndpointWasUpdatedNotification foundation.NSNotificationName
	// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpointManager/functionBlockWasUpdatedNotification
	MIDIUMPFunctionBlockWasUpdatedNotification foundation.NSNotificationName
)

var (
	// MIDINetworkBonjourServiceType is the Bonjour service type.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkBonjourServiceType
	MIDINetworkBonjourServiceType string
	// MIDINetworkNotificationContactsDidChange is indicates that the list of contacts changed.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkNotificationContactsDidChange
	MIDINetworkNotificationContactsDidChange string
	// MIDINetworkNotificationSessionDidChange is indicates that other aspects of the session changed, such as the connection list, connection policy, and so on.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkNotificationSessionDidChange
	MIDINetworkNotificationSessionDidChange string
)

var (
	// KMIDIDriverPropertyUsesSerial is a value that indicates whether the driver uses serial ports and is eligible to have serial ports assigned to it.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIDriverPropertyUsesSerial
	KMIDIDriverPropertyUsesSerial string
	// KMIDIPropertyAdvanceScheduleTimeMuSec is the recommended number of microseconds in advance that clients should schedule output.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyAdvanceScheduleTimeMuSec
	KMIDIPropertyAdvanceScheduleTimeMuSec string
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyAssociatedEndpoint
	KMIDIPropertyAssociatedEndpoint string
	// KMIDIPropertyCanRoute is a Boolean value that indicates whether the device or entity can route messages to or from external MIDI devices.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyCanRoute
	KMIDIPropertyCanRoute string
	// KMIDIPropertyConnectionUniqueID is the unique identifier of an external device attached to this connection.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyConnectionUniqueID
	KMIDIPropertyConnectionUniqueID string
	// KMIDIPropertyDeviceID is the user-visible System Exclusive (SysEx) identifier of a device or entity.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyDeviceID
	KMIDIPropertyDeviceID string
	// KMIDIPropertyDisplayName is the user-visible name for an endpoint that combines the device and endpoint names.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyDisplayName
	KMIDIPropertyDisplayName string
	// KMIDIPropertyDriverDeviceEditorApp is the full path to an app on the system that configures driver-owned devices.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyDriverDeviceEditorApp
	KMIDIPropertyDriverDeviceEditorApp string
	// KMIDIPropertyDriverOwner is the name of the driver that owns a device, entity, or endpoint.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyDriverOwner
	KMIDIPropertyDriverOwner string
	// KMIDIPropertyDriverVersion is the version of the driver that owns a device, entity, or endpoint.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyDriverVersion
	KMIDIPropertyDriverVersion string
	// KMIDIPropertyImage is the full path to a device icon on the system.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyImage
	KMIDIPropertyImage string
	// KMIDIPropertyIsBroadcast is a Boolean value that indicates whether the endpoint broadcasts messages to all of the other endpoints in the device.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyIsBroadcast
	KMIDIPropertyIsBroadcast string
	// KMIDIPropertyIsDrumMachine is a Boolean value that indicates whether the device or entity’s samples aren’t transposable, as with a drum kit.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyIsDrumMachine
	KMIDIPropertyIsDrumMachine string
	// KMIDIPropertyIsEffectUnit is a Boolean value that indicates whether the device or entity primarily acts as a MIDI-controlled audio effect.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyIsEffectUnit
	KMIDIPropertyIsEffectUnit string
	// KMIDIPropertyIsEmbeddedEntity is a Boolean value that indicates whether this entity or endpoint has external MIDI connections.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyIsEmbeddedEntity
	KMIDIPropertyIsEmbeddedEntity string
	// KMIDIPropertyIsMixer is a Boolean value that indicates whether the device or entity mixes external audio signals.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyIsMixer
	KMIDIPropertyIsMixer string
	// KMIDIPropertyIsSampler is a Boolean value that indicates whether the device or entity plays audio samples in response to MIDI note messages.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyIsSampler
	KMIDIPropertyIsSampler string
	// KMIDIPropertyManufacturer is the manufacturer name of a device or endpoint.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyManufacturer
	KMIDIPropertyManufacturer string
	// KMIDIPropertyMaxReceiveChannels is the maximum number of MIDI channels on which a device may simultaneously receive channel messages.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyMaxReceiveChannels
	KMIDIPropertyMaxReceiveChannels string
	// KMIDIPropertyMaxSysExSpeed is the maximum rate, in bytes per second, at which the system may reliably send System Exclusive (SysEx) messages to this object.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyMaxSysExSpeed
	KMIDIPropertyMaxSysExSpeed string
	// KMIDIPropertyMaxTransmitChannels is the maximum number of MIDI channels on which a device may simultaneously transmit channel messages.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyMaxTransmitChannels
	KMIDIPropertyMaxTransmitChannels string
	// KMIDIPropertyModel is the model name of a device or endpoint.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyModel
	KMIDIPropertyModel string
	// KMIDIPropertyName is a name for a device, entity, or endpoint.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyName
	KMIDIPropertyName string
	// KMIDIPropertyNameConfigurationDictionary is the device’s current patch, note, and control name values in MIDINameDocument XML format.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyNameConfigurationDictionary
	KMIDIPropertyNameConfigurationDictionary string
	// KMIDIPropertyOffline is a Boolean value that indicates whether the object is offline.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyOffline
	KMIDIPropertyOffline string
	// KMIDIPropertyPanDisruptsStereo is a Boolean value that indicates whether the MIDI pan messages sent to the device or entity cause undesirable effects when playing stereo sounds.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyPanDisruptsStereo
	KMIDIPropertyPanDisruptsStereo string
	// KMIDIPropertyPrivate is a Boolean value that indicates whether the system hides an endpoint from other clients.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyPrivate
	KMIDIPropertyPrivate string
	// KMIDIPropertyProtocolID is the native protocol in which the endpoint communicates.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyProtocolID
	KMIDIPropertyProtocolID string
	// KMIDIPropertyReceiveChannels is the bitmap of channels on which the object receives messages.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyReceiveChannels
	KMIDIPropertyReceiveChannels string
	// KMIDIPropertyReceivesBankSelectLSB is a Boolean value that indicates whether the device or entity responds to MIDI bank select LSB messages.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyReceivesBankSelectLSB
	KMIDIPropertyReceivesBankSelectLSB string
	// KMIDIPropertyReceivesBankSelectMSB is a Boolean value that indicates whether the device or entity responds to MIDI bank select MSB messages.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyReceivesBankSelectMSB
	KMIDIPropertyReceivesBankSelectMSB string
	// KMIDIPropertyReceivesClock is a Boolean value that indicates whether the device or entity responds to MIDI beat clock messages.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyReceivesClock
	KMIDIPropertyReceivesClock string
	// KMIDIPropertyReceivesMTC is a Boolean value that indicates whether the device or entity responds to MIDI Time Code messages.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyReceivesMTC
	KMIDIPropertyReceivesMTC string
	// KMIDIPropertyReceivesNotes is a Boolean value that indicates whether the device or entity responds to MIDI Note On messages.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyReceivesNotes
	KMIDIPropertyReceivesNotes string
	// KMIDIPropertyReceivesProgramChanges is a Boolean value that indicates whether the device or entity responds to MIDI Program Change messages.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyReceivesProgramChanges
	KMIDIPropertyReceivesProgramChanges string
	// KMIDIPropertySingleRealtimeEntity is the 0-based index of the entity on which incoming real-time messages from the device appear to have originated.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertySingleRealtimeEntity
	KMIDIPropertySingleRealtimeEntity string
	// KMIDIPropertySupportsGeneralMIDI is a Boolean value that indicates whether the device or entity implements the General MIDI specification.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertySupportsGeneralMIDI
	KMIDIPropertySupportsGeneralMIDI string
	// KMIDIPropertySupportsMMC is a Boolean value that indicates whether the device or entity implements the MIDI Machine Control portion of the MIDI specification.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertySupportsMMC
	KMIDIPropertySupportsMMC string
	// KMIDIPropertySupportsShowControl is a Boolean value that indicates whether the device implements the MIDI Show Control specification.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertySupportsShowControl
	KMIDIPropertySupportsShowControl string
	// KMIDIPropertyTransmitChannels is the bitmap of channels on which the object transmits messages.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyTransmitChannels
	KMIDIPropertyTransmitChannels string
	// KMIDIPropertyTransmitsBankSelectLSB is a Boolean value that indicates whether the device or entity transmits MIDI bank select LSB messages.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyTransmitsBankSelectLSB
	KMIDIPropertyTransmitsBankSelectLSB string
	// KMIDIPropertyTransmitsBankSelectMSB is a Boolean value that indicates whether the device or entity transmits MIDI bank select MSB messages.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyTransmitsBankSelectMSB
	KMIDIPropertyTransmitsBankSelectMSB string
	// KMIDIPropertyTransmitsClock is a Boolean value that indicates whether the device or entity transmits MIDI beat clock messages.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyTransmitsClock
	KMIDIPropertyTransmitsClock string
	// KMIDIPropertyTransmitsMTC is a Boolean value that indicates whether the device or entity transmits MIDI Time Code messages.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyTransmitsMTC
	KMIDIPropertyTransmitsMTC string
	// KMIDIPropertyTransmitsNotes is a Boolean value that indicates whether the device or entity transmits MIDI note messages.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyTransmitsNotes
	KMIDIPropertyTransmitsNotes string
	// KMIDIPropertyTransmitsProgramChanges is a Boolean value that indicates whether the device or entity transmits MIDI Program Change messages.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyTransmitsProgramChanges
	KMIDIPropertyTransmitsProgramChanges string
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyUMPActiveGroupBitmap
	KMIDIPropertyUMPActiveGroupBitmap string
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyUMPCanTransmitGroupless
	KMIDIPropertyUMPCanTransmitGroupless string
	// KMIDIPropertyUMPEnabled is kMIDIPropertyUMPEnabled.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyUMPEnabled
	KMIDIPropertyUMPEnabled string
	// KMIDIPropertyUniqueID is the unique identifier of a device, entity, or, endpoint.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/kMIDIPropertyUniqueID
	KMIDIPropertyUniqueID string
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MIDICIDeviceObjectKey"); err == nil && ptr != 0 {
		MIDICIDeviceObjectKey = objc.ValueAt[MIDICIDeviceManagerDictionaryKey](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MIDICIDeviceWasAddedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MIDICIDeviceWasAddedNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MIDICIDeviceWasRemovedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MIDICIDeviceWasRemovedNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MIDICIProfileObjectKey"); err == nil && ptr != 0 {
		MIDICIProfileObjectKey = objc.ValueAt[MIDICIDeviceManagerDictionaryKey](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MIDICIProfileWasRemovedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MIDICIProfileWasRemovedNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MIDICIProfileWasUpdatedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MIDICIProfileWasUpdatedNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MIDINetworkBonjourServiceType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MIDINetworkBonjourServiceType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MIDINetworkNotificationContactsDidChange"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MIDINetworkNotificationContactsDidChange = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MIDINetworkNotificationSessionDidChange"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MIDINetworkNotificationSessionDidChange = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MIDIUMPEndpointWasAddedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MIDIUMPEndpointWasAddedNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MIDIUMPEndpointWasRemovedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MIDIUMPEndpointWasRemovedNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MIDIUMPEndpointWasUpdatedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MIDIUMPEndpointWasUpdatedNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MIDIUMPFunctionBlockWasUpdatedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MIDIUMPFunctionBlockWasUpdatedNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIDriverPropertyUsesSerial"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIDriverPropertyUsesSerial = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyAdvanceScheduleTimeMuSec"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyAdvanceScheduleTimeMuSec = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyAssociatedEndpoint"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyAssociatedEndpoint = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyCanRoute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyCanRoute = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyConnectionUniqueID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyConnectionUniqueID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyDeviceID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyDeviceID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyDisplayName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyDisplayName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyDriverDeviceEditorApp"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyDriverDeviceEditorApp = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyDriverOwner"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyDriverOwner = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyDriverVersion"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyDriverVersion = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyImage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyImage = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyIsBroadcast"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyIsBroadcast = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyIsDrumMachine"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyIsDrumMachine = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyIsEffectUnit"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyIsEffectUnit = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyIsEmbeddedEntity"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyIsEmbeddedEntity = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyIsMixer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyIsMixer = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyIsSampler"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyIsSampler = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyManufacturer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyManufacturer = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyMaxReceiveChannels"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyMaxReceiveChannels = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyMaxSysExSpeed"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyMaxSysExSpeed = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyMaxTransmitChannels"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyMaxTransmitChannels = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyModel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyModel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyNameConfigurationDictionary"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyNameConfigurationDictionary = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyOffline"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyOffline = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyPanDisruptsStereo"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyPanDisruptsStereo = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyPrivate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyPrivate = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyProtocolID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyProtocolID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyReceiveChannels"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyReceiveChannels = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyReceivesBankSelectLSB"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyReceivesBankSelectLSB = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyReceivesBankSelectMSB"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyReceivesBankSelectMSB = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyReceivesClock"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyReceivesClock = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyReceivesMTC"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyReceivesMTC = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyReceivesNotes"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyReceivesNotes = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyReceivesProgramChanges"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyReceivesProgramChanges = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertySingleRealtimeEntity"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertySingleRealtimeEntity = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertySupportsGeneralMIDI"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertySupportsGeneralMIDI = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertySupportsMMC"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertySupportsMMC = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertySupportsShowControl"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertySupportsShowControl = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyTransmitChannels"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyTransmitChannels = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyTransmitsBankSelectLSB"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyTransmitsBankSelectLSB = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyTransmitsBankSelectMSB"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyTransmitsBankSelectMSB = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyTransmitsClock"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyTransmitsClock = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyTransmitsMTC"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyTransmitsMTC = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyTransmitsNotes"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyTransmitsNotes = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyTransmitsProgramChanges"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyTransmitsProgramChanges = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyUMPActiveGroupBitmap"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyUMPActiveGroupBitmap = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyUMPCanTransmitGroupless"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyUMPCanTransmitGroupless = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyUMPEnabled"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyUMPEnabled = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kMIDIPropertyUniqueID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KMIDIPropertyUniqueID = objc.GoString(cstr)
			}
		}
	}

}
