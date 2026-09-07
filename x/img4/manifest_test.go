package img4

import (
	"bytes"
	"encoding/asn1"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
	"reflect"
	"sort"
	"testing"
)

func manifestDER(class, tag int, children ...[]byte) []byte {
	b, err := wrap(class, tag, bytes.Join(children, nil))
	if err != nil {
		panic(err)
	}
	return b
}
func manifestSet(children ...[]byte) []byte {
	sort.Slice(children, func(i, j int) bool { return bytes.Compare(children[i], children[j]) < 0 })
	return manifestDER(asn1.ClassUniversal, asn1.TagSet, children...)
}
func manifestField(name string, value []byte) []byte {
	return manifestDER(asn1.ClassPrivate, int(binary.BigEndian.Uint32([]byte(name))), manifestDER(asn1.ClassUniversal, asn1.TagSequence, ia5(name), value))
}
func manifestScalar(v any) []byte {
	b, err := asn1.Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}
func manifestTicket(groups ...[]byte) []byte {
	return manifestDER(asn1.ClassUniversal, asn1.TagSequence, ia5("IM4M"), []byte{2, 1, 0}, manifestSet(manifestField("MANB", manifestSet(groups...))), []byte{4, 0}, []byte{48, 0})
}
func TestParseManifest(t *testing.T) {
	ticket := manifestTicket(
		manifestField("MANP", manifestSet(
			manifestField("ECID", manifestScalar(new(big.Int).SetUint64(1<<63|17))),
			manifestField("BORD", manifestScalar(10)),
			manifestField("BNCH", manifestScalar([]byte{1, 2, 3})),
			manifestField("CPRO", manifestScalar(true)),
			manifestField("TEST", ia5("text")),
		)),
		manifestField("krnl", manifestSet(manifestField("DGST", manifestScalar([]byte{4, 5, 6})))),
	)
	got, err := ParseManifest(ticket)
	if err != nil {
		t.Fatal(err)
	}
	want := Manifest{Properties: map[string]any{"ECID": uint64(1<<63 | 17), "BORD": uint64(10), "BNCH": []byte{1, 2, 3}, "CPRO": true, "TEST": "text"}, Images: map[string]map[string]any{"krnl": {"DGST": []byte{4, 5, 6}}}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v", got)
	}
	clear(ticket)
	if !reflect.DeepEqual(got, want) {
		t.Fatal("returned values alias ticket")
	}
}
func TestParseManifestReject(t *testing.T) {
	field := manifestField("ECID", manifestScalar(1))
	group := manifestField("MANP", manifestSet(field))
	wrongTag := manifestDER(asn1.ClassPrivate, 1, manifestDER(asn1.ClassUniversal, asn1.TagSequence, ia5("ECID"), manifestScalar(1)))
	for _, tt := range []struct {
		name  string
		input []byte
	}{
		{"empty", nil},
		{"trailing", append(manifestTicket(group), 0)},
		{"missing body", decodeHex(t, ticketHex)},
		{"missing MANP", manifestTicket()},
		{"duplicate group", manifestTicket(group, group)},
		{"duplicate property", manifestTicket(manifestField("MANP", manifestSet(field, field)))},
		{"mismatched tag", manifestTicket(manifestField("MANP", manifestSet(wrongTag)))},
		{"non-set group", manifestTicket(manifestField("MANP", manifestScalar(1)))},
		{"negative integer", manifestTicket(manifestField("MANP", manifestSet(manifestField("ECID", manifestScalar(-1)))))},
		{"oversize integer", manifestTicket(manifestField("MANP", manifestSet(manifestField("ECID", manifestScalar(new(big.Int).Lsh(big.NewInt(1), 64))))))},
		{"nested scalar", manifestTicket(manifestField("MANP", manifestSet(manifestField("ECID", manifestSet()))))},
		{"null scalar", manifestTicket(manifestField("MANP", manifestSet(manifestField("ECID", []byte{5, 0}))))},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := ParseManifest(tt.input); err == nil {
				t.Fatal("accepted invalid manifest")
			}
		})
	}
	valid := manifestTicket(group)
	t.Logf("fixture: %x", valid)
	for i := 0; i < len(valid); i++ {
		if _, err := ParseManifest(valid[:i]); err == nil {
			t.Fatalf("accepted truncation at %d", i)
		}
	}
}
func ExampleParseManifest() {
	// A synthetic manifest asserts ECID 1; its signature is empty.
	ticket, _ := hex.DecodeString("30431604494d344d0201003134ff84ea859c422d302b16044d414e423123ff84ea859c501c301a16044d414e503112ff84aa8d92440b300916044543494402010104003000")
	manifest, err := ParseManifest(ticket)
	fmt.Println(manifest.Properties["ECID"], err)
	// Output: 1 <nil>
}
func FuzzParseManifest(f *testing.F) {
	f.Add(manifestTicket(manifestField("MANP", manifestSet(manifestField("ECID", manifestScalar(1))))))
	f.Add([]byte{})
	f.Fuzz(func(t *testing.T, b []byte) { _, _ = ParseManifest(b) })
}
