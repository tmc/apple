// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CSPerson] class.
var (
	_CSPersonClass     CSPersonClass
	_CSPersonClassOnce sync.Once
)

func getCSPersonClass() CSPersonClass {
	_CSPersonClassOnce.Do(func() {
		_CSPersonClass = CSPersonClass{class: objc.GetClass("CSPerson")}
	})
	return _CSPersonClass
}

// GetCSPersonClass returns the class object for CSPerson.
func GetCSPersonClass() CSPersonClass {
	return getCSPersonClass()
}

type CSPersonClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CSPersonClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CSPersonClass) Alloc() CSPerson {
	rv := objc.Send[CSPerson](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a person in the context of search results.
//
// # Overview
//
// A [CSPerson] object represents a person in the context of search results.
// You can create a [CSPerson] object when you have a display name and a
// contact handle of some kind, such as an email address or phone number.
//
// If you create a [CSPerson] object to represent a specific contact, you can
// use the value of the contact’s identifier property for the person
// object’s [CSPerson.ContactIdentifier] property. Using the same value lets
// you avoid using names or phone numbers to look up the contact that’s
// associated with a person.
//
// # Initializing a person object
//
//   - [CSPerson.InitWithDisplayNameHandlesHandleIdentifier]: Returns a new [CSPerson] object initialized with the specified display name and contact attributes.
//   - [CSPerson.InitWithCoder]
//
// # Accessing person properties
//
//   - [CSPerson.ContactIdentifier]: The identifier for the contact associated with the person.
//   - [CSPerson.SetContactIdentifier]
//   - [CSPerson.DisplayName]: A display name for the person.
//   - [CSPerson.HandleIdentifier]: A key that identifies the type of contact property represented by the person object’s handle.
//   - [CSPerson.Handles]: An array of contact handles related to the person.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSPerson
type CSPerson struct {
	objectivec.Object
}

// CSPersonFromID constructs a [CSPerson] from an objc.ID.
//
// An object that represents a person in the context of search results.
func CSPersonFromID(id objc.ID) CSPerson {
	return CSPerson{objectivec.Object{ID: id}}
}

// NOTE: CSPerson adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CSPerson] class.
//
// # Initializing a person object
//
//   - [ICSPerson.InitWithDisplayNameHandlesHandleIdentifier]: Returns a new [CSPerson] object initialized with the specified display name and contact attributes.
//   - [ICSPerson.InitWithCoder]
//
// # Accessing person properties
//
//   - [ICSPerson.ContactIdentifier]: The identifier for the contact associated with the person.
//   - [ICSPerson.SetContactIdentifier]
//   - [ICSPerson.DisplayName]: A display name for the person.
//   - [ICSPerson.HandleIdentifier]: A key that identifies the type of contact property represented by the person object’s handle.
//   - [ICSPerson.Handles]: An array of contact handles related to the person.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSPerson
type ICSPerson interface {
	objectivec.IObject

	// Topic: Initializing a person object

	// Returns a new [CSPerson] object initialized with the specified display name and contact attributes.
	InitWithDisplayNameHandlesHandleIdentifier(displayName string, handles []string, handleIdentifier string) CSPerson
	InitWithCoder(coder foundation.INSCoder) CSPerson

	// Topic: Accessing person properties

	// The identifier for the contact associated with the person.
	ContactIdentifier() string
	SetContactIdentifier(value string)
	// A display name for the person.
	DisplayName() string
	// A key that identifies the type of contact property represented by the person object’s handle.
	HandleIdentifier() string
	// An array of contact handles related to the person.
	Handles() []string

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c CSPerson) Init() CSPerson {
	rv := objc.Send[CSPerson](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CSPerson) Autorelease() CSPerson {
	rv := objc.Send[CSPerson](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCSPerson creates a new CSPerson instance.
func NewCSPerson() CSPerson {
	class := getCSPersonClass()
	rv := objc.Send[CSPerson](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreSpotlight/CSPerson/init(coder:)
func NewCSPersonWithCoder(coder foundation.INSCoder) CSPerson {
	instance := getCSPersonClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return CSPersonFromID(rv)
}

// Returns a new [CSPerson] object initialized with the specified display name
// and contact attributes.
//
// displayName: The name of the person in a user-displayable string.
//
// handles: An array of contact handles, such as phone number or email address.
//
// handleIdentifier: A property key that specifies a handle type, such as
// [CNContactEmailAddressesKey].
//
// # Return Value
//
// An initialized person object that represents a user’s contact.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSPerson/init(displayName:handles:handleIdentifier:)
//
// [CNContactEmailAddressesKey]: https://developer.apple.com/documentation/Contacts/CNContactEmailAddressesKey
func NewCSPersonWithDisplayNameHandlesHandleIdentifier(displayName string, handles []string, handleIdentifier string) CSPerson {
	instance := getCSPersonClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDisplayName:handles:handleIdentifier:"), objc.String(displayName), objectivec.StringSliceToNSArray(handles), objc.String(handleIdentifier))
	return CSPersonFromID(rv)
}

// Returns a new [CSPerson] object initialized with the specified display name
// and contact attributes.
//
// displayName: The name of the person in a user-displayable string.
//
// handles: An array of contact handles, such as phone number or email address.
//
// handleIdentifier: A property key that specifies a handle type, such as
// [CNContactEmailAddressesKey].
//
// # Return Value
//
// An initialized person object that represents a user’s contact.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSPerson/init(displayName:handles:handleIdentifier:)
//
// [CNContactEmailAddressesKey]: https://developer.apple.com/documentation/Contacts/CNContactEmailAddressesKey
func (c CSPerson) InitWithDisplayNameHandlesHandleIdentifier(displayName string, handles []string, handleIdentifier string) CSPerson {
	rv := objc.Send[CSPerson](c.ID, objc.Sel("initWithDisplayName:handles:handleIdentifier:"), objc.String(displayName), objectivec.StringSliceToNSArray(handles), objc.String(handleIdentifier))
	return rv
}

// See: https://developer.apple.com/documentation/CoreSpotlight/CSPerson/init(coder:)
func (c CSPerson) InitWithCoder(coder foundation.INSCoder) CSPerson {
	rv := objc.Send[CSPerson](c.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (c CSPerson) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The identifier for the contact associated with the person.
//
// # Discussion
//
// When you use the contact’s [identifier] value for the optional
// [CSPerson.ContactIdentifier] property, it enables a direct way to look up
// the associated contact.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSPerson/contactIdentifier
//
// [identifier]: https://developer.apple.com/documentation/Contacts/CNContact/identifier
func (c CSPerson) ContactIdentifier() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("contactIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSPerson) SetContactIdentifier(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setContactIdentifier:"), objc.String(value))
}

// A display name for the person.
//
// # Discussion
//
// Use this optional property to provide a name for the person object that can
// be displayed to users.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSPerson/displayName
func (c CSPerson) DisplayName() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("displayName"))
	return foundation.NSStringFromID(rv).String()
}

// A key that identifies the type of contact property represented by the
// person object’s handle.
//
// # Discussion
//
// The value of this property is a [CNContact] property key, such as
// [CNContactPhoneNumbersKey] or [CNContactEmailAddressesKey].
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSPerson/handleIdentifier
//
// [CNContactEmailAddressesKey]: https://developer.apple.com/documentation/Contacts/CNContactEmailAddressesKey
// [CNContactPhoneNumbersKey]: https://developer.apple.com/documentation/Contacts/CNContactPhoneNumbersKey
// [CNContact]: https://developer.apple.com/documentation/Contacts/CNContact
func (c CSPerson) HandleIdentifier() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("handleIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

// An array of contact handles related to the person.
//
// # Discussion
//
// Contact handles can include phone numbers, email addresses, and URLs. For
// additional contact handles, see Metadata Keys in [CNContact].
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSPerson/handles
//
// [CNContact]: https://developer.apple.com/documentation/Contacts/CNContact
func (c CSPerson) Handles() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("handles"))
	return objc.ConvertSliceToStrings(rv)
}
