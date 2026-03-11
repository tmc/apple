package coremlcompiler

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// metadata generates the metadata.json for a compiled model bundle.
// This is informational and may not be required by the runtime, but
// Apple's compiler always generates it.
func buildMetadataJSON(m *Model) ([]byte, error) {
	inputs := make([]map[string]string, len(m.Description.Inputs))
	for i, in := range m.Description.Inputs {
		inputs[i] = metadataFeatureSchema(in, false)
	}
	outputs := make([]map[string]string, len(m.Description.Outputs))
	for i, out := range m.Description.Outputs {
		outputs[i] = metadataFeatureSchema(out, false)
	}
	states := make([]map[string]string, len(m.Description.States))
	for i, state := range m.Description.States {
		states[i] = metadataFeatureSchema(state, true)
	}

	modelType := "mlProgram"
	if m.MLProgram == nil {
		modelType = "neuralNetwork"
	}

	meta := []map[string]any{
		{
			"metadataOutputVersion": "3.0",
			"storagePrecision":      "Float16",
			"specificationVersion":  m.SpecVersion,
			"modelType":             map[string]string{"name": fmt.Sprintf("MLModelType_%s", modelType)},
			"inputSchema":           inputs,
			"outputSchema":          outputs,
			"stateSchema":           states,
			"modelParameters":       []any{},
			"userDefinedMetadata":   map[string]string{},
			"generatorVersion":      "github.com/tmc/apple/x/coremlcompiler",
		},
	}

	return json.MarshalIndent(meta, "", "  ")
}

func metadataFeatureSchema(fd FeatureDescription, isState bool) map[string]string {
	dataType := "Float32"
	shape := "[]"
	formatted := "MultiArray (Float32)"
	typ := "MultiArray"
	if isState {
		typ = "State"
		formatted = "State (Float32)"
	}
	if fd.Type != nil && fd.Type.MultiArrayType != nil {
		dataType = metadataArrayDataType(fd.Type.MultiArrayType.DataType)
		shape = metadataShape(fd.Type.MultiArrayType.Shape)
		formatted = fmt.Sprintf("%s (%s%s)", typ, dataType, metadataFormattedShape(fd.Type.MultiArrayType.Shape))
	}
	return map[string]string{
		"name":                fd.Name,
		"type":                typ,
		"dataType":            dataType,
		"formattedType":       formatted,
		"shortDescription":    "",
		"shape":               shape,
		"isOptional":          boolString(fd.Type != nil && fd.Type.IsOptional),
		"hasShapeFlexibility": "0",
	}
}

func metadataArrayDataType(dt ArrayDataType) string {
	switch dt {
	case ArrayDataTypeFloat16:
		return "Float16"
	case ArrayDataTypeFloat32:
		return "Float32"
	case ArrayDataTypeDouble:
		return "Double"
	case ArrayDataTypeInt32:
		return "Int32"
	default:
		return "Float32"
	}
}

func metadataShape(shape []int64) string {
	if len(shape) == 0 {
		return "[]"
	}
	parts := make([]string, len(shape))
	for i, dim := range shape {
		parts[i] = strconv.FormatInt(dim, 10)
	}
	return "[" + strings.Join(parts, ", ") + "]"
}

func metadataFormattedShape(shape []int64) string {
	if len(shape) == 0 {
		return ""
	}
	parts := make([]string, len(shape))
	for i, dim := range shape {
		parts[i] = strconv.FormatInt(dim, 10)
	}
	return " " + strings.Join(parts, " × ")
}

func boolString(v bool) string {
	if v {
		return "1"
	}
	return "0"
}
