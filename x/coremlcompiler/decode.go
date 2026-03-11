package coremlcompiler

import "fmt"

// decodeModel reads a CoreML Model from protobuf bytes.
func decodeModel(data []byte) (*Model, error) {
	m := &Model{}
	r := newProtoReader(data)
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return nil, fmt.Errorf("decode model: %w", err)
		}
		switch field {
		case 1: // specificationVersion: int32
			v, err := r.readVarint()
			if err != nil {
				return nil, err
			}
			m.SpecVersion = int32(v)
		case 2: // description: ModelDescription
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			m.descriptionRaw = raw
			desc, err := decodeModelDescription(raw)
			if err != nil {
				return nil, fmt.Errorf("decode model description: %w", err)
			}
			m.Description = *desc
		case 502: // mlProgram: MILSpec.Program
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			prog, err := decodeProgram(raw)
			if err != nil {
				return nil, fmt.Errorf("decode mlprogram: %w", err)
			}
			m.MLProgram = prog
		default:
			if err := r.skip(wire); err != nil {
				return nil, fmt.Errorf("decode model: skip field %d: %w", field, err)
			}
		}
	}
	return m, nil
}

func decodeModelDescription(data []byte) (*ModelDescription, error) {
	d := &ModelDescription{}
	r := newProtoReader(data)
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		switch field {
		case 1: // input: repeated FeatureDescription
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			fd, err := decodeFeatureDescription(raw)
			if err != nil {
				return nil, err
			}
			d.Inputs = append(d.Inputs, *fd)
		case 10: // output: repeated FeatureDescription
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			fd, err := decodeFeatureDescription(raw)
			if err != nil {
				return nil, err
			}
			d.Outputs = append(d.Outputs, *fd)
		case 13: // state: repeated FeatureDescription
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			fd, err := decodeFeatureDescription(raw)
			if err != nil {
				return nil, err
			}
			d.States = append(d.States, *fd)
		default:
			if err := r.skip(wire); err != nil {
				return nil, err
			}
		}
	}
	return d, nil
}

func decodeFeatureDescription(data []byte) (*FeatureDescription, error) {
	fd := &FeatureDescription{}
	r := newProtoReader(data)
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		switch field {
		case 1: // name: string
			fd.Name, err = r.readString()
			if err != nil {
				return nil, err
			}
		default:
			if err := r.skip(wire); err != nil {
				return nil, err
			}
		}
	}
	return fd, nil
}

func decodeProgram(data []byte) (*program, error) {
	p := &program{
		Functions:  make(map[string]*function),
		Attributes: make(map[string]*value),
	}
	r := newProtoReader(data)
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		switch field {
		case 1: // version: int64
			v, err := r.readVarint()
			if err != nil {
				return nil, err
			}
			p.Version = int64(v)
		case 2: // functions: map<string, Function>
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			key, val, err := decodeMapEntry(raw, decodeFunction)
			if err != nil {
				return nil, fmt.Errorf("decode program function: %w", err)
			}
			p.Functions[key] = val
		case 4: // attributes: map<string, Value>
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			key, val, err := decodeMapEntry(raw, decodeValue)
			if err != nil {
				return nil, fmt.Errorf("decode program attribute: %w", err)
			}
			p.Attributes[key] = val
		default:
			if err := r.skip(wire); err != nil {
				return nil, err
			}
		}
	}
	return p, nil
}

func decodeFunction(data []byte) (*function, error) {
	f := &function{
		BlockSpecializations: make(map[string]*block),
		Attributes:           make(map[string]*value),
	}
	r := newProtoReader(data)
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		switch field {
		case 1: // inputs: repeated NamedValueType
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			nvt, err := decodeNamedValueType(raw)
			if err != nil {
				return nil, err
			}
			f.Inputs = append(f.Inputs, *nvt)
		case 2: // opset: string
			f.OpSet, err = r.readString()
			if err != nil {
				return nil, err
			}
		case 3: // block_specializations: map<string, Block>
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			key, val, err := decodeMapEntry(raw, decodeBlock)
			if err != nil {
				return nil, err
			}
			f.BlockSpecializations[key] = val
		case 4: // attributes: map<string, Value>
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			key, val, err := decodeMapEntry(raw, decodeValue)
			if err != nil {
				return nil, err
			}
			f.Attributes[key] = val
		default:
			if err := r.skip(wire); err != nil {
				return nil, err
			}
		}
	}
	return f, nil
}

func decodeBlock(data []byte) (*block, error) {
	b := &block{}
	r := newProtoReader(data)
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		switch field {
		case 1: // inputs: repeated NamedValueType
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			nvt, err := decodeNamedValueType(raw)
			if err != nil {
				return nil, err
			}
			b.Inputs = append(b.Inputs, *nvt)
		case 2: // outputs: repeated string
			s, err := r.readString()
			if err != nil {
				return nil, err
			}
			b.Outputs = append(b.Outputs, s)
		case 3: // operations: repeated Operation
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			op, err := decodeOperation(raw)
			if err != nil {
				return nil, err
			}
			b.Operations = append(b.Operations, op)
		default:
			if err := r.skip(wire); err != nil {
				return nil, err
			}
		}
	}
	return b, nil
}

func decodeOperation(data []byte) (*operation, error) {
	op := &operation{
		Inputs:     make(map[string]*argument),
		Attributes: make(map[string]*value),
	}
	r := newProtoReader(data)
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		switch field {
		case 1: // type: string
			op.Type, err = r.readString()
			if err != nil {
				return nil, err
			}
		case 2: // inputs: map<string, Argument>
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			key, val, err := decodeMapEntry(raw, decodeArgument)
			if err != nil {
				return nil, err
			}
			op.Inputs[key] = val
		case 3: // outputs: repeated NamedValueType
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			nvt, err := decodeNamedValueType(raw)
			if err != nil {
				return nil, err
			}
			op.Outputs = append(op.Outputs, *nvt)
		case 4: // blocks: repeated Block
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			blk, err := decodeBlock(raw)
			if err != nil {
				return nil, err
			}
			op.Blocks = append(op.Blocks, blk)
		case 5: // attributes: map<string, Value>
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			key, val, err := decodeMapEntry(raw, decodeValue)
			if err != nil {
				return nil, err
			}
			op.Attributes[key] = val
		default:
			if err := r.skip(wire); err != nil {
				return nil, err
			}
		}
	}
	return op, nil
}

func decodeNamedValueType(data []byte) (*namedValueType, error) {
	nvt := &namedValueType{}
	r := newProtoReader(data)
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		switch field {
		case 1: // name: string
			nvt.Name, err = r.readString()
			if err != nil {
				return nil, err
			}
		case 2: // type: ValueType
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			vt, err := decodeValueType(raw)
			if err != nil {
				return nil, err
			}
			nvt.Type = vt
		default:
			if err := r.skip(wire); err != nil {
				return nil, err
			}
		}
	}
	return nvt, nil
}

func decodeValueType(data []byte) (*valueType, error) {
	vt := &valueType{}
	r := newProtoReader(data)
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		switch field {
		case 1: // tensorType: TensorType
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			tt, err := decodeTensorType(raw)
			if err != nil {
				return nil, err
			}
			vt.TensorType = tt
		case 5: // stateType: StateType
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			st, err := decodeStateType(raw)
			if err != nil {
				return nil, err
			}
			vt.StateType = st
		default:
			if err := r.skip(wire); err != nil {
				return nil, err
			}
		}
	}
	return vt, nil
}

func decodeTensorType(data []byte) (*tensorType, error) {
	tt := &tensorType{}
	r := newProtoReader(data)
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		switch field {
		case 1: // dataType: DataType
			v, err := r.readVarint()
			if err != nil {
				return nil, err
			}
			tt.DataType = DataType(v)
		case 2: // rank: int64
			v, err := r.readVarint()
			if err != nil {
				return nil, err
			}
			tt.Rank = int64(v)
		case 3: // dimensions: repeated Dimension
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			dim, err := decodeDimension(raw)
			if err != nil {
				return nil, err
			}
			tt.Dimensions = append(tt.Dimensions, *dim)
		default:
			if err := r.skip(wire); err != nil {
				return nil, err
			}
		}
	}
	return tt, nil
}

func decodeStateType(data []byte) (*stateType, error) {
	st := &stateType{}
	r := newProtoReader(data)
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		switch field {
		case 1: // wrappedType: ValueType
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			vt, err := decodeValueType(raw)
			if err != nil {
				return nil, err
			}
			st.WrappedType = vt
		default:
			if err := r.skip(wire); err != nil {
				return nil, err
			}
		}
	}
	return st, nil
}

func decodeDimension(data []byte) (*dimension, error) {
	d := &dimension{}
	r := newProtoReader(data)
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		switch field {
		case 1: // constant: ConstantDimension
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			cr := newProtoReader(raw)
			for !cr.done() {
				cf, cw, err := cr.readTag()
				if err != nil {
					return nil, err
				}
				if cf == 1 { // size: uint64
					v, err := cr.readVarint()
					if err != nil {
						return nil, err
					}
					d.Constant = v
				} else {
					if err := cr.skip(cw); err != nil {
						return nil, err
					}
				}
			}
		case 2: // unknown: UnknownDimension
			d.Unknown = true
			if wire == wireBytes {
				if _, err := r.readBytes(); err != nil {
					return nil, err
				}
			} else {
				if err := r.skip(wire); err != nil {
					return nil, err
				}
			}
		default:
			if err := r.skip(wire); err != nil {
				return nil, err
			}
		}
	}
	return d, nil
}

func decodeValue(data []byte) (*value, error) {
	v := &value{}
	r := newProtoReader(data)
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		switch field {
		case 2: // type: ValueType
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			vt, err := decodeValueType(raw)
			if err != nil {
				return nil, err
			}
			v.Type = vt
		case 3: // immediateValue: ImmediateValue
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			iv, err := decodeImmediateValue(raw)
			if err != nil {
				return nil, err
			}
			v.Immediate = iv
		case 5: // blobFileValue: BlobFileValue
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			bf, err := decodeBlobFileValue(raw)
			if err != nil {
				return nil, err
			}
			v.BlobFile = bf
		default:
			if err := r.skip(wire); err != nil {
				return nil, err
			}
		}
	}
	return v, nil
}

func decodeImmediateValue(data []byte) (*immediateValue, error) {
	iv := &immediateValue{}
	r := newProtoReader(data)
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		switch field {
		case 1: // tensor: TensorValue
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			tv, err := decodeTensorValue(raw)
			if err != nil {
				return nil, err
			}
			iv.Tensor = tv
		default:
			if err := r.skip(wire); err != nil {
				return nil, err
			}
		}
	}
	return iv, nil
}

func decodeTensorValue(data []byte) (*tensorValue, error) {
	tv := &tensorValue{}
	r := newProtoReader(data)
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		switch field {
		case 1: // floats: RepeatedFloats
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			tv.Floats, err = decodeRepeatedFloats(raw)
			if err != nil {
				return nil, err
			}
		case 2: // ints: RepeatedInts
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			tv.Ints, err = decodeRepeatedInts(raw)
			if err != nil {
				return nil, err
			}
		case 3: // bools: RepeatedBools
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			tv.Bools, err = decodeRepeatedBools(raw)
			if err != nil {
				return nil, err
			}
		case 4: // strings: RepeatedStrings
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			tv.Strings, err = decodeRepeatedStrings(raw)
			if err != nil {
				return nil, err
			}
		case 5: // longs: RepeatedLongInts
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			tv.Longs, err = decodeRepeatedLongs(raw)
			if err != nil {
				return nil, err
			}
		case 6: // doubles: RepeatedDoubles
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			tv.Doubles, err = decodeRepeatedDoubles(raw)
			if err != nil {
				return nil, err
			}
		case 7: // bytes: RepeatedBytes
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			// RepeatedBytes has one field: bytes values = 1
			br := newProtoReader(raw)
			for !br.done() {
				bf, bw, err := br.readTag()
				if err != nil {
					return nil, err
				}
				if bf == 1 {
					tv.Bytes, err = br.readBytes()
					if err != nil {
						return nil, err
					}
				} else {
					if err := br.skip(bw); err != nil {
						return nil, err
					}
				}
			}
		default:
			if err := r.skip(wire); err != nil {
				return nil, err
			}
		}
	}
	return tv, nil
}

// decodeRepeatedFloats decodes a RepeatedFloats submessage (field 1: packed float).
func decodeRepeatedFloats(data []byte) ([]float32, error) {
	r := newProtoReader(data)
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		if field == 1 {
			return r.readPackedFloat32()
		}
		if err := r.skip(wire); err != nil {
			return nil, err
		}
	}
	return nil, nil
}

// decodeRepeatedDoubles decodes a RepeatedDoubles submessage.
func decodeRepeatedDoubles(data []byte) ([]float64, error) {
	r := newProtoReader(data)
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		if field == 1 {
			return r.readPackedFloat64()
		}
		if err := r.skip(wire); err != nil {
			return nil, err
		}
	}
	return nil, nil
}

// decodeRepeatedInts decodes a RepeatedInts submessage (field 1: packed int32 varint).
func decodeRepeatedInts(data []byte) ([]int32, error) {
	r := newProtoReader(data)
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		if field == 1 {
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			return readPackedInt32(raw)
		}
		if err := r.skip(wire); err != nil {
			return nil, err
		}
	}
	return nil, nil
}

// decodeRepeatedLongs decodes a RepeatedLongInts submessage.
func decodeRepeatedLongs(data []byte) ([]int64, error) {
	r := newProtoReader(data)
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		if field == 1 {
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			return readPackedInt64(raw)
		}
		if err := r.skip(wire); err != nil {
			return nil, err
		}
	}
	return nil, nil
}

// decodeRepeatedBools decodes a RepeatedBools submessage (field 1: packed bool varint).
func decodeRepeatedBools(data []byte) ([]bool, error) {
	r := newProtoReader(data)
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		if field == 1 {
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			pr := newProtoReader(raw)
			var out []bool
			for !pr.done() {
				v, err := pr.readVarint()
				if err != nil {
					return nil, err
				}
				out = append(out, v != 0)
			}
			return out, nil
		}
		if err := r.skip(wire); err != nil {
			return nil, err
		}
	}
	return nil, nil
}

// decodeRepeatedStrings decodes a RepeatedStrings submessage (field 1: repeated string).
func decodeRepeatedStrings(data []byte) ([]string, error) {
	r := newProtoReader(data)
	var out []string
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		if field == 1 {
			s, err := r.readString()
			if err != nil {
				return nil, err
			}
			out = append(out, s)
		} else {
			if err := r.skip(wire); err != nil {
				return nil, err
			}
		}
	}
	return out, nil
}

func decodeBlobFileValue(data []byte) (*blobFileValue, error) {
	bf := &blobFileValue{}
	r := newProtoReader(data)
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		switch field {
		case 1: // fileName: string
			bf.FileName, err = r.readString()
			if err != nil {
				return nil, err
			}
		case 2: // offset: uint64
			bf.Offset, err = r.readVarint()
			if err != nil {
				return nil, err
			}
		default:
			if err := r.skip(wire); err != nil {
				return nil, err
			}
		}
	}
	return bf, nil
}

func decodeArgument(data []byte) (*argument, error) {
	arg := &argument{}
	r := newProtoReader(data)
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		switch field {
		case 1: // arguments: repeated Binding
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			b, err := decodeBinding(raw)
			if err != nil {
				return nil, err
			}
			arg.Bindings = append(arg.Bindings, *b)
		default:
			if err := r.skip(wire); err != nil {
				return nil, err
			}
		}
	}
	return arg, nil
}

func decodeBinding(data []byte) (*binding, error) {
	b := &binding{}
	r := newProtoReader(data)
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		switch field {
		case 1: // name: string
			b.Name, err = r.readString()
			if err != nil {
				return nil, err
			}
		case 2: // value: Value
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			v, err := decodeValue(raw)
			if err != nil {
				return nil, err
			}
			b.Value = v
		default:
			if err := r.skip(wire); err != nil {
				return nil, err
			}
		}
	}
	return b, nil
}

// decodeMapEntry decodes a protobuf map entry (key=1 string, value=2 submessage).
func decodeMapEntry[T any](data []byte, decodeFn func([]byte) (T, error)) (string, T, error) {
	var key string
	var val T
	var valRaw []byte
	r := newProtoReader(data)
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			var zero T
			return "", zero, err
		}
		switch field {
		case 1: // key: string
			key, err = r.readString()
			if err != nil {
				var zero T
				return "", zero, err
			}
		case 2: // value: submessage
			valRaw, err = r.readBytes()
			if err != nil {
				var zero T
				return "", zero, err
			}
		default:
			if err := r.skip(wire); err != nil {
				var zero T
				return "", zero, err
			}
		}
	}
	if valRaw != nil {
		var err error
		val, err = decodeFn(valRaw)
		if err != nil {
			var zero T
			return "", zero, err
		}
	}
	return key, val, nil
}
