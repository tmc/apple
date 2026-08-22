// Code generated from Apple documentation. DO NOT EDIT.

package accessibility

import (
	"github.com/tmc/apple/objc"
)

// AXBrailleMapHandler is the signature for a completion handler block.
type AXBrailleMapHandler = func(*AXBrailleMap)

// NewAXBrailleMapBlock wraps a Go [AXBrailleMapHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewAXBrailleMapBlock(handler AXBrailleMapHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *AXBrailleMap
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := AXBrailleMapFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// AXCustomContentVoidHandler is the signature for a completion handler block.
type AXCustomContentVoidHandler = func() []AXCustomContent

// StringFloat64Handler handles a primitive value and returns a primitive value.
//
// Used by:
//   - [AXNumericDataAxisDescriptor.InitWithAttributedTitleLowerBoundUpperBoundGridlinePositionsValueDescriptionProvider]
//   - [AXNumericDataAxisDescriptor.InitWithTitleLowerBoundUpperBoundGridlinePositionsValueDescriptionProvider]
type StringFloat64Handler = func(float64) string

// NewStringFloat64Block wraps a Go [StringFloat64Handler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AXNumericDataAxisDescriptor.InitWithAttributedTitleLowerBoundUpperBoundGridlinePositionsValueDescriptionProvider]
//   - [AXNumericDataAxisDescriptor.InitWithTitleLowerBoundUpperBoundGridlinePositionsValueDescriptionProvider]
func NewStringFloat64Block(handler StringFloat64Handler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal float64) string {
		return handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}
