// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKContainer] class.
var (
	_CKContainerClass     CKContainerClass
	_CKContainerClassOnce sync.Once
)

func getCKContainerClass() CKContainerClass {
	_CKContainerClassOnce.Do(func() {
		_CKContainerClass = CKContainerClass{class: objc.GetClass("CKContainer")}
	})
	return _CKContainerClass
}

// GetCKContainerClass returns the class object for CKContainer.
func GetCKContainerClass() CKContainerClass {
	return getCKContainerClass()
}

type CKContainerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKContainerClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKContainerClass) Alloc() CKContainer {
	rv := objc.Send[CKContainer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A conduit to your app’s databases.
//
// # Overview
//
// A container manages all explicit and implicit attempts to access its
// contents.
//
// Every app has a default container that manages its own content. If you
// develop a suite of apps, you can access any containers that you have the
// appropriate entitlements for. Each new container distinguishes between
// public and private data. CloudKit always stores private data in the
// appropriate container directory in the user’s iCloud account.
//
// # Interacting with a Container
//
// A container coordinates all interactions between your app and the server.
// Most of these interactions involve the following tasks:
//
// - Determining whether the user has an iCloud account, which lets you know
// if you can write data to the user’s personal storage. - With the user’s
// permission, discovering other users who the current user knows, and making
// the current user’s information discoverable. - Getting the container or
// one of its databases to use with an operation.
//
// # Public and Private Databases
//
// Each container provides a public and a private database for storing data.
// The contents of the public database are accessible to all users of the app,
// whereas the contents of the private database are, by default, visible only
// to the current user. Content that is specific to a single user usually
// belongs in that user’s private database, whereas app-related content that
// you provide (or that users want to share) belongs in the public database.
//
// The public database is always available, regardless of whether the device
// has an active iCloud account. When there isn’t an iCloud account, your
// app can fetch records from and query the public database, but it can’t
// save changes. Saving records to the public database requires an active
// iCloud account to identify the owner of those records. Access to the
// private database always requires an active iCloud account on the device.
//
// # Using iCloud
//
// Whenever possible, design your app to run gracefully with or without an
// active iCloud account. Even without an active iCloud account, apps can
// fetch records from the public database and display that information to the
// user. If your app requires the ability to write to the public database or
// requires access to the private database, notify the user of the reason and
// encourage them to enable iCloud. You can even provide a button that takes
// the user directly to Settings so that they can enable iCloud. To implement
// such a button, have the button’s action open the URL that the
// [CKContainer.OpenSettingsURLString] constant provides.
//
// # User Records and Permissions
//
// When a user accesses a container for the first time, CloudKit assigns them
// a unique identifier and uses it to create two user records — one in the
// app’s public database and another in that user’s private database. By
// default, these records don’t contain any identifying personal
// information, but you can use the record in the user’s private database to
// store additional, nonsensitive information about that user. Because the
// public database’s user record is accessible to all users of your app,
// don’t use it to store information about the user.
//
// While a user record isn’t the same as the user’s [CKUserIdentity], the
// identity does provide the identifier of their user record that you can use
// to fetch that record from either the public database or the user’s
// private database. For more information, see [CKContainer.UserRecordID].
//
// # Testing Your Code Using the Development Container
//
// At runtime, CloudKit uses your app’s
// `com.AppleXCUIElementTypeDeveloperXCUIElementTypeIcloud()-container-environment`
// entitlement to discover whether you’re using a [Development] or
// [Production] version of your provisioning profile. When you configure the
// entitlement for development, CloudKit configures the app’s containers to
// use the development server. The development environment is a safe place to
// make changes during the development process without disrupting users of
// your app. You can add new fields to records programmatically, and you can
// delete or modify fields using iCloud Dashboard.
//
// Before shipping your app, always test your app’s behavior in the
// production environment. The production server generates errors when your
// app tries to add record types or add new fields to existing record types.
// Testing in the production environment helps you find and fix the places in
// your code where you’re making those types of changes. You can use
// CloudKit Dashboard to modify record types in the development environment,
// and then migrate those changes to the production environment.
//
// # Getting the Public and Private Databases
//
//   - [CKContainer.PrivateCloudDatabase]: The user’s private database.
//   - [CKContainer.PublicCloudDatabase]: The app’s public database.
//   - [CKContainer.SharedCloudDatabase]: The database that contains shared data.
//   - [CKContainer.DatabaseWithDatabaseScope]: Returns the database with the specified scope.
//
// # Getting the Container’s Identifier
//
//   - [CKContainer.ContainerIdentifier]: The container’s unique identifier.
//
// # Determining the User’s iCloud Access Status
//
//   - [CKContainer.AccountStatusWithCompletionHandler]: Determines whether the system can access the user’s iCloud account.
//
// # Performing Operations on the Container
//
//   - [CKContainer.AddOperation]: Adds an operation to the container’s queue.
//
// # Discovering User Records
//
//   - [CKContainer.FetchShareParticipantWithEmailAddressCompletionHandler]: Fetches the share participant with the specified email address.
//   - [CKContainer.FetchShareParticipantWithPhoneNumberCompletionHandler]: Fetches the share participant with the specified phone number.
//   - [CKContainer.FetchShareParticipantWithUserRecordIDCompletionHandler]: Fetches the share participant with the specified user record ID.
//   - [CKContainer.FetchUserRecordIDWithCompletionHandler]: Fetches the user record ID of the current user.
//   - [CKContainer.CKCurrentUserDefaultName]: A constant that provides the current user’s default name.
//   - [CKContainer.CKOwnerDefaultName]: A constant that provides the default owner’s name.
//
// # Accessing Container Metadata
//
//   - [CKContainer.FetchShareMetadataWithURLCompletionHandler]: Fetches the share metadata for the specified share URL.
//   - [CKContainer.AcceptShareMetadataCompletionHandler]: Accepts the specified share metadata.
//   - [CKContainer.CKAccountChanged]: A notification that a container posts when the status of an iCloud account changes.
//
// See: https://developer.apple.com/documentation/CloudKit/CKContainer
type CKContainer struct {
	objectivec.Object
}

// CKContainerFromID constructs a [CKContainer] from an objc.ID.
//
// A conduit to your app’s databases.
func CKContainerFromID(id objc.ID) CKContainer {
	return CKContainer{objectivec.Object{ID: id}}
}

// NOTE: CKContainer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKContainer] class.
//
// # Getting the Public and Private Databases
//
//   - [ICKContainer.PrivateCloudDatabase]: The user’s private database.
//   - [ICKContainer.PublicCloudDatabase]: The app’s public database.
//   - [ICKContainer.SharedCloudDatabase]: The database that contains shared data.
//   - [ICKContainer.DatabaseWithDatabaseScope]: Returns the database with the specified scope.
//
// # Getting the Container’s Identifier
//
//   - [ICKContainer.ContainerIdentifier]: The container’s unique identifier.
//
// # Determining the User’s iCloud Access Status
//
//   - [ICKContainer.AccountStatusWithCompletionHandler]: Determines whether the system can access the user’s iCloud account.
//
// # Performing Operations on the Container
//
//   - [ICKContainer.AddOperation]: Adds an operation to the container’s queue.
//
// # Discovering User Records
//
//   - [ICKContainer.FetchShareParticipantWithEmailAddressCompletionHandler]: Fetches the share participant with the specified email address.
//   - [ICKContainer.FetchShareParticipantWithPhoneNumberCompletionHandler]: Fetches the share participant with the specified phone number.
//   - [ICKContainer.FetchShareParticipantWithUserRecordIDCompletionHandler]: Fetches the share participant with the specified user record ID.
//   - [ICKContainer.FetchUserRecordIDWithCompletionHandler]: Fetches the user record ID of the current user.
//   - [ICKContainer.CKCurrentUserDefaultName]: A constant that provides the current user’s default name.
//   - [ICKContainer.CKOwnerDefaultName]: A constant that provides the default owner’s name.
//
// # Accessing Container Metadata
//
//   - [ICKContainer.FetchShareMetadataWithURLCompletionHandler]: Fetches the share metadata for the specified share URL.
//   - [ICKContainer.AcceptShareMetadataCompletionHandler]: Accepts the specified share metadata.
//   - [ICKContainer.CKAccountChanged]: A notification that a container posts when the status of an iCloud account changes.
//
// See: https://developer.apple.com/documentation/CloudKit/CKContainer
type ICKContainer interface {
	objectivec.IObject

	// Topic: Getting the Public and Private Databases

	// The user’s private database.
	PrivateCloudDatabase() ICKDatabase
	// The app’s public database.
	PublicCloudDatabase() ICKDatabase
	// The database that contains shared data.
	SharedCloudDatabase() ICKDatabase
	// Returns the database with the specified scope.
	DatabaseWithDatabaseScope(databaseScope CKDatabaseScope) ICKDatabase

	// Topic: Getting the Container’s Identifier

	// The container’s unique identifier.
	ContainerIdentifier() string

	// Topic: Determining the User’s iCloud Access Status

	// Determines whether the system can access the user’s iCloud account.
	AccountStatusWithCompletionHandler(completionHandler CKAccountStatusErrorHandler)

	// Topic: Performing Operations on the Container

	// Adds an operation to the container’s queue.
	AddOperation(operation ICKOperation)

	// Topic: Discovering User Records

	// Fetches the share participant with the specified email address.
	FetchShareParticipantWithEmailAddressCompletionHandler(emailAddress string, completionHandler CKShareParticipantErrorHandler)
	// Fetches the share participant with the specified phone number.
	FetchShareParticipantWithPhoneNumberCompletionHandler(phoneNumber string, completionHandler CKShareParticipantErrorHandler)
	// Fetches the share participant with the specified user record ID.
	FetchShareParticipantWithUserRecordIDCompletionHandler(userRecordID ICKRecordID, completionHandler CKShareParticipantErrorHandler)
	// Fetches the user record ID of the current user.
	FetchUserRecordIDWithCompletionHandler(completionHandler CKRecordIDErrorHandler)
	// A constant that provides the current user’s default name.
	CKCurrentUserDefaultName() string
	// A constant that provides the default owner’s name.
	CKOwnerDefaultName() string

	// Topic: Accessing Container Metadata

	// Fetches the share metadata for the specified share URL.
	FetchShareMetadataWithURLCompletionHandler(url foundation.INSURL, completionHandler CKShareMetadataErrorHandler)
	// Accepts the specified share metadata.
	AcceptShareMetadataCompletionHandler(metadata ICKShareMetadata, completionHandler CKShareErrorHandler)
	// A notification that a container posts when the status of an iCloud account changes.
	CKAccountChanged() foundation.NSString

	// The user record ID for the corresponding user record.
	UserRecordID() ICKRecordID
	SetUserRecordID(value ICKRecordID)
}

// Init initializes the instance.
func (c CKContainer) Init() CKContainer {
	rv := objc.Send[CKContainer](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKContainer) Autorelease() CKContainer {
	rv := objc.Send[CKContainer](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKContainer creates a new CKContainer instance.
func NewCKContainer() CKContainer {
	class := getCKContainerClass()
	rv := objc.Send[CKContainer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a container for the specified identifier.
//
// containerIdentifier: The bundle identifier of the app with the container that you want to
// access. The bundle identifier must be in the app’s
// `com.AppleXCUIElementTypeDeveloperXCUIElementTypeIcloud()-container-identifiers`
// entitlement. This parameter must not be `nil`.
//
// # Discussion
//
// The specified identifier must correspond to one of the ubiquity containers
// in the iCloud capabilities section of your Xcode project. Including the
// identifier with your app’s capabilities adds the corresponding
// entitlements to your app. To access your app’s default container, use the
// [DefaultContainer] method instead.
//
// See: https://developer.apple.com/documentation/CloudKit/CKContainer/init(identifier:)
func NewCKContainerWithIdentifier(containerIdentifier string) CKContainer {
	rv := objc.Send[objc.ID](objc.ID(getCKContainerClass().class), objc.Sel("containerWithIdentifier:"), objc.String(containerIdentifier))
	return CKContainerFromID(rv)
}

// Returns the database with the specified scope.
//
// databaseScope: The database’s scope. See [CKDatabase.Scope] for the available options.
//
// See: https://developer.apple.com/documentation/CloudKit/CKContainer/database(with:)
//
// [CKDatabase.Scope]: https://developer.apple.com/documentation/CloudKit/CKDatabase/Scope
func (c CKContainer) DatabaseWithDatabaseScope(databaseScope CKDatabaseScope) ICKDatabase {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("databaseWithDatabaseScope:"), databaseScope)
	return CKDatabaseFromID(rv)
}

// Determines whether the system can access the user’s iCloud account.
//
// completionHandler: The handler to execute when the call completes.
//
// # Discussion
//
// The closure has no return value and takes the following parameters:
//
// - The status of the user’s iCloud account. - An error that describes the
// failure, or `nil` if the system successfully determines the status.
//
// This method determines the status of the user’s iCloud account
// asynchronously, passing the results to the closure that you provide. Call
// this method before accessing the private database to determine whether that
// database is available. While your app is running, use the
// [CKAccountChanged] notification to detect account changes, and call this
// method again to determine the status of the new account.
//
// See: https://developer.apple.com/documentation/CloudKit/CKContainer/accountStatus(completionHandler:)
func (c CKContainer) AccountStatusWithCompletionHandler(completionHandler CKAccountStatusErrorHandler) {
	_block0, _ := NewCKAccountStatusErrorBlock(completionHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("accountStatusWithCompletionHandler:"), _block0)
}

// Adds an operation to the container’s queue.
//
// operation: The operation to add to the queue. Make sure you fully configure the
// operation and have it ready to execute. Don’t change the operation’s
// configuration after you queue it.
//
// # Discussion
//
// This method adds the operation to a queue that the container manages. The
// queue’s operations execute on background threads concurrently, and with
// default priorities. When you add an operation to the queue, its container
// becomes the current container.
//
// See: https://developer.apple.com/documentation/CloudKit/CKContainer/add(_:)
func (c CKContainer) AddOperation(operation ICKOperation) {
	objc.Send[objc.ID](c.ID, objc.Sel("addOperation:"), operation)
}

// Fetches the share participant with the specified email address.
//
// emailAddress: The share participant’s email address.
//
// completionHandler: The handler to execute with the fetch results.
//
// # Discussion
//
// The closure doesn’t return a value and takes the following parameters:
//
// - The share participant, or `nil` if CloudKit can’t find the participant.
// - An error if a problem occurs, or `nil` if CloudKit successfully retrieves
// the participant.
//
// CloudKit can translate any valid email address into a share participant. If
// the email address doesn’t correspond to a known iCloud account, then at
// share-accept-time, CloudKit offers the accepting participant a vetting
// process. The accepting participant uses this vetting process to link the
// email address to an iCloud account.
//
// This method searches for the share participant asynchronously and with a
// low priority. If you want the task to execute with a higher priority,
// create an instance of [CKFetchShareParticipantsOperation] and configure it
// to use the necessary priority.
//
// See: https://developer.apple.com/documentation/CloudKit/CKContainer/fetchShareParticipant(withEmailAddress:completionHandler:)
func (c CKContainer) FetchShareParticipantWithEmailAddressCompletionHandler(emailAddress string, completionHandler CKShareParticipantErrorHandler) {
	_block1, _ := NewCKShareParticipantErrorBlock(completionHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("fetchShareParticipantWithEmailAddress:completionHandler:"), objc.String(emailAddress), _block1)
}

// Fetches the share participant with the specified phone number.
//
// phoneNumber: The share participant’s phone number.
//
// completionHandler: The handler to execute with the fetch results.
//
// # Discussion
//
// The closure doesn’t return a value and takes the following parameters:
//
// - The share participant, or `nil` if CloudKit can’t find the participant.
// - An error if a problem occurs, or `nil` if CloudKit successfully retrieves
// the participant.
//
// CloudKit can translate any valid phone number into a share participant. If
// the phone number doesn’t correspond to a known iCloud account, then at
// share-accept-time, CloudKit offers the accepting participant a vetting
// process. The accepting participant uses this vetting process to link the
// phone number to an iCloud account.
//
// This method searches for the share participant asynchronously and with a
// low priority. If you want the task to execute with a higher priority,
// create an instance of [CKFetchShareParticipantsOperation] and configure it
// to use the necessary priority.
//
// See: https://developer.apple.com/documentation/CloudKit/CKContainer/fetchShareParticipant(withPhoneNumber:completionHandler:)
func (c CKContainer) FetchShareParticipantWithPhoneNumberCompletionHandler(phoneNumber string, completionHandler CKShareParticipantErrorHandler) {
	_block1, _ := NewCKShareParticipantErrorBlock(completionHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("fetchShareParticipantWithPhoneNumber:completionHandler:"), objc.String(phoneNumber), _block1)
}

// Fetches the share participant with the specified user record ID.
//
// userRecordID: The share participant’s user record ID.
//
// completionHandler: The handler to execute with the fetch results.
//
// # Discussion
//
// The closure doesn’t return a value and takes the following parameters:
//
// - The share participant, or `nil` if CloudKit can’t find the participant.
// - An error if a problem occurs, or `nil` if CloudKit successfully retrieves
// the participant.
//
// This method searches for the share participant asynchronously and with a
// low priority. If you want the task to execute with a higher priority,
// create an instance of [CKFetchShareParticipantsOperation] and configure it
// to use the necessary priority.
//
// See: https://developer.apple.com/documentation/CloudKit/CKContainer/fetchShareParticipant(withUserRecordID:completionHandler:)
func (c CKContainer) FetchShareParticipantWithUserRecordIDCompletionHandler(userRecordID ICKRecordID, completionHandler CKShareParticipantErrorHandler) {
	_block1, _ := NewCKShareParticipantErrorBlock(completionHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("fetchShareParticipantWithUserRecordID:completionHandler:"), userRecordID, _block1)
}

// Fetches the user record ID of the current user.
//
// completionHandler: The handler to execute with the fetch results.
//
// # Discussion
//
// The closure doesn’t return a value and takes the following parameters:
//
// - The user record ID, or `nil` if the user disables iCloud or the device
// doesn’t have an iCloud account. - An error if a problem occurs, or `nil`
// if CloudKit successfully retrieves the user record ID.
//
// CloudKit returns a [CKError.Code.notAuthenticated] error when any of the
// following conditions are met:
//
// - The device has an iCloud account but the user disables iCloud. - The
// device has an iCloud account with restricted access. - The device doesn’t
// have an iCloud account.
//
// See: https://developer.apple.com/documentation/CloudKit/CKContainer/fetchUserRecordID(completionHandler:)
//
// [CKError.Code.notAuthenticated]: https://developer.apple.com/documentation/CloudKit/CKError/Code/notAuthenticated
func (c CKContainer) FetchUserRecordIDWithCompletionHandler(completionHandler CKRecordIDErrorHandler) {
	_block0, _ := NewCKRecordIDErrorBlock(completionHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("fetchUserRecordIDWithCompletionHandler:"), _block0)
}

// Fetches the share metadata for the specified share URL.
//
// url: The share URL that CloudKit uses to locate the metadata.
//
// completionHandler: The handler to execute with the fetch results.
//
// # Discussion
//
// The closure doesn’t return a value and takes the following parameters:
//
// - The share metadata, or `nil` if CloudKit can’t find the metadata. - An
// error if a problem occurs, or `nil` if CloudKit successfully retrieves the
// metadata.
//
// See: https://developer.apple.com/documentation/CloudKit/CKContainer/fetchShareMetadata(with:completionHandler:)
func (c CKContainer) FetchShareMetadataWithURLCompletionHandler(url foundation.INSURL, completionHandler CKShareMetadataErrorHandler) {
	_block1, _ := NewCKShareMetadataErrorBlock(completionHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("fetchShareMetadataWithURL:completionHandler:"), url, _block1)
}

// Accepts the specified share metadata.
//
// metadata: The metadata of the share to accept.
//
// completionHandler: The handler to execute when the process finishes.
//
// # Discussion
//
// The closure doesn’t return a value and takes the following parameters:
//
// - The corresponding share, or `nil` if CloudKit can’t accept the
// metadata. - An error if a problem occurs, or `nil` if CloudKit successfully
// accepts the metadata.
//
// See: https://developer.apple.com/documentation/CloudKit/CKContainer/accept(_:completionHandler:)-949ea
func (c CKContainer) AcceptShareMetadataCompletionHandler(metadata ICKShareMetadata, completionHandler CKShareErrorHandler) {
	_block1, _ := NewCKShareErrorBlock(completionHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("acceptShareMetadata:completionHandler:"), metadata, _block1)
}

// Returns the app’s default container.
//
// # Discussion
//
// Use this method to retrieve your app’s default container. This is the one
// you typically use to store your app’s data. If you want the container for
// a different app, create a container using the [ContainerWithIdentifier]
// method.
//
// During development, the container uses the development environment. When
// you release your app, the container uses the production environment.
//
// See: https://developer.apple.com/documentation/CloudKit/CKContainer/default()
func (_CKContainerClass CKContainerClass) DefaultContainer() CKContainer {
	rv := objc.Send[objc.ID](objc.ID(_CKContainerClass.class), objc.Sel("defaultContainer"))
	return CKContainerFromID(rv)
}

// The user’s private database.
//
// # Discussion
//
// The user’s private database is only available if the device has an iCloud
// account. Only the user can access their private database, by default. They
// own all of the database’s content and can view and modify that content.
// Data in the private database isn’t visible in the developer portal.
//
// Data in the private database counts toward the user’s iCloud storage
// quota.
//
// If there isn’t an iCloud account on the user’s device, this property
// still returns a database, but any attempt to use it results in an error. To
// determine if there is an iCloud account on the device, use the
// [AccountStatusWithCompletionHandler] method.
//
// See: https://developer.apple.com/documentation/CloudKit/CKContainer/privateCloudDatabase
func (c CKContainer) PrivateCloudDatabase() ICKDatabase {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("privateCloudDatabase"))
	return CKDatabaseFromID(objc.ID(rv))
}

// The app’s public database.
//
// # Discussion
//
// This database is available regardless of whether the user’s device has an
// iCloud account. The contents of the public database are readable by all
// users of the app, and users have write access to the records, and other
// objects, they create. The public database’s contents are visible in the
// developer portal, where you can assign roles to users and restrict access
// as necessary.
//
// Data in the public database counts toward your app’s iCloud storage
// quota.
//
// See: https://developer.apple.com/documentation/CloudKit/CKContainer/publicCloudDatabase
func (c CKContainer) PublicCloudDatabase() ICKDatabase {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("publicCloudDatabase"))
	return CKDatabaseFromID(objc.ID(rv))
}

// The database that contains shared data.
//
// # Discussion
//
// This database is only available if the device has an iCloud account.
// Permissions on the database are available only to the user according to the
// permissions of the enclosing [CKShare] instance, which represents the
// shared record. The current user doesn’t own the content in the shared
// database, and can view and modify that content only if the necessary
// permissions exist. Data in the shared database isn’t visible in the
// developer portal or to any user who doesn’t have access.
//
// Data in the shared database counts toward your app’s iCloud storage
// quota.
//
// If there isn’t an iCloud account on the user’s device, this property
// still returns a database, but any attempt to use it results in an error. To
// determine if there is an iCloud account on the device, use the
// [AccountStatusWithCompletionHandler] method.
//
// See: https://developer.apple.com/documentation/CloudKit/CKContainer/sharedCloudDatabase
func (c CKContainer) SharedCloudDatabase() ICKDatabase {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("sharedCloudDatabase"))
	return CKDatabaseFromID(objc.ID(rv))
}

// The container’s unique identifier.
//
// # Discussion
//
// Use this property’s value to distinguish different containers in your
// app.
//
// See: https://developer.apple.com/documentation/CloudKit/CKContainer/containerIdentifier
func (c CKContainer) ContainerIdentifier() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("containerIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

// A constant that provides the current user’s default name.
//
// See: https://developer.apple.com/documentation/cloudkit/ckcurrentuserdefaultname
func (c CKContainer) CKCurrentUserDefaultName() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("CKCurrentUserDefaultName"))
	return foundation.NSStringFromID(rv).String()
}

// A constant that provides the default owner’s name.
//
// See: https://developer.apple.com/documentation/cloudkit/ckownerdefaultname
func (c CKContainer) CKOwnerDefaultName() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("CKOwnerDefaultName"))
	return foundation.NSStringFromID(rv).String()
}

// A notification that a container posts when the status of an iCloud account
// changes.
//
// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/CKAccountChanged
func (c CKContainer) CKAccountChanged() foundation.NSString {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("CKAccountChanged"))
	return foundation.NSStringFromID(objc.ID(rv))
}

// The user record ID for the corresponding user record.
//
// See: https://developer.apple.com/documentation/cloudkit/ckuseridentity/userrecordid
func (c CKContainer) UserRecordID() ICKRecordID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("userRecordID"))
	return CKRecordIDFromID(objc.ID(rv))
}
func (c CKContainer) SetUserRecordID(value ICKRecordID) {
	objc.Send[struct{}](c.ID, objc.Sel("setUserRecordID:"), value)
}

// The URL string you use to deep link to your app’s custom settings in the
// Settings app.
//
// See: https://developer.apple.com/documentation/UIKit/UIApplication/openSettingsURLString
func (_CKContainerClass CKContainerClass) OpenSettingsURLString() string {
	rv := objc.Send[objc.ID](objc.ID(_CKContainerClass.class), objc.Sel("openSettingsURLString"))
	return foundation.NSStringFromID(rv).String()
}

// AccountStatus is a synchronous wrapper around [CKContainer.AccountStatusWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c CKContainer) AccountStatus(ctx context.Context) (CKAccountStatus, error) {
	type result struct {
		val CKAccountStatus
		err error
	}
	done := make(chan result, 1)
	c.AccountStatusWithCompletionHandler(func(val CKAccountStatus, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return *new(CKAccountStatus), ctx.Err()
	}
}

// FetchShareParticipantWithEmailAddress is a synchronous wrapper around [CKContainer.FetchShareParticipantWithEmailAddressCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c CKContainer) FetchShareParticipantWithEmailAddress(ctx context.Context, emailAddress string) (*CKShareParticipant, error) {
	type result struct {
		val *CKShareParticipant
		err error
	}
	done := make(chan result, 1)
	c.FetchShareParticipantWithEmailAddressCompletionHandler(emailAddress, func(val *CKShareParticipant, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// FetchShareParticipantWithPhoneNumber is a synchronous wrapper around [CKContainer.FetchShareParticipantWithPhoneNumberCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c CKContainer) FetchShareParticipantWithPhoneNumber(ctx context.Context, phoneNumber string) (*CKShareParticipant, error) {
	type result struct {
		val *CKShareParticipant
		err error
	}
	done := make(chan result, 1)
	c.FetchShareParticipantWithPhoneNumberCompletionHandler(phoneNumber, func(val *CKShareParticipant, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// FetchShareParticipantWithUserRecordID is a synchronous wrapper around [CKContainer.FetchShareParticipantWithUserRecordIDCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c CKContainer) FetchShareParticipantWithUserRecordID(ctx context.Context, userRecordID ICKRecordID) (*CKShareParticipant, error) {
	type result struct {
		val *CKShareParticipant
		err error
	}
	done := make(chan result, 1)
	c.FetchShareParticipantWithUserRecordIDCompletionHandler(userRecordID, func(val *CKShareParticipant, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// FetchUserRecordID is a synchronous wrapper around [CKContainer.FetchUserRecordIDWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c CKContainer) FetchUserRecordID(ctx context.Context) (*CKRecordID, error) {
	type result struct {
		val *CKRecordID
		err error
	}
	done := make(chan result, 1)
	c.FetchUserRecordIDWithCompletionHandler(func(val *CKRecordID, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// FetchShareMetadataWithURL is a synchronous wrapper around [CKContainer.FetchShareMetadataWithURLCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c CKContainer) FetchShareMetadataWithURL(ctx context.Context, url foundation.INSURL) (*CKShareMetadata, error) {
	type result struct {
		val *CKShareMetadata
		err error
	}
	done := make(chan result, 1)
	c.FetchShareMetadataWithURLCompletionHandler(url, func(val *CKShareMetadata, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// AcceptShareMetadata is a synchronous wrapper around [CKContainer.AcceptShareMetadataCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c CKContainer) AcceptShareMetadata(ctx context.Context, metadata ICKShareMetadata) (*CKShare, error) {
	type result struct {
		val *CKShare
		err error
	}
	done := make(chan result, 1)
	c.AcceptShareMetadataCompletionHandler(metadata, func(val *CKShare, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
