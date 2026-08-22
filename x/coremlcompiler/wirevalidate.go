package coremlcompiler

// Preconditions that the protobuf wire format cannot express but Core ML
// enforces when it loads a model. Catching them here turns a confusing
// load-time rejection into an error at the point the model is built.
//
// Rules and their enforcement sites in coremltools:
//   - specificationVersion != 0                mlmodel/src/Model.cpp:66
//   - functions vs model-level features        Validation/InterfaceValidators.cpp:533
//   - defaultFunctionName names a function     Validation/InterfaceValidators.cpp:521
//   - feature name and type present            Validation/InterfaceValidators.cpp:36
//   - image colorSpace is a valid member       Validation/InterfaceValidators.cpp:358
//   - state feature constraints                Validation/InterfaceValidators.cpp:410
// Oneof exclusivity is not validated by coremltools at all: its generated
// setters make it unrepresentable. We hand-roll the encoder, so we check.

import "fmt"

// validateModelInterface checks the ModelDescription preconditions Core ML
// enforces at load time, plus the oneof exclusivity that generated protobuf
// code would guarantee for us.
func validateModelInterface(m *Model) error {
	if m.SpecVersion == 0 {
		return fmt.Errorf("model specification version is unset")
	}
	desc := m.Description
	multi := len(desc.Functions) > 0 || desc.DefaultFunctionName != ""
	if multi {
		// Model.proto reserves the model-level feature lists for the
		// single-function case; Core ML rejects a model that sets both.
		if len(desc.Inputs) > 0 || len(desc.Outputs) > 0 || len(desc.States) > 0 {
			return fmt.Errorf("multi-function description must leave model-level inputs, outputs and states empty")
		}
		found := false
		for _, fn := range desc.Functions {
			if fn.Name == desc.DefaultFunctionName {
				found = true
			}
			for _, group := range [][]FeatureDescription{fn.Inputs, fn.Outputs} {
				for _, fd := range group {
					if err := validateFeatureDescription(fd, false); err != nil {
						return fmt.Errorf("function %q: %w", fn.Name, err)
					}
				}
			}
			for _, fd := range fn.States {
				if err := validateFeatureDescription(fd, true); err != nil {
					return fmt.Errorf("function %q: %w", fn.Name, err)
				}
			}
		}
		if !found {
			return fmt.Errorf("default function name %q does not name a function description", desc.DefaultFunctionName)
		}
		return nil
	}
	for _, group := range [][]FeatureDescription{desc.Inputs, desc.Outputs} {
		for _, fd := range group {
			if err := validateFeatureDescription(fd, false); err != nil {
				return err
			}
		}
	}
	for _, fd := range desc.States {
		if err := validateFeatureDescription(fd, true); err != nil {
			return err
		}
	}
	return nil
}

func validateFeatureDescription(fd FeatureDescription, isState bool) error {
	if fd.Name == "" {
		return fmt.Errorf("feature description has an empty name")
	}
	if fd.Type == nil {
		return fmt.Errorf("feature %q has no type", fd.Name)
	}
	if err := validateFeatureType(fd.Type, isState); err != nil {
		return fmt.Errorf("feature %q: %w", fd.Name, err)
	}
	return nil
}

func validateFeatureType(ft *FeatureType, isState bool) error {
	set := 0
	for _, on := range []bool{
		ft.Int64Type, ft.DoubleType, ft.StringType,
		ft.ImageType != nil, ft.MultiArrayType != nil,
		ft.DictionaryType != nil, ft.SequenceType != nil,
		ft.StateArrayType != nil,
	} {
		if on {
			set++
		}
	}
	if set == 0 {
		return fmt.Errorf("feature type is unset")
	}
	if set > 1 {
		return fmt.Errorf("feature type sets %d members of a oneof, exactly one is allowed", set)
	}
	if img := ft.ImageType; img != nil {
		switch img.ColorSpace {
		case ColorSpaceGrayscale, ColorSpaceRGB, ColorSpaceBGR, ColorSpaceGrayscaleFloat16:
		default:
			return fmt.Errorf("image color space %d is not GRAYSCALE, RGB, BGR or GRAYSCALE_FLOAT16", int(img.ColorSpace))
		}
	}
	if seq := ft.SequenceType; seq != nil {
		et := seq.ElementType
		if et == nil || (!et.Int64Type && !et.StringType) {
			return fmt.Errorf("sequence element type must be int64 or string")
		}
	}
	if dict := ft.DictionaryType; dict != nil {
		if dict.KeyType != "int64" && dict.KeyType != "string" {
			return fmt.Errorf("dictionary key type %q must be int64 or string", dict.KeyType)
		}
	}
	// A state feature is carried as stateType; either spelling reaches
	// field 8, so both must satisfy the state constraints.
	arr := ft.StateArrayType
	if arr == nil && isState {
		arr = ft.MultiArrayType
	}
	if arr != nil {
		if ft.IsOptional {
			return fmt.Errorf("state feature cannot be optional")
		}
		if len(arr.Shape) == 0 {
			return fmt.Errorf("state feature has no shape")
		}
		for _, d := range arr.Shape {
			if d < 0 {
				return fmt.Errorf("state feature shape has negative dimension %d", d)
			}
		}
		if arr.DataType != ArrayDataTypeFloat16 && arr.DataType != ArrayDataTypeInt8 {
			return fmt.Errorf("state feature data type %d must be FLOAT16 or INT8", int(arr.DataType))
		}
	}
	return nil
}
