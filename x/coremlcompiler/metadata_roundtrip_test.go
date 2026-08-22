package coremlcompiler

import (
	"encoding/json"
	"testing"
)

func TestModelMetadataRoundTrip(t *testing.T) {
	m := &Model{
		SpecVersion: 8,
		Description: ModelDescription{
			Inputs: []FeatureDescription{{Name: "x"}},
			Metadata: &ModelMetadata{
				ShortDescription: "a test model",
				VersionString:    "1.2",
				Author:           "tmc",
				License:          "BSD",
				UserDefined: map[string]string{
					"com.github.apple.coremltools.version": "8.0",
					"custom":                               "value",
				},
			},
		},
	}

	got, err := decodeModel(EncodeModel(m))
	if err != nil {
		t.Fatalf("decode model: %v", err)
	}
	md := got.Description.Metadata
	if md == nil {
		t.Fatal("metadata dropped in round trip")
	}
	want := m.Description.Metadata
	if md.ShortDescription != want.ShortDescription || md.VersionString != want.VersionString ||
		md.Author != want.Author || md.License != want.License {
		t.Errorf("scalar metadata mismatch: got %+v want %+v", *md, *want)
	}
	for k, v := range want.UserDefined {
		if md.UserDefined[k] != v {
			t.Errorf("userDefined[%q] = %q, want %q", k, md.UserDefined[k], v)
		}
	}

	data, err := buildMetadataJSON(got)
	if err != nil {
		t.Fatalf("build metadata json: %v", err)
	}
	var entries []struct {
		UserDefinedMetadata map[string]string `json:"userDefinedMetadata"`
	}
	if err := json.Unmarshal(data, &entries); err != nil {
		t.Fatalf("unmarshal metadata json: %v", err)
	}
	if len(entries) != 1 {
		t.Fatalf("got %d metadata entries, want 1", len(entries))
	}
	if entries[0].UserDefinedMetadata["custom"] != "value" {
		t.Errorf("metadata.json userDefinedMetadata = %v, want custom=value", entries[0].UserDefinedMetadata)
	}
}
