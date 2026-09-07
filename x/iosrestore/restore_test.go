package iosrestore

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReceiveRejectsInvalidFrames(t *testing.T) {
	for _, b := range [][]byte{{0, 0, 0, 0}, {1, 0, 0, 1}, {0, 0, 0, 3, 'x'}} {
		if _, err := Receive(bytes.NewReader(b)); err == nil {
			t.Fatal("accepted invalid frame")
		}
	}
}
func TestTickets(t *testing.T) {
	for _, tt := range []struct {
		name, body string
		status     int
		ok         bool
	}{
		{"success", "STATUS=0&MESSAGE=SUCCESS&REQUEST_STRING=<?xml version=\"1.0\"?><plist version=\"1.0\"><dict><key>ApImg4Ticket</key><data>AQID</data></dict></plist>", 200, true},
		{"denied", "STATUS=94&MESSAGE=not+eligible", 200, false},
		{"missing status", "MESSAGE=SUCCESS", 200, false},
		{"duplicate status", "STATUS=0&STATUS=94", 200, false},
		{"missing plist", "STATUS=0&MESSAGE=SUCCESS", 200, false},
		{"http failure", "", 503, false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "POST" {
					t.Error(r.Method)
				}
				w.WriteHeader(tt.status)
				fmt.Fprint(w, tt.body)
			}))
			defer server.Close()
			result, err := Tickets(context.Background(), server.Client(), server.URL, map[string]any{"ApECID": 123})
			if (err == nil) != tt.ok {
				t.Fatal(result, err)
			}
			if tt.ok && !bytes.Equal(result["ApImg4Ticket"].([]byte), []byte{1, 2, 3}) {
				t.Fatal(result)
			}
		})
	}
}
func ExampleSend() {
	var stream bytes.Buffer
	err := Send(&stream, map[string]any{"Request": "QueryType"})
	fmt.Println(err, stream.Len() > 4)
	// Output: <nil> true
}
func ExampleReceive() {
	var stream bytes.Buffer
	Send(&stream, map[string]any{"Request": "QueryType"})
	message, err := Receive(&stream)
	fmt.Println(message["Request"], err)
	// Output: QueryType <nil>
}
func ExampleTickets() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := Tickets(ctx, nil, "https://example.invalid", map[string]any{})
	fmt.Println(err != nil)
	// Output: true
}
