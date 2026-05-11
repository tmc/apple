// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import "github.com/tmc/apple/objc"

// NewAVAudioNodeTapBlock wraps a Go [AVAudioNodeTapBlock] as an Objective-C block.
// The caller must call the returned cleanup function after [AVAudioNode.RemoveTapOnBus].
func NewAVAudioNodeTapBlock(handler AVAudioNodeTapBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(_ objc.Block, bufferID objc.ID, timeID objc.ID) {
		handler(AVAudioPCMBufferFromID(bufferID), AVAudioTimeFromID(timeID))
	})
	return objc.ID(block), func() { block.Release() }
}

// InstallTapOnBus installs an audio tap and keeps the block alive until cleanup.
// Call the returned cleanup function after [AVAudioNode.RemoveTapOnBus].
func InstallTapOnBus(node IAVAudioNode, bus AVAudioNodeBus, bufferSize AVAudioFrameCount, format IAVAudioFormat, tap AVAudioNodeTapBlock) func() {
	blockID, cleanup := NewAVAudioNodeTapBlock(tap)
	objc.Send[objc.ID](node.GetID(), objc.Sel("installTapOnBus:bufferSize:format:block:"), bus, bufferSize, format, blockID)
	return cleanup
}
