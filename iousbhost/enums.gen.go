// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostAbortOption
type IOUSBHostAbortOption uint

const (
	// IOUSBHostAbortOptionAsynchronous: The option to abort input/output requests asynchronously.
	IOUSBHostAbortOptionAsynchronous IOUSBHostAbortOption = 0
	// IOUSBHostAbortOptionSynchronous: The option to abort input/output requests synchronously.
	IOUSBHostAbortOptionSynchronous IOUSBHostAbortOption = 1
)

func (e IOUSBHostAbortOption) String() string {
	switch e {
	case IOUSBHostAbortOptionAsynchronous:
		return "IOUSBHostAbortOptionAsynchronous"
	case IOUSBHostAbortOptionSynchronous:
		return "IOUSBHostAbortOptionSynchronous"
	default:
		return fmt.Sprintf("IOUSBHostAbortOption(%d)", e)
	}
}

type IOUSBHostCICapabilitiesMessage uint32

const (
	IOUSBHostCICapabilitiesMessageControlPortCount                  IOUSBHostCICapabilitiesMessage = 983040
	IOUSBHostCICapabilitiesMessageControlPortCountPhase             IOUSBHostCICapabilitiesMessage = 16
	IOUSBHostCICapabilitiesMessageData0CommandTimeoutThreshold      IOUSBHostCICapabilitiesMessage = 3
	IOUSBHostCICapabilitiesMessageData0CommandTimeoutThresholdPhase IOUSBHostCICapabilitiesMessage = 0
	IOUSBHostCICapabilitiesMessageData0ConnectionLatency            IOUSBHostCICapabilitiesMessage = 240
	IOUSBHostCICapabilitiesMessageData0ConnectionLatencyPhase       IOUSBHostCICapabilitiesMessage = 4
)

func (e IOUSBHostCICapabilitiesMessage) String() string {
	switch e {
	case IOUSBHostCICapabilitiesMessageControlPortCount:
		return "IOUSBHostCICapabilitiesMessageControlPortCount"
	case IOUSBHostCICapabilitiesMessageControlPortCountPhase:
		return "IOUSBHostCICapabilitiesMessageControlPortCountPhase"
	case IOUSBHostCICapabilitiesMessageData0CommandTimeoutThreshold:
		return "IOUSBHostCICapabilitiesMessageData0CommandTimeoutThreshold"
	case IOUSBHostCICapabilitiesMessageData0CommandTimeoutThresholdPhase:
		return "IOUSBHostCICapabilitiesMessageData0CommandTimeoutThresholdPhase"
	case IOUSBHostCICapabilitiesMessageData0ConnectionLatency:
		return "IOUSBHostCICapabilitiesMessageData0ConnectionLatency"
	case IOUSBHostCICapabilitiesMessageData0ConnectionLatencyPhase:
		return "IOUSBHostCICapabilitiesMessageData0ConnectionLatencyPhase"
	default:
		return fmt.Sprintf("IOUSBHostCICapabilitiesMessage(%d)", e)
	}
}

type IOUSBHostCICommandMessage uint32

const (
	IOUSBHostCICommandMessageControlStatus             IOUSBHostCICommandMessage = 3840
	IOUSBHostCICommandMessageControlStatusPhase        IOUSBHostCICommandMessage = 8
	IOUSBHostCICommandMessageData0DeviceAddress        IOUSBHostCICommandMessage = 255
	IOUSBHostCICommandMessageData0DeviceAddressPhase   IOUSBHostCICommandMessage = 0
	IOUSBHostCICommandMessageData0EndpointAddress      IOUSBHostCICommandMessage = 65280
	IOUSBHostCICommandMessageData0EndpointAddressPhase IOUSBHostCICommandMessage = 8
	IOUSBHostCICommandMessageData0RootPort             IOUSBHostCICommandMessage = 15
	IOUSBHostCICommandMessageData0RootPortPhase        IOUSBHostCICommandMessage = 0
	IOUSBHostCICommandMessageData0StreamID             IOUSBHostCICommandMessage = 4294901760
	IOUSBHostCICommandMessageData0StreamIDPhase        IOUSBHostCICommandMessage = 16
)

func (e IOUSBHostCICommandMessage) String() string {
	switch e {
	case IOUSBHostCICommandMessageControlStatus:
		return "IOUSBHostCICommandMessageControlStatus"
	case IOUSBHostCICommandMessageControlStatusPhase:
		return "IOUSBHostCICommandMessageControlStatusPhase"
	case IOUSBHostCICommandMessageData0DeviceAddress:
		return "IOUSBHostCICommandMessageData0DeviceAddress"
	case IOUSBHostCICommandMessageData0DeviceAddressPhase:
		return "IOUSBHostCICommandMessageData0DeviceAddressPhase"
	case IOUSBHostCICommandMessageData0EndpointAddress:
		return "IOUSBHostCICommandMessageData0EndpointAddress"
	case IOUSBHostCICommandMessageData0RootPort:
		return "IOUSBHostCICommandMessageData0RootPort"
	case IOUSBHostCICommandMessageData0StreamID:
		return "IOUSBHostCICommandMessageData0StreamID"
	case IOUSBHostCICommandMessageData0StreamIDPhase:
		return "IOUSBHostCICommandMessageData0StreamIDPhase"
	default:
		return fmt.Sprintf("IOUSBHostCICommandMessage(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIControllerState
type IOUSBHostCIControllerState uint32

const (
	IOUSBHostCIControllerStateActive IOUSBHostCIControllerState = 2
	IOUSBHostCIControllerStateOff    IOUSBHostCIControllerState = 0
	IOUSBHostCIControllerStatePaused IOUSBHostCIControllerState = 1
)

func (e IOUSBHostCIControllerState) String() string {
	switch e {
	case IOUSBHostCIControllerStateActive:
		return "IOUSBHostCIControllerStateActive"
	case IOUSBHostCIControllerStateOff:
		return "IOUSBHostCIControllerStateOff"
	case IOUSBHostCIControllerStatePaused:
		return "IOUSBHostCIControllerStatePaused"
	default:
		return fmt.Sprintf("IOUSBHostCIControllerState(%d)", e)
	}
}

type IOUSBHostCIDeviceCreateCommandData0 uint32

const (
	IOUSBHostCIDeviceCreateCommandData0RootPort      IOUSBHostCIDeviceCreateCommandData0 = 15
	IOUSBHostCIDeviceCreateCommandData0RootPortPhase IOUSBHostCIDeviceCreateCommandData0 = 0
	IOUSBHostCIDeviceCreateCommandData0Route         IOUSBHostCIDeviceCreateCommandData0 = 16777200
	IOUSBHostCIDeviceCreateCommandData0RoutePhase    IOUSBHostCIDeviceCreateCommandData0 = 4
)

func (e IOUSBHostCIDeviceCreateCommandData0) String() string {
	switch e {
	case IOUSBHostCIDeviceCreateCommandData0RootPort:
		return "IOUSBHostCIDeviceCreateCommandData0RootPort"
	case IOUSBHostCIDeviceCreateCommandData0RootPortPhase:
		return "IOUSBHostCIDeviceCreateCommandData0RootPortPhase"
	case IOUSBHostCIDeviceCreateCommandData0Route:
		return "IOUSBHostCIDeviceCreateCommandData0Route"
	case IOUSBHostCIDeviceCreateCommandData0RoutePhase:
		return "IOUSBHostCIDeviceCreateCommandData0RoutePhase"
	default:
		return fmt.Sprintf("IOUSBHostCIDeviceCreateCommandData0(%d)", e)
	}
}

type IOUSBHostCIDeviceCreateCommandData1Device uint32

const (
	IOUSBHostCIDeviceCreateCommandData1DeviceAddress      IOUSBHostCIDeviceCreateCommandData1Device = 255
	IOUSBHostCIDeviceCreateCommandData1DeviceAddressPhase IOUSBHostCIDeviceCreateCommandData1Device = 0
)

func (e IOUSBHostCIDeviceCreateCommandData1Device) String() string {
	switch e {
	case IOUSBHostCIDeviceCreateCommandData1DeviceAddress:
		return "IOUSBHostCIDeviceCreateCommandData1DeviceAddress"
	case IOUSBHostCIDeviceCreateCommandData1DeviceAddressPhase:
		return "IOUSBHostCIDeviceCreateCommandData1DeviceAddressPhase"
	default:
		return fmt.Sprintf("IOUSBHostCIDeviceCreateCommandData1Device(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIDeviceSpeed
type IOUSBHostCIDeviceSpeed uint32

const (
	IOUSBHostCIDeviceSpeedFull         IOUSBHostCIDeviceSpeed = 1
	IOUSBHostCIDeviceSpeedHigh         IOUSBHostCIDeviceSpeed = 3
	IOUSBHostCIDeviceSpeedLow          IOUSBHostCIDeviceSpeed = 2
	IOUSBHostCIDeviceSpeedNone         IOUSBHostCIDeviceSpeed = 0
	IOUSBHostCIDeviceSpeedSuper        IOUSBHostCIDeviceSpeed = 4
	IOUSBHostCIDeviceSpeedSuperPlus    IOUSBHostCIDeviceSpeed = 5
	IOUSBHostCIDeviceSpeedSuperPlusBy2 IOUSBHostCIDeviceSpeed = 6
)

func (e IOUSBHostCIDeviceSpeed) String() string {
	switch e {
	case IOUSBHostCIDeviceSpeedFull:
		return "IOUSBHostCIDeviceSpeedFull"
	case IOUSBHostCIDeviceSpeedHigh:
		return "IOUSBHostCIDeviceSpeedHigh"
	case IOUSBHostCIDeviceSpeedLow:
		return "IOUSBHostCIDeviceSpeedLow"
	case IOUSBHostCIDeviceSpeedNone:
		return "IOUSBHostCIDeviceSpeedNone"
	case IOUSBHostCIDeviceSpeedSuper:
		return "IOUSBHostCIDeviceSpeedSuper"
	case IOUSBHostCIDeviceSpeedSuperPlus:
		return "IOUSBHostCIDeviceSpeedSuperPlus"
	case IOUSBHostCIDeviceSpeedSuperPlusBy2:
		return "IOUSBHostCIDeviceSpeedSuperPlusBy2"
	default:
		return fmt.Sprintf("IOUSBHostCIDeviceSpeed(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIDeviceState
type IOUSBHostCIDeviceState uint32

const (
	IOUSBHostCIDeviceStateActive    IOUSBHostCIDeviceState = 2
	IOUSBHostCIDeviceStateDestroyed IOUSBHostCIDeviceState = 0
	IOUSBHostCIDeviceStatePaused    IOUSBHostCIDeviceState = 1
)

func (e IOUSBHostCIDeviceState) String() string {
	switch e {
	case IOUSBHostCIDeviceStateActive:
		return "IOUSBHostCIDeviceStateActive"
	case IOUSBHostCIDeviceStateDestroyed:
		return "IOUSBHostCIDeviceStateDestroyed"
	case IOUSBHostCIDeviceStatePaused:
		return "IOUSBHostCIDeviceStatePaused"
	default:
		return fmt.Sprintf("IOUSBHostCIDeviceState(%d)", e)
	}
}

type IOUSBHostCIDeviceUpdateCommandData1Descriptor uint

const (
	IOUSBHostCIDeviceUpdateCommandData1DescriptorAddress      IOUSBHostCIDeviceUpdateCommandData1Descriptor = 18446744073709551615
	IOUSBHostCIDeviceUpdateCommandData1DescriptorAddressPhase IOUSBHostCIDeviceUpdateCommandData1Descriptor = 0
)

func (e IOUSBHostCIDeviceUpdateCommandData1Descriptor) String() string {
	switch e {
	case IOUSBHostCIDeviceUpdateCommandData1DescriptorAddress:
		return "IOUSBHostCIDeviceUpdateCommandData1DescriptorAddress"
	case IOUSBHostCIDeviceUpdateCommandData1DescriptorAddressPhase:
		return "IOUSBHostCIDeviceUpdateCommandData1DescriptorAddressPhase"
	default:
		return fmt.Sprintf("IOUSBHostCIDeviceUpdateCommandData1Descriptor(%d)", e)
	}
}

type IOUSBHostCIDoorbell uint32

const (
	IOUSBHostCIDoorbellDeviceAddress        IOUSBHostCIDoorbell = 255
	IOUSBHostCIDoorbellDeviceAddressPhase   IOUSBHostCIDoorbell = 0
	IOUSBHostCIDoorbellEndpointAddress      IOUSBHostCIDoorbell = 65280
	IOUSBHostCIDoorbellEndpointAddressPhase IOUSBHostCIDoorbell = 8
	IOUSBHostCIDoorbellStreamID             IOUSBHostCIDoorbell = 4294901760
	IOUSBHostCIDoorbellStreamIDPhase        IOUSBHostCIDoorbell = 16
)

func (e IOUSBHostCIDoorbell) String() string {
	switch e {
	case IOUSBHostCIDoorbellDeviceAddress:
		return "IOUSBHostCIDoorbellDeviceAddress"
	case IOUSBHostCIDoorbellDeviceAddressPhase:
		return "IOUSBHostCIDoorbellDeviceAddressPhase"
	case IOUSBHostCIDoorbellEndpointAddress:
		return "IOUSBHostCIDoorbellEndpointAddress"
	case IOUSBHostCIDoorbellEndpointAddressPhase:
		return "IOUSBHostCIDoorbellEndpointAddressPhase"
	case IOUSBHostCIDoorbellStreamID:
		return "IOUSBHostCIDoorbellStreamID"
	case IOUSBHostCIDoorbellStreamIDPhase:
		return "IOUSBHostCIDoorbellStreamIDPhase"
	default:
		return fmt.Sprintf("IOUSBHostCIDoorbell(%d)", e)
	}
}

type IOUSBHostCIEndpointCreateCommandData1 uint

const (
	IOUSBHostCIEndpointCreateCommandData1Descriptor      IOUSBHostCIEndpointCreateCommandData1 = 18446744073709551615
	IOUSBHostCIEndpointCreateCommandData1DescriptorPhase IOUSBHostCIEndpointCreateCommandData1 = 0
)

func (e IOUSBHostCIEndpointCreateCommandData1) String() string {
	switch e {
	case IOUSBHostCIEndpointCreateCommandData1Descriptor:
		return "IOUSBHostCIEndpointCreateCommandData1Descriptor"
	case IOUSBHostCIEndpointCreateCommandData1DescriptorPhase:
		return "IOUSBHostCIEndpointCreateCommandData1DescriptorPhase"
	default:
		return fmt.Sprintf("IOUSBHostCIEndpointCreateCommandData1(%d)", e)
	}
}

type IOUSBHostCIEndpointResetCommandData1ClearStateConstants uint32

const (
	IOUSBHostCIEndpointResetCommandData1ClearState IOUSBHostCIEndpointResetCommandData1ClearStateConstants = 1
)

func (e IOUSBHostCIEndpointResetCommandData1ClearStateConstants) String() string {
	switch e {
	case IOUSBHostCIEndpointResetCommandData1ClearState:
		return "IOUSBHostCIEndpointResetCommandData1ClearState"
	default:
		return fmt.Sprintf("IOUSBHostCIEndpointResetCommandData1ClearStateConstants(%d)", e)
	}
}

type IOUSBHostCIEndpointSetNextTransferCommandData1 uint

const (
	IOUSBHostCIEndpointSetNextTransferCommandData1Address      IOUSBHostCIEndpointSetNextTransferCommandData1 = 18446744073709551615
	IOUSBHostCIEndpointSetNextTransferCommandData1AddressPhase IOUSBHostCIEndpointSetNextTransferCommandData1 = 0
)

func (e IOUSBHostCIEndpointSetNextTransferCommandData1) String() string {
	switch e {
	case IOUSBHostCIEndpointSetNextTransferCommandData1Address:
		return "IOUSBHostCIEndpointSetNextTransferCommandData1Address"
	case IOUSBHostCIEndpointSetNextTransferCommandData1AddressPhase:
		return "IOUSBHostCIEndpointSetNextTransferCommandData1AddressPhase"
	default:
		return fmt.Sprintf("IOUSBHostCIEndpointSetNextTransferCommandData1(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIEndpointState
type IOUSBHostCIEndpointState uint32

const (
	IOUSBHostCIEndpointStateActive    IOUSBHostCIEndpointState = 3
	IOUSBHostCIEndpointStateDestroyed IOUSBHostCIEndpointState = 0
	IOUSBHostCIEndpointStateHalted    IOUSBHostCIEndpointState = 1
	IOUSBHostCIEndpointStatePaused    IOUSBHostCIEndpointState = 2
)

func (e IOUSBHostCIEndpointState) String() string {
	switch e {
	case IOUSBHostCIEndpointStateActive:
		return "IOUSBHostCIEndpointStateActive"
	case IOUSBHostCIEndpointStateDestroyed:
		return "IOUSBHostCIEndpointStateDestroyed"
	case IOUSBHostCIEndpointStateHalted:
		return "IOUSBHostCIEndpointStateHalted"
	case IOUSBHostCIEndpointStatePaused:
		return "IOUSBHostCIEndpointStatePaused"
	default:
		return fmt.Sprintf("IOUSBHostCIEndpointState(%d)", e)
	}
}

type IOUSBHostCIEndpointUpdateCommandData1 uint

const (
	IOUSBHostCIEndpointUpdateCommandData1Descriptor      IOUSBHostCIEndpointUpdateCommandData1 = 18446744073709551615
	IOUSBHostCIEndpointUpdateCommandData1DescriptorPhase IOUSBHostCIEndpointUpdateCommandData1 = 0
)

func (e IOUSBHostCIEndpointUpdateCommandData1) String() string {
	switch e {
	case IOUSBHostCIEndpointUpdateCommandData1Descriptor:
		return "IOUSBHostCIEndpointUpdateCommandData1Descriptor"
	case IOUSBHostCIEndpointUpdateCommandData1DescriptorPhase:
		return "IOUSBHostCIEndpointUpdateCommandData1DescriptorPhase"
	default:
		return fmt.Sprintf("IOUSBHostCIEndpointUpdateCommandData1(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIExceptionType
type IOUSBHostCIExceptionType uint32

const (
	IOUSBHostCIExceptionTypeCapabilitiesInvalid   IOUSBHostCIExceptionType = 1
	IOUSBHostCIExceptionTypeCommandFailure        IOUSBHostCIExceptionType = 6
	IOUSBHostCIExceptionTypeCommandReadCollision  IOUSBHostCIExceptionType = 3
	IOUSBHostCIExceptionTypeCommandTimeout        IOUSBHostCIExceptionType = 5
	IOUSBHostCIExceptionTypeCommandWriteFailed    IOUSBHostCIExceptionType = 4
	IOUSBHostCIExceptionTypeDoorbellOverflow      IOUSBHostCIExceptionType = 10
	IOUSBHostCIExceptionTypeDoorbellReadCollision IOUSBHostCIExceptionType = 9
	IOUSBHostCIExceptionTypeFrameUpdateError      IOUSBHostCIExceptionType = 12
	IOUSBHostCIExceptionTypeInterruptInvalid      IOUSBHostCIExceptionType = 7
	IOUSBHostCIExceptionTypeInterruptOverflow     IOUSBHostCIExceptionType = 8
	IOUSBHostCIExceptionTypeProtocolError         IOUSBHostCIExceptionType = 11
	IOUSBHostCIExceptionTypeTerminated            IOUSBHostCIExceptionType = 2
	IOUSBHostCIExceptionTypeUnknown               IOUSBHostCIExceptionType = 0
)

func (e IOUSBHostCIExceptionType) String() string {
	switch e {
	case IOUSBHostCIExceptionTypeCapabilitiesInvalid:
		return "IOUSBHostCIExceptionTypeCapabilitiesInvalid"
	case IOUSBHostCIExceptionTypeCommandFailure:
		return "IOUSBHostCIExceptionTypeCommandFailure"
	case IOUSBHostCIExceptionTypeCommandReadCollision:
		return "IOUSBHostCIExceptionTypeCommandReadCollision"
	case IOUSBHostCIExceptionTypeCommandTimeout:
		return "IOUSBHostCIExceptionTypeCommandTimeout"
	case IOUSBHostCIExceptionTypeCommandWriteFailed:
		return "IOUSBHostCIExceptionTypeCommandWriteFailed"
	case IOUSBHostCIExceptionTypeDoorbellOverflow:
		return "IOUSBHostCIExceptionTypeDoorbellOverflow"
	case IOUSBHostCIExceptionTypeDoorbellReadCollision:
		return "IOUSBHostCIExceptionTypeDoorbellReadCollision"
	case IOUSBHostCIExceptionTypeFrameUpdateError:
		return "IOUSBHostCIExceptionTypeFrameUpdateError"
	case IOUSBHostCIExceptionTypeInterruptInvalid:
		return "IOUSBHostCIExceptionTypeInterruptInvalid"
	case IOUSBHostCIExceptionTypeInterruptOverflow:
		return "IOUSBHostCIExceptionTypeInterruptOverflow"
	case IOUSBHostCIExceptionTypeProtocolError:
		return "IOUSBHostCIExceptionTypeProtocolError"
	case IOUSBHostCIExceptionTypeTerminated:
		return "IOUSBHostCIExceptionTypeTerminated"
	case IOUSBHostCIExceptionTypeUnknown:
		return "IOUSBHostCIExceptionTypeUnknown"
	default:
		return fmt.Sprintf("IOUSBHostCIExceptionType(%d)", e)
	}
}

type IOUSBHostCIIsochronousTransferControl uint32

const (
	IOUSBHostCIIsochronousTransferControlASAP             IOUSBHostCIIsochronousTransferControl = 0x1000000
	IOUSBHostCIIsochronousTransferControlFrameNumber      IOUSBHostCIIsochronousTransferControl = 16711680
	IOUSBHostCIIsochronousTransferControlFrameNumberPhase IOUSBHostCIIsochronousTransferControl = 16
)

func (e IOUSBHostCIIsochronousTransferControl) String() string {
	switch e {
	case IOUSBHostCIIsochronousTransferControlASAP:
		return "IOUSBHostCIIsochronousTransferControlASAP"
	case IOUSBHostCIIsochronousTransferControlFrameNumber:
		return "IOUSBHostCIIsochronousTransferControlFrameNumber"
	case IOUSBHostCIIsochronousTransferControlFrameNumberPhase:
		return "IOUSBHostCIIsochronousTransferControlFrameNumberPhase"
	default:
		return fmt.Sprintf("IOUSBHostCIIsochronousTransferControl(%d)", e)
	}
}

type IOUSBHostCIIsochronousTransferData0 uint32

const (
	IOUSBHostCIIsochronousTransferData0Length      IOUSBHostCIIsochronousTransferData0 = 268435455
	IOUSBHostCIIsochronousTransferData0LengthPhase IOUSBHostCIIsochronousTransferData0 = 0
)

func (e IOUSBHostCIIsochronousTransferData0) String() string {
	switch e {
	case IOUSBHostCIIsochronousTransferData0Length:
		return "IOUSBHostCIIsochronousTransferData0Length"
	case IOUSBHostCIIsochronousTransferData0LengthPhase:
		return "IOUSBHostCIIsochronousTransferData0LengthPhase"
	default:
		return fmt.Sprintf("IOUSBHostCIIsochronousTransferData0(%d)", e)
	}
}

type IOUSBHostCIIsochronousTransferData1 uint

const (
	IOUSBHostCIIsochronousTransferData1Buffer      IOUSBHostCIIsochronousTransferData1 = 18446744073709551615
	IOUSBHostCIIsochronousTransferData1BufferPhase IOUSBHostCIIsochronousTransferData1 = 0
)

func (e IOUSBHostCIIsochronousTransferData1) String() string {
	switch e {
	case IOUSBHostCIIsochronousTransferData1Buffer:
		return "IOUSBHostCIIsochronousTransferData1Buffer"
	case IOUSBHostCIIsochronousTransferData1BufferPhase:
		return "IOUSBHostCIIsochronousTransferData1BufferPhase"
	default:
		return fmt.Sprintf("IOUSBHostCIIsochronousTransferData1(%d)", e)
	}
}

type IOUSBHostCILinkData1TransferStructure uint

const (
	IOUSBHostCILinkData1TransferStructureAddress      IOUSBHostCILinkData1TransferStructure = 18446744073709551615
	IOUSBHostCILinkData1TransferStructureAddressPhase IOUSBHostCILinkData1TransferStructure = 0
)

func (e IOUSBHostCILinkData1TransferStructure) String() string {
	switch e {
	case IOUSBHostCILinkData1TransferStructureAddress:
		return "IOUSBHostCILinkData1TransferStructureAddress"
	case IOUSBHostCILinkData1TransferStructureAddressPhase:
		return "IOUSBHostCILinkData1TransferStructureAddressPhase"
	default:
		return fmt.Sprintf("IOUSBHostCILinkData1TransferStructure(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCILinkState
type IOUSBHostCILinkState uint32

const (
	IOUSBHostCILinkStateCompliance IOUSBHostCILinkState = 10
	IOUSBHostCILinkStateDisabled   IOUSBHostCILinkState = 4
	IOUSBHostCILinkStateInactive   IOUSBHostCILinkState = 6
	IOUSBHostCILinkStatePolling    IOUSBHostCILinkState = 7
	IOUSBHostCILinkStateRecovery   IOUSBHostCILinkState = 8
	IOUSBHostCILinkStateReset      IOUSBHostCILinkState = 9
	IOUSBHostCILinkStateResume     IOUSBHostCILinkState = 15
	IOUSBHostCILinkStateRxDetect   IOUSBHostCILinkState = 5
	IOUSBHostCILinkStateTest       IOUSBHostCILinkState = 11
	IOUSBHostCILinkStateU0         IOUSBHostCILinkState = 0
	IOUSBHostCILinkStateU1         IOUSBHostCILinkState = 1
	IOUSBHostCILinkStateU2         IOUSBHostCILinkState = 2
	IOUSBHostCILinkStateU3         IOUSBHostCILinkState = 3
)

func (e IOUSBHostCILinkState) String() string {
	switch e {
	case IOUSBHostCILinkStateCompliance:
		return "IOUSBHostCILinkStateCompliance"
	case IOUSBHostCILinkStateDisabled:
		return "IOUSBHostCILinkStateDisabled"
	case IOUSBHostCILinkStateInactive:
		return "IOUSBHostCILinkStateInactive"
	case IOUSBHostCILinkStatePolling:
		return "IOUSBHostCILinkStatePolling"
	case IOUSBHostCILinkStateRecovery:
		return "IOUSBHostCILinkStateRecovery"
	case IOUSBHostCILinkStateReset:
		return "IOUSBHostCILinkStateReset"
	case IOUSBHostCILinkStateResume:
		return "IOUSBHostCILinkStateResume"
	case IOUSBHostCILinkStateRxDetect:
		return "IOUSBHostCILinkStateRxDetect"
	case IOUSBHostCILinkStateTest:
		return "IOUSBHostCILinkStateTest"
	case IOUSBHostCILinkStateU0:
		return "IOUSBHostCILinkStateU0"
	case IOUSBHostCILinkStateU1:
		return "IOUSBHostCILinkStateU1"
	case IOUSBHostCILinkStateU2:
		return "IOUSBHostCILinkStateU2"
	case IOUSBHostCILinkStateU3:
		return "IOUSBHostCILinkStateU3"
	default:
		return fmt.Sprintf("IOUSBHostCILinkState(%d)", e)
	}
}

type IOUSBHostCIMessageControl uint32

const (
	IOUSBHostCIMessageControlNoResponse  IOUSBHostCIMessageControl = 0x4000
	IOUSBHostCIMessageControlStatus      IOUSBHostCIMessageControl = 3840
	IOUSBHostCIMessageControlStatusPhase IOUSBHostCIMessageControl = 8
	IOUSBHostCIMessageControlType        IOUSBHostCIMessageControl = 63
	IOUSBHostCIMessageControlTypePhase   IOUSBHostCIMessageControl = 0
	IOUSBHostCIMessageControlValid       IOUSBHostCIMessageControl = 0x8000
)

func (e IOUSBHostCIMessageControl) String() string {
	switch e {
	case IOUSBHostCIMessageControlNoResponse:
		return "IOUSBHostCIMessageControlNoResponse"
	case IOUSBHostCIMessageControlStatus:
		return "IOUSBHostCIMessageControlStatus"
	case IOUSBHostCIMessageControlStatusPhase:
		return "IOUSBHostCIMessageControlStatusPhase"
	case IOUSBHostCIMessageControlType:
		return "IOUSBHostCIMessageControlType"
	case IOUSBHostCIMessageControlTypePhase:
		return "IOUSBHostCIMessageControlTypePhase"
	case IOUSBHostCIMessageControlValid:
		return "IOUSBHostCIMessageControlValid"
	default:
		return fmt.Sprintf("IOUSBHostCIMessageControl(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIMessageStatus
type IOUSBHostCIMessageStatus uint32

const (
	IOUSBHostCIMessageStatusBadArgument        IOUSBHostCIMessageStatus = 4
	IOUSBHostCIMessageStatusEndpointStopped    IOUSBHostCIMessageStatus = 7
	IOUSBHostCIMessageStatusError              IOUSBHostCIMessageStatus = 13
	IOUSBHostCIMessageStatusMissedServiceError IOUSBHostCIMessageStatus = 12
	IOUSBHostCIMessageStatusNoResources        IOUSBHostCIMessageStatus = 6
	IOUSBHostCIMessageStatusNotPermitted       IOUSBHostCIMessageStatus = 3
	IOUSBHostCIMessageStatusOffline            IOUSBHostCIMessageStatus = 2
	IOUSBHostCIMessageStatusOverrunError       IOUSBHostCIMessageStatus = 10
	IOUSBHostCIMessageStatusProtocolError      IOUSBHostCIMessageStatus = 8
	IOUSBHostCIMessageStatusReserved           IOUSBHostCIMessageStatus = 0
	IOUSBHostCIMessageStatusStallError         IOUSBHostCIMessageStatus = 11
	IOUSBHostCIMessageStatusSuccess            IOUSBHostCIMessageStatus = 1
	IOUSBHostCIMessageStatusTimeout            IOUSBHostCIMessageStatus = 5
	IOUSBHostCIMessageStatusTransactionError   IOUSBHostCIMessageStatus = 9
)

func (e IOUSBHostCIMessageStatus) String() string {
	switch e {
	case IOUSBHostCIMessageStatusBadArgument:
		return "IOUSBHostCIMessageStatusBadArgument"
	case IOUSBHostCIMessageStatusEndpointStopped:
		return "IOUSBHostCIMessageStatusEndpointStopped"
	case IOUSBHostCIMessageStatusError:
		return "IOUSBHostCIMessageStatusError"
	case IOUSBHostCIMessageStatusMissedServiceError:
		return "IOUSBHostCIMessageStatusMissedServiceError"
	case IOUSBHostCIMessageStatusNoResources:
		return "IOUSBHostCIMessageStatusNoResources"
	case IOUSBHostCIMessageStatusNotPermitted:
		return "IOUSBHostCIMessageStatusNotPermitted"
	case IOUSBHostCIMessageStatusOffline:
		return "IOUSBHostCIMessageStatusOffline"
	case IOUSBHostCIMessageStatusOverrunError:
		return "IOUSBHostCIMessageStatusOverrunError"
	case IOUSBHostCIMessageStatusProtocolError:
		return "IOUSBHostCIMessageStatusProtocolError"
	case IOUSBHostCIMessageStatusReserved:
		return "IOUSBHostCIMessageStatusReserved"
	case IOUSBHostCIMessageStatusStallError:
		return "IOUSBHostCIMessageStatusStallError"
	case IOUSBHostCIMessageStatusSuccess:
		return "IOUSBHostCIMessageStatusSuccess"
	case IOUSBHostCIMessageStatusTimeout:
		return "IOUSBHostCIMessageStatusTimeout"
	case IOUSBHostCIMessageStatusTransactionError:
		return "IOUSBHostCIMessageStatusTransactionError"
	default:
		return fmt.Sprintf("IOUSBHostCIMessageStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIMessageType
type IOUSBHostCIMessageType uint32

const (
	IOUSBHostCIMessageTypeCommandMax              IOUSBHostCIMessageType = 0x37
	IOUSBHostCIMessageTypeCommandMin              IOUSBHostCIMessageType = 0x10
	IOUSBHostCIMessageTypeControllerCapabilities  IOUSBHostCIMessageType = 0
	IOUSBHostCIMessageTypeControllerFrameNumber   IOUSBHostCIMessageType = 20
	IOUSBHostCIMessageTypeControllerPause         IOUSBHostCIMessageType = 19
	IOUSBHostCIMessageTypeControllerPowerOff      IOUSBHostCIMessageType = 17
	IOUSBHostCIMessageTypeControllerPowerOn       IOUSBHostCIMessageType = 0x10
	IOUSBHostCIMessageTypeControllerStart         IOUSBHostCIMessageType = 18
	IOUSBHostCIMessageTypeDeviceCreate            IOUSBHostCIMessageType = 0x20
	IOUSBHostCIMessageTypeDeviceDestroy           IOUSBHostCIMessageType = 33
	IOUSBHostCIMessageTypeDevicePause             IOUSBHostCIMessageType = 35
	IOUSBHostCIMessageTypeDeviceStart             IOUSBHostCIMessageType = 34
	IOUSBHostCIMessageTypeDeviceUpdate            IOUSBHostCIMessageType = 36
	IOUSBHostCIMessageTypeEndpointCreate          IOUSBHostCIMessageType = 0x28
	IOUSBHostCIMessageTypeEndpointDestroy         IOUSBHostCIMessageType = 41
	IOUSBHostCIMessageTypeEndpointPause           IOUSBHostCIMessageType = 43
	IOUSBHostCIMessageTypeEndpointReset           IOUSBHostCIMessageType = 45
	IOUSBHostCIMessageTypeEndpointSetNextTransfer IOUSBHostCIMessageType = 46
	IOUSBHostCIMessageTypeEndpointUpdate          IOUSBHostCIMessageType = 44
	IOUSBHostCIMessageTypeEndpoint_reserved_      IOUSBHostCIMessageType = 42
	IOUSBHostCIMessageTypeFrameNumberUpdate       IOUSBHostCIMessageType = 9
	IOUSBHostCIMessageTypeFrameTimestampUpdate    IOUSBHostCIMessageType = 10
	IOUSBHostCIMessageTypeIsochronousTransfer     IOUSBHostCIMessageType = 59
	IOUSBHostCIMessageTypeLink                    IOUSBHostCIMessageType = 60
	IOUSBHostCIMessageTypeNormalTransfer          IOUSBHostCIMessageType = 57
	IOUSBHostCIMessageTypePortCapabilities        IOUSBHostCIMessageType = 1
	IOUSBHostCIMessageTypePortDisable             IOUSBHostCIMessageType = 29
	IOUSBHostCIMessageTypePortEvent               IOUSBHostCIMessageType = 0x8
	IOUSBHostCIMessageTypePortPowerOff            IOUSBHostCIMessageType = 25
	IOUSBHostCIMessageTypePortPowerOn             IOUSBHostCIMessageType = 0x18
	IOUSBHostCIMessageTypePortReset               IOUSBHostCIMessageType = 28
	IOUSBHostCIMessageTypePortResume              IOUSBHostCIMessageType = 26
	IOUSBHostCIMessageTypePortStatus              IOUSBHostCIMessageType = 30
	IOUSBHostCIMessageTypePortSuspend             IOUSBHostCIMessageType = 27
	IOUSBHostCIMessageTypeSetupTransfer           IOUSBHostCIMessageType = 0x38
	IOUSBHostCIMessageTypeStatusTransfer          IOUSBHostCIMessageType = 58
	IOUSBHostCIMessageTypeTransferComplete        IOUSBHostCIMessageType = 61
)

func (e IOUSBHostCIMessageType) String() string {
	switch e {
	case IOUSBHostCIMessageTypeCommandMax:
		return "IOUSBHostCIMessageTypeCommandMax"
	case IOUSBHostCIMessageTypeCommandMin:
		return "IOUSBHostCIMessageTypeCommandMin"
	case IOUSBHostCIMessageTypeControllerCapabilities:
		return "IOUSBHostCIMessageTypeControllerCapabilities"
	case IOUSBHostCIMessageTypeControllerFrameNumber:
		return "IOUSBHostCIMessageTypeControllerFrameNumber"
	case IOUSBHostCIMessageTypeControllerPause:
		return "IOUSBHostCIMessageTypeControllerPause"
	case IOUSBHostCIMessageTypeControllerPowerOff:
		return "IOUSBHostCIMessageTypeControllerPowerOff"
	case IOUSBHostCIMessageTypeControllerStart:
		return "IOUSBHostCIMessageTypeControllerStart"
	case IOUSBHostCIMessageTypeDeviceCreate:
		return "IOUSBHostCIMessageTypeDeviceCreate"
	case IOUSBHostCIMessageTypeDeviceDestroy:
		return "IOUSBHostCIMessageTypeDeviceDestroy"
	case IOUSBHostCIMessageTypeDevicePause:
		return "IOUSBHostCIMessageTypeDevicePause"
	case IOUSBHostCIMessageTypeDeviceStart:
		return "IOUSBHostCIMessageTypeDeviceStart"
	case IOUSBHostCIMessageTypeDeviceUpdate:
		return "IOUSBHostCIMessageTypeDeviceUpdate"
	case IOUSBHostCIMessageTypeEndpointCreate:
		return "IOUSBHostCIMessageTypeEndpointCreate"
	case IOUSBHostCIMessageTypeEndpointDestroy:
		return "IOUSBHostCIMessageTypeEndpointDestroy"
	case IOUSBHostCIMessageTypeEndpointPause:
		return "IOUSBHostCIMessageTypeEndpointPause"
	case IOUSBHostCIMessageTypeEndpointReset:
		return "IOUSBHostCIMessageTypeEndpointReset"
	case IOUSBHostCIMessageTypeEndpointSetNextTransfer:
		return "IOUSBHostCIMessageTypeEndpointSetNextTransfer"
	case IOUSBHostCIMessageTypeEndpointUpdate:
		return "IOUSBHostCIMessageTypeEndpointUpdate"
	case IOUSBHostCIMessageTypeEndpoint_reserved_:
		return "IOUSBHostCIMessageTypeEndpoint_reserved_"
	case IOUSBHostCIMessageTypeFrameNumberUpdate:
		return "IOUSBHostCIMessageTypeFrameNumberUpdate"
	case IOUSBHostCIMessageTypeFrameTimestampUpdate:
		return "IOUSBHostCIMessageTypeFrameTimestampUpdate"
	case IOUSBHostCIMessageTypeIsochronousTransfer:
		return "IOUSBHostCIMessageTypeIsochronousTransfer"
	case IOUSBHostCIMessageTypeLink:
		return "IOUSBHostCIMessageTypeLink"
	case IOUSBHostCIMessageTypeNormalTransfer:
		return "IOUSBHostCIMessageTypeNormalTransfer"
	case IOUSBHostCIMessageTypePortCapabilities:
		return "IOUSBHostCIMessageTypePortCapabilities"
	case IOUSBHostCIMessageTypePortDisable:
		return "IOUSBHostCIMessageTypePortDisable"
	case IOUSBHostCIMessageTypePortEvent:
		return "IOUSBHostCIMessageTypePortEvent"
	case IOUSBHostCIMessageTypePortPowerOff:
		return "IOUSBHostCIMessageTypePortPowerOff"
	case IOUSBHostCIMessageTypePortPowerOn:
		return "IOUSBHostCIMessageTypePortPowerOn"
	case IOUSBHostCIMessageTypePortReset:
		return "IOUSBHostCIMessageTypePortReset"
	case IOUSBHostCIMessageTypePortResume:
		return "IOUSBHostCIMessageTypePortResume"
	case IOUSBHostCIMessageTypePortStatus:
		return "IOUSBHostCIMessageTypePortStatus"
	case IOUSBHostCIMessageTypePortSuspend:
		return "IOUSBHostCIMessageTypePortSuspend"
	case IOUSBHostCIMessageTypeSetupTransfer:
		return "IOUSBHostCIMessageTypeSetupTransfer"
	case IOUSBHostCIMessageTypeStatusTransfer:
		return "IOUSBHostCIMessageTypeStatusTransfer"
	case IOUSBHostCIMessageTypeTransferComplete:
		return "IOUSBHostCIMessageTypeTransferComplete"
	default:
		return fmt.Sprintf("IOUSBHostCIMessageType(%d)", e)
	}
}

type IOUSBHostCINormalTransferData0 uint32

const (
	IOUSBHostCINormalTransferData0Length      IOUSBHostCINormalTransferData0 = 268435455
	IOUSBHostCINormalTransferData0LengthPhase IOUSBHostCINormalTransferData0 = 0
)

func (e IOUSBHostCINormalTransferData0) String() string {
	switch e {
	case IOUSBHostCINormalTransferData0Length:
		return "IOUSBHostCINormalTransferData0Length"
	case IOUSBHostCINormalTransferData0LengthPhase:
		return "IOUSBHostCINormalTransferData0LengthPhase"
	default:
		return fmt.Sprintf("IOUSBHostCINormalTransferData0(%d)", e)
	}
}

type IOUSBHostCINormalTransferData1 uint

const (
	IOUSBHostCINormalTransferData1Buffer      IOUSBHostCINormalTransferData1 = 18446744073709551615
	IOUSBHostCINormalTransferData1BufferPhase IOUSBHostCINormalTransferData1 = 0
)

func (e IOUSBHostCINormalTransferData1) String() string {
	switch e {
	case IOUSBHostCINormalTransferData1Buffer:
		return "IOUSBHostCINormalTransferData1Buffer"
	case IOUSBHostCINormalTransferData1BufferPhase:
		return "IOUSBHostCINormalTransferData1BufferPhase"
	default:
		return fmt.Sprintf("IOUSBHostCINormalTransferData1(%d)", e)
	}
}

type IOUSBHostCIPortCapabilitiesMessage uint32

const (
	IOUSBHostCIPortCapabilitiesMessageControlConnectorType      IOUSBHostCIPortCapabilitiesMessage = 4278190080
	IOUSBHostCIPortCapabilitiesMessageControlConnectorTypePhase IOUSBHostCIPortCapabilitiesMessage = 24
	IOUSBHostCIPortCapabilitiesMessageControlInternalConnector  IOUSBHostCIPortCapabilitiesMessage = 0x800000
	IOUSBHostCIPortCapabilitiesMessageControlPortNumber         IOUSBHostCIPortCapabilitiesMessage = 983040
	IOUSBHostCIPortCapabilitiesMessageControlPortNumberPhase    IOUSBHostCIPortCapabilitiesMessage = 16
	IOUSBHostCIPortCapabilitiesMessageData0MaxPower             IOUSBHostCIPortCapabilitiesMessage = 255
	IOUSBHostCIPortCapabilitiesMessageData0MaxPowerPhase        IOUSBHostCIPortCapabilitiesMessage = 0
)

func (e IOUSBHostCIPortCapabilitiesMessage) String() string {
	switch e {
	case IOUSBHostCIPortCapabilitiesMessageControlConnectorType:
		return "IOUSBHostCIPortCapabilitiesMessageControlConnectorType"
	case IOUSBHostCIPortCapabilitiesMessageControlConnectorTypePhase:
		return "IOUSBHostCIPortCapabilitiesMessageControlConnectorTypePhase"
	case IOUSBHostCIPortCapabilitiesMessageControlInternalConnector:
		return "IOUSBHostCIPortCapabilitiesMessageControlInternalConnector"
	case IOUSBHostCIPortCapabilitiesMessageControlPortNumber:
		return "IOUSBHostCIPortCapabilitiesMessageControlPortNumber"
	case IOUSBHostCIPortCapabilitiesMessageControlPortNumberPhase:
		return "IOUSBHostCIPortCapabilitiesMessageControlPortNumberPhase"
	case IOUSBHostCIPortCapabilitiesMessageData0MaxPower:
		return "IOUSBHostCIPortCapabilitiesMessageData0MaxPower"
	case IOUSBHostCIPortCapabilitiesMessageData0MaxPowerPhase:
		return "IOUSBHostCIPortCapabilitiesMessageData0MaxPowerPhase"
	default:
		return fmt.Sprintf("IOUSBHostCIPortCapabilitiesMessage(%d)", e)
	}
}

type IOUSBHostCIPortEventMessageData0Port uint32

const (
	IOUSBHostCIPortEventMessageData0PortNumber      IOUSBHostCIPortEventMessageData0Port = 15
	IOUSBHostCIPortEventMessageData0PortNumberPhase IOUSBHostCIPortEventMessageData0Port = 0
)

func (e IOUSBHostCIPortEventMessageData0Port) String() string {
	switch e {
	case IOUSBHostCIPortEventMessageData0PortNumber:
		return "IOUSBHostCIPortEventMessageData0PortNumber"
	case IOUSBHostCIPortEventMessageData0PortNumberPhase:
		return "IOUSBHostCIPortEventMessageData0PortNumberPhase"
	default:
		return fmt.Sprintf("IOUSBHostCIPortEventMessageData0Port(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortState
type IOUSBHostCIPortState uint32

const (
	IOUSBHostCIPortStateActive    IOUSBHostCIPortState = 3
	IOUSBHostCIPortStateOff       IOUSBHostCIPortState = 0
	IOUSBHostCIPortStatePowered   IOUSBHostCIPortState = 1
	IOUSBHostCIPortStateSuspended IOUSBHostCIPortState = 2
)

func (e IOUSBHostCIPortState) String() string {
	switch e {
	case IOUSBHostCIPortStateActive:
		return "IOUSBHostCIPortStateActive"
	case IOUSBHostCIPortStateOff:
		return "IOUSBHostCIPortStateOff"
	case IOUSBHostCIPortStatePowered:
		return "IOUSBHostCIPortStatePowered"
	case IOUSBHostCIPortStateSuspended:
		return "IOUSBHostCIPortStateSuspended"
	default:
		return fmt.Sprintf("IOUSBHostCIPortState(%d)", e)
	}
}

type IOUSBHostCIPortStatus uint32

const (
	IOUSBHostCIPortStatusChangeMask        IOUSBHostCIPortStatus = 1441792
	IOUSBHostCIPortStatusConnectChange     IOUSBHostCIPortStatus = 0x40000
	IOUSBHostCIPortStatusConnected         IOUSBHostCIPortStatus = 4
	IOUSBHostCIPortStatusLinkState         IOUSBHostCIPortStatus = 240
	IOUSBHostCIPortStatusLinkStateChange   IOUSBHostCIPortStatus = 0x100000
	IOUSBHostCIPortStatusLinkStatePhase    IOUSBHostCIPortStatus = 4
	IOUSBHostCIPortStatusOvercurrent       IOUSBHostCIPortStatus = 2
	IOUSBHostCIPortStatusOvercurrentChange IOUSBHostCIPortStatus = 0x20000
	IOUSBHostCIPortStatusPowered           IOUSBHostCIPortStatus = 1
	IOUSBHostCIPortStatusSpeed             IOUSBHostCIPortStatus = 1792
	IOUSBHostCIPortStatusSpeedPhase        IOUSBHostCIPortStatus = 8
)

func (e IOUSBHostCIPortStatus) String() string {
	switch e {
	case IOUSBHostCIPortStatusChangeMask:
		return "IOUSBHostCIPortStatusChangeMask"
	case IOUSBHostCIPortStatusConnectChange:
		return "IOUSBHostCIPortStatusConnectChange"
	case IOUSBHostCIPortStatusConnected:
		return "IOUSBHostCIPortStatusConnected"
	case IOUSBHostCIPortStatusLinkState:
		return "IOUSBHostCIPortStatusLinkState"
	case IOUSBHostCIPortStatusLinkStateChange:
		return "IOUSBHostCIPortStatusLinkStateChange"
	case IOUSBHostCIPortStatusOvercurrent:
		return "IOUSBHostCIPortStatusOvercurrent"
	case IOUSBHostCIPortStatusOvercurrentChange:
		return "IOUSBHostCIPortStatusOvercurrentChange"
	case IOUSBHostCIPortStatusPowered:
		return "IOUSBHostCIPortStatusPowered"
	case IOUSBHostCIPortStatusSpeed:
		return "IOUSBHostCIPortStatusSpeed"
	case IOUSBHostCIPortStatusSpeedPhase:
		return "IOUSBHostCIPortStatusSpeedPhase"
	default:
		return fmt.Sprintf("IOUSBHostCIPortStatus(%d)", e)
	}
}

type IOUSBHostCIPortStatusCommandData1 uint32

const (
	IOUSBHostCIPortStatusCommandData1ChangeMask        IOUSBHostCIPortStatusCommandData1 = 1441792
	IOUSBHostCIPortStatusCommandData1ConnectChange     IOUSBHostCIPortStatusCommandData1 = 262144
	IOUSBHostCIPortStatusCommandData1Connected         IOUSBHostCIPortStatusCommandData1 = 4
	IOUSBHostCIPortStatusCommandData1LinkState         IOUSBHostCIPortStatusCommandData1 = 240
	IOUSBHostCIPortStatusCommandData1LinkStateChange   IOUSBHostCIPortStatusCommandData1 = 1048576
	IOUSBHostCIPortStatusCommandData1LinkStatePhase    IOUSBHostCIPortStatusCommandData1 = 4
	IOUSBHostCIPortStatusCommandData1Overcurrent       IOUSBHostCIPortStatusCommandData1 = 2
	IOUSBHostCIPortStatusCommandData1OvercurrentChange IOUSBHostCIPortStatusCommandData1 = 131072
	IOUSBHostCIPortStatusCommandData1Powered           IOUSBHostCIPortStatusCommandData1 = 1
	IOUSBHostCIPortStatusCommandData1Speed             IOUSBHostCIPortStatusCommandData1 = 1792
	IOUSBHostCIPortStatusCommandData1SpeedPhase        IOUSBHostCIPortStatusCommandData1 = 8
)

func (e IOUSBHostCIPortStatusCommandData1) String() string {
	switch e {
	case IOUSBHostCIPortStatusCommandData1ChangeMask:
		return "IOUSBHostCIPortStatusCommandData1ChangeMask"
	case IOUSBHostCIPortStatusCommandData1ConnectChange:
		return "IOUSBHostCIPortStatusCommandData1ConnectChange"
	case IOUSBHostCIPortStatusCommandData1Connected:
		return "IOUSBHostCIPortStatusCommandData1Connected"
	case IOUSBHostCIPortStatusCommandData1LinkState:
		return "IOUSBHostCIPortStatusCommandData1LinkState"
	case IOUSBHostCIPortStatusCommandData1LinkStateChange:
		return "IOUSBHostCIPortStatusCommandData1LinkStateChange"
	case IOUSBHostCIPortStatusCommandData1Overcurrent:
		return "IOUSBHostCIPortStatusCommandData1Overcurrent"
	case IOUSBHostCIPortStatusCommandData1OvercurrentChange:
		return "IOUSBHostCIPortStatusCommandData1OvercurrentChange"
	case IOUSBHostCIPortStatusCommandData1Powered:
		return "IOUSBHostCIPortStatusCommandData1Powered"
	case IOUSBHostCIPortStatusCommandData1Speed:
		return "IOUSBHostCIPortStatusCommandData1Speed"
	case IOUSBHostCIPortStatusCommandData1SpeedPhase:
		return "IOUSBHostCIPortStatusCommandData1SpeedPhase"
	default:
		return fmt.Sprintf("IOUSBHostCIPortStatusCommandData1(%d)", e)
	}
}

type IOUSBHostCISetupTransfer uint

const (
	IOUSBHostCISetupTransferData1bRequest           IOUSBHostCISetupTransfer = 65280
	IOUSBHostCISetupTransferData1bRequestPhase      IOUSBHostCISetupTransfer = 8
	IOUSBHostCISetupTransferData1bmRequestType      IOUSBHostCISetupTransfer = 255
	IOUSBHostCISetupTransferData1bmRequestTypePhase IOUSBHostCISetupTransfer = 0
	IOUSBHostCISetupTransferData1wIndex             IOUSBHostCISetupTransfer = 281470681743360
	IOUSBHostCISetupTransferData1wIndexPhase        IOUSBHostCISetupTransfer = 32
	IOUSBHostCISetupTransferData1wLength            IOUSBHostCISetupTransfer = 18446462598732840960
	IOUSBHostCISetupTransferData1wLengthPhase       IOUSBHostCISetupTransfer = 48
	IOUSBHostCISetupTransferData1wValue             IOUSBHostCISetupTransfer = 4294901760
	IOUSBHostCISetupTransferData1wValuePhase        IOUSBHostCISetupTransfer = 16
)

func (e IOUSBHostCISetupTransfer) String() string {
	switch e {
	case IOUSBHostCISetupTransferData1bRequest:
		return "IOUSBHostCISetupTransferData1bRequest"
	case IOUSBHostCISetupTransferData1bRequestPhase:
		return "IOUSBHostCISetupTransferData1bRequestPhase"
	case IOUSBHostCISetupTransferData1bmRequestType:
		return "IOUSBHostCISetupTransferData1bmRequestType"
	case IOUSBHostCISetupTransferData1bmRequestTypePhase:
		return "IOUSBHostCISetupTransferData1bmRequestTypePhase"
	case IOUSBHostCISetupTransferData1wIndex:
		return "IOUSBHostCISetupTransferData1wIndex"
	case IOUSBHostCISetupTransferData1wIndexPhase:
		return "IOUSBHostCISetupTransferData1wIndexPhase"
	case IOUSBHostCISetupTransferData1wLength:
		return "IOUSBHostCISetupTransferData1wLength"
	case IOUSBHostCISetupTransferData1wLengthPhase:
		return "IOUSBHostCISetupTransferData1wLengthPhase"
	case IOUSBHostCISetupTransferData1wValue:
		return "IOUSBHostCISetupTransferData1wValue"
	case IOUSBHostCISetupTransferData1wValuePhase:
		return "IOUSBHostCISetupTransferData1wValuePhase"
	default:
		return fmt.Sprintf("IOUSBHostCISetupTransfer(%d)", e)
	}
}

type IOUSBHostCITransferCompletionMessageControl uint32

const (
	IOUSBHostCITransferCompletionMessageControlDeviceAddress        IOUSBHostCITransferCompletionMessageControl = 16711680
	IOUSBHostCITransferCompletionMessageControlDeviceAddressPhase   IOUSBHostCITransferCompletionMessageControl = 16
	IOUSBHostCITransferCompletionMessageControlEndpointAddress      IOUSBHostCITransferCompletionMessageControl = 4278190080
	IOUSBHostCITransferCompletionMessageControlEndpointAddressPhase IOUSBHostCITransferCompletionMessageControl = 24
	IOUSBHostCITransferCompletionMessageControlStatus               IOUSBHostCITransferCompletionMessageControl = 3840
	IOUSBHostCITransferCompletionMessageControlStatusPhase          IOUSBHostCITransferCompletionMessageControl = 8
)

func (e IOUSBHostCITransferCompletionMessageControl) String() string {
	switch e {
	case IOUSBHostCITransferCompletionMessageControlDeviceAddress:
		return "IOUSBHostCITransferCompletionMessageControlDeviceAddress"
	case IOUSBHostCITransferCompletionMessageControlDeviceAddressPhase:
		return "IOUSBHostCITransferCompletionMessageControlDeviceAddressPhase"
	case IOUSBHostCITransferCompletionMessageControlEndpointAddress:
		return "IOUSBHostCITransferCompletionMessageControlEndpointAddress"
	case IOUSBHostCITransferCompletionMessageControlEndpointAddressPhase:
		return "IOUSBHostCITransferCompletionMessageControlEndpointAddressPhase"
	case IOUSBHostCITransferCompletionMessageControlStatus:
		return "IOUSBHostCITransferCompletionMessageControlStatus"
	case IOUSBHostCITransferCompletionMessageControlStatusPhase:
		return "IOUSBHostCITransferCompletionMessageControlStatusPhase"
	default:
		return fmt.Sprintf("IOUSBHostCITransferCompletionMessageControl(%d)", e)
	}
}

type IOUSBHostCITransferCompletionMessageData0Transfer uint32

const (
	IOUSBHostCITransferCompletionMessageData0TransferLength      IOUSBHostCITransferCompletionMessageData0Transfer = 268435455
	IOUSBHostCITransferCompletionMessageData0TransferLengthPhase IOUSBHostCITransferCompletionMessageData0Transfer = 0
)

func (e IOUSBHostCITransferCompletionMessageData0Transfer) String() string {
	switch e {
	case IOUSBHostCITransferCompletionMessageData0TransferLength:
		return "IOUSBHostCITransferCompletionMessageData0TransferLength"
	case IOUSBHostCITransferCompletionMessageData0TransferLengthPhase:
		return "IOUSBHostCITransferCompletionMessageData0TransferLengthPhase"
	default:
		return fmt.Sprintf("IOUSBHostCITransferCompletionMessageData0Transfer(%d)", e)
	}
}

type IOUSBHostCITransferCompletionMessageData1Transfer uint

const (
	IOUSBHostCITransferCompletionMessageData1TransferStructure      IOUSBHostCITransferCompletionMessageData1Transfer = 18446744073709551615
	IOUSBHostCITransferCompletionMessageData1TransferStructurePhase IOUSBHostCITransferCompletionMessageData1Transfer = 0
)

func (e IOUSBHostCITransferCompletionMessageData1Transfer) String() string {
	switch e {
	case IOUSBHostCITransferCompletionMessageData1TransferStructure:
		return "IOUSBHostCITransferCompletionMessageData1TransferStructure"
	case IOUSBHostCITransferCompletionMessageData1TransferStructurePhase:
		return "IOUSBHostCITransferCompletionMessageData1TransferStructurePhase"
	default:
		return fmt.Sprintf("IOUSBHostCITransferCompletionMessageData1Transfer(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIUserClientVersion
type IOUSBHostCIUserClientVersion uint32

const (
	IOUSBHostCIUserClientVersion100 IOUSBHostCIUserClientVersion = 0
)

func (e IOUSBHostCIUserClientVersion) String() string {
	switch e {
	case IOUSBHostCIUserClientVersion100:
		return "IOUSBHostCIUserClientVersion100"
	default:
		return fmt.Sprintf("IOUSBHostCIUserClientVersion(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostIsochronousTransactionOptions
type IOUSBHostIsochronousTransactionOptions uint32

const (
	IOUSBHostIsochronousTransactionOptionsNone IOUSBHostIsochronousTransactionOptions = 0
	IOUSBHostIsochronousTransactionOptionsWrap IOUSBHostIsochronousTransactionOptions = 1
)

func (e IOUSBHostIsochronousTransactionOptions) String() string {
	switch e {
	case IOUSBHostIsochronousTransactionOptionsNone:
		return "IOUSBHostIsochronousTransactionOptionsNone"
	case IOUSBHostIsochronousTransactionOptionsWrap:
		return "IOUSBHostIsochronousTransactionOptionsWrap"
	default:
		return fmt.Sprintf("IOUSBHostIsochronousTransactionOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostIsochronousTransferOptions
type IOUSBHostIsochronousTransferOptions uint32

const (
	IOUSBHostIsochronousTransferOptionsNone IOUSBHostIsochronousTransferOptions = 0
)

func (e IOUSBHostIsochronousTransferOptions) String() string {
	switch e {
	case IOUSBHostIsochronousTransferOptionsNone:
		return "IOUSBHostIsochronousTransferOptionsNone"
	default:
		return fmt.Sprintf("IOUSBHostIsochronousTransferOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObjectDataOptions
type IOUSBHostObjectDataOptions uint

const (
	IOUSBHostObjectDataOptionsKernelUserShared IOUSBHostObjectDataOptions = 0
)

func (e IOUSBHostObjectDataOptions) String() string {
	switch e {
	case IOUSBHostObjectDataOptionsKernelUserShared:
		return "IOUSBHostObjectDataOptionsKernelUserShared"
	default:
		return fmt.Sprintf("IOUSBHostObjectDataOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObjectDestroyOptions
type IOUSBHostObjectDestroyOptions uint

const (
	IOUSBHostObjectDestroyOptionsDeviceSurrender IOUSBHostObjectDestroyOptions = 1
	IOUSBHostObjectDestroyOptionsNone            IOUSBHostObjectDestroyOptions = 0
)

func (e IOUSBHostObjectDestroyOptions) String() string {
	switch e {
	case IOUSBHostObjectDestroyOptionsDeviceSurrender:
		return "IOUSBHostObjectDestroyOptionsDeviceSurrender"
	case IOUSBHostObjectDestroyOptionsNone:
		return "IOUSBHostObjectDestroyOptionsNone"
	default:
		return fmt.Sprintf("IOUSBHostObjectDestroyOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObjectInitOptions
type IOUSBHostObjectInitOptions uint

const (
	// IOUSBHostObjectInitOptionsDeviceCapture: The option to capture the device and terminate existing drivers.
	IOUSBHostObjectInitOptionsDeviceCapture IOUSBHostObjectInitOptions = 1
	IOUSBHostObjectInitOptionsDeviceSeize   IOUSBHostObjectInitOptions = 2
	// IOUSBHostObjectInitOptionsNone: The default argument for initializing the host object.
	IOUSBHostObjectInitOptionsNone IOUSBHostObjectInitOptions = 0
)

func (e IOUSBHostObjectInitOptions) String() string {
	switch e {
	case IOUSBHostObjectInitOptionsDeviceCapture:
		return "IOUSBHostObjectInitOptionsDeviceCapture"
	case IOUSBHostObjectInitOptionsDeviceSeize:
		return "IOUSBHostObjectInitOptionsDeviceSeize"
	case IOUSBHostObjectInitOptionsNone:
		return "IOUSBHostObjectInitOptionsNone"
	default:
		return fmt.Sprintf("IOUSBHostObjectInitOptions(%d)", e)
	}
}
