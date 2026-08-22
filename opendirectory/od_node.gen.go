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

// The class instance for the [ODNode] class.
var (
	_ODNodeClass     ODNodeClass
	_ODNodeClassOnce sync.Once
)

func getODNodeClass() ODNodeClass {
	_ODNodeClassOnce.Do(func() {
		_ODNodeClass = ODNodeClass{class: objc.GetClass("ODNode")}
	})
	return _ODNodeClass
}

// GetODNodeClass returns the class object for ODNode.
func GetODNodeClass() ODNodeClass {
	return getODNodeClass()
}

type ODNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (oc ODNodeClass) Class() objc.Class {
	return oc.class
}

// Alloc allocates memory for a new instance of the class.
func (oc ODNodeClass) Alloc() ODNode {
	rv := objc.Send[ODNode](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// An [ODNode] object serves as a Cocoa wrapper for an Open Directory node.
//
// # Creating and Initializing a Node
//
//   - [ODNode.InitWithSessionNameError]: Creates a node object with a specified session and name.
//   - [ODNode.InitWithSessionTypeError]: Creates a node object with a specified session and type.
//
// # Querying a Node
//
//   - [ODNode.CustomCallSendDataError]: Returns the result of a custom call to the node.
//   - [ODNode.NodeDetailsForKeysError]: Returns a dictionary containing details about a node.
//   - [ODNode.NodeName]: The node’s name.
//   - [ODNode.SubnodeNamesAndReturnError]: Returns the names of subnodes for the node.
//   - [ODNode.UnreachableSubnodeNamesAndReturnError]: Returns an array of the subnodes of a given node that are currently unreachable.
//
// # Setting Node Credentials
//
//   - [ODNode.SetCredentialsWithRecordTypeRecordNamePasswordError]: Sets credentials for interacting with the node.
//
// # Managing Node Records
//
//   - [ODNode.CreateRecordWithRecordTypeNameAttributesError]: Creates a record in a specified node with specified properties.
//   - [ODNode.RecordWithRecordTypeNameAttributesError]: Returns a record from the node with a specified type and name.
//   - [ODNode.SupportedAttributesForRecordTypeError]: Returns an array of attribute types supported by the node’s records.
//   - [ODNode.SupportedRecordTypesAndReturnError]: Returns an array of the record types supported by the node.
//
// # Instance Properties
//
//   - [ODNode.Configuration]
//
// # Instance Methods
//
//   - [ODNode.AccountPoliciesAndReturnError]
//   - [ODNode.AddAccountPolicyToCategoryError]
//   - [ODNode.CustomFunctionPayloadError]
//   - [ODNode.PasswordContentCheckForRecordNameError]
//   - [ODNode.RemoveAccountPolicyFromCategoryError]
//   - [ODNode.SetAccountPoliciesError]
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNode
type ODNode struct {
	objectivec.Object
}

// ODNodeFromID constructs a [ODNode] from an objc.ID.
//
// An [ODNode] object serves as a Cocoa wrapper for an Open Directory node.
func ODNodeFromID(id objc.ID) ODNode {
	return ODNode{objectivec.Object{ID: id}}
}

// NOTE: ODNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [ODNode] class.
//
// # Creating and Initializing a Node
//
//   - [IODNode.InitWithSessionNameError]: Creates a node object with a specified session and name.
//   - [IODNode.InitWithSessionTypeError]: Creates a node object with a specified session and type.
//
// # Querying a Node
//
//   - [IODNode.CustomCallSendDataError]: Returns the result of a custom call to the node.
//   - [IODNode.NodeDetailsForKeysError]: Returns a dictionary containing details about a node.
//   - [IODNode.NodeName]: The node’s name.
//   - [IODNode.SubnodeNamesAndReturnError]: Returns the names of subnodes for the node.
//   - [IODNode.UnreachableSubnodeNamesAndReturnError]: Returns an array of the subnodes of a given node that are currently unreachable.
//
// # Setting Node Credentials
//
//   - [IODNode.SetCredentialsWithRecordTypeRecordNamePasswordError]: Sets credentials for interacting with the node.
//
// # Managing Node Records
//
//   - [IODNode.CreateRecordWithRecordTypeNameAttributesError]: Creates a record in a specified node with specified properties.
//   - [IODNode.RecordWithRecordTypeNameAttributesError]: Returns a record from the node with a specified type and name.
//   - [IODNode.SupportedAttributesForRecordTypeError]: Returns an array of attribute types supported by the node’s records.
//   - [IODNode.SupportedRecordTypesAndReturnError]: Returns an array of the record types supported by the node.
//
// # Instance Properties
//
//   - [IODNode.Configuration]
//
// # Instance Methods
//
//   - [IODNode.AccountPoliciesAndReturnError]
//   - [IODNode.AddAccountPolicyToCategoryError]
//   - [IODNode.CustomFunctionPayloadError]
//   - [IODNode.PasswordContentCheckForRecordNameError]
//   - [IODNode.RemoveAccountPolicyFromCategoryError]
//   - [IODNode.SetAccountPoliciesError]
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNode
type IODNode interface {
	objectivec.IObject

	// Topic: Creating and Initializing a Node

	// Creates a node object with a specified session and name.
	InitWithSessionNameError(inSession IODSession, inName string) (ODNode, error)
	// Creates a node object with a specified session and type.
	InitWithSessionTypeError(inSession IODSession, inType ODNodeType) (ODNode, error)

	// Topic: Querying a Node

	// Returns the result of a custom call to the node.
	CustomCallSendDataError(inCustomCode int, inSendData foundation.NSData) (foundation.NSData, error)
	// Returns a dictionary containing details about a node.
	NodeDetailsForKeysError(inKeys foundation.INSArray) (foundation.INSDictionary, error)
	// The node’s name.
	NodeName() string
	// Returns the names of subnodes for the node.
	SubnodeNamesAndReturnError() (foundation.INSArray, error)
	// Returns an array of the subnodes of a given node that are currently unreachable.
	UnreachableSubnodeNamesAndReturnError() (foundation.INSArray, error)

	// Topic: Setting Node Credentials

	// Sets credentials for interacting with the node.
	SetCredentialsWithRecordTypeRecordNamePasswordError(inRecordType unsafe.Pointer, inRecordName string, inPassword string) (bool, error)

	// Topic: Managing Node Records

	// Creates a record in a specified node with specified properties.
	CreateRecordWithRecordTypeNameAttributesError(inRecordType unsafe.Pointer, inRecordName string, inAttributes foundation.INSDictionary) (IODRecord, error)
	// Returns a record from the node with a specified type and name.
	RecordWithRecordTypeNameAttributesError(inRecordType unsafe.Pointer, inRecordName string, inAttributes objectivec.IObject) (IODRecord, error)
	// Returns an array of attribute types supported by the node’s records.
	SupportedAttributesForRecordTypeError(inRecordType unsafe.Pointer) (foundation.INSArray, error)
	// Returns an array of the record types supported by the node.
	SupportedRecordTypesAndReturnError() (foundation.INSArray, error)

	// Topic: Instance Properties

	Configuration() IODConfiguration

	// Topic: Instance Methods

	AccountPoliciesAndReturnError() (foundation.INSDictionary, error)
	AddAccountPolicyToCategoryError(policy foundation.INSDictionary, category ODPolicyCategoryType) (bool, error)
	CustomFunctionPayloadError(function string, payload objectivec.IObject) (objectivec.IObject, error)
	PasswordContentCheckForRecordNameError(password string, recordName string) (bool, error)
	RemoveAccountPolicyFromCategoryError(policy foundation.INSDictionary, category ODPolicyCategoryType) (bool, error)
	SetAccountPoliciesError(policies foundation.INSDictionary) (bool, error)
}

// Init initializes the instance.
func (o ODNode) Init() ODNode {
	rv := objc.Send[ODNode](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o ODNode) Autorelease() ODNode {
	rv := objc.Send[ODNode](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewODNode creates a new ODNode instance.
func NewODNode() ODNode {
	class := getODNodeClass()
	rv := objc.Send[ODNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a node object with a specified session and name.
//
// inSession: The session.
//
// inName: The name of the node.
//
// # Return Value
//
// The created node object.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNode/init(session:name:)
func NewODNodeWithSessionNameError(inSession IODSession, inName string) (ODNode, error) {
	var errorPtr objc.ID
	instance := getODNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSession:name:error:"), inSession, objc.String(inName), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return ODNode{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return ODNode{}, objc.ErrInitFailed
	}
	return ODNodeFromID(rv), nil
}

// Creates a node object with a specified session and type.
//
// inSession: The session.
//
// inType: The node type.
//
// # Return Value
//
// The created node object.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNode/init(session:type:)
func NewODNodeWithSessionTypeError(inSession IODSession, inType ODNodeType) (ODNode, error) {
	var errorPtr objc.ID
	instance := getODNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSession:type:error:"), inSession, inType, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return ODNode{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return ODNode{}, objc.ErrInitFailed
	}
	return ODNodeFromID(rv), nil
}

// Creates a node object with a specified session and name.
//
// inSession: The session.
//
// inName: The name of the node.
//
// # Return Value
//
// The created node object.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNode/init(session:name:)
func (o ODNode) InitWithSessionNameError(inSession IODSession, inName string) (ODNode, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](o.ID, objc.Sel("initWithSession:name:error:"), inSession, objc.String(inName), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return ODNode{}, foundation.NSErrorFrom(errorPtr)
	}
	return ODNodeFromID(rv), nil

}

// Creates a node object with a specified session and type.
//
// inSession: The session.
//
// inType: The node type.
//
// # Return Value
//
// The created node object.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNode/init(session:type:)
func (o ODNode) InitWithSessionTypeError(inSession IODSession, inType ODNodeType) (ODNode, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](o.ID, objc.Sel("initWithSession:type:error:"), inSession, inType, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return ODNode{}, foundation.NSErrorFrom(errorPtr)
	}
	return ODNodeFromID(rv), nil

}

// Returns the result of a custom call to the node.
//
// inCustomCode: The custom code to send to the node.
//
// inSendData: Data required by `inCustomCode`. Can be `nil`.
//
// # Return Value
//
// The result of the custom call.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNode/customCall(_:send:)
func (o ODNode) CustomCallSendDataError(inCustomCode int, inSendData foundation.NSData) (foundation.NSData, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](o.ID, objc.Sel("customCall:sendData:error:"), inCustomCode, inSendData, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return foundation.NSData{}, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSDataFromID(rv), nil

}

// Returns a dictionary containing details about a node.
//
// inKeys: An array of keys corresponding to the values returned in the dictionary.
//
// # Return Value
//
// A dictionary containing details about the node corresponding to keys
// specified by `inKeys`.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNode/nodeDetails(forKeys:)
func (o ODNode) NodeDetailsForKeysError(inKeys foundation.INSArray) (foundation.INSDictionary, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](o.ID, objc.Sel("nodeDetailsForKeys:error:"), inKeys, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSDictionaryFromID(rv), nil

}

// Returns the names of subnodes for the node.
//
// # Return Value
//
// An array of subnode names.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNode/subnodeNames()
func (o ODNode) SubnodeNamesAndReturnError() (foundation.INSArray, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](o.ID, objc.Sel("subnodeNamesAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSArrayFromID(rv), nil

}

// Returns an array of the subnodes of a given node that are currently
// unreachable.
//
// # Return Value
//
// An array of unreachable subnodes.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNode/unreachableSubnodeNames()
func (o ODNode) UnreachableSubnodeNamesAndReturnError() (foundation.INSArray, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](o.ID, objc.Sel("unreachableSubnodeNamesAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSArrayFromID(rv), nil

}

// Sets credentials for interacting with the node.
//
// inRecordType: The record type that uses the credentials. Can be `nil`. The default value
// is `kODRecordTypeUsers`.
//
// inRecordName: The username to use to authenticate with the node.
//
// inPassword: The password to use to authenticate with the node.
//
// # Discussion
//
// If this function fails, the previous credentials for the node are used.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNode/setCredentialsWithRecordType(_:recordName:password:)
func (o ODNode) SetCredentialsWithRecordTypeRecordNamePasswordError(inRecordType unsafe.Pointer, inRecordName string, inPassword string) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](o.ID, objc.Sel("setCredentialsWithRecordType:recordName:password:error:"), inRecordType, objc.String(inRecordName), objc.String(inPassword), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setCredentialsWithRecordType:recordName:password:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Creates a record in a specified node with specified properties.
//
// inRecordType: The record’s type.
//
// inRecordName: The record’s name.
//
// inAttributes: A dictionary of key-value pairs representing attributes for the record. Can
// be `nil`.
//
// # Return Value
//
// The created record.
//
// # Discussion
//
// The record is automatically assigned a UUID. This UUID can be overridden if
// one is specified in `inAttributes`.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNode/createRecord(withRecordType:name:attributes:)
func (o ODNode) CreateRecordWithRecordTypeNameAttributesError(inRecordType unsafe.Pointer, inRecordName string, inAttributes foundation.INSDictionary) (IODRecord, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](o.ID, objc.Sel("createRecordWithRecordType:name:attributes:error:"), inRecordType, objc.String(inRecordName), inAttributes, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return ODRecord{}, foundation.NSErrorFrom(errorPtr)
	}
	return ODRecordFromID(rv), nil

}

// Returns a record from the node with a specified type and name.
//
// inRecordType: The type of the record.
//
// inRecordName: The name of the record.
//
// inAttributes: An array of record attributes to be cached before the record is returned.
// Can be `nil`.
//
// # Return Value
//
// The requested record.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNode/record(withRecordType:name:attributes:)
func (o ODNode) RecordWithRecordTypeNameAttributesError(inRecordType unsafe.Pointer, inRecordName string, inAttributes objectivec.IObject) (IODRecord, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](o.ID, objc.Sel("recordWithRecordType:name:attributes:error:"), inRecordType, objc.String(inRecordName), inAttributes, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return ODRecord{}, foundation.NSErrorFrom(errorPtr)
	}
	return ODRecordFromID(rv), nil

}

// Returns an array of attribute types supported by the node’s records.
//
// inRecordType: The record type to list supported attribute types for. Can be `nil`.
//
// # Return Value
//
// An array of supported attribute types.
//
// # Discussion
//
// If `inRecordType` is `nil`, this function returns all attribute types
// supported by all record types of the node; otherwise, only attribute types
// specific to `inRecordType` are returned.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNode/supportedAttributes(forRecordType:)
func (o ODNode) SupportedAttributesForRecordTypeError(inRecordType unsafe.Pointer) (foundation.INSArray, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](o.ID, objc.Sel("supportedAttributesForRecordType:error:"), inRecordType, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSArrayFromID(rv), nil

}

// Returns an array of the record types supported by the node.
//
// # Return Value
//
// An array of supported record types.
//
// # Discussion
//
// If the node does not support checking for supported record types, all
// possible record types are returned.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNode/supportedRecordTypes()
func (o ODNode) SupportedRecordTypesAndReturnError() (foundation.INSArray, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](o.ID, objc.Sel("supportedRecordTypesAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSArrayFromID(rv), nil

}

// See: https://developer.apple.com/documentation/OpenDirectory/ODNode/accountPolicies()
func (o ODNode) AccountPoliciesAndReturnError() (foundation.INSDictionary, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accountPoliciesAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSDictionaryFromID(rv), nil

}

// See: https://developer.apple.com/documentation/OpenDirectory/ODNode/addAccountPolicy(_:toCategory:)
func (o ODNode) AddAccountPolicyToCategoryError(policy foundation.INSDictionary, category ODPolicyCategoryType) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](o.ID, objc.Sel("addAccountPolicy:toCategory:error:"), policy, objc.String(string(category)), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("addAccountPolicy:toCategory:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/OpenDirectory/ODNode/customFunction(_:payload:)
func (o ODNode) CustomFunctionPayloadError(function string, payload objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](o.ID, objc.Sel("customFunction:payload:error:"), objc.String(function), payload, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/OpenDirectory/ODNode/passwordContentCheck(_:forRecordName:)
func (o ODNode) PasswordContentCheckForRecordNameError(password string, recordName string) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](o.ID, objc.Sel("passwordContentCheck:forRecordName:error:"), objc.String(password), objc.String(recordName), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("passwordContentCheck:forRecordName:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/OpenDirectory/ODNode/removeAccountPolicy(_:fromCategory:)
func (o ODNode) RemoveAccountPolicyFromCategoryError(policy foundation.INSDictionary, category ODPolicyCategoryType) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](o.ID, objc.Sel("removeAccountPolicy:fromCategory:error:"), policy, objc.String(string(category)), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("removeAccountPolicy:fromCategory:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/OpenDirectory/ODNode/setAccountPolicies(_:)
func (o ODNode) SetAccountPoliciesError(policies foundation.INSDictionary) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](o.ID, objc.Sel("setAccountPolicies:error:"), policies, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setAccountPolicies:error: returned NO with nil NSError")
	}
	return rv, nil

}

// The node’s name.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNode/nodeName
func (o ODNode) NodeName() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("nodeName"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODNode/configuration
func (o ODNode) Configuration() IODConfiguration {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("configuration"))
	return ODConfigurationFromID(objc.ID(rv))
}
