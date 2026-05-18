// Code generated from Apple documentation for ApplicationServices. DO NOT EDIT.

package applicationservices

import (
	"unsafe"
)

// Point is the QuickDraw point type used by legacy Carbon APIs.
type Point struct {
	V int16
	H int16
}

// Rect is the QuickDraw rectangle type used by legacy Carbon APIs.
type Rect struct {
	Top    int16
	Left   int16
	Bottom int16
	Right  int16
}

// ProcessSerialNumber identifies a process in legacy Process Manager APIs.
type ProcessSerialNumber struct {
	HighLongOfPSN uint32
	LowLongOfPSN  uint32
}

type ICInstance = unsafe.Pointer
type ICProfileID = uint32
type ICAttr = uint32
type ICPerm = uint32
type OptionBits = uint32
type PasteboardRef = uintptr

type SpeechSyncUPP = unsafe.Pointer
type SpeechWordUPP = unsafe.Pointer
type SpeechWordProcPtr = func(*SpeechChannelRecord, uintptr, uint, uint16)
type SpeechPhonemeUPP = unsafe.Pointer
type SpeechErrorUPP = unsafe.Pointer
type SpeechDoneUPP = unsafe.Pointer
type SpeechTextDoneUPP = unsafe.Pointer
type SpeechDoneProcPtr = func(*SpeechChannelRecord, uintptr)
type SpeechPhonemeProcPtr = func(*SpeechChannelRecord, uintptr, int16)
type SpeechErrorProcPtr = func(*SpeechChannelRecord, uintptr, int16, int)
type SpeechTextDoneProcPtr = func(*SpeechChannelRecord, uintptr, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
type SpeechSyncProcPtr = func(*SpeechChannelRecord, uintptr, uint32)

type LaunchPBPtr = unsafe.Pointer

type QDStdTextUPP = unsafe.Pointer
type QDJShieldCursorUPP = unsafe.Pointer
type QDTextWidthUPP = unsafe.Pointer
