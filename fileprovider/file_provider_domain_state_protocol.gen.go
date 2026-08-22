// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An object that contains global state data about the domain.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomainState
type NSFileProviderDomainState interface {
	objectivec.IObject

	// An opaque object that uniquely identifies the domain’s version.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomainState/domainVersion
	DomainVersion() INSFileProviderDomainVersion

	// Global state information about the current domain version.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomainState/userInfo
	UserInfo() foundation.INSDictionary
}

// NSFileProviderDomainStateObject wraps an existing Objective-C object that conforms to the NSFileProviderDomainState protocol.
type NSFileProviderDomainStateObject struct {
	objectivec.Object
}

func (o NSFileProviderDomainStateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSFileProviderDomainStateObjectFromID constructs a [NSFileProviderDomainStateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSFileProviderDomainStateObjectFromID(id objc.ID) NSFileProviderDomainStateObject {
	return NSFileProviderDomainStateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// An opaque object that uniquely identifies the domain’s version.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomainState/domainVersion
func (o NSFileProviderDomainStateObject) DomainVersion() INSFileProviderDomainVersion {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("domainVersion"))
	return NSFileProviderDomainVersionFromID(rv)
}

// Global state information about the current domain version.
//
// # Discussion
//
// Use this dictionary to add state information to the domain. You can then
// access the [UserInfo] dictionary in predicates for user interactions, file
// provider actions, and [File Provider UI] actions using the `domainUserInfo`
// context key.
//
// This dictionary must only contain the following types for both its keys and
// values:
//
// - [NSString]
// - [NSNumber]
// - [NSDate]
// - [NSPersonNameComponents]
//
// The system expects you to update the `domainVersion` whenever the value of
// the [UserInfo] dictionary changes.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomainState/userInfo
//
// [File Provider UI]: https://developer.apple.com/documentation/FileProviderUI
// [NSDate]: https://developer.apple.com/documentation/Foundation/NSDate
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
// [NSPersonNameComponents]: https://developer.apple.com/documentation/Foundation/NSPersonNameComponents
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
func (o NSFileProviderDomainStateObject) UserInfo() foundation.INSDictionary {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("userInfo"))
	return foundation.NSDictionaryFromID(rv)
}
