// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"context"
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/coremidi"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/os"
)

// The class instance for the [AUAudioUnit] class.
var (
	_AUAudioUnitClass     AUAudioUnitClass
	_AUAudioUnitClassOnce sync.Once
)

func getAUAudioUnitClass() AUAudioUnitClass {
	_AUAudioUnitClassOnce.Do(func() {
		_AUAudioUnitClass = AUAudioUnitClass{class: objc.GetClass("AUAudioUnit")}
	})
	return _AUAudioUnitClass
}

// GetAUAudioUnitClass returns the class object for AUAudioUnit.
func GetAUAudioUnitClass() AUAudioUnitClass {
	return getAUAudioUnitClass()
}

type AUAudioUnitClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AUAudioUnitClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AUAudioUnitClass) Alloc() AUAudioUnit {
	rv := objc.Send[AUAudioUnit](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A class that defines a host’s interface to an audio unit.
//
// # Overview
//
// Hosts can instantiate either version 3 or version 2 audio units with this
// class, and to some extent control whether an audio unit is instantiated
// in-process or in a separate extension process.
//
// Version 3 audio units should subclass the [AUAudioUnit] class. Version 3
// audio unit components can be registered in the following ways:
//
// - Package the component into an app extension containing an
// [AudioComponents] `Info.Plist()` entry. The principal class must conform to
// the [AUAudioUnitFactory] protocol, which will typically instantiate an
// [AUAudioUnit] subclass. - Call the
// [AUAudioUnitClass.RegisterSubclassAsComponentDescriptionNameVersion] method
// to associate a component description with an [AUAudioUnit] subclass. Use
// the convention `:` when naming your audio unit component.
//
// Version 2 audio units should subclass the [AUAudioUnitV2Bridge] class
// instead. Version 2 audio unit components can be registered in the following
// ways:
//
// - Package the component into a component bundle containing an
// [AudioComponents] `Info.Plist()` entry, referring to an
// [AudioComponentFactoryFunction] function. - Call the
// [AudioComponentRegister] function to associate a component description with
// an [AudioComponentFactoryFunction] function.
//
// A host does not need to be aware of the concrete [AUAudioUnit] subclass
// that is being instantiated. The
// [AUAudioUnitV2Bridge.InitWithComponentDescriptionOptionsError] method
// ensures that the proper subclass is used.
//
// # Creating an Audio Unit
//
//   - [AUAudioUnit.InitWithComponentDescriptionError]: Synchronously initializes a new audio unit object.
//   - [AUAudioUnit.InitWithComponentDescriptionOptionsError]: Synchronously initializes a new audio unit object.
//
// # Returning the Audio Busses
//
//   - [AUAudioUnit.InputBusses]: An array containing the audio unit’s input connection points.
//   - [AUAudioUnit.OutputBusses]: An array containing the audio unit’s output connection points.
//
// # Customizing the Audio Unit Behavior
//
//   - [AUAudioUnit.ShouldChangeToFormatForBus]: This is called when you set the format on a bus.
//   - [AUAudioUnit.InternalRenderBlock]: The block which you must provide, via a getter, in order to implement rendering.
//   - [AUAudioUnit.MIDIOutputBufferSizeHint]
//   - [AUAudioUnit.SetMIDIOutputBufferSizeHint]
//
// # Querying Parameters
//
//   - [AUAudioUnit.ParameterTree]: An audio unit’s parameters, organized in a tree hierarchy.
//   - [AUAudioUnit.SetParameterTree]
//   - [AUAudioUnit.AllParameterValues]: Special read-only property for KVO.
//   - [AUAudioUnit.ParametersForOverviewWithCount]: Returns the audio unit’s most important parameters.
//
// # Providing Data to the Host
//
//   - [AUAudioUnit.MusicalContextBlock]: A callback to the host for musical context information.
//   - [AUAudioUnit.SetMusicalContextBlock]
//   - [AUAudioUnit.TransportStateBlock]: A callback to the host for transport state information.
//   - [AUAudioUnit.SetTransportStateBlock]
//   - [AUAudioUnit.ContextName]: Information about the host context in which the audio unit is connected, for display in the audio unit’s view.
//   - [AUAudioUnit.SetContextName]
//   - [AUAudioUnit.SupportsMPE]: A Boolean value that indicates whether the audio unit supports multi-dimensional polyphonic expression.
//
// # Managing MIDI Events
//
//   - [AUAudioUnit.IsMusicDeviceOrEffect]: Specifies whether an audio unit responds to MIDI events.
//   - [AUAudioUnit.VirtualMIDICableCount]: The number of virtual MIDI cables implemented by a music device or effect.
//   - [AUAudioUnit.ScheduleMIDIEventBlock]: A block used to schedule MIDI events.
//   - [AUAudioUnit.MIDIOutputEventBlock]
//   - [AUAudioUnit.SetMIDIOutputEventBlock]
//   - [AUAudioUnit.MIDIOutputNames]: The names of the MIDI outputs.
//
// # Managing Presets
//
//   - [AUAudioUnit.FullState]: A persistable snapshot of the audio unit’s properties and parameters, suitable for saving as a user preset.
//   - [AUAudioUnit.SetFullState]
//   - [AUAudioUnit.FullStateForDocument]: A persistable snapshot of the audio unit’s properties and parameters, suitable for saving in a user’s document.
//   - [AUAudioUnit.SetFullStateForDocument]
//   - [AUAudioUnit.FactoryPresets]: A collection of presets provided by the audio unit’s developer.
//   - [AUAudioUnit.CurrentPreset]: The audio unit’s last-selected preset.
//   - [AUAudioUnit.SetCurrentPreset]
//   - [AUAudioUnit.SupportsUserPresets]
//   - [AUAudioUnit.UserPresets]
//   - [AUAudioUnit.SaveUserPresetError]
//   - [AUAudioUnit.DeleteUserPresetError]
//   - [AUAudioUnit.PresetStateForError]
//
// # Managing the Render Cycle
//
//   - [AUAudioUnit.AllocateRenderResourcesAndReturnError]: Allocates resources required to render audio.
//   - [AUAudioUnit.DeallocateRenderResources]: Deallocates resources required to render audio.
//   - [AUAudioUnit.Reset]: Resets transitory rendering state to its initial state.
//   - [AUAudioUnit.RenderResourcesAllocated]: Determines whether the audio unit has allocated render resources.
//   - [AUAudioUnit.RenderBlock]: The block that hosts use to ask the audio unit to render audio.
//   - [AUAudioUnit.ScheduleParameterBlock]: The block that hosts use to schedule parameters.
//   - [AUAudioUnit.MaximumFramesToRender]: The maximum number of frames that the audio unit can render at once.
//   - [AUAudioUnit.SetMaximumFramesToRender]
//   - [AUAudioUnit.TokenByAddingRenderObserver]: Adds a block to be called on each render cycle.
//   - [AUAudioUnit.RemoveRenderObserver]: Removes an observer block previously added to the render cycle.
//
// # Messaging Channels
//
//   - [AUAudioUnit.MessageChannelFor]: Returns an object for bidirectional communication between an audio unit and its host.
//
// # Optimizing Performance
//
//   - [AUAudioUnit.Latency]: The audio unit’s processing latency, in seconds.
//   - [AUAudioUnit.TailTime]: The audio unit’s tail time, in seconds.
//   - [AUAudioUnit.RenderQuality]: Provides a trade-off between rendering quality and CPU load.
//   - [AUAudioUnit.SetRenderQuality]
//   - [AUAudioUnit.ShouldBypassEffect]: Determines whether an effect should route input directly to output, without any processing.
//   - [AUAudioUnit.SetShouldBypassEffect]
//   - [AUAudioUnit.CanProcessInPlace]: Determines whether an audio unit can process in place.
//   - [AUAudioUnit.IsRenderingOffline]: Communicates to an audio unit that it is rendering offline.
//   - [AUAudioUnit.SetRenderingOffline]
//
// # Describing the Audio Unit
//
//   - [AUAudioUnit.ComponentDescription]: The component description with which the audio unit was created.
//   - [AUAudioUnit.Component]: The component found in the component description with which the audio unit was created.
//   - [AUAudioUnit.ComponentName]: The audio unit’s component’s name.
//   - [AUAudioUnit.ComponentVersion]: The audio unit’s component’s version.
//   - [AUAudioUnit.AudioUnitName]: The audio unit’s name, derived from the component’s name.
//   - [AUAudioUnit.AudioUnitShortName]
//   - [AUAudioUnit.ManufacturerName]: The manufacturer’s name, derived from the component’s name.
//
// # Configuring the Channel Capabilities
//
//   - [AUAudioUnit.ChannelCapabilities]: Expresses valid combinations of input and output channels.
//   - [AUAudioUnit.ChannelMap]
//   - [AUAudioUnit.SetChannelMap]
//   - [AUAudioUnit.ProfileStateForCableChannel]
//   - [AUAudioUnit.EnableProfileCableOnChannelError]
//   - [AUAudioUnit.DisableProfileCableOnChannelError]
//   - [AUAudioUnit.ProfileChangedBlock]
//   - [AUAudioUnit.SetProfileChangedBlock]
//
// # Configuring the Device
//
//   - [AUAudioUnit.DeviceID]: Gets the I/O hardware device.
//   - [AUAudioUnit.SetDeviceIDError]: Sets the I/O hardware device.
//   - [AUAudioUnit.CanPerformInput]: Determines whether the I/O device can perform input.
//   - [AUAudioUnit.CanPerformOutput]: Determines whether the I/O device can perform output.
//   - [AUAudioUnit.IsInputEnabled]: A flag enabling audio input from the unit.
//   - [AUAudioUnit.SetInputEnabled]
//   - [AUAudioUnit.IsOutputEnabled]: A flag enabling audio output from the unit.
//   - [AUAudioUnit.SetOutputEnabled]
//   - [AUAudioUnit.InputHandler]: The block that the output unit will call to notify when input is available.
//   - [AUAudioUnit.SetInputHandler]
//   - [AUAudioUnit.OutputProvider]: The block that the output unit will call to get audio to send to the output.
//   - [AUAudioUnit.SetOutputProvider]
//   - [AUAudioUnit.DeviceInputLatency]: The audio device’s input latency, in seconds.
//   - [AUAudioUnit.DeviceOutputLatency]: The audio devic’s output latency, in seconds.
//   - [AUAudioUnit.StartHardwareAndReturnError]: Starts the audio hardware.
//   - [AUAudioUnit.StopHardware]: Stops the audio hardware.
//
// # Configuring the User Interface
//
//   - [AUAudioUnit.ProvidesUserInterface]: A Boolean that indicates whether the audio unit provides a user interface, normally in the form of a view controller.
//   - [AUAudioUnit.SupportedViewConfigurations]
//   - [AUAudioUnit.SelectViewConfiguration]
//
// # Getting the Runtime Behavior
//
//   - [AUAudioUnit.IsRunning]
//   - [AUAudioUnit.IsLoadedInProcess]
//
// # Instance properties
//
//   - [AUAudioUnit.AudioUnitMIDIProtocol]
//   - [AUAudioUnit.HostMIDIProtocol]
//   - [AUAudioUnit.SetHostMIDIProtocol]
//   - [AUAudioUnit.MIDIOutputEventListBlock]
//   - [AUAudioUnit.SetMIDIOutputEventListBlock]
//   - [AUAudioUnit.MigrateFromPlugin]
//   - [AUAudioUnit.ScheduleMIDIEventListBlock]
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit
type AUAudioUnit struct {
	objectivec.Object
}

// AUAudioUnitFromID constructs a [AUAudioUnit] from an objc.ID.
//
// A class that defines a host’s interface to an audio unit.
func AUAudioUnitFromID(id objc.ID) AUAudioUnit {
	return AUAudioUnit{objectivec.Object{ID: id}}
}

// NOTE: AUAudioUnit adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AUAudioUnit] class.
//
// # Creating an Audio Unit
//
//   - [IAUAudioUnit.InitWithComponentDescriptionError]: Synchronously initializes a new audio unit object.
//   - [IAUAudioUnit.InitWithComponentDescriptionOptionsError]: Synchronously initializes a new audio unit object.
//
// # Returning the Audio Busses
//
//   - [IAUAudioUnit.InputBusses]: An array containing the audio unit’s input connection points.
//   - [IAUAudioUnit.OutputBusses]: An array containing the audio unit’s output connection points.
//
// # Customizing the Audio Unit Behavior
//
//   - [IAUAudioUnit.ShouldChangeToFormatForBus]: This is called when you set the format on a bus.
//   - [IAUAudioUnit.InternalRenderBlock]: The block which you must provide, via a getter, in order to implement rendering.
//   - [IAUAudioUnit.MIDIOutputBufferSizeHint]
//   - [IAUAudioUnit.SetMIDIOutputBufferSizeHint]
//
// # Querying Parameters
//
//   - [IAUAudioUnit.ParameterTree]: An audio unit’s parameters, organized in a tree hierarchy.
//   - [IAUAudioUnit.SetParameterTree]
//   - [IAUAudioUnit.AllParameterValues]: Special read-only property for KVO.
//   - [IAUAudioUnit.ParametersForOverviewWithCount]: Returns the audio unit’s most important parameters.
//
// # Providing Data to the Host
//
//   - [IAUAudioUnit.MusicalContextBlock]: A callback to the host for musical context information.
//   - [IAUAudioUnit.SetMusicalContextBlock]
//   - [IAUAudioUnit.TransportStateBlock]: A callback to the host for transport state information.
//   - [IAUAudioUnit.SetTransportStateBlock]
//   - [IAUAudioUnit.ContextName]: Information about the host context in which the audio unit is connected, for display in the audio unit’s view.
//   - [IAUAudioUnit.SetContextName]
//   - [IAUAudioUnit.SupportsMPE]: A Boolean value that indicates whether the audio unit supports multi-dimensional polyphonic expression.
//
// # Managing MIDI Events
//
//   - [IAUAudioUnit.IsMusicDeviceOrEffect]: Specifies whether an audio unit responds to MIDI events.
//   - [IAUAudioUnit.VirtualMIDICableCount]: The number of virtual MIDI cables implemented by a music device or effect.
//   - [IAUAudioUnit.ScheduleMIDIEventBlock]: A block used to schedule MIDI events.
//   - [IAUAudioUnit.MIDIOutputEventBlock]
//   - [IAUAudioUnit.SetMIDIOutputEventBlock]
//   - [IAUAudioUnit.MIDIOutputNames]: The names of the MIDI outputs.
//
// # Managing Presets
//
//   - [IAUAudioUnit.FullState]: A persistable snapshot of the audio unit’s properties and parameters, suitable for saving as a user preset.
//   - [IAUAudioUnit.SetFullState]
//   - [IAUAudioUnit.FullStateForDocument]: A persistable snapshot of the audio unit’s properties and parameters, suitable for saving in a user’s document.
//   - [IAUAudioUnit.SetFullStateForDocument]
//   - [IAUAudioUnit.FactoryPresets]: A collection of presets provided by the audio unit’s developer.
//   - [IAUAudioUnit.CurrentPreset]: The audio unit’s last-selected preset.
//   - [IAUAudioUnit.SetCurrentPreset]
//   - [IAUAudioUnit.SupportsUserPresets]
//   - [IAUAudioUnit.UserPresets]
//   - [IAUAudioUnit.SaveUserPresetError]
//   - [IAUAudioUnit.DeleteUserPresetError]
//   - [IAUAudioUnit.PresetStateForError]
//
// # Managing the Render Cycle
//
//   - [IAUAudioUnit.AllocateRenderResourcesAndReturnError]: Allocates resources required to render audio.
//   - [IAUAudioUnit.DeallocateRenderResources]: Deallocates resources required to render audio.
//   - [IAUAudioUnit.Reset]: Resets transitory rendering state to its initial state.
//   - [IAUAudioUnit.RenderResourcesAllocated]: Determines whether the audio unit has allocated render resources.
//   - [IAUAudioUnit.RenderBlock]: The block that hosts use to ask the audio unit to render audio.
//   - [IAUAudioUnit.ScheduleParameterBlock]: The block that hosts use to schedule parameters.
//   - [IAUAudioUnit.MaximumFramesToRender]: The maximum number of frames that the audio unit can render at once.
//   - [IAUAudioUnit.SetMaximumFramesToRender]
//   - [IAUAudioUnit.TokenByAddingRenderObserver]: Adds a block to be called on each render cycle.
//   - [IAUAudioUnit.RemoveRenderObserver]: Removes an observer block previously added to the render cycle.
//
// # Messaging Channels
//
//   - [IAUAudioUnit.MessageChannelFor]: Returns an object for bidirectional communication between an audio unit and its host.
//
// # Optimizing Performance
//
//   - [IAUAudioUnit.Latency]: The audio unit’s processing latency, in seconds.
//   - [IAUAudioUnit.TailTime]: The audio unit’s tail time, in seconds.
//   - [IAUAudioUnit.RenderQuality]: Provides a trade-off between rendering quality and CPU load.
//   - [IAUAudioUnit.SetRenderQuality]
//   - [IAUAudioUnit.ShouldBypassEffect]: Determines whether an effect should route input directly to output, without any processing.
//   - [IAUAudioUnit.SetShouldBypassEffect]
//   - [IAUAudioUnit.CanProcessInPlace]: Determines whether an audio unit can process in place.
//   - [IAUAudioUnit.IsRenderingOffline]: Communicates to an audio unit that it is rendering offline.
//   - [IAUAudioUnit.SetRenderingOffline]
//
// # Describing the Audio Unit
//
//   - [IAUAudioUnit.ComponentDescription]: The component description with which the audio unit was created.
//   - [IAUAudioUnit.Component]: The component found in the component description with which the audio unit was created.
//   - [IAUAudioUnit.ComponentName]: The audio unit’s component’s name.
//   - [IAUAudioUnit.ComponentVersion]: The audio unit’s component’s version.
//   - [IAUAudioUnit.AudioUnitName]: The audio unit’s name, derived from the component’s name.
//   - [IAUAudioUnit.AudioUnitShortName]
//   - [IAUAudioUnit.ManufacturerName]: The manufacturer’s name, derived from the component’s name.
//
// # Configuring the Channel Capabilities
//
//   - [IAUAudioUnit.ChannelCapabilities]: Expresses valid combinations of input and output channels.
//   - [IAUAudioUnit.ChannelMap]
//   - [IAUAudioUnit.SetChannelMap]
//   - [IAUAudioUnit.ProfileStateForCableChannel]
//   - [IAUAudioUnit.EnableProfileCableOnChannelError]
//   - [IAUAudioUnit.DisableProfileCableOnChannelError]
//   - [IAUAudioUnit.ProfileChangedBlock]
//   - [IAUAudioUnit.SetProfileChangedBlock]
//
// # Configuring the Device
//
//   - [IAUAudioUnit.DeviceID]: Gets the I/O hardware device.
//   - [IAUAudioUnit.SetDeviceIDError]: Sets the I/O hardware device.
//   - [IAUAudioUnit.CanPerformInput]: Determines whether the I/O device can perform input.
//   - [IAUAudioUnit.CanPerformOutput]: Determines whether the I/O device can perform output.
//   - [IAUAudioUnit.IsInputEnabled]: A flag enabling audio input from the unit.
//   - [IAUAudioUnit.SetInputEnabled]
//   - [IAUAudioUnit.IsOutputEnabled]: A flag enabling audio output from the unit.
//   - [IAUAudioUnit.SetOutputEnabled]
//   - [IAUAudioUnit.InputHandler]: The block that the output unit will call to notify when input is available.
//   - [IAUAudioUnit.SetInputHandler]
//   - [IAUAudioUnit.OutputProvider]: The block that the output unit will call to get audio to send to the output.
//   - [IAUAudioUnit.SetOutputProvider]
//   - [IAUAudioUnit.DeviceInputLatency]: The audio device’s input latency, in seconds.
//   - [IAUAudioUnit.DeviceOutputLatency]: The audio devic’s output latency, in seconds.
//   - [IAUAudioUnit.StartHardwareAndReturnError]: Starts the audio hardware.
//   - [IAUAudioUnit.StopHardware]: Stops the audio hardware.
//
// # Configuring the User Interface
//
//   - [IAUAudioUnit.ProvidesUserInterface]: A Boolean that indicates whether the audio unit provides a user interface, normally in the form of a view controller.
//   - [IAUAudioUnit.SupportedViewConfigurations]
//   - [IAUAudioUnit.SelectViewConfiguration]
//
// # Getting the Runtime Behavior
//
//   - [IAUAudioUnit.IsRunning]
//   - [IAUAudioUnit.IsLoadedInProcess]
//
// # Instance properties
//
//   - [IAUAudioUnit.AudioUnitMIDIProtocol]
//   - [IAUAudioUnit.HostMIDIProtocol]
//   - [IAUAudioUnit.SetHostMIDIProtocol]
//   - [IAUAudioUnit.MIDIOutputEventListBlock]
//   - [IAUAudioUnit.SetMIDIOutputEventListBlock]
//   - [IAUAudioUnit.MigrateFromPlugin]
//   - [IAUAudioUnit.ScheduleMIDIEventListBlock]
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit
type IAUAudioUnit interface {
	objectivec.IObject

	// Topic: Creating an Audio Unit

	// Synchronously initializes a new audio unit object.
	InitWithComponentDescriptionError(componentDescription AudioComponentDescription) (AUAudioUnit, error)
	// Synchronously initializes a new audio unit object.
	InitWithComponentDescriptionOptionsError(componentDescription AudioComponentDescription, options AudioComponentInstantiationOptions) (AUAudioUnit, error)

	// Topic: Returning the Audio Busses

	// An array containing the audio unit’s input connection points.
	InputBusses() IAUAudioUnitBusArray
	// An array containing the audio unit’s output connection points.
	OutputBusses() IAUAudioUnitBusArray

	// Topic: Customizing the Audio Unit Behavior

	// This is called when you set the format on a bus.
	ShouldChangeToFormatForBus(format objectivec.IObject, bus IAUAudioUnitBus) bool
	// The block which you must provide, via a getter, in order to implement rendering.
	InternalRenderBlock() IntAudioUnitRenderActionFlagsHandler
	MIDIOutputBufferSizeHint() int
	SetMIDIOutputBufferSizeHint(value int)

	// Topic: Querying Parameters

	// An audio unit’s parameters, organized in a tree hierarchy.
	ParameterTree() IAUParameterTree
	SetParameterTree(value IAUParameterTree)
	// Special read-only property for KVO.
	AllParameterValues() bool
	// Returns the audio unit’s most important parameters.
	ParametersForOverviewWithCount(count int) []foundation.NSNumber

	// Topic: Providing Data to the Host

	// A callback to the host for musical context information.
	MusicalContextBlock() AUHostMusicalContextBlock
	SetMusicalContextBlock(value objc.ID)
	// A callback to the host for transport state information.
	TransportStateBlock() AUHostTransportStateBlock
	SetTransportStateBlock(value objc.ID)
	// Information about the host context in which the audio unit is connected, for display in the audio unit’s view.
	ContextName() string
	SetContextName(value string)
	// A Boolean value that indicates whether the audio unit supports multi-dimensional polyphonic expression.
	SupportsMPE() bool

	// Topic: Managing MIDI Events

	// Specifies whether an audio unit responds to MIDI events.
	IsMusicDeviceOrEffect() bool
	// The number of virtual MIDI cables implemented by a music device or effect.
	VirtualMIDICableCount() int
	// A block used to schedule MIDI events.
	ScheduleMIDIEventBlock() AUScheduleMIDIEventBlock
	MIDIOutputEventBlock() IntInt64Handler
	SetMIDIOutputEventBlock(value IntInt64Handler)
	// The names of the MIDI outputs.
	MIDIOutputNames() []string

	// Topic: Managing Presets

	// A persistable snapshot of the audio unit’s properties and parameters, suitable for saving as a user preset.
	FullState() foundation.INSDictionary
	SetFullState(value foundation.INSDictionary)
	// A persistable snapshot of the audio unit’s properties and parameters, suitable for saving in a user’s document.
	FullStateForDocument() foundation.INSDictionary
	SetFullStateForDocument(value foundation.INSDictionary)
	// A collection of presets provided by the audio unit’s developer.
	FactoryPresets() []AUAudioUnitPreset
	// The audio unit’s last-selected preset.
	CurrentPreset() IAUAudioUnitPreset
	SetCurrentPreset(value IAUAudioUnitPreset)
	SupportsUserPresets() bool
	UserPresets() []AUAudioUnitPreset
	SaveUserPresetError(userPreset IAUAudioUnitPreset) (bool, error)
	DeleteUserPresetError(userPreset IAUAudioUnitPreset) (bool, error)
	PresetStateForError(userPreset IAUAudioUnitPreset) (foundation.INSDictionary, error)

	// Topic: Managing the Render Cycle

	// Allocates resources required to render audio.
	AllocateRenderResourcesAndReturnError() (bool, error)
	// Deallocates resources required to render audio.
	DeallocateRenderResources()
	// Resets transitory rendering state to its initial state.
	Reset()
	// Determines whether the audio unit has allocated render resources.
	RenderResourcesAllocated() bool
	// The block that hosts use to ask the audio unit to render audio.
	RenderBlock() IntAudioUnitRenderActionFlagsHandler
	// The block that hosts use to schedule parameters.
	ScheduleParameterBlock() Int64Uint32Uint64Float32Handler
	// The maximum number of frames that the audio unit can render at once.
	MaximumFramesToRender() AUAudioFrameCount
	SetMaximumFramesToRender(value AUAudioFrameCount)
	// Adds a block to be called on each render cycle.
	TokenByAddingRenderObserver(observer IntAudioTimeStampUint32Int64Handler) int
	// Removes an observer block previously added to the render cycle.
	RemoveRenderObserver(token int)

	// Topic: Messaging Channels

	// Returns an object for bidirectional communication between an audio unit and its host.
	MessageChannelFor(channelName string) AUMessageChannel

	// Topic: Optimizing Performance

	// The audio unit’s processing latency, in seconds.
	Latency() foundation.NSTimeInterval
	// The audio unit’s tail time, in seconds.
	TailTime() foundation.NSTimeInterval
	// Provides a trade-off between rendering quality and CPU load.
	RenderQuality() int
	SetRenderQuality(value int)
	// Determines whether an effect should route input directly to output, without any processing.
	ShouldBypassEffect() bool
	SetShouldBypassEffect(value bool)
	// Determines whether an audio unit can process in place.
	CanProcessInPlace() bool
	// Communicates to an audio unit that it is rendering offline.
	IsRenderingOffline() bool
	SetRenderingOffline(value bool)

	// Topic: Describing the Audio Unit

	// The component description with which the audio unit was created.
	ComponentDescription() AudioComponentDescription
	// The component found in the component description with which the audio unit was created.
	Component() AudioComponent
	// The audio unit’s component’s name.
	ComponentName() string
	// The audio unit’s component’s version.
	ComponentVersion() uint32
	// The audio unit’s name, derived from the component’s name.
	AudioUnitName() string
	AudioUnitShortName() string
	// The manufacturer’s name, derived from the component’s name.
	ManufacturerName() string

	// Topic: Configuring the Channel Capabilities

	// Expresses valid combinations of input and output channels.
	ChannelCapabilities() []foundation.NSNumber
	ChannelMap() []foundation.NSNumber
	SetChannelMap(value []foundation.NSNumber)
	ProfileStateForCableChannel(cable uint8, channel MIDIChannelNumber) coremidi.MIDICIProfileState
	EnableProfileCableOnChannelError(profile *coremidi.MIDICIProfile, cable uint8, channel MIDIChannelNumber) (bool, error)
	DisableProfileCableOnChannelError(profile *coremidi.MIDICIProfile, cable uint8, channel MIDIChannelNumber) (bool, error)
	ProfileChangedBlock() Uint32Uint32MIDICIProfileBoolHandler
	SetProfileChangedBlock(value Uint32Uint32MIDICIProfileBoolHandler)

	// Topic: Configuring the Device

	// Gets the I/O hardware device.
	DeviceID() AUAudioObjectID
	// Sets the I/O hardware device.
	SetDeviceIDError(deviceID AUAudioObjectID) (bool, error)
	// Determines whether the I/O device can perform input.
	CanPerformInput() bool
	// Determines whether the I/O device can perform output.
	CanPerformOutput() bool
	// A flag enabling audio input from the unit.
	IsInputEnabled() bool
	SetInputEnabled(value bool)
	// A flag enabling audio output from the unit.
	IsOutputEnabled() bool
	SetOutputEnabled(value bool)
	// The block that the output unit will call to notify when input is available.
	InputHandler() AudioUnitRenderActionFlagsAudioTimeStampUint32Int64Handler
	SetInputHandler(value AudioUnitRenderActionFlagsAudioTimeStampUint32Int64Handler)
	// The block that the output unit will call to get audio to send to the output.
	OutputProvider() IntAudioUnitRenderActionFlagsHandler
	SetOutputProvider(value IntAudioUnitRenderActionFlagsHandler)
	// The audio device’s input latency, in seconds.
	DeviceInputLatency() foundation.NSTimeInterval
	// The audio devic’s output latency, in seconds.
	DeviceOutputLatency() foundation.NSTimeInterval
	// Starts the audio hardware.
	StartHardwareAndReturnError() (bool, error)
	// Stops the audio hardware.
	StopHardware()

	// Topic: Configuring the User Interface

	// A Boolean that indicates whether the audio unit provides a user interface, normally in the form of a view controller.
	ProvidesUserInterface() bool
	SupportedViewConfigurations(availableViewConfigurations []objectivec.IObject) foundation.NSIndexSet
	SelectViewConfiguration(viewConfiguration objectivec.IObject)

	// Topic: Getting the Runtime Behavior

	IsRunning() bool
	IsLoadedInProcess() bool

	// Topic: Instance properties

	AudioUnitMIDIProtocol() int32
	HostMIDIProtocol() int32
	SetHostMIDIProtocol(value int32)
	MIDIOutputEventListBlock() IntInt64Handler
	SetMIDIOutputEventListBlock(value IntInt64Handler)
	MigrateFromPlugin() foundation.INSArray
	ScheduleMIDIEventListBlock() IntInt64Handler

	// The workgroup associated with the audio device underlying this Audio Unit.
	OsWorkgroup() os.OSWorkgroup
	// The block that the system calls when the rendering context changes.
	RenderContextObserver() conststructAudioUnitRenderContextHandler
}

// Init initializes the instance.
func (a AUAudioUnit) Init() AUAudioUnit {
	rv := objc.Send[AUAudioUnit](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AUAudioUnit) Autorelease() AUAudioUnit {
	rv := objc.Send[AUAudioUnit](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAUAudioUnit creates a new AUAudioUnit instance.
func NewAUAudioUnit() AUAudioUnit {
	class := getAUAudioUnitClass()
	rv := objc.Send[AUAudioUnit](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Synchronously initializes a new audio unit object.
//
// componentDescription: The component to instantiate.
//
// # Return Value
//
// An initialized audio unit, or `nil` if initialization failed.
//
// # Discussion
//
// This is the convenience initializer.
//
// A single audio unit subclass may implement multiple audio units—for
// example, an effect that can also function as a generator, or a cluster of
// related effects.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/init(componentDescription:)
func NewAudioUnitWithComponentDescriptionError(componentDescription AudioComponentDescription) (AUAudioUnit, error) {
	var errorPtr objc.ID
	instance := getAUAudioUnitClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithComponentDescription:error:"), componentDescription, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AUAudioUnit{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return AUAudioUnit{}, objc.ErrInitFailed
	}
	return AUAudioUnitFromID(rv), nil
}

// Synchronously initializes a new audio unit object.
//
// componentDescription: The component to instantiate.
//
// options: Options for loading the unit in-process or out-of-process.
//
// # Return Value
//
// An initialized audio unit, or `nil` if initialization failed.
//
// # Discussion
//
// This is the designated initializer.
//
// A single audio unit subclass may implement multiple audio units—for
// example, an effect that can also function as a generator, or a cluster of
// related effects.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/init(componentDescription:options:)
func NewAudioUnitWithComponentDescriptionOptionsError(componentDescription AudioComponentDescription, options AudioComponentInstantiationOptions) (AUAudioUnit, error) {
	var errorPtr objc.ID
	instance := getAUAudioUnitClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithComponentDescription:options:error:"), componentDescription, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AUAudioUnit{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return AUAudioUnit{}, objc.ErrInitFailed
	}
	return AUAudioUnitFromID(rv), nil
}

// Synchronously initializes a new audio unit object.
//
// componentDescription: The component to instantiate.
//
// # Return Value
//
// An initialized audio unit, or `nil` if initialization failed.
//
// # Discussion
//
// This is the convenience initializer.
//
// A single audio unit subclass may implement multiple audio units—for
// example, an effect that can also function as a generator, or a cluster of
// related effects.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/init(componentDescription:)
func (a AUAudioUnit) InitWithComponentDescriptionError(componentDescription AudioComponentDescription) (AUAudioUnit, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithComponentDescription:error:"), componentDescription, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AUAudioUnit{}, foundation.NSErrorFrom(errorPtr)
	}
	return AUAudioUnitFromID(rv), nil

}

// Synchronously initializes a new audio unit object.
//
// componentDescription: The component to instantiate.
//
// options: Options for loading the unit in-process or out-of-process.
//
// # Return Value
//
// An initialized audio unit, or `nil` if initialization failed.
//
// # Discussion
//
// This is the designated initializer.
//
// A single audio unit subclass may implement multiple audio units—for
// example, an effect that can also function as a generator, or a cluster of
// related effects.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/init(componentDescription:options:)
func (a AUAudioUnit) InitWithComponentDescriptionOptionsError(componentDescription AudioComponentDescription, options AudioComponentInstantiationOptions) (AUAudioUnit, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithComponentDescription:options:error:"), componentDescription, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AUAudioUnit{}, foundation.NSErrorFrom(errorPtr)
	}
	return AUAudioUnitFromID(rv), nil

}

// This is called when you set the format on a bus.
//
// format: The proposed new format.
//
// bus: The bus on which the format will be changed.
//
// format is a [*avfaudio.AVAudioFormat].
//
// # Return Value
//
// - true if the new format will be set on the bus.
//
// # Discussion
//
// - false if the new format will not be set on the bus.
//
// # Discussion
//
// The bus has already checked that the format meets its channel constraints.
// The audio unit can override this method to check the format before allowing
// it to be set on the bus.
//
// The default implementation returns false if the audio unit’s
// [AUAudioUnit.RenderResourcesAllocated] value is true.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/shouldChange(to:for:)
func (a AUAudioUnit) ShouldChangeToFormatForBus(format objectivec.IObject, bus IAUAudioUnitBus) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("shouldChangeToFormat:forBus:"), format, bus)
	return rv
}

// Returns the audio unit’s most important parameters.
//
// count: The number of parameters to return.
//
// # Return Value
//
// An array of addresses representing the audio unit’s most important
// parameters.
//
// # Discussion
//
// This method allows a host to query an audio unit for a small number of its
// most important parameters, to be displayed in a compact generic view.
//
// This version 3 method is partially bridged to the version 2
// `kAudioUnitProperty_ParametersForOverview` API.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/parametersForOverview(withCount:)
func (a AUAudioUnit) ParametersForOverviewWithCount(count int) []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("parametersForOverviewWithCount:"), count)
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/saveUserPreset(_:)
func (a AUAudioUnit) SaveUserPresetError(userPreset IAUAudioUnitPreset) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("saveUserPreset:error:"), userPreset, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("saveUserPreset:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/deleteUserPreset(_:)
func (a AUAudioUnit) DeleteUserPresetError(userPreset IAUAudioUnitPreset) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("deleteUserPreset:error:"), userPreset, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("deleteUserPreset:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/presetState(for:)
func (a AUAudioUnit) PresetStateForError(userPreset IAUAudioUnitPreset) (foundation.INSDictionary, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("presetStateFor:error:"), userPreset, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSDictionaryFromID(rv), nil

}

// Allocates resources required to render audio.
//
// # Discussion
//
// - false if the operation failed.
//
// # Discussion
//
// Hosts must call this before beginning to render. Subclasses should call the
// superclass implementation.
//
// This version 3 method is bridged to the version 2 [AudioUnitInitialize]
// API.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/allocateRenderResources()
func (a AUAudioUnit) AllocateRenderResourcesAndReturnError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("allocateRenderResourcesAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("allocateRenderResourcesAndReturnError: returned NO with nil NSError")
	}
	return rv, nil

}

// Deallocates resources required to render audio.
//
// # Discussion
//
// Hosts should call this after finishing rendering. Subclasses should call
// the superclass implementation.
//
// This version 3 method is bridged to the version 2 [AudioUnitUninitialize]
// API.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/deallocateRenderResources()
func (a AUAudioUnit) DeallocateRenderResources() {
	objc.Send[objc.ID](a.ID, objc.Sel("deallocateRenderResources"))
}

// Resets transitory rendering state to its initial state.
//
// # Discussion
//
// Hosts should call this at the point of a discontinuity in the input stream
// being provided to an audio unit—for example, when seeking forward or
// backward within a track. In response, audio units should clear delay lines,
// filters, etc. Subclasses should call the superclass implementation.
//
// This version 3 method is bridged to the version 2 [AudioUnitReset] API, in
// the global scope.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/reset()
func (a AUAudioUnit) Reset() {
	objc.Send[objc.ID](a.ID, objc.Sel("reset"))
}

// Adds a block to be called on each render cycle.
//
// observer: The block to call.
//
// # Return Value
//
// A token to be used when removing the observer.
//
// # Discussion
//
// The supplied block is called at the beginning and ending of each render
// cycle. It should not make any blocking calls.
//
// This method is implemented in the [AUAudioUnit] base class and should not
// be overridden.
//
// This version 3 method is bridged to the version 2
// [AudioUnitAddRenderNotify] API.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/token(byAddingRenderObserver:)
func (a AUAudioUnit) TokenByAddingRenderObserver(observer IntAudioTimeStampUint32Int64Handler) int {
	_block0, _ := NewIntAudioTimeStampUint32Int64Block(observer)
	rv := objc.Send[int](a.ID, objc.Sel("tokenByAddingRenderObserver:"), _block0)
	return rv
}

// Removes an observer block previously added to the render cycle.
//
// token: The token associated with the block.
//
// # Discussion
//
// This version 3 property is bridged to the version 2
// [AudioUnitRemoveRenderNotify] API.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/removeRenderObserver(_:)
func (a AUAudioUnit) RemoveRenderObserver(token int) {
	objc.Send[objc.ID](a.ID, objc.Sel("removeRenderObserver:"), token)
}

// Returns an object for bidirectional communication between an audio unit and
// its host.
//
// channelName: The name of the message channel the audio unit returns.
//
// # Return Value
//
// An object that conforms to [AUMessageChannel].
//
// # Discussion
//
// Message channels provide a way for custom data exchanges between an audio
// unit and its host. An audio unit may support multiple message channels.
//
// The host manages the message channel object’s lifetime. Design message
// channel objects so they can outlive the audio unit that vended them.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/messageChannel(for:)
func (a AUAudioUnit) MessageChannelFor(channelName string) AUMessageChannel {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("messageChannelFor:"), objc.String(channelName))
	return AUMessageChannelObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/profileState(forCable:channel:)
func (a AUAudioUnit) ProfileStateForCableChannel(cable uint8, channel MIDIChannelNumber) coremidi.MIDICIProfileState {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("profileStateForCable:channel:"), cable, channel)
	return coremidi.MIDICIProfileStateFromID(rv)
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/enable(_:cable:onChannel:)
func (a AUAudioUnit) EnableProfileCableOnChannelError(profile *coremidi.MIDICIProfile, cable uint8, channel MIDIChannelNumber) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("enableProfile:cable:onChannel:error:"), profile.ID, cable, channel, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("enableProfile:cable:onChannel:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/disableProfile(_:cable:onChannel:)
func (a AUAudioUnit) DisableProfileCableOnChannelError(profile *coremidi.MIDICIProfile, cable uint8, channel MIDIChannelNumber) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("disableProfile:cable:onChannel:error:"), profile.ID, cable, channel, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("disableProfile:cable:onChannel:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Sets the I/O hardware device.
//
// deviceID: The device to set.
//
// # Discussion
//
// - false if the operation failed.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/setDeviceID(_:)
func (a AUAudioUnit) SetDeviceIDError(deviceID AUAudioObjectID) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("setDeviceID:error:"), deviceID, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setDeviceID:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Starts the audio hardware.
//
// # Discussion
//
// - false if the operation failed.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/startHardware()
func (a AUAudioUnit) StartHardwareAndReturnError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("startHardwareAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("startHardwareAndReturnError: returned NO with nil NSError")
	}
	return rv, nil

}

// Stops the audio hardware.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/stopHardware()
func (a AUAudioUnit) StopHardware() {
	objc.Send[objc.ID](a.ID, objc.Sel("stopHardware"))
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/supportedViewConfigurations(_:)
func (a AUAudioUnit) SupportedViewConfigurations(availableViewConfigurations []objectivec.IObject) foundation.NSIndexSet {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("supportedViewConfigurations:"), objectivec.IObjectSliceToNSArray(availableViewConfigurations))
	return foundation.NSIndexSetFromID(rv)
}

// viewConfiguration is a [*coreaudiokit.AUAudioUnitViewConfiguration].
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/select(_:)
func (a AUAudioUnit) SelectViewConfiguration(viewConfiguration objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("selectViewConfiguration:"), viewConfiguration)
}

// Asynchronously creates an audio unit instance.
//
// componentDescription: The component to instantiate.
//
// options: Options for loading the unit in-process or out-of-process.
//
// completionHandler: The block called when instantiation has completed. The block parameters are
// defined as follows:
//
// audioUnit: An initialized audio unit if the operation succeeded, or `nil`
// if it failed. error: An error if the operation failed, or `nil` if it
// succeeded.
//
// # Discussion
//
// Certain types of audio units must be instantiated asynchronously, such as
// version 3 units with a view.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/instantiate(with:options:completionHandler:)
func (_AUAudioUnitClass AUAudioUnitClass) InstantiateWithComponentDescriptionOptionsCompletionHandler(componentDescription AudioComponentDescription, options AudioComponentInstantiationOptions, completionHandler AudioComponentInstanceErrorHandler) {
	_block2, _ := NewAudioComponentInstanceErrorBlock(completionHandler)
	objc.Send[objc.ID](objc.ID(_AUAudioUnitClass.class), objc.Sel("instantiateWithComponentDescription:options:completionHandler:"), componentDescription, options, _block2)
}

// Registers an audio unit subclass.
//
// cls: An [AUAudioUnit] subclass.
//
// componentDescription: The component to register.
//
// name: The component’s name, using the convention `:`.
//
// version: The component’s version.
//
// # Discussion
//
// This method dynamically registers the supplied [AUAudioUnit] subclass with
// the Audio Component system, in the context of the current process only.
// After you’ve registered the subclass, instantiate it by calling one of
// the following:
//
// - The [AUAudioUnitV2Bridge.InitWithComponentDescriptionError] method. - The
// [AUAudioUnitV2Bridge.InitWithComponentDescriptionOptionsError] method. -
// The [AVAudioUnit] [instantiate(with:options:completionHandler:)] method. -
// The [AudioComponentInstanceNew] function.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/registerSubclass(_:as:name:version:)
//
// [AVAudioUnit]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnit
// [instantiate(with:options:completionHandler:)]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnit/instantiate(with:options:completionHandler:)
func (_AUAudioUnitClass AUAudioUnitClass) RegisterSubclassAsComponentDescriptionNameVersion(cls objectivec.Class, componentDescription AudioComponentDescription, name string, version uint32) {
	objc.Send[objc.ID](objc.ID(_AUAudioUnitClass.class), objc.Sel("registerSubclass:asComponentDescription:name:version:"), cls, componentDescription, objc.String(name), version)
}

// An array containing the audio unit’s input connection points.
//
// # Discussion
//
// Subclasses must override this property’s getter. The audio unit should
// return the same object every time it is asked for it, since hosts can
// install KVO observers on it.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/inputBusses
func (a AUAudioUnit) InputBusses() IAUAudioUnitBusArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("inputBusses"))
	return AUAudioUnitBusArrayFromID(objc.ID(rv))
}

// An array containing the audio unit’s output connection points.
//
// # Discussion
//
// Subclasses must override this property’s getter. The audio unit should
// return the same object every time it is asked for it, since hosts can
// install KVO observers on it.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/outputBusses
func (a AUAudioUnit) OutputBusses() IAUAudioUnitBusArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("outputBusses"))
	return AUAudioUnitBusArrayFromID(objc.ID(rv))
}

// The block which you must provide, via a getter, in order to implement
// rendering.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/internalRenderBlock
func (a AUAudioUnit) InternalRenderBlock() IntAudioUnitRenderActionFlagsHandler {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("internalRenderBlock"))
	_ = rv
	return nil
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/midiOutputBufferSizeHint
func (a AUAudioUnit) MIDIOutputBufferSizeHint() int {
	rv := objc.Send[int](a.ID, objc.Sel("MIDIOutputBufferSizeHint"))
	return rv
}
func (a AUAudioUnit) SetMIDIOutputBufferSizeHint(value int) {
	objc.Send[struct{}](a.ID, objc.Sel("setMIDIOutputBufferSizeHint:"), value)
}

// An audio unit’s parameters, organized in a tree hierarchy.
//
// # Discussion
//
// Hosts can fetch this property to discover a unit’s parameters. KVO
// notifications are issued on this member to notify the host of changes to
// the set of available parameters.
//
// Subclasses should implement this property to expose parameters to hosts.
// They should then cache as much data as possible and send KVO notifications
// on this property when altering the structure of the tree or the static
// information of parameters.
//
// This version 3 property is similar to the version 2
// `kAudioUnitProperty_ParameterList` and `kAudioUnitProperty_ParameterInfo`
// APIs.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/parameterTree
func (a AUAudioUnit) ParameterTree() IAUParameterTree {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("parameterTree"))
	return AUParameterTreeFromID(objc.ID(rv))
}
func (a AUAudioUnit) SetParameterTree(value IAUParameterTree) {
	objc.Send[struct{}](a.ID, objc.Sel("setParameterTree:"), value)
}

// Special read-only property for KVO.
//
// # Discussion
//
// KVO notifications are issued on this property in response to certain events
// where potentially all parameter values are invalidated—for example, the
// selection of a preset.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/allParameterValues
func (a AUAudioUnit) AllParameterValues() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("allParameterValues"))
	return rv
}

// A callback to the host for musical context information.
//
// # Discussion
//
// An audio unit accessing this property should cache it in realtime-safe
// storage before beginning to render.
//
// This version 3 property is bridged to the version 2
// [HostCallback_GetBeatAndTempo] and [HostCallback_GetMusicalTimeLocation]
// callback members in the `kAudioUnitProperty_HostCallbacks` API.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/musicalContextBlock
func (a AUAudioUnit) MusicalContextBlock() AUHostMusicalContextBlock {
	rv := objc.Send[AUHostMusicalContextBlock](a.ID, objc.Sel("musicalContextBlock"))
	return AUHostMusicalContextBlock(rv)
}
func (a AUAudioUnit) SetMusicalContextBlock(value objc.ID) {
	objc.Send[struct{}](a.ID, objc.Sel("setMusicalContextBlock:"), value)
}

// A callback to the host for transport state information.
//
// # Discussion
//
// An audio unit accessing this property should cache it in realtime-safe
// storage before beginning to render.
//
// This version 3 property is bridged to the version 2
// [HostCallback_GetTransportState] and `HostCallback_GetTransportState2`
// callback members in the `kAudioUnitProperty_HostCallbacks` API.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/transportStateBlock
func (a AUAudioUnit) TransportStateBlock() AUHostTransportStateBlock {
	rv := objc.Send[AUHostTransportStateBlock](a.ID, objc.Sel("transportStateBlock"))
	return AUHostTransportStateBlock(rv)
}
func (a AUAudioUnit) SetTransportStateBlock(value objc.ID) {
	objc.Send[struct{}](a.ID, objc.Sel("setTransportStateBlock:"), value)
}

// Information about the host context in which the audio unit is connected,
// for display in the audio unit’s view.
//
// # Discussion
//
// For example—a host could set “track 3” as the context, so that the
// audio unit’s view could then display “My audio unit on track 3”.
//
// This version 3 property is bridged to the version 2
// `kAudioUnitProperty_ContextName` API.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/contextName
func (a AUAudioUnit) ContextName() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("contextName"))
	return foundation.NSStringFromID(rv).String()
}
func (a AUAudioUnit) SetContextName(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setContextName:"), objc.String(value))
}

// A Boolean value that indicates whether the audio unit supports
// multi-dimensional polyphonic expression.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/supportsMPE
func (a AUAudioUnit) SupportsMPE() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("supportsMPE"))
	return rv
}

// Specifies whether an audio unit responds to MIDI events.
//
// # Discussion
//
// Returns true if the audio unit is a music device or effect.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/isMusicDeviceOrEffect
func (a AUAudioUnit) IsMusicDeviceOrEffect() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isMusicDeviceOrEffect"))
	return rv
}

// The number of virtual MIDI cables implemented by a music device or effect.
//
// # Discussion
//
// A music device or effect can support up to 256 virtual MIDI cables of
// input.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/virtualMIDICableCount
func (a AUAudioUnit) VirtualMIDICableCount() int {
	rv := objc.Send[int](a.ID, objc.Sel("virtualMIDICableCount"))
	return rv
}

// A block used to schedule MIDI events.
//
// # Discussion
//
// As with the render block, a host should fetch this block before beginning
// to render, if it intends to schedule MIDI events.
//
// This property is implemented in the [AUAudioUnit] base class. If the audio
// unit is not a music device or effect, this property is `nil`.
//
// Subclasses should not override this property. When hosts schedule events
// via this block, they are delivered to the audio unit via the list of render
// events delivered to the [AUAudioUnit.InternalRenderBlock] implementation.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/scheduleMIDIEventBlock
func (a AUAudioUnit) ScheduleMIDIEventBlock() AUScheduleMIDIEventBlock {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("scheduleMIDIEventBlock"))
	_ = rv
	return nil
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/midiOutputEventBlock
func (a AUAudioUnit) MIDIOutputEventBlock() IntInt64Handler {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("MIDIOutputEventBlock"))
	_ = rv
	return nil
}
func (a AUAudioUnit) SetMIDIOutputEventBlock(value IntInt64Handler) {
	block, cleanup := NewIntInt64Block(value)
	defer cleanup()
	objc.Send[struct{}](a.ID, objc.Sel("setMIDIOutputEventBlock:"), block)
}

// The names of the MIDI outputs.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/midiOutputNames
func (a AUAudioUnit) MIDIOutputNames() []string {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("MIDIOutputNames"))
	return objc.ConvertSliceToStrings(rv)
}

// A persistable snapshot of the audio unit’s properties and parameters,
// suitable for saving as a user preset.
//
// # Discussion
//
// Hosts may use this property to save and restore the state of an audio unit
// being used in a user preset or document. The audio unit should not persist
// transitory properties such as stream formats, but should save and restore
// all other properties.
//
// The base class implementation of this property saves the values of all
// parameters currently in the parameter tree. A subclass which dynamically
// produces multiple variants of the parameter tree needs to be aware that the
// serialization method does a depth-first preorder traversal of the tree.
//
// This version 3 property is bridged to the version 2
// `kAudioUnitProperty_ClassInfo` API.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/fullState
func (a AUAudioUnit) FullState() foundation.INSDictionary {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("fullState"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (a AUAudioUnit) SetFullState(value foundation.INSDictionary) {
	objc.Send[struct{}](a.ID, objc.Sel("setFullState:"), value)
}

// A persistable snapshot of the audio unit’s properties and parameters,
// suitable for saving in a user’s document.
//
// # Discussion
//
// Hosts may use this property to save and restore the state of an audio unit
// being used. Some state, such as a parameter value, is suitable for saving
// in a user preset. Other state, such as a synthesizer’s primary tuning
// setting, could be considered global state suitable for saving in a user
// document.
//
// Subclasses that do not implement this property interface with the
// [AUAudioUnit.FullState] property instead.
//
// This version 3 property is bridged to the version 2
// `kAudioUnitProperty_ClassInfoFromDocument` API.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/fullStateForDocument
func (a AUAudioUnit) FullStateForDocument() foundation.INSDictionary {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("fullStateForDocument"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (a AUAudioUnit) SetFullStateForDocument(value foundation.INSDictionary) {
	objc.Send[struct{}](a.ID, objc.Sel("setFullStateForDocument:"), value)
}

// A collection of presets provided by the audio unit’s developer.
//
// # Discussion
//
// A preset provides audio unit users with an easily-selectable, fine-tuned
// set of parameters defined by the developer.
//
// This version 3 property is bridged to the version 2
// `kAudioUnitProperty_FactoryPresets` API.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/factoryPresets
func (a AUAudioUnit) FactoryPresets() []AUAudioUnitPreset {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("factoryPresets"))
	return objc.ConvertSlice(rv, func(id objc.ID) AUAudioUnitPreset {
		return AUAudioUnitPresetFromID(id)
	})
}

// The audio unit’s last-selected preset.
//
// # Discussion
//
// Hosts can let the user select a preset by setting this property. When
// getting this property, the preset does not reflect whether parameters may
// have been modified since it was selected.
//
// This version 3 property is bridged to the version 2
// `kAudioUnitProperty_PresentPreset` API.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/currentPreset
func (a AUAudioUnit) CurrentPreset() IAUAudioUnitPreset {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("currentPreset"))
	return AUAudioUnitPresetFromID(objc.ID(rv))
}
func (a AUAudioUnit) SetCurrentPreset(value IAUAudioUnitPreset) {
	objc.Send[struct{}](a.ID, objc.Sel("setCurrentPreset:"), value)
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/supportsUserPresets
func (a AUAudioUnit) SupportsUserPresets() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("supportsUserPresets"))
	return rv
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/userPresets
func (a AUAudioUnit) UserPresets() []AUAudioUnitPreset {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("userPresets"))
	return objc.ConvertSlice(rv, func(id objc.ID) AUAudioUnitPreset {
		return AUAudioUnitPresetFromID(id)
	})
}

// Determines whether the audio unit has allocated render resources.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/renderResourcesAllocated
func (a AUAudioUnit) RenderResourcesAllocated() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("renderResourcesAllocated"))
	return rv
}

// The block that hosts use to ask the audio unit to render audio.
//
// # Discussion
//
// Before invoking an audio unit’s rendering functionality, a host should
// fetch this block and cache the result. The block can then be called from a
// realtime context without the possibility of blocking and causing an
// overload at the Core Audio HAL level.
//
// This block will call a subclass’s [AUAudioUnit.InternalRenderBlock]
// implementation, providing all realtime events scheduled for the current
// render time interval, bracketed by calls to any render observers.
// Subclasses should override their [AUAudioUnit.InternalRenderBlock]
// implementation, not this property.
//
// This version 3 property is bridged to the version 2 [AudioUnitRender] API.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/renderBlock
func (a AUAudioUnit) RenderBlock() IntAudioUnitRenderActionFlagsHandler {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("renderBlock"))
	_ = rv
	return nil
}

// The block that hosts use to schedule parameters.
//
// # Discussion
//
// As with the render block, a host should fetch this block before beginning
// to render, if it intends to schedule parameters.
//
// The block is safe to call from any thread context, including realtime audio
// render threads. Subclasses should not override this; it is implemented in
// the base class and will schedule the events to be provided to the
// [AUAudioUnit.InternalRenderBlock] implementation
//
// This version 3 property is bridged to the version 2
// [AudioUnitScheduleParameters] API.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/scheduleParameterBlock
func (a AUAudioUnit) ScheduleParameterBlock() Int64Uint32Uint64Float32Handler {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("scheduleParameterBlock"))
	_ = rv
	return nil
}

// The maximum number of frames that the audio unit can render at once.
//
// # Discussion
//
// This must be set by the host before render resources are allocated. It
// cannot be changed while render resources are allocated.
//
// This version 3 property is bridged to the version 2
// `kAudioUnitProperty_MaximumFramesPerSlice` API.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/maximumFramesToRender
func (a AUAudioUnit) MaximumFramesToRender() AUAudioFrameCount {
	rv := objc.Send[AUAudioFrameCount](a.ID, objc.Sel("maximumFramesToRender"))
	return AUAudioFrameCount(rv)
}
func (a AUAudioUnit) SetMaximumFramesToRender(value AUAudioFrameCount) {
	objc.Send[struct{}](a.ID, objc.Sel("setMaximumFramesToRender:"), value)
}

// The audio unit’s processing latency, in seconds.
//
// # Discussion
//
// This property reflects the delay between when an impulse arrives in the
// input stream vs. output stream. This should reflect the delay due to signal
// processing (e.g. FFTs), not as an effect (e.g. reverberation).
//
// Note that a latency that varies with parameter settings, including bypass,
// is generally not useful to hosts. A host is usually only prepared to add
// delays before starting to render and those delays need to be fixed. A
// variable delay would introduce artifacts even if the host could track it.
// If an algorithm has a variable latency, it should be adjusted upwards to
// some fixed latency within the audio unit. If for some reason this is not
// possible, then latency could be regarded as an unavoidable consequence of
// the algorithm and left unreported (i.e. a value of `0`).
//
// This version 3 property is bridged to the version 2
// `kAudioUnitProperty_Latency` API.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/latency
func (a AUAudioUnit) Latency() foundation.NSTimeInterval {
	rv := objc.Send[foundation.NSTimeInterval](a.ID, objc.Sel("latency"))
	return foundation.NSTimeInterval(rv)
}

// The audio unit’s tail time, in seconds.
//
// # Discussion
//
// This property reflects the time interval between when the input stream ends
// or otherwise transitions to silence, and when the output stream becomes
// silent. This should also reflect the duration of an effect (e.g.
// reverberation).
//
// This version 3 property is bridged to the version 2
// `kAudioUnitProperty_TailTime` API.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/tailTime
func (a AUAudioUnit) TailTime() foundation.NSTimeInterval {
	rv := objc.Send[foundation.NSTimeInterval](a.ID, objc.Sel("tailTime"))
	return foundation.NSTimeInterval(rv)
}

// Provides a trade-off between rendering quality and CPU load.
//
// # Discussion
//
// The range of valid values is 0 to 127.
//
// This version 3 property is bridged to the version 2
// `kAudioUnitProperty_RenderQuality` API.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/renderQuality
func (a AUAudioUnit) RenderQuality() int {
	rv := objc.Send[int](a.ID, objc.Sel("renderQuality"))
	return rv
}
func (a AUAudioUnit) SetRenderQuality(value int) {
	objc.Send[struct{}](a.ID, objc.Sel("setRenderQuality:"), value)
}

// Determines whether an effect should route input directly to output, without
// any processing.
//
// # Discussion
//
// This version 3 property is bridged to the version 2
// `kAudioUnitProperty_BypassEffect` API.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/shouldBypassEffect
func (a AUAudioUnit) ShouldBypassEffect() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("shouldBypassEffect"))
	return rv
}
func (a AUAudioUnit) SetShouldBypassEffect(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setShouldBypassEffect:"), value)
}

// Determines whether an audio unit can process in place.
//
// # Discussion
//
// In-place processing is the ability for an audio unit to transform an input
// signal to an output signal in-place in the input buffer, without requiring
// a separate output buffer.
//
// A host can express its desire to process in place by using null `mData`
// pointers in the output buffer list. If so, the audio unit may process
// in-place in the input buffers.
//
// This version 3 property is partially bridged to the version 2
// `kAudioUnitProperty_InPlaceProcessing` API. It is not settable in version
// 3.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/canProcessInPlace
func (a AUAudioUnit) CanProcessInPlace() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("canProcessInPlace"))
	return rv
}

// Communicates to an audio unit that it is rendering offline.
//
// # Discussion
//
// A host should use this property when using an audio unit in a context where
// there are no realtime deadlines. An audio unit may respond by using a more
// expensive signal processing algorithm, or allowing itself to block at
// render time if data being generated on secondary work threads is not ready
// in time.
//
// This version 3 property is bridged to the version 2
// `kAudioUnitProperty_OfflineRender` API.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/isRenderingOffline
func (a AUAudioUnit) IsRenderingOffline() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isRenderingOffline"))
	return rv
}
func (a AUAudioUnit) SetRenderingOffline(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setRenderingOffline:"), value)
}

// The component description with which the audio unit was created.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/componentDescription
func (a AUAudioUnit) ComponentDescription() AudioComponentDescription {
	rv := objc.Send[AudioComponentDescription](a.ID, objc.Sel("componentDescription"))
	return AudioComponentDescription(rv)
}

// The component found in the component description with which the audio unit
// was created.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/component
func (a AUAudioUnit) Component() AudioComponent {
	rv := objc.Send[AudioComponent](a.ID, objc.Sel("component"))
	return AudioComponent(rv)
}

// The audio unit’s component’s name.
//
// # Discussion
//
// By convention, an audio unit’s component name is
// “[AUAudioUnit.ManufacturerName]: [AUAudioUnit.AudioUnitName]”.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/componentName
func (a AUAudioUnit) ComponentName() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("componentName"))
	return foundation.NSStringFromID(rv).String()
}

// The audio unit’s component’s version.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/componentVersion
func (a AUAudioUnit) ComponentVersion() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("componentVersion"))
	return rv
}

// The audio unit’s name, derived from the component’s name.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/audioUnitName
func (a AUAudioUnit) AudioUnitName() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("audioUnitName"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/audioUnitShortName
func (a AUAudioUnit) AudioUnitShortName() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("audioUnitShortName"))
	return foundation.NSStringFromID(rv).String()
}

// The manufacturer’s name, derived from the component’s name.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/manufacturerName
func (a AUAudioUnit) ManufacturerName() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("manufacturerName"))
	return foundation.NSStringFromID(rv).String()
}

// Expresses valid combinations of input and output channels.
//
// # Discussion
//
// Array elements are [NSNumber] values containing integers.
//
// The array index alternates between input and output counts, in ascending
// order of input/output channels—for example: [0] = first input count, [1]
// = first output count, [2] = second input count, [3] = second output count,
// etc.
//
// Positive array values specify the number of input and/or output channels
// supported.
//
// Negative array values have particular meanings. An input/output value pair
// of `(-1, -1)` (i.e. [0] = -1, [1] = -1) indicates that any number of
// channels are supported, as long as they are the same number for both input
// and output. An input/output value pair combination of `-1` and `-2` (e.g.
// [0] = -1, [1] = -2) also indicates that any number of channels are
// supported, but without the requirement that the input and output counts are
// the same. A negative value less than `-2` (e.g. [0] = -16) specifies a
// total number of channels across every bus in that scope, regardless of how
// many channels are set on any particular bus.
//
// An array value of `0` (e.g. [0] = 0) specifies that the input/output
// channel is not applicable (though typically only used for input channels).
//
// The table below shows a sample selection of valid input and output channel
// combinations:
//
// [Table data omitted]
//
// This version 3 property is bridged to the version 2
// `kAudioUnitProperty_SupportedNumChannels` API.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/channelCapabilities
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (a AUAudioUnit) ChannelCapabilities() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("channelCapabilities"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/channelMap
func (a AUAudioUnit) ChannelMap() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("channelMap"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
func (a AUAudioUnit) SetChannelMap(value []foundation.NSNumber) {
	objc.Send[struct{}](a.ID, objc.Sel("setChannelMap:"), objectivec.IObjectSliceToNSArray(value))
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/profileChangedBlock
func (a AUAudioUnit) ProfileChangedBlock() Uint32Uint32MIDICIProfileBoolHandler {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("profileChangedBlock"))
	_ = rv
	return nil
}
func (a AUAudioUnit) SetProfileChangedBlock(value Uint32Uint32MIDICIProfileBoolHandler) {
	block, cleanup := NewUint32Uint32MIDICIProfileBoolBlock(value)
	defer cleanup()
	objc.Send[struct{}](a.ID, objc.Sel("setProfileChangedBlock:"), block)
}

// Gets the I/O hardware device.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/deviceID
func (a AUAudioUnit) DeviceID() AUAudioObjectID {
	rv := objc.Send[AUAudioObjectID](a.ID, objc.Sel("deviceID"))
	return AUAudioObjectID(rv)
}

// Determines whether the I/O device can perform input.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/canPerformInput
func (a AUAudioUnit) CanPerformInput() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("canPerformInput"))
	return rv
}

// Determines whether the I/O device can perform output.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/canPerformOutput
func (a AUAudioUnit) CanPerformOutput() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("canPerformOutput"))
	return rv
}

// A flag enabling audio input from the unit.
//
// # Discussion
//
// The default value is false.
//
// If your audio unit desires input audio, this property must be set to true
// and the value of [AUAudioUnit.CanPerformInput] must also be true.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/isInputEnabled
func (a AUAudioUnit) IsInputEnabled() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isInputEnabled"))
	return rv
}
func (a AUAudioUnit) SetInputEnabled(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setInputEnabled:"), value)
}

// A flag enabling audio output from the unit.
//
// # Discussion
//
// The default value is true.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/isOutputEnabled
func (a AUAudioUnit) IsOutputEnabled() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isOutputEnabled"))
	return rv
}
func (a AUAudioUnit) SetOutputEnabled(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setOutputEnabled:"), value)
}

// The block that the output unit will call to notify when input is available.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/inputHandler
func (a AUAudioUnit) InputHandler() AudioUnitRenderActionFlagsAudioTimeStampUint32Int64Handler {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("inputHandler"))
	_ = rv
	return nil
}
func (a AUAudioUnit) SetInputHandler(value AudioUnitRenderActionFlagsAudioTimeStampUint32Int64Handler) {
	block, cleanup := NewAudioUnitRenderActionFlagsAudioTimeStampUint32Int64Block(value)
	defer cleanup()
	objc.Send[struct{}](a.ID, objc.Sel("setInputHandler:"), block)
}

// The block that the output unit will call to get audio to send to the
// output.
//
// # Discussion
//
// This block must be set if output is enabled.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/outputProvider
func (a AUAudioUnit) OutputProvider() IntAudioUnitRenderActionFlagsHandler {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("outputProvider"))
	_ = rv
	return nil
}
func (a AUAudioUnit) SetOutputProvider(value IntAudioUnitRenderActionFlagsHandler) {
	block, cleanup := NewIntAudioUnitRenderActionFlagsBlock(value)
	defer cleanup()
	objc.Send[struct{}](a.ID, objc.Sel("setOutputProvider:"), block)
}

// The audio device’s input latency, in seconds.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/deviceInputLatency
func (a AUAudioUnit) DeviceInputLatency() foundation.NSTimeInterval {
	rv := objc.Send[foundation.NSTimeInterval](a.ID, objc.Sel("deviceInputLatency"))
	return foundation.NSTimeInterval(rv)
}

// The audio devic’s output latency, in seconds.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/deviceOutputLatency
func (a AUAudioUnit) DeviceOutputLatency() foundation.NSTimeInterval {
	rv := objc.Send[foundation.NSTimeInterval](a.ID, objc.Sel("deviceOutputLatency"))
	return foundation.NSTimeInterval(rv)
}

// A Boolean that indicates whether the audio unit provides a user interface,
// normally in the form of a view controller.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/providesUserInterface
func (a AUAudioUnit) ProvidesUserInterface() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("providesUserInterface"))
	return rv
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/isRunning
func (a AUAudioUnit) IsRunning() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isRunning"))
	return rv
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/isLoadedInProcess
func (a AUAudioUnit) IsLoadedInProcess() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isLoadedInProcess"))
	return rv
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/audioUnitMIDIProtocol
func (a AUAudioUnit) AudioUnitMIDIProtocol() int32 {
	rv := objc.Send[int32](a.ID, objc.Sel("AudioUnitMIDIProtocol"))
	return rv
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/hostMIDIProtocol
func (a AUAudioUnit) HostMIDIProtocol() int32 {
	rv := objc.Send[int32](a.ID, objc.Sel("hostMIDIProtocol"))
	return rv
}
func (a AUAudioUnit) SetHostMIDIProtocol(value int32) {
	objc.Send[struct{}](a.ID, objc.Sel("setHostMIDIProtocol:"), value)
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/midiOutputEventListBlock
func (a AUAudioUnit) MIDIOutputEventListBlock() IntInt64Handler {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("MIDIOutputEventListBlock"))
	_ = rv
	return nil
}
func (a AUAudioUnit) SetMIDIOutputEventListBlock(value IntInt64Handler) {
	block, cleanup := NewIntInt64Block(value)
	defer cleanup()
	objc.Send[struct{}](a.ID, objc.Sel("setMIDIOutputEventListBlock:"), block)
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/migrateFromPlugin
func (a AUAudioUnit) MigrateFromPlugin() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("migrateFromPlugin"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/scheduleMIDIEventListBlock
func (a AUAudioUnit) ScheduleMIDIEventListBlock() IntInt64Handler {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("scheduleMIDIEventListBlock"))
	_ = rv
	return nil
}

// The workgroup associated with the audio device underlying this Audio Unit.
//
// # Discussion
//
// Workgroups allow multiple threads to coordinate their activities for
// realtime operations. For Audio Units, this coordination occurs between the
// Audio Unit and other processes, such as the audio server and host app. The
// system uses the workgroup to observe the threads’ CPU usage and
// dynamically balance the competing considerations of power consumption and
// real-time rendering capacity.
//
// This version 3 property is bridged to the version 2
// [KAudioOutputUnitProperty_OSWorkgroup] property.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/osWorkgroup
func (a AUAudioUnit) OsWorkgroup() os.OSWorkgroup {
	rv := objc.Send[os.OSWorkgroup](a.ID, objc.Sel("osWorkgroup"))
	return os.OSWorkgroup(rv)
}

// The block that the system calls when the rendering context changes.
//
// # Discussion
//
// Implement this property if your Audio Unit creates auxilliary realtime
// rendering threads. Return a block for the system to call when the rendering
// context changes.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/renderContextObserver
func (a AUAudioUnit) RenderContextObserver() conststructAudioUnitRenderContextHandler {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("renderContextObserver"))
	_ = rv
	return nil
}

// InstantiateWithComponentDescriptionOptions is a synchronous wrapper around [AUAudioUnit.InstantiateWithComponentDescriptionOptionsCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (ac AUAudioUnitClass) InstantiateWithComponentDescriptionOptions(ctx context.Context, componentDescription AudioComponentDescription, options AudioComponentInstantiationOptions) (AudioComponentInstance, error) {
	type result struct {
		val AudioComponentInstance
		err error
	}
	done := make(chan result, 1)
	ac.InstantiateWithComponentDescriptionOptionsCompletionHandler(componentDescription, options, func(val AudioComponentInstance, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return *new(AudioComponentInstance), ctx.Err()
	}
}
