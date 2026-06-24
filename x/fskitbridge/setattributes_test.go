package fskitbridge

import (
	"syscall"
	"testing"

	"github.com/tmc/apple/fskit"
	"github.com/tmc/apple/objc"
)

// TestSetAttributesConsumesOnlyWhatWasApplied checks that the Server reports an
// attribute to FSKit as consumed only if the Volume returned it as applied.
//
// FSKit reads an attribute the file system left unconsumed as one the file
// system does not support, so consumption is how a caller learns that a chown
// did not happen. A Server that consumed each attribute as it read the request,
// before the Volume ran, reported a successful chown on a volume that cannot
// change ownership. Nothing surfaced an error, which is why this is worth a
// test rather than a comment.
func TestSetAttributesConsumesOnlyWhatWasApplied(t *testing.T) {
	s, err := newTestServer("NinePFSSetAttributes")
	if err != nil {
		t.Skipf("no Objective-C runtime for this test: %v", err)
	}
	s.reply = NewReplyBlocks()
	s.cfg.Logf = t.Logf

	// A Volume that applies mode but can neither change ownership nor set
	// file flags, so it reports back everything except those.
	vol := &setAttrVolume{applied: func(set SetAttributes) SetAttributes {
		set.UID, set.GID, set.Flags = nil, nil, nil
		return set
	}}

	// volumeFor and itemFor are map lookups, so the volume can be registered
	// under any FSVolume identifier; only the FSItem has to be a real object.
	const self = objc.ID(1)
	s.volumes.Store(self, &serverVolume{impl: vol})
	type file struct{ name string }
	item := s.newItem(vol, &file{"f"})

	request := fskit.NewFSItemSetAttributesRequest()
	request.SetMode(0o644)
	request.SetUid(501)
	request.SetGid(20)
	request.SetFlags(1)

	s.impVolumeSetAttributes(self, 0, request.GetID(), item, 0)

	if vol.got == nil {
		t.Fatal("the Server did not call SetAttributes")
	}
	// Reading the request is the half a chown needs to reach the Volume at all.
	if vol.got.Mode == nil {
		t.Error("SetAttributes got no Mode, want the requested mode")
	}
	if vol.got.UID == nil || vol.got.GID == nil {
		t.Error("SetAttributes got no UID/GID, want the requested owner")
	}
	if vol.got.Flags == nil {
		t.Error("SetAttributes got no Flags, want the requested flags")
	}

	consumed := request.ConsumedAttributes()
	if consumed&fskit.FSItemAttributeMode == 0 {
		t.Error("mode was not consumed, want it consumed: the Volume applied it")
	}
	if consumed&fskit.FSItemAttributeUID != 0 || consumed&fskit.FSItemAttributeGID != 0 {
		t.Error("owner was consumed, want it left alone: the Volume did not apply it")
	}
	if consumed&fskit.FSItemAttributeFlags != 0 {
		t.Error("flags were consumed, want them left alone: the Volume did not apply them")
	}
}

// TestSetAttributesClaimsNothingByDefault checks that a Volume returning the
// zero SetAttributes consumes nothing.
//
// This is the property that chose the signature. Reporting the applied set
// rather than clearing an in-out parameter makes forgetting to report cheap:
// an implementer who gets it wrong under-claims, so a caller is told a change
// may not have happened when it did, rather than the reverse.
func TestSetAttributesClaimsNothingByDefault(t *testing.T) {
	s, err := newTestServer("NinePFSSetAttributesDefault")
	if err != nil {
		t.Skipf("no Objective-C runtime for this test: %v", err)
	}
	s.reply = NewReplyBlocks()
	s.cfg.Logf = t.Logf

	vol := &setAttrVolume{applied: func(SetAttributes) SetAttributes {
		return SetAttributes{}
	}}
	const self = objc.ID(1)
	s.volumes.Store(self, &serverVolume{impl: vol})
	type file struct{ name string }
	item := s.newItem(vol, &file{"f"})

	request := fskit.NewFSItemSetAttributesRequest()
	request.SetMode(0o644)
	request.SetUid(501)

	s.impVolumeSetAttributes(self, 0, request.GetID(), item, 0)

	if consumed := request.ConsumedAttributes(); consumed != 0 {
		t.Errorf("consumed attributes = %d, want 0: the Volume claimed nothing", consumed)
	}
}

// A setAttrVolume is a MutableVolume that records the set-attributes request it
// was handed and reports applied(request) as the subset it applied. Every other
// operation is unreachable from the set-attributes path.
type setAttrVolume struct {
	applied func(SetAttributes) SetAttributes
	got     *SetAttributes
}

func (v *setAttrVolume) SetAttributes(item Item, set SetAttributes) (SetAttributes, error) {
	got := set
	v.got = &got
	if v.applied == nil {
		return set, nil
	}
	return v.applied(set), nil
}

func (v *setAttrVolume) Attributes(item Item) (fskit.FSItemAttributes, error) {
	return fskit.NewFSItemAttributes(), nil
}

func (v *setAttrVolume) VolumeName() string                { return "setAttrVolume" }
func (v *setAttrVolume) Root() (Item, error)               { return nil, syscall.ENOSYS }
func (v *setAttrVolume) Lookup(Item, string) (Item, error) { return nil, syscall.ENOSYS }
func (v *setAttrVolume) Reclaim(Item)                      {}
func (v *setAttrVolume) ReadDir(Item) ([]DirEntry, error)  { return nil, syscall.ENOSYS }
func (v *setAttrVolume) Read(Item, int64, []byte) (int, error) {
	return 0, syscall.ENOSYS
}

func (v *setAttrVolume) Create(Item, string, fskit.FSItemType, uint32) (Item, error) {
	return nil, syscall.ENOSYS
}
func (v *setAttrVolume) Remove(Item, string, Item) error { return syscall.ENOSYS }
func (v *setAttrVolume) Rename(Item, Item, string, Item, string, Item) error {
	return syscall.ENOSYS
}
func (v *setAttrVolume) Write(Item, int64, []byte) (int, error) { return 0, syscall.ENOSYS }

var _ MutableVolume = (*setAttrVolume)(nil)
