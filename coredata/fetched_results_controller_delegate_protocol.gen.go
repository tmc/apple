// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// A delegate protocol that describes the methods that the associated fetched results controller calls when the fetch results change.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchedResultsControllerDelegate
type NSFetchedResultsControllerDelegate interface {
	objectivec.IObject
}

// NSFetchedResultsControllerDelegateObject wraps an existing Objective-C object that conforms to the NSFetchedResultsControllerDelegate protocol.
type NSFetchedResultsControllerDelegateObject struct {
	objectivec.Object
}

func (o NSFetchedResultsControllerDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSFetchedResultsControllerDelegateObjectFromID constructs a [NSFetchedResultsControllerDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSFetchedResultsControllerDelegateObjectFromID(id objc.ID) NSFetchedResultsControllerDelegateObject {
	return NSFetchedResultsControllerDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Notifies the receiver about changes to the content in the fetched results
// controller, by using a diffable data source snapshot.
//
// # Discussion
//
// To apply the changes, call [applySnapshot(_:animatingDifferences:)] on the
// collection or table view’s data source.
//
// If this method is implemented, no other delegate methods are invoked.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchedResultsControllerDelegate/controller(_:didChangeContentWith:)-4kezq
//
// [applySnapshot(_:animatingDifferences:)]: https://developer.apple.com/documentation/UIKit/UITableViewDiffableDataSourceReference/applySnapshot(_:animatingDifferences:)
func (o NSFetchedResultsControllerDelegateObject) ControllerDidChangeContentWithSnapshot(controller INSFetchedResultsController, snapshot objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("controller:didChangeContentWithSnapshot:"), controller, snapshot)
}

// Notifies the receiver about changes to the content in the fetched results
// controller, by using a collection difference.
//
// # Discussion
//
// This method is only invoked if the controller’s
// [NSFetchedResultsController.SectionNameKeyPath] property is `nil` and
// [ControllerDidChangeContentWithDifference] is not implemented.
//
// If this method is implemented, no other delegate methods are invoked.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchedResultsControllerDelegate/controller(_:didChangeContentWith:)-5ullb
func (o NSFetchedResultsControllerDelegateObject) ControllerDidChangeContentWithDifference(controller INSFetchedResultsController, diff foundation.NSOrderedCollectionDifference) {
	objc.Send[struct{}](o.ID, objc.Sel("controller:didChangeContentWithDifference:"), controller, diff)
}

// Notifies the receiver that the fetched results controller is about to start
// processing of one or more changes due to an add, remove, move, or update.
//
// controller: The fetched results controller that sent the message.
//
// # Discussion
//
// This method is invoked before all invocations of
// [ControllerDidChangeObjectAtIndexPathForChangeTypeNewIndexPath] and
// [ControllerDidChangeSectionAtIndexForChangeType] have been sent for a given
// change event (such as the controller receiving a
// [NSManagedObjectContextDidSave] notification).
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchedResultsControllerDelegate/controllerWillChangeContent(_:)
//
// [NSManagedObjectContextDidSave]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSManagedObjectContextDidSave
func (o NSFetchedResultsControllerDelegateObject) ControllerWillChangeContent(controller INSFetchedResultsController) {
	objc.Send[struct{}](o.ID, objc.Sel("controllerWillChangeContent:"), controller)
}

// Notifies the receiver that a fetched object has been changed due to an add,
// remove, move, or update.
//
// controller: The fetched results controller that sent the message.
//
// anObject: The object in controller’s fetched results that changed.
//
// indexPath: The index path of the changed object (this value is `nil` for insertions).
//
// type: The type of change. For valid values see [NSFetchedResultsChangeType].
//
// newIndexPath: The destination path for the object for insertions or moves (this value is
// `nil` for a deletion).
//
// # Discussion
//
// The fetched results controller reports changes to its section before
// changes to the fetch result objects.
//
// Changes are reported with the following heuristics:
//
// - On add and remove operations, only the added/removed object is reported.
//
// It’s assumed that all objects that come after the affected object are
// also moved, but these moves are not reported.
//
// - A move is reported when the changed attribute on the object is one of the
// sort descriptors used in the fetch request.
//
// An update of the object is assumed in this case, but no separate update
// message is sent to the delegate.
//
// - An update is reported when an object’s state changes, but the changed
// attributes aren’t part of the sort keys.
//
// # Special Considerations
//
// This method may be invoked many times during an update event (for example,
// if you are importing data on a background thread and adding them to the
// context in a batch). You should consider carefully whether you want to
// update the table view on receipt of each message.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchedResultsControllerDelegate/controller(_:didChange:at:for:newIndexPath:)
//
// [NSFetchedResultsChangeType]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsChangeType
func (o NSFetchedResultsControllerDelegateObject) ControllerDidChangeObjectAtIndexPathForChangeTypeNewIndexPath(controller INSFetchedResultsController, anObject objectivec.IObject, indexPath foundation.NSIndexPath, type_ NSFetchedResultsChangeType, newIndexPath foundation.NSIndexPath) {
	objc.Send[struct{}](o.ID, objc.Sel("controller:didChangeObject:atIndexPath:forChangeType:newIndexPath:"), controller, anObject, indexPath, type_, newIndexPath)
}

// Notifies the receiver of the addition or removal of a section.
//
// controller: The fetched results controller that sent the message.
//
// sectionInfo: The section that changed.
//
// sectionIndex: The index of the changed section.
//
// type: The type of change (insert or delete). Valid values are
// [NSFetchedResultsChangeType.insert] and
// [NSFetchedResultsChangeType.delete].
//
// # Discussion
//
// The fetched results controller reports changes to its section before
// changes to the fetched result objects.
//
// # Special Considerations
//
// This method may be invoked many times during an update event (for example,
// if you are importing data on a background thread and adding them to the
// context in a batch). You should consider carefully whether you want to
// update the table view on receipt of each message.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchedResultsControllerDelegate/controller(_:didChange:atSectionIndex:for:)
//
// [NSFetchedResultsChangeType.delete]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsChangeType/delete
// [NSFetchedResultsChangeType.insert]: https://developer.apple.com/documentation/CoreData/NSFetchedResultsChangeType/insert
func (o NSFetchedResultsControllerDelegateObject) ControllerDidChangeSectionAtIndexForChangeType(controller INSFetchedResultsController, sectionInfo NSFetchedResultsSectionInfo, sectionIndex uint, type_ NSFetchedResultsChangeType) {
	objc.Send[struct{}](o.ID, objc.Sel("controller:didChangeSection:atIndex:forChangeType:"), controller, sectionInfo, sectionIndex, type_)
}

// Notifies the receiver that the fetched results controller has completed
// processing of one or more changes due to an add, remove, move, or update.
//
// controller: The fetched results controller that sent the message.
//
// # Discussion
//
// This method is invoked after all invocations of
// [ControllerDidChangeObjectAtIndexPathForChangeTypeNewIndexPath] and
// [ControllerDidChangeSectionAtIndexForChangeType] have been sent for a given
// change event (such as the controller receiving a
// [NSManagedObjectContextDidSave] notification).
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchedResultsControllerDelegate/controllerDidChangeContent(_:)
//
// [NSManagedObjectContextDidSave]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSManagedObjectContextDidSave
func (o NSFetchedResultsControllerDelegateObject) ControllerDidChangeContent(controller INSFetchedResultsController) {
	objc.Send[struct{}](o.ID, objc.Sel("controllerDidChangeContent:"), controller)
}

// Returns the name for a given section.
//
// controller: The fetched results controller that sent the message.
//
// sectionName: The default name of the section.
//
// # Return Value
//
// The string to use as the name for the specified section.
//
// # Discussion
//
// This method does not enable change tracking. It is only needed if a section
// index is used.
//
// If the delegate doesn’t implement this method, the default implementation
// returns the capitalized first letter of the section name (see
// [NSFetchedResultsController.SectionIndexTitleForSectionName] in
// [NSFetchedResultsController]).
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchedResultsControllerDelegate/controller(_:sectionIndexTitleForSectionName:)
func (o NSFetchedResultsControllerDelegateObject) ControllerSectionIndexTitleForSectionName(controller INSFetchedResultsController, sectionName string) string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("controller:sectionIndexTitleForSectionName:"), controller, objc.String(sectionName))
	return foundation.NSStringFromID(rv).String()
}

// NSFetchedResultsControllerDelegateConfig holds optional typed callbacks for [NSFetchedResultsControllerDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/coredata/nsfetchedresultscontrollerdelegate
type NSFetchedResultsControllerDelegateConfig struct {

	// Responding to Changes
	// ControllerWillChangeContent — Notifies the receiver that the fetched results controller is about to start processing of one or more changes due to an add, remove, move, or update.
	ControllerWillChangeContent func(controller NSFetchedResultsController)
	// ControllerDidChangeContent — Notifies the receiver that the fetched results controller has completed processing of one or more changes due to an add, remove, move, or update.
	ControllerDidChangeContent func(controller NSFetchedResultsController)

	// Other Methods
	// ControllerDidChangeContentWithDifference — Notifies the receiver about changes to the content in the fetched results controller, by using a collection difference.
	ControllerDidChangeContentWithDifference func(controller NSFetchedResultsController, diff foundation.NSOrderedCollectionDifference)
}

// NewNSFetchedResultsControllerDelegate creates an Objective-C object implementing the [NSFetchedResultsControllerDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSFetchedResultsControllerDelegateObject] satisfies the [NSFetchedResultsControllerDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/coredata/nsfetchedresultscontrollerdelegate
func NewNSFetchedResultsControllerDelegate(config NSFetchedResultsControllerDelegateConfig) NSFetchedResultsControllerDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSFetchedResultsControllerDelegate_%d", n)

	var methods []objc.MethodDef

	if config.ControllerDidChangeContentWithDifference != nil {
		fn := config.ControllerDidChangeContentWithDifference
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("controller:didChangeContentWithDifference:"),
			Fn: func(self objc.ID, _cmd objc.SEL, controllerID objc.ID, diffID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("NSFetchedResultsControllerDelegate", "controller:didChangeContentWithDifference:")
					}
				}()
				controller := NSFetchedResultsControllerFromID(controllerID)
				diff := foundation.NSOrderedCollectionDifferenceFromID(diffID)
				fn(controller, diff)
				_delegateDone = true
			},
		})
	}

	if config.ControllerWillChangeContent != nil {
		fn := config.ControllerWillChangeContent
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("controllerWillChangeContent:"),
			Fn: func(self objc.ID, _cmd objc.SEL, controllerID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("NSFetchedResultsControllerDelegate", "controllerWillChangeContent:")
					}
				}()
				controller := NSFetchedResultsControllerFromID(controllerID)
				fn(controller)
				_delegateDone = true
			},
		})
	}

	if config.ControllerDidChangeContent != nil {
		fn := config.ControllerDidChangeContent
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("controllerDidChangeContent:"),
			Fn: func(self objc.ID, _cmd objc.SEL, controllerID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("NSFetchedResultsControllerDelegate", "controllerDidChangeContent:")
					}
				}()
				controller := NSFetchedResultsControllerFromID(controllerID)
				fn(controller)
				_delegateDone = true
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSFetchedResultsControllerDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSFetchedResultsControllerDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSFetchedResultsControllerDelegateObjectFromID(instance)
}
