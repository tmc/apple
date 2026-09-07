// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSDisplaySyncSessionClient] class.
var (
	_SLSDisplaySyncSessionClientClass     SLSDisplaySyncSessionClientClass
	_SLSDisplaySyncSessionClientClassOnce sync.Once
)

func getSLSDisplaySyncSessionClientClass() SLSDisplaySyncSessionClientClass {
	_SLSDisplaySyncSessionClientClassOnce.Do(func() {
		_SLSDisplaySyncSessionClientClass = SLSDisplaySyncSessionClientClass{class: objc.GetClass("SLSDisplaySyncSessionClient")}
	})
	return _SLSDisplaySyncSessionClientClass
}

// GetSLSDisplaySyncSessionClientClass returns the class object for SLSDisplaySyncSessionClient.
func GetSLSDisplaySyncSessionClientClass() SLSDisplaySyncSessionClientClass {
	return getSLSDisplaySyncSessionClientClass()
}

type SLSDisplaySyncSessionClientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSDisplaySyncSessionClientClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSDisplaySyncSessionClientClass) Alloc() SLSDisplaySyncSessionClient {
	rv := objc.SendIfResponds[SLSDisplaySyncSessionClient](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSDisplaySyncSessionClient.DisplaySyncSessionControl]
//   - [SLSDisplaySyncSessionClient.Service]
//   - [SLSDisplaySyncSessionClient.InitDisplaySyncSessionClient]
type SLSDisplaySyncSessionClient struct {
	SLSDisplayControlClient
}

// SLSDisplaySyncSessionClientFromID constructs a [SLSDisplaySyncSessionClient] from an objc.ID.
func SLSDisplaySyncSessionClientFromID(id objc.ID) SLSDisplaySyncSessionClient {
	return SLSDisplaySyncSessionClient{SLSDisplayControlClient: SLSDisplayControlClientFromID(id)}
}

// Ensure SLSDisplaySyncSessionClient implements ISLSDisplaySyncSessionClient.
var _ ISLSDisplaySyncSessionClient = SLSDisplaySyncSessionClient{}

// An interface definition for the [SLSDisplaySyncSessionClient] class.
//
// # Methods
//
//   - [ISLSDisplaySyncSessionClient.DisplaySyncSessionControl]
//   - [ISLSDisplaySyncSessionClient.Service]
//   - [ISLSDisplaySyncSessionClient.InitDisplaySyncSessionClient]
type ISLSDisplaySyncSessionClient interface {
	ISLSDisplayControlClient

	// Topic: Methods

	DisplaySyncSessionControl() unsafe.Pointer
	Service() ISLSXPCService
	InitDisplaySyncSessionClient(client []objectivec.IObject) SLSDisplaySyncSessionClient
}

// Init initializes the instance.
func (s SLSDisplaySyncSessionClient) Init() SLSDisplaySyncSessionClient {
	rv := objc.SendIfResponds[SLSDisplaySyncSessionClient](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSDisplaySyncSessionClient) Autorelease() SLSDisplaySyncSessionClient {
	rv := objc.SendIfResponds[SLSDisplaySyncSessionClient](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSDisplaySyncSessionClient creates a new SLSDisplaySyncSessionClient instance.
func NewSLSDisplaySyncSessionClient() SLSDisplaySyncSessionClient {
	class := getSLSDisplaySyncSessionClientClass()
	rv := objc.SendIfResponds[SLSDisplaySyncSessionClient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLSDisplaySyncSessionClientDisplaySyncSessionClient(client []objectivec.IObject) SLSDisplaySyncSessionClient {
	instance := getSLSDisplaySyncSessionClientClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initDisplaySyncSessionClient:"), objectivec.IObjectSliceToNSArray(client))
	return SLSDisplaySyncSessionClientFromID(rv)
}

func (s SLSDisplaySyncSessionClient) InitDisplaySyncSessionClient(client []objectivec.IObject) SLSDisplaySyncSessionClient {
	rv := objc.SendIfResponds[SLSDisplaySyncSessionClient](s.ID, objc.Sel("initDisplaySyncSessionClient:"), objectivec.IObjectSliceToNSArray(client))
	return rv
}

func (s SLSDisplaySyncSessionClient) DisplaySyncSessionControl() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](s.ID, objc.Sel("displaySyncSessionControl"))
	return rv
}
func (s SLSDisplaySyncSessionClient) Service() ISLSXPCService {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("service"))
	return SLSXPCServiceFromID(objc.ID(rv))
}
