//go:build darwin

package endpointsecurity_test

import (
	"fmt"

	"github.com/tmc/apple/endpointsecurity"
)

func ExampleEsAuthResult() {
	fmt.Println(endpointsecurity.EsAuthResultAllow)
	fmt.Println(endpointsecurity.EsAuthResultDeny)

	// Output:
	// EsAuthResultAllow
	// EsAuthResultDeny
}

func ExampleEsActionType() {
	fmt.Println(endpointsecurity.EsActionTypeAuth)
	fmt.Println(endpointsecurity.EsActionTypeNotify)

	// Output:
	// EsActionTypeAuth
	// EsActionTypeNotify
}
