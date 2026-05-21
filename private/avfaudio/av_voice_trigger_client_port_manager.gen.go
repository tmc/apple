// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"context"
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVVoiceTriggerClientPortManager] class.
var (
	_AVVoiceTriggerClientPortManagerClass     AVVoiceTriggerClientPortManagerClass
	_AVVoiceTriggerClientPortManagerClassOnce sync.Once
)

func getAVVoiceTriggerClientPortManagerClass() AVVoiceTriggerClientPortManagerClass {
	_AVVoiceTriggerClientPortManagerClassOnce.Do(func() {
		_AVVoiceTriggerClientPortManagerClass = AVVoiceTriggerClientPortManagerClass{class: objc.GetClass("AVVoiceTriggerClientPortManager")}
	})
	return _AVVoiceTriggerClientPortManagerClass
}

// GetAVVoiceTriggerClientPortManagerClass returns the class object for AVVoiceTriggerClientPortManager.
func GetAVVoiceTriggerClientPortManagerClass() AVVoiceTriggerClientPortManagerClass {
	return getAVVoiceTriggerClientPortManagerClass()
}

type AVVoiceTriggerClientPortManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVVoiceTriggerClientPortManagerClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVVoiceTriggerClientPortManagerClass) Alloc() AVVoiceTriggerClientPortManager {
	rv := objc.Send[AVVoiceTriggerClientPortManager](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVVoiceTriggerClientPortManager.CallMuteStateChangeNotificationBlock]
//   - [AVVoiceTriggerClientPortManager.CallRunningStateChangeNotificationBlock]
//   - [AVVoiceTriggerClientPortManager.DeviceID]
//   - [AVVoiceTriggerClientPortManager.SetDeviceID]
//   - [AVVoiceTriggerClientPortManager.Generation]
//   - [AVVoiceTriggerClientPortManager.SetGeneration]
//   - [AVVoiceTriggerClientPortManager.HysteresisDurationSeconds]
//   - [AVVoiceTriggerClientPortManager.SetHysteresisDurationSeconds]
//   - [AVVoiceTriggerClientPortManager.LastRunningStateSent]
//   - [AVVoiceTriggerClientPortManager.SetLastRunningStateSent]
//   - [AVVoiceTriggerClientPortManager.ListeningEnabled]
//   - [AVVoiceTriggerClientPortManager.SetListeningEnabled]
//   - [AVVoiceTriggerClientPortManager.MuteStateChangeNotificationRegistered]
//   - [AVVoiceTriggerClientPortManager.SetMuteStateChangeNotificationRegistered]
//   - [AVVoiceTriggerClientPortManager.NotifyMuteStateChanged]
//   - [AVVoiceTriggerClientPortManager.NotifyRunningStateChangedWithHysteresis]
//   - [AVVoiceTriggerClientPortManager.PortType]
//   - [AVVoiceTriggerClientPortManager.SetPortType]
//   - [AVVoiceTriggerClientPortManager.Queue]
//   - [AVVoiceTriggerClientPortManager.SetQueue]
//   - [AVVoiceTriggerClientPortManager.RunningStateChangeNotificationRegistered]
//   - [AVVoiceTriggerClientPortManager.SetRunningStateChangeNotificationRegistered]
//   - [AVVoiceTriggerClientPortManager.SetMuteStateChangeNotificationBlock]
//   - [AVVoiceTriggerClientPortManager.SetRunningStateChangeNotificationBlock]
//   - [AVVoiceTriggerClientPortManager.InitWithSerialQueuePortTypeHysteresisDurationSecondsRunningStateChangeNotificationBlockMuteStateChangeNotificationBlock]
type AVVoiceTriggerClientPortManager struct {
	objectivec.Object
}

// AVVoiceTriggerClientPortManagerFromID constructs a [AVVoiceTriggerClientPortManager] from an objc.ID.
func AVVoiceTriggerClientPortManagerFromID(id objc.ID) AVVoiceTriggerClientPortManager {
	return AVVoiceTriggerClientPortManager{objectivec.Object{ID: id}}
}

// Ensure AVVoiceTriggerClientPortManager implements IAVVoiceTriggerClientPortManager.
var _ IAVVoiceTriggerClientPortManager = AVVoiceTriggerClientPortManager{}

// An interface definition for the [AVVoiceTriggerClientPortManager] class.
//
// # Methods
//
//   - [IAVVoiceTriggerClientPortManager.CallMuteStateChangeNotificationBlock]
//   - [IAVVoiceTriggerClientPortManager.CallRunningStateChangeNotificationBlock]
//   - [IAVVoiceTriggerClientPortManager.DeviceID]
//   - [IAVVoiceTriggerClientPortManager.SetDeviceID]
//   - [IAVVoiceTriggerClientPortManager.Generation]
//   - [IAVVoiceTriggerClientPortManager.SetGeneration]
//   - [IAVVoiceTriggerClientPortManager.HysteresisDurationSeconds]
//   - [IAVVoiceTriggerClientPortManager.SetHysteresisDurationSeconds]
//   - [IAVVoiceTriggerClientPortManager.LastRunningStateSent]
//   - [IAVVoiceTriggerClientPortManager.SetLastRunningStateSent]
//   - [IAVVoiceTriggerClientPortManager.ListeningEnabled]
//   - [IAVVoiceTriggerClientPortManager.SetListeningEnabled]
//   - [IAVVoiceTriggerClientPortManager.MuteStateChangeNotificationRegistered]
//   - [IAVVoiceTriggerClientPortManager.SetMuteStateChangeNotificationRegistered]
//   - [IAVVoiceTriggerClientPortManager.NotifyMuteStateChanged]
//   - [IAVVoiceTriggerClientPortManager.NotifyRunningStateChangedWithHysteresis]
//   - [IAVVoiceTriggerClientPortManager.PortType]
//   - [IAVVoiceTriggerClientPortManager.SetPortType]
//   - [IAVVoiceTriggerClientPortManager.Queue]
//   - [IAVVoiceTriggerClientPortManager.SetQueue]
//   - [IAVVoiceTriggerClientPortManager.RunningStateChangeNotificationRegistered]
//   - [IAVVoiceTriggerClientPortManager.SetRunningStateChangeNotificationRegistered]
//   - [IAVVoiceTriggerClientPortManager.SetMuteStateChangeNotificationBlock]
//   - [IAVVoiceTriggerClientPortManager.SetRunningStateChangeNotificationBlock]
//   - [IAVVoiceTriggerClientPortManager.InitWithSerialQueuePortTypeHysteresisDurationSecondsRunningStateChangeNotificationBlockMuteStateChangeNotificationBlock]
type IAVVoiceTriggerClientPortManager interface {
	objectivec.IObject

	// Topic: Methods

	CallMuteStateChangeNotificationBlock(block bool)
	CallRunningStateChangeNotificationBlock(block bool)
	DeviceID() uint32
	SetDeviceID(value uint32)
	Generation() int64
	SetGeneration(value int64)
	HysteresisDurationSeconds() float32
	SetHysteresisDurationSeconds(value float32)
	LastRunningStateSent() bool
	SetLastRunningStateSent(value bool)
	ListeningEnabled() bool
	SetListeningEnabled(value bool)
	MuteStateChangeNotificationRegistered() bool
	SetMuteStateChangeNotificationRegistered(value bool)
	NotifyMuteStateChanged()
	NotifyRunningStateChangedWithHysteresis()
	PortType() uint64
	SetPortType(value uint64)
	Queue() objectivec.Object
	SetQueue(value objectivec.Object)
	RunningStateChangeNotificationRegistered() bool
	SetRunningStateChangeNotificationRegistered(value bool)
	SetMuteStateChangeNotificationBlock(block VoidHandler)
	SetRunningStateChangeNotificationBlock(block VoidHandler)
	InitWithSerialQueuePortTypeHysteresisDurationSecondsRunningStateChangeNotificationBlockMuteStateChangeNotificationBlock(queue objectivec.IObject, type_ uint64, seconds float32, block VoidHandler, block2 VoidHandler) AVVoiceTriggerClientPortManager
}

// Init initializes the instance.
func (a AVVoiceTriggerClientPortManager) Init() AVVoiceTriggerClientPortManager {
	rv := objc.Send[AVVoiceTriggerClientPortManager](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVVoiceTriggerClientPortManager) Autorelease() AVVoiceTriggerClientPortManager {
	rv := objc.Send[AVVoiceTriggerClientPortManager](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVoiceTriggerClientPortManager creates a new AVVoiceTriggerClientPortManager instance.
func NewAVVoiceTriggerClientPortManager() AVVoiceTriggerClientPortManager {
	class := getAVVoiceTriggerClientPortManagerClass()
	rv := objc.Send[AVVoiceTriggerClientPortManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (a AVVoiceTriggerClientPortManager) CallMuteStateChangeNotificationBlock(block bool) {
	objc.Send[objc.ID](a.ID, objc.Sel("callMuteStateChangeNotificationBlock:"), block)
}
func (a AVVoiceTriggerClientPortManager) CallRunningStateChangeNotificationBlock(block bool) {
	objc.Send[objc.ID](a.ID, objc.Sel("callRunningStateChangeNotificationBlock:"), block)
}
func (a AVVoiceTriggerClientPortManager) NotifyMuteStateChanged() {
	objc.Send[objc.ID](a.ID, objc.Sel("notifyMuteStateChanged"))
}
func (a AVVoiceTriggerClientPortManager) NotifyRunningStateChangedWithHysteresis() {
	objc.Send[objc.ID](a.ID, objc.Sel("notifyRunningStateChangedWithHysteresis"))
}
func (a AVVoiceTriggerClientPortManager) SetMuteStateChangeNotificationBlock(block VoidHandler) {
	_block0, _ := NewVoidBlock(block)
	objc.Send[objc.ID](a.ID, objc.Sel("setMuteStateChangeNotificationBlock:"), _block0)
}
func (a AVVoiceTriggerClientPortManager) SetRunningStateChangeNotificationBlock(block VoidHandler) {
	_block0, _ := NewVoidBlock(block)
	objc.Send[objc.ID](a.ID, objc.Sel("setRunningStateChangeNotificationBlock:"), _block0)
}
func (a AVVoiceTriggerClientPortManager) InitWithSerialQueuePortTypeHysteresisDurationSecondsRunningStateChangeNotificationBlockMuteStateChangeNotificationBlock(queue objectivec.IObject, type_ uint64, seconds float32, block VoidHandler, block2 VoidHandler) AVVoiceTriggerClientPortManager {
	_block3, _ := NewVoidBlock(block)
	_block4, _ := NewVoidBlock(block2)
	rv := objc.Send[AVVoiceTriggerClientPortManager](a.ID, objc.Sel("initWithSerialQueue:portType:hysteresisDurationSeconds:runningStateChangeNotificationBlock:muteStateChangeNotificationBlock:"), queue, type_, seconds, _block3, _block4)
	return rv
}

func (a AVVoiceTriggerClientPortManager) DeviceID() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("deviceID"))
	return rv
}
func (a AVVoiceTriggerClientPortManager) SetDeviceID(value uint32) {
	objc.Send[struct{}](a.ID, objc.Sel("setDeviceID:"), value)
}
func (a AVVoiceTriggerClientPortManager) Generation() int64 {
	rv := objc.Send[int64](a.ID, objc.Sel("generation"))
	return rv
}
func (a AVVoiceTriggerClientPortManager) SetGeneration(value int64) {
	objc.Send[struct{}](a.ID, objc.Sel("setGeneration:"), value)
}
func (a AVVoiceTriggerClientPortManager) HysteresisDurationSeconds() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("hysteresisDurationSeconds"))
	return rv
}
func (a AVVoiceTriggerClientPortManager) SetHysteresisDurationSeconds(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setHysteresisDurationSeconds:"), value)
}
func (a AVVoiceTriggerClientPortManager) LastRunningStateSent() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("lastRunningStateSent"))
	return rv
}
func (a AVVoiceTriggerClientPortManager) SetLastRunningStateSent(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setLastRunningStateSent:"), value)
}
func (a AVVoiceTriggerClientPortManager) ListeningEnabled() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("listeningEnabled"))
	return rv
}
func (a AVVoiceTriggerClientPortManager) SetListeningEnabled(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setListeningEnabled:"), value)
}
func (a AVVoiceTriggerClientPortManager) MuteStateChangeNotificationRegistered() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("muteStateChangeNotificationRegistered"))
	return rv
}
func (a AVVoiceTriggerClientPortManager) SetMuteStateChangeNotificationRegistered(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setMuteStateChangeNotificationRegistered:"), value)
}
func (a AVVoiceTriggerClientPortManager) PortType() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("portType"))
	return rv
}
func (a AVVoiceTriggerClientPortManager) SetPortType(value uint64) {
	objc.Send[struct{}](a.ID, objc.Sel("setPortType:"), value)
}
func (a AVVoiceTriggerClientPortManager) Queue() objectivec.Object {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("queue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (a AVVoiceTriggerClientPortManager) SetQueue(value objectivec.Object) {
	objc.Send[struct{}](a.ID, objc.Sel("setQueue:"), value)
}
func (a AVVoiceTriggerClientPortManager) RunningStateChangeNotificationRegistered() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("runningStateChangeNotificationRegistered"))
	return rv
}
func (a AVVoiceTriggerClientPortManager) SetRunningStateChangeNotificationRegistered(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setRunningStateChangeNotificationRegistered:"), value)
}

// SetMuteStateChangeNotificationBlockSync is a synchronous wrapper around [AVVoiceTriggerClientPortManager.SetMuteStateChangeNotificationBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (a AVVoiceTriggerClientPortManager) SetMuteStateChangeNotificationBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	a.SetMuteStateChangeNotificationBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetRunningStateChangeNotificationBlockSync is a synchronous wrapper around [AVVoiceTriggerClientPortManager.SetRunningStateChangeNotificationBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (a AVVoiceTriggerClientPortManager) SetRunningStateChangeNotificationBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	a.SetRunningStateChangeNotificationBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// InitWithSerialQueuePortTypeHysteresisDurationSecondsRunningStateChangeNotificationBlockMuteStateChangeNotificationBlockSync is a synchronous wrapper around [AVVoiceTriggerClientPortManager.InitWithSerialQueuePortTypeHysteresisDurationSecondsRunningStateChangeNotificationBlockMuteStateChangeNotificationBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (a AVVoiceTriggerClientPortManager) InitWithSerialQueuePortTypeHysteresisDurationSecondsRunningStateChangeNotificationBlockMuteStateChangeNotificationBlockSync(ctx context.Context, queue objectivec.IObject, type_ uint64, seconds float32, block VoidHandler) error {
	done := make(chan struct{}, 1)
	a.InitWithSerialQueuePortTypeHysteresisDurationSecondsRunningStateChangeNotificationBlockMuteStateChangeNotificationBlock(queue, type_, seconds, block, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
