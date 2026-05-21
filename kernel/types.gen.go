// Code generated from Apple documentation for kernel. DO NOT EDIT.

package kernel

import (
	"unsafe"
)

// C struct types

// ATADeviceNub - ATADeviceNub is a concrete implementation of IOATADevice.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/atadevicenub
type ATADeviceNub struct {
	Reserved unsafe.Pointer
	Attach   unsafe.Pointer // override of IOService method.
	Init     unsafe.Pointer // used after creating the nub.

}

// ATATimerEventSource - Extend the timer event source to allow checking for timer expiration from behind the workloop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/atatimereventsource
type ATATimerEventSource struct {
	Reserved unsafe.Pointer
	Disable  unsafe.Pointer // overrides in order to set/clear the timed out flag
	Enable   unsafe.Pointer // overrides in order to set/clear the timed out flag
	Init     unsafe.Pointer
}

// AVCCommandHandlerInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/avccommandhandlerinfo
type AVCCommandHandlerInfo struct {
}

// AVCConnectionRecord
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/avcconnectionrecord
type AVCConnectionRecord struct {
}

// AVCSubunitInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/avcsubunitinfo
type AVCSubunitInfo struct {
	Create unsafe.Pointer
	Free   unsafe.Pointer
	Init   unsafe.Pointer
}

// AppleMacIO
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/applemacio
type AppleMacIO struct {
	Start unsafe.Pointer
}

// AppleMacIODevice
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/applemaciodevice
type AppleMacIODevice struct {
}

// AppleNMI
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/applenmi
type AppleNMI struct {
	Start unsafe.Pointer
}

// ApplePlatformExpert
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/appleplatformexpert
type ApplePlatformExpert struct {
	Configure unsafe.Pointer
	Start     unsafe.Pointer
}

// AsyncPendingTrans
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/asyncpendingtrans
type AsyncPendingTrans struct {
	FAltHandler *IOFWCommand
	FTCode      unsafe.Pointer
	FInUse      unsafe.Pointer
	FHandler    *IOFWAsyncCommand
}

// BigSInt16
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/bigsint16
type BigSInt16 struct {
	Get unsafe.Pointer
}

// BigSInt32
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/bigsint32
type BigSInt32 struct {
	Get unsafe.Pointer
}

// BigSInt64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/bigsint64
type BigSInt64 struct {
	Get unsafe.Pointer
}

// BigUInt16
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/biguint16
type BigUInt16 struct {
	Get unsafe.Pointer
}

// BigUInt32
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/biguint32
type BigUInt32 struct {
	Get unsafe.Pointer
}

// BigUInt64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/biguint64
type BigUInt64 struct {
	Get unsafe.Pointer
}

// DebugKeyAction
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/debugkeyaction
type DebugKeyAction struct {
	Action           unsafe.Pointer
	Parameter        unsafe.Pointer
	Mask             unsafe.Pointer
	DebugArgRequired unsafe.Pointer
}

// FWSegment
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/fwsegment
type FWSegment struct {
	Address unsafe.Pointer
	Length  unsafe.Pointer
}

// FndrExtendedDirInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/fndrextendeddirinfo
type FndrExtendedDirInfo struct {
	Document_id       U_int32_t
	Date_added        U_int32_t
	Extended_flags    U_int16_t
	Reserved3         U_int16_t
	Write_gen_counter U_int32_t
}

// FndrExtendedFileInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/fndrextendedfileinfo
type FndrExtendedFileInfo struct {
	Document_id       U_int32_t
	Date_added        U_int32_t
	Extended_flags    U_int16_t
	Reserved2         U_int16_t
	Write_gen_counter U_int32_t
}

// IOACPIAddressSpaceDescriptor
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioacpiaddressspacedescriptor
type IOACPIAddressSpaceDescriptor struct {
	MaxAddressRange   uint64
	Reserved1         unsafe.Pointer
	TypeSpecificFlags unsafe.Pointer
	GeneralFlags      unsafe.Pointer
	ResourceType      unsafe.Pointer
	TranslationOffset uint64
	Reserved3         uint64
	MinAddressRange   uint64
	Reserved2         uint64
	Granularity       uint64
	AddressLength     uint64
	Reserved4         uint64
}

// IOACPIPlatformDevice
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioacpiplatformdevice
type IOACPIPlatformDevice struct {
	Free unsafe.Pointer
	Init unsafe.Pointer
}

// IOACPIPlatformExpert
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioacpiplatformexpert
type IOACPIPlatformExpert struct {
	Start unsafe.Pointer
}

// IOAGPDevice - An IOService class representing an AGP primary device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioagpdevice
type IOAGPDevice struct {
	Reserved unsafe.Pointer
}

// IOATABusCommand
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioatabuscommand
type IOATABusCommand struct {
	Syncer   unsafe.Pointer
	State    unsafe.Pointer
	Reserved unsafe.Pointer
	Init     unsafe.Pointer // Zeroes all data, returns false if allocation fails. protected.

}

// IOATABusCommand64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioatabuscommand64
type IOATABusCommand64 struct {
	Free unsafe.Pointer
	Init unsafe.Pointer
}

// IOATABusInfo - used to indicate the capabilities of the bus the device is connected to, PIO and DMA modes supported, etc.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioatabusinfo
type IOATABusInfo struct {
	Reserved   unsafe.Pointer
	Atabusinfo unsafe.Pointer // factory method
	Init       unsafe.Pointer
}

// IOATACommand
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioatacommand
type IOATACommand struct {
	Free unsafe.Pointer
	Init unsafe.Pointer
}

// IOATAController - The base class for ata controller family. Provides the interface common to all ata bus controllers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioatacontroller
type IOATAController struct {
	Reserved unsafe.Pointer
	Free     unsafe.Pointer
	Init     unsafe.Pointer
	Probe    unsafe.Pointer
	Start    unsafe.Pointer
}

// IOATADevConfig - used for configuring and communicating the desired transfer modes of a device. A disk driver would typically use this object in conjunction with the 512-bytes of identification data from the drive and the IOATABusInfo object for the bus it is connected to. This object will determine the best matching transfer speeds available. the device driver will then send a series of Set Features commands to configure the drive and this object to the bus through the IOATADevice nub in order to configure the optimum transfer mode. The driver for the disk drive may choose to populate this object with whatever transfer mode desired, in the event that a different mode is required.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioatadevconfig
type IOATADevConfig struct {
	Reserved     unsafe.Pointer
	Atadevconfig unsafe.Pointer // static creator function.
	Init         unsafe.Pointer
}

// IOATADevice - This object implements a relay to an ATA Bus where a drive is attached.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioatadevice
type IOATADevice struct {
	Reserved unsafe.Pointer
}

// IOATAIOReg16
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioataioreg16
type IOATAIOReg16 struct {
}

// IOATAIOReg32
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioataioreg32
type IOATAIOReg32 struct {
}

// IOATAIOReg8
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioataioreg8
type IOATAIOReg8 struct {
}

// IOATAPIProtocolTransport - SCSI Protocol Driver Family for ATAPI Devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioatapiprotocoltransport
type IOATAPIProtocolTransport struct {
	Free    unsafe.Pointer
	Init    unsafe.Pointer
	Message unsafe.Pointer
	Start   unsafe.Pointer
	Stop    unsafe.Pointer
}

// IOATAReg16
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioatareg16
type IOATAReg16 struct {
}

// IOATAReg32
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioatareg32
type IOATAReg32 struct {
}

// IOATAReg8
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioatareg8
type IOATAReg8 struct {
}

// IOAccelerator
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioaccelerator
type IOAccelerator struct {
}

// IOAddressSegment
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioaddresssegment
type IOAddressSegment struct {
	Address unsafe.Pointer
	Length  unsafe.Pointer
}

// IOAppleLabelScheme
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioapplelabelscheme
type IOAppleLabelScheme struct {
	Free  unsafe.Pointer
	Init  unsafe.Pointer
	Probe unsafe.Pointer
	Scan  unsafe.Pointer
	Start unsafe.Pointer
	Stop  unsafe.Pointer
}

// IOApplePartitionScheme
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioapplepartitionscheme
type IOApplePartitionScheme struct {
	Free  unsafe.Pointer
	Init  unsafe.Pointer
	Probe unsafe.Pointer
	Scan  unsafe.Pointer
	Start unsafe.Pointer
	Stop  unsafe.Pointer
}

// IOAudioControl - Represents any controllable attribute of an IOAudioDevice.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioaudiocontrol
type IOAudioControl struct {
	Free  unsafe.Pointer // Frees all of the resources allocated by the IOAudioControl.
	Init  unsafe.Pointer // Initializes a newly allocated IOAudioControl with the given attributes.
	Start unsafe.Pointer // Starts a newly created IOAudioControl.
	Stop  unsafe.Pointer // Stops the control when the provider is going away.

}

// IOAudioControlUserClient
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioaudiocontroluserclient
type IOAudioControlUserClient struct {
	Free unsafe.Pointer
}

// IOAudioDevice - Abstract base class for a single piece of audio hardware. The IOAudioDevice provides the central coordination point for an audio driver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioaudiodevice
type IOAudioDevice struct {
	Free  unsafe.Pointer // Frees resources used by the IOAudioDevice instance
	Init  unsafe.Pointer // Initialize a newly created instance of IOAudioDevice.
	Start unsafe.Pointer // This function is called automatically by the system to tell the driver to start vending services to the rest of the system.
	Stop  unsafe.Pointer // This is responsible for stopping the device after the system is done with it (or if the device is removed from the system).

}

// IOAudioEngine - Abstract base class for a single audio audio / I/O engine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioaudioengine
type IOAudioEngine struct {
	Status unsafe.Pointer
	State  unsafe.Pointer
	Free   unsafe.Pointer // Frees all of the resources allocated by the IOAudioEngine.
	Init   unsafe.Pointer // Performs initialization of a newly allocated IOAudioEngine.
	Start  unsafe.Pointer
	Stop   unsafe.Pointer // Stops the service and prepares for the driver to be terminated.

}

// IOAudioEngineUserClient
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioaudioengineuserclient
type IOAudioEngineUserClient struct {
	Free unsafe.Pointer
	Stop unsafe.Pointer
}

// IOAudioLevelControl
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioaudiolevelcontrol
type IOAudioLevelControl struct {
	Create unsafe.Pointer // Allocates a new level control with the given attributes
	Free   unsafe.Pointer
	Init   unsafe.Pointer // Initializes a newly allocated IOAudioLevelControl with the given attributes

}

// IOAudioPort - Represents a logical or physical port or functional unit in an audio device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioaudioport
type IOAudioPort struct {
	Free  unsafe.Pointer // Frees all of the resources allocated by the IOAudioPort.
	Start unsafe.Pointer // Called to start a newly created IOAudioPort.
	Stop  unsafe.Pointer // Called when the IOAudioDevice is stopping when it is no longer available.

}

// IOAudioSelectorControl
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioaudioselectorcontrol
type IOAudioSelectorControl struct {
	Create unsafe.Pointer
	Free   unsafe.Pointer
	Init   unsafe.Pointer
}

// IOAudioStream - This class wraps a single sample buffer in an audio driver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioaudiostream
type IOAudioStream struct {
	Free unsafe.Pointer
	Stop unsafe.Pointer
}

// IOAudioToggleControl
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioaudiotogglecontrol
type IOAudioToggleControl struct {
	Create unsafe.Pointer // Allocates a new mute control with the given attributes
	Init   unsafe.Pointer // Initializes a newly allocated IOAudioToggleControl with the given attributes

}

// IOBDBlockStorageDevice - The IOBDBlockStorageDevice class is a generic BD block storage device abstraction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iobdblockstoragedevice
type IOBDBlockStorageDevice struct {
	Init unsafe.Pointer
}

// IOBDBlockStorageDriver
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iobdblockstoragedriver
type IOBDBlockStorageDriver struct {
}

// IOBDMedia - The IOBDMedia class is a random-access disk device abstraction for BDs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iobdmedia
type IOBDMedia struct {
}

// IOBDMediaBSDClient
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iobdmediabsdclient
type IOBDMediaBSDClient struct {
	Ioctl unsafe.Pointer
}

// IOBDServices
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iobdservices
type IOBDServices struct {
	Free    unsafe.Pointer
	Message unsafe.Pointer
	Open    unsafe.Pointer
	Start   unsafe.Pointer
}

// IOBacklightDisplay
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iobacklightdisplay
type IOBacklightDisplay struct {
}

// IOBasicOutputQueue - A concrete implementation of an IOOutputQueue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iobasicoutputqueue
type IOBasicOutputQueue struct {
	Dequeue unsafe.Pointer
	Enqueue unsafe.Pointer // Adds a packet, or a chain of packets, to the queue.
	Flush   unsafe.Pointer // Drops and frees all packets currently held by the queue.
	Free    unsafe.Pointer // Frees the IOBasicOutputQueue object.
	Init    unsafe.Pointer // Initializes an IOBasicOutputQueue object.
	Output  unsafe.Pointer // Transfers all packets in the mbuf queue to the target.
	Service unsafe.Pointer // Services a queue that was stalled by the target.
	Start   unsafe.Pointer // Starts up the packet flow between the queue and its target.
	Stop    unsafe.Pointer // Stops the packet flow between the queue and its target.

}

// IOBigMemoryCursor - An IOMemoryCursor subclass that outputs a vector of PhysicalSegments in the big endian byte order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iobigmemorycursor
type IOBigMemoryCursor struct {
}

// IOBlockStorageDevice - A generic block storage device abstraction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioblockstoragedevice
type IOBlockStorageDevice struct {
	Init unsafe.Pointer
}

// IOBlockStorageDeviceExtent - Extent for unmap storage requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioblockstoragedeviceextent
type IOBlockStorageDeviceExtent struct {
	BlockStart uint64 // The starting block number of the operation.
	BlockCount uint64 // The integral number of blocks to be deleted.

}

// IOBlockStorageDriver - The common base class for generic block storage drivers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioblockstoragedriver
type IOBlockStorageDriver struct {
	Free        unsafe.Pointer
	Init        unsafe.Pointer
	Message     unsafe.Pointer
	Read        unsafe.Pointer
	Start       unsafe.Pointer
	Stop        unsafe.Pointer
	Synchronize unsafe.Pointer
	Unmap       unsafe.Pointer
	Write       unsafe.Pointer
	Yield       unsafe.Pointer
}

// IOBlockStorageProvisionDeviceExtent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioblockstorageprovisiondeviceextent
type IOBlockStorageProvisionDeviceExtent struct {
	ProvisionType unsafe.Pointer
	BlockStart    uint64
	BlockCount    uint64
	Reserved      unsafe.Pointer
}

// IOBlockStorageServices
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioblockstorageservices
type IOBlockStorageServices struct {
	Attach    unsafe.Pointer
	Detach    unsafe.Pointer
	Free      unsafe.Pointer
	Message   unsafe.Pointer
	Terminate unsafe.Pointer
}

// IOBufferMemoryDescriptor - A simple memory descriptor that allocates its own buffer memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iobuffermemorydescriptor
type IOBufferMemoryDescriptor struct {
	Free     unsafe.Pointer // Performs any final cleanup for the memory buffer descriptor object.
	Reserved unsafe.Pointer
}

// IOCDBlockStorageDevice - The IOCDBlockStorageDevice class is a generic CD block storage device abstraction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iocdblockstoragedevice
type IOCDBlockStorageDevice struct {
	Init unsafe.Pointer
}

// IOCDBlockStorageDriver
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iocdblockstoragedriver
type IOCDBlockStorageDriver struct {
	Free unsafe.Pointer
	Init unsafe.Pointer
}

// IOCDMedia - The IOCDMedia class is a random-access disk device abstraction for CDs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iocdmedia
type IOCDMedia struct {
	Read  unsafe.Pointer
	Write unsafe.Pointer
}

// IOCDMediaBSDClient
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iocdmediabsdclient
type IOCDMediaBSDClient struct {
	Ioctl unsafe.Pointer
}

// IOCDPartitionScheme
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iocdpartitionscheme
type IOCDPartitionScheme struct {
	Free  unsafe.Pointer
	Init  unsafe.Pointer
	Probe unsafe.Pointer
	Read  unsafe.Pointer
	Scan  unsafe.Pointer
	Start unsafe.Pointer
	Write unsafe.Pointer
}

// IOCatalogue - In-kernel database for IOKit driver personalities.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iocatalogue
type IOCatalogue struct {
	Free       unsafe.Pointer
	Init       unsafe.Pointer
	Initialize unsafe.Pointer
	Reset      unsafe.Pointer
	Serialize  unsafe.Pointer
}

// IOCommand - This class is an abstract class which represents an I/O command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iocommand
type IOCommand struct {
	Free unsafe.Pointer
	Init unsafe.Pointer
}

// IOCommandGate - Single-threaded work-loop client request mechanism.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iocommandgate
type IOCommandGate struct {
	Reserved unsafe.Pointer
	Disable  unsafe.Pointer // Disable the command gate
	Enable   unsafe.Pointer // Enable command gate, this will unblock any blocked Commands and Actions.
	Free     unsafe.Pointer
	Init     unsafe.Pointer // Class initialiser.

}

// IOCommandPool - Manipulates a pool of commands which inherit from IOCommand.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iocommandpool
type IOCommandPool struct {
	Reserved unsafe.Pointer
	Free     unsafe.Pointer
	Init     unsafe.Pointer // Should never be used, obsolete. See initWithWorkLoop.

}

// IOCompactDiscServices
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iocompactdiscservices
type IOCompactDiscServices struct {
	Free    unsafe.Pointer
	Message unsafe.Pointer
	Open    unsafe.Pointer
	Start   unsafe.Pointer
}

// IOConditionLock
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioconditionlock
type IOConditionLock struct {
	Free   unsafe.Pointer
	Lock   unsafe.Pointer
	Unlock unsafe.Pointer
}

// IOConfigDirectory
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioconfigdirectory
type IOConfigDirectory struct {
	Reserved unsafe.Pointer
	Update   unsafe.Pointer
}

// IODCLProgram
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iodclprogram
type IODCLProgram struct {
	Compile unsafe.Pointer
	Free    unsafe.Pointer
	Init    unsafe.Pointer
	Notify  unsafe.Pointer
	Pause   unsafe.Pointer
	Resume  unsafe.Pointer
	Start   unsafe.Pointer
	Stop    unsafe.Pointer
}

// IODCLTranslateListen
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iodcltranslatelisten
type IODCLTranslateListen struct {
	Compile unsafe.Pointer
	Start   unsafe.Pointer
}

// IODCLTranslateTalk
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iodcltranslatetalk
type IODCLTranslateTalk struct {
	Compile unsafe.Pointer
	Start   unsafe.Pointer
}

// IODCLTranslator
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iodcltranslator
type IODCLTranslator struct {
	Init   unsafe.Pointer
	Notify unsafe.Pointer
	Stop   unsafe.Pointer
}

// IODMACommand - An object that converts memory references to I/O bus addresses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iodmacommand
type IODMACommand struct {
	Complete    unsafe.Pointer // Complete processing of DMA mappings after an I/O transfer is finished.
	Free        unsafe.Pointer
	Init        unsafe.Pointer
	Prepare     unsafe.Pointer // Prepare the memory for an I/O transfer.
	Synchronize unsafe.Pointer // Bring IOMemoryDescriptor and IODMACommand buffers into sync.

}

// IODMACommandSpecification
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iodmacommandspecification
type IODMACommandSpecification struct {
	MaxAddressBits unsafe.Pointer
	Options        unsafe.Pointer
}

// IODMAController
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iodmacontroller
type IODMAController struct {
	Start unsafe.Pointer
}

// IODMAEventSource
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iodmaeventsource
type IODMAEventSource struct {
	Free unsafe.Pointer
	Init unsafe.Pointer
}

// IODMAMapPageList
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iodmamappagelist
type IODMAMapPageList struct {
	PageList      *Upl_page_info_t
	PageOffset    unsafe.Pointer
	PageListCount unsafe.Pointer
}

// IODMAMapSpecification
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iodmamapspecification
type IODMAMapSpecification struct {
	Device         *IOService
	Options        unsafe.Pointer
	Alignment      unsafe.Pointer
	NumAddressBits uint8
	ResvA          uint8
	ResvB          unsafe.Pointer
}

// IODTNVRAM
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iodtnvram
type IODTNVRAM struct {
	Init   unsafe.Pointer
	Reload unsafe.Pointer
	Start  unsafe.Pointer
	Sync   unsafe.Pointer
}

// IODTPlatformExpert
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iodtplatformexpert
type IODTPlatformExpert struct {
	Configure unsafe.Pointer
	Probe     unsafe.Pointer
}

// IODVDBlockStorageDevice - The IODVDBlockStorageDevice class is a generic DVD block storage device abstraction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iodvdblockstoragedevice
type IODVDBlockStorageDevice struct {
	Init unsafe.Pointer
}

// IODVDBlockStorageDriver
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iodvdblockstoragedriver
type IODVDBlockStorageDriver struct {
}

// IODVDMedia - The IODVDMedia class is a random-access disk device abstraction for DVDs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iodvdmedia
type IODVDMedia struct {
}

// IODVDMediaBSDClient
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iodvdmediabsdclient
type IODVDMediaBSDClient struct {
	Ioctl unsafe.Pointer
}

// IODVDServices
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iodvdservices
type IODVDServices struct {
	Free    unsafe.Pointer
	Message unsafe.Pointer
	Open    unsafe.Pointer
	Start   unsafe.Pointer
}

// IODataQueue - A generic queue designed to pass data from the kernel to a user process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iodataqueue
type IODataQueue struct {
	Enqueue unsafe.Pointer // Enqueues a new entry on the queue.
	Free    unsafe.Pointer
}

// IODataQueueDispatchSource
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iodataqueuedispatchsource
type IODataQueueDispatchSource struct {
	Free unsafe.Pointer // Performs any final cleanup for the data-queue dispatch source.
	Init unsafe.Pointer // Handles the basic initialization of the dispatch source.

}

// IODataQueueDispatchSourceInterface
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iodataqueuedispatchsourceinterface
type IODataQueueDispatchSourceInterface struct {
}

// IODeviceMemory - An IOMemoryDescriptor used for device physical memory ranges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iodevicememory
type IODeviceMemory struct {
}

// IODispatchQueue
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iodispatchqueue
type IODispatchQueue struct {
	Free unsafe.Pointer // Performs any final cleanup for the dispatch queue object.
	Init unsafe.Pointer // Initializes the dispatch queue object.

}

// IODispatchQueueInterface
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iodispatchqueueinterface
type IODispatchQueueInterface struct {
}

// IODispatchSource
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iodispatchsource
type IODispatchSource struct {
	Free unsafe.Pointer // Performs any final cleanup for the dispatch source.
	Init unsafe.Pointer // Handles the basic initialization of the dispatch source.

}

// IODispatchSourceInterface
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iodispatchsourceinterface
type IODispatchSourceInterface struct {
}

// IODisplay
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iodisplay
type IODisplay struct {
	Free       unsafe.Pointer
	Initialize unsafe.Pointer
	Probe      unsafe.Pointer
	Start      unsafe.Pointer
	Stop       unsafe.Pointer
}

// IODisplayConnect
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iodisplayconnect
type IODisplayConnect struct {
}

// IODisplayParameterHandler
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iodisplayparameterhandler
type IODisplayParameterHandler struct {
}

// IOEthernetAddress
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioethernetaddress
type IOEthernetAddress struct {
	Bytes unsafe.Pointer
}

// IOEthernetController - Abstract superclass for Ethernet controllers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioethernetcontroller
type IOEthernetController struct {
	Free       unsafe.Pointer // Frees the IOEthernetController instance.
	Init       unsafe.Pointer // Initializes an IOEthernetController object.
	Initialize unsafe.Pointer // IOEthernetController class initializer.

}

// IOEthernetInterface - The Ethernet interface object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioethernetinterface
type IOEthernetInterface struct {
	Free unsafe.Pointer // Frees the IOEthernetInterface instance.
	Init unsafe.Pointer // Initializes an IOEthernetInterface instance.

}

// IOEventLink
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioeventlink
type IOEventLink struct {
	Free unsafe.Pointer
	Init unsafe.Pointer
}

// IOEventLinkInterface
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioeventlinkinterface
type IOEventLinkInterface struct {
}

// IOEventSource - Abstract class for all work-loop event sources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioeventsource
type IOEventSource struct {
	Reserved unsafe.Pointer
	Refcon   unsafe.Pointer
	Owner    unsafe.Pointer
	Enabled  unsafe.Pointer
	Disable  unsafe.Pointer // Disable event source.
	Enable   unsafe.Pointer // Enable event source.
	Free     unsafe.Pointer
	Init     unsafe.Pointer // Primary initialiser for the IOEventSource class.

}

// IOExtendedLBA - If 48-bit LBAs are supported, IOExtendedLBA is used to represent a 48-bit LBA. The driver examines the ATA identify data to determine if 48-bit addressing is supported.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioextendedlba
type IOExtendedLBA struct {
	Reserved unsafe.Pointer
}

// IOExternalAsyncMethod
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioexternalasyncmethod
type IOExternalAsyncMethod struct {
	Flags  uint32
	Object *IOService
	Count1 uint
	Count0 uint
	Func   IOAsyncMethod
}

// IOExternalMethod
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioexternalmethod
type IOExternalMethod struct {
	Object *IOService
	Count0 uint
	Flags  uint32
	Count1 uint
	Func   IOMethod
}

// IOExternalMethodArguments
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioexternalmethodarguments
type IOExternalMethodArguments struct {
	ScalarInput                   unsafe.Pointer
	Version                       unsafe.Pointer
	ScalarOutput                  unsafe.Pointer
	StructureInput                unsafe.Pointer
	StructureOutputDescriptorSize unsafe.Pointer
	StructureInputSize            unsafe.Pointer
	AsyncReference                *Io_user_reference_t
	Selector                      unsafe.Pointer
	StructureInputDescriptor      *IOMemoryDescriptor
	StructureOutput               unsafe.Pointer
	StructureOutputDescriptor     *IOMemoryDescriptor
	ScalarInputCount              unsafe.Pointer
	AsyncWakePort                 uint32
	StructureVariableOutputData   *OSObject
	StructureOutputSize           unsafe.Pointer
	AsyncReferenceCount           unsafe.Pointer
	ScalarOutputCount             unsafe.Pointer
}

// IOExternalMethodDispatch
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioexternalmethoddispatch
type IOExternalMethodDispatch struct {
	CheckScalarInputCount    unsafe.Pointer
	Function                 unsafe.Pointer
	CheckStructureOutputSize unsafe.Pointer
	CheckStructureInputSize  unsafe.Pointer
	CheckScalarOutputCount   unsafe.Pointer
}

// IOExternalMethodDispatch2022
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioexternalmethoddispatch2022
type IOExternalMethodDispatch2022 struct {
	AllowAsync               uint8
	CheckScalarInputCount    unsafe.Pointer
	CheckScalarOutputCount   unsafe.Pointer
	CheckStructureInputSize  unsafe.Pointer
	CheckStructureOutputSize unsafe.Pointer
	Function                 unsafe.Pointer
	CheckEntitlement         unsafe.Pointer
}

// IOExternalTrap
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioexternaltrap
type IOExternalTrap struct {
	Object *IOService
	Func   IOTrap
}

// IOFDiskPartitionScheme
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofdiskpartitionscheme
type IOFDiskPartitionScheme struct {
	Free  unsafe.Pointer
	Init  unsafe.Pointer
	Probe unsafe.Pointer
	Scan  unsafe.Pointer
	Start unsafe.Pointer
	Stop  unsafe.Pointer
}

// IOFWAddressSpace
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofwaddressspace
type IOFWAddressSpace struct {
	Activate   unsafe.Pointer // Address space is ready for handling requests.
	Contains   unsafe.Pointer // returns number of bytes starting at addr in this space
	Deactivate unsafe.Pointer // Address space request handler is disabled.
	Free       unsafe.Pointer
	Init       unsafe.Pointer
	Intersects unsafe.Pointer // Checks this address space intersects with the given address range. Currently only supports IOFWPsuedoAddressSpaces.

}

// IOFWAddressSpaceAux
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofwaddressspaceaux
type IOFWAddressSpaceAux struct {
	Free       unsafe.Pointer
	Init       unsafe.Pointer
	Intersects unsafe.Pointer
}

// IOFWAsyncCommand - Send an async request to a device
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofwasynccommand
type IOFWAsyncCommand struct {
	Complete unsafe.Pointer
	Free     unsafe.Pointer
	Reinit   unsafe.Pointer
}

// IOFWAsyncPHYCommand - Send an async PHY packet
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofwasyncphycommand
type IOFWAsyncPHYCommand struct {
	Complete unsafe.Pointer
	Execute  unsafe.Pointer
	Free     unsafe.Pointer
	Reinit   unsafe.Pointer
}

// IOFWAsyncStreamCommand - Send an async stream packet
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofwasyncstreamcommand
type IOFWAsyncStreamCommand struct {
	Complete unsafe.Pointer
	Execute  unsafe.Pointer
	Free     unsafe.Pointer
	Reinit   unsafe.Pointer
}

// IOFWBusCommand - Bus control commands
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofwbuscommand
type IOFWBusCommand struct {
	Reserved unsafe.Pointer
	Complete unsafe.Pointer
	Reinit   unsafe.Pointer
}

// IOFWCmdQ - Structure for head of a queue of IOFWCommands
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofwcmdq
type IOFWCmdQ struct {
}

// IOFWCommand - Base class for FireWire commands
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofwcommand
type IOFWCommand struct {
	Cancel   unsafe.Pointer
	Complete unsafe.Pointer
	Execute  unsafe.Pointer
	Free     unsafe.Pointer
	Submit   unsafe.Pointer
}

// IOFWCompareAndSwapCommand
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofwcompareandswapcommand
type IOFWCompareAndSwapCommand struct {
	Execute unsafe.Pointer
	Free    unsafe.Pointer
	Locked  unsafe.Pointer
	Reinit  unsafe.Pointer
}

// IOFWDCL
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofwdcl
type IOFWDCL struct {
	Compile   unsafe.Pointer
	Debug     unsafe.Pointer
	Finalize  unsafe.Pointer
	Free      unsafe.Pointer
	Interrupt unsafe.Pointer
	Link      unsafe.Pointer
	Update    unsafe.Pointer
}

// IOFWDelayCommand - Command to execute some code after a specified delay (in microseconds) All it does is timeout after the specified delay, hence calling the completion callback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofwdelaycommand
type IOFWDelayCommand struct {
	Reserved unsafe.Pointer
	Execute  unsafe.Pointer
	Reinit   unsafe.Pointer
}

// IOFWIsochChannel
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofwisochchannel
type IOFWIsochChannel struct {
	Free  unsafe.Pointer
	Init  unsafe.Pointer
	Start unsafe.Pointer
	Stop  unsafe.Pointer
}

// IOFWIsochPort
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofwisochport
type IOFWIsochPort struct {
	Start unsafe.Pointer
	Stop  unsafe.Pointer
}

// IOFWLocalIsochPort
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofwlocalisochport
type IOFWLocalIsochPort struct {
	Free   unsafe.Pointer
	Init   unsafe.Pointer
	Notify unsafe.Pointer // Informs hardware of a change to the DCL program.
	Start  unsafe.Pointer
	Stop   unsafe.Pointer
}

// IOFWNodeScan
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofwnodescan
type IOFWNodeScan struct {
	Generation unsafe.Pointer
}

// IOFWPHYPacketListener
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofwphypacketlistener
type IOFWPHYPacketListener struct {
	Activate   unsafe.Pointer
	Deactivate unsafe.Pointer
	Free       unsafe.Pointer
}

// IOFWPhysicalAddressSpace
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofwphysicaladdressspace
type IOFWPhysicalAddressSpace struct {
	Complete    unsafe.Pointer // complete the IODMACommand used by this PhysicalAddressSpace.
	Free        unsafe.Pointer
	Init        unsafe.Pointer // Initialize physical address space.
	Prepare     unsafe.Pointer // Prepare the IODMACommand used by this PhysicalAddressSpace.
	Synchronize unsafe.Pointer // synchronize the IODMACommand used by this PhysicalAddressSpace.

}

// IOFWPhysicalAddressSpaceAux
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofwphysicaladdressspaceaux
type IOFWPhysicalAddressSpaceAux struct {
	Complete    unsafe.Pointer
	Free        unsafe.Pointer
	Init        unsafe.Pointer
	Prepare     unsafe.Pointer
	Synchronize unsafe.Pointer
}

// IOFWPseudoAddressSpace
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofwpseudoaddressspace
type IOFWPseudoAddressSpace struct {
	Reserved unsafe.Pointer
	Contains unsafe.Pointer // returns number of bytes starting at addr in this space
	Free     unsafe.Pointer
}

// IOFWPseudoAddressSpaceAux
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofwpseudoaddressspaceaux
type IOFWPseudoAddressSpaceAux struct {
	Free       unsafe.Pointer
	Init       unsafe.Pointer
	Intersects unsafe.Pointer
}

// IOFWReadCommand
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofwreadcommand
type IOFWReadCommand struct {
	Execute unsafe.Pointer
	Reinit  unsafe.Pointer
}

// IOFWReadQuadCommand - An easier to use version of IOFWReadCommand for use when the data to be transferred is an integer number of quads. Note that block read requests will be used for transfers greater than one quad unless setMaxPacket(4) is called.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofwreadquadcommand
type IOFWReadQuadCommand struct {
	Execute unsafe.Pointer
	Free    unsafe.Pointer
	Reinit  unsafe.Pointer
}

// IOFWReceiveDCL
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofwreceivedcl
type IOFWReceiveDCL struct {
	Debug unsafe.Pointer
}

// IOFWSendDCL
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofwsenddcl
type IOFWSendDCL struct {
	Debug unsafe.Pointer
	Free  unsafe.Pointer
}

// IOFWSimpleContiguousPhysicalAddressSpace
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofwsimplecontiguousphysicaladdressspace
type IOFWSimpleContiguousPhysicalAddressSpace struct {
	Free unsafe.Pointer
	Init unsafe.Pointer
}

// IOFWSimplePhysicalAddressSpace
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofwsimplephysicaladdressspace
type IOFWSimplePhysicalAddressSpace struct {
	Free unsafe.Pointer
	Init unsafe.Pointer
}

// IOFWSkipCycleDCL
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofwskipcycledcl
type IOFWSkipCycleDCL struct {
	Debug unsafe.Pointer
	Init  unsafe.Pointer
}

// IOFWSyncer
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofwsyncer
type IOFWSyncer struct {
	Create unsafe.Pointer
	Free   unsafe.Pointer
	Init   unsafe.Pointer
	Reinit unsafe.Pointer
	Signal unsafe.Pointer
	Wait   unsafe.Pointer
}

// IOFWUserObjectExporter
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofwuserobjectexporter
type IOFWUserObjectExporter struct {
	Free      unsafe.Pointer
	Init      unsafe.Pointer
	Lock      unsafe.Pointer
	Serialize unsafe.Pointer
	Unlock    unsafe.Pointer
}

// IOFWWriteCommand
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofwwritecommand
type IOFWWriteCommand struct {
	Execute unsafe.Pointer
	Free    unsafe.Pointer
	Reinit  unsafe.Pointer
}

// IOFWWriteQuadCommand - An easier to use version of IOFWWriteCommand for use when the data to be transferred is small and an integer number of quads. Note that block read requests will be used for transfers greater than one quad unless setMaxPacket(4) is called. kMaxWriteQuads is the largest legal number of quads that this object can be asked to transfer (the data is copied into an internal buffer in init() and reinit()).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofwwritequadcommand
type IOFWWriteQuadCommand struct {
	Execute unsafe.Pointer
	Free    unsafe.Pointer
	Reinit  unsafe.Pointer
}

// IOFilterInterruptEventSource - Filtering varient of the $link IOInterruptEventSource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofilterinterrupteventsource
type IOFilterInterruptEventSource struct {
	Reserved unsafe.Pointer
	Free     unsafe.Pointer
	Init     unsafe.Pointer
}

// IOFilterScheme - The common base class for all filter scheme objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofilterscheme
type IOFilterScheme struct {
	Read        unsafe.Pointer
	Synchronize unsafe.Pointer
	Unmap       unsafe.Pointer
	Write       unsafe.Pointer
}

// IOFireWireAVCAsynchronousCommand
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofirewireavcasynchronouscommand
type IOFireWireAVCAsynchronousCommand struct {
	Cancel unsafe.Pointer
	Free   unsafe.Pointer
	Init   unsafe.Pointer
	Reinit unsafe.Pointer
	Submit unsafe.Pointer
}

// IOFireWireAVCCommand
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofirewireavccommand
type IOFireWireAVCCommand struct {
	Complete unsafe.Pointer
	Execute  unsafe.Pointer
	Free     unsafe.Pointer
	Init     unsafe.Pointer
	Reinit   unsafe.Pointer
	Submit   unsafe.Pointer
}

// IOFireWireAVCNub - nub for AVC devices
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofirewireavcnub
type IOFireWireAVCNub struct {
}

// IOFireWireAVCSubUnit - nub for sub unit of AVC devices. Just for matching, calls the AVC unit for all functions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofirewireavcsubunit
type IOFireWireAVCSubUnit struct {
	Reserved unsafe.Pointer
	Init     unsafe.Pointer
	Message  unsafe.Pointer
}

// IOFireWireAVCTargetSpace - object to centralize the AVC Target mode support
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofirewireavctargetspace
type IOFireWireAVCTargetSpace struct {
	Reserved unsafe.Pointer
	Init     unsafe.Pointer // initializes the IOFireWireAVCTargetSpace object

}

// IOFireWireAVCUnit - nub for AVC devices
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofirewireavcunit
type IOFireWireAVCUnit struct {
	Available unsafe.Pointer
	Free      unsafe.Pointer
	Message   unsafe.Pointer
	Start     unsafe.Pointer
}

// IOFireWireBus - IOFireWireBus is a public class the provides access to general FireWire functionality...
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofirewirebus
type IOFireWireBus struct {
}

// IOFireWireBusAux
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofirewirebusaux
type IOFireWireBusAux struct {
}

// IOFireWireController
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofirewirecontroller
type IOFireWireController struct {
	Finalize unsafe.Pointer
	Free     unsafe.Pointer
	Init     unsafe.Pointer
	Start    unsafe.Pointer
	Stop     unsafe.Pointer
}

// IOFireWireControllerAux
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofirewirecontrolleraux
type IOFireWireControllerAux struct {
	Free unsafe.Pointer
	Init unsafe.Pointer
}

// IOFireWireDevice - Represents a FireWire device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofirewiredevice
type IOFireWireDevice struct {
	Reserved unsafe.Pointer
	Attach   unsafe.Pointer
	Finalize unsafe.Pointer
	Free     unsafe.Pointer
	Init     unsafe.Pointer // Initializes the nub.
	Message  unsafe.Pointer
}

// IOFireWireDeviceAux
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofirewiredeviceaux
type IOFireWireDeviceAux struct {
	Free unsafe.Pointer
	Init unsafe.Pointer
}

// IOFireWireDuplicateGUIDList
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofirewireduplicateguidlist
type IOFireWireDuplicateGUIDList struct {
	Create unsafe.Pointer
	Free   unsafe.Pointer
}

// IOFireWireIRMAllocation
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofirewireirmallocation
type IOFireWireIRMAllocation struct {
	Free    unsafe.Pointer
	Init    unsafe.Pointer
	Release unsafe.Pointer
}

// IOFireWireLocalNode
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofirewirelocalnode
type IOFireWireLocalNode struct {
	Attach  unsafe.Pointer
	Free    unsafe.Pointer
	Init    unsafe.Pointer
	Message unsafe.Pointer
}

// IOFireWireLocalNodeAux
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofirewirelocalnodeaux
type IOFireWireLocalNodeAux struct {
	Free unsafe.Pointer
	Init unsafe.Pointer
}

// IOFireWireMultiIsochReceiveListener
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofirewiremultiisochreceivelistener
type IOFireWireMultiIsochReceiveListener struct {
	Create unsafe.Pointer
	Free   unsafe.Pointer
	Init   unsafe.Pointer
}

// IOFireWireMultiIsochReceivePacket
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofirewiremultiisochreceivepacket
type IOFireWireMultiIsochReceivePacket struct {
	Create unsafe.Pointer
	Free   unsafe.Pointer
	Init   unsafe.Pointer
}

// IOFireWireNub
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofirewirenub
type IOFireWireNub struct {
	Free unsafe.Pointer
	Init unsafe.Pointer
}

// IOFireWireNubAux
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofirewirenubaux
type IOFireWireNubAux struct {
	Free unsafe.Pointer
	Init unsafe.Pointer
}

// IOFireWirePCRSpace - object to multiplex users of the PCR plug registers
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofirewirepcrspace
type IOFireWirePCRSpace struct {
	Reserved   unsafe.Pointer
	Activate   unsafe.Pointer
	Deactivate unsafe.Pointer
	Init       unsafe.Pointer // initializes the IOFireWirePCRSpace object

}

// IOFireWirePowerManager
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofirewirepowermanager
type IOFireWirePowerManager struct {
}

// IOFireWireSBP2LUN - Provider for most drivers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofirewiresbp2lun
type IOFireWireSBP2LUN struct {
	Attach   unsafe.Pointer // Attaches an IOService client to a provider in the registry.
	Finalize unsafe.Pointer
	Free     unsafe.Pointer
	Message  unsafe.Pointer
}

// IOFireWireSBP2Login - Supplies the login maintenance and Normal Command ORB execution portions of the API.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofirewiresbp2login
type IOFireWireSBP2Login struct {
	Free    unsafe.Pointer
	Release unsafe.Pointer // Primary implementation of the release mechanism.

}

// IOFireWireSBP2ManagementORB - Supplies non login related management ORBs. Management ORBs can be executed independent of a login, if necessary. Management ORBs are created using the IOFireWireSBP2LUN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofirewiresbp2managementorb
type IOFireWireSBP2ManagementORB struct {
	Complete unsafe.Pointer
	Execute  unsafe.Pointer
	Free     unsafe.Pointer
	Release  unsafe.Pointer // Primary implementation of the release mechanism.

}

// IOFireWireSBP2ORB - Represents an SBP2 normal command ORB. Supplies the APIs for configuring normal command ORBs. This includes setting the command block and writing the page tables for I/O. The ORBs are executed using the submitORB method in IOFireWireSBP2Login.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofirewiresbp2orb
type IOFireWireSBP2ORB struct {
	Free    unsafe.Pointer
	Release unsafe.Pointer // Primary implementation of the release mechanism.

}

// IOFireWireSBP2Target - Serves as bridge between IOFireWireUnit and IOFireWireLUN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofirewiresbp2target
type IOFireWireSBP2Target struct {
	Finalize unsafe.Pointer
	Free     unsafe.Pointer
	Message  unsafe.Pointer
	Start    unsafe.Pointer // During an IOService instantiation, the start method is called when the IOService has been selected to run on the provider.
	Stop     unsafe.Pointer // During an IOService termination, the stop method is called in its clients before they are detached & it is destroyed.

}

// IOFireWireSBP2UserClient
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofirewiresbp2userclient
type IOFireWireSBP2UserClient struct {
	Close   unsafe.Pointer
	Free    unsafe.Pointer
	Message unsafe.Pointer
	Open    unsafe.Pointer
	Start   unsafe.Pointer
}

// IOFireWireSerialBusProtocolTransport - SCSI Protocol Driver Family for FireWire SBP2 Devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofirewireserialbusprotocoltransport
type IOFireWireSerialBusProtocolTransport struct {
	Finalize unsafe.Pointer // See IOService for discussion.
	Free     unsafe.Pointer
	Init     unsafe.Pointer // See IOService for discussion.
	Login    unsafe.Pointer
	Message  unsafe.Pointer
	Start    unsafe.Pointer
}

// IOFireWireUnit
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofirewireunit
type IOFireWireUnit struct {
	Attach  unsafe.Pointer
	Free    unsafe.Pointer
	Init    unsafe.Pointer
	Message unsafe.Pointer
}

// IOFireWireUnitAux
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iofirewireunitaux
type IOFireWireUnitAux struct {
	Free unsafe.Pointer
	Init unsafe.Pointer
}

// IOFramebuffer - The base class for graphics devices to be made available as part of the desktop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioframebuffer
type IOFramebuffer struct {
	Reserved   unsafe.Pointer
	Attach     unsafe.Pointer
	Close      unsafe.Pointer
	Free       unsafe.Pointer
	Initialize unsafe.Pointer
	Message    unsafe.Pointer
	Open       unsafe.Pointer
	Start      unsafe.Pointer
	Stop       unsafe.Pointer
	Terminate  unsafe.Pointer
}

// IOFramebufferNotificationNotify
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioframebuffernotificationnotify
type IOFramebufferNotificationNotify struct {
	Info  unsafe.Pointer
	Event IOIndex
}

// IOGUIDPartitionScheme
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioguidpartitionscheme
type IOGUIDPartitionScheme struct {
	Free    unsafe.Pointer
	Init    unsafe.Pointer
	Message unsafe.Pointer
	Probe   unsafe.Pointer
	Scan    unsafe.Pointer
	Start   unsafe.Pointer
	Stop    unsafe.Pointer
}

// IOGatedOutputQueue - An extension of an IOBasicOutputQueue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iogatedoutputqueue
type IOGatedOutputQueue struct {
	Free   unsafe.Pointer // Frees the IOGatedOutputQueue object.
	Init   unsafe.Pointer // Initializes an IOGatedOutputQueue object.
	Output unsafe.Pointer
}

// IOGeneralMemoryDescriptor
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iogeneralmemorydescriptor
type IOGeneralMemoryDescriptor struct {
	Complete  unsafe.Pointer
	Free      unsafe.Pointer
	Prepare   unsafe.Pointer
	Serialize unsafe.Pointer
}

// IOGraphicsDevice
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iographicsdevice
type IOGraphicsDevice struct {
}

// IOGuardPageMemoryDescriptor
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioguardpagememorydescriptor
type IOGuardPageMemoryDescriptor struct {
	Free unsafe.Pointer
}

// IOHIDElement
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iohidelement
type IOHIDElement struct {
}

// IOHIDEventDriver
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iohideventdriver
type IOHIDEventDriver struct {
	Free unsafe.Pointer
	Init unsafe.Pointer
}

// IOHIDEventService - IOService represents an device or OS service in IOKit and DriverKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iohideventservice
type IOHIDEventService struct {
	Free    unsafe.Pointer
	Close   unsafe.Pointer
	Init    unsafe.Pointer
	Message unsafe.Pointer
	Open    unsafe.Pointer
	Start   unsafe.Pointer
	Stop    unsafe.Pointer
}

// IOHIDInterface - IOService represents an device or OS service in IOKit and DriverKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iohidinterface
type IOHIDInterface struct {
	Free    unsafe.Pointer // Free the IOHIDInterface object.
	Close   unsafe.Pointer
	Init    unsafe.Pointer // Initialize an IOHIDInterface object.
	Message unsafe.Pointer
	Open    unsafe.Pointer
	Start   unsafe.Pointer // Start up the driver using the given provider.
	Stop    unsafe.Pointer
}

// IOHIDSystem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iohidsystem
type IOHIDSystem struct {
	Attach   unsafe.Pointer
	Detach   unsafe.Pointer
	Free     unsafe.Pointer
	Init     unsafe.Pointer
	Instance unsafe.Pointer
	Message  unsafe.Pointer
	Probe    unsafe.Pointer
	Start    unsafe.Pointer
}

// IOHIDTranslationService
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iohidtranslationservice
type IOHIDTranslationService struct {
	Finalize unsafe.Pointer
	Free     unsafe.Pointer
	Open     unsafe.Pointer
	Start    unsafe.Pointer
	Stop     unsafe.Pointer
}

// IOHIDWorkLoop
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iohidworkloop
type IOHIDWorkLoop struct {
	Free unsafe.Pointer
	Init unsafe.Pointer
}

// IOHIDevice
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iohidevice
type IOHIDevice struct {
	Free  unsafe.Pointer
	Init  unsafe.Pointer
	Open  unsafe.Pointer
	Start unsafe.Pointer
}

// IOHIKeyboard
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iohikeyboard
type IOHIKeyboard struct {
	Close   unsafe.Pointer
	Free    unsafe.Pointer
	Init    unsafe.Pointer
	Message unsafe.Pointer
	Open    unsafe.Pointer
	Start   unsafe.Pointer
	Stop    unsafe.Pointer
}

// IOHIKeyboardMapper
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iohikeyboardmapper
type IOHIKeyboardMapper struct {
	Free      unsafe.Pointer
	Init      unsafe.Pointer
	Mapping   unsafe.Pointer
	Message   unsafe.Pointer
	Serialize unsafe.Pointer
}

// IOHIPointing
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iohipointing
type IOHIPointing struct {
	Close      unsafe.Pointer
	Free       unsafe.Pointer
	Init       unsafe.Pointer
	Message    unsafe.Pointer
	Open       unsafe.Pointer
	Resolution unsafe.Pointer
	Start      unsafe.Pointer
}

// IOInterleavedMemoryDescriptor - The IOInterleavedMemoryDescriptor object describes a memory area made up of portions of several other IOMemoryDescriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iointerleavedmemorydescriptor
type IOInterleavedMemoryDescriptor struct {
	Complete unsafe.Pointer // Complete processing of the memory after an I/O transfer finishes.
	Free     unsafe.Pointer
	Prepare  unsafe.Pointer // Prepare the memory for an I/O transfer.

}

// IOInterruptController
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iointerruptcontroller
type IOInterruptController struct {
}

// IOInterruptDispatchSource
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iointerruptdispatchsource
type IOInterruptDispatchSource struct {
	Free unsafe.Pointer // Performs any final cleanup for the dispatch source.
	Init unsafe.Pointer // Handles the basic initialization of the dispatch source.

}

// IOInterruptDispatchSourceInterface
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iointerruptdispatchsourceinterface
type IOInterruptDispatchSourceInterface struct {
}

// IOInterruptDispatchSourcePayload
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iointerruptdispatchsourcepayload
type IOInterruptDispatchSourcePayload struct {
	Count unsafe.Pointer
	Time  unsafe.Pointer
}

// IOInterruptEventSource - Event source for interrupt delivery to work-loop based drivers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iointerrupteventsource
type IOInterruptEventSource struct {
	Reserved unsafe.Pointer
	Provider unsafe.Pointer
	Disable  unsafe.Pointer // Disable event source.
	Enable   unsafe.Pointer // Enable event source.
	Free     unsafe.Pointer // Sub-class implementation of free method, disconnects from the interrupt source.
	Init     unsafe.Pointer // Primary initialiser for the IOInterruptEventSource class.

}

// IOKernelDebugger - Kernel debugger nub.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iokerneldebugger
type IOKernelDebugger struct {
	Debugger unsafe.Pointer // Factory method that performs allocation and initialization of an IOKernelDebugger object.
	Free     unsafe.Pointer // Frees the IOKernelDebugger instance.
	Init     unsafe.Pointer // Initializes an IOKernelDebugger instance.
	Lock     unsafe.Pointer // Takes the debugger lock conditionally.
	Message  unsafe.Pointer
	Unlock   unsafe.Pointer // Releases the debugger lock.

}

// IOKitDiagnostics
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iokitdiagnostics
type IOKitDiagnostics struct {
	Diagnostics unsafe.Pointer
	Serialize   unsafe.Pointer
}

// IOLittleMemoryCursor - An IOMemoryCursor subclass that outputs a vector of PhysicalSegments in the little endian byte order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iolittlememorycursor
type IOLittleMemoryCursor struct {
}

// IOLocalConfigDirectory
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iolocalconfigdirectory
type IOLocalConfigDirectory struct {
	Reserved unsafe.Pointer
	Compile  unsafe.Pointer
	Create   unsafe.Pointer
	Free     unsafe.Pointer
	Init     unsafe.Pointer
	Update   unsafe.Pointer
}

// IOMapper
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iomapper
type IOMapper struct {
	Free  unsafe.Pointer
	Start unsafe.Pointer
}

// IOMedia - A random-access disk device abstraction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iomedia
type IOMedia struct {
	Free        unsafe.Pointer
	Init        unsafe.Pointer
	Read        unsafe.Pointer
	Synchronize unsafe.Pointer
	Unmap       unsafe.Pointer
	Write       unsafe.Pointer
}

// IOMediaBSDClient
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iomediabsdclient
type IOMediaBSDClient struct {
	Free      unsafe.Pointer
	Init      unsafe.Pointer
	Ioctl     unsafe.Pointer
	Start     unsafe.Pointer
	Terminate unsafe.Pointer
}

// IOMemoryCursor - A mechanism to convert memory references to physical addresses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iomemorycursor
type IOMemoryCursor struct {
}

// IOMemoryDescriptor - An abstract base class that defines common methods for describing physical or virtual memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iomemorydescriptor
type IOMemoryDescriptor struct {
	Free       unsafe.Pointer // Performs any final cleanup for the memory descriptor object.
	Reserved   unsafe.Pointer
	Map        unsafe.Pointer // Maps an IOMemoryDescriptor into the kernel map.
	Complete   unsafe.Pointer // Complete processing of the memory after an I/O transfer finishes.
	Initialize unsafe.Pointer
	Prepare    unsafe.Pointer // Prepare the memory for an I/O transfer.
	Redirect   unsafe.Pointer
}

// IOMemoryMap - A class defining common methods for describing a memory mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iomemorymap
type IOMemoryMap struct {
	Free     unsafe.Pointer // Performs any final cleanup for the memory map object.
	Redirect unsafe.Pointer // Replace the memory mapped in a process with new backing memory.
	Unmap    unsafe.Pointer // Force the IOMemoryMap to unmap, without destroying the object.

}

// IOMultiMemoryDescriptor - The IOMultiMemoryDescriptor object describes a memory area made up of several other IOMemoryDescriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iomultimemorydescriptor
type IOMultiMemoryDescriptor struct {
	Complete unsafe.Pointer // Complete processing of the memory after an I/O transfer finishes.
	Free     unsafe.Pointer
	Prepare  unsafe.Pointer // Prepare the memory for an I/O transfer.

}

// IONDRVControlParameters
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iondrvcontrolparameters
type IONDRVControlParameters struct {
	Code   unsafe.Pointer
	Params unsafe.Pointer
}

// IONDRVFramebuffer
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iondrvframebuffer
type IONDRVFramebuffer struct {
	Free  unsafe.Pointer
	Probe unsafe.Pointer
	Start unsafe.Pointer
	Stop  unsafe.Pointer
}

// IONVRAMController
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ionvramcontroller
type IONVRAMController struct {
	Read   unsafe.Pointer
	Select unsafe.Pointer
	Sync   unsafe.Pointer
	Write  unsafe.Pointer
}

// IONVRAMDescriptor
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ionvramdescriptor
type IONVRAMDescriptor struct {
	Format        unsafe.Pointer
	DeviceNum     unsafe.Pointer
	BusNum        unsafe.Pointer
	BridgeDevices unsafe.Pointer
	FunctionNum   unsafe.Pointer
	BridgeCount   unsafe.Pointer
	Marker        unsafe.Pointer
}

// IONaturalMemoryCursor - An IOMemoryCursor subclass that outputs a vector of PhysicalSegments in the natural byte orientation for the CPU.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ionaturalmemorycursor
type IONaturalMemoryCursor struct {
}

// IONetworkController - Implements the framework for a generic network controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ionetworkcontroller
type IONetworkController struct {
	Disable unsafe.Pointer
	Enable  unsafe.Pointer
	Free    unsafe.Pointer // Frees the IONetworkController object.
	Init    unsafe.Pointer // Initializes the IONetworkController object.
	Message unsafe.Pointer // Receives messages delivered from an attached provider.
	Prepare unsafe.Pointer // Prepares the controller before an IOService is created and attached as a client.
	Start   unsafe.Pointer // Starts the network controller.
	Stop    unsafe.Pointer // Stops the network controller.

}

// IONetworkData - An object that manages a fixed-size named buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ionetworkdata
type IONetworkData struct {
	Free      unsafe.Pointer // Frees the IONetworkData object.
	Init      unsafe.Pointer // Initializes an IONetworkData object.
	Read      unsafe.Pointer // An access method that reads from the data buffer.
	Reset     unsafe.Pointer // An access method that resets the data buffer.
	Serialize unsafe.Pointer // Serializes the IONetworkData object.
	Write     unsafe.Pointer // An access method that writes to the data buffer.

}

// IONetworkInterface - Abstract class that manages the connection between an IONetworkController and the data link interface layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ionetworkinterface
type IONetworkInterface struct {
	Free                unsafe.Pointer // Frees the
	If_detach           unsafe.Pointer
	If_input_poll       unsafe.Pointer
	If_input_poll_gated unsafe.Pointer
	If_ioctl            unsafe.Pointer
	If_output           unsafe.Pointer
	If_set_bpf_tap      unsafe.Pointer
	If_start            unsafe.Pointer
	If_start_gated      unsafe.Pointer
	If_start_precheck   unsafe.Pointer
	Init                unsafe.Pointer // Initializes the
	Lock                unsafe.Pointer // Acquires a recursive lock owned by the interface.
	Message             unsafe.Pointer
	Unlock              unsafe.Pointer // Releases the recursive lock owned by the interface.

}

// IONetworkMedium - An object that encapsulates information about a network medium (i.e. 10Base-T, or 100Base-T Full Duplex).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ionetworkmedium
type IONetworkMedium struct {
	Free      unsafe.Pointer // Frees the IONetworkMedium object.
	Init      unsafe.Pointer // Initializes an IONetworkMedium object.
	Medium    unsafe.Pointer // Factory method that allocates and initializes an IONetworkMedium object.
	Serialize unsafe.Pointer // Serializes the IONetworkMedium object.

}

// IONotifier - An abstract base class defining common methods for controlling a notification request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ionotifier
type IONotifier struct {
	Disable unsafe.Pointer // Disables the notification request.
	Enable  unsafe.Pointer // Sets the enable state of the notification request.
	Remove  unsafe.Pointer // Removes the notification request and releases it.

}

// IOOutputQueue - A packet queue that supports multiple producers and a single consumer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iooutputqueue
type IOOutputQueue struct {
	Enqueue unsafe.Pointer // Adds a packet, or a chain of packets, to the queue.
	Flush   unsafe.Pointer // Drops and frees all packets currently held by the queue.
	Free    unsafe.Pointer // Frees the IOOutputQueue object.
	Init    unsafe.Pointer // Initializes an IOOutputQueue object.
	Service unsafe.Pointer // Services the queue.
	Start   unsafe.Pointer // Starts up the queue.
	Stop    unsafe.Pointer // Stops the queue.

}

// IOPCIDevice - An IOService class representing a PCI device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iopcidevice
type IOPCIDevice struct {
	Reserved unsafe.Pointer
	Attach   unsafe.Pointer
	Detach   unsafe.Pointer
	Flr      unsafe.Pointer
	Free     unsafe.Pointer // Performs any final cleanup for the object.
	Init     unsafe.Pointer
	Relocate unsafe.Pointer
	Reset    unsafe.Pointer
}

// IOPCIEvent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iopcievent
type IOPCIEvent struct {
	Event    unsafe.Pointer
	Reporter *IOPCIDevice
	Data     unsafe.Pointer
}

// IOPCIEventSource
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iopcieventsource
type IOPCIEventSource struct {
	Disable unsafe.Pointer
	Enable  unsafe.Pointer
	Free    unsafe.Pointer
}

// IOPCIPhysicalAddress
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iopciphysicaladdress
type IOPCIPhysicalAddress struct {
	PhysMid  unsafe.Pointer
	PhysHi   unsafe.Pointer
	LengthLo unsafe.Pointer
	LengthHi unsafe.Pointer
	PhysLo   unsafe.Pointer
}

// IOPMPowerSource
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iopmpowersource
type IOPMPowerSource struct {
	Properties   unsafe.Pointer
	Amperage     unsafe.Pointer
	Free         unsafe.Pointer
	Init         unsafe.Pointer
	Location     unsafe.Pointer
	Manufacturer unsafe.Pointer
	Model        unsafe.Pointer
	Serial       unsafe.Pointer
	Voltage      unsafe.Pointer
}

// IOPMPowerSourceList
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iopmpowersourcelist
type IOPMPowerSourceList struct {
	Free       unsafe.Pointer
	Initialize unsafe.Pointer
}

// IOPMrootDomain
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iopmrootdomain
type IOPMrootDomain struct {
	Construct unsafe.Pointer
	Start     unsafe.Pointer
}

// IOPacketQueue - Implements a bounded FIFO queue of mbuf packets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iopacketqueue
type IOPacketQueue struct {
	Dequeue unsafe.Pointer // Removes a single packet from the head of the queue.
	Enqueue unsafe.Pointer
	Flush   unsafe.Pointer // Frees all packets currently held in the queue and releases them back to the free mbuf pool.
	Free    unsafe.Pointer // Frees the IOPacketQueue object.
	Peek    unsafe.Pointer // Examines the packet at the head of the queue without removing it from the queue.
	Prepend unsafe.Pointer
}

// IOPartitionScheme - The common base class for all partition scheme objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iopartitionscheme
type IOPartitionScheme struct {
	Free        unsafe.Pointer
	Init        unsafe.Pointer
	Read        unsafe.Pointer
	Synchronize unsafe.Pointer
	Unmap       unsafe.Pointer
	Write       unsafe.Pointer
}

// IOPlatformDevice
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioplatformdevice
type IOPlatformDevice struct {
}

// IOPlatformExpert
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioplatformexpert
type IOPlatformExpert struct {
	Attach    unsafe.Pointer
	Configure unsafe.Pointer
	Start     unsafe.Pointer
}

// IOPlatformExpertDevice
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioplatformexpertdevice
type IOPlatformExpertDevice struct {
	Free unsafe.Pointer
	Init unsafe.Pointer
}

// IOPlatformIO - The base class for platform I/O drivers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioplatformio
type IOPlatformIO struct {
	Start unsafe.Pointer
}

// IOPowerConnection
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iopowerconnection
type IOPowerConnection struct {
}

// IOPwrController
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iopwrcontroller
type IOPwrController struct {
}

// IORPCMessageErrorReturn
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iorpcmessageerrorreturn
type IORPCMessageErrorReturn struct {
	Content IORPCMessageErrorReturnContent
	Mach    IORPCMessageMach
}

// IORTC
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iortc
type IORTC struct {
}

// IORTCController
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iortccontroller
type IORTCController struct {
}

// IORangeAllocator - A utility class to manage allocations from a range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iorangeallocator
type IORangeAllocator struct {
	Allocate   unsafe.Pointer // Allocates from the free list, at any offset.
	Deallocate unsafe.Pointer // Deallocates a range to the free list.
	Free       unsafe.Pointer
	Init       unsafe.Pointer // Standard initializer for IORangeAllocator.
	Serialize  unsafe.Pointer
}

// IOReducedBlockServices
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioreducedblockservices
type IOReducedBlockServices struct {
	Attach  unsafe.Pointer
	Detach  unsafe.Pointer
	Free    unsafe.Pointer
	Message unsafe.Pointer
}

// IORegistryEntry - The base class for all objects in the registry.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioregistryentry
type IORegistryEntry struct {
	Reserved   unsafe.Pointer
	Free       unsafe.Pointer // Standard free method for all IORegistryEntry subclasses.
	Init       unsafe.Pointer
	Initialize unsafe.Pointer
}

// IORegistryIterator - An iterator over the registry.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioregistryiterator
type IORegistryIterator struct {
	Free  unsafe.Pointer
	Reset unsafe.Pointer // Exits all levels of recursion, restoring the iterator to its state at creation.

}

// IOSCSIBlockCommandsDevice
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioscsiblockcommandsdevice
type IOSCSIBlockCommandsDevice struct {
	Free    unsafe.Pointer
	Message unsafe.Pointer
}

// IOSCSILogicalUnitNub
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioscsilogicalunitnub
type IOSCSILogicalUnitNub struct {
}

// IOSCSIMultimediaCommandsDevice
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioscsimultimediacommandsdevice
type IOSCSIMultimediaCommandsDevice struct {
	Free unsafe.Pointer
}

// IOSCSIParallelInterfaceController - Class that represents a SCSI Host Bus Adapter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioscsiparallelinterfacecontroller
type IOSCSIParallelInterfaceController struct {
	Free  unsafe.Pointer
	Start unsafe.Pointer
	Stop  unsafe.Pointer
}

// IOSCSIPeripheralDeviceNub
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioscsiperipheraldevicenub
type IOSCSIPeripheralDeviceNub struct {
	Free    unsafe.Pointer
	Init    unsafe.Pointer
	Message unsafe.Pointer
	Start   unsafe.Pointer
}

// IOSCSIPeripheralDeviceType00
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioscsiperipheraldevicetype00
type IOSCSIPeripheralDeviceType00 struct {
	Free  unsafe.Pointer
	Init  unsafe.Pointer
	Start unsafe.Pointer
}

// IOSCSIPeripheralDeviceType05
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioscsiperipheraldevicetype05
type IOSCSIPeripheralDeviceType05 struct {
	Init  unsafe.Pointer
	Start unsafe.Pointer
	Stop  unsafe.Pointer
}

// IOSCSIPeripheralDeviceType07
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioscsiperipheraldevicetype07
type IOSCSIPeripheralDeviceType07 struct {
	Init  unsafe.Pointer
	Start unsafe.Pointer
	Stop  unsafe.Pointer
}

// IOSCSIPeripheralDeviceType0E
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioscsiperipheraldevicetype0e
type IOSCSIPeripheralDeviceType0E struct {
	Init  unsafe.Pointer
	Start unsafe.Pointer
	Stop  unsafe.Pointer
}

// IOSCSIPrimaryCommandsDevice
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioscsiprimarycommandsdevice
type IOSCSIPrimaryCommandsDevice struct {
	Free    unsafe.Pointer
	Init    unsafe.Pointer
	Message unsafe.Pointer
	Start   unsafe.Pointer
	Stop    unsafe.Pointer
}

// IOSCSIProtocolInterface - This class defines the public SCSI Protocol Layer API for any class that provides Protocol services or needs to provide the Protocol Service API for passing service requests to a Protocol Service driver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioscsiprotocolinterface
type IOSCSIProtocolInterface struct {
	Finalize unsafe.Pointer // Finalizes the destruction of an IOService object.
	Free     unsafe.Pointer // Called to release all resources held by the object.
	Init     unsafe.Pointer
	Start    unsafe.Pointer // During an IOService object's instantiation, starts the IOService object that has been selected to run on the provider.

}

// IOSCSIProtocolServices - This class defines the public SCSI Protocol Services Layer API for any class that implements SCSI protocol services. A protocol services layer driver is responsible for taking incoming SCSITaskIdentifier objects and translating them to the native command type for the native protocol interface (e.g. SBP-2 ORB on FireWire).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioscsiprotocolservices
type IOSCSIProtocolServices struct {
	Free  unsafe.Pointer // Frees data structures that were allocated during start().
	Init  unsafe.Pointer // Standard init method for all IORegistryEntry subclasses.
	Start unsafe.Pointer // During an IOService object's instantiation, starts the IOService object that has been selected to run on the provider.

}

// IOSCSIReducedBlockCommandsDevice
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioscsireducedblockcommandsdevice
type IOSCSIReducedBlockCommandsDevice struct {
	Free unsafe.Pointer
}

// IOService - The base class for most I/O Kit families, devices, and drivers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioservice
type IOService struct {
	Reserved         unsafe.Pointer
	Start            unsafe.Pointer // During an IOService object's instantiation, starts the IOService object that has been selected to run on the provider.
	Stop             unsafe.Pointer // During an IOService termination, the stop method is called in its clients before they are detached & it is destroyed.
	Free             unsafe.Pointer // Frees data structures that were allocated when power management was initialized on this service.
	Init             unsafe.Pointer
	Attach           unsafe.Pointer // Attaches an IOService client to a provider in the I/O Registry.
	Close            unsafe.Pointer // Releases active access to a provider.
	Command_received unsafe.Pointer
	Detach           unsafe.Pointer // Detaches an IOService client from a provider in the I/O Registry.
	Finalize         unsafe.Pointer // Finalizes the destruction of an IOService object.
	Message          unsafe.Pointer // Receives a generic message delivered from an attached provider.
	Open             unsafe.Pointer // Requests active access to a provider.
	Probe            unsafe.Pointer // During an IOService object's instantiation, probes a matched service to see if it can be used.
	Terminate        unsafe.Pointer // Makes an IOService object inactive and begins its destruction.

}

// IOServiceInterestContent64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioserviceinterestcontent64
type IOServiceInterestContent64 struct {
	MessageType     Natural_t
	MessageArgument Io_user_reference_t
}

// IOServiceNotificationDispatchSource
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioservicenotificationdispatchsource
type IOServiceNotificationDispatchSource struct {
	Free unsafe.Pointer
	Init unsafe.Pointer
}

// IOServiceNotificationDispatchSourceInterface
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioservicenotificationdispatchsourceinterface
type IOServiceNotificationDispatchSourceInterface struct {
}

// IOServiceStateNotificationDispatchSource
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOServiceStateNotificationDispatchSource
type IOServiceStateNotificationDispatchSource struct {
	Free unsafe.Pointer
	Init unsafe.Pointer
}

// IOServiceStateNotificationDispatchSourceInterface
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioservicestatenotificationdispatchsourceinterface
type IOServiceStateNotificationDispatchSourceInterface struct {
}

// IOSharedDataQueue - A generic queue designed to pass data both from the kernel to a user process and from a user process to the kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioshareddataqueue
type IOSharedDataQueue struct {
	Dequeue unsafe.Pointer // Dequeues the next available entry on the queue and copies it into the given data pointer.
	Enqueue unsafe.Pointer
	Free    unsafe.Pointer
	Peek    unsafe.Pointer // Used to peek at the next entry on the queue.

}

// IOSharedInterruptController
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iosharedinterruptcontroller
type IOSharedInterruptController struct {
}

// IOStorage - The common base class for mass storage objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iostorage
type IOStorage struct {
	Attach      unsafe.Pointer
	Complete    unsafe.Pointer
	Discard     unsafe.Pointer
	Open        unsafe.Pointer
	Read        unsafe.Pointer
	Synchronize unsafe.Pointer
	Unmap       unsafe.Pointer
	Write       unsafe.Pointer
}

// IOStorageAttributes
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iostorageattributes
type IOStorageAttributes struct {
	Reserved0064   uint64
	Priority       IOStoragePriority
	Bufattr        Bufattr_t        // Reserved for future use. Set to zero.
	Options        IOStorageOptions // Options for the request. See IOStorageOptions.
	Reserved0032   unsafe.Pointer
	Reserved0024   unsafe.Pointer
	AdjustedOffset uint64
}

// IOStorageCompletion
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iostoragecompletion
type IOStorageCompletion struct {
	Parameter unsafe.Pointer // Opaque client-supplied pointer.
	Target    unsafe.Pointer // Opaque client-supplied pointer (or an instance pointer for a C++ callback).
	Action    unsafe.Pointer // Completion routine to call on completion of the data transfer.

}

// IOStorageExtent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iostorageextent
type IOStorageExtent struct {
	ByteStart uint64 // Starting byte offset for the operation.
	ByteCount uint64 // Size of the operation.

}

// IOStorageProvisionExtent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iostorageprovisionextent
type IOStorageProvisionExtent struct {
	Reserved      unsafe.Pointer
	ByteCount     uint64
	ProvisionType unsafe.Pointer
	ByteStart     uint64
}

// IOStream - A class representing a stream of data buffers passed from kernel to user space and back again.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iostream
type IOStream struct {
	Free unsafe.Pointer
	Init unsafe.Pointer
}

// IOStreamBuffer - A class representing a data buffer that is part of an IOStream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iostreambuffer
type IOStreamBuffer struct {
	Free unsafe.Pointer
}

// IOStreamUserClient
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iostreamuserclient
type IOStreamUserClient struct {
	Free  unsafe.Pointer
	Start unsafe.Pointer
	Stop  unsafe.Pointer
}

// IOSubMemoryDescriptor - The IOSubMemoryDescriptor object describes a memory area made up of a portion of another IOMemoryDescriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iosubmemorydescriptor
type IOSubMemoryDescriptor struct {
	Complete unsafe.Pointer
	Free     unsafe.Pointer
	Prepare  unsafe.Pointer
	Redirect unsafe.Pointer
}

// IOTimeStampIntervalConstantFiltered
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iotimestampintervalconstantfiltered
type IOTimeStampIntervalConstantFiltered struct {
}

// IOTimerEventSource - Time based event source mechanism.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iotimereventsource
type IOTimerEventSource struct {
	Reserved unsafe.Pointer
	Abstime  unsafe.Pointer
	Disable  unsafe.Pointer // Disable a timed callout.
	Enable   unsafe.Pointer // Enables a call to the action.
	Free     unsafe.Pointer // Sub-class implementation of free method, frees calloutEntry
	Init     unsafe.Pointer
	Timeout  unsafe.Pointer // Function that routes the call from the OS' timeout mechanism into a work-loop context.

}

// IOTrackingCallSiteInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iotrackingcallsiteinfo
type IOTrackingCallSiteInfo struct {
	Bt         Mach_vm_address_t
	Size       Mach_vm_size_t
	Count      unsafe.Pointer
	AddressPID int32
	BtPID      int32
	Address    Mach_vm_address_t
}

// IOUSB30HubPortStatusExt - A structure that represents an extension to the USB 3.0 hub port status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iousb30hubportstatusext
type IOUSB30HubPortStatusExt struct {
	WPortStatus     uint16         // The port status field bits.
	WPortChange     uint16         // The port status change bits.
	DwExtPortStatus unsafe.Pointer // The extended port status bits.

}

// IOUSBDevice - An input/output service object that represents a device on the USB bus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iousbdevice
type IOUSBDevice struct {
	Start unsafe.Pointer
	Stop  unsafe.Pointer
}

// IOUSBHostBundledCompletion - The structure that specifies the action to perform when a bulk USB input/output request completes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iousbhostbundledcompletion
type IOUSBHostBundledCompletion struct {
	Action    unsafe.Pointer // The action that executes when the input/output request completes.
	Owner     unsafe.Pointer // A pointer to an object that owns the transfer.
	Parameter unsafe.Pointer // A context pointer within the completion action.

}

// IOUSBHostCompletion - The structure that specifies the action to perform when the USB input/output request completes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iousbhostcompletion
type IOUSBHostCompletion struct {
	Action    unsafe.Pointer // The action to run when the input/output request completes.
	Owner     unsafe.Pointer // A pointer to an object that owns the transfer.
	Parameter unsafe.Pointer // A context pointer within the completion action.

}

// IOUSBHostDevice
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iousbhostdevice
type IOUSBHostDevice struct {
	Attach    unsafe.Pointer
	Close     unsafe.Pointer
	Free      unsafe.Pointer
	Message   unsafe.Pointer
	Open      unsafe.Pointer
	Reset     unsafe.Pointer
	Start     unsafe.Pointer
	Stop      unsafe.Pointer
	Terminate unsafe.Pointer
}

// IOUSBHostIOSource
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iousbhostiosource
type IOUSBHostIOSource struct {
	Abort   unsafe.Pointer
	Close   unsafe.Pointer
	Destroy unsafe.Pointer
	Free    unsafe.Pointer
	Io      unsafe.Pointer
	Open    unsafe.Pointer
}

// IOUSBHostIOSourceClientRecord - The USB host input/output source client record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iousbhostiosourceclientrecord
type IOUSBHostIOSourceClientRecord struct {
	Link          IOUSBHostIOSourceClientRecordLink // The client record link.
	OutstandingIO unsafe.Pointer                    // The outstanding input/output value.
	ForClient     *IOService                        // The client pointer.

}

// IOUSBHostInterface
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iousbhostinterface
type IOUSBHostInterface struct {
	Attach    unsafe.Pointer
	Close     unsafe.Pointer
	Free      unsafe.Pointer
	Message   unsafe.Pointer
	Open      unsafe.Pointer
	Start     unsafe.Pointer
	Stop      unsafe.Pointer
	Terminate unsafe.Pointer
}

// IOUSBHostIsochronousCompletion - A structure describing the completion callback for an asynchronous isochronous operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iousbhostisochronouscompletion
type IOUSBHostIsochronousCompletion struct {
	Action    unsafe.Pointer // The action to run when the input/output request completes.
	Owner     unsafe.Pointer // A pointer to an object that owns the transfer.
	Parameter unsafe.Pointer // A context pointer within the completion action.

}

// IOUSBHostIsochronousTransactionCompletion
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iousbhostisochronoustransactioncompletion
type IOUSBHostIsochronousTransactionCompletion struct {
	Action    unsafe.Pointer
	Owner     unsafe.Pointer
	Parameter unsafe.Pointer
}

// IOUSBHostPipe
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iousbhostpipe
type IOUSBHostPipe struct {
	Abort unsafe.Pointer
	Free  unsafe.Pointer
	Io    unsafe.Pointer
}

// IOUSBHostStream
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iousbhoststream
type IOUSBHostStream struct {
	Abort unsafe.Pointer
	Free  unsafe.Pointer
	Io    unsafe.Pointer
}

// IOUSBInterface - An object that represents an interface of a device on the USB bus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iousbinterface
type IOUSBInterface struct {
}

// IOUSBStandardEndpointDescriptors - A container for descriptors for a single endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iousbstandardendpointdescriptors
type IOUSBStandardEndpointDescriptors struct {
	BcdUSB                 uint16                                                    // A binary-coded decimal value that indicates the USB version that the device supports.
	Descriptor             IOUSBEndpointDescriptor                                   // A valid endpoint descriptor.
	SsCompanionDescriptor  IOUSBSuperSpeedEndpointCompanionDescriptor                // The companion descriptor for SuperSpeed devices.
	SspCompanionDescriptor IOUSBSuperSpeedPlusIsochronousEndpointCompanionDescriptor // The companion descriptor for SuperSpeedPlus devices.

}

// IOUserClient - Provides a basis for communication between client applications and I/O Kit objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iouserclient
type IOUserClient struct {
	Reserved unsafe.Pointer
	Free     unsafe.Pointer
	Init     unsafe.Pointer
	Reserve  unsafe.Pointer
}

// IOUserClientMethodArguments
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iouserclientmethodarguments
type IOUserClientMethodArguments struct {
	Completion                 *OSAction
	ScalarInput                unsafe.Pointer
	ScalarInputCount           unsafe.Pointer
	ScalarOutput               unsafe.Pointer
	ScalarOutputCount          unsafe.Pointer
	Selector                   unsafe.Pointer
	StructureInput             *OSData
	StructureInputDescriptor   *IOMemoryDescriptor
	StructureOutput            *OSData
	StructureOutputDescriptor  *IOMemoryDescriptor
	StructureOutputMaximumSize unsafe.Pointer
	Version                    unsafe.Pointer
}

// IOUserClientMethodDispatch
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iouserclientmethoddispatch
type IOUserClientMethodDispatch struct {
	CheckCompletionExists    unsafe.Pointer
	CheckScalarInputCount    unsafe.Pointer
	CheckScalarOutputCount   unsafe.Pointer
	CheckStructureInputSize  unsafe.Pointer
	CheckStructureOutputSize unsafe.Pointer
	Function                 unsafe.Pointer
}

// IOVideoControlDictionary
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iovideocontroldictionary
type IOVideoControlDictionary struct {
	Create unsafe.Pointer
}

// IOVideoDevice - A class that represents a video device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iovideodevice
type IOVideoDevice struct {
	Free unsafe.Pointer
	Init unsafe.Pointer
}

// IOVideoDeviceUserClient
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iovideodeviceuserclient
type IOVideoDeviceUserClient struct {
	Close unsafe.Pointer
	Open  unsafe.Pointer
	Start unsafe.Pointer
}

// IOVideoDeviceUserClientInit
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iovideodeviceuserclientinit
type IOVideoDeviceUserClientInit struct {
	Start unsafe.Pointer
}

// IOVideoStream - A class representing a stream of video data buffers passed from kernel to user space and back again.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iovideostream
type IOVideoStream struct {
}

// IOVideoStreamDictionary
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iovideostreamdictionary
type IOVideoStreamDictionary struct {
	Create unsafe.Pointer
}

// IOVideoStreamFormatDictionary
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iovideostreamformatdictionary
type IOVideoStreamFormatDictionary struct {
	Create unsafe.Pointer
}

// IOWatchDogTimer
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iowatchdogtimer
type IOWatchDogTimer struct {
	Start unsafe.Pointer
	Stop  unsafe.Pointer
}

// IOWorkGroup
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioworkgroup
type IOWorkGroup struct {
	Free unsafe.Pointer
	Init unsafe.Pointer
}

// IOWorkGroupInterface
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioworkgroupinterface
type IOWorkGroupInterface struct {
}

// IOWorkLoop
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ioworkloop
type IOWorkLoop struct {
	Reserved unsafe.Pointer
	Free     unsafe.Pointer
	Init     unsafe.Pointer
}

// KeyAttribute
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/keyattribute
type KeyAttribute struct {
}

// KeyValueMask
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/keyvaluemask
type KeyValueMask struct {
	Key  unsafe.Pointer
	Mask unsafe.Pointer
}

// LittleSInt16
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/littlesint16
type LittleSInt16 struct {
	Get unsafe.Pointer
}

// LittleSInt32
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/littlesint32
type LittleSInt32 struct {
	Get unsafe.Pointer
}

// LittleSInt64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/littlesint64
type LittleSInt64 struct {
	Get unsafe.Pointer
}

// LittleUInt16
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/littleuint16
type LittleUInt16 struct {
	Get unsafe.Pointer
}

// LittleUInt32
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/littleuint32
type LittleUInt32 struct {
	Get unsafe.Pointer
}

// LittleUInt64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/littleuint64
type LittleUInt64 struct {
	Get unsafe.Pointer
}

// OSAction
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osaction
type OSAction struct {
	Free unsafe.Pointer // Performs any final cleanup for the action object.

}

// OSActionInterface
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osactioninterface
type OSActionInterface struct {
}

// OSAction_IOHIDEventService__CopyEvent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osaction_iohideventservice_copyevent
type OSAction_IOHIDEventService__CopyEvent struct {
}

// OSAction_IOHIDEventService__CopyEventInterface
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osaction_iohideventservice_copyeventinterface
type OSAction_IOHIDEventService__CopyEventInterface struct {
}

// OSAction_IOHIDEventService__SetLED
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osaction_iohideventservice_setled
type OSAction_IOHIDEventService__SetLED struct {
}

// OSAction_IOHIDEventService__SetLEDInterface
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osaction_iohideventservice_setledinterface
type OSAction_IOHIDEventService__SetLEDInterface struct {
}

// OSAction_IOHIDEventService__SetUserProperties
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osaction_iohideventservice_setuserproperties
type OSAction_IOHIDEventService__SetUserProperties struct {
}

// OSAction_IOHIDEventService__SetUserPropertiesInterface
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osaction_iohideventservice_setuserpropertiesinterface
type OSAction_IOHIDEventService__SetUserPropertiesInterface struct {
}

// OSAction_IOUserClient_KernelCompletion
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osaction_iouserclient_kernelcompletion
type OSAction_IOUserClient_KernelCompletion struct {
}

// OSAction_IOUserClient_KernelCompletionInterface
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osaction_iouserclient_kernelcompletioninterface
type OSAction_IOUserClient_KernelCompletionInterface struct {
}

// OSArray - OSArray provides an indexed store of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osarray
type OSArray struct {
	Free      unsafe.Pointer // Deallocates or releases any resources used by the OSArray instance.
	Merge     unsafe.Pointer // Appends the contents of an array onto the receiving array.
	Serialize unsafe.Pointer // Archives the receiver into the provided OSSerialize object.

}

// OSBoolean - OSBoolean wraps a boolean value in a C++ object for use in Libkern collections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osboolean
type OSBoolean struct {
	Free       unsafe.Pointer // Overridden to prevent deallocation of the shared global instances.
	Initialize unsafe.Pointer
	Serialize  unsafe.Pointer // Archives the receiver into the provided OSSerialize object.

}

// OSClassDescription
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osclassdescription
type OSClassDescription struct {
	DescriptionSize         unsafe.Pointer
	DispatchNames           unsafe.Pointer
	Flags                   unsafe.Pointer
	MetaMethodNames         unsafe.Pointer
	MetaMethodNamesOffset   unsafe.Pointer
	MetaMethodNamesSize     unsafe.Pointer
	MetaMethodOptions       unsafe.Pointer
	MetaMethodOptionsOffset unsafe.Pointer
	MetaMethodOptionsSize   unsafe.Pointer
	MethodNames             unsafe.Pointer
	MethodNamesOffset       unsafe.Pointer
	MethodNamesSize         unsafe.Pointer
	MethodOptions           unsafe.Pointer
	MethodOptionsOffset     unsafe.Pointer
	MethodOptionsSize       unsafe.Pointer
	Name                    unsafe.Pointer
	QueueNamesOffset        unsafe.Pointer
	QueueNamesSize          unsafe.Pointer
	Resv1                   unsafe.Pointer
	SuperName               unsafe.Pointer
}

// OSCollection - The abstract superclass for Libkern collections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/oscollection
type OSCollection struct {
	Init unsafe.Pointer // Initializes the OSCollection object.

}

// OSCollectionIterator
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/oscollectioniterator
type OSCollectionIterator struct {
	Free  unsafe.Pointer // Releases or deallocates any resources used by the OSCollectionIterator object.
	Reset unsafe.Pointer // Resets the iterator to the beginning of the collection, as if it had just been created.

}

// OSData - OSData wraps an array of bytes in a C++ object for use in Libkern collections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osdata
type OSData struct {
	Free      unsafe.Pointer // Deallocates or releases any resources used by the OSDictionary instance.
	Serialize unsafe.Pointer // Archives the receiver into the provided OSSerialize object.

}

// OSDictionary - OSDictionary provides an associative store using strings for keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osdictionary
type OSDictionary struct {
	Free      unsafe.Pointer // Deallocates or releases any resources used by the OSDictionary instance.
	Merge     unsafe.Pointer // Merges the contents of a dictionary into the receiver.
	Serialize unsafe.Pointer // Archives the receiver into the provided OSSerialize object.

}

// OSIterator - The abstract superclass for Libkern iterators.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ositerator
type OSIterator struct {
	Reset unsafe.Pointer // Resets the iterator to the beginning of the collection, as if it had just been created.

}

// OSMetaClass
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osmetaclass
type OSMetaClass struct {
	Alloc     unsafe.Pointer
	Release   unsafe.Pointer
	Retain    unsafe.Pointer
	Serialize unsafe.Pointer
}

// OSMetaClassBase - OSMetaClassBase is the abstract bootstrap class for the Libkern and I/O Kit run-time type information system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osmetaclassbase
type OSMetaClassBase struct {
	Initialize unsafe.Pointer
	Release    unsafe.Pointer
	Retain     unsafe.Pointer // Abstract declaration of retain().
	Serialize  unsafe.Pointer // Abstract declaration of serialize.

}

// OSNotificationHeader64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osnotificationheader64
type OSNotificationHeader64 struct {
	Size      uint32
	Type      Natural_t
	Reference unsafe.Pointer
	Content   unsafe.Pointer
}

// OSNumber - OSNumber wraps an integer value in a C++ object for use in Libkern collections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osnumber
type OSNumber struct {
	Free      unsafe.Pointer // Deallocates or releases any resources used by the OSNumber instance.
	Init      unsafe.Pointer
	Serialize unsafe.Pointer // Archives the receiver into the provided OSSerialize object.

}

// OSObject - OSObject is the concrete root class of the Libkern and I/O Kit C++ class hierarchy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osobject
type OSObject struct {
	Free      unsafe.Pointer // Deallocates/releases resources held by the object.
	Init      unsafe.Pointer // Initializes a newly-allocated object.
	Release   unsafe.Pointer
	Retain    unsafe.Pointer // Retains a reference to the object.
	Serialize unsafe.Pointer // Overridden by subclasses to archive the receiver into the provided OSSerialize object.

}

// OSOrderedSet - OSOrderedSet provides an ordered set store of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osorderedset
type OSOrderedSet struct {
	Free   unsafe.Pointer // Deallocatesand releases any resources used by the OSOrderedSet instance.
	Init   unsafe.Pointer
	Member unsafe.Pointer // Checks the ordered set for the presence of an object.

}

// OSSerialize - OSSerialize coordinates serialization of Libkern C++ objects into an XML stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osserialize
type OSSerialize struct {
	Free unsafe.Pointer
	Text unsafe.Pointer // Returns the XML text serialized so far.

}

// OSSerializer
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osserializer
type OSSerializer struct {
	Free      unsafe.Pointer
	Serialize unsafe.Pointer
}

// OSSet - OSSet provides an unordered set store of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osset
type OSSet struct {
	Free      unsafe.Pointer // Deallocates or releases any resources used by the OSSet instance.
	Init      unsafe.Pointer
	Member    unsafe.Pointer // Checks the set for the presence of an object.
	Merge     unsafe.Pointer
	Serialize unsafe.Pointer // Archives the receiver into the provided OSSerialize object.

}

// OSString - OSString wraps a C string in a C++ object for use in Libkern collections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/osstring
type OSString struct {
	Free      unsafe.Pointer // Deallocates or releases any resources used by the OSString instance.
	Serialize unsafe.Pointer // Archives the receiver into the provided OSSerialize object.

}

// OSSymbol - OSSymbol wraps a C string in a unique C++ object for use as keys in Libkern collections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ossymbol
type OSSymbol struct {
	Free       unsafe.Pointer // Overrides OSObject::free to synchronize with the symbol pool.
	Initialize unsafe.Pointer
}

// PassthruInterruptController
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/passthruinterruptcontroller
type PassthruInterruptController struct {
	Init unsafe.Pointer
}

// StdFBShmem_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/stdfbshmem_t
type StdFBShmem_t struct {
	CursorSize IOGSize
	Frame      unsafe.Pointer
}

// U128
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/u128
type U128 struct {
}

// UCInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ucinfo
type UCInfo struct {
}

// Accessx_descriptor
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/accessx_descriptor
type Accessx_descriptor struct {
	Ad_name_offset uint
	Ad_flags       int
	Ad_pad         unsafe.Pointer
}

// Ah
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ah
type Ah struct {
	Ah_spi     U_int32_t
	Ah_nxt     U_int8_t
	Ah_len     U_int8_t
	Ah_reserve U_int16_t
}

// Applelabel
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/applelabel
type Applelabel struct {
	Al_boot0    uint8
	Al_boot1    uint8
	Al_type     uint16
	Al_size     unsafe.Pointer
	Al_checksum unsafe.Pointer
	Al_offset   unsafe.Pointer
	Al_magic    uint16
	Al_flags    unsafe.Pointer
}

// Arcade_upcall_subsystem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/arcade_upcall_subsystem-4t8
type Arcade_upcall_subsystem struct {
	End      int32
	Maxsize  unsafe.Pointer
	Reserved Vm_address_t
	Routine  unsafe.Pointer
	Server   unsafe.Pointer
	Start    int32
}

// Arm_cpmu_state64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/arm_cpmu_state64
type Arm_cpmu_state64 struct {
	Ctrs uint64
}

// Arphdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/arphdr
type Arphdr struct {
	Ar_hrd U_short
	Ar_pln U_char
	Ar_op  U_short
	Ar_pro U_short
}

// Arpreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/arpreq
type Arpreq struct {
	Arp_flags unsafe.Pointer
	Arp_pa    unsafe.Pointer
	Arp_ha    unsafe.Pointer
}

// Arpstat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/arpstat
type Arpstat struct {
	Txrequests  unsafe.Pointer
	Txannounces unsafe.Pointer
	Rxreplies   unsafe.Pointer
	Rxrequests  unsafe.Pointer
	Dropped     unsafe.Pointer
	Inuse       unsafe.Pointer
	Invalidreqs unsafe.Pointer
	Timeouts    unsafe.Pointer
	Reqnobufs   unsafe.Pointer
	Txreplies   unsafe.Pointer
	Received    unsafe.Pointer
	Txurequests unsafe.Pointer
	Dupips      unsafe.Pointer
	Purged      unsafe.Pointer
	Txconflicts unsafe.Pointer
	Held        unsafe.Pointer
}

// Audit_triggers_subsystem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/audit_triggers_subsystem-162
type Audit_triggers_subsystem struct {
	Reserved Vm_address_t
	Routine  unsafe.Pointer
	Server   unsafe.Pointer
	Maxsize  unsafe.Pointer
	End      int32
	Start    int32
}

// Backtrace_control
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/backtrace_control
type Backtrace_control struct {
	Btc_flags             Backtrace_flags_t
	Btc_frame_addr        uintptr
	Btc_user_thread       unsafe.Pointer
	Btc_user_copy         unsafe.Pointer
	Btc_user_copy_context unsafe.Pointer
	Btc_addr_offset       unsafe.Pointer
}

// Backtrace_user_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/backtrace_user_info
type Backtrace_user_info struct {
	Btui_async_frame_addr  uintptr
	Btui_async_start_index unsafe.Pointer
}

// Bdevsw
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/bdevsw
type Bdevsw struct {
	D_open     *Open_close_fcn_t
	D_psize    *Psize_fcn_t
	D_strategy *Strategy_fcn_t
	D_type     unsafe.Pointer
	D_close    *Open_close_fcn_t
	D_ioctl    *Ioctl_fcn_t
	D_dump     *Dump_fcn_t
}

// Bootp
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/bootp
type Bootp struct {
	Bp_chaddr U_char
	Bp_ciaddr In_addr
	Bp_file   U_char
	Bp_giaddr In_addr
	Bp_hlen   U_char
	Bp_hops   U_char
	Bp_htype  U_char
	Bp_op     U_char
	Bp_secs   U_short
	Bp_siaddr In_addr
	Bp_sname  U_char
	Bp_unused U_short
	Bp_vend   U_char
	Bp_xid    U_int32_t
	Bp_yiaddr In_addr
}

// Bootp_packet
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/bootp_packet
type Bootp_packet struct {
	Bp_bootp Bootp
	Bp_ip    Ip
	Bp_udp   Udphdr
}

// Bpf_dltlist
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/bpf_dltlist
type Bpf_dltlist struct {
	Bfl_len U_int32_t
	Bfl_u   unsafe.Pointer
}

// Bpf_hdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/bpf_hdr
type Bpf_hdr struct {
	Bh_caplen  Bpf_u_int32
	Bh_tstamp  Timeval32
	Bh_datalen Bpf_u_int32
	Bh_hdrlen  U_short
}

// Bpf_insn
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/bpf_insn
type Bpf_insn struct {
	Jt   U_char
	Jf   U_char
	K    Bpf_u_int32
	Code U_short
}

// Bpf_program
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/bpf_program
type Bpf_program struct {
	Bf_len   U_int
	Bf_insns *Bpf_insn
}

// Bpf_stat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/bpf_stat
type Bpf_stat struct {
	Bs_drop U_int
	Bs_recv U_int
}

// Bpf_version
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/bpf_version
type Bpf_version struct {
	Bv_major unsafe.Pointer
	Bv_minor unsafe.Pointer
}

// Bt_params
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/bt_params
type Bt_params struct {
	Base_remote_ts unsafe.Pointer
	Rate           unsafe.Pointer
}

// Btinfo_sc_load_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/btinfo_sc_load_info
type Btinfo_sc_load_info struct {
	SharedCacheBaseAddress unsafe.Pointer
	SharedCacheSlide       unsafe.Pointer
	SharedCacheUUID        [16]byte
}

// Btinfo_sc_load_info64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/btinfo_sc_load_info64
type Btinfo_sc_load_info64 struct {
	SharedCacheBaseAddress unsafe.Pointer
	SharedCacheSlide       unsafe.Pointer
	SharedCacheUUID        [16]byte
}

// Btinfo_thread_state_data_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/btinfo_thread_state_data_t
type Btinfo_thread_state_data_t struct {
	Count  unsafe.Pointer
	Flavor unsafe.Pointer
	Tstate unsafe.Pointer
}

// Build_tool_version
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/build_tool_version
type Build_tool_version struct {
	Tool    unsafe.Pointer
	Version unsafe.Pointer
}

// Build_version_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/build_version_command
type Build_version_command struct {
	Cmd      unsafe.Pointer
	Cmdsize  unsafe.Pointer
	Minos    unsafe.Pointer
	Ntools   unsafe.Pointer
	Platform unsafe.Pointer
	Sdk      unsafe.Pointer
}

// Catch_exc_subsystem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/catch_exc_subsystem-t6n
type Catch_exc_subsystem struct {
	Server   unsafe.Pointer
	Reserved Vm_address_t
	Routine  unsafe.Pointer
	End      int32
	Start    int32
	Maxsize  unsafe.Pointer
}

// Catch_mach_exc_subsystem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/catch_mach_exc_subsystem-j9m
type Catch_mach_exc_subsystem struct {
	Maxsize  unsafe.Pointer
	Server   unsafe.Pointer
	Routine  unsafe.Pointer
	Start    int32
	Reserved Vm_address_t
	End      int32
}

// Cdevsw
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/cdevsw
type Cdevsw struct {
	D_reserved_2 *Rsvd_fcn_t
	D_reset      *Reset_fcn_t
	D_mmap       *Mmap_fcn_t
	D_close      *Open_close_fcn_t
	D_stop       *Stop_fcn_t
	D_type       unsafe.Pointer
	D_ioctl      *Ioctl_fcn_t
	D_open       *Open_close_fcn_t
	D_read       *Read_write_fcn_t
	D_select     *Select_fcn_t
	D_write      *Read_write_fcn_t
	D_reserved_1 *Rsvd_fcn_t
	D_ttys       unsafe.Pointer
	D_strategy   *Strategy_fcn_t
}

// Chain_len_stats
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/chain_len_stats
type Chain_len_stats struct {
	Cls_five_or_more unsafe.Pointer
	Cls_four         unsafe.Pointer
}

// Clock_reply_subsystem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/clock_reply_subsystem
type Clock_reply_subsystem struct {
	Maxsize  unsafe.Pointer
	End      int32
	Reserved Vm_address_t
	Routine  unsafe.Pointer
	Server   unsafe.Pointer
	Start    int32
}

// Clockinfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/clockinfo
type Clockinfo struct {
	Tickadj unsafe.Pointer
	Profhz  unsafe.Pointer
	Stathz  unsafe.Pointer
	Hz      unsafe.Pointer
	Tick    unsafe.Pointer
}

// Cmsghdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/cmsghdr
type Cmsghdr struct {
	Cmsg_len   uint32
	Cmsg_type  unsafe.Pointer
	Cmsg_level unsafe.Pointer
}

// Coalition_notification_subsystem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/coalition_notification_subsystem-36b
type Coalition_notification_subsystem struct {
	Start    int32
	Reserved Vm_address_t
	Maxsize  unsafe.Pointer
	End      int32
	Server   unsafe.Pointer
	Routine  unsafe.Pointer
}

// Codesigning_exit_reason_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/codesigning_exit_reason_info
type Codesigning_exit_reason_info struct {
	Ceri_codesig_modtime_nsecs unsafe.Pointer
}

// Componentname
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/componentname
type Componentname struct {
	Cn_pnbuf     unsafe.Pointer
	Cn_flags     unsafe.Pointer
	Cn_nameptr   unsafe.Pointer
	Cn_pnlen     unsafe.Pointer
	Cn_namelen   unsafe.Pointer
	Cn_hash      unsafe.Pointer
	Cn_nameiop   unsafe.Pointer
	Cn_consume   unsafe.Pointer
	Cn_reserved2 unsafe.Pointer
	Cn_reserved1 unsafe.Pointer
}

// Console_ops
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/console_ops
type Console_ops struct {
	Putc unsafe.Pointer
	Getc unsafe.Pointer
}

// Crashinfo_jit_address_range
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/crashinfo_jit_address_range
type Crashinfo_jit_address_range struct {
	End_address   unsafe.Pointer
	Start_address unsafe.Pointer
}

// Crashinfo_mb
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/crashinfo_mb
type Crashinfo_mb struct {
	Data          unsafe.Pointer
	Start_address unsafe.Pointer
}

// Crashinfo_proc_uniqidentifierinfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/crashinfo_proc_uniqidentifierinfo
type Crashinfo_proc_uniqidentifierinfo struct {
	P_reserve4  unsafe.Pointer
	P_puniqueid unsafe.Pointer
	P_reserve2  unsafe.Pointer
	P_uniqueid  unsafe.Pointer
}

// Ctl_event_data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ctl_event_data
type Ctl_event_data struct {
	Ctl_id   U_int32_t // The kernel control id.
	Ctl_unit U_int32_t // The kernel control unit.

}

// Ctl_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ctl_info
type Ctl_info struct {
	Ctl_id   U_int32_t      // The kernel control id, filled out upon return.
	Ctl_name unsafe.Pointer // The kernel control name to find.

}

// Ctlname
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ctlname
type Ctlname struct {
	Ctl_name unsafe.Pointer
	Ctl_type unsafe.Pointer
}

// Data_in_code_entry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/data_in_code_entry
type Data_in_code_entry struct {
	Length uint16
	Kind   uint16
	Offset unsafe.Pointer
}

// Dirent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dirent
type Dirent struct {
	D_reclen uint16
	D_name   unsafe.Pointer
	D_namlen uint8
	D_type   uint8
	D_ino    uint64
}

// Direntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/direntry
type Direntry struct {
	D_ino     uint64
	D_seekoff uint64
	D_reclen  uint16
	D_namlen  uint16
	D_name    unsafe.Pointer
	D_type    uint8
}

// Disk_blk0
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/disk_blk0
type Disk_blk0 struct {
	Parts     Fdisk_part
	Signature unsafe.Pointer
	Bootcode  unsafe.Pointer
}

// Do_notify_subsystem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/do_notify_subsystem-q2j
type Do_notify_subsystem struct {
	Maxsize  unsafe.Pointer
	End      int32
	Start    int32
	Server   unsafe.Pointer
	Reserved Vm_address_t
	Routine  unsafe.Pointer
}

// Doubleagent_subsystem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/doubleagent_subsystem-1pl
type Doubleagent_subsystem struct {
	End      int32
	Maxsize  unsafe.Pointer
	Reserved Vm_address_t
	Routine  unsafe.Pointer
	Server   unsafe.Pointer
	Start    int32
}

// Dyld_aot_cache_uuid_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_aot_cache_uuid_info
type Dyld_aot_cache_uuid_info struct {
	AotSlidBaseAddress unsafe.Pointer
	AotUUID            [16]byte
	X86SlidBaseAddress unsafe.Pointer
	X86UUID            [16]byte
}

// Dyld_chained_fixups_header
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_fixups_header
type Dyld_chained_fixups_header struct {
	Imports_count  unsafe.Pointer
	Imports_format unsafe.Pointer
	Imports_offset unsafe.Pointer
	Starts_offset  unsafe.Pointer
	Symbols_format unsafe.Pointer
	Symbols_offset unsafe.Pointer
}

// Dyld_chained_import
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_import
type Dyld_chained_import struct {
	Lib_ordinal unsafe.Pointer
	Name_offset unsafe.Pointer
	Weak_import unsafe.Pointer
}

// Dyld_chained_import_addend
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_import_addend
type Dyld_chained_import_addend struct {
	Addend      unsafe.Pointer
	Lib_ordinal unsafe.Pointer
	Name_offset unsafe.Pointer
	Weak_import unsafe.Pointer
}

// Dyld_chained_import_addend64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_import_addend64
type Dyld_chained_import_addend64 struct {
	Addend      unsafe.Pointer
	Lib_ordinal unsafe.Pointer
	Name_offset unsafe.Pointer
	Reserved    unsafe.Pointer
	Weak_import unsafe.Pointer
}

// Dyld_chained_ptr_32_bind
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_ptr_32_bind
type Dyld_chained_ptr_32_bind struct {
	Addend  unsafe.Pointer
	Bind    unsafe.Pointer
	Next    unsafe.Pointer
	Ordinal unsafe.Pointer
}

// Dyld_chained_ptr_32_cache_rebase
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_ptr_32_cache_rebase
type Dyld_chained_ptr_32_cache_rebase struct {
	Next unsafe.Pointer
}

// Dyld_chained_ptr_32_firmware_rebase
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_ptr_32_firmware_rebase
type Dyld_chained_ptr_32_firmware_rebase struct {
	Next   unsafe.Pointer
	Target unsafe.Pointer
}

// Dyld_chained_ptr_32_rebase
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_ptr_32_rebase
type Dyld_chained_ptr_32_rebase struct {
	Bind   unsafe.Pointer
	Next   unsafe.Pointer
	Target unsafe.Pointer
}

// Dyld_chained_ptr_64_bind
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_ptr_64_bind
type Dyld_chained_ptr_64_bind struct {
	Addend   unsafe.Pointer
	Bind     unsafe.Pointer
	Next     unsafe.Pointer
	Ordinal  unsafe.Pointer
	Reserved unsafe.Pointer
}

// Dyld_chained_ptr_64_kernel_cache_rebase
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_ptr_64_kernel_cache_rebase
type Dyld_chained_ptr_64_kernel_cache_rebase struct {
	AddrDiv    unsafe.Pointer
	CacheLevel unsafe.Pointer
	Diversity  unsafe.Pointer
	IsAuth     unsafe.Pointer
	Key        unsafe.Pointer
	Next       unsafe.Pointer
	Target     unsafe.Pointer
}

// Dyld_chained_ptr_64_rebase
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_ptr_64_rebase
type Dyld_chained_ptr_64_rebase struct {
	Bind     unsafe.Pointer
	High8    unsafe.Pointer
	Next     unsafe.Pointer
	Reserved unsafe.Pointer
	Target   unsafe.Pointer
}

// Dyld_chained_ptr_arm64e_auth_bind
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_ptr_arm64e_auth_bind
type Dyld_chained_ptr_arm64e_auth_bind struct {
	AddrDiv   unsafe.Pointer
	Auth      unsafe.Pointer
	Bind      unsafe.Pointer
	Diversity unsafe.Pointer
	Key       unsafe.Pointer
	Next      unsafe.Pointer
	Ordinal   unsafe.Pointer
	Zero      unsafe.Pointer
}

// Dyld_chained_ptr_arm64e_auth_bind24
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_ptr_arm64e_auth_bind24
type Dyld_chained_ptr_arm64e_auth_bind24 struct {
	AddrDiv   unsafe.Pointer
	Auth      unsafe.Pointer
	Diversity unsafe.Pointer
	Zero      unsafe.Pointer
}

// Dyld_chained_ptr_arm64e_auth_rebase
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_ptr_arm64e_auth_rebase
type Dyld_chained_ptr_arm64e_auth_rebase struct {
	Key unsafe.Pointer
}

// Dyld_chained_ptr_arm64e_bind
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_ptr_arm64e_bind
type Dyld_chained_ptr_arm64e_bind struct {
	Addend  unsafe.Pointer
	Auth    unsafe.Pointer
	Bind    unsafe.Pointer
	Ordinal unsafe.Pointer
	Zero    unsafe.Pointer
}

// Dyld_chained_ptr_arm64e_bind24
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_ptr_arm64e_bind24
type Dyld_chained_ptr_arm64e_bind24 struct {
	Addend  unsafe.Pointer
	Auth    unsafe.Pointer
	Bind    unsafe.Pointer
	Next    unsafe.Pointer
	Ordinal unsafe.Pointer
	Zero    unsafe.Pointer
}

// Dyld_chained_ptr_arm64e_rebase
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_ptr_arm64e_rebase
type Dyld_chained_ptr_arm64e_rebase struct {
	Auth   unsafe.Pointer
	Bind   unsafe.Pointer
	Target unsafe.Pointer
}

// Dyld_chained_ptr_arm64e_shared_cache_auth_rebase
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_ptr_arm64e_shared_cache_auth_rebase
type Dyld_chained_ptr_arm64e_shared_cache_auth_rebase struct {
	AddrDiv       unsafe.Pointer
	Auth          unsafe.Pointer
	Diversity     unsafe.Pointer
	KeyIsData     unsafe.Pointer
	Next          unsafe.Pointer
	RuntimeOffset unsafe.Pointer
}

// Dyld_chained_ptr_arm64e_shared_cache_rebase
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_ptr_arm64e_shared_cache_rebase
type Dyld_chained_ptr_arm64e_shared_cache_rebase struct {
	Next          unsafe.Pointer
	RuntimeOffset unsafe.Pointer
}

// Dyld_chained_starts_in_image
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_starts_in_image
type Dyld_chained_starts_in_image struct {
	Seg_count       unsafe.Pointer
	Seg_info_offset unsafe.Pointer
}

// Dyld_chained_starts_in_segment
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_starts_in_segment
type Dyld_chained_starts_in_segment struct {
	Max_valid_pointer unsafe.Pointer
	Page_count        uint16
	Page_size         uint16
	Page_start        uint16
	Pointer_format    uint16
	Segment_offset    unsafe.Pointer
	Size              unsafe.Pointer
}

// Dyld_chained_starts_offsets
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_chained_starts_offsets
type Dyld_chained_starts_offsets struct {
	Chain_starts   unsafe.Pointer
	Pointer_format unsafe.Pointer
	Starts_count   unsafe.Pointer
}

// Dyld_info_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_info_command
type Dyld_info_command struct {
	Rebase_size    unsafe.Pointer
	Lazy_bind_off  unsafe.Pointer
	Cmd            unsafe.Pointer
	Lazy_bind_size unsafe.Pointer
	Bind_off       unsafe.Pointer
	Export_size    unsafe.Pointer
	Rebase_off     unsafe.Pointer
	Bind_size      unsafe.Pointer
	Weak_bind_size unsafe.Pointer
	Weak_bind_off  unsafe.Pointer
	Cmdsize        unsafe.Pointer
	Export_off     unsafe.Pointer
}

// Dyld_shared_cache_loadinfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_shared_cache_loadinfo
type Dyld_shared_cache_loadinfo struct {
	SharedCacheSlidFirstMapping          unsafe.Pointer
	SharedCacheSlide                     unsafe.Pointer
	SharedCacheUUID                      [16]byte
	SharedCacheUnreliableSlidBaseAddress unsafe.Pointer
}

// Dyld_shared_cache_loadinfo_v2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_shared_cache_loadinfo_v2
type Dyld_shared_cache_loadinfo_v2 struct {
	SharedCacheFlags                     unsafe.Pointer
	SharedCacheID                        unsafe.Pointer
	SharedCacheSlidFirstMapping          unsafe.Pointer
	SharedCacheSlide                     unsafe.Pointer
	SharedCacheUUID                      [16]byte
	SharedCacheUnreliableSlidBaseAddress unsafe.Pointer
}

// Dyld_uuid_info_32
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_uuid_info_32
type Dyld_uuid_info_32 struct {
	ImageLoadAddress unsafe.Pointer
	ImageUUID        [16]byte
}

// Dyld_uuid_info_64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_uuid_info_64
type Dyld_uuid_info_64 struct {
	ImageLoadAddress unsafe.Pointer
	ImageUUID        [16]byte
}

// Dyld_uuid_info_64_v2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dyld_uuid_info_64_v2
type Dyld_uuid_info_64_v2 struct {
	ImageSlidBaseAddress unsafe.Pointer
	ImageUUID            [16]byte
	ImageLoadAddress     unsafe.Pointer
}

// Dylib
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dylib
type Dylib struct {
	Name      unsafe.Pointer
	Timestamp unsafe.Pointer
}

// Dylib_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dylib_command
type Dylib_command struct {
	Cmdsize unsafe.Pointer
	Cmd     unsafe.Pointer // Common to all load command structures. For this structure, set to either `LC_LOAD_DYLIB`, `LC_LOAD_WEAK_DYLIB`, or `LC_ID_DYLIB`.
	Dylib   Dylib
}

// Dylib_module
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dylib_module
type Dylib_module struct {
	Objc_module_info_size unsafe.Pointer
	Iextrel               unsafe.Pointer // The index into the external relocation table of the first entry provided by this module.
	Ninit_nterm           unsafe.Pointer // Contains both the number of pointers in the module initialization (the low 16 bits) and the number of pointers in the module termination section (the high 16 bits) for this module.
	Iinit_iterm           unsafe.Pointer
	Objc_module_info_addr unsafe.Pointer
	Irefsym               unsafe.Pointer
	Nrefsym               unsafe.Pointer // The number of external reference entries provided by this module.
	Nextrel               unsafe.Pointer
	Iextdefsym            unsafe.Pointer
	Ilocalsym             unsafe.Pointer
	Nextdefsym            unsafe.Pointer
	Nlocalsym             unsafe.Pointer
}

// Dylib_module_64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dylib_module_64
type Dylib_module_64 struct {
	Ilocalsym             unsafe.Pointer
	Iextdefsym            unsafe.Pointer
	Nrefsym               unsafe.Pointer
	Nextdefsym            unsafe.Pointer
	Iextrel               unsafe.Pointer
	Objc_module_info_addr unsafe.Pointer
	Objc_module_info_size unsafe.Pointer
	Irefsym               unsafe.Pointer
	Iinit_iterm           unsafe.Pointer
	Ninit_nterm           unsafe.Pointer // Contains both the number of pointers in the module initialization (the low 16 bits) and the number of pointers in the module termination section (the high 16 bits) for this module.
	Nextrel               unsafe.Pointer
	Nlocalsym             unsafe.Pointer
}

// Dylib_reference
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dylib_reference
type Dylib_reference struct {
	Isym  unsafe.Pointer // An index into the symbol table for the symbol being referenced.
	Flags unsafe.Pointer
}

// Dylib_table_of_contents
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dylib_table_of_contents
type Dylib_table_of_contents struct {
	Symbol_index unsafe.Pointer // An index into the symbol table indicating the defined external symbol to which this entry refers.

}

// Dylinker_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dylinker_command
type Dylinker_command struct {
	Cmdsize unsafe.Pointer
	Name    unsafe.Pointer
	Cmd     unsafe.Pointer
}

// Dysymtab_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/dysymtab_command
type Dysymtab_command struct {
	Iundefsym      unsafe.Pointer
	Nlocalsym      unsafe.Pointer
	Nundefsym      unsafe.Pointer
	Nextdefsym     unsafe.Pointer
	Indirectsymoff unsafe.Pointer
	Ntoc           unsafe.Pointer
	Iextdefsym     unsafe.Pointer
	Nindirectsyms  unsafe.Pointer
	Nextrefsyms    unsafe.Pointer
	Tocoff         unsafe.Pointer
	Cmdsize        unsafe.Pointer
	Locreloff      unsafe.Pointer
	Nlocrel        unsafe.Pointer
	Ilocalsym      unsafe.Pointer
	Modtaboff      unsafe.Pointer
	Nextrel        unsafe.Pointer
	Extrefsymoff   unsafe.Pointer // An integer indicating the byte offset from the start of the file to the external reference table data.
	Cmd            unsafe.Pointer // Common to all load command structures. For this structure, set to `LC_DYSYMTAB`.
	Nmodtab        unsafe.Pointer
	Extreloff      unsafe.Pointer // An integer indicating the byte offset from the start of the file to the external relocation table data.

}

// Ecc_event
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ecc_event
type Ecc_event struct {
	Count uint8
	Id    uint8
	Data  unsafe.Pointer
}

// Efi_aurr_extended_panic_log
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/efi_aurr_extended_panic_log
type Efi_aurr_extended_panic_log struct {
	Efi_aurr_extended_log_buf unsafe.Pointer
	Efi_aurr_log_head         unsafe.Pointer
	Efi_aurr_log_tail         unsafe.Pointer
}

// Efi_aurr_panic_header
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/efi_aurr_panic_header
type Efi_aurr_panic_header struct {
	Efi_aurr_crc         unsafe.Pointer
	Efi_aurr_reset_cause unsafe.Pointer
}

// Embedded_panic_header
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/embedded_panic_header
type Embedded_panic_header struct {
	Eph_crc                    unsafe.Pointer
	Eph_magic                  unsafe.Pointer
	Eph_panic_flags            Eph_panic_flags_t
	Eph_stackshot_len          unsafe.Pointer
	Eph_other_log_len          unsafe.Pointer
	Eph_panic_log_offset       unsafe.Pointer
	Eph_other_log_offset       unsafe.Pointer
	Eph_stackshot_offset       unsafe.Pointer
	Eph_panic_log_len          unsafe.Pointer
	Eph_bootsessionuuid_string Uuid_string_t
	Eph_roots_installed        unsafe.Pointer
	Eph_ext_paniclog_len       unsafe.Pointer
	Eph_ext_paniclog_offset    unsafe.Pointer
	Eph_panic_initiator_len    unsafe.Pointer
	Eph_panic_initiator_offset unsafe.Pointer
}

// Encryption_info_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/encryption_info_command
type Encryption_info_command struct {
	Cmd       unsafe.Pointer
	Cryptsize unsafe.Pointer
	Cmdsize   unsafe.Pointer
	Cryptoff  unsafe.Pointer
	Cryptid   unsafe.Pointer
}

// Encryption_info_command_64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/encryption_info_command_64
type Encryption_info_command_64 struct {
	Cryptsize unsafe.Pointer
	Pad       unsafe.Pointer
	Cmdsize   unsafe.Pointer
	Cryptoff  unsafe.Pointer
	Cryptid   unsafe.Pointer
	Cmd       unsafe.Pointer
}

// Entry_point_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/entry_point_command
type Entry_point_command struct {
	Stacksize unsafe.Pointer
	Cmd       unsafe.Pointer
	Cmdsize   unsafe.Pointer
	Entryoff  unsafe.Pointer
}

// Esp
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/esp
type Esp struct {
	Esp_spi U_int32_t
}

// Esptail
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/esptail
type Esptail struct {
	Esp_padlen U_int8_t
	Esp_nxt    U_int8_t
}

// Ether_arp
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ether_arp
type Ether_arp struct {
	Arp_spa U_char
	Arp_sha U_char
}

// Ether_vlan_header
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ether_vlan_header
type Ether_vlan_header struct {
	Evl_dhost       U_char
	Evl_encap_proto U_int16_t
	Evl_proto       U_int16_t
	Evl_shost       U_char
	Evl_tag         U_int16_t
}

// Exclave_addressspace_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/exclave_addressspace_info
type Exclave_addressspace_info struct {
	Eas_flags    unsafe.Pointer
	Eas_id       unsafe.Pointer
	Eas_layoutid unsafe.Pointer
	Eas_slide    unsafe.Pointer
	Eas_asroot   unsafe.Pointer
}

// Exclave_ipcstackentry_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/exclave_ipcstackentry_info
type Exclave_ipcstackentry_info struct {
	Eise_asid         unsafe.Pointer
	Eise_flags        unsafe.Pointer
	Eise_invocationid unsafe.Pointer
	Eise_tnid         unsafe.Pointer
}

// Exclave_scresult_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/exclave_scresult_info
type Exclave_scresult_info struct {
	Esc_flags unsafe.Pointer
	Esc_id    unsafe.Pointer
}

// Exclave_textlayout_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/exclave_textlayout_info
type Exclave_textlayout_info struct {
	Etl_flags unsafe.Pointer
	Layout_id unsafe.Pointer
}

// Exclave_textlayout_info_v1
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/exclave_textlayout_info_v1
type Exclave_textlayout_info_v1 struct {
}

// Exclave_textlayout_segment
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/exclave_textlayout_segment
type Exclave_textlayout_segment struct {
	LayoutSegment_loadAddress unsafe.Pointer
	LayoutSegment_uuid        [16]byte
}

// Exclave_textlayout_segment_v2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/exclave_textlayout_segment_v2
type Exclave_textlayout_segment_v2 struct {
	LayoutSegment_loadAddress    unsafe.Pointer
	LayoutSegment_rawLoadAddress unsafe.Pointer
	LayoutSegment_uuid           [16]byte
}

// Exit_reason_snapshot
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/exit_reason_snapshot
type Exit_reason_snapshot struct {
	Ers_namespace unsafe.Pointer
	Ers_flags     unsafe.Pointer
	Ers_code      unsafe.Pointer
}

// Experiment_spec
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/experiment_spec
type Experiment_spec struct {
	Max_value      unsafe.Pointer
	Min_value      unsafe.Pointer
	Modified       unsafe.Pointer
	Original_value unsafe.Pointer
	Ptr            unsafe.Pointer
}

// Fairplay_subsystem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/fairplay_subsystem-4tk
type Fairplay_subsystem struct {
	End      int32
	Maxsize  unsafe.Pointer
	Reserved Vm_address_t
	Routine  unsafe.Pointer
	Server   unsafe.Pointer
	Start    int32
}

// Fat_arch - Describes the location within the binary of an object file targeted at a single architecture. Declared in `/usr/include/mach-o/fat.H()`.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/fat_arch
type Fat_arch struct {
	Cpusubtype int32          // An enumeration value of type `cpu_subtype_t`. Specifies the specific member of the CPU family on which this entry may be used or a constant specifying all members.
	Align      unsafe.Pointer // The power of 2 alignment for the offset of the object file for the architecture specified in `cputype` within the binary. This is required to ensure that, if this binary is changed, the contents it retains are correctly aligned for virtual memory paging and other uses.
	Cputype    int32
	Offset     unsafe.Pointer // Offset to the beginning of the data for this CPU.
	Size       unsafe.Pointer // Size of the data for this CPU.

}

// Fat_header
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/fat_header
type Fat_header struct {
	Nfat_arch unsafe.Pointer // An integer specifying the number of [fat_arch](<doc://com.apple.documentation/documentation/kernel/fat_arch>) data structures that follow. This is the number of architectures contained in this binary.
	Magic     unsafe.Pointer // An integer containing the value `0xCAFEBABE` in big-endian byte order format. On a big-endian host CPU, this can be validated using the constant `FAT_MAGIC`; on a little-endian host CPU, it can be validated using the constant `FAT_CIGAM`.

}

// Fdisk_part
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/fdisk_part
type Fdisk_part struct {
	Endcyl  unsafe.Pointer
	Numsect unsafe.Pointer
	Systid  unsafe.Pointer
	Beghead unsafe.Pointer
	Begsect unsafe.Pointer
	Endhead unsafe.Pointer
	Bootid  unsafe.Pointer
	Endsect unsafe.Pointer
	Relsect unsafe.Pointer
	Begcyl  unsafe.Pointer
}

// Fileset_entry_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/fileset_entry_command
type Fileset_entry_command struct {
	Cmd      unsafe.Pointer
	Cmdsize  unsafe.Pointer
	Entry_id unsafe.Pointer
	Fileoff  unsafe.Pointer
	Reserved unsafe.Pointer
	Vmaddr   unsafe.Pointer
}

// Flock
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/flock
type Flock struct {
	L_start  int64
	L_len    int64
	L_pid    int32
	L_type   int16
	L_whence int16
}

// Flocktimeout
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/flocktimeout
type Flocktimeout struct {
	Fl      Flock
	Timeout Timespec
}

// Frmrinfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/frmrinfo
type Frmrinfo struct {
	Frmr_cause       U_int8_t
	Frmr_control     U_int8_t
	Frmr_rej_pdu1    U_int8_t
	Frmr_rej_pdu0    U_int8_t
	Frmr_control_ext U_int8_t
}

// Fs_snapshot_mount_args
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/fs_snapshot_mount_args
type Fs_snapshot_mount_args struct {
	Sm_cnp *Componentname
	Sm_mp  Mount_t
}

// Fs_snapshot_revert_args
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/fs_snapshot_revert_args
type Fs_snapshot_revert_args struct {
	Sr_cnp *Componentname
}

// Fs_snapshot_root_args
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/fs_snapshot_root_args
type Fs_snapshot_root_args struct {
	Sr_cnp *Componentname
}

// Fssearchblock
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/fssearchblock
type Fssearchblock struct {
	Searchparams2       unsafe.Pointer
	Returnattrs         *Attrlist
	Returnbuffer        unsafe.Pointer
	Searchparams1       unsafe.Pointer
	Sizeofsearchparams1 uintptr
	Timelimit           Timeval
	Searchattrs         Attrlist
	Sizeofsearchparams2 uintptr
	Returnbuffersize    uintptr
	Maxmatches          U_long
}

// Fvmfile_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/fvmfile_command
type Fvmfile_command struct {
	Cmd unsafe.Pointer
}

// Fvmlib
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/fvmlib
type Fvmlib struct {
	Header_addr unsafe.Pointer
	Name        unsafe.Pointer
}

// Fvmlib_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/fvmlib_command
type Fvmlib_command struct {
	Cmd unsafe.Pointer
}

// Gpt_ent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/gpt_ent
type Gpt_ent struct {
	Ent_lba_end   unsafe.Pointer
	Ent_uuid      [16]byte
	Ent_type      [16]byte
	Ent_lba_start unsafe.Pointer
	Ent_name      uint16
	Ent_attr      unsafe.Pointer
}

// Gpt_hdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/gpt_hdr
type Gpt_hdr struct {
	Padding       unsafe.Pointer
	Hdr_entries   unsafe.Pointer
	Hdr_crc_table unsafe.Pointer
	Hdr_crc_self  unsafe.Pointer
	Hdr_size      unsafe.Pointer
	Hdr_lba_table unsafe.Pointer
	Hdr_lba_self  unsafe.Pointer
	Hdr_lba_alt   unsafe.Pointer
	Hdr_uuid      [16]byte
	Hdr_lba_start unsafe.Pointer
	Hdr_entsz     unsafe.Pointer
	Hdr_sig       uint8
	Hdr_lba_end   unsafe.Pointer
	Hdr_revision  unsafe.Pointer
}

// Group_req
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/group_req
type Group_req struct {
	Gr_interface unsafe.Pointer
	Gr_group     unsafe.Pointer
}

// Group_source_req
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/group_source_req
type Group_source_req struct {
	Gsr_source    unsafe.Pointer
	Gsr_interface unsafe.Pointer
	Gsr_group     unsafe.Pointer
}

// Hfs_mount_args
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/hfs_mount_args
type Hfs_mount_args struct {
	Hfs_mask             uint16
	Hfs_gid              uint32
	Flags                unsafe.Pointer
	Journal_tbuffer_size unsafe.Pointer
	Hfs_encoding         U_int32_t
	Journal_disable      unsafe.Pointer
	Journal_flags        unsafe.Pointer
	Hfs_uid              uint32
	Hfs_timezone         Timezone
}

// Icmp
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/icmp
type Icmp struct {
	Icmp_cksum U_short
	Icmp_code  U_char
	Icmp_dun   unsafe.Pointer
	Icmp_hun   unsafe.Pointer
	Icmp_type  U_char
}

// Icmp6_filter
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/icmp6_filter
type Icmp6_filter struct {
	Icmp6_filt U_int32_t
}

// Icmp6_hdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/icmp6_hdr
type Icmp6_hdr struct {
	Icmp6_cksum  U_int16_t
	Icmp6_code   U_int8_t
	Icmp6_dataun unsafe.Pointer
	Icmp6_type   U_int8_t
}

// Icmp6_ifstat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/icmp6_ifstat
type Icmp6_ifstat struct {
	Ifs6_in_adminprohib U_quad_t
	Ifs6_in_dstunreach  U_quad_t
	Ifs6_in_msg         U_quad_t
}

// Icmp6_namelookup
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/icmp6_namelookup
type Icmp6_namelookup struct {
	Icmp6_nl_hdr   Icmp6_hdr
	Icmp6_nl_nonce U_int8_t
	Icmp6_nl_ttl   unsafe.Pointer
}

// Icmp6_nodeinfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/icmp6_nodeinfo
type Icmp6_nodeinfo struct {
	Icmp6_ni_hdr   Icmp6_hdr
	Icmp6_ni_nonce U_int8_t
}

// Icmp6_router_renum
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/icmp6_router_renum
type Icmp6_router_renum struct {
	Rr_flags    U_int8_t
	Rr_hdr      Icmp6_hdr
	Rr_maxdelay U_int16_t
	Rr_reserved U_int32_t
	Rr_segnum   U_int8_t
}

// Icmp6errstat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/icmp6errstat
type Icmp6errstat struct {
	Icp6errs_dst_unreach_addr        U_quad_t
	Icp6errs_dst_unreach_admin       U_quad_t
	Icp6errs_dst_unreach_beyondscope U_quad_t
	Icp6errs_dst_unreach_noport      U_quad_t
	Icp6errs_dst_unreach_noroute     U_quad_t
	Icp6errs_packet_too_big          U_quad_t
	Icp6errs_paramprob_header        U_quad_t
	Icp6errs_paramprob_nextheader    U_quad_t
	Icp6errs_paramprob_option        U_quad_t
	Icp6errs_redirect                U_quad_t
	Icp6errs_time_exceed_reassembly  U_quad_t
	Icp6errs_time_exceed_transit     U_quad_t
	Icp6errs_unknown                 U_quad_t
}

// Icmp6stat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/icmp6stat
type Icmp6stat struct {
	Icp6s_badcode       U_quad_t
	Icp6s_badlen        U_quad_t
	Icp6s_badna         U_quad_t
	Icp6s_badns         U_quad_t
	Icp6s_badpkttoobig  U_quad_t
	Icp6s_badra         U_quad_t
	Icp6s_badredirect   U_quad_t
	Icp6s_badrs         U_quad_t
	Icp6s_canterror     U_quad_t
	Icp6s_checksum      U_quad_t
	Icp6s_error         U_quad_t
	Icp6s_inhist        U_quad_t
	Icp6s_nd_badopt     U_quad_t
	Icp6s_nd_toomanyopt U_quad_t
	Icp6s_outerrhist    Icmp6errstat
	Icp6s_outhist       U_quad_t
	Icp6s_pmtuchg       U_quad_t
	Icp6s_reflect       U_quad_t
	Icp6s_rfc6980_drop  U_quad_t
	Icp6s_toofreq       U_quad_t
	Icp6s_tooshort      U_quad_t
}

// Icmp_ra_addr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/icmp_ra_addr
type Icmp_ra_addr struct {
	Ira_addr       U_int32_t
	Ira_preference U_int32_t
}

// Icmpstat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/icmpstat
type Icmpstat struct {
	Icps_badcode      U_int32_t
	Icps_badlen       U_int32_t
	Icps_bmcastecho   U_int32_t
	Icps_bmcasttstamp U_int32_t
	Icps_checksum     U_int32_t
	Icps_error        U_int32_t
	Icps_inhist       U_int32_t
	Icps_oldicmp      U_int32_t
	Icps_oldshort     U_int32_t
	Icps_outhist      U_int32_t
	Icps_reflect      U_int32_t
	Icps_tooshort     U_int32_t
}

// Ident_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ident_command
type Ident_command struct {
	Cmdsize unsafe.Pointer
	Cmd     unsafe.Pointer
}

// If_agentidreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_agentidreq
type If_agentidreq struct {
	Ifar_name unsafe.Pointer
	Ifar_uuid [16]byte
}

// If_agentidsreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_agentidsreq
type If_agentidsreq struct {
	Ifar_count U_int32_t
	Ifar_name  unsafe.Pointer
	Ifar_uuids *[16]byte
}

// If_bandwidths
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_bandwidths
type If_bandwidths struct {
	Eff_bw unsafe.Pointer
	Max_bw unsafe.Pointer
}

// If_cellular_status
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_cellular_status
type If_cellular_status struct {
	If_cell_u unsafe.Pointer
}

// If_cellular_status_v1
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_cellular_status_v1
type If_cellular_status_v1 struct {
	Config_backoff_time    U_int32_t
	Config_inactivity_time U_int32_t
	Dl_effective_bandwidth U_int32_t
	Dl_max_bandwidth       U_int32_t
	Link_quality_metric    U_int32_t
	Mss_recommended        U_int16_t
	Reserved_1             U_int16_t
	Reserved_2             U_int32_t
	Reserved_3             U_int64_t
	Reserved_4             U_int64_t
	Reserved_5             U_int64_t
	Reserved_6             U_int64_t
	Ul_avg_queue_size      U_int32_t
	Ul_bytes_lost          U_int32_t
	Ul_effective_bandwidth U_int32_t
	Ul_effective_latency   U_int32_t
	Ul_max_bandwidth       U_int32_t
	Ul_max_latency         U_int32_t
	Ul_max_queue_size      U_int32_t
	Ul_min_latency         U_int32_t
	Ul_min_queue_size      U_int32_t
	Ul_retxt_level         U_int32_t
	Valid_bitmask          U_int32_t
}

// If_clat46req
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_clat46req
type If_clat46req struct {
	Ifclat46_addr If_ipv6_address
	Ifclat46_name unsafe.Pointer
}

// If_clonereq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_clonereq
type If_clonereq struct {
	Ifcr_count  unsafe.Pointer
	Ifcr_total  unsafe.Pointer
	Ifcr_buffer unsafe.Pointer
}

// If_data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_data
type If_data struct {
	Ifi_unused1    U_char
	Ifi_noproto    U_int32_t
	Ifi_xmitquota  U_char
	Ifi_obytes     U_int32_t
	Ifi_imcasts    U_int32_t
	Ifi_physical   U_char
	Ifi_metric     U_int32_t
	Ifi_iqdrops    U_int32_t
	Ifi_reserved2  U_int32_t
	Ifi_ipackets   U_int32_t
	Ifi_mtu        U_int32_t
	Ifi_reserved1  U_int32_t
	Ifi_recvquota  U_char
	Ifi_typelen    U_char
	Ifi_hdrlen     U_char
	Ifi_recvtiming U_int32_t
	Ifi_addrlen    U_char
	Ifi_omcasts    U_int32_t
	Ifi_hwassist   U_int32_t
	Ifi_ierrors    U_int32_t
	Ifi_collisions U_int32_t
	Ifi_xmittiming U_int32_t
	Ifi_lastchange Timeval32
	Ifi_oerrors    U_int32_t
	Ifi_ibytes     U_int32_t
	Ifi_opackets   U_int32_t
	Ifi_baudrate   U_int32_t
	Ifi_unused2    U_int32_t
	Ifi_type       U_char
}

// If_data64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_data64
type If_data64 struct {
	Ifi_noproto    U_int64_t
	Ifi_addrlen    U_char
	Ifi_baudrate   U_int64_t
	Ifi_type       U_char
	Ifi_ibytes     U_int64_t
	Ifi_unused1    U_char
	Ifi_typelen    U_char
	Ifi_recvtiming U_int32_t
	Ifi_mtu        U_int32_t
	Ifi_opackets   U_int64_t
	Ifi_iqdrops    U_int64_t
	Ifi_imcasts    U_int64_t
	Ifi_ierrors    U_int64_t
	Ifi_oerrors    U_int64_t
	Ifi_hdrlen     U_char
	Ifi_xmitquota  U_char
	Ifi_lastchange Timeval32
	Ifi_obytes     U_int64_t
	Ifi_physical   U_char
	Ifi_metric     U_int32_t
	Ifi_recvquota  U_char
	Ifi_omcasts    U_int64_t
	Ifi_xmittiming U_int32_t
	Ifi_ipackets   U_int64_t
	Ifi_collisions U_int64_t
}

// If_data_extended
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_data_extended
type If_data_extended struct {
	Ifi_alignerrs U_int64_t
	Ifi_dt_bytes  U_int64_t
	Ifi_fbytes    U_int64_t
	Ifi_fpackets  U_int64_t
	Reserved      U_int64_t
}

// If_descreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_descreq
type If_descreq struct {
	Ifdr_desc U_int8_t
	Ifdr_len  U_int32_t
	Ifdr_name unsafe.Pointer
}

// If_description
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_description
type If_description struct {
	Ifd_desc   *U_int8_t
	Ifd_len    U_int32_t
	Ifd_maxlen U_int32_t
}

// If_interface_state
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_interface_state
type If_interface_state struct {
	Interface_availability U_int8_t
	Lqm_state              int8
	Rrc_state              U_int8_t
	Valid_bitmask          U_int8_t
}

// If_ipv6_address
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_ipv6_address
type If_ipv6_address struct {
	V6_address   unsafe.Pointer
	V6_prefixlen unsafe.Pointer
}

// If_latencies
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_latencies
type If_latencies struct {
	Eff_lt U_int64_t
	Max_lt U_int64_t
}

// If_lim_perf_stat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_lim_perf_stat
type If_lim_perf_stat struct {
	Lim_bk_txpkts            U_int64_t
	Lim_conn_attempts        U_int64_t
	Lim_conn_timeout_percent U_int64_t
	Lim_conn_timeouts        U_int64_t
	Lim_dl_detected          U_int64_t
	Lim_dl_max_bandwidth     U_int64_t
	Lim_packet_loss_percent  U_int64_t
	Lim_packet_ooo_percent   U_int64_t
	Lim_rtt_average          U_int64_t
	Lim_rtt_min              U_int64_t
	Lim_rtt_variance         U_int64_t
	Lim_total_oopkts         U_int64_t
	Lim_total_retxpkts       U_int64_t
	Lim_total_rxpkts         U_int64_t
	Lim_total_txpkts         U_int64_t
	Lim_ul_detected          U_int64_t
	Lim_ul_max_bandwidth     U_int64_t
}

// If_link_status
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_link_status
type If_link_status struct {
	Ifsr_len U_int32_t
	Ifsr_u   unsafe.Pointer
}

// If_linkheuristics
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_linkheuristics
type If_linkheuristics struct {
	Iflh_congested_link_cnt       U_int64_t
	Iflh_congested_link_time      U_int64_t
	Iflh_link_heuristics_cnt      U_int64_t
	Iflh_link_heuristics_time     U_int64_t
	Iflh_lqm_bad_cnt              U_int64_t
	Iflh_lqm_bad_time             U_int64_t
	Iflh_lqm_good_cnt             U_int64_t
	Iflh_lqm_good_time            U_int64_t
	Iflh_lqm_min_viable_cnt       U_int64_t
	Iflh_lqm_min_viable_time      U_int64_t
	Iflh_lqm_poor_cnt             U_int64_t
	Iflh_lqm_poor_time            U_int64_t
	Iflh_tcp_linkheur_comprxmt    U_int64_t
	Iflh_tcp_linkheur_noackpri    U_int64_t
	Iflh_tcp_linkheur_rxmtfloor   U_int64_t
	Iflh_tcp_linkheur_stealthdrop U_int64_t
	Iflh_tcp_linkheur_synrxmt     U_int64_t
	Iflh_udp_linkheur_stealthdrop U_int64_t
}

// If_linkparamsreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_linkparamsreq
type If_linkparamsreq struct {
	Iflpr_flags          U_int32_t
	Iflpr_input_bw       If_bandwidths
	Iflpr_input_lt       If_latencies
	Iflpr_input_netem    If_netem_params
	Iflpr_input_tbr_rate U_int64_t
	Iflpr_output_lt      If_latencies
	Iflpr_output_sched   U_int32_t
}

// If_msghdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_msghdr
type If_msghdr struct {
	Ifm_msglen unsafe.Pointer
	Ifm_data   If_data
	Ifm_index  unsafe.Pointer
	Ifm_addrs  unsafe.Pointer
	Ifm_flags  unsafe.Pointer
	Ifm_type   unsafe.Pointer
}

// If_msghdr2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_msghdr2
type If_msghdr2 struct {
	Ifm_snd_len    unsafe.Pointer
	Ifm_type       U_char
	Ifm_timer      unsafe.Pointer
	Ifm_snd_maxlen unsafe.Pointer
	Ifm_msglen     U_short
	Ifm_index      U_short
}

// If_nat64req
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_nat64req
type If_nat64req struct {
	Ifnat64_name     unsafe.Pointer
	Ifnat64_prefixes Ipv6_prefix
}

// If_netem_params
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_netem_params
type If_netem_params struct {
	Ifnetem_bandwidth_bps    unsafe.Pointer
	Ifnetem_corruption_p     unsafe.Pointer
	Ifnetem_duplication_p    unsafe.Pointer
	Ifnetem_jitter_ms        unsafe.Pointer
	Ifnetem_latency_ms       unsafe.Pointer
	Ifnetem_loss_p_bl_br     unsafe.Pointer
	Ifnetem_loss_p_bl_gr     unsafe.Pointer
	Ifnetem_loss_p_br_bl     unsafe.Pointer
	Ifnetem_loss_p_gr_bl     unsafe.Pointer
	Ifnetem_loss_p_gr_gl     unsafe.Pointer
	Ifnetem_loss_recovery_ms unsafe.Pointer
	Ifnetem_model            If_netem_model_t
	Ifnetem_output_ival_ms   unsafe.Pointer
	Ifnetem_reordering_p     unsafe.Pointer
}

// If_netidreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_netidreq
type If_netidreq struct {
	Ifnetid      U_int8_t
	Ifnetid_len  U_int8_t
	Ifnetid_name unsafe.Pointer
}

// If_netif_stats
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_netif_stats
type If_netif_stats struct {
	Ifn_rx_mit_bytes_avg         U_int32_t
	Ifn_rx_mit_bytes_max         U_int32_t
	Ifn_rx_mit_bytes_min         U_int32_t
	Ifn_rx_mit_cfg_bytes_hiwat   U_int32_t
	Ifn_rx_mit_cfg_bytes_lowat   U_int32_t
	Ifn_rx_mit_cfg_idx           U_int32_t
	Ifn_rx_mit_cfg_interval      U_int32_t
	Ifn_rx_mit_cfg_packets_hiwat U_int32_t
	Ifn_rx_mit_cfg_packets_lowat U_int32_t
	Ifn_rx_mit_interval          U_int64_t
	Ifn_rx_mit_mode              U_int32_t
	Ifn_rx_mit_packets_avg       U_int32_t
	Ifn_rx_mit_packets_max       U_int32_t
	Ifn_rx_mit_packets_min       U_int32_t
}

// If_nexusreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_nexusreq
type If_nexusreq struct {
	Ifnr_flags      unsafe.Pointer
	Ifnr_flowswitch [16]byte
	Ifnr_name       unsafe.Pointer
	Ifnr_netif      [16]byte
	Ifnr_reserved   unsafe.Pointer
}

// If_nsreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_nsreq
type If_nsreq struct {
	Ifnsr_data   U_int8_t
	Ifnsr_family U_int8_t
	Ifnsr_flags  U_int16_t
	Ifnsr_len    U_int8_t
	Ifnsr_name   unsafe.Pointer
}

// If_order
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_order
type If_order struct {
	Ifo_count           U_int32_t
	Ifo_ordered_indices Mach_vm_address_t
	Ifo_reserved        U_int32_t
}

// If_packet_stats
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_packet_stats
type If_packet_stats struct {
	Ifi_udp_port_unreach U_int64_t
}

// If_protolistreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_protolistreq
type If_protolistreq struct {
	Ifpl_count    U_int32_t
	Ifpl_list     *U_int32_t
	Ifpl_name     unsafe.Pointer
	Ifpl_reserved U_int32_t
}

// If_qstatsreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_qstatsreq
type If_qstatsreq struct {
	Ifqr_buf     unsafe.Pointer
	Ifqr_grp_idx U_int32_t
	Ifqr_len     unsafe.Pointer
	Ifqr_name    unsafe.Pointer
	Ifqr_slot    U_int32_t
}

// If_rxpoll_stats
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_rxpoll_stats
type If_rxpoll_stats struct {
	Ifi_poll_bytes         U_int64_t
	Ifi_poll_bytes_avg     U_int32_t
	Ifi_poll_bytes_hiwat   U_int32_t
	Ifi_poll_bytes_lowat   U_int32_t
	Ifi_poll_bytes_max     U_int32_t
	Ifi_poll_bytes_min     U_int32_t
	Ifi_poll_interval_time U_int64_t
	Ifi_poll_off_err       U_int32_t
	Ifi_poll_off_req       U_int32_t
	Ifi_poll_on_err        U_int32_t
	Ifi_poll_on_req        U_int32_t
	Ifi_poll_packets       U_int64_t
	Ifi_poll_packets_avg   U_int32_t
	Ifi_poll_packets_hiwat U_int32_t
	Ifi_poll_packets_limit U_int32_t
	Ifi_poll_packets_lowat U_int32_t
	Ifi_poll_packets_max   U_int32_t
	Ifi_poll_packets_min   U_int32_t
	Ifi_poll_wakeups_avg   U_int32_t
	Ifi_poll_wakeups_hiwat U_int32_t
	Ifi_poll_wakeups_lowat U_int32_t
}

// If_tcp_ecn_perf_stat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_tcp_ecn_perf_stat
type If_tcp_ecn_perf_stat struct {
	Oo_percent        U_int64_t
	Reorder_percent   U_int64_t
	Rst_drop          U_int64_t
	Rtt_avg           U_int64_t
	Rtt_var           U_int64_t
	Rxmit_drop        U_int64_t
	Rxmit_percent     U_int64_t
	Sack_episodes     U_int64_t
	Total_oopkts      U_int64_t
	Total_reorderpkts U_int64_t
	Total_rxmitpkts   U_int64_t
	Total_rxpkts      U_int64_t
	Total_txpkts      U_int64_t
}

// If_tcp_ecn_stat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_tcp_ecn_stat
type If_tcp_ecn_stat struct {
	Ecn_client_setup      U_int64_t
	Ecn_client_success    U_int64_t
	Ecn_conn_noplce       U_int64_t
	Ecn_conn_plce         U_int64_t
	Ecn_conn_plnoce       U_int64_t
	Ecn_conn_recv_ce      U_int64_t
	Ecn_conn_recv_ece     U_int64_t
	Ecn_fallback_ce       U_int64_t
	Ecn_fallback_droprst  U_int64_t
	Ecn_fallback_droprxmt U_int64_t
	Ecn_fallback_reorder  U_int64_t
	Ecn_fallback_synloss  U_int64_t
	Ecn_fallback_synrst   U_int64_t
	Ecn_off               If_tcp_ecn_perf_stat
	Ecn_off_conn          U_int64_t
	Ecn_on                If_tcp_ecn_perf_stat
	Ecn_peer_nosupport    U_int64_t
	Ecn_recv_ce           U_int64_t
	Ecn_recv_ece          U_int64_t
	Ecn_server_setup      U_int64_t
	Ecn_server_success    U_int64_t
	Ecn_syn_lost          U_int64_t
	Ecn_synack_lost       U_int64_t
	Ecn_total_conn        U_int64_t
	Timestamp             U_int64_t
}

// If_tdmreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_tdmreq
type If_tdmreq struct {
	Iftdm_len   U_int32_t
	Iftdm_name  unsafe.Pointer
	Iftdm_table unsafe.Pointer
}

// If_throttlereq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_throttlereq
type If_throttlereq struct {
	Ifthr_level U_int32_t
	Ifthr_name  unsafe.Pointer
}

// If_traffic_class
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_traffic_class
type If_traffic_class struct {
	Ifi_ibebytes   U_int64_t
	Ifi_ibepackets U_int64_t
	Ifi_ibkbytes   U_int64_t
	Ifi_ibkpackets U_int64_t
	Ifi_ipvbytes   U_int64_t
	Ifi_ipvpackets U_int64_t
	Ifi_ivibytes   U_int64_t
	Ifi_ivipackets U_int64_t
	Ifi_ivobytes   U_int64_t
	Ifi_ivopackets U_int64_t
	Ifi_obebytes   U_int64_t
	Ifi_obepackets U_int64_t
	Ifi_obkbytes   U_int64_t
	Ifi_obkpackets U_int64_t
	Ifi_opvbytes   U_int64_t
	Ifi_opvpackets U_int64_t
	Ifi_ovibytes   U_int64_t
	Ifi_ovipackets U_int64_t
	Ifi_ovobytes   U_int64_t
	Ifi_ovopackets U_int64_t
}

// If_wifi_status
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_wifi_status
type If_wifi_status struct {
	If_wifi_u unsafe.Pointer
}

// If_wifi_status_v1
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/if_wifi_status_v1
type If_wifi_status_v1 struct {
	Config_frequency       U_int32_t
	Config_multicast_rate  U_int32_t
	Dl_effective_bandwidth U_int32_t
	Dl_effective_latency   U_int32_t
	Dl_error_rate          U_int32_t
	Dl_max_bandwidth       U_int32_t
	Dl_max_latency         U_int32_t
	Dl_min_latency         U_int32_t
	Link_quality_metric    U_int32_t
	Reserved_1             U_int64_t
	Reserved_2             U_int64_t
	Reserved_3             U_int64_t
	Reserved_4             U_int64_t
	Scan_count             U_int32_t
	Scan_duration          U_int32_t
	Ul_bytes_lost          U_int32_t
	Ul_effective_bandwidth U_int32_t
	Ul_effective_latency   U_int32_t
	Ul_error_rate          U_int32_t
	Ul_max_bandwidth       U_int32_t
	Ul_max_latency         U_int32_t
	Ul_min_latency         U_int32_t
	Ul_retxt_level         U_int32_t
	Valid_bitmask          U_int32_t
}

// Ifa_msghdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifa_msghdr
type Ifa_msghdr struct {
	Ifam_addrs  unsafe.Pointer
	Ifam_index  unsafe.Pointer
	Ifam_metric unsafe.Pointer
	Ifam_msglen unsafe.Pointer
	Ifam_type   unsafe.Pointer
	Ifam_flags  unsafe.Pointer
}

// Ifaliasreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifaliasreq
type Ifaliasreq struct {
	Ifra_addr      unsafe.Pointer
	Ifra_mask      unsafe.Pointer
	Ifra_name      unsafe.Pointer
	Ifra_broadaddr unsafe.Pointer
}

// Ifdevmtu
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifdevmtu
type Ifdevmtu struct {
	Ifdm_current unsafe.Pointer
	Ifdm_min     unsafe.Pointer
	Ifdm_max     unsafe.Pointer
}

// Ifdrv
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifdrv
type Ifdrv struct {
	Ifd_data unsafe.Pointer
	Ifd_cmd  unsafe.Pointer
	Ifd_name unsafe.Pointer
	Ifd_len  uintptr
}

// Iff_filter
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iff_filter
type Iff_filter struct {
	Iff_output   unsafe.Pointer    // The filter function to handle outbound packets, may be NULL.
	Iff_detached unsafe.Pointer    // The filter function used to notify the filter that it has been detached.
	Iff_name     unsafe.Pointer    // A filter name used for debugging purposes.
	Iff_input    unsafe.Pointer    // The filter function to handle inbound packets, may be NULL.
	Iff_ioctl    unsafe.Pointer    // The filter function to handle interface ioctls, may be null.
	Iff_protocol Protocol_family_t // The protocol of the packets this filter is interested in. If you specify zero, packets from all protocols will be passed to the filter.
	Iff_event    unsafe.Pointer    // The filter function to handle interface events, may be null.
	Iff_cookie   unsafe.Pointer    // A kext defined cookie that will be passed to all filter functions.

}

// Ifkpi
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifkpi
type Ifkpi struct {
	Ifk_module_id unsafe.Pointer
	Ifk_data      unsafe.Pointer
}

// Ifma_msghdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifma_msghdr
type Ifma_msghdr struct {
	Ifmam_type   unsafe.Pointer
	Ifmam_msglen unsafe.Pointer
	Ifmam_addrs  unsafe.Pointer
	Ifmam_index  unsafe.Pointer
	Ifmam_flags  unsafe.Pointer
}

// Ifma_msghdr2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifma_msghdr2
type Ifma_msghdr2 struct {
	Ifmam_addrs    unsafe.Pointer
	Ifmam_index    U_short
	Ifmam_flags    unsafe.Pointer
	Ifmam_type     U_char
	Ifmam_msglen   U_short
	Ifmam_refcount unsafe.Pointer
}

// Ifmedia_description
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifmedia_description
type Ifmedia_description struct {
	Ifmt_word   int
	Ifmt_string *byte
}

// Ifmibdata
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifmibdata
type Ifmibdata struct {
	Ifmd_data       If_data64
	Ifmd_filler     unsafe.Pointer
	Ifmd_flags      unsafe.Pointer
	Ifmd_name       unsafe.Pointer
	Ifmd_pcount     unsafe.Pointer
	Ifmd_snd_drops  unsafe.Pointer
	Ifmd_snd_len    unsafe.Pointer
	Ifmd_snd_maxlen unsafe.Pointer
}

// Ifmibdata_supplemental
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifmibdata_supplemental
type Ifmibdata_supplemental struct {
	Ifmd_data_extended If_data_extended
	Ifmd_netif_stats   If_netif_stats
	Ifmd_packet_stats  If_packet_stats
	Ifmd_rxpoll_stats  If_rxpoll_stats
	Ifmd_traffic_class If_traffic_class
}

// Ifnet_attach_proto_param
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifnet_attach_proto_param
type Ifnet_attach_proto_param struct {
	Event       unsafe.Pointer // The function to be called for interface events.
	Pre_output  unsafe.Pointer // The function to be called for outbound packets.
	Input       unsafe.Pointer // The function to be called for inbound packets.
	Detached    unsafe.Pointer // The function to be called for handling the detach.
	Resolve     unsafe.Pointer
	Demux_count U_int32_t         // The number of entries in the demux_array array.
	Demux_array *Ifnet_demux_desc // An array of ifnet_demux_desc structures describing the protocol.
	Ioctl       unsafe.Pointer    // The function to be called for ioctls.
	Send_arp    unsafe.Pointer
}

// Ifnet_attach_proto_param_v2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifnet_attach_proto_param_v2
type Ifnet_attach_proto_param_v2 struct {
	Send_arp    unsafe.Pointer
	Detached    unsafe.Pointer
	Demux_count U_int32_t
	Input       unsafe.Pointer
	Pre_output  unsafe.Pointer
	Event       unsafe.Pointer
	Resolve     unsafe.Pointer
	Ioctl       unsafe.Pointer
	Demux_array *Ifnet_demux_desc
}

// Ifnet_demux_desc
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifnet_demux_desc
type Ifnet_demux_desc struct {
	Data    unsafe.Pointer // A pointer to an entry of type (i.e. pointer to 0x0800).
	Datalen U_int32_t      // The number of bytes of data used to describe the packet.
	Type    U_int32_t      // The type of identifier data (i.e. ETHER_DESC_ETYPE2)

}

// Ifnet_init_params
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifnet_init_params
type Ifnet_init_params struct {
	Set_bpf_tap    unsafe.Pointer // The function used to set the bpf_tap function.
	Detach         unsafe.Pointer // The function called to let the driver know the interface has been detached.
	Broadcast_addr unsafe.Pointer // The link-layer broadcast address for this interface.
	Uniqueid       unsafe.Pointer // An identifier unique to this instance of the interface.
	Add_proto      unsafe.Pointer // The function used to attach a protocol to this interface.
	Demux          unsafe.Pointer // The function used to determine the protocol family of an incoming packet.
	Broadcast_len  U_int32_t      // The length of the link-layer broadcast address.
	Check_multi    unsafe.Pointer
	Uniqueid_len   U_int32_t      // The length, in bytes, of the uniqueid.
	Name           unsafe.Pointer // The interface name (i.e. en).
	Ioctl          unsafe.Pointer // The function used to handle ioctls.
	Family         Ifnet_family_t // The interface family.
	Softc          unsafe.Pointer // Driver specific storage. This value can be retrieved from the ifnet using the ifnet_softc function.
	Del_proto      unsafe.Pointer // The function used to remove a protocol from this interface.
	Type           U_int32_t      // The interface type (see sys/if_types.h). Must be less than 256. For new types, use IFT_OTHER.
	Framer         unsafe.Pointer // The function used to frame outbound packets, may be NULL.
	Event          unsafe.Pointer // The function to notify the interface of various interface specific kernel events.
	Output         unsafe.Pointer // The output function for the interface. Every packet the stack attempts to send through this interface will go out through this function.
	Unit           U_int32_t      // The interface unit number (en0's unit number is 0).

}

// Ifnet_interface_advisory
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifnet_interface_advisory
type Ifnet_interface_advisory struct {
}

// Ifnet_interface_advisory_capacity
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifnet_interface_advisory_capacity
type Ifnet_interface_advisory_capacity struct {
	Average_delay      unsafe.Pointer
	Average_throughput unsafe.Pointer
}

// Ifnet_interface_advisory_cell_context
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifnet_interface_advisory_cell_context
type Ifnet_interface_advisory_cell_context struct {
	Bandwidth_limitation_indication uint8
	Cdrx_cycle                      uint16
	Cdrx_state                      uint8
	Downlink_bler                   uint8
	Estimated_outage_period         uint16
	Outage_state                    uint8
	Radio_access_technology         uint8
	Reference_signal_level          int16
	Signal_level                    int16
	Signal_quality                  int8
	Uplink_bler                     uint8
}

// Ifnet_interface_advisory_header
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifnet_interface_advisory_header
type Ifnet_interface_advisory_header struct {
	Direction         IfInterfaceAdvisoryDirection
	Interface_type    IfInterfaceAdvisoryInterfaceType
	Version           Ifnet_interface_advisory_version
	Notification_type unsafe.Pointer
}

// Ifnet_interface_advisory_wifi_context
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifnet_interface_advisory_wifi_context
type Ifnet_interface_advisory_wifi_context struct {
	Bt_coex                       uint8
	Estimated_intermittent_period uint16
	Frequency_band                IfInterfaceAdvisoryFreqBand
	Intermittent_state            uint8
	Quality_score_channel         uint8
	Quality_score_delay           uint8
	Quality_score_loss            uint8
	Radio_coex                    uint8
	Single_outage_period          uint16
	Wifi_observed_tx_bitrate      unsafe.Pointer
	Wlan_duty_cycle               uint16
}

// Ifnet_ip_addr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifnet_ip_addr
type Ifnet_ip_addr struct {
}

// Ifnet_stat_increment_param
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifnet_stat_increment_param
type Ifnet_stat_increment_param struct {
	Bytes_out   U_int32_t // The number of bytes transmitted.
	Collisions  U_int32_t // The number of collisions seen by this interface.
	Errors_out  U_int32_t // The number of transmission errors.
	Bytes_in    U_int32_t // The number of bytes received.
	Dropped     U_int32_t // The number of packets dropped.
	Packets_in  U_int32_t // The number of packets received.
	Errors_in   U_int32_t // The number of receive errors.
	Packets_out U_int32_t // The number of packets transmitted.

}

// Ifnet_stats_param
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifnet_stats_param
type Ifnet_stats_param struct {
	Bytes_in       U_int64_t // The number of bytes received.
	Errors_out     U_int64_t // The number of transmission errors.
	Errors_in      U_int64_t // The number of receive errors.
	Packets_in     U_int64_t // The number of packets received.
	Multicasts_out U_int64_t
	Multicasts_in  U_int64_t
	Dropped        U_int64_t // The number of packets dropped.
	Collisions     U_int64_t // The number of collisions seen by this interface.
	Bytes_out      U_int64_t // The number of bytes transmitted.
	No_protocol    U_int64_t
	Packets_out    U_int64_t // The number of packets transmitted.

}

// Ifnet_stats_per_flow
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifnet_stats_per_flow
type Ifnet_stats_per_flow struct {
	Bk_txpackets U_int64_t
	Bw_rcvbw_max U_int32_t
	Bw_sndbw_max U_int32_t
	Connreset    U_int16_t
	Conntimeout  U_int16_t
	Srtt         U_int32_t
}

// Ifnet_traffic_descriptor_common
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifnet_traffic_descriptor_common
type Ifnet_traffic_descriptor_common struct {
	Itd_flags unsafe.Pointer
	Itd_len   uint16
}

// Ifnet_traffic_descriptor_inet
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifnet_traffic_descriptor_inet
type Ifnet_traffic_descriptor_inet struct {
	Inet_common Ifnet_traffic_descriptor_common
	Inet_ipver  uint8
	Inet_laddr  Ifnet_ip_addr
	Inet_lport  uint16
	Inet_mask   uint8
	Inet_proto  uint8
	Inet_raddr  Ifnet_ip_addr
	Inet_rport  uint16
}

// Ifnet_traffic_rule_action
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifnet_traffic_rule_action
type Ifnet_traffic_rule_action struct {
	Ra_type uint8
}

// Ifnet_traffic_rule_action_steer
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifnet_traffic_rule_action_steer
type Ifnet_traffic_rule_action_steer struct {
	Ras_common  Ifnet_traffic_rule_action
	Ras_qset_id unsafe.Pointer
}

// Ifqueue
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifqueue
type Ifqueue struct {
	Ifq_tail   unsafe.Pointer
	Ifq_drops  unsafe.Pointer
	Ifq_maxlen unsafe.Pointer
	Ifq_head   unsafe.Pointer
	Ifq_len    unsafe.Pointer
}

// Ifreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifreq
type Ifreq struct {
	Ifr_ifru unsafe.Pointer
	Ifr_name unsafe.Pointer
}

// Ifs_iso_8802_3
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifs_iso_8802_3
type Ifs_iso_8802_3 struct {
	Dot3Compliance           U_int32_t
	Dot3StatsAlignmentErrors U_int32_t
}

// Ifstat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ifstat
type Ifstat struct {
	Ifs_name unsafe.Pointer
	Ascii    unsafe.Pointer
}

// Igmp
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/igmp
type Igmp struct {
	Igmp_cksum U_short
	Igmp_code  U_char
	Igmp_group In_addr
	Igmp_type  U_char
}

// Igmp_grouprec
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/igmp_grouprec
type Igmp_grouprec struct {
	Ig_datalen U_char
}

// Igmp_report
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/igmp_report
type Igmp_report struct {
	Ir_cksum   U_short
	Ir_numgrps U_short
	Ir_rsv1    U_char
	Ir_rsv2    U_short
	Ir_type    U_char
}

// Igmpstat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/igmpstat
type Igmpstat struct {
	Igps_rcv_badqueries U_int
	Igps_rcv_badreports U_int
	Igps_rcv_badsum     U_int
	Igps_rcv_ourreports U_int
}

// Igmpstat_v3
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/igmpstat_v3
type Igmpstat_v3 struct {
	Igps_drop_gsr_queries  unsafe.Pointer
	Igps_len               unsafe.Pointer
	Igps_rcv_badqueries    unsafe.Pointer
	Igps_rcv_badreports    unsafe.Pointer
	Igps_rcv_badsum        unsafe.Pointer
	Igps_rcv_badttl        unsafe.Pointer
	Igps_rcv_gen_queries   unsafe.Pointer
	Igps_rcv_group_queries unsafe.Pointer
	Igps_rcv_gsr_queries   unsafe.Pointer
	Igps_rcv_nora          unsafe.Pointer
	Igps_rcv_ourreports    unsafe.Pointer
	Igps_rcv_reports       unsafe.Pointer
	Igps_rcv_tooshort      unsafe.Pointer
	Igps_rcv_total         unsafe.Pointer
	Igps_rcv_v1v2_queries  unsafe.Pointer
	Igps_rcv_v3_queries    unsafe.Pointer
	Igps_snd_reports       unsafe.Pointer
}

// Igmpv3
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/igmpv3
type Igmpv3 struct {
	Igmp_cksum  U_short
	Igmp_code   U_char
	Igmp_group  In_addr
	Igmp_misc   U_char
	Igmp_numsrc U_short
	Igmp_qqi    U_char
	Igmp_type   U_char
}

// Image_params
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/image_params
type Image_params struct {
	Ip_endargv                    unsafe.Pointer
	Ip_startargv                  unsafe.Pointer
	Ip_user_envv                  User_addr_t
	Ip_arch_offset                User_size_t
	Ip_csflags                    unsafe.Pointer
	Ip_interp_buffer              unsafe.Pointer
	Ip_applec                     unsafe.Pointer
	Ip_origcputype                int32
	Ip_user_argv                  User_addr_t
	Ip_interp_sugid_fd            unsafe.Pointer
	Ip_origcpusubtype             int32
	Ip_envc                       unsafe.Pointer
	Ip_scriptlabelp               unsafe.Pointer
	Ip_vp                         unsafe.Pointer
	Ip_vdata                      unsafe.Pointer
	Ip_px_sa                      unsafe.Pointer
	Ip_ndp                        unsafe.Pointer
	Ip_execlabelp                 unsafe.Pointer
	Ip_origvattr                  *Vnode_attr
	Ip_user_fname                 User_addr_t
	Ip_new_thread                 Thread_t
	Ip_strings                    unsafe.Pointer
	Ip_px_sfa                     unsafe.Pointer
	Ip_flags                      unsafe.Pointer
	Ip_mac_return                 unsafe.Pointer
	Ip_vfs_context                unsafe.Pointer
	Ip_argc                       unsafe.Pointer
	Ip_strspace                   unsafe.Pointer
	Ip_strendp                    unsafe.Pointer
	Ip_scriptvp                   unsafe.Pointer
	Ip_px_persona                 unsafe.Pointer
	Ip_argspace                   unsafe.Pointer
	Ip_endenvv                    unsafe.Pointer
	Ip_px_spa                     unsafe.Pointer
	Ip_px_smpx                    unsafe.Pointer
	Ip_arch_size                  User_size_t
	Ip_vattr                      *Vnode_attr
	Ip_seg                        unsafe.Pointer
	Ip_cs_error                   unsafe.Pointer
	Ip_dyld_fsobjid               unsafe.Pointer
	Ip_dyld_fsid                  unsafe.Pointer
	Ip_px_pcred_info              unsafe.Pointer
	Ip_simulator_binary           unsafe.Pointer
	Ip_inherited_jop_pid          unsafe.Pointer
	Ip_inherited_shared_region_id unsafe.Pointer
	Ip_subsystem_root_path        unsafe.Pointer
	Ip_free_map                   Vm_map_t
}

// In6_addrlifetime
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/in6_addrlifetime
type In6_addrlifetime struct {
	Ia6t_expire    int64
	Ia6t_pltime    U_int32_t
	Ia6t_preferred int64
	Ia6t_vltime    U_int32_t
}

// In6_addrpolicy
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/in6_addrpolicy
type In6_addrpolicy struct {
	Addr Sockaddr_in6
}

// In6_aliasreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/in6_aliasreq
type In6_aliasreq struct {
	Ifra_addr       Sockaddr_in6
	Ifra_broadaddr  Sockaddr_in6
	Ifra_flags      unsafe.Pointer
	Ifra_lifetime   In6_addrlifetime
	Ifra_name       unsafe.Pointer
	Ifra_prefixmask Sockaddr_in6
}

// In6_ifreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/in6_ifreq
type In6_ifreq struct {
	Ifr_ifru unsafe.Pointer
	Ifr_name unsafe.Pointer
}

// In6_ifstat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/in6_ifstat
type In6_ifstat struct {
	Ifs6_addr_expiry_cnt   U_quad_t
	Ifs6_atmfrag_rcvd      U_quad_t
	Ifs6_cantfoward_icmp6  U_quad_t
	Ifs6_defrtr_expiry_cnt U_quad_t
	Ifs6_in_addrerr        U_quad_t
	Ifs6_in_deliver        U_quad_t
	Ifs6_in_discard        U_quad_t
	Ifs6_in_hdrerr         U_quad_t
	Ifs6_in_mcast          U_quad_t
	Ifs6_in_noroute        U_quad_t
	Ifs6_in_protounknown   U_quad_t
	Ifs6_in_receive        U_quad_t
	Ifs6_in_toobig         U_quad_t
	Ifs6_in_truncated      U_quad_t
	Ifs6_out_discard       U_quad_t
	Ifs6_out_forward       U_quad_t
	Ifs6_out_fragcreat     U_quad_t
	Ifs6_out_fragfail      U_quad_t
	Ifs6_out_fragok        U_quad_t
	Ifs6_out_mcast         U_quad_t
	Ifs6_out_request       U_quad_t
	Ifs6_pfx_expiry_cnt    U_quad_t
	Ifs6_reass_fail        U_quad_t
	Ifs6_reass_ok          U_quad_t
	Ifs6_reass_reqd        U_quad_t
}

// In6_pktinfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/in6_pktinfo
type In6_pktinfo struct {
	Ipi6_ifindex unsafe.Pointer
	Ipi6_addr    unsafe.Pointer
}

// In6_prefixreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/in6_prefixreq
type In6_prefixreq struct {
	Ipr_name   unsafe.Pointer
	Ipr_origin U_char
}

// In6_prflags
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/in6_prflags
type In6_prflags struct {
	Prf_ra        unsafe.Pointer
	Prf_reserved1 U_char
	Prf_reserved2 U_short
	Prf_reserved3 U_char
}

// In6_rrenumreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/in6_rrenumreq
type In6_rrenumreq struct {
	Irr_m_len    U_char
	Irr_pltime   U_int32_t
	Irr_u_uselen U_char
}

// In_addr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/in_addr
type In_addr struct {
	S_addr uint32
}

// In_addr_4in6
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/in_addr_4in6
type In_addr_4in6 struct {
	Ia46_addr4 In_addr
	Ia46_pad32 U_int32_t
}

// In_aliasreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/in_aliasreq
type In_aliasreq struct {
	Ifra_addr      Sockaddr_in
	Ifra_broadaddr Sockaddr_in
	Ifra_mask      Sockaddr_in
	Ifra_name      unsafe.Pointer
}

// In_pktinfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/in_pktinfo
type In_pktinfo struct {
	Ipi_ifindex  unsafe.Pointer
	Ipi_addr     In_addr
	Ipi_spec_dst In_addr
}

// Info_tuple
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/info_tuple
type Info_tuple struct {
	Itpl_proto      U_int8_t
	Itpl_remoteaddr unsafe.Pointer
}

// Inpcb
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/inpcb
type Inpcb struct {
	Hash_element    unsafe.Pointer
	Inp_depend4     unsafe.Pointer
	Inp_depend6     unsafe.Pointer
	Inp_dependfaddr unsafe.Pointer
	Inp_dependladdr unsafe.Pointer
	Inp_dependroute unsafe.Pointer
	Inp_flags       unsafe.Pointer
	Inp_flow        U_int32_t
	Inp_fport       U_short
	Inp_gencnt      Inp_gen_t
	Inp_hash        unsafe.Pointer
	Inp_ip_p        U_char
	Inp_ip_ttl      U_char
	Inp_list        unsafe.Pointer
	Inp_lport       U_short
	Inp_pcbinfo     U_int32_t
	Inp_phd         U_int32_t
	Inp_portlist    unsafe.Pointer
	Inp_ppcb        U_int32_t
	Inp_saved_ppcb  U_int32_t
	Inp_socket      U_int32_t
	Inp_sp          U_int32_t
	Inp_vflag       U_char
	Nat_cookie      U_int32_t
	Nat_owner       U_char
	Reserved        U_int32_t
	Reserved1       In_addr
	Reserved2       In_addr
}

// Inpcb64_list_entry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/inpcb64_list_entry
type Inpcb64_list_entry struct {
	Le_next U_int64_t
	Le_prev U_int64_t
}

// Instrs_cycles_snapshot
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/instrs_cycles_snapshot
type Instrs_cycles_snapshot struct {
	Ics_cycles       unsafe.Pointer
	Ics_instructions unsafe.Pointer
}

// Instrs_cycles_snapshot_v2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/instrs_cycles_snapshot_v2
type Instrs_cycles_snapshot_v2 struct {
	Ics_cycles         unsafe.Pointer
	Ics_instructions   unsafe.Pointer
	Ics_p_cycles       unsafe.Pointer
	Ics_p_instructions unsafe.Pointer
}

// Internal_state
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/internal_state
type Internal_state struct {
	Dummy unsafe.Pointer
}

// Io_stat_entry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/io_stat_entry
type Io_stat_entry struct {
	Count uint64
	Size  uint64
}

// Io_stats_snapshot
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/io_stats_snapshot
type Io_stats_snapshot struct {
	Ss_non_paging_count  unsafe.Pointer
	Ss_metadata_count    unsafe.Pointer
	Ss_disk_reads_size   unsafe.Pointer
	Ss_io_priority_size  unsafe.Pointer
	Ss_paging_count      unsafe.Pointer
	Ss_disk_writes_size  unsafe.Pointer
	Ss_disk_reads_count  unsafe.Pointer
	Ss_data_count        unsafe.Pointer
	Ss_data_size         unsafe.Pointer
	Ss_metadata_size     unsafe.Pointer
	Ss_disk_writes_count unsafe.Pointer
	Ss_paging_size       unsafe.Pointer
	Ss_io_priority_count unsafe.Pointer
	Ss_non_paging_size   unsafe.Pointer
}

// Iocompressionstats_notification_subsystem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iocompressionstats_notification_subsystem-oj5
type Iocompressionstats_notification_subsystem struct {
	End      int32
	Maxsize  unsafe.Pointer
	Reserved Vm_address_t
	Routine  unsafe.Pointer
	Server   unsafe.Pointer
	Start    int32
}

// Iocs_store_buffer_entry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iocs_store_buffer_entry
type Iocs_store_buffer_entry struct {
	Iocs      unsafe.Pointer
	Path_name unsafe.Pointer
}

// Iovec
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/iovec
type Iovec struct {
	Iov_len  uintptr
	Iov_base unsafe.Pointer
}

// Ip
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip
type Ip struct {
	Ip_tos U_char
	Ip_off U_short
	Ip_p   U_char
	Ip_hl  U_int
	Ip_dst In_addr
	Ip_ttl U_char
	Ip_src In_addr
	Ip_len U_short
	Ip_id  U_short
	Ip_v   U_int
	Ip_sum U_short
}

// Ip6_dest
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip6_dest
type Ip6_dest struct {
	Ip6d_len U_int8_t
	Ip6d_nxt U_int8_t
}

// Ip6_ext
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip6_ext
type Ip6_ext struct {
	Ip6e_len U_int8_t
	Ip6e_nxt U_int8_t
}

// Ip6_frag
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip6_frag
type Ip6_frag struct {
	Ip6f_ident    U_int32_t
	Ip6f_nxt      U_int8_t
	Ip6f_offlg    U_int16_t
	Ip6f_reserved U_int8_t
}

// Ip6_hbh
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip6_hbh
type Ip6_hbh struct {
	Ip6h_len U_int8_t
	Ip6h_nxt U_int8_t
}

// Ip6_hdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip6_hdr
type Ip6_hdr struct {
	Ip6_ctlun unsafe.Pointer
	Ip6_dst   unsafe.Pointer
	Ip6_src   unsafe.Pointer
}

// Ip6_mtuinfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip6_mtuinfo
type Ip6_mtuinfo struct {
	Ip6m_addr Sockaddr_in6
	Ip6m_mtu  unsafe.Pointer
}

// Ip6_opt
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip6_opt
type Ip6_opt struct {
	Ip6o_len  U_int8_t
	Ip6o_type U_int8_t
}

// Ip6_opt_jumbo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip6_opt_jumbo
type Ip6_opt_jumbo struct {
	Ip6oj_jumbo_len U_int8_t
}

// Ip6_opt_nsap
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip6_opt_nsap
type Ip6_opt_nsap struct {
	Ip6on_dst_nsap_len U_int8_t
	Ip6on_len          U_int8_t
	Ip6on_src_nsap_len U_int8_t
	Ip6on_type         U_int8_t
}

// Ip6_opt_router
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip6_opt_router
type Ip6_opt_router struct {
	Ip6or_len   U_int8_t
	Ip6or_type  U_int8_t
	Ip6or_value U_int8_t
}

// Ip6_opt_tunnel
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip6_opt_tunnel
type Ip6_opt_tunnel struct {
	Ip6ot_encap_limit U_int8_t
	Ip6ot_len         U_int8_t
}

// Ip6_rthdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip6_rthdr
type Ip6_rthdr struct {
	Ip6r_len     U_int8_t
	Ip6r_nxt     U_int8_t
	Ip6r_segleft U_int8_t
	Ip6r_type    U_int8_t
}

// Ip6_rthdr0
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip6_rthdr0
type Ip6_rthdr0 struct {
	Ip6r0_len      U_int8_t
	Ip6r0_nxt      U_int8_t
	Ip6r0_reserved U_int32_t
	Ip6r0_segleft  U_int8_t
	Ip6r0_type     U_int8_t
}

// Ip_linklocal_stat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip_linklocal_stat
type Ip_linklocal_stat struct {
	Iplls_in_badttl  U_int32_t
	Iplls_in_total   U_int32_t
	Iplls_out_badttl U_int32_t
	Iplls_out_total  U_int32_t
}

// Ip_mreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip_mreq
type Ip_mreq struct {
	Imr_interface In_addr
	Imr_multiaddr In_addr
}

// Ip_mreq_source
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip_mreq_source
type Ip_mreq_source struct {
	Imr_interface  In_addr
	Imr_multiaddr  In_addr
	Imr_sourceaddr In_addr
}

// Ip_mreqn
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip_mreqn
type Ip_mreqn struct {
	Imr_multiaddr In_addr
	Imr_address   In_addr
	Imr_ifindex   unsafe.Pointer
}

// Ip_opts
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip_opts
type Ip_opts struct {
	Ip_opts unsafe.Pointer
	Ip_dst  In_addr
}

// Ip_timestamp
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ip_timestamp
type Ip_timestamp struct {
	Ipt_oflw      U_int
	Ipt_len       U_char
	Ipt_timestamp unsafe.Pointer
	Ipt_flg       U_int
	Ipt_ptr       U_char
	Ipt_code      U_char
}

// Ipc_perm
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ipc_perm
type Ipc_perm struct {
	Cuid uint32
	Mode uint16
	Uid  uint32
	Gid  uint32
	Cgid uint32
}

// Ipcomp
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ipcomp
type Ipcomp struct {
	Comp_nxt   U_int8_t
	Comp_flags U_int8_t
	Comp_cpi   U_int16_t
}

// Ipovly
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ipovly
type Ipovly struct {
	Ih_dst In_addr
	Ih_len U_short
	Ih_pr  U_char
	Ih_src In_addr
	Ih_x1  U_char
}

// Ipsec_stats_param
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ipsec_stats_param
type Ipsec_stats_param struct {
	Utsp_bytes   U_int64_t
	Utsp_errors  U_int64_t
	Utsp_packets U_int64_t
}

// Ipsec_wake_pkt_event_data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ipsec_wake_pkt_event_data
type Ipsec_wake_pkt_event_data struct {
	Wake_uuid Uuid_string_t
}

// Ipsec_wake_pkt_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ipsec_wake_pkt_info
type Ipsec_wake_pkt_info struct {
	Wake_pkt     U_int8_t
	Wake_pkt_len U_int16_t
	Wake_pkt_seq U_int32_t
	Wake_pkt_spi U_int32_t
	Wake_uuid    Uuid_string_t
}

// Ipsecstat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ipsecstat
type Ipsecstat struct {
	In_ahauthfail  U_quad_t
	In_success     U_quad_t
	In_ahauthsucc  U_quad_t
	In_espreplay   U_quad_t
	In_esphist     U_quad_t
	In_espauthsucc U_quad_t
	In_nosa        U_quad_t
	In_inval       U_quad_t
	In_ahhist      U_quad_t
	In_ahreplay    U_quad_t
	In_comphist    U_quad_t
	Out_inval      U_quad_t
	Out_comphist   U_quad_t
	In_nomem       U_quad_t
	Out_ahhist     U_quad_t
	Out_nosa       U_quad_t
	Out_esphist    U_quad_t
	Out_nomem      U_quad_t
	In_polvio      U_quad_t
	Out_success    U_quad_t
	In_espauthfail U_quad_t
	In_badspi      U_quad_t
	Out_noroute    U_quad_t
	Out_polvio     U_quad_t
}

// Ipstat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ipstat
type Ipstat struct {
	Ips_odropped     U_int32_t
	Ips_rawout       U_int32_t
	Ips_redirectsent U_int32_t
	Ips_toosmall     U_int32_t
}

// Ipv6_mreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ipv6_mreq
type Ipv6_mreq struct {
	Ipv6mr_interface unsafe.Pointer
	Ipv6mr_multiaddr unsafe.Pointer
}

// Ipv6_prefix
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ipv6_prefix
type Ipv6_prefix struct {
	Ipv6_prefix unsafe.Pointer
	Prefix_len  unsafe.Pointer
}

// Itimerval
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/itimerval
type Itimerval struct {
	It_interval Timeval
	It_value    Timeval
}

// Jetsam_coalition_snapshot
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/jetsam_coalition_snapshot
type Jetsam_coalition_snapshot struct {
	Jcs_flags                unsafe.Pointer
	Jcs_thread_group         unsafe.Pointer
	Jcs_leader_task_uniqueid unsafe.Pointer
	Jcs_id                   unsafe.Pointer
}

// Kauth_cache_sizes
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kauth_cache_sizes
type Kauth_cache_sizes struct {
	Kcs_id_size    U_int32_t
	Kcs_group_size U_int32_t
}

// Kauth_identity_extlookup
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kauth_identity_extlookup
type Kauth_identity_extlookup struct {
	El_uguid           unsafe.Pointer
	El_seqno           U_int32_t
	El_usid            Ntsid_t
	El_info_reserved_1 U_int32_t
	El_sup_grp_cnt     U_int32_t
	El_gguid_valid     U_int32_t
	El_result          U_int32_t
	El_member_valid    U_int32_t
	El_extend          U_int64_t
	El_gsid            Ntsid_t
	El_gguid           unsafe.Pointer
	El_info_pid        unsafe.Pointer
	El_usid_valid      U_int32_t
	El_gid             uint32
	El_uid             uint32
	El_uguid_valid     U_int32_t
	El_flags           U_int32_t
	El_gsid_valid      U_int32_t
	El_sup_groups      uint32
}

// Kcdata_type_definition
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kcdata_type_definition
type Kcdata_type_definition struct {
	Kct_type_identifier unsafe.Pointer
	Kct_num_elements    unsafe.Pointer
	Kct_name            unsafe.Pointer
	Kct_elements        unsafe.Pointer
}

// Kern_ctl_reg
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kern_ctl_reg
type Kern_ctl_reg struct {
	Ctl_name       unsafe.Pointer // A Bundle ID string of up to MAX_KCTL_NAME bytes (including the ending zero). This string should not be empty.
	Ctl_id         U_int32_t      // The control ID may be dynamically assigned or it can be a 32-bit creator code assigned by DTS. For a DTS assigned creator code the CTL_FLAG_REG_ID_UNIT flag must be set. For a dynamically assigned control ID, do not set the CTL_FLAG_REG_ID_UNIT flag. The value of the dynamically assigned control ID is set to this field when the registration succeeds.
	Ctl_unit       U_int32_t      // A separate unit number to register multiple units that share the same control ID with DTS assigned creator code when the CTL_FLAG_REG_ID_UNIT flag is set. This field is ignored for a dynamically assigned control ID.
	Ctl_flags      U_int32_t      // CTL_FLAG_PRIVILEGED and/or CTL_FLAG_REG_ID_UNIT.
	Ctl_sendsize   U_int32_t      // Override the default send size. If set to zero, the default send size will be used, and this default value is set to this field to be retrieved by the caller.
	Ctl_recvsize   U_int32_t      // Override the default receive size. If set to zero, the default receive size will be used, and this default value is set to this field to be retrieved by the caller.
	Ctl_connect    unsafe.Pointer // Specify the function to be called whenever a client connects to the kernel control. This field must be specified.
	Ctl_disconnect unsafe.Pointer // Specify a function to be called whenever a client disconnects from the kernel control.
	Ctl_send       unsafe.Pointer // Specify a function to handle data send from the client to the kernel control.
	Ctl_setopt     unsafe.Pointer // Specify a function to handle set socket option operations for the kernel control.
	Ctl_getopt     unsafe.Pointer // Specify a function to handle get socket option operations for the kernel control.

}

// Kern_event_msg
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kern_event_msg
type Kern_event_msg struct {
	Total_size   U_int32_t // Total size of the kernel event message including the header.
	Event_code   U_int32_t // The event code.
	Vendor_code  U_int32_t // The vendor code indicates which vendor generated the kernel event. This gives every vendor a unique set of classes and subclasses to use. Use the SIOCGKEVVENDOR ioctl to look up vendor codes for vendors other than Apple. Apple uses KEV_VENDOR_APPLE.
	Kev_subclass U_int32_t // The subclass of the kernel event.
	Kev_class    U_int32_t // The class of the kernel event.
	Id           U_int32_t // Monotonically increasing value.
	Event_data   U_int32_t // Any additional data about this event. Format will depend on the vendor_code, kev_class, kev_subclass, and event_code. The length of the event_data can be determined using total_size - KEV_MSG_HEADER_SIZE.

}

// Kernel_triage_info_v1
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kernel_triage_info_v1
type Kernel_triage_info_v1 struct {
	Triage_string1 unsafe.Pointer
	Triage_string2 unsafe.Pointer
	Triage_string3 unsafe.Pointer
	Triage_string4 unsafe.Pointer
	Triage_string5 unsafe.Pointer
}

// Kev_d_vectors
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_d_vectors
type Kev_d_vectors struct {
	Data_length U_int32_t // The length of data.

}

// Kev_dl_issues
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_dl_issues
type Kev_dl_issues struct {
	Info      U_int8_t
	Link_data Net_event_data
	Modid     U_int8_t
	Timestamp U_int64_t
}

// Kev_dl_link_quality_metric_data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_dl_link_quality_metric_data
type Kev_dl_link_quality_metric_data struct {
	Link_data           Net_event_data
	Link_quality_metric unsafe.Pointer
}

// Kev_dl_low_power_mode
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_dl_low_power_mode
type Kev_dl_low_power_mode struct {
	Link_data       Net_event_data
	Low_power_event unsafe.Pointer
}

// Kev_dl_node_absence
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_dl_node_absence
type Kev_dl_node_absence struct {
	Link_data         Net_event_data
	Sdl_node_address  Sockaddr_dl
	Sin6_node_address Sockaddr_in6
}

// Kev_dl_node_presence
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_dl_node_presence
type Kev_dl_node_presence struct {
	Link_data             Net_event_data
	Link_quality_metric   unsafe.Pointer
	Node_proximity_metric unsafe.Pointer
	Node_service_info     U_int8_t
	Rssi                  unsafe.Pointer
	Sdl_node_address      Sockaddr_dl
	Sin6_node_address     Sockaddr_in6
}

// Kev_dl_proto_data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_dl_proto_data
type Kev_dl_proto_data struct {
	Proto_family          U_int32_t
	Proto_remaining_count U_int32_t
	Link_data             Net_event_data
}

// Kev_dl_rrc_state
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_dl_rrc_state
type Kev_dl_rrc_state struct {
	Link_data Net_event_data
	Rrc_state U_int32_t
}

// Kev_in6_addrlifetime
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_in6_addrlifetime
type Kev_in6_addrlifetime struct {
	Ia6t_expire    U_int32_t
	Ia6t_pltime    U_int32_t
	Ia6t_preferred U_int32_t
	Ia6t_vltime    U_int32_t
}

// Kev_in6_data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_in6_data
type Kev_in6_data struct {
	Ia6_flags     U_int32_t
	Ia_addr       Sockaddr_in6
	Ia_dstaddr    Sockaddr_in6
	Ia_lifetime   Kev_in6_addrlifetime
	Ia_mac        uint8
	Ia_net        Sockaddr_in6
	Ia_plen       U_int32_t
	Ia_prefixmask Sockaddr_in6
	Link_data     Net_event_data
}

// Kev_in_arpalive
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_in_arpalive
type Kev_in_arpalive struct {
	Link_data Net_event_data
}

// Kev_in_arpfailure
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_in_arpfailure
type Kev_in_arpfailure struct {
	Link_data Net_event_data
}

// Kev_in_collision
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_in_collision
type Kev_in_collision struct {
	Hw_addr   U_char
	Hw_len    U_char
	Ia_ipaddr In_addr
	Link_data Net_event_data
}

// Kev_in_data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_in_data
type Kev_in_data struct {
	Ia_addr         In_addr
	Ia_dstaddr      In_addr
	Ia_net          U_int32_t
	Ia_netbroadcast In_addr
	Ia_netmask      U_int32_t
	Ia_subnet       U_int32_t
	Ia_subnetmask   U_int32_t
	Link_data       Net_event_data
}

// Kev_in_portinuse
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_in_portinuse
type Kev_in_portinuse struct {
	Port     U_int16_t
	Req_pid  U_int32_t
	Reserved U_int32_t
}

// Kev_msg
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_msg
type Kev_msg struct {
	Kev_class    U_int32_t     // The event's class.
	Vendor_code  U_int32_t     // The vendor code assigned by kev_vendor_code_find.
	Kev_subclass U_int32_t     // The event's subclass.
	Event_code   U_int32_t     // The event's code.
	Dv           Kev_d_vectors // An array of vectors describing additional data to be appended to the kernel event.

}

// Kev_netevent_apnfallbk_data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_netevent_apnfallbk_data
type Kev_netevent_apnfallbk_data struct {
	Epid  int32
	Euuid [16]byte
}

// Kev_netevent_clat46_data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_netevent_clat46_data
type Kev_netevent_clat46_data struct {
	Clat46_event_code In6_clat46_evhdlr_code_t
	Epid              int32
	Euuid             [16]byte
}

// Kev_request
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_request
type Kev_request struct {
	Vendor_code  U_int32_t // All kernel events that don't match this vendor code will be ignored. KEV_ANY_VENDOR can be used to receive kernel events with any vendor code.
	Kev_class    U_int32_t // All kernel events that don't match this class will be ignored. KEV_ANY_CLASS can be used to receive kernel events with any class.
	Kev_subclass U_int32_t // All kernel events that don't match this subclass will be ignored. KEV_ANY_SUBCLASS can be used to receive kernel events with any subclass.

}

// Kev_vendor_code
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kev_vendor_code
type Kev_vendor_code struct {
	Vendor_string unsafe.Pointer // A bundle style identifier.
	Vendor_code   U_int32_t      // After making the SIOCGKEVVENDOR ioctl call, this will be filled in with the vendor code if there is one.

}

// Kevent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kevent
type Kevent struct {
	Data int
}

// Kevent64_s
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kevent64_s
type Kevent64_s struct {
	Filter int16
	Ext    unsafe.Pointer
	Data   unsafe.Pointer
	Flags  uint16
	Ident  unsafe.Pointer
	Udata  unsafe.Pointer
	Fflags unsafe.Pointer
}

// Kpc_config_remote
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kpc_config_remote
type Kpc_config_remote struct {
	Classes  unsafe.Pointer
	Configv  *Kpc_config_t
	Pmc_mask unsafe.Pointer
	Secure   unsafe.Pointer
}

// Kpc_get_counters_remote
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kpc_get_counters_remote
type Kpc_get_counters_remote struct {
	Buf         unsafe.Pointer
	Nb_counters unsafe.Pointer
	Buf_stride  unsafe.Pointer
	Classes     unsafe.Pointer
}

// Kpc_running_remote
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/kpc_running_remote
type Kpc_running_remote struct {
	Cfg_state_mask  unsafe.Pointer
	Cfg_target_mask unsafe.Pointer
}

// Ledger_entry_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ledger_entry_info
type Ledger_entry_info struct {
	Lei_balance       unsafe.Pointer
	Lei_credit        unsafe.Pointer
	Lei_refill_period unsafe.Pointer
	Lei_limit         unsafe.Pointer
	Lei_last_refill   unsafe.Pointer
	Lei_debit         unsafe.Pointer
}

// Ledger_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ledger_info
type Ledger_info struct {
	Li_id unsafe.Pointer
}

// Ledger_limit_args
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ledger_limit_args
type Ledger_limit_args struct {
	Lla_limit         unsafe.Pointer
	Lla_refill_period unsafe.Pointer
	Lla_name          unsafe.Pointer
}

// Ledger_template_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ledger_template_info
type Ledger_template_info struct {
	Lti_name  unsafe.Pointer
	Lti_units unsafe.Pointer
	Lti_group unsafe.Pointer
}

// Linger
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/linger
type Linger struct {
	L_linger unsafe.Pointer
	L_onoff  unsafe.Pointer
}

// Linkedit_data_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/linkedit_data_command
type Linkedit_data_command struct {
	Cmd      unsafe.Pointer
	Datasize unsafe.Pointer
	Cmdsize  unsafe.Pointer
}

// Linker_option_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/linker_option_command
type Linker_option_command struct {
	Cmdsize unsafe.Pointer
	Count   unsafe.Pointer
	Cmd     unsafe.Pointer
}

// Llc
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/llc
type Llc struct {
	Llc_un   unsafe.Pointer
	Llc_ssap U_int8_t
	Llc_dsap U_int8_t
}

// Load_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/load_command
type Load_command struct {
	Cmdsize unsafe.Pointer
	Cmd     unsafe.Pointer
}

// Lockd_ans
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/lockd_ans
type Lockd_ans struct {
	La_errno  unsafe.Pointer
	La_fh     U_int8_t
	La_fh_len unsafe.Pointer
	La_flags  unsafe.Pointer
	La_len    int64
	La_pid    int32
	La_start  int64
	La_xid    U_int64_t
}

// Lockd_notify
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/lockd_notify
type Lockd_notify struct {
	Ln_addr      unsafe.Pointer
	Ln_addrcount unsafe.Pointer
	Ln_flags     unsafe.Pointer
	Ln_pad       unsafe.Pointer
}

// Lockf
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/lockf
type Lockf struct {
	Lf_blkhd Locklist
	Lf_id    Caddr_t
	Lf_owner unsafe.Pointer
	Lf_type  unsafe.Pointer
}

// Locklist
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/locklist
type Locklist struct {
	Tqh_first *Lockf
	Tqh_last  *Lockf
}

// Log2phys
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/log2phys
type Log2phys struct {
	L2p_flags       uint
	L2p_contigbytes int64
	L2p_devoffset   int64
}

// Ltchars
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ltchars
type Ltchars struct {
	T_werasc unsafe.Pointer
	T_rprntc unsafe.Pointer
	T_suspc  unsafe.Pointer
	T_flushc unsafe.Pointer
	T_dsuspc unsafe.Pointer
	T_lnextc unsafe.Pointer
}

// Mach_assert_3x
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mach_assert_3x
type Mach_assert_3x struct {
	A   unsafe.Pointer
	B   unsafe.Pointer
	Hdr Mach_assert_hdr
}

// Mach_assert_default
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mach_assert_default
type Mach_assert_default struct {
	Expr unsafe.Pointer
	Hdr  Mach_assert_hdr
}

// Mach_assert_hdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mach_assert_hdr
type Mach_assert_hdr struct {
	Filename unsafe.Pointer
	Lineno   unsafe.Pointer
	Type     Mach_assert_type_t
}

// Mach_core_details
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mach_core_details
type Mach_core_details struct {
	Core_name   unsafe.Pointer
	Gzip_length unsafe.Pointer
	Gzip_offset unsafe.Pointer
}

// Mach_core_details_v2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mach_core_details_v2
type Mach_core_details_v2 struct {
	Core_name unsafe.Pointer
}

// Mach_core_fileheader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mach_core_fileheader
type Mach_core_fileheader struct {
	Signature  unsafe.Pointer
	Log_length unsafe.Pointer
	Log_offset unsafe.Pointer
	Num_files  unsafe.Pointer
	Files      Mach_core_details
}

// Mach_core_fileheader_base
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mach_core_fileheader_base
type Mach_core_fileheader_base struct {
	Signature unsafe.Pointer
	Version   unsafe.Pointer
}

// Mach_core_fileheader_v2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mach_core_fileheader_v2
type Mach_core_fileheader_v2 struct {
	Files          Mach_core_details_v2
	Flags          unsafe.Pointer
	Log_length     unsafe.Pointer
	Log_offset     unsafe.Pointer
	Num_files      unsafe.Pointer
	Pub_key_length uint16
	Pub_key_offset unsafe.Pointer
	Signature      unsafe.Pointer
	Version        unsafe.Pointer
}

// Mach_header
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mach_header
type Mach_header struct {
	Cputype    int32
	Ncmds      unsafe.Pointer
	Flags      unsafe.Pointer
	Sizeofcmds unsafe.Pointer
	Magic      unsafe.Pointer // An integer containing a value identifying this file as a 32-bit Mach-O file. Use the constant `MH_MAGIC` if the file is intended for use on a CPU with the same endianness as the computer on which the compiler is running. The constant `MH_CIGAM` can be used when the byte ordering scheme of the target machine is the reverse of the host CPU.
	Filetype   unsafe.Pointer
	Cpusubtype int32
}

// Mach_header_64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mach_header_64
type Mach_header_64 struct {
	Ncmds      unsafe.Pointer
	Sizeofcmds unsafe.Pointer
	Magic      unsafe.Pointer // An integer containing a value identifying this file as a 64-bit Mach-O file. Use the constant `MH_MAGIC_64` if the file is intended for use on a CPU with the same endianness as the computer on which the compiler is running. The constant `MH_CIGAM_64` can be used when the byte ordering scheme of the target machine is the reverse of the host CPU.
	Cpusubtype int32          // An integer specifying the exact model of the CPU. To run on all PowerPC processors supported by the macOS kernel, this should be set to `CPU_SUBTYPE_POWERPC_ALL`.
	Cputype    int32
	Reserved   unsafe.Pointer
	Filetype   unsafe.Pointer
	Flags      unsafe.Pointer // An integer containing a set of bit flags that indicate the state of certain optional features of the Mach-O file format. These are the masks you can use to manipulate this field:

}

// Macos_panic_header
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/macos_panic_header
type Macos_panic_header struct {
	Mph_padding          unsafe.Pointer
	Mph_crc              unsafe.Pointer
	Mph_data             unsafe.Pointer
	Mph_panic_flags      Mph_panic_flags_t
	Mph_magic            unsafe.Pointer
	Mph_panic_log_len    unsafe.Pointer
	Mph_panic_log_offset unsafe.Pointer
	Mph_stackshot_len    unsafe.Pointer
	Mph_other_log_offset unsafe.Pointer
	Mph_stackshot_offset unsafe.Pointer
	Mph_other_log_len    unsafe.Pointer
	Mph_roots_installed  unsafe.Pointer
}

// Mbstat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mbstat
type Mbstat struct {
	M_drops        U_int32_t
	M_clusters     U_int32_t
	M_bigclusters  U_int32_t
	M_forcedefunct U_int32_t
}

// Mbuf_stat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mbuf_stat
type Mbuf_stat struct {
	Bigmclbytes U_int32_t // Length of a big mbuf cluster.
	Mtypes      U_short   // An array of counts of each type of mbuf allocated.
	Drain       U_int32_t // Number of times protocol drain functions were called.
	Mpfail      U_int32_t // Number of times m_pullup failed.
	Msize       U_int32_t // Length of an mbuf.
	Mlen        U_int32_t // Length of data in an mbuf.
	Mhlen       U_int32_t // Length of data in an mbuf with a packet header.
	Mbufs       U_int32_t // Number of mbufs (free or otherwise).
	Clusters    U_int32_t // Number of clusters (free or otherwise).
	Mcfail      U_int32_t // Number of times m_copym failed.
	Clfree      U_int32_t // Number of free clusters.
	Mclbytes    U_int32_t // Length of an mbuf cluster.
	Bigclfree   U_int32_t // Number of unused big clusters.
	Drops       U_int32_t // Number of times allocation failed.
	Bigclusters U_int32_t // Number of big clusters.
	Wait        U_int32_t // Number of times allocation blocked.
	Minclsize   U_int32_t // Minimum length of data to allocate a cluster. Anything smaller than this should be placed in chained mbufs.

}

// Mcontext32
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mcontext32
type Mcontext32 struct {
	Es unsafe.Pointer
	Fs unsafe.Pointer
	Ss unsafe.Pointer
}

// Mcontext64_full
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mcontext64_full
type Mcontext64_full struct {
	Es unsafe.Pointer
	Fs unsafe.Pointer
	Ss unsafe.Pointer
}

// Mcontext_avx32
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mcontext_avx32
type Mcontext_avx32 struct {
	Fs unsafe.Pointer
	Ss unsafe.Pointer
	Es unsafe.Pointer
}

// Mcontext_avx512_32
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mcontext_avx512_32
type Mcontext_avx512_32 struct {
	Es unsafe.Pointer
	Ss unsafe.Pointer
	Fs unsafe.Pointer
}

// Mcontext_avx512_64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mcontext_avx512_64
type Mcontext_avx512_64 struct {
	Fs unsafe.Pointer
	Es unsafe.Pointer
	Ss unsafe.Pointer
}

// Mcontext_avx512_64_full
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mcontext_avx512_64_full
type Mcontext_avx512_64_full struct {
	Es unsafe.Pointer
	Fs unsafe.Pointer
	Ss unsafe.Pointer
}

// Mcontext_avx64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mcontext_avx64
type Mcontext_avx64 struct {
	Ss unsafe.Pointer
	Es unsafe.Pointer
	Fs unsafe.Pointer
}

// Mcontext_avx64_full
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mcontext_avx64_full
type Mcontext_avx64_full struct {
	Es unsafe.Pointer
}

// Mem_and_io_snapshot
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mem_and_io_snapshot
type Mem_and_io_snapshot struct {
	Snapshot_magic               unsafe.Pointer
	Compressor_size              unsafe.Pointer
	Speculative_pages            unsafe.Pointer
	Purgeable_pages              unsafe.Pointer
	Wired_pages                  unsafe.Pointer
	Pages_reclaimed              unsafe.Pointer
	Throttled_pages              unsafe.Pointer
	Pages_wanted                 unsafe.Pointer
	Active_pages                 unsafe.Pointer
	Decompressions               unsafe.Pointer
	Pages_wanted_reclaimed_valid uint8
	Free_pages                   unsafe.Pointer
	Compressions                 unsafe.Pointer
	Filebacked_pages             unsafe.Pointer
	Inactive_pages               unsafe.Pointer
	Busy_buffer_count            unsafe.Pointer
}

// Memory_error_notification_subsystem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/memory_error_notification_subsystem-b6h
type Memory_error_notification_subsystem struct {
	End      int32
	Maxsize  unsafe.Pointer
	Reserved Vm_address_t
	Routine  unsafe.Pointer
	Server   unsafe.Pointer
	Start    int32
}

// Micro_snapshot
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/micro_snapshot
type Micro_snapshot struct {
	Ms_flags uint8
	Ms_cpu   unsafe.Pointer
}

// Mld_hdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mld_hdr
type Mld_hdr struct {
	Mld_addr      unsafe.Pointer
	Mld_icmp6_hdr Icmp6_hdr
}

// Mmst_reg
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mmst_reg
type Mmst_reg struct {
	Mmst_reg  unsafe.Pointer
	Mmst_rsrv unsafe.Pointer
}

// Mptcp_itf_stats
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mptcp_itf_stats
type Mptcp_itf_stats struct {
	Ifindex            U_short
	Is_expensive       unsafe.Pointer
	Mpis_cell_rxbytes  unsafe.Pointer
	Mpis_cell_txbytes  unsafe.Pointer
	Mpis_rxbytes       unsafe.Pointer
	Mpis_txbytes       unsafe.Pointer
	Mpis_wifi_rxbytes  unsafe.Pointer
	Mpis_wifi_txbytes  unsafe.Pointer
	Mpis_wired_rxbytes unsafe.Pointer
	Mpis_wired_txbytes unsafe.Pointer
	Switches           uint16
}

// Msg
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/msg
type Msg struct {
	Label    unsafe.Pointer
	Msg_ts   unsafe.Pointer
	Msg_spot unsafe.Pointer
	Msg_type unsafe.Pointer
	Msg_next *Msg
}

// Msghdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/msghdr
type Msghdr struct {
	Msg_flags      unsafe.Pointer
	Msg_control    unsafe.Pointer
	Msg_controllen uint32
	Msg_iovlen     unsafe.Pointer
	Msg_name       unsafe.Pointer
	Msg_iov        *Iovec
	Msg_namelen    uint32
}

// Msginfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/msginfo-lp5
type Msginfo struct {
	Msgseg unsafe.Pointer
	Msgmni unsafe.Pointer
	Msgmnb unsafe.Pointer
	Msgmax unsafe.Pointer
	Msgtql unsafe.Pointer
	Msgssz unsafe.Pointer
}

// Msgmap
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/msgmap
type Msgmap struct {
	Next unsafe.Pointer
}

// Msqid_kernel
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/msqid_kernel
type Msqid_kernel struct {
	U     User_msqid_ds
	Label unsafe.Pointer
}

// Mwl_info_hdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mwl_info_hdr
type Mwl_info_hdr struct {
	Mwli_binds_count    unsafe.Pointer
	Mwli_chains_offset  unsafe.Pointer
	Mwli_chains_size    unsafe.Pointer
	Mwli_image_address  unsafe.Pointer
	Mwli_page_size      uint16
	Mwli_pointer_format uint16
}

// Mwl_region
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mwl_region
type Mwl_region struct {
	Mwlr_address     Mach_vm_address_t
	Mwlr_fd          unsafe.Pointer
	Mwlr_file_offset unsafe.Pointer
	Mwlr_protections Vm_prot_t
	Mwlr_size        Mach_vm_size_t
}

// Mymsg
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/mymsg
type Mymsg struct {
	Mtext unsafe.Pointer
	Mtype unsafe.Pointer
}

// Nd_neighbor_advert
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nd_neighbor_advert
type Nd_neighbor_advert struct {
	Nd_na_hdr    Icmp6_hdr
	Nd_na_target unsafe.Pointer
}

// Nd_neighbor_solicit
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nd_neighbor_solicit
type Nd_neighbor_solicit struct {
	Nd_ns_hdr    Icmp6_hdr
	Nd_ns_target unsafe.Pointer
}

// Nd_opt_dnssl
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nd_opt_dnssl
type Nd_opt_dnssl struct {
	Nd_opt_dnssl_domains  U_int8_t
	Nd_opt_dnssl_len      U_int8_t
	Nd_opt_dnssl_lifetime U_int32_t
	Nd_opt_dnssl_reserved U_int16_t
	Nd_opt_dnssl_type     U_int8_t
}

// Nd_opt_hdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nd_opt_hdr
type Nd_opt_hdr struct {
	Nd_opt_len  U_int8_t
	Nd_opt_type U_int8_t
}

// Nd_opt_mtu
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nd_opt_mtu
type Nd_opt_mtu struct {
	Nd_opt_mtu_len      U_int8_t
	Nd_opt_mtu_mtu      U_int32_t
	Nd_opt_mtu_reserved U_int16_t
	Nd_opt_mtu_type     U_int8_t
}

// Nd_opt_nonce
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nd_opt_nonce
type Nd_opt_nonce struct {
	Nd_opt_nonce      U_int8_t
	Nd_opt_nonce_len  U_int8_t
	Nd_opt_nonce_type U_int8_t
}

// Nd_opt_pref64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nd_opt_pref64
type Nd_opt_pref64 struct {
	Nd_opt_pref64_len                 U_int8_t
	Nd_opt_pref64_prefix              U_int32_t
	Nd_opt_pref64_scaled_lifetime_plc U_int16_t
	Nd_opt_pref64_type                U_int8_t
}

// Nd_opt_prefix_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nd_opt_prefix_info
type Nd_opt_prefix_info struct {
	Nd_opt_pi_flags_reserved U_int8_t
	Nd_opt_pi_len            U_int8_t
	Nd_opt_pi_preferred_time U_int32_t
	Nd_opt_pi_prefix         unsafe.Pointer
	Nd_opt_pi_prefix_len     U_int8_t
	Nd_opt_pi_reserved2      U_int32_t
	Nd_opt_pi_type           U_int8_t
	Nd_opt_pi_valid_time     U_int32_t
}

// Nd_opt_pvd
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nd_opt_pvd
type Nd_opt_pvd struct {
	Nd_opt_flags_delay U_int8_t
	Nd_opt_pvd_id      U_int8_t
	Nd_opt_pvd_len     U_int8_t
	Nd_opt_pvd_seq     U_int16_t
	Nd_opt_pvd_type    U_int8_t
}

// Nd_opt_rd_hdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nd_opt_rd_hdr
type Nd_opt_rd_hdr struct {
	Nd_opt_rh_len       U_int8_t
	Nd_opt_rh_reserved1 U_int16_t
	Nd_opt_rh_reserved2 U_int32_t
	Nd_opt_rh_type      U_int8_t
}

// Nd_opt_rdnss
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nd_opt_rdnss
type Nd_opt_rdnss struct {
	Nd_opt_rdnss_addr     unsafe.Pointer
	Nd_opt_rdnss_len      U_int8_t
	Nd_opt_rdnss_lifetime U_int32_t
	Nd_opt_rdnss_reserved U_int16_t
	Nd_opt_rdnss_type     U_int8_t
}

// Nd_opt_route_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nd_opt_route_info
type Nd_opt_route_info struct {
	Nd_opt_rti_flags     U_int8_t
	Nd_opt_rti_len       U_int8_t
	Nd_opt_rti_lifetime  U_int32_t
	Nd_opt_rti_prefixlen U_int8_t
	Nd_opt_rti_type      U_int8_t
}

// Nd_redirect
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nd_redirect
type Nd_redirect struct {
	Nd_rd_dst    unsafe.Pointer
	Nd_rd_hdr    Icmp6_hdr
	Nd_rd_target unsafe.Pointer
}

// Nd_router_advert
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nd_router_advert
type Nd_router_advert struct {
	Nd_ra_hdr        Icmp6_hdr
	Nd_ra_reachable  U_int32_t
	Nd_ra_retransmit U_int32_t
}

// Nd_router_solicit
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nd_router_solicit
type Nd_router_solicit struct {
	Nd_rs_hdr Icmp6_hdr
}

// Ndrv_demux_desc
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ndrv_demux_desc
type Ndrv_demux_desc struct {
	Data   unsafe.Pointer
	Length U_int16_t
	Type   U_int16_t
}

// Ndrv_protocol_desc
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ndrv_protocol_desc
type Ndrv_protocol_desc struct {
	Demux_count     U_int32_t
	Demux_list      *Ndrv_demux_desc
	Protocol_family U_int32_t
	Version         U_int32_t
}

// Net_event_data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/net_event_data
type Net_event_data struct {
	If_name   unsafe.Pointer
	If_unit   U_int32_t
	If_family U_int32_t
}

// Netfs_status
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/netfs_status
type Netfs_status struct {
	Ns_mountopts unsafe.Pointer
}

// Newah
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/newah
type Newah struct {
	Ah_reserve U_int16_t
	Ah_spi     U_int32_t
	Ah_seq     U_int32_t
	Ah_nxt     U_int8_t
	Ah_len     U_int8_t
}

// Newesp
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/newesp
type Newesp struct {
	Esp_spi U_int32_t
	Esp_seq U_int32_t
}

// Nextvend
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nextvend
type Nextvend struct {
	Nv_U     unsafe.Pointer
	Nv_magic U_char
}

// Nfs_args
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nfs_args
type Nfs_args struct {
	Acdirmax     unsafe.Pointer
	Acdirmin     unsafe.Pointer
	Acregmax     unsafe.Pointer
	Acregmin     unsafe.Pointer
	Addr         User32_addr_t
	Addrlen      uint8
	Auth         unsafe.Pointer
	Deadthresh   unsafe.Pointer
	Deadtimeout  unsafe.Pointer
	Fh           User32_addr_t
	Fhsize       unsafe.Pointer
	Flags        unsafe.Pointer
	Hostname     User32_addr_t
	Leaseterm    unsafe.Pointer
	Maxgrouplist unsafe.Pointer
	Proto        unsafe.Pointer
	Readahead    unsafe.Pointer
	Readdirsize  unsafe.Pointer
	Retrans      unsafe.Pointer
	Rsize        unsafe.Pointer
	Sotype       unsafe.Pointer
	Timeo        unsafe.Pointer
	Version      unsafe.Pointer
	Wsize        unsafe.Pointer
}

// Nfs_etype
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nfs_etype
type Nfs_etype struct {
	Count    unsafe.Pointer
	Etypes   Nfs_supported_kerberos_etypes
	Selected unsafe.Pointer
}

// Nfs_exphandle
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nfs_exphandle
type Nfs_exphandle struct {
	Nxh_expid    unsafe.Pointer
	Nxh_fidlen   unsafe.Pointer
	Nxh_flags    uint16
	Nxh_fsid     unsafe.Pointer
	Nxh_reserved uint8
}

// Nfs_export_args
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nfs_export_args
type Nfs_export_args struct {
	Nxa_expid    unsafe.Pointer
	Nxa_exppath  User32_addr_t
	Nxa_flags    unsafe.Pointer
	Nxa_fsid     unsafe.Pointer
	Nxa_fspath   User32_addr_t
	Nxa_netcount unsafe.Pointer
	Nxa_nets     User32_addr_t
}

// Nfs_export_net_args
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nfs_export_net_args
type Nfs_export_net_args struct {
	Nxna_addr  unsafe.Pointer
	Nxna_cred  Xucred
	Nxna_flags unsafe.Pointer
	Nxna_mask  unsafe.Pointer
}

// Nfs_export_stat_desc
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nfs_export_stat_desc
type Nfs_export_stat_desc struct {
	Rec_count unsafe.Pointer
	Rec_vers  unsafe.Pointer
}

// Nfs_export_stat_rec
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nfs_export_stat_rec
type Nfs_export_stat_rec struct {
	Bytes_read    unsafe.Pointer
	Bytes_written unsafe.Pointer
	Ops           unsafe.Pointer
	Path          unsafe.Pointer
}

// Nfs_filehandle
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nfs_filehandle
type Nfs_filehandle struct {
	Nfh_fhp unsafe.Pointer
	Nfh_fid unsafe.Pointer
	Nfh_len unsafe.Pointer
	Nfh_xh  Nfs_exphandle
}

// Nfs_sec
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nfs_sec
type Nfs_sec struct {
	Count   unsafe.Pointer
	Flavors unsafe.Pointer
}

// Nfs_testmapid
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nfs_testmapid
type Nfs_testmapid struct {
	Ntm_grpflag unsafe.Pointer
	Ntm_guid    unsafe.Pointer
	Ntm_id      unsafe.Pointer
	Ntm_lookup  unsafe.Pointer
	Ntm_name    unsafe.Pointer
	Pad         unsafe.Pointer
}

// Nfs_user_stat_desc
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nfs_user_stat_desc
type Nfs_user_stat_desc struct {
	Rec_count unsafe.Pointer
	Rec_vers  unsafe.Pointer
}

// Nfs_user_stat_path_rec
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nfs_user_stat_path_rec
type Nfs_user_stat_path_rec struct {
	Path     unsafe.Pointer
	Rec_type U_char
}

// Nfs_user_stat_user_rec
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nfs_user_stat_user_rec
type Nfs_user_stat_user_rec struct {
	Bytes_read    unsafe.Pointer
	Bytes_written unsafe.Pointer
	Ops           unsafe.Pointer
	Rec_type      U_char
	Sock          unsafe.Pointer
	Tm_last       int64
	Tm_start      int64
	Uid           uint32
}

// Nfsclntstats
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nfsclntstats
type Nfsclntstats struct {
	Attrcache_hits     unsafe.Pointer
	Attrcache_misses   unsafe.Pointer
	Biocache_readdirs  unsafe.Pointer
	Biocache_readlinks unsafe.Pointer
	Biocache_reads     unsafe.Pointer
	Biocache_writes    unsafe.Pointer
	Direofcache_hits   unsafe.Pointer
	Direofcache_misses unsafe.Pointer
	Lookupcache_hits   unsafe.Pointer
	Lookupcache_misses unsafe.Pointer
	Opcntv4            unsafe.Pointer
	Pageins            unsafe.Pointer
	Pageouts           unsafe.Pointer
	Read_bios          unsafe.Pointer
	Read_physios       unsafe.Pointer
	Readdir_bios       unsafe.Pointer
	Readlink_bios      unsafe.Pointer
	Rpccntv3           unsafe.Pointer
	Rpcinvalid         unsafe.Pointer
	Rpcrequests        unsafe.Pointer
	Rpcretries         unsafe.Pointer
	Rpctimeouts        unsafe.Pointer
	Rpcunexpected      unsafe.Pointer
	Write_bios         unsafe.Pointer
	Write_physios      unsafe.Pointer
	Accesscache_hits   unsafe.Pointer
	Accesscache_misses unsafe.Pointer
	Nfs_errs           unsafe.Pointer
	Nlmcnt             unsafe.Pointer
	Cbopcntv4          unsafe.Pointer
}

// Nfsd_args
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nfsd_args
type Nfsd_args struct {
	Name    User32_addr_t
	Namelen unsafe.Pointer
	Sock    unsafe.Pointer
}

// Nfsrvstats
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nfsrvstats
type Nfsrvstats struct {
	Srv_errs                 unsafe.Pointer
	Srvcache_idemdonehits    unsafe.Pointer
	Srvcache_inproghits      unsafe.Pointer
	Srvcache_misses          unsafe.Pointer
	Srvcache_nonidemdonehits unsafe.Pointer
	Srvrpc_errs              unsafe.Pointer
	Srvrpccntv3              unsafe.Pointer
	Srvvop_writes            unsafe.Pointer
	Nfs_errs                 unsafe.Pointer
}

// Ni_reply_fqdn
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ni_reply_fqdn
type Ni_reply_fqdn struct {
	Ni_fqdn_name    U_int8_t
	Ni_fqdn_namelen U_int8_t
	Ni_fqdn_ttl     U_int32_t
}

// Nlist - Describes an entry in the symbol table for 32-bit architectures. Declared in `/usr/include/mach-o/nlist.H()`. See also [nlist_64](<doc://com.apple.documentation/documentation/kernel/nlist_64>).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nlist
type Nlist struct {
	N_un    unsafe.Pointer // A union that holds an index into the string table, `n_strx`. To specify an empty string (`""`), set this value to 0. The `n_name` field is not used in Mach-O files.
	N_sect  uint8          // An integer specifying the number of the section that this symbol can be found in, or `NO_SECT` if the symbol is not to be found in any section of this image. The sections are contiguously numbered across segments, starting from 1, according to the order they appear in the `LC_SEGMENT` load commands.
	N_value unsafe.Pointer // An integer that contains the value of the symbol. The format of this value is different for each type of symbol table entry (as specified by the `n_type` field). For the `N_SECT` symbol type, `n_value` is the address of the symbol. See the description of the `n_type` field for information on other possible values.
	N_desc  int16          // A 16-bit value providing additional information about the nature of this symbol for non-stab symbols. The reference flags can be accessed using the `REFERENCE_TYPE` mask (0xF) and are defined as follows:
	N_type  uint8          // A byte value consisting of data accessed using four bit masks:

}

// Nlist_64 - Describes an entry in the symbol table for 64-bit architectures. Declared in `/usr/include/mach-o/nlist.H()`.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/nlist_64
type Nlist_64 struct {
	N_un    unsafe.Pointer // A union that holds an index into the string table, `n_strx`. To specify an empty string (`""`), set this value to 0.
	N_value unsafe.Pointer
	N_type  uint8  // A byte value consisting of data accessed using four bit masks:
	N_desc  uint16 // A 16-bit value providing additional information about the nature of this symbol. The reference flags can be accessed using the `REFERENCE_TYPE` mask (0xF) and are defined as follows:
	N_sect  uint8  // An integer specifying the number of the section that this symbol can be found in, or `NO_SECT` if the symbol is not to be found in any section of this image. The sections are contiguously numbered across segments, starting from 1, according to the order they appear in the `LC_SEGMENT` load commands.

}

// Note_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/note_command
type Note_command struct {
	Data_owner unsafe.Pointer
	Cmdsize    unsafe.Pointer
	Offset     unsafe.Pointer
	Size       unsafe.Pointer
	Cmd        unsafe.Pointer
}

// Ntptimeval
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ntptimeval
type Ntptimeval struct {
	Time_state unsafe.Pointer
	Maxerror   unsafe.Pointer
	Esterror   unsafe.Pointer
}

// Ombstat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ombstat
type Ombstat struct {
	M_mhlen     U_int32_t
	M_wait      U_int32_t
	M_mlen      U_int32_t
	M_minclsize U_int32_t
	M_spare     U_int32_t
	M_clfree    U_int32_t
	M_msize     U_int32_t
	M_clusters  U_int32_t
	M_drops     U_int32_t
	M_mtypes    U_short
	M_mclbytes  U_int32_t
	M_mpfail    U_int32_t
	M_mcfail    U_int32_t
	M_drain     U_int32_t
	M_mbufs     U_int32_t
}

// Opmask_reg
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/opmask_reg
type Opmask_reg struct {
	Opmask_reg unsafe.Pointer
}

// Ostat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ostat
type Ostat struct {
	St_size      int32
	St_ctimespec Timespec
	St_dev       uint16
	St_mode      uint16
	St_blocks    int32
	St_flags     uint32
	St_rdev      uint16
	St_gen       uint32
	St_mtimespec Timespec
	St_ino       uint64
	St_uid       uint16
	St_blksize   int32
	St_nlink     uint16
	St_atimespec Timespec
	St_gid       uint16
}

// Persona_modify_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/persona_modify_info
type Persona_modify_info struct {
	Persona_id unsafe.Pointer
	Unique_pid unsafe.Pointer
}

// Persona_token
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/persona_token
type Persona_token struct {
	Proximate  Proc_persona_info
	Originator Proc_persona_info
}

// Portlabel_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/portlabel_info
type Portlabel_info struct {
	Portlabel_domain uint8
	Portlabel_flags  uint16
	Portlabel_id     int16
}

// Prebind_cksum_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/prebind_cksum_command
type Prebind_cksum_command struct {
	Cksum   unsafe.Pointer
	Cmd     unsafe.Pointer
	Cmdsize unsafe.Pointer
}

// Prebound_dylib_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/prebound_dylib_command
type Prebound_dylib_command struct {
	Cmd unsafe.Pointer // Common to all load command structures. For this structure, set to `LC_PREBOUND_DYLIB`.

}

// Priority_queue_deadline_max
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/priority_queue_deadline_max
type Priority_queue_deadline_max struct {
	Pq_root unsafe.Pointer
}

// Priority_queue_deadline_min
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/priority_queue_deadline_min
type Priority_queue_deadline_min struct {
	Pq_root unsafe.Pointer
}

// Priority_queue_max
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/priority_queue_max
type Priority_queue_max struct {
	Pq_cmp_fn Priority_queue_compare_fn_t
	Pq_root   unsafe.Pointer
}

// Priority_queue_min
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/priority_queue_min
type Priority_queue_min struct {
	Pq_cmp_fn Priority_queue_compare_fn_t
	Pq_root   unsafe.Pointer
}

// Priority_queue_sched_max
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/priority_queue_sched_max
type Priority_queue_sched_max struct {
	Pq_root unsafe.Pointer
}

// Priority_queue_sched_min
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/priority_queue_sched_min
type Priority_queue_sched_min struct {
	Pq_root unsafe.Pointer
}

// Priority_queue_sched_stable_max
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/priority_queue_sched_stable_max
type Priority_queue_sched_stable_max struct {
	Pq_root unsafe.Pointer
}

// Priority_queue_sched_stable_min
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/priority_queue_sched_stable_min
type Priority_queue_sched_stable_min struct {
	Pq_root unsafe.Pointer
}

// Proc_persona_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/proc_persona_info
type Proc_persona_info struct {
	Macho_uuid uint8
	Unique_pid unsafe.Pointer
	Flags      unsafe.Pointer
	Uid        unsafe.Pointer
	Persona_id unsafe.Pointer
	Pidversion unsafe.Pointer
	Gid        unsafe.Pointer
	Pid        unsafe.Pointer
}

// Proc_rlimit_control_wakeupmon
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/proc_rlimit_control_wakeupmon
type Proc_rlimit_control_wakeupmon struct {
	Wm_flags uint32
	Wm_rate  int32
}

// Pseminfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/pseminfo
type Pseminfo struct {
	Psem_name             unsafe.Pointer
	Psem_gid              uint32
	Psem_creator_uniqueid unsafe.Pointer
	Psem_mode             uint16
	Psem_uid              uint32
	Psem_creator_pid      int32
	Psem_label            unsafe.Pointer
	Psem_flags            unsafe.Pointer
	Psem_semobject        unsafe.Pointer
	Psem_usecount         unsafe.Pointer
}

// Pshminfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/pshminfo
type Pshminfo struct {
	Pshm_gid       uint32
	Pshm_length    int64
	Pshm_usecount  unsafe.Pointer
	Pshm_label     unsafe.Pointer
	Pshm_memobject unsafe.Pointer
	Pshm_mode      uint16
	Pshm_uid       uint32
	Pshm_name      unsafe.Pointer
	Pshm_flags     unsafe.Pointer
}

// Radvisory
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/radvisory
type Radvisory struct {
	Ra_offset int64
	Ra_count  int
}

// Receive_sysdiagnose_notification_subsystem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/receive_sysdiagnose_notification_subsystem
type Receive_sysdiagnose_notification_subsystem struct {
	End      int32
	Server   unsafe.Pointer
	Maxsize  unsafe.Pointer
	Reserved Vm_address_t
	Start    int32
	Routine  unsafe.Pointer
}

// Receive_vfs_nspace_subsystem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/receive_vfs_nspace_subsystem-rhm
type Receive_vfs_nspace_subsystem struct {
	End      int32
	Maxsize  unsafe.Pointer
	Reserved Vm_address_t
	Routine  unsafe.Pointer
	Server   unsafe.Pointer
	Start    int32
}

// Reg_desc
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/reg_desc
type Reg_desc struct {
	Rd_mask   unsafe.Pointer
	Rd_values *Reg_values
	Rd_format unsafe.Pointer
	Rd_shift  unsafe.Pointer
	Rd_name   unsafe.Pointer
}

// Reg_values
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/reg_values
type Reg_values struct {
	Rv_name  unsafe.Pointer
	Rv_value unsafe.Pointer
}

// Relocation_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/relocation_info
type Relocation_info struct {
	R_pcrel     unsafe.Pointer // Indicates whether the item containing the address to be relocated is part of a CPU instruction that uses PC-relative addressing.
	R_type      unsafe.Pointer
	R_length    unsafe.Pointer // Indicates the length of the item containing the address to be relocated.
	R_address   unsafe.Pointer // In `MH_OBJECT` files, an offset from the start of the section to the item containing the address requiring relocation.
	R_symbolnum unsafe.Pointer // Indicates either an index into the symbol table (when the `r_extern` field is set to 1) or a section number (when the `r_extern` field is set to 0). As previously mentioned, sections are ordered from 1 to 255 in the order in which they appear in the `LC_SEGMENT` load commands. This field is set to `R_ABS` for relocation entries for absolute symbols, which need no relocation.

}

// Rip6stat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rip6stat
type Rip6stat struct {
	Rip6s_fullsock U_quad_t
	Rip6s_ipackets U_quad_t
	Rip6s_isum     U_quad_t
	Rip6s_nosock   U_quad_t
}

// Rlimit
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rlimit
type Rlimit struct {
	Rlim_cur Rlim_t
	Rlim_max Rlim_t
}

// Route_in6_old
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/route_in6_old
type Route_in6_old struct {
	Ro_dst   Sockaddr_in6
	Ro_flags unsafe.Pointer
	Ro_rt    unsafe.Pointer
}

// Route_old
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/route_old
type Route_old struct {
	Ro_dst   unsafe.Pointer
	Ro_flags unsafe.Pointer
	Ro_rt    unsafe.Pointer
}

// Routines_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/routines_command
type Routines_command struct {
	Cmdsize      unsafe.Pointer
	Reserved2    unsafe.Pointer // Reserved for future use. Set this field to `0`.
	Reserved6    unsafe.Pointer
	Reserved4    unsafe.Pointer
	Reserved5    unsafe.Pointer
	Reserved3    unsafe.Pointer // Reserved for future use. Set this field to `0`.
	Reserved1    unsafe.Pointer
	Init_address unsafe.Pointer // An integer specifying the virtual memory address of the initialization function.
	Init_module  unsafe.Pointer
	Cmd          unsafe.Pointer
}

// Routines_command_64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/routines_command_64
type Routines_command_64 struct {
	Init_module  unsafe.Pointer
	Reserved4    unsafe.Pointer
	Init_address unsafe.Pointer
	Reserved3    unsafe.Pointer
	Reserved1    unsafe.Pointer
	Reserved5    unsafe.Pointer // Reserved for future use. Set this field to `0`.
	Cmdsize      unsafe.Pointer
	Cmd          unsafe.Pointer // Common to all load command structures. For this structure, set to `LC_ROUTINES_64`.
	Reserved6    unsafe.Pointer
	Reserved2    unsafe.Pointer // Reserved for future use. Set this field to `0`.

}

// Rpath_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rpath_command
type Rpath_command struct {
	Cmdsize unsafe.Pointer
	Path    unsafe.Pointer
	Cmd     unsafe.Pointer
}

// Rpc_signature
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rpc_signature
type Rpc_signature struct {
	Rd  unsafe.Pointer
	Rad unsafe.Pointer
}

// Rr_pco_match
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rr_pco_match
type Rr_pco_match struct {
	Rpm_code     U_int8_t
	Rpm_len      U_int8_t
	Rpm_matchlen U_int8_t
	Rpm_maxlen   U_int8_t
	Rpm_minlen   U_int8_t
	Rpm_ordinal  U_int8_t
	Rpm_prefix   unsafe.Pointer
	Rpm_reserved U_int16_t
}

// Rr_pco_use
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rr_pco_use
type Rr_pco_use struct {
	Rpu_flags   U_int32_t
	Rpu_keeplen U_int8_t
	Rpu_pltime  U_int32_t
}

// Rr_result
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rr_result
type Rr_result struct {
	Rrr_flags      U_int16_t
	Rrr_ifid       U_int32_t
	Rrr_matchedlen U_int8_t
	Rrr_ordinal    U_int8_t
	Rrr_prefix     unsafe.Pointer
}

// Rslvmulti_req
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rslvmulti_req
type Rslvmulti_req struct {
	Sa   unsafe.Pointer
	Llsa unsafe.Pointer
}

// Rt_addrinfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rt_addrinfo
type Rt_addrinfo struct {
	Rti_info  unsafe.Pointer
	Rti_addrs unsafe.Pointer
}

// Rt_addrinfo_ext
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rt_addrinfo_ext
type Rt_addrinfo_ext struct {
	Rtix_info      Rt_addrinfo
	Rtix_next_tiny uint8
	Rtix_tiny_addr unsafe.Pointer
}

// Rt_metrics
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rt_metrics
type Rt_metrics struct {
	Rmx_rttvar U_int32_t
}

// Rt_msghdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rt_msghdr
type Rt_msghdr struct {
	Rtm_errno unsafe.Pointer
	Rtm_addrs unsafe.Pointer
}

// Rt_msghdr2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rt_msghdr2
type Rt_msghdr2 struct {
	Rtm_index       U_short
	Rtm_type        U_char
	Rtm_msglen      U_short
	Rtm_flags       unsafe.Pointer
	Rtm_rmx         Rt_metrics
	Rtm_inits       U_int32_t
	Rtm_parentflags unsafe.Pointer
	Rtm_refcnt      unsafe.Pointer
	Rtm_use         unsafe.Pointer
	Rtm_addrs       unsafe.Pointer
	Rtm_reserved    unsafe.Pointer
}

// Rt_msghdr_common
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rt_msghdr_common
type Rt_msghdr_common struct {
	Rtm_errno  unsafe.Pointer
	Rtm_flags  unsafe.Pointer
	Rtm_index  U_short
	Rtm_msglen U_short
	Rtm_type   U_char
}

// Rt_msghdr_ext
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rt_msghdr_ext
type Rt_msghdr_ext struct {
	Rtm_addrs    U_int32_t
	Rtm_errno    unsafe.Pointer
	Rtm_flags    U_int32_t
	Rtm_index    U_int32_t
	Rtm_inits    U_int32_t
	Rtm_msglen   U_short
	Rtm_pid      int32
	Rtm_reserved U_int32_t
	Rtm_ri       Rt_reach_info
	Rtm_rmx      Rt_metrics
	Rtm_seq      unsafe.Pointer
	Rtm_type     U_char
	Rtm_use      U_int32_t
}

// Rt_msghdr_prelude
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rt_msghdr_prelude
type Rt_msghdr_prelude struct {
	Rtm_msglen U_short
}

// Rt_reach_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rt_reach_info
type Rt_reach_info struct {
	Ri_lqm        unsafe.Pointer
	Ri_npm        unsafe.Pointer
	Ri_probes     U_int32_t
	Ri_rcv_expire U_int64_t
	Ri_refcnt     U_int32_t
	Ri_rssi       unsafe.Pointer
	Ri_snd_expire U_int64_t
}

// Rtstat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rtstat
type Rtstat struct {
	Rts_dynamic unsafe.Pointer
}

// Rtstat_64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rtstat_64
type Rtstat_64 struct {
	Rts_badredirect  unsafe.Pointer
	Rts_badrtgwroute unsafe.Pointer
	Rts_dynamic      unsafe.Pointer
	Rts_newgateway   unsafe.Pointer
	Rts_unreach      unsafe.Pointer
	Rts_wildcard     unsafe.Pointer
}

// Rusage
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rusage
type Rusage struct {
	Ru_utime    Timeval
	Ru_stime    Timeval
	Ru_maxrss   int
	Ru_ixrss    int
	Ru_idrss    int
	Ru_isrss    int
	Ru_minflt   int
	Ru_majflt   int
	Ru_nswap    int
	Ru_inblock  int
	Ru_oublock  int
	Ru_msgsnd   int
	Ru_msgrcv   int
	Ru_nsignals int
	Ru_nvcsw    int
	Ru_nivcsw   int
}

// Rusage_info_child
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rusage_info_child
type Rusage_info_child struct {
	Ri_child_interrupt_wkups unsafe.Pointer
	Ri_child_user_time       unsafe.Pointer
	Ri_child_system_time     unsafe.Pointer
	Ri_child_elapsed_abstime unsafe.Pointer
	Ri_child_pageins         unsafe.Pointer
	Ri_child_pkg_idle_wkups  unsafe.Pointer
}

// Rusage_info_v0
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rusage_info_v0
type Rusage_info_v0 struct {
	Ri_user_time          uint64
	Ri_system_time        uint64
	Ri_pkg_idle_wkups     uint64
	Ri_interrupt_wkups    uint64
	Ri_pageins            uint64
	Ri_wired_size         uint64
	Ri_resident_size      uint64
	Ri_phys_footprint     uint64
	Ri_proc_start_abstime uint64
	Ri_proc_exit_abstime  uint64
	Ri_uuid               uint8
}

// Rusage_info_v1
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rusage_info_v1
type Rusage_info_v1 struct {
	Ri_user_time             uint64
	Ri_system_time           uint64
	Ri_pkg_idle_wkups        uint64
	Ri_interrupt_wkups       uint64
	Ri_pageins               uint64
	Ri_wired_size            uint64
	Ri_resident_size         uint64
	Ri_phys_footprint        uint64
	Ri_proc_start_abstime    uint64
	Ri_proc_exit_abstime     uint64
	Ri_child_user_time       uint64
	Ri_child_system_time     uint64
	Ri_child_pkg_idle_wkups  uint64
	Ri_child_interrupt_wkups uint64
	Ri_child_pageins         uint64
	Ri_child_elapsed_abstime uint64
	Ri_uuid                  unsafe.Pointer
}

// Rusage_info_v2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rusage_info_v2
type Rusage_info_v2 struct {
	Ri_user_time             uint64
	Ri_system_time           uint64
	Ri_pkg_idle_wkups        uint64
	Ri_interrupt_wkups       uint64
	Ri_pageins               uint64
	Ri_wired_size            uint64
	Ri_resident_size         uint64
	Ri_phys_footprint        uint64
	Ri_proc_start_abstime    uint64
	Ri_proc_exit_abstime     uint64
	Ri_child_user_time       uint64
	Ri_child_system_time     uint64
	Ri_child_pkg_idle_wkups  uint64
	Ri_child_interrupt_wkups uint64
	Ri_child_pageins         uint64
	Ri_child_elapsed_abstime uint64
	Ri_diskio_bytesread      uint64
	Ri_diskio_byteswritten   uint64
	Ri_uuid                  uint8
}

// Rusage_info_v3
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rusage_info_v3
type Rusage_info_v3 struct {
	Ri_user_time                     uint64
	Ri_system_time                   uint64
	Ri_pkg_idle_wkups                uint64
	Ri_interrupt_wkups               uint64
	Ri_pageins                       uint64
	Ri_wired_size                    uint64
	Ri_resident_size                 uint64
	Ri_phys_footprint                uint64
	Ri_proc_start_abstime            uint64
	Ri_proc_exit_abstime             uint64
	Ri_child_user_time               uint64
	Ri_child_system_time             uint64
	Ri_child_pkg_idle_wkups          uint64
	Ri_child_interrupt_wkups         uint64
	Ri_child_pageins                 uint64
	Ri_child_elapsed_abstime         uint64
	Ri_diskio_bytesread              uint64
	Ri_diskio_byteswritten           uint64
	Ri_cpu_time_qos_default          uint64
	Ri_cpu_time_qos_maintenance      uint64
	Ri_cpu_time_qos_background       uint64
	Ri_cpu_time_qos_utility          uint64
	Ri_cpu_time_qos_legacy           uint64
	Ri_cpu_time_qos_user_initiated   uint64
	Ri_cpu_time_qos_user_interactive uint64
	Ri_billed_system_time            uint64
	Ri_serviced_system_time          uint64
	Ri_uuid                          unsafe.Pointer
}

// Rusage_info_v4
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rusage_info_v4
type Rusage_info_v4 struct {
	Ri_user_time                     uint64
	Ri_system_time                   uint64
	Ri_pkg_idle_wkups                uint64
	Ri_interrupt_wkups               uint64
	Ri_pageins                       uint64
	Ri_wired_size                    uint64
	Ri_resident_size                 uint64
	Ri_phys_footprint                uint64
	Ri_proc_start_abstime            uint64
	Ri_proc_exit_abstime             uint64
	Ri_child_user_time               uint64
	Ri_child_system_time             uint64
	Ri_child_pkg_idle_wkups          uint64
	Ri_child_interrupt_wkups         uint64
	Ri_child_pageins                 uint64
	Ri_child_elapsed_abstime         uint64
	Ri_diskio_bytesread              uint64
	Ri_diskio_byteswritten           uint64
	Ri_cpu_time_qos_default          uint64
	Ri_cpu_time_qos_maintenance      uint64
	Ri_cpu_time_qos_background       uint64
	Ri_cpu_time_qos_utility          uint64
	Ri_cpu_time_qos_legacy           uint64
	Ri_cpu_time_qos_user_initiated   uint64
	Ri_cpu_time_qos_user_interactive uint64
	Ri_billed_system_time            uint64
	Ri_serviced_system_time          uint64
	Ri_logical_writes                uint64
	Ri_lifetime_max_phys_footprint   uint64
	Ri_instructions                  uint64
	Ri_cycles                        uint64
	Ri_billed_energy                 uint64
	Ri_serviced_energy               uint64
	Ri_interval_max_phys_footprint   uint64
	Ri_runnable_time                 uint64
	Ri_uuid                          unsafe.Pointer
}

// Rusage_info_v5
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rusage_info_v5
type Rusage_info_v5 struct {
	Ri_user_time                     uint64
	Ri_system_time                   uint64
	Ri_pkg_idle_wkups                uint64
	Ri_interrupt_wkups               uint64
	Ri_pageins                       uint64
	Ri_wired_size                    uint64
	Ri_resident_size                 uint64
	Ri_phys_footprint                uint64
	Ri_proc_start_abstime            uint64
	Ri_proc_exit_abstime             uint64
	Ri_child_user_time               uint64
	Ri_child_system_time             uint64
	Ri_child_pkg_idle_wkups          uint64
	Ri_child_interrupt_wkups         uint64
	Ri_child_pageins                 uint64
	Ri_child_elapsed_abstime         uint64
	Ri_diskio_bytesread              uint64
	Ri_diskio_byteswritten           uint64
	Ri_cpu_time_qos_default          uint64
	Ri_cpu_time_qos_maintenance      uint64
	Ri_cpu_time_qos_background       uint64
	Ri_cpu_time_qos_utility          uint64
	Ri_cpu_time_qos_legacy           uint64
	Ri_cpu_time_qos_user_initiated   uint64
	Ri_cpu_time_qos_user_interactive uint64
	Ri_billed_system_time            uint64
	Ri_serviced_system_time          uint64
	Ri_logical_writes                uint64
	Ri_lifetime_max_phys_footprint   uint64
	Ri_instructions                  uint64
	Ri_cycles                        uint64
	Ri_billed_energy                 uint64
	Ri_serviced_energy               uint64
	Ri_interval_max_phys_footprint   uint64
	Ri_runnable_time                 uint64
	Ri_flags                         uint64
	Ri_uuid                          unsafe.Pointer
}

// Rusage_superset
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/rusage_superset
type Rusage_superset struct {
	Ri Rusage_info_current
	Ru Rusage
}

// Sadb_address
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sadb_address
type Sadb_address struct {
	Sadb_address_prefixlen U_int8_t
	Sadb_address_reserved  U_int16_t
	Sadb_address_exttype   U_int16_t
	Sadb_address_proto     U_int8_t
	Sadb_address_len       U_int16_t
}

// Sadb_alg
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sadb_alg
type Sadb_alg struct {
	Sadb_alg_ivlen    U_int8_t
	Sadb_alg_id       U_int8_t
	Sadb_alg_reserved U_int16_t
	Sadb_alg_minbits  U_int16_t
	Sadb_alg_maxbits  U_int16_t
}

// Sadb_comb
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sadb_comb
type Sadb_comb struct {
	Sadb_comb_encrypt_minbits  U_int16_t
	Sadb_comb_soft_addtime     U_int64_t
	Sadb_comb_reserved         U_int32_t
	Sadb_comb_auth             U_int8_t
	Sadb_comb_auth_maxbits     U_int16_t
	Sadb_comb_encrypt          U_int8_t
	Sadb_comb_hard_allocations U_int32_t
	Sadb_comb_hard_usetime     U_int64_t
	Sadb_comb_soft_bytes       U_int64_t
	Sadb_comb_soft_allocations U_int32_t
	Sadb_comb_hard_bytes       U_int64_t
	Sadb_comb_auth_minbits     U_int16_t
	Sadb_comb_flags            U_int16_t
	Sadb_comb_hard_addtime     U_int64_t
	Sadb_comb_encrypt_maxbits  U_int16_t
	Sadb_comb_soft_usetime     U_int64_t
}

// Sadb_ext
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sadb_ext
type Sadb_ext struct {
	Sadb_ext_type U_int16_t
	Sadb_ext_len  U_int16_t
}

// Sadb_ident
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sadb_ident
type Sadb_ident struct {
	Sadb_ident_len      U_int16_t
	Sadb_ident_type     U_int16_t
	Sadb_ident_id       U_int64_t
	Sadb_ident_exttype  U_int16_t
	Sadb_ident_reserved U_int16_t
}

// Sadb_key
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sadb_key
type Sadb_key struct {
	Sadb_key_bits     U_int16_t
	Sadb_key_exttype  U_int16_t
	Sadb_key_len      U_int16_t
	Sadb_key_reserved U_int16_t
}

// Sadb_lifetime
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sadb_lifetime
type Sadb_lifetime struct {
	Sadb_lifetime_addtime     U_int64_t
	Sadb_lifetime_exttype     U_int16_t
	Sadb_lifetime_allocations U_int32_t
	Sadb_lifetime_bytes       U_int64_t
}

// Sadb_msg
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sadb_msg
type Sadb_msg struct {
	Sadb_msg_satype   U_int8_t
	Sadb_msg_reserved U_int16_t
}

// Sadb_prop
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sadb_prop
type Sadb_prop struct {
	Sadb_prop_replay   U_int8_t
	Sadb_prop_exttype  U_int16_t
	Sadb_prop_len      U_int16_t
	Sadb_prop_reserved U_int8_t
}

// Sadb_sa
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sadb_sa
type Sadb_sa struct {
	Sadb_sa_auth    U_int8_t
	Sadb_sa_exttype U_int16_t
	Sadb_sa_flags   U_int32_t
	Sadb_sa_replay  U_int8_t
	Sadb_sa_encrypt U_int8_t
	Sadb_sa_state   U_int8_t
	Sadb_sa_spi     U_int32_t
	Sadb_sa_len     U_int16_t
}

// Sadb_sastat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sadb_sastat
type Sadb_sastat struct {
	Sadb_sastat_exttype  U_int16_t
	Sadb_sastat_dir      U_int32_t
	Sadb_sastat_list_len U_int32_t
	Sadb_sastat_reserved U_int32_t
	Sadb_sastat_len      U_int16_t
}

// Sadb_sens
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sadb_sens
type Sadb_sens struct {
	Sadb_sens_sens_len    U_int8_t
	Sadb_sens_sens_level  U_int8_t
	Sadb_sens_exttype     U_int16_t
	Sadb_sens_reserved    U_int32_t
	Sadb_sens_integ_len   U_int8_t
	Sadb_sens_integ_level U_int8_t
	Sadb_sens_len         U_int16_t
	Sadb_sens_dpd         U_int32_t
}

// Sadb_session_id
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sadb_session_id
type Sadb_session_id struct {
	Sadb_session_id_v       U_int64_t
	Sadb_session_id_len     U_int16_t
	Sadb_session_id_exttype U_int16_t
}

// Sadb_spirange
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sadb_spirange
type Sadb_spirange struct {
	Sadb_spirange_min      U_int32_t
	Sadb_spirange_reserved U_int32_t
	Sadb_spirange_len      U_int16_t
	Sadb_spirange_max      U_int32_t
	Sadb_spirange_exttype  U_int16_t
}

// Sadb_supported
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sadb_supported
type Sadb_supported struct {
	Sadb_supported_exttype  U_int16_t
	Sadb_supported_len      U_int16_t
	Sadb_supported_reserved U_int32_t
}

// Sadb_x_ipsecrequest
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sadb_x_ipsecrequest
type Sadb_x_ipsecrequest struct {
	Sadb_x_ipsecrequest_level U_int8_t
	Sadb_x_ipsecrequest_mode  U_int8_t
	Sadb_x_ipsecrequest_len   U_int16_t
	Sadb_x_ipsecrequest_proto U_int16_t
}

// Sadb_x_kmprivate
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sadb_x_kmprivate
type Sadb_x_kmprivate struct {
	Sadb_x_kmprivate_len      U_int16_t
	Sadb_x_kmprivate_exttype  U_int16_t
	Sadb_x_kmprivate_reserved U_int32_t
}

// Sadb_x_policy
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sadb_x_policy
type Sadb_x_policy struct {
	Sadb_x_policy_dir       U_int8_t
	Sadb_x_policy_reserved  U_int8_t
	Sadb_x_policy_len       U_int16_t
	Sadb_x_policy_id        U_int32_t
	Sadb_x_policy_exttype   U_int16_t
	Sadb_x_policy_reserved2 U_int32_t
	Sadb_x_policy_type      U_int16_t
}

// Sadb_x_sa2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sadb_x_sa2
type Sadb_x_sa2 struct {
	Sadb_x_sa2_len     U_int16_t
	Sadb_x_sa2_mode    U_int8_t
	Sadb_x_sa2_exttype U_int16_t
	Sadb_x_sa2_reqid   U_int32_t
}

// Sastat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sastat
type Sastat struct {
	Spi     U_int32_t
	Lft_c   Sadb_lifetime
	Created U_int32_t
}

// Sbuf
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sbuf
type Sbuf struct {
	S_len    unsafe.Pointer
	S_unused unsafe.Pointer
	S_buf    unsafe.Pointer
	S_flags  unsafe.Pointer
	S_size   unsafe.Pointer
}

// Scattered_relocation_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/scattered_relocation_info
type Scattered_relocation_info struct {
	R_pcrel     unsafe.Pointer // Indicates whether the item containing the address to be relocated is part of a CPU instruction that uses PC-relative addressing.
	R_address   unsafe.Pointer
	R_length    unsafe.Pointer
	R_scattered unsafe.Pointer // If this bit is 0, this structure is actually a [relocation_info](<doc://com.apple.documentation/documentation/kernel/relocation_info>) structure.
	R_type      unsafe.Pointer // Indicates the type of relocation to be performed. Possible values for this field are shared between this structure and the `relocation_info` data structure; see the description of the `r_type` field in the [relocation_info](<doc://com.apple.documentation/documentation/kernel/relocation_info>) data structure for more details.
	R_value     unsafe.Pointer // The address of the relocatable expression for the item in the file that needs to be updated if the address is changed. For relocatable expressions with the difference of two section addresses, the address from which to subtract (in mathematical terms, the minuend) is contained in the first relocation entry and the address to subtract (the subtrahend) is contained in the second relocation entry.

}

// Searchstate
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/searchstate
type Searchstate struct {
	Ss_union_flags unsafe.Pointer
}

// Section
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/section
type Section struct {
	Sectname  unsafe.Pointer
	Reloff    unsafe.Pointer
	Segname   unsafe.Pointer
	Addr      unsafe.Pointer
	Reserved2 unsafe.Pointer
	Nreloc    unsafe.Pointer
	Flags     unsafe.Pointer
	Align     unsafe.Pointer
	Offset    unsafe.Pointer
	Reserved1 unsafe.Pointer
	Size      unsafe.Pointer
}

// Section_64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/section_64
type Section_64 struct {
	Nreloc    unsafe.Pointer
	Segname   unsafe.Pointer // A string specifying the name of the segment that should eventually contain this section. For compactness, intermediate object files—files of type `MH_OBJECT`—contain only one segment, in which all sections are placed. The static linker places each section in the named segment when building the final product (any file that is not of type `MH_OBJECT`).
	Flags     unsafe.Pointer
	Reserved1 unsafe.Pointer
	Reserved2 unsafe.Pointer
	Align     unsafe.Pointer
	Offset    unsafe.Pointer
	Size      unsafe.Pointer
	Addr      unsafe.Pointer
	Reloff    unsafe.Pointer
	Sectname  unsafe.Pointer
	Reserved3 unsafe.Pointer
}

// Segment_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/segment_command
type Segment_command struct {
	Initprot Vm_prot_t
	Flags    unsafe.Pointer
	Cmd      unsafe.Pointer // Common to all load command structures. Set to `LC_SEGMENT` for this structure.
	Vmsize   unsafe.Pointer // Indicates the number of bytes of virtual memory occupied by this segment. See also the description of `filesize`, below.
	Fileoff  unsafe.Pointer // Indicates the offset in this file of the data to be mapped at `vmaddr`.
	Cmdsize  unsafe.Pointer
	Segname  unsafe.Pointer
	Filesize unsafe.Pointer
	Vmaddr   unsafe.Pointer
	Nsects   unsafe.Pointer
	Maxprot  Vm_prot_t
}

// Segment_command_64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/segment_command_64
type Segment_command_64 struct {
	Cmd      unsafe.Pointer
	Maxprot  Vm_prot_t
	Nsects   unsafe.Pointer // Indicates the number of section data structures following this load command.
	Filesize unsafe.Pointer // Indicates the number of bytes occupied by this segment on disk. For segments that require more memory at runtime than they do at build time, `vmsize` can be larger than `filesize`. For example, the `__PAGEZERO` segment generated by the linker for `MH_EXECUTABLE` files has a `vmsize` of 0x1000 but a `filesize` of 0. Because `__PAGEZERO` contains no data, there is no need for it to occupy any space until runtime. Also, the static linker often allocates uninitialized data at the end of the `__DATA` segment; in this case, the `vmsize` is larger than the `filesize`. The loader guarantees that any memory of this sort is initialized with zeros.
	Fileoff  unsafe.Pointer
	Vmsize   unsafe.Pointer // Indicates the number of bytes of virtual memory occupied by this segment. See also the description of `filesize`, below.
	Segname  unsafe.Pointer
	Flags    unsafe.Pointer
	Cmdsize  unsafe.Pointer
	Initprot Vm_prot_t
	Vmaddr   unsafe.Pointer
}

// Sem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sem
type Sem struct {
	Semncnt unsafe.Pointer
	Semval  unsafe.Pointer
	Semzcnt unsafe.Pointer
	Sempid  int32
}

// Sembuf
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sembuf
type Sembuf struct {
	Sem_flg unsafe.Pointer
	Sem_op  unsafe.Pointer
	Sem_num unsafe.Pointer
}

// Sf_hdtr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sf_hdtr
type Sf_hdtr struct {
	Headers  *Iovec
	Trailers *Iovec
	Hdr_cnt  unsafe.Pointer
	Trl_cnt  unsafe.Pointer
}

// Sflt_filter
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sflt_filter
type Sflt_filter struct {
	Sf_data_out     unsafe.Pointer // Your function for handling outgoing data. May be null.
	Sf_attach       unsafe.Pointer // Your function for handling attaches to sockets.
	Sf_data_in      unsafe.Pointer // Your function for handling incoming data. May be null.
	Sf_ext          unsafe.Pointer
	Sf_getsockname  unsafe.Pointer
	Sf_handle       Sflt_handle    // A value used to find socket filters by applications. An application can use this value to specify that this filter should be attached when using the SO_NKE socket option.
	Sf_getoption    unsafe.Pointer // Your function for handling getsockopt. May be null.
	Sf_unregistered unsafe.Pointer // Your function for being notified when your filter has been unregistered.
	Sf_flags        unsafe.Pointer // Indicate whether this filter should be attached to all new sockets or just those that request the filter be attached using the SO_NKE socket option. If this filter utilizes the socket filter extension fields, it must also set SFLT_EXTENDED.
	Sf_listen       unsafe.Pointer // Your function for handling listen. May be null.
	Sf_notify       unsafe.Pointer // Your function for handling events. May be null.
	Sf_setoption    unsafe.Pointer // Your function for handling setsockopt. May be null.
	Sf_bind         unsafe.Pointer // Your function for handling binds. May be null.
	Sf_name         unsafe.Pointer // A name used for debug purposes.
	Sf_getpeername  unsafe.Pointer
	Sf_detach       unsafe.Pointer // Your function for handling detaches from sockets.
	Sf_connect_out  unsafe.Pointer // Your function for handling outbound connections. May be null.
	Sf_connect_in   unsafe.Pointer // Your function for handling inbound connections. May be null.
	Sf_ioctl        unsafe.Pointer // Your function for handling ioctls. May be null.

}

// Sgttyb
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sgttyb
type Sgttyb struct {
	Sg_kill   unsafe.Pointer
	Sg_erase  unsafe.Pointer
	Sg_flags  unsafe.Pointer
	Sg_ospeed unsafe.Pointer
	Sg_ispeed unsafe.Pointer
}

// Shared_file_mapping_np
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/shared_file_mapping_np
type Shared_file_mapping_np struct {
	Sfm_max_prot    Vm_prot_t
	Sfm_init_prot   Vm_prot_t
	Sfm_size        Mach_vm_size_t
	Sfm_file_offset Mach_vm_offset_t
	Sfm_address     Mach_vm_address_t
}

// Shared_file_np
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/shared_file_np
type Shared_file_np struct {
	Sf_fd             unsafe.Pointer
	Sf_mappings_count unsafe.Pointer
	Sf_slide          unsafe.Pointer
}

// Shared_region_range_np
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/shared_region_range_np
type Shared_region_range_np struct {
	Srr_address Mach_vm_address_t
	Srr_size    Mach_vm_size_t
}

// Sigaction
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sigaction
type Sigaction struct {
	__sigaction_u unsafe.Pointer
	Sa_mask       Sigset_t
	Sa_flags      int
}

// Sigstack
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sigstack
type Sigstack struct {
	Ss_onstack int
	Ss_sp      *byte
}

// Sigvec
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sigvec
type Sigvec struct {
	Sv_mask    int
	Sv_flags   int
	Sv_handler unsafe.Pointer
}

// Smrq_link
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/smrq_link
type Smrq_link struct {
	Next unsafe.Pointer
	Prev unsafe.Pointer
}

// Smrq_list_head
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/smrq_list_head
type Smrq_list_head struct {
	First unsafe.Pointer
}

// Smrq_slink
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/smrq_slink
type Smrq_slink struct {
	Next unsafe.Pointer
}

// Smrq_slist_head
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/smrq_slist_head
type Smrq_slist_head struct {
	First unsafe.Pointer
}

// Smrq_stailq_head
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/smrq_stailq_head
type Smrq_stailq_head struct {
	First unsafe.Pointer
	Last  unsafe.Pointer
}

// Smrq_tailq_head
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/smrq_tailq_head
type Smrq_tailq_head struct {
	First unsafe.Pointer
	Last  unsafe.Pointer
}

// So_nke
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/so_nke
type So_nke struct {
	Reserved   U_int32_t
	Nke_flags  unsafe.Pointer
	Nke_handle unsafe.Pointer
	Nke_where  unsafe.Pointer
}

// So_np_extensions
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/so_np_extensions
type So_np_extensions struct {
	Npx_mask  U_int32_t
	Npx_flags U_int32_t
}

// Sockaddr_ctl
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sockaddr_ctl
type Sockaddr_ctl struct {
	Sc_len    U_char // The length of the structure.
	Sc_family U_char // AF_SYSTEM.

}

// Sockaddr_dl
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sockaddr_dl
type Sockaddr_dl struct {
	Sdl_slen   U_char
	Sdl_family U_char
	Sdl_data   unsafe.Pointer
	Sdl_nlen   U_char
	Sdl_index  U_short
	Sdl_type   U_char
	Sdl_len    U_char
	Sdl_alen   U_char
}

// Sockaddr_in
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sockaddr_in
type Sockaddr_in struct {
	Sin_addr   In_addr
	Sin_len    uint8
	Sin_family uint8
	Sin_zero   unsafe.Pointer
	Sin_port   uint16
}

// Sockaddr_in6
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sockaddr_in6
type Sockaddr_in6 struct {
	Sin6_scope_id uint32
	Sin6_flowinfo uint32
	Sin6_len      uint8
	Sin6_addr     unsafe.Pointer
	Sin6_port     uint16
	Sin6_family   uint8
}

// Sockaddr_inarp
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sockaddr_inarp
type Sockaddr_inarp struct {
	Sin_len     U_char
	Sin_srcaddr In_addr
	Sin_family  U_char
	Sin_addr    In_addr
	Sin_port    U_short
	Sin_tos     U_short
	Sin_other   U_short
}

// Sockaddr_inifscope
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sockaddr_inifscope
type Sockaddr_inifscope struct {
	Sin_addr   In_addr
	Sin_family uint8
	Sin_len    uint8
	Sin_port   uint16
	Un         unsafe.Pointer
}

// Sockaddr_ndrv
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sockaddr_ndrv
type Sockaddr_ndrv struct {
	Snd_family unsafe.Pointer
	Snd_len    unsafe.Pointer
	Snd_name   unsafe.Pointer
}

// Sockaddr_sys
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sockaddr_sys
type Sockaddr_sys struct {
	Ss_reserved U_int32_t
	Ss_len      U_char
	Ss_family   U_char
	Ss_sysaddr  U_int16_t
}

// Sockaddr_un
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sockaddr_un
type Sockaddr_un struct {
	Sun_len    unsafe.Pointer
	Sun_path   unsafe.Pointer
	Sun_family uint8
}

// Sockaddr_vm
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sockaddr_vm
type Sockaddr_vm struct {
	Svm_cid       uint32
	Svm_family    uint8
	Svm_len       uint8
	Svm_port      uint32
	Svm_reserved1 uint16
}

// Sockproto
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sockproto
type Sockproto struct {
	Sp_family uint16
}

// Source_version_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/source_version_command
type Source_version_command struct {
	Cmd     unsafe.Pointer
	Cmdsize unsafe.Pointer
	Version unsafe.Pointer
}

// Specinfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/specinfo
type Specinfo struct {
	Si_isssd         U_int16_t
	Si_hashchain     unsafe.Pointer
	Si_specnext      unsafe.Pointer
	Si_rdev          int32
	Si_size          Daddr_t
	Si_lastr         Daddr64_t
	Si_flags         unsafe.Pointer
	Si_throttle_mask U_int64_t
	Si_opencount     unsafe.Pointer
	Si_devsize       U_int64_t
	Si_throttleable  U_int8_t
	Si_initted       U_int8_t
	Si_devbsdunit    U_int32_t
	Si_mountingowner Thread_t
}

// Stack_snapshot_frame32
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/stack_snapshot_frame32
type Stack_snapshot_frame32 struct {
	Sp unsafe.Pointer
	Lr unsafe.Pointer
}

// Stack_snapshot_frame64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/stack_snapshot_frame64
type Stack_snapshot_frame64 struct {
	Sp unsafe.Pointer
	Lr unsafe.Pointer
}

// Stack_snapshot_stacktop
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/stack_snapshot_stacktop
type Stack_snapshot_stacktop struct {
	Sp             unsafe.Pointer
	Stack_contents uint8
}

// Stackshot_cpu_architecture
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/stackshot_cpu_architecture
type Stackshot_cpu_architecture struct {
	Cpusubtype unsafe.Pointer
}

// Stackshot_cpu_times
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/stackshot_cpu_times
type Stackshot_cpu_times struct {
	User_usec   unsafe.Pointer
	System_usec unsafe.Pointer
}

// Stackshot_cpu_times_v2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/stackshot_cpu_times_v2
type Stackshot_cpu_times_v2 struct {
	Runnable_usec unsafe.Pointer
	System_usec   unsafe.Pointer
	User_usec     unsafe.Pointer
}

// Stackshot_duration
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/stackshot_duration
type Stackshot_duration struct {
	Stackshot_duration_outer unsafe.Pointer
	Stackshot_duration       unsafe.Pointer
}

// Stackshot_duration_v2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/stackshot_duration_v2
type Stackshot_duration_v2 struct {
	Stackshot_duration       unsafe.Pointer
	Stackshot_duration_outer unsafe.Pointer
	Stackshot_duration_prior unsafe.Pointer
}

// Stackshot_fault_stats
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/stackshot_fault_stats
type Stackshot_fault_stats struct {
	Sfs_system_max_fault_time unsafe.Pointer
	Sfs_stopped_faulting      uint8
	Sfs_time_spent_faulting   unsafe.Pointer
	Sfs_pages_faulted_in      unsafe.Pointer
}

// Stackshot_latency_collection
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/stackshot_latency_collection
type Stackshot_latency_collection struct {
	Setup_latency                           unsafe.Pointer
	Total_task_iteration_latency            unsafe.Pointer
	Total_terminated_task_iteration_latency unsafe.Pointer
}

// Stackshot_latency_collection_v2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/stackshot_latency_collection_v2
type Stackshot_latency_collection_v2 struct {
	Buffer_count                               unsafe.Pointer
	Buffer_overhead                            unsafe.Pointer
	Buffer_size                                unsafe.Pointer
	Buffer_used                                unsafe.Pointer
	Calling_cpu_number                         unsafe.Pointer
	Cpu_wait_latency_mt                        unsafe.Pointer
	Main_cpu_number                            unsafe.Pointer
	Setup_latency_mt                           unsafe.Pointer
	Task_queue_building_latency_mt             unsafe.Pointer
	Terminated_task_queue_building_latency_mt  unsafe.Pointer
	Total_task_iteration_latency_mt            unsafe.Pointer
	Total_terminated_task_iteration_latency_mt unsafe.Pointer
}

// Stackshot_latency_cpu
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/stackshot_latency_cpu
type Stackshot_latency_cpu struct {
	Cluster_type          unsafe.Pointer
	Cpu_number            unsafe.Pointer
	Faulting_time_mt      unsafe.Pointer
	Init_latency_mt       unsafe.Pointer
	Intercluster_buf_used unsafe.Pointer
	Tasks_processed       unsafe.Pointer
	Threads_processed     unsafe.Pointer
	Total_buf             unsafe.Pointer
	Total_cycles          unsafe.Pointer
	Total_instrs          unsafe.Pointer
	Total_latency_mt      unsafe.Pointer
	Workqueue_latency_mt  unsafe.Pointer
}

// Stackshot_latency_task
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/stackshot_latency_task
type Stackshot_latency_task struct {
	Bsd_proc_ids_latency           unsafe.Pointer
	Cur_tsnap_latency              unsafe.Pointer
	End_latency                    unsafe.Pointer
	Misc2_latency                  unsafe.Pointer
	Misc_latency                   unsafe.Pointer
	Pmap_latency                   unsafe.Pointer
	Setup_latency                  unsafe.Pointer
	Task_thread_count_loop_latency unsafe.Pointer
	Task_thread_data_loop_latency  unsafe.Pointer
	Task_uniqueid                  unsafe.Pointer
}

// Stackshot_latency_thread
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/stackshot_latency_thread
type Stackshot_latency_thread struct {
	Cur_thsnap1_latency     unsafe.Pointer
	Cur_thsnap2_latency     unsafe.Pointer
	Dispatch_label_latency  unsafe.Pointer
	Dispatch_serial_latency unsafe.Pointer
	Kernel_stack_latency    unsafe.Pointer
	Misc_latency            unsafe.Pointer
	Sur_times_latency       unsafe.Pointer
	Thread_id               unsafe.Pointer
	Thread_name_latency     unsafe.Pointer
	User_stack_latency      unsafe.Pointer
}

// Stackshot_suspension_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/stackshot_suspension_info
type Stackshot_suspension_info struct {
	Tss_count      unsafe.Pointer
	Tss_duration   unsafe.Pointer
	Tss_last_end   unsafe.Pointer
	Tss_last_start unsafe.Pointer
}

// Stackshot_suspension_source
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/stackshot_suspension_source
type Stackshot_suspension_source struct {
	Tss_pid      unsafe.Pointer
	Tss_procname unsafe.Pointer
	Tss_tid      unsafe.Pointer
	Tss_time     unsafe.Pointer
}

// Stackshot_task_codesigning_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/stackshot_task_codesigning_info
type Stackshot_task_codesigning_info struct {
	Cs_trust_level unsafe.Pointer
	Csflags        unsafe.Pointer
}

// Stat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/stat
type Stat struct {
	St_dev       int32
	St_atimespec Timespec
	St_nlink     uint16
	St_size      int64
	St_gid       uint32
	St_uid       uint32
	St_lspare    int32
	St_ino       uint64
	St_mtimespec Timespec
	St_rdev      int32
	St_qspare    int64
	St_mode      uint16
	St_gen       uint32
	St_flags     uint32
	St_ctimespec Timespec
	St_blocks    int64
	St_blksize   int32
}

// Statfs64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/statfs64
type Statfs64 struct {
	F_bfree       unsafe.Pointer
	F_fsid        Fsid_t
	F_iosize      unsafe.Pointer
	F_flags       unsafe.Pointer
	F_type        unsafe.Pointer
	F_mntonname   unsafe.Pointer
	F_owner       uint32
	F_bsize       unsafe.Pointer
	F_blocks      unsafe.Pointer
	F_mntfromname unsafe.Pointer
	F_reserved    unsafe.Pointer
	F_ffree       unsafe.Pointer
	F_fstypename  unsafe.Pointer
	F_fssubtype   unsafe.Pointer
	F_files       unsafe.Pointer
	F_bavail      unsafe.Pointer
	F_flags_ext   unsafe.Pointer
}

// Sub_client_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sub_client_command
type Sub_client_command struct {
	Cmd     unsafe.Pointer
	Client  unsafe.Pointer
	Cmdsize unsafe.Pointer
}

// Sub_framework_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sub_framework_command
type Sub_framework_command struct {
	Cmd      unsafe.Pointer
	Umbrella unsafe.Pointer
	Cmdsize  unsafe.Pointer
}

// Sub_library_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sub_library_command
type Sub_library_command struct {
	Cmd         unsafe.Pointer
	Cmdsize     unsafe.Pointer
	Sub_library unsafe.Pointer
}

// Sub_umbrella_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sub_umbrella_command
type Sub_umbrella_command struct {
	Cmd          unsafe.Pointer
	Sub_umbrella unsafe.Pointer // A data structure of type `lc_str`. Specifies the name of the umbrella framework of which this file is a member.
	Cmdsize      unsafe.Pointer
}

// Symseg_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/symseg_command
type Symseg_command struct {
	Cmdsize unsafe.Pointer
	Offset  unsafe.Pointer
	Size    unsafe.Pointer
	Cmd     unsafe.Pointer
}

// Symtab_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/symtab_command
type Symtab_command struct {
	Nsyms   unsafe.Pointer
	Cmd     unsafe.Pointer
	Cmdsize unsafe.Pointer
	Strsize unsafe.Pointer
	Symoff  unsafe.Pointer
	Stroff  unsafe.Pointer
}

// Sysctl_oid
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sysctl_oid
type Sysctl_oid struct {
	Oid_name    unsafe.Pointer
	Oid_refcnt  unsafe.Pointer
	Oid_handler *Sysctl_req
	Oid_link    unsafe.Pointer
	Oid_arg2    unsafe.Pointer
	Oid_fmt     unsafe.Pointer
	Oid_descr   unsafe.Pointer
	Oid_parent  unsafe.Pointer
	Oid_number  unsafe.Pointer
	Oid_arg1    unsafe.Pointer
	Oid_kind    unsafe.Pointer
}

// Sysctl_oid_list
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sysctl_oid_list
type Sysctl_oid_list struct {
	Slh_first *Sysctl_oid
}

// Sysctl_req
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/sysctl_req
type Sysctl_req struct {
	Oldfunc uintptr
	Oldlen  uintptr
	Lock    unsafe.Pointer
	Newidx  uintptr
	Newlen  uintptr
	Newfunc uintptr
	P       unsafe.Pointer
	Oldidx  uintptr
	Oldptr  User_addr_t
	Newptr  User_addr_t
}

// TPacketFilterMetadata
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/tpacketfiltermetadata
type TPacketFilterMetadata struct {
	DeviceAddress       uint8
	EndpointAddress     uint8
	EndpointType        uint8
	InterfaceAltSetting uint8
	InterfaceClass      uint8
	InterfaceProtocol   uint8
	InterfaceSubclass   uint8
	LocationID          unsafe.Pointer
	Pid                 uint16
	Speed               uint8
	Vid                 uint16
}

// Task_access_subsystem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/task_access_subsystem
type Task_access_subsystem struct {
	Maxsize  unsafe.Pointer
	Server   unsafe.Pointer
	Routine  unsafe.Pointer
	End      int32
	Start    int32
	Reserved Vm_address_t
}

// Task_delta_snapshot_v2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/task_delta_snapshot_v2
type Task_delta_snapshot_v2 struct {
	Tds_did_throttle                      unsafe.Pointer
	Tds_ss_flags                          unsafe.Pointer
	Tds_user_time_in_terminated_threads   unsafe.Pointer
	Tds_task_size                         unsafe.Pointer
	Tds_latency_qos                       unsafe.Pointer
	Tds_pageins                           unsafe.Pointer
	Tds_unique_pid                        unsafe.Pointer
	Tds_max_resident_size                 unsafe.Pointer
	Tds_cow_faults                        unsafe.Pointer
	Tds_system_time_in_terminated_threads unsafe.Pointer
	Tds_faults                            unsafe.Pointer
	Tds_was_throttled                     unsafe.Pointer
	Tds_suspend_count                     unsafe.Pointer
}

// Task_snapshot
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/task_snapshot
type Task_snapshot struct {
	Suspend_count                     unsafe.Pointer
	Latency_qos                       unsafe.Pointer
	Snapshot_magic                    unsafe.Pointer
	P_start_sec                       unsafe.Pointer
	User_time_in_terminated_threads   unsafe.Pointer
	Disk_writes_count                 unsafe.Pointer
	Pid                               unsafe.Pointer
	Paging_count                      unsafe.Pointer
	P_start_usec                      unsafe.Pointer
	Nloadinfos                        unsafe.Pointer
	Did_throttle                      unsafe.Pointer
	Disk_writes_size                  unsafe.Pointer
	Task_size                         unsafe.Pointer
	Shared_cache_slide                unsafe.Pointer
	P_comm                            unsafe.Pointer
	Cow_faults                        unsafe.Pointer
	Non_paging_size                   unsafe.Pointer
	Metadata_size                     unsafe.Pointer
	Was_throttled                     unsafe.Pointer
	Disk_reads_count                  unsafe.Pointer
	Io_priority_size                  unsafe.Pointer
	Uniqueid                          unsafe.Pointer
	Disk_reads_size                   unsafe.Pointer
	Donating_pid_count                unsafe.Pointer
	Non_paging_count                  unsafe.Pointer
	Faults                            unsafe.Pointer
	Paging_size                       unsafe.Pointer
	System_time_in_terminated_threads unsafe.Pointer
	Pageins                           unsafe.Pointer
	Data_size                         unsafe.Pointer
	Data_count                        unsafe.Pointer
	Io_priority_count                 unsafe.Pointer
	Ss_flags                          unsafe.Pointer
	Metadata_count                    unsafe.Pointer
	Shared_cache_identifier           uint8
}

// Task_snapshot_v2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/task_snapshot_v2
type Task_snapshot_v2 struct {
	Ts_pid                               unsafe.Pointer
	Ts_ss_flags                          unsafe.Pointer
	Ts_suspend_count                     unsafe.Pointer
	Ts_cow_faults                        unsafe.Pointer
	Ts_faults                            unsafe.Pointer
	Ts_max_resident_size                 unsafe.Pointer
	Ts_unique_pid                        unsafe.Pointer
	Ts_user_time_in_terminated_threads   unsafe.Pointer
	Ts_did_throttle                      unsafe.Pointer
	Ts_latency_qos                       unsafe.Pointer
	Ts_p_start_sec                       unsafe.Pointer
	Ts_pageins                           unsafe.Pointer
	Ts_system_time_in_terminated_threads unsafe.Pointer
	Ts_p_comm                            unsafe.Pointer
	Ts_was_throttled                     unsafe.Pointer
	Ts_task_size                         unsafe.Pointer
}

// Tchars
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/tchars
type Tchars struct {
	T_stopc unsafe.Pointer
}

// Tcp_conn_status
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/tcp_conn_status
type Tcp_conn_status struct {
}

// Tcp_connection_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/tcp_connection_info
type Tcp_connection_info struct {
	Tcpi_tfo_cookie_req_rcv     U_int32_t
	Tcpi_maxseg                 U_int32_t
	Tcpi_rto                    U_int32_t
	Tcpi_rttcur                 U_int32_t
	Tcpi_tfo_cookie_sent        U_int32_t
	Tcpi_txbytes                U_int64_t
	Tcpi_tfo_syn_data_acked     U_int32_t
	Tcpi_tfo_cookie_rcv         U_int32_t
	Tcpi_tfo_syn_loss           U_int32_t
	Tcpi_tfo_cookie_req         U_int32_t
	Tcpi_options                U_int32_t
	Tcpi_rcv_wscale             U_int8_t
	Tcpi_srtt                   U_int32_t
	Tcpi_rxpackets              U_int64_t
	Tcpi_txretransmitbytes      U_int64_t
	Tcpi_snd_wscale             U_int8_t
	Tcpi_rxbytes                U_int64_t
	Tcpi_tfo_syn_data_rcv       U_int32_t
	Tcpi_snd_ssthresh           U_int32_t
	Tcpi_tfo_cookie_invalid     U_int32_t
	Tcpi_rttvar                 U_int32_t
	Tcpi_flags                  U_int32_t
	Tcpi_state                  U_int8_t
	Tcpi_txpackets              U_int64_t
	Tcpi_snd_sbbytes            U_int32_t
	Tcpi_tfo_syn_data_sent      U_int32_t
	Tcpi_snd_wnd                U_int32_t
	Tcpi_snd_cwnd               U_int32_t
	Tcpi_rxoutoforderbytes      U_int64_t
	Tcpi_rcv_wnd                U_int32_t
	Tcpi_tfo_cookie_wrong       U_int32_t
	Tcpi_tfo_heuristics_disable U_int32_t
	Tcpi_tfo_send_blackhole     U_int32_t
	Tcpi_tfo_no_cookie_rcv      U_int32_t
	Tcpi_tfo_recv_blackhole     U_int32_t
	Tcpi_txretransmitpackets    U_int64_t
	Tcpi_tfo_onebyte_proxy      U_int32_t
}

// Tcp_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/tcp_info
type Tcp_info struct {
	Tcpi_cell_rxbytes               U_int64_t
	Tcpi_cell_rxpackets             U_int64_t
	Tcpi_cell_txbytes               U_int64_t
	Tcpi_cell_txpackets             U_int64_t
	Tcpi_connstatus                 Tcp_conn_status
	Tcpi_dsack_recvd                U_int32_t
	Tcpi_dsack_sent                 U_int32_t
	Tcpi_ecn_client_setup           U_int16_t
	Tcpi_ecn_lost_syn               U_int16_t
	Tcpi_ecn_lost_synack            U_int16_t
	Tcpi_ecn_recv_ce                U_int32_t
	Tcpi_ecn_recv_cwr               U_int32_t
	Tcpi_ecn_server_setup           U_int16_t
	Tcpi_ecn_success                U_int16_t
	Tcpi_flags                      U_int32_t
	Tcpi_flowhash                   U_int32_t
	Tcpi_if_cell                    U_int16_t
	Tcpi_if_wifi                    U_int16_t
	Tcpi_if_wifi_awdl               U_int16_t
	Tcpi_if_wifi_infra              U_int16_t
	Tcpi_if_wired                   U_int16_t
	Tcpi_last_outif                 unsafe.Pointer
	Tcpi_options                    U_int8_t
	Tcpi_pawsdrop                   U_int32_t
	Tcpi_rcv_background             U_int16_t
	Tcpi_rcv_mss                    U_int32_t
	Tcpi_rcv_nxt                    U_int32_t
	Tcpi_rcv_space                  U_int32_t
	Tcpi_rcv_srtt                   unsafe.Pointer
	Tcpi_rcv_wscale                 U_int8_t
	Tcpi_rcvoopack                  U_int32_t
	Tcpi_reordered_pkts             U_int32_t
	Tcpi_rto                        U_int32_t
	Tcpi_rttbest                    U_int32_t
	Tcpi_rttcur                     U_int32_t
	Tcpi_rttvar                     U_int32_t
	Tcpi_rxbytes                    U_int64_t
	Tcpi_rxduplicatebytes           U_int64_t
	Tcpi_rxoutoforderbytes          U_int64_t
	Tcpi_rxpackets                  U_int64_t
	Tcpi_sack_recovery_episode      U_int32_t
	Tcpi_snd_background             U_int16_t
	Tcpi_snd_bw                     U_int64_t
	Tcpi_snd_cwnd                   U_int32_t
	Tcpi_snd_mss                    U_int32_t
	Tcpi_snd_nxt                    U_int32_t
	Tcpi_snd_sbbytes                U_int32_t
	Tcpi_snd_ssthresh               U_int32_t
	Tcpi_snd_wnd                    U_int32_t
	Tcpi_snd_wscale                 U_int8_t
	Tcpi_srtt                       U_int32_t
	Tcpi_state                      U_int8_t
	Tcpi_synrexmits                 U_int8_t
	Tcpi_tfo_cookie_invalid         U_int16_t
	Tcpi_tfo_cookie_rcv             U_int16_t
	Tcpi_tfo_cookie_req             U_int16_t
	Tcpi_tfo_cookie_req_rcv         U_int16_t
	Tcpi_tfo_cookie_sent            U_int16_t
	Tcpi_tfo_cookie_wrong           U_int16_t
	Tcpi_tfo_heuristics_disable     U_int16_t
	Tcpi_tfo_no_cookie_rcv          U_int16_t
	Tcpi_tfo_onebyte_proxy          U_int16_t
	Tcpi_tfo_recv_blackhole         U_int16_t
	Tcpi_tfo_send_blackhole         U_int16_t
	Tcpi_tfo_syn_data_acked         U_int16_t
	Tcpi_tfo_syn_data_rcv           U_int16_t
	Tcpi_tfo_syn_data_sent          U_int16_t
	Tcpi_tfo_syn_loss               U_int16_t
	Tcpi_txbytes                    U_int64_t
	Tcpi_txpackets                  U_int64_t
	Tcpi_txretransmitbytes          U_int64_t
	Tcpi_txretransmitpackets        U_int64_t
	Tcpi_txunacked                  U_int64_t
	Tcpi_unused1                    U_int8_t
	Tcpi_unused2                    U_int16_t
	Tcpi_wifi_rxbytes               U_int64_t
	Tcpi_wifi_rxpackets             U_int64_t
	Tcpi_wifi_txbytes               U_int64_t
	Tcpi_wifi_txpackets             U_int64_t
	Tcpi_wired_rxbytes              U_int64_t
	Tcpi_wired_rxpackets            U_int64_t
	Tcpi_wired_txbytes              U_int64_t
	Tcpi_wired_txpackets            U_int64_t
	Tcpi_ecn_capable_packets_acked  unsafe.Pointer
	Tcpi_ecn_capable_packets_lost   unsafe.Pointer
	Tcpi_ecn_capable_packets_marked unsafe.Pointer
	Tcpi_client_accecn_state        unsafe.Pointer
	Tcpi_ecn_capable_packets_sent   unsafe.Pointer
	Tcpi_server_accecn_state        unsafe.Pointer
	Tcpi_flow_control_total_time    unsafe.Pointer
	Tcpi_rcvwnd_limited_total_time  unsafe.Pointer
	Tcpi_delivered_ce_bytes         unsafe.Pointer
	Tcpi_delivered_ect0_bytes       unsafe.Pointer
	Tcpi_delivered_ect1_bytes       unsafe.Pointer
	Tcpi_received_ce_bytes          unsafe.Pointer
	Tcpi_received_ce_packets        unsafe.Pointer
	Tcpi_received_ect0_bytes        unsafe.Pointer
	Tcpi_received_ect1_bytes        unsafe.Pointer
	Tcpi_l4s_enabled                U_int16_t
}

// Tcp_measure_bw_burst
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/tcp_measure_bw_burst
type Tcp_measure_bw_burst struct {
	Max_burst_size U_int32_t
	Min_burst_size U_int32_t
}

// Tcp_notify_ack_complete
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/tcp_notify_ack_complete
type Tcp_notify_ack_complete struct {
	Notify_complete_count U_int32_t
	Notify_complete_id    Tcp_notify_ack_id_t
	Notify_pending        U_int32_t
}

// Tcpcb
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/tcpcb
type Tcpcb struct {
	Cc_recv Tcp_cc
	Cc_send Tcp_cc
	Irs     Tcp_seq
	Rcv_nxt Tcp_seq
}

// Tcphdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/tcphdr
type Tcphdr struct {
	Th_urp   unsafe.Pointer
	Th_seq   Tcp_seq
	Th_x2    unsafe.Pointer
	Th_sport unsafe.Pointer
	Th_win   unsafe.Pointer
	Th_dport unsafe.Pointer
	Th_flags unsafe.Pointer
	Th_ack   Tcp_seq
	Th_sum   unsafe.Pointer
	Th_off   unsafe.Pointer
}

// Tcpiphdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/tcpiphdr
type Tcpiphdr struct {
	Ti_i Ipovly
	Ti_t Tcphdr
}

// Tcpstat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/tcpstat
type Tcpstat struct {
	Tcps_accepts                                  U_int32_t
	Tcps_avoid_rxmt                               U_int32_t
	Tcps_badrst                                   U_int32_t
	Tcps_badsyn                                   U_int32_t
	Tcps_bg_rcvtotal                              U_int32_t
	Tcps_cachedrtt                                U_int32_t
	Tcps_cachedrttvar                             U_int32_t
	Tcps_cachedssthresh                           U_int32_t
	Tcps_closed                                   U_int32_t
	Tcps_connattempt                              U_int32_t
	Tcps_conndrops                                U_int32_t
	Tcps_connects                                 U_int32_t
	Tcps_delack                                   U_int32_t
	Tcps_delay_recovery                           U_int32_t
	Tcps_detect_reordering                        U_int32_t
	Tcps_drop_after_sleep                         U_int32_t
	Tcps_drops                                    U_int32_t
	Tcps_dsack_ackloss                            U_int32_t
	Tcps_dsack_badrexmt                           U_int32_t
	Tcps_dsack_disable                            U_int32_t
	Tcps_dsack_recvd                              U_int32_t
	Tcps_dsack_recvd_old                          U_int32_t
	Tcps_dsack_sent                               U_int32_t
	Tcps_early_rexmt                              U_int32_t
	Tcps_ecn_client_setup                         U_int32_t
	Tcps_ecn_client_success                       U_int32_t
	Tcps_ecn_conn_nopl_ce                         U_int32_t
	Tcps_ecn_conn_pl_ce                           U_int32_t
	Tcps_ecn_conn_plnoce                          U_int32_t
	Tcps_ecn_conn_recv_ce                         U_int32_t
	Tcps_ecn_conn_recv_ece                        U_int32_t
	Tcps_ecn_fallback_ce                          U_int32_t
	Tcps_ecn_fallback_droprst                     U_int32_t
	Tcps_ecn_fallback_droprxmt                    U_int32_t
	Tcps_ecn_fallback_reorder                     U_int32_t
	Tcps_ecn_fallback_synloss                     U_int32_t
	Tcps_ecn_fallback_synrst                      U_int32_t
	Tcps_ecn_lost_syn                             U_int32_t
	Tcps_ecn_lost_synack                          U_int32_t
	Tcps_ecn_not_supported                        U_int32_t
	Tcps_ecn_recv_ce                              U_int32_t
	Tcps_ecn_recv_ece                             U_int32_t
	Tcps_ecn_sent_ece                             U_int32_t
	Tcps_ecn_server_setup                         U_int32_t
	Tcps_ecn_server_success                       U_int32_t
	Tcps_estab_fallback                           U_int32_t
	Tcps_fcholdpacket                             U_int32_t
	Tcps_hc_added                                 U_int32_t
	Tcps_hc_bucketoverflow                        U_int32_t
	Tcps_invalid_joins                            U_int32_t
	Tcps_invalid_mpcap                            U_int32_t
	Tcps_invalid_opt                              U_int32_t
	Tcps_join_fallback                            U_int32_t
	Tcps_join_rxmts                               U_int32_t
	Tcps_ka_offload_drops                         U_int32_t
	Tcps_keepdrops                                U_int32_t
	Tcps_keepprobe                                U_int32_t
	Tcps_keeptimeo                                U_int32_t
	Tcps_limited_txt                              U_int32_t
	Tcps_listendrop                               U_int32_t
	Tcps_minmssdrops                              U_int32_t
	Tcps_mp_badcsum                               U_int32_t
	Tcps_mp_num_probes                            U_int32_t
	Tcps_mp_oodata                                U_int32_t
	Tcps_mp_outofwin                              U_int32_t
	Tcps_mp_rcvbytes                              U_int32_t
	Tcps_mp_rcvtotal                              U_int32_t
	Tcps_mp_reducedwin                            U_int32_t
	Tcps_mp_sel_peer                              U_int32_t
	Tcps_mp_sel_rto                               U_int32_t
	Tcps_mp_sel_rtt                               U_int32_t
	Tcps_mp_sel_symtomsd                          U_int32_t
	Tcps_mp_sndbytes                              U_int32_t
	Tcps_mp_sndpacks                              U_int32_t
	Tcps_mp_switches                              U_int32_t
	Tcps_mp_verdowngrade                          U_int32_t
	Tcps_mpcap_fallback                           U_int32_t
	Tcps_mptcp_aggregate_all_bytes                U_int64_t
	Tcps_mptcp_aggregate_attempt                  U_int32_t
	Tcps_mptcp_aggregate_cell_bytes               U_int64_t
	Tcps_mptcp_aggregate_success                  U_int32_t
	Tcps_mptcp_back_to_wifi                       U_int32_t
	Tcps_mptcp_cell_proxy                         U_int32_t
	Tcps_mptcp_fp_aggregate_attempt               U_int32_t
	Tcps_mptcp_fp_aggregate_success               U_int32_t
	Tcps_mptcp_fp_handover_attempt                U_int32_t
	Tcps_mptcp_fp_handover_success_cell           U_int32_t
	Tcps_mptcp_fp_handover_success_wifi           U_int32_t
	Tcps_mptcp_fp_heuristic_fallback              U_int32_t
	Tcps_mptcp_fp_interactive_attempt             U_int32_t
	Tcps_mptcp_fp_interactive_success             U_int32_t
	Tcps_mptcp_handover_all_bytes                 U_int64_t
	Tcps_mptcp_handover_attempt                   U_int32_t
	Tcps_mptcp_handover_cell_bytes                U_int64_t
	Tcps_mptcp_handover_cell_from_wifi            U_int32_t
	Tcps_mptcp_handover_success_cell              U_int32_t
	Tcps_mptcp_handover_success_wifi              U_int32_t
	Tcps_mptcp_handover_wifi_from_cell            U_int32_t
	Tcps_mptcp_heuristic_fallback                 U_int32_t
	Tcps_mptcp_interactive_all_bytes              U_int64_t
	Tcps_mptcp_interactive_attempt                U_int32_t
	Tcps_mptcp_interactive_cell_bytes             U_int64_t
	Tcps_mptcp_interactive_cell_from_wifi         U_int32_t
	Tcps_mptcp_interactive_success                U_int32_t
	Tcps_mptcp_rcvduppack                         U_int32_t
	Tcps_mptcp_rcvmemdrop                         U_int32_t
	Tcps_mptcp_rcvpackafterwin                    U_int32_t
	Tcps_mptcp_triggered_cell                     U_int32_t
	Tcps_mptcp_wifi_proxy                         U_int32_t
	Tcps_mss_to_default                           U_int32_t
	Tcps_mss_to_low                               U_int32_t
	Tcps_mss_to_medium                            U_int32_t
	Tcps_mturesent                                U_int32_t
	Tcps_nostretchack                             U_int32_t
	Tcps_pawsdrop                                 U_int32_t
	Tcps_pcbcachemiss                             U_int32_t
	Tcps_persistdrop                              U_int32_t
	Tcps_persisttimeo                             U_int32_t
	Tcps_pmtudbh_reverted                         U_int32_t
	Tcps_predack                                  U_int32_t
	Tcps_preddat                                  U_int32_t
	Tcps_probe_if                                 U_int32_t
	Tcps_probe_if_conflict                        U_int32_t
	Tcps_pto                                      U_int32_t
	Tcps_pto_in_recovery                          U_int32_t
	Tcps_rcv6_swcsum                              U_int32_t
	Tcps_rcv6_swcsum_bytes                        U_int32_t
	Tcps_rcv_swcsum                               U_int32_t
	Tcps_rcv_swcsum_bytes                         U_int32_t
	Tcps_rcvackbyte                               U_int32_t
	Tcps_rcvackpack                               U_int32_t
	Tcps_rcvacktoomuch                            U_int32_t
	Tcps_rcvafterclose                            U_int32_t
	Tcps_rcvbadoff                                U_int32_t
	Tcps_rcvbadsum                                U_int32_t
	Tcps_rcvbyte                                  U_int32_t
	Tcps_rcvbyteafterwin                          U_int32_t
	Tcps_rcvdupack                                U_int32_t
	Tcps_rcvdupbyte                               U_int32_t
	Tcps_rcvduppack                               U_int32_t
	Tcps_rcvmemdrop                               U_int32_t
	Tcps_rcvoobyte                                U_int32_t
	Tcps_rcvoopack                                U_int32_t
	Tcps_rcvpack                                  U_int32_t
	Tcps_rcvpackafterwin                          U_int32_t
	Tcps_rcvpartdupbyte                           U_int32_t
	Tcps_rcvpartduppack                           U_int32_t
	Tcps_rcvshort                                 U_int32_t
	Tcps_rcvtotal                                 U_int32_t
	Tcps_rcvwinprobe                              U_int32_t
	Tcps_rcvwinupd                                U_int32_t
	Tcps_recovered_pkts                           U_int32_t
	Tcps_reordered_pkts                           U_int32_t
	Tcps_rescue_rxmt                              U_int32_t
	Tcps_rexmttimeo                               U_int32_t
	Tcps_rstchallenge                             U_int32_t
	Tcps_rto_after_pto                            U_int32_t
	Tcps_rttupdated                               U_int32_t
	Tcps_rxtfindrop                               U_int32_t
	Tcps_sack_ackadv                              U_int32_t
	Tcps_sack_rcv_blocks                          U_int32_t
	Tcps_sack_recovery_episode                    U_int32_t
	Tcps_sack_rexmit_bytes                        U_int32_t
	Tcps_sack_rexmits                             U_int32_t
	Tcps_sack_sboverflow                          U_int32_t
	Tcps_sack_send_blocks                         U_int32_t
	Tcps_sc_aborted                               U_int32_t
	Tcps_sc_added                                 U_int32_t
	Tcps_sc_badack                                U_int32_t
	Tcps_sc_bucketoverflow                        U_int32_t
	Tcps_sc_cacheoverflow                         U_int32_t
	Tcps_sc_completed                             U_int32_t
	Tcps_sc_dropped                               U_int32_t
	Tcps_sc_dupsyn                                U_int32_t
	Tcps_sc_recvcookie                            U_int32_t
	Tcps_sc_reset                                 U_int32_t
	Tcps_sc_retransmitted                         U_int32_t
	Tcps_sc_sendcookie                            U_int32_t
	Tcps_sc_stale                                 U_int32_t
	Tcps_sc_unreach                               U_int32_t
	Tcps_sc_zonefail                              U_int32_t
	Tcps_segstimed                                U_int32_t
	Tcps_snd6_swcsum                              U_int32_t
	Tcps_snd6_swcsum_bytes                        U_int32_t
	Tcps_snd_swcsum                               U_int32_t
	Tcps_snd_swcsum_bytes                         U_int32_t
	Tcps_sndacks                                  U_int32_t
	Tcps_sndbyte                                  U_int32_t
	Tcps_sndctrl                                  U_int32_t
	Tcps_sndpack                                  U_int32_t
	Tcps_sndprobe                                 U_int32_t
	Tcps_sndrexmitbad                             U_int32_t
	Tcps_sndrexmitbyte                            U_int32_t
	Tcps_sndrexmitpack                            U_int32_t
	Tcps_sndtotal                                 U_int32_t
	Tcps_sndurg                                   U_int32_t
	Tcps_sndwinup                                 U_int32_t
	Tcps_synchallenge                             U_int32_t
	Tcps_tailloss_rto                             U_int32_t
	Tcps_tfo_blackhole                            U_int32_t
	Tcps_tfo_cookie_invalid                       U_int32_t
	Tcps_tfo_cookie_rcv                           U_int32_t
	Tcps_tfo_cookie_req                           U_int32_t
	Tcps_tfo_cookie_req_rcv                       U_int32_t
	Tcps_tfo_cookie_sent                          U_int32_t
	Tcps_tfo_cookie_wrong                         U_int32_t
	Tcps_tfo_heuristics_disable                   U_int32_t
	Tcps_tfo_no_cookie_rcv                        U_int32_t
	Tcps_tfo_sndblackhole                         U_int32_t
	Tcps_tfo_syn_data_acked                       U_int32_t
	Tcps_tfo_syn_data_rcv                         U_int32_t
	Tcps_tfo_syn_data_sent                        U_int32_t
	Tcps_tfo_syn_loss                             U_int32_t
	Tcps_timeoutdrop                              U_int32_t
	Tcps_timer_drift_gt_1000_ms                   U_int32_t
	Tcps_timer_drift_le_1000_ms                   U_int32_t
	Tcps_timer_drift_le_100_ms                    U_int32_t
	Tcps_timer_drift_le_10_ms                     U_int32_t
	Tcps_timer_drift_le_1_ms                      U_int32_t
	Tcps_timer_drift_le_200_ms                    U_int32_t
	Tcps_timer_drift_le_20_ms                     U_int32_t
	Tcps_timer_drift_le_500_ms                    U_int32_t
	Tcps_timer_drift_le_50_ms                     U_int32_t
	Tcps_tlp_recoverlastpkt                       U_int32_t
	Tcps_tlp_recovery                             U_int32_t
	Tcps_unnecessary_rxmt                         U_int32_t
	Tcps_unused_1                                 U_int32_t
	Tcps_unused_2                                 U_int32_t
	Tcps_unused_3                                 U_int32_t
	Tcps_usedrtt                                  U_int32_t
	Tcps_usedrttvar                               U_int32_t
	Tcps_usedssthresh                             U_int32_t
	Tcps_fin_timeout_drops                        U_int32_t
	Tcps_ecn_ace_recv_ce                          U_int32_t
	Tcps_ecn_ace_syn_ce                           U_int32_t
	Tcps_ecn_ace_syn_ect0                         U_int32_t
	Tcps_ecn_ace_syn_ect1                         U_int32_t
	Tcps_ecn_ace_syn_not_ect                      U_int32_t
	Tcps_rack_recovery_episode                    U_int32_t
	Tcps_rack_reordering_timeout_recovery_episode U_int32_t
	Tcps_rack_rexmits                             U_int32_t
}

// Tcpstat_local
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/tcpstat_local
type Tcpstat_local struct {
	Badformat            unsafe.Pointer
	Badformatipsec       unsafe.Pointer
	Cleanup              unsafe.Pointer
	Deprecate6           unsafe.Pointer
	Dospacket            unsafe.Pointer
	Icmp6unreach         unsafe.Pointer
	Linkheur_comprxmt    unsafe.Pointer
	Linkheur_noackpri    unsafe.Pointer
	Linkheur_rxmtfloor   unsafe.Pointer
	Linkheur_stealthdrop unsafe.Pointer
	Linkheur_synrxmt     unsafe.Pointer
	Listbadsyn           unsafe.Pointer
	Noconnlist           unsafe.Pointer
	Noconnnolist         unsafe.Pointer
	Ooopacket            unsafe.Pointer
	Rstinsynrcv          unsafe.Pointer
	Synfin               unsafe.Pointer
	Synwindow            unsafe.Pointer
	Unspecv6             unsafe.Pointer
}

// Telemetry_notification_subsystem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/telemetry_notification_subsystem
type Telemetry_notification_subsystem struct {
	End      int32
	Start    int32
	Maxsize  unsafe.Pointer
	Routine  unsafe.Pointer
	Server   unsafe.Pointer
	Reserved Vm_address_t
}

// Termios
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/termios
type Termios struct {
	C_ispeed Speed_t
	C_iflag  Tcflag_t
	C_oflag  Tcflag_t
	C_lflag  Tcflag_t
	C_cc     Cc_t
	C_cflag  Tcflag_t
}

// Termios32
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/termios32
type Termios32 struct {
	C_cflag  uint32
	C_lflag  uint32
	C_ispeed uint32
	C_ospeed uint32
	C_oflag  uint32
	C_iflag  uint32
	C_cc     Cc_t
}

// Thread_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/thread_command
type Thread_command struct {
	Cmdsize unsafe.Pointer
	Cmd     unsafe.Pointer // Common to all load command structures. For this structure, set to `LC_THREAD` or `LC_UNIXTHREAD`.

}

// Thread_crash_exclaves_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/thread_crash_exclaves_info
type Thread_crash_exclaves_info struct {
	Tcei_flags     unsafe.Pointer
	Tcei_scid      unsafe.Pointer
	Tcei_thread_id unsafe.Pointer
}

// Thread_delta_snapshot_v2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/thread_delta_snapshot_v2
type Thread_delta_snapshot_v2 struct {
	Tds_eqos                    uint8
	Tds_sched_flags             unsafe.Pointer
	Tds_state                   unsafe.Pointer
	Tds_rqos_override           uint8
	Tds_rqos                    uint8
	Tds_voucher_identifier      unsafe.Pointer
	Tds_ss_flags                unsafe.Pointer
	Tds_io_tier                 uint8
	Tds_base_priority           int16
	Tds_sched_priority          int16
	Tds_thread_id               unsafe.Pointer
	Tds_last_made_runnable_time unsafe.Pointer
}

// Thread_delta_snapshot_v3
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/thread_delta_snapshot_v3
type Thread_delta_snapshot_v3 struct {
	Tds_eqos                    uint8
	Tds_rqos                    uint8
	Tds_last_made_runnable_time unsafe.Pointer
	Tds_base_priority           int16
	Tds_effective_policy        unsafe.Pointer
	Tds_state                   unsafe.Pointer
	Tds_thread_id               unsafe.Pointer
	Tds_sched_flags             unsafe.Pointer
	Tds_voucher_identifier      unsafe.Pointer
	Tds_ss_flags                unsafe.Pointer
	Tds_sched_priority          int16
	Tds_rqos_override           uint8
	Tds_requested_policy        unsafe.Pointer
	Tds_io_tier                 uint8
}

// Thread_exclaves_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/thread_exclaves_info
type Thread_exclaves_info struct {
	Tei_flags         unsafe.Pointer
	Tei_scid          unsafe.Pointer
	Tei_thread_offset unsafe.Pointer
}

// Thread_group_snapshot
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/thread_group_snapshot
type Thread_group_snapshot struct {
	Tgs_id   unsafe.Pointer
	Tgs_name unsafe.Pointer
}

// Thread_group_snapshot_v2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/thread_group_snapshot_v2
type Thread_group_snapshot_v2 struct {
	Tgs_flags unsafe.Pointer
	Tgs_id    unsafe.Pointer
	Tgs_name  unsafe.Pointer
}

// Thread_group_snapshot_v3
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/thread_group_snapshot_v3
type Thread_group_snapshot_v3 struct {
	Tgs_flags     unsafe.Pointer
	Tgs_id        unsafe.Pointer
	Tgs_name      unsafe.Pointer
	Tgs_name_cont unsafe.Pointer
}

// Thread_snapshot
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/thread_snapshot
type Thread_snapshot struct {
	Wait_event         unsafe.Pointer
	Ss_flags           unsafe.Pointer
	Continuation       unsafe.Pointer
	Paging_count       unsafe.Pointer
	Io_priority_size   unsafe.Pointer
	Metadata_count     unsafe.Pointer
	Ts_qos             unsafe.Pointer
	Disk_writes_size   unsafe.Pointer
	Total_syscalls     unsafe.Pointer
	Disk_writes_count  unsafe.Pointer
	Ts_rqos            unsafe.Pointer
	Disk_reads_count   unsafe.Pointer
	Sched_flags        unsafe.Pointer
	Non_paging_size    unsafe.Pointer
	User_time          unsafe.Pointer
	Nkern_frames       unsafe.Pointer
	Sched_pri          unsafe.Pointer
	Ts_rqos_override   unsafe.Pointer
	System_time        unsafe.Pointer
	Data_count         unsafe.Pointer
	Priority           unsafe.Pointer
	Paging_size        unsafe.Pointer
	Pth_name           unsafe.Pointer
	Nuser_frames       unsafe.Pointer
	Io_priority_count  unsafe.Pointer
	Metadata_size      unsafe.Pointer
	State              unsafe.Pointer
	Snapshot_magic     unsafe.Pointer
	Disk_reads_size    unsafe.Pointer
	Io_tier            unsafe.Pointer
	Data_size          unsafe.Pointer
	Thread_id          unsafe.Pointer
	Voucher_identifier unsafe.Pointer
	Non_paging_count   unsafe.Pointer
}

// Thread_snapshot_v2
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/thread_snapshot_v2
type Thread_snapshot_v2 struct {
	Ths_wait_event              unsafe.Pointer
	Ths_io_tier                 uint8
	Ths_sys_time                unsafe.Pointer
	Ths_base_priority           int16
	Ths_rqos                    uint8
	Ths_ss_flags                unsafe.Pointer
	Ths_last_made_runnable_time unsafe.Pointer
	Ths_continuation            unsafe.Pointer
	Ths_user_time               unsafe.Pointer
	Ths_voucher_identifier      unsafe.Pointer
	Ths_sched_priority          int16
	Ths_total_syscalls          unsafe.Pointer
	Ths_dqserialnum             unsafe.Pointer
	Ths_last_run_time           unsafe.Pointer
	Ths_rqos_override           uint8
	Ths_eqos                    uint8
	Ths_thread_id               unsafe.Pointer
	Ths_sched_flags             unsafe.Pointer
	Ths_state                   unsafe.Pointer
}

// Thread_snapshot_v3
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/thread_snapshot_v3
type Thread_snapshot_v3 struct {
	Ths_thread_t                unsafe.Pointer
	Ths_thread_id               unsafe.Pointer
	Ths_last_made_runnable_time unsafe.Pointer
	Ths_dqserialnum             unsafe.Pointer
	Ths_eqos                    uint8
	Ths_last_run_time           unsafe.Pointer
	Ths_voucher_identifier      unsafe.Pointer
	Ths_continuation            unsafe.Pointer
	Ths_io_tier                 uint8
	Ths_base_priority           int16
	Ths_wait_event              unsafe.Pointer
	Ths_sched_flags             unsafe.Pointer
	Ths_total_syscalls          unsafe.Pointer
	Ths_ss_flags                unsafe.Pointer
	Ths_sys_time                unsafe.Pointer
	Ths_rqos_override           uint8
	Ths_sched_priority          int16
	Ths_rqos                    uint8
	Ths_state                   unsafe.Pointer
	Ths_user_time               unsafe.Pointer
}

// Thread_snapshot_v4
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/thread_snapshot_v4
type Thread_snapshot_v4 struct {
	Ths_thread_t                unsafe.Pointer
	Ths_effective_policy        unsafe.Pointer
	Ths_last_made_runnable_time unsafe.Pointer
	Ths_eqos                    uint8
	Ths_base_priority           int16
	Ths_thread_id               unsafe.Pointer
	Ths_last_run_time           unsafe.Pointer
	Ths_state                   unsafe.Pointer
	Ths_voucher_identifier      unsafe.Pointer
	Ths_requested_policy        unsafe.Pointer
	Ths_wait_event              unsafe.Pointer
	Ths_user_time               unsafe.Pointer
	Ths_dqserialnum             unsafe.Pointer
	Ths_ss_flags                unsafe.Pointer
	Ths_rqos                    uint8
	Ths_rqos_override           uint8
	Ths_sched_priority          int16
	Ths_continuation            unsafe.Pointer
	Ths_sys_time                unsafe.Pointer
	Ths_io_tier                 uint8
	Ths_total_syscalls          unsafe.Pointer
	Ths_sched_flags             unsafe.Pointer
}

// Thsc_cpi
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/thsc_cpi
type Thsc_cpi struct {
	Tcpi_cycles       unsafe.Pointer
	Tcpi_instructions unsafe.Pointer
}

// Thsc_time_cpi
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/thsc_time_cpi
type Thsc_time_cpi struct {
	Ttci_cycles       unsafe.Pointer
	Ttci_instructions unsafe.Pointer
}

// Thsc_time_energy_cpi
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/thsc_time_energy_cpi
type Thsc_time_energy_cpi struct {
	Ttec_cycles           unsafe.Pointer
	Ttec_energy_nj        unsafe.Pointer
	Ttec_instructions     unsafe.Pointer
	Ttec_system_time_mach unsafe.Pointer
	Ttec_user_time_mach   unsafe.Pointer
}

// Timebase_freq_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/timebase_freq_t
type Timebase_freq_t struct {
	Timebase_num unsafe.Pointer
	Timebase_den unsafe.Pointer
}

// Timespec
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/timespec
type Timespec struct {
	Tv_sec  unsafe.Pointer
	Tv_nsec unsafe.Pointer
}

// Timeval
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/timeval
type Timeval struct {
	Tv_sec  unsafe.Pointer
	Tv_usec unsafe.Pointer
}

// Timeval32
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/timeval32
type Timeval32 struct {
	Tv_sec  int32
	Tv_usec int32
}

// Timeval64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/timeval64
type Timeval64 struct {
	Tv_sec  int64
	Tv_usec int64
}

// Timex
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/timex
type Timex struct {
	Esterror unsafe.Pointer
	Constant unsafe.Pointer
}

// Timezone
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/timezone
type Timezone struct {
	Tz_minuteswest unsafe.Pointer
	Tz_dsttime     unsafe.Pointer
}

// Tlv_descriptor
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/tlv_descriptor
type Tlv_descriptor struct {
	Offset unsafe.Pointer
	Thunk  *Tlv_descriptor
	Key    unsafe.Pointer
}

// Transitioning_task_snapshot
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/transitioning_task_snapshot
type Transitioning_task_snapshot struct {
	Tts_p_comm          unsafe.Pointer
	Tts_pid             unsafe.Pointer
	Tts_ss_flags        unsafe.Pointer
	Tts_transition_type unsafe.Pointer
	Tts_unique_pid      unsafe.Pointer
}

// Trust_cache_entry1
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/trust_cache_entry1
type Trust_cache_entry1 struct {
	Cdhash uint8
	Flags  uint8
}

// Trust_cache_module1
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/trust_cache_module1
type Trust_cache_module1 struct {
	Entries     Trust_cache_entry1
	Num_entries unsafe.Pointer
	Uuid        [16]byte
	Version     unsafe.Pointer
}

// Tsegqe_head
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/tsegqe_head
type Tsegqe_head struct {
	Lh_first U_int32_t
}

// Ttysize
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ttysize
type Ttysize struct {
	Ts_yyy  unsafe.Pointer
	Ts_cols unsafe.Pointer
}

// Twolevel_hint
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/twolevel_hint
type Twolevel_hint struct {
	Isub_image unsafe.Pointer // The subimage in which the symbol is defined. It is an index into the list of images that make up the umbrella image. If this field is 0, the symbol is in the umbrella image itself. If the image is not an umbrella framework or library, this field is 0.
	Itoc       unsafe.Pointer // The symbol index into the table of contents of the image specified by the `isub_image` field.

}

// Twolevel_hints_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/twolevel_hints_command
type Twolevel_hints_command struct {
	Cmd     unsafe.Pointer // Common to all load command structures. Set to `LC_TWOLEVEL_HINTS` for this structure.
	Offset  unsafe.Pointer
	Cmdsize unsafe.Pointer
	Nhints  unsafe.Pointer
}

// Udphdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/udphdr
type Udphdr struct {
	Uh_sum   U_short
	Uh_sport U_short
	Uh_ulen  U_short
	Uh_dport U_short
}

// Udpiphdr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/udpiphdr
type Udpiphdr struct {
	Ui_i Ipovly
	Ui_u Udphdr
}

// Udpstat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/udpstat
type Udpstat struct {
	Udpps_pcbcachemiss     U_int32_t
	Udpps_pcbhashmiss      U_int32_t
	Udps_badlen            U_int32_t
	Udps_badsum            U_int32_t
	Udps_fastout           U_int32_t
	Udps_filtermcast       U_int32_t
	Udps_fullsock          U_int32_t
	Udps_hdrops            U_int32_t
	Udps_ipackets          U_int32_t
	Udps_noport            U_int32_t
	Udps_noportbcast       U_int32_t
	Udps_noportmcast       U_int32_t
	Udps_nosum             U_int32_t
	Udps_opackets          U_int32_t
	Udps_rcv6_swcsum       U_int32_t
	Udps_rcv6_swcsum_bytes U_int32_t
	Udps_rcv_swcsum        U_int32_t
	Udps_rcv_swcsum_bytes  U_int32_t
	Udps_snd6_swcsum       U_int32_t
	Udps_snd6_swcsum_bytes U_int32_t
	Udps_snd_swcsum        U_int32_t
	Udps_snd_swcsum_bytes  U_int32_t
}

// User32_dyld_uuid_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/user32_dyld_uuid_info
type User32_dyld_uuid_info struct {
	ImageUUID        [16]byte
	ImageLoadAddress unsafe.Pointer
}

// User32_fssearchblock
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/user32_fssearchblock
type User32_fssearchblock struct {
	Returnbuffer        User32_addr_t
	Searchparams2       User32_addr_t
	Returnbuffersize    User32_size_t
	Timelimit           User32_timeval
	Searchparams1       User32_addr_t
	Maxmatches          User32_ulong_t
	Sizeofsearchparams2 User32_size_t
	Sizeofsearchparams1 User32_size_t
	Returnattrs         User32_addr_t
	Searchattrs         Attrlist
}

// User32_itimerval
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/user32_itimerval
type User32_itimerval struct {
	It_value    User32_timeval
	It_interval User32_timeval
}

// User32_msqid_ds
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/user32_msqid_ds
type User32_msqid_ds struct {
	Msg_last   int32
	Msg_rtime  User32_time_t
	Msg_perm   Ipc_perm
	Msg_pad2   int32
	Msg_lspid  int32
	Msg_lrpid  int32
	Msg_qnum   User32_msgqnum_t
	Msg_pad4   int32
	Msg_qbytes User32_msglen_t
	Msg_first  int32
	Msg_cbytes User32_msglen_t
	Msg_pad1   int32
	Msg_ctime  User32_time_t
	Msg_stime  User32_time_t
	Msg_pad3   int32
}

// User32_rusage
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/user32_rusage
type User32_rusage struct {
	Ru_isrss    User32_long_t
	Ru_ixrss    User32_long_t
	Ru_oublock  User32_long_t
	Ru_nivcsw   User32_long_t
	Ru_majflt   User32_long_t
	Ru_utime    User32_timeval
	Ru_maxrss   User32_long_t
	Ru_stime    User32_timeval
	Ru_inblock  User32_long_t
	Ru_idrss    User32_long_t
	Ru_nvcsw    User32_long_t
	Ru_nswap    User32_long_t
	Ru_msgrcv   User32_long_t
	Ru_minflt   User32_long_t
	Ru_msgsnd   User32_long_t
	Ru_nsignals User32_long_t
}

// User32_sf_hdtr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/user32_sf_hdtr
type User32_sf_hdtr struct {
	Trailers User32_addr_t
	Headers  User32_addr_t
	Trl_cnt  unsafe.Pointer
	Hdr_cnt  unsafe.Pointer
}

// User32_timespec
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/user32_timespec
type User32_timespec struct {
	Tv_nsec User32_long_t
	Tv_sec  User32_time_t
}

// User32_timeval
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/user32_timeval
type User32_timeval struct {
	Tv_usec int32
	Tv_sec  User32_time_t
}

// User32_timex
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/user32_timex
type User32_timex struct {
	Shift     int32
	Tolerance User32_long_t
	Constant  User32_long_t
	Status    int32
	Jitcnt    User32_long_t
	Offset    User32_long_t
	Freq      User32_long_t
	Jitter    User32_long_t
	Maxerror  User32_long_t
	Stabil    User32_long_t
	Stbcnt    User32_long_t
	Errcnt    User32_long_t
	Ppsfreq   User32_long_t
	Esterror  User32_long_t
	Modes     U_int32_t
	Calcnt    User32_long_t
	Precision User32_long_t
}

// User32_vfsidctl
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/user32_vfsidctl
type User32_vfsidctl struct {
	Vc_ptr   User32_addr_t
	Vc_spare U_int32_t
	Vc_vers  unsafe.Pointer
	Vc_len   User32_size_t
	Vc_fsid  Fsid_t
}

// User64_dyld_aot_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/user64_dyld_aot_info
type User64_dyld_aot_info struct {
	AotImageKey  uint8
	AotImageSize unsafe.Pointer
}

// User64_dyld_uuid_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/user64_dyld_uuid_info
type User64_dyld_uuid_info struct {
	ImageUUID        [16]byte
	ImageLoadAddress unsafe.Pointer
}

// User64_fssearchblock
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/user64_fssearchblock
type User64_fssearchblock struct {
	Returnbuffersize    User64_size_t
	Searchparams1       User64_addr_t
	Sizeofsearchparams2 User64_size_t
	Returnbuffer        User64_addr_t
	Searchparams2       User64_addr_t
	Searchattrs         Attrlist
	Sizeofsearchparams1 User64_size_t
	Returnattrs         User64_addr_t
	Timelimit           User64_timeval
	Maxmatches          User64_ulong_t
}

// User64_itimerval
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/user64_itimerval
type User64_itimerval struct {
	It_interval User64_timeval
}

// User64_msqid_ds
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/user64_msqid_ds
type User64_msqid_ds struct {
	Msg_qbytes User64_msglen_t
	Msg_last   int32
}

// User64_rusage
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/user64_rusage
type User64_rusage struct {
	Ru_nvcsw    User64_long_t
	Ru_majflt   User64_long_t
	Ru_oublock  User64_long_t
	Ru_nswap    User64_long_t
	Ru_nsignals User64_long_t
	Ru_nivcsw   User64_long_t
	Ru_inblock  User64_long_t
	Ru_maxrss   User64_long_t
	Ru_stime    User64_timeval
	Ru_isrss    User64_long_t
	Ru_msgrcv   User64_long_t
	Ru_ixrss    User64_long_t
	Ru_msgsnd   User64_long_t
	Ru_utime    User64_timeval
	Ru_idrss    User64_long_t
	Ru_minflt   User64_long_t
}

// User64_sf_hdtr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/user64_sf_hdtr
type User64_sf_hdtr struct {
	Trailers User64_addr_t
	Hdr_cnt  unsafe.Pointer
	Trl_cnt  unsafe.Pointer
	Headers  User64_addr_t
}

// User64_timespec
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/user64_timespec
type User64_timespec struct {
	Tv_sec  User64_time_t
	Tv_nsec User64_long_t
}

// User64_timeval
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/user64_timeval
type User64_timeval struct {
	Tv_usec int32
	Tv_sec  User64_time_t
}

// User64_timex
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/user64_timex
type User64_timex struct {
	Shift     int32
	Ppsfreq   User64_long_t
	Offset    User64_long_t
	Maxerror  User64_long_t
	Freq      User64_long_t
	Precision User64_long_t
	Status    int32
	Jitter    User64_long_t
	Esterror  User64_long_t
	Jitcnt    User64_long_t
	Constant  User64_long_t
	Stbcnt    User64_long_t
	Modes     U_int32_t
	Stabil    User64_long_t
	Calcnt    User64_long_t
	Errcnt    User64_long_t
	Tolerance User64_long_t
}

// User_msqid_ds
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/user_msqid_ds
type User_msqid_ds struct {
	Msg_cbytes User_msglen_t
	Msg_stime  User_time_t
	Msg_perm   Ipc_perm
	Msg_qnum   User_msgqnum_t
	Msg_last   *Msg
	Msg_pad1   int32
	Msg_first  *Msg
	Msg_rtime  User_time_t
	Msg_pad4   int32
	Msg_pad3   int32
	Msg_lrpid  int32
	Msg_ctime  User_time_t
	Msg_lspid  int32
	Msg_qbytes User_msglen_t
	Msg_pad2   int32
}

// User_nfs_args
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/user_nfs_args
type User_nfs_args struct {
	Auth       unsafe.Pointer
	Deadthresh unsafe.Pointer
	Sotype     unsafe.Pointer
}

// User_nfs_export_args
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/user_nfs_export_args
type User_nfs_export_args struct {
	Nxa_expid    unsafe.Pointer
	Nxa_exppath  User_addr_t
	Nxa_flags    unsafe.Pointer
	Nxa_fsid     unsafe.Pointer
	Nxa_fspath   User_addr_t
	Nxa_netcount unsafe.Pointer
	Nxa_nets     User_addr_t
}

// User_nfsd_args
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/user_nfsd_args
type User_nfsd_args struct {
	Name    User_addr_t
	Namelen unsafe.Pointer
	Sock    unsafe.Pointer
}

// User_sf_hdtr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/user_sf_hdtr
type User_sf_hdtr struct {
	Headers  User_addr_t
	Hdr_cnt  unsafe.Pointer
	Trl_cnt  unsafe.Pointer
	Trailers User_addr_t
}

// User_termios
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/user_termios
type User_termios struct {
	C_iflag  User_tcflag_t
	C_ospeed User_speed_t
	C_oflag  User_tcflag_t
	C_ispeed User_speed_t
	C_cflag  User_tcflag_t
	C_cc     Cc_t
	C_lflag  User_tcflag_t
}

// User_timespec
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/user_timespec
type User_timespec struct {
	Tv_nsec User_long_t
	Tv_sec  User_time_t
}

// User_timeval
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/user_timeval
type User_timeval struct {
	Tv_usec int32
	Tv_sec  User_time_t
}

// User_vfsidctl
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/user_vfsidctl
type User_vfsidctl struct {
	Vc_fsid  Fsid_t
	Vc_spare U_int32_t
	Vc_vers  unsafe.Pointer
	Vc_len   User_size_t
	Vc_ptr   User_addr_t
}

// Utun_stats_param
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/utun_stats_param
type Utun_stats_param struct {
	Utsp_packets U_int64_t
	Utsp_bytes   U_int64_t
	Utsp_errors  U_int64_t
}

// Uuid_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/uuid_command
type Uuid_command struct {
	Uuid    uint8
	Cmdsize unsafe.Pointer
	Cmd     unsafe.Pointer
}

// Vend
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vend
type Vend struct {
	V_flags U_int32_t
	V_magic U_char
}

// Version_min_command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/version_min_command
type Version_min_command struct {
	Version unsafe.Pointer
}

// Vfs_attr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vfs_attr
type Vfs_attr struct {
	F_files       unsafe.Pointer
	F_bsize       unsafe.Pointer
	F_create_time Timespec
	F_fssubtype   unsafe.Pointer
	F_reserved    unsafe.Pointer
}

// Vfs_fsentry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vfs_fsentry
type Vfs_fsentry struct {
	Vfe_flags     unsafe.Pointer
	Vfe_fsname    unsafe.Pointer
	Vfe_opvdescs  *Vnodeopv_desc
	Vfe_fstypenum unsafe.Pointer
	Vfe_vopcnt    unsafe.Pointer
	Vfe_reserv    unsafe.Pointer
	Vfe_vfsops    *Vfsops
}

// Vfs_server
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vfs_server
type Vfs_server struct {
	Vs_server_name U_int8_t
	Vs_minutes     unsafe.Pointer
}

// Vfsconf
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vfsconf
type Vfsconf struct {
	Vfc_name  unsafe.Pointer
	Vfc_flags unsafe.Pointer
}

// Vfsidctl
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vfsidctl
type Vfsidctl struct {
	Vc_spare U_int32_t
	Vc_len   uintptr
	Vc_fsid  Fsid_t
	Vc_vers  unsafe.Pointer
	Vc_ptr   unsafe.Pointer
}

// Vfsioattr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vfsioattr
type Vfsioattr struct {
	Io_maxwritecnt     U_int32_t
	Io_flags           U_int32_t
	Io_maxsegreadsize  U_int32_t
	Io_segwritecnt     U_int32_t
	Io_segreadcnt      U_int32_t
	Io_maxsegwritesize U_int32_t
	Io_devblocksize    U_int32_t
	Io_maxreadcnt      U_int32_t
}

// Vfsops
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vfsops
type Vfsops struct {
	Vfs_setattr      Vfs_context_t
	Vfs_root         Vfs_context_t
	Vfs_quotactl     Vfs_context_t
	Vfs_start        Vfs_context_t
	Vfs_unmount      Vfs_context_t
	Vfs_vptofh       Vfs_context_t
	Vfs_sync         Vfs_context_t
	Vfs_vget         Vfs_context_t
	Vfs_init         *Vfsconf
	Vfs_getattr      Vfs_context_t
	Vfs_mount        Vfs_context_t
	Vfs_sysctl       Vfs_context_t
	Vfs_fhtovp       Vfs_context_t
	Vfs_reserved4    unsafe.Pointer
	Vfs_reserved2    unsafe.Pointer
	Vfs_reserved1    unsafe.Pointer
	Vfs_vget_snapdir Vfs_context_t
	Vfs_reserved3    unsafe.Pointer
	Vfs_reserved5    unsafe.Pointer
	Vfs_ioctl        Vfs_context_t
}

// Vfsquery
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vfsquery
type Vfsquery struct {
	Vq_flags U_int32_t
	Vq_spare U_int32_t
}

// Vfsstatfs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vfsstatfs
type Vfsstatfs struct {
	F_ffree       unsafe.Pointer
	F_mntfromname unsafe.Pointer
	F_bused       unsafe.Pointer
	F_files       unsafe.Pointer
	F_blocks      unsafe.Pointer
}

// Vlanreq
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vlanreq
type Vlanreq struct {
	Vlr_parent unsafe.Pointer
	Vlr_tag    U_short
}

// Vmspace
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vmspace
type Vmspace struct {
	Dummy4 Caddr_t
	Dummy  unsafe.Pointer
	Dummy2 Caddr_t
	Dummy3 unsafe.Pointer
}

// Vnode_attr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnode_attr
type Vnode_attr struct {
	Va_fileid             unsafe.Pointer
	Va_iosize             unsafe.Pointer
	Va_finderinfo         uint8
	Va_objtag             unsafe.Pointer
	Va_dirlinkcount       unsafe.Pointer
	Va_reserved1          unsafe.Pointer
	Va_data_alloc         unsafe.Pointer
	Va_vaflags            unsafe.Pointer
	Va_nchildren          unsafe.Pointer
	Va_dataprotect_flags  unsafe.Pointer
	Va_type               Vtype
	Va_data_size          unsafe.Pointer
	Va_access_time        Timespec
	Va_name               unsafe.Pointer
	Va_total_size         unsafe.Pointer
	Va_acl                unsafe.Pointer
	Va_rdev               int32
	Va_change_time        Timespec
	Va_modify_time        Timespec
	Va_dataprotect_class  unsafe.Pointer
	Va_parentid           unsafe.Pointer
	Va_gen                unsafe.Pointer
	Va_flags              unsafe.Pointer
	Va_linkid             unsafe.Pointer
	Va_document_id        unsafe.Pointer
	Va_devid              unsafe.Pointer
	Va_fsid64             Fsid_t
	Va_backup_time        Timespec
	Va_write_gencount     unsafe.Pointer
	Va_total_alloc        unsafe.Pointer
	Va_supported          unsafe.Pointer
	Va_mode               uint16
	Va_encoding           unsafe.Pointer
	Va_rsrc_alloc         unsafe.Pointer
	Va_uuuid              unsafe.Pointer
	Va_fsid               unsafe.Pointer
	Va_objtype            unsafe.Pointer
	Va_uid                uint32
	Va_rsrc_length        unsafe.Pointer
	Va_user_access        unsafe.Pointer
	Va_gid                uint32
	Va_nlink              unsafe.Pointer
	Va_create_time        Timespec
	Va_active             unsafe.Pointer
	Va_filerev            unsafe.Pointer
	Va_addedtime          Timespec
	Va_guuid              unsafe.Pointer
	Va_private_size       unsafe.Pointer
	Va_clone_id           unsafe.Pointer
	Va_extflags           unsafe.Pointer
	Va_recursive_gencount unsafe.Pointer
	Va_attribution_tag    unsafe.Pointer
	Va_clone_refcnt       unsafe.Pointer
}

// Vnode_fsparam
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnode_fsparam
type Vnode_fsparam struct {
	Vnfs_mp         unsafe.Pointer
	Vnfs_vtype      Vtype
	Vnfs_fsnode     unsafe.Pointer
	Vnfs_vops       unsafe.Pointer
	Vnfs_cnp        *Componentname
	Vnfs_rdev       int32
	Vnfs_flags      unsafe.Pointer
	Vnfs_marksystem unsafe.Pointer
	Vnfs_str        unsafe.Pointer
	Vnfs_filesize   int64
	Vnfs_dvp        unsafe.Pointer
	Vnfs_markroot   unsafe.Pointer
}

// Vnodeopv_desc
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnodeopv_desc
type Vnodeopv_desc struct {
	Opv_desc_vector_p unsafe.Pointer
	Opv_desc_ops      *Vnodeopv_entry_desc
}

// Vnodeopv_entry_desc
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnodeopv_entry_desc
type Vnodeopv_entry_desc struct {
	Opve_impl unsafe.Pointer
	Opve_op   unsafe.Pointer
}

// Vnop_access_args - Call down to a filesystem to close a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_access_args
type Vnop_access_args struct {
	A_vp      Vnode_t
	A_desc    unsafe.Pointer
	A_action  unsafe.Pointer
	A_context Vfs_context_t
}

// Vnop_advlock_args - Query a filesystem for path properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_advlock_args
type Vnop_advlock_args struct {
	A_id      Caddr_t
	A_vp      Vnode_t
	A_desc    unsafe.Pointer
	A_context Vfs_context_t
	A_op      unsafe.Pointer
	A_timeout *Timespec
	A_flags   unsafe.Pointer
	A_fl      *Flock
}

// Vnop_allocate_args - Aquire or release and advisory lock on a vnode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_allocate_args
type Vnop_allocate_args struct {
	A_flags          U_int32_t
	A_length         int64
	A_vp             Vnode_t
	A_context        Vfs_context_t
	A_bytesallocated *int64
	A_offset         int64
	A_desc           unsafe.Pointer
}

// Vnop_blktooff_args - List extended attribute keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_blktooff_args
type Vnop_blktooff_args struct {
	A_desc   unsafe.Pointer
	A_lblkno Daddr64_t
	A_vp     Vnode_t
	A_offset *int64
}

// Vnop_blockmap_args - Call down to a filesystem to convert a file offset to a logical block number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_blockmap_args
type Vnop_blockmap_args struct {
	A_run     *uintptr
	A_size    uintptr
	A_context Vfs_context_t
	A_vp      Vnode_t
	A_desc    unsafe.Pointer
	A_poff    unsafe.Pointer
	A_bpn     *Daddr64_t
	A_flags   unsafe.Pointer
	A_foffset int64
}

// Vnop_bwrite_args
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_bwrite_args
type Vnop_bwrite_args struct {
	A_bp   Buf_t
	A_desc unsafe.Pointer
}

// Vnop_clonefile_args
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_clonefile_args
type Vnop_clonefile_args struct {
	A_context              Vfs_context_t
	A_cnp                  *Componentname
	A_desc                 unsafe.Pointer
	A_vpp                  *Vnode_t
	A_vap                  *Vnode_attr
	A_fvp                  Vnode_t
	A_dvp                  Vnode_t
	A_flags                unsafe.Pointer
	A_reserved             unsafe.Pointer
	A_dir_clone_authorizer Vfs_context_t
}

// Vnop_close_args - Call down to a filesystem to open a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_close_args
type Vnop_close_args struct {
	A_context Vfs_context_t
	A_desc    unsafe.Pointer
	A_vp      Vnode_t
	A_fflag   unsafe.Pointer
}

// Vnop_copyfile_args - Write data from a mapped file back to disk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_copyfile_args
type Vnop_copyfile_args struct {
	A_fvp     Vnode_t
	A_context Vfs_context_t
	A_tcnp    *Componentname
	A_mode    unsafe.Pointer
	A_tvp     Vnode_t
	A_desc    unsafe.Pointer
	A_flags   unsafe.Pointer
	A_tdvp    Vnode_t
}

// Vnop_create_args - Call down to a filesystem to look for a directory entry by name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_create_args
type Vnop_create_args struct {
	A_vpp     *Vnode_t
	A_context Vfs_context_t
	A_cnp     *Componentname
	A_dvp     Vnode_t
	A_vap     *Vnode_attr
	A_desc    unsafe.Pointer
}

// Vnop_exchange_args - Call down to a filesystem or device to check if a file is ready for I/O and request later notification if it is not currently ready.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_exchange_args
type Vnop_exchange_args struct {
	A_context Vfs_context_t
	A_fvp     Vnode_t
	A_options unsafe.Pointer
	A_tvp     Vnode_t
	A_desc    unsafe.Pointer
}

// Vnop_fsync_args - Inform a filesystem that a file is no longer mapped.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_fsync_args
type Vnop_fsync_args struct {
	A_waitfor unsafe.Pointer
	A_vp      Vnode_t
	A_context Vfs_context_t
	A_desc    unsafe.Pointer
}

// Vnop_generic_args
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_generic_args
type Vnop_generic_args struct {
	A_desc unsafe.Pointer
}

// Vnop_getattr_args - Call down to a filesystem to see if a kauth-style operation is permitted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_getattr_args
type Vnop_getattr_args struct {
	A_vp      Vnode_t
	A_vap     *Vnode_attr
	A_context Vfs_context_t
	A_desc    unsafe.Pointer
}

// Vnop_getattrlistbulk_args
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_getattrlistbulk_args
type Vnop_getattrlistbulk_args struct {
	A_context     Vfs_context_t
	A_options     unsafe.Pointer
	A_desc        unsafe.Pointer
	A_alist       *Attrlist
	A_eofflag     unsafe.Pointer
	A_private     unsafe.Pointer
	A_actualcount unsafe.Pointer
	A_uio         unsafe.Pointer
	A_vp          Vnode_t
	A_vap         *Vnode_attr
}

// Vnop_getxattr_args - Write data from a mapped file back to disk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_getxattr_args
type Vnop_getxattr_args struct {
	A_uio     Uio_t
	A_size    *uintptr
	A_options unsafe.Pointer
	A_name    unsafe.Pointer
	A_context Vfs_context_t
	A_desc    unsafe.Pointer
	A_vp      Vnode_t
}

// Vnop_inactive_args - Call down to a filesystem to get the pathname represented by a symbolic link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_inactive_args
type Vnop_inactive_args struct {
	A_desc    unsafe.Pointer
	A_context Vfs_context_t
	A_vp      Vnode_t
}

// Vnop_ioctl_args
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_ioctl_args
type Vnop_ioctl_args struct {
	A_desc    unsafe.Pointer
	A_context Vfs_context_t
	A_fflag   unsafe.Pointer
	A_data    Caddr_t
	A_vp      Vnode_t
	A_command U_long
}

// Vnop_kqfilt_add_args
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_kqfilt_add_args
type Vnop_kqfilt_add_args struct {
	A_context Vfs_context_t
	A_desc    unsafe.Pointer
	A_vp      unsafe.Pointer
	A_kn      unsafe.Pointer
}

// Vnop_kqfilt_remove_args
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_kqfilt_remove_args
type Vnop_kqfilt_remove_args struct {
	A_vp      unsafe.Pointer
	A_desc    unsafe.Pointer
	A_context Vfs_context_t
	A_ident   uintptr
}

// Vnop_link_args - Call down to a filesystem to delete a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_link_args
type Vnop_link_args struct {
	A_context Vfs_context_t
	A_desc    unsafe.Pointer
	A_tdvp    Vnode_t
	A_vp      Vnode_t
	A_cnp     *Componentname
}

// Vnop_listxattr_args - Remove extended file attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_listxattr_args
type Vnop_listxattr_args struct {
	A_context Vfs_context_t
	A_desc    unsafe.Pointer
	A_uio     Uio_t
	A_size    *uintptr
	A_options unsafe.Pointer
	A_vp      Vnode_t
}

// Vnop_lookup_args
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_lookup_args
type Vnop_lookup_args struct {
	A_vpp     *Vnode_t
	A_context Vfs_context_t
	A_desc    unsafe.Pointer
	A_dvp     Vnode_t
	A_cnp     *Componentname
}

// Vnop_mkdir_args - Call down to a filesystem to rename a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_mkdir_args
type Vnop_mkdir_args struct {
	A_vpp     *Vnode_t
	A_vap     *Vnode_attr
	A_context Vfs_context_t
	A_dvp     Vnode_t
	A_cnp     *Componentname
	A_desc    unsafe.Pointer
}

// Vnop_mknod_args - Call down to a filesystem to create a whiteout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_mknod_args
type Vnop_mknod_args struct {
	A_desc    unsafe.Pointer
	A_cnp     *Componentname
	A_vap     *Vnode_attr
	A_dvp     Vnode_t
	A_vpp     *Vnode_t
	A_context Vfs_context_t
}

// Vnop_mmap_args - Call down to a filesystem to invalidate all open file descriptors for a vnode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_mmap_args
type Vnop_mmap_args struct {
	A_vp      Vnode_t
	A_context Vfs_context_t
	A_fflags  unsafe.Pointer
	A_desc    unsafe.Pointer
}

// Vnop_mmap_check_args
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_mmap_check_args
type Vnop_mmap_check_args struct {
	A_context Vfs_context_t
	A_desc    unsafe.Pointer
	A_flags   unsafe.Pointer
	A_vp      Vnode_t
}

// Vnop_mnomap_args - Notify a filesystem that a file is being mmap-ed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_mnomap_args
type Vnop_mnomap_args struct {
	A_context Vfs_context_t
	A_desc    unsafe.Pointer
	A_vp      Vnode_t
}

// Vnop_monitor_args
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_monitor_args
type Vnop_monitor_args struct {
	A_context Vfs_context_t
	A_desc    unsafe.Pointer
	A_events  unsafe.Pointer
	A_flags   unsafe.Pointer
	A_handle  unsafe.Pointer
	A_vp      Vnode_t
}

// Vnop_offtoblk_args - Call down to a filesystem to convert a logical block number to a file offset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_offtoblk_args
type Vnop_offtoblk_args struct {
	A_desc   unsafe.Pointer
	A_offset int64
	A_vp     Vnode_t
	A_lblkno *Daddr64_t
}

// Vnop_open_args - Call down to a filesystem to create a special file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_open_args
type Vnop_open_args struct {
	A_mode    unsafe.Pointer
	A_vp      Vnode_t
	A_desc    unsafe.Pointer
	A_context Vfs_context_t
}

// Vnop_pagein_args - Pre-allocate space for a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_pagein_args
type Vnop_pagein_args struct {
	A_flags     unsafe.Pointer
	A_desc      unsafe.Pointer
	A_vp        Vnode_t
	A_context   Vfs_context_t
	A_pl        Upl_t
	A_f_offset  int64
	A_pl_offset Upl_offset_t
	A_size      uintptr
}

// Vnop_pageout_args - Pull file data into memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_pageout_args
type Vnop_pageout_args struct {
	A_pl_offset Upl_offset_t
	A_size      uintptr
	A_pl        Upl_t
	A_f_offset  int64
	A_flags     unsafe.Pointer
	A_vp        Vnode_t
	A_context   Vfs_context_t
	A_desc      unsafe.Pointer
}

// Vnop_pathconf_args - Release filesystem-internal resources for a vnode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_pathconf_args
type Vnop_pathconf_args struct {
	A_retval  unsafe.Pointer
	A_desc    unsafe.Pointer
	A_vp      Vnode_t
	A_name    unsafe.Pointer
	A_context Vfs_context_t
}

// Vnop_read_args - Call down to a filesystem to set vnode attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_read_args
type Vnop_read_args struct {
	A_ioflag  unsafe.Pointer
	A_context Vfs_context_t
	A_vp      Vnode_t
	A_desc    unsafe.Pointer
	A_uio     unsafe.Pointer
}

// Vnop_readdir_args - Call down to a filesystem to create a symbolic link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_readdir_args
type Vnop_readdir_args struct {
	A_flags     unsafe.Pointer
	A_vp        Vnode_t
	A_desc      unsafe.Pointer
	A_context   Vfs_context_t
	A_uio       unsafe.Pointer
	A_numdirent unsafe.Pointer
	A_eofflag   unsafe.Pointer
}

// Vnop_readdirattr_args - Call down to a filesystem to enumerate directory entries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_readdirattr_args
type Vnop_readdirattr_args struct {
	A_actualcount unsafe.Pointer
	A_desc        unsafe.Pointer
	A_alist       *Attrlist
	A_context     Vfs_context_t
	A_uio         unsafe.Pointer
	A_newstate    unsafe.Pointer
	A_vp          Vnode_t
	A_eofflag     unsafe.Pointer
	A_options     unsafe.Pointer
	A_maxcount    unsafe.Pointer
}

// Vnop_readlink_args - Call down to get file attributes for many files in a directory at once.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_readlink_args
type Vnop_readlink_args struct {
	A_desc    unsafe.Pointer
	A_uio     unsafe.Pointer
	A_context Vfs_context_t
	A_vp      Vnode_t
}

// Vnop_reclaim_args - Notify a filesystem that the last usecount (persistent reference) on a vnode has been dropped.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_reclaim_args
type Vnop_reclaim_args struct {
	A_desc    unsafe.Pointer
	A_vp      Vnode_t
	A_context Vfs_context_t
}

// Vnop_remove_args
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_remove_args
type Vnop_remove_args struct {
	A_flags   unsafe.Pointer
	A_context Vfs_context_t
	A_cnp     *Componentname
	A_dvp     Vnode_t
	A_desc    unsafe.Pointer
	A_vp      Vnode_t
}

// Vnop_removexattr_args
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_removexattr_args
type Vnop_removexattr_args struct {
	A_name    unsafe.Pointer
	A_options unsafe.Pointer
	A_context Vfs_context_t
	A_vp      Vnode_t
	A_desc    unsafe.Pointer
}

// Vnop_rename_args - Call down to a filesystem to create a hardlink to a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_rename_args
type Vnop_rename_args struct {
	A_context Vfs_context_t
	A_tdvp    Vnode_t
	A_fvp     Vnode_t
	A_tvp     Vnode_t
	A_desc    unsafe.Pointer
	A_fdvp    Vnode_t
	A_tcnp    *Componentname
	A_fcnp    *Componentname
}

// Vnop_renamex_args
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_renamex_args
type Vnop_renamex_args struct {
	A_context Vfs_context_t
	A_tcnp    *Componentname
	A_desc    unsafe.Pointer
	A_flags   Vfs_rename_flags_t
	A_tvp     Vnode_t
	A_fdvp    Vnode_t
	A_vap     *Vnode_attr
	A_tdvp    Vnode_t
	A_fvp     Vnode_t
	A_fcnp    *Componentname
}

// Vnop_revoke_args - Call down to a filesystem to atomically exchange the data of two files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_revoke_args
type Vnop_revoke_args struct {
	A_desc    unsafe.Pointer
	A_context Vfs_context_t
	A_flags   unsafe.Pointer
	A_vp      Vnode_t
}

// Vnop_rmdir_args - Call down to a filesystem to create a directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_rmdir_args
type Vnop_rmdir_args struct {
	A_context Vfs_context_t
	A_cnp     *Componentname
	A_dvp     Vnode_t
	A_vp      Vnode_t
	A_desc    unsafe.Pointer
}

// Vnop_searchfs_args - Write data from a mapped file back to disk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_searchfs_args
type Vnop_searchfs_args struct {
	A_context       Vfs_context_t
	A_scriptcode    unsafe.Pointer
	A_returnattrs   *Attrlist
	A_timelimit     *Timeval
	A_nummatches    unsafe.Pointer
	A_vp            Vnode_t
	A_searchstate   *Searchstate
	A_searchparams2 unsafe.Pointer
	A_searchattrs   *Attrlist
	A_options       unsafe.Pointer
	A_uio           unsafe.Pointer
	A_searchparams1 unsafe.Pointer
	A_desc          unsafe.Pointer
	A_maxmatches    unsafe.Pointer
}

// Vnop_select_args
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_select_args
type Vnop_select_args struct {
	A_fflags  unsafe.Pointer
	A_desc    unsafe.Pointer
	A_vp      Vnode_t
	A_wql     unsafe.Pointer
	A_which   unsafe.Pointer
	A_context Vfs_context_t
}

// Vnop_setattr_args - Call down to a filesystem to get vnode attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_setattr_args
type Vnop_setattr_args struct {
	A_context Vfs_context_t
	A_vap     *Vnode_attr
	A_desc    unsafe.Pointer
	A_vp      Vnode_t
}

// Vnop_setlabel_args
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_setlabel_args
type Vnop_setlabel_args struct {
	A_vl      unsafe.Pointer
	A_vp      unsafe.Pointer
	A_desc    unsafe.Pointer
	A_context Vfs_context_t
}

// Vnop_setxattr_args
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_setxattr_args
type Vnop_setxattr_args struct {
	A_uio     Uio_t
	A_options unsafe.Pointer
	A_desc    unsafe.Pointer
	A_context Vfs_context_t
	A_vp      Vnode_t
	A_name    unsafe.Pointer
}

// Vnop_strategy_args - Call down to a filesystem to get information about the on-disk layout of a file region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_strategy_args
type Vnop_strategy_args struct {
	A_bp   unsafe.Pointer
	A_desc unsafe.Pointer
}

// Vnop_symlink_args - Call down to a filesystem to delete a directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_symlink_args
type Vnop_symlink_args struct {
	A_context Vfs_context_t
	A_vpp     *Vnode_t
	A_target  unsafe.Pointer
	A_dvp     Vnode_t
	A_vap     *Vnode_attr
	A_cnp     *Componentname
	A_desc    unsafe.Pointer
}

// Vnop_verify_args
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_verify_args
type Vnop_verify_args struct {
	A_buf           *uint8
	A_bufsize       uintptr
	A_context       Vfs_context_t
	A_desc          unsafe.Pointer
	A_flags         Vnode_verify_flags_t
	A_foffset       int64
	A_verifyblksize *uintptr
	A_vp            Vnode_t
	A_verify_ctxp   unsafe.Pointer
}

// Vnop_whiteout_args - Call down to a filesystem to create a regular file (VREG).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_whiteout_args
type Vnop_whiteout_args struct {
	A_flags   unsafe.Pointer
	A_cnp     *Componentname
	A_desc    unsafe.Pointer
	A_dvp     Vnode_t
	A_context Vfs_context_t
}

// Vnop_write_args
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/vnop_write_args
type Vnop_write_args struct {
	A_context Vfs_context_t
	A_vp      Vnode_t
	A_uio     unsafe.Pointer
	A_desc    unsafe.Pointer
	A_ioflag  unsafe.Pointer
}

// Winsize
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/winsize
type Winsize struct {
	Ws_row    unsafe.Pointer
	Ws_col    unsafe.Pointer
	Ws_ypixel unsafe.Pointer
	Ws_xpixel unsafe.Pointer
}

// X86_cpmu_state64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/x86_cpmu_state64
type X86_cpmu_state64 struct {
	Ctrs uint64
}

// Xdrbuf
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/xdrbuf
type Xdrbuf struct {
	Xb_flags    unsafe.Pointer
	Xb_growsize uintptr
	Xb_left     uintptr
	Xb_ptr      unsafe.Pointer
	Xb_type     Xdrbuf_type
	Xb_u        unsafe.Pointer
}

// Xinpcb
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/xinpcb
type Xinpcb struct {
	Xi_alignment_hack U_quad_t
	Xi_inp            Inpcb
	Xi_len            U_int32_t
	Xi_socket         Xsocket
}

// Xinpcb64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/xinpcb64
type Xinpcb64 struct {
	Inp_flags  unsafe.Pointer
	Inp_ip_p   U_char
	Inp_ip_ttl U_char
	Inp_phd    U_int64_t
	Inp_ppcb   U_int64_t
	Inp_vflag  U_char
}

// Xinpgen
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/xinpgen
type Xinpgen struct {
	Xig_count U_int
	Xig_gen   Inp_gen_t
	Xig_len   U_int32_t
	Xig_sogen So_gen_t
}

// Xmm_reg
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/xmm_reg
type Xmm_reg struct {
	Xmm_reg unsafe.Pointer
}

// Xsockbuf
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/xsockbuf
type Xsockbuf struct {
	Sb_mbcnt U_int32_t
	Sb_hiwat U_int32_t
	Sb_lowat unsafe.Pointer
	Sb_mbmax U_int32_t
	Sb_cc    U_int32_t
	Sb_timeo unsafe.Pointer
	Sb_flags unsafe.Pointer
}

// Xsocket
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/xsocket
type Xsocket struct {
	So_incqlen   unsafe.Pointer
	Xso_family   unsafe.Pointer
	Xso_len      U_int32_t
	So_oobmark   U_int32_t
	So_qlen      unsafe.Pointer
	Xso_protocol unsafe.Pointer
	Xso_so       U_int32_t
	So_options   unsafe.Pointer
	So_snd       Xsockbuf
	So_pcb       U_int32_t
	So_qlimit    unsafe.Pointer
	So_linger    unsafe.Pointer
	So_pgid      int32
	So_type      unsafe.Pointer
	So_rcv       Xsockbuf
	So_timeo     unsafe.Pointer
	So_error     U_short
	So_state     unsafe.Pointer
	So_uid       uint32
}

// Xsocket64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/xsocket64
type Xsocket64 struct {
	So_oobmark   U_int32_t
	So_rcv       Xsockbuf
	Xso_family   unsafe.Pointer
	So_pgid      int32
	So_linger    unsafe.Pointer
	So_state     unsafe.Pointer
	So_incqlen   unsafe.Pointer
	So_options   unsafe.Pointer
	So_snd       Xsockbuf
	So_error     U_short
	Xso_protocol unsafe.Pointer
	So_timeo     unsafe.Pointer
	So_uid       uint32
	So_qlen      unsafe.Pointer
	So_qlimit    unsafe.Pointer
	So_type      unsafe.Pointer
	Xso_so       U_int64_t
	Xso_len      U_int32_t
	So_pcb       U_int64_t
}

// Xtcpcb
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/xtcpcb
type Xtcpcb struct {
	Xt_alignment_hack U_quad_t
}

// Xtcpcb64
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/xtcpcb64
type Xtcpcb64 struct {
	Cc_recv    Tcp_cc
	Cc_send    Tcp_cc
	Max_sndwnd U_int32_t
}

// Xucred
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/xucred
type Xucred struct {
	Cr_ngroups unsafe.Pointer
	Cr_uid     uint32
	Cr_groups  uint32
}

// Xunpgen
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/xunpgen
type Xunpgen struct {
	Xug_len   U_int32_t
	Xug_gen   Unp_gen_t
	Xug_count U_int
}

// Xvsockpcb
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/xvsockpcb
type Xvsockpcb struct {
	Xv_len         U_int32_t
	Xv_socket      Xsocket
	Xvp_gencnt     Vsock_gen_t
	Xvp_remote_cid U_int32_t
}

// Xvsockpgen
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/xvsockpgen
type Xvsockpgen struct {
	Xvg_count U_int64_t
	Xvg_gen   Vsock_gen_t
	Xvg_len   U_int32_t
	Xvg_sogen So_gen_t
}

// Ymm_reg
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/ymm_reg
type Ymm_reg struct {
	Ymm_reg unsafe.Pointer
}

// Zmm_reg
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/zmm_reg
type Zmm_reg struct {
	Zmm_reg unsafe.Pointer
}
