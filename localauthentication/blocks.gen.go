// Code generated from Apple documentation. DO NOT EDIT.

package localauthentication

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// BoolErrorHandler handles A closure that is executed when policy evaluation finishes.
//   - success: [true](<doc://com.apple.documentation/documentation/Swift/true>) if policy evaluation succeeded, otherwise [false](<doc://com.apple.documentation/documentation/Swift/false>).
//   - error: `nil` if policy evaluation succeeded, an error object that should be presented to the user otherwise. See [LAError.Code](<doc://com.apple.localauthentication/documentation/LocalAuthentication/LAError-swift.struct/Code>) for possible error codes
//
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [LAContext.EvaluateAccessControlOperationLocalizedReasonReply]
//   - [LAContext.EvaluatePolicyLocalizedReasonReply]
type BoolErrorHandler = func(bool, error)

// NewBoolErrorBlock wraps a Go [BoolErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [LAContext.EvaluateAccessControlOperationLocalizedReasonReply]
//   - [LAContext.EvaluatePolicyLocalizedReasonReply]
func NewBoolErrorBlock(handler BoolErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal bool, errID objc.ID) {
		handler(primitiveVal, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// DataErrorHandler handles A completion handler to call when the decryption operation completes.
//   - data: The result of the key exchange operation.
//   - error: An error object that indicates why the key exchange failed, or `nil` if the exchange succeeded.
//
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [LAPrivateKey.DecryptDataSecKeyAlgorithmCompletion]
//   - [LAPrivateKey.ExchangeKeysWithPublicKeySecKeyAlgorithmSecKeyParametersCompletion]
//   - [LAPrivateKey.SignDataSecKeyAlgorithmCompletion]
//   - [LAPublicKey.EncryptDataSecKeyAlgorithmCompletion]
//   - [LAPublicKey.ExportBytesWithCompletion]
//   - [LASecret.LoadDataWithCompletion]
type DataErrorHandler = func(*foundation.NSData, error)

// NewDataErrorBlock wraps a Go [DataErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [LAPrivateKey.DecryptDataSecKeyAlgorithmCompletion]
//   - [LAPrivateKey.ExchangeKeysWithPublicKeySecKeyAlgorithmSecKeyParametersCompletion]
//   - [LAPrivateKey.SignDataSecKeyAlgorithmCompletion]
//   - [LAPublicKey.EncryptDataSecKeyAlgorithmCompletion]
//   - [LAPublicKey.ExportBytesWithCompletion]
//   - [LASecret.LoadDataWithCompletion]
func NewDataErrorBlock(handler DataErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *foundation.NSData
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := foundation.NSDataFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// ErrorHandler handles A completion handler to call when the verification operation completes.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [LAPublicKey.VerifyDataSignatureSecKeyAlgorithmCompletion]
//   - [LARight.AuthorizeWithLocalizedReasonCompletion]
//   - [LARight.AuthorizeWithLocalizedReasonInPresentationContextCompletion]
//   - [LARight.CheckCanAuthorizeWithCompletion]
//   - [LARightStore.RemoveAllRightsWithCompletion]
//   - [LARightStore.RemoveRightCompletion]
//   - [LARightStore.RemoveRightForIdentifierCompletion]
type ErrorHandler = func(error)

// NewErrorBlock wraps a Go [ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [LAPublicKey.VerifyDataSignatureSecKeyAlgorithmCompletion]
//   - [LARight.AuthorizeWithLocalizedReasonCompletion]
//   - [LARight.AuthorizeWithLocalizedReasonInPresentationContextCompletion]
//   - [LARight.CheckCanAuthorizeWithCompletion]
//   - [LARightStore.RemoveAllRightsWithCompletion]
//   - [LARightStore.RemoveRightCompletion]
//   - [LARightStore.RemoveRightForIdentifierCompletion]
func NewErrorBlock(handler ErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, errID objc.ID) {
		handler(foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// LAPersistedRightErrorHandler handles A completion handler to call when the right access completes.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [LARightStore.RightForIdentifierCompletion]
//   - [LARightStore.SaveRightIdentifierCompletion]
//   - [LARightStore.SaveRightIdentifierSecretCompletion]
type LAPersistedRightErrorHandler = func(*LAPersistedRight, error)

// NewLAPersistedRightErrorBlock wraps a Go [LAPersistedRightErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [LARightStore.RightForIdentifierCompletion]
//   - [LARightStore.SaveRightIdentifierCompletion]
//   - [LARightStore.SaveRightIdentifierSecretCompletion]
func NewLAPersistedRightErrorBlock(handler LAPersistedRightErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *LAPersistedRight
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := LAPersistedRightFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// VoidHandler is the signature for a completion handler block.
//
// Used by:
//   - [LARight.DeauthorizeWithCompletion]
type VoidHandler = func()

// NewVoidBlock wraps a Go [VoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [LARight.DeauthorizeWithCompletion]
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}
