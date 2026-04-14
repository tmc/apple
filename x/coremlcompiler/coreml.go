package coremlcompiler

import (
	"unsafe"

	"github.com/tmc/apple/coreml"
)

// PredictInput is a named tensor input for CoreML prediction.
type PredictInput struct {
	Name    string
	Data    unsafe.Pointer
	Shape   []int
	Strides []int
	DType   coreml.MLMultiArrayDataType
}

// PredictOutput holds the result of a CoreML prediction.
type PredictOutput struct {
	Data  unsafe.Pointer // owned by CoreML, valid until model is closed
	Shape []int
	DType coreml.MLMultiArrayDataType
}
