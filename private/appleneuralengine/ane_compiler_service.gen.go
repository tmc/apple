// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANECompilerService] class.
var (
	_ANECompilerServiceClass     ANECompilerServiceClass
	_ANECompilerServiceClassOnce sync.Once
)

func getANECompilerServiceClass() ANECompilerServiceClass {
	_ANECompilerServiceClassOnce.Do(func() {
		_ANECompilerServiceClass = ANECompilerServiceClass{class: objc.GetClass("_ANECompilerService")}
	})
	return _ANECompilerServiceClass
}

// GetANECompilerServiceClass returns the class object for _ANECompilerService.
func GetANECompilerServiceClass() ANECompilerServiceClass {
	return getANECompilerServiceClass()
}

type ANECompilerServiceClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANECompilerServiceClass) Alloc() ANECompilerService {
	rv := objc.Send[ANECompilerService](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [ANECompilerService.CompileModelAtCsIdentitySandboxExtensionOptionsTempDirectoryCloneDirectoryOutputURLAotModelBinaryPathWithReply]
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECompilerService
type ANECompilerService struct {
	objectivec.Object
}

// ANECompilerServiceFromID constructs a [ANECompilerService] from an objc.ID.
func ANECompilerServiceFromID(id objc.ID) ANECompilerService {
	return ANECompilerService{objectivec.Object{ID: id}}
}
// Ensure ANECompilerService implements IANECompilerService.
var _ IANECompilerService = ANECompilerService{}

// An interface definition for the [ANECompilerService] class.
//
// # Methods
//
//   - [IANECompilerService.CompileModelAtCsIdentitySandboxExtensionOptionsTempDirectoryCloneDirectoryOutputURLAotModelBinaryPathWithReply]
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECompilerService
type IANECompilerService interface {
	objectivec.IObject

	// Topic: Methods

	CompileModelAtCsIdentitySandboxExtensionOptionsTempDirectoryCloneDirectoryOutputURLAotModelBinaryPathWithReply(at objectivec.IObject, identity objectivec.IObject, extension objectivec.IObject, options objectivec.IObject, directory objectivec.IObject, directory2 objectivec.IObject, url foundation.INSURL, path objectivec.IObject, reply VoidHandler)
}

// Init initializes the instance.
func (a ANECompilerService) Init() ANECompilerService {
	rv := objc.Send[ANECompilerService](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANECompilerService) Autorelease() ANECompilerService {
	rv := objc.Send[ANECompilerService](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANECompilerService creates a new ANECompilerService instance.
func NewANECompilerService() ANECompilerService {
	class := getANECompilerServiceClass()
	rv := objc.Send[ANECompilerService](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECompilerService/compileModelAt:csIdentity:sandboxExtension:options:tempDirectory:cloneDirectory:outputURL:aotModelBinaryPath:withReply:
func (a ANECompilerService) CompileModelAtCsIdentitySandboxExtensionOptionsTempDirectoryCloneDirectoryOutputURLAotModelBinaryPathWithReply(at objectivec.IObject, identity objectivec.IObject, extension objectivec.IObject, options objectivec.IObject, directory objectivec.IObject, directory2 objectivec.IObject, url foundation.INSURL, path objectivec.IObject, reply VoidHandler) {
_block8, _cleanup8 := NewVoidBlock(reply)
	defer _cleanup8()
	objc.Send[objc.ID](a.ID, objc.Sel("compileModelAt:csIdentity:sandboxExtension:options:tempDirectory:cloneDirectory:outputURL:aotModelBinaryPath:withReply:"), at, identity, extension, options, directory, directory2, url, path, _block8)
}

// CompileModelAtCsIdentitySandboxExtensionOptionsTempDirectoryCloneDirectoryOutputURLAotModelBinaryPathWithReplySync is a synchronous wrapper around [ANECompilerService.CompileModelAtCsIdentitySandboxExtensionOptionsTempDirectoryCloneDirectoryOutputURLAotModelBinaryPathWithReply].
// It blocks until the completion handler fires or the context is cancelled.
func (a ANECompilerService) CompileModelAtCsIdentitySandboxExtensionOptionsTempDirectoryCloneDirectoryOutputURLAotModelBinaryPathWithReplySync(ctx context.Context, at objectivec.IObject, identity objectivec.IObject, extension objectivec.IObject, options objectivec.IObject, directory objectivec.IObject, directory2 objectivec.IObject, url foundation.INSURL, path objectivec.IObject) error {
	done := make(chan struct{}, 1)
	a.CompileModelAtCsIdentitySandboxExtensionOptionsTempDirectoryCloneDirectoryOutputURLAotModelBinaryPathWithReply(at, identity, extension, options, directory, directory2, url, path, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

