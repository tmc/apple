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

// The class instance for the [AVCaptureIndexPicker] class.
var (
	_AVCaptureIndexPickerClass     AVCaptureIndexPickerClass
	_AVCaptureIndexPickerClassOnce sync.Once
)

func getAVCaptureIndexPickerClass() AVCaptureIndexPickerClass {
	_AVCaptureIndexPickerClassOnce.Do(func() {
		_AVCaptureIndexPickerClass = AVCaptureIndexPickerClass{class: objc.GetClass("AVCaptureIndexPicker")}
	})
	return _AVCaptureIndexPickerClass
}

// GetAVCaptureIndexPickerClass returns the class object for AVCaptureIndexPicker.
func GetAVCaptureIndexPickerClass() AVCaptureIndexPickerClass {
	return getAVCaptureIndexPickerClass()
}

type AVCaptureIndexPickerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptureIndexPickerClass) Alloc() AVCaptureIndexPicker {
	rv := objc.Send[AVCaptureIndexPicker](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A control for selecting from a set of mutually exclusive values by index.
//
// # Overview
// 
// Index pickers are appropriate for controls that provide an indexed
// container of values.
//
// # Creating an index picker
//
//   - [AVCaptureIndexPicker.InitWithLocalizedTitleSymbolNameNumberOfIndexes]: Creates a control to pick a value from the specified number of indexes.
//   - [AVCaptureIndexPicker.InitWithLocalizedTitleSymbolNameNumberOfIndexesLocalizedTitleTransform]: Creates a control to pick a value from the specified number of indices.
//   - [AVCaptureIndexPicker.InitWithLocalizedTitleSymbolNameLocalizedIndexTitles]: Creates an object to select an index from a set of values.
//
// # Accessing the control value
//
//   - [AVCaptureIndexPicker.SelectedIndex]: The currently selected index.
//   - [AVCaptureIndexPicker.SetSelectedIndex]
//   - [AVCaptureIndexPicker.NumberOfIndexes]: The number of index values the control provides.
//
// # Inspecting presentation attributes
//
//   - [AVCaptureIndexPicker.SymbolName]: The name of the SF Symbol that represents this control.
//   - [AVCaptureIndexPicker.LocalizedTitle]: A localized title that describes the control’s action.
//   - [AVCaptureIndexPicker.LocalizedIndexTitles]: The titles to present for each index.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureIndexPicker
type AVCaptureIndexPicker struct {
	AVCaptureControl
}

// AVCaptureIndexPickerFromID constructs a [AVCaptureIndexPicker] from an objc.ID.
//
// A control for selecting from a set of mutually exclusive values by index.
func AVCaptureIndexPickerFromID(id objc.ID) AVCaptureIndexPicker {
	return AVCaptureIndexPicker{AVCaptureControl: AVCaptureControlFromID(id)}
}
// NOTE: AVCaptureIndexPicker adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptureIndexPicker] class.
//
// # Creating an index picker
//
//   - [IAVCaptureIndexPicker.InitWithLocalizedTitleSymbolNameNumberOfIndexes]: Creates a control to pick a value from the specified number of indexes.
//   - [IAVCaptureIndexPicker.InitWithLocalizedTitleSymbolNameNumberOfIndexesLocalizedTitleTransform]: Creates a control to pick a value from the specified number of indices.
//   - [IAVCaptureIndexPicker.InitWithLocalizedTitleSymbolNameLocalizedIndexTitles]: Creates an object to select an index from a set of values.
//
// # Accessing the control value
//
//   - [IAVCaptureIndexPicker.SelectedIndex]: The currently selected index.
//   - [IAVCaptureIndexPicker.SetSelectedIndex]
//   - [IAVCaptureIndexPicker.NumberOfIndexes]: The number of index values the control provides.
//
// # Inspecting presentation attributes
//
//   - [IAVCaptureIndexPicker.SymbolName]: The name of the SF Symbol that represents this control.
//   - [IAVCaptureIndexPicker.LocalizedTitle]: A localized title that describes the control’s action.
//   - [IAVCaptureIndexPicker.LocalizedIndexTitles]: The titles to present for each index.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureIndexPicker
type IAVCaptureIndexPicker interface {
	IAVCaptureControl

	// Topic: Creating an index picker

	// Creates a control to pick a value from the specified number of indexes.
	InitWithLocalizedTitleSymbolNameNumberOfIndexes(localizedTitle string, symbolName string, numberOfIndexes int) AVCaptureIndexPicker
	// Creates a control to pick a value from the specified number of indices.
	InitWithLocalizedTitleSymbolNameNumberOfIndexesLocalizedTitleTransform(localizedTitle string, symbolName string, numberOfIndexes int, localizedTitleTransform IntHandler) AVCaptureIndexPicker
	// Creates an object to select an index from a set of values.
	InitWithLocalizedTitleSymbolNameLocalizedIndexTitles(localizedTitle string, symbolName string, localizedIndexTitles []string) AVCaptureIndexPicker

	// Topic: Accessing the control value

	// The currently selected index.
	SelectedIndex() int
	SetSelectedIndex(value int)
	// The number of index values the control provides.
	NumberOfIndexes() int

	// Topic: Inspecting presentation attributes

	// The name of the SF Symbol that represents this control.
	SymbolName() string
	// A localized title that describes the control’s action.
	LocalizedTitle() string
	// The titles to present for each index.
	LocalizedIndexTitles() []string

	// Sets the action to perform on the specified dispatch queue when the control’s value changes.
	SetActionQueueAction(actionQueue dispatch.Queue, action IntHandler)
}

// Init initializes the instance.
func (c AVCaptureIndexPicker) Init() AVCaptureIndexPicker {
	rv := objc.Send[AVCaptureIndexPicker](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptureIndexPicker) Autorelease() AVCaptureIndexPicker {
	rv := objc.Send[AVCaptureIndexPicker](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureIndexPicker creates a new AVCaptureIndexPicker instance.
func NewAVCaptureIndexPicker() AVCaptureIndexPicker {
	class := getAVCaptureIndexPickerClass()
	rv := objc.Send[AVCaptureIndexPicker](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an object to select an index from a set of values.
//
// localizedTitle: A localized title that describes the control’s action.
//
// symbolName: The name of an SF Symbol that represents the control.
//
// localizedIndexTitles: The titles to use for each index. The array must not be empty.
//
// # Discussion
// 
// Create a picker with this initializer when you already have an array
// containing a title for each picked value.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureIndexPicker/init(_:symbolName:localizedIndexTitles:)
func NewCaptureIndexPickerWithLocalizedTitleSymbolNameLocalizedIndexTitles(localizedTitle string, symbolName string, localizedIndexTitles []string) AVCaptureIndexPicker {
	instance := getAVCaptureIndexPickerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLocalizedTitle:symbolName:localizedIndexTitles:"), objc.String(localizedTitle), objc.String(symbolName), objectivec.StringSliceToNSArray(localizedIndexTitles))
	return AVCaptureIndexPickerFromID(rv)
}

// Creates a control to pick a value from the specified number of indexes.
//
// localizedTitle: A localized title that describes the picker’s action.
//
// symbolName: The name of the symbol from the SF Symbols library to use to represent this
// control.
//
// numberOfIndexes: The number of indexes to pick between. This value must be greater than `0`.
//
// # Discussion
// 
// Create a picker with this initializer when the control’s values don’t
// require titles.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureIndexPicker/init(_:symbolName:numberOfIndexes:)
func NewCaptureIndexPickerWithLocalizedTitleSymbolNameNumberOfIndexes(localizedTitle string, symbolName string, numberOfIndexes int) AVCaptureIndexPicker {
	instance := getAVCaptureIndexPickerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLocalizedTitle:symbolName:numberOfIndexes:"), objc.String(localizedTitle), objc.String(symbolName), numberOfIndexes)
	return AVCaptureIndexPickerFromID(rv)
}

// Creates a control to pick a value from the specified number of indexes.
//
// localizedTitle: A localized title that describes the picker’s action.
//
// symbolName: The name of the symbol from the SF Symbols library to use to represent this
// control.
//
// numberOfIndexes: The number of indexes to pick between. This value must be greater than `0`.
//
// # Discussion
// 
// Create a picker with this initializer when the control’s values don’t
// require titles.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureIndexPicker/init(_:symbolName:numberOfIndexes:)
func (c AVCaptureIndexPicker) InitWithLocalizedTitleSymbolNameNumberOfIndexes(localizedTitle string, symbolName string, numberOfIndexes int) AVCaptureIndexPicker {
	rv := objc.Send[AVCaptureIndexPicker](c.ID, objc.Sel("initWithLocalizedTitle:symbolName:numberOfIndexes:"), objc.String(localizedTitle), objc.String(symbolName), numberOfIndexes)
	return rv
}
// Creates a control to pick a value from the specified number of indices.
//
// localizedTitle: A localized title that describes the picker’s action.
//
// symbolName: The name of the symbol from the SF Symbols library to use to represent this
// control.
//
// numberOfIndexes: The number of indexes to pick between. This value must be greater than `0`.
//
// localizedTitleTransform: A transformation from index to localized title.
//
// # Discussion
// 
// Create a picker with this initializer if your app requires specifying a
// title for each value lazily.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureIndexPicker/init(_:symbolName:numberOfIndexes:localizedTitleTransform:)
func (c AVCaptureIndexPicker) InitWithLocalizedTitleSymbolNameNumberOfIndexesLocalizedTitleTransform(localizedTitle string, symbolName string, numberOfIndexes int, localizedTitleTransform IntHandler) AVCaptureIndexPicker {
_block3, _cleanup3 := NewIntBlock(localizedTitleTransform)
	defer _cleanup3()
	rv := objc.Send[objc.ID](c.ID, objc.Sel("initWithLocalizedTitle:symbolName:numberOfIndexes:localizedTitleTransform:"), objc.String(localizedTitle), objc.String(symbolName), numberOfIndexes, _block3)
	return AVCaptureIndexPickerFromID(rv)
}
// Creates an object to select an index from a set of values.
//
// localizedTitle: A localized title that describes the control’s action.
//
// symbolName: The name of an SF Symbol that represents the control.
//
// localizedIndexTitles: The titles to use for each index. The array must not be empty.
//
// # Discussion
// 
// Create a picker with this initializer when you already have an array
// containing a title for each picked value.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureIndexPicker/init(_:symbolName:localizedIndexTitles:)
func (c AVCaptureIndexPicker) InitWithLocalizedTitleSymbolNameLocalizedIndexTitles(localizedTitle string, symbolName string, localizedIndexTitles []string) AVCaptureIndexPicker {
	rv := objc.Send[AVCaptureIndexPicker](c.ID, objc.Sel("initWithLocalizedTitle:symbolName:localizedIndexTitles:"), objc.String(localizedTitle), objc.String(symbolName), objectivec.StringSliceToNSArray(localizedIndexTitles))
	return rv
}
// Sets the action to perform on the specified dispatch queue when the
// control’s value changes.
//
// actionQueue: A dispatch queue on which to call the action.
//
// action: The action to perform in response to changes to the control’s value.
//
// # Discussion
// 
// If the action modifies a property of the camera system, the specified
// dispatch queue must represent the camera system’s same exclusive
// execution context (see [isSameExclusiveExecutionContext(other:)]).
//
// [isSameExclusiveExecutionContext(other:)]: https://developer.apple.com/documentation/Swift/SerialExecutor/isSameExclusiveExecutionContext(other:)-3ptya
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureIndexPicker/setActionQueue:action:
func (c AVCaptureIndexPicker) SetActionQueueAction(actionQueue dispatch.Queue, action IntHandler) {
_block1, _ := NewIntBlock(action)
	objc.Send[objc.ID](c.ID, objc.Sel("setActionQueue:action:"), uintptr(actionQueue.Handle()), _block1)
}

// The currently selected index.
//
// # Discussion
// 
// The default value is `0`. You can only set a value that’s greater than or
// equal to `0` and less than [NumberOfIndexes].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureIndexPicker/selectedIndex
func (c AVCaptureIndexPicker) SelectedIndex() int {
	rv := objc.Send[int](c.ID, objc.Sel("selectedIndex"))
	return rv
}
func (c AVCaptureIndexPicker) SetSelectedIndex(value int) {
	objc.Send[struct{}](c.ID, objc.Sel("setSelectedIndex:"), value)
}
// The number of index values the control provides.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureIndexPicker/numberOfIndexes
func (c AVCaptureIndexPicker) NumberOfIndexes() int {
	rv := objc.Send[int](c.ID, objc.Sel("numberOfIndexes"))
	return rv
}
// The name of the SF Symbol that represents this control.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureIndexPicker/symbolName
func (c AVCaptureIndexPicker) SymbolName() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("symbolName"))
	return foundation.NSStringFromID(rv).String()
}
// A localized title that describes the control’s action.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureIndexPicker/localizedTitle
func (c AVCaptureIndexPicker) LocalizedTitle() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("localizedTitle"))
	return foundation.NSStringFromID(rv).String()
}
// The titles to present for each index.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureIndexPicker/localizedIndexTitles
func (c AVCaptureIndexPicker) LocalizedIndexTitles() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("localizedIndexTitles"))
	return objc.ConvertSliceToStrings(rv)
}

// InitWithLocalizedTitleSymbolNameNumberOfIndexesLocalizedTitleTransformSync is a synchronous wrapper around [AVCaptureIndexPicker.InitWithLocalizedTitleSymbolNameNumberOfIndexesLocalizedTitleTransform].
// It blocks until the completion handler fires or the context is cancelled.
func (c AVCaptureIndexPicker) InitWithLocalizedTitleSymbolNameNumberOfIndexesLocalizedTitleTransformSync(ctx context.Context, localizedTitle string, symbolName string, numberOfIndexes int) (int, error) {
	done := make(chan int, 1)
	c.InitWithLocalizedTitleSymbolNameNumberOfIndexesLocalizedTitleTransform(localizedTitle, symbolName, numberOfIndexes, func(val int) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

// SetActionQueueActionSync is a synchronous wrapper around [AVCaptureIndexPicker.SetActionQueueAction].
// It blocks until the completion handler fires or the context is cancelled.
func (c AVCaptureIndexPicker) SetActionQueueActionSync(ctx context.Context, actionQueue dispatch.Queue) (int, error) {
	done := make(chan int, 1)
	c.SetActionQueueAction(actionQueue, func(val int) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

