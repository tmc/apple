package rdma

import "testing"

func TestClassifyCompletionStatus(t *testing.T) {
	tests := []struct {
		name   string
		status int32
		want   CompletionStatusClass
	}{
		{"success", IBV_WC_SUCCESS, CompletionSuccess},
		{"local protection", IBV_WC_LOC_PROT_ERR, CompletionProtection},
		{"local access", IBV_WC_LOC_ACCESS_ERR, CompletionProtection},
		{"remote access", IBV_WC_REM_ACCESS_ERR, CompletionProtection},
		{"other failure", 5, CompletionFailure},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := ClassifyCompletionStatus(test.status); got != test.want {
				t.Fatalf("ClassifyCompletionStatus(%d) = %q, want %q", test.status, got, test.want)
			}
		})
	}
}

func TestCompletionOpcodes(t *testing.T) {
	if IBV_WC_SEND != 0 {
		t.Fatalf("IBV_WC_SEND = %d, want 0", IBV_WC_SEND)
	}
	if IBV_WC_RECV != 128 {
		t.Fatalf("IBV_WC_RECV = %d, want 128", IBV_WC_RECV)
	}
}
