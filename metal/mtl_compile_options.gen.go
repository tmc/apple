// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLCompileOptions] class.
var (
	_MTLCompileOptionsClass     MTLCompileOptionsClass
	_MTLCompileOptionsClassOnce sync.Once
)

func getMTLCompileOptionsClass() MTLCompileOptionsClass {
	_MTLCompileOptionsClassOnce.Do(func() {
		_MTLCompileOptionsClass = MTLCompileOptionsClass{class: objc.GetClass("MTLCompileOptions")}
	})
	return _MTLCompileOptionsClass
}

// GetMTLCompileOptionsClass returns the class object for MTLCompileOptions.
func GetMTLCompileOptionsClass() MTLCompileOptionsClass {
	return getMTLCompileOptionsClass()
}

type MTLCompileOptionsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLCompileOptionsClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLCompileOptionsClass) Alloc() MTLCompileOptions {
	rv := objc.Send[MTLCompileOptions](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Compilation settings for a Metal shader library.
//
// # Overview
//
// You can configure the Metal compiler’s options by setting any or all of
// an [MTLCompileOptions] instance’s properties, including the following:
//
// - Target previous OS releases by assigning the [MTLCompileOptions.LanguageVersion] property
// to an [MTLLanguageVersion] case. - Set preprocessor macros for the Metal
// compiler by assigning a dictionary to the [MTLCompileOptions.PreprocessorMacros] property. -
// Choose what the Metal compiler’s optimizer prioritizes by setting the
// [MTLCompileOptions.OptimizationLevel] property to an [MTLLibraryOptimizationLevel] case. -
// Allow the compiler to optimize for floating-point arithmetic that may
// violate the IEEE 754 standard by setting [MTLCompileOptions.MathMode] to [MTLMathMode.fast].
//
// You can compile a library with your compile options instance by calling an
// [MTLDevice] instance’s [NewLibraryWithSourceOptionsError] or
// [NewLibraryWithSourceOptionsCompletionHandler] method.
//
// # Configuring the compiler options
//
//   - [MTLCompileOptions.EnableLogging]: A Boolean value that enables shader logging.
//   - [MTLCompileOptions.SetEnableLogging]
//   - [MTLCompileOptions.MathMode]: An indication of whether the compiler can perform optimizations for floating-point arithmetic that may violate the IEEE 754 standard.
//   - [MTLCompileOptions.SetMathMode]
//   - [MTLCompileOptions.MathFloatingPointFunctions]: The FP32 math functions Metal uses.
//   - [MTLCompileOptions.SetMathFloatingPointFunctions]
//   - [MTLCompileOptions.PreserveInvariance]: A Boolean value that indicates whether the compiler compiles vertex shaders conservatively to generate consistent position calculations.
//   - [MTLCompileOptions.SetPreserveInvariance]
//   - [MTLCompileOptions.LanguageVersion]: The language version for interpreting the library source code.
//   - [MTLCompileOptions.SetLanguageVersion]
//   - [MTLCompileOptions.PreprocessorMacros]: A list of preprocessor macros to apply when compiling the library source.
//   - [MTLCompileOptions.SetPreprocessorMacros]
//   - [MTLCompileOptions.OptimizationLevel]: An option that tells the compiler what to prioritize when it compiles Metal shader code.
//   - [MTLCompileOptions.SetOptimizationLevel]
//   - [MTLCompileOptions.Libraries]: An array of dynamic libraries the Metal compiler links against.
//   - [MTLCompileOptions.SetLibraries]
//   - [MTLCompileOptions.FastMathEnabled]: A Boolean value that indicates whether the compiler can perform optimizations for floating-point arithmetic that may violate the IEEE 754 standard.
//   - [MTLCompileOptions.SetFastMathEnabled]
//
// # Configuring the library output options
//
//   - [MTLCompileOptions.LibraryType]: The kind of library to create.
//   - [MTLCompileOptions.SetLibraryType]
//   - [MTLCompileOptions.InstallName]: For a dynamic library, the name to use when installing the library.
//   - [MTLCompileOptions.SetInstallName]
//
// # Instance Properties
//
//   - [MTLCompileOptions.AllowReferencingUndefinedSymbols]
//   - [MTLCompileOptions.SetAllowReferencingUndefinedSymbols]
//   - [MTLCompileOptions.CompileSymbolVisibility]
//   - [MTLCompileOptions.SetCompileSymbolVisibility]
//   - [MTLCompileOptions.MaxTotalThreadsPerThreadgroup]
//   - [MTLCompileOptions.SetMaxTotalThreadsPerThreadgroup]
//   - [MTLCompileOptions.RequiredThreadsPerThreadgroup]
//   - [MTLCompileOptions.SetRequiredThreadsPerThreadgroup]
//
// See: https://developer.apple.com/documentation/Metal/MTLCompileOptions
//
// [MTLLanguageVersion]: https://developer.apple.com/documentation/Metal/MTLLanguageVersion
// [MTLLibraryOptimizationLevel]: https://developer.apple.com/documentation/Metal/MTLLibraryOptimizationLevel
// [MTLMathMode.fast]: https://developer.apple.com/documentation/Metal/MTLMathMode/fast
type MTLCompileOptions struct {
	objectivec.Object
}

// MTLCompileOptionsFromID constructs a [MTLCompileOptions] from an objc.ID.
//
// Compilation settings for a Metal shader library.
func MTLCompileOptionsFromID(id objc.ID) MTLCompileOptions {
	return MTLCompileOptions{objectivec.Object{ID: id}}
}

// NOTE: MTLCompileOptions adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLCompileOptions] class.
//
// # Configuring the compiler options
//
//   - [IMTLCompileOptions.EnableLogging]: A Boolean value that enables shader logging.
//   - [IMTLCompileOptions.SetEnableLogging]
//   - [IMTLCompileOptions.MathMode]: An indication of whether the compiler can perform optimizations for floating-point arithmetic that may violate the IEEE 754 standard.
//   - [IMTLCompileOptions.SetMathMode]
//   - [IMTLCompileOptions.MathFloatingPointFunctions]: The FP32 math functions Metal uses.
//   - [IMTLCompileOptions.SetMathFloatingPointFunctions]
//   - [IMTLCompileOptions.PreserveInvariance]: A Boolean value that indicates whether the compiler compiles vertex shaders conservatively to generate consistent position calculations.
//   - [IMTLCompileOptions.SetPreserveInvariance]
//   - [IMTLCompileOptions.LanguageVersion]: The language version for interpreting the library source code.
//   - [IMTLCompileOptions.SetLanguageVersion]
//   - [IMTLCompileOptions.PreprocessorMacros]: A list of preprocessor macros to apply when compiling the library source.
//   - [IMTLCompileOptions.SetPreprocessorMacros]
//   - [IMTLCompileOptions.OptimizationLevel]: An option that tells the compiler what to prioritize when it compiles Metal shader code.
//   - [IMTLCompileOptions.SetOptimizationLevel]
//   - [IMTLCompileOptions.Libraries]: An array of dynamic libraries the Metal compiler links against.
//   - [IMTLCompileOptions.SetLibraries]
//   - [IMTLCompileOptions.FastMathEnabled]: A Boolean value that indicates whether the compiler can perform optimizations for floating-point arithmetic that may violate the IEEE 754 standard.
//   - [IMTLCompileOptions.SetFastMathEnabled]
//
// # Configuring the library output options
//
//   - [IMTLCompileOptions.LibraryType]: The kind of library to create.
//   - [IMTLCompileOptions.SetLibraryType]
//   - [IMTLCompileOptions.InstallName]: For a dynamic library, the name to use when installing the library.
//   - [IMTLCompileOptions.SetInstallName]
//
// # Instance Properties
//
//   - [IMTLCompileOptions.AllowReferencingUndefinedSymbols]
//   - [IMTLCompileOptions.SetAllowReferencingUndefinedSymbols]
//   - [IMTLCompileOptions.CompileSymbolVisibility]
//   - [IMTLCompileOptions.SetCompileSymbolVisibility]
//   - [IMTLCompileOptions.MaxTotalThreadsPerThreadgroup]
//   - [IMTLCompileOptions.SetMaxTotalThreadsPerThreadgroup]
//   - [IMTLCompileOptions.RequiredThreadsPerThreadgroup]
//   - [IMTLCompileOptions.SetRequiredThreadsPerThreadgroup]
//
// See: https://developer.apple.com/documentation/Metal/MTLCompileOptions
type IMTLCompileOptions interface {
	objectivec.IObject

	// Topic: Configuring the compiler options

	// A Boolean value that enables shader logging.
	EnableLogging() bool
	SetEnableLogging(value bool)
	// An indication of whether the compiler can perform optimizations for floating-point arithmetic that may violate the IEEE 754 standard.
	MathMode() MTLMathMode
	SetMathMode(value MTLMathMode)
	// The FP32 math functions Metal uses.
	MathFloatingPointFunctions() MTLMathFloatingPointFunctions
	SetMathFloatingPointFunctions(value MTLMathFloatingPointFunctions)
	// A Boolean value that indicates whether the compiler compiles vertex shaders conservatively to generate consistent position calculations.
	PreserveInvariance() bool
	SetPreserveInvariance(value bool)
	// The language version for interpreting the library source code.
	LanguageVersion() MTLLanguageVersion
	SetLanguageVersion(value MTLLanguageVersion)
	// A list of preprocessor macros to apply when compiling the library source.
	PreprocessorMacros() foundation.INSDictionary
	SetPreprocessorMacros(value foundation.INSDictionary)
	// An option that tells the compiler what to prioritize when it compiles Metal shader code.
	OptimizationLevel() MTLLibraryOptimizationLevel
	SetOptimizationLevel(value MTLLibraryOptimizationLevel)
	// An array of dynamic libraries the Metal compiler links against.
	Libraries() []objectivec.IObject
	SetLibraries(value []objectivec.IObject)
	// A Boolean value that indicates whether the compiler can perform optimizations for floating-point arithmetic that may violate the IEEE 754 standard.
	FastMathEnabled() bool
	SetFastMathEnabled(value bool)

	// Topic: Configuring the library output options

	// The kind of library to create.
	LibraryType() MTLLibraryType
	SetLibraryType(value MTLLibraryType)
	// For a dynamic library, the name to use when installing the library.
	InstallName() string
	SetInstallName(value string)

	// Topic: Instance Properties

	AllowReferencingUndefinedSymbols() bool
	SetAllowReferencingUndefinedSymbols(value bool)
	CompileSymbolVisibility() MTLCompileSymbolVisibility
	SetCompileSymbolVisibility(value MTLCompileSymbolVisibility)
	MaxTotalThreadsPerThreadgroup() uint
	SetMaxTotalThreadsPerThreadgroup(value uint)
	RequiredThreadsPerThreadgroup() MTLSize
	SetRequiredThreadsPerThreadgroup(value MTLSize)
}

// Init initializes the instance.
func (c MTLCompileOptions) Init() MTLCompileOptions {
	rv := objc.Send[MTLCompileOptions](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MTLCompileOptions) Autorelease() MTLCompileOptions {
	rv := objc.Send[MTLCompileOptions](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLCompileOptions creates a new MTLCompileOptions instance.
func NewMTLCompileOptions() MTLCompileOptions {
	class := getMTLCompileOptionsClass()
	rv := objc.Send[MTLCompileOptions](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Boolean value that enables shader logging.
//
// # Discussion
//
// Because logging incurs overhead, regardless of whether the system prints
// messages, you need to explicitly enable logging.
//
// See: https://developer.apple.com/documentation/Metal/MTLCompileOptions/enableLogging
func (c MTLCompileOptions) EnableLogging() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("enableLogging"))
	return rv
}
func (c MTLCompileOptions) SetEnableLogging(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setEnableLogging:"), value)
}

// An indication of whether the compiler can perform optimizations for
// floating-point arithmetic that may violate the IEEE 754 standard.
//
// # Discussion
//
// This property replaces the [FastMathEnabled] property.
//
// If [FastMathEnabled] is `true`, the system sets [MathMode] to
// [MTLMathMode.fast] and [MathFloatingPointFunctions] to
// [MTLMathFloatingPointFunctions.fast].
//
// If [FastMathEnabled] is `false`, the system sets [MathMode] to
// [MTLMathModeSafe] and [MathFloatingPointFunctions] to
// [MTLMathFloatingPointFunctionsPrecise].
//
// Subsequent calls to [MathMode] or [MathFloatingPointFunctions] set the
// variables directly.
//
// See: https://developer.apple.com/documentation/Metal/MTLCompileOptions/mathMode
//
// [MTLMathFloatingPointFunctions.fast]: https://developer.apple.com/documentation/Metal/MTLMathFloatingPointFunctions/fast
// [MTLMathMode.fast]: https://developer.apple.com/documentation/Metal/MTLMathMode/fast
func (c MTLCompileOptions) MathMode() MTLMathMode {
	rv := objc.Send[MTLMathMode](c.ID, objc.Sel("mathMode"))
	return MTLMathMode(rv)
}
func (c MTLCompileOptions) SetMathMode(value MTLMathMode) {
	objc.Send[struct{}](c.ID, objc.Sel("setMathMode:"), value)
}

// The FP32 math functions Metal uses.
//
// See: https://developer.apple.com/documentation/Metal/MTLCompileOptions/mathFloatingPointFunctions
func (c MTLCompileOptions) MathFloatingPointFunctions() MTLMathFloatingPointFunctions {
	rv := objc.Send[MTLMathFloatingPointFunctions](c.ID, objc.Sel("mathFloatingPointFunctions"))
	return MTLMathFloatingPointFunctions(rv)
}
func (c MTLCompileOptions) SetMathFloatingPointFunctions(value MTLMathFloatingPointFunctions) {
	objc.Send[struct{}](c.ID, objc.Sel("setMathFloatingPointFunctions:"), value)
}

// A Boolean value that indicates whether the compiler compiles vertex shaders
// conservatively to generate consistent position calculations.
//
// # Discussion
//
// The default value is false. When true, the Metal shader compiler looks at
// the position value in all vertex output structures that it compiles. If the
// position value also has the `[[invariant]]` attribute, the compiler
// compiles the corresponding vertex shader conservatively to guarantee that
// the GPU performs the calculations the same way. You need to preserve
// invariance when your renderer contains multiple render passes and requires
// the same position calculations in each render pass.
//
// See: https://developer.apple.com/documentation/Metal/MTLCompileOptions/preserveInvariance
func (c MTLCompileOptions) PreserveInvariance() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("preserveInvariance"))
	return rv
}
func (c MTLCompileOptions) SetPreserveInvariance(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setPreserveInvariance:"), value)
}

// The language version for interpreting the library source code.
//
// # Discussion
//
// By default, Metal uses the most recent language version.
//
// See: https://developer.apple.com/documentation/Metal/MTLCompileOptions/languageVersion
func (c MTLCompileOptions) LanguageVersion() MTLLanguageVersion {
	rv := objc.Send[MTLLanguageVersion](c.ID, objc.Sel("languageVersion"))
	return MTLLanguageVersion(rv)
}
func (c MTLCompileOptions) SetLanguageVersion(value MTLLanguageVersion) {
	objc.Send[struct{}](c.ID, objc.Sel("setLanguageVersion:"), value)
}

// A list of preprocessor macros to apply when compiling the library source.
//
// # Discussion
//
// Define the macros as a dictionary where each key is a string, and the
// values can be either an [NSString] or [NSNumber] instance.
//
// The default value is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLCompileOptions/preprocessorMacros
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
func (c MTLCompileOptions) PreprocessorMacros() foundation.INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("preprocessorMacros"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (c MTLCompileOptions) SetPreprocessorMacros(value foundation.INSDictionary) {
	objc.Send[struct{}](c.ID, objc.Sel("setPreprocessorMacros:"), value)
}

// An option that tells the compiler what to prioritize when it compiles Metal
// shader code.
//
// See: https://developer.apple.com/documentation/Metal/MTLCompileOptions/optimizationLevel
func (c MTLCompileOptions) OptimizationLevel() MTLLibraryOptimizationLevel {
	rv := objc.Send[MTLLibraryOptimizationLevel](c.ID, objc.Sel("optimizationLevel"))
	return MTLLibraryOptimizationLevel(rv)
}
func (c MTLCompileOptions) SetOptimizationLevel(value MTLLibraryOptimizationLevel) {
	objc.Send[struct{}](c.ID, objc.Sel("setOptimizationLevel:"), value)
}

// An array of dynamic libraries the Metal compiler links against.
//
// See: https://developer.apple.com/documentation/Metal/MTLCompileOptions/libraries
func (c MTLCompileOptions) Libraries() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("libraries"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
func (c MTLCompileOptions) SetLibraries(value []objectivec.IObject) {
	objc.Send[struct{}](c.ID, objc.Sel("setLibraries:"), objectivec.IObjectSliceToNSArray(value))
}

// A Boolean value that indicates whether the compiler can perform
// optimizations for floating-point arithmetic that may violate the IEEE 754
// standard.
//
// # Discussion
//
// The default value is true. A true value also enables the high-precision
// variant of math functions for single-precision floating-point scalar and
// vector types.
//
// See: https://developer.apple.com/documentation/Metal/MTLCompileOptions/fastMathEnabled
func (c MTLCompileOptions) FastMathEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("fastMathEnabled"))
	return rv
}
func (c MTLCompileOptions) SetFastMathEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setFastMathEnabled:"), value)
}

// The kind of library to create.
//
// # Discussion
//
// The default value is [MTLLibraryType.executable].
//
// See: https://developer.apple.com/documentation/Metal/MTLCompileOptions/libraryType
//
// [MTLLibraryType.executable]: https://developer.apple.com/documentation/Metal/MTLLibraryType/executable
func (c MTLCompileOptions) LibraryType() MTLLibraryType {
	rv := objc.Send[MTLLibraryType](c.ID, objc.Sel("libraryType"))
	return MTLLibraryType(rv)
}
func (c MTLCompileOptions) SetLibraryType(value MTLLibraryType) {
	objc.Send[struct{}](c.ID, objc.Sel("setLibraryType:"), value)
}

// For a dynamic library, the name to use when installing the library.
//
// See: https://developer.apple.com/documentation/Metal/MTLCompileOptions/installName
func (c MTLCompileOptions) InstallName() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("installName"))
	return foundation.NSStringFromID(rv).String()
}
func (c MTLCompileOptions) SetInstallName(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setInstallName:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/Metal/MTLCompileOptions/allowReferencingUndefinedSymbols
func (c MTLCompileOptions) AllowReferencingUndefinedSymbols() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("allowReferencingUndefinedSymbols"))
	return rv
}
func (c MTLCompileOptions) SetAllowReferencingUndefinedSymbols(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setAllowReferencingUndefinedSymbols:"), value)
}

// See: https://developer.apple.com/documentation/Metal/MTLCompileOptions/compileSymbolVisibility
func (c MTLCompileOptions) CompileSymbolVisibility() MTLCompileSymbolVisibility {
	rv := objc.Send[MTLCompileSymbolVisibility](c.ID, objc.Sel("compileSymbolVisibility"))
	return MTLCompileSymbolVisibility(rv)
}
func (c MTLCompileOptions) SetCompileSymbolVisibility(value MTLCompileSymbolVisibility) {
	objc.Send[struct{}](c.ID, objc.Sel("setCompileSymbolVisibility:"), value)
}

// See: https://developer.apple.com/documentation/Metal/MTLCompileOptions/maxTotalThreadsPerThreadgroup
func (c MTLCompileOptions) MaxTotalThreadsPerThreadgroup() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("maxTotalThreadsPerThreadgroup"))
	return rv
}
func (c MTLCompileOptions) SetMaxTotalThreadsPerThreadgroup(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setMaxTotalThreadsPerThreadgroup:"), value)
}

// # Discussion
//
// Sets the required threads-per-threadgroup during dispatches. The
// `threadsPerThreadgroup` argument of any dispatch must match this value if
// it is set. Optional, unless the pipeline is going to use CooperativeTensors
// in which case this must be set. Setting this to a size of 0 in every
// dimension disables this property
//
// See: https://developer.apple.com/documentation/Metal/MTLCompileOptions/requiredThreadsPerThreadgroup
func (c MTLCompileOptions) RequiredThreadsPerThreadgroup() MTLSize {
	rv := objc.Send[MTLSize](c.ID, objc.Sel("requiredThreadsPerThreadgroup"))
	return MTLSize(rv)
}
func (c MTLCompileOptions) SetRequiredThreadsPerThreadgroup(value MTLSize) {
	objc.Send[struct{}](c.ID, objc.Sel("setRequiredThreadsPerThreadgroup:"), value)
}
