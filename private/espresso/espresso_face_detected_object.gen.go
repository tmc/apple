// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [EspressoFaceDetectedObject] class.
var (
	_EspressoFaceDetectedObjectClass     EspressoFaceDetectedObjectClass
	_EspressoFaceDetectedObjectClassOnce sync.Once
)

func getEspressoFaceDetectedObjectClass() EspressoFaceDetectedObjectClass {
	_EspressoFaceDetectedObjectClassOnce.Do(func() {
		_EspressoFaceDetectedObjectClass = EspressoFaceDetectedObjectClass{class: objc.GetClass("EspressoFaceDetectedObject")}
	})
	return _EspressoFaceDetectedObjectClass
}

// GetEspressoFaceDetectedObjectClass returns the class object for EspressoFaceDetectedObject.
func GetEspressoFaceDetectedObjectClass() EspressoFaceDetectedObjectClass {
	return getEspressoFaceDetectedObjectClass()
}

type EspressoFaceDetectedObjectClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoFaceDetectedObjectClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoFaceDetectedObjectClass) Alloc() EspressoFaceDetectedObject {
	rv := objc.Send[EspressoFaceDetectedObject](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [EspressoFaceDetectedObject.Bounds]
//   - [EspressoFaceDetectedObject.SetBounds]
//   - [EspressoFaceDetectedObject.Center]
//   - [EspressoFaceDetectedObject.SetCenter]
//   - [EspressoFaceDetectedObject.Confidence]
//   - [EspressoFaceDetectedObject.SetConfidence]
//   - [EspressoFaceDetectedObject.ObjectType]
//   - [EspressoFaceDetectedObject.SetObjectType]
//   - [EspressoFaceDetectedObject.InitWithOptionsXlocYlocSizeConfidence]
//   - [EspressoFaceDetectedObject.DebugDescription]
//   - [EspressoFaceDetectedObject.Description]
//   - [EspressoFaceDetectedObject.Hash]
//   - [EspressoFaceDetectedObject.Superclass]
// See: https://developer.apple.com/documentation/Espresso/EspressoFaceDetectedObject
type EspressoFaceDetectedObject struct {
	objectivec.Object
}

// EspressoFaceDetectedObjectFromID constructs a [EspressoFaceDetectedObject] from an objc.ID.
func EspressoFaceDetectedObjectFromID(id objc.ID) EspressoFaceDetectedObject {
	return EspressoFaceDetectedObject{objectivec.Object{ID: id}}
}
// Ensure EspressoFaceDetectedObject implements IEspressoFaceDetectedObject.
var _ IEspressoFaceDetectedObject = EspressoFaceDetectedObject{}

// An interface definition for the [EspressoFaceDetectedObject] class.
//
// # Methods
//
//   - [IEspressoFaceDetectedObject.Bounds]
//   - [IEspressoFaceDetectedObject.SetBounds]
//   - [IEspressoFaceDetectedObject.Center]
//   - [IEspressoFaceDetectedObject.SetCenter]
//   - [IEspressoFaceDetectedObject.Confidence]
//   - [IEspressoFaceDetectedObject.SetConfidence]
//   - [IEspressoFaceDetectedObject.ObjectType]
//   - [IEspressoFaceDetectedObject.SetObjectType]
//   - [IEspressoFaceDetectedObject.InitWithOptionsXlocYlocSizeConfidence]
//   - [IEspressoFaceDetectedObject.DebugDescription]
//   - [IEspressoFaceDetectedObject.Description]
//   - [IEspressoFaceDetectedObject.Hash]
//   - [IEspressoFaceDetectedObject.Superclass]
//
// See: https://developer.apple.com/documentation/Espresso/EspressoFaceDetectedObject
type IEspressoFaceDetectedObject interface {
	objectivec.IObject

	// Topic: Methods

	Bounds() corefoundation.CGRect
	SetBounds(value corefoundation.CGRect)
	Center() corefoundation.CGPoint
	SetCenter(value corefoundation.CGPoint)
	Confidence() float32
	SetConfidence(value float32)
	ObjectType() int64
	SetObjectType(value int64)
	InitWithOptionsXlocYlocSizeConfidence(xloc float32, yloc float32, size float32, confidence float32) EspressoFaceDetectedObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (e EspressoFaceDetectedObject) Init() EspressoFaceDetectedObject {
	rv := objc.Send[EspressoFaceDetectedObject](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoFaceDetectedObject) Autorelease() EspressoFaceDetectedObject {
	rv := objc.Send[EspressoFaceDetectedObject](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoFaceDetectedObject creates a new EspressoFaceDetectedObject instance.
func NewEspressoFaceDetectedObject() EspressoFaceDetectedObject {
	class := getEspressoFaceDetectedObjectClass()
	rv := objc.Send[EspressoFaceDetectedObject](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoFaceDetectedObject/initWithOptionsXloc:yloc:size:confidence:
func NewEspressoFaceDetectedObjectWithOptionsXlocYlocSizeConfidence(xloc float32, yloc float32, size float32, confidence float32) EspressoFaceDetectedObject {
	instance := getEspressoFaceDetectedObjectClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithOptionsXloc:yloc:size:confidence:"), xloc, yloc, size, confidence)
	return EspressoFaceDetectedObjectFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoFaceDetectedObject/initWithOptionsXloc:yloc:size:confidence:
func (e EspressoFaceDetectedObject) InitWithOptionsXlocYlocSizeConfidence(xloc float32, yloc float32, size float32, confidence float32) EspressoFaceDetectedObject {
	rv := objc.Send[EspressoFaceDetectedObject](e.ID, objc.Sel("initWithOptionsXloc:yloc:size:confidence:"), xloc, yloc, size, confidence)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoFaceDetectedObject/bounds
func (e EspressoFaceDetectedObject) Bounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](e.ID, objc.Sel("bounds"))
	return corefoundation.CGRect(rv)
}
func (e EspressoFaceDetectedObject) SetBounds(value corefoundation.CGRect) {
	objc.Send[struct{}](e.ID, objc.Sel("setBounds:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/EspressoFaceDetectedObject/center
func (e EspressoFaceDetectedObject) Center() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](e.ID, objc.Sel("center"))
	return corefoundation.CGPoint(rv)
}
func (e EspressoFaceDetectedObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](e.ID, objc.Sel("setCenter:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/EspressoFaceDetectedObject/confidence
func (e EspressoFaceDetectedObject) Confidence() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("confidence"))
	return rv
}
func (e EspressoFaceDetectedObject) SetConfidence(value float32) {
	objc.Send[struct{}](e.ID, objc.Sel("setConfidence:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/EspressoFaceDetectedObject/debugDescription
func (e EspressoFaceDetectedObject) DebugDescription() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Espresso/EspressoFaceDetectedObject/description
func (e EspressoFaceDetectedObject) Description() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Espresso/EspressoFaceDetectedObject/hash
func (e EspressoFaceDetectedObject) Hash() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/Espresso/EspressoFaceDetectedObject/objectType
func (e EspressoFaceDetectedObject) ObjectType() int64 {
	rv := objc.Send[int64](e.ID, objc.Sel("objectType"))
	return rv
}
func (e EspressoFaceDetectedObject) SetObjectType(value int64) {
	objc.Send[struct{}](e.ID, objc.Sel("setObjectType:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/EspressoFaceDetectedObject/superclass
func (e EspressoFaceDetectedObject) Superclass() objc.Class {
	rv := objc.Send[objc.Class](e.ID, objc.Sel("superclass"))
	return rv
}

