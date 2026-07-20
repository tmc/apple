// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"unsafe"

	"github.com/tmc/apple/coreaudiotypes"
	"github.com/tmc/apple/objectivec"
)

// C struct types

// AVVoiceTriggerClientImpl
type AVVoiceTriggerClientImpl struct {
	Field1  unsafe.Pointer
	Field2  uint
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

// AudioComponentDescription
type AudioComponentDescription struct {
	Field1 uint
	Field2 uint
	Field3 uint
	Field4 uint
	Field5 uint
}

// AudioQueueBuffer
type AudioQueueBuffer struct {
	Field1 uint
	Field2 unsafe.Pointer
	Field3 uint
	Field4 unsafe.Pointer
	Field5 uint
	Field6 AudioStreamPacketDescriptionRef
	Field7 uint
}

// CAStreamBasicDescription
type CAStreamBasicDescription struct {
	Field1 float64
	Field2 uint
	Field3 uint
	Field4 uint
	Field5 uint
	Field6 uint
	Field7 uint
	Field8 uint
	Field9 uint
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
	Field1  uint
	Field2  unsafe.Pointer
	Field3  uint
	Field4  unsafe.Pointer
	Field5  uint
	Field6  AudioStreamPacketDescriptionRef
	Field7  uint
	Field8  coreaudiotypes.AudioStreamBasicDescription
	Field9  bool
	Field10 uint8
	Field11 uint8
	Field12 bool
}

// CFDictionary
type CFDictionary struct {
}

// IOSurface
type IOSurface struct {
}

// MachPort
type MachPort struct {
	Field1 uint
}

// Mach_port is a type alias for MachPort for use in objc.Send[T] calls.
type Mach_port = MachPort

// RecursiveMutex
type RecursiveMutex struct {
	__m_ unsafe.Pointer
}

// Recursive_mutex is a type alias for RecursiveMutex for use in objc.Send[T] calls.
type Recursive_mutex = RecursiveMutex

// XPCTypeS
type XPCTypeS struct {
}

// Xpc_type_s is a type alias for XPCTypeS for use in objc.Send[T] calls.
type Xpc_type_s = XPCTypeS
