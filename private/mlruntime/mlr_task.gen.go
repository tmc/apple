// Code generated from Apple documentation for mlruntime. DO NOT EDIT.

package mlruntime

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLRTask] class.
var (
	_MLRTaskClass     MLRTaskClass
	_MLRTaskClassOnce sync.Once
)

func getMLRTaskClass() MLRTaskClass {
	_MLRTaskClassOnce.Do(func() {
		_MLRTaskClass = MLRTaskClass{class: objc.GetClass("MLRTask")}
	})
	return _MLRTaskClass
}

// GetMLRTaskClass returns the class object for MLRTask.
func GetMLRTaskClass() MLRTaskClass {
	return getMLRTaskClass()
}

type MLRTaskClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLRTaskClass) Alloc() MLRTask {
	rv := objc.Send[MLRTask](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLRTask.Attachments]
//   - [MLRTask.EncodeWithCoder]
//   - [MLRTask.Parameters]
//   - [MLRTask.InitWithCoder]
//   - [MLRTask.InitWithParametersAttachments]
//   - [MLRTask.InitWithParametersDict]
// See: https://developer.apple.com/documentation/MLRuntime/MLRTask
type MLRTask struct {
	objectivec.Object
}

// MLRTaskFromID constructs a [MLRTask] from an objc.ID.
func MLRTaskFromID(id objc.ID) MLRTask {
	return MLRTask{objectivec.Object{id}}
}
// Ensure MLRTask implements IMLRTask.
var _ IMLRTask = MLRTask{}

// An interface definition for the [MLRTask] class.
//
// # Methods
//
//   - [IMLRTask.Attachments]
//   - [IMLRTask.EncodeWithCoder]
//   - [IMLRTask.Parameters]
//   - [IMLRTask.InitWithCoder]
//   - [IMLRTask.InitWithParametersAttachments]
//   - [IMLRTask.InitWithParametersDict]
//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTask
type IMLRTask interface {
	objectivec.IObject

	// Topic: Methods

	Attachments() IMLRTaskAttachments
	EncodeWithCoder(coder objectivec.IObject)
	Parameters() IMLRTaskParameters
	InitWithCoder(coder objectivec.IObject) MLRTask
	InitWithParametersAttachments(parameters objectivec.IObject, attachments objectivec.IObject) MLRTask
	InitWithParametersDict(dict objectivec.IObject) MLRTask
}

// Init initializes the instance.
func (r MLRTask) Init() MLRTask {
	rv := objc.Send[MLRTask](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MLRTask) Autorelease() MLRTask {
	rv := objc.Send[MLRTask](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLRTask creates a new MLRTask instance.
func NewMLRTask() MLRTask {
	class := getMLRTaskClass()
	rv := objc.Send[MLRTask](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTask/initWithCoder:
func NewRTaskWithCoder(coder objectivec.IObject) MLRTask {
	instance := getMLRTaskClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLRTaskFromID(rv)
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTask/initWithParameters:attachments:
func NewRTaskWithParametersAttachments(parameters objectivec.IObject, attachments objectivec.IObject) MLRTask {
	instance := getMLRTaskClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameters:attachments:"), parameters, attachments)
	return MLRTaskFromID(rv)
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTask/initWithParametersDict:
func NewRTaskWithParametersDict(dict objectivec.IObject) MLRTask {
	instance := getMLRTaskClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParametersDict:"), dict)
	return MLRTaskFromID(rv)
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTask/encodeWithCoder:
func (r MLRTask) EncodeWithCoder(coder objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("encodeWithCoder:"), coder)
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTask/initWithCoder:
func (r MLRTask) InitWithCoder(coder objectivec.IObject) MLRTask {
	rv := objc.Send[MLRTask](r.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTask/initWithParameters:attachments:
func (r MLRTask) InitWithParametersAttachments(parameters objectivec.IObject, attachments objectivec.IObject) MLRTask {
	rv := objc.Send[MLRTask](r.ID, objc.Sel("initWithParameters:attachments:"), parameters, attachments)
	return rv
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTask/initWithParametersDict:
func (r MLRTask) InitWithParametersDict(dict objectivec.IObject) MLRTask {
	rv := objc.Send[MLRTask](r.ID, objc.Sel("initWithParametersDict:"), dict)
	return rv
}

// See: https://developer.apple.com/documentation/MLRuntime/MLRTask/supportsSecureCoding
func (_MLRTaskClass MLRTaskClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MLRTaskClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/MLRuntime/MLRTask/attachments
func (r MLRTask) Attachments() IMLRTaskAttachments {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("attachments"))
	return MLRTaskAttachmentsFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/MLRuntime/MLRTask/parameters
func (r MLRTask) Parameters() IMLRTaskParameters {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("parameters"))
	return MLRTaskParametersFromID(objc.ID(rv))
}

