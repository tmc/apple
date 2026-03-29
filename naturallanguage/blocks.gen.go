// Code generated from Apple documentation. DO NOT EDIT.

package naturallanguage

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// NLContextualEmbeddingAssetsResultErrorHandler handles A completion handler the system calls after it finishes the request.
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
	block := objc.NewBlock(func(b objc.Block, primitiveVal NLContextualEmbeddingAssetsResult, errID objc.ID) {
		handler(primitiveVal, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// NLDistanceHandler handles A block with the following parameters:
//
// Used by:
//   - [NLEmbedding.EnumerateNeighborsForStringMaximumCountDistanceTypeUsingBlock]
//   - [NLEmbedding.EnumerateNeighborsForStringMaximumCountMaximumDistanceDistanceTypeUsingBlock]
//   - [NLEmbedding.EnumerateNeighborsForVectorMaximumCountDistanceTypeUsingBlock]
//   - [NLEmbedding.EnumerateNeighborsForVectorMaximumCountMaximumDistanceDistanceTypeUsingBlock]
type NLDistanceHandler = func(*string)

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
	block := objc.NewBlock(func(b objc.Block, primitiveVal NLTaggerAssetsResult, errID objc.ID) {
		handler(primitiveVal, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// RangeHandler handles The closure to call after each token; return false if processing should stop.
//   - tag: The tag of the token.
//   - tokenRange: The range of the token.
//   - stop: A reference to a Boolean value. The block can set the value to `true` to stop further processing of the set. The `stop` argument is an out-only argument. You should only ever set this Boolean to `true` within the block.
//
// Used by:
//   - [NLContextualEmbeddingResult.EnumerateTokenVectorsInRangeUsingBlock]
//   - [NLTagger.EnumerateTagsInRangeUnitSchemeOptionsUsingBlock]
//   - [NLTokenizer.EnumerateTokensInRangeUsingBlock]
type RangeHandler = func(foundation.NSRange)

// NewRangeBlock wraps a Go [RangeHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NLContextualEmbeddingResult.EnumerateTokenVectorsInRangeUsingBlock]
//   - [NLTagger.EnumerateTagsInRangeUnitSchemeOptionsUsingBlock]
//   - [NLTokenizer.EnumerateTokensInRangeUsingBlock]
func NewRangeBlock(handler RangeHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal foundation.NSRange) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

