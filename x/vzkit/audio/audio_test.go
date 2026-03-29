package audio

import "testing"

func TestCreateDeviceRequiresStream(t *testing.T) {
	_, err := CreateDevice(Config{})
	if err == nil {
		t.Fatal("CreateDevice(Config{}) error = nil")
	}
}
