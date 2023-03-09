// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package riscv64asm

import (
	"encoding/hex"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func testDecode(t *testing.T, syntax string) {
	input := filepath.Join("testdata", syntax+".txt")
	data, err := os.ReadFile(input)
	if err != nil {
		t.Fatal(err)
	}
	all := string(data)
	for strings.Contains(all, "\t\t") {
		all = strings.ReplaceAll(all, "\t\t", "\t")
	}
	for _, line := range strings.Split(all, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		f := strings.SplitN(line, "\t", 2)
		i := strings.Index(f[0], "|")
		if i < 0 {
			t.Errorf("parsing %q: missing | separator", f[0])
			continue
		}
		if i%2 != 0 {
			t.Errorf("parsing %q: misaligned | separator", f[0])
		}
		code, err := hex.DecodeString(f[0][:i] + f[0][i+1:])
		if err != nil {
			t.Errorf("parsing %q: %v", f[0], err)
			continue
		}
		asm := f[1]
		inst, decodeErr := Decode(code)
		if decodeErr != nil && decodeErr != errUnknown {
			// Some rarely used system instructions are not supported
			// Following logical will filter such unknown instructions
			t.Errorf("parsing %x: %s", code, decodeErr)
			continue
		}
		var out string
		switch syntax {
		case "gnu", "gnu_aliases":
			out = GNUSyntax(inst, 0x10000)
		case "plan9":
			out = GoSyntax(inst, 0x10000, nil)
		default:
			t.Errorf("unknown syntax %q", syntax)
			continue
		}
		if strings.Replace(out, " ", "", -1) != strings.Replace(asm, " ", "", -1) {
			t.Errorf("Decode(%s) [%s] = %s, want %s", strings.Trim(f[0], "|"), syntax, out, asm)
		}
	}
}

func TestDecodeGNUSyntax(t *testing.T) {
	testDecode(t, "gnu")
	testDecode(t, "gnu_aliases")
}

func TestDecodeGoSyntax(t *testing.T) {
	testDecode(t, "plan9")
}
