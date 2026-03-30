// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSRegex] class.
var (
	_TTSRegexClass     TTSRegexClass
	_TTSRegexClassOnce sync.Once
)

func getTTSRegexClass() TTSRegexClass {
	_TTSRegexClassOnce.Do(func() {
		_TTSRegexClass = TTSRegexClass{class: objc.GetClass("TTSRegex")}
	})
	return _TTSRegexClass
}

// GetTTSRegexClass returns the class object for TTSRegex.
func GetTTSRegexClass() TTSRegexClass {
	return getTTSRegexClass()
}

type TTSRegexClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSRegexClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSRegexClass) Alloc() TTSRegex {
	rv := objc.Send[TTSRegex](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [TTSRegex._matchFromOvectorMatchesStringLength]
//   - [TTSRegex.CompiledPCRERegex]
//   - [TTSRegex.SetCompiledPCRERegex]
//   - [TTSRegex.EnumerateMatchesInCStringLengthUsingBlock]
//   - [TTSRegex.EnumerateMatchesInCStringRangesUsingBlock]
//   - [TTSRegex.EnumerateMatchesInCStringStartOffsetLengthUsingBlock]
//   - [TTSRegex.MatchesInCStringLength]
//   - [TTSRegex.InitWithCStringPattern]
//   - [TTSRegex.InitWithCStringPatternOptions]
//   - [TTSRegex.InitWithPattern]
//   - [TTSRegex.InitWithPatternOptions]
//   - [TTSRegex.InitWithPerlPattern]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSRegex
type TTSRegex struct {
	objectivec.Object
}

// TTSRegexFromID constructs a [TTSRegex] from an objc.ID.
func TTSRegexFromID(id objc.ID) TTSRegex {
	return TTSRegex{objectivec.Object{ID: id}}
}

// Ensure TTSRegex implements ITTSRegex.
var _ ITTSRegex = TTSRegex{}

// An interface definition for the [TTSRegex] class.
//
// # Methods
//
//   - [ITTSRegex._matchFromOvectorMatchesStringLength]
//   - [ITTSRegex.CompiledPCRERegex]
//   - [ITTSRegex.SetCompiledPCRERegex]
//   - [ITTSRegex.EnumerateMatchesInCStringLengthUsingBlock]
//   - [ITTSRegex.EnumerateMatchesInCStringRangesUsingBlock]
//   - [ITTSRegex.EnumerateMatchesInCStringStartOffsetLengthUsingBlock]
//   - [ITTSRegex.MatchesInCStringLength]
//   - [ITTSRegex.InitWithCStringPattern]
//   - [ITTSRegex.InitWithCStringPatternOptions]
//   - [ITTSRegex.InitWithPattern]
//   - [ITTSRegex.InitWithPatternOptions]
//   - [ITTSRegex.InitWithPerlPattern]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSRegex
type ITTSRegex interface {
	objectivec.IObject

	// Topic: Methods

	_matchFromOvectorMatchesStringLength(ovector unsafe.Pointer, matches int, string_ string, length uint64) objectivec.IObject
	CompiledPCRERegex() objectivec.IObject
	SetCompiledPCRERegex(value objectivec.IObject)
	EnumerateMatchesInCStringLengthUsingBlock(cString string, length uint64, block VoidHandler)
	EnumerateMatchesInCStringRangesUsingBlock(cString string, ranges objectivec.IObject, block VoidHandler)
	EnumerateMatchesInCStringStartOffsetLengthUsingBlock(cString string, offset uint64, length uint64, block VoidHandler)
	MatchesInCStringLength(cString string, length uint64) objectivec.IObject
	InitWithCStringPattern(pattern string) TTSRegex
	InitWithCStringPatternOptions(pattern string, options uint64) TTSRegex
	InitWithPattern(pattern objectivec.IObject) TTSRegex
	InitWithPatternOptions(pattern objectivec.IObject, options uint64) TTSRegex
	InitWithPerlPattern(pattern objectivec.IObject) TTSRegex
}

// Init initializes the instance.
func (t TTSRegex) Init() TTSRegex {
	rv := objc.Send[TTSRegex](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSRegex) Autorelease() TTSRegex {
	rv := objc.Send[TTSRegex](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSRegex creates a new TTSRegex instance.
func NewTTSRegex() TTSRegex {
	class := getTTSRegexClass()
	rv := objc.Send[TTSRegex](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRegex/initWithCStringPattern:
func NewTTSRegexWithCStringPattern(pattern string) TTSRegex {
	instance := getTTSRegexClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCStringPattern:"), unsafe.Pointer(unsafe.StringData(pattern+"\x00")))
	return TTSRegexFromID(rv)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRegex/initWithCStringPattern:options:
func NewTTSRegexWithCStringPatternOptions(pattern string, options uint64) TTSRegex {
	instance := getTTSRegexClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCStringPattern:options:"), unsafe.Pointer(unsafe.StringData(pattern+"\x00")), options)
	return TTSRegexFromID(rv)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRegex/initWithPattern:
func NewTTSRegexWithPattern(pattern objectivec.IObject) TTSRegex {
	instance := getTTSRegexClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPattern:"), pattern)
	return TTSRegexFromID(rv)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRegex/initWithPattern:options:
func NewTTSRegexWithPatternOptions(pattern objectivec.IObject, options uint64) TTSRegex {
	instance := getTTSRegexClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPattern:options:"), pattern, options)
	return TTSRegexFromID(rv)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRegex/initWithPerlPattern:
func NewTTSRegexWithPerlPattern(pattern objectivec.IObject) TTSRegex {
	instance := getTTSRegexClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPerlPattern:"), pattern)
	return TTSRegexFromID(rv)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRegex/_matchFromOvector:matches:string:length:
func (t TTSRegex) _matchFromOvectorMatchesStringLength(ovector unsafe.Pointer, matches int, string_ string, length uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("_matchFromOvector:matches:string:length:"), ovector, matches, unsafe.Pointer(unsafe.StringData(string_+"\x00")), length)
	return objectivec.Object{ID: rv}
}

// MatchFromOvectorMatchesStringLength is an exported wrapper for the private method _matchFromOvectorMatchesStringLength.
func (t TTSRegex) MatchFromOvectorMatchesStringLength(ovector unsafe.Pointer, matches int, string_ string, length uint64) objectivec.IObject {
	return t._matchFromOvectorMatchesStringLength(ovector, matches, string_, length)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRegex/enumerateMatchesInCString:length:usingBlock:
func (t TTSRegex) EnumerateMatchesInCStringLengthUsingBlock(cString string, length uint64, block VoidHandler) {
	_block2, _ := NewVoidBlock(block)
	objc.Send[objc.ID](t.ID, objc.Sel("enumerateMatchesInCString:length:usingBlock:"), unsafe.Pointer(unsafe.StringData(cString+"\x00")), length, _block2)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRegex/enumerateMatchesInCString:ranges:usingBlock:
func (t TTSRegex) EnumerateMatchesInCStringRangesUsingBlock(cString string, ranges objectivec.IObject, block VoidHandler) {
	_block2, _ := NewVoidBlock(block)
	objc.Send[objc.ID](t.ID, objc.Sel("enumerateMatchesInCString:ranges:usingBlock:"), unsafe.Pointer(unsafe.StringData(cString+"\x00")), ranges, _block2)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRegex/enumerateMatchesInCString:startOffset:length:usingBlock:
func (t TTSRegex) EnumerateMatchesInCStringStartOffsetLengthUsingBlock(cString string, offset uint64, length uint64, block VoidHandler) {
	_block3, _ := NewVoidBlock(block)
	objc.Send[objc.ID](t.ID, objc.Sel("enumerateMatchesInCString:startOffset:length:usingBlock:"), unsafe.Pointer(unsafe.StringData(cString+"\x00")), offset, length, _block3)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRegex/matchesInCString:length:
func (t TTSRegex) MatchesInCStringLength(cString string, length uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("matchesInCString:length:"), unsafe.Pointer(unsafe.StringData(cString+"\x00")), length)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRegex/initWithCStringPattern:
func (t TTSRegex) InitWithCStringPattern(pattern string) TTSRegex {
	rv := objc.Send[TTSRegex](t.ID, objc.Sel("initWithCStringPattern:"), unsafe.Pointer(unsafe.StringData(pattern+"\x00")))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRegex/initWithCStringPattern:options:
func (t TTSRegex) InitWithCStringPatternOptions(pattern string, options uint64) TTSRegex {
	rv := objc.Send[TTSRegex](t.ID, objc.Sel("initWithCStringPattern:options:"), unsafe.Pointer(unsafe.StringData(pattern+"\x00")), options)
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRegex/initWithPattern:
func (t TTSRegex) InitWithPattern(pattern objectivec.IObject) TTSRegex {
	rv := objc.Send[TTSRegex](t.ID, objc.Sel("initWithPattern:"), pattern)
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRegex/initWithPattern:options:
func (t TTSRegex) InitWithPatternOptions(pattern objectivec.IObject, options uint64) TTSRegex {
	rv := objc.Send[TTSRegex](t.ID, objc.Sel("initWithPattern:options:"), pattern, options)
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRegex/initWithPerlPattern:
func (t TTSRegex) InitWithPerlPattern(pattern objectivec.IObject) TTSRegex {
	rv := objc.Send[TTSRegex](t.ID, objc.Sel("initWithPerlPattern:"), pattern)
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSRegex/compiledPCRERegex
func (t TTSRegex) CompiledPCRERegex() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("compiledPCRERegex"))
	return objectivec.Object{ID: rv}
}
func (t TTSRegex) SetCompiledPCRERegex(value objectivec.IObject) {
	objc.Send[struct{}](t.ID, objc.Sel("setCompiledPCRERegex:"), value)
}

// EnumerateMatchesInCStringLengthUsingBlockSync is a synchronous wrapper around [TTSRegex.EnumerateMatchesInCStringLengthUsingBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (t TTSRegex) EnumerateMatchesInCStringLengthUsingBlockSync(ctx context.Context, cString string, length uint64) error {
	done := make(chan struct{}, 1)
	t.EnumerateMatchesInCStringLengthUsingBlock(cString, length, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumerateMatchesInCStringRangesUsingBlockSync is a synchronous wrapper around [TTSRegex.EnumerateMatchesInCStringRangesUsingBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (t TTSRegex) EnumerateMatchesInCStringRangesUsingBlockSync(ctx context.Context, cString string, ranges objectivec.IObject) error {
	done := make(chan struct{}, 1)
	t.EnumerateMatchesInCStringRangesUsingBlock(cString, ranges, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumerateMatchesInCStringStartOffsetLengthUsingBlockSync is a synchronous wrapper around [TTSRegex.EnumerateMatchesInCStringStartOffsetLengthUsingBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (t TTSRegex) EnumerateMatchesInCStringStartOffsetLengthUsingBlockSync(ctx context.Context, cString string, offset uint64, length uint64) error {
	done := make(chan struct{}, 1)
	t.EnumerateMatchesInCStringStartOffsetLengthUsingBlock(cString, offset, length, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
