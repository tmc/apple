package coremlcompiler

import "testing"

func TestDataTypeToBlobDataType(t *testing.T) {
	tests := []struct {
		name    string
		dt      DataType
		want    BlobDataType
		wantErr bool
	}{
		{"float16", DataTypeFloat16, BlobDataTypeFloat16, false},
		{"int32", DataTypeInt32, BlobDataTypeInt32, false},
		{"uint32", DataTypeUInt32, BlobDataTypeUInt32, false},
		// MILBlob has no 64-bit integer type; mapping uint64 to a 32-bit tag
		// would make the reader see twice as many bogus elements.
		{"uint64", DataTypeUInt64, 0, true},
		{"int64", DataTypeInt64, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DataTypeToBlobDataType(tt.dt)
			if (err != nil) != tt.wantErr {
				t.Fatalf("DataTypeToBlobDataType(%v) err = %v, wantErr %v", tt.dt, err, tt.wantErr)
			}
			if err == nil && got != tt.want {
				t.Fatalf("DataTypeToBlobDataType(%v) = %v, want %v", tt.dt, got, tt.want)
			}
		})
	}
}

func TestShouldUseWeightFile(t *testing.T) {
	tests := []struct {
		name        string
		dt          DataType
		numElements int
		specVersion int32
		want        bool
	}{
		{"float32 large", DataTypeFloat32, 10, 7, true},
		{"float32 small", DataTypeFloat32, 9, 7, false},
		{"int32 pre-ios18", DataTypeInt32, 100, 8, false},
		{"int32 ios18", DataTypeInt32, 100, 9, true},
		{"uint16 ios18", DataTypeUInt16, 100, 9, true},
		{"uint64 ios18", DataTypeUInt64, 100, 9, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShouldUseWeightFile(tt.dt, tt.numElements, tt.specVersion); got != tt.want {
				t.Fatalf("ShouldUseWeightFile(%v, %d, %d) = %v, want %v", tt.dt, tt.numElements, tt.specVersion, got, tt.want)
			}
		})
	}
}
