// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

// Package corespotlight provides Go bindings for the CoreSpotlight framework.
//
// Add search capabilities to your app, and index your content so people can
// find it from Spotlight and Safari.
//
// Help people access activities and items within your app by adding details
// about those items to a Core Spotlight index. The framework provides APIs to
// add your content to an index, and search for items in that index. You
// decide what content makes sense to index, but typically you index anything
// that someone might look for in your app. For example, you might index
// photos, contacts, the items someone purchased, or data they see in your
// interface. You can then use Core Spotlight to search for your indexed
// content and display those results in your app.
//
// # Essentials
//
//   - [Adding your app’s content to Spotlight indexes]: Create a description for your app’s content and add it to a Spotlight index to make it searchable.
//
// # Searchable items
//
//   - [CSSearchableItem]: The details of your app-specific content that someone might search for on their devices.
//   - [CSSearchableItemAttributeSet]: The detailed metadata for a searchable item.
//   - [CSCustomAttributeKey]: A key associated with a custom attribute for a searchable item.
//   - [CSLocalizedString]: An object that displays localized text in search results related to your app.
//   - [CSPerson]: An object that represents a person in the context of search results.
//
// # Indexes
//
//   - [Generating summary and priority data for indexed items]: Summarize mail, message, and audio transcripts or assess the priority of mail and messages using Spotlight and Apple Intelligence.
//   - [CSSearchableIndex]: An on-device index for your app’s searchable content. ([CSSearchableIndexDelegate])
//   - [CSSearchableIndexDelegate]: A protocol that defines methods a delegate object or app extension uses to handle communication from the on-device index.
//
// # Spotlight app extensions
//
//   - [Regenerating your app’s indexes on demand]: Create an app extension to maintain your app’s indexes and regenerate them as needed.
//   - [CSIndexExtensionRequestHandler]: An interface that implements an index-maintenance app extension.
//   - [CSImportExtension]: An object that provides searchable attributes for file types that the app supports.
//
// # Queries
//
//   - [Building a search interface for your app]: Add a search interface to your app to execute Spotlight queries and offer suggested text completions.
//   - [Searching for information in your app]: Search for app-specific content and refine search results using predicates and filters.
//   - [CSUserQuery]: A type you use to initiate searches from your interface and offer suggested text completions.
//   - [CSUserQueryContext]: The configuration details to apply to a user query.
//   - [CSSearchQuery]: A type you use to programmatically search the indexed app content.
//   - [CSSearchQueryContext]: The behavior configuration to use for a search query.
//   - [CSSuggestion]: The kind of suggestion to use in a query.
//
// # Errors
//
//   - [CSIndex Errors]: Index error codes and error domain.
//   - [CSSearchQuery Errors]: Search query error codes and error domain.
//
// # Version
//
//   - [CoreSpotlightVersionNumber]: The project version number for Core Spotlight.
//   - [CoreSpotlightVersionString]: The project version string for Core Spotlight.
//
// # Variables
//
//   - [CSSuggestionHighlightAttributeName]//
//
// # Key Types
//
//   - [CSSearchableItemAttributeSet] - The detailed metadata for a searchable item.
//   - [CSSearchableIndex] - An on-device index for your app’s searchable content.
//   - [CSUserQuery] - A type you use to initiate searches from your interface and offer suggested text completions.
//   - [CSIndexExtensionRequestHandler] - An interface that implements an index-maintenance app extension.
//   - [CSSearchQuery] - A type you use to programmatically search the indexed app content.
//   - [CSSearchableItem] - The details of your app-specific content that someone might search for on their devices.
//   - [CSCustomAttributeKey] - A key associated with a custom attribute for a searchable item.
//   - [CSUserQueryContext] - The configuration details to apply to a user query.
//   - [CSPerson] - An object that represents a person in the context of search results.
//   - [CSSearchQueryContext] - The behavior configuration to use for a search query.
//
// [Adding your app’s content to Spotlight indexes]: https://developer.apple.com/documentation/corespotlight/adding-your-app-s-content-to-spotlight-indexes
// [Building a search interface for your app]: https://developer.apple.com/documentation/corespotlight/building-a-search-interface-for-your-app
// [CSIndex Errors]: https://developer.apple.com/documentation/corespotlight/csindex-errors
// [CSSearchQuery Errors]: https://developer.apple.com/documentation/corespotlight/cssearchquery-errors
// [Generating summary and priority data for indexed items]: https://developer.apple.com/documentation/corespotlight/generating-summary-and-priority-data-for-indexed-items
// [Regenerating your app’s indexes on demand]: https://developer.apple.com/documentation/corespotlight/regenerating-your-app-s-indexes-on-demand
// [Searching for information in your app]: https://developer.apple.com/documentation/corespotlight/searching-for-information-in-your-app
package corespotlight

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the CoreSpotlight library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/CoreSpotlight.framework/CoreSpotlight",
	"/usr/lib/libCoreSpotlight.dylib",
}

// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	for _, path := range frameworkPaths {
		h, err := purego.Dlopen(path, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
		if err == nil {
			frameworkHandle = h
			return
		}
	}
	// Loading is best-effort: the warning is silent by default because a missing
	// framework is harmless unless one of its symbols is actually called. Set
	// APPLE_FRAMEWORK_LOAD_DEBUG to surface load failures while diagnosing.
	if os.Getenv("APPLE_FRAMEWORK_LOAD_DEBUG") != "" {
		fmt.Fprintf(os.Stderr, "warning: CoreSpotlight: failed to load framework from any known path\n")
	}
}
