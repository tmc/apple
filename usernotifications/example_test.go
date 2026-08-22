//go:build darwin

package usernotifications_test

import (
	"fmt"

	"github.com/tmc/apple/usernotifications"
)

func ExampleUNAuthorizationOptions() {
	opts := usernotifications.UNAuthorizationOptionAlert | usernotifications.UNAuthorizationOptionSound

	fmt.Println("Option Alert:", usernotifications.UNAuthorizationOptionAlert)
	fmt.Println("Option Sound:", usernotifications.UNAuthorizationOptionSound)
	fmt.Println("Combined Mask:", uint(opts))

	// Output:
	// Option Alert: UNAuthorizationOptionAlert
	// Option Sound: UNAuthorizationOptionSound
	// Combined Mask: 6
}

func ExampleUNTimeIntervalNotificationTrigger() {
	trigger := usernotifications.NewUNTimeIntervalNotificationTriggerWithTimeIntervalRepeats(60, false)
	fmt.Printf("TimeInterval: %.0f, Repeats: %t\n", trigger.TimeInterval(), trigger.Repeats())
	// Output:
	// TimeInterval: 60, Repeats: false
}

func ExampleUNNotificationRequest() {
	content := usernotifications.NewUNMutableNotificationContent()
	trigger := usernotifications.NewUNTimeIntervalNotificationTriggerWithTimeIntervalRepeats(300, false)
	request := usernotifications.NewUNNotificationRequestWithIdentifierContentTrigger("local-reminder-42", content, trigger)

	fmt.Println("Identifier:", request.Identifier())
	fmt.Println("Trigger repeats:", request.Trigger().Repeats())
	// Output:
	// Identifier: local-reminder-42
	// Trigger repeats: false
}

func ExampleUNNotificationActionOptions() {
	opts := usernotifications.UNNotificationActionOptionForeground
	fmt.Println(opts)
	// Output:
	// UNNotificationActionOptionForeground
}
