// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A service that provides a custom communication channel between the host app and the File Provider extension.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderServiceSource
type NSFileProviderServiceSource interface {
	objectivec.IObject

	// Returns an endpoint object that lets the host app communicate with the File Provider extension.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderServiceSource/makeListenerEndpoint()
	MakeListenerEndpointAndReturnError() (foundation.NSXPCListenerEndpoint, error)

	// A name that uniquely identifies the service (reverse domain name notation is recommended).
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderServiceSource/serviceName
	ServiceName() foundation.NSFileProviderServiceName

	// restricted protocol.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderServiceSource/isRestricted
	IsRestricted() bool
}

// NSFileProviderServiceSourceObject wraps an existing Objective-C object that conforms to the NSFileProviderServiceSource protocol.
type NSFileProviderServiceSourceObject struct {
	objectivec.Object
}

func (o NSFileProviderServiceSourceObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSFileProviderServiceSourceObjectFromID constructs a [NSFileProviderServiceSourceObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSFileProviderServiceSourceObjectFromID(id objc.ID) NSFileProviderServiceSourceObject {
	return NSFileProviderServiceSourceObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns an endpoint object that lets the host app communicate with the File
// Provider extension.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderServiceSource/makeListenerEndpoint()
func (o NSFileProviderServiceSourceObject) MakeListenerEndpointAndReturnError() (foundation.NSXPCListenerEndpoint, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("makeListenerEndpointAndReturnError:"))
	if err != nil {
		return foundation.NSXPCListenerEndpoint{}, err
	}
	return foundation.NSXPCListenerEndpointFromID(rv), nil
}

// A name that uniquely identifies the service (reverse domain name notation
// is recommended).
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderServiceSource/serviceName
func (o NSFileProviderServiceSourceObject) ServiceName() foundation.NSFileProviderServiceName {
	rv := objc.Send[foundation.NSFileProviderServiceName](o.ID, objc.Sel("serviceName"))
	return foundation.NSFileProviderServiceName(rv)
}

// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderServiceSource/isRestricted
func (o NSFileProviderServiceSourceObject) IsRestricted() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isRestricted"))
	return bool(rv)
}
