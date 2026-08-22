// Code generated from Apple documentation for SoundAnalysis. DO NOT EDIT.

package soundanalysis

import (
	"context"
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SNAudioFileAnalyzer] class.
var (
	_SNAudioFileAnalyzerClass     SNAudioFileAnalyzerClass
	_SNAudioFileAnalyzerClassOnce sync.Once
)

func getSNAudioFileAnalyzerClass() SNAudioFileAnalyzerClass {
	_SNAudioFileAnalyzerClassOnce.Do(func() {
		_SNAudioFileAnalyzerClass = SNAudioFileAnalyzerClass{class: objc.GetClass("SNAudioFileAnalyzer")}
	})
	return _SNAudioFileAnalyzerClass
}

// GetSNAudioFileAnalyzerClass returns the class object for SNAudioFileAnalyzer.
func GetSNAudioFileAnalyzerClass() SNAudioFileAnalyzerClass {
	return getSNAudioFileAnalyzerClass()
}

type SNAudioFileAnalyzerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SNAudioFileAnalyzerClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SNAudioFileAnalyzerClass) Alloc() SNAudioFileAnalyzer {
	rv := objc.Send[SNAudioFileAnalyzer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// An analyzer that runs sound classification requests on an audio file.
//
// # Overview
//
// Run an [SNRequest] on an audio file by creating an [SNAudioFileAnalyzer].
// You can run the same sound analysis request on multiple file analyzers, and
// each analyzer can process multiple requests. An audio file analyzer
// generates an [SNResult] each time any of its active requests recognizes a
// sound.
//
// # Creating an Analyzer
//
//   - [SNAudioFileAnalyzer.InitWithURLError]: Creates a new audio file analyzer.
//
// # Managing Requests
//
//   - [SNAudioFileAnalyzer.AddRequestWithObserverError]: Adds a new analysis request to the audio file analyzer.
//   - [SNAudioFileAnalyzer.RemoveRequest]: Removes an existing request from the audio file analyzer.
//   - [SNAudioFileAnalyzer.RemoveAllRequests]: Removes all the sound analysis requests from the audio file analyzer.
//
// # Analyzing Data
//
//   - [SNAudioFileAnalyzer.Analyze]: Analyzes the audio file synchronously.
//   - [SNAudioFileAnalyzer.AnalyzeWithCompletionHandler]: Analyzes the audio file asynchronously.
//   - [SNAudioFileAnalyzer.CancelAnalysis]: Cancels all the asynchronous sound analysis requests the analyzer is currently processing.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNAudioFileAnalyzer
type SNAudioFileAnalyzer struct {
	objectivec.Object
}

// SNAudioFileAnalyzerFromID constructs a [SNAudioFileAnalyzer] from an objc.ID.
//
// An analyzer that runs sound classification requests on an audio file.
func SNAudioFileAnalyzerFromID(id objc.ID) SNAudioFileAnalyzer {
	return SNAudioFileAnalyzer{objectivec.Object{ID: id}}
}

// NOTE: SNAudioFileAnalyzer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SNAudioFileAnalyzer] class.
//
// # Creating an Analyzer
//
//   - [ISNAudioFileAnalyzer.InitWithURLError]: Creates a new audio file analyzer.
//
// # Managing Requests
//
//   - [ISNAudioFileAnalyzer.AddRequestWithObserverError]: Adds a new analysis request to the audio file analyzer.
//   - [ISNAudioFileAnalyzer.RemoveRequest]: Removes an existing request from the audio file analyzer.
//   - [ISNAudioFileAnalyzer.RemoveAllRequests]: Removes all the sound analysis requests from the audio file analyzer.
//
// # Analyzing Data
//
//   - [ISNAudioFileAnalyzer.Analyze]: Analyzes the audio file synchronously.
//   - [ISNAudioFileAnalyzer.AnalyzeWithCompletionHandler]: Analyzes the audio file asynchronously.
//   - [ISNAudioFileAnalyzer.CancelAnalysis]: Cancels all the asynchronous sound analysis requests the analyzer is currently processing.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNAudioFileAnalyzer
type ISNAudioFileAnalyzer interface {
	objectivec.IObject

	// Topic: Creating an Analyzer

	// Creates a new audio file analyzer.
	InitWithURLError(url foundation.NSURL) (SNAudioFileAnalyzer, error)

	// Topic: Managing Requests

	// Adds a new analysis request to the audio file analyzer.
	AddRequestWithObserverError(request SNRequest, observer SNResultsObserving) (bool, error)
	// Removes an existing request from the audio file analyzer.
	RemoveRequest(request SNRequest)
	// Removes all the sound analysis requests from the audio file analyzer.
	RemoveAllRequests()

	// Topic: Analyzing Data

	// Analyzes the audio file synchronously.
	Analyze()
	// Analyzes the audio file asynchronously.
	AnalyzeWithCompletionHandler(completionHandler BoolHandler)
	// Cancels all the asynchronous sound analysis requests the analyzer is currently processing.
	CancelAnalysis()
}

// Init initializes the instance.
func (a SNAudioFileAnalyzer) Init() SNAudioFileAnalyzer {
	rv := objc.Send[SNAudioFileAnalyzer](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a SNAudioFileAnalyzer) Autorelease() SNAudioFileAnalyzer {
	rv := objc.Send[SNAudioFileAnalyzer](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewSNAudioFileAnalyzer creates a new SNAudioFileAnalyzer instance.
func NewSNAudioFileAnalyzer() SNAudioFileAnalyzer {
	class := getSNAudioFileAnalyzerClass()
	rv := objc.Send[SNAudioFileAnalyzer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new audio file analyzer.
//
// url: A path to an audio file.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNAudioFileAnalyzer/init(url:)
func NewAudioFileAnalyzerWithURLError(url foundation.NSURL) (SNAudioFileAnalyzer, error) {
	var errorPtr objc.ID
	instance := getSNAudioFileAnalyzerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SNAudioFileAnalyzer{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return SNAudioFileAnalyzer{}, objc.ErrInitFailed
	}
	return SNAudioFileAnalyzerFromID(rv), nil
}

// Creates a new audio file analyzer.
//
// url: A path to an audio file.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNAudioFileAnalyzer/init(url:)
func (a SNAudioFileAnalyzer) InitWithURLError(url foundation.NSURL) (SNAudioFileAnalyzer, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SNAudioFileAnalyzer{}, foundation.NSErrorFrom(errorPtr)
	}
	return SNAudioFileAnalyzerFromID(rv), nil

}

// Adds a new analysis request to the audio file analyzer.
//
// request: A sound analysis request.
//
// observer: An [SNResultsObserving] instance that receives the analyzer’s results.
// The analyzer maintains a weak reference to the observer.
//
// # Discussion
//
// The method throws an error (Swift) or returns an error (Objective-C) if the
// analyzer is actively processing the file.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNAudioFileAnalyzer/add(_:withObserver:)
func (a SNAudioFileAnalyzer) AddRequestWithObserverError(request SNRequest, observer SNResultsObserving) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("addRequest:withObserver:error:"), request, observer, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("addRequest:withObserver:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Removes an existing request from the audio file analyzer.
//
// request: A sound analysis request.
//
// # Discussion
//
// You can remove a request while the analyzer is processing it. The analyzer
// stops sending results to the observer after the method removes the request.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNAudioFileAnalyzer/remove(_:)
func (a SNAudioFileAnalyzer) RemoveRequest(request SNRequest) {
	objc.Send[objc.ID](a.ID, objc.Sel("removeRequest:"), request)
}

// Removes all the sound analysis requests from the audio file analyzer.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNAudioFileAnalyzer/removeAllRequests()
func (a SNAudioFileAnalyzer) RemoveAllRequests() {
	objc.Send[objc.ID](a.ID, objc.Sel("removeAllRequests"))
}

// Analyzes the audio file synchronously.
//
// # Discussion
//
// This method executes synchronously and may block the calling thread for a
// long time.
//
// The audio file analyzer sends errors to each request’s results observer.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNAudioFileAnalyzer/analyze()
func (a SNAudioFileAnalyzer) Analyze() {
	objc.Send[objc.ID](a.ID, objc.Sel("analyze"))
}

// Analyzes the audio file asynchronously.
//
// completionHandler: A completion closure (Swift) or block (Objective-C) the analyzer calls when
// it finishes analyzing a file.
//
// # Discussion
//
// The method executes asynchronously and calls the completion handler after
// the analyzer finishes analyzing the entire file. The audio file analyzer
// sends errors to each request’s results observer.
//
// If you call the [SNAudioFileAnalyzer.CancelAnalysis] method, the analyzer
// calls your completion handler and passes `false` because it can’t reach
// the end of the file.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNAudioFileAnalyzer/analyze(completionHandler:)
func (a SNAudioFileAnalyzer) AnalyzeWithCompletionHandler(completionHandler BoolHandler) {
	_block0, _ := NewBoolBlock(completionHandler)
	objc.Send[objc.ID](a.ID, objc.Sel("analyzeWithCompletionHandler:"), _block0)
}

// Cancels all the asynchronous sound analysis requests the analyzer is
// currently processing.
//
// # Discussion
//
// The method executes asynchronously, and when it completes, the analyzer
// calls the completion handler you provide to the
// [SNAudioFileAnalyzer.AnalyzeWithCompletionHandler] method.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNAudioFileAnalyzer/cancelAnalysis()
func (a SNAudioFileAnalyzer) CancelAnalysis() {
	objc.Send[objc.ID](a.ID, objc.Sel("cancelAnalysis"))
}

// AnalyzeSync is a synchronous wrapper around [SNAudioFileAnalyzer.AnalyzeWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (a SNAudioFileAnalyzer) AnalyzeSync(ctx context.Context) (bool, error) {
	done := make(chan bool, 1)
	a.AnalyzeWithCompletionHandler(func(val bool) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return false, ctx.Err()
	}
}
