export const meta = {
  name: 'pureentry-research',
  description: 'Drive the 9pfs no-Swift FSKit entrypoint experiment to a definitive, verified conclusion',
  phases: [
    { title: 'Understand', detail: 'map EF/FSKit load-time requirements and the existing pipeline' },
    { title: 'Reproduce', detail: 'build the current pureentry .goentry and find the exact stall' },
    { title: 'Synthesize', detail: 'attempt to produce the swift metadata surface without swiftc' },
    { title: 'Verify', detail: 'adversarially test every blocker/success claim' },
    { title: 'Synthesize report', detail: 'reconcile findings into a verdict' },
  ],
}

const ROOT = '/Volumes/tmc/go/src/github.com/tmc/apple-worktrees-fskit'
const RES = ROOT + '/examples/fskit/9pfs-research'
const NINE = ROOT + '/examples/fskit/9pfs'

const REQ_SCHEMA = {
  type: 'object',
  required: ['summary', 'requirements', 'pipelineSteps', 'openQuestions'],
  properties: {
    summary: { type: 'string' },
    requirements: {
      type: 'array',
      description: 'Things ExtensionFoundation/FSKit demand of the extension binary at load time',
      items: {
        type: 'object',
        required: ['what', 'evidence', 'satisfiableWithoutSwiftc'],
        properties: {
          what: { type: 'string' },
          evidence: { type: 'string', description: 'file:line or symbol/section that proves it' },
          satisfiableWithoutSwiftc: { type: 'string', enum: ['yes', 'no', 'unknown'] },
          why: { type: 'string' },
        },
      },
    },
    pipelineSteps: {
      type: 'array',
      description: 'Ordered steps the original build-appex.sh pureentry branch performs',
      items: { type: 'string' },
    },
    openQuestions: { type: 'array', items: { type: 'string' } },
  },
}

const BUILD_SCHEMA = {
  type: 'object',
  required: ['ran', 'reachedStep', 'stalledAt', 'stallReason', 'commands', 'rawTail'],
  properties: {
    ran: { type: 'boolean', description: 'did you actually execute builds (not just read)' },
    reachedStep: { type: 'string', description: 'furthest pipeline step that succeeded' },
    stalledAt: { type: 'string', description: 'first step that failed, or "none — built clean"' },
    stallReason: { type: 'string', description: 'the precise error/blocker, with the actual message' },
    commands: { type: 'array', items: { type: 'string' }, description: 'exact commands run' },
    artifacts: { type: 'array', items: { type: 'string' }, description: 'paths produced (binary, syso, etc.)' },
    rawTail: { type: 'string', description: 'last ~25 lines of the most relevant build/link output' },
  },
}

const SYNTH_SCHEMA = {
  type: 'object',
  required: ['attempted', 'approach', 'opaqueBytes', 'verdict', 'evidence'],
  properties: {
    attempted: { type: 'boolean' },
    approach: { type: 'string', description: 'how you tried to produce the metadata without swiftc' },
    opaqueBytes: {
      type: 'array',
      description: 'metadata sections whose CONTENT cannot be derived from the declarative manifest alone',
      items: {
        type: 'object',
        required: ['section', 'whatItContains', 'derivableWithoutSwiftc'],
        properties: {
          section: { type: 'string' },
          whatItContains: { type: 'string' },
          derivableWithoutSwiftc: { type: 'string', enum: ['yes', 'no', 'partial'] },
        },
      },
    },
    verdict: { type: 'string', enum: ['synthesis-feasible', 'synthesis-blocked', 'partial'] },
    evidence: { type: 'string' },
  },
}

const VERDICT_SCHEMA = {
  type: 'object',
  required: ['claim', 'holds', 'confidence', 'reasoning'],
  properties: {
    claim: { type: 'string' },
    holds: { type: 'boolean', description: 'true if the claim survives scrutiny' },
    confidence: { type: 'string', enum: ['high', 'medium', 'low'] },
    reasoning: { type: 'string' },
    counterevidence: { type: 'string' },
  },
}

// ---- Phase 1: Understand the requirement boundary (parallel readers) ----
phase('Understand')
const understanding = await parallel([
  () => agent(
    `Read these files and determine EXACTLY what Apple's ExtensionFoundation and FSKit require of an ` +
    `FSKit app-extension binary at load time (the metadata/conformance/entrypoint surface that lets ` +
    `pluginkit/NSExtension find and instantiate the extension principal class). Focus on what is ` +
    `consumed from the Mach-O image (e.g. __swift5_proto / __swift5_protos / __swift5_types / ` +
    `conformance records / the NSExtensionPrincipalClass plist key) versus what is just runtime linkage.\n\n` +
    `Read:\n- ${RES}/RESEARCH.md\n- ${RES}/pureentry/NinePFSExtension.swift\n- ${NINE}/appex/NinePFSExtension.swift (if present)\n` +
    `- ${NINE}/Info.plist and any *.plist under ${NINE}\n- ${RES}/pureentry_darwin.go\n- ${RES}/nsext_arm64.s\n- ${RES}/dynimport_pureentry_darwin.go\n\n` +
    `Use Bash freely: otool -l, nm, jq on the manifest, and 'man' pages. Report each requirement with ` +
    `concrete evidence (a symbol, a section, a plist key, a file:line). For each, judge whether it can ` +
    `be satisfied WITHOUT invoking swiftc/clang (i.e. by emitting bytes directly).`,
    { label: 'understand:requirements', phase: 'Understand', schema: REQ_SCHEMA, agentType: 'Explore' }
  ),
  () => agent(
    `Reconstruct the FULL pureentry build pipeline as it was actually run. The original build script is ` +
    `at /tmp/orig-build-appex.sh (the 'pureentry == yes' branch, lines ~155-335). Read it and the tools ` +
    `it drives under ${RES}/pureentry/. Produce the ordered list of pipeline steps (swiftc -> metadata ` +
    `object -> dump/emit round-trip -> rename _main -> Go external link -> patch LC_MAIN -> verify). ` +
    `Note for each step which tool runs and what artifact it produces. Also enumerate, using ` +
    `'python3 -c' or jq, the sections in ${RES}/pureentry/swiftmeta_manifest.json: segment, name, ` +
    `size, #relocations, and whether data is inline or in a data_file. This tells us how much of the ` +
    `metadata is declaratively described vs opaque.`,
    { label: 'understand:pipeline', phase: 'Understand', schema: REQ_SCHEMA, agentType: 'Explore' }
  ),
]).then(r => r.filter(Boolean))

// ---- Phase 2: Reproduce the build empirically, find the stall ----
phase('Reproduce')
const reqText = JSON.stringify(understanding, null, 2).slice(0, 6000)
const build = await agent(
  `You are reproducing the 9pfs pureentry build on THIS machine (macOS 26.5 arm64, Swift 6.3.2, Go 1.26.3, ` +
  `FSKit + ExtensionFoundation present). Goal: build a Go-entrypoint FSKit extension binary and find the ` +
  `EXACT step where it succeeds or stalls. Do NOT install or mount in this phase.\n\n` +
  `Context from phase 1:\n${reqText}\n\n` +
  `The original build pipeline is /tmp/orig-build-appex.sh (pureentry branch). The pureentry source ` +
  `fragments live in ${RES} (pureentry_darwin.go, nsext_arm64.s, dynimport_pureentry_darwin.go) and the ` +
  `core 9pfs main package is in ${NINE}. The tools are go-run programs in ${RES}/pureentry/.\n\n` +
  `Work in a scratch dir under /tmp (e.g. /tmp/pureentry-build). You will need to overlay the ${RES} ` +
  `pureentry fragments onto a copy of the ${NINE} source (the fragments are build-tagged 'pureentry' and ` +
  `are fragments of the 9pfs main package). Drive the pipeline manually following /tmp/orig-build-appex.sh:\n` +
  `  1. swiftc-compile NinePFSExtension.swift to get the metadata object (swiftmeta.o)\n` +
  `  2. round-trip it through dump_swift_metadata.go / emit_macho_object.go to prove the emitter reproduces it\n` +
  `  3. rename _main -> _swift_unused_main with rename_macho_symbol.go\n` +
  `  4. CGO_ENABLED=0 go build -tags pureentry, then link with go tool link -linkmode=external + the EF/FSKit frameworks\n` +
  `  5. patch_lcmain.go to point LC_MAIN at main.nsextMainEntry.abi0\n` +
  `  6. check_swift_metadata.go / dump_swift_metadata.go to verify the final binary still carries the metadata\n\n` +
  `CRITICAL: report whether the build actually completed, the furthest step reached, the FIRST failing step, ` +
  `and the precise error text. Note that THIS still uses swiftc in step 1 — that is expected; we are first ` +
  `confirming the swiftc-derived pipeline still builds on this stack before attempting to remove swiftc. ` +
  `Keep all artifacts; list their paths.`,
  { label: 'reproduce:build', phase: 'Reproduce', schema: BUILD_SCHEMA }
)

// ---- Phase 3: Attempt synthesis WITHOUT swiftc (the actual goal) ----
phase('Synthesize')
const buildText = JSON.stringify(build, null, 2).slice(0, 4000)
const synth = await parallel([
  // 3a: can the manifest-only (swiftmeta_bundled) path produce a working metadata object with NO swiftc?
  () => agent(
    `THE CORE QUESTION. The success criterion is a strictly no-swiftc, no-cgo Go entrypoint that ` +
    `ExtensionFoundation actually loads. The 'swiftmeta_bundled' path in /tmp/orig-build-appex.sh uses ` +
    `write_swiftmeta_fixture.go + emit_macho_object.go to build the metadata object from the CHECKED-IN ` +
    `fixture ${RES}/pureentry/swiftmeta_manifest.json — i.e. WITHOUT calling swiftc at build time.\n\n` +
    `Empirically determine: does that fixture fully describe the metadata, or does it contain opaque ` +
    `byte blobs (in data_file sections) that were originally produced by swiftc and merely captured? ` +
    `Examine every section's data_file under any extracted sections dir, and reason about whether the ` +
    `BYTES (not just the relocation structure) of __swift5_* sections, the protocol conformance ` +
    `descriptor, the type metadata, and the ExtensionFoundation/FSKit conformance records could be ` +
    `regenerated from first principles (the Swift ABI mangling + record layout) vs. requiring captured ` +
    `swiftc output. Build context:\n${buildText}\n\n` +
    `Try it: drive write_swiftmeta_fixture.go -> emit_macho_object.go from the checked-in fixture to ` +
    `produce a metadata object WITHOUT swiftc, link it into the Go binary, and report whether it carries ` +
    `a complete-enough metadata surface. The honest distinction we need: is this 'no swiftc AT BUILD ` +
    `TIME but byte-for-byte captured FROM a past swiftc run' (a checked-in blob), or genuinely synthesized? ` +
    `That distinction decides whether the success criterion is truly met.`,
    { label: 'synth:bundled-path', phase: 'Synthesize', schema: SYNTH_SCHEMA }
  ),
  // 3b: independent analysis of what swift5 metadata actually requires to be hand-synthesized
  () => agent(
    `Independent of the existing fixtures: analyze whether the Swift metadata records an FSKit ` +
    `UnaryFileSystemExtension needs (protocol conformance descriptor for ExtensionFoundation.AppExtension ` +
    `and FSKit.UnaryFileSystemExtension, nominal type descriptor for the principal struct, the ` +
    `__swift5_proto/__swift5_protos/__swift5_types section entries) can be SYNTHESIZED from first ` +
    `principles by emitting the documented Swift runtime record layouts directly in Go, with no swiftc.\n\n` +
    `Read ${RES}/pureentry/encode_swift_metadata.go and ${RES}/pureentry/write_swiftmeta_fixture.go to see ` +
    `how far the existing encoder already goes toward synthesis. Use WebSearch/WebFetch for the Swift ABI ` +
    `(swift/docs/ABI/Mangling.rst, TypeMetadata.rst, the runtime's ProtocolConformanceDescriptor / ` +
    `TargetTypeContextDescriptor layouts). Judge per record type: derivable-without-swiftc yes/no/partial, ` +
    `and what specifically blocks the 'no' cases. Be concrete about the hardest record (the protocol ` +
    `conformance descriptor pointing at an external protocol in a system dylib).`,
    { label: 'synth:abi-analysis', phase: 'Synthesize', schema: SYNTH_SCHEMA }
  ),
]).then(r => r.filter(Boolean))

// ---- Phase 4: Adversarially verify the load-bearing claims ----
phase('Verify')
const claims = []
for (const s of synth) {
  claims.push(`Synthesis verdict '${s.verdict}': ${s.evidence}`)
  for (const ob of (s.opaqueBytes || [])) {
    claims.push(`Section ${ob.section} (${ob.whatItContains}) is derivableWithoutSwiftc=${ob.derivableWithoutSwiftc}`)
  }
}
claims.push(`Build stall claim: stalledAt='${build.stalledAt}' reason='${build.stallReason}'`)
const verified = await parallel(claims.slice(0, 10).map((c, i) => () =>
  agent(
    `Adversarially verify this claim about the 9pfs pureentry experiment. Default to skepticism: try to ` +
    `REFUTE it. If you cannot refute it with concrete evidence (a section's bytes, a symbol, a build error, ` +
    `an Apple ABI doc), then it holds. Claim:\n\n"${c}"\n\n` +
    `Files: ${RES} and its pureentry/ tools and fixtures. /tmp/orig-build-appex.sh has the pipeline. ` +
    `You may run otool/nm/python3/go run to check. State whether the claim holds, your confidence, and ` +
    `any counterevidence you found.`,
    { label: `verify:${i}`, phase: 'Verify', schema: VERDICT_SCHEMA }
  )
)).then(r => r.filter(Boolean))

// ---- Phase 5: Reconcile into a verdict ----
phase('Synthesize report')
const report = await agent(
  `Write the definitive conclusion for the 9pfs no-Swift-entrypoint research. The success criterion was: ` +
  `a strictly no-swiftc, no-cgo Go entrypoint that ExtensionFoundation actually loads and mounts.\n\n` +
  `Inputs:\n` +
  `REQUIREMENTS:\n${JSON.stringify(understanding).slice(0, 4000)}\n\n` +
  `BUILD REPRODUCTION:\n${JSON.stringify(build).slice(0, 3000)}\n\n` +
  `SYNTHESIS ATTEMPTS:\n${JSON.stringify(synth).slice(0, 4000)}\n\n` +
  `ADVERSARIAL VERIFICATION:\n${JSON.stringify(verified).slice(0, 4000)}\n\n` +
  `Produce a clear verdict answering: (1) Did the swiftc-derived pipeline still build on this stack? ` +
  `(2) Can the metadata surface be produced at build time WITHOUT swiftc — and is the checked-in fixture ` +
  `genuine synthesis or a captured swiftc blob? (3) Is the strict success criterion (no-swiftc AND ` +
  `no-cgo AND actually loaded by EF) reachable, and if not, name the single hardest blocking requirement ` +
  `with its evidence. (4) The concrete next experiment that would settle any remaining doubt. Be honest ` +
  `and specific; cite sections/symbols/errors. Return the full report as markdown text suitable for ` +
  `RESEARCH.md's "Goal and current state" section.`,
  { label: 'report', phase: 'Synthesize report' }
)

return { understanding, build, synth, verified, report }
