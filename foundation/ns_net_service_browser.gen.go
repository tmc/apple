// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NetServiceBrowser] class.
var (
	_NetServiceBrowserClass     NetServiceBrowserClass
	_NetServiceBrowserClassOnce sync.Once
)

func getNetServiceBrowserClass() NetServiceBrowserClass {
	_NetServiceBrowserClassOnce.Do(func() {
		_NetServiceBrowserClass = NetServiceBrowserClass{class: objc.GetClass("NSNetServiceBrowser")}
	})
	return _NetServiceBrowserClass
}

// GetNetServiceBrowserClass returns the class object for NSNetServiceBrowser.
func GetNetServiceBrowserClass() NetServiceBrowserClass {
	return getNetServiceBrowserClass()
}

type NetServiceBrowserClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NetServiceBrowserClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NetServiceBrowserClass) Alloc() NetServiceBrowser {
	rv := objc.Send[NetServiceBrowser](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A network service browser that finds published services on a network using
// multicast DNS.
//
// # Overview
//
// Services can range from standard services, such as HTTP and FTP, to custom
// services defined by other applications. You can use a network service
// browser in your code to obtain the list of accessible domains and then to
// obtain an [NSNetService] object for each discovered service. Each network
// service browser performs one search at a time, so if you want to perform
// multiple simultaneous searches, use multiple network service browsers.
//
// A network service browser performs all searches asynchronously using the
// current run loop to execute the search in the background. Results from a
// search are returned through the associated delegate object, which your
// client application must provide. Searching proceeds in the background until
// the object receives a [Stop] message.
//
// To use an [NSNetServiceBrowser] object to search for services, allocate it,
// initialize it, and assign a delegate. (If you wish, you can also use the
// [ScheduleInRunLoopForMode] and [RemoveFromRunLoopForMode] methods to
// execute searches on a run loop other than the current one.) Once your
// object is ready, you begin by gathering the list of accessible domains
// using either the [SearchForRegistrationDomains] or
// [SearchForBrowsableDomains] methods. From the list of returned domains, you
// can pick one and use the [SearchForServicesOfTypeInDomain] method to search
// for services in that domain.
//
// The [NSNetServiceBrowser] class provides two ways to search for domains. In
// most cases, your client should use the [SearchForRegistrationDomains]
// method to search only for local domains to which the host machine has
// registration authority. This is the preferred method for accessing domains
// as it guarantees that the host machine can connect to services in the
// returned domains. Access to domains outside this list may be more limited.
//
// # Configuring Network Service Browsers
//
//   - [NetServiceBrowser.Delegate]: The delegate object for this instance.
//   - [NetServiceBrowser.SetDelegate]
//   - [NetServiceBrowser.IncludesPeerToPeer]: Whether to browse over peer-to-peer Bluetooth and Wi-Fi, if available.
//   - [NetServiceBrowser.SetIncludesPeerToPeer]
//
// See: https://developer.apple.com/documentation/Foundation/NetServiceBrowser
type NetServiceBrowser struct {
	objectivec.Object
}

// NetServiceBrowserFromID constructs a [NetServiceBrowser] from an objc.ID.
//
// A network service browser that finds published services on a network using
// multicast DNS.
func NetServiceBrowserFromID(id objc.ID) NetServiceBrowser {
	return NSNetServiceBrowser{objectivec.Object{ID: id}}
}

// NSNetServiceBrowserFromID is an alias for [NetServiceBrowserFromID] for cross-framework compatibility.
func NSNetServiceBrowserFromID(id objc.ID) NetServiceBrowser { return NetServiceBrowserFromID(id) }

// NOTE: NetServiceBrowser adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NetServiceBrowser] class.
//
// # Configuring Network Service Browsers
//
//   - [INetServiceBrowser.Delegate]: The delegate object for this instance.
//   - [INetServiceBrowser.SetDelegate]
//   - [INetServiceBrowser.IncludesPeerToPeer]: Whether to browse over peer-to-peer Bluetooth and Wi-Fi, if available.
//   - [INetServiceBrowser.SetIncludesPeerToPeer]
//
// See: https://developer.apple.com/documentation/Foundation/NetServiceBrowser
type INetServiceBrowser interface {
	objectivec.IObject

	// Topic: Configuring Network Service Browsers

	// The delegate object for this instance.
	Delegate() NSNetServiceBrowserDelegate
	SetDelegate(value NSNetServiceBrowserDelegate)
	// Whether to browse over peer-to-peer Bluetooth and Wi-Fi, if available.
	IncludesPeerToPeer() bool
	SetIncludesPeerToPeer(value bool)
}

// Init initializes the instance.
func (n NetServiceBrowser) Init() NetServiceBrowser {
	rv := objc.Send[NetServiceBrowser](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NetServiceBrowser) Autorelease() NetServiceBrowser {
	rv := objc.Send[NetServiceBrowser](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNetServiceBrowser creates a new NetServiceBrowser instance.
func NewNetServiceBrowser() NetServiceBrowser {
	class := getNetServiceBrowserClass()
	rv := objc.Send[NetServiceBrowser](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The delegate object for this instance.
//
// See: https://developer.apple.com/documentation/Foundation/NetServiceBrowser/delegate
func (n NetServiceBrowser) Delegate() NSNetServiceBrowserDelegate {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("delegate"))
	return NSNetServiceBrowserDelegateObjectFromID(rv)
}
func (n NetServiceBrowser) SetDelegate(value NSNetServiceBrowserDelegate) {
	objc.Send[struct{}](n.ID, objc.Sel("setDelegate:"), value)
}

// Whether to browse over peer-to-peer Bluetooth and Wi-Fi, if available.
//
// # Discussion
//
// This property must be set before initiating a search to have an effect.
//
// See: https://developer.apple.com/documentation/Foundation/NetServiceBrowser/includesPeerToPeer
func (n NetServiceBrowser) IncludesPeerToPeer() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("includesPeerToPeer"))
	return rv
}
func (n NetServiceBrowser) SetIncludesPeerToPeer(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setIncludesPeerToPeer:"), value)
}
