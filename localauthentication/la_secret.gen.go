// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [LASecret] class.
var (
	_LASecretClass     LASecretClass
	_LASecretClassOnce sync.Once
)

func getLASecretClass() LASecretClass {
	_LASecretClassOnce.Do(func() {
		_LASecretClass = LASecretClass{class: objc.GetClass("LASecret")}
	})
	return _LASecretClass
}

// GetLASecretClass returns the class object for LASecret.
func GetLASecretClass() LASecretClass {
	return getLASecretClass()
}

type LASecretClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (lc LASecretClass) Class() objc.Class {
	return lc.class
}

// Alloc allocates memory for a new instance of the class.
func (lc LASecretClass) Alloc() LASecret {
	rv := objc.Send[LASecret](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// Data that’s protected by a persisted right.
//
// # Overview
//
// You create [LASecret] instances when you store an [LAPersistedRight]; you
// can’t create them directly.
//
// # Loading secret data
//
//   - [LASecret.LoadDataWithCompletion]: Retrieves data stored in a secret.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LASecret
type LASecret struct {
	objectivec.Object
}

// LASecretFromID constructs a [LASecret] from an objc.ID.
//
// Data that’s protected by a persisted right.
func LASecretFromID(id objc.ID) LASecret {
	return LASecret{objectivec.Object{ID: id}}
}

// NOTE: LASecret adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [LASecret] class.
//
// # Loading secret data
//
//   - [ILASecret.LoadDataWithCompletion]: Retrieves data stored in a secret.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LASecret
type ILASecret interface {
	objectivec.IObject

	// Topic: Loading secret data

	// Retrieves data stored in a secret.
	LoadDataWithCompletion(handler DataErrorHandler)
}

// Init initializes the instance.
func (s LASecret) Init() LASecret {
	rv := objc.Send[LASecret](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s LASecret) Autorelease() LASecret {
	rv := objc.Send[LASecret](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewLASecret creates a new LASecret instance.
func NewLASecret() LASecret {
	class := getLASecretClass()
	rv := objc.Send[LASecret](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Retrieves data stored in a secret.
//
// handler: A completion handler that provides the data stored in a secret.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LASecret/loadData(completion:)
func (s LASecret) LoadDataWithCompletion(handler DataErrorHandler) {
	_block0, _ := NewDataErrorBlock(handler)
	objc.Send[objc.ID](s.ID, objc.Sel("loadDataWithCompletion:"), _block0)
}

// LoadData is a synchronous wrapper around [LASecret.LoadDataWithCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (s LASecret) LoadData(ctx context.Context) (*foundation.NSData, error) {
	type result struct {
		val *foundation.NSData
		err error
	}
	done := make(chan result, 1)
	s.LoadDataWithCompletion(func(val *foundation.NSData, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
