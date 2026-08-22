// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZMacAuxiliaryStorage] class.
var (
	_VZMacAuxiliaryStorageClass     VZMacAuxiliaryStorageClass
	_VZMacAuxiliaryStorageClassOnce sync.Once
)

func getVZMacAuxiliaryStorageClass() VZMacAuxiliaryStorageClass {
	_VZMacAuxiliaryStorageClassOnce.Do(func() {
		_VZMacAuxiliaryStorageClass = VZMacAuxiliaryStorageClass{class: objc.GetClass("VZMacAuxiliaryStorage")}
	})
	return _VZMacAuxiliaryStorageClass
}

// GetVZMacAuxiliaryStorageClass returns the class object for VZMacAuxiliaryStorage.
func GetVZMacAuxiliaryStorageClass() VZMacAuxiliaryStorageClass {
	return getVZMacAuxiliaryStorageClass()
}

type VZMacAuxiliaryStorageClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMacAuxiliaryStorageClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacAuxiliaryStorageClass) Alloc() VZMacAuxiliaryStorage {
	rv := objc.SendIfResponds[VZMacAuxiliaryStorage](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZMacAuxiliaryStorage._allNVRAMDataVariablesInPartitionError]
//   - [VZMacAuxiliaryStorage._allNVRAMDataVariablesWithError]
//   - [VZMacAuxiliaryStorage._allNVRAMVariablesInPartitionError]
//   - [VZMacAuxiliaryStorage._allNVRAMVariablesWithError]
//   - [VZMacAuxiliaryStorage._dataValueForNVRAMVariableNamedError]
//   - [VZMacAuxiliaryStorage._initializeUIDKeyWithWrappingKeyError]
//   - [VZMacAuxiliaryStorage._removeNVRAMVariableNamedError]
//   - [VZMacAuxiliaryStorage._setDataValueForNVRAMVariableNamedError]
//   - [VZMacAuxiliaryStorage._setValueForNVRAMVariableNamedError]
//   - [VZMacAuxiliaryStorage._valueForNVRAMVariableNamedError]
type VZMacAuxiliaryStorage struct {
	objectivec.Object
}

// VZMacAuxiliaryStorageFromID constructs a [VZMacAuxiliaryStorage] from an objc.ID.
func VZMacAuxiliaryStorageFromID(id objc.ID) VZMacAuxiliaryStorage {
	return VZMacAuxiliaryStorage{objectivec.Object{ID: id}}
}

// Ensure VZMacAuxiliaryStorage implements IVZMacAuxiliaryStorage.
var _ IVZMacAuxiliaryStorage = VZMacAuxiliaryStorage{}

// An interface definition for the [VZMacAuxiliaryStorage] class.
//
// # Methods
//
//   - [IVZMacAuxiliaryStorage._allNVRAMDataVariablesInPartitionError]
//   - [IVZMacAuxiliaryStorage._allNVRAMDataVariablesWithError]
//   - [IVZMacAuxiliaryStorage._allNVRAMVariablesInPartitionError]
//   - [IVZMacAuxiliaryStorage._allNVRAMVariablesWithError]
//   - [IVZMacAuxiliaryStorage._dataValueForNVRAMVariableNamedError]
//   - [IVZMacAuxiliaryStorage._initializeUIDKeyWithWrappingKeyError]
//   - [IVZMacAuxiliaryStorage._removeNVRAMVariableNamedError]
//   - [IVZMacAuxiliaryStorage._setDataValueForNVRAMVariableNamedError]
//   - [IVZMacAuxiliaryStorage._setValueForNVRAMVariableNamedError]
//   - [IVZMacAuxiliaryStorage._valueForNVRAMVariableNamedError]
type IVZMacAuxiliaryStorage interface {
	objectivec.IObject

	// Topic: Methods

	_allNVRAMDataVariablesInPartitionError(partition uint64) (objectivec.IObject, error)
	_allNVRAMDataVariablesWithError() (objectivec.IObject, error)
	_allNVRAMVariablesInPartitionError(partition uint64) (objectivec.IObject, error)
	_allNVRAMVariablesWithError() (objectivec.IObject, error)
	_dataValueForNVRAMVariableNamedError(named objectivec.IObject) (objectivec.IObject, error)
	_initializeUIDKeyWithWrappingKeyError(key objectivec.IObject) (bool, error)
	_removeNVRAMVariableNamedError(named objectivec.IObject) (bool, error)
	_setDataValueForNVRAMVariableNamedError(value objectivec.IObject, named objectivec.IObject) (bool, error)
	_setValueForNVRAMVariableNamedError(value objectivec.IObject, named objectivec.IObject) (bool, error)
	_valueForNVRAMVariableNamedError(named objectivec.IObject) (objectivec.IObject, error)
}

// Init initializes the instance.
func (v VZMacAuxiliaryStorage) Init() VZMacAuxiliaryStorage {
	rv := objc.SendIfResponds[VZMacAuxiliaryStorage](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMacAuxiliaryStorage) Autorelease() VZMacAuxiliaryStorage {
	rv := objc.SendIfResponds[VZMacAuxiliaryStorage](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacAuxiliaryStorage creates a new VZMacAuxiliaryStorage instance.
func NewVZMacAuxiliaryStorage() VZMacAuxiliaryStorage {
	class := getVZMacAuxiliaryStorageClass()
	rv := objc.SendIfResponds[VZMacAuxiliaryStorage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZMacAuxiliaryStorage) _allNVRAMDataVariablesInPartitionError(partition uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_allNVRAMDataVariablesInPartition:error:"), partition, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// AllNVRAMDataVariablesInPartitionError is an exported wrapper for the private method _allNVRAMDataVariablesInPartitionError.
func (v VZMacAuxiliaryStorage) AllNVRAMDataVariablesInPartitionError(partition uint64) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_allNVRAMDataVariablesInPartition:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_allNVRAMDataVariablesInPartition:error:"}
		return nil, err
	}
	return v._allNVRAMDataVariablesInPartitionError(partition)
}

// CanAllNVRAMDataVariablesInPartitionError reports whether the receiver responds to the private selector _allNVRAMDataVariablesInPartition:error:.
func (v VZMacAuxiliaryStorage) CanAllNVRAMDataVariablesInPartitionError() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_allNVRAMDataVariablesInPartition:error:"))
}
func (v VZMacAuxiliaryStorage) _allNVRAMDataVariablesWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_allNVRAMDataVariablesWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// AllNVRAMDataVariablesWithError is an exported wrapper for the private method _allNVRAMDataVariablesWithError.
func (v VZMacAuxiliaryStorage) AllNVRAMDataVariablesWithError() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_allNVRAMDataVariablesWithError:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_allNVRAMDataVariablesWithError:"}
		return nil, err
	}
	return v._allNVRAMDataVariablesWithError()
}

// CanAllNVRAMDataVariablesWithError reports whether the receiver responds to the private selector _allNVRAMDataVariablesWithError:.
func (v VZMacAuxiliaryStorage) CanAllNVRAMDataVariablesWithError() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_allNVRAMDataVariablesWithError:"))
}
func (v VZMacAuxiliaryStorage) _allNVRAMVariablesInPartitionError(partition uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_allNVRAMVariablesInPartition:error:"), partition, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// AllNVRAMVariablesInPartitionError is an exported wrapper for the private method _allNVRAMVariablesInPartitionError.
func (v VZMacAuxiliaryStorage) AllNVRAMVariablesInPartitionError(partition uint64) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_allNVRAMVariablesInPartition:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_allNVRAMVariablesInPartition:error:"}
		return nil, err
	}
	return v._allNVRAMVariablesInPartitionError(partition)
}

// CanAllNVRAMVariablesInPartitionError reports whether the receiver responds to the private selector _allNVRAMVariablesInPartition:error:.
func (v VZMacAuxiliaryStorage) CanAllNVRAMVariablesInPartitionError() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_allNVRAMVariablesInPartition:error:"))
}
func (v VZMacAuxiliaryStorage) _allNVRAMVariablesWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_allNVRAMVariablesWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// AllNVRAMVariablesWithError is an exported wrapper for the private method _allNVRAMVariablesWithError.
func (v VZMacAuxiliaryStorage) AllNVRAMVariablesWithError() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_allNVRAMVariablesWithError:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_allNVRAMVariablesWithError:"}
		return nil, err
	}
	return v._allNVRAMVariablesWithError()
}

// CanAllNVRAMVariablesWithError reports whether the receiver responds to the private selector _allNVRAMVariablesWithError:.
func (v VZMacAuxiliaryStorage) CanAllNVRAMVariablesWithError() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_allNVRAMVariablesWithError:"))
}
func (v VZMacAuxiliaryStorage) _dataValueForNVRAMVariableNamedError(named objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_dataValueForNVRAMVariableNamed:error:"), named, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// DataValueForNVRAMVariableNamedError is an exported wrapper for the private method _dataValueForNVRAMVariableNamedError.
func (v VZMacAuxiliaryStorage) DataValueForNVRAMVariableNamedError(named objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_dataValueForNVRAMVariableNamed:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_dataValueForNVRAMVariableNamed:error:"}
		return nil, err
	}
	return v._dataValueForNVRAMVariableNamedError(named)
}

// CanDataValueForNVRAMVariableNamedError reports whether the receiver responds to the private selector _dataValueForNVRAMVariableNamed:error:.
func (v VZMacAuxiliaryStorage) CanDataValueForNVRAMVariableNamedError() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_dataValueForNVRAMVariableNamed:error:"))
}
func (v VZMacAuxiliaryStorage) _initializeUIDKeyWithWrappingKeyError(key objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("_initializeUIDKeyWithWrappingKey:error:"), key, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_initializeUIDKeyWithWrappingKey:error: returned NO with nil NSError")
	}
	return rv, nil

}

// InitializeUIDKeyWithWrappingKeyError is an exported wrapper for the private method _initializeUIDKeyWithWrappingKeyError.
func (v VZMacAuxiliaryStorage) InitializeUIDKeyWithWrappingKeyError(key objectivec.IObject) (bool, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_initializeUIDKeyWithWrappingKey:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_initializeUIDKeyWithWrappingKey:error:"}
		return false, err
	}
	return v._initializeUIDKeyWithWrappingKeyError(key)
}

// CanInitializeUIDKeyWithWrappingKeyError reports whether the receiver responds to the private selector _initializeUIDKeyWithWrappingKey:error:.
func (v VZMacAuxiliaryStorage) CanInitializeUIDKeyWithWrappingKeyError() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_initializeUIDKeyWithWrappingKey:error:"))
}
func (v VZMacAuxiliaryStorage) _removeNVRAMVariableNamedError(named objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("_removeNVRAMVariableNamed:error:"), named, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_removeNVRAMVariableNamed:error: returned NO with nil NSError")
	}
	return rv, nil

}

// RemoveNVRAMVariableNamedError is an exported wrapper for the private method _removeNVRAMVariableNamedError.
func (v VZMacAuxiliaryStorage) RemoveNVRAMVariableNamedError(named objectivec.IObject) (bool, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_removeNVRAMVariableNamed:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_removeNVRAMVariableNamed:error:"}
		return false, err
	}
	return v._removeNVRAMVariableNamedError(named)
}

// CanRemoveNVRAMVariableNamedError reports whether the receiver responds to the private selector _removeNVRAMVariableNamed:error:.
func (v VZMacAuxiliaryStorage) CanRemoveNVRAMVariableNamedError() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_removeNVRAMVariableNamed:error:"))
}
func (v VZMacAuxiliaryStorage) _setDataValueForNVRAMVariableNamedError(value objectivec.IObject, named objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("_setDataValue:forNVRAMVariableNamed:error:"), value, named, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_setDataValue:forNVRAMVariableNamed:error: returned NO with nil NSError")
	}
	return rv, nil

}

// SetDataValueForNVRAMVariableNamedError is an exported wrapper for the private method _setDataValueForNVRAMVariableNamedError.
func (v VZMacAuxiliaryStorage) SetDataValueForNVRAMVariableNamedError(value objectivec.IObject, named objectivec.IObject) (bool, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setDataValue:forNVRAMVariableNamed:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setDataValue:forNVRAMVariableNamed:error:"}
		return false, err
	}
	return v._setDataValueForNVRAMVariableNamedError(value, named)
}

// CanSetDataValueForNVRAMVariableNamedError reports whether the receiver responds to the private selector _setDataValue:forNVRAMVariableNamed:error:.
func (v VZMacAuxiliaryStorage) CanSetDataValueForNVRAMVariableNamedError() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setDataValue:forNVRAMVariableNamed:error:"))
}
func (v VZMacAuxiliaryStorage) _setValueForNVRAMVariableNamedError(value objectivec.IObject, named objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("_setValue:forNVRAMVariableNamed:error:"), value, named, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_setValue:forNVRAMVariableNamed:error: returned NO with nil NSError")
	}
	return rv, nil

}

// SetValueForNVRAMVariableNamedError is an exported wrapper for the private method _setValueForNVRAMVariableNamedError.
func (v VZMacAuxiliaryStorage) SetValueForNVRAMVariableNamedError(value objectivec.IObject, named objectivec.IObject) (bool, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setValue:forNVRAMVariableNamed:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setValue:forNVRAMVariableNamed:error:"}
		return false, err
	}
	return v._setValueForNVRAMVariableNamedError(value, named)
}

// CanSetValueForNVRAMVariableNamedError reports whether the receiver responds to the private selector _setValue:forNVRAMVariableNamed:error:.
func (v VZMacAuxiliaryStorage) CanSetValueForNVRAMVariableNamedError() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setValue:forNVRAMVariableNamed:error:"))
}
func (v VZMacAuxiliaryStorage) _valueForNVRAMVariableNamedError(named objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_valueForNVRAMVariableNamed:error:"), named, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// ValueForNVRAMVariableNamedError is an exported wrapper for the private method _valueForNVRAMVariableNamedError.
func (v VZMacAuxiliaryStorage) ValueForNVRAMVariableNamedError(named objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_valueForNVRAMVariableNamed:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_valueForNVRAMVariableNamed:error:"}
		return nil, err
	}
	return v._valueForNVRAMVariableNamedError(named)
}

// CanValueForNVRAMVariableNamedError reports whether the receiver responds to the private selector _valueForNVRAMVariableNamed:error:.
func (v VZMacAuxiliaryStorage) CanValueForNVRAMVariableNamedError() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_valueForNVRAMVariableNamed:error:"))
}
