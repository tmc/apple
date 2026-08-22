// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// The [ODQueryDelegate] protocol defines methods for receiving results returned from an Open Directory query.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODQueryDelegate
type ODQueryDelegate interface {
	objectivec.IObject

	// The delegate method called as results are returned from a query scheduled in a run loop.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/ODQueryDelegate/query(_:foundResults:error:)
	QueryFoundResultsError(inQuery IODQuery, inResults foundation.INSArray, inError foundation.NSError)
}

// ODQueryDelegateObject wraps an existing Objective-C object that conforms to the ODQueryDelegate protocol.
type ODQueryDelegateObject struct {
	objectivec.Object
}

func (o ODQueryDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// ODQueryDelegateObjectFromID constructs a [ODQueryDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func ODQueryDelegateObjectFromID(id objc.ID) ODQueryDelegateObject {
	return ODQueryDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The delegate method called as results are returned from a query scheduled
// in a run loop.
//
// inQuery: The query.
//
// inResults: Partial results returned from the query.
//
// inError: An error reference for error details.
//
// # Discussion
//
// This method is called as soon as any results become available. Results must
// be retained or copied. If both `inResults` and `inError` are `nil`, the
// query has completed.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODQueryDelegate/query(_:foundResults:error:)
func (o ODQueryDelegateObject) QueryFoundResultsError(inQuery IODQuery, inResults foundation.INSArray, inError foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("query:foundResults:error:"), inQuery, inResults, inError)
}

// ODQueryDelegateConfig holds optional typed callbacks for [ODQueryDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/opendirectory/odquerydelegate
type ODQueryDelegateConfig struct {

	// Receiving results from a scheduled query
	// QueryFoundResultsError — The delegate method called as results are returned from a query scheduled in a run loop.
	QueryFoundResultsError func(inQuery ODQuery, inResults foundation.INSArray, inError foundation.NSError)
}

// NewODQueryDelegate creates an Objective-C object implementing the [ODQueryDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [ODQueryDelegateObject] satisfies the [ODQueryDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/opendirectory/odquerydelegate
func NewODQueryDelegate(config ODQueryDelegateConfig) ODQueryDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoODQueryDelegate_%d", n)

	var methods []objc.MethodDef

	if config.QueryFoundResultsError != nil {
		fn := config.QueryFoundResultsError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("query:foundResults:error:"),
			Fn: func(self objc.ID, _cmd objc.SEL, inQueryID objc.ID, inResultsID objc.ID, inErrorID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("ODQueryDelegate", "query:foundResults:error:")
					}
				}()
				inQuery := ODQueryFromID(inQueryID)
				inResults := foundation.NSArrayFromID(inResultsID)
				inError := foundation.NSErrorFromID(inErrorID)
				fn(inQuery, inResults, inError)
				_delegateDone = true
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("ODQueryDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewODQueryDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return ODQueryDelegateObjectFromID(instance)
}
