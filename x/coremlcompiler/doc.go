// Package coremlcompiler compiles CoreML model packages (.mlpackage)
// into compiled model bundles (.mlmodelc) without requiring xcrun or
// Apple's proprietary coremlcompiler binary.
//
// The primary path converts mlprogram models: the protobuf-encoded MIL
// program is serialized to MIL text, weights are copied byte-for-byte,
// and a coremldata.bin header is generated. The result can be loaded
// directly by CoreML, the ANE runtime, or x/ane.
//
//	err := coremlcompiler.Compile("model.mlpackage", "model.mlmodelc")
package coremlcompiler
