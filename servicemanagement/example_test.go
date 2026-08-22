//go:build darwin

package servicemanagement_test

import (
	"fmt"

	"github.com/tmc/apple/servicemanagement"
)

func ExampleSMAppServiceStatus() {
	fmt.Println(servicemanagement.SMAppServiceStatusNotRegistered)
	fmt.Println(servicemanagement.SMAppServiceStatusEnabled)
	fmt.Println(servicemanagement.SMAppServiceStatusRequiresApproval)
	fmt.Println(servicemanagement.SMAppServiceStatusNotFound)
	// Output:
	// SMAppServiceStatusNotRegistered
	// SMAppServiceStatusEnabled
	// SMAppServiceStatusRequiresApproval
	// SMAppServiceStatusNotFound
}

func ExampleKSMError() {
	fmt.Println(servicemanagement.KSMErrorAlreadyRegistered)
	fmt.Println(servicemanagement.KSMErrorJobNotFound)
	fmt.Println(servicemanagement.KSMErrorInvalidPlist)
	fmt.Println(servicemanagement.KSMErrorServiceUnavailable)
	// Output:
	// KSMErrorAlreadyRegistered
	// KSMErrorJobNotFound
	// KSMErrorInvalidPlist
	// KSMErrorServiceUnavailable
}

func ExampleSMAppServiceClass_MainAppService() {
	svc := servicemanagement.GetSMAppServiceClass().MainAppService()
	status := svc.Status()
	fmt.Printf("Main app service status: %v\n", status)
	// Output:
	// Main app service status: SMAppServiceStatusNotFound
}

func ExampleSMAppServiceClass_AgentServiceWithPlistName() {
	svc := servicemanagement.GetSMAppServiceClass().AgentServiceWithPlistName("com.example.agent.plist")
	status := svc.Status()
	fmt.Printf("Agent service status: %v\n", status)
	// Output:
	// Agent service status: SMAppServiceStatusNotFound
}
