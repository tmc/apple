// Code generated from Apple documentation for Security. DO NOT EDIT.

package security

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
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
		return fmt.Sprintf("Security: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("Security: symbol %s unavailable on this system", e.symbol)
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
		return fmt.Errorf("Security: symbol %s unavailable because the framework could not be loaded", name)
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
			*errDst = fmt.Errorf("Security: register symbol %s: %v", name, r)
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

var _authorizationCopyInfo func(authorization AuthorizationRef, tag AuthorizationString, info *AuthorizationItemSet) int32
var _authorizationCopyInfoErr error

func tryAuthorizationCopyInfo(authorization AuthorizationRef, tag AuthorizationString, info *AuthorizationItemSet) (int32, error) {
	if _authorizationCopyInfo == nil {
		return 0, symbolCallError("AuthorizationCopyInfo", "10.0", _authorizationCopyInfoErr)
	}
	return _authorizationCopyInfo(authorization, tag, info), nil
}

// AuthorizationCopyInfo retrieves supporting data such as the user name and other information gathered during evaluation of authorization.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationCopyInfo(_:_:_:)
func AuthorizationCopyInfo(authorization AuthorizationRef, tag AuthorizationString, info *AuthorizationItemSet) int32 {
	result, callErr := tryAuthorizationCopyInfo(authorization, tag, info)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _authorizationCopyRights func(authorization AuthorizationRef, rights *AuthorizationRights, environment *AuthorizationEnvironment, flags AuthorizationFlags, authorizedRights *AuthorizationRights) int32
var _authorizationCopyRightsErr error

func tryAuthorizationCopyRights(authorization AuthorizationRef, rights *AuthorizationRights, environment *AuthorizationEnvironment, flags AuthorizationFlags, authorizedRights *AuthorizationRights) (int32, error) {
	if _authorizationCopyRights == nil {
		return 0, symbolCallError("AuthorizationCopyRights", "10.0", _authorizationCopyRightsErr)
	}
	return _authorizationCopyRights(authorization, rights, environment, flags, authorizedRights), nil
}

// AuthorizationCopyRights authorizes and preauthorizes rights synchronously.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationCopyRights(_:_:_:_:_:)
func AuthorizationCopyRights(authorization AuthorizationRef, rights *AuthorizationRights, environment *AuthorizationEnvironment, flags AuthorizationFlags, authorizedRights *AuthorizationRights) int32 {
	result, callErr := tryAuthorizationCopyRights(authorization, rights, environment, flags, authorizedRights)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _authorizationCopyRightsAsync func(authorization AuthorizationRef, rights *AuthorizationRights, environment *AuthorizationEnvironment, flags AuthorizationFlags, callbackBlock unsafe.Pointer)
var _authorizationCopyRightsAsyncErr error

func tryAuthorizationCopyRightsAsync(authorization AuthorizationRef, rights *AuthorizationRights, environment *AuthorizationEnvironment, flags AuthorizationFlags, callbackBlock AuthorizationAsyncCallback) error {
	if _authorizationCopyRightsAsync == nil {
		return symbolCallError("AuthorizationCopyRightsAsync", "10.7", _authorizationCopyRightsAsyncErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 int, blockArg1 *AuthorizationItemSet) {
		callbackBlock(blockArg0, blockArg1)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_authorizationCopyRightsAsync(authorization, rights, environment, flags, _block0)
	return nil
}

// AuthorizationCopyRightsAsync authorizes and preauthorizes rights asynchronously.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationCopyRightsAsync(_:_:_:_:_:)
func AuthorizationCopyRightsAsync(authorization AuthorizationRef, rights *AuthorizationRights, environment *AuthorizationEnvironment, flags AuthorizationFlags, callbackBlock AuthorizationAsyncCallback) {
	if callErr := tryAuthorizationCopyRightsAsync(authorization, rights, environment, flags, callbackBlock); callErr != nil {
		panic(callErr)
	}
}

var _authorizationCreate func(rights *AuthorizationRights, environment *AuthorizationEnvironment, flags AuthorizationFlags, authorization *AuthorizationRef) int32
var _authorizationCreateErr error

func tryAuthorizationCreate(rights *AuthorizationRights, environment *AuthorizationEnvironment, flags AuthorizationFlags, authorization *AuthorizationRef) (int32, error) {
	if _authorizationCreate == nil {
		return 0, symbolCallError("AuthorizationCreate", "10.0", _authorizationCreateErr)
	}
	return _authorizationCreate(rights, environment, flags, authorization), nil
}

// AuthorizationCreate creates a new authorization reference and provides an option to authorize or preauthorize rights.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationCreate(_:_:_:_:)
func AuthorizationCreate(rights *AuthorizationRights, environment *AuthorizationEnvironment, flags AuthorizationFlags, authorization *AuthorizationRef) int32 {
	result, callErr := tryAuthorizationCreate(rights, environment, flags, authorization)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _authorizationCreateFromExternalForm func(extForm *AuthorizationExternalForm, authorization *AuthorizationRef) int32
var _authorizationCreateFromExternalFormErr error

func tryAuthorizationCreateFromExternalForm(extForm *AuthorizationExternalForm, authorization *AuthorizationRef) (int32, error) {
	if _authorizationCreateFromExternalForm == nil {
		return 0, symbolCallError("AuthorizationCreateFromExternalForm", "10.0", _authorizationCreateFromExternalFormErr)
	}
	return _authorizationCreateFromExternalForm(extForm, authorization), nil
}

// AuthorizationCreateFromExternalForm internalizes the external representation of an authorization reference.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationCreateFromExternalForm(_:_:)
func AuthorizationCreateFromExternalForm(extForm *AuthorizationExternalForm, authorization *AuthorizationRef) int32 {
	result, callErr := tryAuthorizationCreateFromExternalForm(extForm, authorization)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _authorizationFree func(authorization AuthorizationRef, flags AuthorizationFlags) int32
var _authorizationFreeErr error

func tryAuthorizationFree(authorization AuthorizationRef, flags AuthorizationFlags) (int32, error) {
	if _authorizationFree == nil {
		return 0, symbolCallError("AuthorizationFree", "10.0", _authorizationFreeErr)
	}
	return _authorizationFree(authorization, flags), nil
}

// AuthorizationFree frees the memory associated with an authorization reference.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationFree(_:_:)
func AuthorizationFree(authorization AuthorizationRef, flags AuthorizationFlags) int32 {
	result, callErr := tryAuthorizationFree(authorization, flags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _authorizationFreeItemSet func(set *AuthorizationItemSet) int32
var _authorizationFreeItemSetErr error

func tryAuthorizationFreeItemSet(set *AuthorizationItemSet) (int32, error) {
	if _authorizationFreeItemSet == nil {
		return 0, symbolCallError("AuthorizationFreeItemSet", "10.0", _authorizationFreeItemSetErr)
	}
	return _authorizationFreeItemSet(set), nil
}

// AuthorizationFreeItemSet frees the memory associated with a set of authorization items.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationFreeItemSet(_:)
func AuthorizationFreeItemSet(set *AuthorizationItemSet) int32 {
	result, callErr := tryAuthorizationFreeItemSet(set)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _authorizationMakeExternalForm func(authorization AuthorizationRef, extForm *AuthorizationExternalForm) int32
var _authorizationMakeExternalFormErr error

func tryAuthorizationMakeExternalForm(authorization AuthorizationRef, extForm *AuthorizationExternalForm) (int32, error) {
	if _authorizationMakeExternalForm == nil {
		return 0, symbolCallError("AuthorizationMakeExternalForm", "10.0", _authorizationMakeExternalFormErr)
	}
	return _authorizationMakeExternalForm(authorization, extForm), nil
}

// AuthorizationMakeExternalForm creates an external representation of an authorization reference.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationMakeExternalForm(_:_:)
func AuthorizationMakeExternalForm(authorization AuthorizationRef, extForm *AuthorizationExternalForm) int32 {
	result, callErr := tryAuthorizationMakeExternalForm(authorization, extForm)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _authorizationRightGet func(rightName string, rightDefinition *corefoundation.CFDictionaryRef) int32
var _authorizationRightGetErr error

func tryAuthorizationRightGet(rightName string, rightDefinition *corefoundation.CFDictionaryRef) (int32, error) {
	if _authorizationRightGet == nil {
		return 0, symbolCallError("AuthorizationRightGet", "10.0", _authorizationRightGetErr)
	}
	return _authorizationRightGet(rightName, rightDefinition), nil
}

// AuthorizationRightGet retrieves a right definition as a dictionary.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationRightGet(_:_:)
func AuthorizationRightGet(rightName string, rightDefinition *corefoundation.CFDictionaryRef) int32 {
	result, callErr := tryAuthorizationRightGet(rightName, rightDefinition)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _authorizationRightRemove func(authRef AuthorizationRef, rightName string) int32
var _authorizationRightRemoveErr error

func tryAuthorizationRightRemove(authRef AuthorizationRef, rightName string) (int32, error) {
	if _authorizationRightRemove == nil {
		return 0, symbolCallError("AuthorizationRightRemove", "10.0", _authorizationRightRemoveErr)
	}
	return _authorizationRightRemove(authRef, rightName), nil
}

// AuthorizationRightRemove removes a right from the policy database.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationRightRemove(_:_:)
func AuthorizationRightRemove(authRef AuthorizationRef, rightName string) int32 {
	result, callErr := tryAuthorizationRightRemove(authRef, rightName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _authorizationRightSet func(authRef AuthorizationRef, rightName string, rightDefinition corefoundation.CFTypeRef, descriptionKey corefoundation.CFStringRef, bundle corefoundation.CFBundle, localeTableName corefoundation.CFStringRef) int32
var _authorizationRightSetErr error

func tryAuthorizationRightSet(authRef AuthorizationRef, rightName string, rightDefinition corefoundation.CFTypeRef, descriptionKey corefoundation.CFStringRef, bundle corefoundation.CFBundle, localeTableName corefoundation.CFStringRef) (int32, error) {
	if _authorizationRightSet == nil {
		return 0, symbolCallError("AuthorizationRightSet", "10.0", _authorizationRightSetErr)
	}
	return _authorizationRightSet(authRef, rightName, rightDefinition, descriptionKey, bundle, localeTableName), nil
}

// AuthorizationRightSet creates or updates a right entry in the policy database.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationRightSet(_:_:_:_:_:_:)
func AuthorizationRightSet(authRef AuthorizationRef, rightName string, rightDefinition corefoundation.CFTypeRef, descriptionKey corefoundation.CFStringRef, bundle corefoundation.CFBundle, localeTableName corefoundation.CFStringRef) int32 {
	result, callErr := tryAuthorizationRightSet(authRef, rightName, rightDefinition, descriptionKey, bundle, localeTableName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSDecoderCopyAllCerts func(cmsDecoder CMSDecoderRef, certsOut *corefoundation.CFArrayRef) int32
var _cMSDecoderCopyAllCertsErr error

func tryCMSDecoderCopyAllCerts(cmsDecoder CMSDecoderRef, certsOut *corefoundation.CFArrayRef) (int32, error) {
	if _cMSDecoderCopyAllCerts == nil {
		return 0, symbolCallError("CMSDecoderCopyAllCerts", "10.5", _cMSDecoderCopyAllCertsErr)
	}
	return _cMSDecoderCopyAllCerts(cmsDecoder, certsOut), nil
}

// CMSDecoderCopyAllCerts obtains an array of all of the certificates in a message.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderCopyAllCerts(_:_:)
func CMSDecoderCopyAllCerts(cmsDecoder CMSDecoderRef, certsOut *corefoundation.CFArrayRef) int32 {
	result, callErr := tryCMSDecoderCopyAllCerts(cmsDecoder, certsOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSDecoderCopyContent func(cmsDecoder CMSDecoderRef, contentOut *corefoundation.CFDataRef) int32
var _cMSDecoderCopyContentErr error

func tryCMSDecoderCopyContent(cmsDecoder CMSDecoderRef, contentOut *corefoundation.CFDataRef) (int32, error) {
	if _cMSDecoderCopyContent == nil {
		return 0, symbolCallError("CMSDecoderCopyContent", "10.5", _cMSDecoderCopyContentErr)
	}
	return _cMSDecoderCopyContent(cmsDecoder, contentOut), nil
}

// CMSDecoderCopyContent obtains the message content, if any.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderCopyContent(_:_:)
func CMSDecoderCopyContent(cmsDecoder CMSDecoderRef, contentOut *corefoundation.CFDataRef) int32 {
	result, callErr := tryCMSDecoderCopyContent(cmsDecoder, contentOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSDecoderCopyDetachedContent func(cmsDecoder CMSDecoderRef, detachedContentOut *corefoundation.CFDataRef) int32
var _cMSDecoderCopyDetachedContentErr error

func tryCMSDecoderCopyDetachedContent(cmsDecoder CMSDecoderRef, detachedContentOut *corefoundation.CFDataRef) (int32, error) {
	if _cMSDecoderCopyDetachedContent == nil {
		return 0, symbolCallError("CMSDecoderCopyDetachedContent", "10.5", _cMSDecoderCopyDetachedContentErr)
	}
	return _cMSDecoderCopyDetachedContent(cmsDecoder, detachedContentOut), nil
}

// CMSDecoderCopyDetachedContent obtains the detached content specified with the [CMSDecoderSetDetachedContent] function.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderCopyDetachedContent(_:_:)
func CMSDecoderCopyDetachedContent(cmsDecoder CMSDecoderRef, detachedContentOut *corefoundation.CFDataRef) int32 {
	result, callErr := tryCMSDecoderCopyDetachedContent(cmsDecoder, detachedContentOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSDecoderCopyEncapsulatedContentType func(cmsDecoder CMSDecoderRef, eContentTypeOut *corefoundation.CFDataRef) int32
var _cMSDecoderCopyEncapsulatedContentTypeErr error

func tryCMSDecoderCopyEncapsulatedContentType(cmsDecoder CMSDecoderRef, eContentTypeOut *corefoundation.CFDataRef) (int32, error) {
	if _cMSDecoderCopyEncapsulatedContentType == nil {
		return 0, symbolCallError("CMSDecoderCopyEncapsulatedContentType", "10.5", _cMSDecoderCopyEncapsulatedContentTypeErr)
	}
	return _cMSDecoderCopyEncapsulatedContentType(cmsDecoder, eContentTypeOut), nil
}

// CMSDecoderCopyEncapsulatedContentType obtains the object identifier for the encapsulated data of a signed message.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderCopyEncapsulatedContentType(_:_:)
func CMSDecoderCopyEncapsulatedContentType(cmsDecoder CMSDecoderRef, eContentTypeOut *corefoundation.CFDataRef) int32 {
	result, callErr := tryCMSDecoderCopyEncapsulatedContentType(cmsDecoder, eContentTypeOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSDecoderCopySignerCert func(cmsDecoder CMSDecoderRef, signerIndex uintptr, signerCertOut *SecCertificateRef) int32
var _cMSDecoderCopySignerCertErr error

func tryCMSDecoderCopySignerCert(cmsDecoder CMSDecoderRef, signerIndex uintptr, signerCertOut *SecCertificateRef) (int32, error) {
	if _cMSDecoderCopySignerCert == nil {
		return 0, symbolCallError("CMSDecoderCopySignerCert", "10.5", _cMSDecoderCopySignerCertErr)
	}
	return _cMSDecoderCopySignerCert(cmsDecoder, signerIndex, signerCertOut), nil
}

// CMSDecoderCopySignerCert obtains the certificate of the specified signer of a CMS message.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderCopySignerCert(_:_:_:)
func CMSDecoderCopySignerCert(cmsDecoder CMSDecoderRef, signerIndex uintptr, signerCertOut *SecCertificateRef) int32 {
	result, callErr := tryCMSDecoderCopySignerCert(cmsDecoder, signerIndex, signerCertOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSDecoderCopySignerEmailAddress func(cmsDecoder CMSDecoderRef, signerIndex uintptr, signerEmailAddressOut *corefoundation.CFStringRef) int32
var _cMSDecoderCopySignerEmailAddressErr error

func tryCMSDecoderCopySignerEmailAddress(cmsDecoder CMSDecoderRef, signerIndex uintptr, signerEmailAddressOut *corefoundation.CFStringRef) (int32, error) {
	if _cMSDecoderCopySignerEmailAddress == nil {
		return 0, symbolCallError("CMSDecoderCopySignerEmailAddress", "10.5", _cMSDecoderCopySignerEmailAddressErr)
	}
	return _cMSDecoderCopySignerEmailAddress(cmsDecoder, signerIndex, signerEmailAddressOut), nil
}

// CMSDecoderCopySignerEmailAddress obtains the email address of the specified signer of a CMS message.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderCopySignerEmailAddress(_:_:_:)
func CMSDecoderCopySignerEmailAddress(cmsDecoder CMSDecoderRef, signerIndex uintptr, signerEmailAddressOut *corefoundation.CFStringRef) int32 {
	result, callErr := tryCMSDecoderCopySignerEmailAddress(cmsDecoder, signerIndex, signerEmailAddressOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSDecoderCopySignerSigningTime func(cmsDecoder CMSDecoderRef, signerIndex uintptr, signingTime *corefoundation.CFAbsoluteTime) int32
var _cMSDecoderCopySignerSigningTimeErr error

func tryCMSDecoderCopySignerSigningTime(cmsDecoder CMSDecoderRef, signerIndex uintptr, signingTime *corefoundation.CFAbsoluteTime) (int32, error) {
	if _cMSDecoderCopySignerSigningTime == nil {
		return 0, symbolCallError("CMSDecoderCopySignerSigningTime", "10.8", _cMSDecoderCopySignerSigningTimeErr)
	}
	return _cMSDecoderCopySignerSigningTime(cmsDecoder, signerIndex, signingTime), nil
}

// CMSDecoderCopySignerSigningTime obtains the signing time of a CMS message, if present.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderCopySignerSigningTime(_:_:_:)
func CMSDecoderCopySignerSigningTime(cmsDecoder CMSDecoderRef, signerIndex uintptr, signingTime *corefoundation.CFAbsoluteTime) int32 {
	result, callErr := tryCMSDecoderCopySignerSigningTime(cmsDecoder, signerIndex, signingTime)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSDecoderCopySignerStatus func(cmsDecoder CMSDecoderRef, signerIndex uintptr, policyOrArray corefoundation.CFTypeRef, evaluateSecTrust bool, signerStatusOut *CMSSignerStatus, secTrustOut *SecTrustRef, certVerifyResultCodeOut *int32) int32
var _cMSDecoderCopySignerStatusErr error

func tryCMSDecoderCopySignerStatus(cmsDecoder CMSDecoderRef, signerIndex uintptr, policyOrArray corefoundation.CFTypeRef, evaluateSecTrust bool, signerStatusOut *CMSSignerStatus, secTrustOut *SecTrustRef, certVerifyResultCodeOut *int32) (int32, error) {
	if _cMSDecoderCopySignerStatus == nil {
		return 0, symbolCallError("CMSDecoderCopySignerStatus", "10.5", _cMSDecoderCopySignerStatusErr)
	}
	return _cMSDecoderCopySignerStatus(cmsDecoder, signerIndex, policyOrArray, evaluateSecTrust, signerStatusOut, secTrustOut, certVerifyResultCodeOut), nil
}

// CMSDecoderCopySignerStatus obtains the status of a CMS message’s signature.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderCopySignerStatus(_:_:_:_:_:_:_:)
func CMSDecoderCopySignerStatus(cmsDecoder CMSDecoderRef, signerIndex uintptr, policyOrArray corefoundation.CFTypeRef, evaluateSecTrust bool, signerStatusOut *CMSSignerStatus, secTrustOut *SecTrustRef, certVerifyResultCodeOut *int32) int32 {
	result, callErr := tryCMSDecoderCopySignerStatus(cmsDecoder, signerIndex, policyOrArray, evaluateSecTrust, signerStatusOut, secTrustOut, certVerifyResultCodeOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSDecoderCopySignerTimestamp func(cmsDecoder CMSDecoderRef, signerIndex uintptr, timestamp *corefoundation.CFAbsoluteTime) int32
var _cMSDecoderCopySignerTimestampErr error

func tryCMSDecoderCopySignerTimestamp(cmsDecoder CMSDecoderRef, signerIndex uintptr, timestamp *corefoundation.CFAbsoluteTime) (int32, error) {
	if _cMSDecoderCopySignerTimestamp == nil {
		return 0, symbolCallError("CMSDecoderCopySignerTimestamp", "10.8", _cMSDecoderCopySignerTimestampErr)
	}
	return _cMSDecoderCopySignerTimestamp(cmsDecoder, signerIndex, timestamp), nil
}

// CMSDecoderCopySignerTimestamp returns the timestamp of a signer of a CMS message, if present.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderCopySignerTimestamp(_:_:_:)
func CMSDecoderCopySignerTimestamp(cmsDecoder CMSDecoderRef, signerIndex uintptr, timestamp *corefoundation.CFAbsoluteTime) int32 {
	result, callErr := tryCMSDecoderCopySignerTimestamp(cmsDecoder, signerIndex, timestamp)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSDecoderCopySignerTimestampCertificates func(cmsDecoder CMSDecoderRef, signerIndex uintptr, certificateRefs *corefoundation.CFArrayRef) int32
var _cMSDecoderCopySignerTimestampCertificatesErr error

func tryCMSDecoderCopySignerTimestampCertificates(cmsDecoder CMSDecoderRef, signerIndex uintptr, certificateRefs *corefoundation.CFArrayRef) (int32, error) {
	if _cMSDecoderCopySignerTimestampCertificates == nil {
		return 0, symbolCallError("CMSDecoderCopySignerTimestampCertificates", "10.8", _cMSDecoderCopySignerTimestampCertificatesErr)
	}
	return _cMSDecoderCopySignerTimestampCertificates(cmsDecoder, signerIndex, certificateRefs), nil
}

// CMSDecoderCopySignerTimestampCertificates returns an array containing the certificates from a timestamp response.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderCopySignerTimestampCertificates(_:_:_:)
func CMSDecoderCopySignerTimestampCertificates(cmsDecoder CMSDecoderRef, signerIndex uintptr, certificateRefs *corefoundation.CFArrayRef) int32 {
	result, callErr := tryCMSDecoderCopySignerTimestampCertificates(cmsDecoder, signerIndex, certificateRefs)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSDecoderCopySignerTimestampWithPolicy func(cmsDecoder CMSDecoderRef, timeStampPolicy corefoundation.CFTypeRef, signerIndex uintptr, timestamp *corefoundation.CFAbsoluteTime) int32
var _cMSDecoderCopySignerTimestampWithPolicyErr error

func tryCMSDecoderCopySignerTimestampWithPolicy(cmsDecoder CMSDecoderRef, timeStampPolicy corefoundation.CFTypeRef, signerIndex uintptr, timestamp *corefoundation.CFAbsoluteTime) (int32, error) {
	if _cMSDecoderCopySignerTimestampWithPolicy == nil {
		return 0, symbolCallError("CMSDecoderCopySignerTimestampWithPolicy", "10.10", _cMSDecoderCopySignerTimestampWithPolicyErr)
	}
	return _cMSDecoderCopySignerTimestampWithPolicy(cmsDecoder, timeStampPolicy, signerIndex, timestamp), nil
}

// CMSDecoderCopySignerTimestampWithPolicy returns the timestamp of a signer of a CMS message using a given policy, if present.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderCopySignerTimestampWithPolicy(_:_:_:_:)
func CMSDecoderCopySignerTimestampWithPolicy(cmsDecoder CMSDecoderRef, timeStampPolicy corefoundation.CFTypeRef, signerIndex uintptr, timestamp *corefoundation.CFAbsoluteTime) int32 {
	result, callErr := tryCMSDecoderCopySignerTimestampWithPolicy(cmsDecoder, timeStampPolicy, signerIndex, timestamp)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSDecoderCreate func(cmsDecoderOut *CMSDecoderRef) int32
var _cMSDecoderCreateErr error

func tryCMSDecoderCreate(cmsDecoderOut *CMSDecoderRef) (int32, error) {
	if _cMSDecoderCreate == nil {
		return 0, symbolCallError("CMSDecoderCreate", "10.5", _cMSDecoderCreateErr)
	}
	return _cMSDecoderCreate(cmsDecoderOut), nil
}

// CMSDecoderCreate creates a CMSDecoder reference.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderCreate(_:)
func CMSDecoderCreate(cmsDecoderOut *CMSDecoderRef) int32 {
	result, callErr := tryCMSDecoderCreate(cmsDecoderOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSDecoderFinalizeMessage func(cmsDecoder CMSDecoderRef) int32
var _cMSDecoderFinalizeMessageErr error

func tryCMSDecoderFinalizeMessage(cmsDecoder CMSDecoderRef) (int32, error) {
	if _cMSDecoderFinalizeMessage == nil {
		return 0, symbolCallError("CMSDecoderFinalizeMessage", "10.5", _cMSDecoderFinalizeMessageErr)
	}
	return _cMSDecoderFinalizeMessage(cmsDecoder), nil
}

// CMSDecoderFinalizeMessage indicates that there is no more data to decode.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderFinalizeMessage(_:)
func CMSDecoderFinalizeMessage(cmsDecoder CMSDecoderRef) int32 {
	result, callErr := tryCMSDecoderFinalizeMessage(cmsDecoder)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSDecoderGetNumSigners func(cmsDecoder CMSDecoderRef, numSignersOut *uintptr) int32
var _cMSDecoderGetNumSignersErr error

func tryCMSDecoderGetNumSigners(cmsDecoder CMSDecoderRef, numSignersOut *uintptr) (int32, error) {
	if _cMSDecoderGetNumSigners == nil {
		return 0, symbolCallError("CMSDecoderGetNumSigners", "10.5", _cMSDecoderGetNumSignersErr)
	}
	return _cMSDecoderGetNumSigners(cmsDecoder, numSignersOut), nil
}

// CMSDecoderGetNumSigners obtains the number of signers of a message.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderGetNumSigners(_:_:)
func CMSDecoderGetNumSigners(cmsDecoder CMSDecoderRef, numSignersOut *uintptr) int32 {
	result, callErr := tryCMSDecoderGetNumSigners(cmsDecoder, numSignersOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSDecoderGetTypeID func() uint
var _cMSDecoderGetTypeIDErr error

func tryCMSDecoderGetTypeID() (uint, error) {
	if _cMSDecoderGetTypeID == nil {
		return 0, symbolCallError("CMSDecoderGetTypeID", "10.0", _cMSDecoderGetTypeIDErr)
	}
	return _cMSDecoderGetTypeID(), nil
}

// CMSDecoderGetTypeID returns the type identifier for the CMSDecoder opaque type.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderGetTypeID()
func CMSDecoderGetTypeID() uint {
	result, callErr := tryCMSDecoderGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSDecoderIsContentEncrypted func(cmsDecoder CMSDecoderRef, isEncryptedOut *bool) int32
var _cMSDecoderIsContentEncryptedErr error

func tryCMSDecoderIsContentEncrypted(cmsDecoder CMSDecoderRef, isEncryptedOut *bool) (int32, error) {
	if _cMSDecoderIsContentEncrypted == nil {
		return 0, symbolCallError("CMSDecoderIsContentEncrypted", "10.5", _cMSDecoderIsContentEncryptedErr)
	}
	return _cMSDecoderIsContentEncrypted(cmsDecoder, isEncryptedOut), nil
}

// CMSDecoderIsContentEncrypted determines whether a CMS message was encrypted.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderIsContentEncrypted(_:_:)
func CMSDecoderIsContentEncrypted(cmsDecoder CMSDecoderRef, isEncryptedOut *bool) int32 {
	result, callErr := tryCMSDecoderIsContentEncrypted(cmsDecoder, isEncryptedOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSDecoderSetDetachedContent func(cmsDecoder CMSDecoderRef, detachedContent corefoundation.CFDataRef) int32
var _cMSDecoderSetDetachedContentErr error

func tryCMSDecoderSetDetachedContent(cmsDecoder CMSDecoderRef, detachedContent corefoundation.CFDataRef) (int32, error) {
	if _cMSDecoderSetDetachedContent == nil {
		return 0, symbolCallError("CMSDecoderSetDetachedContent", "10.5", _cMSDecoderSetDetachedContentErr)
	}
	return _cMSDecoderSetDetachedContent(cmsDecoder, detachedContent), nil
}

// CMSDecoderSetDetachedContent specifies the message’s detached content, if any.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderSetDetachedContent(_:_:)
func CMSDecoderSetDetachedContent(cmsDecoder CMSDecoderRef, detachedContent corefoundation.CFDataRef) int32 {
	result, callErr := tryCMSDecoderSetDetachedContent(cmsDecoder, detachedContent)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSDecoderSetSearchKeychain func(cmsDecoder CMSDecoderRef, keychainOrArray corefoundation.CFTypeRef) int32
var _cMSDecoderSetSearchKeychainErr error

func tryCMSDecoderSetSearchKeychain(cmsDecoder CMSDecoderRef, keychainOrArray corefoundation.CFTypeRef) (int32, error) {
	if _cMSDecoderSetSearchKeychain == nil {
		return 0, symbolCallError("CMSDecoderSetSearchKeychain", "10.5", _cMSDecoderSetSearchKeychainErr)
	}
	return _cMSDecoderSetSearchKeychain(cmsDecoder, keychainOrArray), nil
}

// CMSDecoderSetSearchKeychain specifies the keychains to search for intermediate certificates to be used in verifying a signed message’s signer certificates.
//
// Deprecated: Deprecated since macOS 10.13.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderSetSearchKeychain(_:_:)
func CMSDecoderSetSearchKeychain(cmsDecoder CMSDecoderRef, keychainOrArray corefoundation.CFTypeRef) int32 {
	result, callErr := tryCMSDecoderSetSearchKeychain(cmsDecoder, keychainOrArray)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSDecoderUpdateMessage func(cmsDecoder CMSDecoderRef, msgBytes unsafe.Pointer, msgBytesLen uintptr) int32
var _cMSDecoderUpdateMessageErr error

func tryCMSDecoderUpdateMessage(cmsDecoder CMSDecoderRef, msgBytes unsafe.Pointer, msgBytesLen uintptr) (int32, error) {
	if _cMSDecoderUpdateMessage == nil {
		return 0, symbolCallError("CMSDecoderUpdateMessage", "10.5", _cMSDecoderUpdateMessageErr)
	}
	return _cMSDecoderUpdateMessage(cmsDecoder, msgBytes, msgBytesLen), nil
}

// CMSDecoderUpdateMessage feeds raw bytes of the message to be decoded into the decoder.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderUpdateMessage(_:_:_:)
func CMSDecoderUpdateMessage(cmsDecoder CMSDecoderRef, msgBytes unsafe.Pointer, msgBytesLen uintptr) int32 {
	result, callErr := tryCMSDecoderUpdateMessage(cmsDecoder, msgBytes, msgBytesLen)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncode func(signers corefoundation.CFTypeRef, recipients corefoundation.CFTypeRef, eContentType unsafe.Pointer, detachedContent bool, signedAttributes CMSSignedAttributes, content unsafe.Pointer, contentLen uintptr, encodedContentOut *corefoundation.CFDataRef) int32
var _cMSEncodeErr error

func tryCMSEncode(signers corefoundation.CFTypeRef, recipients corefoundation.CFTypeRef, eContentType unsafe.Pointer, detachedContent bool, signedAttributes CMSSignedAttributes, content unsafe.Pointer, contentLen uintptr, encodedContentOut *corefoundation.CFDataRef) (int32, error) {
	if _cMSEncode == nil {
		return 0, symbolCallError("CMSEncode", "10.5", _cMSEncodeErr)
	}
	return _cMSEncode(signers, recipients, eContentType, detachedContent, signedAttributes, content, contentLen, encodedContentOut), nil
}

// CMSEncode encodes a message and obtains the result in one high-level function call.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CMSEncode
func CMSEncode(signers corefoundation.CFTypeRef, recipients corefoundation.CFTypeRef, eContentType unsafe.Pointer, detachedContent bool, signedAttributes CMSSignedAttributes, content unsafe.Pointer, contentLen uintptr, encodedContentOut *corefoundation.CFDataRef) int32 {
	result, callErr := tryCMSEncode(signers, recipients, eContentType, detachedContent, signedAttributes, content, contentLen, encodedContentOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncodeContent func(signers corefoundation.CFTypeRef, recipients corefoundation.CFTypeRef, eContentTypeOID corefoundation.CFTypeRef, detachedContent bool, signedAttributes CMSSignedAttributes, content unsafe.Pointer, contentLen uintptr, encodedContentOut *corefoundation.CFDataRef) int32
var _cMSEncodeContentErr error

func tryCMSEncodeContent(signers corefoundation.CFTypeRef, recipients corefoundation.CFTypeRef, eContentTypeOID corefoundation.CFTypeRef, detachedContent bool, signedAttributes CMSSignedAttributes, content unsafe.Pointer, contentLen uintptr, encodedContentOut *corefoundation.CFDataRef) (int32, error) {
	if _cMSEncodeContent == nil {
		return 0, symbolCallError("CMSEncodeContent", "10.7", _cMSEncodeContentErr)
	}
	return _cMSEncodeContent(signers, recipients, eContentTypeOID, detachedContent, signedAttributes, content, contentLen, encodedContentOut), nil
}

// CMSEncodeContent encodes a message and obtains the result in one high-level function call.
//
// See: https://developer.apple.com/documentation/Security/CMSEncodeContent(_:_:_:_:_:_:_:_:)
func CMSEncodeContent(signers corefoundation.CFTypeRef, recipients corefoundation.CFTypeRef, eContentTypeOID corefoundation.CFTypeRef, detachedContent bool, signedAttributes CMSSignedAttributes, content unsafe.Pointer, contentLen uintptr, encodedContentOut *corefoundation.CFDataRef) int32 {
	result, callErr := tryCMSEncodeContent(signers, recipients, eContentTypeOID, detachedContent, signedAttributes, content, contentLen, encodedContentOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderAddRecipients func(cmsEncoder CMSEncoderRef, recipientOrArray corefoundation.CFTypeRef) int32
var _cMSEncoderAddRecipientsErr error

func tryCMSEncoderAddRecipients(cmsEncoder CMSEncoderRef, recipientOrArray corefoundation.CFTypeRef) (int32, error) {
	if _cMSEncoderAddRecipients == nil {
		return 0, symbolCallError("CMSEncoderAddRecipients", "10.5", _cMSEncoderAddRecipientsErr)
	}
	return _cMSEncoderAddRecipients(cmsEncoder, recipientOrArray), nil
}

// CMSEncoderAddRecipients specifies a message is to be encrypted and specifies the recipients of the message.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderAddRecipients(_:_:)
func CMSEncoderAddRecipients(cmsEncoder CMSEncoderRef, recipientOrArray corefoundation.CFTypeRef) int32 {
	result, callErr := tryCMSEncoderAddRecipients(cmsEncoder, recipientOrArray)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderAddSignedAttributes func(cmsEncoder CMSEncoderRef, signedAttributes CMSSignedAttributes) int32
var _cMSEncoderAddSignedAttributesErr error

func tryCMSEncoderAddSignedAttributes(cmsEncoder CMSEncoderRef, signedAttributes CMSSignedAttributes) (int32, error) {
	if _cMSEncoderAddSignedAttributes == nil {
		return 0, symbolCallError("CMSEncoderAddSignedAttributes", "10.5", _cMSEncoderAddSignedAttributesErr)
	}
	return _cMSEncoderAddSignedAttributes(cmsEncoder, signedAttributes), nil
}

// CMSEncoderAddSignedAttributes specifies attributes for a signed message.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderAddSignedAttributes(_:_:)
func CMSEncoderAddSignedAttributes(cmsEncoder CMSEncoderRef, signedAttributes CMSSignedAttributes) int32 {
	result, callErr := tryCMSEncoderAddSignedAttributes(cmsEncoder, signedAttributes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderAddSigners func(cmsEncoder CMSEncoderRef, signerOrArray corefoundation.CFTypeRef) int32
var _cMSEncoderAddSignersErr error

func tryCMSEncoderAddSigners(cmsEncoder CMSEncoderRef, signerOrArray corefoundation.CFTypeRef) (int32, error) {
	if _cMSEncoderAddSigners == nil {
		return 0, symbolCallError("CMSEncoderAddSigners", "10.5", _cMSEncoderAddSignersErr)
	}
	return _cMSEncoderAddSigners(cmsEncoder, signerOrArray), nil
}

// CMSEncoderAddSigners specifies signers of the message.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderAddSigners(_:_:)
func CMSEncoderAddSigners(cmsEncoder CMSEncoderRef, signerOrArray corefoundation.CFTypeRef) int32 {
	result, callErr := tryCMSEncoderAddSigners(cmsEncoder, signerOrArray)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderAddSupportingCerts func(cmsEncoder CMSEncoderRef, certOrArray corefoundation.CFTypeRef) int32
var _cMSEncoderAddSupportingCertsErr error

func tryCMSEncoderAddSupportingCerts(cmsEncoder CMSEncoderRef, certOrArray corefoundation.CFTypeRef) (int32, error) {
	if _cMSEncoderAddSupportingCerts == nil {
		return 0, symbolCallError("CMSEncoderAddSupportingCerts", "10.5", _cMSEncoderAddSupportingCertsErr)
	}
	return _cMSEncoderAddSupportingCerts(cmsEncoder, certOrArray), nil
}

// CMSEncoderAddSupportingCerts adds certificates to a message.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderAddSupportingCerts(_:_:)
func CMSEncoderAddSupportingCerts(cmsEncoder CMSEncoderRef, certOrArray corefoundation.CFTypeRef) int32 {
	result, callErr := tryCMSEncoderAddSupportingCerts(cmsEncoder, certOrArray)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderCopyEncapsulatedContentType func(cmsEncoder CMSEncoderRef, eContentTypeOut *corefoundation.CFDataRef) int32
var _cMSEncoderCopyEncapsulatedContentTypeErr error

func tryCMSEncoderCopyEncapsulatedContentType(cmsEncoder CMSEncoderRef, eContentTypeOut *corefoundation.CFDataRef) (int32, error) {
	if _cMSEncoderCopyEncapsulatedContentType == nil {
		return 0, symbolCallError("CMSEncoderCopyEncapsulatedContentType", "10.5", _cMSEncoderCopyEncapsulatedContentTypeErr)
	}
	return _cMSEncoderCopyEncapsulatedContentType(cmsEncoder, eContentTypeOut), nil
}

// CMSEncoderCopyEncapsulatedContentType obtains the object identifier for the encapsulated data of a signed message.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderCopyEncapsulatedContentType(_:_:)
func CMSEncoderCopyEncapsulatedContentType(cmsEncoder CMSEncoderRef, eContentTypeOut *corefoundation.CFDataRef) int32 {
	result, callErr := tryCMSEncoderCopyEncapsulatedContentType(cmsEncoder, eContentTypeOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderCopyEncodedContent func(cmsEncoder CMSEncoderRef, encodedContentOut *corefoundation.CFDataRef) int32
var _cMSEncoderCopyEncodedContentErr error

func tryCMSEncoderCopyEncodedContent(cmsEncoder CMSEncoderRef, encodedContentOut *corefoundation.CFDataRef) (int32, error) {
	if _cMSEncoderCopyEncodedContent == nil {
		return 0, symbolCallError("CMSEncoderCopyEncodedContent", "10.5", _cMSEncoderCopyEncodedContentErr)
	}
	return _cMSEncoderCopyEncodedContent(cmsEncoder, encodedContentOut), nil
}

// CMSEncoderCopyEncodedContent finishes encoding the message and obtains the encoded result.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderCopyEncodedContent(_:_:)
func CMSEncoderCopyEncodedContent(cmsEncoder CMSEncoderRef, encodedContentOut *corefoundation.CFDataRef) int32 {
	result, callErr := tryCMSEncoderCopyEncodedContent(cmsEncoder, encodedContentOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderCopyRecipients func(cmsEncoder CMSEncoderRef, recipientsOut *corefoundation.CFArrayRef) int32
var _cMSEncoderCopyRecipientsErr error

func tryCMSEncoderCopyRecipients(cmsEncoder CMSEncoderRef, recipientsOut *corefoundation.CFArrayRef) (int32, error) {
	if _cMSEncoderCopyRecipients == nil {
		return 0, symbolCallError("CMSEncoderCopyRecipients", "10.5", _cMSEncoderCopyRecipientsErr)
	}
	return _cMSEncoderCopyRecipients(cmsEncoder, recipientsOut), nil
}

// CMSEncoderCopyRecipients obtains the array of recipients specified with the [CMSEncoderAddRecipients] function.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderCopyRecipients(_:_:)
func CMSEncoderCopyRecipients(cmsEncoder CMSEncoderRef, recipientsOut *corefoundation.CFArrayRef) int32 {
	result, callErr := tryCMSEncoderCopyRecipients(cmsEncoder, recipientsOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderCopySignerTimestamp func(cmsEncoder CMSEncoderRef, signerIndex uintptr, timestamp *corefoundation.CFAbsoluteTime) int32
var _cMSEncoderCopySignerTimestampErr error

func tryCMSEncoderCopySignerTimestamp(cmsEncoder CMSEncoderRef, signerIndex uintptr, timestamp *corefoundation.CFAbsoluteTime) (int32, error) {
	if _cMSEncoderCopySignerTimestamp == nil {
		return 0, symbolCallError("CMSEncoderCopySignerTimestamp", "10.8", _cMSEncoderCopySignerTimestampErr)
	}
	return _cMSEncoderCopySignerTimestamp(cmsEncoder, signerIndex, timestamp), nil
}

// CMSEncoderCopySignerTimestamp returns the timestamp of a signer of a CMS message, if present.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderCopySignerTimestamp(_:_:_:)
func CMSEncoderCopySignerTimestamp(cmsEncoder CMSEncoderRef, signerIndex uintptr, timestamp *corefoundation.CFAbsoluteTime) int32 {
	result, callErr := tryCMSEncoderCopySignerTimestamp(cmsEncoder, signerIndex, timestamp)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderCopySignerTimestampWithPolicy func(cmsEncoder CMSEncoderRef, timeStampPolicy corefoundation.CFTypeRef, signerIndex uintptr, timestamp *corefoundation.CFAbsoluteTime) int32
var _cMSEncoderCopySignerTimestampWithPolicyErr error

func tryCMSEncoderCopySignerTimestampWithPolicy(cmsEncoder CMSEncoderRef, timeStampPolicy corefoundation.CFTypeRef, signerIndex uintptr, timestamp *corefoundation.CFAbsoluteTime) (int32, error) {
	if _cMSEncoderCopySignerTimestampWithPolicy == nil {
		return 0, symbolCallError("CMSEncoderCopySignerTimestampWithPolicy", "10.10", _cMSEncoderCopySignerTimestampWithPolicyErr)
	}
	return _cMSEncoderCopySignerTimestampWithPolicy(cmsEncoder, timeStampPolicy, signerIndex, timestamp), nil
}

// CMSEncoderCopySignerTimestampWithPolicy returns the timestamp of a signer of a CMS message using a particular policy, if present.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderCopySignerTimestampWithPolicy(_:_:_:_:)
func CMSEncoderCopySignerTimestampWithPolicy(cmsEncoder CMSEncoderRef, timeStampPolicy corefoundation.CFTypeRef, signerIndex uintptr, timestamp *corefoundation.CFAbsoluteTime) int32 {
	result, callErr := tryCMSEncoderCopySignerTimestampWithPolicy(cmsEncoder, timeStampPolicy, signerIndex, timestamp)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderCopySigners func(cmsEncoder CMSEncoderRef, signersOut *corefoundation.CFArrayRef) int32
var _cMSEncoderCopySignersErr error

func tryCMSEncoderCopySigners(cmsEncoder CMSEncoderRef, signersOut *corefoundation.CFArrayRef) (int32, error) {
	if _cMSEncoderCopySigners == nil {
		return 0, symbolCallError("CMSEncoderCopySigners", "10.5", _cMSEncoderCopySignersErr)
	}
	return _cMSEncoderCopySigners(cmsEncoder, signersOut), nil
}

// CMSEncoderCopySigners obtains the array of signers specified with the [CMSEncoderAddSigners] function.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderCopySigners(_:_:)
func CMSEncoderCopySigners(cmsEncoder CMSEncoderRef, signersOut *corefoundation.CFArrayRef) int32 {
	result, callErr := tryCMSEncoderCopySigners(cmsEncoder, signersOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderCopySupportingCerts func(cmsEncoder CMSEncoderRef, certsOut *corefoundation.CFArrayRef) int32
var _cMSEncoderCopySupportingCertsErr error

func tryCMSEncoderCopySupportingCerts(cmsEncoder CMSEncoderRef, certsOut *corefoundation.CFArrayRef) (int32, error) {
	if _cMSEncoderCopySupportingCerts == nil {
		return 0, symbolCallError("CMSEncoderCopySupportingCerts", "10.5", _cMSEncoderCopySupportingCertsErr)
	}
	return _cMSEncoderCopySupportingCerts(cmsEncoder, certsOut), nil
}

// CMSEncoderCopySupportingCerts obtains the certificates added to a message with [CMSEncoderAddSupportingCerts].
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderCopySupportingCerts(_:_:)
func CMSEncoderCopySupportingCerts(cmsEncoder CMSEncoderRef, certsOut *corefoundation.CFArrayRef) int32 {
	result, callErr := tryCMSEncoderCopySupportingCerts(cmsEncoder, certsOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderCreate func(cmsEncoderOut *CMSEncoderRef) int32
var _cMSEncoderCreateErr error

func tryCMSEncoderCreate(cmsEncoderOut *CMSEncoderRef) (int32, error) {
	if _cMSEncoderCreate == nil {
		return 0, symbolCallError("CMSEncoderCreate", "10.5", _cMSEncoderCreateErr)
	}
	return _cMSEncoderCreate(cmsEncoderOut), nil
}

// CMSEncoderCreate creates a CMSEncoder reference.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderCreate(_:)
func CMSEncoderCreate(cmsEncoderOut *CMSEncoderRef) int32 {
	result, callErr := tryCMSEncoderCreate(cmsEncoderOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderGetCertificateChainMode func(cmsEncoder CMSEncoderRef, chainModeOut *CMSCertificateChainMode) int32
var _cMSEncoderGetCertificateChainModeErr error

func tryCMSEncoderGetCertificateChainMode(cmsEncoder CMSEncoderRef, chainModeOut *CMSCertificateChainMode) (int32, error) {
	if _cMSEncoderGetCertificateChainMode == nil {
		return 0, symbolCallError("CMSEncoderGetCertificateChainMode", "10.5", _cMSEncoderGetCertificateChainModeErr)
	}
	return _cMSEncoderGetCertificateChainMode(cmsEncoder, chainModeOut), nil
}

// CMSEncoderGetCertificateChainMode obtains a constant that indicates which certificates are to be included in a signed CMS message.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderGetCertificateChainMode(_:_:)
func CMSEncoderGetCertificateChainMode(cmsEncoder CMSEncoderRef, chainModeOut *CMSCertificateChainMode) int32 {
	result, callErr := tryCMSEncoderGetCertificateChainMode(cmsEncoder, chainModeOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderGetHasDetachedContent func(cmsEncoder CMSEncoderRef, detachedContentOut *bool) int32
var _cMSEncoderGetHasDetachedContentErr error

func tryCMSEncoderGetHasDetachedContent(cmsEncoder CMSEncoderRef, detachedContentOut *bool) (int32, error) {
	if _cMSEncoderGetHasDetachedContent == nil {
		return 0, symbolCallError("CMSEncoderGetHasDetachedContent", "10.5", _cMSEncoderGetHasDetachedContentErr)
	}
	return _cMSEncoderGetHasDetachedContent(cmsEncoder, detachedContentOut), nil
}

// CMSEncoderGetHasDetachedContent indicates whether the message is to have detached content.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderGetHasDetachedContent(_:_:)
func CMSEncoderGetHasDetachedContent(cmsEncoder CMSEncoderRef, detachedContentOut *bool) int32 {
	result, callErr := tryCMSEncoderGetHasDetachedContent(cmsEncoder, detachedContentOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderGetTypeID func() uint
var _cMSEncoderGetTypeIDErr error

func tryCMSEncoderGetTypeID() (uint, error) {
	if _cMSEncoderGetTypeID == nil {
		return 0, symbolCallError("CMSEncoderGetTypeID", "10.5", _cMSEncoderGetTypeIDErr)
	}
	return _cMSEncoderGetTypeID(), nil
}

// CMSEncoderGetTypeID returns the type identifier for the CMSEncoder opaque type.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderGetTypeID()
func CMSEncoderGetTypeID() uint {
	result, callErr := tryCMSEncoderGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderSetCertificateChainMode func(cmsEncoder CMSEncoderRef, chainMode CMSCertificateChainMode) int32
var _cMSEncoderSetCertificateChainModeErr error

func tryCMSEncoderSetCertificateChainMode(cmsEncoder CMSEncoderRef, chainMode CMSCertificateChainMode) (int32, error) {
	if _cMSEncoderSetCertificateChainMode == nil {
		return 0, symbolCallError("CMSEncoderSetCertificateChainMode", "10.5", _cMSEncoderSetCertificateChainModeErr)
	}
	return _cMSEncoderSetCertificateChainMode(cmsEncoder, chainMode), nil
}

// CMSEncoderSetCertificateChainMode specifies which certificates to include in a signed CMS message.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderSetCertificateChainMode(_:_:)
func CMSEncoderSetCertificateChainMode(cmsEncoder CMSEncoderRef, chainMode CMSCertificateChainMode) int32 {
	result, callErr := tryCMSEncoderSetCertificateChainMode(cmsEncoder, chainMode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderSetEncapsulatedContentType func(cmsEncoder CMSEncoderRef, eContentType unsafe.Pointer) int32
var _cMSEncoderSetEncapsulatedContentTypeErr error

func tryCMSEncoderSetEncapsulatedContentType(cmsEncoder CMSEncoderRef, eContentType unsafe.Pointer) (int32, error) {
	if _cMSEncoderSetEncapsulatedContentType == nil {
		return 0, symbolCallError("CMSEncoderSetEncapsulatedContentType", "10.5", _cMSEncoderSetEncapsulatedContentTypeErr)
	}
	return _cMSEncoderSetEncapsulatedContentType(cmsEncoder, eContentType), nil
}

// CMSEncoderSetEncapsulatedContentType specifies an object identifier for the encapsulated data of a signed message.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderSetEncapsulatedContentType
func CMSEncoderSetEncapsulatedContentType(cmsEncoder CMSEncoderRef, eContentType unsafe.Pointer) int32 {
	result, callErr := tryCMSEncoderSetEncapsulatedContentType(cmsEncoder, eContentType)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderSetEncapsulatedContentTypeOID func(cmsEncoder CMSEncoderRef, eContentTypeOID corefoundation.CFTypeRef) int32
var _cMSEncoderSetEncapsulatedContentTypeOIDErr error

func tryCMSEncoderSetEncapsulatedContentTypeOID(cmsEncoder CMSEncoderRef, eContentTypeOID corefoundation.CFTypeRef) (int32, error) {
	if _cMSEncoderSetEncapsulatedContentTypeOID == nil {
		return 0, symbolCallError("CMSEncoderSetEncapsulatedContentTypeOID", "10.7", _cMSEncoderSetEncapsulatedContentTypeOIDErr)
	}
	return _cMSEncoderSetEncapsulatedContentTypeOID(cmsEncoder, eContentTypeOID), nil
}

// CMSEncoderSetEncapsulatedContentTypeOID specifies an object identifier for the encapsulated data of a signed message.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderSetEncapsulatedContentTypeOID(_:_:)
func CMSEncoderSetEncapsulatedContentTypeOID(cmsEncoder CMSEncoderRef, eContentTypeOID corefoundation.CFTypeRef) int32 {
	result, callErr := tryCMSEncoderSetEncapsulatedContentTypeOID(cmsEncoder, eContentTypeOID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderSetHasDetachedContent func(cmsEncoder CMSEncoderRef, detachedContent bool) int32
var _cMSEncoderSetHasDetachedContentErr error

func tryCMSEncoderSetHasDetachedContent(cmsEncoder CMSEncoderRef, detachedContent bool) (int32, error) {
	if _cMSEncoderSetHasDetachedContent == nil {
		return 0, symbolCallError("CMSEncoderSetHasDetachedContent", "10.5", _cMSEncoderSetHasDetachedContentErr)
	}
	return _cMSEncoderSetHasDetachedContent(cmsEncoder, detachedContent), nil
}

// CMSEncoderSetHasDetachedContent specifies whether the signed data is to be separate from the message.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderSetHasDetachedContent(_:_:)
func CMSEncoderSetHasDetachedContent(cmsEncoder CMSEncoderRef, detachedContent bool) int32 {
	result, callErr := tryCMSEncoderSetHasDetachedContent(cmsEncoder, detachedContent)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderSetSignerAlgorithm func(cmsEncoder CMSEncoderRef, digestAlgorithm corefoundation.CFStringRef) int32
var _cMSEncoderSetSignerAlgorithmErr error

func tryCMSEncoderSetSignerAlgorithm(cmsEncoder CMSEncoderRef, digestAlgorithm corefoundation.CFStringRef) (int32, error) {
	if _cMSEncoderSetSignerAlgorithm == nil {
		return 0, symbolCallError("CMSEncoderSetSignerAlgorithm", "10.11", _cMSEncoderSetSignerAlgorithmErr)
	}
	return _cMSEncoderSetSignerAlgorithm(cmsEncoder, digestAlgorithm), nil
}

// CMSEncoderSetSignerAlgorithm sets the digest algorithm to use for the signer.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderSetSignerAlgorithm(_:_:)
func CMSEncoderSetSignerAlgorithm(cmsEncoder CMSEncoderRef, digestAlgorithm corefoundation.CFStringRef) int32 {
	result, callErr := tryCMSEncoderSetSignerAlgorithm(cmsEncoder, digestAlgorithm)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderUpdateContent func(cmsEncoder CMSEncoderRef, content unsafe.Pointer, contentLen uintptr) int32
var _cMSEncoderUpdateContentErr error

func tryCMSEncoderUpdateContent(cmsEncoder CMSEncoderRef, content unsafe.Pointer, contentLen uintptr) (int32, error) {
	if _cMSEncoderUpdateContent == nil {
		return 0, symbolCallError("CMSEncoderUpdateContent", "10.5", _cMSEncoderUpdateContentErr)
	}
	return _cMSEncoderUpdateContent(cmsEncoder, content, contentLen), nil
}

// CMSEncoderUpdateContent feeds content bytes into the encoder.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderUpdateContent(_:_:_:)
func CMSEncoderUpdateContent(cmsEncoder CMSEncoderRef, content unsafe.Pointer, contentLen uintptr) int32 {
	result, callErr := tryCMSEncoderUpdateContent(cmsEncoder, content, contentLen)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_AC_AuthCompute func(ACHandle CSSM_AC_HANDLE, BaseAuthorizations unsafe.Pointer, Credentials unsafe.Pointer, NumberOfRequestors uint32, Requestors unsafe.Pointer, RequestedAuthorizationPeriod unsafe.Pointer, RequestedAuthorization unsafe.Pointer, AuthorizationResult unsafe.Pointer) CSSM_RETURN
var _cSSM_AC_AuthComputeErr error

func tryCSSM_AC_AuthCompute(ACHandle CSSM_AC_HANDLE, BaseAuthorizations unsafe.Pointer, Credentials unsafe.Pointer, NumberOfRequestors uint32, Requestors unsafe.Pointer, RequestedAuthorizationPeriod unsafe.Pointer, RequestedAuthorization unsafe.Pointer, AuthorizationResult unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_AC_AuthCompute == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_AC_AuthCompute", "10.0", _cSSM_AC_AuthComputeErr)
	}
	return _cSSM_AC_AuthCompute(ACHandle, BaseAuthorizations, Credentials, NumberOfRequestors, Requestors, RequestedAuthorizationPeriod, RequestedAuthorization, AuthorizationResult), nil
}

// CSSM_AC_AuthCompute.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_AC_AuthCompute
func CSSM_AC_AuthCompute(ACHandle CSSM_AC_HANDLE, BaseAuthorizations unsafe.Pointer, Credentials unsafe.Pointer, NumberOfRequestors uint32, Requestors unsafe.Pointer, RequestedAuthorizationPeriod unsafe.Pointer, RequestedAuthorization unsafe.Pointer, AuthorizationResult unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_AC_AuthCompute(ACHandle, BaseAuthorizations, Credentials, NumberOfRequestors, Requestors, RequestedAuthorizationPeriod, RequestedAuthorization, AuthorizationResult)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_AC_PassThrough func(ACHandle CSSM_AC_HANDLE, TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, DBList unsafe.Pointer, PassThroughId uint32, InputParams unsafe.Pointer, OutputParams unsafe.Pointer) CSSM_RETURN
var _cSSM_AC_PassThroughErr error

func tryCSSM_AC_PassThrough(ACHandle CSSM_AC_HANDLE, TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, DBList unsafe.Pointer, PassThroughId uint32, InputParams unsafe.Pointer, OutputParams unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_AC_PassThrough == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_AC_PassThrough", "10.0", _cSSM_AC_PassThroughErr)
	}
	return _cSSM_AC_PassThrough(ACHandle, TPHandle, CLHandle, CCHandle, DBList, PassThroughId, InputParams, OutputParams), nil
}

// CSSM_AC_PassThrough.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_AC_PassThrough
func CSSM_AC_PassThrough(ACHandle CSSM_AC_HANDLE, TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, DBList unsafe.Pointer, PassThroughId uint32, InputParams unsafe.Pointer, OutputParams unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_AC_PassThrough(ACHandle, TPHandle, CLHandle, CCHandle, DBList, PassThroughId, InputParams, OutputParams)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CertAbortCache func(CLHandle CSSM_CL_HANDLE, CertHandle CSSM_HANDLE) CSSM_RETURN
var _cSSM_CL_CertAbortCacheErr error

func tryCSSM_CL_CertAbortCache(CLHandle CSSM_CL_HANDLE, CertHandle CSSM_HANDLE) (CSSM_RETURN, error) {
	if _cSSM_CL_CertAbortCache == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CL_CertAbortCache", "10.0", _cSSM_CL_CertAbortCacheErr)
	}
	return _cSSM_CL_CertAbortCache(CLHandle, CertHandle), nil
}

// CSSM_CL_CertAbortCache.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertAbortCache
func CSSM_CL_CertAbortCache(CLHandle CSSM_CL_HANDLE, CertHandle CSSM_HANDLE) CSSM_RETURN {
	result, callErr := tryCSSM_CL_CertAbortCache(CLHandle, CertHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CertAbortQuery func(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE) CSSM_RETURN
var _cSSM_CL_CertAbortQueryErr error

func tryCSSM_CL_CertAbortQuery(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE) (CSSM_RETURN, error) {
	if _cSSM_CL_CertAbortQuery == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CL_CertAbortQuery", "10.0", _cSSM_CL_CertAbortQueryErr)
	}
	return _cSSM_CL_CertAbortQuery(CLHandle, ResultsHandle), nil
}

// CSSM_CL_CertAbortQuery.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertAbortQuery
func CSSM_CL_CertAbortQuery(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE) CSSM_RETURN {
	result, callErr := tryCSSM_CL_CertAbortQuery(CLHandle, ResultsHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CertCache func(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, CertHandle CSSM_HANDLE_PTR) CSSM_RETURN
var _cSSM_CL_CertCacheErr error

func tryCSSM_CL_CertCache(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, CertHandle CSSM_HANDLE_PTR) (CSSM_RETURN, error) {
	if _cSSM_CL_CertCache == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CL_CertCache", "10.0", _cSSM_CL_CertCacheErr)
	}
	return _cSSM_CL_CertCache(CLHandle, Cert, CertHandle), nil
}

// CSSM_CL_CertCache.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertCache
func CSSM_CL_CertCache(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, CertHandle CSSM_HANDLE_PTR) CSSM_RETURN {
	result, callErr := tryCSSM_CL_CertCache(CLHandle, Cert, CertHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CertCreateTemplate func(CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CertFields unsafe.Pointer, CertTemplate unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CertCreateTemplateErr error

func tryCSSM_CL_CertCreateTemplate(CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CertFields unsafe.Pointer, CertTemplate unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CL_CertCreateTemplate == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CL_CertCreateTemplate", "10.0", _cSSM_CL_CertCreateTemplateErr)
	}
	return _cSSM_CL_CertCreateTemplate(CLHandle, NumberOfFields, CertFields, CertTemplate), nil
}

// CSSM_CL_CertCreateTemplate.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertCreateTemplate
func CSSM_CL_CertCreateTemplate(CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CertFields unsafe.Pointer, CertTemplate unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CL_CertCreateTemplate(CLHandle, NumberOfFields, CertFields, CertTemplate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CertDescribeFormat func(CLHandle CSSM_CL_HANDLE, NumberOfFields *uint32, OidList unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CertDescribeFormatErr error

func tryCSSM_CL_CertDescribeFormat(CLHandle CSSM_CL_HANDLE, NumberOfFields *uint32, OidList unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CL_CertDescribeFormat == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CL_CertDescribeFormat", "10.0", _cSSM_CL_CertDescribeFormatErr)
	}
	return _cSSM_CL_CertDescribeFormat(CLHandle, NumberOfFields, OidList), nil
}

// CSSM_CL_CertDescribeFormat.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertDescribeFormat
func CSSM_CL_CertDescribeFormat(CLHandle CSSM_CL_HANDLE, NumberOfFields *uint32, OidList unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CL_CertDescribeFormat(CLHandle, NumberOfFields, OidList)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CertGetAllFields func(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, NumberOfFields *uint32, CertFields unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CertGetAllFieldsErr error

func tryCSSM_CL_CertGetAllFields(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, NumberOfFields *uint32, CertFields unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CL_CertGetAllFields == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CL_CertGetAllFields", "10.0", _cSSM_CL_CertGetAllFieldsErr)
	}
	return _cSSM_CL_CertGetAllFields(CLHandle, Cert, NumberOfFields, CertFields), nil
}

// CSSM_CL_CertGetAllFields.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertGetAllFields
func CSSM_CL_CertGetAllFields(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, NumberOfFields *uint32, CertFields unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CL_CertGetAllFields(CLHandle, Cert, NumberOfFields, CertFields)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CertGetAllTemplateFields func(CLHandle CSSM_CL_HANDLE, CertTemplate unsafe.Pointer, NumberOfFields *uint32, CertFields unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CertGetAllTemplateFieldsErr error

func tryCSSM_CL_CertGetAllTemplateFields(CLHandle CSSM_CL_HANDLE, CertTemplate unsafe.Pointer, NumberOfFields *uint32, CertFields unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CL_CertGetAllTemplateFields == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CL_CertGetAllTemplateFields", "10.0", _cSSM_CL_CertGetAllTemplateFieldsErr)
	}
	return _cSSM_CL_CertGetAllTemplateFields(CLHandle, CertTemplate, NumberOfFields, CertFields), nil
}

// CSSM_CL_CertGetAllTemplateFields.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertGetAllTemplateFields
func CSSM_CL_CertGetAllTemplateFields(CLHandle CSSM_CL_HANDLE, CertTemplate unsafe.Pointer, NumberOfFields *uint32, CertFields unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CL_CertGetAllTemplateFields(CLHandle, CertTemplate, NumberOfFields, CertFields)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CertGetFirstCachedFieldValue func(CLHandle CSSM_CL_HANDLE, CertHandle CSSM_HANDLE, CertField unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, NumberOfMatchedFields *uint32, Value unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CertGetFirstCachedFieldValueErr error

func tryCSSM_CL_CertGetFirstCachedFieldValue(CLHandle CSSM_CL_HANDLE, CertHandle CSSM_HANDLE, CertField unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, NumberOfMatchedFields *uint32, Value unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CL_CertGetFirstCachedFieldValue == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CL_CertGetFirstCachedFieldValue", "10.0", _cSSM_CL_CertGetFirstCachedFieldValueErr)
	}
	return _cSSM_CL_CertGetFirstCachedFieldValue(CLHandle, CertHandle, CertField, ResultsHandle, NumberOfMatchedFields, Value), nil
}

// CSSM_CL_CertGetFirstCachedFieldValue.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertGetFirstCachedFieldValue
func CSSM_CL_CertGetFirstCachedFieldValue(CLHandle CSSM_CL_HANDLE, CertHandle CSSM_HANDLE, CertField unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, NumberOfMatchedFields *uint32, Value unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CL_CertGetFirstCachedFieldValue(CLHandle, CertHandle, CertField, ResultsHandle, NumberOfMatchedFields, Value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CertGetFirstFieldValue func(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, CertField unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, NumberOfMatchedFields *uint32, Value unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CertGetFirstFieldValueErr error

func tryCSSM_CL_CertGetFirstFieldValue(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, CertField unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, NumberOfMatchedFields *uint32, Value unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CL_CertGetFirstFieldValue == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CL_CertGetFirstFieldValue", "10.0", _cSSM_CL_CertGetFirstFieldValueErr)
	}
	return _cSSM_CL_CertGetFirstFieldValue(CLHandle, Cert, CertField, ResultsHandle, NumberOfMatchedFields, Value), nil
}

// CSSM_CL_CertGetFirstFieldValue.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertGetFirstFieldValue
func CSSM_CL_CertGetFirstFieldValue(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, CertField unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, NumberOfMatchedFields *uint32, Value unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CL_CertGetFirstFieldValue(CLHandle, Cert, CertField, ResultsHandle, NumberOfMatchedFields, Value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CertGetKeyInfo func(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, Key unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CertGetKeyInfoErr error

func tryCSSM_CL_CertGetKeyInfo(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, Key unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CL_CertGetKeyInfo == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CL_CertGetKeyInfo", "10.0", _cSSM_CL_CertGetKeyInfoErr)
	}
	return _cSSM_CL_CertGetKeyInfo(CLHandle, Cert, Key), nil
}

// CSSM_CL_CertGetKeyInfo.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertGetKeyInfo
func CSSM_CL_CertGetKeyInfo(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, Key unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CL_CertGetKeyInfo(CLHandle, Cert, Key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CertGetNextCachedFieldValue func(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE, Value unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CertGetNextCachedFieldValueErr error

func tryCSSM_CL_CertGetNextCachedFieldValue(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE, Value unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CL_CertGetNextCachedFieldValue == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CL_CertGetNextCachedFieldValue", "10.0", _cSSM_CL_CertGetNextCachedFieldValueErr)
	}
	return _cSSM_CL_CertGetNextCachedFieldValue(CLHandle, ResultsHandle, Value), nil
}

// CSSM_CL_CertGetNextCachedFieldValue.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertGetNextCachedFieldValue
func CSSM_CL_CertGetNextCachedFieldValue(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE, Value unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CL_CertGetNextCachedFieldValue(CLHandle, ResultsHandle, Value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CertGetNextFieldValue func(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE, Value unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CertGetNextFieldValueErr error

func tryCSSM_CL_CertGetNextFieldValue(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE, Value unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CL_CertGetNextFieldValue == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CL_CertGetNextFieldValue", "10.0", _cSSM_CL_CertGetNextFieldValueErr)
	}
	return _cSSM_CL_CertGetNextFieldValue(CLHandle, ResultsHandle, Value), nil
}

// CSSM_CL_CertGetNextFieldValue.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertGetNextFieldValue
func CSSM_CL_CertGetNextFieldValue(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE, Value unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CL_CertGetNextFieldValue(CLHandle, ResultsHandle, Value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CertGroupFromVerifiedBundle func(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertBundle unsafe.Pointer, SignerCert unsafe.Pointer, CertGroup unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CertGroupFromVerifiedBundleErr error

func tryCSSM_CL_CertGroupFromVerifiedBundle(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertBundle unsafe.Pointer, SignerCert unsafe.Pointer, CertGroup unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CL_CertGroupFromVerifiedBundle == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CL_CertGroupFromVerifiedBundle", "10.0", _cSSM_CL_CertGroupFromVerifiedBundleErr)
	}
	return _cSSM_CL_CertGroupFromVerifiedBundle(CLHandle, CCHandle, CertBundle, SignerCert, CertGroup), nil
}

// CSSM_CL_CertGroupFromVerifiedBundle.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertGroupFromVerifiedBundle
func CSSM_CL_CertGroupFromVerifiedBundle(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertBundle unsafe.Pointer, SignerCert unsafe.Pointer, CertGroup unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CL_CertGroupFromVerifiedBundle(CLHandle, CCHandle, CertBundle, SignerCert, CertGroup)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CertGroupToSignedBundle func(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertGroupToBundle unsafe.Pointer, BundleInfo unsafe.Pointer, SignedBundle unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CertGroupToSignedBundleErr error

func tryCSSM_CL_CertGroupToSignedBundle(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertGroupToBundle unsafe.Pointer, BundleInfo unsafe.Pointer, SignedBundle unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CL_CertGroupToSignedBundle == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CL_CertGroupToSignedBundle", "10.0", _cSSM_CL_CertGroupToSignedBundleErr)
	}
	return _cSSM_CL_CertGroupToSignedBundle(CLHandle, CCHandle, CertGroupToBundle, BundleInfo, SignedBundle), nil
}

// CSSM_CL_CertGroupToSignedBundle.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertGroupToSignedBundle
func CSSM_CL_CertGroupToSignedBundle(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertGroupToBundle unsafe.Pointer, BundleInfo unsafe.Pointer, SignedBundle unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CL_CertGroupToSignedBundle(CLHandle, CCHandle, CertGroupToBundle, BundleInfo, SignedBundle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CertSign func(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertTemplate unsafe.Pointer, SignScope unsafe.Pointer, ScopeSize uint32, SignedCert unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CertSignErr error

func tryCSSM_CL_CertSign(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertTemplate unsafe.Pointer, SignScope unsafe.Pointer, ScopeSize uint32, SignedCert unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CL_CertSign == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CL_CertSign", "10.0", _cSSM_CL_CertSignErr)
	}
	return _cSSM_CL_CertSign(CLHandle, CCHandle, CertTemplate, SignScope, ScopeSize, SignedCert), nil
}

// CSSM_CL_CertSign.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertSign
func CSSM_CL_CertSign(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertTemplate unsafe.Pointer, SignScope unsafe.Pointer, ScopeSize uint32, SignedCert unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CL_CertSign(CLHandle, CCHandle, CertTemplate, SignScope, ScopeSize, SignedCert)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CertVerify func(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertToBeVerified unsafe.Pointer, SignerCert unsafe.Pointer, VerifyScope unsafe.Pointer, ScopeSize uint32) CSSM_RETURN
var _cSSM_CL_CertVerifyErr error

func tryCSSM_CL_CertVerify(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertToBeVerified unsafe.Pointer, SignerCert unsafe.Pointer, VerifyScope unsafe.Pointer, ScopeSize uint32) (CSSM_RETURN, error) {
	if _cSSM_CL_CertVerify == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CL_CertVerify", "10.0", _cSSM_CL_CertVerifyErr)
	}
	return _cSSM_CL_CertVerify(CLHandle, CCHandle, CertToBeVerified, SignerCert, VerifyScope, ScopeSize), nil
}

// CSSM_CL_CertVerify.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertVerify
func CSSM_CL_CertVerify(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertToBeVerified unsafe.Pointer, SignerCert unsafe.Pointer, VerifyScope unsafe.Pointer, ScopeSize uint32) CSSM_RETURN {
	result, callErr := tryCSSM_CL_CertVerify(CLHandle, CCHandle, CertToBeVerified, SignerCert, VerifyScope, ScopeSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CertVerifyWithKey func(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertToBeVerified unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CertVerifyWithKeyErr error

func tryCSSM_CL_CertVerifyWithKey(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertToBeVerified unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CL_CertVerifyWithKey == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CL_CertVerifyWithKey", "10.0", _cSSM_CL_CertVerifyWithKeyErr)
	}
	return _cSSM_CL_CertVerifyWithKey(CLHandle, CCHandle, CertToBeVerified), nil
}

// CSSM_CL_CertVerifyWithKey.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertVerifyWithKey
func CSSM_CL_CertVerifyWithKey(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertToBeVerified unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CL_CertVerifyWithKey(CLHandle, CCHandle, CertToBeVerified)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CrlAbortCache func(CLHandle CSSM_CL_HANDLE, CrlHandle CSSM_HANDLE) CSSM_RETURN
var _cSSM_CL_CrlAbortCacheErr error

func tryCSSM_CL_CrlAbortCache(CLHandle CSSM_CL_HANDLE, CrlHandle CSSM_HANDLE) (CSSM_RETURN, error) {
	if _cSSM_CL_CrlAbortCache == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CL_CrlAbortCache", "10.0", _cSSM_CL_CrlAbortCacheErr)
	}
	return _cSSM_CL_CrlAbortCache(CLHandle, CrlHandle), nil
}

// CSSM_CL_CrlAbortCache.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlAbortCache
func CSSM_CL_CrlAbortCache(CLHandle CSSM_CL_HANDLE, CrlHandle CSSM_HANDLE) CSSM_RETURN {
	result, callErr := tryCSSM_CL_CrlAbortCache(CLHandle, CrlHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CrlAbortQuery func(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE) CSSM_RETURN
var _cSSM_CL_CrlAbortQueryErr error

func tryCSSM_CL_CrlAbortQuery(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE) (CSSM_RETURN, error) {
	if _cSSM_CL_CrlAbortQuery == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CL_CrlAbortQuery", "10.0", _cSSM_CL_CrlAbortQueryErr)
	}
	return _cSSM_CL_CrlAbortQuery(CLHandle, ResultsHandle), nil
}

// CSSM_CL_CrlAbortQuery.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlAbortQuery
func CSSM_CL_CrlAbortQuery(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE) CSSM_RETURN {
	result, callErr := tryCSSM_CL_CrlAbortQuery(CLHandle, ResultsHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CrlAddCert func(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, Cert unsafe.Pointer, NumberOfFields uint32, CrlEntryFields unsafe.Pointer, OldCrl unsafe.Pointer, NewCrl unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CrlAddCertErr error

func tryCSSM_CL_CrlAddCert(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, Cert unsafe.Pointer, NumberOfFields uint32, CrlEntryFields unsafe.Pointer, OldCrl unsafe.Pointer, NewCrl unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CL_CrlAddCert == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CL_CrlAddCert", "10.0", _cSSM_CL_CrlAddCertErr)
	}
	return _cSSM_CL_CrlAddCert(CLHandle, CCHandle, Cert, NumberOfFields, CrlEntryFields, OldCrl, NewCrl), nil
}

// CSSM_CL_CrlAddCert.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlAddCert
func CSSM_CL_CrlAddCert(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, Cert unsafe.Pointer, NumberOfFields uint32, CrlEntryFields unsafe.Pointer, OldCrl unsafe.Pointer, NewCrl unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CL_CrlAddCert(CLHandle, CCHandle, Cert, NumberOfFields, CrlEntryFields, OldCrl, NewCrl)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CrlCache func(CLHandle CSSM_CL_HANDLE, Crl unsafe.Pointer, CrlHandle CSSM_HANDLE_PTR) CSSM_RETURN
var _cSSM_CL_CrlCacheErr error

func tryCSSM_CL_CrlCache(CLHandle CSSM_CL_HANDLE, Crl unsafe.Pointer, CrlHandle CSSM_HANDLE_PTR) (CSSM_RETURN, error) {
	if _cSSM_CL_CrlCache == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CL_CrlCache", "10.0", _cSSM_CL_CrlCacheErr)
	}
	return _cSSM_CL_CrlCache(CLHandle, Crl, CrlHandle), nil
}

// CSSM_CL_CrlCache.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlCache
func CSSM_CL_CrlCache(CLHandle CSSM_CL_HANDLE, Crl unsafe.Pointer, CrlHandle CSSM_HANDLE_PTR) CSSM_RETURN {
	result, callErr := tryCSSM_CL_CrlCache(CLHandle, Crl, CrlHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CrlCreateTemplate func(CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CrlTemplate unsafe.Pointer, NewCrl unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CrlCreateTemplateErr error

func tryCSSM_CL_CrlCreateTemplate(CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CrlTemplate unsafe.Pointer, NewCrl unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CL_CrlCreateTemplate == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CL_CrlCreateTemplate", "10.0", _cSSM_CL_CrlCreateTemplateErr)
	}
	return _cSSM_CL_CrlCreateTemplate(CLHandle, NumberOfFields, CrlTemplate, NewCrl), nil
}

// CSSM_CL_CrlCreateTemplate.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlCreateTemplate
func CSSM_CL_CrlCreateTemplate(CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CrlTemplate unsafe.Pointer, NewCrl unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CL_CrlCreateTemplate(CLHandle, NumberOfFields, CrlTemplate, NewCrl)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CrlDescribeFormat func(CLHandle CSSM_CL_HANDLE, NumberOfFields *uint32, OidList unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CrlDescribeFormatErr error

func tryCSSM_CL_CrlDescribeFormat(CLHandle CSSM_CL_HANDLE, NumberOfFields *uint32, OidList unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CL_CrlDescribeFormat == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CL_CrlDescribeFormat", "10.0", _cSSM_CL_CrlDescribeFormatErr)
	}
	return _cSSM_CL_CrlDescribeFormat(CLHandle, NumberOfFields, OidList), nil
}

// CSSM_CL_CrlDescribeFormat.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlDescribeFormat
func CSSM_CL_CrlDescribeFormat(CLHandle CSSM_CL_HANDLE, NumberOfFields *uint32, OidList unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CL_CrlDescribeFormat(CLHandle, NumberOfFields, OidList)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CrlGetAllCachedRecordFields func(CLHandle CSSM_CL_HANDLE, CrlHandle CSSM_HANDLE, CrlRecordIndex unsafe.Pointer, NumberOfFields *uint32, CrlFields unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CrlGetAllCachedRecordFieldsErr error

func tryCSSM_CL_CrlGetAllCachedRecordFields(CLHandle CSSM_CL_HANDLE, CrlHandle CSSM_HANDLE, CrlRecordIndex unsafe.Pointer, NumberOfFields *uint32, CrlFields unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CL_CrlGetAllCachedRecordFields == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CL_CrlGetAllCachedRecordFields", "10.0", _cSSM_CL_CrlGetAllCachedRecordFieldsErr)
	}
	return _cSSM_CL_CrlGetAllCachedRecordFields(CLHandle, CrlHandle, CrlRecordIndex, NumberOfFields, CrlFields), nil
}

// CSSM_CL_CrlGetAllCachedRecordFields.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlGetAllCachedRecordFields
func CSSM_CL_CrlGetAllCachedRecordFields(CLHandle CSSM_CL_HANDLE, CrlHandle CSSM_HANDLE, CrlRecordIndex unsafe.Pointer, NumberOfFields *uint32, CrlFields unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CL_CrlGetAllCachedRecordFields(CLHandle, CrlHandle, CrlRecordIndex, NumberOfFields, CrlFields)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CrlGetAllFields func(CLHandle CSSM_CL_HANDLE, Crl unsafe.Pointer, NumberOfCrlFields *uint32, CrlFields unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CrlGetAllFieldsErr error

func tryCSSM_CL_CrlGetAllFields(CLHandle CSSM_CL_HANDLE, Crl unsafe.Pointer, NumberOfCrlFields *uint32, CrlFields unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CL_CrlGetAllFields == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CL_CrlGetAllFields", "10.0", _cSSM_CL_CrlGetAllFieldsErr)
	}
	return _cSSM_CL_CrlGetAllFields(CLHandle, Crl, NumberOfCrlFields, CrlFields), nil
}

// CSSM_CL_CrlGetAllFields.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlGetAllFields
func CSSM_CL_CrlGetAllFields(CLHandle CSSM_CL_HANDLE, Crl unsafe.Pointer, NumberOfCrlFields *uint32, CrlFields unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CL_CrlGetAllFields(CLHandle, Crl, NumberOfCrlFields, CrlFields)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CrlGetFirstCachedFieldValue func(CLHandle CSSM_CL_HANDLE, CrlHandle CSSM_HANDLE, CrlRecordIndex unsafe.Pointer, CrlField unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, NumberOfMatchedFields *uint32, Value unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CrlGetFirstCachedFieldValueErr error

func tryCSSM_CL_CrlGetFirstCachedFieldValue(CLHandle CSSM_CL_HANDLE, CrlHandle CSSM_HANDLE, CrlRecordIndex unsafe.Pointer, CrlField unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, NumberOfMatchedFields *uint32, Value unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CL_CrlGetFirstCachedFieldValue == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CL_CrlGetFirstCachedFieldValue", "10.0", _cSSM_CL_CrlGetFirstCachedFieldValueErr)
	}
	return _cSSM_CL_CrlGetFirstCachedFieldValue(CLHandle, CrlHandle, CrlRecordIndex, CrlField, ResultsHandle, NumberOfMatchedFields, Value), nil
}

// CSSM_CL_CrlGetFirstCachedFieldValue.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlGetFirstCachedFieldValue
func CSSM_CL_CrlGetFirstCachedFieldValue(CLHandle CSSM_CL_HANDLE, CrlHandle CSSM_HANDLE, CrlRecordIndex unsafe.Pointer, CrlField unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, NumberOfMatchedFields *uint32, Value unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CL_CrlGetFirstCachedFieldValue(CLHandle, CrlHandle, CrlRecordIndex, CrlField, ResultsHandle, NumberOfMatchedFields, Value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CrlGetFirstFieldValue func(CLHandle CSSM_CL_HANDLE, Crl unsafe.Pointer, CrlField unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, NumberOfMatchedFields *uint32, Value unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CrlGetFirstFieldValueErr error

func tryCSSM_CL_CrlGetFirstFieldValue(CLHandle CSSM_CL_HANDLE, Crl unsafe.Pointer, CrlField unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, NumberOfMatchedFields *uint32, Value unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CL_CrlGetFirstFieldValue == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CL_CrlGetFirstFieldValue", "10.0", _cSSM_CL_CrlGetFirstFieldValueErr)
	}
	return _cSSM_CL_CrlGetFirstFieldValue(CLHandle, Crl, CrlField, ResultsHandle, NumberOfMatchedFields, Value), nil
}

// CSSM_CL_CrlGetFirstFieldValue.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlGetFirstFieldValue
func CSSM_CL_CrlGetFirstFieldValue(CLHandle CSSM_CL_HANDLE, Crl unsafe.Pointer, CrlField unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, NumberOfMatchedFields *uint32, Value unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CL_CrlGetFirstFieldValue(CLHandle, Crl, CrlField, ResultsHandle, NumberOfMatchedFields, Value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CrlGetNextCachedFieldValue func(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE, Value unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CrlGetNextCachedFieldValueErr error

func tryCSSM_CL_CrlGetNextCachedFieldValue(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE, Value unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CL_CrlGetNextCachedFieldValue == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CL_CrlGetNextCachedFieldValue", "10.0", _cSSM_CL_CrlGetNextCachedFieldValueErr)
	}
	return _cSSM_CL_CrlGetNextCachedFieldValue(CLHandle, ResultsHandle, Value), nil
}

// CSSM_CL_CrlGetNextCachedFieldValue.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlGetNextCachedFieldValue
func CSSM_CL_CrlGetNextCachedFieldValue(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE, Value unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CL_CrlGetNextCachedFieldValue(CLHandle, ResultsHandle, Value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CrlGetNextFieldValue func(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE, Value unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CrlGetNextFieldValueErr error

func tryCSSM_CL_CrlGetNextFieldValue(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE, Value unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CL_CrlGetNextFieldValue == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CL_CrlGetNextFieldValue", "10.0", _cSSM_CL_CrlGetNextFieldValueErr)
	}
	return _cSSM_CL_CrlGetNextFieldValue(CLHandle, ResultsHandle, Value), nil
}

// CSSM_CL_CrlGetNextFieldValue.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlGetNextFieldValue
func CSSM_CL_CrlGetNextFieldValue(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE, Value unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CL_CrlGetNextFieldValue(CLHandle, ResultsHandle, Value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CrlRemoveCert func(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, OldCrl unsafe.Pointer, NewCrl unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CrlRemoveCertErr error

func tryCSSM_CL_CrlRemoveCert(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, OldCrl unsafe.Pointer, NewCrl unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CL_CrlRemoveCert == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CL_CrlRemoveCert", "10.0", _cSSM_CL_CrlRemoveCertErr)
	}
	return _cSSM_CL_CrlRemoveCert(CLHandle, Cert, OldCrl, NewCrl), nil
}

// CSSM_CL_CrlRemoveCert.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlRemoveCert
func CSSM_CL_CrlRemoveCert(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, OldCrl unsafe.Pointer, NewCrl unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CL_CrlRemoveCert(CLHandle, Cert, OldCrl, NewCrl)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CrlSetFields func(CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CrlTemplate unsafe.Pointer, OldCrl unsafe.Pointer, ModifiedCrl unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CrlSetFieldsErr error

func tryCSSM_CL_CrlSetFields(CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CrlTemplate unsafe.Pointer, OldCrl unsafe.Pointer, ModifiedCrl unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CL_CrlSetFields == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CL_CrlSetFields", "10.0", _cSSM_CL_CrlSetFieldsErr)
	}
	return _cSSM_CL_CrlSetFields(CLHandle, NumberOfFields, CrlTemplate, OldCrl, ModifiedCrl), nil
}

// CSSM_CL_CrlSetFields.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlSetFields
func CSSM_CL_CrlSetFields(CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CrlTemplate unsafe.Pointer, OldCrl unsafe.Pointer, ModifiedCrl unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CL_CrlSetFields(CLHandle, NumberOfFields, CrlTemplate, OldCrl, ModifiedCrl)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CrlSign func(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, UnsignedCrl unsafe.Pointer, SignScope unsafe.Pointer, ScopeSize uint32, SignedCrl unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CrlSignErr error

func tryCSSM_CL_CrlSign(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, UnsignedCrl unsafe.Pointer, SignScope unsafe.Pointer, ScopeSize uint32, SignedCrl unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CL_CrlSign == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CL_CrlSign", "10.0", _cSSM_CL_CrlSignErr)
	}
	return _cSSM_CL_CrlSign(CLHandle, CCHandle, UnsignedCrl, SignScope, ScopeSize, SignedCrl), nil
}

// CSSM_CL_CrlSign.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlSign
func CSSM_CL_CrlSign(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, UnsignedCrl unsafe.Pointer, SignScope unsafe.Pointer, ScopeSize uint32, SignedCrl unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CL_CrlSign(CLHandle, CCHandle, UnsignedCrl, SignScope, ScopeSize, SignedCrl)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CrlVerify func(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CrlToBeVerified unsafe.Pointer, SignerCert unsafe.Pointer, VerifyScope unsafe.Pointer, ScopeSize uint32) CSSM_RETURN
var _cSSM_CL_CrlVerifyErr error

func tryCSSM_CL_CrlVerify(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CrlToBeVerified unsafe.Pointer, SignerCert unsafe.Pointer, VerifyScope unsafe.Pointer, ScopeSize uint32) (CSSM_RETURN, error) {
	if _cSSM_CL_CrlVerify == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CL_CrlVerify", "10.0", _cSSM_CL_CrlVerifyErr)
	}
	return _cSSM_CL_CrlVerify(CLHandle, CCHandle, CrlToBeVerified, SignerCert, VerifyScope, ScopeSize), nil
}

// CSSM_CL_CrlVerify.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlVerify
func CSSM_CL_CrlVerify(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CrlToBeVerified unsafe.Pointer, SignerCert unsafe.Pointer, VerifyScope unsafe.Pointer, ScopeSize uint32) CSSM_RETURN {
	result, callErr := tryCSSM_CL_CrlVerify(CLHandle, CCHandle, CrlToBeVerified, SignerCert, VerifyScope, ScopeSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CrlVerifyWithKey func(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CrlToBeVerified unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CrlVerifyWithKeyErr error

func tryCSSM_CL_CrlVerifyWithKey(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CrlToBeVerified unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CL_CrlVerifyWithKey == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CL_CrlVerifyWithKey", "10.0", _cSSM_CL_CrlVerifyWithKeyErr)
	}
	return _cSSM_CL_CrlVerifyWithKey(CLHandle, CCHandle, CrlToBeVerified), nil
}

// CSSM_CL_CrlVerifyWithKey.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlVerifyWithKey
func CSSM_CL_CrlVerifyWithKey(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CrlToBeVerified unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CL_CrlVerifyWithKey(CLHandle, CCHandle, CrlToBeVerified)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_FreeFieldValue func(CLHandle CSSM_CL_HANDLE, CertOrCrlOid unsafe.Pointer, Value unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_FreeFieldValueErr error

func tryCSSM_CL_FreeFieldValue(CLHandle CSSM_CL_HANDLE, CertOrCrlOid unsafe.Pointer, Value unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CL_FreeFieldValue == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CL_FreeFieldValue", "10.0", _cSSM_CL_FreeFieldValueErr)
	}
	return _cSSM_CL_FreeFieldValue(CLHandle, CertOrCrlOid, Value), nil
}

// CSSM_CL_FreeFieldValue.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_FreeFieldValue
func CSSM_CL_FreeFieldValue(CLHandle CSSM_CL_HANDLE, CertOrCrlOid unsafe.Pointer, Value unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CL_FreeFieldValue(CLHandle, CertOrCrlOid, Value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_FreeFields func(CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, Fields unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_FreeFieldsErr error

func tryCSSM_CL_FreeFields(CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, Fields unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CL_FreeFields == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CL_FreeFields", "10.0", _cSSM_CL_FreeFieldsErr)
	}
	return _cSSM_CL_FreeFields(CLHandle, NumberOfFields, Fields), nil
}

// CSSM_CL_FreeFields.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_FreeFields
func CSSM_CL_FreeFields(CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, Fields unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CL_FreeFields(CLHandle, NumberOfFields, Fields)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_IsCertInCachedCrl func(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, CrlHandle CSSM_HANDLE, CertFound unsafe.Pointer, CrlRecordIndex unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_IsCertInCachedCrlErr error

func tryCSSM_CL_IsCertInCachedCrl(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, CrlHandle CSSM_HANDLE, CertFound unsafe.Pointer, CrlRecordIndex unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CL_IsCertInCachedCrl == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CL_IsCertInCachedCrl", "10.0", _cSSM_CL_IsCertInCachedCrlErr)
	}
	return _cSSM_CL_IsCertInCachedCrl(CLHandle, Cert, CrlHandle, CertFound, CrlRecordIndex), nil
}

// CSSM_CL_IsCertInCachedCrl.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_IsCertInCachedCrl
func CSSM_CL_IsCertInCachedCrl(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, CrlHandle CSSM_HANDLE, CertFound unsafe.Pointer, CrlRecordIndex unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CL_IsCertInCachedCrl(CLHandle, Cert, CrlHandle, CertFound, CrlRecordIndex)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_IsCertInCrl func(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, Crl unsafe.Pointer, CertFound unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_IsCertInCrlErr error

func tryCSSM_CL_IsCertInCrl(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, Crl unsafe.Pointer, CertFound unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CL_IsCertInCrl == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CL_IsCertInCrl", "10.0", _cSSM_CL_IsCertInCrlErr)
	}
	return _cSSM_CL_IsCertInCrl(CLHandle, Cert, Crl, CertFound), nil
}

// CSSM_CL_IsCertInCrl.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_IsCertInCrl
func CSSM_CL_IsCertInCrl(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, Crl unsafe.Pointer, CertFound unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CL_IsCertInCrl(CLHandle, Cert, Crl, CertFound)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_PassThrough func(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, PassThroughId uint32, InputParams unsafe.Pointer, OutputParams unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_PassThroughErr error

func tryCSSM_CL_PassThrough(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, PassThroughId uint32, InputParams unsafe.Pointer, OutputParams unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CL_PassThrough == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CL_PassThrough", "10.0", _cSSM_CL_PassThroughErr)
	}
	return _cSSM_CL_PassThrough(CLHandle, CCHandle, PassThroughId, InputParams, OutputParams), nil
}

// CSSM_CL_PassThrough.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_PassThrough
func CSSM_CL_PassThrough(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, PassThroughId uint32, InputParams unsafe.Pointer, OutputParams unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CL_PassThrough(CLHandle, CCHandle, PassThroughId, InputParams, OutputParams)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CSP_ChangeLoginAcl func(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, AclEdit unsafe.Pointer) CSSM_RETURN
var _cSSM_CSP_ChangeLoginAclErr error

func tryCSSM_CSP_ChangeLoginAcl(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, AclEdit unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CSP_ChangeLoginAcl == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CSP_ChangeLoginAcl", "10.0", _cSSM_CSP_ChangeLoginAclErr)
	}
	return _cSSM_CSP_ChangeLoginAcl(CSPHandle, AccessCred, AclEdit), nil
}

// CSSM_CSP_ChangeLoginAcl.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_ChangeLoginAcl
func CSSM_CSP_ChangeLoginAcl(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, AclEdit unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CSP_ChangeLoginAcl(CSPHandle, AccessCred, AclEdit)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CSP_ChangeLoginOwner func(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, NewOwner unsafe.Pointer) CSSM_RETURN
var _cSSM_CSP_ChangeLoginOwnerErr error

func tryCSSM_CSP_ChangeLoginOwner(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, NewOwner unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CSP_ChangeLoginOwner == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CSP_ChangeLoginOwner", "10.0", _cSSM_CSP_ChangeLoginOwnerErr)
	}
	return _cSSM_CSP_ChangeLoginOwner(CSPHandle, AccessCred, NewOwner), nil
}

// CSSM_CSP_ChangeLoginOwner.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_ChangeLoginOwner
func CSSM_CSP_ChangeLoginOwner(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, NewOwner unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CSP_ChangeLoginOwner(CSPHandle, AccessCred, NewOwner)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CSP_CreateAsymmetricContext func(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, AccessCred unsafe.Pointer, Key unsafe.Pointer, Padding CSSM_PADDING, NewContextHandle unsafe.Pointer) CSSM_RETURN
var _cSSM_CSP_CreateAsymmetricContextErr error

func tryCSSM_CSP_CreateAsymmetricContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, AccessCred unsafe.Pointer, Key unsafe.Pointer, Padding CSSM_PADDING, NewContextHandle unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CSP_CreateAsymmetricContext == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CSP_CreateAsymmetricContext", "10.0", _cSSM_CSP_CreateAsymmetricContextErr)
	}
	return _cSSM_CSP_CreateAsymmetricContext(CSPHandle, AlgorithmID, AccessCred, Key, Padding, NewContextHandle), nil
}

// CSSM_CSP_CreateAsymmetricContext.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_CreateAsymmetricContext
func CSSM_CSP_CreateAsymmetricContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, AccessCred unsafe.Pointer, Key unsafe.Pointer, Padding CSSM_PADDING, NewContextHandle unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CSP_CreateAsymmetricContext(CSPHandle, AlgorithmID, AccessCred, Key, Padding, NewContextHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CSP_CreateDeriveKeyContext func(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, DeriveKeyType CSSM_KEY_TYPE, DeriveKeyLengthInBits uint32, AccessCred unsafe.Pointer, BaseKey unsafe.Pointer, IterationCount uint32, Salt unsafe.Pointer, Seed unsafe.Pointer, NewContextHandle unsafe.Pointer) CSSM_RETURN
var _cSSM_CSP_CreateDeriveKeyContextErr error

func tryCSSM_CSP_CreateDeriveKeyContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, DeriveKeyType CSSM_KEY_TYPE, DeriveKeyLengthInBits uint32, AccessCred unsafe.Pointer, BaseKey unsafe.Pointer, IterationCount uint32, Salt unsafe.Pointer, Seed unsafe.Pointer, NewContextHandle unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CSP_CreateDeriveKeyContext == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CSP_CreateDeriveKeyContext", "10.0", _cSSM_CSP_CreateDeriveKeyContextErr)
	}
	return _cSSM_CSP_CreateDeriveKeyContext(CSPHandle, AlgorithmID, DeriveKeyType, DeriveKeyLengthInBits, AccessCred, BaseKey, IterationCount, Salt, Seed, NewContextHandle), nil
}

// CSSM_CSP_CreateDeriveKeyContext.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_CreateDeriveKeyContext
func CSSM_CSP_CreateDeriveKeyContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, DeriveKeyType CSSM_KEY_TYPE, DeriveKeyLengthInBits uint32, AccessCred unsafe.Pointer, BaseKey unsafe.Pointer, IterationCount uint32, Salt unsafe.Pointer, Seed unsafe.Pointer, NewContextHandle unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CSP_CreateDeriveKeyContext(CSPHandle, AlgorithmID, DeriveKeyType, DeriveKeyLengthInBits, AccessCred, BaseKey, IterationCount, Salt, Seed, NewContextHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CSP_CreateDigestContext func(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, NewContextHandle unsafe.Pointer) CSSM_RETURN
var _cSSM_CSP_CreateDigestContextErr error

func tryCSSM_CSP_CreateDigestContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, NewContextHandle unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CSP_CreateDigestContext == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CSP_CreateDigestContext", "10.0", _cSSM_CSP_CreateDigestContextErr)
	}
	return _cSSM_CSP_CreateDigestContext(CSPHandle, AlgorithmID, NewContextHandle), nil
}

// CSSM_CSP_CreateDigestContext.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_CreateDigestContext
func CSSM_CSP_CreateDigestContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, NewContextHandle unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CSP_CreateDigestContext(CSPHandle, AlgorithmID, NewContextHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CSP_CreateKeyGenContext func(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, KeySizeInBits uint32, Seed unsafe.Pointer, Salt unsafe.Pointer, StartDate unsafe.Pointer, EndDate unsafe.Pointer, Params unsafe.Pointer, NewContextHandle unsafe.Pointer) CSSM_RETURN
var _cSSM_CSP_CreateKeyGenContextErr error

func tryCSSM_CSP_CreateKeyGenContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, KeySizeInBits uint32, Seed unsafe.Pointer, Salt unsafe.Pointer, StartDate unsafe.Pointer, EndDate unsafe.Pointer, Params unsafe.Pointer, NewContextHandle unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CSP_CreateKeyGenContext == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CSP_CreateKeyGenContext", "10.0", _cSSM_CSP_CreateKeyGenContextErr)
	}
	return _cSSM_CSP_CreateKeyGenContext(CSPHandle, AlgorithmID, KeySizeInBits, Seed, Salt, StartDate, EndDate, Params, NewContextHandle), nil
}

// CSSM_CSP_CreateKeyGenContext.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_CreateKeyGenContext
func CSSM_CSP_CreateKeyGenContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, KeySizeInBits uint32, Seed unsafe.Pointer, Salt unsafe.Pointer, StartDate unsafe.Pointer, EndDate unsafe.Pointer, Params unsafe.Pointer, NewContextHandle unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CSP_CreateKeyGenContext(CSPHandle, AlgorithmID, KeySizeInBits, Seed, Salt, StartDate, EndDate, Params, NewContextHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CSP_CreateMacContext func(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, Key unsafe.Pointer, NewContextHandle unsafe.Pointer) CSSM_RETURN
var _cSSM_CSP_CreateMacContextErr error

func tryCSSM_CSP_CreateMacContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, Key unsafe.Pointer, NewContextHandle unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CSP_CreateMacContext == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CSP_CreateMacContext", "10.0", _cSSM_CSP_CreateMacContextErr)
	}
	return _cSSM_CSP_CreateMacContext(CSPHandle, AlgorithmID, Key, NewContextHandle), nil
}

// CSSM_CSP_CreateMacContext.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_CreateMacContext
func CSSM_CSP_CreateMacContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, Key unsafe.Pointer, NewContextHandle unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CSP_CreateMacContext(CSPHandle, AlgorithmID, Key, NewContextHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CSP_CreatePassThroughContext func(CSPHandle CSSM_CSP_HANDLE, Key unsafe.Pointer, NewContextHandle unsafe.Pointer) CSSM_RETURN
var _cSSM_CSP_CreatePassThroughContextErr error

func tryCSSM_CSP_CreatePassThroughContext(CSPHandle CSSM_CSP_HANDLE, Key unsafe.Pointer, NewContextHandle unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CSP_CreatePassThroughContext == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CSP_CreatePassThroughContext", "10.0", _cSSM_CSP_CreatePassThroughContextErr)
	}
	return _cSSM_CSP_CreatePassThroughContext(CSPHandle, Key, NewContextHandle), nil
}

// CSSM_CSP_CreatePassThroughContext.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_CreatePassThroughContext
func CSSM_CSP_CreatePassThroughContext(CSPHandle CSSM_CSP_HANDLE, Key unsafe.Pointer, NewContextHandle unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CSP_CreatePassThroughContext(CSPHandle, Key, NewContextHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CSP_CreateRandomGenContext func(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, Seed unsafe.Pointer, Length CSSM_SIZE, NewContextHandle unsafe.Pointer) CSSM_RETURN
var _cSSM_CSP_CreateRandomGenContextErr error

func tryCSSM_CSP_CreateRandomGenContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, Seed unsafe.Pointer, Length CSSM_SIZE, NewContextHandle unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CSP_CreateRandomGenContext == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CSP_CreateRandomGenContext", "10.0", _cSSM_CSP_CreateRandomGenContextErr)
	}
	return _cSSM_CSP_CreateRandomGenContext(CSPHandle, AlgorithmID, Seed, Length, NewContextHandle), nil
}

// CSSM_CSP_CreateRandomGenContext.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_CreateRandomGenContext
func CSSM_CSP_CreateRandomGenContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, Seed unsafe.Pointer, Length CSSM_SIZE, NewContextHandle unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CSP_CreateRandomGenContext(CSPHandle, AlgorithmID, Seed, Length, NewContextHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CSP_CreateSignatureContext func(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, AccessCred unsafe.Pointer, Key unsafe.Pointer, NewContextHandle unsafe.Pointer) CSSM_RETURN
var _cSSM_CSP_CreateSignatureContextErr error

func tryCSSM_CSP_CreateSignatureContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, AccessCred unsafe.Pointer, Key unsafe.Pointer, NewContextHandle unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CSP_CreateSignatureContext == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CSP_CreateSignatureContext", "10.0", _cSSM_CSP_CreateSignatureContextErr)
	}
	return _cSSM_CSP_CreateSignatureContext(CSPHandle, AlgorithmID, AccessCred, Key, NewContextHandle), nil
}

// CSSM_CSP_CreateSignatureContext.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_CreateSignatureContext
func CSSM_CSP_CreateSignatureContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, AccessCred unsafe.Pointer, Key unsafe.Pointer, NewContextHandle unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CSP_CreateSignatureContext(CSPHandle, AlgorithmID, AccessCred, Key, NewContextHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CSP_CreateSymmetricContext func(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, Mode CSSM_ENCRYPT_MODE, AccessCred unsafe.Pointer, Key unsafe.Pointer, InitVector unsafe.Pointer, Padding CSSM_PADDING, Reserved unsafe.Pointer, NewContextHandle unsafe.Pointer) CSSM_RETURN
var _cSSM_CSP_CreateSymmetricContextErr error

func tryCSSM_CSP_CreateSymmetricContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, Mode CSSM_ENCRYPT_MODE, AccessCred unsafe.Pointer, Key unsafe.Pointer, InitVector unsafe.Pointer, Padding CSSM_PADDING, Reserved unsafe.Pointer, NewContextHandle unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CSP_CreateSymmetricContext == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CSP_CreateSymmetricContext", "10.0", _cSSM_CSP_CreateSymmetricContextErr)
	}
	return _cSSM_CSP_CreateSymmetricContext(CSPHandle, AlgorithmID, Mode, AccessCred, Key, InitVector, Padding, Reserved, NewContextHandle), nil
}

// CSSM_CSP_CreateSymmetricContext.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_CreateSymmetricContext
func CSSM_CSP_CreateSymmetricContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, Mode CSSM_ENCRYPT_MODE, AccessCred unsafe.Pointer, Key unsafe.Pointer, InitVector unsafe.Pointer, Padding CSSM_PADDING, Reserved unsafe.Pointer, NewContextHandle unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CSP_CreateSymmetricContext(CSPHandle, AlgorithmID, Mode, AccessCred, Key, InitVector, Padding, Reserved, NewContextHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CSP_GetLoginAcl func(CSPHandle CSSM_CSP_HANDLE, SelectionTag unsafe.Pointer, NumberOfAclInfos *uint32, AclInfos unsafe.Pointer) CSSM_RETURN
var _cSSM_CSP_GetLoginAclErr error

func tryCSSM_CSP_GetLoginAcl(CSPHandle CSSM_CSP_HANDLE, SelectionTag unsafe.Pointer, NumberOfAclInfos *uint32, AclInfos unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CSP_GetLoginAcl == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CSP_GetLoginAcl", "10.0", _cSSM_CSP_GetLoginAclErr)
	}
	return _cSSM_CSP_GetLoginAcl(CSPHandle, SelectionTag, NumberOfAclInfos, AclInfos), nil
}

// CSSM_CSP_GetLoginAcl.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_GetLoginAcl
func CSSM_CSP_GetLoginAcl(CSPHandle CSSM_CSP_HANDLE, SelectionTag unsafe.Pointer, NumberOfAclInfos *uint32, AclInfos unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CSP_GetLoginAcl(CSPHandle, SelectionTag, NumberOfAclInfos, AclInfos)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CSP_GetLoginOwner func(CSPHandle CSSM_CSP_HANDLE, Owner unsafe.Pointer) CSSM_RETURN
var _cSSM_CSP_GetLoginOwnerErr error

func tryCSSM_CSP_GetLoginOwner(CSPHandle CSSM_CSP_HANDLE, Owner unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CSP_GetLoginOwner == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CSP_GetLoginOwner", "10.0", _cSSM_CSP_GetLoginOwnerErr)
	}
	return _cSSM_CSP_GetLoginOwner(CSPHandle, Owner), nil
}

// CSSM_CSP_GetLoginOwner.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_GetLoginOwner
func CSSM_CSP_GetLoginOwner(CSPHandle CSSM_CSP_HANDLE, Owner unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CSP_GetLoginOwner(CSPHandle, Owner)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CSP_GetOperationalStatistics func(CSPHandle CSSM_CSP_HANDLE, Statistics unsafe.Pointer) CSSM_RETURN
var _cSSM_CSP_GetOperationalStatisticsErr error

func tryCSSM_CSP_GetOperationalStatistics(CSPHandle CSSM_CSP_HANDLE, Statistics unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CSP_GetOperationalStatistics == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CSP_GetOperationalStatistics", "10.0", _cSSM_CSP_GetOperationalStatisticsErr)
	}
	return _cSSM_CSP_GetOperationalStatistics(CSPHandle, Statistics), nil
}

// CSSM_CSP_GetOperationalStatistics.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_GetOperationalStatistics
func CSSM_CSP_GetOperationalStatistics(CSPHandle CSSM_CSP_HANDLE, Statistics unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CSP_GetOperationalStatistics(CSPHandle, Statistics)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CSP_Login func(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, LoginName unsafe.Pointer, Reserved unsafe.Pointer) CSSM_RETURN
var _cSSM_CSP_LoginErr error

func tryCSSM_CSP_Login(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, LoginName unsafe.Pointer, Reserved unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CSP_Login == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CSP_Login", "10.0", _cSSM_CSP_LoginErr)
	}
	return _cSSM_CSP_Login(CSPHandle, AccessCred, LoginName, Reserved), nil
}

// CSSM_CSP_Login.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_Login
func CSSM_CSP_Login(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, LoginName unsafe.Pointer, Reserved unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CSP_Login(CSPHandle, AccessCred, LoginName, Reserved)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CSP_Logout func(CSPHandle CSSM_CSP_HANDLE) CSSM_RETURN
var _cSSM_CSP_LogoutErr error

func tryCSSM_CSP_Logout(CSPHandle CSSM_CSP_HANDLE) (CSSM_RETURN, error) {
	if _cSSM_CSP_Logout == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CSP_Logout", "10.0", _cSSM_CSP_LogoutErr)
	}
	return _cSSM_CSP_Logout(CSPHandle), nil
}

// CSSM_CSP_Logout.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_Logout
func CSSM_CSP_Logout(CSPHandle CSSM_CSP_HANDLE) CSSM_RETURN {
	result, callErr := tryCSSM_CSP_Logout(CSPHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CSP_ObtainPrivateKeyFromPublicKey func(CSPHandle CSSM_CSP_HANDLE, PublicKey unsafe.Pointer, PrivateKey unsafe.Pointer) CSSM_RETURN
var _cSSM_CSP_ObtainPrivateKeyFromPublicKeyErr error

func tryCSSM_CSP_ObtainPrivateKeyFromPublicKey(CSPHandle CSSM_CSP_HANDLE, PublicKey unsafe.Pointer, PrivateKey unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CSP_ObtainPrivateKeyFromPublicKey == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CSP_ObtainPrivateKeyFromPublicKey", "10.0", _cSSM_CSP_ObtainPrivateKeyFromPublicKeyErr)
	}
	return _cSSM_CSP_ObtainPrivateKeyFromPublicKey(CSPHandle, PublicKey, PrivateKey), nil
}

// CSSM_CSP_ObtainPrivateKeyFromPublicKey.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_ObtainPrivateKeyFromPublicKey
func CSSM_CSP_ObtainPrivateKeyFromPublicKey(CSPHandle CSSM_CSP_HANDLE, PublicKey unsafe.Pointer, PrivateKey unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CSP_ObtainPrivateKeyFromPublicKey(CSPHandle, PublicKey, PrivateKey)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CSP_PassThrough func(CCHandle CSSM_CC_HANDLE, PassThroughId uint32, InData unsafe.Pointer, OutData unsafe.Pointer) CSSM_RETURN
var _cSSM_CSP_PassThroughErr error

func tryCSSM_CSP_PassThrough(CCHandle CSSM_CC_HANDLE, PassThroughId uint32, InData unsafe.Pointer, OutData unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_CSP_PassThrough == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_CSP_PassThrough", "10.0", _cSSM_CSP_PassThroughErr)
	}
	return _cSSM_CSP_PassThrough(CCHandle, PassThroughId, InData, OutData), nil
}

// CSSM_CSP_PassThrough.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_PassThrough
func CSSM_CSP_PassThrough(CCHandle CSSM_CC_HANDLE, PassThroughId uint32, InData unsafe.Pointer, OutData unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_CSP_PassThrough(CCHandle, PassThroughId, InData, OutData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_ChangeKeyAcl func(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, AclEdit unsafe.Pointer, Key unsafe.Pointer) CSSM_RETURN
var _cSSM_ChangeKeyAclErr error

func tryCSSM_ChangeKeyAcl(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, AclEdit unsafe.Pointer, Key unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_ChangeKeyAcl == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_ChangeKeyAcl", "10.0", _cSSM_ChangeKeyAclErr)
	}
	return _cSSM_ChangeKeyAcl(CSPHandle, AccessCred, AclEdit, Key), nil
}

// CSSM_ChangeKeyAcl.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_ChangeKeyAcl
func CSSM_ChangeKeyAcl(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, AclEdit unsafe.Pointer, Key unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_ChangeKeyAcl(CSPHandle, AccessCred, AclEdit, Key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_ChangeKeyOwner func(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, Key unsafe.Pointer, NewOwner unsafe.Pointer) CSSM_RETURN
var _cSSM_ChangeKeyOwnerErr error

func tryCSSM_ChangeKeyOwner(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, Key unsafe.Pointer, NewOwner unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_ChangeKeyOwner == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_ChangeKeyOwner", "10.0", _cSSM_ChangeKeyOwnerErr)
	}
	return _cSSM_ChangeKeyOwner(CSPHandle, AccessCred, Key, NewOwner), nil
}

// CSSM_ChangeKeyOwner.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_ChangeKeyOwner
func CSSM_ChangeKeyOwner(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, Key unsafe.Pointer, NewOwner unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_ChangeKeyOwner(CSPHandle, AccessCred, Key, NewOwner)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_Authenticate func(DLDBHandle unsafe.Pointer, AccessRequest CSSM_DB_ACCESS_TYPE, AccessCred unsafe.Pointer) CSSM_RETURN
var _cSSM_DL_AuthenticateErr error

func tryCSSM_DL_Authenticate(DLDBHandle unsafe.Pointer, AccessRequest CSSM_DB_ACCESS_TYPE, AccessCred unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_DL_Authenticate == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_DL_Authenticate", "10.0", _cSSM_DL_AuthenticateErr)
	}
	return _cSSM_DL_Authenticate(DLDBHandle, AccessRequest, AccessCred), nil
}

// CSSM_DL_Authenticate.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_Authenticate
func CSSM_DL_Authenticate(DLDBHandle unsafe.Pointer, AccessRequest CSSM_DB_ACCESS_TYPE, AccessCred unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_DL_Authenticate(DLDBHandle, AccessRequest, AccessCred)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_ChangeDbAcl func(DLDBHandle unsafe.Pointer, AccessCred unsafe.Pointer, AclEdit unsafe.Pointer) CSSM_RETURN
var _cSSM_DL_ChangeDbAclErr error

func tryCSSM_DL_ChangeDbAcl(DLDBHandle unsafe.Pointer, AccessCred unsafe.Pointer, AclEdit unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_DL_ChangeDbAcl == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_DL_ChangeDbAcl", "10.0", _cSSM_DL_ChangeDbAclErr)
	}
	return _cSSM_DL_ChangeDbAcl(DLDBHandle, AccessCred, AclEdit), nil
}

// CSSM_DL_ChangeDbAcl.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_ChangeDbAcl
func CSSM_DL_ChangeDbAcl(DLDBHandle unsafe.Pointer, AccessCred unsafe.Pointer, AclEdit unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_DL_ChangeDbAcl(DLDBHandle, AccessCred, AclEdit)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_ChangeDbOwner func(DLDBHandle unsafe.Pointer, AccessCred unsafe.Pointer, NewOwner unsafe.Pointer) CSSM_RETURN
var _cSSM_DL_ChangeDbOwnerErr error

func tryCSSM_DL_ChangeDbOwner(DLDBHandle unsafe.Pointer, AccessCred unsafe.Pointer, NewOwner unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_DL_ChangeDbOwner == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_DL_ChangeDbOwner", "10.0", _cSSM_DL_ChangeDbOwnerErr)
	}
	return _cSSM_DL_ChangeDbOwner(DLDBHandle, AccessCred, NewOwner), nil
}

// CSSM_DL_ChangeDbOwner.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_ChangeDbOwner
func CSSM_DL_ChangeDbOwner(DLDBHandle unsafe.Pointer, AccessCred unsafe.Pointer, NewOwner unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_DL_ChangeDbOwner(DLDBHandle, AccessCred, NewOwner)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_CreateRelation func(DLDBHandle unsafe.Pointer, RelationID CSSM_DB_RECORDTYPE, RelationName string, NumberOfAttributes uint32, pAttributeInfo unsafe.Pointer, NumberOfIndexes uint32, pIndexInfo unsafe.Pointer) CSSM_RETURN
var _cSSM_DL_CreateRelationErr error

func tryCSSM_DL_CreateRelation(DLDBHandle unsafe.Pointer, RelationID CSSM_DB_RECORDTYPE, RelationName string, NumberOfAttributes uint32, pAttributeInfo unsafe.Pointer, NumberOfIndexes uint32, pIndexInfo unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_DL_CreateRelation == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_DL_CreateRelation", "10.0", _cSSM_DL_CreateRelationErr)
	}
	return _cSSM_DL_CreateRelation(DLDBHandle, RelationID, RelationName, NumberOfAttributes, pAttributeInfo, NumberOfIndexes, pIndexInfo), nil
}

// CSSM_DL_CreateRelation.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_CreateRelation
func CSSM_DL_CreateRelation(DLDBHandle unsafe.Pointer, RelationID CSSM_DB_RECORDTYPE, RelationName string, NumberOfAttributes uint32, pAttributeInfo unsafe.Pointer, NumberOfIndexes uint32, pIndexInfo unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_DL_CreateRelation(DLDBHandle, RelationID, RelationName, NumberOfAttributes, pAttributeInfo, NumberOfIndexes, pIndexInfo)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_DataAbortQuery func(DLDBHandle unsafe.Pointer, ResultsHandle CSSM_HANDLE) CSSM_RETURN
var _cSSM_DL_DataAbortQueryErr error

func tryCSSM_DL_DataAbortQuery(DLDBHandle unsafe.Pointer, ResultsHandle CSSM_HANDLE) (CSSM_RETURN, error) {
	if _cSSM_DL_DataAbortQuery == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_DL_DataAbortQuery", "10.0", _cSSM_DL_DataAbortQueryErr)
	}
	return _cSSM_DL_DataAbortQuery(DLDBHandle, ResultsHandle), nil
}

// CSSM_DL_DataAbortQuery.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_DataAbortQuery
func CSSM_DL_DataAbortQuery(DLDBHandle unsafe.Pointer, ResultsHandle CSSM_HANDLE) CSSM_RETURN {
	result, callErr := tryCSSM_DL_DataAbortQuery(DLDBHandle, ResultsHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_DataDelete func(DLDBHandle unsafe.Pointer, UniqueRecordIdentifier unsafe.Pointer) CSSM_RETURN
var _cSSM_DL_DataDeleteErr error

func tryCSSM_DL_DataDelete(DLDBHandle unsafe.Pointer, UniqueRecordIdentifier unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_DL_DataDelete == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_DL_DataDelete", "10.0", _cSSM_DL_DataDeleteErr)
	}
	return _cSSM_DL_DataDelete(DLDBHandle, UniqueRecordIdentifier), nil
}

// CSSM_DL_DataDelete.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_DataDelete
func CSSM_DL_DataDelete(DLDBHandle unsafe.Pointer, UniqueRecordIdentifier unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_DL_DataDelete(DLDBHandle, UniqueRecordIdentifier)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_DataGetFirst func(DLDBHandle unsafe.Pointer, Query unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, Attributes unsafe.Pointer, Data unsafe.Pointer, UniqueId unsafe.Pointer) CSSM_RETURN
var _cSSM_DL_DataGetFirstErr error

func tryCSSM_DL_DataGetFirst(DLDBHandle unsafe.Pointer, Query unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, Attributes unsafe.Pointer, Data unsafe.Pointer, UniqueId unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_DL_DataGetFirst == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_DL_DataGetFirst", "10.0", _cSSM_DL_DataGetFirstErr)
	}
	return _cSSM_DL_DataGetFirst(DLDBHandle, Query, ResultsHandle, Attributes, Data, UniqueId), nil
}

// CSSM_DL_DataGetFirst.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_DataGetFirst
func CSSM_DL_DataGetFirst(DLDBHandle unsafe.Pointer, Query unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, Attributes unsafe.Pointer, Data unsafe.Pointer, UniqueId unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_DL_DataGetFirst(DLDBHandle, Query, ResultsHandle, Attributes, Data, UniqueId)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_DataGetFromUniqueRecordId func(DLDBHandle unsafe.Pointer, UniqueRecord unsafe.Pointer, Attributes unsafe.Pointer, Data unsafe.Pointer) CSSM_RETURN
var _cSSM_DL_DataGetFromUniqueRecordIdErr error

func tryCSSM_DL_DataGetFromUniqueRecordId(DLDBHandle unsafe.Pointer, UniqueRecord unsafe.Pointer, Attributes unsafe.Pointer, Data unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_DL_DataGetFromUniqueRecordId == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_DL_DataGetFromUniqueRecordId", "10.0", _cSSM_DL_DataGetFromUniqueRecordIdErr)
	}
	return _cSSM_DL_DataGetFromUniqueRecordId(DLDBHandle, UniqueRecord, Attributes, Data), nil
}

// CSSM_DL_DataGetFromUniqueRecordId.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_DataGetFromUniqueRecordId
func CSSM_DL_DataGetFromUniqueRecordId(DLDBHandle unsafe.Pointer, UniqueRecord unsafe.Pointer, Attributes unsafe.Pointer, Data unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_DL_DataGetFromUniqueRecordId(DLDBHandle, UniqueRecord, Attributes, Data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_DataGetNext func(DLDBHandle unsafe.Pointer, ResultsHandle CSSM_HANDLE, Attributes unsafe.Pointer, Data unsafe.Pointer, UniqueId unsafe.Pointer) CSSM_RETURN
var _cSSM_DL_DataGetNextErr error

func tryCSSM_DL_DataGetNext(DLDBHandle unsafe.Pointer, ResultsHandle CSSM_HANDLE, Attributes unsafe.Pointer, Data unsafe.Pointer, UniqueId unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_DL_DataGetNext == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_DL_DataGetNext", "10.0", _cSSM_DL_DataGetNextErr)
	}
	return _cSSM_DL_DataGetNext(DLDBHandle, ResultsHandle, Attributes, Data, UniqueId), nil
}

// CSSM_DL_DataGetNext.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_DataGetNext
func CSSM_DL_DataGetNext(DLDBHandle unsafe.Pointer, ResultsHandle CSSM_HANDLE, Attributes unsafe.Pointer, Data unsafe.Pointer, UniqueId unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_DL_DataGetNext(DLDBHandle, ResultsHandle, Attributes, Data, UniqueId)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_DataInsert func(DLDBHandle unsafe.Pointer, RecordType CSSM_DB_RECORDTYPE, Attributes unsafe.Pointer, Data unsafe.Pointer, UniqueId unsafe.Pointer) CSSM_RETURN
var _cSSM_DL_DataInsertErr error

func tryCSSM_DL_DataInsert(DLDBHandle unsafe.Pointer, RecordType CSSM_DB_RECORDTYPE, Attributes unsafe.Pointer, Data unsafe.Pointer, UniqueId unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_DL_DataInsert == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_DL_DataInsert", "10.0", _cSSM_DL_DataInsertErr)
	}
	return _cSSM_DL_DataInsert(DLDBHandle, RecordType, Attributes, Data, UniqueId), nil
}

// CSSM_DL_DataInsert.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_DataInsert
func CSSM_DL_DataInsert(DLDBHandle unsafe.Pointer, RecordType CSSM_DB_RECORDTYPE, Attributes unsafe.Pointer, Data unsafe.Pointer, UniqueId unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_DL_DataInsert(DLDBHandle, RecordType, Attributes, Data, UniqueId)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_DataModify func(DLDBHandle unsafe.Pointer, RecordType CSSM_DB_RECORDTYPE, UniqueRecordIdentifier unsafe.Pointer, AttributesToBeModified unsafe.Pointer, DataToBeModified unsafe.Pointer, ModifyMode CSSM_DB_MODIFY_MODE) CSSM_RETURN
var _cSSM_DL_DataModifyErr error

func tryCSSM_DL_DataModify(DLDBHandle unsafe.Pointer, RecordType CSSM_DB_RECORDTYPE, UniqueRecordIdentifier unsafe.Pointer, AttributesToBeModified unsafe.Pointer, DataToBeModified unsafe.Pointer, ModifyMode CSSM_DB_MODIFY_MODE) (CSSM_RETURN, error) {
	if _cSSM_DL_DataModify == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_DL_DataModify", "10.0", _cSSM_DL_DataModifyErr)
	}
	return _cSSM_DL_DataModify(DLDBHandle, RecordType, UniqueRecordIdentifier, AttributesToBeModified, DataToBeModified, ModifyMode), nil
}

// CSSM_DL_DataModify.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_DataModify
func CSSM_DL_DataModify(DLDBHandle unsafe.Pointer, RecordType CSSM_DB_RECORDTYPE, UniqueRecordIdentifier unsafe.Pointer, AttributesToBeModified unsafe.Pointer, DataToBeModified unsafe.Pointer, ModifyMode CSSM_DB_MODIFY_MODE) CSSM_RETURN {
	result, callErr := tryCSSM_DL_DataModify(DLDBHandle, RecordType, UniqueRecordIdentifier, AttributesToBeModified, DataToBeModified, ModifyMode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_DbClose func(DLDBHandle unsafe.Pointer) CSSM_RETURN
var _cSSM_DL_DbCloseErr error

func tryCSSM_DL_DbClose(DLDBHandle unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_DL_DbClose == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_DL_DbClose", "10.0", _cSSM_DL_DbCloseErr)
	}
	return _cSSM_DL_DbClose(DLDBHandle), nil
}

// CSSM_DL_DbClose.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_DbClose
func CSSM_DL_DbClose(DLDBHandle unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_DL_DbClose(DLDBHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_DbCreate func(DLHandle CSSM_DL_HANDLE, DbName string, DbLocation unsafe.Pointer, DBInfo unsafe.Pointer, AccessRequest CSSM_DB_ACCESS_TYPE, CredAndAclEntry unsafe.Pointer, OpenParameters unsafe.Pointer, DbHandle unsafe.Pointer) CSSM_RETURN
var _cSSM_DL_DbCreateErr error

func tryCSSM_DL_DbCreate(DLHandle CSSM_DL_HANDLE, DbName string, DbLocation unsafe.Pointer, DBInfo unsafe.Pointer, AccessRequest CSSM_DB_ACCESS_TYPE, CredAndAclEntry unsafe.Pointer, OpenParameters unsafe.Pointer, DbHandle unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_DL_DbCreate == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_DL_DbCreate", "10.0", _cSSM_DL_DbCreateErr)
	}
	return _cSSM_DL_DbCreate(DLHandle, DbName, DbLocation, DBInfo, AccessRequest, CredAndAclEntry, OpenParameters, DbHandle), nil
}

// CSSM_DL_DbCreate.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_DbCreate
func CSSM_DL_DbCreate(DLHandle CSSM_DL_HANDLE, DbName string, DbLocation unsafe.Pointer, DBInfo unsafe.Pointer, AccessRequest CSSM_DB_ACCESS_TYPE, CredAndAclEntry unsafe.Pointer, OpenParameters unsafe.Pointer, DbHandle unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_DL_DbCreate(DLHandle, DbName, DbLocation, DBInfo, AccessRequest, CredAndAclEntry, OpenParameters, DbHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_DbDelete func(DLHandle CSSM_DL_HANDLE, DbName string, DbLocation unsafe.Pointer, AccessCred unsafe.Pointer) CSSM_RETURN
var _cSSM_DL_DbDeleteErr error

func tryCSSM_DL_DbDelete(DLHandle CSSM_DL_HANDLE, DbName string, DbLocation unsafe.Pointer, AccessCred unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_DL_DbDelete == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_DL_DbDelete", "10.0", _cSSM_DL_DbDeleteErr)
	}
	return _cSSM_DL_DbDelete(DLHandle, DbName, DbLocation, AccessCred), nil
}

// CSSM_DL_DbDelete.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_DbDelete
func CSSM_DL_DbDelete(DLHandle CSSM_DL_HANDLE, DbName string, DbLocation unsafe.Pointer, AccessCred unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_DL_DbDelete(DLHandle, DbName, DbLocation, AccessCred)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_DbOpen func(DLHandle CSSM_DL_HANDLE, DbName string, DbLocation unsafe.Pointer, AccessRequest CSSM_DB_ACCESS_TYPE, AccessCred unsafe.Pointer, OpenParameters unsafe.Pointer, DbHandle unsafe.Pointer) CSSM_RETURN
var _cSSM_DL_DbOpenErr error

func tryCSSM_DL_DbOpen(DLHandle CSSM_DL_HANDLE, DbName string, DbLocation unsafe.Pointer, AccessRequest CSSM_DB_ACCESS_TYPE, AccessCred unsafe.Pointer, OpenParameters unsafe.Pointer, DbHandle unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_DL_DbOpen == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_DL_DbOpen", "10.0", _cSSM_DL_DbOpenErr)
	}
	return _cSSM_DL_DbOpen(DLHandle, DbName, DbLocation, AccessRequest, AccessCred, OpenParameters, DbHandle), nil
}

// CSSM_DL_DbOpen.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_DbOpen
func CSSM_DL_DbOpen(DLHandle CSSM_DL_HANDLE, DbName string, DbLocation unsafe.Pointer, AccessRequest CSSM_DB_ACCESS_TYPE, AccessCred unsafe.Pointer, OpenParameters unsafe.Pointer, DbHandle unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_DL_DbOpen(DLHandle, DbName, DbLocation, AccessRequest, AccessCred, OpenParameters, DbHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_DestroyRelation func(DLDBHandle unsafe.Pointer, RelationID CSSM_DB_RECORDTYPE) CSSM_RETURN
var _cSSM_DL_DestroyRelationErr error

func tryCSSM_DL_DestroyRelation(DLDBHandle unsafe.Pointer, RelationID CSSM_DB_RECORDTYPE) (CSSM_RETURN, error) {
	if _cSSM_DL_DestroyRelation == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_DL_DestroyRelation", "10.0", _cSSM_DL_DestroyRelationErr)
	}
	return _cSSM_DL_DestroyRelation(DLDBHandle, RelationID), nil
}

// CSSM_DL_DestroyRelation.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_DestroyRelation
func CSSM_DL_DestroyRelation(DLDBHandle unsafe.Pointer, RelationID CSSM_DB_RECORDTYPE) CSSM_RETURN {
	result, callErr := tryCSSM_DL_DestroyRelation(DLDBHandle, RelationID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_FreeNameList func(DLHandle CSSM_DL_HANDLE, NameList unsafe.Pointer) CSSM_RETURN
var _cSSM_DL_FreeNameListErr error

func tryCSSM_DL_FreeNameList(DLHandle CSSM_DL_HANDLE, NameList unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_DL_FreeNameList == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_DL_FreeNameList", "10.0", _cSSM_DL_FreeNameListErr)
	}
	return _cSSM_DL_FreeNameList(DLHandle, NameList), nil
}

// CSSM_DL_FreeNameList.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_FreeNameList
func CSSM_DL_FreeNameList(DLHandle CSSM_DL_HANDLE, NameList unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_DL_FreeNameList(DLHandle, NameList)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_FreeUniqueRecord func(DLDBHandle unsafe.Pointer, UniqueRecord unsafe.Pointer) CSSM_RETURN
var _cSSM_DL_FreeUniqueRecordErr error

func tryCSSM_DL_FreeUniqueRecord(DLDBHandle unsafe.Pointer, UniqueRecord unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_DL_FreeUniqueRecord == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_DL_FreeUniqueRecord", "10.0", _cSSM_DL_FreeUniqueRecordErr)
	}
	return _cSSM_DL_FreeUniqueRecord(DLDBHandle, UniqueRecord), nil
}

// CSSM_DL_FreeUniqueRecord.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_FreeUniqueRecord
func CSSM_DL_FreeUniqueRecord(DLDBHandle unsafe.Pointer, UniqueRecord unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_DL_FreeUniqueRecord(DLDBHandle, UniqueRecord)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_GetDbAcl func(DLDBHandle unsafe.Pointer, SelectionTag unsafe.Pointer, NumberOfAclInfos *uint32, AclInfos unsafe.Pointer) CSSM_RETURN
var _cSSM_DL_GetDbAclErr error

func tryCSSM_DL_GetDbAcl(DLDBHandle unsafe.Pointer, SelectionTag unsafe.Pointer, NumberOfAclInfos *uint32, AclInfos unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_DL_GetDbAcl == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_DL_GetDbAcl", "10.0", _cSSM_DL_GetDbAclErr)
	}
	return _cSSM_DL_GetDbAcl(DLDBHandle, SelectionTag, NumberOfAclInfos, AclInfos), nil
}

// CSSM_DL_GetDbAcl.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_GetDbAcl
func CSSM_DL_GetDbAcl(DLDBHandle unsafe.Pointer, SelectionTag unsafe.Pointer, NumberOfAclInfos *uint32, AclInfos unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_DL_GetDbAcl(DLDBHandle, SelectionTag, NumberOfAclInfos, AclInfos)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_GetDbNameFromHandle func(DLDBHandle unsafe.Pointer, DbName string) CSSM_RETURN
var _cSSM_DL_GetDbNameFromHandleErr error

func tryCSSM_DL_GetDbNameFromHandle(DLDBHandle unsafe.Pointer, DbName string) (CSSM_RETURN, error) {
	if _cSSM_DL_GetDbNameFromHandle == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_DL_GetDbNameFromHandle", "10.0", _cSSM_DL_GetDbNameFromHandleErr)
	}
	return _cSSM_DL_GetDbNameFromHandle(DLDBHandle, DbName), nil
}

// CSSM_DL_GetDbNameFromHandle.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_GetDbNameFromHandle
func CSSM_DL_GetDbNameFromHandle(DLDBHandle unsafe.Pointer, DbName string) CSSM_RETURN {
	result, callErr := tryCSSM_DL_GetDbNameFromHandle(DLDBHandle, DbName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_GetDbNames func(DLHandle CSSM_DL_HANDLE, NameList unsafe.Pointer) CSSM_RETURN
var _cSSM_DL_GetDbNamesErr error

func tryCSSM_DL_GetDbNames(DLHandle CSSM_DL_HANDLE, NameList unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_DL_GetDbNames == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_DL_GetDbNames", "10.0", _cSSM_DL_GetDbNamesErr)
	}
	return _cSSM_DL_GetDbNames(DLHandle, NameList), nil
}

// CSSM_DL_GetDbNames.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_GetDbNames
func CSSM_DL_GetDbNames(DLHandle CSSM_DL_HANDLE, NameList unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_DL_GetDbNames(DLHandle, NameList)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_GetDbOwner func(DLDBHandle unsafe.Pointer, Owner unsafe.Pointer) CSSM_RETURN
var _cSSM_DL_GetDbOwnerErr error

func tryCSSM_DL_GetDbOwner(DLDBHandle unsafe.Pointer, Owner unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_DL_GetDbOwner == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_DL_GetDbOwner", "10.0", _cSSM_DL_GetDbOwnerErr)
	}
	return _cSSM_DL_GetDbOwner(DLDBHandle, Owner), nil
}

// CSSM_DL_GetDbOwner.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_GetDbOwner
func CSSM_DL_GetDbOwner(DLDBHandle unsafe.Pointer, Owner unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_DL_GetDbOwner(DLDBHandle, Owner)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_PassThrough func(DLDBHandle unsafe.Pointer, PassThroughId uint32, InputParams unsafe.Pointer, OutputParams unsafe.Pointer) CSSM_RETURN
var _cSSM_DL_PassThroughErr error

func tryCSSM_DL_PassThrough(DLDBHandle unsafe.Pointer, PassThroughId uint32, InputParams unsafe.Pointer, OutputParams unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_DL_PassThrough == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_DL_PassThrough", "10.0", _cSSM_DL_PassThroughErr)
	}
	return _cSSM_DL_PassThrough(DLDBHandle, PassThroughId, InputParams, OutputParams), nil
}

// CSSM_DL_PassThrough.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_PassThrough
func CSSM_DL_PassThrough(DLDBHandle unsafe.Pointer, PassThroughId uint32, InputParams unsafe.Pointer, OutputParams unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_DL_PassThrough(DLDBHandle, PassThroughId, InputParams, OutputParams)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DecryptData func(CCHandle CSSM_CC_HANDLE, CipherBufs unsafe.Pointer, CipherBufCount uint32, ClearBufs unsafe.Pointer, ClearBufCount uint32, bytesDecrypted unsafe.Pointer, RemData unsafe.Pointer) CSSM_RETURN
var _cSSM_DecryptDataErr error

func tryCSSM_DecryptData(CCHandle CSSM_CC_HANDLE, CipherBufs unsafe.Pointer, CipherBufCount uint32, ClearBufs unsafe.Pointer, ClearBufCount uint32, bytesDecrypted unsafe.Pointer, RemData unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_DecryptData == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_DecryptData", "10.0", _cSSM_DecryptDataErr)
	}
	return _cSSM_DecryptData(CCHandle, CipherBufs, CipherBufCount, ClearBufs, ClearBufCount, bytesDecrypted, RemData), nil
}

// CSSM_DecryptData.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DecryptData
func CSSM_DecryptData(CCHandle CSSM_CC_HANDLE, CipherBufs unsafe.Pointer, CipherBufCount uint32, ClearBufs unsafe.Pointer, ClearBufCount uint32, bytesDecrypted unsafe.Pointer, RemData unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_DecryptData(CCHandle, CipherBufs, CipherBufCount, ClearBufs, ClearBufCount, bytesDecrypted, RemData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DecryptDataFinal func(CCHandle CSSM_CC_HANDLE, RemData unsafe.Pointer) CSSM_RETURN
var _cSSM_DecryptDataFinalErr error

func tryCSSM_DecryptDataFinal(CCHandle CSSM_CC_HANDLE, RemData unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_DecryptDataFinal == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_DecryptDataFinal", "10.0", _cSSM_DecryptDataFinalErr)
	}
	return _cSSM_DecryptDataFinal(CCHandle, RemData), nil
}

// CSSM_DecryptDataFinal.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DecryptDataFinal
func CSSM_DecryptDataFinal(CCHandle CSSM_CC_HANDLE, RemData unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_DecryptDataFinal(CCHandle, RemData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DecryptDataInit func(CCHandle CSSM_CC_HANDLE) CSSM_RETURN
var _cSSM_DecryptDataInitErr error

func tryCSSM_DecryptDataInit(CCHandle CSSM_CC_HANDLE) (CSSM_RETURN, error) {
	if _cSSM_DecryptDataInit == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_DecryptDataInit", "10.0", _cSSM_DecryptDataInitErr)
	}
	return _cSSM_DecryptDataInit(CCHandle), nil
}

// CSSM_DecryptDataInit.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DecryptDataInit
func CSSM_DecryptDataInit(CCHandle CSSM_CC_HANDLE) CSSM_RETURN {
	result, callErr := tryCSSM_DecryptDataInit(CCHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DecryptDataInitP func(CCHandle CSSM_CC_HANDLE, Privilege CSSM_PRIVILEGE) CSSM_RETURN
var _cSSM_DecryptDataInitPErr error

func tryCSSM_DecryptDataInitP(CCHandle CSSM_CC_HANDLE, Privilege CSSM_PRIVILEGE) (CSSM_RETURN, error) {
	if _cSSM_DecryptDataInitP == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_DecryptDataInitP", "10.0", _cSSM_DecryptDataInitPErr)
	}
	return _cSSM_DecryptDataInitP(CCHandle, Privilege), nil
}

// CSSM_DecryptDataInitP.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DecryptDataInitP
func CSSM_DecryptDataInitP(CCHandle CSSM_CC_HANDLE, Privilege CSSM_PRIVILEGE) CSSM_RETURN {
	result, callErr := tryCSSM_DecryptDataInitP(CCHandle, Privilege)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DecryptDataP func(CCHandle CSSM_CC_HANDLE, CipherBufs unsafe.Pointer, CipherBufCount uint32, ClearBufs unsafe.Pointer, ClearBufCount uint32, bytesDecrypted unsafe.Pointer, RemData unsafe.Pointer, Privilege CSSM_PRIVILEGE) CSSM_RETURN
var _cSSM_DecryptDataPErr error

func tryCSSM_DecryptDataP(CCHandle CSSM_CC_HANDLE, CipherBufs unsafe.Pointer, CipherBufCount uint32, ClearBufs unsafe.Pointer, ClearBufCount uint32, bytesDecrypted unsafe.Pointer, RemData unsafe.Pointer, Privilege CSSM_PRIVILEGE) (CSSM_RETURN, error) {
	if _cSSM_DecryptDataP == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_DecryptDataP", "10.0", _cSSM_DecryptDataPErr)
	}
	return _cSSM_DecryptDataP(CCHandle, CipherBufs, CipherBufCount, ClearBufs, ClearBufCount, bytesDecrypted, RemData, Privilege), nil
}

// CSSM_DecryptDataP.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DecryptDataP
func CSSM_DecryptDataP(CCHandle CSSM_CC_HANDLE, CipherBufs unsafe.Pointer, CipherBufCount uint32, ClearBufs unsafe.Pointer, ClearBufCount uint32, bytesDecrypted unsafe.Pointer, RemData unsafe.Pointer, Privilege CSSM_PRIVILEGE) CSSM_RETURN {
	result, callErr := tryCSSM_DecryptDataP(CCHandle, CipherBufs, CipherBufCount, ClearBufs, ClearBufCount, bytesDecrypted, RemData, Privilege)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DecryptDataUpdate func(CCHandle CSSM_CC_HANDLE, CipherBufs unsafe.Pointer, CipherBufCount uint32, ClearBufs unsafe.Pointer, ClearBufCount uint32, bytesDecrypted unsafe.Pointer) CSSM_RETURN
var _cSSM_DecryptDataUpdateErr error

func tryCSSM_DecryptDataUpdate(CCHandle CSSM_CC_HANDLE, CipherBufs unsafe.Pointer, CipherBufCount uint32, ClearBufs unsafe.Pointer, ClearBufCount uint32, bytesDecrypted unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_DecryptDataUpdate == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_DecryptDataUpdate", "10.0", _cSSM_DecryptDataUpdateErr)
	}
	return _cSSM_DecryptDataUpdate(CCHandle, CipherBufs, CipherBufCount, ClearBufs, ClearBufCount, bytesDecrypted), nil
}

// CSSM_DecryptDataUpdate.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DecryptDataUpdate
func CSSM_DecryptDataUpdate(CCHandle CSSM_CC_HANDLE, CipherBufs unsafe.Pointer, CipherBufCount uint32, ClearBufs unsafe.Pointer, ClearBufCount uint32, bytesDecrypted unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_DecryptDataUpdate(CCHandle, CipherBufs, CipherBufCount, ClearBufs, ClearBufCount, bytesDecrypted)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DeleteContext func(CCHandle CSSM_CC_HANDLE) CSSM_RETURN
var _cSSM_DeleteContextErr error

func tryCSSM_DeleteContext(CCHandle CSSM_CC_HANDLE) (CSSM_RETURN, error) {
	if _cSSM_DeleteContext == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_DeleteContext", "10.0", _cSSM_DeleteContextErr)
	}
	return _cSSM_DeleteContext(CCHandle), nil
}

// CSSM_DeleteContext.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DeleteContext
func CSSM_DeleteContext(CCHandle CSSM_CC_HANDLE) CSSM_RETURN {
	result, callErr := tryCSSM_DeleteContext(CCHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DeleteContextAttributes func(CCHandle CSSM_CC_HANDLE, NumberOfAttributes uint32, ContextAttributes unsafe.Pointer) CSSM_RETURN
var _cSSM_DeleteContextAttributesErr error

func tryCSSM_DeleteContextAttributes(CCHandle CSSM_CC_HANDLE, NumberOfAttributes uint32, ContextAttributes unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_DeleteContextAttributes == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_DeleteContextAttributes", "10.0", _cSSM_DeleteContextAttributesErr)
	}
	return _cSSM_DeleteContextAttributes(CCHandle, NumberOfAttributes, ContextAttributes), nil
}

// CSSM_DeleteContextAttributes.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DeleteContextAttributes
func CSSM_DeleteContextAttributes(CCHandle CSSM_CC_HANDLE, NumberOfAttributes uint32, ContextAttributes unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_DeleteContextAttributes(CCHandle, NumberOfAttributes, ContextAttributes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DeriveKey func(CCHandle CSSM_CC_HANDLE, Param unsafe.Pointer, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, DerivedKey unsafe.Pointer) CSSM_RETURN
var _cSSM_DeriveKeyErr error

func tryCSSM_DeriveKey(CCHandle CSSM_CC_HANDLE, Param unsafe.Pointer, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, DerivedKey unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_DeriveKey == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_DeriveKey", "10.0", _cSSM_DeriveKeyErr)
	}
	return _cSSM_DeriveKey(CCHandle, Param, KeyUsage, KeyAttr, KeyLabel, CredAndAclEntry, DerivedKey), nil
}

// CSSM_DeriveKey.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DeriveKey
func CSSM_DeriveKey(CCHandle CSSM_CC_HANDLE, Param unsafe.Pointer, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, DerivedKey unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_DeriveKey(CCHandle, Param, KeyUsage, KeyAttr, KeyLabel, CredAndAclEntry, DerivedKey)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DigestData func(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, Digest unsafe.Pointer) CSSM_RETURN
var _cSSM_DigestDataErr error

func tryCSSM_DigestData(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, Digest unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_DigestData == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_DigestData", "10.0", _cSSM_DigestDataErr)
	}
	return _cSSM_DigestData(CCHandle, DataBufs, DataBufCount, Digest), nil
}

// CSSM_DigestData.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DigestData
func CSSM_DigestData(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, Digest unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_DigestData(CCHandle, DataBufs, DataBufCount, Digest)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DigestDataClone func(CCHandle CSSM_CC_HANDLE, ClonednewCCHandle unsafe.Pointer) CSSM_RETURN
var _cSSM_DigestDataCloneErr error

func tryCSSM_DigestDataClone(CCHandle CSSM_CC_HANDLE, ClonednewCCHandle unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_DigestDataClone == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_DigestDataClone", "10.0", _cSSM_DigestDataCloneErr)
	}
	return _cSSM_DigestDataClone(CCHandle, ClonednewCCHandle), nil
}

// CSSM_DigestDataClone.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DigestDataClone
func CSSM_DigestDataClone(CCHandle CSSM_CC_HANDLE, ClonednewCCHandle unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_DigestDataClone(CCHandle, ClonednewCCHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DigestDataFinal func(CCHandle CSSM_CC_HANDLE, Digest unsafe.Pointer) CSSM_RETURN
var _cSSM_DigestDataFinalErr error

func tryCSSM_DigestDataFinal(CCHandle CSSM_CC_HANDLE, Digest unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_DigestDataFinal == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_DigestDataFinal", "10.0", _cSSM_DigestDataFinalErr)
	}
	return _cSSM_DigestDataFinal(CCHandle, Digest), nil
}

// CSSM_DigestDataFinal.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DigestDataFinal
func CSSM_DigestDataFinal(CCHandle CSSM_CC_HANDLE, Digest unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_DigestDataFinal(CCHandle, Digest)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DigestDataInit func(CCHandle CSSM_CC_HANDLE) CSSM_RETURN
var _cSSM_DigestDataInitErr error

func tryCSSM_DigestDataInit(CCHandle CSSM_CC_HANDLE) (CSSM_RETURN, error) {
	if _cSSM_DigestDataInit == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_DigestDataInit", "10.0", _cSSM_DigestDataInitErr)
	}
	return _cSSM_DigestDataInit(CCHandle), nil
}

// CSSM_DigestDataInit.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DigestDataInit
func CSSM_DigestDataInit(CCHandle CSSM_CC_HANDLE) CSSM_RETURN {
	result, callErr := tryCSSM_DigestDataInit(CCHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DigestDataUpdate func(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) CSSM_RETURN
var _cSSM_DigestDataUpdateErr error

func tryCSSM_DigestDataUpdate(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) (CSSM_RETURN, error) {
	if _cSSM_DigestDataUpdate == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_DigestDataUpdate", "10.0", _cSSM_DigestDataUpdateErr)
	}
	return _cSSM_DigestDataUpdate(CCHandle, DataBufs, DataBufCount), nil
}

// CSSM_DigestDataUpdate.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DigestDataUpdate
func CSSM_DigestDataUpdate(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) CSSM_RETURN {
	result, callErr := tryCSSM_DigestDataUpdate(CCHandle, DataBufs, DataBufCount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_EncryptData func(CCHandle CSSM_CC_HANDLE, ClearBufs unsafe.Pointer, ClearBufCount uint32, CipherBufs unsafe.Pointer, CipherBufCount uint32, bytesEncrypted unsafe.Pointer, RemData unsafe.Pointer) CSSM_RETURN
var _cSSM_EncryptDataErr error

func tryCSSM_EncryptData(CCHandle CSSM_CC_HANDLE, ClearBufs unsafe.Pointer, ClearBufCount uint32, CipherBufs unsafe.Pointer, CipherBufCount uint32, bytesEncrypted unsafe.Pointer, RemData unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_EncryptData == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_EncryptData", "10.0", _cSSM_EncryptDataErr)
	}
	return _cSSM_EncryptData(CCHandle, ClearBufs, ClearBufCount, CipherBufs, CipherBufCount, bytesEncrypted, RemData), nil
}

// CSSM_EncryptData.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_EncryptData
func CSSM_EncryptData(CCHandle CSSM_CC_HANDLE, ClearBufs unsafe.Pointer, ClearBufCount uint32, CipherBufs unsafe.Pointer, CipherBufCount uint32, bytesEncrypted unsafe.Pointer, RemData unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_EncryptData(CCHandle, ClearBufs, ClearBufCount, CipherBufs, CipherBufCount, bytesEncrypted, RemData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_EncryptDataFinal func(CCHandle CSSM_CC_HANDLE, RemData unsafe.Pointer) CSSM_RETURN
var _cSSM_EncryptDataFinalErr error

func tryCSSM_EncryptDataFinal(CCHandle CSSM_CC_HANDLE, RemData unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_EncryptDataFinal == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_EncryptDataFinal", "10.0", _cSSM_EncryptDataFinalErr)
	}
	return _cSSM_EncryptDataFinal(CCHandle, RemData), nil
}

// CSSM_EncryptDataFinal.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_EncryptDataFinal
func CSSM_EncryptDataFinal(CCHandle CSSM_CC_HANDLE, RemData unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_EncryptDataFinal(CCHandle, RemData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_EncryptDataInit func(CCHandle CSSM_CC_HANDLE) CSSM_RETURN
var _cSSM_EncryptDataInitErr error

func tryCSSM_EncryptDataInit(CCHandle CSSM_CC_HANDLE) (CSSM_RETURN, error) {
	if _cSSM_EncryptDataInit == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_EncryptDataInit", "10.0", _cSSM_EncryptDataInitErr)
	}
	return _cSSM_EncryptDataInit(CCHandle), nil
}

// CSSM_EncryptDataInit.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_EncryptDataInit
func CSSM_EncryptDataInit(CCHandle CSSM_CC_HANDLE) CSSM_RETURN {
	result, callErr := tryCSSM_EncryptDataInit(CCHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_EncryptDataInitP func(CCHandle CSSM_CC_HANDLE, Privilege CSSM_PRIVILEGE) CSSM_RETURN
var _cSSM_EncryptDataInitPErr error

func tryCSSM_EncryptDataInitP(CCHandle CSSM_CC_HANDLE, Privilege CSSM_PRIVILEGE) (CSSM_RETURN, error) {
	if _cSSM_EncryptDataInitP == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_EncryptDataInitP", "10.0", _cSSM_EncryptDataInitPErr)
	}
	return _cSSM_EncryptDataInitP(CCHandle, Privilege), nil
}

// CSSM_EncryptDataInitP.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_EncryptDataInitP
func CSSM_EncryptDataInitP(CCHandle CSSM_CC_HANDLE, Privilege CSSM_PRIVILEGE) CSSM_RETURN {
	result, callErr := tryCSSM_EncryptDataInitP(CCHandle, Privilege)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_EncryptDataP func(CCHandle CSSM_CC_HANDLE, ClearBufs unsafe.Pointer, ClearBufCount uint32, CipherBufs unsafe.Pointer, CipherBufCount uint32, bytesEncrypted unsafe.Pointer, RemData unsafe.Pointer, Privilege CSSM_PRIVILEGE) CSSM_RETURN
var _cSSM_EncryptDataPErr error

func tryCSSM_EncryptDataP(CCHandle CSSM_CC_HANDLE, ClearBufs unsafe.Pointer, ClearBufCount uint32, CipherBufs unsafe.Pointer, CipherBufCount uint32, bytesEncrypted unsafe.Pointer, RemData unsafe.Pointer, Privilege CSSM_PRIVILEGE) (CSSM_RETURN, error) {
	if _cSSM_EncryptDataP == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_EncryptDataP", "10.0", _cSSM_EncryptDataPErr)
	}
	return _cSSM_EncryptDataP(CCHandle, ClearBufs, ClearBufCount, CipherBufs, CipherBufCount, bytesEncrypted, RemData, Privilege), nil
}

// CSSM_EncryptDataP.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_EncryptDataP
func CSSM_EncryptDataP(CCHandle CSSM_CC_HANDLE, ClearBufs unsafe.Pointer, ClearBufCount uint32, CipherBufs unsafe.Pointer, CipherBufCount uint32, bytesEncrypted unsafe.Pointer, RemData unsafe.Pointer, Privilege CSSM_PRIVILEGE) CSSM_RETURN {
	result, callErr := tryCSSM_EncryptDataP(CCHandle, ClearBufs, ClearBufCount, CipherBufs, CipherBufCount, bytesEncrypted, RemData, Privilege)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_EncryptDataUpdate func(CCHandle CSSM_CC_HANDLE, ClearBufs unsafe.Pointer, ClearBufCount uint32, CipherBufs unsafe.Pointer, CipherBufCount uint32, bytesEncrypted unsafe.Pointer) CSSM_RETURN
var _cSSM_EncryptDataUpdateErr error

func tryCSSM_EncryptDataUpdate(CCHandle CSSM_CC_HANDLE, ClearBufs unsafe.Pointer, ClearBufCount uint32, CipherBufs unsafe.Pointer, CipherBufCount uint32, bytesEncrypted unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_EncryptDataUpdate == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_EncryptDataUpdate", "10.0", _cSSM_EncryptDataUpdateErr)
	}
	return _cSSM_EncryptDataUpdate(CCHandle, ClearBufs, ClearBufCount, CipherBufs, CipherBufCount, bytesEncrypted), nil
}

// CSSM_EncryptDataUpdate.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_EncryptDataUpdate
func CSSM_EncryptDataUpdate(CCHandle CSSM_CC_HANDLE, ClearBufs unsafe.Pointer, ClearBufCount uint32, CipherBufs unsafe.Pointer, CipherBufCount uint32, bytesEncrypted unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_EncryptDataUpdate(CCHandle, ClearBufs, ClearBufCount, CipherBufs, CipherBufCount, bytesEncrypted)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_FreeContext func(Context unsafe.Pointer) CSSM_RETURN
var _cSSM_FreeContextErr error

func tryCSSM_FreeContext(Context unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_FreeContext == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_FreeContext", "10.0", _cSSM_FreeContextErr)
	}
	return _cSSM_FreeContext(Context), nil
}

// CSSM_FreeContext.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_FreeContext
func CSSM_FreeContext(Context unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_FreeContext(Context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_FreeKey func(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, KeyPtr unsafe.Pointer, Delete CSSM_BOOL) CSSM_RETURN
var _cSSM_FreeKeyErr error

func tryCSSM_FreeKey(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, KeyPtr unsafe.Pointer, Delete CSSM_BOOL) (CSSM_RETURN, error) {
	if _cSSM_FreeKey == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_FreeKey", "10.0", _cSSM_FreeKeyErr)
	}
	return _cSSM_FreeKey(CSPHandle, AccessCred, KeyPtr, Delete), nil
}

// CSSM_FreeKey.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_FreeKey
func CSSM_FreeKey(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, KeyPtr unsafe.Pointer, Delete CSSM_BOOL) CSSM_RETURN {
	result, callErr := tryCSSM_FreeKey(CSPHandle, AccessCred, KeyPtr, Delete)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_GenerateAlgorithmParams func(CCHandle CSSM_CC_HANDLE, ParamBits uint32, Param unsafe.Pointer) CSSM_RETURN
var _cSSM_GenerateAlgorithmParamsErr error

func tryCSSM_GenerateAlgorithmParams(CCHandle CSSM_CC_HANDLE, ParamBits uint32, Param unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_GenerateAlgorithmParams == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_GenerateAlgorithmParams", "10.0", _cSSM_GenerateAlgorithmParamsErr)
	}
	return _cSSM_GenerateAlgorithmParams(CCHandle, ParamBits, Param), nil
}

// CSSM_GenerateAlgorithmParams.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GenerateAlgorithmParams
func CSSM_GenerateAlgorithmParams(CCHandle CSSM_CC_HANDLE, ParamBits uint32, Param unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_GenerateAlgorithmParams(CCHandle, ParamBits, Param)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_GenerateKey func(CCHandle CSSM_CC_HANDLE, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, Key unsafe.Pointer) CSSM_RETURN
var _cSSM_GenerateKeyErr error

func tryCSSM_GenerateKey(CCHandle CSSM_CC_HANDLE, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, Key unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_GenerateKey == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_GenerateKey", "10.0", _cSSM_GenerateKeyErr)
	}
	return _cSSM_GenerateKey(CCHandle, KeyUsage, KeyAttr, KeyLabel, CredAndAclEntry, Key), nil
}

// CSSM_GenerateKey.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GenerateKey
func CSSM_GenerateKey(CCHandle CSSM_CC_HANDLE, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, Key unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_GenerateKey(CCHandle, KeyUsage, KeyAttr, KeyLabel, CredAndAclEntry, Key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_GenerateKeyP func(CCHandle CSSM_CC_HANDLE, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, Key unsafe.Pointer, Privilege CSSM_PRIVILEGE) CSSM_RETURN
var _cSSM_GenerateKeyPErr error

func tryCSSM_GenerateKeyP(CCHandle CSSM_CC_HANDLE, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, Key unsafe.Pointer, Privilege CSSM_PRIVILEGE) (CSSM_RETURN, error) {
	if _cSSM_GenerateKeyP == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_GenerateKeyP", "10.0", _cSSM_GenerateKeyPErr)
	}
	return _cSSM_GenerateKeyP(CCHandle, KeyUsage, KeyAttr, KeyLabel, CredAndAclEntry, Key, Privilege), nil
}

// CSSM_GenerateKeyP.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GenerateKeyP
func CSSM_GenerateKeyP(CCHandle CSSM_CC_HANDLE, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, Key unsafe.Pointer, Privilege CSSM_PRIVILEGE) CSSM_RETURN {
	result, callErr := tryCSSM_GenerateKeyP(CCHandle, KeyUsage, KeyAttr, KeyLabel, CredAndAclEntry, Key, Privilege)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_GenerateKeyPair func(CCHandle CSSM_CC_HANDLE, PublicKeyUsage uint32, PublicKeyAttr uint32, PublicKeyLabel unsafe.Pointer, PublicKey unsafe.Pointer, PrivateKeyUsage uint32, PrivateKeyAttr uint32, PrivateKeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, PrivateKey unsafe.Pointer) CSSM_RETURN
var _cSSM_GenerateKeyPairErr error

func tryCSSM_GenerateKeyPair(CCHandle CSSM_CC_HANDLE, PublicKeyUsage uint32, PublicKeyAttr uint32, PublicKeyLabel unsafe.Pointer, PublicKey unsafe.Pointer, PrivateKeyUsage uint32, PrivateKeyAttr uint32, PrivateKeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, PrivateKey unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_GenerateKeyPair == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_GenerateKeyPair", "10.0", _cSSM_GenerateKeyPairErr)
	}
	return _cSSM_GenerateKeyPair(CCHandle, PublicKeyUsage, PublicKeyAttr, PublicKeyLabel, PublicKey, PrivateKeyUsage, PrivateKeyAttr, PrivateKeyLabel, CredAndAclEntry, PrivateKey), nil
}

// CSSM_GenerateKeyPair.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GenerateKeyPair
func CSSM_GenerateKeyPair(CCHandle CSSM_CC_HANDLE, PublicKeyUsage uint32, PublicKeyAttr uint32, PublicKeyLabel unsafe.Pointer, PublicKey unsafe.Pointer, PrivateKeyUsage uint32, PrivateKeyAttr uint32, PrivateKeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, PrivateKey unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_GenerateKeyPair(CCHandle, PublicKeyUsage, PublicKeyAttr, PublicKeyLabel, PublicKey, PrivateKeyUsage, PrivateKeyAttr, PrivateKeyLabel, CredAndAclEntry, PrivateKey)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_GenerateKeyPairP func(CCHandle CSSM_CC_HANDLE, PublicKeyUsage uint32, PublicKeyAttr uint32, PublicKeyLabel unsafe.Pointer, PublicKey unsafe.Pointer, PrivateKeyUsage uint32, PrivateKeyAttr uint32, PrivateKeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, PrivateKey unsafe.Pointer, Privilege CSSM_PRIVILEGE) CSSM_RETURN
var _cSSM_GenerateKeyPairPErr error

func tryCSSM_GenerateKeyPairP(CCHandle CSSM_CC_HANDLE, PublicKeyUsage uint32, PublicKeyAttr uint32, PublicKeyLabel unsafe.Pointer, PublicKey unsafe.Pointer, PrivateKeyUsage uint32, PrivateKeyAttr uint32, PrivateKeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, PrivateKey unsafe.Pointer, Privilege CSSM_PRIVILEGE) (CSSM_RETURN, error) {
	if _cSSM_GenerateKeyPairP == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_GenerateKeyPairP", "10.0", _cSSM_GenerateKeyPairPErr)
	}
	return _cSSM_GenerateKeyPairP(CCHandle, PublicKeyUsage, PublicKeyAttr, PublicKeyLabel, PublicKey, PrivateKeyUsage, PrivateKeyAttr, PrivateKeyLabel, CredAndAclEntry, PrivateKey, Privilege), nil
}

// CSSM_GenerateKeyPairP.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GenerateKeyPairP
func CSSM_GenerateKeyPairP(CCHandle CSSM_CC_HANDLE, PublicKeyUsage uint32, PublicKeyAttr uint32, PublicKeyLabel unsafe.Pointer, PublicKey unsafe.Pointer, PrivateKeyUsage uint32, PrivateKeyAttr uint32, PrivateKeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, PrivateKey unsafe.Pointer, Privilege CSSM_PRIVILEGE) CSSM_RETURN {
	result, callErr := tryCSSM_GenerateKeyPairP(CCHandle, PublicKeyUsage, PublicKeyAttr, PublicKeyLabel, PublicKey, PrivateKeyUsage, PrivateKeyAttr, PrivateKeyLabel, CredAndAclEntry, PrivateKey, Privilege)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_GenerateMac func(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, Mac unsafe.Pointer) CSSM_RETURN
var _cSSM_GenerateMacErr error

func tryCSSM_GenerateMac(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, Mac unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_GenerateMac == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_GenerateMac", "10.0", _cSSM_GenerateMacErr)
	}
	return _cSSM_GenerateMac(CCHandle, DataBufs, DataBufCount, Mac), nil
}

// CSSM_GenerateMac.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GenerateMac
func CSSM_GenerateMac(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, Mac unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_GenerateMac(CCHandle, DataBufs, DataBufCount, Mac)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_GenerateMacFinal func(CCHandle CSSM_CC_HANDLE, Mac unsafe.Pointer) CSSM_RETURN
var _cSSM_GenerateMacFinalErr error

func tryCSSM_GenerateMacFinal(CCHandle CSSM_CC_HANDLE, Mac unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_GenerateMacFinal == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_GenerateMacFinal", "10.0", _cSSM_GenerateMacFinalErr)
	}
	return _cSSM_GenerateMacFinal(CCHandle, Mac), nil
}

// CSSM_GenerateMacFinal.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GenerateMacFinal
func CSSM_GenerateMacFinal(CCHandle CSSM_CC_HANDLE, Mac unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_GenerateMacFinal(CCHandle, Mac)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_GenerateMacInit func(CCHandle CSSM_CC_HANDLE) CSSM_RETURN
var _cSSM_GenerateMacInitErr error

func tryCSSM_GenerateMacInit(CCHandle CSSM_CC_HANDLE) (CSSM_RETURN, error) {
	if _cSSM_GenerateMacInit == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_GenerateMacInit", "10.0", _cSSM_GenerateMacInitErr)
	}
	return _cSSM_GenerateMacInit(CCHandle), nil
}

// CSSM_GenerateMacInit.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GenerateMacInit
func CSSM_GenerateMacInit(CCHandle CSSM_CC_HANDLE) CSSM_RETURN {
	result, callErr := tryCSSM_GenerateMacInit(CCHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_GenerateMacUpdate func(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) CSSM_RETURN
var _cSSM_GenerateMacUpdateErr error

func tryCSSM_GenerateMacUpdate(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) (CSSM_RETURN, error) {
	if _cSSM_GenerateMacUpdate == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_GenerateMacUpdate", "10.0", _cSSM_GenerateMacUpdateErr)
	}
	return _cSSM_GenerateMacUpdate(CCHandle, DataBufs, DataBufCount), nil
}

// CSSM_GenerateMacUpdate.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GenerateMacUpdate
func CSSM_GenerateMacUpdate(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) CSSM_RETURN {
	result, callErr := tryCSSM_GenerateMacUpdate(CCHandle, DataBufs, DataBufCount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_GenerateRandom func(CCHandle CSSM_CC_HANDLE, RandomNumber unsafe.Pointer) CSSM_RETURN
var _cSSM_GenerateRandomErr error

func tryCSSM_GenerateRandom(CCHandle CSSM_CC_HANDLE, RandomNumber unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_GenerateRandom == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_GenerateRandom", "10.0", _cSSM_GenerateRandomErr)
	}
	return _cSSM_GenerateRandom(CCHandle, RandomNumber), nil
}

// CSSM_GenerateRandom.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GenerateRandom
func CSSM_GenerateRandom(CCHandle CSSM_CC_HANDLE, RandomNumber unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_GenerateRandom(CCHandle, RandomNumber)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_GetAPIMemoryFunctions func(AddInHandle CSSM_MODULE_HANDLE, AppMemoryFuncs unsafe.Pointer) CSSM_RETURN
var _cSSM_GetAPIMemoryFunctionsErr error

func tryCSSM_GetAPIMemoryFunctions(AddInHandle CSSM_MODULE_HANDLE, AppMemoryFuncs unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_GetAPIMemoryFunctions == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_GetAPIMemoryFunctions", "10.0", _cSSM_GetAPIMemoryFunctionsErr)
	}
	return _cSSM_GetAPIMemoryFunctions(AddInHandle, AppMemoryFuncs), nil
}

// CSSM_GetAPIMemoryFunctions.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GetAPIMemoryFunctions
func CSSM_GetAPIMemoryFunctions(AddInHandle CSSM_MODULE_HANDLE, AppMemoryFuncs unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_GetAPIMemoryFunctions(AddInHandle, AppMemoryFuncs)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_GetContext func(CCHandle CSSM_CC_HANDLE, Context unsafe.Pointer) CSSM_RETURN
var _cSSM_GetContextErr error

func tryCSSM_GetContext(CCHandle CSSM_CC_HANDLE, Context unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_GetContext == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_GetContext", "10.0", _cSSM_GetContextErr)
	}
	return _cSSM_GetContext(CCHandle, Context), nil
}

// CSSM_GetContext.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GetContext
func CSSM_GetContext(CCHandle CSSM_CC_HANDLE, Context unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_GetContext(CCHandle, Context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_GetContextAttribute func(Context unsafe.Pointer, AttributeType uint32, ContextAttribute unsafe.Pointer) CSSM_RETURN
var _cSSM_GetContextAttributeErr error

func tryCSSM_GetContextAttribute(Context unsafe.Pointer, AttributeType uint32, ContextAttribute unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_GetContextAttribute == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_GetContextAttribute", "10.0", _cSSM_GetContextAttributeErr)
	}
	return _cSSM_GetContextAttribute(Context, AttributeType, ContextAttribute), nil
}

// CSSM_GetContextAttribute.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GetContextAttribute
func CSSM_GetContextAttribute(Context unsafe.Pointer, AttributeType uint32, ContextAttribute unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_GetContextAttribute(Context, AttributeType, ContextAttribute)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_GetKeyAcl func(CSPHandle CSSM_CSP_HANDLE, Key unsafe.Pointer, SelectionTag unsafe.Pointer, NumberOfAclInfos *uint32, AclInfos unsafe.Pointer) CSSM_RETURN
var _cSSM_GetKeyAclErr error

func tryCSSM_GetKeyAcl(CSPHandle CSSM_CSP_HANDLE, Key unsafe.Pointer, SelectionTag unsafe.Pointer, NumberOfAclInfos *uint32, AclInfos unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_GetKeyAcl == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_GetKeyAcl", "10.0", _cSSM_GetKeyAclErr)
	}
	return _cSSM_GetKeyAcl(CSPHandle, Key, SelectionTag, NumberOfAclInfos, AclInfos), nil
}

// CSSM_GetKeyAcl.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GetKeyAcl
func CSSM_GetKeyAcl(CSPHandle CSSM_CSP_HANDLE, Key unsafe.Pointer, SelectionTag unsafe.Pointer, NumberOfAclInfos *uint32, AclInfos unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_GetKeyAcl(CSPHandle, Key, SelectionTag, NumberOfAclInfos, AclInfos)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_GetKeyOwner func(CSPHandle CSSM_CSP_HANDLE, Key unsafe.Pointer, Owner unsafe.Pointer) CSSM_RETURN
var _cSSM_GetKeyOwnerErr error

func tryCSSM_GetKeyOwner(CSPHandle CSSM_CSP_HANDLE, Key unsafe.Pointer, Owner unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_GetKeyOwner == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_GetKeyOwner", "10.0", _cSSM_GetKeyOwnerErr)
	}
	return _cSSM_GetKeyOwner(CSPHandle, Key, Owner), nil
}

// CSSM_GetKeyOwner.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GetKeyOwner
func CSSM_GetKeyOwner(CSPHandle CSSM_CSP_HANDLE, Key unsafe.Pointer, Owner unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_GetKeyOwner(CSPHandle, Key, Owner)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_GetModuleGUIDFromHandle func(ModuleHandle CSSM_MODULE_HANDLE, ModuleGUID unsafe.Pointer) CSSM_RETURN
var _cSSM_GetModuleGUIDFromHandleErr error

func tryCSSM_GetModuleGUIDFromHandle(ModuleHandle CSSM_MODULE_HANDLE, ModuleGUID unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_GetModuleGUIDFromHandle == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_GetModuleGUIDFromHandle", "10.0", _cSSM_GetModuleGUIDFromHandleErr)
	}
	return _cSSM_GetModuleGUIDFromHandle(ModuleHandle, ModuleGUID), nil
}

// CSSM_GetModuleGUIDFromHandle.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GetModuleGUIDFromHandle
func CSSM_GetModuleGUIDFromHandle(ModuleHandle CSSM_MODULE_HANDLE, ModuleGUID unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_GetModuleGUIDFromHandle(ModuleHandle, ModuleGUID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_GetPrivilege func(Privilege unsafe.Pointer) CSSM_RETURN
var _cSSM_GetPrivilegeErr error

func tryCSSM_GetPrivilege(Privilege unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_GetPrivilege == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_GetPrivilege", "10.0", _cSSM_GetPrivilegeErr)
	}
	return _cSSM_GetPrivilege(Privilege), nil
}

// CSSM_GetPrivilege.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GetPrivilege
func CSSM_GetPrivilege(Privilege unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_GetPrivilege(Privilege)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_GetSubserviceUIDFromHandle func(ModuleHandle CSSM_MODULE_HANDLE, SubserviceUID unsafe.Pointer) CSSM_RETURN
var _cSSM_GetSubserviceUIDFromHandleErr error

func tryCSSM_GetSubserviceUIDFromHandle(ModuleHandle CSSM_MODULE_HANDLE, SubserviceUID unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_GetSubserviceUIDFromHandle == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_GetSubserviceUIDFromHandle", "10.0", _cSSM_GetSubserviceUIDFromHandleErr)
	}
	return _cSSM_GetSubserviceUIDFromHandle(ModuleHandle, SubserviceUID), nil
}

// CSSM_GetSubserviceUIDFromHandle.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GetSubserviceUIDFromHandle
func CSSM_GetSubserviceUIDFromHandle(ModuleHandle CSSM_MODULE_HANDLE, SubserviceUID unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_GetSubserviceUIDFromHandle(ModuleHandle, SubserviceUID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_GetTimeValue func(CSPHandle CSSM_CSP_HANDLE, TimeAlgorithm CSSM_ALGORITHMS, TimeData unsafe.Pointer) CSSM_RETURN
var _cSSM_GetTimeValueErr error

func tryCSSM_GetTimeValue(CSPHandle CSSM_CSP_HANDLE, TimeAlgorithm CSSM_ALGORITHMS, TimeData unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_GetTimeValue == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_GetTimeValue", "10.0", _cSSM_GetTimeValueErr)
	}
	return _cSSM_GetTimeValue(CSPHandle, TimeAlgorithm, TimeData), nil
}

// CSSM_GetTimeValue.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GetTimeValue
func CSSM_GetTimeValue(CSPHandle CSSM_CSP_HANDLE, TimeAlgorithm CSSM_ALGORITHMS, TimeData unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_GetTimeValue(CSPHandle, TimeAlgorithm, TimeData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_Init func(Version unsafe.Pointer, Scope CSSM_PRIVILEGE_SCOPE, CallerGuid unsafe.Pointer, KeyHierarchy CSSM_KEY_HIERARCHY, PvcPolicy unsafe.Pointer, Reserved unsafe.Pointer) CSSM_RETURN
var _cSSM_InitErr error

func tryCSSM_Init(Version unsafe.Pointer, Scope CSSM_PRIVILEGE_SCOPE, CallerGuid unsafe.Pointer, KeyHierarchy CSSM_KEY_HIERARCHY, PvcPolicy unsafe.Pointer, Reserved unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_Init == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_Init", "10.0", _cSSM_InitErr)
	}
	return _cSSM_Init(Version, Scope, CallerGuid, KeyHierarchy, PvcPolicy, Reserved), nil
}

// CSSM_Init.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_Init
func CSSM_Init(Version unsafe.Pointer, Scope CSSM_PRIVILEGE_SCOPE, CallerGuid unsafe.Pointer, KeyHierarchy CSSM_KEY_HIERARCHY, PvcPolicy unsafe.Pointer, Reserved unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_Init(Version, Scope, CallerGuid, KeyHierarchy, PvcPolicy, Reserved)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_Introduce func(ModuleID unsafe.Pointer, KeyHierarchy CSSM_KEY_HIERARCHY) CSSM_RETURN
var _cSSM_IntroduceErr error

func tryCSSM_Introduce(ModuleID unsafe.Pointer, KeyHierarchy CSSM_KEY_HIERARCHY) (CSSM_RETURN, error) {
	if _cSSM_Introduce == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_Introduce", "10.0", _cSSM_IntroduceErr)
	}
	return _cSSM_Introduce(ModuleID, KeyHierarchy), nil
}

// CSSM_Introduce.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_Introduce
func CSSM_Introduce(ModuleID unsafe.Pointer, KeyHierarchy CSSM_KEY_HIERARCHY) CSSM_RETURN {
	result, callErr := tryCSSM_Introduce(ModuleID, KeyHierarchy)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_ListAttachedModuleManagers func(NumberOfModuleManagers *uint32, ModuleManagerGuids unsafe.Pointer) CSSM_RETURN
var _cSSM_ListAttachedModuleManagersErr error

func tryCSSM_ListAttachedModuleManagers(NumberOfModuleManagers *uint32, ModuleManagerGuids unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_ListAttachedModuleManagers == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_ListAttachedModuleManagers", "10.0", _cSSM_ListAttachedModuleManagersErr)
	}
	return _cSSM_ListAttachedModuleManagers(NumberOfModuleManagers, ModuleManagerGuids), nil
}

// CSSM_ListAttachedModuleManagers.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_ListAttachedModuleManagers
func CSSM_ListAttachedModuleManagers(NumberOfModuleManagers *uint32, ModuleManagerGuids unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_ListAttachedModuleManagers(NumberOfModuleManagers, ModuleManagerGuids)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_ModuleAttach func(ModuleGuid unsafe.Pointer, Version unsafe.Pointer, MemoryFuncs unsafe.Pointer, SubserviceID uint32, SubServiceType CSSM_SERVICE_TYPE, AttachFlags CSSM_ATTACH_FLAGS, KeyHierarchy CSSM_KEY_HIERARCHY, FunctionTable unsafe.Pointer, NumFunctionTable uint32, Reserved unsafe.Pointer, NewModuleHandle CSSM_MODULE_HANDLE_PTR) CSSM_RETURN
var _cSSM_ModuleAttachErr error

func tryCSSM_ModuleAttach(ModuleGuid unsafe.Pointer, Version unsafe.Pointer, MemoryFuncs unsafe.Pointer, SubserviceID uint32, SubServiceType CSSM_SERVICE_TYPE, AttachFlags CSSM_ATTACH_FLAGS, KeyHierarchy CSSM_KEY_HIERARCHY, FunctionTable unsafe.Pointer, NumFunctionTable uint32, Reserved unsafe.Pointer, NewModuleHandle CSSM_MODULE_HANDLE_PTR) (CSSM_RETURN, error) {
	if _cSSM_ModuleAttach == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_ModuleAttach", "10.0", _cSSM_ModuleAttachErr)
	}
	return _cSSM_ModuleAttach(ModuleGuid, Version, MemoryFuncs, SubserviceID, SubServiceType, AttachFlags, KeyHierarchy, FunctionTable, NumFunctionTable, Reserved, NewModuleHandle), nil
}

// CSSM_ModuleAttach.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_ModuleAttach
func CSSM_ModuleAttach(ModuleGuid unsafe.Pointer, Version unsafe.Pointer, MemoryFuncs unsafe.Pointer, SubserviceID uint32, SubServiceType CSSM_SERVICE_TYPE, AttachFlags CSSM_ATTACH_FLAGS, KeyHierarchy CSSM_KEY_HIERARCHY, FunctionTable unsafe.Pointer, NumFunctionTable uint32, Reserved unsafe.Pointer, NewModuleHandle CSSM_MODULE_HANDLE_PTR) CSSM_RETURN {
	result, callErr := tryCSSM_ModuleAttach(ModuleGuid, Version, MemoryFuncs, SubserviceID, SubServiceType, AttachFlags, KeyHierarchy, FunctionTable, NumFunctionTable, Reserved, NewModuleHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_ModuleDetach func(ModuleHandle CSSM_MODULE_HANDLE) CSSM_RETURN
var _cSSM_ModuleDetachErr error

func tryCSSM_ModuleDetach(ModuleHandle CSSM_MODULE_HANDLE) (CSSM_RETURN, error) {
	if _cSSM_ModuleDetach == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_ModuleDetach", "10.0", _cSSM_ModuleDetachErr)
	}
	return _cSSM_ModuleDetach(ModuleHandle), nil
}

// CSSM_ModuleDetach.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_ModuleDetach
func CSSM_ModuleDetach(ModuleHandle CSSM_MODULE_HANDLE) CSSM_RETURN {
	result, callErr := tryCSSM_ModuleDetach(ModuleHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_ModuleLoad func(ModuleGuid unsafe.Pointer, KeyHierarchy CSSM_KEY_HIERARCHY, AppNotifyCallback unsafe.Pointer, AppNotifyCallbackCtx unsafe.Pointer) CSSM_RETURN
var _cSSM_ModuleLoadErr error

func tryCSSM_ModuleLoad(ModuleGuid unsafe.Pointer, KeyHierarchy CSSM_KEY_HIERARCHY, AppNotifyCallback unsafe.Pointer, AppNotifyCallbackCtx unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_ModuleLoad == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_ModuleLoad", "10.0", _cSSM_ModuleLoadErr)
	}
	return _cSSM_ModuleLoad(ModuleGuid, KeyHierarchy, AppNotifyCallback, AppNotifyCallbackCtx), nil
}

// CSSM_ModuleLoad.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_ModuleLoad
func CSSM_ModuleLoad(ModuleGuid unsafe.Pointer, KeyHierarchy CSSM_KEY_HIERARCHY, AppNotifyCallback unsafe.Pointer, AppNotifyCallbackCtx unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_ModuleLoad(ModuleGuid, KeyHierarchy, AppNotifyCallback, AppNotifyCallbackCtx)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_ModuleUnload func(ModuleGuid unsafe.Pointer, AppNotifyCallback unsafe.Pointer, AppNotifyCallbackCtx unsafe.Pointer) CSSM_RETURN
var _cSSM_ModuleUnloadErr error

func tryCSSM_ModuleUnload(ModuleGuid unsafe.Pointer, AppNotifyCallback unsafe.Pointer, AppNotifyCallbackCtx unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_ModuleUnload == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_ModuleUnload", "10.0", _cSSM_ModuleUnloadErr)
	}
	return _cSSM_ModuleUnload(ModuleGuid, AppNotifyCallback, AppNotifyCallbackCtx), nil
}

// CSSM_ModuleUnload.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_ModuleUnload
func CSSM_ModuleUnload(ModuleGuid unsafe.Pointer, AppNotifyCallback unsafe.Pointer, AppNotifyCallbackCtx unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_ModuleUnload(ModuleGuid, AppNotifyCallback, AppNotifyCallbackCtx)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_QueryKeySizeInBits func(CSPHandle CSSM_CSP_HANDLE, CCHandle CSSM_CC_HANDLE, Key unsafe.Pointer, KeySize unsafe.Pointer) CSSM_RETURN
var _cSSM_QueryKeySizeInBitsErr error

func tryCSSM_QueryKeySizeInBits(CSPHandle CSSM_CSP_HANDLE, CCHandle CSSM_CC_HANDLE, Key unsafe.Pointer, KeySize unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_QueryKeySizeInBits == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_QueryKeySizeInBits", "10.0", _cSSM_QueryKeySizeInBitsErr)
	}
	return _cSSM_QueryKeySizeInBits(CSPHandle, CCHandle, Key, KeySize), nil
}

// CSSM_QueryKeySizeInBits.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_QueryKeySizeInBits
func CSSM_QueryKeySizeInBits(CSPHandle CSSM_CSP_HANDLE, CCHandle CSSM_CC_HANDLE, Key unsafe.Pointer, KeySize unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_QueryKeySizeInBits(CSPHandle, CCHandle, Key, KeySize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_QuerySize func(CCHandle CSSM_CC_HANDLE, Encrypt CSSM_BOOL, QuerySizeCount uint32, DataBlockSizes unsafe.Pointer) CSSM_RETURN
var _cSSM_QuerySizeErr error

func tryCSSM_QuerySize(CCHandle CSSM_CC_HANDLE, Encrypt CSSM_BOOL, QuerySizeCount uint32, DataBlockSizes unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_QuerySize == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_QuerySize", "10.0", _cSSM_QuerySizeErr)
	}
	return _cSSM_QuerySize(CCHandle, Encrypt, QuerySizeCount, DataBlockSizes), nil
}

// CSSM_QuerySize.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_QuerySize
func CSSM_QuerySize(CCHandle CSSM_CC_HANDLE, Encrypt CSSM_BOOL, QuerySizeCount uint32, DataBlockSizes unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_QuerySize(CCHandle, Encrypt, QuerySizeCount, DataBlockSizes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_RetrieveCounter func(CSPHandle CSSM_CSP_HANDLE, Counter unsafe.Pointer) CSSM_RETURN
var _cSSM_RetrieveCounterErr error

func tryCSSM_RetrieveCounter(CSPHandle CSSM_CSP_HANDLE, Counter unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_RetrieveCounter == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_RetrieveCounter", "10.0", _cSSM_RetrieveCounterErr)
	}
	return _cSSM_RetrieveCounter(CSPHandle, Counter), nil
}

// CSSM_RetrieveCounter.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_RetrieveCounter
func CSSM_RetrieveCounter(CSPHandle CSSM_CSP_HANDLE, Counter unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_RetrieveCounter(CSPHandle, Counter)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_RetrieveUniqueId func(CSPHandle CSSM_CSP_HANDLE, UniqueID unsafe.Pointer) CSSM_RETURN
var _cSSM_RetrieveUniqueIdErr error

func tryCSSM_RetrieveUniqueId(CSPHandle CSSM_CSP_HANDLE, UniqueID unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_RetrieveUniqueId == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_RetrieveUniqueId", "10.0", _cSSM_RetrieveUniqueIdErr)
	}
	return _cSSM_RetrieveUniqueId(CSPHandle, UniqueID), nil
}

// CSSM_RetrieveUniqueId.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_RetrieveUniqueId
func CSSM_RetrieveUniqueId(CSPHandle CSSM_CSP_HANDLE, UniqueID unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_RetrieveUniqueId(CSPHandle, UniqueID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_SetContext func(CCHandle CSSM_CC_HANDLE, Context unsafe.Pointer) CSSM_RETURN
var _cSSM_SetContextErr error

func tryCSSM_SetContext(CCHandle CSSM_CC_HANDLE, Context unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_SetContext == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_SetContext", "10.0", _cSSM_SetContextErr)
	}
	return _cSSM_SetContext(CCHandle, Context), nil
}

// CSSM_SetContext.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_SetContext
func CSSM_SetContext(CCHandle CSSM_CC_HANDLE, Context unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_SetContext(CCHandle, Context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_SetPrivilege func(Privilege CSSM_PRIVILEGE) CSSM_RETURN
var _cSSM_SetPrivilegeErr error

func tryCSSM_SetPrivilege(Privilege CSSM_PRIVILEGE) (CSSM_RETURN, error) {
	if _cSSM_SetPrivilege == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_SetPrivilege", "10.0", _cSSM_SetPrivilegeErr)
	}
	return _cSSM_SetPrivilege(Privilege), nil
}

// CSSM_SetPrivilege.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_SetPrivilege
func CSSM_SetPrivilege(Privilege CSSM_PRIVILEGE) CSSM_RETURN {
	result, callErr := tryCSSM_SetPrivilege(Privilege)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_SignData func(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, DigestAlgorithm CSSM_ALGORITHMS, Signature unsafe.Pointer) CSSM_RETURN
var _cSSM_SignDataErr error

func tryCSSM_SignData(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, DigestAlgorithm CSSM_ALGORITHMS, Signature unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_SignData == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_SignData", "10.0", _cSSM_SignDataErr)
	}
	return _cSSM_SignData(CCHandle, DataBufs, DataBufCount, DigestAlgorithm, Signature), nil
}

// CSSM_SignData.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_SignData
func CSSM_SignData(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, DigestAlgorithm CSSM_ALGORITHMS, Signature unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_SignData(CCHandle, DataBufs, DataBufCount, DigestAlgorithm, Signature)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_SignDataFinal func(CCHandle CSSM_CC_HANDLE, Signature unsafe.Pointer) CSSM_RETURN
var _cSSM_SignDataFinalErr error

func tryCSSM_SignDataFinal(CCHandle CSSM_CC_HANDLE, Signature unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_SignDataFinal == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_SignDataFinal", "10.0", _cSSM_SignDataFinalErr)
	}
	return _cSSM_SignDataFinal(CCHandle, Signature), nil
}

// CSSM_SignDataFinal.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_SignDataFinal
func CSSM_SignDataFinal(CCHandle CSSM_CC_HANDLE, Signature unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_SignDataFinal(CCHandle, Signature)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_SignDataInit func(CCHandle CSSM_CC_HANDLE) CSSM_RETURN
var _cSSM_SignDataInitErr error

func tryCSSM_SignDataInit(CCHandle CSSM_CC_HANDLE) (CSSM_RETURN, error) {
	if _cSSM_SignDataInit == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_SignDataInit", "10.0", _cSSM_SignDataInitErr)
	}
	return _cSSM_SignDataInit(CCHandle), nil
}

// CSSM_SignDataInit.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_SignDataInit
func CSSM_SignDataInit(CCHandle CSSM_CC_HANDLE) CSSM_RETURN {
	result, callErr := tryCSSM_SignDataInit(CCHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_SignDataUpdate func(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) CSSM_RETURN
var _cSSM_SignDataUpdateErr error

func tryCSSM_SignDataUpdate(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) (CSSM_RETURN, error) {
	if _cSSM_SignDataUpdate == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_SignDataUpdate", "10.0", _cSSM_SignDataUpdateErr)
	}
	return _cSSM_SignDataUpdate(CCHandle, DataBufs, DataBufCount), nil
}

// CSSM_SignDataUpdate.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_SignDataUpdate
func CSSM_SignDataUpdate(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) CSSM_RETURN {
	result, callErr := tryCSSM_SignDataUpdate(CCHandle, DataBufs, DataBufCount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_ApplyCrlToDb func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, CrlToBeApplied unsafe.Pointer, SignerCertGroup unsafe.Pointer, ApplyCrlVerifyContext unsafe.Pointer, ApplyCrlVerifyResult unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_ApplyCrlToDbErr error

func tryCSSM_TP_ApplyCrlToDb(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, CrlToBeApplied unsafe.Pointer, SignerCertGroup unsafe.Pointer, ApplyCrlVerifyContext unsafe.Pointer, ApplyCrlVerifyResult unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_TP_ApplyCrlToDb == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_TP_ApplyCrlToDb", "10.0", _cSSM_TP_ApplyCrlToDbErr)
	}
	return _cSSM_TP_ApplyCrlToDb(TPHandle, CLHandle, CSPHandle, CrlToBeApplied, SignerCertGroup, ApplyCrlVerifyContext, ApplyCrlVerifyResult), nil
}

// CSSM_TP_ApplyCrlToDb.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_ApplyCrlToDb
func CSSM_TP_ApplyCrlToDb(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, CrlToBeApplied unsafe.Pointer, SignerCertGroup unsafe.Pointer, ApplyCrlVerifyContext unsafe.Pointer, ApplyCrlVerifyResult unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_TP_ApplyCrlToDb(TPHandle, CLHandle, CSPHandle, CrlToBeApplied, SignerCertGroup, ApplyCrlVerifyContext, ApplyCrlVerifyResult)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_CertCreateTemplate func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CertFields unsafe.Pointer, CertTemplate unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_CertCreateTemplateErr error

func tryCSSM_TP_CertCreateTemplate(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CertFields unsafe.Pointer, CertTemplate unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_TP_CertCreateTemplate == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_TP_CertCreateTemplate", "10.0", _cSSM_TP_CertCreateTemplateErr)
	}
	return _cSSM_TP_CertCreateTemplate(TPHandle, CLHandle, NumberOfFields, CertFields, CertTemplate), nil
}

// CSSM_TP_CertCreateTemplate.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CertCreateTemplate
func CSSM_TP_CertCreateTemplate(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CertFields unsafe.Pointer, CertTemplate unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_TP_CertCreateTemplate(TPHandle, CLHandle, NumberOfFields, CertFields, CertTemplate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_CertGetAllTemplateFields func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CertTemplate unsafe.Pointer, NumberOfFields *uint32, CertFields unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_CertGetAllTemplateFieldsErr error

func tryCSSM_TP_CertGetAllTemplateFields(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CertTemplate unsafe.Pointer, NumberOfFields *uint32, CertFields unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_TP_CertGetAllTemplateFields == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_TP_CertGetAllTemplateFields", "10.0", _cSSM_TP_CertGetAllTemplateFieldsErr)
	}
	return _cSSM_TP_CertGetAllTemplateFields(TPHandle, CLHandle, CertTemplate, NumberOfFields, CertFields), nil
}

// CSSM_TP_CertGetAllTemplateFields.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CertGetAllTemplateFields
func CSSM_TP_CertGetAllTemplateFields(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CertTemplate unsafe.Pointer, NumberOfFields *uint32, CertFields unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_TP_CertGetAllTemplateFields(TPHandle, CLHandle, CertTemplate, NumberOfFields, CertFields)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_CertGroupConstruct func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, DBList unsafe.Pointer, ConstructParams unsafe.Pointer, CertGroupFrag unsafe.Pointer, CertGroup unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_CertGroupConstructErr error

func tryCSSM_TP_CertGroupConstruct(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, DBList unsafe.Pointer, ConstructParams unsafe.Pointer, CertGroupFrag unsafe.Pointer, CertGroup unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_TP_CertGroupConstruct == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_TP_CertGroupConstruct", "10.0", _cSSM_TP_CertGroupConstructErr)
	}
	return _cSSM_TP_CertGroupConstruct(TPHandle, CLHandle, CSPHandle, DBList, ConstructParams, CertGroupFrag, CertGroup), nil
}

// CSSM_TP_CertGroupConstruct.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CertGroupConstruct
func CSSM_TP_CertGroupConstruct(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, DBList unsafe.Pointer, ConstructParams unsafe.Pointer, CertGroupFrag unsafe.Pointer, CertGroup unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_TP_CertGroupConstruct(TPHandle, CLHandle, CSPHandle, DBList, ConstructParams, CertGroupFrag, CertGroup)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_CertGroupPrune func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, DBList unsafe.Pointer, OrderedCertGroup unsafe.Pointer, PrunedCertGroup unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_CertGroupPruneErr error

func tryCSSM_TP_CertGroupPrune(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, DBList unsafe.Pointer, OrderedCertGroup unsafe.Pointer, PrunedCertGroup unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_TP_CertGroupPrune == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_TP_CertGroupPrune", "10.0", _cSSM_TP_CertGroupPruneErr)
	}
	return _cSSM_TP_CertGroupPrune(TPHandle, CLHandle, DBList, OrderedCertGroup, PrunedCertGroup), nil
}

// CSSM_TP_CertGroupPrune.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CertGroupPrune
func CSSM_TP_CertGroupPrune(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, DBList unsafe.Pointer, OrderedCertGroup unsafe.Pointer, PrunedCertGroup unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_TP_CertGroupPrune(TPHandle, CLHandle, DBList, OrderedCertGroup, PrunedCertGroup)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_CertGroupToTupleGroup func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CertGroup unsafe.Pointer, TupleGroup unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_CertGroupToTupleGroupErr error

func tryCSSM_TP_CertGroupToTupleGroup(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CertGroup unsafe.Pointer, TupleGroup unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_TP_CertGroupToTupleGroup == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_TP_CertGroupToTupleGroup", "10.0", _cSSM_TP_CertGroupToTupleGroupErr)
	}
	return _cSSM_TP_CertGroupToTupleGroup(TPHandle, CLHandle, CertGroup, TupleGroup), nil
}

// CSSM_TP_CertGroupToTupleGroup.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CertGroupToTupleGroup
func CSSM_TP_CertGroupToTupleGroup(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CertGroup unsafe.Pointer, TupleGroup unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_TP_CertGroupToTupleGroup(TPHandle, CLHandle, CertGroup, TupleGroup)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_CertGroupVerify func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, CertGroupToBeVerified unsafe.Pointer, VerifyContext unsafe.Pointer, VerifyContextResult unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_CertGroupVerifyErr error

func tryCSSM_TP_CertGroupVerify(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, CertGroupToBeVerified unsafe.Pointer, VerifyContext unsafe.Pointer, VerifyContextResult unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_TP_CertGroupVerify == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_TP_CertGroupVerify", "10.0", _cSSM_TP_CertGroupVerifyErr)
	}
	return _cSSM_TP_CertGroupVerify(TPHandle, CLHandle, CSPHandle, CertGroupToBeVerified, VerifyContext, VerifyContextResult), nil
}

// CSSM_TP_CertGroupVerify.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CertGroupVerify
func CSSM_TP_CertGroupVerify(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, CertGroupToBeVerified unsafe.Pointer, VerifyContext unsafe.Pointer, VerifyContextResult unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_TP_CertGroupVerify(TPHandle, CLHandle, CSPHandle, CertGroupToBeVerified, VerifyContext, VerifyContextResult)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_CertReclaimAbort func(TPHandle CSSM_TP_HANDLE, KeyCacheHandle CSSM_LONG_HANDLE) CSSM_RETURN
var _cSSM_TP_CertReclaimAbortErr error

func tryCSSM_TP_CertReclaimAbort(TPHandle CSSM_TP_HANDLE, KeyCacheHandle CSSM_LONG_HANDLE) (CSSM_RETURN, error) {
	if _cSSM_TP_CertReclaimAbort == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_TP_CertReclaimAbort", "10.0", _cSSM_TP_CertReclaimAbortErr)
	}
	return _cSSM_TP_CertReclaimAbort(TPHandle, KeyCacheHandle), nil
}

// CSSM_TP_CertReclaimAbort.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CertReclaimAbort
func CSSM_TP_CertReclaimAbort(TPHandle CSSM_TP_HANDLE, KeyCacheHandle CSSM_LONG_HANDLE) CSSM_RETURN {
	result, callErr := tryCSSM_TP_CertReclaimAbort(TPHandle, KeyCacheHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_CertReclaimKey func(TPHandle CSSM_TP_HANDLE, CertGroup unsafe.Pointer, CertIndex uint32, KeyCacheHandle CSSM_LONG_HANDLE, CSPHandle CSSM_CSP_HANDLE, CredAndAclEntry unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_CertReclaimKeyErr error

func tryCSSM_TP_CertReclaimKey(TPHandle CSSM_TP_HANDLE, CertGroup unsafe.Pointer, CertIndex uint32, KeyCacheHandle CSSM_LONG_HANDLE, CSPHandle CSSM_CSP_HANDLE, CredAndAclEntry unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_TP_CertReclaimKey == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_TP_CertReclaimKey", "10.0", _cSSM_TP_CertReclaimKeyErr)
	}
	return _cSSM_TP_CertReclaimKey(TPHandle, CertGroup, CertIndex, KeyCacheHandle, CSPHandle, CredAndAclEntry), nil
}

// CSSM_TP_CertReclaimKey.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CertReclaimKey
func CSSM_TP_CertReclaimKey(TPHandle CSSM_TP_HANDLE, CertGroup unsafe.Pointer, CertIndex uint32, KeyCacheHandle CSSM_LONG_HANDLE, CSPHandle CSSM_CSP_HANDLE, CredAndAclEntry unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_TP_CertReclaimKey(TPHandle, CertGroup, CertIndex, KeyCacheHandle, CSPHandle, CredAndAclEntry)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_CertRemoveFromCrlTemplate func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, OldCrlTemplate unsafe.Pointer, CertGroupToBeRemoved unsafe.Pointer, RevokerCertGroup unsafe.Pointer, RevokerVerifyContext unsafe.Pointer, RevokerVerifyResult unsafe.Pointer, NewCrlTemplate unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_CertRemoveFromCrlTemplateErr error

func tryCSSM_TP_CertRemoveFromCrlTemplate(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, OldCrlTemplate unsafe.Pointer, CertGroupToBeRemoved unsafe.Pointer, RevokerCertGroup unsafe.Pointer, RevokerVerifyContext unsafe.Pointer, RevokerVerifyResult unsafe.Pointer, NewCrlTemplate unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_TP_CertRemoveFromCrlTemplate == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_TP_CertRemoveFromCrlTemplate", "10.0", _cSSM_TP_CertRemoveFromCrlTemplateErr)
	}
	return _cSSM_TP_CertRemoveFromCrlTemplate(TPHandle, CLHandle, CSPHandle, OldCrlTemplate, CertGroupToBeRemoved, RevokerCertGroup, RevokerVerifyContext, RevokerVerifyResult, NewCrlTemplate), nil
}

// CSSM_TP_CertRemoveFromCrlTemplate.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CertRemoveFromCrlTemplate
func CSSM_TP_CertRemoveFromCrlTemplate(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, OldCrlTemplate unsafe.Pointer, CertGroupToBeRemoved unsafe.Pointer, RevokerCertGroup unsafe.Pointer, RevokerVerifyContext unsafe.Pointer, RevokerVerifyResult unsafe.Pointer, NewCrlTemplate unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_TP_CertRemoveFromCrlTemplate(TPHandle, CLHandle, CSPHandle, OldCrlTemplate, CertGroupToBeRemoved, RevokerCertGroup, RevokerVerifyContext, RevokerVerifyResult, NewCrlTemplate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_CertRevoke func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, OldCrlTemplate unsafe.Pointer, CertGroupToBeRevoked unsafe.Pointer, RevokerCertGroup unsafe.Pointer, RevokerVerifyContext unsafe.Pointer, RevokerVerifyResult unsafe.Pointer, Reason CSSM_TP_CERTCHANGE_REASON, NewCrlTemplate unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_CertRevokeErr error

func tryCSSM_TP_CertRevoke(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, OldCrlTemplate unsafe.Pointer, CertGroupToBeRevoked unsafe.Pointer, RevokerCertGroup unsafe.Pointer, RevokerVerifyContext unsafe.Pointer, RevokerVerifyResult unsafe.Pointer, Reason CSSM_TP_CERTCHANGE_REASON, NewCrlTemplate unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_TP_CertRevoke == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_TP_CertRevoke", "10.0", _cSSM_TP_CertRevokeErr)
	}
	return _cSSM_TP_CertRevoke(TPHandle, CLHandle, CSPHandle, OldCrlTemplate, CertGroupToBeRevoked, RevokerCertGroup, RevokerVerifyContext, RevokerVerifyResult, Reason, NewCrlTemplate), nil
}

// CSSM_TP_CertRevoke.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CertRevoke
func CSSM_TP_CertRevoke(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, OldCrlTemplate unsafe.Pointer, CertGroupToBeRevoked unsafe.Pointer, RevokerCertGroup unsafe.Pointer, RevokerVerifyContext unsafe.Pointer, RevokerVerifyResult unsafe.Pointer, Reason CSSM_TP_CERTCHANGE_REASON, NewCrlTemplate unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_TP_CertRevoke(TPHandle, CLHandle, CSPHandle, OldCrlTemplate, CertGroupToBeRevoked, RevokerCertGroup, RevokerVerifyContext, RevokerVerifyResult, Reason, NewCrlTemplate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_CertSign func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertTemplateToBeSigned unsafe.Pointer, SignerCertGroup unsafe.Pointer, SignerVerifyContext unsafe.Pointer, SignerVerifyResult unsafe.Pointer, SignedCert unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_CertSignErr error

func tryCSSM_TP_CertSign(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertTemplateToBeSigned unsafe.Pointer, SignerCertGroup unsafe.Pointer, SignerVerifyContext unsafe.Pointer, SignerVerifyResult unsafe.Pointer, SignedCert unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_TP_CertSign == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_TP_CertSign", "10.0", _cSSM_TP_CertSignErr)
	}
	return _cSSM_TP_CertSign(TPHandle, CLHandle, CCHandle, CertTemplateToBeSigned, SignerCertGroup, SignerVerifyContext, SignerVerifyResult, SignedCert), nil
}

// CSSM_TP_CertSign.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CertSign
func CSSM_TP_CertSign(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertTemplateToBeSigned unsafe.Pointer, SignerCertGroup unsafe.Pointer, SignerVerifyContext unsafe.Pointer, SignerVerifyResult unsafe.Pointer, SignedCert unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_TP_CertSign(TPHandle, CLHandle, CCHandle, CertTemplateToBeSigned, SignerCertGroup, SignerVerifyContext, SignerVerifyResult, SignedCert)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_ConfirmCredResult func(TPHandle CSSM_TP_HANDLE, ReferenceIdentifier unsafe.Pointer, CallerAuthCredentials unsafe.Pointer, Responses unsafe.Pointer, PreferredAuthority unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_ConfirmCredResultErr error

func tryCSSM_TP_ConfirmCredResult(TPHandle CSSM_TP_HANDLE, ReferenceIdentifier unsafe.Pointer, CallerAuthCredentials unsafe.Pointer, Responses unsafe.Pointer, PreferredAuthority unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_TP_ConfirmCredResult == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_TP_ConfirmCredResult", "10.0", _cSSM_TP_ConfirmCredResultErr)
	}
	return _cSSM_TP_ConfirmCredResult(TPHandle, ReferenceIdentifier, CallerAuthCredentials, Responses, PreferredAuthority), nil
}

// CSSM_TP_ConfirmCredResult.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_ConfirmCredResult
func CSSM_TP_ConfirmCredResult(TPHandle CSSM_TP_HANDLE, ReferenceIdentifier unsafe.Pointer, CallerAuthCredentials unsafe.Pointer, Responses unsafe.Pointer, PreferredAuthority unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_TP_ConfirmCredResult(TPHandle, ReferenceIdentifier, CallerAuthCredentials, Responses, PreferredAuthority)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_CrlCreateTemplate func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CrlFields unsafe.Pointer, NewCrlTemplate unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_CrlCreateTemplateErr error

func tryCSSM_TP_CrlCreateTemplate(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CrlFields unsafe.Pointer, NewCrlTemplate unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_TP_CrlCreateTemplate == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_TP_CrlCreateTemplate", "10.0", _cSSM_TP_CrlCreateTemplateErr)
	}
	return _cSSM_TP_CrlCreateTemplate(TPHandle, CLHandle, NumberOfFields, CrlFields, NewCrlTemplate), nil
}

// CSSM_TP_CrlCreateTemplate.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CrlCreateTemplate
func CSSM_TP_CrlCreateTemplate(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CrlFields unsafe.Pointer, NewCrlTemplate unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_TP_CrlCreateTemplate(TPHandle, CLHandle, NumberOfFields, CrlFields, NewCrlTemplate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_CrlSign func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CrlToBeSigned unsafe.Pointer, SignerCertGroup unsafe.Pointer, SignerVerifyContext unsafe.Pointer, SignerVerifyResult unsafe.Pointer, SignedCrl unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_CrlSignErr error

func tryCSSM_TP_CrlSign(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CrlToBeSigned unsafe.Pointer, SignerCertGroup unsafe.Pointer, SignerVerifyContext unsafe.Pointer, SignerVerifyResult unsafe.Pointer, SignedCrl unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_TP_CrlSign == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_TP_CrlSign", "10.0", _cSSM_TP_CrlSignErr)
	}
	return _cSSM_TP_CrlSign(TPHandle, CLHandle, CCHandle, CrlToBeSigned, SignerCertGroup, SignerVerifyContext, SignerVerifyResult, SignedCrl), nil
}

// CSSM_TP_CrlSign.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CrlSign
func CSSM_TP_CrlSign(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CrlToBeSigned unsafe.Pointer, SignerCertGroup unsafe.Pointer, SignerVerifyContext unsafe.Pointer, SignerVerifyResult unsafe.Pointer, SignedCrl unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_TP_CrlSign(TPHandle, CLHandle, CCHandle, CrlToBeSigned, SignerCertGroup, SignerVerifyContext, SignerVerifyResult, SignedCrl)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_CrlVerify func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, CrlToBeVerified unsafe.Pointer, SignerCertGroup unsafe.Pointer, VerifyContext unsafe.Pointer, RevokerVerifyResult unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_CrlVerifyErr error

func tryCSSM_TP_CrlVerify(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, CrlToBeVerified unsafe.Pointer, SignerCertGroup unsafe.Pointer, VerifyContext unsafe.Pointer, RevokerVerifyResult unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_TP_CrlVerify == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_TP_CrlVerify", "10.0", _cSSM_TP_CrlVerifyErr)
	}
	return _cSSM_TP_CrlVerify(TPHandle, CLHandle, CSPHandle, CrlToBeVerified, SignerCertGroup, VerifyContext, RevokerVerifyResult), nil
}

// CSSM_TP_CrlVerify.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CrlVerify
func CSSM_TP_CrlVerify(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, CrlToBeVerified unsafe.Pointer, SignerCertGroup unsafe.Pointer, VerifyContext unsafe.Pointer, RevokerVerifyResult unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_TP_CrlVerify(TPHandle, CLHandle, CSPHandle, CrlToBeVerified, SignerCertGroup, VerifyContext, RevokerVerifyResult)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_FormRequest func(TPHandle CSSM_TP_HANDLE, PreferredAuthority unsafe.Pointer, FormType CSSM_TP_FORM_TYPE, BlankForm unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_FormRequestErr error

func tryCSSM_TP_FormRequest(TPHandle CSSM_TP_HANDLE, PreferredAuthority unsafe.Pointer, FormType CSSM_TP_FORM_TYPE, BlankForm unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_TP_FormRequest == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_TP_FormRequest", "10.0", _cSSM_TP_FormRequestErr)
	}
	return _cSSM_TP_FormRequest(TPHandle, PreferredAuthority, FormType, BlankForm), nil
}

// CSSM_TP_FormRequest.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_FormRequest
func CSSM_TP_FormRequest(TPHandle CSSM_TP_HANDLE, PreferredAuthority unsafe.Pointer, FormType CSSM_TP_FORM_TYPE, BlankForm unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_TP_FormRequest(TPHandle, PreferredAuthority, FormType, BlankForm)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_FormSubmit func(TPHandle CSSM_TP_HANDLE, FormType CSSM_TP_FORM_TYPE, Form unsafe.Pointer, ClearanceAuthority unsafe.Pointer, RepresentedAuthority unsafe.Pointer, Credentials unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_FormSubmitErr error

func tryCSSM_TP_FormSubmit(TPHandle CSSM_TP_HANDLE, FormType CSSM_TP_FORM_TYPE, Form unsafe.Pointer, ClearanceAuthority unsafe.Pointer, RepresentedAuthority unsafe.Pointer, Credentials unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_TP_FormSubmit == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_TP_FormSubmit", "10.0", _cSSM_TP_FormSubmitErr)
	}
	return _cSSM_TP_FormSubmit(TPHandle, FormType, Form, ClearanceAuthority, RepresentedAuthority, Credentials), nil
}

// CSSM_TP_FormSubmit.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_FormSubmit
func CSSM_TP_FormSubmit(TPHandle CSSM_TP_HANDLE, FormType CSSM_TP_FORM_TYPE, Form unsafe.Pointer, ClearanceAuthority unsafe.Pointer, RepresentedAuthority unsafe.Pointer, Credentials unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_TP_FormSubmit(TPHandle, FormType, Form, ClearanceAuthority, RepresentedAuthority, Credentials)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_PassThrough func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, DBList unsafe.Pointer, PassThroughId uint32, InputParams unsafe.Pointer, OutputParams unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_PassThroughErr error

func tryCSSM_TP_PassThrough(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, DBList unsafe.Pointer, PassThroughId uint32, InputParams unsafe.Pointer, OutputParams unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_TP_PassThrough == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_TP_PassThrough", "10.0", _cSSM_TP_PassThroughErr)
	}
	return _cSSM_TP_PassThrough(TPHandle, CLHandle, CCHandle, DBList, PassThroughId, InputParams, OutputParams), nil
}

// CSSM_TP_PassThrough.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_PassThrough
func CSSM_TP_PassThrough(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, DBList unsafe.Pointer, PassThroughId uint32, InputParams unsafe.Pointer, OutputParams unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_TP_PassThrough(TPHandle, CLHandle, CCHandle, DBList, PassThroughId, InputParams, OutputParams)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_ReceiveConfirmation func(TPHandle CSSM_TP_HANDLE, ReferenceIdentifier unsafe.Pointer, Responses unsafe.Pointer, ElapsedTime *int32) CSSM_RETURN
var _cSSM_TP_ReceiveConfirmationErr error

func tryCSSM_TP_ReceiveConfirmation(TPHandle CSSM_TP_HANDLE, ReferenceIdentifier unsafe.Pointer, Responses unsafe.Pointer, ElapsedTime *int32) (CSSM_RETURN, error) {
	if _cSSM_TP_ReceiveConfirmation == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_TP_ReceiveConfirmation", "10.0", _cSSM_TP_ReceiveConfirmationErr)
	}
	return _cSSM_TP_ReceiveConfirmation(TPHandle, ReferenceIdentifier, Responses, ElapsedTime), nil
}

// CSSM_TP_ReceiveConfirmation.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_ReceiveConfirmation
func CSSM_TP_ReceiveConfirmation(TPHandle CSSM_TP_HANDLE, ReferenceIdentifier unsafe.Pointer, Responses unsafe.Pointer, ElapsedTime *int32) CSSM_RETURN {
	result, callErr := tryCSSM_TP_ReceiveConfirmation(TPHandle, ReferenceIdentifier, Responses, ElapsedTime)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_RetrieveCredResult func(TPHandle CSSM_TP_HANDLE, ReferenceIdentifier unsafe.Pointer, CallerAuthCredentials unsafe.Pointer, EstimatedTime *int32, ConfirmationRequired unsafe.Pointer, RetrieveOutput unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_RetrieveCredResultErr error

func tryCSSM_TP_RetrieveCredResult(TPHandle CSSM_TP_HANDLE, ReferenceIdentifier unsafe.Pointer, CallerAuthCredentials unsafe.Pointer, EstimatedTime *int32, ConfirmationRequired unsafe.Pointer, RetrieveOutput unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_TP_RetrieveCredResult == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_TP_RetrieveCredResult", "10.0", _cSSM_TP_RetrieveCredResultErr)
	}
	return _cSSM_TP_RetrieveCredResult(TPHandle, ReferenceIdentifier, CallerAuthCredentials, EstimatedTime, ConfirmationRequired, RetrieveOutput), nil
}

// CSSM_TP_RetrieveCredResult.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_RetrieveCredResult
func CSSM_TP_RetrieveCredResult(TPHandle CSSM_TP_HANDLE, ReferenceIdentifier unsafe.Pointer, CallerAuthCredentials unsafe.Pointer, EstimatedTime *int32, ConfirmationRequired unsafe.Pointer, RetrieveOutput unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_TP_RetrieveCredResult(TPHandle, ReferenceIdentifier, CallerAuthCredentials, EstimatedTime, ConfirmationRequired, RetrieveOutput)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_SubmitCredRequest func(TPHandle CSSM_TP_HANDLE, PreferredAuthority unsafe.Pointer, RequestType CSSM_TP_AUTHORITY_REQUEST_TYPE, RequestInput unsafe.Pointer, CallerAuthContext unsafe.Pointer, EstimatedTime *int32, ReferenceIdentifier unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_SubmitCredRequestErr error

func tryCSSM_TP_SubmitCredRequest(TPHandle CSSM_TP_HANDLE, PreferredAuthority unsafe.Pointer, RequestType CSSM_TP_AUTHORITY_REQUEST_TYPE, RequestInput unsafe.Pointer, CallerAuthContext unsafe.Pointer, EstimatedTime *int32, ReferenceIdentifier unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_TP_SubmitCredRequest == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_TP_SubmitCredRequest", "10.0", _cSSM_TP_SubmitCredRequestErr)
	}
	return _cSSM_TP_SubmitCredRequest(TPHandle, PreferredAuthority, RequestType, RequestInput, CallerAuthContext, EstimatedTime, ReferenceIdentifier), nil
}

// CSSM_TP_SubmitCredRequest.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_SubmitCredRequest
func CSSM_TP_SubmitCredRequest(TPHandle CSSM_TP_HANDLE, PreferredAuthority unsafe.Pointer, RequestType CSSM_TP_AUTHORITY_REQUEST_TYPE, RequestInput unsafe.Pointer, CallerAuthContext unsafe.Pointer, EstimatedTime *int32, ReferenceIdentifier unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_TP_SubmitCredRequest(TPHandle, PreferredAuthority, RequestType, RequestInput, CallerAuthContext, EstimatedTime, ReferenceIdentifier)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_TupleGroupToCertGroup func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, TupleGroup unsafe.Pointer, CertTemplates unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_TupleGroupToCertGroupErr error

func tryCSSM_TP_TupleGroupToCertGroup(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, TupleGroup unsafe.Pointer, CertTemplates unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_TP_TupleGroupToCertGroup == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_TP_TupleGroupToCertGroup", "10.0", _cSSM_TP_TupleGroupToCertGroupErr)
	}
	return _cSSM_TP_TupleGroupToCertGroup(TPHandle, CLHandle, TupleGroup, CertTemplates), nil
}

// CSSM_TP_TupleGroupToCertGroup.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_TupleGroupToCertGroup
func CSSM_TP_TupleGroupToCertGroup(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, TupleGroup unsafe.Pointer, CertTemplates unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_TP_TupleGroupToCertGroup(TPHandle, CLHandle, TupleGroup, CertTemplates)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_Terminate func() CSSM_RETURN
var _cSSM_TerminateErr error

func tryCSSM_Terminate() (CSSM_RETURN, error) {
	if _cSSM_Terminate == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_Terminate", "10.0", _cSSM_TerminateErr)
	}
	return _cSSM_Terminate(), nil
}

// CSSM_Terminate.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_Terminate
func CSSM_Terminate() CSSM_RETURN {
	result, callErr := tryCSSM_Terminate()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_Unintroduce func(ModuleID unsafe.Pointer) CSSM_RETURN
var _cSSM_UnintroduceErr error

func tryCSSM_Unintroduce(ModuleID unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_Unintroduce == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_Unintroduce", "10.0", _cSSM_UnintroduceErr)
	}
	return _cSSM_Unintroduce(ModuleID), nil
}

// CSSM_Unintroduce.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_Unintroduce
func CSSM_Unintroduce(ModuleID unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_Unintroduce(ModuleID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_UnwrapKey func(CCHandle CSSM_CC_HANDLE, PublicKey unsafe.Pointer, WrappedKey unsafe.Pointer, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, UnwrappedKey unsafe.Pointer, DescriptiveData unsafe.Pointer) CSSM_RETURN
var _cSSM_UnwrapKeyErr error

func tryCSSM_UnwrapKey(CCHandle CSSM_CC_HANDLE, PublicKey unsafe.Pointer, WrappedKey unsafe.Pointer, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, UnwrappedKey unsafe.Pointer, DescriptiveData unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_UnwrapKey == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_UnwrapKey", "10.0", _cSSM_UnwrapKeyErr)
	}
	return _cSSM_UnwrapKey(CCHandle, PublicKey, WrappedKey, KeyUsage, KeyAttr, KeyLabel, CredAndAclEntry, UnwrappedKey, DescriptiveData), nil
}

// CSSM_UnwrapKey.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_UnwrapKey
func CSSM_UnwrapKey(CCHandle CSSM_CC_HANDLE, PublicKey unsafe.Pointer, WrappedKey unsafe.Pointer, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, UnwrappedKey unsafe.Pointer, DescriptiveData unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_UnwrapKey(CCHandle, PublicKey, WrappedKey, KeyUsage, KeyAttr, KeyLabel, CredAndAclEntry, UnwrappedKey, DescriptiveData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_UnwrapKeyP func(CCHandle CSSM_CC_HANDLE, PublicKey unsafe.Pointer, WrappedKey unsafe.Pointer, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, UnwrappedKey unsafe.Pointer, DescriptiveData unsafe.Pointer, Privilege CSSM_PRIVILEGE) CSSM_RETURN
var _cSSM_UnwrapKeyPErr error

func tryCSSM_UnwrapKeyP(CCHandle CSSM_CC_HANDLE, PublicKey unsafe.Pointer, WrappedKey unsafe.Pointer, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, UnwrappedKey unsafe.Pointer, DescriptiveData unsafe.Pointer, Privilege CSSM_PRIVILEGE) (CSSM_RETURN, error) {
	if _cSSM_UnwrapKeyP == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_UnwrapKeyP", "10.0", _cSSM_UnwrapKeyPErr)
	}
	return _cSSM_UnwrapKeyP(CCHandle, PublicKey, WrappedKey, KeyUsage, KeyAttr, KeyLabel, CredAndAclEntry, UnwrappedKey, DescriptiveData, Privilege), nil
}

// CSSM_UnwrapKeyP.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_UnwrapKeyP
func CSSM_UnwrapKeyP(CCHandle CSSM_CC_HANDLE, PublicKey unsafe.Pointer, WrappedKey unsafe.Pointer, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, UnwrappedKey unsafe.Pointer, DescriptiveData unsafe.Pointer, Privilege CSSM_PRIVILEGE) CSSM_RETURN {
	result, callErr := tryCSSM_UnwrapKeyP(CCHandle, PublicKey, WrappedKey, KeyUsage, KeyAttr, KeyLabel, CredAndAclEntry, UnwrappedKey, DescriptiveData, Privilege)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_UpdateContextAttributes func(CCHandle CSSM_CC_HANDLE, NumberOfAttributes uint32, ContextAttributes unsafe.Pointer) CSSM_RETURN
var _cSSM_UpdateContextAttributesErr error

func tryCSSM_UpdateContextAttributes(CCHandle CSSM_CC_HANDLE, NumberOfAttributes uint32, ContextAttributes unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_UpdateContextAttributes == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_UpdateContextAttributes", "10.0", _cSSM_UpdateContextAttributesErr)
	}
	return _cSSM_UpdateContextAttributes(CCHandle, NumberOfAttributes, ContextAttributes), nil
}

// CSSM_UpdateContextAttributes.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_UpdateContextAttributes
func CSSM_UpdateContextAttributes(CCHandle CSSM_CC_HANDLE, NumberOfAttributes uint32, ContextAttributes unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_UpdateContextAttributes(CCHandle, NumberOfAttributes, ContextAttributes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_VerifyData func(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, DigestAlgorithm CSSM_ALGORITHMS, Signature unsafe.Pointer) CSSM_RETURN
var _cSSM_VerifyDataErr error

func tryCSSM_VerifyData(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, DigestAlgorithm CSSM_ALGORITHMS, Signature unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_VerifyData == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_VerifyData", "10.0", _cSSM_VerifyDataErr)
	}
	return _cSSM_VerifyData(CCHandle, DataBufs, DataBufCount, DigestAlgorithm, Signature), nil
}

// CSSM_VerifyData.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_VerifyData
func CSSM_VerifyData(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, DigestAlgorithm CSSM_ALGORITHMS, Signature unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_VerifyData(CCHandle, DataBufs, DataBufCount, DigestAlgorithm, Signature)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_VerifyDataFinal func(CCHandle CSSM_CC_HANDLE, Signature unsafe.Pointer) CSSM_RETURN
var _cSSM_VerifyDataFinalErr error

func tryCSSM_VerifyDataFinal(CCHandle CSSM_CC_HANDLE, Signature unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_VerifyDataFinal == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_VerifyDataFinal", "10.0", _cSSM_VerifyDataFinalErr)
	}
	return _cSSM_VerifyDataFinal(CCHandle, Signature), nil
}

// CSSM_VerifyDataFinal.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_VerifyDataFinal
func CSSM_VerifyDataFinal(CCHandle CSSM_CC_HANDLE, Signature unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_VerifyDataFinal(CCHandle, Signature)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_VerifyDataInit func(CCHandle CSSM_CC_HANDLE) CSSM_RETURN
var _cSSM_VerifyDataInitErr error

func tryCSSM_VerifyDataInit(CCHandle CSSM_CC_HANDLE) (CSSM_RETURN, error) {
	if _cSSM_VerifyDataInit == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_VerifyDataInit", "10.0", _cSSM_VerifyDataInitErr)
	}
	return _cSSM_VerifyDataInit(CCHandle), nil
}

// CSSM_VerifyDataInit.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_VerifyDataInit
func CSSM_VerifyDataInit(CCHandle CSSM_CC_HANDLE) CSSM_RETURN {
	result, callErr := tryCSSM_VerifyDataInit(CCHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_VerifyDataUpdate func(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) CSSM_RETURN
var _cSSM_VerifyDataUpdateErr error

func tryCSSM_VerifyDataUpdate(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) (CSSM_RETURN, error) {
	if _cSSM_VerifyDataUpdate == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_VerifyDataUpdate", "10.0", _cSSM_VerifyDataUpdateErr)
	}
	return _cSSM_VerifyDataUpdate(CCHandle, DataBufs, DataBufCount), nil
}

// CSSM_VerifyDataUpdate.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_VerifyDataUpdate
func CSSM_VerifyDataUpdate(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) CSSM_RETURN {
	result, callErr := tryCSSM_VerifyDataUpdate(CCHandle, DataBufs, DataBufCount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_VerifyDevice func(CSPHandle CSSM_CSP_HANDLE, DeviceCert unsafe.Pointer) CSSM_RETURN
var _cSSM_VerifyDeviceErr error

func tryCSSM_VerifyDevice(CSPHandle CSSM_CSP_HANDLE, DeviceCert unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_VerifyDevice == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_VerifyDevice", "10.0", _cSSM_VerifyDeviceErr)
	}
	return _cSSM_VerifyDevice(CSPHandle, DeviceCert), nil
}

// CSSM_VerifyDevice.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_VerifyDevice
func CSSM_VerifyDevice(CSPHandle CSSM_CSP_HANDLE, DeviceCert unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_VerifyDevice(CSPHandle, DeviceCert)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_VerifyMac func(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, Mac unsafe.Pointer) CSSM_RETURN
var _cSSM_VerifyMacErr error

func tryCSSM_VerifyMac(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, Mac unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_VerifyMac == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_VerifyMac", "10.0", _cSSM_VerifyMacErr)
	}
	return _cSSM_VerifyMac(CCHandle, DataBufs, DataBufCount, Mac), nil
}

// CSSM_VerifyMac.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_VerifyMac
func CSSM_VerifyMac(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, Mac unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_VerifyMac(CCHandle, DataBufs, DataBufCount, Mac)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_VerifyMacFinal func(CCHandle CSSM_CC_HANDLE, Mac unsafe.Pointer) CSSM_RETURN
var _cSSM_VerifyMacFinalErr error

func tryCSSM_VerifyMacFinal(CCHandle CSSM_CC_HANDLE, Mac unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_VerifyMacFinal == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_VerifyMacFinal", "10.0", _cSSM_VerifyMacFinalErr)
	}
	return _cSSM_VerifyMacFinal(CCHandle, Mac), nil
}

// CSSM_VerifyMacFinal.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_VerifyMacFinal
func CSSM_VerifyMacFinal(CCHandle CSSM_CC_HANDLE, Mac unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_VerifyMacFinal(CCHandle, Mac)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_VerifyMacInit func(CCHandle CSSM_CC_HANDLE) CSSM_RETURN
var _cSSM_VerifyMacInitErr error

func tryCSSM_VerifyMacInit(CCHandle CSSM_CC_HANDLE) (CSSM_RETURN, error) {
	if _cSSM_VerifyMacInit == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_VerifyMacInit", "10.0", _cSSM_VerifyMacInitErr)
	}
	return _cSSM_VerifyMacInit(CCHandle), nil
}

// CSSM_VerifyMacInit.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_VerifyMacInit
func CSSM_VerifyMacInit(CCHandle CSSM_CC_HANDLE) CSSM_RETURN {
	result, callErr := tryCSSM_VerifyMacInit(CCHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_VerifyMacUpdate func(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) CSSM_RETURN
var _cSSM_VerifyMacUpdateErr error

func tryCSSM_VerifyMacUpdate(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) (CSSM_RETURN, error) {
	if _cSSM_VerifyMacUpdate == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_VerifyMacUpdate", "10.0", _cSSM_VerifyMacUpdateErr)
	}
	return _cSSM_VerifyMacUpdate(CCHandle, DataBufs, DataBufCount), nil
}

// CSSM_VerifyMacUpdate.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_VerifyMacUpdate
func CSSM_VerifyMacUpdate(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) CSSM_RETURN {
	result, callErr := tryCSSM_VerifyMacUpdate(CCHandle, DataBufs, DataBufCount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_WrapKey func(CCHandle CSSM_CC_HANDLE, AccessCred unsafe.Pointer, Key unsafe.Pointer, DescriptiveData unsafe.Pointer, WrappedKey unsafe.Pointer) CSSM_RETURN
var _cSSM_WrapKeyErr error

func tryCSSM_WrapKey(CCHandle CSSM_CC_HANDLE, AccessCred unsafe.Pointer, Key unsafe.Pointer, DescriptiveData unsafe.Pointer, WrappedKey unsafe.Pointer) (CSSM_RETURN, error) {
	if _cSSM_WrapKey == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_WrapKey", "10.0", _cSSM_WrapKeyErr)
	}
	return _cSSM_WrapKey(CCHandle, AccessCred, Key, DescriptiveData, WrappedKey), nil
}

// CSSM_WrapKey.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_WrapKey
func CSSM_WrapKey(CCHandle CSSM_CC_HANDLE, AccessCred unsafe.Pointer, Key unsafe.Pointer, DescriptiveData unsafe.Pointer, WrappedKey unsafe.Pointer) CSSM_RETURN {
	result, callErr := tryCSSM_WrapKey(CCHandle, AccessCred, Key, DescriptiveData, WrappedKey)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_WrapKeyP func(CCHandle CSSM_CC_HANDLE, AccessCred unsafe.Pointer, Key unsafe.Pointer, DescriptiveData unsafe.Pointer, WrappedKey unsafe.Pointer, Privilege CSSM_PRIVILEGE) CSSM_RETURN
var _cSSM_WrapKeyPErr error

func tryCSSM_WrapKeyP(CCHandle CSSM_CC_HANDLE, AccessCred unsafe.Pointer, Key unsafe.Pointer, DescriptiveData unsafe.Pointer, WrappedKey unsafe.Pointer, Privilege CSSM_PRIVILEGE) (CSSM_RETURN, error) {
	if _cSSM_WrapKeyP == nil {
		return *new(CSSM_RETURN), symbolCallError("CSSM_WrapKeyP", "10.0", _cSSM_WrapKeyPErr)
	}
	return _cSSM_WrapKeyP(CCHandle, AccessCred, Key, DescriptiveData, WrappedKey, Privilege), nil
}

// CSSM_WrapKeyP.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_WrapKeyP
func CSSM_WrapKeyP(CCHandle CSSM_CC_HANDLE, AccessCred unsafe.Pointer, Key unsafe.Pointer, DescriptiveData unsafe.Pointer, WrappedKey unsafe.Pointer, Privilege CSSM_PRIVILEGE) CSSM_RETURN {
	result, callErr := tryCSSM_WrapKeyP(CCHandle, AccessCred, Key, DescriptiveData, WrappedKey, Privilege)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mDS_Initialize func(pCallerGuid unsafe.Pointer, pMemoryFunctions unsafe.Pointer, pDlFunctions unsafe.Pointer, hMds *MDS_HANDLE) CSSM_RETURN
var _mDS_InitializeErr error

func tryMDS_Initialize(pCallerGuid unsafe.Pointer, pMemoryFunctions unsafe.Pointer, pDlFunctions unsafe.Pointer, hMds *MDS_HANDLE) (CSSM_RETURN, error) {
	if _mDS_Initialize == nil {
		return *new(CSSM_RETURN), symbolCallError("MDS_Initialize", "10.0", _mDS_InitializeErr)
	}
	return _mDS_Initialize(pCallerGuid, pMemoryFunctions, pDlFunctions, hMds), nil
}

// MDS_Initialize.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/MDS_Initialize
func MDS_Initialize(pCallerGuid unsafe.Pointer, pMemoryFunctions unsafe.Pointer, pDlFunctions unsafe.Pointer, hMds *MDS_HANDLE) CSSM_RETURN {
	result, callErr := tryMDS_Initialize(pCallerGuid, pMemoryFunctions, pDlFunctions, hMds)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mDS_Install func(MdsHandle MDS_HANDLE) CSSM_RETURN
var _mDS_InstallErr error

func tryMDS_Install(MdsHandle MDS_HANDLE) (CSSM_RETURN, error) {
	if _mDS_Install == nil {
		return *new(CSSM_RETURN), symbolCallError("MDS_Install", "10.0", _mDS_InstallErr)
	}
	return _mDS_Install(MdsHandle), nil
}

// MDS_Install.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/MDS_Install
func MDS_Install(MdsHandle MDS_HANDLE) CSSM_RETURN {
	result, callErr := tryMDS_Install(MdsHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mDS_Terminate func(MdsHandle MDS_HANDLE) CSSM_RETURN
var _mDS_TerminateErr error

func tryMDS_Terminate(MdsHandle MDS_HANDLE) (CSSM_RETURN, error) {
	if _mDS_Terminate == nil {
		return *new(CSSM_RETURN), symbolCallError("MDS_Terminate", "10.0", _mDS_TerminateErr)
	}
	return _mDS_Terminate(MdsHandle), nil
}

// MDS_Terminate.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/MDS_Terminate
func MDS_Terminate(MdsHandle MDS_HANDLE) CSSM_RETURN {
	result, callErr := tryMDS_Terminate(MdsHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mDS_Uninstall func(MdsHandle MDS_HANDLE) CSSM_RETURN
var _mDS_UninstallErr error

func tryMDS_Uninstall(MdsHandle MDS_HANDLE) (CSSM_RETURN, error) {
	if _mDS_Uninstall == nil {
		return *new(CSSM_RETURN), symbolCallError("MDS_Uninstall", "10.0", _mDS_UninstallErr)
	}
	return _mDS_Uninstall(MdsHandle), nil
}

// MDS_Uninstall.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/MDS_Uninstall
func MDS_Uninstall(MdsHandle MDS_HANDLE) CSSM_RETURN {
	result, callErr := tryMDS_Uninstall(MdsHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secAccessControlCreateWithFlags func(allocator corefoundation.CFAllocatorRef, protection corefoundation.CFTypeRef, flags SecAccessControlCreateFlags, err *corefoundation.CFErrorRef) SecAccessControlRef
var _secAccessControlCreateWithFlagsErr error

func trySecAccessControlCreateWithFlags(allocator corefoundation.CFAllocatorRef, protection corefoundation.CFTypeRef, flags SecAccessControlCreateFlags, err *corefoundation.CFErrorRef) (SecAccessControlRef, error) {
	if _secAccessControlCreateWithFlags == nil {
		return 0, symbolCallError("SecAccessControlCreateWithFlags", "10.10", _secAccessControlCreateWithFlagsErr)
	}
	return _secAccessControlCreateWithFlags(allocator, protection, flags, err), nil
}

// SecAccessControlCreateWithFlags creates a new access control object with the specified protection type and flags.
//
// See: https://developer.apple.com/documentation/Security/SecAccessControlCreateWithFlags(_:_:_:_:)
func SecAccessControlCreateWithFlags(allocator corefoundation.CFAllocatorRef, protection corefoundation.CFTypeRef, flags SecAccessControlCreateFlags, err *corefoundation.CFErrorRef) SecAccessControlRef {
	result, callErr := trySecAccessControlCreateWithFlags(allocator, protection, flags, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secAccessControlGetTypeID func() uint
var _secAccessControlGetTypeIDErr error

func trySecAccessControlGetTypeID() (uint, error) {
	if _secAccessControlGetTypeID == nil {
		return 0, symbolCallError("SecAccessControlGetTypeID", "10.10", _secAccessControlGetTypeIDErr)
	}
	return _secAccessControlGetTypeID(), nil
}

// SecAccessControlGetTypeID returns the unique identifier of the opaque type to which a keychain item access control object belongs.
//
// See: https://developer.apple.com/documentation/Security/SecAccessControlGetTypeID()
func SecAccessControlGetTypeID() uint {
	result, callErr := trySecAccessControlGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secAddSharedWebCredential func(fqdn corefoundation.CFStringRef, account corefoundation.CFStringRef, password corefoundation.CFStringRef)
var _secAddSharedWebCredentialErr error

func trySecAddSharedWebCredential(fqdn corefoundation.CFStringRef, account corefoundation.CFStringRef, password corefoundation.CFStringRef) error {
	if _secAddSharedWebCredential == nil {
		return symbolCallError("SecAddSharedWebCredential", "11.0", _secAddSharedWebCredentialErr)
	}
	_secAddSharedWebCredential(fqdn, account, password)
	return nil
}

// SecAddSharedWebCredential asynchronously stores (or updates) a shared password for a website.
//
// Deprecated: Deprecated since macOS 26.2. Use ASCredentialDataManager.save(password:for:title:anchor:) (AuthenticationServices framework)
//
// See: https://developer.apple.com/documentation/Security/SecAddSharedWebCredential(_:_:_:_:)
func SecAddSharedWebCredential(fqdn corefoundation.CFStringRef, account corefoundation.CFStringRef, password corefoundation.CFStringRef) {
	if callErr := trySecAddSharedWebCredential(fqdn, account, password); callErr != nil {
		panic(callErr)
	}
}

var _secAsn1AllocCopy func(coder unsafe.Pointer, src unsafe.Pointer, len_ uintptr, dest unsafe.Pointer) int32
var _secAsn1AllocCopyErr error

func trySecAsn1AllocCopy(coder unsafe.Pointer, src unsafe.Pointer, len_ uintptr, dest unsafe.Pointer) (int32, error) {
	if _secAsn1AllocCopy == nil {
		return 0, symbolCallError("SecAsn1AllocCopy", "10.0", _secAsn1AllocCopyErr)
	}
	return _secAsn1AllocCopy(coder, src, len_, dest), nil
}

// SecAsn1AllocCopy allocates memory for an item’s data field in the coder object’s memory pool and copies in a block of data.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecAsn1AllocCopy
func SecAsn1AllocCopy(coder unsafe.Pointer, src unsafe.Pointer, len_ uintptr, dest unsafe.Pointer) int32 {
	result, callErr := trySecAsn1AllocCopy(coder, src, len_, dest)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secAsn1AllocCopyItem func(coder unsafe.Pointer, src unsafe.Pointer, dest unsafe.Pointer) int32
var _secAsn1AllocCopyItemErr error

func trySecAsn1AllocCopyItem(coder unsafe.Pointer, src unsafe.Pointer, dest unsafe.Pointer) (int32, error) {
	if _secAsn1AllocCopyItem == nil {
		return 0, symbolCallError("SecAsn1AllocCopyItem", "10.0", _secAsn1AllocCopyItemErr)
	}
	return _secAsn1AllocCopyItem(coder, src, dest), nil
}

// SecAsn1AllocCopyItem allocates memory for an item’s data field in the coder object’s memory pool and copies in a block of data from another item.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecAsn1AllocCopyItem
func SecAsn1AllocCopyItem(coder unsafe.Pointer, src unsafe.Pointer, dest unsafe.Pointer) int32 {
	result, callErr := trySecAsn1AllocCopyItem(coder, src, dest)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secAsn1AllocItem func(coder unsafe.Pointer, item unsafe.Pointer, len_ uintptr) int32
var _secAsn1AllocItemErr error

func trySecAsn1AllocItem(coder unsafe.Pointer, item unsafe.Pointer, len_ uintptr) (int32, error) {
	if _secAsn1AllocItem == nil {
		return 0, symbolCallError("SecAsn1AllocItem", "10.0", _secAsn1AllocItemErr)
	}
	return _secAsn1AllocItem(coder, item, len_), nil
}

// SecAsn1AllocItem allocates memory for an item’s data field in the coder object’s memory pool.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecAsn1AllocItem
func SecAsn1AllocItem(coder unsafe.Pointer, item unsafe.Pointer, len_ uintptr) int32 {
	result, callErr := trySecAsn1AllocItem(coder, item, len_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secAsn1CoderCreate func(coder unsafe.Pointer) int32
var _secAsn1CoderCreateErr error

func trySecAsn1CoderCreate(coder unsafe.Pointer) (int32, error) {
	if _secAsn1CoderCreate == nil {
		return 0, symbolCallError("SecAsn1CoderCreate", "10.0", _secAsn1CoderCreateErr)
	}
	return _secAsn1CoderCreate(coder), nil
}

// SecAsn1CoderCreate creates an ASN.1 coder object.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecAsn1CoderCreate
func SecAsn1CoderCreate(coder unsafe.Pointer) int32 {
	result, callErr := trySecAsn1CoderCreate(coder)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secAsn1CoderRelease func(coder unsafe.Pointer) int32
var _secAsn1CoderReleaseErr error

func trySecAsn1CoderRelease(coder unsafe.Pointer) (int32, error) {
	if _secAsn1CoderRelease == nil {
		return 0, symbolCallError("SecAsn1CoderRelease", "10.0", _secAsn1CoderReleaseErr)
	}
	return _secAsn1CoderRelease(coder), nil
}

// SecAsn1CoderRelease destroys an ASN.1 coder object and releases all of its memory.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecAsn1CoderRelease
func SecAsn1CoderRelease(coder unsafe.Pointer) int32 {
	result, callErr := trySecAsn1CoderRelease(coder)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secAsn1Decode func(coder unsafe.Pointer, src unsafe.Pointer, len_ uintptr, templates unsafe.Pointer, dest unsafe.Pointer) int32
var _secAsn1DecodeErr error

func trySecAsn1Decode(coder unsafe.Pointer, src unsafe.Pointer, len_ uintptr, templates unsafe.Pointer, dest unsafe.Pointer) (int32, error) {
	if _secAsn1Decode == nil {
		return 0, symbolCallError("SecAsn1Decode", "10.0", _secAsn1DecodeErr)
	}
	return _secAsn1Decode(coder, src, len_, templates, dest), nil
}

// SecAsn1Decode decodes untyped DER data.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecAsn1Decode
func SecAsn1Decode(coder unsafe.Pointer, src unsafe.Pointer, len_ uintptr, templates unsafe.Pointer, dest unsafe.Pointer) int32 {
	result, callErr := trySecAsn1Decode(coder, src, len_, templates, dest)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secAsn1DecodeData func(coder unsafe.Pointer, src unsafe.Pointer, templ unsafe.Pointer, dest unsafe.Pointer) int32
var _secAsn1DecodeDataErr error

func trySecAsn1DecodeData(coder unsafe.Pointer, src unsafe.Pointer, templ unsafe.Pointer, dest unsafe.Pointer) (int32, error) {
	if _secAsn1DecodeData == nil {
		return 0, symbolCallError("SecAsn1DecodeData", "10.0", _secAsn1DecodeDataErr)
	}
	return _secAsn1DecodeData(coder, src, templ, dest), nil
}

// SecAsn1DecodeData decodes an ASN.1 item in DER format.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecAsn1DecodeData
func SecAsn1DecodeData(coder unsafe.Pointer, src unsafe.Pointer, templ unsafe.Pointer, dest unsafe.Pointer) int32 {
	result, callErr := trySecAsn1DecodeData(coder, src, templ, dest)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secAsn1EncodeItem func(coder unsafe.Pointer, src unsafe.Pointer, templates unsafe.Pointer, dest unsafe.Pointer) int32
var _secAsn1EncodeItemErr error

func trySecAsn1EncodeItem(coder unsafe.Pointer, src unsafe.Pointer, templates unsafe.Pointer, dest unsafe.Pointer) (int32, error) {
	if _secAsn1EncodeItem == nil {
		return 0, symbolCallError("SecAsn1EncodeItem", "10.0", _secAsn1EncodeItemErr)
	}
	return _secAsn1EncodeItem(coder, src, templates, dest), nil
}

// SecAsn1EncodeItem encodes data in DER format.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecAsn1EncodeItem
func SecAsn1EncodeItem(coder unsafe.Pointer, src unsafe.Pointer, templates unsafe.Pointer, dest unsafe.Pointer) int32 {
	result, callErr := trySecAsn1EncodeItem(coder, src, templates, dest)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secAsn1Malloc func(coder unsafe.Pointer, len_ uintptr) unsafe.Pointer
var _secAsn1MallocErr error

func trySecAsn1Malloc(coder unsafe.Pointer, len_ uintptr) (unsafe.Pointer, error) {
	if _secAsn1Malloc == nil {
		return nil, symbolCallError("SecAsn1Malloc", "10.0", _secAsn1MallocErr)
	}
	return _secAsn1Malloc(coder, len_), nil
}

// SecAsn1Malloc allocates memory in the coder object’s memory pool.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecAsn1Malloc
func SecAsn1Malloc(coder unsafe.Pointer, len_ uintptr) unsafe.Pointer {
	result, callErr := trySecAsn1Malloc(coder, len_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secAsn1OidCompare func(oid1 unsafe.Pointer, oid2 unsafe.Pointer) bool
var _secAsn1OidCompareErr error

func trySecAsn1OidCompare(oid1 unsafe.Pointer, oid2 unsafe.Pointer) (bool, error) {
	if _secAsn1OidCompare == nil {
		return false, symbolCallError("SecAsn1OidCompare", "10.0", _secAsn1OidCompareErr)
	}
	return _secAsn1OidCompare(oid1, oid2), nil
}

// SecAsn1OidCompare compares two decoded object identifiers.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecAsn1OidCompare
func SecAsn1OidCompare(oid1 unsafe.Pointer, oid2 unsafe.Pointer) bool {
	result, callErr := trySecAsn1OidCompare(oid1, oid2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateAddToKeychain func(certificate SecCertificateRef, keychain SecKeychainRef) int32
var _secCertificateAddToKeychainErr error

func trySecCertificateAddToKeychain(certificate SecCertificateRef, keychain SecKeychainRef) (int32, error) {
	if _secCertificateAddToKeychain == nil {
		return 0, symbolCallError("SecCertificateAddToKeychain", "10.3", _secCertificateAddToKeychainErr)
	}
	return _secCertificateAddToKeychain(certificate, keychain), nil
}

// SecCertificateAddToKeychain adds a certificate to a keychain.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateAddToKeychain(_:_:)
func SecCertificateAddToKeychain(certificate SecCertificateRef, keychain SecKeychainRef) int32 {
	result, callErr := trySecCertificateAddToKeychain(certificate, keychain)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateCopyCommonName func(certificate SecCertificateRef, commonName *corefoundation.CFStringRef) int32
var _secCertificateCopyCommonNameErr error

func trySecCertificateCopyCommonName(certificate SecCertificateRef, commonName *corefoundation.CFStringRef) (int32, error) {
	if _secCertificateCopyCommonName == nil {
		return 0, symbolCallError("SecCertificateCopyCommonName", "10.5", _secCertificateCopyCommonNameErr)
	}
	return _secCertificateCopyCommonName(certificate, commonName), nil
}

// SecCertificateCopyCommonName retrieves the common name of the subject of a certificate.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopyCommonName(_:_:)
func SecCertificateCopyCommonName(certificate SecCertificateRef, commonName *corefoundation.CFStringRef) int32 {
	result, callErr := trySecCertificateCopyCommonName(certificate, commonName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateCopyData func(certificate SecCertificateRef) corefoundation.CFDataRef
var _secCertificateCopyDataErr error

func trySecCertificateCopyData(certificate SecCertificateRef) (corefoundation.CFDataRef, error) {
	if _secCertificateCopyData == nil {
		return 0, symbolCallError("SecCertificateCopyData", "10.6", _secCertificateCopyDataErr)
	}
	return _secCertificateCopyData(certificate), nil
}

// SecCertificateCopyData returns a DER representation of a certificate given a certificate object.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopyData(_:)
func SecCertificateCopyData(certificate SecCertificateRef) corefoundation.CFDataRef {
	result, callErr := trySecCertificateCopyData(certificate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateCopyEmailAddresses func(certificate SecCertificateRef, emailAddresses *corefoundation.CFArrayRef) int32
var _secCertificateCopyEmailAddressesErr error

func trySecCertificateCopyEmailAddresses(certificate SecCertificateRef, emailAddresses *corefoundation.CFArrayRef) (int32, error) {
	if _secCertificateCopyEmailAddresses == nil {
		return 0, symbolCallError("SecCertificateCopyEmailAddresses", "10.5", _secCertificateCopyEmailAddressesErr)
	}
	return _secCertificateCopyEmailAddresses(certificate, emailAddresses), nil
}

// SecCertificateCopyEmailAddresses retrieves the email addresses for the subject of a certificate.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopyEmailAddresses(_:_:)
func SecCertificateCopyEmailAddresses(certificate SecCertificateRef, emailAddresses *corefoundation.CFArrayRef) int32 {
	result, callErr := trySecCertificateCopyEmailAddresses(certificate, emailAddresses)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateCopyKey func(certificate SecCertificateRef) SecKeyRef
var _secCertificateCopyKeyErr error

func trySecCertificateCopyKey(certificate SecCertificateRef) (SecKeyRef, error) {
	if _secCertificateCopyKey == nil {
		return 0, symbolCallError("SecCertificateCopyKey", "10.14", _secCertificateCopyKeyErr)
	}
	return _secCertificateCopyKey(certificate), nil
}

// SecCertificateCopyKey retrieves the public key for a given certificate.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopyKey(_:)
func SecCertificateCopyKey(certificate SecCertificateRef) SecKeyRef {
	result, callErr := trySecCertificateCopyKey(certificate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateCopyLongDescription func(alloc corefoundation.CFAllocatorRef, certificate SecCertificateRef, err *corefoundation.CFErrorRef) corefoundation.CFStringRef
var _secCertificateCopyLongDescriptionErr error

func trySecCertificateCopyLongDescription(alloc corefoundation.CFAllocatorRef, certificate SecCertificateRef, err *corefoundation.CFErrorRef) (corefoundation.CFStringRef, error) {
	if _secCertificateCopyLongDescription == nil {
		return 0, symbolCallError("SecCertificateCopyLongDescription", "10.7", _secCertificateCopyLongDescriptionErr)
	}
	return _secCertificateCopyLongDescription(alloc, certificate, err), nil
}

// SecCertificateCopyLongDescription returns a copy of the long description of a certificate.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopyLongDescription(_:_:_:)
func SecCertificateCopyLongDescription(alloc corefoundation.CFAllocatorRef, certificate SecCertificateRef, err *corefoundation.CFErrorRef) corefoundation.CFStringRef {
	result, callErr := trySecCertificateCopyLongDescription(alloc, certificate, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateCopyNormalizedIssuerSequence func(certificate SecCertificateRef) corefoundation.CFDataRef
var _secCertificateCopyNormalizedIssuerSequenceErr error

func trySecCertificateCopyNormalizedIssuerSequence(certificate SecCertificateRef) (corefoundation.CFDataRef, error) {
	if _secCertificateCopyNormalizedIssuerSequence == nil {
		return 0, symbolCallError("SecCertificateCopyNormalizedIssuerSequence", "10.12.4", _secCertificateCopyNormalizedIssuerSequenceErr)
	}
	return _secCertificateCopyNormalizedIssuerSequence(certificate), nil
}

// SecCertificateCopyNormalizedIssuerSequence retrieves the normalized issuer sequence from a certificate.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopyNormalizedIssuerSequence(_:)
func SecCertificateCopyNormalizedIssuerSequence(certificate SecCertificateRef) corefoundation.CFDataRef {
	result, callErr := trySecCertificateCopyNormalizedIssuerSequence(certificate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateCopyNormalizedSubjectSequence func(certificate SecCertificateRef) corefoundation.CFDataRef
var _secCertificateCopyNormalizedSubjectSequenceErr error

func trySecCertificateCopyNormalizedSubjectSequence(certificate SecCertificateRef) (corefoundation.CFDataRef, error) {
	if _secCertificateCopyNormalizedSubjectSequence == nil {
		return 0, symbolCallError("SecCertificateCopyNormalizedSubjectSequence", "10.12.4", _secCertificateCopyNormalizedSubjectSequenceErr)
	}
	return _secCertificateCopyNormalizedSubjectSequence(certificate), nil
}

// SecCertificateCopyNormalizedSubjectSequence retrieves the normalized subject sequence from a certificate.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopyNormalizedSubjectSequence(_:)
func SecCertificateCopyNormalizedSubjectSequence(certificate SecCertificateRef) corefoundation.CFDataRef {
	result, callErr := trySecCertificateCopyNormalizedSubjectSequence(certificate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateCopyNotValidAfterDate func(certificate SecCertificateRef) corefoundation.CFDateRef
var _secCertificateCopyNotValidAfterDateErr error

func trySecCertificateCopyNotValidAfterDate(certificate SecCertificateRef) (corefoundation.CFDateRef, error) {
	if _secCertificateCopyNotValidAfterDate == nil {
		return 0, symbolCallError("SecCertificateCopyNotValidAfterDate", "15.0", _secCertificateCopyNotValidAfterDateErr)
	}
	return _secCertificateCopyNotValidAfterDate(certificate), nil
}

// SecCertificateCopyNotValidAfterDate.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopyNotValidAfterDate(_:)
func SecCertificateCopyNotValidAfterDate(certificate SecCertificateRef) corefoundation.CFDateRef {
	result, callErr := trySecCertificateCopyNotValidAfterDate(certificate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateCopyNotValidBeforeDate func(certificate SecCertificateRef) corefoundation.CFDateRef
var _secCertificateCopyNotValidBeforeDateErr error

func trySecCertificateCopyNotValidBeforeDate(certificate SecCertificateRef) (corefoundation.CFDateRef, error) {
	if _secCertificateCopyNotValidBeforeDate == nil {
		return 0, symbolCallError("SecCertificateCopyNotValidBeforeDate", "15.0", _secCertificateCopyNotValidBeforeDateErr)
	}
	return _secCertificateCopyNotValidBeforeDate(certificate), nil
}

// SecCertificateCopyNotValidBeforeDate.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopyNotValidBeforeDate(_:)
func SecCertificateCopyNotValidBeforeDate(certificate SecCertificateRef) corefoundation.CFDateRef {
	result, callErr := trySecCertificateCopyNotValidBeforeDate(certificate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateCopyPreference func(name corefoundation.CFStringRef, keyUsage uint32, certificate *SecCertificateRef) int32
var _secCertificateCopyPreferenceErr error

func trySecCertificateCopyPreference(name corefoundation.CFStringRef, keyUsage uint32, certificate *SecCertificateRef) (int32, error) {
	if _secCertificateCopyPreference == nil {
		return 0, symbolCallError("SecCertificateCopyPreference", "10.0", _secCertificateCopyPreferenceErr)
	}
	return _secCertificateCopyPreference(name, keyUsage, certificate), nil
}

// SecCertificateCopyPreference retrieves the preferred certificate for the specified name and key use.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopyPreference
func SecCertificateCopyPreference(name corefoundation.CFStringRef, keyUsage uint32, certificate *SecCertificateRef) int32 {
	result, callErr := trySecCertificateCopyPreference(name, keyUsage, certificate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateCopyPreferred func(name corefoundation.CFStringRef, keyUsage corefoundation.CFArrayRef) SecCertificateRef
var _secCertificateCopyPreferredErr error

func trySecCertificateCopyPreferred(name corefoundation.CFStringRef, keyUsage corefoundation.CFArrayRef) (SecCertificateRef, error) {
	if _secCertificateCopyPreferred == nil {
		return 0, symbolCallError("SecCertificateCopyPreferred", "10.7", _secCertificateCopyPreferredErr)
	}
	return _secCertificateCopyPreferred(name, keyUsage), nil
}

// SecCertificateCopyPreferred returns the preferred certificate for the specified name and key usage.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopyPreferred(_:_:)
func SecCertificateCopyPreferred(name corefoundation.CFStringRef, keyUsage corefoundation.CFArrayRef) SecCertificateRef {
	result, callErr := trySecCertificateCopyPreferred(name, keyUsage)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateCopySerialNumberData func(certificate SecCertificateRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef
var _secCertificateCopySerialNumberDataErr error

func trySecCertificateCopySerialNumberData(certificate SecCertificateRef, err *corefoundation.CFErrorRef) (corefoundation.CFDataRef, error) {
	if _secCertificateCopySerialNumberData == nil {
		return 0, symbolCallError("SecCertificateCopySerialNumberData", "10.13", _secCertificateCopySerialNumberDataErr)
	}
	return _secCertificateCopySerialNumberData(certificate, err), nil
}

// SecCertificateCopySerialNumberData returns the certificate’s serial number.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopySerialNumberData(_:_:)
func SecCertificateCopySerialNumberData(certificate SecCertificateRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef {
	result, callErr := trySecCertificateCopySerialNumberData(certificate, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateCopyShortDescription func(alloc corefoundation.CFAllocatorRef, certificate SecCertificateRef, err *corefoundation.CFErrorRef) corefoundation.CFStringRef
var _secCertificateCopyShortDescriptionErr error

func trySecCertificateCopyShortDescription(alloc corefoundation.CFAllocatorRef, certificate SecCertificateRef, err *corefoundation.CFErrorRef) (corefoundation.CFStringRef, error) {
	if _secCertificateCopyShortDescription == nil {
		return 0, symbolCallError("SecCertificateCopyShortDescription", "10.7", _secCertificateCopyShortDescriptionErr)
	}
	return _secCertificateCopyShortDescription(alloc, certificate, err), nil
}

// SecCertificateCopyShortDescription returns a copy of the short description of a certificate.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopyShortDescription(_:_:_:)
func SecCertificateCopyShortDescription(alloc corefoundation.CFAllocatorRef, certificate SecCertificateRef, err *corefoundation.CFErrorRef) corefoundation.CFStringRef {
	result, callErr := trySecCertificateCopyShortDescription(alloc, certificate, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateCopySubjectSummary func(certificate SecCertificateRef) corefoundation.CFStringRef
var _secCertificateCopySubjectSummaryErr error

func trySecCertificateCopySubjectSummary(certificate SecCertificateRef) (corefoundation.CFStringRef, error) {
	if _secCertificateCopySubjectSummary == nil {
		return 0, symbolCallError("SecCertificateCopySubjectSummary", "10.6", _secCertificateCopySubjectSummaryErr)
	}
	return _secCertificateCopySubjectSummary(certificate), nil
}

// SecCertificateCopySubjectSummary returns a human-readable summary of a certificate.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopySubjectSummary(_:)
func SecCertificateCopySubjectSummary(certificate SecCertificateRef) corefoundation.CFStringRef {
	result, callErr := trySecCertificateCopySubjectSummary(certificate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateCopyValues func(certificate SecCertificateRef, keys corefoundation.CFArrayRef, err *corefoundation.CFErrorRef) corefoundation.CFDictionaryRef
var _secCertificateCopyValuesErr error

func trySecCertificateCopyValues(certificate SecCertificateRef, keys corefoundation.CFArrayRef, err *corefoundation.CFErrorRef) (corefoundation.CFDictionaryRef, error) {
	if _secCertificateCopyValues == nil {
		return 0, symbolCallError("SecCertificateCopyValues", "10.7", _secCertificateCopyValuesErr)
	}
	return _secCertificateCopyValues(certificate, keys, err), nil
}

// SecCertificateCopyValues creates a dictionary that represents a certificate’s contents.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopyValues(_:_:_:)
func SecCertificateCopyValues(certificate SecCertificateRef, keys corefoundation.CFArrayRef, err *corefoundation.CFErrorRef) corefoundation.CFDictionaryRef {
	result, callErr := trySecCertificateCopyValues(certificate, keys, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateCreateWithData func(allocator corefoundation.CFAllocatorRef, data corefoundation.CFDataRef) SecCertificateRef
var _secCertificateCreateWithDataErr error

func trySecCertificateCreateWithData(allocator corefoundation.CFAllocatorRef, data corefoundation.CFDataRef) (SecCertificateRef, error) {
	if _secCertificateCreateWithData == nil {
		return 0, symbolCallError("SecCertificateCreateWithData", "10.6", _secCertificateCreateWithDataErr)
	}
	return _secCertificateCreateWithData(allocator, data), nil
}

// SecCertificateCreateWithData creates a certificate object from a DER representation of a certificate.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCreateWithData(_:_:)
func SecCertificateCreateWithData(allocator corefoundation.CFAllocatorRef, data corefoundation.CFDataRef) SecCertificateRef {
	result, callErr := trySecCertificateCreateWithData(allocator, data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateGetAlgorithmID func(certificate SecCertificateRef, algid unsafe.Pointer) int32
var _secCertificateGetAlgorithmIDErr error

func trySecCertificateGetAlgorithmID(certificate SecCertificateRef, algid unsafe.Pointer) (int32, error) {
	if _secCertificateGetAlgorithmID == nil {
		return 0, symbolCallError("SecCertificateGetAlgorithmID", "10.0", _secCertificateGetAlgorithmIDErr)
	}
	return _secCertificateGetAlgorithmID(certificate, algid), nil
}

// SecCertificateGetAlgorithmID retrieves the algorithm identifier for a certificate.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateGetAlgorithmID
func SecCertificateGetAlgorithmID(certificate SecCertificateRef, algid unsafe.Pointer) int32 {
	result, callErr := trySecCertificateGetAlgorithmID(certificate, algid)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateGetCLHandle func(certificate SecCertificateRef, clHandle unsafe.Pointer) int32
var _secCertificateGetCLHandleErr error

func trySecCertificateGetCLHandle(certificate SecCertificateRef, clHandle unsafe.Pointer) (int32, error) {
	if _secCertificateGetCLHandle == nil {
		return 0, symbolCallError("SecCertificateGetCLHandle", "10.0", _secCertificateGetCLHandleErr)
	}
	return _secCertificateGetCLHandle(certificate, clHandle), nil
}

// SecCertificateGetCLHandle retrieves the certificate library handle from a certificate object.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateGetCLHandle
func SecCertificateGetCLHandle(certificate SecCertificateRef, clHandle unsafe.Pointer) int32 {
	result, callErr := trySecCertificateGetCLHandle(certificate, clHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateGetData func(certificate SecCertificateRef, data unsafe.Pointer) int32
var _secCertificateGetDataErr error

func trySecCertificateGetData(certificate SecCertificateRef, data unsafe.Pointer) (int32, error) {
	if _secCertificateGetData == nil {
		return 0, symbolCallError("SecCertificateGetData", "10.0", _secCertificateGetDataErr)
	}
	return _secCertificateGetData(certificate, data), nil
}

// SecCertificateGetData retrieves the data for a certificate.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateGetData
func SecCertificateGetData(certificate SecCertificateRef, data unsafe.Pointer) int32 {
	result, callErr := trySecCertificateGetData(certificate, data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateGetIssuer func(certificate SecCertificateRef, issuer unsafe.Pointer) int32
var _secCertificateGetIssuerErr error

func trySecCertificateGetIssuer(certificate SecCertificateRef, issuer unsafe.Pointer) (int32, error) {
	if _secCertificateGetIssuer == nil {
		return 0, symbolCallError("SecCertificateGetIssuer", "10.0", _secCertificateGetIssuerErr)
	}
	return _secCertificateGetIssuer(certificate, issuer), nil
}

// SecCertificateGetIssuer unsupported.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateGetIssuer
func SecCertificateGetIssuer(certificate SecCertificateRef, issuer unsafe.Pointer) int32 {
	result, callErr := trySecCertificateGetIssuer(certificate, issuer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateGetSubject func(certificate SecCertificateRef, subject unsafe.Pointer) int32
var _secCertificateGetSubjectErr error

func trySecCertificateGetSubject(certificate SecCertificateRef, subject unsafe.Pointer) (int32, error) {
	if _secCertificateGetSubject == nil {
		return 0, symbolCallError("SecCertificateGetSubject", "10.0", _secCertificateGetSubjectErr)
	}
	return _secCertificateGetSubject(certificate, subject), nil
}

// SecCertificateGetSubject unsupported.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateGetSubject
func SecCertificateGetSubject(certificate SecCertificateRef, subject unsafe.Pointer) int32 {
	result, callErr := trySecCertificateGetSubject(certificate, subject)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateGetType func(certificate SecCertificateRef, certificateType unsafe.Pointer) int32
var _secCertificateGetTypeErr error

func trySecCertificateGetType(certificate SecCertificateRef, certificateType unsafe.Pointer) (int32, error) {
	if _secCertificateGetType == nil {
		return 0, symbolCallError("SecCertificateGetType", "10.0", _secCertificateGetTypeErr)
	}
	return _secCertificateGetType(certificate, certificateType), nil
}

// SecCertificateGetType retrieves the type of a specified certificate.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateGetType
func SecCertificateGetType(certificate SecCertificateRef, certificateType unsafe.Pointer) int32 {
	result, callErr := trySecCertificateGetType(certificate, certificateType)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateGetTypeID func() uint
var _secCertificateGetTypeIDErr error

func trySecCertificateGetTypeID() (uint, error) {
	if _secCertificateGetTypeID == nil {
		return 0, symbolCallError("SecCertificateGetTypeID", "10.3", _secCertificateGetTypeIDErr)
	}
	return _secCertificateGetTypeID(), nil
}

// SecCertificateGetTypeID returns the unique identifier of the opaque type to which a certificate object belongs.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateGetTypeID()
func SecCertificateGetTypeID() uint {
	result, callErr := trySecCertificateGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateSetPreference func(certificate SecCertificateRef, name corefoundation.CFStringRef, keyUsage uint32, date corefoundation.CFDateRef) int32
var _secCertificateSetPreferenceErr error

func trySecCertificateSetPreference(certificate SecCertificateRef, name corefoundation.CFStringRef, keyUsage uint32, date corefoundation.CFDateRef) (int32, error) {
	if _secCertificateSetPreference == nil {
		return 0, symbolCallError("SecCertificateSetPreference", "10.0", _secCertificateSetPreferenceErr)
	}
	return _secCertificateSetPreference(certificate, name, keyUsage, date), nil
}

// SecCertificateSetPreference sets the preferred certificate for a specified name, key use, and date.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateSetPreference
func SecCertificateSetPreference(certificate SecCertificateRef, name corefoundation.CFStringRef, keyUsage uint32, date corefoundation.CFDateRef) int32 {
	result, callErr := trySecCertificateSetPreference(certificate, name, keyUsage, date)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateSetPreferred func(certificate SecCertificateRef, name corefoundation.CFStringRef, keyUsage corefoundation.CFArrayRef) int32
var _secCertificateSetPreferredErr error

func trySecCertificateSetPreferred(certificate SecCertificateRef, name corefoundation.CFStringRef, keyUsage corefoundation.CFArrayRef) (int32, error) {
	if _secCertificateSetPreferred == nil {
		return 0, symbolCallError("SecCertificateSetPreferred", "10.7", _secCertificateSetPreferredErr)
	}
	return _secCertificateSetPreferred(certificate, name, keyUsage), nil
}

// SecCertificateSetPreferred sets the certificate that should be preferred for the specified name and key use.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateSetPreferred(_:_:_:)
func SecCertificateSetPreferred(certificate SecCertificateRef, name corefoundation.CFStringRef, keyUsage corefoundation.CFArrayRef) int32 {
	result, callErr := trySecCertificateSetPreferred(certificate, name, keyUsage)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCodeCheckValidity func(code SecCodeRef, flags SecCSFlags, requirement SecRequirementRef) int32
var _secCodeCheckValidityErr error

func trySecCodeCheckValidity(code SecCodeRef, flags SecCSFlags, requirement SecRequirementRef) (int32, error) {
	if _secCodeCheckValidity == nil {
		return 0, symbolCallError("SecCodeCheckValidity", "10.0", _secCodeCheckValidityErr)
	}
	return _secCodeCheckValidity(code, flags, requirement), nil
}

// SecCodeCheckValidity performs dynamic validation of signed code.
//
// See: https://developer.apple.com/documentation/Security/SecCodeCheckValidity(_:_:_:)
func SecCodeCheckValidity(code SecCodeRef, flags SecCSFlags, requirement SecRequirementRef) int32 {
	result, callErr := trySecCodeCheckValidity(code, flags, requirement)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCodeCheckValidityWithErrors func(code SecCodeRef, flags SecCSFlags, requirement SecRequirementRef, errors *corefoundation.CFErrorRef) int32
var _secCodeCheckValidityWithErrorsErr error

func trySecCodeCheckValidityWithErrors(code SecCodeRef, flags SecCSFlags, requirement SecRequirementRef, errors *corefoundation.CFErrorRef) (int32, error) {
	if _secCodeCheckValidityWithErrors == nil {
		return 0, symbolCallError("SecCodeCheckValidityWithErrors", "10.0", _secCodeCheckValidityWithErrorsErr)
	}
	return _secCodeCheckValidityWithErrors(code, flags, requirement, errors), nil
}

// SecCodeCheckValidityWithErrors performs dynamic validation of signed code and returns detailed error information in the case of failure.
//
// See: https://developer.apple.com/documentation/Security/SecCodeCheckValidityWithErrors(_:_:_:_:)
func SecCodeCheckValidityWithErrors(code SecCodeRef, flags SecCSFlags, requirement SecRequirementRef, errors *corefoundation.CFErrorRef) int32 {
	result, callErr := trySecCodeCheckValidityWithErrors(code, flags, requirement, errors)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCodeCopyDesignatedRequirement func(code SecStaticCodeRef, flags SecCSFlags, requirement *SecRequirementRef) int32
var _secCodeCopyDesignatedRequirementErr error

func trySecCodeCopyDesignatedRequirement(code SecStaticCodeRef, flags SecCSFlags, requirement *SecRequirementRef) (int32, error) {
	if _secCodeCopyDesignatedRequirement == nil {
		return 0, symbolCallError("SecCodeCopyDesignatedRequirement", "10.0", _secCodeCopyDesignatedRequirementErr)
	}
	return _secCodeCopyDesignatedRequirement(code, flags, requirement), nil
}

// SecCodeCopyDesignatedRequirement retrieves the designated code requirement of signed code.
//
// See: https://developer.apple.com/documentation/Security/SecCodeCopyDesignatedRequirement(_:_:_:)
func SecCodeCopyDesignatedRequirement(code SecStaticCodeRef, flags SecCSFlags, requirement *SecRequirementRef) int32 {
	result, callErr := trySecCodeCopyDesignatedRequirement(code, flags, requirement)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCodeCopyGuestWithAttributes func(host SecCodeRef, attributes corefoundation.CFDictionaryRef, flags SecCSFlags, guest *SecCodeRef) int32
var _secCodeCopyGuestWithAttributesErr error

func trySecCodeCopyGuestWithAttributes(host SecCodeRef, attributes corefoundation.CFDictionaryRef, flags SecCSFlags, guest *SecCodeRef) (int32, error) {
	if _secCodeCopyGuestWithAttributes == nil {
		return 0, symbolCallError("SecCodeCopyGuestWithAttributes", "10.0", _secCodeCopyGuestWithAttributesErr)
	}
	return _secCodeCopyGuestWithAttributes(host, attributes, flags, guest), nil
}

// SecCodeCopyGuestWithAttributes asks a code host to identify one of its guests given the type and value of specific attributes of the guest code.
//
// See: https://developer.apple.com/documentation/Security/SecCodeCopyGuestWithAttributes(_:_:_:_:)
func SecCodeCopyGuestWithAttributes(host SecCodeRef, attributes corefoundation.CFDictionaryRef, flags SecCSFlags, guest *SecCodeRef) int32 {
	result, callErr := trySecCodeCopyGuestWithAttributes(host, attributes, flags, guest)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCodeCopyHost func(guest SecCodeRef, flags SecCSFlags, host *SecCodeRef) int32
var _secCodeCopyHostErr error

func trySecCodeCopyHost(guest SecCodeRef, flags SecCSFlags, host *SecCodeRef) (int32, error) {
	if _secCodeCopyHost == nil {
		return 0, symbolCallError("SecCodeCopyHost", "10.0", _secCodeCopyHostErr)
	}
	return _secCodeCopyHost(guest, flags, host), nil
}

// SecCodeCopyHost retrieves the code object for the host of specified guest code.
//
// See: https://developer.apple.com/documentation/Security/SecCodeCopyHost(_:_:_:)
func SecCodeCopyHost(guest SecCodeRef, flags SecCSFlags, host *SecCodeRef) int32 {
	result, callErr := trySecCodeCopyHost(guest, flags, host)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCodeCopyPath func(staticCode SecStaticCodeRef, flags SecCSFlags, path *corefoundation.CFURLRef) int32
var _secCodeCopyPathErr error

func trySecCodeCopyPath(staticCode SecStaticCodeRef, flags SecCSFlags, path *corefoundation.CFURLRef) (int32, error) {
	if _secCodeCopyPath == nil {
		return 0, symbolCallError("SecCodeCopyPath", "10.0", _secCodeCopyPathErr)
	}
	return _secCodeCopyPath(staticCode, flags, path), nil
}

// SecCodeCopyPath retrieves the location on disk of signed code, given a code or static code object.
//
// See: https://developer.apple.com/documentation/Security/SecCodeCopyPath(_:_:_:)
func SecCodeCopyPath(staticCode SecStaticCodeRef, flags SecCSFlags, path *corefoundation.CFURLRef) int32 {
	result, callErr := trySecCodeCopyPath(staticCode, flags, path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCodeCopySelf func(flags SecCSFlags, self *SecCodeRef) int32
var _secCodeCopySelfErr error

func trySecCodeCopySelf(flags SecCSFlags, self *SecCodeRef) (int32, error) {
	if _secCodeCopySelf == nil {
		return 0, symbolCallError("SecCodeCopySelf", "10.0", _secCodeCopySelfErr)
	}
	return _secCodeCopySelf(flags, self), nil
}

// SecCodeCopySelf retrieves the code object for the code making the call.
//
// See: https://developer.apple.com/documentation/Security/SecCodeCopySelf(_:_:)
func SecCodeCopySelf(flags SecCSFlags, self *SecCodeRef) int32 {
	result, callErr := trySecCodeCopySelf(flags, self)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCodeCopySigningInformation func(code SecStaticCodeRef, flags SecCSFlags, information *corefoundation.CFDictionaryRef) int32
var _secCodeCopySigningInformationErr error

func trySecCodeCopySigningInformation(code SecStaticCodeRef, flags SecCSFlags, information *corefoundation.CFDictionaryRef) (int32, error) {
	if _secCodeCopySigningInformation == nil {
		return 0, symbolCallError("SecCodeCopySigningInformation", "10.0", _secCodeCopySigningInformationErr)
	}
	return _secCodeCopySigningInformation(code, flags, information), nil
}

// SecCodeCopySigningInformation retrieves various pieces of information from a code signature.
//
// See: https://developer.apple.com/documentation/Security/SecCodeCopySigningInformation(_:_:_:)
func SecCodeCopySigningInformation(code SecStaticCodeRef, flags SecCSFlags, information *corefoundation.CFDictionaryRef) int32 {
	result, callErr := trySecCodeCopySigningInformation(code, flags, information)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCodeCopyStaticCode func(code SecCodeRef, flags SecCSFlags, staticCode *SecStaticCodeRef) int32
var _secCodeCopyStaticCodeErr error

func trySecCodeCopyStaticCode(code SecCodeRef, flags SecCSFlags, staticCode *SecStaticCodeRef) (int32, error) {
	if _secCodeCopyStaticCode == nil {
		return 0, symbolCallError("SecCodeCopyStaticCode", "10.0", _secCodeCopyStaticCodeErr)
	}
	return _secCodeCopyStaticCode(code, flags, staticCode), nil
}

// SecCodeCopyStaticCode returns a static code object representing the on-disk version of the given running code.
//
// See: https://developer.apple.com/documentation/Security/SecCodeCopyStaticCode(_:_:_:)
func SecCodeCopyStaticCode(code SecCodeRef, flags SecCSFlags, staticCode *SecStaticCodeRef) int32 {
	result, callErr := trySecCodeCopyStaticCode(code, flags, staticCode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCodeCreateWithXPCMessage func(message unsafe.Pointer, flags SecCSFlags, target *SecCodeRef) int32
var _secCodeCreateWithXPCMessageErr error

func trySecCodeCreateWithXPCMessage(message unsafe.Pointer, flags SecCSFlags, target *SecCodeRef) (int32, error) {
	if _secCodeCreateWithXPCMessage == nil {
		return 0, symbolCallError("SecCodeCreateWithXPCMessage", "10.0", _secCodeCreateWithXPCMessageErr)
	}
	return _secCodeCreateWithXPCMessage(message, flags, target), nil
}

// SecCodeCreateWithXPCMessage.
//
// See: https://developer.apple.com/documentation/Security/SecCodeCreateWithXPCMessage(_:_:_:)
func SecCodeCreateWithXPCMessage(message unsafe.Pointer, flags SecCSFlags, target *SecCodeRef) int32 {
	result, callErr := trySecCodeCreateWithXPCMessage(message, flags, target)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCodeGetTypeID func() uint
var _secCodeGetTypeIDErr error

func trySecCodeGetTypeID() (uint, error) {
	if _secCodeGetTypeID == nil {
		return 0, symbolCallError("SecCodeGetTypeID", "10.0", _secCodeGetTypeIDErr)
	}
	return _secCodeGetTypeID(), nil
}

// SecCodeGetTypeID returns the unique identifier of the opaque type to which a code object belongs.
//
// See: https://developer.apple.com/documentation/Security/SecCodeGetTypeID()
func SecCodeGetTypeID() uint {
	result, callErr := trySecCodeGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCodeMapMemory func(code SecStaticCodeRef, flags SecCSFlags) int32
var _secCodeMapMemoryErr error

func trySecCodeMapMemory(code SecStaticCodeRef, flags SecCSFlags) (int32, error) {
	if _secCodeMapMemory == nil {
		return 0, symbolCallError("SecCodeMapMemory", "10.0", _secCodeMapMemoryErr)
	}
	return _secCodeMapMemory(code, flags), nil
}

// SecCodeMapMemory asks the kernel to accept the signing information currently attached to a code object and uses it to validate memory page-ins.
//
// See: https://developer.apple.com/documentation/Security/SecCodeMapMemory(_:_:)
func SecCodeMapMemory(code SecStaticCodeRef, flags SecCSFlags) int32 {
	result, callErr := trySecCodeMapMemory(code, flags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCodeValidateFileResource func(code SecStaticCodeRef, relativePath corefoundation.CFStringRef, fileData corefoundation.CFDataRef, flags SecCSFlags) int32
var _secCodeValidateFileResourceErr error

func trySecCodeValidateFileResource(code SecStaticCodeRef, relativePath corefoundation.CFStringRef, fileData corefoundation.CFDataRef, flags SecCSFlags) (int32, error) {
	if _secCodeValidateFileResource == nil {
		return 0, symbolCallError("SecCodeValidateFileResource", "10.13", _secCodeValidateFileResourceErr)
	}
	return _secCodeValidateFileResource(code, relativePath, fileData, flags), nil
}

// SecCodeValidateFileResource.
//
// See: https://developer.apple.com/documentation/Security/SecCodeValidateFileResource(_:_:_:_:)
func SecCodeValidateFileResource(code SecStaticCodeRef, relativePath corefoundation.CFStringRef, fileData corefoundation.CFDataRef, flags SecCSFlags) int32 {
	result, callErr := trySecCodeValidateFileResource(code, relativePath, fileData, flags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCopyErrorMessageString func(status int32, reserved unsafe.Pointer) corefoundation.CFStringRef
var _secCopyErrorMessageStringErr error

func trySecCopyErrorMessageString(status int32, reserved unsafe.Pointer) (corefoundation.CFStringRef, error) {
	if _secCopyErrorMessageString == nil {
		return 0, symbolCallError("SecCopyErrorMessageString", "10.3", _secCopyErrorMessageStringErr)
	}
	return _secCopyErrorMessageString(status, reserved), nil
}

// SecCopyErrorMessageString returns a string explaining the meaning of a security result code.
//
// See: https://developer.apple.com/documentation/Security/SecCopyErrorMessageString(_:_:)
func SecCopyErrorMessageString(status int32, reserved unsafe.Pointer) corefoundation.CFStringRef {
	result, callErr := trySecCopyErrorMessageString(status, reserved)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCreateSharedWebCredentialPassword func() corefoundation.CFStringRef
var _secCreateSharedWebCredentialPasswordErr error

func trySecCreateSharedWebCredentialPassword() (corefoundation.CFStringRef, error) {
	if _secCreateSharedWebCredentialPassword == nil {
		return 0, symbolCallError("SecCreateSharedWebCredentialPassword", "11.0", _secCreateSharedWebCredentialPasswordErr)
	}
	return _secCreateSharedWebCredentialPassword(), nil
}

// SecCreateSharedWebCredentialPassword returns a randomly generated password.
//
// See: https://developer.apple.com/documentation/Security/SecCreateSharedWebCredentialPassword()
func SecCreateSharedWebCredentialPassword() corefoundation.CFStringRef {
	result, callErr := trySecCreateSharedWebCredentialPassword()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secHostCreateGuest func(host SecGuestRef, status uint32, path corefoundation.CFURLRef, attributes corefoundation.CFDictionaryRef, flags SecCSFlags, newGuest *SecGuestRef) int32
var _secHostCreateGuestErr error

func trySecHostCreateGuest(host SecGuestRef, status uint32, path corefoundation.CFURLRef, attributes corefoundation.CFDictionaryRef, flags SecCSFlags, newGuest *SecGuestRef) (int32, error) {
	if _secHostCreateGuest == nil {
		return 0, symbolCallError("SecHostCreateGuest", "10.5", _secHostCreateGuestErr)
	}
	return _secHostCreateGuest(host, status, path, attributes, flags, newGuest), nil
}

// SecHostCreateGuest creates a new guest and describes its initial properties.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/Security/SecHostCreateGuest
func SecHostCreateGuest(host SecGuestRef, status uint32, path corefoundation.CFURLRef, attributes corefoundation.CFDictionaryRef, flags SecCSFlags, newGuest *SecGuestRef) int32 {
	result, callErr := trySecHostCreateGuest(host, status, path, attributes, flags, newGuest)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secHostRemoveGuest func(host SecGuestRef, guest SecGuestRef, flags SecCSFlags) int32
var _secHostRemoveGuestErr error

func trySecHostRemoveGuest(host SecGuestRef, guest SecGuestRef, flags SecCSFlags) (int32, error) {
	if _secHostRemoveGuest == nil {
		return 0, symbolCallError("SecHostRemoveGuest", "10.5", _secHostRemoveGuestErr)
	}
	return _secHostRemoveGuest(host, guest, flags), nil
}

// SecHostRemoveGuest removes a guest from a host.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/Security/SecHostRemoveGuest
func SecHostRemoveGuest(host SecGuestRef, guest SecGuestRef, flags SecCSFlags) int32 {
	result, callErr := trySecHostRemoveGuest(host, guest, flags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secHostSelectGuest func(guestRef SecGuestRef, flags SecCSFlags) int32
var _secHostSelectGuestErr error

func trySecHostSelectGuest(guestRef SecGuestRef, flags SecCSFlags) (int32, error) {
	if _secHostSelectGuest == nil {
		return 0, symbolCallError("SecHostSelectGuest", "10.5", _secHostSelectGuestErr)
	}
	return _secHostSelectGuest(guestRef, flags), nil
}

// SecHostSelectGuest makes the calling thread the proxy for a specified guest.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/Security/SecHostSelectGuest
func SecHostSelectGuest(guestRef SecGuestRef, flags SecCSFlags) int32 {
	result, callErr := trySecHostSelectGuest(guestRef, flags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secHostSelectedGuest func(flags SecCSFlags, guestRef *SecGuestRef) int32
var _secHostSelectedGuestErr error

func trySecHostSelectedGuest(flags SecCSFlags, guestRef *SecGuestRef) (int32, error) {
	if _secHostSelectedGuest == nil {
		return 0, symbolCallError("SecHostSelectedGuest", "10.5", _secHostSelectedGuestErr)
	}
	return _secHostSelectedGuest(flags, guestRef), nil
}

// SecHostSelectedGuest retrieves the handle for the guest currently selected for the calling thread.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/Security/SecHostSelectedGuest
func SecHostSelectedGuest(flags SecCSFlags, guestRef *SecGuestRef) int32 {
	result, callErr := trySecHostSelectedGuest(flags, guestRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secHostSetGuestStatus func(guestRef SecGuestRef, status uint32, attributes corefoundation.CFDictionaryRef, flags SecCSFlags) int32
var _secHostSetGuestStatusErr error

func trySecHostSetGuestStatus(guestRef SecGuestRef, status uint32, attributes corefoundation.CFDictionaryRef, flags SecCSFlags) (int32, error) {
	if _secHostSetGuestStatus == nil {
		return 0, symbolCallError("SecHostSetGuestStatus", "10.5", _secHostSetGuestStatusErr)
	}
	return _secHostSetGuestStatus(guestRef, status, attributes, flags), nil
}

// SecHostSetGuestStatus updates the status and attributes of a particular guest.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/Security/SecHostSetGuestStatus
func SecHostSetGuestStatus(guestRef SecGuestRef, status uint32, attributes corefoundation.CFDictionaryRef, flags SecCSFlags) int32 {
	result, callErr := trySecHostSetGuestStatus(guestRef, status, attributes, flags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secHostSetHostingPort func(hostingPort uint32, flags SecCSFlags) int32
var _secHostSetHostingPortErr error

func trySecHostSetHostingPort(hostingPort uint32, flags SecCSFlags) (int32, error) {
	if _secHostSetHostingPort == nil {
		return 0, symbolCallError("SecHostSetHostingPort", "10.5", _secHostSetHostingPortErr)
	}
	return _secHostSetHostingPort(hostingPort, flags), nil
}

// SecHostSetHostingPort tells code signing services that the calling code will directly respond to hosting inquiries over the given port.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/Security/SecHostSetHostingPort
func SecHostSetHostingPort(hostingPort uint32, flags SecCSFlags) int32 {
	result, callErr := trySecHostSetHostingPort(hostingPort, flags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secIdentityCopyCertificate func(identityRef SecIdentityRef, certificateRef *SecCertificateRef) int32
var _secIdentityCopyCertificateErr error

func trySecIdentityCopyCertificate(identityRef SecIdentityRef, certificateRef *SecCertificateRef) (int32, error) {
	if _secIdentityCopyCertificate == nil {
		return 0, symbolCallError("SecIdentityCopyCertificate", "10.3", _secIdentityCopyCertificateErr)
	}
	return _secIdentityCopyCertificate(identityRef, certificateRef), nil
}

// SecIdentityCopyCertificate retrieves a certificate associated with an identity.
//
// See: https://developer.apple.com/documentation/Security/SecIdentityCopyCertificate(_:_:)
func SecIdentityCopyCertificate(identityRef SecIdentityRef, certificateRef *SecCertificateRef) int32 {
	result, callErr := trySecIdentityCopyCertificate(identityRef, certificateRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secIdentityCopyPreference func(name corefoundation.CFStringRef, keyUsage CSSM_KEYUSE, validIssuers corefoundation.CFArrayRef, identity *SecIdentityRef) int32
var _secIdentityCopyPreferenceErr error

func trySecIdentityCopyPreference(name corefoundation.CFStringRef, keyUsage CSSM_KEYUSE, validIssuers corefoundation.CFArrayRef, identity *SecIdentityRef) (int32, error) {
	if _secIdentityCopyPreference == nil {
		return 0, symbolCallError("SecIdentityCopyPreference", "10.0", _secIdentityCopyPreferenceErr)
	}
	return _secIdentityCopyPreference(name, keyUsage, validIssuers, identity), nil
}

// SecIdentityCopyPreference returns the preferred identity for the specified name and key use.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecIdentityCopyPreference
func SecIdentityCopyPreference(name corefoundation.CFStringRef, keyUsage CSSM_KEYUSE, validIssuers corefoundation.CFArrayRef, identity *SecIdentityRef) int32 {
	result, callErr := trySecIdentityCopyPreference(name, keyUsage, validIssuers, identity)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secIdentityCopyPreferred func(name corefoundation.CFStringRef, keyUsage corefoundation.CFArrayRef, validIssuers corefoundation.CFArrayRef) SecIdentityRef
var _secIdentityCopyPreferredErr error

func trySecIdentityCopyPreferred(name corefoundation.CFStringRef, keyUsage corefoundation.CFArrayRef, validIssuers corefoundation.CFArrayRef) (SecIdentityRef, error) {
	if _secIdentityCopyPreferred == nil {
		return 0, symbolCallError("SecIdentityCopyPreferred", "10.7", _secIdentityCopyPreferredErr)
	}
	return _secIdentityCopyPreferred(name, keyUsage, validIssuers), nil
}

// SecIdentityCopyPreferred retrieves the preferred identity for the specified name and key use.
//
// See: https://developer.apple.com/documentation/Security/SecIdentityCopyPreferred(_:_:_:)
func SecIdentityCopyPreferred(name corefoundation.CFStringRef, keyUsage corefoundation.CFArrayRef, validIssuers corefoundation.CFArrayRef) SecIdentityRef {
	result, callErr := trySecIdentityCopyPreferred(name, keyUsage, validIssuers)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secIdentityCopyPrivateKey func(identityRef SecIdentityRef, privateKeyRef *SecKeyRef) int32
var _secIdentityCopyPrivateKeyErr error

func trySecIdentityCopyPrivateKey(identityRef SecIdentityRef, privateKeyRef *SecKeyRef) (int32, error) {
	if _secIdentityCopyPrivateKey == nil {
		return 0, symbolCallError("SecIdentityCopyPrivateKey", "10.3", _secIdentityCopyPrivateKeyErr)
	}
	return _secIdentityCopyPrivateKey(identityRef, privateKeyRef), nil
}

// SecIdentityCopyPrivateKey retrieves the private key associated with an identity.
//
// See: https://developer.apple.com/documentation/Security/SecIdentityCopyPrivateKey(_:_:)
func SecIdentityCopyPrivateKey(identityRef SecIdentityRef, privateKeyRef *SecKeyRef) int32 {
	result, callErr := trySecIdentityCopyPrivateKey(identityRef, privateKeyRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secIdentityCopySystemIdentity func(domain corefoundation.CFStringRef, idRef *SecIdentityRef, actualDomain *corefoundation.CFStringRef) int32
var _secIdentityCopySystemIdentityErr error

func trySecIdentityCopySystemIdentity(domain corefoundation.CFStringRef, idRef *SecIdentityRef, actualDomain *corefoundation.CFStringRef) (int32, error) {
	if _secIdentityCopySystemIdentity == nil {
		return 0, symbolCallError("SecIdentityCopySystemIdentity", "10.5", _secIdentityCopySystemIdentityErr)
	}
	return _secIdentityCopySystemIdentity(domain, idRef, actualDomain), nil
}

// SecIdentityCopySystemIdentity obtains the system identity associated with a specified domain.
//
// See: https://developer.apple.com/documentation/Security/SecIdentityCopySystemIdentity(_:_:_:)
func SecIdentityCopySystemIdentity(domain corefoundation.CFStringRef, idRef *SecIdentityRef, actualDomain *corefoundation.CFStringRef) int32 {
	result, callErr := trySecIdentityCopySystemIdentity(domain, idRef, actualDomain)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secIdentityCreate func(allocator corefoundation.CFAllocatorRef, certificate SecCertificateRef, privateKey SecKeyRef) SecIdentityRef
var _secIdentityCreateErr error

func trySecIdentityCreate(allocator corefoundation.CFAllocatorRef, certificate SecCertificateRef, privateKey SecKeyRef) (SecIdentityRef, error) {
	if _secIdentityCreate == nil {
		return 0, symbolCallError("SecIdentityCreate", "10.12", _secIdentityCreateErr)
	}
	return _secIdentityCreate(allocator, certificate, privateKey), nil
}

// SecIdentityCreate.
//
// See: https://developer.apple.com/documentation/Security/SecIdentityCreate(_:_:_:)
func SecIdentityCreate(allocator corefoundation.CFAllocatorRef, certificate SecCertificateRef, privateKey SecKeyRef) SecIdentityRef {
	result, callErr := trySecIdentityCreate(allocator, certificate, privateKey)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secIdentityCreateWithCertificate func(keychainOrArray corefoundation.CFTypeRef, certificateRef SecCertificateRef, identityRef *SecIdentityRef) int32
var _secIdentityCreateWithCertificateErr error

func trySecIdentityCreateWithCertificate(keychainOrArray corefoundation.CFTypeRef, certificateRef SecCertificateRef, identityRef *SecIdentityRef) (int32, error) {
	if _secIdentityCreateWithCertificate == nil {
		return 0, symbolCallError("SecIdentityCreateWithCertificate", "10.5", _secIdentityCreateWithCertificateErr)
	}
	return _secIdentityCreateWithCertificate(keychainOrArray, certificateRef, identityRef), nil
}

// SecIdentityCreateWithCertificate creates a new identity for a certificate and its associated private key.
//
// See: https://developer.apple.com/documentation/Security/SecIdentityCreateWithCertificate(_:_:_:)
func SecIdentityCreateWithCertificate(keychainOrArray corefoundation.CFTypeRef, certificateRef SecCertificateRef, identityRef *SecIdentityRef) int32 {
	result, callErr := trySecIdentityCreateWithCertificate(keychainOrArray, certificateRef, identityRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secIdentityGetTypeID func() uint
var _secIdentityGetTypeIDErr error

func trySecIdentityGetTypeID() (uint, error) {
	if _secIdentityGetTypeID == nil {
		return 0, symbolCallError("SecIdentityGetTypeID", "10.3", _secIdentityGetTypeIDErr)
	}
	return _secIdentityGetTypeID(), nil
}

// SecIdentityGetTypeID returns the unique identifier of the opaque type to which an identity object belongs.
//
// See: https://developer.apple.com/documentation/Security/SecIdentityGetTypeID()
func SecIdentityGetTypeID() uint {
	result, callErr := trySecIdentityGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secIdentitySearchCopyNext func(searchRef SecIdentitySearchRef, identity *SecIdentityRef) int32
var _secIdentitySearchCopyNextErr error

func trySecIdentitySearchCopyNext(searchRef SecIdentitySearchRef, identity *SecIdentityRef) (int32, error) {
	if _secIdentitySearchCopyNext == nil {
		return 0, symbolCallError("SecIdentitySearchCopyNext", "10.0", _secIdentitySearchCopyNextErr)
	}
	return _secIdentitySearchCopyNext(searchRef, identity), nil
}

// SecIdentitySearchCopyNext finds the next identity matching specified search criteria
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecIdentitySearchCopyNext
func SecIdentitySearchCopyNext(searchRef SecIdentitySearchRef, identity *SecIdentityRef) int32 {
	result, callErr := trySecIdentitySearchCopyNext(searchRef, identity)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secIdentitySearchCreate func(keychainOrArray corefoundation.CFTypeRef, keyUsage CSSM_KEYUSE, searchRef *SecIdentitySearchRef) int32
var _secIdentitySearchCreateErr error

func trySecIdentitySearchCreate(keychainOrArray corefoundation.CFTypeRef, keyUsage CSSM_KEYUSE, searchRef *SecIdentitySearchRef) (int32, error) {
	if _secIdentitySearchCreate == nil {
		return 0, symbolCallError("SecIdentitySearchCreate", "10.0", _secIdentitySearchCreateErr)
	}
	return _secIdentitySearchCreate(keychainOrArray, keyUsage, searchRef), nil
}

// SecIdentitySearchCreate creates a search object for finding identities.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecIdentitySearchCreate
func SecIdentitySearchCreate(keychainOrArray corefoundation.CFTypeRef, keyUsage CSSM_KEYUSE, searchRef *SecIdentitySearchRef) int32 {
	result, callErr := trySecIdentitySearchCreate(keychainOrArray, keyUsage, searchRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secIdentitySearchGetTypeID func() uint
var _secIdentitySearchGetTypeIDErr error

func trySecIdentitySearchGetTypeID() (uint, error) {
	if _secIdentitySearchGetTypeID == nil {
		return 0, symbolCallError("SecIdentitySearchGetTypeID", "10.0", _secIdentitySearchGetTypeIDErr)
	}
	return _secIdentitySearchGetTypeID(), nil
}

// SecIdentitySearchGetTypeID returns the unique identifier of the opaque type to which a [SecIdentitySearch] object belongs.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecIdentitySearchGetTypeID
func SecIdentitySearchGetTypeID() uint {
	result, callErr := trySecIdentitySearchGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secIdentitySetPreference func(identity SecIdentityRef, name corefoundation.CFStringRef, keyUsage CSSM_KEYUSE) int32
var _secIdentitySetPreferenceErr error

func trySecIdentitySetPreference(identity SecIdentityRef, name corefoundation.CFStringRef, keyUsage CSSM_KEYUSE) (int32, error) {
	if _secIdentitySetPreference == nil {
		return 0, symbolCallError("SecIdentitySetPreference", "10.0", _secIdentitySetPreferenceErr)
	}
	return _secIdentitySetPreference(identity, name, keyUsage), nil
}

// SecIdentitySetPreference sets the preferred identity for the specified name and key use.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecIdentitySetPreference
func SecIdentitySetPreference(identity SecIdentityRef, name corefoundation.CFStringRef, keyUsage CSSM_KEYUSE) int32 {
	result, callErr := trySecIdentitySetPreference(identity, name, keyUsage)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secIdentitySetPreferred func(identity SecIdentityRef, name corefoundation.CFStringRef, keyUsage corefoundation.CFArrayRef) int32
var _secIdentitySetPreferredErr error

func trySecIdentitySetPreferred(identity SecIdentityRef, name corefoundation.CFStringRef, keyUsage corefoundation.CFArrayRef) (int32, error) {
	if _secIdentitySetPreferred == nil {
		return 0, symbolCallError("SecIdentitySetPreferred", "10.7", _secIdentitySetPreferredErr)
	}
	return _secIdentitySetPreferred(identity, name, keyUsage), nil
}

// SecIdentitySetPreferred sets the identity that should be preferred for the specified name and key use.
//
// See: https://developer.apple.com/documentation/Security/SecIdentitySetPreferred(_:_:_:)
func SecIdentitySetPreferred(identity SecIdentityRef, name corefoundation.CFStringRef, keyUsage corefoundation.CFArrayRef) int32 {
	result, callErr := trySecIdentitySetPreferred(identity, name, keyUsage)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secIdentitySetSystemIdentity func(domain corefoundation.CFStringRef, idRef SecIdentityRef) int32
var _secIdentitySetSystemIdentityErr error

func trySecIdentitySetSystemIdentity(domain corefoundation.CFStringRef, idRef SecIdentityRef) (int32, error) {
	if _secIdentitySetSystemIdentity == nil {
		return 0, symbolCallError("SecIdentitySetSystemIdentity", "10.5", _secIdentitySetSystemIdentityErr)
	}
	return _secIdentitySetSystemIdentity(domain, idRef), nil
}

// SecIdentitySetSystemIdentity assigns the system identity to be associated with a specified domain.
//
// See: https://developer.apple.com/documentation/Security/SecIdentitySetSystemIdentity(_:_:)
func SecIdentitySetSystemIdentity(domain corefoundation.CFStringRef, idRef SecIdentityRef) int32 {
	result, callErr := trySecIdentitySetSystemIdentity(domain, idRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secItemAdd func(attributes corefoundation.CFDictionaryRef, result *corefoundation.CFTypeRef) int32
var _secItemAddErr error

func trySecItemAdd(attributes corefoundation.CFDictionaryRef, result *corefoundation.CFTypeRef) (int32, error) {
	if _secItemAdd == nil {
		return 0, symbolCallError("SecItemAdd", "10.6", _secItemAddErr)
	}
	return _secItemAdd(attributes, result), nil
}

// SecItemAdd adds one or more items to a keychain.
//
// See: https://developer.apple.com/documentation/Security/SecItemAdd(_:_:)
func SecItemAdd(attributes corefoundation.CFDictionaryRef, result *corefoundation.CFTypeRef) int32 {
	result0, callErr := trySecItemAdd(attributes, result)
	if callErr != nil {
		panic(callErr)
	}
	return result0
}

var _secItemCopyMatching func(query corefoundation.CFDictionaryRef, result *corefoundation.CFTypeRef) int32
var _secItemCopyMatchingErr error

func trySecItemCopyMatching(query corefoundation.CFDictionaryRef, result *corefoundation.CFTypeRef) (int32, error) {
	if _secItemCopyMatching == nil {
		return 0, symbolCallError("SecItemCopyMatching", "10.6", _secItemCopyMatchingErr)
	}
	return _secItemCopyMatching(query, result), nil
}

// SecItemCopyMatching returns one or more keychain items that match a search query, or copies attributes of specific keychain items.
//
// See: https://developer.apple.com/documentation/Security/SecItemCopyMatching(_:_:)
func SecItemCopyMatching(query corefoundation.CFDictionaryRef, result *corefoundation.CFTypeRef) int32 {
	result0, callErr := trySecItemCopyMatching(query, result)
	if callErr != nil {
		panic(callErr)
	}
	return result0
}

var _secItemDelete func(query corefoundation.CFDictionaryRef) int32
var _secItemDeleteErr error

func trySecItemDelete(query corefoundation.CFDictionaryRef) (int32, error) {
	if _secItemDelete == nil {
		return 0, symbolCallError("SecItemDelete", "10.6", _secItemDeleteErr)
	}
	return _secItemDelete(query), nil
}

// SecItemDelete deletes items that match a search query.
//
// See: https://developer.apple.com/documentation/Security/SecItemDelete(_:)
func SecItemDelete(query corefoundation.CFDictionaryRef) int32 {
	result, callErr := trySecItemDelete(query)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secItemExport func(secItemOrArray corefoundation.CFTypeRef, outputFormat SecExternalFormat, flags SecItemImportExportFlags, keyParams *SecItemImportExportKeyParameters, exportedData *corefoundation.CFDataRef) int32
var _secItemExportErr error

func trySecItemExport(secItemOrArray corefoundation.CFTypeRef, outputFormat SecExternalFormat, flags SecItemImportExportFlags, keyParams *SecItemImportExportKeyParameters, exportedData *corefoundation.CFDataRef) (int32, error) {
	if _secItemExport == nil {
		return 0, symbolCallError("SecItemExport", "10.7", _secItemExportErr)
	}
	return _secItemExport(secItemOrArray, outputFormat, flags, keyParams, exportedData), nil
}

// SecItemExport exports one or more certificates, keys, or identities.
//
// See: https://developer.apple.com/documentation/Security/SecItemExport(_:_:_:_:_:)
func SecItemExport(secItemOrArray corefoundation.CFTypeRef, outputFormat SecExternalFormat, flags SecItemImportExportFlags, keyParams *SecItemImportExportKeyParameters, exportedData *corefoundation.CFDataRef) int32 {
	result, callErr := trySecItemExport(secItemOrArray, outputFormat, flags, keyParams, exportedData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secItemImport func(importedData corefoundation.CFDataRef, fileNameOrExtension corefoundation.CFStringRef, inputFormat *SecExternalFormat, itemType *SecExternalItemType, flags SecItemImportExportFlags, keyParams *SecItemImportExportKeyParameters, importKeychain SecKeychainRef, outItems *corefoundation.CFArrayRef) int32
var _secItemImportErr error

func trySecItemImport(importedData corefoundation.CFDataRef, fileNameOrExtension corefoundation.CFStringRef, inputFormat *SecExternalFormat, itemType *SecExternalItemType, flags SecItemImportExportFlags, keyParams *SecItemImportExportKeyParameters, importKeychain SecKeychainRef, outItems *corefoundation.CFArrayRef) (int32, error) {
	if _secItemImport == nil {
		return 0, symbolCallError("SecItemImport", "10.7", _secItemImportErr)
	}
	return _secItemImport(importedData, fileNameOrExtension, inputFormat, itemType, flags, keyParams, importKeychain, outItems), nil
}

// SecItemImport imports one or more certificates, keys, or identities and optionally adds them to a keychain.
//
// See: https://developer.apple.com/documentation/Security/SecItemImport(_:_:_:_:_:_:_:_:)
func SecItemImport(importedData corefoundation.CFDataRef, fileNameOrExtension corefoundation.CFStringRef, inputFormat *SecExternalFormat, itemType *SecExternalItemType, flags SecItemImportExportFlags, keyParams *SecItemImportExportKeyParameters, importKeychain SecKeychainRef, outItems *corefoundation.CFArrayRef) int32 {
	result, callErr := trySecItemImport(importedData, fileNameOrExtension, inputFormat, itemType, flags, keyParams, importKeychain, outItems)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secItemUpdate func(query corefoundation.CFDictionaryRef, attributesToUpdate corefoundation.CFDictionaryRef) int32
var _secItemUpdateErr error

func trySecItemUpdate(query corefoundation.CFDictionaryRef, attributesToUpdate corefoundation.CFDictionaryRef) (int32, error) {
	if _secItemUpdate == nil {
		return 0, symbolCallError("SecItemUpdate", "10.6", _secItemUpdateErr)
	}
	return _secItemUpdate(query, attributesToUpdate), nil
}

// SecItemUpdate modifies items that match a search query.
//
// See: https://developer.apple.com/documentation/Security/SecItemUpdate(_:_:)
func SecItemUpdate(query corefoundation.CFDictionaryRef, attributesToUpdate corefoundation.CFDictionaryRef) int32 {
	result, callErr := trySecItemUpdate(query, attributesToUpdate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secKeyCopyAttributes func(key SecKeyRef) corefoundation.CFDictionaryRef
var _secKeyCopyAttributesErr error

func trySecKeyCopyAttributes(key SecKeyRef) (corefoundation.CFDictionaryRef, error) {
	if _secKeyCopyAttributes == nil {
		return 0, symbolCallError("SecKeyCopyAttributes", "10.12", _secKeyCopyAttributesErr)
	}
	return _secKeyCopyAttributes(key), nil
}

// SecKeyCopyAttributes gets the attributes of a given key.
//
// See: https://developer.apple.com/documentation/Security/SecKeyCopyAttributes(_:)
func SecKeyCopyAttributes(key SecKeyRef) corefoundation.CFDictionaryRef {
	result, callErr := trySecKeyCopyAttributes(key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secKeyCopyExternalRepresentation func(key SecKeyRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef
var _secKeyCopyExternalRepresentationErr error

func trySecKeyCopyExternalRepresentation(key SecKeyRef, err *corefoundation.CFErrorRef) (corefoundation.CFDataRef, error) {
	if _secKeyCopyExternalRepresentation == nil {
		return 0, symbolCallError("SecKeyCopyExternalRepresentation", "10.12", _secKeyCopyExternalRepresentationErr)
	}
	return _secKeyCopyExternalRepresentation(key, err), nil
}

// SecKeyCopyExternalRepresentation returns an external representation of the given key suitable for the key’s type.
//
// See: https://developer.apple.com/documentation/Security/SecKeyCopyExternalRepresentation(_:_:)
func SecKeyCopyExternalRepresentation(key SecKeyRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef {
	result, callErr := trySecKeyCopyExternalRepresentation(key, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secKeyCopyKeyExchangeResult func(privateKey SecKeyRef, algorithm SecKeyAlgorithm, publicKey SecKeyRef, parameters corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef
var _secKeyCopyKeyExchangeResultErr error

func trySecKeyCopyKeyExchangeResult(privateKey SecKeyRef, algorithm SecKeyAlgorithm, publicKey SecKeyRef, parameters corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) (corefoundation.CFDataRef, error) {
	if _secKeyCopyKeyExchangeResult == nil {
		return 0, symbolCallError("SecKeyCopyKeyExchangeResult", "10.12", _secKeyCopyKeyExchangeResultErr)
	}
	return _secKeyCopyKeyExchangeResult(privateKey, algorithm, publicKey, parameters, err), nil
}

// SecKeyCopyKeyExchangeResult performs the Diffie-Hellman style of key exchange with optional key-derivation steps.
//
// See: https://developer.apple.com/documentation/Security/SecKeyCopyKeyExchangeResult(_:_:_:_:_:)
func SecKeyCopyKeyExchangeResult(privateKey SecKeyRef, algorithm SecKeyAlgorithm, publicKey SecKeyRef, parameters corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef {
	result, callErr := trySecKeyCopyKeyExchangeResult(privateKey, algorithm, publicKey, parameters, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secKeyCopyPublicKey func(key SecKeyRef) SecKeyRef
var _secKeyCopyPublicKeyErr error

func trySecKeyCopyPublicKey(key SecKeyRef) (SecKeyRef, error) {
	if _secKeyCopyPublicKey == nil {
		return 0, symbolCallError("SecKeyCopyPublicKey", "10.12", _secKeyCopyPublicKeyErr)
	}
	return _secKeyCopyPublicKey(key), nil
}

// SecKeyCopyPublicKey gets the public key associated with the given private key.
//
// See: https://developer.apple.com/documentation/Security/SecKeyCopyPublicKey(_:)
func SecKeyCopyPublicKey(key SecKeyRef) SecKeyRef {
	result, callErr := trySecKeyCopyPublicKey(key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secKeyCreateDecryptedData func(key SecKeyRef, algorithm SecKeyAlgorithm, ciphertext corefoundation.CFDataRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef
var _secKeyCreateDecryptedDataErr error

func trySecKeyCreateDecryptedData(key SecKeyRef, algorithm SecKeyAlgorithm, ciphertext corefoundation.CFDataRef, err *corefoundation.CFErrorRef) (corefoundation.CFDataRef, error) {
	if _secKeyCreateDecryptedData == nil {
		return 0, symbolCallError("SecKeyCreateDecryptedData", "10.12", _secKeyCreateDecryptedDataErr)
	}
	return _secKeyCreateDecryptedData(key, algorithm, ciphertext, err), nil
}

// SecKeyCreateDecryptedData decrypts a block of data using a private key and specified algorithm.
//
// See: https://developer.apple.com/documentation/Security/SecKeyCreateDecryptedData(_:_:_:_:)
func SecKeyCreateDecryptedData(key SecKeyRef, algorithm SecKeyAlgorithm, ciphertext corefoundation.CFDataRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef {
	result, callErr := trySecKeyCreateDecryptedData(key, algorithm, ciphertext, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secKeyCreateEncryptedData func(key SecKeyRef, algorithm SecKeyAlgorithm, plaintext corefoundation.CFDataRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef
var _secKeyCreateEncryptedDataErr error

func trySecKeyCreateEncryptedData(key SecKeyRef, algorithm SecKeyAlgorithm, plaintext corefoundation.CFDataRef, err *corefoundation.CFErrorRef) (corefoundation.CFDataRef, error) {
	if _secKeyCreateEncryptedData == nil {
		return 0, symbolCallError("SecKeyCreateEncryptedData", "10.12", _secKeyCreateEncryptedDataErr)
	}
	return _secKeyCreateEncryptedData(key, algorithm, plaintext, err), nil
}

// SecKeyCreateEncryptedData encrypts a block of data using a public key and specified algorithm.
//
// See: https://developer.apple.com/documentation/Security/SecKeyCreateEncryptedData(_:_:_:_:)
func SecKeyCreateEncryptedData(key SecKeyRef, algorithm SecKeyAlgorithm, plaintext corefoundation.CFDataRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef {
	result, callErr := trySecKeyCreateEncryptedData(key, algorithm, plaintext, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secKeyCreatePair func(keychainRef SecKeychainRef, algorithm CSSM_ALGORITHMS, keySizeInBits uint32, contextHandle CSSM_CC_HANDLE, publicKeyUsage CSSM_KEYUSE, publicKeyAttr uint32, privateKeyUsage CSSM_KEYUSE, privateKeyAttr uint32, initialAccess SecAccessRef, publicKey *SecKeyRef, privateKey *SecKeyRef) int32
var _secKeyCreatePairErr error

func trySecKeyCreatePair(keychainRef SecKeychainRef, algorithm CSSM_ALGORITHMS, keySizeInBits uint32, contextHandle CSSM_CC_HANDLE, publicKeyUsage CSSM_KEYUSE, publicKeyAttr uint32, privateKeyUsage CSSM_KEYUSE, privateKeyAttr uint32, initialAccess SecAccessRef, publicKey *SecKeyRef, privateKey *SecKeyRef) (int32, error) {
	if _secKeyCreatePair == nil {
		return 0, symbolCallError("SecKeyCreatePair", "10.0", _secKeyCreatePairErr)
	}
	return _secKeyCreatePair(keychainRef, algorithm, keySizeInBits, contextHandle, publicKeyUsage, publicKeyAttr, privateKeyUsage, privateKeyAttr, initialAccess, publicKey, privateKey), nil
}

// SecKeyCreatePair creates an asymmetric key pair and stores it in a keychain.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecKeyCreatePair
func SecKeyCreatePair(keychainRef SecKeychainRef, algorithm CSSM_ALGORITHMS, keySizeInBits uint32, contextHandle CSSM_CC_HANDLE, publicKeyUsage CSSM_KEYUSE, publicKeyAttr uint32, privateKeyUsage CSSM_KEYUSE, privateKeyAttr uint32, initialAccess SecAccessRef, publicKey *SecKeyRef, privateKey *SecKeyRef) int32 {
	result, callErr := trySecKeyCreatePair(keychainRef, algorithm, keySizeInBits, contextHandle, publicKeyUsage, publicKeyAttr, privateKeyUsage, privateKeyAttr, initialAccess, publicKey, privateKey)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secKeyCreateRandomKey func(parameters corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) SecKeyRef
var _secKeyCreateRandomKeyErr error

func trySecKeyCreateRandomKey(parameters corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) (SecKeyRef, error) {
	if _secKeyCreateRandomKey == nil {
		return 0, symbolCallError("SecKeyCreateRandomKey", "10.12", _secKeyCreateRandomKeyErr)
	}
	return _secKeyCreateRandomKey(parameters, err), nil
}

// SecKeyCreateRandomKey generates a new public-private key pair.
//
// See: https://developer.apple.com/documentation/Security/SecKeyCreateRandomKey(_:_:)
func SecKeyCreateRandomKey(parameters corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) SecKeyRef {
	result, callErr := trySecKeyCreateRandomKey(parameters, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secKeyCreateSignature func(key SecKeyRef, algorithm SecKeyAlgorithm, dataToSign corefoundation.CFDataRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef
var _secKeyCreateSignatureErr error

func trySecKeyCreateSignature(key SecKeyRef, algorithm SecKeyAlgorithm, dataToSign corefoundation.CFDataRef, err *corefoundation.CFErrorRef) (corefoundation.CFDataRef, error) {
	if _secKeyCreateSignature == nil {
		return 0, symbolCallError("SecKeyCreateSignature", "10.12", _secKeyCreateSignatureErr)
	}
	return _secKeyCreateSignature(key, algorithm, dataToSign, err), nil
}

// SecKeyCreateSignature creates the cryptographic signature for a block of data using a private key and specified algorithm.
//
// See: https://developer.apple.com/documentation/Security/SecKeyCreateSignature(_:_:_:_:)
func SecKeyCreateSignature(key SecKeyRef, algorithm SecKeyAlgorithm, dataToSign corefoundation.CFDataRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef {
	result, callErr := trySecKeyCreateSignature(key, algorithm, dataToSign, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secKeyCreateWithData func(keyData corefoundation.CFDataRef, attributes corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) SecKeyRef
var _secKeyCreateWithDataErr error

func trySecKeyCreateWithData(keyData corefoundation.CFDataRef, attributes corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) (SecKeyRef, error) {
	if _secKeyCreateWithData == nil {
		return 0, symbolCallError("SecKeyCreateWithData", "10.12", _secKeyCreateWithDataErr)
	}
	return _secKeyCreateWithData(keyData, attributes, err), nil
}

// SecKeyCreateWithData restores a key from an external representation of that key.
//
// See: https://developer.apple.com/documentation/Security/SecKeyCreateWithData(_:_:_:)
func SecKeyCreateWithData(keyData corefoundation.CFDataRef, attributes corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) SecKeyRef {
	result, callErr := trySecKeyCreateWithData(keyData, attributes, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secKeyGenerate func(keychainRef SecKeychainRef, algorithm CSSM_ALGORITHMS, keySizeInBits uint32, contextHandle CSSM_CC_HANDLE, keyUsage CSSM_KEYUSE, keyAttr uint32, initialAccess SecAccessRef, keyRef *SecKeyRef) int32
var _secKeyGenerateErr error

func trySecKeyGenerate(keychainRef SecKeychainRef, algorithm CSSM_ALGORITHMS, keySizeInBits uint32, contextHandle CSSM_CC_HANDLE, keyUsage CSSM_KEYUSE, keyAttr uint32, initialAccess SecAccessRef, keyRef *SecKeyRef) (int32, error) {
	if _secKeyGenerate == nil {
		return 0, symbolCallError("SecKeyGenerate", "10.0", _secKeyGenerateErr)
	}
	return _secKeyGenerate(keychainRef, algorithm, keySizeInBits, contextHandle, keyUsage, keyAttr, initialAccess, keyRef), nil
}

// SecKeyGenerate creates a symmetric key and optionally stores it in a keychain.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecKeyGenerate
func SecKeyGenerate(keychainRef SecKeychainRef, algorithm CSSM_ALGORITHMS, keySizeInBits uint32, contextHandle CSSM_CC_HANDLE, keyUsage CSSM_KEYUSE, keyAttr uint32, initialAccess SecAccessRef, keyRef *SecKeyRef) int32 {
	result, callErr := trySecKeyGenerate(keychainRef, algorithm, keySizeInBits, contextHandle, keyUsage, keyAttr, initialAccess, keyRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secKeyGetBlockSize func(key SecKeyRef) uintptr
var _secKeyGetBlockSizeErr error

func trySecKeyGetBlockSize(key SecKeyRef) (uintptr, error) {
	if _secKeyGetBlockSize == nil {
		return 0, symbolCallError("SecKeyGetBlockSize", "10.6", _secKeyGetBlockSizeErr)
	}
	return _secKeyGetBlockSize(key), nil
}

// SecKeyGetBlockSize gets the block length associated with a cryptographic key.
//
// See: https://developer.apple.com/documentation/Security/SecKeyGetBlockSize(_:)
func SecKeyGetBlockSize(key SecKeyRef) uintptr {
	result, callErr := trySecKeyGetBlockSize(key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secKeyGetCSPHandle func(keyRef SecKeyRef, cspHandle unsafe.Pointer) int32
var _secKeyGetCSPHandleErr error

func trySecKeyGetCSPHandle(keyRef SecKeyRef, cspHandle unsafe.Pointer) (int32, error) {
	if _secKeyGetCSPHandle == nil {
		return 0, symbolCallError("SecKeyGetCSPHandle", "10.0", _secKeyGetCSPHandleErr)
	}
	return _secKeyGetCSPHandle(keyRef, cspHandle), nil
}

// SecKeyGetCSPHandle returns the CSSM CSP handle for a key.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecKeyGetCSPHandle
func SecKeyGetCSPHandle(keyRef SecKeyRef, cspHandle unsafe.Pointer) int32 {
	result, callErr := trySecKeyGetCSPHandle(keyRef, cspHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secKeyGetCSSMKey func(key SecKeyRef, cssmKey unsafe.Pointer) int32
var _secKeyGetCSSMKeyErr error

func trySecKeyGetCSSMKey(key SecKeyRef, cssmKey unsafe.Pointer) (int32, error) {
	if _secKeyGetCSSMKey == nil {
		return 0, symbolCallError("SecKeyGetCSSMKey", "10.0", _secKeyGetCSSMKeyErr)
	}
	return _secKeyGetCSSMKey(key, cssmKey), nil
}

// SecKeyGetCSSMKey retrieves a pointer to the `CSSM_KEY` structure containing the key stored in a keychain item.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecKeyGetCSSMKey
func SecKeyGetCSSMKey(key SecKeyRef, cssmKey unsafe.Pointer) int32 {
	result, callErr := trySecKeyGetCSSMKey(key, cssmKey)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secKeyGetCredentials func(keyRef SecKeyRef, operation CSSM_ACL_AUTHORIZATION_TAG, credentialType SecCredentialType, outCredentials unsafe.Pointer) int32
var _secKeyGetCredentialsErr error

func trySecKeyGetCredentials(keyRef SecKeyRef, operation CSSM_ACL_AUTHORIZATION_TAG, credentialType SecCredentialType, outCredentials unsafe.Pointer) (int32, error) {
	if _secKeyGetCredentials == nil {
		return 0, symbolCallError("SecKeyGetCredentials", "10.0", _secKeyGetCredentialsErr)
	}
	return _secKeyGetCredentials(keyRef, operation, credentialType, outCredentials), nil
}

// SecKeyGetCredentials returns an access credential for a key.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecKeyGetCredentials
func SecKeyGetCredentials(keyRef SecKeyRef, operation CSSM_ACL_AUTHORIZATION_TAG, credentialType SecCredentialType, outCredentials unsafe.Pointer) int32 {
	result, callErr := trySecKeyGetCredentials(keyRef, operation, credentialType, outCredentials)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secKeyGetTypeID func() uint
var _secKeyGetTypeIDErr error

func trySecKeyGetTypeID() (uint, error) {
	if _secKeyGetTypeID == nil {
		return 0, symbolCallError("SecKeyGetTypeID", "10.3", _secKeyGetTypeIDErr)
	}
	return _secKeyGetTypeID(), nil
}

// SecKeyGetTypeID returns the unique identifier of the opaque type to which a key object belongs.
//
// See: https://developer.apple.com/documentation/Security/SecKeyGetTypeID()
func SecKeyGetTypeID() uint {
	result, callErr := trySecKeyGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secKeyIsAlgorithmSupported func(key SecKeyRef, operation SecKeyOperationType, algorithm SecKeyAlgorithm) bool
var _secKeyIsAlgorithmSupportedErr error

func trySecKeyIsAlgorithmSupported(key SecKeyRef, operation SecKeyOperationType, algorithm SecKeyAlgorithm) (bool, error) {
	if _secKeyIsAlgorithmSupported == nil {
		return false, symbolCallError("SecKeyIsAlgorithmSupported", "10.12", _secKeyIsAlgorithmSupportedErr)
	}
	return _secKeyIsAlgorithmSupported(key, operation, algorithm), nil
}

// SecKeyIsAlgorithmSupported returns a Boolean indicating whether a key is suitable for an operation using a certain algorithm.
//
// See: https://developer.apple.com/documentation/Security/SecKeyIsAlgorithmSupported(_:_:_:)
func SecKeyIsAlgorithmSupported(key SecKeyRef, operation SecKeyOperationType, algorithm SecKeyAlgorithm) bool {
	result, callErr := trySecKeyIsAlgorithmSupported(key, operation, algorithm)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secKeyVerifySignature func(key SecKeyRef, algorithm SecKeyAlgorithm, signedData corefoundation.CFDataRef, signature corefoundation.CFDataRef, err *corefoundation.CFErrorRef) bool
var _secKeyVerifySignatureErr error

func trySecKeyVerifySignature(key SecKeyRef, algorithm SecKeyAlgorithm, signedData corefoundation.CFDataRef, signature corefoundation.CFDataRef, err *corefoundation.CFErrorRef) (bool, error) {
	if _secKeyVerifySignature == nil {
		return false, symbolCallError("SecKeyVerifySignature", "10.12", _secKeyVerifySignatureErr)
	}
	return _secKeyVerifySignature(key, algorithm, signedData, signature, err), nil
}

// SecKeyVerifySignature verifies the cryptographic signature of a block of data using a public key and specified algorithm.
//
// See: https://developer.apple.com/documentation/Security/SecKeyVerifySignature(_:_:_:_:_:)
func SecKeyVerifySignature(key SecKeyRef, algorithm SecKeyAlgorithm, signedData corefoundation.CFDataRef, signature corefoundation.CFDataRef, err *corefoundation.CFErrorRef) bool {
	result, callErr := trySecKeyVerifySignature(key, algorithm, signedData, signature, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secKeychainSearchGetTypeID func() uint
var _secKeychainSearchGetTypeIDErr error

func trySecKeychainSearchGetTypeID() (uint, error) {
	if _secKeychainSearchGetTypeID == nil {
		return 0, symbolCallError("SecKeychainSearchGetTypeID", "10.0", _secKeychainSearchGetTypeIDErr)
	}
	return _secKeychainSearchGetTypeID(), nil
}

// SecKeychainSearchGetTypeID returns the unique identifier of the opaque type to which a keychain search object belongs.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecKeychainSearchGetTypeID
func SecKeychainSearchGetTypeID() uint {
	result, callErr := trySecKeychainSearchGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secPKCS12Import func(pkcs12_data corefoundation.CFDataRef, options corefoundation.CFDictionaryRef, items *corefoundation.CFArrayRef) int32
var _secPKCS12ImportErr error

func trySecPKCS12Import(pkcs12_data corefoundation.CFDataRef, options corefoundation.CFDictionaryRef, items *corefoundation.CFArrayRef) (int32, error) {
	if _secPKCS12Import == nil {
		return 0, symbolCallError("SecPKCS12Import", "10.6", _secPKCS12ImportErr)
	}
	return _secPKCS12Import(pkcs12_data, options, items), nil
}

// SecPKCS12Import returns the identities and certificates in a PKCS #12-formatted blob.
//
// See: https://developer.apple.com/documentation/Security/SecPKCS12Import(_:_:_:)
func SecPKCS12Import(pkcs12_data corefoundation.CFDataRef, options corefoundation.CFDictionaryRef, items *corefoundation.CFArrayRef) int32 {
	result, callErr := trySecPKCS12Import(pkcs12_data, options, items)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secPolicyCopyProperties func(policyRef SecPolicyRef) corefoundation.CFDictionaryRef
var _secPolicyCopyPropertiesErr error

func trySecPolicyCopyProperties(policyRef SecPolicyRef) (corefoundation.CFDictionaryRef, error) {
	if _secPolicyCopyProperties == nil {
		return 0, symbolCallError("SecPolicyCopyProperties", "10.7", _secPolicyCopyPropertiesErr)
	}
	return _secPolicyCopyProperties(policyRef), nil
}

// SecPolicyCopyProperties returns a dictionary containing a policy’s properties.
//
// See: https://developer.apple.com/documentation/Security/SecPolicyCopyProperties(_:)
func SecPolicyCopyProperties(policyRef SecPolicyRef) corefoundation.CFDictionaryRef {
	result, callErr := trySecPolicyCopyProperties(policyRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secPolicyCreateBasicX509 func() SecPolicyRef
var _secPolicyCreateBasicX509Err error

func trySecPolicyCreateBasicX509() (SecPolicyRef, error) {
	if _secPolicyCreateBasicX509 == nil {
		return 0, symbolCallError("SecPolicyCreateBasicX509", "10.6", _secPolicyCreateBasicX509Err)
	}
	return _secPolicyCreateBasicX509(), nil
}

// SecPolicyCreateBasicX509 returns a policy object for the default X.509 policy.
//
// See: https://developer.apple.com/documentation/Security/SecPolicyCreateBasicX509()
func SecPolicyCreateBasicX509() SecPolicyRef {
	result, callErr := trySecPolicyCreateBasicX509()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secPolicyCreateRevocation func(revocationFlags uint64) SecPolicyRef
var _secPolicyCreateRevocationErr error

func trySecPolicyCreateRevocation(revocationFlags uint64) (SecPolicyRef, error) {
	if _secPolicyCreateRevocation == nil {
		return 0, symbolCallError("SecPolicyCreateRevocation", "10.9", _secPolicyCreateRevocationErr)
	}
	return _secPolicyCreateRevocation(revocationFlags), nil
}

// SecPolicyCreateRevocation returns a policy object for checking revocation of certificates.
//
// See: https://developer.apple.com/documentation/Security/SecPolicyCreateRevocation(_:)
func SecPolicyCreateRevocation(revocationFlags uint64) SecPolicyRef {
	result, callErr := trySecPolicyCreateRevocation(revocationFlags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secPolicyCreateSSL func(server bool, hostname corefoundation.CFStringRef) SecPolicyRef
var _secPolicyCreateSSLErr error

func trySecPolicyCreateSSL(server bool, hostname corefoundation.CFStringRef) (SecPolicyRef, error) {
	if _secPolicyCreateSSL == nil {
		return 0, symbolCallError("SecPolicyCreateSSL", "10.6", _secPolicyCreateSSLErr)
	}
	return _secPolicyCreateSSL(server, hostname), nil
}

// SecPolicyCreateSSL returns a policy object for evaluating SSL certificate chains.
//
// See: https://developer.apple.com/documentation/Security/SecPolicyCreateSSL(_:_:)
func SecPolicyCreateSSL(server bool, hostname corefoundation.CFStringRef) SecPolicyRef {
	result, callErr := trySecPolicyCreateSSL(server, hostname)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secPolicyCreateWithOID func(policyOID corefoundation.CFTypeRef) SecPolicyRef
var _secPolicyCreateWithOIDErr error

func trySecPolicyCreateWithOID(policyOID corefoundation.CFTypeRef) (SecPolicyRef, error) {
	if _secPolicyCreateWithOID == nil {
		return 0, symbolCallError("SecPolicyCreateWithOID", "10.7", _secPolicyCreateWithOIDErr)
	}
	return _secPolicyCreateWithOID(policyOID), nil
}

// SecPolicyCreateWithOID returns a policy object for the specified policy type object identifier.
//
// Deprecated: Deprecated since macOS 10.9.
//
// See: https://developer.apple.com/documentation/Security/SecPolicyCreateWithOID
func SecPolicyCreateWithOID(policyOID corefoundation.CFTypeRef) SecPolicyRef {
	result, callErr := trySecPolicyCreateWithOID(policyOID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secPolicyCreateWithProperties func(policyIdentifier corefoundation.CFTypeRef, properties corefoundation.CFDictionaryRef) SecPolicyRef
var _secPolicyCreateWithPropertiesErr error

func trySecPolicyCreateWithProperties(policyIdentifier corefoundation.CFTypeRef, properties corefoundation.CFDictionaryRef) (SecPolicyRef, error) {
	if _secPolicyCreateWithProperties == nil {
		return 0, symbolCallError("SecPolicyCreateWithProperties", "10.9", _secPolicyCreateWithPropertiesErr)
	}
	return _secPolicyCreateWithProperties(policyIdentifier, properties), nil
}

// SecPolicyCreateWithProperties returns a policy object based on an object identifier for the policy type.
//
// See: https://developer.apple.com/documentation/Security/SecPolicyCreateWithProperties(_:_:)
func SecPolicyCreateWithProperties(policyIdentifier corefoundation.CFTypeRef, properties corefoundation.CFDictionaryRef) SecPolicyRef {
	result, callErr := trySecPolicyCreateWithProperties(policyIdentifier, properties)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secPolicyGetOID func(policyRef SecPolicyRef, oid unsafe.Pointer) int32
var _secPolicyGetOIDErr error

func trySecPolicyGetOID(policyRef SecPolicyRef, oid unsafe.Pointer) (int32, error) {
	if _secPolicyGetOID == nil {
		return 0, symbolCallError("SecPolicyGetOID", "10.2", _secPolicyGetOIDErr)
	}
	return _secPolicyGetOID(policyRef, oid), nil
}

// SecPolicyGetOID retrieves a policy’s object identifier.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecPolicyGetOID
func SecPolicyGetOID(policyRef SecPolicyRef, oid unsafe.Pointer) int32 {
	result, callErr := trySecPolicyGetOID(policyRef, oid)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secPolicyGetTPHandle func(policyRef SecPolicyRef, tpHandle unsafe.Pointer) int32
var _secPolicyGetTPHandleErr error

func trySecPolicyGetTPHandle(policyRef SecPolicyRef, tpHandle unsafe.Pointer) (int32, error) {
	if _secPolicyGetTPHandle == nil {
		return 0, symbolCallError("SecPolicyGetTPHandle", "10.2", _secPolicyGetTPHandleErr)
	}
	return _secPolicyGetTPHandle(policyRef, tpHandle), nil
}

// SecPolicyGetTPHandle retrieves the trust policy handle for a policy object.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecPolicyGetTPHandle
func SecPolicyGetTPHandle(policyRef SecPolicyRef, tpHandle unsafe.Pointer) int32 {
	result, callErr := trySecPolicyGetTPHandle(policyRef, tpHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secPolicyGetTypeID func() uint
var _secPolicyGetTypeIDErr error

func trySecPolicyGetTypeID() (uint, error) {
	if _secPolicyGetTypeID == nil {
		return 0, symbolCallError("SecPolicyGetTypeID", "10.3", _secPolicyGetTypeIDErr)
	}
	return _secPolicyGetTypeID(), nil
}

// SecPolicyGetTypeID returns the unique identifier of the opaque type to which a policy object belongs.
//
// See: https://developer.apple.com/documentation/Security/SecPolicyGetTypeID()
func SecPolicyGetTypeID() uint {
	result, callErr := trySecPolicyGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secPolicyGetValue func(policyRef SecPolicyRef, value unsafe.Pointer) int32
var _secPolicyGetValueErr error

func trySecPolicyGetValue(policyRef SecPolicyRef, value unsafe.Pointer) (int32, error) {
	if _secPolicyGetValue == nil {
		return 0, symbolCallError("SecPolicyGetValue", "10.2", _secPolicyGetValueErr)
	}
	return _secPolicyGetValue(policyRef, value), nil
}

// SecPolicyGetValue retrieves a policy’s value.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecPolicyGetValue
func SecPolicyGetValue(policyRef SecPolicyRef, value unsafe.Pointer) int32 {
	result, callErr := trySecPolicyGetValue(policyRef, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secPolicySearchCopyNext func(searchRef SecPolicySearchRef, policyRef *SecPolicyRef) int32
var _secPolicySearchCopyNextErr error

func trySecPolicySearchCopyNext(searchRef SecPolicySearchRef, policyRef *SecPolicyRef) (int32, error) {
	if _secPolicySearchCopyNext == nil {
		return 0, symbolCallError("SecPolicySearchCopyNext", "10.0", _secPolicySearchCopyNextErr)
	}
	return _secPolicySearchCopyNext(searchRef, policyRef), nil
}

// SecPolicySearchCopyNext retrieves a policy object for the next policy matching specified search criteria.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecPolicySearchCopyNext
func SecPolicySearchCopyNext(searchRef SecPolicySearchRef, policyRef *SecPolicyRef) int32 {
	result, callErr := trySecPolicySearchCopyNext(searchRef, policyRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secPolicySearchCreate func(certType CSSM_CERT_TYPE, policyOID unsafe.Pointer, value unsafe.Pointer, searchRef *SecPolicySearchRef) int32
var _secPolicySearchCreateErr error

func trySecPolicySearchCreate(certType CSSM_CERT_TYPE, policyOID unsafe.Pointer, value unsafe.Pointer, searchRef *SecPolicySearchRef) (int32, error) {
	if _secPolicySearchCreate == nil {
		return 0, symbolCallError("SecPolicySearchCreate", "10.0", _secPolicySearchCreateErr)
	}
	return _secPolicySearchCreate(certType, policyOID, value, searchRef), nil
}

// SecPolicySearchCreate creates a search object for finding policies.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecPolicySearchCreate
func SecPolicySearchCreate(certType CSSM_CERT_TYPE, policyOID unsafe.Pointer, value unsafe.Pointer, searchRef *SecPolicySearchRef) int32 {
	result, callErr := trySecPolicySearchCreate(certType, policyOID, value, searchRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secPolicySearchGetTypeID func() uint
var _secPolicySearchGetTypeIDErr error

func trySecPolicySearchGetTypeID() (uint, error) {
	if _secPolicySearchGetTypeID == nil {
		return 0, symbolCallError("SecPolicySearchGetTypeID", "10.0", _secPolicySearchGetTypeIDErr)
	}
	return _secPolicySearchGetTypeID(), nil
}

// SecPolicySearchGetTypeID returns the unique identifier of the opaque type to which a [SecPolicySearch] object belongs.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecPolicySearchGetTypeID
func SecPolicySearchGetTypeID() uint {
	result, callErr := trySecPolicySearchGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secPolicySetProperties func(policyRef SecPolicyRef, properties corefoundation.CFDictionaryRef) int32
var _secPolicySetPropertiesErr error

func trySecPolicySetProperties(policyRef SecPolicyRef, properties corefoundation.CFDictionaryRef) (int32, error) {
	if _secPolicySetProperties == nil {
		return 0, symbolCallError("SecPolicySetProperties", "10.7", _secPolicySetPropertiesErr)
	}
	return _secPolicySetProperties(policyRef, properties), nil
}

// SecPolicySetProperties sets properties for a policy.
//
// Deprecated: Deprecated since macOS 10.9.
//
// See: https://developer.apple.com/documentation/Security/SecPolicySetProperties
func SecPolicySetProperties(policyRef SecPolicyRef, properties corefoundation.CFDictionaryRef) int32 {
	result, callErr := trySecPolicySetProperties(policyRef, properties)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secPolicySetValue func(policyRef SecPolicyRef, value unsafe.Pointer) int32
var _secPolicySetValueErr error

func trySecPolicySetValue(policyRef SecPolicyRef, value unsafe.Pointer) (int32, error) {
	if _secPolicySetValue == nil {
		return 0, symbolCallError("SecPolicySetValue", "10.2", _secPolicySetValueErr)
	}
	return _secPolicySetValue(policyRef, value), nil
}

// SecPolicySetValue sets a policy’s value.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecPolicySetValue
func SecPolicySetValue(policyRef SecPolicyRef, value unsafe.Pointer) int32 {
	result, callErr := trySecPolicySetValue(policyRef, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secRandomCopyBytes func(rnd SecRandomRef, count uintptr, bytes unsafe.Pointer) int
var _secRandomCopyBytesErr error

func trySecRandomCopyBytes(rnd SecRandomRef, count uintptr, bytes unsafe.Pointer) (int, error) {
	if _secRandomCopyBytes == nil {
		return 0, symbolCallError("SecRandomCopyBytes", "10.7", _secRandomCopyBytesErr)
	}
	return _secRandomCopyBytes(rnd, count, bytes), nil
}

// SecRandomCopyBytes generates an array of cryptographically secure random bytes.
//
// See: https://developer.apple.com/documentation/Security/SecRandomCopyBytes(_:_:_:)
func SecRandomCopyBytes(rnd SecRandomRef, count uintptr, bytes unsafe.Pointer) int {
	result, callErr := trySecRandomCopyBytes(rnd, count, bytes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secRequirementCopyData func(requirement SecRequirementRef, flags SecCSFlags, data *corefoundation.CFDataRef) int32
var _secRequirementCopyDataErr error

func trySecRequirementCopyData(requirement SecRequirementRef, flags SecCSFlags, data *corefoundation.CFDataRef) (int32, error) {
	if _secRequirementCopyData == nil {
		return 0, symbolCallError("SecRequirementCopyData", "10.0", _secRequirementCopyDataErr)
	}
	return _secRequirementCopyData(requirement, flags, data), nil
}

// SecRequirementCopyData extracts a binary form of a code requirement from a code requirement object.
//
// See: https://developer.apple.com/documentation/Security/SecRequirementCopyData(_:_:_:)
func SecRequirementCopyData(requirement SecRequirementRef, flags SecCSFlags, data *corefoundation.CFDataRef) int32 {
	result, callErr := trySecRequirementCopyData(requirement, flags, data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secRequirementCopyString func(requirement SecRequirementRef, flags SecCSFlags, text *corefoundation.CFStringRef) int32
var _secRequirementCopyStringErr error

func trySecRequirementCopyString(requirement SecRequirementRef, flags SecCSFlags, text *corefoundation.CFStringRef) (int32, error) {
	if _secRequirementCopyString == nil {
		return 0, symbolCallError("SecRequirementCopyString", "10.0", _secRequirementCopyStringErr)
	}
	return _secRequirementCopyString(requirement, flags, text), nil
}

// SecRequirementCopyString converts a code requirement object into text form.
//
// See: https://developer.apple.com/documentation/Security/SecRequirementCopyString(_:_:_:)
func SecRequirementCopyString(requirement SecRequirementRef, flags SecCSFlags, text *corefoundation.CFStringRef) int32 {
	result, callErr := trySecRequirementCopyString(requirement, flags, text)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secRequirementCreateWithData func(data corefoundation.CFDataRef, flags SecCSFlags, requirement *SecRequirementRef) int32
var _secRequirementCreateWithDataErr error

func trySecRequirementCreateWithData(data corefoundation.CFDataRef, flags SecCSFlags, requirement *SecRequirementRef) (int32, error) {
	if _secRequirementCreateWithData == nil {
		return 0, symbolCallError("SecRequirementCreateWithData", "10.0", _secRequirementCreateWithDataErr)
	}
	return _secRequirementCreateWithData(data, flags, requirement), nil
}

// SecRequirementCreateWithData creates a code requirement object from the binary form of a code requirement.
//
// See: https://developer.apple.com/documentation/Security/SecRequirementCreateWithData(_:_:_:)
func SecRequirementCreateWithData(data corefoundation.CFDataRef, flags SecCSFlags, requirement *SecRequirementRef) int32 {
	result, callErr := trySecRequirementCreateWithData(data, flags, requirement)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secRequirementCreateWithString func(text corefoundation.CFStringRef, flags SecCSFlags, requirement *SecRequirementRef) int32
var _secRequirementCreateWithStringErr error

func trySecRequirementCreateWithString(text corefoundation.CFStringRef, flags SecCSFlags, requirement *SecRequirementRef) (int32, error) {
	if _secRequirementCreateWithString == nil {
		return 0, symbolCallError("SecRequirementCreateWithString", "10.0", _secRequirementCreateWithStringErr)
	}
	return _secRequirementCreateWithString(text, flags, requirement), nil
}

// SecRequirementCreateWithString creates a code requirement object by compiling a valid text representation of a code requirement.
//
// See: https://developer.apple.com/documentation/Security/SecRequirementCreateWithString(_:_:_:)
func SecRequirementCreateWithString(text corefoundation.CFStringRef, flags SecCSFlags, requirement *SecRequirementRef) int32 {
	result, callErr := trySecRequirementCreateWithString(text, flags, requirement)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secRequirementCreateWithStringAndErrors func(text corefoundation.CFStringRef, flags SecCSFlags, errors *corefoundation.CFErrorRef, requirement *SecRequirementRef) int32
var _secRequirementCreateWithStringAndErrorsErr error

func trySecRequirementCreateWithStringAndErrors(text corefoundation.CFStringRef, flags SecCSFlags, errors *corefoundation.CFErrorRef, requirement *SecRequirementRef) (int32, error) {
	if _secRequirementCreateWithStringAndErrors == nil {
		return 0, symbolCallError("SecRequirementCreateWithStringAndErrors", "10.0", _secRequirementCreateWithStringAndErrorsErr)
	}
	return _secRequirementCreateWithStringAndErrors(text, flags, errors, requirement), nil
}

// SecRequirementCreateWithStringAndErrors creates a code requirement object by compiling a valid text representation of a code requirement and returns detailed error information in the case of failure.
//
// See: https://developer.apple.com/documentation/Security/SecRequirementCreateWithStringAndErrors(_:_:_:_:)
func SecRequirementCreateWithStringAndErrors(text corefoundation.CFStringRef, flags SecCSFlags, errors *corefoundation.CFErrorRef, requirement *SecRequirementRef) int32 {
	result, callErr := trySecRequirementCreateWithStringAndErrors(text, flags, errors, requirement)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secRequirementGetTypeID func() uint
var _secRequirementGetTypeIDErr error

func trySecRequirementGetTypeID() (uint, error) {
	if _secRequirementGetTypeID == nil {
		return 0, symbolCallError("SecRequirementGetTypeID", "10.0", _secRequirementGetTypeIDErr)
	}
	return _secRequirementGetTypeID(), nil
}

// SecRequirementGetTypeID returns the unique identifier of the opaque type to which a code requirement object belongs.
//
// See: https://developer.apple.com/documentation/Security/SecRequirementGetTypeID()
func SecRequirementGetTypeID() uint {
	result, callErr := trySecRequirementGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secStaticCodeCheckValidity func(staticCode SecStaticCodeRef, flags SecCSFlags, requirement SecRequirementRef) int32
var _secStaticCodeCheckValidityErr error

func trySecStaticCodeCheckValidity(staticCode SecStaticCodeRef, flags SecCSFlags, requirement SecRequirementRef) (int32, error) {
	if _secStaticCodeCheckValidity == nil {
		return 0, symbolCallError("SecStaticCodeCheckValidity", "10.0", _secStaticCodeCheckValidityErr)
	}
	return _secStaticCodeCheckValidity(staticCode, flags, requirement), nil
}

// SecStaticCodeCheckValidity validates a static code object.
//
// See: https://developer.apple.com/documentation/Security/SecStaticCodeCheckValidity(_:_:_:)
func SecStaticCodeCheckValidity(staticCode SecStaticCodeRef, flags SecCSFlags, requirement SecRequirementRef) int32 {
	result, callErr := trySecStaticCodeCheckValidity(staticCode, flags, requirement)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secStaticCodeCheckValidityWithErrors func(staticCode SecStaticCodeRef, flags SecCSFlags, requirement SecRequirementRef, errors *corefoundation.CFErrorRef) int32
var _secStaticCodeCheckValidityWithErrorsErr error

func trySecStaticCodeCheckValidityWithErrors(staticCode SecStaticCodeRef, flags SecCSFlags, requirement SecRequirementRef, errors *corefoundation.CFErrorRef) (int32, error) {
	if _secStaticCodeCheckValidityWithErrors == nil {
		return 0, symbolCallError("SecStaticCodeCheckValidityWithErrors", "10.0", _secStaticCodeCheckValidityWithErrorsErr)
	}
	return _secStaticCodeCheckValidityWithErrors(staticCode, flags, requirement, errors), nil
}

// SecStaticCodeCheckValidityWithErrors performs static validation of static signed code and returns detailed error information in the case of failure.
//
// See: https://developer.apple.com/documentation/Security/SecStaticCodeCheckValidityWithErrors(_:_:_:_:)
func SecStaticCodeCheckValidityWithErrors(staticCode SecStaticCodeRef, flags SecCSFlags, requirement SecRequirementRef, errors *corefoundation.CFErrorRef) int32 {
	result, callErr := trySecStaticCodeCheckValidityWithErrors(staticCode, flags, requirement, errors)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secStaticCodeCreateWithPath func(path corefoundation.CFURLRef, flags SecCSFlags, staticCode *SecStaticCodeRef) int32
var _secStaticCodeCreateWithPathErr error

func trySecStaticCodeCreateWithPath(path corefoundation.CFURLRef, flags SecCSFlags, staticCode *SecStaticCodeRef) (int32, error) {
	if _secStaticCodeCreateWithPath == nil {
		return 0, symbolCallError("SecStaticCodeCreateWithPath", "10.0", _secStaticCodeCreateWithPathErr)
	}
	return _secStaticCodeCreateWithPath(path, flags, staticCode), nil
}

// SecStaticCodeCreateWithPath creates a static code object representing the code at a specified file system path.
//
// See: https://developer.apple.com/documentation/Security/SecStaticCodeCreateWithPath(_:_:_:)
func SecStaticCodeCreateWithPath(path corefoundation.CFURLRef, flags SecCSFlags, staticCode *SecStaticCodeRef) int32 {
	result, callErr := trySecStaticCodeCreateWithPath(path, flags, staticCode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secStaticCodeCreateWithPathAndAttributes func(path corefoundation.CFURLRef, flags SecCSFlags, attributes corefoundation.CFDictionaryRef, staticCode *SecStaticCodeRef) int32
var _secStaticCodeCreateWithPathAndAttributesErr error

func trySecStaticCodeCreateWithPathAndAttributes(path corefoundation.CFURLRef, flags SecCSFlags, attributes corefoundation.CFDictionaryRef, staticCode *SecStaticCodeRef) (int32, error) {
	if _secStaticCodeCreateWithPathAndAttributes == nil {
		return 0, symbolCallError("SecStaticCodeCreateWithPathAndAttributes", "10.0", _secStaticCodeCreateWithPathAndAttributesErr)
	}
	return _secStaticCodeCreateWithPathAndAttributes(path, flags, attributes, staticCode), nil
}

// SecStaticCodeCreateWithPathAndAttributes creates a static code object representing the code at a specified file system path using an attributes dictionary.
//
// See: https://developer.apple.com/documentation/Security/SecStaticCodeCreateWithPathAndAttributes(_:_:_:_:)
func SecStaticCodeCreateWithPathAndAttributes(path corefoundation.CFURLRef, flags SecCSFlags, attributes corefoundation.CFDictionaryRef, staticCode *SecStaticCodeRef) int32 {
	result, callErr := trySecStaticCodeCreateWithPathAndAttributes(path, flags, attributes, staticCode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secStaticCodeGetTypeID func() uint
var _secStaticCodeGetTypeIDErr error

func trySecStaticCodeGetTypeID() (uint, error) {
	if _secStaticCodeGetTypeID == nil {
		return 0, symbolCallError("SecStaticCodeGetTypeID", "10.0", _secStaticCodeGetTypeIDErr)
	}
	return _secStaticCodeGetTypeID(), nil
}

// SecStaticCodeGetTypeID returns the unique identifier of the opaque type to which a static code object belongs.
//
// See: https://developer.apple.com/documentation/Security/SecStaticCodeGetTypeID()
func SecStaticCodeGetTypeID() uint {
	result, callErr := trySecStaticCodeGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTaskCopySigningIdentifier func(task SecTaskRef, err *corefoundation.CFErrorRef) corefoundation.CFStringRef
var _secTaskCopySigningIdentifierErr error

func trySecTaskCopySigningIdentifier(task SecTaskRef, err *corefoundation.CFErrorRef) (corefoundation.CFStringRef, error) {
	if _secTaskCopySigningIdentifier == nil {
		return 0, symbolCallError("SecTaskCopySigningIdentifier", "10.0", _secTaskCopySigningIdentifierErr)
	}
	return _secTaskCopySigningIdentifier(task, err), nil
}

// SecTaskCopySigningIdentifier returns the value of the code signing identifier.
//
// See: https://developer.apple.com/documentation/Security/SecTaskCopySigningIdentifier(_:_:)
func SecTaskCopySigningIdentifier(task SecTaskRef, err *corefoundation.CFErrorRef) corefoundation.CFStringRef {
	result, callErr := trySecTaskCopySigningIdentifier(task, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTaskCopyValueForEntitlement func(task SecTaskRef, entitlement corefoundation.CFStringRef, err *corefoundation.CFErrorRef) corefoundation.CFTypeRef
var _secTaskCopyValueForEntitlementErr error

func trySecTaskCopyValueForEntitlement(task SecTaskRef, entitlement corefoundation.CFStringRef, err *corefoundation.CFErrorRef) (corefoundation.CFTypeRef, error) {
	if _secTaskCopyValueForEntitlement == nil {
		return 0, symbolCallError("SecTaskCopyValueForEntitlement", "10.0", _secTaskCopyValueForEntitlementErr)
	}
	return _secTaskCopyValueForEntitlement(task, entitlement, err), nil
}

// SecTaskCopyValueForEntitlement returns the value of a single entitlement for the represented task.
//
// See: https://developer.apple.com/documentation/Security/SecTaskCopyValueForEntitlement(_:_:_:)
func SecTaskCopyValueForEntitlement(task SecTaskRef, entitlement corefoundation.CFStringRef, err *corefoundation.CFErrorRef) corefoundation.CFTypeRef {
	result, callErr := trySecTaskCopyValueForEntitlement(task, entitlement, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTaskCopyValuesForEntitlements func(task SecTaskRef, entitlements corefoundation.CFArrayRef, err *corefoundation.CFErrorRef) corefoundation.CFDictionaryRef
var _secTaskCopyValuesForEntitlementsErr error

func trySecTaskCopyValuesForEntitlements(task SecTaskRef, entitlements corefoundation.CFArrayRef, err *corefoundation.CFErrorRef) (corefoundation.CFDictionaryRef, error) {
	if _secTaskCopyValuesForEntitlements == nil {
		return 0, symbolCallError("SecTaskCopyValuesForEntitlements", "10.0", _secTaskCopyValuesForEntitlementsErr)
	}
	return _secTaskCopyValuesForEntitlements(task, entitlements, err), nil
}

// SecTaskCopyValuesForEntitlements returns the values of multiple entitlements for the represented task.
//
// See: https://developer.apple.com/documentation/Security/SecTaskCopyValuesForEntitlements(_:_:_:)
func SecTaskCopyValuesForEntitlements(task SecTaskRef, entitlements corefoundation.CFArrayRef, err *corefoundation.CFErrorRef) corefoundation.CFDictionaryRef {
	result, callErr := trySecTaskCopyValuesForEntitlements(task, entitlements, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTaskCreateFromSelf func(allocator corefoundation.CFAllocatorRef) SecTaskRef
var _secTaskCreateFromSelfErr error

func trySecTaskCreateFromSelf(allocator corefoundation.CFAllocatorRef) (SecTaskRef, error) {
	if _secTaskCreateFromSelf == nil {
		return 0, symbolCallError("SecTaskCreateFromSelf", "10.0", _secTaskCreateFromSelfErr)
	}
	return _secTaskCreateFromSelf(allocator), nil
}

// SecTaskCreateFromSelf creates a task object for the current task.
//
// See: https://developer.apple.com/documentation/Security/SecTaskCreateFromSelf(_:)
func SecTaskCreateFromSelf(allocator corefoundation.CFAllocatorRef) SecTaskRef {
	result, callErr := trySecTaskCreateFromSelf(allocator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTaskGetTypeID func() uint
var _secTaskGetTypeIDErr error

func trySecTaskGetTypeID() (uint, error) {
	if _secTaskGetTypeID == nil {
		return 0, symbolCallError("SecTaskGetTypeID", "10.0", _secTaskGetTypeIDErr)
	}
	return _secTaskGetTypeID(), nil
}

// SecTaskGetTypeID returns the unique identifier of the opaque type to which a task object belongs.
//
// See: https://developer.apple.com/documentation/Security/SecTaskGetTypeID()
func SecTaskGetTypeID() uint {
	result, callErr := trySecTaskGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustCopyAnchorCertificates func(anchors *corefoundation.CFArrayRef) int32
var _secTrustCopyAnchorCertificatesErr error

func trySecTrustCopyAnchorCertificates(anchors *corefoundation.CFArrayRef) (int32, error) {
	if _secTrustCopyAnchorCertificates == nil {
		return 0, symbolCallError("SecTrustCopyAnchorCertificates", "10.3", _secTrustCopyAnchorCertificatesErr)
	}
	return _secTrustCopyAnchorCertificates(anchors), nil
}

// SecTrustCopyAnchorCertificates retrieves the anchor (root) certificates stored by macOS.
//
// See: https://developer.apple.com/documentation/Security/SecTrustCopyAnchorCertificates(_:)
func SecTrustCopyAnchorCertificates(anchors *corefoundation.CFArrayRef) int32 {
	result, callErr := trySecTrustCopyAnchorCertificates(anchors)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustCopyCertificateChain func(trust SecTrustRef) corefoundation.CFArrayRef
var _secTrustCopyCertificateChainErr error

func trySecTrustCopyCertificateChain(trust SecTrustRef) (corefoundation.CFArrayRef, error) {
	if _secTrustCopyCertificateChain == nil {
		return 0, symbolCallError("SecTrustCopyCertificateChain", "12.0", _secTrustCopyCertificateChainErr)
	}
	return _secTrustCopyCertificateChain(trust), nil
}

// SecTrustCopyCertificateChain.
//
// See: https://developer.apple.com/documentation/Security/SecTrustCopyCertificateChain(_:)
func SecTrustCopyCertificateChain(trust SecTrustRef) corefoundation.CFArrayRef {
	result, callErr := trySecTrustCopyCertificateChain(trust)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustCopyCustomAnchorCertificates func(trust SecTrustRef, anchors *corefoundation.CFArrayRef) int32
var _secTrustCopyCustomAnchorCertificatesErr error

func trySecTrustCopyCustomAnchorCertificates(trust SecTrustRef, anchors *corefoundation.CFArrayRef) (int32, error) {
	if _secTrustCopyCustomAnchorCertificates == nil {
		return 0, symbolCallError("SecTrustCopyCustomAnchorCertificates", "10.5", _secTrustCopyCustomAnchorCertificatesErr)
	}
	return _secTrustCopyCustomAnchorCertificates(trust, anchors), nil
}

// SecTrustCopyCustomAnchorCertificates retrieves the custom anchor certificates, if any, used by a given trust.
//
// See: https://developer.apple.com/documentation/Security/SecTrustCopyCustomAnchorCertificates(_:_:)
func SecTrustCopyCustomAnchorCertificates(trust SecTrustRef, anchors *corefoundation.CFArrayRef) int32 {
	result, callErr := trySecTrustCopyCustomAnchorCertificates(trust, anchors)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustCopyExceptions func(trust SecTrustRef) corefoundation.CFDataRef
var _secTrustCopyExceptionsErr error

func trySecTrustCopyExceptions(trust SecTrustRef) (corefoundation.CFDataRef, error) {
	if _secTrustCopyExceptions == nil {
		return 0, symbolCallError("SecTrustCopyExceptions", "10.9", _secTrustCopyExceptionsErr)
	}
	return _secTrustCopyExceptions(trust), nil
}

// SecTrustCopyExceptions returns an opaque cookie containing exceptions to trust policies that will allow future evaluations of the current certificate to succeed.
//
// See: https://developer.apple.com/documentation/Security/SecTrustCopyExceptions(_:)
func SecTrustCopyExceptions(trust SecTrustRef) corefoundation.CFDataRef {
	result, callErr := trySecTrustCopyExceptions(trust)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustCopyKey func(trust SecTrustRef) SecKeyRef
var _secTrustCopyKeyErr error

func trySecTrustCopyKey(trust SecTrustRef) (SecKeyRef, error) {
	if _secTrustCopyKey == nil {
		return 0, symbolCallError("SecTrustCopyKey", "11.0", _secTrustCopyKeyErr)
	}
	return _secTrustCopyKey(trust), nil
}

// SecTrustCopyKey.
//
// See: https://developer.apple.com/documentation/Security/SecTrustCopyKey(_:)
func SecTrustCopyKey(trust SecTrustRef) SecKeyRef {
	result, callErr := trySecTrustCopyKey(trust)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustCopyPolicies func(trust SecTrustRef, policies *corefoundation.CFArrayRef) int32
var _secTrustCopyPoliciesErr error

func trySecTrustCopyPolicies(trust SecTrustRef, policies *corefoundation.CFArrayRef) (int32, error) {
	if _secTrustCopyPolicies == nil {
		return 0, symbolCallError("SecTrustCopyPolicies", "10.3", _secTrustCopyPoliciesErr)
	}
	return _secTrustCopyPolicies(trust, policies), nil
}

// SecTrustCopyPolicies retrieves the policies used by a given trust management object.
//
// See: https://developer.apple.com/documentation/Security/SecTrustCopyPolicies(_:_:)
func SecTrustCopyPolicies(trust SecTrustRef, policies *corefoundation.CFArrayRef) int32 {
	result, callErr := trySecTrustCopyPolicies(trust, policies)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustCopyProperties func(trust SecTrustRef) corefoundation.CFArrayRef
var _secTrustCopyPropertiesErr error

func trySecTrustCopyProperties(trust SecTrustRef) (corefoundation.CFArrayRef, error) {
	if _secTrustCopyProperties == nil {
		return 0, symbolCallError("SecTrustCopyProperties", "10.7", _secTrustCopyPropertiesErr)
	}
	return _secTrustCopyProperties(trust), nil
}

// SecTrustCopyProperties returns an array containing the properties of a trust object.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecTrustCopyProperties(_:)
func SecTrustCopyProperties(trust SecTrustRef) corefoundation.CFArrayRef {
	result, callErr := trySecTrustCopyProperties(trust)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustCopyPublicKey func(trust SecTrustRef) SecKeyRef
var _secTrustCopyPublicKeyErr error

func trySecTrustCopyPublicKey(trust SecTrustRef) (SecKeyRef, error) {
	if _secTrustCopyPublicKey == nil {
		return 0, symbolCallError("SecTrustCopyPublicKey", "10.7", _secTrustCopyPublicKeyErr)
	}
	return _secTrustCopyPublicKey(trust), nil
}

// SecTrustCopyPublicKey returns the public key for a leaf certificate after it has been evaluated.
//
// Deprecated: Deprecated since macOS 11.0.
//
// See: https://developer.apple.com/documentation/Security/SecTrustCopyPublicKey(_:)
func SecTrustCopyPublicKey(trust SecTrustRef) SecKeyRef {
	result, callErr := trySecTrustCopyPublicKey(trust)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustCopyResult func(trust SecTrustRef) corefoundation.CFDictionaryRef
var _secTrustCopyResultErr error

func trySecTrustCopyResult(trust SecTrustRef) (corefoundation.CFDictionaryRef, error) {
	if _secTrustCopyResult == nil {
		return 0, symbolCallError("SecTrustCopyResult", "10.9", _secTrustCopyResultErr)
	}
	return _secTrustCopyResult(trust), nil
}

// SecTrustCopyResult returns a dictionary containing information about an evaluated trust.
//
// See: https://developer.apple.com/documentation/Security/SecTrustCopyResult(_:)
func SecTrustCopyResult(trust SecTrustRef) corefoundation.CFDictionaryRef {
	result, callErr := trySecTrustCopyResult(trust)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustCreateWithCertificates func(certificates corefoundation.CFTypeRef, policies corefoundation.CFTypeRef, trust *SecTrustRef) int32
var _secTrustCreateWithCertificatesErr error

func trySecTrustCreateWithCertificates(certificates corefoundation.CFTypeRef, policies corefoundation.CFTypeRef, trust *SecTrustRef) (int32, error) {
	if _secTrustCreateWithCertificates == nil {
		return 0, symbolCallError("SecTrustCreateWithCertificates", "10.3", _secTrustCreateWithCertificatesErr)
	}
	return _secTrustCreateWithCertificates(certificates, policies, trust), nil
}

// SecTrustCreateWithCertificates creates a trust management object based on certificates and policies.
//
// See: https://developer.apple.com/documentation/Security/SecTrustCreateWithCertificates(_:_:_:)
func SecTrustCreateWithCertificates(certificates corefoundation.CFTypeRef, policies corefoundation.CFTypeRef, trust *SecTrustRef) int32 {
	result, callErr := trySecTrustCreateWithCertificates(certificates, policies, trust)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustEvaluateAsyncWithError func(trust SecTrustRef, queue uintptr, result unsafe.Pointer) int32
var _secTrustEvaluateAsyncWithErrorErr error

func trySecTrustEvaluateAsyncWithError(trust SecTrustRef, queue dispatch.Queue, result SecTrustWithErrorCallback) (int32, error) {
	if _secTrustEvaluateAsyncWithError == nil {
		return 0, symbolCallError("SecTrustEvaluateAsyncWithError", "10.15", _secTrustEvaluateAsyncWithErrorErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objectivec.IObject, blockArg1 bool, blockArg2 objectivec.IObject) {
		result(blockArg0, blockArg1, blockArg2)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _secTrustEvaluateAsyncWithError(trust, uintptr(queue.Handle()), _block0), nil
}

// SecTrustEvaluateAsyncWithError evaluates a trust object asynchronously on the specified dispatch queue.
//
// See: https://developer.apple.com/documentation/Security/SecTrustEvaluateAsyncWithError(_:_:_:)
func SecTrustEvaluateAsyncWithError(trust SecTrustRef, queue dispatch.Queue, result SecTrustWithErrorCallback) int32 {
	result0, callErr := trySecTrustEvaluateAsyncWithError(trust, queue, result)
	if callErr != nil {
		panic(callErr)
	}
	return result0
}

var _secTrustEvaluateWithError func(trust SecTrustRef, err *corefoundation.CFErrorRef) bool
var _secTrustEvaluateWithErrorErr error

func trySecTrustEvaluateWithError(trust SecTrustRef, err *corefoundation.CFErrorRef) (bool, error) {
	if _secTrustEvaluateWithError == nil {
		return false, symbolCallError("SecTrustEvaluateWithError", "10.14", _secTrustEvaluateWithErrorErr)
	}
	return _secTrustEvaluateWithError(trust, err), nil
}

// SecTrustEvaluateWithError evaluates trust for the specified certificate and policies.
//
// See: https://developer.apple.com/documentation/Security/SecTrustEvaluateWithError(_:_:)
func SecTrustEvaluateWithError(trust SecTrustRef, err *corefoundation.CFErrorRef) bool {
	result, callErr := trySecTrustEvaluateWithError(trust, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustGetCertificateAtIndex func(trust SecTrustRef, ix int) SecCertificateRef
var _secTrustGetCertificateAtIndexErr error

func trySecTrustGetCertificateAtIndex(trust SecTrustRef, ix int) (SecCertificateRef, error) {
	if _secTrustGetCertificateAtIndex == nil {
		return 0, symbolCallError("SecTrustGetCertificateAtIndex", "10.7", _secTrustGetCertificateAtIndexErr)
	}
	return _secTrustGetCertificateAtIndex(trust, ix), nil
}

// SecTrustGetCertificateAtIndex returns a specific certificate from the certificate chain used to evaluate trust.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecTrustGetCertificateAtIndex(_:_:)
func SecTrustGetCertificateAtIndex(trust SecTrustRef, ix int) SecCertificateRef {
	result, callErr := trySecTrustGetCertificateAtIndex(trust, ix)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustGetCertificateCount func(trust SecTrustRef) int
var _secTrustGetCertificateCountErr error

func trySecTrustGetCertificateCount(trust SecTrustRef) (int, error) {
	if _secTrustGetCertificateCount == nil {
		return 0, symbolCallError("SecTrustGetCertificateCount", "10.7", _secTrustGetCertificateCountErr)
	}
	return _secTrustGetCertificateCount(trust), nil
}

// SecTrustGetCertificateCount returns the number of certificates in an evaluated certificate chain.
//
// See: https://developer.apple.com/documentation/Security/SecTrustGetCertificateCount(_:)
func SecTrustGetCertificateCount(trust SecTrustRef) int {
	result, callErr := trySecTrustGetCertificateCount(trust)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustGetCssmResult func(trust SecTrustRef, result unsafe.Pointer) int32
var _secTrustGetCssmResultErr error

func trySecTrustGetCssmResult(trust SecTrustRef, result unsafe.Pointer) (int32, error) {
	if _secTrustGetCssmResult == nil {
		return 0, symbolCallError("SecTrustGetCssmResult", "10.2", _secTrustGetCssmResultErr)
	}
	return _secTrustGetCssmResult(trust, result), nil
}

// SecTrustGetCssmResult retrieves the CSSM trust result.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecTrustGetCssmResult
func SecTrustGetCssmResult(trust SecTrustRef, result unsafe.Pointer) int32 {
	result0, callErr := trySecTrustGetCssmResult(trust, result)
	if callErr != nil {
		panic(callErr)
	}
	return result0
}

var _secTrustGetCssmResultCode func(trust SecTrustRef, resultCode *int32) int32
var _secTrustGetCssmResultCodeErr error

func trySecTrustGetCssmResultCode(trust SecTrustRef, resultCode *int32) (int32, error) {
	if _secTrustGetCssmResultCode == nil {
		return 0, symbolCallError("SecTrustGetCssmResultCode", "10.2", _secTrustGetCssmResultCodeErr)
	}
	return _secTrustGetCssmResultCode(trust, resultCode), nil
}

// SecTrustGetCssmResultCode retrieves the CSSM result code from the most recent trust evaluation for a trust management object.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecTrustGetCssmResultCode
func SecTrustGetCssmResultCode(trust SecTrustRef, resultCode *int32) int32 {
	result, callErr := trySecTrustGetCssmResultCode(trust, resultCode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustGetNetworkFetchAllowed func(trust SecTrustRef, allowFetch *bool) int32
var _secTrustGetNetworkFetchAllowedErr error

func trySecTrustGetNetworkFetchAllowed(trust SecTrustRef, allowFetch *bool) (int32, error) {
	if _secTrustGetNetworkFetchAllowed == nil {
		return 0, symbolCallError("SecTrustGetNetworkFetchAllowed", "10.9", _secTrustGetNetworkFetchAllowedErr)
	}
	return _secTrustGetNetworkFetchAllowed(trust, allowFetch), nil
}

// SecTrustGetNetworkFetchAllowed indicates whether a trust evaluation is permitted to fetch missing intermediate certificates from the network.
//
// See: https://developer.apple.com/documentation/Security/SecTrustGetNetworkFetchAllowed(_:_:)
func SecTrustGetNetworkFetchAllowed(trust SecTrustRef, allowFetch *bool) int32 {
	result, callErr := trySecTrustGetNetworkFetchAllowed(trust, allowFetch)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustGetResult func(trustRef SecTrustRef, result *SecTrustResultType, certChain *corefoundation.CFArrayRef, statusChain unsafe.Pointer) int32
var _secTrustGetResultErr error

func trySecTrustGetResult(trustRef SecTrustRef, result *SecTrustResultType, certChain *corefoundation.CFArrayRef, statusChain unsafe.Pointer) (int32, error) {
	if _secTrustGetResult == nil {
		return 0, symbolCallError("SecTrustGetResult", "10.2", _secTrustGetResultErr)
	}
	return _secTrustGetResult(trustRef, result, certChain, statusChain), nil
}

// SecTrustGetResult retrieves details on the outcome of a call to the function [SecTrustEvaluate].
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecTrustGetResult
func SecTrustGetResult(trustRef SecTrustRef, result *SecTrustResultType, certChain *corefoundation.CFArrayRef, statusChain unsafe.Pointer) int32 {
	result0, callErr := trySecTrustGetResult(trustRef, result, certChain, statusChain)
	if callErr != nil {
		panic(callErr)
	}
	return result0
}

var _secTrustGetTPHandle func(trust SecTrustRef, handle unsafe.Pointer) int32
var _secTrustGetTPHandleErr error

func trySecTrustGetTPHandle(trust SecTrustRef, handle unsafe.Pointer) (int32, error) {
	if _secTrustGetTPHandle == nil {
		return 0, symbolCallError("SecTrustGetTPHandle", "10.2", _secTrustGetTPHandleErr)
	}
	return _secTrustGetTPHandle(trust, handle), nil
}

// SecTrustGetTPHandle retrieves the trust policy handle.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecTrustGetTPHandle
func SecTrustGetTPHandle(trust SecTrustRef, handle unsafe.Pointer) int32 {
	result, callErr := trySecTrustGetTPHandle(trust, handle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustGetTrustResult func(trust SecTrustRef, result *SecTrustResultType) int32
var _secTrustGetTrustResultErr error

func trySecTrustGetTrustResult(trust SecTrustRef, result *SecTrustResultType) (int32, error) {
	if _secTrustGetTrustResult == nil {
		return 0, symbolCallError("SecTrustGetTrustResult", "10.7", _secTrustGetTrustResultErr)
	}
	return _secTrustGetTrustResult(trust, result), nil
}

// SecTrustGetTrustResult returns the result code from the most recent trust evaluation.
//
// See: https://developer.apple.com/documentation/Security/SecTrustGetTrustResult(_:_:)
func SecTrustGetTrustResult(trust SecTrustRef, result *SecTrustResultType) int32 {
	result0, callErr := trySecTrustGetTrustResult(trust, result)
	if callErr != nil {
		panic(callErr)
	}
	return result0
}

var _secTrustGetTypeID func() uint
var _secTrustGetTypeIDErr error

func trySecTrustGetTypeID() (uint, error) {
	if _secTrustGetTypeID == nil {
		return 0, symbolCallError("SecTrustGetTypeID", "10.3", _secTrustGetTypeIDErr)
	}
	return _secTrustGetTypeID(), nil
}

// SecTrustGetTypeID returns the unique identifier of the opaque type to which a trust object belongs.
//
// See: https://developer.apple.com/documentation/Security/SecTrustGetTypeID()
func SecTrustGetTypeID() uint {
	result, callErr := trySecTrustGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustGetVerifyTime func(trust SecTrustRef) corefoundation.CFAbsoluteTime
var _secTrustGetVerifyTimeErr error

func trySecTrustGetVerifyTime(trust SecTrustRef) (corefoundation.CFAbsoluteTime, error) {
	if _secTrustGetVerifyTime == nil {
		return *new(corefoundation.CFAbsoluteTime), symbolCallError("SecTrustGetVerifyTime", "10.6", _secTrustGetVerifyTimeErr)
	}
	return _secTrustGetVerifyTime(trust), nil
}

// SecTrustGetVerifyTime gets the absolute time against which the certificates in a trust management object are verified.
//
// See: https://developer.apple.com/documentation/Security/SecTrustGetVerifyTime(_:)
func SecTrustGetVerifyTime(trust SecTrustRef) corefoundation.CFAbsoluteTime {
	result, callErr := trySecTrustGetVerifyTime(trust)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustSetAnchorCertificates func(trust SecTrustRef, anchorCertificates corefoundation.CFArrayRef) int32
var _secTrustSetAnchorCertificatesErr error

func trySecTrustSetAnchorCertificates(trust SecTrustRef, anchorCertificates corefoundation.CFArrayRef) (int32, error) {
	if _secTrustSetAnchorCertificates == nil {
		return 0, symbolCallError("SecTrustSetAnchorCertificates", "10.3", _secTrustSetAnchorCertificatesErr)
	}
	return _secTrustSetAnchorCertificates(trust, anchorCertificates), nil
}

// SecTrustSetAnchorCertificates sets the anchor certificates used when evaluating a trust management object.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSetAnchorCertificates(_:_:)
func SecTrustSetAnchorCertificates(trust SecTrustRef, anchorCertificates corefoundation.CFArrayRef) int32 {
	result, callErr := trySecTrustSetAnchorCertificates(trust, anchorCertificates)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustSetAnchorCertificatesOnly func(trust SecTrustRef, anchorCertificatesOnly bool) int32
var _secTrustSetAnchorCertificatesOnlyErr error

func trySecTrustSetAnchorCertificatesOnly(trust SecTrustRef, anchorCertificatesOnly bool) (int32, error) {
	if _secTrustSetAnchorCertificatesOnly == nil {
		return 0, symbolCallError("SecTrustSetAnchorCertificatesOnly", "10.6", _secTrustSetAnchorCertificatesOnlyErr)
	}
	return _secTrustSetAnchorCertificatesOnly(trust, anchorCertificatesOnly), nil
}

// SecTrustSetAnchorCertificatesOnly reenables trusting built-in anchor certificates.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSetAnchorCertificatesOnly(_:_:)
func SecTrustSetAnchorCertificatesOnly(trust SecTrustRef, anchorCertificatesOnly bool) int32 {
	result, callErr := trySecTrustSetAnchorCertificatesOnly(trust, anchorCertificatesOnly)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustSetExceptions func(trust SecTrustRef, exceptions corefoundation.CFDataRef) bool
var _secTrustSetExceptionsErr error

func trySecTrustSetExceptions(trust SecTrustRef, exceptions corefoundation.CFDataRef) (bool, error) {
	if _secTrustSetExceptions == nil {
		return false, symbolCallError("SecTrustSetExceptions", "10.9", _secTrustSetExceptionsErr)
	}
	return _secTrustSetExceptions(trust, exceptions), nil
}

// SecTrustSetExceptions sets a list of exceptions that should be ignored when the certificate is evaluated.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSetExceptions(_:_:)
func SecTrustSetExceptions(trust SecTrustRef, exceptions corefoundation.CFDataRef) bool {
	result, callErr := trySecTrustSetExceptions(trust, exceptions)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustSetKeychains func(trust SecTrustRef, keychainOrArray corefoundation.CFTypeRef) int32
var _secTrustSetKeychainsErr error

func trySecTrustSetKeychains(trust SecTrustRef, keychainOrArray corefoundation.CFTypeRef) (int32, error) {
	if _secTrustSetKeychains == nil {
		return 0, symbolCallError("SecTrustSetKeychains", "10.3", _secTrustSetKeychainsErr)
	}
	return _secTrustSetKeychains(trust, keychainOrArray), nil
}

// SecTrustSetKeychains sets the keychains searched for intermediate certificates when evaluating a trust management object.
//
// Deprecated: Deprecated since macOS 10.13.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSetKeychains(_:_:)
func SecTrustSetKeychains(trust SecTrustRef, keychainOrArray corefoundation.CFTypeRef) int32 {
	result, callErr := trySecTrustSetKeychains(trust, keychainOrArray)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustSetNetworkFetchAllowed func(trust SecTrustRef, allowFetch bool) int32
var _secTrustSetNetworkFetchAllowedErr error

func trySecTrustSetNetworkFetchAllowed(trust SecTrustRef, allowFetch bool) (int32, error) {
	if _secTrustSetNetworkFetchAllowed == nil {
		return 0, symbolCallError("SecTrustSetNetworkFetchAllowed", "10.9", _secTrustSetNetworkFetchAllowedErr)
	}
	return _secTrustSetNetworkFetchAllowed(trust, allowFetch), nil
}

// SecTrustSetNetworkFetchAllowed specifies whether a trust evaluation is permitted to fetch missing intermediate certificates from the network.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSetNetworkFetchAllowed(_:_:)
func SecTrustSetNetworkFetchAllowed(trust SecTrustRef, allowFetch bool) int32 {
	result, callErr := trySecTrustSetNetworkFetchAllowed(trust, allowFetch)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustSetOCSPResponse func(trust SecTrustRef, responseData corefoundation.CFTypeRef) int32
var _secTrustSetOCSPResponseErr error

func trySecTrustSetOCSPResponse(trust SecTrustRef, responseData corefoundation.CFTypeRef) (int32, error) {
	if _secTrustSetOCSPResponse == nil {
		return 0, symbolCallError("SecTrustSetOCSPResponse", "10.9", _secTrustSetOCSPResponseErr)
	}
	return _secTrustSetOCSPResponse(trust, responseData), nil
}

// SecTrustSetOCSPResponse attaches Online Certificate Status Protocol (OSCP) response data to a trust object.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSetOCSPResponse(_:_:)
func SecTrustSetOCSPResponse(trust SecTrustRef, responseData corefoundation.CFTypeRef) int32 {
	result, callErr := trySecTrustSetOCSPResponse(trust, responseData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustSetOptions func(trustRef SecTrustRef, options SecTrustOptionFlags) int32
var _secTrustSetOptionsErr error

func trySecTrustSetOptions(trustRef SecTrustRef, options SecTrustOptionFlags) (int32, error) {
	if _secTrustSetOptions == nil {
		return 0, symbolCallError("SecTrustSetOptions", "10.7", _secTrustSetOptionsErr)
	}
	return _secTrustSetOptions(trustRef, options), nil
}

// SecTrustSetOptions sets option flags for customizing evaluation of a trust object.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSetOptions(_:_:)
func SecTrustSetOptions(trustRef SecTrustRef, options SecTrustOptionFlags) int32 {
	result, callErr := trySecTrustSetOptions(trustRef, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustSetParameters func(trustRef SecTrustRef, action CSSM_TP_ACTION, actionData corefoundation.CFDataRef) int32
var _secTrustSetParametersErr error

func trySecTrustSetParameters(trustRef SecTrustRef, action CSSM_TP_ACTION, actionData corefoundation.CFDataRef) (int32, error) {
	if _secTrustSetParameters == nil {
		return 0, symbolCallError("SecTrustSetParameters", "10.2", _secTrustSetParametersErr)
	}
	return _secTrustSetParameters(trustRef, action, actionData), nil
}

// SecTrustSetParameters sets the action and action data for a trust management object.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSetParameters
func SecTrustSetParameters(trustRef SecTrustRef, action CSSM_TP_ACTION, actionData corefoundation.CFDataRef) int32 {
	result, callErr := trySecTrustSetParameters(trustRef, action, actionData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustSetPolicies func(trust SecTrustRef, policies corefoundation.CFTypeRef) int32
var _secTrustSetPoliciesErr error

func trySecTrustSetPolicies(trust SecTrustRef, policies corefoundation.CFTypeRef) (int32, error) {
	if _secTrustSetPolicies == nil {
		return 0, symbolCallError("SecTrustSetPolicies", "10.3", _secTrustSetPoliciesErr)
	}
	return _secTrustSetPolicies(trust, policies), nil
}

// SecTrustSetPolicies sets the policies to use in an evaluation.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSetPolicies(_:_:)
func SecTrustSetPolicies(trust SecTrustRef, policies corefoundation.CFTypeRef) int32 {
	result, callErr := trySecTrustSetPolicies(trust, policies)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustSetSignedCertificateTimestamps func(trust SecTrustRef, sctArray corefoundation.CFArrayRef) int32
var _secTrustSetSignedCertificateTimestampsErr error

func trySecTrustSetSignedCertificateTimestamps(trust SecTrustRef, sctArray corefoundation.CFArrayRef) (int32, error) {
	if _secTrustSetSignedCertificateTimestamps == nil {
		return 0, symbolCallError("SecTrustSetSignedCertificateTimestamps", "10.14.2", _secTrustSetSignedCertificateTimestampsErr)
	}
	return _secTrustSetSignedCertificateTimestamps(trust, sctArray), nil
}

// SecTrustSetSignedCertificateTimestamps attaches signed certificate timestamp data to a trust object.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSetSignedCertificateTimestamps(_:_:)
func SecTrustSetSignedCertificateTimestamps(trust SecTrustRef, sctArray corefoundation.CFArrayRef) int32 {
	result, callErr := trySecTrustSetSignedCertificateTimestamps(trust, sctArray)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustSetVerifyDate func(trust SecTrustRef, verifyDate corefoundation.CFDateRef) int32
var _secTrustSetVerifyDateErr error

func trySecTrustSetVerifyDate(trust SecTrustRef, verifyDate corefoundation.CFDateRef) (int32, error) {
	if _secTrustSetVerifyDate == nil {
		return 0, symbolCallError("SecTrustSetVerifyDate", "10.3", _secTrustSetVerifyDateErr)
	}
	return _secTrustSetVerifyDate(trust, verifyDate), nil
}

// SecTrustSetVerifyDate sets the date and time against which the certificates in a trust management object are verified.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSetVerifyDate(_:_:)
func SecTrustSetVerifyDate(trust SecTrustRef, verifyDate corefoundation.CFDateRef) int32 {
	result, callErr := trySecTrustSetVerifyDate(trust, verifyDate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustSettingsCopyCertificates func(domain SecTrustSettingsDomain, certArray *corefoundation.CFArrayRef) int32
var _secTrustSettingsCopyCertificatesErr error

func trySecTrustSettingsCopyCertificates(domain SecTrustSettingsDomain, certArray *corefoundation.CFArrayRef) (int32, error) {
	if _secTrustSettingsCopyCertificates == nil {
		return 0, symbolCallError("SecTrustSettingsCopyCertificates", "10.0", _secTrustSettingsCopyCertificatesErr)
	}
	return _secTrustSettingsCopyCertificates(domain, certArray), nil
}

// SecTrustSettingsCopyCertificates obtains an array of all certificates that have trust settings in a specific trust settings domain.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSettingsCopyCertificates(_:_:)
func SecTrustSettingsCopyCertificates(domain SecTrustSettingsDomain, certArray *corefoundation.CFArrayRef) int32 {
	result, callErr := trySecTrustSettingsCopyCertificates(domain, certArray)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustSettingsCopyModificationDate func(certRef SecCertificateRef, domain SecTrustSettingsDomain, modificationDate *corefoundation.CFDateRef) int32
var _secTrustSettingsCopyModificationDateErr error

func trySecTrustSettingsCopyModificationDate(certRef SecCertificateRef, domain SecTrustSettingsDomain, modificationDate *corefoundation.CFDateRef) (int32, error) {
	if _secTrustSettingsCopyModificationDate == nil {
		return 0, symbolCallError("SecTrustSettingsCopyModificationDate", "10.0", _secTrustSettingsCopyModificationDateErr)
	}
	return _secTrustSettingsCopyModificationDate(certRef, domain, modificationDate), nil
}

// SecTrustSettingsCopyModificationDate obtains the date and time at which a certificate’s trust settings were last modified.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSettingsCopyModificationDate(_:_:_:)
func SecTrustSettingsCopyModificationDate(certRef SecCertificateRef, domain SecTrustSettingsDomain, modificationDate *corefoundation.CFDateRef) int32 {
	result, callErr := trySecTrustSettingsCopyModificationDate(certRef, domain, modificationDate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustSettingsCopyTrustSettings func(certRef SecCertificateRef, domain SecTrustSettingsDomain, trustSettings *corefoundation.CFArrayRef) int32
var _secTrustSettingsCopyTrustSettingsErr error

func trySecTrustSettingsCopyTrustSettings(certRef SecCertificateRef, domain SecTrustSettingsDomain, trustSettings *corefoundation.CFArrayRef) (int32, error) {
	if _secTrustSettingsCopyTrustSettings == nil {
		return 0, symbolCallError("SecTrustSettingsCopyTrustSettings", "10.0", _secTrustSettingsCopyTrustSettingsErr)
	}
	return _secTrustSettingsCopyTrustSettings(certRef, domain, trustSettings), nil
}

// SecTrustSettingsCopyTrustSettings obtains the trust settings for a certificate.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSettingsCopyTrustSettings(_:_:_:)
func SecTrustSettingsCopyTrustSettings(certRef SecCertificateRef, domain SecTrustSettingsDomain, trustSettings *corefoundation.CFArrayRef) int32 {
	result, callErr := trySecTrustSettingsCopyTrustSettings(certRef, domain, trustSettings)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustSettingsCreateExternalRepresentation func(domain SecTrustSettingsDomain, trustSettings *corefoundation.CFDataRef) int32
var _secTrustSettingsCreateExternalRepresentationErr error

func trySecTrustSettingsCreateExternalRepresentation(domain SecTrustSettingsDomain, trustSettings *corefoundation.CFDataRef) (int32, error) {
	if _secTrustSettingsCreateExternalRepresentation == nil {
		return 0, symbolCallError("SecTrustSettingsCreateExternalRepresentation", "10.0", _secTrustSettingsCreateExternalRepresentationErr)
	}
	return _secTrustSettingsCreateExternalRepresentation(domain, trustSettings), nil
}

// SecTrustSettingsCreateExternalRepresentation obtains an external, portable representation of the specified domain’s trust settings.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSettingsCreateExternalRepresentation(_:_:)
func SecTrustSettingsCreateExternalRepresentation(domain SecTrustSettingsDomain, trustSettings *corefoundation.CFDataRef) int32 {
	result, callErr := trySecTrustSettingsCreateExternalRepresentation(domain, trustSettings)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustSettingsImportExternalRepresentation func(domain SecTrustSettingsDomain, trustSettings corefoundation.CFDataRef) int32
var _secTrustSettingsImportExternalRepresentationErr error

func trySecTrustSettingsImportExternalRepresentation(domain SecTrustSettingsDomain, trustSettings corefoundation.CFDataRef) (int32, error) {
	if _secTrustSettingsImportExternalRepresentation == nil {
		return 0, symbolCallError("SecTrustSettingsImportExternalRepresentation", "10.0", _secTrustSettingsImportExternalRepresentationErr)
	}
	return _secTrustSettingsImportExternalRepresentation(domain, trustSettings), nil
}

// SecTrustSettingsImportExternalRepresentation imports trust settings into a trust domain.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSettingsImportExternalRepresentation(_:_:)
func SecTrustSettingsImportExternalRepresentation(domain SecTrustSettingsDomain, trustSettings corefoundation.CFDataRef) int32 {
	result, callErr := trySecTrustSettingsImportExternalRepresentation(domain, trustSettings)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustSettingsRemoveTrustSettings func(certRef SecCertificateRef, domain SecTrustSettingsDomain) int32
var _secTrustSettingsRemoveTrustSettingsErr error

func trySecTrustSettingsRemoveTrustSettings(certRef SecCertificateRef, domain SecTrustSettingsDomain) (int32, error) {
	if _secTrustSettingsRemoveTrustSettings == nil {
		return 0, symbolCallError("SecTrustSettingsRemoveTrustSettings", "10.0", _secTrustSettingsRemoveTrustSettingsErr)
	}
	return _secTrustSettingsRemoveTrustSettings(certRef, domain), nil
}

// SecTrustSettingsRemoveTrustSettings deletes the trust settings for a certificate.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSettingsRemoveTrustSettings(_:_:)
func SecTrustSettingsRemoveTrustSettings(certRef SecCertificateRef, domain SecTrustSettingsDomain) int32 {
	result, callErr := trySecTrustSettingsRemoveTrustSettings(certRef, domain)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustSettingsSetTrustSettings func(certRef SecCertificateRef, domain SecTrustSettingsDomain, trustSettingsDictOrArray corefoundation.CFTypeRef) int32
var _secTrustSettingsSetTrustSettingsErr error

func trySecTrustSettingsSetTrustSettings(certRef SecCertificateRef, domain SecTrustSettingsDomain, trustSettingsDictOrArray corefoundation.CFTypeRef) (int32, error) {
	if _secTrustSettingsSetTrustSettings == nil {
		return 0, symbolCallError("SecTrustSettingsSetTrustSettings", "10.0", _secTrustSettingsSetTrustSettingsErr)
	}
	return _secTrustSettingsSetTrustSettings(certRef, domain, trustSettingsDictOrArray), nil
}

// SecTrustSettingsSetTrustSettings specifies trust settings for a certificate.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSettingsSetTrustSettings(_:_:_:)
func SecTrustSettingsSetTrustSettings(certRef SecCertificateRef, domain SecTrustSettingsDomain, trustSettingsDictOrArray corefoundation.CFTypeRef) int32 {
	result, callErr := trySecTrustSettingsSetTrustSettings(certRef, domain, trustSettingsDictOrArray)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secureDownloadCopyCreationDate func(downloadRef SecureDownloadRef, date *corefoundation.CFDateRef) int32
var _secureDownloadCopyCreationDateErr error

func trySecureDownloadCopyCreationDate(downloadRef SecureDownloadRef, date *corefoundation.CFDateRef) (int32, error) {
	if _secureDownloadCopyCreationDate == nil {
		return 0, symbolCallError("SecureDownloadCopyCreationDate", "10.5", _secureDownloadCopyCreationDateErr)
	}
	return _secureDownloadCopyCreationDate(downloadRef, date), nil
}

// SecureDownloadCopyCreationDate returns download ticket’s creation date.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecureDownloadCopyCreationDate
func SecureDownloadCopyCreationDate(downloadRef SecureDownloadRef, date *corefoundation.CFDateRef) int32 {
	result, callErr := trySecureDownloadCopyCreationDate(downloadRef, date)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secureDownloadCopyName func(downloadRef SecureDownloadRef, name *corefoundation.CFStringRef) int32
var _secureDownloadCopyNameErr error

func trySecureDownloadCopyName(downloadRef SecureDownloadRef, name *corefoundation.CFStringRef) (int32, error) {
	if _secureDownloadCopyName == nil {
		return 0, symbolCallError("SecureDownloadCopyName", "10.5", _secureDownloadCopyNameErr)
	}
	return _secureDownloadCopyName(downloadRef, name), nil
}

// SecureDownloadCopyName returns the printable name of the download ticket.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecureDownloadCopyName
func SecureDownloadCopyName(downloadRef SecureDownloadRef, name *corefoundation.CFStringRef) int32 {
	result, callErr := trySecureDownloadCopyName(downloadRef, name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secureDownloadCopyTicketLocation func(url corefoundation.CFURLRef, ticketLocation *corefoundation.CFURLRef) int32
var _secureDownloadCopyTicketLocationErr error

func trySecureDownloadCopyTicketLocation(url corefoundation.CFURLRef, ticketLocation *corefoundation.CFURLRef) (int32, error) {
	if _secureDownloadCopyTicketLocation == nil {
		return 0, symbolCallError("SecureDownloadCopyTicketLocation", "10.5", _secureDownloadCopyTicketLocationErr)
	}
	return _secureDownloadCopyTicketLocation(url, ticketLocation), nil
}

// SecureDownloadCopyTicketLocation copies the ticket location from a secure download URL.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecureDownloadCopyTicketLocation
func SecureDownloadCopyTicketLocation(url corefoundation.CFURLRef, ticketLocation *corefoundation.CFURLRef) int32 {
	result, callErr := trySecureDownloadCopyTicketLocation(url, ticketLocation)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secureDownloadCopyURLs func(downloadRef SecureDownloadRef, urls *corefoundation.CFArrayRef) int32
var _secureDownloadCopyURLsErr error

func trySecureDownloadCopyURLs(downloadRef SecureDownloadRef, urls *corefoundation.CFArrayRef) (int32, error) {
	if _secureDownloadCopyURLs == nil {
		return 0, symbolCallError("SecureDownloadCopyURLs", "10.5", _secureDownloadCopyURLsErr)
	}
	return _secureDownloadCopyURLs(downloadRef, urls), nil
}

// SecureDownloadCopyURLs returns a list of URLs from which the data can be downloaded.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecureDownloadCopyURLs
func SecureDownloadCopyURLs(downloadRef SecureDownloadRef, urls *corefoundation.CFArrayRef) int32 {
	result, callErr := trySecureDownloadCopyURLs(downloadRef, urls)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secureDownloadCreateWithTicket func(ticket corefoundation.CFDataRef, setup unsafe.Pointer, setupContext unsafe.Pointer, evaluate unsafe.Pointer, evaluateContext unsafe.Pointer, downloadRef *SecureDownloadRef) int32
var _secureDownloadCreateWithTicketErr error

func trySecureDownloadCreateWithTicket(ticket corefoundation.CFDataRef, setup unsafe.Pointer, setupContext unsafe.Pointer, evaluate unsafe.Pointer, evaluateContext unsafe.Pointer, downloadRef *SecureDownloadRef) (int32, error) {
	if _secureDownloadCreateWithTicket == nil {
		return 0, symbolCallError("SecureDownloadCreateWithTicket", "10.5", _secureDownloadCreateWithTicketErr)
	}
	return _secureDownloadCreateWithTicket(ticket, setup, setupContext, evaluate, evaluateContext, downloadRef), nil
}

// SecureDownloadCreateWithTicket creates a secure download object for use during the download process.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecureDownloadCreateWithTicket
func SecureDownloadCreateWithTicket(ticket corefoundation.CFDataRef, setup unsafe.Pointer, setupContext unsafe.Pointer, evaluate unsafe.Pointer, evaluateContext unsafe.Pointer, downloadRef *SecureDownloadRef) int32 {
	result, callErr := trySecureDownloadCreateWithTicket(ticket, setup, setupContext, evaluate, evaluateContext, downloadRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secureDownloadFinished func(downloadRef SecureDownloadRef) int32
var _secureDownloadFinishedErr error

func trySecureDownloadFinished(downloadRef SecureDownloadRef) (int32, error) {
	if _secureDownloadFinished == nil {
		return 0, symbolCallError("SecureDownloadFinished", "10.5", _secureDownloadFinishedErr)
	}
	return _secureDownloadFinished(downloadRef), nil
}

// SecureDownloadFinished concludes the secure download process.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecureDownloadFinished
func SecureDownloadFinished(downloadRef SecureDownloadRef) int32 {
	result, callErr := trySecureDownloadFinished(downloadRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secureDownloadGetDownloadSize func(downloadRef SecureDownloadRef, downloadSize *int64) int32
var _secureDownloadGetDownloadSizeErr error

func trySecureDownloadGetDownloadSize(downloadRef SecureDownloadRef, downloadSize *int64) (int32, error) {
	if _secureDownloadGetDownloadSize == nil {
		return 0, symbolCallError("SecureDownloadGetDownloadSize", "10.5", _secureDownloadGetDownloadSizeErr)
	}
	return _secureDownloadGetDownloadSize(downloadRef, downloadSize), nil
}

// SecureDownloadGetDownloadSize returns the size of the expected download.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecureDownloadGetDownloadSize
func SecureDownloadGetDownloadSize(downloadRef SecureDownloadRef, downloadSize *int64) int32 {
	result, callErr := trySecureDownloadGetDownloadSize(downloadRef, downloadSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secureDownloadRelease func(downloadRef SecureDownloadRef) int32
var _secureDownloadReleaseErr error

func trySecureDownloadRelease(downloadRef SecureDownloadRef) (int32, error) {
	if _secureDownloadRelease == nil {
		return 0, symbolCallError("SecureDownloadRelease", "10.5", _secureDownloadReleaseErr)
	}
	return _secureDownloadRelease(downloadRef), nil
}

// SecureDownloadRelease releases the memory associated with a secure download object.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecureDownloadRelease
func SecureDownloadRelease(downloadRef SecureDownloadRef) int32 {
	result, callErr := trySecureDownloadRelease(downloadRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secureDownloadUpdateWithData func(downloadRef SecureDownloadRef, data corefoundation.CFDataRef) int32
var _secureDownloadUpdateWithDataErr error

func trySecureDownloadUpdateWithData(downloadRef SecureDownloadRef, data corefoundation.CFDataRef) (int32, error) {
	if _secureDownloadUpdateWithData == nil {
		return 0, symbolCallError("SecureDownloadUpdateWithData", "10.5", _secureDownloadUpdateWithDataErr)
	}
	return _secureDownloadUpdateWithData(downloadRef, data), nil
}

// SecureDownloadUpdateWithData checks data received during download for validity.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecureDownloadUpdateWithData
func SecureDownloadUpdateWithData(downloadRef SecureDownloadRef, data corefoundation.CFDataRef) int32 {
	result, callErr := trySecureDownloadUpdateWithData(downloadRef, data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sessionCreate func(flags SessionCreationFlags, attributes SessionAttributeBits) int32
var _sessionCreateErr error

func trySessionCreate(flags SessionCreationFlags, attributes SessionAttributeBits) (int32, error) {
	if _sessionCreate == nil {
		return 0, symbolCallError("SessionCreate", "10.0", _sessionCreateErr)
	}
	return _sessionCreate(flags, attributes), nil
}

// SessionCreate creates a security session.
//
// See: https://developer.apple.com/documentation/Security/SessionCreate(_:_:)
func SessionCreate(flags SessionCreationFlags, attributes SessionAttributeBits) int32 {
	result, callErr := trySessionCreate(flags, attributes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sessionGetInfo func(session SecuritySessionId, sessionId *SecuritySessionId, attributes *SessionAttributeBits) int32
var _sessionGetInfoErr error

func trySessionGetInfo(session SecuritySessionId, sessionId *SecuritySessionId, attributes *SessionAttributeBits) (int32, error) {
	if _sessionGetInfo == nil {
		return 0, symbolCallError("SessionGetInfo", "10.0", _sessionGetInfoErr)
	}
	return _sessionGetInfo(session, sessionId, attributes), nil
}

// SessionGetInfo obtains information about a security session.
//
// See: https://developer.apple.com/documentation/Security/SessionGetInfo(_:_:_:)
func SessionGetInfo(session SecuritySessionId, sessionId *SecuritySessionId, attributes *SessionAttributeBits) int32 {
	result, callErr := trySessionGetInfo(session, sessionId, attributes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cssmAlgToOid func(algId CSSM_ALGORITHMS) unsafe.Pointer
var _cssmAlgToOidErr error

func tryCssmAlgToOid(algId CSSM_ALGORITHMS) (unsafe.Pointer, error) {
	if _cssmAlgToOid == nil {
		return nil, symbolCallError("cssmAlgToOid", "10.0", _cssmAlgToOidErr)
	}
	return _cssmAlgToOid(algId), nil
}

// CssmAlgToOid.
//
// See: https://developer.apple.com/documentation/Security/cssmAlgToOid(_:)
func CssmAlgToOid(algId CSSM_ALGORITHMS) unsafe.Pointer {
	result, callErr := tryCssmAlgToOid(algId)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cssmOidToAlg func(oid unsafe.Pointer, alg unsafe.Pointer) bool
var _cssmOidToAlgErr error

func tryCssmOidToAlg(oid unsafe.Pointer, alg unsafe.Pointer) (bool, error) {
	if _cssmOidToAlg == nil {
		return false, symbolCallError("cssmOidToAlg", "10.0", _cssmOidToAlgErr)
	}
	return _cssmOidToAlg(oid, alg), nil
}

// CssmOidToAlg.
//
// See: https://developer.apple.com/documentation/Security/cssmOidToAlg(_:_:)
func CssmOidToAlg(oid unsafe.Pointer, alg unsafe.Pointer) bool {
	result, callErr := tryCssmOidToAlg(oid, alg)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cssmPerror func(how string, err CSSM_RETURN)
var _cssmPerrorErr error

func tryCssmPerror(how string, err CSSM_RETURN) error {
	if _cssmPerror == nil {
		return symbolCallError("cssmPerror", "10.0", _cssmPerrorErr)
	}
	_cssmPerror(how, err)
	return nil
}

// CssmPerror.
//
// See: https://developer.apple.com/documentation/Security/cssmPerror(_:_:)
func CssmPerror(how string, err CSSM_RETURN) {
	if callErr := tryCssmPerror(how, err); callErr != nil {
		panic(callErr)
	}
}

var _sec_certificate_copy_ref func(certificate Sec_certificate_t) SecCertificateRef
var _sec_certificate_copy_refErr error

func trySec_certificate_copy_ref(certificate Sec_certificate_t) (SecCertificateRef, error) {
	if _sec_certificate_copy_ref == nil {
		return 0, symbolCallError("sec_certificate_copy_ref", "10.14", _sec_certificate_copy_refErr)
	}
	return _sec_certificate_copy_ref(certificate), nil
}

// Sec_certificate_copy_ref.
//
// See: https://developer.apple.com/documentation/Security/sec_certificate_copy_ref(_:)
func Sec_certificate_copy_ref(certificate Sec_certificate_t) SecCertificateRef {
	result, callErr := trySec_certificate_copy_ref(certificate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_certificate_create func(certificate SecCertificateRef) Sec_certificate_t
var _sec_certificate_createErr error

func trySec_certificate_create(certificate SecCertificateRef) (Sec_certificate_t, error) {
	if _sec_certificate_create == nil {
		return *new(Sec_certificate_t), symbolCallError("sec_certificate_create", "10.14", _sec_certificate_createErr)
	}
	return _sec_certificate_create(certificate), nil
}

// Sec_certificate_create.
//
// See: https://developer.apple.com/documentation/Security/sec_certificate_create(_:)
func Sec_certificate_create(certificate SecCertificateRef) Sec_certificate_t {
	result, callErr := trySec_certificate_create(certificate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_identity_access_certificates func(identity Sec_identity_t) bool
var _sec_identity_access_certificatesErr error

func trySec_identity_access_certificates(identity Sec_identity_t) (bool, error) {
	if _sec_identity_access_certificates == nil {
		return false, symbolCallError("sec_identity_access_certificates", "10.15", _sec_identity_access_certificatesErr)
	}
	return _sec_identity_access_certificates(identity), nil
}

// Sec_identity_access_certificates.
//
// See: https://developer.apple.com/documentation/Security/sec_identity_access_certificates(_:_:)
func Sec_identity_access_certificates(identity Sec_identity_t) bool {
	result, callErr := trySec_identity_access_certificates(identity)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_identity_copy_certificates_ref func(identity Sec_identity_t) corefoundation.CFArrayRef
var _sec_identity_copy_certificates_refErr error

func trySec_identity_copy_certificates_ref(identity Sec_identity_t) (corefoundation.CFArrayRef, error) {
	if _sec_identity_copy_certificates_ref == nil {
		return 0, symbolCallError("sec_identity_copy_certificates_ref", "10.14", _sec_identity_copy_certificates_refErr)
	}
	return _sec_identity_copy_certificates_ref(identity), nil
}

// Sec_identity_copy_certificates_ref.
//
// See: https://developer.apple.com/documentation/Security/sec_identity_copy_certificates_ref(_:)
func Sec_identity_copy_certificates_ref(identity Sec_identity_t) corefoundation.CFArrayRef {
	result, callErr := trySec_identity_copy_certificates_ref(identity)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_identity_copy_ref func(identity Sec_identity_t) SecIdentityRef
var _sec_identity_copy_refErr error

func trySec_identity_copy_ref(identity Sec_identity_t) (SecIdentityRef, error) {
	if _sec_identity_copy_ref == nil {
		return 0, symbolCallError("sec_identity_copy_ref", "10.14", _sec_identity_copy_refErr)
	}
	return _sec_identity_copy_ref(identity), nil
}

// Sec_identity_copy_ref.
//
// See: https://developer.apple.com/documentation/Security/sec_identity_copy_ref(_:)
func Sec_identity_copy_ref(identity Sec_identity_t) SecIdentityRef {
	result, callErr := trySec_identity_copy_ref(identity)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_identity_create func(identity SecIdentityRef) Sec_identity_t
var _sec_identity_createErr error

func trySec_identity_create(identity SecIdentityRef) (Sec_identity_t, error) {
	if _sec_identity_create == nil {
		return *new(Sec_identity_t), symbolCallError("sec_identity_create", "10.14", _sec_identity_createErr)
	}
	return _sec_identity_create(identity), nil
}

// Sec_identity_create.
//
// See: https://developer.apple.com/documentation/Security/sec_identity_create(_:)
func Sec_identity_create(identity SecIdentityRef) Sec_identity_t {
	result, callErr := trySec_identity_create(identity)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_identity_create_with_certificates func(identity SecIdentityRef, certificates corefoundation.CFArrayRef) Sec_identity_t
var _sec_identity_create_with_certificatesErr error

func trySec_identity_create_with_certificates(identity SecIdentityRef, certificates corefoundation.CFArrayRef) (Sec_identity_t, error) {
	if _sec_identity_create_with_certificates == nil {
		return *new(Sec_identity_t), symbolCallError("sec_identity_create_with_certificates", "10.14", _sec_identity_create_with_certificatesErr)
	}
	return _sec_identity_create_with_certificates(identity, certificates), nil
}

// Sec_identity_create_with_certificates.
//
// See: https://developer.apple.com/documentation/Security/sec_identity_create_with_certificates(_:_:)
func Sec_identity_create_with_certificates(identity SecIdentityRef, certificates corefoundation.CFArrayRef) Sec_identity_t {
	result, callErr := trySec_identity_create_with_certificates(identity, certificates)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_metadata_access_distinguished_names func(metadata Sec_protocol_metadata_t) bool
var _sec_protocol_metadata_access_distinguished_namesErr error

func trySec_protocol_metadata_access_distinguished_names(metadata Sec_protocol_metadata_t) (bool, error) {
	if _sec_protocol_metadata_access_distinguished_names == nil {
		return false, symbolCallError("sec_protocol_metadata_access_distinguished_names", "10.14", _sec_protocol_metadata_access_distinguished_namesErr)
	}
	return _sec_protocol_metadata_access_distinguished_names(metadata), nil
}

// Sec_protocol_metadata_access_distinguished_names.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_access_distinguished_names(_:_:)
func Sec_protocol_metadata_access_distinguished_names(metadata Sec_protocol_metadata_t) bool {
	result, callErr := trySec_protocol_metadata_access_distinguished_names(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_metadata_access_ocsp_response func(metadata Sec_protocol_metadata_t) bool
var _sec_protocol_metadata_access_ocsp_responseErr error

func trySec_protocol_metadata_access_ocsp_response(metadata Sec_protocol_metadata_t) (bool, error) {
	if _sec_protocol_metadata_access_ocsp_response == nil {
		return false, symbolCallError("sec_protocol_metadata_access_ocsp_response", "10.14", _sec_protocol_metadata_access_ocsp_responseErr)
	}
	return _sec_protocol_metadata_access_ocsp_response(metadata), nil
}

// Sec_protocol_metadata_access_ocsp_response.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_access_ocsp_response(_:_:)
func Sec_protocol_metadata_access_ocsp_response(metadata Sec_protocol_metadata_t) bool {
	result, callErr := trySec_protocol_metadata_access_ocsp_response(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_metadata_access_peer_certificate_chain func(metadata Sec_protocol_metadata_t) bool
var _sec_protocol_metadata_access_peer_certificate_chainErr error

func trySec_protocol_metadata_access_peer_certificate_chain(metadata Sec_protocol_metadata_t) (bool, error) {
	if _sec_protocol_metadata_access_peer_certificate_chain == nil {
		return false, symbolCallError("sec_protocol_metadata_access_peer_certificate_chain", "10.14", _sec_protocol_metadata_access_peer_certificate_chainErr)
	}
	return _sec_protocol_metadata_access_peer_certificate_chain(metadata), nil
}

// Sec_protocol_metadata_access_peer_certificate_chain.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_access_peer_certificate_chain(_:_:)
func Sec_protocol_metadata_access_peer_certificate_chain(metadata Sec_protocol_metadata_t) bool {
	result, callErr := trySec_protocol_metadata_access_peer_certificate_chain(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_metadata_access_pre_shared_keys func(metadata Sec_protocol_metadata_t) bool
var _sec_protocol_metadata_access_pre_shared_keysErr error

func trySec_protocol_metadata_access_pre_shared_keys(metadata Sec_protocol_metadata_t) (bool, error) {
	if _sec_protocol_metadata_access_pre_shared_keys == nil {
		return false, symbolCallError("sec_protocol_metadata_access_pre_shared_keys", "10.15", _sec_protocol_metadata_access_pre_shared_keysErr)
	}
	return _sec_protocol_metadata_access_pre_shared_keys(metadata), nil
}

// Sec_protocol_metadata_access_pre_shared_keys.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_access_pre_shared_keys(_:_:)
func Sec_protocol_metadata_access_pre_shared_keys(metadata Sec_protocol_metadata_t) bool {
	result, callErr := trySec_protocol_metadata_access_pre_shared_keys(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_metadata_access_supported_signature_algorithms func(metadata Sec_protocol_metadata_t) bool
var _sec_protocol_metadata_access_supported_signature_algorithmsErr error

func trySec_protocol_metadata_access_supported_signature_algorithms(metadata Sec_protocol_metadata_t) (bool, error) {
	if _sec_protocol_metadata_access_supported_signature_algorithms == nil {
		return false, symbolCallError("sec_protocol_metadata_access_supported_signature_algorithms", "10.14", _sec_protocol_metadata_access_supported_signature_algorithmsErr)
	}
	return _sec_protocol_metadata_access_supported_signature_algorithms(metadata), nil
}

// Sec_protocol_metadata_access_supported_signature_algorithms.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_access_supported_signature_algorithms(_:_:)
func Sec_protocol_metadata_access_supported_signature_algorithms(metadata Sec_protocol_metadata_t) bool {
	result, callErr := trySec_protocol_metadata_access_supported_signature_algorithms(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_metadata_challenge_parameters_are_equal func(metadataA Sec_protocol_metadata_t, metadataB Sec_protocol_metadata_t) bool
var _sec_protocol_metadata_challenge_parameters_are_equalErr error

func trySec_protocol_metadata_challenge_parameters_are_equal(metadataA Sec_protocol_metadata_t, metadataB Sec_protocol_metadata_t) (bool, error) {
	if _sec_protocol_metadata_challenge_parameters_are_equal == nil {
		return false, symbolCallError("sec_protocol_metadata_challenge_parameters_are_equal", "10.14", _sec_protocol_metadata_challenge_parameters_are_equalErr)
	}
	return _sec_protocol_metadata_challenge_parameters_are_equal(metadataA, metadataB), nil
}

// Sec_protocol_metadata_challenge_parameters_are_equal.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_challenge_parameters_are_equal(_:_:)
func Sec_protocol_metadata_challenge_parameters_are_equal(metadataA Sec_protocol_metadata_t, metadataB Sec_protocol_metadata_t) bool {
	result, callErr := trySec_protocol_metadata_challenge_parameters_are_equal(metadataA, metadataB)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_metadata_copy_negotiated_protocol func(metadata Sec_protocol_metadata_t) *byte
var _sec_protocol_metadata_copy_negotiated_protocolErr error

func trySec_protocol_metadata_copy_negotiated_protocol(metadata Sec_protocol_metadata_t) (*byte, error) {
	if _sec_protocol_metadata_copy_negotiated_protocol == nil {
		return nil, symbolCallError("sec_protocol_metadata_copy_negotiated_protocol", "15.5", _sec_protocol_metadata_copy_negotiated_protocolErr)
	}
	return _sec_protocol_metadata_copy_negotiated_protocol(metadata), nil
}

// Sec_protocol_metadata_copy_negotiated_protocol.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_copy_negotiated_protocol(_:)
func Sec_protocol_metadata_copy_negotiated_protocol(metadata Sec_protocol_metadata_t) *byte {
	result, callErr := trySec_protocol_metadata_copy_negotiated_protocol(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_metadata_copy_peer_public_key func(metadata Sec_protocol_metadata_t) uintptr
var _sec_protocol_metadata_copy_peer_public_keyErr error

func trySec_protocol_metadata_copy_peer_public_key(metadata Sec_protocol_metadata_t) (dispatch.Data, error) {
	if _sec_protocol_metadata_copy_peer_public_key == nil {
		return dispatch.DataFromHandle(0), symbolCallError("sec_protocol_metadata_copy_peer_public_key", "10.14", _sec_protocol_metadata_copy_peer_public_keyErr)
	}
	return dispatch.DataFromHandle(_sec_protocol_metadata_copy_peer_public_key(metadata)), nil
}

// Sec_protocol_metadata_copy_peer_public_key.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_copy_peer_public_key(_:)
func Sec_protocol_metadata_copy_peer_public_key(metadata Sec_protocol_metadata_t) dispatch.Data {
	result, callErr := trySec_protocol_metadata_copy_peer_public_key(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_metadata_copy_server_name func(metadata Sec_protocol_metadata_t) *byte
var _sec_protocol_metadata_copy_server_nameErr error

func trySec_protocol_metadata_copy_server_name(metadata Sec_protocol_metadata_t) (*byte, error) {
	if _sec_protocol_metadata_copy_server_name == nil {
		return nil, symbolCallError("sec_protocol_metadata_copy_server_name", "15.5", _sec_protocol_metadata_copy_server_nameErr)
	}
	return _sec_protocol_metadata_copy_server_name(metadata), nil
}

// Sec_protocol_metadata_copy_server_name.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_copy_server_name(_:)
func Sec_protocol_metadata_copy_server_name(metadata Sec_protocol_metadata_t) *byte {
	result, callErr := trySec_protocol_metadata_copy_server_name(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_metadata_create_secret func(metadata Sec_protocol_metadata_t, label_len uintptr, label string, exporter_length uintptr) uintptr
var _sec_protocol_metadata_create_secretErr error

func trySec_protocol_metadata_create_secret(metadata Sec_protocol_metadata_t, label_len uintptr, label string, exporter_length uintptr) (dispatch.Data, error) {
	if _sec_protocol_metadata_create_secret == nil {
		return dispatch.DataFromHandle(0), symbolCallError("sec_protocol_metadata_create_secret", "10.14", _sec_protocol_metadata_create_secretErr)
	}
	return dispatch.DataFromHandle(_sec_protocol_metadata_create_secret(metadata, label_len, label, exporter_length)), nil
}

// Sec_protocol_metadata_create_secret.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_create_secret(_:_:_:_:)
func Sec_protocol_metadata_create_secret(metadata Sec_protocol_metadata_t, label_len uintptr, label string, exporter_length uintptr) dispatch.Data {
	result, callErr := trySec_protocol_metadata_create_secret(metadata, label_len, label, exporter_length)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_metadata_create_secret_with_context func(metadata Sec_protocol_metadata_t, label_len uintptr, label string, context_len uintptr, context *uint8, exporter_length uintptr) uintptr
var _sec_protocol_metadata_create_secret_with_contextErr error

func trySec_protocol_metadata_create_secret_with_context(metadata Sec_protocol_metadata_t, label_len uintptr, label string, context_len uintptr, context *uint8, exporter_length uintptr) (dispatch.Data, error) {
	if _sec_protocol_metadata_create_secret_with_context == nil {
		return dispatch.DataFromHandle(0), symbolCallError("sec_protocol_metadata_create_secret_with_context", "10.14", _sec_protocol_metadata_create_secret_with_contextErr)
	}
	return dispatch.DataFromHandle(_sec_protocol_metadata_create_secret_with_context(metadata, label_len, label, context_len, context, exporter_length)), nil
}

// Sec_protocol_metadata_create_secret_with_context.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_create_secret_with_context(_:_:_:_:_:_:)
func Sec_protocol_metadata_create_secret_with_context(metadata Sec_protocol_metadata_t, label_len uintptr, label string, context_len uintptr, context *uint8, exporter_length uintptr) dispatch.Data {
	result, callErr := trySec_protocol_metadata_create_secret_with_context(metadata, label_len, label, context_len, context, exporter_length)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_metadata_get_early_data_accepted func(metadata Sec_protocol_metadata_t) bool
var _sec_protocol_metadata_get_early_data_acceptedErr error

func trySec_protocol_metadata_get_early_data_accepted(metadata Sec_protocol_metadata_t) (bool, error) {
	if _sec_protocol_metadata_get_early_data_accepted == nil {
		return false, symbolCallError("sec_protocol_metadata_get_early_data_accepted", "10.14", _sec_protocol_metadata_get_early_data_acceptedErr)
	}
	return _sec_protocol_metadata_get_early_data_accepted(metadata), nil
}

// Sec_protocol_metadata_get_early_data_accepted.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_get_early_data_accepted(_:)
func Sec_protocol_metadata_get_early_data_accepted(metadata Sec_protocol_metadata_t) bool {
	result, callErr := trySec_protocol_metadata_get_early_data_accepted(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_metadata_get_negotiated_ciphersuite func(metadata Sec_protocol_metadata_t) SSLCipherSuite
var _sec_protocol_metadata_get_negotiated_ciphersuiteErr error

func trySec_protocol_metadata_get_negotiated_ciphersuite(metadata Sec_protocol_metadata_t) (SSLCipherSuite, error) {
	if _sec_protocol_metadata_get_negotiated_ciphersuite == nil {
		return *new(SSLCipherSuite), symbolCallError("sec_protocol_metadata_get_negotiated_ciphersuite", "10.14", _sec_protocol_metadata_get_negotiated_ciphersuiteErr)
	}
	return _sec_protocol_metadata_get_negotiated_ciphersuite(metadata), nil
}

// Sec_protocol_metadata_get_negotiated_ciphersuite.
//
// Deprecated: Deprecated since macOS 10.15.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_get_negotiated_ciphersuite(_:)
func Sec_protocol_metadata_get_negotiated_ciphersuite(metadata Sec_protocol_metadata_t) SSLCipherSuite {
	result, callErr := trySec_protocol_metadata_get_negotiated_ciphersuite(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_metadata_get_negotiated_protocol func(metadata Sec_protocol_metadata_t) *byte
var _sec_protocol_metadata_get_negotiated_protocolErr error

func trySec_protocol_metadata_get_negotiated_protocol(metadata Sec_protocol_metadata_t) (*byte, error) {
	if _sec_protocol_metadata_get_negotiated_protocol == nil {
		return nil, symbolCallError("sec_protocol_metadata_get_negotiated_protocol", "10.14", _sec_protocol_metadata_get_negotiated_protocolErr)
	}
	return _sec_protocol_metadata_get_negotiated_protocol(metadata), nil
}

// Sec_protocol_metadata_get_negotiated_protocol.
//
// Deprecated: Deprecated since macOS 15.5.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_get_negotiated_protocol(_:)
func Sec_protocol_metadata_get_negotiated_protocol(metadata Sec_protocol_metadata_t) *byte {
	result, callErr := trySec_protocol_metadata_get_negotiated_protocol(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_metadata_get_negotiated_protocol_version func(metadata Sec_protocol_metadata_t) SSLProtocol
var _sec_protocol_metadata_get_negotiated_protocol_versionErr error

func trySec_protocol_metadata_get_negotiated_protocol_version(metadata Sec_protocol_metadata_t) (SSLProtocol, error) {
	if _sec_protocol_metadata_get_negotiated_protocol_version == nil {
		return *new(SSLProtocol), symbolCallError("sec_protocol_metadata_get_negotiated_protocol_version", "10.14", _sec_protocol_metadata_get_negotiated_protocol_versionErr)
	}
	return _sec_protocol_metadata_get_negotiated_protocol_version(metadata), nil
}

// Sec_protocol_metadata_get_negotiated_protocol_version.
//
// Deprecated: Deprecated since macOS 10.15.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_get_negotiated_protocol_version(_:)
func Sec_protocol_metadata_get_negotiated_protocol_version(metadata Sec_protocol_metadata_t) SSLProtocol {
	result, callErr := trySec_protocol_metadata_get_negotiated_protocol_version(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_metadata_get_negotiated_tls_ciphersuite func(metadata Sec_protocol_metadata_t) Tls_ciphersuite_t
var _sec_protocol_metadata_get_negotiated_tls_ciphersuiteErr error

func trySec_protocol_metadata_get_negotiated_tls_ciphersuite(metadata Sec_protocol_metadata_t) (Tls_ciphersuite_t, error) {
	if _sec_protocol_metadata_get_negotiated_tls_ciphersuite == nil {
		return *new(Tls_ciphersuite_t), symbolCallError("sec_protocol_metadata_get_negotiated_tls_ciphersuite", "10.15", _sec_protocol_metadata_get_negotiated_tls_ciphersuiteErr)
	}
	return _sec_protocol_metadata_get_negotiated_tls_ciphersuite(metadata), nil
}

// Sec_protocol_metadata_get_negotiated_tls_ciphersuite.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_get_negotiated_tls_ciphersuite(_:)
func Sec_protocol_metadata_get_negotiated_tls_ciphersuite(metadata Sec_protocol_metadata_t) Tls_ciphersuite_t {
	result, callErr := trySec_protocol_metadata_get_negotiated_tls_ciphersuite(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_metadata_get_negotiated_tls_protocol_version func(metadata Sec_protocol_metadata_t) Tls_protocol_version_t
var _sec_protocol_metadata_get_negotiated_tls_protocol_versionErr error

func trySec_protocol_metadata_get_negotiated_tls_protocol_version(metadata Sec_protocol_metadata_t) (Tls_protocol_version_t, error) {
	if _sec_protocol_metadata_get_negotiated_tls_protocol_version == nil {
		return *new(Tls_protocol_version_t), symbolCallError("sec_protocol_metadata_get_negotiated_tls_protocol_version", "10.15", _sec_protocol_metadata_get_negotiated_tls_protocol_versionErr)
	}
	return _sec_protocol_metadata_get_negotiated_tls_protocol_version(metadata), nil
}

// Sec_protocol_metadata_get_negotiated_tls_protocol_version.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_get_negotiated_tls_protocol_version(_:)
func Sec_protocol_metadata_get_negotiated_tls_protocol_version(metadata Sec_protocol_metadata_t) Tls_protocol_version_t {
	result, callErr := trySec_protocol_metadata_get_negotiated_tls_protocol_version(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_metadata_get_server_name func(metadata Sec_protocol_metadata_t) *byte
var _sec_protocol_metadata_get_server_nameErr error

func trySec_protocol_metadata_get_server_name(metadata Sec_protocol_metadata_t) (*byte, error) {
	if _sec_protocol_metadata_get_server_name == nil {
		return nil, symbolCallError("sec_protocol_metadata_get_server_name", "10.14", _sec_protocol_metadata_get_server_nameErr)
	}
	return _sec_protocol_metadata_get_server_name(metadata), nil
}

// Sec_protocol_metadata_get_server_name.
//
// Deprecated: Deprecated since macOS 15.5.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_get_server_name(_:)
func Sec_protocol_metadata_get_server_name(metadata Sec_protocol_metadata_t) *byte {
	result, callErr := trySec_protocol_metadata_get_server_name(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_metadata_peers_are_equal func(metadataA Sec_protocol_metadata_t, metadataB Sec_protocol_metadata_t) bool
var _sec_protocol_metadata_peers_are_equalErr error

func trySec_protocol_metadata_peers_are_equal(metadataA Sec_protocol_metadata_t, metadataB Sec_protocol_metadata_t) (bool, error) {
	if _sec_protocol_metadata_peers_are_equal == nil {
		return false, symbolCallError("sec_protocol_metadata_peers_are_equal", "10.14", _sec_protocol_metadata_peers_are_equalErr)
	}
	return _sec_protocol_metadata_peers_are_equal(metadataA, metadataB), nil
}

// Sec_protocol_metadata_peers_are_equal.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_peers_are_equal(_:_:)
func Sec_protocol_metadata_peers_are_equal(metadataA Sec_protocol_metadata_t, metadataB Sec_protocol_metadata_t) bool {
	result, callErr := trySec_protocol_metadata_peers_are_equal(metadataA, metadataB)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_options_add_pre_shared_key func(options Sec_protocol_options_t, psk uintptr, psk_identity uintptr)
var _sec_protocol_options_add_pre_shared_keyErr error

func trySec_protocol_options_add_pre_shared_key(options Sec_protocol_options_t, psk dispatch.Data, psk_identity dispatch.Data) error {
	if _sec_protocol_options_add_pre_shared_key == nil {
		return symbolCallError("sec_protocol_options_add_pre_shared_key", "10.14", _sec_protocol_options_add_pre_shared_keyErr)
	}
	_sec_protocol_options_add_pre_shared_key(options, uintptr(psk.Handle()), uintptr(psk_identity.Handle()))
	return nil
}

// Sec_protocol_options_add_pre_shared_key.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_add_pre_shared_key(_:_:_:)
func Sec_protocol_options_add_pre_shared_key(options Sec_protocol_options_t, psk dispatch.Data, psk_identity dispatch.Data) {
	if callErr := trySec_protocol_options_add_pre_shared_key(options, psk, psk_identity); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_add_tls_application_protocol func(options Sec_protocol_options_t, application_protocol string)
var _sec_protocol_options_add_tls_application_protocolErr error

func trySec_protocol_options_add_tls_application_protocol(options Sec_protocol_options_t, application_protocol string) error {
	if _sec_protocol_options_add_tls_application_protocol == nil {
		return symbolCallError("sec_protocol_options_add_tls_application_protocol", "10.14", _sec_protocol_options_add_tls_application_protocolErr)
	}
	_sec_protocol_options_add_tls_application_protocol(options, application_protocol)
	return nil
}

// Sec_protocol_options_add_tls_application_protocol.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_add_tls_application_protocol(_:_:)
func Sec_protocol_options_add_tls_application_protocol(options Sec_protocol_options_t, application_protocol string) {
	if callErr := trySec_protocol_options_add_tls_application_protocol(options, application_protocol); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_append_tls_ciphersuite func(options Sec_protocol_options_t, ciphersuite Tls_ciphersuite_t)
var _sec_protocol_options_append_tls_ciphersuiteErr error

func trySec_protocol_options_append_tls_ciphersuite(options Sec_protocol_options_t, ciphersuite Tls_ciphersuite_t) error {
	if _sec_protocol_options_append_tls_ciphersuite == nil {
		return symbolCallError("sec_protocol_options_append_tls_ciphersuite", "10.15", _sec_protocol_options_append_tls_ciphersuiteErr)
	}
	_sec_protocol_options_append_tls_ciphersuite(options, ciphersuite)
	return nil
}

// Sec_protocol_options_append_tls_ciphersuite.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_append_tls_ciphersuite(_:_:)
func Sec_protocol_options_append_tls_ciphersuite(options Sec_protocol_options_t, ciphersuite Tls_ciphersuite_t) {
	if callErr := trySec_protocol_options_append_tls_ciphersuite(options, ciphersuite); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_append_tls_ciphersuite_group func(options Sec_protocol_options_t, group Tls_ciphersuite_group_t)
var _sec_protocol_options_append_tls_ciphersuite_groupErr error

func trySec_protocol_options_append_tls_ciphersuite_group(options Sec_protocol_options_t, group Tls_ciphersuite_group_t) error {
	if _sec_protocol_options_append_tls_ciphersuite_group == nil {
		return symbolCallError("sec_protocol_options_append_tls_ciphersuite_group", "10.15", _sec_protocol_options_append_tls_ciphersuite_groupErr)
	}
	_sec_protocol_options_append_tls_ciphersuite_group(options, group)
	return nil
}

// Sec_protocol_options_append_tls_ciphersuite_group.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_append_tls_ciphersuite_group(_:_:)
func Sec_protocol_options_append_tls_ciphersuite_group(options Sec_protocol_options_t, group Tls_ciphersuite_group_t) {
	if callErr := trySec_protocol_options_append_tls_ciphersuite_group(options, group); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_are_equal func(optionsA Sec_protocol_options_t, optionsB Sec_protocol_options_t) bool
var _sec_protocol_options_are_equalErr error

func trySec_protocol_options_are_equal(optionsA Sec_protocol_options_t, optionsB Sec_protocol_options_t) (bool, error) {
	if _sec_protocol_options_are_equal == nil {
		return false, symbolCallError("sec_protocol_options_are_equal", "10.15", _sec_protocol_options_are_equalErr)
	}
	return _sec_protocol_options_are_equal(optionsA, optionsB), nil
}

// Sec_protocol_options_are_equal.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_are_equal(_:_:)
func Sec_protocol_options_are_equal(optionsA Sec_protocol_options_t, optionsB Sec_protocol_options_t) bool {
	result, callErr := trySec_protocol_options_are_equal(optionsA, optionsB)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_options_get_default_max_dtls_protocol_version func() Tls_protocol_version_t
var _sec_protocol_options_get_default_max_dtls_protocol_versionErr error

func trySec_protocol_options_get_default_max_dtls_protocol_version() (Tls_protocol_version_t, error) {
	if _sec_protocol_options_get_default_max_dtls_protocol_version == nil {
		return *new(Tls_protocol_version_t), symbolCallError("sec_protocol_options_get_default_max_dtls_protocol_version", "10.15", _sec_protocol_options_get_default_max_dtls_protocol_versionErr)
	}
	return _sec_protocol_options_get_default_max_dtls_protocol_version(), nil
}

// Sec_protocol_options_get_default_max_dtls_protocol_version.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_get_default_max_dtls_protocol_version()
func Sec_protocol_options_get_default_max_dtls_protocol_version() Tls_protocol_version_t {
	result, callErr := trySec_protocol_options_get_default_max_dtls_protocol_version()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_options_get_default_max_tls_protocol_version func() Tls_protocol_version_t
var _sec_protocol_options_get_default_max_tls_protocol_versionErr error

func trySec_protocol_options_get_default_max_tls_protocol_version() (Tls_protocol_version_t, error) {
	if _sec_protocol_options_get_default_max_tls_protocol_version == nil {
		return *new(Tls_protocol_version_t), symbolCallError("sec_protocol_options_get_default_max_tls_protocol_version", "10.15", _sec_protocol_options_get_default_max_tls_protocol_versionErr)
	}
	return _sec_protocol_options_get_default_max_tls_protocol_version(), nil
}

// Sec_protocol_options_get_default_max_tls_protocol_version.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_get_default_max_tls_protocol_version()
func Sec_protocol_options_get_default_max_tls_protocol_version() Tls_protocol_version_t {
	result, callErr := trySec_protocol_options_get_default_max_tls_protocol_version()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_options_get_default_min_dtls_protocol_version func() Tls_protocol_version_t
var _sec_protocol_options_get_default_min_dtls_protocol_versionErr error

func trySec_protocol_options_get_default_min_dtls_protocol_version() (Tls_protocol_version_t, error) {
	if _sec_protocol_options_get_default_min_dtls_protocol_version == nil {
		return *new(Tls_protocol_version_t), symbolCallError("sec_protocol_options_get_default_min_dtls_protocol_version", "10.15", _sec_protocol_options_get_default_min_dtls_protocol_versionErr)
	}
	return _sec_protocol_options_get_default_min_dtls_protocol_version(), nil
}

// Sec_protocol_options_get_default_min_dtls_protocol_version.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_get_default_min_dtls_protocol_version()
func Sec_protocol_options_get_default_min_dtls_protocol_version() Tls_protocol_version_t {
	result, callErr := trySec_protocol_options_get_default_min_dtls_protocol_version()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_options_get_default_min_tls_protocol_version func() Tls_protocol_version_t
var _sec_protocol_options_get_default_min_tls_protocol_versionErr error

func trySec_protocol_options_get_default_min_tls_protocol_version() (Tls_protocol_version_t, error) {
	if _sec_protocol_options_get_default_min_tls_protocol_version == nil {
		return *new(Tls_protocol_version_t), symbolCallError("sec_protocol_options_get_default_min_tls_protocol_version", "10.15", _sec_protocol_options_get_default_min_tls_protocol_versionErr)
	}
	return _sec_protocol_options_get_default_min_tls_protocol_version(), nil
}

// Sec_protocol_options_get_default_min_tls_protocol_version.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_get_default_min_tls_protocol_version()
func Sec_protocol_options_get_default_min_tls_protocol_version() Tls_protocol_version_t {
	result, callErr := trySec_protocol_options_get_default_min_tls_protocol_version()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_options_get_enable_encrypted_client_hello func(options Sec_protocol_options_t) bool
var _sec_protocol_options_get_enable_encrypted_client_helloErr error

func trySec_protocol_options_get_enable_encrypted_client_hello(options Sec_protocol_options_t) (bool, error) {
	if _sec_protocol_options_get_enable_encrypted_client_hello == nil {
		return false, symbolCallError("sec_protocol_options_get_enable_encrypted_client_hello", "", _sec_protocol_options_get_enable_encrypted_client_helloErr)
	}
	return _sec_protocol_options_get_enable_encrypted_client_hello(options), nil
}

// Sec_protocol_options_get_enable_encrypted_client_hello.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_get_enable_encrypted_client_hello
func Sec_protocol_options_get_enable_encrypted_client_hello(options Sec_protocol_options_t) bool {
	result, callErr := trySec_protocol_options_get_enable_encrypted_client_hello(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_options_get_quic_use_legacy_codepoint func(options Sec_protocol_options_t) bool
var _sec_protocol_options_get_quic_use_legacy_codepointErr error

func trySec_protocol_options_get_quic_use_legacy_codepoint(options Sec_protocol_options_t) (bool, error) {
	if _sec_protocol_options_get_quic_use_legacy_codepoint == nil {
		return false, symbolCallError("sec_protocol_options_get_quic_use_legacy_codepoint", "", _sec_protocol_options_get_quic_use_legacy_codepointErr)
	}
	return _sec_protocol_options_get_quic_use_legacy_codepoint(options), nil
}

// Sec_protocol_options_get_quic_use_legacy_codepoint.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_get_quic_use_legacy_codepoint
func Sec_protocol_options_get_quic_use_legacy_codepoint(options Sec_protocol_options_t) bool {
	result, callErr := trySec_protocol_options_get_quic_use_legacy_codepoint(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_options_set_challenge_block func(options Sec_protocol_options_t, challenge_block unsafe.Pointer, challenge_queue uintptr)
var _sec_protocol_options_set_challenge_blockErr error

func trySec_protocol_options_set_challenge_block(options Sec_protocol_options_t, challenge_block Sec_protocol_challenge_t, challenge_queue dispatch.Queue) error {
	if _sec_protocol_options_set_challenge_block == nil {
		return symbolCallError("sec_protocol_options_set_challenge_block", "10.14", _sec_protocol_options_set_challenge_blockErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objectivec.Object, blockArg1 func(*objectivec.Object)) {
		challenge_block(blockArg0, blockArg1)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_sec_protocol_options_set_challenge_block(options, _block0, uintptr(challenge_queue.Handle()))
	return nil
}

// Sec_protocol_options_set_challenge_block.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_challenge_block(_:_:_:)
func Sec_protocol_options_set_challenge_block(options Sec_protocol_options_t, challenge_block Sec_protocol_challenge_t, challenge_queue dispatch.Queue) {
	if callErr := trySec_protocol_options_set_challenge_block(options, challenge_block, challenge_queue); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_enable_encrypted_client_hello func(options Sec_protocol_options_t, enable_encrypted_client_hello bool)
var _sec_protocol_options_set_enable_encrypted_client_helloErr error

func trySec_protocol_options_set_enable_encrypted_client_hello(options Sec_protocol_options_t, enable_encrypted_client_hello bool) error {
	if _sec_protocol_options_set_enable_encrypted_client_hello == nil {
		return symbolCallError("sec_protocol_options_set_enable_encrypted_client_hello", "", _sec_protocol_options_set_enable_encrypted_client_helloErr)
	}
	_sec_protocol_options_set_enable_encrypted_client_hello(options, enable_encrypted_client_hello)
	return nil
}

// Sec_protocol_options_set_enable_encrypted_client_hello.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_enable_encrypted_client_hello
func Sec_protocol_options_set_enable_encrypted_client_hello(options Sec_protocol_options_t, enable_encrypted_client_hello bool) {
	if callErr := trySec_protocol_options_set_enable_encrypted_client_hello(options, enable_encrypted_client_hello); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_key_update_block func(options Sec_protocol_options_t, key_update_block unsafe.Pointer, key_update_queue uintptr)
var _sec_protocol_options_set_key_update_blockErr error

func trySec_protocol_options_set_key_update_block(options Sec_protocol_options_t, key_update_block Sec_protocol_key_update_t, key_update_queue dispatch.Queue) error {
	if _sec_protocol_options_set_key_update_block == nil {
		return symbolCallError("sec_protocol_options_set_key_update_block", "10.14", _sec_protocol_options_set_key_update_blockErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objectivec.Object, blockArg1 func()) {
		key_update_block(blockArg0, blockArg1)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_sec_protocol_options_set_key_update_block(options, _block0, uintptr(key_update_queue.Handle()))
	return nil
}

// Sec_protocol_options_set_key_update_block.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_key_update_block(_:_:_:)
func Sec_protocol_options_set_key_update_block(options Sec_protocol_options_t, key_update_block Sec_protocol_key_update_t, key_update_queue dispatch.Queue) {
	if callErr := trySec_protocol_options_set_key_update_block(options, key_update_block, key_update_queue); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_local_identity func(options Sec_protocol_options_t, identity Sec_identity_t)
var _sec_protocol_options_set_local_identityErr error

func trySec_protocol_options_set_local_identity(options Sec_protocol_options_t, identity Sec_identity_t) error {
	if _sec_protocol_options_set_local_identity == nil {
		return symbolCallError("sec_protocol_options_set_local_identity", "10.14", _sec_protocol_options_set_local_identityErr)
	}
	_sec_protocol_options_set_local_identity(options, identity)
	return nil
}

// Sec_protocol_options_set_local_identity.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_local_identity(_:_:)
func Sec_protocol_options_set_local_identity(options Sec_protocol_options_t, identity Sec_identity_t) {
	if callErr := trySec_protocol_options_set_local_identity(options, identity); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_max_tls_protocol_version func(options Sec_protocol_options_t, version Tls_protocol_version_t)
var _sec_protocol_options_set_max_tls_protocol_versionErr error

func trySec_protocol_options_set_max_tls_protocol_version(options Sec_protocol_options_t, version Tls_protocol_version_t) error {
	if _sec_protocol_options_set_max_tls_protocol_version == nil {
		return symbolCallError("sec_protocol_options_set_max_tls_protocol_version", "10.15", _sec_protocol_options_set_max_tls_protocol_versionErr)
	}
	_sec_protocol_options_set_max_tls_protocol_version(options, version)
	return nil
}

// Sec_protocol_options_set_max_tls_protocol_version.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_max_tls_protocol_version(_:_:)
func Sec_protocol_options_set_max_tls_protocol_version(options Sec_protocol_options_t, version Tls_protocol_version_t) {
	if callErr := trySec_protocol_options_set_max_tls_protocol_version(options, version); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_min_tls_protocol_version func(options Sec_protocol_options_t, version Tls_protocol_version_t)
var _sec_protocol_options_set_min_tls_protocol_versionErr error

func trySec_protocol_options_set_min_tls_protocol_version(options Sec_protocol_options_t, version Tls_protocol_version_t) error {
	if _sec_protocol_options_set_min_tls_protocol_version == nil {
		return symbolCallError("sec_protocol_options_set_min_tls_protocol_version", "10.15", _sec_protocol_options_set_min_tls_protocol_versionErr)
	}
	_sec_protocol_options_set_min_tls_protocol_version(options, version)
	return nil
}

// Sec_protocol_options_set_min_tls_protocol_version.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_min_tls_protocol_version(_:_:)
func Sec_protocol_options_set_min_tls_protocol_version(options Sec_protocol_options_t, version Tls_protocol_version_t) {
	if callErr := trySec_protocol_options_set_min_tls_protocol_version(options, version); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_peer_authentication_optional func(options Sec_protocol_options_t, peer_authentication_optional bool)
var _sec_protocol_options_set_peer_authentication_optionalErr error

func trySec_protocol_options_set_peer_authentication_optional(options Sec_protocol_options_t, peer_authentication_optional bool) error {
	if _sec_protocol_options_set_peer_authentication_optional == nil {
		return symbolCallError("sec_protocol_options_set_peer_authentication_optional", "", _sec_protocol_options_set_peer_authentication_optionalErr)
	}
	_sec_protocol_options_set_peer_authentication_optional(options, peer_authentication_optional)
	return nil
}

// Sec_protocol_options_set_peer_authentication_optional.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_peer_authentication_optional
func Sec_protocol_options_set_peer_authentication_optional(options Sec_protocol_options_t, peer_authentication_optional bool) {
	if callErr := trySec_protocol_options_set_peer_authentication_optional(options, peer_authentication_optional); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_peer_authentication_required func(options Sec_protocol_options_t, peer_authentication_required bool)
var _sec_protocol_options_set_peer_authentication_requiredErr error

func trySec_protocol_options_set_peer_authentication_required(options Sec_protocol_options_t, peer_authentication_required bool) error {
	if _sec_protocol_options_set_peer_authentication_required == nil {
		return symbolCallError("sec_protocol_options_set_peer_authentication_required", "10.14", _sec_protocol_options_set_peer_authentication_requiredErr)
	}
	_sec_protocol_options_set_peer_authentication_required(options, peer_authentication_required)
	return nil
}

// Sec_protocol_options_set_peer_authentication_required.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_peer_authentication_required(_:_:)
func Sec_protocol_options_set_peer_authentication_required(options Sec_protocol_options_t, peer_authentication_required bool) {
	if callErr := trySec_protocol_options_set_peer_authentication_required(options, peer_authentication_required); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_pre_shared_key_selection_block func(options Sec_protocol_options_t, psk_selection_block unsafe.Pointer, psk_selection_queue uintptr)
var _sec_protocol_options_set_pre_shared_key_selection_blockErr error

func trySec_protocol_options_set_pre_shared_key_selection_block(options Sec_protocol_options_t, psk_selection_block Sec_protocol_pre_shared_key_selection_t, psk_selection_queue dispatch.Queue) error {
	if _sec_protocol_options_set_pre_shared_key_selection_block == nil {
		return symbolCallError("sec_protocol_options_set_pre_shared_key_selection_block", "10.15", _sec_protocol_options_set_pre_shared_key_selection_blockErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objectivec.Object, blockArg1 objectivec.Object, blockArg2 func(*objectivec.Object)) {
		psk_selection_block(blockArg0, blockArg1, blockArg2)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_sec_protocol_options_set_pre_shared_key_selection_block(options, _block0, uintptr(psk_selection_queue.Handle()))
	return nil
}

// Sec_protocol_options_set_pre_shared_key_selection_block.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_pre_shared_key_selection_block(_:_:_:)
func Sec_protocol_options_set_pre_shared_key_selection_block(options Sec_protocol_options_t, psk_selection_block Sec_protocol_pre_shared_key_selection_t, psk_selection_queue dispatch.Queue) {
	if callErr := trySec_protocol_options_set_pre_shared_key_selection_block(options, psk_selection_block, psk_selection_queue); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_quic_use_legacy_codepoint func(options Sec_protocol_options_t, quic_use_legacy_codepoint bool)
var _sec_protocol_options_set_quic_use_legacy_codepointErr error

func trySec_protocol_options_set_quic_use_legacy_codepoint(options Sec_protocol_options_t, quic_use_legacy_codepoint bool) error {
	if _sec_protocol_options_set_quic_use_legacy_codepoint == nil {
		return symbolCallError("sec_protocol_options_set_quic_use_legacy_codepoint", "", _sec_protocol_options_set_quic_use_legacy_codepointErr)
	}
	_sec_protocol_options_set_quic_use_legacy_codepoint(options, quic_use_legacy_codepoint)
	return nil
}

// Sec_protocol_options_set_quic_use_legacy_codepoint.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_quic_use_legacy_codepoint
func Sec_protocol_options_set_quic_use_legacy_codepoint(options Sec_protocol_options_t, quic_use_legacy_codepoint bool) {
	if callErr := trySec_protocol_options_set_quic_use_legacy_codepoint(options, quic_use_legacy_codepoint); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_tls_false_start_enabled func(options Sec_protocol_options_t, false_start_enabled bool)
var _sec_protocol_options_set_tls_false_start_enabledErr error

func trySec_protocol_options_set_tls_false_start_enabled(options Sec_protocol_options_t, false_start_enabled bool) error {
	if _sec_protocol_options_set_tls_false_start_enabled == nil {
		return symbolCallError("sec_protocol_options_set_tls_false_start_enabled", "10.14", _sec_protocol_options_set_tls_false_start_enabledErr)
	}
	_sec_protocol_options_set_tls_false_start_enabled(options, false_start_enabled)
	return nil
}

// Sec_protocol_options_set_tls_false_start_enabled.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_false_start_enabled(_:_:)
func Sec_protocol_options_set_tls_false_start_enabled(options Sec_protocol_options_t, false_start_enabled bool) {
	if callErr := trySec_protocol_options_set_tls_false_start_enabled(options, false_start_enabled); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_tls_is_fallback_attempt func(options Sec_protocol_options_t, is_fallback_attempt bool)
var _sec_protocol_options_set_tls_is_fallback_attemptErr error

func trySec_protocol_options_set_tls_is_fallback_attempt(options Sec_protocol_options_t, is_fallback_attempt bool) error {
	if _sec_protocol_options_set_tls_is_fallback_attempt == nil {
		return symbolCallError("sec_protocol_options_set_tls_is_fallback_attempt", "10.14", _sec_protocol_options_set_tls_is_fallback_attemptErr)
	}
	_sec_protocol_options_set_tls_is_fallback_attempt(options, is_fallback_attempt)
	return nil
}

// Sec_protocol_options_set_tls_is_fallback_attempt.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_is_fallback_attempt(_:_:)
func Sec_protocol_options_set_tls_is_fallback_attempt(options Sec_protocol_options_t, is_fallback_attempt bool) {
	if callErr := trySec_protocol_options_set_tls_is_fallback_attempt(options, is_fallback_attempt); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_tls_max_version func(options Sec_protocol_options_t, version SSLProtocol)
var _sec_protocol_options_set_tls_max_versionErr error

func trySec_protocol_options_set_tls_max_version(options Sec_protocol_options_t, version SSLProtocol) error {
	if _sec_protocol_options_set_tls_max_version == nil {
		return symbolCallError("sec_protocol_options_set_tls_max_version", "10.14", _sec_protocol_options_set_tls_max_versionErr)
	}
	_sec_protocol_options_set_tls_max_version(options, version)
	return nil
}

// Sec_protocol_options_set_tls_max_version.
//
// Deprecated: Deprecated since macOS 10.15.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_max_version(_:_:)
func Sec_protocol_options_set_tls_max_version(options Sec_protocol_options_t, version SSLProtocol) {
	if callErr := trySec_protocol_options_set_tls_max_version(options, version); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_tls_min_version func(options Sec_protocol_options_t, version SSLProtocol)
var _sec_protocol_options_set_tls_min_versionErr error

func trySec_protocol_options_set_tls_min_version(options Sec_protocol_options_t, version SSLProtocol) error {
	if _sec_protocol_options_set_tls_min_version == nil {
		return symbolCallError("sec_protocol_options_set_tls_min_version", "10.14", _sec_protocol_options_set_tls_min_versionErr)
	}
	_sec_protocol_options_set_tls_min_version(options, version)
	return nil
}

// Sec_protocol_options_set_tls_min_version.
//
// Deprecated: Deprecated since macOS 10.15.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_min_version(_:_:)
func Sec_protocol_options_set_tls_min_version(options Sec_protocol_options_t, version SSLProtocol) {
	if callErr := trySec_protocol_options_set_tls_min_version(options, version); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_tls_ocsp_enabled func(options Sec_protocol_options_t, ocsp_enabled bool)
var _sec_protocol_options_set_tls_ocsp_enabledErr error

func trySec_protocol_options_set_tls_ocsp_enabled(options Sec_protocol_options_t, ocsp_enabled bool) error {
	if _sec_protocol_options_set_tls_ocsp_enabled == nil {
		return symbolCallError("sec_protocol_options_set_tls_ocsp_enabled", "10.14", _sec_protocol_options_set_tls_ocsp_enabledErr)
	}
	_sec_protocol_options_set_tls_ocsp_enabled(options, ocsp_enabled)
	return nil
}

// Sec_protocol_options_set_tls_ocsp_enabled.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_ocsp_enabled(_:_:)
func Sec_protocol_options_set_tls_ocsp_enabled(options Sec_protocol_options_t, ocsp_enabled bool) {
	if callErr := trySec_protocol_options_set_tls_ocsp_enabled(options, ocsp_enabled); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_tls_pre_shared_key_identity_hint func(options Sec_protocol_options_t, psk_identity_hint uintptr)
var _sec_protocol_options_set_tls_pre_shared_key_identity_hintErr error

func trySec_protocol_options_set_tls_pre_shared_key_identity_hint(options Sec_protocol_options_t, psk_identity_hint dispatch.Data) error {
	if _sec_protocol_options_set_tls_pre_shared_key_identity_hint == nil {
		return symbolCallError("sec_protocol_options_set_tls_pre_shared_key_identity_hint", "10.15", _sec_protocol_options_set_tls_pre_shared_key_identity_hintErr)
	}
	_sec_protocol_options_set_tls_pre_shared_key_identity_hint(options, uintptr(psk_identity_hint.Handle()))
	return nil
}

// Sec_protocol_options_set_tls_pre_shared_key_identity_hint.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_pre_shared_key_identity_hint(_:_:)
func Sec_protocol_options_set_tls_pre_shared_key_identity_hint(options Sec_protocol_options_t, psk_identity_hint dispatch.Data) {
	if callErr := trySec_protocol_options_set_tls_pre_shared_key_identity_hint(options, psk_identity_hint); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_tls_renegotiation_enabled func(options Sec_protocol_options_t, renegotiation_enabled bool)
var _sec_protocol_options_set_tls_renegotiation_enabledErr error

func trySec_protocol_options_set_tls_renegotiation_enabled(options Sec_protocol_options_t, renegotiation_enabled bool) error {
	if _sec_protocol_options_set_tls_renegotiation_enabled == nil {
		return symbolCallError("sec_protocol_options_set_tls_renegotiation_enabled", "10.14", _sec_protocol_options_set_tls_renegotiation_enabledErr)
	}
	_sec_protocol_options_set_tls_renegotiation_enabled(options, renegotiation_enabled)
	return nil
}

// Sec_protocol_options_set_tls_renegotiation_enabled.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_renegotiation_enabled(_:_:)
func Sec_protocol_options_set_tls_renegotiation_enabled(options Sec_protocol_options_t, renegotiation_enabled bool) {
	if callErr := trySec_protocol_options_set_tls_renegotiation_enabled(options, renegotiation_enabled); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_tls_resumption_enabled func(options Sec_protocol_options_t, resumption_enabled bool)
var _sec_protocol_options_set_tls_resumption_enabledErr error

func trySec_protocol_options_set_tls_resumption_enabled(options Sec_protocol_options_t, resumption_enabled bool) error {
	if _sec_protocol_options_set_tls_resumption_enabled == nil {
		return symbolCallError("sec_protocol_options_set_tls_resumption_enabled", "10.14", _sec_protocol_options_set_tls_resumption_enabledErr)
	}
	_sec_protocol_options_set_tls_resumption_enabled(options, resumption_enabled)
	return nil
}

// Sec_protocol_options_set_tls_resumption_enabled.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_resumption_enabled(_:_:)
func Sec_protocol_options_set_tls_resumption_enabled(options Sec_protocol_options_t, resumption_enabled bool) {
	if callErr := trySec_protocol_options_set_tls_resumption_enabled(options, resumption_enabled); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_tls_sct_enabled func(options Sec_protocol_options_t, sct_enabled bool)
var _sec_protocol_options_set_tls_sct_enabledErr error

func trySec_protocol_options_set_tls_sct_enabled(options Sec_protocol_options_t, sct_enabled bool) error {
	if _sec_protocol_options_set_tls_sct_enabled == nil {
		return symbolCallError("sec_protocol_options_set_tls_sct_enabled", "10.14", _sec_protocol_options_set_tls_sct_enabledErr)
	}
	_sec_protocol_options_set_tls_sct_enabled(options, sct_enabled)
	return nil
}

// Sec_protocol_options_set_tls_sct_enabled.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_sct_enabled(_:_:)
func Sec_protocol_options_set_tls_sct_enabled(options Sec_protocol_options_t, sct_enabled bool) {
	if callErr := trySec_protocol_options_set_tls_sct_enabled(options, sct_enabled); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_tls_server_name func(options Sec_protocol_options_t, server_name string)
var _sec_protocol_options_set_tls_server_nameErr error

func trySec_protocol_options_set_tls_server_name(options Sec_protocol_options_t, server_name string) error {
	if _sec_protocol_options_set_tls_server_name == nil {
		return symbolCallError("sec_protocol_options_set_tls_server_name", "10.14", _sec_protocol_options_set_tls_server_nameErr)
	}
	_sec_protocol_options_set_tls_server_name(options, server_name)
	return nil
}

// Sec_protocol_options_set_tls_server_name.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_server_name(_:_:)
func Sec_protocol_options_set_tls_server_name(options Sec_protocol_options_t, server_name string) {
	if callErr := trySec_protocol_options_set_tls_server_name(options, server_name); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_tls_tickets_enabled func(options Sec_protocol_options_t, tickets_enabled bool)
var _sec_protocol_options_set_tls_tickets_enabledErr error

func trySec_protocol_options_set_tls_tickets_enabled(options Sec_protocol_options_t, tickets_enabled bool) error {
	if _sec_protocol_options_set_tls_tickets_enabled == nil {
		return symbolCallError("sec_protocol_options_set_tls_tickets_enabled", "10.14", _sec_protocol_options_set_tls_tickets_enabledErr)
	}
	_sec_protocol_options_set_tls_tickets_enabled(options, tickets_enabled)
	return nil
}

// Sec_protocol_options_set_tls_tickets_enabled.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_tickets_enabled(_:_:)
func Sec_protocol_options_set_tls_tickets_enabled(options Sec_protocol_options_t, tickets_enabled bool) {
	if callErr := trySec_protocol_options_set_tls_tickets_enabled(options, tickets_enabled); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_verify_block func(options Sec_protocol_options_t, verify_block unsafe.Pointer, verify_block_queue uintptr)
var _sec_protocol_options_set_verify_blockErr error

func trySec_protocol_options_set_verify_block(options Sec_protocol_options_t, verify_block Sec_protocol_verify_t, verify_block_queue dispatch.Queue) error {
	if _sec_protocol_options_set_verify_block == nil {
		return symbolCallError("sec_protocol_options_set_verify_block", "10.14", _sec_protocol_options_set_verify_blockErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 objectivec.Object, blockArg1 objectivec.Object, blockArg2 func(bool)) {
		verify_block(blockArg0, blockArg1, blockArg2)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_sec_protocol_options_set_verify_block(options, _block0, uintptr(verify_block_queue.Handle()))
	return nil
}

// Sec_protocol_options_set_verify_block.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_verify_block(_:_:_:)
func Sec_protocol_options_set_verify_block(options Sec_protocol_options_t, verify_block Sec_protocol_verify_t, verify_block_queue dispatch.Queue) {
	if callErr := trySec_protocol_options_set_verify_block(options, verify_block, verify_block_queue); callErr != nil {
		panic(callErr)
	}
}

var _sec_release func(obj unsafe.Pointer)
var _sec_releaseErr error

func trySec_release(obj unsafe.Pointer) error {
	if _sec_release == nil {
		return symbolCallError("sec_release", "10.0", _sec_releaseErr)
	}
	_sec_release(obj)
	return nil
}

// Sec_release.
//
// See: https://developer.apple.com/documentation/Security/sec_release(_:)
func Sec_release(obj unsafe.Pointer) {
	if callErr := trySec_release(obj); callErr != nil {
		panic(callErr)
	}
}

var _sec_retain func(obj unsafe.Pointer) unsafe.Pointer
var _sec_retainErr error

func trySec_retain(obj unsafe.Pointer) (unsafe.Pointer, error) {
	if _sec_retain == nil {
		return nil, symbolCallError("sec_retain", "10.0", _sec_retainErr)
	}
	return _sec_retain(obj), nil
}

// Sec_retain.
//
// See: https://developer.apple.com/documentation/Security/sec_retain(_:)
func Sec_retain(obj unsafe.Pointer) unsafe.Pointer {
	result, callErr := trySec_retain(obj)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_trust_copy_ref func(trust Sec_trust_t) SecTrustRef
var _sec_trust_copy_refErr error

func trySec_trust_copy_ref(trust Sec_trust_t) (SecTrustRef, error) {
	if _sec_trust_copy_ref == nil {
		return 0, symbolCallError("sec_trust_copy_ref", "10.14", _sec_trust_copy_refErr)
	}
	return _sec_trust_copy_ref(trust), nil
}

// Sec_trust_copy_ref.
//
// See: https://developer.apple.com/documentation/Security/sec_trust_copy_ref(_:)
func Sec_trust_copy_ref(trust Sec_trust_t) SecTrustRef {
	result, callErr := trySec_trust_copy_ref(trust)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_trust_create func(trust SecTrustRef) Sec_trust_t
var _sec_trust_createErr error

func trySec_trust_create(trust SecTrustRef) (Sec_trust_t, error) {
	if _sec_trust_create == nil {
		return *new(Sec_trust_t), symbolCallError("sec_trust_create", "10.14", _sec_trust_createErr)
	}
	return _sec_trust_create(trust), nil
}

// Sec_trust_create.
//
// See: https://developer.apple.com/documentation/Security/sec_trust_create(_:)
func Sec_trust_create(trust SecTrustRef) Sec_trust_t {
	result, callErr := trySec_trust_create(trust)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_authorizationCopyInfo, &_authorizationCopyInfoErr, frameworkHandle, "AuthorizationCopyInfo", "10.0")
	registerFunc(&_authorizationCopyRights, &_authorizationCopyRightsErr, frameworkHandle, "AuthorizationCopyRights", "10.0")
	registerFunc(&_authorizationCopyRightsAsync, &_authorizationCopyRightsAsyncErr, frameworkHandle, "AuthorizationCopyRightsAsync", "10.7")
	registerFunc(&_authorizationCreate, &_authorizationCreateErr, frameworkHandle, "AuthorizationCreate", "10.0")
	registerFunc(&_authorizationCreateFromExternalForm, &_authorizationCreateFromExternalFormErr, frameworkHandle, "AuthorizationCreateFromExternalForm", "10.0")
	registerFunc(&_authorizationFree, &_authorizationFreeErr, frameworkHandle, "AuthorizationFree", "10.0")
	registerFunc(&_authorizationFreeItemSet, &_authorizationFreeItemSetErr, frameworkHandle, "AuthorizationFreeItemSet", "10.0")
	registerFunc(&_authorizationMakeExternalForm, &_authorizationMakeExternalFormErr, frameworkHandle, "AuthorizationMakeExternalForm", "10.0")
	registerFunc(&_authorizationRightGet, &_authorizationRightGetErr, frameworkHandle, "AuthorizationRightGet", "10.0")
	registerFunc(&_authorizationRightRemove, &_authorizationRightRemoveErr, frameworkHandle, "AuthorizationRightRemove", "10.0")
	registerFunc(&_authorizationRightSet, &_authorizationRightSetErr, frameworkHandle, "AuthorizationRightSet", "10.0")
	registerFunc(&_cMSDecoderCopyAllCerts, &_cMSDecoderCopyAllCertsErr, frameworkHandle, "CMSDecoderCopyAllCerts", "10.5")
	registerFunc(&_cMSDecoderCopyContent, &_cMSDecoderCopyContentErr, frameworkHandle, "CMSDecoderCopyContent", "10.5")
	registerFunc(&_cMSDecoderCopyDetachedContent, &_cMSDecoderCopyDetachedContentErr, frameworkHandle, "CMSDecoderCopyDetachedContent", "10.5")
	registerFunc(&_cMSDecoderCopyEncapsulatedContentType, &_cMSDecoderCopyEncapsulatedContentTypeErr, frameworkHandle, "CMSDecoderCopyEncapsulatedContentType", "10.5")
	registerFunc(&_cMSDecoderCopySignerCert, &_cMSDecoderCopySignerCertErr, frameworkHandle, "CMSDecoderCopySignerCert", "10.5")
	registerFunc(&_cMSDecoderCopySignerEmailAddress, &_cMSDecoderCopySignerEmailAddressErr, frameworkHandle, "CMSDecoderCopySignerEmailAddress", "10.5")
	registerFunc(&_cMSDecoderCopySignerSigningTime, &_cMSDecoderCopySignerSigningTimeErr, frameworkHandle, "CMSDecoderCopySignerSigningTime", "10.8")
	registerFunc(&_cMSDecoderCopySignerStatus, &_cMSDecoderCopySignerStatusErr, frameworkHandle, "CMSDecoderCopySignerStatus", "10.5")
	registerFunc(&_cMSDecoderCopySignerTimestamp, &_cMSDecoderCopySignerTimestampErr, frameworkHandle, "CMSDecoderCopySignerTimestamp", "10.8")
	registerFunc(&_cMSDecoderCopySignerTimestampCertificates, &_cMSDecoderCopySignerTimestampCertificatesErr, frameworkHandle, "CMSDecoderCopySignerTimestampCertificates", "10.8")
	registerFunc(&_cMSDecoderCopySignerTimestampWithPolicy, &_cMSDecoderCopySignerTimestampWithPolicyErr, frameworkHandle, "CMSDecoderCopySignerTimestampWithPolicy", "10.10")
	registerFunc(&_cMSDecoderCreate, &_cMSDecoderCreateErr, frameworkHandle, "CMSDecoderCreate", "10.5")
	registerFunc(&_cMSDecoderFinalizeMessage, &_cMSDecoderFinalizeMessageErr, frameworkHandle, "CMSDecoderFinalizeMessage", "10.5")
	registerFunc(&_cMSDecoderGetNumSigners, &_cMSDecoderGetNumSignersErr, frameworkHandle, "CMSDecoderGetNumSigners", "10.5")
	registerFunc(&_cMSDecoderGetTypeID, &_cMSDecoderGetTypeIDErr, frameworkHandle, "CMSDecoderGetTypeID", "10.0")
	registerFunc(&_cMSDecoderIsContentEncrypted, &_cMSDecoderIsContentEncryptedErr, frameworkHandle, "CMSDecoderIsContentEncrypted", "10.5")
	registerFunc(&_cMSDecoderSetDetachedContent, &_cMSDecoderSetDetachedContentErr, frameworkHandle, "CMSDecoderSetDetachedContent", "10.5")
	registerFunc(&_cMSDecoderSetSearchKeychain, &_cMSDecoderSetSearchKeychainErr, frameworkHandle, "CMSDecoderSetSearchKeychain", "10.5")
	registerFunc(&_cMSDecoderUpdateMessage, &_cMSDecoderUpdateMessageErr, frameworkHandle, "CMSDecoderUpdateMessage", "10.5")
	registerFunc(&_cMSEncode, &_cMSEncodeErr, frameworkHandle, "CMSEncode", "10.5")
	registerFunc(&_cMSEncodeContent, &_cMSEncodeContentErr, frameworkHandle, "CMSEncodeContent", "10.7")
	registerFunc(&_cMSEncoderAddRecipients, &_cMSEncoderAddRecipientsErr, frameworkHandle, "CMSEncoderAddRecipients", "10.5")
	registerFunc(&_cMSEncoderAddSignedAttributes, &_cMSEncoderAddSignedAttributesErr, frameworkHandle, "CMSEncoderAddSignedAttributes", "10.5")
	registerFunc(&_cMSEncoderAddSigners, &_cMSEncoderAddSignersErr, frameworkHandle, "CMSEncoderAddSigners", "10.5")
	registerFunc(&_cMSEncoderAddSupportingCerts, &_cMSEncoderAddSupportingCertsErr, frameworkHandle, "CMSEncoderAddSupportingCerts", "10.5")
	registerFunc(&_cMSEncoderCopyEncapsulatedContentType, &_cMSEncoderCopyEncapsulatedContentTypeErr, frameworkHandle, "CMSEncoderCopyEncapsulatedContentType", "10.5")
	registerFunc(&_cMSEncoderCopyEncodedContent, &_cMSEncoderCopyEncodedContentErr, frameworkHandle, "CMSEncoderCopyEncodedContent", "10.5")
	registerFunc(&_cMSEncoderCopyRecipients, &_cMSEncoderCopyRecipientsErr, frameworkHandle, "CMSEncoderCopyRecipients", "10.5")
	registerFunc(&_cMSEncoderCopySignerTimestamp, &_cMSEncoderCopySignerTimestampErr, frameworkHandle, "CMSEncoderCopySignerTimestamp", "10.8")
	registerFunc(&_cMSEncoderCopySignerTimestampWithPolicy, &_cMSEncoderCopySignerTimestampWithPolicyErr, frameworkHandle, "CMSEncoderCopySignerTimestampWithPolicy", "10.10")
	registerFunc(&_cMSEncoderCopySigners, &_cMSEncoderCopySignersErr, frameworkHandle, "CMSEncoderCopySigners", "10.5")
	registerFunc(&_cMSEncoderCopySupportingCerts, &_cMSEncoderCopySupportingCertsErr, frameworkHandle, "CMSEncoderCopySupportingCerts", "10.5")
	registerFunc(&_cMSEncoderCreate, &_cMSEncoderCreateErr, frameworkHandle, "CMSEncoderCreate", "10.5")
	registerFunc(&_cMSEncoderGetCertificateChainMode, &_cMSEncoderGetCertificateChainModeErr, frameworkHandle, "CMSEncoderGetCertificateChainMode", "10.5")
	registerFunc(&_cMSEncoderGetHasDetachedContent, &_cMSEncoderGetHasDetachedContentErr, frameworkHandle, "CMSEncoderGetHasDetachedContent", "10.5")
	registerFunc(&_cMSEncoderGetTypeID, &_cMSEncoderGetTypeIDErr, frameworkHandle, "CMSEncoderGetTypeID", "10.5")
	registerFunc(&_cMSEncoderSetCertificateChainMode, &_cMSEncoderSetCertificateChainModeErr, frameworkHandle, "CMSEncoderSetCertificateChainMode", "10.5")
	registerFunc(&_cMSEncoderSetEncapsulatedContentType, &_cMSEncoderSetEncapsulatedContentTypeErr, frameworkHandle, "CMSEncoderSetEncapsulatedContentType", "10.5")
	registerFunc(&_cMSEncoderSetEncapsulatedContentTypeOID, &_cMSEncoderSetEncapsulatedContentTypeOIDErr, frameworkHandle, "CMSEncoderSetEncapsulatedContentTypeOID", "10.7")
	registerFunc(&_cMSEncoderSetHasDetachedContent, &_cMSEncoderSetHasDetachedContentErr, frameworkHandle, "CMSEncoderSetHasDetachedContent", "10.5")
	registerFunc(&_cMSEncoderSetSignerAlgorithm, &_cMSEncoderSetSignerAlgorithmErr, frameworkHandle, "CMSEncoderSetSignerAlgorithm", "10.11")
	registerFunc(&_cMSEncoderUpdateContent, &_cMSEncoderUpdateContentErr, frameworkHandle, "CMSEncoderUpdateContent", "10.5")
	registerFunc(&_cSSM_AC_AuthCompute, &_cSSM_AC_AuthComputeErr, frameworkHandle, "CSSM_AC_AuthCompute", "10.0")
	registerFunc(&_cSSM_AC_PassThrough, &_cSSM_AC_PassThroughErr, frameworkHandle, "CSSM_AC_PassThrough", "10.0")
	registerFunc(&_cSSM_CL_CertAbortCache, &_cSSM_CL_CertAbortCacheErr, frameworkHandle, "CSSM_CL_CertAbortCache", "10.0")
	registerFunc(&_cSSM_CL_CertAbortQuery, &_cSSM_CL_CertAbortQueryErr, frameworkHandle, "CSSM_CL_CertAbortQuery", "10.0")
	registerFunc(&_cSSM_CL_CertCache, &_cSSM_CL_CertCacheErr, frameworkHandle, "CSSM_CL_CertCache", "10.0")
	registerFunc(&_cSSM_CL_CertCreateTemplate, &_cSSM_CL_CertCreateTemplateErr, frameworkHandle, "CSSM_CL_CertCreateTemplate", "10.0")
	registerFunc(&_cSSM_CL_CertDescribeFormat, &_cSSM_CL_CertDescribeFormatErr, frameworkHandle, "CSSM_CL_CertDescribeFormat", "10.0")
	registerFunc(&_cSSM_CL_CertGetAllFields, &_cSSM_CL_CertGetAllFieldsErr, frameworkHandle, "CSSM_CL_CertGetAllFields", "10.0")
	registerFunc(&_cSSM_CL_CertGetAllTemplateFields, &_cSSM_CL_CertGetAllTemplateFieldsErr, frameworkHandle, "CSSM_CL_CertGetAllTemplateFields", "10.0")
	registerFunc(&_cSSM_CL_CertGetFirstCachedFieldValue, &_cSSM_CL_CertGetFirstCachedFieldValueErr, frameworkHandle, "CSSM_CL_CertGetFirstCachedFieldValue", "10.0")
	registerFunc(&_cSSM_CL_CertGetFirstFieldValue, &_cSSM_CL_CertGetFirstFieldValueErr, frameworkHandle, "CSSM_CL_CertGetFirstFieldValue", "10.0")
	registerFunc(&_cSSM_CL_CertGetKeyInfo, &_cSSM_CL_CertGetKeyInfoErr, frameworkHandle, "CSSM_CL_CertGetKeyInfo", "10.0")
	registerFunc(&_cSSM_CL_CertGetNextCachedFieldValue, &_cSSM_CL_CertGetNextCachedFieldValueErr, frameworkHandle, "CSSM_CL_CertGetNextCachedFieldValue", "10.0")
	registerFunc(&_cSSM_CL_CertGetNextFieldValue, &_cSSM_CL_CertGetNextFieldValueErr, frameworkHandle, "CSSM_CL_CertGetNextFieldValue", "10.0")
	registerFunc(&_cSSM_CL_CertGroupFromVerifiedBundle, &_cSSM_CL_CertGroupFromVerifiedBundleErr, frameworkHandle, "CSSM_CL_CertGroupFromVerifiedBundle", "10.0")
	registerFunc(&_cSSM_CL_CertGroupToSignedBundle, &_cSSM_CL_CertGroupToSignedBundleErr, frameworkHandle, "CSSM_CL_CertGroupToSignedBundle", "10.0")
	registerFunc(&_cSSM_CL_CertSign, &_cSSM_CL_CertSignErr, frameworkHandle, "CSSM_CL_CertSign", "10.0")
	registerFunc(&_cSSM_CL_CertVerify, &_cSSM_CL_CertVerifyErr, frameworkHandle, "CSSM_CL_CertVerify", "10.0")
	registerFunc(&_cSSM_CL_CertVerifyWithKey, &_cSSM_CL_CertVerifyWithKeyErr, frameworkHandle, "CSSM_CL_CertVerifyWithKey", "10.0")
	registerFunc(&_cSSM_CL_CrlAbortCache, &_cSSM_CL_CrlAbortCacheErr, frameworkHandle, "CSSM_CL_CrlAbortCache", "10.0")
	registerFunc(&_cSSM_CL_CrlAbortQuery, &_cSSM_CL_CrlAbortQueryErr, frameworkHandle, "CSSM_CL_CrlAbortQuery", "10.0")
	registerFunc(&_cSSM_CL_CrlAddCert, &_cSSM_CL_CrlAddCertErr, frameworkHandle, "CSSM_CL_CrlAddCert", "10.0")
	registerFunc(&_cSSM_CL_CrlCache, &_cSSM_CL_CrlCacheErr, frameworkHandle, "CSSM_CL_CrlCache", "10.0")
	registerFunc(&_cSSM_CL_CrlCreateTemplate, &_cSSM_CL_CrlCreateTemplateErr, frameworkHandle, "CSSM_CL_CrlCreateTemplate", "10.0")
	registerFunc(&_cSSM_CL_CrlDescribeFormat, &_cSSM_CL_CrlDescribeFormatErr, frameworkHandle, "CSSM_CL_CrlDescribeFormat", "10.0")
	registerFunc(&_cSSM_CL_CrlGetAllCachedRecordFields, &_cSSM_CL_CrlGetAllCachedRecordFieldsErr, frameworkHandle, "CSSM_CL_CrlGetAllCachedRecordFields", "10.0")
	registerFunc(&_cSSM_CL_CrlGetAllFields, &_cSSM_CL_CrlGetAllFieldsErr, frameworkHandle, "CSSM_CL_CrlGetAllFields", "10.0")
	registerFunc(&_cSSM_CL_CrlGetFirstCachedFieldValue, &_cSSM_CL_CrlGetFirstCachedFieldValueErr, frameworkHandle, "CSSM_CL_CrlGetFirstCachedFieldValue", "10.0")
	registerFunc(&_cSSM_CL_CrlGetFirstFieldValue, &_cSSM_CL_CrlGetFirstFieldValueErr, frameworkHandle, "CSSM_CL_CrlGetFirstFieldValue", "10.0")
	registerFunc(&_cSSM_CL_CrlGetNextCachedFieldValue, &_cSSM_CL_CrlGetNextCachedFieldValueErr, frameworkHandle, "CSSM_CL_CrlGetNextCachedFieldValue", "10.0")
	registerFunc(&_cSSM_CL_CrlGetNextFieldValue, &_cSSM_CL_CrlGetNextFieldValueErr, frameworkHandle, "CSSM_CL_CrlGetNextFieldValue", "10.0")
	registerFunc(&_cSSM_CL_CrlRemoveCert, &_cSSM_CL_CrlRemoveCertErr, frameworkHandle, "CSSM_CL_CrlRemoveCert", "10.0")
	registerFunc(&_cSSM_CL_CrlSetFields, &_cSSM_CL_CrlSetFieldsErr, frameworkHandle, "CSSM_CL_CrlSetFields", "10.0")
	registerFunc(&_cSSM_CL_CrlSign, &_cSSM_CL_CrlSignErr, frameworkHandle, "CSSM_CL_CrlSign", "10.0")
	registerFunc(&_cSSM_CL_CrlVerify, &_cSSM_CL_CrlVerifyErr, frameworkHandle, "CSSM_CL_CrlVerify", "10.0")
	registerFunc(&_cSSM_CL_CrlVerifyWithKey, &_cSSM_CL_CrlVerifyWithKeyErr, frameworkHandle, "CSSM_CL_CrlVerifyWithKey", "10.0")
	registerFunc(&_cSSM_CL_FreeFieldValue, &_cSSM_CL_FreeFieldValueErr, frameworkHandle, "CSSM_CL_FreeFieldValue", "10.0")
	registerFunc(&_cSSM_CL_FreeFields, &_cSSM_CL_FreeFieldsErr, frameworkHandle, "CSSM_CL_FreeFields", "10.0")
	registerFunc(&_cSSM_CL_IsCertInCachedCrl, &_cSSM_CL_IsCertInCachedCrlErr, frameworkHandle, "CSSM_CL_IsCertInCachedCrl", "10.0")
	registerFunc(&_cSSM_CL_IsCertInCrl, &_cSSM_CL_IsCertInCrlErr, frameworkHandle, "CSSM_CL_IsCertInCrl", "10.0")
	registerFunc(&_cSSM_CL_PassThrough, &_cSSM_CL_PassThroughErr, frameworkHandle, "CSSM_CL_PassThrough", "10.0")
	registerFunc(&_cSSM_CSP_ChangeLoginAcl, &_cSSM_CSP_ChangeLoginAclErr, frameworkHandle, "CSSM_CSP_ChangeLoginAcl", "10.0")
	registerFunc(&_cSSM_CSP_ChangeLoginOwner, &_cSSM_CSP_ChangeLoginOwnerErr, frameworkHandle, "CSSM_CSP_ChangeLoginOwner", "10.0")
	registerFunc(&_cSSM_CSP_CreateAsymmetricContext, &_cSSM_CSP_CreateAsymmetricContextErr, frameworkHandle, "CSSM_CSP_CreateAsymmetricContext", "10.0")
	registerFunc(&_cSSM_CSP_CreateDeriveKeyContext, &_cSSM_CSP_CreateDeriveKeyContextErr, frameworkHandle, "CSSM_CSP_CreateDeriveKeyContext", "10.0")
	registerFunc(&_cSSM_CSP_CreateDigestContext, &_cSSM_CSP_CreateDigestContextErr, frameworkHandle, "CSSM_CSP_CreateDigestContext", "10.0")
	registerFunc(&_cSSM_CSP_CreateKeyGenContext, &_cSSM_CSP_CreateKeyGenContextErr, frameworkHandle, "CSSM_CSP_CreateKeyGenContext", "10.0")
	registerFunc(&_cSSM_CSP_CreateMacContext, &_cSSM_CSP_CreateMacContextErr, frameworkHandle, "CSSM_CSP_CreateMacContext", "10.0")
	registerFunc(&_cSSM_CSP_CreatePassThroughContext, &_cSSM_CSP_CreatePassThroughContextErr, frameworkHandle, "CSSM_CSP_CreatePassThroughContext", "10.0")
	registerFunc(&_cSSM_CSP_CreateRandomGenContext, &_cSSM_CSP_CreateRandomGenContextErr, frameworkHandle, "CSSM_CSP_CreateRandomGenContext", "10.0")
	registerFunc(&_cSSM_CSP_CreateSignatureContext, &_cSSM_CSP_CreateSignatureContextErr, frameworkHandle, "CSSM_CSP_CreateSignatureContext", "10.0")
	registerFunc(&_cSSM_CSP_CreateSymmetricContext, &_cSSM_CSP_CreateSymmetricContextErr, frameworkHandle, "CSSM_CSP_CreateSymmetricContext", "10.0")
	registerFunc(&_cSSM_CSP_GetLoginAcl, &_cSSM_CSP_GetLoginAclErr, frameworkHandle, "CSSM_CSP_GetLoginAcl", "10.0")
	registerFunc(&_cSSM_CSP_GetLoginOwner, &_cSSM_CSP_GetLoginOwnerErr, frameworkHandle, "CSSM_CSP_GetLoginOwner", "10.0")
	registerFunc(&_cSSM_CSP_GetOperationalStatistics, &_cSSM_CSP_GetOperationalStatisticsErr, frameworkHandle, "CSSM_CSP_GetOperationalStatistics", "10.0")
	registerFunc(&_cSSM_CSP_Login, &_cSSM_CSP_LoginErr, frameworkHandle, "CSSM_CSP_Login", "10.0")
	registerFunc(&_cSSM_CSP_Logout, &_cSSM_CSP_LogoutErr, frameworkHandle, "CSSM_CSP_Logout", "10.0")
	registerFunc(&_cSSM_CSP_ObtainPrivateKeyFromPublicKey, &_cSSM_CSP_ObtainPrivateKeyFromPublicKeyErr, frameworkHandle, "CSSM_CSP_ObtainPrivateKeyFromPublicKey", "10.0")
	registerFunc(&_cSSM_CSP_PassThrough, &_cSSM_CSP_PassThroughErr, frameworkHandle, "CSSM_CSP_PassThrough", "10.0")
	registerFunc(&_cSSM_ChangeKeyAcl, &_cSSM_ChangeKeyAclErr, frameworkHandle, "CSSM_ChangeKeyAcl", "10.0")
	registerFunc(&_cSSM_ChangeKeyOwner, &_cSSM_ChangeKeyOwnerErr, frameworkHandle, "CSSM_ChangeKeyOwner", "10.0")
	registerFunc(&_cSSM_DL_Authenticate, &_cSSM_DL_AuthenticateErr, frameworkHandle, "CSSM_DL_Authenticate", "10.0")
	registerFunc(&_cSSM_DL_ChangeDbAcl, &_cSSM_DL_ChangeDbAclErr, frameworkHandle, "CSSM_DL_ChangeDbAcl", "10.0")
	registerFunc(&_cSSM_DL_ChangeDbOwner, &_cSSM_DL_ChangeDbOwnerErr, frameworkHandle, "CSSM_DL_ChangeDbOwner", "10.0")
	registerFunc(&_cSSM_DL_CreateRelation, &_cSSM_DL_CreateRelationErr, frameworkHandle, "CSSM_DL_CreateRelation", "10.0")
	registerFunc(&_cSSM_DL_DataAbortQuery, &_cSSM_DL_DataAbortQueryErr, frameworkHandle, "CSSM_DL_DataAbortQuery", "10.0")
	registerFunc(&_cSSM_DL_DataDelete, &_cSSM_DL_DataDeleteErr, frameworkHandle, "CSSM_DL_DataDelete", "10.0")
	registerFunc(&_cSSM_DL_DataGetFirst, &_cSSM_DL_DataGetFirstErr, frameworkHandle, "CSSM_DL_DataGetFirst", "10.0")
	registerFunc(&_cSSM_DL_DataGetFromUniqueRecordId, &_cSSM_DL_DataGetFromUniqueRecordIdErr, frameworkHandle, "CSSM_DL_DataGetFromUniqueRecordId", "10.0")
	registerFunc(&_cSSM_DL_DataGetNext, &_cSSM_DL_DataGetNextErr, frameworkHandle, "CSSM_DL_DataGetNext", "10.0")
	registerFunc(&_cSSM_DL_DataInsert, &_cSSM_DL_DataInsertErr, frameworkHandle, "CSSM_DL_DataInsert", "10.0")
	registerFunc(&_cSSM_DL_DataModify, &_cSSM_DL_DataModifyErr, frameworkHandle, "CSSM_DL_DataModify", "10.0")
	registerFunc(&_cSSM_DL_DbClose, &_cSSM_DL_DbCloseErr, frameworkHandle, "CSSM_DL_DbClose", "10.0")
	registerFunc(&_cSSM_DL_DbCreate, &_cSSM_DL_DbCreateErr, frameworkHandle, "CSSM_DL_DbCreate", "10.0")
	registerFunc(&_cSSM_DL_DbDelete, &_cSSM_DL_DbDeleteErr, frameworkHandle, "CSSM_DL_DbDelete", "10.0")
	registerFunc(&_cSSM_DL_DbOpen, &_cSSM_DL_DbOpenErr, frameworkHandle, "CSSM_DL_DbOpen", "10.0")
	registerFunc(&_cSSM_DL_DestroyRelation, &_cSSM_DL_DestroyRelationErr, frameworkHandle, "CSSM_DL_DestroyRelation", "10.0")
	registerFunc(&_cSSM_DL_FreeNameList, &_cSSM_DL_FreeNameListErr, frameworkHandle, "CSSM_DL_FreeNameList", "10.0")
	registerFunc(&_cSSM_DL_FreeUniqueRecord, &_cSSM_DL_FreeUniqueRecordErr, frameworkHandle, "CSSM_DL_FreeUniqueRecord", "10.0")
	registerFunc(&_cSSM_DL_GetDbAcl, &_cSSM_DL_GetDbAclErr, frameworkHandle, "CSSM_DL_GetDbAcl", "10.0")
	registerFunc(&_cSSM_DL_GetDbNameFromHandle, &_cSSM_DL_GetDbNameFromHandleErr, frameworkHandle, "CSSM_DL_GetDbNameFromHandle", "10.0")
	registerFunc(&_cSSM_DL_GetDbNames, &_cSSM_DL_GetDbNamesErr, frameworkHandle, "CSSM_DL_GetDbNames", "10.0")
	registerFunc(&_cSSM_DL_GetDbOwner, &_cSSM_DL_GetDbOwnerErr, frameworkHandle, "CSSM_DL_GetDbOwner", "10.0")
	registerFunc(&_cSSM_DL_PassThrough, &_cSSM_DL_PassThroughErr, frameworkHandle, "CSSM_DL_PassThrough", "10.0")
	registerFunc(&_cSSM_DecryptData, &_cSSM_DecryptDataErr, frameworkHandle, "CSSM_DecryptData", "10.0")
	registerFunc(&_cSSM_DecryptDataFinal, &_cSSM_DecryptDataFinalErr, frameworkHandle, "CSSM_DecryptDataFinal", "10.0")
	registerFunc(&_cSSM_DecryptDataInit, &_cSSM_DecryptDataInitErr, frameworkHandle, "CSSM_DecryptDataInit", "10.0")
	registerFunc(&_cSSM_DecryptDataInitP, &_cSSM_DecryptDataInitPErr, frameworkHandle, "CSSM_DecryptDataInitP", "10.0")
	registerFunc(&_cSSM_DecryptDataP, &_cSSM_DecryptDataPErr, frameworkHandle, "CSSM_DecryptDataP", "10.0")
	registerFunc(&_cSSM_DecryptDataUpdate, &_cSSM_DecryptDataUpdateErr, frameworkHandle, "CSSM_DecryptDataUpdate", "10.0")
	registerFunc(&_cSSM_DeleteContext, &_cSSM_DeleteContextErr, frameworkHandle, "CSSM_DeleteContext", "10.0")
	registerFunc(&_cSSM_DeleteContextAttributes, &_cSSM_DeleteContextAttributesErr, frameworkHandle, "CSSM_DeleteContextAttributes", "10.0")
	registerFunc(&_cSSM_DeriveKey, &_cSSM_DeriveKeyErr, frameworkHandle, "CSSM_DeriveKey", "10.0")
	registerFunc(&_cSSM_DigestData, &_cSSM_DigestDataErr, frameworkHandle, "CSSM_DigestData", "10.0")
	registerFunc(&_cSSM_DigestDataClone, &_cSSM_DigestDataCloneErr, frameworkHandle, "CSSM_DigestDataClone", "10.0")
	registerFunc(&_cSSM_DigestDataFinal, &_cSSM_DigestDataFinalErr, frameworkHandle, "CSSM_DigestDataFinal", "10.0")
	registerFunc(&_cSSM_DigestDataInit, &_cSSM_DigestDataInitErr, frameworkHandle, "CSSM_DigestDataInit", "10.0")
	registerFunc(&_cSSM_DigestDataUpdate, &_cSSM_DigestDataUpdateErr, frameworkHandle, "CSSM_DigestDataUpdate", "10.0")
	registerFunc(&_cSSM_EncryptData, &_cSSM_EncryptDataErr, frameworkHandle, "CSSM_EncryptData", "10.0")
	registerFunc(&_cSSM_EncryptDataFinal, &_cSSM_EncryptDataFinalErr, frameworkHandle, "CSSM_EncryptDataFinal", "10.0")
	registerFunc(&_cSSM_EncryptDataInit, &_cSSM_EncryptDataInitErr, frameworkHandle, "CSSM_EncryptDataInit", "10.0")
	registerFunc(&_cSSM_EncryptDataInitP, &_cSSM_EncryptDataInitPErr, frameworkHandle, "CSSM_EncryptDataInitP", "10.0")
	registerFunc(&_cSSM_EncryptDataP, &_cSSM_EncryptDataPErr, frameworkHandle, "CSSM_EncryptDataP", "10.0")
	registerFunc(&_cSSM_EncryptDataUpdate, &_cSSM_EncryptDataUpdateErr, frameworkHandle, "CSSM_EncryptDataUpdate", "10.0")
	registerFunc(&_cSSM_FreeContext, &_cSSM_FreeContextErr, frameworkHandle, "CSSM_FreeContext", "10.0")
	registerFunc(&_cSSM_FreeKey, &_cSSM_FreeKeyErr, frameworkHandle, "CSSM_FreeKey", "10.0")
	registerFunc(&_cSSM_GenerateAlgorithmParams, &_cSSM_GenerateAlgorithmParamsErr, frameworkHandle, "CSSM_GenerateAlgorithmParams", "10.0")
	registerFunc(&_cSSM_GenerateKey, &_cSSM_GenerateKeyErr, frameworkHandle, "CSSM_GenerateKey", "10.0")
	registerFunc(&_cSSM_GenerateKeyP, &_cSSM_GenerateKeyPErr, frameworkHandle, "CSSM_GenerateKeyP", "10.0")
	registerFunc(&_cSSM_GenerateKeyPair, &_cSSM_GenerateKeyPairErr, frameworkHandle, "CSSM_GenerateKeyPair", "10.0")
	registerFunc(&_cSSM_GenerateKeyPairP, &_cSSM_GenerateKeyPairPErr, frameworkHandle, "CSSM_GenerateKeyPairP", "10.0")
	registerFunc(&_cSSM_GenerateMac, &_cSSM_GenerateMacErr, frameworkHandle, "CSSM_GenerateMac", "10.0")
	registerFunc(&_cSSM_GenerateMacFinal, &_cSSM_GenerateMacFinalErr, frameworkHandle, "CSSM_GenerateMacFinal", "10.0")
	registerFunc(&_cSSM_GenerateMacInit, &_cSSM_GenerateMacInitErr, frameworkHandle, "CSSM_GenerateMacInit", "10.0")
	registerFunc(&_cSSM_GenerateMacUpdate, &_cSSM_GenerateMacUpdateErr, frameworkHandle, "CSSM_GenerateMacUpdate", "10.0")
	registerFunc(&_cSSM_GenerateRandom, &_cSSM_GenerateRandomErr, frameworkHandle, "CSSM_GenerateRandom", "10.0")
	registerFunc(&_cSSM_GetAPIMemoryFunctions, &_cSSM_GetAPIMemoryFunctionsErr, frameworkHandle, "CSSM_GetAPIMemoryFunctions", "10.0")
	registerFunc(&_cSSM_GetContext, &_cSSM_GetContextErr, frameworkHandle, "CSSM_GetContext", "10.0")
	registerFunc(&_cSSM_GetContextAttribute, &_cSSM_GetContextAttributeErr, frameworkHandle, "CSSM_GetContextAttribute", "10.0")
	registerFunc(&_cSSM_GetKeyAcl, &_cSSM_GetKeyAclErr, frameworkHandle, "CSSM_GetKeyAcl", "10.0")
	registerFunc(&_cSSM_GetKeyOwner, &_cSSM_GetKeyOwnerErr, frameworkHandle, "CSSM_GetKeyOwner", "10.0")
	registerFunc(&_cSSM_GetModuleGUIDFromHandle, &_cSSM_GetModuleGUIDFromHandleErr, frameworkHandle, "CSSM_GetModuleGUIDFromHandle", "10.0")
	registerFunc(&_cSSM_GetPrivilege, &_cSSM_GetPrivilegeErr, frameworkHandle, "CSSM_GetPrivilege", "10.0")
	registerFunc(&_cSSM_GetSubserviceUIDFromHandle, &_cSSM_GetSubserviceUIDFromHandleErr, frameworkHandle, "CSSM_GetSubserviceUIDFromHandle", "10.0")
	registerFunc(&_cSSM_GetTimeValue, &_cSSM_GetTimeValueErr, frameworkHandle, "CSSM_GetTimeValue", "10.0")
	registerFunc(&_cSSM_Init, &_cSSM_InitErr, frameworkHandle, "CSSM_Init", "10.0")
	registerFunc(&_cSSM_Introduce, &_cSSM_IntroduceErr, frameworkHandle, "CSSM_Introduce", "10.0")
	registerFunc(&_cSSM_ListAttachedModuleManagers, &_cSSM_ListAttachedModuleManagersErr, frameworkHandle, "CSSM_ListAttachedModuleManagers", "10.0")
	registerFunc(&_cSSM_ModuleAttach, &_cSSM_ModuleAttachErr, frameworkHandle, "CSSM_ModuleAttach", "10.0")
	registerFunc(&_cSSM_ModuleDetach, &_cSSM_ModuleDetachErr, frameworkHandle, "CSSM_ModuleDetach", "10.0")
	registerFunc(&_cSSM_ModuleLoad, &_cSSM_ModuleLoadErr, frameworkHandle, "CSSM_ModuleLoad", "10.0")
	registerFunc(&_cSSM_ModuleUnload, &_cSSM_ModuleUnloadErr, frameworkHandle, "CSSM_ModuleUnload", "10.0")
	registerFunc(&_cSSM_QueryKeySizeInBits, &_cSSM_QueryKeySizeInBitsErr, frameworkHandle, "CSSM_QueryKeySizeInBits", "10.0")
	registerFunc(&_cSSM_QuerySize, &_cSSM_QuerySizeErr, frameworkHandle, "CSSM_QuerySize", "10.0")
	registerFunc(&_cSSM_RetrieveCounter, &_cSSM_RetrieveCounterErr, frameworkHandle, "CSSM_RetrieveCounter", "10.0")
	registerFunc(&_cSSM_RetrieveUniqueId, &_cSSM_RetrieveUniqueIdErr, frameworkHandle, "CSSM_RetrieveUniqueId", "10.0")
	registerFunc(&_cSSM_SetContext, &_cSSM_SetContextErr, frameworkHandle, "CSSM_SetContext", "10.0")
	registerFunc(&_cSSM_SetPrivilege, &_cSSM_SetPrivilegeErr, frameworkHandle, "CSSM_SetPrivilege", "10.0")
	registerFunc(&_cSSM_SignData, &_cSSM_SignDataErr, frameworkHandle, "CSSM_SignData", "10.0")
	registerFunc(&_cSSM_SignDataFinal, &_cSSM_SignDataFinalErr, frameworkHandle, "CSSM_SignDataFinal", "10.0")
	registerFunc(&_cSSM_SignDataInit, &_cSSM_SignDataInitErr, frameworkHandle, "CSSM_SignDataInit", "10.0")
	registerFunc(&_cSSM_SignDataUpdate, &_cSSM_SignDataUpdateErr, frameworkHandle, "CSSM_SignDataUpdate", "10.0")
	registerFunc(&_cSSM_TP_ApplyCrlToDb, &_cSSM_TP_ApplyCrlToDbErr, frameworkHandle, "CSSM_TP_ApplyCrlToDb", "10.0")
	registerFunc(&_cSSM_TP_CertCreateTemplate, &_cSSM_TP_CertCreateTemplateErr, frameworkHandle, "CSSM_TP_CertCreateTemplate", "10.0")
	registerFunc(&_cSSM_TP_CertGetAllTemplateFields, &_cSSM_TP_CertGetAllTemplateFieldsErr, frameworkHandle, "CSSM_TP_CertGetAllTemplateFields", "10.0")
	registerFunc(&_cSSM_TP_CertGroupConstruct, &_cSSM_TP_CertGroupConstructErr, frameworkHandle, "CSSM_TP_CertGroupConstruct", "10.0")
	registerFunc(&_cSSM_TP_CertGroupPrune, &_cSSM_TP_CertGroupPruneErr, frameworkHandle, "CSSM_TP_CertGroupPrune", "10.0")
	registerFunc(&_cSSM_TP_CertGroupToTupleGroup, &_cSSM_TP_CertGroupToTupleGroupErr, frameworkHandle, "CSSM_TP_CertGroupToTupleGroup", "10.0")
	registerFunc(&_cSSM_TP_CertGroupVerify, &_cSSM_TP_CertGroupVerifyErr, frameworkHandle, "CSSM_TP_CertGroupVerify", "10.0")
	registerFunc(&_cSSM_TP_CertReclaimAbort, &_cSSM_TP_CertReclaimAbortErr, frameworkHandle, "CSSM_TP_CertReclaimAbort", "10.0")
	registerFunc(&_cSSM_TP_CertReclaimKey, &_cSSM_TP_CertReclaimKeyErr, frameworkHandle, "CSSM_TP_CertReclaimKey", "10.0")
	registerFunc(&_cSSM_TP_CertRemoveFromCrlTemplate, &_cSSM_TP_CertRemoveFromCrlTemplateErr, frameworkHandle, "CSSM_TP_CertRemoveFromCrlTemplate", "10.0")
	registerFunc(&_cSSM_TP_CertRevoke, &_cSSM_TP_CertRevokeErr, frameworkHandle, "CSSM_TP_CertRevoke", "10.0")
	registerFunc(&_cSSM_TP_CertSign, &_cSSM_TP_CertSignErr, frameworkHandle, "CSSM_TP_CertSign", "10.0")
	registerFunc(&_cSSM_TP_ConfirmCredResult, &_cSSM_TP_ConfirmCredResultErr, frameworkHandle, "CSSM_TP_ConfirmCredResult", "10.0")
	registerFunc(&_cSSM_TP_CrlCreateTemplate, &_cSSM_TP_CrlCreateTemplateErr, frameworkHandle, "CSSM_TP_CrlCreateTemplate", "10.0")
	registerFunc(&_cSSM_TP_CrlSign, &_cSSM_TP_CrlSignErr, frameworkHandle, "CSSM_TP_CrlSign", "10.0")
	registerFunc(&_cSSM_TP_CrlVerify, &_cSSM_TP_CrlVerifyErr, frameworkHandle, "CSSM_TP_CrlVerify", "10.0")
	registerFunc(&_cSSM_TP_FormRequest, &_cSSM_TP_FormRequestErr, frameworkHandle, "CSSM_TP_FormRequest", "10.0")
	registerFunc(&_cSSM_TP_FormSubmit, &_cSSM_TP_FormSubmitErr, frameworkHandle, "CSSM_TP_FormSubmit", "10.0")
	registerFunc(&_cSSM_TP_PassThrough, &_cSSM_TP_PassThroughErr, frameworkHandle, "CSSM_TP_PassThrough", "10.0")
	registerFunc(&_cSSM_TP_ReceiveConfirmation, &_cSSM_TP_ReceiveConfirmationErr, frameworkHandle, "CSSM_TP_ReceiveConfirmation", "10.0")
	registerFunc(&_cSSM_TP_RetrieveCredResult, &_cSSM_TP_RetrieveCredResultErr, frameworkHandle, "CSSM_TP_RetrieveCredResult", "10.0")
	registerFunc(&_cSSM_TP_SubmitCredRequest, &_cSSM_TP_SubmitCredRequestErr, frameworkHandle, "CSSM_TP_SubmitCredRequest", "10.0")
	registerFunc(&_cSSM_TP_TupleGroupToCertGroup, &_cSSM_TP_TupleGroupToCertGroupErr, frameworkHandle, "CSSM_TP_TupleGroupToCertGroup", "10.0")
	registerFunc(&_cSSM_Terminate, &_cSSM_TerminateErr, frameworkHandle, "CSSM_Terminate", "10.0")
	registerFunc(&_cSSM_Unintroduce, &_cSSM_UnintroduceErr, frameworkHandle, "CSSM_Unintroduce", "10.0")
	registerFunc(&_cSSM_UnwrapKey, &_cSSM_UnwrapKeyErr, frameworkHandle, "CSSM_UnwrapKey", "10.0")
	registerFunc(&_cSSM_UnwrapKeyP, &_cSSM_UnwrapKeyPErr, frameworkHandle, "CSSM_UnwrapKeyP", "10.0")
	registerFunc(&_cSSM_UpdateContextAttributes, &_cSSM_UpdateContextAttributesErr, frameworkHandle, "CSSM_UpdateContextAttributes", "10.0")
	registerFunc(&_cSSM_VerifyData, &_cSSM_VerifyDataErr, frameworkHandle, "CSSM_VerifyData", "10.0")
	registerFunc(&_cSSM_VerifyDataFinal, &_cSSM_VerifyDataFinalErr, frameworkHandle, "CSSM_VerifyDataFinal", "10.0")
	registerFunc(&_cSSM_VerifyDataInit, &_cSSM_VerifyDataInitErr, frameworkHandle, "CSSM_VerifyDataInit", "10.0")
	registerFunc(&_cSSM_VerifyDataUpdate, &_cSSM_VerifyDataUpdateErr, frameworkHandle, "CSSM_VerifyDataUpdate", "10.0")
	registerFunc(&_cSSM_VerifyDevice, &_cSSM_VerifyDeviceErr, frameworkHandle, "CSSM_VerifyDevice", "10.0")
	registerFunc(&_cSSM_VerifyMac, &_cSSM_VerifyMacErr, frameworkHandle, "CSSM_VerifyMac", "10.0")
	registerFunc(&_cSSM_VerifyMacFinal, &_cSSM_VerifyMacFinalErr, frameworkHandle, "CSSM_VerifyMacFinal", "10.0")
	registerFunc(&_cSSM_VerifyMacInit, &_cSSM_VerifyMacInitErr, frameworkHandle, "CSSM_VerifyMacInit", "10.0")
	registerFunc(&_cSSM_VerifyMacUpdate, &_cSSM_VerifyMacUpdateErr, frameworkHandle, "CSSM_VerifyMacUpdate", "10.0")
	registerFunc(&_cSSM_WrapKey, &_cSSM_WrapKeyErr, frameworkHandle, "CSSM_WrapKey", "10.0")
	registerFunc(&_cSSM_WrapKeyP, &_cSSM_WrapKeyPErr, frameworkHandle, "CSSM_WrapKeyP", "10.0")
	registerFunc(&_mDS_Initialize, &_mDS_InitializeErr, frameworkHandle, "MDS_Initialize", "10.0")
	registerFunc(&_mDS_Install, &_mDS_InstallErr, frameworkHandle, "MDS_Install", "10.0")
	registerFunc(&_mDS_Terminate, &_mDS_TerminateErr, frameworkHandle, "MDS_Terminate", "10.0")
	registerFunc(&_mDS_Uninstall, &_mDS_UninstallErr, frameworkHandle, "MDS_Uninstall", "10.0")
	registerFunc(&_secAccessControlCreateWithFlags, &_secAccessControlCreateWithFlagsErr, frameworkHandle, "SecAccessControlCreateWithFlags", "10.10")
	registerFunc(&_secAccessControlGetTypeID, &_secAccessControlGetTypeIDErr, frameworkHandle, "SecAccessControlGetTypeID", "10.10")
	registerFunc(&_secAddSharedWebCredential, &_secAddSharedWebCredentialErr, frameworkHandle, "SecAddSharedWebCredential", "11.0")
	registerFunc(&_secAsn1AllocCopy, &_secAsn1AllocCopyErr, frameworkHandle, "SecAsn1AllocCopy", "10.0")
	registerFunc(&_secAsn1AllocCopyItem, &_secAsn1AllocCopyItemErr, frameworkHandle, "SecAsn1AllocCopyItem", "10.0")
	registerFunc(&_secAsn1AllocItem, &_secAsn1AllocItemErr, frameworkHandle, "SecAsn1AllocItem", "10.0")
	registerFunc(&_secAsn1CoderCreate, &_secAsn1CoderCreateErr, frameworkHandle, "SecAsn1CoderCreate", "10.0")
	registerFunc(&_secAsn1CoderRelease, &_secAsn1CoderReleaseErr, frameworkHandle, "SecAsn1CoderRelease", "10.0")
	registerFunc(&_secAsn1Decode, &_secAsn1DecodeErr, frameworkHandle, "SecAsn1Decode", "10.0")
	registerFunc(&_secAsn1DecodeData, &_secAsn1DecodeDataErr, frameworkHandle, "SecAsn1DecodeData", "10.0")
	registerFunc(&_secAsn1EncodeItem, &_secAsn1EncodeItemErr, frameworkHandle, "SecAsn1EncodeItem", "10.0")
	registerFunc(&_secAsn1Malloc, &_secAsn1MallocErr, frameworkHandle, "SecAsn1Malloc", "10.0")
	registerFunc(&_secAsn1OidCompare, &_secAsn1OidCompareErr, frameworkHandle, "SecAsn1OidCompare", "10.0")
	registerFunc(&_secCertificateAddToKeychain, &_secCertificateAddToKeychainErr, frameworkHandle, "SecCertificateAddToKeychain", "10.3")
	registerFunc(&_secCertificateCopyCommonName, &_secCertificateCopyCommonNameErr, frameworkHandle, "SecCertificateCopyCommonName", "10.5")
	registerFunc(&_secCertificateCopyData, &_secCertificateCopyDataErr, frameworkHandle, "SecCertificateCopyData", "10.6")
	registerFunc(&_secCertificateCopyEmailAddresses, &_secCertificateCopyEmailAddressesErr, frameworkHandle, "SecCertificateCopyEmailAddresses", "10.5")
	registerFunc(&_secCertificateCopyKey, &_secCertificateCopyKeyErr, frameworkHandle, "SecCertificateCopyKey", "10.14")
	registerFunc(&_secCertificateCopyLongDescription, &_secCertificateCopyLongDescriptionErr, frameworkHandle, "SecCertificateCopyLongDescription", "10.7")
	registerFunc(&_secCertificateCopyNormalizedIssuerSequence, &_secCertificateCopyNormalizedIssuerSequenceErr, frameworkHandle, "SecCertificateCopyNormalizedIssuerSequence", "10.12.4")
	registerFunc(&_secCertificateCopyNormalizedSubjectSequence, &_secCertificateCopyNormalizedSubjectSequenceErr, frameworkHandle, "SecCertificateCopyNormalizedSubjectSequence", "10.12.4")
	registerFunc(&_secCertificateCopyNotValidAfterDate, &_secCertificateCopyNotValidAfterDateErr, frameworkHandle, "SecCertificateCopyNotValidAfterDate", "15.0")
	registerFunc(&_secCertificateCopyNotValidBeforeDate, &_secCertificateCopyNotValidBeforeDateErr, frameworkHandle, "SecCertificateCopyNotValidBeforeDate", "15.0")
	registerFunc(&_secCertificateCopyPreference, &_secCertificateCopyPreferenceErr, frameworkHandle, "SecCertificateCopyPreference", "10.0")
	registerFunc(&_secCertificateCopyPreferred, &_secCertificateCopyPreferredErr, frameworkHandle, "SecCertificateCopyPreferred", "10.7")
	registerFunc(&_secCertificateCopySerialNumberData, &_secCertificateCopySerialNumberDataErr, frameworkHandle, "SecCertificateCopySerialNumberData", "10.13")
	registerFunc(&_secCertificateCopyShortDescription, &_secCertificateCopyShortDescriptionErr, frameworkHandle, "SecCertificateCopyShortDescription", "10.7")
	registerFunc(&_secCertificateCopySubjectSummary, &_secCertificateCopySubjectSummaryErr, frameworkHandle, "SecCertificateCopySubjectSummary", "10.6")
	registerFunc(&_secCertificateCopyValues, &_secCertificateCopyValuesErr, frameworkHandle, "SecCertificateCopyValues", "10.7")
	registerFunc(&_secCertificateCreateWithData, &_secCertificateCreateWithDataErr, frameworkHandle, "SecCertificateCreateWithData", "10.6")
	registerFunc(&_secCertificateGetAlgorithmID, &_secCertificateGetAlgorithmIDErr, frameworkHandle, "SecCertificateGetAlgorithmID", "10.0")
	registerFunc(&_secCertificateGetCLHandle, &_secCertificateGetCLHandleErr, frameworkHandle, "SecCertificateGetCLHandle", "10.0")
	registerFunc(&_secCertificateGetData, &_secCertificateGetDataErr, frameworkHandle, "SecCertificateGetData", "10.0")
	registerFunc(&_secCertificateGetIssuer, &_secCertificateGetIssuerErr, frameworkHandle, "SecCertificateGetIssuer", "10.0")
	registerFunc(&_secCertificateGetSubject, &_secCertificateGetSubjectErr, frameworkHandle, "SecCertificateGetSubject", "10.0")
	registerFunc(&_secCertificateGetType, &_secCertificateGetTypeErr, frameworkHandle, "SecCertificateGetType", "10.0")
	registerFunc(&_secCertificateGetTypeID, &_secCertificateGetTypeIDErr, frameworkHandle, "SecCertificateGetTypeID", "10.3")
	registerFunc(&_secCertificateSetPreference, &_secCertificateSetPreferenceErr, frameworkHandle, "SecCertificateSetPreference", "10.0")
	registerFunc(&_secCertificateSetPreferred, &_secCertificateSetPreferredErr, frameworkHandle, "SecCertificateSetPreferred", "10.7")
	registerFunc(&_secCodeCheckValidity, &_secCodeCheckValidityErr, frameworkHandle, "SecCodeCheckValidity", "10.0")
	registerFunc(&_secCodeCheckValidityWithErrors, &_secCodeCheckValidityWithErrorsErr, frameworkHandle, "SecCodeCheckValidityWithErrors", "10.0")
	registerFunc(&_secCodeCopyDesignatedRequirement, &_secCodeCopyDesignatedRequirementErr, frameworkHandle, "SecCodeCopyDesignatedRequirement", "10.0")
	registerFunc(&_secCodeCopyGuestWithAttributes, &_secCodeCopyGuestWithAttributesErr, frameworkHandle, "SecCodeCopyGuestWithAttributes", "10.0")
	registerFunc(&_secCodeCopyHost, &_secCodeCopyHostErr, frameworkHandle, "SecCodeCopyHost", "10.0")
	registerFunc(&_secCodeCopyPath, &_secCodeCopyPathErr, frameworkHandle, "SecCodeCopyPath", "10.0")
	registerFunc(&_secCodeCopySelf, &_secCodeCopySelfErr, frameworkHandle, "SecCodeCopySelf", "10.0")
	registerFunc(&_secCodeCopySigningInformation, &_secCodeCopySigningInformationErr, frameworkHandle, "SecCodeCopySigningInformation", "10.0")
	registerFunc(&_secCodeCopyStaticCode, &_secCodeCopyStaticCodeErr, frameworkHandle, "SecCodeCopyStaticCode", "10.0")
	registerFunc(&_secCodeCreateWithXPCMessage, &_secCodeCreateWithXPCMessageErr, frameworkHandle, "SecCodeCreateWithXPCMessage", "10.0")
	registerFunc(&_secCodeGetTypeID, &_secCodeGetTypeIDErr, frameworkHandle, "SecCodeGetTypeID", "10.0")
	registerFunc(&_secCodeMapMemory, &_secCodeMapMemoryErr, frameworkHandle, "SecCodeMapMemory", "10.0")
	registerFunc(&_secCodeValidateFileResource, &_secCodeValidateFileResourceErr, frameworkHandle, "SecCodeValidateFileResource", "10.13")
	registerFunc(&_secCopyErrorMessageString, &_secCopyErrorMessageStringErr, frameworkHandle, "SecCopyErrorMessageString", "10.3")
	registerFunc(&_secCreateSharedWebCredentialPassword, &_secCreateSharedWebCredentialPasswordErr, frameworkHandle, "SecCreateSharedWebCredentialPassword", "11.0")
	registerFunc(&_secHostCreateGuest, &_secHostCreateGuestErr, frameworkHandle, "SecHostCreateGuest", "10.5")
	registerFunc(&_secHostRemoveGuest, &_secHostRemoveGuestErr, frameworkHandle, "SecHostRemoveGuest", "10.5")
	registerFunc(&_secHostSelectGuest, &_secHostSelectGuestErr, frameworkHandle, "SecHostSelectGuest", "10.5")
	registerFunc(&_secHostSelectedGuest, &_secHostSelectedGuestErr, frameworkHandle, "SecHostSelectedGuest", "10.5")
	registerFunc(&_secHostSetGuestStatus, &_secHostSetGuestStatusErr, frameworkHandle, "SecHostSetGuestStatus", "10.5")
	registerFunc(&_secHostSetHostingPort, &_secHostSetHostingPortErr, frameworkHandle, "SecHostSetHostingPort", "10.5")
	registerFunc(&_secIdentityCopyCertificate, &_secIdentityCopyCertificateErr, frameworkHandle, "SecIdentityCopyCertificate", "10.3")
	registerFunc(&_secIdentityCopyPreference, &_secIdentityCopyPreferenceErr, frameworkHandle, "SecIdentityCopyPreference", "10.0")
	registerFunc(&_secIdentityCopyPreferred, &_secIdentityCopyPreferredErr, frameworkHandle, "SecIdentityCopyPreferred", "10.7")
	registerFunc(&_secIdentityCopyPrivateKey, &_secIdentityCopyPrivateKeyErr, frameworkHandle, "SecIdentityCopyPrivateKey", "10.3")
	registerFunc(&_secIdentityCopySystemIdentity, &_secIdentityCopySystemIdentityErr, frameworkHandle, "SecIdentityCopySystemIdentity", "10.5")
	registerFunc(&_secIdentityCreate, &_secIdentityCreateErr, frameworkHandle, "SecIdentityCreate", "10.12")
	registerFunc(&_secIdentityCreateWithCertificate, &_secIdentityCreateWithCertificateErr, frameworkHandle, "SecIdentityCreateWithCertificate", "10.5")
	registerFunc(&_secIdentityGetTypeID, &_secIdentityGetTypeIDErr, frameworkHandle, "SecIdentityGetTypeID", "10.3")
	registerFunc(&_secIdentitySearchCopyNext, &_secIdentitySearchCopyNextErr, frameworkHandle, "SecIdentitySearchCopyNext", "10.0")
	registerFunc(&_secIdentitySearchCreate, &_secIdentitySearchCreateErr, frameworkHandle, "SecIdentitySearchCreate", "10.0")
	registerFunc(&_secIdentitySearchGetTypeID, &_secIdentitySearchGetTypeIDErr, frameworkHandle, "SecIdentitySearchGetTypeID", "10.0")
	registerFunc(&_secIdentitySetPreference, &_secIdentitySetPreferenceErr, frameworkHandle, "SecIdentitySetPreference", "10.0")
	registerFunc(&_secIdentitySetPreferred, &_secIdentitySetPreferredErr, frameworkHandle, "SecIdentitySetPreferred", "10.7")
	registerFunc(&_secIdentitySetSystemIdentity, &_secIdentitySetSystemIdentityErr, frameworkHandle, "SecIdentitySetSystemIdentity", "10.5")
	registerFunc(&_secItemAdd, &_secItemAddErr, frameworkHandle, "SecItemAdd", "10.6")
	registerFunc(&_secItemCopyMatching, &_secItemCopyMatchingErr, frameworkHandle, "SecItemCopyMatching", "10.6")
	registerFunc(&_secItemDelete, &_secItemDeleteErr, frameworkHandle, "SecItemDelete", "10.6")
	registerFunc(&_secItemExport, &_secItemExportErr, frameworkHandle, "SecItemExport", "10.7")
	registerFunc(&_secItemImport, &_secItemImportErr, frameworkHandle, "SecItemImport", "10.7")
	registerFunc(&_secItemUpdate, &_secItemUpdateErr, frameworkHandle, "SecItemUpdate", "10.6")
	registerFunc(&_secKeyCopyAttributes, &_secKeyCopyAttributesErr, frameworkHandle, "SecKeyCopyAttributes", "10.12")
	registerFunc(&_secKeyCopyExternalRepresentation, &_secKeyCopyExternalRepresentationErr, frameworkHandle, "SecKeyCopyExternalRepresentation", "10.12")
	registerFunc(&_secKeyCopyKeyExchangeResult, &_secKeyCopyKeyExchangeResultErr, frameworkHandle, "SecKeyCopyKeyExchangeResult", "10.12")
	registerFunc(&_secKeyCopyPublicKey, &_secKeyCopyPublicKeyErr, frameworkHandle, "SecKeyCopyPublicKey", "10.12")
	registerFunc(&_secKeyCreateDecryptedData, &_secKeyCreateDecryptedDataErr, frameworkHandle, "SecKeyCreateDecryptedData", "10.12")
	registerFunc(&_secKeyCreateEncryptedData, &_secKeyCreateEncryptedDataErr, frameworkHandle, "SecKeyCreateEncryptedData", "10.12")
	registerFunc(&_secKeyCreatePair, &_secKeyCreatePairErr, frameworkHandle, "SecKeyCreatePair", "10.0")
	registerFunc(&_secKeyCreateRandomKey, &_secKeyCreateRandomKeyErr, frameworkHandle, "SecKeyCreateRandomKey", "10.12")
	registerFunc(&_secKeyCreateSignature, &_secKeyCreateSignatureErr, frameworkHandle, "SecKeyCreateSignature", "10.12")
	registerFunc(&_secKeyCreateWithData, &_secKeyCreateWithDataErr, frameworkHandle, "SecKeyCreateWithData", "10.12")
	registerFunc(&_secKeyGenerate, &_secKeyGenerateErr, frameworkHandle, "SecKeyGenerate", "10.0")
	registerFunc(&_secKeyGetBlockSize, &_secKeyGetBlockSizeErr, frameworkHandle, "SecKeyGetBlockSize", "10.6")
	registerFunc(&_secKeyGetCSPHandle, &_secKeyGetCSPHandleErr, frameworkHandle, "SecKeyGetCSPHandle", "10.0")
	registerFunc(&_secKeyGetCSSMKey, &_secKeyGetCSSMKeyErr, frameworkHandle, "SecKeyGetCSSMKey", "10.0")
	registerFunc(&_secKeyGetCredentials, &_secKeyGetCredentialsErr, frameworkHandle, "SecKeyGetCredentials", "10.0")
	registerFunc(&_secKeyGetTypeID, &_secKeyGetTypeIDErr, frameworkHandle, "SecKeyGetTypeID", "10.3")
	registerFunc(&_secKeyIsAlgorithmSupported, &_secKeyIsAlgorithmSupportedErr, frameworkHandle, "SecKeyIsAlgorithmSupported", "10.12")
	registerFunc(&_secKeyVerifySignature, &_secKeyVerifySignatureErr, frameworkHandle, "SecKeyVerifySignature", "10.12")
	registerFunc(&_secKeychainSearchGetTypeID, &_secKeychainSearchGetTypeIDErr, frameworkHandle, "SecKeychainSearchGetTypeID", "10.0")
	registerFunc(&_secPKCS12Import, &_secPKCS12ImportErr, frameworkHandle, "SecPKCS12Import", "10.6")
	registerFunc(&_secPolicyCopyProperties, &_secPolicyCopyPropertiesErr, frameworkHandle, "SecPolicyCopyProperties", "10.7")
	registerFunc(&_secPolicyCreateBasicX509, &_secPolicyCreateBasicX509Err, frameworkHandle, "SecPolicyCreateBasicX509", "10.6")
	registerFunc(&_secPolicyCreateRevocation, &_secPolicyCreateRevocationErr, frameworkHandle, "SecPolicyCreateRevocation", "10.9")
	registerFunc(&_secPolicyCreateSSL, &_secPolicyCreateSSLErr, frameworkHandle, "SecPolicyCreateSSL", "10.6")
	registerFunc(&_secPolicyCreateWithOID, &_secPolicyCreateWithOIDErr, frameworkHandle, "SecPolicyCreateWithOID", "10.7")
	registerFunc(&_secPolicyCreateWithProperties, &_secPolicyCreateWithPropertiesErr, frameworkHandle, "SecPolicyCreateWithProperties", "10.9")
	registerFunc(&_secPolicyGetOID, &_secPolicyGetOIDErr, frameworkHandle, "SecPolicyGetOID", "10.2")
	registerFunc(&_secPolicyGetTPHandle, &_secPolicyGetTPHandleErr, frameworkHandle, "SecPolicyGetTPHandle", "10.2")
	registerFunc(&_secPolicyGetTypeID, &_secPolicyGetTypeIDErr, frameworkHandle, "SecPolicyGetTypeID", "10.3")
	registerFunc(&_secPolicyGetValue, &_secPolicyGetValueErr, frameworkHandle, "SecPolicyGetValue", "10.2")
	registerFunc(&_secPolicySearchCopyNext, &_secPolicySearchCopyNextErr, frameworkHandle, "SecPolicySearchCopyNext", "10.0")
	registerFunc(&_secPolicySearchCreate, &_secPolicySearchCreateErr, frameworkHandle, "SecPolicySearchCreate", "10.0")
	registerFunc(&_secPolicySearchGetTypeID, &_secPolicySearchGetTypeIDErr, frameworkHandle, "SecPolicySearchGetTypeID", "10.0")
	registerFunc(&_secPolicySetProperties, &_secPolicySetPropertiesErr, frameworkHandle, "SecPolicySetProperties", "10.7")
	registerFunc(&_secPolicySetValue, &_secPolicySetValueErr, frameworkHandle, "SecPolicySetValue", "10.2")
	registerFunc(&_secRandomCopyBytes, &_secRandomCopyBytesErr, frameworkHandle, "SecRandomCopyBytes", "10.7")
	registerFunc(&_secRequirementCopyData, &_secRequirementCopyDataErr, frameworkHandle, "SecRequirementCopyData", "10.0")
	registerFunc(&_secRequirementCopyString, &_secRequirementCopyStringErr, frameworkHandle, "SecRequirementCopyString", "10.0")
	registerFunc(&_secRequirementCreateWithData, &_secRequirementCreateWithDataErr, frameworkHandle, "SecRequirementCreateWithData", "10.0")
	registerFunc(&_secRequirementCreateWithString, &_secRequirementCreateWithStringErr, frameworkHandle, "SecRequirementCreateWithString", "10.0")
	registerFunc(&_secRequirementCreateWithStringAndErrors, &_secRequirementCreateWithStringAndErrorsErr, frameworkHandle, "SecRequirementCreateWithStringAndErrors", "10.0")
	registerFunc(&_secRequirementGetTypeID, &_secRequirementGetTypeIDErr, frameworkHandle, "SecRequirementGetTypeID", "10.0")
	registerFunc(&_secStaticCodeCheckValidity, &_secStaticCodeCheckValidityErr, frameworkHandle, "SecStaticCodeCheckValidity", "10.0")
	registerFunc(&_secStaticCodeCheckValidityWithErrors, &_secStaticCodeCheckValidityWithErrorsErr, frameworkHandle, "SecStaticCodeCheckValidityWithErrors", "10.0")
	registerFunc(&_secStaticCodeCreateWithPath, &_secStaticCodeCreateWithPathErr, frameworkHandle, "SecStaticCodeCreateWithPath", "10.0")
	registerFunc(&_secStaticCodeCreateWithPathAndAttributes, &_secStaticCodeCreateWithPathAndAttributesErr, frameworkHandle, "SecStaticCodeCreateWithPathAndAttributes", "10.0")
	registerFunc(&_secStaticCodeGetTypeID, &_secStaticCodeGetTypeIDErr, frameworkHandle, "SecStaticCodeGetTypeID", "10.0")
	registerFunc(&_secTaskCopySigningIdentifier, &_secTaskCopySigningIdentifierErr, frameworkHandle, "SecTaskCopySigningIdentifier", "10.0")
	registerFunc(&_secTaskCopyValueForEntitlement, &_secTaskCopyValueForEntitlementErr, frameworkHandle, "SecTaskCopyValueForEntitlement", "10.0")
	registerFunc(&_secTaskCopyValuesForEntitlements, &_secTaskCopyValuesForEntitlementsErr, frameworkHandle, "SecTaskCopyValuesForEntitlements", "10.0")
	registerFunc(&_secTaskCreateFromSelf, &_secTaskCreateFromSelfErr, frameworkHandle, "SecTaskCreateFromSelf", "10.0")
	registerFunc(&_secTaskGetTypeID, &_secTaskGetTypeIDErr, frameworkHandle, "SecTaskGetTypeID", "10.0")
	registerFunc(&_secTrustCopyAnchorCertificates, &_secTrustCopyAnchorCertificatesErr, frameworkHandle, "SecTrustCopyAnchorCertificates", "10.3")
	registerFunc(&_secTrustCopyCertificateChain, &_secTrustCopyCertificateChainErr, frameworkHandle, "SecTrustCopyCertificateChain", "12.0")
	registerFunc(&_secTrustCopyCustomAnchorCertificates, &_secTrustCopyCustomAnchorCertificatesErr, frameworkHandle, "SecTrustCopyCustomAnchorCertificates", "10.5")
	registerFunc(&_secTrustCopyExceptions, &_secTrustCopyExceptionsErr, frameworkHandle, "SecTrustCopyExceptions", "10.9")
	registerFunc(&_secTrustCopyKey, &_secTrustCopyKeyErr, frameworkHandle, "SecTrustCopyKey", "11.0")
	registerFunc(&_secTrustCopyPolicies, &_secTrustCopyPoliciesErr, frameworkHandle, "SecTrustCopyPolicies", "10.3")
	registerFunc(&_secTrustCopyProperties, &_secTrustCopyPropertiesErr, frameworkHandle, "SecTrustCopyProperties", "10.7")
	registerFunc(&_secTrustCopyPublicKey, &_secTrustCopyPublicKeyErr, frameworkHandle, "SecTrustCopyPublicKey", "10.7")
	registerFunc(&_secTrustCopyResult, &_secTrustCopyResultErr, frameworkHandle, "SecTrustCopyResult", "10.9")
	registerFunc(&_secTrustCreateWithCertificates, &_secTrustCreateWithCertificatesErr, frameworkHandle, "SecTrustCreateWithCertificates", "10.3")
	registerFunc(&_secTrustEvaluateAsyncWithError, &_secTrustEvaluateAsyncWithErrorErr, frameworkHandle, "SecTrustEvaluateAsyncWithError", "10.15")
	registerFunc(&_secTrustEvaluateWithError, &_secTrustEvaluateWithErrorErr, frameworkHandle, "SecTrustEvaluateWithError", "10.14")
	registerFunc(&_secTrustGetCertificateAtIndex, &_secTrustGetCertificateAtIndexErr, frameworkHandle, "SecTrustGetCertificateAtIndex", "10.7")
	registerFunc(&_secTrustGetCertificateCount, &_secTrustGetCertificateCountErr, frameworkHandle, "SecTrustGetCertificateCount", "10.7")
	registerFunc(&_secTrustGetCssmResult, &_secTrustGetCssmResultErr, frameworkHandle, "SecTrustGetCssmResult", "10.2")
	registerFunc(&_secTrustGetCssmResultCode, &_secTrustGetCssmResultCodeErr, frameworkHandle, "SecTrustGetCssmResultCode", "10.2")
	registerFunc(&_secTrustGetNetworkFetchAllowed, &_secTrustGetNetworkFetchAllowedErr, frameworkHandle, "SecTrustGetNetworkFetchAllowed", "10.9")
	registerFunc(&_secTrustGetResult, &_secTrustGetResultErr, frameworkHandle, "SecTrustGetResult", "10.2")
	registerFunc(&_secTrustGetTPHandle, &_secTrustGetTPHandleErr, frameworkHandle, "SecTrustGetTPHandle", "10.2")
	registerFunc(&_secTrustGetTrustResult, &_secTrustGetTrustResultErr, frameworkHandle, "SecTrustGetTrustResult", "10.7")
	registerFunc(&_secTrustGetTypeID, &_secTrustGetTypeIDErr, frameworkHandle, "SecTrustGetTypeID", "10.3")
	registerFunc(&_secTrustGetVerifyTime, &_secTrustGetVerifyTimeErr, frameworkHandle, "SecTrustGetVerifyTime", "10.6")
	registerFunc(&_secTrustSetAnchorCertificates, &_secTrustSetAnchorCertificatesErr, frameworkHandle, "SecTrustSetAnchorCertificates", "10.3")
	registerFunc(&_secTrustSetAnchorCertificatesOnly, &_secTrustSetAnchorCertificatesOnlyErr, frameworkHandle, "SecTrustSetAnchorCertificatesOnly", "10.6")
	registerFunc(&_secTrustSetExceptions, &_secTrustSetExceptionsErr, frameworkHandle, "SecTrustSetExceptions", "10.9")
	registerFunc(&_secTrustSetKeychains, &_secTrustSetKeychainsErr, frameworkHandle, "SecTrustSetKeychains", "10.3")
	registerFunc(&_secTrustSetNetworkFetchAllowed, &_secTrustSetNetworkFetchAllowedErr, frameworkHandle, "SecTrustSetNetworkFetchAllowed", "10.9")
	registerFunc(&_secTrustSetOCSPResponse, &_secTrustSetOCSPResponseErr, frameworkHandle, "SecTrustSetOCSPResponse", "10.9")
	registerFunc(&_secTrustSetOptions, &_secTrustSetOptionsErr, frameworkHandle, "SecTrustSetOptions", "10.7")
	registerFunc(&_secTrustSetParameters, &_secTrustSetParametersErr, frameworkHandle, "SecTrustSetParameters", "10.2")
	registerFunc(&_secTrustSetPolicies, &_secTrustSetPoliciesErr, frameworkHandle, "SecTrustSetPolicies", "10.3")
	registerFunc(&_secTrustSetSignedCertificateTimestamps, &_secTrustSetSignedCertificateTimestampsErr, frameworkHandle, "SecTrustSetSignedCertificateTimestamps", "10.14.2")
	registerFunc(&_secTrustSetVerifyDate, &_secTrustSetVerifyDateErr, frameworkHandle, "SecTrustSetVerifyDate", "10.3")
	registerFunc(&_secTrustSettingsCopyCertificates, &_secTrustSettingsCopyCertificatesErr, frameworkHandle, "SecTrustSettingsCopyCertificates", "10.0")
	registerFunc(&_secTrustSettingsCopyModificationDate, &_secTrustSettingsCopyModificationDateErr, frameworkHandle, "SecTrustSettingsCopyModificationDate", "10.0")
	registerFunc(&_secTrustSettingsCopyTrustSettings, &_secTrustSettingsCopyTrustSettingsErr, frameworkHandle, "SecTrustSettingsCopyTrustSettings", "10.0")
	registerFunc(&_secTrustSettingsCreateExternalRepresentation, &_secTrustSettingsCreateExternalRepresentationErr, frameworkHandle, "SecTrustSettingsCreateExternalRepresentation", "10.0")
	registerFunc(&_secTrustSettingsImportExternalRepresentation, &_secTrustSettingsImportExternalRepresentationErr, frameworkHandle, "SecTrustSettingsImportExternalRepresentation", "10.0")
	registerFunc(&_secTrustSettingsRemoveTrustSettings, &_secTrustSettingsRemoveTrustSettingsErr, frameworkHandle, "SecTrustSettingsRemoveTrustSettings", "10.0")
	registerFunc(&_secTrustSettingsSetTrustSettings, &_secTrustSettingsSetTrustSettingsErr, frameworkHandle, "SecTrustSettingsSetTrustSettings", "10.0")
	registerFunc(&_secureDownloadCopyCreationDate, &_secureDownloadCopyCreationDateErr, frameworkHandle, "SecureDownloadCopyCreationDate", "10.5")
	registerFunc(&_secureDownloadCopyName, &_secureDownloadCopyNameErr, frameworkHandle, "SecureDownloadCopyName", "10.5")
	registerFunc(&_secureDownloadCopyTicketLocation, &_secureDownloadCopyTicketLocationErr, frameworkHandle, "SecureDownloadCopyTicketLocation", "10.5")
	registerFunc(&_secureDownloadCopyURLs, &_secureDownloadCopyURLsErr, frameworkHandle, "SecureDownloadCopyURLs", "10.5")
	registerFunc(&_secureDownloadCreateWithTicket, &_secureDownloadCreateWithTicketErr, frameworkHandle, "SecureDownloadCreateWithTicket", "10.5")
	registerFunc(&_secureDownloadFinished, &_secureDownloadFinishedErr, frameworkHandle, "SecureDownloadFinished", "10.5")
	registerFunc(&_secureDownloadGetDownloadSize, &_secureDownloadGetDownloadSizeErr, frameworkHandle, "SecureDownloadGetDownloadSize", "10.5")
	registerFunc(&_secureDownloadRelease, &_secureDownloadReleaseErr, frameworkHandle, "SecureDownloadRelease", "10.5")
	registerFunc(&_secureDownloadUpdateWithData, &_secureDownloadUpdateWithDataErr, frameworkHandle, "SecureDownloadUpdateWithData", "10.5")
	registerFunc(&_sessionCreate, &_sessionCreateErr, frameworkHandle, "SessionCreate", "10.0")
	registerFunc(&_sessionGetInfo, &_sessionGetInfoErr, frameworkHandle, "SessionGetInfo", "10.0")
	registerFunc(&_cssmAlgToOid, &_cssmAlgToOidErr, frameworkHandle, "cssmAlgToOid", "10.0")
	registerFunc(&_cssmOidToAlg, &_cssmOidToAlgErr, frameworkHandle, "cssmOidToAlg", "10.0")
	registerFunc(&_cssmPerror, &_cssmPerrorErr, frameworkHandle, "cssmPerror", "10.0")
	registerFunc(&_sec_certificate_copy_ref, &_sec_certificate_copy_refErr, frameworkHandle, "sec_certificate_copy_ref", "10.14")
	registerFunc(&_sec_certificate_create, &_sec_certificate_createErr, frameworkHandle, "sec_certificate_create", "10.14")
	registerFunc(&_sec_identity_access_certificates, &_sec_identity_access_certificatesErr, frameworkHandle, "sec_identity_access_certificates", "10.15")
	registerFunc(&_sec_identity_copy_certificates_ref, &_sec_identity_copy_certificates_refErr, frameworkHandle, "sec_identity_copy_certificates_ref", "10.14")
	registerFunc(&_sec_identity_copy_ref, &_sec_identity_copy_refErr, frameworkHandle, "sec_identity_copy_ref", "10.14")
	registerFunc(&_sec_identity_create, &_sec_identity_createErr, frameworkHandle, "sec_identity_create", "10.14")
	registerFunc(&_sec_identity_create_with_certificates, &_sec_identity_create_with_certificatesErr, frameworkHandle, "sec_identity_create_with_certificates", "10.14")
	registerFunc(&_sec_protocol_metadata_access_distinguished_names, &_sec_protocol_metadata_access_distinguished_namesErr, frameworkHandle, "sec_protocol_metadata_access_distinguished_names", "10.14")
	registerFunc(&_sec_protocol_metadata_access_ocsp_response, &_sec_protocol_metadata_access_ocsp_responseErr, frameworkHandle, "sec_protocol_metadata_access_ocsp_response", "10.14")
	registerFunc(&_sec_protocol_metadata_access_peer_certificate_chain, &_sec_protocol_metadata_access_peer_certificate_chainErr, frameworkHandle, "sec_protocol_metadata_access_peer_certificate_chain", "10.14")
	registerFunc(&_sec_protocol_metadata_access_pre_shared_keys, &_sec_protocol_metadata_access_pre_shared_keysErr, frameworkHandle, "sec_protocol_metadata_access_pre_shared_keys", "10.15")
	registerFunc(&_sec_protocol_metadata_access_supported_signature_algorithms, &_sec_protocol_metadata_access_supported_signature_algorithmsErr, frameworkHandle, "sec_protocol_metadata_access_supported_signature_algorithms", "10.14")
	registerFunc(&_sec_protocol_metadata_challenge_parameters_are_equal, &_sec_protocol_metadata_challenge_parameters_are_equalErr, frameworkHandle, "sec_protocol_metadata_challenge_parameters_are_equal", "10.14")
	registerFunc(&_sec_protocol_metadata_copy_negotiated_protocol, &_sec_protocol_metadata_copy_negotiated_protocolErr, frameworkHandle, "sec_protocol_metadata_copy_negotiated_protocol", "15.5")
	registerFunc(&_sec_protocol_metadata_copy_peer_public_key, &_sec_protocol_metadata_copy_peer_public_keyErr, frameworkHandle, "sec_protocol_metadata_copy_peer_public_key", "10.14")
	registerFunc(&_sec_protocol_metadata_copy_server_name, &_sec_protocol_metadata_copy_server_nameErr, frameworkHandle, "sec_protocol_metadata_copy_server_name", "15.5")
	registerFunc(&_sec_protocol_metadata_create_secret, &_sec_protocol_metadata_create_secretErr, frameworkHandle, "sec_protocol_metadata_create_secret", "10.14")
	registerFunc(&_sec_protocol_metadata_create_secret_with_context, &_sec_protocol_metadata_create_secret_with_contextErr, frameworkHandle, "sec_protocol_metadata_create_secret_with_context", "10.14")
	registerFunc(&_sec_protocol_metadata_get_early_data_accepted, &_sec_protocol_metadata_get_early_data_acceptedErr, frameworkHandle, "sec_protocol_metadata_get_early_data_accepted", "10.14")
	registerFunc(&_sec_protocol_metadata_get_negotiated_ciphersuite, &_sec_protocol_metadata_get_negotiated_ciphersuiteErr, frameworkHandle, "sec_protocol_metadata_get_negotiated_ciphersuite", "10.14")
	registerFunc(&_sec_protocol_metadata_get_negotiated_protocol, &_sec_protocol_metadata_get_negotiated_protocolErr, frameworkHandle, "sec_protocol_metadata_get_negotiated_protocol", "10.14")
	registerFunc(&_sec_protocol_metadata_get_negotiated_protocol_version, &_sec_protocol_metadata_get_negotiated_protocol_versionErr, frameworkHandle, "sec_protocol_metadata_get_negotiated_protocol_version", "10.14")
	registerFunc(&_sec_protocol_metadata_get_negotiated_tls_ciphersuite, &_sec_protocol_metadata_get_negotiated_tls_ciphersuiteErr, frameworkHandle, "sec_protocol_metadata_get_negotiated_tls_ciphersuite", "10.15")
	registerFunc(&_sec_protocol_metadata_get_negotiated_tls_protocol_version, &_sec_protocol_metadata_get_negotiated_tls_protocol_versionErr, frameworkHandle, "sec_protocol_metadata_get_negotiated_tls_protocol_version", "10.15")
	registerFunc(&_sec_protocol_metadata_get_server_name, &_sec_protocol_metadata_get_server_nameErr, frameworkHandle, "sec_protocol_metadata_get_server_name", "10.14")
	registerFunc(&_sec_protocol_metadata_peers_are_equal, &_sec_protocol_metadata_peers_are_equalErr, frameworkHandle, "sec_protocol_metadata_peers_are_equal", "10.14")
	registerFunc(&_sec_protocol_options_add_pre_shared_key, &_sec_protocol_options_add_pre_shared_keyErr, frameworkHandle, "sec_protocol_options_add_pre_shared_key", "10.14")
	registerFunc(&_sec_protocol_options_add_tls_application_protocol, &_sec_protocol_options_add_tls_application_protocolErr, frameworkHandle, "sec_protocol_options_add_tls_application_protocol", "10.14")
	registerFunc(&_sec_protocol_options_append_tls_ciphersuite, &_sec_protocol_options_append_tls_ciphersuiteErr, frameworkHandle, "sec_protocol_options_append_tls_ciphersuite", "10.15")
	registerFunc(&_sec_protocol_options_append_tls_ciphersuite_group, &_sec_protocol_options_append_tls_ciphersuite_groupErr, frameworkHandle, "sec_protocol_options_append_tls_ciphersuite_group", "10.15")
	registerFunc(&_sec_protocol_options_are_equal, &_sec_protocol_options_are_equalErr, frameworkHandle, "sec_protocol_options_are_equal", "10.15")
	registerFunc(&_sec_protocol_options_get_default_max_dtls_protocol_version, &_sec_protocol_options_get_default_max_dtls_protocol_versionErr, frameworkHandle, "sec_protocol_options_get_default_max_dtls_protocol_version", "10.15")
	registerFunc(&_sec_protocol_options_get_default_max_tls_protocol_version, &_sec_protocol_options_get_default_max_tls_protocol_versionErr, frameworkHandle, "sec_protocol_options_get_default_max_tls_protocol_version", "10.15")
	registerFunc(&_sec_protocol_options_get_default_min_dtls_protocol_version, &_sec_protocol_options_get_default_min_dtls_protocol_versionErr, frameworkHandle, "sec_protocol_options_get_default_min_dtls_protocol_version", "10.15")
	registerFunc(&_sec_protocol_options_get_default_min_tls_protocol_version, &_sec_protocol_options_get_default_min_tls_protocol_versionErr, frameworkHandle, "sec_protocol_options_get_default_min_tls_protocol_version", "10.15")
	registerFunc(&_sec_protocol_options_get_enable_encrypted_client_hello, &_sec_protocol_options_get_enable_encrypted_client_helloErr, frameworkHandle, "sec_protocol_options_get_enable_encrypted_client_hello", "")
	registerFunc(&_sec_protocol_options_get_quic_use_legacy_codepoint, &_sec_protocol_options_get_quic_use_legacy_codepointErr, frameworkHandle, "sec_protocol_options_get_quic_use_legacy_codepoint", "")
	registerFunc(&_sec_protocol_options_set_challenge_block, &_sec_protocol_options_set_challenge_blockErr, frameworkHandle, "sec_protocol_options_set_challenge_block", "10.14")
	registerFunc(&_sec_protocol_options_set_enable_encrypted_client_hello, &_sec_protocol_options_set_enable_encrypted_client_helloErr, frameworkHandle, "sec_protocol_options_set_enable_encrypted_client_hello", "")
	registerFunc(&_sec_protocol_options_set_key_update_block, &_sec_protocol_options_set_key_update_blockErr, frameworkHandle, "sec_protocol_options_set_key_update_block", "10.14")
	registerFunc(&_sec_protocol_options_set_local_identity, &_sec_protocol_options_set_local_identityErr, frameworkHandle, "sec_protocol_options_set_local_identity", "10.14")
	registerFunc(&_sec_protocol_options_set_max_tls_protocol_version, &_sec_protocol_options_set_max_tls_protocol_versionErr, frameworkHandle, "sec_protocol_options_set_max_tls_protocol_version", "10.15")
	registerFunc(&_sec_protocol_options_set_min_tls_protocol_version, &_sec_protocol_options_set_min_tls_protocol_versionErr, frameworkHandle, "sec_protocol_options_set_min_tls_protocol_version", "10.15")
	registerFunc(&_sec_protocol_options_set_peer_authentication_optional, &_sec_protocol_options_set_peer_authentication_optionalErr, frameworkHandle, "sec_protocol_options_set_peer_authentication_optional", "")
	registerFunc(&_sec_protocol_options_set_peer_authentication_required, &_sec_protocol_options_set_peer_authentication_requiredErr, frameworkHandle, "sec_protocol_options_set_peer_authentication_required", "10.14")
	registerFunc(&_sec_protocol_options_set_pre_shared_key_selection_block, &_sec_protocol_options_set_pre_shared_key_selection_blockErr, frameworkHandle, "sec_protocol_options_set_pre_shared_key_selection_block", "10.15")
	registerFunc(&_sec_protocol_options_set_quic_use_legacy_codepoint, &_sec_protocol_options_set_quic_use_legacy_codepointErr, frameworkHandle, "sec_protocol_options_set_quic_use_legacy_codepoint", "")
	registerFunc(&_sec_protocol_options_set_tls_false_start_enabled, &_sec_protocol_options_set_tls_false_start_enabledErr, frameworkHandle, "sec_protocol_options_set_tls_false_start_enabled", "10.14")
	registerFunc(&_sec_protocol_options_set_tls_is_fallback_attempt, &_sec_protocol_options_set_tls_is_fallback_attemptErr, frameworkHandle, "sec_protocol_options_set_tls_is_fallback_attempt", "10.14")
	registerFunc(&_sec_protocol_options_set_tls_max_version, &_sec_protocol_options_set_tls_max_versionErr, frameworkHandle, "sec_protocol_options_set_tls_max_version", "10.14")
	registerFunc(&_sec_protocol_options_set_tls_min_version, &_sec_protocol_options_set_tls_min_versionErr, frameworkHandle, "sec_protocol_options_set_tls_min_version", "10.14")
	registerFunc(&_sec_protocol_options_set_tls_ocsp_enabled, &_sec_protocol_options_set_tls_ocsp_enabledErr, frameworkHandle, "sec_protocol_options_set_tls_ocsp_enabled", "10.14")
	registerFunc(&_sec_protocol_options_set_tls_pre_shared_key_identity_hint, &_sec_protocol_options_set_tls_pre_shared_key_identity_hintErr, frameworkHandle, "sec_protocol_options_set_tls_pre_shared_key_identity_hint", "10.15")
	registerFunc(&_sec_protocol_options_set_tls_renegotiation_enabled, &_sec_protocol_options_set_tls_renegotiation_enabledErr, frameworkHandle, "sec_protocol_options_set_tls_renegotiation_enabled", "10.14")
	registerFunc(&_sec_protocol_options_set_tls_resumption_enabled, &_sec_protocol_options_set_tls_resumption_enabledErr, frameworkHandle, "sec_protocol_options_set_tls_resumption_enabled", "10.14")
	registerFunc(&_sec_protocol_options_set_tls_sct_enabled, &_sec_protocol_options_set_tls_sct_enabledErr, frameworkHandle, "sec_protocol_options_set_tls_sct_enabled", "10.14")
	registerFunc(&_sec_protocol_options_set_tls_server_name, &_sec_protocol_options_set_tls_server_nameErr, frameworkHandle, "sec_protocol_options_set_tls_server_name", "10.14")
	registerFunc(&_sec_protocol_options_set_tls_tickets_enabled, &_sec_protocol_options_set_tls_tickets_enabledErr, frameworkHandle, "sec_protocol_options_set_tls_tickets_enabled", "10.14")
	registerFunc(&_sec_protocol_options_set_verify_block, &_sec_protocol_options_set_verify_blockErr, frameworkHandle, "sec_protocol_options_set_verify_block", "10.14")
	registerFunc(&_sec_release, &_sec_releaseErr, frameworkHandle, "sec_release", "10.0")
	registerFunc(&_sec_retain, &_sec_retainErr, frameworkHandle, "sec_retain", "10.0")
	registerFunc(&_sec_trust_copy_ref, &_sec_trust_copy_refErr, frameworkHandle, "sec_trust_copy_ref", "10.14")
	registerFunc(&_sec_trust_create, &_sec_trust_createErr, frameworkHandle, "sec_trust_create", "10.14")
}
