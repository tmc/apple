// Command sysinfo prints system information and optionally lists running applications.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/foundation"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	jsonOut := flag.Bool("j", false, "JSON output")
	apps := flag.Bool("apps", false, "list running applications")
	flag.Parse()

	if *apps {
		listApps(*jsonOut)
		return
	}
	showInfo(*jsonOut)
}

func showInfo(jsonOut bool) {
	pi := foundation.GetProcessInfoClass().ProcessInfo()

	if jsonOut {
		v := map[string]any{
			"hostname":     pi.HostName(),
			"os_version":   pi.OperatingSystemVersionString(),
			"processors":   pi.ProcessorCount(),
			"memory_bytes": pi.PhysicalMemory(),
			"pid":          pi.ProcessIdentifier(),
			"process_name": pi.ProcessName(),
		}
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		if err := enc.Encode(v); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(2)
		}
		return
	}

	fmt.Printf("hostname=%s\n", pi.HostName())
	fmt.Printf("os=%s\n", pi.OperatingSystemVersionString())
	fmt.Printf("processors=%d\n", pi.ProcessorCount())
	fmt.Printf("memory=%d\n", pi.PhysicalMemory())
	fmt.Printf("pid=%d\n", pi.ProcessIdentifier())
	fmt.Printf("name=%s\n", pi.ProcessName())
}

func listApps(jsonOut bool) {
	ws := appkit.GetNSWorkspaceClass().SharedWorkspace()
	running := ws.RunningApplications()

	if jsonOut {
		var result []map[string]any
		for _, app := range running {
			result = append(result, map[string]any{
				"name":      app.LocalizedName(),
				"pid":       app.ProcessIdentifier(),
				"bundle_id": app.BundleIdentifier(),
			})
		}
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		if err := enc.Encode(result); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(2)
		}
		return
	}

	for _, app := range running {
		fmt.Printf("%s\tpid=%d\tbundle=%s\n", app.LocalizedName(), app.ProcessIdentifier(), app.BundleIdentifier())
	}
}
