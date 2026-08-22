//go:build darwin

package oslog_test

import (
	"fmt"

	"github.com/tmc/apple/oslog"
)

func ExampleOSLogStoreScope() {
	fmt.Println(oslog.OSLogStoreCurrentProcessIdentifier)
	fmt.Println(oslog.OSLogStoreSystem)
	// Output:
	// OSLogStoreCurrentProcessIdentifier
	// OSLogStoreSystem
}

func ExampleOSLogEntryLogLevel() {
	fmt.Println(oslog.OSLogEntryLogLevelInfo)
	fmt.Println(oslog.OSLogEntryLogLevelNotice)
	fmt.Println(oslog.OSLogEntryLogLevelError)
	fmt.Println(oslog.OSLogEntryLogLevelFault)
	// Output:
	// OSLogEntryLogLevelInfo
	// OSLogEntryLogLevelNotice
	// OSLogEntryLogLevelError
	// OSLogEntryLogLevelFault
}

func ExampleOSLogEntryStoreCategory() {
	fmt.Println(oslog.OSLogEntryStoreCategoryShortTerm)
	fmt.Println(oslog.OSLogEntryStoreCategoryLongTerm1)
	fmt.Println(oslog.OSLogEntryStoreCategoryLongTerm7)
	// Output:
	// OSLogEntryStoreCategoryShortTerm
	// OSLogEntryStoreCategoryLongTerm1
	// OSLogEntryStoreCategoryLongTerm7
}

func ExampleNewOSLogStoreWithScopeError() {
	store, err := oslog.NewOSLogStoreWithScopeError(oslog.OSLogStoreCurrentProcessIdentifier)
	if err != nil {
		fmt.Printf("failed to create log store: %v\n", err)
		return
	}
	_ = store
	fmt.Println("Created process OSLogStore successfully")
	// Output:
	// Created process OSLogStore successfully
}
