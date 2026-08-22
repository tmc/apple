// Package main implements a virtual audio loopback driver equivalent to Existential Audio BlackHole.
package main

import (
	"errors"
	"fmt"
	"math"
	"sync"

	"github.com/tmc/apple/coreaudio"
	"github.com/tmc/apple/coreaudiotypes"
)

// Driver configuration constants matching Existential Audio BlackHole.
const (
	DriverName       = "BlackHole"
	ManufacturerName = "Existential Audio Inc."
	PlugInBundleID   = "audio.existential.BlackHole2ch"
	DeviceName       = "BlackHole 2ch"
	Device2Name      = "BlackHole 2ch 2"
	DeviceUID        = "BlackHole2ch_UID"
	Device2UID       = "BlackHole2ch_2_UID"
	BoxUID           = "BlackHole2ch_UID"
	ModelUID         = "BlackHole2ch_ModelUID"

	DefaultSampleRate    float64 = 48000.0
	DefaultChannelCount  uint32  = 2
	RingBufferFrameCount uint32  = 65536

	MinVolumeDB float32 = -64.0
	MaxVolumeDB float32 = 0.0
)

// SupportedSampleRates lists all sample rates matching BlackHole driver specs (8kHz to 768kHz).
var SupportedSampleRates = []float64{
	8000, 16000, 24000, 44100, 48000, 88200, 96000, 176400, 192000, 352800, 384000, 705600, 768000,
}

// Object IDs matching BlackHole CoreAudio plugin structure.
const (
	ObjectIDPlugin             coreaudio.AudioObjectID = 1
	ObjectIDBox                coreaudio.AudioObjectID = 2
	ObjectIDDevice             coreaudio.AudioObjectID = 3
	ObjectIDStreamInput        coreaudio.AudioObjectID = 4
	ObjectIDVolumeInputMaster  coreaudio.AudioObjectID = 5
	ObjectIDMuteInputMaster    coreaudio.AudioObjectID = 6
	ObjectIDStreamOutput       coreaudio.AudioObjectID = 7
	ObjectIDVolumeOutputMaster coreaudio.AudioObjectID = 8
	ObjectIDMuteOutputMaster   coreaudio.AudioObjectID = 9
	ObjectIDPitchAdjust        coreaudio.AudioObjectID = 10
	ObjectIDClockSource        coreaudio.AudioObjectID = 11
	ObjectIDDevice2            coreaudio.AudioObjectID = 12 // Mirror Device
)

// Clock source options.
const (
	ClockSourceInternalFixed      uint32 = 0
	ClockSourceInternalAdjustable uint32 = 1
)

// CoreAudio property selectors (4-character codes).
const (
	PropClass                         uint32 = 0x70636c73 // 'pcls'
	PropOwner                         uint32 = 0x73746476 // 'stdv'
	PropName                          uint32 = 0x6c6e616d // 'lnam'
	PropManufacturer                  uint32 = 0x6c6d616e // 'lman'
	PropModelUID                      uint32 = 0x6d756964 // 'muid'
	PropDeviceList                    uint32 = 0x64657623 // 'dev#'
	PropBoxList                       uint32 = 0x626f7823 // 'box#'
	PropResourceBundle                uint32 = 0x72737263 // 'rsrc'
	PropBoxUID                        uint32 = 0x62756964 // 'buid'
	PropIsAcquired                    uint32 = 0x62617164 // 'baqd'
	PropDeviceUID                     uint32 = 0x75696420 // 'uid '
	PropTransportType                 uint32 = 0x74726e73 // 'trns'
	PropIsHidden                      uint32 = 0x6869646e // 'hidn'
	PropNominalSampleRate             uint32 = 0x6e737274 // 'nsrt'
	PropAvailableNominalSampleRates   uint32 = 0x6e737261 // 'nsra'
	PropStreams                       uint32 = 0x73746d23 // 'stm#'
	PropControls                      uint32 = 0x63746c23 // 'ctl#'
	PropSafetyOffset                  uint32 = 0x73616674 // 'saft'
	PropLatency                       uint32 = 0x6c746e63 // 'ltnc'
	PropRingBufferFrameSize           uint32 = 0x72626673 // 'rbfs'
	PropIsRunning                     uint32 = 0x676f696e // 'goin'
	PropCanBeDefaultDevice            uint32 = 0x64666c74 // 'dflt'
	PropCanBeDefaultSystemDevice      uint32 = 0x7364666c // 'sdfl'
	PropDeviceIsAlive                 uint32 = 0x6c697665 // 'live'
	PropStreamDirection               uint32 = 0x73646972 // 'sdir'
	PropStreamTerminalType            uint32 = 0x7465726d // 'term'
	PropStreamStartingChannel         uint32 = 0x7363686e // 'schn'
	PropStreamVirtualFormat           uint32 = 0x76666d74 // 'vfmt'
	PropStreamPhysicalFormat          uint32 = 0x70666d74 // 'pfmt'
	PropLevelControlScalarValue       uint32 = 0x6c766c73 // 'lvls'
	PropLevelControlDecibelValue      uint32 = 0x6c766c64 // 'lvld'
	PropLevelControlDecibelRange      uint32 = 0x6c766c72 // 'lvlr'
	PropBooleanControlValue           uint32 = 0x6263766c // 'bcvl'
	PropSelectorControlValue          uint32 = 0x736c6374 // 'slct'
	PropSelectorControlAvailableItems uint32 = 0x7369746d // 'sitm'
	PropSelectorControlItemName       uint32 = 0x73696e6d // 'sinm'
)

// VolumeToDecibel converts a linear volume (0.0 to 1.0) to decibels (-64.0 dB to 0.0 dB).
func VolumeToDecibel(volume float32) float32 {
	minLinear := float32(math.Pow(10, float64(MinVolumeDB)/20.0))
	if volume <= minLinear {
		return MinVolumeDB
	}
	return 20.0 * float32(math.Log10(float64(volume)))
}

// VolumeFromDecibel converts decibels (-64.0 dB to 0.0 dB) to linear volume (0.0 to 1.0).
func VolumeFromDecibel(decibel float32) float32 {
	if decibel <= MinVolumeDB {
		return 0.0
	}
	return float32(math.Pow(10, float64(decibel)/20.0))
}

// VolumeToScalar converts linear volume to normalized scalar (0.0 to 1.0).
func VolumeToScalar(volume float32) float32 {
	db := VolumeToDecibel(volume)
	return (db - MinVolumeDB) / (MaxVolumeDB - MinVolumeDB)
}

// VolumeFromScalar converts normalized scalar (0.0 to 1.0) to linear volume.
func VolumeFromScalar(scalar float32) float32 {
	if scalar <= 0 {
		return 0.0
	}
	if scalar >= 1.0 {
		return 1.0
	}
	db := scalar*(MaxVolumeDB-MinVolumeDB) + MinVolumeDB
	return VolumeFromDecibel(db)
}

// ClientInfo represents an attached CoreAudio process client.
type ClientInfo struct {
	ClientID  uint32
	ProcessID int32
	IsActive  bool
}

// Driver represents the virtual audio loopback driver engine matching BlackHole.
type Driver struct {
	mu                   sync.RWMutex
	sampleRate           float64
	numChannels          uint32
	ringBufferSizeFrames uint32
	ringBuffer           []float32
	latencyFrames        uint32

	ioRunning            bool
	masterVolume         float32
	masterMute           bool
	pitchAdjust          float32
	pitchAdjustEnabled   bool
	clockSource          uint32
	deviceHidden         bool
	mirrorDeviceHidden   bool
	lastOutputSampleTime uint64
	bufferIsClear        bool

	clients map[uint32]*ClientInfo

	anchorSampleTime float64
	anchorHostTime   uint64
	hostClockFreq    float64
}

// NewDriver creates and initializes a new BlackHole virtual audio driver.
func NewDriver(sampleRate float64, numChannels uint32) (*Driver, error) {
	if sampleRate <= 0 {
		sampleRate = DefaultSampleRate
	}
	if numChannels == 0 {
		numChannels = DefaultChannelCount
	}

	validRate := false
	for _, r := range SupportedSampleRates {
		if r == sampleRate {
			validRate = true
			break
		}
	}
	if !validRate {
		return nil, fmt.Errorf("new driver: unsupported sample rate %.0f", sampleRate)
	}

	freq := coreaudio.AudioGetHostClockFrequency()
	if freq <= 0 {
		freq = 24000000.0 // Default 24MHz host clock if uninitialized
	}

	d := &Driver{
		sampleRate:           sampleRate,
		numChannels:          numChannels,
		ringBufferSizeFrames: RingBufferFrameCount,
		ringBuffer:           make([]float32, RingBufferFrameCount*numChannels),
		latencyFrames:        0,
		masterVolume:         1.0,
		masterMute:           false,
		pitchAdjust:          0.5,
		pitchAdjustEnabled:   false,
		clockSource:          ClockSourceInternalFixed,
		deviceHidden:         false,
		mirrorDeviceHidden:   true,
		bufferIsClear:        true,
		clients:              make(map[uint32]*ClientInfo),
		hostClockFreq:        freq,
	}
	return d, nil
}

// StartIO starts the audio IO processing cycle.
func (d *Driver) StartIO() error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.ioRunning {
		return errors.New("start IO: already running")
	}
	d.ioRunning = true
	d.anchorSampleTime = 0
	d.anchorHostTime = coreaudio.AudioGetCurrentHostTime()
	return nil
}

// StopIO stops the audio IO processing cycle.
func (d *Driver) StopIO() error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if !d.ioRunning {
		return errors.New("stop IO: not running")
	}
	d.ioRunning = false
	return nil
}

// IsIORunning returns whether the driver IO engine is active.
func (d *Driver) IsIORunning() bool {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.ioRunning
}

// SetSampleRate updates the driver nominal sample rate.
func (d *Driver) SetSampleRate(rate float64) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	valid := false
	for _, r := range SupportedSampleRates {
		if r == rate {
			valid = true
			break
		}
	}
	if !valid {
		return fmt.Errorf("set sample rate: invalid sample rate %.0f", rate)
	}

	d.sampleRate = rate
	d.anchorSampleTime = 0
	d.anchorHostTime = coreaudio.AudioGetCurrentHostTime()
	return nil
}

// SampleRate returns the current nominal sample rate.
func (d *Driver) SampleRate() float64 {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.sampleRate
}

// SetVolume sets the master output/input volume scalar (0.0 to 1.0).
func (d *Driver) SetVolume(scalar float32) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if scalar < 0.0 {
		scalar = 0.0
	}
	if scalar > 1.0 {
		scalar = 1.0
	}
	d.masterVolume = scalar
}

// Volume returns the current master volume scalar (0.0 to 1.0).
func (d *Driver) Volume() float32 {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.masterVolume
}

// SetMute enables or disables master mute.
func (d *Driver) SetMute(mute bool) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.masterMute = mute
}

// Mute returns whether master mute is enabled.
func (d *Driver) Mute() bool {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.masterMute
}

// SetPitchAdjust sets the pitch adjustment value (0.0 to 1.0, default 0.5).
func (d *Driver) SetPitchAdjust(pitch float32) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if pitch < 0.0 {
		pitch = 0.0
	}
	if pitch > 1.0 {
		pitch = 1.0
	}
	d.pitchAdjust = pitch
}

// PitchAdjust returns the current pitch adjustment value.
func (d *Driver) PitchAdjust() float32 {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.pitchAdjust
}

// SetPitchAdjustEnabled enables or disables pitch adjustment control.
func (d *Driver) SetPitchAdjustEnabled(enabled bool) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.pitchAdjustEnabled = enabled
}

// PitchAdjustEnabled returns whether pitch control is enabled.
func (d *Driver) PitchAdjustEnabled() bool {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.pitchAdjustEnabled
}

// SetClockSource sets the clock source (Internal Fixed or Internal Adjustable).
func (d *Driver) SetClockSource(source uint32) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if source != ClockSourceInternalFixed && source != ClockSourceInternalAdjustable {
		return fmt.Errorf("set clock source: invalid clock source %d", source)
	}
	d.clockSource = source
	return nil
}

// ClockSource returns the current clock source selector.
func (d *Driver) ClockSource() uint32 {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.clockSource
}

// AddClient registers a new process client ID with the driver.
func (d *Driver) AddClient(clientID uint32, pid int32) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.clients[clientID] = &ClientInfo{
		ClientID:  clientID,
		ProcessID: pid,
		IsActive:  true,
	}
}

// RemoveClient unregisters a client ID from the driver.
func (d *Driver) RemoveClient(clientID uint32) {
	d.mu.Lock()
	defer d.mu.Unlock()
	delete(d.clients, clientID)
}

// ClientCount returns the number of active attached process clients.
func (d *Driver) ClientCount() int {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return len(d.clients)
}

// WriteMix writes output PCM float32 samples from an audio application into the BlackHole ring buffer.
func (d *Driver) WriteMix(sampleTime uint64, numFrames uint32, ioBuffer []float32) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if !d.ioRunning {
		return errors.New("write mix: driver IO not running")
	}

	expectedSamples := numFrames * d.numChannels
	if uint32(len(ioBuffer)) < expectedSamples {
		return fmt.Errorf("write mix: buffer length %d smaller than required %d", len(ioBuffer), expectedSamples)
	}

	start := uint32(sampleTime % uint64(d.ringBufferSizeFrames))
	firstPartFrames := d.ringBufferSizeFrames - start
	if firstPartFrames > numFrames {
		firstPartFrames = numFrames
	}
	secondPartFrames := numFrames - firstPartFrames

	firstPartSamples := firstPartFrames * d.numChannels
	secondPartSamples := secondPartFrames * d.numChannels

	ringStartSample := start * d.numChannels
	copy(d.ringBuffer[ringStartSample:ringStartSample+firstPartSamples], ioBuffer[:firstPartSamples])
	if secondPartSamples > 0 {
		copy(d.ringBuffer[:secondPartSamples], ioBuffer[firstPartSamples:firstPartSamples+secondPartSamples])
	}

	d.lastOutputSampleTime = sampleTime + uint64(numFrames)
	d.bufferIsClear = false
	return nil
}

// ReadInput reads loopback PCM float32 samples from the BlackHole ring buffer into an input buffer for receiving applications.
func (d *Driver) ReadInput(sampleTime uint64, numFrames uint32, ioBuffer []float32) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if !d.ioRunning {
		return errors.New("read input: driver IO not running")
	}

	expectedSamples := numFrames * d.numChannels
	if uint32(len(ioBuffer)) < expectedSamples {
		return fmt.Errorf("read input: buffer length %d smaller than required %d", len(ioBuffer), expectedSamples)
	}

	// Fills zeroes if muted or if no audio output has been written recently.
	if d.masterMute || (d.lastOutputSampleTime > 0 && sampleTime >= d.lastOutputSampleTime+uint64(d.ringBufferSizeFrames)) {
		for i := uint32(0); i < expectedSamples; i++ {
			ioBuffer[i] = 0
		}
		if !d.bufferIsClear {
			for i := range d.ringBuffer {
				d.ringBuffer[i] = 0
			}
			d.bufferIsClear = true
		}
		return nil
	}

	start := uint32(sampleTime % uint64(d.ringBufferSizeFrames))
	firstPartFrames := d.ringBufferSizeFrames - start
	if firstPartFrames > numFrames {
		firstPartFrames = numFrames
	}
	secondPartFrames := numFrames - firstPartFrames

	firstPartSamples := firstPartFrames * d.numChannels
	secondPartSamples := secondPartFrames * d.numChannels

	ringStartSample := start * d.numChannels
	copy(ioBuffer[:firstPartSamples], d.ringBuffer[ringStartSample:ringStartSample+firstPartSamples])
	if secondPartSamples > 0 {
		copy(ioBuffer[firstPartSamples:firstPartSamples+secondPartSamples], d.ringBuffer[:secondPartSamples])
	}

	// Apply master volume scaling
	vol := d.masterVolume
	if vol != 1.0 {
		for i := uint32(0); i < expectedSamples; i++ {
			ioBuffer[i] *= vol
		}
	}
	return nil
}

// GetZeroTimeStamp calculates the current sample time and host time mapping for CoreAudio synchronization.
func (d *Driver) GetZeroTimeStamp() (sampleTime float64, hostTime uint64) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	currentHost := coreaudio.AudioGetCurrentHostTime()
	if d.anchorHostTime == 0 {
		return 0, currentHost
	}

	ticksPerFrame := d.hostClockFreq / d.sampleRate
	elapsedHost := float64(currentHost - d.anchorHostTime)
	sampleTime = d.anchorSampleTime + elapsedHost/ticksPerFrame
	return sampleTime, currentHost
}

// AudioStreamBasicDescription returns the stream format for BlackHole (32-bit float LPCM).
func (d *Driver) AudioStreamBasicDescription() coreaudiotypes.AudioStreamBasicDescription {
	d.mu.RLock()
	defer d.mu.RUnlock()

	bytesPerSample := uint32(4)
	return coreaudiotypes.AudioStreamBasicDescription{
		MSampleRate:       d.sampleRate,
		MFormatID:         0x6c70636d, // 'lpcm'
		MFormatFlags:      0x9,        // kAudioFormatFlagIsFloat | kAudioFormatFlagIsPacked
		MBytesPerPacket:   bytesPerSample * d.numChannels,
		MFramesPerPacket:  1,
		MBytesPerFrame:    bytesPerSample * d.numChannels,
		MChannelsPerFrame: d.numChannels,
		MBitsPerChannel:   32,
		MReserved:         0,
	}
}

// AudioObjectPropertyAddress represents a property address in CoreAudio queries.
type AudioObjectPropertyAddress struct {
	Selector uint32
	Scope    uint32
	Element  uint32
}

// HasProperty returns whether the driver supports a given property address for an object ID.
func (d *Driver) HasProperty(objectID coreaudio.AudioObjectID, addr AudioObjectPropertyAddress) bool {
	switch objectID {
	case ObjectIDPlugin:
		switch addr.Selector {
		case PropClass, PropOwner, PropManufacturer, PropDeviceList, PropBoxList, PropResourceBundle:
			return true
		}
	case ObjectIDBox:
		switch addr.Selector {
		case PropClass, PropOwner, PropName, PropModelUID, PropBoxUID, PropIsAcquired:
			return true
		}
	case ObjectIDDevice, ObjectIDDevice2:
		switch addr.Selector {
		case PropClass, PropOwner, PropName, PropManufacturer, PropDeviceUID, PropModelUID,
			PropTransportType, PropIsHidden, PropNominalSampleRate, PropAvailableNominalSampleRates,
			PropStreams, PropControls, PropSafetyOffset, PropLatency, PropRingBufferFrameSize,
			PropIsRunning, PropCanBeDefaultDevice, PropCanBeDefaultSystemDevice, PropDeviceIsAlive:
			return true
		}
	case ObjectIDStreamInput, ObjectIDStreamOutput:
		switch addr.Selector {
		case PropClass, PropOwner, PropStreamDirection, PropStreamTerminalType, PropStreamStartingChannel, PropStreamVirtualFormat, PropStreamPhysicalFormat:
			return true
		}
	case ObjectIDVolumeInputMaster, ObjectIDVolumeOutputMaster:
		switch addr.Selector {
		case PropClass, PropOwner, PropLevelControlScalarValue, PropLevelControlDecibelValue, PropLevelControlDecibelRange:
			return true
		}
	case ObjectIDMuteInputMaster, ObjectIDMuteOutputMaster:
		switch addr.Selector {
		case PropClass, PropOwner, PropBooleanControlValue:
			return true
		}
	case ObjectIDPitchAdjust:
		switch addr.Selector {
		case PropClass, PropOwner, PropLevelControlScalarValue:
			return true
		}
	case ObjectIDClockSource:
		switch addr.Selector {
		case PropClass, PropOwner, PropSelectorControlValue, PropSelectorControlAvailableItems, PropSelectorControlItemName:
			return true
		}
	}
	return false
}

// GetPropertyData retrieves property data matching CoreAudio HAL query interface.
func (d *Driver) GetPropertyData(objectID coreaudio.AudioObjectID, addr AudioObjectPropertyAddress) (interface{}, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	if !d.HasProperty(objectID, addr) {
		return nil, fmt.Errorf("get property: unknown property 0x%x for object %d", addr.Selector, objectID)
	}

	switch addr.Selector {
	case PropName:
		switch objectID {
		case ObjectIDBox:
			return DriverName, nil
		case ObjectIDDevice:
			return fmt.Sprintf("%s %dch", DriverName, d.numChannels), nil
		case ObjectIDDevice2:
			return fmt.Sprintf("%s %dch 2", DriverName, d.numChannels), nil
		}
	case PropManufacturer:
		return ManufacturerName, nil
	case PropDeviceUID:
		if objectID == ObjectIDDevice2 {
			return Device2UID, nil
		}
		return DeviceUID, nil
	case PropNominalSampleRate:
		return d.sampleRate, nil
	case PropAvailableNominalSampleRates:
		return SupportedSampleRates, nil
	case PropIsRunning:
		if d.ioRunning {
			return uint32(1), nil
		}
		return uint32(0), nil
	case PropIsHidden:
		if objectID == ObjectIDDevice2 {
			return d.mirrorDeviceHidden, nil
		}
		return d.deviceHidden, nil
	case PropLevelControlScalarValue:
		if objectID == ObjectIDPitchAdjust {
			return d.pitchAdjust, nil
		}
		return d.masterVolume, nil
	case PropLevelControlDecibelValue:
		return VolumeToDecibel(d.masterVolume), nil
	case PropBooleanControlValue:
		return d.masterMute, nil
	case PropSelectorControlValue:
		return d.clockSource, nil
	case PropLatency:
		return d.latencyFrames, nil
	case PropRingBufferFrameSize:
		return d.ringBufferSizeFrames, nil
	}
	return nil, nil
}

// SetPropertyData mutates property data matching CoreAudio HAL property setter interface.
func (d *Driver) SetPropertyData(objectID coreaudio.AudioObjectID, addr AudioObjectPropertyAddress, val interface{}) error {
	switch addr.Selector {
	case PropNominalSampleRate:
		if rate, ok := val.(float64); ok {
			return d.SetSampleRate(rate)
		}
		return errors.New("set nominal sample rate: value must be float64")
	case PropLevelControlScalarValue:
		if scalar, ok := val.(float32); ok {
			if objectID == ObjectIDPitchAdjust {
				d.SetPitchAdjust(scalar)
				return nil
			}
			d.SetVolume(scalar)
			return nil
		}
		return errors.New("set volume scalar: value must be float32")
	case PropBooleanControlValue:
		if mute, ok := val.(bool); ok {
			d.SetMute(mute)
			return nil
		}
		return errors.New("set mute: value must be bool")
	case PropSelectorControlValue:
		if src, ok := val.(uint32); ok {
			return d.SetClockSource(src)
		}
		return errors.New("set clock source: value must be uint32")
	}
	return fmt.Errorf("set property: unsupported property 0x%x for object %d", addr.Selector, objectID)
}

// String returns a descriptive overview of the driver state.
func (d *Driver) String() string {
	d.mu.RLock()
	defer d.mu.RUnlock()

	status := "stopped"
	if d.ioRunning {
		status = "running"
	}
	return fmt.Sprintf("%s (%d channels, %.0f Hz, status: %s, volume: %.2f, muted: %v)",
		DeviceName, d.numChannels, d.sampleRate, status, d.masterVolume, d.masterMute)
}
