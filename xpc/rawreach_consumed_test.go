package xpc

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

// This file is hand-maintained and deliberately NOT generated. It closes two
// holes that a green generated-output gate cannot see:
//
//   - The generator emits one rawSyms_<Entry> var per entry point, but nothing
//     forces a guard to CONSUME it. Sixteen emitted vars had no consumer at
//     all, and every existing gate stayed green, because they all measure the
//     emitted set against itself.
//   - rawReachAllowed is checked only in the reported-edge -> allowlist
//     direction. A key that matches nothing is invisible, which is how six
//     dangling keys naming functions deleted by the codec change survived.
//
// Both gates below are paired with a mutation control that runs the same
// instrument over a MUTATED COPY of the input, in memory. Neither control
// edits a .gen.go file: a hand-edit to generated source is deleted by the next
// regen, so a control that required one would be a control nobody could run.

// --- instrument -----------------------------------------------------------

// rawSymsVarUsage is what one pass over the package source found.
type rawSymsVarUsage struct {
	emitted  map[string]string // var name -> file it was declared in
	consumed map[string]int    // var name -> number of call sites naming it
	files    int
	sites    int
}

// scanRawSymsUsage parses every non-test Go file in the package and records
// which rawSyms_ vars are declared and which are named as an argument to
// requireRawSymbols. src maps a file name to overriding contents; a file whose
// name maps to an entry there is parsed from that string instead of from disk,
// which is how the mutation controls inject a change without writing to a
// generated file.
func scanRawSymsUsage(t *testing.T, src map[string]string) rawSymsVarUsage {
	t.Helper()
	u := rawSymsVarUsage{emitted: map[string]string{}, consumed: map[string]int{}}
	names, err := filepath.Glob("*.go")
	if err != nil {
		t.Fatalf("glob package sources: %v", err)
	}
	fset := token.NewFileSet()
	for _, name := range names {
		if strings.HasSuffix(name, "_test.go") {
			continue
		}
		var data any
		if override, ok := src[name]; ok {
			data = override
		} else {
			b, err := os.ReadFile(name)
			if err != nil {
				t.Fatalf("read %s: %v", name, err)
			}
			data = b
		}
		f, err := parser.ParseFile(fset, name, data, 0)
		if err != nil {
			t.Fatalf("parse %s: %v", name, err)
		}
		u.files++

		for _, d := range f.Decls {
			gd, ok := d.(*ast.GenDecl)
			if !ok || gd.Tok != token.VAR {
				continue
			}
			for _, spec := range gd.Specs {
				vs, ok := spec.(*ast.ValueSpec)
				if !ok {
					continue
				}
				for _, id := range vs.Names {
					if strings.HasPrefix(id.Name, "rawSyms_") {
						u.emitted[id.Name] = name
					}
				}
			}
		}

		ast.Inspect(f, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}
			id, ok := call.Fun.(*ast.Ident)
			if !ok || id.Name != "requireRawSymbols" {
				return true
			}
			for _, arg := range call.Args {
				a, ok := arg.(*ast.Ident)
				if ok && strings.HasPrefix(a.Name, "rawSyms_") {
					u.consumed[a.Name]++
					u.sites++
				}
			}
			return true
		})
	}
	if u.files == 0 {
		t.Fatal("scanned zero source files: a zero from this instrument would come from not looking")
	}
	if len(u.emitted) == 0 {
		t.Fatal("found zero rawSyms_ vars: the instrument is not matching the generated output")
	}
	return u
}

// rawSymsUnconsumed returns the emitted rawSyms_ vars that no requireRawSymbols
// call site names and that rawSymsIntentionallyUnconsumed does not excuse,
// plus the excuse keys that name no emitted var.
func rawSymsUnconsumed(u rawSymsVarUsage, excused map[string]string) (unguarded, danglingExcuses []string) {
	for name := range u.emitted {
		if u.consumed[name] > 0 {
			continue
		}
		if excused[name] != "" {
			continue
		}
		unguarded = append(unguarded, name)
	}
	for name := range excused {
		if _, ok := u.emitted[name]; !ok {
			danglingExcuses = append(danglingExcuses, name)
		}
	}
	sort.Strings(unguarded)
	sort.Strings(danglingExcuses)
	return unguarded, danglingExcuses
}

// --- TEST 1: every emitted rawSyms_ var is consumed or excused -------------

// rawSymsIntentionallyUnconsumed records, per emitted rawSyms_ var, why no
// guard sources its symbol set from that var. A reason here is an assertion
// that the entry point's symbols are checked SOMEWHERE ELSE, and it names
// where. It is not a licence to skip a guard.
//
// Classification vocabulary:
//
//	DELEGATES        the exported entry point does no raw call of its own; it
//	                 forwards to another exported entry point that consumes
//	                 its own rawSyms_ var. The reachable sets are identical by
//	                 construction, so a second guard would be redundant.
//	GUARDED-VIA-CALLEE
//	                 the raw calls happen in a shared unexported helper that
//	                 guards. NOTE: every helper below guards with HAND-WRITTEN
//	                 LITERAL symbol names, not with the emitted set, so its
//	                 coverage is a strict subset of the emitted set. That is a
//	                 real gap, tracked per row.
//	GENUINELY UNGUARDED
//	                 at least one symbol the entry point reaches is checked by
//	                 no requireRawSymbols call on any path to it. This is the
//	                 only category that is a defect.
//	NONREPORTING     the entry point has no error return, so a guard here
//	                 would have nowhere to report. Orthogonal to the above: a
//	                 guard with nowhere to report is a different problem from a
//	                 missing guard, and both are noted where they apply.
//
// Counts at the time of writing: 4 DELEGATES, 9 GUARDED-VIA-CALLEE,
// 3 GENUINELY UNGUARDED.
var rawSymsIntentionallyUnconsumed = map[string]string{
	// DELEGATES: pure forwarders. Each encodes/forwards and returns the
	// callee's result; the callee consumes its own rawSyms_ var.
	"rawSyms_Session_Call":           "DELEGATES: Call = encodeMessage + CallDictionary, which consumes rawSyms_Session_CallDictionary at xpc.highlevel.gen.go:1151 and :1182; the two emitted sets are identical",
	"rawSyms_Session_Notify":         "DELEGATES: Notify = encodeMessage + NotifyDictionary, which consumes rawSyms_Session_NotifyDictionary at xpc.highlevel.gen.go:1073; the two emitted sets are identical",
	"rawSyms_ReceivedMessage_Decode": "GENUINELY UNGUARDED (inherited): Decode = decodeMessage, whose only raw reach is (ReceivedMessage).Dictionary; the two emitted sets are identical, so it inherits Dictionary's unguarded getters. Decode DOES have an error channel, so the defect is not the missing return path here but Dictionary swallowing rawObjectToDictionary's error at xpc.highlevel.gen.go:1240",
	"rawSyms_DialXPCService":         "DELEGATES: DialXPCService = dialSession(\"xpc\", ...); dialSession guards, but with two literal names only (see rawSyms_DialMachService)",
	"rawSyms_NewServiceListener":     "DELEGATES: NewServiceListener = newListener(service, ..., true); newListener guards, but with seven literal names only (see rawSyms_NewAnonymousListener)",

	// GUARDED-VIA-CALLEE, with a measured literal-vs-emitted shortfall.
	"rawSyms_DialMachService":                      "GUARDED-VIA-CALLEE: dialSession guards at xpc.highlevel.gen.go:864 with the literals xpc_session_create_{xpc,mach}_service; the emitted set has 7 symbols, so xpc_release, xpc_rich_error_can_retry, xpc_rich_error_copy_description, xpc_session_activate and xpc_session_set_peer_requirement are reached unguarded on this path (the last two are re-guarded by SetPeerRequirement/Activate, which do consume their vars)",
	"rawSyms_NewAnonymousListener":                 "GUARDED-VIA-CALLEE: newListener guards at xpc.highlevel.gen.go:700 with 7 literals plus a conditional xpc_listener_set_peer_requirement; the emitted set has 29 symbols, so the whole codec surface (xpc_*_create, xpc_dictionary_set_*, xpc_release, xpc_rich_error_*) is reached unguarded on this path",
	"rawSyms_NewEntitlementExistsRequirement":      "GUARDED-VIA-CALLEE: newRequirement guards at xpc.highlevel.gen.go:347 with the one literal its caller passes; here that literal equals the create symbol, and the remaining 2 emitted symbols are the rich-error pair",
	"rawSyms_NewPlatformBinaryRequirement":         "GUARDED-VIA-CALLEE: as above; 1 of 3 emitted symbols guarded, remainder is the rich-error pair",
	"rawSyms_NewPlatformBinarySignedAsRequirement": "GUARDED-VIA-CALLEE: as above; 1 of 3 emitted symbols guarded, remainder is the rich-error pair",
	"rawSyms_NewSameTeamRequirement":               "GUARDED-VIA-CALLEE: as above; 1 of 3 emitted symbols guarded, remainder is the rich-error pair",
	"rawSyms_NewSameTeamSignedAsRequirement":       "GUARDED-VIA-CALLEE: as above; 1 of 3 emitted symbols guarded, remainder is the rich-error pair",
	"rawSyms_NewEntitlementMatchesRequirement":     "GUARDED-VIA-CALLEE: newRequirement guards 1 of 7 emitted symbols; entitlementValueToRawObject runs BEFORE the guard and calls xpc_{bool,int64,string}_create unguarded",
	"rawSyms_NewLightweightCodeRequirement":        "GUARDED-VIA-CALLEE: newRequirement guards 1 of 21 emitted symbols; dictionaryToRawObject runs BEFORE the guard and calls the whole codec surface unguarded",

	// GENUINELY UNGUARDED. These three are defects, not redundancies. They are
	// recorded here so the gate stays green on a known state, not excused away.
	"rawSyms_Listener_Cancel":            "GENUINELY UNGUARDED on one construction path, and NONREPORTING: Cancel returns nothing, so a guard would have nowhere to report. xpc_listener_cancel is guarded by literal in newListener at xpc.highlevel.gen.go:700, but ListenerFromHandle (xpc.highlevel.gen.go:443) sets l.raw without passing through newListener, so on that path the call is both unguarded and unreportable",
	"rawSyms_ReceivedMessage_Dictionary": "GENUINELY UNGUARDED for 9 of 11 emitted symbols, and NONREPORTING: rawObjectToDictionary guards only the literal xpc_dictionary_apply (xpc.highlevel.gen.go:1607) and rawArrayToSlice only xpc_array_apply (:1625); rawObjectToValue (:1639) and copyRawData (:1670) call xpc_get_type, xpc_{bool,int64,uint64,double}_get_value, xpc_string_get_string_ptr, xpc_data_get_{length,bytes_ptr} and xpc_copy_description with no guard on any path. Dictionary returns Dictionary with no error, so it also discards the two guard errors that DO exist",
}

func TestEveryEmittedRawSymsVarIsConsumed(t *testing.T) {
	u := scanRawSymsUsage(t, nil)
	unguarded, dangling := rawSymsUnconsumed(u, rawSymsIntentionallyUnconsumed)

	t.Logf("rawSyms_ vars: %d emitted, %d consumed by %d call site(s), %d excused, %d unaccounted",
		len(u.emitted), len(u.consumed), u.sites, len(rawSymsIntentionallyUnconsumed), len(unguarded))

	for _, name := range unguarded {
		t.Errorf("%s is emitted by the generator but no requireRawSymbols call site names it, and rawSymsIntentionallyUnconsumed gives no reason: the entry point's availability guard is not derived from the emitted set", name)
	}
	// Liveness, in the other direction: an excuse for a var the generator no
	// longer emits is a stale claim about code that is gone.
	for _, name := range dangling {
		t.Errorf("rawSymsIntentionallyUnconsumed excuses %s, which the generator does not emit: delete the stale entry", name)
	}
}

func TestEveryEmittedRawSymsVarIsConsumedIsSensitive(t *testing.T) {
	base := scanRawSymsUsage(t, nil)
	if len(base.consumed) == 0 {
		t.Fatal("no consumed vars to remove: the control would pass by never looking")
	}
	// Pick a deterministic consumed var and delete its call sites from an
	// in-memory copy of the file that consumes them.
	var victim string
	for name := range base.consumed {
		if victim == "" || name < victim {
			victim = name
		}
	}
	const guardFile = "xpc.highlevel.gen.go"
	b, err := os.ReadFile(guardFile)
	if err != nil {
		t.Fatalf("read %s: %v", guardFile, err)
	}
	mutated := strings.ReplaceAll(string(b),
		"requireRawSymbols("+victim+"...)",
		"requireRawSymbols(\"xpc_release\")")
	if mutated == string(b) {
		t.Fatalf("mutation control could not remove the consumer of %s from %s: the instrument and the source disagree about where guards live", victim, guardFile)
	}

	u := scanRawSymsUsage(t, map[string]string{guardFile: mutated})
	if u.consumed[victim] != 0 {
		t.Fatalf("mutation control failed to take effect: %s still has %d consumer(s)", victim, u.consumed[victim])
	}
	unguarded, _ := rawSymsUnconsumed(u, rawSymsIntentionallyUnconsumed)
	found := false
	for _, name := range unguarded {
		if name == victim {
			found = true
		}
	}
	if !found {
		t.Fatalf("mutation control failed: removed the only consumer of %s, but the gate reported it as accounted for (examined %d emitted vars, %d reported unguarded)", victim, len(u.emitted), len(unguarded))
	}
	t.Logf("mutation control: examined %d emitted vars; removing the consumer of %s made the gate report it (%d unguarded under mutation, %d without)",
		len(u.emitted), victim, len(unguarded), 0)
}

// --- TEST 2: allowlist liveness -------------------------------------------

// rawReachAllowlistDeadKeys returns the rawReachAllowed keys that match no
// reported edge. TestRawReachUnresolvedIsAllowlisted checks only the other
// direction (every edge has a key), so without this a key naming a deleted
// function survives forever.
func rawReachAllowlistDeadKeys(allowed map[string]string, edges []rawReachEdge) []string {
	live := map[string]bool{}
	for _, e := range edges {
		live[e.Entry+"|"+e.Expr] = true
	}
	var dead []string
	for key := range allowed {
		if !live[key] {
			dead = append(dead, key)
		}
	}
	sort.Strings(dead)
	return dead
}

func TestRawReachAllowlistHasNoDeadKeys(t *testing.T) {
	if len(rawReachAllowed) == 0 {
		t.Fatal("rawReachAllowed is empty: a green run here would prove nothing")
	}
	if len(rawReachUnresolved) == 0 {
		t.Fatal("the generator reported no unfollowable edges: every allowlist key would look dead, so this gate cannot run")
	}
	dead := rawReachAllowlistDeadKeys(rawReachAllowed, rawReachUnresolved)
	t.Logf("allowlist liveness: %d key(s), %d reported edge(s), %d dead key(s)",
		len(rawReachAllowed), len(rawReachUnresolved), len(dead))
	for _, key := range dead {
		t.Errorf("rawReachAllowed key %q matches no edge the generator reports: it justifies a call edge that no longer exists (renamed, deleted, or a typo), so it is not protecting anything -- delete it", key)
	}
}

func TestRawReachAllowlistHasNoDeadKeysIsSensitive(t *testing.T) {
	mutated := make(map[string]string, len(rawReachAllowed)+1)
	for k, v := range rawReachAllowed {
		mutated[k] = v
	}
	const bogus = "functionDeletedByTheCodecChange|somethingRaw(...)"
	mutated[bogus] = "bogus key injected by the mutation control"

	dead := rawReachAllowlistDeadKeys(mutated, rawReachUnresolved)
	found := false
	for _, key := range dead {
		if key == bogus {
			found = true
		}
	}
	if !found {
		t.Fatalf("mutation control failed: a bogus allowlist key was reported as live (examined %d keys against %d edges, %d reported dead)",
			len(mutated), len(rawReachUnresolved), len(dead))
	}
	t.Logf("mutation control: examined %d keys against %d edges; bogus key correctly reported dead", len(mutated), len(rawReachUnresolved))
}
