// Code generated from Apple documentation. DO NOT EDIT.

package iobluetooth

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

var (
	// IOBluetoothHandsFreeCallDirection is a value that indicates whether a call is incoming or outgoing.
	//
	// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeCallDirection
	IOBluetoothHandsFreeCallDirection string
	// IOBluetoothHandsFreeCallIndex is the index of the call, starting with `1`.
	//
	// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeCallIndex
	IOBluetoothHandsFreeCallIndex string
	// IOBluetoothHandsFreeCallMode is the type of call data.
	//
	// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeCallMode
	IOBluetoothHandsFreeCallMode string
	// IOBluetoothHandsFreeCallMultiparty is a value that indicates whether the call is multiple-party.
	//
	// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeCallMultiparty
	IOBluetoothHandsFreeCallMultiparty string
	// IOBluetoothHandsFreeCallName is the name of the caller.
	//
	// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeCallName
	IOBluetoothHandsFreeCallName string
	// IOBluetoothHandsFreeCallNumber is the caller’s phone number.
	//
	// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeCallNumber
	IOBluetoothHandsFreeCallNumber string
	// IOBluetoothHandsFreeCallStatus is the current state of the call.
	//
	// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeCallStatus
	IOBluetoothHandsFreeCallStatus string
	// IOBluetoothHandsFreeCallType is the format of the caller’s phone number.
	//
	// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeCallType
	IOBluetoothHandsFreeCallType string
	// IOBluetoothHandsFreeIndicatorBattChg is the command string you use to show a battery charge level indicator on a phone.
	//
	// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeIndicatorBattChg
	IOBluetoothHandsFreeIndicatorBattChg string
	// IOBluetoothHandsFreeIndicatorCall is the command string you use to show an active call indicator on a phone.
	//
	// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeIndicatorCall
	IOBluetoothHandsFreeIndicatorCall string
	// IOBluetoothHandsFreeIndicatorCallHeld is the command string you use to show a call hold status indicator on a phone.
	//
	// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeIndicatorCallHeld
	IOBluetoothHandsFreeIndicatorCallHeld string
	// IOBluetoothHandsFreeIndicatorCallSetup is the command string you use to show a call setup status indicator on a phone.
	//
	// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeIndicatorCallSetup
	IOBluetoothHandsFreeIndicatorCallSetup string
	// IOBluetoothHandsFreeIndicatorRoam is the command string you use to show a roaming status indicator on a phone.
	//
	// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeIndicatorRoam
	IOBluetoothHandsFreeIndicatorRoam string
	// IOBluetoothHandsFreeIndicatorService is the command string you use to show a carrier network connection indicator on a phone.
	//
	// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeIndicatorService
	IOBluetoothHandsFreeIndicatorService string
	// IOBluetoothHandsFreeIndicatorSignal is the command string you use to show a signal strength indicator on a phone.
	//
	// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeIndicatorSignal
	IOBluetoothHandsFreeIndicatorSignal string
	// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHostControllerPoweredOffNotification
	IOBluetoothHostControllerPoweredOffNotification string
	// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHostControllerPoweredOnNotification
	IOBluetoothHostControllerPoweredOnNotification string
	// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannelPublishedNotification
	IOBluetoothL2CAPChannelPublishedNotification string
	// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannelTerminatedNotification
	IOBluetoothL2CAPChannelTerminatedNotification string
	// IOBluetoothPDUEncoding is the encoding for the text message.
	//
	// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothPDUEncoding
	IOBluetoothPDUEncoding string
	// IOBluetoothPDUOriginatingAddress is the phone number for the sender of the text message.
	//
	// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothPDUOriginatingAddress
	IOBluetoothPDUOriginatingAddress string
	// IOBluetoothPDUOriginatingAddressType is the format of the phone number for the sender of the text message.
	//
	// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothPDUOriginatingAddressType
	IOBluetoothPDUOriginatingAddressType string
	// IOBluetoothPDUProtocolID is the protocol of the text message content.
	//
	// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothPDUProtocolID
	IOBluetoothPDUProtocolID string
	// IOBluetoothPDUServicCenterAddress is the phone number for the service center that stored and then delivered the message.
	//
	// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothPDUServicCenterAddress
	IOBluetoothPDUServicCenterAddress string
	// IOBluetoothPDUServiceCenterAddressType is the format of the phone number for the service center.
	//
	// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothPDUServiceCenterAddressType
	IOBluetoothPDUServiceCenterAddressType string
	// IOBluetoothPDUTimestamp is the time the text message was sent.
	//
	// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothPDUTimestamp
	IOBluetoothPDUTimestamp string
	// IOBluetoothPDUType is the GSM type of the text message.
	//
	// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothPDUType
	IOBluetoothPDUType string
	// IOBluetoothPDUUserData is the content of the text message.
	//
	// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothPDUUserData
	IOBluetoothPDUUserData string
)

var (
	// KFTSListingNameKey is nSString value. This key is used with the array of NSDictionary’s returned through the delegate method fileTransferServicesGetListingComplete: after calling getFolderListing.
	//
	// See: https://developer.apple.com/documentation/IOBluetooth/kFTSListingNameKey
	KFTSListingNameKey string
	// KFTSListingSizeKey is int value. This key is used with the array of NSDictionary’s returned through the delegate method fileTransferServicesGetListingComplete: after calling getFolderListing.
	//
	// See: https://developer.apple.com/documentation/IOBluetooth/kFTSListingSizeKey
	KFTSListingSizeKey string
	// KFTSListingTypeKey is fTSFileType value. This key is used with the array of NSDictionary’s returned through the delegate method fileTransferServicesGetListingComplete: after calling getFolderListing.
	//
	// See: https://developer.apple.com/documentation/IOBluetooth/kFTSListingTypeKey
	KFTSListingTypeKey string
	// KFTSProgressBytesTotalKey is nSNumber integer value. This key is used with the NSDictionary returned from the fileTransferServicesPutProgress: and fileTransferServicesGetProgress: delegate methods.
	//
	// See: https://developer.apple.com/documentation/IOBluetooth/kFTSProgressBytesTotalKey
	KFTSProgressBytesTotalKey string
	// KFTSProgressBytesTransferredKey is nSNumber integer value. This key is used with the NSDictionary returned from the fileTransferServicesPutProgress: and fileTransferServicesGetProgress: delegate methods.
	//
	// See: https://developer.apple.com/documentation/IOBluetooth/kFTSProgressBytesTransferredKey
	KFTSProgressBytesTransferredKey string
	// KFTSProgressEstimatedTimeKey is nSNumber double value. This key is used with the NSDictionary returned from the fileTransferServicesPutProgress: and fileTransferServicesGetProgress: delegate methods.
	//
	// See: https://developer.apple.com/documentation/IOBluetooth/kFTSProgressEstimatedTimeKey
	KFTSProgressEstimatedTimeKey string
	// KFTSProgressPercentageKey is nSNumber float value. This key is used with the NSDictionary returned from the fileTransferServicesPutProgress: and fileTransferServicesGetProgress: delegate methods.
	//
	// See: https://developer.apple.com/documentation/IOBluetooth/kFTSProgressPercentageKey
	KFTSProgressPercentageKey string
	// See: https://developer.apple.com/documentation/IOBluetooth/kFTSProgressPrecentageKey
	KFTSProgressPrecentageKey string
	// KFTSProgressTimeElapsedKey is nSNumber int value. This key is used with the NSDictionary returned from the fileTransferServicesPutProgress: and fileTransferServicesGetProgress: delegate methods.
	//
	// See: https://developer.apple.com/documentation/IOBluetooth/kFTSProgressTimeElapsedKey
	KFTSProgressTimeElapsedKey string
	// KFTSProgressTransferRateKey is nSNumber float value. This key is used with the NSDictionary returned from the fileTransferServicesPutProgress: and fileTransferServicesGetProgress: delegate methods.
	//
	// See: https://developer.apple.com/documentation/IOBluetooth/kFTSProgressTransferRateKey
	KFTSProgressTransferRateKey string
	// See: https://developer.apple.com/documentation/IOBluetooth/kOBEXHeaderIDKeyAppParameters
	KOBEXHeaderIDKeyAppParameters string
	// See: https://developer.apple.com/documentation/IOBluetooth/kOBEXHeaderIDKeyAuthorizationChallenge
	KOBEXHeaderIDKeyAuthorizationChallenge string
	// See: https://developer.apple.com/documentation/IOBluetooth/kOBEXHeaderIDKeyAuthorizationResponse
	KOBEXHeaderIDKeyAuthorizationResponse string
	// See: https://developer.apple.com/documentation/IOBluetooth/kOBEXHeaderIDKeyBody
	KOBEXHeaderIDKeyBody string
	// See: https://developer.apple.com/documentation/IOBluetooth/kOBEXHeaderIDKeyByteSequence
	KOBEXHeaderIDKeyByteSequence string
	// See: https://developer.apple.com/documentation/IOBluetooth/kOBEXHeaderIDKeyConnectionID
	KOBEXHeaderIDKeyConnectionID string
	// See: https://developer.apple.com/documentation/IOBluetooth/kOBEXHeaderIDKeyCount
	KOBEXHeaderIDKeyCount string
	// See: https://developer.apple.com/documentation/IOBluetooth/kOBEXHeaderIDKeyDescription
	KOBEXHeaderIDKeyDescription string
	// See: https://developer.apple.com/documentation/IOBluetooth/kOBEXHeaderIDKeyEndOfBody
	KOBEXHeaderIDKeyEndOfBody string
	// See: https://developer.apple.com/documentation/IOBluetooth/kOBEXHeaderIDKeyHTTP
	KOBEXHeaderIDKeyHTTP string
	// See: https://developer.apple.com/documentation/IOBluetooth/kOBEXHeaderIDKeyLength
	KOBEXHeaderIDKeyLength string
	// See: https://developer.apple.com/documentation/IOBluetooth/kOBEXHeaderIDKeyName
	KOBEXHeaderIDKeyName string
	// See: https://developer.apple.com/documentation/IOBluetooth/kOBEXHeaderIDKeyObjectClass
	KOBEXHeaderIDKeyObjectClass string
	// See: https://developer.apple.com/documentation/IOBluetooth/kOBEXHeaderIDKeyTarget
	KOBEXHeaderIDKeyTarget string
	// See: https://developer.apple.com/documentation/IOBluetooth/kOBEXHeaderIDKeyTime4Byte
	KOBEXHeaderIDKeyTime4Byte string
	// See: https://developer.apple.com/documentation/IOBluetooth/kOBEXHeaderIDKeyTimeISO
	KOBEXHeaderIDKeyTimeISO string
	// See: https://developer.apple.com/documentation/IOBluetooth/kOBEXHeaderIDKeyType
	KOBEXHeaderIDKeyType string
	// See: https://developer.apple.com/documentation/IOBluetooth/kOBEXHeaderIDKeyUnknown1ByteQuantity
	KOBEXHeaderIDKeyUnknown1ByteQuantity string
	// See: https://developer.apple.com/documentation/IOBluetooth/kOBEXHeaderIDKeyUnknown4ByteQuantity
	KOBEXHeaderIDKeyUnknown4ByteQuantity string
	// See: https://developer.apple.com/documentation/IOBluetooth/kOBEXHeaderIDKeyUnknownByteSequence
	KOBEXHeaderIDKeyUnknownByteSequence string
	// See: https://developer.apple.com/documentation/IOBluetooth/kOBEXHeaderIDKeyUnknownUnicodeText
	KOBEXHeaderIDKeyUnknownUnicodeText string
	// See: https://developer.apple.com/documentation/IOBluetooth/kOBEXHeaderIDKeyUserDefined
	KOBEXHeaderIDKeyUserDefined string
	// See: https://developer.apple.com/documentation/IOBluetooth/kOBEXHeaderIDKeyWho
	KOBEXHeaderIDKeyWho string
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOBluetoothHandsFreeCallDirection"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOBluetoothHandsFreeCallDirection = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOBluetoothHandsFreeCallIndex"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOBluetoothHandsFreeCallIndex = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOBluetoothHandsFreeCallMode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOBluetoothHandsFreeCallMode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOBluetoothHandsFreeCallMultiparty"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOBluetoothHandsFreeCallMultiparty = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOBluetoothHandsFreeCallName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOBluetoothHandsFreeCallName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOBluetoothHandsFreeCallNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOBluetoothHandsFreeCallNumber = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOBluetoothHandsFreeCallStatus"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOBluetoothHandsFreeCallStatus = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOBluetoothHandsFreeCallType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOBluetoothHandsFreeCallType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOBluetoothHandsFreeIndicatorBattChg"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOBluetoothHandsFreeIndicatorBattChg = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOBluetoothHandsFreeIndicatorCall"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOBluetoothHandsFreeIndicatorCall = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOBluetoothHandsFreeIndicatorCallHeld"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOBluetoothHandsFreeIndicatorCallHeld = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOBluetoothHandsFreeIndicatorCallSetup"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOBluetoothHandsFreeIndicatorCallSetup = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOBluetoothHandsFreeIndicatorRoam"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOBluetoothHandsFreeIndicatorRoam = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOBluetoothHandsFreeIndicatorService"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOBluetoothHandsFreeIndicatorService = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOBluetoothHandsFreeIndicatorSignal"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOBluetoothHandsFreeIndicatorSignal = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOBluetoothHostControllerPoweredOffNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOBluetoothHostControllerPoweredOffNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOBluetoothHostControllerPoweredOnNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOBluetoothHostControllerPoweredOnNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOBluetoothL2CAPChannelPublishedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOBluetoothL2CAPChannelPublishedNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOBluetoothL2CAPChannelTerminatedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOBluetoothL2CAPChannelTerminatedNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOBluetoothPDUEncoding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOBluetoothPDUEncoding = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOBluetoothPDUOriginatingAddress"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOBluetoothPDUOriginatingAddress = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOBluetoothPDUOriginatingAddressType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOBluetoothPDUOriginatingAddressType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOBluetoothPDUProtocolID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOBluetoothPDUProtocolID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOBluetoothPDUServicCenterAddress"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOBluetoothPDUServicCenterAddress = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOBluetoothPDUServiceCenterAddressType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOBluetoothPDUServiceCenterAddressType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOBluetoothPDUTimestamp"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOBluetoothPDUTimestamp = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOBluetoothPDUType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOBluetoothPDUType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOBluetoothPDUUserData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOBluetoothPDUUserData = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kFTSListingNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KFTSListingNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kFTSListingSizeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KFTSListingSizeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kFTSListingTypeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KFTSListingTypeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kFTSProgressBytesTotalKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KFTSProgressBytesTotalKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kFTSProgressBytesTransferredKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KFTSProgressBytesTransferredKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kFTSProgressEstimatedTimeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KFTSProgressEstimatedTimeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kFTSProgressPercentageKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KFTSProgressPercentageKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kFTSProgressPrecentageKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KFTSProgressPrecentageKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kFTSProgressTimeElapsedKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KFTSProgressTimeElapsedKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kFTSProgressTransferRateKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KFTSProgressTransferRateKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kOBEXHeaderIDKeyAppParameters"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KOBEXHeaderIDKeyAppParameters = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kOBEXHeaderIDKeyAuthorizationChallenge"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KOBEXHeaderIDKeyAuthorizationChallenge = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kOBEXHeaderIDKeyAuthorizationResponse"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KOBEXHeaderIDKeyAuthorizationResponse = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kOBEXHeaderIDKeyBody"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KOBEXHeaderIDKeyBody = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kOBEXHeaderIDKeyByteSequence"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KOBEXHeaderIDKeyByteSequence = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kOBEXHeaderIDKeyConnectionID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KOBEXHeaderIDKeyConnectionID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kOBEXHeaderIDKeyCount"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KOBEXHeaderIDKeyCount = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kOBEXHeaderIDKeyDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KOBEXHeaderIDKeyDescription = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kOBEXHeaderIDKeyEndOfBody"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KOBEXHeaderIDKeyEndOfBody = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kOBEXHeaderIDKeyHTTP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KOBEXHeaderIDKeyHTTP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kOBEXHeaderIDKeyLength"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KOBEXHeaderIDKeyLength = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kOBEXHeaderIDKeyName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KOBEXHeaderIDKeyName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kOBEXHeaderIDKeyObjectClass"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KOBEXHeaderIDKeyObjectClass = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kOBEXHeaderIDKeyTarget"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KOBEXHeaderIDKeyTarget = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kOBEXHeaderIDKeyTime4Byte"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KOBEXHeaderIDKeyTime4Byte = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kOBEXHeaderIDKeyTimeISO"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KOBEXHeaderIDKeyTimeISO = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kOBEXHeaderIDKeyType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KOBEXHeaderIDKeyType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kOBEXHeaderIDKeyUnknown1ByteQuantity"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KOBEXHeaderIDKeyUnknown1ByteQuantity = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kOBEXHeaderIDKeyUnknown4ByteQuantity"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KOBEXHeaderIDKeyUnknown4ByteQuantity = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kOBEXHeaderIDKeyUnknownByteSequence"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KOBEXHeaderIDKeyUnknownByteSequence = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kOBEXHeaderIDKeyUnknownUnicodeText"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KOBEXHeaderIDKeyUnknownUnicodeText = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kOBEXHeaderIDKeyUserDefined"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KOBEXHeaderIDKeyUserDefined = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kOBEXHeaderIDKeyWho"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KOBEXHeaderIDKeyWho = objc.GoString(cstr)
			}
		}
	}

}
