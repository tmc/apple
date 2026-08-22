//go:build darwin

package security_test

import (
	"fmt"

	"github.com/tmc/apple/security"
)

func ExampleSecPolicyCreateBasicX509() {
	policy := security.SecPolicyCreateBasicX509()
	typeID := security.SecPolicyGetTypeID()
	fmt.Printf("Basic X.509 policy created: %t\n", policy != 0 && typeID > 0)
	// Output:
	// Basic X.509 policy created: true
}

func ExampleAuthorizationFlags() {
	fmt.Println(security.KAuthorizationFlagDefaults)
	fmt.Println(security.KAuthorizationFlagInteractionAllowed)
	fmt.Println(security.KAuthorizationFlagExtendRights)
	// Output:
	// KAuthorizationFlagDefaults
	// KAuthorizationFlagInteractionAllowed
	// KAuthorizationFlagExtendRights
}

func ExampleCMSSignerStatus() {
	fmt.Println(security.KCMSSignerValid)
	fmt.Println(security.KCMSSignerUnsigned)
	fmt.Println(security.KCMSSignerInvalidSignature)
	// Output:
	// KCMSSignerValid
	// KCMSSignerUnsigned
	// KCMSSignerInvalidSignature
}

func ExampleCMSCertificateChainMode() {
	fmt.Println(security.KCMSCertificateChain)
	fmt.Println(security.KCMSCertificateSignerOnly)
	fmt.Println(security.KCMSCertificateNone)
	// Output:
	// KCMSCertificateChain
	// KCMSCertificateSignerOnly
	// KCMSCertificateNone
}
