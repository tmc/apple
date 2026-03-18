// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [URLSessionTaskMetrics] class.
var (
	_URLSessionTaskMetricsClass     URLSessionTaskMetricsClass
	_URLSessionTaskMetricsClassOnce sync.Once
)

func getURLSessionTaskMetricsClass() URLSessionTaskMetricsClass {
	_URLSessionTaskMetricsClassOnce.Do(func() {
		_URLSessionTaskMetricsClass = URLSessionTaskMetricsClass{class: objc.GetClass("NSURLSessionTaskMetrics")}
	})
	return _URLSessionTaskMetricsClass
}

// GetURLSessionTaskMetricsClass returns the class object for NSURLSessionTaskMetrics.
func GetURLSessionTaskMetricsClass() URLSessionTaskMetricsClass {
	return getURLSessionTaskMetricsClass()
}

type URLSessionTaskMetricsClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (uc URLSessionTaskMetricsClass) Alloc() URLSessionTaskMetrics {
	rv := objc.Send[URLSessionTaskMetrics](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// An object encapsulating the metrics for a session task.
//
// # Overview
// 
// Each [NSURLSessionTaskMetrics] object contains the [TaskInterval] and
// [RedirectCount], as well as metrics for each request-and-response
// transaction made during the execution of the task.
//
// # Accessing task metrics
//
//   - [URLSessionTaskMetrics.TransactionMetrics]: An array of metrics for each individual request-response transaction made during the execution of the task.
//   - [URLSessionTaskMetrics.TaskInterval]: The time interval between when a task is instantiated and when the task is completed.
//   - [URLSessionTaskMetrics.RedirectCount]: The number of redirects that occurred during the execution of the task.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskMetrics
type URLSessionTaskMetrics struct {
	objectivec.Object
}

// URLSessionTaskMetricsFromID constructs a [URLSessionTaskMetrics] from an objc.ID.
//
// An object encapsulating the metrics for a session task.
func URLSessionTaskMetricsFromID(id objc.ID) URLSessionTaskMetrics {
	return NSURLSessionTaskMetrics{objectivec.Object{ID: id}}
}

// NSURLSessionTaskMetricsFromID is an alias for [URLSessionTaskMetricsFromID] for cross-framework compatibility.
func NSURLSessionTaskMetricsFromID(id objc.ID) URLSessionTaskMetrics { return URLSessionTaskMetricsFromID(id) }
// NOTE: URLSessionTaskMetrics adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [URLSessionTaskMetrics] class.
//
// # Accessing task metrics
//
//   - [IURLSessionTaskMetrics.TransactionMetrics]: An array of metrics for each individual request-response transaction made during the execution of the task.
//   - [IURLSessionTaskMetrics.TaskInterval]: The time interval between when a task is instantiated and when the task is completed.
//   - [IURLSessionTaskMetrics.RedirectCount]: The number of redirects that occurred during the execution of the task.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskMetrics
type IURLSessionTaskMetrics interface {
	objectivec.IObject

	// Topic: Accessing task metrics

	// An array of metrics for each individual request-response transaction made during the execution of the task.
	TransactionMetrics() []NSURLSessionTaskTransactionMetrics
	// The time interval between when a task is instantiated and when the task is completed.
	TaskInterval() INSDateInterval
	// The number of redirects that occurred during the execution of the task.
	RedirectCount() uint
}

// Init initializes the instance.
func (u URLSessionTaskMetrics) Init() URLSessionTaskMetrics {
	rv := objc.Send[URLSessionTaskMetrics](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u URLSessionTaskMetrics) Autorelease() URLSessionTaskMetrics {
	rv := objc.Send[URLSessionTaskMetrics](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLSessionTaskMetrics creates a new URLSessionTaskMetrics instance.
func NewURLSessionTaskMetrics() URLSessionTaskMetrics {
	class := getURLSessionTaskMetricsClass()
	rv := objc.Send[URLSessionTaskMetrics](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// An array of metrics for each individual request-response transaction made
// during the execution of the task.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskMetrics/transactionMetrics
func (u URLSessionTaskMetrics) TransactionMetrics() []NSURLSessionTaskTransactionMetrics {
	rv := objc.Send[[]objc.ID](u.ID, objc.Sel("transactionMetrics"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSURLSessionTaskTransactionMetrics {
		return NSURLSessionTaskTransactionMetricsFromID(id)
	})
}

// The time interval between when a task is instantiated and when the task is
// completed.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskMetrics/taskInterval
func (u URLSessionTaskMetrics) TaskInterval() INSDateInterval {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("taskInterval"))
	return NSDateIntervalFromID(objc.ID(rv))
}

// The number of redirects that occurred during the execution of the task.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskMetrics/redirectCount
func (u URLSessionTaskMetrics) RedirectCount() uint {
	rv := objc.Send[uint](u.ID, objc.Sel("redirectCount"))
	return rv
}

