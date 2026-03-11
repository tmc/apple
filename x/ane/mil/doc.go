//go:build darwin

// Package mil generates MIL (Model Intermediate Language) programs and weight
// blobs for Apple Neural Engine compilation.
//
//	text := mil.GenConv(16, 16, 1)
//	blob, _ := mil.BuildWeightBlob(weights, 16, 16)
package mil
