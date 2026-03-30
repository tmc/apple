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
	rv := objc.Send[VZMacAuxiliaryStorage](objc.ID(vc.class), objc.Sel("alloc"))
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
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacAuxiliaryStorage
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
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacAuxiliaryStorage
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
func (m VZMacAuxiliaryStorage) Init() VZMacAuxiliaryStorage {
	rv := objc.Send[VZMacAuxiliaryStorage](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m VZMacAuxiliaryStorage) Autorelease() VZMacAuxiliaryStorage {
	rv := objc.Send[VZMacAuxiliaryStorage](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacAuxiliaryStorage creates a new VZMacAuxiliaryStorage instance.
func NewVZMacAuxiliaryStorage() VZMacAuxiliaryStorage {
	class := getVZMacAuxiliaryStorageClass()
	rv := objc.Send[VZMacAuxiliaryStorage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacAuxiliaryStorage/_allNVRAMDataVariablesInPartition:error:
func (m VZMacAuxiliaryStorage) _allNVRAMDataVariablesInPartitionError(partition uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_allNVRAMDataVariablesInPartition:error:"), partition, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// AllNVRAMDataVariablesInPartitionError is an exported wrapper for the private method _allNVRAMDataVariablesInPartitionError.
func (m VZMacAuxiliaryStorage) AllNVRAMDataVariablesInPartitionError(partition uint64) (objectivec.IObject, error) {
	return m._allNVRAMDataVariablesInPartitionError(partition)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacAuxiliaryStorage/_allNVRAMDataVariablesWithError:
func (m VZMacAuxiliaryStorage) _allNVRAMDataVariablesWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_allNVRAMDataVariablesWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// AllNVRAMDataVariablesWithError is an exported wrapper for the private method _allNVRAMDataVariablesWithError.
func (m VZMacAuxiliaryStorage) AllNVRAMDataVariablesWithError() (objectivec.IObject, error) {
	return m._allNVRAMDataVariablesWithError()
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacAuxiliaryStorage/_allNVRAMVariablesInPartition:error:
func (m VZMacAuxiliaryStorage) _allNVRAMVariablesInPartitionError(partition uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_allNVRAMVariablesInPartition:error:"), partition, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// AllNVRAMVariablesInPartitionError is an exported wrapper for the private method _allNVRAMVariablesInPartitionError.
func (m VZMacAuxiliaryStorage) AllNVRAMVariablesInPartitionError(partition uint64) (objectivec.IObject, error) {
	return m._allNVRAMVariablesInPartitionError(partition)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacAuxiliaryStorage/_allNVRAMVariablesWithError:
func (m VZMacAuxiliaryStorage) _allNVRAMVariablesWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_allNVRAMVariablesWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// AllNVRAMVariablesWithError is an exported wrapper for the private method _allNVRAMVariablesWithError.
func (m VZMacAuxiliaryStorage) AllNVRAMVariablesWithError() (objectivec.IObject, error) {
	return m._allNVRAMVariablesWithError()
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacAuxiliaryStorage/_dataValueForNVRAMVariableNamed:error:
func (m VZMacAuxiliaryStorage) _dataValueForNVRAMVariableNamedError(named objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_dataValueForNVRAMVariableNamed:error:"), named, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// DataValueForNVRAMVariableNamedError is an exported wrapper for the private method _dataValueForNVRAMVariableNamedError.
func (m VZMacAuxiliaryStorage) DataValueForNVRAMVariableNamedError(named objectivec.IObject) (objectivec.IObject, error) {
	return m._dataValueForNVRAMVariableNamedError(named)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacAuxiliaryStorage/_initializeUIDKeyWithWrappingKey:error:
func (m VZMacAuxiliaryStorage) _initializeUIDKeyWithWrappingKeyError(key objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("_initializeUIDKeyWithWrappingKey:error:"), key, unsafe.Pointer(&errorPtr))
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
func (m VZMacAuxiliaryStorage) InitializeUIDKeyWithWrappingKeyError(key objectivec.IObject) (bool, error) {
	return m._initializeUIDKeyWithWrappingKeyError(key)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacAuxiliaryStorage/_removeNVRAMVariableNamed:error:
func (m VZMacAuxiliaryStorage) _removeNVRAMVariableNamedError(named objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("_removeNVRAMVariableNamed:error:"), named, unsafe.Pointer(&errorPtr))
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
func (m VZMacAuxiliaryStorage) RemoveNVRAMVariableNamedError(named objectivec.IObject) (bool, error) {
	return m._removeNVRAMVariableNamedError(named)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacAuxiliaryStorage/_setDataValue:forNVRAMVariableNamed:error:
func (m VZMacAuxiliaryStorage) _setDataValueForNVRAMVariableNamedError(value objectivec.IObject, named objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("_setDataValue:forNVRAMVariableNamed:error:"), value, named, unsafe.Pointer(&errorPtr))
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
func (m VZMacAuxiliaryStorage) SetDataValueForNVRAMVariableNamedError(value objectivec.IObject, named objectivec.IObject) (bool, error) {
	return m._setDataValueForNVRAMVariableNamedError(value, named)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacAuxiliaryStorage/_setValue:forNVRAMVariableNamed:error:
func (m VZMacAuxiliaryStorage) _setValueForNVRAMVariableNamedError(value objectivec.IObject, named objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("_setValue:forNVRAMVariableNamed:error:"), value, named, unsafe.Pointer(&errorPtr))
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
func (m VZMacAuxiliaryStorage) SetValueForNVRAMVariableNamedError(value objectivec.IObject, named objectivec.IObject) (bool, error) {
	return m._setValueForNVRAMVariableNamedError(value, named)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacAuxiliaryStorage/_valueForNVRAMVariableNamed:error:
func (m VZMacAuxiliaryStorage) _valueForNVRAMVariableNamedError(named objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_valueForNVRAMVariableNamed:error:"), named, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// ValueForNVRAMVariableNamedError is an exported wrapper for the private method _valueForNVRAMVariableNamedError.
func (m VZMacAuxiliaryStorage) ValueForNVRAMVariableNamedError(named objectivec.IObject) (objectivec.IObject, error) {
	return m._valueForNVRAMVariableNamedError(named)
}
