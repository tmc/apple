package img4

import (
	"bytes"
	"encoding/asn1"
	"encoding/binary"
	"fmt"
	"math/big"
	"sort"
)

// Options controls container assembly. An empty FourCC retains the IM4P type.
// RestoreProperties are encoded as IM4R private-tagged properties; supported
// values are bool, int, int64, uint64 and []byte. Byte slices are wire bytes and
// are not reversed, including BNCN. Nil properties omit IM4R; an empty non-nil
// map is rejected. Options does not select components or validate ticket policy.
type Options struct {
	FourCC            string
	RestoreProperties map[string]any
}

// Personalize combines complete DER-encoded IM4P and IM4M objects into IMG4.
// It validates their outer structure but does not authenticate the ticket or
// check its identity/digest properties. Inputs are never modified. The payload,
// optional IM4P fields (including PAYP), and the complete ticket remain byte
// identical; only the requested four-character payload type can change.
func Personalize(payload, ticket []byte, options Options) ([]byte, error) {
	fields, err := sequence(payload, "IM4P")
	if err != nil {
		return nil, err
	}
	if len(fields) < 4 || !primitive(fields[1], asn1.TagIA5String) || !fourCC(string(fields[1].Bytes)) || !primitive(fields[2], asn1.TagIA5String) || !primitive(fields[3], asn1.TagOctetString) {
		return nil, fmt.Errorf("invalid IM4P fields")
	}
	manifest, err := sequence(ticket, "IM4M")
	if err != nil {
		return nil, err
	}
	if len(manifest) != 5 || !primitive(manifest[1], asn1.TagInteger) || !compound(manifest[2], asn1.TagSet) || !primitive(manifest[3], asn1.TagOctetString) || !compound(manifest[4], asn1.TagSequence) {
		return nil, fmt.Errorf("invalid IM4M fields")
	}
	var version int
	if rest, err := asn1.Unmarshal(manifest[1].FullBytes, &version); err != nil || len(rest) != 0 || version != 0 {
		return nil, fmt.Errorf("unsupported IM4M version")
	}
	if options.FourCC != "" {
		if !fourCC(options.FourCC) {
			return nil, fmt.Errorf("payload type must be four ASCII bytes")
		}
		content := append([]byte(nil), fields[0].FullBytes...)
		content = append(content, ia5(options.FourCC)...)
		for _, field := range fields[2:] {
			content = append(content, field.FullBytes...)
		}
		payload, err = wrap(asn1.ClassUniversal, asn1.TagSequence, content)
		if err != nil {
			return nil, err
		}
	}
	content := append(ia5("IMG4"), payload...)
	signed, err := wrap(asn1.ClassContextSpecific, 0, ticket)
	if err != nil {
		return nil, err
	}
	content = append(content, signed...)
	if options.RestoreProperties != nil {
		restore, err := restoreInfo(options.RestoreProperties)
		if err != nil {
			return nil, err
		}
		tagged, err := wrap(asn1.ClassContextSpecific, 1, restore)
		if err != nil {
			return nil, err
		}
		content = append(content, tagged...)
	}
	return wrap(asn1.ClassUniversal, asn1.TagSequence, content)
}

func sequence(data []byte, magic string) ([]asn1.RawValue, error) {
	var root asn1.RawValue
	rest, err := asn1.Unmarshal(data, &root)
	if err != nil {
		return nil, fmt.Errorf("parse %s: %w", magic, err)
	}
	if len(rest) != 0 || !compound(root, asn1.TagSequence) {
		return nil, fmt.Errorf("%s must be one complete DER sequence", magic)
	}
	var fields []asn1.RawValue
	data = root.Bytes
	for len(data) > 0 {
		if len(fields) >= 32 {
			return nil, fmt.Errorf("too many %s fields", magic)
		}
		var field asn1.RawValue
		data, err = asn1.Unmarshal(data, &field)
		if err != nil {
			return nil, fmt.Errorf("parse %s field: %w", magic, err)
		}
		fields = append(fields, field)
	}
	if len(fields) == 0 || !primitive(fields[0], asn1.TagIA5String) || string(fields[0].Bytes) != magic {
		return nil, fmt.Errorf("expected %s magic", magic)
	}
	return fields, nil
}
func primitive(v asn1.RawValue, tag int) bool {
	return v.Class == asn1.ClassUniversal && v.Tag == tag && !v.IsCompound
}
func compound(v asn1.RawValue, tag int) bool {
	return v.Class == asn1.ClassUniversal && v.Tag == tag && v.IsCompound
}
func fourCC(s string) bool {
	if len(s) != 4 {
		return false
	}
	for i := range s {
		if s[i] > 127 {
			return false
		}
	}
	return true
}
func ia5(s string) []byte { return append([]byte{asn1.TagIA5String, byte(len(s))}, s...) }
func wrap(class, tag int, content []byte) ([]byte, error) {
	return asn1.Marshal(asn1.RawValue{Class: class, Tag: tag, IsCompound: true, Bytes: content})
}

func restoreInfo(properties map[string]any) ([]byte, error) {
	if len(properties) == 0 {
		return nil, fmt.Errorf("IM4R properties are empty")
	}
	var encoded [][]byte
	for name, value := range properties {
		if !fourCC(name) {
			return nil, fmt.Errorf("restore property name must be four ASCII bytes")
		}
		switch v := value.(type) {
		case bool, int, int64, []byte:
		case uint64:
			value = new(big.Int).SetUint64(v)
		default:
			return nil, fmt.Errorf("unsupported restore property %s type %T", name, value)
		}
		data, err := asn1.Marshal(value)
		if err != nil {
			return nil, fmt.Errorf("encode restore property %s: %w", name, err)
		}
		field, err := wrap(asn1.ClassUniversal, asn1.TagSequence, append(ia5(name), data...))
		if err != nil {
			return nil, err
		}
		field, err = wrap(asn1.ClassPrivate, int(binary.BigEndian.Uint32([]byte(name))), field)
		if err != nil {
			return nil, err
		}
		encoded = append(encoded, field)
	}
	// A DER SET's elements are sorted by their complete encodings.
	sort.Slice(encoded, func(i, j int) bool { return bytes.Compare(encoded[i], encoded[j]) < 0 })
	set, err := wrap(asn1.ClassUniversal, asn1.TagSet, bytes.Join(encoded, nil))
	if err != nil {
		return nil, err
	}
	return wrap(asn1.ClassUniversal, asn1.TagSequence, append(ia5("IM4R"), set...))
}
