package plist

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

func TestFormatRoundTrips(t *testing.T) {
	now := time.Date(2025, 6, 15, 12, 0, 0, 0, time.UTC)

	orig := map[string]any{
		"name":    "Test App",
		"version": int64(3),
		"pi":      3.14159,
		"enabled": true,
		"tags":    []any{"alpha", "beta"},
		"data":    []byte{0xCA, 0xFE},
		"created": now,
		"nested": map[string]any{
			"inner": "value",
			"count": int64(7),
		},
	}

	for _, format := range []struct {
		name string
		fmt  Format
	}{
		{"XML", FormatXML},
		{"Binary", FormatBinary},
		{"JSON", FormatJSON},
	} {
		t.Run(format.name, func(t *testing.T) {
			data, err := Marshal(orig, format.fmt)
			if err != nil {
				t.Fatalf("Marshal: %v", err)
			}
			if DetectFormat(data) != format.fmt {
				t.Fatalf("DetectFormat mismatch: got %v", DetectFormat(data))
			}
			parsed, err := ParseBytes(data)
			if err != nil {
				t.Fatalf("ParseBytes: %v", err)
			}
			m, ok := parsed.(map[string]any)
			if !ok {
				t.Fatalf("expected map, got %T", parsed)
			}
			if String(m, "name") != "Test App" {
				t.Errorf("name = %q", String(m, "name"))
			}
			if Int(m, "version") != 3 {
				t.Errorf("version = %d", Int(m, "version"))
			}
			if Bool(m, "enabled") != true {
				t.Errorf("enabled = %v", Bool(m, "enabled"))
			}
			nested := Dict(m, "nested")
			if nested == nil {
				t.Fatal("nested is nil")
			}
			if String(nested, "inner") != "value" {
				t.Errorf("nested.inner = %q", String(nested, "inner"))
			}
		})
	}
}

func TestGetSetRemove(t *testing.T) {
	root := CreateEmpty()
	root["app"] = map[string]any{
		"name":    "demo",
		"version": int64(1),
	}

	if String(root, "app.name") != "demo" {
		t.Errorf("Get app.name = %q", String(root, "app.name"))
	}
	if Int(root, "app.version") != 1 {
		t.Errorf("Get app.version = %d", Int(root, "app.version"))
	}

	updated, err := Set(root, "app.version", int64(2))
	if err != nil {
		t.Fatal(err)
	}
	if Int(updated, "app.version") != 2 {
		t.Errorf("after Set, app.version = %d", Int(updated, "app.version"))
	}

	updated, err = Set(updated, "app.settings.theme", "dark")
	if err != nil {
		t.Fatal(err)
	}
	if String(updated, "app.settings.theme") != "dark" {
		t.Errorf("deep Set failed: %v", Get(updated, "app.settings.theme"))
	}

	updated, err = Remove(updated, "app.settings.theme")
	if err != nil {
		t.Fatal(err)
	}
	if Get(updated, "app.settings.theme") != nil {
		t.Error("Remove failed")
	}
}

func TestArrayAccess(t *testing.T) {
	root := map[string]any{
		"items": []any{"zero", "one", "two"},
	}

	arr := Array(root, "items")
	if len(arr) != 3 {
		t.Fatalf("len(items) = %d", len(arr))
	}

	// Access by index.
	if Get(root, "items", "1") != "one" {
		t.Errorf("items.1 = %v", Get(root, "items", "1"))
	}
	if Get(root, "items", "99") != nil {
		t.Error("out of bounds should return nil")
	}
}

func TestTypeOf(t *testing.T) {
	tests := []struct {
		value any
		want  string
	}{
		{"hello", "string"},
		{int64(42), "integer"},
		{3.14, "float"},
		{true, "bool"},
		{[]byte{1}, "data"},
		{time.Now(), "date"},
		{[]any{}, "array"},
		{map[string]any{}, "dictionary"},
		{nil, "unknown"},
	}
	for _, tt := range tests {
		if got := TypeOf(tt.value); got != tt.want {
			t.Errorf("TypeOf(%T) = %q, want %q", tt.value, got, tt.want)
		}
	}
}

func TestMarshalUnmarshalComplexStruct(t *testing.T) {
	type Bookmark struct {
		URL   string `plist:"url"`
		Title string `plist:"title"`
		Order int    `plist:"order"`
	}
	type Browser struct {
		Name      string     `plist:"name"`
		Version   string     `plist:"version"`
		Default   bool       `plist:"default"`
		Bookmarks []Bookmark `plist:"bookmarks"`
		Settings  map[string]string `plist:"settings,omitempty"`
	}

	orig := Browser{
		Name:    "Safari",
		Version: "18.0",
		Default: true,
		Bookmarks: []Bookmark{
			{URL: "https://apple.com", Title: "Apple", Order: 1},
			{URL: "https://go.dev", Title: "Go", Order: 2},
		},
		Settings: map[string]string{
			"homepage": "https://apple.com",
			"engine":   "google",
		},
	}

	for _, format := range []Format{FormatXML, FormatBinary} {
		t.Run(formatName(format), func(t *testing.T) {
			data, err := Marshal(orig, format)
			if err != nil {
				t.Fatalf("Marshal: %v", err)
			}

			var got Browser
			if _, err := Unmarshal(data, &got); err != nil {
				t.Fatalf("Unmarshal: %v", err)
			}

			if got.Name != orig.Name {
				t.Errorf("Name = %q", got.Name)
			}
			if got.Default != orig.Default {
				t.Errorf("Default = %v", got.Default)
			}
			if len(got.Bookmarks) != 2 {
				t.Fatalf("len(Bookmarks) = %d", len(got.Bookmarks))
			}
			if got.Bookmarks[0].URL != "https://apple.com" {
				t.Errorf("Bookmarks[0].URL = %q", got.Bookmarks[0].URL)
			}
			if got.Bookmarks[1].Title != "Go" {
				t.Errorf("Bookmarks[1].Title = %q", got.Bookmarks[1].Title)
			}
			if got.Settings["homepage"] != "https://apple.com" {
				t.Errorf("Settings[homepage] = %q", got.Settings["homepage"])
			}
		})
	}
}

func TestWriteXMLRoundTrip(t *testing.T) {
	orig := map[string]any{
		"escaped": "< & >",
		"nested": map[string]any{
			"list": []any{int64(1), int64(2), int64(3)},
		},
	}

	var buf bytes.Buffer
	if err := WriteXML(&buf, orig); err != nil {
		t.Fatal(err)
	}
	xml := buf.String()
	if !strings.Contains(xml, "&lt;") {
		t.Error("XML escaping missing for <")
	}
	if !strings.Contains(xml, "&amp;") {
		t.Error("XML escaping missing for &")
	}

	parsed, err := ParseBytes(buf.Bytes())
	if err != nil {
		t.Fatalf("re-parse: %v", err)
	}
	if String(parsed, "escaped") != "< & >" {
		t.Errorf("escaped = %q", String(parsed, "escaped"))
	}
}

func TestWriteJSONRoundTrip(t *testing.T) {
	orig := map[string]any{
		"name":  "test",
		"items": []any{int64(1), int64(2)},
	}
	var buf bytes.Buffer
	if err := WriteJSON(&buf, orig); err != nil {
		t.Fatal(err)
	}
	parsed, err := ParseBytes(buf.Bytes())
	if err != nil {
		t.Fatal(err)
	}
	if String(parsed, "name") != "test" {
		t.Errorf("name = %q", String(parsed, "name"))
	}
}

func TestParseFromReader(t *testing.T) {
	xml := `<?xml version="1.0"?>
<plist version="1.0">
<dict>
	<key>hello</key>
	<string>world</string>
</dict>
</plist>`

	v, err := Parse(strings.NewReader(xml))
	if err != nil {
		t.Fatal(err)
	}
	if String(v, "hello") != "world" {
		t.Errorf("hello = %q", String(v, "hello"))
	}
}

func TestBinaryUnicodeString(t *testing.T) {
	orig := map[string]any{
		"emoji": "Hello 🌍",
		"cjk":   "日本語",
	}
	data, err := Marshal(orig, FormatBinary)
	if err != nil {
		t.Fatal(err)
	}
	parsed, err := ParseBytes(data)
	if err != nil {
		t.Fatal(err)
	}
	if String(parsed, "emoji") != "Hello 🌍" {
		t.Errorf("emoji = %q", String(parsed, "emoji"))
	}
	if String(parsed, "cjk") != "日本語" {
		t.Errorf("cjk = %q", String(parsed, "cjk"))
	}
}

func formatName(f Format) string {
	switch f {
	case FormatXML:
		return "XML"
	case FormatBinary:
		return "Binary"
	case FormatJSON:
		return "JSON"
	default:
		return "Unknown"
	}
}
