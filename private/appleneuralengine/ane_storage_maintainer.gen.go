// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANEStorageMaintainer] class.
var (
	_ANEStorageMaintainerClass     ANEStorageMaintainerClass
	_ANEStorageMaintainerClassOnce sync.Once
)

func getANEStorageMaintainerClass() ANEStorageMaintainerClass {
	_ANEStorageMaintainerClassOnce.Do(func() {
		_ANEStorageMaintainerClass = ANEStorageMaintainerClass{class: objc.GetClass("_ANEStorageMaintainer")}
	})
	return _ANEStorageMaintainerClass
}

// GetANEStorageMaintainerClass returns the class object for _ANEStorageMaintainer.
func GetANEStorageMaintainerClass() ANEStorageMaintainerClass {
	return getANEStorageMaintainerClass()
}

type ANEStorageMaintainerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEStorageMaintainerClass) Alloc() ANEStorageMaintainer {
	rv := objc.Send[ANEStorageMaintainer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [ANEStorageMaintainer.PurgeDanglingModelsAtWithReply]
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStorageMaintainer
type ANEStorageMaintainer struct {
	objectivec.Object
}

// ANEStorageMaintainerFromID constructs a [ANEStorageMaintainer] from an objc.ID.
func ANEStorageMaintainerFromID(id objc.ID) ANEStorageMaintainer {
	return ANEStorageMaintainer{objectivec.Object{id}}
}
// Ensure ANEStorageMaintainer implements IANEStorageMaintainer.
var _ IANEStorageMaintainer = ANEStorageMaintainer{}

// An interface definition for the [ANEStorageMaintainer] class.
//
// # Methods
//
//   - [IANEStorageMaintainer.PurgeDanglingModelsAtWithReply]
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStorageMaintainer
type IANEStorageMaintainer interface {
	objectivec.IObject

	// Topic: Methods

	PurgeDanglingModelsAtWithReply(at objectivec.IObject, reply VoidHandler)
}

// Init initializes the instance.
func (a ANEStorageMaintainer) Init() ANEStorageMaintainer {
	rv := objc.Send[ANEStorageMaintainer](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEStorageMaintainer) Autorelease() ANEStorageMaintainer {
	rv := objc.Send[ANEStorageMaintainer](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEStorageMaintainer creates a new ANEStorageMaintainer instance.
func NewANEStorageMaintainer() ANEStorageMaintainer {
	class := getANEStorageMaintainerClass()
	rv := objc.Send[ANEStorageMaintainer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStorageMaintainer/purgeDanglingModelsAt:withReply:
func (a ANEStorageMaintainer) PurgeDanglingModelsAtWithReply(at objectivec.IObject, reply VoidHandler) {
_block1, _cleanup1 := NewVoidBlock(reply)
	defer _cleanup1()
	objc.Send[objc.ID](a.ID, objc.Sel("purgeDanglingModelsAt:withReply:"), at, _block1)
}

// PurgeDanglingModelsAtWithReplySync is a synchronous wrapper around [ANEStorageMaintainer.PurgeDanglingModelsAtWithReply].
// It blocks until the completion handler fires or the context is cancelled.
func (a ANEStorageMaintainer) PurgeDanglingModelsAtWithReplySync(ctx context.Context, at objectivec.IObject) error {
	done := make(chan struct{}, 1)
	a.PurgeDanglingModelsAtWithReply(at, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

