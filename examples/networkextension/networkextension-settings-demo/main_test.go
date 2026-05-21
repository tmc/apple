package main

import "testing"

func TestRunSmoke(t *testing.T) {
	if err := runSmoke(); err != nil {
		t.Fatal(err)
	}
}
