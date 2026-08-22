// Code generated from Apple documentation. DO NOT EDIT.

package soundanalysis

import (
	"github.com/tmc/apple/objc"
)

// BoolHandler handles A completion closure (Swift) or block (Objective-C) the analyzer calls when it finishes analyzing a file.
//
// Used by:
//   - [SNAudioFileAnalyzer.AnalyzeWithCompletionHandler]
type BoolHandler = func(bool)

// NewBoolBlock wraps a Go [BoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [SNAudioFileAnalyzer.AnalyzeWithCompletionHandler]
func NewBoolBlock(handler BoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal bool) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}
