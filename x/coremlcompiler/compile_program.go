package coremlcompiler

import "fmt"

// CompileProgram compiles a [Program] built in Go into a compiled bundle at
// outputPath.
//
// It is the path for a producer that has a structured program: unlike
// [CompileMILText] it validates before emitting, so the program is held to the
// same rules a model decoded from a .mlpackage is. It runs [ValidateProgram],
// the per-op input specifications, the model interface rules, and the check
// that desc matches the program's own signature; then emits the MIL text
// itself, so text and description cannot disagree.
//
// desc names the model inputs, outputs, and states stored in coremldata.bin
// and metadata.json; [TensorFeatureDescription] and [StateFeatureDescription]
// build its entries. specVersion selects the MIL dialect, defaulting to the
// iOS18 opset. weightRoot, when non-empty, is the directory holding the weight
// files the program's BLOBFILE references name, as written by
// [WriteWeightRoot]; only referenced files are copied into the bundle.
func CompileProgram(prog *Program, specVersion int32, desc ModelDescription, weightRoot, outputPath string) error {
	if prog == nil {
		return fmt.Errorf("coremlcompiler: program is nil")
	}
	if len(prog.Functions) == 0 {
		return fmt.Errorf("coremlcompiler: program has no functions")
	}
	if specVersion == 0 {
		specVersion = defaultMILSpecVersion
	}
	model := &Model{
		SpecVersion:    specVersion,
		Description:    desc,
		MLProgram:      prog,
		descriptionRaw: encodeModelDescription(desc),
	}
	return compileMLProgram(model, weightRoot, outputPath)
}
