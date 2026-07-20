// Code generated from Apple documentation. DO NOT EDIT.

package coreaudio

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// AudioDeviceIOBlock handles completion with primitive and object results.

// NewAudioDeviceIOBlock wraps a Go [AudioDeviceIOBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewAudioDeviceIOBlock(handler AudioDeviceIOBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive objectivec.IObject, extra0 objectivec.IObject, extra1 objectivec.IObject, extra2 objectivec.IObject, extra3 objectivec.IObject) {
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
