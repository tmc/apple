// Code generated from Apple documentation. DO NOT EDIT.

package coreaudio

import (
	"github.com/tmc/apple/coreaudiotypes"
	"github.com/tmc/apple/objc"
)

// AudioDeviceIOBlock handles completion with primitive and object results.

// NewAudioDeviceIOBlock wraps a Go [AudioDeviceIOBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewAudioDeviceIOBlock(handler AudioDeviceIOBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive *coreaudiotypes.AudioTimeStamp, extra0 *coreaudiotypes.AudioBufferList, extra1 *coreaudiotypes.AudioTimeStamp, extra2 *coreaudiotypes.AudioBufferList, extra3 *coreaudiotypes.AudioTimeStamp) {
		handler(primitive, extra0, extra1, extra2, extra3)
	})
	return objc.ID(block), func() { block.Release() }
}

// AudioObjectPropertyListenerBlock handles completion with primitive and object results.

// NewAudioObjectPropertyListenerBlock wraps a Go [AudioObjectPropertyListenerBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewAudioObjectPropertyListenerBlock(handler AudioObjectPropertyListenerBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive uint32, extra0 *AudioObjectPropertyAddress) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}
