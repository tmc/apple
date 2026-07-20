// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"unsafe"
)

// C struct types

// AnalyticsData
type AnalyticsData struct {
	Field1 uint
	Field2 uint
	Field3 unsafe.Pointer
}

// AnalyticsGroupInfo
type AnalyticsGroupInfo struct {
	Field1 uint
	Field2 uint64
	Field3 uint
	Field4 uint64
}

// AnalyticsLayerInfo
type AnalyticsLayerInfo struct {
	Field1 unsafe.Pointer
	Field2 unsafe.Pointer
	Field3 float32
}

// AnalyticsProcedureInfo
type AnalyticsProcedureInfo struct {
	Field1 uint
	Field2 uint
	Field3 uint
	Field4 uint
	Field5 uint
	Field6 uint64
	Field7 uint
	Field8 uint64
}

// AnalyticsTaskInfo
type AnalyticsTaskInfo struct {
	Field1 uint
	Field2 uint64
}

// CGColorSpace
type CGColorSpace struct {
}

// CGImage
type CGImage struct {
}

// CVBuffer
type CVBuffer struct {
}

// IOSurface
type IOSurface struct {
}

// ConvolutionUniforms
type ConvolutionUniforms struct {
	Field1  int
	Field2  int
	Field3  int
	Field4  int
	Field5  int
	Field6  int
	Field7  int
	Field8  int
	Field9  int
	Field10 float32
	Field11 int
	Field12 int
	Field13 int
	Field14 int
	Field15 int
	Field16 int
	Field17 float32
	Field18 float32
	Field19 unsafe.Pointer
	Field20 int16
	Field21 int16
	Field22 int16
	Field23 int16
	Field24 uint16
	Field25 uint16
	Field26 int
	Field27 int
	Field28 int
	Field29 int
	Field30 int
	Field31 uint16
	Field32 uint16
	Field33 uint16
	Field34 uint16
	Field35 uint16
	Field36 int16
	Field37 int
	Field38 int
	Field39 int
	Field40 int
	Field41 int
	Field42 int16
	Field43 int
	Field44 bool
}

// Convolution_uniforms is a type alias for ConvolutionUniforms for use in objc.Send[T] calls.
type Convolution_uniforms = ConvolutionUniforms

// FloatBuffer
type FloatBuffer struct {
	Field1 []float32
	Field2 uint64
	Field3 bool
}

// Float_buffer_t is a type alias for FloatBuffer for use in objc.Send[T] calls.
type Float_buffer_t = FloatBuffer

// InnerProductUniforms
type InnerProductUniforms struct {
	Field1  uint
	Field2  uint
	Field3  int
	Field4  int
	Field5  int
	Field6  float32
	Field7  float32
	Field8  int
	Field9  int
	Field10 int
	Field11 bool
	Field12 int
	Field13 int
	Field14 int
	Field15 float32
	Field16 float32
	Field17 uint
	Field18 uint
	Field19 uint
	Field20 uint
	Field21 uint
	Field22 uint
	Field23 uint
	Field24 uint
	Field25 uint
	Field26 int
	Field27 int
	Field28 int
	Field29 int
	Field30 int
	Field31 int
	Field32 int
	Field33 int
}

// Inner_product_uniforms is a type alias for InnerProductUniforms for use in objc.Send[T] calls.
type Inner_product_uniforms = InnerProductUniforms

// MxnetToolsImageHeaderT
type MxnetToolsImageHeaderT struct {
	Field1 uint
	Field2 float32
	Field3 MxnetToolsImageIDT
}

// MxnetTools_imageHeader_t_ is a type alias for MxnetToolsImageHeaderT for use in objc.Send[T] calls.
type MxnetTools_imageHeader_t_ = MxnetToolsImageHeaderT

// MxnetToolsImageIDT
type MxnetToolsImageIDT struct {
	Field1 unsafe.Pointer
}

// MxnetTools_imageID_t_ is a type alias for MxnetToolsImageIDT for use in objc.Send[T] calls.
type MxnetTools_imageID_t_ = MxnetToolsImageIDT

// MxnetToolsRecordHeaderT
type MxnetToolsRecordHeaderT struct {
	Field1 uint
	Field2 uint
}

// MxnetTools_recordHeader_t_ is a type alias for MxnetToolsRecordHeaderT for use in objc.Send[T] calls.
type MxnetTools_recordHeader_t_ = MxnetToolsRecordHeaderT

// NetStridesConfiguration
type NetStridesConfiguration struct {
	Field1 int
	Field2 unsafe.Pointer
	Field3 unsafe.Pointer
	Field4 unsafe.Pointer
}

// Net_strides_configuration is a type alias for NetStridesConfiguration for use in objc.Send[T] calls.
type Net_strides_configuration = NetStridesConfiguration
