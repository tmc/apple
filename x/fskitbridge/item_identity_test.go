package fskitbridge

import "testing"

// TestNewItemIsStable checks that the same Item always maps to the same FSItem.
//
// FSKit correlates items by object identity: overItem names an object "as
// discovered in a prior lookup", and reclaimItem: retires one specific object.
// A Server that allocated a fresh FSItem per lookup left FSKit unable to match
// a looked-up item against the one an operation names, and it responded by
// declining to dispatch those operations at all -- unlink reported success and
// removed nothing, a rename onto an existing name failed with ENOENT, and
// neither reached the file system. Nothing surfaced an error, which is why this
// is worth a test rather than a comment.
func TestNewItemIsStable(t *testing.T) {
	s, err := newTestServer("NinePFSIdentity")
	if err != nil {
		t.Skipf("no Objective-C runtime for this test: %v", err)
	}

	type file struct{ name string }
	a, b := &file{"a"}, &file{"b"}

	first := s.newItem(nil, a)
	if first == 0 {
		t.Fatal("newItem returned a nil FSItem")
	}
	if again := s.newItem(nil, a); again != first {
		t.Fatalf("newItem(a) = %v on the second call, want %v", again, first)
	}
	if other := s.newItem(nil, b); other == first {
		t.Fatal("newItem(b) returned the FSItem belonging to a")
	}

	// Reclaim retires the object, so the next lookup of the same file must get
	// a new one rather than the reclaimed one.
	s.objs.Delete(a)
	s.items.Delete(first)
	if after := s.newItem(nil, a); after == first {
		t.Fatal("newItem(a) returned the reclaimed FSItem")
	}
}

// newTestServer returns a Server with a registered class set and no file
// system implementation behind it.
//
// Each caller passes its own prefix: RegisterClasses reuses an item or volume
// class that already exists but always registers the file system class afresh,
// so two tests sharing a prefix would collide on the second one and report the
// collision as a missing Objective-C runtime.
func newTestServer(prefix string) (*Server, error) {
	s := &Server{}
	classes, err := RegisterClasses(ClassConfig{
		FileSystemName: prefix + "TestFileSystem",
		VolumeName:     prefix + "TestVolume",
		ItemName:       prefix + "TestItem",
	})
	if err != nil {
		return nil, err
	}
	s.classes = classes
	return s, nil
}
