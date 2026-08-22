// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// A protocol that defines methods a delegate object or app extension uses to handle communication from the on-device index.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndexDelegate
type CSSearchableIndexDelegate interface {
	objectivec.IObject

	// Tells the delegate to reindex all searchable data and clear all local state information.
	//
	// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndexDelegate/searchableIndex(_:reindexAllSearchableItemsWithAcknowledgementHandler:)
	SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler(searchableIndex ICSSearchableIndex, acknowledgementHandler VoidHandler)

	// Tells the delegate to reindex the searchable items associated with the specified identifiers.
	//
	// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndexDelegate/searchableIndex(_:reindexSearchableItemsWithIdentifiers:acknowledgementHandler:)
	SearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler(searchableIndex ICSSearchableIndex, identifiers []string, acknowledgementHandler VoidHandler)
}

// CSSearchableIndexDelegateObject wraps an existing Objective-C object that conforms to the CSSearchableIndexDelegate protocol.
type CSSearchableIndexDelegateObject struct {
	objectivec.Object
}

func (o CSSearchableIndexDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// CSSearchableIndexDelegateObjectFromID constructs a [CSSearchableIndexDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CSSearchableIndexDelegateObjectFromID(id objc.ID) CSSearchableIndexDelegateObject {
	return CSSearchableIndexDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate to reindex all searchable data and clear all local state
// information.
//
// searchableIndex: The index in which to reindex the searchable data. The delegate or app
// extension should pass `searchableIndex` to
// [CSSearchableIndex.IndexSearchableItemsCompletionHandler].
//
// acknowledgementHandler: The handler to call after all client state has been saved. Note that if the
// app passes client state information in a batch (for example, by calling
// [CSSearchableIndex.EndIndexBatchWithClientStateCompletionHandler]), the
// acknowledgement handler can be called immediately.
//
// The delegate or app extension must call the acknowledgement handler after
// all client state information has been saved, so that the indexer can call
// this method again in case of a crash.
//
// # Discussion
//
// Typically, the index tells the delegate to reindex its searchable data and
// clear local state when the index has been lost. An app extension should not
// use the index passed in `searchableIndex` when a custom data protection
// class is needed.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndexDelegate/searchableIndex(_:reindexAllSearchableItemsWithAcknowledgementHandler:)
func (o CSSearchableIndexDelegateObject) SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler(searchableIndex ICSSearchableIndex, acknowledgementHandler VoidHandler) {
	_block1, _ := NewVoidBlock(acknowledgementHandler)
	objc.Send[struct{}](o.ID, objc.Sel("searchableIndex:reindexAllSearchableItemsWithAcknowledgementHandler:"), searchableIndex, _block1)
}

// Tells the delegate to reindex the searchable items associated with the
// specified identifiers.
//
// searchableIndex: The index in which to reindex the specified searchable data. To update the
// state of the items, the delegate or app extension should call
// [CSSearchableIndex.IndexSearchableItemsCompletionHandler] passing in
// `searchableIndex`.
//
// identifiers: An array of identifiers that specify searchable items.
//
// acknowledgementHandler: The handler to call after all client state has been saved. Note that if the
// app passes client state information in a batch (for example, by calling
// [CSSearchableIndex.EndIndexBatchWithClientStateCompletionHandler]), the
// acknowledgement handler can be called immediately.
//
// The delegate or app extension must call the acknowledgement handler after
// all client state information has been saved, so that the indexer can call
// this method again in case of a crash.
//
// # Discussion
//
// An app extension should not use the index passed in `searchableIndex` when
// a custom data protection class is needed.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndexDelegate/searchableIndex(_:reindexSearchableItemsWithIdentifiers:acknowledgementHandler:)
func (o CSSearchableIndexDelegateObject) SearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler(searchableIndex ICSSearchableIndex, identifiers []string, acknowledgementHandler VoidHandler) {
	_block2, _ := NewVoidBlock(acknowledgementHandler)
	objc.Send[struct{}](o.ID, objc.Sel("searchableIndex:reindexSearchableItemsWithIdentifiers:acknowledgementHandler:"), searchableIndex, identifiers, _block2)
}

// Tells the delegate that the framework updated the list of searchable items.
//
// items: The items the framework updated.
//
// # Discussion
//
// The framework calls this method when it updates an item with specific
// attributes; see [CSSearchableItem.UpdateListenerOptions] for Apple
// Intelligence attributes.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndexDelegate/searchableItemsDidUpdate(_:)
//
// [CSSearchableItem.UpdateListenerOptions]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/UpdateListenerOptions-swift.struct
func (o CSSearchableIndexDelegateObject) SearchableItemsDidUpdate(items []CSSearchableItem) {
	objc.Send[struct{}](o.ID, objc.Sel("searchableItemsDidUpdate:"), objectivec.IObjectSliceToNSArray(items))
}

// Requests that the delegate provide searchable items for the provided
// identifiers.
//
// identifiers: An array of strings that represent the identifiers.
//
// searchableItemsHandler: A method the framework calls that provides an array of [CSSearchableItem]
// objects.
//
// # Discussion
//
// Use this method to provide the framework with a list of identifiers to
// search for.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndexDelegate/searchableItems(forIdentifiers:searchableItemsHandler:)
func (o CSSearchableIndexDelegateObject) SearchableItemsForIdentifiersSearchableItemsHandler(identifiers []string, searchableItemsHandler CSSearchableItemArrayHandler) {
	_block1, _ := NewCSSearchableItemArrayBlock(searchableItemsHandler)
	objc.Send[struct{}](o.ID, objc.Sel("searchableItemsForIdentifiers:searchableItemsHandler:"), identifiers, _block1)
}

// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndexDelegate/searchableItems(forIdentifiers:protectionClass:searchableItemsHandler:)
func (o CSSearchableIndexDelegateObject) SearchableItemsForIdentifiersProtectionClassSearchableItemsHandler(identifiers []string, protectionClass foundation.NSFileProtectionType, searchableItemsHandler CSSearchableItemArrayHandler) {
	_block2, _ := NewCSSearchableItemArrayBlock(searchableItemsHandler)
	objc.Send[struct{}](o.ID, objc.Sel("searchableItemsForIdentifiers:protectionClass:searchableItemsHandler:"), identifiers, protectionClass, _block2)
}

// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndexDelegate/data(for:itemIdentifier:typeIdentifier:)
func (o CSSearchableIndexDelegateObject) DataForSearchableIndexItemIdentifierTypeIdentifierError(searchableIndex ICSSearchableIndex, itemIdentifier string, typeIdentifier string) (foundation.NSData, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("dataForSearchableIndex:itemIdentifier:typeIdentifier:error:"), searchableIndex, objc.String(itemIdentifier), objc.String(typeIdentifier))
	if err != nil {
		return foundation.NSData{}, err
	}
	return foundation.NSDataFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndexDelegate/fileURL(for:itemIdentifier:typeIdentifier:inPlace:)
func (o CSSearchableIndexDelegateObject) FileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError(searchableIndex ICSSearchableIndex, itemIdentifier string, typeIdentifier string, inPlace bool) (foundation.NSURL, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("fileURLForSearchableIndex:itemIdentifier:typeIdentifier:inPlace:error:"), searchableIndex, objc.String(itemIdentifier), objc.String(typeIdentifier), inPlace)
	if err != nil {
		return foundation.NSURL{}, err
	}
	return foundation.NSURLFromID(rv), nil
}

// Tells the delegate that indexing is being throttled.
//
// searchableIndex: The indexing that’s being throttled.
//
// # Discussion
//
// In some situations, such as when the device is using battery only, the
// system may throttle indexing to save power. You can implement this method
// to be notified of this situation so that you can respond by, for example,
// prioritizing the items to index.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndexDelegate/searchableIndexDidThrottle(_:)
func (o CSSearchableIndexDelegateObject) SearchableIndexDidThrottle(searchableIndex ICSSearchableIndex) {
	objc.Send[struct{}](o.ID, objc.Sel("searchableIndexDidThrottle:"), searchableIndex)
}

// Tells the delegate that the index throttling has finished.
//
// searchableIndex: The index that was throttled.
//
// # Discussion
//
// In some situations, such as when the device is using battery only, the
// system may throttle indexing to save power. You can implement this method
// to be notified when throttling is finished so that your app can resume its
// standard indexing behavior.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndexDelegate/searchableIndexDidFinishThrottle(_:)
func (o CSSearchableIndexDelegateObject) SearchableIndexDidFinishThrottle(searchableIndex ICSSearchableIndex) {
	objc.Send[struct{}](o.ID, objc.Sel("searchableIndexDidFinishThrottle:"), searchableIndex)
}

// CSSearchableIndexDelegateConfig holds optional typed callbacks for [CSSearchableIndexDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/corespotlight/cssearchableindexdelegate
type CSSearchableIndexDelegateConfig struct {

	// Monitoring Spotlight status
	// SearchableIndexDidThrottle — Tells the delegate that indexing is being throttled.
	SearchableIndexDidThrottle func(searchableIndex CSSearchableIndex)
	// SearchableIndexDidFinishThrottle — Tells the delegate that the index throttling has finished.
	SearchableIndexDidFinishThrottle func(searchableIndex CSSearchableIndex)
}

// NewCSSearchableIndexDelegate creates an Objective-C object implementing the [CSSearchableIndexDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [CSSearchableIndexDelegateObject] satisfies the [CSSearchableIndexDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/corespotlight/cssearchableindexdelegate
func NewCSSearchableIndexDelegate(config CSSearchableIndexDelegateConfig) CSSearchableIndexDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoCSSearchableIndexDelegate_%d", n)

	var methods []objc.MethodDef

	if config.SearchableIndexDidThrottle != nil {
		fn := config.SearchableIndexDidThrottle
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("searchableIndexDidThrottle:"),
			Fn: func(self objc.ID, _cmd objc.SEL, searchableIndexID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CSSearchableIndexDelegate", "searchableIndexDidThrottle:")
					}
				}()
				searchableIndex := CSSearchableIndexFromID(searchableIndexID)
				fn(searchableIndex)
				_delegateDone = true
			},
		})
	}

	if config.SearchableIndexDidFinishThrottle != nil {
		fn := config.SearchableIndexDidFinishThrottle
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("searchableIndexDidFinishThrottle:"),
			Fn: func(self objc.ID, _cmd objc.SEL, searchableIndexID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CSSearchableIndexDelegate", "searchableIndexDidFinishThrottle:")
					}
				}()
				searchableIndex := CSSearchableIndexFromID(searchableIndexID)
				fn(searchableIndex)
				_delegateDone = true
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("CSSearchableIndexDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewCSSearchableIndexDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return CSSearchableIndexDelegateObjectFromID(instance)
}
