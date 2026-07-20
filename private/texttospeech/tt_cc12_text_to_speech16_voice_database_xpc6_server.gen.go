// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [Server] class.
var (
	_ServerClass     ServerClass
	_ServerClassOnce sync.Once
)

func getServerClass() ServerClass {
	_ServerClassOnce.Do(func() {
		_ServerClass = ServerClass{class: objc.GetClass("_TtCC12TextToSpeech16VoiceDatabaseXPC6Server")}
	})
	return _ServerClass
}

// GetServerClass returns the class object for _TtCC12TextToSpeech16VoiceDatabaseXPC6Server.
func GetServerClass() ServerClass {
	return getServerClass()
}

type ServerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc ServerClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc ServerClass) Alloc() Server {
	rv := objc.Send[Server](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

type Server struct {
	objectivec.Object
}

// ServerFromID constructs a [Server] from an objc.ID.
func ServerFromID(id objc.ID) Server {
	return Server{objectivec.Object{ID: id}}
}

// NOTE: Server struct embeds objectivec.Object (parent type unavailable) but
// IServer embeds the parent interface; skip compile-time assertion.

// An interface definition for the [Server] class.
type IServer interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s Server) Init() Server {
	rv := objc.Send[Server](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s Server) Autorelease() Server {
	rv := objc.Send[Server](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewServer creates a new Server instance.
func NewServer() Server {
	class := getServerClass()
	rv := objc.Send[Server](objc.ID(class.class), objc.Sel("new"))
	return rv
}
