// Package experiment defines the portable contract, arm, sample, and receipt
// types shared by the mlxperf optimization arena and its measurement
// harnesses.
//
// A [Receipt] is the canonical result of one measured experiment. It binds a
// [Workload] identity, the raw [Sample] records in execution order, every
// typed [Refusal] raised while measuring or analyzing, and a final verdict.
// Receipts are sealed with a SHA-256 digest over their canonical JSON form so
// a receipt cannot be edited after the fact without detection.
//
// The zero value of every type here is not meaningful; construct receipts by
// measuring an experiment (see github.com/tmc/apple/x/armbench) or by
// decoding a sealed receipt with [ParseReceipt], which verifies the digest.
package experiment
