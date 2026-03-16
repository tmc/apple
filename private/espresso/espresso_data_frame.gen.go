// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [EspressoDataFrame] class.
var (
	_EspressoDataFrameClass     EspressoDataFrameClass
	_EspressoDataFrameClassOnce sync.Once
)

func getEspressoDataFrameClass() EspressoDataFrameClass {
	_EspressoDataFrameClassOnce.Do(func() {
		_EspressoDataFrameClass = EspressoDataFrameClass{class: objc.GetClass("EspressoDataFrame")}
	})
	return _EspressoDataFrameClass
}

// GetEspressoDataFrameClass returns the class object for EspressoDataFrame.
func GetEspressoDataFrameClass() EspressoDataFrameClass {
	return getEspressoDataFrameClass()
}

type EspressoDataFrameClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoDataFrameClass) Alloc() EspressoDataFrame {
	rv := objc.Send[EspressoDataFrame](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [EspressoDataFrame.Function_name]
//   - [EspressoDataFrame.SetFunction_name]
//   - [EspressoDataFrame.GetFunctionName]
//   - [EspressoDataFrame.GetGroundTruthAttachment]
//   - [EspressoDataFrame.GetInputAttachment]
//   - [EspressoDataFrame.GetOutputAttachment]
//   - [EspressoDataFrame.GroundTruthAttachmentNames]
//   - [EspressoDataFrame.GroundTruthAttachments]
//   - [EspressoDataFrame.SetGroundTruthAttachments]
//   - [EspressoDataFrame.InputAttachmentNames]
//   - [EspressoDataFrame.InputAttachments]
//   - [EspressoDataFrame.SetInputAttachments]
//   - [EspressoDataFrame.LoadFromDictFrameStorage]
//   - [EspressoDataFrame.OutputAttachmentNames]
//   - [EspressoDataFrame.OutputAttachments]
//   - [EspressoDataFrame.SetOutputAttachments]
// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrame
type EspressoDataFrame struct {
	objectivec.Object
}

// EspressoDataFrameFromID constructs a [EspressoDataFrame] from an objc.ID.
func EspressoDataFrameFromID(id objc.ID) EspressoDataFrame {
	return EspressoDataFrame{objectivec.Object{id}}
}
// Ensure EspressoDataFrame implements IEspressoDataFrame.
var _ IEspressoDataFrame = EspressoDataFrame{}

// An interface definition for the [EspressoDataFrame] class.
//
// # Methods
//
//   - [IEspressoDataFrame.Function_name]
//   - [IEspressoDataFrame.SetFunction_name]
//   - [IEspressoDataFrame.GetFunctionName]
//   - [IEspressoDataFrame.GetGroundTruthAttachment]
//   - [IEspressoDataFrame.GetInputAttachment]
//   - [IEspressoDataFrame.GetOutputAttachment]
//   - [IEspressoDataFrame.GroundTruthAttachmentNames]
//   - [IEspressoDataFrame.GroundTruthAttachments]
//   - [IEspressoDataFrame.SetGroundTruthAttachments]
//   - [IEspressoDataFrame.InputAttachmentNames]
//   - [IEspressoDataFrame.InputAttachments]
//   - [IEspressoDataFrame.SetInputAttachments]
//   - [IEspressoDataFrame.LoadFromDictFrameStorage]
//   - [IEspressoDataFrame.OutputAttachmentNames]
//   - [IEspressoDataFrame.OutputAttachments]
//   - [IEspressoDataFrame.SetOutputAttachments]
//
// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrame
type IEspressoDataFrame interface {
	objectivec.IObject

	// Topic: Methods

	Function_name() string
	SetFunction_name(value string)
	GetFunctionName() objectivec.IObject
	GetGroundTruthAttachment(attachment objectivec.IObject) objectivec.IObject
	GetInputAttachment(attachment objectivec.IObject) objectivec.IObject
	GetOutputAttachment(attachment objectivec.IObject) objectivec.IObject
	GroundTruthAttachmentNames() foundation.INSArray
	GroundTruthAttachments() foundation.INSDictionary
	SetGroundTruthAttachments(value foundation.INSDictionary)
	InputAttachmentNames() foundation.INSArray
	InputAttachments() foundation.INSDictionary
	SetInputAttachments(value foundation.INSDictionary)
	LoadFromDictFrameStorage(dict objectivec.IObject, storage objectivec.IObject)
	OutputAttachmentNames() foundation.INSArray
	OutputAttachments() foundation.INSDictionary
	SetOutputAttachments(value foundation.INSDictionary)
}

// Init initializes the instance.
func (e EspressoDataFrame) Init() EspressoDataFrame {
	rv := objc.Send[EspressoDataFrame](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoDataFrame) Autorelease() EspressoDataFrame {
	rv := objc.Send[EspressoDataFrame](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoDataFrame creates a new EspressoDataFrame instance.
func NewEspressoDataFrame() EspressoDataFrame {
	class := getEspressoDataFrameClass()
	rv := objc.Send[EspressoDataFrame](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrame/getFunctionName
func (e EspressoDataFrame) GetFunctionName() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("getFunctionName"))
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrame/getGroundTruthAttachment:
func (e EspressoDataFrame) GetGroundTruthAttachment(attachment objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("getGroundTruthAttachment:"), attachment)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrame/getInputAttachment:
func (e EspressoDataFrame) GetInputAttachment(attachment objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("getInputAttachment:"), attachment)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrame/getOutputAttachment:
func (e EspressoDataFrame) GetOutputAttachment(attachment objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("getOutputAttachment:"), attachment)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrame/loadFromDict:frameStorage:
func (e EspressoDataFrame) LoadFromDictFrameStorage(dict objectivec.IObject, storage objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("loadFromDict:frameStorage:"), dict, storage)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrame/function_name
func (e EspressoDataFrame) Function_name() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("function_name"))
	return foundation.NSStringFromID(rv).String()
}
func (e EspressoDataFrame) SetFunction_name(value string) {
	objc.Send[struct{}](e.ID, objc.Sel("setFunction_name:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrame/groundTruthAttachmentNames
func (e EspressoDataFrame) GroundTruthAttachmentNames() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("groundTruthAttachmentNames"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrame/groundTruthAttachments
func (e EspressoDataFrame) GroundTruthAttachments() foundation.INSDictionary {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("groundTruthAttachments"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (e EspressoDataFrame) SetGroundTruthAttachments(value foundation.INSDictionary) {
	objc.Send[struct{}](e.ID, objc.Sel("setGroundTruthAttachments:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrame/inputAttachmentNames
func (e EspressoDataFrame) InputAttachmentNames() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("inputAttachmentNames"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrame/inputAttachments
func (e EspressoDataFrame) InputAttachments() foundation.INSDictionary {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("inputAttachments"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (e EspressoDataFrame) SetInputAttachments(value foundation.INSDictionary) {
	objc.Send[struct{}](e.ID, objc.Sel("setInputAttachments:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrame/outputAttachmentNames
func (e EspressoDataFrame) OutputAttachmentNames() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("outputAttachmentNames"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrame/outputAttachments
func (e EspressoDataFrame) OutputAttachments() foundation.INSDictionary {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("outputAttachments"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (e EspressoDataFrame) SetOutputAttachments(value foundation.INSDictionary) {
	objc.Send[struct{}](e.ID, objc.Sel("setOutputAttachments:"), value)
}

