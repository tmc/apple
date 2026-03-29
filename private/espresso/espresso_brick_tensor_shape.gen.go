// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [EspressoBrickTensorShape] class.
var (
	_EspressoBrickTensorShapeClass     EspressoBrickTensorShapeClass
	_EspressoBrickTensorShapeClassOnce sync.Once
)

func getEspressoBrickTensorShapeClass() EspressoBrickTensorShapeClass {
	_EspressoBrickTensorShapeClassOnce.Do(func() {
		_EspressoBrickTensorShapeClass = EspressoBrickTensorShapeClass{class: objc.GetClass("EspressoBrickTensorShape")}
	})
	return _EspressoBrickTensorShapeClass
}

// GetEspressoBrickTensorShapeClass returns the class object for EspressoBrickTensorShape.
func GetEspressoBrickTensorShapeClass() EspressoBrickTensorShapeClass {
	return getEspressoBrickTensorShapeClass()
}

type EspressoBrickTensorShapeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoBrickTensorShapeClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoBrickTensorShapeClass) Alloc() EspressoBrickTensorShape {
	rv := objc.Send[EspressoBrickTensorShape](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [EspressoBrickTensorShape.Batch]
//   - [EspressoBrickTensorShape.SetBatch]
//   - [EspressoBrickTensorShape.Channels]
//   - [EspressoBrickTensorShape.SetChannels]
//   - [EspressoBrickTensorShape.Height]
//   - [EspressoBrickTensorShape.SetHeight]
//   - [EspressoBrickTensorShape.Rank]
//   - [EspressoBrickTensorShape.SetRank]
//   - [EspressoBrickTensorShape.Sequence]
//   - [EspressoBrickTensorShape.SetSequence]
//   - [EspressoBrickTensorShape.Width]
//   - [EspressoBrickTensorShape.SetWidth]
// See: https://developer.apple.com/documentation/Espresso/EspressoBrickTensorShape
type EspressoBrickTensorShape struct {
	objectivec.Object
}

// EspressoBrickTensorShapeFromID constructs a [EspressoBrickTensorShape] from an objc.ID.
func EspressoBrickTensorShapeFromID(id objc.ID) EspressoBrickTensorShape {
	return EspressoBrickTensorShape{objectivec.Object{ID: id}}
}
// Ensure EspressoBrickTensorShape implements IEspressoBrickTensorShape.
var _ IEspressoBrickTensorShape = EspressoBrickTensorShape{}

// An interface definition for the [EspressoBrickTensorShape] class.
//
// # Methods
//
//   - [IEspressoBrickTensorShape.Batch]
//   - [IEspressoBrickTensorShape.SetBatch]
//   - [IEspressoBrickTensorShape.Channels]
//   - [IEspressoBrickTensorShape.SetChannels]
//   - [IEspressoBrickTensorShape.Height]
//   - [IEspressoBrickTensorShape.SetHeight]
//   - [IEspressoBrickTensorShape.Rank]
//   - [IEspressoBrickTensorShape.SetRank]
//   - [IEspressoBrickTensorShape.Sequence]
//   - [IEspressoBrickTensorShape.SetSequence]
//   - [IEspressoBrickTensorShape.Width]
//   - [IEspressoBrickTensorShape.SetWidth]
//
// See: https://developer.apple.com/documentation/Espresso/EspressoBrickTensorShape
type IEspressoBrickTensorShape interface {
	objectivec.IObject

	// Topic: Methods

	Batch() int
	SetBatch(value int)
	Channels() int
	SetChannels(value int)
	Height() int
	SetHeight(value int)
	Rank() int
	SetRank(value int)
	Sequence() int
	SetSequence(value int)
	Width() int
	SetWidth(value int)
}

// Init initializes the instance.
func (e EspressoBrickTensorShape) Init() EspressoBrickTensorShape {
	rv := objc.Send[EspressoBrickTensorShape](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoBrickTensorShape) Autorelease() EspressoBrickTensorShape {
	rv := objc.Send[EspressoBrickTensorShape](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoBrickTensorShape creates a new EspressoBrickTensorShape instance.
func NewEspressoBrickTensorShape() EspressoBrickTensorShape {
	class := getEspressoBrickTensorShapeClass()
	rv := objc.Send[EspressoBrickTensorShape](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoBrickTensorShape/batch
func (e EspressoBrickTensorShape) Batch() int {
	rv := objc.Send[int](e.ID, objc.Sel("batch"))
	return rv
}
func (e EspressoBrickTensorShape) SetBatch(value int) {
	objc.Send[struct{}](e.ID, objc.Sel("setBatch:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/EspressoBrickTensorShape/channels
func (e EspressoBrickTensorShape) Channels() int {
	rv := objc.Send[int](e.ID, objc.Sel("channels"))
	return rv
}
func (e EspressoBrickTensorShape) SetChannels(value int) {
	objc.Send[struct{}](e.ID, objc.Sel("setChannels:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/EspressoBrickTensorShape/height
func (e EspressoBrickTensorShape) Height() int {
	rv := objc.Send[int](e.ID, objc.Sel("height"))
	return rv
}
func (e EspressoBrickTensorShape) SetHeight(value int) {
	objc.Send[struct{}](e.ID, objc.Sel("setHeight:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/EspressoBrickTensorShape/rank
func (e EspressoBrickTensorShape) Rank() int {
	rv := objc.Send[int](e.ID, objc.Sel("rank"))
	return rv
}
func (e EspressoBrickTensorShape) SetRank(value int) {
	objc.Send[struct{}](e.ID, objc.Sel("setRank:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/EspressoBrickTensorShape/sequence
func (e EspressoBrickTensorShape) Sequence() int {
	rv := objc.Send[int](e.ID, objc.Sel("sequence"))
	return rv
}
func (e EspressoBrickTensorShape) SetSequence(value int) {
	objc.Send[struct{}](e.ID, objc.Sel("setSequence:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/EspressoBrickTensorShape/width
func (e EspressoBrickTensorShape) Width() int {
	rv := objc.Send[int](e.ID, objc.Sel("width"))
	return rv
}
func (e EspressoBrickTensorShape) SetWidth(value int) {
	objc.Send[struct{}](e.ID, objc.Sel("setWidth:"), value)
}

