// Code generated from Apple documentation for SoundAnalysis. DO NOT EDIT.

package soundanalysis

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The interface your app implements to receive the results of an analysis request.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNResultsObserving
type SNResultsObserving interface {
	objectivec.IObject

	// Provides a new analysis result to your app with the specified time range.
	//
	// See: https://developer.apple.com/documentation/SoundAnalysis/SNResultsObserving/request(_:didProduce:)
	RequestDidProduceResult(request SNRequest, result SNResult)
}

// SNResultsObservingObject wraps an existing Objective-C object that conforms to the SNResultsObserving protocol.
type SNResultsObservingObject struct {
	objectivec.Object
}

func (o SNResultsObservingObject) BaseObject() objectivec.Object {
	return o.Object
}

// SNResultsObservingObjectFromID constructs a [SNResultsObservingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func SNResultsObservingObjectFromID(id objc.ID) SNResultsObservingObject {
	return SNResultsObservingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Provides a new analysis result to your app with the specified time range.
//
// request: The request that produces the result.
//
// result: The result of the analysis request.
//
// # Discussion
//
// The analyzer calls this function each time a new analysis result is
// available. Different types of analyses may produce results at different
// rates, spanning different time ranges.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNResultsObserving/request(_:didProduce:)
func (o SNResultsObservingObject) RequestDidProduceResult(request SNRequest, result SNResult) {
	objc.Send[struct{}](o.ID, objc.Sel("request:didProduceResult:"), request, result)
}

// Provides any errors that occur during processing of the request.
//
// request: The request that produces the error.
//
// error: The error that occurs during the request.
//
// # Discussion
//
// The analyzer stops processing a specific request when it encounters an
// error, and doesn’t call [RequestDidComplete].
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNResultsObserving/request(_:didFailWithError:)
func (o SNResultsObservingObject) RequestDidFailWithError(request SNRequest, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("request:didFailWithError:"), request, error_)
}

// Notifies your app when the analysis request completes normally.
//
// request: The request that’s completing.
//
// # Discussion
//
// The analyzer calls this method when it finishes processing the request.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNResultsObserving/requestDidComplete(_:)
func (o SNResultsObservingObject) RequestDidComplete(request SNRequest) {
	objc.Send[struct{}](o.ID, objc.Sel("requestDidComplete:"), request)
}
