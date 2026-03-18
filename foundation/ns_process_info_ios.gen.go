// Code generated from Apple documentation for Foundation. DO NOT EDIT.
// +build ios

package foundation

import (
"github.com/tmc/apple/objc"
)

// Performs the specified block asynchronously and notifies you if the process
// is about to be suspended.
//
// reason: A string used in debugging to indicate the reason the activity began. This
// parameter must not be `nil` or an empty string.
//
// block: A block containing the work to be performed by the activity. The block has
// no return value and takes the following parameter:
// 
// expired: A Boolean indicating whether the process is about to be suspended.
// If the value is [true], the process is about to be suspended so you should
// take whatever steps are needed to stop in progress work. If it is [false],
// start the planned tasks.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method to perform tasks when your process is executing in the
// background. This method queues `block` for asynchronous execution on a
// concurrent queue. When your process is in the background, the method tries
// to take a task assertion to ensure that your block has time to execute. If
// it is unable to take a task assertion, or if the time allotted for the task
// assertion expires, the system executes your block with the parameter set to
// [true]. If it is able to take the task assertion, it executes the block and
// passes [false] for the expired parameter.
// 
// If your block is still executing and the system need to suspend the
// process, the system executes your block a second time with the `expired`
// parameter set to [true]. Your block must be prepared to handle this case.
// When the expired parameter is [true], stop any in-progress tasks as quickly
// as possible.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/ProcessInfo/performExpiringActivity(withReason:using:)
func (p ProcessInfo) PerformExpiringActivityWithReasonUsingBlock(reason string, block bool) {
objc.Send[objc.ID](p.ID, objc.Sel("performExpiringActivityWithReason:usingBlock:"), objc.String(reason), block)
}

