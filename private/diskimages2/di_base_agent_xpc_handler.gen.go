// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [DIBaseAgentXPCHandler] class.
var (
	_DIBaseAgentXPCHandlerClass     DIBaseAgentXPCHandlerClass
	_DIBaseAgentXPCHandlerClassOnce sync.Once
)

func getDIBaseAgentXPCHandlerClass() DIBaseAgentXPCHandlerClass {
	_DIBaseAgentXPCHandlerClassOnce.Do(func() {
		_DIBaseAgentXPCHandlerClass = DIBaseAgentXPCHandlerClass{class: objc.GetClass("DIBaseAgentXPCHandler")}
	})
	return _DIBaseAgentXPCHandlerClass
}

// GetDIBaseAgentXPCHandlerClass returns the class object for DIBaseAgentXPCHandler.
func GetDIBaseAgentXPCHandlerClass() DIBaseAgentXPCHandlerClass {
	return getDIBaseAgentXPCHandlerClass()
}

type DIBaseAgentXPCHandlerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIBaseAgentXPCHandlerClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIBaseAgentXPCHandlerClass) Alloc() DIBaseAgentXPCHandler {
	rv := objc.Send[DIBaseAgentXPCHandler](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [DIBaseAgentXPCHandler.SetConnectionMode]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseAgentXPCHandler
type DIBaseAgentXPCHandler struct {
	DIBaseXPCHandler
}

// DIBaseAgentXPCHandlerFromID constructs a [DIBaseAgentXPCHandler] from an objc.ID.
func DIBaseAgentXPCHandlerFromID(id objc.ID) DIBaseAgentXPCHandler {
	return DIBaseAgentXPCHandler{DIBaseXPCHandler: DIBaseXPCHandlerFromID(id)}
}

// Ensure DIBaseAgentXPCHandler implements IDIBaseAgentXPCHandler.
var _ IDIBaseAgentXPCHandler = DIBaseAgentXPCHandler{}

// An interface definition for the [DIBaseAgentXPCHandler] class.
//
// # Methods
//
//   - [IDIBaseAgentXPCHandler.SetConnectionMode]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseAgentXPCHandler
type IDIBaseAgentXPCHandler interface {
	IDIBaseXPCHandler

	// Topic: Methods

	SetConnectionMode()
}

// Init initializes the instance.
func (d DIBaseAgentXPCHandler) Init() DIBaseAgentXPCHandler {
	rv := objc.Send[DIBaseAgentXPCHandler](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIBaseAgentXPCHandler) Autorelease() DIBaseAgentXPCHandler {
	rv := objc.Send[DIBaseAgentXPCHandler](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIBaseAgentXPCHandler creates a new DIBaseAgentXPCHandler instance.
func NewDIBaseAgentXPCHandler() DIBaseAgentXPCHandler {
	class := getDIBaseAgentXPCHandlerClass()
	rv := objc.Send[DIBaseAgentXPCHandler](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIBaseAgentXPCHandler/setConnectionMode
func (d DIBaseAgentXPCHandler) SetConnectionMode() {
	objc.Send[objc.ID](d.ID, objc.Sel("setConnectionMode"))
}
