// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [RequestHandler] class.
var (
	_RequestHandlerClass     RequestHandlerClass
	_RequestHandlerClassOnce sync.Once
)

func getRequestHandlerClass() RequestHandlerClass {
	_RequestHandlerClassOnce.Do(func() {
		_RequestHandlerClass = RequestHandlerClass{class: objc.GetClass("_TtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandler")}
	})
	return _RequestHandlerClass
}

// GetRequestHandlerClass returns the class object for _TtCCC12TextToSpeech16VoiceDatabaseXPC6Server14RequestHandler.
func GetRequestHandlerClass() RequestHandlerClass {
	return getRequestHandlerClass()
}

type RequestHandlerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (rc RequestHandlerClass) Class() objc.Class {
	return rc.class
}

// Alloc allocates memory for a new instance of the class.
func (rc RequestHandlerClass) Alloc() RequestHandler {
	rv := objc.Send[RequestHandler](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

type RequestHandler struct {
	objectivec.Object
}

// RequestHandlerFromID constructs a [RequestHandler] from an objc.ID.
func RequestHandlerFromID(id objc.ID) RequestHandler {
	return RequestHandler{objectivec.Object{ID: id}}
}

// NOTE: RequestHandler struct embeds objectivec.Object (parent type unavailable) but
// IRequestHandler embeds the parent interface; skip compile-time assertion.

// An interface definition for the [RequestHandler] class.
type IRequestHandler interface {
	objectivec.IObject
}

// Init initializes the instance.
func (r RequestHandler) Init() RequestHandler {
	rv := objc.Send[RequestHandler](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r RequestHandler) Autorelease() RequestHandler {
	rv := objc.Send[RequestHandler](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewRequestHandler creates a new RequestHandler instance.
func NewRequestHandler() RequestHandler {
	class := getRequestHandlerClass()
	rv := objc.Send[RequestHandler](objc.ID(class.class), objc.Sel("new"))
	return rv
}
