// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [URLCredential] class.
var (
	_URLCredentialClass     URLCredentialClass
	_URLCredentialClassOnce sync.Once
)

func getURLCredentialClass() URLCredentialClass {
	_URLCredentialClassOnce.Do(func() {
		_URLCredentialClass = URLCredentialClass{class: objc.GetClass("NSURLCredential")}
	})
	return _URLCredentialClass
}

// GetURLCredentialClass returns the class object for NSURLCredential.
func GetURLCredentialClass() URLCredentialClass {
	return getURLCredentialClass()
}

type URLCredentialClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (uc URLCredentialClass) Class() objc.Class {
	return uc.class
}

// Alloc allocates memory for a new instance of the class.
func (uc URLCredentialClass) Alloc() URLCredential {
	rv := objc.Send[URLCredential](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// [A]n authentication credential consisting of information specific to the
// type of credential and the type of persistent storage to use, if any.
//
// # Overview
//
// The URL Loading System supports password-based user credentials,
// certificate-based user credentials, and certificate-based server
// credentials.
//
// When you create a credential, you can specify it for a single request,
// persist it temporarily (until your app quits), or persist it permanently.
// Permanent persistence can be local persistence in the keychain, or
// synchronized persistence across the user’s devices, based on their Apple
// ID.
//
// # Creating a credential
//
//   - [URLCredential.InitWithIdentityCertificatesPersistence]: Creates a URL credential instance for resolving a client certificate authentication challenge.
//   - [URLCredential.InitWithTrust]: Creates a URL credential instance for server trust authentication, initialized with a accepted trust.
//   - [URLCredential.InitWithUserPasswordPersistence]: Creates a URL credential instance initialized with a given user name and password, using a given persistence setting.
//
// # Getting credential properties
//
//   - [URLCredential.User]: The credential’s user name.
//   - [URLCredential.Certificates]: The intermediate certificates of the credential, if it is a client certificate credential.
//   - [URLCredential.HasPassword]: A Boolean value that indicates whether the credential has a password.
//   - [URLCredential.Password]: The credential’s password.
//   - [URLCredential.Identity]: The identity of this credential if it is a client certificate credential.
//   - [URLCredential.Persistence]: The credential’s persistence setting.
//
// See: https://developer.apple.com/documentation/Foundation/URLCredential
type URLCredential struct {
	objectivec.Object
}

// URLCredentialFromID constructs a [URLCredential] from an objc.ID.
//
// [A]n authentication credential consisting of information specific to the
// type of credential and the type of persistent storage to use, if any.
func URLCredentialFromID(id objc.ID) URLCredential {
	return NSURLCredential{objectivec.Object{ID: id}}
}

// NSURLCredentialFromID is an alias for [URLCredentialFromID] for cross-framework compatibility.
func NSURLCredentialFromID(id objc.ID) URLCredential { return URLCredentialFromID(id) }

// NOTE: URLCredential adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [URLCredential] class.
//
// # Creating a credential
//
//   - [IURLCredential.InitWithIdentityCertificatesPersistence]: Creates a URL credential instance for resolving a client certificate authentication challenge.
//   - [IURLCredential.InitWithTrust]: Creates a URL credential instance for server trust authentication, initialized with a accepted trust.
//   - [IURLCredential.InitWithUserPasswordPersistence]: Creates a URL credential instance initialized with a given user name and password, using a given persistence setting.
//
// # Getting credential properties
//
//   - [IURLCredential.User]: The credential’s user name.
//   - [IURLCredential.Certificates]: The intermediate certificates of the credential, if it is a client certificate credential.
//   - [IURLCredential.HasPassword]: A Boolean value that indicates whether the credential has a password.
//   - [IURLCredential.Password]: The credential’s password.
//   - [IURLCredential.Identity]: The identity of this credential if it is a client certificate credential.
//   - [IURLCredential.Persistence]: The credential’s persistence setting.
//
// See: https://developer.apple.com/documentation/Foundation/URLCredential
type IURLCredential interface {
	objectivec.IObject
	NSCoding
	NSCopying
	NSSecureCoding

	// Topic: Creating a credential

	// Creates a URL credential instance for resolving a client certificate authentication challenge.
	InitWithIdentityCertificatesPersistence(identity objectivec.IObject, certArray INSArray, persistence NSURLCredentialPersistence) URLCredential
	// Creates a URL credential instance for server trust authentication, initialized with a accepted trust.
	InitWithTrust(trust objectivec.IObject) URLCredential
	// Creates a URL credential instance initialized with a given user name and password, using a given persistence setting.
	InitWithUserPasswordPersistence(user string, password string, persistence NSURLCredentialPersistence) URLCredential

	// Topic: Getting credential properties

	// The credential’s user name.
	User() string
	// The intermediate certificates of the credential, if it is a client certificate credential.
	Certificates() INSArray
	// A Boolean value that indicates whether the credential has a password.
	HasPassword() bool
	// The credential’s password.
	Password() string
	// The identity of this credential if it is a client certificate credential.
	Identity() objectivec.IObject
	// The credential’s persistence setting.
	Persistence() NSURLCredentialPersistence
}

// Init initializes the instance.
func (u URLCredential) Init() URLCredential {
	rv := objc.Send[URLCredential](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u URLCredential) Autorelease() URLCredential {
	rv := objc.Send[URLCredential](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLCredential creates a new URLCredential instance.
func NewURLCredential() URLCredential {
	class := getURLCredentialClass()
	rv := objc.Send[URLCredential](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewURLCredentialWithCoder(coder INSCoder) URLCredential {
	instance := getURLCredentialClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return URLCredentialFromID(rv)
}

// Creates a URL credential instance for resolving a client certificate
// authentication challenge.
//
// identity: The identity for the credential.
//
// certArray: An array of one or more [SecCertificateRef] objects representing
// intermediate certificates leading from the identity’s certificate to a
// trusted root, or `nil` if the server does not need any intermediate
// certificates to authenticate the client.
//
// persistence: The method ignores this parameter; you should supply a value of
// [NSURLCredentialPersistenceForSession] because that most accurately
// reflects the actual behaviour.
//
// # Return Value
//
// A new URL credential object, using the provided identity and, optionally,
// an array of intermediate certificates.
//
// # Discussion
//
// When you receive a client certificate authentication challenge
// ([NSURLAuthenticationMethodClientCertificate]) and want to resolve it
// successfully, you must supply a credential created using this initializer.
//
// In most cases you should pass `nil` to the `certArray` parameter. You only
// need to supply an array of intermediate certificates if the server needs
// those intermediate certificates to authenticate the client. Typically this
// isn’t necessary because the server already has a copy of the relevant
// intermediate certificates.
//
// See: https://developer.apple.com/documentation/Foundation/URLCredential/init(identity:certificates:persistence:)
//
// [NSURLAuthenticationMethodClientCertificate]: https://developer.apple.com/documentation/Foundation/NSURLAuthenticationMethodClientCertificate
func NewURLCredentialWithIdentityCertificatesPersistence(identity objectivec.IObject, certArray INSArray, persistence NSURLCredentialPersistence) URLCredential {
	instance := getURLCredentialClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIdentity:certificates:persistence:"), identity, certArray, persistence)
	return URLCredentialFromID(rv)
}

// Creates a URL credential instance for server trust authentication,
// initialized with a accepted trust.
//
// trust: The accepted trust.
//
// # Return Value
//
// A new URL credential object, containing the provided server trust.
//
// # Discussion
//
// Before your implementation of
// [URLSessionTaskDidReceiveChallengeCompletionHandler] uses this initializer
// to create a server trust credential, you are responsible for evaluating the
// received [SecTrust] instance. You get this [ServerTrust] from the
// [ProtectionSpace] of the [NSURLAuthenticationChallenge] parameter that is
// passed to your delegate method. Pass the trust instance to
// [SecTrustEvaluate(_:_:)] to evaluate it. If this call indicates the trust
// is invalid, you should cancel the challenge by passing the
// [NSURLSessionAuthChallengeCancelAuthenticationChallenge] disposition to the
// completion handler.
//
// See: https://developer.apple.com/documentation/Foundation/URLCredential/init(trust:)
//
// [SecTrustEvaluate(_:_:)]: https://developer.apple.com/documentation/Security/SecTrustEvaluate(_:_:)
// [SecTrust]: https://developer.apple.com/documentation/Security/SecTrust
func NewURLCredentialWithTrust(trust objectivec.IObject) URLCredential {
	instance := getURLCredentialClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTrust:"), trust)
	return URLCredentialFromID(rv)
}

// Creates a URL credential instance initialized with a given user name and
// password, using a given persistence setting.
//
// user: The user for the credential.
//
// password: The password for `user`.
//
// persistence: A [URLCredential.Persistence] value indicating whether the credential
// should be stored permanently, for the duration of the current session, or
// not at all.
//
// # Return Value
//
// An instance of [NSURLCredential], initialized with user name `user`,
// password `password`, and using persistence setting `persistence`.
//
// # Discussion
//
// If `persistence` is [NSURLCredentialPersistencePermanent], the credential
// is stored in the keychain. If `persistence` is
// [NSURLCredentialPersistenceSynchronizable], it is also stored to the
// user’s other devices.
//
// See: https://developer.apple.com/documentation/Foundation/URLCredential/init(user:password:persistence:)
//
// [URLCredential.Persistence]: https://developer.apple.com/documentation/Foundation/URLCredential/Persistence-swift.enum
func NewURLCredentialWithUserPasswordPersistence(user string, password string, persistence NSURLCredentialPersistence) URLCredential {
	instance := getURLCredentialClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUser:password:persistence:"), objc.String(user), objc.String(password), persistence)
	return URLCredentialFromID(rv)
}

// Creates a URL credential instance for resolving a client certificate
// authentication challenge.
//
// identity: The identity for the credential.
//
// certArray: An array of one or more [SecCertificateRef] objects representing
// intermediate certificates leading from the identity’s certificate to a
// trusted root, or `nil` if the server does not need any intermediate
// certificates to authenticate the client.
//
// persistence: The method ignores this parameter; you should supply a value of
// [NSURLCredentialPersistenceForSession] because that most accurately
// reflects the actual behaviour.
//
// # Return Value
//
// A new URL credential object, using the provided identity and, optionally,
// an array of intermediate certificates.
//
// # Discussion
//
// When you receive a client certificate authentication challenge
// ([NSURLAuthenticationMethodClientCertificate]) and want to resolve it
// successfully, you must supply a credential created using this initializer.
//
// In most cases you should pass `nil` to the `certArray` parameter. You only
// need to supply an array of intermediate certificates if the server needs
// those intermediate certificates to authenticate the client. Typically this
// isn’t necessary because the server already has a copy of the relevant
// intermediate certificates.
//
// See: https://developer.apple.com/documentation/Foundation/URLCredential/init(identity:certificates:persistence:)
//
// [NSURLAuthenticationMethodClientCertificate]: https://developer.apple.com/documentation/Foundation/NSURLAuthenticationMethodClientCertificate
func (u URLCredential) InitWithIdentityCertificatesPersistence(identity objectivec.IObject, certArray INSArray, persistence NSURLCredentialPersistence) URLCredential {
	rv := objc.Send[URLCredential](u.ID, objc.Sel("initWithIdentity:certificates:persistence:"), identity, certArray, persistence)
	return rv
}

// Creates a URL credential instance for server trust authentication,
// initialized with a accepted trust.
//
// trust: The accepted trust.
//
// # Return Value
//
// A new URL credential object, containing the provided server trust.
//
// # Discussion
//
// Before your implementation of
// [URLSessionTaskDidReceiveChallengeCompletionHandler] uses this initializer
// to create a server trust credential, you are responsible for evaluating the
// received [SecTrust] instance. You get this [ServerTrust] from the
// [ProtectionSpace] of the [NSURLAuthenticationChallenge] parameter that is
// passed to your delegate method. Pass the trust instance to
// [SecTrustEvaluate(_:_:)] to evaluate it. If this call indicates the trust
// is invalid, you should cancel the challenge by passing the
// [NSURLSessionAuthChallengeCancelAuthenticationChallenge] disposition to the
// completion handler.
//
// See: https://developer.apple.com/documentation/Foundation/URLCredential/init(trust:)
//
// [SecTrustEvaluate(_:_:)]: https://developer.apple.com/documentation/Security/SecTrustEvaluate(_:_:)
// [SecTrust]: https://developer.apple.com/documentation/Security/SecTrust
func (u URLCredential) InitWithTrust(trust objectivec.IObject) URLCredential {
	rv := objc.Send[URLCredential](u.ID, objc.Sel("initWithTrust:"), trust)
	return rv
}

// Creates a URL credential instance initialized with a given user name and
// password, using a given persistence setting.
//
// user: The user for the credential.
//
// password: The password for `user`.
//
// persistence: A [URLCredential.Persistence] value indicating whether the credential
// should be stored permanently, for the duration of the current session, or
// not at all.
//
// # Return Value
//
// An instance of [NSURLCredential], initialized with user name `user`,
// password `password`, and using persistence setting `persistence`.
//
// # Discussion
//
// If `persistence` is [NSURLCredentialPersistencePermanent], the credential
// is stored in the keychain. If `persistence` is
// [NSURLCredentialPersistenceSynchronizable], it is also stored to the
// user’s other devices.
//
// See: https://developer.apple.com/documentation/Foundation/URLCredential/init(user:password:persistence:)
//
// [URLCredential.Persistence]: https://developer.apple.com/documentation/Foundation/URLCredential/Persistence-swift.enum
func (u URLCredential) InitWithUserPasswordPersistence(user string, password string, persistence NSURLCredentialPersistence) URLCredential {
	rv := objc.Send[URLCredential](u.ID, objc.Sel("initWithUser:password:persistence:"), objc.String(user), objc.String(password), persistence)
	return rv
}

// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (u URLCredential) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](u.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (u URLCredential) InitWithCoder(coder INSCoder) URLCredential {
	rv := objc.Send[URLCredential](u.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// # Return Value
//
// # The new autoreleased NSURLCredential
//
// # Discussion
//
// Create a new NSURLCredential which specifies that a handshake has been
// trusted.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLCredential/credentialForTrust:
func (_URLCredentialClass URLCredentialClass) CredentialForTrust(trust objectivec.IObject) URLCredential {
	rv := objc.Send[objc.ID](objc.ID(_URLCredentialClass.class), objc.Sel("credentialForTrust:"), trust)
	return NSURLCredentialFromID(rv)
}

// Creates a URL credential instance for resolving a client certificate
// authentication challenge.
//
// identity: The identity for the credential.
//
// certArray: An array of one or more [SecCertificateRef] objects representing
// intermediate certificates leading from the identity’s certificate to a
// trusted root, or `nil` if the server does not need any intermediate
// certificates to authenticate the client.
//
// persistence: The method ignores this parameter; you should supply a value of
// [NSURLCredentialPersistenceForSession] because that most accurately
// reflects the actual behaviour.
//
// # Return Value
//
// A new URL credential object, using the provided identity and, optionally,
// an array of intermediate certificates.
//
// # Discussion
//
// When you receive a client certificate authentication challenge
// ([NSURLAuthenticationMethodClientCertificate]) and want to resolve it
// successfully, you must supply a credential created using this method.
//
// In most cases you should pass `nil` to the `certArray` parameter. You only
// need to supply an array of intermediate certificates if the server needs
// those intermediate certificates to authenticate the client. Typically this
// isn’t necessary because the server already has a copy of the relevant
// intermediate certificates.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLCredential/credentialWithIdentity:certificates:persistence:
//
// [NSURLAuthenticationMethodClientCertificate]: https://developer.apple.com/documentation/Foundation/NSURLAuthenticationMethodClientCertificate
func (_URLCredentialClass URLCredentialClass) CredentialWithIdentityCertificatesPersistence(identity objectivec.IObject, certArray INSArray, persistence NSURLCredentialPersistence) URLCredential {
	rv := objc.Send[objc.ID](objc.ID(_URLCredentialClass.class), objc.Sel("credentialWithIdentity:certificates:persistence:"), identity, certArray, persistence)
	return NSURLCredentialFromID(rv)
}

// Creates a URL credential instance for internet password authentication with
// a given user name and password, using a given persistence setting.
//
// user: The user for the credential.
//
// password: The password for `user`.
//
// persistence: A [URLCredential.Persistence] value indicating whether the credential
// should be stored permanently, for the duration of the current session, or
// not at all.
//
// # Return Value
//
// A new URL credential object with user name `user`, password `password`, and
// using persistence setting `persistence`.
//
// # Discussion
//
// If `persistence` is [NSURLCredentialPersistencePermanent], the credential
// is stored in the keychain. If `persistence` is
// [NSURLCredentialPersistenceSynchronizable], it is also synchronized to the
// user’s other devices.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLCredential/credentialWithUser:password:persistence:
//
// [URLCredential.Persistence]: https://developer.apple.com/documentation/Foundation/URLCredential/Persistence-swift.enum
func (_URLCredentialClass URLCredentialClass) CredentialWithUserPasswordPersistence(user string, password string, persistence NSURLCredentialPersistence) URLCredential {
	rv := objc.Send[objc.ID](objc.ID(_URLCredentialClass.class), objc.Sel("credentialWithUser:password:persistence:"), objc.String(user), objc.String(password), persistence)
	return NSURLCredentialFromID(rv)
}

// The credential’s user name.
//
// See: https://developer.apple.com/documentation/Foundation/URLCredential/user
func (u URLCredential) User() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("user"))
	return NSStringFromID(rv).String()
}

// The intermediate certificates of the credential, if it is a client
// certificate credential.
//
// # Discussion
//
// The certificates are [SecCertificate] objects representing the intermediate
// certificates of the credential. This value is `nil` if this is not a client
// certificate credential or if the credential was created with no
// intermediate certificates.
//
// See: https://developer.apple.com/documentation/Foundation/URLCredential/certificates
//
// [SecCertificate]: https://developer.apple.com/documentation/Security/SecCertificate
func (u URLCredential) Certificates() INSArray {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("certificates"))
	return NSArrayFromID(objc.ID(rv))
}

// A Boolean value that indicates whether the credential has a password.
//
// # Discussion
//
// This value is true if the receiver has a password, false otherwise.
//
// This method does not attempt to retrieve the password.
//
// If this credential’s password is stored in the user’s keychain,
// [Password] may return `nil` even if this method returns true—getting the
// password may fail, or the user may refuse access.
//
// See: https://developer.apple.com/documentation/Foundation/URLCredential/hasPassword
func (u URLCredential) HasPassword() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("hasPassword"))
	return rv
}

// The credential’s password.
//
// # Discussion
//
// You should only access this property if you need the actual password value.
// If you only need to know if there is a password, use [HasPassword].
// Accessing this property may result in prompting the user for access—for
// example, if the password is stored in the user’s keychain.
//
// See: https://developer.apple.com/documentation/Foundation/URLCredential/password
func (u URLCredential) Password() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("password"))
	return NSStringFromID(rv).String()
}

// The identity of this credential if it is a client certificate credential.
//
// # Discussion
//
// This value is `nil` if the credential is not a client certificate
// credential.
//
// See: https://developer.apple.com/documentation/Foundation/URLCredential/identity
func (u URLCredential) Identity() objectivec.IObject {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("identity"))
	return objectivec.Object{ID: rv}
}

// The credential’s persistence setting.
//
// See: https://developer.apple.com/documentation/Foundation/URLCredential/persistence-swift.property
func (u URLCredential) Persistence() NSURLCredentialPersistence {
	rv := objc.Send[NSURLCredentialPersistence](u.ID, objc.Sel("persistence"))
	return NSURLCredentialPersistence(rv)
}

// Protocol methods for NSCopying

// Protocol methods for NSSecureCoding
