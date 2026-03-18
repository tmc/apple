// Code generated from Apple documentation for Security. DO NOT EDIT.

package security

import (
	"fmt"
	"os"
	"unsafe"
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/dispatch"
)

// registerFunc resolves a framework symbol and registers it as a Go function.
// If the symbol is not found, a warning is printed and the function pointer is left nil.
func registerFunc(fptr any, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: Security: symbol %s not found\n", name)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "warning: Security: symbol %s registration skipped: %v\n", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: Security: symbol %s not found\n", name)
		return
	}
	*dst = sym
}

var _authorizationCopyInfo func(authorization AuthorizationRef, tag AuthorizationString, info *AuthorizationItemSet) int32

// AuthorizationCopyInfo retrieves supporting data such as the user name and other information gathered during evaluation of authorization.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationCopyInfo(_:_:_:)
func AuthorizationCopyInfo(authorization AuthorizationRef, tag AuthorizationString, info *AuthorizationItemSet) int32 {
	if _authorizationCopyInfo == nil {
		panic("Security: symbol AuthorizationCopyInfo not loaded")
	}
	return _authorizationCopyInfo(authorization, tag, info)
}

var _authorizationCopyRights func(authorization AuthorizationRef, rights *AuthorizationRights, environment *AuthorizationEnvironment, flags AuthorizationFlags, authorizedRights *AuthorizationRights) int32

// AuthorizationCopyRights authorizes and preauthorizes rights synchronously.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationCopyRights(_:_:_:_:_:)
func AuthorizationCopyRights(authorization AuthorizationRef, rights *AuthorizationRights, environment *AuthorizationEnvironment, flags AuthorizationFlags, authorizedRights *AuthorizationRights) int32 {
	if _authorizationCopyRights == nil {
		panic("Security: symbol AuthorizationCopyRights not loaded")
	}
	return _authorizationCopyRights(authorization, rights, environment, flags, authorizedRights)
}

var _authorizationCopyRightsAsync func(authorization AuthorizationRef, rights *AuthorizationRights, environment *AuthorizationEnvironment, flags AuthorizationFlags, callbackBlock AuthorizationAsyncCallback)

// AuthorizationCopyRightsAsync authorizes and preauthorizes rights asynchronously.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationCopyRightsAsync(_:_:_:_:_:)
func AuthorizationCopyRightsAsync(authorization AuthorizationRef, rights *AuthorizationRights, environment *AuthorizationEnvironment, flags AuthorizationFlags, callbackBlock AuthorizationAsyncCallback) {
	if _authorizationCopyRightsAsync == nil {
		panic("Security: symbol AuthorizationCopyRightsAsync not loaded")
	}
	_authorizationCopyRightsAsync(authorization, rights, environment, flags, callbackBlock)
}

var _authorizationCreate func(rights *AuthorizationRights, environment *AuthorizationEnvironment, flags AuthorizationFlags, authorization *AuthorizationRef) int32

// AuthorizationCreate creates a new authorization reference and provides an option to authorize or preauthorize rights.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationCreate(_:_:_:_:)
func AuthorizationCreate(rights *AuthorizationRights, environment *AuthorizationEnvironment, flags AuthorizationFlags, authorization *AuthorizationRef) int32 {
	if _authorizationCreate == nil {
		panic("Security: symbol AuthorizationCreate not loaded")
	}
	return _authorizationCreate(rights, environment, flags, authorization)
}

var _authorizationCreateFromExternalForm func(extForm *AuthorizationExternalForm, authorization *AuthorizationRef) int32

// AuthorizationCreateFromExternalForm internalizes the external representation of an authorization reference.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationCreateFromExternalForm(_:_:)
func AuthorizationCreateFromExternalForm(extForm *AuthorizationExternalForm, authorization *AuthorizationRef) int32 {
	if _authorizationCreateFromExternalForm == nil {
		panic("Security: symbol AuthorizationCreateFromExternalForm not loaded")
	}
	return _authorizationCreateFromExternalForm(extForm, authorization)
}

var _authorizationFree func(authorization AuthorizationRef, flags AuthorizationFlags) int32

// AuthorizationFree frees the memory associated with an authorization reference.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationFree(_:_:)
func AuthorizationFree(authorization AuthorizationRef, flags AuthorizationFlags) int32 {
	if _authorizationFree == nil {
		panic("Security: symbol AuthorizationFree not loaded")
	}
	return _authorizationFree(authorization, flags)
}

var _authorizationFreeItemSet func(set *AuthorizationItemSet) int32

// AuthorizationFreeItemSet frees the memory associated with a set of authorization items.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationFreeItemSet(_:)
func AuthorizationFreeItemSet(set *AuthorizationItemSet) int32 {
	if _authorizationFreeItemSet == nil {
		panic("Security: symbol AuthorizationFreeItemSet not loaded")
	}
	return _authorizationFreeItemSet(set)
}

var _authorizationMakeExternalForm func(authorization AuthorizationRef, extForm *AuthorizationExternalForm) int32

// AuthorizationMakeExternalForm creates an external representation of an authorization reference.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationMakeExternalForm(_:_:)
func AuthorizationMakeExternalForm(authorization AuthorizationRef, extForm *AuthorizationExternalForm) int32 {
	if _authorizationMakeExternalForm == nil {
		panic("Security: symbol AuthorizationMakeExternalForm not loaded")
	}
	return _authorizationMakeExternalForm(authorization, extForm)
}

var _authorizationRightGet func(rightName *byte, rightDefinition *corefoundation.CFDictionaryRef) int32

// AuthorizationRightGet retrieves a right definition as a dictionary.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationRightGet(_:_:)
func AuthorizationRightGet(rightName *byte, rightDefinition *corefoundation.CFDictionaryRef) int32 {
	if _authorizationRightGet == nil {
		panic("Security: symbol AuthorizationRightGet not loaded")
	}
	return _authorizationRightGet(rightName, rightDefinition)
}

var _authorizationRightRemove func(authRef AuthorizationRef, rightName *byte) int32

// AuthorizationRightRemove removes a right from the policy database.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationRightRemove(_:_:)
func AuthorizationRightRemove(authRef AuthorizationRef, rightName *byte) int32 {
	if _authorizationRightRemove == nil {
		panic("Security: symbol AuthorizationRightRemove not loaded")
	}
	return _authorizationRightRemove(authRef, rightName)
}

var _authorizationRightSet func(authRef AuthorizationRef, rightName *byte, rightDefinition corefoundation.CFTypeRef, descriptionKey corefoundation.CFStringRef, bundle corefoundation.CFBundle, localeTableName corefoundation.CFStringRef) int32

// AuthorizationRightSet creates or updates a right entry in the policy database.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationRightSet(_:_:_:_:_:_:)
func AuthorizationRightSet(authRef AuthorizationRef, rightName *byte, rightDefinition corefoundation.CFTypeRef, descriptionKey corefoundation.CFStringRef, bundle corefoundation.CFBundle, localeTableName corefoundation.CFStringRef) int32 {
	if _authorizationRightSet == nil {
		panic("Security: symbol AuthorizationRightSet not loaded")
	}
	return _authorizationRightSet(authRef, rightName, rightDefinition, descriptionKey, bundle, localeTableName)
}

var _cMSDecoderCopyAllCerts func(cmsDecoder CMSDecoderRef, certsOut *corefoundation.CFArrayRef) int32

// CMSDecoderCopyAllCerts obtains an array of all of the certificates in a message.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderCopyAllCerts(_:_:)
func CMSDecoderCopyAllCerts(cmsDecoder CMSDecoderRef, certsOut *corefoundation.CFArrayRef) int32 {
	if _cMSDecoderCopyAllCerts == nil {
		panic("Security: symbol CMSDecoderCopyAllCerts not loaded")
	}
	return _cMSDecoderCopyAllCerts(cmsDecoder, certsOut)
}

var _cMSDecoderCopyContent func(cmsDecoder CMSDecoderRef, contentOut *corefoundation.CFDataRef) int32

// CMSDecoderCopyContent obtains the message content, if any.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderCopyContent(_:_:)
func CMSDecoderCopyContent(cmsDecoder CMSDecoderRef, contentOut *corefoundation.CFDataRef) int32 {
	if _cMSDecoderCopyContent == nil {
		panic("Security: symbol CMSDecoderCopyContent not loaded")
	}
	return _cMSDecoderCopyContent(cmsDecoder, contentOut)
}

var _cMSDecoderCopyDetachedContent func(cmsDecoder CMSDecoderRef, detachedContentOut *corefoundation.CFDataRef) int32

// CMSDecoderCopyDetachedContent obtains the detached content specified with the [CMSDecoderSetDetachedContent] function.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderCopyDetachedContent(_:_:)
func CMSDecoderCopyDetachedContent(cmsDecoder CMSDecoderRef, detachedContentOut *corefoundation.CFDataRef) int32 {
	if _cMSDecoderCopyDetachedContent == nil {
		panic("Security: symbol CMSDecoderCopyDetachedContent not loaded")
	}
	return _cMSDecoderCopyDetachedContent(cmsDecoder, detachedContentOut)
}

var _cMSDecoderCopyEncapsulatedContentType func(cmsDecoder CMSDecoderRef, eContentTypeOut *corefoundation.CFDataRef) int32

// CMSDecoderCopyEncapsulatedContentType obtains the object identifier for the encapsulated data of a signed message.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderCopyEncapsulatedContentType(_:_:)
func CMSDecoderCopyEncapsulatedContentType(cmsDecoder CMSDecoderRef, eContentTypeOut *corefoundation.CFDataRef) int32 {
	if _cMSDecoderCopyEncapsulatedContentType == nil {
		panic("Security: symbol CMSDecoderCopyEncapsulatedContentType not loaded")
	}
	return _cMSDecoderCopyEncapsulatedContentType(cmsDecoder, eContentTypeOut)
}

var _cMSDecoderCopySignerCert func(cmsDecoder CMSDecoderRef, signerIndex uintptr, signerCertOut *SecCertificateRef) int32

// CMSDecoderCopySignerCert obtains the certificate of the specified signer of a CMS message.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderCopySignerCert(_:_:_:)
func CMSDecoderCopySignerCert(cmsDecoder CMSDecoderRef, signerIndex uintptr, signerCertOut *SecCertificateRef) int32 {
	if _cMSDecoderCopySignerCert == nil {
		panic("Security: symbol CMSDecoderCopySignerCert not loaded")
	}
	return _cMSDecoderCopySignerCert(cmsDecoder, signerIndex, signerCertOut)
}

var _cMSDecoderCopySignerEmailAddress func(cmsDecoder CMSDecoderRef, signerIndex uintptr, signerEmailAddressOut *corefoundation.CFStringRef) int32

// CMSDecoderCopySignerEmailAddress obtains the email address of the specified signer of a CMS message.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderCopySignerEmailAddress(_:_:_:)
func CMSDecoderCopySignerEmailAddress(cmsDecoder CMSDecoderRef, signerIndex uintptr, signerEmailAddressOut *corefoundation.CFStringRef) int32 {
	if _cMSDecoderCopySignerEmailAddress == nil {
		panic("Security: symbol CMSDecoderCopySignerEmailAddress not loaded")
	}
	return _cMSDecoderCopySignerEmailAddress(cmsDecoder, signerIndex, signerEmailAddressOut)
}

var _cMSDecoderCopySignerSigningTime func(cmsDecoder CMSDecoderRef, signerIndex uintptr, signingTime *corefoundation.CFAbsoluteTime) int32

// CMSDecoderCopySignerSigningTime obtains the signing time of a CMS message, if present.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderCopySignerSigningTime(_:_:_:)
func CMSDecoderCopySignerSigningTime(cmsDecoder CMSDecoderRef, signerIndex uintptr, signingTime *corefoundation.CFAbsoluteTime) int32 {
	if _cMSDecoderCopySignerSigningTime == nil {
		panic("Security: symbol CMSDecoderCopySignerSigningTime not loaded")
	}
	return _cMSDecoderCopySignerSigningTime(cmsDecoder, signerIndex, signingTime)
}

var _cMSDecoderCopySignerStatus func(cmsDecoder CMSDecoderRef, signerIndex uintptr, policyOrArray corefoundation.CFTypeRef, evaluateSecTrust bool, signerStatusOut *CMSSignerStatus, secTrustOut *SecTrustRef, certVerifyResultCodeOut *int32) int32

// CMSDecoderCopySignerStatus obtains the status of a CMS message’s signature.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderCopySignerStatus(_:_:_:_:_:_:_:)
func CMSDecoderCopySignerStatus(cmsDecoder CMSDecoderRef, signerIndex uintptr, policyOrArray corefoundation.CFTypeRef, evaluateSecTrust bool, signerStatusOut *CMSSignerStatus, secTrustOut *SecTrustRef, certVerifyResultCodeOut *int32) int32 {
	if _cMSDecoderCopySignerStatus == nil {
		panic("Security: symbol CMSDecoderCopySignerStatus not loaded")
	}
	return _cMSDecoderCopySignerStatus(cmsDecoder, signerIndex, policyOrArray, evaluateSecTrust, signerStatusOut, secTrustOut, certVerifyResultCodeOut)
}

var _cMSDecoderCopySignerTimestamp func(cmsDecoder CMSDecoderRef, signerIndex uintptr, timestamp *corefoundation.CFAbsoluteTime) int32

// CMSDecoderCopySignerTimestamp returns the timestamp of a signer of a CMS message, if present.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderCopySignerTimestamp(_:_:_:)
func CMSDecoderCopySignerTimestamp(cmsDecoder CMSDecoderRef, signerIndex uintptr, timestamp *corefoundation.CFAbsoluteTime) int32 {
	if _cMSDecoderCopySignerTimestamp == nil {
		panic("Security: symbol CMSDecoderCopySignerTimestamp not loaded")
	}
	return _cMSDecoderCopySignerTimestamp(cmsDecoder, signerIndex, timestamp)
}

var _cMSDecoderCopySignerTimestampCertificates func(cmsDecoder CMSDecoderRef, signerIndex uintptr, certificateRefs *corefoundation.CFArrayRef) int32

// CMSDecoderCopySignerTimestampCertificates returns an array containing the certificates from a timestamp response.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderCopySignerTimestampCertificates(_:_:_:)
func CMSDecoderCopySignerTimestampCertificates(cmsDecoder CMSDecoderRef, signerIndex uintptr, certificateRefs *corefoundation.CFArrayRef) int32 {
	if _cMSDecoderCopySignerTimestampCertificates == nil {
		panic("Security: symbol CMSDecoderCopySignerTimestampCertificates not loaded")
	}
	return _cMSDecoderCopySignerTimestampCertificates(cmsDecoder, signerIndex, certificateRefs)
}

var _cMSDecoderCopySignerTimestampWithPolicy func(cmsDecoder CMSDecoderRef, timeStampPolicy corefoundation.CFTypeRef, signerIndex uintptr, timestamp *corefoundation.CFAbsoluteTime) int32

// CMSDecoderCopySignerTimestampWithPolicy returns the timestamp of a signer of a CMS message using a given policy, if present.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderCopySignerTimestampWithPolicy(_:_:_:_:)
func CMSDecoderCopySignerTimestampWithPolicy(cmsDecoder CMSDecoderRef, timeStampPolicy corefoundation.CFTypeRef, signerIndex uintptr, timestamp *corefoundation.CFAbsoluteTime) int32 {
	if _cMSDecoderCopySignerTimestampWithPolicy == nil {
		panic("Security: symbol CMSDecoderCopySignerTimestampWithPolicy not loaded")
	}
	return _cMSDecoderCopySignerTimestampWithPolicy(cmsDecoder, timeStampPolicy, signerIndex, timestamp)
}

var _cMSDecoderCreate func(cmsDecoderOut *CMSDecoderRef) int32

// CMSDecoderCreate creates a CMSDecoder reference.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderCreate(_:)
func CMSDecoderCreate(cmsDecoderOut *CMSDecoderRef) int32 {
	if _cMSDecoderCreate == nil {
		panic("Security: symbol CMSDecoderCreate not loaded")
	}
	return _cMSDecoderCreate(cmsDecoderOut)
}

var _cMSDecoderFinalizeMessage func(cmsDecoder CMSDecoderRef) int32

// CMSDecoderFinalizeMessage indicates that there is no more data to decode.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderFinalizeMessage(_:)
func CMSDecoderFinalizeMessage(cmsDecoder CMSDecoderRef) int32 {
	if _cMSDecoderFinalizeMessage == nil {
		panic("Security: symbol CMSDecoderFinalizeMessage not loaded")
	}
	return _cMSDecoderFinalizeMessage(cmsDecoder)
}

var _cMSDecoderGetNumSigners func(cmsDecoder CMSDecoderRef, numSignersOut *uintptr) int32

// CMSDecoderGetNumSigners obtains the number of signers of a message.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderGetNumSigners(_:_:)
func CMSDecoderGetNumSigners(cmsDecoder CMSDecoderRef, numSignersOut *uintptr) int32 {
	if _cMSDecoderGetNumSigners == nil {
		panic("Security: symbol CMSDecoderGetNumSigners not loaded")
	}
	return _cMSDecoderGetNumSigners(cmsDecoder, numSignersOut)
}

var _cMSDecoderGetTypeID func() uint

// CMSDecoderGetTypeID returns the type identifier for the CMSDecoder opaque type.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderGetTypeID()
func CMSDecoderGetTypeID() uint {
	if _cMSDecoderGetTypeID == nil {
		panic("Security: symbol CMSDecoderGetTypeID not loaded")
	}
	return _cMSDecoderGetTypeID()
}

var _cMSDecoderIsContentEncrypted func(cmsDecoder CMSDecoderRef, isEncryptedOut *bool) int32

// CMSDecoderIsContentEncrypted determines whether a CMS message was encrypted.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderIsContentEncrypted(_:_:)
func CMSDecoderIsContentEncrypted(cmsDecoder CMSDecoderRef, isEncryptedOut *bool) int32 {
	if _cMSDecoderIsContentEncrypted == nil {
		panic("Security: symbol CMSDecoderIsContentEncrypted not loaded")
	}
	return _cMSDecoderIsContentEncrypted(cmsDecoder, isEncryptedOut)
}

var _cMSDecoderSetDetachedContent func(cmsDecoder CMSDecoderRef, detachedContent corefoundation.CFDataRef) int32

// CMSDecoderSetDetachedContent specifies the message’s detached content, if any.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderSetDetachedContent(_:_:)
func CMSDecoderSetDetachedContent(cmsDecoder CMSDecoderRef, detachedContent corefoundation.CFDataRef) int32 {
	if _cMSDecoderSetDetachedContent == nil {
		panic("Security: symbol CMSDecoderSetDetachedContent not loaded")
	}
	return _cMSDecoderSetDetachedContent(cmsDecoder, detachedContent)
}

var _cMSDecoderSetSearchKeychain func(cmsDecoder CMSDecoderRef, keychainOrArray corefoundation.CFTypeRef) int32

// CMSDecoderSetSearchKeychain specifies the keychains to search for intermediate certificates to be used in verifying a signed message’s signer certificates.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderSetSearchKeychain(_:_:)
func CMSDecoderSetSearchKeychain(cmsDecoder CMSDecoderRef, keychainOrArray corefoundation.CFTypeRef) int32 {
	if _cMSDecoderSetSearchKeychain == nil {
		panic("Security: symbol CMSDecoderSetSearchKeychain not loaded")
	}
	return _cMSDecoderSetSearchKeychain(cmsDecoder, keychainOrArray)
}

var _cMSDecoderUpdateMessage func(cmsDecoder CMSDecoderRef, msgBytes unsafe.Pointer, msgBytesLen uintptr) int32

// CMSDecoderUpdateMessage feeds raw bytes of the message to be decoded into the decoder.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoderUpdateMessage(_:_:_:)
func CMSDecoderUpdateMessage(cmsDecoder CMSDecoderRef, msgBytes unsafe.Pointer, msgBytesLen uintptr) int32 {
	if _cMSDecoderUpdateMessage == nil {
		panic("Security: symbol CMSDecoderUpdateMessage not loaded")
	}
	return _cMSDecoderUpdateMessage(cmsDecoder, msgBytes, msgBytesLen)
}

var _cMSEncode func(signers corefoundation.CFTypeRef, recipients corefoundation.CFTypeRef, eContentType unsafe.Pointer, detachedContent bool, signedAttributes CMSSignedAttributes, content unsafe.Pointer, contentLen uintptr, encodedContentOut *corefoundation.CFDataRef) int32

// CMSEncode encodes a message and obtains the result in one high-level function call.
//
// See: https://developer.apple.com/documentation/Security/CMSEncode
func CMSEncode(signers corefoundation.CFTypeRef, recipients corefoundation.CFTypeRef, eContentType unsafe.Pointer, detachedContent bool, signedAttributes CMSSignedAttributes, content unsafe.Pointer, contentLen uintptr, encodedContentOut *corefoundation.CFDataRef) int32 {
	if _cMSEncode == nil {
		panic("Security: symbol CMSEncode not loaded")
	}
	return _cMSEncode(signers, recipients, eContentType, detachedContent, signedAttributes, content, contentLen, encodedContentOut)
}

var _cMSEncodeContent func(signers corefoundation.CFTypeRef, recipients corefoundation.CFTypeRef, eContentTypeOID corefoundation.CFTypeRef, detachedContent bool, signedAttributes CMSSignedAttributes, content unsafe.Pointer, contentLen uintptr, encodedContentOut *corefoundation.CFDataRef) int32

// CMSEncodeContent encodes a message and obtains the result in one high-level function call.
//
// See: https://developer.apple.com/documentation/Security/CMSEncodeContent(_:_:_:_:_:_:_:_:)
func CMSEncodeContent(signers corefoundation.CFTypeRef, recipients corefoundation.CFTypeRef, eContentTypeOID corefoundation.CFTypeRef, detachedContent bool, signedAttributes CMSSignedAttributes, content unsafe.Pointer, contentLen uintptr, encodedContentOut *corefoundation.CFDataRef) int32 {
	if _cMSEncodeContent == nil {
		panic("Security: symbol CMSEncodeContent not loaded")
	}
	return _cMSEncodeContent(signers, recipients, eContentTypeOID, detachedContent, signedAttributes, content, contentLen, encodedContentOut)
}

var _cMSEncoderAddRecipients func(cmsEncoder CMSEncoderRef, recipientOrArray corefoundation.CFTypeRef) int32

// CMSEncoderAddRecipients specifies a message is to be encrypted and specifies the recipients of the message.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderAddRecipients(_:_:)
func CMSEncoderAddRecipients(cmsEncoder CMSEncoderRef, recipientOrArray corefoundation.CFTypeRef) int32 {
	if _cMSEncoderAddRecipients == nil {
		panic("Security: symbol CMSEncoderAddRecipients not loaded")
	}
	return _cMSEncoderAddRecipients(cmsEncoder, recipientOrArray)
}

var _cMSEncoderAddSignedAttributes func(cmsEncoder CMSEncoderRef, signedAttributes CMSSignedAttributes) int32

// CMSEncoderAddSignedAttributes specifies attributes for a signed message.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderAddSignedAttributes(_:_:)
func CMSEncoderAddSignedAttributes(cmsEncoder CMSEncoderRef, signedAttributes CMSSignedAttributes) int32 {
	if _cMSEncoderAddSignedAttributes == nil {
		panic("Security: symbol CMSEncoderAddSignedAttributes not loaded")
	}
	return _cMSEncoderAddSignedAttributes(cmsEncoder, signedAttributes)
}

var _cMSEncoderAddSigners func(cmsEncoder CMSEncoderRef, signerOrArray corefoundation.CFTypeRef) int32

// CMSEncoderAddSigners specifies signers of the message.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderAddSigners(_:_:)
func CMSEncoderAddSigners(cmsEncoder CMSEncoderRef, signerOrArray corefoundation.CFTypeRef) int32 {
	if _cMSEncoderAddSigners == nil {
		panic("Security: symbol CMSEncoderAddSigners not loaded")
	}
	return _cMSEncoderAddSigners(cmsEncoder, signerOrArray)
}

var _cMSEncoderAddSupportingCerts func(cmsEncoder CMSEncoderRef, certOrArray corefoundation.CFTypeRef) int32

// CMSEncoderAddSupportingCerts adds certificates to a message.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderAddSupportingCerts(_:_:)
func CMSEncoderAddSupportingCerts(cmsEncoder CMSEncoderRef, certOrArray corefoundation.CFTypeRef) int32 {
	if _cMSEncoderAddSupportingCerts == nil {
		panic("Security: symbol CMSEncoderAddSupportingCerts not loaded")
	}
	return _cMSEncoderAddSupportingCerts(cmsEncoder, certOrArray)
}

var _cMSEncoderCopyEncapsulatedContentType func(cmsEncoder CMSEncoderRef, eContentTypeOut *corefoundation.CFDataRef) int32

// CMSEncoderCopyEncapsulatedContentType obtains the object identifier for the encapsulated data of a signed message.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderCopyEncapsulatedContentType(_:_:)
func CMSEncoderCopyEncapsulatedContentType(cmsEncoder CMSEncoderRef, eContentTypeOut *corefoundation.CFDataRef) int32 {
	if _cMSEncoderCopyEncapsulatedContentType == nil {
		panic("Security: symbol CMSEncoderCopyEncapsulatedContentType not loaded")
	}
	return _cMSEncoderCopyEncapsulatedContentType(cmsEncoder, eContentTypeOut)
}

var _cMSEncoderCopyEncodedContent func(cmsEncoder CMSEncoderRef, encodedContentOut *corefoundation.CFDataRef) int32

// CMSEncoderCopyEncodedContent finishes encoding the message and obtains the encoded result.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderCopyEncodedContent(_:_:)
func CMSEncoderCopyEncodedContent(cmsEncoder CMSEncoderRef, encodedContentOut *corefoundation.CFDataRef) int32 {
	if _cMSEncoderCopyEncodedContent == nil {
		panic("Security: symbol CMSEncoderCopyEncodedContent not loaded")
	}
	return _cMSEncoderCopyEncodedContent(cmsEncoder, encodedContentOut)
}

var _cMSEncoderCopyRecipients func(cmsEncoder CMSEncoderRef, recipientsOut *corefoundation.CFArrayRef) int32

// CMSEncoderCopyRecipients obtains the array of recipients specified with the [CMSEncoderAddRecipients] function.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderCopyRecipients(_:_:)
func CMSEncoderCopyRecipients(cmsEncoder CMSEncoderRef, recipientsOut *corefoundation.CFArrayRef) int32 {
	if _cMSEncoderCopyRecipients == nil {
		panic("Security: symbol CMSEncoderCopyRecipients not loaded")
	}
	return _cMSEncoderCopyRecipients(cmsEncoder, recipientsOut)
}

var _cMSEncoderCopySignerTimestamp func(cmsEncoder CMSEncoderRef, signerIndex uintptr, timestamp *corefoundation.CFAbsoluteTime) int32

// CMSEncoderCopySignerTimestamp returns the timestamp of a signer of a CMS message, if present.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderCopySignerTimestamp(_:_:_:)
func CMSEncoderCopySignerTimestamp(cmsEncoder CMSEncoderRef, signerIndex uintptr, timestamp *corefoundation.CFAbsoluteTime) int32 {
	if _cMSEncoderCopySignerTimestamp == nil {
		panic("Security: symbol CMSEncoderCopySignerTimestamp not loaded")
	}
	return _cMSEncoderCopySignerTimestamp(cmsEncoder, signerIndex, timestamp)
}

var _cMSEncoderCopySignerTimestampWithPolicy func(cmsEncoder CMSEncoderRef, timeStampPolicy corefoundation.CFTypeRef, signerIndex uintptr, timestamp *corefoundation.CFAbsoluteTime) int32

// CMSEncoderCopySignerTimestampWithPolicy returns the timestamp of a signer of a CMS message using a particular policy, if present.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderCopySignerTimestampWithPolicy(_:_:_:_:)
func CMSEncoderCopySignerTimestampWithPolicy(cmsEncoder CMSEncoderRef, timeStampPolicy corefoundation.CFTypeRef, signerIndex uintptr, timestamp *corefoundation.CFAbsoluteTime) int32 {
	if _cMSEncoderCopySignerTimestampWithPolicy == nil {
		panic("Security: symbol CMSEncoderCopySignerTimestampWithPolicy not loaded")
	}
	return _cMSEncoderCopySignerTimestampWithPolicy(cmsEncoder, timeStampPolicy, signerIndex, timestamp)
}

var _cMSEncoderCopySigners func(cmsEncoder CMSEncoderRef, signersOut *corefoundation.CFArrayRef) int32

// CMSEncoderCopySigners obtains the array of signers specified with the [CMSEncoderAddSigners] function.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderCopySigners(_:_:)
func CMSEncoderCopySigners(cmsEncoder CMSEncoderRef, signersOut *corefoundation.CFArrayRef) int32 {
	if _cMSEncoderCopySigners == nil {
		panic("Security: symbol CMSEncoderCopySigners not loaded")
	}
	return _cMSEncoderCopySigners(cmsEncoder, signersOut)
}

var _cMSEncoderCopySupportingCerts func(cmsEncoder CMSEncoderRef, certsOut *corefoundation.CFArrayRef) int32

// CMSEncoderCopySupportingCerts obtains the certificates added to a message with [CMSEncoderAddSupportingCerts].
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderCopySupportingCerts(_:_:)
func CMSEncoderCopySupportingCerts(cmsEncoder CMSEncoderRef, certsOut *corefoundation.CFArrayRef) int32 {
	if _cMSEncoderCopySupportingCerts == nil {
		panic("Security: symbol CMSEncoderCopySupportingCerts not loaded")
	}
	return _cMSEncoderCopySupportingCerts(cmsEncoder, certsOut)
}

var _cMSEncoderCreate func(cmsEncoderOut *CMSEncoderRef) int32

// CMSEncoderCreate creates a CMSEncoder reference.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderCreate(_:)
func CMSEncoderCreate(cmsEncoderOut *CMSEncoderRef) int32 {
	if _cMSEncoderCreate == nil {
		panic("Security: symbol CMSEncoderCreate not loaded")
	}
	return _cMSEncoderCreate(cmsEncoderOut)
}

var _cMSEncoderGetCertificateChainMode func(cmsEncoder CMSEncoderRef, chainModeOut *CMSCertificateChainMode) int32

// CMSEncoderGetCertificateChainMode obtains a constant that indicates which certificates are to be included in a signed CMS message.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderGetCertificateChainMode(_:_:)
func CMSEncoderGetCertificateChainMode(cmsEncoder CMSEncoderRef, chainModeOut *CMSCertificateChainMode) int32 {
	if _cMSEncoderGetCertificateChainMode == nil {
		panic("Security: symbol CMSEncoderGetCertificateChainMode not loaded")
	}
	return _cMSEncoderGetCertificateChainMode(cmsEncoder, chainModeOut)
}

var _cMSEncoderGetHasDetachedContent func(cmsEncoder CMSEncoderRef, detachedContentOut *bool) int32

// CMSEncoderGetHasDetachedContent indicates whether the message is to have detached content.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderGetHasDetachedContent(_:_:)
func CMSEncoderGetHasDetachedContent(cmsEncoder CMSEncoderRef, detachedContentOut *bool) int32 {
	if _cMSEncoderGetHasDetachedContent == nil {
		panic("Security: symbol CMSEncoderGetHasDetachedContent not loaded")
	}
	return _cMSEncoderGetHasDetachedContent(cmsEncoder, detachedContentOut)
}

var _cMSEncoderGetTypeID func() uint

// CMSEncoderGetTypeID returns the type identifier for the CMSEncoder opaque type.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderGetTypeID()
func CMSEncoderGetTypeID() uint {
	if _cMSEncoderGetTypeID == nil {
		panic("Security: symbol CMSEncoderGetTypeID not loaded")
	}
	return _cMSEncoderGetTypeID()
}

var _cMSEncoderSetCertificateChainMode func(cmsEncoder CMSEncoderRef, chainMode CMSCertificateChainMode) int32

// CMSEncoderSetCertificateChainMode specifies which certificates to include in a signed CMS message.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderSetCertificateChainMode(_:_:)
func CMSEncoderSetCertificateChainMode(cmsEncoder CMSEncoderRef, chainMode CMSCertificateChainMode) int32 {
	if _cMSEncoderSetCertificateChainMode == nil {
		panic("Security: symbol CMSEncoderSetCertificateChainMode not loaded")
	}
	return _cMSEncoderSetCertificateChainMode(cmsEncoder, chainMode)
}

var _cMSEncoderSetEncapsulatedContentType func(cmsEncoder CMSEncoderRef, eContentType unsafe.Pointer) int32

// CMSEncoderSetEncapsulatedContentType specifies an object identifier for the encapsulated data of a signed message.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderSetEncapsulatedContentType
func CMSEncoderSetEncapsulatedContentType(cmsEncoder CMSEncoderRef, eContentType unsafe.Pointer) int32 {
	if _cMSEncoderSetEncapsulatedContentType == nil {
		panic("Security: symbol CMSEncoderSetEncapsulatedContentType not loaded")
	}
	return _cMSEncoderSetEncapsulatedContentType(cmsEncoder, eContentType)
}

var _cMSEncoderSetEncapsulatedContentTypeOID func(cmsEncoder CMSEncoderRef, eContentTypeOID corefoundation.CFTypeRef) int32

// CMSEncoderSetEncapsulatedContentTypeOID specifies an object identifier for the encapsulated data of a signed message.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderSetEncapsulatedContentTypeOID(_:_:)
func CMSEncoderSetEncapsulatedContentTypeOID(cmsEncoder CMSEncoderRef, eContentTypeOID corefoundation.CFTypeRef) int32 {
	if _cMSEncoderSetEncapsulatedContentTypeOID == nil {
		panic("Security: symbol CMSEncoderSetEncapsulatedContentTypeOID not loaded")
	}
	return _cMSEncoderSetEncapsulatedContentTypeOID(cmsEncoder, eContentTypeOID)
}

var _cMSEncoderSetHasDetachedContent func(cmsEncoder CMSEncoderRef, detachedContent bool) int32

// CMSEncoderSetHasDetachedContent specifies whether the signed data is to be separate from the message.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderSetHasDetachedContent(_:_:)
func CMSEncoderSetHasDetachedContent(cmsEncoder CMSEncoderRef, detachedContent bool) int32 {
	if _cMSEncoderSetHasDetachedContent == nil {
		panic("Security: symbol CMSEncoderSetHasDetachedContent not loaded")
	}
	return _cMSEncoderSetHasDetachedContent(cmsEncoder, detachedContent)
}

var _cMSEncoderSetSignerAlgorithm func(cmsEncoder CMSEncoderRef, digestAlgorithm corefoundation.CFStringRef) int32

// CMSEncoderSetSignerAlgorithm sets the digest algorithm to use for the signer.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderSetSignerAlgorithm(_:_:)
func CMSEncoderSetSignerAlgorithm(cmsEncoder CMSEncoderRef, digestAlgorithm corefoundation.CFStringRef) int32 {
	if _cMSEncoderSetSignerAlgorithm == nil {
		panic("Security: symbol CMSEncoderSetSignerAlgorithm not loaded")
	}
	return _cMSEncoderSetSignerAlgorithm(cmsEncoder, digestAlgorithm)
}

var _cMSEncoderUpdateContent func(cmsEncoder CMSEncoderRef, content unsafe.Pointer, contentLen uintptr) int32

// CMSEncoderUpdateContent feeds content bytes into the encoder.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoderUpdateContent(_:_:_:)
func CMSEncoderUpdateContent(cmsEncoder CMSEncoderRef, content unsafe.Pointer, contentLen uintptr) int32 {
	if _cMSEncoderUpdateContent == nil {
		panic("Security: symbol CMSEncoderUpdateContent not loaded")
	}
	return _cMSEncoderUpdateContent(cmsEncoder, content, contentLen)
}

var _cSSM_AC_AuthCompute func(ACHandle CSSM_AC_HANDLE, BaseAuthorizations unsafe.Pointer, Credentials unsafe.Pointer, NumberOfRequestors uint32, Requestors unsafe.Pointer, RequestedAuthorizationPeriod unsafe.Pointer, RequestedAuthorization unsafe.Pointer, AuthorizationResult unsafe.Pointer) CSSM_RETURN

// CSSM_AC_AuthCompute.
//
// See: https://developer.apple.com/documentation/Security/CSSM_AC_AuthCompute
func CSSM_AC_AuthCompute(ACHandle CSSM_AC_HANDLE, BaseAuthorizations unsafe.Pointer, Credentials unsafe.Pointer, NumberOfRequestors uint32, Requestors unsafe.Pointer, RequestedAuthorizationPeriod unsafe.Pointer, RequestedAuthorization unsafe.Pointer, AuthorizationResult unsafe.Pointer) CSSM_RETURN {
	if _cSSM_AC_AuthCompute == nil {
		panic("Security: symbol CSSM_AC_AuthCompute not loaded")
	}
	return _cSSM_AC_AuthCompute(ACHandle, BaseAuthorizations, Credentials, NumberOfRequestors, Requestors, RequestedAuthorizationPeriod, RequestedAuthorization, AuthorizationResult)
}

var _cSSM_AC_PassThrough func(ACHandle CSSM_AC_HANDLE, TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, DBList unsafe.Pointer, PassThroughId uint32, InputParams unsafe.Pointer, OutputParams unsafe.Pointer) CSSM_RETURN

// CSSM_AC_PassThrough.
//
// See: https://developer.apple.com/documentation/Security/CSSM_AC_PassThrough
func CSSM_AC_PassThrough(ACHandle CSSM_AC_HANDLE, TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, DBList unsafe.Pointer, PassThroughId uint32, InputParams unsafe.Pointer, OutputParams unsafe.Pointer) CSSM_RETURN {
	if _cSSM_AC_PassThrough == nil {
		panic("Security: symbol CSSM_AC_PassThrough not loaded")
	}
	return _cSSM_AC_PassThrough(ACHandle, TPHandle, CLHandle, CCHandle, DBList, PassThroughId, InputParams, OutputParams)
}

var _cSSM_CL_CertAbortCache func(CLHandle CSSM_CL_HANDLE, CertHandle CSSM_HANDLE) CSSM_RETURN

// CSSM_CL_CertAbortCache.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertAbortCache
func CSSM_CL_CertAbortCache(CLHandle CSSM_CL_HANDLE, CertHandle CSSM_HANDLE) CSSM_RETURN {
	if _cSSM_CL_CertAbortCache == nil {
		panic("Security: symbol CSSM_CL_CertAbortCache not loaded")
	}
	return _cSSM_CL_CertAbortCache(CLHandle, CertHandle)
}

var _cSSM_CL_CertAbortQuery func(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE) CSSM_RETURN

// CSSM_CL_CertAbortQuery.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertAbortQuery
func CSSM_CL_CertAbortQuery(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE) CSSM_RETURN {
	if _cSSM_CL_CertAbortQuery == nil {
		panic("Security: symbol CSSM_CL_CertAbortQuery not loaded")
	}
	return _cSSM_CL_CertAbortQuery(CLHandle, ResultsHandle)
}

var _cSSM_CL_CertCache func(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, CertHandle CSSM_HANDLE_PTR) CSSM_RETURN

// CSSM_CL_CertCache.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertCache
func CSSM_CL_CertCache(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, CertHandle CSSM_HANDLE_PTR) CSSM_RETURN {
	if _cSSM_CL_CertCache == nil {
		panic("Security: symbol CSSM_CL_CertCache not loaded")
	}
	return _cSSM_CL_CertCache(CLHandle, Cert, CertHandle)
}

var _cSSM_CL_CertCreateTemplate func(CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CertFields unsafe.Pointer, CertTemplate unsafe.Pointer) CSSM_RETURN

// CSSM_CL_CertCreateTemplate.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertCreateTemplate
func CSSM_CL_CertCreateTemplate(CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CertFields unsafe.Pointer, CertTemplate unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CL_CertCreateTemplate == nil {
		panic("Security: symbol CSSM_CL_CertCreateTemplate not loaded")
	}
	return _cSSM_CL_CertCreateTemplate(CLHandle, NumberOfFields, CertFields, CertTemplate)
}

var _cSSM_CL_CertDescribeFormat func(CLHandle CSSM_CL_HANDLE, NumberOfFields *uint32, OidList unsafe.Pointer) CSSM_RETURN

// CSSM_CL_CertDescribeFormat.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertDescribeFormat
func CSSM_CL_CertDescribeFormat(CLHandle CSSM_CL_HANDLE, NumberOfFields *uint32, OidList unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CL_CertDescribeFormat == nil {
		panic("Security: symbol CSSM_CL_CertDescribeFormat not loaded")
	}
	return _cSSM_CL_CertDescribeFormat(CLHandle, NumberOfFields, OidList)
}

var _cSSM_CL_CertGetAllFields func(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, NumberOfFields *uint32, CertFields unsafe.Pointer) CSSM_RETURN

// CSSM_CL_CertGetAllFields.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertGetAllFields
func CSSM_CL_CertGetAllFields(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, NumberOfFields *uint32, CertFields unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CL_CertGetAllFields == nil {
		panic("Security: symbol CSSM_CL_CertGetAllFields not loaded")
	}
	return _cSSM_CL_CertGetAllFields(CLHandle, Cert, NumberOfFields, CertFields)
}

var _cSSM_CL_CertGetAllTemplateFields func(CLHandle CSSM_CL_HANDLE, CertTemplate unsafe.Pointer, NumberOfFields *uint32, CertFields unsafe.Pointer) CSSM_RETURN

// CSSM_CL_CertGetAllTemplateFields.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertGetAllTemplateFields
func CSSM_CL_CertGetAllTemplateFields(CLHandle CSSM_CL_HANDLE, CertTemplate unsafe.Pointer, NumberOfFields *uint32, CertFields unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CL_CertGetAllTemplateFields == nil {
		panic("Security: symbol CSSM_CL_CertGetAllTemplateFields not loaded")
	}
	return _cSSM_CL_CertGetAllTemplateFields(CLHandle, CertTemplate, NumberOfFields, CertFields)
}

var _cSSM_CL_CertGetFirstCachedFieldValue func(CLHandle CSSM_CL_HANDLE, CertHandle CSSM_HANDLE, CertField unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, NumberOfMatchedFields *uint32, Value unsafe.Pointer) CSSM_RETURN

// CSSM_CL_CertGetFirstCachedFieldValue.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertGetFirstCachedFieldValue
func CSSM_CL_CertGetFirstCachedFieldValue(CLHandle CSSM_CL_HANDLE, CertHandle CSSM_HANDLE, CertField unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, NumberOfMatchedFields *uint32, Value unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CL_CertGetFirstCachedFieldValue == nil {
		panic("Security: symbol CSSM_CL_CertGetFirstCachedFieldValue not loaded")
	}
	return _cSSM_CL_CertGetFirstCachedFieldValue(CLHandle, CertHandle, CertField, ResultsHandle, NumberOfMatchedFields, Value)
}

var _cSSM_CL_CertGetFirstFieldValue func(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, CertField unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, NumberOfMatchedFields *uint32, Value unsafe.Pointer) CSSM_RETURN

// CSSM_CL_CertGetFirstFieldValue.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertGetFirstFieldValue
func CSSM_CL_CertGetFirstFieldValue(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, CertField unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, NumberOfMatchedFields *uint32, Value unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CL_CertGetFirstFieldValue == nil {
		panic("Security: symbol CSSM_CL_CertGetFirstFieldValue not loaded")
	}
	return _cSSM_CL_CertGetFirstFieldValue(CLHandle, Cert, CertField, ResultsHandle, NumberOfMatchedFields, Value)
}

var _cSSM_CL_CertGetKeyInfo func(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, Key unsafe.Pointer) CSSM_RETURN

// CSSM_CL_CertGetKeyInfo.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertGetKeyInfo
func CSSM_CL_CertGetKeyInfo(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, Key unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CL_CertGetKeyInfo == nil {
		panic("Security: symbol CSSM_CL_CertGetKeyInfo not loaded")
	}
	return _cSSM_CL_CertGetKeyInfo(CLHandle, Cert, Key)
}

var _cSSM_CL_CertGetNextCachedFieldValue func(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE, Value unsafe.Pointer) CSSM_RETURN

// CSSM_CL_CertGetNextCachedFieldValue.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertGetNextCachedFieldValue
func CSSM_CL_CertGetNextCachedFieldValue(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE, Value unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CL_CertGetNextCachedFieldValue == nil {
		panic("Security: symbol CSSM_CL_CertGetNextCachedFieldValue not loaded")
	}
	return _cSSM_CL_CertGetNextCachedFieldValue(CLHandle, ResultsHandle, Value)
}

var _cSSM_CL_CertGetNextFieldValue func(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE, Value unsafe.Pointer) CSSM_RETURN

// CSSM_CL_CertGetNextFieldValue.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertGetNextFieldValue
func CSSM_CL_CertGetNextFieldValue(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE, Value unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CL_CertGetNextFieldValue == nil {
		panic("Security: symbol CSSM_CL_CertGetNextFieldValue not loaded")
	}
	return _cSSM_CL_CertGetNextFieldValue(CLHandle, ResultsHandle, Value)
}

var _cSSM_CL_CertGroupFromVerifiedBundle func(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertBundle unsafe.Pointer, SignerCert unsafe.Pointer, CertGroup unsafe.Pointer) CSSM_RETURN

// CSSM_CL_CertGroupFromVerifiedBundle.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertGroupFromVerifiedBundle
func CSSM_CL_CertGroupFromVerifiedBundle(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertBundle unsafe.Pointer, SignerCert unsafe.Pointer, CertGroup unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CL_CertGroupFromVerifiedBundle == nil {
		panic("Security: symbol CSSM_CL_CertGroupFromVerifiedBundle not loaded")
	}
	return _cSSM_CL_CertGroupFromVerifiedBundle(CLHandle, CCHandle, CertBundle, SignerCert, CertGroup)
}

var _cSSM_CL_CertGroupToSignedBundle func(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertGroupToBundle unsafe.Pointer, BundleInfo unsafe.Pointer, SignedBundle unsafe.Pointer) CSSM_RETURN

// CSSM_CL_CertGroupToSignedBundle.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertGroupToSignedBundle
func CSSM_CL_CertGroupToSignedBundle(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertGroupToBundle unsafe.Pointer, BundleInfo unsafe.Pointer, SignedBundle unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CL_CertGroupToSignedBundle == nil {
		panic("Security: symbol CSSM_CL_CertGroupToSignedBundle not loaded")
	}
	return _cSSM_CL_CertGroupToSignedBundle(CLHandle, CCHandle, CertGroupToBundle, BundleInfo, SignedBundle)
}

var _cSSM_CL_CertSign func(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertTemplate unsafe.Pointer, SignScope unsafe.Pointer, ScopeSize uint32, SignedCert unsafe.Pointer) CSSM_RETURN

// CSSM_CL_CertSign.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertSign
func CSSM_CL_CertSign(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertTemplate unsafe.Pointer, SignScope unsafe.Pointer, ScopeSize uint32, SignedCert unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CL_CertSign == nil {
		panic("Security: symbol CSSM_CL_CertSign not loaded")
	}
	return _cSSM_CL_CertSign(CLHandle, CCHandle, CertTemplate, SignScope, ScopeSize, SignedCert)
}

var _cSSM_CL_CertVerify func(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertToBeVerified unsafe.Pointer, SignerCert unsafe.Pointer, VerifyScope unsafe.Pointer, ScopeSize uint32) CSSM_RETURN

// CSSM_CL_CertVerify.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertVerify
func CSSM_CL_CertVerify(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertToBeVerified unsafe.Pointer, SignerCert unsafe.Pointer, VerifyScope unsafe.Pointer, ScopeSize uint32) CSSM_RETURN {
	if _cSSM_CL_CertVerify == nil {
		panic("Security: symbol CSSM_CL_CertVerify not loaded")
	}
	return _cSSM_CL_CertVerify(CLHandle, CCHandle, CertToBeVerified, SignerCert, VerifyScope, ScopeSize)
}

var _cSSM_CL_CertVerifyWithKey func(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertToBeVerified unsafe.Pointer) CSSM_RETURN

// CSSM_CL_CertVerifyWithKey.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CertVerifyWithKey
func CSSM_CL_CertVerifyWithKey(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertToBeVerified unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CL_CertVerifyWithKey == nil {
		panic("Security: symbol CSSM_CL_CertVerifyWithKey not loaded")
	}
	return _cSSM_CL_CertVerifyWithKey(CLHandle, CCHandle, CertToBeVerified)
}

var _cSSM_CL_CrlAbortCache func(CLHandle CSSM_CL_HANDLE, CrlHandle CSSM_HANDLE) CSSM_RETURN

// CSSM_CL_CrlAbortCache.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlAbortCache
func CSSM_CL_CrlAbortCache(CLHandle CSSM_CL_HANDLE, CrlHandle CSSM_HANDLE) CSSM_RETURN {
	if _cSSM_CL_CrlAbortCache == nil {
		panic("Security: symbol CSSM_CL_CrlAbortCache not loaded")
	}
	return _cSSM_CL_CrlAbortCache(CLHandle, CrlHandle)
}

var _cSSM_CL_CrlAbortQuery func(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE) CSSM_RETURN

// CSSM_CL_CrlAbortQuery.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlAbortQuery
func CSSM_CL_CrlAbortQuery(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE) CSSM_RETURN {
	if _cSSM_CL_CrlAbortQuery == nil {
		panic("Security: symbol CSSM_CL_CrlAbortQuery not loaded")
	}
	return _cSSM_CL_CrlAbortQuery(CLHandle, ResultsHandle)
}

var _cSSM_CL_CrlAddCert func(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, Cert unsafe.Pointer, NumberOfFields uint32, CrlEntryFields unsafe.Pointer, OldCrl unsafe.Pointer, NewCrl unsafe.Pointer) CSSM_RETURN

// CSSM_CL_CrlAddCert.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlAddCert
func CSSM_CL_CrlAddCert(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, Cert unsafe.Pointer, NumberOfFields uint32, CrlEntryFields unsafe.Pointer, OldCrl unsafe.Pointer, NewCrl unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CL_CrlAddCert == nil {
		panic("Security: symbol CSSM_CL_CrlAddCert not loaded")
	}
	return _cSSM_CL_CrlAddCert(CLHandle, CCHandle, Cert, NumberOfFields, CrlEntryFields, OldCrl, NewCrl)
}

var _cSSM_CL_CrlCache func(CLHandle CSSM_CL_HANDLE, Crl unsafe.Pointer, CrlHandle CSSM_HANDLE_PTR) CSSM_RETURN

// CSSM_CL_CrlCache.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlCache
func CSSM_CL_CrlCache(CLHandle CSSM_CL_HANDLE, Crl unsafe.Pointer, CrlHandle CSSM_HANDLE_PTR) CSSM_RETURN {
	if _cSSM_CL_CrlCache == nil {
		panic("Security: symbol CSSM_CL_CrlCache not loaded")
	}
	return _cSSM_CL_CrlCache(CLHandle, Crl, CrlHandle)
}

var _cSSM_CL_CrlCreateTemplate func(CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CrlTemplate unsafe.Pointer, NewCrl unsafe.Pointer) CSSM_RETURN

// CSSM_CL_CrlCreateTemplate.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlCreateTemplate
func CSSM_CL_CrlCreateTemplate(CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CrlTemplate unsafe.Pointer, NewCrl unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CL_CrlCreateTemplate == nil {
		panic("Security: symbol CSSM_CL_CrlCreateTemplate not loaded")
	}
	return _cSSM_CL_CrlCreateTemplate(CLHandle, NumberOfFields, CrlTemplate, NewCrl)
}

var _cSSM_CL_CrlDescribeFormat func(CLHandle CSSM_CL_HANDLE, NumberOfFields *uint32, OidList unsafe.Pointer) CSSM_RETURN

// CSSM_CL_CrlDescribeFormat.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlDescribeFormat
func CSSM_CL_CrlDescribeFormat(CLHandle CSSM_CL_HANDLE, NumberOfFields *uint32, OidList unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CL_CrlDescribeFormat == nil {
		panic("Security: symbol CSSM_CL_CrlDescribeFormat not loaded")
	}
	return _cSSM_CL_CrlDescribeFormat(CLHandle, NumberOfFields, OidList)
}

var _cSSM_CL_CrlGetAllCachedRecordFields func(CLHandle CSSM_CL_HANDLE, CrlHandle CSSM_HANDLE, CrlRecordIndex unsafe.Pointer, NumberOfFields *uint32, CrlFields unsafe.Pointer) CSSM_RETURN

// CSSM_CL_CrlGetAllCachedRecordFields.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlGetAllCachedRecordFields
func CSSM_CL_CrlGetAllCachedRecordFields(CLHandle CSSM_CL_HANDLE, CrlHandle CSSM_HANDLE, CrlRecordIndex unsafe.Pointer, NumberOfFields *uint32, CrlFields unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CL_CrlGetAllCachedRecordFields == nil {
		panic("Security: symbol CSSM_CL_CrlGetAllCachedRecordFields not loaded")
	}
	return _cSSM_CL_CrlGetAllCachedRecordFields(CLHandle, CrlHandle, CrlRecordIndex, NumberOfFields, CrlFields)
}

var _cSSM_CL_CrlGetAllFields func(CLHandle CSSM_CL_HANDLE, Crl unsafe.Pointer, NumberOfCrlFields *uint32, CrlFields unsafe.Pointer) CSSM_RETURN

// CSSM_CL_CrlGetAllFields.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlGetAllFields
func CSSM_CL_CrlGetAllFields(CLHandle CSSM_CL_HANDLE, Crl unsafe.Pointer, NumberOfCrlFields *uint32, CrlFields unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CL_CrlGetAllFields == nil {
		panic("Security: symbol CSSM_CL_CrlGetAllFields not loaded")
	}
	return _cSSM_CL_CrlGetAllFields(CLHandle, Crl, NumberOfCrlFields, CrlFields)
}

var _cSSM_CL_CrlGetFirstCachedFieldValue func(CLHandle CSSM_CL_HANDLE, CrlHandle CSSM_HANDLE, CrlRecordIndex unsafe.Pointer, CrlField unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, NumberOfMatchedFields *uint32, Value unsafe.Pointer) CSSM_RETURN

// CSSM_CL_CrlGetFirstCachedFieldValue.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlGetFirstCachedFieldValue
func CSSM_CL_CrlGetFirstCachedFieldValue(CLHandle CSSM_CL_HANDLE, CrlHandle CSSM_HANDLE, CrlRecordIndex unsafe.Pointer, CrlField unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, NumberOfMatchedFields *uint32, Value unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CL_CrlGetFirstCachedFieldValue == nil {
		panic("Security: symbol CSSM_CL_CrlGetFirstCachedFieldValue not loaded")
	}
	return _cSSM_CL_CrlGetFirstCachedFieldValue(CLHandle, CrlHandle, CrlRecordIndex, CrlField, ResultsHandle, NumberOfMatchedFields, Value)
}

var _cSSM_CL_CrlGetFirstFieldValue func(CLHandle CSSM_CL_HANDLE, Crl unsafe.Pointer, CrlField unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, NumberOfMatchedFields *uint32, Value unsafe.Pointer) CSSM_RETURN

// CSSM_CL_CrlGetFirstFieldValue.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlGetFirstFieldValue
func CSSM_CL_CrlGetFirstFieldValue(CLHandle CSSM_CL_HANDLE, Crl unsafe.Pointer, CrlField unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, NumberOfMatchedFields *uint32, Value unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CL_CrlGetFirstFieldValue == nil {
		panic("Security: symbol CSSM_CL_CrlGetFirstFieldValue not loaded")
	}
	return _cSSM_CL_CrlGetFirstFieldValue(CLHandle, Crl, CrlField, ResultsHandle, NumberOfMatchedFields, Value)
}

var _cSSM_CL_CrlGetNextCachedFieldValue func(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE, Value unsafe.Pointer) CSSM_RETURN

// CSSM_CL_CrlGetNextCachedFieldValue.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlGetNextCachedFieldValue
func CSSM_CL_CrlGetNextCachedFieldValue(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE, Value unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CL_CrlGetNextCachedFieldValue == nil {
		panic("Security: symbol CSSM_CL_CrlGetNextCachedFieldValue not loaded")
	}
	return _cSSM_CL_CrlGetNextCachedFieldValue(CLHandle, ResultsHandle, Value)
}

var _cSSM_CL_CrlGetNextFieldValue func(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE, Value unsafe.Pointer) CSSM_RETURN

// CSSM_CL_CrlGetNextFieldValue.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlGetNextFieldValue
func CSSM_CL_CrlGetNextFieldValue(CLHandle CSSM_CL_HANDLE, ResultsHandle CSSM_HANDLE, Value unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CL_CrlGetNextFieldValue == nil {
		panic("Security: symbol CSSM_CL_CrlGetNextFieldValue not loaded")
	}
	return _cSSM_CL_CrlGetNextFieldValue(CLHandle, ResultsHandle, Value)
}

var _cSSM_CL_CrlRemoveCert func(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, OldCrl unsafe.Pointer, NewCrl unsafe.Pointer) CSSM_RETURN

// CSSM_CL_CrlRemoveCert.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlRemoveCert
func CSSM_CL_CrlRemoveCert(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, OldCrl unsafe.Pointer, NewCrl unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CL_CrlRemoveCert == nil {
		panic("Security: symbol CSSM_CL_CrlRemoveCert not loaded")
	}
	return _cSSM_CL_CrlRemoveCert(CLHandle, Cert, OldCrl, NewCrl)
}

var _cSSM_CL_CrlSetFields func(CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CrlTemplate unsafe.Pointer, OldCrl unsafe.Pointer, ModifiedCrl unsafe.Pointer) CSSM_RETURN

// CSSM_CL_CrlSetFields.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlSetFields
func CSSM_CL_CrlSetFields(CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CrlTemplate unsafe.Pointer, OldCrl unsafe.Pointer, ModifiedCrl unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CL_CrlSetFields == nil {
		panic("Security: symbol CSSM_CL_CrlSetFields not loaded")
	}
	return _cSSM_CL_CrlSetFields(CLHandle, NumberOfFields, CrlTemplate, OldCrl, ModifiedCrl)
}

var _cSSM_CL_CrlSign func(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, UnsignedCrl unsafe.Pointer, SignScope unsafe.Pointer, ScopeSize uint32, SignedCrl unsafe.Pointer) CSSM_RETURN

// CSSM_CL_CrlSign.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlSign
func CSSM_CL_CrlSign(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, UnsignedCrl unsafe.Pointer, SignScope unsafe.Pointer, ScopeSize uint32, SignedCrl unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CL_CrlSign == nil {
		panic("Security: symbol CSSM_CL_CrlSign not loaded")
	}
	return _cSSM_CL_CrlSign(CLHandle, CCHandle, UnsignedCrl, SignScope, ScopeSize, SignedCrl)
}

var _cSSM_CL_CrlVerify func(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CrlToBeVerified unsafe.Pointer, SignerCert unsafe.Pointer, VerifyScope unsafe.Pointer, ScopeSize uint32) CSSM_RETURN

// CSSM_CL_CrlVerify.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlVerify
func CSSM_CL_CrlVerify(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CrlToBeVerified unsafe.Pointer, SignerCert unsafe.Pointer, VerifyScope unsafe.Pointer, ScopeSize uint32) CSSM_RETURN {
	if _cSSM_CL_CrlVerify == nil {
		panic("Security: symbol CSSM_CL_CrlVerify not loaded")
	}
	return _cSSM_CL_CrlVerify(CLHandle, CCHandle, CrlToBeVerified, SignerCert, VerifyScope, ScopeSize)
}

var _cSSM_CL_CrlVerifyWithKey func(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CrlToBeVerified unsafe.Pointer) CSSM_RETURN

// CSSM_CL_CrlVerifyWithKey.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_CrlVerifyWithKey
func CSSM_CL_CrlVerifyWithKey(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CrlToBeVerified unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CL_CrlVerifyWithKey == nil {
		panic("Security: symbol CSSM_CL_CrlVerifyWithKey not loaded")
	}
	return _cSSM_CL_CrlVerifyWithKey(CLHandle, CCHandle, CrlToBeVerified)
}

var _cSSM_CL_FreeFieldValue func(CLHandle CSSM_CL_HANDLE, CertOrCrlOid unsafe.Pointer, Value unsafe.Pointer) CSSM_RETURN

// CSSM_CL_FreeFieldValue.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_FreeFieldValue
func CSSM_CL_FreeFieldValue(CLHandle CSSM_CL_HANDLE, CertOrCrlOid unsafe.Pointer, Value unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CL_FreeFieldValue == nil {
		panic("Security: symbol CSSM_CL_FreeFieldValue not loaded")
	}
	return _cSSM_CL_FreeFieldValue(CLHandle, CertOrCrlOid, Value)
}

var _cSSM_CL_FreeFields func(CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, Fields unsafe.Pointer) CSSM_RETURN

// CSSM_CL_FreeFields.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_FreeFields
func CSSM_CL_FreeFields(CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, Fields unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CL_FreeFields == nil {
		panic("Security: symbol CSSM_CL_FreeFields not loaded")
	}
	return _cSSM_CL_FreeFields(CLHandle, NumberOfFields, Fields)
}

var _cSSM_CL_IsCertInCachedCrl func(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, CrlHandle CSSM_HANDLE, CertFound unsafe.Pointer, CrlRecordIndex unsafe.Pointer) CSSM_RETURN

// CSSM_CL_IsCertInCachedCrl.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_IsCertInCachedCrl
func CSSM_CL_IsCertInCachedCrl(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, CrlHandle CSSM_HANDLE, CertFound unsafe.Pointer, CrlRecordIndex unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CL_IsCertInCachedCrl == nil {
		panic("Security: symbol CSSM_CL_IsCertInCachedCrl not loaded")
	}
	return _cSSM_CL_IsCertInCachedCrl(CLHandle, Cert, CrlHandle, CertFound, CrlRecordIndex)
}

var _cSSM_CL_IsCertInCrl func(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, Crl unsafe.Pointer, CertFound unsafe.Pointer) CSSM_RETURN

// CSSM_CL_IsCertInCrl.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_IsCertInCrl
func CSSM_CL_IsCertInCrl(CLHandle CSSM_CL_HANDLE, Cert unsafe.Pointer, Crl unsafe.Pointer, CertFound unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CL_IsCertInCrl == nil {
		panic("Security: symbol CSSM_CL_IsCertInCrl not loaded")
	}
	return _cSSM_CL_IsCertInCrl(CLHandle, Cert, Crl, CertFound)
}

var _cSSM_CL_PassThrough func(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, PassThroughId uint32, InputParams unsafe.Pointer, OutputParams unsafe.Pointer) CSSM_RETURN

// CSSM_CL_PassThrough.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CL_PassThrough
func CSSM_CL_PassThrough(CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, PassThroughId uint32, InputParams unsafe.Pointer, OutputParams unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CL_PassThrough == nil {
		panic("Security: symbol CSSM_CL_PassThrough not loaded")
	}
	return _cSSM_CL_PassThrough(CLHandle, CCHandle, PassThroughId, InputParams, OutputParams)
}

var _cSSM_CSP_ChangeLoginAcl func(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, AclEdit unsafe.Pointer) CSSM_RETURN

// CSSM_CSP_ChangeLoginAcl.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_ChangeLoginAcl
func CSSM_CSP_ChangeLoginAcl(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, AclEdit unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CSP_ChangeLoginAcl == nil {
		panic("Security: symbol CSSM_CSP_ChangeLoginAcl not loaded")
	}
	return _cSSM_CSP_ChangeLoginAcl(CSPHandle, AccessCred, AclEdit)
}

var _cSSM_CSP_ChangeLoginOwner func(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, NewOwner unsafe.Pointer) CSSM_RETURN

// CSSM_CSP_ChangeLoginOwner.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_ChangeLoginOwner
func CSSM_CSP_ChangeLoginOwner(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, NewOwner unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CSP_ChangeLoginOwner == nil {
		panic("Security: symbol CSSM_CSP_ChangeLoginOwner not loaded")
	}
	return _cSSM_CSP_ChangeLoginOwner(CSPHandle, AccessCred, NewOwner)
}

var _cSSM_CSP_CreateAsymmetricContext func(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, AccessCred unsafe.Pointer, Key unsafe.Pointer, Padding CSSM_PADDING, NewContextHandle unsafe.Pointer) CSSM_RETURN

// CSSM_CSP_CreateAsymmetricContext.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_CreateAsymmetricContext
func CSSM_CSP_CreateAsymmetricContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, AccessCred unsafe.Pointer, Key unsafe.Pointer, Padding CSSM_PADDING, NewContextHandle unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CSP_CreateAsymmetricContext == nil {
		panic("Security: symbol CSSM_CSP_CreateAsymmetricContext not loaded")
	}
	return _cSSM_CSP_CreateAsymmetricContext(CSPHandle, AlgorithmID, AccessCred, Key, Padding, NewContextHandle)
}

var _cSSM_CSP_CreateDeriveKeyContext func(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, DeriveKeyType CSSM_KEY_TYPE, DeriveKeyLengthInBits uint32, AccessCred unsafe.Pointer, BaseKey unsafe.Pointer, IterationCount uint32, Salt unsafe.Pointer, Seed unsafe.Pointer, NewContextHandle unsafe.Pointer) CSSM_RETURN

// CSSM_CSP_CreateDeriveKeyContext.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_CreateDeriveKeyContext
func CSSM_CSP_CreateDeriveKeyContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, DeriveKeyType CSSM_KEY_TYPE, DeriveKeyLengthInBits uint32, AccessCred unsafe.Pointer, BaseKey unsafe.Pointer, IterationCount uint32, Salt unsafe.Pointer, Seed unsafe.Pointer, NewContextHandle unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CSP_CreateDeriveKeyContext == nil {
		panic("Security: symbol CSSM_CSP_CreateDeriveKeyContext not loaded")
	}
	return _cSSM_CSP_CreateDeriveKeyContext(CSPHandle, AlgorithmID, DeriveKeyType, DeriveKeyLengthInBits, AccessCred, BaseKey, IterationCount, Salt, Seed, NewContextHandle)
}

var _cSSM_CSP_CreateDigestContext func(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, NewContextHandle unsafe.Pointer) CSSM_RETURN

// CSSM_CSP_CreateDigestContext.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_CreateDigestContext
func CSSM_CSP_CreateDigestContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, NewContextHandle unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CSP_CreateDigestContext == nil {
		panic("Security: symbol CSSM_CSP_CreateDigestContext not loaded")
	}
	return _cSSM_CSP_CreateDigestContext(CSPHandle, AlgorithmID, NewContextHandle)
}

var _cSSM_CSP_CreateKeyGenContext func(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, KeySizeInBits uint32, Seed unsafe.Pointer, Salt unsafe.Pointer, StartDate unsafe.Pointer, EndDate unsafe.Pointer, Params unsafe.Pointer, NewContextHandle unsafe.Pointer) CSSM_RETURN

// CSSM_CSP_CreateKeyGenContext.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_CreateKeyGenContext
func CSSM_CSP_CreateKeyGenContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, KeySizeInBits uint32, Seed unsafe.Pointer, Salt unsafe.Pointer, StartDate unsafe.Pointer, EndDate unsafe.Pointer, Params unsafe.Pointer, NewContextHandle unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CSP_CreateKeyGenContext == nil {
		panic("Security: symbol CSSM_CSP_CreateKeyGenContext not loaded")
	}
	return _cSSM_CSP_CreateKeyGenContext(CSPHandle, AlgorithmID, KeySizeInBits, Seed, Salt, StartDate, EndDate, Params, NewContextHandle)
}

var _cSSM_CSP_CreateMacContext func(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, Key unsafe.Pointer, NewContextHandle unsafe.Pointer) CSSM_RETURN

// CSSM_CSP_CreateMacContext.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_CreateMacContext
func CSSM_CSP_CreateMacContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, Key unsafe.Pointer, NewContextHandle unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CSP_CreateMacContext == nil {
		panic("Security: symbol CSSM_CSP_CreateMacContext not loaded")
	}
	return _cSSM_CSP_CreateMacContext(CSPHandle, AlgorithmID, Key, NewContextHandle)
}

var _cSSM_CSP_CreatePassThroughContext func(CSPHandle CSSM_CSP_HANDLE, Key unsafe.Pointer, NewContextHandle unsafe.Pointer) CSSM_RETURN

// CSSM_CSP_CreatePassThroughContext.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_CreatePassThroughContext
func CSSM_CSP_CreatePassThroughContext(CSPHandle CSSM_CSP_HANDLE, Key unsafe.Pointer, NewContextHandle unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CSP_CreatePassThroughContext == nil {
		panic("Security: symbol CSSM_CSP_CreatePassThroughContext not loaded")
	}
	return _cSSM_CSP_CreatePassThroughContext(CSPHandle, Key, NewContextHandle)
}

var _cSSM_CSP_CreateRandomGenContext func(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, Seed unsafe.Pointer, Length CSSM_SIZE, NewContextHandle unsafe.Pointer) CSSM_RETURN

// CSSM_CSP_CreateRandomGenContext.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_CreateRandomGenContext
func CSSM_CSP_CreateRandomGenContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, Seed unsafe.Pointer, Length CSSM_SIZE, NewContextHandle unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CSP_CreateRandomGenContext == nil {
		panic("Security: symbol CSSM_CSP_CreateRandomGenContext not loaded")
	}
	return _cSSM_CSP_CreateRandomGenContext(CSPHandle, AlgorithmID, Seed, Length, NewContextHandle)
}

var _cSSM_CSP_CreateSignatureContext func(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, AccessCred unsafe.Pointer, Key unsafe.Pointer, NewContextHandle unsafe.Pointer) CSSM_RETURN

// CSSM_CSP_CreateSignatureContext.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_CreateSignatureContext
func CSSM_CSP_CreateSignatureContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, AccessCred unsafe.Pointer, Key unsafe.Pointer, NewContextHandle unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CSP_CreateSignatureContext == nil {
		panic("Security: symbol CSSM_CSP_CreateSignatureContext not loaded")
	}
	return _cSSM_CSP_CreateSignatureContext(CSPHandle, AlgorithmID, AccessCred, Key, NewContextHandle)
}

var _cSSM_CSP_CreateSymmetricContext func(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, Mode CSSM_ENCRYPT_MODE, AccessCred unsafe.Pointer, Key unsafe.Pointer, InitVector unsafe.Pointer, Padding CSSM_PADDING, Reserved unsafe.Pointer, NewContextHandle unsafe.Pointer) CSSM_RETURN

// CSSM_CSP_CreateSymmetricContext.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_CreateSymmetricContext
func CSSM_CSP_CreateSymmetricContext(CSPHandle CSSM_CSP_HANDLE, AlgorithmID CSSM_ALGORITHMS, Mode CSSM_ENCRYPT_MODE, AccessCred unsafe.Pointer, Key unsafe.Pointer, InitVector unsafe.Pointer, Padding CSSM_PADDING, Reserved unsafe.Pointer, NewContextHandle unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CSP_CreateSymmetricContext == nil {
		panic("Security: symbol CSSM_CSP_CreateSymmetricContext not loaded")
	}
	return _cSSM_CSP_CreateSymmetricContext(CSPHandle, AlgorithmID, Mode, AccessCred, Key, InitVector, Padding, Reserved, NewContextHandle)
}

var _cSSM_CSP_GetLoginAcl func(CSPHandle CSSM_CSP_HANDLE, SelectionTag unsafe.Pointer, NumberOfAclInfos *uint32, AclInfos unsafe.Pointer) CSSM_RETURN

// CSSM_CSP_GetLoginAcl.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_GetLoginAcl
func CSSM_CSP_GetLoginAcl(CSPHandle CSSM_CSP_HANDLE, SelectionTag unsafe.Pointer, NumberOfAclInfos *uint32, AclInfos unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CSP_GetLoginAcl == nil {
		panic("Security: symbol CSSM_CSP_GetLoginAcl not loaded")
	}
	return _cSSM_CSP_GetLoginAcl(CSPHandle, SelectionTag, NumberOfAclInfos, AclInfos)
}

var _cSSM_CSP_GetLoginOwner func(CSPHandle CSSM_CSP_HANDLE, Owner unsafe.Pointer) CSSM_RETURN

// CSSM_CSP_GetLoginOwner.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_GetLoginOwner
func CSSM_CSP_GetLoginOwner(CSPHandle CSSM_CSP_HANDLE, Owner unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CSP_GetLoginOwner == nil {
		panic("Security: symbol CSSM_CSP_GetLoginOwner not loaded")
	}
	return _cSSM_CSP_GetLoginOwner(CSPHandle, Owner)
}

var _cSSM_CSP_GetOperationalStatistics func(CSPHandle CSSM_CSP_HANDLE, Statistics unsafe.Pointer) CSSM_RETURN

// CSSM_CSP_GetOperationalStatistics.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_GetOperationalStatistics
func CSSM_CSP_GetOperationalStatistics(CSPHandle CSSM_CSP_HANDLE, Statistics unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CSP_GetOperationalStatistics == nil {
		panic("Security: symbol CSSM_CSP_GetOperationalStatistics not loaded")
	}
	return _cSSM_CSP_GetOperationalStatistics(CSPHandle, Statistics)
}

var _cSSM_CSP_Login func(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, LoginName unsafe.Pointer, Reserved unsafe.Pointer) CSSM_RETURN

// CSSM_CSP_Login.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_Login
func CSSM_CSP_Login(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, LoginName unsafe.Pointer, Reserved unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CSP_Login == nil {
		panic("Security: symbol CSSM_CSP_Login not loaded")
	}
	return _cSSM_CSP_Login(CSPHandle, AccessCred, LoginName, Reserved)
}

var _cSSM_CSP_Logout func(CSPHandle CSSM_CSP_HANDLE) CSSM_RETURN

// CSSM_CSP_Logout.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_Logout
func CSSM_CSP_Logout(CSPHandle CSSM_CSP_HANDLE) CSSM_RETURN {
	if _cSSM_CSP_Logout == nil {
		panic("Security: symbol CSSM_CSP_Logout not loaded")
	}
	return _cSSM_CSP_Logout(CSPHandle)
}

var _cSSM_CSP_ObtainPrivateKeyFromPublicKey func(CSPHandle CSSM_CSP_HANDLE, PublicKey unsafe.Pointer, PrivateKey unsafe.Pointer) CSSM_RETURN

// CSSM_CSP_ObtainPrivateKeyFromPublicKey.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_ObtainPrivateKeyFromPublicKey
func CSSM_CSP_ObtainPrivateKeyFromPublicKey(CSPHandle CSSM_CSP_HANDLE, PublicKey unsafe.Pointer, PrivateKey unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CSP_ObtainPrivateKeyFromPublicKey == nil {
		panic("Security: symbol CSSM_CSP_ObtainPrivateKeyFromPublicKey not loaded")
	}
	return _cSSM_CSP_ObtainPrivateKeyFromPublicKey(CSPHandle, PublicKey, PrivateKey)
}

var _cSSM_CSP_PassThrough func(CCHandle CSSM_CC_HANDLE, PassThroughId uint32, InData unsafe.Pointer, OutData unsafe.Pointer) CSSM_RETURN

// CSSM_CSP_PassThrough.
//
// See: https://developer.apple.com/documentation/Security/CSSM_CSP_PassThrough
func CSSM_CSP_PassThrough(CCHandle CSSM_CC_HANDLE, PassThroughId uint32, InData unsafe.Pointer, OutData unsafe.Pointer) CSSM_RETURN {
	if _cSSM_CSP_PassThrough == nil {
		panic("Security: symbol CSSM_CSP_PassThrough not loaded")
	}
	return _cSSM_CSP_PassThrough(CCHandle, PassThroughId, InData, OutData)
}

var _cSSM_ChangeKeyAcl func(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, AclEdit unsafe.Pointer, Key unsafe.Pointer) CSSM_RETURN

// CSSM_ChangeKeyAcl.
//
// See: https://developer.apple.com/documentation/Security/CSSM_ChangeKeyAcl
func CSSM_ChangeKeyAcl(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, AclEdit unsafe.Pointer, Key unsafe.Pointer) CSSM_RETURN {
	if _cSSM_ChangeKeyAcl == nil {
		panic("Security: symbol CSSM_ChangeKeyAcl not loaded")
	}
	return _cSSM_ChangeKeyAcl(CSPHandle, AccessCred, AclEdit, Key)
}

var _cSSM_ChangeKeyOwner func(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, Key unsafe.Pointer, NewOwner unsafe.Pointer) CSSM_RETURN

// CSSM_ChangeKeyOwner.
//
// See: https://developer.apple.com/documentation/Security/CSSM_ChangeKeyOwner
func CSSM_ChangeKeyOwner(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, Key unsafe.Pointer, NewOwner unsafe.Pointer) CSSM_RETURN {
	if _cSSM_ChangeKeyOwner == nil {
		panic("Security: symbol CSSM_ChangeKeyOwner not loaded")
	}
	return _cSSM_ChangeKeyOwner(CSPHandle, AccessCred, Key, NewOwner)
}

var _cSSM_DL_Authenticate func(DLDBHandle unsafe.Pointer, AccessRequest CSSM_DB_ACCESS_TYPE, AccessCred unsafe.Pointer) CSSM_RETURN

// CSSM_DL_Authenticate.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_Authenticate
func CSSM_DL_Authenticate(DLDBHandle unsafe.Pointer, AccessRequest CSSM_DB_ACCESS_TYPE, AccessCred unsafe.Pointer) CSSM_RETURN {
	if _cSSM_DL_Authenticate == nil {
		panic("Security: symbol CSSM_DL_Authenticate not loaded")
	}
	return _cSSM_DL_Authenticate(DLDBHandle, AccessRequest, AccessCred)
}

var _cSSM_DL_ChangeDbAcl func(DLDBHandle unsafe.Pointer, AccessCred unsafe.Pointer, AclEdit unsafe.Pointer) CSSM_RETURN

// CSSM_DL_ChangeDbAcl.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_ChangeDbAcl
func CSSM_DL_ChangeDbAcl(DLDBHandle unsafe.Pointer, AccessCred unsafe.Pointer, AclEdit unsafe.Pointer) CSSM_RETURN {
	if _cSSM_DL_ChangeDbAcl == nil {
		panic("Security: symbol CSSM_DL_ChangeDbAcl not loaded")
	}
	return _cSSM_DL_ChangeDbAcl(DLDBHandle, AccessCred, AclEdit)
}

var _cSSM_DL_ChangeDbOwner func(DLDBHandle unsafe.Pointer, AccessCred unsafe.Pointer, NewOwner unsafe.Pointer) CSSM_RETURN

// CSSM_DL_ChangeDbOwner.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_ChangeDbOwner
func CSSM_DL_ChangeDbOwner(DLDBHandle unsafe.Pointer, AccessCred unsafe.Pointer, NewOwner unsafe.Pointer) CSSM_RETURN {
	if _cSSM_DL_ChangeDbOwner == nil {
		panic("Security: symbol CSSM_DL_ChangeDbOwner not loaded")
	}
	return _cSSM_DL_ChangeDbOwner(DLDBHandle, AccessCred, NewOwner)
}

var _cSSM_DL_CreateRelation func(DLDBHandle unsafe.Pointer, RelationID CSSM_DB_RECORDTYPE, RelationName *byte, NumberOfAttributes uint32, pAttributeInfo unsafe.Pointer, NumberOfIndexes uint32, pIndexInfo unsafe.Pointer) CSSM_RETURN

// CSSM_DL_CreateRelation.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_CreateRelation
func CSSM_DL_CreateRelation(DLDBHandle unsafe.Pointer, RelationID CSSM_DB_RECORDTYPE, RelationName *byte, NumberOfAttributes uint32, pAttributeInfo unsafe.Pointer, NumberOfIndexes uint32, pIndexInfo unsafe.Pointer) CSSM_RETURN {
	if _cSSM_DL_CreateRelation == nil {
		panic("Security: symbol CSSM_DL_CreateRelation not loaded")
	}
	return _cSSM_DL_CreateRelation(DLDBHandle, RelationID, RelationName, NumberOfAttributes, pAttributeInfo, NumberOfIndexes, pIndexInfo)
}

var _cSSM_DL_DataAbortQuery func(DLDBHandle unsafe.Pointer, ResultsHandle CSSM_HANDLE) CSSM_RETURN

// CSSM_DL_DataAbortQuery.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_DataAbortQuery
func CSSM_DL_DataAbortQuery(DLDBHandle unsafe.Pointer, ResultsHandle CSSM_HANDLE) CSSM_RETURN {
	if _cSSM_DL_DataAbortQuery == nil {
		panic("Security: symbol CSSM_DL_DataAbortQuery not loaded")
	}
	return _cSSM_DL_DataAbortQuery(DLDBHandle, ResultsHandle)
}

var _cSSM_DL_DataDelete func(DLDBHandle unsafe.Pointer, UniqueRecordIdentifier unsafe.Pointer) CSSM_RETURN

// CSSM_DL_DataDelete.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_DataDelete
func CSSM_DL_DataDelete(DLDBHandle unsafe.Pointer, UniqueRecordIdentifier unsafe.Pointer) CSSM_RETURN {
	if _cSSM_DL_DataDelete == nil {
		panic("Security: symbol CSSM_DL_DataDelete not loaded")
	}
	return _cSSM_DL_DataDelete(DLDBHandle, UniqueRecordIdentifier)
}

var _cSSM_DL_DataGetFirst func(DLDBHandle unsafe.Pointer, Query unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, Attributes unsafe.Pointer, Data unsafe.Pointer, UniqueId unsafe.Pointer) CSSM_RETURN

// CSSM_DL_DataGetFirst.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_DataGetFirst
func CSSM_DL_DataGetFirst(DLDBHandle unsafe.Pointer, Query unsafe.Pointer, ResultsHandle CSSM_HANDLE_PTR, Attributes unsafe.Pointer, Data unsafe.Pointer, UniqueId unsafe.Pointer) CSSM_RETURN {
	if _cSSM_DL_DataGetFirst == nil {
		panic("Security: symbol CSSM_DL_DataGetFirst not loaded")
	}
	return _cSSM_DL_DataGetFirst(DLDBHandle, Query, ResultsHandle, Attributes, Data, UniqueId)
}

var _cSSM_DL_DataGetFromUniqueRecordId func(DLDBHandle unsafe.Pointer, UniqueRecord unsafe.Pointer, Attributes unsafe.Pointer, Data unsafe.Pointer) CSSM_RETURN

// CSSM_DL_DataGetFromUniqueRecordId.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_DataGetFromUniqueRecordId
func CSSM_DL_DataGetFromUniqueRecordId(DLDBHandle unsafe.Pointer, UniqueRecord unsafe.Pointer, Attributes unsafe.Pointer, Data unsafe.Pointer) CSSM_RETURN {
	if _cSSM_DL_DataGetFromUniqueRecordId == nil {
		panic("Security: symbol CSSM_DL_DataGetFromUniqueRecordId not loaded")
	}
	return _cSSM_DL_DataGetFromUniqueRecordId(DLDBHandle, UniqueRecord, Attributes, Data)
}

var _cSSM_DL_DataGetNext func(DLDBHandle unsafe.Pointer, ResultsHandle CSSM_HANDLE, Attributes unsafe.Pointer, Data unsafe.Pointer, UniqueId unsafe.Pointer) CSSM_RETURN

// CSSM_DL_DataGetNext.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_DataGetNext
func CSSM_DL_DataGetNext(DLDBHandle unsafe.Pointer, ResultsHandle CSSM_HANDLE, Attributes unsafe.Pointer, Data unsafe.Pointer, UniqueId unsafe.Pointer) CSSM_RETURN {
	if _cSSM_DL_DataGetNext == nil {
		panic("Security: symbol CSSM_DL_DataGetNext not loaded")
	}
	return _cSSM_DL_DataGetNext(DLDBHandle, ResultsHandle, Attributes, Data, UniqueId)
}

var _cSSM_DL_DataInsert func(DLDBHandle unsafe.Pointer, RecordType CSSM_DB_RECORDTYPE, Attributes unsafe.Pointer, Data unsafe.Pointer, UniqueId unsafe.Pointer) CSSM_RETURN

// CSSM_DL_DataInsert.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_DataInsert
func CSSM_DL_DataInsert(DLDBHandle unsafe.Pointer, RecordType CSSM_DB_RECORDTYPE, Attributes unsafe.Pointer, Data unsafe.Pointer, UniqueId unsafe.Pointer) CSSM_RETURN {
	if _cSSM_DL_DataInsert == nil {
		panic("Security: symbol CSSM_DL_DataInsert not loaded")
	}
	return _cSSM_DL_DataInsert(DLDBHandle, RecordType, Attributes, Data, UniqueId)
}

var _cSSM_DL_DataModify func(DLDBHandle unsafe.Pointer, RecordType CSSM_DB_RECORDTYPE, UniqueRecordIdentifier unsafe.Pointer, AttributesToBeModified unsafe.Pointer, DataToBeModified unsafe.Pointer, ModifyMode CSSM_DB_MODIFY_MODE) CSSM_RETURN

// CSSM_DL_DataModify.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_DataModify
func CSSM_DL_DataModify(DLDBHandle unsafe.Pointer, RecordType CSSM_DB_RECORDTYPE, UniqueRecordIdentifier unsafe.Pointer, AttributesToBeModified unsafe.Pointer, DataToBeModified unsafe.Pointer, ModifyMode CSSM_DB_MODIFY_MODE) CSSM_RETURN {
	if _cSSM_DL_DataModify == nil {
		panic("Security: symbol CSSM_DL_DataModify not loaded")
	}
	return _cSSM_DL_DataModify(DLDBHandle, RecordType, UniqueRecordIdentifier, AttributesToBeModified, DataToBeModified, ModifyMode)
}

var _cSSM_DL_DbClose func(DLDBHandle unsafe.Pointer) CSSM_RETURN

// CSSM_DL_DbClose.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_DbClose
func CSSM_DL_DbClose(DLDBHandle unsafe.Pointer) CSSM_RETURN {
	if _cSSM_DL_DbClose == nil {
		panic("Security: symbol CSSM_DL_DbClose not loaded")
	}
	return _cSSM_DL_DbClose(DLDBHandle)
}

var _cSSM_DL_DbCreate func(DLHandle CSSM_DL_HANDLE, DbName *byte, DbLocation unsafe.Pointer, DBInfo unsafe.Pointer, AccessRequest CSSM_DB_ACCESS_TYPE, CredAndAclEntry unsafe.Pointer, OpenParameters unsafe.Pointer, DbHandle unsafe.Pointer) CSSM_RETURN

// CSSM_DL_DbCreate.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_DbCreate
func CSSM_DL_DbCreate(DLHandle CSSM_DL_HANDLE, DbName *byte, DbLocation unsafe.Pointer, DBInfo unsafe.Pointer, AccessRequest CSSM_DB_ACCESS_TYPE, CredAndAclEntry unsafe.Pointer, OpenParameters unsafe.Pointer, DbHandle unsafe.Pointer) CSSM_RETURN {
	if _cSSM_DL_DbCreate == nil {
		panic("Security: symbol CSSM_DL_DbCreate not loaded")
	}
	return _cSSM_DL_DbCreate(DLHandle, DbName, DbLocation, DBInfo, AccessRequest, CredAndAclEntry, OpenParameters, DbHandle)
}

var _cSSM_DL_DbDelete func(DLHandle CSSM_DL_HANDLE, DbName *byte, DbLocation unsafe.Pointer, AccessCred unsafe.Pointer) CSSM_RETURN

// CSSM_DL_DbDelete.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_DbDelete
func CSSM_DL_DbDelete(DLHandle CSSM_DL_HANDLE, DbName *byte, DbLocation unsafe.Pointer, AccessCred unsafe.Pointer) CSSM_RETURN {
	if _cSSM_DL_DbDelete == nil {
		panic("Security: symbol CSSM_DL_DbDelete not loaded")
	}
	return _cSSM_DL_DbDelete(DLHandle, DbName, DbLocation, AccessCred)
}

var _cSSM_DL_DbOpen func(DLHandle CSSM_DL_HANDLE, DbName *byte, DbLocation unsafe.Pointer, AccessRequest CSSM_DB_ACCESS_TYPE, AccessCred unsafe.Pointer, OpenParameters unsafe.Pointer, DbHandle unsafe.Pointer) CSSM_RETURN

// CSSM_DL_DbOpen.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_DbOpen
func CSSM_DL_DbOpen(DLHandle CSSM_DL_HANDLE, DbName *byte, DbLocation unsafe.Pointer, AccessRequest CSSM_DB_ACCESS_TYPE, AccessCred unsafe.Pointer, OpenParameters unsafe.Pointer, DbHandle unsafe.Pointer) CSSM_RETURN {
	if _cSSM_DL_DbOpen == nil {
		panic("Security: symbol CSSM_DL_DbOpen not loaded")
	}
	return _cSSM_DL_DbOpen(DLHandle, DbName, DbLocation, AccessRequest, AccessCred, OpenParameters, DbHandle)
}

var _cSSM_DL_DestroyRelation func(DLDBHandle unsafe.Pointer, RelationID CSSM_DB_RECORDTYPE) CSSM_RETURN

// CSSM_DL_DestroyRelation.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_DestroyRelation
func CSSM_DL_DestroyRelation(DLDBHandle unsafe.Pointer, RelationID CSSM_DB_RECORDTYPE) CSSM_RETURN {
	if _cSSM_DL_DestroyRelation == nil {
		panic("Security: symbol CSSM_DL_DestroyRelation not loaded")
	}
	return _cSSM_DL_DestroyRelation(DLDBHandle, RelationID)
}

var _cSSM_DL_FreeNameList func(DLHandle CSSM_DL_HANDLE, NameList unsafe.Pointer) CSSM_RETURN

// CSSM_DL_FreeNameList.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_FreeNameList
func CSSM_DL_FreeNameList(DLHandle CSSM_DL_HANDLE, NameList unsafe.Pointer) CSSM_RETURN {
	if _cSSM_DL_FreeNameList == nil {
		panic("Security: symbol CSSM_DL_FreeNameList not loaded")
	}
	return _cSSM_DL_FreeNameList(DLHandle, NameList)
}

var _cSSM_DL_FreeUniqueRecord func(DLDBHandle unsafe.Pointer, UniqueRecord unsafe.Pointer) CSSM_RETURN

// CSSM_DL_FreeUniqueRecord.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_FreeUniqueRecord
func CSSM_DL_FreeUniqueRecord(DLDBHandle unsafe.Pointer, UniqueRecord unsafe.Pointer) CSSM_RETURN {
	if _cSSM_DL_FreeUniqueRecord == nil {
		panic("Security: symbol CSSM_DL_FreeUniqueRecord not loaded")
	}
	return _cSSM_DL_FreeUniqueRecord(DLDBHandle, UniqueRecord)
}

var _cSSM_DL_GetDbAcl func(DLDBHandle unsafe.Pointer, SelectionTag unsafe.Pointer, NumberOfAclInfos *uint32, AclInfos unsafe.Pointer) CSSM_RETURN

// CSSM_DL_GetDbAcl.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_GetDbAcl
func CSSM_DL_GetDbAcl(DLDBHandle unsafe.Pointer, SelectionTag unsafe.Pointer, NumberOfAclInfos *uint32, AclInfos unsafe.Pointer) CSSM_RETURN {
	if _cSSM_DL_GetDbAcl == nil {
		panic("Security: symbol CSSM_DL_GetDbAcl not loaded")
	}
	return _cSSM_DL_GetDbAcl(DLDBHandle, SelectionTag, NumberOfAclInfos, AclInfos)
}

var _cSSM_DL_GetDbNameFromHandle func(DLDBHandle unsafe.Pointer, DbName *byte) CSSM_RETURN

// CSSM_DL_GetDbNameFromHandle.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_GetDbNameFromHandle
func CSSM_DL_GetDbNameFromHandle(DLDBHandle unsafe.Pointer, DbName *byte) CSSM_RETURN {
	if _cSSM_DL_GetDbNameFromHandle == nil {
		panic("Security: symbol CSSM_DL_GetDbNameFromHandle not loaded")
	}
	return _cSSM_DL_GetDbNameFromHandle(DLDBHandle, DbName)
}

var _cSSM_DL_GetDbNames func(DLHandle CSSM_DL_HANDLE, NameList unsafe.Pointer) CSSM_RETURN

// CSSM_DL_GetDbNames.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_GetDbNames
func CSSM_DL_GetDbNames(DLHandle CSSM_DL_HANDLE, NameList unsafe.Pointer) CSSM_RETURN {
	if _cSSM_DL_GetDbNames == nil {
		panic("Security: symbol CSSM_DL_GetDbNames not loaded")
	}
	return _cSSM_DL_GetDbNames(DLHandle, NameList)
}

var _cSSM_DL_GetDbOwner func(DLDBHandle unsafe.Pointer, Owner unsafe.Pointer) CSSM_RETURN

// CSSM_DL_GetDbOwner.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_GetDbOwner
func CSSM_DL_GetDbOwner(DLDBHandle unsafe.Pointer, Owner unsafe.Pointer) CSSM_RETURN {
	if _cSSM_DL_GetDbOwner == nil {
		panic("Security: symbol CSSM_DL_GetDbOwner not loaded")
	}
	return _cSSM_DL_GetDbOwner(DLDBHandle, Owner)
}

var _cSSM_DL_PassThrough func(DLDBHandle unsafe.Pointer, PassThroughId uint32, InputParams unsafe.Pointer, OutputParams unsafe.Pointer) CSSM_RETURN

// CSSM_DL_PassThrough.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DL_PassThrough
func CSSM_DL_PassThrough(DLDBHandle unsafe.Pointer, PassThroughId uint32, InputParams unsafe.Pointer, OutputParams unsafe.Pointer) CSSM_RETURN {
	if _cSSM_DL_PassThrough == nil {
		panic("Security: symbol CSSM_DL_PassThrough not loaded")
	}
	return _cSSM_DL_PassThrough(DLDBHandle, PassThroughId, InputParams, OutputParams)
}

var _cSSM_DecryptData func(CCHandle CSSM_CC_HANDLE, CipherBufs unsafe.Pointer, CipherBufCount uint32, ClearBufs unsafe.Pointer, ClearBufCount uint32, bytesDecrypted unsafe.Pointer, RemData unsafe.Pointer) CSSM_RETURN

// CSSM_DecryptData.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DecryptData
func CSSM_DecryptData(CCHandle CSSM_CC_HANDLE, CipherBufs unsafe.Pointer, CipherBufCount uint32, ClearBufs unsafe.Pointer, ClearBufCount uint32, bytesDecrypted unsafe.Pointer, RemData unsafe.Pointer) CSSM_RETURN {
	if _cSSM_DecryptData == nil {
		panic("Security: symbol CSSM_DecryptData not loaded")
	}
	return _cSSM_DecryptData(CCHandle, CipherBufs, CipherBufCount, ClearBufs, ClearBufCount, bytesDecrypted, RemData)
}

var _cSSM_DecryptDataFinal func(CCHandle CSSM_CC_HANDLE, RemData unsafe.Pointer) CSSM_RETURN

// CSSM_DecryptDataFinal.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DecryptDataFinal
func CSSM_DecryptDataFinal(CCHandle CSSM_CC_HANDLE, RemData unsafe.Pointer) CSSM_RETURN {
	if _cSSM_DecryptDataFinal == nil {
		panic("Security: symbol CSSM_DecryptDataFinal not loaded")
	}
	return _cSSM_DecryptDataFinal(CCHandle, RemData)
}

var _cSSM_DecryptDataInit func(CCHandle CSSM_CC_HANDLE) CSSM_RETURN

// CSSM_DecryptDataInit.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DecryptDataInit
func CSSM_DecryptDataInit(CCHandle CSSM_CC_HANDLE) CSSM_RETURN {
	if _cSSM_DecryptDataInit == nil {
		panic("Security: symbol CSSM_DecryptDataInit not loaded")
	}
	return _cSSM_DecryptDataInit(CCHandle)
}

var _cSSM_DecryptDataInitP func(CCHandle CSSM_CC_HANDLE, Privilege CSSM_PRIVILEGE) CSSM_RETURN

// CSSM_DecryptDataInitP.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DecryptDataInitP
func CSSM_DecryptDataInitP(CCHandle CSSM_CC_HANDLE, Privilege CSSM_PRIVILEGE) CSSM_RETURN {
	if _cSSM_DecryptDataInitP == nil {
		panic("Security: symbol CSSM_DecryptDataInitP not loaded")
	}
	return _cSSM_DecryptDataInitP(CCHandle, Privilege)
}

var _cSSM_DecryptDataP func(CCHandle CSSM_CC_HANDLE, CipherBufs unsafe.Pointer, CipherBufCount uint32, ClearBufs unsafe.Pointer, ClearBufCount uint32, bytesDecrypted unsafe.Pointer, RemData unsafe.Pointer, Privilege CSSM_PRIVILEGE) CSSM_RETURN

// CSSM_DecryptDataP.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DecryptDataP
func CSSM_DecryptDataP(CCHandle CSSM_CC_HANDLE, CipherBufs unsafe.Pointer, CipherBufCount uint32, ClearBufs unsafe.Pointer, ClearBufCount uint32, bytesDecrypted unsafe.Pointer, RemData unsafe.Pointer, Privilege CSSM_PRIVILEGE) CSSM_RETURN {
	if _cSSM_DecryptDataP == nil {
		panic("Security: symbol CSSM_DecryptDataP not loaded")
	}
	return _cSSM_DecryptDataP(CCHandle, CipherBufs, CipherBufCount, ClearBufs, ClearBufCount, bytesDecrypted, RemData, Privilege)
}

var _cSSM_DecryptDataUpdate func(CCHandle CSSM_CC_HANDLE, CipherBufs unsafe.Pointer, CipherBufCount uint32, ClearBufs unsafe.Pointer, ClearBufCount uint32, bytesDecrypted unsafe.Pointer) CSSM_RETURN

// CSSM_DecryptDataUpdate.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DecryptDataUpdate
func CSSM_DecryptDataUpdate(CCHandle CSSM_CC_HANDLE, CipherBufs unsafe.Pointer, CipherBufCount uint32, ClearBufs unsafe.Pointer, ClearBufCount uint32, bytesDecrypted unsafe.Pointer) CSSM_RETURN {
	if _cSSM_DecryptDataUpdate == nil {
		panic("Security: symbol CSSM_DecryptDataUpdate not loaded")
	}
	return _cSSM_DecryptDataUpdate(CCHandle, CipherBufs, CipherBufCount, ClearBufs, ClearBufCount, bytesDecrypted)
}

var _cSSM_DeleteContext func(CCHandle CSSM_CC_HANDLE) CSSM_RETURN

// CSSM_DeleteContext.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DeleteContext
func CSSM_DeleteContext(CCHandle CSSM_CC_HANDLE) CSSM_RETURN {
	if _cSSM_DeleteContext == nil {
		panic("Security: symbol CSSM_DeleteContext not loaded")
	}
	return _cSSM_DeleteContext(CCHandle)
}

var _cSSM_DeleteContextAttributes func(CCHandle CSSM_CC_HANDLE, NumberOfAttributes uint32, ContextAttributes unsafe.Pointer) CSSM_RETURN

// CSSM_DeleteContextAttributes.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DeleteContextAttributes
func CSSM_DeleteContextAttributes(CCHandle CSSM_CC_HANDLE, NumberOfAttributes uint32, ContextAttributes unsafe.Pointer) CSSM_RETURN {
	if _cSSM_DeleteContextAttributes == nil {
		panic("Security: symbol CSSM_DeleteContextAttributes not loaded")
	}
	return _cSSM_DeleteContextAttributes(CCHandle, NumberOfAttributes, ContextAttributes)
}

var _cSSM_DeriveKey func(CCHandle CSSM_CC_HANDLE, Param unsafe.Pointer, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, DerivedKey unsafe.Pointer) CSSM_RETURN

// CSSM_DeriveKey.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DeriveKey
func CSSM_DeriveKey(CCHandle CSSM_CC_HANDLE, Param unsafe.Pointer, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, DerivedKey unsafe.Pointer) CSSM_RETURN {
	if _cSSM_DeriveKey == nil {
		panic("Security: symbol CSSM_DeriveKey not loaded")
	}
	return _cSSM_DeriveKey(CCHandle, Param, KeyUsage, KeyAttr, KeyLabel, CredAndAclEntry, DerivedKey)
}

var _cSSM_DigestData func(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, Digest unsafe.Pointer) CSSM_RETURN

// CSSM_DigestData.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DigestData
func CSSM_DigestData(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, Digest unsafe.Pointer) CSSM_RETURN {
	if _cSSM_DigestData == nil {
		panic("Security: symbol CSSM_DigestData not loaded")
	}
	return _cSSM_DigestData(CCHandle, DataBufs, DataBufCount, Digest)
}

var _cSSM_DigestDataClone func(CCHandle CSSM_CC_HANDLE, ClonednewCCHandle unsafe.Pointer) CSSM_RETURN

// CSSM_DigestDataClone.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DigestDataClone
func CSSM_DigestDataClone(CCHandle CSSM_CC_HANDLE, ClonednewCCHandle unsafe.Pointer) CSSM_RETURN {
	if _cSSM_DigestDataClone == nil {
		panic("Security: symbol CSSM_DigestDataClone not loaded")
	}
	return _cSSM_DigestDataClone(CCHandle, ClonednewCCHandle)
}

var _cSSM_DigestDataFinal func(CCHandle CSSM_CC_HANDLE, Digest unsafe.Pointer) CSSM_RETURN

// CSSM_DigestDataFinal.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DigestDataFinal
func CSSM_DigestDataFinal(CCHandle CSSM_CC_HANDLE, Digest unsafe.Pointer) CSSM_RETURN {
	if _cSSM_DigestDataFinal == nil {
		panic("Security: symbol CSSM_DigestDataFinal not loaded")
	}
	return _cSSM_DigestDataFinal(CCHandle, Digest)
}

var _cSSM_DigestDataInit func(CCHandle CSSM_CC_HANDLE) CSSM_RETURN

// CSSM_DigestDataInit.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DigestDataInit
func CSSM_DigestDataInit(CCHandle CSSM_CC_HANDLE) CSSM_RETURN {
	if _cSSM_DigestDataInit == nil {
		panic("Security: symbol CSSM_DigestDataInit not loaded")
	}
	return _cSSM_DigestDataInit(CCHandle)
}

var _cSSM_DigestDataUpdate func(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) CSSM_RETURN

// CSSM_DigestDataUpdate.
//
// See: https://developer.apple.com/documentation/Security/CSSM_DigestDataUpdate
func CSSM_DigestDataUpdate(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) CSSM_RETURN {
	if _cSSM_DigestDataUpdate == nil {
		panic("Security: symbol CSSM_DigestDataUpdate not loaded")
	}
	return _cSSM_DigestDataUpdate(CCHandle, DataBufs, DataBufCount)
}

var _cSSM_EncryptData func(CCHandle CSSM_CC_HANDLE, ClearBufs unsafe.Pointer, ClearBufCount uint32, CipherBufs unsafe.Pointer, CipherBufCount uint32, bytesEncrypted unsafe.Pointer, RemData unsafe.Pointer) CSSM_RETURN

// CSSM_EncryptData.
//
// See: https://developer.apple.com/documentation/Security/CSSM_EncryptData
func CSSM_EncryptData(CCHandle CSSM_CC_HANDLE, ClearBufs unsafe.Pointer, ClearBufCount uint32, CipherBufs unsafe.Pointer, CipherBufCount uint32, bytesEncrypted unsafe.Pointer, RemData unsafe.Pointer) CSSM_RETURN {
	if _cSSM_EncryptData == nil {
		panic("Security: symbol CSSM_EncryptData not loaded")
	}
	return _cSSM_EncryptData(CCHandle, ClearBufs, ClearBufCount, CipherBufs, CipherBufCount, bytesEncrypted, RemData)
}

var _cSSM_EncryptDataFinal func(CCHandle CSSM_CC_HANDLE, RemData unsafe.Pointer) CSSM_RETURN

// CSSM_EncryptDataFinal.
//
// See: https://developer.apple.com/documentation/Security/CSSM_EncryptDataFinal
func CSSM_EncryptDataFinal(CCHandle CSSM_CC_HANDLE, RemData unsafe.Pointer) CSSM_RETURN {
	if _cSSM_EncryptDataFinal == nil {
		panic("Security: symbol CSSM_EncryptDataFinal not loaded")
	}
	return _cSSM_EncryptDataFinal(CCHandle, RemData)
}

var _cSSM_EncryptDataInit func(CCHandle CSSM_CC_HANDLE) CSSM_RETURN

// CSSM_EncryptDataInit.
//
// See: https://developer.apple.com/documentation/Security/CSSM_EncryptDataInit
func CSSM_EncryptDataInit(CCHandle CSSM_CC_HANDLE) CSSM_RETURN {
	if _cSSM_EncryptDataInit == nil {
		panic("Security: symbol CSSM_EncryptDataInit not loaded")
	}
	return _cSSM_EncryptDataInit(CCHandle)
}

var _cSSM_EncryptDataInitP func(CCHandle CSSM_CC_HANDLE, Privilege CSSM_PRIVILEGE) CSSM_RETURN

// CSSM_EncryptDataInitP.
//
// See: https://developer.apple.com/documentation/Security/CSSM_EncryptDataInitP
func CSSM_EncryptDataInitP(CCHandle CSSM_CC_HANDLE, Privilege CSSM_PRIVILEGE) CSSM_RETURN {
	if _cSSM_EncryptDataInitP == nil {
		panic("Security: symbol CSSM_EncryptDataInitP not loaded")
	}
	return _cSSM_EncryptDataInitP(CCHandle, Privilege)
}

var _cSSM_EncryptDataP func(CCHandle CSSM_CC_HANDLE, ClearBufs unsafe.Pointer, ClearBufCount uint32, CipherBufs unsafe.Pointer, CipherBufCount uint32, bytesEncrypted unsafe.Pointer, RemData unsafe.Pointer, Privilege CSSM_PRIVILEGE) CSSM_RETURN

// CSSM_EncryptDataP.
//
// See: https://developer.apple.com/documentation/Security/CSSM_EncryptDataP
func CSSM_EncryptDataP(CCHandle CSSM_CC_HANDLE, ClearBufs unsafe.Pointer, ClearBufCount uint32, CipherBufs unsafe.Pointer, CipherBufCount uint32, bytesEncrypted unsafe.Pointer, RemData unsafe.Pointer, Privilege CSSM_PRIVILEGE) CSSM_RETURN {
	if _cSSM_EncryptDataP == nil {
		panic("Security: symbol CSSM_EncryptDataP not loaded")
	}
	return _cSSM_EncryptDataP(CCHandle, ClearBufs, ClearBufCount, CipherBufs, CipherBufCount, bytesEncrypted, RemData, Privilege)
}

var _cSSM_EncryptDataUpdate func(CCHandle CSSM_CC_HANDLE, ClearBufs unsafe.Pointer, ClearBufCount uint32, CipherBufs unsafe.Pointer, CipherBufCount uint32, bytesEncrypted unsafe.Pointer) CSSM_RETURN

// CSSM_EncryptDataUpdate.
//
// See: https://developer.apple.com/documentation/Security/CSSM_EncryptDataUpdate
func CSSM_EncryptDataUpdate(CCHandle CSSM_CC_HANDLE, ClearBufs unsafe.Pointer, ClearBufCount uint32, CipherBufs unsafe.Pointer, CipherBufCount uint32, bytesEncrypted unsafe.Pointer) CSSM_RETURN {
	if _cSSM_EncryptDataUpdate == nil {
		panic("Security: symbol CSSM_EncryptDataUpdate not loaded")
	}
	return _cSSM_EncryptDataUpdate(CCHandle, ClearBufs, ClearBufCount, CipherBufs, CipherBufCount, bytesEncrypted)
}

var _cSSM_FreeContext func(Context unsafe.Pointer) CSSM_RETURN

// CSSM_FreeContext.
//
// See: https://developer.apple.com/documentation/Security/CSSM_FreeContext
func CSSM_FreeContext(Context unsafe.Pointer) CSSM_RETURN {
	if _cSSM_FreeContext == nil {
		panic("Security: symbol CSSM_FreeContext not loaded")
	}
	return _cSSM_FreeContext(Context)
}

var _cSSM_FreeKey func(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, KeyPtr unsafe.Pointer, Delete CSSM_BOOL) CSSM_RETURN

// CSSM_FreeKey.
//
// See: https://developer.apple.com/documentation/Security/CSSM_FreeKey
func CSSM_FreeKey(CSPHandle CSSM_CSP_HANDLE, AccessCred unsafe.Pointer, KeyPtr unsafe.Pointer, Delete CSSM_BOOL) CSSM_RETURN {
	if _cSSM_FreeKey == nil {
		panic("Security: symbol CSSM_FreeKey not loaded")
	}
	return _cSSM_FreeKey(CSPHandle, AccessCred, KeyPtr, Delete)
}

var _cSSM_GenerateAlgorithmParams func(CCHandle CSSM_CC_HANDLE, ParamBits uint32, Param unsafe.Pointer) CSSM_RETURN

// CSSM_GenerateAlgorithmParams.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GenerateAlgorithmParams
func CSSM_GenerateAlgorithmParams(CCHandle CSSM_CC_HANDLE, ParamBits uint32, Param unsafe.Pointer) CSSM_RETURN {
	if _cSSM_GenerateAlgorithmParams == nil {
		panic("Security: symbol CSSM_GenerateAlgorithmParams not loaded")
	}
	return _cSSM_GenerateAlgorithmParams(CCHandle, ParamBits, Param)
}

var _cSSM_GenerateKey func(CCHandle CSSM_CC_HANDLE, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, Key unsafe.Pointer) CSSM_RETURN

// CSSM_GenerateKey.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GenerateKey
func CSSM_GenerateKey(CCHandle CSSM_CC_HANDLE, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, Key unsafe.Pointer) CSSM_RETURN {
	if _cSSM_GenerateKey == nil {
		panic("Security: symbol CSSM_GenerateKey not loaded")
	}
	return _cSSM_GenerateKey(CCHandle, KeyUsage, KeyAttr, KeyLabel, CredAndAclEntry, Key)
}

var _cSSM_GenerateKeyP func(CCHandle CSSM_CC_HANDLE, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, Key unsafe.Pointer, Privilege CSSM_PRIVILEGE) CSSM_RETURN

// CSSM_GenerateKeyP.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GenerateKeyP
func CSSM_GenerateKeyP(CCHandle CSSM_CC_HANDLE, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, Key unsafe.Pointer, Privilege CSSM_PRIVILEGE) CSSM_RETURN {
	if _cSSM_GenerateKeyP == nil {
		panic("Security: symbol CSSM_GenerateKeyP not loaded")
	}
	return _cSSM_GenerateKeyP(CCHandle, KeyUsage, KeyAttr, KeyLabel, CredAndAclEntry, Key, Privilege)
}

var _cSSM_GenerateKeyPair func(CCHandle CSSM_CC_HANDLE, PublicKeyUsage uint32, PublicKeyAttr uint32, PublicKeyLabel unsafe.Pointer, PublicKey unsafe.Pointer, PrivateKeyUsage uint32, PrivateKeyAttr uint32, PrivateKeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, PrivateKey unsafe.Pointer) CSSM_RETURN

// CSSM_GenerateKeyPair.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GenerateKeyPair
func CSSM_GenerateKeyPair(CCHandle CSSM_CC_HANDLE, PublicKeyUsage uint32, PublicKeyAttr uint32, PublicKeyLabel unsafe.Pointer, PublicKey unsafe.Pointer, PrivateKeyUsage uint32, PrivateKeyAttr uint32, PrivateKeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, PrivateKey unsafe.Pointer) CSSM_RETURN {
	if _cSSM_GenerateKeyPair == nil {
		panic("Security: symbol CSSM_GenerateKeyPair not loaded")
	}
	return _cSSM_GenerateKeyPair(CCHandle, PublicKeyUsage, PublicKeyAttr, PublicKeyLabel, PublicKey, PrivateKeyUsage, PrivateKeyAttr, PrivateKeyLabel, CredAndAclEntry, PrivateKey)
}

var _cSSM_GenerateKeyPairP func(CCHandle CSSM_CC_HANDLE, PublicKeyUsage uint32, PublicKeyAttr uint32, PublicKeyLabel unsafe.Pointer, PublicKey unsafe.Pointer, PrivateKeyUsage uint32, PrivateKeyAttr uint32, PrivateKeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, PrivateKey unsafe.Pointer, Privilege CSSM_PRIVILEGE) CSSM_RETURN

// CSSM_GenerateKeyPairP.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GenerateKeyPairP
func CSSM_GenerateKeyPairP(CCHandle CSSM_CC_HANDLE, PublicKeyUsage uint32, PublicKeyAttr uint32, PublicKeyLabel unsafe.Pointer, PublicKey unsafe.Pointer, PrivateKeyUsage uint32, PrivateKeyAttr uint32, PrivateKeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, PrivateKey unsafe.Pointer, Privilege CSSM_PRIVILEGE) CSSM_RETURN {
	if _cSSM_GenerateKeyPairP == nil {
		panic("Security: symbol CSSM_GenerateKeyPairP not loaded")
	}
	return _cSSM_GenerateKeyPairP(CCHandle, PublicKeyUsage, PublicKeyAttr, PublicKeyLabel, PublicKey, PrivateKeyUsage, PrivateKeyAttr, PrivateKeyLabel, CredAndAclEntry, PrivateKey, Privilege)
}

var _cSSM_GenerateMac func(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, Mac unsafe.Pointer) CSSM_RETURN

// CSSM_GenerateMac.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GenerateMac
func CSSM_GenerateMac(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, Mac unsafe.Pointer) CSSM_RETURN {
	if _cSSM_GenerateMac == nil {
		panic("Security: symbol CSSM_GenerateMac not loaded")
	}
	return _cSSM_GenerateMac(CCHandle, DataBufs, DataBufCount, Mac)
}

var _cSSM_GenerateMacFinal func(CCHandle CSSM_CC_HANDLE, Mac unsafe.Pointer) CSSM_RETURN

// CSSM_GenerateMacFinal.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GenerateMacFinal
func CSSM_GenerateMacFinal(CCHandle CSSM_CC_HANDLE, Mac unsafe.Pointer) CSSM_RETURN {
	if _cSSM_GenerateMacFinal == nil {
		panic("Security: symbol CSSM_GenerateMacFinal not loaded")
	}
	return _cSSM_GenerateMacFinal(CCHandle, Mac)
}

var _cSSM_GenerateMacInit func(CCHandle CSSM_CC_HANDLE) CSSM_RETURN

// CSSM_GenerateMacInit.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GenerateMacInit
func CSSM_GenerateMacInit(CCHandle CSSM_CC_HANDLE) CSSM_RETURN {
	if _cSSM_GenerateMacInit == nil {
		panic("Security: symbol CSSM_GenerateMacInit not loaded")
	}
	return _cSSM_GenerateMacInit(CCHandle)
}

var _cSSM_GenerateMacUpdate func(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) CSSM_RETURN

// CSSM_GenerateMacUpdate.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GenerateMacUpdate
func CSSM_GenerateMacUpdate(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) CSSM_RETURN {
	if _cSSM_GenerateMacUpdate == nil {
		panic("Security: symbol CSSM_GenerateMacUpdate not loaded")
	}
	return _cSSM_GenerateMacUpdate(CCHandle, DataBufs, DataBufCount)
}

var _cSSM_GenerateRandom func(CCHandle CSSM_CC_HANDLE, RandomNumber unsafe.Pointer) CSSM_RETURN

// CSSM_GenerateRandom.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GenerateRandom
func CSSM_GenerateRandom(CCHandle CSSM_CC_HANDLE, RandomNumber unsafe.Pointer) CSSM_RETURN {
	if _cSSM_GenerateRandom == nil {
		panic("Security: symbol CSSM_GenerateRandom not loaded")
	}
	return _cSSM_GenerateRandom(CCHandle, RandomNumber)
}

var _cSSM_GetAPIMemoryFunctions func(AddInHandle CSSM_MODULE_HANDLE, AppMemoryFuncs unsafe.Pointer) CSSM_RETURN

// CSSM_GetAPIMemoryFunctions.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GetAPIMemoryFunctions
func CSSM_GetAPIMemoryFunctions(AddInHandle CSSM_MODULE_HANDLE, AppMemoryFuncs unsafe.Pointer) CSSM_RETURN {
	if _cSSM_GetAPIMemoryFunctions == nil {
		panic("Security: symbol CSSM_GetAPIMemoryFunctions not loaded")
	}
	return _cSSM_GetAPIMemoryFunctions(AddInHandle, AppMemoryFuncs)
}

var _cSSM_GetContext func(CCHandle CSSM_CC_HANDLE, Context unsafe.Pointer) CSSM_RETURN

// CSSM_GetContext.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GetContext
func CSSM_GetContext(CCHandle CSSM_CC_HANDLE, Context unsafe.Pointer) CSSM_RETURN {
	if _cSSM_GetContext == nil {
		panic("Security: symbol CSSM_GetContext not loaded")
	}
	return _cSSM_GetContext(CCHandle, Context)
}

var _cSSM_GetContextAttribute func(Context unsafe.Pointer, AttributeType uint32, ContextAttribute unsafe.Pointer) CSSM_RETURN

// CSSM_GetContextAttribute.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GetContextAttribute
func CSSM_GetContextAttribute(Context unsafe.Pointer, AttributeType uint32, ContextAttribute unsafe.Pointer) CSSM_RETURN {
	if _cSSM_GetContextAttribute == nil {
		panic("Security: symbol CSSM_GetContextAttribute not loaded")
	}
	return _cSSM_GetContextAttribute(Context, AttributeType, ContextAttribute)
}

var _cSSM_GetKeyAcl func(CSPHandle CSSM_CSP_HANDLE, Key unsafe.Pointer, SelectionTag unsafe.Pointer, NumberOfAclInfos *uint32, AclInfos unsafe.Pointer) CSSM_RETURN

// CSSM_GetKeyAcl.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GetKeyAcl
func CSSM_GetKeyAcl(CSPHandle CSSM_CSP_HANDLE, Key unsafe.Pointer, SelectionTag unsafe.Pointer, NumberOfAclInfos *uint32, AclInfos unsafe.Pointer) CSSM_RETURN {
	if _cSSM_GetKeyAcl == nil {
		panic("Security: symbol CSSM_GetKeyAcl not loaded")
	}
	return _cSSM_GetKeyAcl(CSPHandle, Key, SelectionTag, NumberOfAclInfos, AclInfos)
}

var _cSSM_GetKeyOwner func(CSPHandle CSSM_CSP_HANDLE, Key unsafe.Pointer, Owner unsafe.Pointer) CSSM_RETURN

// CSSM_GetKeyOwner.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GetKeyOwner
func CSSM_GetKeyOwner(CSPHandle CSSM_CSP_HANDLE, Key unsafe.Pointer, Owner unsafe.Pointer) CSSM_RETURN {
	if _cSSM_GetKeyOwner == nil {
		panic("Security: symbol CSSM_GetKeyOwner not loaded")
	}
	return _cSSM_GetKeyOwner(CSPHandle, Key, Owner)
}

var _cSSM_GetModuleGUIDFromHandle func(ModuleHandle CSSM_MODULE_HANDLE, ModuleGUID unsafe.Pointer) CSSM_RETURN

// CSSM_GetModuleGUIDFromHandle.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GetModuleGUIDFromHandle
func CSSM_GetModuleGUIDFromHandle(ModuleHandle CSSM_MODULE_HANDLE, ModuleGUID unsafe.Pointer) CSSM_RETURN {
	if _cSSM_GetModuleGUIDFromHandle == nil {
		panic("Security: symbol CSSM_GetModuleGUIDFromHandle not loaded")
	}
	return _cSSM_GetModuleGUIDFromHandle(ModuleHandle, ModuleGUID)
}

var _cSSM_GetPrivilege func(Privilege unsafe.Pointer) CSSM_RETURN

// CSSM_GetPrivilege.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GetPrivilege
func CSSM_GetPrivilege(Privilege unsafe.Pointer) CSSM_RETURN {
	if _cSSM_GetPrivilege == nil {
		panic("Security: symbol CSSM_GetPrivilege not loaded")
	}
	return _cSSM_GetPrivilege(Privilege)
}

var _cSSM_GetSubserviceUIDFromHandle func(ModuleHandle CSSM_MODULE_HANDLE, SubserviceUID unsafe.Pointer) CSSM_RETURN

// CSSM_GetSubserviceUIDFromHandle.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GetSubserviceUIDFromHandle
func CSSM_GetSubserviceUIDFromHandle(ModuleHandle CSSM_MODULE_HANDLE, SubserviceUID unsafe.Pointer) CSSM_RETURN {
	if _cSSM_GetSubserviceUIDFromHandle == nil {
		panic("Security: symbol CSSM_GetSubserviceUIDFromHandle not loaded")
	}
	return _cSSM_GetSubserviceUIDFromHandle(ModuleHandle, SubserviceUID)
}

var _cSSM_GetTimeValue func(CSPHandle CSSM_CSP_HANDLE, TimeAlgorithm CSSM_ALGORITHMS, TimeData unsafe.Pointer) CSSM_RETURN

// CSSM_GetTimeValue.
//
// See: https://developer.apple.com/documentation/Security/CSSM_GetTimeValue
func CSSM_GetTimeValue(CSPHandle CSSM_CSP_HANDLE, TimeAlgorithm CSSM_ALGORITHMS, TimeData unsafe.Pointer) CSSM_RETURN {
	if _cSSM_GetTimeValue == nil {
		panic("Security: symbol CSSM_GetTimeValue not loaded")
	}
	return _cSSM_GetTimeValue(CSPHandle, TimeAlgorithm, TimeData)
}

var _cSSM_Init func(Version unsafe.Pointer, Scope CSSM_PRIVILEGE_SCOPE, CallerGuid unsafe.Pointer, KeyHierarchy CSSM_KEY_HIERARCHY, PvcPolicy unsafe.Pointer, Reserved unsafe.Pointer) CSSM_RETURN

// CSSM_Init.
//
// See: https://developer.apple.com/documentation/Security/CSSM_Init
func CSSM_Init(Version unsafe.Pointer, Scope CSSM_PRIVILEGE_SCOPE, CallerGuid unsafe.Pointer, KeyHierarchy CSSM_KEY_HIERARCHY, PvcPolicy unsafe.Pointer, Reserved unsafe.Pointer) CSSM_RETURN {
	if _cSSM_Init == nil {
		panic("Security: symbol CSSM_Init not loaded")
	}
	return _cSSM_Init(Version, Scope, CallerGuid, KeyHierarchy, PvcPolicy, Reserved)
}

var _cSSM_Introduce func(ModuleID unsafe.Pointer, KeyHierarchy CSSM_KEY_HIERARCHY) CSSM_RETURN

// CSSM_Introduce.
//
// See: https://developer.apple.com/documentation/Security/CSSM_Introduce
func CSSM_Introduce(ModuleID unsafe.Pointer, KeyHierarchy CSSM_KEY_HIERARCHY) CSSM_RETURN {
	if _cSSM_Introduce == nil {
		panic("Security: symbol CSSM_Introduce not loaded")
	}
	return _cSSM_Introduce(ModuleID, KeyHierarchy)
}

var _cSSM_ListAttachedModuleManagers func(NumberOfModuleManagers *uint32, ModuleManagerGuids unsafe.Pointer) CSSM_RETURN

// CSSM_ListAttachedModuleManagers.
//
// See: https://developer.apple.com/documentation/Security/CSSM_ListAttachedModuleManagers
func CSSM_ListAttachedModuleManagers(NumberOfModuleManagers *uint32, ModuleManagerGuids unsafe.Pointer) CSSM_RETURN {
	if _cSSM_ListAttachedModuleManagers == nil {
		panic("Security: symbol CSSM_ListAttachedModuleManagers not loaded")
	}
	return _cSSM_ListAttachedModuleManagers(NumberOfModuleManagers, ModuleManagerGuids)
}

var _cSSM_ModuleAttach func(ModuleGuid unsafe.Pointer, Version unsafe.Pointer, MemoryFuncs unsafe.Pointer, SubserviceID uint32, SubServiceType CSSM_SERVICE_TYPE, AttachFlags CSSM_ATTACH_FLAGS, KeyHierarchy CSSM_KEY_HIERARCHY, FunctionTable unsafe.Pointer, NumFunctionTable uint32, Reserved unsafe.Pointer, NewModuleHandle CSSM_MODULE_HANDLE_PTR) CSSM_RETURN

// CSSM_ModuleAttach.
//
// See: https://developer.apple.com/documentation/Security/CSSM_ModuleAttach
func CSSM_ModuleAttach(ModuleGuid unsafe.Pointer, Version unsafe.Pointer, MemoryFuncs unsafe.Pointer, SubserviceID uint32, SubServiceType CSSM_SERVICE_TYPE, AttachFlags CSSM_ATTACH_FLAGS, KeyHierarchy CSSM_KEY_HIERARCHY, FunctionTable unsafe.Pointer, NumFunctionTable uint32, Reserved unsafe.Pointer, NewModuleHandle CSSM_MODULE_HANDLE_PTR) CSSM_RETURN {
	if _cSSM_ModuleAttach == nil {
		panic("Security: symbol CSSM_ModuleAttach not loaded")
	}
	return _cSSM_ModuleAttach(ModuleGuid, Version, MemoryFuncs, SubserviceID, SubServiceType, AttachFlags, KeyHierarchy, FunctionTable, NumFunctionTable, Reserved, NewModuleHandle)
}

var _cSSM_ModuleDetach func(ModuleHandle CSSM_MODULE_HANDLE) CSSM_RETURN

// CSSM_ModuleDetach.
//
// See: https://developer.apple.com/documentation/Security/CSSM_ModuleDetach
func CSSM_ModuleDetach(ModuleHandle CSSM_MODULE_HANDLE) CSSM_RETURN {
	if _cSSM_ModuleDetach == nil {
		panic("Security: symbol CSSM_ModuleDetach not loaded")
	}
	return _cSSM_ModuleDetach(ModuleHandle)
}

var _cSSM_ModuleLoad func(ModuleGuid unsafe.Pointer, KeyHierarchy CSSM_KEY_HIERARCHY, AppNotifyCallback unsafe.Pointer, AppNotifyCallbackCtx unsafe.Pointer) CSSM_RETURN

// CSSM_ModuleLoad.
//
// See: https://developer.apple.com/documentation/Security/CSSM_ModuleLoad
func CSSM_ModuleLoad(ModuleGuid unsafe.Pointer, KeyHierarchy CSSM_KEY_HIERARCHY, AppNotifyCallback unsafe.Pointer, AppNotifyCallbackCtx unsafe.Pointer) CSSM_RETURN {
	if _cSSM_ModuleLoad == nil {
		panic("Security: symbol CSSM_ModuleLoad not loaded")
	}
	return _cSSM_ModuleLoad(ModuleGuid, KeyHierarchy, AppNotifyCallback, AppNotifyCallbackCtx)
}

var _cSSM_ModuleUnload func(ModuleGuid unsafe.Pointer, AppNotifyCallback unsafe.Pointer, AppNotifyCallbackCtx unsafe.Pointer) CSSM_RETURN

// CSSM_ModuleUnload.
//
// See: https://developer.apple.com/documentation/Security/CSSM_ModuleUnload
func CSSM_ModuleUnload(ModuleGuid unsafe.Pointer, AppNotifyCallback unsafe.Pointer, AppNotifyCallbackCtx unsafe.Pointer) CSSM_RETURN {
	if _cSSM_ModuleUnload == nil {
		panic("Security: symbol CSSM_ModuleUnload not loaded")
	}
	return _cSSM_ModuleUnload(ModuleGuid, AppNotifyCallback, AppNotifyCallbackCtx)
}

var _cSSM_QueryKeySizeInBits func(CSPHandle CSSM_CSP_HANDLE, CCHandle CSSM_CC_HANDLE, Key unsafe.Pointer, KeySize unsafe.Pointer) CSSM_RETURN

// CSSM_QueryKeySizeInBits.
//
// See: https://developer.apple.com/documentation/Security/CSSM_QueryKeySizeInBits
func CSSM_QueryKeySizeInBits(CSPHandle CSSM_CSP_HANDLE, CCHandle CSSM_CC_HANDLE, Key unsafe.Pointer, KeySize unsafe.Pointer) CSSM_RETURN {
	if _cSSM_QueryKeySizeInBits == nil {
		panic("Security: symbol CSSM_QueryKeySizeInBits not loaded")
	}
	return _cSSM_QueryKeySizeInBits(CSPHandle, CCHandle, Key, KeySize)
}

var _cSSM_QuerySize func(CCHandle CSSM_CC_HANDLE, Encrypt CSSM_BOOL, QuerySizeCount uint32, DataBlockSizes unsafe.Pointer) CSSM_RETURN

// CSSM_QuerySize.
//
// See: https://developer.apple.com/documentation/Security/CSSM_QuerySize
func CSSM_QuerySize(CCHandle CSSM_CC_HANDLE, Encrypt CSSM_BOOL, QuerySizeCount uint32, DataBlockSizes unsafe.Pointer) CSSM_RETURN {
	if _cSSM_QuerySize == nil {
		panic("Security: symbol CSSM_QuerySize not loaded")
	}
	return _cSSM_QuerySize(CCHandle, Encrypt, QuerySizeCount, DataBlockSizes)
}

var _cSSM_RetrieveCounter func(CSPHandle CSSM_CSP_HANDLE, Counter unsafe.Pointer) CSSM_RETURN

// CSSM_RetrieveCounter.
//
// See: https://developer.apple.com/documentation/Security/CSSM_RetrieveCounter
func CSSM_RetrieveCounter(CSPHandle CSSM_CSP_HANDLE, Counter unsafe.Pointer) CSSM_RETURN {
	if _cSSM_RetrieveCounter == nil {
		panic("Security: symbol CSSM_RetrieveCounter not loaded")
	}
	return _cSSM_RetrieveCounter(CSPHandle, Counter)
}

var _cSSM_RetrieveUniqueId func(CSPHandle CSSM_CSP_HANDLE, UniqueID unsafe.Pointer) CSSM_RETURN

// CSSM_RetrieveUniqueId.
//
// See: https://developer.apple.com/documentation/Security/CSSM_RetrieveUniqueId
func CSSM_RetrieveUniqueId(CSPHandle CSSM_CSP_HANDLE, UniqueID unsafe.Pointer) CSSM_RETURN {
	if _cSSM_RetrieveUniqueId == nil {
		panic("Security: symbol CSSM_RetrieveUniqueId not loaded")
	}
	return _cSSM_RetrieveUniqueId(CSPHandle, UniqueID)
}

var _cSSM_SetContext func(CCHandle CSSM_CC_HANDLE, Context unsafe.Pointer) CSSM_RETURN

// CSSM_SetContext.
//
// See: https://developer.apple.com/documentation/Security/CSSM_SetContext
func CSSM_SetContext(CCHandle CSSM_CC_HANDLE, Context unsafe.Pointer) CSSM_RETURN {
	if _cSSM_SetContext == nil {
		panic("Security: symbol CSSM_SetContext not loaded")
	}
	return _cSSM_SetContext(CCHandle, Context)
}

var _cSSM_SetPrivilege func(Privilege CSSM_PRIVILEGE) CSSM_RETURN

// CSSM_SetPrivilege.
//
// See: https://developer.apple.com/documentation/Security/CSSM_SetPrivilege
func CSSM_SetPrivilege(Privilege CSSM_PRIVILEGE) CSSM_RETURN {
	if _cSSM_SetPrivilege == nil {
		panic("Security: symbol CSSM_SetPrivilege not loaded")
	}
	return _cSSM_SetPrivilege(Privilege)
}

var _cSSM_SignData func(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, DigestAlgorithm CSSM_ALGORITHMS, Signature unsafe.Pointer) CSSM_RETURN

// CSSM_SignData.
//
// See: https://developer.apple.com/documentation/Security/CSSM_SignData
func CSSM_SignData(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, DigestAlgorithm CSSM_ALGORITHMS, Signature unsafe.Pointer) CSSM_RETURN {
	if _cSSM_SignData == nil {
		panic("Security: symbol CSSM_SignData not loaded")
	}
	return _cSSM_SignData(CCHandle, DataBufs, DataBufCount, DigestAlgorithm, Signature)
}

var _cSSM_SignDataFinal func(CCHandle CSSM_CC_HANDLE, Signature unsafe.Pointer) CSSM_RETURN

// CSSM_SignDataFinal.
//
// See: https://developer.apple.com/documentation/Security/CSSM_SignDataFinal
func CSSM_SignDataFinal(CCHandle CSSM_CC_HANDLE, Signature unsafe.Pointer) CSSM_RETURN {
	if _cSSM_SignDataFinal == nil {
		panic("Security: symbol CSSM_SignDataFinal not loaded")
	}
	return _cSSM_SignDataFinal(CCHandle, Signature)
}

var _cSSM_SignDataInit func(CCHandle CSSM_CC_HANDLE) CSSM_RETURN

// CSSM_SignDataInit.
//
// See: https://developer.apple.com/documentation/Security/CSSM_SignDataInit
func CSSM_SignDataInit(CCHandle CSSM_CC_HANDLE) CSSM_RETURN {
	if _cSSM_SignDataInit == nil {
		panic("Security: symbol CSSM_SignDataInit not loaded")
	}
	return _cSSM_SignDataInit(CCHandle)
}

var _cSSM_SignDataUpdate func(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) CSSM_RETURN

// CSSM_SignDataUpdate.
//
// See: https://developer.apple.com/documentation/Security/CSSM_SignDataUpdate
func CSSM_SignDataUpdate(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) CSSM_RETURN {
	if _cSSM_SignDataUpdate == nil {
		panic("Security: symbol CSSM_SignDataUpdate not loaded")
	}
	return _cSSM_SignDataUpdate(CCHandle, DataBufs, DataBufCount)
}

var _cSSM_TP_ApplyCrlToDb func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, CrlToBeApplied unsafe.Pointer, SignerCertGroup unsafe.Pointer, ApplyCrlVerifyContext unsafe.Pointer, ApplyCrlVerifyResult unsafe.Pointer) CSSM_RETURN

// CSSM_TP_ApplyCrlToDb.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_ApplyCrlToDb
func CSSM_TP_ApplyCrlToDb(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, CrlToBeApplied unsafe.Pointer, SignerCertGroup unsafe.Pointer, ApplyCrlVerifyContext unsafe.Pointer, ApplyCrlVerifyResult unsafe.Pointer) CSSM_RETURN {
	if _cSSM_TP_ApplyCrlToDb == nil {
		panic("Security: symbol CSSM_TP_ApplyCrlToDb not loaded")
	}
	return _cSSM_TP_ApplyCrlToDb(TPHandle, CLHandle, CSPHandle, CrlToBeApplied, SignerCertGroup, ApplyCrlVerifyContext, ApplyCrlVerifyResult)
}

var _cSSM_TP_CertCreateTemplate func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CertFields unsafe.Pointer, CertTemplate unsafe.Pointer) CSSM_RETURN

// CSSM_TP_CertCreateTemplate.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CertCreateTemplate
func CSSM_TP_CertCreateTemplate(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CertFields unsafe.Pointer, CertTemplate unsafe.Pointer) CSSM_RETURN {
	if _cSSM_TP_CertCreateTemplate == nil {
		panic("Security: symbol CSSM_TP_CertCreateTemplate not loaded")
	}
	return _cSSM_TP_CertCreateTemplate(TPHandle, CLHandle, NumberOfFields, CertFields, CertTemplate)
}

var _cSSM_TP_CertGetAllTemplateFields func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CertTemplate unsafe.Pointer, NumberOfFields *uint32, CertFields unsafe.Pointer) CSSM_RETURN

// CSSM_TP_CertGetAllTemplateFields.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CertGetAllTemplateFields
func CSSM_TP_CertGetAllTemplateFields(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CertTemplate unsafe.Pointer, NumberOfFields *uint32, CertFields unsafe.Pointer) CSSM_RETURN {
	if _cSSM_TP_CertGetAllTemplateFields == nil {
		panic("Security: symbol CSSM_TP_CertGetAllTemplateFields not loaded")
	}
	return _cSSM_TP_CertGetAllTemplateFields(TPHandle, CLHandle, CertTemplate, NumberOfFields, CertFields)
}

var _cSSM_TP_CertGroupConstruct func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, DBList unsafe.Pointer, ConstructParams unsafe.Pointer, CertGroupFrag unsafe.Pointer, CertGroup unsafe.Pointer) CSSM_RETURN

// CSSM_TP_CertGroupConstruct.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CertGroupConstruct
func CSSM_TP_CertGroupConstruct(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, DBList unsafe.Pointer, ConstructParams unsafe.Pointer, CertGroupFrag unsafe.Pointer, CertGroup unsafe.Pointer) CSSM_RETURN {
	if _cSSM_TP_CertGroupConstruct == nil {
		panic("Security: symbol CSSM_TP_CertGroupConstruct not loaded")
	}
	return _cSSM_TP_CertGroupConstruct(TPHandle, CLHandle, CSPHandle, DBList, ConstructParams, CertGroupFrag, CertGroup)
}

var _cSSM_TP_CertGroupPrune func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, DBList unsafe.Pointer, OrderedCertGroup unsafe.Pointer, PrunedCertGroup unsafe.Pointer) CSSM_RETURN

// CSSM_TP_CertGroupPrune.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CertGroupPrune
func CSSM_TP_CertGroupPrune(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, DBList unsafe.Pointer, OrderedCertGroup unsafe.Pointer, PrunedCertGroup unsafe.Pointer) CSSM_RETURN {
	if _cSSM_TP_CertGroupPrune == nil {
		panic("Security: symbol CSSM_TP_CertGroupPrune not loaded")
	}
	return _cSSM_TP_CertGroupPrune(TPHandle, CLHandle, DBList, OrderedCertGroup, PrunedCertGroup)
}

var _cSSM_TP_CertGroupToTupleGroup func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CertGroup unsafe.Pointer, TupleGroup unsafe.Pointer) CSSM_RETURN

// CSSM_TP_CertGroupToTupleGroup.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CertGroupToTupleGroup
func CSSM_TP_CertGroupToTupleGroup(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CertGroup unsafe.Pointer, TupleGroup unsafe.Pointer) CSSM_RETURN {
	if _cSSM_TP_CertGroupToTupleGroup == nil {
		panic("Security: symbol CSSM_TP_CertGroupToTupleGroup not loaded")
	}
	return _cSSM_TP_CertGroupToTupleGroup(TPHandle, CLHandle, CertGroup, TupleGroup)
}

var _cSSM_TP_CertGroupVerify func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, CertGroupToBeVerified unsafe.Pointer, VerifyContext unsafe.Pointer, VerifyContextResult unsafe.Pointer) CSSM_RETURN

// CSSM_TP_CertGroupVerify.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CertGroupVerify
func CSSM_TP_CertGroupVerify(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, CertGroupToBeVerified unsafe.Pointer, VerifyContext unsafe.Pointer, VerifyContextResult unsafe.Pointer) CSSM_RETURN {
	if _cSSM_TP_CertGroupVerify == nil {
		panic("Security: symbol CSSM_TP_CertGroupVerify not loaded")
	}
	return _cSSM_TP_CertGroupVerify(TPHandle, CLHandle, CSPHandle, CertGroupToBeVerified, VerifyContext, VerifyContextResult)
}

var _cSSM_TP_CertReclaimAbort func(TPHandle CSSM_TP_HANDLE, KeyCacheHandle CSSM_LONG_HANDLE) CSSM_RETURN

// CSSM_TP_CertReclaimAbort.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CertReclaimAbort
func CSSM_TP_CertReclaimAbort(TPHandle CSSM_TP_HANDLE, KeyCacheHandle CSSM_LONG_HANDLE) CSSM_RETURN {
	if _cSSM_TP_CertReclaimAbort == nil {
		panic("Security: symbol CSSM_TP_CertReclaimAbort not loaded")
	}
	return _cSSM_TP_CertReclaimAbort(TPHandle, KeyCacheHandle)
}

var _cSSM_TP_CertReclaimKey func(TPHandle CSSM_TP_HANDLE, CertGroup unsafe.Pointer, CertIndex uint32, KeyCacheHandle CSSM_LONG_HANDLE, CSPHandle CSSM_CSP_HANDLE, CredAndAclEntry unsafe.Pointer) CSSM_RETURN

// CSSM_TP_CertReclaimKey.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CertReclaimKey
func CSSM_TP_CertReclaimKey(TPHandle CSSM_TP_HANDLE, CertGroup unsafe.Pointer, CertIndex uint32, KeyCacheHandle CSSM_LONG_HANDLE, CSPHandle CSSM_CSP_HANDLE, CredAndAclEntry unsafe.Pointer) CSSM_RETURN {
	if _cSSM_TP_CertReclaimKey == nil {
		panic("Security: symbol CSSM_TP_CertReclaimKey not loaded")
	}
	return _cSSM_TP_CertReclaimKey(TPHandle, CertGroup, CertIndex, KeyCacheHandle, CSPHandle, CredAndAclEntry)
}

var _cSSM_TP_CertRemoveFromCrlTemplate func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, OldCrlTemplate unsafe.Pointer, CertGroupToBeRemoved unsafe.Pointer, RevokerCertGroup unsafe.Pointer, RevokerVerifyContext unsafe.Pointer, RevokerVerifyResult unsafe.Pointer, NewCrlTemplate unsafe.Pointer) CSSM_RETURN

// CSSM_TP_CertRemoveFromCrlTemplate.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CertRemoveFromCrlTemplate
func CSSM_TP_CertRemoveFromCrlTemplate(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, OldCrlTemplate unsafe.Pointer, CertGroupToBeRemoved unsafe.Pointer, RevokerCertGroup unsafe.Pointer, RevokerVerifyContext unsafe.Pointer, RevokerVerifyResult unsafe.Pointer, NewCrlTemplate unsafe.Pointer) CSSM_RETURN {
	if _cSSM_TP_CertRemoveFromCrlTemplate == nil {
		panic("Security: symbol CSSM_TP_CertRemoveFromCrlTemplate not loaded")
	}
	return _cSSM_TP_CertRemoveFromCrlTemplate(TPHandle, CLHandle, CSPHandle, OldCrlTemplate, CertGroupToBeRemoved, RevokerCertGroup, RevokerVerifyContext, RevokerVerifyResult, NewCrlTemplate)
}

var _cSSM_TP_CertRevoke func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, OldCrlTemplate unsafe.Pointer, CertGroupToBeRevoked unsafe.Pointer, RevokerCertGroup unsafe.Pointer, RevokerVerifyContext unsafe.Pointer, RevokerVerifyResult unsafe.Pointer, Reason CSSM_TP_CERTCHANGE_REASON, NewCrlTemplate unsafe.Pointer) CSSM_RETURN

// CSSM_TP_CertRevoke.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CertRevoke
func CSSM_TP_CertRevoke(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, OldCrlTemplate unsafe.Pointer, CertGroupToBeRevoked unsafe.Pointer, RevokerCertGroup unsafe.Pointer, RevokerVerifyContext unsafe.Pointer, RevokerVerifyResult unsafe.Pointer, Reason CSSM_TP_CERTCHANGE_REASON, NewCrlTemplate unsafe.Pointer) CSSM_RETURN {
	if _cSSM_TP_CertRevoke == nil {
		panic("Security: symbol CSSM_TP_CertRevoke not loaded")
	}
	return _cSSM_TP_CertRevoke(TPHandle, CLHandle, CSPHandle, OldCrlTemplate, CertGroupToBeRevoked, RevokerCertGroup, RevokerVerifyContext, RevokerVerifyResult, Reason, NewCrlTemplate)
}

var _cSSM_TP_CertSign func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertTemplateToBeSigned unsafe.Pointer, SignerCertGroup unsafe.Pointer, SignerVerifyContext unsafe.Pointer, SignerVerifyResult unsafe.Pointer, SignedCert unsafe.Pointer) CSSM_RETURN

// CSSM_TP_CertSign.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CertSign
func CSSM_TP_CertSign(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CertTemplateToBeSigned unsafe.Pointer, SignerCertGroup unsafe.Pointer, SignerVerifyContext unsafe.Pointer, SignerVerifyResult unsafe.Pointer, SignedCert unsafe.Pointer) CSSM_RETURN {
	if _cSSM_TP_CertSign == nil {
		panic("Security: symbol CSSM_TP_CertSign not loaded")
	}
	return _cSSM_TP_CertSign(TPHandle, CLHandle, CCHandle, CertTemplateToBeSigned, SignerCertGroup, SignerVerifyContext, SignerVerifyResult, SignedCert)
}

var _cSSM_TP_ConfirmCredResult func(TPHandle CSSM_TP_HANDLE, ReferenceIdentifier unsafe.Pointer, CallerAuthCredentials unsafe.Pointer, Responses unsafe.Pointer, PreferredAuthority unsafe.Pointer) CSSM_RETURN

// CSSM_TP_ConfirmCredResult.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_ConfirmCredResult
func CSSM_TP_ConfirmCredResult(TPHandle CSSM_TP_HANDLE, ReferenceIdentifier unsafe.Pointer, CallerAuthCredentials unsafe.Pointer, Responses unsafe.Pointer, PreferredAuthority unsafe.Pointer) CSSM_RETURN {
	if _cSSM_TP_ConfirmCredResult == nil {
		panic("Security: symbol CSSM_TP_ConfirmCredResult not loaded")
	}
	return _cSSM_TP_ConfirmCredResult(TPHandle, ReferenceIdentifier, CallerAuthCredentials, Responses, PreferredAuthority)
}

var _cSSM_TP_CrlCreateTemplate func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CrlFields unsafe.Pointer, NewCrlTemplate unsafe.Pointer) CSSM_RETURN

// CSSM_TP_CrlCreateTemplate.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CrlCreateTemplate
func CSSM_TP_CrlCreateTemplate(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, NumberOfFields uint32, CrlFields unsafe.Pointer, NewCrlTemplate unsafe.Pointer) CSSM_RETURN {
	if _cSSM_TP_CrlCreateTemplate == nil {
		panic("Security: symbol CSSM_TP_CrlCreateTemplate not loaded")
	}
	return _cSSM_TP_CrlCreateTemplate(TPHandle, CLHandle, NumberOfFields, CrlFields, NewCrlTemplate)
}

var _cSSM_TP_CrlSign func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CrlToBeSigned unsafe.Pointer, SignerCertGroup unsafe.Pointer, SignerVerifyContext unsafe.Pointer, SignerVerifyResult unsafe.Pointer, SignedCrl unsafe.Pointer) CSSM_RETURN

// CSSM_TP_CrlSign.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CrlSign
func CSSM_TP_CrlSign(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, CrlToBeSigned unsafe.Pointer, SignerCertGroup unsafe.Pointer, SignerVerifyContext unsafe.Pointer, SignerVerifyResult unsafe.Pointer, SignedCrl unsafe.Pointer) CSSM_RETURN {
	if _cSSM_TP_CrlSign == nil {
		panic("Security: symbol CSSM_TP_CrlSign not loaded")
	}
	return _cSSM_TP_CrlSign(TPHandle, CLHandle, CCHandle, CrlToBeSigned, SignerCertGroup, SignerVerifyContext, SignerVerifyResult, SignedCrl)
}

var _cSSM_TP_CrlVerify func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, CrlToBeVerified unsafe.Pointer, SignerCertGroup unsafe.Pointer, VerifyContext unsafe.Pointer, RevokerVerifyResult unsafe.Pointer) CSSM_RETURN

// CSSM_TP_CrlVerify.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_CrlVerify
func CSSM_TP_CrlVerify(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CSPHandle CSSM_CSP_HANDLE, CrlToBeVerified unsafe.Pointer, SignerCertGroup unsafe.Pointer, VerifyContext unsafe.Pointer, RevokerVerifyResult unsafe.Pointer) CSSM_RETURN {
	if _cSSM_TP_CrlVerify == nil {
		panic("Security: symbol CSSM_TP_CrlVerify not loaded")
	}
	return _cSSM_TP_CrlVerify(TPHandle, CLHandle, CSPHandle, CrlToBeVerified, SignerCertGroup, VerifyContext, RevokerVerifyResult)
}

var _cSSM_TP_FormRequest func(TPHandle CSSM_TP_HANDLE, PreferredAuthority unsafe.Pointer, FormType CSSM_TP_FORM_TYPE, BlankForm unsafe.Pointer) CSSM_RETURN

// CSSM_TP_FormRequest.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_FormRequest
func CSSM_TP_FormRequest(TPHandle CSSM_TP_HANDLE, PreferredAuthority unsafe.Pointer, FormType CSSM_TP_FORM_TYPE, BlankForm unsafe.Pointer) CSSM_RETURN {
	if _cSSM_TP_FormRequest == nil {
		panic("Security: symbol CSSM_TP_FormRequest not loaded")
	}
	return _cSSM_TP_FormRequest(TPHandle, PreferredAuthority, FormType, BlankForm)
}

var _cSSM_TP_FormSubmit func(TPHandle CSSM_TP_HANDLE, FormType CSSM_TP_FORM_TYPE, Form unsafe.Pointer, ClearanceAuthority unsafe.Pointer, RepresentedAuthority unsafe.Pointer, Credentials unsafe.Pointer) CSSM_RETURN

// CSSM_TP_FormSubmit.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_FormSubmit
func CSSM_TP_FormSubmit(TPHandle CSSM_TP_HANDLE, FormType CSSM_TP_FORM_TYPE, Form unsafe.Pointer, ClearanceAuthority unsafe.Pointer, RepresentedAuthority unsafe.Pointer, Credentials unsafe.Pointer) CSSM_RETURN {
	if _cSSM_TP_FormSubmit == nil {
		panic("Security: symbol CSSM_TP_FormSubmit not loaded")
	}
	return _cSSM_TP_FormSubmit(TPHandle, FormType, Form, ClearanceAuthority, RepresentedAuthority, Credentials)
}

var _cSSM_TP_PassThrough func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, DBList unsafe.Pointer, PassThroughId uint32, InputParams unsafe.Pointer, OutputParams unsafe.Pointer) CSSM_RETURN

// CSSM_TP_PassThrough.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_PassThrough
func CSSM_TP_PassThrough(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, CCHandle CSSM_CC_HANDLE, DBList unsafe.Pointer, PassThroughId uint32, InputParams unsafe.Pointer, OutputParams unsafe.Pointer) CSSM_RETURN {
	if _cSSM_TP_PassThrough == nil {
		panic("Security: symbol CSSM_TP_PassThrough not loaded")
	}
	return _cSSM_TP_PassThrough(TPHandle, CLHandle, CCHandle, DBList, PassThroughId, InputParams, OutputParams)
}

var _cSSM_TP_ReceiveConfirmation func(TPHandle CSSM_TP_HANDLE, ReferenceIdentifier unsafe.Pointer, Responses unsafe.Pointer, ElapsedTime *int32) CSSM_RETURN

// CSSM_TP_ReceiveConfirmation.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_ReceiveConfirmation
func CSSM_TP_ReceiveConfirmation(TPHandle CSSM_TP_HANDLE, ReferenceIdentifier unsafe.Pointer, Responses unsafe.Pointer, ElapsedTime *int32) CSSM_RETURN {
	if _cSSM_TP_ReceiveConfirmation == nil {
		panic("Security: symbol CSSM_TP_ReceiveConfirmation not loaded")
	}
	return _cSSM_TP_ReceiveConfirmation(TPHandle, ReferenceIdentifier, Responses, ElapsedTime)
}

var _cSSM_TP_RetrieveCredResult func(TPHandle CSSM_TP_HANDLE, ReferenceIdentifier unsafe.Pointer, CallerAuthCredentials unsafe.Pointer, EstimatedTime *int32, ConfirmationRequired unsafe.Pointer, RetrieveOutput unsafe.Pointer) CSSM_RETURN

// CSSM_TP_RetrieveCredResult.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_RetrieveCredResult
func CSSM_TP_RetrieveCredResult(TPHandle CSSM_TP_HANDLE, ReferenceIdentifier unsafe.Pointer, CallerAuthCredentials unsafe.Pointer, EstimatedTime *int32, ConfirmationRequired unsafe.Pointer, RetrieveOutput unsafe.Pointer) CSSM_RETURN {
	if _cSSM_TP_RetrieveCredResult == nil {
		panic("Security: symbol CSSM_TP_RetrieveCredResult not loaded")
	}
	return _cSSM_TP_RetrieveCredResult(TPHandle, ReferenceIdentifier, CallerAuthCredentials, EstimatedTime, ConfirmationRequired, RetrieveOutput)
}

var _cSSM_TP_SubmitCredRequest func(TPHandle CSSM_TP_HANDLE, PreferredAuthority unsafe.Pointer, RequestType CSSM_TP_AUTHORITY_REQUEST_TYPE, RequestInput unsafe.Pointer, CallerAuthContext unsafe.Pointer, EstimatedTime *int32, ReferenceIdentifier unsafe.Pointer) CSSM_RETURN

// CSSM_TP_SubmitCredRequest.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_SubmitCredRequest
func CSSM_TP_SubmitCredRequest(TPHandle CSSM_TP_HANDLE, PreferredAuthority unsafe.Pointer, RequestType CSSM_TP_AUTHORITY_REQUEST_TYPE, RequestInput unsafe.Pointer, CallerAuthContext unsafe.Pointer, EstimatedTime *int32, ReferenceIdentifier unsafe.Pointer) CSSM_RETURN {
	if _cSSM_TP_SubmitCredRequest == nil {
		panic("Security: symbol CSSM_TP_SubmitCredRequest not loaded")
	}
	return _cSSM_TP_SubmitCredRequest(TPHandle, PreferredAuthority, RequestType, RequestInput, CallerAuthContext, EstimatedTime, ReferenceIdentifier)
}

var _cSSM_TP_TupleGroupToCertGroup func(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, TupleGroup unsafe.Pointer, CertTemplates unsafe.Pointer) CSSM_RETURN

// CSSM_TP_TupleGroupToCertGroup.
//
// See: https://developer.apple.com/documentation/Security/CSSM_TP_TupleGroupToCertGroup
func CSSM_TP_TupleGroupToCertGroup(TPHandle CSSM_TP_HANDLE, CLHandle CSSM_CL_HANDLE, TupleGroup unsafe.Pointer, CertTemplates unsafe.Pointer) CSSM_RETURN {
	if _cSSM_TP_TupleGroupToCertGroup == nil {
		panic("Security: symbol CSSM_TP_TupleGroupToCertGroup not loaded")
	}
	return _cSSM_TP_TupleGroupToCertGroup(TPHandle, CLHandle, TupleGroup, CertTemplates)
}

var _cSSM_Terminate func() CSSM_RETURN

// CSSM_Terminate.
//
// See: https://developer.apple.com/documentation/Security/CSSM_Terminate
func CSSM_Terminate() CSSM_RETURN {
	if _cSSM_Terminate == nil {
		panic("Security: symbol CSSM_Terminate not loaded")
	}
	return _cSSM_Terminate()
}

var _cSSM_Unintroduce func(ModuleID unsafe.Pointer) CSSM_RETURN

// CSSM_Unintroduce.
//
// See: https://developer.apple.com/documentation/Security/CSSM_Unintroduce
func CSSM_Unintroduce(ModuleID unsafe.Pointer) CSSM_RETURN {
	if _cSSM_Unintroduce == nil {
		panic("Security: symbol CSSM_Unintroduce not loaded")
	}
	return _cSSM_Unintroduce(ModuleID)
}

var _cSSM_UnwrapKey func(CCHandle CSSM_CC_HANDLE, PublicKey unsafe.Pointer, WrappedKey unsafe.Pointer, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, UnwrappedKey unsafe.Pointer, DescriptiveData unsafe.Pointer) CSSM_RETURN

// CSSM_UnwrapKey.
//
// See: https://developer.apple.com/documentation/Security/CSSM_UnwrapKey
func CSSM_UnwrapKey(CCHandle CSSM_CC_HANDLE, PublicKey unsafe.Pointer, WrappedKey unsafe.Pointer, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, UnwrappedKey unsafe.Pointer, DescriptiveData unsafe.Pointer) CSSM_RETURN {
	if _cSSM_UnwrapKey == nil {
		panic("Security: symbol CSSM_UnwrapKey not loaded")
	}
	return _cSSM_UnwrapKey(CCHandle, PublicKey, WrappedKey, KeyUsage, KeyAttr, KeyLabel, CredAndAclEntry, UnwrappedKey, DescriptiveData)
}

var _cSSM_UnwrapKeyP func(CCHandle CSSM_CC_HANDLE, PublicKey unsafe.Pointer, WrappedKey unsafe.Pointer, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, UnwrappedKey unsafe.Pointer, DescriptiveData unsafe.Pointer, Privilege CSSM_PRIVILEGE) CSSM_RETURN

// CSSM_UnwrapKeyP.
//
// See: https://developer.apple.com/documentation/Security/CSSM_UnwrapKeyP
func CSSM_UnwrapKeyP(CCHandle CSSM_CC_HANDLE, PublicKey unsafe.Pointer, WrappedKey unsafe.Pointer, KeyUsage uint32, KeyAttr uint32, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, UnwrappedKey unsafe.Pointer, DescriptiveData unsafe.Pointer, Privilege CSSM_PRIVILEGE) CSSM_RETURN {
	if _cSSM_UnwrapKeyP == nil {
		panic("Security: symbol CSSM_UnwrapKeyP not loaded")
	}
	return _cSSM_UnwrapKeyP(CCHandle, PublicKey, WrappedKey, KeyUsage, KeyAttr, KeyLabel, CredAndAclEntry, UnwrappedKey, DescriptiveData, Privilege)
}

var _cSSM_UpdateContextAttributes func(CCHandle CSSM_CC_HANDLE, NumberOfAttributes uint32, ContextAttributes unsafe.Pointer) CSSM_RETURN

// CSSM_UpdateContextAttributes.
//
// See: https://developer.apple.com/documentation/Security/CSSM_UpdateContextAttributes
func CSSM_UpdateContextAttributes(CCHandle CSSM_CC_HANDLE, NumberOfAttributes uint32, ContextAttributes unsafe.Pointer) CSSM_RETURN {
	if _cSSM_UpdateContextAttributes == nil {
		panic("Security: symbol CSSM_UpdateContextAttributes not loaded")
	}
	return _cSSM_UpdateContextAttributes(CCHandle, NumberOfAttributes, ContextAttributes)
}

var _cSSM_VerifyData func(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, DigestAlgorithm CSSM_ALGORITHMS, Signature unsafe.Pointer) CSSM_RETURN

// CSSM_VerifyData.
//
// See: https://developer.apple.com/documentation/Security/CSSM_VerifyData
func CSSM_VerifyData(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, DigestAlgorithm CSSM_ALGORITHMS, Signature unsafe.Pointer) CSSM_RETURN {
	if _cSSM_VerifyData == nil {
		panic("Security: symbol CSSM_VerifyData not loaded")
	}
	return _cSSM_VerifyData(CCHandle, DataBufs, DataBufCount, DigestAlgorithm, Signature)
}

var _cSSM_VerifyDataFinal func(CCHandle CSSM_CC_HANDLE, Signature unsafe.Pointer) CSSM_RETURN

// CSSM_VerifyDataFinal.
//
// See: https://developer.apple.com/documentation/Security/CSSM_VerifyDataFinal
func CSSM_VerifyDataFinal(CCHandle CSSM_CC_HANDLE, Signature unsafe.Pointer) CSSM_RETURN {
	if _cSSM_VerifyDataFinal == nil {
		panic("Security: symbol CSSM_VerifyDataFinal not loaded")
	}
	return _cSSM_VerifyDataFinal(CCHandle, Signature)
}

var _cSSM_VerifyDataInit func(CCHandle CSSM_CC_HANDLE) CSSM_RETURN

// CSSM_VerifyDataInit.
//
// See: https://developer.apple.com/documentation/Security/CSSM_VerifyDataInit
func CSSM_VerifyDataInit(CCHandle CSSM_CC_HANDLE) CSSM_RETURN {
	if _cSSM_VerifyDataInit == nil {
		panic("Security: symbol CSSM_VerifyDataInit not loaded")
	}
	return _cSSM_VerifyDataInit(CCHandle)
}

var _cSSM_VerifyDataUpdate func(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) CSSM_RETURN

// CSSM_VerifyDataUpdate.
//
// See: https://developer.apple.com/documentation/Security/CSSM_VerifyDataUpdate
func CSSM_VerifyDataUpdate(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) CSSM_RETURN {
	if _cSSM_VerifyDataUpdate == nil {
		panic("Security: symbol CSSM_VerifyDataUpdate not loaded")
	}
	return _cSSM_VerifyDataUpdate(CCHandle, DataBufs, DataBufCount)
}

var _cSSM_VerifyDevice func(CSPHandle CSSM_CSP_HANDLE, DeviceCert unsafe.Pointer) CSSM_RETURN

// CSSM_VerifyDevice.
//
// See: https://developer.apple.com/documentation/Security/CSSM_VerifyDevice
func CSSM_VerifyDevice(CSPHandle CSSM_CSP_HANDLE, DeviceCert unsafe.Pointer) CSSM_RETURN {
	if _cSSM_VerifyDevice == nil {
		panic("Security: symbol CSSM_VerifyDevice not loaded")
	}
	return _cSSM_VerifyDevice(CSPHandle, DeviceCert)
}

var _cSSM_VerifyMac func(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, Mac unsafe.Pointer) CSSM_RETURN

// CSSM_VerifyMac.
//
// See: https://developer.apple.com/documentation/Security/CSSM_VerifyMac
func CSSM_VerifyMac(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32, Mac unsafe.Pointer) CSSM_RETURN {
	if _cSSM_VerifyMac == nil {
		panic("Security: symbol CSSM_VerifyMac not loaded")
	}
	return _cSSM_VerifyMac(CCHandle, DataBufs, DataBufCount, Mac)
}

var _cSSM_VerifyMacFinal func(CCHandle CSSM_CC_HANDLE, Mac unsafe.Pointer) CSSM_RETURN

// CSSM_VerifyMacFinal.
//
// See: https://developer.apple.com/documentation/Security/CSSM_VerifyMacFinal
func CSSM_VerifyMacFinal(CCHandle CSSM_CC_HANDLE, Mac unsafe.Pointer) CSSM_RETURN {
	if _cSSM_VerifyMacFinal == nil {
		panic("Security: symbol CSSM_VerifyMacFinal not loaded")
	}
	return _cSSM_VerifyMacFinal(CCHandle, Mac)
}

var _cSSM_VerifyMacInit func(CCHandle CSSM_CC_HANDLE) CSSM_RETURN

// CSSM_VerifyMacInit.
//
// See: https://developer.apple.com/documentation/Security/CSSM_VerifyMacInit
func CSSM_VerifyMacInit(CCHandle CSSM_CC_HANDLE) CSSM_RETURN {
	if _cSSM_VerifyMacInit == nil {
		panic("Security: symbol CSSM_VerifyMacInit not loaded")
	}
	return _cSSM_VerifyMacInit(CCHandle)
}

var _cSSM_VerifyMacUpdate func(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) CSSM_RETURN

// CSSM_VerifyMacUpdate.
//
// See: https://developer.apple.com/documentation/Security/CSSM_VerifyMacUpdate
func CSSM_VerifyMacUpdate(CCHandle CSSM_CC_HANDLE, DataBufs unsafe.Pointer, DataBufCount uint32) CSSM_RETURN {
	if _cSSM_VerifyMacUpdate == nil {
		panic("Security: symbol CSSM_VerifyMacUpdate not loaded")
	}
	return _cSSM_VerifyMacUpdate(CCHandle, DataBufs, DataBufCount)
}

var _cSSM_WrapKey func(CCHandle CSSM_CC_HANDLE, AccessCred unsafe.Pointer, Key unsafe.Pointer, DescriptiveData unsafe.Pointer, WrappedKey unsafe.Pointer) CSSM_RETURN

// CSSM_WrapKey.
//
// See: https://developer.apple.com/documentation/Security/CSSM_WrapKey
func CSSM_WrapKey(CCHandle CSSM_CC_HANDLE, AccessCred unsafe.Pointer, Key unsafe.Pointer, DescriptiveData unsafe.Pointer, WrappedKey unsafe.Pointer) CSSM_RETURN {
	if _cSSM_WrapKey == nil {
		panic("Security: symbol CSSM_WrapKey not loaded")
	}
	return _cSSM_WrapKey(CCHandle, AccessCred, Key, DescriptiveData, WrappedKey)
}

var _cSSM_WrapKeyP func(CCHandle CSSM_CC_HANDLE, AccessCred unsafe.Pointer, Key unsafe.Pointer, DescriptiveData unsafe.Pointer, WrappedKey unsafe.Pointer, Privilege CSSM_PRIVILEGE) CSSM_RETURN

// CSSM_WrapKeyP.
//
// See: https://developer.apple.com/documentation/Security/CSSM_WrapKeyP
func CSSM_WrapKeyP(CCHandle CSSM_CC_HANDLE, AccessCred unsafe.Pointer, Key unsafe.Pointer, DescriptiveData unsafe.Pointer, WrappedKey unsafe.Pointer, Privilege CSSM_PRIVILEGE) CSSM_RETURN {
	if _cSSM_WrapKeyP == nil {
		panic("Security: symbol CSSM_WrapKeyP not loaded")
	}
	return _cSSM_WrapKeyP(CCHandle, AccessCred, Key, DescriptiveData, WrappedKey, Privilege)
}

var _mDS_Initialize func(pCallerGuid unsafe.Pointer, pMemoryFunctions unsafe.Pointer, pDlFunctions unsafe.Pointer, hMds *MDS_HANDLE) CSSM_RETURN

// MDS_Initialize.
//
// See: https://developer.apple.com/documentation/Security/MDS_Initialize
func MDS_Initialize(pCallerGuid unsafe.Pointer, pMemoryFunctions unsafe.Pointer, pDlFunctions unsafe.Pointer, hMds *MDS_HANDLE) CSSM_RETURN {
	if _mDS_Initialize == nil {
		panic("Security: symbol MDS_Initialize not loaded")
	}
	return _mDS_Initialize(pCallerGuid, pMemoryFunctions, pDlFunctions, hMds)
}

var _mDS_Install func(MdsHandle MDS_HANDLE) CSSM_RETURN

// MDS_Install.
//
// See: https://developer.apple.com/documentation/Security/MDS_Install
func MDS_Install(MdsHandle MDS_HANDLE) CSSM_RETURN {
	if _mDS_Install == nil {
		panic("Security: symbol MDS_Install not loaded")
	}
	return _mDS_Install(MdsHandle)
}

var _mDS_Terminate func(MdsHandle MDS_HANDLE) CSSM_RETURN

// MDS_Terminate.
//
// See: https://developer.apple.com/documentation/Security/MDS_Terminate
func MDS_Terminate(MdsHandle MDS_HANDLE) CSSM_RETURN {
	if _mDS_Terminate == nil {
		panic("Security: symbol MDS_Terminate not loaded")
	}
	return _mDS_Terminate(MdsHandle)
}

var _mDS_Uninstall func(MdsHandle MDS_HANDLE) CSSM_RETURN

// MDS_Uninstall.
//
// See: https://developer.apple.com/documentation/Security/MDS_Uninstall
func MDS_Uninstall(MdsHandle MDS_HANDLE) CSSM_RETURN {
	if _mDS_Uninstall == nil {
		panic("Security: symbol MDS_Uninstall not loaded")
	}
	return _mDS_Uninstall(MdsHandle)
}

var _secAccessControlCreateWithFlags func(allocator corefoundation.CFAllocatorRef, protection corefoundation.CFTypeRef, flags SecAccessControlCreateFlags, err *corefoundation.CFErrorRef) SecAccessControlRef

// SecAccessControlCreateWithFlags creates a new access control object with the specified protection type and flags.
//
// See: https://developer.apple.com/documentation/Security/SecAccessControlCreateWithFlags(_:_:_:_:)
func SecAccessControlCreateWithFlags(allocator corefoundation.CFAllocatorRef, protection corefoundation.CFTypeRef, flags SecAccessControlCreateFlags, err *corefoundation.CFErrorRef) SecAccessControlRef {
	if _secAccessControlCreateWithFlags == nil {
		panic("Security: symbol SecAccessControlCreateWithFlags not loaded")
	}
	return _secAccessControlCreateWithFlags(allocator, protection, flags, err)
}

var _secAccessControlGetTypeID func() uint

// SecAccessControlGetTypeID returns the unique identifier of the opaque type to which a keychain item access control object belongs.
//
// See: https://developer.apple.com/documentation/Security/SecAccessControlGetTypeID()
func SecAccessControlGetTypeID() uint {
	if _secAccessControlGetTypeID == nil {
		panic("Security: symbol SecAccessControlGetTypeID not loaded")
	}
	return _secAccessControlGetTypeID()
}

var _secAddSharedWebCredential func(fqdn corefoundation.CFStringRef, account corefoundation.CFStringRef, password corefoundation.CFStringRef)

// SecAddSharedWebCredential asynchronously stores (or updates) a shared password for a website.
//
// Deprecated: Deprecated since macOS 26.2. Use ASCredentialDataManager.save(password:for:title:anchor:) (AuthenticationServices framework)
//
// See: https://developer.apple.com/documentation/Security/SecAddSharedWebCredential(_:_:_:_:)
func SecAddSharedWebCredential(fqdn corefoundation.CFStringRef, account corefoundation.CFStringRef, password corefoundation.CFStringRef) {
	if _secAddSharedWebCredential == nil {
		panic("Security: symbol SecAddSharedWebCredential not loaded")
	}
	_secAddSharedWebCredential(fqdn, account, password)
}

var _secAsn1AllocCopy func(coder uintptr, src unsafe.Pointer, len_ uintptr, dest unsafe.Pointer) int32

// SecAsn1AllocCopy allocates memory for an item’s data field in the coder object’s memory pool and copies in a block of data.
//
// See: https://developer.apple.com/documentation/Security/SecAsn1AllocCopy
func SecAsn1AllocCopy(coder uintptr, src unsafe.Pointer, len_ uintptr, dest unsafe.Pointer) int32 {
	if _secAsn1AllocCopy == nil {
		panic("Security: symbol SecAsn1AllocCopy not loaded")
	}
	return _secAsn1AllocCopy(coder, src, len_, dest)
}

var _secAsn1AllocCopyItem func(coder uintptr, src unsafe.Pointer, dest unsafe.Pointer) int32

// SecAsn1AllocCopyItem allocates memory for an item’s data field in the coder object’s memory pool and copies in a block of data from another item.
//
// See: https://developer.apple.com/documentation/Security/SecAsn1AllocCopyItem
func SecAsn1AllocCopyItem(coder uintptr, src unsafe.Pointer, dest unsafe.Pointer) int32 {
	if _secAsn1AllocCopyItem == nil {
		panic("Security: symbol SecAsn1AllocCopyItem not loaded")
	}
	return _secAsn1AllocCopyItem(coder, src, dest)
}

var _secAsn1AllocItem func(coder uintptr, item unsafe.Pointer, len_ uintptr) int32

// SecAsn1AllocItem allocates memory for an item’s data field in the coder object’s memory pool.
//
// See: https://developer.apple.com/documentation/Security/SecAsn1AllocItem
func SecAsn1AllocItem(coder uintptr, item unsafe.Pointer, len_ uintptr) int32 {
	if _secAsn1AllocItem == nil {
		panic("Security: symbol SecAsn1AllocItem not loaded")
	}
	return _secAsn1AllocItem(coder, item, len_)
}

var _secAsn1CoderCreate func(coder unsafe.Pointer) int32

// SecAsn1CoderCreate creates an ASN.1 coder object.
//
// See: https://developer.apple.com/documentation/Security/SecAsn1CoderCreate
func SecAsn1CoderCreate(coder unsafe.Pointer) int32 {
	if _secAsn1CoderCreate == nil {
		panic("Security: symbol SecAsn1CoderCreate not loaded")
	}
	return _secAsn1CoderCreate(coder)
}

var _secAsn1CoderRelease func(coder uintptr) int32

// SecAsn1CoderRelease destroys an ASN.1 coder object and releases all of its memory.
//
// See: https://developer.apple.com/documentation/Security/SecAsn1CoderRelease
func SecAsn1CoderRelease(coder uintptr) int32 {
	if _secAsn1CoderRelease == nil {
		panic("Security: symbol SecAsn1CoderRelease not loaded")
	}
	return _secAsn1CoderRelease(coder)
}

var _secAsn1Decode func(coder uintptr, src unsafe.Pointer, len_ uintptr, templates unsafe.Pointer, dest unsafe.Pointer) int32

// SecAsn1Decode decodes untyped DER data.
//
// See: https://developer.apple.com/documentation/Security/SecAsn1Decode
func SecAsn1Decode(coder uintptr, src unsafe.Pointer, len_ uintptr, templates unsafe.Pointer, dest unsafe.Pointer) int32 {
	if _secAsn1Decode == nil {
		panic("Security: symbol SecAsn1Decode not loaded")
	}
	return _secAsn1Decode(coder, src, len_, templates, dest)
}

var _secAsn1DecodeData func(coder uintptr, src unsafe.Pointer, templ unsafe.Pointer, dest unsafe.Pointer) int32

// SecAsn1DecodeData decodes an ASN.1 item in DER format.
//
// See: https://developer.apple.com/documentation/Security/SecAsn1DecodeData
func SecAsn1DecodeData(coder uintptr, src unsafe.Pointer, templ unsafe.Pointer, dest unsafe.Pointer) int32 {
	if _secAsn1DecodeData == nil {
		panic("Security: symbol SecAsn1DecodeData not loaded")
	}
	return _secAsn1DecodeData(coder, src, templ, dest)
}

var _secAsn1EncodeItem func(coder uintptr, src unsafe.Pointer, templates unsafe.Pointer, dest unsafe.Pointer) int32

// SecAsn1EncodeItem encodes data in DER format.
//
// See: https://developer.apple.com/documentation/Security/SecAsn1EncodeItem
func SecAsn1EncodeItem(coder uintptr, src unsafe.Pointer, templates unsafe.Pointer, dest unsafe.Pointer) int32 {
	if _secAsn1EncodeItem == nil {
		panic("Security: symbol SecAsn1EncodeItem not loaded")
	}
	return _secAsn1EncodeItem(coder, src, templates, dest)
}

var _secAsn1Malloc func(coder uintptr, len_ uintptr) unsafe.Pointer

// SecAsn1Malloc allocates memory in the coder object’s memory pool.
//
// See: https://developer.apple.com/documentation/Security/SecAsn1Malloc
func SecAsn1Malloc(coder uintptr, len_ uintptr) unsafe.Pointer {
	if _secAsn1Malloc == nil {
		panic("Security: symbol SecAsn1Malloc not loaded")
	}
	return _secAsn1Malloc(coder, len_)
}

var _secAsn1OidCompare func(oid1 unsafe.Pointer, oid2 unsafe.Pointer) bool

// SecAsn1OidCompare compares two decoded object identifiers.
//
// See: https://developer.apple.com/documentation/Security/SecAsn1OidCompare
func SecAsn1OidCompare(oid1 unsafe.Pointer, oid2 unsafe.Pointer) bool {
	if _secAsn1OidCompare == nil {
		panic("Security: symbol SecAsn1OidCompare not loaded")
	}
	return _secAsn1OidCompare(oid1, oid2)
}

var _secCertificateAddToKeychain func(certificate SecCertificateRef, keychain SecKeychainRef) int32

// SecCertificateAddToKeychain adds a certificate to a keychain.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateAddToKeychain(_:_:)
func SecCertificateAddToKeychain(certificate SecCertificateRef, keychain SecKeychainRef) int32 {
	if _secCertificateAddToKeychain == nil {
		panic("Security: symbol SecCertificateAddToKeychain not loaded")
	}
	return _secCertificateAddToKeychain(certificate, keychain)
}

var _secCertificateCopyCommonName func(certificate SecCertificateRef, commonName *corefoundation.CFStringRef) int32

// SecCertificateCopyCommonName retrieves the common name of the subject of a certificate.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopyCommonName(_:_:)
func SecCertificateCopyCommonName(certificate SecCertificateRef, commonName *corefoundation.CFStringRef) int32 {
	if _secCertificateCopyCommonName == nil {
		panic("Security: symbol SecCertificateCopyCommonName not loaded")
	}
	return _secCertificateCopyCommonName(certificate, commonName)
}

var _secCertificateCopyData func(certificate SecCertificateRef) corefoundation.CFDataRef

// SecCertificateCopyData returns a DER representation of a certificate given a certificate object.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopyData(_:)
func SecCertificateCopyData(certificate SecCertificateRef) corefoundation.CFDataRef {
	if _secCertificateCopyData == nil {
		panic("Security: symbol SecCertificateCopyData not loaded")
	}
	return _secCertificateCopyData(certificate)
}

var _secCertificateCopyEmailAddresses func(certificate SecCertificateRef, emailAddresses *corefoundation.CFArrayRef) int32

// SecCertificateCopyEmailAddresses retrieves the email addresses for the subject of a certificate.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopyEmailAddresses(_:_:)
func SecCertificateCopyEmailAddresses(certificate SecCertificateRef, emailAddresses *corefoundation.CFArrayRef) int32 {
	if _secCertificateCopyEmailAddresses == nil {
		panic("Security: symbol SecCertificateCopyEmailAddresses not loaded")
	}
	return _secCertificateCopyEmailAddresses(certificate, emailAddresses)
}

var _secCertificateCopyKey func(certificate SecCertificateRef) SecKeyRef

// SecCertificateCopyKey retrieves the public key for a given certificate.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopyKey(_:)
func SecCertificateCopyKey(certificate SecCertificateRef) SecKeyRef {
	if _secCertificateCopyKey == nil {
		panic("Security: symbol SecCertificateCopyKey not loaded")
	}
	return _secCertificateCopyKey(certificate)
}

var _secCertificateCopyLongDescription func(alloc corefoundation.CFAllocatorRef, certificate SecCertificateRef, err *corefoundation.CFErrorRef) corefoundation.CFStringRef

// SecCertificateCopyLongDescription returns a copy of the long description of a certificate.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopyLongDescription(_:_:_:)
func SecCertificateCopyLongDescription(alloc corefoundation.CFAllocatorRef, certificate SecCertificateRef, err *corefoundation.CFErrorRef) corefoundation.CFStringRef {
	if _secCertificateCopyLongDescription == nil {
		panic("Security: symbol SecCertificateCopyLongDescription not loaded")
	}
	return _secCertificateCopyLongDescription(alloc, certificate, err)
}

var _secCertificateCopyNormalizedIssuerSequence func(certificate SecCertificateRef) corefoundation.CFDataRef

// SecCertificateCopyNormalizedIssuerSequence retrieves the normalized issuer sequence from a certificate.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopyNormalizedIssuerSequence(_:)
func SecCertificateCopyNormalizedIssuerSequence(certificate SecCertificateRef) corefoundation.CFDataRef {
	if _secCertificateCopyNormalizedIssuerSequence == nil {
		panic("Security: symbol SecCertificateCopyNormalizedIssuerSequence not loaded")
	}
	return _secCertificateCopyNormalizedIssuerSequence(certificate)
}

var _secCertificateCopyNormalizedSubjectSequence func(certificate SecCertificateRef) corefoundation.CFDataRef

// SecCertificateCopyNormalizedSubjectSequence retrieves the normalized subject sequence from a certificate.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopyNormalizedSubjectSequence(_:)
func SecCertificateCopyNormalizedSubjectSequence(certificate SecCertificateRef) corefoundation.CFDataRef {
	if _secCertificateCopyNormalizedSubjectSequence == nil {
		panic("Security: symbol SecCertificateCopyNormalizedSubjectSequence not loaded")
	}
	return _secCertificateCopyNormalizedSubjectSequence(certificate)
}

var _secCertificateCopyNotValidAfterDate func(certificate SecCertificateRef) corefoundation.CFDateRef

// SecCertificateCopyNotValidAfterDate.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopyNotValidAfterDate(_:)
func SecCertificateCopyNotValidAfterDate(certificate SecCertificateRef) corefoundation.CFDateRef {
	if _secCertificateCopyNotValidAfterDate == nil {
		panic("Security: symbol SecCertificateCopyNotValidAfterDate not loaded")
	}
	return _secCertificateCopyNotValidAfterDate(certificate)
}

var _secCertificateCopyNotValidBeforeDate func(certificate SecCertificateRef) corefoundation.CFDateRef

// SecCertificateCopyNotValidBeforeDate.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopyNotValidBeforeDate(_:)
func SecCertificateCopyNotValidBeforeDate(certificate SecCertificateRef) corefoundation.CFDateRef {
	if _secCertificateCopyNotValidBeforeDate == nil {
		panic("Security: symbol SecCertificateCopyNotValidBeforeDate not loaded")
	}
	return _secCertificateCopyNotValidBeforeDate(certificate)
}

var _secCertificateCopyPreference func(name corefoundation.CFStringRef, keyUsage uint32, certificate *SecCertificateRef) int32

// SecCertificateCopyPreference retrieves the preferred certificate for the specified name and key use.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopyPreference
func SecCertificateCopyPreference(name corefoundation.CFStringRef, keyUsage uint32, certificate *SecCertificateRef) int32 {
	if _secCertificateCopyPreference == nil {
		panic("Security: symbol SecCertificateCopyPreference not loaded")
	}
	return _secCertificateCopyPreference(name, keyUsage, certificate)
}

var _secCertificateCopyPreferred func(name corefoundation.CFStringRef, keyUsage corefoundation.CFArrayRef) SecCertificateRef

// SecCertificateCopyPreferred returns the preferred certificate for the specified name and key usage.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopyPreferred(_:_:)
func SecCertificateCopyPreferred(name corefoundation.CFStringRef, keyUsage corefoundation.CFArrayRef) SecCertificateRef {
	if _secCertificateCopyPreferred == nil {
		panic("Security: symbol SecCertificateCopyPreferred not loaded")
	}
	return _secCertificateCopyPreferred(name, keyUsage)
}

var _secCertificateCopySerialNumberData func(certificate SecCertificateRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef

// SecCertificateCopySerialNumberData returns the certificate’s serial number.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopySerialNumberData(_:_:)
func SecCertificateCopySerialNumberData(certificate SecCertificateRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef {
	if _secCertificateCopySerialNumberData == nil {
		panic("Security: symbol SecCertificateCopySerialNumberData not loaded")
	}
	return _secCertificateCopySerialNumberData(certificate, err)
}

var _secCertificateCopyShortDescription func(alloc corefoundation.CFAllocatorRef, certificate SecCertificateRef, err *corefoundation.CFErrorRef) corefoundation.CFStringRef

// SecCertificateCopyShortDescription returns a copy of the short description of a certificate.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopyShortDescription(_:_:_:)
func SecCertificateCopyShortDescription(alloc corefoundation.CFAllocatorRef, certificate SecCertificateRef, err *corefoundation.CFErrorRef) corefoundation.CFStringRef {
	if _secCertificateCopyShortDescription == nil {
		panic("Security: symbol SecCertificateCopyShortDescription not loaded")
	}
	return _secCertificateCopyShortDescription(alloc, certificate, err)
}

var _secCertificateCopySubjectSummary func(certificate SecCertificateRef) corefoundation.CFStringRef

// SecCertificateCopySubjectSummary returns a human-readable summary of a certificate.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopySubjectSummary(_:)
func SecCertificateCopySubjectSummary(certificate SecCertificateRef) corefoundation.CFStringRef {
	if _secCertificateCopySubjectSummary == nil {
		panic("Security: symbol SecCertificateCopySubjectSummary not loaded")
	}
	return _secCertificateCopySubjectSummary(certificate)
}

var _secCertificateCopyValues func(certificate SecCertificateRef, keys corefoundation.CFArrayRef, err *corefoundation.CFErrorRef) corefoundation.CFDictionaryRef

// SecCertificateCopyValues creates a dictionary that represents a certificate’s contents.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCopyValues(_:_:_:)
func SecCertificateCopyValues(certificate SecCertificateRef, keys corefoundation.CFArrayRef, err *corefoundation.CFErrorRef) corefoundation.CFDictionaryRef {
	if _secCertificateCopyValues == nil {
		panic("Security: symbol SecCertificateCopyValues not loaded")
	}
	return _secCertificateCopyValues(certificate, keys, err)
}

var _secCertificateCreateWithData func(allocator corefoundation.CFAllocatorRef, data corefoundation.CFDataRef) SecCertificateRef

// SecCertificateCreateWithData creates a certificate object from a DER representation of a certificate.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateCreateWithData(_:_:)
func SecCertificateCreateWithData(allocator corefoundation.CFAllocatorRef, data corefoundation.CFDataRef) SecCertificateRef {
	if _secCertificateCreateWithData == nil {
		panic("Security: symbol SecCertificateCreateWithData not loaded")
	}
	return _secCertificateCreateWithData(allocator, data)
}

var _secCertificateGetAlgorithmID func(certificate SecCertificateRef, algid unsafe.Pointer) int32

// SecCertificateGetAlgorithmID retrieves the algorithm identifier for a certificate.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateGetAlgorithmID
func SecCertificateGetAlgorithmID(certificate SecCertificateRef, algid unsafe.Pointer) int32 {
	if _secCertificateGetAlgorithmID == nil {
		panic("Security: symbol SecCertificateGetAlgorithmID not loaded")
	}
	return _secCertificateGetAlgorithmID(certificate, algid)
}

var _secCertificateGetCLHandle func(certificate SecCertificateRef, clHandle unsafe.Pointer) int32

// SecCertificateGetCLHandle retrieves the certificate library handle from a certificate object.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateGetCLHandle
func SecCertificateGetCLHandle(certificate SecCertificateRef, clHandle unsafe.Pointer) int32 {
	if _secCertificateGetCLHandle == nil {
		panic("Security: symbol SecCertificateGetCLHandle not loaded")
	}
	return _secCertificateGetCLHandle(certificate, clHandle)
}

var _secCertificateGetData func(certificate SecCertificateRef, data unsafe.Pointer) int32

// SecCertificateGetData retrieves the data for a certificate.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateGetData
func SecCertificateGetData(certificate SecCertificateRef, data unsafe.Pointer) int32 {
	if _secCertificateGetData == nil {
		panic("Security: symbol SecCertificateGetData not loaded")
	}
	return _secCertificateGetData(certificate, data)
}

var _secCertificateGetIssuer func(certificate SecCertificateRef, issuer unsafe.Pointer) int32

// SecCertificateGetIssuer unsupported.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateGetIssuer
func SecCertificateGetIssuer(certificate SecCertificateRef, issuer unsafe.Pointer) int32 {
	if _secCertificateGetIssuer == nil {
		panic("Security: symbol SecCertificateGetIssuer not loaded")
	}
	return _secCertificateGetIssuer(certificate, issuer)
}

var _secCertificateGetSubject func(certificate SecCertificateRef, subject unsafe.Pointer) int32

// SecCertificateGetSubject unsupported.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateGetSubject
func SecCertificateGetSubject(certificate SecCertificateRef, subject unsafe.Pointer) int32 {
	if _secCertificateGetSubject == nil {
		panic("Security: symbol SecCertificateGetSubject not loaded")
	}
	return _secCertificateGetSubject(certificate, subject)
}

var _secCertificateGetType func(certificate SecCertificateRef, certificateType unsafe.Pointer) int32

// SecCertificateGetType retrieves the type of a specified certificate.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateGetType
func SecCertificateGetType(certificate SecCertificateRef, certificateType unsafe.Pointer) int32 {
	if _secCertificateGetType == nil {
		panic("Security: symbol SecCertificateGetType not loaded")
	}
	return _secCertificateGetType(certificate, certificateType)
}

var _secCertificateGetTypeID func() uint

// SecCertificateGetTypeID returns the unique identifier of the opaque type to which a certificate object belongs.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateGetTypeID()
func SecCertificateGetTypeID() uint {
	if _secCertificateGetTypeID == nil {
		panic("Security: symbol SecCertificateGetTypeID not loaded")
	}
	return _secCertificateGetTypeID()
}

var _secCertificateSetPreference func(certificate SecCertificateRef, name corefoundation.CFStringRef, keyUsage uint32, date corefoundation.CFDateRef) int32

// SecCertificateSetPreference sets the preferred certificate for a specified name, key use, and date.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateSetPreference
func SecCertificateSetPreference(certificate SecCertificateRef, name corefoundation.CFStringRef, keyUsage uint32, date corefoundation.CFDateRef) int32 {
	if _secCertificateSetPreference == nil {
		panic("Security: symbol SecCertificateSetPreference not loaded")
	}
	return _secCertificateSetPreference(certificate, name, keyUsage, date)
}

var _secCertificateSetPreferred func(certificate SecCertificateRef, name corefoundation.CFStringRef, keyUsage corefoundation.CFArrayRef) int32

// SecCertificateSetPreferred sets the certificate that should be preferred for the specified name and key use.
//
// See: https://developer.apple.com/documentation/Security/SecCertificateSetPreferred(_:_:_:)
func SecCertificateSetPreferred(certificate SecCertificateRef, name corefoundation.CFStringRef, keyUsage corefoundation.CFArrayRef) int32 {
	if _secCertificateSetPreferred == nil {
		panic("Security: symbol SecCertificateSetPreferred not loaded")
	}
	return _secCertificateSetPreferred(certificate, name, keyUsage)
}

var _secCodeCheckValidity func(code SecCodeRef, flags SecCSFlags, requirement SecRequirementRef) int32

// SecCodeCheckValidity performs dynamic validation of signed code.
//
// See: https://developer.apple.com/documentation/Security/SecCodeCheckValidity(_:_:_:)
func SecCodeCheckValidity(code SecCodeRef, flags SecCSFlags, requirement SecRequirementRef) int32 {
	if _secCodeCheckValidity == nil {
		panic("Security: symbol SecCodeCheckValidity not loaded")
	}
	return _secCodeCheckValidity(code, flags, requirement)
}

var _secCodeCheckValidityWithErrors func(code SecCodeRef, flags SecCSFlags, requirement SecRequirementRef, errors *corefoundation.CFErrorRef) int32

// SecCodeCheckValidityWithErrors performs dynamic validation of signed code and returns detailed error information in the case of failure.
//
// See: https://developer.apple.com/documentation/Security/SecCodeCheckValidityWithErrors(_:_:_:_:)
func SecCodeCheckValidityWithErrors(code SecCodeRef, flags SecCSFlags, requirement SecRequirementRef, errors *corefoundation.CFErrorRef) int32 {
	if _secCodeCheckValidityWithErrors == nil {
		panic("Security: symbol SecCodeCheckValidityWithErrors not loaded")
	}
	return _secCodeCheckValidityWithErrors(code, flags, requirement, errors)
}

var _secCodeCopyDesignatedRequirement func(code SecStaticCodeRef, flags SecCSFlags, requirement *SecRequirementRef) int32

// SecCodeCopyDesignatedRequirement retrieves the designated code requirement of signed code.
//
// See: https://developer.apple.com/documentation/Security/SecCodeCopyDesignatedRequirement(_:_:_:)
func SecCodeCopyDesignatedRequirement(code SecStaticCodeRef, flags SecCSFlags, requirement *SecRequirementRef) int32 {
	if _secCodeCopyDesignatedRequirement == nil {
		panic("Security: symbol SecCodeCopyDesignatedRequirement not loaded")
	}
	return _secCodeCopyDesignatedRequirement(code, flags, requirement)
}

var _secCodeCopyGuestWithAttributes func(host SecCodeRef, attributes corefoundation.CFDictionaryRef, flags SecCSFlags, guest *SecCodeRef) int32

// SecCodeCopyGuestWithAttributes asks a code host to identify one of its guests given the type and value of specific attributes of the guest code.
//
// See: https://developer.apple.com/documentation/Security/SecCodeCopyGuestWithAttributes(_:_:_:_:)
func SecCodeCopyGuestWithAttributes(host SecCodeRef, attributes corefoundation.CFDictionaryRef, flags SecCSFlags, guest *SecCodeRef) int32 {
	if _secCodeCopyGuestWithAttributes == nil {
		panic("Security: symbol SecCodeCopyGuestWithAttributes not loaded")
	}
	return _secCodeCopyGuestWithAttributes(host, attributes, flags, guest)
}

var _secCodeCopyHost func(guest SecCodeRef, flags SecCSFlags, host *SecCodeRef) int32

// SecCodeCopyHost retrieves the code object for the host of specified guest code.
//
// See: https://developer.apple.com/documentation/Security/SecCodeCopyHost(_:_:_:)
func SecCodeCopyHost(guest SecCodeRef, flags SecCSFlags, host *SecCodeRef) int32 {
	if _secCodeCopyHost == nil {
		panic("Security: symbol SecCodeCopyHost not loaded")
	}
	return _secCodeCopyHost(guest, flags, host)
}

var _secCodeCopyPath func(staticCode SecStaticCodeRef, flags SecCSFlags, path *corefoundation.CFURLRef) int32

// SecCodeCopyPath retrieves the location on disk of signed code, given a code or static code object.
//
// See: https://developer.apple.com/documentation/Security/SecCodeCopyPath(_:_:_:)
func SecCodeCopyPath(staticCode SecStaticCodeRef, flags SecCSFlags, path *corefoundation.CFURLRef) int32 {
	if _secCodeCopyPath == nil {
		panic("Security: symbol SecCodeCopyPath not loaded")
	}
	return _secCodeCopyPath(staticCode, flags, path)
}

var _secCodeCopySelf func(flags SecCSFlags, self *SecCodeRef) int32

// SecCodeCopySelf retrieves the code object for the code making the call.
//
// See: https://developer.apple.com/documentation/Security/SecCodeCopySelf(_:_:)
func SecCodeCopySelf(flags SecCSFlags, self *SecCodeRef) int32 {
	if _secCodeCopySelf == nil {
		panic("Security: symbol SecCodeCopySelf not loaded")
	}
	return _secCodeCopySelf(flags, self)
}

var _secCodeCopySigningInformation func(code SecStaticCodeRef, flags SecCSFlags, information *corefoundation.CFDictionaryRef) int32

// SecCodeCopySigningInformation retrieves various pieces of information from a code signature.
//
// See: https://developer.apple.com/documentation/Security/SecCodeCopySigningInformation(_:_:_:)
func SecCodeCopySigningInformation(code SecStaticCodeRef, flags SecCSFlags, information *corefoundation.CFDictionaryRef) int32 {
	if _secCodeCopySigningInformation == nil {
		panic("Security: symbol SecCodeCopySigningInformation not loaded")
	}
	return _secCodeCopySigningInformation(code, flags, information)
}

var _secCodeCopyStaticCode func(code SecCodeRef, flags SecCSFlags, staticCode *SecStaticCodeRef) int32

// SecCodeCopyStaticCode returns a static code object representing the on-disk version of the given running code.
//
// See: https://developer.apple.com/documentation/Security/SecCodeCopyStaticCode(_:_:_:)
func SecCodeCopyStaticCode(code SecCodeRef, flags SecCSFlags, staticCode *SecStaticCodeRef) int32 {
	if _secCodeCopyStaticCode == nil {
		panic("Security: symbol SecCodeCopyStaticCode not loaded")
	}
	return _secCodeCopyStaticCode(code, flags, staticCode)
}

var _secCodeCreateWithXPCMessage func(message unsafe.Pointer, flags SecCSFlags, target *SecCodeRef) int32

// SecCodeCreateWithXPCMessage.
//
// See: https://developer.apple.com/documentation/Security/SecCodeCreateWithXPCMessage(_:_:_:)
func SecCodeCreateWithXPCMessage(message unsafe.Pointer, flags SecCSFlags, target *SecCodeRef) int32 {
	if _secCodeCreateWithXPCMessage == nil {
		panic("Security: symbol SecCodeCreateWithXPCMessage not loaded")
	}
	return _secCodeCreateWithXPCMessage(message, flags, target)
}

var _secCodeGetTypeID func() uint

// SecCodeGetTypeID returns the unique identifier of the opaque type to which a code object belongs.
//
// See: https://developer.apple.com/documentation/Security/SecCodeGetTypeID()
func SecCodeGetTypeID() uint {
	if _secCodeGetTypeID == nil {
		panic("Security: symbol SecCodeGetTypeID not loaded")
	}
	return _secCodeGetTypeID()
}

var _secCodeMapMemory func(code SecStaticCodeRef, flags SecCSFlags) int32

// SecCodeMapMemory asks the kernel to accept the signing information currently attached to a code object and uses it to validate memory page-ins.
//
// See: https://developer.apple.com/documentation/Security/SecCodeMapMemory(_:_:)
func SecCodeMapMemory(code SecStaticCodeRef, flags SecCSFlags) int32 {
	if _secCodeMapMemory == nil {
		panic("Security: symbol SecCodeMapMemory not loaded")
	}
	return _secCodeMapMemory(code, flags)
}

var _secCodeValidateFileResource func(code SecStaticCodeRef, relativePath corefoundation.CFStringRef, fileData corefoundation.CFDataRef, flags SecCSFlags) int32

// SecCodeValidateFileResource.
//
// See: https://developer.apple.com/documentation/Security/SecCodeValidateFileResource(_:_:_:_:)
func SecCodeValidateFileResource(code SecStaticCodeRef, relativePath corefoundation.CFStringRef, fileData corefoundation.CFDataRef, flags SecCSFlags) int32 {
	if _secCodeValidateFileResource == nil {
		panic("Security: symbol SecCodeValidateFileResource not loaded")
	}
	return _secCodeValidateFileResource(code, relativePath, fileData, flags)
}

var _secCopyErrorMessageString func(status int32, reserved unsafe.Pointer) corefoundation.CFStringRef

// SecCopyErrorMessageString returns a string explaining the meaning of a security result code.
//
// See: https://developer.apple.com/documentation/Security/SecCopyErrorMessageString(_:_:)
func SecCopyErrorMessageString(status int32, reserved unsafe.Pointer) corefoundation.CFStringRef {
	if _secCopyErrorMessageString == nil {
		panic("Security: symbol SecCopyErrorMessageString not loaded")
	}
	return _secCopyErrorMessageString(status, reserved)
}

var _secCreateSharedWebCredentialPassword func() corefoundation.CFStringRef

// SecCreateSharedWebCredentialPassword returns a randomly generated password.
//
// See: https://developer.apple.com/documentation/Security/SecCreateSharedWebCredentialPassword()
func SecCreateSharedWebCredentialPassword() corefoundation.CFStringRef {
	if _secCreateSharedWebCredentialPassword == nil {
		panic("Security: symbol SecCreateSharedWebCredentialPassword not loaded")
	}
	return _secCreateSharedWebCredentialPassword()
}

var _secHostCreateGuest func(host SecGuestRef, status uint32, path corefoundation.CFURLRef, attributes corefoundation.CFDictionaryRef, flags SecCSFlags, newGuest *SecGuestRef) int32

// SecHostCreateGuest creates a new guest and describes its initial properties.
//
// See: https://developer.apple.com/documentation/Security/SecHostCreateGuest
func SecHostCreateGuest(host SecGuestRef, status uint32, path corefoundation.CFURLRef, attributes corefoundation.CFDictionaryRef, flags SecCSFlags, newGuest *SecGuestRef) int32 {
	if _secHostCreateGuest == nil {
		panic("Security: symbol SecHostCreateGuest not loaded")
	}
	return _secHostCreateGuest(host, status, path, attributes, flags, newGuest)
}

var _secHostRemoveGuest func(host SecGuestRef, guest SecGuestRef, flags SecCSFlags) int32

// SecHostRemoveGuest removes a guest from a host.
//
// See: https://developer.apple.com/documentation/Security/SecHostRemoveGuest
func SecHostRemoveGuest(host SecGuestRef, guest SecGuestRef, flags SecCSFlags) int32 {
	if _secHostRemoveGuest == nil {
		panic("Security: symbol SecHostRemoveGuest not loaded")
	}
	return _secHostRemoveGuest(host, guest, flags)
}

var _secHostSelectGuest func(guestRef SecGuestRef, flags SecCSFlags) int32

// SecHostSelectGuest makes the calling thread the proxy for a specified guest.
//
// See: https://developer.apple.com/documentation/Security/SecHostSelectGuest
func SecHostSelectGuest(guestRef SecGuestRef, flags SecCSFlags) int32 {
	if _secHostSelectGuest == nil {
		panic("Security: symbol SecHostSelectGuest not loaded")
	}
	return _secHostSelectGuest(guestRef, flags)
}

var _secHostSelectedGuest func(flags SecCSFlags, guestRef *SecGuestRef) int32

// SecHostSelectedGuest retrieves the handle for the guest currently selected for the calling thread.
//
// See: https://developer.apple.com/documentation/Security/SecHostSelectedGuest
func SecHostSelectedGuest(flags SecCSFlags, guestRef *SecGuestRef) int32 {
	if _secHostSelectedGuest == nil {
		panic("Security: symbol SecHostSelectedGuest not loaded")
	}
	return _secHostSelectedGuest(flags, guestRef)
}

var _secHostSetGuestStatus func(guestRef SecGuestRef, status uint32, attributes corefoundation.CFDictionaryRef, flags SecCSFlags) int32

// SecHostSetGuestStatus updates the status and attributes of a particular guest.
//
// See: https://developer.apple.com/documentation/Security/SecHostSetGuestStatus
func SecHostSetGuestStatus(guestRef SecGuestRef, status uint32, attributes corefoundation.CFDictionaryRef, flags SecCSFlags) int32 {
	if _secHostSetGuestStatus == nil {
		panic("Security: symbol SecHostSetGuestStatus not loaded")
	}
	return _secHostSetGuestStatus(guestRef, status, attributes, flags)
}

var _secHostSetHostingPort func(hostingPort uint32, flags SecCSFlags) int32

// SecHostSetHostingPort tells code signing services that the calling code will directly respond to hosting inquiries over the given port.
//
// See: https://developer.apple.com/documentation/Security/SecHostSetHostingPort
func SecHostSetHostingPort(hostingPort uint32, flags SecCSFlags) int32 {
	if _secHostSetHostingPort == nil {
		panic("Security: symbol SecHostSetHostingPort not loaded")
	}
	return _secHostSetHostingPort(hostingPort, flags)
}

var _secIdentityCopyCertificate func(identityRef SecIdentityRef, certificateRef *SecCertificateRef) int32

// SecIdentityCopyCertificate retrieves a certificate associated with an identity.
//
// See: https://developer.apple.com/documentation/Security/SecIdentityCopyCertificate(_:_:)
func SecIdentityCopyCertificate(identityRef SecIdentityRef, certificateRef *SecCertificateRef) int32 {
	if _secIdentityCopyCertificate == nil {
		panic("Security: symbol SecIdentityCopyCertificate not loaded")
	}
	return _secIdentityCopyCertificate(identityRef, certificateRef)
}

var _secIdentityCopyPreference func(name corefoundation.CFStringRef, keyUsage CSSM_KEYUSE, validIssuers corefoundation.CFArrayRef, identity *SecIdentityRef) int32

// SecIdentityCopyPreference returns the preferred identity for the specified name and key use.
//
// See: https://developer.apple.com/documentation/Security/SecIdentityCopyPreference
func SecIdentityCopyPreference(name corefoundation.CFStringRef, keyUsage CSSM_KEYUSE, validIssuers corefoundation.CFArrayRef, identity *SecIdentityRef) int32 {
	if _secIdentityCopyPreference == nil {
		panic("Security: symbol SecIdentityCopyPreference not loaded")
	}
	return _secIdentityCopyPreference(name, keyUsage, validIssuers, identity)
}

var _secIdentityCopyPreferred func(name corefoundation.CFStringRef, keyUsage corefoundation.CFArrayRef, validIssuers corefoundation.CFArrayRef) SecIdentityRef

// SecIdentityCopyPreferred retrieves the preferred identity for the specified name and key use.
//
// See: https://developer.apple.com/documentation/Security/SecIdentityCopyPreferred(_:_:_:)
func SecIdentityCopyPreferred(name corefoundation.CFStringRef, keyUsage corefoundation.CFArrayRef, validIssuers corefoundation.CFArrayRef) SecIdentityRef {
	if _secIdentityCopyPreferred == nil {
		panic("Security: symbol SecIdentityCopyPreferred not loaded")
	}
	return _secIdentityCopyPreferred(name, keyUsage, validIssuers)
}

var _secIdentityCopyPrivateKey func(identityRef SecIdentityRef, privateKeyRef *SecKeyRef) int32

// SecIdentityCopyPrivateKey retrieves the private key associated with an identity.
//
// See: https://developer.apple.com/documentation/Security/SecIdentityCopyPrivateKey(_:_:)
func SecIdentityCopyPrivateKey(identityRef SecIdentityRef, privateKeyRef *SecKeyRef) int32 {
	if _secIdentityCopyPrivateKey == nil {
		panic("Security: symbol SecIdentityCopyPrivateKey not loaded")
	}
	return _secIdentityCopyPrivateKey(identityRef, privateKeyRef)
}

var _secIdentityCopySystemIdentity func(domain corefoundation.CFStringRef, idRef *SecIdentityRef, actualDomain *corefoundation.CFStringRef) int32

// SecIdentityCopySystemIdentity obtains the system identity associated with a specified domain.
//
// See: https://developer.apple.com/documentation/Security/SecIdentityCopySystemIdentity(_:_:_:)
func SecIdentityCopySystemIdentity(domain corefoundation.CFStringRef, idRef *SecIdentityRef, actualDomain *corefoundation.CFStringRef) int32 {
	if _secIdentityCopySystemIdentity == nil {
		panic("Security: symbol SecIdentityCopySystemIdentity not loaded")
	}
	return _secIdentityCopySystemIdentity(domain, idRef, actualDomain)
}

var _secIdentityCreate func(allocator corefoundation.CFAllocatorRef, certificate SecCertificateRef, privateKey SecKeyRef) SecIdentityRef

// SecIdentityCreate.
//
// See: https://developer.apple.com/documentation/Security/SecIdentityCreate(_:_:_:)
func SecIdentityCreate(allocator corefoundation.CFAllocatorRef, certificate SecCertificateRef, privateKey SecKeyRef) SecIdentityRef {
	if _secIdentityCreate == nil {
		panic("Security: symbol SecIdentityCreate not loaded")
	}
	return _secIdentityCreate(allocator, certificate, privateKey)
}

var _secIdentityCreateWithCertificate func(keychainOrArray corefoundation.CFTypeRef, certificateRef SecCertificateRef, identityRef *SecIdentityRef) int32

// SecIdentityCreateWithCertificate creates a new identity for a certificate and its associated private key.
//
// See: https://developer.apple.com/documentation/Security/SecIdentityCreateWithCertificate(_:_:_:)
func SecIdentityCreateWithCertificate(keychainOrArray corefoundation.CFTypeRef, certificateRef SecCertificateRef, identityRef *SecIdentityRef) int32 {
	if _secIdentityCreateWithCertificate == nil {
		panic("Security: symbol SecIdentityCreateWithCertificate not loaded")
	}
	return _secIdentityCreateWithCertificate(keychainOrArray, certificateRef, identityRef)
}

var _secIdentityGetTypeID func() uint

// SecIdentityGetTypeID returns the unique identifier of the opaque type to which an identity object belongs.
//
// See: https://developer.apple.com/documentation/Security/SecIdentityGetTypeID()
func SecIdentityGetTypeID() uint {
	if _secIdentityGetTypeID == nil {
		panic("Security: symbol SecIdentityGetTypeID not loaded")
	}
	return _secIdentityGetTypeID()
}

var _secIdentitySearchCopyNext func(searchRef SecIdentitySearchRef, identity *SecIdentityRef) int32

// SecIdentitySearchCopyNext finds the next identity matching specified search criteria
//
// See: https://developer.apple.com/documentation/Security/SecIdentitySearchCopyNext
func SecIdentitySearchCopyNext(searchRef SecIdentitySearchRef, identity *SecIdentityRef) int32 {
	if _secIdentitySearchCopyNext == nil {
		panic("Security: symbol SecIdentitySearchCopyNext not loaded")
	}
	return _secIdentitySearchCopyNext(searchRef, identity)
}

var _secIdentitySearchCreate func(keychainOrArray corefoundation.CFTypeRef, keyUsage CSSM_KEYUSE, searchRef *SecIdentitySearchRef) int32

// SecIdentitySearchCreate creates a search object for finding identities.
//
// See: https://developer.apple.com/documentation/Security/SecIdentitySearchCreate
func SecIdentitySearchCreate(keychainOrArray corefoundation.CFTypeRef, keyUsage CSSM_KEYUSE, searchRef *SecIdentitySearchRef) int32 {
	if _secIdentitySearchCreate == nil {
		panic("Security: symbol SecIdentitySearchCreate not loaded")
	}
	return _secIdentitySearchCreate(keychainOrArray, keyUsage, searchRef)
}

var _secIdentitySearchGetTypeID func() uint

// SecIdentitySearchGetTypeID returns the unique identifier of the opaque type to which a [SecIdentitySearch] object belongs.
//
// See: https://developer.apple.com/documentation/Security/SecIdentitySearchGetTypeID
func SecIdentitySearchGetTypeID() uint {
	if _secIdentitySearchGetTypeID == nil {
		panic("Security: symbol SecIdentitySearchGetTypeID not loaded")
	}
	return _secIdentitySearchGetTypeID()
}

var _secIdentitySetPreference func(identity SecIdentityRef, name corefoundation.CFStringRef, keyUsage CSSM_KEYUSE) int32

// SecIdentitySetPreference sets the preferred identity for the specified name and key use.
//
// See: https://developer.apple.com/documentation/Security/SecIdentitySetPreference
func SecIdentitySetPreference(identity SecIdentityRef, name corefoundation.CFStringRef, keyUsage CSSM_KEYUSE) int32 {
	if _secIdentitySetPreference == nil {
		panic("Security: symbol SecIdentitySetPreference not loaded")
	}
	return _secIdentitySetPreference(identity, name, keyUsage)
}

var _secIdentitySetPreferred func(identity SecIdentityRef, name corefoundation.CFStringRef, keyUsage corefoundation.CFArrayRef) int32

// SecIdentitySetPreferred sets the identity that should be preferred for the specified name and key use.
//
// See: https://developer.apple.com/documentation/Security/SecIdentitySetPreferred(_:_:_:)
func SecIdentitySetPreferred(identity SecIdentityRef, name corefoundation.CFStringRef, keyUsage corefoundation.CFArrayRef) int32 {
	if _secIdentitySetPreferred == nil {
		panic("Security: symbol SecIdentitySetPreferred not loaded")
	}
	return _secIdentitySetPreferred(identity, name, keyUsage)
}

var _secIdentitySetSystemIdentity func(domain corefoundation.CFStringRef, idRef SecIdentityRef) int32

// SecIdentitySetSystemIdentity assigns the system identity to be associated with a specified domain.
//
// See: https://developer.apple.com/documentation/Security/SecIdentitySetSystemIdentity(_:_:)
func SecIdentitySetSystemIdentity(domain corefoundation.CFStringRef, idRef SecIdentityRef) int32 {
	if _secIdentitySetSystemIdentity == nil {
		panic("Security: symbol SecIdentitySetSystemIdentity not loaded")
	}
	return _secIdentitySetSystemIdentity(domain, idRef)
}

var _secItemAdd func(attributes corefoundation.CFDictionaryRef, result *corefoundation.CFTypeRef) int32

// SecItemAdd adds one or more items to a keychain.
//
// See: https://developer.apple.com/documentation/Security/SecItemAdd(_:_:)
func SecItemAdd(attributes corefoundation.CFDictionaryRef, result *corefoundation.CFTypeRef) int32 {
	if _secItemAdd == nil {
		panic("Security: symbol SecItemAdd not loaded")
	}
	return _secItemAdd(attributes, result)
}

var _secItemCopyMatching func(query corefoundation.CFDictionaryRef, result *corefoundation.CFTypeRef) int32

// SecItemCopyMatching returns one or more keychain items that match a search query, or copies attributes of specific keychain items.
//
// See: https://developer.apple.com/documentation/Security/SecItemCopyMatching(_:_:)
func SecItemCopyMatching(query corefoundation.CFDictionaryRef, result *corefoundation.CFTypeRef) int32 {
	if _secItemCopyMatching == nil {
		panic("Security: symbol SecItemCopyMatching not loaded")
	}
	return _secItemCopyMatching(query, result)
}

var _secItemDelete func(query corefoundation.CFDictionaryRef) int32

// SecItemDelete deletes items that match a search query.
//
// See: https://developer.apple.com/documentation/Security/SecItemDelete(_:)
func SecItemDelete(query corefoundation.CFDictionaryRef) int32 {
	if _secItemDelete == nil {
		panic("Security: symbol SecItemDelete not loaded")
	}
	return _secItemDelete(query)
}

var _secItemExport func(secItemOrArray corefoundation.CFTypeRef, outputFormat SecExternalFormat, flags SecItemImportExportFlags, keyParams *SecItemImportExportKeyParameters, exportedData *corefoundation.CFDataRef) int32

// SecItemExport exports one or more certificates, keys, or identities.
//
// See: https://developer.apple.com/documentation/Security/SecItemExport(_:_:_:_:_:)
func SecItemExport(secItemOrArray corefoundation.CFTypeRef, outputFormat SecExternalFormat, flags SecItemImportExportFlags, keyParams *SecItemImportExportKeyParameters, exportedData *corefoundation.CFDataRef) int32 {
	if _secItemExport == nil {
		panic("Security: symbol SecItemExport not loaded")
	}
	return _secItemExport(secItemOrArray, outputFormat, flags, keyParams, exportedData)
}

var _secItemImport func(importedData corefoundation.CFDataRef, fileNameOrExtension corefoundation.CFStringRef, inputFormat *SecExternalFormat, itemType *SecExternalItemType, flags SecItemImportExportFlags, keyParams *SecItemImportExportKeyParameters, importKeychain SecKeychainRef, outItems *corefoundation.CFArrayRef) int32

// SecItemImport imports one or more certificates, keys, or identities and optionally adds them to a keychain.
//
// See: https://developer.apple.com/documentation/Security/SecItemImport(_:_:_:_:_:_:_:_:)
func SecItemImport(importedData corefoundation.CFDataRef, fileNameOrExtension corefoundation.CFStringRef, inputFormat *SecExternalFormat, itemType *SecExternalItemType, flags SecItemImportExportFlags, keyParams *SecItemImportExportKeyParameters, importKeychain SecKeychainRef, outItems *corefoundation.CFArrayRef) int32 {
	if _secItemImport == nil {
		panic("Security: symbol SecItemImport not loaded")
	}
	return _secItemImport(importedData, fileNameOrExtension, inputFormat, itemType, flags, keyParams, importKeychain, outItems)
}

var _secItemUpdate func(query corefoundation.CFDictionaryRef, attributesToUpdate corefoundation.CFDictionaryRef) int32

// SecItemUpdate modifies items that match a search query.
//
// See: https://developer.apple.com/documentation/Security/SecItemUpdate(_:_:)
func SecItemUpdate(query corefoundation.CFDictionaryRef, attributesToUpdate corefoundation.CFDictionaryRef) int32 {
	if _secItemUpdate == nil {
		panic("Security: symbol SecItemUpdate not loaded")
	}
	return _secItemUpdate(query, attributesToUpdate)
}

var _secKeyCopyAttributes func(key SecKeyRef) corefoundation.CFDictionaryRef

// SecKeyCopyAttributes gets the attributes of a given key.
//
// See: https://developer.apple.com/documentation/Security/SecKeyCopyAttributes(_:)
func SecKeyCopyAttributes(key SecKeyRef) corefoundation.CFDictionaryRef {
	if _secKeyCopyAttributes == nil {
		panic("Security: symbol SecKeyCopyAttributes not loaded")
	}
	return _secKeyCopyAttributes(key)
}

var _secKeyCopyExternalRepresentation func(key SecKeyRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef

// SecKeyCopyExternalRepresentation returns an external representation of the given key suitable for the key’s type.
//
// See: https://developer.apple.com/documentation/Security/SecKeyCopyExternalRepresentation(_:_:)
func SecKeyCopyExternalRepresentation(key SecKeyRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef {
	if _secKeyCopyExternalRepresentation == nil {
		panic("Security: symbol SecKeyCopyExternalRepresentation not loaded")
	}
	return _secKeyCopyExternalRepresentation(key, err)
}

var _secKeyCopyKeyExchangeResult func(privateKey SecKeyRef, algorithm SecKeyAlgorithm, publicKey SecKeyRef, parameters corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef

// SecKeyCopyKeyExchangeResult performs the Diffie-Hellman style of key exchange with optional key-derivation steps.
//
// See: https://developer.apple.com/documentation/Security/SecKeyCopyKeyExchangeResult(_:_:_:_:_:)
func SecKeyCopyKeyExchangeResult(privateKey SecKeyRef, algorithm SecKeyAlgorithm, publicKey SecKeyRef, parameters corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef {
	if _secKeyCopyKeyExchangeResult == nil {
		panic("Security: symbol SecKeyCopyKeyExchangeResult not loaded")
	}
	return _secKeyCopyKeyExchangeResult(privateKey, algorithm, publicKey, parameters, err)
}

var _secKeyCopyPublicKey func(key SecKeyRef) SecKeyRef

// SecKeyCopyPublicKey gets the public key associated with the given private key.
//
// See: https://developer.apple.com/documentation/Security/SecKeyCopyPublicKey(_:)
func SecKeyCopyPublicKey(key SecKeyRef) SecKeyRef {
	if _secKeyCopyPublicKey == nil {
		panic("Security: symbol SecKeyCopyPublicKey not loaded")
	}
	return _secKeyCopyPublicKey(key)
}

var _secKeyCreateDecryptedData func(key SecKeyRef, algorithm SecKeyAlgorithm, ciphertext corefoundation.CFDataRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef

// SecKeyCreateDecryptedData decrypts a block of data using a private key and specified algorithm.
//
// See: https://developer.apple.com/documentation/Security/SecKeyCreateDecryptedData(_:_:_:_:)
func SecKeyCreateDecryptedData(key SecKeyRef, algorithm SecKeyAlgorithm, ciphertext corefoundation.CFDataRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef {
	if _secKeyCreateDecryptedData == nil {
		panic("Security: symbol SecKeyCreateDecryptedData not loaded")
	}
	return _secKeyCreateDecryptedData(key, algorithm, ciphertext, err)
}

var _secKeyCreateEncryptedData func(key SecKeyRef, algorithm SecKeyAlgorithm, plaintext corefoundation.CFDataRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef

// SecKeyCreateEncryptedData encrypts a block of data using a public key and specified algorithm.
//
// See: https://developer.apple.com/documentation/Security/SecKeyCreateEncryptedData(_:_:_:_:)
func SecKeyCreateEncryptedData(key SecKeyRef, algorithm SecKeyAlgorithm, plaintext corefoundation.CFDataRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef {
	if _secKeyCreateEncryptedData == nil {
		panic("Security: symbol SecKeyCreateEncryptedData not loaded")
	}
	return _secKeyCreateEncryptedData(key, algorithm, plaintext, err)
}

var _secKeyCreatePair func(keychainRef SecKeychainRef, algorithm CSSM_ALGORITHMS, keySizeInBits uint32, contextHandle CSSM_CC_HANDLE, publicKeyUsage CSSM_KEYUSE, publicKeyAttr uint32, privateKeyUsage CSSM_KEYUSE, privateKeyAttr uint32, initialAccess SecAccessRef, publicKey *SecKeyRef, privateKey *SecKeyRef) int32

// SecKeyCreatePair creates an asymmetric key pair and stores it in a keychain.
//
// See: https://developer.apple.com/documentation/Security/SecKeyCreatePair
func SecKeyCreatePair(keychainRef SecKeychainRef, algorithm CSSM_ALGORITHMS, keySizeInBits uint32, contextHandle CSSM_CC_HANDLE, publicKeyUsage CSSM_KEYUSE, publicKeyAttr uint32, privateKeyUsage CSSM_KEYUSE, privateKeyAttr uint32, initialAccess SecAccessRef, publicKey *SecKeyRef, privateKey *SecKeyRef) int32 {
	if _secKeyCreatePair == nil {
		panic("Security: symbol SecKeyCreatePair not loaded")
	}
	return _secKeyCreatePair(keychainRef, algorithm, keySizeInBits, contextHandle, publicKeyUsage, publicKeyAttr, privateKeyUsage, privateKeyAttr, initialAccess, publicKey, privateKey)
}

var _secKeyCreateRandomKey func(parameters corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) SecKeyRef

// SecKeyCreateRandomKey generates a new public-private key pair.
//
// See: https://developer.apple.com/documentation/Security/SecKeyCreateRandomKey(_:_:)
func SecKeyCreateRandomKey(parameters corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) SecKeyRef {
	if _secKeyCreateRandomKey == nil {
		panic("Security: symbol SecKeyCreateRandomKey not loaded")
	}
	return _secKeyCreateRandomKey(parameters, err)
}

var _secKeyCreateSignature func(key SecKeyRef, algorithm SecKeyAlgorithm, dataToSign corefoundation.CFDataRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef

// SecKeyCreateSignature creates the cryptographic signature for a block of data using a private key and specified algorithm.
//
// See: https://developer.apple.com/documentation/Security/SecKeyCreateSignature(_:_:_:_:)
func SecKeyCreateSignature(key SecKeyRef, algorithm SecKeyAlgorithm, dataToSign corefoundation.CFDataRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef {
	if _secKeyCreateSignature == nil {
		panic("Security: symbol SecKeyCreateSignature not loaded")
	}
	return _secKeyCreateSignature(key, algorithm, dataToSign, err)
}

var _secKeyCreateWithData func(keyData corefoundation.CFDataRef, attributes corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) SecKeyRef

// SecKeyCreateWithData restores a key from an external representation of that key.
//
// See: https://developer.apple.com/documentation/Security/SecKeyCreateWithData(_:_:_:)
func SecKeyCreateWithData(keyData corefoundation.CFDataRef, attributes corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) SecKeyRef {
	if _secKeyCreateWithData == nil {
		panic("Security: symbol SecKeyCreateWithData not loaded")
	}
	return _secKeyCreateWithData(keyData, attributes, err)
}

var _secKeyGenerate func(keychainRef SecKeychainRef, algorithm CSSM_ALGORITHMS, keySizeInBits uint32, contextHandle CSSM_CC_HANDLE, keyUsage CSSM_KEYUSE, keyAttr uint32, initialAccess SecAccessRef, keyRef *SecKeyRef) int32

// SecKeyGenerate creates a symmetric key and optionally stores it in a keychain.
//
// See: https://developer.apple.com/documentation/Security/SecKeyGenerate
func SecKeyGenerate(keychainRef SecKeychainRef, algorithm CSSM_ALGORITHMS, keySizeInBits uint32, contextHandle CSSM_CC_HANDLE, keyUsage CSSM_KEYUSE, keyAttr uint32, initialAccess SecAccessRef, keyRef *SecKeyRef) int32 {
	if _secKeyGenerate == nil {
		panic("Security: symbol SecKeyGenerate not loaded")
	}
	return _secKeyGenerate(keychainRef, algorithm, keySizeInBits, contextHandle, keyUsage, keyAttr, initialAccess, keyRef)
}

var _secKeyGetBlockSize func(key SecKeyRef) uintptr

// SecKeyGetBlockSize gets the block length associated with a cryptographic key.
//
// See: https://developer.apple.com/documentation/Security/SecKeyGetBlockSize(_:)
func SecKeyGetBlockSize(key SecKeyRef) uintptr {
	if _secKeyGetBlockSize == nil {
		panic("Security: symbol SecKeyGetBlockSize not loaded")
	}
	return _secKeyGetBlockSize(key)
}

var _secKeyGetCSPHandle func(keyRef SecKeyRef, cspHandle unsafe.Pointer) int32

// SecKeyGetCSPHandle returns the CSSM CSP handle for a key.
//
// See: https://developer.apple.com/documentation/Security/SecKeyGetCSPHandle
func SecKeyGetCSPHandle(keyRef SecKeyRef, cspHandle unsafe.Pointer) int32 {
	if _secKeyGetCSPHandle == nil {
		panic("Security: symbol SecKeyGetCSPHandle not loaded")
	}
	return _secKeyGetCSPHandle(keyRef, cspHandle)
}

var _secKeyGetCSSMKey func(key SecKeyRef, cssmKey unsafe.Pointer) int32

// SecKeyGetCSSMKey retrieves a pointer to the `CSSM_KEY` structure containing the key stored in a keychain item.
//
// See: https://developer.apple.com/documentation/Security/SecKeyGetCSSMKey
func SecKeyGetCSSMKey(key SecKeyRef, cssmKey unsafe.Pointer) int32 {
	if _secKeyGetCSSMKey == nil {
		panic("Security: symbol SecKeyGetCSSMKey not loaded")
	}
	return _secKeyGetCSSMKey(key, cssmKey)
}

var _secKeyGetCredentials func(keyRef SecKeyRef, operation CSSM_ACL_AUTHORIZATION_TAG, credentialType SecCredentialType, outCredentials unsafe.Pointer) int32

// SecKeyGetCredentials returns an access credential for a key.
//
// See: https://developer.apple.com/documentation/Security/SecKeyGetCredentials
func SecKeyGetCredentials(keyRef SecKeyRef, operation CSSM_ACL_AUTHORIZATION_TAG, credentialType SecCredentialType, outCredentials unsafe.Pointer) int32 {
	if _secKeyGetCredentials == nil {
		panic("Security: symbol SecKeyGetCredentials not loaded")
	}
	return _secKeyGetCredentials(keyRef, operation, credentialType, outCredentials)
}

var _secKeyGetTypeID func() uint

// SecKeyGetTypeID returns the unique identifier of the opaque type to which a key object belongs.
//
// See: https://developer.apple.com/documentation/Security/SecKeyGetTypeID()
func SecKeyGetTypeID() uint {
	if _secKeyGetTypeID == nil {
		panic("Security: symbol SecKeyGetTypeID not loaded")
	}
	return _secKeyGetTypeID()
}

var _secKeyIsAlgorithmSupported func(key SecKeyRef, operation SecKeyOperationType, algorithm SecKeyAlgorithm) bool

// SecKeyIsAlgorithmSupported returns a Boolean indicating whether a key is suitable for an operation using a certain algorithm.
//
// See: https://developer.apple.com/documentation/Security/SecKeyIsAlgorithmSupported(_:_:_:)
func SecKeyIsAlgorithmSupported(key SecKeyRef, operation SecKeyOperationType, algorithm SecKeyAlgorithm) bool {
	if _secKeyIsAlgorithmSupported == nil {
		panic("Security: symbol SecKeyIsAlgorithmSupported not loaded")
	}
	return _secKeyIsAlgorithmSupported(key, operation, algorithm)
}

var _secKeyVerifySignature func(key SecKeyRef, algorithm SecKeyAlgorithm, signedData corefoundation.CFDataRef, signature corefoundation.CFDataRef, err *corefoundation.CFErrorRef) bool

// SecKeyVerifySignature verifies the cryptographic signature of a block of data using a public key and specified algorithm.
//
// See: https://developer.apple.com/documentation/Security/SecKeyVerifySignature(_:_:_:_:_:)
func SecKeyVerifySignature(key SecKeyRef, algorithm SecKeyAlgorithm, signedData corefoundation.CFDataRef, signature corefoundation.CFDataRef, err *corefoundation.CFErrorRef) bool {
	if _secKeyVerifySignature == nil {
		panic("Security: symbol SecKeyVerifySignature not loaded")
	}
	return _secKeyVerifySignature(key, algorithm, signedData, signature, err)
}

var _secKeychainSearchGetTypeID func() uint

// SecKeychainSearchGetTypeID returns the unique identifier of the opaque type to which a keychain search object belongs.
//
// See: https://developer.apple.com/documentation/Security/SecKeychainSearchGetTypeID
func SecKeychainSearchGetTypeID() uint {
	if _secKeychainSearchGetTypeID == nil {
		panic("Security: symbol SecKeychainSearchGetTypeID not loaded")
	}
	return _secKeychainSearchGetTypeID()
}

var _secPKCS12Import func(pkcs12_data corefoundation.CFDataRef, options corefoundation.CFDictionaryRef, items *corefoundation.CFArrayRef) int32

// SecPKCS12Import returns the identities and certificates in a PKCS #12-formatted blob.
//
// See: https://developer.apple.com/documentation/Security/SecPKCS12Import(_:_:_:)
func SecPKCS12Import(pkcs12_data corefoundation.CFDataRef, options corefoundation.CFDictionaryRef, items *corefoundation.CFArrayRef) int32 {
	if _secPKCS12Import == nil {
		panic("Security: symbol SecPKCS12Import not loaded")
	}
	return _secPKCS12Import(pkcs12_data, options, items)
}

var _secPolicyCopyProperties func(policyRef SecPolicyRef) corefoundation.CFDictionaryRef

// SecPolicyCopyProperties returns a dictionary containing a policy’s properties.
//
// See: https://developer.apple.com/documentation/Security/SecPolicyCopyProperties(_:)
func SecPolicyCopyProperties(policyRef SecPolicyRef) corefoundation.CFDictionaryRef {
	if _secPolicyCopyProperties == nil {
		panic("Security: symbol SecPolicyCopyProperties not loaded")
	}
	return _secPolicyCopyProperties(policyRef)
}

var _secPolicyCreateBasicX509 func() SecPolicyRef

// SecPolicyCreateBasicX509 returns a policy object for the default X.509 policy.
//
// See: https://developer.apple.com/documentation/Security/SecPolicyCreateBasicX509()
func SecPolicyCreateBasicX509() SecPolicyRef {
	if _secPolicyCreateBasicX509 == nil {
		panic("Security: symbol SecPolicyCreateBasicX509 not loaded")
	}
	return _secPolicyCreateBasicX509()
}

var _secPolicyCreateRevocation func(revocationFlags uint64) SecPolicyRef

// SecPolicyCreateRevocation returns a policy object for checking revocation of certificates.
//
// See: https://developer.apple.com/documentation/Security/SecPolicyCreateRevocation(_:)
func SecPolicyCreateRevocation(revocationFlags uint64) SecPolicyRef {
	if _secPolicyCreateRevocation == nil {
		panic("Security: symbol SecPolicyCreateRevocation not loaded")
	}
	return _secPolicyCreateRevocation(revocationFlags)
}

var _secPolicyCreateSSL func(server bool, hostname corefoundation.CFStringRef) SecPolicyRef

// SecPolicyCreateSSL returns a policy object for evaluating SSL certificate chains.
//
// See: https://developer.apple.com/documentation/Security/SecPolicyCreateSSL(_:_:)
func SecPolicyCreateSSL(server bool, hostname corefoundation.CFStringRef) SecPolicyRef {
	if _secPolicyCreateSSL == nil {
		panic("Security: symbol SecPolicyCreateSSL not loaded")
	}
	return _secPolicyCreateSSL(server, hostname)
}

var _secPolicyCreateWithOID func(policyOID corefoundation.CFTypeRef) SecPolicyRef

// SecPolicyCreateWithOID returns a policy object for the specified policy type object identifier.
//
// See: https://developer.apple.com/documentation/Security/SecPolicyCreateWithOID
func SecPolicyCreateWithOID(policyOID corefoundation.CFTypeRef) SecPolicyRef {
	if _secPolicyCreateWithOID == nil {
		panic("Security: symbol SecPolicyCreateWithOID not loaded")
	}
	return _secPolicyCreateWithOID(policyOID)
}

var _secPolicyCreateWithProperties func(policyIdentifier corefoundation.CFTypeRef, properties corefoundation.CFDictionaryRef) SecPolicyRef

// SecPolicyCreateWithProperties returns a policy object based on an object identifier for the policy type.
//
// See: https://developer.apple.com/documentation/Security/SecPolicyCreateWithProperties(_:_:)
func SecPolicyCreateWithProperties(policyIdentifier corefoundation.CFTypeRef, properties corefoundation.CFDictionaryRef) SecPolicyRef {
	if _secPolicyCreateWithProperties == nil {
		panic("Security: symbol SecPolicyCreateWithProperties not loaded")
	}
	return _secPolicyCreateWithProperties(policyIdentifier, properties)
}

var _secPolicyGetOID func(policyRef SecPolicyRef, oid unsafe.Pointer) int32

// SecPolicyGetOID retrieves a policy’s object identifier.
//
// See: https://developer.apple.com/documentation/Security/SecPolicyGetOID
func SecPolicyGetOID(policyRef SecPolicyRef, oid unsafe.Pointer) int32 {
	if _secPolicyGetOID == nil {
		panic("Security: symbol SecPolicyGetOID not loaded")
	}
	return _secPolicyGetOID(policyRef, oid)
}

var _secPolicyGetTPHandle func(policyRef SecPolicyRef, tpHandle unsafe.Pointer) int32

// SecPolicyGetTPHandle retrieves the trust policy handle for a policy object.
//
// See: https://developer.apple.com/documentation/Security/SecPolicyGetTPHandle
func SecPolicyGetTPHandle(policyRef SecPolicyRef, tpHandle unsafe.Pointer) int32 {
	if _secPolicyGetTPHandle == nil {
		panic("Security: symbol SecPolicyGetTPHandle not loaded")
	}
	return _secPolicyGetTPHandle(policyRef, tpHandle)
}

var _secPolicyGetTypeID func() uint

// SecPolicyGetTypeID returns the unique identifier of the opaque type to which a policy object belongs.
//
// See: https://developer.apple.com/documentation/Security/SecPolicyGetTypeID()
func SecPolicyGetTypeID() uint {
	if _secPolicyGetTypeID == nil {
		panic("Security: symbol SecPolicyGetTypeID not loaded")
	}
	return _secPolicyGetTypeID()
}

var _secPolicyGetValue func(policyRef SecPolicyRef, value unsafe.Pointer) int32

// SecPolicyGetValue retrieves a policy’s value.
//
// See: https://developer.apple.com/documentation/Security/SecPolicyGetValue
func SecPolicyGetValue(policyRef SecPolicyRef, value unsafe.Pointer) int32 {
	if _secPolicyGetValue == nil {
		panic("Security: symbol SecPolicyGetValue not loaded")
	}
	return _secPolicyGetValue(policyRef, value)
}

var _secPolicySearchCopyNext func(searchRef SecPolicySearchRef, policyRef *SecPolicyRef) int32

// SecPolicySearchCopyNext retrieves a policy object for the next policy matching specified search criteria.
//
// See: https://developer.apple.com/documentation/Security/SecPolicySearchCopyNext
func SecPolicySearchCopyNext(searchRef SecPolicySearchRef, policyRef *SecPolicyRef) int32 {
	if _secPolicySearchCopyNext == nil {
		panic("Security: symbol SecPolicySearchCopyNext not loaded")
	}
	return _secPolicySearchCopyNext(searchRef, policyRef)
}

var _secPolicySearchCreate func(certType CSSM_CERT_TYPE, policyOID unsafe.Pointer, value unsafe.Pointer, searchRef *SecPolicySearchRef) int32

// SecPolicySearchCreate creates a search object for finding policies.
//
// See: https://developer.apple.com/documentation/Security/SecPolicySearchCreate
func SecPolicySearchCreate(certType CSSM_CERT_TYPE, policyOID unsafe.Pointer, value unsafe.Pointer, searchRef *SecPolicySearchRef) int32 {
	if _secPolicySearchCreate == nil {
		panic("Security: symbol SecPolicySearchCreate not loaded")
	}
	return _secPolicySearchCreate(certType, policyOID, value, searchRef)
}

var _secPolicySearchGetTypeID func() uint

// SecPolicySearchGetTypeID returns the unique identifier of the opaque type to which a [SecPolicySearch] object belongs.
//
// See: https://developer.apple.com/documentation/Security/SecPolicySearchGetTypeID
func SecPolicySearchGetTypeID() uint {
	if _secPolicySearchGetTypeID == nil {
		panic("Security: symbol SecPolicySearchGetTypeID not loaded")
	}
	return _secPolicySearchGetTypeID()
}

var _secPolicySetProperties func(policyRef SecPolicyRef, properties corefoundation.CFDictionaryRef) int32

// SecPolicySetProperties sets properties for a policy.
//
// See: https://developer.apple.com/documentation/Security/SecPolicySetProperties
func SecPolicySetProperties(policyRef SecPolicyRef, properties corefoundation.CFDictionaryRef) int32 {
	if _secPolicySetProperties == nil {
		panic("Security: symbol SecPolicySetProperties not loaded")
	}
	return _secPolicySetProperties(policyRef, properties)
}

var _secPolicySetValue func(policyRef SecPolicyRef, value unsafe.Pointer) int32

// SecPolicySetValue sets a policy’s value.
//
// See: https://developer.apple.com/documentation/Security/SecPolicySetValue
func SecPolicySetValue(policyRef SecPolicyRef, value unsafe.Pointer) int32 {
	if _secPolicySetValue == nil {
		panic("Security: symbol SecPolicySetValue not loaded")
	}
	return _secPolicySetValue(policyRef, value)
}

var _secRandomCopyBytes func(rnd SecRandomRef, count uintptr, bytes unsafe.Pointer) int

// SecRandomCopyBytes generates an array of cryptographically secure random bytes.
//
// See: https://developer.apple.com/documentation/Security/SecRandomCopyBytes(_:_:_:)
func SecRandomCopyBytes(rnd SecRandomRef, count uintptr, bytes unsafe.Pointer) int {
	if _secRandomCopyBytes == nil {
		panic("Security: symbol SecRandomCopyBytes not loaded")
	}
	return _secRandomCopyBytes(rnd, count, bytes)
}

var _secRequirementCopyData func(requirement SecRequirementRef, flags SecCSFlags, data *corefoundation.CFDataRef) int32

// SecRequirementCopyData extracts a binary form of a code requirement from a code requirement object.
//
// See: https://developer.apple.com/documentation/Security/SecRequirementCopyData(_:_:_:)
func SecRequirementCopyData(requirement SecRequirementRef, flags SecCSFlags, data *corefoundation.CFDataRef) int32 {
	if _secRequirementCopyData == nil {
		panic("Security: symbol SecRequirementCopyData not loaded")
	}
	return _secRequirementCopyData(requirement, flags, data)
}

var _secRequirementCopyString func(requirement SecRequirementRef, flags SecCSFlags, text *corefoundation.CFStringRef) int32

// SecRequirementCopyString converts a code requirement object into text form.
//
// See: https://developer.apple.com/documentation/Security/SecRequirementCopyString(_:_:_:)
func SecRequirementCopyString(requirement SecRequirementRef, flags SecCSFlags, text *corefoundation.CFStringRef) int32 {
	if _secRequirementCopyString == nil {
		panic("Security: symbol SecRequirementCopyString not loaded")
	}
	return _secRequirementCopyString(requirement, flags, text)
}

var _secRequirementCreateWithData func(data corefoundation.CFDataRef, flags SecCSFlags, requirement *SecRequirementRef) int32

// SecRequirementCreateWithData creates a code requirement object from the binary form of a code requirement.
//
// See: https://developer.apple.com/documentation/Security/SecRequirementCreateWithData(_:_:_:)
func SecRequirementCreateWithData(data corefoundation.CFDataRef, flags SecCSFlags, requirement *SecRequirementRef) int32 {
	if _secRequirementCreateWithData == nil {
		panic("Security: symbol SecRequirementCreateWithData not loaded")
	}
	return _secRequirementCreateWithData(data, flags, requirement)
}

var _secRequirementCreateWithString func(text corefoundation.CFStringRef, flags SecCSFlags, requirement *SecRequirementRef) int32

// SecRequirementCreateWithString creates a code requirement object by compiling a valid text representation of a code requirement.
//
// See: https://developer.apple.com/documentation/Security/SecRequirementCreateWithString(_:_:_:)
func SecRequirementCreateWithString(text corefoundation.CFStringRef, flags SecCSFlags, requirement *SecRequirementRef) int32 {
	if _secRequirementCreateWithString == nil {
		panic("Security: symbol SecRequirementCreateWithString not loaded")
	}
	return _secRequirementCreateWithString(text, flags, requirement)
}

var _secRequirementCreateWithStringAndErrors func(text corefoundation.CFStringRef, flags SecCSFlags, errors *corefoundation.CFErrorRef, requirement *SecRequirementRef) int32

// SecRequirementCreateWithStringAndErrors creates a code requirement object by compiling a valid text representation of a code requirement and returns detailed error information in the case of failure.
//
// See: https://developer.apple.com/documentation/Security/SecRequirementCreateWithStringAndErrors(_:_:_:_:)
func SecRequirementCreateWithStringAndErrors(text corefoundation.CFStringRef, flags SecCSFlags, errors *corefoundation.CFErrorRef, requirement *SecRequirementRef) int32 {
	if _secRequirementCreateWithStringAndErrors == nil {
		panic("Security: symbol SecRequirementCreateWithStringAndErrors not loaded")
	}
	return _secRequirementCreateWithStringAndErrors(text, flags, errors, requirement)
}

var _secRequirementGetTypeID func() uint

// SecRequirementGetTypeID returns the unique identifier of the opaque type to which a code requirement object belongs.
//
// See: https://developer.apple.com/documentation/Security/SecRequirementGetTypeID()
func SecRequirementGetTypeID() uint {
	if _secRequirementGetTypeID == nil {
		panic("Security: symbol SecRequirementGetTypeID not loaded")
	}
	return _secRequirementGetTypeID()
}

var _secStaticCodeCheckValidity func(staticCode SecStaticCodeRef, flags SecCSFlags, requirement SecRequirementRef) int32

// SecStaticCodeCheckValidity validates a static code object.
//
// See: https://developer.apple.com/documentation/Security/SecStaticCodeCheckValidity(_:_:_:)
func SecStaticCodeCheckValidity(staticCode SecStaticCodeRef, flags SecCSFlags, requirement SecRequirementRef) int32 {
	if _secStaticCodeCheckValidity == nil {
		panic("Security: symbol SecStaticCodeCheckValidity not loaded")
	}
	return _secStaticCodeCheckValidity(staticCode, flags, requirement)
}

var _secStaticCodeCheckValidityWithErrors func(staticCode SecStaticCodeRef, flags SecCSFlags, requirement SecRequirementRef, errors *corefoundation.CFErrorRef) int32

// SecStaticCodeCheckValidityWithErrors performs static validation of static signed code and returns detailed error information in the case of failure.
//
// See: https://developer.apple.com/documentation/Security/SecStaticCodeCheckValidityWithErrors(_:_:_:_:)
func SecStaticCodeCheckValidityWithErrors(staticCode SecStaticCodeRef, flags SecCSFlags, requirement SecRequirementRef, errors *corefoundation.CFErrorRef) int32 {
	if _secStaticCodeCheckValidityWithErrors == nil {
		panic("Security: symbol SecStaticCodeCheckValidityWithErrors not loaded")
	}
	return _secStaticCodeCheckValidityWithErrors(staticCode, flags, requirement, errors)
}

var _secStaticCodeCreateWithPath func(path corefoundation.CFURLRef, flags SecCSFlags, staticCode *SecStaticCodeRef) int32

// SecStaticCodeCreateWithPath creates a static code object representing the code at a specified file system path.
//
// See: https://developer.apple.com/documentation/Security/SecStaticCodeCreateWithPath(_:_:_:)
func SecStaticCodeCreateWithPath(path corefoundation.CFURLRef, flags SecCSFlags, staticCode *SecStaticCodeRef) int32 {
	if _secStaticCodeCreateWithPath == nil {
		panic("Security: symbol SecStaticCodeCreateWithPath not loaded")
	}
	return _secStaticCodeCreateWithPath(path, flags, staticCode)
}

var _secStaticCodeCreateWithPathAndAttributes func(path corefoundation.CFURLRef, flags SecCSFlags, attributes corefoundation.CFDictionaryRef, staticCode *SecStaticCodeRef) int32

// SecStaticCodeCreateWithPathAndAttributes creates a static code object representing the code at a specified file system path using an attributes dictionary.
//
// See: https://developer.apple.com/documentation/Security/SecStaticCodeCreateWithPathAndAttributes(_:_:_:_:)
func SecStaticCodeCreateWithPathAndAttributes(path corefoundation.CFURLRef, flags SecCSFlags, attributes corefoundation.CFDictionaryRef, staticCode *SecStaticCodeRef) int32 {
	if _secStaticCodeCreateWithPathAndAttributes == nil {
		panic("Security: symbol SecStaticCodeCreateWithPathAndAttributes not loaded")
	}
	return _secStaticCodeCreateWithPathAndAttributes(path, flags, attributes, staticCode)
}

var _secStaticCodeGetTypeID func() uint

// SecStaticCodeGetTypeID returns the unique identifier of the opaque type to which a static code object belongs.
//
// See: https://developer.apple.com/documentation/Security/SecStaticCodeGetTypeID()
func SecStaticCodeGetTypeID() uint {
	if _secStaticCodeGetTypeID == nil {
		panic("Security: symbol SecStaticCodeGetTypeID not loaded")
	}
	return _secStaticCodeGetTypeID()
}

var _secTaskCopySigningIdentifier func(task SecTaskRef, err *corefoundation.CFErrorRef) corefoundation.CFStringRef

// SecTaskCopySigningIdentifier returns the value of the code signing identifier.
//
// See: https://developer.apple.com/documentation/Security/SecTaskCopySigningIdentifier(_:_:)
func SecTaskCopySigningIdentifier(task SecTaskRef, err *corefoundation.CFErrorRef) corefoundation.CFStringRef {
	if _secTaskCopySigningIdentifier == nil {
		panic("Security: symbol SecTaskCopySigningIdentifier not loaded")
	}
	return _secTaskCopySigningIdentifier(task, err)
}

var _secTaskCopyValueForEntitlement func(task SecTaskRef, entitlement corefoundation.CFStringRef, err *corefoundation.CFErrorRef) corefoundation.CFTypeRef

// SecTaskCopyValueForEntitlement returns the value of a single entitlement for the represented task.
//
// See: https://developer.apple.com/documentation/Security/SecTaskCopyValueForEntitlement(_:_:_:)
func SecTaskCopyValueForEntitlement(task SecTaskRef, entitlement corefoundation.CFStringRef, err *corefoundation.CFErrorRef) corefoundation.CFTypeRef {
	if _secTaskCopyValueForEntitlement == nil {
		panic("Security: symbol SecTaskCopyValueForEntitlement not loaded")
	}
	return _secTaskCopyValueForEntitlement(task, entitlement, err)
}

var _secTaskCopyValuesForEntitlements func(task SecTaskRef, entitlements corefoundation.CFArrayRef, err *corefoundation.CFErrorRef) corefoundation.CFDictionaryRef

// SecTaskCopyValuesForEntitlements returns the values of multiple entitlements for the represented task.
//
// See: https://developer.apple.com/documentation/Security/SecTaskCopyValuesForEntitlements(_:_:_:)
func SecTaskCopyValuesForEntitlements(task SecTaskRef, entitlements corefoundation.CFArrayRef, err *corefoundation.CFErrorRef) corefoundation.CFDictionaryRef {
	if _secTaskCopyValuesForEntitlements == nil {
		panic("Security: symbol SecTaskCopyValuesForEntitlements not loaded")
	}
	return _secTaskCopyValuesForEntitlements(task, entitlements, err)
}

var _secTaskCreateFromSelf func(allocator corefoundation.CFAllocatorRef) SecTaskRef

// SecTaskCreateFromSelf creates a task object for the current task.
//
// See: https://developer.apple.com/documentation/Security/SecTaskCreateFromSelf(_:)
func SecTaskCreateFromSelf(allocator corefoundation.CFAllocatorRef) SecTaskRef {
	if _secTaskCreateFromSelf == nil {
		panic("Security: symbol SecTaskCreateFromSelf not loaded")
	}
	return _secTaskCreateFromSelf(allocator)
}

var _secTaskCreateWithAuditToken func(allocator corefoundation.CFAllocatorRef, token uintptr) SecTaskRef

// SecTaskCreateWithAuditToken creates a task object for the task that sent the Mach message represented by the audit token.
//
// See: https://developer.apple.com/documentation/Security/SecTaskCreateWithAuditToken(_:_:)
func SecTaskCreateWithAuditToken(allocator corefoundation.CFAllocatorRef, token uintptr) SecTaskRef {
	if _secTaskCreateWithAuditToken == nil {
		panic("Security: symbol SecTaskCreateWithAuditToken not loaded")
	}
	return _secTaskCreateWithAuditToken(allocator, token)
}

var _secTaskGetTypeID func() uint

// SecTaskGetTypeID returns the unique identifier of the opaque type to which a task object belongs.
//
// See: https://developer.apple.com/documentation/Security/SecTaskGetTypeID()
func SecTaskGetTypeID() uint {
	if _secTaskGetTypeID == nil {
		panic("Security: symbol SecTaskGetTypeID not loaded")
	}
	return _secTaskGetTypeID()
}

var _secTrustCopyAnchorCertificates func(anchors *corefoundation.CFArrayRef) int32

// SecTrustCopyAnchorCertificates retrieves the anchor (root) certificates stored by macOS.
//
// See: https://developer.apple.com/documentation/Security/SecTrustCopyAnchorCertificates(_:)
func SecTrustCopyAnchorCertificates(anchors *corefoundation.CFArrayRef) int32 {
	if _secTrustCopyAnchorCertificates == nil {
		panic("Security: symbol SecTrustCopyAnchorCertificates not loaded")
	}
	return _secTrustCopyAnchorCertificates(anchors)
}

var _secTrustCopyCertificateChain func(trust SecTrustRef) corefoundation.CFArrayRef

// SecTrustCopyCertificateChain.
//
// See: https://developer.apple.com/documentation/Security/SecTrustCopyCertificateChain(_:)
func SecTrustCopyCertificateChain(trust SecTrustRef) corefoundation.CFArrayRef {
	if _secTrustCopyCertificateChain == nil {
		panic("Security: symbol SecTrustCopyCertificateChain not loaded")
	}
	return _secTrustCopyCertificateChain(trust)
}

var _secTrustCopyCustomAnchorCertificates func(trust SecTrustRef, anchors *corefoundation.CFArrayRef) int32

// SecTrustCopyCustomAnchorCertificates retrieves the custom anchor certificates, if any, used by a given trust.
//
// See: https://developer.apple.com/documentation/Security/SecTrustCopyCustomAnchorCertificates(_:_:)
func SecTrustCopyCustomAnchorCertificates(trust SecTrustRef, anchors *corefoundation.CFArrayRef) int32 {
	if _secTrustCopyCustomAnchorCertificates == nil {
		panic("Security: symbol SecTrustCopyCustomAnchorCertificates not loaded")
	}
	return _secTrustCopyCustomAnchorCertificates(trust, anchors)
}

var _secTrustCopyExceptions func(trust SecTrustRef) corefoundation.CFDataRef

// SecTrustCopyExceptions returns an opaque cookie containing exceptions to trust policies that will allow future evaluations of the current certificate to succeed.
//
// See: https://developer.apple.com/documentation/Security/SecTrustCopyExceptions(_:)
func SecTrustCopyExceptions(trust SecTrustRef) corefoundation.CFDataRef {
	if _secTrustCopyExceptions == nil {
		panic("Security: symbol SecTrustCopyExceptions not loaded")
	}
	return _secTrustCopyExceptions(trust)
}

var _secTrustCopyKey func(trust SecTrustRef) SecKeyRef

// SecTrustCopyKey.
//
// See: https://developer.apple.com/documentation/Security/SecTrustCopyKey(_:)
func SecTrustCopyKey(trust SecTrustRef) SecKeyRef {
	if _secTrustCopyKey == nil {
		panic("Security: symbol SecTrustCopyKey not loaded")
	}
	return _secTrustCopyKey(trust)
}

var _secTrustCopyPolicies func(trust SecTrustRef, policies *corefoundation.CFArrayRef) int32

// SecTrustCopyPolicies retrieves the policies used by a given trust management object.
//
// See: https://developer.apple.com/documentation/Security/SecTrustCopyPolicies(_:_:)
func SecTrustCopyPolicies(trust SecTrustRef, policies *corefoundation.CFArrayRef) int32 {
	if _secTrustCopyPolicies == nil {
		panic("Security: symbol SecTrustCopyPolicies not loaded")
	}
	return _secTrustCopyPolicies(trust, policies)
}

var _secTrustCopyProperties func(trust SecTrustRef) corefoundation.CFArrayRef

// SecTrustCopyProperties returns an array containing the properties of a trust object.
//
// See: https://developer.apple.com/documentation/Security/SecTrustCopyProperties(_:)
func SecTrustCopyProperties(trust SecTrustRef) corefoundation.CFArrayRef {
	if _secTrustCopyProperties == nil {
		panic("Security: symbol SecTrustCopyProperties not loaded")
	}
	return _secTrustCopyProperties(trust)
}

var _secTrustCopyPublicKey func(trust SecTrustRef) SecKeyRef

// SecTrustCopyPublicKey returns the public key for a leaf certificate after it has been evaluated.
//
// See: https://developer.apple.com/documentation/Security/SecTrustCopyPublicKey(_:)
func SecTrustCopyPublicKey(trust SecTrustRef) SecKeyRef {
	if _secTrustCopyPublicKey == nil {
		panic("Security: symbol SecTrustCopyPublicKey not loaded")
	}
	return _secTrustCopyPublicKey(trust)
}

var _secTrustCopyResult func(trust SecTrustRef) corefoundation.CFDictionaryRef

// SecTrustCopyResult returns a dictionary containing information about an evaluated trust.
//
// See: https://developer.apple.com/documentation/Security/SecTrustCopyResult(_:)
func SecTrustCopyResult(trust SecTrustRef) corefoundation.CFDictionaryRef {
	if _secTrustCopyResult == nil {
		panic("Security: symbol SecTrustCopyResult not loaded")
	}
	return _secTrustCopyResult(trust)
}

var _secTrustCreateWithCertificates func(certificates corefoundation.CFTypeRef, policies corefoundation.CFTypeRef, trust *SecTrustRef) int32

// SecTrustCreateWithCertificates creates a trust management object based on certificates and policies.
//
// See: https://developer.apple.com/documentation/Security/SecTrustCreateWithCertificates(_:_:_:)
func SecTrustCreateWithCertificates(certificates corefoundation.CFTypeRef, policies corefoundation.CFTypeRef, trust *SecTrustRef) int32 {
	if _secTrustCreateWithCertificates == nil {
		panic("Security: symbol SecTrustCreateWithCertificates not loaded")
	}
	return _secTrustCreateWithCertificates(certificates, policies, trust)
}

var _secTrustEvaluateAsyncWithError func(trust SecTrustRef, queue uintptr, result SecTrustWithErrorCallback) int32

// SecTrustEvaluateAsyncWithError evaluates a trust object asynchronously on the specified dispatch queue.
//
// See: https://developer.apple.com/documentation/Security/SecTrustEvaluateAsyncWithError(_:_:_:)
func SecTrustEvaluateAsyncWithError(trust SecTrustRef, queue dispatch.Queue, result SecTrustWithErrorCallback) int32 {
	if _secTrustEvaluateAsyncWithError == nil {
		panic("Security: symbol SecTrustEvaluateAsyncWithError not loaded")
	}
	return _secTrustEvaluateAsyncWithError(trust, uintptr(queue.Handle()), result)
}

var _secTrustEvaluateWithError func(trust SecTrustRef, err *corefoundation.CFErrorRef) bool

// SecTrustEvaluateWithError evaluates trust for the specified certificate and policies.
//
// See: https://developer.apple.com/documentation/Security/SecTrustEvaluateWithError(_:_:)
func SecTrustEvaluateWithError(trust SecTrustRef, err *corefoundation.CFErrorRef) bool {
	if _secTrustEvaluateWithError == nil {
		panic("Security: symbol SecTrustEvaluateWithError not loaded")
	}
	return _secTrustEvaluateWithError(trust, err)
}

var _secTrustGetCertificateAtIndex func(trust SecTrustRef, ix int) SecCertificateRef

// SecTrustGetCertificateAtIndex returns a specific certificate from the certificate chain used to evaluate trust.
//
// See: https://developer.apple.com/documentation/Security/SecTrustGetCertificateAtIndex(_:_:)
func SecTrustGetCertificateAtIndex(trust SecTrustRef, ix int) SecCertificateRef {
	if _secTrustGetCertificateAtIndex == nil {
		panic("Security: symbol SecTrustGetCertificateAtIndex not loaded")
	}
	return _secTrustGetCertificateAtIndex(trust, ix)
}

var _secTrustGetCertificateCount func(trust SecTrustRef) int

// SecTrustGetCertificateCount returns the number of certificates in an evaluated certificate chain.
//
// See: https://developer.apple.com/documentation/Security/SecTrustGetCertificateCount(_:)
func SecTrustGetCertificateCount(trust SecTrustRef) int {
	if _secTrustGetCertificateCount == nil {
		panic("Security: symbol SecTrustGetCertificateCount not loaded")
	}
	return _secTrustGetCertificateCount(trust)
}

var _secTrustGetCssmResult func(trust SecTrustRef, result unsafe.Pointer) int32

// SecTrustGetCssmResult retrieves the CSSM trust result.
//
// See: https://developer.apple.com/documentation/Security/SecTrustGetCssmResult
func SecTrustGetCssmResult(trust SecTrustRef, result unsafe.Pointer) int32 {
	if _secTrustGetCssmResult == nil {
		panic("Security: symbol SecTrustGetCssmResult not loaded")
	}
	return _secTrustGetCssmResult(trust, result)
}

var _secTrustGetCssmResultCode func(trust SecTrustRef, resultCode *int32) int32

// SecTrustGetCssmResultCode retrieves the CSSM result code from the most recent trust evaluation for a trust management object.
//
// See: https://developer.apple.com/documentation/Security/SecTrustGetCssmResultCode
func SecTrustGetCssmResultCode(trust SecTrustRef, resultCode *int32) int32 {
	if _secTrustGetCssmResultCode == nil {
		panic("Security: symbol SecTrustGetCssmResultCode not loaded")
	}
	return _secTrustGetCssmResultCode(trust, resultCode)
}

var _secTrustGetNetworkFetchAllowed func(trust SecTrustRef, allowFetch *bool) int32

// SecTrustGetNetworkFetchAllowed indicates whether a trust evaluation is permitted to fetch missing intermediate certificates from the network.
//
// See: https://developer.apple.com/documentation/Security/SecTrustGetNetworkFetchAllowed(_:_:)
func SecTrustGetNetworkFetchAllowed(trust SecTrustRef, allowFetch *bool) int32 {
	if _secTrustGetNetworkFetchAllowed == nil {
		panic("Security: symbol SecTrustGetNetworkFetchAllowed not loaded")
	}
	return _secTrustGetNetworkFetchAllowed(trust, allowFetch)
}

var _secTrustGetResult func(trustRef SecTrustRef, result *SecTrustResultType, certChain *corefoundation.CFArrayRef, statusChain unsafe.Pointer) int32

// SecTrustGetResult retrieves details on the outcome of a call to the function [SecTrustEvaluate].
//
// See: https://developer.apple.com/documentation/Security/SecTrustGetResult
func SecTrustGetResult(trustRef SecTrustRef, result *SecTrustResultType, certChain *corefoundation.CFArrayRef, statusChain unsafe.Pointer) int32 {
	if _secTrustGetResult == nil {
		panic("Security: symbol SecTrustGetResult not loaded")
	}
	return _secTrustGetResult(trustRef, result, certChain, statusChain)
}

var _secTrustGetTPHandle func(trust SecTrustRef, handle unsafe.Pointer) int32

// SecTrustGetTPHandle retrieves the trust policy handle.
//
// See: https://developer.apple.com/documentation/Security/SecTrustGetTPHandle
func SecTrustGetTPHandle(trust SecTrustRef, handle unsafe.Pointer) int32 {
	if _secTrustGetTPHandle == nil {
		panic("Security: symbol SecTrustGetTPHandle not loaded")
	}
	return _secTrustGetTPHandle(trust, handle)
}

var _secTrustGetTrustResult func(trust SecTrustRef, result *SecTrustResultType) int32

// SecTrustGetTrustResult returns the result code from the most recent trust evaluation.
//
// See: https://developer.apple.com/documentation/Security/SecTrustGetTrustResult(_:_:)
func SecTrustGetTrustResult(trust SecTrustRef, result *SecTrustResultType) int32 {
	if _secTrustGetTrustResult == nil {
		panic("Security: symbol SecTrustGetTrustResult not loaded")
	}
	return _secTrustGetTrustResult(trust, result)
}

var _secTrustGetTypeID func() uint

// SecTrustGetTypeID returns the unique identifier of the opaque type to which a trust object belongs.
//
// See: https://developer.apple.com/documentation/Security/SecTrustGetTypeID()
func SecTrustGetTypeID() uint {
	if _secTrustGetTypeID == nil {
		panic("Security: symbol SecTrustGetTypeID not loaded")
	}
	return _secTrustGetTypeID()
}

var _secTrustGetVerifyTime func(trust SecTrustRef) corefoundation.CFAbsoluteTime

// SecTrustGetVerifyTime gets the absolute time against which the certificates in a trust management object are verified.
//
// See: https://developer.apple.com/documentation/Security/SecTrustGetVerifyTime(_:)
func SecTrustGetVerifyTime(trust SecTrustRef) corefoundation.CFAbsoluteTime {
	if _secTrustGetVerifyTime == nil {
		panic("Security: symbol SecTrustGetVerifyTime not loaded")
	}
	return _secTrustGetVerifyTime(trust)
}

var _secTrustSetAnchorCertificates func(trust SecTrustRef, anchorCertificates corefoundation.CFArrayRef) int32

// SecTrustSetAnchorCertificates sets the anchor certificates used when evaluating a trust management object.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSetAnchorCertificates(_:_:)
func SecTrustSetAnchorCertificates(trust SecTrustRef, anchorCertificates corefoundation.CFArrayRef) int32 {
	if _secTrustSetAnchorCertificates == nil {
		panic("Security: symbol SecTrustSetAnchorCertificates not loaded")
	}
	return _secTrustSetAnchorCertificates(trust, anchorCertificates)
}

var _secTrustSetAnchorCertificatesOnly func(trust SecTrustRef, anchorCertificatesOnly bool) int32

// SecTrustSetAnchorCertificatesOnly reenables trusting built-in anchor certificates.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSetAnchorCertificatesOnly(_:_:)
func SecTrustSetAnchorCertificatesOnly(trust SecTrustRef, anchorCertificatesOnly bool) int32 {
	if _secTrustSetAnchorCertificatesOnly == nil {
		panic("Security: symbol SecTrustSetAnchorCertificatesOnly not loaded")
	}
	return _secTrustSetAnchorCertificatesOnly(trust, anchorCertificatesOnly)
}

var _secTrustSetExceptions func(trust SecTrustRef, exceptions corefoundation.CFDataRef) bool

// SecTrustSetExceptions sets a list of exceptions that should be ignored when the certificate is evaluated.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSetExceptions(_:_:)
func SecTrustSetExceptions(trust SecTrustRef, exceptions corefoundation.CFDataRef) bool {
	if _secTrustSetExceptions == nil {
		panic("Security: symbol SecTrustSetExceptions not loaded")
	}
	return _secTrustSetExceptions(trust, exceptions)
}

var _secTrustSetKeychains func(trust SecTrustRef, keychainOrArray corefoundation.CFTypeRef) int32

// SecTrustSetKeychains sets the keychains searched for intermediate certificates when evaluating a trust management object.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSetKeychains(_:_:)
func SecTrustSetKeychains(trust SecTrustRef, keychainOrArray corefoundation.CFTypeRef) int32 {
	if _secTrustSetKeychains == nil {
		panic("Security: symbol SecTrustSetKeychains not loaded")
	}
	return _secTrustSetKeychains(trust, keychainOrArray)
}

var _secTrustSetNetworkFetchAllowed func(trust SecTrustRef, allowFetch bool) int32

// SecTrustSetNetworkFetchAllowed specifies whether a trust evaluation is permitted to fetch missing intermediate certificates from the network.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSetNetworkFetchAllowed(_:_:)
func SecTrustSetNetworkFetchAllowed(trust SecTrustRef, allowFetch bool) int32 {
	if _secTrustSetNetworkFetchAllowed == nil {
		panic("Security: symbol SecTrustSetNetworkFetchAllowed not loaded")
	}
	return _secTrustSetNetworkFetchAllowed(trust, allowFetch)
}

var _secTrustSetOCSPResponse func(trust SecTrustRef, responseData corefoundation.CFTypeRef) int32

// SecTrustSetOCSPResponse attaches Online Certificate Status Protocol (OSCP) response data to a trust object.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSetOCSPResponse(_:_:)
func SecTrustSetOCSPResponse(trust SecTrustRef, responseData corefoundation.CFTypeRef) int32 {
	if _secTrustSetOCSPResponse == nil {
		panic("Security: symbol SecTrustSetOCSPResponse not loaded")
	}
	return _secTrustSetOCSPResponse(trust, responseData)
}

var _secTrustSetOptions func(trustRef SecTrustRef, options SecTrustOptionFlags) int32

// SecTrustSetOptions sets option flags for customizing evaluation of a trust object.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSetOptions(_:_:)
func SecTrustSetOptions(trustRef SecTrustRef, options SecTrustOptionFlags) int32 {
	if _secTrustSetOptions == nil {
		panic("Security: symbol SecTrustSetOptions not loaded")
	}
	return _secTrustSetOptions(trustRef, options)
}

var _secTrustSetParameters func(trustRef SecTrustRef, action CSSM_TP_ACTION, actionData corefoundation.CFDataRef) int32

// SecTrustSetParameters sets the action and action data for a trust management object.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSetParameters
func SecTrustSetParameters(trustRef SecTrustRef, action CSSM_TP_ACTION, actionData corefoundation.CFDataRef) int32 {
	if _secTrustSetParameters == nil {
		panic("Security: symbol SecTrustSetParameters not loaded")
	}
	return _secTrustSetParameters(trustRef, action, actionData)
}

var _secTrustSetPolicies func(trust SecTrustRef, policies corefoundation.CFTypeRef) int32

// SecTrustSetPolicies sets the policies to use in an evaluation.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSetPolicies(_:_:)
func SecTrustSetPolicies(trust SecTrustRef, policies corefoundation.CFTypeRef) int32 {
	if _secTrustSetPolicies == nil {
		panic("Security: symbol SecTrustSetPolicies not loaded")
	}
	return _secTrustSetPolicies(trust, policies)
}

var _secTrustSetSignedCertificateTimestamps func(trust SecTrustRef, sctArray corefoundation.CFArrayRef) int32

// SecTrustSetSignedCertificateTimestamps attaches signed certificate timestamp data to a trust object.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSetSignedCertificateTimestamps(_:_:)
func SecTrustSetSignedCertificateTimestamps(trust SecTrustRef, sctArray corefoundation.CFArrayRef) int32 {
	if _secTrustSetSignedCertificateTimestamps == nil {
		panic("Security: symbol SecTrustSetSignedCertificateTimestamps not loaded")
	}
	return _secTrustSetSignedCertificateTimestamps(trust, sctArray)
}

var _secTrustSetVerifyDate func(trust SecTrustRef, verifyDate corefoundation.CFDateRef) int32

// SecTrustSetVerifyDate sets the date and time against which the certificates in a trust management object are verified.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSetVerifyDate(_:_:)
func SecTrustSetVerifyDate(trust SecTrustRef, verifyDate corefoundation.CFDateRef) int32 {
	if _secTrustSetVerifyDate == nil {
		panic("Security: symbol SecTrustSetVerifyDate not loaded")
	}
	return _secTrustSetVerifyDate(trust, verifyDate)
}

var _secTrustSettingsCopyCertificates func(domain SecTrustSettingsDomain, certArray *corefoundation.CFArrayRef) int32

// SecTrustSettingsCopyCertificates obtains an array of all certificates that have trust settings in a specific trust settings domain.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSettingsCopyCertificates(_:_:)
func SecTrustSettingsCopyCertificates(domain SecTrustSettingsDomain, certArray *corefoundation.CFArrayRef) int32 {
	if _secTrustSettingsCopyCertificates == nil {
		panic("Security: symbol SecTrustSettingsCopyCertificates not loaded")
	}
	return _secTrustSettingsCopyCertificates(domain, certArray)
}

var _secTrustSettingsCopyModificationDate func(certRef SecCertificateRef, domain SecTrustSettingsDomain, modificationDate *corefoundation.CFDateRef) int32

// SecTrustSettingsCopyModificationDate obtains the date and time at which a certificate’s trust settings were last modified.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSettingsCopyModificationDate(_:_:_:)
func SecTrustSettingsCopyModificationDate(certRef SecCertificateRef, domain SecTrustSettingsDomain, modificationDate *corefoundation.CFDateRef) int32 {
	if _secTrustSettingsCopyModificationDate == nil {
		panic("Security: symbol SecTrustSettingsCopyModificationDate not loaded")
	}
	return _secTrustSettingsCopyModificationDate(certRef, domain, modificationDate)
}

var _secTrustSettingsCopyTrustSettings func(certRef SecCertificateRef, domain SecTrustSettingsDomain, trustSettings *corefoundation.CFArrayRef) int32

// SecTrustSettingsCopyTrustSettings obtains the trust settings for a certificate.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSettingsCopyTrustSettings(_:_:_:)
func SecTrustSettingsCopyTrustSettings(certRef SecCertificateRef, domain SecTrustSettingsDomain, trustSettings *corefoundation.CFArrayRef) int32 {
	if _secTrustSettingsCopyTrustSettings == nil {
		panic("Security: symbol SecTrustSettingsCopyTrustSettings not loaded")
	}
	return _secTrustSettingsCopyTrustSettings(certRef, domain, trustSettings)
}

var _secTrustSettingsCreateExternalRepresentation func(domain SecTrustSettingsDomain, trustSettings *corefoundation.CFDataRef) int32

// SecTrustSettingsCreateExternalRepresentation obtains an external, portable representation of the specified domain’s trust settings.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSettingsCreateExternalRepresentation(_:_:)
func SecTrustSettingsCreateExternalRepresentation(domain SecTrustSettingsDomain, trustSettings *corefoundation.CFDataRef) int32 {
	if _secTrustSettingsCreateExternalRepresentation == nil {
		panic("Security: symbol SecTrustSettingsCreateExternalRepresentation not loaded")
	}
	return _secTrustSettingsCreateExternalRepresentation(domain, trustSettings)
}

var _secTrustSettingsImportExternalRepresentation func(domain SecTrustSettingsDomain, trustSettings corefoundation.CFDataRef) int32

// SecTrustSettingsImportExternalRepresentation imports trust settings into a trust domain.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSettingsImportExternalRepresentation(_:_:)
func SecTrustSettingsImportExternalRepresentation(domain SecTrustSettingsDomain, trustSettings corefoundation.CFDataRef) int32 {
	if _secTrustSettingsImportExternalRepresentation == nil {
		panic("Security: symbol SecTrustSettingsImportExternalRepresentation not loaded")
	}
	return _secTrustSettingsImportExternalRepresentation(domain, trustSettings)
}

var _secTrustSettingsRemoveTrustSettings func(certRef SecCertificateRef, domain SecTrustSettingsDomain) int32

// SecTrustSettingsRemoveTrustSettings deletes the trust settings for a certificate.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSettingsRemoveTrustSettings(_:_:)
func SecTrustSettingsRemoveTrustSettings(certRef SecCertificateRef, domain SecTrustSettingsDomain) int32 {
	if _secTrustSettingsRemoveTrustSettings == nil {
		panic("Security: symbol SecTrustSettingsRemoveTrustSettings not loaded")
	}
	return _secTrustSettingsRemoveTrustSettings(certRef, domain)
}

var _secTrustSettingsSetTrustSettings func(certRef SecCertificateRef, domain SecTrustSettingsDomain, trustSettingsDictOrArray corefoundation.CFTypeRef) int32

// SecTrustSettingsSetTrustSettings specifies trust settings for a certificate.
//
// See: https://developer.apple.com/documentation/Security/SecTrustSettingsSetTrustSettings(_:_:_:)
func SecTrustSettingsSetTrustSettings(certRef SecCertificateRef, domain SecTrustSettingsDomain, trustSettingsDictOrArray corefoundation.CFTypeRef) int32 {
	if _secTrustSettingsSetTrustSettings == nil {
		panic("Security: symbol SecTrustSettingsSetTrustSettings not loaded")
	}
	return _secTrustSettingsSetTrustSettings(certRef, domain, trustSettingsDictOrArray)
}

var _secureDownloadCopyCreationDate func(downloadRef SecureDownloadRef, date *corefoundation.CFDateRef) int32

// SecureDownloadCopyCreationDate returns download ticket’s creation date.
//
// See: https://developer.apple.com/documentation/Security/SecureDownloadCopyCreationDate
func SecureDownloadCopyCreationDate(downloadRef SecureDownloadRef, date *corefoundation.CFDateRef) int32 {
	if _secureDownloadCopyCreationDate == nil {
		panic("Security: symbol SecureDownloadCopyCreationDate not loaded")
	}
	return _secureDownloadCopyCreationDate(downloadRef, date)
}

var _secureDownloadCopyName func(downloadRef SecureDownloadRef, name *corefoundation.CFStringRef) int32

// SecureDownloadCopyName returns the printable name of the download ticket.
//
// See: https://developer.apple.com/documentation/Security/SecureDownloadCopyName
func SecureDownloadCopyName(downloadRef SecureDownloadRef, name *corefoundation.CFStringRef) int32 {
	if _secureDownloadCopyName == nil {
		panic("Security: symbol SecureDownloadCopyName not loaded")
	}
	return _secureDownloadCopyName(downloadRef, name)
}

var _secureDownloadCopyTicketLocation func(url corefoundation.CFURLRef, ticketLocation *corefoundation.CFURLRef) int32

// SecureDownloadCopyTicketLocation copies the ticket location from a secure download URL.
//
// See: https://developer.apple.com/documentation/Security/SecureDownloadCopyTicketLocation
func SecureDownloadCopyTicketLocation(url corefoundation.CFURLRef, ticketLocation *corefoundation.CFURLRef) int32 {
	if _secureDownloadCopyTicketLocation == nil {
		panic("Security: symbol SecureDownloadCopyTicketLocation not loaded")
	}
	return _secureDownloadCopyTicketLocation(url, ticketLocation)
}

var _secureDownloadCopyURLs func(downloadRef SecureDownloadRef, urls *corefoundation.CFArrayRef) int32

// SecureDownloadCopyURLs returns a list of URLs from which the data can be downloaded.
//
// See: https://developer.apple.com/documentation/Security/SecureDownloadCopyURLs
func SecureDownloadCopyURLs(downloadRef SecureDownloadRef, urls *corefoundation.CFArrayRef) int32 {
	if _secureDownloadCopyURLs == nil {
		panic("Security: symbol SecureDownloadCopyURLs not loaded")
	}
	return _secureDownloadCopyURLs(downloadRef, urls)
}

var _secureDownloadCreateWithTicket func(ticket corefoundation.CFDataRef, setup unsafe.Pointer, setupContext unsafe.Pointer, evaluate unsafe.Pointer, evaluateContext unsafe.Pointer, downloadRef *SecureDownloadRef) int32

// SecureDownloadCreateWithTicket creates a secure download object for use during the download process.
//
// See: https://developer.apple.com/documentation/Security/SecureDownloadCreateWithTicket
func SecureDownloadCreateWithTicket(ticket corefoundation.CFDataRef, setup unsafe.Pointer, setupContext unsafe.Pointer, evaluate unsafe.Pointer, evaluateContext unsafe.Pointer, downloadRef *SecureDownloadRef) int32 {
	if _secureDownloadCreateWithTicket == nil {
		panic("Security: symbol SecureDownloadCreateWithTicket not loaded")
	}
	return _secureDownloadCreateWithTicket(ticket, setup, setupContext, evaluate, evaluateContext, downloadRef)
}

var _secureDownloadFinished func(downloadRef SecureDownloadRef) int32

// SecureDownloadFinished concludes the secure download process.
//
// See: https://developer.apple.com/documentation/Security/SecureDownloadFinished
func SecureDownloadFinished(downloadRef SecureDownloadRef) int32 {
	if _secureDownloadFinished == nil {
		panic("Security: symbol SecureDownloadFinished not loaded")
	}
	return _secureDownloadFinished(downloadRef)
}

var _secureDownloadGetDownloadSize func(downloadRef SecureDownloadRef, downloadSize *int64) int32

// SecureDownloadGetDownloadSize returns the size of the expected download.
//
// See: https://developer.apple.com/documentation/Security/SecureDownloadGetDownloadSize
func SecureDownloadGetDownloadSize(downloadRef SecureDownloadRef, downloadSize *int64) int32 {
	if _secureDownloadGetDownloadSize == nil {
		panic("Security: symbol SecureDownloadGetDownloadSize not loaded")
	}
	return _secureDownloadGetDownloadSize(downloadRef, downloadSize)
}

var _secureDownloadRelease func(downloadRef SecureDownloadRef) int32

// SecureDownloadRelease releases the memory associated with a secure download object.
//
// See: https://developer.apple.com/documentation/Security/SecureDownloadRelease
func SecureDownloadRelease(downloadRef SecureDownloadRef) int32 {
	if _secureDownloadRelease == nil {
		panic("Security: symbol SecureDownloadRelease not loaded")
	}
	return _secureDownloadRelease(downloadRef)
}

var _secureDownloadUpdateWithData func(downloadRef SecureDownloadRef, data corefoundation.CFDataRef) int32

// SecureDownloadUpdateWithData checks data received during download for validity.
//
// See: https://developer.apple.com/documentation/Security/SecureDownloadUpdateWithData
func SecureDownloadUpdateWithData(downloadRef SecureDownloadRef, data corefoundation.CFDataRef) int32 {
	if _secureDownloadUpdateWithData == nil {
		panic("Security: symbol SecureDownloadUpdateWithData not loaded")
	}
	return _secureDownloadUpdateWithData(downloadRef, data)
}

var _sessionCreate func(flags SessionCreationFlags, attributes SessionAttributeBits) int32

// SessionCreate creates a security session.
//
// See: https://developer.apple.com/documentation/Security/SessionCreate(_:_:)
func SessionCreate(flags SessionCreationFlags, attributes SessionAttributeBits) int32 {
	if _sessionCreate == nil {
		panic("Security: symbol SessionCreate not loaded")
	}
	return _sessionCreate(flags, attributes)
}

var _sessionGetInfo func(session SecuritySessionId, sessionId *SecuritySessionId, attributes *SessionAttributeBits) int32

// SessionGetInfo obtains information about a security session.
//
// See: https://developer.apple.com/documentation/Security/SessionGetInfo(_:_:_:)
func SessionGetInfo(session SecuritySessionId, sessionId *SecuritySessionId, attributes *SessionAttributeBits) int32 {
	if _sessionGetInfo == nil {
		panic("Security: symbol SessionGetInfo not loaded")
	}
	return _sessionGetInfo(session, sessionId, attributes)
}

var _cssmAlgToOid func(algId CSSM_ALGORITHMS) unsafe.Pointer

// CssmAlgToOid.
//
// See: https://developer.apple.com/documentation/Security/cssmAlgToOid(_:)
func CssmAlgToOid(algId CSSM_ALGORITHMS) unsafe.Pointer {
	if _cssmAlgToOid == nil {
		panic("Security: symbol cssmAlgToOid not loaded")
	}
	return _cssmAlgToOid(algId)
}

var _cssmOidToAlg func(oid unsafe.Pointer, alg unsafe.Pointer) bool

// CssmOidToAlg.
//
// See: https://developer.apple.com/documentation/Security/cssmOidToAlg(_:_:)
func CssmOidToAlg(oid unsafe.Pointer, alg unsafe.Pointer) bool {
	if _cssmOidToAlg == nil {
		panic("Security: symbol cssmOidToAlg not loaded")
	}
	return _cssmOidToAlg(oid, alg)
}

var _cssmPerror func(how *byte, err CSSM_RETURN)

// CssmPerror.
//
// See: https://developer.apple.com/documentation/Security/cssmPerror(_:_:)
func CssmPerror(how *byte, err CSSM_RETURN) {
	if _cssmPerror == nil {
		panic("Security: symbol cssmPerror not loaded")
	}
	_cssmPerror(how, err)
}

var _sec_certificate_copy_ref func(certificate Sec_certificate_t) SecCertificateRef

// Sec_certificate_copy_ref.
//
// See: https://developer.apple.com/documentation/Security/sec_certificate_copy_ref(_:)
func Sec_certificate_copy_ref(certificate Sec_certificate_t) SecCertificateRef {
	if _sec_certificate_copy_ref == nil {
		panic("Security: symbol sec_certificate_copy_ref not loaded")
	}
	return _sec_certificate_copy_ref(certificate)
}

var _sec_certificate_create func(certificate SecCertificateRef) Sec_certificate_t

// Sec_certificate_create.
//
// See: https://developer.apple.com/documentation/Security/sec_certificate_create(_:)
func Sec_certificate_create(certificate SecCertificateRef) Sec_certificate_t {
	if _sec_certificate_create == nil {
		panic("Security: symbol sec_certificate_create not loaded")
	}
	return _sec_certificate_create(certificate)
}

var _sec_identity_access_certificates func(identity Sec_identity_t) bool

// Sec_identity_access_certificates.
//
// See: https://developer.apple.com/documentation/Security/sec_identity_access_certificates(_:_:)
func Sec_identity_access_certificates(identity Sec_identity_t) bool {
	if _sec_identity_access_certificates == nil {
		panic("Security: symbol sec_identity_access_certificates not loaded")
	}
	return _sec_identity_access_certificates(identity)
}

var _sec_identity_copy_certificates_ref func(identity Sec_identity_t) corefoundation.CFArrayRef

// Sec_identity_copy_certificates_ref.
//
// See: https://developer.apple.com/documentation/Security/sec_identity_copy_certificates_ref(_:)
func Sec_identity_copy_certificates_ref(identity Sec_identity_t) corefoundation.CFArrayRef {
	if _sec_identity_copy_certificates_ref == nil {
		panic("Security: symbol sec_identity_copy_certificates_ref not loaded")
	}
	return _sec_identity_copy_certificates_ref(identity)
}

var _sec_identity_copy_ref func(identity Sec_identity_t) SecIdentityRef

// Sec_identity_copy_ref.
//
// See: https://developer.apple.com/documentation/Security/sec_identity_copy_ref(_:)
func Sec_identity_copy_ref(identity Sec_identity_t) SecIdentityRef {
	if _sec_identity_copy_ref == nil {
		panic("Security: symbol sec_identity_copy_ref not loaded")
	}
	return _sec_identity_copy_ref(identity)
}

var _sec_identity_create func(identity SecIdentityRef) Sec_identity_t

// Sec_identity_create.
//
// See: https://developer.apple.com/documentation/Security/sec_identity_create(_:)
func Sec_identity_create(identity SecIdentityRef) Sec_identity_t {
	if _sec_identity_create == nil {
		panic("Security: symbol sec_identity_create not loaded")
	}
	return _sec_identity_create(identity)
}

var _sec_identity_create_with_certificates func(identity SecIdentityRef, certificates corefoundation.CFArrayRef) Sec_identity_t

// Sec_identity_create_with_certificates.
//
// See: https://developer.apple.com/documentation/Security/sec_identity_create_with_certificates(_:_:)
func Sec_identity_create_with_certificates(identity SecIdentityRef, certificates corefoundation.CFArrayRef) Sec_identity_t {
	if _sec_identity_create_with_certificates == nil {
		panic("Security: symbol sec_identity_create_with_certificates not loaded")
	}
	return _sec_identity_create_with_certificates(identity, certificates)
}

var _sec_protocol_metadata_access_distinguished_names func(metadata Sec_protocol_metadata_t) bool

// Sec_protocol_metadata_access_distinguished_names.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_access_distinguished_names(_:_:)
func Sec_protocol_metadata_access_distinguished_names(metadata Sec_protocol_metadata_t) bool {
	if _sec_protocol_metadata_access_distinguished_names == nil {
		panic("Security: symbol sec_protocol_metadata_access_distinguished_names not loaded")
	}
	return _sec_protocol_metadata_access_distinguished_names(metadata)
}

var _sec_protocol_metadata_access_ocsp_response func(metadata Sec_protocol_metadata_t) bool

// Sec_protocol_metadata_access_ocsp_response.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_access_ocsp_response(_:_:)
func Sec_protocol_metadata_access_ocsp_response(metadata Sec_protocol_metadata_t) bool {
	if _sec_protocol_metadata_access_ocsp_response == nil {
		panic("Security: symbol sec_protocol_metadata_access_ocsp_response not loaded")
	}
	return _sec_protocol_metadata_access_ocsp_response(metadata)
}

var _sec_protocol_metadata_access_peer_certificate_chain func(metadata Sec_protocol_metadata_t) bool

// Sec_protocol_metadata_access_peer_certificate_chain.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_access_peer_certificate_chain(_:_:)
func Sec_protocol_metadata_access_peer_certificate_chain(metadata Sec_protocol_metadata_t) bool {
	if _sec_protocol_metadata_access_peer_certificate_chain == nil {
		panic("Security: symbol sec_protocol_metadata_access_peer_certificate_chain not loaded")
	}
	return _sec_protocol_metadata_access_peer_certificate_chain(metadata)
}

var _sec_protocol_metadata_access_pre_shared_keys func(metadata Sec_protocol_metadata_t) bool

// Sec_protocol_metadata_access_pre_shared_keys.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_access_pre_shared_keys(_:_:)
func Sec_protocol_metadata_access_pre_shared_keys(metadata Sec_protocol_metadata_t) bool {
	if _sec_protocol_metadata_access_pre_shared_keys == nil {
		panic("Security: symbol sec_protocol_metadata_access_pre_shared_keys not loaded")
	}
	return _sec_protocol_metadata_access_pre_shared_keys(metadata)
}

var _sec_protocol_metadata_access_supported_signature_algorithms func(metadata Sec_protocol_metadata_t) bool

// Sec_protocol_metadata_access_supported_signature_algorithms.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_access_supported_signature_algorithms(_:_:)
func Sec_protocol_metadata_access_supported_signature_algorithms(metadata Sec_protocol_metadata_t) bool {
	if _sec_protocol_metadata_access_supported_signature_algorithms == nil {
		panic("Security: symbol sec_protocol_metadata_access_supported_signature_algorithms not loaded")
	}
	return _sec_protocol_metadata_access_supported_signature_algorithms(metadata)
}

var _sec_protocol_metadata_challenge_parameters_are_equal func(metadataA Sec_protocol_metadata_t, metadataB Sec_protocol_metadata_t) bool

// Sec_protocol_metadata_challenge_parameters_are_equal.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_challenge_parameters_are_equal(_:_:)
func Sec_protocol_metadata_challenge_parameters_are_equal(metadataA Sec_protocol_metadata_t, metadataB Sec_protocol_metadata_t) bool {
	if _sec_protocol_metadata_challenge_parameters_are_equal == nil {
		panic("Security: symbol sec_protocol_metadata_challenge_parameters_are_equal not loaded")
	}
	return _sec_protocol_metadata_challenge_parameters_are_equal(metadataA, metadataB)
}

var _sec_protocol_metadata_copy_negotiated_protocol func(metadata Sec_protocol_metadata_t) *byte

// Sec_protocol_metadata_copy_negotiated_protocol.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_copy_negotiated_protocol(_:)
func Sec_protocol_metadata_copy_negotiated_protocol(metadata Sec_protocol_metadata_t) *byte {
	if _sec_protocol_metadata_copy_negotiated_protocol == nil {
		panic("Security: symbol sec_protocol_metadata_copy_negotiated_protocol not loaded")
	}
	return _sec_protocol_metadata_copy_negotiated_protocol(metadata)
}

var _sec_protocol_metadata_copy_peer_public_key func(metadata Sec_protocol_metadata_t) uintptr

// Sec_protocol_metadata_copy_peer_public_key.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_copy_peer_public_key(_:)
func Sec_protocol_metadata_copy_peer_public_key(metadata Sec_protocol_metadata_t) dispatch.Data {
	if _sec_protocol_metadata_copy_peer_public_key == nil {
		panic("Security: symbol sec_protocol_metadata_copy_peer_public_key not loaded")
	}
	return dispatch.DataFromHandle(_sec_protocol_metadata_copy_peer_public_key(metadata))
}

var _sec_protocol_metadata_copy_server_name func(metadata Sec_protocol_metadata_t) *byte

// Sec_protocol_metadata_copy_server_name.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_copy_server_name(_:)
func Sec_protocol_metadata_copy_server_name(metadata Sec_protocol_metadata_t) *byte {
	if _sec_protocol_metadata_copy_server_name == nil {
		panic("Security: symbol sec_protocol_metadata_copy_server_name not loaded")
	}
	return _sec_protocol_metadata_copy_server_name(metadata)
}

var _sec_protocol_metadata_create_secret func(metadata Sec_protocol_metadata_t, label_len uintptr, label *byte, exporter_length uintptr) uintptr

// Sec_protocol_metadata_create_secret.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_create_secret(_:_:_:_:)
func Sec_protocol_metadata_create_secret(metadata Sec_protocol_metadata_t, label_len uintptr, label *byte, exporter_length uintptr) dispatch.Data {
	if _sec_protocol_metadata_create_secret == nil {
		panic("Security: symbol sec_protocol_metadata_create_secret not loaded")
	}
	return dispatch.DataFromHandle(_sec_protocol_metadata_create_secret(metadata, label_len, label, exporter_length))
}

var _sec_protocol_metadata_create_secret_with_context func(metadata Sec_protocol_metadata_t, label_len uintptr, label *byte, context_len uintptr, context *uint8, exporter_length uintptr) uintptr

// Sec_protocol_metadata_create_secret_with_context.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_create_secret_with_context(_:_:_:_:_:_:)
func Sec_protocol_metadata_create_secret_with_context(metadata Sec_protocol_metadata_t, label_len uintptr, label *byte, context_len uintptr, context *uint8, exporter_length uintptr) dispatch.Data {
	if _sec_protocol_metadata_create_secret_with_context == nil {
		panic("Security: symbol sec_protocol_metadata_create_secret_with_context not loaded")
	}
	return dispatch.DataFromHandle(_sec_protocol_metadata_create_secret_with_context(metadata, label_len, label, context_len, context, exporter_length))
}

var _sec_protocol_metadata_get_early_data_accepted func(metadata Sec_protocol_metadata_t) bool

// Sec_protocol_metadata_get_early_data_accepted.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_get_early_data_accepted(_:)
func Sec_protocol_metadata_get_early_data_accepted(metadata Sec_protocol_metadata_t) bool {
	if _sec_protocol_metadata_get_early_data_accepted == nil {
		panic("Security: symbol sec_protocol_metadata_get_early_data_accepted not loaded")
	}
	return _sec_protocol_metadata_get_early_data_accepted(metadata)
}

var _sec_protocol_metadata_get_negotiated_ciphersuite func(metadata Sec_protocol_metadata_t) SSLCipherSuite

// Sec_protocol_metadata_get_negotiated_ciphersuite.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_get_negotiated_ciphersuite(_:)
func Sec_protocol_metadata_get_negotiated_ciphersuite(metadata Sec_protocol_metadata_t) SSLCipherSuite {
	if _sec_protocol_metadata_get_negotiated_ciphersuite == nil {
		panic("Security: symbol sec_protocol_metadata_get_negotiated_ciphersuite not loaded")
	}
	return _sec_protocol_metadata_get_negotiated_ciphersuite(metadata)
}

var _sec_protocol_metadata_get_negotiated_protocol func(metadata Sec_protocol_metadata_t) *byte

// Sec_protocol_metadata_get_negotiated_protocol.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_get_negotiated_protocol(_:)
func Sec_protocol_metadata_get_negotiated_protocol(metadata Sec_protocol_metadata_t) *byte {
	if _sec_protocol_metadata_get_negotiated_protocol == nil {
		panic("Security: symbol sec_protocol_metadata_get_negotiated_protocol not loaded")
	}
	return _sec_protocol_metadata_get_negotiated_protocol(metadata)
}

var _sec_protocol_metadata_get_negotiated_protocol_version func(metadata Sec_protocol_metadata_t) SSLProtocol

// Sec_protocol_metadata_get_negotiated_protocol_version.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_get_negotiated_protocol_version(_:)
func Sec_protocol_metadata_get_negotiated_protocol_version(metadata Sec_protocol_metadata_t) SSLProtocol {
	if _sec_protocol_metadata_get_negotiated_protocol_version == nil {
		panic("Security: symbol sec_protocol_metadata_get_negotiated_protocol_version not loaded")
	}
	return _sec_protocol_metadata_get_negotiated_protocol_version(metadata)
}

var _sec_protocol_metadata_get_negotiated_tls_ciphersuite func(metadata Sec_protocol_metadata_t) uint16

// Sec_protocol_metadata_get_negotiated_tls_ciphersuite.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_get_negotiated_tls_ciphersuite(_:)
func Sec_protocol_metadata_get_negotiated_tls_ciphersuite(metadata Sec_protocol_metadata_t) uint16 {
	if _sec_protocol_metadata_get_negotiated_tls_ciphersuite == nil {
		panic("Security: symbol sec_protocol_metadata_get_negotiated_tls_ciphersuite not loaded")
	}
	return _sec_protocol_metadata_get_negotiated_tls_ciphersuite(metadata)
}

var _sec_protocol_metadata_get_negotiated_tls_protocol_version func(metadata Sec_protocol_metadata_t) uint16

// Sec_protocol_metadata_get_negotiated_tls_protocol_version.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_get_negotiated_tls_protocol_version(_:)
func Sec_protocol_metadata_get_negotiated_tls_protocol_version(metadata Sec_protocol_metadata_t) uint16 {
	if _sec_protocol_metadata_get_negotiated_tls_protocol_version == nil {
		panic("Security: symbol sec_protocol_metadata_get_negotiated_tls_protocol_version not loaded")
	}
	return _sec_protocol_metadata_get_negotiated_tls_protocol_version(metadata)
}

var _sec_protocol_metadata_get_server_name func(metadata Sec_protocol_metadata_t) *byte

// Sec_protocol_metadata_get_server_name.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_get_server_name(_:)
func Sec_protocol_metadata_get_server_name(metadata Sec_protocol_metadata_t) *byte {
	if _sec_protocol_metadata_get_server_name == nil {
		panic("Security: symbol sec_protocol_metadata_get_server_name not loaded")
	}
	return _sec_protocol_metadata_get_server_name(metadata)
}

var _sec_protocol_metadata_peers_are_equal func(metadataA Sec_protocol_metadata_t, metadataB Sec_protocol_metadata_t) bool

// Sec_protocol_metadata_peers_are_equal.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_peers_are_equal(_:_:)
func Sec_protocol_metadata_peers_are_equal(metadataA Sec_protocol_metadata_t, metadataB Sec_protocol_metadata_t) bool {
	if _sec_protocol_metadata_peers_are_equal == nil {
		panic("Security: symbol sec_protocol_metadata_peers_are_equal not loaded")
	}
	return _sec_protocol_metadata_peers_are_equal(metadataA, metadataB)
}

var _sec_protocol_options_add_pre_shared_key func(options Sec_protocol_options_t, psk uintptr, psk_identity uintptr)

// Sec_protocol_options_add_pre_shared_key.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_add_pre_shared_key(_:_:_:)
func Sec_protocol_options_add_pre_shared_key(options Sec_protocol_options_t, psk dispatch.Data, psk_identity dispatch.Data) {
	if _sec_protocol_options_add_pre_shared_key == nil {
		panic("Security: symbol sec_protocol_options_add_pre_shared_key not loaded")
	}
	_sec_protocol_options_add_pre_shared_key(options, uintptr(psk.Handle()), uintptr(psk_identity.Handle()))
}

var _sec_protocol_options_add_tls_application_protocol func(options Sec_protocol_options_t, application_protocol *byte)

// Sec_protocol_options_add_tls_application_protocol.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_add_tls_application_protocol(_:_:)
func Sec_protocol_options_add_tls_application_protocol(options Sec_protocol_options_t, application_protocol *byte) {
	if _sec_protocol_options_add_tls_application_protocol == nil {
		panic("Security: symbol sec_protocol_options_add_tls_application_protocol not loaded")
	}
	_sec_protocol_options_add_tls_application_protocol(options, application_protocol)
}

var _sec_protocol_options_append_tls_ciphersuite func(options Sec_protocol_options_t, ciphersuite uint16)

// Sec_protocol_options_append_tls_ciphersuite.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_append_tls_ciphersuite(_:_:)
func Sec_protocol_options_append_tls_ciphersuite(options Sec_protocol_options_t, ciphersuite uint16) {
	if _sec_protocol_options_append_tls_ciphersuite == nil {
		panic("Security: symbol sec_protocol_options_append_tls_ciphersuite not loaded")
	}
	_sec_protocol_options_append_tls_ciphersuite(options, ciphersuite)
}

var _sec_protocol_options_append_tls_ciphersuite_group func(options Sec_protocol_options_t, group uint16)

// Sec_protocol_options_append_tls_ciphersuite_group.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_append_tls_ciphersuite_group(_:_:)
func Sec_protocol_options_append_tls_ciphersuite_group(options Sec_protocol_options_t, group uint16) {
	if _sec_protocol_options_append_tls_ciphersuite_group == nil {
		panic("Security: symbol sec_protocol_options_append_tls_ciphersuite_group not loaded")
	}
	_sec_protocol_options_append_tls_ciphersuite_group(options, group)
}

var _sec_protocol_options_are_equal func(optionsA Sec_protocol_options_t, optionsB Sec_protocol_options_t) bool

// Sec_protocol_options_are_equal.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_are_equal(_:_:)
func Sec_protocol_options_are_equal(optionsA Sec_protocol_options_t, optionsB Sec_protocol_options_t) bool {
	if _sec_protocol_options_are_equal == nil {
		panic("Security: symbol sec_protocol_options_are_equal not loaded")
	}
	return _sec_protocol_options_are_equal(optionsA, optionsB)
}

var _sec_protocol_options_get_default_max_dtls_protocol_version func() uint16

// Sec_protocol_options_get_default_max_dtls_protocol_version.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_get_default_max_dtls_protocol_version()
func Sec_protocol_options_get_default_max_dtls_protocol_version() uint16 {
	if _sec_protocol_options_get_default_max_dtls_protocol_version == nil {
		panic("Security: symbol sec_protocol_options_get_default_max_dtls_protocol_version not loaded")
	}
	return _sec_protocol_options_get_default_max_dtls_protocol_version()
}

var _sec_protocol_options_get_default_max_tls_protocol_version func() uint16

// Sec_protocol_options_get_default_max_tls_protocol_version.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_get_default_max_tls_protocol_version()
func Sec_protocol_options_get_default_max_tls_protocol_version() uint16 {
	if _sec_protocol_options_get_default_max_tls_protocol_version == nil {
		panic("Security: symbol sec_protocol_options_get_default_max_tls_protocol_version not loaded")
	}
	return _sec_protocol_options_get_default_max_tls_protocol_version()
}

var _sec_protocol_options_get_default_min_dtls_protocol_version func() uint16

// Sec_protocol_options_get_default_min_dtls_protocol_version.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_get_default_min_dtls_protocol_version()
func Sec_protocol_options_get_default_min_dtls_protocol_version() uint16 {
	if _sec_protocol_options_get_default_min_dtls_protocol_version == nil {
		panic("Security: symbol sec_protocol_options_get_default_min_dtls_protocol_version not loaded")
	}
	return _sec_protocol_options_get_default_min_dtls_protocol_version()
}

var _sec_protocol_options_get_default_min_tls_protocol_version func() uint16

// Sec_protocol_options_get_default_min_tls_protocol_version.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_get_default_min_tls_protocol_version()
func Sec_protocol_options_get_default_min_tls_protocol_version() uint16 {
	if _sec_protocol_options_get_default_min_tls_protocol_version == nil {
		panic("Security: symbol sec_protocol_options_get_default_min_tls_protocol_version not loaded")
	}
	return _sec_protocol_options_get_default_min_tls_protocol_version()
}

var _sec_protocol_options_get_enable_encrypted_client_hello func(options Sec_protocol_options_t) bool

// Sec_protocol_options_get_enable_encrypted_client_hello.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_get_enable_encrypted_client_hello
func Sec_protocol_options_get_enable_encrypted_client_hello(options Sec_protocol_options_t) bool {
	if _sec_protocol_options_get_enable_encrypted_client_hello == nil {
		panic("Security: symbol sec_protocol_options_get_enable_encrypted_client_hello not loaded")
	}
	return _sec_protocol_options_get_enable_encrypted_client_hello(options)
}

var _sec_protocol_options_get_quic_use_legacy_codepoint func(options Sec_protocol_options_t) bool

// Sec_protocol_options_get_quic_use_legacy_codepoint.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_get_quic_use_legacy_codepoint
func Sec_protocol_options_get_quic_use_legacy_codepoint(options Sec_protocol_options_t) bool {
	if _sec_protocol_options_get_quic_use_legacy_codepoint == nil {
		panic("Security: symbol sec_protocol_options_get_quic_use_legacy_codepoint not loaded")
	}
	return _sec_protocol_options_get_quic_use_legacy_codepoint(options)
}

var _sec_protocol_options_set_challenge_block func(options Sec_protocol_options_t, challenge_block Sec_protocol_challenge_t, challenge_queue uintptr)

// Sec_protocol_options_set_challenge_block.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_challenge_block(_:_:_:)
func Sec_protocol_options_set_challenge_block(options Sec_protocol_options_t, challenge_block Sec_protocol_challenge_t, challenge_queue dispatch.Queue) {
	if _sec_protocol_options_set_challenge_block == nil {
		panic("Security: symbol sec_protocol_options_set_challenge_block not loaded")
	}
	_sec_protocol_options_set_challenge_block(options, challenge_block, uintptr(challenge_queue.Handle()))
}

var _sec_protocol_options_set_enable_encrypted_client_hello func(options Sec_protocol_options_t, enable_encrypted_client_hello bool)

// Sec_protocol_options_set_enable_encrypted_client_hello.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_enable_encrypted_client_hello
func Sec_protocol_options_set_enable_encrypted_client_hello(options Sec_protocol_options_t, enable_encrypted_client_hello bool) {
	if _sec_protocol_options_set_enable_encrypted_client_hello == nil {
		panic("Security: symbol sec_protocol_options_set_enable_encrypted_client_hello not loaded")
	}
	_sec_protocol_options_set_enable_encrypted_client_hello(options, enable_encrypted_client_hello)
}

var _sec_protocol_options_set_key_update_block func(options Sec_protocol_options_t, key_update_block Sec_protocol_key_update_t, key_update_queue uintptr)

// Sec_protocol_options_set_key_update_block.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_key_update_block(_:_:_:)
func Sec_protocol_options_set_key_update_block(options Sec_protocol_options_t, key_update_block Sec_protocol_key_update_t, key_update_queue dispatch.Queue) {
	if _sec_protocol_options_set_key_update_block == nil {
		panic("Security: symbol sec_protocol_options_set_key_update_block not loaded")
	}
	_sec_protocol_options_set_key_update_block(options, key_update_block, uintptr(key_update_queue.Handle()))
}

var _sec_protocol_options_set_local_identity func(options Sec_protocol_options_t, identity Sec_identity_t)

// Sec_protocol_options_set_local_identity.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_local_identity(_:_:)
func Sec_protocol_options_set_local_identity(options Sec_protocol_options_t, identity Sec_identity_t) {
	if _sec_protocol_options_set_local_identity == nil {
		panic("Security: symbol sec_protocol_options_set_local_identity not loaded")
	}
	_sec_protocol_options_set_local_identity(options, identity)
}

var _sec_protocol_options_set_max_tls_protocol_version func(options Sec_protocol_options_t, version uint16)

// Sec_protocol_options_set_max_tls_protocol_version.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_max_tls_protocol_version(_:_:)
func Sec_protocol_options_set_max_tls_protocol_version(options Sec_protocol_options_t, version uint16) {
	if _sec_protocol_options_set_max_tls_protocol_version == nil {
		panic("Security: symbol sec_protocol_options_set_max_tls_protocol_version not loaded")
	}
	_sec_protocol_options_set_max_tls_protocol_version(options, version)
}

var _sec_protocol_options_set_min_tls_protocol_version func(options Sec_protocol_options_t, version uint16)

// Sec_protocol_options_set_min_tls_protocol_version.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_min_tls_protocol_version(_:_:)
func Sec_protocol_options_set_min_tls_protocol_version(options Sec_protocol_options_t, version uint16) {
	if _sec_protocol_options_set_min_tls_protocol_version == nil {
		panic("Security: symbol sec_protocol_options_set_min_tls_protocol_version not loaded")
	}
	_sec_protocol_options_set_min_tls_protocol_version(options, version)
}

var _sec_protocol_options_set_peer_authentication_optional func(options Sec_protocol_options_t, peer_authentication_optional bool)

// Sec_protocol_options_set_peer_authentication_optional.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_peer_authentication_optional
func Sec_protocol_options_set_peer_authentication_optional(options Sec_protocol_options_t, peer_authentication_optional bool) {
	if _sec_protocol_options_set_peer_authentication_optional == nil {
		panic("Security: symbol sec_protocol_options_set_peer_authentication_optional not loaded")
	}
	_sec_protocol_options_set_peer_authentication_optional(options, peer_authentication_optional)
}

var _sec_protocol_options_set_peer_authentication_required func(options Sec_protocol_options_t, peer_authentication_required bool)

// Sec_protocol_options_set_peer_authentication_required.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_peer_authentication_required(_:_:)
func Sec_protocol_options_set_peer_authentication_required(options Sec_protocol_options_t, peer_authentication_required bool) {
	if _sec_protocol_options_set_peer_authentication_required == nil {
		panic("Security: symbol sec_protocol_options_set_peer_authentication_required not loaded")
	}
	_sec_protocol_options_set_peer_authentication_required(options, peer_authentication_required)
}

var _sec_protocol_options_set_pre_shared_key_selection_block func(options Sec_protocol_options_t, psk_selection_block Sec_protocol_pre_shared_key_selection_t, psk_selection_queue uintptr)

// Sec_protocol_options_set_pre_shared_key_selection_block.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_pre_shared_key_selection_block(_:_:_:)
func Sec_protocol_options_set_pre_shared_key_selection_block(options Sec_protocol_options_t, psk_selection_block Sec_protocol_pre_shared_key_selection_t, psk_selection_queue dispatch.Queue) {
	if _sec_protocol_options_set_pre_shared_key_selection_block == nil {
		panic("Security: symbol sec_protocol_options_set_pre_shared_key_selection_block not loaded")
	}
	_sec_protocol_options_set_pre_shared_key_selection_block(options, psk_selection_block, uintptr(psk_selection_queue.Handle()))
}

var _sec_protocol_options_set_quic_use_legacy_codepoint func(options Sec_protocol_options_t, quic_use_legacy_codepoint bool)

// Sec_protocol_options_set_quic_use_legacy_codepoint.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_quic_use_legacy_codepoint
func Sec_protocol_options_set_quic_use_legacy_codepoint(options Sec_protocol_options_t, quic_use_legacy_codepoint bool) {
	if _sec_protocol_options_set_quic_use_legacy_codepoint == nil {
		panic("Security: symbol sec_protocol_options_set_quic_use_legacy_codepoint not loaded")
	}
	_sec_protocol_options_set_quic_use_legacy_codepoint(options, quic_use_legacy_codepoint)
}

var _sec_protocol_options_set_tls_false_start_enabled func(options Sec_protocol_options_t, false_start_enabled bool)

// Sec_protocol_options_set_tls_false_start_enabled.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_false_start_enabled(_:_:)
func Sec_protocol_options_set_tls_false_start_enabled(options Sec_protocol_options_t, false_start_enabled bool) {
	if _sec_protocol_options_set_tls_false_start_enabled == nil {
		panic("Security: symbol sec_protocol_options_set_tls_false_start_enabled not loaded")
	}
	_sec_protocol_options_set_tls_false_start_enabled(options, false_start_enabled)
}

var _sec_protocol_options_set_tls_is_fallback_attempt func(options Sec_protocol_options_t, is_fallback_attempt bool)

// Sec_protocol_options_set_tls_is_fallback_attempt.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_is_fallback_attempt(_:_:)
func Sec_protocol_options_set_tls_is_fallback_attempt(options Sec_protocol_options_t, is_fallback_attempt bool) {
	if _sec_protocol_options_set_tls_is_fallback_attempt == nil {
		panic("Security: symbol sec_protocol_options_set_tls_is_fallback_attempt not loaded")
	}
	_sec_protocol_options_set_tls_is_fallback_attempt(options, is_fallback_attempt)
}

var _sec_protocol_options_set_tls_max_version func(options Sec_protocol_options_t, version SSLProtocol)

// Sec_protocol_options_set_tls_max_version.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_max_version(_:_:)
func Sec_protocol_options_set_tls_max_version(options Sec_protocol_options_t, version SSLProtocol) {
	if _sec_protocol_options_set_tls_max_version == nil {
		panic("Security: symbol sec_protocol_options_set_tls_max_version not loaded")
	}
	_sec_protocol_options_set_tls_max_version(options, version)
}

var _sec_protocol_options_set_tls_min_version func(options Sec_protocol_options_t, version SSLProtocol)

// Sec_protocol_options_set_tls_min_version.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_min_version(_:_:)
func Sec_protocol_options_set_tls_min_version(options Sec_protocol_options_t, version SSLProtocol) {
	if _sec_protocol_options_set_tls_min_version == nil {
		panic("Security: symbol sec_protocol_options_set_tls_min_version not loaded")
	}
	_sec_protocol_options_set_tls_min_version(options, version)
}

var _sec_protocol_options_set_tls_ocsp_enabled func(options Sec_protocol_options_t, ocsp_enabled bool)

// Sec_protocol_options_set_tls_ocsp_enabled.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_ocsp_enabled(_:_:)
func Sec_protocol_options_set_tls_ocsp_enabled(options Sec_protocol_options_t, ocsp_enabled bool) {
	if _sec_protocol_options_set_tls_ocsp_enabled == nil {
		panic("Security: symbol sec_protocol_options_set_tls_ocsp_enabled not loaded")
	}
	_sec_protocol_options_set_tls_ocsp_enabled(options, ocsp_enabled)
}

var _sec_protocol_options_set_tls_pre_shared_key_identity_hint func(options Sec_protocol_options_t, psk_identity_hint uintptr)

// Sec_protocol_options_set_tls_pre_shared_key_identity_hint.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_pre_shared_key_identity_hint(_:_:)
func Sec_protocol_options_set_tls_pre_shared_key_identity_hint(options Sec_protocol_options_t, psk_identity_hint dispatch.Data) {
	if _sec_protocol_options_set_tls_pre_shared_key_identity_hint == nil {
		panic("Security: symbol sec_protocol_options_set_tls_pre_shared_key_identity_hint not loaded")
	}
	_sec_protocol_options_set_tls_pre_shared_key_identity_hint(options, uintptr(psk_identity_hint.Handle()))
}

var _sec_protocol_options_set_tls_renegotiation_enabled func(options Sec_protocol_options_t, renegotiation_enabled bool)

// Sec_protocol_options_set_tls_renegotiation_enabled.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_renegotiation_enabled(_:_:)
func Sec_protocol_options_set_tls_renegotiation_enabled(options Sec_protocol_options_t, renegotiation_enabled bool) {
	if _sec_protocol_options_set_tls_renegotiation_enabled == nil {
		panic("Security: symbol sec_protocol_options_set_tls_renegotiation_enabled not loaded")
	}
	_sec_protocol_options_set_tls_renegotiation_enabled(options, renegotiation_enabled)
}

var _sec_protocol_options_set_tls_resumption_enabled func(options Sec_protocol_options_t, resumption_enabled bool)

// Sec_protocol_options_set_tls_resumption_enabled.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_resumption_enabled(_:_:)
func Sec_protocol_options_set_tls_resumption_enabled(options Sec_protocol_options_t, resumption_enabled bool) {
	if _sec_protocol_options_set_tls_resumption_enabled == nil {
		panic("Security: symbol sec_protocol_options_set_tls_resumption_enabled not loaded")
	}
	_sec_protocol_options_set_tls_resumption_enabled(options, resumption_enabled)
}

var _sec_protocol_options_set_tls_sct_enabled func(options Sec_protocol_options_t, sct_enabled bool)

// Sec_protocol_options_set_tls_sct_enabled.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_sct_enabled(_:_:)
func Sec_protocol_options_set_tls_sct_enabled(options Sec_protocol_options_t, sct_enabled bool) {
	if _sec_protocol_options_set_tls_sct_enabled == nil {
		panic("Security: symbol sec_protocol_options_set_tls_sct_enabled not loaded")
	}
	_sec_protocol_options_set_tls_sct_enabled(options, sct_enabled)
}

var _sec_protocol_options_set_tls_server_name func(options Sec_protocol_options_t, server_name *byte)

// Sec_protocol_options_set_tls_server_name.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_server_name(_:_:)
func Sec_protocol_options_set_tls_server_name(options Sec_protocol_options_t, server_name *byte) {
	if _sec_protocol_options_set_tls_server_name == nil {
		panic("Security: symbol sec_protocol_options_set_tls_server_name not loaded")
	}
	_sec_protocol_options_set_tls_server_name(options, server_name)
}

var _sec_protocol_options_set_tls_tickets_enabled func(options Sec_protocol_options_t, tickets_enabled bool)

// Sec_protocol_options_set_tls_tickets_enabled.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_tickets_enabled(_:_:)
func Sec_protocol_options_set_tls_tickets_enabled(options Sec_protocol_options_t, tickets_enabled bool) {
	if _sec_protocol_options_set_tls_tickets_enabled == nil {
		panic("Security: symbol sec_protocol_options_set_tls_tickets_enabled not loaded")
	}
	_sec_protocol_options_set_tls_tickets_enabled(options, tickets_enabled)
}

var _sec_protocol_options_set_verify_block func(options Sec_protocol_options_t, verify_block Sec_protocol_verify_t, verify_block_queue uintptr)

// Sec_protocol_options_set_verify_block.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_set_verify_block(_:_:_:)
func Sec_protocol_options_set_verify_block(options Sec_protocol_options_t, verify_block Sec_protocol_verify_t, verify_block_queue dispatch.Queue) {
	if _sec_protocol_options_set_verify_block == nil {
		panic("Security: symbol sec_protocol_options_set_verify_block not loaded")
	}
	_sec_protocol_options_set_verify_block(options, verify_block, uintptr(verify_block_queue.Handle()))
}

var _sec_release func(obj unsafe.Pointer)

// Sec_release.
//
// See: https://developer.apple.com/documentation/Security/sec_release(_:)
func Sec_release(obj unsafe.Pointer) {
	if _sec_release == nil {
		panic("Security: symbol sec_release not loaded")
	}
	_sec_release(obj)
}

var _sec_retain func(obj unsafe.Pointer) unsafe.Pointer

// Sec_retain.
//
// See: https://developer.apple.com/documentation/Security/sec_retain(_:)
func Sec_retain(obj unsafe.Pointer) unsafe.Pointer {
	if _sec_retain == nil {
		panic("Security: symbol sec_retain not loaded")
	}
	return _sec_retain(obj)
}

var _sec_trust_copy_ref func(trust Sec_trust_t) SecTrustRef

// Sec_trust_copy_ref.
//
// See: https://developer.apple.com/documentation/Security/sec_trust_copy_ref(_:)
func Sec_trust_copy_ref(trust Sec_trust_t) SecTrustRef {
	if _sec_trust_copy_ref == nil {
		panic("Security: symbol sec_trust_copy_ref not loaded")
	}
	return _sec_trust_copy_ref(trust)
}

var _sec_trust_create func(trust SecTrustRef) Sec_trust_t

// Sec_trust_create.
//
// See: https://developer.apple.com/documentation/Security/sec_trust_create(_:)
func Sec_trust_create(trust SecTrustRef) Sec_trust_t {
	if _sec_trust_create == nil {
		panic("Security: symbol sec_trust_create not loaded")
	}
	return _sec_trust_create(trust)
}

func init() {
	if frameworkHandle == 0 {
		return
	}
		registerFunc(&_authorizationCopyInfo, frameworkHandle, "AuthorizationCopyInfo")
		registerFunc(&_authorizationCopyRights, frameworkHandle, "AuthorizationCopyRights")
		registerFunc(&_authorizationCopyRightsAsync, frameworkHandle, "AuthorizationCopyRightsAsync")
		registerFunc(&_authorizationCreate, frameworkHandle, "AuthorizationCreate")
		registerFunc(&_authorizationCreateFromExternalForm, frameworkHandle, "AuthorizationCreateFromExternalForm")
		registerFunc(&_authorizationFree, frameworkHandle, "AuthorizationFree")
		registerFunc(&_authorizationFreeItemSet, frameworkHandle, "AuthorizationFreeItemSet")
		registerFunc(&_authorizationMakeExternalForm, frameworkHandle, "AuthorizationMakeExternalForm")
		registerFunc(&_authorizationRightGet, frameworkHandle, "AuthorizationRightGet")
		registerFunc(&_authorizationRightRemove, frameworkHandle, "AuthorizationRightRemove")
		registerFunc(&_authorizationRightSet, frameworkHandle, "AuthorizationRightSet")
		registerFunc(&_cMSDecoderCopyAllCerts, frameworkHandle, "CMSDecoderCopyAllCerts")
		registerFunc(&_cMSDecoderCopyContent, frameworkHandle, "CMSDecoderCopyContent")
		registerFunc(&_cMSDecoderCopyDetachedContent, frameworkHandle, "CMSDecoderCopyDetachedContent")
		registerFunc(&_cMSDecoderCopyEncapsulatedContentType, frameworkHandle, "CMSDecoderCopyEncapsulatedContentType")
		registerFunc(&_cMSDecoderCopySignerCert, frameworkHandle, "CMSDecoderCopySignerCert")
		registerFunc(&_cMSDecoderCopySignerEmailAddress, frameworkHandle, "CMSDecoderCopySignerEmailAddress")
		registerFunc(&_cMSDecoderCopySignerSigningTime, frameworkHandle, "CMSDecoderCopySignerSigningTime")
		registerFunc(&_cMSDecoderCopySignerStatus, frameworkHandle, "CMSDecoderCopySignerStatus")
		registerFunc(&_cMSDecoderCopySignerTimestamp, frameworkHandle, "CMSDecoderCopySignerTimestamp")
		registerFunc(&_cMSDecoderCopySignerTimestampCertificates, frameworkHandle, "CMSDecoderCopySignerTimestampCertificates")
		registerFunc(&_cMSDecoderCopySignerTimestampWithPolicy, frameworkHandle, "CMSDecoderCopySignerTimestampWithPolicy")
		registerFunc(&_cMSDecoderCreate, frameworkHandle, "CMSDecoderCreate")
		registerFunc(&_cMSDecoderFinalizeMessage, frameworkHandle, "CMSDecoderFinalizeMessage")
		registerFunc(&_cMSDecoderGetNumSigners, frameworkHandle, "CMSDecoderGetNumSigners")
		registerFunc(&_cMSDecoderGetTypeID, frameworkHandle, "CMSDecoderGetTypeID")
		registerFunc(&_cMSDecoderIsContentEncrypted, frameworkHandle, "CMSDecoderIsContentEncrypted")
		registerFunc(&_cMSDecoderSetDetachedContent, frameworkHandle, "CMSDecoderSetDetachedContent")
		registerFunc(&_cMSDecoderSetSearchKeychain, frameworkHandle, "CMSDecoderSetSearchKeychain")
		registerFunc(&_cMSDecoderUpdateMessage, frameworkHandle, "CMSDecoderUpdateMessage")
		registerFunc(&_cMSEncode, frameworkHandle, "CMSEncode")
		registerFunc(&_cMSEncodeContent, frameworkHandle, "CMSEncodeContent")
		registerFunc(&_cMSEncoderAddRecipients, frameworkHandle, "CMSEncoderAddRecipients")
		registerFunc(&_cMSEncoderAddSignedAttributes, frameworkHandle, "CMSEncoderAddSignedAttributes")
		registerFunc(&_cMSEncoderAddSigners, frameworkHandle, "CMSEncoderAddSigners")
		registerFunc(&_cMSEncoderAddSupportingCerts, frameworkHandle, "CMSEncoderAddSupportingCerts")
		registerFunc(&_cMSEncoderCopyEncapsulatedContentType, frameworkHandle, "CMSEncoderCopyEncapsulatedContentType")
		registerFunc(&_cMSEncoderCopyEncodedContent, frameworkHandle, "CMSEncoderCopyEncodedContent")
		registerFunc(&_cMSEncoderCopyRecipients, frameworkHandle, "CMSEncoderCopyRecipients")
		registerFunc(&_cMSEncoderCopySignerTimestamp, frameworkHandle, "CMSEncoderCopySignerTimestamp")
		registerFunc(&_cMSEncoderCopySignerTimestampWithPolicy, frameworkHandle, "CMSEncoderCopySignerTimestampWithPolicy")
		registerFunc(&_cMSEncoderCopySigners, frameworkHandle, "CMSEncoderCopySigners")
		registerFunc(&_cMSEncoderCopySupportingCerts, frameworkHandle, "CMSEncoderCopySupportingCerts")
		registerFunc(&_cMSEncoderCreate, frameworkHandle, "CMSEncoderCreate")
		registerFunc(&_cMSEncoderGetCertificateChainMode, frameworkHandle, "CMSEncoderGetCertificateChainMode")
		registerFunc(&_cMSEncoderGetHasDetachedContent, frameworkHandle, "CMSEncoderGetHasDetachedContent")
		registerFunc(&_cMSEncoderGetTypeID, frameworkHandle, "CMSEncoderGetTypeID")
		registerFunc(&_cMSEncoderSetCertificateChainMode, frameworkHandle, "CMSEncoderSetCertificateChainMode")
		registerFunc(&_cMSEncoderSetEncapsulatedContentType, frameworkHandle, "CMSEncoderSetEncapsulatedContentType")
		registerFunc(&_cMSEncoderSetEncapsulatedContentTypeOID, frameworkHandle, "CMSEncoderSetEncapsulatedContentTypeOID")
		registerFunc(&_cMSEncoderSetHasDetachedContent, frameworkHandle, "CMSEncoderSetHasDetachedContent")
		registerFunc(&_cMSEncoderSetSignerAlgorithm, frameworkHandle, "CMSEncoderSetSignerAlgorithm")
		registerFunc(&_cMSEncoderUpdateContent, frameworkHandle, "CMSEncoderUpdateContent")
		registerFunc(&_cSSM_AC_AuthCompute, frameworkHandle, "CSSM_AC_AuthCompute")
		registerFunc(&_cSSM_AC_PassThrough, frameworkHandle, "CSSM_AC_PassThrough")
		registerFunc(&_cSSM_CL_CertAbortCache, frameworkHandle, "CSSM_CL_CertAbortCache")
		registerFunc(&_cSSM_CL_CertAbortQuery, frameworkHandle, "CSSM_CL_CertAbortQuery")
		registerFunc(&_cSSM_CL_CertCache, frameworkHandle, "CSSM_CL_CertCache")
		registerFunc(&_cSSM_CL_CertCreateTemplate, frameworkHandle, "CSSM_CL_CertCreateTemplate")
		registerFunc(&_cSSM_CL_CertDescribeFormat, frameworkHandle, "CSSM_CL_CertDescribeFormat")
		registerFunc(&_cSSM_CL_CertGetAllFields, frameworkHandle, "CSSM_CL_CertGetAllFields")
		registerFunc(&_cSSM_CL_CertGetAllTemplateFields, frameworkHandle, "CSSM_CL_CertGetAllTemplateFields")
		registerFunc(&_cSSM_CL_CertGetFirstCachedFieldValue, frameworkHandle, "CSSM_CL_CertGetFirstCachedFieldValue")
		registerFunc(&_cSSM_CL_CertGetFirstFieldValue, frameworkHandle, "CSSM_CL_CertGetFirstFieldValue")
		registerFunc(&_cSSM_CL_CertGetKeyInfo, frameworkHandle, "CSSM_CL_CertGetKeyInfo")
		registerFunc(&_cSSM_CL_CertGetNextCachedFieldValue, frameworkHandle, "CSSM_CL_CertGetNextCachedFieldValue")
		registerFunc(&_cSSM_CL_CertGetNextFieldValue, frameworkHandle, "CSSM_CL_CertGetNextFieldValue")
		registerFunc(&_cSSM_CL_CertGroupFromVerifiedBundle, frameworkHandle, "CSSM_CL_CertGroupFromVerifiedBundle")
		registerFunc(&_cSSM_CL_CertGroupToSignedBundle, frameworkHandle, "CSSM_CL_CertGroupToSignedBundle")
		registerFunc(&_cSSM_CL_CertSign, frameworkHandle, "CSSM_CL_CertSign")
		registerFunc(&_cSSM_CL_CertVerify, frameworkHandle, "CSSM_CL_CertVerify")
		registerFunc(&_cSSM_CL_CertVerifyWithKey, frameworkHandle, "CSSM_CL_CertVerifyWithKey")
		registerFunc(&_cSSM_CL_CrlAbortCache, frameworkHandle, "CSSM_CL_CrlAbortCache")
		registerFunc(&_cSSM_CL_CrlAbortQuery, frameworkHandle, "CSSM_CL_CrlAbortQuery")
		registerFunc(&_cSSM_CL_CrlAddCert, frameworkHandle, "CSSM_CL_CrlAddCert")
		registerFunc(&_cSSM_CL_CrlCache, frameworkHandle, "CSSM_CL_CrlCache")
		registerFunc(&_cSSM_CL_CrlCreateTemplate, frameworkHandle, "CSSM_CL_CrlCreateTemplate")
		registerFunc(&_cSSM_CL_CrlDescribeFormat, frameworkHandle, "CSSM_CL_CrlDescribeFormat")
		registerFunc(&_cSSM_CL_CrlGetAllCachedRecordFields, frameworkHandle, "CSSM_CL_CrlGetAllCachedRecordFields")
		registerFunc(&_cSSM_CL_CrlGetAllFields, frameworkHandle, "CSSM_CL_CrlGetAllFields")
		registerFunc(&_cSSM_CL_CrlGetFirstCachedFieldValue, frameworkHandle, "CSSM_CL_CrlGetFirstCachedFieldValue")
		registerFunc(&_cSSM_CL_CrlGetFirstFieldValue, frameworkHandle, "CSSM_CL_CrlGetFirstFieldValue")
		registerFunc(&_cSSM_CL_CrlGetNextCachedFieldValue, frameworkHandle, "CSSM_CL_CrlGetNextCachedFieldValue")
		registerFunc(&_cSSM_CL_CrlGetNextFieldValue, frameworkHandle, "CSSM_CL_CrlGetNextFieldValue")
		registerFunc(&_cSSM_CL_CrlRemoveCert, frameworkHandle, "CSSM_CL_CrlRemoveCert")
		registerFunc(&_cSSM_CL_CrlSetFields, frameworkHandle, "CSSM_CL_CrlSetFields")
		registerFunc(&_cSSM_CL_CrlSign, frameworkHandle, "CSSM_CL_CrlSign")
		registerFunc(&_cSSM_CL_CrlVerify, frameworkHandle, "CSSM_CL_CrlVerify")
		registerFunc(&_cSSM_CL_CrlVerifyWithKey, frameworkHandle, "CSSM_CL_CrlVerifyWithKey")
		registerFunc(&_cSSM_CL_FreeFieldValue, frameworkHandle, "CSSM_CL_FreeFieldValue")
		registerFunc(&_cSSM_CL_FreeFields, frameworkHandle, "CSSM_CL_FreeFields")
		registerFunc(&_cSSM_CL_IsCertInCachedCrl, frameworkHandle, "CSSM_CL_IsCertInCachedCrl")
		registerFunc(&_cSSM_CL_IsCertInCrl, frameworkHandle, "CSSM_CL_IsCertInCrl")
		registerFunc(&_cSSM_CL_PassThrough, frameworkHandle, "CSSM_CL_PassThrough")
		registerFunc(&_cSSM_CSP_ChangeLoginAcl, frameworkHandle, "CSSM_CSP_ChangeLoginAcl")
		registerFunc(&_cSSM_CSP_ChangeLoginOwner, frameworkHandle, "CSSM_CSP_ChangeLoginOwner")
		registerFunc(&_cSSM_CSP_CreateAsymmetricContext, frameworkHandle, "CSSM_CSP_CreateAsymmetricContext")
		registerFunc(&_cSSM_CSP_CreateDeriveKeyContext, frameworkHandle, "CSSM_CSP_CreateDeriveKeyContext")
		registerFunc(&_cSSM_CSP_CreateDigestContext, frameworkHandle, "CSSM_CSP_CreateDigestContext")
		registerFunc(&_cSSM_CSP_CreateKeyGenContext, frameworkHandle, "CSSM_CSP_CreateKeyGenContext")
		registerFunc(&_cSSM_CSP_CreateMacContext, frameworkHandle, "CSSM_CSP_CreateMacContext")
		registerFunc(&_cSSM_CSP_CreatePassThroughContext, frameworkHandle, "CSSM_CSP_CreatePassThroughContext")
		registerFunc(&_cSSM_CSP_CreateRandomGenContext, frameworkHandle, "CSSM_CSP_CreateRandomGenContext")
		registerFunc(&_cSSM_CSP_CreateSignatureContext, frameworkHandle, "CSSM_CSP_CreateSignatureContext")
		registerFunc(&_cSSM_CSP_CreateSymmetricContext, frameworkHandle, "CSSM_CSP_CreateSymmetricContext")
		registerFunc(&_cSSM_CSP_GetLoginAcl, frameworkHandle, "CSSM_CSP_GetLoginAcl")
		registerFunc(&_cSSM_CSP_GetLoginOwner, frameworkHandle, "CSSM_CSP_GetLoginOwner")
		registerFunc(&_cSSM_CSP_GetOperationalStatistics, frameworkHandle, "CSSM_CSP_GetOperationalStatistics")
		registerFunc(&_cSSM_CSP_Login, frameworkHandle, "CSSM_CSP_Login")
		registerFunc(&_cSSM_CSP_Logout, frameworkHandle, "CSSM_CSP_Logout")
		registerFunc(&_cSSM_CSP_ObtainPrivateKeyFromPublicKey, frameworkHandle, "CSSM_CSP_ObtainPrivateKeyFromPublicKey")
		registerFunc(&_cSSM_CSP_PassThrough, frameworkHandle, "CSSM_CSP_PassThrough")
		registerFunc(&_cSSM_ChangeKeyAcl, frameworkHandle, "CSSM_ChangeKeyAcl")
		registerFunc(&_cSSM_ChangeKeyOwner, frameworkHandle, "CSSM_ChangeKeyOwner")
		registerFunc(&_cSSM_DL_Authenticate, frameworkHandle, "CSSM_DL_Authenticate")
		registerFunc(&_cSSM_DL_ChangeDbAcl, frameworkHandle, "CSSM_DL_ChangeDbAcl")
		registerFunc(&_cSSM_DL_ChangeDbOwner, frameworkHandle, "CSSM_DL_ChangeDbOwner")
		registerFunc(&_cSSM_DL_CreateRelation, frameworkHandle, "CSSM_DL_CreateRelation")
		registerFunc(&_cSSM_DL_DataAbortQuery, frameworkHandle, "CSSM_DL_DataAbortQuery")
		registerFunc(&_cSSM_DL_DataDelete, frameworkHandle, "CSSM_DL_DataDelete")
		registerFunc(&_cSSM_DL_DataGetFirst, frameworkHandle, "CSSM_DL_DataGetFirst")
		registerFunc(&_cSSM_DL_DataGetFromUniqueRecordId, frameworkHandle, "CSSM_DL_DataGetFromUniqueRecordId")
		registerFunc(&_cSSM_DL_DataGetNext, frameworkHandle, "CSSM_DL_DataGetNext")
		registerFunc(&_cSSM_DL_DataInsert, frameworkHandle, "CSSM_DL_DataInsert")
		registerFunc(&_cSSM_DL_DataModify, frameworkHandle, "CSSM_DL_DataModify")
		registerFunc(&_cSSM_DL_DbClose, frameworkHandle, "CSSM_DL_DbClose")
		registerFunc(&_cSSM_DL_DbCreate, frameworkHandle, "CSSM_DL_DbCreate")
		registerFunc(&_cSSM_DL_DbDelete, frameworkHandle, "CSSM_DL_DbDelete")
		registerFunc(&_cSSM_DL_DbOpen, frameworkHandle, "CSSM_DL_DbOpen")
		registerFunc(&_cSSM_DL_DestroyRelation, frameworkHandle, "CSSM_DL_DestroyRelation")
		registerFunc(&_cSSM_DL_FreeNameList, frameworkHandle, "CSSM_DL_FreeNameList")
		registerFunc(&_cSSM_DL_FreeUniqueRecord, frameworkHandle, "CSSM_DL_FreeUniqueRecord")
		registerFunc(&_cSSM_DL_GetDbAcl, frameworkHandle, "CSSM_DL_GetDbAcl")
		registerFunc(&_cSSM_DL_GetDbNameFromHandle, frameworkHandle, "CSSM_DL_GetDbNameFromHandle")
		registerFunc(&_cSSM_DL_GetDbNames, frameworkHandle, "CSSM_DL_GetDbNames")
		registerFunc(&_cSSM_DL_GetDbOwner, frameworkHandle, "CSSM_DL_GetDbOwner")
		registerFunc(&_cSSM_DL_PassThrough, frameworkHandle, "CSSM_DL_PassThrough")
		registerFunc(&_cSSM_DecryptData, frameworkHandle, "CSSM_DecryptData")
		registerFunc(&_cSSM_DecryptDataFinal, frameworkHandle, "CSSM_DecryptDataFinal")
		registerFunc(&_cSSM_DecryptDataInit, frameworkHandle, "CSSM_DecryptDataInit")
		registerFunc(&_cSSM_DecryptDataInitP, frameworkHandle, "CSSM_DecryptDataInitP")
		registerFunc(&_cSSM_DecryptDataP, frameworkHandle, "CSSM_DecryptDataP")
		registerFunc(&_cSSM_DecryptDataUpdate, frameworkHandle, "CSSM_DecryptDataUpdate")
		registerFunc(&_cSSM_DeleteContext, frameworkHandle, "CSSM_DeleteContext")
		registerFunc(&_cSSM_DeleteContextAttributes, frameworkHandle, "CSSM_DeleteContextAttributes")
		registerFunc(&_cSSM_DeriveKey, frameworkHandle, "CSSM_DeriveKey")
		registerFunc(&_cSSM_DigestData, frameworkHandle, "CSSM_DigestData")
		registerFunc(&_cSSM_DigestDataClone, frameworkHandle, "CSSM_DigestDataClone")
		registerFunc(&_cSSM_DigestDataFinal, frameworkHandle, "CSSM_DigestDataFinal")
		registerFunc(&_cSSM_DigestDataInit, frameworkHandle, "CSSM_DigestDataInit")
		registerFunc(&_cSSM_DigestDataUpdate, frameworkHandle, "CSSM_DigestDataUpdate")
		registerFunc(&_cSSM_EncryptData, frameworkHandle, "CSSM_EncryptData")
		registerFunc(&_cSSM_EncryptDataFinal, frameworkHandle, "CSSM_EncryptDataFinal")
		registerFunc(&_cSSM_EncryptDataInit, frameworkHandle, "CSSM_EncryptDataInit")
		registerFunc(&_cSSM_EncryptDataInitP, frameworkHandle, "CSSM_EncryptDataInitP")
		registerFunc(&_cSSM_EncryptDataP, frameworkHandle, "CSSM_EncryptDataP")
		registerFunc(&_cSSM_EncryptDataUpdate, frameworkHandle, "CSSM_EncryptDataUpdate")
		registerFunc(&_cSSM_FreeContext, frameworkHandle, "CSSM_FreeContext")
		registerFunc(&_cSSM_FreeKey, frameworkHandle, "CSSM_FreeKey")
		registerFunc(&_cSSM_GenerateAlgorithmParams, frameworkHandle, "CSSM_GenerateAlgorithmParams")
		registerFunc(&_cSSM_GenerateKey, frameworkHandle, "CSSM_GenerateKey")
		registerFunc(&_cSSM_GenerateKeyP, frameworkHandle, "CSSM_GenerateKeyP")
		registerFunc(&_cSSM_GenerateKeyPair, frameworkHandle, "CSSM_GenerateKeyPair")
		registerFunc(&_cSSM_GenerateKeyPairP, frameworkHandle, "CSSM_GenerateKeyPairP")
		registerFunc(&_cSSM_GenerateMac, frameworkHandle, "CSSM_GenerateMac")
		registerFunc(&_cSSM_GenerateMacFinal, frameworkHandle, "CSSM_GenerateMacFinal")
		registerFunc(&_cSSM_GenerateMacInit, frameworkHandle, "CSSM_GenerateMacInit")
		registerFunc(&_cSSM_GenerateMacUpdate, frameworkHandle, "CSSM_GenerateMacUpdate")
		registerFunc(&_cSSM_GenerateRandom, frameworkHandle, "CSSM_GenerateRandom")
		registerFunc(&_cSSM_GetAPIMemoryFunctions, frameworkHandle, "CSSM_GetAPIMemoryFunctions")
		registerFunc(&_cSSM_GetContext, frameworkHandle, "CSSM_GetContext")
		registerFunc(&_cSSM_GetContextAttribute, frameworkHandle, "CSSM_GetContextAttribute")
		registerFunc(&_cSSM_GetKeyAcl, frameworkHandle, "CSSM_GetKeyAcl")
		registerFunc(&_cSSM_GetKeyOwner, frameworkHandle, "CSSM_GetKeyOwner")
		registerFunc(&_cSSM_GetModuleGUIDFromHandle, frameworkHandle, "CSSM_GetModuleGUIDFromHandle")
		registerFunc(&_cSSM_GetPrivilege, frameworkHandle, "CSSM_GetPrivilege")
		registerFunc(&_cSSM_GetSubserviceUIDFromHandle, frameworkHandle, "CSSM_GetSubserviceUIDFromHandle")
		registerFunc(&_cSSM_GetTimeValue, frameworkHandle, "CSSM_GetTimeValue")
		registerFunc(&_cSSM_Init, frameworkHandle, "CSSM_Init")
		registerFunc(&_cSSM_Introduce, frameworkHandle, "CSSM_Introduce")
		registerFunc(&_cSSM_ListAttachedModuleManagers, frameworkHandle, "CSSM_ListAttachedModuleManagers")
		registerFunc(&_cSSM_ModuleAttach, frameworkHandle, "CSSM_ModuleAttach")
		registerFunc(&_cSSM_ModuleDetach, frameworkHandle, "CSSM_ModuleDetach")
		registerFunc(&_cSSM_ModuleLoad, frameworkHandle, "CSSM_ModuleLoad")
		registerFunc(&_cSSM_ModuleUnload, frameworkHandle, "CSSM_ModuleUnload")
		registerFunc(&_cSSM_QueryKeySizeInBits, frameworkHandle, "CSSM_QueryKeySizeInBits")
		registerFunc(&_cSSM_QuerySize, frameworkHandle, "CSSM_QuerySize")
		registerFunc(&_cSSM_RetrieveCounter, frameworkHandle, "CSSM_RetrieveCounter")
		registerFunc(&_cSSM_RetrieveUniqueId, frameworkHandle, "CSSM_RetrieveUniqueId")
		registerFunc(&_cSSM_SetContext, frameworkHandle, "CSSM_SetContext")
		registerFunc(&_cSSM_SetPrivilege, frameworkHandle, "CSSM_SetPrivilege")
		registerFunc(&_cSSM_SignData, frameworkHandle, "CSSM_SignData")
		registerFunc(&_cSSM_SignDataFinal, frameworkHandle, "CSSM_SignDataFinal")
		registerFunc(&_cSSM_SignDataInit, frameworkHandle, "CSSM_SignDataInit")
		registerFunc(&_cSSM_SignDataUpdate, frameworkHandle, "CSSM_SignDataUpdate")
		registerFunc(&_cSSM_TP_ApplyCrlToDb, frameworkHandle, "CSSM_TP_ApplyCrlToDb")
		registerFunc(&_cSSM_TP_CertCreateTemplate, frameworkHandle, "CSSM_TP_CertCreateTemplate")
		registerFunc(&_cSSM_TP_CertGetAllTemplateFields, frameworkHandle, "CSSM_TP_CertGetAllTemplateFields")
		registerFunc(&_cSSM_TP_CertGroupConstruct, frameworkHandle, "CSSM_TP_CertGroupConstruct")
		registerFunc(&_cSSM_TP_CertGroupPrune, frameworkHandle, "CSSM_TP_CertGroupPrune")
		registerFunc(&_cSSM_TP_CertGroupToTupleGroup, frameworkHandle, "CSSM_TP_CertGroupToTupleGroup")
		registerFunc(&_cSSM_TP_CertGroupVerify, frameworkHandle, "CSSM_TP_CertGroupVerify")
		registerFunc(&_cSSM_TP_CertReclaimAbort, frameworkHandle, "CSSM_TP_CertReclaimAbort")
		registerFunc(&_cSSM_TP_CertReclaimKey, frameworkHandle, "CSSM_TP_CertReclaimKey")
		registerFunc(&_cSSM_TP_CertRemoveFromCrlTemplate, frameworkHandle, "CSSM_TP_CertRemoveFromCrlTemplate")
		registerFunc(&_cSSM_TP_CertRevoke, frameworkHandle, "CSSM_TP_CertRevoke")
		registerFunc(&_cSSM_TP_CertSign, frameworkHandle, "CSSM_TP_CertSign")
		registerFunc(&_cSSM_TP_ConfirmCredResult, frameworkHandle, "CSSM_TP_ConfirmCredResult")
		registerFunc(&_cSSM_TP_CrlCreateTemplate, frameworkHandle, "CSSM_TP_CrlCreateTemplate")
		registerFunc(&_cSSM_TP_CrlSign, frameworkHandle, "CSSM_TP_CrlSign")
		registerFunc(&_cSSM_TP_CrlVerify, frameworkHandle, "CSSM_TP_CrlVerify")
		registerFunc(&_cSSM_TP_FormRequest, frameworkHandle, "CSSM_TP_FormRequest")
		registerFunc(&_cSSM_TP_FormSubmit, frameworkHandle, "CSSM_TP_FormSubmit")
		registerFunc(&_cSSM_TP_PassThrough, frameworkHandle, "CSSM_TP_PassThrough")
		registerFunc(&_cSSM_TP_ReceiveConfirmation, frameworkHandle, "CSSM_TP_ReceiveConfirmation")
		registerFunc(&_cSSM_TP_RetrieveCredResult, frameworkHandle, "CSSM_TP_RetrieveCredResult")
		registerFunc(&_cSSM_TP_SubmitCredRequest, frameworkHandle, "CSSM_TP_SubmitCredRequest")
		registerFunc(&_cSSM_TP_TupleGroupToCertGroup, frameworkHandle, "CSSM_TP_TupleGroupToCertGroup")
		registerFunc(&_cSSM_Terminate, frameworkHandle, "CSSM_Terminate")
		registerFunc(&_cSSM_Unintroduce, frameworkHandle, "CSSM_Unintroduce")
		registerFunc(&_cSSM_UnwrapKey, frameworkHandle, "CSSM_UnwrapKey")
		registerFunc(&_cSSM_UnwrapKeyP, frameworkHandle, "CSSM_UnwrapKeyP")
		registerFunc(&_cSSM_UpdateContextAttributes, frameworkHandle, "CSSM_UpdateContextAttributes")
		registerFunc(&_cSSM_VerifyData, frameworkHandle, "CSSM_VerifyData")
		registerFunc(&_cSSM_VerifyDataFinal, frameworkHandle, "CSSM_VerifyDataFinal")
		registerFunc(&_cSSM_VerifyDataInit, frameworkHandle, "CSSM_VerifyDataInit")
		registerFunc(&_cSSM_VerifyDataUpdate, frameworkHandle, "CSSM_VerifyDataUpdate")
		registerFunc(&_cSSM_VerifyDevice, frameworkHandle, "CSSM_VerifyDevice")
		registerFunc(&_cSSM_VerifyMac, frameworkHandle, "CSSM_VerifyMac")
		registerFunc(&_cSSM_VerifyMacFinal, frameworkHandle, "CSSM_VerifyMacFinal")
		registerFunc(&_cSSM_VerifyMacInit, frameworkHandle, "CSSM_VerifyMacInit")
		registerFunc(&_cSSM_VerifyMacUpdate, frameworkHandle, "CSSM_VerifyMacUpdate")
		registerFunc(&_cSSM_WrapKey, frameworkHandle, "CSSM_WrapKey")
		registerFunc(&_cSSM_WrapKeyP, frameworkHandle, "CSSM_WrapKeyP")
		registerFunc(&_mDS_Initialize, frameworkHandle, "MDS_Initialize")
		registerFunc(&_mDS_Install, frameworkHandle, "MDS_Install")
		registerFunc(&_mDS_Terminate, frameworkHandle, "MDS_Terminate")
		registerFunc(&_mDS_Uninstall, frameworkHandle, "MDS_Uninstall")
		registerFunc(&_secAccessControlCreateWithFlags, frameworkHandle, "SecAccessControlCreateWithFlags")
		registerFunc(&_secAccessControlGetTypeID, frameworkHandle, "SecAccessControlGetTypeID")
		registerFunc(&_secAddSharedWebCredential, frameworkHandle, "SecAddSharedWebCredential")
		registerFunc(&_secAsn1AllocCopy, frameworkHandle, "SecAsn1AllocCopy")
		registerFunc(&_secAsn1AllocCopyItem, frameworkHandle, "SecAsn1AllocCopyItem")
		registerFunc(&_secAsn1AllocItem, frameworkHandle, "SecAsn1AllocItem")
		registerFunc(&_secAsn1CoderCreate, frameworkHandle, "SecAsn1CoderCreate")
		registerFunc(&_secAsn1CoderRelease, frameworkHandle, "SecAsn1CoderRelease")
		registerFunc(&_secAsn1Decode, frameworkHandle, "SecAsn1Decode")
		registerFunc(&_secAsn1DecodeData, frameworkHandle, "SecAsn1DecodeData")
		registerFunc(&_secAsn1EncodeItem, frameworkHandle, "SecAsn1EncodeItem")
		registerFunc(&_secAsn1Malloc, frameworkHandle, "SecAsn1Malloc")
		registerFunc(&_secAsn1OidCompare, frameworkHandle, "SecAsn1OidCompare")
		registerFunc(&_secCertificateAddToKeychain, frameworkHandle, "SecCertificateAddToKeychain")
		registerFunc(&_secCertificateCopyCommonName, frameworkHandle, "SecCertificateCopyCommonName")
		registerFunc(&_secCertificateCopyData, frameworkHandle, "SecCertificateCopyData")
		registerFunc(&_secCertificateCopyEmailAddresses, frameworkHandle, "SecCertificateCopyEmailAddresses")
		registerFunc(&_secCertificateCopyKey, frameworkHandle, "SecCertificateCopyKey")
		registerFunc(&_secCertificateCopyLongDescription, frameworkHandle, "SecCertificateCopyLongDescription")
		registerFunc(&_secCertificateCopyNormalizedIssuerSequence, frameworkHandle, "SecCertificateCopyNormalizedIssuerSequence")
		registerFunc(&_secCertificateCopyNormalizedSubjectSequence, frameworkHandle, "SecCertificateCopyNormalizedSubjectSequence")
		registerFunc(&_secCertificateCopyNotValidAfterDate, frameworkHandle, "SecCertificateCopyNotValidAfterDate")
		registerFunc(&_secCertificateCopyNotValidBeforeDate, frameworkHandle, "SecCertificateCopyNotValidBeforeDate")
		registerFunc(&_secCertificateCopyPreference, frameworkHandle, "SecCertificateCopyPreference")
		registerFunc(&_secCertificateCopyPreferred, frameworkHandle, "SecCertificateCopyPreferred")
		registerFunc(&_secCertificateCopySerialNumberData, frameworkHandle, "SecCertificateCopySerialNumberData")
		registerFunc(&_secCertificateCopyShortDescription, frameworkHandle, "SecCertificateCopyShortDescription")
		registerFunc(&_secCertificateCopySubjectSummary, frameworkHandle, "SecCertificateCopySubjectSummary")
		registerFunc(&_secCertificateCopyValues, frameworkHandle, "SecCertificateCopyValues")
		registerFunc(&_secCertificateCreateWithData, frameworkHandle, "SecCertificateCreateWithData")
		registerFunc(&_secCertificateGetAlgorithmID, frameworkHandle, "SecCertificateGetAlgorithmID")
		registerFunc(&_secCertificateGetCLHandle, frameworkHandle, "SecCertificateGetCLHandle")
		registerFunc(&_secCertificateGetData, frameworkHandle, "SecCertificateGetData")
		registerFunc(&_secCertificateGetIssuer, frameworkHandle, "SecCertificateGetIssuer")
		registerFunc(&_secCertificateGetSubject, frameworkHandle, "SecCertificateGetSubject")
		registerFunc(&_secCertificateGetType, frameworkHandle, "SecCertificateGetType")
		registerFunc(&_secCertificateGetTypeID, frameworkHandle, "SecCertificateGetTypeID")
		registerFunc(&_secCertificateSetPreference, frameworkHandle, "SecCertificateSetPreference")
		registerFunc(&_secCertificateSetPreferred, frameworkHandle, "SecCertificateSetPreferred")
		registerFunc(&_secCodeCheckValidity, frameworkHandle, "SecCodeCheckValidity")
		registerFunc(&_secCodeCheckValidityWithErrors, frameworkHandle, "SecCodeCheckValidityWithErrors")
		registerFunc(&_secCodeCopyDesignatedRequirement, frameworkHandle, "SecCodeCopyDesignatedRequirement")
		registerFunc(&_secCodeCopyGuestWithAttributes, frameworkHandle, "SecCodeCopyGuestWithAttributes")
		registerFunc(&_secCodeCopyHost, frameworkHandle, "SecCodeCopyHost")
		registerFunc(&_secCodeCopyPath, frameworkHandle, "SecCodeCopyPath")
		registerFunc(&_secCodeCopySelf, frameworkHandle, "SecCodeCopySelf")
		registerFunc(&_secCodeCopySigningInformation, frameworkHandle, "SecCodeCopySigningInformation")
		registerFunc(&_secCodeCopyStaticCode, frameworkHandle, "SecCodeCopyStaticCode")
		registerFunc(&_secCodeCreateWithXPCMessage, frameworkHandle, "SecCodeCreateWithXPCMessage")
		registerFunc(&_secCodeGetTypeID, frameworkHandle, "SecCodeGetTypeID")
		registerFunc(&_secCodeMapMemory, frameworkHandle, "SecCodeMapMemory")
		registerFunc(&_secCodeValidateFileResource, frameworkHandle, "SecCodeValidateFileResource")
		registerFunc(&_secCopyErrorMessageString, frameworkHandle, "SecCopyErrorMessageString")
		registerFunc(&_secCreateSharedWebCredentialPassword, frameworkHandle, "SecCreateSharedWebCredentialPassword")
		registerFunc(&_secHostCreateGuest, frameworkHandle, "SecHostCreateGuest")
		registerFunc(&_secHostRemoveGuest, frameworkHandle, "SecHostRemoveGuest")
		registerFunc(&_secHostSelectGuest, frameworkHandle, "SecHostSelectGuest")
		registerFunc(&_secHostSelectedGuest, frameworkHandle, "SecHostSelectedGuest")
		registerFunc(&_secHostSetGuestStatus, frameworkHandle, "SecHostSetGuestStatus")
		registerFunc(&_secHostSetHostingPort, frameworkHandle, "SecHostSetHostingPort")
		registerFunc(&_secIdentityCopyCertificate, frameworkHandle, "SecIdentityCopyCertificate")
		registerFunc(&_secIdentityCopyPreference, frameworkHandle, "SecIdentityCopyPreference")
		registerFunc(&_secIdentityCopyPreferred, frameworkHandle, "SecIdentityCopyPreferred")
		registerFunc(&_secIdentityCopyPrivateKey, frameworkHandle, "SecIdentityCopyPrivateKey")
		registerFunc(&_secIdentityCopySystemIdentity, frameworkHandle, "SecIdentityCopySystemIdentity")
		registerFunc(&_secIdentityCreate, frameworkHandle, "SecIdentityCreate")
		registerFunc(&_secIdentityCreateWithCertificate, frameworkHandle, "SecIdentityCreateWithCertificate")
		registerFunc(&_secIdentityGetTypeID, frameworkHandle, "SecIdentityGetTypeID")
		registerFunc(&_secIdentitySearchCopyNext, frameworkHandle, "SecIdentitySearchCopyNext")
		registerFunc(&_secIdentitySearchCreate, frameworkHandle, "SecIdentitySearchCreate")
		registerFunc(&_secIdentitySearchGetTypeID, frameworkHandle, "SecIdentitySearchGetTypeID")
		registerFunc(&_secIdentitySetPreference, frameworkHandle, "SecIdentitySetPreference")
		registerFunc(&_secIdentitySetPreferred, frameworkHandle, "SecIdentitySetPreferred")
		registerFunc(&_secIdentitySetSystemIdentity, frameworkHandle, "SecIdentitySetSystemIdentity")
		registerFunc(&_secItemAdd, frameworkHandle, "SecItemAdd")
		registerFunc(&_secItemCopyMatching, frameworkHandle, "SecItemCopyMatching")
		registerFunc(&_secItemDelete, frameworkHandle, "SecItemDelete")
		registerFunc(&_secItemExport, frameworkHandle, "SecItemExport")
		registerFunc(&_secItemImport, frameworkHandle, "SecItemImport")
		registerFunc(&_secItemUpdate, frameworkHandle, "SecItemUpdate")
		registerFunc(&_secKeyCopyAttributes, frameworkHandle, "SecKeyCopyAttributes")
		registerFunc(&_secKeyCopyExternalRepresentation, frameworkHandle, "SecKeyCopyExternalRepresentation")
		registerFunc(&_secKeyCopyKeyExchangeResult, frameworkHandle, "SecKeyCopyKeyExchangeResult")
		registerFunc(&_secKeyCopyPublicKey, frameworkHandle, "SecKeyCopyPublicKey")
		registerFunc(&_secKeyCreateDecryptedData, frameworkHandle, "SecKeyCreateDecryptedData")
		registerFunc(&_secKeyCreateEncryptedData, frameworkHandle, "SecKeyCreateEncryptedData")
		registerFunc(&_secKeyCreatePair, frameworkHandle, "SecKeyCreatePair")
		registerFunc(&_secKeyCreateRandomKey, frameworkHandle, "SecKeyCreateRandomKey")
		registerFunc(&_secKeyCreateSignature, frameworkHandle, "SecKeyCreateSignature")
		registerFunc(&_secKeyCreateWithData, frameworkHandle, "SecKeyCreateWithData")
		registerFunc(&_secKeyGenerate, frameworkHandle, "SecKeyGenerate")
		registerFunc(&_secKeyGetBlockSize, frameworkHandle, "SecKeyGetBlockSize")
		registerFunc(&_secKeyGetCSPHandle, frameworkHandle, "SecKeyGetCSPHandle")
		registerFunc(&_secKeyGetCSSMKey, frameworkHandle, "SecKeyGetCSSMKey")
		registerFunc(&_secKeyGetCredentials, frameworkHandle, "SecKeyGetCredentials")
		registerFunc(&_secKeyGetTypeID, frameworkHandle, "SecKeyGetTypeID")
		registerFunc(&_secKeyIsAlgorithmSupported, frameworkHandle, "SecKeyIsAlgorithmSupported")
		registerFunc(&_secKeyVerifySignature, frameworkHandle, "SecKeyVerifySignature")
		registerFunc(&_secKeychainSearchGetTypeID, frameworkHandle, "SecKeychainSearchGetTypeID")
		registerFunc(&_secPKCS12Import, frameworkHandle, "SecPKCS12Import")
		registerFunc(&_secPolicyCopyProperties, frameworkHandle, "SecPolicyCopyProperties")
		registerFunc(&_secPolicyCreateBasicX509, frameworkHandle, "SecPolicyCreateBasicX509")
		registerFunc(&_secPolicyCreateRevocation, frameworkHandle, "SecPolicyCreateRevocation")
		registerFunc(&_secPolicyCreateSSL, frameworkHandle, "SecPolicyCreateSSL")
		registerFunc(&_secPolicyCreateWithOID, frameworkHandle, "SecPolicyCreateWithOID")
		registerFunc(&_secPolicyCreateWithProperties, frameworkHandle, "SecPolicyCreateWithProperties")
		registerFunc(&_secPolicyGetOID, frameworkHandle, "SecPolicyGetOID")
		registerFunc(&_secPolicyGetTPHandle, frameworkHandle, "SecPolicyGetTPHandle")
		registerFunc(&_secPolicyGetTypeID, frameworkHandle, "SecPolicyGetTypeID")
		registerFunc(&_secPolicyGetValue, frameworkHandle, "SecPolicyGetValue")
		registerFunc(&_secPolicySearchCopyNext, frameworkHandle, "SecPolicySearchCopyNext")
		registerFunc(&_secPolicySearchCreate, frameworkHandle, "SecPolicySearchCreate")
		registerFunc(&_secPolicySearchGetTypeID, frameworkHandle, "SecPolicySearchGetTypeID")
		registerFunc(&_secPolicySetProperties, frameworkHandle, "SecPolicySetProperties")
		registerFunc(&_secPolicySetValue, frameworkHandle, "SecPolicySetValue")
		registerFunc(&_secRandomCopyBytes, frameworkHandle, "SecRandomCopyBytes")
		registerFunc(&_secRequirementCopyData, frameworkHandle, "SecRequirementCopyData")
		registerFunc(&_secRequirementCopyString, frameworkHandle, "SecRequirementCopyString")
		registerFunc(&_secRequirementCreateWithData, frameworkHandle, "SecRequirementCreateWithData")
		registerFunc(&_secRequirementCreateWithString, frameworkHandle, "SecRequirementCreateWithString")
		registerFunc(&_secRequirementCreateWithStringAndErrors, frameworkHandle, "SecRequirementCreateWithStringAndErrors")
		registerFunc(&_secRequirementGetTypeID, frameworkHandle, "SecRequirementGetTypeID")
		registerFunc(&_secStaticCodeCheckValidity, frameworkHandle, "SecStaticCodeCheckValidity")
		registerFunc(&_secStaticCodeCheckValidityWithErrors, frameworkHandle, "SecStaticCodeCheckValidityWithErrors")
		registerFunc(&_secStaticCodeCreateWithPath, frameworkHandle, "SecStaticCodeCreateWithPath")
		registerFunc(&_secStaticCodeCreateWithPathAndAttributes, frameworkHandle, "SecStaticCodeCreateWithPathAndAttributes")
		registerFunc(&_secStaticCodeGetTypeID, frameworkHandle, "SecStaticCodeGetTypeID")
		registerFunc(&_secTaskCopySigningIdentifier, frameworkHandle, "SecTaskCopySigningIdentifier")
		registerFunc(&_secTaskCopyValueForEntitlement, frameworkHandle, "SecTaskCopyValueForEntitlement")
		registerFunc(&_secTaskCopyValuesForEntitlements, frameworkHandle, "SecTaskCopyValuesForEntitlements")
		registerFunc(&_secTaskCreateFromSelf, frameworkHandle, "SecTaskCreateFromSelf")
		registerFunc(&_secTaskCreateWithAuditToken, frameworkHandle, "SecTaskCreateWithAuditToken")
		registerFunc(&_secTaskGetTypeID, frameworkHandle, "SecTaskGetTypeID")
		registerFunc(&_secTrustCopyAnchorCertificates, frameworkHandle, "SecTrustCopyAnchorCertificates")
		registerFunc(&_secTrustCopyCertificateChain, frameworkHandle, "SecTrustCopyCertificateChain")
		registerFunc(&_secTrustCopyCustomAnchorCertificates, frameworkHandle, "SecTrustCopyCustomAnchorCertificates")
		registerFunc(&_secTrustCopyExceptions, frameworkHandle, "SecTrustCopyExceptions")
		registerFunc(&_secTrustCopyKey, frameworkHandle, "SecTrustCopyKey")
		registerFunc(&_secTrustCopyPolicies, frameworkHandle, "SecTrustCopyPolicies")
		registerFunc(&_secTrustCopyProperties, frameworkHandle, "SecTrustCopyProperties")
		registerFunc(&_secTrustCopyPublicKey, frameworkHandle, "SecTrustCopyPublicKey")
		registerFunc(&_secTrustCopyResult, frameworkHandle, "SecTrustCopyResult")
		registerFunc(&_secTrustCreateWithCertificates, frameworkHandle, "SecTrustCreateWithCertificates")
		registerFunc(&_secTrustEvaluateAsyncWithError, frameworkHandle, "SecTrustEvaluateAsyncWithError")
		registerFunc(&_secTrustEvaluateWithError, frameworkHandle, "SecTrustEvaluateWithError")
		registerFunc(&_secTrustGetCertificateAtIndex, frameworkHandle, "SecTrustGetCertificateAtIndex")
		registerFunc(&_secTrustGetCertificateCount, frameworkHandle, "SecTrustGetCertificateCount")
		registerFunc(&_secTrustGetCssmResult, frameworkHandle, "SecTrustGetCssmResult")
		registerFunc(&_secTrustGetCssmResultCode, frameworkHandle, "SecTrustGetCssmResultCode")
		registerFunc(&_secTrustGetNetworkFetchAllowed, frameworkHandle, "SecTrustGetNetworkFetchAllowed")
		registerFunc(&_secTrustGetResult, frameworkHandle, "SecTrustGetResult")
		registerFunc(&_secTrustGetTPHandle, frameworkHandle, "SecTrustGetTPHandle")
		registerFunc(&_secTrustGetTrustResult, frameworkHandle, "SecTrustGetTrustResult")
		registerFunc(&_secTrustGetTypeID, frameworkHandle, "SecTrustGetTypeID")
		registerFunc(&_secTrustGetVerifyTime, frameworkHandle, "SecTrustGetVerifyTime")
		registerFunc(&_secTrustSetAnchorCertificates, frameworkHandle, "SecTrustSetAnchorCertificates")
		registerFunc(&_secTrustSetAnchorCertificatesOnly, frameworkHandle, "SecTrustSetAnchorCertificatesOnly")
		registerFunc(&_secTrustSetExceptions, frameworkHandle, "SecTrustSetExceptions")
		registerFunc(&_secTrustSetKeychains, frameworkHandle, "SecTrustSetKeychains")
		registerFunc(&_secTrustSetNetworkFetchAllowed, frameworkHandle, "SecTrustSetNetworkFetchAllowed")
		registerFunc(&_secTrustSetOCSPResponse, frameworkHandle, "SecTrustSetOCSPResponse")
		registerFunc(&_secTrustSetOptions, frameworkHandle, "SecTrustSetOptions")
		registerFunc(&_secTrustSetParameters, frameworkHandle, "SecTrustSetParameters")
		registerFunc(&_secTrustSetPolicies, frameworkHandle, "SecTrustSetPolicies")
		registerFunc(&_secTrustSetSignedCertificateTimestamps, frameworkHandle, "SecTrustSetSignedCertificateTimestamps")
		registerFunc(&_secTrustSetVerifyDate, frameworkHandle, "SecTrustSetVerifyDate")
		registerFunc(&_secTrustSettingsCopyCertificates, frameworkHandle, "SecTrustSettingsCopyCertificates")
		registerFunc(&_secTrustSettingsCopyModificationDate, frameworkHandle, "SecTrustSettingsCopyModificationDate")
		registerFunc(&_secTrustSettingsCopyTrustSettings, frameworkHandle, "SecTrustSettingsCopyTrustSettings")
		registerFunc(&_secTrustSettingsCreateExternalRepresentation, frameworkHandle, "SecTrustSettingsCreateExternalRepresentation")
		registerFunc(&_secTrustSettingsImportExternalRepresentation, frameworkHandle, "SecTrustSettingsImportExternalRepresentation")
		registerFunc(&_secTrustSettingsRemoveTrustSettings, frameworkHandle, "SecTrustSettingsRemoveTrustSettings")
		registerFunc(&_secTrustSettingsSetTrustSettings, frameworkHandle, "SecTrustSettingsSetTrustSettings")
		registerFunc(&_secureDownloadCopyCreationDate, frameworkHandle, "SecureDownloadCopyCreationDate")
		registerFunc(&_secureDownloadCopyName, frameworkHandle, "SecureDownloadCopyName")
		registerFunc(&_secureDownloadCopyTicketLocation, frameworkHandle, "SecureDownloadCopyTicketLocation")
		registerFunc(&_secureDownloadCopyURLs, frameworkHandle, "SecureDownloadCopyURLs")
		registerFunc(&_secureDownloadCreateWithTicket, frameworkHandle, "SecureDownloadCreateWithTicket")
		registerFunc(&_secureDownloadFinished, frameworkHandle, "SecureDownloadFinished")
		registerFunc(&_secureDownloadGetDownloadSize, frameworkHandle, "SecureDownloadGetDownloadSize")
		registerFunc(&_secureDownloadRelease, frameworkHandle, "SecureDownloadRelease")
		registerFunc(&_secureDownloadUpdateWithData, frameworkHandle, "SecureDownloadUpdateWithData")
		registerFunc(&_sessionCreate, frameworkHandle, "SessionCreate")
		registerFunc(&_sessionGetInfo, frameworkHandle, "SessionGetInfo")
		registerFunc(&_cssmAlgToOid, frameworkHandle, "cssmAlgToOid")
		registerFunc(&_cssmOidToAlg, frameworkHandle, "cssmOidToAlg")
		registerFunc(&_cssmPerror, frameworkHandle, "cssmPerror")
		registerFunc(&_sec_certificate_copy_ref, frameworkHandle, "sec_certificate_copy_ref")
		registerFunc(&_sec_certificate_create, frameworkHandle, "sec_certificate_create")
		registerFunc(&_sec_identity_access_certificates, frameworkHandle, "sec_identity_access_certificates")
		registerFunc(&_sec_identity_copy_certificates_ref, frameworkHandle, "sec_identity_copy_certificates_ref")
		registerFunc(&_sec_identity_copy_ref, frameworkHandle, "sec_identity_copy_ref")
		registerFunc(&_sec_identity_create, frameworkHandle, "sec_identity_create")
		registerFunc(&_sec_identity_create_with_certificates, frameworkHandle, "sec_identity_create_with_certificates")
		registerFunc(&_sec_protocol_metadata_access_distinguished_names, frameworkHandle, "sec_protocol_metadata_access_distinguished_names")
		registerFunc(&_sec_protocol_metadata_access_ocsp_response, frameworkHandle, "sec_protocol_metadata_access_ocsp_response")
		registerFunc(&_sec_protocol_metadata_access_peer_certificate_chain, frameworkHandle, "sec_protocol_metadata_access_peer_certificate_chain")
		registerFunc(&_sec_protocol_metadata_access_pre_shared_keys, frameworkHandle, "sec_protocol_metadata_access_pre_shared_keys")
		registerFunc(&_sec_protocol_metadata_access_supported_signature_algorithms, frameworkHandle, "sec_protocol_metadata_access_supported_signature_algorithms")
		registerFunc(&_sec_protocol_metadata_challenge_parameters_are_equal, frameworkHandle, "sec_protocol_metadata_challenge_parameters_are_equal")
		registerFunc(&_sec_protocol_metadata_copy_negotiated_protocol, frameworkHandle, "sec_protocol_metadata_copy_negotiated_protocol")
		registerFunc(&_sec_protocol_metadata_copy_peer_public_key, frameworkHandle, "sec_protocol_metadata_copy_peer_public_key")
		registerFunc(&_sec_protocol_metadata_copy_server_name, frameworkHandle, "sec_protocol_metadata_copy_server_name")
		registerFunc(&_sec_protocol_metadata_create_secret, frameworkHandle, "sec_protocol_metadata_create_secret")
		registerFunc(&_sec_protocol_metadata_create_secret_with_context, frameworkHandle, "sec_protocol_metadata_create_secret_with_context")
		registerFunc(&_sec_protocol_metadata_get_early_data_accepted, frameworkHandle, "sec_protocol_metadata_get_early_data_accepted")
		registerFunc(&_sec_protocol_metadata_get_negotiated_ciphersuite, frameworkHandle, "sec_protocol_metadata_get_negotiated_ciphersuite")
		registerFunc(&_sec_protocol_metadata_get_negotiated_protocol, frameworkHandle, "sec_protocol_metadata_get_negotiated_protocol")
		registerFunc(&_sec_protocol_metadata_get_negotiated_protocol_version, frameworkHandle, "sec_protocol_metadata_get_negotiated_protocol_version")
		registerFunc(&_sec_protocol_metadata_get_negotiated_tls_ciphersuite, frameworkHandle, "sec_protocol_metadata_get_negotiated_tls_ciphersuite")
		registerFunc(&_sec_protocol_metadata_get_negotiated_tls_protocol_version, frameworkHandle, "sec_protocol_metadata_get_negotiated_tls_protocol_version")
		registerFunc(&_sec_protocol_metadata_get_server_name, frameworkHandle, "sec_protocol_metadata_get_server_name")
		registerFunc(&_sec_protocol_metadata_peers_are_equal, frameworkHandle, "sec_protocol_metadata_peers_are_equal")
		registerFunc(&_sec_protocol_options_add_pre_shared_key, frameworkHandle, "sec_protocol_options_add_pre_shared_key")
		registerFunc(&_sec_protocol_options_add_tls_application_protocol, frameworkHandle, "sec_protocol_options_add_tls_application_protocol")
		registerFunc(&_sec_protocol_options_append_tls_ciphersuite, frameworkHandle, "sec_protocol_options_append_tls_ciphersuite")
		registerFunc(&_sec_protocol_options_append_tls_ciphersuite_group, frameworkHandle, "sec_protocol_options_append_tls_ciphersuite_group")
		registerFunc(&_sec_protocol_options_are_equal, frameworkHandle, "sec_protocol_options_are_equal")
		registerFunc(&_sec_protocol_options_get_default_max_dtls_protocol_version, frameworkHandle, "sec_protocol_options_get_default_max_dtls_protocol_version")
		registerFunc(&_sec_protocol_options_get_default_max_tls_protocol_version, frameworkHandle, "sec_protocol_options_get_default_max_tls_protocol_version")
		registerFunc(&_sec_protocol_options_get_default_min_dtls_protocol_version, frameworkHandle, "sec_protocol_options_get_default_min_dtls_protocol_version")
		registerFunc(&_sec_protocol_options_get_default_min_tls_protocol_version, frameworkHandle, "sec_protocol_options_get_default_min_tls_protocol_version")
		registerFunc(&_sec_protocol_options_get_enable_encrypted_client_hello, frameworkHandle, "sec_protocol_options_get_enable_encrypted_client_hello")
		registerFunc(&_sec_protocol_options_get_quic_use_legacy_codepoint, frameworkHandle, "sec_protocol_options_get_quic_use_legacy_codepoint")
		registerFunc(&_sec_protocol_options_set_challenge_block, frameworkHandle, "sec_protocol_options_set_challenge_block")
		registerFunc(&_sec_protocol_options_set_enable_encrypted_client_hello, frameworkHandle, "sec_protocol_options_set_enable_encrypted_client_hello")
		registerFunc(&_sec_protocol_options_set_key_update_block, frameworkHandle, "sec_protocol_options_set_key_update_block")
		registerFunc(&_sec_protocol_options_set_local_identity, frameworkHandle, "sec_protocol_options_set_local_identity")
		registerFunc(&_sec_protocol_options_set_max_tls_protocol_version, frameworkHandle, "sec_protocol_options_set_max_tls_protocol_version")
		registerFunc(&_sec_protocol_options_set_min_tls_protocol_version, frameworkHandle, "sec_protocol_options_set_min_tls_protocol_version")
		registerFunc(&_sec_protocol_options_set_peer_authentication_optional, frameworkHandle, "sec_protocol_options_set_peer_authentication_optional")
		registerFunc(&_sec_protocol_options_set_peer_authentication_required, frameworkHandle, "sec_protocol_options_set_peer_authentication_required")
		registerFunc(&_sec_protocol_options_set_pre_shared_key_selection_block, frameworkHandle, "sec_protocol_options_set_pre_shared_key_selection_block")
		registerFunc(&_sec_protocol_options_set_quic_use_legacy_codepoint, frameworkHandle, "sec_protocol_options_set_quic_use_legacy_codepoint")
		registerFunc(&_sec_protocol_options_set_tls_false_start_enabled, frameworkHandle, "sec_protocol_options_set_tls_false_start_enabled")
		registerFunc(&_sec_protocol_options_set_tls_is_fallback_attempt, frameworkHandle, "sec_protocol_options_set_tls_is_fallback_attempt")
		registerFunc(&_sec_protocol_options_set_tls_max_version, frameworkHandle, "sec_protocol_options_set_tls_max_version")
		registerFunc(&_sec_protocol_options_set_tls_min_version, frameworkHandle, "sec_protocol_options_set_tls_min_version")
		registerFunc(&_sec_protocol_options_set_tls_ocsp_enabled, frameworkHandle, "sec_protocol_options_set_tls_ocsp_enabled")
		registerFunc(&_sec_protocol_options_set_tls_pre_shared_key_identity_hint, frameworkHandle, "sec_protocol_options_set_tls_pre_shared_key_identity_hint")
		registerFunc(&_sec_protocol_options_set_tls_renegotiation_enabled, frameworkHandle, "sec_protocol_options_set_tls_renegotiation_enabled")
		registerFunc(&_sec_protocol_options_set_tls_resumption_enabled, frameworkHandle, "sec_protocol_options_set_tls_resumption_enabled")
		registerFunc(&_sec_protocol_options_set_tls_sct_enabled, frameworkHandle, "sec_protocol_options_set_tls_sct_enabled")
		registerFunc(&_sec_protocol_options_set_tls_server_name, frameworkHandle, "sec_protocol_options_set_tls_server_name")
		registerFunc(&_sec_protocol_options_set_tls_tickets_enabled, frameworkHandle, "sec_protocol_options_set_tls_tickets_enabled")
		registerFunc(&_sec_protocol_options_set_verify_block, frameworkHandle, "sec_protocol_options_set_verify_block")
		registerFunc(&_sec_release, frameworkHandle, "sec_release")
		registerFunc(&_sec_retain, frameworkHandle, "sec_retain")
		registerFunc(&_sec_trust_copy_ref, frameworkHandle, "sec_trust_copy_ref")
		registerFunc(&_sec_trust_create, frameworkHandle, "sec_trust_create")
	}

