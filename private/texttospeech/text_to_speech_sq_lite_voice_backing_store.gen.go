// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechSQLiteVoiceBackingStore] class.
var (
	_TextToSpeechSQLiteVoiceBackingStoreClass     TextToSpeechSQLiteVoiceBackingStoreClass
	_TextToSpeechSQLiteVoiceBackingStoreClassOnce sync.Once
)

func getTextToSpeechSQLiteVoiceBackingStoreClass() TextToSpeechSQLiteVoiceBackingStoreClass {
	_TextToSpeechSQLiteVoiceBackingStoreClassOnce.Do(func() {
		_TextToSpeechSQLiteVoiceBackingStoreClass = TextToSpeechSQLiteVoiceBackingStoreClass{class: objc.GetClass("TextToSpeech.SQLiteVoiceBackingStore")}
	})
	return _TextToSpeechSQLiteVoiceBackingStoreClass
}

// GetTextToSpeechSQLiteVoiceBackingStoreClass returns the class object for TextToSpeech.SQLiteVoiceBackingStore.
func GetTextToSpeechSQLiteVoiceBackingStoreClass() TextToSpeechSQLiteVoiceBackingStoreClass {
	return getTextToSpeechSQLiteVoiceBackingStoreClass()
}

type TextToSpeechSQLiteVoiceBackingStoreClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TextToSpeechSQLiteVoiceBackingStoreClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechSQLiteVoiceBackingStoreClass) Alloc() TextToSpeechSQLiteVoiceBackingStore {
	rv := objc.Send[TextToSpeechSQLiteVoiceBackingStore](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.SQLiteVoiceBackingStore
type TextToSpeechSQLiteVoiceBackingStore struct {
	objectivec.Object
}

// TextToSpeechSQLiteVoiceBackingStoreFromID constructs a [TextToSpeechSQLiteVoiceBackingStore] from an objc.ID.
func TextToSpeechSQLiteVoiceBackingStoreFromID(id objc.ID) TextToSpeechSQLiteVoiceBackingStore {
	return TextToSpeechSQLiteVoiceBackingStore{objectivec.Object{ID: id}}
}
// NOTE: TextToSpeechSQLiteVoiceBackingStore struct embeds objectivec.Object (parent type unavailable) but
// ITextToSpeechSQLiteVoiceBackingStore embeds the parent interface; skip compile-time assertion.

// An interface definition for the [TextToSpeechSQLiteVoiceBackingStore] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.SQLiteVoiceBackingStore
type ITextToSpeechSQLiteVoiceBackingStore interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechSQLiteVoiceBackingStore) Init() TextToSpeechSQLiteVoiceBackingStore {
	rv := objc.Send[TextToSpeechSQLiteVoiceBackingStore](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechSQLiteVoiceBackingStore) Autorelease() TextToSpeechSQLiteVoiceBackingStore {
	rv := objc.Send[TextToSpeechSQLiteVoiceBackingStore](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechSQLiteVoiceBackingStore creates a new TextToSpeechSQLiteVoiceBackingStore instance.
func NewTextToSpeechSQLiteVoiceBackingStore() TextToSpeechSQLiteVoiceBackingStore {
	class := getTextToSpeechSQLiteVoiceBackingStoreClass()
	rv := objc.Send[TextToSpeechSQLiteVoiceBackingStore](objc.ID(class.class), objc.Sel("new"))
	return rv
}

