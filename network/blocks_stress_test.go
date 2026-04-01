package network

import (
	"testing"
)

const networkBlockStressCount = 4096

func newTestTXTRecord(t testing.TB) Nw_txt_record_t {
	t.Helper()

	record := Nw_txt_record_create_dictionary()
	if record.ID == 0 {
		t.Fatal("Nw_txt_record_create_dictionary returned nil")
	}

	value := []byte("value")
	if !Nw_txt_record_set_key(record, "key", &value[0], uintptr(len(value))) {
		record.Release()
		t.Fatal("Nw_txt_record_set_key failed")
	}
	if !Nw_txt_record_is_dictionary(record) {
		record.Release()
		t.Fatal("Nw_txt_record_is_dictionary = false")
	}
	if got := Nw_txt_record_get_key_count(record); got != 1 {
		record.Release()
		t.Fatalf("Nw_txt_record_get_key_count = %d, want 1", got)
	}

	return record
}

func TestTXTRecordAccessBytesMoreThanBlockLimit(t *testing.T) {
	record := newTestTXTRecord(t)
	defer record.Release()

	callbacks := 0
	for i := 0; i < networkBlockStressCount; i++ {
		ok := Nw_txt_record_access_bytes(record, func(value *uint8, valueLen uint32) bool {
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
			t.Fatalf("Nw_txt_record_access_bytes failed at iteration %d", i)
		}
	}

	if callbacks != networkBlockStressCount {
		t.Fatalf("callback count = %d, want %d", callbacks, networkBlockStressCount)
	}
}

func TestParametersIterateProhibitedInterfaceTypesMoreThanBlockLimit(t *testing.T) {
	parameters := Nw_parameters_create()
	if parameters.ID == 0 {
		t.Fatal("Nw_parameters_create returned nil")
	}
	defer parameters.Release()

	Nw_parameters_prohibit_interface_type(parameters, Nw_interface_type_wifi)

	callbacks := 0
	for i := 0; i < networkBlockStressCount; i++ {
		iterationCallbacks := 0
		Nw_parameters_iterate_prohibited_interface_types(parameters, func(interfaceType NwInterfaceType) bool {
			if interfaceType != Nw_interface_type_wifi {
				t.Fatalf("callback %d interface type = %v, want %v", i, interfaceType, Nw_interface_type_wifi)
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
		ok := Nw_txt_record_access_key(record, "key", func(key string, status NwTxtRecordFindKey, value *uint8, valueLen uint32) bool {
			if key != "key" {
				t.Fatalf("callback %d key = %q, want %q", i, key, "key")
			}
			if status != Nw_txt_record_find_key_non_empty_value {
				t.Fatalf("callback %d status = %v, want %v", i, status, Nw_txt_record_find_key_non_empty_value)
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
			t.Fatalf("Nw_txt_record_access_key failed at iteration %d", i)
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
		ok := Nw_txt_record_apply(record, func(key string, status NwTxtRecordFindKey, value *uint8, valueLen uint32) bool {
			if key != "key" {
				t.Fatalf("callback %d key = %q, want %q", i, key, "key")
			}
			if status != Nw_txt_record_find_key_non_empty_value {
				t.Fatalf("callback %d status = %v, want %v", i, status, Nw_txt_record_find_key_non_empty_value)
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
			t.Fatalf("Nw_txt_record_apply failed at iteration %d", i)
		}
		if iterationCallbacks != 1 {
			t.Fatalf("iteration %d callback count = %d, want 1", i, iterationCallbacks)
		}
	}

	if callbacks != networkBlockStressCount {
		t.Fatalf("callback count = %d, want %d", callbacks, networkBlockStressCount)
	}
}
