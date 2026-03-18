// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// The interface a net service browser uses to inform a delegate about the state of service discovery.
//
// See: https://developer.apple.com/documentation/Foundation/NetServiceBrowserDelegate
type NSNetServiceBrowserDelegate interface {
	objectivec.IObject
}

// NSNetServiceBrowserDelegateObject wraps an existing Objective-C object that conforms to the NSNetServiceBrowserDelegate protocol.
type NSNetServiceBrowserDelegateObject struct {
	objectivec.Object
}
func (o NSNetServiceBrowserDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSNetServiceBrowserDelegateObjectFromID constructs a [NSNetServiceBrowserDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSNetServiceBrowserDelegateObjectFromID(id objc.ID) NSNetServiceBrowserDelegateObject {
	return NSNetServiceBrowserDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate the sender found a domain.
//
// browser: Sender of this delegate message.
//
// domainString: Name of the domain found by `netServiceBrowser`.
//
// moreComing: [true] when `netServiceBrowser` is waiting for additional domains. [false]
// when there are no additional domains.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The delegate uses this message to compile a list of available domains. It
// should wait until `moreDomainsComing` is [false] to do a bulk update of
// user interface elements.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Foundation/NetServiceBrowserDelegate/netServiceBrowser(_:didFindDomain:moreComing:)

func (o NSNetServiceBrowserDelegateObject) NetServiceBrowserDidFindDomainMoreComing(browser INSNetServiceBrowser, domainString string, moreComing bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("netServiceBrowser:didFindDomain:moreComing:"), browser, objc.String(domainString), moreComing)
	}

// Tells the delegate the a domain has disappeared or has become unavailable.
//
// browser: Sender of this delegate message.
//
// domainString: Name of the domain that became unavailable.
//
// moreComing: [true] when `netServiceBrowser` is waiting for additional domains. [false]
// when there are no additional domains.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The delegate uses this message to compile a list of unavailable domains. It
// should wait until `moreDomainsComing` is [false] to do a bulk update of
// user interface elements.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Foundation/NetServiceBrowserDelegate/netServiceBrowser(_:didRemoveDomain:moreComing:)

func (o NSNetServiceBrowserDelegateObject) NetServiceBrowserDidRemoveDomainMoreComing(browser INSNetServiceBrowser, domainString string, moreComing bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("netServiceBrowser:didRemoveDomain:moreComing:"), browser, objc.String(domainString), moreComing)
	}

// Tells the delegate the sender found a service.
//
// browser: Sender of this delegate message.
//
// service: Network service found by `netServiceBrowser`. The delegate can use this
// object to connect to and use the service.
//
// moreComing: [true] when `netServiceBrowser` is waiting for additional services. [false]
// when there are no additional services.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The delegate uses this message to compile a list of available services. It
// should wait until `moreServicesComing` is [false] to do a bulk update of
// user interface elements.
// 
// # Special Considerations
// 
// If the delegate chooses to resolve `netService`, it should retain
// `netService` and set itself as that service’s delegate. The delegate
// should, therefore, release that service when it receives the
// [NetServiceDidResolveAddress] or [NetServiceDidNotResolve] delegate
// messages of the [NSNetService] class.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Foundation/NetServiceBrowserDelegate/netServiceBrowser(_:didFind:moreComing:)

func (o NSNetServiceBrowserDelegateObject) NetServiceBrowserDidFindServiceMoreComing(browser INSNetServiceBrowser, service INSNetService, moreComing bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("netServiceBrowser:didFindService:moreComing:"), browser, service, moreComing)
	}

// Tells the delegate a service has disappeared or has become unavailable.
//
// browser: Sender of this delegate message.
//
// service: Network service that has become unavailable.
//
// moreComing: [true] when `netServiceBrowser` is waiting for additional services. [false]
// when there are no additional services.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The delegate uses this message to compile a list of unavailable services.
// It should wait until `moreServicesComing` is [false] to do a bulk update of
// user interface elements.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Foundation/NetServiceBrowserDelegate/netServiceBrowser(_:didRemove:moreComing:)

func (o NSNetServiceBrowserDelegateObject) NetServiceBrowserDidRemoveServiceMoreComing(browser INSNetServiceBrowser, service INSNetService, moreComing bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("netServiceBrowser:didRemoveService:moreComing:"), browser, service, moreComing)
	}

// Tells the delegate that a search is commencing.
//
// browser: Sender of this delegate message.
//
// # Discussion
// 
// This message is sent to the delegate only if the underlying network layer
// is ready to begin a search. The delegate can use this notification to
// prepare its data structures to receive data.
//
// See: https://developer.apple.com/documentation/Foundation/NetServiceBrowserDelegate/netServiceBrowserWillSearch(_:)

func (o NSNetServiceBrowserDelegateObject) NetServiceBrowserWillSearch(browser INSNetServiceBrowser) {
	
	objc.Send[struct{}](o.ID, objc.Sel("netServiceBrowserWillSearch:"), browser)
	}

// Tells the delegate that a search was not successful.
//
// browser: Sender of this delegate message.
//
// errorDict: Dictionary with the reasons the search was unsuccessful. Use the dictionary
// keys [errorCode] and [errorDomain] to retrieve the error information from
// the dictionary.
// //
// [errorCode]: https://developer.apple.com/documentation/Foundation/NetService/errorCode-swift.type.property
// [errorDomain]: https://developer.apple.com/documentation/Foundation/NetService/errorDomain
//
// See: https://developer.apple.com/documentation/Foundation/NetServiceBrowserDelegate/netServiceBrowser(_:didNotSearch:)

func (o NSNetServiceBrowserDelegateObject) NetServiceBrowserDidNotSearch(browser INSNetServiceBrowser, errorDict INSDictionary) {
	
	objc.Send[struct{}](o.ID, objc.Sel("netServiceBrowser:didNotSearch:"), browser, errorDict)
	}

// Tells the delegate that a search was stopped.
//
// browser: Sender of this delegate message.
//
// # Discussion
// 
// When `netServiceBrowser` receives a [Stop] message from its client,
// `netServiceBrowser` sends a `` message to its delegate. The delegate then
// performs any necessary cleanup.
//
// See: https://developer.apple.com/documentation/Foundation/NetServiceBrowserDelegate/netServiceBrowserDidStopSearch(_:)

func (o NSNetServiceBrowserDelegateObject) NetServiceBrowserDidStopSearch(browser INSNetServiceBrowser) {
	
	objc.Send[struct{}](o.ID, objc.Sel("netServiceBrowserDidStopSearch:"), browser)
	}

// NSNetServiceBrowserDelegateConfig holds optional typed callbacks for [NSNetServiceBrowserDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/foundation/nsnetservicebrowserdelegate
type NSNetServiceBrowserDelegateConfig struct {

	// Using Network Service Browsers
	// NetServiceBrowserWillSearch — Tells the delegate that a search is commencing.
	NetServiceBrowserWillSearch func(browser NSNetServiceBrowser)
	// NetServiceBrowserDidStopSearch — Tells the delegate that a search was stopped.
	NetServiceBrowserDidStopSearch func(browser NSNetServiceBrowser)

	// Other Methods
	// NetServiceBrowserDidFindServiceMoreComing — Tells the delegate the sender found a service.
	NetServiceBrowserDidFindServiceMoreComing func(browser NSNetServiceBrowser, service NSNetService, moreComing bool)
	// NetServiceBrowserDidRemoveServiceMoreComing — Tells the delegate a service has disappeared or has become unavailable.
	NetServiceBrowserDidRemoveServiceMoreComing func(browser NSNetServiceBrowser, service NSNetService, moreComing bool)
}

// NewNSNetServiceBrowserDelegate creates an Objective-C object implementing the [NSNetServiceBrowserDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSNetServiceBrowserDelegateObject] satisfies the [NSNetServiceBrowserDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/foundation/nsnetservicebrowserdelegate
func NewNSNetServiceBrowserDelegate(config NSNetServiceBrowserDelegateConfig) NSNetServiceBrowserDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSNetServiceBrowserDelegate_%d", n)

	var methods []objc.MethodDef

	if config.NetServiceBrowserDidFindServiceMoreComing != nil {
		fn := config.NetServiceBrowserDidFindServiceMoreComing
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("netServiceBrowser:didFindService:moreComing:"),
			Fn: func(self objc.ID, _cmd objc.SEL, browserID objc.ID, serviceID objc.ID, moreComing bool) {
				browser := NSNetServiceBrowserFromID(browserID)
				service := NSNetServiceFromID(serviceID)
				fn(browser, service, moreComing)
			},
		})
	}

	if config.NetServiceBrowserDidRemoveServiceMoreComing != nil {
		fn := config.NetServiceBrowserDidRemoveServiceMoreComing
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("netServiceBrowser:didRemoveService:moreComing:"),
			Fn: func(self objc.ID, _cmd objc.SEL, browserID objc.ID, serviceID objc.ID, moreComing bool) {
				browser := NSNetServiceBrowserFromID(browserID)
				service := NSNetServiceFromID(serviceID)
				fn(browser, service, moreComing)
			},
		})
	}

	if config.NetServiceBrowserWillSearch != nil {
		fn := config.NetServiceBrowserWillSearch
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("netServiceBrowserWillSearch:"),
			Fn: func(self objc.ID, _cmd objc.SEL, browserID objc.ID) {
				browser := NSNetServiceBrowserFromID(browserID)
				fn(browser)
			},
		})
	}

	if config.NetServiceBrowserDidStopSearch != nil {
		fn := config.NetServiceBrowserDidStopSearch
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("netServiceBrowserDidStopSearch:"),
			Fn: func(self objc.ID, _cmd objc.SEL, browserID objc.ID) {
				browser := NSNetServiceBrowserFromID(browserID)
				fn(browser)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSNetServiceBrowserDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSNetServiceBrowserDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSNetServiceBrowserDelegateObjectFromID(instance)
}

