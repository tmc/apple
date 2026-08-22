// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that declares the minimum interface necessary for an accessibility element to act as a chart.
//
// See: https://developer.apple.com/documentation/Accessibility/AXChart
type AXChart interface {
	objectivec.IObject

	// A semantic description of an accessible chart or graph in the form of a chart descriptor.
	//
	// See: https://developer.apple.com/documentation/Accessibility/AXChart/accessibilityChartDescriptor
	AccessibilityChartDescriptor() IAXChartDescriptor
	SetAccessibilityChartDescriptor(value IAXChartDescriptor)
}

// AXChartObject wraps an existing Objective-C object that conforms to the AXChart protocol.
type AXChartObject struct {
	objectivec.Object
}

func (o AXChartObject) BaseObject() objectivec.Object {
	return o.Object
}

// AXChartObjectFromID constructs a [AXChartObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AXChartObjectFromID(id objc.ID) AXChartObject {
	return AXChartObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A semantic description of an accessible chart or graph in the form of a
// chart descriptor.
//
// See: https://developer.apple.com/documentation/Accessibility/AXChart/accessibilityChartDescriptor
func (o AXChartObject) AccessibilityChartDescriptor() IAXChartDescriptor {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityChartDescriptor"))
	return AXChartDescriptorFromID(rv)
}

func (o AXChartObject) SetAccessibilityChartDescriptor(value IAXChartDescriptor) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityChartDescriptor:"), value)
}
