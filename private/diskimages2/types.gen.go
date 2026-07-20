// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"unsafe"
)

// C struct types

// AAAsyncByteStreamImpl
type AAAsyncByteStreamImpl struct {
	Field1 unsafe.Pointer
	Field2 unsafe.Pointer
	Field3 unsafe.Pointer
	Field4 unsafe.Pointer
	Field5 unsafe.Pointer
	Field6 unsafe.Pointer
	Field7 unsafe.Pointer
}

// AAAsyncByteStream_impl is a type alias for AAAsyncByteStreamImpl for use in objc.Send[T] calls.
type AAAsyncByteStream_impl = AAAsyncByteStreamImpl

// AuthData
type AuthData struct {
	Field1 AEAAuthDataImplRef
}

// AuthorizationOpaqueRef
type AuthorizationOpaqueRef struct {
}

// DADisk
type DADisk struct {
}

// DASession
type DASession struct {
}

// MKMedia
type MKMedia struct {
}

// SecCertificate
type SecCertificate struct {
}

// SecKeychainItem
type SecKeychainItem struct {
}

// DiskimageOpenParams
type DiskimageOpenParams struct {
	Field1 DiskimageOpenParamsImplRef
}

// Diskimage_open_params is a type alias for DiskimageOpenParams for use in objc.Send[T] calls.
type Diskimage_open_params = DiskimageOpenParams

// QtnFile
type QtnFile struct {
}

// Qtn_file is a type alias for QtnFile for use in objc.Send[T] calls.
type Qtn_file = QtnFile

// Statfs
type Statfs struct {
	Field1  uint
	Field2  int
	Field3  uint64
	Field4  uint64
	Field5  uint64
	Field6  uint64
	Field7  uint64
	Field8  unsafe.Pointer
	Field9  uint
	Field10 uint
	Field11 uint
	Field12 uint
	Field13 unsafe.Pointer
	Field14 unsafe.Pointer
	Field15 unsafe.Pointer
	Field16 uint
	Field17 unsafe.Pointer
}
