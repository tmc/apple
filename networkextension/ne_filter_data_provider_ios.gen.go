// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.
// +build ios

package networkextension

import (
"github.com/tmc/apple/objc"
)

// Handle a remediation request.
//
// flow: An [NEFilterFlow] object containing information about the flow.
//
// # Return Value
// 
// An [NEFilterRemediationVerdict] object indicating how the system should
// handle the flow of network content.
//
// [NEFilterRemediationVerdict]: https://developer.apple.com/documentation/NetworkExtension/NEFilterRemediationVerdict
//
// # Discussion
// 
// The system calls this method when the user taps or clicks on the
// remediation link in the “block” web page in a WebKit browser object and
// the target of the remediation link is not set to a web page.
// 
// [NEFilterDataProvider] subclasses must override this method.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataProvider/handleRemediation(for:)
func (f NEFilterDataProvider) HandleRemediationForFlow(flow INEFilterFlow) unsafe.Pointer {
rv := objc.Send[unsafe.Pointer](f.ID, objc.Sel("handleRemediationForFlow:"), flow)
return *uintptr(rv)
}
// Handle a rules changed event.
//
// # Discussion
// 
// The system calls this method when the Filter Control Provider calls the
// `notifyRulesChanged` method or returns a [NEFilterControlVerdict] with the
// [updateRules()] property set to YES.
//
// [NEFilterControlVerdict]: https://developer.apple.com/documentation/NetworkExtension/NEFilterControlVerdict
// [updateRules()]: https://developer.apple.com/documentation/NetworkExtension/NEFilterControlVerdict/updateRules()
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataProvider/handleRulesChanged()
func (f NEFilterDataProvider) HandleRulesChanged() {
objc.Send[objc.ID](f.ID, objc.Sel("handleRulesChanged"))
}

