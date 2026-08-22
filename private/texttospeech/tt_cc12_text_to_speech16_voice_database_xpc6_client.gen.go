// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [Client] class.
var (
	_ClientClass     ClientClass
	_ClientClassOnce sync.Once
)

func getClientClass() ClientClass {
	_ClientClassOnce.Do(func() {
		_ClientClass = ClientClass{class: objc.GetClass("_TtCC12TextToSpeech16VoiceDatabaseXPC6Client")}
	})
	return _ClientClass
}

// GetClientClass returns the class object for _TtCC12TextToSpeech16VoiceDatabaseXPC6Client.
func GetClientClass() ClientClass {
	return getClientClass()
}

type ClientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc ClientClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc ClientClass) Alloc() Client {
	rv := objc.SendIfResponds[Client](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

type Client struct {
	objectivec.Object
}

// ClientFromID constructs a [Client] from an objc.ID.
func ClientFromID(id objc.ID) Client {
	return Client{objectivec.Object{ID: id}}
}

// Ensure Client implements IClient.
var _ IClient = Client{}

// An interface definition for the [Client] class.
type IClient interface {
	objectivec.IObject
}

// Init initializes the instance.
func (c Client) Init() Client {
	rv := objc.SendIfResponds[Client](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c Client) Autorelease() Client {
	rv := objc.SendIfResponds[Client](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewClient creates a new Client instance.
func NewClient() Client {
	class := getClientClass()
	rv := objc.SendIfResponds[Client](objc.ID(class.class), objc.Sel("new"))
	return rv
}
