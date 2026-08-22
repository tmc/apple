// Command authcheck reports which local authentication policies this machine
// can evaluate, using the LocalAuthentication framework.
//
// By default it only queries capabilities and never prompts the user. With the
// -evaluate flag it additionally evaluates the device owner authentication
// policy, which presents the system authentication dialog.
//
// Usage: authcheck [-evaluate] [-reason text]
package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/tmc/apple/localauthentication"
)

var policies = []struct {
	name   string
	policy localauthentication.LAPolicy
}{
	{"biometrics", localauthentication.LAPolicyDeviceOwnerAuthenticationWithBiometrics},
	{"biometrics or companion", localauthentication.LAPolicyDeviceOwnerAuthenticationWithBiometricsOrCompanion},
	{"companion", localauthentication.LAPolicyDeviceOwnerAuthenticationWithCompanion},
	{"biometrics or password", localauthentication.LAPolicyDeviceOwnerAuthentication},
}

func main() {
	evaluate := flag.Bool("evaluate", false, "evaluate the device owner authentication policy, prompting the user")
	reason := flag.String("reason", "Demonstrate local authentication", "reason shown in the authentication dialog")
	flag.Parse()

	lac := localauthentication.NewLAContext()
	defer lac.Invalidate()

	fmt.Printf("biometry type: %v\n", lac.BiometryType())
	for _, p := range policies {
		ok, err := lac.CanEvaluatePolicyError(p.policy)
		switch {
		case err != nil:
			fmt.Printf("%-24s unavailable: %v\n", p.name+":", err)
		case ok:
			fmt.Printf("%-24s available\n", p.name+":")
		default:
			fmt.Printf("%-24s unavailable\n", p.name+":")
		}
	}

	if !*evaluate {
		return
	}

	policy := localauthentication.LAPolicyDeviceOwnerAuthentication
	ok, err := lac.EvaluatePolicyLocalizedReasonReplySync(context.Background(), policy, *reason)
	if err != nil {
		fmt.Fprintf(os.Stderr, "evaluate: %v\n", err)
		os.Exit(1)
	}
	if !ok {
		fmt.Fprintf(os.Stderr, "authentication failed\n")
		os.Exit(1)
	}
	fmt.Println("authenticated")
}
