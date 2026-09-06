package jaccl

import "fmt"

// DType is an element representation used by AllReduce. Its values match the
// standalone JACCL Dtype enum so an external parity fixture can use the same
// type table in both implementations.
type DType uint8

const (
	// Bool is a boolean element.
	Bool DType = iota
	// Int8 is a signed 8-bit integer element.
	Int8
	// Int16 is a signed 16-bit integer element.
	Int16
	// Int32 is a signed 32-bit integer element.
	Int32
	// Int64 is a signed 64-bit integer element.
	Int64
	// UInt8 is an unsigned 8-bit integer element.
	UInt8
	// UInt16 is an unsigned 16-bit integer element.
	UInt16
	// UInt32 is an unsigned 32-bit integer element.
	UInt32
	// UInt64 is an unsigned 64-bit integer element.
	UInt64
	// Float16 is an IEEE 754 binary16 element.
	Float16
	// BFloat16 is a bfloat16 element.
	BFloat16
	// Float32 is an IEEE 754 binary32 element.
	Float32
	// Float64 is an IEEE 754 binary64 element.
	Float64
	// Complex64 is a pair of IEEE 754 binary32 values.
	Complex64
)

// Uint8 is kept as the spelling used by the first Go API.
const Uint8 = UInt8

// Size reports the number of bytes in one value of d.
func (d DType) Size() (int, error) {
	switch d {
	case Bool, Int8, UInt8:
		return 1, nil
	case Int16, UInt16, Float16, BFloat16:
		return 2, nil
	case Int32, UInt32, Float32:
		return 4, nil
	case Int64, UInt64, Float64, Complex64:
		return 8, nil
	default:
		return 0, fmt.Errorf("dtype %d: %w", d, ErrProtocol)
	}
}

// ReduceOp selects the reduction performed by AllReduce.
type ReduceOp uint8

const (
	// Sum adds corresponding elements.
	Sum ReduceOp = iota + 1
	// Min selects the smaller corresponding element.
	Min
	// Max selects the larger corresponding element.
	Max
)

func (op ReduceOp) valid() bool {
	return op == Sum || op == Min || op == Max
}
