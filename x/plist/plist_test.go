package plist

import (
	"bytes"
	"testing"
	"time"
)

func TestDetectFormat(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want Format
	}{
		{"binary", []byte("bplist00..."), FormatBinary},
		{"xml", []byte(`<?xml version="1.0"?><plist><dict></dict></plist>`), FormatXML},
		{"json", []byte(`{"key": "val"}`), FormatJSON},
		{"json array", []byte(`[1,2,3]`), FormatJSON},
		{"xml with whitespace", []byte("  <?xml"), FormatXML},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DetectFormat(tt.data); got != tt.want {
				t.Errorf("DetectFormat = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseXML(t *testing.T) {
	data := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>Name</key>
	<string>Test</string>
	<key>Count</key>
	<integer>42</integer>
	<key>Enabled</key>
	<true/>
	<key>Items</key>
	<array>
		<string>a</string>
		<string>b</string>
	</array>
</dict>
</plist>`)

	v, err := ParseBytes(data)
	if err != nil {
		t.Fatal(err)
	}
	m, ok := v.(map[string]any)
	if !ok {
		t.Fatalf("expected map, got %T", v)
	}
	if m["Name"] != "Test" {
		t.Errorf("Name = %v", m["Name"])
	}
	if m["Count"] != int64(42) {
		t.Errorf("Count = %v", m["Count"])
	}
	if m["Enabled"] != true {
		t.Errorf("Enabled = %v", m["Enabled"])
	}
}

func TestMarshalUnmarshalXMLStruct(t *testing.T) {
	type Inner struct {
		Value int `plist:"value"`
	}
	type Config struct {
		Name    string   `plist:"name"`
		Count   int      `plist:"count"`
		Enabled bool     `plist:"enabled"`
		Tags    []string `plist:"tags"`
		Inner   *Inner   `plist:"inner"`
		Secret  string   `plist:"secret,omitempty"`
	}

	orig := Config{
		Name:    "test",
		Count:   42,
		Enabled: true,
		Tags:    []string{"a", "b"},
		Inner:   &Inner{Value: 7},
	}

	data, err := Marshal(orig, FormatXML)
	if err != nil {
		t.Fatal(err)
	}

	var got Config
	if _, err := Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}

	if got.Name != orig.Name {
		t.Errorf("Name = %q, want %q", got.Name, orig.Name)
	}
	if got.Count != orig.Count {
		t.Errorf("Count = %d, want %d", got.Count, orig.Count)
	}
	if got.Enabled != orig.Enabled {
		t.Errorf("Enabled = %v, want %v", got.Enabled, orig.Enabled)
	}
	if len(got.Tags) != 2 || got.Tags[0] != "a" || got.Tags[1] != "b" {
		t.Errorf("Tags = %v, want [a b]", got.Tags)
	}
	if got.Inner == nil || got.Inner.Value != 7 {
		t.Errorf("Inner = %v, want &{7}", got.Inner)
	}
	if got.Secret != "" {
		t.Errorf("Secret = %q, want empty (omitempty)", got.Secret)
	}
}

func TestMarshalUnmarshalBinaryStruct(t *testing.T) {
	type PBKDF2Hash struct {
		Entropy    []byte `plist:"entropy"`
		Iterations int    `plist:"iterations"`
		Salt       []byte `plist:"salt"`
	}
	type ShadowHashData struct {
		SaltedSHA512PBKDF2 *PBKDF2Hash `plist:"SALTED-SHA512-PBKDF2,omitempty"`
	}

	orig := ShadowHashData{
		SaltedSHA512PBKDF2: &PBKDF2Hash{
			Entropy:    []byte{0x01, 0x02, 0x03, 0x04},
			Iterations: 45000,
			Salt:       []byte{0xAA, 0xBB, 0xCC},
		},
	}

	data, err := Marshal(orig, FormatBinary)
	if err != nil {
		t.Fatal(err)
	}

	// Verify magic header.
	if !bytes.HasPrefix(data, []byte("bplist00")) {
		t.Fatalf("missing bplist00 header")
	}

	var got ShadowHashData
	if _, err := Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}

	if got.SaltedSHA512PBKDF2 == nil {
		t.Fatal("SaltedSHA512PBKDF2 is nil")
	}
	if got.SaltedSHA512PBKDF2.Iterations != 45000 {
		t.Errorf("Iterations = %d, want 45000", got.SaltedSHA512PBKDF2.Iterations)
	}
	if !bytes.Equal(got.SaltedSHA512PBKDF2.Entropy, orig.SaltedSHA512PBKDF2.Entropy) {
		t.Errorf("Entropy mismatch")
	}
	if !bytes.Equal(got.SaltedSHA512PBKDF2.Salt, orig.SaltedSHA512PBKDF2.Salt) {
		t.Errorf("Salt mismatch")
	}
}

func TestMarshalUnmarshalBinaryMap(t *testing.T) {
	orig := map[string]any{
		"name":  []any{"admin"},
		"gid":   []any{"80"},
		"users": []any{"testuser", "root"},
	}

	data, err := Marshal(orig, FormatBinary)
	if err != nil {
		t.Fatal(err)
	}

	var got map[string]any
	if _, err := Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}

	users, ok := got["users"].([]any)
	if !ok {
		t.Fatalf("users is %T, want []any", got["users"])
	}
	if len(users) != 2 {
		t.Fatalf("len(users) = %d, want 2", len(users))
	}
	if users[0] != "testuser" {
		t.Errorf("users[0] = %v, want testuser", users[0])
	}
}

func TestMarshalUnmarshalBinaryHdiutilInfo(t *testing.T) {
	type hdiutilImagePart struct {
		DevEntry   string `plist:"dev-entry"`
		MountPoint string `plist:"mount-point"`
	}
	type hdiutilImage struct {
		ImagePath   string             `plist:"image-path"`
		SystemImage bool               `plist:"system-image"`
		Entities    []hdiutilImagePart `plist:"system-entities"`
	}
	type hdiutilInfo struct {
		Images []hdiutilImage `plist:"images"`
	}

	orig := hdiutilInfo{
		Images: []hdiutilImage{
			{
				ImagePath:   "/path/to/disk.img",
				SystemImage: false,
				Entities: []hdiutilImagePart{
					{DevEntry: "/dev/disk5", MountPoint: "/Volumes/Data"},
				},
			},
		},
	}

	// Test XML round-trip (hdiutil outputs XML).
	data, err := Marshal(orig, FormatXML)
	if err != nil {
		t.Fatal(err)
	}

	var got hdiutilInfo
	if _, err := Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}

	if len(got.Images) != 1 {
		t.Fatalf("len(Images) = %d, want 1", len(got.Images))
	}
	if got.Images[0].ImagePath != "/path/to/disk.img" {
		t.Errorf("ImagePath = %q", got.Images[0].ImagePath)
	}
	if got.Images[0].SystemImage != false {
		t.Errorf("SystemImage = %v", got.Images[0].SystemImage)
	}
	if len(got.Images[0].Entities) != 1 {
		t.Fatalf("len(Entities) = %d, want 1", len(got.Images[0].Entities))
	}
	if got.Images[0].Entities[0].DevEntry != "/dev/disk5" {
		t.Errorf("DevEntry = %q", got.Images[0].Entities[0].DevEntry)
	}
}

func TestMarshalOmitempty(t *testing.T) {
	type S struct {
		A string   `plist:"a,omitempty"`
		B int      `plist:"b,omitempty"`
		C []string `plist:"c,omitempty"`
		D bool     `plist:"d,omitempty"`
		E string   `plist:"e"`
	}

	data, err := Marshal(S{E: ""}, FormatXML)
	if err != nil {
		t.Fatal(err)
	}
	s := string(data)
	// a, b, c, d should be omitted; e should be present.
	if bytes.Contains(data, []byte("<key>a</key>")) {
		t.Errorf("omitempty string present: %s", s)
	}
	if bytes.Contains(data, []byte("<key>b</key>")) {
		t.Errorf("omitempty int present: %s", s)
	}
	if bytes.Contains(data, []byte("<key>c</key>")) {
		t.Errorf("omitempty slice present: %s", s)
	}
	if bytes.Contains(data, []byte("<key>d</key>")) {
		t.Errorf("omitempty bool present: %s", s)
	}
	if !bytes.Contains(data, []byte("<key>e</key>")) {
		t.Errorf("non-omitempty field missing: %s", s)
	}
}

func TestMarshalUnmarshalByteSliceSlice(t *testing.T) {
	// This mirrors ShadowHashData [][]byte in UserPlist.
	type U struct {
		Data [][]byte `plist:"data"`
	}

	inner := []byte{0xDE, 0xAD, 0xBE, 0xEF}
	orig := U{Data: [][]byte{inner}}

	data, err := Marshal(orig, FormatBinary)
	if err != nil {
		t.Fatal(err)
	}

	var got U
	if _, err := Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}

	if len(got.Data) != 1 {
		t.Fatalf("len(Data) = %d, want 1", len(got.Data))
	}
	if !bytes.Equal(got.Data[0], inner) {
		t.Errorf("Data[0] = %x, want %x", got.Data[0], inner)
	}
}

func TestWriteXML(t *testing.T) {
	v := map[string]any{
		"key": "value",
		"num": int64(42),
	}
	var buf bytes.Buffer
	if err := WriteXML(&buf, v); err != nil {
		t.Fatal(err)
	}
	s := buf.String()
	if !bytes.Contains([]byte(s), []byte("<key>key</key>")) {
		t.Errorf("missing key in XML output")
	}
	if !bytes.Contains([]byte(s), []byte("<string>value</string>")) {
		t.Errorf("missing value in XML output")
	}
}

func TestBinaryRoundTripDate(t *testing.T) {
	now := time.Date(2025, 6, 15, 12, 0, 0, 0, time.UTC)
	v := map[string]any{
		"date": now,
	}
	data, err := Marshal(v, FormatBinary)
	if err != nil {
		t.Fatal(err)
	}
	parsed, err := ParseBytes(data)
	if err != nil {
		t.Fatal(err)
	}
	m := parsed.(map[string]any)
	got := m["date"].(time.Time)
	if !got.Equal(now) {
		t.Errorf("date = %v, want %v", got, now)
	}
}

func TestBinaryRoundTripFloat(t *testing.T) {
	v := map[string]any{
		"pi": 3.14159,
	}
	data, err := Marshal(v, FormatBinary)
	if err != nil {
		t.Fatal(err)
	}
	parsed, err := ParseBytes(data)
	if err != nil {
		t.Fatal(err)
	}
	m := parsed.(map[string]any)
	if m["pi"] != 3.14159 {
		t.Errorf("pi = %v", m["pi"])
	}
}

func TestMarshalUnmarshalJSONStruct(t *testing.T) {
	type Config struct {
		Name    string   `plist:"name"`
		Count   int      `plist:"count"`
		Enabled bool     `plist:"enabled"`
		Tags    []string `plist:"tags"`
	}

	orig := Config{
		Name:    "test",
		Count:   42,
		Enabled: true,
		Tags:    []string{"a", "b"},
	}

	data, err := Marshal(orig, FormatJSON)
	if err != nil {
		t.Fatal(err)
	}

	// Verify it's valid JSON.
	if data[0] != '{' {
		t.Fatalf("expected JSON object, got %q", data[:10])
	}

	var got Config
	if _, err := Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}

	if got.Name != orig.Name {
		t.Errorf("Name = %q, want %q", got.Name, orig.Name)
	}
	if got.Count != orig.Count {
		t.Errorf("Count = %d, want %d", got.Count, orig.Count)
	}
	if got.Enabled != orig.Enabled {
		t.Errorf("Enabled = %v, want %v", got.Enabled, orig.Enabled)
	}
	if len(got.Tags) != 2 || got.Tags[0] != "a" || got.Tags[1] != "b" {
		t.Errorf("Tags = %v, want [a b]", got.Tags)
	}
}

func TestUnmarshalNonPointer(t *testing.T) {
	var s struct{}
	_, err := Unmarshal([]byte("<plist><dict/></plist>"), s)
	if err == nil {
		t.Error("expected error for non-pointer")
	}
}
