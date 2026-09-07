// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [BufferSchedulingRecord] class.
var (
	_BufferSchedulingRecordClass     BufferSchedulingRecordClass
	_BufferSchedulingRecordClassOnce sync.Once
)

func getBufferSchedulingRecordClass() BufferSchedulingRecordClass {
	_BufferSchedulingRecordClassOnce.Do(func() {
		_BufferSchedulingRecordClass = BufferSchedulingRecordClass{class: objc.GetClass("_TtCC12TextToSpeech17TTSCoreAudioQueue22BufferSchedulingRecord")}
	})
	return _BufferSchedulingRecordClass
}

// GetBufferSchedulingRecordClass returns the class object for _TtCC12TextToSpeech17TTSCoreAudioQueue22BufferSchedulingRecord.
func GetBufferSchedulingRecordClass() BufferSchedulingRecordClass {
	return getBufferSchedulingRecordClass()
}

type BufferSchedulingRecordClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (bc BufferSchedulingRecordClass) Class() objc.Class {
	return bc.class
}

// Alloc allocates memory for a new instance of the class.
func (bc BufferSchedulingRecordClass) Alloc() BufferSchedulingRecord {
	rv := objc.SendIfResponds[BufferSchedulingRecord](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

type BufferSchedulingRecord struct {
	objectivec.Object
}

// BufferSchedulingRecordFromID constructs a [BufferSchedulingRecord] from an objc.ID.
func BufferSchedulingRecordFromID(id objc.ID) BufferSchedulingRecord {
	return BufferSchedulingRecord{objectivec.Object{ID: id}}
}

// Ensure BufferSchedulingRecord implements IBufferSchedulingRecord.
var _ IBufferSchedulingRecord = BufferSchedulingRecord{}

// An interface definition for the [BufferSchedulingRecord] class.
type IBufferSchedulingRecord interface {
	objectivec.IObject
}

// Init initializes the instance.
func (b BufferSchedulingRecord) Init() BufferSchedulingRecord {
	rv := objc.SendIfResponds[BufferSchedulingRecord](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b BufferSchedulingRecord) Autorelease() BufferSchedulingRecord {
	rv := objc.SendIfResponds[BufferSchedulingRecord](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewBufferSchedulingRecord creates a new BufferSchedulingRecord instance.
func NewBufferSchedulingRecord() BufferSchedulingRecord {
	class := getBufferSchedulingRecordClass()
	rv := objc.SendIfResponds[BufferSchedulingRecord](objc.ID(class.class), objc.Sel("new"))
	return rv
}
