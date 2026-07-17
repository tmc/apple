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
