// Code generated from Apple documentation for mlruntime. DO NOT EDIT.

package mlruntime

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLRTrialTask] class.
var (
	_MLRTrialTaskClass     MLRTrialTaskClass
	_MLRTrialTaskClassOnce sync.Once
)

func getMLRTrialTaskClass() MLRTrialTaskClass {
	_MLRTrialTaskClassOnce.Do(func() {
		_MLRTrialTaskClass = MLRTrialTaskClass{class: objc.GetClass("MLRTrialTask")}
	})
	return _MLRTrialTaskClass
}

// GetMLRTrialTaskClass returns the class object for MLRTrialTask.
func GetMLRTrialTaskClass() MLRTrialTaskClass {
	return getMLRTrialTaskClass()
}

type MLRTrialTaskClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLRTrialTaskClass) Alloc() MLRTrialTask {
	rv := objc.Send[MLRTrialTask](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MLRTrialTask.TriClient]
//   - [MLRTrialTask.InitWithTriClient]
// See: https://developer.apple.com/documentation/MLRuntime/MLRTrialTask
type MLRTrialTask struct {
	objectivec.Object
}

// MLRTrialTaskFromID constructs a [MLRTrialTask] from an objc.ID.
func MLRTrialTaskFromID(id objc.ID) MLRTrialTask {
	return MLRTrialTask{objectivec.Object{id}}
}
// Ensure MLRTrialTask implements IMLRTrialTask.
var _ IMLRTrialTask = MLRTrialTask{}





// An interface definition for the [MLRTrialTask] class.
//
// # Methods
//
//   - [IMLRTrialTask.TriClient]
//   - [IMLRTrialTask.InitWithTriClient]
//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTrialTask
type IMLRTrialTask interface {
	objectivec.IObject

	// Topic: Methods

	TriClient() unsafe.Pointer
	InitWithTriClient(client objectivec.IObject) MLRTrialTask
}





// Init initializes the instance.
func (r MLRTrialTask) Init() MLRTrialTask {
	rv := objc.Send[MLRTrialTask](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MLRTrialTask) Autorelease() MLRTrialTask {
	rv := objc.Send[MLRTrialTask](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLRTrialTask creates a new MLRTrialTask instance.
func NewMLRTrialTask() MLRTrialTask {
	class := getMLRTrialTaskClass()
	rv := objc.Send[MLRTrialTask](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTrialTask/initWithTriClient:
func NewRTrialTaskWithTriClient(client objectivec.IObject) MLRTrialTask {
	instance := getMLRTrialTaskClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTriClient:"), client)
	return MLRTrialTaskFromID(rv)
}







//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTrialTask/initWithTriClient:
func (r MLRTrialTask) InitWithTriClient(client objectivec.IObject) MLRTrialTask {
	rv := objc.Send[MLRTrialTask](r.ID, objc.Sel("initWithTriClient:"), client)
	return rv
}












// See: https://developer.apple.com/documentation/MLRuntime/MLRTrialTask/triClient
func (r MLRTrialTask) TriClient() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r.ID, objc.Sel("triClient"))
	return rv
}

















