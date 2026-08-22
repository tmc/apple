// Code generated from Apple documentation. DO NOT EDIT.

package fileprovider

import (
	"github.com/tmc/apple/foundation"
)

// NSFileProviderDomainIdentifier is a unique identifier for a file provider’s domain.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomainIdentifier
type NSFileProviderDomainIdentifier = string

// NSFileProviderExtensionActionIdentifier is an identifier for custom actions.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtensionActionIdentifier
type NSFileProviderExtensionActionIdentifier = string

// NSFileProviderItemDecorationIdentifier is a decoration identifier defined in the File Provider extension’s information property list.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemDecorationIdentifier
type NSFileProviderItemDecorationIdentifier = string

// NSFileProviderItemIdentifier is a unique identifier for an item managed by the File Provider extension.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemIdentifier
type NSFileProviderItemIdentifier = string

// NSFileProviderPage is a synchronization point that represents the next batch of items to be returned by an enumerator.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderPage
type NSFileProviderPage = *foundation.NSData

// NSFileProviderSyncAnchor is a synchronization point that represents the last batch of changes returned by the enumerator.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderSyncAnchor
type NSFileProviderSyncAnchor = *foundation.NSData

// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderUserInfoKey
type NSFileProviderUserInfoKey = string
