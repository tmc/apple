// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [URLCredentialStorage] class.
var (
	_URLCredentialStorageClass     URLCredentialStorageClass
	_URLCredentialStorageClassOnce sync.Once
)

func getURLCredentialStorageClass() URLCredentialStorageClass {
	_URLCredentialStorageClassOnce.Do(func() {
		_URLCredentialStorageClass = URLCredentialStorageClass{class: objc.GetClass("NSURLCredentialStorage")}
	})
	return _URLCredentialStorageClass
}

// GetURLCredentialStorageClass returns the class object for NSURLCredentialStorage.
func GetURLCredentialStorageClass() URLCredentialStorageClass {
	return getURLCredentialStorageClass()
}

type URLCredentialStorageClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (uc URLCredentialStorageClass) Alloc() URLCredentialStorage {
	rv := objc.Send[URLCredentialStorage](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// The manager of a shared credentials cache.
//
// # Overview
// 
// The shared cache stores and retrieves instances of [NSURLCredential]. You
// can store password-based credentials permanently, based on the
// [URLCredential.Persistence] they were created with. Certificate-based
// credentials are never stored permanently.
// 
// # Subclassing notes
// 
// The [NSURLCredentialStorage] class is meant to be used as-is, but you can
// subclass it if you have specific needs, such as screening which credentials
// are stored.
// 
// When overriding methods of this class, be aware that methods that take a
// `task` parameter are preferred to equivalent methods that do not.
// Therefore, you should override the task-based methods when subclassing, as
// follows:
// 
// - Setting credentials — Override [SetCredentialForProtectionSpaceTask]
// instead of or in addition to [SetCredentialForProtectionSpace]. - Getting
// credentials — Override
// [GetCredentialsForProtectionSpaceTaskCompletionHandler] instead of or in
// addition to [CredentialsForProtectionSpace]. - Removing credentials —
// Override [RemoveCredentialForProtectionSpaceOptionsTask] instead of or in
// addition to [RemoveCredentialForProtectionSpaceOptions] and
// [RemoveCredentialForProtectionSpace]. - Setting default credentials —
// Override [SetDefaultCredentialForProtectionSpaceTask] instead of or in
// addition to [SetDefaultCredentialForProtectionSpace]. - Getting default
// credentials — Override
// [GetDefaultCredentialForProtectionSpaceTaskCompletionHandler] instead of or
// in addition to [DefaultCredentialForProtectionSpace].
//
// [URLCredential.Persistence]: https://developer.apple.com/documentation/Foundation/URLCredential/Persistence-swift.enum
//
// # Getting and setting default credentials
//
//   - [URLCredentialStorage.DefaultCredentialForProtectionSpace]: Returns the default credential for the specified protection space.
//   - [URLCredentialStorage.GetDefaultCredentialForProtectionSpaceTaskCompletionHandler]: Gets the default credential for the specified protection space, which is being accessed by the given task, and passes it to the provided completion handler.
//   - [URLCredentialStorage.SetDefaultCredentialForProtectionSpace]: Sets the default credential for a specified protection space.
//   - [URLCredentialStorage.SetDefaultCredentialForProtectionSpaceTask]: Sets the default credential for a given protection space, which is being accessed by the given task.
//
// # Adding and removing credentials
//
//   - [URLCredentialStorage.RemoveCredentialForProtectionSpace]: Removes the specified credential from the credential storage for the specified protection space.
//   - [URLCredentialStorage.RemoveCredentialForProtectionSpaceOptions]: Removes the specified credential from the credential storage for the specified protection space using the given options.
//   - [URLCredentialStorage.RemoveCredentialForProtectionSpaceOptionsTask]: Removes the specified credential from the credential storage for the specified protection space, on behalf of the given task and using the given options.
//   - [URLCredentialStorage.SetCredentialForProtectionSpace]: Adds a credential to the credential storage for the specified protection space.
//   - [URLCredentialStorage.SetCredentialForProtectionSpaceTask]: Adds a credential to the credential storage for the specified protection space, on behalf of the specified task.
//
// # Retrieving credentials
//
//   - [URLCredentialStorage.AllCredentials]: The credentials for all available protection spaces.
//   - [URLCredentialStorage.CredentialsForProtectionSpace]: Returns a dictionary containing the credentials for the specified protection space.
//
// # Tracking credential storage changes
//
//   - [URLCredentialStorage.NSURLCredentialStorageChanged]: A notification posted when the set of stored credentials changes.
//
// See: https://developer.apple.com/documentation/Foundation/URLCredentialStorage
type URLCredentialStorage struct {
	objectivec.Object
}

// URLCredentialStorageFromID constructs a [URLCredentialStorage] from an objc.ID.
//
// The manager of a shared credentials cache.
func URLCredentialStorageFromID(id objc.ID) URLCredentialStorage {
	return NSURLCredentialStorage{objectivec.Object{ID: id}}
}

// NSURLCredentialStorageFromID is an alias for [URLCredentialStorageFromID] for cross-framework compatibility.
func NSURLCredentialStorageFromID(id objc.ID) URLCredentialStorage { return URLCredentialStorageFromID(id) }
// NOTE: URLCredentialStorage adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [URLCredentialStorage] class.
//
// # Getting and setting default credentials
//
//   - [IURLCredentialStorage.DefaultCredentialForProtectionSpace]: Returns the default credential for the specified protection space.
//   - [IURLCredentialStorage.GetDefaultCredentialForProtectionSpaceTaskCompletionHandler]: Gets the default credential for the specified protection space, which is being accessed by the given task, and passes it to the provided completion handler.
//   - [IURLCredentialStorage.SetDefaultCredentialForProtectionSpace]: Sets the default credential for a specified protection space.
//   - [IURLCredentialStorage.SetDefaultCredentialForProtectionSpaceTask]: Sets the default credential for a given protection space, which is being accessed by the given task.
//
// # Adding and removing credentials
//
//   - [IURLCredentialStorage.RemoveCredentialForProtectionSpace]: Removes the specified credential from the credential storage for the specified protection space.
//   - [IURLCredentialStorage.RemoveCredentialForProtectionSpaceOptions]: Removes the specified credential from the credential storage for the specified protection space using the given options.
//   - [IURLCredentialStorage.RemoveCredentialForProtectionSpaceOptionsTask]: Removes the specified credential from the credential storage for the specified protection space, on behalf of the given task and using the given options.
//   - [IURLCredentialStorage.SetCredentialForProtectionSpace]: Adds a credential to the credential storage for the specified protection space.
//   - [IURLCredentialStorage.SetCredentialForProtectionSpaceTask]: Adds a credential to the credential storage for the specified protection space, on behalf of the specified task.
//
// # Retrieving credentials
//
//   - [IURLCredentialStorage.AllCredentials]: The credentials for all available protection spaces.
//   - [IURLCredentialStorage.CredentialsForProtectionSpace]: Returns a dictionary containing the credentials for the specified protection space.
//
// # Tracking credential storage changes
//
//   - [IURLCredentialStorage.NSURLCredentialStorageChanged]: A notification posted when the set of stored credentials changes.
//
// See: https://developer.apple.com/documentation/Foundation/URLCredentialStorage
type IURLCredentialStorage interface {
	objectivec.IObject

	// Topic: Getting and setting default credentials

	// Returns the default credential for the specified protection space.
	DefaultCredentialForProtectionSpace(space INSURLProtectionSpace) INSURLCredential
	// Gets the default credential for the specified protection space, which is being accessed by the given task, and passes it to the provided completion handler.
	GetDefaultCredentialForProtectionSpaceTaskCompletionHandler(space INSURLProtectionSpace, task INSURLSessionTask, completionHandler URLCredentialHandler)
	// Sets the default credential for a specified protection space.
	SetDefaultCredentialForProtectionSpace(credential INSURLCredential, space INSURLProtectionSpace)
	// Sets the default credential for a given protection space, which is being accessed by the given task.
	SetDefaultCredentialForProtectionSpaceTask(credential INSURLCredential, protectionSpace INSURLProtectionSpace, task INSURLSessionTask)

	// Topic: Adding and removing credentials

	// Removes the specified credential from the credential storage for the specified protection space.
	RemoveCredentialForProtectionSpace(credential INSURLCredential, space INSURLProtectionSpace)
	// Removes the specified credential from the credential storage for the specified protection space using the given options.
	RemoveCredentialForProtectionSpaceOptions(credential INSURLCredential, space INSURLProtectionSpace, options INSDictionary)
	// Removes the specified credential from the credential storage for the specified protection space, on behalf of the given task and using the given options.
	RemoveCredentialForProtectionSpaceOptionsTask(credential INSURLCredential, protectionSpace INSURLProtectionSpace, options INSDictionary, task INSURLSessionTask)
	// Adds a credential to the credential storage for the specified protection space.
	SetCredentialForProtectionSpace(credential INSURLCredential, space INSURLProtectionSpace)
	// Adds a credential to the credential storage for the specified protection space, on behalf of the specified task.
	SetCredentialForProtectionSpaceTask(credential INSURLCredential, protectionSpace INSURLProtectionSpace, task INSURLSessionTask)

	// Topic: Retrieving credentials

	// The credentials for all available protection spaces.
	AllCredentials() INSDictionary
	// Returns a dictionary containing the credentials for the specified protection space.
	CredentialsForProtectionSpace(space INSURLProtectionSpace) INSDictionary

	// Topic: Tracking credential storage changes

	// A notification posted when the set of stored credentials changes.
	NSURLCredentialStorageChanged() NSNotificationName
}

// Init initializes the instance.
func (u URLCredentialStorage) Init() URLCredentialStorage {
	rv := objc.Send[URLCredentialStorage](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u URLCredentialStorage) Autorelease() URLCredentialStorage {
	rv := objc.Send[URLCredentialStorage](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLCredentialStorage creates a new URLCredentialStorage instance.
func NewURLCredentialStorage() URLCredentialStorage {
	class := getURLCredentialStorageClass()
	rv := objc.Send[URLCredentialStorage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the default credential for the specified protection space.
//
// space: The URL protection space of interest.
//
// # Return Value
// 
// The default credential for `space` or `nil` if no default has been set.
//
// # Discussion
// 
// If you override this method, also override
// [GetDefaultCredentialForProtectionSpaceTaskCompletionHandler].
//
// See: https://developer.apple.com/documentation/Foundation/URLCredentialStorage/defaultCredential(for:)
func (u URLCredentialStorage) DefaultCredentialForProtectionSpace(space INSURLProtectionSpace) INSURLCredential {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("defaultCredentialForProtectionSpace:"), space)
	return NSURLCredentialFromID(rv)
}
// Gets the default credential for the specified protection space, which is
// being accessed by the given task, and passes it to the provided completion
// handler.
//
// space: The protection space of interest.
//
// task: The task seeking to use the protection space
//
// completionHandler: A completion handler that receives the default credential as its argument,
// or `nil` if there is no default credential for this combination of
// protection space and task.
//
// See: https://developer.apple.com/documentation/Foundation/URLCredentialStorage/getDefaultCredential(for:task:completionHandler:)
func (u URLCredentialStorage) GetDefaultCredentialForProtectionSpaceTaskCompletionHandler(space INSURLProtectionSpace, task INSURLSessionTask, completionHandler URLCredentialHandler) {
_block2, _cleanup2 := NewURLCredentialBlock(completionHandler)
	defer _cleanup2()
	objc.Send[objc.ID](u.ID, objc.Sel("getDefaultCredentialForProtectionSpace:task:completionHandler:"), space, task, _block2)
}
// Sets the default credential for a specified protection space.
//
// credential: The URL credential to set as the default for `space`. If the receiver does
// not contain `credential` in the specified protection space it will be
// added.
//
// space: The protection space whose default credential is being set.
//
// # Discussion
// 
// If you override this method, also override
// [SetDefaultCredentialForProtectionSpaceTask].
//
// See: https://developer.apple.com/documentation/Foundation/URLCredentialStorage/setDefaultCredential(_:for:)
func (u URLCredentialStorage) SetDefaultCredentialForProtectionSpace(credential INSURLCredential, space INSURLProtectionSpace) {
	objc.Send[objc.ID](u.ID, objc.Sel("setDefaultCredential:forProtectionSpace:"), credential, space)
}
// Sets the default credential for a given protection space, which is being
// accessed by the given task.
//
// credential: The URL credential to set as the default for the protection space. If the
// receiver does not contain `credential` in the specified protection space it
// will be added.
//
// protectionSpace: The protection space whose default credential is being set.
//
// task: The task accessing the specified protection space. Subclasses of
// [NSURLCredentialStorage] may use the request URL or other properties of
// this task to affect how the default credential is stored.
//
// See: https://developer.apple.com/documentation/Foundation/URLCredentialStorage/setDefaultCredential(_:for:task:)
func (u URLCredentialStorage) SetDefaultCredentialForProtectionSpaceTask(credential INSURLCredential, protectionSpace INSURLProtectionSpace, task INSURLSessionTask) {
	objc.Send[objc.ID](u.ID, objc.Sel("setDefaultCredential:forProtectionSpace:task:"), credential, protectionSpace, task)
}
// Removes the specified credential from the credential storage for the
// specified protection space.
//
// credential: The credential to remove.
//
// space: The protection space from which to remove the credential.
//
// # Discussion
// 
// If you override this method, also override
// [RemoveCredentialForProtectionSpaceOptionsTask].
//
// See: https://developer.apple.com/documentation/Foundation/URLCredentialStorage/remove(_:for:)
func (u URLCredentialStorage) RemoveCredentialForProtectionSpace(credential INSURLCredential, space INSURLProtectionSpace) {
	objc.Send[objc.ID](u.ID, objc.Sel("removeCredential:forProtectionSpace:"), credential, space)
}
// Removes the specified credential from the credential storage for the
// specified protection space using the given options.
//
// credential: The credential to remove.
//
// space: The protection space from which to remove the credential.
//
// options: A dictionary containing options to consider when removing the credential.
// 
// For possible keys, see [Dictionary key for credential removal options]. You
// should use this when trying to delete a credential that has the
// [URLCredentialPersistenceSynchronizable] policy.
// //
// [Dictionary key for credential removal options]: https://developer.apple.com/documentation/Foundation/dictionary-key-for-credential-removal-options
//
// # Discussion
// 
// The credential is removed from both persistent and temporary storage.
// 
// If you override this method, also override
// [RemoveCredentialForProtectionSpaceOptionsTask].
//
// See: https://developer.apple.com/documentation/Foundation/URLCredentialStorage/remove(_:for:options:)
func (u URLCredentialStorage) RemoveCredentialForProtectionSpaceOptions(credential INSURLCredential, space INSURLProtectionSpace, options INSDictionary) {
	objc.Send[objc.ID](u.ID, objc.Sel("removeCredential:forProtectionSpace:options:"), credential, space, options)
}
// Removes the specified credential from the credential storage for the
// specified protection space, on behalf of the given task and using the given
// options.
//
// credential: The credential to remove.
//
// protectionSpace: The protection space from which to remove the credential.
//
// options: A dictionary containing options to consider when removing the credential.
// 
// For possible keys, see [Dictionary key for credential removal options]. You
// should use this when trying to delete a credential that has the
// [URLCredentialPersistenceSynchronizable] policy.
// //
// [Dictionary key for credential removal options]: https://developer.apple.com/documentation/Foundation/dictionary-key-for-credential-removal-options
//
// task: The task using the protection space that you wish to remove the credential
// for.
//
// # Discussion
// 
// The credential is removed from both persistent and temporary storage.
//
// See: https://developer.apple.com/documentation/Foundation/URLCredentialStorage/remove(_:for:options:task:)
func (u URLCredentialStorage) RemoveCredentialForProtectionSpaceOptionsTask(credential INSURLCredential, protectionSpace INSURLProtectionSpace, options INSDictionary, task INSURLSessionTask) {
	objc.Send[objc.ID](u.ID, objc.Sel("removeCredential:forProtectionSpace:options:task:"), credential, protectionSpace, options, task)
}
// Adds a credential to the credential storage for the specified protection
// space.
//
// credential: The credential to add. If a credential with the same user name already
// exists in `space`, then `credential` replaces the existing object.
//
// space: The protection space to which to add the credential.
//
// # Discussion
// 
// If the credential is not yet in the set for the protection space, it will
// be added to it.
// 
// If you override this method, also override
// [SetCredentialForProtectionSpaceTask].
//
// See: https://developer.apple.com/documentation/Foundation/URLCredentialStorage/set(_:for:)
func (u URLCredentialStorage) SetCredentialForProtectionSpace(credential INSURLCredential, space INSURLProtectionSpace) {
	objc.Send[objc.ID](u.ID, objc.Sel("setCredential:forProtectionSpace:"), credential, space)
}
// Adds a credential to the credential storage for the specified protection
// space, on behalf of the specified task.
//
// credential: The credential to add. If a credential with the same user name already
// exists in `space`, then `credential` replaces the existing object.
//
// protectionSpace: The protection space to which to add the credential.
//
// task: The task accessing the specified protection space. Subclasses of
// [NSURLCredentialStorage] may use the request URL or other properties of
// this task to affect how the default credential is stored.
//
// See: https://developer.apple.com/documentation/Foundation/URLCredentialStorage/set(_:for:task:)
func (u URLCredentialStorage) SetCredentialForProtectionSpaceTask(credential INSURLCredential, protectionSpace INSURLProtectionSpace, task INSURLSessionTask) {
	objc.Send[objc.ID](u.ID, objc.Sel("setCredential:forProtectionSpace:task:"), credential, protectionSpace, task)
}
// Returns a dictionary containing the credentials for the specified
// protection space.
//
// space: The protection space whose credentials you want to retrieve.
//
// # Return Value
// 
// A dictionary containing the credentials for the specified protection space.
// The dictionary’s keys are user name strings, and each value is the
// corresponding [NSURLCredential].
//
// # Discussion
// 
// If you override this method, also override
// [GetCredentialsForProtectionSpaceTaskCompletionHandler].
//
// See: https://developer.apple.com/documentation/Foundation/URLCredentialStorage/credentials(for:)
func (u URLCredentialStorage) CredentialsForProtectionSpace(space INSURLProtectionSpace) INSDictionary {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("credentialsForProtectionSpace:"), space)
	return NSDictionaryFromID(rv)
}

// The credentials for all available protection spaces.
//
// # Discussion
// 
// The dictionary has keys corresponding to the [NSURLProtectionSpace]
// instances. The values are dictionaries where the keys are user name
// strings, and each value is the corresponding [NSURLCredential] instances.
//
// See: https://developer.apple.com/documentation/Foundation/URLCredentialStorage/allCredentials
func (u URLCredentialStorage) AllCredentials() INSDictionary {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("allCredentials"))
	return NSDictionaryFromID(objc.ID(rv))
}
// A notification posted when the set of stored credentials changes.
//
// See: https://developer.apple.com/documentation/foundation/nsnotification/name-swift.struct/nsurlcredentialstoragechanged
func (u URLCredentialStorage) NSURLCredentialStorageChanged() NSNotificationName {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("NSURLCredentialStorageChangedNotification"))
	return NSNotificationName(NSStringFromID(rv).String())
}

// The shared URL credential storage instance.
//
// See: https://developer.apple.com/documentation/Foundation/URLCredentialStorage/shared
func (_URLCredentialStorageClass URLCredentialStorageClass) SharedCredentialStorage() URLCredentialStorage {
	rv := objc.Send[objc.ID](objc.ID(_URLCredentialStorageClass.class), objc.Sel("sharedCredentialStorage"))
	return NSURLCredentialStorageFromID(objc.ID(rv))
}

// GetDefaultCredentialForProtectionSpaceTask is a synchronous wrapper around [URLCredentialStorage.GetDefaultCredentialForProtectionSpaceTaskCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (u URLCredentialStorage) GetDefaultCredentialForProtectionSpaceTask(ctx context.Context, space INSURLProtectionSpace, task INSURLSessionTask) (*NSURLCredential, error) {
	done := make(chan *NSURLCredential, 1)
	u.GetDefaultCredentialForProtectionSpaceTaskCompletionHandler(space, task, func(val *NSURLCredential) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

