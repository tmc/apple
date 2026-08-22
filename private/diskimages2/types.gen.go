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

// AAS3ContextImpl
type AAS3ContextImpl struct {
}

// AAS3Context_impl is a type alias for AAS3ContextImpl for use in objc.Send[T] calls.
type AAS3Context_impl = AAS3ContextImpl

// AEAAuthDataImpl
type AEAAuthDataImpl struct {
}

// AEAAuthData_impl is a type alias for AEAAuthDataImpl for use in objc.Send[T] calls.
type AEAAuthData_impl = AEAAuthDataImpl

// AuthData
type AuthData struct {
	Field1 AEAAuthDataImplRef
}

// AuthorizationOpaqueRef
type AuthorizationOpaqueRef struct {
}

// Backend
type Backend struct {
}

// BaseFolderCopier
type BaseFolderCopier struct {
}

// DiskImageInfo
type DiskImageInfo struct {
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

// Repr
type Repr struct {
	Field1 [2]uint64
	Field2 bool
}

// SharedWeakCount
type SharedWeakCount struct {
}

// Union
type Union struct {
	Field1 [2]uint64
	Field2 ErrorCode
}

// AuthEntryDescriptor
type AuthEntryDescriptor struct {
}

// Auth_entry_descriptor is a type alias for AuthEntryDescriptor for use in objc.Send[T] calls.
type Auth_entry_descriptor = AuthEntryDescriptor

// AuthTable
type AuthTable struct {
	Descriptors [3]uint64
	Reader      [2]uint64
}

// Auth_table is a type alias for AuthTable for use in objc.Send[T] calls.
type Auth_table = AuthTable

// AuthTableReader
type AuthTableReader struct {
}

// Auth_table_reader_t is a type alias for AuthTableReader for use in objc.Send[T] calls.
type Auth_table_reader_t = AuthTableReader

// CryptoSerializer
type CryptoSerializer struct {
	Field1 unsafe.Pointer
	Field2 [3]uint64
	Field3 FormatRef
	Field4 [2]uint64
}

// Crypto_serializer_t is a type alias for CryptoSerializer for use in objc.Send[T] calls.
type Crypto_serializer_t = CryptoSerializer

// Diskimage
type Diskimage struct {
}

// DiskimageOpenParams
type DiskimageOpenParams struct {
	Field1 DiskimageOpenParamsImplRef
}

// Diskimage_open_params is a type alias for DiskimageOpenParams for use in objc.Send[T] calls.
type Diskimage_open_params = DiskimageOpenParams

// DiskimageOpenParamsImpl
type DiskimageOpenParamsImpl struct {
}

// Diskimage_open_params_impl is a type alias for DiskimageOpenParamsImpl for use in objc.Send[T] calls.
type Diskimage_open_params_impl = DiskimageOpenParamsImpl

// ErrorCategory
type ErrorCategory struct {
}

// Error_category is a type alias for ErrorCategory for use in objc.Send[T] calls.
type Error_category = ErrorCategory

// ErrorCode
type ErrorCode struct {
	Field1 int32
	Field2 ErrorCategoryRef
}

// Error_code is a type alias for ErrorCode for use in objc.Send[T] calls.
type Error_code = ErrorCode

// Format
type Format struct {
}

// Fsid
type Fsid struct {
	Field1 [2]int32
}

// Header
type Header struct {
}

// PassphraseHeader
type PassphraseHeader struct {
}

// Passphrase_header is a type alias for PassphraseHeader for use in objc.Send[T] calls.
type Passphrase_header = PassphraseHeader

// QtnFile
type QtnFile struct {
}

// Qtn_file is a type alias for QtnFile for use in objc.Send[T] calls.
type Qtn_file = QtnFile

// Statfs
type Statfs struct {
	Field1  uint32
	Field2  int32
	Field3  uint64
	Field4  uint64
	Field5  uint64
	Field6  uint64
	Field7  uint64
	Field8  Fsid
	Field9  uint32
	Field10 uint32
	Field11 uint32
	Field12 uint32
	Field13 [16]int8
	Field14 [1024]int8
	Field15 [1024]int8
	Field16 uint32
	Field17 [7]uint32
}

// Type
type Type struct {
	Field1 [24]uint8
}
