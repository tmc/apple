// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.
//go:build ios
// +build ios

package networkextension

import (
	"github.com/tmc/apple/objc"
)

// Create a verdict that indicates to the system that all of the new flow’s
// data should be dropped, but allow the user to request access by tapping or
// clicking on a URL.
//
// remediationURLMapKey: The key in the Filter Control Provider’s [remediationMap] dictionary
// corresponding to the URL of the remediation link to give to the user.
//
// remediationButtonTextMapKey: The key in the Filter Control Provider’s [remediationMap] dictionary
// corresponding to the text of the remediation link text to give to the user.
//
// # Return Value
//
// An [NEFilterNewFlowVerdict] object.
//
// # Discussion
//
// When the Filter Data Provider returns this verdict from its “ method, the
// system uses the verdict’s `remediationURLMapKey` and
// `remediationButtonTextMapKey` to look up the remediation URL parameters in
// the [remediationMap] dictionary set by the Filter Control Provider. The
// remediation URL parameters are then inserted into the block page which is
// presented to the user.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterNewFlowVerdict/remediateVerdict(withRemediationURLMapKey:remediationButtonTextMapKey:)
//
// [remediationMap]: https://developer.apple.com/documentation/NetworkExtension/NEFilterControlProvider/remediationMap
//
// [remediationMap]: https://developer.apple.com/documentation/NetworkExtension/NEFilterControlProvider/remediationMap
// [remediationMap]: https://developer.apple.com/documentation/NetworkExtension/NEFilterControlProvider/remediationMap
func (_NEFilterNewFlowVerdictClass NEFilterNewFlowVerdictClass) RemediateVerdictWithRemediationURLMapKeyRemediationButtonTextMapKey(remediationURLMapKey string, remediationButtonTextMapKey string) NEFilterNewFlowVerdict {
	rv := objc.Send[objc.ID](objc.ID(_NEFilterNewFlowVerdictClass.class), objc.Sel("remediateVerdictWithRemediationURLMapKey:remediationButtonTextMapKey:"), objc.String(remediationURLMapKey), objc.String(remediationButtonTextMapKey))
	return NEFilterNewFlowVerdictFromID(rv)
}

// Create a verdict that indicates to the system that the Filter Data Provider
// needs more information before it can make a decision about a new flow.
//
// # Return Value
//
// A [NEFilterNewFlowVerdict] object.
//
// # Discussion
//
// When the Filter Data Provider returns this object from its “ method, the
// system calls the Filter Control Provider’s
// [handleNewFlow(_:completionHandler:)] method.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterNewFlowVerdict/needRules()
//
// [handleNewFlow(_:completionHandler:)]: https://developer.apple.com/documentation/NetworkExtension/NEFilterControlProvider/handleNewFlow(_:completionHandler:)
func (_NEFilterNewFlowVerdictClass NEFilterNewFlowVerdictClass) NeedRulesVerdict() NEFilterNewFlowVerdict {
	rv := objc.Send[objc.ID](objc.ID(_NEFilterNewFlowVerdictClass.class), objc.Sel("needRulesVerdict"))
	return NEFilterNewFlowVerdictFromID(rv)
}

// Create a verdict that indicates to the system that all of the new flow’s
// data should be allowed to pass to its final destination, but a string
// should first be appended to the new flow’s request URL.
//
// urlAppendMapKey: The key in the Filter Control Provider’s [urlAppendStringMap] dictionary
// corresponding to the string to append to the new flow’s request URL.
//
// # Return Value
//
// A [NEFilterNewFlowVerdict] object.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterNewFlowVerdict/urlAppendStringVerdict(withMapKey:)
//
// [urlAppendStringMap]: https://developer.apple.com/documentation/NetworkExtension/NEFilterControlProvider/urlAppendStringMap
func (_NEFilterNewFlowVerdictClass NEFilterNewFlowVerdictClass) URLAppendStringVerdictWithMapKey(urlAppendMapKey string) NEFilterNewFlowVerdict {
	rv := objc.Send[objc.ID](objc.ID(_NEFilterNewFlowVerdictClass.class), objc.Sel("URLAppendStringVerdictWithMapKey:"), objc.String(urlAppendMapKey))
	return NEFilterNewFlowVerdictFromID(rv)
}
