// Code generated from Apple documentation for SystemExtensions. DO NOT EDIT.

package systemextensions

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// A type that receives updates about the progress of a request.
//
// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionRequestDelegate
type OSSystemExtensionRequestDelegate interface {
	objectivec.IObject

	// Tells the delegate that the manager completed the request.
	//
	// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionRequestDelegate/request(_:didFinishWithResult:)
	RequestDidFinishWithResult(request IOSSystemExtensionRequest, result OSSystemExtensionRequestResult)

	// Tells the delegate the manager failed to complete the request.
	//
	// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionRequestDelegate/request(_:didFailWithError:)
	RequestDidFailWithError(request IOSSystemExtensionRequest, error_ foundation.INSError)

	// Tells the delegate that the user must grant approval before the manager can activate the extension.
	//
	// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionRequestDelegate/requestNeedsUserApproval(_:)
	RequestNeedsUserApproval(request IOSSystemExtensionRequest)

	// Tells the delegate that the user has a different version of the extension installed on their system.
	//
	// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionRequestDelegate/request(_:actionForReplacingExtension:withExtension:)
	RequestActionForReplacingExtensionWithExtension(request IOSSystemExtensionRequest, existing IOSSystemExtensionProperties, ext IOSSystemExtensionProperties) OSSystemExtensionReplacementAction
}

// OSSystemExtensionRequestDelegateObject wraps an existing Objective-C object that conforms to the OSSystemExtensionRequestDelegate protocol.
type OSSystemExtensionRequestDelegateObject struct {
	objectivec.Object
}

func (o OSSystemExtensionRequestDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// OSSystemExtensionRequestDelegateObjectFromID constructs a [OSSystemExtensionRequestDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OSSystemExtensionRequestDelegateObjectFromID(id objc.ID) OSSystemExtensionRequestDelegateObject {
	return OSSystemExtensionRequestDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate that the manager completed the request.
//
// request: The request that completed.
//
// result: Additional information about the completion state.
//
// # Discussion
//
// If the request completes with the
// [OSSystemExtensionRequestWillCompleteAfterReboot] result, then the
// extension isn’t active until after the next restart. After restarting,
// the most recently-processed request determines the extension’s state.
// Consider the following scenarios:
//
// - Activate extension and restart: the extension is active upon restarting.
// - Activate extension, deactivate extension, and restart: the extension is
// inactive upon restarting.
//
// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionRequestDelegate/request(_:didFinishWithResult:)
func (o OSSystemExtensionRequestDelegateObject) RequestDidFinishWithResult(request IOSSystemExtensionRequest, result OSSystemExtensionRequestResult) {
	objc.Send[struct{}](o.ID, objc.Sel("request:didFinishWithResult:"), request, result)
}

// Tells the delegate the manager failed to complete the request.
//
// request: The request that failed.
//
// error: The reason the request failed.
//
// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionRequestDelegate/request(_:didFailWithError:)
func (o OSSystemExtensionRequestDelegateObject) RequestDidFailWithError(request IOSSystemExtensionRequest, error_ foundation.INSError) {
	objc.Send[struct{}](o.ID, objc.Sel("request:didFailWithError:"), request, error_)
}

// Tells the delegate that the user must grant approval before the manager can
// activate the extension.
//
// # Discussion
//
// Activating an extension may require explicit user approval to proceed. For
// example, this occurs when the user hasn’t approved the extension. The
// manager calls this method to notify the delegate. Activation remains
// pending until the user grants or denies permission, or until the app quits.
//
// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionRequestDelegate/requestNeedsUserApproval(_:)
func (o OSSystemExtensionRequestDelegateObject) RequestNeedsUserApproval(request IOSSystemExtensionRequest) {
	objc.Send[struct{}](o.ID, objc.Sel("requestNeedsUserApproval:"), request)
}

// Tells the delegate that the user has a different version of the extension
// installed on their system.
//
// request: The request that encountered a conflict.
//
// existing: A properties object that describes the installed version of the extension.
//
// ext: A properties object that describes the updated version of the extension.
//
// # Return Value
//
// A replacement action the manager should take to resolve the conflict.
//
// # Discussion
//
// The manager calls this method when it encounters an existing extension with
// the same team and bundle identifiers, but with different version
// identifiers. It uses the [CFBundleVersion] and [CFBundleShortVersionString]
// identifiers to determine if the existing and new versions differ. The
// delegate must make a decision on whether to replace the existing extension.
//
// Implement this method to return an
// [OSSystemExtensionRequest.ReplacementAction], which tells the manager what
// to do. If you return [OSSystemExtensionReplacementActionCancel], the
// manager aborts the installation and calls [RequestDidFailWithError], with
// the [OSSystemExtensionErrorRequestCanceled] error code.
//
// If the local system has System Extension developer mode enabled, the
// manager always calls this method when it finds an existing installation,
// even if the version identifiers match.
//
// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionRequestDelegate/request(_:actionForReplacingExtension:withExtension:)
//
// [CFBundleShortVersionString]: https://developer.apple.com/documentation/BundleResources/Information-Property-List/CFBundleShortVersionString
// [CFBundleVersion]: https://developer.apple.com/documentation/BundleResources/Information-Property-List/CFBundleVersion
// [OSSystemExtensionRequest.ReplacementAction]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionRequest/ReplacementAction
func (o OSSystemExtensionRequestDelegateObject) RequestActionForReplacingExtensionWithExtension(request IOSSystemExtensionRequest, existing IOSSystemExtensionProperties, ext IOSSystemExtensionProperties) OSSystemExtensionReplacementAction {
	rv := objc.Send[OSSystemExtensionReplacementAction](o.ID, objc.Sel("request:actionForReplacingExtension:withExtension:"), request, existing, ext)
	return rv
}

// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionRequestDelegate/request(_:foundProperties:)
func (o OSSystemExtensionRequestDelegateObject) RequestFoundProperties(request IOSSystemExtensionRequest, properties []OSSystemExtensionProperties) {
	objc.Send[struct{}](o.ID, objc.Sel("request:foundProperties:"), request, objectivec.IObjectSliceToNSArray(properties))
}

// OSSystemExtensionRequestDelegateConfig holds optional typed callbacks for [OSSystemExtensionRequestDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/systemextensions/ossystemextensionrequestdelegate
type OSSystemExtensionRequestDelegateConfig struct {

	// Handling Success and Failure
	// RequestDidFinishWithResult — Tells the delegate that the manager completed the request.
	RequestDidFinishWithResult func(request OSSystemExtensionRequest, result OSSystemExtensionRequestResult)
	// RequestDidFailWithError — Tells the delegate the manager failed to complete the request.
	RequestDidFailWithError func(request OSSystemExtensionRequest, error_ foundation.NSError)

	// Handling Indeterminate Installs
	// RequestNeedsUserApproval — Tells the delegate that the user must grant approval before the manager can activate the extension.
	RequestNeedsUserApproval func(request OSSystemExtensionRequest)
	// RequestActionForReplacingExtensionWithExtension — Tells the delegate that the user has a different version of the extension installed on their system.
	RequestActionForReplacingExtensionWithExtension func(request OSSystemExtensionRequest, existing OSSystemExtensionProperties, ext OSSystemExtensionProperties) OSSystemExtensionReplacementAction
}

// NewOSSystemExtensionRequestDelegate creates an Objective-C object implementing the [OSSystemExtensionRequestDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [OSSystemExtensionRequestDelegateObject] satisfies the [OSSystemExtensionRequestDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/systemextensions/ossystemextensionrequestdelegate
func NewOSSystemExtensionRequestDelegate(config OSSystemExtensionRequestDelegateConfig) OSSystemExtensionRequestDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoOSSystemExtensionRequestDelegate_%d", n)

	var methods []objc.MethodDef

	if config.RequestDidFinishWithResult != nil {
		fn := config.RequestDidFinishWithResult
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("request:didFinishWithResult:"),
			Fn: func(self objc.ID, _cmd objc.SEL, requestID objc.ID, result OSSystemExtensionRequestResult) {
				request := OSSystemExtensionRequestFromID(requestID)
				fn(request, result)
			},
		})
	}

	if config.RequestDidFailWithError != nil {
		fn := config.RequestDidFailWithError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("request:didFailWithError:"),
			Fn: func(self objc.ID, _cmd objc.SEL, requestID objc.ID, error_ID objc.ID) {
				request := OSSystemExtensionRequestFromID(requestID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(request, error_)
			},
		})
	}

	if config.RequestNeedsUserApproval != nil {
		fn := config.RequestNeedsUserApproval
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("requestNeedsUserApproval:"),
			Fn: func(self objc.ID, _cmd objc.SEL, requestID objc.ID) {
				request := OSSystemExtensionRequestFromID(requestID)
				fn(request)
			},
		})
	}

	if config.RequestActionForReplacingExtensionWithExtension != nil {
		fn := config.RequestActionForReplacingExtensionWithExtension
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("request:actionForReplacingExtension:withExtension:"),
			Fn: func(self objc.ID, _cmd objc.SEL, requestID objc.ID, existingID objc.ID, extID objc.ID) OSSystemExtensionReplacementAction {
				request := OSSystemExtensionRequestFromID(requestID)
				existing := OSSystemExtensionPropertiesFromID(existingID)
				ext := OSSystemExtensionPropertiesFromID(extID)
				return fn(request, existing, ext)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("OSSystemExtensionRequestDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewOSSystemExtensionRequestDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return OSSystemExtensionRequestDelegateObjectFromID(instance)
}
