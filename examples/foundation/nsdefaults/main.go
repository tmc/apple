// Command nsdefaults reads and writes NSUserDefaults values.
//
// Usage:
//
//	nsdefaults read <key>
//	nsdefaults write <key> <value>
//	nsdefaults write -bool <key> <true|false>
//	nsdefaults write -int <key> <number>
//	nsdefaults delete <key>
//	nsdefaults -j read <key>
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/tmc/apple/foundation"
)

func usage() {
	fmt.Fprintln(os.Stderr, "usage: nsdefaults [-j] <read|write|delete> [-bool|-int] <key> [value]")
	os.Exit(2)
}

func main() {
	// Parse args manually since flags and subcommands are interleaved.
	var jsonOut bool
	var boolType, intType bool
	var positional []string

	for _, arg := range os.Args[1:] {
		switch arg {
		case "-j":
			jsonOut = true
		case "-bool":
			boolType = true
		case "-int":
			intType = true
		default:
			positional = append(positional, arg)
		}
	}

	if len(positional) < 1 {
		usage()
	}

	cmd := positional[0]
	args := positional[1:]
	ud := foundation.GetUserDefaultsClass().StandardUserDefaults()

	switch cmd {
	case "read":
		if len(args) < 1 {
			fmt.Fprintln(os.Stderr, "usage: nsdefaults read <key>")
			os.Exit(2)
		}
		key := args[0]
		obj := ud.ObjectForKey(key)
		if obj.GetID() == 0 {
			fmt.Fprintf(os.Stderr, "key not found: %s\n", key)
			os.Exit(1)
		}
		val := readValue(ud, key)
		if jsonOut {
			printJSON(map[string]any{"key": key, "value": val})
		} else {
			fmt.Println(val)
		}

	case "write":
		if boolType {
			if len(args) < 2 {
				fmt.Fprintln(os.Stderr, "usage: nsdefaults write -bool <key> <true|false>")
				os.Exit(2)
			}
			b, err := strconv.ParseBool(args[1])
			if err != nil {
				fmt.Fprintf(os.Stderr, "invalid bool: %s\n", args[1])
				os.Exit(2)
			}
			ud.SetBoolForKey(b, args[0])
		} else if intType {
			if len(args) < 2 {
				fmt.Fprintln(os.Stderr, "usage: nsdefaults write -int <key> <number>")
				os.Exit(2)
			}
			n, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Fprintf(os.Stderr, "invalid int: %s\n", args[1])
				os.Exit(2)
			}
			ud.SetIntegerForKey(n, args[0])
		} else {
			if len(args) < 2 {
				fmt.Fprintln(os.Stderr, "usage: nsdefaults write <key> <value>")
				os.Exit(2)
			}
			nsStr := foundation.GetNSStringClass().Alloc().InitWithString(args[1])
			ud.SetObjectForKey(nsStr, args[0])
		}
		ud.Synchronize()

	case "delete":
		if len(args) < 1 {
			fmt.Fprintln(os.Stderr, "usage: nsdefaults delete <key>")
			os.Exit(2)
		}
		key := args[0]
		ud.RemoveObjectForKey(key)
		ud.Synchronize()
		if jsonOut {
			printJSON(map[string]any{"deleted": key})
		} else {
			fmt.Printf("deleted: %s\n", key)
		}

	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", cmd)
		os.Exit(2)
	}
}

func readValue(ud foundation.NSUserDefaults, key string) any {
	if s := ud.StringForKey(key); s != "" {
		return s
	}
	if n := ud.IntegerForKey(key); n != 0 {
		return n
	}
	return ud.BoolForKey(key)
}

func printJSON(v any) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(v); err != nil {
		fmt.Fprintf(os.Stderr, "encoding json: %v\n", err)
		os.Exit(2)
	}
}
