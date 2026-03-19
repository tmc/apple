// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// A set of methods that provides the name of the promised file and writes the file to the destination directory when the file promise is fulfilled.
//
// See: https://developer.apple.com/documentation/AppKit/NSFilePromiseProviderDelegate
type NSFilePromiseProviderDelegate interface {
	objectivec.IObject

	// Provides the drag destination file’s name.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFilePromiseProviderDelegate/filePromiseProvider(_:fileNameForType:)
	FilePromiseProviderFileNameForType(filePromiseProvider INSFilePromiseProvider, fileType string) string

	// Writes the contents of a promise to the specified URL.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSFilePromiseProviderDelegate/filePromiseProvider(_:writePromiseTo:completionHandler:)
	FilePromiseProviderWritePromiseToURLCompletionHandler(filePromiseProvider INSFilePromiseProvider, url foundation.INSURL, completionHandler ErrorHandler)
}

// NSFilePromiseProviderDelegateObject wraps an existing Objective-C object that conforms to the NSFilePromiseProviderDelegate protocol.
type NSFilePromiseProviderDelegateObject struct {
	objectivec.Object
}
func (o NSFilePromiseProviderDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSFilePromiseProviderDelegateObjectFromID constructs a [NSFilePromiseProviderDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSFilePromiseProviderDelegateObjectFromID(id objc.ID) NSFilePromiseProviderDelegateObject {
	return NSFilePromiseProviderDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Provides the drag destination file’s name.
//
// filePromiseProvider: The file promise provider.
//
// fileType: A string describing the type of file being provided.
//
// # Discussion
// 
// This method is called when the drag destination fulfills the file promise.
// Determine and return the final filename (a base filename, not a full path).
//
// See: https://developer.apple.com/documentation/AppKit/NSFilePromiseProviderDelegate/filePromiseProvider(_:fileNameForType:)
func (o NSFilePromiseProviderDelegateObject) FilePromiseProviderFileNameForType(filePromiseProvider INSFilePromiseProvider, fileType string) string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("filePromiseProvider:fileNameForType:"), filePromiseProvider, objc.String(fileType))
	return foundation.NSStringFromID(rv).String()
	}
// Writes the contents of a promise to the specified URL.
//
// filePromiseProvider: The file promise provider.
//
// url: The destination URL to write to.
//
// completionHandler: A completion handler to execute after the file has been written.
//
// # Discussion
// 
// This method is called after the drag is complete. The request executes on
// the [OperationQueue] supplied by [OperationQueueForFilePromiseProvider].
// 
// Call the completion handler with the file contents wrapped in
// [NSFileCoordinator]. Be sure to write your file to the input `url`
// parameter.
//
// [NSFileCoordinator]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator
// [OperationQueue]: https://developer.apple.com/documentation/Foundation/OperationQueue
//
// See: https://developer.apple.com/documentation/AppKit/NSFilePromiseProviderDelegate/filePromiseProvider(_:writePromiseTo:completionHandler:)
func (o NSFilePromiseProviderDelegateObject) FilePromiseProviderWritePromiseToURLCompletionHandler(filePromiseProvider INSFilePromiseProvider, url foundation.INSURL, completionHandler ErrorHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("filePromiseProvider:writePromiseToURL:completionHandler:"), filePromiseProvider, url, completionHandler)
	}
// Returns the operation queue from which to issue the write request.
//
// filePromiseProvider: The file promise provider for the operation queue.
//
// # Discussion
// 
// If this method isn’t implemented, the main operation queue is used.
// Although this method is optional, to avoid blocking your main thread,
// provide an operation queue other than the main operation queue.
//
// See: https://developer.apple.com/documentation/AppKit/NSFilePromiseProviderDelegate/operationQueue(for:)
func (o NSFilePromiseProviderDelegateObject) OperationQueueForFilePromiseProvider(filePromiseProvider INSFilePromiseProvider) foundation.NSOperationQueue {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("operationQueueForFilePromiseProvider:"), filePromiseProvider)
	return foundation.NSOperationQueueFromID(rv)
	}

// NSFilePromiseProviderDelegateConfig holds optional typed callbacks for [NSFilePromiseProviderDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nsfilepromiseproviderdelegate
type NSFilePromiseProviderDelegateConfig struct {

	// Other Methods
	// OperationQueueForFilePromiseProvider — Returns the operation queue from which to issue the write request.
	OperationQueueForFilePromiseProvider func(filePromiseProvider NSFilePromiseProvider) foundation.NSOperationQueue
}

// NewNSFilePromiseProviderDelegate creates an Objective-C object implementing the [NSFilePromiseProviderDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSFilePromiseProviderDelegateObject] satisfies the [NSFilePromiseProviderDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nsfilepromiseproviderdelegate
func NewNSFilePromiseProviderDelegate(config NSFilePromiseProviderDelegateConfig) NSFilePromiseProviderDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSFilePromiseProviderDelegate_%d", n)

	var methods []objc.MethodDef

	if config.OperationQueueForFilePromiseProvider != nil {
		fn := config.OperationQueueForFilePromiseProvider
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("operationQueueForFilePromiseProvider:"),
			Fn: func(self objc.ID, _cmd objc.SEL, filePromiseProviderID objc.ID) objc.ID {
				filePromiseProvider := NSFilePromiseProviderFromID(filePromiseProviderID)
				return fn(filePromiseProvider).GetID()
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSFilePromiseProviderDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSFilePromiseProviderDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSFilePromiseProviderDelegateObjectFromID(instance)
}

