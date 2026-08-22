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
	rv := objc.SendIfResponds[MLCompilerOptions](objc.ID(mc.class), objc.Sel("alloc"))
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
	Iv() foundation.NSData
	SetIv(value foundation.NSData)
	Key() foundation.NSData
	SetKey(value foundation.NSData)
	KeyID() string
	SetKeyID(value string)
	KeyInfoVersion() foundation.NSNumber
	SetKeyInfoVersion(value foundation.NSNumber)
	MlProgramAddDuringCompilationMode() int
	SetMlProgramAddDuringCompilationMode(value int)
	Mlsinf() foundation.NSData
	SetMlsinf(value foundation.NSData)
	Platform() string
	SetPlatform(value string)
	PlatformVersion() string
	SetPlatformVersion(value string)
	Sinf() foundation.NSData
	SetSinf(value foundation.NSData)
	SpecURL() foundation.NSURL
	SetSpecURL(value foundation.NSURL)
	TrainWithMLCompute() bool
	SetTrainWithMLCompute(value bool)
	UsesCodeSigningIdentityForEncryption() bool
	SetUsesCodeSigningIdentityForEncryption(value bool)
	Warnings() foundation.INSArray
	SetWarnings(value foundation.INSArray)
}

// Init initializes the instance.
func (m MLCompilerOptions) Init() MLCompilerOptions {
	rv := objc.SendIfResponds[MLCompilerOptions](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLCompilerOptions) Autorelease() MLCompilerOptions {
	rv := objc.SendIfResponds[MLCompilerOptions](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLCompilerOptions creates a new MLCompilerOptions instance.
func NewMLCompilerOptions() MLCompilerOptions {
	class := getMLCompilerOptionsClass()
	rv := objc.SendIfResponds[MLCompilerOptions](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (_MLCompilerOptionsClass MLCompilerOptionsClass) DefaultOptions() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLCompilerOptionsClass.class), objc.Sel("defaultOptions"))
	return objectivec.Object{ID: rv}
}

func (m MLCompilerOptions) AllowMultipleInputsWithEnumeratedShapes() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("allowMultipleInputsWithEnumeratedShapes"))
	return rv
}
func (m MLCompilerOptions) SetAllowMultipleInputsWithEnumeratedShapes(value bool) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setAllowMultipleInputsWithEnumeratedShapes:"), value)
}
func (m MLCompilerOptions) AllowsPixelBufferDirectBinding() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("allowsPixelBufferDirectBinding"))
	return rv
}
func (m MLCompilerOptions) SetAllowsPixelBufferDirectBinding(value bool) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setAllowsPixelBufferDirectBinding:"), value)
}
func (m MLCompilerOptions) ContainerIsCloud() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("containerIsCloud"))
	return rv
}
func (m MLCompilerOptions) SetContainerIsCloud(value bool) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setContainerIsCloud:"), value)
}
func (m MLCompilerOptions) DryRun() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("dryRun"))
	return rv
}
func (m MLCompilerOptions) SetDryRun(value bool) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setDryRun:"), value)
}
func (m MLCompilerOptions) EncryptModel() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("encryptModel"))
	return rv
}
func (m MLCompilerOptions) SetEncryptModel(value bool) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setEncryptModel:"), value)
}
func (m MLCompilerOptions) Iv() foundation.NSData {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("iv"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (m MLCompilerOptions) SetIv(value foundation.NSData) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setIv:"), value)
}
func (m MLCompilerOptions) Key() foundation.NSData {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("key"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (m MLCompilerOptions) SetKey(value foundation.NSData) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setKey:"), value)
}
func (m MLCompilerOptions) KeyID() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("keyID"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLCompilerOptions) SetKeyID(value string) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setKeyID:"), objc.String(value))
}
func (m MLCompilerOptions) KeyInfoVersion() foundation.NSNumber {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("keyInfoVersion"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (m MLCompilerOptions) SetKeyInfoVersion(value foundation.NSNumber) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setKeyInfoVersion:"), value)
}
func (m MLCompilerOptions) MlProgramAddDuringCompilationMode() int {
	rv := objc.SendIfResponds[int](m.ID, objc.Sel("mlProgramAddDuringCompilationMode"))
	return rv
}
func (m MLCompilerOptions) SetMlProgramAddDuringCompilationMode(value int) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setMlProgramAddDuringCompilationMode:"), value)
}
func (m MLCompilerOptions) Mlsinf() foundation.NSData {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("mlsinf"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (m MLCompilerOptions) SetMlsinf(value foundation.NSData) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setMlsinf:"), value)
}
func (m MLCompilerOptions) Platform() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("platform"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLCompilerOptions) SetPlatform(value string) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setPlatform:"), objc.String(value))
}
func (m MLCompilerOptions) PlatformVersion() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("platformVersion"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLCompilerOptions) SetPlatformVersion(value string) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setPlatformVersion:"), objc.String(value))
}
func (m MLCompilerOptions) Sinf() foundation.NSData {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("sinf"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (m MLCompilerOptions) SetSinf(value foundation.NSData) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setSinf:"), value)
}
func (m MLCompilerOptions) SpecURL() foundation.NSURL {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("specURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (m MLCompilerOptions) SetSpecURL(value foundation.NSURL) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setSpecURL:"), value)
}
func (m MLCompilerOptions) TrainWithMLCompute() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("trainWithMLCompute"))
	return rv
}
func (m MLCompilerOptions) SetTrainWithMLCompute(value bool) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setTrainWithMLCompute:"), value)
}
func (m MLCompilerOptions) UsesCodeSigningIdentityForEncryption() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("usesCodeSigningIdentityForEncryption"))
	return rv
}
func (m MLCompilerOptions) SetUsesCodeSigningIdentityForEncryption(value bool) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setUsesCodeSigningIdentityForEncryption:"), value)
}
func (m MLCompilerOptions) Warnings() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("warnings"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLCompilerOptions) SetWarnings(value foundation.INSArray) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setWarnings:"), value)
}
