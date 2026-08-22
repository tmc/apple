// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/CoreSpotlight/CSIndexError/Code
type CSIndexErrorCode int

const (
	// CSIndexErrorCodeIndexUnavailableError: The indexer is unavailable.
	CSIndexErrorCodeIndexUnavailableError CSIndexErrorCode = -1000
	// CSIndexErrorCodeIndexingUnsupported: Indexing isn’t supported on the device.
	CSIndexErrorCodeIndexingUnsupported CSIndexErrorCode = -1005
	// CSIndexErrorCodeInvalidClientStateError: The provided client state data is invalid.
	CSIndexErrorCodeInvalidClientStateError CSIndexErrorCode = -1002
	// CSIndexErrorCodeInvalidItemError: The searchable item object is invalid.
	CSIndexErrorCodeInvalidItemError CSIndexErrorCode = -1001
	// CSIndexErrorCodeMismatchedClientState: The provided client state did not match the information in the index.
	CSIndexErrorCodeMismatchedClientState CSIndexErrorCode = -1006
	// CSIndexErrorCodeQuotaExceeded: The quota for the bundle has been exceeded.
	CSIndexErrorCodeQuotaExceeded CSIndexErrorCode = -1004
	// CSIndexErrorCodeRemoteConnectionError: An error occurred while communicating with the remote process.
	CSIndexErrorCodeRemoteConnectionError CSIndexErrorCode = -1003
	// CSIndexErrorCodeUnknownError: An unknown error occurred.
	CSIndexErrorCodeUnknownError CSIndexErrorCode = -1
)

func (e CSIndexErrorCode) String() string {
	switch e {
	case CSIndexErrorCodeIndexUnavailableError:
		return "CSIndexErrorCodeIndexUnavailableError"
	case CSIndexErrorCodeIndexingUnsupported:
		return "CSIndexErrorCodeIndexingUnsupported"
	case CSIndexErrorCodeInvalidClientStateError:
		return "CSIndexErrorCodeInvalidClientStateError"
	case CSIndexErrorCodeInvalidItemError:
		return "CSIndexErrorCodeInvalidItemError"
	case CSIndexErrorCodeMismatchedClientState:
		return "CSIndexErrorCodeMismatchedClientState"
	case CSIndexErrorCodeQuotaExceeded:
		return "CSIndexErrorCodeQuotaExceeded"
	case CSIndexErrorCodeRemoteConnectionError:
		return "CSIndexErrorCodeRemoteConnectionError"
	case CSIndexErrorCodeUnknownError:
		return "CSIndexErrorCodeUnknownError"
	default:
		return fmt.Sprintf("CSIndexErrorCode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQueryError/Code
type CSSearchQueryErrorCode int

const (
	// CSSearchQueryErrorCodeCancelled: The query stopped because someone canceled it.
	CSSearchQueryErrorCodeCancelled CSSearchQueryErrorCode = -2003
	// CSSearchQueryErrorCodeIndexUnreachable: The index is unreachable.
	CSSearchQueryErrorCodeIndexUnreachable CSSearchQueryErrorCode = -2001
	// CSSearchQueryErrorCodeInvalidQuery: The query is syntactically invalid or specifies items that your app doesn’t have access to.
	CSSearchQueryErrorCodeInvalidQuery CSSearchQueryErrorCode = -2002
	// CSSearchQueryErrorCodeUnknown: An unknown error occurred.
	CSSearchQueryErrorCodeUnknown CSSearchQueryErrorCode = -2000
)

func (e CSSearchQueryErrorCode) String() string {
	switch e {
	case CSSearchQueryErrorCodeCancelled:
		return "CSSearchQueryErrorCodeCancelled"
	case CSSearchQueryErrorCodeIndexUnreachable:
		return "CSSearchQueryErrorCodeIndexUnreachable"
	case CSSearchQueryErrorCodeInvalidQuery:
		return "CSSearchQueryErrorCodeInvalidQuery"
	case CSSearchQueryErrorCodeUnknown:
		return "CSSearchQueryErrorCodeUnknown"
	default:
		return fmt.Sprintf("CSSearchQueryErrorCode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQueryContext/SourceOptions-swift.struct
type CSSearchQuerySourceOptions uint

const (
	// CSSearchQuerySourceOptionAllowMail: The query allows Mail messages in the search.
	CSSearchQuerySourceOptionAllowMail CSSearchQuerySourceOptions = 1
	// CSSearchQuerySourceOptionDefault: The query uses the default search option that excludes Mail messages.
	CSSearchQuerySourceOptionDefault CSSearchQuerySourceOptions = 0
)

func (e CSSearchQuerySourceOptions) String() string {
	switch e {
	case CSSearchQuerySourceOptionAllowMail:
		return "CSSearchQuerySourceOptionAllowMail"
	case CSSearchQuerySourceOptionDefault:
		return "CSSearchQuerySourceOptionDefault"
	default:
		return fmt.Sprintf("CSSearchQuerySourceOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/UpdateListenerOptions-swift.struct
type CSSearchableItemUpdateListenerOptions uint

const (
	CSSearchableItemUpdateListenerOptionDefault CSSearchableItemUpdateListenerOptions = 0
	// CSSearchableItemUpdateListenerOptionPriority: A value that describes the listener priority options.
	CSSearchableItemUpdateListenerOptionPriority CSSearchableItemUpdateListenerOptions = 4
	// CSSearchableItemUpdateListenerOptionSummarization: A value that describes the listener summarization options.
	CSSearchableItemUpdateListenerOptionSummarization CSSearchableItemUpdateListenerOptions = 2
)

func (e CSSearchableItemUpdateListenerOptions) String() string {
	switch e {
	case CSSearchableItemUpdateListenerOptionDefault:
		return "CSSearchableItemUpdateListenerOptionDefault"
	case CSSearchableItemUpdateListenerOptionPriority:
		return "CSSearchableItemUpdateListenerOptionPriority"
	case CSSearchableItemUpdateListenerOptionSummarization:
		return "CSSearchableItemUpdateListenerOptionSummarization"
	default:
		return fmt.Sprintf("CSSearchableItemUpdateListenerOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreSpotlight/CSSuggestion/SuggestionKind-swift.enum
type CSSuggestionKind int

const (
	// CSSuggestionKindCustom: Sorts the custom suggestions together.
	CSSuggestionKindCustom CSSuggestionKind = 1
	// CSSuggestionKindDefault: Displays the suggestion normally.
	CSSuggestionKindDefault CSSuggestionKind = 2
	// CSSuggestionKindNone: Blocks the system from displaying the suggestion.
	CSSuggestionKindNone CSSuggestionKind = 0
)

func (e CSSuggestionKind) String() string {
	switch e {
	case CSSuggestionKindCustom:
		return "CSSuggestionKindCustom"
	case CSSuggestionKindDefault:
		return "CSSuggestionKindDefault"
	case CSSuggestionKindNone:
		return "CSSuggestionKindNone"
	default:
		return fmt.Sprintf("CSSuggestionKind(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreSpotlight/CSUserQuery/UserInteractionKind
type CSUserInteraction int

const (
	CSUserInteractionDefault CSUserInteraction = 0
	CSUserInteractionFocus   CSUserInteraction = 1
	CSUserInteractionSelect  CSUserInteraction = 0
)

func (e CSUserInteraction) String() string {
	switch e {
	case CSUserInteractionDefault:
		return "CSUserInteractionDefault"
	case CSUserInteractionFocus:
		return "CSUserInteractionFocus"
	default:
		return fmt.Sprintf("CSUserInteraction(%d)", e)
	}
}
