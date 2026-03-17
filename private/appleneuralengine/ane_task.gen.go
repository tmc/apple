// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANETask] class.
var (
	_ANETaskClass     ANETaskClass
	_ANETaskClassOnce sync.Once
)

func getANETaskClass() ANETaskClass {
	_ANETaskClassOnce.Do(func() {
		_ANETaskClass = ANETaskClass{class: objc.GetClass("_ANETask")}
	})
	return _ANETaskClass
}

// GetANETaskClass returns the class object for _ANETask.
func GetANETaskClass() ANETaskClass {
	return getANETaskClass()
}

type ANETaskClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANETaskClass) Alloc() ANETask {
	rv := objc.Send[ANETask](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [ANETask.ExecutionCriteria]
//   - [ANETask.Handler]
//   - [ANETask.Name]
//   - [ANETask.PeriodSeconds]
//   - [ANETask.Queue]
//   - [ANETask.InitWithNamePeriodHandler]
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANETask
type ANETask struct {
	objectivec.Object
}

// ANETaskFromID constructs a [ANETask] from an objc.ID.
func ANETaskFromID(id objc.ID) ANETask {
	return ANETask{objectivec.Object{ID: id}}
}
// Ensure ANETask implements IANETask.
var _ IANETask = ANETask{}

// An interface definition for the [ANETask] class.
//
// # Methods
//
//   - [IANETask.ExecutionCriteria]
//   - [IANETask.Handler]
//   - [IANETask.Name]
//   - [IANETask.PeriodSeconds]
//   - [IANETask.Queue]
//   - [IANETask.InitWithNamePeriodHandler]
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANETask
type IANETask interface {
	objectivec.IObject

	// Topic: Methods

	ExecutionCriteria() objectivec.Object
	Handler() VoidHandler
	Name() string
	PeriodSeconds() uint64
	Queue() objectivec.Object
	InitWithNamePeriodHandler(name objectivec.IObject, period uint64, handler VoidHandler) ANETask
}

// Init initializes the instance.
func (a ANETask) Init() ANETask {
	rv := objc.Send[ANETask](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANETask) Autorelease() ANETask {
	rv := objc.Send[ANETask](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANETask creates a new ANETask instance.
func NewANETask() ANETask {
	class := getANETaskClass()
	rv := objc.Send[ANETask](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANETask/initWithName:period:handler:
func (a ANETask) InitWithNamePeriodHandler(name objectivec.IObject, period uint64, handler VoidHandler) ANETask {
_block2, _cleanup2 := NewVoidBlock(handler)
	defer _cleanup2()
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithName:period:handler:"), name, period, _block2)
	return ANETaskFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANETask/taskWithName:period:handler:
func (_ANETaskClass ANETaskClass) TaskWithNamePeriodHandler(name objectivec.IObject, period uint64, handler VoidHandler) objectivec.IObject {
_block2, _cleanup2 := NewVoidBlock(handler)
	defer _cleanup2()
	rv := objc.Send[objc.ID](objc.ID(_ANETaskClass.class), objc.Sel("taskWithName:period:handler:"), name, period, _block2)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANETask/executionCriteria
func (a ANETask) ExecutionCriteria() objectivec.Object {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("executionCriteria"))
	return objectivec.ObjectFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANETask/handler
func (a ANETask) Handler() VoidHandler {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("handler"))
	_ = rv
	return nil
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANETask/name
func (a ANETask) Name() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANETask/periodSeconds
func (a ANETask) PeriodSeconds() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("periodSeconds"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANETask/queue
func (a ANETask) Queue() objectivec.Object {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("queue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}

// InitWithNamePeriodHandlerSync is a synchronous wrapper around [ANETask.InitWithNamePeriodHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (a ANETask) InitWithNamePeriodHandlerSync(ctx context.Context, name objectivec.IObject, period uint64) error {
	done := make(chan struct{}, 1)
	a.InitWithNamePeriodHandler(name, period, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// TaskWithNamePeriodHandlerSync is a synchronous wrapper around [ANETask.TaskWithNamePeriodHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (ac ANETaskClass) TaskWithNamePeriodHandlerSync(ctx context.Context, name objectivec.IObject, period uint64) error {
	done := make(chan struct{}, 1)
	ac.TaskWithNamePeriodHandler(name, period, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

