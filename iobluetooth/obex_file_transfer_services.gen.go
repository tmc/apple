// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [OBEXFileTransferServices] class.
var (
	_OBEXFileTransferServicesClass     OBEXFileTransferServicesClass
	_OBEXFileTransferServicesClassOnce sync.Once
)

func getOBEXFileTransferServicesClass() OBEXFileTransferServicesClass {
	_OBEXFileTransferServicesClassOnce.Do(func() {
		_OBEXFileTransferServicesClass = OBEXFileTransferServicesClass{class: objc.GetClass("OBEXFileTransferServices")}
	})
	return _OBEXFileTransferServicesClass
}

// GetOBEXFileTransferServicesClass returns the class object for OBEXFileTransferServices.
func GetOBEXFileTransferServicesClass() OBEXFileTransferServicesClass {
	return getOBEXFileTransferServicesClass()
}

type OBEXFileTransferServicesClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (oc OBEXFileTransferServicesClass) Class() objc.Class {
	return oc.class
}

// Alloc allocates memory for a new instance of the class.
func (oc OBEXFileTransferServicesClass) Alloc() OBEXFileTransferServices {
	rv := objc.Send[OBEXFileTransferServices](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// Implements advanced OBEX operations in addition to simple PUT and GET.
//
// # Overview
//
// All operations are asynchronous and will callback over a respective
// delegate method if the initial return value is successful. The initial
// return value usually concerns the state of this object where as the
// delegate return value reflects the response of the remote device.
//
// # Initializers
//
//   - [OBEXFileTransferServices.InitWithOBEXSession]: Create a new OBEXFileTransferServices object
//
// # Instance Properties
//
//   - [OBEXFileTransferServices.Delegate]
//   - [OBEXFileTransferServices.SetDelegate]
//
// # Instance Methods
//
//   - [OBEXFileTransferServices.Abort]: Abort the current operation
//   - [OBEXFileTransferServices.ChangeCurrentFolderBackward]: Change to the directory above the current level if not at the root
//   - [OBEXFileTransferServices.ChangeCurrentFolderForwardToPath]: Change the remote path
//   - [OBEXFileTransferServices.ChangeCurrentFolderToRoot]: Asynchronously change to the remote root directory
//   - [OBEXFileTransferServices.ConnectToFTPService]: Connect to a remote device for FTP operations
//   - [OBEXFileTransferServices.ConnectToObjectPushService]: Connect to a remote device for ObjectPush operations. Most of the FTP functionality of this object will be disabled.
//   - [OBEXFileTransferServices.CopyRemoteFileToLocalPath]: Copy a remote file to a local path
//   - [OBEXFileTransferServices.CreateFolder]: Create a folder on the remote target
//   - [OBEXFileTransferServices.CurrentPath]: Get the remote current directory path during an FTP session
//   - [OBEXFileTransferServices.Disconnect]: Disconnect from the remote device
//   - [OBEXFileTransferServices.GetDefaultVCard]: Get the remote default VCard, if it is supported
//   - [OBEXFileTransferServices.IsBusy]: Get the action state of the module
//   - [OBEXFileTransferServices.IsConnected]: Get the connected state of this module.
//   - [OBEXFileTransferServices.RemoveItem]: Remove a remote item.
//   - [OBEXFileTransferServices.RetrieveFolderListing]: Get a remote directory listing
//   - [OBEXFileTransferServices.SendDataTypeName]: Send data to a remote target
//   - [OBEXFileTransferServices.SendFile]: Put a local file to the remote target
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices
type OBEXFileTransferServices struct {
	objectivec.Object
}

// OBEXFileTransferServicesFromID constructs a [OBEXFileTransferServices] from an objc.ID.
//
// Implements advanced OBEX operations in addition to simple PUT and GET.
func OBEXFileTransferServicesFromID(id objc.ID) OBEXFileTransferServices {
	return OBEXFileTransferServices{objectivec.Object{ID: id}}
}

// NOTE: OBEXFileTransferServices adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [OBEXFileTransferServices] class.
//
// # Initializers
//
//   - [IOBEXFileTransferServices.InitWithOBEXSession]: Create a new OBEXFileTransferServices object
//
// # Instance Properties
//
//   - [IOBEXFileTransferServices.Delegate]
//   - [IOBEXFileTransferServices.SetDelegate]
//
// # Instance Methods
//
//   - [IOBEXFileTransferServices.Abort]: Abort the current operation
//   - [IOBEXFileTransferServices.ChangeCurrentFolderBackward]: Change to the directory above the current level if not at the root
//   - [IOBEXFileTransferServices.ChangeCurrentFolderForwardToPath]: Change the remote path
//   - [IOBEXFileTransferServices.ChangeCurrentFolderToRoot]: Asynchronously change to the remote root directory
//   - [IOBEXFileTransferServices.ConnectToFTPService]: Connect to a remote device for FTP operations
//   - [IOBEXFileTransferServices.ConnectToObjectPushService]: Connect to a remote device for ObjectPush operations. Most of the FTP functionality of this object will be disabled.
//   - [IOBEXFileTransferServices.CopyRemoteFileToLocalPath]: Copy a remote file to a local path
//   - [IOBEXFileTransferServices.CreateFolder]: Create a folder on the remote target
//   - [IOBEXFileTransferServices.CurrentPath]: Get the remote current directory path during an FTP session
//   - [IOBEXFileTransferServices.Disconnect]: Disconnect from the remote device
//   - [IOBEXFileTransferServices.GetDefaultVCard]: Get the remote default VCard, if it is supported
//   - [IOBEXFileTransferServices.IsBusy]: Get the action state of the module
//   - [IOBEXFileTransferServices.IsConnected]: Get the connected state of this module.
//   - [IOBEXFileTransferServices.RemoveItem]: Remove a remote item.
//   - [IOBEXFileTransferServices.RetrieveFolderListing]: Get a remote directory listing
//   - [IOBEXFileTransferServices.SendDataTypeName]: Send data to a remote target
//   - [IOBEXFileTransferServices.SendFile]: Put a local file to the remote target
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices
type IOBEXFileTransferServices interface {
	objectivec.IObject

	// Topic: Initializers

	// Create a new OBEXFileTransferServices object
	InitWithOBEXSession(inOBEXSession IIOBluetoothOBEXSession) OBEXFileTransferServices

	// Topic: Instance Properties

	Delegate() objectivec.IObject
	SetDelegate(value objectivec.IObject)

	// Topic: Instance Methods

	// Abort the current operation
	Abort() OBEXError
	// Change to the directory above the current level if not at the root
	ChangeCurrentFolderBackward() OBEXError
	// Change the remote path
	ChangeCurrentFolderForwardToPath(inDirName string) OBEXError
	// Asynchronously change to the remote root directory
	ChangeCurrentFolderToRoot() OBEXError
	// Connect to a remote device for FTP operations
	ConnectToFTPService() OBEXError
	// Connect to a remote device for ObjectPush operations. Most of the FTP functionality of this object will be disabled.
	ConnectToObjectPushService() OBEXError
	// Copy a remote file to a local path
	CopyRemoteFileToLocalPath(inRemoteFileName string, inLocalPathAndName string) OBEXError
	// Create a folder on the remote target
	CreateFolder(inDirName string) OBEXError
	// Get the remote current directory path during an FTP session
	CurrentPath() string
	// Disconnect from the remote device
	Disconnect() OBEXError
	// Get the remote default VCard, if it is supported
	GetDefaultVCard(inLocalPathAndName string) OBEXError
	// Get the action state of the module
	IsBusy() bool
	// Get the connected state of this module.
	IsConnected() bool
	// Remove a remote item.
	RemoveItem(inItemName string) OBEXError
	// Get a remote directory listing
	RetrieveFolderListing() OBEXError
	// Send data to a remote target
	SendDataTypeName(inData foundation.NSData, inType string, inName string) OBEXError
	// Put a local file to the remote target
	SendFile(inLocalPathAndName string) OBEXError
}

// Init initializes the instance.
func (o OBEXFileTransferServices) Init() OBEXFileTransferServices {
	rv := objc.Send[OBEXFileTransferServices](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o OBEXFileTransferServices) Autorelease() OBEXFileTransferServices {
	rv := objc.Send[OBEXFileTransferServices](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewOBEXFileTransferServices creates a new OBEXFileTransferServices instance.
func NewOBEXFileTransferServices() OBEXFileTransferServices {
	class := getOBEXFileTransferServicesClass()
	rv := objc.Send[OBEXFileTransferServices](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Create a new OBEXFileTransferServices object
//
// inOBEXSession: A valid IOBluetoothOBEXSession
//
// # Return Value
//
// # A newly created OBEXFileTransferServices object on success, nil on failure
//
// # Discussion
//
// This object must be constructed with a valid IOBluetoothOBEXSession. The
// given IOBluetoothOBEXSession does not need to be connected to the remote
// server. OBEXFileTransferServices can be manually connected through the
// provided connection methods.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/init(obexSession:)
func NewOBEXFileTransferServicesWithOBEXSession(inOBEXSession IIOBluetoothOBEXSession) OBEXFileTransferServices {
	instance := getOBEXFileTransferServicesClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithOBEXSession:"), inOBEXSession)
	return OBEXFileTransferServicesFromID(rv)
}

// Create a new OBEXFileTransferServices object
//
// inOBEXSession: A valid IOBluetoothOBEXSession
//
// # Return Value
//
// # A newly created OBEXFileTransferServices object on success, nil on failure
//
// # Discussion
//
// This object must be constructed with a valid IOBluetoothOBEXSession. The
// given IOBluetoothOBEXSession does not need to be connected to the remote
// server. OBEXFileTransferServices can be manually connected through the
// provided connection methods.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/init(obexSession:)
func (o OBEXFileTransferServices) InitWithOBEXSession(inOBEXSession IIOBluetoothOBEXSession) OBEXFileTransferServices {
	rv := objc.Send[OBEXFileTransferServices](o.ID, objc.Sel("initWithOBEXSession:"), inOBEXSession)
	return rv
}

// Abort the current operation
//
// # Return Value
//
// kOBEXSuccess, or kOBEXGeneralError if no command is in progress. ABORT
// commands can only be sent on our turn, meaning we may have to timeout if
// the target side never responds to the command in progress. In that case
// this object will call back with a status of kOBEXTimeoutError and an error.
// Further results returned through the fileTransferServicesAbortComplete:
// delegate method if initially successful.
//
// # Discussion
//
// Attempts send an abort request to the remote device. Returns the
// OBEXFileTransferServices object to an idle state though the state of the
// remote device is not guaranteed.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/abort()
func (o OBEXFileTransferServices) Abort() OBEXError {
	rv := objc.Send[OBEXError](o.ID, objc.Sel("abort"))
	return OBEXError(rv)
}

// Change to the directory above the current level if not at the root
//
// # Return Value
//
// kOBEXSuccess or kOBEXSessionBusyError initially. Further results returned
// through the fileTransferServicesPathChangeComplete: delegate method if
// initially successful.
//
// # Discussion
//
// Equivalent to ‘cd ..’ only if remote path is not already at root.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/changeCurrentFolderBackward()
func (o OBEXFileTransferServices) ChangeCurrentFolderBackward() OBEXError {
	rv := objc.Send[OBEXError](o.ID, objc.Sel("changeCurrentFolderBackward"))
	return OBEXError(rv)
}

// Change the remote path
//
// inDirName: The name of the remote folder to be set as current
//
// # Return Value
//
// kOBEXSuccess, kOBEXSessionBusyError, or kOBEXBadArgumentError initially.
// Further results returned through the
// fileTransferServicesPathChangeComplete: delegate method if initially
// successful.
//
// # Discussion
//
// Equivalent to ‘cd dirName’.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/changeCurrentFolderForward(toPath:)
func (o OBEXFileTransferServices) ChangeCurrentFolderForwardToPath(inDirName string) OBEXError {
	rv := objc.Send[OBEXError](o.ID, objc.Sel("changeCurrentFolderForwardToPath:"), objc.String(inDirName))
	return OBEXError(rv)
}

// Asynchronously change to the remote root directory
//
// # Discussion
//
// Equivalent to ‘cd ~/’
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/changeCurrentFolderToRoot()
func (o OBEXFileTransferServices) ChangeCurrentFolderToRoot() OBEXError {
	rv := objc.Send[OBEXError](o.ID, objc.Sel("changeCurrentFolderToRoot"))
	return OBEXError(rv)
}

// Connect to a remote device for FTP operations
//
// # Return Value
//
// kOBEXSuccess, kOBEXSessionBusyError, or kOBEXSessionAlreadyConnectedError,
// kOBEXNoResourcesError initially. Further results returned through the
// fileTransferServicesConnectionComplete: delegate method if initially
// successful.
//
// # Discussion
//
// If the OBEXSession given to OBEXFileTransferServices on creation is not
// connected it can be manually connected through this method.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/connectToFTPService()
func (o OBEXFileTransferServices) ConnectToFTPService() OBEXError {
	rv := objc.Send[OBEXError](o.ID, objc.Sel("connectToFTPService"))
	return OBEXError(rv)
}

// Connect to a remote device for ObjectPush operations. Most of the FTP
// functionality of this object will be disabled.
//
// # Return Value
//
// kOBEXSuccess, kOBEXSessionBusyError, or kOBEXSessionAlreadyConnectedError,
// kOBEXNoResourcesError initially. Further results returned through the
// fileTransferServicesConnectionComplete: delegate method if initially
// successful.
//
// # Discussion
//
// If the OBEXSession given to OBEXFileTransferServices on creation is not
// connected it can be manually connected through this method.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/connectToObjectPushService()
func (o OBEXFileTransferServices) ConnectToObjectPushService() OBEXError {
	rv := objc.Send[OBEXError](o.ID, objc.Sel("connectToObjectPushService"))
	return OBEXError(rv)
}

// Copy a remote file to a local path
//
// inRemoteFileName: The name of the remote file to get
//
// inLocalPathAndName: The path and name of where the received file will go
//
// # Return Value
//
// kOBEXSuccess, kOBEXSessionBusyError, or kOBEXBadArgumentError. initially.
// Further results returned through the fileTransferServicesGetComplete: and
// fileTransferServicesGetProgress: delegate methods if initially successful.
//
// # Discussion
//
// Equivalent to ‘cp remotePath/remoteFileName localPathAndName’.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/copyRemoteFile(_:toLocalPath:)
func (o OBEXFileTransferServices) CopyRemoteFileToLocalPath(inRemoteFileName string, inLocalPathAndName string) OBEXError {
	rv := objc.Send[OBEXError](o.ID, objc.Sel("copyRemoteFile:toLocalPath:"), objc.String(inRemoteFileName), objc.String(inLocalPathAndName))
	return OBEXError(rv)
}

// Create a folder on the remote target
//
// inDirName: The name of the folder to be created
//
// # Return Value
//
// kOBEXSuccess, kOBEXSessionBusyError, or kOBEXBadArgumentError initially.
// Further results returned through the
// fileTransferServicesCreateFolderComplete delegate method if initially
// successful.
//
// # Discussion
//
// Equivalent to ‘mkdir dirName’.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/createFolder(_:)
func (o OBEXFileTransferServices) CreateFolder(inDirName string) OBEXError {
	rv := objc.Send[OBEXError](o.ID, objc.Sel("createFolder:"), objc.String(inDirName))
	return OBEXError(rv)
}

// Get the remote current directory path during an FTP session
//
// # Return Value
//
// # The current path being browsed over FTP
//
// # Discussion
//
// This path is changed with each path-specific command called on
// OBEXFileTransferServices.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/currentPath()
func (o OBEXFileTransferServices) CurrentPath() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("currentPath"))
	return foundation.NSStringFromID(rv).String()
}

// Disconnect from the remote device
//
// # Return Value
//
// kOBEXSuccess, kOBEXSessionNotConnectedError, or kOBEXSessionBusyError
// initially. Further results returned through the
// fileTransferServicesDisconnectionComplete: delegate method if initially
// successful.
//
// # Discussion
//
// The user can manually disconnect the OBEXSession from the remote device if
// they want to. OBEXFileTransferServices will disconnect the OBEXSession at
// release only if it was responsible for opening the connection via a connect
// method.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/disconnect()
func (o OBEXFileTransferServices) Disconnect() OBEXError {
	rv := objc.Send[OBEXError](o.ID, objc.Sel("disconnect"))
	return OBEXError(rv)
}

// Get the remote default VCard, if it is supported
//
// inLocalPathAndName: The path and name of where the received file will go
//
// # Return Value
//
// kOBEXSuccess, kOBEXSessionBusyError, or kOBEXBadArgumentError initially.
// Further results returned through the fileTransferServicesGetComplete: and
// fileTransferServicesGetProgress: delegate methods if initially successful.
//
// # Discussion
//
// # Some devices such as cellphones and computers support default VCards
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/getDefaultVCard(_:)
func (o OBEXFileTransferServices) GetDefaultVCard(inLocalPathAndName string) OBEXError {
	rv := objc.Send[OBEXError](o.ID, objc.Sel("getDefaultVCard:"), objc.String(inLocalPathAndName))
	return OBEXError(rv)
}

// Get the action state of the module
//
// # Return Value
//
// Success or failure code.
//
// # Discussion
//
// OBEXFileTransferServices will be considered “busy” when an operation in
// taking place or has not completed. Calling abort: on this module will not
// automatically reset its busy state. The user will have to wait for the
// operation to complete or for the current operation to timeout.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/isBusy()
func (o OBEXFileTransferServices) IsBusy() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isBusy"))
	return rv
}

// Get the connected state of this module.
//
// # Return Value
//
// Success or failure code.
//
// # Discussion
//
// Asks the OBEXSession that was passed to it on creation if it has an open
// OBEX connection
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/isConnected()
func (o OBEXFileTransferServices) IsConnected() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isConnected"))
	return rv
}

// Remove a remote item.
//
// inItemName: The name of the remote item to be removed
//
// # Return Value
//
// kOBEXSuccess, kOBEXSessionBusyError, or kOBEXBadArgumentError initially.
// Further results returned through the
// fileTransferServicesRemoveItemComplete: delegate method if initially
// successful.
//
// # Discussion
//
// # Not supported for use on Apple computer targets
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/removeItem(_:)
func (o OBEXFileTransferServices) RemoveItem(inItemName string) OBEXError {
	rv := objc.Send[OBEXError](o.ID, objc.Sel("removeItem:"), objc.String(inItemName))
	return OBEXError(rv)
}

// Get a remote directory listing
//
// # Return Value
//
// kOBEXSuccess or kOBEXSessionBusyError initially. Further results returned
// through the fileTransferServicesRetrieveFolderListingComplete: delegate
// method if initially successful.
//
// # Discussion
//
// Equivalent to ‘ls’.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/retrieveFolderListing()
func (o OBEXFileTransferServices) RetrieveFolderListing() OBEXError {
	rv := objc.Send[OBEXError](o.ID, objc.Sel("retrieveFolderListing"))
	return OBEXError(rv)
}

// Send data to a remote target
//
// inData: The data to be sent
//
// inType: The type of the data to be sent that will be used in the OBEX type header,
// usually a mime-type. For example, use “text/x-vCard” when sending
// vCards. This argument is optional.
//
// inName: The name of the file that the data can be referenced as.
//
// # Return Value
//
// kOBEXSuccess, kOBEXSessionBusyError, or kOBEXBadArgumentError initially.
// Further results returned through the fileTransferServicesSendComplete: and
// fileTransferServicesSendProgress: delegate methods if initially successful.
//
// # Discussion
//
// Use this method when you have data to send but no file to read from.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/send(_:type:name:)
func (o OBEXFileTransferServices) SendDataTypeName(inData foundation.NSData, inType string, inName string) OBEXError {
	rv := objc.Send[OBEXError](o.ID, objc.Sel("sendData:type:name:"), inData, objc.String(inType), objc.String(inName))
	return OBEXError(rv)
}

// Put a local file to the remote target
//
// inLocalPathAndName: The name and path of the file to be sent an instance of OBEXFilePut.
//
// # Return Value
//
// kOBEXSuccess, kOBEXSessionBusyError, or kOBEXBadArgumentError initially.
// Further results returned through the fileTransferServicesSendComplete: and
// fileTransferServicesSendProgress: delegate methods if initially successful.
//
// # Discussion
//
// Equivalent to ‘mv inLocalFilePath remoteCurrentPath’.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/sendFile(_:)
func (o OBEXFileTransferServices) SendFile(inLocalPathAndName string) OBEXError {
	rv := objc.Send[OBEXError](o.ID, objc.Sel("sendFile:"), objc.String(inLocalPathAndName))
	return OBEXError(rv)
}

// Create a new OBEXFileTransferServices object
//
// inOBEXSession: A valid IOBluetoothOBEXSession
//
// # Return Value
//
// # A newly created OBEXFileTransferServices object on success, nil on failure
//
// # Discussion
//
// This object must be constructed with a valid IOBluetoothOBEXSession. The
// given IOBluetoothOBEXSession does not need to be connected to the remote
// server. This module can be manually connected through the connect() method.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/withOBEXSession(_:)
func (_OBEXFileTransferServicesClass OBEXFileTransferServicesClass) WithOBEXSession(inOBEXSession IIOBluetoothOBEXSession) OBEXFileTransferServices {
	rv := objc.Send[objc.ID](objc.ID(_OBEXFileTransferServicesClass.class), objc.Sel("withOBEXSession:"), inOBEXSession)
	return OBEXFileTransferServicesFromID(rv)
}

// See: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/delegate
func (o OBEXFileTransferServices) Delegate() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("delegate"))
	return objectivec.Object{ID: rv}
}
func (o OBEXFileTransferServices) SetDelegate(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setDelegate:"), value)
}
