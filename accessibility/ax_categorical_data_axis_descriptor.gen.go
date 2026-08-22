// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AXCategoricalDataAxisDescriptor] class.
var (
	_AXCategoricalDataAxisDescriptorClass     AXCategoricalDataAxisDescriptorClass
	_AXCategoricalDataAxisDescriptorClassOnce sync.Once
)

func getAXCategoricalDataAxisDescriptorClass() AXCategoricalDataAxisDescriptorClass {
	_AXCategoricalDataAxisDescriptorClassOnce.Do(func() {
		_AXCategoricalDataAxisDescriptorClass = AXCategoricalDataAxisDescriptorClass{class: objc.GetClass("AXCategoricalDataAxisDescriptor")}
	})
	return _AXCategoricalDataAxisDescriptorClass
}

// GetAXCategoricalDataAxisDescriptorClass returns the class object for AXCategoricalDataAxisDescriptor.
func GetAXCategoricalDataAxisDescriptorClass() AXCategoricalDataAxisDescriptorClass {
	return getAXCategoricalDataAxisDescriptorClass()
}

type AXCategoricalDataAxisDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AXCategoricalDataAxisDescriptorClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AXCategoricalDataAxisDescriptorClass) Alloc() AXCategoricalDataAxisDescriptor {
	rv := objc.Send[AXCategoricalDataAxisDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents an axis of categorical data.
//
// # Overview
//
// A categorical data axis divides information into groups, or categories. For
// example, a categorical axis may represent blood type data divided into the
// possible categories AB, A, B, and O.
//
// # Creating a categorical data axis
//
//   - [AXCategoricalDataAxisDescriptor.InitWithTitleCategoryOrder]: Creates a categorical data axis with the specified title and an array of categories in the specified order.
//   - [AXCategoricalDataAxisDescriptor.InitWithAttributedTitleCategoryOrder]: Creates a categorical data axis with the specified attributed title and an array of categories in the specified order.
//
// # Configuring the order of categories
//
//   - [AXCategoricalDataAxisDescriptor.CategoryOrder]: A list of every category value for the axis in the order they appear visually in the graph or legend.
//   - [AXCategoricalDataAxisDescriptor.SetCategoryOrder]
//
// See: https://developer.apple.com/documentation/Accessibility/AXCategoricalDataAxisDescriptor
type AXCategoricalDataAxisDescriptor struct {
	objectivec.Object
}

// AXCategoricalDataAxisDescriptorFromID constructs a [AXCategoricalDataAxisDescriptor] from an objc.ID.
//
// An object that represents an axis of categorical data.
func AXCategoricalDataAxisDescriptorFromID(id objc.ID) AXCategoricalDataAxisDescriptor {
	return AXCategoricalDataAxisDescriptor{objectivec.Object{ID: id}}
}

// NOTE: AXCategoricalDataAxisDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AXCategoricalDataAxisDescriptor] class.
//
// # Creating a categorical data axis
//
//   - [IAXCategoricalDataAxisDescriptor.InitWithTitleCategoryOrder]: Creates a categorical data axis with the specified title and an array of categories in the specified order.
//   - [IAXCategoricalDataAxisDescriptor.InitWithAttributedTitleCategoryOrder]: Creates a categorical data axis with the specified attributed title and an array of categories in the specified order.
//
// # Configuring the order of categories
//
//   - [IAXCategoricalDataAxisDescriptor.CategoryOrder]: A list of every category value for the axis in the order they appear visually in the graph or legend.
//   - [IAXCategoricalDataAxisDescriptor.SetCategoryOrder]
//
// See: https://developer.apple.com/documentation/Accessibility/AXCategoricalDataAxisDescriptor
type IAXCategoricalDataAxisDescriptor interface {
	objectivec.IObject

	// Topic: Creating a categorical data axis

	// Creates a categorical data axis with the specified title and an array of categories in the specified order.
	InitWithTitleCategoryOrder(title string, categoryOrder []string) AXCategoricalDataAxisDescriptor
	// Creates a categorical data axis with the specified attributed title and an array of categories in the specified order.
	InitWithAttributedTitleCategoryOrder(attributedTitle foundation.NSAttributedString, categoryOrder []string) AXCategoricalDataAxisDescriptor

	// Topic: Configuring the order of categories

	// A list of every category value for the axis in the order they appear visually in the graph or legend.
	CategoryOrder() []string
	SetCategoryOrder(value []string)

	// An attributed version of the axis title.
	AttributedTitle() foundation.NSAttributedString
	// The title of the axis.
	Title() string
}

// Init initializes the instance.
func (a AXCategoricalDataAxisDescriptor) Init() AXCategoricalDataAxisDescriptor {
	rv := objc.Send[AXCategoricalDataAxisDescriptor](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AXCategoricalDataAxisDescriptor) Autorelease() AXCategoricalDataAxisDescriptor {
	rv := objc.Send[AXCategoricalDataAxisDescriptor](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXCategoricalDataAxisDescriptor creates a new AXCategoricalDataAxisDescriptor instance.
func NewAXCategoricalDataAxisDescriptor() AXCategoricalDataAxisDescriptor {
	class := getAXCategoricalDataAxisDescriptorClass()
	rv := objc.Send[AXCategoricalDataAxisDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a categorical data axis with the specified attributed title and an
// array of categories in the specified order.
//
// See: https://developer.apple.com/documentation/Accessibility/AXCategoricalDataAxisDescriptor/init(attributedTitle:categoryOrder:)
func NewAXCategoricalDataAxisDescriptorWithAttributedTitleCategoryOrder(attributedTitle foundation.NSAttributedString, categoryOrder []string) AXCategoricalDataAxisDescriptor {
	instance := getAXCategoricalDataAxisDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAttributedTitle:categoryOrder:"), attributedTitle, objectivec.StringSliceToNSArray(categoryOrder))
	return AXCategoricalDataAxisDescriptorFromID(rv)
}

// Creates a categorical data axis with the specified title and an array of
// categories in the specified order.
//
// See: https://developer.apple.com/documentation/Accessibility/AXCategoricalDataAxisDescriptor/init(title:categoryOrder:)
func NewAXCategoricalDataAxisDescriptorWithTitleCategoryOrder(title string, categoryOrder []string) AXCategoricalDataAxisDescriptor {
	instance := getAXCategoricalDataAxisDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTitle:categoryOrder:"), objc.String(title), objectivec.StringSliceToNSArray(categoryOrder))
	return AXCategoricalDataAxisDescriptorFromID(rv)
}

// Creates a categorical data axis with the specified title and an array of
// categories in the specified order.
//
// See: https://developer.apple.com/documentation/Accessibility/AXCategoricalDataAxisDescriptor/init(title:categoryOrder:)
func (a AXCategoricalDataAxisDescriptor) InitWithTitleCategoryOrder(title string, categoryOrder []string) AXCategoricalDataAxisDescriptor {
	rv := objc.Send[AXCategoricalDataAxisDescriptor](a.ID, objc.Sel("initWithTitle:categoryOrder:"), objc.String(title), objectivec.StringSliceToNSArray(categoryOrder))
	return rv
}

// Creates a categorical data axis with the specified attributed title and an
// array of categories in the specified order.
//
// See: https://developer.apple.com/documentation/Accessibility/AXCategoricalDataAxisDescriptor/init(attributedTitle:categoryOrder:)
func (a AXCategoricalDataAxisDescriptor) InitWithAttributedTitleCategoryOrder(attributedTitle foundation.NSAttributedString, categoryOrder []string) AXCategoricalDataAxisDescriptor {
	rv := objc.Send[AXCategoricalDataAxisDescriptor](a.ID, objc.Sel("initWithAttributedTitle:categoryOrder:"), attributedTitle, objectivec.StringSliceToNSArray(categoryOrder))
	return rv
}

// An attributed version of the axis title.
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataAxisDescriptor/attributedTitle
func (a AXCategoricalDataAxisDescriptor) AttributedTitle() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("attributedTitle"))
	return foundation.NSAttributedStringFromID(rv)
}

// The title of the axis.
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataAxisDescriptor/title
func (a AXCategoricalDataAxisDescriptor) Title() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("title"))
	return foundation.NSStringFromID(rv).String()
}

// A list of every category value for the axis in the order they appear
// visually in the graph or legend.
//
// # Discussion
//
// If your categorical axis represents, for example, blood type data, and the
// legend lists AB, A, B, O in that order, provide an array that contains the
// strings `"AB"`, `"A"`, `"B"`, and `"O"` in the same order.
//
// See: https://developer.apple.com/documentation/Accessibility/AXCategoricalDataAxisDescriptor/categoryOrder
func (a AXCategoricalDataAxisDescriptor) CategoryOrder() []string {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("categoryOrder"))
	return objc.ConvertSliceToStrings(rv)
}
func (a AXCategoricalDataAxisDescriptor) SetCategoryOrder(value []string) {
	objc.Send[struct{}](a.ID, objc.Sel("setCategoryOrder:"), objectivec.StringSliceToNSArray(value))
}

// Protocol methods for AXDataAxisDescriptor

// The title of the axis.
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataAxisDescriptor/title
func (o AXCategoricalDataAxisDescriptor) SetTitle(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setTitle:"), objc.String(value))
}

// An attributed version of the axis title.
//
// # Discussion
//
// If you set the value of this property, the system uses this value instead
// of [Title].
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataAxisDescriptor/attributedTitle
func (o AXCategoricalDataAxisDescriptor) SetAttributedTitle(value foundation.NSAttributedString) {
	objc.Send[struct{}](o.ID, objc.Sel("setAttributedTitle:"), value)
}
