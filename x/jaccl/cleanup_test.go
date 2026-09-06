package jaccl

import (
	"errors"
	"testing"
)

func TestCloseStackReverseOrder(t *testing.T) {
	var got []string
	var stack closeStack
	for _, name := range []string{"context", "pd", "cq", "mr", "qp"} {
		name := name
		stack.add(func() error {
			got = append(got, name)
			return nil
		})
	}
	if err := stack.close(); err != nil {
		t.Fatal(err)
	}
	want := []string{"qp", "mr", "cq", "pd", "context"}
	if len(got) != len(want) {
		t.Fatalf("close order = %v, want %v", got, want)
	}
	for i := range got {
		if got[i] != want[i] {
			t.Fatalf("close order = %v, want %v", got, want)
		}
	}
}

func TestCloseStackJoinsErrors(t *testing.T) {
	first := errors.New("first")
	second := errors.New("second")
	var stack closeStack
	stack.add(func() error { return first })
	stack.add(func() error { return second })
	err := stack.close()
	if !errors.Is(err, first) || !errors.Is(err, second) {
		t.Fatalf("close error = %v, want both errors", err)
	}
	if err := stack.close(); err != nil {
		t.Fatalf("second close error = %v, want nil", err)
	}
}
