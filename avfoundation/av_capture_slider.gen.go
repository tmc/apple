// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCaptureSlider] class.
var (
	_AVCaptureSliderClass     AVCaptureSliderClass
	_AVCaptureSliderClassOnce sync.Once
)

func getAVCaptureSliderClass() AVCaptureSliderClass {
	_AVCaptureSliderClassOnce.Do(func() {
		_AVCaptureSliderClass = AVCaptureSliderClass{class: objc.GetClass("AVCaptureSlider")}
	})
	return _AVCaptureSliderClass
}

// GetAVCaptureSliderClass returns the class object for AVCaptureSlider.
func GetAVCaptureSliderClass() AVCaptureSliderClass {
	return getAVCaptureSliderClass()
}

type AVCaptureSliderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCaptureSliderClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptureSliderClass) Alloc() AVCaptureSlider {
	rv := objc.Send[AVCaptureSlider](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A slider control that selects a value from a bounded range.
//
// # Overview
// 
// Sliders are appropriate for controls that provide a single float value.
//
// # Accessing the control value
//
//   - [AVCaptureSlider.Value]: The current value of the slider.
//   - [AVCaptureSlider.SetValue]
//   - [AVCaptureSlider.LocalizedValueFormat]: A localized string that defines the presentation of the slider’s value.
//   - [AVCaptureSlider.SetLocalizedValueFormat]
//
// # Inspecting presentation attributes
//
//   - [AVCaptureSlider.SymbolName]: The name of the SF Symbol that represents this control.
//   - [AVCaptureSlider.LocalizedTitle]: A localized title that describes the control’s action.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSlider
type AVCaptureSlider struct {
	AVCaptureControl
}

// AVCaptureSliderFromID constructs a [AVCaptureSlider] from an objc.ID.
//
// A slider control that selects a value from a bounded range.
func AVCaptureSliderFromID(id objc.ID) AVCaptureSlider {
	return AVCaptureSlider{AVCaptureControl: AVCaptureControlFromID(id)}
}
// NOTE: AVCaptureSlider adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptureSlider] class.
//
// # Accessing the control value
//
//   - [IAVCaptureSlider.Value]: The current value of the slider.
//   - [IAVCaptureSlider.SetValue]
//   - [IAVCaptureSlider.LocalizedValueFormat]: A localized string that defines the presentation of the slider’s value.
//   - [IAVCaptureSlider.SetLocalizedValueFormat]
//
// # Inspecting presentation attributes
//
//   - [IAVCaptureSlider.SymbolName]: The name of the SF Symbol that represents this control.
//   - [IAVCaptureSlider.LocalizedTitle]: A localized title that describes the control’s action.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSlider
type IAVCaptureSlider interface {
	IAVCaptureControl

	// Topic: Accessing the control value

	// The current value of the slider.
	Value() float32
	SetValue(value float32)
	// A localized string that defines the presentation of the slider’s value.
	LocalizedValueFormat() string
	SetLocalizedValueFormat(value string)

	// Topic: Inspecting presentation attributes

	// The name of the SF Symbol that represents this control.
	SymbolName() string
	// A localized title that describes the control’s action.
	LocalizedTitle() string

	// Values in this array may receive unique visual representations or behaviors.
	ProminentValues() []foundation.NSNumber
	SetProminentValues(value []foundation.NSNumber)
	// Creates a continuous slider control that selects a value from a bounded range.
	InitWithLocalizedTitleSymbolNameMinValueMaxValue(localizedTitle string, symbolName string, minValue float32, maxValue float32) AVCaptureSlider
	// Creates a discrete slider control that selects a stepped value from a bounded range.
	InitWithLocalizedTitleSymbolNameMinValueMaxValueStep(localizedTitle string, symbolName string, minValue float32, maxValue float32, step float32) AVCaptureSlider
	// Creates a discrete slider control that selects a value from a list.
	InitWithLocalizedTitleSymbolNameValues(localizedTitle string, symbolName string, values []foundation.NSNumber) AVCaptureSlider
	// Sets the action to perform on the specified dispatch queue when the slider’s value changes.
	SetActionQueueAction(actionQueue dispatch.Queue, action Float32Handler)
}

// Init initializes the instance.
func (c AVCaptureSlider) Init() AVCaptureSlider {
	rv := objc.Send[AVCaptureSlider](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptureSlider) Autorelease() AVCaptureSlider {
	rv := objc.Send[AVCaptureSlider](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureSlider creates a new AVCaptureSlider instance.
func NewAVCaptureSlider() AVCaptureSlider {
	class := getAVCaptureSliderClass()
	rv := objc.Send[AVCaptureSlider](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a continuous slider control that selects a value from a bounded
// range.
//
// localizedTitle: A localized title that describes the slider’s action.
//
// symbolName: A symbol name from the SF Symbols library.
//
// minValue: The lower bound of the range.
//
// maxValue: The upper bound of the range.
//
// # Discussion
// 
// Use continuous sliders when your use case supports selecting any value in
// the specified range.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSlider/initWithLocalizedTitle:symbolName:minValue:maxValue:
func NewCaptureSliderWithLocalizedTitleSymbolNameMinValueMaxValue(localizedTitle string, symbolName string, minValue float32, maxValue float32) AVCaptureSlider {
	instance := getAVCaptureSliderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLocalizedTitle:symbolName:minValue:maxValue:"), objc.String(localizedTitle), objc.String(symbolName), minValue, maxValue)
	return AVCaptureSliderFromID(rv)
}

// Creates a discrete slider control that selects a stepped value from a
// bounded range.
//
// localizedTitle: A localized title that describes the slider’s action.
//
// symbolName: A symbol name from the SF Symbols library.
//
// minValue: The lower bound of the range.
//
// maxValue: The upper bound of the range.
//
// step: The distance between each valid value. This specified value must be greater
// than `0` or the system throws an [InvalidArgumentException].
//
// # Discussion
// 
// Use discrete sliders when your use case supports selecting stepped values
// within the specified range.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSlider/initWithLocalizedTitle:symbolName:minValue:maxValue:step:
func NewCaptureSliderWithLocalizedTitleSymbolNameMinValueMaxValueStep(localizedTitle string, symbolName string, minValue float32, maxValue float32, step float32) AVCaptureSlider {
	instance := getAVCaptureSliderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLocalizedTitle:symbolName:minValue:maxValue:step:"), objc.String(localizedTitle), objc.String(symbolName), minValue, maxValue, step)
	return AVCaptureSliderFromID(rv)
}

// Creates a discrete slider control that selects a value from a list.
//
// localizedTitle: A localized title that describes the slider’s action.
//
// symbolName: A symbol name from the SF Symbols library.
//
// values: An array of floating-point values.
//
// # Discussion
// 
// Use discrete sliders when your app supports selecting from a specific list
// of values.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSlider/initWithLocalizedTitle:symbolName:values:
func NewCaptureSliderWithLocalizedTitleSymbolNameValues(localizedTitle string, symbolName string, values []foundation.NSNumber) AVCaptureSlider {
	instance := getAVCaptureSliderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLocalizedTitle:symbolName:values:"), objc.String(localizedTitle), objc.String(symbolName), objectivec.IObjectSliceToNSArray(values))
	return AVCaptureSliderFromID(rv)
}

// Creates a continuous slider control that selects a value from a bounded
// range.
//
// localizedTitle: A localized title that describes the slider’s action.
//
// symbolName: A symbol name from the SF Symbols library.
//
// minValue: The lower bound of the range.
//
// maxValue: The upper bound of the range.
//
// # Discussion
// 
// Use continuous sliders when your use case supports selecting any value in
// the specified range.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSlider/initWithLocalizedTitle:symbolName:minValue:maxValue:
func (c AVCaptureSlider) InitWithLocalizedTitleSymbolNameMinValueMaxValue(localizedTitle string, symbolName string, minValue float32, maxValue float32) AVCaptureSlider {
	rv := objc.Send[AVCaptureSlider](c.ID, objc.Sel("initWithLocalizedTitle:symbolName:minValue:maxValue:"), objc.String(localizedTitle), objc.String(symbolName), minValue, maxValue)
	return rv
}
// Creates a discrete slider control that selects a stepped value from a
// bounded range.
//
// localizedTitle: A localized title that describes the slider’s action.
//
// symbolName: A symbol name from the SF Symbols library.
//
// minValue: The lower bound of the range.
//
// maxValue: The upper bound of the range.
//
// step: The distance between each valid value. This specified value must be greater
// than `0` or the system throws an [InvalidArgumentException].
//
// # Discussion
// 
// Use discrete sliders when your use case supports selecting stepped values
// within the specified range.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSlider/initWithLocalizedTitle:symbolName:minValue:maxValue:step:
func (c AVCaptureSlider) InitWithLocalizedTitleSymbolNameMinValueMaxValueStep(localizedTitle string, symbolName string, minValue float32, maxValue float32, step float32) AVCaptureSlider {
	rv := objc.Send[AVCaptureSlider](c.ID, objc.Sel("initWithLocalizedTitle:symbolName:minValue:maxValue:step:"), objc.String(localizedTitle), objc.String(symbolName), minValue, maxValue, step)
	return rv
}
// Creates a discrete slider control that selects a value from a list.
//
// localizedTitle: A localized title that describes the slider’s action.
//
// symbolName: A symbol name from the SF Symbols library.
//
// values: An array of floating-point values.
//
// # Discussion
// 
// Use discrete sliders when your app supports selecting from a specific list
// of values.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSlider/initWithLocalizedTitle:symbolName:values:
func (c AVCaptureSlider) InitWithLocalizedTitleSymbolNameValues(localizedTitle string, symbolName string, values []foundation.NSNumber) AVCaptureSlider {
	rv := objc.Send[AVCaptureSlider](c.ID, objc.Sel("initWithLocalizedTitle:symbolName:values:"), objc.String(localizedTitle), objc.String(symbolName), objectivec.IObjectSliceToNSArray(values))
	return rv
}
// Sets the action to perform on the specified dispatch queue when the
// slider’s value changes.
//
// actionQueue: A dispatch queue on which to call the action.
//
// action: The action to perform in response to changes to the slider’s value.
//
// # Discussion
// 
// If the action modifies a property of the camera system, the specified
// dispatch queue must represent the camera system’s same exclusive
// execution context (see [isSameExclusiveExecutionContext(other:)]).
//
// [isSameExclusiveExecutionContext(other:)]: https://developer.apple.com/documentation/Swift/SerialExecutor/isSameExclusiveExecutionContext(other:)-3ptya
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSlider/setActionQueue:action:
func (c AVCaptureSlider) SetActionQueueAction(actionQueue dispatch.Queue, action Float32Handler) {
_block1, _ := NewFloat32Block(action)
	objc.Send[objc.ID](c.ID, objc.Sel("setActionQueue:action:"), uintptr(actionQueue.Handle()), _block1)
}

// The current value of the slider.
//
// # Discussion
// 
// The default value is the slider’s minimum value. You may set a value only
// if it’s within the slider’s minimum and maximum values, otherwise the
// system throws an exception.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSlider/value
func (c AVCaptureSlider) Value() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("value"))
	return rv
}
func (c AVCaptureSlider) SetValue(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setValue:"), value)
}
// A localized string that defines the presentation of the slider’s value.
//
// # Discussion
// 
// Specify a format string to modify the presentation of a slider’s value.
// The format string may only contain `%@` and no other placeholders like
// `%d`, `%s`, and so on. Setting an Invalid format string results in the
// value’s default presentation.
// 
// Examples of valid format strings are:
// 
// - “%@%” for “40%” - “%@ fps” for “60 fps” - “+ %@” for
// “+ 20”
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSlider/localizedValueFormat
func (c AVCaptureSlider) LocalizedValueFormat() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("localizedValueFormat"))
	return foundation.NSStringFromID(rv).String()
}
func (c AVCaptureSlider) SetLocalizedValueFormat(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setLocalizedValueFormat:"), objc.String(value))
}
// The name of the SF Symbol that represents this control.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSlider/symbolName
func (c AVCaptureSlider) SymbolName() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("symbolName"))
	return foundation.NSStringFromID(rv).String()
}
// A localized title that describes the control’s action.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSlider/localizedTitle
func (c AVCaptureSlider) LocalizedTitle() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("localizedTitle"))
	return foundation.NSStringFromID(rv).String()
}
// Values in this array may receive unique visual representations or
// behaviors.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSlider/prominentValues-7usgc
func (c AVCaptureSlider) ProminentValues() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("prominentValues"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
func (c AVCaptureSlider) SetProminentValues(value []foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setProminentValues:"), objectivec.IObjectSliceToNSArray(value))
}

// SetActionQueueActionSync is a synchronous wrapper around [AVCaptureSlider.SetActionQueueAction].
// It blocks until the completion handler fires or the context is cancelled.
func (c AVCaptureSlider) SetActionQueueActionSync(ctx context.Context, actionQueue dispatch.Queue) (float32, error) {
	done := make(chan float32, 1)
	c.SetActionQueueAction(actionQueue, func(val float32) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return 0.0, ctx.Err()
	}
}

