// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"unsafe"

	"github.com/tmc/apple/coreaudiotypes"
	"github.com/tmc/apple/objectivec"
)

// C struct types

// AVVCRecordingEngine
type AVVCRecordingEngine struct {
}

// AVVoiceTriggerClientImpl
type AVVoiceTriggerClientImpl struct {
	Field1  unsafe.Pointer
	Field2  uint32
	Field3  bool
	Field4  bool
	Field5  bool
	Field6  bool
	Field7  bool
	Field8  objectivec.Object
	Field9  bool
	Field10 unsafe.Pointer
	Field11 unsafe.Pointer
	Field12 objectivec.Object
	Field13 bool
	Field14 objectivec.Object
	Field15 objectivec.Object
	Field16 bool
}

// Averager
type Averager struct {
}

// CAStreamBasicDescription
type CAStreamBasicDescription struct {
	Field1 float64
	Field2 uint32
	Field3 uint32
	Field4 uint32
	Field5 uint32
	Field6 uint32
	Field7 uint32
	Field8 uint32
	Field9 uint32
}

// ControllerImpl
type ControllerImpl struct {
}

// Impl
type Impl struct {
	Field1 unsafe.Pointer
	Field2 AveragerRef
}

// MusicTrackEventIteratorImpl
type MusicTrackEventIteratorImpl struct {
	Field1 OpaqueMusicEventIteratorRef
}

// MyAudioQueueBuffer
type MyAudioQueueBuffer struct {
	Field1  uint32
	Field2  unsafe.Pointer
	Field3  uint32
	Field4  unsafe.Pointer
	Field5  uint32
	Field6  AudioStreamPacketDescriptionRef
	Field7  uint32
	Field8  coreaudiotypes.AudioStreamBasicDescription
	Field9  bool
	Field10 uint8
	Field11 uint8
	Field12 bool
}

// OpaqueMusicEventIterator
type OpaqueMusicEventIterator struct {
}

// CFDictionary
type CFDictionary struct {
}

// IOSurface
type IOSurface struct {
}

// SharedWeakCount
type SharedWeakCount struct {
}

// MachPort
type MachPort struct {
	Field1 uint32
}

// Mach_port is a type alias for MachPort for use in objc.Send[T] calls.
type Mach_port = MachPort

// OpaquePthreadMutex
type OpaquePthreadMutex struct {
	__sig    int64
	__opaque [56]int8
}

// Opaque_pthread_mutex_t is a type alias for OpaquePthreadMutex for use in objc.Send[T] calls.
type Opaque_pthread_mutex_t = OpaquePthreadMutex

// RecursiveMutex
type RecursiveMutex struct {
	__m_ [8]uint64
}

// Recursive_mutex is a type alias for RecursiveMutex for use in objc.Send[T] calls.
type Recursive_mutex = RecursiveMutex

// XPCTypeS
type XPCTypeS struct {
}

// Xpc_type_s is a type alias for XPCTypeS for use in objc.Send[T] calls.
type Xpc_type_s = XPCTypeS
