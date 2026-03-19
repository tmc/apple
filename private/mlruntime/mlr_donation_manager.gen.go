// Code generated from Apple documentation for mlruntime. DO NOT EDIT.

package mlruntime

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLRDonationManager] class.
var (
	_MLRDonationManagerClass     MLRDonationManagerClass
	_MLRDonationManagerClassOnce sync.Once
)

func getMLRDonationManagerClass() MLRDonationManagerClass {
	_MLRDonationManagerClassOnce.Do(func() {
		_MLRDonationManagerClass = MLRDonationManagerClass{class: objc.GetClass("MLRDonationManager")}
	})
	return _MLRDonationManagerClass
}

// GetMLRDonationManagerClass returns the class object for MLRDonationManager.
func GetMLRDonationManagerClass() MLRDonationManagerClass {
	return getMLRDonationManagerClass()
}

type MLRDonationManagerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLRDonationManagerClass) Alloc() MLRDonationManager {
	rv := objc.Send[MLRDonationManager](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLRDonationManager.EncodeAndUploadToDediscoWithIdentifierMeasurementsWithEncodingSchemasMetadataCompletion]
//   - [MLRDonationManager.Queue]
//   - [MLRDonationManager.SetQueue]
//   - [MLRDonationManager.RecordDataEncodingSchemaMetadataErrorOut]
// See: https://developer.apple.com/documentation/MLRuntime/MLRDonationManager
type MLRDonationManager struct {
	objectivec.Object
}

// MLRDonationManagerFromID constructs a [MLRDonationManager] from an objc.ID.
func MLRDonationManagerFromID(id objc.ID) MLRDonationManager {
	return MLRDonationManager{objectivec.Object{ID: id}}
}
// Ensure MLRDonationManager implements IMLRDonationManager.
var _ IMLRDonationManager = MLRDonationManager{}

// An interface definition for the [MLRDonationManager] class.
//
// # Methods
//
//   - [IMLRDonationManager.EncodeAndUploadToDediscoWithIdentifierMeasurementsWithEncodingSchemasMetadataCompletion]
//   - [IMLRDonationManager.Queue]
//   - [IMLRDonationManager.SetQueue]
//   - [IMLRDonationManager.RecordDataEncodingSchemaMetadataErrorOut]
//
// See: https://developer.apple.com/documentation/MLRuntime/MLRDonationManager
type IMLRDonationManager interface {
	objectivec.IObject

	// Topic: Methods

	EncodeAndUploadToDediscoWithIdentifierMeasurementsWithEncodingSchemasMetadataCompletion(identifier objectivec.IObject, measurements objectivec.IObject, schemas objectivec.IObject, metadata objectivec.IObject, completion ErrorHandler)
	Queue() objectivec.Object
	SetQueue(value objectivec.Object)
	RecordDataEncodingSchemaMetadataErrorOut(record objectivec.IObject, data objectivec.IObject, schema objectivec.IObject, metadata objectivec.IObject, out []objectivec.IObject) bool
}

// Init initializes the instance.
func (r MLRDonationManager) Init() MLRDonationManager {
	rv := objc.Send[MLRDonationManager](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MLRDonationManager) Autorelease() MLRDonationManager {
	rv := objc.Send[MLRDonationManager](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLRDonationManager creates a new MLRDonationManager instance.
func NewMLRDonationManager() MLRDonationManager {
	class := getMLRDonationManagerClass()
	rv := objc.Send[MLRDonationManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRDonationManager/encodeAndUploadToDediscoWithIdentifier:measurements:withEncodingSchemas:metadata:completion:
func (r MLRDonationManager) EncodeAndUploadToDediscoWithIdentifierMeasurementsWithEncodingSchemasMetadataCompletion(identifier objectivec.IObject, measurements objectivec.IObject, schemas objectivec.IObject, metadata objectivec.IObject, completion ErrorHandler) {
_block4, _cleanup4 := NewErrorBlock(completion)
	defer _cleanup4()
	objc.Send[objc.ID](r.ID, objc.Sel("encodeAndUploadToDediscoWithIdentifier:measurements:withEncodingSchemas:metadata:completion:"), identifier, measurements, schemas, metadata, _block4)
}
//
// See: https://developer.apple.com/documentation/MLRuntime/MLRDonationManager/record:data:encodingSchema:metadata:errorOut:
func (r MLRDonationManager) RecordDataEncodingSchemaMetadataErrorOut(record objectivec.IObject, data objectivec.IObject, schema objectivec.IObject, metadata objectivec.IObject, out []objectivec.IObject) bool {
	rv := objc.Send[bool](r.ID, objc.Sel("record:data:encodingSchema:metadata:errorOut:"), record, data, schema, metadata, objectivec.IObjectSliceToNSArray(out))
	return rv
}

// See: https://developer.apple.com/documentation/MLRuntime/MLRDonationManager/defaultManager
func (_MLRDonationManagerClass MLRDonationManagerClass) DefaultManager() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLRDonationManagerClass.class), objc.Sel("defaultManager"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/MLRuntime/MLRDonationManager/queue
func (r MLRDonationManager) Queue() objectivec.Object {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("queue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (r MLRDonationManager) SetQueue(value objectivec.Object) {
	objc.Send[struct{}](r.ID, objc.Sel("setQueue:"), value)
}

