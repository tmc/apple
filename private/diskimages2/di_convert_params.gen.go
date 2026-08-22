// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"context"
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DIConvertParams] class.
var (
	_DIConvertParamsClass     DIConvertParamsClass
	_DIConvertParamsClassOnce sync.Once
)

func getDIConvertParamsClass() DIConvertParamsClass {
	_DIConvertParamsClassOnce.Do(func() {
		_DIConvertParamsClass = DIConvertParamsClass{class: objc.GetClass("DIConvertParams")}
	})
	return _DIConvertParamsClass
}

// GetDIConvertParamsClass returns the class object for DIConvertParams.
func GetDIConvertParamsClass() DIConvertParamsClass {
	return getDIConvertParamsClass()
}

type DIConvertParamsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIConvertParamsClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIConvertParamsClass) Alloc() DIConvertParams {
	rv := objc.SendIfResponds[DIConvertParams](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [DIConvertParams.Certificate]
//   - [DIConvertParams.SetCertificate]
//   - [DIConvertParams.ConversionMethod]
//   - [DIConvertParams.SetConversionMethod]
//   - [DIConvertParams.ConvertWithCompletionBlock]
//   - [DIConvertParams.CopyUpdatedOutputURLWithError]
//   - [DIConvertParams.EncryptionMethod]
//   - [DIConvertParams.SetEncryptionMethod]
//   - [DIConvertParams.InPlaceConversion]
//   - [DIConvertParams.IsInputURLDevice]
//   - [DIConvertParams.MaxRawUDIFRunSize]
//   - [DIConvertParams.SetMaxRawUDIFRunSize]
//   - [DIConvertParams.OnConvertCompletionWithInErrorOutError]
//   - [DIConvertParams.OpenDeviceAsRootWithError]
//   - [DIConvertParams.OutputFormat]
//   - [DIConvertParams.SetOutputFormat]
//   - [DIConvertParams.OutputParams]
//   - [DIConvertParams.SetOutputParams]
//   - [DIConvertParams.OutputURL]
//   - [DIConvertParams.SetOutputURL]
//   - [DIConvertParams.Passphrase]
//   - [DIConvertParams.SetPassphrase]
//   - [DIConvertParams.PrepareConvertWithError]
//   - [DIConvertParams.PrepareParamsForSquashWithError]
//   - [DIConvertParams.PrepareParamsWithError]
//   - [DIConvertParams.PublicKey]
//   - [DIConvertParams.SetPublicKey]
//   - [DIConvertParams.SetPassphraseEncryptionMethodError]
//   - [DIConvertParams.SetPassphraseError]
//   - [DIConvertParams.ShadowURLs]
//   - [DIConvertParams.ShouldPerformInplaceSquash]
//   - [DIConvertParams.ShouldValidateShadows]
//   - [DIConvertParams.SetShouldValidateShadows]
//   - [DIConvertParams.SparseBundleBandSize]
//   - [DIConvertParams.SetSparseBundleBandSize]
//   - [DIConvertParams.TemporaryPassphrase]
//   - [DIConvertParams.UseFormatMappingInfo]
//   - [DIConvertParams.SetUseFormatMappingInfo]
//   - [DIConvertParams.ValidateFileWithURLError]
//   - [DIConvertParams.ValidateSquashFormats]
//   - [DIConvertParams.InitForInplaceWithExistingParamsError]
//   - [DIConvertParams.InitForInplaceWithURLError]
//   - [DIConvertParams.InitWithInputURLOutputURLError]
//   - [DIConvertParams.InitWithInputURLOutputURLShadowURLsError]
type DIConvertParams struct {
	DIBaseParams
}

// DIConvertParamsFromID constructs a [DIConvertParams] from an objc.ID.
func DIConvertParamsFromID(id objc.ID) DIConvertParams {
	return DIConvertParams{DIBaseParams: DIBaseParamsFromID(id)}
}

// Ensure DIConvertParams implements IDIConvertParams.
var _ IDIConvertParams = DIConvertParams{}

// An interface definition for the [DIConvertParams] class.
//
// # Methods
//
//   - [IDIConvertParams.Certificate]
//   - [IDIConvertParams.SetCertificate]
//   - [IDIConvertParams.ConversionMethod]
//   - [IDIConvertParams.SetConversionMethod]
//   - [IDIConvertParams.ConvertWithCompletionBlock]
//   - [IDIConvertParams.CopyUpdatedOutputURLWithError]
//   - [IDIConvertParams.EncryptionMethod]
//   - [IDIConvertParams.SetEncryptionMethod]
//   - [IDIConvertParams.InPlaceConversion]
//   - [IDIConvertParams.IsInputURLDevice]
//   - [IDIConvertParams.MaxRawUDIFRunSize]
//   - [IDIConvertParams.SetMaxRawUDIFRunSize]
//   - [IDIConvertParams.OnConvertCompletionWithInErrorOutError]
//   - [IDIConvertParams.OpenDeviceAsRootWithError]
//   - [IDIConvertParams.OutputFormat]
//   - [IDIConvertParams.SetOutputFormat]
//   - [IDIConvertParams.OutputParams]
//   - [IDIConvertParams.SetOutputParams]
//   - [IDIConvertParams.OutputURL]
//   - [IDIConvertParams.SetOutputURL]
//   - [IDIConvertParams.Passphrase]
//   - [IDIConvertParams.SetPassphrase]
//   - [IDIConvertParams.PrepareConvertWithError]
//   - [IDIConvertParams.PrepareParamsForSquashWithError]
//   - [IDIConvertParams.PrepareParamsWithError]
//   - [IDIConvertParams.PublicKey]
//   - [IDIConvertParams.SetPublicKey]
//   - [IDIConvertParams.SetPassphraseEncryptionMethodError]
//   - [IDIConvertParams.SetPassphraseError]
//   - [IDIConvertParams.ShadowURLs]
//   - [IDIConvertParams.ShouldPerformInplaceSquash]
//   - [IDIConvertParams.ShouldValidateShadows]
//   - [IDIConvertParams.SetShouldValidateShadows]
//   - [IDIConvertParams.SparseBundleBandSize]
//   - [IDIConvertParams.SetSparseBundleBandSize]
//   - [IDIConvertParams.TemporaryPassphrase]
//   - [IDIConvertParams.UseFormatMappingInfo]
//   - [IDIConvertParams.SetUseFormatMappingInfo]
//   - [IDIConvertParams.ValidateFileWithURLError]
//   - [IDIConvertParams.ValidateSquashFormats]
//   - [IDIConvertParams.InitForInplaceWithExistingParamsError]
//   - [IDIConvertParams.InitForInplaceWithURLError]
//   - [IDIConvertParams.InitWithInputURLOutputURLError]
//   - [IDIConvertParams.InitWithInputURLOutputURLShadowURLsError]
type IDIConvertParams interface {
	IDIBaseParams

	// Topic: Methods

	Certificate() string
	SetCertificate(value string)
	ConversionMethod() uint64
	SetConversionMethod(value uint64)
	ConvertWithCompletionBlock(block VoidHandler) objectivec.IObject
	CopyUpdatedOutputURLWithError() (objectivec.IObject, error)
	EncryptionMethod() uint64
	SetEncryptionMethod(value uint64)
	InPlaceConversion() bool
	IsInputURLDevice() bool
	MaxRawUDIFRunSize() uint64
	SetMaxRawUDIFRunSize(value uint64)
	OnConvertCompletionWithInErrorOutError(error_ objectivec.IObject) (bool, error)
	OpenDeviceAsRootWithError() (objectivec.IObject, error)
	OutputFormat() int64
	SetOutputFormat(value int64)
	OutputParams() IDIBaseParams
	SetOutputParams(value IDIBaseParams)
	OutputURL() IDIURL
	SetOutputURL(value IDIURL)
	Passphrase() bool
	SetPassphrase(value bool)
	PrepareConvertWithError() (objectivec.IObject, error)
	PrepareParamsForSquashWithError() (bool, error)
	PrepareParamsWithError() (bool, error)
	PublicKey() string
	SetPublicKey(value string)
	SetPassphraseEncryptionMethodError(passphrase string, method uint64) (bool, error)
	SetPassphraseError(passphrase string) (bool, error)
	ShadowURLs() foundation.INSArray
	ShouldPerformInplaceSquash() bool
	ShouldValidateShadows() bool
	SetShouldValidateShadows(value bool)
	SparseBundleBandSize() uint64
	SetSparseBundleBandSize(value uint64)
	TemporaryPassphrase() IDITemporaryPassphrase
	UseFormatMappingInfo() bool
	SetUseFormatMappingInfo(value bool)
	ValidateFileWithURLError(url foundation.NSURL) (bool, error)
	ValidateSquashFormats() bool
	InitForInplaceWithExistingParamsError(params objectivec.IObject) (DIConvertParams, error)
	InitForInplaceWithURLError(url foundation.NSURL) (DIConvertParams, error)
	InitWithInputURLOutputURLError(url foundation.NSURL, url2 foundation.NSURL) (DIConvertParams, error)
	InitWithInputURLOutputURLShadowURLsError(url foundation.NSURL, url2 foundation.NSURL, uRLs objectivec.IObject) (DIConvertParams, error)
}

// Init initializes the instance.
func (d DIConvertParams) Init() DIConvertParams {
	rv := objc.SendIfResponds[DIConvertParams](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIConvertParams) Autorelease() DIConvertParams {
	rv := objc.SendIfResponds[DIConvertParams](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIConvertParams creates a new DIConvertParams instance.
func NewDIConvertParams() DIConvertParams {
	class := getDIConvertParamsClass()
	rv := objc.SendIfResponds[DIConvertParams](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewDIConvertParamsForInplaceWithExistingParamsError(params objectivec.IObject) (DIConvertParams, error) {
	var errorPtr objc.ID
	instance := getDIConvertParamsClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initForInplaceWithExistingParams:error:"), params, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIConvertParams{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return DIConvertParams{}, objc.ErrInitFailed
	}
	return DIConvertParamsFromID(rv), nil
}

func NewDIConvertParamsForInplaceWithURLError(url foundation.NSURL) (DIConvertParams, error) {
	var errorPtr objc.ID
	instance := getDIConvertParamsClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initForInplaceWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIConvertParams{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return DIConvertParams{}, objc.ErrInitFailed
	}
	return DIConvertParamsFromID(rv), nil
}

func NewDIConvertParamsWithCoder(coder objectivec.IObject) DIConvertParams {
	instance := getDIConvertParamsClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DIConvertParamsFromID(rv)
}

func NewDIConvertParamsWithInputURLOutputURLError(url foundation.NSURL, url2 foundation.NSURL) (DIConvertParams, error) {
	var errorPtr objc.ID
	instance := getDIConvertParamsClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithInputURL:outputURL:error:"), url, url2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIConvertParams{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return DIConvertParams{}, objc.ErrInitFailed
	}
	return DIConvertParamsFromID(rv), nil
}

func NewDIConvertParamsWithInputURLOutputURLShadowURLsError(url foundation.NSURL, url2 foundation.NSURL, uRLs objectivec.IObject) (DIConvertParams, error) {
	var errorPtr objc.ID
	instance := getDIConvertParamsClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithInputURL:outputURL:shadowURLs:error:"), url, url2, uRLs, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIConvertParams{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return DIConvertParams{}, objc.ErrInitFailed
	}
	return DIConvertParamsFromID(rv), nil
}

func NewDIConvertParamsWithURLError(url foundation.NSURL) (DIConvertParams, error) {
	var errorPtr objc.ID
	instance := getDIConvertParamsClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIConvertParams{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return DIConvertParams{}, objc.ErrInitFailed
	}
	return DIConvertParamsFromID(rv), nil
}

func (d DIConvertParams) ConvertWithCompletionBlock(block VoidHandler) objectivec.IObject {
	_block0, _ := NewVoidBlock(block)
	rv := objc.SendIfResponds[objc.ID](d.ID, objc.Sel("convertWithCompletionBlock:"), _block0)
	return objectivec.Object{ID: rv}
}
func (d DIConvertParams) CopyUpdatedOutputURLWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("copyUpdatedOutputURLWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (d DIConvertParams) IsInputURLDevice() bool {
	rv := objc.SendIfResponds[bool](d.ID, objc.Sel("isInputURLDevice"))
	return rv
}
func (d DIConvertParams) OnConvertCompletionWithInErrorOutError(error_ objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("onConvertCompletionWithInError:outError:"), error_, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("onConvertCompletionWithInError:outError: returned NO with nil NSError")
	}
	return rv, nil

}
func (d DIConvertParams) OpenDeviceAsRootWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("openDeviceAsRootWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (d DIConvertParams) PrepareConvertWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("prepareConvertWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (d DIConvertParams) PrepareParamsForSquashWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("prepareParamsForSquashWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("prepareParamsForSquashWithError: returned NO with nil NSError")
	}
	return rv, nil

}
func (d DIConvertParams) PrepareParamsWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("prepareParamsWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("prepareParamsWithError: returned NO with nil NSError")
	}
	return rv, nil

}
func (d DIConvertParams) SetPassphraseEncryptionMethodError(passphrase string, method uint64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("setPassphrase:encryptionMethod:error:"), unsafe.Pointer(unsafe.StringData(passphrase+"\x00")), method, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setPassphrase:encryptionMethod:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (d DIConvertParams) SetPassphraseError(passphrase string) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("setPassphrase:error:"), unsafe.Pointer(unsafe.StringData(passphrase+"\x00")), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setPassphrase:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (d DIConvertParams) ShouldPerformInplaceSquash() bool {
	rv := objc.SendIfResponds[bool](d.ID, objc.Sel("shouldPerformInplaceSquash"))
	return rv
}
func (d DIConvertParams) ValidateFileWithURLError(url foundation.NSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("validateFileWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("validateFileWithURL:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (d DIConvertParams) ValidateSquashFormats() bool {
	rv := objc.SendIfResponds[bool](d.ID, objc.Sel("validateSquashFormats"))
	return rv
}
func (d DIConvertParams) InitForInplaceWithExistingParamsError(params objectivec.IObject) (DIConvertParams, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initForInplaceWithExistingParams:error:"), params, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIConvertParams{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIConvertParamsFromID(rv), nil

}
func (d DIConvertParams) InitForInplaceWithURLError(url foundation.NSURL) (DIConvertParams, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initForInplaceWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIConvertParams{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIConvertParamsFromID(rv), nil

}
func (d DIConvertParams) InitWithInputURLOutputURLError(url foundation.NSURL, url2 foundation.NSURL) (DIConvertParams, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithInputURL:outputURL:error:"), url, url2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIConvertParams{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIConvertParamsFromID(rv), nil

}
func (d DIConvertParams) InitWithInputURLOutputURLShadowURLsError(url foundation.NSURL, url2 foundation.NSURL, uRLs objectivec.IObject) (DIConvertParams, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithInputURL:outputURL:shadowURLs:error:"), url, url2, uRLs, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIConvertParams{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIConvertParamsFromID(rv), nil

}

func (d DIConvertParams) Certificate() string {
	rv := objc.SendIfResponds[objc.ID](d.ID, objc.Sel("certificate"))
	return foundation.NSStringFromID(rv).String()
}
func (d DIConvertParams) SetCertificate(value string) {
	objc.SendIfResponds[struct{}](d.ID, objc.Sel("setCertificate:"), objc.String(value))
}
func (d DIConvertParams) ConversionMethod() uint64 {
	rv := objc.SendIfResponds[uint64](d.ID, objc.Sel("conversionMethod"))
	return rv
}
func (d DIConvertParams) SetConversionMethod(value uint64) {
	objc.SendIfResponds[struct{}](d.ID, objc.Sel("setConversionMethod:"), value)
}
func (d DIConvertParams) EncryptionMethod() uint64 {
	rv := objc.SendIfResponds[uint64](d.ID, objc.Sel("encryptionMethod"))
	return rv
}
func (d DIConvertParams) SetEncryptionMethod(value uint64) {
	objc.SendIfResponds[struct{}](d.ID, objc.Sel("setEncryptionMethod:"), value)
}
func (d DIConvertParams) InPlaceConversion() bool {
	rv := objc.SendIfResponds[bool](d.ID, objc.Sel("inPlaceConversion"))
	return rv
}
func (d DIConvertParams) MaxRawUDIFRunSize() uint64 {
	rv := objc.SendIfResponds[uint64](d.ID, objc.Sel("maxRawUDIFRunSize"))
	return rv
}
func (d DIConvertParams) SetMaxRawUDIFRunSize(value uint64) {
	objc.SendIfResponds[struct{}](d.ID, objc.Sel("setMaxRawUDIFRunSize:"), value)
}
func (d DIConvertParams) OutputFormat() int64 {
	rv := objc.SendIfResponds[int64](d.ID, objc.Sel("outputFormat"))
	return rv
}
func (d DIConvertParams) SetOutputFormat(value int64) {
	objc.SendIfResponds[struct{}](d.ID, objc.Sel("setOutputFormat:"), value)
}
func (d DIConvertParams) OutputParams() IDIBaseParams {
	rv := objc.SendIfResponds[objc.ID](d.ID, objc.Sel("outputParams"))
	return DIBaseParamsFromID(objc.ID(rv))
}
func (d DIConvertParams) SetOutputParams(value IDIBaseParams) {
	objc.SendIfResponds[struct{}](d.ID, objc.Sel("setOutputParams:"), value)
}
func (d DIConvertParams) OutputURL() IDIURL {
	rv := objc.SendIfResponds[objc.ID](d.ID, objc.Sel("outputURL"))
	return DIURLFromID(objc.ID(rv))
}
func (d DIConvertParams) SetOutputURL(value IDIURL) {
	objc.SendIfResponds[struct{}](d.ID, objc.Sel("setOutputURL:"), value)
}
func (d DIConvertParams) Passphrase() bool {
	rv := objc.SendIfResponds[bool](d.ID, objc.Sel("passphrase"))
	return rv
}
func (d DIConvertParams) SetPassphrase(value bool) {
	objc.SendIfResponds[struct{}](d.ID, objc.Sel("setPassphrase:"), value)
}
func (d DIConvertParams) PublicKey() string {
	rv := objc.SendIfResponds[objc.ID](d.ID, objc.Sel("publicKey"))
	return foundation.NSStringFromID(rv).String()
}
func (d DIConvertParams) SetPublicKey(value string) {
	objc.SendIfResponds[struct{}](d.ID, objc.Sel("setPublicKey:"), objc.String(value))
}
func (d DIConvertParams) ShadowURLs() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](d.ID, objc.Sel("shadowURLs"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (d DIConvertParams) ShouldValidateShadows() bool {
	rv := objc.SendIfResponds[bool](d.ID, objc.Sel("shouldValidateShadows"))
	return rv
}
func (d DIConvertParams) SetShouldValidateShadows(value bool) {
	objc.SendIfResponds[struct{}](d.ID, objc.Sel("setShouldValidateShadows:"), value)
}
func (d DIConvertParams) SparseBundleBandSize() uint64 {
	rv := objc.SendIfResponds[uint64](d.ID, objc.Sel("sparseBundleBandSize"))
	return rv
}
func (d DIConvertParams) SetSparseBundleBandSize(value uint64) {
	objc.SendIfResponds[struct{}](d.ID, objc.Sel("setSparseBundleBandSize:"), value)
}
func (d DIConvertParams) TemporaryPassphrase() IDITemporaryPassphrase {
	rv := objc.SendIfResponds[objc.ID](d.ID, objc.Sel("temporaryPassphrase"))
	return DITemporaryPassphraseFromID(objc.ID(rv))
}
func (d DIConvertParams) UseFormatMappingInfo() bool {
	rv := objc.SendIfResponds[bool](d.ID, objc.Sel("useFormatMappingInfo"))
	return rv
}
func (d DIConvertParams) SetUseFormatMappingInfo(value bool) {
	objc.SendIfResponds[struct{}](d.ID, objc.Sel("setUseFormatMappingInfo:"), value)
}

// ConvertWithCompletionBlockSync is a synchronous wrapper around [DIConvertParams.ConvertWithCompletionBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (d DIConvertParams) ConvertWithCompletionBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	d.ConvertWithCompletionBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
