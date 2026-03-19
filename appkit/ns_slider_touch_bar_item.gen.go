// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSSliderTouchBarItem] class.
var (
	_NSSliderTouchBarItemClass     NSSliderTouchBarItemClass
	_NSSliderTouchBarItemClassOnce sync.Once
)

func getNSSliderTouchBarItemClass() NSSliderTouchBarItemClass {
	_NSSliderTouchBarItemClassOnce.Do(func() {
		_NSSliderTouchBarItemClass = NSSliderTouchBarItemClass{class: objc.GetClass("NSSliderTouchBarItem")}
	})
	return _NSSliderTouchBarItemClass
}

// GetNSSliderTouchBarItemClass returns the class object for NSSliderTouchBarItem.
func GetNSSliderTouchBarItemClass() NSSliderTouchBarItemClass {
	return getNSSliderTouchBarItemClass()
}

type NSSliderTouchBarItemClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSliderTouchBarItemClass) Alloc() NSSliderTouchBarItem {
	rv := objc.Send[NSSliderTouchBarItem](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A bar item that provides a slider control for choosing a value in a range.
//
// # Configuring the slider
//
//   - [NSSliderTouchBarItem.Slider]: The slider displayed by the bar item.
//   - [NSSliderTouchBarItem.SetSlider]
//   - [NSSliderTouchBarItem.Label]: The text displayed alongside the slider.
//   - [NSSliderTouchBarItem.SetLabel]
//
// # Configuring slider accessories
//
//   - [NSSliderTouchBarItem.MinimumValueAccessory]: The accessory that appears at the end of the slider with the minimum value.
//   - [NSSliderTouchBarItem.SetMinimumValueAccessory]
//   - [NSSliderTouchBarItem.MaximumValueAccessory]: The accessory that appears at the end of the slider with the maximum value.
//   - [NSSliderTouchBarItem.SetMaximumValueAccessory]
//   - [NSSliderTouchBarItem.ValueAccessoryWidth]: The width of the value accessories that appear at either end of the slider.
//   - [NSSliderTouchBarItem.SetValueAccessoryWidth]
//
// # Managing the slider’s value
//
//   - [NSSliderTouchBarItem.MinimumSliderWidth]: The minimum width of the slider’s track.
//   - [NSSliderTouchBarItem.SetMinimumSliderWidth]
//   - [NSSliderTouchBarItem.MaximumSliderWidth]: The maximum width of the slider’s track.
//   - [NSSliderTouchBarItem.SetMaximumSliderWidth]
//   - [NSSliderTouchBarItem.DoubleValue]: The double value of the slider.
//   - [NSSliderTouchBarItem.SetDoubleValue]
//
// # Handling slider interaction
//
//   - [NSSliderTouchBarItem.Target]: An object that is notified when a user interacts with the slider or either of the accessories.
//   - [NSSliderTouchBarItem.SetTarget]
//   - [NSSliderTouchBarItem.Action]: The selector on the target object that is invoked when a user interacts with the slider or either of the accessories.
//   - [NSSliderTouchBarItem.SetAction]
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem
type NSSliderTouchBarItem struct {
	NSTouchBarItem
}

// NSSliderTouchBarItemFromID constructs a [NSSliderTouchBarItem] from an objc.ID.
//
// A bar item that provides a slider control for choosing a value in a range.
func NSSliderTouchBarItemFromID(id objc.ID) NSSliderTouchBarItem {
	return NSSliderTouchBarItem{NSTouchBarItem: NSTouchBarItemFromID(id)}
}
// NOTE: NSSliderTouchBarItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSSliderTouchBarItem] class.
//
// # Configuring the slider
//
//   - [INSSliderTouchBarItem.Slider]: The slider displayed by the bar item.
//   - [INSSliderTouchBarItem.SetSlider]
//   - [INSSliderTouchBarItem.Label]: The text displayed alongside the slider.
//   - [INSSliderTouchBarItem.SetLabel]
//
// # Configuring slider accessories
//
//   - [INSSliderTouchBarItem.MinimumValueAccessory]: The accessory that appears at the end of the slider with the minimum value.
//   - [INSSliderTouchBarItem.SetMinimumValueAccessory]
//   - [INSSliderTouchBarItem.MaximumValueAccessory]: The accessory that appears at the end of the slider with the maximum value.
//   - [INSSliderTouchBarItem.SetMaximumValueAccessory]
//   - [INSSliderTouchBarItem.ValueAccessoryWidth]: The width of the value accessories that appear at either end of the slider.
//   - [INSSliderTouchBarItem.SetValueAccessoryWidth]
//
// # Managing the slider’s value
//
//   - [INSSliderTouchBarItem.MinimumSliderWidth]: The minimum width of the slider’s track.
//   - [INSSliderTouchBarItem.SetMinimumSliderWidth]
//   - [INSSliderTouchBarItem.MaximumSliderWidth]: The maximum width of the slider’s track.
//   - [INSSliderTouchBarItem.SetMaximumSliderWidth]
//   - [INSSliderTouchBarItem.DoubleValue]: The double value of the slider.
//   - [INSSliderTouchBarItem.SetDoubleValue]
//
// # Handling slider interaction
//
//   - [INSSliderTouchBarItem.Target]: An object that is notified when a user interacts with the slider or either of the accessories.
//   - [INSSliderTouchBarItem.SetTarget]
//   - [INSSliderTouchBarItem.Action]: The selector on the target object that is invoked when a user interacts with the slider or either of the accessories.
//   - [INSSliderTouchBarItem.SetAction]
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem
type INSSliderTouchBarItem interface {
	INSTouchBarItem

	// Topic: Configuring the slider

	// The slider displayed by the bar item.
	Slider() INSSlider
	SetSlider(value INSSlider)
	// The text displayed alongside the slider.
	Label() string
	SetLabel(value string)

	// Topic: Configuring slider accessories

	// The accessory that appears at the end of the slider with the minimum value.
	MinimumValueAccessory() INSSliderAccessory
	SetMinimumValueAccessory(value INSSliderAccessory)
	// The accessory that appears at the end of the slider with the maximum value.
	MaximumValueAccessory() INSSliderAccessory
	SetMaximumValueAccessory(value INSSliderAccessory)
	// The width of the value accessories that appear at either end of the slider.
	ValueAccessoryWidth() NSSliderAccessoryWidth
	SetValueAccessoryWidth(value NSSliderAccessoryWidth)

	// Topic: Managing the slider’s value

	// The minimum width of the slider’s track.
	MinimumSliderWidth() float64
	SetMinimumSliderWidth(value float64)
	// The maximum width of the slider’s track.
	MaximumSliderWidth() float64
	SetMaximumSliderWidth(value float64)
	// The double value of the slider.
	DoubleValue() float64
	SetDoubleValue(value float64)

	// Topic: Handling slider interaction

	// An object that is notified when a user interacts with the slider or either of the accessories.
	Target() objectivec.IObject
	SetTarget(value objectivec.IObject)
	// The selector on the target object that is invoked when a user interacts with the slider or either of the accessories.
	Action() objc.SEL
	SetAction(value objc.SEL)
}

// Init initializes the instance.
func (s NSSliderTouchBarItem) Init() NSSliderTouchBarItem {
	rv := objc.Send[NSSliderTouchBarItem](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSliderTouchBarItem) Autorelease() NSSliderTouchBarItem {
	rv := objc.Send[NSSliderTouchBarItem](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSliderTouchBarItem creates a new NSSliderTouchBarItem instance.
func NewNSSliderTouchBarItem() NSSliderTouchBarItem {
	class := getNSSliderTouchBarItemClass()
	rv := objc.Send[NSSliderTouchBarItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes and returns a new item from a storyboard or nib file.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/init(coder:)
func NewSliderTouchBarItemWithCoder(coder foundation.INSCoder) NSSliderTouchBarItem {
	instance := getNSSliderTouchBarItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSSliderTouchBarItemFromID(rv)
}

// Creates a new item with the specified identifier.
//
// # Discussion
// 
// The designated initializer. The identifier must be globally unique for
// every item, except for space items.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/init(identifier:)
func NewSliderTouchBarItemWithIdentifier(identifier NSTouchBarItemIdentifier) NSSliderTouchBarItem {
	instance := getNSSliderTouchBarItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIdentifier:"), objc.String(string(identifier)))
	return NSSliderTouchBarItemFromID(rv)
}

// The slider displayed by the bar item.
//
// # Discussion
// 
// The slider is automatically created by the item, but you can provide a
// custom subclass to this property to customize behavior.
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/slider
func (s NSSliderTouchBarItem) Slider() INSSlider {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("slider"))
	return NSSliderFromID(objc.ID(rv))
}
func (s NSSliderTouchBarItem) SetSlider(value INSSlider) {
	objc.Send[struct{}](s.ID, objc.Sel("setSlider:"), value)
}
// The text displayed alongside the slider.
//
// # Discussion
// 
// Set this property to `nil` to hide the label.
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/label
func (s NSSliderTouchBarItem) Label() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}
func (s NSSliderTouchBarItem) SetLabel(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setLabel:"), objc.String(value))
}
// The accessory that appears at the end of the slider with the minimum value.
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/minimumValueAccessory
func (s NSSliderTouchBarItem) MinimumValueAccessory() INSSliderAccessory {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("minimumValueAccessory"))
	return NSSliderAccessoryFromID(objc.ID(rv))
}
func (s NSSliderTouchBarItem) SetMinimumValueAccessory(value INSSliderAccessory) {
	objc.Send[struct{}](s.ID, objc.Sel("setMinimumValueAccessory:"), value)
}
// The accessory that appears at the end of the slider with the maximum value.
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/maximumValueAccessory
func (s NSSliderTouchBarItem) MaximumValueAccessory() INSSliderAccessory {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("maximumValueAccessory"))
	return NSSliderAccessoryFromID(objc.ID(rv))
}
func (s NSSliderTouchBarItem) SetMaximumValueAccessory(value INSSliderAccessory) {
	objc.Send[struct{}](s.ID, objc.Sel("setMaximumValueAccessory:"), value)
}
// The width of the value accessories that appear at either end of the slider.
//
// # Discussion
// 
// You can provide your own custom width, or use the system-provided [default]
// or [wide] options.
// 
// The default value is [default].
//
// [default]: https://developer.apple.com/documentation/AppKit/NSSliderAccessory/Width/default
// [wide]: https://developer.apple.com/documentation/AppKit/NSSliderAccessory/Width/wide
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/valueAccessoryWidth
func (s NSSliderTouchBarItem) ValueAccessoryWidth() NSSliderAccessoryWidth {
	rv := objc.Send[NSSliderAccessoryWidth](s.ID, objc.Sel("valueAccessoryWidth"))
	return NSSliderAccessoryWidth(rv)
}
func (s NSSliderTouchBarItem) SetValueAccessoryWidth(value NSSliderAccessoryWidth) {
	objc.Send[struct{}](s.ID, objc.Sel("setValueAccessoryWidth:"), value)
}
// The minimum width of the slider’s track.
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/minimumSliderWidth
func (s NSSliderTouchBarItem) MinimumSliderWidth() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("minimumSliderWidth"))
	return rv
}
func (s NSSliderTouchBarItem) SetMinimumSliderWidth(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setMinimumSliderWidth:"), value)
}
// The maximum width of the slider’s track.
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/maximumSliderWidth
func (s NSSliderTouchBarItem) MaximumSliderWidth() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("maximumSliderWidth"))
	return rv
}
func (s NSSliderTouchBarItem) SetMaximumSliderWidth(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setMaximumSliderWidth:"), value)
}
// The double value of the slider.
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/doubleValue
func (s NSSliderTouchBarItem) DoubleValue() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("doubleValue"))
	return rv
}
func (s NSSliderTouchBarItem) SetDoubleValue(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setDoubleValue:"), value)
}
// An object that is notified when a user interacts with the slider or either
// of the accessories.
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/target
func (s NSSliderTouchBarItem) Target() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("target"))
	return objectivec.Object{ID: rv}
}
func (s NSSliderTouchBarItem) SetTarget(value objectivec.IObject) {
	objc.Send[struct{}](s.ID, objc.Sel("setTarget:"), value)
}
// The selector on the target object that is invoked when a user interacts
// with the slider or either of the accessories.
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem/action
func (s NSSliderTouchBarItem) Action() objc.SEL {
	rv := objc.Send[objc.SEL](s.ID, objc.Sel("action"))
	return rv
}
func (s NSSliderTouchBarItem) SetAction(value objc.SEL) {
	objc.Send[struct{}](s.ID, objc.Sel("setAction:"), value)
}

