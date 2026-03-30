// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/WebKit/DOMEventExceptionCode
type DOMEventExceptionCode uint32

const (
	DOM_UNSPECIFIED_EVENT_TYPE_ERR DOMEventExceptionCode = 0
)

func (e DOMEventExceptionCode) String() string {
	switch e {
	case DOM_UNSPECIFIED_EVENT_TYPE_ERR:
		return "DOM_UNSPECIFIED_EVENT_TYPE_ERR"
	default:
		return fmt.Sprintf("DOMEventExceptionCode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/DOMExceptionCode
type DOMExceptionCode uint32

const (
	DOM_DOMSTRING_SIZE_ERR          DOMExceptionCode = 2
	DOM_HIERARCHY_REQUEST_ERR       DOMExceptionCode = 3
	DOM_INDEX_SIZE_ERR              DOMExceptionCode = 1
	DOM_INUSE_ATTRIBUTE_ERR         DOMExceptionCode = 10
	DOM_INVALID_ACCESS_ERR          DOMExceptionCode = 15
	DOM_INVALID_CHARACTER_ERR       DOMExceptionCode = 5
	DOM_INVALID_MODIFICATION_ERR    DOMExceptionCode = 13
	DOM_INVALID_STATE_ERR           DOMExceptionCode = 11
	DOM_NAMESPACE_ERR               DOMExceptionCode = 14
	DOM_NOT_FOUND_ERR               DOMExceptionCode = 8
	DOM_NOT_SUPPORTED_ERR           DOMExceptionCode = 9
	DOM_NO_DATA_ALLOWED_ERR         DOMExceptionCode = 6
	DOM_NO_MODIFICATION_ALLOWED_ERR DOMExceptionCode = 7
	DOM_SYNTAX_ERR                  DOMExceptionCode = 12
	DOM_WRONG_DOCUMENT_ERR          DOMExceptionCode = 4
)

func (e DOMExceptionCode) String() string {
	switch e {
	case DOM_DOMSTRING_SIZE_ERR:
		return "DOM_DOMSTRING_SIZE_ERR"
	case DOM_HIERARCHY_REQUEST_ERR:
		return "DOM_HIERARCHY_REQUEST_ERR"
	case DOM_INDEX_SIZE_ERR:
		return "DOM_INDEX_SIZE_ERR"
	case DOM_INUSE_ATTRIBUTE_ERR:
		return "DOM_INUSE_ATTRIBUTE_ERR"
	case DOM_INVALID_ACCESS_ERR:
		return "DOM_INVALID_ACCESS_ERR"
	case DOM_INVALID_CHARACTER_ERR:
		return "DOM_INVALID_CHARACTER_ERR"
	case DOM_INVALID_MODIFICATION_ERR:
		return "DOM_INVALID_MODIFICATION_ERR"
	case DOM_INVALID_STATE_ERR:
		return "DOM_INVALID_STATE_ERR"
	case DOM_NAMESPACE_ERR:
		return "DOM_NAMESPACE_ERR"
	case DOM_NOT_FOUND_ERR:
		return "DOM_NOT_FOUND_ERR"
	case DOM_NOT_SUPPORTED_ERR:
		return "DOM_NOT_SUPPORTED_ERR"
	case DOM_NO_DATA_ALLOWED_ERR:
		return "DOM_NO_DATA_ALLOWED_ERR"
	case DOM_NO_MODIFICATION_ALLOWED_ERR:
		return "DOM_NO_MODIFICATION_ALLOWED_ERR"
	case DOM_SYNTAX_ERR:
		return "DOM_SYNTAX_ERR"
	case DOM_WRONG_DOCUMENT_ERR:
		return "DOM_WRONG_DOCUMENT_ERR"
	default:
		return fmt.Sprintf("DOMExceptionCode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/DOMRangeExceptionCode
type DOMRangeExceptionCode uint32

const (
	DOM_BAD_BOUNDARYPOINTS_ERR DOMRangeExceptionCode = 1
	DOM_INVALID_NODE_TYPE_ERR  DOMRangeExceptionCode = 2
)

func (e DOMRangeExceptionCode) String() string {
	switch e {
	case DOM_BAD_BOUNDARYPOINTS_ERR:
		return "DOM_BAD_BOUNDARYPOINTS_ERR"
	case DOM_INVALID_NODE_TYPE_ERR:
		return "DOM_INVALID_NODE_TYPE_ERR"
	default:
		return fmt.Sprintf("DOMRangeExceptionCode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/DOMXPathExceptionCode
type DOMXPathExceptionCode uint32

const (
	DOM_INVALID_EXPRESSION_ERR DOMXPathExceptionCode = 51
	DOM_TYPE_ERR               DOMXPathExceptionCode = 52
)

func (e DOMXPathExceptionCode) String() string {
	switch e {
	case DOM_INVALID_EXPRESSION_ERR:
		return "DOM_INVALID_EXPRESSION_ERR"
	case DOM_TYPE_ERR:
		return "DOM_TYPE_ERR"
	default:
		return fmt.Sprintf("DOMXPathExceptionCode(%d)", e)
	}
}

type Dom uint

const (
	DOM_ADDITION                                  Dom = 2
	DOM_ANY_TYPE                                  Dom = 0
	DOM_ANY_UNORDERED_NODE_TYPE                   Dom = 8
	DOM_ATTRIBUTE_NODE                            Dom = 2
	DOM_AT_TARGET                                 Dom = 2
	DOM_BOOLEAN_TYPE                              Dom = 3
	DOM_BOTH                                      Dom = 2
	DOM_BUBBLING_PHASE                            Dom = 3
	DOM_CAPTURING_PHASE                           Dom = 1
	DOM_CDATA_SECTION_NODE                        Dom = 4
	DOM_CHARSET_RULE                              Dom = 2
	DOM_COMMENT_NODE                              Dom = 8
	DOM_DOCUMENT_FRAGMENT_NODE                    Dom = 11
	DOM_DOCUMENT_NODE                             Dom = 9
	DOM_DOCUMENT_POSITION_CONTAINED_BY            Dom = 16
	DOM_DOCUMENT_POSITION_CONTAINS                Dom = 8
	DOM_DOCUMENT_POSITION_DISCONNECTED            Dom = 1
	DOM_DOCUMENT_POSITION_FOLLOWING               Dom = 4
	DOM_DOCUMENT_POSITION_IMPLEMENTATION_SPECIFIC Dom = 32
	DOM_DOCUMENT_POSITION_PRECEDING               Dom = 2
	DOM_DOCUMENT_TYPE_NODE                        Dom = 10
	DOM_ELEMENT_NODE                              Dom = 1
	DOM_END_TO_END                                Dom = 2
	DOM_END_TO_START                              Dom = 3
	DOM_ENTITY_NODE                               Dom = 6
	DOM_ENTITY_REFERENCE_NODE                     Dom = 5
	DOM_FILTER_ACCEPT                             Dom = 1
	DOM_FILTER_REJECT                             Dom = 2
	DOM_FILTER_SKIP                               Dom = 3
	DOM_FIRST_ORDERED_NODE_TYPE                   Dom = 9
	DOM_FONT_FACE_RULE                            Dom = 5
	DOM_HORIZONTAL                                Dom = 0
	DOM_IMPORT_RULE                               Dom = 3
	DOM_KEYFRAMES_RULE                            Dom = 7
	DOM_KEYFRAME_RULE                             Dom = 8
	DOM_MEDIA_RULE                                Dom = 4
	DOM_MODIFICATION                              Dom = 1
	DOM_NAMESPACE_RULE                            Dom = 10
	DOM_NODE_AFTER                                Dom = 1
	DOM_NODE_BEFORE                               Dom = 0
	DOM_NODE_BEFORE_AND_AFTER                     Dom = 2
	DOM_NODE_INSIDE                               Dom = 3
	DOM_NONE                                      Dom = 0
	DOM_NOTATION_NODE                             Dom = 12
	DOM_NUMBER_TYPE                               Dom = 1
	DOM_ORDERED_NODE_ITERATOR_TYPE                Dom = 5
	DOM_ORDERED_NODE_SNAPSHOT_TYPE                Dom = 7
	DOM_PAGE_RULE                                 Dom = 6
	DOM_PROCESSING_INSTRUCTION_NODE               Dom = 7
	DOM_REMOVAL                                   Dom = 3
	DOM_SHOW_ALL                                  Dom = 4294967295
	DOM_SHOW_ATTRIBUTE                            Dom = 2
	DOM_SHOW_CDATA_SECTION                        Dom = 8
	DOM_SHOW_COMMENT                              Dom = 128
	DOM_SHOW_DOCUMENT                             Dom = 256
	DOM_SHOW_DOCUMENT_FRAGMENT                    Dom = 1024
	DOM_SHOW_DOCUMENT_TYPE                        Dom = 512
	DOM_SHOW_ELEMENT                              Dom = 1
	DOM_SHOW_ENTITY                               Dom = 32
	DOM_SHOW_ENTITY_REFERENCE                     Dom = 16
	DOM_SHOW_NOTATION                             Dom = 2048
	DOM_SHOW_PROCESSING_INSTRUCTION               Dom = 64
	DOM_SHOW_TEXT                                 Dom = 4
	DOM_START_TO_END                              Dom = 1
	DOM_START_TO_START                            Dom = 0
	DOM_STRING_TYPE                               Dom = 2
	DOM_STYLE_RULE                                Dom = 1
	DOM_SUPPORTS_RULE                             Dom = 12
	DOM_TEXT_NODE                                 Dom = 3
	DOM_UNKNOWN_RULE                              Dom = 0
	DOM_UNORDERED_NODE_ITERATOR_TYPE              Dom = 4
	DOM_UNORDERED_NODE_SNAPSHOT_TYPE              Dom = 6
	DOM_VERTICAL                                  Dom = 1
	DOM_WEBKIT_KEYFRAMES_RULE                     Dom = 7
	DOM_WEBKIT_KEYFRAME_RULE                      Dom = 8
	DOM_WEBKIT_REGION_RULE                        Dom = 16
)

func (e Dom) String() string {
	switch e {
	case DOM_ADDITION:
		return "DOM_ADDITION"
	case DOM_ANY_TYPE:
		return "DOM_ANY_TYPE"
	case DOM_ANY_UNORDERED_NODE_TYPE:
		return "DOM_ANY_UNORDERED_NODE_TYPE"
	case DOM_BOOLEAN_TYPE:
		return "DOM_BOOLEAN_TYPE"
	case DOM_CAPTURING_PHASE:
		return "DOM_CAPTURING_PHASE"
	case DOM_CDATA_SECTION_NODE:
		return "DOM_CDATA_SECTION_NODE"
	case DOM_DOCUMENT_FRAGMENT_NODE:
		return "DOM_DOCUMENT_FRAGMENT_NODE"
	case DOM_DOCUMENT_NODE:
		return "DOM_DOCUMENT_NODE"
	case DOM_DOCUMENT_POSITION_CONTAINED_BY:
		return "DOM_DOCUMENT_POSITION_CONTAINED_BY"
	case DOM_DOCUMENT_POSITION_IMPLEMENTATION_SPECIFIC:
		return "DOM_DOCUMENT_POSITION_IMPLEMENTATION_SPECIFIC"
	case DOM_DOCUMENT_TYPE_NODE:
		return "DOM_DOCUMENT_TYPE_NODE"
	case DOM_ENTITY_NODE:
		return "DOM_ENTITY_NODE"
	case DOM_ENTITY_REFERENCE_NODE:
		return "DOM_ENTITY_REFERENCE_NODE"
	case DOM_KEYFRAMES_RULE:
		return "DOM_KEYFRAMES_RULE"
	case DOM_NOTATION_NODE:
		return "DOM_NOTATION_NODE"
	case DOM_SHOW_ALL:
		return "DOM_SHOW_ALL"
	case DOM_SHOW_COMMENT:
		return "DOM_SHOW_COMMENT"
	case DOM_SHOW_DOCUMENT:
		return "DOM_SHOW_DOCUMENT"
	case DOM_SHOW_DOCUMENT_FRAGMENT:
		return "DOM_SHOW_DOCUMENT_FRAGMENT"
	case DOM_SHOW_DOCUMENT_TYPE:
		return "DOM_SHOW_DOCUMENT_TYPE"
	case DOM_SHOW_NOTATION:
		return "DOM_SHOW_NOTATION"
	case DOM_SHOW_PROCESSING_INSTRUCTION:
		return "DOM_SHOW_PROCESSING_INSTRUCTION"
	default:
		return fmt.Sprintf("Dom(%d)", e)
	}
}

type DomAllowKeyboard uint

const (
	DOM_ALLOW_KEYBOARD_INPUT DomAllowKeyboard = 1
)

func (e DomAllowKeyboard) String() string {
	switch e {
	case DOM_ALLOW_KEYBOARD_INPUT:
		return "DOM_ALLOW_KEYBOARD_INPUT"
	default:
		return fmt.Sprintf("DomAllowKeyboard(%d)", e)
	}
}

type DomCSS uint

const (
	DOM_CSS_ATTR            DomCSS = 22
	DOM_CSS_CM              DomCSS = 6
	DOM_CSS_COUNTER         DomCSS = 23
	DOM_CSS_CUSTOM          DomCSS = 3
	DOM_CSS_DEG             DomCSS = 11
	DOM_CSS_DIMENSION       DomCSS = 18
	DOM_CSS_EMS             DomCSS = 3
	DOM_CSS_EXS             DomCSS = 4
	DOM_CSS_GRAD            DomCSS = 13
	DOM_CSS_HZ              DomCSS = 16
	DOM_CSS_IDENT           DomCSS = 21
	DOM_CSS_IN              DomCSS = 8
	DOM_CSS_INHERIT         DomCSS = 0
	DOM_CSS_KHZ             DomCSS = 17
	DOM_CSS_MM              DomCSS = 7
	DOM_CSS_MS              DomCSS = 14
	DOM_CSS_NUMBER          DomCSS = 1
	DOM_CSS_PC              DomCSS = 10
	DOM_CSS_PERCENTAGE      DomCSS = 2
	DOM_CSS_PRIMITIVE_VALUE DomCSS = 1
	DOM_CSS_PT              DomCSS = 9
	DOM_CSS_PX              DomCSS = 5
	DOM_CSS_RAD             DomCSS = 12
	DOM_CSS_RECT            DomCSS = 24
	DOM_CSS_RGBCOLOR        DomCSS = 25
	DOM_CSS_S               DomCSS = 15
	DOM_CSS_STRING          DomCSS = 19
	DOM_CSS_UNKNOWN         DomCSS = 0
	DOM_CSS_URI             DomCSS = 20
	DOM_CSS_VALUE_LIST      DomCSS = 2
	DOM_CSS_VH              DomCSS = 27
	DOM_CSS_VMAX            DomCSS = 29
	DOM_CSS_VMIN            DomCSS = 28
	DOM_CSS_VW              DomCSS = 26
)

func (e DomCSS) String() string {
	switch e {
	case DOM_CSS_ATTR:
		return "DOM_CSS_ATTR"
	case DOM_CSS_CM:
		return "DOM_CSS_CM"
	case DOM_CSS_COUNTER:
		return "DOM_CSS_COUNTER"
	case DOM_CSS_CUSTOM:
		return "DOM_CSS_CUSTOM"
	case DOM_CSS_DEG:
		return "DOM_CSS_DEG"
	case DOM_CSS_DIMENSION:
		return "DOM_CSS_DIMENSION"
	case DOM_CSS_EXS:
		return "DOM_CSS_EXS"
	case DOM_CSS_GRAD:
		return "DOM_CSS_GRAD"
	case DOM_CSS_HZ:
		return "DOM_CSS_HZ"
	case DOM_CSS_IDENT:
		return "DOM_CSS_IDENT"
	case DOM_CSS_IN:
		return "DOM_CSS_IN"
	case DOM_CSS_INHERIT:
		return "DOM_CSS_INHERIT"
	case DOM_CSS_KHZ:
		return "DOM_CSS_KHZ"
	case DOM_CSS_MM:
		return "DOM_CSS_MM"
	case DOM_CSS_MS:
		return "DOM_CSS_MS"
	case DOM_CSS_NUMBER:
		return "DOM_CSS_NUMBER"
	case DOM_CSS_PC:
		return "DOM_CSS_PC"
	case DOM_CSS_PERCENTAGE:
		return "DOM_CSS_PERCENTAGE"
	case DOM_CSS_PT:
		return "DOM_CSS_PT"
	case DOM_CSS_PX:
		return "DOM_CSS_PX"
	case DOM_CSS_RAD:
		return "DOM_CSS_RAD"
	case DOM_CSS_RECT:
		return "DOM_CSS_RECT"
	case DOM_CSS_RGBCOLOR:
		return "DOM_CSS_RGBCOLOR"
	case DOM_CSS_S:
		return "DOM_CSS_S"
	case DOM_CSS_STRING:
		return "DOM_CSS_STRING"
	case DOM_CSS_URI:
		return "DOM_CSS_URI"
	case DOM_CSS_VH:
		return "DOM_CSS_VH"
	case DOM_CSS_VMAX:
		return "DOM_CSS_VMAX"
	case DOM_CSS_VMIN:
		return "DOM_CSS_VMIN"
	case DOM_CSS_VW:
		return "DOM_CSS_VW"
	default:
		return fmt.Sprintf("DomCSS(%d)", e)
	}
}

type DomDomDelta uint

const (
	DOM_DOM_DELTA_LINE  DomDomDelta = 1
	DOM_DOM_DELTA_PAGE  DomDomDelta = 2
	DOM_DOM_DELTA_PIXEL DomDomDelta = 0
)

func (e DomDomDelta) String() string {
	switch e {
	case DOM_DOM_DELTA_LINE:
		return "DOM_DOM_DELTA_LINE"
	case DOM_DOM_DELTA_PAGE:
		return "DOM_DOM_DELTA_PAGE"
	case DOM_DOM_DELTA_PIXEL:
		return "DOM_DOM_DELTA_PIXEL"
	default:
		return fmt.Sprintf("DomDomDelta(%d)", e)
	}
}

type DomKeyLocation uint

const (
	DOM_KEY_LOCATION_LEFT     DomKeyLocation = 1
	DOM_KEY_LOCATION_NUMPAD   DomKeyLocation = 3
	DOM_KEY_LOCATION_RIGHT    DomKeyLocation = 2
	DOM_KEY_LOCATION_STANDARD DomKeyLocation = 0
)

func (e DomKeyLocation) String() string {
	switch e {
	case DOM_KEY_LOCATION_LEFT:
		return "DOM_KEY_LOCATION_LEFT"
	case DOM_KEY_LOCATION_NUMPAD:
		return "DOM_KEY_LOCATION_NUMPAD"
	case DOM_KEY_LOCATION_RIGHT:
		return "DOM_KEY_LOCATION_RIGHT"
	case DOM_KEY_LOCATION_STANDARD:
		return "DOM_KEY_LOCATION_STANDARD"
	default:
		return fmt.Sprintf("DomKeyLocation(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/WKAudiovisualMediaTypes
type WKAudiovisualMediaTypes uint

const (
	// WKAudiovisualMediaTypeAll: All media types require a user gesture to begin playing.
	WKAudiovisualMediaTypeAll WKAudiovisualMediaTypes = 0
	// WKAudiovisualMediaTypeAudio: Media types that contain audio require a user gesture to begin playing.
	WKAudiovisualMediaTypeAudio WKAudiovisualMediaTypes = 1
	// WKAudiovisualMediaTypeNone: No media types require a user gesture to begin playing.
	WKAudiovisualMediaTypeNone WKAudiovisualMediaTypes = 0
	// WKAudiovisualMediaTypeVideo: Media types that contain video require a user gesture to begin playing.
	WKAudiovisualMediaTypeVideo WKAudiovisualMediaTypes = 2
)

func (e WKAudiovisualMediaTypes) String() string {
	switch e {
	case WKAudiovisualMediaTypeAll:
		return "WKAudiovisualMediaTypeAll"
	case WKAudiovisualMediaTypeAudio:
		return "WKAudiovisualMediaTypeAudio"
	case WKAudiovisualMediaTypeVideo:
		return "WKAudiovisualMediaTypeVideo"
	default:
		return fmt.Sprintf("WKAudiovisualMediaTypes(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/WKWebpagePreferences/ContentMode
type WKContentMode int

const (
	// WKContentModeDesktop: The content mode that represents a desktop experience.
	WKContentModeDesktop WKContentMode = 2
	// WKContentModeMobile: The content mode that represents a mobile experience.
	WKContentModeMobile WKContentMode = 1
	// WKContentModeRecommended: The content mode that is appropriate for the current device.
	WKContentModeRecommended WKContentMode = 0
)

func (e WKContentMode) String() string {
	switch e {
	case WKContentModeDesktop:
		return "WKContentModeDesktop"
	case WKContentModeMobile:
		return "WKContentModeMobile"
	case WKContentModeRecommended:
		return "WKContentModeRecommended"
	default:
		return fmt.Sprintf("WKContentMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/WKHTTPCookieStore/CookiePolicy
type WKCookiePolicy int

const (
	// WKCookiePolicyAllow: A case that indicates the cookie store allows cookie storage.
	WKCookiePolicyAllow WKCookiePolicy = 0
	// WKCookiePolicyDisallow: A case that indicates the cookie store does not allow cookie storage.
	WKCookiePolicyDisallow WKCookiePolicy = 1
)

func (e WKCookiePolicy) String() string {
	switch e {
	case WKCookiePolicyAllow:
		return "WKCookiePolicyAllow"
	case WKCookiePolicyDisallow:
		return "WKCookiePolicyDisallow"
	default:
		return fmt.Sprintf("WKCookiePolicy(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/WKDataDetectorTypes
type WKDataDetectorTypes uint

const (
	// WKDataDetectorTypeAddress: Detect addresses in text and turn them into links to display the location.
	WKDataDetectorTypeAddress WKDataDetectorTypes = 0
	// WKDataDetectorTypeAll: Detect all data types and turn them into links.
	WKDataDetectorTypeAll WKDataDetectorTypes = 0
	// WKDataDetectorTypeCalendarEvent: Turn future dates and times into links to create calendar events.
	WKDataDetectorTypeCalendarEvent WKDataDetectorTypes = 0
	// WKDataDetectorTypeFlightNumber: Detect flight numbers in text and turn them into links.
	WKDataDetectorTypeFlightNumber WKDataDetectorTypes = 0
	// WKDataDetectorTypeLink: Detect URLs in text and turn them into links.
	WKDataDetectorTypeLink WKDataDetectorTypes = 0
	// WKDataDetectorTypeLookupSuggestion: Detect Spotlight suggestions and turn them into links.
	WKDataDetectorTypeLookupSuggestion WKDataDetectorTypes = 0
	// WKDataDetectorTypeNone: No data detection.
	WKDataDetectorTypeNone WKDataDetectorTypes = 0
	// WKDataDetectorTypePhoneNumber: Detect phone numbers in text and create a link to call the specified number.
	WKDataDetectorTypePhoneNumber WKDataDetectorTypes = 0
	// WKDataDetectorTypeTrackingNumber: Detect tracking numbers in text and turn them into links.
	WKDataDetectorTypeTrackingNumber WKDataDetectorTypes = 0
	// Deprecated.
	WKDataDetectorTypeSpotlightSuggestion WKDataDetectorTypes = 0
)

func (e WKDataDetectorTypes) String() string {
	switch e {
	case WKDataDetectorTypeAddress:
		return "WKDataDetectorTypeAddress"
	default:
		return fmt.Sprintf("WKDataDetectorTypes(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/WKDialogResult
type WKDialogResult int

const (
	// WKDialogResultAskAgain: A result that indicates the delegate didn’t display a message, so other web views should check again.
	WKDialogResultAskAgain WKDialogResult = 2
	// WKDialogResultHandled: A result that indicates the delegate displayed the first use message.
	WKDialogResultHandled WKDialogResult = 3
	// WKDialogResultShowDefault: A result that indicates the delegate didn’t display a message, so the web view should show the default Lockdown Mode message.
	WKDialogResultShowDefault WKDialogResult = 1
)

func (e WKDialogResult) String() string {
	switch e {
	case WKDialogResultAskAgain:
		return "WKDialogResultAskAgain"
	case WKDialogResultHandled:
		return "WKDialogResultHandled"
	case WKDialogResultShowDefault:
		return "WKDialogResultShowDefault"
	default:
		return fmt.Sprintf("WKDialogResult(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/WKDownload/PlaceholderPolicy
type WKDownloadPlaceholderPolicy int

const (
	WKDownloadPlaceholderPolicyDisable WKDownloadPlaceholderPolicy = 0
	WKDownloadPlaceholderPolicyEnable  WKDownloadPlaceholderPolicy = 1
)

func (e WKDownloadPlaceholderPolicy) String() string {
	switch e {
	case WKDownloadPlaceholderPolicyDisable:
		return "WKDownloadPlaceholderPolicyDisable"
	case WKDownloadPlaceholderPolicyEnable:
		return "WKDownloadPlaceholderPolicyEnable"
	default:
		return fmt.Sprintf("WKDownloadPlaceholderPolicy(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/WKDownload/RedirectPolicy
type WKDownloadRedirectPolicy int

const (
	// WKDownloadRedirectPolicyAllow: Allow a redirect to proceed.
	WKDownloadRedirectPolicyAllow WKDownloadRedirectPolicy = 1
	// WKDownloadRedirectPolicyCancel: Cancel the redirect action.
	WKDownloadRedirectPolicyCancel WKDownloadRedirectPolicy = 0
)

func (e WKDownloadRedirectPolicy) String() string {
	switch e {
	case WKDownloadRedirectPolicyAllow:
		return "WKDownloadRedirectPolicyAllow"
	case WKDownloadRedirectPolicyCancel:
		return "WKDownloadRedirectPolicyCancel"
	default:
		return fmt.Sprintf("WKDownloadRedirectPolicy(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/WKError/Code
type WKErrorCode int

const (
	// WKErrorAttributedStringContentFailedToLoad: An error that indicates the failure to navigate to web content from an attributed string.
	WKErrorAttributedStringContentFailedToLoad WKErrorCode = 10
	// WKErrorAttributedStringContentLoadTimedOut: An error that indicates a timeout occurred while trying to load web content from an attributed string.
	WKErrorAttributedStringContentLoadTimedOut WKErrorCode = 11
	// WKErrorContentRuleListStoreCompileFailed: An error that indicates the compilation of a rule list failed.
	WKErrorContentRuleListStoreCompileFailed WKErrorCode = 6
	// WKErrorContentRuleListStoreLookUpFailed: An error that indicates a content rule list data store didn’t find a rule list with the specified identifier.
	WKErrorContentRuleListStoreLookUpFailed WKErrorCode = 7
	// WKErrorContentRuleListStoreRemoveFailed: An error that indicates a failure to remove a content rule list from the rule list data store object.
	WKErrorContentRuleListStoreRemoveFailed WKErrorCode = 8
	// WKErrorContentRuleListStoreVersionMismatch: An error that indicates the rule list version is outdated and cannot be read.
	WKErrorContentRuleListStoreVersionMismatch WKErrorCode = 9
	// WKErrorCredentialNotFound: An error that indicates the system could not find a passkey during an export.
	WKErrorCredentialNotFound WKErrorCode = 17
	// WKErrorDuplicateCredential: An error that indicates the system found a duplicate passkey during an import.
	WKErrorDuplicateCredential WKErrorCode = 15
	// WKErrorJavaScriptAppBoundDomain: An error that indicates JavaScript execution failed due to an app-bound domain restriction.
	WKErrorJavaScriptAppBoundDomain WKErrorCode = 14
	// WKErrorJavaScriptExceptionOccurred: An error that indicates a JavaScript exception occurred.
	WKErrorJavaScriptExceptionOccurred WKErrorCode = 4
	// WKErrorJavaScriptInvalidFrameTarget: An error that indicates your content referenced an invalid web frame.
	WKErrorJavaScriptInvalidFrameTarget WKErrorCode = 12
	// WKErrorJavaScriptResultTypeIsUnsupported: An error that indicates the result of JavaScript execution could not be returned.
	WKErrorJavaScriptResultTypeIsUnsupported WKErrorCode = 5
	// WKErrorMalformedCredential: An error that indicates the system could not parse passkey data during an import.
	WKErrorMalformedCredential WKErrorCode = 16
	// WKErrorNavigationAppBoundDomain: An error that indicates navigation failed due to an app-bound domain restriction.
	WKErrorNavigationAppBoundDomain WKErrorCode = 13
	// WKErrorUnknown: An error that indicates an unknown issue occurred.
	WKErrorUnknown WKErrorCode = 1
	// WKErrorWebContentProcessTerminated: An error that indicates the web process that contains the content is no longer running.
	WKErrorWebContentProcessTerminated WKErrorCode = 2
	// WKErrorWebViewInvalidated: An error that indicates the web view was invalidated.
	WKErrorWebViewInvalidated WKErrorCode = 3
)

func (e WKErrorCode) String() string {
	switch e {
	case WKErrorAttributedStringContentFailedToLoad:
		return "WKErrorAttributedStringContentFailedToLoad"
	case WKErrorAttributedStringContentLoadTimedOut:
		return "WKErrorAttributedStringContentLoadTimedOut"
	case WKErrorContentRuleListStoreCompileFailed:
		return "WKErrorContentRuleListStoreCompileFailed"
	case WKErrorContentRuleListStoreLookUpFailed:
		return "WKErrorContentRuleListStoreLookUpFailed"
	case WKErrorContentRuleListStoreRemoveFailed:
		return "WKErrorContentRuleListStoreRemoveFailed"
	case WKErrorContentRuleListStoreVersionMismatch:
		return "WKErrorContentRuleListStoreVersionMismatch"
	case WKErrorCredentialNotFound:
		return "WKErrorCredentialNotFound"
	case WKErrorDuplicateCredential:
		return "WKErrorDuplicateCredential"
	case WKErrorJavaScriptAppBoundDomain:
		return "WKErrorJavaScriptAppBoundDomain"
	case WKErrorJavaScriptExceptionOccurred:
		return "WKErrorJavaScriptExceptionOccurred"
	case WKErrorJavaScriptInvalidFrameTarget:
		return "WKErrorJavaScriptInvalidFrameTarget"
	case WKErrorJavaScriptResultTypeIsUnsupported:
		return "WKErrorJavaScriptResultTypeIsUnsupported"
	case WKErrorMalformedCredential:
		return "WKErrorMalformedCredential"
	case WKErrorNavigationAppBoundDomain:
		return "WKErrorNavigationAppBoundDomain"
	case WKErrorUnknown:
		return "WKErrorUnknown"
	case WKErrorWebContentProcessTerminated:
		return "WKErrorWebContentProcessTerminated"
	case WKErrorWebViewInvalidated:
		return "WKErrorWebViewInvalidated"
	default:
		return fmt.Sprintf("WKErrorCode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/WKWebView/FullscreenState-swift.enum
type WKFullscreenState int

const (
	WKFullscreenStateEnteringFullscreen WKFullscreenState = 1
	WKFullscreenStateExitingFullscreen  WKFullscreenState = 3
	WKFullscreenStateInFullscreen       WKFullscreenState = 2
	WKFullscreenStateNotInFullscreen    WKFullscreenState = 0
)

func (e WKFullscreenState) String() string {
	switch e {
	case WKFullscreenStateEnteringFullscreen:
		return "WKFullscreenStateEnteringFullscreen"
	case WKFullscreenStateExitingFullscreen:
		return "WKFullscreenStateExitingFullscreen"
	case WKFullscreenStateInFullscreen:
		return "WKFullscreenStateInFullscreen"
	case WKFullscreenStateNotInFullscreen:
		return "WKFullscreenStateNotInFullscreen"
	default:
		return fmt.Sprintf("WKFullscreenState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/WKPreferences/InactiveSchedulingPolicy-swift.enum
type WKInactiveSchedulingPolicy int

const (
	// WKInactiveSchedulingPolicyNone: A policy where a web view that’s not in a window runs tasks normally.
	WKInactiveSchedulingPolicyNone WKInactiveSchedulingPolicy = 2
	// WKInactiveSchedulingPolicySuspend: A policy where a web view that’s not in a window fully suspends tasks.
	WKInactiveSchedulingPolicySuspend WKInactiveSchedulingPolicy = 0
	// WKInactiveSchedulingPolicyThrottle: A policy where a web view that’s not in a window limits processing, but does not fully suspend tasks.
	WKInactiveSchedulingPolicyThrottle WKInactiveSchedulingPolicy = 1
)

func (e WKInactiveSchedulingPolicy) String() string {
	switch e {
	case WKInactiveSchedulingPolicyNone:
		return "WKInactiveSchedulingPolicyNone"
	case WKInactiveSchedulingPolicySuspend:
		return "WKInactiveSchedulingPolicySuspend"
	case WKInactiveSchedulingPolicyThrottle:
		return "WKInactiveSchedulingPolicyThrottle"
	default:
		return fmt.Sprintf("WKInactiveSchedulingPolicy(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/WKMediaCaptureState
type WKMediaCaptureState int

const (
	// WKMediaCaptureStateActive: The media device is actively capturing audio or video.
	WKMediaCaptureStateActive WKMediaCaptureState = 1
	// WKMediaCaptureStateMuted: The media device is muted, and not actively capturing audio or video.
	WKMediaCaptureStateMuted WKMediaCaptureState = 2
	// WKMediaCaptureStateNone: The media device is off.
	WKMediaCaptureStateNone WKMediaCaptureState = 0
)

func (e WKMediaCaptureState) String() string {
	switch e {
	case WKMediaCaptureStateActive:
		return "WKMediaCaptureStateActive"
	case WKMediaCaptureStateMuted:
		return "WKMediaCaptureStateMuted"
	case WKMediaCaptureStateNone:
		return "WKMediaCaptureStateNone"
	default:
		return fmt.Sprintf("WKMediaCaptureState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/WKMediaCaptureType
type WKMediaCaptureType int

const (
	// WKMediaCaptureTypeCamera: A media device that can capture video.
	WKMediaCaptureTypeCamera WKMediaCaptureType = 0
	// WKMediaCaptureTypeCameraAndMicrophone: A media device or devices that can capture audio and video.
	WKMediaCaptureTypeCameraAndMicrophone WKMediaCaptureType = 2
	// WKMediaCaptureTypeMicrophone: A media device that can capture audio.
	WKMediaCaptureTypeMicrophone WKMediaCaptureType = 1
)

func (e WKMediaCaptureType) String() string {
	switch e {
	case WKMediaCaptureTypeCamera:
		return "WKMediaCaptureTypeCamera"
	case WKMediaCaptureTypeCameraAndMicrophone:
		return "WKMediaCaptureTypeCameraAndMicrophone"
	case WKMediaCaptureTypeMicrophone:
		return "WKMediaCaptureTypeMicrophone"
	default:
		return fmt.Sprintf("WKMediaCaptureType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/WKMediaPlaybackState
type WKMediaPlaybackState int

const (
	// WKMediaPlaybackStateNone: There is no media to play back.
	WKMediaPlaybackStateNone WKMediaPlaybackState = 0
	// WKMediaPlaybackStatePaused: The media playback is paused.
	WKMediaPlaybackStatePaused WKMediaPlaybackState = 2
	// WKMediaPlaybackStatePlaying: The media is playing.
	WKMediaPlaybackStatePlaying WKMediaPlaybackState = 1
	// WKMediaPlaybackStateSuspended: The media is not playing, and cannot be resumed until the user revokes the suspension.
	WKMediaPlaybackStateSuspended WKMediaPlaybackState = 3
)

func (e WKMediaPlaybackState) String() string {
	switch e {
	case WKMediaPlaybackStateNone:
		return "WKMediaPlaybackStateNone"
	case WKMediaPlaybackStatePaused:
		return "WKMediaPlaybackStatePaused"
	case WKMediaPlaybackStatePlaying:
		return "WKMediaPlaybackStatePlaying"
	case WKMediaPlaybackStateSuspended:
		return "WKMediaPlaybackStateSuspended"
	default:
		return fmt.Sprintf("WKMediaPlaybackState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/WKNavigationActionPolicy
type WKNavigationActionPolicy int

const (
	// WKNavigationActionPolicyAllow: Allow the navigation to continue.
	WKNavigationActionPolicyAllow WKNavigationActionPolicy = 1
	// WKNavigationActionPolicyCancel: Cancel the navigation.
	WKNavigationActionPolicyCancel WKNavigationActionPolicy = 0
	// WKNavigationActionPolicyDownload: Allow the download to proceed.
	WKNavigationActionPolicyDownload WKNavigationActionPolicy = 2
)

func (e WKNavigationActionPolicy) String() string {
	switch e {
	case WKNavigationActionPolicyAllow:
		return "WKNavigationActionPolicyAllow"
	case WKNavigationActionPolicyCancel:
		return "WKNavigationActionPolicyCancel"
	case WKNavigationActionPolicyDownload:
		return "WKNavigationActionPolicyDownload"
	default:
		return fmt.Sprintf("WKNavigationActionPolicy(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/WKNavigationResponsePolicy
type WKNavigationResponsePolicy int

const (
	// WKNavigationResponsePolicyAllow: Allow the navigation to continue.
	WKNavigationResponsePolicyAllow WKNavigationResponsePolicy = 1
	// WKNavigationResponsePolicyCancel: Cancel the navigation.
	WKNavigationResponsePolicyCancel WKNavigationResponsePolicy = 0
	// WKNavigationResponsePolicyDownload: Allow the download to proceed.
	WKNavigationResponsePolicyDownload WKNavigationResponsePolicy = 2
)

func (e WKNavigationResponsePolicy) String() string {
	switch e {
	case WKNavigationResponsePolicyAllow:
		return "WKNavigationResponsePolicyAllow"
	case WKNavigationResponsePolicyCancel:
		return "WKNavigationResponsePolicyCancel"
	case WKNavigationResponsePolicyDownload:
		return "WKNavigationResponsePolicyDownload"
	default:
		return fmt.Sprintf("WKNavigationResponsePolicy(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/WKNavigationType
type WKNavigationType int

const (
	// WKNavigationTypeBackForward: A request for the frame’s next or previous item.
	WKNavigationTypeBackForward WKNavigationType = 2
	// WKNavigationTypeFormResubmitted: A request to resubmit a form.
	WKNavigationTypeFormResubmitted WKNavigationType = 4
	// WKNavigationTypeFormSubmitted: A request to submit a form.
	WKNavigationTypeFormSubmitted WKNavigationType = 1
	// WKNavigationTypeLinkActivated: A link activation.
	WKNavigationTypeLinkActivated WKNavigationType = 0
	// WKNavigationTypeOther: A navigation request that originates for some other reason.
	WKNavigationTypeOther WKNavigationType = -1
	// WKNavigationTypeReload: A request to reload the webpage.
	WKNavigationTypeReload WKNavigationType = 3
)

func (e WKNavigationType) String() string {
	switch e {
	case WKNavigationTypeBackForward:
		return "WKNavigationTypeBackForward"
	case WKNavigationTypeFormResubmitted:
		return "WKNavigationTypeFormResubmitted"
	case WKNavigationTypeFormSubmitted:
		return "WKNavigationTypeFormSubmitted"
	case WKNavigationTypeLinkActivated:
		return "WKNavigationTypeLinkActivated"
	case WKNavigationTypeOther:
		return "WKNavigationTypeOther"
	case WKNavigationTypeReload:
		return "WKNavigationTypeReload"
	default:
		return fmt.Sprintf("WKNavigationType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/WKPermissionDecision
type WKPermissionDecision int

const (
	// WKPermissionDecisionDeny: Deny permission for the requested resource.
	WKPermissionDecisionDeny WKPermissionDecision = 2
	// WKPermissionDecisionGrant: Grant permission for the requested resource.
	WKPermissionDecisionGrant WKPermissionDecision = 1
	// WKPermissionDecisionPrompt: Prompt the user for permission for the requested resource.
	WKPermissionDecisionPrompt WKPermissionDecision = 0
)

func (e WKPermissionDecision) String() string {
	switch e {
	case WKPermissionDecisionDeny:
		return "WKPermissionDecisionDeny"
	case WKPermissionDecisionGrant:
		return "WKPermissionDecisionGrant"
	case WKPermissionDecisionPrompt:
		return "WKPermissionDecisionPrompt"
	default:
		return fmt.Sprintf("WKPermissionDecision(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/WKSecurityRestrictionMode
type WKSecurityRestrictionMode int

const (
	WKSecurityRestrictionModeLockdown              WKSecurityRestrictionMode = 2
	WKSecurityRestrictionModeMaximizeCompatibility WKSecurityRestrictionMode = 1
	WKSecurityRestrictionModeNone                  WKSecurityRestrictionMode = 0
)

func (e WKSecurityRestrictionMode) String() string {
	switch e {
	case WKSecurityRestrictionModeLockdown:
		return "WKSecurityRestrictionModeLockdown"
	case WKSecurityRestrictionModeMaximizeCompatibility:
		return "WKSecurityRestrictionModeMaximizeCompatibility"
	case WKSecurityRestrictionModeNone:
		return "WKSecurityRestrictionModeNone"
	default:
		return fmt.Sprintf("WKSecurityRestrictionMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/WKSelectionGranularity
type WKSelectionGranularity int

const (
	// Deprecated.
	WKSelectionGranularityCharacter WKSelectionGranularity = 0
	// Deprecated.
	WKSelectionGranularityDynamic WKSelectionGranularity = 0
)

func (e WKSelectionGranularity) String() string {
	switch e {
	case WKSelectionGranularityCharacter:
		return "WKSelectionGranularityCharacter"
	default:
		return fmt.Sprintf("WKSelectionGranularity(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/WKUserInterfaceDirectionPolicy
type WKUserInterfaceDirectionPolicy int

const (
	// WKUserInterfaceDirectionPolicyContent: The directionality follows the CSS/HTML/XHTML specifications.
	WKUserInterfaceDirectionPolicyContent WKUserInterfaceDirectionPolicy = 0
	// WKUserInterfaceDirectionPolicySystem: The directionality follows the view’s user interface layout direction.
	WKUserInterfaceDirectionPolicySystem WKUserInterfaceDirectionPolicy = 1
)

func (e WKUserInterfaceDirectionPolicy) String() string {
	switch e {
	case WKUserInterfaceDirectionPolicyContent:
		return "WKUserInterfaceDirectionPolicyContent"
	case WKUserInterfaceDirectionPolicySystem:
		return "WKUserInterfaceDirectionPolicySystem"
	default:
		return fmt.Sprintf("WKUserInterfaceDirectionPolicy(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/WKUserScriptInjectionTime
type WKUserScriptInjectionTime int

const (
	// WKUserScriptInjectionTimeAtDocumentEnd: A constant to inject the script after the document finishes loading, but before loading any other subresources.
	WKUserScriptInjectionTimeAtDocumentEnd WKUserScriptInjectionTime = 1
	// WKUserScriptInjectionTimeAtDocumentStart: A constant to inject the script after the creation of the webpage’s document element, but before loading any other content.
	WKUserScriptInjectionTimeAtDocumentStart WKUserScriptInjectionTime = 0
)

func (e WKUserScriptInjectionTime) String() string {
	switch e {
	case WKUserScriptInjectionTimeAtDocumentEnd:
		return "WKUserScriptInjectionTimeAtDocumentEnd"
	case WKUserScriptInjectionTimeAtDocumentStart:
		return "WKUserScriptInjectionTimeAtDocumentStart"
	default:
		return fmt.Sprintf("WKUserScriptInjectionTime(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/Error/Code
type WKWebExtensionContextError int

const (
	// WKWebExtensionContextErrorAlreadyLoaded: Indicates that the context is already loaded by a WKWebExtensionController.
	WKWebExtensionContextErrorAlreadyLoaded WKWebExtensionContextError = 2
	// WKWebExtensionContextErrorBackgroundContentFailedToLoad: Indicates that an error occurred loading the background content.
	WKWebExtensionContextErrorBackgroundContentFailedToLoad WKWebExtensionContextError = 6
	// WKWebExtensionContextErrorBaseURLAlreadyInUse: Indicates that another context is already using the specified base URL.
	WKWebExtensionContextErrorBaseURLAlreadyInUse WKWebExtensionContextError = 4
	// WKWebExtensionContextErrorNoBackgroundContent: Indicates that the extension does not have background content.
	WKWebExtensionContextErrorNoBackgroundContent WKWebExtensionContextError = 5
	// WKWebExtensionContextErrorNotLoaded: Indicates that the context is not loaded by a WKWebExtensionController.
	WKWebExtensionContextErrorNotLoaded WKWebExtensionContextError = 3
	// WKWebExtensionContextErrorUnknown: Indicates that an unknown error occurred.
	WKWebExtensionContextErrorUnknown WKWebExtensionContextError = 1
)

func (e WKWebExtensionContextError) String() string {
	switch e {
	case WKWebExtensionContextErrorAlreadyLoaded:
		return "WKWebExtensionContextErrorAlreadyLoaded"
	case WKWebExtensionContextErrorBackgroundContentFailedToLoad:
		return "WKWebExtensionContextErrorBackgroundContentFailedToLoad"
	case WKWebExtensionContextErrorBaseURLAlreadyInUse:
		return "WKWebExtensionContextErrorBaseURLAlreadyInUse"
	case WKWebExtensionContextErrorNoBackgroundContent:
		return "WKWebExtensionContextErrorNoBackgroundContent"
	case WKWebExtensionContextErrorNotLoaded:
		return "WKWebExtensionContextErrorNotLoaded"
	case WKWebExtensionContextErrorUnknown:
		return "WKWebExtensionContextErrorUnknown"
	default:
		return fmt.Sprintf("WKWebExtensionContextError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/PermissionStatus
type WKWebExtensionContextPermissionStatus int

const (
	// WKWebExtensionContextPermissionStatusDeniedExplicitly: Indicates that the permission was explicitly denied.
	WKWebExtensionContextPermissionStatusDeniedExplicitly WKWebExtensionContextPermissionStatus = -3
	// WKWebExtensionContextPermissionStatusDeniedImplicitly: Indicates that the permission was implicitly denied because of another explicitly denied permission.
	WKWebExtensionContextPermissionStatusDeniedImplicitly WKWebExtensionContextPermissionStatus = -2
	// WKWebExtensionContextPermissionStatusGrantedExplicitly: Indicates that the permission was explicitly granted permission.
	WKWebExtensionContextPermissionStatusGrantedExplicitly WKWebExtensionContextPermissionStatus = 3
	// WKWebExtensionContextPermissionStatusGrantedImplicitly: Indicates that the permission was implicitly granted because of another explicitly granted permission.
	WKWebExtensionContextPermissionStatusGrantedImplicitly WKWebExtensionContextPermissionStatus = 2
	// WKWebExtensionContextPermissionStatusRequestedExplicitly: Indicates that the permission was explicitly requested.
	WKWebExtensionContextPermissionStatusRequestedExplicitly WKWebExtensionContextPermissionStatus = 1
	// WKWebExtensionContextPermissionStatusRequestedImplicitly: Indicates that the permission was implicitly requested because of another explicitly requested permission.
	WKWebExtensionContextPermissionStatusRequestedImplicitly WKWebExtensionContextPermissionStatus = -1
	// WKWebExtensionContextPermissionStatusUnknown: Indicates an unknown permission status.
	WKWebExtensionContextPermissionStatusUnknown WKWebExtensionContextPermissionStatus = 0
)

func (e WKWebExtensionContextPermissionStatus) String() string {
	switch e {
	case WKWebExtensionContextPermissionStatusDeniedExplicitly:
		return "WKWebExtensionContextPermissionStatusDeniedExplicitly"
	case WKWebExtensionContextPermissionStatusDeniedImplicitly:
		return "WKWebExtensionContextPermissionStatusDeniedImplicitly"
	case WKWebExtensionContextPermissionStatusGrantedExplicitly:
		return "WKWebExtensionContextPermissionStatusGrantedExplicitly"
	case WKWebExtensionContextPermissionStatusGrantedImplicitly:
		return "WKWebExtensionContextPermissionStatusGrantedImplicitly"
	case WKWebExtensionContextPermissionStatusRequestedExplicitly:
		return "WKWebExtensionContextPermissionStatusRequestedExplicitly"
	case WKWebExtensionContextPermissionStatusRequestedImplicitly:
		return "WKWebExtensionContextPermissionStatusRequestedImplicitly"
	case WKWebExtensionContextPermissionStatusUnknown:
		return "WKWebExtensionContextPermissionStatusUnknown"
	default:
		return fmt.Sprintf("WKWebExtensionContextPermissionStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/DataRecord/Error/Code
type WKWebExtensionDataRecordError int

const (
	// WKWebExtensionDataRecordErrorLocalStorageFailed: Indicates a failure occurred when either deleting or calculating local storage.
	WKWebExtensionDataRecordErrorLocalStorageFailed WKWebExtensionDataRecordError = 2
	// WKWebExtensionDataRecordErrorSessionStorageFailed: Indicates a failure occurred when either deleting or calculating session storage.
	WKWebExtensionDataRecordErrorSessionStorageFailed WKWebExtensionDataRecordError = 3
	// WKWebExtensionDataRecordErrorSynchronizedStorageFailed: Indicates a failure occurred when either deleting or calculating synchronized storage.
	WKWebExtensionDataRecordErrorSynchronizedStorageFailed WKWebExtensionDataRecordError = 4
	// WKWebExtensionDataRecordErrorUnknown: Indicates that an unknown error occurred.
	WKWebExtensionDataRecordErrorUnknown WKWebExtensionDataRecordError = 1
)

func (e WKWebExtensionDataRecordError) String() string {
	switch e {
	case WKWebExtensionDataRecordErrorLocalStorageFailed:
		return "WKWebExtensionDataRecordErrorLocalStorageFailed"
	case WKWebExtensionDataRecordErrorSessionStorageFailed:
		return "WKWebExtensionDataRecordErrorSessionStorageFailed"
	case WKWebExtensionDataRecordErrorSynchronizedStorageFailed:
		return "WKWebExtensionDataRecordErrorSynchronizedStorageFailed"
	case WKWebExtensionDataRecordErrorUnknown:
		return "WKWebExtensionDataRecordErrorUnknown"
	default:
		return fmt.Sprintf("WKWebExtensionDataRecordError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/Error/Code
type WKWebExtensionError int

const (
	// WKWebExtensionErrorInvalidArchive: Indicates that the archive file is invalid or corrupt.
	WKWebExtensionErrorInvalidArchive WKWebExtensionError = 9
	// WKWebExtensionErrorInvalidBackgroundPersistence: Indicates that the extension specified background persistence that was not compatible with the platform or features requested.
	WKWebExtensionErrorInvalidBackgroundPersistence WKWebExtensionError = 8
	// WKWebExtensionErrorInvalidDeclarativeNetRequestEntry: Indicates that an invalid declarative net request entry was encountered.
	WKWebExtensionErrorInvalidDeclarativeNetRequestEntry WKWebExtensionError = 7
	// WKWebExtensionErrorInvalidManifest: Indicates that an invalid `manifest.Json()` was encountered.
	WKWebExtensionErrorInvalidManifest WKWebExtensionError = 4
	// WKWebExtensionErrorInvalidManifestEntry: Indicates that an invalid manifest entry was encountered.
	WKWebExtensionErrorInvalidManifestEntry WKWebExtensionError = 6
	// WKWebExtensionErrorInvalidResourceCodeSignature: Indicates that a resource failed the bundle’s code signature checks.
	WKWebExtensionErrorInvalidResourceCodeSignature WKWebExtensionError = 3
	// WKWebExtensionErrorResourceNotFound: Indicates that a specified resource was not found on disk.
	WKWebExtensionErrorResourceNotFound WKWebExtensionError = 2
	// WKWebExtensionErrorUnknown: Indicates that an unknown error occurred.
	WKWebExtensionErrorUnknown WKWebExtensionError = 1
	// WKWebExtensionErrorUnsupportedManifestVersion: Indicates that the manifest version is not supported.
	WKWebExtensionErrorUnsupportedManifestVersion WKWebExtensionError = 5
)

func (e WKWebExtensionError) String() string {
	switch e {
	case WKWebExtensionErrorInvalidArchive:
		return "WKWebExtensionErrorInvalidArchive"
	case WKWebExtensionErrorInvalidBackgroundPersistence:
		return "WKWebExtensionErrorInvalidBackgroundPersistence"
	case WKWebExtensionErrorInvalidDeclarativeNetRequestEntry:
		return "WKWebExtensionErrorInvalidDeclarativeNetRequestEntry"
	case WKWebExtensionErrorInvalidManifest:
		return "WKWebExtensionErrorInvalidManifest"
	case WKWebExtensionErrorInvalidManifestEntry:
		return "WKWebExtensionErrorInvalidManifestEntry"
	case WKWebExtensionErrorInvalidResourceCodeSignature:
		return "WKWebExtensionErrorInvalidResourceCodeSignature"
	case WKWebExtensionErrorResourceNotFound:
		return "WKWebExtensionErrorResourceNotFound"
	case WKWebExtensionErrorUnknown:
		return "WKWebExtensionErrorUnknown"
	case WKWebExtensionErrorUnsupportedManifestVersion:
		return "WKWebExtensionErrorUnsupportedManifestVersion"
	default:
		return fmt.Sprintf("WKWebExtensionError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/Error/Code
type WKWebExtensionMatchPatternError int

const (
	// WKWebExtensionMatchPatternErrorInvalidHost: Indicates that the host component was invalid.
	WKWebExtensionMatchPatternErrorInvalidHost WKWebExtensionMatchPatternError = 3
	// WKWebExtensionMatchPatternErrorInvalidPath: Indicates that the path component was invalid.
	WKWebExtensionMatchPatternErrorInvalidPath WKWebExtensionMatchPatternError = 4
	// WKWebExtensionMatchPatternErrorInvalidScheme: Indicates that the scheme component was invalid.
	WKWebExtensionMatchPatternErrorInvalidScheme WKWebExtensionMatchPatternError = 2
	// WKWebExtensionMatchPatternErrorUnknown: Indicates that an unknown error occurred.
	WKWebExtensionMatchPatternErrorUnknown WKWebExtensionMatchPatternError = 1
)

func (e WKWebExtensionMatchPatternError) String() string {
	switch e {
	case WKWebExtensionMatchPatternErrorInvalidHost:
		return "WKWebExtensionMatchPatternErrorInvalidHost"
	case WKWebExtensionMatchPatternErrorInvalidPath:
		return "WKWebExtensionMatchPatternErrorInvalidPath"
	case WKWebExtensionMatchPatternErrorInvalidScheme:
		return "WKWebExtensionMatchPatternErrorInvalidScheme"
	case WKWebExtensionMatchPatternErrorUnknown:
		return "WKWebExtensionMatchPatternErrorUnknown"
	default:
		return fmt.Sprintf("WKWebExtensionMatchPatternError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/Options
type WKWebExtensionMatchPatternOptions uint

const (
	// WKWebExtensionMatchPatternOptionsIgnorePaths: Indicates that the host components should be ignored while matching.
	WKWebExtensionMatchPatternOptionsIgnorePaths WKWebExtensionMatchPatternOptions = 2
	// WKWebExtensionMatchPatternOptionsIgnoreSchemes: Indicates that the scheme components should be ignored while matching.
	WKWebExtensionMatchPatternOptionsIgnoreSchemes WKWebExtensionMatchPatternOptions = 1
	// WKWebExtensionMatchPatternOptionsMatchBidirectionally: Indicates that two patterns should be checked in either direction while matching.
	WKWebExtensionMatchPatternOptionsMatchBidirectionally WKWebExtensionMatchPatternOptions = 4
	// WKWebExtensionMatchPatternOptionsNone: Indicates no special matching options.
	WKWebExtensionMatchPatternOptionsNone WKWebExtensionMatchPatternOptions = 0
)

func (e WKWebExtensionMatchPatternOptions) String() string {
	switch e {
	case WKWebExtensionMatchPatternOptionsIgnorePaths:
		return "WKWebExtensionMatchPatternOptionsIgnorePaths"
	case WKWebExtensionMatchPatternOptionsIgnoreSchemes:
		return "WKWebExtensionMatchPatternOptionsIgnoreSchemes"
	case WKWebExtensionMatchPatternOptionsMatchBidirectionally:
		return "WKWebExtensionMatchPatternOptionsMatchBidirectionally"
	case WKWebExtensionMatchPatternOptionsNone:
		return "WKWebExtensionMatchPatternOptionsNone"
	default:
		return fmt.Sprintf("WKWebExtensionMatchPatternOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/MessagePort/Error/Code
type WKWebExtensionMessagePortError int

const (
	// WKWebExtensionMessagePortErrorMessageInvalid: Indicates that the message is invalid.
	WKWebExtensionMessagePortErrorMessageInvalid WKWebExtensionMessagePortError = 3
	// WKWebExtensionMessagePortErrorNotConnected: Indicates that the message port is disconnected.
	WKWebExtensionMessagePortErrorNotConnected WKWebExtensionMessagePortError = 2
	// WKWebExtensionMessagePortErrorUnknown: Indicates that an unknown error occurred.
	WKWebExtensionMessagePortErrorUnknown WKWebExtensionMessagePortError = 1
)

func (e WKWebExtensionMessagePortError) String() string {
	switch e {
	case WKWebExtensionMessagePortErrorMessageInvalid:
		return "WKWebExtensionMessagePortErrorMessageInvalid"
	case WKWebExtensionMessagePortErrorNotConnected:
		return "WKWebExtensionMessagePortErrorNotConnected"
	case WKWebExtensionMessagePortErrorUnknown:
		return "WKWebExtensionMessagePortErrorUnknown"
	default:
		return fmt.Sprintf("WKWebExtensionMessagePortError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/TabChangedProperties
type WKWebExtensionTabChangedProperties uint

const (
	// WKWebExtensionTabChangedPropertiesLoading: Indicates the loading state changed.
	WKWebExtensionTabChangedPropertiesLoading WKWebExtensionTabChangedProperties = 2
	// WKWebExtensionTabChangedPropertiesMuted: Indicates the muted state changed.
	WKWebExtensionTabChangedPropertiesMuted WKWebExtensionTabChangedProperties = 4
	// WKWebExtensionTabChangedPropertiesNone: Indicates nothing changed.
	WKWebExtensionTabChangedPropertiesNone WKWebExtensionTabChangedProperties = 0
	// WKWebExtensionTabChangedPropertiesPinned: Indicates the pinned state changed.
	WKWebExtensionTabChangedPropertiesPinned WKWebExtensionTabChangedProperties = 8
	// WKWebExtensionTabChangedPropertiesPlayingAudio: Indicates the audio playback state changed.
	WKWebExtensionTabChangedPropertiesPlayingAudio WKWebExtensionTabChangedProperties = 16
	// WKWebExtensionTabChangedPropertiesReaderMode: Indicates the reader mode state changed.
	WKWebExtensionTabChangedPropertiesReaderMode WKWebExtensionTabChangedProperties = 32
	// WKWebExtensionTabChangedPropertiesSize: Indicates the size changed.
	WKWebExtensionTabChangedPropertiesSize WKWebExtensionTabChangedProperties = 64
	// WKWebExtensionTabChangedPropertiesTitle: Indicates the title changed.
	WKWebExtensionTabChangedPropertiesTitle WKWebExtensionTabChangedProperties = 128
	// WKWebExtensionTabChangedPropertiesURL: Indicates the URL changed.
	WKWebExtensionTabChangedPropertiesURL WKWebExtensionTabChangedProperties = 256
	// WKWebExtensionTabChangedPropertiesZoomFactor: Indicates the zoom factor changed.
	WKWebExtensionTabChangedPropertiesZoomFactor WKWebExtensionTabChangedProperties = 512
)

func (e WKWebExtensionTabChangedProperties) String() string {
	switch e {
	case WKWebExtensionTabChangedPropertiesLoading:
		return "WKWebExtensionTabChangedPropertiesLoading"
	case WKWebExtensionTabChangedPropertiesMuted:
		return "WKWebExtensionTabChangedPropertiesMuted"
	case WKWebExtensionTabChangedPropertiesNone:
		return "WKWebExtensionTabChangedPropertiesNone"
	case WKWebExtensionTabChangedPropertiesPinned:
		return "WKWebExtensionTabChangedPropertiesPinned"
	case WKWebExtensionTabChangedPropertiesPlayingAudio:
		return "WKWebExtensionTabChangedPropertiesPlayingAudio"
	case WKWebExtensionTabChangedPropertiesReaderMode:
		return "WKWebExtensionTabChangedPropertiesReaderMode"
	case WKWebExtensionTabChangedPropertiesSize:
		return "WKWebExtensionTabChangedPropertiesSize"
	case WKWebExtensionTabChangedPropertiesTitle:
		return "WKWebExtensionTabChangedPropertiesTitle"
	case WKWebExtensionTabChangedPropertiesURL:
		return "WKWebExtensionTabChangedPropertiesURL"
	case WKWebExtensionTabChangedPropertiesZoomFactor:
		return "WKWebExtensionTabChangedPropertiesZoomFactor"
	default:
		return fmt.Sprintf("WKWebExtensionTabChangedProperties(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/WindowState
type WKWebExtensionWindowState int

const (
	// WKWebExtensionWindowStateFullscreen: Indicates a window is in full-screen mode.
	WKWebExtensionWindowStateFullscreen WKWebExtensionWindowState = 3
	// WKWebExtensionWindowStateMaximized: Indicates a window is maximized.
	WKWebExtensionWindowStateMaximized WKWebExtensionWindowState = 2
	// WKWebExtensionWindowStateMinimized: Indicates a window is minimized.
	WKWebExtensionWindowStateMinimized WKWebExtensionWindowState = 1
	// WKWebExtensionWindowStateNormal: Indicates a window is in its normal state.
	WKWebExtensionWindowStateNormal WKWebExtensionWindowState = 0
)

func (e WKWebExtensionWindowState) String() string {
	switch e {
	case WKWebExtensionWindowStateFullscreen:
		return "WKWebExtensionWindowStateFullscreen"
	case WKWebExtensionWindowStateMaximized:
		return "WKWebExtensionWindowStateMaximized"
	case WKWebExtensionWindowStateMinimized:
		return "WKWebExtensionWindowStateMinimized"
	case WKWebExtensionWindowStateNormal:
		return "WKWebExtensionWindowStateNormal"
	default:
		return fmt.Sprintf("WKWebExtensionWindowState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/WindowType
type WKWebExtensionWindowType int

const (
	// WKWebExtensionWindowTypeNormal: Indicates a normal window.
	WKWebExtensionWindowTypeNormal WKWebExtensionWindowType = 0
	// WKWebExtensionWindowTypePopup: Indicates a pop-up window.
	WKWebExtensionWindowTypePopup WKWebExtensionWindowType = 1
)

func (e WKWebExtensionWindowType) String() string {
	switch e {
	case WKWebExtensionWindowTypeNormal:
		return "WKWebExtensionWindowTypeNormal"
	case WKWebExtensionWindowTypePopup:
		return "WKWebExtensionWindowTypePopup"
	default:
		return fmt.Sprintf("WKWebExtensionWindowType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/WKWebViewDataType
type WKWebViewDataType uint

const (
	WKWebViewDataTypeSessionStorage WKWebViewDataType = 1
)

func (e WKWebViewDataType) String() string {
	switch e {
	case WKWebViewDataTypeSessionStorage:
		return "WKWebViewDataTypeSessionStorage"
	default:
		return fmt.Sprintf("WKWebViewDataType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/WKWebpagePreferences/UpgradeToHTTPSPolicy
type WKWebpagePreferencesUpgradeToHTTPSPolicy int

const (
	WKWebpagePreferencesUpgradeToHTTPSPolicyAutomaticFallbackToHTTP    WKWebpagePreferencesUpgradeToHTTPSPolicy = 1
	WKWebpagePreferencesUpgradeToHTTPSPolicyErrorOnFailure             WKWebpagePreferencesUpgradeToHTTPSPolicy = 3
	WKWebpagePreferencesUpgradeToHTTPSPolicyKeepAsRequested            WKWebpagePreferencesUpgradeToHTTPSPolicy = 0
	WKWebpagePreferencesUpgradeToHTTPSPolicyUserMediatedFallbackToHTTP WKWebpagePreferencesUpgradeToHTTPSPolicy = 2
)

func (e WKWebpagePreferencesUpgradeToHTTPSPolicy) String() string {
	switch e {
	case WKWebpagePreferencesUpgradeToHTTPSPolicyAutomaticFallbackToHTTP:
		return "WKWebpagePreferencesUpgradeToHTTPSPolicyAutomaticFallbackToHTTP"
	case WKWebpagePreferencesUpgradeToHTTPSPolicyErrorOnFailure:
		return "WKWebpagePreferencesUpgradeToHTTPSPolicyErrorOnFailure"
	case WKWebpagePreferencesUpgradeToHTTPSPolicyKeepAsRequested:
		return "WKWebpagePreferencesUpgradeToHTTPSPolicyKeepAsRequested"
	case WKWebpagePreferencesUpgradeToHTTPSPolicyUserMediatedFallbackToHTTP:
		return "WKWebpagePreferencesUpgradeToHTTPSPolicyUserMediatedFallbackToHTTP"
	default:
		return fmt.Sprintf("WKWebpagePreferencesUpgradeToHTTPSPolicy(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/WebCacheModel
type WebCacheModel uint

const (
	// Deprecated.
	WebCacheModelDocumentBrowser WebCacheModel = 1
	// Deprecated.
	WebCacheModelDocumentViewer WebCacheModel = 0
	// Deprecated.
	WebCacheModelPrimaryWebBrowser WebCacheModel = 2
)

func (e WebCacheModel) String() string {
	switch e {
	case WebCacheModelDocumentBrowser:
		return "WebCacheModelDocumentBrowser"
	case WebCacheModelDocumentViewer:
		return "WebCacheModelDocumentViewer"
	case WebCacheModelPrimaryWebBrowser:
		return "WebCacheModelPrimaryWebBrowser"
	default:
		return fmt.Sprintf("WebCacheModel(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/WebDragDestinationAction
type WebDragDestinationAction uint

const (
	// Deprecated.
	WebDragDestinationActionAny WebDragDestinationAction = 0
	// Deprecated.
	WebDragDestinationActionDHTML WebDragDestinationAction = 1
	// Deprecated.
	WebDragDestinationActionEdit WebDragDestinationAction = 2
	// Deprecated.
	WebDragDestinationActionLoad WebDragDestinationAction = 4
	// Deprecated.
	WebDragDestinationActionNone WebDragDestinationAction = 0
)

func (e WebDragDestinationAction) String() string {
	switch e {
	case WebDragDestinationActionAny:
		return "WebDragDestinationActionAny"
	case WebDragDestinationActionDHTML:
		return "WebDragDestinationActionDHTML"
	case WebDragDestinationActionEdit:
		return "WebDragDestinationActionEdit"
	case WebDragDestinationActionLoad:
		return "WebDragDestinationActionLoad"
	default:
		return fmt.Sprintf("WebDragDestinationAction(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/WebDragSourceAction
type WebDragSourceAction uint

const (
	// Deprecated.
	WebDragSourceActionAny WebDragSourceAction = 0
	// Deprecated.
	WebDragSourceActionDHTML WebDragSourceAction = 1
	// Deprecated.
	WebDragSourceActionImage WebDragSourceAction = 2
	// Deprecated.
	WebDragSourceActionLink WebDragSourceAction = 4
	// Deprecated.
	WebDragSourceActionNone WebDragSourceAction = 0
	// Deprecated.
	WebDragSourceActionSelection WebDragSourceAction = 8
)

func (e WebDragSourceAction) String() string {
	switch e {
	case WebDragSourceActionAny:
		return "WebDragSourceActionAny"
	case WebDragSourceActionDHTML:
		return "WebDragSourceActionDHTML"
	case WebDragSourceActionImage:
		return "WebDragSourceActionImage"
	case WebDragSourceActionLink:
		return "WebDragSourceActionLink"
	case WebDragSourceActionSelection:
		return "WebDragSourceActionSelection"
	default:
		return fmt.Sprintf("WebDragSourceAction(%d)", e)
	}
}

type WebKitError int

const (
	WebKitErrorBlockedPlugInVersion               WebKitError = 203
	WebKitErrorCannotFindPlugIn                   WebKitError = 200
	WebKitErrorCannotLoadPlugIn                   WebKitError = 201
	WebKitErrorCannotShowMIMEType                 WebKitError = 100
	WebKitErrorCannotShowURL                      WebKitError = 101
	WebKitErrorFrameLoadInterruptedByPolicyChange WebKitError = 102
	WebKitErrorJavaUnavailable                    WebKitError = 202
)

func (e WebKitError) String() string {
	switch e {
	case WebKitErrorBlockedPlugInVersion:
		return "WebKitErrorBlockedPlugInVersion"
	case WebKitErrorCannotFindPlugIn:
		return "WebKitErrorCannotFindPlugIn"
	case WebKitErrorCannotLoadPlugIn:
		return "WebKitErrorCannotLoadPlugIn"
	case WebKitErrorCannotShowMIMEType:
		return "WebKitErrorCannotShowMIMEType"
	case WebKitErrorCannotShowURL:
		return "WebKitErrorCannotShowURL"
	case WebKitErrorFrameLoadInterruptedByPolicyChange:
		return "WebKitErrorFrameLoadInterruptedByPolicyChange"
	case WebKitErrorJavaUnavailable:
		return "WebKitErrorJavaUnavailable"
	default:
		return fmt.Sprintf("WebKitError(%d)", e)
	}
}

type WebMenuItem uint

const (
	// WebMenuItemPDFActualSize: Display a PDF document at its original size.
	WebMenuItemPDFActualSize WebMenuItem = 24
	// WebMenuItemPDFAutoSize: Display a PDF document at a user-specified size.
	WebMenuItemPDFAutoSize WebMenuItem = 27
	// WebMenuItemPDFContinuous: Display all pages in a PDF document continuously, using a vertical scroll bar, if necessary.
	WebMenuItemPDFContinuous WebMenuItem = 30
	// WebMenuItemPDFFacingPages: Display a PDF document two pages at a time.
	WebMenuItemPDFFacingPages WebMenuItem = 29
	// WebMenuItemPDFNextPage: Display the next page of a PDF document.
	WebMenuItemPDFNextPage WebMenuItem = 31
	// WebMenuItemPDFPreviousPage: Display the previous page of a PDF document.
	WebMenuItemPDFPreviousPage WebMenuItem = 32
	// WebMenuItemPDFSinglePage: Display a PDF document one page at a time.
	WebMenuItemPDFSinglePage WebMenuItem = 28
	// WebMenuItemPDFZoomIn: Scale up a PDF document.
	WebMenuItemPDFZoomIn WebMenuItem = 25
	// WebMenuItemPDFZoomOut: Scale down a PDF document.
	WebMenuItemPDFZoomOut WebMenuItem = 26
	// WebMenuItemTagCopy: Copy the element to the clipboard.
	WebMenuItemTagCopy WebMenuItem = 8
	// WebMenuItemTagCopyImageToClipboard: Copy the image to the clipboard.
	WebMenuItemTagCopyImageToClipboard WebMenuItem = 6
	// WebMenuItemTagCopyLinkToClipboard: Copy the link to the clipboard.
	WebMenuItemTagCopyLinkToClipboard WebMenuItem = 3
	// WebMenuItemTagCut: Cut the currently selected content.
	WebMenuItemTagCut WebMenuItem = 13
	// WebMenuItemTagDownloadImageToDisk: Download the image to disk.
	WebMenuItemTagDownloadImageToDisk WebMenuItem = 5
	// WebMenuItemTagDownloadLinkToDisk: Download the link to a disk.
	WebMenuItemTagDownloadLinkToDisk WebMenuItem = 2
	// WebMenuItemTagGoBack: Load the previous page.
	WebMenuItemTagGoBack WebMenuItem = 9
	// WebMenuItemTagGoForward: Load the next page.
	WebMenuItemTagGoForward WebMenuItem = 10
	// WebMenuItemTagIgnoreSpelling: Ignore the misspelled word.
	WebMenuItemTagIgnoreSpelling WebMenuItem = 17
	// WebMenuItemTagLearnSpelling: Add the misspelled word to the user’s list of acceptable words.
	WebMenuItemTagLearnSpelling WebMenuItem = 18
	// WebMenuItemTagLookUpInDictionary: Look up the current selection in the Dictionary.
	WebMenuItemTagLookUpInDictionary WebMenuItem = 22
	// WebMenuItemTagNoGuessesFound: Indicate whether any suggested spellings for the misspelled word could be found.
	WebMenuItemTagNoGuessesFound WebMenuItem = 16
	// WebMenuItemTagOpenFrameInNewWindow: Open the frame in a new window.
	WebMenuItemTagOpenFrameInNewWindow WebMenuItem = 7
	// WebMenuItemTagOpenImageInNewWindow: Open the image in a new window.
	WebMenuItemTagOpenImageInNewWindow WebMenuItem = 4
	// WebMenuItemTagOpenLinkInNewWindow: Open the link in a new window.
	WebMenuItemTagOpenLinkInNewWindow WebMenuItem = 1
	// WebMenuItemTagOpenWithDefaultApplication: Open the current selection using the default application.
	WebMenuItemTagOpenWithDefaultApplication WebMenuItem = 23
	// WebMenuItemTagOther: Used when a tag for an item in the context menu can’t be determined.
	WebMenuItemTagOther WebMenuItem = 19
	// WebMenuItemTagPaste: Paste the content on the clipboard onto the current selection.
	WebMenuItemTagPaste WebMenuItem = 14
	// WebMenuItemTagReload: Reload the current page.
	WebMenuItemTagReload WebMenuItem = 12
	// WebMenuItemTagSearchInSpotlight: Search SpotLight for the current selection.
	WebMenuItemTagSearchInSpotlight WebMenuItem = 20
	// WebMenuItemTagSearchWeb: Search the web for the current selection.
	WebMenuItemTagSearchWeb WebMenuItem = 21
	// WebMenuItemTagSpellingGuess: Suggest spellings for the misspelled word.
	WebMenuItemTagSpellingGuess WebMenuItem = 15
	// WebMenuItemTagStop: Stop loading the current page.
	WebMenuItemTagStop WebMenuItem = 11
)

func (e WebMenuItem) String() string {
	switch e {
	case WebMenuItemPDFActualSize:
		return "WebMenuItemPDFActualSize"
	case WebMenuItemPDFAutoSize:
		return "WebMenuItemPDFAutoSize"
	case WebMenuItemPDFContinuous:
		return "WebMenuItemPDFContinuous"
	case WebMenuItemPDFFacingPages:
		return "WebMenuItemPDFFacingPages"
	case WebMenuItemPDFNextPage:
		return "WebMenuItemPDFNextPage"
	case WebMenuItemPDFPreviousPage:
		return "WebMenuItemPDFPreviousPage"
	case WebMenuItemPDFSinglePage:
		return "WebMenuItemPDFSinglePage"
	case WebMenuItemPDFZoomIn:
		return "WebMenuItemPDFZoomIn"
	case WebMenuItemPDFZoomOut:
		return "WebMenuItemPDFZoomOut"
	case WebMenuItemTagCopy:
		return "WebMenuItemTagCopy"
	case WebMenuItemTagCopyImageToClipboard:
		return "WebMenuItemTagCopyImageToClipboard"
	case WebMenuItemTagCopyLinkToClipboard:
		return "WebMenuItemTagCopyLinkToClipboard"
	case WebMenuItemTagCut:
		return "WebMenuItemTagCut"
	case WebMenuItemTagDownloadImageToDisk:
		return "WebMenuItemTagDownloadImageToDisk"
	case WebMenuItemTagDownloadLinkToDisk:
		return "WebMenuItemTagDownloadLinkToDisk"
	case WebMenuItemTagGoBack:
		return "WebMenuItemTagGoBack"
	case WebMenuItemTagGoForward:
		return "WebMenuItemTagGoForward"
	case WebMenuItemTagIgnoreSpelling:
		return "WebMenuItemTagIgnoreSpelling"
	case WebMenuItemTagLearnSpelling:
		return "WebMenuItemTagLearnSpelling"
	case WebMenuItemTagLookUpInDictionary:
		return "WebMenuItemTagLookUpInDictionary"
	case WebMenuItemTagNoGuessesFound:
		return "WebMenuItemTagNoGuessesFound"
	case WebMenuItemTagOpenFrameInNewWindow:
		return "WebMenuItemTagOpenFrameInNewWindow"
	case WebMenuItemTagOpenImageInNewWindow:
		return "WebMenuItemTagOpenImageInNewWindow"
	case WebMenuItemTagOpenLinkInNewWindow:
		return "WebMenuItemTagOpenLinkInNewWindow"
	case WebMenuItemTagOpenWithDefaultApplication:
		return "WebMenuItemTagOpenWithDefaultApplication"
	case WebMenuItemTagOther:
		return "WebMenuItemTagOther"
	case WebMenuItemTagPaste:
		return "WebMenuItemTagPaste"
	case WebMenuItemTagReload:
		return "WebMenuItemTagReload"
	case WebMenuItemTagSearchInSpotlight:
		return "WebMenuItemTagSearchInSpotlight"
	case WebMenuItemTagSearchWeb:
		return "WebMenuItemTagSearchWeb"
	case WebMenuItemTagSpellingGuess:
		return "WebMenuItemTagSpellingGuess"
	case WebMenuItemTagStop:
		return "WebMenuItemTagStop"
	default:
		return fmt.Sprintf("WebMenuItem(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/WebNavigationType
type WebNavigationType int

const (
	// Deprecated.
	WebNavigationTypeBackForward WebNavigationType = 2
	// Deprecated.
	WebNavigationTypeFormResubmitted WebNavigationType = 4
	// Deprecated.
	WebNavigationTypeFormSubmitted WebNavigationType = 1
	// Deprecated.
	WebNavigationTypeLinkClicked WebNavigationType = 0
	// Deprecated.
	WebNavigationTypeOther WebNavigationType = 5
	// Deprecated.
	WebNavigationTypeReload WebNavigationType = 3
)

func (e WebNavigationType) String() string {
	switch e {
	case WebNavigationTypeBackForward:
		return "WebNavigationTypeBackForward"
	case WebNavigationTypeFormResubmitted:
		return "WebNavigationTypeFormResubmitted"
	case WebNavigationTypeFormSubmitted:
		return "WebNavigationTypeFormSubmitted"
	case WebNavigationTypeLinkClicked:
		return "WebNavigationTypeLinkClicked"
	case WebNavigationTypeOther:
		return "WebNavigationTypeOther"
	case WebNavigationTypeReload:
		return "WebNavigationTypeReload"
	default:
		return fmt.Sprintf("WebNavigationType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/WebKit/WebViewInsertAction
type WebViewInsertAction int

const (
	// Deprecated.
	WebViewInsertActionDropped WebViewInsertAction = 2
	// Deprecated.
	WebViewInsertActionPasted WebViewInsertAction = 1
	// Deprecated.
	WebViewInsertActionTyped WebViewInsertAction = 0
)

func (e WebViewInsertAction) String() string {
	switch e {
	case WebViewInsertActionDropped:
		return "WebViewInsertActionDropped"
	case WebViewInsertActionPasted:
		return "WebViewInsertActionPasted"
	case WebViewInsertActionTyped:
		return "WebViewInsertActionTyped"
	default:
		return fmt.Sprintf("WebViewInsertAction(%d)", e)
	}
}
