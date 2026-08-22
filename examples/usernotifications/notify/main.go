// Command notify requests notification authorization and schedules a local
// notification using the UserNotifications framework.
//
// UserNotifications requires a bundled application: an executable run from a
// plain directory has no main bundle identifier, and asking for the current
// notification center in that state raises an Objective-C exception that
// terminates the process. notify therefore checks the main bundle first and
// reports what is missing instead of crashing.
//
// Usage:
//
//	notify [-title t] [-body b] [-delay d]
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/usernotifications"
)

func main() {
	title := flag.String("title", "Hello from Go", "notification title")
	body := flag.String("body", "Scheduled by the notify example.", "notification body")
	delay := flag.Duration("delay", 5*time.Second, "delay before delivery")
	flag.Parse()

	if err := run(*title, *body, *delay); err != nil {
		fmt.Fprintf(os.Stderr, "notify: %v\n", err)
		os.Exit(1)
	}
}

func run(title, body string, delay time.Duration) error {
	id := foundation.GetBundleClass().MainBundle().BundleIdentifier()
	if id == "" {
		fmt.Println("no main bundle identifier: UserNotifications is unavailable.")
		fmt.Println("run this program from inside a signed .app bundle to post notifications.")
		return nil
	}
	fmt.Println("bundle:", id)

	if delay < time.Second {
		return fmt.Errorf("delay must be at least 1s")
	}

	center := usernotifications.GetUNUserNotificationCenterClass().CurrentNotificationCenter()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	opts := usernotifications.UNAuthorizationOptionAlert | usernotifications.UNAuthorizationOptionSound
	granted, err := center.RequestAuthorizationWithOptions(ctx, opts)
	if err != nil {
		return fmt.Errorf("request authorization: %w", err)
	}
	if !granted {
		fmt.Println("authorization denied; nothing scheduled")
		return nil
	}

	settings, err := center.GetNotificationSettings(ctx)
	if err != nil {
		return fmt.Errorf("get settings: %w", err)
	}
	if settings != nil {
		fmt.Println("authorization status:", settings.AuthorizationStatus())
		fmt.Println("alert setting:", settings.AlertSetting())
	}

	content := usernotifications.NewUNMutableNotificationContent()
	// The generated bindings expose only getters on UNNotificationContent, so
	// set the editable fields through their Objective-C setters.
	objc.Send[struct{}](content.ID, objc.Sel("setTitle:"), objc.String(title))
	objc.Send[struct{}](content.ID, objc.Sel("setBody:"), objc.String(body))

	trigger := usernotifications.NewUNTimeIntervalNotificationTriggerWithTimeIntervalRepeats(delay.Seconds(), false)
	ident := fmt.Sprintf("notify-%d", time.Now().UnixNano())
	request := usernotifications.NewUNNotificationRequestWithIdentifierContentTrigger(ident, content, trigger)

	if err := center.AddNotificationRequest(ctx, request); err != nil {
		return fmt.Errorf("add request: %w", err)
	}
	fmt.Printf("scheduled %s in %s\n", ident, delay)

	pending, err := center.GetPendingNotificationRequests(ctx)
	if err != nil {
		return fmt.Errorf("get pending requests: %w", err)
	}
	for _, p := range pending {
		fmt.Printf("pending: %s %q\n", p.Identifier(), p.Content().Title())
	}
	return nil
}
