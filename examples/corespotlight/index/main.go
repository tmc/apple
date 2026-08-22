// Command index indexes a few items into Spotlight with Core Spotlight and
// then deletes them again by unique identifier, leaving the index as it was.
//
// Usage: index [-n count] [-keep]
//
// With -keep the items are left in the index; otherwise they are removed
// before the program exits.
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/tmc/apple/corespotlight"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

const (
	domainIdentifier = "dev.tmc.apple.examples.corespotlight"
	waitTimeout      = 10 * time.Second
)

func main() {
	count := flag.Int("n", 3, "number of items to index")
	keep := flag.Bool("keep", false, "leave the items in the index instead of deleting them")
	flag.Parse()

	if *count < 1 {
		fmt.Fprintf(os.Stderr, "index: -n must be at least 1\n")
		os.Exit(1)
	}

	if !corespotlight.GetCSSearchableIndexClass().IsIndexingAvailable() {
		fmt.Fprintf(os.Stderr, "index: Spotlight indexing is not available on this device\n")
		os.Exit(1)
	}
	idx := corespotlight.GetCSSearchableIndexClass().DefaultSearchableIndex()
	if idx.ID == 0 {
		fmt.Fprintf(os.Stderr, "index: no default searchable index\n")
		os.Exit(1)
	}

	items := make([]corespotlight.CSSearchableItem, 0, *count)
	ids := make([]string, 0, *count)
	for i := 0; i < *count; i++ {
		attrs := corespotlight.NewCSSearchableItemAttributeSet()
		attrs.SetContentType("public.plain-text")
		attrs.SetTitle(fmt.Sprintf("Go Core Spotlight example %d", i+1))
		attrs.SetDisplayName(fmt.Sprintf("Example item %d", i+1))
		attrs.SetContentDescription("Temporary item indexed by the apple/corespotlight example.")
		attrs.SetKeywords([]string{"golang", "purego", "corespotlight"})

		id := fmt.Sprintf("%s.item.%d", domainIdentifier, i+1)
		item := corespotlight.NewCSSearchableItemWithUniqueIdentifierDomainIdentifierAttributeSet(id, domainIdentifier, attrs)
		items = append(items, item)
		ids = append(ids, id)
	}

	if err := indexItems(idx, items); err != nil {
		fmt.Fprintf(os.Stderr, "index: indexing failed: %v\n", err)
		os.Exit(1)
	}
	for _, id := range ids {
		fmt.Printf("indexed %s\n", id)
	}

	if *keep {
		fmt.Printf("keeping %d item(s) in the index; remove them with:\n", len(ids))
		fmt.Printf("  mdfind \"kMDItemDomainIdentifier == '%s'\"\n", domainIdentifier)
		return
	}

	if err := deleteItems(idx, ids); err != nil {
		fmt.Fprintf(os.Stderr, "index: delete failed: %v\n", err)
		os.Exit(1)
	}
	for _, id := range ids {
		fmt.Printf("deleted %s\n", id)
	}
}

// indexItems adds items to idx and waits for the index to acknowledge them.
//
// It does not use [corespotlight.CSSearchableIndex.IndexSearchableItemsCompletionHandler]:
// that binding passes the Go slice straight to objc_msgSend instead of an
// NSArray, which crashes. Build the NSArray here and send the selector
// directly.
func indexItems(idx corespotlight.CSSearchableIndex, items []corespotlight.CSSearchableItem) error {
	array := foundation.NewNSMutableArray()
	for _, item := range items {
		array.AddObject(item)
	}

	done := make(chan error, 1)
	block, release := corespotlight.NewErrorBlock(func(err error) {
		done <- err
	})
	defer release()
	objc.Send[objc.ID](idx.ID, objc.Sel("indexSearchableItems:completionHandler:"), array.ID, block)
	return wait(done, "indexSearchableItems")
}

// deleteItems removes the items with the given identifiers from idx and waits
// for the index to acknowledge the deletion.
//
// As in indexItems, the generated binding takes a Go slice and passes it
// through unconverted, so the NSArray is built here instead.
func deleteItems(idx corespotlight.CSSearchableIndex, ids []string) error {
	array := foundation.NewNSMutableArray()
	for _, id := range ids {
		objc.Send[objc.ID](array.ID, objc.Sel("addObject:"), objc.String(id))
	}

	done := make(chan error, 1)
	block, release := corespotlight.NewErrorBlock(func(err error) {
		done <- err
	})
	defer release()
	objc.Send[objc.ID](idx.ID, objc.Sel("deleteSearchableItemsWithIdentifiers:completionHandler:"), array.ID, block)
	return wait(done, "deleteSearchableItemsWithIdentifiers")
}

func wait(done <-chan error, what string) error {
	select {
	case err := <-done:
		return err
	case <-time.After(waitTimeout):
		return errors.New(what + ": timed out waiting for the Spotlight index; the indexing daemon may not be running or the process may lack permission")
	}
}
