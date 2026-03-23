// Package coremlcompiler compiles CoreML model packages (.mlpackage)
// into compiled model bundles (.mlmodelc) without requiring xcrun or
// Apple's proprietary coremlcompiler binary.
//
// Only mlprogram models (spec version 5+) are supported. Legacy
// NeuralNetwork models must first be converted to mlprogram format
// using coremltools (convert_to="mlprogram").
//
// The compilation path converts mlprogram models: the protobuf-encoded
// MIL program is serialized to MIL text, weights are copied byte-for-byte,
// and a coremldata.bin header is generated. The result can be loaded
// directly by CoreML, the ANE runtime, or x/ane.
//
// Three entry points cover different use cases:
//
// [Compile] takes a .mlpackage directory or .mlmodel file on disk and
// produces a compiled .mlmodelc bundle. This is the simplest path and
// handles reading the protobuf, locating weights, and writing all
// output files:
//
//	err := coremlcompiler.Compile("model.mlpackage", "model.mlmodelc")
//
// [CompileMLProgram] accepts raw model protobuf bytes and a weight
// directory. Use this when you already have the serialized model in
// memory or need to supply the weight directory separately:
//
//	err := coremlcompiler.CompileMLProgram(protoBytes, weightDir, "model.mlmodelc")
//
// [CompileMILText] takes pre-built MIL text, a [ModelDescription], and
// an optional weight root directory. Use this when you generate MIL
// text programmatically rather than decoding it from a protobuf:
//
//	err := coremlcompiler.CompileMILText(milText, 8, desc, weightRoot, "model.mlmodelc")
package coremlcompiler
