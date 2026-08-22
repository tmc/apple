// Copyright 2026 The tmc/apple Authors. All rights reserved.

//go:build typedefgate

// This file is the acceptance gate for rendering a typedef'd C record as a
// usable Go type. It does not build by default: UTCDateTime is currently
//
//	type UTCDateTime unsafe.Pointer
//
// so UCConvertCFAbsoluteTimeToUTCDateTime's out-parameter is *unsafe.Pointer
// and no caller can read what the call wrote. The test cannot be written
// against the current rendering at all, which is the point -- it is red as a
// compile error today and green on a value assertion once the record is real.
//
// Build it with:
//
//	go test -tags typedefgate ./coreservices
//
// Drop the build tag when the fix lands.
package coreservices_test

import (
	"testing"
	"unsafe"

	"github.com/tmc/apple/coreservices"
)

// TestUTCDateTimeLayout pins the record's shape as the C compiler measures it:
//
//	sizeof=8 align=2 off(highSeconds)=0 off(lowSeconds)=2 off(fraction)=6
//
// Alignment 2 with a four-byte member at offset 2 makes this a packed record.
// Go cannot place these as ordinary fields, so the expected rendering is
// opaque storage plus accessors, as CMTime already does.
func TestUTCDateTimeLayout(t *testing.T) {
	var d coreservices.UTCDateTime
	if got, want := unsafe.Sizeof(d), uintptr(8); got != want {
		t.Errorf("sizeof(UTCDateTime) = %d, want %d", got, want)
	}
	if got, want := unsafe.Alignof(d), uintptr(2); got != want {
		t.Errorf("alignof(UTCDateTime) = %d, want %d", got, want)
	}
}

// TestUCConvertCFAbsoluteTimeToUTCDateTime checks the members against values
// fixed by the two epoch definitions rather than against a round trip. A round
// trip agrees with itself when a pair of bugs is symmetric; these numbers come
// from outside our code.
//
// UTCDateTime counts from 1904-01-01, CFAbsoluteTime from 2001-01-01, and
// 3061152000 is the seconds between them (97 years of 365.25 days). fraction
// is a 1/65536 share of a second, so half a second is 32768.
//
// Every case pins a different member, which is what makes the test detect a
// wrong field order rather than merely a wrong field name. Declaration order
// is highSeconds, lowSeconds, fraction; alphabetical order would lead with
// fraction. Because the members have different widths, reading them in the
// wrong order yields wrong numbers, not just wrong labels.
func TestUCConvertCFAbsoluteTimeToUTCDateTime(t *testing.T) {
	for _, tt := range []struct {
		name           string
		in             float64
		high, fraction uint16
		low            uint32
	}{
		{"cf epoch", 0, 0, 0, 3061152000},
		{"mac epoch", -3061152000, 0, 0, 0},
		{"half second", 0.5, 0, 32768, 3061152000},
		{"past 2^32 seconds", 1233815296, 1, 0, 0},
	} {
		t.Run(tt.name, func(t *testing.T) {
			var d coreservices.UTCDateTime
			if s := coreservices.UCConvertCFAbsoluteTimeToUTCDateTime(tt.in, &d); s != 0 {
				t.Fatalf("UCConvertCFAbsoluteTimeToUTCDateTime(%v) = %d, want 0", tt.in, s)
			}
			if got := d.HighSeconds(); got != tt.high {
				t.Errorf("HighSeconds() = %d, want %d", got, tt.high)
			}
			if got := d.LowSeconds(); got != tt.low {
				t.Errorf("LowSeconds() = %d, want %d", got, tt.low)
			}
			if got := d.Fraction(); got != tt.fraction {
				t.Errorf("Fraction() = %d, want %d", got, tt.fraction)
			}
		})
	}
}

// TestUCConvertUTCDateTimeToCFAbsoluteTime runs the other direction, where the
// caller fills the record and the framework reads it. This is the half that
// makes the current rendering useless rather than merely awkward: with
// UTCDateTime opaque there is no way to set a member, so the call has no
// reachable input.
func TestUCConvertUTCDateTimeToCFAbsoluteTime(t *testing.T) {
	var d coreservices.UTCDateTime
	d.SetLowSeconds(3061152000)
	d.SetFraction(32768)

	var got float64
	if s := coreservices.UCConvertUTCDateTimeToCFAbsoluteTime(&d, &got); s != 0 {
		t.Fatalf("UCConvertUTCDateTimeToCFAbsoluteTime = %d, want 0", s)
	}
	if want := 0.5; got != want {
		t.Errorf("CFAbsoluteTime = %v, want %v", got, want)
	}
}
