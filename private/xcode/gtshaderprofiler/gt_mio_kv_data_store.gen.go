// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTMioKVDataStore] class.
var (
	_GTMioKVDataStoreClass     GTMioKVDataStoreClass
	_GTMioKVDataStoreClassOnce sync.Once
)

func getGTMioKVDataStoreClass() GTMioKVDataStoreClass {
	_GTMioKVDataStoreClassOnce.Do(func() {
		_GTMioKVDataStoreClass = GTMioKVDataStoreClass{class: objc.GetClass("GTMioKVDataStore")}
	})
	return _GTMioKVDataStoreClass
}

// GetGTMioKVDataStoreClass returns the class object for GTMioKVDataStore.
func GetGTMioKVDataStoreClass() GTMioKVDataStoreClass {
	return getGTMioKVDataStoreClass()
}

type GTMioKVDataStoreClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioKVDataStoreClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioKVDataStoreClass) Alloc() GTMioKVDataStore {
	rv := objc.Send[GTMioKVDataStore](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [GTMioKVDataStore._blockForName]
//   - [GTMioKVDataStore._enumerateBlocks]
//   - [GTMioKVDataStore._serializeData]
//   - [GTMioKVDataStore._valueForName]
//   - [GTMioKVDataStore.AddChildForKey]
//   - [GTMioKVDataStore.AddChildArrayForKey]
//   - [GTMioKVDataStore.AddDataForKey]
//   - [GTMioKVDataStore.AddDataArrayForKey]
//   - [GTMioKVDataStore.AddMetaForKey]
//   - [GTMioKVDataStore.AddStringArrayForKey]
//   - [GTMioKVDataStore.CompressBlocks]
//   - [GTMioKVDataStore.SetCompressBlocks]
//   - [GTMioKVDataStore.DescriptionWithLevel]
//   - [GTMioKVDataStore.EnumerateBlocks]
//   - [GTMioKVDataStore.GetChild]
//   - [GTMioKVDataStore.GetChildArray]
//   - [GTMioKVDataStore.GetData]
//   - [GTMioKVDataStore.GetDataArray]
//   - [GTMioKVDataStore.GetMeta]
//   - [GTMioKVDataStore.GetStringArray]
//   - [GTMioKVDataStore.Serialize]
//   - [GTMioKVDataStore.SerializeToFile]
//   - [GTMioKVDataStore.InitWithBlockCompression]
//   - [GTMioKVDataStore.InitWithData]
//   - [GTMioKVDataStore.InitWithURL]
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioKVDataStore
type GTMioKVDataStore struct {
	objectivec.Object
}

// GTMioKVDataStoreFromID constructs a [GTMioKVDataStore] from an objc.ID.
func GTMioKVDataStoreFromID(id objc.ID) GTMioKVDataStore {
	return GTMioKVDataStore{objectivec.Object{ID: id}}
}
// Ensure GTMioKVDataStore implements IGTMioKVDataStore.
var _ IGTMioKVDataStore = GTMioKVDataStore{}

// An interface definition for the [GTMioKVDataStore] class.
//
// # Methods
//
//   - [IGTMioKVDataStore._blockForName]
//   - [IGTMioKVDataStore._enumerateBlocks]
//   - [IGTMioKVDataStore._serializeData]
//   - [IGTMioKVDataStore._valueForName]
//   - [IGTMioKVDataStore.AddChildForKey]
//   - [IGTMioKVDataStore.AddChildArrayForKey]
//   - [IGTMioKVDataStore.AddDataForKey]
//   - [IGTMioKVDataStore.AddDataArrayForKey]
//   - [IGTMioKVDataStore.AddMetaForKey]
//   - [IGTMioKVDataStore.AddStringArrayForKey]
//   - [IGTMioKVDataStore.CompressBlocks]
//   - [IGTMioKVDataStore.SetCompressBlocks]
//   - [IGTMioKVDataStore.DescriptionWithLevel]
//   - [IGTMioKVDataStore.EnumerateBlocks]
//   - [IGTMioKVDataStore.GetChild]
//   - [IGTMioKVDataStore.GetChildArray]
//   - [IGTMioKVDataStore.GetData]
//   - [IGTMioKVDataStore.GetDataArray]
//   - [IGTMioKVDataStore.GetMeta]
//   - [IGTMioKVDataStore.GetStringArray]
//   - [IGTMioKVDataStore.Serialize]
//   - [IGTMioKVDataStore.SerializeToFile]
//   - [IGTMioKVDataStore.InitWithBlockCompression]
//   - [IGTMioKVDataStore.InitWithData]
//   - [IGTMioKVDataStore.InitWithURL]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioKVDataStore
type IGTMioKVDataStore interface {
	objectivec.IObject

	// Topic: Methods

	_blockForName(name objectivec.IObject) unsafe.Pointer
	_enumerateBlocks(blocks objectivec.IObject) bool
	_serializeData(_serialize objectivec.IObject, data []objectivec.IObject) bool
	_valueForName(name objectivec.IObject) objectivec.IObject
	AddChildForKey(child objectivec.IObject, key objectivec.IObject) bool
	AddChildArrayForKey(array objectivec.IObject, key objectivec.IObject) bool
	AddDataForKey(data objectivec.IObject, key objectivec.IObject) bool
	AddDataArrayForKey(array objectivec.IObject, key objectivec.IObject) bool
	AddMetaForKey(meta objectivec.IObject, key objectivec.IObject) bool
	AddStringArrayForKey(array objectivec.IObject, key objectivec.IObject) bool
	CompressBlocks() bool
	SetCompressBlocks(value bool)
	DescriptionWithLevel(level uint32) objectivec.IObject
	EnumerateBlocks(blocks objectivec.IObject)
	GetChild(child objectivec.IObject) objectivec.IObject
	GetChildArray(array objectivec.IObject) objectivec.IObject
	GetData(data objectivec.IObject) objectivec.IObject
	GetDataArray(array objectivec.IObject) objectivec.IObject
	GetMeta(meta objectivec.IObject) objectivec.IObject
	GetStringArray(array objectivec.IObject) objectivec.IObject
	Serialize() objectivec.IObject
	SerializeToFile(file objectivec.IObject) bool
	InitWithBlockCompression(compression bool) GTMioKVDataStore
	InitWithData(data objectivec.IObject) GTMioKVDataStore
	InitWithURL(url foundation.INSURL) GTMioKVDataStore
}

// Init initializes the instance.
func (g GTMioKVDataStore) Init() GTMioKVDataStore {
	rv := objc.Send[GTMioKVDataStore](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioKVDataStore) Autorelease() GTMioKVDataStore {
	rv := objc.Send[GTMioKVDataStore](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioKVDataStore creates a new GTMioKVDataStore instance.
func NewGTMioKVDataStore() GTMioKVDataStore {
	class := getGTMioKVDataStoreClass()
	rv := objc.Send[GTMioKVDataStore](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioKVDataStore/initWithBlockCompression:
func NewGTMioKVDataStoreWithBlockCompression(compression bool) GTMioKVDataStore {
	instance := getGTMioKVDataStoreClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBlockCompression:"), compression)
	return GTMioKVDataStoreFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioKVDataStore/initWithData:
func NewGTMioKVDataStoreWithData(data objectivec.IObject) GTMioKVDataStore {
	instance := getGTMioKVDataStoreClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:"), data)
	return GTMioKVDataStoreFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioKVDataStore/initWithURL:
func NewGTMioKVDataStoreWithURL(url foundation.INSURL) GTMioKVDataStore {
	instance := getGTMioKVDataStoreClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:"), url)
	return GTMioKVDataStoreFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioKVDataStore/_blockForName:
func (g GTMioKVDataStore) _blockForName(name objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("_blockForName:"), name)
	return rv
}

// BlockForName is an exported wrapper for the private method _blockForName.
func (g GTMioKVDataStore) BlockForName(name objectivec.IObject) unsafe.Pointer {
	return g._blockForName(name)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioKVDataStore/_enumerateBlocks:
func (g GTMioKVDataStore) _enumerateBlocks(blocks objectivec.IObject) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("_enumerateBlocks:"), blocks)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioKVDataStore/_serialize:data:
func (g GTMioKVDataStore) _serializeData(_serialize objectivec.IObject, data []objectivec.IObject) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("_serialize:data:"), _serialize, objectivec.IObjectSliceToNSArray(data))
	return rv
}

// SerializeData is an exported wrapper for the private method _serializeData.
func (g GTMioKVDataStore) SerializeData(_serialize objectivec.IObject, data []objectivec.IObject) bool {
	return g._serializeData(_serialize, data)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioKVDataStore/_valueForName:
func (g GTMioKVDataStore) _valueForName(name objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("_valueForName:"), name)
	return objectivec.Object{ID: rv}
}

// ValueForName is an exported wrapper for the private method _valueForName.
func (g GTMioKVDataStore) ValueForName(name objectivec.IObject) objectivec.IObject {
	return g._valueForName(name)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioKVDataStore/addChild:forKey:
func (g GTMioKVDataStore) AddChildForKey(child objectivec.IObject, key objectivec.IObject) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("addChild:forKey:"), child, key)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioKVDataStore/addChildArray:forKey:
func (g GTMioKVDataStore) AddChildArrayForKey(array objectivec.IObject, key objectivec.IObject) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("addChildArray:forKey:"), array, key)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioKVDataStore/addData:forKey:
func (g GTMioKVDataStore) AddDataForKey(data objectivec.IObject, key objectivec.IObject) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("addData:forKey:"), data, key)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioKVDataStore/addDataArray:forKey:
func (g GTMioKVDataStore) AddDataArrayForKey(array objectivec.IObject, key objectivec.IObject) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("addDataArray:forKey:"), array, key)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioKVDataStore/addMeta:forKey:
func (g GTMioKVDataStore) AddMetaForKey(meta objectivec.IObject, key objectivec.IObject) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("addMeta:forKey:"), meta, key)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioKVDataStore/addStringArray:forKey:
func (g GTMioKVDataStore) AddStringArrayForKey(array objectivec.IObject, key objectivec.IObject) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("addStringArray:forKey:"), array, key)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioKVDataStore/descriptionWithLevel:
func (g GTMioKVDataStore) DescriptionWithLevel(level uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("descriptionWithLevel:"), level)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioKVDataStore/enumerateBlocks:
func (g GTMioKVDataStore) EnumerateBlocks(blocks objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateBlocks:"), blocks)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioKVDataStore/getChild:
func (g GTMioKVDataStore) GetChild(child objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("getChild:"), child)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioKVDataStore/getChildArray:
func (g GTMioKVDataStore) GetChildArray(array objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("getChildArray:"), array)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioKVDataStore/getData:
func (g GTMioKVDataStore) GetData(data objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("getData:"), data)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioKVDataStore/getDataArray:
func (g GTMioKVDataStore) GetDataArray(array objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("getDataArray:"), array)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioKVDataStore/getMeta:
func (g GTMioKVDataStore) GetMeta(meta objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("getMeta:"), meta)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioKVDataStore/getStringArray:
func (g GTMioKVDataStore) GetStringArray(array objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("getStringArray:"), array)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioKVDataStore/serialize
func (g GTMioKVDataStore) Serialize() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("serialize"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioKVDataStore/serializeToFile:
func (g GTMioKVDataStore) SerializeToFile(file objectivec.IObject) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("serializeToFile:"), file)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioKVDataStore/initWithBlockCompression:
func (g GTMioKVDataStore) InitWithBlockCompression(compression bool) GTMioKVDataStore {
	rv := objc.Send[GTMioKVDataStore](g.ID, objc.Sel("initWithBlockCompression:"), compression)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioKVDataStore/initWithData:
func (g GTMioKVDataStore) InitWithData(data objectivec.IObject) GTMioKVDataStore {
	rv := objc.Send[GTMioKVDataStore](g.ID, objc.Sel("initWithData:"), data)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioKVDataStore/initWithURL:
func (g GTMioKVDataStore) InitWithURL(url foundation.INSURL) GTMioKVDataStore {
	rv := objc.Send[GTMioKVDataStore](g.ID, objc.Sel("initWithURL:"), url)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioKVDataStore/compressBlocks
func (g GTMioKVDataStore) CompressBlocks() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("compressBlocks"))
	return rv
}
func (g GTMioKVDataStore) SetCompressBlocks(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setCompressBlocks:"), value)
}

