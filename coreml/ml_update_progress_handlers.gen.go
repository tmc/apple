// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"context"
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLUpdateProgressHandlers] class.
var (
	_MLUpdateProgressHandlersClass     MLUpdateProgressHandlersClass
	_MLUpdateProgressHandlersClassOnce sync.Once
)

func getMLUpdateProgressHandlersClass() MLUpdateProgressHandlersClass {
	_MLUpdateProgressHandlersClassOnce.Do(func() {
		_MLUpdateProgressHandlersClass = MLUpdateProgressHandlersClass{class: objc.GetClass("MLUpdateProgressHandlers")}
	})
	return _MLUpdateProgressHandlersClass
}

// GetMLUpdateProgressHandlersClass returns the class object for MLUpdateProgressHandlers.
func GetMLUpdateProgressHandlersClass() MLUpdateProgressHandlersClass {
	return getMLUpdateProgressHandlersClass()
}

type MLUpdateProgressHandlersClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLUpdateProgressHandlersClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLUpdateProgressHandlersClass) Alloc() MLUpdateProgressHandlers {
	rv := objc.Send[MLUpdateProgressHandlers](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A collection of closures an update task uses to notify your app of its
// progress.
//
// # Creating progress handlers
//
//   - [MLUpdateProgressHandlers.InitForEventsProgressHandlerCompletionHandler]: Creates the collection of closures an update task uses to notify your app of its progress.
//
// See: https://developer.apple.com/documentation/CoreML/MLUpdateProgressHandlers
type MLUpdateProgressHandlers struct {
	objectivec.Object
}

// MLUpdateProgressHandlersFromID constructs a [MLUpdateProgressHandlers] from an objc.ID.
//
// A collection of closures an update task uses to notify your app of its
// progress.
func MLUpdateProgressHandlersFromID(id objc.ID) MLUpdateProgressHandlers {
	return MLUpdateProgressHandlers{objectivec.Object{ID: id}}
}

// NOTE: MLUpdateProgressHandlers adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MLUpdateProgressHandlers] class.
//
// # Creating progress handlers
//
//   - [IMLUpdateProgressHandlers.InitForEventsProgressHandlerCompletionHandler]: Creates the collection of closures an update task uses to notify your app of its progress.
//
// See: https://developer.apple.com/documentation/CoreML/MLUpdateProgressHandlers
type IMLUpdateProgressHandlers interface {
	objectivec.IObject

	// Topic: Creating progress handlers

	// Creates the collection of closures an update task uses to notify your app of its progress.
	InitForEventsProgressHandlerCompletionHandler(interestedEvents MLUpdateProgressEvent, progressHandler MLUpdateContextHandler, completionHandler MLUpdateContextHandler) MLUpdateProgressHandlers
}

// Init initializes the instance.
func (u MLUpdateProgressHandlers) Init() MLUpdateProgressHandlers {
	rv := objc.Send[MLUpdateProgressHandlers](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u MLUpdateProgressHandlers) Autorelease() MLUpdateProgressHandlers {
	rv := objc.Send[MLUpdateProgressHandlers](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLUpdateProgressHandlers creates a new MLUpdateProgressHandlers instance.
func NewMLUpdateProgressHandlers() MLUpdateProgressHandlers {
	class := getMLUpdateProgressHandlersClass()
	rv := objc.Send[MLUpdateProgressHandlers](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates the collection of closures an update task uses to notify your app
// of its progress.
//
// interestedEvents: The events for which the update task will call your closures for, contained
// in an option set.
//
// progressHandler: The closure an update task uses to notify your app. The update task only
// uses this closure for the events you specified in `interestedEvents`.
//
// completionHandler: The closure that an update tasks uses to notify you when it is complete.
//
// See: https://developer.apple.com/documentation/CoreML/MLUpdateProgressHandlers/init(forEvents:progressHandler:completionHandler:)
func (u MLUpdateProgressHandlers) InitForEventsProgressHandlerCompletionHandler(interestedEvents MLUpdateProgressEvent, progressHandler MLUpdateContextHandler, completionHandler MLUpdateContextHandler) MLUpdateProgressHandlers {
	_block1, _ := NewMLUpdateContextBlock(progressHandler)
	_block2, _ := NewMLUpdateContextBlock(completionHandler)
	rv := objc.Send[objc.ID](u.ID, objc.Sel("initForEvents:progressHandler:completionHandler:"), interestedEvents, _block1, _block2)
	return MLUpdateProgressHandlersFromID(rv)
}

// InitForEventsProgressHandler is a synchronous wrapper around [MLUpdateProgressHandlers.InitForEventsProgressHandlerCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (u MLUpdateProgressHandlers) InitForEventsProgressHandler(ctx context.Context, interestedEvents MLUpdateProgressEvent, progressHandler MLUpdateContextHandler) (*MLUpdateContext, error) {
	done := make(chan *MLUpdateContext, 1)
	u.InitForEventsProgressHandlerCompletionHandler(interestedEvents, progressHandler, func(val *MLUpdateContext) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
