//go:build darwin && arm64

package jaccl

import (
	"errors"
	"testing"

	"github.com/tmc/apple/rdma"
)

func TestNativeDestinationValidate(t *testing.T) {
	valid := nativeDestination{
		QPN:       1,
		PSN:       1,
		LID:       1,
		GIDIndex:  1,
		GID:       rdma.IbvGID{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xff, 0xff, 10, 0, 0, 1},
		ActiveMTU: rdma.IBV_MTU_1024,
	}
	tests := []struct {
		name string
		edit func(*nativeDestination)
	}{
		{"qpn", func(d *nativeDestination) { d.QPN = 0 }},
		{"psn", func(d *nativeDestination) { d.PSN = maxPSN + 1 }},
		{"route", func(d *nativeDestination) { d.LID = 0; d.GID = rdma.IbvGID{} }},
		{"gid index", func(d *nativeDestination) { d.GIDIndex = 256 }},
		{"mtu", func(d *nativeDestination) { d.ActiveMTU = 99 }},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			destination := valid
			test.edit(&destination)
			if err := destination.validate(); !errors.Is(err, ErrProtocol) {
				t.Fatalf("validate error = %v, want ErrProtocol", err)
			}
		})
	}
	if err := valid.validate(); err != nil {
		t.Fatal(err)
	}
}

func TestNamedDevice(t *testing.T) {
	devices := []rdma.Device{{Name: "one"}, {Name: "two"}}
	got := namedDevice(devices, "two")
	if len(got) != 1 || got[0].Name != "two" {
		t.Fatalf("named device = %v, want two", got)
	}
	if got := namedDevice(devices, "missing"); got != nil {
		t.Fatalf("missing device = %v, want nil", got)
	}
}

func TestNativeDestinationRoundTrip(t *testing.T) {
	want := nativeDestination{
		QPN:       1,
		PSN:       2,
		LID:       3,
		GIDIndex:  1,
		GID:       rdma.IbvGID{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xff, 0xff, 10, 0, 0, 1},
		ActiveMTU: rdma.IBV_MTU_1024,
		LinkLayer: rdma.IBV_LINK_LAYER_THUNDERBOLT,
	}
	data, err := encodeNativeDestination(want)
	if err != nil {
		t.Fatal(err)
	}
	got, err := decodeNativeDestination(data)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Fatalf("destination = %#v, want %#v", got, want)
	}
}

func TestNativeRTRAttrUsesSelectedLocalGIDIndex(t *testing.T) {
	local := nativeDestination{
		QPN:       1,
		PSN:       2,
		LID:       3,
		GIDIndex:  2,
		GID:       rdma.IbvGID{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xff, 0xff, 192, 168, 0, 2},
		ActiveMTU: rdma.IBV_MTU_1024,
		LinkLayer: rdma.IBV_LINK_LAYER_THUNDERBOLT,
	}
	remote := local
	remote.GIDIndex = 1
	remote.GID = rdma.IbvGID{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xff, 0xff, 192, 168, 0, 1}
	attr, _, err := nativeRTRAttr(1, local, remote)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := attr.AHAttr.GRH.SGIDIndex, uint8(2); got != want {
		t.Fatalf("source GID index = %d, want selected local index %d", got, want)
	}
	if got, want := attr.AHAttr.GRH.DGID, remote.GID; got != want {
		t.Fatalf("destination GID = %x, want %x", got, want)
	}
}
