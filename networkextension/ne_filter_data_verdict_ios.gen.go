// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.
//go:build ios
// +build ios

package networkextension

import (
	"github.com/tmc/apple/objc"
)

// Creates a verdict to drop the current chunk of network data and all
// subsequent data for the current flow, and provides a remediation URL.
//
// remediationURLMapKey: The key in the Filter Control Provider’s [remediationMap] dictionary
// corresponding to the URL of the remediation link to give to the user.
//
// remediationButtonTextMapKey: The key in the Filter Control Provider’s `remediationMap` dictionary that
// corresponds to the text of the remediation link text to give to the user.
//
// # Return Value
//
// A [NEFilterDataVerdict] object.
//
// # Discussion
//
// When the Filter Data Provider returns this verdict from one of its data
// filtering methods, the system resolves the verdict as follows:
//
// - The system uses the verdict’s `remediationURLMapKey` and
// `remediationButtonTextMapKey` to look up the remediation URL parameters in
// the [remediationMap] dictionary set by the Filter Control Provider. - The
// system then inserts the remediation URL parameters into the block page and
// presents it to the user.
//
// The user can tap the URL to appeal the decision to drop the flow. This
// starts the remediation process, if your app provides one.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataVerdict/remediateVerdict(withRemediationURLMapKey:remediationButtonTextMapKey:)
//
// [remediationMap]: https://developer.apple.com/documentation/NetworkExtension/NEFilterControlProvider/remediationMap
//
// [remediationMap]: https://developer.apple.com/documentation/NetworkExtension/NEFilterControlProvider/remediationMap
func (_NEFilterDataVerdictClass NEFilterDataVerdictClass) RemediateVerdictWithRemediationURLMapKeyRemediationButtonTextMapKey(remediationURLMapKey string, remediationButtonTextMapKey string) NEFilterDataVerdict {
	rv := objc.Send[objc.ID](objc.ID(_NEFilterDataVerdictClass.class), objc.Sel("remediateVerdictWithRemediationURLMapKey:remediationButtonTextMapKey:"), objc.String(remediationURLMapKey), objc.String(remediationButtonTextMapKey))
	return NEFilterDataVerdictFromID(rv)
}

// Creates a verdict that tells the system that the Filter Control Provider
// needs to update the rules before making a decision about the flow’s data.
//
// # Return Value
//
// A [NEFilterDataVerdict] object.
//
// # Discussion
//
// When the Filter Data Provider returns this verdict, the system passes the
// flow to the Filter Control Provider’s
// [handleNewFlow(_:completionHandler:)] method.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataVerdict/needRules()
//
// [handleNewFlow(_:completionHandler:)]: https://developer.apple.com/documentation/NetworkExtension/NEFilterControlProvider/handleNewFlow(_:completionHandler:)
func (_NEFilterDataVerdictClass NEFilterDataVerdictClass) NeedRulesVerdict() NEFilterDataVerdict {
	rv := objc.Send[objc.ID](objc.ID(_NEFilterDataVerdictClass.class), objc.Sel("needRulesVerdict"))
	return NEFilterDataVerdictFromID(rv)
}
