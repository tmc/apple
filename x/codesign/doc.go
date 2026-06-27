// Package codesign ad-hoc signs the current macOS executable with a set of
// entitlements and re-execs it once, so a binary built without signing can
// still acquire entitlements it needs at runtime (notably
// com.apple.security.virtualization).
//
// It is idempotent and re-exec-guarded: if the executable already carries every
// required entitlement, or the guard environment variable is set (indicating a
// prior re-exec), [EnsureSigned] is a no-op. Call it once at startup:
//
//	err := codesign.EnsureSigned(codesign.Options{
//		Entitlements: entitlementsPlist, // plist bytes (e.g. go:embed)
//		RequireKeys:  []string{"com.apple.security.virtualization"},
//		GuardEnv:     "MYAPP_CODESIGN_DONE",
//	})
//
// This package performs ad-hoc signing (codesign -s -); it requires no signing
// identity. For signing-identity discovery and signature verification, see
// github.com/tmc/macgo/codesign, which is a separate concern.
package codesign
