package coremlcompiler

// Cross-check of the ModelDescription against the MIL function signature.
//
// coremltools never needs this check: it derives the description from the MIL
// vars it just built (backend/mil/load.py:640-700 get_func_input,
// load.py:860-1000 get_func_output), so the two sides cannot disagree. We
// start from an already-serialized model, where the description is input, so
// the two sides can and do disagree. Core ML rejects such a model at load time
// without naming the offending feature.
//
// The rules mirrored here:
//   - state inputs are the MIL inputs whose type is stateType
//     (load.py:617-635 _decouple_state_and_input)
//   - the default function is "main" unless the description names another
//     (mil/mil/program.py:34)
//   - element type mapping (backend/mil/helper.py:304-324
//     cast_to_framework_io_dtype)
//   - image features have no MIL image type; the MIL side must be a rank-4
//     tensor (load.py:665-673, backend/backend_helper.py:58-62)
//   - a symbolic dimension carries no shape obligation: coremltools
//     substitutes the user's default shape for symbolic inputs
//     (load.py:687-693) and writes no shape at all for symbolic outputs
//     (load.py:925-929)

import "fmt"

// validateDescriptionSignature checks that each described function's features
// name and type the same values as the MIL function's signature.
func validateDescriptionSignature(m *Model) error {
	if m == nil || m.MLProgram == nil {
		return nil
	}
	desc := m.Description
	if len(desc.Functions) > 0 {
		for _, fd := range desc.Functions {
			fn, ok := m.MLProgram.Functions[fd.Name]
			if !ok {
				// validateFunctionDescriptions owns this rule.
				continue
			}
			if err := matchFunctionSignature(fd.Name, fn, fd.Inputs, fd.Outputs, fd.States); err != nil {
				return err
			}
		}
		return nil
	}
	name := desc.DefaultFunctionName
	if name == "" {
		name = "main"
	}
	fn, ok := m.MLProgram.Functions[name]
	if !ok {
		return nil
	}
	return matchFunctionSignature(name, fn, desc.Inputs, desc.Outputs, desc.States)
}

// matchFunctionSignature compares one function's MIL signature with the
// description features that address it.
func matchFunctionSignature(name string, fn *Function, inputs, outputs, states []FeatureDescription) error {
	if fn == nil {
		return nil
	}
	milInputs := make(map[string]*ValueType)
	milStates := make(map[string]*ValueType)
	var inputOrder, stateOrder []string
	for _, in := range fn.Inputs {
		if in.Type != nil && in.Type.StateType != nil {
			milStates[in.Name] = in.Type
			stateOrder = append(stateOrder, in.Name)
			continue
		}
		milInputs[in.Name] = in.Type
		inputOrder = append(inputOrder, in.Name)
	}

	milOutputs := make(map[string]*ValueType)
	var outputOrder []string
	if blk := fn.BlockSpecializations[fn.OpSet]; blk != nil {
		defs := make(map[string]*ValueType, len(fn.Inputs))
		for _, in := range fn.Inputs {
			defs[in.Name] = in.Type
		}
		for _, op := range blk.Operations {
			for _, out := range op.Outputs {
				defs[out.Name] = out.Type
			}
		}
		for _, out := range blk.Outputs {
			t, ok := defs[out]
			if !ok {
				// validateBlockInvariants owns def visibility.
				continue
			}
			milOutputs[out] = t
			outputOrder = append(outputOrder, out)
		}
	}

	for _, c := range []struct {
		kind  string
		feats []FeatureDescription
		mil   map[string]*ValueType
		order []string
	}{
		{"input", inputs, milInputs, inputOrder},
		{"output", outputs, milOutputs, outputOrder},
		{"state", states, milStates, stateOrder},
	} {
		declared := make(map[string]bool, len(c.feats))
		for _, fd := range c.feats {
			declared[fd.Name] = true
			t, ok := c.mil[fd.Name]
			if !ok {
				return fmt.Errorf("function %q: description declares %s %q, which the program's signature does not have", name, c.kind, fd.Name)
			}
			if err := matchFeature(name, c.kind, fd, t); err != nil {
				return err
			}
		}
		for _, mn := range c.order {
			if !declared[mn] {
				return fmt.Errorf("function %q: program %s %q has no description feature", name, c.kind, mn)
			}
		}
	}
	return nil
}

// matchFeature compares one description feature with the MIL type of the value
// it names.
func matchFeature(fname, kind string, fd FeatureDescription, t *ValueType) error {
	if fd.Type == nil || t == nil {
		return nil
	}
	if st := t.StateType; st != nil {
		t = st.WrappedType
		if t == nil {
			return nil
		}
	}
	tensor := t.TensorType
	if img := fd.Type.ImageType; img != nil {
		// An image feature has no MIL counterpart type; the MIL side must be
		// a rank-4 tensor.
		if tensor == nil {
			return fmt.Errorf("function %q: image %s %q is not a tensor in the program", fname, kind, fd.Name)
		}
		if tensorRank(tensor) != 4 {
			return fmt.Errorf("function %q: image %s %q must have rank 4, program has rank %d", fname, kind, fd.Name, tensorRank(tensor))
		}
		return nil
	}
	arr := fd.Type.StateArrayType
	if arr == nil {
		arr = fd.Type.MultiArrayType
	}
	if arr == nil || tensor == nil {
		// Dictionary, sequence and scalar features have no array type to
		// compare; classifier outputs take that shape.
		return nil
	}

	want, err := featureDataType(tensor.DataType)
	if err != nil {
		return fmt.Errorf("function %q: %s %q: %w", fname, kind, fd.Name, err)
	}
	if arr.DataType != want {
		return fmt.Errorf("function %q: %s %q has description data type %d, program has %v (%d)", fname, kind, fd.Name, int(arr.DataType), tensor.DataType, int(want))
	}

	if len(arr.Shape) == 0 {
		return nil
	}
	dims := make([]int64, 0, len(tensor.Dimensions))
	for _, d := range tensor.Dimensions {
		if d.Unknown || d.Variadic {
			return nil
		}
		dims = append(dims, int64(d.Constant))
	}
	if len(dims) != len(arr.Shape) {
		return fmt.Errorf("function %q: %s %q has description shape %v, program has %v", fname, kind, fd.Name, arr.Shape, dims)
	}
	for i, d := range dims {
		if arr.Shape[i] != d {
			return fmt.Errorf("function %q: %s %q has description shape %v, program has %v", fname, kind, fd.Name, arr.Shape, dims)
		}
	}
	return nil
}

// featureDataType maps a MIL element type to the Core ML feature data type,
// mirroring backend/mil/helper.py:304-324 cast_to_framework_io_dtype.
func featureDataType(dt DataType) (ArrayDataType, error) {
	switch dt {
	case DataTypeFloat32:
		return ArrayDataTypeFloat32, nil
	case DataTypeInt32:
		return ArrayDataTypeInt32, nil
	case DataTypeFloat16:
		return ArrayDataTypeFloat16, nil
	case DataTypeInt8:
		return ArrayDataTypeInt8, nil
	}
	return 0, fmt.Errorf("data type %v has no core ml feature data type; ml program models only support fp32, fp16, int32 and int8 at the boundary", dt)
}
