// Code generated from Apple documentation for speechobjects. DO NOT EDIT.

package speechobjects

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SOSpeechInstallationManager] class.
var (
	_SOSpeechInstallationManagerClass     SOSpeechInstallationManagerClass
	_SOSpeechInstallationManagerClassOnce sync.Once
)

func getSOSpeechInstallationManagerClass() SOSpeechInstallationManagerClass {
	_SOSpeechInstallationManagerClassOnce.Do(func() {
		_SOSpeechInstallationManagerClass = SOSpeechInstallationManagerClass{class: objc.GetClass("SOSpeechInstallationManager")}
	})
	return _SOSpeechInstallationManagerClass
}

// GetSOSpeechInstallationManagerClass returns the class object for SOSpeechInstallationManager.
func GetSOSpeechInstallationManagerClass() SOSpeechInstallationManagerClass {
	return getSOSpeechInstallationManagerClass()
}

type SOSpeechInstallationManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SOSpeechInstallationManagerClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SOSpeechInstallationManagerClass) Alloc() SOSpeechInstallationManager {
	rv := objc.Send[SOSpeechInstallationManager](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SOSpeechInstallationManager._clientHasRightsToCustomVoices]
//   - [SOSpeechInstallationManager._createLocalPort]
//   - [SOSpeechInstallationManager._createServerPortIfNeeded]
//   - [SOSpeechInstallationManager._overriddenVoiceIdentifierDictionary]
//   - [SOSpeechInstallationManager._sendMessageWithDataWithReply]
//   - [SOSpeechInstallationManager._serverIsRunning]
//   - [SOSpeechInstallationManager._startDownloadingHighestQualityIfNecessaryForVoiceIdentifierRequireACPowerInitiator]
//   - [SOSpeechInstallationManager._voiceIdentifierForRootVoiceIdentifierStartDownloadingRequireACPowerInitiatorHighestQuality]
//   - [SOSpeechInstallationManager.ActiveInstallations]
//   - [SOSpeechInstallationManager.ActiveVoiceInstallations]
//   - [SOSpeechInstallationManager.AreBackgroundDownloadsAllowed]
//   - [SOSpeechInstallationManager.BundleForRemovableLanguagePassingBackTagName]
//   - [SOSpeechInstallationManager.BundleForRemovableVoicePassingBackTagName]
//   - [SOSpeechInstallationManager.CancelAllInstallations]
//   - [SOSpeechInstallationManager.CancelPurgingByInitiator]
//   - [SOSpeechInstallationManager.CancelPurgingOfSpeechBundleIdentifiersInitiator]
//   - [SOSpeechInstallationManager.Delegate]
//   - [SOSpeechInstallationManager.SetDelegate]
//   - [SOSpeechInstallationManager.DoesLanguageDataNeedToBeDownloadedForLocaleIdentifier]
//   - [SOSpeechInstallationManager.DownloadableSRLanguageItems]
//   - [SOSpeechInstallationManager.DownloadableSpeechItemUpgradeForExistingSpeechItemFromDownloadableItems]
//   - [SOSpeechInstallationManager.DownloadableVoiceUpgradeForVoiceFromDownloadableVoices]
//   - [SOSpeechInstallationManager.DownloadableVoiceUpgradeForVoiceIdentifierDesirabilityFromDownloadableVoices]
//   - [SOSpeechInstallationManager.DownloadableVoices]
//   - [SOSpeechInstallationManager.HighestQualityVoiceIdentifierForRootVoiceIdentifierStartDownloading]
//   - [SOSpeechInstallationManager.HighestQualityVoiceIdentifierForRootVoiceIdentifierStartDownloadingRequireACPowerInitiator]
//   - [SOSpeechInstallationManager.InstallationLogEntryForTag]
//   - [SOSpeechInstallationManager.InstallationLogEntryForTagPreferenceDomain]
//   - [SOSpeechInstallationManager.IsAutoDownloadProhibitedForDownloadableVoiceObject]
//   - [SOSpeechInstallationManager.IsRunningOnACPower]
//   - [SOSpeechInstallationManager.LowestQualityVoiceIdentifierForRootVoiceIdentifierStartDownloadingRequireACPowerInitiator]
//   - [SOSpeechInstallationManager.MarkDownloadableVoiceObjectWithAutoDownloadProhibitFlag]
//   - [SOSpeechInstallationManager.OnDiskVersionForBundlePath]
//   - [SOSpeechInstallationManager.PercentageOfBatteryCharge]
//   - [SOSpeechInstallationManager.ResetAllVoiceDownloadAttemptsForInitiator]
//   - [SOSpeechInstallationManager.ShowProgressWindow]
//   - [SOSpeechInstallationManager.ShowSRLanguagesSelectionSheetForWindowNetworkSupportedLocaleIdentifiersRequiredLocaleIdentifierSupportDownloadsShowOnlyNetworkSupportedItems]
//   - [SOSpeechInstallationManager.ShowVoiceSelectionSheetForWindowShowIndividualVoiceQualitiesVoiceIdentifiersNotToBeRemoved]
//   - [SOSpeechInstallationManager.SrLanguagesSelectionWindowController]
//   - [SOSpeechInstallationManager.StartInstallingDownloadableSpeechItemsUserInteractionMode]
//   - [SOSpeechInstallationManager.StartInstallingDownloadableSpeechItemsUserInteractionModeInitiator]
//   - [SOSpeechInstallationManager.StartInstallingDownloadableVoicesUseRootNamesUserInteractionMode]
//   - [SOSpeechInstallationManager.StartInstallingDownloadableVoicesUseRootNamesUserInteractionModeInitiator]
//   - [SOSpeechInstallationManager.StartRemovingDownloadedSpeechBundleIdentifiersInitiatorImmediately]
//   - [SOSpeechInstallationManager.StartRemovingDownloadedSpeechItems]
//   - [SOSpeechInstallationManager.StartRemovingDownloadedSpeechItemsInitiatorImmediately]
//   - [SOSpeechInstallationManager.VoiceSelectionWindowController]
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager
type SOSpeechInstallationManager struct {
	objectivec.Object
}

// SOSpeechInstallationManagerFromID constructs a [SOSpeechInstallationManager] from an objc.ID.
func SOSpeechInstallationManagerFromID(id objc.ID) SOSpeechInstallationManager {
	return SOSpeechInstallationManager{objectivec.Object{ID: id}}
}

// Ensure SOSpeechInstallationManager implements ISOSpeechInstallationManager.
var _ ISOSpeechInstallationManager = SOSpeechInstallationManager{}

// An interface definition for the [SOSpeechInstallationManager] class.
//
// # Methods
//
//   - [ISOSpeechInstallationManager._clientHasRightsToCustomVoices]
//   - [ISOSpeechInstallationManager._createLocalPort]
//   - [ISOSpeechInstallationManager._createServerPortIfNeeded]
//   - [ISOSpeechInstallationManager._overriddenVoiceIdentifierDictionary]
//   - [ISOSpeechInstallationManager._sendMessageWithDataWithReply]
//   - [ISOSpeechInstallationManager._serverIsRunning]
//   - [ISOSpeechInstallationManager._startDownloadingHighestQualityIfNecessaryForVoiceIdentifierRequireACPowerInitiator]
//   - [ISOSpeechInstallationManager._voiceIdentifierForRootVoiceIdentifierStartDownloadingRequireACPowerInitiatorHighestQuality]
//   - [ISOSpeechInstallationManager.ActiveInstallations]
//   - [ISOSpeechInstallationManager.ActiveVoiceInstallations]
//   - [ISOSpeechInstallationManager.AreBackgroundDownloadsAllowed]
//   - [ISOSpeechInstallationManager.BundleForRemovableLanguagePassingBackTagName]
//   - [ISOSpeechInstallationManager.BundleForRemovableVoicePassingBackTagName]
//   - [ISOSpeechInstallationManager.CancelAllInstallations]
//   - [ISOSpeechInstallationManager.CancelPurgingByInitiator]
//   - [ISOSpeechInstallationManager.CancelPurgingOfSpeechBundleIdentifiersInitiator]
//   - [ISOSpeechInstallationManager.Delegate]
//   - [ISOSpeechInstallationManager.SetDelegate]
//   - [ISOSpeechInstallationManager.DoesLanguageDataNeedToBeDownloadedForLocaleIdentifier]
//   - [ISOSpeechInstallationManager.DownloadableSRLanguageItems]
//   - [ISOSpeechInstallationManager.DownloadableSpeechItemUpgradeForExistingSpeechItemFromDownloadableItems]
//   - [ISOSpeechInstallationManager.DownloadableVoiceUpgradeForVoiceFromDownloadableVoices]
//   - [ISOSpeechInstallationManager.DownloadableVoiceUpgradeForVoiceIdentifierDesirabilityFromDownloadableVoices]
//   - [ISOSpeechInstallationManager.DownloadableVoices]
//   - [ISOSpeechInstallationManager.HighestQualityVoiceIdentifierForRootVoiceIdentifierStartDownloading]
//   - [ISOSpeechInstallationManager.HighestQualityVoiceIdentifierForRootVoiceIdentifierStartDownloadingRequireACPowerInitiator]
//   - [ISOSpeechInstallationManager.InstallationLogEntryForTag]
//   - [ISOSpeechInstallationManager.InstallationLogEntryForTagPreferenceDomain]
//   - [ISOSpeechInstallationManager.IsAutoDownloadProhibitedForDownloadableVoiceObject]
//   - [ISOSpeechInstallationManager.IsRunningOnACPower]
//   - [ISOSpeechInstallationManager.LowestQualityVoiceIdentifierForRootVoiceIdentifierStartDownloadingRequireACPowerInitiator]
//   - [ISOSpeechInstallationManager.MarkDownloadableVoiceObjectWithAutoDownloadProhibitFlag]
//   - [ISOSpeechInstallationManager.OnDiskVersionForBundlePath]
//   - [ISOSpeechInstallationManager.PercentageOfBatteryCharge]
//   - [ISOSpeechInstallationManager.ResetAllVoiceDownloadAttemptsForInitiator]
//   - [ISOSpeechInstallationManager.ShowProgressWindow]
//   - [ISOSpeechInstallationManager.ShowSRLanguagesSelectionSheetForWindowNetworkSupportedLocaleIdentifiersRequiredLocaleIdentifierSupportDownloadsShowOnlyNetworkSupportedItems]
//   - [ISOSpeechInstallationManager.ShowVoiceSelectionSheetForWindowShowIndividualVoiceQualitiesVoiceIdentifiersNotToBeRemoved]
//   - [ISOSpeechInstallationManager.SrLanguagesSelectionWindowController]
//   - [ISOSpeechInstallationManager.StartInstallingDownloadableSpeechItemsUserInteractionMode]
//   - [ISOSpeechInstallationManager.StartInstallingDownloadableSpeechItemsUserInteractionModeInitiator]
//   - [ISOSpeechInstallationManager.StartInstallingDownloadableVoicesUseRootNamesUserInteractionMode]
//   - [ISOSpeechInstallationManager.StartInstallingDownloadableVoicesUseRootNamesUserInteractionModeInitiator]
//   - [ISOSpeechInstallationManager.StartRemovingDownloadedSpeechBundleIdentifiersInitiatorImmediately]
//   - [ISOSpeechInstallationManager.StartRemovingDownloadedSpeechItems]
//   - [ISOSpeechInstallationManager.StartRemovingDownloadedSpeechItemsInitiatorImmediately]
//   - [ISOSpeechInstallationManager.VoiceSelectionWindowController]
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager
type ISOSpeechInstallationManager interface {
	objectivec.IObject

	// Topic: Methods

	_clientHasRightsToCustomVoices() bool
	_createLocalPort() corefoundation.CFMessagePortRef
	_createServerPortIfNeeded() bool
	_overriddenVoiceIdentifierDictionary() objectivec.IObject
	_sendMessageWithDataWithReply(message int, data unsafe.Pointer, reply bool) objectivec.IObject
	_serverIsRunning() bool
	_startDownloadingHighestQualityIfNecessaryForVoiceIdentifierRequireACPowerInitiator(identifier objectivec.IObject, aCPower bool, initiator objectivec.IObject) bool
	_voiceIdentifierForRootVoiceIdentifierStartDownloadingRequireACPowerInitiatorHighestQuality(identifier objectivec.IObject, downloading bool, aCPower bool, initiator objectivec.IObject, quality bool) objectivec.IObject
	ActiveInstallations() objectivec.IObject
	ActiveVoiceInstallations() objectivec.IObject
	AreBackgroundDownloadsAllowed() bool
	BundleForRemovableLanguagePassingBackTagName(language objectivec.IObject, name []objectivec.IObject) objectivec.IObject
	BundleForRemovableVoicePassingBackTagName(voice objectivec.IObject, name []objectivec.IObject) objectivec.IObject
	CancelAllInstallations()
	CancelPurgingByInitiator(initiator objectivec.IObject)
	CancelPurgingOfSpeechBundleIdentifiersInitiator(identifiers objectivec.IObject, initiator objectivec.IObject)
	Delegate() objectivec.IObject
	SetDelegate(value objectivec.IObject)
	DoesLanguageDataNeedToBeDownloadedForLocaleIdentifier(identifier objectivec.IObject) bool
	DownloadableSRLanguageItems() objectivec.IObject
	DownloadableSpeechItemUpgradeForExistingSpeechItemFromDownloadableItems(item objectivec.IObject, items objectivec.IObject) objectivec.IObject
	DownloadableVoiceUpgradeForVoiceFromDownloadableVoices(voice objectivec.IObject, voices objectivec.IObject) objectivec.IObject
	DownloadableVoiceUpgradeForVoiceIdentifierDesirabilityFromDownloadableVoices(identifier objectivec.IObject, desirability int64, voices objectivec.IObject) objectivec.IObject
	DownloadableVoices() objectivec.IObject
	HighestQualityVoiceIdentifierForRootVoiceIdentifierStartDownloading(identifier objectivec.IObject, downloading bool) objectivec.IObject
	HighestQualityVoiceIdentifierForRootVoiceIdentifierStartDownloadingRequireACPowerInitiator(identifier objectivec.IObject, downloading bool, aCPower bool, initiator objectivec.IObject) objectivec.IObject
	InstallationLogEntryForTag(tag objectivec.IObject) objectivec.IObject
	InstallationLogEntryForTagPreferenceDomain(tag objectivec.IObject, domain objectivec.IObject) objectivec.IObject
	IsAutoDownloadProhibitedForDownloadableVoiceObject(object objectivec.IObject) bool
	IsRunningOnACPower() bool
	LowestQualityVoiceIdentifierForRootVoiceIdentifierStartDownloadingRequireACPowerInitiator(identifier objectivec.IObject, downloading bool, aCPower bool, initiator objectivec.IObject) objectivec.IObject
	MarkDownloadableVoiceObjectWithAutoDownloadProhibitFlag(object objectivec.IObject, flag bool)
	OnDiskVersionForBundlePath(path objectivec.IObject) objectivec.IObject
	PercentageOfBatteryCharge() float64
	ResetAllVoiceDownloadAttemptsForInitiator(initiator objectivec.IObject)
	ShowProgressWindow()
	ShowSRLanguagesSelectionSheetForWindowNetworkSupportedLocaleIdentifiersRequiredLocaleIdentifierSupportDownloadsShowOnlyNetworkSupportedItems(window objectivec.IObject, identifiers objectivec.IObject, identifier objectivec.IObject, downloads bool, items bool)
	ShowVoiceSelectionSheetForWindowShowIndividualVoiceQualitiesVoiceIdentifiersNotToBeRemoved(window objectivec.IObject, qualities bool, removed objectivec.IObject)
	SrLanguagesSelectionWindowController() ISOCustomizeSRLanguagesWindowController
	StartInstallingDownloadableSpeechItemsUserInteractionMode(items objectivec.IObject, mode uint32) bool
	StartInstallingDownloadableSpeechItemsUserInteractionModeInitiator(items objectivec.IObject, mode uint32, initiator objectivec.IObject) bool
	StartInstallingDownloadableVoicesUseRootNamesUserInteractionMode(voices objectivec.IObject, names bool, mode uint32) bool
	StartInstallingDownloadableVoicesUseRootNamesUserInteractionModeInitiator(voices objectivec.IObject, names bool, mode uint32, initiator objectivec.IObject) bool
	StartRemovingDownloadedSpeechBundleIdentifiersInitiatorImmediately(identifiers objectivec.IObject, initiator objectivec.IObject, immediately bool)
	StartRemovingDownloadedSpeechItems(items objectivec.IObject)
	StartRemovingDownloadedSpeechItemsInitiatorImmediately(items objectivec.IObject, initiator objectivec.IObject, immediately bool)
	VoiceSelectionWindowController() ICustomizeVoicesWindowController
}

// Init initializes the instance.
func (s SOSpeechInstallationManager) Init() SOSpeechInstallationManager {
	rv := objc.Send[SOSpeechInstallationManager](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SOSpeechInstallationManager) Autorelease() SOSpeechInstallationManager {
	rv := objc.Send[SOSpeechInstallationManager](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSOSpeechInstallationManager creates a new SOSpeechInstallationManager instance.
func NewSOSpeechInstallationManager() SOSpeechInstallationManager {
	class := getSOSpeechInstallationManagerClass()
	rv := objc.Send[SOSpeechInstallationManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/_clientHasRightsToCustomVoices
func (s SOSpeechInstallationManager) _clientHasRightsToCustomVoices() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("_clientHasRightsToCustomVoices"))
	return rv
}

// ClientHasRightsToCustomVoices is an exported wrapper for the private method _clientHasRightsToCustomVoices.
func (s SOSpeechInstallationManager) ClientHasRightsToCustomVoices() bool {
	return s._clientHasRightsToCustomVoices()
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/_createLocalPort
func (s SOSpeechInstallationManager) _createLocalPort() corefoundation.CFMessagePortRef {
	rv := objc.Send[corefoundation.CFMessagePortRef](s.ID, objc.Sel("_createLocalPort"))
	return corefoundation.CFMessagePortRef(rv)
}

// CreateLocalPort is an exported wrapper for the private method _createLocalPort.
func (s SOSpeechInstallationManager) CreateLocalPort() corefoundation.CFMessagePortRef {
	return s._createLocalPort()
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/_createServerPortIfNeeded
func (s SOSpeechInstallationManager) _createServerPortIfNeeded() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("_createServerPortIfNeeded"))
	return rv
}

// CreateServerPortIfNeeded is an exported wrapper for the private method _createServerPortIfNeeded.
func (s SOSpeechInstallationManager) CreateServerPortIfNeeded() bool {
	return s._createServerPortIfNeeded()
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/_overriddenVoiceIdentifierDictionary
func (s SOSpeechInstallationManager) _overriddenVoiceIdentifierDictionary() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_overriddenVoiceIdentifierDictionary"))
	return objectivec.Object{ID: rv}
}

// OverriddenVoiceIdentifierDictionary is an exported wrapper for the private method _overriddenVoiceIdentifierDictionary.
func (s SOSpeechInstallationManager) OverriddenVoiceIdentifierDictionary() objectivec.IObject {
	return s._overriddenVoiceIdentifierDictionary()
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/_sendMessage:withData:withReply:
func (s SOSpeechInstallationManager) _sendMessageWithDataWithReply(message int, data unsafe.Pointer, reply bool) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_sendMessage:withData:withReply:"), message, data, reply)
	return objectivec.Object{ID: rv}
}

// SendMessageWithDataWithReply is an exported wrapper for the private method _sendMessageWithDataWithReply.
func (s SOSpeechInstallationManager) SendMessageWithDataWithReply(message int, data unsafe.Pointer, reply bool) objectivec.IObject {
	return s._sendMessageWithDataWithReply(message, data, reply)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/_serverIsRunning
func (s SOSpeechInstallationManager) _serverIsRunning() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("_serverIsRunning"))
	return rv
}

// ServerIsRunning is an exported wrapper for the private method _serverIsRunning.
func (s SOSpeechInstallationManager) ServerIsRunning() bool {
	return s._serverIsRunning()
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/_startDownloadingHighestQualityIfNecessaryForVoiceIdentifier:requireACPower:initiator:
func (s SOSpeechInstallationManager) _startDownloadingHighestQualityIfNecessaryForVoiceIdentifierRequireACPowerInitiator(identifier objectivec.IObject, aCPower bool, initiator objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("_startDownloadingHighestQualityIfNecessaryForVoiceIdentifier:requireACPower:initiator:"), identifier, aCPower, initiator)
	return rv
}

// StartDownloadingHighestQualityIfNecessaryForVoiceIdentifierRequireACPowerInitiator is an exported wrapper for the private method _startDownloadingHighestQualityIfNecessaryForVoiceIdentifierRequireACPowerInitiator.
func (s SOSpeechInstallationManager) StartDownloadingHighestQualityIfNecessaryForVoiceIdentifierRequireACPowerInitiator(identifier objectivec.IObject, aCPower bool, initiator objectivec.IObject) bool {
	return s._startDownloadingHighestQualityIfNecessaryForVoiceIdentifierRequireACPowerInitiator(identifier, aCPower, initiator)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/_voiceIdentifierForRootVoiceIdentifier:startDownloading:requireACPower:initiator:highestQuality:
func (s SOSpeechInstallationManager) _voiceIdentifierForRootVoiceIdentifierStartDownloadingRequireACPowerInitiatorHighestQuality(identifier objectivec.IObject, downloading bool, aCPower bool, initiator objectivec.IObject, quality bool) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_voiceIdentifierForRootVoiceIdentifier:startDownloading:requireACPower:initiator:highestQuality:"), identifier, downloading, aCPower, initiator, quality)
	return objectivec.Object{ID: rv}
}

// VoiceIdentifierForRootVoiceIdentifierStartDownloadingRequireACPowerInitiatorHighestQuality is an exported wrapper for the private method _voiceIdentifierForRootVoiceIdentifierStartDownloadingRequireACPowerInitiatorHighestQuality.
func (s SOSpeechInstallationManager) VoiceIdentifierForRootVoiceIdentifierStartDownloadingRequireACPowerInitiatorHighestQuality(identifier objectivec.IObject, downloading bool, aCPower bool, initiator objectivec.IObject, quality bool) objectivec.IObject {
	return s._voiceIdentifierForRootVoiceIdentifierStartDownloadingRequireACPowerInitiatorHighestQuality(identifier, downloading, aCPower, initiator, quality)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/activeInstallations
func (s SOSpeechInstallationManager) ActiveInstallations() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("activeInstallations"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/activeVoiceInstallations
func (s SOSpeechInstallationManager) ActiveVoiceInstallations() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("activeVoiceInstallations"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/areBackgroundDownloadsAllowed
func (s SOSpeechInstallationManager) AreBackgroundDownloadsAllowed() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("areBackgroundDownloadsAllowed"))
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/bundleForRemovableLanguage:passingBackTagName:
func (s SOSpeechInstallationManager) BundleForRemovableLanguagePassingBackTagName(language objectivec.IObject, name []objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("bundleForRemovableLanguage:passingBackTagName:"), language, objectivec.IObjectSliceToNSArray(name))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/bundleForRemovableVoice:passingBackTagName:
func (s SOSpeechInstallationManager) BundleForRemovableVoicePassingBackTagName(voice objectivec.IObject, name []objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("bundleForRemovableVoice:passingBackTagName:"), voice, objectivec.IObjectSliceToNSArray(name))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/cancelAllInstallations
func (s SOSpeechInstallationManager) CancelAllInstallations() {
	objc.Send[objc.ID](s.ID, objc.Sel("cancelAllInstallations"))
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/cancelPurgingByInitiator:
func (s SOSpeechInstallationManager) CancelPurgingByInitiator(initiator objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("cancelPurgingByInitiator:"), initiator)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/cancelPurgingOfSpeechBundleIdentifiers:initiator:
func (s SOSpeechInstallationManager) CancelPurgingOfSpeechBundleIdentifiersInitiator(identifiers objectivec.IObject, initiator objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("cancelPurgingOfSpeechBundleIdentifiers:initiator:"), identifiers, initiator)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/doesLanguageDataNeedToBeDownloadedForLocaleIdentifier:
func (s SOSpeechInstallationManager) DoesLanguageDataNeedToBeDownloadedForLocaleIdentifier(identifier objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("doesLanguageDataNeedToBeDownloadedForLocaleIdentifier:"), identifier)
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/downloadableSRLanguageItems
func (s SOSpeechInstallationManager) DownloadableSRLanguageItems() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("downloadableSRLanguageItems"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/downloadableSpeechItemUpgradeForExistingSpeechItem:fromDownloadableItems:
func (s SOSpeechInstallationManager) DownloadableSpeechItemUpgradeForExistingSpeechItemFromDownloadableItems(item objectivec.IObject, items objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("downloadableSpeechItemUpgradeForExistingSpeechItem:fromDownloadableItems:"), item, items)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/downloadableVoiceUpgradeForVoice:fromDownloadableVoices:
func (s SOSpeechInstallationManager) DownloadableVoiceUpgradeForVoiceFromDownloadableVoices(voice objectivec.IObject, voices objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("downloadableVoiceUpgradeForVoice:fromDownloadableVoices:"), voice, voices)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/downloadableVoiceUpgradeForVoiceIdentifier:desirability:fromDownloadableVoices:
func (s SOSpeechInstallationManager) DownloadableVoiceUpgradeForVoiceIdentifierDesirabilityFromDownloadableVoices(identifier objectivec.IObject, desirability int64, voices objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("downloadableVoiceUpgradeForVoiceIdentifier:desirability:fromDownloadableVoices:"), identifier, desirability, voices)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/downloadableVoices
func (s SOSpeechInstallationManager) DownloadableVoices() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("downloadableVoices"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/highestQualityVoiceIdentifierForRootVoiceIdentifier:startDownloading:
func (s SOSpeechInstallationManager) HighestQualityVoiceIdentifierForRootVoiceIdentifierStartDownloading(identifier objectivec.IObject, downloading bool) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("highestQualityVoiceIdentifierForRootVoiceIdentifier:startDownloading:"), identifier, downloading)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/highestQualityVoiceIdentifierForRootVoiceIdentifier:startDownloading:requireACPower:initiator:
func (s SOSpeechInstallationManager) HighestQualityVoiceIdentifierForRootVoiceIdentifierStartDownloadingRequireACPowerInitiator(identifier objectivec.IObject, downloading bool, aCPower bool, initiator objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("highestQualityVoiceIdentifierForRootVoiceIdentifier:startDownloading:requireACPower:initiator:"), identifier, downloading, aCPower, initiator)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/installationLogEntryForTag:
func (s SOSpeechInstallationManager) InstallationLogEntryForTag(tag objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("installationLogEntryForTag:"), tag)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/installationLogEntryForTag:preferenceDomain:
func (s SOSpeechInstallationManager) InstallationLogEntryForTagPreferenceDomain(tag objectivec.IObject, domain objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("installationLogEntryForTag:preferenceDomain:"), tag, domain)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/isAutoDownloadProhibitedForDownloadableVoiceObject:
func (s SOSpeechInstallationManager) IsAutoDownloadProhibitedForDownloadableVoiceObject(object objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isAutoDownloadProhibitedForDownloadableVoiceObject:"), object)
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/isRunningOnACPower
func (s SOSpeechInstallationManager) IsRunningOnACPower() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isRunningOnACPower"))
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/lowestQualityVoiceIdentifierForRootVoiceIdentifier:startDownloading:requireACPower:initiator:
func (s SOSpeechInstallationManager) LowestQualityVoiceIdentifierForRootVoiceIdentifierStartDownloadingRequireACPowerInitiator(identifier objectivec.IObject, downloading bool, aCPower bool, initiator objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("lowestQualityVoiceIdentifierForRootVoiceIdentifier:startDownloading:requireACPower:initiator:"), identifier, downloading, aCPower, initiator)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/markDownloadableVoiceObject:withAutoDownloadProhibitFlag:
func (s SOSpeechInstallationManager) MarkDownloadableVoiceObjectWithAutoDownloadProhibitFlag(object objectivec.IObject, flag bool) {
	objc.Send[objc.ID](s.ID, objc.Sel("markDownloadableVoiceObject:withAutoDownloadProhibitFlag:"), object, flag)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/onDiskVersionForBundlePath:
func (s SOSpeechInstallationManager) OnDiskVersionForBundlePath(path objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("onDiskVersionForBundlePath:"), path)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/percentageOfBatteryCharge
func (s SOSpeechInstallationManager) PercentageOfBatteryCharge() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("percentageOfBatteryCharge"))
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/resetAllVoiceDownloadAttemptsForInitiator:
func (s SOSpeechInstallationManager) ResetAllVoiceDownloadAttemptsForInitiator(initiator objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("resetAllVoiceDownloadAttemptsForInitiator:"), initiator)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/showProgressWindow
func (s SOSpeechInstallationManager) ShowProgressWindow() {
	objc.Send[objc.ID](s.ID, objc.Sel("showProgressWindow"))
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/showSRLanguagesSelectionSheetForWindow:networkSupportedLocaleIdentifiers:requiredLocaleIdentifier:supportDownloads:showOnlyNetworkSupportedItems:
func (s SOSpeechInstallationManager) ShowSRLanguagesSelectionSheetForWindowNetworkSupportedLocaleIdentifiersRequiredLocaleIdentifierSupportDownloadsShowOnlyNetworkSupportedItems(window objectivec.IObject, identifiers objectivec.IObject, identifier objectivec.IObject, downloads bool, items bool) {
	objc.Send[objc.ID](s.ID, objc.Sel("showSRLanguagesSelectionSheetForWindow:networkSupportedLocaleIdentifiers:requiredLocaleIdentifier:supportDownloads:showOnlyNetworkSupportedItems:"), window, identifiers, identifier, downloads, items)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/showVoiceSelectionSheetForWindow:showIndividualVoiceQualities:voiceIdentifiersNotToBeRemoved:
func (s SOSpeechInstallationManager) ShowVoiceSelectionSheetForWindowShowIndividualVoiceQualitiesVoiceIdentifiersNotToBeRemoved(window objectivec.IObject, qualities bool, removed objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("showVoiceSelectionSheetForWindow:showIndividualVoiceQualities:voiceIdentifiersNotToBeRemoved:"), window, qualities, removed)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/startInstallingDownloadableSpeechItems:userInteractionMode:
func (s SOSpeechInstallationManager) StartInstallingDownloadableSpeechItemsUserInteractionMode(items objectivec.IObject, mode uint32) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("startInstallingDownloadableSpeechItems:userInteractionMode:"), items, mode)
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/startInstallingDownloadableSpeechItems:userInteractionMode:initiator:
func (s SOSpeechInstallationManager) StartInstallingDownloadableSpeechItemsUserInteractionModeInitiator(items objectivec.IObject, mode uint32, initiator objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("startInstallingDownloadableSpeechItems:userInteractionMode:initiator:"), items, mode, initiator)
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/startInstallingDownloadableVoices:useRootNames:userInteractionMode:
func (s SOSpeechInstallationManager) StartInstallingDownloadableVoicesUseRootNamesUserInteractionMode(voices objectivec.IObject, names bool, mode uint32) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("startInstallingDownloadableVoices:useRootNames:userInteractionMode:"), voices, names, mode)
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/startInstallingDownloadableVoices:useRootNames:userInteractionMode:initiator:
func (s SOSpeechInstallationManager) StartInstallingDownloadableVoicesUseRootNamesUserInteractionModeInitiator(voices objectivec.IObject, names bool, mode uint32, initiator objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("startInstallingDownloadableVoices:useRootNames:userInteractionMode:initiator:"), voices, names, mode, initiator)
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/startRemovingDownloadedSpeechBundleIdentifiers:initiator:immediately:
func (s SOSpeechInstallationManager) StartRemovingDownloadedSpeechBundleIdentifiersInitiatorImmediately(identifiers objectivec.IObject, initiator objectivec.IObject, immediately bool) {
	objc.Send[objc.ID](s.ID, objc.Sel("startRemovingDownloadedSpeechBundleIdentifiers:initiator:immediately:"), identifiers, initiator, immediately)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/startRemovingDownloadedSpeechItems:
func (s SOSpeechInstallationManager) StartRemovingDownloadedSpeechItems(items objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("startRemovingDownloadedSpeechItems:"), items)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/startRemovingDownloadedSpeechItems:initiator:immediately:
func (s SOSpeechInstallationManager) StartRemovingDownloadedSpeechItemsInitiatorImmediately(items objectivec.IObject, initiator objectivec.IObject, immediately bool) {
	objc.Send[objc.ID](s.ID, objc.Sel("startRemovingDownloadedSpeechItems:initiator:immediately:"), items, initiator, immediately)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/sharedManager
func (_SOSpeechInstallationManagerClass SOSpeechInstallationManagerClass) SharedManager() SOSpeechInstallationManager {
	rv := objc.Send[objc.ID](objc.ID(_SOSpeechInstallationManagerClass.class), objc.Sel("sharedManager"))
	return SOSpeechInstallationManagerFromID(rv)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/delegate
func (s SOSpeechInstallationManager) Delegate() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("delegate"))
	return objectivec.Object{ID: rv}
}
func (s SOSpeechInstallationManager) SetDelegate(value objectivec.IObject) {
	objc.Send[struct{}](s.ID, objc.Sel("setDelegate:"), value)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/srLanguagesSelectionWindowController
func (s SOSpeechInstallationManager) SrLanguagesSelectionWindowController() ISOCustomizeSRLanguagesWindowController {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("srLanguagesSelectionWindowController"))
	return SOCustomizeSRLanguagesWindowControllerFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechInstallationManager/voiceSelectionWindowController
func (s SOSpeechInstallationManager) VoiceSelectionWindowController() ICustomizeVoicesWindowController {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("voiceSelectionWindowController"))
	return CustomizeVoicesWindowControllerFromID(objc.ID(rv))
}
