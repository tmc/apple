// Code generated from Apple documentation. DO NOT EDIT.

package webkit

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// BoolHandler handles The completion handler block to execute with the decision.
//   - decision: A Boolean value that indicates whether to continue to use a deprecated version of TLS. Specify [true](<doc://com.apple.documentation/documentation/Swift/true>) to continue, or [false](<doc://com.apple.documentation/documentation/Swift/false>) to reject the connection.
//
// Used by:
//   - [WKNavigationDelegate.WebViewAuthenticationChallengeShouldAllowDeprecatedTLS]
//   - [WKNavigationDelegate.WebViewShouldGoToBackForwardListItemWillUseInstantBackCompletionHandler]
//   - [WKUIDelegate.WebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler]
type BoolHandler = func(bool)

// NewBoolBlock wraps a Go [BoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [WKNavigationDelegate.WebViewAuthenticationChallengeShouldAllowDeprecatedTLS]
//   - [WKNavigationDelegate.WebViewShouldGoToBackForwardListItemWillUseInstantBackCompletionHandler]
//   - [WKUIDelegate.WebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler]
func NewBoolBlock(handler BoolHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal bool) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// DataErrorHandler handles The completion handler to call when the data is ready.
//   - pdfDocumentData: A data object that contains the PDF data to use for rendering the contents of the web view.
//   - error: An error object if a problem occurred, or `nil` on success.
//
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [WKWebView.CreatePDFWithConfigurationCompletionHandler]
//   - [WKWebView.CreateWebArchiveDataWithCompletionHandler]
//   - [WKWebView.FetchDataOfTypesCompletionHandler]
//   - [WKWebsiteDataStore.FetchDataOfTypesCompletionHandler]
type DataErrorHandler = func(*foundation.NSData, error)

// NewDataErrorBlock wraps a Go [DataErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [WKWebView.CreatePDFWithConfigurationCompletionHandler]
//   - [WKWebView.CreateWebArchiveDataWithCompletionHandler]
//   - [WKWebView.FetchDataOfTypesCompletionHandler]
//   - [WKWebsiteDataStore.FetchDataOfTypesCompletionHandler]
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

// DataHandler handles A closure you provide to capture and store data so that you can resume the download later.
//
// Used by:
//   - [WKDownload.Cancel]
type DataHandler = func(*foundation.NSData)

// NewDataBlock wraps a Go [DataHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [WKDownload.Cancel]
func NewDataBlock(handler DataHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *foundation.NSData
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := foundation.NSDataFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// ErrorHandler handles A completion handler block to call after the removal of the content rule list.
//   - error: `nil` on success, or an error object if the store encountered an error when deleting the rule list.
//
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [WKContentRuleListStore.RemoveContentRuleListForIdentifierCompletionHandler]
//   - [WKWebExtensionContext.LoadBackgroundContentWithCompletionHandler]
//   - [WKWebExtensionControllerDelegate.WebExtensionControllerConnectUsingMessagePortForExtensionContextCompletionHandler]
//   - [WKWebExtensionControllerDelegate.WebExtensionControllerOpenOptionsPageForExtensionContextCompletionHandler]
//   - [WKWebExtensionControllerDelegate.WebExtensionControllerPresentPopupForActionForExtensionContextCompletionHandler]
//   - [WKWebExtensionMessagePort.SendMessageCompletionHandler]
//   - [WKWebExtensionTab.ActivateForWebExtensionContextCompletionHandler]
//   - [WKWebExtensionTab.CloseForWebExtensionContextCompletionHandler]
//   - [WKWebExtensionTab.GoBackForWebExtensionContextCompletionHandler]
//   - [WKWebExtensionTab.GoForwardForWebExtensionContextCompletionHandler]
//   - [WKWebExtensionTab.LoadURLForWebExtensionContextCompletionHandler]
//   - [WKWebExtensionTab.ReloadFromOriginForWebExtensionContextCompletionHandler]
//   - [WKWebExtensionTab.SetMutedForWebExtensionContextCompletionHandler]
//   - [WKWebExtensionTab.SetParentTabForWebExtensionContextCompletionHandler]
//   - [WKWebExtensionTab.SetPinnedForWebExtensionContextCompletionHandler]
//   - [WKWebExtensionTab.SetReaderModeActiveForWebExtensionContextCompletionHandler]
//   - [WKWebExtensionTab.SetSelectedForWebExtensionContextCompletionHandler]
//   - [WKWebExtensionTab.SetZoomFactorForWebExtensionContextCompletionHandler]
//   - [WKWebExtensionWindow.CloseForWebExtensionContextCompletionHandler]
//   - [WKWebExtensionWindow.FocusForWebExtensionContextCompletionHandler]
//   - [WKWebExtensionWindow.SetFrameForWebExtensionContextCompletionHandler]
//   - [WKWebExtensionWindow.SetWindowStateForWebExtensionContextCompletionHandler]
//   - [WKWebView.RestoreDataCompletionHandler]
//   - [WKWebsiteDataStore.RemoveDataStoreForIdentifierCompletionHandler]
//   - [WKWebsiteDataStore.RestoreDataCompletionHandler]
type ErrorHandler = func(error)

// NewErrorBlock wraps a Go [ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [WKContentRuleListStore.RemoveContentRuleListForIdentifierCompletionHandler]
//   - [WKWebExtensionContext.LoadBackgroundContentWithCompletionHandler]
//   - [WKWebExtensionControllerDelegate.WebExtensionControllerConnectUsingMessagePortForExtensionContextCompletionHandler]
//   - [WKWebExtensionControllerDelegate.WebExtensionControllerOpenOptionsPageForExtensionContextCompletionHandler]
//   - [WKWebExtensionControllerDelegate.WebExtensionControllerPresentPopupForActionForExtensionContextCompletionHandler]
//   - [WKWebExtensionMessagePort.SendMessageCompletionHandler]
//   - [WKWebExtensionTab.ActivateForWebExtensionContextCompletionHandler]
//   - [WKWebExtensionTab.CloseForWebExtensionContextCompletionHandler]
//   - [WKWebExtensionTab.GoBackForWebExtensionContextCompletionHandler]
//   - [WKWebExtensionTab.GoForwardForWebExtensionContextCompletionHandler]
//   - [WKWebExtensionTab.LoadURLForWebExtensionContextCompletionHandler]
//   - [WKWebExtensionTab.ReloadFromOriginForWebExtensionContextCompletionHandler]
//   - [WKWebExtensionTab.SetMutedForWebExtensionContextCompletionHandler]
//   - [WKWebExtensionTab.SetParentTabForWebExtensionContextCompletionHandler]
//   - [WKWebExtensionTab.SetPinnedForWebExtensionContextCompletionHandler]
//   - [WKWebExtensionTab.SetReaderModeActiveForWebExtensionContextCompletionHandler]
//   - [WKWebExtensionTab.SetSelectedForWebExtensionContextCompletionHandler]
//   - [WKWebExtensionTab.SetZoomFactorForWebExtensionContextCompletionHandler]
//   - [WKWebExtensionWindow.CloseForWebExtensionContextCompletionHandler]
//   - [WKWebExtensionWindow.FocusForWebExtensionContextCompletionHandler]
//   - [WKWebExtensionWindow.SetFrameForWebExtensionContextCompletionHandler]
//   - [WKWebExtensionWindow.SetWindowStateForWebExtensionContextCompletionHandler]
//   - [WKWebView.RestoreDataCompletionHandler]
//   - [WKWebsiteDataStore.RemoveDataStoreForIdentifierCompletionHandler]
//   - [WKWebsiteDataStore.RestoreDataCompletionHandler]
func NewErrorBlock(handler ErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, errID objc.ID) {
		handler(foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// LocaleErrorHandler handles A block that must be called upon completion.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [WKWebExtensionTab.DetectWebpageLocaleForWebExtensionContextCompletionHandler]
type LocaleErrorHandler = func(*foundation.NSLocale, error)

// NewLocaleErrorBlock wraps a Go [LocaleErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [WKWebExtensionTab.DetectWebpageLocaleForWebExtensionContextCompletionHandler]
func NewLocaleErrorBlock(handler LocaleErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *foundation.NSLocale
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := foundation.NSLocaleFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// NSHTTPCookieArrayHandler handles A completion handler block to execute asynchronously with the results.
//   - cookieArray: An array of [HTTPCookie](<doc://com.apple.documentation/documentation/Foundation/HTTPCookie>) objects. If the store contains no cookies, this parameter contains an empty array.
//
// Used by:
//   - [WKHTTPCookieStore.GetAllCookies]
type NSHTTPCookieArrayHandler = func(*[]foundation.NSHTTPCookie)

// NSURLArrayHandler handles The completion handler the system calls after a person dismisses the open panel.
//
// Used by:
//   - [WKUIDelegate.WebViewRunOpenPanelWithParametersInitiatedByFrameCompletionHandler]
type NSURLArrayHandler = func(*[]foundation.NSURL)

// NSUUIDArrayHandler handles A block to invoke with the fetched list of unique identifiers.
//
// Used by:
//   - [WKWebsiteDataStore.FetchAllDataStoreIdentifiers]
type NSUUIDArrayHandler = func(*[]foundation.NSUUID)

// ObjectErrorHandler handles A handler block to execute when script evaluation finishes.
//   - object: The result of the script evaluation, or `nil` if an error occurred.
//   - error: `nil` on success, or an error object with information about the problem.
//
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [WKWebExtensionControllerDelegate.WebExtensionControllerSendMessageToApplicationWithIdentifierForExtensionContextReplyHandler]
//   - [WKWebView.EvaluateJavaScriptCompletionHandler]
type ObjectErrorHandler = func(objectivec.IObject, error)

// NewObjectErrorBlock wraps a Go [ObjectErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [WKWebExtensionControllerDelegate.WebExtensionControllerSendMessageToApplicationWithIdentifierForExtensionContextReplyHandler]
//   - [WKWebView.EvaluateJavaScriptCompletionHandler]
func NewObjectErrorBlock(handler ObjectErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, valID objc.ID, errID objc.ID) {
		var val objectivec.IObject
		if valID != 0 {
			objc.Send[objc.ID](valID, objc.Sel("retain"))
			obj := objectivec.ObjectFromID(valID)
			val = &obj
		}
		handler(val, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// StringHandler handles A reply handler block to execute with the response to send back to the webpage.
//   - reply: An object that contains the data to return to the webpage. Allowed types for this parameter are [NSNumber](<doc://com.apple.documentation/documentation/Foundation/NSNumber>), [NSString](<doc://com.apple.documentation/documentation/Foundation/NSString>), [NSDate](<doc://com.apple.documentation/documentation/Foundation/NSDate>), [NSArray](<doc://com.apple.documentation/documentation/Foundation/NSArray>), [NSDictionary](<doc://com.apple.documentation/documentation/Foundation/NSDictionary>), and [NSNull](<doc://com.apple.documentation/documentation/Foundation/NSNull>). Specify `nil` if an error occurred.
//   - errorMessage: `nil` on success, or a string that describes the error that occurred.
//
// Used by:
//   - [WKScriptMessageHandlerWithReply.UserContentControllerDidReceiveScriptMessageReplyHandler]
//   - [WKUIDelegate.WebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrameCompletionHandler]
type StringHandler = func(*string)

// StringSetDateHandler handles A block to be called with the set of allowed permissions and an optional expiration date.
//
// Used by:
//   - [WKWebExtensionControllerDelegate.WebExtensionControllerPromptForPermissionsInTabForExtensionContextCompletionHandler]
type StringSetDateHandler = func(*foundation.INSSet, *foundation.NSDate)

// UIContextMenuConfigurationHandler handles The completion handler for you to call with information about how you want to handle the interaction.
//   - configuration: The [UIContextMenuConfiguration](<doc://com.apple.documentation/documentation/UIKit/UIContextMenuConfiguration>) object that contains the details of how you want to handle the interaction. Specify `nil` for this parameter if you don’t want to show a contextual menu.
//
// Used by:
//   - [WKUIDelegate.WebViewContextMenuConfigurationForElementCompletionHandler]
type UIContextMenuConfigurationHandler = func(*objectivec.Object)

// NewUIContextMenuConfigurationBlock wraps a Go [UIContextMenuConfigurationHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [WKUIDelegate.WebViewContextMenuConfigurationForElementCompletionHandler]
func NewUIContextMenuConfigurationBlock(handler UIContextMenuConfigurationHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *objectivec.Object
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := objectivec.ObjectFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// UIImageErrorHandler handles The completion handler to call when the image is ready.
//   - snapshotImage: A platform-native image that contains the specified portion of the web view.
//   - error: An error object if a problem occurred, or `nil` on success.
//
// Used by:
//   - [WKWebExtensionTab.TakeSnapshotUsingConfigurationForWebExtensionContextCompletionHandler]
//   - [WKWebView.TakeSnapshotWithConfigurationCompletionHandler]
type UIImageErrorHandler = func(*objc.ID, error)

// URLCredentialHandler handles A closure you must invoke to respond to the authentication challenge.
//   - disposition: The option to use to handle the challenge. For a list of options, see [URLSession.AuthChallengeDisposition](<doc://com.apple.documentation/documentation/Foundation/URLSession/AuthChallengeDisposition>).
//   - credential: The credential to use for authentication when the `disposition` parameter contains the value [URLSession.AuthChallengeDisposition.useCredential](<doc://com.apple.documentation/documentation/Foundation/URLSession/AuthChallengeDisposition/useCredential>). Specify `nil` to continue without a credential.
//
// Used by:
//   - [WKDownloadDelegate.DownloadDidReceiveAuthenticationChallengeCompletionHandler]
//   - [WKNavigationDelegate.WebViewDidReceiveAuthenticationChallengeCompletionHandler]
type URLCredentialHandler = func(*foundation.NSURLCredential)

// NewURLCredentialBlock wraps a Go [URLCredentialHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [WKDownloadDelegate.DownloadDidReceiveAuthenticationChallengeCompletionHandler]
//   - [WKNavigationDelegate.WebViewDidReceiveAuthenticationChallengeCompletionHandler]
func NewURLCredentialBlock(handler URLCredentialHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *foundation.NSURLCredential
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := foundation.NSURLCredentialFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// URLHandler handles A closure you invoke with a destination file URL to begin the download, or `nil` to cancel the download.
//
// Used by:
//   - [WKDownloadDelegate.DownloadDecideDestinationUsingResponseSuggestedFilenameCompletionHandler]
//   - [WKDownloadDelegate.DownloadDecidePlaceholderPolicy]
type URLHandler = func(*foundation.NSURL)

// NewURLBlock wraps a Go [URLHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [WKDownloadDelegate.DownloadDecideDestinationUsingResponseSuggestedFilenameCompletionHandler]
//   - [WKDownloadDelegate.DownloadDecidePlaceholderPolicy]
func NewURLBlock(handler URLHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *foundation.NSURL
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := foundation.NSURLFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// URLSetDateHandler handles A block to be called with the set of allowed URLs and an optional expiration date.
//
// Used by:
//   - [WKWebExtensionControllerDelegate.WebExtensionControllerPromptForPermissionToAccessURLsInTabForExtensionContextCompletionHandler]
type URLSetDateHandler = func(*foundation.INSSet, *foundation.NSDate)

// VoidHandler handles A completion handler block to call with the results.
//   - identifierArray: An array of strings, each of which corresponds to an identifier for a rule list in the data store. Use each string to look up the associated [WKContentRuleList](<doc://com.apple.webkit/documentation/WebKit/WKContentRuleList>) object. If the data store has no rule lists, the array is empty.
//
// Used by:
//   - [WKContentRuleListStore.GetAvailableContentRuleListIdentifiers]
//   - [WKDownloadDelegate.DownloadDidReceivePlaceholderURLCompletionHandler]
//   - [WKHTTPCookieStore.DeleteCookieCompletionHandler]
//   - [WKHTTPCookieStore.SetCookieCompletionHandler]
//   - [WKHTTPCookieStore.SetCookiePolicyCompletionHandler]
//   - [WKHTTPCookieStore.SetCookiesCompletionHandler]
//   - [WKUIDelegate.WebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler]
//   - [WKWebExtensionController.RemoveDataOfTypesFromDataRecordsCompletionHandler]
//   - [WKWebView.CloseAllMediaPresentationsWithCompletionHandler]
//   - [WKWebView.PauseAllMediaPlaybackWithCompletionHandler]
//   - [WKWebView.SetAllMediaPlaybackSuspendedCompletionHandler]
//   - [WKWebView.SetCameraCaptureStateCompletionHandler]
//   - [WKWebView.SetMicrophoneCaptureStateCompletionHandler]
//   - [WKWebsiteDataStore.RemoveDataOfTypesForDataRecordsCompletionHandler]
//   - [WKWebsiteDataStore.RemoveDataOfTypesModifiedSinceCompletionHandler]
type VoidHandler = func()

// NewVoidBlock wraps a Go [VoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [WKContentRuleListStore.GetAvailableContentRuleListIdentifiers]
//   - [WKDownloadDelegate.DownloadDidReceivePlaceholderURLCompletionHandler]
//   - [WKHTTPCookieStore.DeleteCookieCompletionHandler]
//   - [WKHTTPCookieStore.SetCookieCompletionHandler]
//   - [WKHTTPCookieStore.SetCookiePolicyCompletionHandler]
//   - [WKHTTPCookieStore.SetCookiesCompletionHandler]
//   - [WKUIDelegate.WebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler]
//   - [WKWebExtensionController.RemoveDataOfTypesFromDataRecordsCompletionHandler]
//   - [WKWebView.CloseAllMediaPresentationsWithCompletionHandler]
//   - [WKWebView.PauseAllMediaPlaybackWithCompletionHandler]
//   - [WKWebView.SetAllMediaPlaybackSuspendedCompletionHandler]
//   - [WKWebView.SetCameraCaptureStateCompletionHandler]
//   - [WKWebView.SetMicrophoneCaptureStateCompletionHandler]
//   - [WKWebsiteDataStore.RemoveDataOfTypesForDataRecordsCompletionHandler]
//   - [WKWebsiteDataStore.RemoveDataOfTypesModifiedSinceCompletionHandler]
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

// WKContentRuleListErrorHandler handles A completion handler block to call after compilation finishes.
//   - ruleList: The [WKContentRuleList](<doc://com.apple.webkit/documentation/WebKit/WKContentRuleList>) object that encapsulates the compiled rules derived from the `encodedContentRuleList` parameter. This parameter is `nil` if an error occurs during compilation.
//   - error: `nil` on success, or an error object if a problem occurred.
//
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [WKContentRuleListStore.CompileContentRuleListForIdentifierEncodedContentRuleListCompletionHandler]
//   - [WKContentRuleListStore.LookUpContentRuleListForIdentifierCompletionHandler]
type WKContentRuleListErrorHandler = func(*WKContentRuleList, error)

// NewWKContentRuleListErrorBlock wraps a Go [WKContentRuleListErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [WKContentRuleListStore.CompileContentRuleListForIdentifierEncodedContentRuleListCompletionHandler]
//   - [WKContentRuleListStore.LookUpContentRuleListForIdentifierCompletionHandler]
func NewWKContentRuleListErrorBlock(handler WKContentRuleListErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *WKContentRuleList
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := WKContentRuleListFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// WKCookiePolicyHandler handles The completion handler block to execute asynchronously with the cookie policy.
//   - cookiePolicy: A [WKHTTPCookieStore.CookiePolicy](<doc://com.apple.webkit/documentation/WebKit/WKHTTPCookieStore/CookiePolicy>) case that indicates whether the cookie store allows cookie storage.
//
// Used by:
//   - [WKHTTPCookieStore.GetCookiePolicy]
type WKCookiePolicyHandler = func(int)

// NewWKCookiePolicyBlock wraps a Go [WKCookiePolicyHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [WKHTTPCookieStore.GetCookiePolicy]
func NewWKCookiePolicyBlock(handler WKCookiePolicyHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal int) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// WKDialogResultHandler handles A block you must invoke to resume after the web view displays the first use dialog.
//   - dialogResult: A display result case that indicates how the method handled the display request.
//
// Used by:
//   - [WKUIDelegate.WebViewShowLockdownModeFirstUseMessageCompletionHandler]
type WKDialogResultHandler = func(int)

// NewWKDialogResultBlock wraps a Go [WKDialogResultHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [WKUIDelegate.WebViewShowLockdownModeFirstUseMessageCompletionHandler]
func NewWKDialogResultBlock(handler WKDialogResultHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal int) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// WKDownloadHandler handles A closure the system executes when it has started to download the resource.
//
// Used by:
//   - [WKWebView.ResumeDownloadFromResumeDataCompletionHandler]
//   - [WKWebView.StartDownloadUsingRequestCompletionHandler]
type WKDownloadHandler = func(*WKDownload)

// NewWKDownloadBlock wraps a Go [WKDownloadHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [WKWebView.ResumeDownloadFromResumeDataCompletionHandler]
//   - [WKWebView.StartDownloadUsingRequestCompletionHandler]
func NewWKDownloadBlock(handler WKDownloadHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *WKDownload
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := WKDownloadFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// WKDownloadRedirectPolicyHandler handles A closure you must invoke, providing a download redirect policy that indicates whether to proceed with the redirect.
//
// Used by:
//   - [WKDownloadDelegate.DownloadWillPerformHTTPRedirectionNewRequestDecisionHandler]
type WKDownloadRedirectPolicyHandler = func(int)

// NewWKDownloadRedirectPolicyBlock wraps a Go [WKDownloadRedirectPolicyHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [WKDownloadDelegate.DownloadWillPerformHTTPRedirectionNewRequestDecisionHandler]
func NewWKDownloadRedirectPolicyBlock(handler WKDownloadRedirectPolicyHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal int) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// WKFindResultHandler handles The completion handler to call with the results of the search.
//   - result: The object that contains the results of the search.
//
// Used by:
//   - [WKWebView.FindStringWithConfigurationCompletionHandler]
type WKFindResultHandler = func(*WKFindResult)

// NewWKFindResultBlock wraps a Go [WKFindResultHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [WKWebView.FindStringWithConfigurationCompletionHandler]
func NewWKFindResultBlock(handler WKFindResultHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *WKFindResult
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := WKFindResultFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// WKMediaPlaybackStateHandler handles A closure the system executes after the web view determines the current state of media playback.
//
// Used by:
//   - [WKWebView.RequestMediaPlaybackStateWithCompletionHandler]
type WKMediaPlaybackStateHandler = func(int)

// NewWKMediaPlaybackStateBlock wraps a Go [WKMediaPlaybackStateHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [WKWebView.RequestMediaPlaybackStateWithCompletionHandler]
func NewWKMediaPlaybackStateBlock(handler WKMediaPlaybackStateHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal int) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// WKNavigationActionPolicyHandler handles A completion handler block to call with the results about whether to allow or cancel the navigation.
//   - policy: A constant that indicates whether to cancel or allow the navigation. For a list of possible values, see [WKNavigationActionPolicy](<doc://com.apple.webkit/documentation/WebKit/WKNavigationActionPolicy>).
//
// Used by:
//   - [WKNavigationDelegate.WebViewDecidePolicyForNavigationActionDecisionHandler]
type WKNavigationActionPolicyHandler = func(int)

// NewWKNavigationActionPolicyBlock wraps a Go [WKNavigationActionPolicyHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [WKNavigationDelegate.WebViewDecidePolicyForNavigationActionDecisionHandler]
func NewWKNavigationActionPolicyBlock(handler WKNavigationActionPolicyHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal int) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// WKNavigationResponsePolicyHandler handles A completion handler block to call with the results about whether to allow or cancel the navigation.
//   - policy: A constant that indicates whether to cancel or allow the navigation. For a list of possible values, see [WKNavigationResponsePolicy](<doc://com.apple.webkit/documentation/WebKit/WKNavigationResponsePolicy>).
//
// Used by:
//   - [WKNavigationDelegate.WebViewDecidePolicyForNavigationResponseDecisionHandler]
type WKNavigationResponsePolicyHandler = func(int)

// NewWKNavigationResponsePolicyBlock wraps a Go [WKNavigationResponsePolicyHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [WKNavigationDelegate.WebViewDecidePolicyForNavigationResponseDecisionHandler]
func NewWKNavigationResponsePolicyBlock(handler WKNavigationResponsePolicyHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal int) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// WKPermissionDecisionHandler handles A closure that you call from your delegate method.
//
// Used by:
//   - [WKUIDelegate.WebViewRequestDeviceOrientationAndMotionPermissionForOriginInitiatedByFrameDecisionHandler]
//   - [WKUIDelegate.WebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler]
type WKPermissionDecisionHandler = func(int)

// NewWKPermissionDecisionBlock wraps a Go [WKPermissionDecisionHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [WKUIDelegate.WebViewRequestDeviceOrientationAndMotionPermissionForOriginInitiatedByFrameDecisionHandler]
//   - [WKUIDelegate.WebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler]
func NewWKPermissionDecisionBlock(handler WKPermissionDecisionHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal int) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// WKWebExtensionDataRecordArrayHandler handles A block to invoke when the data records have been fetched.
//
// Used by:
//   - [WKWebExtensionController.FetchDataRecordsOfTypesCompletionHandler]
type WKWebExtensionDataRecordArrayHandler = func(*[]WKWebExtensionDataRecord)

// WKWebExtensionDataRecordHandler handles A block to invoke when the data record has been fetched.
//
// Used by:
//   - [WKWebExtensionController.FetchDataRecordOfTypesForExtensionContextCompletionHandler]
type WKWebExtensionDataRecordHandler = func(*WKWebExtensionDataRecord)

// NewWKWebExtensionDataRecordBlock wraps a Go [WKWebExtensionDataRecordHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [WKWebExtensionController.FetchDataRecordOfTypesForExtensionContextCompletionHandler]
func NewWKWebExtensionDataRecordBlock(handler WKWebExtensionDataRecordHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *WKWebExtensionDataRecord
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := WKWebExtensionDataRecordFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// WKWebExtensionErrorHandler handles A block to be called with an initialized web extension, or `nil` if the object could not be initialized due to an error.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [WKWebExtension.ExtensionWithAppExtensionBundleCompletionHandler]
//   - [WKWebExtension.ExtensionWithResourceBaseURLCompletionHandler]
type WKWebExtensionErrorHandler = func(*WKWebExtension, error)

// NewWKWebExtensionErrorBlock wraps a Go [WKWebExtensionErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [WKWebExtension.ExtensionWithAppExtensionBundleCompletionHandler]
//   - [WKWebExtension.ExtensionWithResourceBaseURLCompletionHandler]
func NewWKWebExtensionErrorBlock(handler WKWebExtensionErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *WKWebExtension
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := WKWebExtensionFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// WKWebExtensionMatchPatternSetDateHandler handles A block to be called with the set of allowed match patterns and an optional expiration date.
//
// Used by:
//   - [WKWebExtensionControllerDelegate.WebExtensionControllerPromptForPermissionMatchPatternsInTabForExtensionContextCompletionHandler]
type WKWebExtensionMatchPatternSetDateHandler = func(*foundation.INSSet, *foundation.NSDate)

// WKWebExtensionTabErrorHandler handles A block to be called with the newly created tab or `nil` if the tab wasn’t created.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [WKWebExtensionControllerDelegate.WebExtensionControllerOpenNewTabUsingConfigurationForExtensionContextCompletionHandler]
//   - [WKWebExtensionTab.DuplicateUsingConfigurationForWebExtensionContextCompletionHandler]
type WKWebExtensionTabErrorHandler = func(WKWebExtensionTab, error)

// NewWKWebExtensionTabErrorBlock wraps a Go [WKWebExtensionTabErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [WKWebExtensionControllerDelegate.WebExtensionControllerOpenNewTabUsingConfigurationForExtensionContextCompletionHandler]
//   - [WKWebExtensionTab.DuplicateUsingConfigurationForWebExtensionContextCompletionHandler]
func NewWKWebExtensionTabErrorBlock(handler WKWebExtensionTabErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result WKWebExtensionTab
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			result = WKWebExtensionTabObjectFromID(resultID)
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// WKWebExtensionWindowErrorHandler handles A block to be called with the newly created window or `nil` if the window wasn’t created.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [WKWebExtensionControllerDelegate.WebExtensionControllerOpenNewWindowUsingConfigurationForExtensionContextCompletionHandler]
type WKWebExtensionWindowErrorHandler = func(WKWebExtensionWindow, error)

// NewWKWebExtensionWindowErrorBlock wraps a Go [WKWebExtensionWindowErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [WKWebExtensionControllerDelegate.WebExtensionControllerOpenNewWindowUsingConfigurationForExtensionContextCompletionHandler]
func NewWKWebExtensionWindowErrorBlock(handler WKWebExtensionWindowErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result WKWebExtensionWindow
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			result = WKWebExtensionWindowObjectFromID(resultID)
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// WKWebpagePreferencesHandler handles A completion handler block to call with the results about whether to allow or cancel the navigation.
//   - policy: A constant that indicates whether to cancel or allow the navigation. For a list of possible values, see [WKNavigationActionPolicy](<doc://com.apple.webkit/documentation/WebKit/WKNavigationActionPolicy>).
//   - preferences: The set of preferences to apply to the page if the navigation is allowed. You may pass the object from the `preferences` parameter or configure a new preferences object and pass it instead.
//
// Used by:
//   - [WKNavigationDelegate.WebViewDecidePolicyForNavigationActionPreferencesDecisionHandler]
type WKWebpagePreferencesHandler = func(*WKWebpagePreferences)

// NewWKWebpagePreferencesBlock wraps a Go [WKWebpagePreferencesHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [WKNavigationDelegate.WebViewDecidePolicyForNavigationActionPreferencesDecisionHandler]
func NewWKWebpagePreferencesBlock(handler WKWebpagePreferencesHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *WKWebpagePreferences
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := WKWebpagePreferencesFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// WKWebsiteDataRecordArrayHandler handles The completion handler block to execute asynchronously with the results.
//   - dataRecordArray: An array of [WKWebsiteDataRecord](<doc://com.apple.webkit/documentation/WebKit/WKWebsiteDataRecord>) objects. Each object in this array corresponds to data for one of the requested types. If no records of the requested types exist, this array is empty.
//
// Used by:
//   - [WKWebsiteDataStore.FetchDataRecordsOfTypesCompletionHandler]
type WKWebsiteDataRecordArrayHandler = func(*[]WKWebsiteDataRecord)
