package img4

import (
	"bytes"
	"encoding/asn1"
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"
)

// Synthetic DER fixtures contain no firmware or signed ticket material.
const payloadHex = "30131604494d345016046b726e6c16000403616263"
const ticketHex = "30121604494d344d020100310004037369673000"
const imageHex = "30311604494d473430131604494d345016046b726e6c16000403616263a01430121604494d344d020100310004037369673000"

func decodeHex(t *testing.T, s string) []byte {
	t.Helper()
	b, err := hex.DecodeString(s)
	if err != nil {
		t.Fatal(err)
	}
	return b
}

func TestPersonalizeGolden(t *testing.T) {
	payload, ticket := decodeHex(t, payloadHex), decodeHex(t, ticketHex)
	got, err := Personalize(payload, ticket, Options{})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got, decodeHex(t, imageHex)) {
		t.Fatalf("got %x", got)
	}
	if hex.EncodeToString(payload) != payloadHex || hex.EncodeToString(ticket) != ticketHex {
		t.Fatal("mutated input")
	}
}
func TestPersonalizeRetagAndPreserveFields(t *testing.T) {
	payload := decodeHex(t, payloadHex)
	// Opaque, well-framed PAYP property field remains inside the IM4P sequence.
	payp := decodeHex(t, "a00a30081604504159503100")
	payload = append(payload, payp...)
	payload[1] += byte(len(payp))
	original := append([]byte(nil), payload...)
	got, err := Personalize(payload, decodeHex(t, ticketHex), Options{FourCC: "rkrn"})
	if err != nil {
		t.Fatal(err)
	}
	fields, err := sequence(got, "IMG4")
	if err != nil {
		t.Fatal(err)
	}
	want := append([]byte(nil), payload...)
	copy(want[10:14], "rkrn")
	if !bytes.Equal(fields[1].FullBytes, want) {
		t.Fatalf("changed payload fields: %x", fields[1].FullBytes)
	}
	if !bytes.Equal(payload, original) {
		t.Fatal("modified caller payload")
	}
	if !bytes.Equal(fields[2].Bytes, decodeHex(t, ticketHex)) {
		t.Fatal("changed signed ticket")
	}
}
func TestRestoreProperties(t *testing.T) {
	props := map[string]any{"anid": uint64(2), "BNCN": []byte{1, 2, 3, 4, 5, 6, 7, 8}, "snid": uint64(1)<<63 | 17, "TEST": true}
	encoded, err := restoreInfo(props)
	if err != nil {
		t.Fatal(err)
	}
	fields, err := sequence(encoded, "IM4R")
	if err != nil {
		t.Fatal(err)
	}
	data := fields[1].Bytes
	var previous []byte
	seen := 0
	for len(data) > 0 {
		var property asn1.RawValue
		data, err = asn1.Unmarshal(data, &property)
		if err != nil {
			t.Fatal(err)
		}
		if property.Class != asn1.ClassPrivate || !property.IsCompound {
			t.Fatal(property)
		}
		if bytes.Compare(previous, property.FullBytes) > 0 {
			t.Fatal("SET is not sorted")
		}
		previous = property.FullBytes
		var seq struct {
			Name  string `asn1:"ia5"`
			Value asn1.RawValue
		}
		if _, err := asn1.Unmarshal(property.Bytes, &seq); err != nil {
			t.Fatal(err)
		}
		switch seq.Name {
		case "anid":
			if !bytes.Equal(seq.Value.FullBytes, []byte{2, 1, 2}) {
				t.Fatal(seq.Value)
			}
		case "snid":
			var n *big.Int
			if _, err := asn1.Unmarshal(seq.Value.FullBytes, &n); err != nil || n.Uint64() != uint64(1)<<63|17 {
				t.Fatal(n, err)
			}
		case "BNCN":
			if !bytes.Equal(seq.Value.Bytes, props["BNCN"].([]byte)) {
				t.Fatal("reversed wire nonce")
			}
		case "TEST":
			if !bytes.Equal(seq.Value.FullBytes, []byte{1, 1, 255}) {
				t.Fatal(seq.Value)
			}
		default:
			t.Fatal(seq.Name)
		}
		seen++
	}
	if seen != 4 {
		t.Fatal(seen)
	}
	// High-tag-number PRIVATE anid wraps SEQUENCE { IA5String anid, INTEGER 2 }.
	single, err := restoreInfo(map[string]any{"anid": 2})
	if err != nil {
		t.Fatal(err)
	}
	if want := "301a1604494d34523112ff868bb9d2640b30091604616e6964020102"; hex.EncodeToString(single) != want {
		t.Fatalf("IM4R got %x", single)
	}
}
func TestPersonalizeRejectsMalformed(t *testing.T) {
	payload, ticket := decodeHex(t, payloadHex), decodeHex(t, ticketHex)
	for _, tt := range []struct {
		name            string
		payload, ticket []byte
		options         Options
	}{
		{"truncated", payload[:len(payload)-1], ticket, Options{}},
		{"trailing", append(append([]byte(nil), payload...), 0), ticket, Options{}},
		{"wrong objects", ticket, payload, Options{}},
		{"bad type", payload, ticket, Options{FourCC: "longer"}},
		{"unicode", payload, ticket, Options{FourCC: "éab"}},
		{"empty restore", payload, ticket, Options{RestoreProperties: map[string]any{}}},
		{"bad property", payload, ticket, Options{RestoreProperties: map[string]any{"bad": 1}}},
		{"bad value", payload, ticket, Options{RestoreProperties: map[string]any{"anid": 1.5}}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := Personalize(tt.payload, tt.ticket, tt.options); err == nil {
				t.Fatal("accepted invalid container")
			}
		})
	}
}
func TestPersonalizeLongLengths(t *testing.T) {
	for _, size := range []int{127, 128, 255, 256, 65536} {
		payload := struct {
			Magic       string `asn1:"ia5"`
			Type        string `asn1:"ia5"`
			Description string `asn1:"ia5"`
			Data        []byte
		}{"IM4P", "krnl", "research", bytes.Repeat([]byte{7}, size)}
		encoded, err := asn1.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}
		image, err := Personalize(encoded, decodeHex(t, ticketHex), Options{RestoreProperties: map[string]any{"anid": 0}})
		if err != nil {
			t.Fatal(err)
		}
		fields, err := sequence(image, "IMG4")
		if err != nil || len(fields) != 4 {
			t.Fatal(err)
		}
		if !bytes.Equal(fields[1].FullBytes, encoded) {
			t.Fatal("changed long payload")
		}
	}
}
func ExamplePersonalize() {
	payload, _ := hex.DecodeString(payloadHex)
	ticket, _ := hex.DecodeString(ticketHex)
	image, err := Personalize(payload, ticket, Options{FourCC: "rkrn"})
	fmt.Println(len(image), err)
	// Output: 51 <nil>
}

func FuzzPersonalize(f *testing.F) {
	payload, _ := hex.DecodeString(payloadHex)
	ticket, _ := hex.DecodeString(ticketHex)
	f.Add(payload, ticket, "rkrn")
	f.Add([]byte{0x30, 0x80}, ticket, "")
	f.Fuzz(func(t *testing.T, payload, ticket []byte, fourcc string) {
		if len(payload) > 1<<20 || len(ticket) > 1<<20 || len(fourcc) > 100 {
			t.Skip()
		}
		original := append([]byte(nil), payload...)
		output, err := Personalize(payload, ticket, Options{FourCC: fourcc})
		if !bytes.Equal(original, payload) {
			t.Fatal("modified input")
		}
		if err != nil {
			return
		}
		fields, err := sequence(output, "IMG4")
		if err != nil || len(fields) != 3 {
			t.Fatal("invalid IMG4 output", err)
		}
		if !bytes.Equal(fields[2].Bytes, ticket) {
			t.Fatal("changed ticket")
		}
		if fourcc == "" && !bytes.Equal(fields[1].FullBytes, payload) {
			t.Fatal("changed untouched IM4P")
		}
	})
}
