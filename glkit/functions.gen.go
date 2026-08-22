// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"fmt"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

type unavailableSymbolError struct {
	symbol     string
	introduced string
	cause      error
}

func (e *unavailableSymbolError) Error() string {
	if e == nil {
		return ""
	}
	if e.introduced != "" {
		return fmt.Sprintf("GLKit: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("GLKit: symbol %s unavailable on this system", e.symbol)
}

func (e *unavailableSymbolError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.cause
}

func missingSymbolError(name, introduced string, cause error) error {
	return &unavailableSymbolError{
		symbol:     name,
		introduced: introduced,
		cause:      cause,
	}
}

func symbolCallError(name, introduced string, err error) error {
	if err != nil {
		return err
	}
	if frameworkHandle == 0 {
		return fmt.Errorf("GLKit: symbol %s unavailable because the framework could not be loaded", name)
	}
	return missingSymbolError(name, introduced, nil)
}

// registerFunc resolves a framework symbol and registers it as a Go function.
func registerFunc(fptr any, errDst *error, handle uintptr, name, introduced string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil || sym == 0 {
		*errDst = missingSymbolError(name, introduced, err)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			*errDst = fmt.Errorf("GLKit: register symbol %s: %v", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
	*errDst = nil
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, errDst *error, handle uintptr, name, introduced string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil || sym == 0 {
		*errDst = missingSymbolError(name, introduced, err)
		return
	}
	*dst = sym
	*errDst = nil
}

var _gLKMathProject func(object GLKVector3, model GLKMatrix4, projection GLKMatrix4, viewport *int32) GLKVector3
var _gLKMathProjectErr error

func tryGLKMathProject(object GLKVector3, model GLKMatrix4, projection GLKMatrix4, viewport *int32) (GLKVector3, error) {
	if _gLKMathProject == nil {
		return *new(GLKVector3), symbolCallError("GLKMathProject", "10.8", _gLKMathProjectErr)
	}
	return _gLKMathProject(object, model, projection, viewport), nil
}

// GLKMathProject projects a point in object space into the window coordinate system.
//
// See: https://developer.apple.com/documentation/GLKit/GLKMathProject(_:_:_:_:)
func GLKMathProject(object GLKVector3, model GLKMatrix4, projection GLKMatrix4, viewport *int32) GLKVector3 {
	result, callErr := tryGLKMathProject(object, model, projection, viewport)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gLKMathUnproject func(window GLKVector3, model GLKMatrix4, projection GLKMatrix4, viewport *int32, success *bool) GLKVector3
var _gLKMathUnprojectErr error

func tryGLKMathUnproject(window GLKVector3, model GLKMatrix4, projection GLKMatrix4, viewport *int32, success *bool) (GLKVector3, error) {
	if _gLKMathUnproject == nil {
		return *new(GLKVector3), symbolCallError("GLKMathUnproject", "10.8", _gLKMathUnprojectErr)
	}
	return _gLKMathUnproject(window, model, projection, viewport, success), nil
}

// GLKMathUnproject projects a point in view space into object space.
//
// See: https://developer.apple.com/documentation/GLKit/GLKMathUnproject(_:_:_:_:_:)
func GLKMathUnproject(window GLKVector3, model GLKMatrix4, projection GLKMatrix4, viewport *int32, success *bool) GLKVector3 {
	result, callErr := tryGLKMathUnproject(window, model, projection, viewport, success)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gLKMatrix3Invert func(matrix GLKMatrix3, isInvertible *bool) GLKMatrix3
var _gLKMatrix3InvertErr error

func tryGLKMatrix3Invert(matrix GLKMatrix3, isInvertible *bool) (GLKMatrix3, error) {
	if _gLKMatrix3Invert == nil {
		return *new(GLKMatrix3), symbolCallError("GLKMatrix3Invert", "10.8", _gLKMatrix3InvertErr)
	}
	return _gLKMatrix3Invert(matrix, isInvertible), nil
}

// GLKMatrix3Invert returns the inverse of a matrix.
//
// See: https://developer.apple.com/documentation/GLKit/GLKMatrix3Invert(_:_:)
func GLKMatrix3Invert(matrix GLKMatrix3, isInvertible *bool) GLKMatrix3 {
	result, callErr := tryGLKMatrix3Invert(matrix, isInvertible)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gLKMatrix3InvertAndTranspose func(matrix GLKMatrix3, isInvertible *bool) GLKMatrix3
var _gLKMatrix3InvertAndTransposeErr error

func tryGLKMatrix3InvertAndTranspose(matrix GLKMatrix3, isInvertible *bool) (GLKMatrix3, error) {
	if _gLKMatrix3InvertAndTranspose == nil {
		return *new(GLKMatrix3), symbolCallError("GLKMatrix3InvertAndTranspose", "10.8", _gLKMatrix3InvertAndTransposeErr)
	}
	return _gLKMatrix3InvertAndTranspose(matrix, isInvertible), nil
}

// GLKMatrix3InvertAndTranspose returns the inverse transpose of a matrix.
//
// See: https://developer.apple.com/documentation/GLKit/GLKMatrix3InvertAndTranspose(_:_:)
func GLKMatrix3InvertAndTranspose(matrix GLKMatrix3, isInvertible *bool) GLKMatrix3 {
	result, callErr := tryGLKMatrix3InvertAndTranspose(matrix, isInvertible)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gLKMatrix4Invert func(matrix GLKMatrix4, isInvertible *bool) GLKMatrix4
var _gLKMatrix4InvertErr error

func tryGLKMatrix4Invert(matrix GLKMatrix4, isInvertible *bool) (GLKMatrix4, error) {
	if _gLKMatrix4Invert == nil {
		return *new(GLKMatrix4), symbolCallError("GLKMatrix4Invert", "10.8", _gLKMatrix4InvertErr)
	}
	return _gLKMatrix4Invert(matrix, isInvertible), nil
}

// GLKMatrix4Invert returns the inverse of a matrix.
//
// See: https://developer.apple.com/documentation/GLKit/GLKMatrix4Invert(_:_:)
func GLKMatrix4Invert(matrix GLKMatrix4, isInvertible *bool) GLKMatrix4 {
	result, callErr := tryGLKMatrix4Invert(matrix, isInvertible)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gLKMatrix4InvertAndTranspose func(matrix GLKMatrix4, isInvertible *bool) GLKMatrix4
var _gLKMatrix4InvertAndTransposeErr error

func tryGLKMatrix4InvertAndTranspose(matrix GLKMatrix4, isInvertible *bool) (GLKMatrix4, error) {
	if _gLKMatrix4InvertAndTranspose == nil {
		return *new(GLKMatrix4), symbolCallError("GLKMatrix4InvertAndTranspose", "10.8", _gLKMatrix4InvertAndTransposeErr)
	}
	return _gLKMatrix4InvertAndTranspose(matrix, isInvertible), nil
}

// GLKMatrix4InvertAndTranspose returns the inverse transpose of a matrix.
//
// See: https://developer.apple.com/documentation/GLKit/GLKMatrix4InvertAndTranspose(_:_:)
func GLKMatrix4InvertAndTranspose(matrix GLKMatrix4, isInvertible *bool) GLKMatrix4 {
	result, callErr := tryGLKMatrix4InvertAndTranspose(matrix, isInvertible)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gLKMatrixStackCreate func(alloc corefoundation.CFAllocatorRef) GLKMatrixStackRef
var _gLKMatrixStackCreateErr error

func tryGLKMatrixStackCreate(alloc corefoundation.CFAllocatorRef) (GLKMatrixStackRef, error) {
	if _gLKMatrixStackCreate == nil {
		return *new(GLKMatrixStackRef), symbolCallError("GLKMatrixStackCreate", "10.8", _gLKMatrixStackCreateErr)
	}
	return _gLKMatrixStackCreate(alloc), nil
}

// GLKMatrixStackCreate allocates and returns a new matrix stack.
//
// See: https://developer.apple.com/documentation/GLKit/GLKMatrixStackCreate(_:)
func GLKMatrixStackCreate(alloc corefoundation.CFAllocatorRef) GLKMatrixStackRef {
	result, callErr := tryGLKMatrixStackCreate(alloc)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gLKMatrixStackGetMatrix2 func(stack GLKMatrixStackRef) GLKMatrix2
var _gLKMatrixStackGetMatrix2Err error

func tryGLKMatrixStackGetMatrix2(stack GLKMatrixStackRef) (GLKMatrix2, error) {
	if _gLKMatrixStackGetMatrix2 == nil {
		return *new(GLKMatrix2), symbolCallError("GLKMatrixStackGetMatrix2", "10.8", _gLKMatrixStackGetMatrix2Err)
	}
	return _gLKMatrixStackGetMatrix2(stack), nil
}

// GLKMatrixStackGetMatrix2 returns the top-left `2x2` corner of the top matrix.
//
// See: https://developer.apple.com/documentation/GLKit/GLKMatrixStackGetMatrix2(_:)
func GLKMatrixStackGetMatrix2(stack GLKMatrixStackRef) GLKMatrix2 {
	result, callErr := tryGLKMatrixStackGetMatrix2(stack)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gLKMatrixStackGetMatrix3 func(stack GLKMatrixStackRef) GLKMatrix3
var _gLKMatrixStackGetMatrix3Err error

func tryGLKMatrixStackGetMatrix3(stack GLKMatrixStackRef) (GLKMatrix3, error) {
	if _gLKMatrixStackGetMatrix3 == nil {
		return *new(GLKMatrix3), symbolCallError("GLKMatrixStackGetMatrix3", "10.8", _gLKMatrixStackGetMatrix3Err)
	}
	return _gLKMatrixStackGetMatrix3(stack), nil
}

// GLKMatrixStackGetMatrix3 returns the top-left `3x3` corner of the top matrix.
//
// See: https://developer.apple.com/documentation/GLKit/GLKMatrixStackGetMatrix3(_:)
func GLKMatrixStackGetMatrix3(stack GLKMatrixStackRef) GLKMatrix3 {
	result, callErr := tryGLKMatrixStackGetMatrix3(stack)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gLKMatrixStackGetMatrix3Inverse func(stack GLKMatrixStackRef) GLKMatrix3
var _gLKMatrixStackGetMatrix3InverseErr error

func tryGLKMatrixStackGetMatrix3Inverse(stack GLKMatrixStackRef) (GLKMatrix3, error) {
	if _gLKMatrixStackGetMatrix3Inverse == nil {
		return *new(GLKMatrix3), symbolCallError("GLKMatrixStackGetMatrix3Inverse", "10.8", _gLKMatrixStackGetMatrix3InverseErr)
	}
	return _gLKMatrixStackGetMatrix3Inverse(stack), nil
}

// GLKMatrixStackGetMatrix3Inverse fetches the top-left `3x3` corner of the top matrix and returns its inverse.
//
// See: https://developer.apple.com/documentation/GLKit/GLKMatrixStackGetMatrix3Inverse(_:)
func GLKMatrixStackGetMatrix3Inverse(stack GLKMatrixStackRef) GLKMatrix3 {
	result, callErr := tryGLKMatrixStackGetMatrix3Inverse(stack)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gLKMatrixStackGetMatrix3InverseTranspose func(stack GLKMatrixStackRef) GLKMatrix3
var _gLKMatrixStackGetMatrix3InverseTransposeErr error

func tryGLKMatrixStackGetMatrix3InverseTranspose(stack GLKMatrixStackRef) (GLKMatrix3, error) {
	if _gLKMatrixStackGetMatrix3InverseTranspose == nil {
		return *new(GLKMatrix3), symbolCallError("GLKMatrixStackGetMatrix3InverseTranspose", "10.8", _gLKMatrixStackGetMatrix3InverseTransposeErr)
	}
	return _gLKMatrixStackGetMatrix3InverseTranspose(stack), nil
}

// GLKMatrixStackGetMatrix3InverseTranspose fetches the top-left `3x3` corner of the top matrix and returns its inverse transpose.
//
// See: https://developer.apple.com/documentation/GLKit/GLKMatrixStackGetMatrix3InverseTranspose(_:)
func GLKMatrixStackGetMatrix3InverseTranspose(stack GLKMatrixStackRef) GLKMatrix3 {
	result, callErr := tryGLKMatrixStackGetMatrix3InverseTranspose(stack)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gLKMatrixStackGetMatrix4 func(stack GLKMatrixStackRef) GLKMatrix4
var _gLKMatrixStackGetMatrix4Err error

func tryGLKMatrixStackGetMatrix4(stack GLKMatrixStackRef) (GLKMatrix4, error) {
	if _gLKMatrixStackGetMatrix4 == nil {
		return *new(GLKMatrix4), symbolCallError("GLKMatrixStackGetMatrix4", "10.8", _gLKMatrixStackGetMatrix4Err)
	}
	return _gLKMatrixStackGetMatrix4(stack), nil
}

// GLKMatrixStackGetMatrix4 returns a copy of the top matrix on the stack.
//
// See: https://developer.apple.com/documentation/GLKit/GLKMatrixStackGetMatrix4(_:)
func GLKMatrixStackGetMatrix4(stack GLKMatrixStackRef) GLKMatrix4 {
	result, callErr := tryGLKMatrixStackGetMatrix4(stack)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gLKMatrixStackGetMatrix4Inverse func(stack GLKMatrixStackRef) GLKMatrix4
var _gLKMatrixStackGetMatrix4InverseErr error

func tryGLKMatrixStackGetMatrix4Inverse(stack GLKMatrixStackRef) (GLKMatrix4, error) {
	if _gLKMatrixStackGetMatrix4Inverse == nil {
		return *new(GLKMatrix4), symbolCallError("GLKMatrixStackGetMatrix4Inverse", "10.8", _gLKMatrixStackGetMatrix4InverseErr)
	}
	return _gLKMatrixStackGetMatrix4Inverse(stack), nil
}

// GLKMatrixStackGetMatrix4Inverse returns the inverse of the top matrix.
//
// See: https://developer.apple.com/documentation/GLKit/GLKMatrixStackGetMatrix4Inverse(_:)
func GLKMatrixStackGetMatrix4Inverse(stack GLKMatrixStackRef) GLKMatrix4 {
	result, callErr := tryGLKMatrixStackGetMatrix4Inverse(stack)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gLKMatrixStackGetMatrix4InverseTranspose func(stack GLKMatrixStackRef) GLKMatrix4
var _gLKMatrixStackGetMatrix4InverseTransposeErr error

func tryGLKMatrixStackGetMatrix4InverseTranspose(stack GLKMatrixStackRef) (GLKMatrix4, error) {
	if _gLKMatrixStackGetMatrix4InverseTranspose == nil {
		return *new(GLKMatrix4), symbolCallError("GLKMatrixStackGetMatrix4InverseTranspose", "10.8", _gLKMatrixStackGetMatrix4InverseTransposeErr)
	}
	return _gLKMatrixStackGetMatrix4InverseTranspose(stack), nil
}

// GLKMatrixStackGetMatrix4InverseTranspose returns the inverse transpose of the top matrix.
//
// See: https://developer.apple.com/documentation/GLKit/GLKMatrixStackGetMatrix4InverseTranspose(_:)
func GLKMatrixStackGetMatrix4InverseTranspose(stack GLKMatrixStackRef) GLKMatrix4 {
	result, callErr := tryGLKMatrixStackGetMatrix4InverseTranspose(stack)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gLKMatrixStackGetTypeID func() uint
var _gLKMatrixStackGetTypeIDErr error

func tryGLKMatrixStackGetTypeID() (uint, error) {
	if _gLKMatrixStackGetTypeID == nil {
		return 0, symbolCallError("GLKMatrixStackGetTypeID", "10.8", _gLKMatrixStackGetTypeIDErr)
	}
	return _gLKMatrixStackGetTypeID(), nil
}

// GLKMatrixStackGetTypeID returns the Core Foundation type for a matrix stack.
//
// See: https://developer.apple.com/documentation/GLKit/GLKMatrixStackGetTypeID()
func GLKMatrixStackGetTypeID() uint {
	result, callErr := tryGLKMatrixStackGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gLKMatrixStackLoadMatrix4 func(stack GLKMatrixStackRef, matrix GLKMatrix4)
var _gLKMatrixStackLoadMatrix4Err error

func tryGLKMatrixStackLoadMatrix4(stack GLKMatrixStackRef, matrix GLKMatrix4) error {
	if _gLKMatrixStackLoadMatrix4 == nil {
		return symbolCallError("GLKMatrixStackLoadMatrix4", "10.8", _gLKMatrixStackLoadMatrix4Err)
	}
	_gLKMatrixStackLoadMatrix4(stack, matrix)
	return nil
}

// GLKMatrixStackLoadMatrix4 replaces the contents of the top matrix with a new matrix.
//
// See: https://developer.apple.com/documentation/GLKit/GLKMatrixStackLoadMatrix4(_:_:)
func GLKMatrixStackLoadMatrix4(stack GLKMatrixStackRef, matrix GLKMatrix4) {
	if callErr := tryGLKMatrixStackLoadMatrix4(stack, matrix); callErr != nil {
		panic(callErr)
	}
}

var _gLKMatrixStackMultiplyMatrix4 func(stack GLKMatrixStackRef, matrix GLKMatrix4)
var _gLKMatrixStackMultiplyMatrix4Err error

func tryGLKMatrixStackMultiplyMatrix4(stack GLKMatrixStackRef, matrix GLKMatrix4) error {
	if _gLKMatrixStackMultiplyMatrix4 == nil {
		return symbolCallError("GLKMatrixStackMultiplyMatrix4", "10.8", _gLKMatrixStackMultiplyMatrix4Err)
	}
	_gLKMatrixStackMultiplyMatrix4(stack, matrix)
	return nil
}

// GLKMatrixStackMultiplyMatrix4 replaces the contents of the top matrix with a matrix calculated by multiplying the contents of the top matrix by another matrix.
//
// See: https://developer.apple.com/documentation/GLKit/GLKMatrixStackMultiplyMatrix4(_:_:)
func GLKMatrixStackMultiplyMatrix4(stack GLKMatrixStackRef, matrix GLKMatrix4) {
	if callErr := tryGLKMatrixStackMultiplyMatrix4(stack, matrix); callErr != nil {
		panic(callErr)
	}
}

var _gLKMatrixStackMultiplyMatrixStack func(stackLeft GLKMatrixStackRef, stackRight GLKMatrixStackRef)
var _gLKMatrixStackMultiplyMatrixStackErr error

func tryGLKMatrixStackMultiplyMatrixStack(stackLeft GLKMatrixStackRef, stackRight GLKMatrixStackRef) error {
	if _gLKMatrixStackMultiplyMatrixStack == nil {
		return symbolCallError("GLKMatrixStackMultiplyMatrixStack", "10.8", _gLKMatrixStackMultiplyMatrixStackErr)
	}
	_gLKMatrixStackMultiplyMatrixStack(stackLeft, stackRight)
	return nil
}

// GLKMatrixStackMultiplyMatrixStack replaces the contents of the top matrix with a matrix calculated by multiplying the contents of the top matrix by the top matrix of another matrix stack.
//
// See: https://developer.apple.com/documentation/GLKit/GLKMatrixStackMultiplyMatrixStack(_:_:)
func GLKMatrixStackMultiplyMatrixStack(stackLeft GLKMatrixStackRef, stackRight GLKMatrixStackRef) {
	if callErr := tryGLKMatrixStackMultiplyMatrixStack(stackLeft, stackRight); callErr != nil {
		panic(callErr)
	}
}

var _gLKMatrixStackPop func(stack GLKMatrixStackRef)
var _gLKMatrixStackPopErr error

func tryGLKMatrixStackPop(stack GLKMatrixStackRef) error {
	if _gLKMatrixStackPop == nil {
		return symbolCallError("GLKMatrixStackPop", "10.8", _gLKMatrixStackPopErr)
	}
	_gLKMatrixStackPop(stack)
	return nil
}

// GLKMatrixStackPop removes the topmost entry from the stack.
//
// See: https://developer.apple.com/documentation/GLKit/GLKMatrixStackPop(_:)
func GLKMatrixStackPop(stack GLKMatrixStackRef) {
	if callErr := tryGLKMatrixStackPop(stack); callErr != nil {
		panic(callErr)
	}
}

var _gLKMatrixStackPush func(stack GLKMatrixStackRef)
var _gLKMatrixStackPushErr error

func tryGLKMatrixStackPush(stack GLKMatrixStackRef) error {
	if _gLKMatrixStackPush == nil {
		return symbolCallError("GLKMatrixStackPush", "10.8", _gLKMatrixStackPushErr)
	}
	_gLKMatrixStackPush(stack)
	return nil
}

// GLKMatrixStackPush push a copy of the topmost matrix onto the top of the stack.
//
// See: https://developer.apple.com/documentation/GLKit/GLKMatrixStackPush(_:)
func GLKMatrixStackPush(stack GLKMatrixStackRef) {
	if callErr := tryGLKMatrixStackPush(stack); callErr != nil {
		panic(callErr)
	}
}

var _gLKMatrixStackRotate func(stack GLKMatrixStackRef, radians float32, x float32, y float32, z float32)
var _gLKMatrixStackRotateErr error

func tryGLKMatrixStackRotate(stack GLKMatrixStackRef, radians float32, x float32, y float32, z float32) error {
	if _gLKMatrixStackRotate == nil {
		return symbolCallError("GLKMatrixStackRotate", "10.8", _gLKMatrixStackRotateErr)
	}
	_gLKMatrixStackRotate(stack, radians, x, y, z)
	return nil
}

// GLKMatrixStackRotate replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a rotation around an arbitrary axis.
//
// See: https://developer.apple.com/documentation/GLKit/GLKMatrixStackRotate(_:_:_:_:_:)
func GLKMatrixStackRotate(stack GLKMatrixStackRef, radians float32, x float32, y float32, z float32) {
	if callErr := tryGLKMatrixStackRotate(stack, radians, x, y, z); callErr != nil {
		panic(callErr)
	}
}

var _gLKMatrixStackRotateWithVector3 func(stack GLKMatrixStackRef, radians float32, axisVector GLKVector3)
var _gLKMatrixStackRotateWithVector3Err error

func tryGLKMatrixStackRotateWithVector3(stack GLKMatrixStackRef, radians float32, axisVector GLKVector3) error {
	if _gLKMatrixStackRotateWithVector3 == nil {
		return symbolCallError("GLKMatrixStackRotateWithVector3", "10.8", _gLKMatrixStackRotateWithVector3Err)
	}
	_gLKMatrixStackRotateWithVector3(stack, radians, axisVector)
	return nil
}

// GLKMatrixStackRotateWithVector3 replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a rotation around an arbitrary axis.
//
// See: https://developer.apple.com/documentation/GLKit/GLKMatrixStackRotateWithVector3(_:_:_:)
func GLKMatrixStackRotateWithVector3(stack GLKMatrixStackRef, radians float32, axisVector GLKVector3) {
	if callErr := tryGLKMatrixStackRotateWithVector3(stack, radians, axisVector); callErr != nil {
		panic(callErr)
	}
}

var _gLKMatrixStackRotateWithVector4 func(stack GLKMatrixStackRef, radians float32, axisVector GLKVector4)
var _gLKMatrixStackRotateWithVector4Err error

func tryGLKMatrixStackRotateWithVector4(stack GLKMatrixStackRef, radians float32, axisVector GLKVector4) error {
	if _gLKMatrixStackRotateWithVector4 == nil {
		return symbolCallError("GLKMatrixStackRotateWithVector4", "10.8", _gLKMatrixStackRotateWithVector4Err)
	}
	_gLKMatrixStackRotateWithVector4(stack, radians, axisVector)
	return nil
}

// GLKMatrixStackRotateWithVector4 replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a rotation around an arbitrary axis.
//
// See: https://developer.apple.com/documentation/GLKit/GLKMatrixStackRotateWithVector4(_:_:_:)
func GLKMatrixStackRotateWithVector4(stack GLKMatrixStackRef, radians float32, axisVector GLKVector4) {
	if callErr := tryGLKMatrixStackRotateWithVector4(stack, radians, axisVector); callErr != nil {
		panic(callErr)
	}
}

var _gLKMatrixStackRotateX func(stack GLKMatrixStackRef, radians float32)
var _gLKMatrixStackRotateXErr error

func tryGLKMatrixStackRotateX(stack GLKMatrixStackRef, radians float32) error {
	if _gLKMatrixStackRotateX == nil {
		return symbolCallError("GLKMatrixStackRotateX", "10.8", _gLKMatrixStackRotateXErr)
	}
	_gLKMatrixStackRotateX(stack, radians)
	return nil
}

// GLKMatrixStackRotateX replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a rotation around the positive-x axis.
//
// See: https://developer.apple.com/documentation/GLKit/GLKMatrixStackRotateX(_:_:)
func GLKMatrixStackRotateX(stack GLKMatrixStackRef, radians float32) {
	if callErr := tryGLKMatrixStackRotateX(stack, radians); callErr != nil {
		panic(callErr)
	}
}

var _gLKMatrixStackRotateY func(stack GLKMatrixStackRef, radians float32)
var _gLKMatrixStackRotateYErr error

func tryGLKMatrixStackRotateY(stack GLKMatrixStackRef, radians float32) error {
	if _gLKMatrixStackRotateY == nil {
		return symbolCallError("GLKMatrixStackRotateY", "10.8", _gLKMatrixStackRotateYErr)
	}
	_gLKMatrixStackRotateY(stack, radians)
	return nil
}

// GLKMatrixStackRotateY replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a rotation around the positive-y axis.
//
// See: https://developer.apple.com/documentation/GLKit/GLKMatrixStackRotateY(_:_:)
func GLKMatrixStackRotateY(stack GLKMatrixStackRef, radians float32) {
	if callErr := tryGLKMatrixStackRotateY(stack, radians); callErr != nil {
		panic(callErr)
	}
}

var _gLKMatrixStackRotateZ func(stack GLKMatrixStackRef, radians float32)
var _gLKMatrixStackRotateZErr error

func tryGLKMatrixStackRotateZ(stack GLKMatrixStackRef, radians float32) error {
	if _gLKMatrixStackRotateZ == nil {
		return symbolCallError("GLKMatrixStackRotateZ", "10.8", _gLKMatrixStackRotateZErr)
	}
	_gLKMatrixStackRotateZ(stack, radians)
	return nil
}

// GLKMatrixStackRotateZ replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a rotation around the positive-z axis.
//
// See: https://developer.apple.com/documentation/GLKit/GLKMatrixStackRotateZ(_:_:)
func GLKMatrixStackRotateZ(stack GLKMatrixStackRef, radians float32) {
	if callErr := tryGLKMatrixStackRotateZ(stack, radians); callErr != nil {
		panic(callErr)
	}
}

var _gLKMatrixStackScale func(stack GLKMatrixStackRef, sx float32, sy float32, sz float32)
var _gLKMatrixStackScaleErr error

func tryGLKMatrixStackScale(stack GLKMatrixStackRef, sx float32, sy float32, sz float32) error {
	if _gLKMatrixStackScale == nil {
		return symbolCallError("GLKMatrixStackScale", "10.8", _gLKMatrixStackScaleErr)
	}
	_gLKMatrixStackScale(stack, sx, sy, sz)
	return nil
}

// GLKMatrixStackScale replaces the contents of the top matrix with a matrix calculated by scaling the contents of the top matrix.
//
// See: https://developer.apple.com/documentation/GLKit/GLKMatrixStackScale(_:_:_:_:)
func GLKMatrixStackScale(stack GLKMatrixStackRef, sx float32, sy float32, sz float32) {
	if callErr := tryGLKMatrixStackScale(stack, sx, sy, sz); callErr != nil {
		panic(callErr)
	}
}

var _gLKMatrixStackScaleWithVector3 func(stack GLKMatrixStackRef, scaleVector GLKVector3)
var _gLKMatrixStackScaleWithVector3Err error

func tryGLKMatrixStackScaleWithVector3(stack GLKMatrixStackRef, scaleVector GLKVector3) error {
	if _gLKMatrixStackScaleWithVector3 == nil {
		return symbolCallError("GLKMatrixStackScaleWithVector3", "10.8", _gLKMatrixStackScaleWithVector3Err)
	}
	_gLKMatrixStackScaleWithVector3(stack, scaleVector)
	return nil
}

// GLKMatrixStackScaleWithVector3 replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a scaling operation.
//
// See: https://developer.apple.com/documentation/GLKit/GLKMatrixStackScaleWithVector3(_:_:)
func GLKMatrixStackScaleWithVector3(stack GLKMatrixStackRef, scaleVector GLKVector3) {
	if callErr := tryGLKMatrixStackScaleWithVector3(stack, scaleVector); callErr != nil {
		panic(callErr)
	}
}

var _gLKMatrixStackScaleWithVector4 func(stack GLKMatrixStackRef, scaleVector GLKVector4)
var _gLKMatrixStackScaleWithVector4Err error

func tryGLKMatrixStackScaleWithVector4(stack GLKMatrixStackRef, scaleVector GLKVector4) error {
	if _gLKMatrixStackScaleWithVector4 == nil {
		return symbolCallError("GLKMatrixStackScaleWithVector4", "10.8", _gLKMatrixStackScaleWithVector4Err)
	}
	_gLKMatrixStackScaleWithVector4(stack, scaleVector)
	return nil
}

// GLKMatrixStackScaleWithVector4 replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a scaling operation defined by a vector.
//
// See: https://developer.apple.com/documentation/GLKit/GLKMatrixStackScaleWithVector4(_:_:)
func GLKMatrixStackScaleWithVector4(stack GLKMatrixStackRef, scaleVector GLKVector4) {
	if callErr := tryGLKMatrixStackScaleWithVector4(stack, scaleVector); callErr != nil {
		panic(callErr)
	}
}

var _gLKMatrixStackSize func(stack GLKMatrixStackRef) int32
var _gLKMatrixStackSizeErr error

func tryGLKMatrixStackSize(stack GLKMatrixStackRef) (int32, error) {
	if _gLKMatrixStackSize == nil {
		return 0, symbolCallError("GLKMatrixStackSize", "10.8", _gLKMatrixStackSizeErr)
	}
	return _gLKMatrixStackSize(stack), nil
}

// GLKMatrixStackSize returns the number of matrices present on the matrix stack.
//
// See: https://developer.apple.com/documentation/GLKit/GLKMatrixStackSize(_:)
func GLKMatrixStackSize(stack GLKMatrixStackRef) int32 {
	result, callErr := tryGLKMatrixStackSize(stack)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gLKMatrixStackTranslate func(stack GLKMatrixStackRef, tx float32, ty float32, tz float32)
var _gLKMatrixStackTranslateErr error

func tryGLKMatrixStackTranslate(stack GLKMatrixStackRef, tx float32, ty float32, tz float32) error {
	if _gLKMatrixStackTranslate == nil {
		return symbolCallError("GLKMatrixStackTranslate", "10.8", _gLKMatrixStackTranslateErr)
	}
	_gLKMatrixStackTranslate(stack, tx, ty, tz)
	return nil
}

// GLKMatrixStackTranslate replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a translation operation.
//
// See: https://developer.apple.com/documentation/GLKit/GLKMatrixStackTranslate(_:_:_:_:)
func GLKMatrixStackTranslate(stack GLKMatrixStackRef, tx float32, ty float32, tz float32) {
	if callErr := tryGLKMatrixStackTranslate(stack, tx, ty, tz); callErr != nil {
		panic(callErr)
	}
}

var _gLKMatrixStackTranslateWithVector3 func(stack GLKMatrixStackRef, translationVector GLKVector3)
var _gLKMatrixStackTranslateWithVector3Err error

func tryGLKMatrixStackTranslateWithVector3(stack GLKMatrixStackRef, translationVector GLKVector3) error {
	if _gLKMatrixStackTranslateWithVector3 == nil {
		return symbolCallError("GLKMatrixStackTranslateWithVector3", "10.8", _gLKMatrixStackTranslateWithVector3Err)
	}
	_gLKMatrixStackTranslateWithVector3(stack, translationVector)
	return nil
}

// GLKMatrixStackTranslateWithVector3 replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a translation defined by a vector.
//
// See: https://developer.apple.com/documentation/GLKit/GLKMatrixStackTranslateWithVector3(_:_:)
func GLKMatrixStackTranslateWithVector3(stack GLKMatrixStackRef, translationVector GLKVector3) {
	if callErr := tryGLKMatrixStackTranslateWithVector3(stack, translationVector); callErr != nil {
		panic(callErr)
	}
}

var _gLKMatrixStackTranslateWithVector4 func(stack GLKMatrixStackRef, translationVector GLKVector4)
var _gLKMatrixStackTranslateWithVector4Err error

func tryGLKMatrixStackTranslateWithVector4(stack GLKMatrixStackRef, translationVector GLKVector4) error {
	if _gLKMatrixStackTranslateWithVector4 == nil {
		return symbolCallError("GLKMatrixStackTranslateWithVector4", "10.8", _gLKMatrixStackTranslateWithVector4Err)
	}
	_gLKMatrixStackTranslateWithVector4(stack, translationVector)
	return nil
}

// GLKMatrixStackTranslateWithVector4 replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a translation defined by a vector.
//
// See: https://developer.apple.com/documentation/GLKit/GLKMatrixStackTranslateWithVector4(_:_:)
func GLKMatrixStackTranslateWithVector4(stack GLKMatrixStackRef, translationVector GLKVector4) {
	if callErr := tryGLKMatrixStackTranslateWithVector4(stack, translationVector); callErr != nil {
		panic(callErr)
	}
}

var _gLKQuaternionAngle func(quaternion GLKQuaternion) float32
var _gLKQuaternionAngleErr error

func tryGLKQuaternionAngle(quaternion GLKQuaternion) (float32, error) {
	if _gLKQuaternionAngle == nil {
		return 0.0, symbolCallError("GLKQuaternionAngle", "10.8", _gLKQuaternionAngleErr)
	}
	return _gLKQuaternionAngle(quaternion), nil
}

// GLKQuaternionAngle returns the rotation angle of a quaternion.
//
// See: https://developer.apple.com/documentation/GLKit/GLKQuaternionAngle(_:)
func GLKQuaternionAngle(quaternion GLKQuaternion) float32 {
	result, callErr := tryGLKQuaternionAngle(quaternion)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gLKQuaternionAxis func(quaternion GLKQuaternion) GLKVector3
var _gLKQuaternionAxisErr error

func tryGLKQuaternionAxis(quaternion GLKQuaternion) (GLKVector3, error) {
	if _gLKQuaternionAxis == nil {
		return *new(GLKVector3), symbolCallError("GLKQuaternionAxis", "10.8", _gLKQuaternionAxisErr)
	}
	return _gLKQuaternionAxis(quaternion), nil
}

// GLKQuaternionAxis returns the axis of rotation of a quaternion.
//
// See: https://developer.apple.com/documentation/GLKit/GLKQuaternionAxis(_:)
func GLKQuaternionAxis(quaternion GLKQuaternion) GLKVector3 {
	result, callErr := tryGLKQuaternionAxis(quaternion)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gLKQuaternionMakeWithMatrix3 func(matrix GLKMatrix3) GLKQuaternion
var _gLKQuaternionMakeWithMatrix3Err error

func tryGLKQuaternionMakeWithMatrix3(matrix GLKMatrix3) (GLKQuaternion, error) {
	if _gLKQuaternionMakeWithMatrix3 == nil {
		return *new(GLKQuaternion), symbolCallError("GLKQuaternionMakeWithMatrix3", "10.8", _gLKQuaternionMakeWithMatrix3Err)
	}
	return _gLKQuaternionMakeWithMatrix3(matrix), nil
}

// GLKQuaternionMakeWithMatrix3 creates a quaternion from a rotation matrix.
//
// See: https://developer.apple.com/documentation/GLKit/GLKQuaternionMakeWithMatrix3(_:)
func GLKQuaternionMakeWithMatrix3(matrix GLKMatrix3) GLKQuaternion {
	result, callErr := tryGLKQuaternionMakeWithMatrix3(matrix)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gLKQuaternionMakeWithMatrix4 func(matrix GLKMatrix4) GLKQuaternion
var _gLKQuaternionMakeWithMatrix4Err error

func tryGLKQuaternionMakeWithMatrix4(matrix GLKMatrix4) (GLKQuaternion, error) {
	if _gLKQuaternionMakeWithMatrix4 == nil {
		return *new(GLKQuaternion), symbolCallError("GLKQuaternionMakeWithMatrix4", "10.8", _gLKQuaternionMakeWithMatrix4Err)
	}
	return _gLKQuaternionMakeWithMatrix4(matrix), nil
}

// GLKQuaternionMakeWithMatrix4 creates a quaternion from a rotation matrix.
//
// See: https://developer.apple.com/documentation/GLKit/GLKQuaternionMakeWithMatrix4(_:)
func GLKQuaternionMakeWithMatrix4(matrix GLKMatrix4) GLKQuaternion {
	result, callErr := tryGLKQuaternionMakeWithMatrix4(matrix)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gLKQuaternionRotateVector3Array func(quaternion GLKQuaternion, vectors *GLKVector3, vectorCount uintptr)
var _gLKQuaternionRotateVector3ArrayErr error

func tryGLKQuaternionRotateVector3Array(quaternion GLKQuaternion, vectors *GLKVector3, vectorCount uintptr) error {
	if _gLKQuaternionRotateVector3Array == nil {
		return symbolCallError("GLKQuaternionRotateVector3Array", "10.8", _gLKQuaternionRotateVector3ArrayErr)
	}
	_gLKQuaternionRotateVector3Array(quaternion, vectors, vectorCount)
	return nil
}

// GLKQuaternionRotateVector3Array applies a quaternion rotation to an array of vectors.
//
// See: https://developer.apple.com/documentation/GLKit/GLKQuaternionRotateVector3Array(_:_:_:)
func GLKQuaternionRotateVector3Array(quaternion GLKQuaternion, vectors *GLKVector3, vectorCount uintptr) {
	if callErr := tryGLKQuaternionRotateVector3Array(quaternion, vectors, vectorCount); callErr != nil {
		panic(callErr)
	}
}

var _gLKQuaternionRotateVector4Array func(quaternion GLKQuaternion, vectors *GLKVector4, vectorCount uintptr)
var _gLKQuaternionRotateVector4ArrayErr error

func tryGLKQuaternionRotateVector4Array(quaternion GLKQuaternion, vectors *GLKVector4, vectorCount uintptr) error {
	if _gLKQuaternionRotateVector4Array == nil {
		return symbolCallError("GLKQuaternionRotateVector4Array", "10.8", _gLKQuaternionRotateVector4ArrayErr)
	}
	_gLKQuaternionRotateVector4Array(quaternion, vectors, vectorCount)
	return nil
}

// GLKQuaternionRotateVector4Array applies a quaternion rotation to an array of vectors.
//
// See: https://developer.apple.com/documentation/GLKit/GLKQuaternionRotateVector4Array(_:_:_:)
func GLKQuaternionRotateVector4Array(quaternion GLKQuaternion, vectors *GLKVector4, vectorCount uintptr) {
	if callErr := tryGLKQuaternionRotateVector4Array(quaternion, vectors, vectorCount); callErr != nil {
		panic(callErr)
	}
}

var _gLKQuaternionSlerp func(quaternionStart GLKQuaternion, quaternionEnd GLKQuaternion, t float32) GLKQuaternion
var _gLKQuaternionSlerpErr error

func tryGLKQuaternionSlerp(quaternionStart GLKQuaternion, quaternionEnd GLKQuaternion, t float32) (GLKQuaternion, error) {
	if _gLKQuaternionSlerp == nil {
		return *new(GLKQuaternion), symbolCallError("GLKQuaternionSlerp", "10.8", _gLKQuaternionSlerpErr)
	}
	return _gLKQuaternionSlerp(quaternionStart, quaternionEnd, t), nil
}

// GLKQuaternionSlerp returns the spherical linear interpolation of two quaternions.
//
// See: https://developer.apple.com/documentation/GLKit/GLKQuaternionSlerp(_:_:_:)
func GLKQuaternionSlerp(quaternionStart GLKQuaternion, quaternionEnd GLKQuaternion, t float32) GLKQuaternion {
	result, callErr := tryGLKQuaternionSlerp(quaternionStart, quaternionEnd, t)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gLKVertexAttributeParametersFromModelIO func(vertexFormat uint) GLKVertexAttributeParameters
var _gLKVertexAttributeParametersFromModelIOErr error

func tryGLKVertexAttributeParametersFromModelIO(vertexFormat uint) (GLKVertexAttributeParameters, error) {
	if _gLKVertexAttributeParametersFromModelIO == nil {
		return *new(GLKVertexAttributeParameters), symbolCallError("GLKVertexAttributeParametersFromModelIO", "10.8", _gLKVertexAttributeParametersFromModelIOErr)
	}
	return _gLKVertexAttributeParametersFromModelIO(vertexFormat), nil
}

// GLKVertexAttributeParametersFromModelIO.
//
// See: https://developer.apple.com/documentation/GLKit/GLKVertexAttributeParametersFromModelIO(_:)
func GLKVertexAttributeParametersFromModelIO(vertexFormat uint) GLKVertexAttributeParameters {
	result, callErr := tryGLKVertexAttributeParametersFromModelIO(vertexFormat)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSStringFromGLKMatrix2 func(matrix GLKMatrix2) *foundation.NSString
var _nSStringFromGLKMatrix2Err error

func tryNSStringFromGLKMatrix2(matrix GLKMatrix2) (*foundation.NSString, error) {
	if _nSStringFromGLKMatrix2 == nil {
		return nil, symbolCallError("NSStringFromGLKMatrix2", "10.8", _nSStringFromGLKMatrix2Err)
	}
	return _nSStringFromGLKMatrix2(matrix), nil
}

// NSStringFromGLKMatrix2 returns a string that represents the contents of a matrix.
//
// See: https://developer.apple.com/documentation/GLKit/NSStringFromGLKMatrix2(_:)
func NSStringFromGLKMatrix2(matrix GLKMatrix2) *foundation.NSString {
	result, callErr := tryNSStringFromGLKMatrix2(matrix)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSStringFromGLKMatrix3 func(matrix GLKMatrix3) *foundation.NSString
var _nSStringFromGLKMatrix3Err error

func tryNSStringFromGLKMatrix3(matrix GLKMatrix3) (*foundation.NSString, error) {
	if _nSStringFromGLKMatrix3 == nil {
		return nil, symbolCallError("NSStringFromGLKMatrix3", "10.8", _nSStringFromGLKMatrix3Err)
	}
	return _nSStringFromGLKMatrix3(matrix), nil
}

// NSStringFromGLKMatrix3 returns a string that represents the contents of a matrix.
//
// See: https://developer.apple.com/documentation/GLKit/NSStringFromGLKMatrix3(_:)
func NSStringFromGLKMatrix3(matrix GLKMatrix3) *foundation.NSString {
	result, callErr := tryNSStringFromGLKMatrix3(matrix)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSStringFromGLKMatrix4 func(matrix GLKMatrix4) *foundation.NSString
var _nSStringFromGLKMatrix4Err error

func tryNSStringFromGLKMatrix4(matrix GLKMatrix4) (*foundation.NSString, error) {
	if _nSStringFromGLKMatrix4 == nil {
		return nil, symbolCallError("NSStringFromGLKMatrix4", "10.8", _nSStringFromGLKMatrix4Err)
	}
	return _nSStringFromGLKMatrix4(matrix), nil
}

// NSStringFromGLKMatrix4 returns a string that represents the contents of a matrix.
//
// See: https://developer.apple.com/documentation/GLKit/NSStringFromGLKMatrix4(_:)
func NSStringFromGLKMatrix4(matrix GLKMatrix4) *foundation.NSString {
	result, callErr := tryNSStringFromGLKMatrix4(matrix)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSStringFromGLKQuaternion func(quaternion GLKQuaternion) *foundation.NSString
var _nSStringFromGLKQuaternionErr error

func tryNSStringFromGLKQuaternion(quaternion GLKQuaternion) (*foundation.NSString, error) {
	if _nSStringFromGLKQuaternion == nil {
		return nil, symbolCallError("NSStringFromGLKQuaternion", "10.8", _nSStringFromGLKQuaternionErr)
	}
	return _nSStringFromGLKQuaternion(quaternion), nil
}

// NSStringFromGLKQuaternion returns a string that represents the contents of a quaternion.
//
// See: https://developer.apple.com/documentation/GLKit/NSStringFromGLKQuaternion(_:)
func NSStringFromGLKQuaternion(quaternion GLKQuaternion) *foundation.NSString {
	result, callErr := tryNSStringFromGLKQuaternion(quaternion)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSStringFromGLKVector2 func(vector GLKVector2) *foundation.NSString
var _nSStringFromGLKVector2Err error

func tryNSStringFromGLKVector2(vector GLKVector2) (*foundation.NSString, error) {
	if _nSStringFromGLKVector2 == nil {
		return nil, symbolCallError("NSStringFromGLKVector2", "10.8", _nSStringFromGLKVector2Err)
	}
	return _nSStringFromGLKVector2(vector), nil
}

// NSStringFromGLKVector2 returns a string that represents the contents of a vector.
//
// See: https://developer.apple.com/documentation/GLKit/NSStringFromGLKVector2(_:)
func NSStringFromGLKVector2(vector GLKVector2) *foundation.NSString {
	result, callErr := tryNSStringFromGLKVector2(vector)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSStringFromGLKVector3 func(vector GLKVector3) *foundation.NSString
var _nSStringFromGLKVector3Err error

func tryNSStringFromGLKVector3(vector GLKVector3) (*foundation.NSString, error) {
	if _nSStringFromGLKVector3 == nil {
		return nil, symbolCallError("NSStringFromGLKVector3", "10.8", _nSStringFromGLKVector3Err)
	}
	return _nSStringFromGLKVector3(vector), nil
}

// NSStringFromGLKVector3 returns a string that represents the contents of a vector.
//
// See: https://developer.apple.com/documentation/GLKit/NSStringFromGLKVector3(_:)
func NSStringFromGLKVector3(vector GLKVector3) *foundation.NSString {
	result, callErr := tryNSStringFromGLKVector3(vector)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSStringFromGLKVector4 func(vector GLKVector4) *foundation.NSString
var _nSStringFromGLKVector4Err error

func tryNSStringFromGLKVector4(vector GLKVector4) (*foundation.NSString, error) {
	if _nSStringFromGLKVector4 == nil {
		return nil, symbolCallError("NSStringFromGLKVector4", "10.8", _nSStringFromGLKVector4Err)
	}
	return _nSStringFromGLKVector4(vector), nil
}

// NSStringFromGLKVector4 returns a string that represents the contents of a vector.
//
// See: https://developer.apple.com/documentation/GLKit/NSStringFromGLKVector4(_:)
func NSStringFromGLKVector4(vector GLKVector4) *foundation.NSString {
	result, callErr := tryNSStringFromGLKVector4(vector)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_gLKMathProject, &_gLKMathProjectErr, frameworkHandle, "GLKMathProject", "10.8")
	registerFunc(&_gLKMathUnproject, &_gLKMathUnprojectErr, frameworkHandle, "GLKMathUnproject", "10.8")
	registerFunc(&_gLKMatrix3Invert, &_gLKMatrix3InvertErr, frameworkHandle, "GLKMatrix3Invert", "10.8")
	registerFunc(&_gLKMatrix3InvertAndTranspose, &_gLKMatrix3InvertAndTransposeErr, frameworkHandle, "GLKMatrix3InvertAndTranspose", "10.8")
	registerFunc(&_gLKMatrix4Invert, &_gLKMatrix4InvertErr, frameworkHandle, "GLKMatrix4Invert", "10.8")
	registerFunc(&_gLKMatrix4InvertAndTranspose, &_gLKMatrix4InvertAndTransposeErr, frameworkHandle, "GLKMatrix4InvertAndTranspose", "10.8")
	registerFunc(&_gLKMatrixStackCreate, &_gLKMatrixStackCreateErr, frameworkHandle, "GLKMatrixStackCreate", "10.8")
	registerFunc(&_gLKMatrixStackGetMatrix2, &_gLKMatrixStackGetMatrix2Err, frameworkHandle, "GLKMatrixStackGetMatrix2", "10.8")
	registerFunc(&_gLKMatrixStackGetMatrix3, &_gLKMatrixStackGetMatrix3Err, frameworkHandle, "GLKMatrixStackGetMatrix3", "10.8")
	registerFunc(&_gLKMatrixStackGetMatrix3Inverse, &_gLKMatrixStackGetMatrix3InverseErr, frameworkHandle, "GLKMatrixStackGetMatrix3Inverse", "10.8")
	registerFunc(&_gLKMatrixStackGetMatrix3InverseTranspose, &_gLKMatrixStackGetMatrix3InverseTransposeErr, frameworkHandle, "GLKMatrixStackGetMatrix3InverseTranspose", "10.8")
	registerFunc(&_gLKMatrixStackGetMatrix4, &_gLKMatrixStackGetMatrix4Err, frameworkHandle, "GLKMatrixStackGetMatrix4", "10.8")
	registerFunc(&_gLKMatrixStackGetMatrix4Inverse, &_gLKMatrixStackGetMatrix4InverseErr, frameworkHandle, "GLKMatrixStackGetMatrix4Inverse", "10.8")
	registerFunc(&_gLKMatrixStackGetMatrix4InverseTranspose, &_gLKMatrixStackGetMatrix4InverseTransposeErr, frameworkHandle, "GLKMatrixStackGetMatrix4InverseTranspose", "10.8")
	registerFunc(&_gLKMatrixStackGetTypeID, &_gLKMatrixStackGetTypeIDErr, frameworkHandle, "GLKMatrixStackGetTypeID", "10.8")
	registerFunc(&_gLKMatrixStackLoadMatrix4, &_gLKMatrixStackLoadMatrix4Err, frameworkHandle, "GLKMatrixStackLoadMatrix4", "10.8")
	registerFunc(&_gLKMatrixStackMultiplyMatrix4, &_gLKMatrixStackMultiplyMatrix4Err, frameworkHandle, "GLKMatrixStackMultiplyMatrix4", "10.8")
	registerFunc(&_gLKMatrixStackMultiplyMatrixStack, &_gLKMatrixStackMultiplyMatrixStackErr, frameworkHandle, "GLKMatrixStackMultiplyMatrixStack", "10.8")
	registerFunc(&_gLKMatrixStackPop, &_gLKMatrixStackPopErr, frameworkHandle, "GLKMatrixStackPop", "10.8")
	registerFunc(&_gLKMatrixStackPush, &_gLKMatrixStackPushErr, frameworkHandle, "GLKMatrixStackPush", "10.8")
	registerFunc(&_gLKMatrixStackRotate, &_gLKMatrixStackRotateErr, frameworkHandle, "GLKMatrixStackRotate", "10.8")
	registerFunc(&_gLKMatrixStackRotateWithVector3, &_gLKMatrixStackRotateWithVector3Err, frameworkHandle, "GLKMatrixStackRotateWithVector3", "10.8")
	registerFunc(&_gLKMatrixStackRotateWithVector4, &_gLKMatrixStackRotateWithVector4Err, frameworkHandle, "GLKMatrixStackRotateWithVector4", "10.8")
	registerFunc(&_gLKMatrixStackRotateX, &_gLKMatrixStackRotateXErr, frameworkHandle, "GLKMatrixStackRotateX", "10.8")
	registerFunc(&_gLKMatrixStackRotateY, &_gLKMatrixStackRotateYErr, frameworkHandle, "GLKMatrixStackRotateY", "10.8")
	registerFunc(&_gLKMatrixStackRotateZ, &_gLKMatrixStackRotateZErr, frameworkHandle, "GLKMatrixStackRotateZ", "10.8")
	registerFunc(&_gLKMatrixStackScale, &_gLKMatrixStackScaleErr, frameworkHandle, "GLKMatrixStackScale", "10.8")
	registerFunc(&_gLKMatrixStackScaleWithVector3, &_gLKMatrixStackScaleWithVector3Err, frameworkHandle, "GLKMatrixStackScaleWithVector3", "10.8")
	registerFunc(&_gLKMatrixStackScaleWithVector4, &_gLKMatrixStackScaleWithVector4Err, frameworkHandle, "GLKMatrixStackScaleWithVector4", "10.8")
	registerFunc(&_gLKMatrixStackSize, &_gLKMatrixStackSizeErr, frameworkHandle, "GLKMatrixStackSize", "10.8")
	registerFunc(&_gLKMatrixStackTranslate, &_gLKMatrixStackTranslateErr, frameworkHandle, "GLKMatrixStackTranslate", "10.8")
	registerFunc(&_gLKMatrixStackTranslateWithVector3, &_gLKMatrixStackTranslateWithVector3Err, frameworkHandle, "GLKMatrixStackTranslateWithVector3", "10.8")
	registerFunc(&_gLKMatrixStackTranslateWithVector4, &_gLKMatrixStackTranslateWithVector4Err, frameworkHandle, "GLKMatrixStackTranslateWithVector4", "10.8")
	registerFunc(&_gLKQuaternionAngle, &_gLKQuaternionAngleErr, frameworkHandle, "GLKQuaternionAngle", "10.8")
	registerFunc(&_gLKQuaternionAxis, &_gLKQuaternionAxisErr, frameworkHandle, "GLKQuaternionAxis", "10.8")
	registerFunc(&_gLKQuaternionMakeWithMatrix3, &_gLKQuaternionMakeWithMatrix3Err, frameworkHandle, "GLKQuaternionMakeWithMatrix3", "10.8")
	registerFunc(&_gLKQuaternionMakeWithMatrix4, &_gLKQuaternionMakeWithMatrix4Err, frameworkHandle, "GLKQuaternionMakeWithMatrix4", "10.8")
	registerFunc(&_gLKQuaternionRotateVector3Array, &_gLKQuaternionRotateVector3ArrayErr, frameworkHandle, "GLKQuaternionRotateVector3Array", "10.8")
	registerFunc(&_gLKQuaternionRotateVector4Array, &_gLKQuaternionRotateVector4ArrayErr, frameworkHandle, "GLKQuaternionRotateVector4Array", "10.8")
	registerFunc(&_gLKQuaternionSlerp, &_gLKQuaternionSlerpErr, frameworkHandle, "GLKQuaternionSlerp", "10.8")
	registerFunc(&_gLKVertexAttributeParametersFromModelIO, &_gLKVertexAttributeParametersFromModelIOErr, frameworkHandle, "GLKVertexAttributeParametersFromModelIO", "10.8")
	registerFunc(&_nSStringFromGLKMatrix2, &_nSStringFromGLKMatrix2Err, frameworkHandle, "NSStringFromGLKMatrix2", "10.8")
	registerFunc(&_nSStringFromGLKMatrix3, &_nSStringFromGLKMatrix3Err, frameworkHandle, "NSStringFromGLKMatrix3", "10.8")
	registerFunc(&_nSStringFromGLKMatrix4, &_nSStringFromGLKMatrix4Err, frameworkHandle, "NSStringFromGLKMatrix4", "10.8")
	registerFunc(&_nSStringFromGLKQuaternion, &_nSStringFromGLKQuaternionErr, frameworkHandle, "NSStringFromGLKQuaternion", "10.8")
	registerFunc(&_nSStringFromGLKVector2, &_nSStringFromGLKVector2Err, frameworkHandle, "NSStringFromGLKVector2", "10.8")
	registerFunc(&_nSStringFromGLKVector3, &_nSStringFromGLKVector3Err, frameworkHandle, "NSStringFromGLKVector3", "10.8")
	registerFunc(&_nSStringFromGLKVector4, &_nSStringFromGLKVector4Err, frameworkHandle, "NSStringFromGLKVector4", "10.8")
}
