package network

import (
	"testing"
)

const networkBlockStressCount = 4096

func newTestTXTRecord(t testing.TB) NWTXTRecord {
	t.Helper()

	record := NWTXTRecordCreateDictionary()
	if record.ID == 0 {
		t.Fatal("NWTXTRecordCreateDictionary returned nil")
	}

	value := []byte("value")
	if !NWTXTRecordSetKey(record, "key", value, uintptr(len(value))) {
		record.Release()
		t.Fatal("NWTXTRecordSetKey failed")
	}
	if !NWTXTRecordIsDictionary(record) {
		record.Release()
		t.Fatal("NWTXTRecordIsDictionary = false")
	}
	if got := NWTXTRecordGetKeyCount(record); got != 1 {
		record.Release()
		t.Fatalf("NWTXTRecordGetKeyCount = %d, want 1", got)
	}

	return record
}

func TestTXTRecordAccessBytesMoreThanBlockLimit(t *testing.T) {
	record := newTestTXTRecord(t)
	defer record.Release()

	callbacks := 0
	for i := 0; i < networkBlockStressCount; i++ {
		ok := NWTXTRecordAccessBytesFunc(record, func(value *uint8, valueLen uint32) bool {
			if valueLen == 0 {
				t.Fatalf("callback %d returned empty data", i)
			}
			if value == nil {
				t.Fatalf("callback %d returned nil data", i)
			}
			callbacks++
			return true
		})
		if !ok {
			t.Fatalf("NWTXTRecordAccessBytesFunc failed at iteration %d", i)
		}
	}

	if callbacks != networkBlockStressCount {
		t.Fatalf("callback count = %d, want %d", callbacks, networkBlockStressCount)
	}
}

func TestParametersIterateProhibitedInterfaceTypesMoreThanBlockLimit(t *testing.T) {
	parameters := NWParametersCreate()
	if parameters.ID == 0 {
		t.Fatal("NWParametersCreate returned nil")
	}
	defer parameters.Release()

	NWParametersProhibitInterfaceType(parameters, NWInterfaceTypeWifi)

	callbacks := 0
	for i := 0; i < networkBlockStressCount; i++ {
		iterationCallbacks := 0
		NWParametersIterateProhibitedInterfaceTypes(parameters, func(interfaceType NWInterfaceType) bool {
			if interfaceType != NWInterfaceTypeWifi {
				t.Fatalf("callback %d interface type = %v, want %v", i, interfaceType, NWInterfaceTypeWifi)
			}
			iterationCallbacks++
			callbacks++
			return true
		})
		if iterationCallbacks != 1 {
			t.Fatalf("iteration %d callback count = %d, want 1", i, iterationCallbacks)
		}
	}

	if callbacks != networkBlockStressCount {
		t.Fatalf("callback count = %d, want %d", callbacks, networkBlockStressCount)
	}
}

func TestTXTRecordAccessKeyStringCallbackMoreThanBlockLimit(t *testing.T) {
	record := newTestTXTRecord(t)
	defer record.Release()

	callbacks := 0
	for i := 0; i < networkBlockStressCount; i++ {
		ok := NWTXTRecordAccessKeyFunc(record, "key", func(key string, status NWTXTRecordFindKey, value *uint8, valueLen uint32) bool {
			if key != "key" {
				t.Fatalf("callback %d key = %q, want %q", i, key, "key")
			}
			if status != NWTXTRecordFindKeyNonEmptyValue {
				t.Fatalf("callback %d status = %v, want %v", i, status, NWTXTRecordFindKeyNonEmptyValue)
			}
			if value == nil {
				t.Fatalf("callback %d returned nil value", i)
			}
			if valueLen == 0 {
				t.Fatalf("callback %d returned empty value", i)
			}
			callbacks++
			return true
		})
		if !ok {
			t.Fatalf("NWTXTRecordAccessKeyFunc failed at iteration %d", i)
		}
	}

	if callbacks != networkBlockStressCount {
		t.Fatalf("callback count = %d, want %d", callbacks, networkBlockStressCount)
	}
}

func TestTXTRecordApplyStringCallbackMoreThanBlockLimit(t *testing.T) {
	record := newTestTXTRecord(t)
	defer record.Release()

	callbacks := 0
	for i := 0; i < networkBlockStressCount; i++ {
		iterationCallbacks := 0
		ok := NWTXTRecordApply(record, func(key string, status NWTXTRecordFindKey, value *uint8, valueLen uint32) bool {
			if key != "key" {
				t.Fatalf("callback %d key = %q, want %q", i, key, "key")
			}
			if status != NWTXTRecordFindKeyNonEmptyValue {
				t.Fatalf("callback %d status = %v, want %v", i, status, NWTXTRecordFindKeyNonEmptyValue)
			}
			if value == nil {
				t.Fatalf("callback %d returned nil value", i)
			}
			if valueLen == 0 {
				t.Fatalf("callback %d returned empty value", i)
			}
			iterationCallbacks++
			callbacks++
			return true
		})
		if !ok {
			t.Fatalf("NWTXTRecordApply failed at iteration %d", i)
		}
		if iterationCallbacks != 1 {
			t.Fatalf("iteration %d callback count = %d, want 1", i, iterationCallbacks)
		}
	}

	if callbacks != networkBlockStressCount {
		t.Fatalf("callback count = %d, want %d", callbacks, networkBlockStressCount)
	}
}
