// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [LARightStore] class.
var (
	_LARightStoreClass     LARightStoreClass
	_LARightStoreClassOnce sync.Once
)

func getLARightStoreClass() LARightStoreClass {
	_LARightStoreClassOnce.Do(func() {
		_LARightStoreClass = LARightStoreClass{class: objc.GetClass("LARightStore")}
	})
	return _LARightStoreClass
}

// GetLARightStoreClass returns the class object for LARightStore.
func GetLARightStoreClass() LARightStoreClass {
	return getLARightStoreClass()
}

type LARightStoreClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (lc LARightStoreClass) Class() objc.Class {
	return lc.class
}

// Alloc allocates memory for a new instance of the class.
func (lc LARightStoreClass) Alloc() LARightStore {
	rv := objc.Send[LARightStore](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// A container for data protected by a right.
//
// # Overview
//
// Use an [LARightStore] along with an [LARight] to make secrets accessible
// only after certain conditions, including authentication, are met. Storing
// secrets this way lets you tie the availability of sensitive resources to
// the authorization status of the user.
//
// The following stores a named access token behind the default authorization
// requirements:
//
// The system stores your secret in the keychain and protects it with a unique
// key in the Secure Enclave. The system associates the key with your right
// and with an access control list that ensures that the data is only
// accessible after your access requirements are met.
//
// You can retrieve stored secrets later using the right’s identifier:
//
// # Accessing rights
//
//   - [LARightStore.RightForIdentifierCompletion]: Fetches a previously stored right from the shared right store.
//
// # Storing rights
//
//   - [LARightStore.SaveRightIdentifierCompletion]: Saves a right to a persistent right store.
//   - [LARightStore.SaveRightIdentifierSecretCompletion]: Saves a right to a persistent store along with secret data you supply.
//
// # Removing stored rights
//
//   - [LARightStore.RemoveRightCompletion]: Removes a right from the right store given an instance of that right.
//   - [LARightStore.RemoveRightForIdentifierCompletion]: Removes a right from the right store given its unique identifier.
//   - [LARightStore.RemoveAllRightsWithCompletion]: Removes all rights associated with this client from the right store.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LARightStore
type LARightStore struct {
	objectivec.Object
}

// LARightStoreFromID constructs a [LARightStore] from an objc.ID.
//
// A container for data protected by a right.
func LARightStoreFromID(id objc.ID) LARightStore {
	return LARightStore{objectivec.Object{ID: id}}
}

// NOTE: LARightStore adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [LARightStore] class.
//
// # Accessing rights
//
//   - [ILARightStore.RightForIdentifierCompletion]: Fetches a previously stored right from the shared right store.
//
// # Storing rights
//
//   - [ILARightStore.SaveRightIdentifierCompletion]: Saves a right to a persistent right store.
//   - [ILARightStore.SaveRightIdentifierSecretCompletion]: Saves a right to a persistent store along with secret data you supply.
//
// # Removing stored rights
//
//   - [ILARightStore.RemoveRightCompletion]: Removes a right from the right store given an instance of that right.
//   - [ILARightStore.RemoveRightForIdentifierCompletion]: Removes a right from the right store given its unique identifier.
//   - [ILARightStore.RemoveAllRightsWithCompletion]: Removes all rights associated with this client from the right store.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LARightStore
type ILARightStore interface {
	objectivec.IObject

	// Topic: Accessing rights

	// Fetches a previously stored right from the shared right store.
	RightForIdentifierCompletion(identifier string, handler LAPersistedRightErrorHandler)

	// Topic: Storing rights

	// Saves a right to a persistent right store.
	SaveRightIdentifierCompletion(right ILARight, identifier string, handler LAPersistedRightErrorHandler)
	// Saves a right to a persistent store along with secret data you supply.
	SaveRightIdentifierSecretCompletion(right ILARight, identifier string, secret foundation.INSData, handler LAPersistedRightErrorHandler)

	// Topic: Removing stored rights

	// Removes a right from the right store given an instance of that right.
	RemoveRightCompletion(right ILAPersistedRight, handler ErrorHandler)
	// Removes a right from the right store given its unique identifier.
	RemoveRightForIdentifierCompletion(identifier string, handler ErrorHandler)
	// Removes all rights associated with this client from the right store.
	RemoveAllRightsWithCompletion(handler ErrorHandler)
}

// Init initializes the instance.
func (r LARightStore) Init() LARightStore {
	rv := objc.Send[LARightStore](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r LARightStore) Autorelease() LARightStore {
	rv := objc.Send[LARightStore](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewLARightStore creates a new LARightStore instance.
func NewLARightStore() LARightStore {
	class := getLARightStoreClass()
	rv := objc.Send[LARightStore](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Fetches a previously stored right from the shared right store.
//
// identifier: The unique identifier of the right.
//
// handler: A completion handler to call when the right access completes.
//
// `right`: The right that matches the `identifier` you supply. `error`: An
// error object that indicates why the `right` parameter is `nil`, or `nil` if
// the right parameter is non-`nil`.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LARightStore/right(forIdentifier:completion:)
func (r LARightStore) RightForIdentifierCompletion(identifier string, handler LAPersistedRightErrorHandler) {
	_block1, _ := NewLAPersistedRightErrorBlock(handler)
	objc.Send[objc.ID](r.ID, objc.Sel("rightForIdentifier:completion:"), objc.String(identifier), _block1)
}

// Saves a right to a persistent right store.
//
// right: The right to store.
//
// identifier: A unique identifier for the right.
//
// handler: A completion handler to call when the save operation completes.
//
// `right`: The persisted form of the right that the save operation stores.
// `error`: An error object that indicates why the `right` parameter is `nil`,
// or `nil` if the right parameter is non-`nil`.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LARightStore/saveRight(_:identifier:completion:)
func (r LARightStore) SaveRightIdentifierCompletion(right ILARight, identifier string, handler LAPersistedRightErrorHandler) {
	_block2, _ := NewLAPersistedRightErrorBlock(handler)
	objc.Send[objc.ID](r.ID, objc.Sel("saveRight:identifier:completion:"), right, objc.String(identifier), _block2)
}

// Saves a right to a persistent store along with secret data you supply.
//
// right: The right to store.
//
// identifier: A unique identifier for the right.
//
// secret: Secret data that’s stored with the right.
//
// handler: A completion handler to call when the save operation completes.
//
// `right`: The persisted form of the right that the save operation stores.
// `error`: An error object that indicates why the `right` parameter is `nil`,
// or `nil` if the right parameter is non-`nil`.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LARightStore/saveRight(_:identifier:secret:completion:)
func (r LARightStore) SaveRightIdentifierSecretCompletion(right ILARight, identifier string, secret foundation.INSData, handler LAPersistedRightErrorHandler) {
	_block3, _ := NewLAPersistedRightErrorBlock(handler)
	objc.Send[objc.ID](r.ID, objc.Sel("saveRight:identifier:secret:completion:"), right, objc.String(identifier), secret, _block3)
}

// Removes a right from the right store given an instance of that right.
//
// right: An instance that represents the right to remove.
//
// handler: A completion handler to call when the removal operation completes.
//
// `error`: An error object that indicates why the removal operation failed,
// or `nil` if it succeeded.
//
// # Discussion
//
// Removing a right also removes any resources stored along with the right,
// such as secrets.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LARightStore/removeRight(_:completion:)
func (r LARightStore) RemoveRightCompletion(right ILAPersistedRight, handler ErrorHandler) {
	_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](r.ID, objc.Sel("removeRight:completion:"), right, _block1)
}

// Removes a right from the right store given its unique identifier.
//
// identifier: The identifier for the right to remove.
//
// handler: A completion handler to call when the removal operation completes.
//
// `error`: An error object that indicates why the removal operation failed,
// or `nil` if it succeeded.
//
// # Discussion
//
// Removing a right also removes any resources stored along with the right,
// such as secrets.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LARightStore/removeRight(forIdentifier:completion:)
func (r LARightStore) RemoveRightForIdentifierCompletion(identifier string, handler ErrorHandler) {
	_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](r.ID, objc.Sel("removeRightForIdentifier:completion:"), objc.String(identifier), _block1)
}

// Removes all rights associated with this client from the right store.
//
// handler: A completion handler to call when the removal operation completes.
//
// `error`: An error object that indicates why the removal operation failed,
// or `nil` if it succeeded.
//
// # Discussion
//
// Removing rights also removes any resources stored along with the rights,
// such as secrets.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LARightStore/removeAllRights(completion:)
func (r LARightStore) RemoveAllRightsWithCompletion(handler ErrorHandler) {
	_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](r.ID, objc.Sel("removeAllRightsWithCompletion:"), _block0)
}

// A shared object that stores rights.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LARightStore/shared
func (_LARightStoreClass LARightStoreClass) SharedStore() LARightStore {
	rv := objc.Send[objc.ID](objc.ID(_LARightStoreClass.class), objc.Sel("sharedStore"))
	return LARightStoreFromID(objc.ID(rv))
}

// RightForIdentifierCompletionSync is a synchronous wrapper around [LARightStore.RightForIdentifierCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (r LARightStore) RightForIdentifierCompletionSync(ctx context.Context, identifier string) (*LAPersistedRight, error) {
	type result struct {
		val *LAPersistedRight
		err error
	}
	done := make(chan result, 1)
	r.RightForIdentifierCompletion(identifier, func(val *LAPersistedRight, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// SaveRightIdentifierCompletionSync is a synchronous wrapper around [LARightStore.SaveRightIdentifierCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (r LARightStore) SaveRightIdentifierCompletionSync(ctx context.Context, right ILARight, identifier string) (*LAPersistedRight, error) {
	type result struct {
		val *LAPersistedRight
		err error
	}
	done := make(chan result, 1)
	r.SaveRightIdentifierCompletion(right, identifier, func(val *LAPersistedRight, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// SaveRightIdentifierSecretCompletionSync is a synchronous wrapper around [LARightStore.SaveRightIdentifierSecretCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (r LARightStore) SaveRightIdentifierSecretCompletionSync(ctx context.Context, right ILARight, identifier string, secret foundation.INSData) (*LAPersistedRight, error) {
	type result struct {
		val *LAPersistedRight
		err error
	}
	done := make(chan result, 1)
	r.SaveRightIdentifierSecretCompletion(right, identifier, secret, func(val *LAPersistedRight, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// RemoveRightCompletionSync is a synchronous wrapper around [LARightStore.RemoveRightCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (r LARightStore) RemoveRightCompletionSync(ctx context.Context, right ILAPersistedRight) error {
	done := make(chan error, 1)
	r.RemoveRightCompletion(right, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RemoveRightForIdentifierCompletionSync is a synchronous wrapper around [LARightStore.RemoveRightForIdentifierCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (r LARightStore) RemoveRightForIdentifierCompletionSync(ctx context.Context, identifier string) error {
	done := make(chan error, 1)
	r.RemoveRightForIdentifierCompletion(identifier, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RemoveAllRights is a synchronous wrapper around [LARightStore.RemoveAllRightsWithCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (r LARightStore) RemoveAllRights(ctx context.Context) error {
	done := make(chan error, 1)
	r.RemoveAllRightsWithCompletion(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
