package codesign

import (
	"strings"
	"testing"
)

func TestEnsureSignedValidation(t *testing.T) {
	tests := []struct {
		name string
		opts Options
		want string // substring the error must contain
	}{
		{
			name: "missing guard env",
			opts: Options{RequireKeys: []string{"k"}, Entitlements: []byte("x")},
			want: "GuardEnv",
		},
		{
			name: "missing require keys",
			opts: Options{GuardEnv: "G", Entitlements: []byte("x")},
			want: "RequireKeys",
		},
		{
			name: "missing entitlements",
			opts: Options{GuardEnv: "G", RequireKeys: []string{"k"}},
			want: "Entitlements",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := EnsureSigned(tt.opts)
			if err == nil || !strings.Contains(err.Error(), tt.want) {
				t.Fatalf("EnsureSigned() error = %v, want substring %q", err, tt.want)
			}
		})
	}
}

func TestEnsureSignedGuardShortCircuits(t *testing.T) {
	t.Setenv("CODESIGN_TEST_GUARD", "1")
	called := false
	opts := Options{
		Entitlements: []byte("x"),
		RequireKeys:  []string{"com.apple.security.virtualization"},
		GuardEnv:     "CODESIGN_TEST_GUARD",
		runner: func(args ...string) ([]byte, error) {
			called = true
			return nil, nil
		},
	}
	if err := EnsureSigned(opts); err != nil {
		t.Fatalf("EnsureSigned() = %v, want nil when guard set", err)
	}
	if called {
		t.Error("runner was invoked despite the guard being set")
	}
}

func TestHasAllEntitlements(t *testing.T) {
	// The codesign -d output a signed virtualization binary would produce.
	const signed = `<plist><dict>
		<key>com.apple.security.virtualization</key><true/>
		<key>com.apple.security.network.client</key><true/>
	</dict></plist>`

	tests := []struct {
		name string
		out  string
		err  error
		keys []string
		want bool
	}{
		{
			name: "single key present",
			out:  signed,
			keys: []string{"com.apple.security.virtualization"},
			want: true,
		},
		{
			name: "all keys present",
			out:  signed,
			keys: []string{"com.apple.security.virtualization", "com.apple.security.network.client"},
			want: true,
		},
		{
			// The cove bug this package fixes: virtualization present but a
			// required network key missing must NOT count as already signed.
			name: "one of several keys missing",
			out:  signed,
			keys: []string{"com.apple.security.virtualization", "com.apple.security.network.server"},
			want: false,
		},
		{
			name: "codesign error means unsigned",
			out:  "",
			err:  errStub,
			keys: []string{"com.apple.security.virtualization"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			run := func(args ...string) ([]byte, error) {
				return []byte(tt.out), tt.err
			}
			if got := hasAllEntitlements(run, "/bin/exe", tt.keys); got != tt.want {
				t.Errorf("hasAllEntitlements() = %v, want %v", got, tt.want)
			}
		})
	}
}

var errStub = stubErr("codesign: no signature")

type stubErr string

func (e stubErr) Error() string { return string(e) }
