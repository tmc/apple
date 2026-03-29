// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLModelCollection] class.
var (
	_MLModelCollectionClass     MLModelCollectionClass
	_MLModelCollectionClassOnce sync.Once
)

func getMLModelCollectionClass() MLModelCollectionClass {
	_MLModelCollectionClassOnce.Do(func() {
		_MLModelCollectionClass = MLModelCollectionClass{class: objc.GetClass("MLModelCollection")}
	})
	return _MLModelCollectionClass
}

// GetMLModelCollectionClass returns the class object for MLModelCollection.
func GetMLModelCollectionClass() MLModelCollectionClass {
	return getMLModelCollectionClass()
}

type MLModelCollectionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLModelCollectionClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelCollectionClass) Alloc() MLModelCollection {
	rv := objc.Send[MLModelCollection](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLModelCollection._downloadOptions]
//   - [MLModelCollection._downloadWithProgress]
//   - [MLModelCollection._endAccess]
//   - [MLModelCollection._handleTrialUpdateForNamespaceName]
//   - [MLModelCollection._populateEntries]
//   - [MLModelCollection._register]
//   - [MLModelCollection._registerForUpdates]
//   - [MLModelCollection._setDeploymentID]
//   - [MLModelCollection.DeploymentID]
//   - [MLModelCollection.SetDeploymentID]
//   - [MLModelCollection.Entries]
//   - [MLModelCollection.SetEntries]
//   - [MLModelCollection.Identifier]
//   - [MLModelCollection.NamespaceName]
//   - [MLModelCollection.TrialClient]
//   - [MLModelCollection.InitWithIdentifier]
// See: https://developer.apple.com/documentation/CoreML/MLModelCollection
type MLModelCollection struct {
	objectivec.Object
}

// MLModelCollectionFromID constructs a [MLModelCollection] from an objc.ID.
func MLModelCollectionFromID(id objc.ID) MLModelCollection {
	return MLModelCollection{objectivec.Object{ID: id}}
}
// Ensure MLModelCollection implements IMLModelCollection.
var _ IMLModelCollection = MLModelCollection{}

// An interface definition for the [MLModelCollection] class.
//
// # Methods
//
//   - [IMLModelCollection._downloadOptions]
//   - [IMLModelCollection._downloadWithProgress]
//   - [IMLModelCollection._endAccess]
//   - [IMLModelCollection._handleTrialUpdateForNamespaceName]
//   - [IMLModelCollection._populateEntries]
//   - [IMLModelCollection._register]
//   - [IMLModelCollection._registerForUpdates]
//   - [IMLModelCollection._setDeploymentID]
//   - [IMLModelCollection.DeploymentID]
//   - [IMLModelCollection.SetDeploymentID]
//   - [IMLModelCollection.Entries]
//   - [IMLModelCollection.SetEntries]
//   - [IMLModelCollection.Identifier]
//   - [IMLModelCollection.NamespaceName]
//   - [IMLModelCollection.TrialClient]
//   - [IMLModelCollection.InitWithIdentifier]
//
// See: https://developer.apple.com/documentation/CoreML/MLModelCollection
type IMLModelCollection interface {
	objectivec.IObject

	// Topic: Methods

	_downloadOptions() objectivec.IObject
	_downloadWithProgress(progress VoidHandler) bool
	_endAccess() bool
	_handleTrialUpdateForNamespaceName(name objectivec.IObject)
	_populateEntries()
	_register() bool
	_registerForUpdates()
	_setDeploymentID()
	DeploymentID() string
	SetDeploymentID(value string)
	Entries() foundation.INSDictionary
	SetEntries(value foundation.INSDictionary)
	Identifier() string
	NamespaceName() string
	TrialClient() unsafe.Pointer
	InitWithIdentifier(identifier objectivec.IObject) MLModelCollection
}

// Init initializes the instance.
func (m MLModelCollection) Init() MLModelCollection {
	rv := objc.Send[MLModelCollection](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelCollection) Autorelease() MLModelCollection {
	rv := objc.Send[MLModelCollection](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelCollection creates a new MLModelCollection instance.
func NewMLModelCollection() MLModelCollection {
	class := getMLModelCollectionClass()
	rv := objc.Send[MLModelCollection](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelCollection/initWithIdentifier:
func NewModelCollectionWithIdentifier(identifier objectivec.IObject) MLModelCollection {
	instance := getMLModelCollectionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIdentifier:"), identifier)
	return MLModelCollectionFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelCollection/_downloadOptions
func (m MLModelCollection) _downloadOptions() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_downloadOptions"))
	return objectivec.Object{ID: rv}
}

// DownloadOptions is an exported wrapper for the private method _downloadOptions.
func (m MLModelCollection) DownloadOptions() objectivec.IObject {
	return m._downloadOptions()
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelCollection/_downloadWithProgress:
func (m MLModelCollection) _downloadWithProgress(progress VoidHandler) bool {
_block0, _ := NewVoidBlock(progress)
	rv := objc.Send[bool](m.ID, objc.Sel("_downloadWithProgress:"), _block0)
	return rv
}

// DownloadWithProgress is an exported wrapper for the private method _downloadWithProgress.
func (m MLModelCollection) DownloadWithProgress(progress VoidHandler) bool {
	return m._downloadWithProgress(progress)
}
// See: https://developer.apple.com/documentation/CoreML/MLModelCollection/_endAccess
func (m MLModelCollection) _endAccess() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("_endAccess"))
	return rv
}

// EndAccess is an exported wrapper for the private method _endAccess.
func (m MLModelCollection) EndAccess() bool {
	return m._endAccess()
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelCollection/_handleTrialUpdateForNamespaceName:
func (m MLModelCollection) _handleTrialUpdateForNamespaceName(name objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("_handleTrialUpdateForNamespaceName:"), name)
}

// HandleTrialUpdateForNamespaceName is an exported wrapper for the private method _handleTrialUpdateForNamespaceName.
func (m MLModelCollection) HandleTrialUpdateForNamespaceName(name objectivec.IObject) {
	m._handleTrialUpdateForNamespaceName(name)
}
// See: https://developer.apple.com/documentation/CoreML/MLModelCollection/_populateEntries
func (m MLModelCollection) _populateEntries() {
	objc.Send[objc.ID](m.ID, objc.Sel("_populateEntries"))
}

// PopulateEntries is an exported wrapper for the private method _populateEntries.
func (m MLModelCollection) PopulateEntries() {
	m._populateEntries()
}
// See: https://developer.apple.com/documentation/CoreML/MLModelCollection/_register
func (m MLModelCollection) _register() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("_register"))
	return rv
}

// Register is an exported wrapper for the private method _register.
func (m MLModelCollection) Register() bool {
	return m._register()
}
// See: https://developer.apple.com/documentation/CoreML/MLModelCollection/_registerForUpdates
func (m MLModelCollection) _registerForUpdates() {
	objc.Send[objc.ID](m.ID, objc.Sel("_registerForUpdates"))
}

// RegisterForUpdates is an exported wrapper for the private method _registerForUpdates.
func (m MLModelCollection) RegisterForUpdates() {
	m._registerForUpdates()
}
// See: https://developer.apple.com/documentation/CoreML/MLModelCollection/_setDeploymentID
func (m MLModelCollection) _setDeploymentID() {
	objc.Send[objc.ID](m.ID, objc.Sel("_setDeploymentID"))
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelCollection/initWithIdentifier:
func (m MLModelCollection) InitWithIdentifier(identifier objectivec.IObject) MLModelCollection {
	rv := objc.Send[MLModelCollection](m.ID, objc.Sel("initWithIdentifier:"), identifier)
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelCollection/_namespaceNameFromCollectionIdentifier:
func (_MLModelCollectionClass MLModelCollectionClass) _namespaceNameFromCollectionIdentifier(identifier objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLModelCollectionClass.class), objc.Sel("_namespaceNameFromCollectionIdentifier:"), identifier)
	return objectivec.Object{ID: rv}
}

// NamespaceNameFromCollectionIdentifier is an exported wrapper for the private method _namespaceNameFromCollectionIdentifier.
func (_MLModelCollectionClass MLModelCollectionClass) NamespaceNameFromCollectionIdentifier(identifier objectivec.IObject) objectivec.IObject {
	return _MLModelCollectionClass._namespaceNameFromCollectionIdentifier(identifier)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelCollection/beginAccessingModelCollectionWithIdentifier:completionHandler:
func (_MLModelCollectionClass MLModelCollectionClass) BeginAccessingModelCollectionWithIdentifierCompletionHandler(identifier objectivec.IObject, handler ErrorHandler) objectivec.IObject {
_block1, _ := NewErrorBlock(handler)
	rv := objc.Send[objc.ID](objc.ID(_MLModelCollectionClass.class), objc.Sel("beginAccessingModelCollectionWithIdentifier:completionHandler:"), identifier, _block1)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelCollection/endAccessingModelCollectionWithIdentifier:completionHandler:
func (_MLModelCollectionClass MLModelCollectionClass) EndAccessingModelCollectionWithIdentifierCompletionHandler(identifier objectivec.IObject, handler ErrorHandler) {
_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](objc.ID(_MLModelCollectionClass.class), objc.Sel("endAccessingModelCollectionWithIdentifier:completionHandler:"), identifier, _block1)
}
// See: https://developer.apple.com/documentation/CoreML/MLModelCollection/getTrialClientClass
func (_MLModelCollectionClass MLModelCollectionClass) GetTrialClientClass() objc.Class {
	rv := objc.Send[objc.Class](objc.ID(_MLModelCollectionClass.class), objc.Sel("getTrialClientClass"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLModelCollection/getTrialDownloadOptionsClass
func (_MLModelCollectionClass MLModelCollectionClass) GetTrialDownloadOptionsClass() objc.Class {
	rv := objc.Send[objc.Class](objc.ID(_MLModelCollectionClass.class), objc.Sel("getTrialDownloadOptionsClass"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLModelCollection/getTrialExperimentIdentifiersClass
func (_MLModelCollectionClass MLModelCollectionClass) GetTrialExperimentIdentifiersClass() objc.Class {
	rv := objc.Send[objc.Class](objc.ID(_MLModelCollectionClass.class), objc.Sel("getTrialExperimentIdentifiersClass"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLModelCollection/getTrialFactorClass
func (_MLModelCollectionClass MLModelCollectionClass) GetTrialFactorClass() objc.Class {
	rv := objc.Send[objc.Class](objc.ID(_MLModelCollectionClass.class), objc.Sel("getTrialFactorClass"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLModelCollection/getTrialFactorLevelClass
func (_MLModelCollectionClass MLModelCollectionClass) GetTrialFactorLevelClass() objc.Class {
	rv := objc.Send[objc.Class](objc.ID(_MLModelCollectionClass.class), objc.Sel("getTrialFactorLevelClass"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLModelCollection/getTrialFileClass
func (_MLModelCollectionClass MLModelCollectionClass) GetTrialFileClass() objc.Class {
	rv := objc.Send[objc.Class](objc.ID(_MLModelCollectionClass.class), objc.Sel("getTrialFileClass"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLModelCollection/getTrialLevelClass
func (_MLModelCollectionClass MLModelCollectionClass) GetTrialLevelClass() objc.Class {
	rv := objc.Send[objc.Class](objc.ID(_MLModelCollectionClass.class), objc.Sel("getTrialLevelClass"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelCollection/deploymentID
func (m MLModelCollection) DeploymentID() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("deploymentID"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLModelCollection) SetDeploymentID(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setDeploymentID:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/CoreML/MLModelCollection/entries
func (m MLModelCollection) Entries() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("entries"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLModelCollection) SetEntries(value foundation.INSDictionary) {
	objc.Send[struct{}](m.ID, objc.Sel("setEntries:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLModelCollection/identifier
func (m MLModelCollection) Identifier() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("identifier"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLModelCollection/namespaceName
func (m MLModelCollection) NamespaceName() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("namespaceName"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLModelCollection/trialClient
func (m MLModelCollection) TrialClient() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m.ID, objc.Sel("trialClient"))
	return rv
}

// _downloadWithProgressSync is a synchronous wrapper around [MLModelCollection._downloadWithProgress].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLModelCollection) _downloadWithProgressSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	m._downloadWithProgress(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// BeginAccessingModelCollectionWithIdentifier is a synchronous wrapper around [MLModelCollection.BeginAccessingModelCollectionWithIdentifierCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (mc MLModelCollectionClass) BeginAccessingModelCollectionWithIdentifier(ctx context.Context, identifier objectivec.IObject) error {
	done := make(chan error, 1)
	mc.BeginAccessingModelCollectionWithIdentifierCompletionHandler(identifier, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EndAccessingModelCollectionWithIdentifier is a synchronous wrapper around [MLModelCollection.EndAccessingModelCollectionWithIdentifierCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (mc MLModelCollectionClass) EndAccessingModelCollectionWithIdentifier(ctx context.Context, identifier objectivec.IObject) error {
	done := make(chan error, 1)
	mc.EndAccessingModelCollectionWithIdentifierCompletionHandler(identifier, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

