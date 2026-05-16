// Command txtrecord creates and reads a Network TXT record.
//
// It demonstrates the []byte value parameter for NWTXTRecordSetKey and
// the typed block callbacks used to inspect TXT record contents.
package main

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/tmc/apple/network"
)

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	record := network.NWTXTRecordCreateDictionary()
	if record.ID == 0 {
		panic("nw_txt_record_create_dictionary returned nil")
	}
	defer record.Release()

	value := []byte("apple")
	if !network.NWTXTRecordSetKey(record, "framework", value, uintptr(len(value))) {
		panic("nw_txt_record_set_key failed")
	}

	ok := network.NWTXTRecordAccessKeyFunc(record, "framework",
		func(key string, status network.NWTXTRecordFindKey, value *uint8, valueLen uint32) bool {
			bytes := unsafe.Slice(value, valueLen)
			fmt.Printf("%s=%s status=%s\n", key, string(bytes), status)
			return true
		})
	if !ok {
		panic("nw_txt_record_access_key failed")
	}
}
