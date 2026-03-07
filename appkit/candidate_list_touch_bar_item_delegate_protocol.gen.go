// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf


// A set of methods that a candidate list item delegate uses to enable selection state and list visibility.
//
// See: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItemDelegate
type NSCandidateListTouchBarItemDelegate interface {
	objectivec.IObject
}



// NSCandidateListTouchBarItemDelegateObject wraps an existing Objective-C object that conforms to the NSCandidateListTouchBarItemDelegate protocol.
type NSCandidateListTouchBarItemDelegateObject struct {
	objectivec.Object
}
func (o NSCandidateListTouchBarItemDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}



// NSCandidateListTouchBarItemDelegateObjectFromID constructs a [NSCandidateListTouchBarItemDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSCandidateListTouchBarItemDelegateObjectFromID(id objc.ID) NSCandidateListTouchBarItemDelegateObject {
	return NSCandidateListTouchBarItemDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}




// Tells the delegate that the user has started touching one of the candidates
// in the candidate list item.
//
// anItem: The candidate list bar item that the user is interacting with.
//
// index: The index of the candidate that the user is currently touching.
//
// See: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItemDelegate/candidateListTouchBarItem(_:beginSelectingCandidateAt:)

func (o NSCandidateListTouchBarItemDelegateObject) CandidateListTouchBarItemBeginSelectingCandidateAtIndex(anItem INSCandidateListTouchBarItem, index int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("candidateListTouchBarItem:beginSelectingCandidateAtIndex:"), anItem, index)
	}

// Tells the delegate that user has moved from touching one candidate in the
// candidate list item to another.
//
// anItem: The candidate list item that the user is interacting with.
//
// previousIndex: The index of the candidate that the user was previously touching.
//
// index: The index of the candidate that the user is currently touching.
//
// See: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItemDelegate/candidateListTouchBarItem(_:changeSelectionFromCandidateAt:to:)

func (o NSCandidateListTouchBarItemDelegateObject) CandidateListTouchBarItemChangeSelectionFromCandidateAtIndexToIndex(anItem INSCandidateListTouchBarItem, previousIndex int, index int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("candidateListTouchBarItem:changeSelectionFromCandidateAtIndex:toIndex:"), anItem, previousIndex, index)
	}

// Tells the delegate that a user has stopped touching candidates in the
// candidate list item.
//
// anItem: The candidate list item that the user is interacting with.
//
// index: The index of the candidate that the user was touching when they lifted
// their finger.
//
// # Discussion
// 
// If `index` is equal to [NSNotFound] then the user didn’t select a
// candidate.
//
// See: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItemDelegate/candidateListTouchBarItem(_:endSelectingCandidateAt:)

func (o NSCandidateListTouchBarItemDelegateObject) CandidateListTouchBarItemEndSelectingCandidateAtIndex(anItem INSCandidateListTouchBarItem, index int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("candidateListTouchBarItem:endSelectingCandidateAtIndex:"), anItem, index)
	}

// Tells the delegate that the visibility of the candidate list has changed.
//
// anItem: The candidate list item whose candidate list’s visibility has changed.
//
// isVisible: A Boolean value that specifies whether or not the candidate list is
// visible. If [true] then the candidate list is visible, [false] otherwise.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItemDelegate/candidateListTouchBarItem(_:changedCandidateListVisibility:)

func (o NSCandidateListTouchBarItemDelegateObject) CandidateListTouchBarItemChangedCandidateListVisibility(anItem INSCandidateListTouchBarItem, isVisible bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("candidateListTouchBarItem:changedCandidateListVisibility:"), anItem, isVisible)
	}





// NSCandidateListTouchBarItemDelegateConfig holds optional typed callbacks for [NSCandidateListTouchBarItemDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritemdelegate
type NSCandidateListTouchBarItemDelegateConfig struct {

	// Handling visibility changes
	// CandidateListTouchBarItemChangedCandidateListVisibility — Tells the delegate that the visibility of the candidate list has changed.
	CandidateListTouchBarItemChangedCandidateListVisibility func(anItem NSCandidateListTouchBarItem, isVisible bool)

	// Other Methods
	// CandidateListTouchBarItemBeginSelectingCandidateAtIndex — Tells the delegate that the user has started touching one of the candidates in the candidate list item.
	CandidateListTouchBarItemBeginSelectingCandidateAtIndex func(anItem NSCandidateListTouchBarItem, index int)
	// CandidateListTouchBarItemChangeSelectionFromCandidateAtIndexToIndex — Tells the delegate that user has moved from touching one candidate in the candidate list item to another.
	CandidateListTouchBarItemChangeSelectionFromCandidateAtIndexToIndex func(anItem NSCandidateListTouchBarItem, previousIndex int, index int)
	// CandidateListTouchBarItemEndSelectingCandidateAtIndex — Tells the delegate that a user has stopped touching candidates in the candidate list item.
	CandidateListTouchBarItemEndSelectingCandidateAtIndex func(anItem NSCandidateListTouchBarItem, index int)
}

// NewNSCandidateListTouchBarItemDelegate creates an Objective-C object implementing the [NSCandidateListTouchBarItemDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSCandidateListTouchBarItemDelegateObject] satisfies the [NSCandidateListTouchBarItemDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritemdelegate
func NewNSCandidateListTouchBarItemDelegate(config NSCandidateListTouchBarItemDelegateConfig) NSCandidateListTouchBarItemDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSCandidateListTouchBarItemDelegate_%d", n)

	var methods []objc.MethodDef

	if config.CandidateListTouchBarItemBeginSelectingCandidateAtIndex != nil {
		fn := config.CandidateListTouchBarItemBeginSelectingCandidateAtIndex
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("candidateListTouchBarItem:beginSelectingCandidateAtIndex:"),
			Fn: func(self objc.ID, _cmd objc.SEL, anItemID objc.ID, index int) {
				anItem := NSCandidateListTouchBarItemFromID(anItemID)
				fn(anItem, index)
			},
		})
	}

	if config.CandidateListTouchBarItemChangeSelectionFromCandidateAtIndexToIndex != nil {
		fn := config.CandidateListTouchBarItemChangeSelectionFromCandidateAtIndexToIndex
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("candidateListTouchBarItem:changeSelectionFromCandidateAtIndex:toIndex:"),
			Fn: func(self objc.ID, _cmd objc.SEL, anItemID objc.ID, previousIndex int, index int) {
				anItem := NSCandidateListTouchBarItemFromID(anItemID)
				fn(anItem, previousIndex, index)
			},
		})
	}

	if config.CandidateListTouchBarItemEndSelectingCandidateAtIndex != nil {
		fn := config.CandidateListTouchBarItemEndSelectingCandidateAtIndex
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("candidateListTouchBarItem:endSelectingCandidateAtIndex:"),
			Fn: func(self objc.ID, _cmd objc.SEL, anItemID objc.ID, index int) {
				anItem := NSCandidateListTouchBarItemFromID(anItemID)
				fn(anItem, index)
			},
		})
	}

	if config.CandidateListTouchBarItemChangedCandidateListVisibility != nil {
		fn := config.CandidateListTouchBarItemChangedCandidateListVisibility
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("candidateListTouchBarItem:changedCandidateListVisibility:"),
			Fn: func(self objc.ID, _cmd objc.SEL, anItemID objc.ID, isVisible bool) {
				anItem := NSCandidateListTouchBarItemFromID(anItemID)
				fn(anItem, isVisible)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSCandidateListTouchBarItemDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSCandidateListTouchBarItemDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSCandidateListTouchBarItemDelegateObjectFromID(instance)
}





