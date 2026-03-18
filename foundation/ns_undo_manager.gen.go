// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [UndoManager] class.
var (
	_UndoManagerClass     UndoManagerClass
	_UndoManagerClassOnce sync.Once
)

func getUndoManagerClass() UndoManagerClass {
	_UndoManagerClassOnce.Do(func() {
		_UndoManagerClass = UndoManagerClass{class: objc.GetClass("NSUndoManager")}
	})
	return _UndoManagerClass
}

// GetUndoManagerClass returns the class object for NSUndoManager.
func GetUndoManagerClass() UndoManagerClass {
	return getUndoManagerClass()
}

type UndoManagerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (uc UndoManagerClass) Alloc() UndoManager {
	rv := objc.Send[UndoManager](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// A general-purpose recorder of operations that enables undo and redo.
//
// # Overview
// 
// You register an undo operation by calling one of the methods described in
// Registering undo operations. You specify the name of the object that’s
// changing (or the owner of that object) and provide a closure, method, or
// invocation to revert its state.
// 
// After you register an undo operation, you can call [Undo] on the undo
// manager to revert to the state of the last undo operation. When undoing an
// action, [NSUndoManager] saves the operations you revert to so that you can
// call [Redo] automatically.
// 
// Typically, apps with UI interactions work with [NSUndoManager]. For
// example, UIKit implements undo and redo in its text view object, making it
// easy for you to undo and redo actions in objects along the responder chain.
// [NSUndoManager] also serves as a general-purpose state manager, which you
// can use to undo and redo many kinds of actions. For example, an interactive
// command-line utility can use this class to undo the last command run, or a
// networking library can undo a request by sending another request that
// invalidates the previous one.
//
// # Registering undo operations
//
//   - [UndoManager.RegisterUndoWithTargetSelectorObject]: Registers the selector of the specified target to implement a single undo operation that the target receives.
//   - [UndoManager.PrepareWithInvocationTarget]: Prepares the undo manager for invocation-based undo with the given target as the subject of the next undo operation.
//
// # Checking undo ability
//
//   - [UndoManager.CanUndo]: A Boolean value that indicates whether the manager has any actions to undo.
//   - [UndoManager.CanRedo]: A Boolean value that indicates whether the manager has any actions to redo.
//
// # Performing undo and redo
//
//   - [UndoManager.Undo]: Closes the top-level undo group if necessary, and then performs undo operations on the group.
//   - [UndoManager.UndoNestedGroup]: Performs the undo operations in the last undo group (whether top-level or nested), recording the operations on the redo stack as a single group.
//   - [UndoManager.Redo]: Performs the operations in the last group on the redo stack, if there are any, recording them on the undo stack as a single group.
//
// # Managing undo and redo stack depth
//
//   - [UndoManager.LevelsOfUndo]: The maximum number of top-level undo groups the undo manager holds.
//   - [UndoManager.SetLevelsOfUndo]
//   - [UndoManager.UndoCount]: The number of times you can invoke undo before there are no actions left to undo.
//   - [UndoManager.RedoCount]: The number of times you can invoke redo before there are no actions left to redo.
//
// # Creating undo groups
//
//   - [UndoManager.BeginUndoGrouping]: Marks the beginning of an undo group.
//   - [UndoManager.EndUndoGrouping]: Marks the end of an undo group.
//   - [UndoManager.GroupsByEvent]: A Boolean value that indicates whether the manager automatically creates undo groups around each pass of the run loop.
//   - [UndoManager.SetGroupsByEvent]
//   - [UndoManager.GroupingLevel]: The number of nested undo groups (or redo groups, if redo is the most recent operation) in the current event loop.
//
// # Enabling and disabling undo
//
//   - [UndoManager.DisableUndoRegistration]: Disables the recording of undo operations.
//   - [UndoManager.EnableUndoRegistration]: Enables the recording of undo operations.
//   - [UndoManager.UndoRegistrationEnabled]: A Boolean value that indicates whether the recording of undo operations is enabled.
//
// # Checking whether undo or redo is in process
//
//   - [UndoManager.Undoing]: Returns a Boolean value that indicates whether the manager is in the process of performing an undo action.
//   - [UndoManager.Redoing]: Returns a Boolean value that indicates whether the manager is in the process of performing a redo action.
//
// # Clearing undo operations
//
//   - [UndoManager.RemoveAllActions]: Clears the undo and redo stacks and reenables the manager.
//   - [UndoManager.RemoveAllActionsWithTarget]: Clears the undo and redo stacks of all operations involving the specified target as the recipient of the undo message.
//
// # Managing the action name
//
//   - [UndoManager.UndoActionName]: The name identifying the undo action.
//   - [UndoManager.RedoActionName]: The name identifying the redo action.
//   - [UndoManager.SetActionName]: Sets the name of the action associated with the Undo or Redo command.
//
// # Getting and localizing the menu item title
//
//   - [UndoManager.UndoMenuItemTitle]: The title of the Undo menu command, such as Undo Paste.
//   - [UndoManager.RedoMenuItemTitle]: The title of the Redo menu command, such as Redo Paste.
//   - [UndoManager.UndoMenuTitleForUndoActionName]: Returns the localized title of the Undo menu command for the identified action.
//   - [UndoManager.RedoMenuTitleForUndoActionName]: Returns the localized title of the Redo menu command for the identified action.
//
// # Working with user info
//
//   - [UndoManager.SetActionUserInfoValueForKey]: Sets a user info value for an undo or redo action.
//   - [UndoManager.UndoActionUserInfoValueForKey]: Retrieves the undo action’s user info value for the given key.
//   - [UndoManager.RedoActionUserInfoValueForKey]: Retrieves the redo action’s user info value for the given key.
//
// # Working with run loops
//
//   - [UndoManager.RunLoopModes]: The modes governing the types of input to handle during a cycle of the run loop.
//   - [UndoManager.SetRunLoopModes]
//   - [UndoManager.NSUndoCloseGroupingRunLoopOrdering]: A priority to use when using a run loop to close an undo group.
//   - [UndoManager.SetNSUndoCloseGroupingRunLoopOrdering]
//
// # Using discardable undo and redo actions
//
//   - [UndoManager.SetActionIsDiscardable]: Sets whether the next undo or redo action is discardable.
//   - [UndoManager.UndoActionIsDiscardable]: A Boolean value that indicates whether the next undo action is discardable.
//   - [UndoManager.RedoActionIsDiscardable]: A Boolean value that indicates whether the next redo action is discardable.
//
// # Working with notifications
//
//   - [UndoManager.NSUndoManagerWillUndoChange]: Posted just before an undo manager performs an undo operation.
//   - [UndoManager.NSUndoManagerDidUndoChange]: Posted just after an undo manager performs an undo operation.
//   - [UndoManager.NSUndoManagerWillRedoChange]: Posted just before an undo manager performs a redo operation.
//   - [UndoManager.NSUndoManagerDidRedoChange]: Posted just after an undo manager performs a redo operation.
//   - [UndoManager.NSUndoManagerCheckpoint]: Posted whenever an undo manager opens or closes an undo group (except when it opens a top-level group) and when checking the redo stack.
//   - [UndoManager.NSUndoManagerDidOpenUndoGroup]: Posted whenever an undo manager opens an undo group.
//   - [UndoManager.NSUndoManagerWillCloseUndoGroup]: Posted before an undo manager closes an undo group.
//   - [UndoManager.NSUndoManagerDidCloseUndoGroup]: Posted after an undo manager closes an undo group.
//   - [UndoManager.NSUndoManagerGroupIsDiscardableKey]: A key, used in a notification’s user info, that indicates the undo group contains only discardable actions.
//
// See: https://developer.apple.com/documentation/Foundation/UndoManager
type UndoManager struct {
	objectivec.Object
}

// UndoManagerFromID constructs a [UndoManager] from an objc.ID.
//
// A general-purpose recorder of operations that enables undo and redo.
func UndoManagerFromID(id objc.ID) UndoManager {
	return NSUndoManager{objectivec.Object{ID: id}}
}

// NSUndoManagerFromID is an alias for [UndoManagerFromID] for cross-framework compatibility.
func NSUndoManagerFromID(id objc.ID) UndoManager { return UndoManagerFromID(id) }
// NOTE: UndoManager adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UndoManager] class.
//
// # Registering undo operations
//
//   - [IUndoManager.RegisterUndoWithTargetSelectorObject]: Registers the selector of the specified target to implement a single undo operation that the target receives.
//   - [IUndoManager.PrepareWithInvocationTarget]: Prepares the undo manager for invocation-based undo with the given target as the subject of the next undo operation.
//
// # Checking undo ability
//
//   - [IUndoManager.CanUndo]: A Boolean value that indicates whether the manager has any actions to undo.
//   - [IUndoManager.CanRedo]: A Boolean value that indicates whether the manager has any actions to redo.
//
// # Performing undo and redo
//
//   - [IUndoManager.Undo]: Closes the top-level undo group if necessary, and then performs undo operations on the group.
//   - [IUndoManager.UndoNestedGroup]: Performs the undo operations in the last undo group (whether top-level or nested), recording the operations on the redo stack as a single group.
//   - [IUndoManager.Redo]: Performs the operations in the last group on the redo stack, if there are any, recording them on the undo stack as a single group.
//
// # Managing undo and redo stack depth
//
//   - [IUndoManager.LevelsOfUndo]: The maximum number of top-level undo groups the undo manager holds.
//   - [IUndoManager.SetLevelsOfUndo]
//   - [IUndoManager.UndoCount]: The number of times you can invoke undo before there are no actions left to undo.
//   - [IUndoManager.RedoCount]: The number of times you can invoke redo before there are no actions left to redo.
//
// # Creating undo groups
//
//   - [IUndoManager.BeginUndoGrouping]: Marks the beginning of an undo group.
//   - [IUndoManager.EndUndoGrouping]: Marks the end of an undo group.
//   - [IUndoManager.GroupsByEvent]: A Boolean value that indicates whether the manager automatically creates undo groups around each pass of the run loop.
//   - [IUndoManager.SetGroupsByEvent]
//   - [IUndoManager.GroupingLevel]: The number of nested undo groups (or redo groups, if redo is the most recent operation) in the current event loop.
//
// # Enabling and disabling undo
//
//   - [IUndoManager.DisableUndoRegistration]: Disables the recording of undo operations.
//   - [IUndoManager.EnableUndoRegistration]: Enables the recording of undo operations.
//   - [IUndoManager.UndoRegistrationEnabled]: A Boolean value that indicates whether the recording of undo operations is enabled.
//
// # Checking whether undo or redo is in process
//
//   - [IUndoManager.Undoing]: Returns a Boolean value that indicates whether the manager is in the process of performing an undo action.
//   - [IUndoManager.Redoing]: Returns a Boolean value that indicates whether the manager is in the process of performing a redo action.
//
// # Clearing undo operations
//
//   - [IUndoManager.RemoveAllActions]: Clears the undo and redo stacks and reenables the manager.
//   - [IUndoManager.RemoveAllActionsWithTarget]: Clears the undo and redo stacks of all operations involving the specified target as the recipient of the undo message.
//
// # Managing the action name
//
//   - [IUndoManager.UndoActionName]: The name identifying the undo action.
//   - [IUndoManager.RedoActionName]: The name identifying the redo action.
//   - [IUndoManager.SetActionName]: Sets the name of the action associated with the Undo or Redo command.
//
// # Getting and localizing the menu item title
//
//   - [IUndoManager.UndoMenuItemTitle]: The title of the Undo menu command, such as Undo Paste.
//   - [IUndoManager.RedoMenuItemTitle]: The title of the Redo menu command, such as Redo Paste.
//   - [IUndoManager.UndoMenuTitleForUndoActionName]: Returns the localized title of the Undo menu command for the identified action.
//   - [IUndoManager.RedoMenuTitleForUndoActionName]: Returns the localized title of the Redo menu command for the identified action.
//
// # Working with user info
//
//   - [IUndoManager.SetActionUserInfoValueForKey]: Sets a user info value for an undo or redo action.
//   - [IUndoManager.UndoActionUserInfoValueForKey]: Retrieves the undo action’s user info value for the given key.
//   - [IUndoManager.RedoActionUserInfoValueForKey]: Retrieves the redo action’s user info value for the given key.
//
// # Working with run loops
//
//   - [IUndoManager.RunLoopModes]: The modes governing the types of input to handle during a cycle of the run loop.
//   - [IUndoManager.SetRunLoopModes]
//   - [IUndoManager.NSUndoCloseGroupingRunLoopOrdering]: A priority to use when using a run loop to close an undo group.
//   - [IUndoManager.SetNSUndoCloseGroupingRunLoopOrdering]
//
// # Using discardable undo and redo actions
//
//   - [IUndoManager.SetActionIsDiscardable]: Sets whether the next undo or redo action is discardable.
//   - [IUndoManager.UndoActionIsDiscardable]: A Boolean value that indicates whether the next undo action is discardable.
//   - [IUndoManager.RedoActionIsDiscardable]: A Boolean value that indicates whether the next redo action is discardable.
//
// # Working with notifications
//
//   - [IUndoManager.NSUndoManagerWillUndoChange]: Posted just before an undo manager performs an undo operation.
//   - [IUndoManager.NSUndoManagerDidUndoChange]: Posted just after an undo manager performs an undo operation.
//   - [IUndoManager.NSUndoManagerWillRedoChange]: Posted just before an undo manager performs a redo operation.
//   - [IUndoManager.NSUndoManagerDidRedoChange]: Posted just after an undo manager performs a redo operation.
//   - [IUndoManager.NSUndoManagerCheckpoint]: Posted whenever an undo manager opens or closes an undo group (except when it opens a top-level group) and when checking the redo stack.
//   - [IUndoManager.NSUndoManagerDidOpenUndoGroup]: Posted whenever an undo manager opens an undo group.
//   - [IUndoManager.NSUndoManagerWillCloseUndoGroup]: Posted before an undo manager closes an undo group.
//   - [IUndoManager.NSUndoManagerDidCloseUndoGroup]: Posted after an undo manager closes an undo group.
//   - [IUndoManager.NSUndoManagerGroupIsDiscardableKey]: A key, used in a notification’s user info, that indicates the undo group contains only discardable actions.
//
// See: https://developer.apple.com/documentation/Foundation/UndoManager
type IUndoManager interface {
	objectivec.IObject

	// Topic: Registering undo operations

	// Registers the selector of the specified target to implement a single undo operation that the target receives.
	RegisterUndoWithTargetSelectorObject(target objectivec.IObject, selector objc.SEL, object objectivec.IObject)
	// Prepares the undo manager for invocation-based undo with the given target as the subject of the next undo operation.
	PrepareWithInvocationTarget(target objectivec.IObject) objectivec.IObject

	// Topic: Checking undo ability

	// A Boolean value that indicates whether the manager has any actions to undo.
	CanUndo() bool
	// A Boolean value that indicates whether the manager has any actions to redo.
	CanRedo() bool

	// Topic: Performing undo and redo

	// Closes the top-level undo group if necessary, and then performs undo operations on the group.
	Undo()
	// Performs the undo operations in the last undo group (whether top-level or nested), recording the operations on the redo stack as a single group.
	UndoNestedGroup()
	// Performs the operations in the last group on the redo stack, if there are any, recording them on the undo stack as a single group.
	Redo()

	// Topic: Managing undo and redo stack depth

	// The maximum number of top-level undo groups the undo manager holds.
	LevelsOfUndo() uint
	SetLevelsOfUndo(value uint)
	// The number of times you can invoke undo before there are no actions left to undo.
	UndoCount() uint
	// The number of times you can invoke redo before there are no actions left to redo.
	RedoCount() uint

	// Topic: Creating undo groups

	// Marks the beginning of an undo group.
	BeginUndoGrouping()
	// Marks the end of an undo group.
	EndUndoGrouping()
	// A Boolean value that indicates whether the manager automatically creates undo groups around each pass of the run loop.
	GroupsByEvent() bool
	SetGroupsByEvent(value bool)
	// The number of nested undo groups (or redo groups, if redo is the most recent operation) in the current event loop.
	GroupingLevel() int

	// Topic: Enabling and disabling undo

	// Disables the recording of undo operations.
	DisableUndoRegistration()
	// Enables the recording of undo operations.
	EnableUndoRegistration()
	// A Boolean value that indicates whether the recording of undo operations is enabled.
	UndoRegistrationEnabled() bool

	// Topic: Checking whether undo or redo is in process

	// Returns a Boolean value that indicates whether the manager is in the process of performing an undo action.
	Undoing() bool
	// Returns a Boolean value that indicates whether the manager is in the process of performing a redo action.
	Redoing() bool

	// Topic: Clearing undo operations

	// Clears the undo and redo stacks and reenables the manager.
	RemoveAllActions()
	// Clears the undo and redo stacks of all operations involving the specified target as the recipient of the undo message.
	RemoveAllActionsWithTarget(target objectivec.IObject)

	// Topic: Managing the action name

	// The name identifying the undo action.
	UndoActionName() string
	// The name identifying the redo action.
	RedoActionName() string
	// Sets the name of the action associated with the Undo or Redo command.
	SetActionName(actionName string)

	// Topic: Getting and localizing the menu item title

	// The title of the Undo menu command, such as Undo Paste.
	UndoMenuItemTitle() string
	// The title of the Redo menu command, such as Redo Paste.
	RedoMenuItemTitle() string
	// Returns the localized title of the Undo menu command for the identified action.
	UndoMenuTitleForUndoActionName(actionName string) string
	// Returns the localized title of the Redo menu command for the identified action.
	RedoMenuTitleForUndoActionName(actionName string) string

	// Topic: Working with user info

	// Sets a user info value for an undo or redo action.
	SetActionUserInfoValueForKey(info objectivec.IObject, key NSUndoManagerUserInfoKey)
	// Retrieves the undo action’s user info value for the given key.
	UndoActionUserInfoValueForKey(key NSUndoManagerUserInfoKey) objectivec.IObject
	// Retrieves the redo action’s user info value for the given key.
	RedoActionUserInfoValueForKey(key NSUndoManagerUserInfoKey) objectivec.IObject

	// Topic: Working with run loops

	// The modes governing the types of input to handle during a cycle of the run loop.
	RunLoopModes() []string
	SetRunLoopModes(value []string)
	// A priority to use when using a run loop to close an undo group.
	NSUndoCloseGroupingRunLoopOrdering() int
	SetNSUndoCloseGroupingRunLoopOrdering(value int)

	// Topic: Using discardable undo and redo actions

	// Sets whether the next undo or redo action is discardable.
	SetActionIsDiscardable(discardable bool)
	// A Boolean value that indicates whether the next undo action is discardable.
	UndoActionIsDiscardable() bool
	// A Boolean value that indicates whether the next redo action is discardable.
	RedoActionIsDiscardable() bool

	// Topic: Working with notifications

	// Posted just before an undo manager performs an undo operation.
	NSUndoManagerWillUndoChange() NSNotificationName
	// Posted just after an undo manager performs an undo operation.
	NSUndoManagerDidUndoChange() NSNotificationName
	// Posted just before an undo manager performs a redo operation.
	NSUndoManagerWillRedoChange() NSNotificationName
	// Posted just after an undo manager performs a redo operation.
	NSUndoManagerDidRedoChange() NSNotificationName
	// Posted whenever an undo manager opens or closes an undo group (except when it opens a top-level group) and when checking the redo stack.
	NSUndoManagerCheckpoint() NSNotificationName
	// Posted whenever an undo manager opens an undo group.
	NSUndoManagerDidOpenUndoGroup() NSNotificationName
	// Posted before an undo manager closes an undo group.
	NSUndoManagerWillCloseUndoGroup() NSNotificationName
	// Posted after an undo manager closes an undo group.
	NSUndoManagerDidCloseUndoGroup() NSNotificationName
	// A key, used in a notification’s user info, that indicates the undo group contains only discardable actions.
	NSUndoManagerGroupIsDiscardableKey() string

	// Records a single undo operation for a given target so that when the manager performs an undo, it executes the specified block.
	RegisterUndoWithTargetHandler(target objectivec.IObject, undoHandler ObjectHandler)
}

// Init initializes the instance.
func (u UndoManager) Init() UndoManager {
	rv := objc.Send[UndoManager](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UndoManager) Autorelease() UndoManager {
	rv := objc.Send[UndoManager](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUndoManager creates a new UndoManager instance.
func NewUndoManager() UndoManager {
	class := getUndoManagerClass()
	rv := objc.Send[UndoManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Registers the selector of the specified target to implement a single undo
// operation that the target receives.
//
// target: The target of the undo operation.
// 
// The undo manager maintains an unowned reference to `target` to prevent
// retain cycles.
//
// selector: The selector for the undo operation.
//
// object: The argument sent with the selector.
// 
// The undo manager maintains a strong reference to `anO``bject`.
//
// # Discussion
// 
// Use [RegisterUndoWithTargetSelectorObject] to register a selector for an
// undo operation. To register a selector on the undo stack, you also need to
// make the method available to the Objective-C runtime by applying the
// @`objc` attribute to the method. For more on how to create a selector, see
// [Selectors].
// 
// Calling this method also clears the redo stack.
// 
// The following example demonstrates how to register and use a selector on
// the undo stack by modeling a [Garden] class with two methods: `plant()` and
// `pluck()`. The `plant()` method removes a flower from the garden while
// `pluck()` adds a flower such that effectively, the two methods are inverse
// operations of each other. This inverse quality makes it ideal to register
// `plant()` and `pluck()` to be each other’s undo operation.
//
// [Selectors]: https://developer.apple.com/library/archive/documentation/Swift/Conceptual/BuildingCocoaApps/InteractingWithObjective-CAPIs.html#//apple_ref/doc/uid/TP40014216-CH4-ID59
//
// See: https://developer.apple.com/documentation/Foundation/UndoManager/registerUndo(withTarget:selector:object:)
func (u UndoManager) RegisterUndoWithTargetSelectorObject(target objectivec.IObject, selector objc.SEL, object objectivec.IObject) {
	objc.Send[objc.ID](u.ID, objc.Sel("registerUndoWithTarget:selector:object:"), target, selector, object)
}

// Prepares the undo manager for invocation-based undo with the given target
// as the subject of the next undo operation.
//
// target: The target of the undo operation.
// 
// The undo manager maintains a weak reference to `target`.
//
// # Return Value
// 
// A proxy object that forwards messages to the undo manager for recording as
// undo actions.
//
// See: https://developer.apple.com/documentation/Foundation/UndoManager/prepare(withInvocationTarget:)
func (u UndoManager) PrepareWithInvocationTarget(target objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("prepareWithInvocationTarget:"), target)
	return objectivec.Object{ID: rv}
}

// Closes the top-level undo group if necessary, and then performs undo
// operations on the group.
//
// # Discussion
// 
// After closing the top-level undo group, this method invokes
// [UndoNestedGroup].
// 
// This method also invokes [EndUndoGrouping] if the nesting level is 1.
// Raises an [NSInternalInconsistencyException] if more than one undo group is
// open (that is, if the last group isn’t at the top level).
// 
// This method posts an [NSUndoManagerCheckpoint].
//
// [NSUndoManagerCheckpoint]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSUndoManagerCheckpoint
//
// See: https://developer.apple.com/documentation/Foundation/UndoManager/undo()
func (u UndoManager) Undo() {
	objc.Send[objc.ID](u.ID, objc.Sel("undo"))
}

// Performs the undo operations in the last undo group (whether top-level or
// nested), recording the operations on the redo stack as a single group.
//
// # Discussion
// 
// Raises an [NSInternalInconsistencyException] if any undo operations have
// been registered since the last [EnableUndoRegistration] message.
// 
// This method posts an [NSUndoManagerCheckpoint] and
// [NSUndoManagerWillUndoChange] before it performs the undo operation, and it
// posts an [NSUndoManagerDidUndoChange] after it performs the undo operation.
//
// [NSUndoManagerCheckpoint]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSUndoManagerCheckpoint
// [NSUndoManagerDidUndoChange]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSUndoManagerDidUndoChange
// [NSUndoManagerWillUndoChange]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSUndoManagerWillUndoChange
//
// See: https://developer.apple.com/documentation/Foundation/UndoManager/undoNestedGroup()
func (u UndoManager) UndoNestedGroup() {
	objc.Send[objc.ID](u.ID, objc.Sel("undoNestedGroup"))
}

// Performs the operations in the last group on the redo stack, if there are
// any, recording them on the undo stack as a single group.
//
// # Discussion
// 
// Raises an [internalInconsistencyException] if the method is invoked during
// an undo operation.
// 
// This method posts an [NSUndoManagerCheckpoint] and
// [NSUndoManagerWillRedoChange] before it performs the redo operation, and it
// posts the [NSUndoManagerDidRedoChange] after it performs the redo
// operation.
//
// [NSUndoManagerCheckpoint]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSUndoManagerCheckpoint
// [NSUndoManagerDidRedoChange]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSUndoManagerDidRedoChange
// [NSUndoManagerWillRedoChange]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSUndoManagerWillRedoChange
// [internalInconsistencyException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/internalInconsistencyException
//
// See: https://developer.apple.com/documentation/Foundation/UndoManager/redo()
func (u UndoManager) Redo() {
	objc.Send[objc.ID](u.ID, objc.Sel("redo"))
}

// Marks the beginning of an undo group.
//
// # Discussion
// 
// All individual undo operations before a subsequent [EndUndoGrouping]
// message are grouped together and reversed by a later [Undo] message. By
// default undo groups are begun automatically at the start of the event loop,
// but you can begin your own undo groups with this method, and nest them
// within other groups.
// 
// This method posts an [NSUndoManagerCheckpoint] unless a top-level undo is
// in progress. It posts an [NSUndoManagerDidOpenUndoGroup] if a new group was
// successfully created.
//
// [NSUndoManagerCheckpoint]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSUndoManagerCheckpoint
// [NSUndoManagerDidOpenUndoGroup]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSUndoManagerDidOpenUndoGroup
//
// See: https://developer.apple.com/documentation/Foundation/UndoManager/beginUndoGrouping()
func (u UndoManager) BeginUndoGrouping() {
	objc.Send[objc.ID](u.ID, objc.Sel("beginUndoGrouping"))
}

// Marks the end of an undo group.
//
// # Discussion
// 
// All individual undo operations back to the matching [BeginUndoGrouping]
// message are grouped together and reversed by a later [Undo] or
// [UndoNestedGroup] message. Undo groups can be nested, thus providing
// functionality similar to nested transactions. Raises an
// [NSInternalInconsistencyException] if there’s no [BeginUndoGrouping]
// message in effect.
// 
// This method posts an [NSUndoManagerCheckpoint] and an
// [NSUndoManagerDidCloseUndoGroup] just before the group is closed.
//
// [NSUndoManagerCheckpoint]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSUndoManagerCheckpoint
// [NSUndoManagerDidCloseUndoGroup]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSUndoManagerDidCloseUndoGroup
//
// See: https://developer.apple.com/documentation/Foundation/UndoManager/endUndoGrouping()
func (u UndoManager) EndUndoGrouping() {
	objc.Send[objc.ID](u.ID, objc.Sel("endUndoGrouping"))
}

// Disables the recording of undo operations.
//
// # Discussion
// 
// This method disables undos recorded by
// [RegisterUndoWithTargetSelectorObject] or invocation-based undo.
// 
// This method can be invoked multiple times by multiple clients. The
// [EnableUndoRegistration] method must be invoked an equal number of times to
// re-enable undo registration.
//
// See: https://developer.apple.com/documentation/Foundation/UndoManager/disableUndoRegistration()
func (u UndoManager) DisableUndoRegistration() {
	objc.Send[objc.ID](u.ID, objc.Sel("disableUndoRegistration"))
}

// Enables the recording of undo operations.
//
// # Discussion
// 
// Because undo registration is enabled by default, it is often used to
// balance a prior [DisableUndoRegistration] message. Undo registration
// isn’t actually re-enabled until an enable message balances the last
// disable message in effect. Raises an [NSInternalInconsistencyException] if
// invoked while no [DisableUndoRegistration] message is in effect.
//
// See: https://developer.apple.com/documentation/Foundation/UndoManager/enableUndoRegistration()
func (u UndoManager) EnableUndoRegistration() {
	objc.Send[objc.ID](u.ID, objc.Sel("enableUndoRegistration"))
}

// Clears the undo and redo stacks and reenables the manager.
//
// See: https://developer.apple.com/documentation/Foundation/UndoManager/removeAllActions()
func (u UndoManager) RemoveAllActions() {
	objc.Send[objc.ID](u.ID, objc.Sel("removeAllActions"))
}

// Clears the undo and redo stacks of all operations involving the specified
// target as the recipient of the undo message.
//
// target: The recipient of the undo messages to be removed.
//
// # Discussion
// 
// Doesn’t re-enable the manager if it’s disabled.
//
// See: https://developer.apple.com/documentation/Foundation/UndoManager/removeAllActions(withTarget:)
func (u UndoManager) RemoveAllActionsWithTarget(target objectivec.IObject) {
	objc.Send[objc.ID](u.ID, objc.Sel("removeAllActionsWithTarget:"), target)
}

// Sets the name of the action associated with the Undo or Redo command.
//
// actionName: The name of the action.
//
// # Discussion
// 
// If `actionName` is an empty string, the undo manager removes the action
// name currently associated with the menu command.
//
// See: https://developer.apple.com/documentation/Foundation/UndoManager/setActionName(_:)-8lzip
func (u UndoManager) SetActionName(actionName string) {
	objc.Send[objc.ID](u.ID, objc.Sel("setActionName:"), objc.String(actionName))
}

// Returns the localized title of the Undo menu command for the identified
// action.
//
// actionName: The name of the undo action.
//
// # Return Value
// 
// The localized title of the undo menu item.
//
// # Discussion
// 
// Override this method if you want to customize the localization behavior.
// This method is invoked by [UndoMenuItemTitle].
//
// See: https://developer.apple.com/documentation/Foundation/UndoManager/undoMenuTitle(forUndoActionName:)
func (u UndoManager) UndoMenuTitleForUndoActionName(actionName string) string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("undoMenuTitleForUndoActionName:"), objc.String(actionName))
	return NSStringFromID(rv).String()
}

// Returns the localized title of the Redo menu command for the identified
// action.
//
// actionName: The name of the undo action.
//
// # Return Value
// 
// The localized title of the redo menu item.
//
// # Discussion
// 
// Override this method if you want to customize the localization behavior.
// This method is invoked by [RedoMenuItemTitle].
//
// See: https://developer.apple.com/documentation/Foundation/UndoManager/redoMenuTitle(forUndoActionName:)
func (u UndoManager) RedoMenuTitleForUndoActionName(actionName string) string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("redoMenuTitleForUndoActionName:"), objc.String(actionName))
	return NSStringFromID(rv).String()
}

// Sets a user info value for an undo or redo action.
//
// info: The value to save in the action’s user info.
//
// key: The key to associate with the user info value.
//
// # Discussion
// 
// Set user info on an undo action to provide custom information to the action
// beyond its action name. You can use this for things like an icon to
// represent the undoable action, or a timestamp of when the undo manager
// registers the action.
// 
// Start by extending [NSUndoManagerUserInfoKey] with key names to identify
// the user info values you want to associate with undo actions.
// 
// Then set the user info value with this key as part of registering the
// undoable action. In this example, an app’s `insertLayer()` method
// provides a custom icon before setting up an undo action that calls the
// app’s `removeLayer()` method:
//
// See: https://developer.apple.com/documentation/Foundation/UndoManager/setActionUserInfoValue(_:forKey:)
func (u UndoManager) SetActionUserInfoValueForKey(info objectivec.IObject, key NSUndoManagerUserInfoKey) {
	objc.Send[objc.ID](u.ID, objc.Sel("setActionUserInfoValue:forKey:"), info, objc.String(string(key)))
}

// Retrieves the undo action’s user info value for the given key.
//
// key: The key associated with the value to return.
//
// # Return Value
// 
// The value that you previously registered to this key with
// [SetActionUserInfoValueForKey], or `nil` if the key is absent.
//
// # Discussion
// 
// Use this method to retrieve a user info value for the undo action you
// previously set with [SetActionUserInfoValueForKey].
// 
// In this example, an app’s `undoButton()` method provides a SwiftUI view
// that incorporates a previously assigned icon for the undo action:
//
// See: https://developer.apple.com/documentation/Foundation/UndoManager/undoActionUserInfoValue(forKey:)
func (u UndoManager) UndoActionUserInfoValueForKey(key NSUndoManagerUserInfoKey) objectivec.IObject {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("undoActionUserInfoValueForKey:"), objc.String(string(key)))
	return objectivec.Object{ID: rv}
}

// Retrieves the redo action’s user info value for the given key.
//
// key: The key associated with the value to return.
//
// # Return Value
// 
// The value that you previously registered to this key with
// [SetActionUserInfoValueForKey], or `nil` if the key is absent.
//
// # Discussion
// 
// Use this method to retrieve a user info value for the redo action you
// previously set with [SetActionUserInfoValueForKey].
// 
// In this example, an app’s `redoButton()` method provides a SwiftUI view
// that incorporates a previously assigned icon for the action:
//
// See: https://developer.apple.com/documentation/Foundation/UndoManager/redoActionUserInfoValue(forKey:)
func (u UndoManager) RedoActionUserInfoValueForKey(key NSUndoManagerUserInfoKey) objectivec.IObject {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("redoActionUserInfoValueForKey:"), objc.String(string(key)))
	return objectivec.Object{ID: rv}
}

// Sets whether the next undo or redo action is discardable.
//
// discardable: Specifies if the action is discardable. [true] if the next undo or redo
// action can be discarded; [false] otherwise.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Specifies that the latest undo action may be safely discarded when a
// document can not be saved for any reason.
// 
// An example might be an undo action that changes the viewable area of a
// document.
// 
// To find out if an undo group contains only discardable actions, look for
// the [NSUndoManagerGroupIsDiscardableKey] in the [UserInfo] dictionary of
// the [NSUndoManagerWillCloseUndoGroup].
//
// [NSUndoManagerWillCloseUndoGroup]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSUndoManagerWillCloseUndoGroup
//
// See: https://developer.apple.com/documentation/Foundation/UndoManager/setActionIsDiscardable(_:)
func (u UndoManager) SetActionIsDiscardable(discardable bool) {
	objc.Send[objc.ID](u.ID, objc.Sel("setActionIsDiscardable:"), discardable)
}

// Records a single undo operation for a given target so that when the manager
// performs an undo, it executes the specified block.
//
// target: The target of the undo operation.
//
// undoHandler: A block to be executed when an operation is undone.
// 
// The block takes a single argument, the target of the undo operation.
//
// See: https://developer.apple.com/documentation/Foundation/NSUndoManager/registerUndoWithTarget:handler:
func (u UndoManager) RegisterUndoWithTargetHandler(target objectivec.IObject, undoHandler ObjectHandler) {
_block1, _cleanup1 := NewObjectBlock(undoHandler)
	defer _cleanup1()
	objc.Send[objc.ID](u.ID, objc.Sel("registerUndoWithTarget:handler:"), target, _block1)
}

// A Boolean value that indicates whether the manager has any actions to undo.
//
// # Discussion
// 
// [true] if the manager has any actions to undo, otherwise [false].
// 
// The return value doesn’t mean you can safely invoke [Undo] or
// [UndoNestedGroup]—you may have to close open undo groups first.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/UndoManager/canUndo
func (u UndoManager) CanUndo() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("canUndo"))
	return rv
}

// A Boolean value that indicates whether the manager has any actions to redo.
//
// # Discussion
// 
// [true] if the manager has any actions to redo, otherwise [false].
// 
// Because any undo operation registered clears the redo stack, this method
// posts an [NSUndoManagerCheckpoint] to allow clients to apply their pending
// operations before testing the redo stack.
//
// [NSUndoManagerCheckpoint]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSUndoManagerCheckpoint
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/UndoManager/canRedo
func (u UndoManager) CanRedo() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("canRedo"))
	return rv
}

// The maximum number of top-level undo groups the undo manager holds.
//
// # Discussion
// 
// An integer specifying the number of undo groups. A limit of `0` indicates
// no limit, so the manager never drops old undo groups.
// 
// When ending an undo group results in the number of groups exceeding this
// limit, the manager drops the oldest groups from the stack. The default is
// `0`.
// 
// If you change the limit to a level below the prior limit, the manager
// immediately drops old undo groups.
//
// See: https://developer.apple.com/documentation/Foundation/UndoManager/levelsOfUndo
func (u UndoManager) LevelsOfUndo() uint {
	rv := objc.Send[uint](u.ID, objc.Sel("levelsOfUndo"))
	return rv
}
func (u UndoManager) SetLevelsOfUndo(value uint) {
	objc.Send[struct{}](u.ID, objc.Sel("setLevelsOfUndo:"), value)
}

// The number of times you can invoke undo before there are no actions left to
// undo.
//
// # Discussion
// 
// A nonzero value doesn’t imply you can safely invoke [Undo] immediately,
// because you may have to close open undo groups first.
//
// See: https://developer.apple.com/documentation/Foundation/UndoManager/undoCount
func (u UndoManager) UndoCount() uint {
	rv := objc.Send[uint](u.ID, objc.Sel("undoCount"))
	return rv
}

// The number of times you can invoke redo before there are no actions left to
// redo.
//
// See: https://developer.apple.com/documentation/Foundation/UndoManager/redoCount
func (u UndoManager) RedoCount() uint {
	rv := objc.Send[uint](u.ID, objc.Sel("redoCount"))
	return rv
}

// A Boolean value that indicates whether the manager automatically creates
// undo groups around each pass of the run loop.
//
// # Discussion
// 
// [true] if the manager automatically creates undo groups around each pass of
// the run loop, otherwise [false].
// 
// The default is [true]. If you turn automatic grouping off, you must close
// groups explicitly before invoking either [Undo] or [UndoNestedGroup].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/UndoManager/groupsByEvent
func (u UndoManager) GroupsByEvent() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("groupsByEvent"))
	return rv
}
func (u UndoManager) SetGroupsByEvent(value bool) {
	objc.Send[struct{}](u.ID, objc.Sel("setGroupsByEvent:"), value)
}

// The number of nested undo groups (or redo groups, if redo is the most
// recent operation) in the current event loop.
//
// # Discussion
// 
// An integer indicating the number of nested groups. If 0 is returned, there
// is no open undo or redo group.
//
// See: https://developer.apple.com/documentation/Foundation/UndoManager/groupingLevel
func (u UndoManager) GroupingLevel() int {
	rv := objc.Send[int](u.ID, objc.Sel("groupingLevel"))
	return rv
}

// A Boolean value that indicates whether the recording of undo operations is
// enabled.
//
// # Discussion
// 
// [true] if registration is enabled; otherwise, [false].
// 
// The default is [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/UndoManager/isUndoRegistrationEnabled
func (u UndoManager) UndoRegistrationEnabled() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("isUndoRegistrationEnabled"))
	return rv
}

// Returns a Boolean value that indicates whether the manager is in the
// process of performing an undo action.
//
// # Discussion
// 
// The value is [true] if the manager is performing its [Undo] or
// [UndoNestedGroup] method, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/UndoManager/isUndoing
func (u UndoManager) Undoing() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("isUndoing"))
	return rv
}

// Returns a Boolean value that indicates whether the manager is in the
// process of performing a redo action.
//
// # Discussion
// 
// The value is [true] if the manager is performing its [Redo] method,
// otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/UndoManager/isRedoing
func (u UndoManager) Redoing() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("isRedoing"))
	return rv
}

// The name identifying the undo action.
//
// # Discussion
// 
// The undo action name. Returns an empty string (`@""`) if no action name has
// been assigned or if there is nothing to undo.
// 
// For example, if the menu title is “Undo Delete,” the string returned is
// “Delete.”
//
// See: https://developer.apple.com/documentation/Foundation/UndoManager/undoActionName
func (u UndoManager) UndoActionName() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("undoActionName"))
	return NSStringFromID(rv).String()
}

// The name identifying the redo action.
//
// # Discussion
// 
// The redo action name. Returns an empty string (`@""`) if no action name has
// been assigned or if there is nothing to redo.
// 
// For example, if the menu title is “Redo Delete,” the string returned is
// “Delete.”
//
// See: https://developer.apple.com/documentation/Foundation/UndoManager/redoActionName
func (u UndoManager) RedoActionName() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("redoActionName"))
	return NSStringFromID(rv).String()
}

// The title of the Undo menu command, such as Undo Paste.
//
// # Discussion
// 
// Returns “Undo” if no action name has been assigned or `nil` if there is
// nothing to undo.
//
// See: https://developer.apple.com/documentation/Foundation/UndoManager/undoMenuItemTitle
func (u UndoManager) UndoMenuItemTitle() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("undoMenuItemTitle"))
	return NSStringFromID(rv).String()
}

// The title of the Redo menu command, such as Redo Paste.
//
// # Discussion
// 
// Returns “Redo” if no action name has been assigned or `nil` if there is
// nothing to redo.
//
// See: https://developer.apple.com/documentation/Foundation/UndoManager/redoMenuItemTitle
func (u UndoManager) RedoMenuItemTitle() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("redoMenuItemTitle"))
	return NSStringFromID(rv).String()
}

// The modes governing the types of input to handle during a cycle of the run
// loop.
//
// # Discussion
// 
// An array of string constants specifying the current run-loop modes.
// 
// By default, the sole run-loop mode is [NSDefaultRunLoopMode] (which
// excludes data from [NSConnection] objects). Some examples of other uses are
// to limit the input to data received during a mouse-tracking session by
// setting the mode to [NSEventTrackingRunLoopMode], or limit it to data
// received from a modal panel with [NSModalPanelRunLoopMode].
//
// See: https://developer.apple.com/documentation/Foundation/UndoManager/runLoopModes
func (u UndoManager) RunLoopModes() []string {
	rv := objc.Send[[]objc.ID](u.ID, objc.Sel("runLoopModes"))
	return objc.ConvertSliceToStrings(rv)
}
func (u UndoManager) SetRunLoopModes(value []string) {
	objc.Send[struct{}](u.ID, objc.Sel("setRunLoopModes:"), objectivec.StringSliceToNSArray(value))
}

// A priority to use when using a run loop to close an undo group.
//
// See: https://developer.apple.com/documentation/foundation/nsundoclosegroupingrunloopordering
func (u UndoManager) NSUndoCloseGroupingRunLoopOrdering() int {
	rv := objc.Send[int](u.ID, objc.Sel("NSUndoCloseGroupingRunLoopOrdering"))
	return rv
}
func (u UndoManager) SetNSUndoCloseGroupingRunLoopOrdering(value int) {
	objc.Send[struct{}](u.ID, objc.Sel("setNSUndoCloseGroupingRunLoopOrdering:"), value)
}

// A Boolean value that indicates whether the next undo action is discardable.
//
// # Discussion
// 
// [true] if the action is discardable; [false] otherwise.
// 
// Specifies that the latest undo action may be safely discarded when a
// document can not be saved for any reason. These are typically actions that
// don’t affect persistent state.
// 
// An example might be an undo action that changes the viewable area of a
// document.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/UndoManager/undoActionIsDiscardable
func (u UndoManager) UndoActionIsDiscardable() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("undoActionIsDiscardable"))
	return rv
}

// A Boolean value that indicates whether the next redo action is discardable.
//
// # Discussion
// 
// [true] if the action is discardable; [false] otherwise.
// 
// Specifies that the latest redo action may be safely discarded when a
// document can not be saved for any reason. These are typically actions that
// don’t affect persistent state.
// 
// An example might be an redo action that changes the viewable area of a
// document.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/UndoManager/redoActionIsDiscardable
func (u UndoManager) RedoActionIsDiscardable() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("redoActionIsDiscardable"))
	return rv
}

// Posted just before an undo manager performs an undo operation.
//
// See: https://developer.apple.com/documentation/foundation/nsnotification/name-swift.struct/nsundomanagerwillundochange
func (u UndoManager) NSUndoManagerWillUndoChange() NSNotificationName {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("NSUndoManagerWillUndoChangeNotification"))
	return NSNotificationName(NSStringFromID(rv).String())
}

// Posted just after an undo manager performs an undo operation.
//
// See: https://developer.apple.com/documentation/foundation/nsnotification/name-swift.struct/nsundomanagerdidundochange
func (u UndoManager) NSUndoManagerDidUndoChange() NSNotificationName {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("NSUndoManagerDidUndoChangeNotification"))
	return NSNotificationName(NSStringFromID(rv).String())
}

// Posted just before an undo manager performs a redo operation.
//
// See: https://developer.apple.com/documentation/foundation/nsnotification/name-swift.struct/nsundomanagerwillredochange
func (u UndoManager) NSUndoManagerWillRedoChange() NSNotificationName {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("NSUndoManagerWillRedoChangeNotification"))
	return NSNotificationName(NSStringFromID(rv).String())
}

// Posted just after an undo manager performs a redo operation.
//
// See: https://developer.apple.com/documentation/foundation/nsnotification/name-swift.struct/nsundomanagerdidredochange
func (u UndoManager) NSUndoManagerDidRedoChange() NSNotificationName {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("NSUndoManagerDidRedoChangeNotification"))
	return NSNotificationName(NSStringFromID(rv).String())
}

// Posted whenever an undo manager opens or closes an undo group (except when
// it opens a top-level group) and when checking the redo stack.
//
// See: https://developer.apple.com/documentation/foundation/nsnotification/name-swift.struct/nsundomanagercheckpoint
func (u UndoManager) NSUndoManagerCheckpoint() NSNotificationName {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("NSUndoManagerCheckpointNotification"))
	return NSNotificationName(NSStringFromID(rv).String())
}

// Posted whenever an undo manager opens an undo group.
//
// See: https://developer.apple.com/documentation/foundation/nsnotification/name-swift.struct/nsundomanagerdidopenundogroup
func (u UndoManager) NSUndoManagerDidOpenUndoGroup() NSNotificationName {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("NSUndoManagerDidOpenUndoGroupNotification"))
	return NSNotificationName(NSStringFromID(rv).String())
}

// Posted before an undo manager closes an undo group.
//
// See: https://developer.apple.com/documentation/foundation/nsnotification/name-swift.struct/nsundomanagerwillcloseundogroup
func (u UndoManager) NSUndoManagerWillCloseUndoGroup() NSNotificationName {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("NSUndoManagerWillCloseUndoGroupNotification"))
	return NSNotificationName(NSStringFromID(rv).String())
}

// Posted after an undo manager closes an undo group.
//
// See: https://developer.apple.com/documentation/foundation/nsnotification/name-swift.struct/nsundomanagerdidcloseundogroup
func (u UndoManager) NSUndoManagerDidCloseUndoGroup() NSNotificationName {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("NSUndoManagerDidCloseUndoGroupNotification"))
	return NSNotificationName(NSStringFromID(rv).String())
}

// A key, used in a notification’s user info, that indicates the undo group
// contains only discardable actions.
//
// See: https://developer.apple.com/documentation/foundation/nsundomanagergroupisdiscardablekey
func (u UndoManager) NSUndoManagerGroupIsDiscardableKey() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("NSUndoManagerGroupIsDiscardableKey"))
	return NSStringFromID(rv).String()
}

// RegisterUndoWithTargetHandlerSync is a synchronous wrapper around [UndoManager.RegisterUndoWithTargetHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (u UndoManager) RegisterUndoWithTargetHandlerSync(ctx context.Context, target objectivec.IObject) (objectivec.IObject, error) {
	done := make(chan objectivec.IObject, 1)
	u.RegisterUndoWithTargetHandler(target, func(val objectivec.IObject) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

