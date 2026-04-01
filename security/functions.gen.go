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

// CanCallAuthorizationCopyInfo reports whether the symbol for AuthorizationCopyInfo is available on this system.
func CanCallAuthorizationCopyInfo() bool {
	return _authorizationCopyInfo != nil
}

// AuthorizationCopyInfoCallError returns the symbol lookup or registration error for AuthorizationCopyInfo.
func AuthorizationCopyInfoCallError() error {
	if _authorizationCopyInfo != nil {
		return nil
	}
	return symbolCallError("AuthorizationCopyInfo", "10.0", _authorizationCopyInfoErr)
}

// TryAuthorizationCopyInfo calls AuthorizationCopyInfo without panicking when the symbol is unavailable.
func TryAuthorizationCopyInfo(authorization AuthorizationRef, tag AuthorizationString, info *AuthorizationItemSet) (int32, error) {
	if err := AuthorizationCopyInfoCallError(); err != nil {
		return 0, err
	}
	return _authorizationCopyInfo(authorization, tag, info), nil
}

// AuthorizationCopyInfo retrieves supporting data such as the user name and other information gathered during evaluation of authorization.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationCopyInfo(_:_:_:)
func AuthorizationCopyInfo(authorization AuthorizationRef, tag AuthorizationString, info *AuthorizationItemSet) int32 {
	result, callErr := TryAuthorizationCopyInfo(authorization, tag, info)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _authorizationCopyRights func(authorization AuthorizationRef, rights *AuthorizationRights, environment *AuthorizationEnvironment, flags AuthorizationFlags, authorizedRights *AuthorizationRights) int32
var _authorizationCopyRightsErr error

// CanCallAuthorizationCopyRights reports whether the symbol for AuthorizationCopyRights is available on this system.
func CanCallAuthorizationCopyRights() bool {
	return _authorizationCopyRights != nil
}

// AuthorizationCopyRightsCallError returns the symbol lookup or registration error for AuthorizationCopyRights.
func AuthorizationCopyRightsCallError() error {
	if _authorizationCopyRights != nil {
		return nil
	}
	return symbolCallError("AuthorizationCopyRights", "10.0", _authorizationCopyRightsErr)
}

// TryAuthorizationCopyRights calls AuthorizationCopyRights without panicking when the symbol is unavailable.
func TryAuthorizationCopyRights(authorization AuthorizationRef, rights *AuthorizationRights, environment *AuthorizationEnvironment, flags AuthorizationFlags, authorizedRights *AuthorizationRights) (int32, error) {
	if err := AuthorizationCopyRightsCallError(); err != nil {
		return 0, err
	}
	return _authorizationCopyRights(authorization, rights, environment, flags, authorizedRights), nil
}

// AuthorizationCopyRights authorizes and preauthorizes rights synchronously.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationCopyRights(_:_:_:_:_:)
func AuthorizationCopyRights(authorization AuthorizationRef, rights *AuthorizationRights, environment *AuthorizationEnvironment, flags AuthorizationFlags, authorizedRights *AuthorizationRights) int32 {
	result, callErr := TryAuthorizationCopyRights(authorization, rights, environment, flags, authorizedRights)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _authorizationCopyRightsAsync func(authorization AuthorizationRef, rights *AuthorizationRights, environment *AuthorizationEnvironment, flags AuthorizationFlags, callbackBlock unsafe.Pointer)
var _authorizationCopyRightsAsyncErr error

// CanCallAuthorizationCopyRightsAsync reports whether the symbol for AuthorizationCopyRightsAsync is available on this system.
func CanCallAuthorizationCopyRightsAsync() bool {
	return _authorizationCopyRightsAsync != nil
}

// AuthorizationCopyRightsAsyncCallError returns the symbol lookup or registration error for AuthorizationCopyRightsAsync.
func AuthorizationCopyRightsAsyncCallError() error {
	if _authorizationCopyRightsAsync != nil {
		return nil
	}
	return symbolCallError("AuthorizationCopyRightsAsync", "10.7", _authorizationCopyRightsAsyncErr)
}

// TryAuthorizationCopyRightsAsync calls AuthorizationCopyRightsAsync without panicking when the symbol is unavailable.
func TryAuthorizationCopyRightsAsync(authorization AuthorizationRef, rights *AuthorizationRights, environment *AuthorizationEnvironment, flags AuthorizationFlags, callbackBlock AuthorizationAsyncCallback) error {
	if err := AuthorizationCopyRightsAsyncCallError(); err != nil {
		return err
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 int, arg1 *AuthorizationItemSet) { callbackBlock(arg0, arg1) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_authorizationCopyRightsAsync(authorization, rights, environment, flags, _block0)
	return nil
}

// AuthorizationCopyRightsAsync authorizes and preauthorizes rights asynchronously.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationCopyRightsAsync(_:_:_:_:_:)
func AuthorizationCopyRightsAsync(authorization AuthorizationRef, rights *AuthorizationRights, environment *AuthorizationEnvironment, flags AuthorizationFlags, callbackBlock AuthorizationAsyncCallback) {
	if callErr := TryAuthorizationCopyRightsAsync(authorization, rights, environment, flags, callbackBlock); callErr != nil {
		panic(callErr)
	}
}

var _authorizationCreate func(rights *AuthorizationRights, environment *AuthorizationEnvironment, flags AuthorizationFlags, authorization *AuthorizationRef) int32
var _authorizationCreateErr error

// CanCallAuthorizationCreate reports whether the symbol for AuthorizationCreate is available on this system.
func CanCallAuthorizationCreate() bool {
	return _authorizationCreate != nil
}

// AuthorizationCreateCallError returns the symbol lookup or registration error for AuthorizationCreate.
func AuthorizationCreateCallError() error {
	if _authorizationCreate != nil {
		return nil
	}
	return symbolCallError("AuthorizationCreate", "10.0", _authorizationCreateErr)
}

// TryAuthorizationCreate calls AuthorizationCreate without panicking when the symbol is unavailable.
func TryAuthorizationCreate(rights *AuthorizationRights, environment *AuthorizationEnvironment, flags AuthorizationFlags, authorization *AuthorizationRef) (int32, error) {
	if err := AuthorizationCreateCallError(); err != nil {
		return 0, err
	}
	return _authorizationCreate(rights, environment, flags, authorization), nil
}

// AuthorizationCreate creates a new authorization reference and provides an option to authorize or preauthorize rights.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationCreate(_:_:_:_:)
func AuthorizationCreate(rights *AuthorizationRights, environment *AuthorizationEnvironment, flags AuthorizationFlags, authorization *AuthorizationRef) int32 {
	result, callErr := TryAuthorizationCreate(rights, environment, flags, authorization)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _authorizationCreateFromExternalForm func(extForm *AuthorizationExternalForm, authorization *AuthorizationRef) int32
var _authorizationCreateFromExternalFormErr error

// CanCallAuthorizationCreateFromExternalForm reports whether the symbol for AuthorizationCreateFromExternalForm is available on this system.
func CanCallAuthorizationCreateFromExternalForm() bool {
	return _authorizationCreateFromExternalForm != nil
}

// AuthorizationCreateFromExternalFormCallError returns the symbol lookup or registration error for AuthorizationCreateFromExternalForm.
func AuthorizationCreateFromExternalFormCallError() error {
	if _authorizationCreateFromExternalForm != nil {
		return nil
	}
	return symbolCallError("AuthorizationCreateFromExternalForm", "10.0", _authorizationCreateFromExternalFormErr)
}

// TryAuthorizationCreateFromExternalForm calls AuthorizationCreateFromExternalForm without panicking when the symbol is unavailable.
func TryAuthorizationCreateFromExternalForm(extForm *AuthorizationExternalForm, authorization *AuthorizationRef) (int32, error) {
	if err := AuthorizationCreateFromExternalFormCallError(); err != nil {
		return 0, err
	}
	return _authorizationCreateFromExternalForm(extForm, authorization), nil
}

// AuthorizationCreateFromExternalForm internalizes the external representation of an authorization reference.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationCreateFromExternalForm(_:_:)
func AuthorizationCreateFromExternalForm(extForm *AuthorizationExternalForm, authorization *AuthorizationRef) int32 {
	result, callErr := TryAuthorizationCreateFromExternalForm(extForm, authorization)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _authorizationFree func(authorization AuthorizationRef, flags AuthorizationFlags) int32
var _authorizationFreeErr error

// CanCallAuthorizationFree reports whether the symbol for AuthorizationFree is available on this system.
func CanCallAuthorizationFree() bool {
	return _authorizationFree != nil
}

// AuthorizationFreeCallError returns the symbol lookup or registration error for AuthorizationFree.
func AuthorizationFreeCallError() error {
	if _authorizationFree != nil {
		return nil
	}
	return symbolCallError("AuthorizationFree", "10.0", _authorizationFreeErr)
}

// TryAuthorizationFree calls AuthorizationFree without panicking when the symbol is unavailable.
func TryAuthorizationFree(authorization AuthorizationRef, flags AuthorizationFlags) (int32, error) {
	if err := AuthorizationFreeCallError(); err != nil {
		return 0, err
	}
	return _authorizationFree(authorization, flags), nil
}

// AuthorizationFree frees the memory associated with an authorization reference.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationFree(_:_:)
func AuthorizationFree(authorization AuthorizationRef, flags AuthorizationFlags) int32 {
	result, callErr := TryAuthorizationFree(authorization, flags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _authorizationFreeItemSet func(set *AuthorizationItemSet) int32
var _authorizationFreeItemSetErr error

// CanCallAuthorizationFreeItemSet reports whether the symbol for AuthorizationFreeItemSet is available on this system.
func CanCallAuthorizationFreeItemSet() bool {
	return _authorizationFreeItemSet != nil
}

// AuthorizationFreeItemSetCallError returns the symbol lookup or registration error for AuthorizationFreeItemSet.
func AuthorizationFreeItemSetCallError() error {
	if _authorizationFreeItemSet != nil {
		return nil
	}
	return symbolCallError("AuthorizationFreeItemSet", "10.0", _authorizationFreeItemSetErr)
}

// TryAuthorizationFreeItemSet calls AuthorizationFreeItemSet without panicking when the symbol is unavailable.
func TryAuthorizationFreeItemSet(set *AuthorizationItemSet) (int32, error) {
	if err := AuthorizationFreeItemSetCallError(); err != nil {
		return 0, err
	}
	return _authorizationFreeItemSet(set), nil
}

// AuthorizationFreeItemSet frees the memory associated with a set of authorization items.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationFreeItemSet(_:)
func AuthorizationFreeItemSet(set *AuthorizationItemSet) int32 {
	result, callErr := TryAuthorizationFreeItemSet(set)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _authorizationMakeExternalForm func(authorization AuthorizationRef, extForm *AuthorizationExternalForm) int32
var _authorizationMakeExternalFormErr error

// CanCallAuthorizationMakeExternalForm reports whether the symbol for AuthorizationMakeExternalForm is available on this system.
func CanCallAuthorizationMakeExternalForm() bool {
	return _authorizationMakeExternalForm != nil
}

// AuthorizationMakeExternalFormCallError returns the symbol lookup or registration error for AuthorizationMakeExternalForm.
func AuthorizationMakeExternalFormCallError() error {
	if _authorizationMakeExternalForm != nil {
		return nil
	}
	return symbolCallError("AuthorizationMakeExternalForm", "10.0", _authorizationMakeExternalFormErr)
}

// TryAuthorizationMakeExternalForm calls AuthorizationMakeExternalForm without panicking when the symbol is unavailable.
func TryAuthorizationMakeExternalForm(authorization AuthorizationRef, extForm *AuthorizationExternalForm) (int32, error) {
	if err := AuthorizationMakeExternalFormCallError(); err != nil {
		return 0, err
	}
	return _authorizationMakeExternalForm(authorization, extForm), nil
}

// AuthorizationMakeExternalForm creates an external representation of an authorization reference.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationMakeExternalForm(_:_:)
func AuthorizationMakeExternalForm(authorization AuthorizationRef, extForm *AuthorizationExternalForm) int32 {
	result, callErr := TryAuthorizationMakeExternalForm(authorization, extForm)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _authorizationRightGet func(rightName string, rightDefinition *corefoundation.CFDictionaryRef) int32
var _authorizationRightGetErr error

// CanCallAuthorizationRightGet reports whether the symbol for AuthorizationRightGet is available on this system.
func CanCallAuthorizationRightGet() bool {
	return _authorizationRightGet != nil
}

// AuthorizationRightGetCallError returns the symbol lookup or registration error for AuthorizationRightGet.
func AuthorizationRightGetCallError() error {
	if _authorizationRightGet != nil {
		return nil
	}
	return symbolCallError("AuthorizationRightGet", "10.0", _authorizationRightGetErr)
}

// TryAuthorizationRightGet calls AuthorizationRightGet without panicking when the symbol is unavailable.
func TryAuthorizationRightGet(rightName string, rightDefinition *corefoundation.CFDictionaryRef) (int32, error) {
	if err := AuthorizationRightGetCallError(); err != nil {
		return 0, err
	}
	return _authorizationRightGet(rightName, rightDefinition), nil
}

// AuthorizationRightGet retrieves a right definition as a dictionary.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationRightGet(_:_:)
func AuthorizationRightGet(rightName string, rightDefinition *corefoundation.CFDictionaryRef) int32 {
	result, callErr := TryAuthorizationRightGet(rightName, rightDefinition)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _authorizationRightRemove func(authRef AuthorizationRef, rightName string) int32
var _authorizationRightRemoveErr error

// CanCallAuthorizationRightRemove reports whether the symbol for AuthorizationRightRemove is available on this system.
func CanCallAuthorizationRightRemove() bool {
	return _authorizationRightRemove != nil
}

// AuthorizationRightRemoveCallError returns the symbol lookup or registration error for AuthorizationRightRemove.
func AuthorizationRightRemoveCallError() error {
	if _authorizationRightRemove != nil {
		return nil
	}
	return symbolCallError("AuthorizationRightRemove", "10.0", _authorizationRightRemoveErr)
}

// TryAuthorizationRightRemove calls AuthorizationRightRemove without panicking when the symbol is unavailable.
func TryAuthorizationRightRemove(authRef AuthorizationRef, rightName string) (int32, error) {
	if err := AuthorizationRightRemoveCallError(); err != nil {
		return 0, err
	}
	return _authorizationRightRemove(authRef, rightName), nil
}

// AuthorizationRightRemove removes a right from the policy database.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationRightRemove(_:_:)
func AuthorizationRightRemove(authRef AuthorizationRef, rightName string) int32 {
	result, callErr := TryAuthorizationRightRemove(authRef, rightName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _authorizationRightSet func(authRef AuthorizationRef, rightName string, rightDefinition corefoundation.CFTypeRef, descriptionKey corefoundation.CFStringRef, bundle corefoundation.CFBundle, localeTableName corefoundation.CFStringRef) int32
var _authorizationRightSetErr error

// CanCallAuthorizationRightSet reports whether the symbol for AuthorizationRightSet is available on this system.
func CanCallAuthorizationRightSet() bool {
	return _authorizationRightSet != nil
}

// AuthorizationRightSetCallError returns the symbol lookup or registration error for AuthorizationRightSet.
func AuthorizationRightSetCallError() error {
	if _authorizationRightSet != nil {
		return nil
	}
	return symbolCallError("AuthorizationRightSet", "10.0", _authorizationRightSetErr)
}

// TryAuthorizationRightSet calls AuthorizationRightSet without panicking when the symbol is unavailable.
func TryAuthorizationRightSet(authRef AuthorizationRef, rightName string, rightDefinition corefoundation.CFTypeRef, descriptionKey corefoundation.CFStringRef, bundle corefoundation.CFBundle, localeTableName corefoundation.CFStringRef) (int32, error) {
	if err := AuthorizationRightSetCallError(); err != nil {
		return 0, err
	}
	return _authorizationRightSet(authRef, rightName, rightDefinition, descriptionKey, bundle, localeTableName), nil
}

// AuthorizationRightSet creates or updates a right entry in the policy database.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationRightSet(_:_:_:_:_:_:)
func AuthorizationRightSet(authRef AuthorizationRef, rightName string, rightDefinition corefoundation.CFTypeRef, descriptionKey corefoundation.CFStringRef, bundle corefoundation.CFBundle, localeTableName corefoundation.CFStringRef) int32 {
	result, callErr := TryAuthorizationRightSet(authRef, rightName, rightDefinition, descriptionKey, bundle, localeTableName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSDecoderCopyAllCerts func(cmsDecoder CMSDecoderRef, certsOut *corefoundation.CFArrayRef) int32
var _cMSDecoderCopyAllCertsErr error

// CanCallCMSDecoderCopyAllCerts reports whether the symbol for CMSDecoderCopyAllCerts is available on this system.
func CanCallCMSDecoderCopyAllCerts() bool {
	return _cMSDecoderCopyAllCerts != nil
}

// CMSDecoderCopyAllCertsCallError returns the symbol lookup or registration error for CMSDecoderCopyAllCerts.
func CMSDecoderCopyAllCertsCallError() error {
	if _cMSDecoderCopyAllCerts != nil {
		return nil
	}
	return symbolCallError("CMSDecoderCopyAllCerts", "10.5", _cMSDecoderCopyAllCertsErr)
}

// TryCMSDecoderCopyAllCerts calls CMSDecoderCopyAllCerts without panicking when the symbol is unavailable.
func TryCMSDecoderCopyAllCerts(cmsDecoder CMSDecoderRef, certsOut *corefoundation.CFArrayRef) (int32, error) {
	if err := CMSDecoderCopyAllCertsCallError(); err != nil {
		return 0, err
	}
	return _cMSDecoderCopyAllCerts(cmsDecoder, certsOut), nil
}

// CMSDecoderCopyAllCerts obtains an array of all of the certificates in a message.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderCopyAllCerts(_:_:)
func CMSDecoderCopyAllCerts(cmsDecoder CMSDecoderRef, certsOut *corefoundation.CFArrayRef) int32 {
	result, callErr := TryCMSDecoderCopyAllCerts(cmsDecoder, certsOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSDecoderCopyContent func(cmsDecoder CMSDecoderRef, contentOut *corefoundation.CFDataRef) int32
var _cMSDecoderCopyContentErr error

// CanCallCMSDecoderCopyContent reports whether the symbol for CMSDecoderCopyContent is available on this system.
func CanCallCMSDecoderCopyContent() bool {
	return _cMSDecoderCopyContent != nil
}

// CMSDecoderCopyContentCallError returns the symbol lookup or registration error for CMSDecoderCopyContent.
func CMSDecoderCopyContentCallError() error {
	if _cMSDecoderCopyContent != nil {
		return nil
	}
	return symbolCallError("CMSDecoderCopyContent", "10.5", _cMSDecoderCopyContentErr)
}

// TryCMSDecoderCopyContent calls CMSDecoderCopyContent without panicking when the symbol is unavailable.
func TryCMSDecoderCopyContent(cmsDecoder CMSDecoderRef, contentOut *corefoundation.CFDataRef) (int32, error) {
	if err := CMSDecoderCopyContentCallError(); err != nil {
		return 0, err
	}
	return _cMSDecoderCopyContent(cmsDecoder, contentOut), nil
}

// CMSDecoderCopyContent obtains the message content, if any.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderCopyContent(_:_:)
func CMSDecoderCopyContent(cmsDecoder CMSDecoderRef, contentOut *corefoundation.CFDataRef) int32 {
	result, callErr := TryCMSDecoderCopyContent(cmsDecoder, contentOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSDecoderCopyDetachedContent func(cmsDecoder CMSDecoderRef, detachedContentOut *corefoundation.CFDataRef) int32
var _cMSDecoderCopyDetachedContentErr error

// CanCallCMSDecoderCopyDetachedContent reports whether the symbol for CMSDecoderCopyDetachedContent is available on this system.
func CanCallCMSDecoderCopyDetachedContent() bool {
	return _cMSDecoderCopyDetachedContent != nil
}

// CMSDecoderCopyDetachedContentCallError returns the symbol lookup or registration error for CMSDecoderCopyDetachedContent.
func CMSDecoderCopyDetachedContentCallError() error {
	if _cMSDecoderCopyDetachedContent != nil {
		return nil
	}
	return symbolCallError("CMSDecoderCopyDetachedContent", "10.5", _cMSDecoderCopyDetachedContentErr)
}

// TryCMSDecoderCopyDetachedContent calls CMSDecoderCopyDetachedContent without panicking when the symbol is unavailable.
func TryCMSDecoderCopyDetachedContent(cmsDecoder CMSDecoderRef, detachedContentOut *corefoundation.CFDataRef) (int32, error) {
	if err := CMSDecoderCopyDetachedContentCallError(); err != nil {
		return 0, err
	}
	return _cMSDecoderCopyDetachedContent(cmsDecoder, detachedContentOut), nil
}

// CMSDecoderCopyDetachedContent obtains the detached content specified with the [CMSDecoderSetDetachedContent] function.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderCopyDetachedContent(_:_:)
func CMSDecoderCopyDetachedContent(cmsDecoder CMSDecoderRef, detachedContentOut *corefoundation.CFDataRef) int32 {
	result, callErr := TryCMSDecoderCopyDetachedContent(cmsDecoder, detachedContentOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSDecoderCopyEncapsulatedContentType func(cmsDecoder CMSDecoderRef, eContentTypeOut *corefoundation.CFDataRef) int32
var _cMSDecoderCopyEncapsulatedContentTypeErr error

// CanCallCMSDecoderCopyEncapsulatedContentType reports whether the symbol for CMSDecoderCopyEncapsulatedContentType is available on this system.
func CanCallCMSDecoderCopyEncapsulatedContentType() bool {
	return _cMSDecoderCopyEncapsulatedContentType != nil
}

// CMSDecoderCopyEncapsulatedContentTypeCallError returns the symbol lookup or registration error for CMSDecoderCopyEncapsulatedContentType.
func CMSDecoderCopyEncapsulatedContentTypeCallError() error {
	if _cMSDecoderCopyEncapsulatedContentType != nil {
		return nil
	}
	return symbolCallError("CMSDecoderCopyEncapsulatedContentType", "10.5", _cMSDecoderCopyEncapsulatedContentTypeErr)
}

// TryCMSDecoderCopyEncapsulatedContentType calls CMSDecoderCopyEncapsulatedContentType without panicking when the symbol is unavailable.
func TryCMSDecoderCopyEncapsulatedContentType(cmsDecoder CMSDecoderRef, eContentTypeOut *corefoundation.CFDataRef) (int32, error) {
	if err := CMSDecoderCopyEncapsulatedContentTypeCallError(); err != nil {
		return 0, err
	}
	return _cMSDecoderCopyEncapsulatedContentType(cmsDecoder, eContentTypeOut), nil
}

// CMSDecoderCopyEncapsulatedContentType obtains the object identifier for the encapsulated data of a signed message.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderCopyEncapsulatedContentType(_:_:)
func CMSDecoderCopyEncapsulatedContentType(cmsDecoder CMSDecoderRef, eContentTypeOut *corefoundation.CFDataRef) int32 {
	result, callErr := TryCMSDecoderCopyEncapsulatedContentType(cmsDecoder, eContentTypeOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSDecoderCopySignerCert func(cmsDecoder CMSDecoderRef, signerIndex uintptr, signerCertOut *SecCertificateRef) int32
var _cMSDecoderCopySignerCertErr error

// CanCallCMSDecoderCopySignerCert reports whether the symbol for CMSDecoderCopySignerCert is available on this system.
func CanCallCMSDecoderCopySignerCert() bool {
	return _cMSDecoderCopySignerCert != nil
}

// CMSDecoderCopySignerCertCallError returns the symbol lookup or registration error for CMSDecoderCopySignerCert.
func CMSDecoderCopySignerCertCallError() error {
	if _cMSDecoderCopySignerCert != nil {
		return nil
	}
	return symbolCallError("CMSDecoderCopySignerCert", "10.5", _cMSDecoderCopySignerCertErr)
}

// TryCMSDecoderCopySignerCert calls CMSDecoderCopySignerCert without panicking when the symbol is unavailable.
func TryCMSDecoderCopySignerCert(cmsDecoder CMSDecoderRef, signerIndex uintptr, signerCertOut *SecCertificateRef) (int32, error) {
	if err := CMSDecoderCopySignerCertCallError(); err != nil {
		return 0, err
	}
	return _cMSDecoderCopySignerCert(cmsDecoder, signerIndex, signerCertOut), nil
}

// CMSDecoderCopySignerCert obtains the certificate of the specified signer of a CMS message.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderCopySignerCert(_:_:_:)
func CMSDecoderCopySignerCert(cmsDecoder CMSDecoderRef, signerIndex uintptr, signerCertOut *SecCertificateRef) int32 {
	result, callErr := TryCMSDecoderCopySignerCert(cmsDecoder, signerIndex, signerCertOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSDecoderCopySignerEmailAddress func(cmsDecoder CMSDecoderRef, signerIndex uintptr, signerEmailAddressOut *corefoundation.CFStringRef) int32
var _cMSDecoderCopySignerEmailAddressErr error

// CanCallCMSDecoderCopySignerEmailAddress reports whether the symbol for CMSDecoderCopySignerEmailAddress is available on this system.
func CanCallCMSDecoderCopySignerEmailAddress() bool {
	return _cMSDecoderCopySignerEmailAddress != nil
}

// CMSDecoderCopySignerEmailAddressCallError returns the symbol lookup or registration error for CMSDecoderCopySignerEmailAddress.
func CMSDecoderCopySignerEmailAddressCallError() error {
	if _cMSDecoderCopySignerEmailAddress != nil {
		return nil
	}
	return symbolCallError("CMSDecoderCopySignerEmailAddress", "10.5", _cMSDecoderCopySignerEmailAddressErr)
}

// TryCMSDecoderCopySignerEmailAddress calls CMSDecoderCopySignerEmailAddress without panicking when the symbol is unavailable.
func TryCMSDecoderCopySignerEmailAddress(cmsDecoder CMSDecoderRef, signerIndex uintptr, signerEmailAddressOut *corefoundation.CFStringRef) (int32, error) {
	if err := CMSDecoderCopySignerEmailAddressCallError(); err != nil {
		return 0, err
	}
	return _cMSDecoderCopySignerEmailAddress(cmsDecoder, signerIndex, signerEmailAddressOut), nil
}

// CMSDecoderCopySignerEmailAddress obtains the email address of the specified signer of a CMS message.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderCopySignerEmailAddress(_:_:_:)
func CMSDecoderCopySignerEmailAddress(cmsDecoder CMSDecoderRef, signerIndex uintptr, signerEmailAddressOut *corefoundation.CFStringRef) int32 {
	result, callErr := TryCMSDecoderCopySignerEmailAddress(cmsDecoder, signerIndex, signerEmailAddressOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSDecoderCopySignerSigningTime func(cmsDecoder CMSDecoderRef, signerIndex uintptr, signingTime *corefoundation.CFAbsoluteTime) int32
var _cMSDecoderCopySignerSigningTimeErr error

// CanCallCMSDecoderCopySignerSigningTime reports whether the symbol for CMSDecoderCopySignerSigningTime is available on this system.
func CanCallCMSDecoderCopySignerSigningTime() bool {
	return _cMSDecoderCopySignerSigningTime != nil
}

// CMSDecoderCopySignerSigningTimeCallError returns the symbol lookup or registration error for CMSDecoderCopySignerSigningTime.
func CMSDecoderCopySignerSigningTimeCallError() error {
	if _cMSDecoderCopySignerSigningTime != nil {
		return nil
	}
	return symbolCallError("CMSDecoderCopySignerSigningTime", "10.8", _cMSDecoderCopySignerSigningTimeErr)
}

// TryCMSDecoderCopySignerSigningTime calls CMSDecoderCopySignerSigningTime without panicking when the symbol is unavailable.
func TryCMSDecoderCopySignerSigningTime(cmsDecoder CMSDecoderRef, signerIndex uintptr, signingTime *corefoundation.CFAbsoluteTime) (int32, error) {
	if err := CMSDecoderCopySignerSigningTimeCallError(); err != nil {
		return 0, err
	}
	return _cMSDecoderCopySignerSigningTime(cmsDecoder, signerIndex, signingTime), nil
}

// CMSDecoderCopySignerSigningTime obtains the signing time of a CMS message, if present.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderCopySignerSigningTime(_:_:_:)
func CMSDecoderCopySignerSigningTime(cmsDecoder CMSDecoderRef, signerIndex uintptr, signingTime *corefoundation.CFAbsoluteTime) int32 {
	result, callErr := TryCMSDecoderCopySignerSigningTime(cmsDecoder, signerIndex, signingTime)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSDecoderCopySignerStatus func(cmsDecoder CMSDecoderRef, signerIndex uintptr, policyOrArray corefoundation.CFTypeRef, evaluateSecTrust bool, signerStatusOut *CMSSignerStatus, secTrustOut *SecTrustRef, certVerifyResultCodeOut *int32) int32
var _cMSDecoderCopySignerStatusErr error

// CanCallCMSDecoderCopySignerStatus reports whether the symbol for CMSDecoderCopySignerStatus is available on this system.
func CanCallCMSDecoderCopySignerStatus() bool {
	return _cMSDecoderCopySignerStatus != nil
}

// CMSDecoderCopySignerStatusCallError returns the symbol lookup or registration error for CMSDecoderCopySignerStatus.
func CMSDecoderCopySignerStatusCallError() error {
	if _cMSDecoderCopySignerStatus != nil {
		return nil
	}
	return symbolCallError("CMSDecoderCopySignerStatus", "10.5", _cMSDecoderCopySignerStatusErr)
}

// TryCMSDecoderCopySignerStatus calls CMSDecoderCopySignerStatus without panicking when the symbol is unavailable.
func TryCMSDecoderCopySignerStatus(cmsDecoder CMSDecoderRef, signerIndex uintptr, policyOrArray corefoundation.CFTypeRef, evaluateSecTrust bool, signerStatusOut *CMSSignerStatus, secTrustOut *SecTrustRef, certVerifyResultCodeOut *int32) (int32, error) {
	if err := CMSDecoderCopySignerStatusCallError(); err != nil {
		return 0, err
	}
	return _cMSDecoderCopySignerStatus(cmsDecoder, signerIndex, policyOrArray, evaluateSecTrust, signerStatusOut, secTrustOut, certVerifyResultCodeOut), nil
}

// CMSDecoderCopySignerStatus obtains the status of a CMS message’s signature.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderCopySignerStatus(_:_:_:_:_:_:_:)
func CMSDecoderCopySignerStatus(cmsDecoder CMSDecoderRef, signerIndex uintptr, policyOrArray corefoundation.CFTypeRef, evaluateSecTrust bool, signerStatusOut *CMSSignerStatus, secTrustOut *SecTrustRef, certVerifyResultCodeOut *int32) int32 {
	result, callErr := TryCMSDecoderCopySignerStatus(cmsDecoder, signerIndex, policyOrArray, evaluateSecTrust, signerStatusOut, secTrustOut, certVerifyResultCodeOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSDecoderCopySignerTimestamp func(cmsDecoder CMSDecoderRef, signerIndex uintptr, timestamp *corefoundation.CFAbsoluteTime) int32
var _cMSDecoderCopySignerTimestampErr error

// CanCallCMSDecoderCopySignerTimestamp reports whether the symbol for CMSDecoderCopySignerTimestamp is available on this system.
func CanCallCMSDecoderCopySignerTimestamp() bool {
	return _cMSDecoderCopySignerTimestamp != nil
}

// CMSDecoderCopySignerTimestampCallError returns the symbol lookup or registration error for CMSDecoderCopySignerTimestamp.
func CMSDecoderCopySignerTimestampCallError() error {
	if _cMSDecoderCopySignerTimestamp != nil {
		return nil
	}
	return symbolCallError("CMSDecoderCopySignerTimestamp", "10.8", _cMSDecoderCopySignerTimestampErr)
}

// TryCMSDecoderCopySignerTimestamp calls CMSDecoderCopySignerTimestamp without panicking when the symbol is unavailable.
func TryCMSDecoderCopySignerTimestamp(cmsDecoder CMSDecoderRef, signerIndex uintptr, timestamp *corefoundation.CFAbsoluteTime) (int32, error) {
	if err := CMSDecoderCopySignerTimestampCallError(); err != nil {
		return 0, err
	}
	return _cMSDecoderCopySignerTimestamp(cmsDecoder, signerIndex, timestamp), nil
}

// CMSDecoderCopySignerTimestamp returns the timestamp of a signer of a CMS message, if present.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderCopySignerTimestamp(_:_:_:)
func CMSDecoderCopySignerTimestamp(cmsDecoder CMSDecoderRef, signerIndex uintptr, timestamp *corefoundation.CFAbsoluteTime) int32 {
	result, callErr := TryCMSDecoderCopySignerTimestamp(cmsDecoder, signerIndex, timestamp)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSDecoderCopySignerTimestampCertificates func(cmsDecoder CMSDecoderRef, signerIndex uintptr, certificateRefs *corefoundation.CFArrayRef) int32
var _cMSDecoderCopySignerTimestampCertificatesErr error

// CanCallCMSDecoderCopySignerTimestampCertificates reports whether the symbol for CMSDecoderCopySignerTimestampCertificates is available on this system.
func CanCallCMSDecoderCopySignerTimestampCertificates() bool {
	return _cMSDecoderCopySignerTimestampCertificates != nil
}

// CMSDecoderCopySignerTimestampCertificatesCallError returns the symbol lookup or registration error for CMSDecoderCopySignerTimestampCertificates.
func CMSDecoderCopySignerTimestampCertificatesCallError() error {
	if _cMSDecoderCopySignerTimestampCertificates != nil {
		return nil
	}
	return symbolCallError("CMSDecoderCopySignerTimestampCertificates", "10.8", _cMSDecoderCopySignerTimestampCertificatesErr)
}

// TryCMSDecoderCopySignerTimestampCertificates calls CMSDecoderCopySignerTimestampCertificates without panicking when the symbol is unavailable.
func TryCMSDecoderCopySignerTimestampCertificates(cmsDecoder CMSDecoderRef, signerIndex uintptr, certificateRefs *corefoundation.CFArrayRef) (int32, error) {
	if err := CMSDecoderCopySignerTimestampCertificatesCallError(); err != nil {
		return 0, err
	}
	return _cMSDecoderCopySignerTimestampCertificates(cmsDecoder, signerIndex, certificateRefs), nil
}

// CMSDecoderCopySignerTimestampCertificates returns an array containing the certificates from a timestamp response.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderCopySignerTimestampCertificates(_:_:_:)
func CMSDecoderCopySignerTimestampCertificates(cmsDecoder CMSDecoderRef, signerIndex uintptr, certificateRefs *corefoundation.CFArrayRef) int32 {
	result, callErr := TryCMSDecoderCopySignerTimestampCertificates(cmsDecoder, signerIndex, certificateRefs)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSDecoderCopySignerTimestampWithPolicy func(cmsDecoder CMSDecoderRef, timeStampPolicy corefoundation.CFTypeRef, signerIndex uintptr, timestamp *corefoundation.CFAbsoluteTime) int32
var _cMSDecoderCopySignerTimestampWithPolicyErr error

// CanCallCMSDecoderCopySignerTimestampWithPolicy reports whether the symbol for CMSDecoderCopySignerTimestampWithPolicy is available on this system.
func CanCallCMSDecoderCopySignerTimestampWithPolicy() bool {
	return _cMSDecoderCopySignerTimestampWithPolicy != nil
}

// CMSDecoderCopySignerTimestampWithPolicyCallError returns the symbol lookup or registration error for CMSDecoderCopySignerTimestampWithPolicy.
func CMSDecoderCopySignerTimestampWithPolicyCallError() error {
	if _cMSDecoderCopySignerTimestampWithPolicy != nil {
		return nil
	}
	return symbolCallError("CMSDecoderCopySignerTimestampWithPolicy", "10.10", _cMSDecoderCopySignerTimestampWithPolicyErr)
}

// TryCMSDecoderCopySignerTimestampWithPolicy calls CMSDecoderCopySignerTimestampWithPolicy without panicking when the symbol is unavailable.
func TryCMSDecoderCopySignerTimestampWithPolicy(cmsDecoder CMSDecoderRef, timeStampPolicy corefoundation.CFTypeRef, signerIndex uintptr, timestamp *corefoundation.CFAbsoluteTime) (int32, error) {
	if err := CMSDecoderCopySignerTimestampWithPolicyCallError(); err != nil {
		return 0, err
	}
	return _cMSDecoderCopySignerTimestampWithPolicy(cmsDecoder, timeStampPolicy, signerIndex, timestamp), nil
}

// CMSDecoderCopySignerTimestampWithPolicy returns the timestamp of a signer of a CMS message using a given policy, if present.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderCopySignerTimestampWithPolicy(_:_:_:_:)
func CMSDecoderCopySignerTimestampWithPolicy(cmsDecoder CMSDecoderRef, timeStampPolicy corefoundation.CFTypeRef, signerIndex uintptr, timestamp *corefoundation.CFAbsoluteTime) int32 {
	result, callErr := TryCMSDecoderCopySignerTimestampWithPolicy(cmsDecoder, timeStampPolicy, signerIndex, timestamp)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSDecoderCreate func(cmsDecoderOut *CMSDecoderRef) int32
var _cMSDecoderCreateErr error

// CanCallCMSDecoderCreate reports whether the symbol for CMSDecoderCreate is available on this system.
func CanCallCMSDecoderCreate() bool {
	return _cMSDecoderCreate != nil
}

// CMSDecoderCreateCallError returns the symbol lookup or registration error for CMSDecoderCreate.
func CMSDecoderCreateCallError() error {
	if _cMSDecoderCreate != nil {
		return nil
	}
	return symbolCallError("CMSDecoderCreate", "10.5", _cMSDecoderCreateErr)
}

// TryCMSDecoderCreate calls CMSDecoderCreate without panicking when the symbol is unavailable.
func TryCMSDecoderCreate(cmsDecoderOut *CMSDecoderRef) (int32, error) {
	if err := CMSDecoderCreateCallError(); err != nil {
		return 0, err
	}
	return _cMSDecoderCreate(cmsDecoderOut), nil
}

// CMSDecoderCreate creates a CMSDecoder reference.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderCreate(_:)
func CMSDecoderCreate(cmsDecoderOut *CMSDecoderRef) int32 {
	result, callErr := TryCMSDecoderCreate(cmsDecoderOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSDecoderFinalizeMessage func(cmsDecoder CMSDecoderRef) int32
var _cMSDecoderFinalizeMessageErr error

// CanCallCMSDecoderFinalizeMessage reports whether the symbol for CMSDecoderFinalizeMessage is available on this system.
func CanCallCMSDecoderFinalizeMessage() bool {
	return _cMSDecoderFinalizeMessage != nil
}

// CMSDecoderFinalizeMessageCallError returns the symbol lookup or registration error for CMSDecoderFinalizeMessage.
func CMSDecoderFinalizeMessageCallError() error {
	if _cMSDecoderFinalizeMessage != nil {
		return nil
	}
	return symbolCallError("CMSDecoderFinalizeMessage", "10.5", _cMSDecoderFinalizeMessageErr)
}

// TryCMSDecoderFinalizeMessage calls CMSDecoderFinalizeMessage without panicking when the symbol is unavailable.
func TryCMSDecoderFinalizeMessage(cmsDecoder CMSDecoderRef) (int32, error) {
	if err := CMSDecoderFinalizeMessageCallError(); err != nil {
		return 0, err
	}
	return _cMSDecoderFinalizeMessage(cmsDecoder), nil
}

// CMSDecoderFinalizeMessage indicates that there is no more data to decode.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderFinalizeMessage(_:)
func CMSDecoderFinalizeMessage(cmsDecoder CMSDecoderRef) int32 {
	result, callErr := TryCMSDecoderFinalizeMessage(cmsDecoder)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSDecoderGetNumSigners func(cmsDecoder CMSDecoderRef, numSignersOut *uintptr) int32
var _cMSDecoderGetNumSignersErr error

// CanCallCMSDecoderGetNumSigners reports whether the symbol for CMSDecoderGetNumSigners is available on this system.
func CanCallCMSDecoderGetNumSigners() bool {
	return _cMSDecoderGetNumSigners != nil
}

// CMSDecoderGetNumSignersCallError returns the symbol lookup or registration error for CMSDecoderGetNumSigners.
func CMSDecoderGetNumSignersCallError() error {
	if _cMSDecoderGetNumSigners != nil {
		return nil
	}
	return symbolCallError("CMSDecoderGetNumSigners", "10.5", _cMSDecoderGetNumSignersErr)
}

// TryCMSDecoderGetNumSigners calls CMSDecoderGetNumSigners without panicking when the symbol is unavailable.
func TryCMSDecoderGetNumSigners(cmsDecoder CMSDecoderRef, numSignersOut *uintptr) (int32, error) {
	if err := CMSDecoderGetNumSignersCallError(); err != nil {
		return 0, err
	}
	return _cMSDecoderGetNumSigners(cmsDecoder, numSignersOut), nil
}

// CMSDecoderGetNumSigners obtains the number of signers of a message.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderGetNumSigners(_:_:)
func CMSDecoderGetNumSigners(cmsDecoder CMSDecoderRef, numSignersOut *uintptr) int32 {
	result, callErr := TryCMSDecoderGetNumSigners(cmsDecoder, numSignersOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSDecoderGetTypeID func() uint
var _cMSDecoderGetTypeIDErr error

// CanCallCMSDecoderGetTypeID reports whether the symbol for CMSDecoderGetTypeID is available on this system.
func CanCallCMSDecoderGetTypeID() bool {
	return _cMSDecoderGetTypeID != nil
}

// CMSDecoderGetTypeIDCallError returns the symbol lookup or registration error for CMSDecoderGetTypeID.
func CMSDecoderGetTypeIDCallError() error {
	if _cMSDecoderGetTypeID != nil {
		return nil
	}
	return symbolCallError("CMSDecoderGetTypeID", "10.0", _cMSDecoderGetTypeIDErr)
}

// TryCMSDecoderGetTypeID calls CMSDecoderGetTypeID without panicking when the symbol is unavailable.
func TryCMSDecoderGetTypeID() (uint, error) {
	if err := CMSDecoderGetTypeIDCallError(); err != nil {
		return 0, err
	}
	return _cMSDecoderGetTypeID(), nil
}

// CMSDecoderGetTypeID returns the type identifier for the CMSDecoder opaque type.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderGetTypeID()
func CMSDecoderGetTypeID() uint {
	result, callErr := TryCMSDecoderGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSDecoderIsContentEncrypted func(cmsDecoder CMSDecoderRef, isEncryptedOut *bool) int32
var _cMSDecoderIsContentEncryptedErr error

// CanCallCMSDecoderIsContentEncrypted reports whether the symbol for CMSDecoderIsContentEncrypted is available on this system.
func CanCallCMSDecoderIsContentEncrypted() bool {
	return _cMSDecoderIsContentEncrypted != nil
}

// CMSDecoderIsContentEncryptedCallError returns the symbol lookup or registration error for CMSDecoderIsContentEncrypted.
func CMSDecoderIsContentEncryptedCallError() error {
	if _cMSDecoderIsContentEncrypted != nil {
		return nil
	}
	return symbolCallError("CMSDecoderIsContentEncrypted", "10.5", _cMSDecoderIsContentEncryptedErr)
}

// TryCMSDecoderIsContentEncrypted calls CMSDecoderIsContentEncrypted without panicking when the symbol is unavailable.
func TryCMSDecoderIsContentEncrypted(cmsDecoder CMSDecoderRef, isEncryptedOut *bool) (int32, error) {
	if err := CMSDecoderIsContentEncryptedCallError(); err != nil {
		return 0, err
	}
	return _cMSDecoderIsContentEncrypted(cmsDecoder, isEncryptedOut), nil
}

// CMSDecoderIsContentEncrypted determines whether a CMS message was encrypted.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderIsContentEncrypted(_:_:)
func CMSDecoderIsContentEncrypted(cmsDecoder CMSDecoderRef, isEncryptedOut *bool) int32 {
	result, callErr := TryCMSDecoderIsContentEncrypted(cmsDecoder, isEncryptedOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSDecoderSetDetachedContent func(cmsDecoder CMSDecoderRef, detachedContent corefoundation.CFDataRef) int32
var _cMSDecoderSetDetachedContentErr error

// CanCallCMSDecoderSetDetachedContent reports whether the symbol for CMSDecoderSetDetachedContent is available on this system.
func CanCallCMSDecoderSetDetachedContent() bool {
	return _cMSDecoderSetDetachedContent != nil
}

// CMSDecoderSetDetachedContentCallError returns the symbol lookup or registration error for CMSDecoderSetDetachedContent.
func CMSDecoderSetDetachedContentCallError() error {
	if _cMSDecoderSetDetachedContent != nil {
		return nil
	}
	return symbolCallError("CMSDecoderSetDetachedContent", "10.5", _cMSDecoderSetDetachedContentErr)
}

// TryCMSDecoderSetDetachedContent calls CMSDecoderSetDetachedContent without panicking when the symbol is unavailable.
func TryCMSDecoderSetDetachedContent(cmsDecoder CMSDecoderRef, detachedContent corefoundation.CFDataRef) (int32, error) {
	if err := CMSDecoderSetDetachedContentCallError(); err != nil {
		return 0, err
	}
	return _cMSDecoderSetDetachedContent(cmsDecoder, detachedContent), nil
}

// CMSDecoderSetDetachedContent specifies the message’s detached content, if any.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderSetDetachedContent(_:_:)
func CMSDecoderSetDetachedContent(cmsDecoder CMSDecoderRef, detachedContent corefoundation.CFDataRef) int32 {
	result, callErr := TryCMSDecoderSetDetachedContent(cmsDecoder, detachedContent)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSDecoderSetSearchKeychain func(cmsDecoder CMSDecoderRef, keychainOrArray corefoundation.CFTypeRef) int32
var _cMSDecoderSetSearchKeychainErr error

// CanCallCMSDecoderSetSearchKeychain reports whether the symbol for CMSDecoderSetSearchKeychain is available on this system.
func CanCallCMSDecoderSetSearchKeychain() bool {
	return _cMSDecoderSetSearchKeychain != nil
}

// CMSDecoderSetSearchKeychainCallError returns the symbol lookup or registration error for CMSDecoderSetSearchKeychain.
func CMSDecoderSetSearchKeychainCallError() error {
	if _cMSDecoderSetSearchKeychain != nil {
		return nil
	}
	return symbolCallError("CMSDecoderSetSearchKeychain", "10.5", _cMSDecoderSetSearchKeychainErr)
}

// TryCMSDecoderSetSearchKeychain calls CMSDecoderSetSearchKeychain without panicking when the symbol is unavailable.
func TryCMSDecoderSetSearchKeychain(cmsDecoder CMSDecoderRef, keychainOrArray corefoundation.CFTypeRef) (int32, error) {
	if err := CMSDecoderSetSearchKeychainCallError(); err != nil {
		return 0, err
	}
	return _cMSDecoderSetSearchKeychain(cmsDecoder, keychainOrArray), nil
}

// CMSDecoderSetSearchKeychain specifies the keychains to search for intermediate certificates to be used in verifying a signed message’s signer certificates.
//
// Deprecated: Deprecated since macOS 10.13.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderSetSearchKeychain(_:_:)
func CMSDecoderSetSearchKeychain(cmsDecoder CMSDecoderRef, keychainOrArray corefoundation.CFTypeRef) int32 {
	result, callErr := TryCMSDecoderSetSearchKeychain(cmsDecoder, keychainOrArray)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSDecoderUpdateMessage func(cmsDecoder CMSDecoderRef, msgBytes unsafe.Pointer, msgBytesLen uintptr) int32
var _cMSDecoderUpdateMessageErr error

// CanCallCMSDecoderUpdateMessage reports whether the symbol for CMSDecoderUpdateMessage is available on this system.
func CanCallCMSDecoderUpdateMessage() bool {
	return _cMSDecoderUpdateMessage != nil
}

// CMSDecoderUpdateMessageCallError returns the symbol lookup or registration error for CMSDecoderUpdateMessage.
func CMSDecoderUpdateMessageCallError() error {
	if _cMSDecoderUpdateMessage != nil {
		return nil
	}
	return symbolCallError("CMSDecoderUpdateMessage", "10.5", _cMSDecoderUpdateMessageErr)
}

// TryCMSDecoderUpdateMessage calls CMSDecoderUpdateMessage without panicking when the symbol is unavailable.
func TryCMSDecoderUpdateMessage(cmsDecoder CMSDecoderRef, msgBytes unsafe.Pointer, msgBytesLen uintptr) (int32, error) {
	if err := CMSDecoderUpdateMessageCallError(); err != nil {
		return 0, err
	}
	return _cMSDecoderUpdateMessage(cmsDecoder, msgBytes, msgBytesLen), nil
}

// CMSDecoderUpdateMessage feeds raw bytes of the message to be decoded into the decoder.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderUpdateMessage(_:_:_:)
func CMSDecoderUpdateMessage(cmsDecoder CMSDecoderRef, msgBytes unsafe.Pointer, msgBytesLen uintptr) int32 {
	result, callErr := TryCMSDecoderUpdateMessage(cmsDecoder, msgBytes, msgBytesLen)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncode func(signers corefoundation.CFTypeRef, recipients corefoundation.CFTypeRef, eContentType unsafe.Pointer, detachedContent bool, signedAttributes CMSSignedAttributes, content unsafe.Pointer, contentLen uintptr, encodedContentOut *corefoundation.CFDataRef) int32
var _cMSEncodeErr error

// CanCallCMSEncode reports whether the symbol for CMSEncode is available on this system.
func CanCallCMSEncode() bool {
	return _cMSEncode != nil
}

// CMSEncodeCallError returns the symbol lookup or registration error for CMSEncode.
func CMSEncodeCallError() error {
	if _cMSEncode != nil {
		return nil
	}
	return symbolCallError("CMSEncode", "10.5", _cMSEncodeErr)
}

// TryCMSEncode calls CMSEncode without panicking when the symbol is unavailable.
func TryCMSEncode(signers corefoundation.CFTypeRef, recipients corefoundation.CFTypeRef, eContentType unsafe.Pointer, detachedContent bool, signedAttributes CMSSignedAttributes, content unsafe.Pointer, contentLen uintptr, encodedContentOut *corefoundation.CFDataRef) (int32, error) {
	if err := CMSEncodeCallError(); err != nil {
		return 0, err
	}
	return _cMSEncode(signers, recipients, eContentType, detachedContent, signedAttributes, content, contentLen, encodedContentOut), nil
}

// CMSEncode encodes a message and obtains the result in one high-level function call.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CMSEncode
func CMSEncode(signers corefoundation.CFTypeRef, recipients corefoundation.CFTypeRef, eContentType unsafe.Pointer, detachedContent bool, signedAttributes CMSSignedAttributes, content unsafe.Pointer, contentLen uintptr, encodedContentOut *corefoundation.CFDataRef) int32 {
	result, callErr := TryCMSEncode(signers, recipients, eContentType, detachedContent, signedAttributes, content, contentLen, encodedContentOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncodeContent func(signers corefoundation.CFTypeRef, recipients corefoundation.CFTypeRef, eContentTypeOID corefoundation.CFTypeRef, detachedContent bool, signedAttributes CMSSignedAttributes, content unsafe.Pointer, contentLen uintptr, encodedContentOut *corefoundation.CFDataRef) int32
var _cMSEncodeContentErr error

// CanCallCMSEncodeContent reports whether the symbol for CMSEncodeContent is available on this system.
func CanCallCMSEncodeContent() bool {
	return _cMSEncodeContent != nil
}

// CMSEncodeContentCallError returns the symbol lookup or registration error for CMSEncodeContent.
func CMSEncodeContentCallError() error {
	if _cMSEncodeContent != nil {
		return nil
	}
	return symbolCallError("CMSEncodeContent", "10.7", _cMSEncodeContentErr)
}

// TryCMSEncodeContent calls CMSEncodeContent without panicking when the symbol is unavailable.
func TryCMSEncodeContent(signers corefoundation.CFTypeRef, recipients corefoundation.CFTypeRef, eContentTypeOID corefoundation.CFTypeRef, detachedContent bool, signedAttributes CMSSignedAttributes, content unsafe.Pointer, contentLen uintptr, encodedContentOut *corefoundation.CFDataRef) (int32, error) {
	if err := CMSEncodeContentCallError(); err != nil {
		return 0, err
	}
	return _cMSEncodeContent(signers, recipients, eContentTypeOID, detachedContent, signedAttributes, content, contentLen, encodedContentOut), nil
}

// CMSEncodeContent encodes a message and obtains the result in one high-level function call.
//
// See: https://developer.apple.com/documentation/Security/CMSEncodeContent(_:_:_:_:_:_:_:_:)
func CMSEncodeContent(signers corefoundation.CFTypeRef, recipients corefoundation.CFTypeRef, eContentTypeOID corefoundation.CFTypeRef, detachedContent bool, signedAttributes CMSSignedAttributes, content unsafe.Pointer, contentLen uintptr, encodedContentOut *corefoundation.CFDataRef) int32 {
	result, callErr := TryCMSEncodeContent(signers, recipients, eContentTypeOID, detachedContent, signedAttributes, content, contentLen, encodedContentOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderAddRecipients func(cmsEncoder CMSEncoderRef, recipientOrArray corefoundation.CFTypeRef) int32
var _cMSEncoderAddRecipientsErr error

// CanCallCMSEncoderAddRecipients reports whether the symbol for CMSEncoderAddRecipients is available on this system.
func CanCallCMSEncoderAddRecipients() bool {
	return _cMSEncoderAddRecipients != nil
}

// CMSEncoderAddRecipientsCallError returns the symbol lookup or registration error for CMSEncoderAddRecipients.
func CMSEncoderAddRecipientsCallError() error {
	if _cMSEncoderAddRecipients != nil {
		return nil
	}
	return symbolCallError("CMSEncoderAddRecipients", "10.5", _cMSEncoderAddRecipientsErr)
}

// TryCMSEncoderAddRecipients calls CMSEncoderAddRecipients without panicking when the symbol is unavailable.
func TryCMSEncoderAddRecipients(cmsEncoder CMSEncoderRef, recipientOrArray corefoundation.CFTypeRef) (int32, error) {
	if err := CMSEncoderAddRecipientsCallError(); err != nil {
		return 0, err
	}
	return _cMSEncoderAddRecipients(cmsEncoder, recipientOrArray), nil
}

// CMSEncoderAddRecipients specifies a message is to be encrypted and specifies the recipients of the message.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderAddRecipients(_:_:)
func CMSEncoderAddRecipients(cmsEncoder CMSEncoderRef, recipientOrArray corefoundation.CFTypeRef) int32 {
	result, callErr := TryCMSEncoderAddRecipients(cmsEncoder, recipientOrArray)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderAddSignedAttributes func(cmsEncoder CMSEncoderRef, signedAttributes CMSSignedAttributes) int32
var _cMSEncoderAddSignedAttributesErr error

// CanCallCMSEncoderAddSignedAttributes reports whether the symbol for CMSEncoderAddSignedAttributes is available on this system.
func CanCallCMSEncoderAddSignedAttributes() bool {
	return _cMSEncoderAddSignedAttributes != nil
}

// CMSEncoderAddSignedAttributesCallError returns the symbol lookup or registration error for CMSEncoderAddSignedAttributes.
func CMSEncoderAddSignedAttributesCallError() error {
	if _cMSEncoderAddSignedAttributes != nil {
		return nil
	}
	return symbolCallError("CMSEncoderAddSignedAttributes", "10.5", _cMSEncoderAddSignedAttributesErr)
}

// TryCMSEncoderAddSignedAttributes calls CMSEncoderAddSignedAttributes without panicking when the symbol is unavailable.
func TryCMSEncoderAddSignedAttributes(cmsEncoder CMSEncoderRef, signedAttributes CMSSignedAttributes) (int32, error) {
	if err := CMSEncoderAddSignedAttributesCallError(); err != nil {
		return 0, err
	}
	return _cMSEncoderAddSignedAttributes(cmsEncoder, signedAttributes), nil
}

// CMSEncoderAddSignedAttributes specifies attributes for a signed message.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderAddSignedAttributes(_:_:)
func CMSEncoderAddSignedAttributes(cmsEncoder CMSEncoderRef, signedAttributes CMSSignedAttributes) int32 {
	result, callErr := TryCMSEncoderAddSignedAttributes(cmsEncoder, signedAttributes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderAddSigners func(cmsEncoder CMSEncoderRef, signerOrArray corefoundation.CFTypeRef) int32
var _cMSEncoderAddSignersErr error

// CanCallCMSEncoderAddSigners reports whether the symbol for CMSEncoderAddSigners is available on this system.
func CanCallCMSEncoderAddSigners() bool {
	return _cMSEncoderAddSigners != nil
}

// CMSEncoderAddSignersCallError returns the symbol lookup or registration error for CMSEncoderAddSigners.
func CMSEncoderAddSignersCallError() error {
	if _cMSEncoderAddSigners != nil {
		return nil
	}
	return symbolCallError("CMSEncoderAddSigners", "10.5", _cMSEncoderAddSignersErr)
}

// TryCMSEncoderAddSigners calls CMSEncoderAddSigners without panicking when the symbol is unavailable.
func TryCMSEncoderAddSigners(cmsEncoder CMSEncoderRef, signerOrArray corefoundation.CFTypeRef) (int32, error) {
	if err := CMSEncoderAddSignersCallError(); err != nil {
		return 0, err
	}
	return _cMSEncoderAddSigners(cmsEncoder, signerOrArray), nil
}

// CMSEncoderAddSigners specifies signers of the message.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderAddSigners(_:_:)
func CMSEncoderAddSigners(cmsEncoder CMSEncoderRef, signerOrArray corefoundation.CFTypeRef) int32 {
	result, callErr := TryCMSEncoderAddSigners(cmsEncoder, signerOrArray)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderAddSupportingCerts func(cmsEncoder CMSEncoderRef, certOrArray corefoundation.CFTypeRef) int32
var _cMSEncoderAddSupportingCertsErr error

// CanCallCMSEncoderAddSupportingCerts reports whether the symbol for CMSEncoderAddSupportingCerts is available on this system.
func CanCallCMSEncoderAddSupportingCerts() bool {
	return _cMSEncoderAddSupportingCerts != nil
}

// CMSEncoderAddSupportingCertsCallError returns the symbol lookup or registration error for CMSEncoderAddSupportingCerts.
func CMSEncoderAddSupportingCertsCallError() error {
	if _cMSEncoderAddSupportingCerts != nil {
		return nil
	}
	return symbolCallError("CMSEncoderAddSupportingCerts", "10.5", _cMSEncoderAddSupportingCertsErr)
}

// TryCMSEncoderAddSupportingCerts calls CMSEncoderAddSupportingCerts without panicking when the symbol is unavailable.
func TryCMSEncoderAddSupportingCerts(cmsEncoder CMSEncoderRef, certOrArray corefoundation.CFTypeRef) (int32, error) {
	if err := CMSEncoderAddSupportingCertsCallError(); err != nil {
		return 0, err
	}
	return _cMSEncoderAddSupportingCerts(cmsEncoder, certOrArray), nil
}

// CMSEncoderAddSupportingCerts adds certificates to a message.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderAddSupportingCerts(_:_:)
func CMSEncoderAddSupportingCerts(cmsEncoder CMSEncoderRef, certOrArray corefoundation.CFTypeRef) int32 {
	result, callErr := TryCMSEncoderAddSupportingCerts(cmsEncoder, certOrArray)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderCopyEncapsulatedContentType func(cmsEncoder CMSEncoderRef, eContentTypeOut *corefoundation.CFDataRef) int32
var _cMSEncoderCopyEncapsulatedContentTypeErr error

// CanCallCMSEncoderCopyEncapsulatedContentType reports whether the symbol for CMSEncoderCopyEncapsulatedContentType is available on this system.
func CanCallCMSEncoderCopyEncapsulatedContentType() bool {
	return _cMSEncoderCopyEncapsulatedContentType != nil
}

// CMSEncoderCopyEncapsulatedContentTypeCallError returns the symbol lookup or registration error for CMSEncoderCopyEncapsulatedContentType.
func CMSEncoderCopyEncapsulatedContentTypeCallError() error {
	if _cMSEncoderCopyEncapsulatedContentType != nil {
		return nil
	}
	return symbolCallError("CMSEncoderCopyEncapsulatedContentType", "10.5", _cMSEncoderCopyEncapsulatedContentTypeErr)
}

// TryCMSEncoderCopyEncapsulatedContentType calls CMSEncoderCopyEncapsulatedContentType without panicking when the symbol is unavailable.
func TryCMSEncoderCopyEncapsulatedContentType(cmsEncoder CMSEncoderRef, eContentTypeOut *corefoundation.CFDataRef) (int32, error) {
	if err := CMSEncoderCopyEncapsulatedContentTypeCallError(); err != nil {
		return 0, err
	}
	return _cMSEncoderCopyEncapsulatedContentType(cmsEncoder, eContentTypeOut), nil
}

// CMSEncoderCopyEncapsulatedContentType obtains the object identifier for the encapsulated data of a signed message.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderCopyEncapsulatedContentType(_:_:)
func CMSEncoderCopyEncapsulatedContentType(cmsEncoder CMSEncoderRef, eContentTypeOut *corefoundation.CFDataRef) int32 {
	result, callErr := TryCMSEncoderCopyEncapsulatedContentType(cmsEncoder, eContentTypeOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderCopyEncodedContent func(cmsEncoder CMSEncoderRef, encodedContentOut *corefoundation.CFDataRef) int32
var _cMSEncoderCopyEncodedContentErr error

// CanCallCMSEncoderCopyEncodedContent reports whether the symbol for CMSEncoderCopyEncodedContent is available on this system.
func CanCallCMSEncoderCopyEncodedContent() bool {
	return _cMSEncoderCopyEncodedContent != nil
}

// CMSEncoderCopyEncodedContentCallError returns the symbol lookup or registration error for CMSEncoderCopyEncodedContent.
func CMSEncoderCopyEncodedContentCallError() error {
	if _cMSEncoderCopyEncodedContent != nil {
		return nil
	}
	return symbolCallError("CMSEncoderCopyEncodedContent", "10.5", _cMSEncoderCopyEncodedContentErr)
}

// TryCMSEncoderCopyEncodedContent calls CMSEncoderCopyEncodedContent without panicking when the symbol is unavailable.
func TryCMSEncoderCopyEncodedContent(cmsEncoder CMSEncoderRef, encodedContentOut *corefoundation.CFDataRef) (int32, error) {
	if err := CMSEncoderCopyEncodedContentCallError(); err != nil {
		return 0, err
	}
	return _cMSEncoderCopyEncodedContent(cmsEncoder, encodedContentOut), nil
}

// CMSEncoderCopyEncodedContent finishes encoding the message and obtains the encoded result.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderCopyEncodedContent(_:_:)
func CMSEncoderCopyEncodedContent(cmsEncoder CMSEncoderRef, encodedContentOut *corefoundation.CFDataRef) int32 {
	result, callErr := TryCMSEncoderCopyEncodedContent(cmsEncoder, encodedContentOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderCopyRecipients func(cmsEncoder CMSEncoderRef, recipientsOut *corefoundation.CFArrayRef) int32
var _cMSEncoderCopyRecipientsErr error

// CanCallCMSEncoderCopyRecipients reports whether the symbol for CMSEncoderCopyRecipients is available on this system.
func CanCallCMSEncoderCopyRecipients() bool {
	return _cMSEncoderCopyRecipients != nil
}

// CMSEncoderCopyRecipientsCallError returns the symbol lookup or registration error for CMSEncoderCopyRecipients.
func CMSEncoderCopyRecipientsCallError() error {
	if _cMSEncoderCopyRecipients != nil {
		return nil
	}
	return symbolCallError("CMSEncoderCopyRecipients", "10.5", _cMSEncoderCopyRecipientsErr)
}

// TryCMSEncoderCopyRecipients calls CMSEncoderCopyRecipients without panicking when the symbol is unavailable.
func TryCMSEncoderCopyRecipients(cmsEncoder CMSEncoderRef, recipientsOut *corefoundation.CFArrayRef) (int32, error) {
	if err := CMSEncoderCopyRecipientsCallError(); err != nil {
		return 0, err
	}
	return _cMSEncoderCopyRecipients(cmsEncoder, recipientsOut), nil
}

// CMSEncoderCopyRecipients obtains the array of recipients specified with the [CMSEncoderAddRecipients] function.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderCopyRecipients(_:_:)
func CMSEncoderCopyRecipients(cmsEncoder CMSEncoderRef, recipientsOut *corefoundation.CFArrayRef) int32 {
	result, callErr := TryCMSEncoderCopyRecipients(cmsEncoder, recipientsOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderCopySignerTimestamp func(cmsEncoder CMSEncoderRef, signerIndex uintptr, timestamp *corefoundation.CFAbsoluteTime) int32
var _cMSEncoderCopySignerTimestampErr error

// CanCallCMSEncoderCopySignerTimestamp reports whether the symbol for CMSEncoderCopySignerTimestamp is available on this system.
func CanCallCMSEncoderCopySignerTimestamp() bool {
	return _cMSEncoderCopySignerTimestamp != nil
}

// CMSEncoderCopySignerTimestampCallError returns the symbol lookup or registration error for CMSEncoderCopySignerTimestamp.
func CMSEncoderCopySignerTimestampCallError() error {
	if _cMSEncoderCopySignerTimestamp != nil {
		return nil
	}
	return symbolCallError("CMSEncoderCopySignerTimestamp", "10.8", _cMSEncoderCopySignerTimestampErr)
}

// TryCMSEncoderCopySignerTimestamp calls CMSEncoderCopySignerTimestamp without panicking when the symbol is unavailable.
func TryCMSEncoderCopySignerTimestamp(cmsEncoder CMSEncoderRef, signerIndex uintptr, timestamp *corefoundation.CFAbsoluteTime) (int32, error) {
	if err := CMSEncoderCopySignerTimestampCallError(); err != nil {
		return 0, err
	}
	return _cMSEncoderCopySignerTimestamp(cmsEncoder, signerIndex, timestamp), nil
}

// CMSEncoderCopySignerTimestamp returns the timestamp of a signer of a CMS message, if present.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderCopySignerTimestamp(_:_:_:)
func CMSEncoderCopySignerTimestamp(cmsEncoder CMSEncoderRef, signerIndex uintptr, timestamp *corefoundation.CFAbsoluteTime) int32 {
	result, callErr := TryCMSEncoderCopySignerTimestamp(cmsEncoder, signerIndex, timestamp)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderCopySignerTimestampWithPolicy func(cmsEncoder CMSEncoderRef, timeStampPolicy corefoundation.CFTypeRef, signerIndex uintptr, timestamp *corefoundation.CFAbsoluteTime) int32
var _cMSEncoderCopySignerTimestampWithPolicyErr error

// CanCallCMSEncoderCopySignerTimestampWithPolicy reports whether the symbol for CMSEncoderCopySignerTimestampWithPolicy is available on this system.
func CanCallCMSEncoderCopySignerTimestampWithPolicy() bool {
	return _cMSEncoderCopySignerTimestampWithPolicy != nil
}

// CMSEncoderCopySignerTimestampWithPolicyCallError returns the symbol lookup or registration error for CMSEncoderCopySignerTimestampWithPolicy.
func CMSEncoderCopySignerTimestampWithPolicyCallError() error {
	if _cMSEncoderCopySignerTimestampWithPolicy != nil {
		return nil
	}
	return symbolCallError("CMSEncoderCopySignerTimestampWithPolicy", "10.10", _cMSEncoderCopySignerTimestampWithPolicyErr)
}

// TryCMSEncoderCopySignerTimestampWithPolicy calls CMSEncoderCopySignerTimestampWithPolicy without panicking when the symbol is unavailable.
func TryCMSEncoderCopySignerTimestampWithPolicy(cmsEncoder CMSEncoderRef, timeStampPolicy corefoundation.CFTypeRef, signerIndex uintptr, timestamp *corefoundation.CFAbsoluteTime) (int32, error) {
	if err := CMSEncoderCopySignerTimestampWithPolicyCallError(); err != nil {
		return 0, err
	}
	return _cMSEncoderCopySignerTimestampWithPolicy(cmsEncoder, timeStampPolicy, signerIndex, timestamp), nil
}

// CMSEncoderCopySignerTimestampWithPolicy returns the timestamp of a signer of a CMS message using a particular policy, if present.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderCopySignerTimestampWithPolicy(_:_:_:_:)
func CMSEncoderCopySignerTimestampWithPolicy(cmsEncoder CMSEncoderRef, timeStampPolicy corefoundation.CFTypeRef, signerIndex uintptr, timestamp *corefoundation.CFAbsoluteTime) int32 {
	result, callErr := TryCMSEncoderCopySignerTimestampWithPolicy(cmsEncoder, timeStampPolicy, signerIndex, timestamp)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderCopySigners func(cmsEncoder CMSEncoderRef, signersOut *corefoundation.CFArrayRef) int32
var _cMSEncoderCopySignersErr error

// CanCallCMSEncoderCopySigners reports whether the symbol for CMSEncoderCopySigners is available on this system.
func CanCallCMSEncoderCopySigners() bool {
	return _cMSEncoderCopySigners != nil
}

// CMSEncoderCopySignersCallError returns the symbol lookup or registration error for CMSEncoderCopySigners.
func CMSEncoderCopySignersCallError() error {
	if _cMSEncoderCopySigners != nil {
		return nil
	}
	return symbolCallError("CMSEncoderCopySigners", "10.5", _cMSEncoderCopySignersErr)
}

// TryCMSEncoderCopySigners calls CMSEncoderCopySigners without panicking when the symbol is unavailable.
func TryCMSEncoderCopySigners(cmsEncoder CMSEncoderRef, signersOut *corefoundation.CFArrayRef) (int32, error) {
	if err := CMSEncoderCopySignersCallError(); err != nil {
		return 0, err
	}
	return _cMSEncoderCopySigners(cmsEncoder, signersOut), nil
}

// CMSEncoderCopySigners obtains the array of signers specified with the [CMSEncoderAddSigners] function.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderCopySigners(_:_:)
func CMSEncoderCopySigners(cmsEncoder CMSEncoderRef, signersOut *corefoundation.CFArrayRef) int32 {
	result, callErr := TryCMSEncoderCopySigners(cmsEncoder, signersOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderCopySupportingCerts func(cmsEncoder CMSEncoderRef, certsOut *corefoundation.CFArrayRef) int32
var _cMSEncoderCopySupportingCertsErr error

// CanCallCMSEncoderCopySupportingCerts reports whether the symbol for CMSEncoderCopySupportingCerts is available on this system.
func CanCallCMSEncoderCopySupportingCerts() bool {
	return _cMSEncoderCopySupportingCerts != nil
}

// CMSEncoderCopySupportingCertsCallError returns the symbol lookup or registration error for CMSEncoderCopySupportingCerts.
func CMSEncoderCopySupportingCertsCallError() error {
	if _cMSEncoderCopySupportingCerts != nil {
		return nil
	}
	return symbolCallError("CMSEncoderCopySupportingCerts", "10.5", _cMSEncoderCopySupportingCertsErr)
}

// TryCMSEncoderCopySupportingCerts calls CMSEncoderCopySupportingCerts without panicking when the symbol is unavailable.
func TryCMSEncoderCopySupportingCerts(cmsEncoder CMSEncoderRef, certsOut *corefoundation.CFArrayRef) (int32, error) {
	if err := CMSEncoderCopySupportingCertsCallError(); err != nil {
		return 0, err
	}
	return _cMSEncoderCopySupportingCerts(cmsEncoder, certsOut), nil
}

// CMSEncoderCopySupportingCerts obtains the certificates added to a message with [CMSEncoderAddSupportingCerts].
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderCopySupportingCerts(_:_:)
func CMSEncoderCopySupportingCerts(cmsEncoder CMSEncoderRef, certsOut *corefoundation.CFArrayRef) int32 {
	result, callErr := TryCMSEncoderCopySupportingCerts(cmsEncoder, certsOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderCreate func(cmsEncoderOut *CMSEncoderRef) int32
var _cMSEncoderCreateErr error

// CanCallCMSEncoderCreate reports whether the symbol for CMSEncoderCreate is available on this system.
func CanCallCMSEncoderCreate() bool {
	return _cMSEncoderCreate != nil
}

// CMSEncoderCreateCallError returns the symbol lookup or registration error for CMSEncoderCreate.
func CMSEncoderCreateCallError() error {
	if _cMSEncoderCreate != nil {
		return nil
	}
	return symbolCallError("CMSEncoderCreate", "10.5", _cMSEncoderCreateErr)
}

// TryCMSEncoderCreate calls CMSEncoderCreate without panicking when the symbol is unavailable.
func TryCMSEncoderCreate(cmsEncoderOut *CMSEncoderRef) (int32, error) {
	if err := CMSEncoderCreateCallError(); err != nil {
		return 0, err
	}
	return _cMSEncoderCreate(cmsEncoderOut), nil
}

// CMSEncoderCreate creates a CMSEncoder reference.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderCreate(_:)
func CMSEncoderCreate(cmsEncoderOut *CMSEncoderRef) int32 {
	result, callErr := TryCMSEncoderCreate(cmsEncoderOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderGetCertificateChainMode func(cmsEncoder CMSEncoderRef, chainModeOut *CMSCertificateChainMode) int32
var _cMSEncoderGetCertificateChainModeErr error

// CanCallCMSEncoderGetCertificateChainMode reports whether the symbol for CMSEncoderGetCertificateChainMode is available on this system.
func CanCallCMSEncoderGetCertificateChainMode() bool {
	return _cMSEncoderGetCertificateChainMode != nil
}

// CMSEncoderGetCertificateChainModeCallError returns the symbol lookup or registration error for CMSEncoderGetCertificateChainMode.
func CMSEncoderGetCertificateChainModeCallError() error {
	if _cMSEncoderGetCertificateChainMode != nil {
		return nil
	}
	return symbolCallError("CMSEncoderGetCertificateChainMode", "10.5", _cMSEncoderGetCertificateChainModeErr)
}

// TryCMSEncoderGetCertificateChainMode calls CMSEncoderGetCertificateChainMode without panicking when the symbol is unavailable.
func TryCMSEncoderGetCertificateChainMode(cmsEncoder CMSEncoderRef, chainModeOut *CMSCertificateChainMode) (int32, error) {
	if err := CMSEncoderGetCertificateChainModeCallError(); err != nil {
		return 0, err
	}
	return _cMSEncoderGetCertificateChainMode(cmsEncoder, chainModeOut), nil
}

// CMSEncoderGetCertificateChainMode obtains a constant that indicates which certificates are to be included in a signed CMS message.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderGetCertificateChainMode(_:_:)
func CMSEncoderGetCertificateChainMode(cmsEncoder CMSEncoderRef, chainModeOut *CMSCertificateChainMode) int32 {
	result, callErr := TryCMSEncoderGetCertificateChainMode(cmsEncoder, chainModeOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderGetHasDetachedContent func(cmsEncoder CMSEncoderRef, detachedContentOut *bool) int32
var _cMSEncoderGetHasDetachedContentErr error

// CanCallCMSEncoderGetHasDetachedContent reports whether the symbol for CMSEncoderGetHasDetachedContent is available on this system.
func CanCallCMSEncoderGetHasDetachedContent() bool {
	return _cMSEncoderGetHasDetachedContent != nil
}

// CMSEncoderGetHasDetachedContentCallError returns the symbol lookup or registration error for CMSEncoderGetHasDetachedContent.
func CMSEncoderGetHasDetachedContentCallError() error {
	if _cMSEncoderGetHasDetachedContent != nil {
		return nil
	}
	return symbolCallError("CMSEncoderGetHasDetachedContent", "10.5", _cMSEncoderGetHasDetachedContentErr)
}

// TryCMSEncoderGetHasDetachedContent calls CMSEncoderGetHasDetachedContent without panicking when the symbol is unavailable.
func TryCMSEncoderGetHasDetachedContent(cmsEncoder CMSEncoderRef, detachedContentOut *bool) (int32, error) {
	if err := CMSEncoderGetHasDetachedContentCallError(); err != nil {
		return 0, err
	}
	return _cMSEncoderGetHasDetachedContent(cmsEncoder, detachedContentOut), nil
}

// CMSEncoderGetHasDetachedContent indicates whether the message is to have detached content.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderGetHasDetachedContent(_:_:)
func CMSEncoderGetHasDetachedContent(cmsEncoder CMSEncoderRef, detachedContentOut *bool) int32 {
	result, callErr := TryCMSEncoderGetHasDetachedContent(cmsEncoder, detachedContentOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderGetTypeID func() uint
var _cMSEncoderGetTypeIDErr error

// CanCallCMSEncoderGetTypeID reports whether the symbol for CMSEncoderGetTypeID is available on this system.
func CanCallCMSEncoderGetTypeID() bool {
	return _cMSEncoderGetTypeID != nil
}

// CMSEncoderGetTypeIDCallError returns the symbol lookup or registration error for CMSEncoderGetTypeID.
func CMSEncoderGetTypeIDCallError() error {
	if _cMSEncoderGetTypeID != nil {
		return nil
	}
	return symbolCallError("CMSEncoderGetTypeID", "10.5", _cMSEncoderGetTypeIDErr)
}

// TryCMSEncoderGetTypeID calls CMSEncoderGetTypeID without panicking when the symbol is unavailable.
func TryCMSEncoderGetTypeID() (uint, error) {
	if err := CMSEncoderGetTypeIDCallError(); err != nil {
		return 0, err
	}
	return _cMSEncoderGetTypeID(), nil
}

// CMSEncoderGetTypeID returns the type identifier for the CMSEncoder opaque type.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderGetTypeID()
func CMSEncoderGetTypeID() uint {
	result, callErr := TryCMSEncoderGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderSetCertificateChainMode func(cmsEncoder CMSEncoderRef, chainMode CMSCertificateChainMode) int32
var _cMSEncoderSetCertificateChainModeErr error

// CanCallCMSEncoderSetCertificateChainMode reports whether the symbol for CMSEncoderSetCertificateChainMode is available on this system.
func CanCallCMSEncoderSetCertificateChainMode() bool {
	return _cMSEncoderSetCertificateChainMode != nil
}

// CMSEncoderSetCertificateChainModeCallError returns the symbol lookup or registration error for CMSEncoderSetCertificateChainMode.
func CMSEncoderSetCertificateChainModeCallError() error {
	if _cMSEncoderSetCertificateChainMode != nil {
		return nil
	}
	return symbolCallError("CMSEncoderSetCertificateChainMode", "10.5", _cMSEncoderSetCertificateChainModeErr)
}

// TryCMSEncoderSetCertificateChainMode calls CMSEncoderSetCertificateChainMode without panicking when the symbol is unavailable.
func TryCMSEncoderSetCertificateChainMode(cmsEncoder CMSEncoderRef, chainMode CMSCertificateChainMode) (int32, error) {
	if err := CMSEncoderSetCertificateChainModeCallError(); err != nil {
		return 0, err
	}
	return _cMSEncoderSetCertificateChainMode(cmsEncoder, chainMode), nil
}

// CMSEncoderSetCertificateChainMode specifies which certificates to include in a signed CMS message.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderSetCertificateChainMode(_:_:)
func CMSEncoderSetCertificateChainMode(cmsEncoder CMSEncoderRef, chainMode CMSCertificateChainMode) int32 {
	result, callErr := TryCMSEncoderSetCertificateChainMode(cmsEncoder, chainMode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderSetEncapsulatedContentType func(cmsEncoder CMSEncoderRef, eContentType unsafe.Pointer) int32
var _cMSEncoderSetEncapsulatedContentTypeErr error

// CanCallCMSEncoderSetEncapsulatedContentType reports whether the symbol for CMSEncoderSetEncapsulatedContentType is available on this system.
func CanCallCMSEncoderSetEncapsulatedContentType() bool {
	return _cMSEncoderSetEncapsulatedContentType != nil
}

// CMSEncoderSetEncapsulatedContentTypeCallError returns the symbol lookup or registration error for CMSEncoderSetEncapsulatedContentType.
func CMSEncoderSetEncapsulatedContentTypeCallError() error {
	if _cMSEncoderSetEncapsulatedContentType != nil {
		return nil
	}
	return symbolCallError("CMSEncoderSetEncapsulatedContentType", "10.5", _cMSEncoderSetEncapsulatedContentTypeErr)
}

// TryCMSEncoderSetEncapsulatedContentType calls CMSEncoderSetEncapsulatedContentType without panicking when the symbol is unavailable.
func TryCMSEncoderSetEncapsulatedContentType(cmsEncoder CMSEncoderRef, eContentType unsafe.Pointer) (int32, error) {
	if err := CMSEncoderSetEncapsulatedContentTypeCallError(); err != nil {
		return 0, err
	}
	return _cMSEncoderSetEncapsulatedContentType(cmsEncoder, eContentType), nil
}

// CMSEncoderSetEncapsulatedContentType specifies an object identifier for the encapsulated data of a signed message.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderSetEncapsulatedContentType
func CMSEncoderSetEncapsulatedContentType(cmsEncoder CMSEncoderRef, eContentType unsafe.Pointer) int32 {
	result, callErr := TryCMSEncoderSetEncapsulatedContentType(cmsEncoder, eContentType)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderSetEncapsulatedContentTypeOID func(cmsEncoder CMSEncoderRef, eContentTypeOID corefoundation.CFTypeRef) int32
var _cMSEncoderSetEncapsulatedContentTypeOIDErr error

// CanCallCMSEncoderSetEncapsulatedContentTypeOID reports whether the symbol for CMSEncoderSetEncapsulatedContentTypeOID is available on this system.
func CanCallCMSEncoderSetEncapsulatedContentTypeOID() bool {
	return _cMSEncoderSetEncapsulatedContentTypeOID != nil
}

// CMSEncoderSetEncapsulatedContentTypeOIDCallError returns the symbol lookup or registration error for CMSEncoderSetEncapsulatedContentTypeOID.
func CMSEncoderSetEncapsulatedContentTypeOIDCallError() error {
	if _cMSEncoderSetEncapsulatedContentTypeOID != nil {
		return nil
	}
	return symbolCallError("CMSEncoderSetEncapsulatedContentTypeOID", "10.7", _cMSEncoderSetEncapsulatedContentTypeOIDErr)
}

// TryCMSEncoderSetEncapsulatedContentTypeOID calls CMSEncoderSetEncapsulatedContentTypeOID without panicking when the symbol is unavailable.
func TryCMSEncoderSetEncapsulatedContentTypeOID(cmsEncoder CMSEncoderRef, eContentTypeOID corefoundation.CFTypeRef) (int32, error) {
	if err := CMSEncoderSetEncapsulatedContentTypeOIDCallError(); err != nil {
		return 0, err
	}
	return _cMSEncoderSetEncapsulatedContentTypeOID(cmsEncoder, eContentTypeOID), nil
}

// CMSEncoderSetEncapsulatedContentTypeOID specifies an object identifier for the encapsulated data of a signed message.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderSetEncapsulatedContentTypeOID(_:_:)
func CMSEncoderSetEncapsulatedContentTypeOID(cmsEncoder CMSEncoderRef, eContentTypeOID corefoundation.CFTypeRef) int32 {
	result, callErr := TryCMSEncoderSetEncapsulatedContentTypeOID(cmsEncoder, eContentTypeOID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderSetHasDetachedContent func(cmsEncoder CMSEncoderRef, detachedContent bool) int32
var _cMSEncoderSetHasDetachedContentErr error

// CanCallCMSEncoderSetHasDetachedContent reports whether the symbol for CMSEncoderSetHasDetachedContent is available on this system.
func CanCallCMSEncoderSetHasDetachedContent() bool {
	return _cMSEncoderSetHasDetachedContent != nil
}

// CMSEncoderSetHasDetachedContentCallError returns the symbol lookup or registration error for CMSEncoderSetHasDetachedContent.
func CMSEncoderSetHasDetachedContentCallError() error {
	if _cMSEncoderSetHasDetachedContent != nil {
		return nil
	}
	return symbolCallError("CMSEncoderSetHasDetachedContent", "10.5", _cMSEncoderSetHasDetachedContentErr)
}

// TryCMSEncoderSetHasDetachedContent calls CMSEncoderSetHasDetachedContent without panicking when the symbol is unavailable.
func TryCMSEncoderSetHasDetachedContent(cmsEncoder CMSEncoderRef, detachedContent bool) (int32, error) {
	if err := CMSEncoderSetHasDetachedContentCallError(); err != nil {
		return 0, err
	}
	return _cMSEncoderSetHasDetachedContent(cmsEncoder, detachedContent), nil
}

// CMSEncoderSetHasDetachedContent specifies whether the signed data is to be separate from the message.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderSetHasDetachedContent(_:_:)
func CMSEncoderSetHasDetachedContent(cmsEncoder CMSEncoderRef, detachedContent bool) int32 {
	result, callErr := TryCMSEncoderSetHasDetachedContent(cmsEncoder, detachedContent)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderSetSignerAlgorithm func(cmsEncoder CMSEncoderRef, digestAlgorithm corefoundation.CFStringRef) int32
var _cMSEncoderSetSignerAlgorithmErr error

// CanCallCMSEncoderSetSignerAlgorithm reports whether the symbol for CMSEncoderSetSignerAlgorithm is available on this system.
func CanCallCMSEncoderSetSignerAlgorithm() bool {
	return _cMSEncoderSetSignerAlgorithm != nil
}

// CMSEncoderSetSignerAlgorithmCallError returns the symbol lookup or registration error for CMSEncoderSetSignerAlgorithm.
func CMSEncoderSetSignerAlgorithmCallError() error {
	if _cMSEncoderSetSignerAlgorithm != nil {
		return nil
	}
	return symbolCallError("CMSEncoderSetSignerAlgorithm", "10.11", _cMSEncoderSetSignerAlgorithmErr)
}

// TryCMSEncoderSetSignerAlgorithm calls CMSEncoderSetSignerAlgorithm without panicking when the symbol is unavailable.
func TryCMSEncoderSetSignerAlgorithm(cmsEncoder CMSEncoderRef, digestAlgorithm corefoundation.CFStringRef) (int32, error) {
	if err := CMSEncoderSetSignerAlgorithmCallError(); err != nil {
		return 0, err
	}
	return _cMSEncoderSetSignerAlgorithm(cmsEncoder, digestAlgorithm), nil
}

// CMSEncoderSetSignerAlgorithm sets the digest algorithm to use for the signer.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderSetSignerAlgorithm(_:_:)
func CMSEncoderSetSignerAlgorithm(cmsEncoder CMSEncoderRef, digestAlgorithm corefoundation.CFStringRef) int32 {
	result, callErr := TryCMSEncoderSetSignerAlgorithm(cmsEncoder, digestAlgorithm)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSEncoderUpdateContent func(cmsEncoder CMSEncoderRef, content unsafe.Pointer, contentLen uintptr) int32
var _cMSEncoderUpdateContentErr error

// CanCallCMSEncoderUpdateContent reports whether the symbol for CMSEncoderUpdateContent is available on this system.
func CanCallCMSEncoderUpdateContent() bool {
	return _cMSEncoderUpdateContent != nil
}

// CMSEncoderUpdateContentCallError returns the symbol lookup or registration error for CMSEncoderUpdateContent.
func CMSEncoderUpdateContentCallError() error {
	if _cMSEncoderUpdateContent != nil {
		return nil
	}
	return symbolCallError("CMSEncoderUpdateContent", "10.5", _cMSEncoderUpdateContentErr)
}

// TryCMSEncoderUpdateContent calls CMSEncoderUpdateContent without panicking when the symbol is unavailable.
func TryCMSEncoderUpdateContent(cmsEncoder CMSEncoderRef, content unsafe.Pointer, contentLen uintptr) (int32, error) {
	if err := CMSEncoderUpdateContentCallError(); err != nil {
		return 0, err
	}
	return _cMSEncoderUpdateContent(cmsEncoder, content, contentLen), nil
}

// CMSEncoderUpdateContent feeds content bytes into the encoder.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderUpdateContent(_:_:_:)
func CMSEncoderUpdateContent(cmsEncoder CMSEncoderRef, content unsafe.Pointer, contentLen uintptr) int32 {
	result, callErr := TryCMSEncoderUpdateContent(cmsEncoder, content, contentLen)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_AC_AuthCompute func(ACHandle CSSM_AC_HANDLE, BaseAuthorizations unsafe.Pointer, Credentials unsafe.Pointer, NumberOfRequestors uint32, Requestors unsafe.Pointer, RequestedAuthorizationPeriod unsafe.Pointer, RequestedAuthorization unsafe.Pointer, AuthorizationResult unsafe.Pointer) CSSM_RETURN
var _cSSM_AC_AuthComputeErr error

// CanCallCSSM_AC_AuthCompute reports whether the symbol for CSSM_AC_AuthCompute is available on this system.
func CanCallCSSM_AC_AuthCompute() bool {
	return _cSSM_AC_AuthCompute != nil
}

// CSSM_AC_AuthComputeCallError returns the symbol lookup or registration error for CSSM_AC_AuthCompute.
func CSSM_AC_AuthComputeCallError() error {
	if _cSSM_AC_AuthCompute != nil {
		return nil
	}
	return symbolCallError("CSSM_AC_AuthCompute", "10.0", _cSSM_AC_AuthComputeErr)
}

// TryCSSM_AC_AuthCompute calls CSSM_AC_AuthCompute without panicking when the symbol is unavailable.
func TryCSSM_AC_AuthCompute(ACHandle CSSM_AC_HANDLE, BaseAuthorizations unsafe.Pointer, Credentials unsafe.Pointer, NumberOfRequestors uint32, Requestors unsafe.Pointer, RequestedAuthorizationPeriod unsafe.Pointer, RequestedAuthorization unsafe.Pointer, AuthorizationResult unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_AC_AuthComputeCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_AC_AuthCompute(ACHandle, BaseAuthorizations, Credentials, NumberOfRequestors, Requestors, RequestedAuthorizationPeriod, RequestedAuthorization, AuthorizationResult), nil
}

// CSSM_AC_AuthCompute.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_AC_AuthCompute
func CSSM_AC_AuthCompute(ACHandle CSSM_AC_HANDLE, BaseAuthorizations unsafe.Pointer, Credentials unsafe.Pointer, NumberOfRequestors uint32, Requestors unsafe.Pointer, RequestedAuthorizationPeriod unsafe.Pointer, RequestedAuthorization unsafe.Pointer, AuthorizationResult unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_AC_AuthCompute(ACHandle, BaseAuthorizations, Credentials, NumberOfRequestors, Requestors, RequestedAuthorizationPeriod, RequestedAuthorization, AuthorizationResult)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_AC_PassThrough func(ACHandle CSSM_AC_HANDLE, TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, DBList unsafe.Pointer, PassThroughId uint32, InputParams unsafe.Pointer, OutputParams unsafe.Pointer) CSSM_RETURN
var _cSSM_AC_PassThroughErr error

// CanCallCSSM_AC_PassThrough reports whether the symbol for CSSM_AC_PassThrough is available on this system.
func CanCallCSSM_AC_PassThrough() bool {
	return _cSSM_AC_PassThrough != nil
}

// CSSM_AC_PassThroughCallError returns the symbol lookup or registration error for CSSM_AC_PassThrough.
func CSSM_AC_PassThroughCallError() error {
	if _cSSM_AC_PassThrough != nil {
		return nil
	}
	return symbolCallError("CSSM_AC_PassThrough", "10.0", _cSSM_AC_PassThroughErr)
}

// TryCSSM_AC_PassThrough calls CSSM_AC_PassThrough without panicking when the symbol is unavailable.
func TryCSSM_AC_PassThrough(ACHandle CSSM_AC_HANDLE, TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, DBList unsafe.Pointer, PassThroughId uint32, InputParams unsafe.Pointer, OutputParams unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_AC_PassThroughCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_AC_PassThrough(ACHandle, TPHandle, CLHandle, CCHandle, DBList, PassThroughId, InputParams, OutputParams), nil
}

// CSSM_AC_PassThrough.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_AC_PassThrough
func CSSM_AC_PassThrough(ACHandle CSSM_AC_HANDLE, TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, DBList unsafe.Pointer, PassThroughId uint32, InputParams unsafe.Pointer, OutputParams unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_AC_PassThrough(ACHandle, TPHandle, CLHandle, CCHandle, DBList, PassThroughId, InputParams, OutputParams)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CertAbortCache func(CLHandle CSSM_CL_HANDLE, CertHandle CSSM_HANDLE) CSSM_RETURN
var _cSSM_CL_CertAbortCacheErr error

// CanCallCSSM_CL_CertAbortCache reports whether the symbol for CSSM_CL_CertAbortCache is available on this system.
func CanCallCSSM_CL_CertAbortCache() bool {
	return _cSSM_CL_CertAbortCache != nil
}

// CSSM_CL_CertAbortCacheCallError returns the symbol lookup or registration error for CSSM_CL_CertAbortCache.
func CSSM_CL_CertAbortCacheCallError() error {
	if _cSSM_CL_CertAbortCache != nil {
		return nil
	}
	return symbolCallError("CSSM_CL_CertAbortCache", "10.0", _cSSM_CL_CertAbortCacheErr)
}

// TryCSSM_CL_CertAbortCache calls CSSM_CL_CertAbortCache without panicking when the symbol is unavailable.
func TryCSSM_CL_CertAbortCache(CLHandle CSSM_CL_HANDLE, CertHandle CSSM_HANDLE) (CSSM_RETURN, error) {
	if err := CSSM_CL_CertAbortCacheCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CL_CertAbortCache(CLHandle, CertHandle), nil
}

// CSSM_CL_CertAbortCache.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertAbortCache
func CSSM_CL_CertAbortCache(CLHandle CSSM_CL_HANDLE, CertHandle CSSM_HANDLE) CSSM_RETURN {
	result, callErr := TryCSSM_CL_CertAbortCache(CLHandle, CertHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CertAbortQuery func(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE) CSSM_RETURN
var _cSSM_CL_CertAbortQueryErr error

// CanCallCSSM_CL_CertAbortQuery reports whether the symbol for CSSM_CL_CertAbortQuery is available on this system.
func CanCallCSSM_CL_CertAbortQuery() bool {
	return _cSSM_CL_CertAbortQuery != nil
}

// CSSM_CL_CertAbortQueryCallError returns the symbol lookup or registration error for CSSM_CL_CertAbortQuery.
func CSSM_CL_CertAbortQueryCallError() error {
	if _cSSM_CL_CertAbortQuery != nil {
		return nil
	}
	return symbolCallError("CSSM_CL_CertAbortQuery", "10.0", _cSSM_CL_CertAbortQueryErr)
}

// TryCSSM_CL_CertAbortQuery calls CSSM_CL_CertAbortQuery without panicking when the symbol is unavailable.
func TryCSSM_CL_CertAbortQuery(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE) (CSSM_RETURN, error) {
	if err := CSSM_CL_CertAbortQueryCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CL_CertAbortQuery(CLHandle, ResultsHandle), nil
}

// CSSM_CL_CertAbortQuery.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertAbortQuery
func CSSM_CL_CertAbortQuery(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE) CSSM_RETURN {
	result, callErr := TryCSSM_CL_CertAbortQuery(CLHandle, ResultsHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CertCache func(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, CertHandle CSSM_HANDLE_PTR) CSSM_RETURN
var _cSSM_CL_CertCacheErr error

// CanCallCSSM_CL_CertCache reports whether the symbol for CSSM_CL_CertCache is available on this system.
func CanCallCSSM_CL_CertCache() bool {
	return _cSSM_CL_CertCache != nil
}

// CSSM_CL_CertCacheCallError returns the symbol lookup or registration error for CSSM_CL_CertCache.
func CSSM_CL_CertCacheCallError() error {
	if _cSSM_CL_CertCache != nil {
		return nil
	}
	return symbolCallError("CSSM_CL_CertCache", "10.0", _cSSM_CL_CertCacheErr)
}

// TryCSSM_CL_CertCache calls CSSM_CL_CertCache without panicking when the symbol is unavailable.
func TryCSSM_CL_CertCache(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, CertHandle CSSM_HANDLE_PTR) (CSSM_RETURN, error) {
	if err := CSSM_CL_CertCacheCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CL_CertCache(CLHandle, Cert, CertHandle), nil
}

// CSSM_CL_CertCache.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertCache
func CSSM_CL_CertCache(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, CertHandle CSSM_HANDLE_PTR) CSSM_RETURN {
	result, callErr := TryCSSM_CL_CertCache(CLHandle, Cert, CertHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CertCreateTemplate func(CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CertFields unsafe.Pointer, CertTemplate unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CertCreateTemplateErr error

// CanCallCSSM_CL_CertCreateTemplate reports whether the symbol for CSSM_CL_CertCreateTemplate is available on this system.
func CanCallCSSM_CL_CertCreateTemplate() bool {
	return _cSSM_CL_CertCreateTemplate != nil
}

// CSSM_CL_CertCreateTemplateCallError returns the symbol lookup or registration error for CSSM_CL_CertCreateTemplate.
func CSSM_CL_CertCreateTemplateCallError() error {
	if _cSSM_CL_CertCreateTemplate != nil {
		return nil
	}
	return symbolCallError("CSSM_CL_CertCreateTemplate", "10.0", _cSSM_CL_CertCreateTemplateErr)
}

// TryCSSM_CL_CertCreateTemplate calls CSSM_CL_CertCreateTemplate without panicking when the symbol is unavailable.
func TryCSSM_CL_CertCreateTemplate(CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CertFields unsafe.Pointer, CertTemplate unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CL_CertCreateTemplateCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CL_CertCreateTemplate(CLHandle, NumberOfFields, CertFields, CertTemplate), nil
}

// CSSM_CL_CertCreateTemplate.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertCreateTemplate
func CSSM_CL_CertCreateTemplate(CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CertFields unsafe.Pointer, CertTemplate unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CL_CertCreateTemplate(CLHandle, NumberOfFields, CertFields, CertTemplate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CertDescribeFormat func(CLHandle CSSM_CL_HANDLE, NumberOfFields *uint32, OidList unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CertDescribeFormatErr error

// CanCallCSSM_CL_CertDescribeFormat reports whether the symbol for CSSM_CL_CertDescribeFormat is available on this system.
func CanCallCSSM_CL_CertDescribeFormat() bool {
	return _cSSM_CL_CertDescribeFormat != nil
}

// CSSM_CL_CertDescribeFormatCallError returns the symbol lookup or registration error for CSSM_CL_CertDescribeFormat.
func CSSM_CL_CertDescribeFormatCallError() error {
	if _cSSM_CL_CertDescribeFormat != nil {
		return nil
	}
	return symbolCallError("CSSM_CL_CertDescribeFormat", "10.0", _cSSM_CL_CertDescribeFormatErr)
}

// TryCSSM_CL_CertDescribeFormat calls CSSM_CL_CertDescribeFormat without panicking when the symbol is unavailable.
func TryCSSM_CL_CertDescribeFormat(CLHandle CSSM_CL_HANDLE, NumberOfFields *uint32, OidList unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CL_CertDescribeFormatCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CL_CertDescribeFormat(CLHandle, NumberOfFields, OidList), nil
}

// CSSM_CL_CertDescribeFormat.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertDescribeFormat
func CSSM_CL_CertDescribeFormat(CLHandle CSSM_CL_HANDLE, NumberOfFields *uint32, OidList unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CL_CertDescribeFormat(CLHandle, NumberOfFields, OidList)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CertGetAllFields func(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, NumberOfFields *uint32, CertFields unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CertGetAllFieldsErr error

// CanCallCSSM_CL_CertGetAllFields reports whether the symbol for CSSM_CL_CertGetAllFields is available on this system.
func CanCallCSSM_CL_CertGetAllFields() bool {
	return _cSSM_CL_CertGetAllFields != nil
}

// CSSM_CL_CertGetAllFieldsCallError returns the symbol lookup or registration error for CSSM_CL_CertGetAllFields.
func CSSM_CL_CertGetAllFieldsCallError() error {
	if _cSSM_CL_CertGetAllFields != nil {
		return nil
	}
	return symbolCallError("CSSM_CL_CertGetAllFields", "10.0", _cSSM_CL_CertGetAllFieldsErr)
}

// TryCSSM_CL_CertGetAllFields calls CSSM_CL_CertGetAllFields without panicking when the symbol is unavailable.
func TryCSSM_CL_CertGetAllFields(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, NumberOfFields *uint32, CertFields unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CL_CertGetAllFieldsCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CL_CertGetAllFields(CLHandle, Cert, NumberOfFields, CertFields), nil
}

// CSSM_CL_CertGetAllFields.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertGetAllFields
func CSSM_CL_CertGetAllFields(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, NumberOfFields *uint32, CertFields unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CL_CertGetAllFields(CLHandle, Cert, NumberOfFields, CertFields)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CertGetAllTemplateFields func(CLHandle CSSM_CL_HANDLE, CertTemplate unsafe.Pointer, NumberOfFields *uint32, CertFields unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CertGetAllTemplateFieldsErr error

// CanCallCSSM_CL_CertGetAllTemplateFields reports whether the symbol for CSSM_CL_CertGetAllTemplateFields is available on this system.
func CanCallCSSM_CL_CertGetAllTemplateFields() bool {
	return _cSSM_CL_CertGetAllTemplateFields != nil
}

// CSSM_CL_CertGetAllTemplateFieldsCallError returns the symbol lookup or registration error for CSSM_CL_CertGetAllTemplateFields.
func CSSM_CL_CertGetAllTemplateFieldsCallError() error {
	if _cSSM_CL_CertGetAllTemplateFields != nil {
		return nil
	}
	return symbolCallError("CSSM_CL_CertGetAllTemplateFields", "10.0", _cSSM_CL_CertGetAllTemplateFieldsErr)
}

// TryCSSM_CL_CertGetAllTemplateFields calls CSSM_CL_CertGetAllTemplateFields without panicking when the symbol is unavailable.
func TryCSSM_CL_CertGetAllTemplateFields(CLHandle CSSM_CL_HANDLE, CertTemplate unsafe.Pointer, NumberOfFields *uint32, CertFields unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CL_CertGetAllTemplateFieldsCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CL_CertGetAllTemplateFields(CLHandle, CertTemplate, NumberOfFields, CertFields), nil
}

// CSSM_CL_CertGetAllTemplateFields.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertGetAllTemplateFields
func CSSM_CL_CertGetAllTemplateFields(CLHandle CSSM_CL_HANDLE, CertTemplate unsafe.Pointer, NumberOfFields *uint32, CertFields unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CL_CertGetAllTemplateFields(CLHandle, CertTemplate, NumberOfFields, CertFields)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CertGetFirstCachedFieldValue func(CLHandle CSSM_CL_HANDLE, CertHandle CSSM_HANDLE, CertField unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, NumberOfMatchedFields *uint32, Value unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CertGetFirstCachedFieldValueErr error

// CanCallCSSM_CL_CertGetFirstCachedFieldValue reports whether the symbol for CSSM_CL_CertGetFirstCachedFieldValue is available on this system.
func CanCallCSSM_CL_CertGetFirstCachedFieldValue() bool {
	return _cSSM_CL_CertGetFirstCachedFieldValue != nil
}

// CSSM_CL_CertGetFirstCachedFieldValueCallError returns the symbol lookup or registration error for CSSM_CL_CertGetFirstCachedFieldValue.
func CSSM_CL_CertGetFirstCachedFieldValueCallError() error {
	if _cSSM_CL_CertGetFirstCachedFieldValue != nil {
		return nil
	}
	return symbolCallError("CSSM_CL_CertGetFirstCachedFieldValue", "10.0", _cSSM_CL_CertGetFirstCachedFieldValueErr)
}

// TryCSSM_CL_CertGetFirstCachedFieldValue calls CSSM_CL_CertGetFirstCachedFieldValue without panicking when the symbol is unavailable.
func TryCSSM_CL_CertGetFirstCachedFieldValue(CLHandle CSSM_CL_HANDLE, CertHandle CSSM_HANDLE, CertField unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, NumberOfMatchedFields *uint32, Value unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CL_CertGetFirstCachedFieldValueCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CL_CertGetFirstCachedFieldValue(CLHandle, CertHandle, CertField, ResultsHandle, NumberOfMatchedFields, Value), nil
}

// CSSM_CL_CertGetFirstCachedFieldValue.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertGetFirstCachedFieldValue
func CSSM_CL_CertGetFirstCachedFieldValue(CLHandle CSSM_CL_HANDLE, CertHandle CSSM_HANDLE, CertField unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, NumberOfMatchedFields *uint32, Value unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CL_CertGetFirstCachedFieldValue(CLHandle, CertHandle, CertField, ResultsHandle, NumberOfMatchedFields, Value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CertGetFirstFieldValue func(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, CertField unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, NumberOfMatchedFields *uint32, Value unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CertGetFirstFieldValueErr error

// CanCallCSSM_CL_CertGetFirstFieldValue reports whether the symbol for CSSM_CL_CertGetFirstFieldValue is available on this system.
func CanCallCSSM_CL_CertGetFirstFieldValue() bool {
	return _cSSM_CL_CertGetFirstFieldValue != nil
}

// CSSM_CL_CertGetFirstFieldValueCallError returns the symbol lookup or registration error for CSSM_CL_CertGetFirstFieldValue.
func CSSM_CL_CertGetFirstFieldValueCallError() error {
	if _cSSM_CL_CertGetFirstFieldValue != nil {
		return nil
	}
	return symbolCallError("CSSM_CL_CertGetFirstFieldValue", "10.0", _cSSM_CL_CertGetFirstFieldValueErr)
}

// TryCSSM_CL_CertGetFirstFieldValue calls CSSM_CL_CertGetFirstFieldValue without panicking when the symbol is unavailable.
func TryCSSM_CL_CertGetFirstFieldValue(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, CertField unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, NumberOfMatchedFields *uint32, Value unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CL_CertGetFirstFieldValueCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CL_CertGetFirstFieldValue(CLHandle, Cert, CertField, ResultsHandle, NumberOfMatchedFields, Value), nil
}

// CSSM_CL_CertGetFirstFieldValue.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertGetFirstFieldValue
func CSSM_CL_CertGetFirstFieldValue(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, CertField unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, NumberOfMatchedFields *uint32, Value unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CL_CertGetFirstFieldValue(CLHandle, Cert, CertField, ResultsHandle, NumberOfMatchedFields, Value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CertGetKeyInfo func(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, Key unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CertGetKeyInfoErr error

// CanCallCSSM_CL_CertGetKeyInfo reports whether the symbol for CSSM_CL_CertGetKeyInfo is available on this system.
func CanCallCSSM_CL_CertGetKeyInfo() bool {
	return _cSSM_CL_CertGetKeyInfo != nil
}

// CSSM_CL_CertGetKeyInfoCallError returns the symbol lookup or registration error for CSSM_CL_CertGetKeyInfo.
func CSSM_CL_CertGetKeyInfoCallError() error {
	if _cSSM_CL_CertGetKeyInfo != nil {
		return nil
	}
	return symbolCallError("CSSM_CL_CertGetKeyInfo", "10.0", _cSSM_CL_CertGetKeyInfoErr)
}

// TryCSSM_CL_CertGetKeyInfo calls CSSM_CL_CertGetKeyInfo without panicking when the symbol is unavailable.
func TryCSSM_CL_CertGetKeyInfo(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, Key unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CL_CertGetKeyInfoCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CL_CertGetKeyInfo(CLHandle, Cert, Key), nil
}

// CSSM_CL_CertGetKeyInfo.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertGetKeyInfo
func CSSM_CL_CertGetKeyInfo(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, Key unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CL_CertGetKeyInfo(CLHandle, Cert, Key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CertGetNextCachedFieldValue func(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE, Value unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CertGetNextCachedFieldValueErr error

// CanCallCSSM_CL_CertGetNextCachedFieldValue reports whether the symbol for CSSM_CL_CertGetNextCachedFieldValue is available on this system.
func CanCallCSSM_CL_CertGetNextCachedFieldValue() bool {
	return _cSSM_CL_CertGetNextCachedFieldValue != nil
}

// CSSM_CL_CertGetNextCachedFieldValueCallError returns the symbol lookup or registration error for CSSM_CL_CertGetNextCachedFieldValue.
func CSSM_CL_CertGetNextCachedFieldValueCallError() error {
	if _cSSM_CL_CertGetNextCachedFieldValue != nil {
		return nil
	}
	return symbolCallError("CSSM_CL_CertGetNextCachedFieldValue", "10.0", _cSSM_CL_CertGetNextCachedFieldValueErr)
}

// TryCSSM_CL_CertGetNextCachedFieldValue calls CSSM_CL_CertGetNextCachedFieldValue without panicking when the symbol is unavailable.
func TryCSSM_CL_CertGetNextCachedFieldValue(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE, Value unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CL_CertGetNextCachedFieldValueCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CL_CertGetNextCachedFieldValue(CLHandle, ResultsHandle, Value), nil
}

// CSSM_CL_CertGetNextCachedFieldValue.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertGetNextCachedFieldValue
func CSSM_CL_CertGetNextCachedFieldValue(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE, Value unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CL_CertGetNextCachedFieldValue(CLHandle, ResultsHandle, Value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CertGetNextFieldValue func(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE, Value unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CertGetNextFieldValueErr error

// CanCallCSSM_CL_CertGetNextFieldValue reports whether the symbol for CSSM_CL_CertGetNextFieldValue is available on this system.
func CanCallCSSM_CL_CertGetNextFieldValue() bool {
	return _cSSM_CL_CertGetNextFieldValue != nil
}

// CSSM_CL_CertGetNextFieldValueCallError returns the symbol lookup or registration error for CSSM_CL_CertGetNextFieldValue.
func CSSM_CL_CertGetNextFieldValueCallError() error {
	if _cSSM_CL_CertGetNextFieldValue != nil {
		return nil
	}
	return symbolCallError("CSSM_CL_CertGetNextFieldValue", "10.0", _cSSM_CL_CertGetNextFieldValueErr)
}

// TryCSSM_CL_CertGetNextFieldValue calls CSSM_CL_CertGetNextFieldValue without panicking when the symbol is unavailable.
func TryCSSM_CL_CertGetNextFieldValue(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE, Value unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CL_CertGetNextFieldValueCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CL_CertGetNextFieldValue(CLHandle, ResultsHandle, Value), nil
}

// CSSM_CL_CertGetNextFieldValue.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertGetNextFieldValue
func CSSM_CL_CertGetNextFieldValue(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE, Value unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CL_CertGetNextFieldValue(CLHandle, ResultsHandle, Value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CertGroupFromVerifiedBundle func(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertBundle unsafe.Pointer, SignerCert unsafe.Pointer, CertGroup unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CertGroupFromVerifiedBundleErr error

// CanCallCSSM_CL_CertGroupFromVerifiedBundle reports whether the symbol for CSSM_CL_CertGroupFromVerifiedBundle is available on this system.
func CanCallCSSM_CL_CertGroupFromVerifiedBundle() bool {
	return _cSSM_CL_CertGroupFromVerifiedBundle != nil
}

// CSSM_CL_CertGroupFromVerifiedBundleCallError returns the symbol lookup or registration error for CSSM_CL_CertGroupFromVerifiedBundle.
func CSSM_CL_CertGroupFromVerifiedBundleCallError() error {
	if _cSSM_CL_CertGroupFromVerifiedBundle != nil {
		return nil
	}
	return symbolCallError("CSSM_CL_CertGroupFromVerifiedBundle", "10.0", _cSSM_CL_CertGroupFromVerifiedBundleErr)
}

// TryCSSM_CL_CertGroupFromVerifiedBundle calls CSSM_CL_CertGroupFromVerifiedBundle without panicking when the symbol is unavailable.
func TryCSSM_CL_CertGroupFromVerifiedBundle(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertBundle unsafe.Pointer, SignerCert unsafe.Pointer, CertGroup unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CL_CertGroupFromVerifiedBundleCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CL_CertGroupFromVerifiedBundle(CLHandle, CCHandle, CertBundle, SignerCert, CertGroup), nil
}

// CSSM_CL_CertGroupFromVerifiedBundle.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertGroupFromVerifiedBundle
func CSSM_CL_CertGroupFromVerifiedBundle(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertBundle unsafe.Pointer, SignerCert unsafe.Pointer, CertGroup unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CL_CertGroupFromVerifiedBundle(CLHandle, CCHandle, CertBundle, SignerCert, CertGroup)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CertGroupToSignedBundle func(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertGroupToBundle unsafe.Pointer, BundleInfo unsafe.Pointer, SignedBundle unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CertGroupToSignedBundleErr error

// CanCallCSSM_CL_CertGroupToSignedBundle reports whether the symbol for CSSM_CL_CertGroupToSignedBundle is available on this system.
func CanCallCSSM_CL_CertGroupToSignedBundle() bool {
	return _cSSM_CL_CertGroupToSignedBundle != nil
}

// CSSM_CL_CertGroupToSignedBundleCallError returns the symbol lookup or registration error for CSSM_CL_CertGroupToSignedBundle.
func CSSM_CL_CertGroupToSignedBundleCallError() error {
	if _cSSM_CL_CertGroupToSignedBundle != nil {
		return nil
	}
	return symbolCallError("CSSM_CL_CertGroupToSignedBundle", "10.0", _cSSM_CL_CertGroupToSignedBundleErr)
}

// TryCSSM_CL_CertGroupToSignedBundle calls CSSM_CL_CertGroupToSignedBundle without panicking when the symbol is unavailable.
func TryCSSM_CL_CertGroupToSignedBundle(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertGroupToBundle unsafe.Pointer, BundleInfo unsafe.Pointer, SignedBundle unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CL_CertGroupToSignedBundleCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CL_CertGroupToSignedBundle(CLHandle, CCHandle, CertGroupToBundle, BundleInfo, SignedBundle), nil
}

// CSSM_CL_CertGroupToSignedBundle.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertGroupToSignedBundle
func CSSM_CL_CertGroupToSignedBundle(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertGroupToBundle unsafe.Pointer, BundleInfo unsafe.Pointer, SignedBundle unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CL_CertGroupToSignedBundle(CLHandle, CCHandle, CertGroupToBundle, BundleInfo, SignedBundle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CertSign func(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertTemplate unsafe.Pointer, SignScope unsafe.Pointer, ScopeSize uint32, SignedCert unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CertSignErr error

// CanCallCSSM_CL_CertSign reports whether the symbol for CSSM_CL_CertSign is available on this system.
func CanCallCSSM_CL_CertSign() bool {
	return _cSSM_CL_CertSign != nil
}

// CSSM_CL_CertSignCallError returns the symbol lookup or registration error for CSSM_CL_CertSign.
func CSSM_CL_CertSignCallError() error {
	if _cSSM_CL_CertSign != nil {
		return nil
	}
	return symbolCallError("CSSM_CL_CertSign", "10.0", _cSSM_CL_CertSignErr)
}

// TryCSSM_CL_CertSign calls CSSM_CL_CertSign without panicking when the symbol is unavailable.
func TryCSSM_CL_CertSign(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertTemplate unsafe.Pointer, SignScope unsafe.Pointer, ScopeSize uint32, SignedCert unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CL_CertSignCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CL_CertSign(CLHandle, CCHandle, CertTemplate, SignScope, ScopeSize, SignedCert), nil
}

// CSSM_CL_CertSign.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertSign
func CSSM_CL_CertSign(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertTemplate unsafe.Pointer, SignScope unsafe.Pointer, ScopeSize uint32, SignedCert unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CL_CertSign(CLHandle, CCHandle, CertTemplate, SignScope, ScopeSize, SignedCert)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CertVerify func(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertToBeVerified unsafe.Pointer, SignerCert unsafe.Pointer, VerifyScope unsafe.Pointer, ScopeSize uint32) CSSM_RETURN
var _cSSM_CL_CertVerifyErr error

// CanCallCSSM_CL_CertVerify reports whether the symbol for CSSM_CL_CertVerify is available on this system.
func CanCallCSSM_CL_CertVerify() bool {
	return _cSSM_CL_CertVerify != nil
}

// CSSM_CL_CertVerifyCallError returns the symbol lookup or registration error for CSSM_CL_CertVerify.
func CSSM_CL_CertVerifyCallError() error {
	if _cSSM_CL_CertVerify != nil {
		return nil
	}
	return symbolCallError("CSSM_CL_CertVerify", "10.0", _cSSM_CL_CertVerifyErr)
}

// TryCSSM_CL_CertVerify calls CSSM_CL_CertVerify without panicking when the symbol is unavailable.
func TryCSSM_CL_CertVerify(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertToBeVerified unsafe.Pointer, SignerCert unsafe.Pointer, VerifyScope unsafe.Pointer, ScopeSize uint32) (CSSM_RETURN, error) {
	if err := CSSM_CL_CertVerifyCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CL_CertVerify(CLHandle, CCHandle, CertToBeVerified, SignerCert, VerifyScope, ScopeSize), nil
}

// CSSM_CL_CertVerify.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertVerify
func CSSM_CL_CertVerify(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertToBeVerified unsafe.Pointer, SignerCert unsafe.Pointer, VerifyScope unsafe.Pointer, ScopeSize uint32) CSSM_RETURN {
	result, callErr := TryCSSM_CL_CertVerify(CLHandle, CCHandle, CertToBeVerified, SignerCert, VerifyScope, ScopeSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CertVerifyWithKey func(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertToBeVerified unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CertVerifyWithKeyErr error

// CanCallCSSM_CL_CertVerifyWithKey reports whether the symbol for CSSM_CL_CertVerifyWithKey is available on this system.
func CanCallCSSM_CL_CertVerifyWithKey() bool {
	return _cSSM_CL_CertVerifyWithKey != nil
}

// CSSM_CL_CertVerifyWithKeyCallError returns the symbol lookup or registration error for CSSM_CL_CertVerifyWithKey.
func CSSM_CL_CertVerifyWithKeyCallError() error {
	if _cSSM_CL_CertVerifyWithKey != nil {
		return nil
	}
	return symbolCallError("CSSM_CL_CertVerifyWithKey", "10.0", _cSSM_CL_CertVerifyWithKeyErr)
}

// TryCSSM_CL_CertVerifyWithKey calls CSSM_CL_CertVerifyWithKey without panicking when the symbol is unavailable.
func TryCSSM_CL_CertVerifyWithKey(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertToBeVerified unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CL_CertVerifyWithKeyCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CL_CertVerifyWithKey(CLHandle, CCHandle, CertToBeVerified), nil
}

// CSSM_CL_CertVerifyWithKey.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertVerifyWithKey
func CSSM_CL_CertVerifyWithKey(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertToBeVerified unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CL_CertVerifyWithKey(CLHandle, CCHandle, CertToBeVerified)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CrlAbortCache func(CLHandle CSSM_CL_HANDLE, CrlHandle CSSM_HANDLE) CSSM_RETURN
var _cSSM_CL_CrlAbortCacheErr error

// CanCallCSSM_CL_CrlAbortCache reports whether the symbol for CSSM_CL_CrlAbortCache is available on this system.
func CanCallCSSM_CL_CrlAbortCache() bool {
	return _cSSM_CL_CrlAbortCache != nil
}

// CSSM_CL_CrlAbortCacheCallError returns the symbol lookup or registration error for CSSM_CL_CrlAbortCache.
func CSSM_CL_CrlAbortCacheCallError() error {
	if _cSSM_CL_CrlAbortCache != nil {
		return nil
	}
	return symbolCallError("CSSM_CL_CrlAbortCache", "10.0", _cSSM_CL_CrlAbortCacheErr)
}

// TryCSSM_CL_CrlAbortCache calls CSSM_CL_CrlAbortCache without panicking when the symbol is unavailable.
func TryCSSM_CL_CrlAbortCache(CLHandle CSSM_CL_HANDLE, CrlHandle CSSM_HANDLE) (CSSM_RETURN, error) {
	if err := CSSM_CL_CrlAbortCacheCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CL_CrlAbortCache(CLHandle, CrlHandle), nil
}

// CSSM_CL_CrlAbortCache.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlAbortCache
func CSSM_CL_CrlAbortCache(CLHandle CSSM_CL_HANDLE, CrlHandle CSSM_HANDLE) CSSM_RETURN {
	result, callErr := TryCSSM_CL_CrlAbortCache(CLHandle, CrlHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CrlAbortQuery func(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE) CSSM_RETURN
var _cSSM_CL_CrlAbortQueryErr error

// CanCallCSSM_CL_CrlAbortQuery reports whether the symbol for CSSM_CL_CrlAbortQuery is available on this system.
func CanCallCSSM_CL_CrlAbortQuery() bool {
	return _cSSM_CL_CrlAbortQuery != nil
}

// CSSM_CL_CrlAbortQueryCallError returns the symbol lookup or registration error for CSSM_CL_CrlAbortQuery.
func CSSM_CL_CrlAbortQueryCallError() error {
	if _cSSM_CL_CrlAbortQuery != nil {
		return nil
	}
	return symbolCallError("CSSM_CL_CrlAbortQuery", "10.0", _cSSM_CL_CrlAbortQueryErr)
}

// TryCSSM_CL_CrlAbortQuery calls CSSM_CL_CrlAbortQuery without panicking when the symbol is unavailable.
func TryCSSM_CL_CrlAbortQuery(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE) (CSSM_RETURN, error) {
	if err := CSSM_CL_CrlAbortQueryCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CL_CrlAbortQuery(CLHandle, ResultsHandle), nil
}

// CSSM_CL_CrlAbortQuery.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlAbortQuery
func CSSM_CL_CrlAbortQuery(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE) CSSM_RETURN {
	result, callErr := TryCSSM_CL_CrlAbortQuery(CLHandle, ResultsHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CrlAddCert func(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, Cert unsafe.Pointer, NumberOfFields uint32, CrlEntryFields unsafe.Pointer, OldCrl unsafe.Pointer, NewCrl unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CrlAddCertErr error

// CanCallCSSM_CL_CrlAddCert reports whether the symbol for CSSM_CL_CrlAddCert is available on this system.
func CanCallCSSM_CL_CrlAddCert() bool {
	return _cSSM_CL_CrlAddCert != nil
}

// CSSM_CL_CrlAddCertCallError returns the symbol lookup or registration error for CSSM_CL_CrlAddCert.
func CSSM_CL_CrlAddCertCallError() error {
	if _cSSM_CL_CrlAddCert != nil {
		return nil
	}
	return symbolCallError("CSSM_CL_CrlAddCert", "10.0", _cSSM_CL_CrlAddCertErr)
}

// TryCSSM_CL_CrlAddCert calls CSSM_CL_CrlAddCert without panicking when the symbol is unavailable.
func TryCSSM_CL_CrlAddCert(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, Cert unsafe.Pointer, NumberOfFields uint32, CrlEntryFields unsafe.Pointer, OldCrl unsafe.Pointer, NewCrl unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CL_CrlAddCertCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CL_CrlAddCert(CLHandle, CCHandle, Cert, NumberOfFields, CrlEntryFields, OldCrl, NewCrl), nil
}

// CSSM_CL_CrlAddCert.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlAddCert
func CSSM_CL_CrlAddCert(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, Cert unsafe.Pointer, NumberOfFields uint32, CrlEntryFields unsafe.Pointer, OldCrl unsafe.Pointer, NewCrl unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CL_CrlAddCert(CLHandle, CCHandle, Cert, NumberOfFields, CrlEntryFields, OldCrl, NewCrl)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CrlCache func(CLHandle CSSM_CL_HANDLE, Crl unsafe.Pointer, CrlHandle CSSM_HANDLE_PTR) CSSM_RETURN
var _cSSM_CL_CrlCacheErr error

// CanCallCSSM_CL_CrlCache reports whether the symbol for CSSM_CL_CrlCache is available on this system.
func CanCallCSSM_CL_CrlCache() bool {
	return _cSSM_CL_CrlCache != nil
}

// CSSM_CL_CrlCacheCallError returns the symbol lookup or registration error for CSSM_CL_CrlCache.
func CSSM_CL_CrlCacheCallError() error {
	if _cSSM_CL_CrlCache != nil {
		return nil
	}
	return symbolCallError("CSSM_CL_CrlCache", "10.0", _cSSM_CL_CrlCacheErr)
}

// TryCSSM_CL_CrlCache calls CSSM_CL_CrlCache without panicking when the symbol is unavailable.
func TryCSSM_CL_CrlCache(CLHandle CSSM_CL_HANDLE, Crl unsafe.Pointer, CrlHandle CSSM_HANDLE_PTR) (CSSM_RETURN, error) {
	if err := CSSM_CL_CrlCacheCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CL_CrlCache(CLHandle, Crl, CrlHandle), nil
}

// CSSM_CL_CrlCache.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlCache
func CSSM_CL_CrlCache(CLHandle CSSM_CL_HANDLE, Crl unsafe.Pointer, CrlHandle CSSM_HANDLE_PTR) CSSM_RETURN {
	result, callErr := TryCSSM_CL_CrlCache(CLHandle, Crl, CrlHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CrlCreateTemplate func(CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CrlTemplate unsafe.Pointer, NewCrl unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CrlCreateTemplateErr error

// CanCallCSSM_CL_CrlCreateTemplate reports whether the symbol for CSSM_CL_CrlCreateTemplate is available on this system.
func CanCallCSSM_CL_CrlCreateTemplate() bool {
	return _cSSM_CL_CrlCreateTemplate != nil
}

// CSSM_CL_CrlCreateTemplateCallError returns the symbol lookup or registration error for CSSM_CL_CrlCreateTemplate.
func CSSM_CL_CrlCreateTemplateCallError() error {
	if _cSSM_CL_CrlCreateTemplate != nil {
		return nil
	}
	return symbolCallError("CSSM_CL_CrlCreateTemplate", "10.0", _cSSM_CL_CrlCreateTemplateErr)
}

// TryCSSM_CL_CrlCreateTemplate calls CSSM_CL_CrlCreateTemplate without panicking when the symbol is unavailable.
func TryCSSM_CL_CrlCreateTemplate(CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CrlTemplate unsafe.Pointer, NewCrl unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CL_CrlCreateTemplateCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CL_CrlCreateTemplate(CLHandle, NumberOfFields, CrlTemplate, NewCrl), nil
}

// CSSM_CL_CrlCreateTemplate.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlCreateTemplate
func CSSM_CL_CrlCreateTemplate(CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CrlTemplate unsafe.Pointer, NewCrl unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CL_CrlCreateTemplate(CLHandle, NumberOfFields, CrlTemplate, NewCrl)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CrlDescribeFormat func(CLHandle CSSM_CL_HANDLE, NumberOfFields *uint32, OidList unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CrlDescribeFormatErr error

// CanCallCSSM_CL_CrlDescribeFormat reports whether the symbol for CSSM_CL_CrlDescribeFormat is available on this system.
func CanCallCSSM_CL_CrlDescribeFormat() bool {
	return _cSSM_CL_CrlDescribeFormat != nil
}

// CSSM_CL_CrlDescribeFormatCallError returns the symbol lookup or registration error for CSSM_CL_CrlDescribeFormat.
func CSSM_CL_CrlDescribeFormatCallError() error {
	if _cSSM_CL_CrlDescribeFormat != nil {
		return nil
	}
	return symbolCallError("CSSM_CL_CrlDescribeFormat", "10.0", _cSSM_CL_CrlDescribeFormatErr)
}

// TryCSSM_CL_CrlDescribeFormat calls CSSM_CL_CrlDescribeFormat without panicking when the symbol is unavailable.
func TryCSSM_CL_CrlDescribeFormat(CLHandle CSSM_CL_HANDLE, NumberOfFields *uint32, OidList unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CL_CrlDescribeFormatCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CL_CrlDescribeFormat(CLHandle, NumberOfFields, OidList), nil
}

// CSSM_CL_CrlDescribeFormat.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlDescribeFormat
func CSSM_CL_CrlDescribeFormat(CLHandle CSSM_CL_HANDLE, NumberOfFields *uint32, OidList unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CL_CrlDescribeFormat(CLHandle, NumberOfFields, OidList)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CrlGetAllCachedRecordFields func(CLHandle CSSM_CL_HANDLE, CrlHandle CSSM_HANDLE, CrlRecordIndex unsafe.Pointer, NumberOfFields *uint32, CrlFields unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CrlGetAllCachedRecordFieldsErr error

// CanCallCSSM_CL_CrlGetAllCachedRecordFields reports whether the symbol for CSSM_CL_CrlGetAllCachedRecordFields is available on this system.
func CanCallCSSM_CL_CrlGetAllCachedRecordFields() bool {
	return _cSSM_CL_CrlGetAllCachedRecordFields != nil
}

// CSSM_CL_CrlGetAllCachedRecordFieldsCallError returns the symbol lookup or registration error for CSSM_CL_CrlGetAllCachedRecordFields.
func CSSM_CL_CrlGetAllCachedRecordFieldsCallError() error {
	if _cSSM_CL_CrlGetAllCachedRecordFields != nil {
		return nil
	}
	return symbolCallError("CSSM_CL_CrlGetAllCachedRecordFields", "10.0", _cSSM_CL_CrlGetAllCachedRecordFieldsErr)
}

// TryCSSM_CL_CrlGetAllCachedRecordFields calls CSSM_CL_CrlGetAllCachedRecordFields without panicking when the symbol is unavailable.
func TryCSSM_CL_CrlGetAllCachedRecordFields(CLHandle CSSM_CL_HANDLE, CrlHandle CSSM_HANDLE, CrlRecordIndex unsafe.Pointer, NumberOfFields *uint32, CrlFields unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CL_CrlGetAllCachedRecordFieldsCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CL_CrlGetAllCachedRecordFields(CLHandle, CrlHandle, CrlRecordIndex, NumberOfFields, CrlFields), nil
}

// CSSM_CL_CrlGetAllCachedRecordFields.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlGetAllCachedRecordFields
func CSSM_CL_CrlGetAllCachedRecordFields(CLHandle CSSM_CL_HANDLE, CrlHandle CSSM_HANDLE, CrlRecordIndex unsafe.Pointer, NumberOfFields *uint32, CrlFields unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CL_CrlGetAllCachedRecordFields(CLHandle, CrlHandle, CrlRecordIndex, NumberOfFields, CrlFields)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CrlGetAllFields func(CLHandle CSSM_CL_HANDLE, Crl unsafe.Pointer, NumberOfCrlFields *uint32, CrlFields unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CrlGetAllFieldsErr error

// CanCallCSSM_CL_CrlGetAllFields reports whether the symbol for CSSM_CL_CrlGetAllFields is available on this system.
func CanCallCSSM_CL_CrlGetAllFields() bool {
	return _cSSM_CL_CrlGetAllFields != nil
}

// CSSM_CL_CrlGetAllFieldsCallError returns the symbol lookup or registration error for CSSM_CL_CrlGetAllFields.
func CSSM_CL_CrlGetAllFieldsCallError() error {
	if _cSSM_CL_CrlGetAllFields != nil {
		return nil
	}
	return symbolCallError("CSSM_CL_CrlGetAllFields", "10.0", _cSSM_CL_CrlGetAllFieldsErr)
}

// TryCSSM_CL_CrlGetAllFields calls CSSM_CL_CrlGetAllFields without panicking when the symbol is unavailable.
func TryCSSM_CL_CrlGetAllFields(CLHandle CSSM_CL_HANDLE, Crl unsafe.Pointer, NumberOfCrlFields *uint32, CrlFields unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CL_CrlGetAllFieldsCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CL_CrlGetAllFields(CLHandle, Crl, NumberOfCrlFields, CrlFields), nil
}

// CSSM_CL_CrlGetAllFields.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlGetAllFields
func CSSM_CL_CrlGetAllFields(CLHandle CSSM_CL_HANDLE, Crl unsafe.Pointer, NumberOfCrlFields *uint32, CrlFields unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CL_CrlGetAllFields(CLHandle, Crl, NumberOfCrlFields, CrlFields)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CrlGetFirstCachedFieldValue func(CLHandle CSSM_CL_HANDLE, CrlHandle CSSM_HANDLE, CrlRecordIndex unsafe.Pointer, CrlField unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, NumberOfMatchedFields *uint32, Value unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CrlGetFirstCachedFieldValueErr error

// CanCallCSSM_CL_CrlGetFirstCachedFieldValue reports whether the symbol for CSSM_CL_CrlGetFirstCachedFieldValue is available on this system.
func CanCallCSSM_CL_CrlGetFirstCachedFieldValue() bool {
	return _cSSM_CL_CrlGetFirstCachedFieldValue != nil
}

// CSSM_CL_CrlGetFirstCachedFieldValueCallError returns the symbol lookup or registration error for CSSM_CL_CrlGetFirstCachedFieldValue.
func CSSM_CL_CrlGetFirstCachedFieldValueCallError() error {
	if _cSSM_CL_CrlGetFirstCachedFieldValue != nil {
		return nil
	}
	return symbolCallError("CSSM_CL_CrlGetFirstCachedFieldValue", "10.0", _cSSM_CL_CrlGetFirstCachedFieldValueErr)
}

// TryCSSM_CL_CrlGetFirstCachedFieldValue calls CSSM_CL_CrlGetFirstCachedFieldValue without panicking when the symbol is unavailable.
func TryCSSM_CL_CrlGetFirstCachedFieldValue(CLHandle CSSM_CL_HANDLE, CrlHandle CSSM_HANDLE, CrlRecordIndex unsafe.Pointer, CrlField unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, NumberOfMatchedFields *uint32, Value unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CL_CrlGetFirstCachedFieldValueCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CL_CrlGetFirstCachedFieldValue(CLHandle, CrlHandle, CrlRecordIndex, CrlField, ResultsHandle, NumberOfMatchedFields, Value), nil
}

// CSSM_CL_CrlGetFirstCachedFieldValue.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlGetFirstCachedFieldValue
func CSSM_CL_CrlGetFirstCachedFieldValue(CLHandle CSSM_CL_HANDLE, CrlHandle CSSM_HANDLE, CrlRecordIndex unsafe.Pointer, CrlField unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, NumberOfMatchedFields *uint32, Value unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CL_CrlGetFirstCachedFieldValue(CLHandle, CrlHandle, CrlRecordIndex, CrlField, ResultsHandle, NumberOfMatchedFields, Value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CrlGetFirstFieldValue func(CLHandle CSSM_CL_HANDLE, Crl unsafe.Pointer, CrlField unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, NumberOfMatchedFields *uint32, Value unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CrlGetFirstFieldValueErr error

// CanCallCSSM_CL_CrlGetFirstFieldValue reports whether the symbol for CSSM_CL_CrlGetFirstFieldValue is available on this system.
func CanCallCSSM_CL_CrlGetFirstFieldValue() bool {
	return _cSSM_CL_CrlGetFirstFieldValue != nil
}

// CSSM_CL_CrlGetFirstFieldValueCallError returns the symbol lookup or registration error for CSSM_CL_CrlGetFirstFieldValue.
func CSSM_CL_CrlGetFirstFieldValueCallError() error {
	if _cSSM_CL_CrlGetFirstFieldValue != nil {
		return nil
	}
	return symbolCallError("CSSM_CL_CrlGetFirstFieldValue", "10.0", _cSSM_CL_CrlGetFirstFieldValueErr)
}

// TryCSSM_CL_CrlGetFirstFieldValue calls CSSM_CL_CrlGetFirstFieldValue without panicking when the symbol is unavailable.
func TryCSSM_CL_CrlGetFirstFieldValue(CLHandle CSSM_CL_HANDLE, Crl unsafe.Pointer, CrlField unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, NumberOfMatchedFields *uint32, Value unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CL_CrlGetFirstFieldValueCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CL_CrlGetFirstFieldValue(CLHandle, Crl, CrlField, ResultsHandle, NumberOfMatchedFields, Value), nil
}

// CSSM_CL_CrlGetFirstFieldValue.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlGetFirstFieldValue
func CSSM_CL_CrlGetFirstFieldValue(CLHandle CSSM_CL_HANDLE, Crl unsafe.Pointer, CrlField unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, NumberOfMatchedFields *uint32, Value unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CL_CrlGetFirstFieldValue(CLHandle, Crl, CrlField, ResultsHandle, NumberOfMatchedFields, Value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CrlGetNextCachedFieldValue func(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE, Value unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CrlGetNextCachedFieldValueErr error

// CanCallCSSM_CL_CrlGetNextCachedFieldValue reports whether the symbol for CSSM_CL_CrlGetNextCachedFieldValue is available on this system.
func CanCallCSSM_CL_CrlGetNextCachedFieldValue() bool {
	return _cSSM_CL_CrlGetNextCachedFieldValue != nil
}

// CSSM_CL_CrlGetNextCachedFieldValueCallError returns the symbol lookup or registration error for CSSM_CL_CrlGetNextCachedFieldValue.
func CSSM_CL_CrlGetNextCachedFieldValueCallError() error {
	if _cSSM_CL_CrlGetNextCachedFieldValue != nil {
		return nil
	}
	return symbolCallError("CSSM_CL_CrlGetNextCachedFieldValue", "10.0", _cSSM_CL_CrlGetNextCachedFieldValueErr)
}

// TryCSSM_CL_CrlGetNextCachedFieldValue calls CSSM_CL_CrlGetNextCachedFieldValue without panicking when the symbol is unavailable.
func TryCSSM_CL_CrlGetNextCachedFieldValue(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE, Value unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CL_CrlGetNextCachedFieldValueCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CL_CrlGetNextCachedFieldValue(CLHandle, ResultsHandle, Value), nil
}

// CSSM_CL_CrlGetNextCachedFieldValue.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlGetNextCachedFieldValue
func CSSM_CL_CrlGetNextCachedFieldValue(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE, Value unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CL_CrlGetNextCachedFieldValue(CLHandle, ResultsHandle, Value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CrlGetNextFieldValue func(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE, Value unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CrlGetNextFieldValueErr error

// CanCallCSSM_CL_CrlGetNextFieldValue reports whether the symbol for CSSM_CL_CrlGetNextFieldValue is available on this system.
func CanCallCSSM_CL_CrlGetNextFieldValue() bool {
	return _cSSM_CL_CrlGetNextFieldValue != nil
}

// CSSM_CL_CrlGetNextFieldValueCallError returns the symbol lookup or registration error for CSSM_CL_CrlGetNextFieldValue.
func CSSM_CL_CrlGetNextFieldValueCallError() error {
	if _cSSM_CL_CrlGetNextFieldValue != nil {
		return nil
	}
	return symbolCallError("CSSM_CL_CrlGetNextFieldValue", "10.0", _cSSM_CL_CrlGetNextFieldValueErr)
}

// TryCSSM_CL_CrlGetNextFieldValue calls CSSM_CL_CrlGetNextFieldValue without panicking when the symbol is unavailable.
func TryCSSM_CL_CrlGetNextFieldValue(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE, Value unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CL_CrlGetNextFieldValueCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CL_CrlGetNextFieldValue(CLHandle, ResultsHandle, Value), nil
}

// CSSM_CL_CrlGetNextFieldValue.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlGetNextFieldValue
func CSSM_CL_CrlGetNextFieldValue(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE, Value unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CL_CrlGetNextFieldValue(CLHandle, ResultsHandle, Value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CrlRemoveCert func(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, OldCrl unsafe.Pointer, NewCrl unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CrlRemoveCertErr error

// CanCallCSSM_CL_CrlRemoveCert reports whether the symbol for CSSM_CL_CrlRemoveCert is available on this system.
func CanCallCSSM_CL_CrlRemoveCert() bool {
	return _cSSM_CL_CrlRemoveCert != nil
}

// CSSM_CL_CrlRemoveCertCallError returns the symbol lookup or registration error for CSSM_CL_CrlRemoveCert.
func CSSM_CL_CrlRemoveCertCallError() error {
	if _cSSM_CL_CrlRemoveCert != nil {
		return nil
	}
	return symbolCallError("CSSM_CL_CrlRemoveCert", "10.0", _cSSM_CL_CrlRemoveCertErr)
}

// TryCSSM_CL_CrlRemoveCert calls CSSM_CL_CrlRemoveCert without panicking when the symbol is unavailable.
func TryCSSM_CL_CrlRemoveCert(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, OldCrl unsafe.Pointer, NewCrl unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CL_CrlRemoveCertCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CL_CrlRemoveCert(CLHandle, Cert, OldCrl, NewCrl), nil
}

// CSSM_CL_CrlRemoveCert.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlRemoveCert
func CSSM_CL_CrlRemoveCert(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, OldCrl unsafe.Pointer, NewCrl unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CL_CrlRemoveCert(CLHandle, Cert, OldCrl, NewCrl)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CrlSetFields func(CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CrlTemplate unsafe.Pointer, OldCrl unsafe.Pointer, ModifiedCrl unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CrlSetFieldsErr error

// CanCallCSSM_CL_CrlSetFields reports whether the symbol for CSSM_CL_CrlSetFields is available on this system.
func CanCallCSSM_CL_CrlSetFields() bool {
	return _cSSM_CL_CrlSetFields != nil
}

// CSSM_CL_CrlSetFieldsCallError returns the symbol lookup or registration error for CSSM_CL_CrlSetFields.
func CSSM_CL_CrlSetFieldsCallError() error {
	if _cSSM_CL_CrlSetFields != nil {
		return nil
	}
	return symbolCallError("CSSM_CL_CrlSetFields", "10.0", _cSSM_CL_CrlSetFieldsErr)
}

// TryCSSM_CL_CrlSetFields calls CSSM_CL_CrlSetFields without panicking when the symbol is unavailable.
func TryCSSM_CL_CrlSetFields(CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CrlTemplate unsafe.Pointer, OldCrl unsafe.Pointer, ModifiedCrl unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CL_CrlSetFieldsCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CL_CrlSetFields(CLHandle, NumberOfFields, CrlTemplate, OldCrl, ModifiedCrl), nil
}

// CSSM_CL_CrlSetFields.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlSetFields
func CSSM_CL_CrlSetFields(CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CrlTemplate unsafe.Pointer, OldCrl unsafe.Pointer, ModifiedCrl unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CL_CrlSetFields(CLHandle, NumberOfFields, CrlTemplate, OldCrl, ModifiedCrl)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CrlSign func(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, UnsignedCrl unsafe.Pointer, SignScope unsafe.Pointer, ScopeSize uint32, SignedCrl unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CrlSignErr error

// CanCallCSSM_CL_CrlSign reports whether the symbol for CSSM_CL_CrlSign is available on this system.
func CanCallCSSM_CL_CrlSign() bool {
	return _cSSM_CL_CrlSign != nil
}

// CSSM_CL_CrlSignCallError returns the symbol lookup or registration error for CSSM_CL_CrlSign.
func CSSM_CL_CrlSignCallError() error {
	if _cSSM_CL_CrlSign != nil {
		return nil
	}
	return symbolCallError("CSSM_CL_CrlSign", "10.0", _cSSM_CL_CrlSignErr)
}

// TryCSSM_CL_CrlSign calls CSSM_CL_CrlSign without panicking when the symbol is unavailable.
func TryCSSM_CL_CrlSign(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, UnsignedCrl unsafe.Pointer, SignScope unsafe.Pointer, ScopeSize uint32, SignedCrl unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CL_CrlSignCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CL_CrlSign(CLHandle, CCHandle, UnsignedCrl, SignScope, ScopeSize, SignedCrl), nil
}

// CSSM_CL_CrlSign.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlSign
func CSSM_CL_CrlSign(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, UnsignedCrl unsafe.Pointer, SignScope unsafe.Pointer, ScopeSize uint32, SignedCrl unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CL_CrlSign(CLHandle, CCHandle, UnsignedCrl, SignScope, ScopeSize, SignedCrl)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CrlVerify func(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CrlToBeVerified unsafe.Pointer, SignerCert unsafe.Pointer, VerifyScope unsafe.Pointer, ScopeSize uint32) CSSM_RETURN
var _cSSM_CL_CrlVerifyErr error

// CanCallCSSM_CL_CrlVerify reports whether the symbol for CSSM_CL_CrlVerify is available on this system.
func CanCallCSSM_CL_CrlVerify() bool {
	return _cSSM_CL_CrlVerify != nil
}

// CSSM_CL_CrlVerifyCallError returns the symbol lookup or registration error for CSSM_CL_CrlVerify.
func CSSM_CL_CrlVerifyCallError() error {
	if _cSSM_CL_CrlVerify != nil {
		return nil
	}
	return symbolCallError("CSSM_CL_CrlVerify", "10.0", _cSSM_CL_CrlVerifyErr)
}

// TryCSSM_CL_CrlVerify calls CSSM_CL_CrlVerify without panicking when the symbol is unavailable.
func TryCSSM_CL_CrlVerify(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CrlToBeVerified unsafe.Pointer, SignerCert unsafe.Pointer, VerifyScope unsafe.Pointer, ScopeSize uint32) (CSSM_RETURN, error) {
	if err := CSSM_CL_CrlVerifyCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CL_CrlVerify(CLHandle, CCHandle, CrlToBeVerified, SignerCert, VerifyScope, ScopeSize), nil
}

// CSSM_CL_CrlVerify.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlVerify
func CSSM_CL_CrlVerify(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CrlToBeVerified unsafe.Pointer, SignerCert unsafe.Pointer, VerifyScope unsafe.Pointer, ScopeSize uint32) CSSM_RETURN {
	result, callErr := TryCSSM_CL_CrlVerify(CLHandle, CCHandle, CrlToBeVerified, SignerCert, VerifyScope, ScopeSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_CrlVerifyWithKey func(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CrlToBeVerified unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_CrlVerifyWithKeyErr error

// CanCallCSSM_CL_CrlVerifyWithKey reports whether the symbol for CSSM_CL_CrlVerifyWithKey is available on this system.
func CanCallCSSM_CL_CrlVerifyWithKey() bool {
	return _cSSM_CL_CrlVerifyWithKey != nil
}

// CSSM_CL_CrlVerifyWithKeyCallError returns the symbol lookup or registration error for CSSM_CL_CrlVerifyWithKey.
func CSSM_CL_CrlVerifyWithKeyCallError() error {
	if _cSSM_CL_CrlVerifyWithKey != nil {
		return nil
	}
	return symbolCallError("CSSM_CL_CrlVerifyWithKey", "10.0", _cSSM_CL_CrlVerifyWithKeyErr)
}

// TryCSSM_CL_CrlVerifyWithKey calls CSSM_CL_CrlVerifyWithKey without panicking when the symbol is unavailable.
func TryCSSM_CL_CrlVerifyWithKey(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CrlToBeVerified unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CL_CrlVerifyWithKeyCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CL_CrlVerifyWithKey(CLHandle, CCHandle, CrlToBeVerified), nil
}

// CSSM_CL_CrlVerifyWithKey.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlVerifyWithKey
func CSSM_CL_CrlVerifyWithKey(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CrlToBeVerified unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CL_CrlVerifyWithKey(CLHandle, CCHandle, CrlToBeVerified)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_FreeFieldValue func(CLHandle CSSM_CL_HANDLE, CertOrCrlOid unsafe.Pointer, Value unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_FreeFieldValueErr error

// CanCallCSSM_CL_FreeFieldValue reports whether the symbol for CSSM_CL_FreeFieldValue is available on this system.
func CanCallCSSM_CL_FreeFieldValue() bool {
	return _cSSM_CL_FreeFieldValue != nil
}

// CSSM_CL_FreeFieldValueCallError returns the symbol lookup or registration error for CSSM_CL_FreeFieldValue.
func CSSM_CL_FreeFieldValueCallError() error {
	if _cSSM_CL_FreeFieldValue != nil {
		return nil
	}
	return symbolCallError("CSSM_CL_FreeFieldValue", "10.0", _cSSM_CL_FreeFieldValueErr)
}

// TryCSSM_CL_FreeFieldValue calls CSSM_CL_FreeFieldValue without panicking when the symbol is unavailable.
func TryCSSM_CL_FreeFieldValue(CLHandle CSSM_CL_HANDLE, CertOrCrlOid unsafe.Pointer, Value unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CL_FreeFieldValueCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CL_FreeFieldValue(CLHandle, CertOrCrlOid, Value), nil
}

// CSSM_CL_FreeFieldValue.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_FreeFieldValue
func CSSM_CL_FreeFieldValue(CLHandle CSSM_CL_HANDLE, CertOrCrlOid unsafe.Pointer, Value unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CL_FreeFieldValue(CLHandle, CertOrCrlOid, Value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_FreeFields func(CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, Fields unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_FreeFieldsErr error

// CanCallCSSM_CL_FreeFields reports whether the symbol for CSSM_CL_FreeFields is available on this system.
func CanCallCSSM_CL_FreeFields() bool {
	return _cSSM_CL_FreeFields != nil
}

// CSSM_CL_FreeFieldsCallError returns the symbol lookup or registration error for CSSM_CL_FreeFields.
func CSSM_CL_FreeFieldsCallError() error {
	if _cSSM_CL_FreeFields != nil {
		return nil
	}
	return symbolCallError("CSSM_CL_FreeFields", "10.0", _cSSM_CL_FreeFieldsErr)
}

// TryCSSM_CL_FreeFields calls CSSM_CL_FreeFields without panicking when the symbol is unavailable.
func TryCSSM_CL_FreeFields(CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, Fields unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CL_FreeFieldsCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CL_FreeFields(CLHandle, NumberOfFields, Fields), nil
}

// CSSM_CL_FreeFields.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_FreeFields
func CSSM_CL_FreeFields(CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, Fields unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CL_FreeFields(CLHandle, NumberOfFields, Fields)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_IsCertInCachedCrl func(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, CrlHandle CSSM_HANDLE, CertFound unsafe.Pointer, CrlRecordIndex unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_IsCertInCachedCrlErr error

// CanCallCSSM_CL_IsCertInCachedCrl reports whether the symbol for CSSM_CL_IsCertInCachedCrl is available on this system.
func CanCallCSSM_CL_IsCertInCachedCrl() bool {
	return _cSSM_CL_IsCertInCachedCrl != nil
}

// CSSM_CL_IsCertInCachedCrlCallError returns the symbol lookup or registration error for CSSM_CL_IsCertInCachedCrl.
func CSSM_CL_IsCertInCachedCrlCallError() error {
	if _cSSM_CL_IsCertInCachedCrl != nil {
		return nil
	}
	return symbolCallError("CSSM_CL_IsCertInCachedCrl", "10.0", _cSSM_CL_IsCertInCachedCrlErr)
}

// TryCSSM_CL_IsCertInCachedCrl calls CSSM_CL_IsCertInCachedCrl without panicking when the symbol is unavailable.
func TryCSSM_CL_IsCertInCachedCrl(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, CrlHandle CSSM_HANDLE, CertFound unsafe.Pointer, CrlRecordIndex unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CL_IsCertInCachedCrlCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CL_IsCertInCachedCrl(CLHandle, Cert, CrlHandle, CertFound, CrlRecordIndex), nil
}

// CSSM_CL_IsCertInCachedCrl.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_IsCertInCachedCrl
func CSSM_CL_IsCertInCachedCrl(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, CrlHandle CSSM_HANDLE, CertFound unsafe.Pointer, CrlRecordIndex unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CL_IsCertInCachedCrl(CLHandle, Cert, CrlHandle, CertFound, CrlRecordIndex)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_IsCertInCrl func(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, Crl unsafe.Pointer, CertFound unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_IsCertInCrlErr error

// CanCallCSSM_CL_IsCertInCrl reports whether the symbol for CSSM_CL_IsCertInCrl is available on this system.
func CanCallCSSM_CL_IsCertInCrl() bool {
	return _cSSM_CL_IsCertInCrl != nil
}

// CSSM_CL_IsCertInCrlCallError returns the symbol lookup or registration error for CSSM_CL_IsCertInCrl.
func CSSM_CL_IsCertInCrlCallError() error {
	if _cSSM_CL_IsCertInCrl != nil {
		return nil
	}
	return symbolCallError("CSSM_CL_IsCertInCrl", "10.0", _cSSM_CL_IsCertInCrlErr)
}

// TryCSSM_CL_IsCertInCrl calls CSSM_CL_IsCertInCrl without panicking when the symbol is unavailable.
func TryCSSM_CL_IsCertInCrl(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, Crl unsafe.Pointer, CertFound unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CL_IsCertInCrlCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CL_IsCertInCrl(CLHandle, Cert, Crl, CertFound), nil
}

// CSSM_CL_IsCertInCrl.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_IsCertInCrl
func CSSM_CL_IsCertInCrl(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, Crl unsafe.Pointer, CertFound unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CL_IsCertInCrl(CLHandle, Cert, Crl, CertFound)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CL_PassThrough func(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, PassThroughId uint32, InputParams unsafe.Pointer, OutputParams unsafe.Pointer) CSSM_RETURN
var _cSSM_CL_PassThroughErr error

// CanCallCSSM_CL_PassThrough reports whether the symbol for CSSM_CL_PassThrough is available on this system.
func CanCallCSSM_CL_PassThrough() bool {
	return _cSSM_CL_PassThrough != nil
}

// CSSM_CL_PassThroughCallError returns the symbol lookup or registration error for CSSM_CL_PassThrough.
func CSSM_CL_PassThroughCallError() error {
	if _cSSM_CL_PassThrough != nil {
		return nil
	}
	return symbolCallError("CSSM_CL_PassThrough", "10.0", _cSSM_CL_PassThroughErr)
}

// TryCSSM_CL_PassThrough calls CSSM_CL_PassThrough without panicking when the symbol is unavailable.
func TryCSSM_CL_PassThrough(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, PassThroughId uint32, InputParams unsafe.Pointer, OutputParams unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CL_PassThroughCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CL_PassThrough(CLHandle, CCHandle, PassThroughId, InputParams, OutputParams), nil
}

// CSSM_CL_PassThrough.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_PassThrough
func CSSM_CL_PassThrough(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, PassThroughId uint32, InputParams unsafe.Pointer, OutputParams unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CL_PassThrough(CLHandle, CCHandle, PassThroughId, InputParams, OutputParams)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CSP_ChangeLoginAcl func(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, AclEdit unsafe.Pointer) CSSM_RETURN
var _cSSM_CSP_ChangeLoginAclErr error

// CanCallCSSM_CSP_ChangeLoginAcl reports whether the symbol for CSSM_CSP_ChangeLoginAcl is available on this system.
func CanCallCSSM_CSP_ChangeLoginAcl() bool {
	return _cSSM_CSP_ChangeLoginAcl != nil
}

// CSSM_CSP_ChangeLoginAclCallError returns the symbol lookup or registration error for CSSM_CSP_ChangeLoginAcl.
func CSSM_CSP_ChangeLoginAclCallError() error {
	if _cSSM_CSP_ChangeLoginAcl != nil {
		return nil
	}
	return symbolCallError("CSSM_CSP_ChangeLoginAcl", "10.0", _cSSM_CSP_ChangeLoginAclErr)
}

// TryCSSM_CSP_ChangeLoginAcl calls CSSM_CSP_ChangeLoginAcl without panicking when the symbol is unavailable.
func TryCSSM_CSP_ChangeLoginAcl(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, AclEdit unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CSP_ChangeLoginAclCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CSP_ChangeLoginAcl(CSPHandle, AccessCred, AclEdit), nil
}

// CSSM_CSP_ChangeLoginAcl.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_ChangeLoginAcl
func CSSM_CSP_ChangeLoginAcl(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, AclEdit unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CSP_ChangeLoginAcl(CSPHandle, AccessCred, AclEdit)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CSP_ChangeLoginOwner func(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, NewOwner unsafe.Pointer) CSSM_RETURN
var _cSSM_CSP_ChangeLoginOwnerErr error

// CanCallCSSM_CSP_ChangeLoginOwner reports whether the symbol for CSSM_CSP_ChangeLoginOwner is available on this system.
func CanCallCSSM_CSP_ChangeLoginOwner() bool {
	return _cSSM_CSP_ChangeLoginOwner != nil
}

// CSSM_CSP_ChangeLoginOwnerCallError returns the symbol lookup or registration error for CSSM_CSP_ChangeLoginOwner.
func CSSM_CSP_ChangeLoginOwnerCallError() error {
	if _cSSM_CSP_ChangeLoginOwner != nil {
		return nil
	}
	return symbolCallError("CSSM_CSP_ChangeLoginOwner", "10.0", _cSSM_CSP_ChangeLoginOwnerErr)
}

// TryCSSM_CSP_ChangeLoginOwner calls CSSM_CSP_ChangeLoginOwner without panicking when the symbol is unavailable.
func TryCSSM_CSP_ChangeLoginOwner(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, NewOwner unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CSP_ChangeLoginOwnerCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CSP_ChangeLoginOwner(CSPHandle, AccessCred, NewOwner), nil
}

// CSSM_CSP_ChangeLoginOwner.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_ChangeLoginOwner
func CSSM_CSP_ChangeLoginOwner(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, NewOwner unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CSP_ChangeLoginOwner(CSPHandle, AccessCred, NewOwner)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CSP_CreateAsymmetricContext func(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, AccessCred unsafe.Pointer, Key unsafe.Pointer, Padding CSSM_PADDING, NewContextHandle unsafe.Pointer) CSSM_RETURN
var _cSSM_CSP_CreateAsymmetricContextErr error

// CanCallCSSM_CSP_CreateAsymmetricContext reports whether the symbol for CSSM_CSP_CreateAsymmetricContext is available on this system.
func CanCallCSSM_CSP_CreateAsymmetricContext() bool {
	return _cSSM_CSP_CreateAsymmetricContext != nil
}

// CSSM_CSP_CreateAsymmetricContextCallError returns the symbol lookup or registration error for CSSM_CSP_CreateAsymmetricContext.
func CSSM_CSP_CreateAsymmetricContextCallError() error {
	if _cSSM_CSP_CreateAsymmetricContext != nil {
		return nil
	}
	return symbolCallError("CSSM_CSP_CreateAsymmetricContext", "10.0", _cSSM_CSP_CreateAsymmetricContextErr)
}

// TryCSSM_CSP_CreateAsymmetricContext calls CSSM_CSP_CreateAsymmetricContext without panicking when the symbol is unavailable.
func TryCSSM_CSP_CreateAsymmetricContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, AccessCred unsafe.Pointer, Key unsafe.Pointer, Padding CSSM_PADDING, NewContextHandle unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CSP_CreateAsymmetricContextCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CSP_CreateAsymmetricContext(CSPHandle, AlgorithmID, AccessCred, Key, Padding, NewContextHandle), nil
}

// CSSM_CSP_CreateAsymmetricContext.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_CreateAsymmetricContext
func CSSM_CSP_CreateAsymmetricContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, AccessCred unsafe.Pointer, Key unsafe.Pointer, Padding CSSM_PADDING, NewContextHandle unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CSP_CreateAsymmetricContext(CSPHandle, AlgorithmID, AccessCred, Key, Padding, NewContextHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CSP_CreateDeriveKeyContext func(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, DeriveKeyType CSSM_KEY_TYPE, DeriveKeyLengthInBits uint32, AccessCred unsafe.Pointer, BaseKey unsafe.Pointer, IterationCount uint32, Salt unsafe.Pointer, Seed unsafe.Pointer, NewContextHandle unsafe.Pointer) CSSM_RETURN
var _cSSM_CSP_CreateDeriveKeyContextErr error

// CanCallCSSM_CSP_CreateDeriveKeyContext reports whether the symbol for CSSM_CSP_CreateDeriveKeyContext is available on this system.
func CanCallCSSM_CSP_CreateDeriveKeyContext() bool {
	return _cSSM_CSP_CreateDeriveKeyContext != nil
}

// CSSM_CSP_CreateDeriveKeyContextCallError returns the symbol lookup or registration error for CSSM_CSP_CreateDeriveKeyContext.
func CSSM_CSP_CreateDeriveKeyContextCallError() error {
	if _cSSM_CSP_CreateDeriveKeyContext != nil {
		return nil
	}
	return symbolCallError("CSSM_CSP_CreateDeriveKeyContext", "10.0", _cSSM_CSP_CreateDeriveKeyContextErr)
}

// TryCSSM_CSP_CreateDeriveKeyContext calls CSSM_CSP_CreateDeriveKeyContext without panicking when the symbol is unavailable.
func TryCSSM_CSP_CreateDeriveKeyContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, DeriveKeyType CSSM_KEY_TYPE, DeriveKeyLengthInBits uint32, AccessCred unsafe.Pointer, BaseKey unsafe.Pointer, IterationCount uint32, Salt unsafe.Pointer, Seed unsafe.Pointer, NewContextHandle unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CSP_CreateDeriveKeyContextCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CSP_CreateDeriveKeyContext(CSPHandle, AlgorithmID, DeriveKeyType, DeriveKeyLengthInBits, AccessCred, BaseKey, IterationCount, Salt, Seed, NewContextHandle), nil
}

// CSSM_CSP_CreateDeriveKeyContext.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_CreateDeriveKeyContext
func CSSM_CSP_CreateDeriveKeyContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, DeriveKeyType CSSM_KEY_TYPE, DeriveKeyLengthInBits uint32, AccessCred unsafe.Pointer, BaseKey unsafe.Pointer, IterationCount uint32, Salt unsafe.Pointer, Seed unsafe.Pointer, NewContextHandle unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CSP_CreateDeriveKeyContext(CSPHandle, AlgorithmID, DeriveKeyType, DeriveKeyLengthInBits, AccessCred, BaseKey, IterationCount, Salt, Seed, NewContextHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CSP_CreateDigestContext func(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, NewContextHandle unsafe.Pointer) CSSM_RETURN
var _cSSM_CSP_CreateDigestContextErr error

// CanCallCSSM_CSP_CreateDigestContext reports whether the symbol for CSSM_CSP_CreateDigestContext is available on this system.
func CanCallCSSM_CSP_CreateDigestContext() bool {
	return _cSSM_CSP_CreateDigestContext != nil
}

// CSSM_CSP_CreateDigestContextCallError returns the symbol lookup or registration error for CSSM_CSP_CreateDigestContext.
func CSSM_CSP_CreateDigestContextCallError() error {
	if _cSSM_CSP_CreateDigestContext != nil {
		return nil
	}
	return symbolCallError("CSSM_CSP_CreateDigestContext", "10.0", _cSSM_CSP_CreateDigestContextErr)
}

// TryCSSM_CSP_CreateDigestContext calls CSSM_CSP_CreateDigestContext without panicking when the symbol is unavailable.
func TryCSSM_CSP_CreateDigestContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, NewContextHandle unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CSP_CreateDigestContextCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CSP_CreateDigestContext(CSPHandle, AlgorithmID, NewContextHandle), nil
}

// CSSM_CSP_CreateDigestContext.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_CreateDigestContext
func CSSM_CSP_CreateDigestContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, NewContextHandle unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CSP_CreateDigestContext(CSPHandle, AlgorithmID, NewContextHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CSP_CreateKeyGenContext func(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, KeySizeInBits uint32, Seed unsafe.Pointer, Salt unsafe.Pointer, StartDate unsafe.Pointer, EndDate unsafe.Pointer, Params unsafe.Pointer, NewContextHandle unsafe.Pointer) CSSM_RETURN
var _cSSM_CSP_CreateKeyGenContextErr error

// CanCallCSSM_CSP_CreateKeyGenContext reports whether the symbol for CSSM_CSP_CreateKeyGenContext is available on this system.
func CanCallCSSM_CSP_CreateKeyGenContext() bool {
	return _cSSM_CSP_CreateKeyGenContext != nil
}

// CSSM_CSP_CreateKeyGenContextCallError returns the symbol lookup or registration error for CSSM_CSP_CreateKeyGenContext.
func CSSM_CSP_CreateKeyGenContextCallError() error {
	if _cSSM_CSP_CreateKeyGenContext != nil {
		return nil
	}
	return symbolCallError("CSSM_CSP_CreateKeyGenContext", "10.0", _cSSM_CSP_CreateKeyGenContextErr)
}

// TryCSSM_CSP_CreateKeyGenContext calls CSSM_CSP_CreateKeyGenContext without panicking when the symbol is unavailable.
func TryCSSM_CSP_CreateKeyGenContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, KeySizeInBits uint32, Seed unsafe.Pointer, Salt unsafe.Pointer, StartDate unsafe.Pointer, EndDate unsafe.Pointer, Params unsafe.Pointer, NewContextHandle unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CSP_CreateKeyGenContextCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CSP_CreateKeyGenContext(CSPHandle, AlgorithmID, KeySizeInBits, Seed, Salt, StartDate, EndDate, Params, NewContextHandle), nil
}

// CSSM_CSP_CreateKeyGenContext.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_CreateKeyGenContext
func CSSM_CSP_CreateKeyGenContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, KeySizeInBits uint32, Seed unsafe.Pointer, Salt unsafe.Pointer, StartDate unsafe.Pointer, EndDate unsafe.Pointer, Params unsafe.Pointer, NewContextHandle unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CSP_CreateKeyGenContext(CSPHandle, AlgorithmID, KeySizeInBits, Seed, Salt, StartDate, EndDate, Params, NewContextHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CSP_CreateMacContext func(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, Key unsafe.Pointer, NewContextHandle unsafe.Pointer) CSSM_RETURN
var _cSSM_CSP_CreateMacContextErr error

// CanCallCSSM_CSP_CreateMacContext reports whether the symbol for CSSM_CSP_CreateMacContext is available on this system.
func CanCallCSSM_CSP_CreateMacContext() bool {
	return _cSSM_CSP_CreateMacContext != nil
}

// CSSM_CSP_CreateMacContextCallError returns the symbol lookup or registration error for CSSM_CSP_CreateMacContext.
func CSSM_CSP_CreateMacContextCallError() error {
	if _cSSM_CSP_CreateMacContext != nil {
		return nil
	}
	return symbolCallError("CSSM_CSP_CreateMacContext", "10.0", _cSSM_CSP_CreateMacContextErr)
}

// TryCSSM_CSP_CreateMacContext calls CSSM_CSP_CreateMacContext without panicking when the symbol is unavailable.
func TryCSSM_CSP_CreateMacContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, Key unsafe.Pointer, NewContextHandle unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CSP_CreateMacContextCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CSP_CreateMacContext(CSPHandle, AlgorithmID, Key, NewContextHandle), nil
}

// CSSM_CSP_CreateMacContext.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_CreateMacContext
func CSSM_CSP_CreateMacContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, Key unsafe.Pointer, NewContextHandle unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CSP_CreateMacContext(CSPHandle, AlgorithmID, Key, NewContextHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CSP_CreatePassThroughContext func(CSPHandle CSSM_CSP_HANDLE, Key unsafe.Pointer, NewContextHandle unsafe.Pointer) CSSM_RETURN
var _cSSM_CSP_CreatePassThroughContextErr error

// CanCallCSSM_CSP_CreatePassThroughContext reports whether the symbol for CSSM_CSP_CreatePassThroughContext is available on this system.
func CanCallCSSM_CSP_CreatePassThroughContext() bool {
	return _cSSM_CSP_CreatePassThroughContext != nil
}

// CSSM_CSP_CreatePassThroughContextCallError returns the symbol lookup or registration error for CSSM_CSP_CreatePassThroughContext.
func CSSM_CSP_CreatePassThroughContextCallError() error {
	if _cSSM_CSP_CreatePassThroughContext != nil {
		return nil
	}
	return symbolCallError("CSSM_CSP_CreatePassThroughContext", "10.0", _cSSM_CSP_CreatePassThroughContextErr)
}

// TryCSSM_CSP_CreatePassThroughContext calls CSSM_CSP_CreatePassThroughContext without panicking when the symbol is unavailable.
func TryCSSM_CSP_CreatePassThroughContext(CSPHandle CSSM_CSP_HANDLE, Key unsafe.Pointer, NewContextHandle unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CSP_CreatePassThroughContextCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CSP_CreatePassThroughContext(CSPHandle, Key, NewContextHandle), nil
}

// CSSM_CSP_CreatePassThroughContext.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_CreatePassThroughContext
func CSSM_CSP_CreatePassThroughContext(CSPHandle CSSM_CSP_HANDLE, Key unsafe.Pointer, NewContextHandle unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CSP_CreatePassThroughContext(CSPHandle, Key, NewContextHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CSP_CreateRandomGenContext func(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, Seed unsafe.Pointer, Length CSSM_SIZE, NewContextHandle unsafe.Pointer) CSSM_RETURN
var _cSSM_CSP_CreateRandomGenContextErr error

// CanCallCSSM_CSP_CreateRandomGenContext reports whether the symbol for CSSM_CSP_CreateRandomGenContext is available on this system.
func CanCallCSSM_CSP_CreateRandomGenContext() bool {
	return _cSSM_CSP_CreateRandomGenContext != nil
}

// CSSM_CSP_CreateRandomGenContextCallError returns the symbol lookup or registration error for CSSM_CSP_CreateRandomGenContext.
func CSSM_CSP_CreateRandomGenContextCallError() error {
	if _cSSM_CSP_CreateRandomGenContext != nil {
		return nil
	}
	return symbolCallError("CSSM_CSP_CreateRandomGenContext", "10.0", _cSSM_CSP_CreateRandomGenContextErr)
}

// TryCSSM_CSP_CreateRandomGenContext calls CSSM_CSP_CreateRandomGenContext without panicking when the symbol is unavailable.
func TryCSSM_CSP_CreateRandomGenContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, Seed unsafe.Pointer, Length CSSM_SIZE, NewContextHandle unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CSP_CreateRandomGenContextCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CSP_CreateRandomGenContext(CSPHandle, AlgorithmID, Seed, Length, NewContextHandle), nil
}

// CSSM_CSP_CreateRandomGenContext.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_CreateRandomGenContext
func CSSM_CSP_CreateRandomGenContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, Seed unsafe.Pointer, Length CSSM_SIZE, NewContextHandle unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CSP_CreateRandomGenContext(CSPHandle, AlgorithmID, Seed, Length, NewContextHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CSP_CreateSignatureContext func(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, AccessCred unsafe.Pointer, Key unsafe.Pointer, NewContextHandle unsafe.Pointer) CSSM_RETURN
var _cSSM_CSP_CreateSignatureContextErr error

// CanCallCSSM_CSP_CreateSignatureContext reports whether the symbol for CSSM_CSP_CreateSignatureContext is available on this system.
func CanCallCSSM_CSP_CreateSignatureContext() bool {
	return _cSSM_CSP_CreateSignatureContext != nil
}

// CSSM_CSP_CreateSignatureContextCallError returns the symbol lookup or registration error for CSSM_CSP_CreateSignatureContext.
func CSSM_CSP_CreateSignatureContextCallError() error {
	if _cSSM_CSP_CreateSignatureContext != nil {
		return nil
	}
	return symbolCallError("CSSM_CSP_CreateSignatureContext", "10.0", _cSSM_CSP_CreateSignatureContextErr)
}

// TryCSSM_CSP_CreateSignatureContext calls CSSM_CSP_CreateSignatureContext without panicking when the symbol is unavailable.
func TryCSSM_CSP_CreateSignatureContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, AccessCred unsafe.Pointer, Key unsafe.Pointer, NewContextHandle unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CSP_CreateSignatureContextCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CSP_CreateSignatureContext(CSPHandle, AlgorithmID, AccessCred, Key, NewContextHandle), nil
}

// CSSM_CSP_CreateSignatureContext.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_CreateSignatureContext
func CSSM_CSP_CreateSignatureContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, AccessCred unsafe.Pointer, Key unsafe.Pointer, NewContextHandle unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CSP_CreateSignatureContext(CSPHandle, AlgorithmID, AccessCred, Key, NewContextHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CSP_CreateSymmetricContext func(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, Mode CSSM_ENCRYPT_MODE, AccessCred unsafe.Pointer, Key unsafe.Pointer, InitVector unsafe.Pointer, Padding CSSM_PADDING, Reserved unsafe.Pointer, NewContextHandle unsafe.Pointer) CSSM_RETURN
var _cSSM_CSP_CreateSymmetricContextErr error

// CanCallCSSM_CSP_CreateSymmetricContext reports whether the symbol for CSSM_CSP_CreateSymmetricContext is available on this system.
func CanCallCSSM_CSP_CreateSymmetricContext() bool {
	return _cSSM_CSP_CreateSymmetricContext != nil
}

// CSSM_CSP_CreateSymmetricContextCallError returns the symbol lookup or registration error for CSSM_CSP_CreateSymmetricContext.
func CSSM_CSP_CreateSymmetricContextCallError() error {
	if _cSSM_CSP_CreateSymmetricContext != nil {
		return nil
	}
	return symbolCallError("CSSM_CSP_CreateSymmetricContext", "10.0", _cSSM_CSP_CreateSymmetricContextErr)
}

// TryCSSM_CSP_CreateSymmetricContext calls CSSM_CSP_CreateSymmetricContext without panicking when the symbol is unavailable.
func TryCSSM_CSP_CreateSymmetricContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, Mode CSSM_ENCRYPT_MODE, AccessCred unsafe.Pointer, Key unsafe.Pointer, InitVector unsafe.Pointer, Padding CSSM_PADDING, Reserved unsafe.Pointer, NewContextHandle unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CSP_CreateSymmetricContextCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CSP_CreateSymmetricContext(CSPHandle, AlgorithmID, Mode, AccessCred, Key, InitVector, Padding, Reserved, NewContextHandle), nil
}

// CSSM_CSP_CreateSymmetricContext.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_CreateSymmetricContext
func CSSM_CSP_CreateSymmetricContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, Mode CSSM_ENCRYPT_MODE, AccessCred unsafe.Pointer, Key unsafe.Pointer, InitVector unsafe.Pointer, Padding CSSM_PADDING, Reserved unsafe.Pointer, NewContextHandle unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CSP_CreateSymmetricContext(CSPHandle, AlgorithmID, Mode, AccessCred, Key, InitVector, Padding, Reserved, NewContextHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CSP_GetLoginAcl func(CSPHandle CSSM_CSP_HANDLE, SelectionTag unsafe.Pointer, NumberOfAclInfos *uint32, AclInfos unsafe.Pointer) CSSM_RETURN
var _cSSM_CSP_GetLoginAclErr error

// CanCallCSSM_CSP_GetLoginAcl reports whether the symbol for CSSM_CSP_GetLoginAcl is available on this system.
func CanCallCSSM_CSP_GetLoginAcl() bool {
	return _cSSM_CSP_GetLoginAcl != nil
}

// CSSM_CSP_GetLoginAclCallError returns the symbol lookup or registration error for CSSM_CSP_GetLoginAcl.
func CSSM_CSP_GetLoginAclCallError() error {
	if _cSSM_CSP_GetLoginAcl != nil {
		return nil
	}
	return symbolCallError("CSSM_CSP_GetLoginAcl", "10.0", _cSSM_CSP_GetLoginAclErr)
}

// TryCSSM_CSP_GetLoginAcl calls CSSM_CSP_GetLoginAcl without panicking when the symbol is unavailable.
func TryCSSM_CSP_GetLoginAcl(CSPHandle CSSM_CSP_HANDLE, SelectionTag unsafe.Pointer, NumberOfAclInfos *uint32, AclInfos unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CSP_GetLoginAclCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CSP_GetLoginAcl(CSPHandle, SelectionTag, NumberOfAclInfos, AclInfos), nil
}

// CSSM_CSP_GetLoginAcl.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_GetLoginAcl
func CSSM_CSP_GetLoginAcl(CSPHandle CSSM_CSP_HANDLE, SelectionTag unsafe.Pointer, NumberOfAclInfos *uint32, AclInfos unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CSP_GetLoginAcl(CSPHandle, SelectionTag, NumberOfAclInfos, AclInfos)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CSP_GetLoginOwner func(CSPHandle CSSM_CSP_HANDLE, Owner unsafe.Pointer) CSSM_RETURN
var _cSSM_CSP_GetLoginOwnerErr error

// CanCallCSSM_CSP_GetLoginOwner reports whether the symbol for CSSM_CSP_GetLoginOwner is available on this system.
func CanCallCSSM_CSP_GetLoginOwner() bool {
	return _cSSM_CSP_GetLoginOwner != nil
}

// CSSM_CSP_GetLoginOwnerCallError returns the symbol lookup or registration error for CSSM_CSP_GetLoginOwner.
func CSSM_CSP_GetLoginOwnerCallError() error {
	if _cSSM_CSP_GetLoginOwner != nil {
		return nil
	}
	return symbolCallError("CSSM_CSP_GetLoginOwner", "10.0", _cSSM_CSP_GetLoginOwnerErr)
}

// TryCSSM_CSP_GetLoginOwner calls CSSM_CSP_GetLoginOwner without panicking when the symbol is unavailable.
func TryCSSM_CSP_GetLoginOwner(CSPHandle CSSM_CSP_HANDLE, Owner unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CSP_GetLoginOwnerCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CSP_GetLoginOwner(CSPHandle, Owner), nil
}

// CSSM_CSP_GetLoginOwner.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_GetLoginOwner
func CSSM_CSP_GetLoginOwner(CSPHandle CSSM_CSP_HANDLE, Owner unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CSP_GetLoginOwner(CSPHandle, Owner)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CSP_GetOperationalStatistics func(CSPHandle CSSM_CSP_HANDLE, Statistics unsafe.Pointer) CSSM_RETURN
var _cSSM_CSP_GetOperationalStatisticsErr error

// CanCallCSSM_CSP_GetOperationalStatistics reports whether the symbol for CSSM_CSP_GetOperationalStatistics is available on this system.
func CanCallCSSM_CSP_GetOperationalStatistics() bool {
	return _cSSM_CSP_GetOperationalStatistics != nil
}

// CSSM_CSP_GetOperationalStatisticsCallError returns the symbol lookup or registration error for CSSM_CSP_GetOperationalStatistics.
func CSSM_CSP_GetOperationalStatisticsCallError() error {
	if _cSSM_CSP_GetOperationalStatistics != nil {
		return nil
	}
	return symbolCallError("CSSM_CSP_GetOperationalStatistics", "10.0", _cSSM_CSP_GetOperationalStatisticsErr)
}

// TryCSSM_CSP_GetOperationalStatistics calls CSSM_CSP_GetOperationalStatistics without panicking when the symbol is unavailable.
func TryCSSM_CSP_GetOperationalStatistics(CSPHandle CSSM_CSP_HANDLE, Statistics unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CSP_GetOperationalStatisticsCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CSP_GetOperationalStatistics(CSPHandle, Statistics), nil
}

// CSSM_CSP_GetOperationalStatistics.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_GetOperationalStatistics
func CSSM_CSP_GetOperationalStatistics(CSPHandle CSSM_CSP_HANDLE, Statistics unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CSP_GetOperationalStatistics(CSPHandle, Statistics)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CSP_Login func(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, LoginName unsafe.Pointer, Reserved unsafe.Pointer) CSSM_RETURN
var _cSSM_CSP_LoginErr error

// CanCallCSSM_CSP_Login reports whether the symbol for CSSM_CSP_Login is available on this system.
func CanCallCSSM_CSP_Login() bool {
	return _cSSM_CSP_Login != nil
}

// CSSM_CSP_LoginCallError returns the symbol lookup or registration error for CSSM_CSP_Login.
func CSSM_CSP_LoginCallError() error {
	if _cSSM_CSP_Login != nil {
		return nil
	}
	return symbolCallError("CSSM_CSP_Login", "10.0", _cSSM_CSP_LoginErr)
}

// TryCSSM_CSP_Login calls CSSM_CSP_Login without panicking when the symbol is unavailable.
func TryCSSM_CSP_Login(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, LoginName unsafe.Pointer, Reserved unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CSP_LoginCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CSP_Login(CSPHandle, AccessCred, LoginName, Reserved), nil
}

// CSSM_CSP_Login.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_Login
func CSSM_CSP_Login(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, LoginName unsafe.Pointer, Reserved unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CSP_Login(CSPHandle, AccessCred, LoginName, Reserved)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CSP_Logout func(CSPHandle CSSM_CSP_HANDLE) CSSM_RETURN
var _cSSM_CSP_LogoutErr error

// CanCallCSSM_CSP_Logout reports whether the symbol for CSSM_CSP_Logout is available on this system.
func CanCallCSSM_CSP_Logout() bool {
	return _cSSM_CSP_Logout != nil
}

// CSSM_CSP_LogoutCallError returns the symbol lookup or registration error for CSSM_CSP_Logout.
func CSSM_CSP_LogoutCallError() error {
	if _cSSM_CSP_Logout != nil {
		return nil
	}
	return symbolCallError("CSSM_CSP_Logout", "10.0", _cSSM_CSP_LogoutErr)
}

// TryCSSM_CSP_Logout calls CSSM_CSP_Logout without panicking when the symbol is unavailable.
func TryCSSM_CSP_Logout(CSPHandle CSSM_CSP_HANDLE) (CSSM_RETURN, error) {
	if err := CSSM_CSP_LogoutCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CSP_Logout(CSPHandle), nil
}

// CSSM_CSP_Logout.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_Logout
func CSSM_CSP_Logout(CSPHandle CSSM_CSP_HANDLE) CSSM_RETURN {
	result, callErr := TryCSSM_CSP_Logout(CSPHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CSP_ObtainPrivateKeyFromPublicKey func(CSPHandle CSSM_CSP_HANDLE, PublicKey unsafe.Pointer, PrivateKey unsafe.Pointer) CSSM_RETURN
var _cSSM_CSP_ObtainPrivateKeyFromPublicKeyErr error

// CanCallCSSM_CSP_ObtainPrivateKeyFromPublicKey reports whether the symbol for CSSM_CSP_ObtainPrivateKeyFromPublicKey is available on this system.
func CanCallCSSM_CSP_ObtainPrivateKeyFromPublicKey() bool {
	return _cSSM_CSP_ObtainPrivateKeyFromPublicKey != nil
}

// CSSM_CSP_ObtainPrivateKeyFromPublicKeyCallError returns the symbol lookup or registration error for CSSM_CSP_ObtainPrivateKeyFromPublicKey.
func CSSM_CSP_ObtainPrivateKeyFromPublicKeyCallError() error {
	if _cSSM_CSP_ObtainPrivateKeyFromPublicKey != nil {
		return nil
	}
	return symbolCallError("CSSM_CSP_ObtainPrivateKeyFromPublicKey", "10.0", _cSSM_CSP_ObtainPrivateKeyFromPublicKeyErr)
}

// TryCSSM_CSP_ObtainPrivateKeyFromPublicKey calls CSSM_CSP_ObtainPrivateKeyFromPublicKey without panicking when the symbol is unavailable.
func TryCSSM_CSP_ObtainPrivateKeyFromPublicKey(CSPHandle CSSM_CSP_HANDLE, PublicKey unsafe.Pointer, PrivateKey unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CSP_ObtainPrivateKeyFromPublicKeyCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CSP_ObtainPrivateKeyFromPublicKey(CSPHandle, PublicKey, PrivateKey), nil
}

// CSSM_CSP_ObtainPrivateKeyFromPublicKey.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_ObtainPrivateKeyFromPublicKey
func CSSM_CSP_ObtainPrivateKeyFromPublicKey(CSPHandle CSSM_CSP_HANDLE, PublicKey unsafe.Pointer, PrivateKey unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CSP_ObtainPrivateKeyFromPublicKey(CSPHandle, PublicKey, PrivateKey)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_CSP_PassThrough func(CCHandle CSSM_CC_HANDLE, PassThroughId uint32, InData unsafe.Pointer, OutData unsafe.Pointer) CSSM_RETURN
var _cSSM_CSP_PassThroughErr error

// CanCallCSSM_CSP_PassThrough reports whether the symbol for CSSM_CSP_PassThrough is available on this system.
func CanCallCSSM_CSP_PassThrough() bool {
	return _cSSM_CSP_PassThrough != nil
}

// CSSM_CSP_PassThroughCallError returns the symbol lookup or registration error for CSSM_CSP_PassThrough.
func CSSM_CSP_PassThroughCallError() error {
	if _cSSM_CSP_PassThrough != nil {
		return nil
	}
	return symbolCallError("CSSM_CSP_PassThrough", "10.0", _cSSM_CSP_PassThroughErr)
}

// TryCSSM_CSP_PassThrough calls CSSM_CSP_PassThrough without panicking when the symbol is unavailable.
func TryCSSM_CSP_PassThrough(CCHandle CSSM_CC_HANDLE, PassThroughId uint32, InData unsafe.Pointer, OutData unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_CSP_PassThroughCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_CSP_PassThrough(CCHandle, PassThroughId, InData, OutData), nil
}

// CSSM_CSP_PassThrough.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_PassThrough
func CSSM_CSP_PassThrough(CCHandle CSSM_CC_HANDLE, PassThroughId uint32, InData unsafe.Pointer, OutData unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_CSP_PassThrough(CCHandle, PassThroughId, InData, OutData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_ChangeKeyAcl func(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, AclEdit unsafe.Pointer, Key unsafe.Pointer) CSSM_RETURN
var _cSSM_ChangeKeyAclErr error

// CanCallCSSM_ChangeKeyAcl reports whether the symbol for CSSM_ChangeKeyAcl is available on this system.
func CanCallCSSM_ChangeKeyAcl() bool {
	return _cSSM_ChangeKeyAcl != nil
}

// CSSM_ChangeKeyAclCallError returns the symbol lookup or registration error for CSSM_ChangeKeyAcl.
func CSSM_ChangeKeyAclCallError() error {
	if _cSSM_ChangeKeyAcl != nil {
		return nil
	}
	return symbolCallError("CSSM_ChangeKeyAcl", "10.0", _cSSM_ChangeKeyAclErr)
}

// TryCSSM_ChangeKeyAcl calls CSSM_ChangeKeyAcl without panicking when the symbol is unavailable.
func TryCSSM_ChangeKeyAcl(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, AclEdit unsafe.Pointer, Key unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_ChangeKeyAclCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_ChangeKeyAcl(CSPHandle, AccessCred, AclEdit, Key), nil
}

// CSSM_ChangeKeyAcl.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_ChangeKeyAcl
func CSSM_ChangeKeyAcl(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, AclEdit unsafe.Pointer, Key unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_ChangeKeyAcl(CSPHandle, AccessCred, AclEdit, Key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_ChangeKeyOwner func(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, Key unsafe.Pointer, NewOwner unsafe.Pointer) CSSM_RETURN
var _cSSM_ChangeKeyOwnerErr error

// CanCallCSSM_ChangeKeyOwner reports whether the symbol for CSSM_ChangeKeyOwner is available on this system.
func CanCallCSSM_ChangeKeyOwner() bool {
	return _cSSM_ChangeKeyOwner != nil
}

// CSSM_ChangeKeyOwnerCallError returns the symbol lookup or registration error for CSSM_ChangeKeyOwner.
func CSSM_ChangeKeyOwnerCallError() error {
	if _cSSM_ChangeKeyOwner != nil {
		return nil
	}
	return symbolCallError("CSSM_ChangeKeyOwner", "10.0", _cSSM_ChangeKeyOwnerErr)
}

// TryCSSM_ChangeKeyOwner calls CSSM_ChangeKeyOwner without panicking when the symbol is unavailable.
func TryCSSM_ChangeKeyOwner(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, Key unsafe.Pointer, NewOwner unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_ChangeKeyOwnerCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_ChangeKeyOwner(CSPHandle, AccessCred, Key, NewOwner), nil
}

// CSSM_ChangeKeyOwner.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_ChangeKeyOwner
func CSSM_ChangeKeyOwner(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, Key unsafe.Pointer, NewOwner unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_ChangeKeyOwner(CSPHandle, AccessCred, Key, NewOwner)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_Authenticate func(DLDBHandle unsafe.Pointer, AccessRequest CSSM_DB_ACCESS_TYPE, AccessCred unsafe.Pointer) CSSM_RETURN
var _cSSM_DL_AuthenticateErr error

// CanCallCSSM_DL_Authenticate reports whether the symbol for CSSM_DL_Authenticate is available on this system.
func CanCallCSSM_DL_Authenticate() bool {
	return _cSSM_DL_Authenticate != nil
}

// CSSM_DL_AuthenticateCallError returns the symbol lookup or registration error for CSSM_DL_Authenticate.
func CSSM_DL_AuthenticateCallError() error {
	if _cSSM_DL_Authenticate != nil {
		return nil
	}
	return symbolCallError("CSSM_DL_Authenticate", "10.0", _cSSM_DL_AuthenticateErr)
}

// TryCSSM_DL_Authenticate calls CSSM_DL_Authenticate without panicking when the symbol is unavailable.
func TryCSSM_DL_Authenticate(DLDBHandle unsafe.Pointer, AccessRequest CSSM_DB_ACCESS_TYPE, AccessCred unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_DL_AuthenticateCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_DL_Authenticate(DLDBHandle, AccessRequest, AccessCred), nil
}

// CSSM_DL_Authenticate.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_Authenticate
func CSSM_DL_Authenticate(DLDBHandle unsafe.Pointer, AccessRequest CSSM_DB_ACCESS_TYPE, AccessCred unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_DL_Authenticate(DLDBHandle, AccessRequest, AccessCred)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_ChangeDbAcl func(DLDBHandle unsafe.Pointer, AccessCred unsafe.Pointer, AclEdit unsafe.Pointer) CSSM_RETURN
var _cSSM_DL_ChangeDbAclErr error

// CanCallCSSM_DL_ChangeDbAcl reports whether the symbol for CSSM_DL_ChangeDbAcl is available on this system.
func CanCallCSSM_DL_ChangeDbAcl() bool {
	return _cSSM_DL_ChangeDbAcl != nil
}

// CSSM_DL_ChangeDbAclCallError returns the symbol lookup or registration error for CSSM_DL_ChangeDbAcl.
func CSSM_DL_ChangeDbAclCallError() error {
	if _cSSM_DL_ChangeDbAcl != nil {
		return nil
	}
	return symbolCallError("CSSM_DL_ChangeDbAcl", "10.0", _cSSM_DL_ChangeDbAclErr)
}

// TryCSSM_DL_ChangeDbAcl calls CSSM_DL_ChangeDbAcl without panicking when the symbol is unavailable.
func TryCSSM_DL_ChangeDbAcl(DLDBHandle unsafe.Pointer, AccessCred unsafe.Pointer, AclEdit unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_DL_ChangeDbAclCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_DL_ChangeDbAcl(DLDBHandle, AccessCred, AclEdit), nil
}

// CSSM_DL_ChangeDbAcl.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_ChangeDbAcl
func CSSM_DL_ChangeDbAcl(DLDBHandle unsafe.Pointer, AccessCred unsafe.Pointer, AclEdit unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_DL_ChangeDbAcl(DLDBHandle, AccessCred, AclEdit)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_ChangeDbOwner func(DLDBHandle unsafe.Pointer, AccessCred unsafe.Pointer, NewOwner unsafe.Pointer) CSSM_RETURN
var _cSSM_DL_ChangeDbOwnerErr error

// CanCallCSSM_DL_ChangeDbOwner reports whether the symbol for CSSM_DL_ChangeDbOwner is available on this system.
func CanCallCSSM_DL_ChangeDbOwner() bool {
	return _cSSM_DL_ChangeDbOwner != nil
}

// CSSM_DL_ChangeDbOwnerCallError returns the symbol lookup or registration error for CSSM_DL_ChangeDbOwner.
func CSSM_DL_ChangeDbOwnerCallError() error {
	if _cSSM_DL_ChangeDbOwner != nil {
		return nil
	}
	return symbolCallError("CSSM_DL_ChangeDbOwner", "10.0", _cSSM_DL_ChangeDbOwnerErr)
}

// TryCSSM_DL_ChangeDbOwner calls CSSM_DL_ChangeDbOwner without panicking when the symbol is unavailable.
func TryCSSM_DL_ChangeDbOwner(DLDBHandle unsafe.Pointer, AccessCred unsafe.Pointer, NewOwner unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_DL_ChangeDbOwnerCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_DL_ChangeDbOwner(DLDBHandle, AccessCred, NewOwner), nil
}

// CSSM_DL_ChangeDbOwner.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_ChangeDbOwner
func CSSM_DL_ChangeDbOwner(DLDBHandle unsafe.Pointer, AccessCred unsafe.Pointer, NewOwner unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_DL_ChangeDbOwner(DLDBHandle, AccessCred, NewOwner)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_CreateRelation func(DLDBHandle unsafe.Pointer, RelationID CSSM_DB_RECORDTYPE, RelationName string, NumberOfAttributes uint32, pAttributeInfo unsafe.Pointer, NumberOfIndexes uint32, pIndexInfo unsafe.Pointer) CSSM_RETURN
var _cSSM_DL_CreateRelationErr error

// CanCallCSSM_DL_CreateRelation reports whether the symbol for CSSM_DL_CreateRelation is available on this system.
func CanCallCSSM_DL_CreateRelation() bool {
	return _cSSM_DL_CreateRelation != nil
}

// CSSM_DL_CreateRelationCallError returns the symbol lookup or registration error for CSSM_DL_CreateRelation.
func CSSM_DL_CreateRelationCallError() error {
	if _cSSM_DL_CreateRelation != nil {
		return nil
	}
	return symbolCallError("CSSM_DL_CreateRelation", "10.0", _cSSM_DL_CreateRelationErr)
}

// TryCSSM_DL_CreateRelation calls CSSM_DL_CreateRelation without panicking when the symbol is unavailable.
func TryCSSM_DL_CreateRelation(DLDBHandle unsafe.Pointer, RelationID CSSM_DB_RECORDTYPE, RelationName string, NumberOfAttributes uint32, pAttributeInfo unsafe.Pointer, NumberOfIndexes uint32, pIndexInfo unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_DL_CreateRelationCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_DL_CreateRelation(DLDBHandle, RelationID, RelationName, NumberOfAttributes, pAttributeInfo, NumberOfIndexes, pIndexInfo), nil
}

// CSSM_DL_CreateRelation.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_CreateRelation
func CSSM_DL_CreateRelation(DLDBHandle unsafe.Pointer, RelationID CSSM_DB_RECORDTYPE, RelationName string, NumberOfAttributes uint32, pAttributeInfo unsafe.Pointer, NumberOfIndexes uint32, pIndexInfo unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_DL_CreateRelation(DLDBHandle, RelationID, RelationName, NumberOfAttributes, pAttributeInfo, NumberOfIndexes, pIndexInfo)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_DataAbortQuery func(DLDBHandle unsafe.Pointer, ResultsHandle CSSM_HANDLE) CSSM_RETURN
var _cSSM_DL_DataAbortQueryErr error

// CanCallCSSM_DL_DataAbortQuery reports whether the symbol for CSSM_DL_DataAbortQuery is available on this system.
func CanCallCSSM_DL_DataAbortQuery() bool {
	return _cSSM_DL_DataAbortQuery != nil
}

// CSSM_DL_DataAbortQueryCallError returns the symbol lookup or registration error for CSSM_DL_DataAbortQuery.
func CSSM_DL_DataAbortQueryCallError() error {
	if _cSSM_DL_DataAbortQuery != nil {
		return nil
	}
	return symbolCallError("CSSM_DL_DataAbortQuery", "10.0", _cSSM_DL_DataAbortQueryErr)
}

// TryCSSM_DL_DataAbortQuery calls CSSM_DL_DataAbortQuery without panicking when the symbol is unavailable.
func TryCSSM_DL_DataAbortQuery(DLDBHandle unsafe.Pointer, ResultsHandle CSSM_HANDLE) (CSSM_RETURN, error) {
	if err := CSSM_DL_DataAbortQueryCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_DL_DataAbortQuery(DLDBHandle, ResultsHandle), nil
}

// CSSM_DL_DataAbortQuery.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_DataAbortQuery
func CSSM_DL_DataAbortQuery(DLDBHandle unsafe.Pointer, ResultsHandle CSSM_HANDLE) CSSM_RETURN {
	result, callErr := TryCSSM_DL_DataAbortQuery(DLDBHandle, ResultsHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_DataDelete func(DLDBHandle unsafe.Pointer, UniqueRecordIdentifier unsafe.Pointer) CSSM_RETURN
var _cSSM_DL_DataDeleteErr error

// CanCallCSSM_DL_DataDelete reports whether the symbol for CSSM_DL_DataDelete is available on this system.
func CanCallCSSM_DL_DataDelete() bool {
	return _cSSM_DL_DataDelete != nil
}

// CSSM_DL_DataDeleteCallError returns the symbol lookup or registration error for CSSM_DL_DataDelete.
func CSSM_DL_DataDeleteCallError() error {
	if _cSSM_DL_DataDelete != nil {
		return nil
	}
	return symbolCallError("CSSM_DL_DataDelete", "10.0", _cSSM_DL_DataDeleteErr)
}

// TryCSSM_DL_DataDelete calls CSSM_DL_DataDelete without panicking when the symbol is unavailable.
func TryCSSM_DL_DataDelete(DLDBHandle unsafe.Pointer, UniqueRecordIdentifier unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_DL_DataDeleteCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_DL_DataDelete(DLDBHandle, UniqueRecordIdentifier), nil
}

// CSSM_DL_DataDelete.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_DataDelete
func CSSM_DL_DataDelete(DLDBHandle unsafe.Pointer, UniqueRecordIdentifier unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_DL_DataDelete(DLDBHandle, UniqueRecordIdentifier)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_DataGetFirst func(DLDBHandle unsafe.Pointer, Query unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, Attributes unsafe.Pointer, Data unsafe.Pointer, UniqueId unsafe.Pointer) CSSM_RETURN
var _cSSM_DL_DataGetFirstErr error

// CanCallCSSM_DL_DataGetFirst reports whether the symbol for CSSM_DL_DataGetFirst is available on this system.
func CanCallCSSM_DL_DataGetFirst() bool {
	return _cSSM_DL_DataGetFirst != nil
}

// CSSM_DL_DataGetFirstCallError returns the symbol lookup or registration error for CSSM_DL_DataGetFirst.
func CSSM_DL_DataGetFirstCallError() error {
	if _cSSM_DL_DataGetFirst != nil {
		return nil
	}
	return symbolCallError("CSSM_DL_DataGetFirst", "10.0", _cSSM_DL_DataGetFirstErr)
}

// TryCSSM_DL_DataGetFirst calls CSSM_DL_DataGetFirst without panicking when the symbol is unavailable.
func TryCSSM_DL_DataGetFirst(DLDBHandle unsafe.Pointer, Query unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, Attributes unsafe.Pointer, Data unsafe.Pointer, UniqueId unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_DL_DataGetFirstCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_DL_DataGetFirst(DLDBHandle, Query, ResultsHandle, Attributes, Data, UniqueId), nil
}

// CSSM_DL_DataGetFirst.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_DataGetFirst
func CSSM_DL_DataGetFirst(DLDBHandle unsafe.Pointer, Query unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, Attributes unsafe.Pointer, Data unsafe.Pointer, UniqueId unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_DL_DataGetFirst(DLDBHandle, Query, ResultsHandle, Attributes, Data, UniqueId)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_DataGetFromUniqueRecordId func(DLDBHandle unsafe.Pointer, UniqueRecord unsafe.Pointer, Attributes unsafe.Pointer, Data unsafe.Pointer) CSSM_RETURN
var _cSSM_DL_DataGetFromUniqueRecordIdErr error

// CanCallCSSM_DL_DataGetFromUniqueRecordId reports whether the symbol for CSSM_DL_DataGetFromUniqueRecordId is available on this system.
func CanCallCSSM_DL_DataGetFromUniqueRecordId() bool {
	return _cSSM_DL_DataGetFromUniqueRecordId != nil
}

// CSSM_DL_DataGetFromUniqueRecordIdCallError returns the symbol lookup or registration error for CSSM_DL_DataGetFromUniqueRecordId.
func CSSM_DL_DataGetFromUniqueRecordIdCallError() error {
	if _cSSM_DL_DataGetFromUniqueRecordId != nil {
		return nil
	}
	return symbolCallError("CSSM_DL_DataGetFromUniqueRecordId", "10.0", _cSSM_DL_DataGetFromUniqueRecordIdErr)
}

// TryCSSM_DL_DataGetFromUniqueRecordId calls CSSM_DL_DataGetFromUniqueRecordId without panicking when the symbol is unavailable.
func TryCSSM_DL_DataGetFromUniqueRecordId(DLDBHandle unsafe.Pointer, UniqueRecord unsafe.Pointer, Attributes unsafe.Pointer, Data unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_DL_DataGetFromUniqueRecordIdCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_DL_DataGetFromUniqueRecordId(DLDBHandle, UniqueRecord, Attributes, Data), nil
}

// CSSM_DL_DataGetFromUniqueRecordId.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_DataGetFromUniqueRecordId
func CSSM_DL_DataGetFromUniqueRecordId(DLDBHandle unsafe.Pointer, UniqueRecord unsafe.Pointer, Attributes unsafe.Pointer, Data unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_DL_DataGetFromUniqueRecordId(DLDBHandle, UniqueRecord, Attributes, Data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_DataGetNext func(DLDBHandle unsafe.Pointer, ResultsHandle CSSM_HANDLE, Attributes unsafe.Pointer, Data unsafe.Pointer, UniqueId unsafe.Pointer) CSSM_RETURN
var _cSSM_DL_DataGetNextErr error

// CanCallCSSM_DL_DataGetNext reports whether the symbol for CSSM_DL_DataGetNext is available on this system.
func CanCallCSSM_DL_DataGetNext() bool {
	return _cSSM_DL_DataGetNext != nil
}

// CSSM_DL_DataGetNextCallError returns the symbol lookup or registration error for CSSM_DL_DataGetNext.
func CSSM_DL_DataGetNextCallError() error {
	if _cSSM_DL_DataGetNext != nil {
		return nil
	}
	return symbolCallError("CSSM_DL_DataGetNext", "10.0", _cSSM_DL_DataGetNextErr)
}

// TryCSSM_DL_DataGetNext calls CSSM_DL_DataGetNext without panicking when the symbol is unavailable.
func TryCSSM_DL_DataGetNext(DLDBHandle unsafe.Pointer, ResultsHandle CSSM_HANDLE, Attributes unsafe.Pointer, Data unsafe.Pointer, UniqueId unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_DL_DataGetNextCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_DL_DataGetNext(DLDBHandle, ResultsHandle, Attributes, Data, UniqueId), nil
}

// CSSM_DL_DataGetNext.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_DataGetNext
func CSSM_DL_DataGetNext(DLDBHandle unsafe.Pointer, ResultsHandle CSSM_HANDLE, Attributes unsafe.Pointer, Data unsafe.Pointer, UniqueId unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_DL_DataGetNext(DLDBHandle, ResultsHandle, Attributes, Data, UniqueId)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_DataInsert func(DLDBHandle unsafe.Pointer, RecordType CSSM_DB_RECORDTYPE, Attributes unsafe.Pointer, Data unsafe.Pointer, UniqueId unsafe.Pointer) CSSM_RETURN
var _cSSM_DL_DataInsertErr error

// CanCallCSSM_DL_DataInsert reports whether the symbol for CSSM_DL_DataInsert is available on this system.
func CanCallCSSM_DL_DataInsert() bool {
	return _cSSM_DL_DataInsert != nil
}

// CSSM_DL_DataInsertCallError returns the symbol lookup or registration error for CSSM_DL_DataInsert.
func CSSM_DL_DataInsertCallError() error {
	if _cSSM_DL_DataInsert != nil {
		return nil
	}
	return symbolCallError("CSSM_DL_DataInsert", "10.0", _cSSM_DL_DataInsertErr)
}

// TryCSSM_DL_DataInsert calls CSSM_DL_DataInsert without panicking when the symbol is unavailable.
func TryCSSM_DL_DataInsert(DLDBHandle unsafe.Pointer, RecordType CSSM_DB_RECORDTYPE, Attributes unsafe.Pointer, Data unsafe.Pointer, UniqueId unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_DL_DataInsertCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_DL_DataInsert(DLDBHandle, RecordType, Attributes, Data, UniqueId), nil
}

// CSSM_DL_DataInsert.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_DataInsert
func CSSM_DL_DataInsert(DLDBHandle unsafe.Pointer, RecordType CSSM_DB_RECORDTYPE, Attributes unsafe.Pointer, Data unsafe.Pointer, UniqueId unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_DL_DataInsert(DLDBHandle, RecordType, Attributes, Data, UniqueId)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_DataModify func(DLDBHandle unsafe.Pointer, RecordType CSSM_DB_RECORDTYPE, UniqueRecordIdentifier unsafe.Pointer, AttributesToBeModified unsafe.Pointer, DataToBeModified unsafe.Pointer, ModifyMode CSSM_DB_MODIFY_MODE) CSSM_RETURN
var _cSSM_DL_DataModifyErr error

// CanCallCSSM_DL_DataModify reports whether the symbol for CSSM_DL_DataModify is available on this system.
func CanCallCSSM_DL_DataModify() bool {
	return _cSSM_DL_DataModify != nil
}

// CSSM_DL_DataModifyCallError returns the symbol lookup or registration error for CSSM_DL_DataModify.
func CSSM_DL_DataModifyCallError() error {
	if _cSSM_DL_DataModify != nil {
		return nil
	}
	return symbolCallError("CSSM_DL_DataModify", "10.0", _cSSM_DL_DataModifyErr)
}

// TryCSSM_DL_DataModify calls CSSM_DL_DataModify without panicking when the symbol is unavailable.
func TryCSSM_DL_DataModify(DLDBHandle unsafe.Pointer, RecordType CSSM_DB_RECORDTYPE, UniqueRecordIdentifier unsafe.Pointer, AttributesToBeModified unsafe.Pointer, DataToBeModified unsafe.Pointer, ModifyMode CSSM_DB_MODIFY_MODE) (CSSM_RETURN, error) {
	if err := CSSM_DL_DataModifyCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_DL_DataModify(DLDBHandle, RecordType, UniqueRecordIdentifier, AttributesToBeModified, DataToBeModified, ModifyMode), nil
}

// CSSM_DL_DataModify.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_DataModify
func CSSM_DL_DataModify(DLDBHandle unsafe.Pointer, RecordType CSSM_DB_RECORDTYPE, UniqueRecordIdentifier unsafe.Pointer, AttributesToBeModified unsafe.Pointer, DataToBeModified unsafe.Pointer, ModifyMode CSSM_DB_MODIFY_MODE) CSSM_RETURN {
	result, callErr := TryCSSM_DL_DataModify(DLDBHandle, RecordType, UniqueRecordIdentifier, AttributesToBeModified, DataToBeModified, ModifyMode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_DbClose func(DLDBHandle unsafe.Pointer) CSSM_RETURN
var _cSSM_DL_DbCloseErr error

// CanCallCSSM_DL_DbClose reports whether the symbol for CSSM_DL_DbClose is available on this system.
func CanCallCSSM_DL_DbClose() bool {
	return _cSSM_DL_DbClose != nil
}

// CSSM_DL_DbCloseCallError returns the symbol lookup or registration error for CSSM_DL_DbClose.
func CSSM_DL_DbCloseCallError() error {
	if _cSSM_DL_DbClose != nil {
		return nil
	}
	return symbolCallError("CSSM_DL_DbClose", "10.0", _cSSM_DL_DbCloseErr)
}

// TryCSSM_DL_DbClose calls CSSM_DL_DbClose without panicking when the symbol is unavailable.
func TryCSSM_DL_DbClose(DLDBHandle unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_DL_DbCloseCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_DL_DbClose(DLDBHandle), nil
}

// CSSM_DL_DbClose.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_DbClose
func CSSM_DL_DbClose(DLDBHandle unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_DL_DbClose(DLDBHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_DbCreate func(DLHandle CSSM_DL_HANDLE, DbName string, DbLocation unsafe.Pointer, DBInfo unsafe.Pointer, AccessRequest CSSM_DB_ACCESS_TYPE, CredAndAclEntry unsafe.Pointer, OpenParameters unsafe.Pointer, DbHandle unsafe.Pointer) CSSM_RETURN
var _cSSM_DL_DbCreateErr error

// CanCallCSSM_DL_DbCreate reports whether the symbol for CSSM_DL_DbCreate is available on this system.
func CanCallCSSM_DL_DbCreate() bool {
	return _cSSM_DL_DbCreate != nil
}

// CSSM_DL_DbCreateCallError returns the symbol lookup or registration error for CSSM_DL_DbCreate.
func CSSM_DL_DbCreateCallError() error {
	if _cSSM_DL_DbCreate != nil {
		return nil
	}
	return symbolCallError("CSSM_DL_DbCreate", "10.0", _cSSM_DL_DbCreateErr)
}

// TryCSSM_DL_DbCreate calls CSSM_DL_DbCreate without panicking when the symbol is unavailable.
func TryCSSM_DL_DbCreate(DLHandle CSSM_DL_HANDLE, DbName string, DbLocation unsafe.Pointer, DBInfo unsafe.Pointer, AccessRequest CSSM_DB_ACCESS_TYPE, CredAndAclEntry unsafe.Pointer, OpenParameters unsafe.Pointer, DbHandle unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_DL_DbCreateCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_DL_DbCreate(DLHandle, DbName, DbLocation, DBInfo, AccessRequest, CredAndAclEntry, OpenParameters, DbHandle), nil
}

// CSSM_DL_DbCreate.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_DbCreate
func CSSM_DL_DbCreate(DLHandle CSSM_DL_HANDLE, DbName string, DbLocation unsafe.Pointer, DBInfo unsafe.Pointer, AccessRequest CSSM_DB_ACCESS_TYPE, CredAndAclEntry unsafe.Pointer, OpenParameters unsafe.Pointer, DbHandle unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_DL_DbCreate(DLHandle, DbName, DbLocation, DBInfo, AccessRequest, CredAndAclEntry, OpenParameters, DbHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_DbDelete func(DLHandle CSSM_DL_HANDLE, DbName string, DbLocation unsafe.Pointer, AccessCred unsafe.Pointer) CSSM_RETURN
var _cSSM_DL_DbDeleteErr error

// CanCallCSSM_DL_DbDelete reports whether the symbol for CSSM_DL_DbDelete is available on this system.
func CanCallCSSM_DL_DbDelete() bool {
	return _cSSM_DL_DbDelete != nil
}

// CSSM_DL_DbDeleteCallError returns the symbol lookup or registration error for CSSM_DL_DbDelete.
func CSSM_DL_DbDeleteCallError() error {
	if _cSSM_DL_DbDelete != nil {
		return nil
	}
	return symbolCallError("CSSM_DL_DbDelete", "10.0", _cSSM_DL_DbDeleteErr)
}

// TryCSSM_DL_DbDelete calls CSSM_DL_DbDelete without panicking when the symbol is unavailable.
func TryCSSM_DL_DbDelete(DLHandle CSSM_DL_HANDLE, DbName string, DbLocation unsafe.Pointer, AccessCred unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_DL_DbDeleteCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_DL_DbDelete(DLHandle, DbName, DbLocation, AccessCred), nil
}

// CSSM_DL_DbDelete.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_DbDelete
func CSSM_DL_DbDelete(DLHandle CSSM_DL_HANDLE, DbName string, DbLocation unsafe.Pointer, AccessCred unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_DL_DbDelete(DLHandle, DbName, DbLocation, AccessCred)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_DbOpen func(DLHandle CSSM_DL_HANDLE, DbName string, DbLocation unsafe.Pointer, AccessRequest CSSM_DB_ACCESS_TYPE, AccessCred unsafe.Pointer, OpenParameters unsafe.Pointer, DbHandle unsafe.Pointer) CSSM_RETURN
var _cSSM_DL_DbOpenErr error

// CanCallCSSM_DL_DbOpen reports whether the symbol for CSSM_DL_DbOpen is available on this system.
func CanCallCSSM_DL_DbOpen() bool {
	return _cSSM_DL_DbOpen != nil
}

// CSSM_DL_DbOpenCallError returns the symbol lookup or registration error for CSSM_DL_DbOpen.
func CSSM_DL_DbOpenCallError() error {
	if _cSSM_DL_DbOpen != nil {
		return nil
	}
	return symbolCallError("CSSM_DL_DbOpen", "10.0", _cSSM_DL_DbOpenErr)
}

// TryCSSM_DL_DbOpen calls CSSM_DL_DbOpen without panicking when the symbol is unavailable.
func TryCSSM_DL_DbOpen(DLHandle CSSM_DL_HANDLE, DbName string, DbLocation unsafe.Pointer, AccessRequest CSSM_DB_ACCESS_TYPE, AccessCred unsafe.Pointer, OpenParameters unsafe.Pointer, DbHandle unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_DL_DbOpenCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_DL_DbOpen(DLHandle, DbName, DbLocation, AccessRequest, AccessCred, OpenParameters, DbHandle), nil
}

// CSSM_DL_DbOpen.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_DbOpen
func CSSM_DL_DbOpen(DLHandle CSSM_DL_HANDLE, DbName string, DbLocation unsafe.Pointer, AccessRequest CSSM_DB_ACCESS_TYPE, AccessCred unsafe.Pointer, OpenParameters unsafe.Pointer, DbHandle unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_DL_DbOpen(DLHandle, DbName, DbLocation, AccessRequest, AccessCred, OpenParameters, DbHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_DestroyRelation func(DLDBHandle unsafe.Pointer, RelationID CSSM_DB_RECORDTYPE) CSSM_RETURN
var _cSSM_DL_DestroyRelationErr error

// CanCallCSSM_DL_DestroyRelation reports whether the symbol for CSSM_DL_DestroyRelation is available on this system.
func CanCallCSSM_DL_DestroyRelation() bool {
	return _cSSM_DL_DestroyRelation != nil
}

// CSSM_DL_DestroyRelationCallError returns the symbol lookup or registration error for CSSM_DL_DestroyRelation.
func CSSM_DL_DestroyRelationCallError() error {
	if _cSSM_DL_DestroyRelation != nil {
		return nil
	}
	return symbolCallError("CSSM_DL_DestroyRelation", "10.0", _cSSM_DL_DestroyRelationErr)
}

// TryCSSM_DL_DestroyRelation calls CSSM_DL_DestroyRelation without panicking when the symbol is unavailable.
func TryCSSM_DL_DestroyRelation(DLDBHandle unsafe.Pointer, RelationID CSSM_DB_RECORDTYPE) (CSSM_RETURN, error) {
	if err := CSSM_DL_DestroyRelationCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_DL_DestroyRelation(DLDBHandle, RelationID), nil
}

// CSSM_DL_DestroyRelation.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_DestroyRelation
func CSSM_DL_DestroyRelation(DLDBHandle unsafe.Pointer, RelationID CSSM_DB_RECORDTYPE) CSSM_RETURN {
	result, callErr := TryCSSM_DL_DestroyRelation(DLDBHandle, RelationID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_FreeNameList func(DLHandle CSSM_DL_HANDLE, NameList unsafe.Pointer) CSSM_RETURN
var _cSSM_DL_FreeNameListErr error

// CanCallCSSM_DL_FreeNameList reports whether the symbol for CSSM_DL_FreeNameList is available on this system.
func CanCallCSSM_DL_FreeNameList() bool {
	return _cSSM_DL_FreeNameList != nil
}

// CSSM_DL_FreeNameListCallError returns the symbol lookup or registration error for CSSM_DL_FreeNameList.
func CSSM_DL_FreeNameListCallError() error {
	if _cSSM_DL_FreeNameList != nil {
		return nil
	}
	return symbolCallError("CSSM_DL_FreeNameList", "10.0", _cSSM_DL_FreeNameListErr)
}

// TryCSSM_DL_FreeNameList calls CSSM_DL_FreeNameList without panicking when the symbol is unavailable.
func TryCSSM_DL_FreeNameList(DLHandle CSSM_DL_HANDLE, NameList unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_DL_FreeNameListCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_DL_FreeNameList(DLHandle, NameList), nil
}

// CSSM_DL_FreeNameList.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_FreeNameList
func CSSM_DL_FreeNameList(DLHandle CSSM_DL_HANDLE, NameList unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_DL_FreeNameList(DLHandle, NameList)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_FreeUniqueRecord func(DLDBHandle unsafe.Pointer, UniqueRecord unsafe.Pointer) CSSM_RETURN
var _cSSM_DL_FreeUniqueRecordErr error

// CanCallCSSM_DL_FreeUniqueRecord reports whether the symbol for CSSM_DL_FreeUniqueRecord is available on this system.
func CanCallCSSM_DL_FreeUniqueRecord() bool {
	return _cSSM_DL_FreeUniqueRecord != nil
}

// CSSM_DL_FreeUniqueRecordCallError returns the symbol lookup or registration error for CSSM_DL_FreeUniqueRecord.
func CSSM_DL_FreeUniqueRecordCallError() error {
	if _cSSM_DL_FreeUniqueRecord != nil {
		return nil
	}
	return symbolCallError("CSSM_DL_FreeUniqueRecord", "10.0", _cSSM_DL_FreeUniqueRecordErr)
}

// TryCSSM_DL_FreeUniqueRecord calls CSSM_DL_FreeUniqueRecord without panicking when the symbol is unavailable.
func TryCSSM_DL_FreeUniqueRecord(DLDBHandle unsafe.Pointer, UniqueRecord unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_DL_FreeUniqueRecordCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_DL_FreeUniqueRecord(DLDBHandle, UniqueRecord), nil
}

// CSSM_DL_FreeUniqueRecord.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_FreeUniqueRecord
func CSSM_DL_FreeUniqueRecord(DLDBHandle unsafe.Pointer, UniqueRecord unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_DL_FreeUniqueRecord(DLDBHandle, UniqueRecord)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_GetDbAcl func(DLDBHandle unsafe.Pointer, SelectionTag unsafe.Pointer, NumberOfAclInfos *uint32, AclInfos unsafe.Pointer) CSSM_RETURN
var _cSSM_DL_GetDbAclErr error

// CanCallCSSM_DL_GetDbAcl reports whether the symbol for CSSM_DL_GetDbAcl is available on this system.
func CanCallCSSM_DL_GetDbAcl() bool {
	return _cSSM_DL_GetDbAcl != nil
}

// CSSM_DL_GetDbAclCallError returns the symbol lookup or registration error for CSSM_DL_GetDbAcl.
func CSSM_DL_GetDbAclCallError() error {
	if _cSSM_DL_GetDbAcl != nil {
		return nil
	}
	return symbolCallError("CSSM_DL_GetDbAcl", "10.0", _cSSM_DL_GetDbAclErr)
}

// TryCSSM_DL_GetDbAcl calls CSSM_DL_GetDbAcl without panicking when the symbol is unavailable.
func TryCSSM_DL_GetDbAcl(DLDBHandle unsafe.Pointer, SelectionTag unsafe.Pointer, NumberOfAclInfos *uint32, AclInfos unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_DL_GetDbAclCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_DL_GetDbAcl(DLDBHandle, SelectionTag, NumberOfAclInfos, AclInfos), nil
}

// CSSM_DL_GetDbAcl.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_GetDbAcl
func CSSM_DL_GetDbAcl(DLDBHandle unsafe.Pointer, SelectionTag unsafe.Pointer, NumberOfAclInfos *uint32, AclInfos unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_DL_GetDbAcl(DLDBHandle, SelectionTag, NumberOfAclInfos, AclInfos)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_GetDbNameFromHandle func(DLDBHandle unsafe.Pointer, DbName string) CSSM_RETURN
var _cSSM_DL_GetDbNameFromHandleErr error

// CanCallCSSM_DL_GetDbNameFromHandle reports whether the symbol for CSSM_DL_GetDbNameFromHandle is available on this system.
func CanCallCSSM_DL_GetDbNameFromHandle() bool {
	return _cSSM_DL_GetDbNameFromHandle != nil
}

// CSSM_DL_GetDbNameFromHandleCallError returns the symbol lookup or registration error for CSSM_DL_GetDbNameFromHandle.
func CSSM_DL_GetDbNameFromHandleCallError() error {
	if _cSSM_DL_GetDbNameFromHandle != nil {
		return nil
	}
	return symbolCallError("CSSM_DL_GetDbNameFromHandle", "10.0", _cSSM_DL_GetDbNameFromHandleErr)
}

// TryCSSM_DL_GetDbNameFromHandle calls CSSM_DL_GetDbNameFromHandle without panicking when the symbol is unavailable.
func TryCSSM_DL_GetDbNameFromHandle(DLDBHandle unsafe.Pointer, DbName string) (CSSM_RETURN, error) {
	if err := CSSM_DL_GetDbNameFromHandleCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_DL_GetDbNameFromHandle(DLDBHandle, DbName), nil
}

// CSSM_DL_GetDbNameFromHandle.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_GetDbNameFromHandle
func CSSM_DL_GetDbNameFromHandle(DLDBHandle unsafe.Pointer, DbName string) CSSM_RETURN {
	result, callErr := TryCSSM_DL_GetDbNameFromHandle(DLDBHandle, DbName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_GetDbNames func(DLHandle CSSM_DL_HANDLE, NameList unsafe.Pointer) CSSM_RETURN
var _cSSM_DL_GetDbNamesErr error

// CanCallCSSM_DL_GetDbNames reports whether the symbol for CSSM_DL_GetDbNames is available on this system.
func CanCallCSSM_DL_GetDbNames() bool {
	return _cSSM_DL_GetDbNames != nil
}

// CSSM_DL_GetDbNamesCallError returns the symbol lookup or registration error for CSSM_DL_GetDbNames.
func CSSM_DL_GetDbNamesCallError() error {
	if _cSSM_DL_GetDbNames != nil {
		return nil
	}
	return symbolCallError("CSSM_DL_GetDbNames", "10.0", _cSSM_DL_GetDbNamesErr)
}

// TryCSSM_DL_GetDbNames calls CSSM_DL_GetDbNames without panicking when the symbol is unavailable.
func TryCSSM_DL_GetDbNames(DLHandle CSSM_DL_HANDLE, NameList unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_DL_GetDbNamesCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_DL_GetDbNames(DLHandle, NameList), nil
}

// CSSM_DL_GetDbNames.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_GetDbNames
func CSSM_DL_GetDbNames(DLHandle CSSM_DL_HANDLE, NameList unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_DL_GetDbNames(DLHandle, NameList)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_GetDbOwner func(DLDBHandle unsafe.Pointer, Owner unsafe.Pointer) CSSM_RETURN
var _cSSM_DL_GetDbOwnerErr error

// CanCallCSSM_DL_GetDbOwner reports whether the symbol for CSSM_DL_GetDbOwner is available on this system.
func CanCallCSSM_DL_GetDbOwner() bool {
	return _cSSM_DL_GetDbOwner != nil
}

// CSSM_DL_GetDbOwnerCallError returns the symbol lookup or registration error for CSSM_DL_GetDbOwner.
func CSSM_DL_GetDbOwnerCallError() error {
	if _cSSM_DL_GetDbOwner != nil {
		return nil
	}
	return symbolCallError("CSSM_DL_GetDbOwner", "10.0", _cSSM_DL_GetDbOwnerErr)
}

// TryCSSM_DL_GetDbOwner calls CSSM_DL_GetDbOwner without panicking when the symbol is unavailable.
func TryCSSM_DL_GetDbOwner(DLDBHandle unsafe.Pointer, Owner unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_DL_GetDbOwnerCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_DL_GetDbOwner(DLDBHandle, Owner), nil
}

// CSSM_DL_GetDbOwner.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_GetDbOwner
func CSSM_DL_GetDbOwner(DLDBHandle unsafe.Pointer, Owner unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_DL_GetDbOwner(DLDBHandle, Owner)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DL_PassThrough func(DLDBHandle unsafe.Pointer, PassThroughId uint32, InputParams unsafe.Pointer, OutputParams unsafe.Pointer) CSSM_RETURN
var _cSSM_DL_PassThroughErr error

// CanCallCSSM_DL_PassThrough reports whether the symbol for CSSM_DL_PassThrough is available on this system.
func CanCallCSSM_DL_PassThrough() bool {
	return _cSSM_DL_PassThrough != nil
}

// CSSM_DL_PassThroughCallError returns the symbol lookup or registration error for CSSM_DL_PassThrough.
func CSSM_DL_PassThroughCallError() error {
	if _cSSM_DL_PassThrough != nil {
		return nil
	}
	return symbolCallError("CSSM_DL_PassThrough", "10.0", _cSSM_DL_PassThroughErr)
}

// TryCSSM_DL_PassThrough calls CSSM_DL_PassThrough without panicking when the symbol is unavailable.
func TryCSSM_DL_PassThrough(DLDBHandle unsafe.Pointer, PassThroughId uint32, InputParams unsafe.Pointer, OutputParams unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_DL_PassThroughCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_DL_PassThrough(DLDBHandle, PassThroughId, InputParams, OutputParams), nil
}

// CSSM_DL_PassThrough.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_PassThrough
func CSSM_DL_PassThrough(DLDBHandle unsafe.Pointer, PassThroughId uint32, InputParams unsafe.Pointer, OutputParams unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_DL_PassThrough(DLDBHandle, PassThroughId, InputParams, OutputParams)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DecryptData func(CCHandle CSSM_CC_HANDLE, CipherBufs unsafe.Pointer, CipherBufCount uint32, ClearBufs unsafe.Pointer, ClearBufCount uint32, bytesDecrypted unsafe.Pointer, RemData unsafe.Pointer) CSSM_RETURN
var _cSSM_DecryptDataErr error

// CanCallCSSM_DecryptData reports whether the symbol for CSSM_DecryptData is available on this system.
func CanCallCSSM_DecryptData() bool {
	return _cSSM_DecryptData != nil
}

// CSSM_DecryptDataCallError returns the symbol lookup or registration error for CSSM_DecryptData.
func CSSM_DecryptDataCallError() error {
	if _cSSM_DecryptData != nil {
		return nil
	}
	return symbolCallError("CSSM_DecryptData", "10.0", _cSSM_DecryptDataErr)
}

// TryCSSM_DecryptData calls CSSM_DecryptData without panicking when the symbol is unavailable.
func TryCSSM_DecryptData(CCHandle CSSM_CC_HANDLE, CipherBufs unsafe.Pointer, CipherBufCount uint32, ClearBufs unsafe.Pointer, ClearBufCount uint32, bytesDecrypted unsafe.Pointer, RemData unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_DecryptDataCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_DecryptData(CCHandle, CipherBufs, CipherBufCount, ClearBufs, ClearBufCount, bytesDecrypted, RemData), nil
}

// CSSM_DecryptData.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DecryptData
func CSSM_DecryptData(CCHandle CSSM_CC_HANDLE, CipherBufs unsafe.Pointer, CipherBufCount uint32, ClearBufs unsafe.Pointer, ClearBufCount uint32, bytesDecrypted unsafe.Pointer, RemData unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_DecryptData(CCHandle, CipherBufs, CipherBufCount, ClearBufs, ClearBufCount, bytesDecrypted, RemData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DecryptDataFinal func(CCHandle CSSM_CC_HANDLE, RemData unsafe.Pointer) CSSM_RETURN
var _cSSM_DecryptDataFinalErr error

// CanCallCSSM_DecryptDataFinal reports whether the symbol for CSSM_DecryptDataFinal is available on this system.
func CanCallCSSM_DecryptDataFinal() bool {
	return _cSSM_DecryptDataFinal != nil
}

// CSSM_DecryptDataFinalCallError returns the symbol lookup or registration error for CSSM_DecryptDataFinal.
func CSSM_DecryptDataFinalCallError() error {
	if _cSSM_DecryptDataFinal != nil {
		return nil
	}
	return symbolCallError("CSSM_DecryptDataFinal", "10.0", _cSSM_DecryptDataFinalErr)
}

// TryCSSM_DecryptDataFinal calls CSSM_DecryptDataFinal without panicking when the symbol is unavailable.
func TryCSSM_DecryptDataFinal(CCHandle CSSM_CC_HANDLE, RemData unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_DecryptDataFinalCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_DecryptDataFinal(CCHandle, RemData), nil
}

// CSSM_DecryptDataFinal.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DecryptDataFinal
func CSSM_DecryptDataFinal(CCHandle CSSM_CC_HANDLE, RemData unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_DecryptDataFinal(CCHandle, RemData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DecryptDataInit func(CCHandle CSSM_CC_HANDLE) CSSM_RETURN
var _cSSM_DecryptDataInitErr error

// CanCallCSSM_DecryptDataInit reports whether the symbol for CSSM_DecryptDataInit is available on this system.
func CanCallCSSM_DecryptDataInit() bool {
	return _cSSM_DecryptDataInit != nil
}

// CSSM_DecryptDataInitCallError returns the symbol lookup or registration error for CSSM_DecryptDataInit.
func CSSM_DecryptDataInitCallError() error {
	if _cSSM_DecryptDataInit != nil {
		return nil
	}
	return symbolCallError("CSSM_DecryptDataInit", "10.0", _cSSM_DecryptDataInitErr)
}

// TryCSSM_DecryptDataInit calls CSSM_DecryptDataInit without panicking when the symbol is unavailable.
func TryCSSM_DecryptDataInit(CCHandle CSSM_CC_HANDLE) (CSSM_RETURN, error) {
	if err := CSSM_DecryptDataInitCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_DecryptDataInit(CCHandle), nil
}

// CSSM_DecryptDataInit.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DecryptDataInit
func CSSM_DecryptDataInit(CCHandle CSSM_CC_HANDLE) CSSM_RETURN {
	result, callErr := TryCSSM_DecryptDataInit(CCHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DecryptDataInitP func(CCHandle CSSM_CC_HANDLE, Privilege CSSM_PRIVILEGE) CSSM_RETURN
var _cSSM_DecryptDataInitPErr error

// CanCallCSSM_DecryptDataInitP reports whether the symbol for CSSM_DecryptDataInitP is available on this system.
func CanCallCSSM_DecryptDataInitP() bool {
	return _cSSM_DecryptDataInitP != nil
}

// CSSM_DecryptDataInitPCallError returns the symbol lookup or registration error for CSSM_DecryptDataInitP.
func CSSM_DecryptDataInitPCallError() error {
	if _cSSM_DecryptDataInitP != nil {
		return nil
	}
	return symbolCallError("CSSM_DecryptDataInitP", "10.0", _cSSM_DecryptDataInitPErr)
}

// TryCSSM_DecryptDataInitP calls CSSM_DecryptDataInitP without panicking when the symbol is unavailable.
func TryCSSM_DecryptDataInitP(CCHandle CSSM_CC_HANDLE, Privilege CSSM_PRIVILEGE) (CSSM_RETURN, error) {
	if err := CSSM_DecryptDataInitPCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_DecryptDataInitP(CCHandle, Privilege), nil
}

// CSSM_DecryptDataInitP.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DecryptDataInitP
func CSSM_DecryptDataInitP(CCHandle CSSM_CC_HANDLE, Privilege CSSM_PRIVILEGE) CSSM_RETURN {
	result, callErr := TryCSSM_DecryptDataInitP(CCHandle, Privilege)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DecryptDataP func(CCHandle CSSM_CC_HANDLE, CipherBufs unsafe.Pointer, CipherBufCount uint32, ClearBufs unsafe.Pointer, ClearBufCount uint32, bytesDecrypted unsafe.Pointer, RemData unsafe.Pointer, Privilege CSSM_PRIVILEGE) CSSM_RETURN
var _cSSM_DecryptDataPErr error

// CanCallCSSM_DecryptDataP reports whether the symbol for CSSM_DecryptDataP is available on this system.
func CanCallCSSM_DecryptDataP() bool {
	return _cSSM_DecryptDataP != nil
}

// CSSM_DecryptDataPCallError returns the symbol lookup or registration error for CSSM_DecryptDataP.
func CSSM_DecryptDataPCallError() error {
	if _cSSM_DecryptDataP != nil {
		return nil
	}
	return symbolCallError("CSSM_DecryptDataP", "10.0", _cSSM_DecryptDataPErr)
}

// TryCSSM_DecryptDataP calls CSSM_DecryptDataP without panicking when the symbol is unavailable.
func TryCSSM_DecryptDataP(CCHandle CSSM_CC_HANDLE, CipherBufs unsafe.Pointer, CipherBufCount uint32, ClearBufs unsafe.Pointer, ClearBufCount uint32, bytesDecrypted unsafe.Pointer, RemData unsafe.Pointer, Privilege CSSM_PRIVILEGE) (CSSM_RETURN, error) {
	if err := CSSM_DecryptDataPCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_DecryptDataP(CCHandle, CipherBufs, CipherBufCount, ClearBufs, ClearBufCount, bytesDecrypted, RemData, Privilege), nil
}

// CSSM_DecryptDataP.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DecryptDataP
func CSSM_DecryptDataP(CCHandle CSSM_CC_HANDLE, CipherBufs unsafe.Pointer, CipherBufCount uint32, ClearBufs unsafe.Pointer, ClearBufCount uint32, bytesDecrypted unsafe.Pointer, RemData unsafe.Pointer, Privilege CSSM_PRIVILEGE) CSSM_RETURN {
	result, callErr := TryCSSM_DecryptDataP(CCHandle, CipherBufs, CipherBufCount, ClearBufs, ClearBufCount, bytesDecrypted, RemData, Privilege)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DecryptDataUpdate func(CCHandle CSSM_CC_HANDLE, CipherBufs unsafe.Pointer, CipherBufCount uint32, ClearBufs unsafe.Pointer, ClearBufCount uint32, bytesDecrypted unsafe.Pointer) CSSM_RETURN
var _cSSM_DecryptDataUpdateErr error

// CanCallCSSM_DecryptDataUpdate reports whether the symbol for CSSM_DecryptDataUpdate is available on this system.
func CanCallCSSM_DecryptDataUpdate() bool {
	return _cSSM_DecryptDataUpdate != nil
}

// CSSM_DecryptDataUpdateCallError returns the symbol lookup or registration error for CSSM_DecryptDataUpdate.
func CSSM_DecryptDataUpdateCallError() error {
	if _cSSM_DecryptDataUpdate != nil {
		return nil
	}
	return symbolCallError("CSSM_DecryptDataUpdate", "10.0", _cSSM_DecryptDataUpdateErr)
}

// TryCSSM_DecryptDataUpdate calls CSSM_DecryptDataUpdate without panicking when the symbol is unavailable.
func TryCSSM_DecryptDataUpdate(CCHandle CSSM_CC_HANDLE, CipherBufs unsafe.Pointer, CipherBufCount uint32, ClearBufs unsafe.Pointer, ClearBufCount uint32, bytesDecrypted unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_DecryptDataUpdateCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_DecryptDataUpdate(CCHandle, CipherBufs, CipherBufCount, ClearBufs, ClearBufCount, bytesDecrypted), nil
}

// CSSM_DecryptDataUpdate.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DecryptDataUpdate
func CSSM_DecryptDataUpdate(CCHandle CSSM_CC_HANDLE, CipherBufs unsafe.Pointer, CipherBufCount uint32, ClearBufs unsafe.Pointer, ClearBufCount uint32, bytesDecrypted unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_DecryptDataUpdate(CCHandle, CipherBufs, CipherBufCount, ClearBufs, ClearBufCount, bytesDecrypted)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DeleteContext func(CCHandle CSSM_CC_HANDLE) CSSM_RETURN
var _cSSM_DeleteContextErr error

// CanCallCSSM_DeleteContext reports whether the symbol for CSSM_DeleteContext is available on this system.
func CanCallCSSM_DeleteContext() bool {
	return _cSSM_DeleteContext != nil
}

// CSSM_DeleteContextCallError returns the symbol lookup or registration error for CSSM_DeleteContext.
func CSSM_DeleteContextCallError() error {
	if _cSSM_DeleteContext != nil {
		return nil
	}
	return symbolCallError("CSSM_DeleteContext", "10.0", _cSSM_DeleteContextErr)
}

// TryCSSM_DeleteContext calls CSSM_DeleteContext without panicking when the symbol is unavailable.
func TryCSSM_DeleteContext(CCHandle CSSM_CC_HANDLE) (CSSM_RETURN, error) {
	if err := CSSM_DeleteContextCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_DeleteContext(CCHandle), nil
}

// CSSM_DeleteContext.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DeleteContext
func CSSM_DeleteContext(CCHandle CSSM_CC_HANDLE) CSSM_RETURN {
	result, callErr := TryCSSM_DeleteContext(CCHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DeleteContextAttributes func(CCHandle CSSM_CC_HANDLE, NumberOfAttributes uint32, ContextAttributes unsafe.Pointer) CSSM_RETURN
var _cSSM_DeleteContextAttributesErr error

// CanCallCSSM_DeleteContextAttributes reports whether the symbol for CSSM_DeleteContextAttributes is available on this system.
func CanCallCSSM_DeleteContextAttributes() bool {
	return _cSSM_DeleteContextAttributes != nil
}

// CSSM_DeleteContextAttributesCallError returns the symbol lookup or registration error for CSSM_DeleteContextAttributes.
func CSSM_DeleteContextAttributesCallError() error {
	if _cSSM_DeleteContextAttributes != nil {
		return nil
	}
	return symbolCallError("CSSM_DeleteContextAttributes", "10.0", _cSSM_DeleteContextAttributesErr)
}

// TryCSSM_DeleteContextAttributes calls CSSM_DeleteContextAttributes without panicking when the symbol is unavailable.
func TryCSSM_DeleteContextAttributes(CCHandle CSSM_CC_HANDLE, NumberOfAttributes uint32, ContextAttributes unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_DeleteContextAttributesCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_DeleteContextAttributes(CCHandle, NumberOfAttributes, ContextAttributes), nil
}

// CSSM_DeleteContextAttributes.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DeleteContextAttributes
func CSSM_DeleteContextAttributes(CCHandle CSSM_CC_HANDLE, NumberOfAttributes uint32, ContextAttributes unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_DeleteContextAttributes(CCHandle, NumberOfAttributes, ContextAttributes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DeriveKey func(CCHandle CSSM_CC_HANDLE, Param unsafe.Pointer, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, DerivedKey unsafe.Pointer) CSSM_RETURN
var _cSSM_DeriveKeyErr error

// CanCallCSSM_DeriveKey reports whether the symbol for CSSM_DeriveKey is available on this system.
func CanCallCSSM_DeriveKey() bool {
	return _cSSM_DeriveKey != nil
}

// CSSM_DeriveKeyCallError returns the symbol lookup or registration error for CSSM_DeriveKey.
func CSSM_DeriveKeyCallError() error {
	if _cSSM_DeriveKey != nil {
		return nil
	}
	return symbolCallError("CSSM_DeriveKey", "10.0", _cSSM_DeriveKeyErr)
}

// TryCSSM_DeriveKey calls CSSM_DeriveKey without panicking when the symbol is unavailable.
func TryCSSM_DeriveKey(CCHandle CSSM_CC_HANDLE, Param unsafe.Pointer, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, DerivedKey unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_DeriveKeyCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_DeriveKey(CCHandle, Param, KeyUsage, KeyAttr, KeyLabel, CredAndAclEntry, DerivedKey), nil
}

// CSSM_DeriveKey.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DeriveKey
func CSSM_DeriveKey(CCHandle CSSM_CC_HANDLE, Param unsafe.Pointer, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, DerivedKey unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_DeriveKey(CCHandle, Param, KeyUsage, KeyAttr, KeyLabel, CredAndAclEntry, DerivedKey)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DigestData func(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, Digest unsafe.Pointer) CSSM_RETURN
var _cSSM_DigestDataErr error

// CanCallCSSM_DigestData reports whether the symbol for CSSM_DigestData is available on this system.
func CanCallCSSM_DigestData() bool {
	return _cSSM_DigestData != nil
}

// CSSM_DigestDataCallError returns the symbol lookup or registration error for CSSM_DigestData.
func CSSM_DigestDataCallError() error {
	if _cSSM_DigestData != nil {
		return nil
	}
	return symbolCallError("CSSM_DigestData", "10.0", _cSSM_DigestDataErr)
}

// TryCSSM_DigestData calls CSSM_DigestData without panicking when the symbol is unavailable.
func TryCSSM_DigestData(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, Digest unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_DigestDataCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_DigestData(CCHandle, DataBufs, DataBufCount, Digest), nil
}

// CSSM_DigestData.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DigestData
func CSSM_DigestData(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, Digest unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_DigestData(CCHandle, DataBufs, DataBufCount, Digest)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DigestDataClone func(CCHandle CSSM_CC_HANDLE, ClonednewCCHandle unsafe.Pointer) CSSM_RETURN
var _cSSM_DigestDataCloneErr error

// CanCallCSSM_DigestDataClone reports whether the symbol for CSSM_DigestDataClone is available on this system.
func CanCallCSSM_DigestDataClone() bool {
	return _cSSM_DigestDataClone != nil
}

// CSSM_DigestDataCloneCallError returns the symbol lookup or registration error for CSSM_DigestDataClone.
func CSSM_DigestDataCloneCallError() error {
	if _cSSM_DigestDataClone != nil {
		return nil
	}
	return symbolCallError("CSSM_DigestDataClone", "10.0", _cSSM_DigestDataCloneErr)
}

// TryCSSM_DigestDataClone calls CSSM_DigestDataClone without panicking when the symbol is unavailable.
func TryCSSM_DigestDataClone(CCHandle CSSM_CC_HANDLE, ClonednewCCHandle unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_DigestDataCloneCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_DigestDataClone(CCHandle, ClonednewCCHandle), nil
}

// CSSM_DigestDataClone.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DigestDataClone
func CSSM_DigestDataClone(CCHandle CSSM_CC_HANDLE, ClonednewCCHandle unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_DigestDataClone(CCHandle, ClonednewCCHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DigestDataFinal func(CCHandle CSSM_CC_HANDLE, Digest unsafe.Pointer) CSSM_RETURN
var _cSSM_DigestDataFinalErr error

// CanCallCSSM_DigestDataFinal reports whether the symbol for CSSM_DigestDataFinal is available on this system.
func CanCallCSSM_DigestDataFinal() bool {
	return _cSSM_DigestDataFinal != nil
}

// CSSM_DigestDataFinalCallError returns the symbol lookup or registration error for CSSM_DigestDataFinal.
func CSSM_DigestDataFinalCallError() error {
	if _cSSM_DigestDataFinal != nil {
		return nil
	}
	return symbolCallError("CSSM_DigestDataFinal", "10.0", _cSSM_DigestDataFinalErr)
}

// TryCSSM_DigestDataFinal calls CSSM_DigestDataFinal without panicking when the symbol is unavailable.
func TryCSSM_DigestDataFinal(CCHandle CSSM_CC_HANDLE, Digest unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_DigestDataFinalCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_DigestDataFinal(CCHandle, Digest), nil
}

// CSSM_DigestDataFinal.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DigestDataFinal
func CSSM_DigestDataFinal(CCHandle CSSM_CC_HANDLE, Digest unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_DigestDataFinal(CCHandle, Digest)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DigestDataInit func(CCHandle CSSM_CC_HANDLE) CSSM_RETURN
var _cSSM_DigestDataInitErr error

// CanCallCSSM_DigestDataInit reports whether the symbol for CSSM_DigestDataInit is available on this system.
func CanCallCSSM_DigestDataInit() bool {
	return _cSSM_DigestDataInit != nil
}

// CSSM_DigestDataInitCallError returns the symbol lookup or registration error for CSSM_DigestDataInit.
func CSSM_DigestDataInitCallError() error {
	if _cSSM_DigestDataInit != nil {
		return nil
	}
	return symbolCallError("CSSM_DigestDataInit", "10.0", _cSSM_DigestDataInitErr)
}

// TryCSSM_DigestDataInit calls CSSM_DigestDataInit without panicking when the symbol is unavailable.
func TryCSSM_DigestDataInit(CCHandle CSSM_CC_HANDLE) (CSSM_RETURN, error) {
	if err := CSSM_DigestDataInitCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_DigestDataInit(CCHandle), nil
}

// CSSM_DigestDataInit.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DigestDataInit
func CSSM_DigestDataInit(CCHandle CSSM_CC_HANDLE) CSSM_RETURN {
	result, callErr := TryCSSM_DigestDataInit(CCHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_DigestDataUpdate func(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) CSSM_RETURN
var _cSSM_DigestDataUpdateErr error

// CanCallCSSM_DigestDataUpdate reports whether the symbol for CSSM_DigestDataUpdate is available on this system.
func CanCallCSSM_DigestDataUpdate() bool {
	return _cSSM_DigestDataUpdate != nil
}

// CSSM_DigestDataUpdateCallError returns the symbol lookup or registration error for CSSM_DigestDataUpdate.
func CSSM_DigestDataUpdateCallError() error {
	if _cSSM_DigestDataUpdate != nil {
		return nil
	}
	return symbolCallError("CSSM_DigestDataUpdate", "10.0", _cSSM_DigestDataUpdateErr)
}

// TryCSSM_DigestDataUpdate calls CSSM_DigestDataUpdate without panicking when the symbol is unavailable.
func TryCSSM_DigestDataUpdate(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) (CSSM_RETURN, error) {
	if err := CSSM_DigestDataUpdateCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_DigestDataUpdate(CCHandle, DataBufs, DataBufCount), nil
}

// CSSM_DigestDataUpdate.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DigestDataUpdate
func CSSM_DigestDataUpdate(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) CSSM_RETURN {
	result, callErr := TryCSSM_DigestDataUpdate(CCHandle, DataBufs, DataBufCount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_EncryptData func(CCHandle CSSM_CC_HANDLE, ClearBufs unsafe.Pointer, ClearBufCount uint32, CipherBufs unsafe.Pointer, CipherBufCount uint32, bytesEncrypted unsafe.Pointer, RemData unsafe.Pointer) CSSM_RETURN
var _cSSM_EncryptDataErr error

// CanCallCSSM_EncryptData reports whether the symbol for CSSM_EncryptData is available on this system.
func CanCallCSSM_EncryptData() bool {
	return _cSSM_EncryptData != nil
}

// CSSM_EncryptDataCallError returns the symbol lookup or registration error for CSSM_EncryptData.
func CSSM_EncryptDataCallError() error {
	if _cSSM_EncryptData != nil {
		return nil
	}
	return symbolCallError("CSSM_EncryptData", "10.0", _cSSM_EncryptDataErr)
}

// TryCSSM_EncryptData calls CSSM_EncryptData without panicking when the symbol is unavailable.
func TryCSSM_EncryptData(CCHandle CSSM_CC_HANDLE, ClearBufs unsafe.Pointer, ClearBufCount uint32, CipherBufs unsafe.Pointer, CipherBufCount uint32, bytesEncrypted unsafe.Pointer, RemData unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_EncryptDataCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_EncryptData(CCHandle, ClearBufs, ClearBufCount, CipherBufs, CipherBufCount, bytesEncrypted, RemData), nil
}

// CSSM_EncryptData.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_EncryptData
func CSSM_EncryptData(CCHandle CSSM_CC_HANDLE, ClearBufs unsafe.Pointer, ClearBufCount uint32, CipherBufs unsafe.Pointer, CipherBufCount uint32, bytesEncrypted unsafe.Pointer, RemData unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_EncryptData(CCHandle, ClearBufs, ClearBufCount, CipherBufs, CipherBufCount, bytesEncrypted, RemData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_EncryptDataFinal func(CCHandle CSSM_CC_HANDLE, RemData unsafe.Pointer) CSSM_RETURN
var _cSSM_EncryptDataFinalErr error

// CanCallCSSM_EncryptDataFinal reports whether the symbol for CSSM_EncryptDataFinal is available on this system.
func CanCallCSSM_EncryptDataFinal() bool {
	return _cSSM_EncryptDataFinal != nil
}

// CSSM_EncryptDataFinalCallError returns the symbol lookup or registration error for CSSM_EncryptDataFinal.
func CSSM_EncryptDataFinalCallError() error {
	if _cSSM_EncryptDataFinal != nil {
		return nil
	}
	return symbolCallError("CSSM_EncryptDataFinal", "10.0", _cSSM_EncryptDataFinalErr)
}

// TryCSSM_EncryptDataFinal calls CSSM_EncryptDataFinal without panicking when the symbol is unavailable.
func TryCSSM_EncryptDataFinal(CCHandle CSSM_CC_HANDLE, RemData unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_EncryptDataFinalCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_EncryptDataFinal(CCHandle, RemData), nil
}

// CSSM_EncryptDataFinal.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_EncryptDataFinal
func CSSM_EncryptDataFinal(CCHandle CSSM_CC_HANDLE, RemData unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_EncryptDataFinal(CCHandle, RemData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_EncryptDataInit func(CCHandle CSSM_CC_HANDLE) CSSM_RETURN
var _cSSM_EncryptDataInitErr error

// CanCallCSSM_EncryptDataInit reports whether the symbol for CSSM_EncryptDataInit is available on this system.
func CanCallCSSM_EncryptDataInit() bool {
	return _cSSM_EncryptDataInit != nil
}

// CSSM_EncryptDataInitCallError returns the symbol lookup or registration error for CSSM_EncryptDataInit.
func CSSM_EncryptDataInitCallError() error {
	if _cSSM_EncryptDataInit != nil {
		return nil
	}
	return symbolCallError("CSSM_EncryptDataInit", "10.0", _cSSM_EncryptDataInitErr)
}

// TryCSSM_EncryptDataInit calls CSSM_EncryptDataInit without panicking when the symbol is unavailable.
func TryCSSM_EncryptDataInit(CCHandle CSSM_CC_HANDLE) (CSSM_RETURN, error) {
	if err := CSSM_EncryptDataInitCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_EncryptDataInit(CCHandle), nil
}

// CSSM_EncryptDataInit.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_EncryptDataInit
func CSSM_EncryptDataInit(CCHandle CSSM_CC_HANDLE) CSSM_RETURN {
	result, callErr := TryCSSM_EncryptDataInit(CCHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_EncryptDataInitP func(CCHandle CSSM_CC_HANDLE, Privilege CSSM_PRIVILEGE) CSSM_RETURN
var _cSSM_EncryptDataInitPErr error

// CanCallCSSM_EncryptDataInitP reports whether the symbol for CSSM_EncryptDataInitP is available on this system.
func CanCallCSSM_EncryptDataInitP() bool {
	return _cSSM_EncryptDataInitP != nil
}

// CSSM_EncryptDataInitPCallError returns the symbol lookup or registration error for CSSM_EncryptDataInitP.
func CSSM_EncryptDataInitPCallError() error {
	if _cSSM_EncryptDataInitP != nil {
		return nil
	}
	return symbolCallError("CSSM_EncryptDataInitP", "10.0", _cSSM_EncryptDataInitPErr)
}

// TryCSSM_EncryptDataInitP calls CSSM_EncryptDataInitP without panicking when the symbol is unavailable.
func TryCSSM_EncryptDataInitP(CCHandle CSSM_CC_HANDLE, Privilege CSSM_PRIVILEGE) (CSSM_RETURN, error) {
	if err := CSSM_EncryptDataInitPCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_EncryptDataInitP(CCHandle, Privilege), nil
}

// CSSM_EncryptDataInitP.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_EncryptDataInitP
func CSSM_EncryptDataInitP(CCHandle CSSM_CC_HANDLE, Privilege CSSM_PRIVILEGE) CSSM_RETURN {
	result, callErr := TryCSSM_EncryptDataInitP(CCHandle, Privilege)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_EncryptDataP func(CCHandle CSSM_CC_HANDLE, ClearBufs unsafe.Pointer, ClearBufCount uint32, CipherBufs unsafe.Pointer, CipherBufCount uint32, bytesEncrypted unsafe.Pointer, RemData unsafe.Pointer, Privilege CSSM_PRIVILEGE) CSSM_RETURN
var _cSSM_EncryptDataPErr error

// CanCallCSSM_EncryptDataP reports whether the symbol for CSSM_EncryptDataP is available on this system.
func CanCallCSSM_EncryptDataP() bool {
	return _cSSM_EncryptDataP != nil
}

// CSSM_EncryptDataPCallError returns the symbol lookup or registration error for CSSM_EncryptDataP.
func CSSM_EncryptDataPCallError() error {
	if _cSSM_EncryptDataP != nil {
		return nil
	}
	return symbolCallError("CSSM_EncryptDataP", "10.0", _cSSM_EncryptDataPErr)
}

// TryCSSM_EncryptDataP calls CSSM_EncryptDataP without panicking when the symbol is unavailable.
func TryCSSM_EncryptDataP(CCHandle CSSM_CC_HANDLE, ClearBufs unsafe.Pointer, ClearBufCount uint32, CipherBufs unsafe.Pointer, CipherBufCount uint32, bytesEncrypted unsafe.Pointer, RemData unsafe.Pointer, Privilege CSSM_PRIVILEGE) (CSSM_RETURN, error) {
	if err := CSSM_EncryptDataPCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_EncryptDataP(CCHandle, ClearBufs, ClearBufCount, CipherBufs, CipherBufCount, bytesEncrypted, RemData, Privilege), nil
}

// CSSM_EncryptDataP.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_EncryptDataP
func CSSM_EncryptDataP(CCHandle CSSM_CC_HANDLE, ClearBufs unsafe.Pointer, ClearBufCount uint32, CipherBufs unsafe.Pointer, CipherBufCount uint32, bytesEncrypted unsafe.Pointer, RemData unsafe.Pointer, Privilege CSSM_PRIVILEGE) CSSM_RETURN {
	result, callErr := TryCSSM_EncryptDataP(CCHandle, ClearBufs, ClearBufCount, CipherBufs, CipherBufCount, bytesEncrypted, RemData, Privilege)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_EncryptDataUpdate func(CCHandle CSSM_CC_HANDLE, ClearBufs unsafe.Pointer, ClearBufCount uint32, CipherBufs unsafe.Pointer, CipherBufCount uint32, bytesEncrypted unsafe.Pointer) CSSM_RETURN
var _cSSM_EncryptDataUpdateErr error

// CanCallCSSM_EncryptDataUpdate reports whether the symbol for CSSM_EncryptDataUpdate is available on this system.
func CanCallCSSM_EncryptDataUpdate() bool {
	return _cSSM_EncryptDataUpdate != nil
}

// CSSM_EncryptDataUpdateCallError returns the symbol lookup or registration error for CSSM_EncryptDataUpdate.
func CSSM_EncryptDataUpdateCallError() error {
	if _cSSM_EncryptDataUpdate != nil {
		return nil
	}
	return symbolCallError("CSSM_EncryptDataUpdate", "10.0", _cSSM_EncryptDataUpdateErr)
}

// TryCSSM_EncryptDataUpdate calls CSSM_EncryptDataUpdate without panicking when the symbol is unavailable.
func TryCSSM_EncryptDataUpdate(CCHandle CSSM_CC_HANDLE, ClearBufs unsafe.Pointer, ClearBufCount uint32, CipherBufs unsafe.Pointer, CipherBufCount uint32, bytesEncrypted unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_EncryptDataUpdateCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_EncryptDataUpdate(CCHandle, ClearBufs, ClearBufCount, CipherBufs, CipherBufCount, bytesEncrypted), nil
}

// CSSM_EncryptDataUpdate.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_EncryptDataUpdate
func CSSM_EncryptDataUpdate(CCHandle CSSM_CC_HANDLE, ClearBufs unsafe.Pointer, ClearBufCount uint32, CipherBufs unsafe.Pointer, CipherBufCount uint32, bytesEncrypted unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_EncryptDataUpdate(CCHandle, ClearBufs, ClearBufCount, CipherBufs, CipherBufCount, bytesEncrypted)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_FreeContext func(Context unsafe.Pointer) CSSM_RETURN
var _cSSM_FreeContextErr error

// CanCallCSSM_FreeContext reports whether the symbol for CSSM_FreeContext is available on this system.
func CanCallCSSM_FreeContext() bool {
	return _cSSM_FreeContext != nil
}

// CSSM_FreeContextCallError returns the symbol lookup or registration error for CSSM_FreeContext.
func CSSM_FreeContextCallError() error {
	if _cSSM_FreeContext != nil {
		return nil
	}
	return symbolCallError("CSSM_FreeContext", "10.0", _cSSM_FreeContextErr)
}

// TryCSSM_FreeContext calls CSSM_FreeContext without panicking when the symbol is unavailable.
func TryCSSM_FreeContext(Context unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_FreeContextCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_FreeContext(Context), nil
}

// CSSM_FreeContext.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_FreeContext
func CSSM_FreeContext(Context unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_FreeContext(Context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_FreeKey func(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, KeyPtr unsafe.Pointer, Delete CSSM_BOOL) CSSM_RETURN
var _cSSM_FreeKeyErr error

// CanCallCSSM_FreeKey reports whether the symbol for CSSM_FreeKey is available on this system.
func CanCallCSSM_FreeKey() bool {
	return _cSSM_FreeKey != nil
}

// CSSM_FreeKeyCallError returns the symbol lookup or registration error for CSSM_FreeKey.
func CSSM_FreeKeyCallError() error {
	if _cSSM_FreeKey != nil {
		return nil
	}
	return symbolCallError("CSSM_FreeKey", "10.0", _cSSM_FreeKeyErr)
}

// TryCSSM_FreeKey calls CSSM_FreeKey without panicking when the symbol is unavailable.
func TryCSSM_FreeKey(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, KeyPtr unsafe.Pointer, Delete CSSM_BOOL) (CSSM_RETURN, error) {
	if err := CSSM_FreeKeyCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_FreeKey(CSPHandle, AccessCred, KeyPtr, Delete), nil
}

// CSSM_FreeKey.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_FreeKey
func CSSM_FreeKey(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, KeyPtr unsafe.Pointer, Delete CSSM_BOOL) CSSM_RETURN {
	result, callErr := TryCSSM_FreeKey(CSPHandle, AccessCred, KeyPtr, Delete)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_GenerateAlgorithmParams func(CCHandle CSSM_CC_HANDLE, ParamBits uint32, Param unsafe.Pointer) CSSM_RETURN
var _cSSM_GenerateAlgorithmParamsErr error

// CanCallCSSM_GenerateAlgorithmParams reports whether the symbol for CSSM_GenerateAlgorithmParams is available on this system.
func CanCallCSSM_GenerateAlgorithmParams() bool {
	return _cSSM_GenerateAlgorithmParams != nil
}

// CSSM_GenerateAlgorithmParamsCallError returns the symbol lookup or registration error for CSSM_GenerateAlgorithmParams.
func CSSM_GenerateAlgorithmParamsCallError() error {
	if _cSSM_GenerateAlgorithmParams != nil {
		return nil
	}
	return symbolCallError("CSSM_GenerateAlgorithmParams", "10.0", _cSSM_GenerateAlgorithmParamsErr)
}

// TryCSSM_GenerateAlgorithmParams calls CSSM_GenerateAlgorithmParams without panicking when the symbol is unavailable.
func TryCSSM_GenerateAlgorithmParams(CCHandle CSSM_CC_HANDLE, ParamBits uint32, Param unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_GenerateAlgorithmParamsCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_GenerateAlgorithmParams(CCHandle, ParamBits, Param), nil
}

// CSSM_GenerateAlgorithmParams.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GenerateAlgorithmParams
func CSSM_GenerateAlgorithmParams(CCHandle CSSM_CC_HANDLE, ParamBits uint32, Param unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_GenerateAlgorithmParams(CCHandle, ParamBits, Param)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_GenerateKey func(CCHandle CSSM_CC_HANDLE, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, Key unsafe.Pointer) CSSM_RETURN
var _cSSM_GenerateKeyErr error

// CanCallCSSM_GenerateKey reports whether the symbol for CSSM_GenerateKey is available on this system.
func CanCallCSSM_GenerateKey() bool {
	return _cSSM_GenerateKey != nil
}

// CSSM_GenerateKeyCallError returns the symbol lookup or registration error for CSSM_GenerateKey.
func CSSM_GenerateKeyCallError() error {
	if _cSSM_GenerateKey != nil {
		return nil
	}
	return symbolCallError("CSSM_GenerateKey", "10.0", _cSSM_GenerateKeyErr)
}

// TryCSSM_GenerateKey calls CSSM_GenerateKey without panicking when the symbol is unavailable.
func TryCSSM_GenerateKey(CCHandle CSSM_CC_HANDLE, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, Key unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_GenerateKeyCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_GenerateKey(CCHandle, KeyUsage, KeyAttr, KeyLabel, CredAndAclEntry, Key), nil
}

// CSSM_GenerateKey.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GenerateKey
func CSSM_GenerateKey(CCHandle CSSM_CC_HANDLE, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, Key unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_GenerateKey(CCHandle, KeyUsage, KeyAttr, KeyLabel, CredAndAclEntry, Key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_GenerateKeyP func(CCHandle CSSM_CC_HANDLE, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, Key unsafe.Pointer, Privilege CSSM_PRIVILEGE) CSSM_RETURN
var _cSSM_GenerateKeyPErr error

// CanCallCSSM_GenerateKeyP reports whether the symbol for CSSM_GenerateKeyP is available on this system.
func CanCallCSSM_GenerateKeyP() bool {
	return _cSSM_GenerateKeyP != nil
}

// CSSM_GenerateKeyPCallError returns the symbol lookup or registration error for CSSM_GenerateKeyP.
func CSSM_GenerateKeyPCallError() error {
	if _cSSM_GenerateKeyP != nil {
		return nil
	}
	return symbolCallError("CSSM_GenerateKeyP", "10.0", _cSSM_GenerateKeyPErr)
}

// TryCSSM_GenerateKeyP calls CSSM_GenerateKeyP without panicking when the symbol is unavailable.
func TryCSSM_GenerateKeyP(CCHandle CSSM_CC_HANDLE, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, Key unsafe.Pointer, Privilege CSSM_PRIVILEGE) (CSSM_RETURN, error) {
	if err := CSSM_GenerateKeyPCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_GenerateKeyP(CCHandle, KeyUsage, KeyAttr, KeyLabel, CredAndAclEntry, Key, Privilege), nil
}

// CSSM_GenerateKeyP.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GenerateKeyP
func CSSM_GenerateKeyP(CCHandle CSSM_CC_HANDLE, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, Key unsafe.Pointer, Privilege CSSM_PRIVILEGE) CSSM_RETURN {
	result, callErr := TryCSSM_GenerateKeyP(CCHandle, KeyUsage, KeyAttr, KeyLabel, CredAndAclEntry, Key, Privilege)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_GenerateKeyPair func(CCHandle CSSM_CC_HANDLE, PublicKeyUsage uint32, PublicKeyAttr uint32, PublicKeyLabel unsafe.Pointer, PublicKey unsafe.Pointer, PrivateKeyUsage uint32, PrivateKeyAttr uint32, PrivateKeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, PrivateKey unsafe.Pointer) CSSM_RETURN
var _cSSM_GenerateKeyPairErr error

// CanCallCSSM_GenerateKeyPair reports whether the symbol for CSSM_GenerateKeyPair is available on this system.
func CanCallCSSM_GenerateKeyPair() bool {
	return _cSSM_GenerateKeyPair != nil
}

// CSSM_GenerateKeyPairCallError returns the symbol lookup or registration error for CSSM_GenerateKeyPair.
func CSSM_GenerateKeyPairCallError() error {
	if _cSSM_GenerateKeyPair != nil {
		return nil
	}
	return symbolCallError("CSSM_GenerateKeyPair", "10.0", _cSSM_GenerateKeyPairErr)
}

// TryCSSM_GenerateKeyPair calls CSSM_GenerateKeyPair without panicking when the symbol is unavailable.
func TryCSSM_GenerateKeyPair(CCHandle CSSM_CC_HANDLE, PublicKeyUsage uint32, PublicKeyAttr uint32, PublicKeyLabel unsafe.Pointer, PublicKey unsafe.Pointer, PrivateKeyUsage uint32, PrivateKeyAttr uint32, PrivateKeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, PrivateKey unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_GenerateKeyPairCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_GenerateKeyPair(CCHandle, PublicKeyUsage, PublicKeyAttr, PublicKeyLabel, PublicKey, PrivateKeyUsage, PrivateKeyAttr, PrivateKeyLabel, CredAndAclEntry, PrivateKey), nil
}

// CSSM_GenerateKeyPair.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GenerateKeyPair
func CSSM_GenerateKeyPair(CCHandle CSSM_CC_HANDLE, PublicKeyUsage uint32, PublicKeyAttr uint32, PublicKeyLabel unsafe.Pointer, PublicKey unsafe.Pointer, PrivateKeyUsage uint32, PrivateKeyAttr uint32, PrivateKeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, PrivateKey unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_GenerateKeyPair(CCHandle, PublicKeyUsage, PublicKeyAttr, PublicKeyLabel, PublicKey, PrivateKeyUsage, PrivateKeyAttr, PrivateKeyLabel, CredAndAclEntry, PrivateKey)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_GenerateKeyPairP func(CCHandle CSSM_CC_HANDLE, PublicKeyUsage uint32, PublicKeyAttr uint32, PublicKeyLabel unsafe.Pointer, PublicKey unsafe.Pointer, PrivateKeyUsage uint32, PrivateKeyAttr uint32, PrivateKeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, PrivateKey unsafe.Pointer, Privilege CSSM_PRIVILEGE) CSSM_RETURN
var _cSSM_GenerateKeyPairPErr error

// CanCallCSSM_GenerateKeyPairP reports whether the symbol for CSSM_GenerateKeyPairP is available on this system.
func CanCallCSSM_GenerateKeyPairP() bool {
	return _cSSM_GenerateKeyPairP != nil
}

// CSSM_GenerateKeyPairPCallError returns the symbol lookup or registration error for CSSM_GenerateKeyPairP.
func CSSM_GenerateKeyPairPCallError() error {
	if _cSSM_GenerateKeyPairP != nil {
		return nil
	}
	return symbolCallError("CSSM_GenerateKeyPairP", "10.0", _cSSM_GenerateKeyPairPErr)
}

// TryCSSM_GenerateKeyPairP calls CSSM_GenerateKeyPairP without panicking when the symbol is unavailable.
func TryCSSM_GenerateKeyPairP(CCHandle CSSM_CC_HANDLE, PublicKeyUsage uint32, PublicKeyAttr uint32, PublicKeyLabel unsafe.Pointer, PublicKey unsafe.Pointer, PrivateKeyUsage uint32, PrivateKeyAttr uint32, PrivateKeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, PrivateKey unsafe.Pointer, Privilege CSSM_PRIVILEGE) (CSSM_RETURN, error) {
	if err := CSSM_GenerateKeyPairPCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_GenerateKeyPairP(CCHandle, PublicKeyUsage, PublicKeyAttr, PublicKeyLabel, PublicKey, PrivateKeyUsage, PrivateKeyAttr, PrivateKeyLabel, CredAndAclEntry, PrivateKey, Privilege), nil
}

// CSSM_GenerateKeyPairP.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GenerateKeyPairP
func CSSM_GenerateKeyPairP(CCHandle CSSM_CC_HANDLE, PublicKeyUsage uint32, PublicKeyAttr uint32, PublicKeyLabel unsafe.Pointer, PublicKey unsafe.Pointer, PrivateKeyUsage uint32, PrivateKeyAttr uint32, PrivateKeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, PrivateKey unsafe.Pointer, Privilege CSSM_PRIVILEGE) CSSM_RETURN {
	result, callErr := TryCSSM_GenerateKeyPairP(CCHandle, PublicKeyUsage, PublicKeyAttr, PublicKeyLabel, PublicKey, PrivateKeyUsage, PrivateKeyAttr, PrivateKeyLabel, CredAndAclEntry, PrivateKey, Privilege)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_GenerateMac func(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, Mac unsafe.Pointer) CSSM_RETURN
var _cSSM_GenerateMacErr error

// CanCallCSSM_GenerateMac reports whether the symbol for CSSM_GenerateMac is available on this system.
func CanCallCSSM_GenerateMac() bool {
	return _cSSM_GenerateMac != nil
}

// CSSM_GenerateMacCallError returns the symbol lookup or registration error for CSSM_GenerateMac.
func CSSM_GenerateMacCallError() error {
	if _cSSM_GenerateMac != nil {
		return nil
	}
	return symbolCallError("CSSM_GenerateMac", "10.0", _cSSM_GenerateMacErr)
}

// TryCSSM_GenerateMac calls CSSM_GenerateMac without panicking when the symbol is unavailable.
func TryCSSM_GenerateMac(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, Mac unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_GenerateMacCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_GenerateMac(CCHandle, DataBufs, DataBufCount, Mac), nil
}

// CSSM_GenerateMac.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GenerateMac
func CSSM_GenerateMac(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, Mac unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_GenerateMac(CCHandle, DataBufs, DataBufCount, Mac)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_GenerateMacFinal func(CCHandle CSSM_CC_HANDLE, Mac unsafe.Pointer) CSSM_RETURN
var _cSSM_GenerateMacFinalErr error

// CanCallCSSM_GenerateMacFinal reports whether the symbol for CSSM_GenerateMacFinal is available on this system.
func CanCallCSSM_GenerateMacFinal() bool {
	return _cSSM_GenerateMacFinal != nil
}

// CSSM_GenerateMacFinalCallError returns the symbol lookup or registration error for CSSM_GenerateMacFinal.
func CSSM_GenerateMacFinalCallError() error {
	if _cSSM_GenerateMacFinal != nil {
		return nil
	}
	return symbolCallError("CSSM_GenerateMacFinal", "10.0", _cSSM_GenerateMacFinalErr)
}

// TryCSSM_GenerateMacFinal calls CSSM_GenerateMacFinal without panicking when the symbol is unavailable.
func TryCSSM_GenerateMacFinal(CCHandle CSSM_CC_HANDLE, Mac unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_GenerateMacFinalCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_GenerateMacFinal(CCHandle, Mac), nil
}

// CSSM_GenerateMacFinal.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GenerateMacFinal
func CSSM_GenerateMacFinal(CCHandle CSSM_CC_HANDLE, Mac unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_GenerateMacFinal(CCHandle, Mac)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_GenerateMacInit func(CCHandle CSSM_CC_HANDLE) CSSM_RETURN
var _cSSM_GenerateMacInitErr error

// CanCallCSSM_GenerateMacInit reports whether the symbol for CSSM_GenerateMacInit is available on this system.
func CanCallCSSM_GenerateMacInit() bool {
	return _cSSM_GenerateMacInit != nil
}

// CSSM_GenerateMacInitCallError returns the symbol lookup or registration error for CSSM_GenerateMacInit.
func CSSM_GenerateMacInitCallError() error {
	if _cSSM_GenerateMacInit != nil {
		return nil
	}
	return symbolCallError("CSSM_GenerateMacInit", "10.0", _cSSM_GenerateMacInitErr)
}

// TryCSSM_GenerateMacInit calls CSSM_GenerateMacInit without panicking when the symbol is unavailable.
func TryCSSM_GenerateMacInit(CCHandle CSSM_CC_HANDLE) (CSSM_RETURN, error) {
	if err := CSSM_GenerateMacInitCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_GenerateMacInit(CCHandle), nil
}

// CSSM_GenerateMacInit.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GenerateMacInit
func CSSM_GenerateMacInit(CCHandle CSSM_CC_HANDLE) CSSM_RETURN {
	result, callErr := TryCSSM_GenerateMacInit(CCHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_GenerateMacUpdate func(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) CSSM_RETURN
var _cSSM_GenerateMacUpdateErr error

// CanCallCSSM_GenerateMacUpdate reports whether the symbol for CSSM_GenerateMacUpdate is available on this system.
func CanCallCSSM_GenerateMacUpdate() bool {
	return _cSSM_GenerateMacUpdate != nil
}

// CSSM_GenerateMacUpdateCallError returns the symbol lookup or registration error for CSSM_GenerateMacUpdate.
func CSSM_GenerateMacUpdateCallError() error {
	if _cSSM_GenerateMacUpdate != nil {
		return nil
	}
	return symbolCallError("CSSM_GenerateMacUpdate", "10.0", _cSSM_GenerateMacUpdateErr)
}

// TryCSSM_GenerateMacUpdate calls CSSM_GenerateMacUpdate without panicking when the symbol is unavailable.
func TryCSSM_GenerateMacUpdate(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) (CSSM_RETURN, error) {
	if err := CSSM_GenerateMacUpdateCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_GenerateMacUpdate(CCHandle, DataBufs, DataBufCount), nil
}

// CSSM_GenerateMacUpdate.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GenerateMacUpdate
func CSSM_GenerateMacUpdate(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) CSSM_RETURN {
	result, callErr := TryCSSM_GenerateMacUpdate(CCHandle, DataBufs, DataBufCount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_GenerateRandom func(CCHandle CSSM_CC_HANDLE, RandomNumber unsafe.Pointer) CSSM_RETURN
var _cSSM_GenerateRandomErr error

// CanCallCSSM_GenerateRandom reports whether the symbol for CSSM_GenerateRandom is available on this system.
func CanCallCSSM_GenerateRandom() bool {
	return _cSSM_GenerateRandom != nil
}

// CSSM_GenerateRandomCallError returns the symbol lookup or registration error for CSSM_GenerateRandom.
func CSSM_GenerateRandomCallError() error {
	if _cSSM_GenerateRandom != nil {
		return nil
	}
	return symbolCallError("CSSM_GenerateRandom", "10.0", _cSSM_GenerateRandomErr)
}

// TryCSSM_GenerateRandom calls CSSM_GenerateRandom without panicking when the symbol is unavailable.
func TryCSSM_GenerateRandom(CCHandle CSSM_CC_HANDLE, RandomNumber unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_GenerateRandomCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_GenerateRandom(CCHandle, RandomNumber), nil
}

// CSSM_GenerateRandom.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GenerateRandom
func CSSM_GenerateRandom(CCHandle CSSM_CC_HANDLE, RandomNumber unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_GenerateRandom(CCHandle, RandomNumber)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_GetAPIMemoryFunctions func(AddInHandle CSSM_MODULE_HANDLE, AppMemoryFuncs unsafe.Pointer) CSSM_RETURN
var _cSSM_GetAPIMemoryFunctionsErr error

// CanCallCSSM_GetAPIMemoryFunctions reports whether the symbol for CSSM_GetAPIMemoryFunctions is available on this system.
func CanCallCSSM_GetAPIMemoryFunctions() bool {
	return _cSSM_GetAPIMemoryFunctions != nil
}

// CSSM_GetAPIMemoryFunctionsCallError returns the symbol lookup or registration error for CSSM_GetAPIMemoryFunctions.
func CSSM_GetAPIMemoryFunctionsCallError() error {
	if _cSSM_GetAPIMemoryFunctions != nil {
		return nil
	}
	return symbolCallError("CSSM_GetAPIMemoryFunctions", "10.0", _cSSM_GetAPIMemoryFunctionsErr)
}

// TryCSSM_GetAPIMemoryFunctions calls CSSM_GetAPIMemoryFunctions without panicking when the symbol is unavailable.
func TryCSSM_GetAPIMemoryFunctions(AddInHandle CSSM_MODULE_HANDLE, AppMemoryFuncs unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_GetAPIMemoryFunctionsCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_GetAPIMemoryFunctions(AddInHandle, AppMemoryFuncs), nil
}

// CSSM_GetAPIMemoryFunctions.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GetAPIMemoryFunctions
func CSSM_GetAPIMemoryFunctions(AddInHandle CSSM_MODULE_HANDLE, AppMemoryFuncs unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_GetAPIMemoryFunctions(AddInHandle, AppMemoryFuncs)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_GetContext func(CCHandle CSSM_CC_HANDLE, Context unsafe.Pointer) CSSM_RETURN
var _cSSM_GetContextErr error

// CanCallCSSM_GetContext reports whether the symbol for CSSM_GetContext is available on this system.
func CanCallCSSM_GetContext() bool {
	return _cSSM_GetContext != nil
}

// CSSM_GetContextCallError returns the symbol lookup or registration error for CSSM_GetContext.
func CSSM_GetContextCallError() error {
	if _cSSM_GetContext != nil {
		return nil
	}
	return symbolCallError("CSSM_GetContext", "10.0", _cSSM_GetContextErr)
}

// TryCSSM_GetContext calls CSSM_GetContext without panicking when the symbol is unavailable.
func TryCSSM_GetContext(CCHandle CSSM_CC_HANDLE, Context unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_GetContextCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_GetContext(CCHandle, Context), nil
}

// CSSM_GetContext.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GetContext
func CSSM_GetContext(CCHandle CSSM_CC_HANDLE, Context unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_GetContext(CCHandle, Context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_GetContextAttribute func(Context unsafe.Pointer, AttributeType uint32, ContextAttribute unsafe.Pointer) CSSM_RETURN
var _cSSM_GetContextAttributeErr error

// CanCallCSSM_GetContextAttribute reports whether the symbol for CSSM_GetContextAttribute is available on this system.
func CanCallCSSM_GetContextAttribute() bool {
	return _cSSM_GetContextAttribute != nil
}

// CSSM_GetContextAttributeCallError returns the symbol lookup or registration error for CSSM_GetContextAttribute.
func CSSM_GetContextAttributeCallError() error {
	if _cSSM_GetContextAttribute != nil {
		return nil
	}
	return symbolCallError("CSSM_GetContextAttribute", "10.0", _cSSM_GetContextAttributeErr)
}

// TryCSSM_GetContextAttribute calls CSSM_GetContextAttribute without panicking when the symbol is unavailable.
func TryCSSM_GetContextAttribute(Context unsafe.Pointer, AttributeType uint32, ContextAttribute unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_GetContextAttributeCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_GetContextAttribute(Context, AttributeType, ContextAttribute), nil
}

// CSSM_GetContextAttribute.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GetContextAttribute
func CSSM_GetContextAttribute(Context unsafe.Pointer, AttributeType uint32, ContextAttribute unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_GetContextAttribute(Context, AttributeType, ContextAttribute)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_GetKeyAcl func(CSPHandle CSSM_CSP_HANDLE, Key unsafe.Pointer, SelectionTag unsafe.Pointer, NumberOfAclInfos *uint32, AclInfos unsafe.Pointer) CSSM_RETURN
var _cSSM_GetKeyAclErr error

// CanCallCSSM_GetKeyAcl reports whether the symbol for CSSM_GetKeyAcl is available on this system.
func CanCallCSSM_GetKeyAcl() bool {
	return _cSSM_GetKeyAcl != nil
}

// CSSM_GetKeyAclCallError returns the symbol lookup or registration error for CSSM_GetKeyAcl.
func CSSM_GetKeyAclCallError() error {
	if _cSSM_GetKeyAcl != nil {
		return nil
	}
	return symbolCallError("CSSM_GetKeyAcl", "10.0", _cSSM_GetKeyAclErr)
}

// TryCSSM_GetKeyAcl calls CSSM_GetKeyAcl without panicking when the symbol is unavailable.
func TryCSSM_GetKeyAcl(CSPHandle CSSM_CSP_HANDLE, Key unsafe.Pointer, SelectionTag unsafe.Pointer, NumberOfAclInfos *uint32, AclInfos unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_GetKeyAclCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_GetKeyAcl(CSPHandle, Key, SelectionTag, NumberOfAclInfos, AclInfos), nil
}

// CSSM_GetKeyAcl.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GetKeyAcl
func CSSM_GetKeyAcl(CSPHandle CSSM_CSP_HANDLE, Key unsafe.Pointer, SelectionTag unsafe.Pointer, NumberOfAclInfos *uint32, AclInfos unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_GetKeyAcl(CSPHandle, Key, SelectionTag, NumberOfAclInfos, AclInfos)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_GetKeyOwner func(CSPHandle CSSM_CSP_HANDLE, Key unsafe.Pointer, Owner unsafe.Pointer) CSSM_RETURN
var _cSSM_GetKeyOwnerErr error

// CanCallCSSM_GetKeyOwner reports whether the symbol for CSSM_GetKeyOwner is available on this system.
func CanCallCSSM_GetKeyOwner() bool {
	return _cSSM_GetKeyOwner != nil
}

// CSSM_GetKeyOwnerCallError returns the symbol lookup or registration error for CSSM_GetKeyOwner.
func CSSM_GetKeyOwnerCallError() error {
	if _cSSM_GetKeyOwner != nil {
		return nil
	}
	return symbolCallError("CSSM_GetKeyOwner", "10.0", _cSSM_GetKeyOwnerErr)
}

// TryCSSM_GetKeyOwner calls CSSM_GetKeyOwner without panicking when the symbol is unavailable.
func TryCSSM_GetKeyOwner(CSPHandle CSSM_CSP_HANDLE, Key unsafe.Pointer, Owner unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_GetKeyOwnerCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_GetKeyOwner(CSPHandle, Key, Owner), nil
}

// CSSM_GetKeyOwner.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GetKeyOwner
func CSSM_GetKeyOwner(CSPHandle CSSM_CSP_HANDLE, Key unsafe.Pointer, Owner unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_GetKeyOwner(CSPHandle, Key, Owner)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_GetModuleGUIDFromHandle func(ModuleHandle CSSM_MODULE_HANDLE, ModuleGUID unsafe.Pointer) CSSM_RETURN
var _cSSM_GetModuleGUIDFromHandleErr error

// CanCallCSSM_GetModuleGUIDFromHandle reports whether the symbol for CSSM_GetModuleGUIDFromHandle is available on this system.
func CanCallCSSM_GetModuleGUIDFromHandle() bool {
	return _cSSM_GetModuleGUIDFromHandle != nil
}

// CSSM_GetModuleGUIDFromHandleCallError returns the symbol lookup or registration error for CSSM_GetModuleGUIDFromHandle.
func CSSM_GetModuleGUIDFromHandleCallError() error {
	if _cSSM_GetModuleGUIDFromHandle != nil {
		return nil
	}
	return symbolCallError("CSSM_GetModuleGUIDFromHandle", "10.0", _cSSM_GetModuleGUIDFromHandleErr)
}

// TryCSSM_GetModuleGUIDFromHandle calls CSSM_GetModuleGUIDFromHandle without panicking when the symbol is unavailable.
func TryCSSM_GetModuleGUIDFromHandle(ModuleHandle CSSM_MODULE_HANDLE, ModuleGUID unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_GetModuleGUIDFromHandleCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_GetModuleGUIDFromHandle(ModuleHandle, ModuleGUID), nil
}

// CSSM_GetModuleGUIDFromHandle.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GetModuleGUIDFromHandle
func CSSM_GetModuleGUIDFromHandle(ModuleHandle CSSM_MODULE_HANDLE, ModuleGUID unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_GetModuleGUIDFromHandle(ModuleHandle, ModuleGUID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_GetPrivilege func(Privilege unsafe.Pointer) CSSM_RETURN
var _cSSM_GetPrivilegeErr error

// CanCallCSSM_GetPrivilege reports whether the symbol for CSSM_GetPrivilege is available on this system.
func CanCallCSSM_GetPrivilege() bool {
	return _cSSM_GetPrivilege != nil
}

// CSSM_GetPrivilegeCallError returns the symbol lookup or registration error for CSSM_GetPrivilege.
func CSSM_GetPrivilegeCallError() error {
	if _cSSM_GetPrivilege != nil {
		return nil
	}
	return symbolCallError("CSSM_GetPrivilege", "10.0", _cSSM_GetPrivilegeErr)
}

// TryCSSM_GetPrivilege calls CSSM_GetPrivilege without panicking when the symbol is unavailable.
func TryCSSM_GetPrivilege(Privilege unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_GetPrivilegeCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_GetPrivilege(Privilege), nil
}

// CSSM_GetPrivilege.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GetPrivilege
func CSSM_GetPrivilege(Privilege unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_GetPrivilege(Privilege)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_GetSubserviceUIDFromHandle func(ModuleHandle CSSM_MODULE_HANDLE, SubserviceUID unsafe.Pointer) CSSM_RETURN
var _cSSM_GetSubserviceUIDFromHandleErr error

// CanCallCSSM_GetSubserviceUIDFromHandle reports whether the symbol for CSSM_GetSubserviceUIDFromHandle is available on this system.
func CanCallCSSM_GetSubserviceUIDFromHandle() bool {
	return _cSSM_GetSubserviceUIDFromHandle != nil
}

// CSSM_GetSubserviceUIDFromHandleCallError returns the symbol lookup or registration error for CSSM_GetSubserviceUIDFromHandle.
func CSSM_GetSubserviceUIDFromHandleCallError() error {
	if _cSSM_GetSubserviceUIDFromHandle != nil {
		return nil
	}
	return symbolCallError("CSSM_GetSubserviceUIDFromHandle", "10.0", _cSSM_GetSubserviceUIDFromHandleErr)
}

// TryCSSM_GetSubserviceUIDFromHandle calls CSSM_GetSubserviceUIDFromHandle without panicking when the symbol is unavailable.
func TryCSSM_GetSubserviceUIDFromHandle(ModuleHandle CSSM_MODULE_HANDLE, SubserviceUID unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_GetSubserviceUIDFromHandleCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_GetSubserviceUIDFromHandle(ModuleHandle, SubserviceUID), nil
}

// CSSM_GetSubserviceUIDFromHandle.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GetSubserviceUIDFromHandle
func CSSM_GetSubserviceUIDFromHandle(ModuleHandle CSSM_MODULE_HANDLE, SubserviceUID unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_GetSubserviceUIDFromHandle(ModuleHandle, SubserviceUID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_GetTimeValue func(CSPHandle CSSM_CSP_HANDLE, TimeAlgorithm CSSM_ALGORITHMS, TimeData unsafe.Pointer) CSSM_RETURN
var _cSSM_GetTimeValueErr error

// CanCallCSSM_GetTimeValue reports whether the symbol for CSSM_GetTimeValue is available on this system.
func CanCallCSSM_GetTimeValue() bool {
	return _cSSM_GetTimeValue != nil
}

// CSSM_GetTimeValueCallError returns the symbol lookup or registration error for CSSM_GetTimeValue.
func CSSM_GetTimeValueCallError() error {
	if _cSSM_GetTimeValue != nil {
		return nil
	}
	return symbolCallError("CSSM_GetTimeValue", "10.0", _cSSM_GetTimeValueErr)
}

// TryCSSM_GetTimeValue calls CSSM_GetTimeValue without panicking when the symbol is unavailable.
func TryCSSM_GetTimeValue(CSPHandle CSSM_CSP_HANDLE, TimeAlgorithm CSSM_ALGORITHMS, TimeData unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_GetTimeValueCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_GetTimeValue(CSPHandle, TimeAlgorithm, TimeData), nil
}

// CSSM_GetTimeValue.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GetTimeValue
func CSSM_GetTimeValue(CSPHandle CSSM_CSP_HANDLE, TimeAlgorithm CSSM_ALGORITHMS, TimeData unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_GetTimeValue(CSPHandle, TimeAlgorithm, TimeData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_Init func(Version unsafe.Pointer, Scope CSSM_PRIVILEGE_SCOPE, CallerGuid unsafe.Pointer, KeyHierarchy CSSM_KEY_HIERARCHY, PvcPolicy unsafe.Pointer, Reserved unsafe.Pointer) CSSM_RETURN
var _cSSM_InitErr error

// CanCallCSSM_Init reports whether the symbol for CSSM_Init is available on this system.
func CanCallCSSM_Init() bool {
	return _cSSM_Init != nil
}

// CSSM_InitCallError returns the symbol lookup or registration error for CSSM_Init.
func CSSM_InitCallError() error {
	if _cSSM_Init != nil {
		return nil
	}
	return symbolCallError("CSSM_Init", "10.0", _cSSM_InitErr)
}

// TryCSSM_Init calls CSSM_Init without panicking when the symbol is unavailable.
func TryCSSM_Init(Version unsafe.Pointer, Scope CSSM_PRIVILEGE_SCOPE, CallerGuid unsafe.Pointer, KeyHierarchy CSSM_KEY_HIERARCHY, PvcPolicy unsafe.Pointer, Reserved unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_InitCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_Init(Version, Scope, CallerGuid, KeyHierarchy, PvcPolicy, Reserved), nil
}

// CSSM_Init.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_Init
func CSSM_Init(Version unsafe.Pointer, Scope CSSM_PRIVILEGE_SCOPE, CallerGuid unsafe.Pointer, KeyHierarchy CSSM_KEY_HIERARCHY, PvcPolicy unsafe.Pointer, Reserved unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_Init(Version, Scope, CallerGuid, KeyHierarchy, PvcPolicy, Reserved)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_Introduce func(ModuleID unsafe.Pointer, KeyHierarchy CSSM_KEY_HIERARCHY) CSSM_RETURN
var _cSSM_IntroduceErr error

// CanCallCSSM_Introduce reports whether the symbol for CSSM_Introduce is available on this system.
func CanCallCSSM_Introduce() bool {
	return _cSSM_Introduce != nil
}

// CSSM_IntroduceCallError returns the symbol lookup or registration error for CSSM_Introduce.
func CSSM_IntroduceCallError() error {
	if _cSSM_Introduce != nil {
		return nil
	}
	return symbolCallError("CSSM_Introduce", "10.0", _cSSM_IntroduceErr)
}

// TryCSSM_Introduce calls CSSM_Introduce without panicking when the symbol is unavailable.
func TryCSSM_Introduce(ModuleID unsafe.Pointer, KeyHierarchy CSSM_KEY_HIERARCHY) (CSSM_RETURN, error) {
	if err := CSSM_IntroduceCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_Introduce(ModuleID, KeyHierarchy), nil
}

// CSSM_Introduce.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_Introduce
func CSSM_Introduce(ModuleID unsafe.Pointer, KeyHierarchy CSSM_KEY_HIERARCHY) CSSM_RETURN {
	result, callErr := TryCSSM_Introduce(ModuleID, KeyHierarchy)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_ListAttachedModuleManagers func(NumberOfModuleManagers *uint32, ModuleManagerGuids unsafe.Pointer) CSSM_RETURN
var _cSSM_ListAttachedModuleManagersErr error

// CanCallCSSM_ListAttachedModuleManagers reports whether the symbol for CSSM_ListAttachedModuleManagers is available on this system.
func CanCallCSSM_ListAttachedModuleManagers() bool {
	return _cSSM_ListAttachedModuleManagers != nil
}

// CSSM_ListAttachedModuleManagersCallError returns the symbol lookup or registration error for CSSM_ListAttachedModuleManagers.
func CSSM_ListAttachedModuleManagersCallError() error {
	if _cSSM_ListAttachedModuleManagers != nil {
		return nil
	}
	return symbolCallError("CSSM_ListAttachedModuleManagers", "10.0", _cSSM_ListAttachedModuleManagersErr)
}

// TryCSSM_ListAttachedModuleManagers calls CSSM_ListAttachedModuleManagers without panicking when the symbol is unavailable.
func TryCSSM_ListAttachedModuleManagers(NumberOfModuleManagers *uint32, ModuleManagerGuids unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_ListAttachedModuleManagersCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_ListAttachedModuleManagers(NumberOfModuleManagers, ModuleManagerGuids), nil
}

// CSSM_ListAttachedModuleManagers.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_ListAttachedModuleManagers
func CSSM_ListAttachedModuleManagers(NumberOfModuleManagers *uint32, ModuleManagerGuids unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_ListAttachedModuleManagers(NumberOfModuleManagers, ModuleManagerGuids)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_ModuleAttach func(ModuleGuid unsafe.Pointer, Version unsafe.Pointer, MemoryFuncs unsafe.Pointer, SubserviceID uint32, SubServiceType CSSM_SERVICE_TYPE, AttachFlags CSSM_ATTACH_FLAGS, KeyHierarchy CSSM_KEY_HIERARCHY, FunctionTable unsafe.Pointer, NumFunctionTable uint32, Reserved unsafe.Pointer, NewModuleHandle CSSM_MODULE_HANDLE_PTR) CSSM_RETURN
var _cSSM_ModuleAttachErr error

// CanCallCSSM_ModuleAttach reports whether the symbol for CSSM_ModuleAttach is available on this system.
func CanCallCSSM_ModuleAttach() bool {
	return _cSSM_ModuleAttach != nil
}

// CSSM_ModuleAttachCallError returns the symbol lookup or registration error for CSSM_ModuleAttach.
func CSSM_ModuleAttachCallError() error {
	if _cSSM_ModuleAttach != nil {
		return nil
	}
	return symbolCallError("CSSM_ModuleAttach", "10.0", _cSSM_ModuleAttachErr)
}

// TryCSSM_ModuleAttach calls CSSM_ModuleAttach without panicking when the symbol is unavailable.
func TryCSSM_ModuleAttach(ModuleGuid unsafe.Pointer, Version unsafe.Pointer, MemoryFuncs unsafe.Pointer, SubserviceID uint32, SubServiceType CSSM_SERVICE_TYPE, AttachFlags CSSM_ATTACH_FLAGS, KeyHierarchy CSSM_KEY_HIERARCHY, FunctionTable unsafe.Pointer, NumFunctionTable uint32, Reserved unsafe.Pointer, NewModuleHandle CSSM_MODULE_HANDLE_PTR) (CSSM_RETURN, error) {
	if err := CSSM_ModuleAttachCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_ModuleAttach(ModuleGuid, Version, MemoryFuncs, SubserviceID, SubServiceType, AttachFlags, KeyHierarchy, FunctionTable, NumFunctionTable, Reserved, NewModuleHandle), nil
}

// CSSM_ModuleAttach.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_ModuleAttach
func CSSM_ModuleAttach(ModuleGuid unsafe.Pointer, Version unsafe.Pointer, MemoryFuncs unsafe.Pointer, SubserviceID uint32, SubServiceType CSSM_SERVICE_TYPE, AttachFlags CSSM_ATTACH_FLAGS, KeyHierarchy CSSM_KEY_HIERARCHY, FunctionTable unsafe.Pointer, NumFunctionTable uint32, Reserved unsafe.Pointer, NewModuleHandle CSSM_MODULE_HANDLE_PTR) CSSM_RETURN {
	result, callErr := TryCSSM_ModuleAttach(ModuleGuid, Version, MemoryFuncs, SubserviceID, SubServiceType, AttachFlags, KeyHierarchy, FunctionTable, NumFunctionTable, Reserved, NewModuleHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_ModuleDetach func(ModuleHandle CSSM_MODULE_HANDLE) CSSM_RETURN
var _cSSM_ModuleDetachErr error

// CanCallCSSM_ModuleDetach reports whether the symbol for CSSM_ModuleDetach is available on this system.
func CanCallCSSM_ModuleDetach() bool {
	return _cSSM_ModuleDetach != nil
}

// CSSM_ModuleDetachCallError returns the symbol lookup or registration error for CSSM_ModuleDetach.
func CSSM_ModuleDetachCallError() error {
	if _cSSM_ModuleDetach != nil {
		return nil
	}
	return symbolCallError("CSSM_ModuleDetach", "10.0", _cSSM_ModuleDetachErr)
}

// TryCSSM_ModuleDetach calls CSSM_ModuleDetach without panicking when the symbol is unavailable.
func TryCSSM_ModuleDetach(ModuleHandle CSSM_MODULE_HANDLE) (CSSM_RETURN, error) {
	if err := CSSM_ModuleDetachCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_ModuleDetach(ModuleHandle), nil
}

// CSSM_ModuleDetach.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_ModuleDetach
func CSSM_ModuleDetach(ModuleHandle CSSM_MODULE_HANDLE) CSSM_RETURN {
	result, callErr := TryCSSM_ModuleDetach(ModuleHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_ModuleLoad func(ModuleGuid unsafe.Pointer, KeyHierarchy CSSM_KEY_HIERARCHY, AppNotifyCallback unsafe.Pointer, AppNotifyCallbackCtx unsafe.Pointer) CSSM_RETURN
var _cSSM_ModuleLoadErr error

// CanCallCSSM_ModuleLoad reports whether the symbol for CSSM_ModuleLoad is available on this system.
func CanCallCSSM_ModuleLoad() bool {
	return _cSSM_ModuleLoad != nil
}

// CSSM_ModuleLoadCallError returns the symbol lookup or registration error for CSSM_ModuleLoad.
func CSSM_ModuleLoadCallError() error {
	if _cSSM_ModuleLoad != nil {
		return nil
	}
	return symbolCallError("CSSM_ModuleLoad", "10.0", _cSSM_ModuleLoadErr)
}

// TryCSSM_ModuleLoad calls CSSM_ModuleLoad without panicking when the symbol is unavailable.
func TryCSSM_ModuleLoad(ModuleGuid unsafe.Pointer, KeyHierarchy CSSM_KEY_HIERARCHY, AppNotifyCallback unsafe.Pointer, AppNotifyCallbackCtx unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_ModuleLoadCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_ModuleLoad(ModuleGuid, KeyHierarchy, AppNotifyCallback, AppNotifyCallbackCtx), nil
}

// CSSM_ModuleLoad.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_ModuleLoad
func CSSM_ModuleLoad(ModuleGuid unsafe.Pointer, KeyHierarchy CSSM_KEY_HIERARCHY, AppNotifyCallback unsafe.Pointer, AppNotifyCallbackCtx unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_ModuleLoad(ModuleGuid, KeyHierarchy, AppNotifyCallback, AppNotifyCallbackCtx)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_ModuleUnload func(ModuleGuid unsafe.Pointer, AppNotifyCallback unsafe.Pointer, AppNotifyCallbackCtx unsafe.Pointer) CSSM_RETURN
var _cSSM_ModuleUnloadErr error

// CanCallCSSM_ModuleUnload reports whether the symbol for CSSM_ModuleUnload is available on this system.
func CanCallCSSM_ModuleUnload() bool {
	return _cSSM_ModuleUnload != nil
}

// CSSM_ModuleUnloadCallError returns the symbol lookup or registration error for CSSM_ModuleUnload.
func CSSM_ModuleUnloadCallError() error {
	if _cSSM_ModuleUnload != nil {
		return nil
	}
	return symbolCallError("CSSM_ModuleUnload", "10.0", _cSSM_ModuleUnloadErr)
}

// TryCSSM_ModuleUnload calls CSSM_ModuleUnload without panicking when the symbol is unavailable.
func TryCSSM_ModuleUnload(ModuleGuid unsafe.Pointer, AppNotifyCallback unsafe.Pointer, AppNotifyCallbackCtx unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_ModuleUnloadCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_ModuleUnload(ModuleGuid, AppNotifyCallback, AppNotifyCallbackCtx), nil
}

// CSSM_ModuleUnload.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_ModuleUnload
func CSSM_ModuleUnload(ModuleGuid unsafe.Pointer, AppNotifyCallback unsafe.Pointer, AppNotifyCallbackCtx unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_ModuleUnload(ModuleGuid, AppNotifyCallback, AppNotifyCallbackCtx)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_QueryKeySizeInBits func(CSPHandle CSSM_CSP_HANDLE, CCHandle CSSM_CC_HANDLE, Key unsafe.Pointer, KeySize unsafe.Pointer) CSSM_RETURN
var _cSSM_QueryKeySizeInBitsErr error

// CanCallCSSM_QueryKeySizeInBits reports whether the symbol for CSSM_QueryKeySizeInBits is available on this system.
func CanCallCSSM_QueryKeySizeInBits() bool {
	return _cSSM_QueryKeySizeInBits != nil
}

// CSSM_QueryKeySizeInBitsCallError returns the symbol lookup or registration error for CSSM_QueryKeySizeInBits.
func CSSM_QueryKeySizeInBitsCallError() error {
	if _cSSM_QueryKeySizeInBits != nil {
		return nil
	}
	return symbolCallError("CSSM_QueryKeySizeInBits", "10.0", _cSSM_QueryKeySizeInBitsErr)
}

// TryCSSM_QueryKeySizeInBits calls CSSM_QueryKeySizeInBits without panicking when the symbol is unavailable.
func TryCSSM_QueryKeySizeInBits(CSPHandle CSSM_CSP_HANDLE, CCHandle CSSM_CC_HANDLE, Key unsafe.Pointer, KeySize unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_QueryKeySizeInBitsCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_QueryKeySizeInBits(CSPHandle, CCHandle, Key, KeySize), nil
}

// CSSM_QueryKeySizeInBits.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_QueryKeySizeInBits
func CSSM_QueryKeySizeInBits(CSPHandle CSSM_CSP_HANDLE, CCHandle CSSM_CC_HANDLE, Key unsafe.Pointer, KeySize unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_QueryKeySizeInBits(CSPHandle, CCHandle, Key, KeySize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_QuerySize func(CCHandle CSSM_CC_HANDLE, Encrypt CSSM_BOOL, QuerySizeCount uint32, DataBlockSizes unsafe.Pointer) CSSM_RETURN
var _cSSM_QuerySizeErr error

// CanCallCSSM_QuerySize reports whether the symbol for CSSM_QuerySize is available on this system.
func CanCallCSSM_QuerySize() bool {
	return _cSSM_QuerySize != nil
}

// CSSM_QuerySizeCallError returns the symbol lookup or registration error for CSSM_QuerySize.
func CSSM_QuerySizeCallError() error {
	if _cSSM_QuerySize != nil {
		return nil
	}
	return symbolCallError("CSSM_QuerySize", "10.0", _cSSM_QuerySizeErr)
}

// TryCSSM_QuerySize calls CSSM_QuerySize without panicking when the symbol is unavailable.
func TryCSSM_QuerySize(CCHandle CSSM_CC_HANDLE, Encrypt CSSM_BOOL, QuerySizeCount uint32, DataBlockSizes unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_QuerySizeCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_QuerySize(CCHandle, Encrypt, QuerySizeCount, DataBlockSizes), nil
}

// CSSM_QuerySize.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_QuerySize
func CSSM_QuerySize(CCHandle CSSM_CC_HANDLE, Encrypt CSSM_BOOL, QuerySizeCount uint32, DataBlockSizes unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_QuerySize(CCHandle, Encrypt, QuerySizeCount, DataBlockSizes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_RetrieveCounter func(CSPHandle CSSM_CSP_HANDLE, Counter unsafe.Pointer) CSSM_RETURN
var _cSSM_RetrieveCounterErr error

// CanCallCSSM_RetrieveCounter reports whether the symbol for CSSM_RetrieveCounter is available on this system.
func CanCallCSSM_RetrieveCounter() bool {
	return _cSSM_RetrieveCounter != nil
}

// CSSM_RetrieveCounterCallError returns the symbol lookup or registration error for CSSM_RetrieveCounter.
func CSSM_RetrieveCounterCallError() error {
	if _cSSM_RetrieveCounter != nil {
		return nil
	}
	return symbolCallError("CSSM_RetrieveCounter", "10.0", _cSSM_RetrieveCounterErr)
}

// TryCSSM_RetrieveCounter calls CSSM_RetrieveCounter without panicking when the symbol is unavailable.
func TryCSSM_RetrieveCounter(CSPHandle CSSM_CSP_HANDLE, Counter unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_RetrieveCounterCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_RetrieveCounter(CSPHandle, Counter), nil
}

// CSSM_RetrieveCounter.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_RetrieveCounter
func CSSM_RetrieveCounter(CSPHandle CSSM_CSP_HANDLE, Counter unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_RetrieveCounter(CSPHandle, Counter)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_RetrieveUniqueId func(CSPHandle CSSM_CSP_HANDLE, UniqueID unsafe.Pointer) CSSM_RETURN
var _cSSM_RetrieveUniqueIdErr error

// CanCallCSSM_RetrieveUniqueId reports whether the symbol for CSSM_RetrieveUniqueId is available on this system.
func CanCallCSSM_RetrieveUniqueId() bool {
	return _cSSM_RetrieveUniqueId != nil
}

// CSSM_RetrieveUniqueIdCallError returns the symbol lookup or registration error for CSSM_RetrieveUniqueId.
func CSSM_RetrieveUniqueIdCallError() error {
	if _cSSM_RetrieveUniqueId != nil {
		return nil
	}
	return symbolCallError("CSSM_RetrieveUniqueId", "10.0", _cSSM_RetrieveUniqueIdErr)
}

// TryCSSM_RetrieveUniqueId calls CSSM_RetrieveUniqueId without panicking when the symbol is unavailable.
func TryCSSM_RetrieveUniqueId(CSPHandle CSSM_CSP_HANDLE, UniqueID unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_RetrieveUniqueIdCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_RetrieveUniqueId(CSPHandle, UniqueID), nil
}

// CSSM_RetrieveUniqueId.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_RetrieveUniqueId
func CSSM_RetrieveUniqueId(CSPHandle CSSM_CSP_HANDLE, UniqueID unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_RetrieveUniqueId(CSPHandle, UniqueID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_SetContext func(CCHandle CSSM_CC_HANDLE, Context unsafe.Pointer) CSSM_RETURN
var _cSSM_SetContextErr error

// CanCallCSSM_SetContext reports whether the symbol for CSSM_SetContext is available on this system.
func CanCallCSSM_SetContext() bool {
	return _cSSM_SetContext != nil
}

// CSSM_SetContextCallError returns the symbol lookup or registration error for CSSM_SetContext.
func CSSM_SetContextCallError() error {
	if _cSSM_SetContext != nil {
		return nil
	}
	return symbolCallError("CSSM_SetContext", "10.0", _cSSM_SetContextErr)
}

// TryCSSM_SetContext calls CSSM_SetContext without panicking when the symbol is unavailable.
func TryCSSM_SetContext(CCHandle CSSM_CC_HANDLE, Context unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_SetContextCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_SetContext(CCHandle, Context), nil
}

// CSSM_SetContext.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_SetContext
func CSSM_SetContext(CCHandle CSSM_CC_HANDLE, Context unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_SetContext(CCHandle, Context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_SetPrivilege func(Privilege CSSM_PRIVILEGE) CSSM_RETURN
var _cSSM_SetPrivilegeErr error

// CanCallCSSM_SetPrivilege reports whether the symbol for CSSM_SetPrivilege is available on this system.
func CanCallCSSM_SetPrivilege() bool {
	return _cSSM_SetPrivilege != nil
}

// CSSM_SetPrivilegeCallError returns the symbol lookup or registration error for CSSM_SetPrivilege.
func CSSM_SetPrivilegeCallError() error {
	if _cSSM_SetPrivilege != nil {
		return nil
	}
	return symbolCallError("CSSM_SetPrivilege", "10.0", _cSSM_SetPrivilegeErr)
}

// TryCSSM_SetPrivilege calls CSSM_SetPrivilege without panicking when the symbol is unavailable.
func TryCSSM_SetPrivilege(Privilege CSSM_PRIVILEGE) (CSSM_RETURN, error) {
	if err := CSSM_SetPrivilegeCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_SetPrivilege(Privilege), nil
}

// CSSM_SetPrivilege.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_SetPrivilege
func CSSM_SetPrivilege(Privilege CSSM_PRIVILEGE) CSSM_RETURN {
	result, callErr := TryCSSM_SetPrivilege(Privilege)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_SignData func(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, DigestAlgorithm CSSM_ALGORITHMS, Signature unsafe.Pointer) CSSM_RETURN
var _cSSM_SignDataErr error

// CanCallCSSM_SignData reports whether the symbol for CSSM_SignData is available on this system.
func CanCallCSSM_SignData() bool {
	return _cSSM_SignData != nil
}

// CSSM_SignDataCallError returns the symbol lookup or registration error for CSSM_SignData.
func CSSM_SignDataCallError() error {
	if _cSSM_SignData != nil {
		return nil
	}
	return symbolCallError("CSSM_SignData", "10.0", _cSSM_SignDataErr)
}

// TryCSSM_SignData calls CSSM_SignData without panicking when the symbol is unavailable.
func TryCSSM_SignData(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, DigestAlgorithm CSSM_ALGORITHMS, Signature unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_SignDataCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_SignData(CCHandle, DataBufs, DataBufCount, DigestAlgorithm, Signature), nil
}

// CSSM_SignData.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_SignData
func CSSM_SignData(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, DigestAlgorithm CSSM_ALGORITHMS, Signature unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_SignData(CCHandle, DataBufs, DataBufCount, DigestAlgorithm, Signature)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_SignDataFinal func(CCHandle CSSM_CC_HANDLE, Signature unsafe.Pointer) CSSM_RETURN
var _cSSM_SignDataFinalErr error

// CanCallCSSM_SignDataFinal reports whether the symbol for CSSM_SignDataFinal is available on this system.
func CanCallCSSM_SignDataFinal() bool {
	return _cSSM_SignDataFinal != nil
}

// CSSM_SignDataFinalCallError returns the symbol lookup or registration error for CSSM_SignDataFinal.
func CSSM_SignDataFinalCallError() error {
	if _cSSM_SignDataFinal != nil {
		return nil
	}
	return symbolCallError("CSSM_SignDataFinal", "10.0", _cSSM_SignDataFinalErr)
}

// TryCSSM_SignDataFinal calls CSSM_SignDataFinal without panicking when the symbol is unavailable.
func TryCSSM_SignDataFinal(CCHandle CSSM_CC_HANDLE, Signature unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_SignDataFinalCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_SignDataFinal(CCHandle, Signature), nil
}

// CSSM_SignDataFinal.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_SignDataFinal
func CSSM_SignDataFinal(CCHandle CSSM_CC_HANDLE, Signature unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_SignDataFinal(CCHandle, Signature)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_SignDataInit func(CCHandle CSSM_CC_HANDLE) CSSM_RETURN
var _cSSM_SignDataInitErr error

// CanCallCSSM_SignDataInit reports whether the symbol for CSSM_SignDataInit is available on this system.
func CanCallCSSM_SignDataInit() bool {
	return _cSSM_SignDataInit != nil
}

// CSSM_SignDataInitCallError returns the symbol lookup or registration error for CSSM_SignDataInit.
func CSSM_SignDataInitCallError() error {
	if _cSSM_SignDataInit != nil {
		return nil
	}
	return symbolCallError("CSSM_SignDataInit", "10.0", _cSSM_SignDataInitErr)
}

// TryCSSM_SignDataInit calls CSSM_SignDataInit without panicking when the symbol is unavailable.
func TryCSSM_SignDataInit(CCHandle CSSM_CC_HANDLE) (CSSM_RETURN, error) {
	if err := CSSM_SignDataInitCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_SignDataInit(CCHandle), nil
}

// CSSM_SignDataInit.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_SignDataInit
func CSSM_SignDataInit(CCHandle CSSM_CC_HANDLE) CSSM_RETURN {
	result, callErr := TryCSSM_SignDataInit(CCHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_SignDataUpdate func(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) CSSM_RETURN
var _cSSM_SignDataUpdateErr error

// CanCallCSSM_SignDataUpdate reports whether the symbol for CSSM_SignDataUpdate is available on this system.
func CanCallCSSM_SignDataUpdate() bool {
	return _cSSM_SignDataUpdate != nil
}

// CSSM_SignDataUpdateCallError returns the symbol lookup or registration error for CSSM_SignDataUpdate.
func CSSM_SignDataUpdateCallError() error {
	if _cSSM_SignDataUpdate != nil {
		return nil
	}
	return symbolCallError("CSSM_SignDataUpdate", "10.0", _cSSM_SignDataUpdateErr)
}

// TryCSSM_SignDataUpdate calls CSSM_SignDataUpdate without panicking when the symbol is unavailable.
func TryCSSM_SignDataUpdate(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) (CSSM_RETURN, error) {
	if err := CSSM_SignDataUpdateCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_SignDataUpdate(CCHandle, DataBufs, DataBufCount), nil
}

// CSSM_SignDataUpdate.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_SignDataUpdate
func CSSM_SignDataUpdate(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) CSSM_RETURN {
	result, callErr := TryCSSM_SignDataUpdate(CCHandle, DataBufs, DataBufCount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_ApplyCrlToDb func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, CrlToBeApplied unsafe.Pointer, SignerCertGroup unsafe.Pointer, ApplyCrlVerifyContext unsafe.Pointer, ApplyCrlVerifyResult unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_ApplyCrlToDbErr error

// CanCallCSSM_TP_ApplyCrlToDb reports whether the symbol for CSSM_TP_ApplyCrlToDb is available on this system.
func CanCallCSSM_TP_ApplyCrlToDb() bool {
	return _cSSM_TP_ApplyCrlToDb != nil
}

// CSSM_TP_ApplyCrlToDbCallError returns the symbol lookup or registration error for CSSM_TP_ApplyCrlToDb.
func CSSM_TP_ApplyCrlToDbCallError() error {
	if _cSSM_TP_ApplyCrlToDb != nil {
		return nil
	}
	return symbolCallError("CSSM_TP_ApplyCrlToDb", "10.0", _cSSM_TP_ApplyCrlToDbErr)
}

// TryCSSM_TP_ApplyCrlToDb calls CSSM_TP_ApplyCrlToDb without panicking when the symbol is unavailable.
func TryCSSM_TP_ApplyCrlToDb(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, CrlToBeApplied unsafe.Pointer, SignerCertGroup unsafe.Pointer, ApplyCrlVerifyContext unsafe.Pointer, ApplyCrlVerifyResult unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_TP_ApplyCrlToDbCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_TP_ApplyCrlToDb(TPHandle, CLHandle, CSPHandle, CrlToBeApplied, SignerCertGroup, ApplyCrlVerifyContext, ApplyCrlVerifyResult), nil
}

// CSSM_TP_ApplyCrlToDb.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_ApplyCrlToDb
func CSSM_TP_ApplyCrlToDb(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, CrlToBeApplied unsafe.Pointer, SignerCertGroup unsafe.Pointer, ApplyCrlVerifyContext unsafe.Pointer, ApplyCrlVerifyResult unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_TP_ApplyCrlToDb(TPHandle, CLHandle, CSPHandle, CrlToBeApplied, SignerCertGroup, ApplyCrlVerifyContext, ApplyCrlVerifyResult)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_CertCreateTemplate func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CertFields unsafe.Pointer, CertTemplate unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_CertCreateTemplateErr error

// CanCallCSSM_TP_CertCreateTemplate reports whether the symbol for CSSM_TP_CertCreateTemplate is available on this system.
func CanCallCSSM_TP_CertCreateTemplate() bool {
	return _cSSM_TP_CertCreateTemplate != nil
}

// CSSM_TP_CertCreateTemplateCallError returns the symbol lookup or registration error for CSSM_TP_CertCreateTemplate.
func CSSM_TP_CertCreateTemplateCallError() error {
	if _cSSM_TP_CertCreateTemplate != nil {
		return nil
	}
	return symbolCallError("CSSM_TP_CertCreateTemplate", "10.0", _cSSM_TP_CertCreateTemplateErr)
}

// TryCSSM_TP_CertCreateTemplate calls CSSM_TP_CertCreateTemplate without panicking when the symbol is unavailable.
func TryCSSM_TP_CertCreateTemplate(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CertFields unsafe.Pointer, CertTemplate unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_TP_CertCreateTemplateCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_TP_CertCreateTemplate(TPHandle, CLHandle, NumberOfFields, CertFields, CertTemplate), nil
}

// CSSM_TP_CertCreateTemplate.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CertCreateTemplate
func CSSM_TP_CertCreateTemplate(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CertFields unsafe.Pointer, CertTemplate unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_TP_CertCreateTemplate(TPHandle, CLHandle, NumberOfFields, CertFields, CertTemplate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_CertGetAllTemplateFields func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CertTemplate unsafe.Pointer, NumberOfFields *uint32, CertFields unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_CertGetAllTemplateFieldsErr error

// CanCallCSSM_TP_CertGetAllTemplateFields reports whether the symbol for CSSM_TP_CertGetAllTemplateFields is available on this system.
func CanCallCSSM_TP_CertGetAllTemplateFields() bool {
	return _cSSM_TP_CertGetAllTemplateFields != nil
}

// CSSM_TP_CertGetAllTemplateFieldsCallError returns the symbol lookup or registration error for CSSM_TP_CertGetAllTemplateFields.
func CSSM_TP_CertGetAllTemplateFieldsCallError() error {
	if _cSSM_TP_CertGetAllTemplateFields != nil {
		return nil
	}
	return symbolCallError("CSSM_TP_CertGetAllTemplateFields", "10.0", _cSSM_TP_CertGetAllTemplateFieldsErr)
}

// TryCSSM_TP_CertGetAllTemplateFields calls CSSM_TP_CertGetAllTemplateFields without panicking when the symbol is unavailable.
func TryCSSM_TP_CertGetAllTemplateFields(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CertTemplate unsafe.Pointer, NumberOfFields *uint32, CertFields unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_TP_CertGetAllTemplateFieldsCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_TP_CertGetAllTemplateFields(TPHandle, CLHandle, CertTemplate, NumberOfFields, CertFields), nil
}

// CSSM_TP_CertGetAllTemplateFields.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CertGetAllTemplateFields
func CSSM_TP_CertGetAllTemplateFields(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CertTemplate unsafe.Pointer, NumberOfFields *uint32, CertFields unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_TP_CertGetAllTemplateFields(TPHandle, CLHandle, CertTemplate, NumberOfFields, CertFields)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_CertGroupConstruct func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, DBList unsafe.Pointer, ConstructParams unsafe.Pointer, CertGroupFrag unsafe.Pointer, CertGroup unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_CertGroupConstructErr error

// CanCallCSSM_TP_CertGroupConstruct reports whether the symbol for CSSM_TP_CertGroupConstruct is available on this system.
func CanCallCSSM_TP_CertGroupConstruct() bool {
	return _cSSM_TP_CertGroupConstruct != nil
}

// CSSM_TP_CertGroupConstructCallError returns the symbol lookup or registration error for CSSM_TP_CertGroupConstruct.
func CSSM_TP_CertGroupConstructCallError() error {
	if _cSSM_TP_CertGroupConstruct != nil {
		return nil
	}
	return symbolCallError("CSSM_TP_CertGroupConstruct", "10.0", _cSSM_TP_CertGroupConstructErr)
}

// TryCSSM_TP_CertGroupConstruct calls CSSM_TP_CertGroupConstruct without panicking when the symbol is unavailable.
func TryCSSM_TP_CertGroupConstruct(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, DBList unsafe.Pointer, ConstructParams unsafe.Pointer, CertGroupFrag unsafe.Pointer, CertGroup unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_TP_CertGroupConstructCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_TP_CertGroupConstruct(TPHandle, CLHandle, CSPHandle, DBList, ConstructParams, CertGroupFrag, CertGroup), nil
}

// CSSM_TP_CertGroupConstruct.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CertGroupConstruct
func CSSM_TP_CertGroupConstruct(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, DBList unsafe.Pointer, ConstructParams unsafe.Pointer, CertGroupFrag unsafe.Pointer, CertGroup unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_TP_CertGroupConstruct(TPHandle, CLHandle, CSPHandle, DBList, ConstructParams, CertGroupFrag, CertGroup)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_CertGroupPrune func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, DBList unsafe.Pointer, OrderedCertGroup unsafe.Pointer, PrunedCertGroup unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_CertGroupPruneErr error

// CanCallCSSM_TP_CertGroupPrune reports whether the symbol for CSSM_TP_CertGroupPrune is available on this system.
func CanCallCSSM_TP_CertGroupPrune() bool {
	return _cSSM_TP_CertGroupPrune != nil
}

// CSSM_TP_CertGroupPruneCallError returns the symbol lookup or registration error for CSSM_TP_CertGroupPrune.
func CSSM_TP_CertGroupPruneCallError() error {
	if _cSSM_TP_CertGroupPrune != nil {
		return nil
	}
	return symbolCallError("CSSM_TP_CertGroupPrune", "10.0", _cSSM_TP_CertGroupPruneErr)
}

// TryCSSM_TP_CertGroupPrune calls CSSM_TP_CertGroupPrune without panicking when the symbol is unavailable.
func TryCSSM_TP_CertGroupPrune(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, DBList unsafe.Pointer, OrderedCertGroup unsafe.Pointer, PrunedCertGroup unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_TP_CertGroupPruneCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_TP_CertGroupPrune(TPHandle, CLHandle, DBList, OrderedCertGroup, PrunedCertGroup), nil
}

// CSSM_TP_CertGroupPrune.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CertGroupPrune
func CSSM_TP_CertGroupPrune(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, DBList unsafe.Pointer, OrderedCertGroup unsafe.Pointer, PrunedCertGroup unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_TP_CertGroupPrune(TPHandle, CLHandle, DBList, OrderedCertGroup, PrunedCertGroup)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_CertGroupToTupleGroup func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CertGroup unsafe.Pointer, TupleGroup unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_CertGroupToTupleGroupErr error

// CanCallCSSM_TP_CertGroupToTupleGroup reports whether the symbol for CSSM_TP_CertGroupToTupleGroup is available on this system.
func CanCallCSSM_TP_CertGroupToTupleGroup() bool {
	return _cSSM_TP_CertGroupToTupleGroup != nil
}

// CSSM_TP_CertGroupToTupleGroupCallError returns the symbol lookup or registration error for CSSM_TP_CertGroupToTupleGroup.
func CSSM_TP_CertGroupToTupleGroupCallError() error {
	if _cSSM_TP_CertGroupToTupleGroup != nil {
		return nil
	}
	return symbolCallError("CSSM_TP_CertGroupToTupleGroup", "10.0", _cSSM_TP_CertGroupToTupleGroupErr)
}

// TryCSSM_TP_CertGroupToTupleGroup calls CSSM_TP_CertGroupToTupleGroup without panicking when the symbol is unavailable.
func TryCSSM_TP_CertGroupToTupleGroup(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CertGroup unsafe.Pointer, TupleGroup unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_TP_CertGroupToTupleGroupCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_TP_CertGroupToTupleGroup(TPHandle, CLHandle, CertGroup, TupleGroup), nil
}

// CSSM_TP_CertGroupToTupleGroup.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CertGroupToTupleGroup
func CSSM_TP_CertGroupToTupleGroup(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CertGroup unsafe.Pointer, TupleGroup unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_TP_CertGroupToTupleGroup(TPHandle, CLHandle, CertGroup, TupleGroup)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_CertGroupVerify func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, CertGroupToBeVerified unsafe.Pointer, VerifyContext unsafe.Pointer, VerifyContextResult unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_CertGroupVerifyErr error

// CanCallCSSM_TP_CertGroupVerify reports whether the symbol for CSSM_TP_CertGroupVerify is available on this system.
func CanCallCSSM_TP_CertGroupVerify() bool {
	return _cSSM_TP_CertGroupVerify != nil
}

// CSSM_TP_CertGroupVerifyCallError returns the symbol lookup or registration error for CSSM_TP_CertGroupVerify.
func CSSM_TP_CertGroupVerifyCallError() error {
	if _cSSM_TP_CertGroupVerify != nil {
		return nil
	}
	return symbolCallError("CSSM_TP_CertGroupVerify", "10.0", _cSSM_TP_CertGroupVerifyErr)
}

// TryCSSM_TP_CertGroupVerify calls CSSM_TP_CertGroupVerify without panicking when the symbol is unavailable.
func TryCSSM_TP_CertGroupVerify(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, CertGroupToBeVerified unsafe.Pointer, VerifyContext unsafe.Pointer, VerifyContextResult unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_TP_CertGroupVerifyCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_TP_CertGroupVerify(TPHandle, CLHandle, CSPHandle, CertGroupToBeVerified, VerifyContext, VerifyContextResult), nil
}

// CSSM_TP_CertGroupVerify.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CertGroupVerify
func CSSM_TP_CertGroupVerify(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, CertGroupToBeVerified unsafe.Pointer, VerifyContext unsafe.Pointer, VerifyContextResult unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_TP_CertGroupVerify(TPHandle, CLHandle, CSPHandle, CertGroupToBeVerified, VerifyContext, VerifyContextResult)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_CertReclaimAbort func(TPHandle CSSM_TP_HANDLE, KeyCacheHandle CSSM_LONG_HANDLE) CSSM_RETURN
var _cSSM_TP_CertReclaimAbortErr error

// CanCallCSSM_TP_CertReclaimAbort reports whether the symbol for CSSM_TP_CertReclaimAbort is available on this system.
func CanCallCSSM_TP_CertReclaimAbort() bool {
	return _cSSM_TP_CertReclaimAbort != nil
}

// CSSM_TP_CertReclaimAbortCallError returns the symbol lookup or registration error for CSSM_TP_CertReclaimAbort.
func CSSM_TP_CertReclaimAbortCallError() error {
	if _cSSM_TP_CertReclaimAbort != nil {
		return nil
	}
	return symbolCallError("CSSM_TP_CertReclaimAbort", "10.0", _cSSM_TP_CertReclaimAbortErr)
}

// TryCSSM_TP_CertReclaimAbort calls CSSM_TP_CertReclaimAbort without panicking when the symbol is unavailable.
func TryCSSM_TP_CertReclaimAbort(TPHandle CSSM_TP_HANDLE, KeyCacheHandle CSSM_LONG_HANDLE) (CSSM_RETURN, error) {
	if err := CSSM_TP_CertReclaimAbortCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_TP_CertReclaimAbort(TPHandle, KeyCacheHandle), nil
}

// CSSM_TP_CertReclaimAbort.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CertReclaimAbort
func CSSM_TP_CertReclaimAbort(TPHandle CSSM_TP_HANDLE, KeyCacheHandle CSSM_LONG_HANDLE) CSSM_RETURN {
	result, callErr := TryCSSM_TP_CertReclaimAbort(TPHandle, KeyCacheHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_CertReclaimKey func(TPHandle CSSM_TP_HANDLE, CertGroup unsafe.Pointer, CertIndex uint32, KeyCacheHandle CSSM_LONG_HANDLE, CSPHandle CSSM_CSP_HANDLE, CredAndAclEntry unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_CertReclaimKeyErr error

// CanCallCSSM_TP_CertReclaimKey reports whether the symbol for CSSM_TP_CertReclaimKey is available on this system.
func CanCallCSSM_TP_CertReclaimKey() bool {
	return _cSSM_TP_CertReclaimKey != nil
}

// CSSM_TP_CertReclaimKeyCallError returns the symbol lookup or registration error for CSSM_TP_CertReclaimKey.
func CSSM_TP_CertReclaimKeyCallError() error {
	if _cSSM_TP_CertReclaimKey != nil {
		return nil
	}
	return symbolCallError("CSSM_TP_CertReclaimKey", "10.0", _cSSM_TP_CertReclaimKeyErr)
}

// TryCSSM_TP_CertReclaimKey calls CSSM_TP_CertReclaimKey without panicking when the symbol is unavailable.
func TryCSSM_TP_CertReclaimKey(TPHandle CSSM_TP_HANDLE, CertGroup unsafe.Pointer, CertIndex uint32, KeyCacheHandle CSSM_LONG_HANDLE, CSPHandle CSSM_CSP_HANDLE, CredAndAclEntry unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_TP_CertReclaimKeyCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_TP_CertReclaimKey(TPHandle, CertGroup, CertIndex, KeyCacheHandle, CSPHandle, CredAndAclEntry), nil
}

// CSSM_TP_CertReclaimKey.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CertReclaimKey
func CSSM_TP_CertReclaimKey(TPHandle CSSM_TP_HANDLE, CertGroup unsafe.Pointer, CertIndex uint32, KeyCacheHandle CSSM_LONG_HANDLE, CSPHandle CSSM_CSP_HANDLE, CredAndAclEntry unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_TP_CertReclaimKey(TPHandle, CertGroup, CertIndex, KeyCacheHandle, CSPHandle, CredAndAclEntry)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_CertRemoveFromCrlTemplate func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, OldCrlTemplate unsafe.Pointer, CertGroupToBeRemoved unsafe.Pointer, RevokerCertGroup unsafe.Pointer, RevokerVerifyContext unsafe.Pointer, RevokerVerifyResult unsafe.Pointer, NewCrlTemplate unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_CertRemoveFromCrlTemplateErr error

// CanCallCSSM_TP_CertRemoveFromCrlTemplate reports whether the symbol for CSSM_TP_CertRemoveFromCrlTemplate is available on this system.
func CanCallCSSM_TP_CertRemoveFromCrlTemplate() bool {
	return _cSSM_TP_CertRemoveFromCrlTemplate != nil
}

// CSSM_TP_CertRemoveFromCrlTemplateCallError returns the symbol lookup or registration error for CSSM_TP_CertRemoveFromCrlTemplate.
func CSSM_TP_CertRemoveFromCrlTemplateCallError() error {
	if _cSSM_TP_CertRemoveFromCrlTemplate != nil {
		return nil
	}
	return symbolCallError("CSSM_TP_CertRemoveFromCrlTemplate", "10.0", _cSSM_TP_CertRemoveFromCrlTemplateErr)
}

// TryCSSM_TP_CertRemoveFromCrlTemplate calls CSSM_TP_CertRemoveFromCrlTemplate without panicking when the symbol is unavailable.
func TryCSSM_TP_CertRemoveFromCrlTemplate(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, OldCrlTemplate unsafe.Pointer, CertGroupToBeRemoved unsafe.Pointer, RevokerCertGroup unsafe.Pointer, RevokerVerifyContext unsafe.Pointer, RevokerVerifyResult unsafe.Pointer, NewCrlTemplate unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_TP_CertRemoveFromCrlTemplateCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_TP_CertRemoveFromCrlTemplate(TPHandle, CLHandle, CSPHandle, OldCrlTemplate, CertGroupToBeRemoved, RevokerCertGroup, RevokerVerifyContext, RevokerVerifyResult, NewCrlTemplate), nil
}

// CSSM_TP_CertRemoveFromCrlTemplate.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CertRemoveFromCrlTemplate
func CSSM_TP_CertRemoveFromCrlTemplate(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, OldCrlTemplate unsafe.Pointer, CertGroupToBeRemoved unsafe.Pointer, RevokerCertGroup unsafe.Pointer, RevokerVerifyContext unsafe.Pointer, RevokerVerifyResult unsafe.Pointer, NewCrlTemplate unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_TP_CertRemoveFromCrlTemplate(TPHandle, CLHandle, CSPHandle, OldCrlTemplate, CertGroupToBeRemoved, RevokerCertGroup, RevokerVerifyContext, RevokerVerifyResult, NewCrlTemplate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_CertRevoke func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, OldCrlTemplate unsafe.Pointer, CertGroupToBeRevoked unsafe.Pointer, RevokerCertGroup unsafe.Pointer, RevokerVerifyContext unsafe.Pointer, RevokerVerifyResult unsafe.Pointer, Reason CSSM_TP_CERTCHANGE_REASON, NewCrlTemplate unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_CertRevokeErr error

// CanCallCSSM_TP_CertRevoke reports whether the symbol for CSSM_TP_CertRevoke is available on this system.
func CanCallCSSM_TP_CertRevoke() bool {
	return _cSSM_TP_CertRevoke != nil
}

// CSSM_TP_CertRevokeCallError returns the symbol lookup or registration error for CSSM_TP_CertRevoke.
func CSSM_TP_CertRevokeCallError() error {
	if _cSSM_TP_CertRevoke != nil {
		return nil
	}
	return symbolCallError("CSSM_TP_CertRevoke", "10.0", _cSSM_TP_CertRevokeErr)
}

// TryCSSM_TP_CertRevoke calls CSSM_TP_CertRevoke without panicking when the symbol is unavailable.
func TryCSSM_TP_CertRevoke(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, OldCrlTemplate unsafe.Pointer, CertGroupToBeRevoked unsafe.Pointer, RevokerCertGroup unsafe.Pointer, RevokerVerifyContext unsafe.Pointer, RevokerVerifyResult unsafe.Pointer, Reason CSSM_TP_CERTCHANGE_REASON, NewCrlTemplate unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_TP_CertRevokeCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_TP_CertRevoke(TPHandle, CLHandle, CSPHandle, OldCrlTemplate, CertGroupToBeRevoked, RevokerCertGroup, RevokerVerifyContext, RevokerVerifyResult, Reason, NewCrlTemplate), nil
}

// CSSM_TP_CertRevoke.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CertRevoke
func CSSM_TP_CertRevoke(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, OldCrlTemplate unsafe.Pointer, CertGroupToBeRevoked unsafe.Pointer, RevokerCertGroup unsafe.Pointer, RevokerVerifyContext unsafe.Pointer, RevokerVerifyResult unsafe.Pointer, Reason CSSM_TP_CERTCHANGE_REASON, NewCrlTemplate unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_TP_CertRevoke(TPHandle, CLHandle, CSPHandle, OldCrlTemplate, CertGroupToBeRevoked, RevokerCertGroup, RevokerVerifyContext, RevokerVerifyResult, Reason, NewCrlTemplate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_CertSign func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertTemplateToBeSigned unsafe.Pointer, SignerCertGroup unsafe.Pointer, SignerVerifyContext unsafe.Pointer, SignerVerifyResult unsafe.Pointer, SignedCert unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_CertSignErr error

// CanCallCSSM_TP_CertSign reports whether the symbol for CSSM_TP_CertSign is available on this system.
func CanCallCSSM_TP_CertSign() bool {
	return _cSSM_TP_CertSign != nil
}

// CSSM_TP_CertSignCallError returns the symbol lookup or registration error for CSSM_TP_CertSign.
func CSSM_TP_CertSignCallError() error {
	if _cSSM_TP_CertSign != nil {
		return nil
	}
	return symbolCallError("CSSM_TP_CertSign", "10.0", _cSSM_TP_CertSignErr)
}

// TryCSSM_TP_CertSign calls CSSM_TP_CertSign without panicking when the symbol is unavailable.
func TryCSSM_TP_CertSign(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertTemplateToBeSigned unsafe.Pointer, SignerCertGroup unsafe.Pointer, SignerVerifyContext unsafe.Pointer, SignerVerifyResult unsafe.Pointer, SignedCert unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_TP_CertSignCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_TP_CertSign(TPHandle, CLHandle, CCHandle, CertTemplateToBeSigned, SignerCertGroup, SignerVerifyContext, SignerVerifyResult, SignedCert), nil
}

// CSSM_TP_CertSign.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CertSign
func CSSM_TP_CertSign(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertTemplateToBeSigned unsafe.Pointer, SignerCertGroup unsafe.Pointer, SignerVerifyContext unsafe.Pointer, SignerVerifyResult unsafe.Pointer, SignedCert unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_TP_CertSign(TPHandle, CLHandle, CCHandle, CertTemplateToBeSigned, SignerCertGroup, SignerVerifyContext, SignerVerifyResult, SignedCert)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_ConfirmCredResult func(TPHandle CSSM_TP_HANDLE, ReferenceIdentifier unsafe.Pointer, CallerAuthCredentials unsafe.Pointer, Responses unsafe.Pointer, PreferredAuthority unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_ConfirmCredResultErr error

// CanCallCSSM_TP_ConfirmCredResult reports whether the symbol for CSSM_TP_ConfirmCredResult is available on this system.
func CanCallCSSM_TP_ConfirmCredResult() bool {
	return _cSSM_TP_ConfirmCredResult != nil
}

// CSSM_TP_ConfirmCredResultCallError returns the symbol lookup or registration error for CSSM_TP_ConfirmCredResult.
func CSSM_TP_ConfirmCredResultCallError() error {
	if _cSSM_TP_ConfirmCredResult != nil {
		return nil
	}
	return symbolCallError("CSSM_TP_ConfirmCredResult", "10.0", _cSSM_TP_ConfirmCredResultErr)
}

// TryCSSM_TP_ConfirmCredResult calls CSSM_TP_ConfirmCredResult without panicking when the symbol is unavailable.
func TryCSSM_TP_ConfirmCredResult(TPHandle CSSM_TP_HANDLE, ReferenceIdentifier unsafe.Pointer, CallerAuthCredentials unsafe.Pointer, Responses unsafe.Pointer, PreferredAuthority unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_TP_ConfirmCredResultCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_TP_ConfirmCredResult(TPHandle, ReferenceIdentifier, CallerAuthCredentials, Responses, PreferredAuthority), nil
}

// CSSM_TP_ConfirmCredResult.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_ConfirmCredResult
func CSSM_TP_ConfirmCredResult(TPHandle CSSM_TP_HANDLE, ReferenceIdentifier unsafe.Pointer, CallerAuthCredentials unsafe.Pointer, Responses unsafe.Pointer, PreferredAuthority unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_TP_ConfirmCredResult(TPHandle, ReferenceIdentifier, CallerAuthCredentials, Responses, PreferredAuthority)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_CrlCreateTemplate func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CrlFields unsafe.Pointer, NewCrlTemplate unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_CrlCreateTemplateErr error

// CanCallCSSM_TP_CrlCreateTemplate reports whether the symbol for CSSM_TP_CrlCreateTemplate is available on this system.
func CanCallCSSM_TP_CrlCreateTemplate() bool {
	return _cSSM_TP_CrlCreateTemplate != nil
}

// CSSM_TP_CrlCreateTemplateCallError returns the symbol lookup or registration error for CSSM_TP_CrlCreateTemplate.
func CSSM_TP_CrlCreateTemplateCallError() error {
	if _cSSM_TP_CrlCreateTemplate != nil {
		return nil
	}
	return symbolCallError("CSSM_TP_CrlCreateTemplate", "10.0", _cSSM_TP_CrlCreateTemplateErr)
}

// TryCSSM_TP_CrlCreateTemplate calls CSSM_TP_CrlCreateTemplate without panicking when the symbol is unavailable.
func TryCSSM_TP_CrlCreateTemplate(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CrlFields unsafe.Pointer, NewCrlTemplate unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_TP_CrlCreateTemplateCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_TP_CrlCreateTemplate(TPHandle, CLHandle, NumberOfFields, CrlFields, NewCrlTemplate), nil
}

// CSSM_TP_CrlCreateTemplate.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CrlCreateTemplate
func CSSM_TP_CrlCreateTemplate(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CrlFields unsafe.Pointer, NewCrlTemplate unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_TP_CrlCreateTemplate(TPHandle, CLHandle, NumberOfFields, CrlFields, NewCrlTemplate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_CrlSign func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CrlToBeSigned unsafe.Pointer, SignerCertGroup unsafe.Pointer, SignerVerifyContext unsafe.Pointer, SignerVerifyResult unsafe.Pointer, SignedCrl unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_CrlSignErr error

// CanCallCSSM_TP_CrlSign reports whether the symbol for CSSM_TP_CrlSign is available on this system.
func CanCallCSSM_TP_CrlSign() bool {
	return _cSSM_TP_CrlSign != nil
}

// CSSM_TP_CrlSignCallError returns the symbol lookup or registration error for CSSM_TP_CrlSign.
func CSSM_TP_CrlSignCallError() error {
	if _cSSM_TP_CrlSign != nil {
		return nil
	}
	return symbolCallError("CSSM_TP_CrlSign", "10.0", _cSSM_TP_CrlSignErr)
}

// TryCSSM_TP_CrlSign calls CSSM_TP_CrlSign without panicking when the symbol is unavailable.
func TryCSSM_TP_CrlSign(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CrlToBeSigned unsafe.Pointer, SignerCertGroup unsafe.Pointer, SignerVerifyContext unsafe.Pointer, SignerVerifyResult unsafe.Pointer, SignedCrl unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_TP_CrlSignCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_TP_CrlSign(TPHandle, CLHandle, CCHandle, CrlToBeSigned, SignerCertGroup, SignerVerifyContext, SignerVerifyResult, SignedCrl), nil
}

// CSSM_TP_CrlSign.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CrlSign
func CSSM_TP_CrlSign(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CrlToBeSigned unsafe.Pointer, SignerCertGroup unsafe.Pointer, SignerVerifyContext unsafe.Pointer, SignerVerifyResult unsafe.Pointer, SignedCrl unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_TP_CrlSign(TPHandle, CLHandle, CCHandle, CrlToBeSigned, SignerCertGroup, SignerVerifyContext, SignerVerifyResult, SignedCrl)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_CrlVerify func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, CrlToBeVerified unsafe.Pointer, SignerCertGroup unsafe.Pointer, VerifyContext unsafe.Pointer, RevokerVerifyResult unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_CrlVerifyErr error

// CanCallCSSM_TP_CrlVerify reports whether the symbol for CSSM_TP_CrlVerify is available on this system.
func CanCallCSSM_TP_CrlVerify() bool {
	return _cSSM_TP_CrlVerify != nil
}

// CSSM_TP_CrlVerifyCallError returns the symbol lookup or registration error for CSSM_TP_CrlVerify.
func CSSM_TP_CrlVerifyCallError() error {
	if _cSSM_TP_CrlVerify != nil {
		return nil
	}
	return symbolCallError("CSSM_TP_CrlVerify", "10.0", _cSSM_TP_CrlVerifyErr)
}

// TryCSSM_TP_CrlVerify calls CSSM_TP_CrlVerify without panicking when the symbol is unavailable.
func TryCSSM_TP_CrlVerify(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, CrlToBeVerified unsafe.Pointer, SignerCertGroup unsafe.Pointer, VerifyContext unsafe.Pointer, RevokerVerifyResult unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_TP_CrlVerifyCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_TP_CrlVerify(TPHandle, CLHandle, CSPHandle, CrlToBeVerified, SignerCertGroup, VerifyContext, RevokerVerifyResult), nil
}

// CSSM_TP_CrlVerify.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CrlVerify
func CSSM_TP_CrlVerify(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, CrlToBeVerified unsafe.Pointer, SignerCertGroup unsafe.Pointer, VerifyContext unsafe.Pointer, RevokerVerifyResult unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_TP_CrlVerify(TPHandle, CLHandle, CSPHandle, CrlToBeVerified, SignerCertGroup, VerifyContext, RevokerVerifyResult)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_FormRequest func(TPHandle CSSM_TP_HANDLE, PreferredAuthority unsafe.Pointer, FormType CSSM_TP_FORM_TYPE, BlankForm unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_FormRequestErr error

// CanCallCSSM_TP_FormRequest reports whether the symbol for CSSM_TP_FormRequest is available on this system.
func CanCallCSSM_TP_FormRequest() bool {
	return _cSSM_TP_FormRequest != nil
}

// CSSM_TP_FormRequestCallError returns the symbol lookup or registration error for CSSM_TP_FormRequest.
func CSSM_TP_FormRequestCallError() error {
	if _cSSM_TP_FormRequest != nil {
		return nil
	}
	return symbolCallError("CSSM_TP_FormRequest", "10.0", _cSSM_TP_FormRequestErr)
}

// TryCSSM_TP_FormRequest calls CSSM_TP_FormRequest without panicking when the symbol is unavailable.
func TryCSSM_TP_FormRequest(TPHandle CSSM_TP_HANDLE, PreferredAuthority unsafe.Pointer, FormType CSSM_TP_FORM_TYPE, BlankForm unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_TP_FormRequestCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_TP_FormRequest(TPHandle, PreferredAuthority, FormType, BlankForm), nil
}

// CSSM_TP_FormRequest.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_FormRequest
func CSSM_TP_FormRequest(TPHandle CSSM_TP_HANDLE, PreferredAuthority unsafe.Pointer, FormType CSSM_TP_FORM_TYPE, BlankForm unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_TP_FormRequest(TPHandle, PreferredAuthority, FormType, BlankForm)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_FormSubmit func(TPHandle CSSM_TP_HANDLE, FormType CSSM_TP_FORM_TYPE, Form unsafe.Pointer, ClearanceAuthority unsafe.Pointer, RepresentedAuthority unsafe.Pointer, Credentials unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_FormSubmitErr error

// CanCallCSSM_TP_FormSubmit reports whether the symbol for CSSM_TP_FormSubmit is available on this system.
func CanCallCSSM_TP_FormSubmit() bool {
	return _cSSM_TP_FormSubmit != nil
}

// CSSM_TP_FormSubmitCallError returns the symbol lookup or registration error for CSSM_TP_FormSubmit.
func CSSM_TP_FormSubmitCallError() error {
	if _cSSM_TP_FormSubmit != nil {
		return nil
	}
	return symbolCallError("CSSM_TP_FormSubmit", "10.0", _cSSM_TP_FormSubmitErr)
}

// TryCSSM_TP_FormSubmit calls CSSM_TP_FormSubmit without panicking when the symbol is unavailable.
func TryCSSM_TP_FormSubmit(TPHandle CSSM_TP_HANDLE, FormType CSSM_TP_FORM_TYPE, Form unsafe.Pointer, ClearanceAuthority unsafe.Pointer, RepresentedAuthority unsafe.Pointer, Credentials unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_TP_FormSubmitCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_TP_FormSubmit(TPHandle, FormType, Form, ClearanceAuthority, RepresentedAuthority, Credentials), nil
}

// CSSM_TP_FormSubmit.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_FormSubmit
func CSSM_TP_FormSubmit(TPHandle CSSM_TP_HANDLE, FormType CSSM_TP_FORM_TYPE, Form unsafe.Pointer, ClearanceAuthority unsafe.Pointer, RepresentedAuthority unsafe.Pointer, Credentials unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_TP_FormSubmit(TPHandle, FormType, Form, ClearanceAuthority, RepresentedAuthority, Credentials)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_PassThrough func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, DBList unsafe.Pointer, PassThroughId uint32, InputParams unsafe.Pointer, OutputParams unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_PassThroughErr error

// CanCallCSSM_TP_PassThrough reports whether the symbol for CSSM_TP_PassThrough is available on this system.
func CanCallCSSM_TP_PassThrough() bool {
	return _cSSM_TP_PassThrough != nil
}

// CSSM_TP_PassThroughCallError returns the symbol lookup or registration error for CSSM_TP_PassThrough.
func CSSM_TP_PassThroughCallError() error {
	if _cSSM_TP_PassThrough != nil {
		return nil
	}
	return symbolCallError("CSSM_TP_PassThrough", "10.0", _cSSM_TP_PassThroughErr)
}

// TryCSSM_TP_PassThrough calls CSSM_TP_PassThrough without panicking when the symbol is unavailable.
func TryCSSM_TP_PassThrough(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, DBList unsafe.Pointer, PassThroughId uint32, InputParams unsafe.Pointer, OutputParams unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_TP_PassThroughCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_TP_PassThrough(TPHandle, CLHandle, CCHandle, DBList, PassThroughId, InputParams, OutputParams), nil
}

// CSSM_TP_PassThrough.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_PassThrough
func CSSM_TP_PassThrough(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, DBList unsafe.Pointer, PassThroughId uint32, InputParams unsafe.Pointer, OutputParams unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_TP_PassThrough(TPHandle, CLHandle, CCHandle, DBList, PassThroughId, InputParams, OutputParams)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_ReceiveConfirmation func(TPHandle CSSM_TP_HANDLE, ReferenceIdentifier unsafe.Pointer, Responses unsafe.Pointer, ElapsedTime *int32) CSSM_RETURN
var _cSSM_TP_ReceiveConfirmationErr error

// CanCallCSSM_TP_ReceiveConfirmation reports whether the symbol for CSSM_TP_ReceiveConfirmation is available on this system.
func CanCallCSSM_TP_ReceiveConfirmation() bool {
	return _cSSM_TP_ReceiveConfirmation != nil
}

// CSSM_TP_ReceiveConfirmationCallError returns the symbol lookup or registration error for CSSM_TP_ReceiveConfirmation.
func CSSM_TP_ReceiveConfirmationCallError() error {
	if _cSSM_TP_ReceiveConfirmation != nil {
		return nil
	}
	return symbolCallError("CSSM_TP_ReceiveConfirmation", "10.0", _cSSM_TP_ReceiveConfirmationErr)
}

// TryCSSM_TP_ReceiveConfirmation calls CSSM_TP_ReceiveConfirmation without panicking when the symbol is unavailable.
func TryCSSM_TP_ReceiveConfirmation(TPHandle CSSM_TP_HANDLE, ReferenceIdentifier unsafe.Pointer, Responses unsafe.Pointer, ElapsedTime *int32) (CSSM_RETURN, error) {
	if err := CSSM_TP_ReceiveConfirmationCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_TP_ReceiveConfirmation(TPHandle, ReferenceIdentifier, Responses, ElapsedTime), nil
}

// CSSM_TP_ReceiveConfirmation.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_ReceiveConfirmation
func CSSM_TP_ReceiveConfirmation(TPHandle CSSM_TP_HANDLE, ReferenceIdentifier unsafe.Pointer, Responses unsafe.Pointer, ElapsedTime *int32) CSSM_RETURN {
	result, callErr := TryCSSM_TP_ReceiveConfirmation(TPHandle, ReferenceIdentifier, Responses, ElapsedTime)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_RetrieveCredResult func(TPHandle CSSM_TP_HANDLE, ReferenceIdentifier unsafe.Pointer, CallerAuthCredentials unsafe.Pointer, EstimatedTime *int32, ConfirmationRequired unsafe.Pointer, RetrieveOutput unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_RetrieveCredResultErr error

// CanCallCSSM_TP_RetrieveCredResult reports whether the symbol for CSSM_TP_RetrieveCredResult is available on this system.
func CanCallCSSM_TP_RetrieveCredResult() bool {
	return _cSSM_TP_RetrieveCredResult != nil
}

// CSSM_TP_RetrieveCredResultCallError returns the symbol lookup or registration error for CSSM_TP_RetrieveCredResult.
func CSSM_TP_RetrieveCredResultCallError() error {
	if _cSSM_TP_RetrieveCredResult != nil {
		return nil
	}
	return symbolCallError("CSSM_TP_RetrieveCredResult", "10.0", _cSSM_TP_RetrieveCredResultErr)
}

// TryCSSM_TP_RetrieveCredResult calls CSSM_TP_RetrieveCredResult without panicking when the symbol is unavailable.
func TryCSSM_TP_RetrieveCredResult(TPHandle CSSM_TP_HANDLE, ReferenceIdentifier unsafe.Pointer, CallerAuthCredentials unsafe.Pointer, EstimatedTime *int32, ConfirmationRequired unsafe.Pointer, RetrieveOutput unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_TP_RetrieveCredResultCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_TP_RetrieveCredResult(TPHandle, ReferenceIdentifier, CallerAuthCredentials, EstimatedTime, ConfirmationRequired, RetrieveOutput), nil
}

// CSSM_TP_RetrieveCredResult.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_RetrieveCredResult
func CSSM_TP_RetrieveCredResult(TPHandle CSSM_TP_HANDLE, ReferenceIdentifier unsafe.Pointer, CallerAuthCredentials unsafe.Pointer, EstimatedTime *int32, ConfirmationRequired unsafe.Pointer, RetrieveOutput unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_TP_RetrieveCredResult(TPHandle, ReferenceIdentifier, CallerAuthCredentials, EstimatedTime, ConfirmationRequired, RetrieveOutput)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_SubmitCredRequest func(TPHandle CSSM_TP_HANDLE, PreferredAuthority unsafe.Pointer, RequestType CSSM_TP_AUTHORITY_REQUEST_TYPE, RequestInput unsafe.Pointer, CallerAuthContext unsafe.Pointer, EstimatedTime *int32, ReferenceIdentifier unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_SubmitCredRequestErr error

// CanCallCSSM_TP_SubmitCredRequest reports whether the symbol for CSSM_TP_SubmitCredRequest is available on this system.
func CanCallCSSM_TP_SubmitCredRequest() bool {
	return _cSSM_TP_SubmitCredRequest != nil
}

// CSSM_TP_SubmitCredRequestCallError returns the symbol lookup or registration error for CSSM_TP_SubmitCredRequest.
func CSSM_TP_SubmitCredRequestCallError() error {
	if _cSSM_TP_SubmitCredRequest != nil {
		return nil
	}
	return symbolCallError("CSSM_TP_SubmitCredRequest", "10.0", _cSSM_TP_SubmitCredRequestErr)
}

// TryCSSM_TP_SubmitCredRequest calls CSSM_TP_SubmitCredRequest without panicking when the symbol is unavailable.
func TryCSSM_TP_SubmitCredRequest(TPHandle CSSM_TP_HANDLE, PreferredAuthority unsafe.Pointer, RequestType CSSM_TP_AUTHORITY_REQUEST_TYPE, RequestInput unsafe.Pointer, CallerAuthContext unsafe.Pointer, EstimatedTime *int32, ReferenceIdentifier unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_TP_SubmitCredRequestCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_TP_SubmitCredRequest(TPHandle, PreferredAuthority, RequestType, RequestInput, CallerAuthContext, EstimatedTime, ReferenceIdentifier), nil
}

// CSSM_TP_SubmitCredRequest.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_SubmitCredRequest
func CSSM_TP_SubmitCredRequest(TPHandle CSSM_TP_HANDLE, PreferredAuthority unsafe.Pointer, RequestType CSSM_TP_AUTHORITY_REQUEST_TYPE, RequestInput unsafe.Pointer, CallerAuthContext unsafe.Pointer, EstimatedTime *int32, ReferenceIdentifier unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_TP_SubmitCredRequest(TPHandle, PreferredAuthority, RequestType, RequestInput, CallerAuthContext, EstimatedTime, ReferenceIdentifier)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_TP_TupleGroupToCertGroup func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, TupleGroup unsafe.Pointer, CertTemplates unsafe.Pointer) CSSM_RETURN
var _cSSM_TP_TupleGroupToCertGroupErr error

// CanCallCSSM_TP_TupleGroupToCertGroup reports whether the symbol for CSSM_TP_TupleGroupToCertGroup is available on this system.
func CanCallCSSM_TP_TupleGroupToCertGroup() bool {
	return _cSSM_TP_TupleGroupToCertGroup != nil
}

// CSSM_TP_TupleGroupToCertGroupCallError returns the symbol lookup or registration error for CSSM_TP_TupleGroupToCertGroup.
func CSSM_TP_TupleGroupToCertGroupCallError() error {
	if _cSSM_TP_TupleGroupToCertGroup != nil {
		return nil
	}
	return symbolCallError("CSSM_TP_TupleGroupToCertGroup", "10.0", _cSSM_TP_TupleGroupToCertGroupErr)
}

// TryCSSM_TP_TupleGroupToCertGroup calls CSSM_TP_TupleGroupToCertGroup without panicking when the symbol is unavailable.
func TryCSSM_TP_TupleGroupToCertGroup(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, TupleGroup unsafe.Pointer, CertTemplates unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_TP_TupleGroupToCertGroupCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_TP_TupleGroupToCertGroup(TPHandle, CLHandle, TupleGroup, CertTemplates), nil
}

// CSSM_TP_TupleGroupToCertGroup.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_TupleGroupToCertGroup
func CSSM_TP_TupleGroupToCertGroup(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, TupleGroup unsafe.Pointer, CertTemplates unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_TP_TupleGroupToCertGroup(TPHandle, CLHandle, TupleGroup, CertTemplates)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_Terminate func() CSSM_RETURN
var _cSSM_TerminateErr error

// CanCallCSSM_Terminate reports whether the symbol for CSSM_Terminate is available on this system.
func CanCallCSSM_Terminate() bool {
	return _cSSM_Terminate != nil
}

// CSSM_TerminateCallError returns the symbol lookup or registration error for CSSM_Terminate.
func CSSM_TerminateCallError() error {
	if _cSSM_Terminate != nil {
		return nil
	}
	return symbolCallError("CSSM_Terminate", "10.0", _cSSM_TerminateErr)
}

// TryCSSM_Terminate calls CSSM_Terminate without panicking when the symbol is unavailable.
func TryCSSM_Terminate() (CSSM_RETURN, error) {
	if err := CSSM_TerminateCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_Terminate(), nil
}

// CSSM_Terminate.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_Terminate
func CSSM_Terminate() CSSM_RETURN {
	result, callErr := TryCSSM_Terminate()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_Unintroduce func(ModuleID unsafe.Pointer) CSSM_RETURN
var _cSSM_UnintroduceErr error

// CanCallCSSM_Unintroduce reports whether the symbol for CSSM_Unintroduce is available on this system.
func CanCallCSSM_Unintroduce() bool {
	return _cSSM_Unintroduce != nil
}

// CSSM_UnintroduceCallError returns the symbol lookup or registration error for CSSM_Unintroduce.
func CSSM_UnintroduceCallError() error {
	if _cSSM_Unintroduce != nil {
		return nil
	}
	return symbolCallError("CSSM_Unintroduce", "10.0", _cSSM_UnintroduceErr)
}

// TryCSSM_Unintroduce calls CSSM_Unintroduce without panicking when the symbol is unavailable.
func TryCSSM_Unintroduce(ModuleID unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_UnintroduceCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_Unintroduce(ModuleID), nil
}

// CSSM_Unintroduce.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_Unintroduce
func CSSM_Unintroduce(ModuleID unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_Unintroduce(ModuleID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_UnwrapKey func(CCHandle CSSM_CC_HANDLE, PublicKey unsafe.Pointer, WrappedKey unsafe.Pointer, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, UnwrappedKey unsafe.Pointer, DescriptiveData unsafe.Pointer) CSSM_RETURN
var _cSSM_UnwrapKeyErr error

// CanCallCSSM_UnwrapKey reports whether the symbol for CSSM_UnwrapKey is available on this system.
func CanCallCSSM_UnwrapKey() bool {
	return _cSSM_UnwrapKey != nil
}

// CSSM_UnwrapKeyCallError returns the symbol lookup or registration error for CSSM_UnwrapKey.
func CSSM_UnwrapKeyCallError() error {
	if _cSSM_UnwrapKey != nil {
		return nil
	}
	return symbolCallError("CSSM_UnwrapKey", "10.0", _cSSM_UnwrapKeyErr)
}

// TryCSSM_UnwrapKey calls CSSM_UnwrapKey without panicking when the symbol is unavailable.
func TryCSSM_UnwrapKey(CCHandle CSSM_CC_HANDLE, PublicKey unsafe.Pointer, WrappedKey unsafe.Pointer, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, UnwrappedKey unsafe.Pointer, DescriptiveData unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_UnwrapKeyCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_UnwrapKey(CCHandle, PublicKey, WrappedKey, KeyUsage, KeyAttr, KeyLabel, CredAndAclEntry, UnwrappedKey, DescriptiveData), nil
}

// CSSM_UnwrapKey.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_UnwrapKey
func CSSM_UnwrapKey(CCHandle CSSM_CC_HANDLE, PublicKey unsafe.Pointer, WrappedKey unsafe.Pointer, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, UnwrappedKey unsafe.Pointer, DescriptiveData unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_UnwrapKey(CCHandle, PublicKey, WrappedKey, KeyUsage, KeyAttr, KeyLabel, CredAndAclEntry, UnwrappedKey, DescriptiveData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_UnwrapKeyP func(CCHandle CSSM_CC_HANDLE, PublicKey unsafe.Pointer, WrappedKey unsafe.Pointer, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, UnwrappedKey unsafe.Pointer, DescriptiveData unsafe.Pointer, Privilege CSSM_PRIVILEGE) CSSM_RETURN
var _cSSM_UnwrapKeyPErr error

// CanCallCSSM_UnwrapKeyP reports whether the symbol for CSSM_UnwrapKeyP is available on this system.
func CanCallCSSM_UnwrapKeyP() bool {
	return _cSSM_UnwrapKeyP != nil
}

// CSSM_UnwrapKeyPCallError returns the symbol lookup or registration error for CSSM_UnwrapKeyP.
func CSSM_UnwrapKeyPCallError() error {
	if _cSSM_UnwrapKeyP != nil {
		return nil
	}
	return symbolCallError("CSSM_UnwrapKeyP", "10.0", _cSSM_UnwrapKeyPErr)
}

// TryCSSM_UnwrapKeyP calls CSSM_UnwrapKeyP without panicking when the symbol is unavailable.
func TryCSSM_UnwrapKeyP(CCHandle CSSM_CC_HANDLE, PublicKey unsafe.Pointer, WrappedKey unsafe.Pointer, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, UnwrappedKey unsafe.Pointer, DescriptiveData unsafe.Pointer, Privilege CSSM_PRIVILEGE) (CSSM_RETURN, error) {
	if err := CSSM_UnwrapKeyPCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_UnwrapKeyP(CCHandle, PublicKey, WrappedKey, KeyUsage, KeyAttr, KeyLabel, CredAndAclEntry, UnwrappedKey, DescriptiveData, Privilege), nil
}

// CSSM_UnwrapKeyP.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_UnwrapKeyP
func CSSM_UnwrapKeyP(CCHandle CSSM_CC_HANDLE, PublicKey unsafe.Pointer, WrappedKey unsafe.Pointer, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, UnwrappedKey unsafe.Pointer, DescriptiveData unsafe.Pointer, Privilege CSSM_PRIVILEGE) CSSM_RETURN {
	result, callErr := TryCSSM_UnwrapKeyP(CCHandle, PublicKey, WrappedKey, KeyUsage, KeyAttr, KeyLabel, CredAndAclEntry, UnwrappedKey, DescriptiveData, Privilege)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_UpdateContextAttributes func(CCHandle CSSM_CC_HANDLE, NumberOfAttributes uint32, ContextAttributes unsafe.Pointer) CSSM_RETURN
var _cSSM_UpdateContextAttributesErr error

// CanCallCSSM_UpdateContextAttributes reports whether the symbol for CSSM_UpdateContextAttributes is available on this system.
func CanCallCSSM_UpdateContextAttributes() bool {
	return _cSSM_UpdateContextAttributes != nil
}

// CSSM_UpdateContextAttributesCallError returns the symbol lookup or registration error for CSSM_UpdateContextAttributes.
func CSSM_UpdateContextAttributesCallError() error {
	if _cSSM_UpdateContextAttributes != nil {
		return nil
	}
	return symbolCallError("CSSM_UpdateContextAttributes", "10.0", _cSSM_UpdateContextAttributesErr)
}

// TryCSSM_UpdateContextAttributes calls CSSM_UpdateContextAttributes without panicking when the symbol is unavailable.
func TryCSSM_UpdateContextAttributes(CCHandle CSSM_CC_HANDLE, NumberOfAttributes uint32, ContextAttributes unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_UpdateContextAttributesCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_UpdateContextAttributes(CCHandle, NumberOfAttributes, ContextAttributes), nil
}

// CSSM_UpdateContextAttributes.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_UpdateContextAttributes
func CSSM_UpdateContextAttributes(CCHandle CSSM_CC_HANDLE, NumberOfAttributes uint32, ContextAttributes unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_UpdateContextAttributes(CCHandle, NumberOfAttributes, ContextAttributes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_VerifyData func(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, DigestAlgorithm CSSM_ALGORITHMS, Signature unsafe.Pointer) CSSM_RETURN
var _cSSM_VerifyDataErr error

// CanCallCSSM_VerifyData reports whether the symbol for CSSM_VerifyData is available on this system.
func CanCallCSSM_VerifyData() bool {
	return _cSSM_VerifyData != nil
}

// CSSM_VerifyDataCallError returns the symbol lookup or registration error for CSSM_VerifyData.
func CSSM_VerifyDataCallError() error {
	if _cSSM_VerifyData != nil {
		return nil
	}
	return symbolCallError("CSSM_VerifyData", "10.0", _cSSM_VerifyDataErr)
}

// TryCSSM_VerifyData calls CSSM_VerifyData without panicking when the symbol is unavailable.
func TryCSSM_VerifyData(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, DigestAlgorithm CSSM_ALGORITHMS, Signature unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_VerifyDataCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_VerifyData(CCHandle, DataBufs, DataBufCount, DigestAlgorithm, Signature), nil
}

// CSSM_VerifyData.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_VerifyData
func CSSM_VerifyData(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, DigestAlgorithm CSSM_ALGORITHMS, Signature unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_VerifyData(CCHandle, DataBufs, DataBufCount, DigestAlgorithm, Signature)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_VerifyDataFinal func(CCHandle CSSM_CC_HANDLE, Signature unsafe.Pointer) CSSM_RETURN
var _cSSM_VerifyDataFinalErr error

// CanCallCSSM_VerifyDataFinal reports whether the symbol for CSSM_VerifyDataFinal is available on this system.
func CanCallCSSM_VerifyDataFinal() bool {
	return _cSSM_VerifyDataFinal != nil
}

// CSSM_VerifyDataFinalCallError returns the symbol lookup or registration error for CSSM_VerifyDataFinal.
func CSSM_VerifyDataFinalCallError() error {
	if _cSSM_VerifyDataFinal != nil {
		return nil
	}
	return symbolCallError("CSSM_VerifyDataFinal", "10.0", _cSSM_VerifyDataFinalErr)
}

// TryCSSM_VerifyDataFinal calls CSSM_VerifyDataFinal without panicking when the symbol is unavailable.
func TryCSSM_VerifyDataFinal(CCHandle CSSM_CC_HANDLE, Signature unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_VerifyDataFinalCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_VerifyDataFinal(CCHandle, Signature), nil
}

// CSSM_VerifyDataFinal.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_VerifyDataFinal
func CSSM_VerifyDataFinal(CCHandle CSSM_CC_HANDLE, Signature unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_VerifyDataFinal(CCHandle, Signature)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_VerifyDataInit func(CCHandle CSSM_CC_HANDLE) CSSM_RETURN
var _cSSM_VerifyDataInitErr error

// CanCallCSSM_VerifyDataInit reports whether the symbol for CSSM_VerifyDataInit is available on this system.
func CanCallCSSM_VerifyDataInit() bool {
	return _cSSM_VerifyDataInit != nil
}

// CSSM_VerifyDataInitCallError returns the symbol lookup or registration error for CSSM_VerifyDataInit.
func CSSM_VerifyDataInitCallError() error {
	if _cSSM_VerifyDataInit != nil {
		return nil
	}
	return symbolCallError("CSSM_VerifyDataInit", "10.0", _cSSM_VerifyDataInitErr)
}

// TryCSSM_VerifyDataInit calls CSSM_VerifyDataInit without panicking when the symbol is unavailable.
func TryCSSM_VerifyDataInit(CCHandle CSSM_CC_HANDLE) (CSSM_RETURN, error) {
	if err := CSSM_VerifyDataInitCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_VerifyDataInit(CCHandle), nil
}

// CSSM_VerifyDataInit.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_VerifyDataInit
func CSSM_VerifyDataInit(CCHandle CSSM_CC_HANDLE) CSSM_RETURN {
	result, callErr := TryCSSM_VerifyDataInit(CCHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_VerifyDataUpdate func(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) CSSM_RETURN
var _cSSM_VerifyDataUpdateErr error

// CanCallCSSM_VerifyDataUpdate reports whether the symbol for CSSM_VerifyDataUpdate is available on this system.
func CanCallCSSM_VerifyDataUpdate() bool {
	return _cSSM_VerifyDataUpdate != nil
}

// CSSM_VerifyDataUpdateCallError returns the symbol lookup or registration error for CSSM_VerifyDataUpdate.
func CSSM_VerifyDataUpdateCallError() error {
	if _cSSM_VerifyDataUpdate != nil {
		return nil
	}
	return symbolCallError("CSSM_VerifyDataUpdate", "10.0", _cSSM_VerifyDataUpdateErr)
}

// TryCSSM_VerifyDataUpdate calls CSSM_VerifyDataUpdate without panicking when the symbol is unavailable.
func TryCSSM_VerifyDataUpdate(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) (CSSM_RETURN, error) {
	if err := CSSM_VerifyDataUpdateCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_VerifyDataUpdate(CCHandle, DataBufs, DataBufCount), nil
}

// CSSM_VerifyDataUpdate.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_VerifyDataUpdate
func CSSM_VerifyDataUpdate(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) CSSM_RETURN {
	result, callErr := TryCSSM_VerifyDataUpdate(CCHandle, DataBufs, DataBufCount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_VerifyDevice func(CSPHandle CSSM_CSP_HANDLE, DeviceCert unsafe.Pointer) CSSM_RETURN
var _cSSM_VerifyDeviceErr error

// CanCallCSSM_VerifyDevice reports whether the symbol for CSSM_VerifyDevice is available on this system.
func CanCallCSSM_VerifyDevice() bool {
	return _cSSM_VerifyDevice != nil
}

// CSSM_VerifyDeviceCallError returns the symbol lookup or registration error for CSSM_VerifyDevice.
func CSSM_VerifyDeviceCallError() error {
	if _cSSM_VerifyDevice != nil {
		return nil
	}
	return symbolCallError("CSSM_VerifyDevice", "10.0", _cSSM_VerifyDeviceErr)
}

// TryCSSM_VerifyDevice calls CSSM_VerifyDevice without panicking when the symbol is unavailable.
func TryCSSM_VerifyDevice(CSPHandle CSSM_CSP_HANDLE, DeviceCert unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_VerifyDeviceCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_VerifyDevice(CSPHandle, DeviceCert), nil
}

// CSSM_VerifyDevice.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_VerifyDevice
func CSSM_VerifyDevice(CSPHandle CSSM_CSP_HANDLE, DeviceCert unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_VerifyDevice(CSPHandle, DeviceCert)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_VerifyMac func(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, Mac unsafe.Pointer) CSSM_RETURN
var _cSSM_VerifyMacErr error

// CanCallCSSM_VerifyMac reports whether the symbol for CSSM_VerifyMac is available on this system.
func CanCallCSSM_VerifyMac() bool {
	return _cSSM_VerifyMac != nil
}

// CSSM_VerifyMacCallError returns the symbol lookup or registration error for CSSM_VerifyMac.
func CSSM_VerifyMacCallError() error {
	if _cSSM_VerifyMac != nil {
		return nil
	}
	return symbolCallError("CSSM_VerifyMac", "10.0", _cSSM_VerifyMacErr)
}

// TryCSSM_VerifyMac calls CSSM_VerifyMac without panicking when the symbol is unavailable.
func TryCSSM_VerifyMac(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, Mac unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_VerifyMacCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_VerifyMac(CCHandle, DataBufs, DataBufCount, Mac), nil
}

// CSSM_VerifyMac.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_VerifyMac
func CSSM_VerifyMac(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, Mac unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_VerifyMac(CCHandle, DataBufs, DataBufCount, Mac)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_VerifyMacFinal func(CCHandle CSSM_CC_HANDLE, Mac unsafe.Pointer) CSSM_RETURN
var _cSSM_VerifyMacFinalErr error

// CanCallCSSM_VerifyMacFinal reports whether the symbol for CSSM_VerifyMacFinal is available on this system.
func CanCallCSSM_VerifyMacFinal() bool {
	return _cSSM_VerifyMacFinal != nil
}

// CSSM_VerifyMacFinalCallError returns the symbol lookup or registration error for CSSM_VerifyMacFinal.
func CSSM_VerifyMacFinalCallError() error {
	if _cSSM_VerifyMacFinal != nil {
		return nil
	}
	return symbolCallError("CSSM_VerifyMacFinal", "10.0", _cSSM_VerifyMacFinalErr)
}

// TryCSSM_VerifyMacFinal calls CSSM_VerifyMacFinal without panicking when the symbol is unavailable.
func TryCSSM_VerifyMacFinal(CCHandle CSSM_CC_HANDLE, Mac unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_VerifyMacFinalCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_VerifyMacFinal(CCHandle, Mac), nil
}

// CSSM_VerifyMacFinal.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_VerifyMacFinal
func CSSM_VerifyMacFinal(CCHandle CSSM_CC_HANDLE, Mac unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_VerifyMacFinal(CCHandle, Mac)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_VerifyMacInit func(CCHandle CSSM_CC_HANDLE) CSSM_RETURN
var _cSSM_VerifyMacInitErr error

// CanCallCSSM_VerifyMacInit reports whether the symbol for CSSM_VerifyMacInit is available on this system.
func CanCallCSSM_VerifyMacInit() bool {
	return _cSSM_VerifyMacInit != nil
}

// CSSM_VerifyMacInitCallError returns the symbol lookup or registration error for CSSM_VerifyMacInit.
func CSSM_VerifyMacInitCallError() error {
	if _cSSM_VerifyMacInit != nil {
		return nil
	}
	return symbolCallError("CSSM_VerifyMacInit", "10.0", _cSSM_VerifyMacInitErr)
}

// TryCSSM_VerifyMacInit calls CSSM_VerifyMacInit without panicking when the symbol is unavailable.
func TryCSSM_VerifyMacInit(CCHandle CSSM_CC_HANDLE) (CSSM_RETURN, error) {
	if err := CSSM_VerifyMacInitCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_VerifyMacInit(CCHandle), nil
}

// CSSM_VerifyMacInit.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_VerifyMacInit
func CSSM_VerifyMacInit(CCHandle CSSM_CC_HANDLE) CSSM_RETURN {
	result, callErr := TryCSSM_VerifyMacInit(CCHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_VerifyMacUpdate func(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) CSSM_RETURN
var _cSSM_VerifyMacUpdateErr error

// CanCallCSSM_VerifyMacUpdate reports whether the symbol for CSSM_VerifyMacUpdate is available on this system.
func CanCallCSSM_VerifyMacUpdate() bool {
	return _cSSM_VerifyMacUpdate != nil
}

// CSSM_VerifyMacUpdateCallError returns the symbol lookup or registration error for CSSM_VerifyMacUpdate.
func CSSM_VerifyMacUpdateCallError() error {
	if _cSSM_VerifyMacUpdate != nil {
		return nil
	}
	return symbolCallError("CSSM_VerifyMacUpdate", "10.0", _cSSM_VerifyMacUpdateErr)
}

// TryCSSM_VerifyMacUpdate calls CSSM_VerifyMacUpdate without panicking when the symbol is unavailable.
func TryCSSM_VerifyMacUpdate(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) (CSSM_RETURN, error) {
	if err := CSSM_VerifyMacUpdateCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_VerifyMacUpdate(CCHandle, DataBufs, DataBufCount), nil
}

// CSSM_VerifyMacUpdate.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_VerifyMacUpdate
func CSSM_VerifyMacUpdate(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) CSSM_RETURN {
	result, callErr := TryCSSM_VerifyMacUpdate(CCHandle, DataBufs, DataBufCount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_WrapKey func(CCHandle CSSM_CC_HANDLE, AccessCred unsafe.Pointer, Key unsafe.Pointer, DescriptiveData unsafe.Pointer, WrappedKey unsafe.Pointer) CSSM_RETURN
var _cSSM_WrapKeyErr error

// CanCallCSSM_WrapKey reports whether the symbol for CSSM_WrapKey is available on this system.
func CanCallCSSM_WrapKey() bool {
	return _cSSM_WrapKey != nil
}

// CSSM_WrapKeyCallError returns the symbol lookup or registration error for CSSM_WrapKey.
func CSSM_WrapKeyCallError() error {
	if _cSSM_WrapKey != nil {
		return nil
	}
	return symbolCallError("CSSM_WrapKey", "10.0", _cSSM_WrapKeyErr)
}

// TryCSSM_WrapKey calls CSSM_WrapKey without panicking when the symbol is unavailable.
func TryCSSM_WrapKey(CCHandle CSSM_CC_HANDLE, AccessCred unsafe.Pointer, Key unsafe.Pointer, DescriptiveData unsafe.Pointer, WrappedKey unsafe.Pointer) (CSSM_RETURN, error) {
	if err := CSSM_WrapKeyCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_WrapKey(CCHandle, AccessCred, Key, DescriptiveData, WrappedKey), nil
}

// CSSM_WrapKey.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_WrapKey
func CSSM_WrapKey(CCHandle CSSM_CC_HANDLE, AccessCred unsafe.Pointer, Key unsafe.Pointer, DescriptiveData unsafe.Pointer, WrappedKey unsafe.Pointer) CSSM_RETURN {
	result, callErr := TryCSSM_WrapKey(CCHandle, AccessCred, Key, DescriptiveData, WrappedKey)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cSSM_WrapKeyP func(CCHandle CSSM_CC_HANDLE, AccessCred unsafe.Pointer, Key unsafe.Pointer, DescriptiveData unsafe.Pointer, WrappedKey unsafe.Pointer, Privilege CSSM_PRIVILEGE) CSSM_RETURN
var _cSSM_WrapKeyPErr error

// CanCallCSSM_WrapKeyP reports whether the symbol for CSSM_WrapKeyP is available on this system.
func CanCallCSSM_WrapKeyP() bool {
	return _cSSM_WrapKeyP != nil
}

// CSSM_WrapKeyPCallError returns the symbol lookup or registration error for CSSM_WrapKeyP.
func CSSM_WrapKeyPCallError() error {
	if _cSSM_WrapKeyP != nil {
		return nil
	}
	return symbolCallError("CSSM_WrapKeyP", "10.0", _cSSM_WrapKeyPErr)
}

// TryCSSM_WrapKeyP calls CSSM_WrapKeyP without panicking when the symbol is unavailable.
func TryCSSM_WrapKeyP(CCHandle CSSM_CC_HANDLE, AccessCred unsafe.Pointer, Key unsafe.Pointer, DescriptiveData unsafe.Pointer, WrappedKey unsafe.Pointer, Privilege CSSM_PRIVILEGE) (CSSM_RETURN, error) {
	if err := CSSM_WrapKeyPCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _cSSM_WrapKeyP(CCHandle, AccessCred, Key, DescriptiveData, WrappedKey, Privilege), nil
}

// CSSM_WrapKeyP.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CSSM_WrapKeyP
func CSSM_WrapKeyP(CCHandle CSSM_CC_HANDLE, AccessCred unsafe.Pointer, Key unsafe.Pointer, DescriptiveData unsafe.Pointer, WrappedKey unsafe.Pointer, Privilege CSSM_PRIVILEGE) CSSM_RETURN {
	result, callErr := TryCSSM_WrapKeyP(CCHandle, AccessCred, Key, DescriptiveData, WrappedKey, Privilege)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mDS_Initialize func(pCallerGuid unsafe.Pointer, pMemoryFunctions unsafe.Pointer, pDlFunctions unsafe.Pointer, hMds *MDS_HANDLE) CSSM_RETURN
var _mDS_InitializeErr error

// CanCallMDS_Initialize reports whether the symbol for MDS_Initialize is available on this system.
func CanCallMDS_Initialize() bool {
	return _mDS_Initialize != nil
}

// MDS_InitializeCallError returns the symbol lookup or registration error for MDS_Initialize.
func MDS_InitializeCallError() error {
	if _mDS_Initialize != nil {
		return nil
	}
	return symbolCallError("MDS_Initialize", "10.0", _mDS_InitializeErr)
}

// TryMDS_Initialize calls MDS_Initialize without panicking when the symbol is unavailable.
func TryMDS_Initialize(pCallerGuid unsafe.Pointer, pMemoryFunctions unsafe.Pointer, pDlFunctions unsafe.Pointer, hMds *MDS_HANDLE) (CSSM_RETURN, error) {
	if err := MDS_InitializeCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _mDS_Initialize(pCallerGuid, pMemoryFunctions, pDlFunctions, hMds), nil
}

// MDS_Initialize.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/MDS_Initialize
func MDS_Initialize(pCallerGuid unsafe.Pointer, pMemoryFunctions unsafe.Pointer, pDlFunctions unsafe.Pointer, hMds *MDS_HANDLE) CSSM_RETURN {
	result, callErr := TryMDS_Initialize(pCallerGuid, pMemoryFunctions, pDlFunctions, hMds)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mDS_Install func(MdsHandle MDS_HANDLE) CSSM_RETURN
var _mDS_InstallErr error

// CanCallMDS_Install reports whether the symbol for MDS_Install is available on this system.
func CanCallMDS_Install() bool {
	return _mDS_Install != nil
}

// MDS_InstallCallError returns the symbol lookup or registration error for MDS_Install.
func MDS_InstallCallError() error {
	if _mDS_Install != nil {
		return nil
	}
	return symbolCallError("MDS_Install", "10.0", _mDS_InstallErr)
}

// TryMDS_Install calls MDS_Install without panicking when the symbol is unavailable.
func TryMDS_Install(MdsHandle MDS_HANDLE) (CSSM_RETURN, error) {
	if err := MDS_InstallCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _mDS_Install(MdsHandle), nil
}

// MDS_Install.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/MDS_Install
func MDS_Install(MdsHandle MDS_HANDLE) CSSM_RETURN {
	result, callErr := TryMDS_Install(MdsHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mDS_Terminate func(MdsHandle MDS_HANDLE) CSSM_RETURN
var _mDS_TerminateErr error

// CanCallMDS_Terminate reports whether the symbol for MDS_Terminate is available on this system.
func CanCallMDS_Terminate() bool {
	return _mDS_Terminate != nil
}

// MDS_TerminateCallError returns the symbol lookup or registration error for MDS_Terminate.
func MDS_TerminateCallError() error {
	if _mDS_Terminate != nil {
		return nil
	}
	return symbolCallError("MDS_Terminate", "10.0", _mDS_TerminateErr)
}

// TryMDS_Terminate calls MDS_Terminate without panicking when the symbol is unavailable.
func TryMDS_Terminate(MdsHandle MDS_HANDLE) (CSSM_RETURN, error) {
	if err := MDS_TerminateCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _mDS_Terminate(MdsHandle), nil
}

// MDS_Terminate.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/MDS_Terminate
func MDS_Terminate(MdsHandle MDS_HANDLE) CSSM_RETURN {
	result, callErr := TryMDS_Terminate(MdsHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mDS_Uninstall func(MdsHandle MDS_HANDLE) CSSM_RETURN
var _mDS_UninstallErr error

// CanCallMDS_Uninstall reports whether the symbol for MDS_Uninstall is available on this system.
func CanCallMDS_Uninstall() bool {
	return _mDS_Uninstall != nil
}

// MDS_UninstallCallError returns the symbol lookup or registration error for MDS_Uninstall.
func MDS_UninstallCallError() error {
	if _mDS_Uninstall != nil {
		return nil
	}
	return symbolCallError("MDS_Uninstall", "10.0", _mDS_UninstallErr)
}

// TryMDS_Uninstall calls MDS_Uninstall without panicking when the symbol is unavailable.
func TryMDS_Uninstall(MdsHandle MDS_HANDLE) (CSSM_RETURN, error) {
	if err := MDS_UninstallCallError(); err != nil {
		return *new(CSSM_RETURN), err
	}
	return _mDS_Uninstall(MdsHandle), nil
}

// MDS_Uninstall.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/MDS_Uninstall
func MDS_Uninstall(MdsHandle MDS_HANDLE) CSSM_RETURN {
	result, callErr := TryMDS_Uninstall(MdsHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secAccessControlCreateWithFlags func(allocator corefoundation.CFAllocatorRef, protection corefoundation.CFTypeRef, flags SecAccessControlCreateFlags, err *corefoundation.CFErrorRef) SecAccessControlRef
var _secAccessControlCreateWithFlagsErr error

// CanCallSecAccessControlCreateWithFlags reports whether the symbol for SecAccessControlCreateWithFlags is available on this system.
func CanCallSecAccessControlCreateWithFlags() bool {
	return _secAccessControlCreateWithFlags != nil
}

// SecAccessControlCreateWithFlagsCallError returns the symbol lookup or registration error for SecAccessControlCreateWithFlags.
func SecAccessControlCreateWithFlagsCallError() error {
	if _secAccessControlCreateWithFlags != nil {
		return nil
	}
	return symbolCallError("SecAccessControlCreateWithFlags", "10.10", _secAccessControlCreateWithFlagsErr)
}

// TrySecAccessControlCreateWithFlags calls SecAccessControlCreateWithFlags without panicking when the symbol is unavailable.
func TrySecAccessControlCreateWithFlags(allocator corefoundation.CFAllocatorRef, protection corefoundation.CFTypeRef, flags SecAccessControlCreateFlags, err *corefoundation.CFErrorRef) (SecAccessControlRef, error) {
	if err := SecAccessControlCreateWithFlagsCallError(); err != nil {
		return 0, err
	}
	return _secAccessControlCreateWithFlags(allocator, protection, flags, err), nil
}

// SecAccessControlCreateWithFlags creates a new access control object with the specified protection type and flags.
//
// See: https://developer.apple.com/documentation/Security/SecAccessControlCreateWithFlags(_:_:_:_:)
func SecAccessControlCreateWithFlags(allocator corefoundation.CFAllocatorRef, protection corefoundation.CFTypeRef, flags SecAccessControlCreateFlags, err *corefoundation.CFErrorRef) SecAccessControlRef {
	result, callErr := TrySecAccessControlCreateWithFlags(allocator, protection, flags, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secAccessControlGetTypeID func() uint
var _secAccessControlGetTypeIDErr error

// CanCallSecAccessControlGetTypeID reports whether the symbol for SecAccessControlGetTypeID is available on this system.
func CanCallSecAccessControlGetTypeID() bool {
	return _secAccessControlGetTypeID != nil
}

// SecAccessControlGetTypeIDCallError returns the symbol lookup or registration error for SecAccessControlGetTypeID.
func SecAccessControlGetTypeIDCallError() error {
	if _secAccessControlGetTypeID != nil {
		return nil
	}
	return symbolCallError("SecAccessControlGetTypeID", "10.10", _secAccessControlGetTypeIDErr)
}

// TrySecAccessControlGetTypeID calls SecAccessControlGetTypeID without panicking when the symbol is unavailable.
func TrySecAccessControlGetTypeID() (uint, error) {
	if err := SecAccessControlGetTypeIDCallError(); err != nil {
		return 0, err
	}
	return _secAccessControlGetTypeID(), nil
}

// SecAccessControlGetTypeID returns the unique identifier of the opaque type to which a keychain item access control object belongs.
//
// See: https://developer.apple.com/documentation/Security/SecAccessControlGetTypeID()
func SecAccessControlGetTypeID() uint {
	result, callErr := TrySecAccessControlGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secAddSharedWebCredential func(fqdn corefoundation.CFStringRef, account corefoundation.CFStringRef, password corefoundation.CFStringRef)
var _secAddSharedWebCredentialErr error

// CanCallSecAddSharedWebCredential reports whether the symbol for SecAddSharedWebCredential is available on this system.
func CanCallSecAddSharedWebCredential() bool {
	return _secAddSharedWebCredential != nil
}

// SecAddSharedWebCredentialCallError returns the symbol lookup or registration error for SecAddSharedWebCredential.
func SecAddSharedWebCredentialCallError() error {
	if _secAddSharedWebCredential != nil {
		return nil
	}
	return symbolCallError("SecAddSharedWebCredential", "11.0", _secAddSharedWebCredentialErr)
}

// TrySecAddSharedWebCredential calls SecAddSharedWebCredential without panicking when the symbol is unavailable.
func TrySecAddSharedWebCredential(fqdn corefoundation.CFStringRef, account corefoundation.CFStringRef, password corefoundation.CFStringRef) error {
	if err := SecAddSharedWebCredentialCallError(); err != nil {
		return err
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
	if callErr := TrySecAddSharedWebCredential(fqdn, account, password); callErr != nil {
		panic(callErr)
	}
}

var _secAsn1AllocCopy func(coder unsafe.Pointer, src unsafe.Pointer, len_ uintptr, dest unsafe.Pointer) int32
var _secAsn1AllocCopyErr error

// CanCallSecAsn1AllocCopy reports whether the symbol for SecAsn1AllocCopy is available on this system.
func CanCallSecAsn1AllocCopy() bool {
	return _secAsn1AllocCopy != nil
}

// SecAsn1AllocCopyCallError returns the symbol lookup or registration error for SecAsn1AllocCopy.
func SecAsn1AllocCopyCallError() error {
	if _secAsn1AllocCopy != nil {
		return nil
	}
	return symbolCallError("SecAsn1AllocCopy", "10.0", _secAsn1AllocCopyErr)
}

// TrySecAsn1AllocCopy calls SecAsn1AllocCopy without panicking when the symbol is unavailable.
func TrySecAsn1AllocCopy(coder unsafe.Pointer, src unsafe.Pointer, len_ uintptr, dest unsafe.Pointer) (int32, error) {
	if err := SecAsn1AllocCopyCallError(); err != nil {
		return 0, err
	}
	return _secAsn1AllocCopy(coder, src, len_, dest), nil
}

// SecAsn1AllocCopy allocates memory for an item’s data field in the coder object’s memory pool and copies in a block of data.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecAsn1AllocCopy
func SecAsn1AllocCopy(coder unsafe.Pointer, src unsafe.Pointer, len_ uintptr, dest unsafe.Pointer) int32 {
	result, callErr := TrySecAsn1AllocCopy(coder, src, len_, dest)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secAsn1AllocCopyItem func(coder unsafe.Pointer, src unsafe.Pointer, dest unsafe.Pointer) int32
var _secAsn1AllocCopyItemErr error

// CanCallSecAsn1AllocCopyItem reports whether the symbol for SecAsn1AllocCopyItem is available on this system.
func CanCallSecAsn1AllocCopyItem() bool {
	return _secAsn1AllocCopyItem != nil
}

// SecAsn1AllocCopyItemCallError returns the symbol lookup or registration error for SecAsn1AllocCopyItem.
func SecAsn1AllocCopyItemCallError() error {
	if _secAsn1AllocCopyItem != nil {
		return nil
	}
	return symbolCallError("SecAsn1AllocCopyItem", "10.0", _secAsn1AllocCopyItemErr)
}

// TrySecAsn1AllocCopyItem calls SecAsn1AllocCopyItem without panicking when the symbol is unavailable.
func TrySecAsn1AllocCopyItem(coder unsafe.Pointer, src unsafe.Pointer, dest unsafe.Pointer) (int32, error) {
	if err := SecAsn1AllocCopyItemCallError(); err != nil {
		return 0, err
	}
	return _secAsn1AllocCopyItem(coder, src, dest), nil
}

// SecAsn1AllocCopyItem allocates memory for an item’s data field in the coder object’s memory pool and copies in a block of data from another item.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecAsn1AllocCopyItem
func SecAsn1AllocCopyItem(coder unsafe.Pointer, src unsafe.Pointer, dest unsafe.Pointer) int32 {
	result, callErr := TrySecAsn1AllocCopyItem(coder, src, dest)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secAsn1AllocItem func(coder unsafe.Pointer, item unsafe.Pointer, len_ uintptr) int32
var _secAsn1AllocItemErr error

// CanCallSecAsn1AllocItem reports whether the symbol for SecAsn1AllocItem is available on this system.
func CanCallSecAsn1AllocItem() bool {
	return _secAsn1AllocItem != nil
}

// SecAsn1AllocItemCallError returns the symbol lookup or registration error for SecAsn1AllocItem.
func SecAsn1AllocItemCallError() error {
	if _secAsn1AllocItem != nil {
		return nil
	}
	return symbolCallError("SecAsn1AllocItem", "10.0", _secAsn1AllocItemErr)
}

// TrySecAsn1AllocItem calls SecAsn1AllocItem without panicking when the symbol is unavailable.
func TrySecAsn1AllocItem(coder unsafe.Pointer, item unsafe.Pointer, len_ uintptr) (int32, error) {
	if err := SecAsn1AllocItemCallError(); err != nil {
		return 0, err
	}
	return _secAsn1AllocItem(coder, item, len_), nil
}

// SecAsn1AllocItem allocates memory for an item’s data field in the coder object’s memory pool.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecAsn1AllocItem
func SecAsn1AllocItem(coder unsafe.Pointer, item unsafe.Pointer, len_ uintptr) int32 {
	result, callErr := TrySecAsn1AllocItem(coder, item, len_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secAsn1CoderCreate func(coder unsafe.Pointer) int32
var _secAsn1CoderCreateErr error

// CanCallSecAsn1CoderCreate reports whether the symbol for SecAsn1CoderCreate is available on this system.
func CanCallSecAsn1CoderCreate() bool {
	return _secAsn1CoderCreate != nil
}

// SecAsn1CoderCreateCallError returns the symbol lookup or registration error for SecAsn1CoderCreate.
func SecAsn1CoderCreateCallError() error {
	if _secAsn1CoderCreate != nil {
		return nil
	}
	return symbolCallError("SecAsn1CoderCreate", "10.0", _secAsn1CoderCreateErr)
}

// TrySecAsn1CoderCreate calls SecAsn1CoderCreate without panicking when the symbol is unavailable.
func TrySecAsn1CoderCreate(coder unsafe.Pointer) (int32, error) {
	if err := SecAsn1CoderCreateCallError(); err != nil {
		return 0, err
	}
	return _secAsn1CoderCreate(coder), nil
}

// SecAsn1CoderCreate creates an ASN.1 coder object.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecAsn1CoderCreate
func SecAsn1CoderCreate(coder unsafe.Pointer) int32 {
	result, callErr := TrySecAsn1CoderCreate(coder)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secAsn1CoderRelease func(coder unsafe.Pointer) int32
var _secAsn1CoderReleaseErr error

// CanCallSecAsn1CoderRelease reports whether the symbol for SecAsn1CoderRelease is available on this system.
func CanCallSecAsn1CoderRelease() bool {
	return _secAsn1CoderRelease != nil
}

// SecAsn1CoderReleaseCallError returns the symbol lookup or registration error for SecAsn1CoderRelease.
func SecAsn1CoderReleaseCallError() error {
	if _secAsn1CoderRelease != nil {
		return nil
	}
	return symbolCallError("SecAsn1CoderRelease", "10.0", _secAsn1CoderReleaseErr)
}

// TrySecAsn1CoderRelease calls SecAsn1CoderRelease without panicking when the symbol is unavailable.
func TrySecAsn1CoderRelease(coder unsafe.Pointer) (int32, error) {
	if err := SecAsn1CoderReleaseCallError(); err != nil {
		return 0, err
	}
	return _secAsn1CoderRelease(coder), nil
}

// SecAsn1CoderRelease destroys an ASN.1 coder object and releases all of its memory.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecAsn1CoderRelease
func SecAsn1CoderRelease(coder unsafe.Pointer) int32 {
	result, callErr := TrySecAsn1CoderRelease(coder)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secAsn1Decode func(coder unsafe.Pointer, src unsafe.Pointer, len_ uintptr, templates unsafe.Pointer, dest unsafe.Pointer) int32
var _secAsn1DecodeErr error

// CanCallSecAsn1Decode reports whether the symbol for SecAsn1Decode is available on this system.
func CanCallSecAsn1Decode() bool {
	return _secAsn1Decode != nil
}

// SecAsn1DecodeCallError returns the symbol lookup or registration error for SecAsn1Decode.
func SecAsn1DecodeCallError() error {
	if _secAsn1Decode != nil {
		return nil
	}
	return symbolCallError("SecAsn1Decode", "10.0", _secAsn1DecodeErr)
}

// TrySecAsn1Decode calls SecAsn1Decode without panicking when the symbol is unavailable.
func TrySecAsn1Decode(coder unsafe.Pointer, src unsafe.Pointer, len_ uintptr, templates unsafe.Pointer, dest unsafe.Pointer) (int32, error) {
	if err := SecAsn1DecodeCallError(); err != nil {
		return 0, err
	}
	return _secAsn1Decode(coder, src, len_, templates, dest), nil
}

// SecAsn1Decode decodes untyped DER data.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecAsn1Decode
func SecAsn1Decode(coder unsafe.Pointer, src unsafe.Pointer, len_ uintptr, templates unsafe.Pointer, dest unsafe.Pointer) int32 {
	result, callErr := TrySecAsn1Decode(coder, src, len_, templates, dest)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secAsn1DecodeData func(coder unsafe.Pointer, src unsafe.Pointer, templ unsafe.Pointer, dest unsafe.Pointer) int32
var _secAsn1DecodeDataErr error

// CanCallSecAsn1DecodeData reports whether the symbol for SecAsn1DecodeData is available on this system.
func CanCallSecAsn1DecodeData() bool {
	return _secAsn1DecodeData != nil
}

// SecAsn1DecodeDataCallError returns the symbol lookup or registration error for SecAsn1DecodeData.
func SecAsn1DecodeDataCallError() error {
	if _secAsn1DecodeData != nil {
		return nil
	}
	return symbolCallError("SecAsn1DecodeData", "10.0", _secAsn1DecodeDataErr)
}

// TrySecAsn1DecodeData calls SecAsn1DecodeData without panicking when the symbol is unavailable.
func TrySecAsn1DecodeData(coder unsafe.Pointer, src unsafe.Pointer, templ unsafe.Pointer, dest unsafe.Pointer) (int32, error) {
	if err := SecAsn1DecodeDataCallError(); err != nil {
		return 0, err
	}
	return _secAsn1DecodeData(coder, src, templ, dest), nil
}

// SecAsn1DecodeData decodes an ASN.1 item in DER format.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecAsn1DecodeData
func SecAsn1DecodeData(coder unsafe.Pointer, src unsafe.Pointer, templ unsafe.Pointer, dest unsafe.Pointer) int32 {
	result, callErr := TrySecAsn1DecodeData(coder, src, templ, dest)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secAsn1EncodeItem func(coder unsafe.Pointer, src unsafe.Pointer, templates unsafe.Pointer, dest unsafe.Pointer) int32
var _secAsn1EncodeItemErr error

// CanCallSecAsn1EncodeItem reports whether the symbol for SecAsn1EncodeItem is available on this system.
func CanCallSecAsn1EncodeItem() bool {
	return _secAsn1EncodeItem != nil
}

// SecAsn1EncodeItemCallError returns the symbol lookup or registration error for SecAsn1EncodeItem.
func SecAsn1EncodeItemCallError() error {
	if _secAsn1EncodeItem != nil {
		return nil
	}
	return symbolCallError("SecAsn1EncodeItem", "10.0", _secAsn1EncodeItemErr)
}

// TrySecAsn1EncodeItem calls SecAsn1EncodeItem without panicking when the symbol is unavailable.
func TrySecAsn1EncodeItem(coder unsafe.Pointer, src unsafe.Pointer, templates unsafe.Pointer, dest unsafe.Pointer) (int32, error) {
	if err := SecAsn1EncodeItemCallError(); err != nil {
		return 0, err
	}
	return _secAsn1EncodeItem(coder, src, templates, dest), nil
}

// SecAsn1EncodeItem encodes data in DER format.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecAsn1EncodeItem
func SecAsn1EncodeItem(coder unsafe.Pointer, src unsafe.Pointer, templates unsafe.Pointer, dest unsafe.Pointer) int32 {
	result, callErr := TrySecAsn1EncodeItem(coder, src, templates, dest)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secAsn1Malloc func(coder unsafe.Pointer, len_ uintptr) unsafe.Pointer
var _secAsn1MallocErr error

// CanCallSecAsn1Malloc reports whether the symbol for SecAsn1Malloc is available on this system.
func CanCallSecAsn1Malloc() bool {
	return _secAsn1Malloc != nil
}

// SecAsn1MallocCallError returns the symbol lookup or registration error for SecAsn1Malloc.
func SecAsn1MallocCallError() error {
	if _secAsn1Malloc != nil {
		return nil
	}
	return symbolCallError("SecAsn1Malloc", "10.0", _secAsn1MallocErr)
}

// TrySecAsn1Malloc calls SecAsn1Malloc without panicking when the symbol is unavailable.
func TrySecAsn1Malloc(coder unsafe.Pointer, len_ uintptr) (unsafe.Pointer, error) {
	if err := SecAsn1MallocCallError(); err != nil {
		return nil, err
	}
	return _secAsn1Malloc(coder, len_), nil
}

// SecAsn1Malloc allocates memory in the coder object’s memory pool.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecAsn1Malloc
func SecAsn1Malloc(coder unsafe.Pointer, len_ uintptr) unsafe.Pointer {
	result, callErr := TrySecAsn1Malloc(coder, len_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secAsn1OidCompare func(oid1 unsafe.Pointer, oid2 unsafe.Pointer) bool
var _secAsn1OidCompareErr error

// CanCallSecAsn1OidCompare reports whether the symbol for SecAsn1OidCompare is available on this system.
func CanCallSecAsn1OidCompare() bool {
	return _secAsn1OidCompare != nil
}

// SecAsn1OidCompareCallError returns the symbol lookup or registration error for SecAsn1OidCompare.
func SecAsn1OidCompareCallError() error {
	if _secAsn1OidCompare != nil {
		return nil
	}
	return symbolCallError("SecAsn1OidCompare", "10.0", _secAsn1OidCompareErr)
}

// TrySecAsn1OidCompare calls SecAsn1OidCompare without panicking when the symbol is unavailable.
func TrySecAsn1OidCompare(oid1 unsafe.Pointer, oid2 unsafe.Pointer) (bool, error) {
	if err := SecAsn1OidCompareCallError(); err != nil {
		return false, err
	}
	return _secAsn1OidCompare(oid1, oid2), nil
}

// SecAsn1OidCompare compares two decoded object identifiers.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecAsn1OidCompare
func SecAsn1OidCompare(oid1 unsafe.Pointer, oid2 unsafe.Pointer) bool {
	result, callErr := TrySecAsn1OidCompare(oid1, oid2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateAddToKeychain func(certificate SecCertificateRef, keychain SecKeychainRef) int32
var _secCertificateAddToKeychainErr error

// CanCallSecCertificateAddToKeychain reports whether the symbol for SecCertificateAddToKeychain is available on this system.
func CanCallSecCertificateAddToKeychain() bool {
	return _secCertificateAddToKeychain != nil
}

// SecCertificateAddToKeychainCallError returns the symbol lookup or registration error for SecCertificateAddToKeychain.
func SecCertificateAddToKeychainCallError() error {
	if _secCertificateAddToKeychain != nil {
		return nil
	}
	return symbolCallError("SecCertificateAddToKeychain", "10.3", _secCertificateAddToKeychainErr)
}

// TrySecCertificateAddToKeychain calls SecCertificateAddToKeychain without panicking when the symbol is unavailable.
func TrySecCertificateAddToKeychain(certificate SecCertificateRef, keychain SecKeychainRef) (int32, error) {
	if err := SecCertificateAddToKeychainCallError(); err != nil {
		return 0, err
	}
	return _secCertificateAddToKeychain(certificate, keychain), nil
}

// SecCertificateAddToKeychain adds a certificate to a keychain.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateAddToKeychain(_:_:)
func SecCertificateAddToKeychain(certificate SecCertificateRef, keychain SecKeychainRef) int32 {
	result, callErr := TrySecCertificateAddToKeychain(certificate, keychain)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateCopyCommonName func(certificate SecCertificateRef, commonName *corefoundation.CFStringRef) int32
var _secCertificateCopyCommonNameErr error

// CanCallSecCertificateCopyCommonName reports whether the symbol for SecCertificateCopyCommonName is available on this system.
func CanCallSecCertificateCopyCommonName() bool {
	return _secCertificateCopyCommonName != nil
}

// SecCertificateCopyCommonNameCallError returns the symbol lookup or registration error for SecCertificateCopyCommonName.
func SecCertificateCopyCommonNameCallError() error {
	if _secCertificateCopyCommonName != nil {
		return nil
	}
	return symbolCallError("SecCertificateCopyCommonName", "10.5", _secCertificateCopyCommonNameErr)
}

// TrySecCertificateCopyCommonName calls SecCertificateCopyCommonName without panicking when the symbol is unavailable.
func TrySecCertificateCopyCommonName(certificate SecCertificateRef, commonName *corefoundation.CFStringRef) (int32, error) {
	if err := SecCertificateCopyCommonNameCallError(); err != nil {
		return 0, err
	}
	return _secCertificateCopyCommonName(certificate, commonName), nil
}

// SecCertificateCopyCommonName retrieves the common name of the subject of a certificate.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopyCommonName(_:_:)
func SecCertificateCopyCommonName(certificate SecCertificateRef, commonName *corefoundation.CFStringRef) int32 {
	result, callErr := TrySecCertificateCopyCommonName(certificate, commonName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateCopyData func(certificate SecCertificateRef) corefoundation.CFDataRef
var _secCertificateCopyDataErr error

// CanCallSecCertificateCopyData reports whether the symbol for SecCertificateCopyData is available on this system.
func CanCallSecCertificateCopyData() bool {
	return _secCertificateCopyData != nil
}

// SecCertificateCopyDataCallError returns the symbol lookup or registration error for SecCertificateCopyData.
func SecCertificateCopyDataCallError() error {
	if _secCertificateCopyData != nil {
		return nil
	}
	return symbolCallError("SecCertificateCopyData", "10.6", _secCertificateCopyDataErr)
}

// TrySecCertificateCopyData calls SecCertificateCopyData without panicking when the symbol is unavailable.
func TrySecCertificateCopyData(certificate SecCertificateRef) (corefoundation.CFDataRef, error) {
	if err := SecCertificateCopyDataCallError(); err != nil {
		return 0, err
	}
	return _secCertificateCopyData(certificate), nil
}

// SecCertificateCopyData returns a DER representation of a certificate given a certificate object.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopyData(_:)
func SecCertificateCopyData(certificate SecCertificateRef) corefoundation.CFDataRef {
	result, callErr := TrySecCertificateCopyData(certificate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateCopyEmailAddresses func(certificate SecCertificateRef, emailAddresses *corefoundation.CFArrayRef) int32
var _secCertificateCopyEmailAddressesErr error

// CanCallSecCertificateCopyEmailAddresses reports whether the symbol for SecCertificateCopyEmailAddresses is available on this system.
func CanCallSecCertificateCopyEmailAddresses() bool {
	return _secCertificateCopyEmailAddresses != nil
}

// SecCertificateCopyEmailAddressesCallError returns the symbol lookup or registration error for SecCertificateCopyEmailAddresses.
func SecCertificateCopyEmailAddressesCallError() error {
	if _secCertificateCopyEmailAddresses != nil {
		return nil
	}
	return symbolCallError("SecCertificateCopyEmailAddresses", "10.5", _secCertificateCopyEmailAddressesErr)
}

// TrySecCertificateCopyEmailAddresses calls SecCertificateCopyEmailAddresses without panicking when the symbol is unavailable.
func TrySecCertificateCopyEmailAddresses(certificate SecCertificateRef, emailAddresses *corefoundation.CFArrayRef) (int32, error) {
	if err := SecCertificateCopyEmailAddressesCallError(); err != nil {
		return 0, err
	}
	return _secCertificateCopyEmailAddresses(certificate, emailAddresses), nil
}

// SecCertificateCopyEmailAddresses retrieves the email addresses for the subject of a certificate.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopyEmailAddresses(_:_:)
func SecCertificateCopyEmailAddresses(certificate SecCertificateRef, emailAddresses *corefoundation.CFArrayRef) int32 {
	result, callErr := TrySecCertificateCopyEmailAddresses(certificate, emailAddresses)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateCopyKey func(certificate SecCertificateRef) SecKeyRef
var _secCertificateCopyKeyErr error

// CanCallSecCertificateCopyKey reports whether the symbol for SecCertificateCopyKey is available on this system.
func CanCallSecCertificateCopyKey() bool {
	return _secCertificateCopyKey != nil
}

// SecCertificateCopyKeyCallError returns the symbol lookup or registration error for SecCertificateCopyKey.
func SecCertificateCopyKeyCallError() error {
	if _secCertificateCopyKey != nil {
		return nil
	}
	return symbolCallError("SecCertificateCopyKey", "10.14", _secCertificateCopyKeyErr)
}

// TrySecCertificateCopyKey calls SecCertificateCopyKey without panicking when the symbol is unavailable.
func TrySecCertificateCopyKey(certificate SecCertificateRef) (SecKeyRef, error) {
	if err := SecCertificateCopyKeyCallError(); err != nil {
		return 0, err
	}
	return _secCertificateCopyKey(certificate), nil
}

// SecCertificateCopyKey retrieves the public key for a given certificate.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopyKey(_:)
func SecCertificateCopyKey(certificate SecCertificateRef) SecKeyRef {
	result, callErr := TrySecCertificateCopyKey(certificate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateCopyLongDescription func(alloc corefoundation.CFAllocatorRef, certificate SecCertificateRef, err *corefoundation.CFErrorRef) corefoundation.CFStringRef
var _secCertificateCopyLongDescriptionErr error

// CanCallSecCertificateCopyLongDescription reports whether the symbol for SecCertificateCopyLongDescription is available on this system.
func CanCallSecCertificateCopyLongDescription() bool {
	return _secCertificateCopyLongDescription != nil
}

// SecCertificateCopyLongDescriptionCallError returns the symbol lookup or registration error for SecCertificateCopyLongDescription.
func SecCertificateCopyLongDescriptionCallError() error {
	if _secCertificateCopyLongDescription != nil {
		return nil
	}
	return symbolCallError("SecCertificateCopyLongDescription", "10.7", _secCertificateCopyLongDescriptionErr)
}

// TrySecCertificateCopyLongDescription calls SecCertificateCopyLongDescription without panicking when the symbol is unavailable.
func TrySecCertificateCopyLongDescription(alloc corefoundation.CFAllocatorRef, certificate SecCertificateRef, err *corefoundation.CFErrorRef) (corefoundation.CFStringRef, error) {
	if err := SecCertificateCopyLongDescriptionCallError(); err != nil {
		return 0, err
	}
	return _secCertificateCopyLongDescription(alloc, certificate, err), nil
}

// SecCertificateCopyLongDescription returns a copy of the long description of a certificate.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopyLongDescription(_:_:_:)
func SecCertificateCopyLongDescription(alloc corefoundation.CFAllocatorRef, certificate SecCertificateRef, err *corefoundation.CFErrorRef) corefoundation.CFStringRef {
	result, callErr := TrySecCertificateCopyLongDescription(alloc, certificate, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateCopyNormalizedIssuerSequence func(certificate SecCertificateRef) corefoundation.CFDataRef
var _secCertificateCopyNormalizedIssuerSequenceErr error

// CanCallSecCertificateCopyNormalizedIssuerSequence reports whether the symbol for SecCertificateCopyNormalizedIssuerSequence is available on this system.
func CanCallSecCertificateCopyNormalizedIssuerSequence() bool {
	return _secCertificateCopyNormalizedIssuerSequence != nil
}

// SecCertificateCopyNormalizedIssuerSequenceCallError returns the symbol lookup or registration error for SecCertificateCopyNormalizedIssuerSequence.
func SecCertificateCopyNormalizedIssuerSequenceCallError() error {
	if _secCertificateCopyNormalizedIssuerSequence != nil {
		return nil
	}
	return symbolCallError("SecCertificateCopyNormalizedIssuerSequence", "10.12.4", _secCertificateCopyNormalizedIssuerSequenceErr)
}

// TrySecCertificateCopyNormalizedIssuerSequence calls SecCertificateCopyNormalizedIssuerSequence without panicking when the symbol is unavailable.
func TrySecCertificateCopyNormalizedIssuerSequence(certificate SecCertificateRef) (corefoundation.CFDataRef, error) {
	if err := SecCertificateCopyNormalizedIssuerSequenceCallError(); err != nil {
		return 0, err
	}
	return _secCertificateCopyNormalizedIssuerSequence(certificate), nil
}

// SecCertificateCopyNormalizedIssuerSequence retrieves the normalized issuer sequence from a certificate.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopyNormalizedIssuerSequence(_:)
func SecCertificateCopyNormalizedIssuerSequence(certificate SecCertificateRef) corefoundation.CFDataRef {
	result, callErr := TrySecCertificateCopyNormalizedIssuerSequence(certificate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateCopyNormalizedSubjectSequence func(certificate SecCertificateRef) corefoundation.CFDataRef
var _secCertificateCopyNormalizedSubjectSequenceErr error

// CanCallSecCertificateCopyNormalizedSubjectSequence reports whether the symbol for SecCertificateCopyNormalizedSubjectSequence is available on this system.
func CanCallSecCertificateCopyNormalizedSubjectSequence() bool {
	return _secCertificateCopyNormalizedSubjectSequence != nil
}

// SecCertificateCopyNormalizedSubjectSequenceCallError returns the symbol lookup or registration error for SecCertificateCopyNormalizedSubjectSequence.
func SecCertificateCopyNormalizedSubjectSequenceCallError() error {
	if _secCertificateCopyNormalizedSubjectSequence != nil {
		return nil
	}
	return symbolCallError("SecCertificateCopyNormalizedSubjectSequence", "10.12.4", _secCertificateCopyNormalizedSubjectSequenceErr)
}

// TrySecCertificateCopyNormalizedSubjectSequence calls SecCertificateCopyNormalizedSubjectSequence without panicking when the symbol is unavailable.
func TrySecCertificateCopyNormalizedSubjectSequence(certificate SecCertificateRef) (corefoundation.CFDataRef, error) {
	if err := SecCertificateCopyNormalizedSubjectSequenceCallError(); err != nil {
		return 0, err
	}
	return _secCertificateCopyNormalizedSubjectSequence(certificate), nil
}

// SecCertificateCopyNormalizedSubjectSequence retrieves the normalized subject sequence from a certificate.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopyNormalizedSubjectSequence(_:)
func SecCertificateCopyNormalizedSubjectSequence(certificate SecCertificateRef) corefoundation.CFDataRef {
	result, callErr := TrySecCertificateCopyNormalizedSubjectSequence(certificate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateCopyNotValidAfterDate func(certificate SecCertificateRef) corefoundation.CFDateRef
var _secCertificateCopyNotValidAfterDateErr error

// CanCallSecCertificateCopyNotValidAfterDate reports whether the symbol for SecCertificateCopyNotValidAfterDate is available on this system.
func CanCallSecCertificateCopyNotValidAfterDate() bool {
	return _secCertificateCopyNotValidAfterDate != nil
}

// SecCertificateCopyNotValidAfterDateCallError returns the symbol lookup or registration error for SecCertificateCopyNotValidAfterDate.
func SecCertificateCopyNotValidAfterDateCallError() error {
	if _secCertificateCopyNotValidAfterDate != nil {
		return nil
	}
	return symbolCallError("SecCertificateCopyNotValidAfterDate", "15.0", _secCertificateCopyNotValidAfterDateErr)
}

// TrySecCertificateCopyNotValidAfterDate calls SecCertificateCopyNotValidAfterDate without panicking when the symbol is unavailable.
func TrySecCertificateCopyNotValidAfterDate(certificate SecCertificateRef) (corefoundation.CFDateRef, error) {
	if err := SecCertificateCopyNotValidAfterDateCallError(); err != nil {
		return 0, err
	}
	return _secCertificateCopyNotValidAfterDate(certificate), nil
}

// SecCertificateCopyNotValidAfterDate.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopyNotValidAfterDate(_:)
func SecCertificateCopyNotValidAfterDate(certificate SecCertificateRef) corefoundation.CFDateRef {
	result, callErr := TrySecCertificateCopyNotValidAfterDate(certificate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateCopyNotValidBeforeDate func(certificate SecCertificateRef) corefoundation.CFDateRef
var _secCertificateCopyNotValidBeforeDateErr error

// CanCallSecCertificateCopyNotValidBeforeDate reports whether the symbol for SecCertificateCopyNotValidBeforeDate is available on this system.
func CanCallSecCertificateCopyNotValidBeforeDate() bool {
	return _secCertificateCopyNotValidBeforeDate != nil
}

// SecCertificateCopyNotValidBeforeDateCallError returns the symbol lookup or registration error for SecCertificateCopyNotValidBeforeDate.
func SecCertificateCopyNotValidBeforeDateCallError() error {
	if _secCertificateCopyNotValidBeforeDate != nil {
		return nil
	}
	return symbolCallError("SecCertificateCopyNotValidBeforeDate", "15.0", _secCertificateCopyNotValidBeforeDateErr)
}

// TrySecCertificateCopyNotValidBeforeDate calls SecCertificateCopyNotValidBeforeDate without panicking when the symbol is unavailable.
func TrySecCertificateCopyNotValidBeforeDate(certificate SecCertificateRef) (corefoundation.CFDateRef, error) {
	if err := SecCertificateCopyNotValidBeforeDateCallError(); err != nil {
		return 0, err
	}
	return _secCertificateCopyNotValidBeforeDate(certificate), nil
}

// SecCertificateCopyNotValidBeforeDate.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopyNotValidBeforeDate(_:)
func SecCertificateCopyNotValidBeforeDate(certificate SecCertificateRef) corefoundation.CFDateRef {
	result, callErr := TrySecCertificateCopyNotValidBeforeDate(certificate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateCopyPreference func(name corefoundation.CFStringRef, keyUsage uint32, certificate *SecCertificateRef) int32
var _secCertificateCopyPreferenceErr error

// CanCallSecCertificateCopyPreference reports whether the symbol for SecCertificateCopyPreference is available on this system.
func CanCallSecCertificateCopyPreference() bool {
	return _secCertificateCopyPreference != nil
}

// SecCertificateCopyPreferenceCallError returns the symbol lookup or registration error for SecCertificateCopyPreference.
func SecCertificateCopyPreferenceCallError() error {
	if _secCertificateCopyPreference != nil {
		return nil
	}
	return symbolCallError("SecCertificateCopyPreference", "10.0", _secCertificateCopyPreferenceErr)
}

// TrySecCertificateCopyPreference calls SecCertificateCopyPreference without panicking when the symbol is unavailable.
func TrySecCertificateCopyPreference(name corefoundation.CFStringRef, keyUsage uint32, certificate *SecCertificateRef) (int32, error) {
	if err := SecCertificateCopyPreferenceCallError(); err != nil {
		return 0, err
	}
	return _secCertificateCopyPreference(name, keyUsage, certificate), nil
}

// SecCertificateCopyPreference retrieves the preferred certificate for the specified name and key use.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopyPreference
func SecCertificateCopyPreference(name corefoundation.CFStringRef, keyUsage uint32, certificate *SecCertificateRef) int32 {
	result, callErr := TrySecCertificateCopyPreference(name, keyUsage, certificate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateCopyPreferred func(name corefoundation.CFStringRef, keyUsage corefoundation.CFArrayRef) SecCertificateRef
var _secCertificateCopyPreferredErr error

// CanCallSecCertificateCopyPreferred reports whether the symbol for SecCertificateCopyPreferred is available on this system.
func CanCallSecCertificateCopyPreferred() bool {
	return _secCertificateCopyPreferred != nil
}

// SecCertificateCopyPreferredCallError returns the symbol lookup or registration error for SecCertificateCopyPreferred.
func SecCertificateCopyPreferredCallError() error {
	if _secCertificateCopyPreferred != nil {
		return nil
	}
	return symbolCallError("SecCertificateCopyPreferred", "10.7", _secCertificateCopyPreferredErr)
}

// TrySecCertificateCopyPreferred calls SecCertificateCopyPreferred without panicking when the symbol is unavailable.
func TrySecCertificateCopyPreferred(name corefoundation.CFStringRef, keyUsage corefoundation.CFArrayRef) (SecCertificateRef, error) {
	if err := SecCertificateCopyPreferredCallError(); err != nil {
		return 0, err
	}
	return _secCertificateCopyPreferred(name, keyUsage), nil
}

// SecCertificateCopyPreferred returns the preferred certificate for the specified name and key usage.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopyPreferred(_:_:)
func SecCertificateCopyPreferred(name corefoundation.CFStringRef, keyUsage corefoundation.CFArrayRef) SecCertificateRef {
	result, callErr := TrySecCertificateCopyPreferred(name, keyUsage)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateCopySerialNumberData func(certificate SecCertificateRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef
var _secCertificateCopySerialNumberDataErr error

// CanCallSecCertificateCopySerialNumberData reports whether the symbol for SecCertificateCopySerialNumberData is available on this system.
func CanCallSecCertificateCopySerialNumberData() bool {
	return _secCertificateCopySerialNumberData != nil
}

// SecCertificateCopySerialNumberDataCallError returns the symbol lookup or registration error for SecCertificateCopySerialNumberData.
func SecCertificateCopySerialNumberDataCallError() error {
	if _secCertificateCopySerialNumberData != nil {
		return nil
	}
	return symbolCallError("SecCertificateCopySerialNumberData", "10.13", _secCertificateCopySerialNumberDataErr)
}

// TrySecCertificateCopySerialNumberData calls SecCertificateCopySerialNumberData without panicking when the symbol is unavailable.
func TrySecCertificateCopySerialNumberData(certificate SecCertificateRef, err *corefoundation.CFErrorRef) (corefoundation.CFDataRef, error) {
	if err := SecCertificateCopySerialNumberDataCallError(); err != nil {
		return 0, err
	}
	return _secCertificateCopySerialNumberData(certificate, err), nil
}

// SecCertificateCopySerialNumberData returns the certificate’s serial number.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopySerialNumberData(_:_:)
func SecCertificateCopySerialNumberData(certificate SecCertificateRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef {
	result, callErr := TrySecCertificateCopySerialNumberData(certificate, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateCopyShortDescription func(alloc corefoundation.CFAllocatorRef, certificate SecCertificateRef, err *corefoundation.CFErrorRef) corefoundation.CFStringRef
var _secCertificateCopyShortDescriptionErr error

// CanCallSecCertificateCopyShortDescription reports whether the symbol for SecCertificateCopyShortDescription is available on this system.
func CanCallSecCertificateCopyShortDescription() bool {
	return _secCertificateCopyShortDescription != nil
}

// SecCertificateCopyShortDescriptionCallError returns the symbol lookup or registration error for SecCertificateCopyShortDescription.
func SecCertificateCopyShortDescriptionCallError() error {
	if _secCertificateCopyShortDescription != nil {
		return nil
	}
	return symbolCallError("SecCertificateCopyShortDescription", "10.7", _secCertificateCopyShortDescriptionErr)
}

// TrySecCertificateCopyShortDescription calls SecCertificateCopyShortDescription without panicking when the symbol is unavailable.
func TrySecCertificateCopyShortDescription(alloc corefoundation.CFAllocatorRef, certificate SecCertificateRef, err *corefoundation.CFErrorRef) (corefoundation.CFStringRef, error) {
	if err := SecCertificateCopyShortDescriptionCallError(); err != nil {
		return 0, err
	}
	return _secCertificateCopyShortDescription(alloc, certificate, err), nil
}

// SecCertificateCopyShortDescription returns a copy of the short description of a certificate.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopyShortDescription(_:_:_:)
func SecCertificateCopyShortDescription(alloc corefoundation.CFAllocatorRef, certificate SecCertificateRef, err *corefoundation.CFErrorRef) corefoundation.CFStringRef {
	result, callErr := TrySecCertificateCopyShortDescription(alloc, certificate, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateCopySubjectSummary func(certificate SecCertificateRef) corefoundation.CFStringRef
var _secCertificateCopySubjectSummaryErr error

// CanCallSecCertificateCopySubjectSummary reports whether the symbol for SecCertificateCopySubjectSummary is available on this system.
func CanCallSecCertificateCopySubjectSummary() bool {
	return _secCertificateCopySubjectSummary != nil
}

// SecCertificateCopySubjectSummaryCallError returns the symbol lookup or registration error for SecCertificateCopySubjectSummary.
func SecCertificateCopySubjectSummaryCallError() error {
	if _secCertificateCopySubjectSummary != nil {
		return nil
	}
	return symbolCallError("SecCertificateCopySubjectSummary", "10.6", _secCertificateCopySubjectSummaryErr)
}

// TrySecCertificateCopySubjectSummary calls SecCertificateCopySubjectSummary without panicking when the symbol is unavailable.
func TrySecCertificateCopySubjectSummary(certificate SecCertificateRef) (corefoundation.CFStringRef, error) {
	if err := SecCertificateCopySubjectSummaryCallError(); err != nil {
		return 0, err
	}
	return _secCertificateCopySubjectSummary(certificate), nil
}

// SecCertificateCopySubjectSummary returns a human-readable summary of a certificate.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopySubjectSummary(_:)
func SecCertificateCopySubjectSummary(certificate SecCertificateRef) corefoundation.CFStringRef {
	result, callErr := TrySecCertificateCopySubjectSummary(certificate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateCopyValues func(certificate SecCertificateRef, keys corefoundation.CFArrayRef, err *corefoundation.CFErrorRef) corefoundation.CFDictionaryRef
var _secCertificateCopyValuesErr error

// CanCallSecCertificateCopyValues reports whether the symbol for SecCertificateCopyValues is available on this system.
func CanCallSecCertificateCopyValues() bool {
	return _secCertificateCopyValues != nil
}

// SecCertificateCopyValuesCallError returns the symbol lookup or registration error for SecCertificateCopyValues.
func SecCertificateCopyValuesCallError() error {
	if _secCertificateCopyValues != nil {
		return nil
	}
	return symbolCallError("SecCertificateCopyValues", "10.7", _secCertificateCopyValuesErr)
}

// TrySecCertificateCopyValues calls SecCertificateCopyValues without panicking when the symbol is unavailable.
func TrySecCertificateCopyValues(certificate SecCertificateRef, keys corefoundation.CFArrayRef, err *corefoundation.CFErrorRef) (corefoundation.CFDictionaryRef, error) {
	if err := SecCertificateCopyValuesCallError(); err != nil {
		return 0, err
	}
	return _secCertificateCopyValues(certificate, keys, err), nil
}

// SecCertificateCopyValues creates a dictionary that represents a certificate’s contents.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopyValues(_:_:_:)
func SecCertificateCopyValues(certificate SecCertificateRef, keys corefoundation.CFArrayRef, err *corefoundation.CFErrorRef) corefoundation.CFDictionaryRef {
	result, callErr := TrySecCertificateCopyValues(certificate, keys, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateCreateWithData func(allocator corefoundation.CFAllocatorRef, data corefoundation.CFDataRef) SecCertificateRef
var _secCertificateCreateWithDataErr error

// CanCallSecCertificateCreateWithData reports whether the symbol for SecCertificateCreateWithData is available on this system.
func CanCallSecCertificateCreateWithData() bool {
	return _secCertificateCreateWithData != nil
}

// SecCertificateCreateWithDataCallError returns the symbol lookup or registration error for SecCertificateCreateWithData.
func SecCertificateCreateWithDataCallError() error {
	if _secCertificateCreateWithData != nil {
		return nil
	}
	return symbolCallError("SecCertificateCreateWithData", "10.6", _secCertificateCreateWithDataErr)
}

// TrySecCertificateCreateWithData calls SecCertificateCreateWithData without panicking when the symbol is unavailable.
func TrySecCertificateCreateWithData(allocator corefoundation.CFAllocatorRef, data corefoundation.CFDataRef) (SecCertificateRef, error) {
	if err := SecCertificateCreateWithDataCallError(); err != nil {
		return 0, err
	}
	return _secCertificateCreateWithData(allocator, data), nil
}

// SecCertificateCreateWithData creates a certificate object from a DER representation of a certificate.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCreateWithData(_:_:)
func SecCertificateCreateWithData(allocator corefoundation.CFAllocatorRef, data corefoundation.CFDataRef) SecCertificateRef {
	result, callErr := TrySecCertificateCreateWithData(allocator, data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateGetAlgorithmID func(certificate SecCertificateRef, algid unsafe.Pointer) int32
var _secCertificateGetAlgorithmIDErr error

// CanCallSecCertificateGetAlgorithmID reports whether the symbol for SecCertificateGetAlgorithmID is available on this system.
func CanCallSecCertificateGetAlgorithmID() bool {
	return _secCertificateGetAlgorithmID != nil
}

// SecCertificateGetAlgorithmIDCallError returns the symbol lookup or registration error for SecCertificateGetAlgorithmID.
func SecCertificateGetAlgorithmIDCallError() error {
	if _secCertificateGetAlgorithmID != nil {
		return nil
	}
	return symbolCallError("SecCertificateGetAlgorithmID", "10.0", _secCertificateGetAlgorithmIDErr)
}

// TrySecCertificateGetAlgorithmID calls SecCertificateGetAlgorithmID without panicking when the symbol is unavailable.
func TrySecCertificateGetAlgorithmID(certificate SecCertificateRef, algid unsafe.Pointer) (int32, error) {
	if err := SecCertificateGetAlgorithmIDCallError(); err != nil {
		return 0, err
	}
	return _secCertificateGetAlgorithmID(certificate, algid), nil
}

// SecCertificateGetAlgorithmID retrieves the algorithm identifier for a certificate.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateGetAlgorithmID
func SecCertificateGetAlgorithmID(certificate SecCertificateRef, algid unsafe.Pointer) int32 {
	result, callErr := TrySecCertificateGetAlgorithmID(certificate, algid)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateGetCLHandle func(certificate SecCertificateRef, clHandle unsafe.Pointer) int32
var _secCertificateGetCLHandleErr error

// CanCallSecCertificateGetCLHandle reports whether the symbol for SecCertificateGetCLHandle is available on this system.
func CanCallSecCertificateGetCLHandle() bool {
	return _secCertificateGetCLHandle != nil
}

// SecCertificateGetCLHandleCallError returns the symbol lookup or registration error for SecCertificateGetCLHandle.
func SecCertificateGetCLHandleCallError() error {
	if _secCertificateGetCLHandle != nil {
		return nil
	}
	return symbolCallError("SecCertificateGetCLHandle", "10.0", _secCertificateGetCLHandleErr)
}

// TrySecCertificateGetCLHandle calls SecCertificateGetCLHandle without panicking when the symbol is unavailable.
func TrySecCertificateGetCLHandle(certificate SecCertificateRef, clHandle unsafe.Pointer) (int32, error) {
	if err := SecCertificateGetCLHandleCallError(); err != nil {
		return 0, err
	}
	return _secCertificateGetCLHandle(certificate, clHandle), nil
}

// SecCertificateGetCLHandle retrieves the certificate library handle from a certificate object.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateGetCLHandle
func SecCertificateGetCLHandle(certificate SecCertificateRef, clHandle unsafe.Pointer) int32 {
	result, callErr := TrySecCertificateGetCLHandle(certificate, clHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateGetData func(certificate SecCertificateRef, data unsafe.Pointer) int32
var _secCertificateGetDataErr error

// CanCallSecCertificateGetData reports whether the symbol for SecCertificateGetData is available on this system.
func CanCallSecCertificateGetData() bool {
	return _secCertificateGetData != nil
}

// SecCertificateGetDataCallError returns the symbol lookup or registration error for SecCertificateGetData.
func SecCertificateGetDataCallError() error {
	if _secCertificateGetData != nil {
		return nil
	}
	return symbolCallError("SecCertificateGetData", "10.0", _secCertificateGetDataErr)
}

// TrySecCertificateGetData calls SecCertificateGetData without panicking when the symbol is unavailable.
func TrySecCertificateGetData(certificate SecCertificateRef, data unsafe.Pointer) (int32, error) {
	if err := SecCertificateGetDataCallError(); err != nil {
		return 0, err
	}
	return _secCertificateGetData(certificate, data), nil
}

// SecCertificateGetData retrieves the data for a certificate.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateGetData
func SecCertificateGetData(certificate SecCertificateRef, data unsafe.Pointer) int32 {
	result, callErr := TrySecCertificateGetData(certificate, data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateGetIssuer func(certificate SecCertificateRef, issuer unsafe.Pointer) int32
var _secCertificateGetIssuerErr error

// CanCallSecCertificateGetIssuer reports whether the symbol for SecCertificateGetIssuer is available on this system.
func CanCallSecCertificateGetIssuer() bool {
	return _secCertificateGetIssuer != nil
}

// SecCertificateGetIssuerCallError returns the symbol lookup or registration error for SecCertificateGetIssuer.
func SecCertificateGetIssuerCallError() error {
	if _secCertificateGetIssuer != nil {
		return nil
	}
	return symbolCallError("SecCertificateGetIssuer", "10.0", _secCertificateGetIssuerErr)
}

// TrySecCertificateGetIssuer calls SecCertificateGetIssuer without panicking when the symbol is unavailable.
func TrySecCertificateGetIssuer(certificate SecCertificateRef, issuer unsafe.Pointer) (int32, error) {
	if err := SecCertificateGetIssuerCallError(); err != nil {
		return 0, err
	}
	return _secCertificateGetIssuer(certificate, issuer), nil
}

// SecCertificateGetIssuer unsupported.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateGetIssuer
func SecCertificateGetIssuer(certificate SecCertificateRef, issuer unsafe.Pointer) int32 {
	result, callErr := TrySecCertificateGetIssuer(certificate, issuer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateGetSubject func(certificate SecCertificateRef, subject unsafe.Pointer) int32
var _secCertificateGetSubjectErr error

// CanCallSecCertificateGetSubject reports whether the symbol for SecCertificateGetSubject is available on this system.
func CanCallSecCertificateGetSubject() bool {
	return _secCertificateGetSubject != nil
}

// SecCertificateGetSubjectCallError returns the symbol lookup or registration error for SecCertificateGetSubject.
func SecCertificateGetSubjectCallError() error {
	if _secCertificateGetSubject != nil {
		return nil
	}
	return symbolCallError("SecCertificateGetSubject", "10.0", _secCertificateGetSubjectErr)
}

// TrySecCertificateGetSubject calls SecCertificateGetSubject without panicking when the symbol is unavailable.
func TrySecCertificateGetSubject(certificate SecCertificateRef, subject unsafe.Pointer) (int32, error) {
	if err := SecCertificateGetSubjectCallError(); err != nil {
		return 0, err
	}
	return _secCertificateGetSubject(certificate, subject), nil
}

// SecCertificateGetSubject unsupported.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateGetSubject
func SecCertificateGetSubject(certificate SecCertificateRef, subject unsafe.Pointer) int32 {
	result, callErr := TrySecCertificateGetSubject(certificate, subject)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateGetType func(certificate SecCertificateRef, certificateType unsafe.Pointer) int32
var _secCertificateGetTypeErr error

// CanCallSecCertificateGetType reports whether the symbol for SecCertificateGetType is available on this system.
func CanCallSecCertificateGetType() bool {
	return _secCertificateGetType != nil
}

// SecCertificateGetTypeCallError returns the symbol lookup or registration error for SecCertificateGetType.
func SecCertificateGetTypeCallError() error {
	if _secCertificateGetType != nil {
		return nil
	}
	return symbolCallError("SecCertificateGetType", "10.0", _secCertificateGetTypeErr)
}

// TrySecCertificateGetType calls SecCertificateGetType without panicking when the symbol is unavailable.
func TrySecCertificateGetType(certificate SecCertificateRef, certificateType unsafe.Pointer) (int32, error) {
	if err := SecCertificateGetTypeCallError(); err != nil {
		return 0, err
	}
	return _secCertificateGetType(certificate, certificateType), nil
}

// SecCertificateGetType retrieves the type of a specified certificate.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateGetType
func SecCertificateGetType(certificate SecCertificateRef, certificateType unsafe.Pointer) int32 {
	result, callErr := TrySecCertificateGetType(certificate, certificateType)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateGetTypeID func() uint
var _secCertificateGetTypeIDErr error

// CanCallSecCertificateGetTypeID reports whether the symbol for SecCertificateGetTypeID is available on this system.
func CanCallSecCertificateGetTypeID() bool {
	return _secCertificateGetTypeID != nil
}

// SecCertificateGetTypeIDCallError returns the symbol lookup or registration error for SecCertificateGetTypeID.
func SecCertificateGetTypeIDCallError() error {
	if _secCertificateGetTypeID != nil {
		return nil
	}
	return symbolCallError("SecCertificateGetTypeID", "10.3", _secCertificateGetTypeIDErr)
}

// TrySecCertificateGetTypeID calls SecCertificateGetTypeID without panicking when the symbol is unavailable.
func TrySecCertificateGetTypeID() (uint, error) {
	if err := SecCertificateGetTypeIDCallError(); err != nil {
		return 0, err
	}
	return _secCertificateGetTypeID(), nil
}

// SecCertificateGetTypeID returns the unique identifier of the opaque type to which a certificate object belongs.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateGetTypeID()
func SecCertificateGetTypeID() uint {
	result, callErr := TrySecCertificateGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateSetPreference func(certificate SecCertificateRef, name corefoundation.CFStringRef, keyUsage uint32, date corefoundation.CFDateRef) int32
var _secCertificateSetPreferenceErr error

// CanCallSecCertificateSetPreference reports whether the symbol for SecCertificateSetPreference is available on this system.
func CanCallSecCertificateSetPreference() bool {
	return _secCertificateSetPreference != nil
}

// SecCertificateSetPreferenceCallError returns the symbol lookup or registration error for SecCertificateSetPreference.
func SecCertificateSetPreferenceCallError() error {
	if _secCertificateSetPreference != nil {
		return nil
	}
	return symbolCallError("SecCertificateSetPreference", "10.0", _secCertificateSetPreferenceErr)
}

// TrySecCertificateSetPreference calls SecCertificateSetPreference without panicking when the symbol is unavailable.
func TrySecCertificateSetPreference(certificate SecCertificateRef, name corefoundation.CFStringRef, keyUsage uint32, date corefoundation.CFDateRef) (int32, error) {
	if err := SecCertificateSetPreferenceCallError(); err != nil {
		return 0, err
	}
	return _secCertificateSetPreference(certificate, name, keyUsage, date), nil
}

// SecCertificateSetPreference sets the preferred certificate for a specified name, key use, and date.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateSetPreference
func SecCertificateSetPreference(certificate SecCertificateRef, name corefoundation.CFStringRef, keyUsage uint32, date corefoundation.CFDateRef) int32 {
	result, callErr := TrySecCertificateSetPreference(certificate, name, keyUsage, date)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCertificateSetPreferred func(certificate SecCertificateRef, name corefoundation.CFStringRef, keyUsage corefoundation.CFArrayRef) int32
var _secCertificateSetPreferredErr error

// CanCallSecCertificateSetPreferred reports whether the symbol for SecCertificateSetPreferred is available on this system.
func CanCallSecCertificateSetPreferred() bool {
	return _secCertificateSetPreferred != nil
}

// SecCertificateSetPreferredCallError returns the symbol lookup or registration error for SecCertificateSetPreferred.
func SecCertificateSetPreferredCallError() error {
	if _secCertificateSetPreferred != nil {
		return nil
	}
	return symbolCallError("SecCertificateSetPreferred", "10.7", _secCertificateSetPreferredErr)
}

// TrySecCertificateSetPreferred calls SecCertificateSetPreferred without panicking when the symbol is unavailable.
func TrySecCertificateSetPreferred(certificate SecCertificateRef, name corefoundation.CFStringRef, keyUsage corefoundation.CFArrayRef) (int32, error) {
	if err := SecCertificateSetPreferredCallError(); err != nil {
		return 0, err
	}
	return _secCertificateSetPreferred(certificate, name, keyUsage), nil
}

// SecCertificateSetPreferred sets the certificate that should be preferred for the specified name and key use.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateSetPreferred(_:_:_:)
func SecCertificateSetPreferred(certificate SecCertificateRef, name corefoundation.CFStringRef, keyUsage corefoundation.CFArrayRef) int32 {
	result, callErr := TrySecCertificateSetPreferred(certificate, name, keyUsage)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCodeCheckValidity func(code SecCodeRef, flags SecCSFlags, requirement SecRequirementRef) int32
var _secCodeCheckValidityErr error

// CanCallSecCodeCheckValidity reports whether the symbol for SecCodeCheckValidity is available on this system.
func CanCallSecCodeCheckValidity() bool {
	return _secCodeCheckValidity != nil
}

// SecCodeCheckValidityCallError returns the symbol lookup or registration error for SecCodeCheckValidity.
func SecCodeCheckValidityCallError() error {
	if _secCodeCheckValidity != nil {
		return nil
	}
	return symbolCallError("SecCodeCheckValidity", "10.0", _secCodeCheckValidityErr)
}

// TrySecCodeCheckValidity calls SecCodeCheckValidity without panicking when the symbol is unavailable.
func TrySecCodeCheckValidity(code SecCodeRef, flags SecCSFlags, requirement SecRequirementRef) (int32, error) {
	if err := SecCodeCheckValidityCallError(); err != nil {
		return 0, err
	}
	return _secCodeCheckValidity(code, flags, requirement), nil
}

// SecCodeCheckValidity performs dynamic validation of signed code.
//
// See: https://developer.apple.com/documentation/Security/SecCodeCheckValidity(_:_:_:)
func SecCodeCheckValidity(code SecCodeRef, flags SecCSFlags, requirement SecRequirementRef) int32 {
	result, callErr := TrySecCodeCheckValidity(code, flags, requirement)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCodeCheckValidityWithErrors func(code SecCodeRef, flags SecCSFlags, requirement SecRequirementRef, errors *corefoundation.CFErrorRef) int32
var _secCodeCheckValidityWithErrorsErr error

// CanCallSecCodeCheckValidityWithErrors reports whether the symbol for SecCodeCheckValidityWithErrors is available on this system.
func CanCallSecCodeCheckValidityWithErrors() bool {
	return _secCodeCheckValidityWithErrors != nil
}

// SecCodeCheckValidityWithErrorsCallError returns the symbol lookup or registration error for SecCodeCheckValidityWithErrors.
func SecCodeCheckValidityWithErrorsCallError() error {
	if _secCodeCheckValidityWithErrors != nil {
		return nil
	}
	return symbolCallError("SecCodeCheckValidityWithErrors", "10.0", _secCodeCheckValidityWithErrorsErr)
}

// TrySecCodeCheckValidityWithErrors calls SecCodeCheckValidityWithErrors without panicking when the symbol is unavailable.
func TrySecCodeCheckValidityWithErrors(code SecCodeRef, flags SecCSFlags, requirement SecRequirementRef, errors *corefoundation.CFErrorRef) (int32, error) {
	if err := SecCodeCheckValidityWithErrorsCallError(); err != nil {
		return 0, err
	}
	return _secCodeCheckValidityWithErrors(code, flags, requirement, errors), nil
}

// SecCodeCheckValidityWithErrors performs dynamic validation of signed code and returns detailed error information in the case of failure.
//
// See: https://developer.apple.com/documentation/Security/SecCodeCheckValidityWithErrors(_:_:_:_:)
func SecCodeCheckValidityWithErrors(code SecCodeRef, flags SecCSFlags, requirement SecRequirementRef, errors *corefoundation.CFErrorRef) int32 {
	result, callErr := TrySecCodeCheckValidityWithErrors(code, flags, requirement, errors)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCodeCopyDesignatedRequirement func(code SecStaticCodeRef, flags SecCSFlags, requirement *SecRequirementRef) int32
var _secCodeCopyDesignatedRequirementErr error

// CanCallSecCodeCopyDesignatedRequirement reports whether the symbol for SecCodeCopyDesignatedRequirement is available on this system.
func CanCallSecCodeCopyDesignatedRequirement() bool {
	return _secCodeCopyDesignatedRequirement != nil
}

// SecCodeCopyDesignatedRequirementCallError returns the symbol lookup or registration error for SecCodeCopyDesignatedRequirement.
func SecCodeCopyDesignatedRequirementCallError() error {
	if _secCodeCopyDesignatedRequirement != nil {
		return nil
	}
	return symbolCallError("SecCodeCopyDesignatedRequirement", "10.0", _secCodeCopyDesignatedRequirementErr)
}

// TrySecCodeCopyDesignatedRequirement calls SecCodeCopyDesignatedRequirement without panicking when the symbol is unavailable.
func TrySecCodeCopyDesignatedRequirement(code SecStaticCodeRef, flags SecCSFlags, requirement *SecRequirementRef) (int32, error) {
	if err := SecCodeCopyDesignatedRequirementCallError(); err != nil {
		return 0, err
	}
	return _secCodeCopyDesignatedRequirement(code, flags, requirement), nil
}

// SecCodeCopyDesignatedRequirement retrieves the designated code requirement of signed code.
//
// See: https://developer.apple.com/documentation/Security/SecCodeCopyDesignatedRequirement(_:_:_:)
func SecCodeCopyDesignatedRequirement(code SecStaticCodeRef, flags SecCSFlags, requirement *SecRequirementRef) int32 {
	result, callErr := TrySecCodeCopyDesignatedRequirement(code, flags, requirement)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCodeCopyGuestWithAttributes func(host SecCodeRef, attributes corefoundation.CFDictionaryRef, flags SecCSFlags, guest *SecCodeRef) int32
var _secCodeCopyGuestWithAttributesErr error

// CanCallSecCodeCopyGuestWithAttributes reports whether the symbol for SecCodeCopyGuestWithAttributes is available on this system.
func CanCallSecCodeCopyGuestWithAttributes() bool {
	return _secCodeCopyGuestWithAttributes != nil
}

// SecCodeCopyGuestWithAttributesCallError returns the symbol lookup or registration error for SecCodeCopyGuestWithAttributes.
func SecCodeCopyGuestWithAttributesCallError() error {
	if _secCodeCopyGuestWithAttributes != nil {
		return nil
	}
	return symbolCallError("SecCodeCopyGuestWithAttributes", "10.0", _secCodeCopyGuestWithAttributesErr)
}

// TrySecCodeCopyGuestWithAttributes calls SecCodeCopyGuestWithAttributes without panicking when the symbol is unavailable.
func TrySecCodeCopyGuestWithAttributes(host SecCodeRef, attributes corefoundation.CFDictionaryRef, flags SecCSFlags, guest *SecCodeRef) (int32, error) {
	if err := SecCodeCopyGuestWithAttributesCallError(); err != nil {
		return 0, err
	}
	return _secCodeCopyGuestWithAttributes(host, attributes, flags, guest), nil
}

// SecCodeCopyGuestWithAttributes asks a code host to identify one of its guests given the type and value of specific attributes of the guest code.
//
// See: https://developer.apple.com/documentation/Security/SecCodeCopyGuestWithAttributes(_:_:_:_:)
func SecCodeCopyGuestWithAttributes(host SecCodeRef, attributes corefoundation.CFDictionaryRef, flags SecCSFlags, guest *SecCodeRef) int32 {
	result, callErr := TrySecCodeCopyGuestWithAttributes(host, attributes, flags, guest)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCodeCopyHost func(guest SecCodeRef, flags SecCSFlags, host *SecCodeRef) int32
var _secCodeCopyHostErr error

// CanCallSecCodeCopyHost reports whether the symbol for SecCodeCopyHost is available on this system.
func CanCallSecCodeCopyHost() bool {
	return _secCodeCopyHost != nil
}

// SecCodeCopyHostCallError returns the symbol lookup or registration error for SecCodeCopyHost.
func SecCodeCopyHostCallError() error {
	if _secCodeCopyHost != nil {
		return nil
	}
	return symbolCallError("SecCodeCopyHost", "10.0", _secCodeCopyHostErr)
}

// TrySecCodeCopyHost calls SecCodeCopyHost without panicking when the symbol is unavailable.
func TrySecCodeCopyHost(guest SecCodeRef, flags SecCSFlags, host *SecCodeRef) (int32, error) {
	if err := SecCodeCopyHostCallError(); err != nil {
		return 0, err
	}
	return _secCodeCopyHost(guest, flags, host), nil
}

// SecCodeCopyHost retrieves the code object for the host of specified guest code.
//
// See: https://developer.apple.com/documentation/Security/SecCodeCopyHost(_:_:_:)
func SecCodeCopyHost(guest SecCodeRef, flags SecCSFlags, host *SecCodeRef) int32 {
	result, callErr := TrySecCodeCopyHost(guest, flags, host)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCodeCopyPath func(staticCode SecStaticCodeRef, flags SecCSFlags, path *corefoundation.CFURLRef) int32
var _secCodeCopyPathErr error

// CanCallSecCodeCopyPath reports whether the symbol for SecCodeCopyPath is available on this system.
func CanCallSecCodeCopyPath() bool {
	return _secCodeCopyPath != nil
}

// SecCodeCopyPathCallError returns the symbol lookup or registration error for SecCodeCopyPath.
func SecCodeCopyPathCallError() error {
	if _secCodeCopyPath != nil {
		return nil
	}
	return symbolCallError("SecCodeCopyPath", "10.0", _secCodeCopyPathErr)
}

// TrySecCodeCopyPath calls SecCodeCopyPath without panicking when the symbol is unavailable.
func TrySecCodeCopyPath(staticCode SecStaticCodeRef, flags SecCSFlags, path *corefoundation.CFURLRef) (int32, error) {
	if err := SecCodeCopyPathCallError(); err != nil {
		return 0, err
	}
	return _secCodeCopyPath(staticCode, flags, path), nil
}

// SecCodeCopyPath retrieves the location on disk of signed code, given a code or static code object.
//
// See: https://developer.apple.com/documentation/Security/SecCodeCopyPath(_:_:_:)
func SecCodeCopyPath(staticCode SecStaticCodeRef, flags SecCSFlags, path *corefoundation.CFURLRef) int32 {
	result, callErr := TrySecCodeCopyPath(staticCode, flags, path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCodeCopySelf func(flags SecCSFlags, self *SecCodeRef) int32
var _secCodeCopySelfErr error

// CanCallSecCodeCopySelf reports whether the symbol for SecCodeCopySelf is available on this system.
func CanCallSecCodeCopySelf() bool {
	return _secCodeCopySelf != nil
}

// SecCodeCopySelfCallError returns the symbol lookup or registration error for SecCodeCopySelf.
func SecCodeCopySelfCallError() error {
	if _secCodeCopySelf != nil {
		return nil
	}
	return symbolCallError("SecCodeCopySelf", "10.0", _secCodeCopySelfErr)
}

// TrySecCodeCopySelf calls SecCodeCopySelf without panicking when the symbol is unavailable.
func TrySecCodeCopySelf(flags SecCSFlags, self *SecCodeRef) (int32, error) {
	if err := SecCodeCopySelfCallError(); err != nil {
		return 0, err
	}
	return _secCodeCopySelf(flags, self), nil
}

// SecCodeCopySelf retrieves the code object for the code making the call.
//
// See: https://developer.apple.com/documentation/Security/SecCodeCopySelf(_:_:)
func SecCodeCopySelf(flags SecCSFlags, self *SecCodeRef) int32 {
	result, callErr := TrySecCodeCopySelf(flags, self)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCodeCopySigningInformation func(code SecStaticCodeRef, flags SecCSFlags, information *corefoundation.CFDictionaryRef) int32
var _secCodeCopySigningInformationErr error

// CanCallSecCodeCopySigningInformation reports whether the symbol for SecCodeCopySigningInformation is available on this system.
func CanCallSecCodeCopySigningInformation() bool {
	return _secCodeCopySigningInformation != nil
}

// SecCodeCopySigningInformationCallError returns the symbol lookup or registration error for SecCodeCopySigningInformation.
func SecCodeCopySigningInformationCallError() error {
	if _secCodeCopySigningInformation != nil {
		return nil
	}
	return symbolCallError("SecCodeCopySigningInformation", "10.0", _secCodeCopySigningInformationErr)
}

// TrySecCodeCopySigningInformation calls SecCodeCopySigningInformation without panicking when the symbol is unavailable.
func TrySecCodeCopySigningInformation(code SecStaticCodeRef, flags SecCSFlags, information *corefoundation.CFDictionaryRef) (int32, error) {
	if err := SecCodeCopySigningInformationCallError(); err != nil {
		return 0, err
	}
	return _secCodeCopySigningInformation(code, flags, information), nil
}

// SecCodeCopySigningInformation retrieves various pieces of information from a code signature.
//
// See: https://developer.apple.com/documentation/Security/SecCodeCopySigningInformation(_:_:_:)
func SecCodeCopySigningInformation(code SecStaticCodeRef, flags SecCSFlags, information *corefoundation.CFDictionaryRef) int32 {
	result, callErr := TrySecCodeCopySigningInformation(code, flags, information)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCodeCopyStaticCode func(code SecCodeRef, flags SecCSFlags, staticCode *SecStaticCodeRef) int32
var _secCodeCopyStaticCodeErr error

// CanCallSecCodeCopyStaticCode reports whether the symbol for SecCodeCopyStaticCode is available on this system.
func CanCallSecCodeCopyStaticCode() bool {
	return _secCodeCopyStaticCode != nil
}

// SecCodeCopyStaticCodeCallError returns the symbol lookup or registration error for SecCodeCopyStaticCode.
func SecCodeCopyStaticCodeCallError() error {
	if _secCodeCopyStaticCode != nil {
		return nil
	}
	return symbolCallError("SecCodeCopyStaticCode", "10.0", _secCodeCopyStaticCodeErr)
}

// TrySecCodeCopyStaticCode calls SecCodeCopyStaticCode without panicking when the symbol is unavailable.
func TrySecCodeCopyStaticCode(code SecCodeRef, flags SecCSFlags, staticCode *SecStaticCodeRef) (int32, error) {
	if err := SecCodeCopyStaticCodeCallError(); err != nil {
		return 0, err
	}
	return _secCodeCopyStaticCode(code, flags, staticCode), nil
}

// SecCodeCopyStaticCode returns a static code object representing the on-disk version of the given running code.
//
// See: https://developer.apple.com/documentation/Security/SecCodeCopyStaticCode(_:_:_:)
func SecCodeCopyStaticCode(code SecCodeRef, flags SecCSFlags, staticCode *SecStaticCodeRef) int32 {
	result, callErr := TrySecCodeCopyStaticCode(code, flags, staticCode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCodeCreateWithXPCMessage func(message unsafe.Pointer, flags SecCSFlags, target *SecCodeRef) int32
var _secCodeCreateWithXPCMessageErr error

// CanCallSecCodeCreateWithXPCMessage reports whether the symbol for SecCodeCreateWithXPCMessage is available on this system.
func CanCallSecCodeCreateWithXPCMessage() bool {
	return _secCodeCreateWithXPCMessage != nil
}

// SecCodeCreateWithXPCMessageCallError returns the symbol lookup or registration error for SecCodeCreateWithXPCMessage.
func SecCodeCreateWithXPCMessageCallError() error {
	if _secCodeCreateWithXPCMessage != nil {
		return nil
	}
	return symbolCallError("SecCodeCreateWithXPCMessage", "10.0", _secCodeCreateWithXPCMessageErr)
}

// TrySecCodeCreateWithXPCMessage calls SecCodeCreateWithXPCMessage without panicking when the symbol is unavailable.
func TrySecCodeCreateWithXPCMessage(message unsafe.Pointer, flags SecCSFlags, target *SecCodeRef) (int32, error) {
	if err := SecCodeCreateWithXPCMessageCallError(); err != nil {
		return 0, err
	}
	return _secCodeCreateWithXPCMessage(message, flags, target), nil
}

// SecCodeCreateWithXPCMessage.
//
// See: https://developer.apple.com/documentation/Security/SecCodeCreateWithXPCMessage(_:_:_:)
func SecCodeCreateWithXPCMessage(message unsafe.Pointer, flags SecCSFlags, target *SecCodeRef) int32 {
	result, callErr := TrySecCodeCreateWithXPCMessage(message, flags, target)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCodeGetTypeID func() uint
var _secCodeGetTypeIDErr error

// CanCallSecCodeGetTypeID reports whether the symbol for SecCodeGetTypeID is available on this system.
func CanCallSecCodeGetTypeID() bool {
	return _secCodeGetTypeID != nil
}

// SecCodeGetTypeIDCallError returns the symbol lookup or registration error for SecCodeGetTypeID.
func SecCodeGetTypeIDCallError() error {
	if _secCodeGetTypeID != nil {
		return nil
	}
	return symbolCallError("SecCodeGetTypeID", "10.0", _secCodeGetTypeIDErr)
}

// TrySecCodeGetTypeID calls SecCodeGetTypeID without panicking when the symbol is unavailable.
func TrySecCodeGetTypeID() (uint, error) {
	if err := SecCodeGetTypeIDCallError(); err != nil {
		return 0, err
	}
	return _secCodeGetTypeID(), nil
}

// SecCodeGetTypeID returns the unique identifier of the opaque type to which a code object belongs.
//
// See: https://developer.apple.com/documentation/Security/SecCodeGetTypeID()
func SecCodeGetTypeID() uint {
	result, callErr := TrySecCodeGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCodeMapMemory func(code SecStaticCodeRef, flags SecCSFlags) int32
var _secCodeMapMemoryErr error

// CanCallSecCodeMapMemory reports whether the symbol for SecCodeMapMemory is available on this system.
func CanCallSecCodeMapMemory() bool {
	return _secCodeMapMemory != nil
}

// SecCodeMapMemoryCallError returns the symbol lookup or registration error for SecCodeMapMemory.
func SecCodeMapMemoryCallError() error {
	if _secCodeMapMemory != nil {
		return nil
	}
	return symbolCallError("SecCodeMapMemory", "10.0", _secCodeMapMemoryErr)
}

// TrySecCodeMapMemory calls SecCodeMapMemory without panicking when the symbol is unavailable.
func TrySecCodeMapMemory(code SecStaticCodeRef, flags SecCSFlags) (int32, error) {
	if err := SecCodeMapMemoryCallError(); err != nil {
		return 0, err
	}
	return _secCodeMapMemory(code, flags), nil
}

// SecCodeMapMemory asks the kernel to accept the signing information currently attached to a code object and uses it to validate memory page-ins.
//
// See: https://developer.apple.com/documentation/Security/SecCodeMapMemory(_:_:)
func SecCodeMapMemory(code SecStaticCodeRef, flags SecCSFlags) int32 {
	result, callErr := TrySecCodeMapMemory(code, flags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCodeValidateFileResource func(code SecStaticCodeRef, relativePath corefoundation.CFStringRef, fileData corefoundation.CFDataRef, flags SecCSFlags) int32
var _secCodeValidateFileResourceErr error

// CanCallSecCodeValidateFileResource reports whether the symbol for SecCodeValidateFileResource is available on this system.
func CanCallSecCodeValidateFileResource() bool {
	return _secCodeValidateFileResource != nil
}

// SecCodeValidateFileResourceCallError returns the symbol lookup or registration error for SecCodeValidateFileResource.
func SecCodeValidateFileResourceCallError() error {
	if _secCodeValidateFileResource != nil {
		return nil
	}
	return symbolCallError("SecCodeValidateFileResource", "10.13", _secCodeValidateFileResourceErr)
}

// TrySecCodeValidateFileResource calls SecCodeValidateFileResource without panicking when the symbol is unavailable.
func TrySecCodeValidateFileResource(code SecStaticCodeRef, relativePath corefoundation.CFStringRef, fileData corefoundation.CFDataRef, flags SecCSFlags) (int32, error) {
	if err := SecCodeValidateFileResourceCallError(); err != nil {
		return 0, err
	}
	return _secCodeValidateFileResource(code, relativePath, fileData, flags), nil
}

// SecCodeValidateFileResource.
//
// See: https://developer.apple.com/documentation/Security/SecCodeValidateFileResource(_:_:_:_:)
func SecCodeValidateFileResource(code SecStaticCodeRef, relativePath corefoundation.CFStringRef, fileData corefoundation.CFDataRef, flags SecCSFlags) int32 {
	result, callErr := TrySecCodeValidateFileResource(code, relativePath, fileData, flags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCopyErrorMessageString func(status int32, reserved unsafe.Pointer) corefoundation.CFStringRef
var _secCopyErrorMessageStringErr error

// CanCallSecCopyErrorMessageString reports whether the symbol for SecCopyErrorMessageString is available on this system.
func CanCallSecCopyErrorMessageString() bool {
	return _secCopyErrorMessageString != nil
}

// SecCopyErrorMessageStringCallError returns the symbol lookup or registration error for SecCopyErrorMessageString.
func SecCopyErrorMessageStringCallError() error {
	if _secCopyErrorMessageString != nil {
		return nil
	}
	return symbolCallError("SecCopyErrorMessageString", "10.3", _secCopyErrorMessageStringErr)
}

// TrySecCopyErrorMessageString calls SecCopyErrorMessageString without panicking when the symbol is unavailable.
func TrySecCopyErrorMessageString(status int32, reserved unsafe.Pointer) (corefoundation.CFStringRef, error) {
	if err := SecCopyErrorMessageStringCallError(); err != nil {
		return 0, err
	}
	return _secCopyErrorMessageString(status, reserved), nil
}

// SecCopyErrorMessageString returns a string explaining the meaning of a security result code.
//
// See: https://developer.apple.com/documentation/Security/SecCopyErrorMessageString(_:_:)
func SecCopyErrorMessageString(status int32, reserved unsafe.Pointer) corefoundation.CFStringRef {
	result, callErr := TrySecCopyErrorMessageString(status, reserved)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secCreateSharedWebCredentialPassword func() corefoundation.CFStringRef
var _secCreateSharedWebCredentialPasswordErr error

// CanCallSecCreateSharedWebCredentialPassword reports whether the symbol for SecCreateSharedWebCredentialPassword is available on this system.
func CanCallSecCreateSharedWebCredentialPassword() bool {
	return _secCreateSharedWebCredentialPassword != nil
}

// SecCreateSharedWebCredentialPasswordCallError returns the symbol lookup or registration error for SecCreateSharedWebCredentialPassword.
func SecCreateSharedWebCredentialPasswordCallError() error {
	if _secCreateSharedWebCredentialPassword != nil {
		return nil
	}
	return symbolCallError("SecCreateSharedWebCredentialPassword", "11.0", _secCreateSharedWebCredentialPasswordErr)
}

// TrySecCreateSharedWebCredentialPassword calls SecCreateSharedWebCredentialPassword without panicking when the symbol is unavailable.
func TrySecCreateSharedWebCredentialPassword() (corefoundation.CFStringRef, error) {
	if err := SecCreateSharedWebCredentialPasswordCallError(); err != nil {
		return 0, err
	}
	return _secCreateSharedWebCredentialPassword(), nil
}

// SecCreateSharedWebCredentialPassword returns a randomly generated password.
//
// See: https://developer.apple.com/documentation/Security/SecCreateSharedWebCredentialPassword()
func SecCreateSharedWebCredentialPassword() corefoundation.CFStringRef {
	result, callErr := TrySecCreateSharedWebCredentialPassword()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secHostCreateGuest func(host SecGuestRef, status uint32, path corefoundation.CFURLRef, attributes corefoundation.CFDictionaryRef, flags SecCSFlags, newGuest *SecGuestRef) int32
var _secHostCreateGuestErr error

// CanCallSecHostCreateGuest reports whether the symbol for SecHostCreateGuest is available on this system.
func CanCallSecHostCreateGuest() bool {
	return _secHostCreateGuest != nil
}

// SecHostCreateGuestCallError returns the symbol lookup or registration error for SecHostCreateGuest.
func SecHostCreateGuestCallError() error {
	if _secHostCreateGuest != nil {
		return nil
	}
	return symbolCallError("SecHostCreateGuest", "10.5", _secHostCreateGuestErr)
}

// TrySecHostCreateGuest calls SecHostCreateGuest without panicking when the symbol is unavailable.
func TrySecHostCreateGuest(host SecGuestRef, status uint32, path corefoundation.CFURLRef, attributes corefoundation.CFDictionaryRef, flags SecCSFlags, newGuest *SecGuestRef) (int32, error) {
	if err := SecHostCreateGuestCallError(); err != nil {
		return 0, err
	}
	return _secHostCreateGuest(host, status, path, attributes, flags, newGuest), nil
}

// SecHostCreateGuest creates a new guest and describes its initial properties.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/Security/SecHostCreateGuest
func SecHostCreateGuest(host SecGuestRef, status uint32, path corefoundation.CFURLRef, attributes corefoundation.CFDictionaryRef, flags SecCSFlags, newGuest *SecGuestRef) int32 {
	result, callErr := TrySecHostCreateGuest(host, status, path, attributes, flags, newGuest)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secHostRemoveGuest func(host SecGuestRef, guest SecGuestRef, flags SecCSFlags) int32
var _secHostRemoveGuestErr error

// CanCallSecHostRemoveGuest reports whether the symbol for SecHostRemoveGuest is available on this system.
func CanCallSecHostRemoveGuest() bool {
	return _secHostRemoveGuest != nil
}

// SecHostRemoveGuestCallError returns the symbol lookup or registration error for SecHostRemoveGuest.
func SecHostRemoveGuestCallError() error {
	if _secHostRemoveGuest != nil {
		return nil
	}
	return symbolCallError("SecHostRemoveGuest", "10.5", _secHostRemoveGuestErr)
}

// TrySecHostRemoveGuest calls SecHostRemoveGuest without panicking when the symbol is unavailable.
func TrySecHostRemoveGuest(host SecGuestRef, guest SecGuestRef, flags SecCSFlags) (int32, error) {
	if err := SecHostRemoveGuestCallError(); err != nil {
		return 0, err
	}
	return _secHostRemoveGuest(host, guest, flags), nil
}

// SecHostRemoveGuest removes a guest from a host.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/Security/SecHostRemoveGuest
func SecHostRemoveGuest(host SecGuestRef, guest SecGuestRef, flags SecCSFlags) int32 {
	result, callErr := TrySecHostRemoveGuest(host, guest, flags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secHostSelectGuest func(guestRef SecGuestRef, flags SecCSFlags) int32
var _secHostSelectGuestErr error

// CanCallSecHostSelectGuest reports whether the symbol for SecHostSelectGuest is available on this system.
func CanCallSecHostSelectGuest() bool {
	return _secHostSelectGuest != nil
}

// SecHostSelectGuestCallError returns the symbol lookup or registration error for SecHostSelectGuest.
func SecHostSelectGuestCallError() error {
	if _secHostSelectGuest != nil {
		return nil
	}
	return symbolCallError("SecHostSelectGuest", "10.5", _secHostSelectGuestErr)
}

// TrySecHostSelectGuest calls SecHostSelectGuest without panicking when the symbol is unavailable.
func TrySecHostSelectGuest(guestRef SecGuestRef, flags SecCSFlags) (int32, error) {
	if err := SecHostSelectGuestCallError(); err != nil {
		return 0, err
	}
	return _secHostSelectGuest(guestRef, flags), nil
}

// SecHostSelectGuest makes the calling thread the proxy for a specified guest.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/Security/SecHostSelectGuest
func SecHostSelectGuest(guestRef SecGuestRef, flags SecCSFlags) int32 {
	result, callErr := TrySecHostSelectGuest(guestRef, flags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secHostSelectedGuest func(flags SecCSFlags, guestRef *SecGuestRef) int32
var _secHostSelectedGuestErr error

// CanCallSecHostSelectedGuest reports whether the symbol for SecHostSelectedGuest is available on this system.
func CanCallSecHostSelectedGuest() bool {
	return _secHostSelectedGuest != nil
}

// SecHostSelectedGuestCallError returns the symbol lookup or registration error for SecHostSelectedGuest.
func SecHostSelectedGuestCallError() error {
	if _secHostSelectedGuest != nil {
		return nil
	}
	return symbolCallError("SecHostSelectedGuest", "10.5", _secHostSelectedGuestErr)
}

// TrySecHostSelectedGuest calls SecHostSelectedGuest without panicking when the symbol is unavailable.
func TrySecHostSelectedGuest(flags SecCSFlags, guestRef *SecGuestRef) (int32, error) {
	if err := SecHostSelectedGuestCallError(); err != nil {
		return 0, err
	}
	return _secHostSelectedGuest(flags, guestRef), nil
}

// SecHostSelectedGuest retrieves the handle for the guest currently selected for the calling thread.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/Security/SecHostSelectedGuest
func SecHostSelectedGuest(flags SecCSFlags, guestRef *SecGuestRef) int32 {
	result, callErr := TrySecHostSelectedGuest(flags, guestRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secHostSetGuestStatus func(guestRef SecGuestRef, status uint32, attributes corefoundation.CFDictionaryRef, flags SecCSFlags) int32
var _secHostSetGuestStatusErr error

// CanCallSecHostSetGuestStatus reports whether the symbol for SecHostSetGuestStatus is available on this system.
func CanCallSecHostSetGuestStatus() bool {
	return _secHostSetGuestStatus != nil
}

// SecHostSetGuestStatusCallError returns the symbol lookup or registration error for SecHostSetGuestStatus.
func SecHostSetGuestStatusCallError() error {
	if _secHostSetGuestStatus != nil {
		return nil
	}
	return symbolCallError("SecHostSetGuestStatus", "10.5", _secHostSetGuestStatusErr)
}

// TrySecHostSetGuestStatus calls SecHostSetGuestStatus without panicking when the symbol is unavailable.
func TrySecHostSetGuestStatus(guestRef SecGuestRef, status uint32, attributes corefoundation.CFDictionaryRef, flags SecCSFlags) (int32, error) {
	if err := SecHostSetGuestStatusCallError(); err != nil {
		return 0, err
	}
	return _secHostSetGuestStatus(guestRef, status, attributes, flags), nil
}

// SecHostSetGuestStatus updates the status and attributes of a particular guest.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/Security/SecHostSetGuestStatus
func SecHostSetGuestStatus(guestRef SecGuestRef, status uint32, attributes corefoundation.CFDictionaryRef, flags SecCSFlags) int32 {
	result, callErr := TrySecHostSetGuestStatus(guestRef, status, attributes, flags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secHostSetHostingPort func(hostingPort uint32, flags SecCSFlags) int32
var _secHostSetHostingPortErr error

// CanCallSecHostSetHostingPort reports whether the symbol for SecHostSetHostingPort is available on this system.
func CanCallSecHostSetHostingPort() bool {
	return _secHostSetHostingPort != nil
}

// SecHostSetHostingPortCallError returns the symbol lookup or registration error for SecHostSetHostingPort.
func SecHostSetHostingPortCallError() error {
	if _secHostSetHostingPort != nil {
		return nil
	}
	return symbolCallError("SecHostSetHostingPort", "10.5", _secHostSetHostingPortErr)
}

// TrySecHostSetHostingPort calls SecHostSetHostingPort without panicking when the symbol is unavailable.
func TrySecHostSetHostingPort(hostingPort uint32, flags SecCSFlags) (int32, error) {
	if err := SecHostSetHostingPortCallError(); err != nil {
		return 0, err
	}
	return _secHostSetHostingPort(hostingPort, flags), nil
}

// SecHostSetHostingPort tells code signing services that the calling code will directly respond to hosting inquiries over the given port.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/Security/SecHostSetHostingPort
func SecHostSetHostingPort(hostingPort uint32, flags SecCSFlags) int32 {
	result, callErr := TrySecHostSetHostingPort(hostingPort, flags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secIdentityCopyCertificate func(identityRef SecIdentityRef, certificateRef *SecCertificateRef) int32
var _secIdentityCopyCertificateErr error

// CanCallSecIdentityCopyCertificate reports whether the symbol for SecIdentityCopyCertificate is available on this system.
func CanCallSecIdentityCopyCertificate() bool {
	return _secIdentityCopyCertificate != nil
}

// SecIdentityCopyCertificateCallError returns the symbol lookup or registration error for SecIdentityCopyCertificate.
func SecIdentityCopyCertificateCallError() error {
	if _secIdentityCopyCertificate != nil {
		return nil
	}
	return symbolCallError("SecIdentityCopyCertificate", "10.3", _secIdentityCopyCertificateErr)
}

// TrySecIdentityCopyCertificate calls SecIdentityCopyCertificate without panicking when the symbol is unavailable.
func TrySecIdentityCopyCertificate(identityRef SecIdentityRef, certificateRef *SecCertificateRef) (int32, error) {
	if err := SecIdentityCopyCertificateCallError(); err != nil {
		return 0, err
	}
	return _secIdentityCopyCertificate(identityRef, certificateRef), nil
}

// SecIdentityCopyCertificate retrieves a certificate associated with an identity.
//
// See: https://developer.apple.com/documentation/Security/SecIdentityCopyCertificate(_:_:)
func SecIdentityCopyCertificate(identityRef SecIdentityRef, certificateRef *SecCertificateRef) int32 {
	result, callErr := TrySecIdentityCopyCertificate(identityRef, certificateRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secIdentityCopyPreference func(name corefoundation.CFStringRef, keyUsage CSSM_KEYUSE, validIssuers corefoundation.CFArrayRef, identity *SecIdentityRef) int32
var _secIdentityCopyPreferenceErr error

// CanCallSecIdentityCopyPreference reports whether the symbol for SecIdentityCopyPreference is available on this system.
func CanCallSecIdentityCopyPreference() bool {
	return _secIdentityCopyPreference != nil
}

// SecIdentityCopyPreferenceCallError returns the symbol lookup or registration error for SecIdentityCopyPreference.
func SecIdentityCopyPreferenceCallError() error {
	if _secIdentityCopyPreference != nil {
		return nil
	}
	return symbolCallError("SecIdentityCopyPreference", "10.0", _secIdentityCopyPreferenceErr)
}

// TrySecIdentityCopyPreference calls SecIdentityCopyPreference without panicking when the symbol is unavailable.
func TrySecIdentityCopyPreference(name corefoundation.CFStringRef, keyUsage CSSM_KEYUSE, validIssuers corefoundation.CFArrayRef, identity *SecIdentityRef) (int32, error) {
	if err := SecIdentityCopyPreferenceCallError(); err != nil {
		return 0, err
	}
	return _secIdentityCopyPreference(name, keyUsage, validIssuers, identity), nil
}

// SecIdentityCopyPreference returns the preferred identity for the specified name and key use.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecIdentityCopyPreference
func SecIdentityCopyPreference(name corefoundation.CFStringRef, keyUsage CSSM_KEYUSE, validIssuers corefoundation.CFArrayRef, identity *SecIdentityRef) int32 {
	result, callErr := TrySecIdentityCopyPreference(name, keyUsage, validIssuers, identity)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secIdentityCopyPreferred func(name corefoundation.CFStringRef, keyUsage corefoundation.CFArrayRef, validIssuers corefoundation.CFArrayRef) SecIdentityRef
var _secIdentityCopyPreferredErr error

// CanCallSecIdentityCopyPreferred reports whether the symbol for SecIdentityCopyPreferred is available on this system.
func CanCallSecIdentityCopyPreferred() bool {
	return _secIdentityCopyPreferred != nil
}

// SecIdentityCopyPreferredCallError returns the symbol lookup or registration error for SecIdentityCopyPreferred.
func SecIdentityCopyPreferredCallError() error {
	if _secIdentityCopyPreferred != nil {
		return nil
	}
	return symbolCallError("SecIdentityCopyPreferred", "10.7", _secIdentityCopyPreferredErr)
}

// TrySecIdentityCopyPreferred calls SecIdentityCopyPreferred without panicking when the symbol is unavailable.
func TrySecIdentityCopyPreferred(name corefoundation.CFStringRef, keyUsage corefoundation.CFArrayRef, validIssuers corefoundation.CFArrayRef) (SecIdentityRef, error) {
	if err := SecIdentityCopyPreferredCallError(); err != nil {
		return 0, err
	}
	return _secIdentityCopyPreferred(name, keyUsage, validIssuers), nil
}

// SecIdentityCopyPreferred retrieves the preferred identity for the specified name and key use.
//
// See: https://developer.apple.com/documentation/Security/SecIdentityCopyPreferred(_:_:_:)
func SecIdentityCopyPreferred(name corefoundation.CFStringRef, keyUsage corefoundation.CFArrayRef, validIssuers corefoundation.CFArrayRef) SecIdentityRef {
	result, callErr := TrySecIdentityCopyPreferred(name, keyUsage, validIssuers)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secIdentityCopyPrivateKey func(identityRef SecIdentityRef, privateKeyRef *SecKeyRef) int32
var _secIdentityCopyPrivateKeyErr error

// CanCallSecIdentityCopyPrivateKey reports whether the symbol for SecIdentityCopyPrivateKey is available on this system.
func CanCallSecIdentityCopyPrivateKey() bool {
	return _secIdentityCopyPrivateKey != nil
}

// SecIdentityCopyPrivateKeyCallError returns the symbol lookup or registration error for SecIdentityCopyPrivateKey.
func SecIdentityCopyPrivateKeyCallError() error {
	if _secIdentityCopyPrivateKey != nil {
		return nil
	}
	return symbolCallError("SecIdentityCopyPrivateKey", "10.3", _secIdentityCopyPrivateKeyErr)
}

// TrySecIdentityCopyPrivateKey calls SecIdentityCopyPrivateKey without panicking when the symbol is unavailable.
func TrySecIdentityCopyPrivateKey(identityRef SecIdentityRef, privateKeyRef *SecKeyRef) (int32, error) {
	if err := SecIdentityCopyPrivateKeyCallError(); err != nil {
		return 0, err
	}
	return _secIdentityCopyPrivateKey(identityRef, privateKeyRef), nil
}

// SecIdentityCopyPrivateKey retrieves the private key associated with an identity.
//
// See: https://developer.apple.com/documentation/Security/SecIdentityCopyPrivateKey(_:_:)
func SecIdentityCopyPrivateKey(identityRef SecIdentityRef, privateKeyRef *SecKeyRef) int32 {
	result, callErr := TrySecIdentityCopyPrivateKey(identityRef, privateKeyRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secIdentityCopySystemIdentity func(domain corefoundation.CFStringRef, idRef *SecIdentityRef, actualDomain *corefoundation.CFStringRef) int32
var _secIdentityCopySystemIdentityErr error

// CanCallSecIdentityCopySystemIdentity reports whether the symbol for SecIdentityCopySystemIdentity is available on this system.
func CanCallSecIdentityCopySystemIdentity() bool {
	return _secIdentityCopySystemIdentity != nil
}

// SecIdentityCopySystemIdentityCallError returns the symbol lookup or registration error for SecIdentityCopySystemIdentity.
func SecIdentityCopySystemIdentityCallError() error {
	if _secIdentityCopySystemIdentity != nil {
		return nil
	}
	return symbolCallError("SecIdentityCopySystemIdentity", "10.5", _secIdentityCopySystemIdentityErr)
}

// TrySecIdentityCopySystemIdentity calls SecIdentityCopySystemIdentity without panicking when the symbol is unavailable.
func TrySecIdentityCopySystemIdentity(domain corefoundation.CFStringRef, idRef *SecIdentityRef, actualDomain *corefoundation.CFStringRef) (int32, error) {
	if err := SecIdentityCopySystemIdentityCallError(); err != nil {
		return 0, err
	}
	return _secIdentityCopySystemIdentity(domain, idRef, actualDomain), nil
}

// SecIdentityCopySystemIdentity obtains the system identity associated with a specified domain.
//
// See: https://developer.apple.com/documentation/Security/SecIdentityCopySystemIdentity(_:_:_:)
func SecIdentityCopySystemIdentity(domain corefoundation.CFStringRef, idRef *SecIdentityRef, actualDomain *corefoundation.CFStringRef) int32 {
	result, callErr := TrySecIdentityCopySystemIdentity(domain, idRef, actualDomain)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secIdentityCreate func(allocator corefoundation.CFAllocatorRef, certificate SecCertificateRef, privateKey SecKeyRef) SecIdentityRef
var _secIdentityCreateErr error

// CanCallSecIdentityCreate reports whether the symbol for SecIdentityCreate is available on this system.
func CanCallSecIdentityCreate() bool {
	return _secIdentityCreate != nil
}

// SecIdentityCreateCallError returns the symbol lookup or registration error for SecIdentityCreate.
func SecIdentityCreateCallError() error {
	if _secIdentityCreate != nil {
		return nil
	}
	return symbolCallError("SecIdentityCreate", "10.12", _secIdentityCreateErr)
}

// TrySecIdentityCreate calls SecIdentityCreate without panicking when the symbol is unavailable.
func TrySecIdentityCreate(allocator corefoundation.CFAllocatorRef, certificate SecCertificateRef, privateKey SecKeyRef) (SecIdentityRef, error) {
	if err := SecIdentityCreateCallError(); err != nil {
		return 0, err
	}
	return _secIdentityCreate(allocator, certificate, privateKey), nil
}

// SecIdentityCreate.
//
// See: https://developer.apple.com/documentation/Security/SecIdentityCreate(_:_:_:)
func SecIdentityCreate(allocator corefoundation.CFAllocatorRef, certificate SecCertificateRef, privateKey SecKeyRef) SecIdentityRef {
	result, callErr := TrySecIdentityCreate(allocator, certificate, privateKey)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secIdentityCreateWithCertificate func(keychainOrArray corefoundation.CFTypeRef, certificateRef SecCertificateRef, identityRef *SecIdentityRef) int32
var _secIdentityCreateWithCertificateErr error

// CanCallSecIdentityCreateWithCertificate reports whether the symbol for SecIdentityCreateWithCertificate is available on this system.
func CanCallSecIdentityCreateWithCertificate() bool {
	return _secIdentityCreateWithCertificate != nil
}

// SecIdentityCreateWithCertificateCallError returns the symbol lookup or registration error for SecIdentityCreateWithCertificate.
func SecIdentityCreateWithCertificateCallError() error {
	if _secIdentityCreateWithCertificate != nil {
		return nil
	}
	return symbolCallError("SecIdentityCreateWithCertificate", "10.5", _secIdentityCreateWithCertificateErr)
}

// TrySecIdentityCreateWithCertificate calls SecIdentityCreateWithCertificate without panicking when the symbol is unavailable.
func TrySecIdentityCreateWithCertificate(keychainOrArray corefoundation.CFTypeRef, certificateRef SecCertificateRef, identityRef *SecIdentityRef) (int32, error) {
	if err := SecIdentityCreateWithCertificateCallError(); err != nil {
		return 0, err
	}
	return _secIdentityCreateWithCertificate(keychainOrArray, certificateRef, identityRef), nil
}

// SecIdentityCreateWithCertificate creates a new identity for a certificate and its associated private key.
//
// See: https://developer.apple.com/documentation/Security/SecIdentityCreateWithCertificate(_:_:_:)
func SecIdentityCreateWithCertificate(keychainOrArray corefoundation.CFTypeRef, certificateRef SecCertificateRef, identityRef *SecIdentityRef) int32 {
	result, callErr := TrySecIdentityCreateWithCertificate(keychainOrArray, certificateRef, identityRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secIdentityGetTypeID func() uint
var _secIdentityGetTypeIDErr error

// CanCallSecIdentityGetTypeID reports whether the symbol for SecIdentityGetTypeID is available on this system.
func CanCallSecIdentityGetTypeID() bool {
	return _secIdentityGetTypeID != nil
}

// SecIdentityGetTypeIDCallError returns the symbol lookup or registration error for SecIdentityGetTypeID.
func SecIdentityGetTypeIDCallError() error {
	if _secIdentityGetTypeID != nil {
		return nil
	}
	return symbolCallError("SecIdentityGetTypeID", "10.3", _secIdentityGetTypeIDErr)
}

// TrySecIdentityGetTypeID calls SecIdentityGetTypeID without panicking when the symbol is unavailable.
func TrySecIdentityGetTypeID() (uint, error) {
	if err := SecIdentityGetTypeIDCallError(); err != nil {
		return 0, err
	}
	return _secIdentityGetTypeID(), nil
}

// SecIdentityGetTypeID returns the unique identifier of the opaque type to which an identity object belongs.
//
// See: https://developer.apple.com/documentation/Security/SecIdentityGetTypeID()
func SecIdentityGetTypeID() uint {
	result, callErr := TrySecIdentityGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secIdentitySearchCopyNext func(searchRef SecIdentitySearchRef, identity *SecIdentityRef) int32
var _secIdentitySearchCopyNextErr error

// CanCallSecIdentitySearchCopyNext reports whether the symbol for SecIdentitySearchCopyNext is available on this system.
func CanCallSecIdentitySearchCopyNext() bool {
	return _secIdentitySearchCopyNext != nil
}

// SecIdentitySearchCopyNextCallError returns the symbol lookup or registration error for SecIdentitySearchCopyNext.
func SecIdentitySearchCopyNextCallError() error {
	if _secIdentitySearchCopyNext != nil {
		return nil
	}
	return symbolCallError("SecIdentitySearchCopyNext", "10.0", _secIdentitySearchCopyNextErr)
}

// TrySecIdentitySearchCopyNext calls SecIdentitySearchCopyNext without panicking when the symbol is unavailable.
func TrySecIdentitySearchCopyNext(searchRef SecIdentitySearchRef, identity *SecIdentityRef) (int32, error) {
	if err := SecIdentitySearchCopyNextCallError(); err != nil {
		return 0, err
	}
	return _secIdentitySearchCopyNext(searchRef, identity), nil
}

// SecIdentitySearchCopyNext finds the next identity matching specified search criteria
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecIdentitySearchCopyNext
func SecIdentitySearchCopyNext(searchRef SecIdentitySearchRef, identity *SecIdentityRef) int32 {
	result, callErr := TrySecIdentitySearchCopyNext(searchRef, identity)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secIdentitySearchCreate func(keychainOrArray corefoundation.CFTypeRef, keyUsage CSSM_KEYUSE, searchRef *SecIdentitySearchRef) int32
var _secIdentitySearchCreateErr error

// CanCallSecIdentitySearchCreate reports whether the symbol for SecIdentitySearchCreate is available on this system.
func CanCallSecIdentitySearchCreate() bool {
	return _secIdentitySearchCreate != nil
}

// SecIdentitySearchCreateCallError returns the symbol lookup or registration error for SecIdentitySearchCreate.
func SecIdentitySearchCreateCallError() error {
	if _secIdentitySearchCreate != nil {
		return nil
	}
	return symbolCallError("SecIdentitySearchCreate", "10.0", _secIdentitySearchCreateErr)
}

// TrySecIdentitySearchCreate calls SecIdentitySearchCreate without panicking when the symbol is unavailable.
func TrySecIdentitySearchCreate(keychainOrArray corefoundation.CFTypeRef, keyUsage CSSM_KEYUSE, searchRef *SecIdentitySearchRef) (int32, error) {
	if err := SecIdentitySearchCreateCallError(); err != nil {
		return 0, err
	}
	return _secIdentitySearchCreate(keychainOrArray, keyUsage, searchRef), nil
}

// SecIdentitySearchCreate creates a search object for finding identities.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecIdentitySearchCreate
func SecIdentitySearchCreate(keychainOrArray corefoundation.CFTypeRef, keyUsage CSSM_KEYUSE, searchRef *SecIdentitySearchRef) int32 {
	result, callErr := TrySecIdentitySearchCreate(keychainOrArray, keyUsage, searchRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secIdentitySearchGetTypeID func() uint
var _secIdentitySearchGetTypeIDErr error

// CanCallSecIdentitySearchGetTypeID reports whether the symbol for SecIdentitySearchGetTypeID is available on this system.
func CanCallSecIdentitySearchGetTypeID() bool {
	return _secIdentitySearchGetTypeID != nil
}

// SecIdentitySearchGetTypeIDCallError returns the symbol lookup or registration error for SecIdentitySearchGetTypeID.
func SecIdentitySearchGetTypeIDCallError() error {
	if _secIdentitySearchGetTypeID != nil {
		return nil
	}
	return symbolCallError("SecIdentitySearchGetTypeID", "10.0", _secIdentitySearchGetTypeIDErr)
}

// TrySecIdentitySearchGetTypeID calls SecIdentitySearchGetTypeID without panicking when the symbol is unavailable.
func TrySecIdentitySearchGetTypeID() (uint, error) {
	if err := SecIdentitySearchGetTypeIDCallError(); err != nil {
		return 0, err
	}
	return _secIdentitySearchGetTypeID(), nil
}

// SecIdentitySearchGetTypeID returns the unique identifier of the opaque type to which a [SecIdentitySearch] object belongs.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecIdentitySearchGetTypeID
func SecIdentitySearchGetTypeID() uint {
	result, callErr := TrySecIdentitySearchGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secIdentitySetPreference func(identity SecIdentityRef, name corefoundation.CFStringRef, keyUsage CSSM_KEYUSE) int32
var _secIdentitySetPreferenceErr error

// CanCallSecIdentitySetPreference reports whether the symbol for SecIdentitySetPreference is available on this system.
func CanCallSecIdentitySetPreference() bool {
	return _secIdentitySetPreference != nil
}

// SecIdentitySetPreferenceCallError returns the symbol lookup or registration error for SecIdentitySetPreference.
func SecIdentitySetPreferenceCallError() error {
	if _secIdentitySetPreference != nil {
		return nil
	}
	return symbolCallError("SecIdentitySetPreference", "10.0", _secIdentitySetPreferenceErr)
}

// TrySecIdentitySetPreference calls SecIdentitySetPreference without panicking when the symbol is unavailable.
func TrySecIdentitySetPreference(identity SecIdentityRef, name corefoundation.CFStringRef, keyUsage CSSM_KEYUSE) (int32, error) {
	if err := SecIdentitySetPreferenceCallError(); err != nil {
		return 0, err
	}
	return _secIdentitySetPreference(identity, name, keyUsage), nil
}

// SecIdentitySetPreference sets the preferred identity for the specified name and key use.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecIdentitySetPreference
func SecIdentitySetPreference(identity SecIdentityRef, name corefoundation.CFStringRef, keyUsage CSSM_KEYUSE) int32 {
	result, callErr := TrySecIdentitySetPreference(identity, name, keyUsage)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secIdentitySetPreferred func(identity SecIdentityRef, name corefoundation.CFStringRef, keyUsage corefoundation.CFArrayRef) int32
var _secIdentitySetPreferredErr error

// CanCallSecIdentitySetPreferred reports whether the symbol for SecIdentitySetPreferred is available on this system.
func CanCallSecIdentitySetPreferred() bool {
	return _secIdentitySetPreferred != nil
}

// SecIdentitySetPreferredCallError returns the symbol lookup or registration error for SecIdentitySetPreferred.
func SecIdentitySetPreferredCallError() error {
	if _secIdentitySetPreferred != nil {
		return nil
	}
	return symbolCallError("SecIdentitySetPreferred", "10.7", _secIdentitySetPreferredErr)
}

// TrySecIdentitySetPreferred calls SecIdentitySetPreferred without panicking when the symbol is unavailable.
func TrySecIdentitySetPreferred(identity SecIdentityRef, name corefoundation.CFStringRef, keyUsage corefoundation.CFArrayRef) (int32, error) {
	if err := SecIdentitySetPreferredCallError(); err != nil {
		return 0, err
	}
	return _secIdentitySetPreferred(identity, name, keyUsage), nil
}

// SecIdentitySetPreferred sets the identity that should be preferred for the specified name and key use.
//
// See: https://developer.apple.com/documentation/Security/SecIdentitySetPreferred(_:_:_:)
func SecIdentitySetPreferred(identity SecIdentityRef, name corefoundation.CFStringRef, keyUsage corefoundation.CFArrayRef) int32 {
	result, callErr := TrySecIdentitySetPreferred(identity, name, keyUsage)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secIdentitySetSystemIdentity func(domain corefoundation.CFStringRef, idRef SecIdentityRef) int32
var _secIdentitySetSystemIdentityErr error

// CanCallSecIdentitySetSystemIdentity reports whether the symbol for SecIdentitySetSystemIdentity is available on this system.
func CanCallSecIdentitySetSystemIdentity() bool {
	return _secIdentitySetSystemIdentity != nil
}

// SecIdentitySetSystemIdentityCallError returns the symbol lookup or registration error for SecIdentitySetSystemIdentity.
func SecIdentitySetSystemIdentityCallError() error {
	if _secIdentitySetSystemIdentity != nil {
		return nil
	}
	return symbolCallError("SecIdentitySetSystemIdentity", "10.5", _secIdentitySetSystemIdentityErr)
}

// TrySecIdentitySetSystemIdentity calls SecIdentitySetSystemIdentity without panicking when the symbol is unavailable.
func TrySecIdentitySetSystemIdentity(domain corefoundation.CFStringRef, idRef SecIdentityRef) (int32, error) {
	if err := SecIdentitySetSystemIdentityCallError(); err != nil {
		return 0, err
	}
	return _secIdentitySetSystemIdentity(domain, idRef), nil
}

// SecIdentitySetSystemIdentity assigns the system identity to be associated with a specified domain.
//
// See: https://developer.apple.com/documentation/Security/SecIdentitySetSystemIdentity(_:_:)
func SecIdentitySetSystemIdentity(domain corefoundation.CFStringRef, idRef SecIdentityRef) int32 {
	result, callErr := TrySecIdentitySetSystemIdentity(domain, idRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secItemAdd func(attributes corefoundation.CFDictionaryRef, result *corefoundation.CFTypeRef) int32
var _secItemAddErr error

// CanCallSecItemAdd reports whether the symbol for SecItemAdd is available on this system.
func CanCallSecItemAdd() bool {
	return _secItemAdd != nil
}

// SecItemAddCallError returns the symbol lookup or registration error for SecItemAdd.
func SecItemAddCallError() error {
	if _secItemAdd != nil {
		return nil
	}
	return symbolCallError("SecItemAdd", "10.6", _secItemAddErr)
}

// TrySecItemAdd calls SecItemAdd without panicking when the symbol is unavailable.
func TrySecItemAdd(attributes corefoundation.CFDictionaryRef, result *corefoundation.CFTypeRef) (int32, error) {
	if err := SecItemAddCallError(); err != nil {
		return 0, err
	}
	return _secItemAdd(attributes, result), nil
}

// SecItemAdd adds one or more items to a keychain.
//
// See: https://developer.apple.com/documentation/Security/SecItemAdd(_:_:)
func SecItemAdd(attributes corefoundation.CFDictionaryRef, result *corefoundation.CFTypeRef) int32 {
	result0, callErr := TrySecItemAdd(attributes, result)
	if callErr != nil {
		panic(callErr)
	}
	return result0
}

var _secItemCopyMatching func(query corefoundation.CFDictionaryRef, result *corefoundation.CFTypeRef) int32
var _secItemCopyMatchingErr error

// CanCallSecItemCopyMatching reports whether the symbol for SecItemCopyMatching is available on this system.
func CanCallSecItemCopyMatching() bool {
	return _secItemCopyMatching != nil
}

// SecItemCopyMatchingCallError returns the symbol lookup or registration error for SecItemCopyMatching.
func SecItemCopyMatchingCallError() error {
	if _secItemCopyMatching != nil {
		return nil
	}
	return symbolCallError("SecItemCopyMatching", "10.6", _secItemCopyMatchingErr)
}

// TrySecItemCopyMatching calls SecItemCopyMatching without panicking when the symbol is unavailable.
func TrySecItemCopyMatching(query corefoundation.CFDictionaryRef, result *corefoundation.CFTypeRef) (int32, error) {
	if err := SecItemCopyMatchingCallError(); err != nil {
		return 0, err
	}
	return _secItemCopyMatching(query, result), nil
}

// SecItemCopyMatching returns one or more keychain items that match a search query, or copies attributes of specific keychain items.
//
// See: https://developer.apple.com/documentation/Security/SecItemCopyMatching(_:_:)
func SecItemCopyMatching(query corefoundation.CFDictionaryRef, result *corefoundation.CFTypeRef) int32 {
	result0, callErr := TrySecItemCopyMatching(query, result)
	if callErr != nil {
		panic(callErr)
	}
	return result0
}

var _secItemDelete func(query corefoundation.CFDictionaryRef) int32
var _secItemDeleteErr error

// CanCallSecItemDelete reports whether the symbol for SecItemDelete is available on this system.
func CanCallSecItemDelete() bool {
	return _secItemDelete != nil
}

// SecItemDeleteCallError returns the symbol lookup or registration error for SecItemDelete.
func SecItemDeleteCallError() error {
	if _secItemDelete != nil {
		return nil
	}
	return symbolCallError("SecItemDelete", "10.6", _secItemDeleteErr)
}

// TrySecItemDelete calls SecItemDelete without panicking when the symbol is unavailable.
func TrySecItemDelete(query corefoundation.CFDictionaryRef) (int32, error) {
	if err := SecItemDeleteCallError(); err != nil {
		return 0, err
	}
	return _secItemDelete(query), nil
}

// SecItemDelete deletes items that match a search query.
//
// See: https://developer.apple.com/documentation/Security/SecItemDelete(_:)
func SecItemDelete(query corefoundation.CFDictionaryRef) int32 {
	result, callErr := TrySecItemDelete(query)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secItemExport func(secItemOrArray corefoundation.CFTypeRef, outputFormat SecExternalFormat, flags SecItemImportExportFlags, keyParams *SecItemImportExportKeyParameters, exportedData *corefoundation.CFDataRef) int32
var _secItemExportErr error

// CanCallSecItemExport reports whether the symbol for SecItemExport is available on this system.
func CanCallSecItemExport() bool {
	return _secItemExport != nil
}

// SecItemExportCallError returns the symbol lookup or registration error for SecItemExport.
func SecItemExportCallError() error {
	if _secItemExport != nil {
		return nil
	}
	return symbolCallError("SecItemExport", "10.7", _secItemExportErr)
}

// TrySecItemExport calls SecItemExport without panicking when the symbol is unavailable.
func TrySecItemExport(secItemOrArray corefoundation.CFTypeRef, outputFormat SecExternalFormat, flags SecItemImportExportFlags, keyParams *SecItemImportExportKeyParameters, exportedData *corefoundation.CFDataRef) (int32, error) {
	if err := SecItemExportCallError(); err != nil {
		return 0, err
	}
	return _secItemExport(secItemOrArray, outputFormat, flags, keyParams, exportedData), nil
}

// SecItemExport exports one or more certificates, keys, or identities.
//
// See: https://developer.apple.com/documentation/Security/SecItemExport(_:_:_:_:_:)
func SecItemExport(secItemOrArray corefoundation.CFTypeRef, outputFormat SecExternalFormat, flags SecItemImportExportFlags, keyParams *SecItemImportExportKeyParameters, exportedData *corefoundation.CFDataRef) int32 {
	result, callErr := TrySecItemExport(secItemOrArray, outputFormat, flags, keyParams, exportedData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secItemImport func(importedData corefoundation.CFDataRef, fileNameOrExtension corefoundation.CFStringRef, inputFormat *SecExternalFormat, itemType *SecExternalItemType, flags SecItemImportExportFlags, keyParams *SecItemImportExportKeyParameters, importKeychain SecKeychainRef, outItems *corefoundation.CFArrayRef) int32
var _secItemImportErr error

// CanCallSecItemImport reports whether the symbol for SecItemImport is available on this system.
func CanCallSecItemImport() bool {
	return _secItemImport != nil
}

// SecItemImportCallError returns the symbol lookup or registration error for SecItemImport.
func SecItemImportCallError() error {
	if _secItemImport != nil {
		return nil
	}
	return symbolCallError("SecItemImport", "10.7", _secItemImportErr)
}

// TrySecItemImport calls SecItemImport without panicking when the symbol is unavailable.
func TrySecItemImport(importedData corefoundation.CFDataRef, fileNameOrExtension corefoundation.CFStringRef, inputFormat *SecExternalFormat, itemType *SecExternalItemType, flags SecItemImportExportFlags, keyParams *SecItemImportExportKeyParameters, importKeychain SecKeychainRef, outItems *corefoundation.CFArrayRef) (int32, error) {
	if err := SecItemImportCallError(); err != nil {
		return 0, err
	}
	return _secItemImport(importedData, fileNameOrExtension, inputFormat, itemType, flags, keyParams, importKeychain, outItems), nil
}

// SecItemImport imports one or more certificates, keys, or identities and optionally adds them to a keychain.
//
// See: https://developer.apple.com/documentation/Security/SecItemImport(_:_:_:_:_:_:_:_:)
func SecItemImport(importedData corefoundation.CFDataRef, fileNameOrExtension corefoundation.CFStringRef, inputFormat *SecExternalFormat, itemType *SecExternalItemType, flags SecItemImportExportFlags, keyParams *SecItemImportExportKeyParameters, importKeychain SecKeychainRef, outItems *corefoundation.CFArrayRef) int32 {
	result, callErr := TrySecItemImport(importedData, fileNameOrExtension, inputFormat, itemType, flags, keyParams, importKeychain, outItems)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secItemUpdate func(query corefoundation.CFDictionaryRef, attributesToUpdate corefoundation.CFDictionaryRef) int32
var _secItemUpdateErr error

// CanCallSecItemUpdate reports whether the symbol for SecItemUpdate is available on this system.
func CanCallSecItemUpdate() bool {
	return _secItemUpdate != nil
}

// SecItemUpdateCallError returns the symbol lookup or registration error for SecItemUpdate.
func SecItemUpdateCallError() error {
	if _secItemUpdate != nil {
		return nil
	}
	return symbolCallError("SecItemUpdate", "10.6", _secItemUpdateErr)
}

// TrySecItemUpdate calls SecItemUpdate without panicking when the symbol is unavailable.
func TrySecItemUpdate(query corefoundation.CFDictionaryRef, attributesToUpdate corefoundation.CFDictionaryRef) (int32, error) {
	if err := SecItemUpdateCallError(); err != nil {
		return 0, err
	}
	return _secItemUpdate(query, attributesToUpdate), nil
}

// SecItemUpdate modifies items that match a search query.
//
// See: https://developer.apple.com/documentation/Security/SecItemUpdate(_:_:)
func SecItemUpdate(query corefoundation.CFDictionaryRef, attributesToUpdate corefoundation.CFDictionaryRef) int32 {
	result, callErr := TrySecItemUpdate(query, attributesToUpdate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secKeyCopyAttributes func(key SecKeyRef) corefoundation.CFDictionaryRef
var _secKeyCopyAttributesErr error

// CanCallSecKeyCopyAttributes reports whether the symbol for SecKeyCopyAttributes is available on this system.
func CanCallSecKeyCopyAttributes() bool {
	return _secKeyCopyAttributes != nil
}

// SecKeyCopyAttributesCallError returns the symbol lookup or registration error for SecKeyCopyAttributes.
func SecKeyCopyAttributesCallError() error {
	if _secKeyCopyAttributes != nil {
		return nil
	}
	return symbolCallError("SecKeyCopyAttributes", "10.12", _secKeyCopyAttributesErr)
}

// TrySecKeyCopyAttributes calls SecKeyCopyAttributes without panicking when the symbol is unavailable.
func TrySecKeyCopyAttributes(key SecKeyRef) (corefoundation.CFDictionaryRef, error) {
	if err := SecKeyCopyAttributesCallError(); err != nil {
		return 0, err
	}
	return _secKeyCopyAttributes(key), nil
}

// SecKeyCopyAttributes gets the attributes of a given key.
//
// See: https://developer.apple.com/documentation/Security/SecKeyCopyAttributes(_:)
func SecKeyCopyAttributes(key SecKeyRef) corefoundation.CFDictionaryRef {
	result, callErr := TrySecKeyCopyAttributes(key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secKeyCopyExternalRepresentation func(key SecKeyRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef
var _secKeyCopyExternalRepresentationErr error

// CanCallSecKeyCopyExternalRepresentation reports whether the symbol for SecKeyCopyExternalRepresentation is available on this system.
func CanCallSecKeyCopyExternalRepresentation() bool {
	return _secKeyCopyExternalRepresentation != nil
}

// SecKeyCopyExternalRepresentationCallError returns the symbol lookup or registration error for SecKeyCopyExternalRepresentation.
func SecKeyCopyExternalRepresentationCallError() error {
	if _secKeyCopyExternalRepresentation != nil {
		return nil
	}
	return symbolCallError("SecKeyCopyExternalRepresentation", "10.12", _secKeyCopyExternalRepresentationErr)
}

// TrySecKeyCopyExternalRepresentation calls SecKeyCopyExternalRepresentation without panicking when the symbol is unavailable.
func TrySecKeyCopyExternalRepresentation(key SecKeyRef, err *corefoundation.CFErrorRef) (corefoundation.CFDataRef, error) {
	if err := SecKeyCopyExternalRepresentationCallError(); err != nil {
		return 0, err
	}
	return _secKeyCopyExternalRepresentation(key, err), nil
}

// SecKeyCopyExternalRepresentation returns an external representation of the given key suitable for the key’s type.
//
// See: https://developer.apple.com/documentation/Security/SecKeyCopyExternalRepresentation(_:_:)
func SecKeyCopyExternalRepresentation(key SecKeyRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef {
	result, callErr := TrySecKeyCopyExternalRepresentation(key, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secKeyCopyKeyExchangeResult func(privateKey SecKeyRef, algorithm SecKeyAlgorithm, publicKey SecKeyRef, parameters corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef
var _secKeyCopyKeyExchangeResultErr error

// CanCallSecKeyCopyKeyExchangeResult reports whether the symbol for SecKeyCopyKeyExchangeResult is available on this system.
func CanCallSecKeyCopyKeyExchangeResult() bool {
	return _secKeyCopyKeyExchangeResult != nil
}

// SecKeyCopyKeyExchangeResultCallError returns the symbol lookup or registration error for SecKeyCopyKeyExchangeResult.
func SecKeyCopyKeyExchangeResultCallError() error {
	if _secKeyCopyKeyExchangeResult != nil {
		return nil
	}
	return symbolCallError("SecKeyCopyKeyExchangeResult", "10.12", _secKeyCopyKeyExchangeResultErr)
}

// TrySecKeyCopyKeyExchangeResult calls SecKeyCopyKeyExchangeResult without panicking when the symbol is unavailable.
func TrySecKeyCopyKeyExchangeResult(privateKey SecKeyRef, algorithm SecKeyAlgorithm, publicKey SecKeyRef, parameters corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) (corefoundation.CFDataRef, error) {
	if err := SecKeyCopyKeyExchangeResultCallError(); err != nil {
		return 0, err
	}
	return _secKeyCopyKeyExchangeResult(privateKey, algorithm, publicKey, parameters, err), nil
}

// SecKeyCopyKeyExchangeResult performs the Diffie-Hellman style of key exchange with optional key-derivation steps.
//
// See: https://developer.apple.com/documentation/Security/SecKeyCopyKeyExchangeResult(_:_:_:_:_:)
func SecKeyCopyKeyExchangeResult(privateKey SecKeyRef, algorithm SecKeyAlgorithm, publicKey SecKeyRef, parameters corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef {
	result, callErr := TrySecKeyCopyKeyExchangeResult(privateKey, algorithm, publicKey, parameters, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secKeyCopyPublicKey func(key SecKeyRef) SecKeyRef
var _secKeyCopyPublicKeyErr error

// CanCallSecKeyCopyPublicKey reports whether the symbol for SecKeyCopyPublicKey is available on this system.
func CanCallSecKeyCopyPublicKey() bool {
	return _secKeyCopyPublicKey != nil
}

// SecKeyCopyPublicKeyCallError returns the symbol lookup or registration error for SecKeyCopyPublicKey.
func SecKeyCopyPublicKeyCallError() error {
	if _secKeyCopyPublicKey != nil {
		return nil
	}
	return symbolCallError("SecKeyCopyPublicKey", "10.12", _secKeyCopyPublicKeyErr)
}

// TrySecKeyCopyPublicKey calls SecKeyCopyPublicKey without panicking when the symbol is unavailable.
func TrySecKeyCopyPublicKey(key SecKeyRef) (SecKeyRef, error) {
	if err := SecKeyCopyPublicKeyCallError(); err != nil {
		return 0, err
	}
	return _secKeyCopyPublicKey(key), nil
}

// SecKeyCopyPublicKey gets the public key associated with the given private key.
//
// See: https://developer.apple.com/documentation/Security/SecKeyCopyPublicKey(_:)
func SecKeyCopyPublicKey(key SecKeyRef) SecKeyRef {
	result, callErr := TrySecKeyCopyPublicKey(key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secKeyCreateDecryptedData func(key SecKeyRef, algorithm SecKeyAlgorithm, ciphertext corefoundation.CFDataRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef
var _secKeyCreateDecryptedDataErr error

// CanCallSecKeyCreateDecryptedData reports whether the symbol for SecKeyCreateDecryptedData is available on this system.
func CanCallSecKeyCreateDecryptedData() bool {
	return _secKeyCreateDecryptedData != nil
}

// SecKeyCreateDecryptedDataCallError returns the symbol lookup or registration error for SecKeyCreateDecryptedData.
func SecKeyCreateDecryptedDataCallError() error {
	if _secKeyCreateDecryptedData != nil {
		return nil
	}
	return symbolCallError("SecKeyCreateDecryptedData", "10.12", _secKeyCreateDecryptedDataErr)
}

// TrySecKeyCreateDecryptedData calls SecKeyCreateDecryptedData without panicking when the symbol is unavailable.
func TrySecKeyCreateDecryptedData(key SecKeyRef, algorithm SecKeyAlgorithm, ciphertext corefoundation.CFDataRef, err *corefoundation.CFErrorRef) (corefoundation.CFDataRef, error) {
	if err := SecKeyCreateDecryptedDataCallError(); err != nil {
		return 0, err
	}
	return _secKeyCreateDecryptedData(key, algorithm, ciphertext, err), nil
}

// SecKeyCreateDecryptedData decrypts a block of data using a private key and specified algorithm.
//
// See: https://developer.apple.com/documentation/Security/SecKeyCreateDecryptedData(_:_:_:_:)
func SecKeyCreateDecryptedData(key SecKeyRef, algorithm SecKeyAlgorithm, ciphertext corefoundation.CFDataRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef {
	result, callErr := TrySecKeyCreateDecryptedData(key, algorithm, ciphertext, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secKeyCreateEncryptedData func(key SecKeyRef, algorithm SecKeyAlgorithm, plaintext corefoundation.CFDataRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef
var _secKeyCreateEncryptedDataErr error

// CanCallSecKeyCreateEncryptedData reports whether the symbol for SecKeyCreateEncryptedData is available on this system.
func CanCallSecKeyCreateEncryptedData() bool {
	return _secKeyCreateEncryptedData != nil
}

// SecKeyCreateEncryptedDataCallError returns the symbol lookup or registration error for SecKeyCreateEncryptedData.
func SecKeyCreateEncryptedDataCallError() error {
	if _secKeyCreateEncryptedData != nil {
		return nil
	}
	return symbolCallError("SecKeyCreateEncryptedData", "10.12", _secKeyCreateEncryptedDataErr)
}

// TrySecKeyCreateEncryptedData calls SecKeyCreateEncryptedData without panicking when the symbol is unavailable.
func TrySecKeyCreateEncryptedData(key SecKeyRef, algorithm SecKeyAlgorithm, plaintext corefoundation.CFDataRef, err *corefoundation.CFErrorRef) (corefoundation.CFDataRef, error) {
	if err := SecKeyCreateEncryptedDataCallError(); err != nil {
		return 0, err
	}
	return _secKeyCreateEncryptedData(key, algorithm, plaintext, err), nil
}

// SecKeyCreateEncryptedData encrypts a block of data using a public key and specified algorithm.
//
// See: https://developer.apple.com/documentation/Security/SecKeyCreateEncryptedData(_:_:_:_:)
func SecKeyCreateEncryptedData(key SecKeyRef, algorithm SecKeyAlgorithm, plaintext corefoundation.CFDataRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef {
	result, callErr := TrySecKeyCreateEncryptedData(key, algorithm, plaintext, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secKeyCreatePair func(keychainRef SecKeychainRef, algorithm CSSM_ALGORITHMS, keySizeInBits uint32, contextHandle CSSM_CC_HANDLE, publicKeyUsage CSSM_KEYUSE, publicKeyAttr uint32, privateKeyUsage CSSM_KEYUSE, privateKeyAttr uint32, initialAccess SecAccessRef, publicKey *SecKeyRef, privateKey *SecKeyRef) int32
var _secKeyCreatePairErr error

// CanCallSecKeyCreatePair reports whether the symbol for SecKeyCreatePair is available on this system.
func CanCallSecKeyCreatePair() bool {
	return _secKeyCreatePair != nil
}

// SecKeyCreatePairCallError returns the symbol lookup or registration error for SecKeyCreatePair.
func SecKeyCreatePairCallError() error {
	if _secKeyCreatePair != nil {
		return nil
	}
	return symbolCallError("SecKeyCreatePair", "10.0", _secKeyCreatePairErr)
}

// TrySecKeyCreatePair calls SecKeyCreatePair without panicking when the symbol is unavailable.
func TrySecKeyCreatePair(keychainRef SecKeychainRef, algorithm CSSM_ALGORITHMS, keySizeInBits uint32, contextHandle CSSM_CC_HANDLE, publicKeyUsage CSSM_KEYUSE, publicKeyAttr uint32, privateKeyUsage CSSM_KEYUSE, privateKeyAttr uint32, initialAccess SecAccessRef, publicKey *SecKeyRef, privateKey *SecKeyRef) (int32, error) {
	if err := SecKeyCreatePairCallError(); err != nil {
		return 0, err
	}
	return _secKeyCreatePair(keychainRef, algorithm, keySizeInBits, contextHandle, publicKeyUsage, publicKeyAttr, privateKeyUsage, privateKeyAttr, initialAccess, publicKey, privateKey), nil
}

// SecKeyCreatePair creates an asymmetric key pair and stores it in a keychain.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecKeyCreatePair
func SecKeyCreatePair(keychainRef SecKeychainRef, algorithm CSSM_ALGORITHMS, keySizeInBits uint32, contextHandle CSSM_CC_HANDLE, publicKeyUsage CSSM_KEYUSE, publicKeyAttr uint32, privateKeyUsage CSSM_KEYUSE, privateKeyAttr uint32, initialAccess SecAccessRef, publicKey *SecKeyRef, privateKey *SecKeyRef) int32 {
	result, callErr := TrySecKeyCreatePair(keychainRef, algorithm, keySizeInBits, contextHandle, publicKeyUsage, publicKeyAttr, privateKeyUsage, privateKeyAttr, initialAccess, publicKey, privateKey)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secKeyCreateRandomKey func(parameters corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) SecKeyRef
var _secKeyCreateRandomKeyErr error

// CanCallSecKeyCreateRandomKey reports whether the symbol for SecKeyCreateRandomKey is available on this system.
func CanCallSecKeyCreateRandomKey() bool {
	return _secKeyCreateRandomKey != nil
}

// SecKeyCreateRandomKeyCallError returns the symbol lookup or registration error for SecKeyCreateRandomKey.
func SecKeyCreateRandomKeyCallError() error {
	if _secKeyCreateRandomKey != nil {
		return nil
	}
	return symbolCallError("SecKeyCreateRandomKey", "10.12", _secKeyCreateRandomKeyErr)
}

// TrySecKeyCreateRandomKey calls SecKeyCreateRandomKey without panicking when the symbol is unavailable.
func TrySecKeyCreateRandomKey(parameters corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) (SecKeyRef, error) {
	if err := SecKeyCreateRandomKeyCallError(); err != nil {
		return 0, err
	}
	return _secKeyCreateRandomKey(parameters, err), nil
}

// SecKeyCreateRandomKey generates a new public-private key pair.
//
// See: https://developer.apple.com/documentation/Security/SecKeyCreateRandomKey(_:_:)
func SecKeyCreateRandomKey(parameters corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) SecKeyRef {
	result, callErr := TrySecKeyCreateRandomKey(parameters, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secKeyCreateSignature func(key SecKeyRef, algorithm SecKeyAlgorithm, dataToSign corefoundation.CFDataRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef
var _secKeyCreateSignatureErr error

// CanCallSecKeyCreateSignature reports whether the symbol for SecKeyCreateSignature is available on this system.
func CanCallSecKeyCreateSignature() bool {
	return _secKeyCreateSignature != nil
}

// SecKeyCreateSignatureCallError returns the symbol lookup or registration error for SecKeyCreateSignature.
func SecKeyCreateSignatureCallError() error {
	if _secKeyCreateSignature != nil {
		return nil
	}
	return symbolCallError("SecKeyCreateSignature", "10.12", _secKeyCreateSignatureErr)
}

// TrySecKeyCreateSignature calls SecKeyCreateSignature without panicking when the symbol is unavailable.
func TrySecKeyCreateSignature(key SecKeyRef, algorithm SecKeyAlgorithm, dataToSign corefoundation.CFDataRef, err *corefoundation.CFErrorRef) (corefoundation.CFDataRef, error) {
	if err := SecKeyCreateSignatureCallError(); err != nil {
		return 0, err
	}
	return _secKeyCreateSignature(key, algorithm, dataToSign, err), nil
}

// SecKeyCreateSignature creates the cryptographic signature for a block of data using a private key and specified algorithm.
//
// See: https://developer.apple.com/documentation/Security/SecKeyCreateSignature(_:_:_:_:)
func SecKeyCreateSignature(key SecKeyRef, algorithm SecKeyAlgorithm, dataToSign corefoundation.CFDataRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef {
	result, callErr := TrySecKeyCreateSignature(key, algorithm, dataToSign, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secKeyCreateWithData func(keyData corefoundation.CFDataRef, attributes corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) SecKeyRef
var _secKeyCreateWithDataErr error

// CanCallSecKeyCreateWithData reports whether the symbol for SecKeyCreateWithData is available on this system.
func CanCallSecKeyCreateWithData() bool {
	return _secKeyCreateWithData != nil
}

// SecKeyCreateWithDataCallError returns the symbol lookup or registration error for SecKeyCreateWithData.
func SecKeyCreateWithDataCallError() error {
	if _secKeyCreateWithData != nil {
		return nil
	}
	return symbolCallError("SecKeyCreateWithData", "10.12", _secKeyCreateWithDataErr)
}

// TrySecKeyCreateWithData calls SecKeyCreateWithData without panicking when the symbol is unavailable.
func TrySecKeyCreateWithData(keyData corefoundation.CFDataRef, attributes corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) (SecKeyRef, error) {
	if err := SecKeyCreateWithDataCallError(); err != nil {
		return 0, err
	}
	return _secKeyCreateWithData(keyData, attributes, err), nil
}

// SecKeyCreateWithData restores a key from an external representation of that key.
//
// See: https://developer.apple.com/documentation/Security/SecKeyCreateWithData(_:_:_:)
func SecKeyCreateWithData(keyData corefoundation.CFDataRef, attributes corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) SecKeyRef {
	result, callErr := TrySecKeyCreateWithData(keyData, attributes, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secKeyGenerate func(keychainRef SecKeychainRef, algorithm CSSM_ALGORITHMS, keySizeInBits uint32, contextHandle CSSM_CC_HANDLE, keyUsage CSSM_KEYUSE, keyAttr uint32, initialAccess SecAccessRef, keyRef *SecKeyRef) int32
var _secKeyGenerateErr error

// CanCallSecKeyGenerate reports whether the symbol for SecKeyGenerate is available on this system.
func CanCallSecKeyGenerate() bool {
	return _secKeyGenerate != nil
}

// SecKeyGenerateCallError returns the symbol lookup or registration error for SecKeyGenerate.
func SecKeyGenerateCallError() error {
	if _secKeyGenerate != nil {
		return nil
	}
	return symbolCallError("SecKeyGenerate", "10.0", _secKeyGenerateErr)
}

// TrySecKeyGenerate calls SecKeyGenerate without panicking when the symbol is unavailable.
func TrySecKeyGenerate(keychainRef SecKeychainRef, algorithm CSSM_ALGORITHMS, keySizeInBits uint32, contextHandle CSSM_CC_HANDLE, keyUsage CSSM_KEYUSE, keyAttr uint32, initialAccess SecAccessRef, keyRef *SecKeyRef) (int32, error) {
	if err := SecKeyGenerateCallError(); err != nil {
		return 0, err
	}
	return _secKeyGenerate(keychainRef, algorithm, keySizeInBits, contextHandle, keyUsage, keyAttr, initialAccess, keyRef), nil
}

// SecKeyGenerate creates a symmetric key and optionally stores it in a keychain.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecKeyGenerate
func SecKeyGenerate(keychainRef SecKeychainRef, algorithm CSSM_ALGORITHMS, keySizeInBits uint32, contextHandle CSSM_CC_HANDLE, keyUsage CSSM_KEYUSE, keyAttr uint32, initialAccess SecAccessRef, keyRef *SecKeyRef) int32 {
	result, callErr := TrySecKeyGenerate(keychainRef, algorithm, keySizeInBits, contextHandle, keyUsage, keyAttr, initialAccess, keyRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secKeyGetBlockSize func(key SecKeyRef) uintptr
var _secKeyGetBlockSizeErr error

// CanCallSecKeyGetBlockSize reports whether the symbol for SecKeyGetBlockSize is available on this system.
func CanCallSecKeyGetBlockSize() bool {
	return _secKeyGetBlockSize != nil
}

// SecKeyGetBlockSizeCallError returns the symbol lookup or registration error for SecKeyGetBlockSize.
func SecKeyGetBlockSizeCallError() error {
	if _secKeyGetBlockSize != nil {
		return nil
	}
	return symbolCallError("SecKeyGetBlockSize", "10.6", _secKeyGetBlockSizeErr)
}

// TrySecKeyGetBlockSize calls SecKeyGetBlockSize without panicking when the symbol is unavailable.
func TrySecKeyGetBlockSize(key SecKeyRef) (uintptr, error) {
	if err := SecKeyGetBlockSizeCallError(); err != nil {
		return 0, err
	}
	return _secKeyGetBlockSize(key), nil
}

// SecKeyGetBlockSize gets the block length associated with a cryptographic key.
//
// See: https://developer.apple.com/documentation/Security/SecKeyGetBlockSize(_:)
func SecKeyGetBlockSize(key SecKeyRef) uintptr {
	result, callErr := TrySecKeyGetBlockSize(key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secKeyGetCSPHandle func(keyRef SecKeyRef, cspHandle unsafe.Pointer) int32
var _secKeyGetCSPHandleErr error

// CanCallSecKeyGetCSPHandle reports whether the symbol for SecKeyGetCSPHandle is available on this system.
func CanCallSecKeyGetCSPHandle() bool {
	return _secKeyGetCSPHandle != nil
}

// SecKeyGetCSPHandleCallError returns the symbol lookup or registration error for SecKeyGetCSPHandle.
func SecKeyGetCSPHandleCallError() error {
	if _secKeyGetCSPHandle != nil {
		return nil
	}
	return symbolCallError("SecKeyGetCSPHandle", "10.0", _secKeyGetCSPHandleErr)
}

// TrySecKeyGetCSPHandle calls SecKeyGetCSPHandle without panicking when the symbol is unavailable.
func TrySecKeyGetCSPHandle(keyRef SecKeyRef, cspHandle unsafe.Pointer) (int32, error) {
	if err := SecKeyGetCSPHandleCallError(); err != nil {
		return 0, err
	}
	return _secKeyGetCSPHandle(keyRef, cspHandle), nil
}

// SecKeyGetCSPHandle returns the CSSM CSP handle for a key.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecKeyGetCSPHandle
func SecKeyGetCSPHandle(keyRef SecKeyRef, cspHandle unsafe.Pointer) int32 {
	result, callErr := TrySecKeyGetCSPHandle(keyRef, cspHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secKeyGetCSSMKey func(key SecKeyRef, cssmKey unsafe.Pointer) int32
var _secKeyGetCSSMKeyErr error

// CanCallSecKeyGetCSSMKey reports whether the symbol for SecKeyGetCSSMKey is available on this system.
func CanCallSecKeyGetCSSMKey() bool {
	return _secKeyGetCSSMKey != nil
}

// SecKeyGetCSSMKeyCallError returns the symbol lookup or registration error for SecKeyGetCSSMKey.
func SecKeyGetCSSMKeyCallError() error {
	if _secKeyGetCSSMKey != nil {
		return nil
	}
	return symbolCallError("SecKeyGetCSSMKey", "10.0", _secKeyGetCSSMKeyErr)
}

// TrySecKeyGetCSSMKey calls SecKeyGetCSSMKey without panicking when the symbol is unavailable.
func TrySecKeyGetCSSMKey(key SecKeyRef, cssmKey unsafe.Pointer) (int32, error) {
	if err := SecKeyGetCSSMKeyCallError(); err != nil {
		return 0, err
	}
	return _secKeyGetCSSMKey(key, cssmKey), nil
}

// SecKeyGetCSSMKey retrieves a pointer to the `CSSM_KEY` structure containing the key stored in a keychain item.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecKeyGetCSSMKey
func SecKeyGetCSSMKey(key SecKeyRef, cssmKey unsafe.Pointer) int32 {
	result, callErr := TrySecKeyGetCSSMKey(key, cssmKey)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secKeyGetCredentials func(keyRef SecKeyRef, operation CSSM_ACL_AUTHORIZATION_TAG, credentialType SecCredentialType, outCredentials unsafe.Pointer) int32
var _secKeyGetCredentialsErr error

// CanCallSecKeyGetCredentials reports whether the symbol for SecKeyGetCredentials is available on this system.
func CanCallSecKeyGetCredentials() bool {
	return _secKeyGetCredentials != nil
}

// SecKeyGetCredentialsCallError returns the symbol lookup or registration error for SecKeyGetCredentials.
func SecKeyGetCredentialsCallError() error {
	if _secKeyGetCredentials != nil {
		return nil
	}
	return symbolCallError("SecKeyGetCredentials", "10.0", _secKeyGetCredentialsErr)
}

// TrySecKeyGetCredentials calls SecKeyGetCredentials without panicking when the symbol is unavailable.
func TrySecKeyGetCredentials(keyRef SecKeyRef, operation CSSM_ACL_AUTHORIZATION_TAG, credentialType SecCredentialType, outCredentials unsafe.Pointer) (int32, error) {
	if err := SecKeyGetCredentialsCallError(); err != nil {
		return 0, err
	}
	return _secKeyGetCredentials(keyRef, operation, credentialType, outCredentials), nil
}

// SecKeyGetCredentials returns an access credential for a key.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecKeyGetCredentials
func SecKeyGetCredentials(keyRef SecKeyRef, operation CSSM_ACL_AUTHORIZATION_TAG, credentialType SecCredentialType, outCredentials unsafe.Pointer) int32 {
	result, callErr := TrySecKeyGetCredentials(keyRef, operation, credentialType, outCredentials)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secKeyGetTypeID func() uint
var _secKeyGetTypeIDErr error

// CanCallSecKeyGetTypeID reports whether the symbol for SecKeyGetTypeID is available on this system.
func CanCallSecKeyGetTypeID() bool {
	return _secKeyGetTypeID != nil
}

// SecKeyGetTypeIDCallError returns the symbol lookup or registration error for SecKeyGetTypeID.
func SecKeyGetTypeIDCallError() error {
	if _secKeyGetTypeID != nil {
		return nil
	}
	return symbolCallError("SecKeyGetTypeID", "10.3", _secKeyGetTypeIDErr)
}

// TrySecKeyGetTypeID calls SecKeyGetTypeID without panicking when the symbol is unavailable.
func TrySecKeyGetTypeID() (uint, error) {
	if err := SecKeyGetTypeIDCallError(); err != nil {
		return 0, err
	}
	return _secKeyGetTypeID(), nil
}

// SecKeyGetTypeID returns the unique identifier of the opaque type to which a key object belongs.
//
// See: https://developer.apple.com/documentation/Security/SecKeyGetTypeID()
func SecKeyGetTypeID() uint {
	result, callErr := TrySecKeyGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secKeyIsAlgorithmSupported func(key SecKeyRef, operation SecKeyOperationType, algorithm SecKeyAlgorithm) bool
var _secKeyIsAlgorithmSupportedErr error

// CanCallSecKeyIsAlgorithmSupported reports whether the symbol for SecKeyIsAlgorithmSupported is available on this system.
func CanCallSecKeyIsAlgorithmSupported() bool {
	return _secKeyIsAlgorithmSupported != nil
}

// SecKeyIsAlgorithmSupportedCallError returns the symbol lookup or registration error for SecKeyIsAlgorithmSupported.
func SecKeyIsAlgorithmSupportedCallError() error {
	if _secKeyIsAlgorithmSupported != nil {
		return nil
	}
	return symbolCallError("SecKeyIsAlgorithmSupported", "10.12", _secKeyIsAlgorithmSupportedErr)
}

// TrySecKeyIsAlgorithmSupported calls SecKeyIsAlgorithmSupported without panicking when the symbol is unavailable.
func TrySecKeyIsAlgorithmSupported(key SecKeyRef, operation SecKeyOperationType, algorithm SecKeyAlgorithm) (bool, error) {
	if err := SecKeyIsAlgorithmSupportedCallError(); err != nil {
		return false, err
	}
	return _secKeyIsAlgorithmSupported(key, operation, algorithm), nil
}

// SecKeyIsAlgorithmSupported returns a Boolean indicating whether a key is suitable for an operation using a certain algorithm.
//
// See: https://developer.apple.com/documentation/Security/SecKeyIsAlgorithmSupported(_:_:_:)
func SecKeyIsAlgorithmSupported(key SecKeyRef, operation SecKeyOperationType, algorithm SecKeyAlgorithm) bool {
	result, callErr := TrySecKeyIsAlgorithmSupported(key, operation, algorithm)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secKeyVerifySignature func(key SecKeyRef, algorithm SecKeyAlgorithm, signedData corefoundation.CFDataRef, signature corefoundation.CFDataRef, err *corefoundation.CFErrorRef) bool
var _secKeyVerifySignatureErr error

// CanCallSecKeyVerifySignature reports whether the symbol for SecKeyVerifySignature is available on this system.
func CanCallSecKeyVerifySignature() bool {
	return _secKeyVerifySignature != nil
}

// SecKeyVerifySignatureCallError returns the symbol lookup or registration error for SecKeyVerifySignature.
func SecKeyVerifySignatureCallError() error {
	if _secKeyVerifySignature != nil {
		return nil
	}
	return symbolCallError("SecKeyVerifySignature", "10.12", _secKeyVerifySignatureErr)
}

// TrySecKeyVerifySignature calls SecKeyVerifySignature without panicking when the symbol is unavailable.
func TrySecKeyVerifySignature(key SecKeyRef, algorithm SecKeyAlgorithm, signedData corefoundation.CFDataRef, signature corefoundation.CFDataRef, err *corefoundation.CFErrorRef) (bool, error) {
	if err := SecKeyVerifySignatureCallError(); err != nil {
		return false, err
	}
	return _secKeyVerifySignature(key, algorithm, signedData, signature, err), nil
}

// SecKeyVerifySignature verifies the cryptographic signature of a block of data using a public key and specified algorithm.
//
// See: https://developer.apple.com/documentation/Security/SecKeyVerifySignature(_:_:_:_:_:)
func SecKeyVerifySignature(key SecKeyRef, algorithm SecKeyAlgorithm, signedData corefoundation.CFDataRef, signature corefoundation.CFDataRef, err *corefoundation.CFErrorRef) bool {
	result, callErr := TrySecKeyVerifySignature(key, algorithm, signedData, signature, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secKeychainSearchGetTypeID func() uint
var _secKeychainSearchGetTypeIDErr error

// CanCallSecKeychainSearchGetTypeID reports whether the symbol for SecKeychainSearchGetTypeID is available on this system.
func CanCallSecKeychainSearchGetTypeID() bool {
	return _secKeychainSearchGetTypeID != nil
}

// SecKeychainSearchGetTypeIDCallError returns the symbol lookup or registration error for SecKeychainSearchGetTypeID.
func SecKeychainSearchGetTypeIDCallError() error {
	if _secKeychainSearchGetTypeID != nil {
		return nil
	}
	return symbolCallError("SecKeychainSearchGetTypeID", "10.0", _secKeychainSearchGetTypeIDErr)
}

// TrySecKeychainSearchGetTypeID calls SecKeychainSearchGetTypeID without panicking when the symbol is unavailable.
func TrySecKeychainSearchGetTypeID() (uint, error) {
	if err := SecKeychainSearchGetTypeIDCallError(); err != nil {
		return 0, err
	}
	return _secKeychainSearchGetTypeID(), nil
}

// SecKeychainSearchGetTypeID returns the unique identifier of the opaque type to which a keychain search object belongs.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecKeychainSearchGetTypeID
func SecKeychainSearchGetTypeID() uint {
	result, callErr := TrySecKeychainSearchGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secPKCS12Import func(pkcs12_data corefoundation.CFDataRef, options corefoundation.CFDictionaryRef, items *corefoundation.CFArrayRef) int32
var _secPKCS12ImportErr error

// CanCallSecPKCS12Import reports whether the symbol for SecPKCS12Import is available on this system.
func CanCallSecPKCS12Import() bool {
	return _secPKCS12Import != nil
}

// SecPKCS12ImportCallError returns the symbol lookup or registration error for SecPKCS12Import.
func SecPKCS12ImportCallError() error {
	if _secPKCS12Import != nil {
		return nil
	}
	return symbolCallError("SecPKCS12Import", "10.6", _secPKCS12ImportErr)
}

// TrySecPKCS12Import calls SecPKCS12Import without panicking when the symbol is unavailable.
func TrySecPKCS12Import(pkcs12_data corefoundation.CFDataRef, options corefoundation.CFDictionaryRef, items *corefoundation.CFArrayRef) (int32, error) {
	if err := SecPKCS12ImportCallError(); err != nil {
		return 0, err
	}
	return _secPKCS12Import(pkcs12_data, options, items), nil
}

// SecPKCS12Import returns the identities and certificates in a PKCS #12-formatted blob.
//
// See: https://developer.apple.com/documentation/Security/SecPKCS12Import(_:_:_:)
func SecPKCS12Import(pkcs12_data corefoundation.CFDataRef, options corefoundation.CFDictionaryRef, items *corefoundation.CFArrayRef) int32 {
	result, callErr := TrySecPKCS12Import(pkcs12_data, options, items)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secPolicyCopyProperties func(policyRef SecPolicyRef) corefoundation.CFDictionaryRef
var _secPolicyCopyPropertiesErr error

// CanCallSecPolicyCopyProperties reports whether the symbol for SecPolicyCopyProperties is available on this system.
func CanCallSecPolicyCopyProperties() bool {
	return _secPolicyCopyProperties != nil
}

// SecPolicyCopyPropertiesCallError returns the symbol lookup or registration error for SecPolicyCopyProperties.
func SecPolicyCopyPropertiesCallError() error {
	if _secPolicyCopyProperties != nil {
		return nil
	}
	return symbolCallError("SecPolicyCopyProperties", "10.7", _secPolicyCopyPropertiesErr)
}

// TrySecPolicyCopyProperties calls SecPolicyCopyProperties without panicking when the symbol is unavailable.
func TrySecPolicyCopyProperties(policyRef SecPolicyRef) (corefoundation.CFDictionaryRef, error) {
	if err := SecPolicyCopyPropertiesCallError(); err != nil {
		return 0, err
	}
	return _secPolicyCopyProperties(policyRef), nil
}

// SecPolicyCopyProperties returns a dictionary containing a policy’s properties.
//
// See: https://developer.apple.com/documentation/Security/SecPolicyCopyProperties(_:)
func SecPolicyCopyProperties(policyRef SecPolicyRef) corefoundation.CFDictionaryRef {
	result, callErr := TrySecPolicyCopyProperties(policyRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secPolicyCreateBasicX509 func() SecPolicyRef
var _secPolicyCreateBasicX509Err error

// CanCallSecPolicyCreateBasicX509 reports whether the symbol for SecPolicyCreateBasicX509 is available on this system.
func CanCallSecPolicyCreateBasicX509() bool {
	return _secPolicyCreateBasicX509 != nil
}

// SecPolicyCreateBasicX509CallError returns the symbol lookup or registration error for SecPolicyCreateBasicX509.
func SecPolicyCreateBasicX509CallError() error {
	if _secPolicyCreateBasicX509 != nil {
		return nil
	}
	return symbolCallError("SecPolicyCreateBasicX509", "10.6", _secPolicyCreateBasicX509Err)
}

// TrySecPolicyCreateBasicX509 calls SecPolicyCreateBasicX509 without panicking when the symbol is unavailable.
func TrySecPolicyCreateBasicX509() (SecPolicyRef, error) {
	if err := SecPolicyCreateBasicX509CallError(); err != nil {
		return 0, err
	}
	return _secPolicyCreateBasicX509(), nil
}

// SecPolicyCreateBasicX509 returns a policy object for the default X.509 policy.
//
// See: https://developer.apple.com/documentation/Security/SecPolicyCreateBasicX509()
func SecPolicyCreateBasicX509() SecPolicyRef {
	result, callErr := TrySecPolicyCreateBasicX509()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secPolicyCreateRevocation func(revocationFlags uint64) SecPolicyRef
var _secPolicyCreateRevocationErr error

// CanCallSecPolicyCreateRevocation reports whether the symbol for SecPolicyCreateRevocation is available on this system.
func CanCallSecPolicyCreateRevocation() bool {
	return _secPolicyCreateRevocation != nil
}

// SecPolicyCreateRevocationCallError returns the symbol lookup or registration error for SecPolicyCreateRevocation.
func SecPolicyCreateRevocationCallError() error {
	if _secPolicyCreateRevocation != nil {
		return nil
	}
	return symbolCallError("SecPolicyCreateRevocation", "10.9", _secPolicyCreateRevocationErr)
}

// TrySecPolicyCreateRevocation calls SecPolicyCreateRevocation without panicking when the symbol is unavailable.
func TrySecPolicyCreateRevocation(revocationFlags uint64) (SecPolicyRef, error) {
	if err := SecPolicyCreateRevocationCallError(); err != nil {
		return 0, err
	}
	return _secPolicyCreateRevocation(revocationFlags), nil
}

// SecPolicyCreateRevocation returns a policy object for checking revocation of certificates.
//
// See: https://developer.apple.com/documentation/Security/SecPolicyCreateRevocation(_:)
func SecPolicyCreateRevocation(revocationFlags uint64) SecPolicyRef {
	result, callErr := TrySecPolicyCreateRevocation(revocationFlags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secPolicyCreateSSL func(server bool, hostname corefoundation.CFStringRef) SecPolicyRef
var _secPolicyCreateSSLErr error

// CanCallSecPolicyCreateSSL reports whether the symbol for SecPolicyCreateSSL is available on this system.
func CanCallSecPolicyCreateSSL() bool {
	return _secPolicyCreateSSL != nil
}

// SecPolicyCreateSSLCallError returns the symbol lookup or registration error for SecPolicyCreateSSL.
func SecPolicyCreateSSLCallError() error {
	if _secPolicyCreateSSL != nil {
		return nil
	}
	return symbolCallError("SecPolicyCreateSSL", "10.6", _secPolicyCreateSSLErr)
}

// TrySecPolicyCreateSSL calls SecPolicyCreateSSL without panicking when the symbol is unavailable.
func TrySecPolicyCreateSSL(server bool, hostname corefoundation.CFStringRef) (SecPolicyRef, error) {
	if err := SecPolicyCreateSSLCallError(); err != nil {
		return 0, err
	}
	return _secPolicyCreateSSL(server, hostname), nil
}

// SecPolicyCreateSSL returns a policy object for evaluating SSL certificate chains.
//
// See: https://developer.apple.com/documentation/Security/SecPolicyCreateSSL(_:_:)
func SecPolicyCreateSSL(server bool, hostname corefoundation.CFStringRef) SecPolicyRef {
	result, callErr := TrySecPolicyCreateSSL(server, hostname)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secPolicyCreateWithOID func(policyOID corefoundation.CFTypeRef) SecPolicyRef
var _secPolicyCreateWithOIDErr error

// CanCallSecPolicyCreateWithOID reports whether the symbol for SecPolicyCreateWithOID is available on this system.
func CanCallSecPolicyCreateWithOID() bool {
	return _secPolicyCreateWithOID != nil
}

// SecPolicyCreateWithOIDCallError returns the symbol lookup or registration error for SecPolicyCreateWithOID.
func SecPolicyCreateWithOIDCallError() error {
	if _secPolicyCreateWithOID != nil {
		return nil
	}
	return symbolCallError("SecPolicyCreateWithOID", "10.7", _secPolicyCreateWithOIDErr)
}

// TrySecPolicyCreateWithOID calls SecPolicyCreateWithOID without panicking when the symbol is unavailable.
func TrySecPolicyCreateWithOID(policyOID corefoundation.CFTypeRef) (SecPolicyRef, error) {
	if err := SecPolicyCreateWithOIDCallError(); err != nil {
		return 0, err
	}
	return _secPolicyCreateWithOID(policyOID), nil
}

// SecPolicyCreateWithOID returns a policy object for the specified policy type object identifier.
//
// Deprecated: Deprecated since macOS 10.9.
//
// See: https://developer.apple.com/documentation/Security/SecPolicyCreateWithOID
func SecPolicyCreateWithOID(policyOID corefoundation.CFTypeRef) SecPolicyRef {
	result, callErr := TrySecPolicyCreateWithOID(policyOID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secPolicyCreateWithProperties func(policyIdentifier corefoundation.CFTypeRef, properties corefoundation.CFDictionaryRef) SecPolicyRef
var _secPolicyCreateWithPropertiesErr error

// CanCallSecPolicyCreateWithProperties reports whether the symbol for SecPolicyCreateWithProperties is available on this system.
func CanCallSecPolicyCreateWithProperties() bool {
	return _secPolicyCreateWithProperties != nil
}

// SecPolicyCreateWithPropertiesCallError returns the symbol lookup or registration error for SecPolicyCreateWithProperties.
func SecPolicyCreateWithPropertiesCallError() error {
	if _secPolicyCreateWithProperties != nil {
		return nil
	}
	return symbolCallError("SecPolicyCreateWithProperties", "10.9", _secPolicyCreateWithPropertiesErr)
}

// TrySecPolicyCreateWithProperties calls SecPolicyCreateWithProperties without panicking when the symbol is unavailable.
func TrySecPolicyCreateWithProperties(policyIdentifier corefoundation.CFTypeRef, properties corefoundation.CFDictionaryRef) (SecPolicyRef, error) {
	if err := SecPolicyCreateWithPropertiesCallError(); err != nil {
		return 0, err
	}
	return _secPolicyCreateWithProperties(policyIdentifier, properties), nil
}

// SecPolicyCreateWithProperties returns a policy object based on an object identifier for the policy type.
//
// See: https://developer.apple.com/documentation/Security/SecPolicyCreateWithProperties(_:_:)
func SecPolicyCreateWithProperties(policyIdentifier corefoundation.CFTypeRef, properties corefoundation.CFDictionaryRef) SecPolicyRef {
	result, callErr := TrySecPolicyCreateWithProperties(policyIdentifier, properties)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secPolicyGetOID func(policyRef SecPolicyRef, oid unsafe.Pointer) int32
var _secPolicyGetOIDErr error

// CanCallSecPolicyGetOID reports whether the symbol for SecPolicyGetOID is available on this system.
func CanCallSecPolicyGetOID() bool {
	return _secPolicyGetOID != nil
}

// SecPolicyGetOIDCallError returns the symbol lookup or registration error for SecPolicyGetOID.
func SecPolicyGetOIDCallError() error {
	if _secPolicyGetOID != nil {
		return nil
	}
	return symbolCallError("SecPolicyGetOID", "10.2", _secPolicyGetOIDErr)
}

// TrySecPolicyGetOID calls SecPolicyGetOID without panicking when the symbol is unavailable.
func TrySecPolicyGetOID(policyRef SecPolicyRef, oid unsafe.Pointer) (int32, error) {
	if err := SecPolicyGetOIDCallError(); err != nil {
		return 0, err
	}
	return _secPolicyGetOID(policyRef, oid), nil
}

// SecPolicyGetOID retrieves a policy’s object identifier.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecPolicyGetOID
func SecPolicyGetOID(policyRef SecPolicyRef, oid unsafe.Pointer) int32 {
	result, callErr := TrySecPolicyGetOID(policyRef, oid)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secPolicyGetTPHandle func(policyRef SecPolicyRef, tpHandle unsafe.Pointer) int32
var _secPolicyGetTPHandleErr error

// CanCallSecPolicyGetTPHandle reports whether the symbol for SecPolicyGetTPHandle is available on this system.
func CanCallSecPolicyGetTPHandle() bool {
	return _secPolicyGetTPHandle != nil
}

// SecPolicyGetTPHandleCallError returns the symbol lookup or registration error for SecPolicyGetTPHandle.
func SecPolicyGetTPHandleCallError() error {
	if _secPolicyGetTPHandle != nil {
		return nil
	}
	return symbolCallError("SecPolicyGetTPHandle", "10.2", _secPolicyGetTPHandleErr)
}

// TrySecPolicyGetTPHandle calls SecPolicyGetTPHandle without panicking when the symbol is unavailable.
func TrySecPolicyGetTPHandle(policyRef SecPolicyRef, tpHandle unsafe.Pointer) (int32, error) {
	if err := SecPolicyGetTPHandleCallError(); err != nil {
		return 0, err
	}
	return _secPolicyGetTPHandle(policyRef, tpHandle), nil
}

// SecPolicyGetTPHandle retrieves the trust policy handle for a policy object.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecPolicyGetTPHandle
func SecPolicyGetTPHandle(policyRef SecPolicyRef, tpHandle unsafe.Pointer) int32 {
	result, callErr := TrySecPolicyGetTPHandle(policyRef, tpHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secPolicyGetTypeID func() uint
var _secPolicyGetTypeIDErr error

// CanCallSecPolicyGetTypeID reports whether the symbol for SecPolicyGetTypeID is available on this system.
func CanCallSecPolicyGetTypeID() bool {
	return _secPolicyGetTypeID != nil
}

// SecPolicyGetTypeIDCallError returns the symbol lookup or registration error for SecPolicyGetTypeID.
func SecPolicyGetTypeIDCallError() error {
	if _secPolicyGetTypeID != nil {
		return nil
	}
	return symbolCallError("SecPolicyGetTypeID", "10.3", _secPolicyGetTypeIDErr)
}

// TrySecPolicyGetTypeID calls SecPolicyGetTypeID without panicking when the symbol is unavailable.
func TrySecPolicyGetTypeID() (uint, error) {
	if err := SecPolicyGetTypeIDCallError(); err != nil {
		return 0, err
	}
	return _secPolicyGetTypeID(), nil
}

// SecPolicyGetTypeID returns the unique identifier of the opaque type to which a policy object belongs.
//
// See: https://developer.apple.com/documentation/Security/SecPolicyGetTypeID()
func SecPolicyGetTypeID() uint {
	result, callErr := TrySecPolicyGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secPolicyGetValue func(policyRef SecPolicyRef, value unsafe.Pointer) int32
var _secPolicyGetValueErr error

// CanCallSecPolicyGetValue reports whether the symbol for SecPolicyGetValue is available on this system.
func CanCallSecPolicyGetValue() bool {
	return _secPolicyGetValue != nil
}

// SecPolicyGetValueCallError returns the symbol lookup or registration error for SecPolicyGetValue.
func SecPolicyGetValueCallError() error {
	if _secPolicyGetValue != nil {
		return nil
	}
	return symbolCallError("SecPolicyGetValue", "10.2", _secPolicyGetValueErr)
}

// TrySecPolicyGetValue calls SecPolicyGetValue without panicking when the symbol is unavailable.
func TrySecPolicyGetValue(policyRef SecPolicyRef, value unsafe.Pointer) (int32, error) {
	if err := SecPolicyGetValueCallError(); err != nil {
		return 0, err
	}
	return _secPolicyGetValue(policyRef, value), nil
}

// SecPolicyGetValue retrieves a policy’s value.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecPolicyGetValue
func SecPolicyGetValue(policyRef SecPolicyRef, value unsafe.Pointer) int32 {
	result, callErr := TrySecPolicyGetValue(policyRef, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secPolicySearchCopyNext func(searchRef SecPolicySearchRef, policyRef *SecPolicyRef) int32
var _secPolicySearchCopyNextErr error

// CanCallSecPolicySearchCopyNext reports whether the symbol for SecPolicySearchCopyNext is available on this system.
func CanCallSecPolicySearchCopyNext() bool {
	return _secPolicySearchCopyNext != nil
}

// SecPolicySearchCopyNextCallError returns the symbol lookup or registration error for SecPolicySearchCopyNext.
func SecPolicySearchCopyNextCallError() error {
	if _secPolicySearchCopyNext != nil {
		return nil
	}
	return symbolCallError("SecPolicySearchCopyNext", "10.0", _secPolicySearchCopyNextErr)
}

// TrySecPolicySearchCopyNext calls SecPolicySearchCopyNext without panicking when the symbol is unavailable.
func TrySecPolicySearchCopyNext(searchRef SecPolicySearchRef, policyRef *SecPolicyRef) (int32, error) {
	if err := SecPolicySearchCopyNextCallError(); err != nil {
		return 0, err
	}
	return _secPolicySearchCopyNext(searchRef, policyRef), nil
}

// SecPolicySearchCopyNext retrieves a policy object for the next policy matching specified search criteria.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecPolicySearchCopyNext
func SecPolicySearchCopyNext(searchRef SecPolicySearchRef, policyRef *SecPolicyRef) int32 {
	result, callErr := TrySecPolicySearchCopyNext(searchRef, policyRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secPolicySearchCreate func(certType CSSM_CERT_TYPE, policyOID unsafe.Pointer, value unsafe.Pointer, searchRef *SecPolicySearchRef) int32
var _secPolicySearchCreateErr error

// CanCallSecPolicySearchCreate reports whether the symbol for SecPolicySearchCreate is available on this system.
func CanCallSecPolicySearchCreate() bool {
	return _secPolicySearchCreate != nil
}

// SecPolicySearchCreateCallError returns the symbol lookup or registration error for SecPolicySearchCreate.
func SecPolicySearchCreateCallError() error {
	if _secPolicySearchCreate != nil {
		return nil
	}
	return symbolCallError("SecPolicySearchCreate", "10.0", _secPolicySearchCreateErr)
}

// TrySecPolicySearchCreate calls SecPolicySearchCreate without panicking when the symbol is unavailable.
func TrySecPolicySearchCreate(certType CSSM_CERT_TYPE, policyOID unsafe.Pointer, value unsafe.Pointer, searchRef *SecPolicySearchRef) (int32, error) {
	if err := SecPolicySearchCreateCallError(); err != nil {
		return 0, err
	}
	return _secPolicySearchCreate(certType, policyOID, value, searchRef), nil
}

// SecPolicySearchCreate creates a search object for finding policies.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecPolicySearchCreate
func SecPolicySearchCreate(certType CSSM_CERT_TYPE, policyOID unsafe.Pointer, value unsafe.Pointer, searchRef *SecPolicySearchRef) int32 {
	result, callErr := TrySecPolicySearchCreate(certType, policyOID, value, searchRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secPolicySearchGetTypeID func() uint
var _secPolicySearchGetTypeIDErr error

// CanCallSecPolicySearchGetTypeID reports whether the symbol for SecPolicySearchGetTypeID is available on this system.
func CanCallSecPolicySearchGetTypeID() bool {
	return _secPolicySearchGetTypeID != nil
}

// SecPolicySearchGetTypeIDCallError returns the symbol lookup or registration error for SecPolicySearchGetTypeID.
func SecPolicySearchGetTypeIDCallError() error {
	if _secPolicySearchGetTypeID != nil {
		return nil
	}
	return symbolCallError("SecPolicySearchGetTypeID", "10.0", _secPolicySearchGetTypeIDErr)
}

// TrySecPolicySearchGetTypeID calls SecPolicySearchGetTypeID without panicking when the symbol is unavailable.
func TrySecPolicySearchGetTypeID() (uint, error) {
	if err := SecPolicySearchGetTypeIDCallError(); err != nil {
		return 0, err
	}
	return _secPolicySearchGetTypeID(), nil
}

// SecPolicySearchGetTypeID returns the unique identifier of the opaque type to which a [SecPolicySearch] object belongs.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecPolicySearchGetTypeID
func SecPolicySearchGetTypeID() uint {
	result, callErr := TrySecPolicySearchGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secPolicySetProperties func(policyRef SecPolicyRef, properties corefoundation.CFDictionaryRef) int32
var _secPolicySetPropertiesErr error

// CanCallSecPolicySetProperties reports whether the symbol for SecPolicySetProperties is available on this system.
func CanCallSecPolicySetProperties() bool {
	return _secPolicySetProperties != nil
}

// SecPolicySetPropertiesCallError returns the symbol lookup or registration error for SecPolicySetProperties.
func SecPolicySetPropertiesCallError() error {
	if _secPolicySetProperties != nil {
		return nil
	}
	return symbolCallError("SecPolicySetProperties", "10.7", _secPolicySetPropertiesErr)
}

// TrySecPolicySetProperties calls SecPolicySetProperties without panicking when the symbol is unavailable.
func TrySecPolicySetProperties(policyRef SecPolicyRef, properties corefoundation.CFDictionaryRef) (int32, error) {
	if err := SecPolicySetPropertiesCallError(); err != nil {
		return 0, err
	}
	return _secPolicySetProperties(policyRef, properties), nil
}

// SecPolicySetProperties sets properties for a policy.
//
// Deprecated: Deprecated since macOS 10.9.
//
// See: https://developer.apple.com/documentation/Security/SecPolicySetProperties
func SecPolicySetProperties(policyRef SecPolicyRef, properties corefoundation.CFDictionaryRef) int32 {
	result, callErr := TrySecPolicySetProperties(policyRef, properties)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secPolicySetValue func(policyRef SecPolicyRef, value unsafe.Pointer) int32
var _secPolicySetValueErr error

// CanCallSecPolicySetValue reports whether the symbol for SecPolicySetValue is available on this system.
func CanCallSecPolicySetValue() bool {
	return _secPolicySetValue != nil
}

// SecPolicySetValueCallError returns the symbol lookup or registration error for SecPolicySetValue.
func SecPolicySetValueCallError() error {
	if _secPolicySetValue != nil {
		return nil
	}
	return symbolCallError("SecPolicySetValue", "10.2", _secPolicySetValueErr)
}

// TrySecPolicySetValue calls SecPolicySetValue without panicking when the symbol is unavailable.
func TrySecPolicySetValue(policyRef SecPolicyRef, value unsafe.Pointer) (int32, error) {
	if err := SecPolicySetValueCallError(); err != nil {
		return 0, err
	}
	return _secPolicySetValue(policyRef, value), nil
}

// SecPolicySetValue sets a policy’s value.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecPolicySetValue
func SecPolicySetValue(policyRef SecPolicyRef, value unsafe.Pointer) int32 {
	result, callErr := TrySecPolicySetValue(policyRef, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secRandomCopyBytes func(rnd SecRandomRef, count uintptr, bytes unsafe.Pointer) int
var _secRandomCopyBytesErr error

// CanCallSecRandomCopyBytes reports whether the symbol for SecRandomCopyBytes is available on this system.
func CanCallSecRandomCopyBytes() bool {
	return _secRandomCopyBytes != nil
}

// SecRandomCopyBytesCallError returns the symbol lookup or registration error for SecRandomCopyBytes.
func SecRandomCopyBytesCallError() error {
	if _secRandomCopyBytes != nil {
		return nil
	}
	return symbolCallError("SecRandomCopyBytes", "10.7", _secRandomCopyBytesErr)
}

// TrySecRandomCopyBytes calls SecRandomCopyBytes without panicking when the symbol is unavailable.
func TrySecRandomCopyBytes(rnd SecRandomRef, count uintptr, bytes unsafe.Pointer) (int, error) {
	if err := SecRandomCopyBytesCallError(); err != nil {
		return 0, err
	}
	return _secRandomCopyBytes(rnd, count, bytes), nil
}

// SecRandomCopyBytes generates an array of cryptographically secure random bytes.
//
// See: https://developer.apple.com/documentation/Security/SecRandomCopyBytes(_:_:_:)
func SecRandomCopyBytes(rnd SecRandomRef, count uintptr, bytes unsafe.Pointer) int {
	result, callErr := TrySecRandomCopyBytes(rnd, count, bytes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secRequirementCopyData func(requirement SecRequirementRef, flags SecCSFlags, data *corefoundation.CFDataRef) int32
var _secRequirementCopyDataErr error

// CanCallSecRequirementCopyData reports whether the symbol for SecRequirementCopyData is available on this system.
func CanCallSecRequirementCopyData() bool {
	return _secRequirementCopyData != nil
}

// SecRequirementCopyDataCallError returns the symbol lookup or registration error for SecRequirementCopyData.
func SecRequirementCopyDataCallError() error {
	if _secRequirementCopyData != nil {
		return nil
	}
	return symbolCallError("SecRequirementCopyData", "10.0", _secRequirementCopyDataErr)
}

// TrySecRequirementCopyData calls SecRequirementCopyData without panicking when the symbol is unavailable.
func TrySecRequirementCopyData(requirement SecRequirementRef, flags SecCSFlags, data *corefoundation.CFDataRef) (int32, error) {
	if err := SecRequirementCopyDataCallError(); err != nil {
		return 0, err
	}
	return _secRequirementCopyData(requirement, flags, data), nil
}

// SecRequirementCopyData extracts a binary form of a code requirement from a code requirement object.
//
// See: https://developer.apple.com/documentation/Security/SecRequirementCopyData(_:_:_:)
func SecRequirementCopyData(requirement SecRequirementRef, flags SecCSFlags, data *corefoundation.CFDataRef) int32 {
	result, callErr := TrySecRequirementCopyData(requirement, flags, data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secRequirementCopyString func(requirement SecRequirementRef, flags SecCSFlags, text *corefoundation.CFStringRef) int32
var _secRequirementCopyStringErr error

// CanCallSecRequirementCopyString reports whether the symbol for SecRequirementCopyString is available on this system.
func CanCallSecRequirementCopyString() bool {
	return _secRequirementCopyString != nil
}

// SecRequirementCopyStringCallError returns the symbol lookup or registration error for SecRequirementCopyString.
func SecRequirementCopyStringCallError() error {
	if _secRequirementCopyString != nil {
		return nil
	}
	return symbolCallError("SecRequirementCopyString", "10.0", _secRequirementCopyStringErr)
}

// TrySecRequirementCopyString calls SecRequirementCopyString without panicking when the symbol is unavailable.
func TrySecRequirementCopyString(requirement SecRequirementRef, flags SecCSFlags, text *corefoundation.CFStringRef) (int32, error) {
	if err := SecRequirementCopyStringCallError(); err != nil {
		return 0, err
	}
	return _secRequirementCopyString(requirement, flags, text), nil
}

// SecRequirementCopyString converts a code requirement object into text form.
//
// See: https://developer.apple.com/documentation/Security/SecRequirementCopyString(_:_:_:)
func SecRequirementCopyString(requirement SecRequirementRef, flags SecCSFlags, text *corefoundation.CFStringRef) int32 {
	result, callErr := TrySecRequirementCopyString(requirement, flags, text)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secRequirementCreateWithData func(data corefoundation.CFDataRef, flags SecCSFlags, requirement *SecRequirementRef) int32
var _secRequirementCreateWithDataErr error

// CanCallSecRequirementCreateWithData reports whether the symbol for SecRequirementCreateWithData is available on this system.
func CanCallSecRequirementCreateWithData() bool {
	return _secRequirementCreateWithData != nil
}

// SecRequirementCreateWithDataCallError returns the symbol lookup or registration error for SecRequirementCreateWithData.
func SecRequirementCreateWithDataCallError() error {
	if _secRequirementCreateWithData != nil {
		return nil
	}
	return symbolCallError("SecRequirementCreateWithData", "10.0", _secRequirementCreateWithDataErr)
}

// TrySecRequirementCreateWithData calls SecRequirementCreateWithData without panicking when the symbol is unavailable.
func TrySecRequirementCreateWithData(data corefoundation.CFDataRef, flags SecCSFlags, requirement *SecRequirementRef) (int32, error) {
	if err := SecRequirementCreateWithDataCallError(); err != nil {
		return 0, err
	}
	return _secRequirementCreateWithData(data, flags, requirement), nil
}

// SecRequirementCreateWithData creates a code requirement object from the binary form of a code requirement.
//
// See: https://developer.apple.com/documentation/Security/SecRequirementCreateWithData(_:_:_:)
func SecRequirementCreateWithData(data corefoundation.CFDataRef, flags SecCSFlags, requirement *SecRequirementRef) int32 {
	result, callErr := TrySecRequirementCreateWithData(data, flags, requirement)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secRequirementCreateWithString func(text corefoundation.CFStringRef, flags SecCSFlags, requirement *SecRequirementRef) int32
var _secRequirementCreateWithStringErr error

// CanCallSecRequirementCreateWithString reports whether the symbol for SecRequirementCreateWithString is available on this system.
func CanCallSecRequirementCreateWithString() bool {
	return _secRequirementCreateWithString != nil
}

// SecRequirementCreateWithStringCallError returns the symbol lookup or registration error for SecRequirementCreateWithString.
func SecRequirementCreateWithStringCallError() error {
	if _secRequirementCreateWithString != nil {
		return nil
	}
	return symbolCallError("SecRequirementCreateWithString", "10.0", _secRequirementCreateWithStringErr)
}

// TrySecRequirementCreateWithString calls SecRequirementCreateWithString without panicking when the symbol is unavailable.
func TrySecRequirementCreateWithString(text corefoundation.CFStringRef, flags SecCSFlags, requirement *SecRequirementRef) (int32, error) {
	if err := SecRequirementCreateWithStringCallError(); err != nil {
		return 0, err
	}
	return _secRequirementCreateWithString(text, flags, requirement), nil
}

// SecRequirementCreateWithString creates a code requirement object by compiling a valid text representation of a code requirement.
//
// See: https://developer.apple.com/documentation/Security/SecRequirementCreateWithString(_:_:_:)
func SecRequirementCreateWithString(text corefoundation.CFStringRef, flags SecCSFlags, requirement *SecRequirementRef) int32 {
	result, callErr := TrySecRequirementCreateWithString(text, flags, requirement)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secRequirementCreateWithStringAndErrors func(text corefoundation.CFStringRef, flags SecCSFlags, errors *corefoundation.CFErrorRef, requirement *SecRequirementRef) int32
var _secRequirementCreateWithStringAndErrorsErr error

// CanCallSecRequirementCreateWithStringAndErrors reports whether the symbol for SecRequirementCreateWithStringAndErrors is available on this system.
func CanCallSecRequirementCreateWithStringAndErrors() bool {
	return _secRequirementCreateWithStringAndErrors != nil
}

// SecRequirementCreateWithStringAndErrorsCallError returns the symbol lookup or registration error for SecRequirementCreateWithStringAndErrors.
func SecRequirementCreateWithStringAndErrorsCallError() error {
	if _secRequirementCreateWithStringAndErrors != nil {
		return nil
	}
	return symbolCallError("SecRequirementCreateWithStringAndErrors", "10.0", _secRequirementCreateWithStringAndErrorsErr)
}

// TrySecRequirementCreateWithStringAndErrors calls SecRequirementCreateWithStringAndErrors without panicking when the symbol is unavailable.
func TrySecRequirementCreateWithStringAndErrors(text corefoundation.CFStringRef, flags SecCSFlags, errors *corefoundation.CFErrorRef, requirement *SecRequirementRef) (int32, error) {
	if err := SecRequirementCreateWithStringAndErrorsCallError(); err != nil {
		return 0, err
	}
	return _secRequirementCreateWithStringAndErrors(text, flags, errors, requirement), nil
}

// SecRequirementCreateWithStringAndErrors creates a code requirement object by compiling a valid text representation of a code requirement and returns detailed error information in the case of failure.
//
// See: https://developer.apple.com/documentation/Security/SecRequirementCreateWithStringAndErrors(_:_:_:_:)
func SecRequirementCreateWithStringAndErrors(text corefoundation.CFStringRef, flags SecCSFlags, errors *corefoundation.CFErrorRef, requirement *SecRequirementRef) int32 {
	result, callErr := TrySecRequirementCreateWithStringAndErrors(text, flags, errors, requirement)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secRequirementGetTypeID func() uint
var _secRequirementGetTypeIDErr error

// CanCallSecRequirementGetTypeID reports whether the symbol for SecRequirementGetTypeID is available on this system.
func CanCallSecRequirementGetTypeID() bool {
	return _secRequirementGetTypeID != nil
}

// SecRequirementGetTypeIDCallError returns the symbol lookup or registration error for SecRequirementGetTypeID.
func SecRequirementGetTypeIDCallError() error {
	if _secRequirementGetTypeID != nil {
		return nil
	}
	return symbolCallError("SecRequirementGetTypeID", "10.0", _secRequirementGetTypeIDErr)
}

// TrySecRequirementGetTypeID calls SecRequirementGetTypeID without panicking when the symbol is unavailable.
func TrySecRequirementGetTypeID() (uint, error) {
	if err := SecRequirementGetTypeIDCallError(); err != nil {
		return 0, err
	}
	return _secRequirementGetTypeID(), nil
}

// SecRequirementGetTypeID returns the unique identifier of the opaque type to which a code requirement object belongs.
//
// See: https://developer.apple.com/documentation/Security/SecRequirementGetTypeID()
func SecRequirementGetTypeID() uint {
	result, callErr := TrySecRequirementGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secStaticCodeCheckValidity func(staticCode SecStaticCodeRef, flags SecCSFlags, requirement SecRequirementRef) int32
var _secStaticCodeCheckValidityErr error

// CanCallSecStaticCodeCheckValidity reports whether the symbol for SecStaticCodeCheckValidity is available on this system.
func CanCallSecStaticCodeCheckValidity() bool {
	return _secStaticCodeCheckValidity != nil
}

// SecStaticCodeCheckValidityCallError returns the symbol lookup or registration error for SecStaticCodeCheckValidity.
func SecStaticCodeCheckValidityCallError() error {
	if _secStaticCodeCheckValidity != nil {
		return nil
	}
	return symbolCallError("SecStaticCodeCheckValidity", "10.0", _secStaticCodeCheckValidityErr)
}

// TrySecStaticCodeCheckValidity calls SecStaticCodeCheckValidity without panicking when the symbol is unavailable.
func TrySecStaticCodeCheckValidity(staticCode SecStaticCodeRef, flags SecCSFlags, requirement SecRequirementRef) (int32, error) {
	if err := SecStaticCodeCheckValidityCallError(); err != nil {
		return 0, err
	}
	return _secStaticCodeCheckValidity(staticCode, flags, requirement), nil
}

// SecStaticCodeCheckValidity validates a static code object.
//
// See: https://developer.apple.com/documentation/Security/SecStaticCodeCheckValidity(_:_:_:)
func SecStaticCodeCheckValidity(staticCode SecStaticCodeRef, flags SecCSFlags, requirement SecRequirementRef) int32 {
	result, callErr := TrySecStaticCodeCheckValidity(staticCode, flags, requirement)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secStaticCodeCheckValidityWithErrors func(staticCode SecStaticCodeRef, flags SecCSFlags, requirement SecRequirementRef, errors *corefoundation.CFErrorRef) int32
var _secStaticCodeCheckValidityWithErrorsErr error

// CanCallSecStaticCodeCheckValidityWithErrors reports whether the symbol for SecStaticCodeCheckValidityWithErrors is available on this system.
func CanCallSecStaticCodeCheckValidityWithErrors() bool {
	return _secStaticCodeCheckValidityWithErrors != nil
}

// SecStaticCodeCheckValidityWithErrorsCallError returns the symbol lookup or registration error for SecStaticCodeCheckValidityWithErrors.
func SecStaticCodeCheckValidityWithErrorsCallError() error {
	if _secStaticCodeCheckValidityWithErrors != nil {
		return nil
	}
	return symbolCallError("SecStaticCodeCheckValidityWithErrors", "10.0", _secStaticCodeCheckValidityWithErrorsErr)
}

// TrySecStaticCodeCheckValidityWithErrors calls SecStaticCodeCheckValidityWithErrors without panicking when the symbol is unavailable.
func TrySecStaticCodeCheckValidityWithErrors(staticCode SecStaticCodeRef, flags SecCSFlags, requirement SecRequirementRef, errors *corefoundation.CFErrorRef) (int32, error) {
	if err := SecStaticCodeCheckValidityWithErrorsCallError(); err != nil {
		return 0, err
	}
	return _secStaticCodeCheckValidityWithErrors(staticCode, flags, requirement, errors), nil
}

// SecStaticCodeCheckValidityWithErrors performs static validation of static signed code and returns detailed error information in the case of failure.
//
// See: https://developer.apple.com/documentation/Security/SecStaticCodeCheckValidityWithErrors(_:_:_:_:)
func SecStaticCodeCheckValidityWithErrors(staticCode SecStaticCodeRef, flags SecCSFlags, requirement SecRequirementRef, errors *corefoundation.CFErrorRef) int32 {
	result, callErr := TrySecStaticCodeCheckValidityWithErrors(staticCode, flags, requirement, errors)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secStaticCodeCreateWithPath func(path corefoundation.CFURLRef, flags SecCSFlags, staticCode *SecStaticCodeRef) int32
var _secStaticCodeCreateWithPathErr error

// CanCallSecStaticCodeCreateWithPath reports whether the symbol for SecStaticCodeCreateWithPath is available on this system.
func CanCallSecStaticCodeCreateWithPath() bool {
	return _secStaticCodeCreateWithPath != nil
}

// SecStaticCodeCreateWithPathCallError returns the symbol lookup or registration error for SecStaticCodeCreateWithPath.
func SecStaticCodeCreateWithPathCallError() error {
	if _secStaticCodeCreateWithPath != nil {
		return nil
	}
	return symbolCallError("SecStaticCodeCreateWithPath", "10.0", _secStaticCodeCreateWithPathErr)
}

// TrySecStaticCodeCreateWithPath calls SecStaticCodeCreateWithPath without panicking when the symbol is unavailable.
func TrySecStaticCodeCreateWithPath(path corefoundation.CFURLRef, flags SecCSFlags, staticCode *SecStaticCodeRef) (int32, error) {
	if err := SecStaticCodeCreateWithPathCallError(); err != nil {
		return 0, err
	}
	return _secStaticCodeCreateWithPath(path, flags, staticCode), nil
}

// SecStaticCodeCreateWithPath creates a static code object representing the code at a specified file system path.
//
// See: https://developer.apple.com/documentation/Security/SecStaticCodeCreateWithPath(_:_:_:)
func SecStaticCodeCreateWithPath(path corefoundation.CFURLRef, flags SecCSFlags, staticCode *SecStaticCodeRef) int32 {
	result, callErr := TrySecStaticCodeCreateWithPath(path, flags, staticCode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secStaticCodeCreateWithPathAndAttributes func(path corefoundation.CFURLRef, flags SecCSFlags, attributes corefoundation.CFDictionaryRef, staticCode *SecStaticCodeRef) int32
var _secStaticCodeCreateWithPathAndAttributesErr error

// CanCallSecStaticCodeCreateWithPathAndAttributes reports whether the symbol for SecStaticCodeCreateWithPathAndAttributes is available on this system.
func CanCallSecStaticCodeCreateWithPathAndAttributes() bool {
	return _secStaticCodeCreateWithPathAndAttributes != nil
}

// SecStaticCodeCreateWithPathAndAttributesCallError returns the symbol lookup or registration error for SecStaticCodeCreateWithPathAndAttributes.
func SecStaticCodeCreateWithPathAndAttributesCallError() error {
	if _secStaticCodeCreateWithPathAndAttributes != nil {
		return nil
	}
	return symbolCallError("SecStaticCodeCreateWithPathAndAttributes", "10.0", _secStaticCodeCreateWithPathAndAttributesErr)
}

// TrySecStaticCodeCreateWithPathAndAttributes calls SecStaticCodeCreateWithPathAndAttributes without panicking when the symbol is unavailable.
func TrySecStaticCodeCreateWithPathAndAttributes(path corefoundation.CFURLRef, flags SecCSFlags, attributes corefoundation.CFDictionaryRef, staticCode *SecStaticCodeRef) (int32, error) {
	if err := SecStaticCodeCreateWithPathAndAttributesCallError(); err != nil {
		return 0, err
	}
	return _secStaticCodeCreateWithPathAndAttributes(path, flags, attributes, staticCode), nil
}

// SecStaticCodeCreateWithPathAndAttributes creates a static code object representing the code at a specified file system path using an attributes dictionary.
//
// See: https://developer.apple.com/documentation/Security/SecStaticCodeCreateWithPathAndAttributes(_:_:_:_:)
func SecStaticCodeCreateWithPathAndAttributes(path corefoundation.CFURLRef, flags SecCSFlags, attributes corefoundation.CFDictionaryRef, staticCode *SecStaticCodeRef) int32 {
	result, callErr := TrySecStaticCodeCreateWithPathAndAttributes(path, flags, attributes, staticCode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secStaticCodeGetTypeID func() uint
var _secStaticCodeGetTypeIDErr error

// CanCallSecStaticCodeGetTypeID reports whether the symbol for SecStaticCodeGetTypeID is available on this system.
func CanCallSecStaticCodeGetTypeID() bool {
	return _secStaticCodeGetTypeID != nil
}

// SecStaticCodeGetTypeIDCallError returns the symbol lookup or registration error for SecStaticCodeGetTypeID.
func SecStaticCodeGetTypeIDCallError() error {
	if _secStaticCodeGetTypeID != nil {
		return nil
	}
	return symbolCallError("SecStaticCodeGetTypeID", "10.0", _secStaticCodeGetTypeIDErr)
}

// TrySecStaticCodeGetTypeID calls SecStaticCodeGetTypeID without panicking when the symbol is unavailable.
func TrySecStaticCodeGetTypeID() (uint, error) {
	if err := SecStaticCodeGetTypeIDCallError(); err != nil {
		return 0, err
	}
	return _secStaticCodeGetTypeID(), nil
}

// SecStaticCodeGetTypeID returns the unique identifier of the opaque type to which a static code object belongs.
//
// See: https://developer.apple.com/documentation/Security/SecStaticCodeGetTypeID()
func SecStaticCodeGetTypeID() uint {
	result, callErr := TrySecStaticCodeGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTaskCopySigningIdentifier func(task SecTaskRef, err *corefoundation.CFErrorRef) corefoundation.CFStringRef
var _secTaskCopySigningIdentifierErr error

// CanCallSecTaskCopySigningIdentifier reports whether the symbol for SecTaskCopySigningIdentifier is available on this system.
func CanCallSecTaskCopySigningIdentifier() bool {
	return _secTaskCopySigningIdentifier != nil
}

// SecTaskCopySigningIdentifierCallError returns the symbol lookup or registration error for SecTaskCopySigningIdentifier.
func SecTaskCopySigningIdentifierCallError() error {
	if _secTaskCopySigningIdentifier != nil {
		return nil
	}
	return symbolCallError("SecTaskCopySigningIdentifier", "10.0", _secTaskCopySigningIdentifierErr)
}

// TrySecTaskCopySigningIdentifier calls SecTaskCopySigningIdentifier without panicking when the symbol is unavailable.
func TrySecTaskCopySigningIdentifier(task SecTaskRef, err *corefoundation.CFErrorRef) (corefoundation.CFStringRef, error) {
	if err := SecTaskCopySigningIdentifierCallError(); err != nil {
		return 0, err
	}
	return _secTaskCopySigningIdentifier(task, err), nil
}

// SecTaskCopySigningIdentifier returns the value of the code signing identifier.
//
// See: https://developer.apple.com/documentation/Security/SecTaskCopySigningIdentifier(_:_:)
func SecTaskCopySigningIdentifier(task SecTaskRef, err *corefoundation.CFErrorRef) corefoundation.CFStringRef {
	result, callErr := TrySecTaskCopySigningIdentifier(task, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTaskCopyValueForEntitlement func(task SecTaskRef, entitlement corefoundation.CFStringRef, err *corefoundation.CFErrorRef) corefoundation.CFTypeRef
var _secTaskCopyValueForEntitlementErr error

// CanCallSecTaskCopyValueForEntitlement reports whether the symbol for SecTaskCopyValueForEntitlement is available on this system.
func CanCallSecTaskCopyValueForEntitlement() bool {
	return _secTaskCopyValueForEntitlement != nil
}

// SecTaskCopyValueForEntitlementCallError returns the symbol lookup or registration error for SecTaskCopyValueForEntitlement.
func SecTaskCopyValueForEntitlementCallError() error {
	if _secTaskCopyValueForEntitlement != nil {
		return nil
	}
	return symbolCallError("SecTaskCopyValueForEntitlement", "10.0", _secTaskCopyValueForEntitlementErr)
}

// TrySecTaskCopyValueForEntitlement calls SecTaskCopyValueForEntitlement without panicking when the symbol is unavailable.
func TrySecTaskCopyValueForEntitlement(task SecTaskRef, entitlement corefoundation.CFStringRef, err *corefoundation.CFErrorRef) (corefoundation.CFTypeRef, error) {
	if err := SecTaskCopyValueForEntitlementCallError(); err != nil {
		return 0, err
	}
	return _secTaskCopyValueForEntitlement(task, entitlement, err), nil
}

// SecTaskCopyValueForEntitlement returns the value of a single entitlement for the represented task.
//
// See: https://developer.apple.com/documentation/Security/SecTaskCopyValueForEntitlement(_:_:_:)
func SecTaskCopyValueForEntitlement(task SecTaskRef, entitlement corefoundation.CFStringRef, err *corefoundation.CFErrorRef) corefoundation.CFTypeRef {
	result, callErr := TrySecTaskCopyValueForEntitlement(task, entitlement, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTaskCopyValuesForEntitlements func(task SecTaskRef, entitlements corefoundation.CFArrayRef, err *corefoundation.CFErrorRef) corefoundation.CFDictionaryRef
var _secTaskCopyValuesForEntitlementsErr error

// CanCallSecTaskCopyValuesForEntitlements reports whether the symbol for SecTaskCopyValuesForEntitlements is available on this system.
func CanCallSecTaskCopyValuesForEntitlements() bool {
	return _secTaskCopyValuesForEntitlements != nil
}

// SecTaskCopyValuesForEntitlementsCallError returns the symbol lookup or registration error for SecTaskCopyValuesForEntitlements.
func SecTaskCopyValuesForEntitlementsCallError() error {
	if _secTaskCopyValuesForEntitlements != nil {
		return nil
	}
	return symbolCallError("SecTaskCopyValuesForEntitlements", "10.0", _secTaskCopyValuesForEntitlementsErr)
}

// TrySecTaskCopyValuesForEntitlements calls SecTaskCopyValuesForEntitlements without panicking when the symbol is unavailable.
func TrySecTaskCopyValuesForEntitlements(task SecTaskRef, entitlements corefoundation.CFArrayRef, err *corefoundation.CFErrorRef) (corefoundation.CFDictionaryRef, error) {
	if err := SecTaskCopyValuesForEntitlementsCallError(); err != nil {
		return 0, err
	}
	return _secTaskCopyValuesForEntitlements(task, entitlements, err), nil
}

// SecTaskCopyValuesForEntitlements returns the values of multiple entitlements for the represented task.
//
// See: https://developer.apple.com/documentation/Security/SecTaskCopyValuesForEntitlements(_:_:_:)
func SecTaskCopyValuesForEntitlements(task SecTaskRef, entitlements corefoundation.CFArrayRef, err *corefoundation.CFErrorRef) corefoundation.CFDictionaryRef {
	result, callErr := TrySecTaskCopyValuesForEntitlements(task, entitlements, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTaskCreateFromSelf func(allocator corefoundation.CFAllocatorRef) SecTaskRef
var _secTaskCreateFromSelfErr error

// CanCallSecTaskCreateFromSelf reports whether the symbol for SecTaskCreateFromSelf is available on this system.
func CanCallSecTaskCreateFromSelf() bool {
	return _secTaskCreateFromSelf != nil
}

// SecTaskCreateFromSelfCallError returns the symbol lookup or registration error for SecTaskCreateFromSelf.
func SecTaskCreateFromSelfCallError() error {
	if _secTaskCreateFromSelf != nil {
		return nil
	}
	return symbolCallError("SecTaskCreateFromSelf", "10.0", _secTaskCreateFromSelfErr)
}

// TrySecTaskCreateFromSelf calls SecTaskCreateFromSelf without panicking when the symbol is unavailable.
func TrySecTaskCreateFromSelf(allocator corefoundation.CFAllocatorRef) (SecTaskRef, error) {
	if err := SecTaskCreateFromSelfCallError(); err != nil {
		return 0, err
	}
	return _secTaskCreateFromSelf(allocator), nil
}

// SecTaskCreateFromSelf creates a task object for the current task.
//
// See: https://developer.apple.com/documentation/Security/SecTaskCreateFromSelf(_:)
func SecTaskCreateFromSelf(allocator corefoundation.CFAllocatorRef) SecTaskRef {
	result, callErr := TrySecTaskCreateFromSelf(allocator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTaskGetTypeID func() uint
var _secTaskGetTypeIDErr error

// CanCallSecTaskGetTypeID reports whether the symbol for SecTaskGetTypeID is available on this system.
func CanCallSecTaskGetTypeID() bool {
	return _secTaskGetTypeID != nil
}

// SecTaskGetTypeIDCallError returns the symbol lookup or registration error for SecTaskGetTypeID.
func SecTaskGetTypeIDCallError() error {
	if _secTaskGetTypeID != nil {
		return nil
	}
	return symbolCallError("SecTaskGetTypeID", "10.0", _secTaskGetTypeIDErr)
}

// TrySecTaskGetTypeID calls SecTaskGetTypeID without panicking when the symbol is unavailable.
func TrySecTaskGetTypeID() (uint, error) {
	if err := SecTaskGetTypeIDCallError(); err != nil {
		return 0, err
	}
	return _secTaskGetTypeID(), nil
}

// SecTaskGetTypeID returns the unique identifier of the opaque type to which a task object belongs.
//
// See: https://developer.apple.com/documentation/Security/SecTaskGetTypeID()
func SecTaskGetTypeID() uint {
	result, callErr := TrySecTaskGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustCopyAnchorCertificates func(anchors *corefoundation.CFArrayRef) int32
var _secTrustCopyAnchorCertificatesErr error

// CanCallSecTrustCopyAnchorCertificates reports whether the symbol for SecTrustCopyAnchorCertificates is available on this system.
func CanCallSecTrustCopyAnchorCertificates() bool {
	return _secTrustCopyAnchorCertificates != nil
}

// SecTrustCopyAnchorCertificatesCallError returns the symbol lookup or registration error for SecTrustCopyAnchorCertificates.
func SecTrustCopyAnchorCertificatesCallError() error {
	if _secTrustCopyAnchorCertificates != nil {
		return nil
	}
	return symbolCallError("SecTrustCopyAnchorCertificates", "10.3", _secTrustCopyAnchorCertificatesErr)
}

// TrySecTrustCopyAnchorCertificates calls SecTrustCopyAnchorCertificates without panicking when the symbol is unavailable.
func TrySecTrustCopyAnchorCertificates(anchors *corefoundation.CFArrayRef) (int32, error) {
	if err := SecTrustCopyAnchorCertificatesCallError(); err != nil {
		return 0, err
	}
	return _secTrustCopyAnchorCertificates(anchors), nil
}

// SecTrustCopyAnchorCertificates retrieves the anchor (root) certificates stored by macOS.
//
// See: https://developer.apple.com/documentation/Security/SecTrustCopyAnchorCertificates(_:)
func SecTrustCopyAnchorCertificates(anchors *corefoundation.CFArrayRef) int32 {
	result, callErr := TrySecTrustCopyAnchorCertificates(anchors)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustCopyCertificateChain func(trust SecTrustRef) corefoundation.CFArrayRef
var _secTrustCopyCertificateChainErr error

// CanCallSecTrustCopyCertificateChain reports whether the symbol for SecTrustCopyCertificateChain is available on this system.
func CanCallSecTrustCopyCertificateChain() bool {
	return _secTrustCopyCertificateChain != nil
}

// SecTrustCopyCertificateChainCallError returns the symbol lookup or registration error for SecTrustCopyCertificateChain.
func SecTrustCopyCertificateChainCallError() error {
	if _secTrustCopyCertificateChain != nil {
		return nil
	}
	return symbolCallError("SecTrustCopyCertificateChain", "12.0", _secTrustCopyCertificateChainErr)
}

// TrySecTrustCopyCertificateChain calls SecTrustCopyCertificateChain without panicking when the symbol is unavailable.
func TrySecTrustCopyCertificateChain(trust SecTrustRef) (corefoundation.CFArrayRef, error) {
	if err := SecTrustCopyCertificateChainCallError(); err != nil {
		return 0, err
	}
	return _secTrustCopyCertificateChain(trust), nil
}

// SecTrustCopyCertificateChain.
//
// See: https://developer.apple.com/documentation/Security/SecTrustCopyCertificateChain(_:)
func SecTrustCopyCertificateChain(trust SecTrustRef) corefoundation.CFArrayRef {
	result, callErr := TrySecTrustCopyCertificateChain(trust)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustCopyCustomAnchorCertificates func(trust SecTrustRef, anchors *corefoundation.CFArrayRef) int32
var _secTrustCopyCustomAnchorCertificatesErr error

// CanCallSecTrustCopyCustomAnchorCertificates reports whether the symbol for SecTrustCopyCustomAnchorCertificates is available on this system.
func CanCallSecTrustCopyCustomAnchorCertificates() bool {
	return _secTrustCopyCustomAnchorCertificates != nil
}

// SecTrustCopyCustomAnchorCertificatesCallError returns the symbol lookup or registration error for SecTrustCopyCustomAnchorCertificates.
func SecTrustCopyCustomAnchorCertificatesCallError() error {
	if _secTrustCopyCustomAnchorCertificates != nil {
		return nil
	}
	return symbolCallError("SecTrustCopyCustomAnchorCertificates", "10.5", _secTrustCopyCustomAnchorCertificatesErr)
}

// TrySecTrustCopyCustomAnchorCertificates calls SecTrustCopyCustomAnchorCertificates without panicking when the symbol is unavailable.
func TrySecTrustCopyCustomAnchorCertificates(trust SecTrustRef, anchors *corefoundation.CFArrayRef) (int32, error) {
	if err := SecTrustCopyCustomAnchorCertificatesCallError(); err != nil {
		return 0, err
	}
	return _secTrustCopyCustomAnchorCertificates(trust, anchors), nil
}

// SecTrustCopyCustomAnchorCertificates retrieves the custom anchor certificates, if any, used by a given trust.
//
// See: https://developer.apple.com/documentation/Security/SecTrustCopyCustomAnchorCertificates(_:_:)
func SecTrustCopyCustomAnchorCertificates(trust SecTrustRef, anchors *corefoundation.CFArrayRef) int32 {
	result, callErr := TrySecTrustCopyCustomAnchorCertificates(trust, anchors)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustCopyExceptions func(trust SecTrustRef) corefoundation.CFDataRef
var _secTrustCopyExceptionsErr error

// CanCallSecTrustCopyExceptions reports whether the symbol for SecTrustCopyExceptions is available on this system.
func CanCallSecTrustCopyExceptions() bool {
	return _secTrustCopyExceptions != nil
}

// SecTrustCopyExceptionsCallError returns the symbol lookup or registration error for SecTrustCopyExceptions.
func SecTrustCopyExceptionsCallError() error {
	if _secTrustCopyExceptions != nil {
		return nil
	}
	return symbolCallError("SecTrustCopyExceptions", "10.9", _secTrustCopyExceptionsErr)
}

// TrySecTrustCopyExceptions calls SecTrustCopyExceptions without panicking when the symbol is unavailable.
func TrySecTrustCopyExceptions(trust SecTrustRef) (corefoundation.CFDataRef, error) {
	if err := SecTrustCopyExceptionsCallError(); err != nil {
		return 0, err
	}
	return _secTrustCopyExceptions(trust), nil
}

// SecTrustCopyExceptions returns an opaque cookie containing exceptions to trust policies that will allow future evaluations of the current certificate to succeed.
//
// See: https://developer.apple.com/documentation/Security/SecTrustCopyExceptions(_:)
func SecTrustCopyExceptions(trust SecTrustRef) corefoundation.CFDataRef {
	result, callErr := TrySecTrustCopyExceptions(trust)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustCopyKey func(trust SecTrustRef) SecKeyRef
var _secTrustCopyKeyErr error

// CanCallSecTrustCopyKey reports whether the symbol for SecTrustCopyKey is available on this system.
func CanCallSecTrustCopyKey() bool {
	return _secTrustCopyKey != nil
}

// SecTrustCopyKeyCallError returns the symbol lookup or registration error for SecTrustCopyKey.
func SecTrustCopyKeyCallError() error {
	if _secTrustCopyKey != nil {
		return nil
	}
	return symbolCallError("SecTrustCopyKey", "11.0", _secTrustCopyKeyErr)
}

// TrySecTrustCopyKey calls SecTrustCopyKey without panicking when the symbol is unavailable.
func TrySecTrustCopyKey(trust SecTrustRef) (SecKeyRef, error) {
	if err := SecTrustCopyKeyCallError(); err != nil {
		return 0, err
	}
	return _secTrustCopyKey(trust), nil
}

// SecTrustCopyKey.
//
// See: https://developer.apple.com/documentation/Security/SecTrustCopyKey(_:)
func SecTrustCopyKey(trust SecTrustRef) SecKeyRef {
	result, callErr := TrySecTrustCopyKey(trust)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustCopyPolicies func(trust SecTrustRef, policies *corefoundation.CFArrayRef) int32
var _secTrustCopyPoliciesErr error

// CanCallSecTrustCopyPolicies reports whether the symbol for SecTrustCopyPolicies is available on this system.
func CanCallSecTrustCopyPolicies() bool {
	return _secTrustCopyPolicies != nil
}

// SecTrustCopyPoliciesCallError returns the symbol lookup or registration error for SecTrustCopyPolicies.
func SecTrustCopyPoliciesCallError() error {
	if _secTrustCopyPolicies != nil {
		return nil
	}
	return symbolCallError("SecTrustCopyPolicies", "10.3", _secTrustCopyPoliciesErr)
}

// TrySecTrustCopyPolicies calls SecTrustCopyPolicies without panicking when the symbol is unavailable.
func TrySecTrustCopyPolicies(trust SecTrustRef, policies *corefoundation.CFArrayRef) (int32, error) {
	if err := SecTrustCopyPoliciesCallError(); err != nil {
		return 0, err
	}
	return _secTrustCopyPolicies(trust, policies), nil
}

// SecTrustCopyPolicies retrieves the policies used by a given trust management object.
//
// See: https://developer.apple.com/documentation/Security/SecTrustCopyPolicies(_:_:)
func SecTrustCopyPolicies(trust SecTrustRef, policies *corefoundation.CFArrayRef) int32 {
	result, callErr := TrySecTrustCopyPolicies(trust, policies)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustCopyProperties func(trust SecTrustRef) corefoundation.CFArrayRef
var _secTrustCopyPropertiesErr error

// CanCallSecTrustCopyProperties reports whether the symbol for SecTrustCopyProperties is available on this system.
func CanCallSecTrustCopyProperties() bool {
	return _secTrustCopyProperties != nil
}

// SecTrustCopyPropertiesCallError returns the symbol lookup or registration error for SecTrustCopyProperties.
func SecTrustCopyPropertiesCallError() error {
	if _secTrustCopyProperties != nil {
		return nil
	}
	return symbolCallError("SecTrustCopyProperties", "10.7", _secTrustCopyPropertiesErr)
}

// TrySecTrustCopyProperties calls SecTrustCopyProperties without panicking when the symbol is unavailable.
func TrySecTrustCopyProperties(trust SecTrustRef) (corefoundation.CFArrayRef, error) {
	if err := SecTrustCopyPropertiesCallError(); err != nil {
		return 0, err
	}
	return _secTrustCopyProperties(trust), nil
}

// SecTrustCopyProperties returns an array containing the properties of a trust object.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecTrustCopyProperties(_:)
func SecTrustCopyProperties(trust SecTrustRef) corefoundation.CFArrayRef {
	result, callErr := TrySecTrustCopyProperties(trust)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustCopyPublicKey func(trust SecTrustRef) SecKeyRef
var _secTrustCopyPublicKeyErr error

// CanCallSecTrustCopyPublicKey reports whether the symbol for SecTrustCopyPublicKey is available on this system.
func CanCallSecTrustCopyPublicKey() bool {
	return _secTrustCopyPublicKey != nil
}

// SecTrustCopyPublicKeyCallError returns the symbol lookup or registration error for SecTrustCopyPublicKey.
func SecTrustCopyPublicKeyCallError() error {
	if _secTrustCopyPublicKey != nil {
		return nil
	}
	return symbolCallError("SecTrustCopyPublicKey", "10.7", _secTrustCopyPublicKeyErr)
}

// TrySecTrustCopyPublicKey calls SecTrustCopyPublicKey without panicking when the symbol is unavailable.
func TrySecTrustCopyPublicKey(trust SecTrustRef) (SecKeyRef, error) {
	if err := SecTrustCopyPublicKeyCallError(); err != nil {
		return 0, err
	}
	return _secTrustCopyPublicKey(trust), nil
}

// SecTrustCopyPublicKey returns the public key for a leaf certificate after it has been evaluated.
//
// Deprecated: Deprecated since macOS 11.0.
//
// See: https://developer.apple.com/documentation/Security/SecTrustCopyPublicKey(_:)
func SecTrustCopyPublicKey(trust SecTrustRef) SecKeyRef {
	result, callErr := TrySecTrustCopyPublicKey(trust)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustCopyResult func(trust SecTrustRef) corefoundation.CFDictionaryRef
var _secTrustCopyResultErr error

// CanCallSecTrustCopyResult reports whether the symbol for SecTrustCopyResult is available on this system.
func CanCallSecTrustCopyResult() bool {
	return _secTrustCopyResult != nil
}

// SecTrustCopyResultCallError returns the symbol lookup or registration error for SecTrustCopyResult.
func SecTrustCopyResultCallError() error {
	if _secTrustCopyResult != nil {
		return nil
	}
	return symbolCallError("SecTrustCopyResult", "10.9", _secTrustCopyResultErr)
}

// TrySecTrustCopyResult calls SecTrustCopyResult without panicking when the symbol is unavailable.
func TrySecTrustCopyResult(trust SecTrustRef) (corefoundation.CFDictionaryRef, error) {
	if err := SecTrustCopyResultCallError(); err != nil {
		return 0, err
	}
	return _secTrustCopyResult(trust), nil
}

// SecTrustCopyResult returns a dictionary containing information about an evaluated trust.
//
// See: https://developer.apple.com/documentation/Security/SecTrustCopyResult(_:)
func SecTrustCopyResult(trust SecTrustRef) corefoundation.CFDictionaryRef {
	result, callErr := TrySecTrustCopyResult(trust)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustCreateWithCertificates func(certificates corefoundation.CFTypeRef, policies corefoundation.CFTypeRef, trust *SecTrustRef) int32
var _secTrustCreateWithCertificatesErr error

// CanCallSecTrustCreateWithCertificates reports whether the symbol for SecTrustCreateWithCertificates is available on this system.
func CanCallSecTrustCreateWithCertificates() bool {
	return _secTrustCreateWithCertificates != nil
}

// SecTrustCreateWithCertificatesCallError returns the symbol lookup or registration error for SecTrustCreateWithCertificates.
func SecTrustCreateWithCertificatesCallError() error {
	if _secTrustCreateWithCertificates != nil {
		return nil
	}
	return symbolCallError("SecTrustCreateWithCertificates", "10.3", _secTrustCreateWithCertificatesErr)
}

// TrySecTrustCreateWithCertificates calls SecTrustCreateWithCertificates without panicking when the symbol is unavailable.
func TrySecTrustCreateWithCertificates(certificates corefoundation.CFTypeRef, policies corefoundation.CFTypeRef, trust *SecTrustRef) (int32, error) {
	if err := SecTrustCreateWithCertificatesCallError(); err != nil {
		return 0, err
	}
	return _secTrustCreateWithCertificates(certificates, policies, trust), nil
}

// SecTrustCreateWithCertificates creates a trust management object based on certificates and policies.
//
// See: https://developer.apple.com/documentation/Security/SecTrustCreateWithCertificates(_:_:_:)
func SecTrustCreateWithCertificates(certificates corefoundation.CFTypeRef, policies corefoundation.CFTypeRef, trust *SecTrustRef) int32 {
	result, callErr := TrySecTrustCreateWithCertificates(certificates, policies, trust)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustEvaluateAsyncWithError func(trust SecTrustRef, queue uintptr, result unsafe.Pointer) int32
var _secTrustEvaluateAsyncWithErrorErr error

// CanCallSecTrustEvaluateAsyncWithError reports whether the symbol for SecTrustEvaluateAsyncWithError is available on this system.
func CanCallSecTrustEvaluateAsyncWithError() bool {
	return _secTrustEvaluateAsyncWithError != nil
}

// SecTrustEvaluateAsyncWithErrorCallError returns the symbol lookup or registration error for SecTrustEvaluateAsyncWithError.
func SecTrustEvaluateAsyncWithErrorCallError() error {
	if _secTrustEvaluateAsyncWithError != nil {
		return nil
	}
	return symbolCallError("SecTrustEvaluateAsyncWithError", "10.15", _secTrustEvaluateAsyncWithErrorErr)
}

// TrySecTrustEvaluateAsyncWithError calls SecTrustEvaluateAsyncWithError without panicking when the symbol is unavailable.
func TrySecTrustEvaluateAsyncWithError(trust SecTrustRef, queue dispatch.Queue, result SecTrustWithErrorCallback) (int32, error) {
	if err := SecTrustEvaluateAsyncWithErrorCallError(); err != nil {
		return 0, err
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.IObject, arg1 bool, arg2 objectivec.IObject) {
		result(arg0, arg1, arg2)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _secTrustEvaluateAsyncWithError(trust, uintptr(queue.Handle()), _block0), nil
}

// SecTrustEvaluateAsyncWithError evaluates a trust object asynchronously on the specified dispatch queue.
//
// See: https://developer.apple.com/documentation/Security/SecTrustEvaluateAsyncWithError(_:_:_:)
func SecTrustEvaluateAsyncWithError(trust SecTrustRef, queue dispatch.Queue, result SecTrustWithErrorCallback) int32 {
	result0, callErr := TrySecTrustEvaluateAsyncWithError(trust, queue, result)
	if callErr != nil {
		panic(callErr)
	}
	return result0
}

var _secTrustEvaluateWithError func(trust SecTrustRef, err *corefoundation.CFErrorRef) bool
var _secTrustEvaluateWithErrorErr error

// CanCallSecTrustEvaluateWithError reports whether the symbol for SecTrustEvaluateWithError is available on this system.
func CanCallSecTrustEvaluateWithError() bool {
	return _secTrustEvaluateWithError != nil
}

// SecTrustEvaluateWithErrorCallError returns the symbol lookup or registration error for SecTrustEvaluateWithError.
func SecTrustEvaluateWithErrorCallError() error {
	if _secTrustEvaluateWithError != nil {
		return nil
	}
	return symbolCallError("SecTrustEvaluateWithError", "10.14", _secTrustEvaluateWithErrorErr)
}

// TrySecTrustEvaluateWithError calls SecTrustEvaluateWithError without panicking when the symbol is unavailable.
func TrySecTrustEvaluateWithError(trust SecTrustRef, err *corefoundation.CFErrorRef) (bool, error) {
	if err := SecTrustEvaluateWithErrorCallError(); err != nil {
		return false, err
	}
	return _secTrustEvaluateWithError(trust, err), nil
}

// SecTrustEvaluateWithError evaluates trust for the specified certificate and policies.
//
// See: https://developer.apple.com/documentation/Security/SecTrustEvaluateWithError(_:_:)
func SecTrustEvaluateWithError(trust SecTrustRef, err *corefoundation.CFErrorRef) bool {
	result, callErr := TrySecTrustEvaluateWithError(trust, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustGetCertificateAtIndex func(trust SecTrustRef, ix int) SecCertificateRef
var _secTrustGetCertificateAtIndexErr error

// CanCallSecTrustGetCertificateAtIndex reports whether the symbol for SecTrustGetCertificateAtIndex is available on this system.
func CanCallSecTrustGetCertificateAtIndex() bool {
	return _secTrustGetCertificateAtIndex != nil
}

// SecTrustGetCertificateAtIndexCallError returns the symbol lookup or registration error for SecTrustGetCertificateAtIndex.
func SecTrustGetCertificateAtIndexCallError() error {
	if _secTrustGetCertificateAtIndex != nil {
		return nil
	}
	return symbolCallError("SecTrustGetCertificateAtIndex", "10.7", _secTrustGetCertificateAtIndexErr)
}

// TrySecTrustGetCertificateAtIndex calls SecTrustGetCertificateAtIndex without panicking when the symbol is unavailable.
func TrySecTrustGetCertificateAtIndex(trust SecTrustRef, ix int) (SecCertificateRef, error) {
	if err := SecTrustGetCertificateAtIndexCallError(); err != nil {
		return 0, err
	}
	return _secTrustGetCertificateAtIndex(trust, ix), nil
}

// SecTrustGetCertificateAtIndex returns a specific certificate from the certificate chain used to evaluate trust.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecTrustGetCertificateAtIndex(_:_:)
func SecTrustGetCertificateAtIndex(trust SecTrustRef, ix int) SecCertificateRef {
	result, callErr := TrySecTrustGetCertificateAtIndex(trust, ix)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustGetCertificateCount func(trust SecTrustRef) int
var _secTrustGetCertificateCountErr error

// CanCallSecTrustGetCertificateCount reports whether the symbol for SecTrustGetCertificateCount is available on this system.
func CanCallSecTrustGetCertificateCount() bool {
	return _secTrustGetCertificateCount != nil
}

// SecTrustGetCertificateCountCallError returns the symbol lookup or registration error for SecTrustGetCertificateCount.
func SecTrustGetCertificateCountCallError() error {
	if _secTrustGetCertificateCount != nil {
		return nil
	}
	return symbolCallError("SecTrustGetCertificateCount", "10.7", _secTrustGetCertificateCountErr)
}

// TrySecTrustGetCertificateCount calls SecTrustGetCertificateCount without panicking when the symbol is unavailable.
func TrySecTrustGetCertificateCount(trust SecTrustRef) (int, error) {
	if err := SecTrustGetCertificateCountCallError(); err != nil {
		return 0, err
	}
	return _secTrustGetCertificateCount(trust), nil
}

// SecTrustGetCertificateCount returns the number of certificates in an evaluated certificate chain.
//
// See: https://developer.apple.com/documentation/Security/SecTrustGetCertificateCount(_:)
func SecTrustGetCertificateCount(trust SecTrustRef) int {
	result, callErr := TrySecTrustGetCertificateCount(trust)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustGetCssmResult func(trust SecTrustRef, result unsafe.Pointer) int32
var _secTrustGetCssmResultErr error

// CanCallSecTrustGetCssmResult reports whether the symbol for SecTrustGetCssmResult is available on this system.
func CanCallSecTrustGetCssmResult() bool {
	return _secTrustGetCssmResult != nil
}

// SecTrustGetCssmResultCallError returns the symbol lookup or registration error for SecTrustGetCssmResult.
func SecTrustGetCssmResultCallError() error {
	if _secTrustGetCssmResult != nil {
		return nil
	}
	return symbolCallError("SecTrustGetCssmResult", "10.2", _secTrustGetCssmResultErr)
}

// TrySecTrustGetCssmResult calls SecTrustGetCssmResult without panicking when the symbol is unavailable.
func TrySecTrustGetCssmResult(trust SecTrustRef, result unsafe.Pointer) (int32, error) {
	if err := SecTrustGetCssmResultCallError(); err != nil {
		return 0, err
	}
	return _secTrustGetCssmResult(trust, result), nil
}

// SecTrustGetCssmResult retrieves the CSSM trust result.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecTrustGetCssmResult
func SecTrustGetCssmResult(trust SecTrustRef, result unsafe.Pointer) int32 {
	result0, callErr := TrySecTrustGetCssmResult(trust, result)
	if callErr != nil {
		panic(callErr)
	}
	return result0
}

var _secTrustGetCssmResultCode func(trust SecTrustRef, resultCode *int32) int32
var _secTrustGetCssmResultCodeErr error

// CanCallSecTrustGetCssmResultCode reports whether the symbol for SecTrustGetCssmResultCode is available on this system.
func CanCallSecTrustGetCssmResultCode() bool {
	return _secTrustGetCssmResultCode != nil
}

// SecTrustGetCssmResultCodeCallError returns the symbol lookup or registration error for SecTrustGetCssmResultCode.
func SecTrustGetCssmResultCodeCallError() error {
	if _secTrustGetCssmResultCode != nil {
		return nil
	}
	return symbolCallError("SecTrustGetCssmResultCode", "10.2", _secTrustGetCssmResultCodeErr)
}

// TrySecTrustGetCssmResultCode calls SecTrustGetCssmResultCode without panicking when the symbol is unavailable.
func TrySecTrustGetCssmResultCode(trust SecTrustRef, resultCode *int32) (int32, error) {
	if err := SecTrustGetCssmResultCodeCallError(); err != nil {
		return 0, err
	}
	return _secTrustGetCssmResultCode(trust, resultCode), nil
}

// SecTrustGetCssmResultCode retrieves the CSSM result code from the most recent trust evaluation for a trust management object.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecTrustGetCssmResultCode
func SecTrustGetCssmResultCode(trust SecTrustRef, resultCode *int32) int32 {
	result, callErr := TrySecTrustGetCssmResultCode(trust, resultCode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustGetNetworkFetchAllowed func(trust SecTrustRef, allowFetch *bool) int32
var _secTrustGetNetworkFetchAllowedErr error

// CanCallSecTrustGetNetworkFetchAllowed reports whether the symbol for SecTrustGetNetworkFetchAllowed is available on this system.
func CanCallSecTrustGetNetworkFetchAllowed() bool {
	return _secTrustGetNetworkFetchAllowed != nil
}

// SecTrustGetNetworkFetchAllowedCallError returns the symbol lookup or registration error for SecTrustGetNetworkFetchAllowed.
func SecTrustGetNetworkFetchAllowedCallError() error {
	if _secTrustGetNetworkFetchAllowed != nil {
		return nil
	}
	return symbolCallError("SecTrustGetNetworkFetchAllowed", "10.9", _secTrustGetNetworkFetchAllowedErr)
}

// TrySecTrustGetNetworkFetchAllowed calls SecTrustGetNetworkFetchAllowed without panicking when the symbol is unavailable.
func TrySecTrustGetNetworkFetchAllowed(trust SecTrustRef, allowFetch *bool) (int32, error) {
	if err := SecTrustGetNetworkFetchAllowedCallError(); err != nil {
		return 0, err
	}
	return _secTrustGetNetworkFetchAllowed(trust, allowFetch), nil
}

// SecTrustGetNetworkFetchAllowed indicates whether a trust evaluation is permitted to fetch missing intermediate certificates from the network.
//
// See: https://developer.apple.com/documentation/Security/SecTrustGetNetworkFetchAllowed(_:_:)
func SecTrustGetNetworkFetchAllowed(trust SecTrustRef, allowFetch *bool) int32 {
	result, callErr := TrySecTrustGetNetworkFetchAllowed(trust, allowFetch)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustGetResult func(trustRef SecTrustRef, result *SecTrustResultType, certChain *corefoundation.CFArrayRef, statusChain unsafe.Pointer) int32
var _secTrustGetResultErr error

// CanCallSecTrustGetResult reports whether the symbol for SecTrustGetResult is available on this system.
func CanCallSecTrustGetResult() bool {
	return _secTrustGetResult != nil
}

// SecTrustGetResultCallError returns the symbol lookup or registration error for SecTrustGetResult.
func SecTrustGetResultCallError() error {
	if _secTrustGetResult != nil {
		return nil
	}
	return symbolCallError("SecTrustGetResult", "10.2", _secTrustGetResultErr)
}

// TrySecTrustGetResult calls SecTrustGetResult without panicking when the symbol is unavailable.
func TrySecTrustGetResult(trustRef SecTrustRef, result *SecTrustResultType, certChain *corefoundation.CFArrayRef, statusChain unsafe.Pointer) (int32, error) {
	if err := SecTrustGetResultCallError(); err != nil {
		return 0, err
	}
	return _secTrustGetResult(trustRef, result, certChain, statusChain), nil
}

// SecTrustGetResult retrieves details on the outcome of a call to the function [SecTrustEvaluate].
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecTrustGetResult
func SecTrustGetResult(trustRef SecTrustRef, result *SecTrustResultType, certChain *corefoundation.CFArrayRef, statusChain unsafe.Pointer) int32 {
	result0, callErr := TrySecTrustGetResult(trustRef, result, certChain, statusChain)
	if callErr != nil {
		panic(callErr)
	}
	return result0
}

var _secTrustGetTPHandle func(trust SecTrustRef, handle unsafe.Pointer) int32
var _secTrustGetTPHandleErr error

// CanCallSecTrustGetTPHandle reports whether the symbol for SecTrustGetTPHandle is available on this system.
func CanCallSecTrustGetTPHandle() bool {
	return _secTrustGetTPHandle != nil
}

// SecTrustGetTPHandleCallError returns the symbol lookup or registration error for SecTrustGetTPHandle.
func SecTrustGetTPHandleCallError() error {
	if _secTrustGetTPHandle != nil {
		return nil
	}
	return symbolCallError("SecTrustGetTPHandle", "10.2", _secTrustGetTPHandleErr)
}

// TrySecTrustGetTPHandle calls SecTrustGetTPHandle without panicking when the symbol is unavailable.
func TrySecTrustGetTPHandle(trust SecTrustRef, handle unsafe.Pointer) (int32, error) {
	if err := SecTrustGetTPHandleCallError(); err != nil {
		return 0, err
	}
	return _secTrustGetTPHandle(trust, handle), nil
}

// SecTrustGetTPHandle retrieves the trust policy handle.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecTrustGetTPHandle
func SecTrustGetTPHandle(trust SecTrustRef, handle unsafe.Pointer) int32 {
	result, callErr := TrySecTrustGetTPHandle(trust, handle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustGetTrustResult func(trust SecTrustRef, result *SecTrustResultType) int32
var _secTrustGetTrustResultErr error

// CanCallSecTrustGetTrustResult reports whether the symbol for SecTrustGetTrustResult is available on this system.
func CanCallSecTrustGetTrustResult() bool {
	return _secTrustGetTrustResult != nil
}

// SecTrustGetTrustResultCallError returns the symbol lookup or registration error for SecTrustGetTrustResult.
func SecTrustGetTrustResultCallError() error {
	if _secTrustGetTrustResult != nil {
		return nil
	}
	return symbolCallError("SecTrustGetTrustResult", "10.7", _secTrustGetTrustResultErr)
}

// TrySecTrustGetTrustResult calls SecTrustGetTrustResult without panicking when the symbol is unavailable.
func TrySecTrustGetTrustResult(trust SecTrustRef, result *SecTrustResultType) (int32, error) {
	if err := SecTrustGetTrustResultCallError(); err != nil {
		return 0, err
	}
	return _secTrustGetTrustResult(trust, result), nil
}

// SecTrustGetTrustResult returns the result code from the most recent trust evaluation.
//
// See: https://developer.apple.com/documentation/Security/SecTrustGetTrustResult(_:_:)
func SecTrustGetTrustResult(trust SecTrustRef, result *SecTrustResultType) int32 {
	result0, callErr := TrySecTrustGetTrustResult(trust, result)
	if callErr != nil {
		panic(callErr)
	}
	return result0
}

var _secTrustGetTypeID func() uint
var _secTrustGetTypeIDErr error

// CanCallSecTrustGetTypeID reports whether the symbol for SecTrustGetTypeID is available on this system.
func CanCallSecTrustGetTypeID() bool {
	return _secTrustGetTypeID != nil
}

// SecTrustGetTypeIDCallError returns the symbol lookup or registration error for SecTrustGetTypeID.
func SecTrustGetTypeIDCallError() error {
	if _secTrustGetTypeID != nil {
		return nil
	}
	return symbolCallError("SecTrustGetTypeID", "10.3", _secTrustGetTypeIDErr)
}

// TrySecTrustGetTypeID calls SecTrustGetTypeID without panicking when the symbol is unavailable.
func TrySecTrustGetTypeID() (uint, error) {
	if err := SecTrustGetTypeIDCallError(); err != nil {
		return 0, err
	}
	return _secTrustGetTypeID(), nil
}

// SecTrustGetTypeID returns the unique identifier of the opaque type to which a trust object belongs.
//
// See: https://developer.apple.com/documentation/Security/SecTrustGetTypeID()
func SecTrustGetTypeID() uint {
	result, callErr := TrySecTrustGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustGetVerifyTime func(trust SecTrustRef) corefoundation.CFAbsoluteTime
var _secTrustGetVerifyTimeErr error

// CanCallSecTrustGetVerifyTime reports whether the symbol for SecTrustGetVerifyTime is available on this system.
func CanCallSecTrustGetVerifyTime() bool {
	return _secTrustGetVerifyTime != nil
}

// SecTrustGetVerifyTimeCallError returns the symbol lookup or registration error for SecTrustGetVerifyTime.
func SecTrustGetVerifyTimeCallError() error {
	if _secTrustGetVerifyTime != nil {
		return nil
	}
	return symbolCallError("SecTrustGetVerifyTime", "10.6", _secTrustGetVerifyTimeErr)
}

// TrySecTrustGetVerifyTime calls SecTrustGetVerifyTime without panicking when the symbol is unavailable.
func TrySecTrustGetVerifyTime(trust SecTrustRef) (corefoundation.CFAbsoluteTime, error) {
	if err := SecTrustGetVerifyTimeCallError(); err != nil {
		return *new(corefoundation.CFAbsoluteTime), err
	}
	return _secTrustGetVerifyTime(trust), nil
}

// SecTrustGetVerifyTime gets the absolute time against which the certificates in a trust management object are verified.
//
// See: https://developer.apple.com/documentation/Security/SecTrustGetVerifyTime(_:)
func SecTrustGetVerifyTime(trust SecTrustRef) corefoundation.CFAbsoluteTime {
	result, callErr := TrySecTrustGetVerifyTime(trust)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustSetAnchorCertificates func(trust SecTrustRef, anchorCertificates corefoundation.CFArrayRef) int32
var _secTrustSetAnchorCertificatesErr error

// CanCallSecTrustSetAnchorCertificates reports whether the symbol for SecTrustSetAnchorCertificates is available on this system.
func CanCallSecTrustSetAnchorCertificates() bool {
	return _secTrustSetAnchorCertificates != nil
}

// SecTrustSetAnchorCertificatesCallError returns the symbol lookup or registration error for SecTrustSetAnchorCertificates.
func SecTrustSetAnchorCertificatesCallError() error {
	if _secTrustSetAnchorCertificates != nil {
		return nil
	}
	return symbolCallError("SecTrustSetAnchorCertificates", "10.3", _secTrustSetAnchorCertificatesErr)
}

// TrySecTrustSetAnchorCertificates calls SecTrustSetAnchorCertificates without panicking when the symbol is unavailable.
func TrySecTrustSetAnchorCertificates(trust SecTrustRef, anchorCertificates corefoundation.CFArrayRef) (int32, error) {
	if err := SecTrustSetAnchorCertificatesCallError(); err != nil {
		return 0, err
	}
	return _secTrustSetAnchorCertificates(trust, anchorCertificates), nil
}

// SecTrustSetAnchorCertificates sets the anchor certificates used when evaluating a trust management object.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSetAnchorCertificates(_:_:)
func SecTrustSetAnchorCertificates(trust SecTrustRef, anchorCertificates corefoundation.CFArrayRef) int32 {
	result, callErr := TrySecTrustSetAnchorCertificates(trust, anchorCertificates)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustSetAnchorCertificatesOnly func(trust SecTrustRef, anchorCertificatesOnly bool) int32
var _secTrustSetAnchorCertificatesOnlyErr error

// CanCallSecTrustSetAnchorCertificatesOnly reports whether the symbol for SecTrustSetAnchorCertificatesOnly is available on this system.
func CanCallSecTrustSetAnchorCertificatesOnly() bool {
	return _secTrustSetAnchorCertificatesOnly != nil
}

// SecTrustSetAnchorCertificatesOnlyCallError returns the symbol lookup or registration error for SecTrustSetAnchorCertificatesOnly.
func SecTrustSetAnchorCertificatesOnlyCallError() error {
	if _secTrustSetAnchorCertificatesOnly != nil {
		return nil
	}
	return symbolCallError("SecTrustSetAnchorCertificatesOnly", "10.6", _secTrustSetAnchorCertificatesOnlyErr)
}

// TrySecTrustSetAnchorCertificatesOnly calls SecTrustSetAnchorCertificatesOnly without panicking when the symbol is unavailable.
func TrySecTrustSetAnchorCertificatesOnly(trust SecTrustRef, anchorCertificatesOnly bool) (int32, error) {
	if err := SecTrustSetAnchorCertificatesOnlyCallError(); err != nil {
		return 0, err
	}
	return _secTrustSetAnchorCertificatesOnly(trust, anchorCertificatesOnly), nil
}

// SecTrustSetAnchorCertificatesOnly reenables trusting built-in anchor certificates.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSetAnchorCertificatesOnly(_:_:)
func SecTrustSetAnchorCertificatesOnly(trust SecTrustRef, anchorCertificatesOnly bool) int32 {
	result, callErr := TrySecTrustSetAnchorCertificatesOnly(trust, anchorCertificatesOnly)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustSetExceptions func(trust SecTrustRef, exceptions corefoundation.CFDataRef) bool
var _secTrustSetExceptionsErr error

// CanCallSecTrustSetExceptions reports whether the symbol for SecTrustSetExceptions is available on this system.
func CanCallSecTrustSetExceptions() bool {
	return _secTrustSetExceptions != nil
}

// SecTrustSetExceptionsCallError returns the symbol lookup or registration error for SecTrustSetExceptions.
func SecTrustSetExceptionsCallError() error {
	if _secTrustSetExceptions != nil {
		return nil
	}
	return symbolCallError("SecTrustSetExceptions", "10.9", _secTrustSetExceptionsErr)
}

// TrySecTrustSetExceptions calls SecTrustSetExceptions without panicking when the symbol is unavailable.
func TrySecTrustSetExceptions(trust SecTrustRef, exceptions corefoundation.CFDataRef) (bool, error) {
	if err := SecTrustSetExceptionsCallError(); err != nil {
		return false, err
	}
	return _secTrustSetExceptions(trust, exceptions), nil
}

// SecTrustSetExceptions sets a list of exceptions that should be ignored when the certificate is evaluated.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSetExceptions(_:_:)
func SecTrustSetExceptions(trust SecTrustRef, exceptions corefoundation.CFDataRef) bool {
	result, callErr := TrySecTrustSetExceptions(trust, exceptions)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustSetKeychains func(trust SecTrustRef, keychainOrArray corefoundation.CFTypeRef) int32
var _secTrustSetKeychainsErr error

// CanCallSecTrustSetKeychains reports whether the symbol for SecTrustSetKeychains is available on this system.
func CanCallSecTrustSetKeychains() bool {
	return _secTrustSetKeychains != nil
}

// SecTrustSetKeychainsCallError returns the symbol lookup or registration error for SecTrustSetKeychains.
func SecTrustSetKeychainsCallError() error {
	if _secTrustSetKeychains != nil {
		return nil
	}
	return symbolCallError("SecTrustSetKeychains", "10.3", _secTrustSetKeychainsErr)
}

// TrySecTrustSetKeychains calls SecTrustSetKeychains without panicking when the symbol is unavailable.
func TrySecTrustSetKeychains(trust SecTrustRef, keychainOrArray corefoundation.CFTypeRef) (int32, error) {
	if err := SecTrustSetKeychainsCallError(); err != nil {
		return 0, err
	}
	return _secTrustSetKeychains(trust, keychainOrArray), nil
}

// SecTrustSetKeychains sets the keychains searched for intermediate certificates when evaluating a trust management object.
//
// Deprecated: Deprecated since macOS 10.13.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSetKeychains(_:_:)
func SecTrustSetKeychains(trust SecTrustRef, keychainOrArray corefoundation.CFTypeRef) int32 {
	result, callErr := TrySecTrustSetKeychains(trust, keychainOrArray)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustSetNetworkFetchAllowed func(trust SecTrustRef, allowFetch bool) int32
var _secTrustSetNetworkFetchAllowedErr error

// CanCallSecTrustSetNetworkFetchAllowed reports whether the symbol for SecTrustSetNetworkFetchAllowed is available on this system.
func CanCallSecTrustSetNetworkFetchAllowed() bool {
	return _secTrustSetNetworkFetchAllowed != nil
}

// SecTrustSetNetworkFetchAllowedCallError returns the symbol lookup or registration error for SecTrustSetNetworkFetchAllowed.
func SecTrustSetNetworkFetchAllowedCallError() error {
	if _secTrustSetNetworkFetchAllowed != nil {
		return nil
	}
	return symbolCallError("SecTrustSetNetworkFetchAllowed", "10.9", _secTrustSetNetworkFetchAllowedErr)
}

// TrySecTrustSetNetworkFetchAllowed calls SecTrustSetNetworkFetchAllowed without panicking when the symbol is unavailable.
func TrySecTrustSetNetworkFetchAllowed(trust SecTrustRef, allowFetch bool) (int32, error) {
	if err := SecTrustSetNetworkFetchAllowedCallError(); err != nil {
		return 0, err
	}
	return _secTrustSetNetworkFetchAllowed(trust, allowFetch), nil
}

// SecTrustSetNetworkFetchAllowed specifies whether a trust evaluation is permitted to fetch missing intermediate certificates from the network.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSetNetworkFetchAllowed(_:_:)
func SecTrustSetNetworkFetchAllowed(trust SecTrustRef, allowFetch bool) int32 {
	result, callErr := TrySecTrustSetNetworkFetchAllowed(trust, allowFetch)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustSetOCSPResponse func(trust SecTrustRef, responseData corefoundation.CFTypeRef) int32
var _secTrustSetOCSPResponseErr error

// CanCallSecTrustSetOCSPResponse reports whether the symbol for SecTrustSetOCSPResponse is available on this system.
func CanCallSecTrustSetOCSPResponse() bool {
	return _secTrustSetOCSPResponse != nil
}

// SecTrustSetOCSPResponseCallError returns the symbol lookup or registration error for SecTrustSetOCSPResponse.
func SecTrustSetOCSPResponseCallError() error {
	if _secTrustSetOCSPResponse != nil {
		return nil
	}
	return symbolCallError("SecTrustSetOCSPResponse", "10.9", _secTrustSetOCSPResponseErr)
}

// TrySecTrustSetOCSPResponse calls SecTrustSetOCSPResponse without panicking when the symbol is unavailable.
func TrySecTrustSetOCSPResponse(trust SecTrustRef, responseData corefoundation.CFTypeRef) (int32, error) {
	if err := SecTrustSetOCSPResponseCallError(); err != nil {
		return 0, err
	}
	return _secTrustSetOCSPResponse(trust, responseData), nil
}

// SecTrustSetOCSPResponse attaches Online Certificate Status Protocol (OSCP) response data to a trust object.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSetOCSPResponse(_:_:)
func SecTrustSetOCSPResponse(trust SecTrustRef, responseData corefoundation.CFTypeRef) int32 {
	result, callErr := TrySecTrustSetOCSPResponse(trust, responseData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustSetOptions func(trustRef SecTrustRef, options SecTrustOptionFlags) int32
var _secTrustSetOptionsErr error

// CanCallSecTrustSetOptions reports whether the symbol for SecTrustSetOptions is available on this system.
func CanCallSecTrustSetOptions() bool {
	return _secTrustSetOptions != nil
}

// SecTrustSetOptionsCallError returns the symbol lookup or registration error for SecTrustSetOptions.
func SecTrustSetOptionsCallError() error {
	if _secTrustSetOptions != nil {
		return nil
	}
	return symbolCallError("SecTrustSetOptions", "10.7", _secTrustSetOptionsErr)
}

// TrySecTrustSetOptions calls SecTrustSetOptions without panicking when the symbol is unavailable.
func TrySecTrustSetOptions(trustRef SecTrustRef, options SecTrustOptionFlags) (int32, error) {
	if err := SecTrustSetOptionsCallError(); err != nil {
		return 0, err
	}
	return _secTrustSetOptions(trustRef, options), nil
}

// SecTrustSetOptions sets option flags for customizing evaluation of a trust object.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSetOptions(_:_:)
func SecTrustSetOptions(trustRef SecTrustRef, options SecTrustOptionFlags) int32 {
	result, callErr := TrySecTrustSetOptions(trustRef, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustSetParameters func(trustRef SecTrustRef, action CSSM_TP_ACTION, actionData corefoundation.CFDataRef) int32
var _secTrustSetParametersErr error

// CanCallSecTrustSetParameters reports whether the symbol for SecTrustSetParameters is available on this system.
func CanCallSecTrustSetParameters() bool {
	return _secTrustSetParameters != nil
}

// SecTrustSetParametersCallError returns the symbol lookup or registration error for SecTrustSetParameters.
func SecTrustSetParametersCallError() error {
	if _secTrustSetParameters != nil {
		return nil
	}
	return symbolCallError("SecTrustSetParameters", "10.2", _secTrustSetParametersErr)
}

// TrySecTrustSetParameters calls SecTrustSetParameters without panicking when the symbol is unavailable.
func TrySecTrustSetParameters(trustRef SecTrustRef, action CSSM_TP_ACTION, actionData corefoundation.CFDataRef) (int32, error) {
	if err := SecTrustSetParametersCallError(); err != nil {
		return 0, err
	}
	return _secTrustSetParameters(trustRef, action, actionData), nil
}

// SecTrustSetParameters sets the action and action data for a trust management object.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSetParameters
func SecTrustSetParameters(trustRef SecTrustRef, action CSSM_TP_ACTION, actionData corefoundation.CFDataRef) int32 {
	result, callErr := TrySecTrustSetParameters(trustRef, action, actionData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustSetPolicies func(trust SecTrustRef, policies corefoundation.CFTypeRef) int32
var _secTrustSetPoliciesErr error

// CanCallSecTrustSetPolicies reports whether the symbol for SecTrustSetPolicies is available on this system.
func CanCallSecTrustSetPolicies() bool {
	return _secTrustSetPolicies != nil
}

// SecTrustSetPoliciesCallError returns the symbol lookup or registration error for SecTrustSetPolicies.
func SecTrustSetPoliciesCallError() error {
	if _secTrustSetPolicies != nil {
		return nil
	}
	return symbolCallError("SecTrustSetPolicies", "10.3", _secTrustSetPoliciesErr)
}

// TrySecTrustSetPolicies calls SecTrustSetPolicies without panicking when the symbol is unavailable.
func TrySecTrustSetPolicies(trust SecTrustRef, policies corefoundation.CFTypeRef) (int32, error) {
	if err := SecTrustSetPoliciesCallError(); err != nil {
		return 0, err
	}
	return _secTrustSetPolicies(trust, policies), nil
}

// SecTrustSetPolicies sets the policies to use in an evaluation.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSetPolicies(_:_:)
func SecTrustSetPolicies(trust SecTrustRef, policies corefoundation.CFTypeRef) int32 {
	result, callErr := TrySecTrustSetPolicies(trust, policies)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustSetSignedCertificateTimestamps func(trust SecTrustRef, sctArray corefoundation.CFArrayRef) int32
var _secTrustSetSignedCertificateTimestampsErr error

// CanCallSecTrustSetSignedCertificateTimestamps reports whether the symbol for SecTrustSetSignedCertificateTimestamps is available on this system.
func CanCallSecTrustSetSignedCertificateTimestamps() bool {
	return _secTrustSetSignedCertificateTimestamps != nil
}

// SecTrustSetSignedCertificateTimestampsCallError returns the symbol lookup or registration error for SecTrustSetSignedCertificateTimestamps.
func SecTrustSetSignedCertificateTimestampsCallError() error {
	if _secTrustSetSignedCertificateTimestamps != nil {
		return nil
	}
	return symbolCallError("SecTrustSetSignedCertificateTimestamps", "10.14.2", _secTrustSetSignedCertificateTimestampsErr)
}

// TrySecTrustSetSignedCertificateTimestamps calls SecTrustSetSignedCertificateTimestamps without panicking when the symbol is unavailable.
func TrySecTrustSetSignedCertificateTimestamps(trust SecTrustRef, sctArray corefoundation.CFArrayRef) (int32, error) {
	if err := SecTrustSetSignedCertificateTimestampsCallError(); err != nil {
		return 0, err
	}
	return _secTrustSetSignedCertificateTimestamps(trust, sctArray), nil
}

// SecTrustSetSignedCertificateTimestamps attaches signed certificate timestamp data to a trust object.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSetSignedCertificateTimestamps(_:_:)
func SecTrustSetSignedCertificateTimestamps(trust SecTrustRef, sctArray corefoundation.CFArrayRef) int32 {
	result, callErr := TrySecTrustSetSignedCertificateTimestamps(trust, sctArray)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustSetVerifyDate func(trust SecTrustRef, verifyDate corefoundation.CFDateRef) int32
var _secTrustSetVerifyDateErr error

// CanCallSecTrustSetVerifyDate reports whether the symbol for SecTrustSetVerifyDate is available on this system.
func CanCallSecTrustSetVerifyDate() bool {
	return _secTrustSetVerifyDate != nil
}

// SecTrustSetVerifyDateCallError returns the symbol lookup or registration error for SecTrustSetVerifyDate.
func SecTrustSetVerifyDateCallError() error {
	if _secTrustSetVerifyDate != nil {
		return nil
	}
	return symbolCallError("SecTrustSetVerifyDate", "10.3", _secTrustSetVerifyDateErr)
}

// TrySecTrustSetVerifyDate calls SecTrustSetVerifyDate without panicking when the symbol is unavailable.
func TrySecTrustSetVerifyDate(trust SecTrustRef, verifyDate corefoundation.CFDateRef) (int32, error) {
	if err := SecTrustSetVerifyDateCallError(); err != nil {
		return 0, err
	}
	return _secTrustSetVerifyDate(trust, verifyDate), nil
}

// SecTrustSetVerifyDate sets the date and time against which the certificates in a trust management object are verified.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSetVerifyDate(_:_:)
func SecTrustSetVerifyDate(trust SecTrustRef, verifyDate corefoundation.CFDateRef) int32 {
	result, callErr := TrySecTrustSetVerifyDate(trust, verifyDate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustSettingsCopyCertificates func(domain SecTrustSettingsDomain, certArray *corefoundation.CFArrayRef) int32
var _secTrustSettingsCopyCertificatesErr error

// CanCallSecTrustSettingsCopyCertificates reports whether the symbol for SecTrustSettingsCopyCertificates is available on this system.
func CanCallSecTrustSettingsCopyCertificates() bool {
	return _secTrustSettingsCopyCertificates != nil
}

// SecTrustSettingsCopyCertificatesCallError returns the symbol lookup or registration error for SecTrustSettingsCopyCertificates.
func SecTrustSettingsCopyCertificatesCallError() error {
	if _secTrustSettingsCopyCertificates != nil {
		return nil
	}
	return symbolCallError("SecTrustSettingsCopyCertificates", "10.0", _secTrustSettingsCopyCertificatesErr)
}

// TrySecTrustSettingsCopyCertificates calls SecTrustSettingsCopyCertificates without panicking when the symbol is unavailable.
func TrySecTrustSettingsCopyCertificates(domain SecTrustSettingsDomain, certArray *corefoundation.CFArrayRef) (int32, error) {
	if err := SecTrustSettingsCopyCertificatesCallError(); err != nil {
		return 0, err
	}
	return _secTrustSettingsCopyCertificates(domain, certArray), nil
}

// SecTrustSettingsCopyCertificates obtains an array of all certificates that have trust settings in a specific trust settings domain.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSettingsCopyCertificates(_:_:)
func SecTrustSettingsCopyCertificates(domain SecTrustSettingsDomain, certArray *corefoundation.CFArrayRef) int32 {
	result, callErr := TrySecTrustSettingsCopyCertificates(domain, certArray)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustSettingsCopyModificationDate func(certRef SecCertificateRef, domain SecTrustSettingsDomain, modificationDate *corefoundation.CFDateRef) int32
var _secTrustSettingsCopyModificationDateErr error

// CanCallSecTrustSettingsCopyModificationDate reports whether the symbol for SecTrustSettingsCopyModificationDate is available on this system.
func CanCallSecTrustSettingsCopyModificationDate() bool {
	return _secTrustSettingsCopyModificationDate != nil
}

// SecTrustSettingsCopyModificationDateCallError returns the symbol lookup or registration error for SecTrustSettingsCopyModificationDate.
func SecTrustSettingsCopyModificationDateCallError() error {
	if _secTrustSettingsCopyModificationDate != nil {
		return nil
	}
	return symbolCallError("SecTrustSettingsCopyModificationDate", "10.0", _secTrustSettingsCopyModificationDateErr)
}

// TrySecTrustSettingsCopyModificationDate calls SecTrustSettingsCopyModificationDate without panicking when the symbol is unavailable.
func TrySecTrustSettingsCopyModificationDate(certRef SecCertificateRef, domain SecTrustSettingsDomain, modificationDate *corefoundation.CFDateRef) (int32, error) {
	if err := SecTrustSettingsCopyModificationDateCallError(); err != nil {
		return 0, err
	}
	return _secTrustSettingsCopyModificationDate(certRef, domain, modificationDate), nil
}

// SecTrustSettingsCopyModificationDate obtains the date and time at which a certificate’s trust settings were last modified.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSettingsCopyModificationDate(_:_:_:)
func SecTrustSettingsCopyModificationDate(certRef SecCertificateRef, domain SecTrustSettingsDomain, modificationDate *corefoundation.CFDateRef) int32 {
	result, callErr := TrySecTrustSettingsCopyModificationDate(certRef, domain, modificationDate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustSettingsCopyTrustSettings func(certRef SecCertificateRef, domain SecTrustSettingsDomain, trustSettings *corefoundation.CFArrayRef) int32
var _secTrustSettingsCopyTrustSettingsErr error

// CanCallSecTrustSettingsCopyTrustSettings reports whether the symbol for SecTrustSettingsCopyTrustSettings is available on this system.
func CanCallSecTrustSettingsCopyTrustSettings() bool {
	return _secTrustSettingsCopyTrustSettings != nil
}

// SecTrustSettingsCopyTrustSettingsCallError returns the symbol lookup or registration error for SecTrustSettingsCopyTrustSettings.
func SecTrustSettingsCopyTrustSettingsCallError() error {
	if _secTrustSettingsCopyTrustSettings != nil {
		return nil
	}
	return symbolCallError("SecTrustSettingsCopyTrustSettings", "10.0", _secTrustSettingsCopyTrustSettingsErr)
}

// TrySecTrustSettingsCopyTrustSettings calls SecTrustSettingsCopyTrustSettings without panicking when the symbol is unavailable.
func TrySecTrustSettingsCopyTrustSettings(certRef SecCertificateRef, domain SecTrustSettingsDomain, trustSettings *corefoundation.CFArrayRef) (int32, error) {
	if err := SecTrustSettingsCopyTrustSettingsCallError(); err != nil {
		return 0, err
	}
	return _secTrustSettingsCopyTrustSettings(certRef, domain, trustSettings), nil
}

// SecTrustSettingsCopyTrustSettings obtains the trust settings for a certificate.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSettingsCopyTrustSettings(_:_:_:)
func SecTrustSettingsCopyTrustSettings(certRef SecCertificateRef, domain SecTrustSettingsDomain, trustSettings *corefoundation.CFArrayRef) int32 {
	result, callErr := TrySecTrustSettingsCopyTrustSettings(certRef, domain, trustSettings)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustSettingsCreateExternalRepresentation func(domain SecTrustSettingsDomain, trustSettings *corefoundation.CFDataRef) int32
var _secTrustSettingsCreateExternalRepresentationErr error

// CanCallSecTrustSettingsCreateExternalRepresentation reports whether the symbol for SecTrustSettingsCreateExternalRepresentation is available on this system.
func CanCallSecTrustSettingsCreateExternalRepresentation() bool {
	return _secTrustSettingsCreateExternalRepresentation != nil
}

// SecTrustSettingsCreateExternalRepresentationCallError returns the symbol lookup or registration error for SecTrustSettingsCreateExternalRepresentation.
func SecTrustSettingsCreateExternalRepresentationCallError() error {
	if _secTrustSettingsCreateExternalRepresentation != nil {
		return nil
	}
	return symbolCallError("SecTrustSettingsCreateExternalRepresentation", "10.0", _secTrustSettingsCreateExternalRepresentationErr)
}

// TrySecTrustSettingsCreateExternalRepresentation calls SecTrustSettingsCreateExternalRepresentation without panicking when the symbol is unavailable.
func TrySecTrustSettingsCreateExternalRepresentation(domain SecTrustSettingsDomain, trustSettings *corefoundation.CFDataRef) (int32, error) {
	if err := SecTrustSettingsCreateExternalRepresentationCallError(); err != nil {
		return 0, err
	}
	return _secTrustSettingsCreateExternalRepresentation(domain, trustSettings), nil
}

// SecTrustSettingsCreateExternalRepresentation obtains an external, portable representation of the specified domain’s trust settings.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSettingsCreateExternalRepresentation(_:_:)
func SecTrustSettingsCreateExternalRepresentation(domain SecTrustSettingsDomain, trustSettings *corefoundation.CFDataRef) int32 {
	result, callErr := TrySecTrustSettingsCreateExternalRepresentation(domain, trustSettings)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustSettingsImportExternalRepresentation func(domain SecTrustSettingsDomain, trustSettings corefoundation.CFDataRef) int32
var _secTrustSettingsImportExternalRepresentationErr error

// CanCallSecTrustSettingsImportExternalRepresentation reports whether the symbol for SecTrustSettingsImportExternalRepresentation is available on this system.
func CanCallSecTrustSettingsImportExternalRepresentation() bool {
	return _secTrustSettingsImportExternalRepresentation != nil
}

// SecTrustSettingsImportExternalRepresentationCallError returns the symbol lookup or registration error for SecTrustSettingsImportExternalRepresentation.
func SecTrustSettingsImportExternalRepresentationCallError() error {
	if _secTrustSettingsImportExternalRepresentation != nil {
		return nil
	}
	return symbolCallError("SecTrustSettingsImportExternalRepresentation", "10.0", _secTrustSettingsImportExternalRepresentationErr)
}

// TrySecTrustSettingsImportExternalRepresentation calls SecTrustSettingsImportExternalRepresentation without panicking when the symbol is unavailable.
func TrySecTrustSettingsImportExternalRepresentation(domain SecTrustSettingsDomain, trustSettings corefoundation.CFDataRef) (int32, error) {
	if err := SecTrustSettingsImportExternalRepresentationCallError(); err != nil {
		return 0, err
	}
	return _secTrustSettingsImportExternalRepresentation(domain, trustSettings), nil
}

// SecTrustSettingsImportExternalRepresentation imports trust settings into a trust domain.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSettingsImportExternalRepresentation(_:_:)
func SecTrustSettingsImportExternalRepresentation(domain SecTrustSettingsDomain, trustSettings corefoundation.CFDataRef) int32 {
	result, callErr := TrySecTrustSettingsImportExternalRepresentation(domain, trustSettings)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustSettingsRemoveTrustSettings func(certRef SecCertificateRef, domain SecTrustSettingsDomain) int32
var _secTrustSettingsRemoveTrustSettingsErr error

// CanCallSecTrustSettingsRemoveTrustSettings reports whether the symbol for SecTrustSettingsRemoveTrustSettings is available on this system.
func CanCallSecTrustSettingsRemoveTrustSettings() bool {
	return _secTrustSettingsRemoveTrustSettings != nil
}

// SecTrustSettingsRemoveTrustSettingsCallError returns the symbol lookup or registration error for SecTrustSettingsRemoveTrustSettings.
func SecTrustSettingsRemoveTrustSettingsCallError() error {
	if _secTrustSettingsRemoveTrustSettings != nil {
		return nil
	}
	return symbolCallError("SecTrustSettingsRemoveTrustSettings", "10.0", _secTrustSettingsRemoveTrustSettingsErr)
}

// TrySecTrustSettingsRemoveTrustSettings calls SecTrustSettingsRemoveTrustSettings without panicking when the symbol is unavailable.
func TrySecTrustSettingsRemoveTrustSettings(certRef SecCertificateRef, domain SecTrustSettingsDomain) (int32, error) {
	if err := SecTrustSettingsRemoveTrustSettingsCallError(); err != nil {
		return 0, err
	}
	return _secTrustSettingsRemoveTrustSettings(certRef, domain), nil
}

// SecTrustSettingsRemoveTrustSettings deletes the trust settings for a certificate.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSettingsRemoveTrustSettings(_:_:)
func SecTrustSettingsRemoveTrustSettings(certRef SecCertificateRef, domain SecTrustSettingsDomain) int32 {
	result, callErr := TrySecTrustSettingsRemoveTrustSettings(certRef, domain)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secTrustSettingsSetTrustSettings func(certRef SecCertificateRef, domain SecTrustSettingsDomain, trustSettingsDictOrArray corefoundation.CFTypeRef) int32
var _secTrustSettingsSetTrustSettingsErr error

// CanCallSecTrustSettingsSetTrustSettings reports whether the symbol for SecTrustSettingsSetTrustSettings is available on this system.
func CanCallSecTrustSettingsSetTrustSettings() bool {
	return _secTrustSettingsSetTrustSettings != nil
}

// SecTrustSettingsSetTrustSettingsCallError returns the symbol lookup or registration error for SecTrustSettingsSetTrustSettings.
func SecTrustSettingsSetTrustSettingsCallError() error {
	if _secTrustSettingsSetTrustSettings != nil {
		return nil
	}
	return symbolCallError("SecTrustSettingsSetTrustSettings", "10.0", _secTrustSettingsSetTrustSettingsErr)
}

// TrySecTrustSettingsSetTrustSettings calls SecTrustSettingsSetTrustSettings without panicking when the symbol is unavailable.
func TrySecTrustSettingsSetTrustSettings(certRef SecCertificateRef, domain SecTrustSettingsDomain, trustSettingsDictOrArray corefoundation.CFTypeRef) (int32, error) {
	if err := SecTrustSettingsSetTrustSettingsCallError(); err != nil {
		return 0, err
	}
	return _secTrustSettingsSetTrustSettings(certRef, domain, trustSettingsDictOrArray), nil
}

// SecTrustSettingsSetTrustSettings specifies trust settings for a certificate.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSettingsSetTrustSettings(_:_:_:)
func SecTrustSettingsSetTrustSettings(certRef SecCertificateRef, domain SecTrustSettingsDomain, trustSettingsDictOrArray corefoundation.CFTypeRef) int32 {
	result, callErr := TrySecTrustSettingsSetTrustSettings(certRef, domain, trustSettingsDictOrArray)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secureDownloadCopyCreationDate func(downloadRef SecureDownloadRef, date *corefoundation.CFDateRef) int32
var _secureDownloadCopyCreationDateErr error

// CanCallSecureDownloadCopyCreationDate reports whether the symbol for SecureDownloadCopyCreationDate is available on this system.
func CanCallSecureDownloadCopyCreationDate() bool {
	return _secureDownloadCopyCreationDate != nil
}

// SecureDownloadCopyCreationDateCallError returns the symbol lookup or registration error for SecureDownloadCopyCreationDate.
func SecureDownloadCopyCreationDateCallError() error {
	if _secureDownloadCopyCreationDate != nil {
		return nil
	}
	return symbolCallError("SecureDownloadCopyCreationDate", "10.5", _secureDownloadCopyCreationDateErr)
}

// TrySecureDownloadCopyCreationDate calls SecureDownloadCopyCreationDate without panicking when the symbol is unavailable.
func TrySecureDownloadCopyCreationDate(downloadRef SecureDownloadRef, date *corefoundation.CFDateRef) (int32, error) {
	if err := SecureDownloadCopyCreationDateCallError(); err != nil {
		return 0, err
	}
	return _secureDownloadCopyCreationDate(downloadRef, date), nil
}

// SecureDownloadCopyCreationDate returns download ticket’s creation date.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecureDownloadCopyCreationDate
func SecureDownloadCopyCreationDate(downloadRef SecureDownloadRef, date *corefoundation.CFDateRef) int32 {
	result, callErr := TrySecureDownloadCopyCreationDate(downloadRef, date)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secureDownloadCopyName func(downloadRef SecureDownloadRef, name *corefoundation.CFStringRef) int32
var _secureDownloadCopyNameErr error

// CanCallSecureDownloadCopyName reports whether the symbol for SecureDownloadCopyName is available on this system.
func CanCallSecureDownloadCopyName() bool {
	return _secureDownloadCopyName != nil
}

// SecureDownloadCopyNameCallError returns the symbol lookup or registration error for SecureDownloadCopyName.
func SecureDownloadCopyNameCallError() error {
	if _secureDownloadCopyName != nil {
		return nil
	}
	return symbolCallError("SecureDownloadCopyName", "10.5", _secureDownloadCopyNameErr)
}

// TrySecureDownloadCopyName calls SecureDownloadCopyName without panicking when the symbol is unavailable.
func TrySecureDownloadCopyName(downloadRef SecureDownloadRef, name *corefoundation.CFStringRef) (int32, error) {
	if err := SecureDownloadCopyNameCallError(); err != nil {
		return 0, err
	}
	return _secureDownloadCopyName(downloadRef, name), nil
}

// SecureDownloadCopyName returns the printable name of the download ticket.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecureDownloadCopyName
func SecureDownloadCopyName(downloadRef SecureDownloadRef, name *corefoundation.CFStringRef) int32 {
	result, callErr := TrySecureDownloadCopyName(downloadRef, name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secureDownloadCopyTicketLocation func(url corefoundation.CFURLRef, ticketLocation *corefoundation.CFURLRef) int32
var _secureDownloadCopyTicketLocationErr error

// CanCallSecureDownloadCopyTicketLocation reports whether the symbol for SecureDownloadCopyTicketLocation is available on this system.
func CanCallSecureDownloadCopyTicketLocation() bool {
	return _secureDownloadCopyTicketLocation != nil
}

// SecureDownloadCopyTicketLocationCallError returns the symbol lookup or registration error for SecureDownloadCopyTicketLocation.
func SecureDownloadCopyTicketLocationCallError() error {
	if _secureDownloadCopyTicketLocation != nil {
		return nil
	}
	return symbolCallError("SecureDownloadCopyTicketLocation", "10.5", _secureDownloadCopyTicketLocationErr)
}

// TrySecureDownloadCopyTicketLocation calls SecureDownloadCopyTicketLocation without panicking when the symbol is unavailable.
func TrySecureDownloadCopyTicketLocation(url corefoundation.CFURLRef, ticketLocation *corefoundation.CFURLRef) (int32, error) {
	if err := SecureDownloadCopyTicketLocationCallError(); err != nil {
		return 0, err
	}
	return _secureDownloadCopyTicketLocation(url, ticketLocation), nil
}

// SecureDownloadCopyTicketLocation copies the ticket location from a secure download URL.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecureDownloadCopyTicketLocation
func SecureDownloadCopyTicketLocation(url corefoundation.CFURLRef, ticketLocation *corefoundation.CFURLRef) int32 {
	result, callErr := TrySecureDownloadCopyTicketLocation(url, ticketLocation)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secureDownloadCopyURLs func(downloadRef SecureDownloadRef, urls *corefoundation.CFArrayRef) int32
var _secureDownloadCopyURLsErr error

// CanCallSecureDownloadCopyURLs reports whether the symbol for SecureDownloadCopyURLs is available on this system.
func CanCallSecureDownloadCopyURLs() bool {
	return _secureDownloadCopyURLs != nil
}

// SecureDownloadCopyURLsCallError returns the symbol lookup or registration error for SecureDownloadCopyURLs.
func SecureDownloadCopyURLsCallError() error {
	if _secureDownloadCopyURLs != nil {
		return nil
	}
	return symbolCallError("SecureDownloadCopyURLs", "10.5", _secureDownloadCopyURLsErr)
}

// TrySecureDownloadCopyURLs calls SecureDownloadCopyURLs without panicking when the symbol is unavailable.
func TrySecureDownloadCopyURLs(downloadRef SecureDownloadRef, urls *corefoundation.CFArrayRef) (int32, error) {
	if err := SecureDownloadCopyURLsCallError(); err != nil {
		return 0, err
	}
	return _secureDownloadCopyURLs(downloadRef, urls), nil
}

// SecureDownloadCopyURLs returns a list of URLs from which the data can be downloaded.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecureDownloadCopyURLs
func SecureDownloadCopyURLs(downloadRef SecureDownloadRef, urls *corefoundation.CFArrayRef) int32 {
	result, callErr := TrySecureDownloadCopyURLs(downloadRef, urls)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secureDownloadCreateWithTicket func(ticket corefoundation.CFDataRef, setup unsafe.Pointer, setupContext unsafe.Pointer, evaluate unsafe.Pointer, evaluateContext unsafe.Pointer, downloadRef *SecureDownloadRef) int32
var _secureDownloadCreateWithTicketErr error

// CanCallSecureDownloadCreateWithTicket reports whether the symbol for SecureDownloadCreateWithTicket is available on this system.
func CanCallSecureDownloadCreateWithTicket() bool {
	return _secureDownloadCreateWithTicket != nil
}

// SecureDownloadCreateWithTicketCallError returns the symbol lookup or registration error for SecureDownloadCreateWithTicket.
func SecureDownloadCreateWithTicketCallError() error {
	if _secureDownloadCreateWithTicket != nil {
		return nil
	}
	return symbolCallError("SecureDownloadCreateWithTicket", "10.5", _secureDownloadCreateWithTicketErr)
}

// TrySecureDownloadCreateWithTicket calls SecureDownloadCreateWithTicket without panicking when the symbol is unavailable.
func TrySecureDownloadCreateWithTicket(ticket corefoundation.CFDataRef, setup unsafe.Pointer, setupContext unsafe.Pointer, evaluate unsafe.Pointer, evaluateContext unsafe.Pointer, downloadRef *SecureDownloadRef) (int32, error) {
	if err := SecureDownloadCreateWithTicketCallError(); err != nil {
		return 0, err
	}
	return _secureDownloadCreateWithTicket(ticket, setup, setupContext, evaluate, evaluateContext, downloadRef), nil
}

// SecureDownloadCreateWithTicket creates a secure download object for use during the download process.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecureDownloadCreateWithTicket
func SecureDownloadCreateWithTicket(ticket corefoundation.CFDataRef, setup unsafe.Pointer, setupContext unsafe.Pointer, evaluate unsafe.Pointer, evaluateContext unsafe.Pointer, downloadRef *SecureDownloadRef) int32 {
	result, callErr := TrySecureDownloadCreateWithTicket(ticket, setup, setupContext, evaluate, evaluateContext, downloadRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secureDownloadFinished func(downloadRef SecureDownloadRef) int32
var _secureDownloadFinishedErr error

// CanCallSecureDownloadFinished reports whether the symbol for SecureDownloadFinished is available on this system.
func CanCallSecureDownloadFinished() bool {
	return _secureDownloadFinished != nil
}

// SecureDownloadFinishedCallError returns the symbol lookup or registration error for SecureDownloadFinished.
func SecureDownloadFinishedCallError() error {
	if _secureDownloadFinished != nil {
		return nil
	}
	return symbolCallError("SecureDownloadFinished", "10.5", _secureDownloadFinishedErr)
}

// TrySecureDownloadFinished calls SecureDownloadFinished without panicking when the symbol is unavailable.
func TrySecureDownloadFinished(downloadRef SecureDownloadRef) (int32, error) {
	if err := SecureDownloadFinishedCallError(); err != nil {
		return 0, err
	}
	return _secureDownloadFinished(downloadRef), nil
}

// SecureDownloadFinished concludes the secure download process.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecureDownloadFinished
func SecureDownloadFinished(downloadRef SecureDownloadRef) int32 {
	result, callErr := TrySecureDownloadFinished(downloadRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secureDownloadGetDownloadSize func(downloadRef SecureDownloadRef, downloadSize *int64) int32
var _secureDownloadGetDownloadSizeErr error

// CanCallSecureDownloadGetDownloadSize reports whether the symbol for SecureDownloadGetDownloadSize is available on this system.
func CanCallSecureDownloadGetDownloadSize() bool {
	return _secureDownloadGetDownloadSize != nil
}

// SecureDownloadGetDownloadSizeCallError returns the symbol lookup or registration error for SecureDownloadGetDownloadSize.
func SecureDownloadGetDownloadSizeCallError() error {
	if _secureDownloadGetDownloadSize != nil {
		return nil
	}
	return symbolCallError("SecureDownloadGetDownloadSize", "10.5", _secureDownloadGetDownloadSizeErr)
}

// TrySecureDownloadGetDownloadSize calls SecureDownloadGetDownloadSize without panicking when the symbol is unavailable.
func TrySecureDownloadGetDownloadSize(downloadRef SecureDownloadRef, downloadSize *int64) (int32, error) {
	if err := SecureDownloadGetDownloadSizeCallError(); err != nil {
		return 0, err
	}
	return _secureDownloadGetDownloadSize(downloadRef, downloadSize), nil
}

// SecureDownloadGetDownloadSize returns the size of the expected download.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecureDownloadGetDownloadSize
func SecureDownloadGetDownloadSize(downloadRef SecureDownloadRef, downloadSize *int64) int32 {
	result, callErr := TrySecureDownloadGetDownloadSize(downloadRef, downloadSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secureDownloadRelease func(downloadRef SecureDownloadRef) int32
var _secureDownloadReleaseErr error

// CanCallSecureDownloadRelease reports whether the symbol for SecureDownloadRelease is available on this system.
func CanCallSecureDownloadRelease() bool {
	return _secureDownloadRelease != nil
}

// SecureDownloadReleaseCallError returns the symbol lookup or registration error for SecureDownloadRelease.
func SecureDownloadReleaseCallError() error {
	if _secureDownloadRelease != nil {
		return nil
	}
	return symbolCallError("SecureDownloadRelease", "10.5", _secureDownloadReleaseErr)
}

// TrySecureDownloadRelease calls SecureDownloadRelease without panicking when the symbol is unavailable.
func TrySecureDownloadRelease(downloadRef SecureDownloadRef) (int32, error) {
	if err := SecureDownloadReleaseCallError(); err != nil {
		return 0, err
	}
	return _secureDownloadRelease(downloadRef), nil
}

// SecureDownloadRelease releases the memory associated with a secure download object.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecureDownloadRelease
func SecureDownloadRelease(downloadRef SecureDownloadRef) int32 {
	result, callErr := TrySecureDownloadRelease(downloadRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _secureDownloadUpdateWithData func(downloadRef SecureDownloadRef, data corefoundation.CFDataRef) int32
var _secureDownloadUpdateWithDataErr error

// CanCallSecureDownloadUpdateWithData reports whether the symbol for SecureDownloadUpdateWithData is available on this system.
func CanCallSecureDownloadUpdateWithData() bool {
	return _secureDownloadUpdateWithData != nil
}

// SecureDownloadUpdateWithDataCallError returns the symbol lookup or registration error for SecureDownloadUpdateWithData.
func SecureDownloadUpdateWithDataCallError() error {
	if _secureDownloadUpdateWithData != nil {
		return nil
	}
	return symbolCallError("SecureDownloadUpdateWithData", "10.5", _secureDownloadUpdateWithDataErr)
}

// TrySecureDownloadUpdateWithData calls SecureDownloadUpdateWithData without panicking when the symbol is unavailable.
func TrySecureDownloadUpdateWithData(downloadRef SecureDownloadRef, data corefoundation.CFDataRef) (int32, error) {
	if err := SecureDownloadUpdateWithDataCallError(); err != nil {
		return 0, err
	}
	return _secureDownloadUpdateWithData(downloadRef, data), nil
}

// SecureDownloadUpdateWithData checks data received during download for validity.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/Security/SecureDownloadUpdateWithData
func SecureDownloadUpdateWithData(downloadRef SecureDownloadRef, data corefoundation.CFDataRef) int32 {
	result, callErr := TrySecureDownloadUpdateWithData(downloadRef, data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sessionCreate func(flags SessionCreationFlags, attributes SessionAttributeBits) int32
var _sessionCreateErr error

// CanCallSessionCreate reports whether the symbol for SessionCreate is available on this system.
func CanCallSessionCreate() bool {
	return _sessionCreate != nil
}

// SessionCreateCallError returns the symbol lookup or registration error for SessionCreate.
func SessionCreateCallError() error {
	if _sessionCreate != nil {
		return nil
	}
	return symbolCallError("SessionCreate", "10.0", _sessionCreateErr)
}

// TrySessionCreate calls SessionCreate without panicking when the symbol is unavailable.
func TrySessionCreate(flags SessionCreationFlags, attributes SessionAttributeBits) (int32, error) {
	if err := SessionCreateCallError(); err != nil {
		return 0, err
	}
	return _sessionCreate(flags, attributes), nil
}

// SessionCreate creates a security session.
//
// See: https://developer.apple.com/documentation/Security/SessionCreate(_:_:)
func SessionCreate(flags SessionCreationFlags, attributes SessionAttributeBits) int32 {
	result, callErr := TrySessionCreate(flags, attributes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sessionGetInfo func(session SecuritySessionId, sessionId *SecuritySessionId, attributes *SessionAttributeBits) int32
var _sessionGetInfoErr error

// CanCallSessionGetInfo reports whether the symbol for SessionGetInfo is available on this system.
func CanCallSessionGetInfo() bool {
	return _sessionGetInfo != nil
}

// SessionGetInfoCallError returns the symbol lookup or registration error for SessionGetInfo.
func SessionGetInfoCallError() error {
	if _sessionGetInfo != nil {
		return nil
	}
	return symbolCallError("SessionGetInfo", "10.0", _sessionGetInfoErr)
}

// TrySessionGetInfo calls SessionGetInfo without panicking when the symbol is unavailable.
func TrySessionGetInfo(session SecuritySessionId, sessionId *SecuritySessionId, attributes *SessionAttributeBits) (int32, error) {
	if err := SessionGetInfoCallError(); err != nil {
		return 0, err
	}
	return _sessionGetInfo(session, sessionId, attributes), nil
}

// SessionGetInfo obtains information about a security session.
//
// See: https://developer.apple.com/documentation/Security/SessionGetInfo(_:_:_:)
func SessionGetInfo(session SecuritySessionId, sessionId *SecuritySessionId, attributes *SessionAttributeBits) int32 {
	result, callErr := TrySessionGetInfo(session, sessionId, attributes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cssmAlgToOid func(algId CSSM_ALGORITHMS) unsafe.Pointer
var _cssmAlgToOidErr error

// CanCallCssmAlgToOid reports whether the symbol for cssmAlgToOid is available on this system.
func CanCallCssmAlgToOid() bool {
	return _cssmAlgToOid != nil
}

// CssmAlgToOidCallError returns the symbol lookup or registration error for cssmAlgToOid.
func CssmAlgToOidCallError() error {
	if _cssmAlgToOid != nil {
		return nil
	}
	return symbolCallError("cssmAlgToOid", "10.0", _cssmAlgToOidErr)
}

// TryCssmAlgToOid calls CssmAlgToOid without panicking when the symbol is unavailable.
func TryCssmAlgToOid(algId CSSM_ALGORITHMS) (unsafe.Pointer, error) {
	if err := CssmAlgToOidCallError(); err != nil {
		return nil, err
	}
	return _cssmAlgToOid(algId), nil
}

// CssmAlgToOid.
//
// See: https://developer.apple.com/documentation/Security/cssmAlgToOid(_:)
func CssmAlgToOid(algId CSSM_ALGORITHMS) unsafe.Pointer {
	result, callErr := TryCssmAlgToOid(algId)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cssmOidToAlg func(oid unsafe.Pointer, alg unsafe.Pointer) bool
var _cssmOidToAlgErr error

// CanCallCssmOidToAlg reports whether the symbol for cssmOidToAlg is available on this system.
func CanCallCssmOidToAlg() bool {
	return _cssmOidToAlg != nil
}

// CssmOidToAlgCallError returns the symbol lookup or registration error for cssmOidToAlg.
func CssmOidToAlgCallError() error {
	if _cssmOidToAlg != nil {
		return nil
	}
	return symbolCallError("cssmOidToAlg", "10.0", _cssmOidToAlgErr)
}

// TryCssmOidToAlg calls CssmOidToAlg without panicking when the symbol is unavailable.
func TryCssmOidToAlg(oid unsafe.Pointer, alg unsafe.Pointer) (bool, error) {
	if err := CssmOidToAlgCallError(); err != nil {
		return false, err
	}
	return _cssmOidToAlg(oid, alg), nil
}

// CssmOidToAlg.
//
// See: https://developer.apple.com/documentation/Security/cssmOidToAlg(_:_:)
func CssmOidToAlg(oid unsafe.Pointer, alg unsafe.Pointer) bool {
	result, callErr := TryCssmOidToAlg(oid, alg)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cssmPerror func(how string, err CSSM_RETURN)
var _cssmPerrorErr error

// CanCallCssmPerror reports whether the symbol for cssmPerror is available on this system.
func CanCallCssmPerror() bool {
	return _cssmPerror != nil
}

// CssmPerrorCallError returns the symbol lookup or registration error for cssmPerror.
func CssmPerrorCallError() error {
	if _cssmPerror != nil {
		return nil
	}
	return symbolCallError("cssmPerror", "10.0", _cssmPerrorErr)
}

// TryCssmPerror calls CssmPerror without panicking when the symbol is unavailable.
func TryCssmPerror(how string, err CSSM_RETURN) error {
	if err := CssmPerrorCallError(); err != nil {
		return err
	}
	_cssmPerror(how, err)
	return nil
}

// CssmPerror.
//
// See: https://developer.apple.com/documentation/Security/cssmPerror(_:_:)
func CssmPerror(how string, err CSSM_RETURN) {
	if callErr := TryCssmPerror(how, err); callErr != nil {
		panic(callErr)
	}
}

var _sec_certificate_copy_ref func(certificate Sec_certificate_t) SecCertificateRef
var _sec_certificate_copy_refErr error

// CanCallSec_certificate_copy_ref reports whether the symbol for sec_certificate_copy_ref is available on this system.
func CanCallSec_certificate_copy_ref() bool {
	return _sec_certificate_copy_ref != nil
}

// Sec_certificate_copy_refCallError returns the symbol lookup or registration error for sec_certificate_copy_ref.
func Sec_certificate_copy_refCallError() error {
	if _sec_certificate_copy_ref != nil {
		return nil
	}
	return symbolCallError("sec_certificate_copy_ref", "10.14", _sec_certificate_copy_refErr)
}

// TrySec_certificate_copy_ref calls Sec_certificate_copy_ref without panicking when the symbol is unavailable.
func TrySec_certificate_copy_ref(certificate Sec_certificate_t) (SecCertificateRef, error) {
	if err := Sec_certificate_copy_refCallError(); err != nil {
		return 0, err
	}
	return _sec_certificate_copy_ref(certificate), nil
}

// Sec_certificate_copy_ref.
//
// See: https://developer.apple.com/documentation/Security/sec_certificate_copy_ref(_:)
func Sec_certificate_copy_ref(certificate Sec_certificate_t) SecCertificateRef {
	result, callErr := TrySec_certificate_copy_ref(certificate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_certificate_create func(certificate SecCertificateRef) Sec_certificate_t
var _sec_certificate_createErr error

// CanCallSec_certificate_create reports whether the symbol for sec_certificate_create is available on this system.
func CanCallSec_certificate_create() bool {
	return _sec_certificate_create != nil
}

// Sec_certificate_createCallError returns the symbol lookup or registration error for sec_certificate_create.
func Sec_certificate_createCallError() error {
	if _sec_certificate_create != nil {
		return nil
	}
	return symbolCallError("sec_certificate_create", "10.14", _sec_certificate_createErr)
}

// TrySec_certificate_create calls Sec_certificate_create without panicking when the symbol is unavailable.
func TrySec_certificate_create(certificate SecCertificateRef) (Sec_certificate_t, error) {
	if err := Sec_certificate_createCallError(); err != nil {
		return *new(Sec_certificate_t), err
	}
	return _sec_certificate_create(certificate), nil
}

// Sec_certificate_create.
//
// See: https://developer.apple.com/documentation/Security/sec_certificate_create(_:)
func Sec_certificate_create(certificate SecCertificateRef) Sec_certificate_t {
	result, callErr := TrySec_certificate_create(certificate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_identity_access_certificates func(identity Sec_identity_t) bool
var _sec_identity_access_certificatesErr error

// CanCallSec_identity_access_certificates reports whether the symbol for sec_identity_access_certificates is available on this system.
func CanCallSec_identity_access_certificates() bool {
	return _sec_identity_access_certificates != nil
}

// Sec_identity_access_certificatesCallError returns the symbol lookup or registration error for sec_identity_access_certificates.
func Sec_identity_access_certificatesCallError() error {
	if _sec_identity_access_certificates != nil {
		return nil
	}
	return symbolCallError("sec_identity_access_certificates", "10.15", _sec_identity_access_certificatesErr)
}

// TrySec_identity_access_certificates calls Sec_identity_access_certificates without panicking when the symbol is unavailable.
func TrySec_identity_access_certificates(identity Sec_identity_t) (bool, error) {
	if err := Sec_identity_access_certificatesCallError(); err != nil {
		return false, err
	}
	return _sec_identity_access_certificates(identity), nil
}

// Sec_identity_access_certificates.
//
// See: https://developer.apple.com/documentation/Security/sec_identity_access_certificates(_:_:)
func Sec_identity_access_certificates(identity Sec_identity_t) bool {
	result, callErr := TrySec_identity_access_certificates(identity)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_identity_copy_certificates_ref func(identity Sec_identity_t) corefoundation.CFArrayRef
var _sec_identity_copy_certificates_refErr error

// CanCallSec_identity_copy_certificates_ref reports whether the symbol for sec_identity_copy_certificates_ref is available on this system.
func CanCallSec_identity_copy_certificates_ref() bool {
	return _sec_identity_copy_certificates_ref != nil
}

// Sec_identity_copy_certificates_refCallError returns the symbol lookup or registration error for sec_identity_copy_certificates_ref.
func Sec_identity_copy_certificates_refCallError() error {
	if _sec_identity_copy_certificates_ref != nil {
		return nil
	}
	return symbolCallError("sec_identity_copy_certificates_ref", "10.14", _sec_identity_copy_certificates_refErr)
}

// TrySec_identity_copy_certificates_ref calls Sec_identity_copy_certificates_ref without panicking when the symbol is unavailable.
func TrySec_identity_copy_certificates_ref(identity Sec_identity_t) (corefoundation.CFArrayRef, error) {
	if err := Sec_identity_copy_certificates_refCallError(); err != nil {
		return 0, err
	}
	return _sec_identity_copy_certificates_ref(identity), nil
}

// Sec_identity_copy_certificates_ref.
//
// See: https://developer.apple.com/documentation/Security/sec_identity_copy_certificates_ref(_:)
func Sec_identity_copy_certificates_ref(identity Sec_identity_t) corefoundation.CFArrayRef {
	result, callErr := TrySec_identity_copy_certificates_ref(identity)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_identity_copy_ref func(identity Sec_identity_t) SecIdentityRef
var _sec_identity_copy_refErr error

// CanCallSec_identity_copy_ref reports whether the symbol for sec_identity_copy_ref is available on this system.
func CanCallSec_identity_copy_ref() bool {
	return _sec_identity_copy_ref != nil
}

// Sec_identity_copy_refCallError returns the symbol lookup or registration error for sec_identity_copy_ref.
func Sec_identity_copy_refCallError() error {
	if _sec_identity_copy_ref != nil {
		return nil
	}
	return symbolCallError("sec_identity_copy_ref", "10.14", _sec_identity_copy_refErr)
}

// TrySec_identity_copy_ref calls Sec_identity_copy_ref without panicking when the symbol is unavailable.
func TrySec_identity_copy_ref(identity Sec_identity_t) (SecIdentityRef, error) {
	if err := Sec_identity_copy_refCallError(); err != nil {
		return 0, err
	}
	return _sec_identity_copy_ref(identity), nil
}

// Sec_identity_copy_ref.
//
// See: https://developer.apple.com/documentation/Security/sec_identity_copy_ref(_:)
func Sec_identity_copy_ref(identity Sec_identity_t) SecIdentityRef {
	result, callErr := TrySec_identity_copy_ref(identity)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_identity_create func(identity SecIdentityRef) Sec_identity_t
var _sec_identity_createErr error

// CanCallSec_identity_create reports whether the symbol for sec_identity_create is available on this system.
func CanCallSec_identity_create() bool {
	return _sec_identity_create != nil
}

// Sec_identity_createCallError returns the symbol lookup or registration error for sec_identity_create.
func Sec_identity_createCallError() error {
	if _sec_identity_create != nil {
		return nil
	}
	return symbolCallError("sec_identity_create", "10.14", _sec_identity_createErr)
}

// TrySec_identity_create calls Sec_identity_create without panicking when the symbol is unavailable.
func TrySec_identity_create(identity SecIdentityRef) (Sec_identity_t, error) {
	if err := Sec_identity_createCallError(); err != nil {
		return *new(Sec_identity_t), err
	}
	return _sec_identity_create(identity), nil
}

// Sec_identity_create.
//
// See: https://developer.apple.com/documentation/Security/sec_identity_create(_:)
func Sec_identity_create(identity SecIdentityRef) Sec_identity_t {
	result, callErr := TrySec_identity_create(identity)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_identity_create_with_certificates func(identity SecIdentityRef, certificates corefoundation.CFArrayRef) Sec_identity_t
var _sec_identity_create_with_certificatesErr error

// CanCallSec_identity_create_with_certificates reports whether the symbol for sec_identity_create_with_certificates is available on this system.
func CanCallSec_identity_create_with_certificates() bool {
	return _sec_identity_create_with_certificates != nil
}

// Sec_identity_create_with_certificatesCallError returns the symbol lookup or registration error for sec_identity_create_with_certificates.
func Sec_identity_create_with_certificatesCallError() error {
	if _sec_identity_create_with_certificates != nil {
		return nil
	}
	return symbolCallError("sec_identity_create_with_certificates", "10.14", _sec_identity_create_with_certificatesErr)
}

// TrySec_identity_create_with_certificates calls Sec_identity_create_with_certificates without panicking when the symbol is unavailable.
func TrySec_identity_create_with_certificates(identity SecIdentityRef, certificates corefoundation.CFArrayRef) (Sec_identity_t, error) {
	if err := Sec_identity_create_with_certificatesCallError(); err != nil {
		return *new(Sec_identity_t), err
	}
	return _sec_identity_create_with_certificates(identity, certificates), nil
}

// Sec_identity_create_with_certificates.
//
// See: https://developer.apple.com/documentation/Security/sec_identity_create_with_certificates(_:_:)
func Sec_identity_create_with_certificates(identity SecIdentityRef, certificates corefoundation.CFArrayRef) Sec_identity_t {
	result, callErr := TrySec_identity_create_with_certificates(identity, certificates)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_metadata_access_distinguished_names func(metadata Sec_protocol_metadata_t) bool
var _sec_protocol_metadata_access_distinguished_namesErr error

// CanCallSec_protocol_metadata_access_distinguished_names reports whether the symbol for sec_protocol_metadata_access_distinguished_names is available on this system.
func CanCallSec_protocol_metadata_access_distinguished_names() bool {
	return _sec_protocol_metadata_access_distinguished_names != nil
}

// Sec_protocol_metadata_access_distinguished_namesCallError returns the symbol lookup or registration error for sec_protocol_metadata_access_distinguished_names.
func Sec_protocol_metadata_access_distinguished_namesCallError() error {
	if _sec_protocol_metadata_access_distinguished_names != nil {
		return nil
	}
	return symbolCallError("sec_protocol_metadata_access_distinguished_names", "10.14", _sec_protocol_metadata_access_distinguished_namesErr)
}

// TrySec_protocol_metadata_access_distinguished_names calls Sec_protocol_metadata_access_distinguished_names without panicking when the symbol is unavailable.
func TrySec_protocol_metadata_access_distinguished_names(metadata Sec_protocol_metadata_t) (bool, error) {
	if err := Sec_protocol_metadata_access_distinguished_namesCallError(); err != nil {
		return false, err
	}
	return _sec_protocol_metadata_access_distinguished_names(metadata), nil
}

// Sec_protocol_metadata_access_distinguished_names.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_access_distinguished_names(_:_:)
func Sec_protocol_metadata_access_distinguished_names(metadata Sec_protocol_metadata_t) bool {
	result, callErr := TrySec_protocol_metadata_access_distinguished_names(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_metadata_access_ocsp_response func(metadata Sec_protocol_metadata_t) bool
var _sec_protocol_metadata_access_ocsp_responseErr error

// CanCallSec_protocol_metadata_access_ocsp_response reports whether the symbol for sec_protocol_metadata_access_ocsp_response is available on this system.
func CanCallSec_protocol_metadata_access_ocsp_response() bool {
	return _sec_protocol_metadata_access_ocsp_response != nil
}

// Sec_protocol_metadata_access_ocsp_responseCallError returns the symbol lookup or registration error for sec_protocol_metadata_access_ocsp_response.
func Sec_protocol_metadata_access_ocsp_responseCallError() error {
	if _sec_protocol_metadata_access_ocsp_response != nil {
		return nil
	}
	return symbolCallError("sec_protocol_metadata_access_ocsp_response", "10.14", _sec_protocol_metadata_access_ocsp_responseErr)
}

// TrySec_protocol_metadata_access_ocsp_response calls Sec_protocol_metadata_access_ocsp_response without panicking when the symbol is unavailable.
func TrySec_protocol_metadata_access_ocsp_response(metadata Sec_protocol_metadata_t) (bool, error) {
	if err := Sec_protocol_metadata_access_ocsp_responseCallError(); err != nil {
		return false, err
	}
	return _sec_protocol_metadata_access_ocsp_response(metadata), nil
}

// Sec_protocol_metadata_access_ocsp_response.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_access_ocsp_response(_:_:)
func Sec_protocol_metadata_access_ocsp_response(metadata Sec_protocol_metadata_t) bool {
	result, callErr := TrySec_protocol_metadata_access_ocsp_response(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_metadata_access_peer_certificate_chain func(metadata Sec_protocol_metadata_t) bool
var _sec_protocol_metadata_access_peer_certificate_chainErr error

// CanCallSec_protocol_metadata_access_peer_certificate_chain reports whether the symbol for sec_protocol_metadata_access_peer_certificate_chain is available on this system.
func CanCallSec_protocol_metadata_access_peer_certificate_chain() bool {
	return _sec_protocol_metadata_access_peer_certificate_chain != nil
}

// Sec_protocol_metadata_access_peer_certificate_chainCallError returns the symbol lookup or registration error for sec_protocol_metadata_access_peer_certificate_chain.
func Sec_protocol_metadata_access_peer_certificate_chainCallError() error {
	if _sec_protocol_metadata_access_peer_certificate_chain != nil {
		return nil
	}
	return symbolCallError("sec_protocol_metadata_access_peer_certificate_chain", "10.14", _sec_protocol_metadata_access_peer_certificate_chainErr)
}

// TrySec_protocol_metadata_access_peer_certificate_chain calls Sec_protocol_metadata_access_peer_certificate_chain without panicking when the symbol is unavailable.
func TrySec_protocol_metadata_access_peer_certificate_chain(metadata Sec_protocol_metadata_t) (bool, error) {
	if err := Sec_protocol_metadata_access_peer_certificate_chainCallError(); err != nil {
		return false, err
	}
	return _sec_protocol_metadata_access_peer_certificate_chain(metadata), nil
}

// Sec_protocol_metadata_access_peer_certificate_chain.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_access_peer_certificate_chain(_:_:)
func Sec_protocol_metadata_access_peer_certificate_chain(metadata Sec_protocol_metadata_t) bool {
	result, callErr := TrySec_protocol_metadata_access_peer_certificate_chain(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_metadata_access_pre_shared_keys func(metadata Sec_protocol_metadata_t) bool
var _sec_protocol_metadata_access_pre_shared_keysErr error

// CanCallSec_protocol_metadata_access_pre_shared_keys reports whether the symbol for sec_protocol_metadata_access_pre_shared_keys is available on this system.
func CanCallSec_protocol_metadata_access_pre_shared_keys() bool {
	return _sec_protocol_metadata_access_pre_shared_keys != nil
}

// Sec_protocol_metadata_access_pre_shared_keysCallError returns the symbol lookup or registration error for sec_protocol_metadata_access_pre_shared_keys.
func Sec_protocol_metadata_access_pre_shared_keysCallError() error {
	if _sec_protocol_metadata_access_pre_shared_keys != nil {
		return nil
	}
	return symbolCallError("sec_protocol_metadata_access_pre_shared_keys", "10.15", _sec_protocol_metadata_access_pre_shared_keysErr)
}

// TrySec_protocol_metadata_access_pre_shared_keys calls Sec_protocol_metadata_access_pre_shared_keys without panicking when the symbol is unavailable.
func TrySec_protocol_metadata_access_pre_shared_keys(metadata Sec_protocol_metadata_t) (bool, error) {
	if err := Sec_protocol_metadata_access_pre_shared_keysCallError(); err != nil {
		return false, err
	}
	return _sec_protocol_metadata_access_pre_shared_keys(metadata), nil
}

// Sec_protocol_metadata_access_pre_shared_keys.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_access_pre_shared_keys(_:_:)
func Sec_protocol_metadata_access_pre_shared_keys(metadata Sec_protocol_metadata_t) bool {
	result, callErr := TrySec_protocol_metadata_access_pre_shared_keys(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_metadata_access_supported_signature_algorithms func(metadata Sec_protocol_metadata_t) bool
var _sec_protocol_metadata_access_supported_signature_algorithmsErr error

// CanCallSec_protocol_metadata_access_supported_signature_algorithms reports whether the symbol for sec_protocol_metadata_access_supported_signature_algorithms is available on this system.
func CanCallSec_protocol_metadata_access_supported_signature_algorithms() bool {
	return _sec_protocol_metadata_access_supported_signature_algorithms != nil
}

// Sec_protocol_metadata_access_supported_signature_algorithmsCallError returns the symbol lookup or registration error for sec_protocol_metadata_access_supported_signature_algorithms.
func Sec_protocol_metadata_access_supported_signature_algorithmsCallError() error {
	if _sec_protocol_metadata_access_supported_signature_algorithms != nil {
		return nil
	}
	return symbolCallError("sec_protocol_metadata_access_supported_signature_algorithms", "10.14", _sec_protocol_metadata_access_supported_signature_algorithmsErr)
}

// TrySec_protocol_metadata_access_supported_signature_algorithms calls Sec_protocol_metadata_access_supported_signature_algorithms without panicking when the symbol is unavailable.
func TrySec_protocol_metadata_access_supported_signature_algorithms(metadata Sec_protocol_metadata_t) (bool, error) {
	if err := Sec_protocol_metadata_access_supported_signature_algorithmsCallError(); err != nil {
		return false, err
	}
	return _sec_protocol_metadata_access_supported_signature_algorithms(metadata), nil
}

// Sec_protocol_metadata_access_supported_signature_algorithms.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_access_supported_signature_algorithms(_:_:)
func Sec_protocol_metadata_access_supported_signature_algorithms(metadata Sec_protocol_metadata_t) bool {
	result, callErr := TrySec_protocol_metadata_access_supported_signature_algorithms(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_metadata_challenge_parameters_are_equal func(metadataA Sec_protocol_metadata_t, metadataB Sec_protocol_metadata_t) bool
var _sec_protocol_metadata_challenge_parameters_are_equalErr error

// CanCallSec_protocol_metadata_challenge_parameters_are_equal reports whether the symbol for sec_protocol_metadata_challenge_parameters_are_equal is available on this system.
func CanCallSec_protocol_metadata_challenge_parameters_are_equal() bool {
	return _sec_protocol_metadata_challenge_parameters_are_equal != nil
}

// Sec_protocol_metadata_challenge_parameters_are_equalCallError returns the symbol lookup or registration error for sec_protocol_metadata_challenge_parameters_are_equal.
func Sec_protocol_metadata_challenge_parameters_are_equalCallError() error {
	if _sec_protocol_metadata_challenge_parameters_are_equal != nil {
		return nil
	}
	return symbolCallError("sec_protocol_metadata_challenge_parameters_are_equal", "10.14", _sec_protocol_metadata_challenge_parameters_are_equalErr)
}

// TrySec_protocol_metadata_challenge_parameters_are_equal calls Sec_protocol_metadata_challenge_parameters_are_equal without panicking when the symbol is unavailable.
func TrySec_protocol_metadata_challenge_parameters_are_equal(metadataA Sec_protocol_metadata_t, metadataB Sec_protocol_metadata_t) (bool, error) {
	if err := Sec_protocol_metadata_challenge_parameters_are_equalCallError(); err != nil {
		return false, err
	}
	return _sec_protocol_metadata_challenge_parameters_are_equal(metadataA, metadataB), nil
}

// Sec_protocol_metadata_challenge_parameters_are_equal.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_challenge_parameters_are_equal(_:_:)
func Sec_protocol_metadata_challenge_parameters_are_equal(metadataA Sec_protocol_metadata_t, metadataB Sec_protocol_metadata_t) bool {
	result, callErr := TrySec_protocol_metadata_challenge_parameters_are_equal(metadataA, metadataB)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_metadata_copy_negotiated_protocol func(metadata Sec_protocol_metadata_t) *byte
var _sec_protocol_metadata_copy_negotiated_protocolErr error

// CanCallSec_protocol_metadata_copy_negotiated_protocol reports whether the symbol for sec_protocol_metadata_copy_negotiated_protocol is available on this system.
func CanCallSec_protocol_metadata_copy_negotiated_protocol() bool {
	return _sec_protocol_metadata_copy_negotiated_protocol != nil
}

// Sec_protocol_metadata_copy_negotiated_protocolCallError returns the symbol lookup or registration error for sec_protocol_metadata_copy_negotiated_protocol.
func Sec_protocol_metadata_copy_negotiated_protocolCallError() error {
	if _sec_protocol_metadata_copy_negotiated_protocol != nil {
		return nil
	}
	return symbolCallError("sec_protocol_metadata_copy_negotiated_protocol", "15.5", _sec_protocol_metadata_copy_negotiated_protocolErr)
}

// TrySec_protocol_metadata_copy_negotiated_protocol calls Sec_protocol_metadata_copy_negotiated_protocol without panicking when the symbol is unavailable.
func TrySec_protocol_metadata_copy_negotiated_protocol(metadata Sec_protocol_metadata_t) (*byte, error) {
	if err := Sec_protocol_metadata_copy_negotiated_protocolCallError(); err != nil {
		return nil, err
	}
	return _sec_protocol_metadata_copy_negotiated_protocol(metadata), nil
}

// Sec_protocol_metadata_copy_negotiated_protocol.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_copy_negotiated_protocol(_:)
func Sec_protocol_metadata_copy_negotiated_protocol(metadata Sec_protocol_metadata_t) *byte {
	result, callErr := TrySec_protocol_metadata_copy_negotiated_protocol(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_metadata_copy_peer_public_key func(metadata Sec_protocol_metadata_t) uintptr
var _sec_protocol_metadata_copy_peer_public_keyErr error

// CanCallSec_protocol_metadata_copy_peer_public_key reports whether the symbol for sec_protocol_metadata_copy_peer_public_key is available on this system.
func CanCallSec_protocol_metadata_copy_peer_public_key() bool {
	return _sec_protocol_metadata_copy_peer_public_key != nil
}

// Sec_protocol_metadata_copy_peer_public_keyCallError returns the symbol lookup or registration error for sec_protocol_metadata_copy_peer_public_key.
func Sec_protocol_metadata_copy_peer_public_keyCallError() error {
	if _sec_protocol_metadata_copy_peer_public_key != nil {
		return nil
	}
	return symbolCallError("sec_protocol_metadata_copy_peer_public_key", "10.14", _sec_protocol_metadata_copy_peer_public_keyErr)
}

// TrySec_protocol_metadata_copy_peer_public_key calls Sec_protocol_metadata_copy_peer_public_key without panicking when the symbol is unavailable.
func TrySec_protocol_metadata_copy_peer_public_key(metadata Sec_protocol_metadata_t) (dispatch.Data, error) {
	if err := Sec_protocol_metadata_copy_peer_public_keyCallError(); err != nil {
		return dispatch.DataFromHandle(0), err
	}
	return dispatch.DataFromHandle(_sec_protocol_metadata_copy_peer_public_key(metadata)), nil
}

// Sec_protocol_metadata_copy_peer_public_key.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_copy_peer_public_key(_:)
func Sec_protocol_metadata_copy_peer_public_key(metadata Sec_protocol_metadata_t) dispatch.Data {
	result, callErr := TrySec_protocol_metadata_copy_peer_public_key(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_metadata_copy_server_name func(metadata Sec_protocol_metadata_t) *byte
var _sec_protocol_metadata_copy_server_nameErr error

// CanCallSec_protocol_metadata_copy_server_name reports whether the symbol for sec_protocol_metadata_copy_server_name is available on this system.
func CanCallSec_protocol_metadata_copy_server_name() bool {
	return _sec_protocol_metadata_copy_server_name != nil
}

// Sec_protocol_metadata_copy_server_nameCallError returns the symbol lookup or registration error for sec_protocol_metadata_copy_server_name.
func Sec_protocol_metadata_copy_server_nameCallError() error {
	if _sec_protocol_metadata_copy_server_name != nil {
		return nil
	}
	return symbolCallError("sec_protocol_metadata_copy_server_name", "15.5", _sec_protocol_metadata_copy_server_nameErr)
}

// TrySec_protocol_metadata_copy_server_name calls Sec_protocol_metadata_copy_server_name without panicking when the symbol is unavailable.
func TrySec_protocol_metadata_copy_server_name(metadata Sec_protocol_metadata_t) (*byte, error) {
	if err := Sec_protocol_metadata_copy_server_nameCallError(); err != nil {
		return nil, err
	}
	return _sec_protocol_metadata_copy_server_name(metadata), nil
}

// Sec_protocol_metadata_copy_server_name.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_copy_server_name(_:)
func Sec_protocol_metadata_copy_server_name(metadata Sec_protocol_metadata_t) *byte {
	result, callErr := TrySec_protocol_metadata_copy_server_name(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_metadata_create_secret func(metadata Sec_protocol_metadata_t, label_len uintptr, label string, exporter_length uintptr) uintptr
var _sec_protocol_metadata_create_secretErr error

// CanCallSec_protocol_metadata_create_secret reports whether the symbol for sec_protocol_metadata_create_secret is available on this system.
func CanCallSec_protocol_metadata_create_secret() bool {
	return _sec_protocol_metadata_create_secret != nil
}

// Sec_protocol_metadata_create_secretCallError returns the symbol lookup or registration error for sec_protocol_metadata_create_secret.
func Sec_protocol_metadata_create_secretCallError() error {
	if _sec_protocol_metadata_create_secret != nil {
		return nil
	}
	return symbolCallError("sec_protocol_metadata_create_secret", "10.14", _sec_protocol_metadata_create_secretErr)
}

// TrySec_protocol_metadata_create_secret calls Sec_protocol_metadata_create_secret without panicking when the symbol is unavailable.
func TrySec_protocol_metadata_create_secret(metadata Sec_protocol_metadata_t, label_len uintptr, label string, exporter_length uintptr) (dispatch.Data, error) {
	if err := Sec_protocol_metadata_create_secretCallError(); err != nil {
		return dispatch.DataFromHandle(0), err
	}
	return dispatch.DataFromHandle(_sec_protocol_metadata_create_secret(metadata, label_len, label, exporter_length)), nil
}

// Sec_protocol_metadata_create_secret.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_create_secret(_:_:_:_:)
func Sec_protocol_metadata_create_secret(metadata Sec_protocol_metadata_t, label_len uintptr, label string, exporter_length uintptr) dispatch.Data {
	result, callErr := TrySec_protocol_metadata_create_secret(metadata, label_len, label, exporter_length)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_metadata_create_secret_with_context func(metadata Sec_protocol_metadata_t, label_len uintptr, label string, context_len uintptr, context *uint8, exporter_length uintptr) uintptr
var _sec_protocol_metadata_create_secret_with_contextErr error

// CanCallSec_protocol_metadata_create_secret_with_context reports whether the symbol for sec_protocol_metadata_create_secret_with_context is available on this system.
func CanCallSec_protocol_metadata_create_secret_with_context() bool {
	return _sec_protocol_metadata_create_secret_with_context != nil
}

// Sec_protocol_metadata_create_secret_with_contextCallError returns the symbol lookup or registration error for sec_protocol_metadata_create_secret_with_context.
func Sec_protocol_metadata_create_secret_with_contextCallError() error {
	if _sec_protocol_metadata_create_secret_with_context != nil {
		return nil
	}
	return symbolCallError("sec_protocol_metadata_create_secret_with_context", "10.14", _sec_protocol_metadata_create_secret_with_contextErr)
}

// TrySec_protocol_metadata_create_secret_with_context calls Sec_protocol_metadata_create_secret_with_context without panicking when the symbol is unavailable.
func TrySec_protocol_metadata_create_secret_with_context(metadata Sec_protocol_metadata_t, label_len uintptr, label string, context_len uintptr, context *uint8, exporter_length uintptr) (dispatch.Data, error) {
	if err := Sec_protocol_metadata_create_secret_with_contextCallError(); err != nil {
		return dispatch.DataFromHandle(0), err
	}
	return dispatch.DataFromHandle(_sec_protocol_metadata_create_secret_with_context(metadata, label_len, label, context_len, context, exporter_length)), nil
}

// Sec_protocol_metadata_create_secret_with_context.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_create_secret_with_context(_:_:_:_:_:_:)
func Sec_protocol_metadata_create_secret_with_context(metadata Sec_protocol_metadata_t, label_len uintptr, label string, context_len uintptr, context *uint8, exporter_length uintptr) dispatch.Data {
	result, callErr := TrySec_protocol_metadata_create_secret_with_context(metadata, label_len, label, context_len, context, exporter_length)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_metadata_get_early_data_accepted func(metadata Sec_protocol_metadata_t) bool
var _sec_protocol_metadata_get_early_data_acceptedErr error

// CanCallSec_protocol_metadata_get_early_data_accepted reports whether the symbol for sec_protocol_metadata_get_early_data_accepted is available on this system.
func CanCallSec_protocol_metadata_get_early_data_accepted() bool {
	return _sec_protocol_metadata_get_early_data_accepted != nil
}

// Sec_protocol_metadata_get_early_data_acceptedCallError returns the symbol lookup or registration error for sec_protocol_metadata_get_early_data_accepted.
func Sec_protocol_metadata_get_early_data_acceptedCallError() error {
	if _sec_protocol_metadata_get_early_data_accepted != nil {
		return nil
	}
	return symbolCallError("sec_protocol_metadata_get_early_data_accepted", "10.14", _sec_protocol_metadata_get_early_data_acceptedErr)
}

// TrySec_protocol_metadata_get_early_data_accepted calls Sec_protocol_metadata_get_early_data_accepted without panicking when the symbol is unavailable.
func TrySec_protocol_metadata_get_early_data_accepted(metadata Sec_protocol_metadata_t) (bool, error) {
	if err := Sec_protocol_metadata_get_early_data_acceptedCallError(); err != nil {
		return false, err
	}
	return _sec_protocol_metadata_get_early_data_accepted(metadata), nil
}

// Sec_protocol_metadata_get_early_data_accepted.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_get_early_data_accepted(_:)
func Sec_protocol_metadata_get_early_data_accepted(metadata Sec_protocol_metadata_t) bool {
	result, callErr := TrySec_protocol_metadata_get_early_data_accepted(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_metadata_get_negotiated_ciphersuite func(metadata Sec_protocol_metadata_t) SSLCipherSuite
var _sec_protocol_metadata_get_negotiated_ciphersuiteErr error

// CanCallSec_protocol_metadata_get_negotiated_ciphersuite reports whether the symbol for sec_protocol_metadata_get_negotiated_ciphersuite is available on this system.
func CanCallSec_protocol_metadata_get_negotiated_ciphersuite() bool {
	return _sec_protocol_metadata_get_negotiated_ciphersuite != nil
}

// Sec_protocol_metadata_get_negotiated_ciphersuiteCallError returns the symbol lookup or registration error for sec_protocol_metadata_get_negotiated_ciphersuite.
func Sec_protocol_metadata_get_negotiated_ciphersuiteCallError() error {
	if _sec_protocol_metadata_get_negotiated_ciphersuite != nil {
		return nil
	}
	return symbolCallError("sec_protocol_metadata_get_negotiated_ciphersuite", "10.14", _sec_protocol_metadata_get_negotiated_ciphersuiteErr)
}

// TrySec_protocol_metadata_get_negotiated_ciphersuite calls Sec_protocol_metadata_get_negotiated_ciphersuite without panicking when the symbol is unavailable.
func TrySec_protocol_metadata_get_negotiated_ciphersuite(metadata Sec_protocol_metadata_t) (SSLCipherSuite, error) {
	if err := Sec_protocol_metadata_get_negotiated_ciphersuiteCallError(); err != nil {
		return *new(SSLCipherSuite), err
	}
	return _sec_protocol_metadata_get_negotiated_ciphersuite(metadata), nil
}

// Sec_protocol_metadata_get_negotiated_ciphersuite.
//
// Deprecated: Deprecated since macOS 10.15.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_get_negotiated_ciphersuite(_:)
func Sec_protocol_metadata_get_negotiated_ciphersuite(metadata Sec_protocol_metadata_t) SSLCipherSuite {
	result, callErr := TrySec_protocol_metadata_get_negotiated_ciphersuite(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_metadata_get_negotiated_protocol func(metadata Sec_protocol_metadata_t) *byte
var _sec_protocol_metadata_get_negotiated_protocolErr error

// CanCallSec_protocol_metadata_get_negotiated_protocol reports whether the symbol for sec_protocol_metadata_get_negotiated_protocol is available on this system.
func CanCallSec_protocol_metadata_get_negotiated_protocol() bool {
	return _sec_protocol_metadata_get_negotiated_protocol != nil
}

// Sec_protocol_metadata_get_negotiated_protocolCallError returns the symbol lookup or registration error for sec_protocol_metadata_get_negotiated_protocol.
func Sec_protocol_metadata_get_negotiated_protocolCallError() error {
	if _sec_protocol_metadata_get_negotiated_protocol != nil {
		return nil
	}
	return symbolCallError("sec_protocol_metadata_get_negotiated_protocol", "10.14", _sec_protocol_metadata_get_negotiated_protocolErr)
}

// TrySec_protocol_metadata_get_negotiated_protocol calls Sec_protocol_metadata_get_negotiated_protocol without panicking when the symbol is unavailable.
func TrySec_protocol_metadata_get_negotiated_protocol(metadata Sec_protocol_metadata_t) (*byte, error) {
	if err := Sec_protocol_metadata_get_negotiated_protocolCallError(); err != nil {
		return nil, err
	}
	return _sec_protocol_metadata_get_negotiated_protocol(metadata), nil
}

// Sec_protocol_metadata_get_negotiated_protocol.
//
// Deprecated: Deprecated since macOS 15.5.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_get_negotiated_protocol(_:)
func Sec_protocol_metadata_get_negotiated_protocol(metadata Sec_protocol_metadata_t) *byte {
	result, callErr := TrySec_protocol_metadata_get_negotiated_protocol(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_metadata_get_negotiated_protocol_version func(metadata Sec_protocol_metadata_t) SSLProtocol
var _sec_protocol_metadata_get_negotiated_protocol_versionErr error

// CanCallSec_protocol_metadata_get_negotiated_protocol_version reports whether the symbol for sec_protocol_metadata_get_negotiated_protocol_version is available on this system.
func CanCallSec_protocol_metadata_get_negotiated_protocol_version() bool {
	return _sec_protocol_metadata_get_negotiated_protocol_version != nil
}

// Sec_protocol_metadata_get_negotiated_protocol_versionCallError returns the symbol lookup or registration error for sec_protocol_metadata_get_negotiated_protocol_version.
func Sec_protocol_metadata_get_negotiated_protocol_versionCallError() error {
	if _sec_protocol_metadata_get_negotiated_protocol_version != nil {
		return nil
	}
	return symbolCallError("sec_protocol_metadata_get_negotiated_protocol_version", "10.14", _sec_protocol_metadata_get_negotiated_protocol_versionErr)
}

// TrySec_protocol_metadata_get_negotiated_protocol_version calls Sec_protocol_metadata_get_negotiated_protocol_version without panicking when the symbol is unavailable.
func TrySec_protocol_metadata_get_negotiated_protocol_version(metadata Sec_protocol_metadata_t) (SSLProtocol, error) {
	if err := Sec_protocol_metadata_get_negotiated_protocol_versionCallError(); err != nil {
		return *new(SSLProtocol), err
	}
	return _sec_protocol_metadata_get_negotiated_protocol_version(metadata), nil
}

// Sec_protocol_metadata_get_negotiated_protocol_version.
//
// Deprecated: Deprecated since macOS 10.15.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_get_negotiated_protocol_version(_:)
func Sec_protocol_metadata_get_negotiated_protocol_version(metadata Sec_protocol_metadata_t) SSLProtocol {
	result, callErr := TrySec_protocol_metadata_get_negotiated_protocol_version(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_metadata_get_negotiated_tls_ciphersuite func(metadata Sec_protocol_metadata_t) Tls_ciphersuite_t
var _sec_protocol_metadata_get_negotiated_tls_ciphersuiteErr error

// CanCallSec_protocol_metadata_get_negotiated_tls_ciphersuite reports whether the symbol for sec_protocol_metadata_get_negotiated_tls_ciphersuite is available on this system.
func CanCallSec_protocol_metadata_get_negotiated_tls_ciphersuite() bool {
	return _sec_protocol_metadata_get_negotiated_tls_ciphersuite != nil
}

// Sec_protocol_metadata_get_negotiated_tls_ciphersuiteCallError returns the symbol lookup or registration error for sec_protocol_metadata_get_negotiated_tls_ciphersuite.
func Sec_protocol_metadata_get_negotiated_tls_ciphersuiteCallError() error {
	if _sec_protocol_metadata_get_negotiated_tls_ciphersuite != nil {
		return nil
	}
	return symbolCallError("sec_protocol_metadata_get_negotiated_tls_ciphersuite", "10.15", _sec_protocol_metadata_get_negotiated_tls_ciphersuiteErr)
}

// TrySec_protocol_metadata_get_negotiated_tls_ciphersuite calls Sec_protocol_metadata_get_negotiated_tls_ciphersuite without panicking when the symbol is unavailable.
func TrySec_protocol_metadata_get_negotiated_tls_ciphersuite(metadata Sec_protocol_metadata_t) (Tls_ciphersuite_t, error) {
	if err := Sec_protocol_metadata_get_negotiated_tls_ciphersuiteCallError(); err != nil {
		return *new(Tls_ciphersuite_t), err
	}
	return _sec_protocol_metadata_get_negotiated_tls_ciphersuite(metadata), nil
}

// Sec_protocol_metadata_get_negotiated_tls_ciphersuite.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_get_negotiated_tls_ciphersuite(_:)
func Sec_protocol_metadata_get_negotiated_tls_ciphersuite(metadata Sec_protocol_metadata_t) Tls_ciphersuite_t {
	result, callErr := TrySec_protocol_metadata_get_negotiated_tls_ciphersuite(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_metadata_get_negotiated_tls_protocol_version func(metadata Sec_protocol_metadata_t) Tls_protocol_version_t
var _sec_protocol_metadata_get_negotiated_tls_protocol_versionErr error

// CanCallSec_protocol_metadata_get_negotiated_tls_protocol_version reports whether the symbol for sec_protocol_metadata_get_negotiated_tls_protocol_version is available on this system.
func CanCallSec_protocol_metadata_get_negotiated_tls_protocol_version() bool {
	return _sec_protocol_metadata_get_negotiated_tls_protocol_version != nil
}

// Sec_protocol_metadata_get_negotiated_tls_protocol_versionCallError returns the symbol lookup or registration error for sec_protocol_metadata_get_negotiated_tls_protocol_version.
func Sec_protocol_metadata_get_negotiated_tls_protocol_versionCallError() error {
	if _sec_protocol_metadata_get_negotiated_tls_protocol_version != nil {
		return nil
	}
	return symbolCallError("sec_protocol_metadata_get_negotiated_tls_protocol_version", "10.15", _sec_protocol_metadata_get_negotiated_tls_protocol_versionErr)
}

// TrySec_protocol_metadata_get_negotiated_tls_protocol_version calls Sec_protocol_metadata_get_negotiated_tls_protocol_version without panicking when the symbol is unavailable.
func TrySec_protocol_metadata_get_negotiated_tls_protocol_version(metadata Sec_protocol_metadata_t) (Tls_protocol_version_t, error) {
	if err := Sec_protocol_metadata_get_negotiated_tls_protocol_versionCallError(); err != nil {
		return *new(Tls_protocol_version_t), err
	}
	return _sec_protocol_metadata_get_negotiated_tls_protocol_version(metadata), nil
}

// Sec_protocol_metadata_get_negotiated_tls_protocol_version.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_get_negotiated_tls_protocol_version(_:)
func Sec_protocol_metadata_get_negotiated_tls_protocol_version(metadata Sec_protocol_metadata_t) Tls_protocol_version_t {
	result, callErr := TrySec_protocol_metadata_get_negotiated_tls_protocol_version(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_metadata_get_server_name func(metadata Sec_protocol_metadata_t) *byte
var _sec_protocol_metadata_get_server_nameErr error

// CanCallSec_protocol_metadata_get_server_name reports whether the symbol for sec_protocol_metadata_get_server_name is available on this system.
func CanCallSec_protocol_metadata_get_server_name() bool {
	return _sec_protocol_metadata_get_server_name != nil
}

// Sec_protocol_metadata_get_server_nameCallError returns the symbol lookup or registration error for sec_protocol_metadata_get_server_name.
func Sec_protocol_metadata_get_server_nameCallError() error {
	if _sec_protocol_metadata_get_server_name != nil {
		return nil
	}
	return symbolCallError("sec_protocol_metadata_get_server_name", "10.14", _sec_protocol_metadata_get_server_nameErr)
}

// TrySec_protocol_metadata_get_server_name calls Sec_protocol_metadata_get_server_name without panicking when the symbol is unavailable.
func TrySec_protocol_metadata_get_server_name(metadata Sec_protocol_metadata_t) (*byte, error) {
	if err := Sec_protocol_metadata_get_server_nameCallError(); err != nil {
		return nil, err
	}
	return _sec_protocol_metadata_get_server_name(metadata), nil
}

// Sec_protocol_metadata_get_server_name.
//
// Deprecated: Deprecated since macOS 15.5.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_get_server_name(_:)
func Sec_protocol_metadata_get_server_name(metadata Sec_protocol_metadata_t) *byte {
	result, callErr := TrySec_protocol_metadata_get_server_name(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_metadata_peers_are_equal func(metadataA Sec_protocol_metadata_t, metadataB Sec_protocol_metadata_t) bool
var _sec_protocol_metadata_peers_are_equalErr error

// CanCallSec_protocol_metadata_peers_are_equal reports whether the symbol for sec_protocol_metadata_peers_are_equal is available on this system.
func CanCallSec_protocol_metadata_peers_are_equal() bool {
	return _sec_protocol_metadata_peers_are_equal != nil
}

// Sec_protocol_metadata_peers_are_equalCallError returns the symbol lookup or registration error for sec_protocol_metadata_peers_are_equal.
func Sec_protocol_metadata_peers_are_equalCallError() error {
	if _sec_protocol_metadata_peers_are_equal != nil {
		return nil
	}
	return symbolCallError("sec_protocol_metadata_peers_are_equal", "10.14", _sec_protocol_metadata_peers_are_equalErr)
}

// TrySec_protocol_metadata_peers_are_equal calls Sec_protocol_metadata_peers_are_equal without panicking when the symbol is unavailable.
func TrySec_protocol_metadata_peers_are_equal(metadataA Sec_protocol_metadata_t, metadataB Sec_protocol_metadata_t) (bool, error) {
	if err := Sec_protocol_metadata_peers_are_equalCallError(); err != nil {
		return false, err
	}
	return _sec_protocol_metadata_peers_are_equal(metadataA, metadataB), nil
}

// Sec_protocol_metadata_peers_are_equal.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_peers_are_equal(_:_:)
func Sec_protocol_metadata_peers_are_equal(metadataA Sec_protocol_metadata_t, metadataB Sec_protocol_metadata_t) bool {
	result, callErr := TrySec_protocol_metadata_peers_are_equal(metadataA, metadataB)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_options_add_pre_shared_key func(options Sec_protocol_options_t, psk uintptr, psk_identity uintptr)
var _sec_protocol_options_add_pre_shared_keyErr error

// CanCallSec_protocol_options_add_pre_shared_key reports whether the symbol for sec_protocol_options_add_pre_shared_key is available on this system.
func CanCallSec_protocol_options_add_pre_shared_key() bool {
	return _sec_protocol_options_add_pre_shared_key != nil
}

// Sec_protocol_options_add_pre_shared_keyCallError returns the symbol lookup or registration error for sec_protocol_options_add_pre_shared_key.
func Sec_protocol_options_add_pre_shared_keyCallError() error {
	if _sec_protocol_options_add_pre_shared_key != nil {
		return nil
	}
	return symbolCallError("sec_protocol_options_add_pre_shared_key", "10.14", _sec_protocol_options_add_pre_shared_keyErr)
}

// TrySec_protocol_options_add_pre_shared_key calls Sec_protocol_options_add_pre_shared_key without panicking when the symbol is unavailable.
func TrySec_protocol_options_add_pre_shared_key(options Sec_protocol_options_t, psk dispatch.Data, psk_identity dispatch.Data) error {
	if err := Sec_protocol_options_add_pre_shared_keyCallError(); err != nil {
		return err
	}
	_sec_protocol_options_add_pre_shared_key(options, uintptr(psk.Handle()), uintptr(psk_identity.Handle()))
	return nil
}

// Sec_protocol_options_add_pre_shared_key.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_add_pre_shared_key(_:_:_:)
func Sec_protocol_options_add_pre_shared_key(options Sec_protocol_options_t, psk dispatch.Data, psk_identity dispatch.Data) {
	if callErr := TrySec_protocol_options_add_pre_shared_key(options, psk, psk_identity); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_add_tls_application_protocol func(options Sec_protocol_options_t, application_protocol string)
var _sec_protocol_options_add_tls_application_protocolErr error

// CanCallSec_protocol_options_add_tls_application_protocol reports whether the symbol for sec_protocol_options_add_tls_application_protocol is available on this system.
func CanCallSec_protocol_options_add_tls_application_protocol() bool {
	return _sec_protocol_options_add_tls_application_protocol != nil
}

// Sec_protocol_options_add_tls_application_protocolCallError returns the symbol lookup or registration error for sec_protocol_options_add_tls_application_protocol.
func Sec_protocol_options_add_tls_application_protocolCallError() error {
	if _sec_protocol_options_add_tls_application_protocol != nil {
		return nil
	}
	return symbolCallError("sec_protocol_options_add_tls_application_protocol", "10.14", _sec_protocol_options_add_tls_application_protocolErr)
}

// TrySec_protocol_options_add_tls_application_protocol calls Sec_protocol_options_add_tls_application_protocol without panicking when the symbol is unavailable.
func TrySec_protocol_options_add_tls_application_protocol(options Sec_protocol_options_t, application_protocol string) error {
	if err := Sec_protocol_options_add_tls_application_protocolCallError(); err != nil {
		return err
	}
	_sec_protocol_options_add_tls_application_protocol(options, application_protocol)
	return nil
}

// Sec_protocol_options_add_tls_application_protocol.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_add_tls_application_protocol(_:_:)
func Sec_protocol_options_add_tls_application_protocol(options Sec_protocol_options_t, application_protocol string) {
	if callErr := TrySec_protocol_options_add_tls_application_protocol(options, application_protocol); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_append_tls_ciphersuite func(options Sec_protocol_options_t, ciphersuite Tls_ciphersuite_t)
var _sec_protocol_options_append_tls_ciphersuiteErr error

// CanCallSec_protocol_options_append_tls_ciphersuite reports whether the symbol for sec_protocol_options_append_tls_ciphersuite is available on this system.
func CanCallSec_protocol_options_append_tls_ciphersuite() bool {
	return _sec_protocol_options_append_tls_ciphersuite != nil
}

// Sec_protocol_options_append_tls_ciphersuiteCallError returns the symbol lookup or registration error for sec_protocol_options_append_tls_ciphersuite.
func Sec_protocol_options_append_tls_ciphersuiteCallError() error {
	if _sec_protocol_options_append_tls_ciphersuite != nil {
		return nil
	}
	return symbolCallError("sec_protocol_options_append_tls_ciphersuite", "10.15", _sec_protocol_options_append_tls_ciphersuiteErr)
}

// TrySec_protocol_options_append_tls_ciphersuite calls Sec_protocol_options_append_tls_ciphersuite without panicking when the symbol is unavailable.
func TrySec_protocol_options_append_tls_ciphersuite(options Sec_protocol_options_t, ciphersuite Tls_ciphersuite_t) error {
	if err := Sec_protocol_options_append_tls_ciphersuiteCallError(); err != nil {
		return err
	}
	_sec_protocol_options_append_tls_ciphersuite(options, ciphersuite)
	return nil
}

// Sec_protocol_options_append_tls_ciphersuite.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_append_tls_ciphersuite(_:_:)
func Sec_protocol_options_append_tls_ciphersuite(options Sec_protocol_options_t, ciphersuite Tls_ciphersuite_t) {
	if callErr := TrySec_protocol_options_append_tls_ciphersuite(options, ciphersuite); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_append_tls_ciphersuite_group func(options Sec_protocol_options_t, group Tls_ciphersuite_group_t)
var _sec_protocol_options_append_tls_ciphersuite_groupErr error

// CanCallSec_protocol_options_append_tls_ciphersuite_group reports whether the symbol for sec_protocol_options_append_tls_ciphersuite_group is available on this system.
func CanCallSec_protocol_options_append_tls_ciphersuite_group() bool {
	return _sec_protocol_options_append_tls_ciphersuite_group != nil
}

// Sec_protocol_options_append_tls_ciphersuite_groupCallError returns the symbol lookup or registration error for sec_protocol_options_append_tls_ciphersuite_group.
func Sec_protocol_options_append_tls_ciphersuite_groupCallError() error {
	if _sec_protocol_options_append_tls_ciphersuite_group != nil {
		return nil
	}
	return symbolCallError("sec_protocol_options_append_tls_ciphersuite_group", "10.15", _sec_protocol_options_append_tls_ciphersuite_groupErr)
}

// TrySec_protocol_options_append_tls_ciphersuite_group calls Sec_protocol_options_append_tls_ciphersuite_group without panicking when the symbol is unavailable.
func TrySec_protocol_options_append_tls_ciphersuite_group(options Sec_protocol_options_t, group Tls_ciphersuite_group_t) error {
	if err := Sec_protocol_options_append_tls_ciphersuite_groupCallError(); err != nil {
		return err
	}
	_sec_protocol_options_append_tls_ciphersuite_group(options, group)
	return nil
}

// Sec_protocol_options_append_tls_ciphersuite_group.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_append_tls_ciphersuite_group(_:_:)
func Sec_protocol_options_append_tls_ciphersuite_group(options Sec_protocol_options_t, group Tls_ciphersuite_group_t) {
	if callErr := TrySec_protocol_options_append_tls_ciphersuite_group(options, group); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_are_equal func(optionsA Sec_protocol_options_t, optionsB Sec_protocol_options_t) bool
var _sec_protocol_options_are_equalErr error

// CanCallSec_protocol_options_are_equal reports whether the symbol for sec_protocol_options_are_equal is available on this system.
func CanCallSec_protocol_options_are_equal() bool {
	return _sec_protocol_options_are_equal != nil
}

// Sec_protocol_options_are_equalCallError returns the symbol lookup or registration error for sec_protocol_options_are_equal.
func Sec_protocol_options_are_equalCallError() error {
	if _sec_protocol_options_are_equal != nil {
		return nil
	}
	return symbolCallError("sec_protocol_options_are_equal", "10.15", _sec_protocol_options_are_equalErr)
}

// TrySec_protocol_options_are_equal calls Sec_protocol_options_are_equal without panicking when the symbol is unavailable.
func TrySec_protocol_options_are_equal(optionsA Sec_protocol_options_t, optionsB Sec_protocol_options_t) (bool, error) {
	if err := Sec_protocol_options_are_equalCallError(); err != nil {
		return false, err
	}
	return _sec_protocol_options_are_equal(optionsA, optionsB), nil
}

// Sec_protocol_options_are_equal.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_are_equal(_:_:)
func Sec_protocol_options_are_equal(optionsA Sec_protocol_options_t, optionsB Sec_protocol_options_t) bool {
	result, callErr := TrySec_protocol_options_are_equal(optionsA, optionsB)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_options_get_default_max_dtls_protocol_version func() Tls_protocol_version_t
var _sec_protocol_options_get_default_max_dtls_protocol_versionErr error

// CanCallSec_protocol_options_get_default_max_dtls_protocol_version reports whether the symbol for sec_protocol_options_get_default_max_dtls_protocol_version is available on this system.
func CanCallSec_protocol_options_get_default_max_dtls_protocol_version() bool {
	return _sec_protocol_options_get_default_max_dtls_protocol_version != nil
}

// Sec_protocol_options_get_default_max_dtls_protocol_versionCallError returns the symbol lookup or registration error for sec_protocol_options_get_default_max_dtls_protocol_version.
func Sec_protocol_options_get_default_max_dtls_protocol_versionCallError() error {
	if _sec_protocol_options_get_default_max_dtls_protocol_version != nil {
		return nil
	}
	return symbolCallError("sec_protocol_options_get_default_max_dtls_protocol_version", "10.15", _sec_protocol_options_get_default_max_dtls_protocol_versionErr)
}

// TrySec_protocol_options_get_default_max_dtls_protocol_version calls Sec_protocol_options_get_default_max_dtls_protocol_version without panicking when the symbol is unavailable.
func TrySec_protocol_options_get_default_max_dtls_protocol_version() (Tls_protocol_version_t, error) {
	if err := Sec_protocol_options_get_default_max_dtls_protocol_versionCallError(); err != nil {
		return *new(Tls_protocol_version_t), err
	}
	return _sec_protocol_options_get_default_max_dtls_protocol_version(), nil
}

// Sec_protocol_options_get_default_max_dtls_protocol_version.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_get_default_max_dtls_protocol_version()
func Sec_protocol_options_get_default_max_dtls_protocol_version() Tls_protocol_version_t {
	result, callErr := TrySec_protocol_options_get_default_max_dtls_protocol_version()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_options_get_default_max_tls_protocol_version func() Tls_protocol_version_t
var _sec_protocol_options_get_default_max_tls_protocol_versionErr error

// CanCallSec_protocol_options_get_default_max_tls_protocol_version reports whether the symbol for sec_protocol_options_get_default_max_tls_protocol_version is available on this system.
func CanCallSec_protocol_options_get_default_max_tls_protocol_version() bool {
	return _sec_protocol_options_get_default_max_tls_protocol_version != nil
}

// Sec_protocol_options_get_default_max_tls_protocol_versionCallError returns the symbol lookup or registration error for sec_protocol_options_get_default_max_tls_protocol_version.
func Sec_protocol_options_get_default_max_tls_protocol_versionCallError() error {
	if _sec_protocol_options_get_default_max_tls_protocol_version != nil {
		return nil
	}
	return symbolCallError("sec_protocol_options_get_default_max_tls_protocol_version", "10.15", _sec_protocol_options_get_default_max_tls_protocol_versionErr)
}

// TrySec_protocol_options_get_default_max_tls_protocol_version calls Sec_protocol_options_get_default_max_tls_protocol_version without panicking when the symbol is unavailable.
func TrySec_protocol_options_get_default_max_tls_protocol_version() (Tls_protocol_version_t, error) {
	if err := Sec_protocol_options_get_default_max_tls_protocol_versionCallError(); err != nil {
		return *new(Tls_protocol_version_t), err
	}
	return _sec_protocol_options_get_default_max_tls_protocol_version(), nil
}

// Sec_protocol_options_get_default_max_tls_protocol_version.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_get_default_max_tls_protocol_version()
func Sec_protocol_options_get_default_max_tls_protocol_version() Tls_protocol_version_t {
	result, callErr := TrySec_protocol_options_get_default_max_tls_protocol_version()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_options_get_default_min_dtls_protocol_version func() Tls_protocol_version_t
var _sec_protocol_options_get_default_min_dtls_protocol_versionErr error

// CanCallSec_protocol_options_get_default_min_dtls_protocol_version reports whether the symbol for sec_protocol_options_get_default_min_dtls_protocol_version is available on this system.
func CanCallSec_protocol_options_get_default_min_dtls_protocol_version() bool {
	return _sec_protocol_options_get_default_min_dtls_protocol_version != nil
}

// Sec_protocol_options_get_default_min_dtls_protocol_versionCallError returns the symbol lookup or registration error for sec_protocol_options_get_default_min_dtls_protocol_version.
func Sec_protocol_options_get_default_min_dtls_protocol_versionCallError() error {
	if _sec_protocol_options_get_default_min_dtls_protocol_version != nil {
		return nil
	}
	return symbolCallError("sec_protocol_options_get_default_min_dtls_protocol_version", "10.15", _sec_protocol_options_get_default_min_dtls_protocol_versionErr)
}

// TrySec_protocol_options_get_default_min_dtls_protocol_version calls Sec_protocol_options_get_default_min_dtls_protocol_version without panicking when the symbol is unavailable.
func TrySec_protocol_options_get_default_min_dtls_protocol_version() (Tls_protocol_version_t, error) {
	if err := Sec_protocol_options_get_default_min_dtls_protocol_versionCallError(); err != nil {
		return *new(Tls_protocol_version_t), err
	}
	return _sec_protocol_options_get_default_min_dtls_protocol_version(), nil
}

// Sec_protocol_options_get_default_min_dtls_protocol_version.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_get_default_min_dtls_protocol_version()
func Sec_protocol_options_get_default_min_dtls_protocol_version() Tls_protocol_version_t {
	result, callErr := TrySec_protocol_options_get_default_min_dtls_protocol_version()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_options_get_default_min_tls_protocol_version func() Tls_protocol_version_t
var _sec_protocol_options_get_default_min_tls_protocol_versionErr error

// CanCallSec_protocol_options_get_default_min_tls_protocol_version reports whether the symbol for sec_protocol_options_get_default_min_tls_protocol_version is available on this system.
func CanCallSec_protocol_options_get_default_min_tls_protocol_version() bool {
	return _sec_protocol_options_get_default_min_tls_protocol_version != nil
}

// Sec_protocol_options_get_default_min_tls_protocol_versionCallError returns the symbol lookup or registration error for sec_protocol_options_get_default_min_tls_protocol_version.
func Sec_protocol_options_get_default_min_tls_protocol_versionCallError() error {
	if _sec_protocol_options_get_default_min_tls_protocol_version != nil {
		return nil
	}
	return symbolCallError("sec_protocol_options_get_default_min_tls_protocol_version", "10.15", _sec_protocol_options_get_default_min_tls_protocol_versionErr)
}

// TrySec_protocol_options_get_default_min_tls_protocol_version calls Sec_protocol_options_get_default_min_tls_protocol_version without panicking when the symbol is unavailable.
func TrySec_protocol_options_get_default_min_tls_protocol_version() (Tls_protocol_version_t, error) {
	if err := Sec_protocol_options_get_default_min_tls_protocol_versionCallError(); err != nil {
		return *new(Tls_protocol_version_t), err
	}
	return _sec_protocol_options_get_default_min_tls_protocol_version(), nil
}

// Sec_protocol_options_get_default_min_tls_protocol_version.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_get_default_min_tls_protocol_version()
func Sec_protocol_options_get_default_min_tls_protocol_version() Tls_protocol_version_t {
	result, callErr := TrySec_protocol_options_get_default_min_tls_protocol_version()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_options_get_enable_encrypted_client_hello func(options Sec_protocol_options_t) bool
var _sec_protocol_options_get_enable_encrypted_client_helloErr error

// CanCallSec_protocol_options_get_enable_encrypted_client_hello reports whether the symbol for sec_protocol_options_get_enable_encrypted_client_hello is available on this system.
func CanCallSec_protocol_options_get_enable_encrypted_client_hello() bool {
	return _sec_protocol_options_get_enable_encrypted_client_hello != nil
}

// Sec_protocol_options_get_enable_encrypted_client_helloCallError returns the symbol lookup or registration error for sec_protocol_options_get_enable_encrypted_client_hello.
func Sec_protocol_options_get_enable_encrypted_client_helloCallError() error {
	if _sec_protocol_options_get_enable_encrypted_client_hello != nil {
		return nil
	}
	return symbolCallError("sec_protocol_options_get_enable_encrypted_client_hello", "", _sec_protocol_options_get_enable_encrypted_client_helloErr)
}

// TrySec_protocol_options_get_enable_encrypted_client_hello calls Sec_protocol_options_get_enable_encrypted_client_hello without panicking when the symbol is unavailable.
func TrySec_protocol_options_get_enable_encrypted_client_hello(options Sec_protocol_options_t) (bool, error) {
	if err := Sec_protocol_options_get_enable_encrypted_client_helloCallError(); err != nil {
		return false, err
	}
	return _sec_protocol_options_get_enable_encrypted_client_hello(options), nil
}

// Sec_protocol_options_get_enable_encrypted_client_hello.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_get_enable_encrypted_client_hello
func Sec_protocol_options_get_enable_encrypted_client_hello(options Sec_protocol_options_t) bool {
	result, callErr := TrySec_protocol_options_get_enable_encrypted_client_hello(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_options_get_quic_use_legacy_codepoint func(options Sec_protocol_options_t) bool
var _sec_protocol_options_get_quic_use_legacy_codepointErr error

// CanCallSec_protocol_options_get_quic_use_legacy_codepoint reports whether the symbol for sec_protocol_options_get_quic_use_legacy_codepoint is available on this system.
func CanCallSec_protocol_options_get_quic_use_legacy_codepoint() bool {
	return _sec_protocol_options_get_quic_use_legacy_codepoint != nil
}

// Sec_protocol_options_get_quic_use_legacy_codepointCallError returns the symbol lookup or registration error for sec_protocol_options_get_quic_use_legacy_codepoint.
func Sec_protocol_options_get_quic_use_legacy_codepointCallError() error {
	if _sec_protocol_options_get_quic_use_legacy_codepoint != nil {
		return nil
	}
	return symbolCallError("sec_protocol_options_get_quic_use_legacy_codepoint", "", _sec_protocol_options_get_quic_use_legacy_codepointErr)
}

// TrySec_protocol_options_get_quic_use_legacy_codepoint calls Sec_protocol_options_get_quic_use_legacy_codepoint without panicking when the symbol is unavailable.
func TrySec_protocol_options_get_quic_use_legacy_codepoint(options Sec_protocol_options_t) (bool, error) {
	if err := Sec_protocol_options_get_quic_use_legacy_codepointCallError(); err != nil {
		return false, err
	}
	return _sec_protocol_options_get_quic_use_legacy_codepoint(options), nil
}

// Sec_protocol_options_get_quic_use_legacy_codepoint.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_get_quic_use_legacy_codepoint
func Sec_protocol_options_get_quic_use_legacy_codepoint(options Sec_protocol_options_t) bool {
	result, callErr := TrySec_protocol_options_get_quic_use_legacy_codepoint(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_protocol_options_set_challenge_block func(options Sec_protocol_options_t, challenge_block unsafe.Pointer, challenge_queue uintptr)
var _sec_protocol_options_set_challenge_blockErr error

// CanCallSec_protocol_options_set_challenge_block reports whether the symbol for sec_protocol_options_set_challenge_block is available on this system.
func CanCallSec_protocol_options_set_challenge_block() bool {
	return _sec_protocol_options_set_challenge_block != nil
}

// Sec_protocol_options_set_challenge_blockCallError returns the symbol lookup or registration error for sec_protocol_options_set_challenge_block.
func Sec_protocol_options_set_challenge_blockCallError() error {
	if _sec_protocol_options_set_challenge_block != nil {
		return nil
	}
	return symbolCallError("sec_protocol_options_set_challenge_block", "10.14", _sec_protocol_options_set_challenge_blockErr)
}

// TrySec_protocol_options_set_challenge_block calls Sec_protocol_options_set_challenge_block without panicking when the symbol is unavailable.
func TrySec_protocol_options_set_challenge_block(options Sec_protocol_options_t, challenge_block Sec_protocol_challenge_t, challenge_queue dispatch.Queue) error {
	if err := Sec_protocol_options_set_challenge_blockCallError(); err != nil {
		return err
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object, arg1 func(*objectivec.Object)) { challenge_block(arg0, arg1) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_sec_protocol_options_set_challenge_block(options, _block0, uintptr(challenge_queue.Handle()))
	return nil
}

// Sec_protocol_options_set_challenge_block.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_challenge_block(_:_:_:)
func Sec_protocol_options_set_challenge_block(options Sec_protocol_options_t, challenge_block Sec_protocol_challenge_t, challenge_queue dispatch.Queue) {
	if callErr := TrySec_protocol_options_set_challenge_block(options, challenge_block, challenge_queue); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_enable_encrypted_client_hello func(options Sec_protocol_options_t, enable_encrypted_client_hello bool)
var _sec_protocol_options_set_enable_encrypted_client_helloErr error

// CanCallSec_protocol_options_set_enable_encrypted_client_hello reports whether the symbol for sec_protocol_options_set_enable_encrypted_client_hello is available on this system.
func CanCallSec_protocol_options_set_enable_encrypted_client_hello() bool {
	return _sec_protocol_options_set_enable_encrypted_client_hello != nil
}

// Sec_protocol_options_set_enable_encrypted_client_helloCallError returns the symbol lookup or registration error for sec_protocol_options_set_enable_encrypted_client_hello.
func Sec_protocol_options_set_enable_encrypted_client_helloCallError() error {
	if _sec_protocol_options_set_enable_encrypted_client_hello != nil {
		return nil
	}
	return symbolCallError("sec_protocol_options_set_enable_encrypted_client_hello", "", _sec_protocol_options_set_enable_encrypted_client_helloErr)
}

// TrySec_protocol_options_set_enable_encrypted_client_hello calls Sec_protocol_options_set_enable_encrypted_client_hello without panicking when the symbol is unavailable.
func TrySec_protocol_options_set_enable_encrypted_client_hello(options Sec_protocol_options_t, enable_encrypted_client_hello bool) error {
	if err := Sec_protocol_options_set_enable_encrypted_client_helloCallError(); err != nil {
		return err
	}
	_sec_protocol_options_set_enable_encrypted_client_hello(options, enable_encrypted_client_hello)
	return nil
}

// Sec_protocol_options_set_enable_encrypted_client_hello.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_enable_encrypted_client_hello
func Sec_protocol_options_set_enable_encrypted_client_hello(options Sec_protocol_options_t, enable_encrypted_client_hello bool) {
	if callErr := TrySec_protocol_options_set_enable_encrypted_client_hello(options, enable_encrypted_client_hello); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_key_update_block func(options Sec_protocol_options_t, key_update_block unsafe.Pointer, key_update_queue uintptr)
var _sec_protocol_options_set_key_update_blockErr error

// CanCallSec_protocol_options_set_key_update_block reports whether the symbol for sec_protocol_options_set_key_update_block is available on this system.
func CanCallSec_protocol_options_set_key_update_block() bool {
	return _sec_protocol_options_set_key_update_block != nil
}

// Sec_protocol_options_set_key_update_blockCallError returns the symbol lookup or registration error for sec_protocol_options_set_key_update_block.
func Sec_protocol_options_set_key_update_blockCallError() error {
	if _sec_protocol_options_set_key_update_block != nil {
		return nil
	}
	return symbolCallError("sec_protocol_options_set_key_update_block", "10.14", _sec_protocol_options_set_key_update_blockErr)
}

// TrySec_protocol_options_set_key_update_block calls Sec_protocol_options_set_key_update_block without panicking when the symbol is unavailable.
func TrySec_protocol_options_set_key_update_block(options Sec_protocol_options_t, key_update_block Sec_protocol_key_update_t, key_update_queue dispatch.Queue) error {
	if err := Sec_protocol_options_set_key_update_blockCallError(); err != nil {
		return err
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object, arg1 func()) { key_update_block(arg0, arg1) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_sec_protocol_options_set_key_update_block(options, _block0, uintptr(key_update_queue.Handle()))
	return nil
}

// Sec_protocol_options_set_key_update_block.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_key_update_block(_:_:_:)
func Sec_protocol_options_set_key_update_block(options Sec_protocol_options_t, key_update_block Sec_protocol_key_update_t, key_update_queue dispatch.Queue) {
	if callErr := TrySec_protocol_options_set_key_update_block(options, key_update_block, key_update_queue); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_local_identity func(options Sec_protocol_options_t, identity Sec_identity_t)
var _sec_protocol_options_set_local_identityErr error

// CanCallSec_protocol_options_set_local_identity reports whether the symbol for sec_protocol_options_set_local_identity is available on this system.
func CanCallSec_protocol_options_set_local_identity() bool {
	return _sec_protocol_options_set_local_identity != nil
}

// Sec_protocol_options_set_local_identityCallError returns the symbol lookup or registration error for sec_protocol_options_set_local_identity.
func Sec_protocol_options_set_local_identityCallError() error {
	if _sec_protocol_options_set_local_identity != nil {
		return nil
	}
	return symbolCallError("sec_protocol_options_set_local_identity", "10.14", _sec_protocol_options_set_local_identityErr)
}

// TrySec_protocol_options_set_local_identity calls Sec_protocol_options_set_local_identity without panicking when the symbol is unavailable.
func TrySec_protocol_options_set_local_identity(options Sec_protocol_options_t, identity Sec_identity_t) error {
	if err := Sec_protocol_options_set_local_identityCallError(); err != nil {
		return err
	}
	_sec_protocol_options_set_local_identity(options, identity)
	return nil
}

// Sec_protocol_options_set_local_identity.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_local_identity(_:_:)
func Sec_protocol_options_set_local_identity(options Sec_protocol_options_t, identity Sec_identity_t) {
	if callErr := TrySec_protocol_options_set_local_identity(options, identity); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_max_tls_protocol_version func(options Sec_protocol_options_t, version Tls_protocol_version_t)
var _sec_protocol_options_set_max_tls_protocol_versionErr error

// CanCallSec_protocol_options_set_max_tls_protocol_version reports whether the symbol for sec_protocol_options_set_max_tls_protocol_version is available on this system.
func CanCallSec_protocol_options_set_max_tls_protocol_version() bool {
	return _sec_protocol_options_set_max_tls_protocol_version != nil
}

// Sec_protocol_options_set_max_tls_protocol_versionCallError returns the symbol lookup or registration error for sec_protocol_options_set_max_tls_protocol_version.
func Sec_protocol_options_set_max_tls_protocol_versionCallError() error {
	if _sec_protocol_options_set_max_tls_protocol_version != nil {
		return nil
	}
	return symbolCallError("sec_protocol_options_set_max_tls_protocol_version", "10.15", _sec_protocol_options_set_max_tls_protocol_versionErr)
}

// TrySec_protocol_options_set_max_tls_protocol_version calls Sec_protocol_options_set_max_tls_protocol_version without panicking when the symbol is unavailable.
func TrySec_protocol_options_set_max_tls_protocol_version(options Sec_protocol_options_t, version Tls_protocol_version_t) error {
	if err := Sec_protocol_options_set_max_tls_protocol_versionCallError(); err != nil {
		return err
	}
	_sec_protocol_options_set_max_tls_protocol_version(options, version)
	return nil
}

// Sec_protocol_options_set_max_tls_protocol_version.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_max_tls_protocol_version(_:_:)
func Sec_protocol_options_set_max_tls_protocol_version(options Sec_protocol_options_t, version Tls_protocol_version_t) {
	if callErr := TrySec_protocol_options_set_max_tls_protocol_version(options, version); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_min_tls_protocol_version func(options Sec_protocol_options_t, version Tls_protocol_version_t)
var _sec_protocol_options_set_min_tls_protocol_versionErr error

// CanCallSec_protocol_options_set_min_tls_protocol_version reports whether the symbol for sec_protocol_options_set_min_tls_protocol_version is available on this system.
func CanCallSec_protocol_options_set_min_tls_protocol_version() bool {
	return _sec_protocol_options_set_min_tls_protocol_version != nil
}

// Sec_protocol_options_set_min_tls_protocol_versionCallError returns the symbol lookup or registration error for sec_protocol_options_set_min_tls_protocol_version.
func Sec_protocol_options_set_min_tls_protocol_versionCallError() error {
	if _sec_protocol_options_set_min_tls_protocol_version != nil {
		return nil
	}
	return symbolCallError("sec_protocol_options_set_min_tls_protocol_version", "10.15", _sec_protocol_options_set_min_tls_protocol_versionErr)
}

// TrySec_protocol_options_set_min_tls_protocol_version calls Sec_protocol_options_set_min_tls_protocol_version without panicking when the symbol is unavailable.
func TrySec_protocol_options_set_min_tls_protocol_version(options Sec_protocol_options_t, version Tls_protocol_version_t) error {
	if err := Sec_protocol_options_set_min_tls_protocol_versionCallError(); err != nil {
		return err
	}
	_sec_protocol_options_set_min_tls_protocol_version(options, version)
	return nil
}

// Sec_protocol_options_set_min_tls_protocol_version.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_min_tls_protocol_version(_:_:)
func Sec_protocol_options_set_min_tls_protocol_version(options Sec_protocol_options_t, version Tls_protocol_version_t) {
	if callErr := TrySec_protocol_options_set_min_tls_protocol_version(options, version); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_peer_authentication_optional func(options Sec_protocol_options_t, peer_authentication_optional bool)
var _sec_protocol_options_set_peer_authentication_optionalErr error

// CanCallSec_protocol_options_set_peer_authentication_optional reports whether the symbol for sec_protocol_options_set_peer_authentication_optional is available on this system.
func CanCallSec_protocol_options_set_peer_authentication_optional() bool {
	return _sec_protocol_options_set_peer_authentication_optional != nil
}

// Sec_protocol_options_set_peer_authentication_optionalCallError returns the symbol lookup or registration error for sec_protocol_options_set_peer_authentication_optional.
func Sec_protocol_options_set_peer_authentication_optionalCallError() error {
	if _sec_protocol_options_set_peer_authentication_optional != nil {
		return nil
	}
	return symbolCallError("sec_protocol_options_set_peer_authentication_optional", "", _sec_protocol_options_set_peer_authentication_optionalErr)
}

// TrySec_protocol_options_set_peer_authentication_optional calls Sec_protocol_options_set_peer_authentication_optional without panicking when the symbol is unavailable.
func TrySec_protocol_options_set_peer_authentication_optional(options Sec_protocol_options_t, peer_authentication_optional bool) error {
	if err := Sec_protocol_options_set_peer_authentication_optionalCallError(); err != nil {
		return err
	}
	_sec_protocol_options_set_peer_authentication_optional(options, peer_authentication_optional)
	return nil
}

// Sec_protocol_options_set_peer_authentication_optional.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_peer_authentication_optional
func Sec_protocol_options_set_peer_authentication_optional(options Sec_protocol_options_t, peer_authentication_optional bool) {
	if callErr := TrySec_protocol_options_set_peer_authentication_optional(options, peer_authentication_optional); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_peer_authentication_required func(options Sec_protocol_options_t, peer_authentication_required bool)
var _sec_protocol_options_set_peer_authentication_requiredErr error

// CanCallSec_protocol_options_set_peer_authentication_required reports whether the symbol for sec_protocol_options_set_peer_authentication_required is available on this system.
func CanCallSec_protocol_options_set_peer_authentication_required() bool {
	return _sec_protocol_options_set_peer_authentication_required != nil
}

// Sec_protocol_options_set_peer_authentication_requiredCallError returns the symbol lookup or registration error for sec_protocol_options_set_peer_authentication_required.
func Sec_protocol_options_set_peer_authentication_requiredCallError() error {
	if _sec_protocol_options_set_peer_authentication_required != nil {
		return nil
	}
	return symbolCallError("sec_protocol_options_set_peer_authentication_required", "10.14", _sec_protocol_options_set_peer_authentication_requiredErr)
}

// TrySec_protocol_options_set_peer_authentication_required calls Sec_protocol_options_set_peer_authentication_required without panicking when the symbol is unavailable.
func TrySec_protocol_options_set_peer_authentication_required(options Sec_protocol_options_t, peer_authentication_required bool) error {
	if err := Sec_protocol_options_set_peer_authentication_requiredCallError(); err != nil {
		return err
	}
	_sec_protocol_options_set_peer_authentication_required(options, peer_authentication_required)
	return nil
}

// Sec_protocol_options_set_peer_authentication_required.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_peer_authentication_required(_:_:)
func Sec_protocol_options_set_peer_authentication_required(options Sec_protocol_options_t, peer_authentication_required bool) {
	if callErr := TrySec_protocol_options_set_peer_authentication_required(options, peer_authentication_required); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_pre_shared_key_selection_block func(options Sec_protocol_options_t, psk_selection_block unsafe.Pointer, psk_selection_queue uintptr)
var _sec_protocol_options_set_pre_shared_key_selection_blockErr error

// CanCallSec_protocol_options_set_pre_shared_key_selection_block reports whether the symbol for sec_protocol_options_set_pre_shared_key_selection_block is available on this system.
func CanCallSec_protocol_options_set_pre_shared_key_selection_block() bool {
	return _sec_protocol_options_set_pre_shared_key_selection_block != nil
}

// Sec_protocol_options_set_pre_shared_key_selection_blockCallError returns the symbol lookup or registration error for sec_protocol_options_set_pre_shared_key_selection_block.
func Sec_protocol_options_set_pre_shared_key_selection_blockCallError() error {
	if _sec_protocol_options_set_pre_shared_key_selection_block != nil {
		return nil
	}
	return symbolCallError("sec_protocol_options_set_pre_shared_key_selection_block", "10.15", _sec_protocol_options_set_pre_shared_key_selection_blockErr)
}

// TrySec_protocol_options_set_pre_shared_key_selection_block calls Sec_protocol_options_set_pre_shared_key_selection_block without panicking when the symbol is unavailable.
func TrySec_protocol_options_set_pre_shared_key_selection_block(options Sec_protocol_options_t, psk_selection_block Sec_protocol_pre_shared_key_selection_t, psk_selection_queue dispatch.Queue) error {
	if err := Sec_protocol_options_set_pre_shared_key_selection_blockCallError(); err != nil {
		return err
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object, arg1 objectivec.Object, arg2 func(*objectivec.Object)) {
		psk_selection_block(arg0, arg1, arg2)
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
	if callErr := TrySec_protocol_options_set_pre_shared_key_selection_block(options, psk_selection_block, psk_selection_queue); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_quic_use_legacy_codepoint func(options Sec_protocol_options_t, quic_use_legacy_codepoint bool)
var _sec_protocol_options_set_quic_use_legacy_codepointErr error

// CanCallSec_protocol_options_set_quic_use_legacy_codepoint reports whether the symbol for sec_protocol_options_set_quic_use_legacy_codepoint is available on this system.
func CanCallSec_protocol_options_set_quic_use_legacy_codepoint() bool {
	return _sec_protocol_options_set_quic_use_legacy_codepoint != nil
}

// Sec_protocol_options_set_quic_use_legacy_codepointCallError returns the symbol lookup or registration error for sec_protocol_options_set_quic_use_legacy_codepoint.
func Sec_protocol_options_set_quic_use_legacy_codepointCallError() error {
	if _sec_protocol_options_set_quic_use_legacy_codepoint != nil {
		return nil
	}
	return symbolCallError("sec_protocol_options_set_quic_use_legacy_codepoint", "", _sec_protocol_options_set_quic_use_legacy_codepointErr)
}

// TrySec_protocol_options_set_quic_use_legacy_codepoint calls Sec_protocol_options_set_quic_use_legacy_codepoint without panicking when the symbol is unavailable.
func TrySec_protocol_options_set_quic_use_legacy_codepoint(options Sec_protocol_options_t, quic_use_legacy_codepoint bool) error {
	if err := Sec_protocol_options_set_quic_use_legacy_codepointCallError(); err != nil {
		return err
	}
	_sec_protocol_options_set_quic_use_legacy_codepoint(options, quic_use_legacy_codepoint)
	return nil
}

// Sec_protocol_options_set_quic_use_legacy_codepoint.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_quic_use_legacy_codepoint
func Sec_protocol_options_set_quic_use_legacy_codepoint(options Sec_protocol_options_t, quic_use_legacy_codepoint bool) {
	if callErr := TrySec_protocol_options_set_quic_use_legacy_codepoint(options, quic_use_legacy_codepoint); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_tls_false_start_enabled func(options Sec_protocol_options_t, false_start_enabled bool)
var _sec_protocol_options_set_tls_false_start_enabledErr error

// CanCallSec_protocol_options_set_tls_false_start_enabled reports whether the symbol for sec_protocol_options_set_tls_false_start_enabled is available on this system.
func CanCallSec_protocol_options_set_tls_false_start_enabled() bool {
	return _sec_protocol_options_set_tls_false_start_enabled != nil
}

// Sec_protocol_options_set_tls_false_start_enabledCallError returns the symbol lookup or registration error for sec_protocol_options_set_tls_false_start_enabled.
func Sec_protocol_options_set_tls_false_start_enabledCallError() error {
	if _sec_protocol_options_set_tls_false_start_enabled != nil {
		return nil
	}
	return symbolCallError("sec_protocol_options_set_tls_false_start_enabled", "10.14", _sec_protocol_options_set_tls_false_start_enabledErr)
}

// TrySec_protocol_options_set_tls_false_start_enabled calls Sec_protocol_options_set_tls_false_start_enabled without panicking when the symbol is unavailable.
func TrySec_protocol_options_set_tls_false_start_enabled(options Sec_protocol_options_t, false_start_enabled bool) error {
	if err := Sec_protocol_options_set_tls_false_start_enabledCallError(); err != nil {
		return err
	}
	_sec_protocol_options_set_tls_false_start_enabled(options, false_start_enabled)
	return nil
}

// Sec_protocol_options_set_tls_false_start_enabled.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_false_start_enabled(_:_:)
func Sec_protocol_options_set_tls_false_start_enabled(options Sec_protocol_options_t, false_start_enabled bool) {
	if callErr := TrySec_protocol_options_set_tls_false_start_enabled(options, false_start_enabled); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_tls_is_fallback_attempt func(options Sec_protocol_options_t, is_fallback_attempt bool)
var _sec_protocol_options_set_tls_is_fallback_attemptErr error

// CanCallSec_protocol_options_set_tls_is_fallback_attempt reports whether the symbol for sec_protocol_options_set_tls_is_fallback_attempt is available on this system.
func CanCallSec_protocol_options_set_tls_is_fallback_attempt() bool {
	return _sec_protocol_options_set_tls_is_fallback_attempt != nil
}

// Sec_protocol_options_set_tls_is_fallback_attemptCallError returns the symbol lookup or registration error for sec_protocol_options_set_tls_is_fallback_attempt.
func Sec_protocol_options_set_tls_is_fallback_attemptCallError() error {
	if _sec_protocol_options_set_tls_is_fallback_attempt != nil {
		return nil
	}
	return symbolCallError("sec_protocol_options_set_tls_is_fallback_attempt", "10.14", _sec_protocol_options_set_tls_is_fallback_attemptErr)
}

// TrySec_protocol_options_set_tls_is_fallback_attempt calls Sec_protocol_options_set_tls_is_fallback_attempt without panicking when the symbol is unavailable.
func TrySec_protocol_options_set_tls_is_fallback_attempt(options Sec_protocol_options_t, is_fallback_attempt bool) error {
	if err := Sec_protocol_options_set_tls_is_fallback_attemptCallError(); err != nil {
		return err
	}
	_sec_protocol_options_set_tls_is_fallback_attempt(options, is_fallback_attempt)
	return nil
}

// Sec_protocol_options_set_tls_is_fallback_attempt.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_is_fallback_attempt(_:_:)
func Sec_protocol_options_set_tls_is_fallback_attempt(options Sec_protocol_options_t, is_fallback_attempt bool) {
	if callErr := TrySec_protocol_options_set_tls_is_fallback_attempt(options, is_fallback_attempt); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_tls_max_version func(options Sec_protocol_options_t, version SSLProtocol)
var _sec_protocol_options_set_tls_max_versionErr error

// CanCallSec_protocol_options_set_tls_max_version reports whether the symbol for sec_protocol_options_set_tls_max_version is available on this system.
func CanCallSec_protocol_options_set_tls_max_version() bool {
	return _sec_protocol_options_set_tls_max_version != nil
}

// Sec_protocol_options_set_tls_max_versionCallError returns the symbol lookup or registration error for sec_protocol_options_set_tls_max_version.
func Sec_protocol_options_set_tls_max_versionCallError() error {
	if _sec_protocol_options_set_tls_max_version != nil {
		return nil
	}
	return symbolCallError("sec_protocol_options_set_tls_max_version", "10.14", _sec_protocol_options_set_tls_max_versionErr)
}

// TrySec_protocol_options_set_tls_max_version calls Sec_protocol_options_set_tls_max_version without panicking when the symbol is unavailable.
func TrySec_protocol_options_set_tls_max_version(options Sec_protocol_options_t, version SSLProtocol) error {
	if err := Sec_protocol_options_set_tls_max_versionCallError(); err != nil {
		return err
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
	if callErr := TrySec_protocol_options_set_tls_max_version(options, version); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_tls_min_version func(options Sec_protocol_options_t, version SSLProtocol)
var _sec_protocol_options_set_tls_min_versionErr error

// CanCallSec_protocol_options_set_tls_min_version reports whether the symbol for sec_protocol_options_set_tls_min_version is available on this system.
func CanCallSec_protocol_options_set_tls_min_version() bool {
	return _sec_protocol_options_set_tls_min_version != nil
}

// Sec_protocol_options_set_tls_min_versionCallError returns the symbol lookup or registration error for sec_protocol_options_set_tls_min_version.
func Sec_protocol_options_set_tls_min_versionCallError() error {
	if _sec_protocol_options_set_tls_min_version != nil {
		return nil
	}
	return symbolCallError("sec_protocol_options_set_tls_min_version", "10.14", _sec_protocol_options_set_tls_min_versionErr)
}

// TrySec_protocol_options_set_tls_min_version calls Sec_protocol_options_set_tls_min_version without panicking when the symbol is unavailable.
func TrySec_protocol_options_set_tls_min_version(options Sec_protocol_options_t, version SSLProtocol) error {
	if err := Sec_protocol_options_set_tls_min_versionCallError(); err != nil {
		return err
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
	if callErr := TrySec_protocol_options_set_tls_min_version(options, version); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_tls_ocsp_enabled func(options Sec_protocol_options_t, ocsp_enabled bool)
var _sec_protocol_options_set_tls_ocsp_enabledErr error

// CanCallSec_protocol_options_set_tls_ocsp_enabled reports whether the symbol for sec_protocol_options_set_tls_ocsp_enabled is available on this system.
func CanCallSec_protocol_options_set_tls_ocsp_enabled() bool {
	return _sec_protocol_options_set_tls_ocsp_enabled != nil
}

// Sec_protocol_options_set_tls_ocsp_enabledCallError returns the symbol lookup or registration error for sec_protocol_options_set_tls_ocsp_enabled.
func Sec_protocol_options_set_tls_ocsp_enabledCallError() error {
	if _sec_protocol_options_set_tls_ocsp_enabled != nil {
		return nil
	}
	return symbolCallError("sec_protocol_options_set_tls_ocsp_enabled", "10.14", _sec_protocol_options_set_tls_ocsp_enabledErr)
}

// TrySec_protocol_options_set_tls_ocsp_enabled calls Sec_protocol_options_set_tls_ocsp_enabled without panicking when the symbol is unavailable.
func TrySec_protocol_options_set_tls_ocsp_enabled(options Sec_protocol_options_t, ocsp_enabled bool) error {
	if err := Sec_protocol_options_set_tls_ocsp_enabledCallError(); err != nil {
		return err
	}
	_sec_protocol_options_set_tls_ocsp_enabled(options, ocsp_enabled)
	return nil
}

// Sec_protocol_options_set_tls_ocsp_enabled.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_ocsp_enabled(_:_:)
func Sec_protocol_options_set_tls_ocsp_enabled(options Sec_protocol_options_t, ocsp_enabled bool) {
	if callErr := TrySec_protocol_options_set_tls_ocsp_enabled(options, ocsp_enabled); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_tls_pre_shared_key_identity_hint func(options Sec_protocol_options_t, psk_identity_hint uintptr)
var _sec_protocol_options_set_tls_pre_shared_key_identity_hintErr error

// CanCallSec_protocol_options_set_tls_pre_shared_key_identity_hint reports whether the symbol for sec_protocol_options_set_tls_pre_shared_key_identity_hint is available on this system.
func CanCallSec_protocol_options_set_tls_pre_shared_key_identity_hint() bool {
	return _sec_protocol_options_set_tls_pre_shared_key_identity_hint != nil
}

// Sec_protocol_options_set_tls_pre_shared_key_identity_hintCallError returns the symbol lookup or registration error for sec_protocol_options_set_tls_pre_shared_key_identity_hint.
func Sec_protocol_options_set_tls_pre_shared_key_identity_hintCallError() error {
	if _sec_protocol_options_set_tls_pre_shared_key_identity_hint != nil {
		return nil
	}
	return symbolCallError("sec_protocol_options_set_tls_pre_shared_key_identity_hint", "10.15", _sec_protocol_options_set_tls_pre_shared_key_identity_hintErr)
}

// TrySec_protocol_options_set_tls_pre_shared_key_identity_hint calls Sec_protocol_options_set_tls_pre_shared_key_identity_hint without panicking when the symbol is unavailable.
func TrySec_protocol_options_set_tls_pre_shared_key_identity_hint(options Sec_protocol_options_t, psk_identity_hint dispatch.Data) error {
	if err := Sec_protocol_options_set_tls_pre_shared_key_identity_hintCallError(); err != nil {
		return err
	}
	_sec_protocol_options_set_tls_pre_shared_key_identity_hint(options, uintptr(psk_identity_hint.Handle()))
	return nil
}

// Sec_protocol_options_set_tls_pre_shared_key_identity_hint.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_pre_shared_key_identity_hint(_:_:)
func Sec_protocol_options_set_tls_pre_shared_key_identity_hint(options Sec_protocol_options_t, psk_identity_hint dispatch.Data) {
	if callErr := TrySec_protocol_options_set_tls_pre_shared_key_identity_hint(options, psk_identity_hint); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_tls_renegotiation_enabled func(options Sec_protocol_options_t, renegotiation_enabled bool)
var _sec_protocol_options_set_tls_renegotiation_enabledErr error

// CanCallSec_protocol_options_set_tls_renegotiation_enabled reports whether the symbol for sec_protocol_options_set_tls_renegotiation_enabled is available on this system.
func CanCallSec_protocol_options_set_tls_renegotiation_enabled() bool {
	return _sec_protocol_options_set_tls_renegotiation_enabled != nil
}

// Sec_protocol_options_set_tls_renegotiation_enabledCallError returns the symbol lookup or registration error for sec_protocol_options_set_tls_renegotiation_enabled.
func Sec_protocol_options_set_tls_renegotiation_enabledCallError() error {
	if _sec_protocol_options_set_tls_renegotiation_enabled != nil {
		return nil
	}
	return symbolCallError("sec_protocol_options_set_tls_renegotiation_enabled", "10.14", _sec_protocol_options_set_tls_renegotiation_enabledErr)
}

// TrySec_protocol_options_set_tls_renegotiation_enabled calls Sec_protocol_options_set_tls_renegotiation_enabled without panicking when the symbol is unavailable.
func TrySec_protocol_options_set_tls_renegotiation_enabled(options Sec_protocol_options_t, renegotiation_enabled bool) error {
	if err := Sec_protocol_options_set_tls_renegotiation_enabledCallError(); err != nil {
		return err
	}
	_sec_protocol_options_set_tls_renegotiation_enabled(options, renegotiation_enabled)
	return nil
}

// Sec_protocol_options_set_tls_renegotiation_enabled.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_renegotiation_enabled(_:_:)
func Sec_protocol_options_set_tls_renegotiation_enabled(options Sec_protocol_options_t, renegotiation_enabled bool) {
	if callErr := TrySec_protocol_options_set_tls_renegotiation_enabled(options, renegotiation_enabled); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_tls_resumption_enabled func(options Sec_protocol_options_t, resumption_enabled bool)
var _sec_protocol_options_set_tls_resumption_enabledErr error

// CanCallSec_protocol_options_set_tls_resumption_enabled reports whether the symbol for sec_protocol_options_set_tls_resumption_enabled is available on this system.
func CanCallSec_protocol_options_set_tls_resumption_enabled() bool {
	return _sec_protocol_options_set_tls_resumption_enabled != nil
}

// Sec_protocol_options_set_tls_resumption_enabledCallError returns the symbol lookup or registration error for sec_protocol_options_set_tls_resumption_enabled.
func Sec_protocol_options_set_tls_resumption_enabledCallError() error {
	if _sec_protocol_options_set_tls_resumption_enabled != nil {
		return nil
	}
	return symbolCallError("sec_protocol_options_set_tls_resumption_enabled", "10.14", _sec_protocol_options_set_tls_resumption_enabledErr)
}

// TrySec_protocol_options_set_tls_resumption_enabled calls Sec_protocol_options_set_tls_resumption_enabled without panicking when the symbol is unavailable.
func TrySec_protocol_options_set_tls_resumption_enabled(options Sec_protocol_options_t, resumption_enabled bool) error {
	if err := Sec_protocol_options_set_tls_resumption_enabledCallError(); err != nil {
		return err
	}
	_sec_protocol_options_set_tls_resumption_enabled(options, resumption_enabled)
	return nil
}

// Sec_protocol_options_set_tls_resumption_enabled.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_resumption_enabled(_:_:)
func Sec_protocol_options_set_tls_resumption_enabled(options Sec_protocol_options_t, resumption_enabled bool) {
	if callErr := TrySec_protocol_options_set_tls_resumption_enabled(options, resumption_enabled); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_tls_sct_enabled func(options Sec_protocol_options_t, sct_enabled bool)
var _sec_protocol_options_set_tls_sct_enabledErr error

// CanCallSec_protocol_options_set_tls_sct_enabled reports whether the symbol for sec_protocol_options_set_tls_sct_enabled is available on this system.
func CanCallSec_protocol_options_set_tls_sct_enabled() bool {
	return _sec_protocol_options_set_tls_sct_enabled != nil
}

// Sec_protocol_options_set_tls_sct_enabledCallError returns the symbol lookup or registration error for sec_protocol_options_set_tls_sct_enabled.
func Sec_protocol_options_set_tls_sct_enabledCallError() error {
	if _sec_protocol_options_set_tls_sct_enabled != nil {
		return nil
	}
	return symbolCallError("sec_protocol_options_set_tls_sct_enabled", "10.14", _sec_protocol_options_set_tls_sct_enabledErr)
}

// TrySec_protocol_options_set_tls_sct_enabled calls Sec_protocol_options_set_tls_sct_enabled without panicking when the symbol is unavailable.
func TrySec_protocol_options_set_tls_sct_enabled(options Sec_protocol_options_t, sct_enabled bool) error {
	if err := Sec_protocol_options_set_tls_sct_enabledCallError(); err != nil {
		return err
	}
	_sec_protocol_options_set_tls_sct_enabled(options, sct_enabled)
	return nil
}

// Sec_protocol_options_set_tls_sct_enabled.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_sct_enabled(_:_:)
func Sec_protocol_options_set_tls_sct_enabled(options Sec_protocol_options_t, sct_enabled bool) {
	if callErr := TrySec_protocol_options_set_tls_sct_enabled(options, sct_enabled); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_tls_server_name func(options Sec_protocol_options_t, server_name string)
var _sec_protocol_options_set_tls_server_nameErr error

// CanCallSec_protocol_options_set_tls_server_name reports whether the symbol for sec_protocol_options_set_tls_server_name is available on this system.
func CanCallSec_protocol_options_set_tls_server_name() bool {
	return _sec_protocol_options_set_tls_server_name != nil
}

// Sec_protocol_options_set_tls_server_nameCallError returns the symbol lookup or registration error for sec_protocol_options_set_tls_server_name.
func Sec_protocol_options_set_tls_server_nameCallError() error {
	if _sec_protocol_options_set_tls_server_name != nil {
		return nil
	}
	return symbolCallError("sec_protocol_options_set_tls_server_name", "10.14", _sec_protocol_options_set_tls_server_nameErr)
}

// TrySec_protocol_options_set_tls_server_name calls Sec_protocol_options_set_tls_server_name without panicking when the symbol is unavailable.
func TrySec_protocol_options_set_tls_server_name(options Sec_protocol_options_t, server_name string) error {
	if err := Sec_protocol_options_set_tls_server_nameCallError(); err != nil {
		return err
	}
	_sec_protocol_options_set_tls_server_name(options, server_name)
	return nil
}

// Sec_protocol_options_set_tls_server_name.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_server_name(_:_:)
func Sec_protocol_options_set_tls_server_name(options Sec_protocol_options_t, server_name string) {
	if callErr := TrySec_protocol_options_set_tls_server_name(options, server_name); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_tls_tickets_enabled func(options Sec_protocol_options_t, tickets_enabled bool)
var _sec_protocol_options_set_tls_tickets_enabledErr error

// CanCallSec_protocol_options_set_tls_tickets_enabled reports whether the symbol for sec_protocol_options_set_tls_tickets_enabled is available on this system.
func CanCallSec_protocol_options_set_tls_tickets_enabled() bool {
	return _sec_protocol_options_set_tls_tickets_enabled != nil
}

// Sec_protocol_options_set_tls_tickets_enabledCallError returns the symbol lookup or registration error for sec_protocol_options_set_tls_tickets_enabled.
func Sec_protocol_options_set_tls_tickets_enabledCallError() error {
	if _sec_protocol_options_set_tls_tickets_enabled != nil {
		return nil
	}
	return symbolCallError("sec_protocol_options_set_tls_tickets_enabled", "10.14", _sec_protocol_options_set_tls_tickets_enabledErr)
}

// TrySec_protocol_options_set_tls_tickets_enabled calls Sec_protocol_options_set_tls_tickets_enabled without panicking when the symbol is unavailable.
func TrySec_protocol_options_set_tls_tickets_enabled(options Sec_protocol_options_t, tickets_enabled bool) error {
	if err := Sec_protocol_options_set_tls_tickets_enabledCallError(); err != nil {
		return err
	}
	_sec_protocol_options_set_tls_tickets_enabled(options, tickets_enabled)
	return nil
}

// Sec_protocol_options_set_tls_tickets_enabled.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_tickets_enabled(_:_:)
func Sec_protocol_options_set_tls_tickets_enabled(options Sec_protocol_options_t, tickets_enabled bool) {
	if callErr := TrySec_protocol_options_set_tls_tickets_enabled(options, tickets_enabled); callErr != nil {
		panic(callErr)
	}
}

var _sec_protocol_options_set_verify_block func(options Sec_protocol_options_t, verify_block unsafe.Pointer, verify_block_queue uintptr)
var _sec_protocol_options_set_verify_blockErr error

// CanCallSec_protocol_options_set_verify_block reports whether the symbol for sec_protocol_options_set_verify_block is available on this system.
func CanCallSec_protocol_options_set_verify_block() bool {
	return _sec_protocol_options_set_verify_block != nil
}

// Sec_protocol_options_set_verify_blockCallError returns the symbol lookup or registration error for sec_protocol_options_set_verify_block.
func Sec_protocol_options_set_verify_blockCallError() error {
	if _sec_protocol_options_set_verify_block != nil {
		return nil
	}
	return symbolCallError("sec_protocol_options_set_verify_block", "10.14", _sec_protocol_options_set_verify_blockErr)
}

// TrySec_protocol_options_set_verify_block calls Sec_protocol_options_set_verify_block without panicking when the symbol is unavailable.
func TrySec_protocol_options_set_verify_block(options Sec_protocol_options_t, verify_block Sec_protocol_verify_t, verify_block_queue dispatch.Queue) error {
	if err := Sec_protocol_options_set_verify_blockCallError(); err != nil {
		return err
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.Object, arg1 objectivec.Object, arg2 func(bool)) {
		verify_block(arg0, arg1, arg2)
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
	if callErr := TrySec_protocol_options_set_verify_block(options, verify_block, verify_block_queue); callErr != nil {
		panic(callErr)
	}
}

var _sec_release func(obj unsafe.Pointer)
var _sec_releaseErr error

// CanCallSec_release reports whether the symbol for sec_release is available on this system.
func CanCallSec_release() bool {
	return _sec_release != nil
}

// Sec_releaseCallError returns the symbol lookup or registration error for sec_release.
func Sec_releaseCallError() error {
	if _sec_release != nil {
		return nil
	}
	return symbolCallError("sec_release", "10.0", _sec_releaseErr)
}

// TrySec_release calls Sec_release without panicking when the symbol is unavailable.
func TrySec_release(obj unsafe.Pointer) error {
	if err := Sec_releaseCallError(); err != nil {
		return err
	}
	_sec_release(obj)
	return nil
}

// Sec_release.
//
// See: https://developer.apple.com/documentation/Security/sec_release(_:)
func Sec_release(obj unsafe.Pointer) {
	if callErr := TrySec_release(obj); callErr != nil {
		panic(callErr)
	}
}

var _sec_retain func(obj unsafe.Pointer) unsafe.Pointer
var _sec_retainErr error

// CanCallSec_retain reports whether the symbol for sec_retain is available on this system.
func CanCallSec_retain() bool {
	return _sec_retain != nil
}

// Sec_retainCallError returns the symbol lookup or registration error for sec_retain.
func Sec_retainCallError() error {
	if _sec_retain != nil {
		return nil
	}
	return symbolCallError("sec_retain", "10.0", _sec_retainErr)
}

// TrySec_retain calls Sec_retain without panicking when the symbol is unavailable.
func TrySec_retain(obj unsafe.Pointer) (unsafe.Pointer, error) {
	if err := Sec_retainCallError(); err != nil {
		return nil, err
	}
	return _sec_retain(obj), nil
}

// Sec_retain.
//
// See: https://developer.apple.com/documentation/Security/sec_retain(_:)
func Sec_retain(obj unsafe.Pointer) unsafe.Pointer {
	result, callErr := TrySec_retain(obj)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_trust_copy_ref func(trust Sec_trust_t) SecTrustRef
var _sec_trust_copy_refErr error

// CanCallSec_trust_copy_ref reports whether the symbol for sec_trust_copy_ref is available on this system.
func CanCallSec_trust_copy_ref() bool {
	return _sec_trust_copy_ref != nil
}

// Sec_trust_copy_refCallError returns the symbol lookup or registration error for sec_trust_copy_ref.
func Sec_trust_copy_refCallError() error {
	if _sec_trust_copy_ref != nil {
		return nil
	}
	return symbolCallError("sec_trust_copy_ref", "10.14", _sec_trust_copy_refErr)
}

// TrySec_trust_copy_ref calls Sec_trust_copy_ref without panicking when the symbol is unavailable.
func TrySec_trust_copy_ref(trust Sec_trust_t) (SecTrustRef, error) {
	if err := Sec_trust_copy_refCallError(); err != nil {
		return 0, err
	}
	return _sec_trust_copy_ref(trust), nil
}

// Sec_trust_copy_ref.
//
// See: https://developer.apple.com/documentation/Security/sec_trust_copy_ref(_:)
func Sec_trust_copy_ref(trust Sec_trust_t) SecTrustRef {
	result, callErr := TrySec_trust_copy_ref(trust)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sec_trust_create func(trust SecTrustRef) Sec_trust_t
var _sec_trust_createErr error

// CanCallSec_trust_create reports whether the symbol for sec_trust_create is available on this system.
func CanCallSec_trust_create() bool {
	return _sec_trust_create != nil
}

// Sec_trust_createCallError returns the symbol lookup or registration error for sec_trust_create.
func Sec_trust_createCallError() error {
	if _sec_trust_create != nil {
		return nil
	}
	return symbolCallError("sec_trust_create", "10.14", _sec_trust_createErr)
}

// TrySec_trust_create calls Sec_trust_create without panicking when the symbol is unavailable.
func TrySec_trust_create(trust SecTrustRef) (Sec_trust_t, error) {
	if err := Sec_trust_createCallError(); err != nil {
		return *new(Sec_trust_t), err
	}
	return _sec_trust_create(trust), nil
}

// Sec_trust_create.
//
// See: https://developer.apple.com/documentation/Security/sec_trust_create(_:)
func Sec_trust_create(trust SecTrustRef) Sec_trust_t {
	result, callErr := TrySec_trust_create(trust)
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
