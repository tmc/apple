// Code generated from Apple documentation for Foundation. DO NOT EDIT.
//go:build ios
// +build ios

package foundation

import (
	"github.com/tmc/apple/objc"
)

// A service type that specifies the Multipath TCP connection policy for
// transmitting data over Wi-Fi and cellular interfaces.
//
// # Discussion
//
// Multipath TCP, defined by the IETF in [RFC 6824], is an extension to TCP
// that permits multiple interfaces to transmit a single data stream. This
// capability allows a seamless handover from Wi-Fi to cellular, aimed at
// making both interfaces more efficient and improving the user experience.
//
// The [MultipathServiceType] property defines which policy the Multipath TCP
// stack uses to schedule traffic across Wi-Fi and cellular interfaces. The
// default value is `none`, meaning Multipath TCP is disabled. You can also
// select handover mode, which provides seamless handover between Wi-Fi and
// cellular.
//
// Multipath TCP requires server support. Resources for Linux-based systems
// are available at [https://mptcp.dev].
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/multipathServiceType-swift.property
//
// [RFC 6824]: https://tools.ietf.org/html/rfc6824
// [https://mptcp.dev]: https://mptcp.dev
func (u URLSessionConfiguration) MultipathServiceType() NSURLSessionMultipathServiceType {
	rv := objc.Send[NSURLSessionMultipathServiceType](u.ID, objc.Sel("multipathServiceType"))
	return NSURLSessionMultipathServiceType(rv)
}
func (u URLSessionConfiguration) SetMultipathServiceType(value NSURLSessionMultipathServiceType) {
	objc.Send[struct{}](u.ID, objc.Sel("setMultipathServiceType:"), value)
}
