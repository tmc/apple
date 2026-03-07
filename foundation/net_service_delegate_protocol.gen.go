// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf


// The interface a net service uses to inform its delegate about the state of the service it offers.
//
// See: https://developer.apple.com/documentation/Foundation/NetServiceDelegate
type NSNetServiceDelegate interface {
	objectivec.IObject
}



// NSNetServiceDelegateObject wraps an existing Objective-C object that conforms to the NSNetServiceDelegate protocol.
type NSNetServiceDelegateObject struct {
	objectivec.Object
}
func (o NSNetServiceDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}



// NSNetServiceDelegateObjectFromID constructs a [NSNetServiceDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSNetServiceDelegateObjectFromID(id objc.ID) NSNetServiceDelegateObject {
	return NSNetServiceDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}




// Notifies the delegate that the network is ready to publish the service.
//
// sender: The service that is ready to publish.
//
// # Discussion
// 
// Publication of the service proceeds asynchronously and may still generate a
// call to the delegate’s [NetServiceDidNotPublish] method if an error
// occurs.
//
// See: https://developer.apple.com/documentation/Foundation/NetServiceDelegate/netServiceWillPublish(_:)

func (o NSNetServiceDelegateObject) NetServiceWillPublish(sender INSNetService) {
	
	objc.Send[struct{}](o.ID, objc.Sel("netServiceWillPublish:"), sender)
	}

// Notifies the delegate that a service could not be published.
//
// sender: The service that could not be published.
//
// errorDict: A dictionary containing information about the problem. The dictionary
// contains the keys [NSNetServicesErrorCode] and [NSNetServicesErrorDomain].
//
// # Discussion
// 
// This method may be called long after a [NetServiceWillPublish] message has
// been delivered to the delegate.
//
// See: https://developer.apple.com/documentation/Foundation/NetServiceDelegate/netService(_:didNotPublish:)

func (o NSNetServiceDelegateObject) NetServiceDidNotPublish(sender INSNetService, errorDict INSDictionary) {
	
	objc.Send[struct{}](o.ID, objc.Sel("netService:didNotPublish:"), sender, errorDict)
	}

// Notifies the delegate that a service was successfully published.
//
// sender: The service that was published.
//
// See: https://developer.apple.com/documentation/Foundation/NetServiceDelegate/netServiceDidPublish(_:)

func (o NSNetServiceDelegateObject) NetServiceDidPublish(sender INSNetService) {
	
	objc.Send[struct{}](o.ID, objc.Sel("netServiceDidPublish:"), sender)
	}

// Notifies the delegate that the network is ready to resolve the service.
//
// sender: The service that the network is ready to resolve.
//
// # Discussion
// 
// Resolution of the service proceeds asynchronously and may still generate a
// call to the delegate’s [NetServiceDidNotResolve] method if an error
// occurs.
//
// See: https://developer.apple.com/documentation/Foundation/NetServiceDelegate/netServiceWillResolve(_:)

func (o NSNetServiceDelegateObject) NetServiceWillResolve(sender INSNetService) {
	
	objc.Send[struct{}](o.ID, objc.Sel("netServiceWillResolve:"), sender)
	}

// Informs the delegate that an error occurred during resolution of a given
// service.
//
// sender: The service that did not resolve.
//
// errorDict: A dictionary containing information about the problem. The dictionary
// contains the keys [errorCode] and [errorDomain].
// //
// [errorCode]: https://developer.apple.com/documentation/Foundation/NetService/errorCode-swift.type.property
// [errorDomain]: https://developer.apple.com/documentation/Foundation/NetService/errorDomain
//
// # Discussion
// 
// Clients may try to resolve again upon receiving this error. For example, a
// DNS rotary may yield different IP addresses on different resolution
// requests. A common error condition is that no addresses were resolved
// during the timeout period specified in [ResolveWithTimeout].
//
// See: https://developer.apple.com/documentation/Foundation/NetServiceDelegate/netService(_:didNotResolve:)

func (o NSNetServiceDelegateObject) NetServiceDidNotResolve(sender INSNetService, errorDict INSDictionary) {
	
	objc.Send[struct{}](o.ID, objc.Sel("netService:didNotResolve:"), sender, errorDict)
	}

// Informs the delegate that the address for a given service was resolved.
//
// sender: The service that was resolved.
//
// # Discussion
// 
// The delegate can use the [Addresses] method to retrieve the service’s
// address. If the delegate needs only one address, it can stop the resolution
// process using [Stop]. Otherwise, the resolution will continue until the
// timeout specified in [ResolveWithTimeout] is reached.
//
// See: https://developer.apple.com/documentation/Foundation/NetServiceDelegate/netServiceDidResolveAddress(_:)

func (o NSNetServiceDelegateObject) NetServiceDidResolveAddress(sender INSNetService) {
	
	objc.Send[struct{}](o.ID, objc.Sel("netServiceDidResolveAddress:"), sender)
	}

// Notifies the delegate that the TXT record for a given service has been
// updated.
//
// sender: The service whose TXT record was updated.
//
// data: The new TXT record.
//
// See: https://developer.apple.com/documentation/Foundation/NetServiceDelegate/netService(_:didUpdateTXTRecord:)

func (o NSNetServiceDelegateObject) NetServiceDidUpdateTXTRecordData(sender INSNetService, data INSData) {
	
	objc.Send[struct{}](o.ID, objc.Sel("netService:didUpdateTXTRecordData:"), sender, data)
	}

// Informs the delegate that a [Publish] or [ResolveWithTimeout] request was
// stopped.
//
// sender: The service that stopped.
//
// See: https://developer.apple.com/documentation/Foundation/NetServiceDelegate/netServiceDidStop(_:)

func (o NSNetServiceDelegateObject) NetServiceDidStop(sender INSNetService) {
	
	objc.Send[struct{}](o.ID, objc.Sel("netServiceDidStop:"), sender)
	}

// Called when a client connects to a service managed by Bonjour.
//
// sender: The net service object that the client connected to.
//
// inputStream: A stream object for receiving data from the client.
//
// outputStream: A stream object for sending data to the client.
//
// # Discussion
// 
// When you publish a service, if you set the [NetServiceListenForConnections]
// flag in the service options, the service object accepts connections on
// behalf of your app. Later, when a client connects to that service, the
// service object calls this method to provide the app with a pair of streams
// for communicating with that client.
//
// See: https://developer.apple.com/documentation/Foundation/NetServiceDelegate/netService(_:didAcceptConnectionWith:outputStream:)

func (o NSNetServiceDelegateObject) NetServiceDidAcceptConnectionWithInputStreamOutputStream(sender INSNetService, inputStream INSInputStream, outputStream INSOutputStream) {
	
	objc.Send[struct{}](o.ID, objc.Sel("netService:didAcceptConnectionWithInputStream:outputStream:"), sender, inputStream, outputStream)
	}





// NSNetServiceDelegateConfig holds optional typed callbacks for [NSNetServiceDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate
type NSNetServiceDelegateConfig struct {

	// Using Network Services
	// NetServiceWillPublish — Notifies the delegate that the network is ready to publish the service.
	NetServiceWillPublish func(sender NSNetService)
	// NetServiceDidPublish — Notifies the delegate that a service was successfully published.
	NetServiceDidPublish func(sender NSNetService)
	// NetServiceWillResolve — Notifies the delegate that the network is ready to resolve the service.
	NetServiceWillResolve func(sender NSNetService)
	// NetServiceDidResolveAddress — Informs the delegate that the address for a given service was resolved.
	NetServiceDidResolveAddress func(sender NSNetService)
	// NetServiceDidStop — Informs the delegate that a [publish()](<doc://com.apple.foundation/documentation/Foundation/NetService/publish()>) or [resolve(withTimeout:)](<doc://com.apple.foundation/documentation/Foundation/NetService/resolve(withTimeout:)>) request was stopped.
	NetServiceDidStop func(sender NSNetService)

	// Other Methods
	// NetServiceDidUpdateTXTRecordData — Notifies the delegate that the TXT record for a given service has been updated.
	NetServiceDidUpdateTXTRecordData func(sender NSNetService, data objectivec.Object)
	// NetServiceDidAcceptConnectionWithInputStreamOutputStream — Called when a client connects to a service managed by Bonjour.
	NetServiceDidAcceptConnectionWithInputStreamOutputStream func(sender NSNetService, inputStream NSInputStream, outputStream NSOutputStream)
}

// NewNSNetServiceDelegate creates an Objective-C object implementing the [NSNetServiceDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSNetServiceDelegateObject] satisfies the [NSNetServiceDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate
func NewNSNetServiceDelegate(config NSNetServiceDelegateConfig) NSNetServiceDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSNetServiceDelegate_%d", n)

	var methods []objc.MethodDef

	if config.NetServiceWillPublish != nil {
		fn := config.NetServiceWillPublish
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("netServiceWillPublish:"),
			Fn: func(self objc.ID, _cmd objc.SEL, senderID objc.ID) {
				sender := NSNetServiceFromID(senderID)
				fn(sender)
			},
		})
	}

	if config.NetServiceDidPublish != nil {
		fn := config.NetServiceDidPublish
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("netServiceDidPublish:"),
			Fn: func(self objc.ID, _cmd objc.SEL, senderID objc.ID) {
				sender := NSNetServiceFromID(senderID)
				fn(sender)
			},
		})
	}

	if config.NetServiceWillResolve != nil {
		fn := config.NetServiceWillResolve
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("netServiceWillResolve:"),
			Fn: func(self objc.ID, _cmd objc.SEL, senderID objc.ID) {
				sender := NSNetServiceFromID(senderID)
				fn(sender)
			},
		})
	}

	if config.NetServiceDidResolveAddress != nil {
		fn := config.NetServiceDidResolveAddress
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("netServiceDidResolveAddress:"),
			Fn: func(self objc.ID, _cmd objc.SEL, senderID objc.ID) {
				sender := NSNetServiceFromID(senderID)
				fn(sender)
			},
		})
	}

	if config.NetServiceDidUpdateTXTRecordData != nil {
		fn := config.NetServiceDidUpdateTXTRecordData
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("netService:didUpdateTXTRecordData:"),
			Fn: func(self objc.ID, _cmd objc.SEL, senderID objc.ID, dataID objc.ID) {
				sender := NSNetServiceFromID(senderID)
				data := objectivec.ObjectFromID(dataID)
				fn(sender, data)
			},
		})
	}

	if config.NetServiceDidStop != nil {
		fn := config.NetServiceDidStop
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("netServiceDidStop:"),
			Fn: func(self objc.ID, _cmd objc.SEL, senderID objc.ID) {
				sender := NSNetServiceFromID(senderID)
				fn(sender)
			},
		})
	}

	if config.NetServiceDidAcceptConnectionWithInputStreamOutputStream != nil {
		fn := config.NetServiceDidAcceptConnectionWithInputStreamOutputStream
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("netService:didAcceptConnectionWithInputStream:outputStream:"),
			Fn: func(self objc.ID, _cmd objc.SEL, senderID objc.ID, inputStreamID objc.ID, outputStreamID objc.ID) {
				sender := NSNetServiceFromID(senderID)
				inputStream := NSInputStreamFromID(inputStreamID)
				outputStream := NSOutputStreamFromID(outputStreamID)
				fn(sender, inputStream, outputStream)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSNetServiceDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSNetServiceDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSNetServiceDelegateObjectFromID(instance)
}





