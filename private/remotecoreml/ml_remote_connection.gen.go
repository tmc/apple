// Code generated from Apple documentation for remotecoreml. DO NOT EDIT.

package remotecoreml

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLRemoteConnection] class.
var (
	_MLRemoteConnectionClass     MLRemoteConnectionClass
	_MLRemoteConnectionClassOnce sync.Once
)

func getMLRemoteConnectionClass() MLRemoteConnectionClass {
	_MLRemoteConnectionClassOnce.Do(func() {
		_MLRemoteConnectionClass = MLRemoteConnectionClass{class: objc.GetClass("_MLRemoteConnection")}
	})
	return _MLRemoteConnectionClass
}

// GetMLRemoteConnectionClass returns the class object for _MLRemoteConnection.
func GetMLRemoteConnectionClass() MLRemoteConnectionClass {
	return getMLRemoteConnectionClass()
}

type MLRemoteConnectionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLRemoteConnectionClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLRemoteConnectionClass) Alloc() MLRemoteConnection {
	rv := objc.Send[MLRemoteConnection](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLRemoteConnection.Connection]
//   - [MLRemoteConnection.DoReceiveContextIsCompleteError]
//   - [MLRemoteConnection.JobCount]
//   - [MLRemoteConnection.LoadFromURLOptionsError]
//   - [MLRemoteConnection.NwObj]
//   - [MLRemoteConnection.NwOptions]
//   - [MLRemoteConnection.OutputResult]
//   - [MLRemoteConnection.SetOutputResult]
//   - [MLRemoteConnection.Packet]
//   - [MLRemoteConnection.PredictionFromURLFeaturesOutputOptionsError]
//   - [MLRemoteConnection.Q]
//   - [MLRemoteConnection.Semaphore]
//   - [MLRemoteConnection.SendOptions]
//   - [MLRemoteConnection.SendDataAndWaitForAcknowledgementOrTimeout]
//   - [MLRemoteConnection.UnloadFromURLOptionsError]
//   - [MLRemoteConnection.InitWithOptions]
//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLRemoteConnection
type MLRemoteConnection struct {
	objectivec.Object
}

// MLRemoteConnectionFromID constructs a [MLRemoteConnection] from an objc.ID.
func MLRemoteConnectionFromID(id objc.ID) MLRemoteConnection {
	return MLRemoteConnection{objectivec.Object{ID: id}}
}

// Ensure MLRemoteConnection implements IMLRemoteConnection.
var _ IMLRemoteConnection = MLRemoteConnection{}

// An interface definition for the [MLRemoteConnection] class.
//
// # Methods
//
//   - [IMLRemoteConnection.Connection]
//   - [IMLRemoteConnection.DoReceiveContextIsCompleteError]
//   - [IMLRemoteConnection.JobCount]
//   - [IMLRemoteConnection.LoadFromURLOptionsError]
//   - [IMLRemoteConnection.NwObj]
//   - [IMLRemoteConnection.NwOptions]
//   - [IMLRemoteConnection.OutputResult]
//   - [IMLRemoteConnection.SetOutputResult]
//   - [IMLRemoteConnection.Packet]
//   - [IMLRemoteConnection.PredictionFromURLFeaturesOutputOptionsError]
//   - [IMLRemoteConnection.Q]
//   - [IMLRemoteConnection.Semaphore]
//   - [IMLRemoteConnection.SendOptions]
//   - [IMLRemoteConnection.SendDataAndWaitForAcknowledgementOrTimeout]
//   - [IMLRemoteConnection.UnloadFromURLOptionsError]
//   - [IMLRemoteConnection.InitWithOptions]
//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLRemoteConnection
type IMLRemoteConnection interface {
	objectivec.IObject

	// Topic: Methods

	Connection() objectivec.Object
	DoReceiveContextIsCompleteError(receive objectivec.IObject, context objectivec.IObject, complete bool, error_ objectivec.IObject)
	JobCount() uint64
	LoadFromURLOptionsError(url foundation.INSURL, options objectivec.IObject) (bool, error)
	NwObj() IMLNetworking
	NwOptions() unsafe.Pointer
	OutputResult() foundation.NSMutableData
	SetOutputResult(value foundation.NSMutableData)
	Packet() IMLNetworkPacket
	PredictionFromURLFeaturesOutputOptionsError(url foundation.INSURL, features objectivec.IObject, output objectivec.IObject, options objectivec.IObject) (bool, error)
	Q() objectivec.Object
	Semaphore() objectivec.Object
	SendOptions(send objectivec.IObject, options objectivec.IObject)
	SendDataAndWaitForAcknowledgementOrTimeout(timeout objectivec.IObject) bool
	UnloadFromURLOptionsError(url foundation.INSURL, options objectivec.IObject) (bool, error)
	InitWithOptions(options objectivec.IObject) MLRemoteConnection
}

// Init initializes the instance.
func (m MLRemoteConnection) Init() MLRemoteConnection {
	rv := objc.Send[MLRemoteConnection](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLRemoteConnection) Autorelease() MLRemoteConnection {
	rv := objc.Send[MLRemoteConnection](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLRemoteConnection creates a new MLRemoteConnection instance.
func NewMLRemoteConnection() MLRemoteConnection {
	class := getMLRemoteConnectionClass()
	rv := objc.Send[MLRemoteConnection](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLRemoteConnection/initWithOptions:
func NewMLRemoteConnectionWithOptions(options objectivec.IObject) MLRemoteConnection {
	instance := getMLRemoteConnectionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithOptions:"), options)
	return MLRemoteConnectionFromID(rv)
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLRemoteConnection/doReceive:context:isComplete:error:
func (m MLRemoteConnection) DoReceiveContextIsCompleteError(receive objectivec.IObject, context objectivec.IObject, complete bool, error_ objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("doReceive:context:isComplete:error:"), receive, context, complete, error_)
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLRemoteConnection/loadFromURL:options:error:
func (m MLRemoteConnection) LoadFromURLOptionsError(url foundation.INSURL, options objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("loadFromURL:options:error:"), url, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("loadFromURL:options:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLRemoteConnection/predictionFromURL:features:output:options:error:
func (m MLRemoteConnection) PredictionFromURLFeaturesOutputOptionsError(url foundation.INSURL, features objectivec.IObject, output objectivec.IObject, options objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("predictionFromURL:features:output:options:error:"), url, features, output, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("predictionFromURL:features:output:options:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLRemoteConnection/send:options:
func (m MLRemoteConnection) SendOptions(send objectivec.IObject, options objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("send:options:"), send, options)
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLRemoteConnection/sendDataAndWaitForAcknowledgementOrTimeout:
func (m MLRemoteConnection) SendDataAndWaitForAcknowledgementOrTimeout(timeout objectivec.IObject) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("sendDataAndWaitForAcknowledgementOrTimeout:"), timeout)
	return rv
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLRemoteConnection/unloadFromURL:options:error:
func (m MLRemoteConnection) UnloadFromURLOptionsError(url foundation.INSURL, options objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("unloadFromURL:options:error:"), url, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("unloadFromURL:options:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLRemoteConnection/initWithOptions:
func (m MLRemoteConnection) InitWithOptions(options objectivec.IObject) MLRemoteConnection {
	rv := objc.Send[MLRemoteConnection](m.ID, objc.Sel("initWithOptions:"), options)
	return rv
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLRemoteConnection/connection
func (m MLRemoteConnection) Connection() objectivec.Object {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("connection"))
	return objectivec.ObjectFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLRemoteConnection/jobCount
func (m MLRemoteConnection) JobCount() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("jobCount"))
	return rv
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLRemoteConnection/nwObj
func (m MLRemoteConnection) NwObj() IMLNetworking {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("nwObj"))
	return MLNetworkingFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLRemoteConnection/nwOptions
func (m MLRemoteConnection) NwOptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m.ID, objc.Sel("nwOptions"))
	return rv
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLRemoteConnection/outputResult
func (m MLRemoteConnection) OutputResult() foundation.NSMutableData {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("outputResult"))
	return foundation.NSMutableDataFromID(objc.ID(rv))
}
func (m MLRemoteConnection) SetOutputResult(value foundation.NSMutableData) {
	objc.Send[struct{}](m.ID, objc.Sel("setOutputResult:"), value)
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLRemoteConnection/packet
func (m MLRemoteConnection) Packet() IMLNetworkPacket {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("packet"))
	return MLNetworkPacketFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLRemoteConnection/q
func (m MLRemoteConnection) Q() objectivec.Object {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("q"))
	return objectivec.ObjectFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLRemoteConnection/semaphore
func (m MLRemoteConnection) Semaphore() objectivec.Object {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("semaphore"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
