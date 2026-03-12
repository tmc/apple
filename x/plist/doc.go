// Package plist provides plist parsing and writing for Apple configuration files.
// Supports XML, binary, and JSON plist formats.
//
// # Parsing
//
// Parse plist data from any format with auto-detection:
//
//	data, _ := os.ReadFile("Info.plist")
//	v, err := plist.ParseBytes(data)
//
//	// Or parse directly from a file
//	v, err := plist.ParseFile("Info.plist")
//
// # Accessing Values
//
// Use typed accessor functions with key paths:
//
//	bundleID := plist.String(v, "CFBundleIdentifier")
//	version := plist.Int(v, "CFBundleVersion")
//
// # Struct Marshal/Unmarshal
//
// Encode and decode Go structs using the "plist" struct tag:
//
//	type Info struct {
//		BundleID string `plist:"CFBundleIdentifier"`
//		Version  int    `plist:"CFBundleVersion"`
//	}
//
//	data, err := plist.Marshal(info, plist.FormatXML)
//
//	var info Info
//	_, err := plist.Unmarshal(data, &info)
//
// # Writing
//
// Write plist data in any supported format:
//
//	plist.WriteXML(os.Stdout, v)
//	plist.WriteJSON(os.Stdout, v)
//	plist.WriteBinary(os.Stdout, v)
//	plist.WriteToFile("output.plist", v, plist.FormatXML)
//
// # NSKeyedArchiver and UID
//
// Binary plists produced by NSKeyedArchiver use UID values to reference
// objects in the $objects array. The [UID] type preserves these references
// when parsing and writing binary plists:
//
//	v, _ := plist.ParseBytes(data)
//	m := v.(map[string]any)
//	top := m["$top"].(map[string]any)
//	ref := top["root"].(plist.UID) // UID referencing an entry in $objects
package plist
