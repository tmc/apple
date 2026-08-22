// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/CryptoTokenKit/TKError/Code
type TKErrorCode int

const (
	TKErrorAuthenticationFailed TKErrorCode = -5
	// TKErrorCodeAuthenticationFailed: Authentication failed.
	TKErrorCodeAuthenticationFailed TKErrorCode = -5
	// TKErrorCodeAuthenticationNeeded: Authentication is needed.
	TKErrorCodeAuthenticationNeeded TKErrorCode = -9
	// TKErrorCodeBadParameter: An invalid parameter was provided.
	TKErrorCodeBadParameter TKErrorCode = -8
	// TKErrorCodeCanceledByUser: The operation was canceled by the user.
	TKErrorCodeCanceledByUser TKErrorCode = -4
	// TKErrorCodeCommunicationError: A communication error occurred.
	TKErrorCodeCommunicationError TKErrorCode = -2
	// TKErrorCodeCorruptedData: The data was corrupted.
	TKErrorCodeCorruptedData TKErrorCode = -3
	// TKErrorCodeNotImplemented: The functionality is not implemented.
	TKErrorCodeNotImplemented TKErrorCode = -1
	// TKErrorCodeObjectNotFound: The object was not found.
	TKErrorCodeObjectNotFound TKErrorCode = -6
	// TKErrorCodeTokenNotFound: The token was not found.
	TKErrorCodeTokenNotFound TKErrorCode = -7
	TKErrorObjectNotFound    TKErrorCode = -6
	TKErrorTokenNotFound     TKErrorCode = -7
)

func (e TKErrorCode) String() string {
	switch e {
	case TKErrorAuthenticationFailed:
		return "TKErrorAuthenticationFailed"
	case TKErrorCodeAuthenticationNeeded:
		return "TKErrorCodeAuthenticationNeeded"
	case TKErrorCodeBadParameter:
		return "TKErrorCodeBadParameter"
	case TKErrorCodeCanceledByUser:
		return "TKErrorCodeCanceledByUser"
	case TKErrorCodeCommunicationError:
		return "TKErrorCodeCommunicationError"
	case TKErrorCodeCorruptedData:
		return "TKErrorCodeCorruptedData"
	case TKErrorCodeNotImplemented:
		return "TKErrorCodeNotImplemented"
	case TKErrorCodeObjectNotFound:
		return "TKErrorCodeObjectNotFound"
	case TKErrorCodeTokenNotFound:
		return "TKErrorCodeTokenNotFound"
	default:
		return fmt.Sprintf("TKErrorCode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/Charset-swift.enum
type TKSmartCardPINCharset int

const (
	TKSmartCardPINCharsetAlphanumeric      TKSmartCardPINCharset = 1
	TKSmartCardPINCharsetNumeric           TKSmartCardPINCharset = 0
	TKSmartCardPINCharsetUpperAlphanumeric TKSmartCardPINCharset = 2
)

func (e TKSmartCardPINCharset) String() string {
	switch e {
	case TKSmartCardPINCharsetAlphanumeric:
		return "TKSmartCardPINCharsetAlphanumeric"
	case TKSmartCardPINCharsetNumeric:
		return "TKSmartCardPINCharsetNumeric"
	case TKSmartCardPINCharsetUpperAlphanumeric:
		return "TKSmartCardPINCharsetUpperAlphanumeric"
	default:
		return fmt.Sprintf("TKSmartCardPINCharset(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionForPINOperation/Completion
type TKSmartCardPINCompletion uint

const (
	TKSmartCardPINCompletionKey       TKSmartCardPINCompletion = 2
	TKSmartCardPINCompletionMaxLength TKSmartCardPINCompletion = 1
	TKSmartCardPINCompletionTimeout   TKSmartCardPINCompletion = 4
)

func (e TKSmartCardPINCompletion) String() string {
	switch e {
	case TKSmartCardPINCompletionKey:
		return "TKSmartCardPINCompletionKey"
	case TKSmartCardPINCompletionMaxLength:
		return "TKSmartCardPINCompletionMaxLength"
	case TKSmartCardPINCompletionTimeout:
		return "TKSmartCardPINCompletionTimeout"
	default:
		return fmt.Sprintf("TKSmartCardPINCompletion(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionForSecurePINChange/Confirmation
type TKSmartCardPINConfirmation uint

const (
	TKSmartCardPINConfirmationCurrent TKSmartCardPINConfirmation = 2
	TKSmartCardPINConfirmationNew     TKSmartCardPINConfirmation = 1
	TKSmartCardPINConfirmationNone    TKSmartCardPINConfirmation = 0
)

func (e TKSmartCardPINConfirmation) String() string {
	switch e {
	case TKSmartCardPINConfirmationCurrent:
		return "TKSmartCardPINConfirmationCurrent"
	case TKSmartCardPINConfirmationNew:
		return "TKSmartCardPINConfirmationNew"
	case TKSmartCardPINConfirmationNone:
		return "TKSmartCardPINConfirmationNone"
	default:
		return fmt.Sprintf("TKSmartCardPINConfirmation(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/Encoding-swift.enum
type TKSmartCardPINEncoding int

const (
	TKSmartCardPINEncodingASCII  TKSmartCardPINEncoding = 1
	TKSmartCardPINEncodingBCD    TKSmartCardPINEncoding = 2
	TKSmartCardPINEncodingBinary TKSmartCardPINEncoding = 0
)

func (e TKSmartCardPINEncoding) String() string {
	switch e {
	case TKSmartCardPINEncodingASCII:
		return "TKSmartCardPINEncodingASCII"
	case TKSmartCardPINEncodingBCD:
		return "TKSmartCardPINEncodingBCD"
	case TKSmartCardPINEncodingBinary:
		return "TKSmartCardPINEncodingBinary"
	default:
		return fmt.Sprintf("TKSmartCardPINEncoding(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/Justification
type TKSmartCardPINJustification int

const (
	TKSmartCardPINJustificationLeft  TKSmartCardPINJustification = 0
	TKSmartCardPINJustificationRight TKSmartCardPINJustification = 1
)

func (e TKSmartCardPINJustification) String() string {
	switch e {
	case TKSmartCardPINJustificationLeft:
		return "TKSmartCardPINJustificationLeft"
	case TKSmartCardPINJustificationRight:
		return "TKSmartCardPINJustificationRight"
	default:
		return fmt.Sprintf("TKSmartCardPINJustification(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardProtocol
type TKSmartCardProtocol uint

const (
	// TKSmartCardProtocolAny: # Discussion
	TKSmartCardProtocolAny  TKSmartCardProtocol = 65535
	TKSmartCardProtocolNone TKSmartCardProtocol = 0
	// TKSmartCardProtocolT0: # Discussion
	TKSmartCardProtocolT0 TKSmartCardProtocol = 1
	// TKSmartCardProtocolT1: # Discussion
	TKSmartCardProtocolT1 TKSmartCardProtocol = 2
	// TKSmartCardProtocolT15: # Discussion
	TKSmartCardProtocolT15 TKSmartCardProtocol = 32768
)

func (e TKSmartCardProtocol) String() string {
	switch e {
	case TKSmartCardProtocolAny:
		return "TKSmartCardProtocolAny"
	case TKSmartCardProtocolNone:
		return "TKSmartCardProtocolNone"
	case TKSmartCardProtocolT0:
		return "TKSmartCardProtocolT0"
	case TKSmartCardProtocolT1:
		return "TKSmartCardProtocolT1"
	case TKSmartCardProtocolT15:
		return "TKSmartCardProtocolT15"
	default:
		return fmt.Sprintf("TKSmartCardProtocol(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlot/State-swift.enum
type TKSmartCardSlotState int

const (
	TKSmartCardSlotStateEmpty     TKSmartCardSlotState = 1
	TKSmartCardSlotStateMissing   TKSmartCardSlotState = 0
	TKSmartCardSlotStateMuteCard  TKSmartCardSlotState = 3
	TKSmartCardSlotStateProbing   TKSmartCardSlotState = 2
	TKSmartCardSlotStateValidCard TKSmartCardSlotState = 4
)

func (e TKSmartCardSlotState) String() string {
	switch e {
	case TKSmartCardSlotStateEmpty:
		return "TKSmartCardSlotStateEmpty"
	case TKSmartCardSlotStateMissing:
		return "TKSmartCardSlotStateMissing"
	case TKSmartCardSlotStateMuteCard:
		return "TKSmartCardSlotStateMuteCard"
	case TKSmartCardSlotStateProbing:
		return "TKSmartCardSlotStateProbing"
	case TKSmartCardSlotStateValidCard:
		return "TKSmartCardSlotStateValidCard"
	default:
		return fmt.Sprintf("TKSmartCardSlotState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenOperation
type TKTokenOperation int

const (
	TKTokenOperationDecryptData        TKTokenOperation = 3
	TKTokenOperationNone               TKTokenOperation = 0
	TKTokenOperationPerformKeyExchange TKTokenOperation = 4
	TKTokenOperationReadData           TKTokenOperation = 1
	TKTokenOperationSignData           TKTokenOperation = 2
)

func (e TKTokenOperation) String() string {
	switch e {
	case TKTokenOperationDecryptData:
		return "TKTokenOperationDecryptData"
	case TKTokenOperationNone:
		return "TKTokenOperationNone"
	case TKTokenOperationPerformKeyExchange:
		return "TKTokenOperationPerformKeyExchange"
	case TKTokenOperationReadData:
		return "TKTokenOperationReadData"
	case TKTokenOperationSignData:
		return "TKTokenOperationSignData"
	default:
		return fmt.Sprintf("TKTokenOperation(%d)", e)
	}
}
