// Code generated from Apple documentation. DO NOT EDIT.

package spritekit

// SKActionTimingFunction is the signature for the custom timing block.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKActionTimingFunction
type SKActionTimingFunction = func(float32) float32

// SKFieldForceEvaluator is the definition for a custom block that processes a single physics body’s interaction with the field.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKFieldForceEvaluator
type SKFieldForceEvaluator = func(float32, float32, float32, float32, float64) float32

// Vector_float3 is a floating point vector type used to perform physics calculations.
//
// See: https://developer.apple.com/documentation/SpriteKit/vector_float3
type Vector_float3 = [3]float32

// VectorFloat3 is a Go-name alias for Vector_float3.
type VectorFloat3 = Vector_float3
