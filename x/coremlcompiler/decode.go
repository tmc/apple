package coremlcompiler

import (
	"fmt"
	"math"
)

// decodeModel reads a CoreML Model from protobuf bytes.
func decodeModel(data []byte) (*Model, error) {
	m := &Model{}
	r := newProtoReader(data)
	for !r.done() {
		start := r.pos
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
			raw, err := r.skipUnknown(start, wire)
			if err != nil {
				return nil, fmt.Errorf("decode model: skip field %d: %w", field, err)
			}
			m.unknown = append(m.unknown, raw...)
		}
	}
	return m, nil
}

func decodeModelDescription(data []byte) (*ModelDescription, error) {
	d := &ModelDescription{}
	r := newProtoReader(data)
	for !r.done() {
		start := r.pos
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		switch field {
		case 20: // functions: repeated FunctionDescription
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			fn, err := decodeFunctionDescription(raw)
			if err != nil {
				return nil, err
			}
			d.Functions = append(d.Functions, *fn)
		case 21: // defaultFunctionName: string
			d.DefaultFunctionName, err = r.readString()
			if err != nil {
				return nil, err
			}
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
		case 100: // metadata: Metadata
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			md, err := decodeModelMetadata(raw)
			if err != nil {
				return nil, err
			}
			d.Metadata = md
		default:
			raw, err := r.skipUnknown(start, wire)
			if err != nil {
				return nil, err
			}
			d.unknown = append(d.unknown, raw...)
		}
	}
	return d, nil
}

func decodeModelMetadata(data []byte) (*ModelMetadata, error) {
	md := &ModelMetadata{}
	r := newProtoReader(data)
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		switch field {
		case 1: // shortDescription: string
			md.ShortDescription, err = r.readString()
			if err != nil {
				return nil, err
			}
		case 2: // versionString: string
			md.VersionString, err = r.readString()
			if err != nil {
				return nil, err
			}
		case 3: // author: string
			md.Author, err = r.readString()
			if err != nil {
				return nil, err
			}
		case 4: // license: string
			md.License, err = r.readString()
			if err != nil {
				return nil, err
			}
		case 100: // userDefined: map<string, string>
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			key, val, err := decodeMapEntry(raw, func(b []byte) (string, error) { return string(b), nil })
			if err != nil {
				return nil, err
			}
			if md.UserDefined == nil {
				md.UserDefined = make(map[string]string)
			}
			md.UserDefined[key] = val
		default:
			if err := r.skip(wire); err != nil {
				return nil, err
			}
		}
	}
	return md, nil
}

func decodeFunctionDescription(data []byte) (*FunctionDescription, error) {
	fn := &FunctionDescription{}
	r := newProtoReader(data)
	for !r.done() {
		start := r.pos
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		switch field {
		case 1: // name: string
			fn.Name, err = r.readString()
			if err != nil {
				return nil, err
			}
		case 2, 3, 6: // input, output, state: repeated FeatureDescription
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			fd, err := decodeFeatureDescription(raw)
			if err != nil {
				return nil, err
			}
			switch field {
			case 2:
				fn.Inputs = append(fn.Inputs, *fd)
			case 3:
				fn.Outputs = append(fn.Outputs, *fd)
			case 6:
				fn.States = append(fn.States, *fd)
			}
		default:
			raw, err := r.skipUnknown(start, wire)
			if err != nil {
				return nil, err
			}
			fn.unknown = append(fn.unknown, raw...)
		}
	}
	return fn, nil
}

func decodeFeatureDescription(data []byte) (*FeatureDescription, error) {
	fd := &FeatureDescription{}
	r := newProtoReader(data)
	for !r.done() {
		start := r.pos
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
		case 3: // type: FeatureType
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			ft, err := decodeFeatureType(raw)
			if err != nil {
				return nil, err
			}
			fd.Type = ft
		default:
			raw, err := r.skipUnknown(start, wire)
			if err != nil {
				return nil, err
			}
			fd.unknown = append(fd.unknown, raw...)
		}
	}
	return fd, nil
}

func decodeFeatureType(data []byte) (*FeatureType, error) {
	ft := &FeatureType{}
	r := newProtoReader(data)
	for !r.done() {
		start := r.pos
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		switch field {
		case 1: // int64Type
			if err := r.skip(wire); err != nil {
				return nil, err
			}
			ft.Int64Type = true
		case 2: // doubleType
			if err := r.skip(wire); err != nil {
				return nil, err
			}
			ft.DoubleType = true
		case 3: // stringType
			if err := r.skip(wire); err != nil {
				return nil, err
			}
			ft.StringType = true
		case 4: // imageType
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			img, err := decodeImageFeatureType(raw)
			if err != nil {
				return nil, err
			}
			ft.ImageType = img
		case 5: // multiArrayType
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			mat, err := decodeArrayFeatureType(raw)
			if err != nil {
				return nil, err
			}
			ft.MultiArrayType = mat
		case 6: // dictionaryType
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			dict, err := decodeDictionaryFeatureType(raw)
			if err != nil {
				return nil, err
			}
			ft.DictionaryType = dict
		case 7: // sequenceType
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			seq, err := decodeSequenceFeatureType(raw)
			if err != nil {
				return nil, err
			}
			ft.SequenceType = seq
		case 8: // stateFeatureType (wraps ArrayFeatureType at field 1)
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			mat, err := decodeStateFeatureType(raw)
			if err != nil {
				return nil, err
			}
			ft.StateArrayType = mat
		case 1000: // isOptional
			v, err := r.readVarint()
			if err != nil {
				return nil, err
			}
			ft.IsOptional = v != 0
		default:
			raw, err := r.skipUnknown(start, wire)
			if err != nil {
				return nil, err
			}
			ft.unknown = append(ft.unknown, raw...)
		}
	}
	return ft, nil
}

func decodeImageFeatureType(data []byte) (*ImageFeatureType, error) {
	img := &ImageFeatureType{}
	r := newProtoReader(data)
	for !r.done() {
		start := r.pos
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		switch field {
		case 1: // width
			v, err := r.readVarint()
			if err != nil {
				return nil, err
			}
			img.Width = int64(v)
		case 2: // height
			v, err := r.readVarint()
			if err != nil {
				return nil, err
			}
			img.Height = int64(v)
		case 3: // colorSpace
			v, err := r.readVarint()
			if err != nil {
				return nil, err
			}
			img.ColorSpace = ColorSpace(v)
		default:
			raw, err := r.skipUnknown(start, wire)
			if err != nil {
				return nil, err
			}
			img.unknown = append(img.unknown, raw...)
		}
	}
	return img, nil
}

func decodeDictionaryFeatureType(data []byte) (*DictionaryFeatureType, error) {
	dict := &DictionaryFeatureType{}
	r := newProtoReader(data)
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		switch field {
		case 1: // int64KeyType
			if err := r.skip(wire); err != nil {
				return nil, err
			}
			dict.KeyType = "int64"
		case 2: // stringKeyType
			if err := r.skip(wire); err != nil {
				return nil, err
			}
			dict.KeyType = "string"
		default:
			if err := r.skip(wire); err != nil {
				return nil, err
			}
		}
	}
	return dict, nil
}

func decodeSequenceFeatureType(data []byte) (*SequenceFeatureType, error) {
	seq := &SequenceFeatureType{}
	r := newProtoReader(data)
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		switch field {
		// The oneof members are empty messages selected by field number.
		case 1: // int64Type
			if err := r.skip(wire); err != nil {
				return nil, err
			}
			seq.ElementType = &FeatureType{Int64Type: true}
		case 3: // stringType
			if err := r.skip(wire); err != nil {
				return nil, err
			}
			seq.ElementType = &FeatureType{StringType: true}
		default:
			if err := r.skip(wire); err != nil {
				return nil, err
			}
		}
	}
	return seq, nil
}

func decodeStateFeatureType(data []byte) (*ArrayFeatureType, error) {
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
			return decodeArrayFeatureType(raw)
		}
		if err := r.skip(wire); err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func decodeArrayFeatureType(data []byte) (*ArrayFeatureType, error) {
	aft := &ArrayFeatureType{}
	r := newProtoReader(data)
	for !r.done() {
		start := r.pos
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		switch field {
		case 1: // shape: repeated int64 (packed or individual)
			if wire == wireBytes {
				raw, err := r.readBytes()
				if err != nil {
					return nil, err
				}
				pr := newProtoReader(raw)
				for !pr.done() {
					v, err := pr.readVarint()
					if err != nil {
						return nil, err
					}
					aft.Shape = append(aft.Shape, int64(v))
				}
			} else {
				v, err := r.readVarint()
				if err != nil {
					return nil, err
				}
				aft.Shape = append(aft.Shape, int64(v))
			}
		case 2: // dataType: ArrayDataType
			v, err := r.readVarint()
			if err != nil {
				return nil, err
			}
			aft.DataType = ArrayDataType(v)
		default:
			raw, err := r.skipUnknown(start, wire)
			if err != nil {
				return nil, err
			}
			aft.unknown = append(aft.unknown, raw...)
		}
	}
	return aft, nil
}

func decodeProgram(data []byte) (*Program, error) {
	p := &Program{
		Functions:  make(map[string]*Function),
		Attributes: make(map[string]*Value),
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
				return nil, fmt.Errorf("decode Program Function: %w", err)
			}
			p.Functions[key] = val
		case 4: // attributes: map<string, Value>
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			key, val, err := decodeMapEntry(raw, decodeValue)
			if err != nil {
				return nil, fmt.Errorf("decode Program attribute: %w", err)
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

func decodeFunction(data []byte) (*Function, error) {
	f := &Function{
		BlockSpecializations: make(map[string]*Block),
		Attributes:           make(map[string]*Value),
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

func decodeBlock(data []byte) (*Block, error) {
	b := &Block{}
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

func decodeOperation(data []byte) (*Operation, error) {
	op := &Operation{
		Inputs:     make(map[string]*Argument),
		Attributes: make(map[string]*Value),
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

func decodeNamedValueType(data []byte) (*NamedValueType, error) {
	nvt := &NamedValueType{}
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

func decodeValueType(data []byte) (*ValueType, error) {
	vt := &ValueType{}
	r := newProtoReader(data)
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		switch field {
		case 1: // TensorType: TensorType
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			tt, err := decodeTensorType(raw)
			if err != nil {
				return nil, err
			}
			vt.TensorType = tt
		case 2: // ListType: ListType
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			lt, err := decodeListType(raw)
			if err != nil {
				return nil, err
			}
			vt.ListType = lt
		case 3: // TupleType: TupleType
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			tup, err := decodeTupleType(raw)
			if err != nil {
				return nil, err
			}
			vt.TupleType = tup
		case 4: // DictionaryType: DictionaryType
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			dt, err := decodeDictionaryType(raw)
			if err != nil {
				return nil, err
			}
			vt.DictionaryType = dt
		case 5: // StateType: StateType
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

func decodeListType(data []byte) (*ListType, error) {
	lt := &ListType{}
	r := newProtoReader(data)
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		switch field {
		case 1: // elementType: ValueType
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			et, err := decodeValueType(raw)
			if err != nil {
				return nil, err
			}
			lt.ElementType = et
		case 2: // length: Dimension
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			d, err := decodeDimension(raw)
			if err != nil {
				return nil, err
			}
			lt.Length = int64(d.Constant)
		default:
			if err := r.skip(wire); err != nil {
				return nil, err
			}
		}
	}
	return lt, nil
}

func decodeTupleType(data []byte) (*TupleType, error) {
	tt := &TupleType{}
	r := newProtoReader(data)
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		switch field {
		case 1: // types: repeated ValueType
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			vt, err := decodeValueType(raw)
			if err != nil {
				return nil, err
			}
			tt.Types = append(tt.Types, vt)
		default:
			if err := r.skip(wire); err != nil {
				return nil, err
			}
		}
	}
	return tt, nil
}

func decodeDictionaryType(data []byte) (*DictionaryType, error) {
	dt := &DictionaryType{}
	r := newProtoReader(data)
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		switch field {
		case 1: // keyType: ValueType
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			vt, err := decodeValueType(raw)
			if err != nil {
				return nil, err
			}
			dt.KeyType = vt
		case 2: // valueType: ValueType
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			vt, err := decodeValueType(raw)
			if err != nil {
				return nil, err
			}
			dt.ValueType = vt
		default:
			if err := r.skip(wire); err != nil {
				return nil, err
			}
		}
	}
	return dt, nil
}

func decodeTensorType(data []byte) (*TensorType, error) {
	tt := &TensorType{}
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

func decodeStateType(data []byte) (*StateType, error) {
	st := &StateType{}
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

func decodeDimension(data []byte) (*Dimension, error) {
	d := &Dimension{}
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
				raw, err := r.readBytes()
				if err != nil {
					return nil, err
				}
				ur := newProtoReader(raw)
				for !ur.done() {
					uf, uw, err := ur.readTag()
					if err != nil {
						return nil, err
					}
					if uf == 1 { // variadic: bool
						v, err := ur.readVarint()
						if err != nil {
							return nil, err
						}
						d.Variadic = v != 0
					} else {
						if err := ur.skip(uw); err != nil {
							return nil, err
						}
					}
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

func decodeValue(data []byte) (*Value, error) {
	v := &Value{}
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
		case 3: // ImmediateValue: ImmediateValue
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			iv, err := decodeImmediateValue(raw)
			if err != nil {
				return nil, err
			}
			v.Immediate = iv
		case 5: // BlobFileValue: BlobFileValue
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

func decodeImmediateValue(data []byte) (*ImmediateValue, error) {
	iv := &ImmediateValue{}
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
		case 2: // tuple: TupleValue
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			tup, err := decodeTupleValue(raw)
			if err != nil {
				return nil, err
			}
			iv.Tuple = tup
		case 3: // list: ListValue
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			lv, err := decodeListValue(raw)
			if err != nil {
				return nil, err
			}
			iv.List = lv
		case 4: // dictionary: DictionaryValue
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			dv, err := decodeDictionaryValue(raw)
			if err != nil {
				return nil, err
			}
			iv.Dictionary = dv
		default:
			if err := r.skip(wire); err != nil {
				return nil, err
			}
		}
	}
	return iv, nil
}

func decodeTupleValue(data []byte) (*TupleValue, error) {
	tv := &TupleValue{}
	r := newProtoReader(data)
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		switch field {
		case 1: // values: repeated Value
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			v, err := decodeValue(raw)
			if err != nil {
				return nil, err
			}
			tv.Values = append(tv.Values, v)
		default:
			if err := r.skip(wire); err != nil {
				return nil, err
			}
		}
	}
	return tv, nil
}

func decodeListValue(data []byte) (*ListValue, error) {
	lv := &ListValue{}
	r := newProtoReader(data)
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		switch field {
		case 1: // values: repeated Value
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			v, err := decodeValue(raw)
			if err != nil {
				return nil, err
			}
			lv.Values = append(lv.Values, v)
		default:
			if err := r.skip(wire); err != nil {
				return nil, err
			}
		}
	}
	return lv, nil
}

func decodeDictionaryValue(data []byte) (*DictionaryValue, error) {
	dv := &DictionaryValue{}
	r := newProtoReader(data)
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		switch field {
		case 1: // values: map<Value, Value> represented as entries
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			entry, err := decodeDictionaryMapEntry(raw)
			if err != nil {
				return nil, err
			}
			dv.Entries = append(dv.Entries, entry)
		default:
			if err := r.skip(wire); err != nil {
				return nil, err
			}
		}
	}
	return dv, nil
}

func decodeDictionaryMapEntry(data []byte) (DictionaryMapEntry, error) {
	var entry DictionaryMapEntry
	r := newProtoReader(data)
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return entry, err
		}
		switch field {
		case 1:
			raw, err := r.readBytes()
			if err != nil {
				return entry, err
			}
			k, err := decodeValue(raw)
			if err != nil {
				return entry, err
			}
			entry.Key = k
		case 2:
			raw, err := r.readBytes()
			if err != nil {
				return entry, err
			}
			v, err := decodeValue(raw)
			if err != nil {
				return entry, err
			}
			entry.Value = v
		default:
			if err := r.skip(wire); err != nil {
				return entry, err
			}
		}
	}
	return entry, nil
}

func decodeTensorValue(data []byte) (*TensorValue, error) {
	tv := &TensorValue{}
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

// The repeated scalar fields inside RepeatedFloats/Ints/... are declared
// packed, but a conforming proto3 parser must also accept the unpacked form
// (one tag per element) and a packed field split across several chunks, which
// are concatenated. These helpers accept all three forms; the caller only
// invokes them when the submessage is present, so an empty-but-non-nil result
// distinguishes a present, empty tensor from an absent field.

// decodeRepeatedFloats decodes a RepeatedFloats submessage (field 1: float).
func decodeRepeatedFloats(data []byte) ([]float32, error) {
	r := newProtoReader(data)
	out := []float32{}
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		if field == 1 && wire == wireBytes {
			vals, err := r.readPackedFloat32()
			if err != nil {
				return nil, err
			}
			out = append(out, vals...)
			continue
		}
		if field == 1 && wire == wireFixed32 {
			v, err := r.readFixed32()
			if err != nil {
				return nil, err
			}
			out = append(out, math.Float32frombits(v))
			continue
		}
		if err := r.skip(wire); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// decodeRepeatedDoubles decodes a RepeatedDoubles submessage.
func decodeRepeatedDoubles(data []byte) ([]float64, error) {
	r := newProtoReader(data)
	out := []float64{}
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		if field == 1 && wire == wireBytes {
			vals, err := r.readPackedFloat64()
			if err != nil {
				return nil, err
			}
			out = append(out, vals...)
			continue
		}
		if field == 1 && wire == wireFixed64 {
			v, err := r.readFixed64()
			if err != nil {
				return nil, err
			}
			out = append(out, math.Float64frombits(v))
			continue
		}
		if err := r.skip(wire); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// decodeRepeatedInts decodes a RepeatedInts submessage (field 1: int32 varint).
func decodeRepeatedInts(data []byte) ([]int32, error) {
	r := newProtoReader(data)
	out := []int32{}
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		if field == 1 && wire == wireBytes {
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			vals, err := readPackedInt32(raw)
			if err != nil {
				return nil, err
			}
			out = append(out, vals...)
			continue
		}
		if field == 1 && wire == wireVarint {
			v, err := r.readVarint()
			if err != nil {
				return nil, err
			}
			out = append(out, int32(v))
			continue
		}
		if err := r.skip(wire); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// decodeRepeatedLongs decodes a RepeatedLongInts submessage.
func decodeRepeatedLongs(data []byte) ([]int64, error) {
	r := newProtoReader(data)
	out := []int64{}
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		if field == 1 && wire == wireBytes {
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			vals, err := readPackedInt64(raw)
			if err != nil {
				return nil, err
			}
			out = append(out, vals...)
			continue
		}
		if field == 1 && wire == wireVarint {
			v, err := r.readVarint()
			if err != nil {
				return nil, err
			}
			out = append(out, int64(v))
			continue
		}
		if err := r.skip(wire); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// decodeRepeatedBools decodes a RepeatedBools submessage (field 1: bool varint).
func decodeRepeatedBools(data []byte) ([]bool, error) {
	r := newProtoReader(data)
	out := []bool{}
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			return nil, err
		}
		if field == 1 && wire == wireBytes {
			raw, err := r.readBytes()
			if err != nil {
				return nil, err
			}
			pr := newProtoReader(raw)
			for !pr.done() {
				v, err := pr.readVarint()
				if err != nil {
					return nil, err
				}
				out = append(out, v != 0)
			}
			continue
		}
		if field == 1 && wire == wireVarint {
			v, err := r.readVarint()
			if err != nil {
				return nil, err
			}
			out = append(out, v != 0)
			continue
		}
		if err := r.skip(wire); err != nil {
			return nil, err
		}
	}
	return out, nil
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

func decodeBlobFileValue(data []byte) (*BlobFileValue, error) {
	bf := &BlobFileValue{}
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

func decodeArgument(data []byte) (*Argument, error) {
	arg := &Argument{}
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

func decodeBinding(data []byte) (*Binding, error) {
	b := &Binding{}
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
