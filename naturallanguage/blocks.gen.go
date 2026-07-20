// Code generated from Apple documentation. DO NOT EDIT.

package naturallanguage

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// NLContextualEmbeddingAssetsResultErrorHandler handles A closure that notifies your app when the asset request completes.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NLContextualEmbedding.RequestEmbeddingAssetsWithCompletionHandler]
type NLContextualEmbeddingAssetsResultErrorHandler = func(NLContextualEmbeddingAssetsResult, error)

// NewNLContextualEmbeddingAssetsResultErrorBlock wraps a Go [NLContextualEmbeddingAssetsResultErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NLContextualEmbedding.RequestEmbeddingAssetsWithCompletionHandler]
func NewNLContextualEmbeddingAssetsResultErrorBlock(handler NLContextualEmbeddingAssetsResultErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal NLContextualEmbeddingAssetsResult, errID objc.ID) {
		handler(primitiveVal, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// NLTaggerAssetsResultErrorHandler handles A closure the framework uses to notify your app when the tag scheme request has completed.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NLTagger.RequestAssetsForLanguageTagSchemeCompletionHandler]
type NLTaggerAssetsResultErrorHandler = func(NLTaggerAssetsResult, error)

// NewNLTaggerAssetsResultErrorBlock wraps a Go [NLTaggerAssetsResultErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NLTagger.RequestAssetsForLanguageTagSchemeCompletionHandler]
func NewNLTaggerAssetsResultErrorBlock(handler NLTaggerAssetsResultErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal NLTaggerAssetsResult, errID objc.ID) {
		handler(primitiveVal, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// NSNumberArrayNSRangeBoolHandler handles A block that contains each token’s embedding vector and its corresponding character range in the string.
//
// Used by:
//   - [NLContextualEmbeddingResult.EnumerateTokenVectorsInRangeUsingBlock]
type NSNumberArrayNSRangeBoolHandler = func(*[]foundation.NSNumber, foundation.NSRange, bool)

// NewNSNumberArrayNSRangeBoolBlock wraps a Go [NSNumberArrayNSRangeBoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NLContextualEmbeddingResult.EnumerateTokenVectorsInRangeUsingBlock]
func NewNSNumberArrayNSRangeBoolBlock(handler NSNumberArrayNSRangeBoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 foundation.NSRange, extra1 bool) {
		var result *[]foundation.NSNumber
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			obj := foundation.NSArrayFromID(resultID)
			count := obj.Count()
			res := make([]foundation.NSNumber, count)
			for i := uint(0); i < count; i++ {
				item := obj.ObjectAtIndex(i)
				res[i] = foundation.NSNumberFromID(item.GetID())
			}
			result = &res
		}
		handler(result, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// NSRangeNLTokenizerAttributesBoolHandler handles The closure to call after each token; return false if processing should stop.
//
// Used by:
//   - [NLTokenizer.EnumerateTokensInRangeUsingBlock]
type NSRangeNLTokenizerAttributesBoolHandler = func(foundation.NSRange, NLTokenizerAttributes, bool)

// NewNSRangeNLTokenizerAttributesBoolBlock wraps a Go [NSRangeNLTokenizerAttributesBoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NLTokenizer.EnumerateTokensInRangeUsingBlock]
func NewNSRangeNLTokenizerAttributesBoolBlock(handler NSRangeNLTokenizerAttributesBoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive foundation.NSRange, extra0 NLTokenizerAttributes, extra1 bool) {
		handler(primitive, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// StringNLDistanceBoolHandler handles A block with the following parameters:
//
// Used by:
//   - [NLEmbedding.EnumerateNeighborsForStringMaximumCountDistanceTypeUsingBlock]
//   - [NLEmbedding.EnumerateNeighborsForStringMaximumCountMaximumDistanceDistanceTypeUsingBlock]
//   - [NLEmbedding.EnumerateNeighborsForVectorMaximumCountDistanceTypeUsingBlock]
//   - [NLEmbedding.EnumerateNeighborsForVectorMaximumCountMaximumDistanceDistanceTypeUsingBlock]
type StringNLDistanceBoolHandler = func(*string, NLDistance, bool)

// NewStringNLDistanceBoolBlock wraps a Go [StringNLDistanceBoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NLEmbedding.EnumerateNeighborsForStringMaximumCountDistanceTypeUsingBlock]
//   - [NLEmbedding.EnumerateNeighborsForStringMaximumCountMaximumDistanceDistanceTypeUsingBlock]
//   - [NLEmbedding.EnumerateNeighborsForVectorMaximumCountDistanceTypeUsingBlock]
//   - [NLEmbedding.EnumerateNeighborsForVectorMaximumCountMaximumDistanceDistanceTypeUsingBlock]
func NewStringNLDistanceBoolBlock(handler StringNLDistanceBoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 NLDistance, extra1 bool) {
		var result *string
		if resultID != 0 {
			v := objc.IDToString(resultID)
			result = &v
		}
		handler(result, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// StringNSRangeBoolHandler handles The block this method uses to iterate over the tagger’s string property.
//   - tag: The tag of the token.
//   - tokenRange: The range of the token.
//   - stop: A reference to a Boolean value. The block can set the value to `true` to stop further processing of the set. The `stop` argument is an out-only argument. You should only ever set this Boolean to `true` within the block.
//
// Used by:
//   - [NLTagger.EnumerateTagsInRangeUnitSchemeOptionsUsingBlock]
type StringNSRangeBoolHandler = func(*string, foundation.NSRange, bool)

// NewStringNSRangeBoolBlock wraps a Go [StringNSRangeBoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NLTagger.EnumerateTagsInRangeUnitSchemeOptionsUsingBlock]
func NewStringNSRangeBoolBlock(handler StringNSRangeBoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 foundation.NSRange, extra1 bool) {
		var result *string
		if resultID != 0 {
			v := objc.IDToString(resultID)
			result = &v
		}
		handler(result, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}
