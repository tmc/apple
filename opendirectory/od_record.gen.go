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

// The class instance for the [ODRecord] class.
var (
	_ODRecordClass     ODRecordClass
	_ODRecordClassOnce sync.Once
)

func getODRecordClass() ODRecordClass {
	_ODRecordClassOnce.Do(func() {
		_ODRecordClass = ODRecordClass{class: objc.GetClass("ODRecord")}
	})
	return _ODRecordClass
}

// GetODRecordClass returns the class object for ODRecord.
func GetODRecordClass() ODRecordClass {
	return getODRecordClass()
}

type ODRecordClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (oc ODRecordClass) Class() objc.Class {
	return oc.class
}

// Alloc allocates memory for a new instance of the class.
func (oc ODRecordClass) Alloc() ODRecord {
	rv := objc.Send[ODRecord](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// An [ODRecord] object serves as a Cocoa wrapper for an Open Directory
// record.
//
// # Managing Authentication
//
//   - [ODRecord.ChangePasswordToPasswordError]: Changes the record’s password.
//   - [ODRecord.SetNodeCredentialsPasswordError]: Sets credentials for the record’s node.
//   - [ODRecord.VerifyPasswordError]: Verifies the password for interaction with the record.
//
// # Managing Group Records
//
//   - [ODRecord.AddMemberRecordError]: Adds a member record to this group record.
//   - [ODRecord.IsMemberRecordError]: Determines whether a given record is a member of this group record.
//   - [ODRecord.RemoveMemberRecordError]: Removes a record as a member of this group record.
//
// # Managing Record Attributes
//
//   - [ODRecord.AddValueToAttributeError]: Adds a value to an attribute of the record.
//   - [ODRecord.RecordDetailsForAttributesError]: Returns a dictionary of attributes with their respective values.
//   - [ODRecord.RecordName]: The official name of the record.
//   - [ODRecord.RecordType]: The record’s type.
//   - [ODRecord.RemoveValuesForAttributeError]: Removes all values from an attribute of the record.
//   - [ODRecord.RemoveValueFromAttributeError]: Removes a value from an attribute of the record.
//   - [ODRecord.SetValueForAttributeError]: Sets the values of an attribute of the record.
//   - [ODRecord.SynchronizeAndReturnError]: Synchronizes the record from the directory to get current data and commit changes.
//   - [ODRecord.ValuesForAttributeError]: Returns the values of an attribute of the record.
//
// # Deleting a Record
//
//   - [ODRecord.DeleteRecordAndReturnError]: Deletes the record from its node and invalidates it.
//
// # Instance Properties
//
//   - [ODRecord.SecondsUntilAuthenticationsExpire]
//   - [ODRecord.SecondsUntilPasswordExpires]
//
// # Instance Methods
//
//   - [ODRecord.AccountPoliciesAndReturnError]
//   - [ODRecord.AddAccountPolicyToCategoryError]
//   - [ODRecord.AuthenticationAllowedAndReturnError]
//   - [ODRecord.PasswordChangeAllowedError]
//   - [ODRecord.RemoveAccountPolicyFromCategoryError]
//   - [ODRecord.SetAccountPoliciesError]
//   - [ODRecord.WillAuthenticationsExpire]
//   - [ODRecord.WillPasswordExpire]
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecord
type ODRecord struct {
	objectivec.Object
}

// ODRecordFromID constructs a [ODRecord] from an objc.ID.
//
// An [ODRecord] object serves as a Cocoa wrapper for an Open Directory
// record.
func ODRecordFromID(id objc.ID) ODRecord {
	return ODRecord{objectivec.Object{ID: id}}
}

// NOTE: ODRecord adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [ODRecord] class.
//
// # Managing Authentication
//
//   - [IODRecord.ChangePasswordToPasswordError]: Changes the record’s password.
//   - [IODRecord.SetNodeCredentialsPasswordError]: Sets credentials for the record’s node.
//   - [IODRecord.VerifyPasswordError]: Verifies the password for interaction with the record.
//
// # Managing Group Records
//
//   - [IODRecord.AddMemberRecordError]: Adds a member record to this group record.
//   - [IODRecord.IsMemberRecordError]: Determines whether a given record is a member of this group record.
//   - [IODRecord.RemoveMemberRecordError]: Removes a record as a member of this group record.
//
// # Managing Record Attributes
//
//   - [IODRecord.AddValueToAttributeError]: Adds a value to an attribute of the record.
//   - [IODRecord.RecordDetailsForAttributesError]: Returns a dictionary of attributes with their respective values.
//   - [IODRecord.RecordName]: The official name of the record.
//   - [IODRecord.RecordType]: The record’s type.
//   - [IODRecord.RemoveValuesForAttributeError]: Removes all values from an attribute of the record.
//   - [IODRecord.RemoveValueFromAttributeError]: Removes a value from an attribute of the record.
//   - [IODRecord.SetValueForAttributeError]: Sets the values of an attribute of the record.
//   - [IODRecord.SynchronizeAndReturnError]: Synchronizes the record from the directory to get current data and commit changes.
//   - [IODRecord.ValuesForAttributeError]: Returns the values of an attribute of the record.
//
// # Deleting a Record
//
//   - [IODRecord.DeleteRecordAndReturnError]: Deletes the record from its node and invalidates it.
//
// # Instance Properties
//
//   - [IODRecord.SecondsUntilAuthenticationsExpire]
//   - [IODRecord.SecondsUntilPasswordExpires]
//
// # Instance Methods
//
//   - [IODRecord.AccountPoliciesAndReturnError]
//   - [IODRecord.AddAccountPolicyToCategoryError]
//   - [IODRecord.AuthenticationAllowedAndReturnError]
//   - [IODRecord.PasswordChangeAllowedError]
//   - [IODRecord.RemoveAccountPolicyFromCategoryError]
//   - [IODRecord.SetAccountPoliciesError]
//   - [IODRecord.WillAuthenticationsExpire]
//   - [IODRecord.WillPasswordExpire]
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecord
type IODRecord interface {
	objectivec.IObject

	// Topic: Managing Authentication

	// Changes the record’s password.
	ChangePasswordToPasswordError(oldPassword string, newPassword string) (bool, error)
	// Sets credentials for the record’s node.
	SetNodeCredentialsPasswordError(inUsername string, inPassword string) (bool, error)
	// Verifies the password for interaction with the record.
	VerifyPasswordError(inPassword string) (bool, error)

	// Topic: Managing Group Records

	// Adds a member record to this group record.
	AddMemberRecordError(inRecord IODRecord) (bool, error)
	// Determines whether a given record is a member of this group record.
	IsMemberRecordError(inRecord IODRecord) (bool, error)
	// Removes a record as a member of this group record.
	RemoveMemberRecordError(inRecord IODRecord) (bool, error)

	// Topic: Managing Record Attributes

	// Adds a value to an attribute of the record.
	AddValueToAttributeError(inValue objectivec.IObject, inAttribute unsafe.Pointer) (bool, error)
	// Returns a dictionary of attributes with their respective values.
	RecordDetailsForAttributesError(inAttributes foundation.INSArray) (foundation.INSDictionary, error)
	// The official name of the record.
	RecordName() string
	// The record’s type.
	RecordType() string
	// Removes all values from an attribute of the record.
	RemoveValuesForAttributeError(inAttribute unsafe.Pointer) (bool, error)
	// Removes a value from an attribute of the record.
	RemoveValueFromAttributeError(inValue objectivec.IObject, inAttribute unsafe.Pointer) (bool, error)
	// Sets the values of an attribute of the record.
	SetValueForAttributeError(inValueOrValues objectivec.IObject, inAttribute unsafe.Pointer) (bool, error)
	// Synchronizes the record from the directory to get current data and commit changes.
	SynchronizeAndReturnError() (bool, error)
	// Returns the values of an attribute of the record.
	ValuesForAttributeError(inAttribute unsafe.Pointer) (foundation.INSArray, error)

	// Topic: Deleting a Record

	// Deletes the record from its node and invalidates it.
	DeleteRecordAndReturnError() (bool, error)

	// Topic: Instance Properties

	SecondsUntilAuthenticationsExpire() int64
	SecondsUntilPasswordExpires() int64

	// Topic: Instance Methods

	AccountPoliciesAndReturnError() (foundation.INSDictionary, error)
	AddAccountPolicyToCategoryError(policy foundation.INSDictionary, category ODPolicyCategoryType) (bool, error)
	AuthenticationAllowedAndReturnError() (bool, error)
	PasswordChangeAllowedError(newPassword string) (bool, error)
	RemoveAccountPolicyFromCategoryError(policy foundation.INSDictionary, category ODPolicyCategoryType) (bool, error)
	SetAccountPoliciesError(policies foundation.INSDictionary) (bool, error)
	WillAuthenticationsExpire(willExpireIn uint64) bool
	WillPasswordExpire(willExpireIn uint64) bool
}

// Init initializes the instance.
func (o ODRecord) Init() ODRecord {
	rv := objc.Send[ODRecord](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o ODRecord) Autorelease() ODRecord {
	rv := objc.Send[ODRecord](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewODRecord creates a new ODRecord instance.
func NewODRecord() ODRecord {
	class := getODRecordClass()
	rv := objc.Send[ODRecord](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Changes the record’s password.
//
// oldPassword: The record’s old password. Can be `nil` if the user has the proper
// permissions.
//
// newPassword: The new password.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecord/changePassword(_:toPassword:)
func (o ODRecord) ChangePasswordToPasswordError(oldPassword string, newPassword string) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](o.ID, objc.Sel("changePassword:toPassword:error:"), objc.String(oldPassword), objc.String(newPassword), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("changePassword:toPassword:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Sets credentials for the record’s node.
//
// inUsername: The username to use to authenticate with the node.
//
// inPassword: The password to use to authenticate with the node.
//
// # Discussion
//
// If this function fails, the previous credentials for the node are used.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecord/setNodeCredentials(_:password:)
func (o ODRecord) SetNodeCredentialsPasswordError(inUsername string, inPassword string) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](o.ID, objc.Sel("setNodeCredentials:password:error:"), objc.String(inUsername), objc.String(inPassword), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setNodeCredentials:password:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Verifies the password for interaction with the record.
//
// inPassword: The password for authenticating with the record.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecord/verifyPassword(_:)
func (o ODRecord) VerifyPasswordError(inPassword string) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](o.ID, objc.Sel("verifyPassword:error:"), objc.String(inPassword), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("verifyPassword:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Adds a member record to this group record.
//
// inRecord: The member record to add.
//
// # Discussion
//
// This method produces an error if this record is not a group record, or if
// `inRecord` is not an appropriate type.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecord/addMemberRecord(_:)
func (o ODRecord) AddMemberRecordError(inRecord IODRecord) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](o.ID, objc.Sel("addMemberRecord:error:"), inRecord, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("addMemberRecord:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Determines whether a given record is a member of this group record.
//
// inRecord: The record to test for membership.
//
// # Discussion
//
// If this record is not a group record, this method returns false.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecord/isMemberRecord(_:)
func (o ODRecord) IsMemberRecordError(inRecord IODRecord) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](o.ID, objc.Sel("isMemberRecord:error:"), inRecord, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("isMemberRecord:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Removes a record as a member of this group record.
//
// inRecord: The member record.
//
// # Discussion
//
// This method produces an error if this record is not a group record, or if
// `inRecord` is not an appropriate type.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecord/removeMemberRecord(_:)
func (o ODRecord) RemoveMemberRecordError(inRecord IODRecord) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](o.ID, objc.Sel("removeMemberRecord:error:"), inRecord, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("removeMemberRecord:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Adds a value to an attribute of the record.
//
// inValue: The value. Should be of type [NSString] or [NSData].
//
// inAttribute: The attribute.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecord/addValue(_:toAttribute:)
func (o ODRecord) AddValueToAttributeError(inValue objectivec.IObject, inAttribute unsafe.Pointer) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](o.ID, objc.Sel("addValue:toAttribute:error:"), inValue, inAttribute, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("addValue:toAttribute:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Returns a dictionary of attributes with their respective values.
//
// inAttributes: An array of attributes. Can be `nil`.
//
// # Return Value
//
// A dictionary of the attributes in `inAttributes` with their respective
// values.
//
// # Discussion
//
// If `inAttributes` is `nil`, all currently retrieved attributes are
// returned.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecord/recordDetails(forAttributes:)
func (o ODRecord) RecordDetailsForAttributesError(inAttributes foundation.INSArray) (foundation.INSDictionary, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](o.ID, objc.Sel("recordDetailsForAttributes:error:"), inAttributes, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSDictionaryFromID(rv), nil

}

// Removes all values from an attribute of the record.
//
// inAttribute: The attribute.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecord/removeValues(forAttribute:)
func (o ODRecord) RemoveValuesForAttributeError(inAttribute unsafe.Pointer) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](o.ID, objc.Sel("removeValuesForAttribute:error:"), inAttribute, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("removeValuesForAttribute:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Removes a value from an attribute of the record.
//
// inValue: The value. Should be of type [NSString] or [NSData].
//
// inAttribute: The attribute.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecord/removeValue(_:fromAttribute:)
func (o ODRecord) RemoveValueFromAttributeError(inValue objectivec.IObject, inAttribute unsafe.Pointer) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](o.ID, objc.Sel("removeValue:fromAttribute:error:"), inValue, inAttribute, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("removeValue:fromAttribute:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Sets the values of an attribute of the record.
//
// inValueOrValues: The value or values. Can be of type [NSString] or [NSData], or an [NSArray]
// with elements of both types.
//
// inAttribute: The attribute.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecord/setValue(_:forAttribute:)
func (o ODRecord) SetValueForAttributeError(inValueOrValues objectivec.IObject, inAttribute unsafe.Pointer) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](o.ID, objc.Sel("setValue:forAttribute:error:"), inValueOrValues, inAttribute, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setValue:forAttribute:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Synchronizes the record from the directory to get current data and commit
// changes.
//
// # Discussion
//
// This method only fetches those attributes that have been fetched before.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecord/synchronize()
func (o ODRecord) SynchronizeAndReturnError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](o.ID, objc.Sel("synchronizeAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("synchronizeAndReturnError: returned NO with nil NSError")
	}
	return rv, nil

}

// Returns the values of an attribute of the record.
//
// inAttribute: The attribute.
//
// # Return Value
//
// An array of attribute values. Elements are of type [NSString] or [NSData].
//
// # Discussion
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecord/values(forAttribute:)
func (o ODRecord) ValuesForAttributeError(inAttribute unsafe.Pointer) (foundation.INSArray, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](o.ID, objc.Sel("valuesForAttribute:error:"), inAttribute, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSArrayFromID(rv), nil

}

// Deletes the record from its node and invalidates it.
//
// # Discussion
//
// The record should be released after this method is called.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecord/delete()
func (o ODRecord) DeleteRecordAndReturnError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](o.ID, objc.Sel("deleteRecordAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("deleteRecordAndReturnError: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/OpenDirectory/ODRecord/accountPolicies()
func (o ODRecord) AccountPoliciesAndReturnError() (foundation.INSDictionary, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accountPoliciesAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSDictionaryFromID(rv), nil

}

// See: https://developer.apple.com/documentation/OpenDirectory/ODRecord/addAccountPolicy(_:toCategory:)
func (o ODRecord) AddAccountPolicyToCategoryError(policy foundation.INSDictionary, category ODPolicyCategoryType) (bool, error) {
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

// See: https://developer.apple.com/documentation/OpenDirectory/ODRecord/authenticationAllowed()
func (o ODRecord) AuthenticationAllowedAndReturnError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](o.ID, objc.Sel("authenticationAllowedAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("authenticationAllowedAndReturnError: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/OpenDirectory/ODRecord/passwordChangeAllowed(_:)
func (o ODRecord) PasswordChangeAllowedError(newPassword string) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](o.ID, objc.Sel("passwordChangeAllowed:error:"), objc.String(newPassword), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("passwordChangeAllowed:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/OpenDirectory/ODRecord/removeAccountPolicy(_:fromCategory:)
func (o ODRecord) RemoveAccountPolicyFromCategoryError(policy foundation.INSDictionary, category ODPolicyCategoryType) (bool, error) {
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

// See: https://developer.apple.com/documentation/OpenDirectory/ODRecord/setAccountPolicies(_:)
func (o ODRecord) SetAccountPoliciesError(policies foundation.INSDictionary) (bool, error) {
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

// See: https://developer.apple.com/documentation/OpenDirectory/ODRecord/willAuthenticationsExpire(_:)
func (o ODRecord) WillAuthenticationsExpire(willExpireIn uint64) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("willAuthenticationsExpire:"), willExpireIn)
	return rv
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODRecord/willPasswordExpire(_:)
func (o ODRecord) WillPasswordExpire(willExpireIn uint64) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("willPasswordExpire:"), willExpireIn)
	return rv
}

// The official name of the record.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecord/recordName
func (o ODRecord) RecordName() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("recordName"))
	return foundation.NSStringFromID(rv).String()
}

// The record’s type.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecord/recordType
func (o ODRecord) RecordType() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("recordType"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODRecord/secondsUntilAuthenticationsExpire
func (o ODRecord) SecondsUntilAuthenticationsExpire() int64 {
	rv := objc.Send[int64](o.ID, objc.Sel("secondsUntilAuthenticationsExpire"))
	return rv
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODRecord/secondsUntilPasswordExpires
func (o ODRecord) SecondsUntilPasswordExpires() int64 {
	rv := objc.Send[int64](o.ID, objc.Sel("secondsUntilPasswordExpires"))
	return rv
}
