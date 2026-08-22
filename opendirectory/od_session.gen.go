// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ODSession] class.
var (
	_ODSessionClass     ODSessionClass
	_ODSessionClassOnce sync.Once
)

func getODSessionClass() ODSessionClass {
	_ODSessionClassOnce.Do(func() {
		_ODSessionClass = ODSessionClass{class: objc.GetClass("ODSession")}
	})
	return _ODSessionClass
}

// GetODSessionClass returns the class object for ODSession.
func GetODSessionClass() ODSessionClass {
	return getODSessionClass()
}

type ODSessionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (oc ODSessionClass) Class() objc.Class {
	return oc.class
}

// Alloc allocates memory for a new instance of the class.
func (oc ODSessionClass) Alloc() ODSession {
	rv := objc.Send[ODSession](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// An [ODSession] object serves as a Cocoa wrapper for an Open Directory
// session.
//
// # Creating and Accessing Sessions
//
//   - [ODSession.InitWithOptionsError]: Creates a session object directed over proxy to another host.
//
// # Accessing Node Information
//
//   - [ODSession.NodeNamesAndReturnError]: Returns the node names that are registered with this session.
//
// # Instance Properties
//
//   - [ODSession.ConfigurationTemplateNames]
//   - [ODSession.MappingTemplateNames]
//
// # Instance Methods
//
//   - [ODSession.AddConfigurationAuthorizationError]
//   - [ODSession.ConfigurationForNodename]
//   - [ODSession.ConfigurationAuthorizationAllowingUserInteractionError]
//   - [ODSession.DeleteConfigurationAuthorizationError]
//   - [ODSession.DeleteConfigurationWithNodenameAuthorizationError]
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODSession
type ODSession struct {
	objectivec.Object
}

// ODSessionFromID constructs a [ODSession] from an objc.ID.
//
// An [ODSession] object serves as a Cocoa wrapper for an Open Directory
// session.
func ODSessionFromID(id objc.ID) ODSession {
	return ODSession{objectivec.Object{ID: id}}
}

// NOTE: ODSession adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [ODSession] class.
//
// # Creating and Accessing Sessions
//
//   - [IODSession.InitWithOptionsError]: Creates a session object directed over proxy to another host.
//
// # Accessing Node Information
//
//   - [IODSession.NodeNamesAndReturnError]: Returns the node names that are registered with this session.
//
// # Instance Properties
//
//   - [IODSession.ConfigurationTemplateNames]
//   - [IODSession.MappingTemplateNames]
//
// # Instance Methods
//
//   - [IODSession.AddConfigurationAuthorizationError]
//   - [IODSession.ConfigurationForNodename]
//   - [IODSession.ConfigurationAuthorizationAllowingUserInteractionError]
//   - [IODSession.DeleteConfigurationAuthorizationError]
//   - [IODSession.DeleteConfigurationWithNodenameAuthorizationError]
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODSession
type IODSession interface {
	objectivec.IObject

	// Topic: Creating and Accessing Sessions

	// Creates a session object directed over proxy to another host.
	InitWithOptionsError(inOptions foundation.INSDictionary) (ODSession, error)

	// Topic: Accessing Node Information

	// Returns the node names that are registered with this session.
	NodeNamesAndReturnError() (foundation.INSArray, error)

	// Topic: Instance Properties

	ConfigurationTemplateNames() foundation.INSArray
	MappingTemplateNames() foundation.INSArray

	// Topic: Instance Methods

	AddConfigurationAuthorizationError(configuration IODConfiguration, authorization objectivec.IObject) (bool, error)
	ConfigurationForNodename(nodename string) IODConfiguration
	ConfigurationAuthorizationAllowingUserInteractionError(allowInteraction bool) (objectivec.IObject, error)
	DeleteConfigurationAuthorizationError(configuration IODConfiguration, authorization objectivec.IObject) (bool, error)
	DeleteConfigurationWithNodenameAuthorizationError(nodename string, authorization objectivec.IObject) (bool, error)
}

// Init initializes the instance.
func (o ODSession) Init() ODSession {
	rv := objc.Send[ODSession](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o ODSession) Autorelease() ODSession {
	rv := objc.Send[ODSession](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewODSession creates a new ODSession instance.
func NewODSession() ODSession {
	class := getODSessionClass()
	rv := objc.Send[ODSession](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a session object directed over proxy to another host.
//
// inOptions: A dictionary of options to associate with the session. Can be `nil`.
//
// # Return Value
//
// The created session object.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODSession/init(options:)
func NewODSessionWithOptionsError(inOptions foundation.INSDictionary) (ODSession, error) {
	var errorPtr objc.ID
	instance := getODSessionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithOptions:error:"), inOptions, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return ODSession{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return ODSession{}, objc.ErrInitFailed
	}
	return ODSessionFromID(rv), nil
}

// Creates a session object directed over proxy to another host.
//
// inOptions: A dictionary of options to associate with the session. Can be `nil`.
//
// # Return Value
//
// The created session object.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODSession/init(options:)
func (o ODSession) InitWithOptionsError(inOptions foundation.INSDictionary) (ODSession, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](o.ID, objc.Sel("initWithOptions:error:"), inOptions, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return ODSession{}, foundation.NSErrorFrom(errorPtr)
	}
	return ODSessionFromID(rv), nil

}

// Returns the node names that are registered with this session.
//
// # Return Value
//
// The node names registered with this session.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODSession/nodeNames()
func (o ODSession) NodeNamesAndReturnError() (foundation.INSArray, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](o.ID, objc.Sel("nodeNamesAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSArrayFromID(rv), nil

}

// authorization is a [*securityfoundation.SFAuthorization].
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODSession/add(_:authorization:)
func (o ODSession) AddConfigurationAuthorizationError(configuration IODConfiguration, authorization objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](o.ID, objc.Sel("addConfiguration:authorization:error:"), configuration, authorization, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("addConfiguration:authorization:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/OpenDirectory/ODSession/configuration(forNodename:)
func (o ODSession) ConfigurationForNodename(nodename string) IODConfiguration {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("configurationForNodename:"), objc.String(nodename))
	return ODConfigurationFromID(rv)
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODSession/configurationAuthorizationAllowingUserInteraction(_:)
func (o ODSession) ConfigurationAuthorizationAllowingUserInteractionError(allowInteraction bool) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](o.ID, objc.Sel("configurationAuthorizationAllowingUserInteraction:error:"), allowInteraction, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// authorization is a [*securityfoundation.SFAuthorization].
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODSession/delete(_:authorization:)
func (o ODSession) DeleteConfigurationAuthorizationError(configuration IODConfiguration, authorization objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](o.ID, objc.Sel("deleteConfiguration:authorization:error:"), configuration, authorization, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("deleteConfiguration:authorization:error: returned NO with nil NSError")
	}
	return rv, nil

}

// authorization is a [*securityfoundation.SFAuthorization].
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODSession/deleteConfiguration(withNodename:authorization:)
func (o ODSession) DeleteConfigurationWithNodenameAuthorizationError(nodename string, authorization objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](o.ID, objc.Sel("deleteConfigurationWithNodename:authorization:error:"), objc.String(nodename), authorization, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("deleteConfigurationWithNodename:authorization:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Returns a shared instance of the local session.
//
// # Return Value
//
// A shared instance of the local session.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODSession/default()
func (_ODSessionClass ODSessionClass) DefaultSession() ODSession {
	rv := objc.Send[objc.ID](objc.ID(_ODSessionClass.class), objc.Sel("defaultSession"))
	return ODSessionFromID(rv)
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODSession/configurationTemplateNames
func (o ODSession) ConfigurationTemplateNames() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("configurationTemplateNames"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODSession/mappingTemplateNames
func (o ODSession) MappingTemplateNames() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("mappingTemplateNames"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
