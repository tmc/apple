// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLCompilerOptions] class.
var (
	_MLCompilerOptionsClass     MLCompilerOptionsClass
	_MLCompilerOptionsClassOnce sync.Once
)

func getMLCompilerOptionsClass() MLCompilerOptionsClass {
	_MLCompilerOptionsClassOnce.Do(func() {
		_MLCompilerOptionsClass = MLCompilerOptionsClass{class: objc.GetClass("MLCompilerOptions")}
	})
	return _MLCompilerOptionsClass
}

// GetMLCompilerOptionsClass returns the class object for MLCompilerOptions.
func GetMLCompilerOptionsClass() MLCompilerOptionsClass {
	return getMLCompilerOptionsClass()
}

type MLCompilerOptionsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLCompilerOptionsClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLCompilerOptionsClass) Alloc() MLCompilerOptions {
	rv := objc.Send[MLCompilerOptions](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLCompilerOptions.AllowMultipleInputsWithEnumeratedShapes]
//   - [MLCompilerOptions.SetAllowMultipleInputsWithEnumeratedShapes]
//   - [MLCompilerOptions.AllowsPixelBufferDirectBinding]
//   - [MLCompilerOptions.SetAllowsPixelBufferDirectBinding]
//   - [MLCompilerOptions.ContainerIsCloud]
//   - [MLCompilerOptions.SetContainerIsCloud]
//   - [MLCompilerOptions.DryRun]
//   - [MLCompilerOptions.SetDryRun]
//   - [MLCompilerOptions.EncryptModel]
//   - [MLCompilerOptions.SetEncryptModel]
//   - [MLCompilerOptions.Iv]
//   - [MLCompilerOptions.SetIv]
//   - [MLCompilerOptions.Key]
//   - [MLCompilerOptions.SetKey]
//   - [MLCompilerOptions.KeyID]
//   - [MLCompilerOptions.SetKeyID]
//   - [MLCompilerOptions.KeyInfoVersion]
//   - [MLCompilerOptions.SetKeyInfoVersion]
//   - [MLCompilerOptions.MlProgramAddDuringCompilationMode]
//   - [MLCompilerOptions.SetMlProgramAddDuringCompilationMode]
//   - [MLCompilerOptions.Mlsinf]
//   - [MLCompilerOptions.SetMlsinf]
//   - [MLCompilerOptions.Platform]
//   - [MLCompilerOptions.SetPlatform]
//   - [MLCompilerOptions.PlatformVersion]
//   - [MLCompilerOptions.SetPlatformVersion]
//   - [MLCompilerOptions.Sinf]
//   - [MLCompilerOptions.SetSinf]
//   - [MLCompilerOptions.SpecURL]
//   - [MLCompilerOptions.SetSpecURL]
//   - [MLCompilerOptions.TrainWithMLCompute]
//   - [MLCompilerOptions.SetTrainWithMLCompute]
//   - [MLCompilerOptions.UsesCodeSigningIdentityForEncryption]
//   - [MLCompilerOptions.SetUsesCodeSigningIdentityForEncryption]
//   - [MLCompilerOptions.Warnings]
//   - [MLCompilerOptions.SetWarnings]
//
// See: https://developer.apple.com/documentation/CoreML/MLCompilerOptions
type MLCompilerOptions struct {
	objectivec.Object
}

// MLCompilerOptionsFromID constructs a [MLCompilerOptions] from an objc.ID.
func MLCompilerOptionsFromID(id objc.ID) MLCompilerOptions {
	return MLCompilerOptions{objectivec.Object{ID: id}}
}

// Ensure MLCompilerOptions implements IMLCompilerOptions.
var _ IMLCompilerOptions = MLCompilerOptions{}

// An interface definition for the [MLCompilerOptions] class.
//
// # Methods
//
//   - [IMLCompilerOptions.AllowMultipleInputsWithEnumeratedShapes]
//   - [IMLCompilerOptions.SetAllowMultipleInputsWithEnumeratedShapes]
//   - [IMLCompilerOptions.AllowsPixelBufferDirectBinding]
//   - [IMLCompilerOptions.SetAllowsPixelBufferDirectBinding]
//   - [IMLCompilerOptions.ContainerIsCloud]
//   - [IMLCompilerOptions.SetContainerIsCloud]
//   - [IMLCompilerOptions.DryRun]
//   - [IMLCompilerOptions.SetDryRun]
//   - [IMLCompilerOptions.EncryptModel]
//   - [IMLCompilerOptions.SetEncryptModel]
//   - [IMLCompilerOptions.Iv]
//   - [IMLCompilerOptions.SetIv]
//   - [IMLCompilerOptions.Key]
//   - [IMLCompilerOptions.SetKey]
//   - [IMLCompilerOptions.KeyID]
//   - [IMLCompilerOptions.SetKeyID]
//   - [IMLCompilerOptions.KeyInfoVersion]
//   - [IMLCompilerOptions.SetKeyInfoVersion]
//   - [IMLCompilerOptions.MlProgramAddDuringCompilationMode]
//   - [IMLCompilerOptions.SetMlProgramAddDuringCompilationMode]
//   - [IMLCompilerOptions.Mlsinf]
//   - [IMLCompilerOptions.SetMlsinf]
//   - [IMLCompilerOptions.Platform]
//   - [IMLCompilerOptions.SetPlatform]
//   - [IMLCompilerOptions.PlatformVersion]
//   - [IMLCompilerOptions.SetPlatformVersion]
//   - [IMLCompilerOptions.Sinf]
//   - [IMLCompilerOptions.SetSinf]
//   - [IMLCompilerOptions.SpecURL]
//   - [IMLCompilerOptions.SetSpecURL]
//   - [IMLCompilerOptions.TrainWithMLCompute]
//   - [IMLCompilerOptions.SetTrainWithMLCompute]
//   - [IMLCompilerOptions.UsesCodeSigningIdentityForEncryption]
//   - [IMLCompilerOptions.SetUsesCodeSigningIdentityForEncryption]
//   - [IMLCompilerOptions.Warnings]
//   - [IMLCompilerOptions.SetWarnings]
//
// See: https://developer.apple.com/documentation/CoreML/MLCompilerOptions
type IMLCompilerOptions interface {
	objectivec.IObject

	// Topic: Methods

	AllowMultipleInputsWithEnumeratedShapes() bool
	SetAllowMultipleInputsWithEnumeratedShapes(value bool)
	AllowsPixelBufferDirectBinding() bool
	SetAllowsPixelBufferDirectBinding(value bool)
	ContainerIsCloud() bool
	SetContainerIsCloud(value bool)
	DryRun() bool
	SetDryRun(value bool)
	EncryptModel() bool
	SetEncryptModel(value bool)
	Iv() foundation.INSData
	SetIv(value foundation.INSData)
	Key() foundation.INSData
	SetKey(value foundation.INSData)
	KeyID() string
	SetKeyID(value string)
	KeyInfoVersion() foundation.NSNumber
	SetKeyInfoVersion(value foundation.NSNumber)
	MlProgramAddDuringCompilationMode() int
	SetMlProgramAddDuringCompilationMode(value int)
	Mlsinf() foundation.INSData
	SetMlsinf(value foundation.INSData)
	Platform() string
	SetPlatform(value string)
	PlatformVersion() string
	SetPlatformVersion(value string)
	Sinf() foundation.INSData
	SetSinf(value foundation.INSData)
	SpecURL() foundation.INSURL
	SetSpecURL(value foundation.INSURL)
	TrainWithMLCompute() bool
	SetTrainWithMLCompute(value bool)
	UsesCodeSigningIdentityForEncryption() bool
	SetUsesCodeSigningIdentityForEncryption(value bool)
	Warnings() foundation.INSArray
	SetWarnings(value foundation.INSArray)
}

// Init initializes the instance.
func (c MLCompilerOptions) Init() MLCompilerOptions {
	rv := objc.Send[MLCompilerOptions](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MLCompilerOptions) Autorelease() MLCompilerOptions {
	rv := objc.Send[MLCompilerOptions](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLCompilerOptions creates a new MLCompilerOptions instance.
func NewMLCompilerOptions() MLCompilerOptions {
	class := getMLCompilerOptionsClass()
	rv := objc.Send[MLCompilerOptions](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLCompilerOptions/defaultOptions
func (_MLCompilerOptionsClass MLCompilerOptionsClass) DefaultOptions() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLCompilerOptionsClass.class), objc.Sel("defaultOptions"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLCompilerOptions/allowMultipleInputsWithEnumeratedShapes
func (c MLCompilerOptions) AllowMultipleInputsWithEnumeratedShapes() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("allowMultipleInputsWithEnumeratedShapes"))
	return rv
}
func (c MLCompilerOptions) SetAllowMultipleInputsWithEnumeratedShapes(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setAllowMultipleInputsWithEnumeratedShapes:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLCompilerOptions/allowsPixelBufferDirectBinding
func (c MLCompilerOptions) AllowsPixelBufferDirectBinding() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("allowsPixelBufferDirectBinding"))
	return rv
}
func (c MLCompilerOptions) SetAllowsPixelBufferDirectBinding(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setAllowsPixelBufferDirectBinding:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLCompilerOptions/containerIsCloud
func (c MLCompilerOptions) ContainerIsCloud() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("containerIsCloud"))
	return rv
}
func (c MLCompilerOptions) SetContainerIsCloud(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setContainerIsCloud:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLCompilerOptions/dryRun
func (c MLCompilerOptions) DryRun() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("dryRun"))
	return rv
}
func (c MLCompilerOptions) SetDryRun(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setDryRun:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLCompilerOptions/encryptModel
func (c MLCompilerOptions) EncryptModel() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("encryptModel"))
	return rv
}
func (c MLCompilerOptions) SetEncryptModel(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setEncryptModel:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLCompilerOptions/iv
func (c MLCompilerOptions) Iv() foundation.INSData {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("iv"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (c MLCompilerOptions) SetIv(value foundation.INSData) {
	objc.Send[struct{}](c.ID, objc.Sel("setIv:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLCompilerOptions/key
func (c MLCompilerOptions) Key() foundation.INSData {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("key"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (c MLCompilerOptions) SetKey(value foundation.INSData) {
	objc.Send[struct{}](c.ID, objc.Sel("setKey:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLCompilerOptions/keyID
func (c MLCompilerOptions) KeyID() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("keyID"))
	return foundation.NSStringFromID(rv).String()
}
func (c MLCompilerOptions) SetKeyID(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setKeyID:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLCompilerOptions/keyInfoVersion
func (c MLCompilerOptions) KeyInfoVersion() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("keyInfoVersion"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c MLCompilerOptions) SetKeyInfoVersion(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setKeyInfoVersion:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLCompilerOptions/mlProgramAddDuringCompilationMode
func (c MLCompilerOptions) MlProgramAddDuringCompilationMode() int {
	rv := objc.Send[int](c.ID, objc.Sel("mlProgramAddDuringCompilationMode"))
	return rv
}
func (c MLCompilerOptions) SetMlProgramAddDuringCompilationMode(value int) {
	objc.Send[struct{}](c.ID, objc.Sel("setMlProgramAddDuringCompilationMode:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLCompilerOptions/mlsinf
func (c MLCompilerOptions) Mlsinf() foundation.INSData {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("mlsinf"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (c MLCompilerOptions) SetMlsinf(value foundation.INSData) {
	objc.Send[struct{}](c.ID, objc.Sel("setMlsinf:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLCompilerOptions/platform
func (c MLCompilerOptions) Platform() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("platform"))
	return foundation.NSStringFromID(rv).String()
}
func (c MLCompilerOptions) SetPlatform(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setPlatform:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLCompilerOptions/platformVersion
func (c MLCompilerOptions) PlatformVersion() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("platformVersion"))
	return foundation.NSStringFromID(rv).String()
}
func (c MLCompilerOptions) SetPlatformVersion(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setPlatformVersion:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLCompilerOptions/sinf
func (c MLCompilerOptions) Sinf() foundation.INSData {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("sinf"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (c MLCompilerOptions) SetSinf(value foundation.INSData) {
	objc.Send[struct{}](c.ID, objc.Sel("setSinf:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLCompilerOptions/specURL
func (c MLCompilerOptions) SpecURL() foundation.INSURL {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("specURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (c MLCompilerOptions) SetSpecURL(value foundation.INSURL) {
	objc.Send[struct{}](c.ID, objc.Sel("setSpecURL:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLCompilerOptions/trainWithMLCompute
func (c MLCompilerOptions) TrainWithMLCompute() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("trainWithMLCompute"))
	return rv
}
func (c MLCompilerOptions) SetTrainWithMLCompute(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setTrainWithMLCompute:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLCompilerOptions/usesCodeSigningIdentityForEncryption
func (c MLCompilerOptions) UsesCodeSigningIdentityForEncryption() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("usesCodeSigningIdentityForEncryption"))
	return rv
}
func (c MLCompilerOptions) SetUsesCodeSigningIdentityForEncryption(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setUsesCodeSigningIdentityForEncryption:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLCompilerOptions/warnings
func (c MLCompilerOptions) Warnings() foundation.INSArray {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("warnings"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (c MLCompilerOptions) SetWarnings(value foundation.INSArray) {
	objc.Send[struct{}](c.ID, objc.Sel("setWarnings:"), value)
}
