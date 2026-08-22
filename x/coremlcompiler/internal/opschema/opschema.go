// Package opschema holds the per-op input specifications MIL ops are declared
// with in coremltools.
//
// Every op class in coremltools/converters/mil/mil/ops/defs declares an
// InputSpec (input_type.py:32-53): an ordered map from input name to an
// _InputType carrying a const and an optional flag (input_type.py:145-160) and,
// for tensors, a type domain (input_type.py:238-282). The spec is what
// coremltools checks a program against before it will serialize it. The table
// in table.gen.go is a dump of those specs for the ops we emit; see
// gen/dumpops.py.
package opschema

//go:generate python3 gen/dumpops.py
//go:generate go run ./gen/gentable.go

// Kind is the _InputType subclass an input was declared with.
type Kind uint8

const (
	KindTensor Kind = iota
	KindList
	KindListOrTensorOrDict
	KindTuple
	KindState
	KindPyFunc
	KindInternal
)

// DataType is a MIL builtin type name, as types.builtin_to_string renders it.
type DataType string

// Param is one declared input of an op.
type Param struct {
	Name     string
	Kind     Kind
	Const    bool // must be const at compile time
	Optional bool
	DomainID string     // type domain id, "" if the spec listed types directly
	Domain   []DataType // allowed types; empty iff DomainID != ""
}

// Op is the input specification of one op at one opset.
type Op struct {
	Type    string
	Params  []Param               // declaration order
	Domains map[string][]DataType // type_domains
}

// Param returns the declared input named name.
func (op *Op) Param(name string) (*Param, bool) {
	for i := range op.Params {
		if op.Params[i].Name == name {
			return &op.Params[i], true
		}
	}
	return nil, false
}

// Known reports whether schemas are recorded for opset, named as MIL text names
// it ("ios15" and up). Callers skip per-op checking for unknown opsets rather
// than reporting an error, so a test asserts the opsets we emit are known.
func Known(opset string) bool {
	_, ok := opsets[opset]
	return ok
}

// Registered reports whether the table records opType for any opset. It is
// false both for an op that does not exist and for one the dump does not
// cover, so it cannot be used to reject an op outright.
func Registered(opType string) bool {
	return registered[opType]
}

// Lookup returns the schema for opType at opset. ok is false if either the
// opset or the op is unknown.
func Lookup(opset, opType string) (*Op, bool) {
	ops, ok := opsets[opset]
	if !ok {
		return nil, false
	}
	op, ok := ops[opType]
	return op, ok
}
