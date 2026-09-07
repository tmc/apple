package img4

import (
	"encoding/asn1"
	"encoding/binary"
	"fmt"
	"math/big"
)

// Manifest contains the asserted properties in an IM4M's MANP and image groups.
// Parsing does not verify the signature, certificate chain or device eligibility.
// Values are uint64, bool, []byte or string. Returned byte values are copies.
type Manifest struct {
	Properties map[string]any
	Images     map[string]map[string]any
}

// ParseManifest reads an IM4M property tree, rejecting duplicate names, mismatched
// private tags and unsupported scalar types. The caller must authenticate the
// ticket separately before relying on the returned assertions.
func ParseManifest(ticket []byte) (Manifest, error) {
	var result Manifest
	fields, err := sequence(ticket, "IM4M")
	if err != nil {
		return result, err
	}
	if len(fields) != 5 || !primitive(fields[1], asn1.TagInteger) || !compound(fields[2], asn1.TagSet) || !primitive(fields[3], asn1.TagOctetString) || !compound(fields[4], asn1.TagSequence) {
		return result, fmt.Errorf("invalid IM4M fields")
	}
	var version int
	if rest, err := asn1.Unmarshal(fields[1].FullBytes, &version); err != nil || len(rest) != 0 || version != 0 {
		return result, fmt.Errorf("unsupported IM4M version")
	}
	roots, err := privateFields(fields[2].Bytes)
	if err != nil {
		return result, err
	}
	root, ok := roots["MANB"]
	if len(roots) != 1 || !ok || !compound(root, asn1.TagSet) {
		return result, fmt.Errorf("IM4M must contain one MANB set")
	}
	groups, err := privateFields(root.Bytes)
	if err != nil {
		return result, err
	}
	result.Images = make(map[string]map[string]any)
	for name, set := range groups {
		if !compound(set, asn1.TagSet) {
			return Manifest{}, fmt.Errorf("manifest group %s is not a set", name)
		}
		fields, err := privateFields(set.Bytes)
		if err != nil {
			return Manifest{}, fmt.Errorf("manifest group %s: %w", name, err)
		}
		properties := make(map[string]any)
		for key, field := range fields {
			value, err := manifestValue(field)
			if err != nil {
				return Manifest{}, fmt.Errorf("manifest property %s.%s: %w", name, key, err)
			}
			properties[key] = value
		}
		if name == "MANP" {
			result.Properties = properties
		} else {
			result.Images[name] = properties
		}
	}
	if result.Properties == nil {
		return Manifest{}, fmt.Errorf("IM4M MANP group is missing")
	}
	return result, nil
}

func privateFields(data []byte) (map[string]asn1.RawValue, error) {
	fields := make(map[string]asn1.RawValue)
	for len(data) > 0 {
		if len(fields) >= 4096 {
			return nil, fmt.Errorf("too many manifest properties")
		}
		var outer, seq, name, value asn1.RawValue
		var err error
		data, err = asn1.Unmarshal(data, &outer)
		if err != nil {
			return nil, err
		}
		if outer.Class != asn1.ClassPrivate || !outer.IsCompound {
			return nil, fmt.Errorf("manifest field is not a private wrapper")
		}
		rest, err := asn1.Unmarshal(outer.Bytes, &seq)
		if err != nil || len(rest) != 0 || !compound(seq, asn1.TagSequence) {
			return nil, fmt.Errorf("invalid manifest property sequence")
		}
		rest, err = asn1.Unmarshal(seq.Bytes, &name)
		if err != nil || !primitive(name, asn1.TagIA5String) || !fourCC(string(name.Bytes)) {
			return nil, fmt.Errorf("invalid manifest property name")
		}
		key := string(name.Bytes)
		if outer.Tag != int(binary.BigEndian.Uint32(name.Bytes)) {
			return nil, fmt.Errorf("manifest private tag differs from %s", key)
		}
		rest, err = asn1.Unmarshal(rest, &value)
		if err != nil || len(rest) != 0 {
			return nil, fmt.Errorf("invalid manifest property %s value", key)
		}
		if _, ok := fields[key]; ok {
			return nil, fmt.Errorf("duplicate manifest property %s", key)
		}
		fields[key] = value
	}
	return fields, nil
}
func manifestValue(v asn1.RawValue) (any, error) {
	if v.Class != asn1.ClassUniversal || v.IsCompound {
		return nil, fmt.Errorf("unsupported manifest value")
	}
	switch v.Tag {
	case asn1.TagOctetString:
		return append([]byte{}, v.Bytes...), nil
	case asn1.TagInteger:
		var n *big.Int
		if _, err := asn1.Unmarshal(v.FullBytes, &n); err != nil {
			return nil, err
		}
		if !n.IsUint64() {
			return nil, fmt.Errorf("manifest integer is outside uint64")
		}
		return n.Uint64(), nil
	case asn1.TagBoolean:
		var value bool
		_, err := asn1.Unmarshal(v.FullBytes, &value)
		return value, err
	case asn1.TagIA5String:
		var value string
		_, err := asn1.UnmarshalWithParams(v.FullBytes, &value, "ia5")
		return value, err
	}
	return nil, fmt.Errorf("unsupported manifest scalar tag %d", v.Tag)
}
