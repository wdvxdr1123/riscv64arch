// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package riscv64asm

import (
	"fmt"
	"strings"
)

// An Op is an RISCV opcode.
type Op uint16

func (op Op) String() string {
	if op >= Op(len(opstr)) || opstr[op] == "" {
		return fmt.Sprintf("Op(%d)", int(op))
	}
	return opstr[op]
}

// An Inst is a single instruction.
type Inst struct {
	Op   Op     // Opcode mnemonic
	Enc  uint32 // Raw encoding bits.
	Args Args   // Instruction arguments, in RISCV manual order.
}

func (i Inst) String() string {
	var args []string
	for _, arg := range i.Args {
		if arg == nil {
			break
		}
		args = append(args, arg.String())
	}
	return i.Op.String() + " " + strings.Join(args, ", ")
}

// An Args holds the instruction arguments.
// If an instruction has fewer than 5 arguments,
// the final elements in the array are nil.
type Args [5]Arg

// An Arg is a single instruction argument, one of these types:
// Reg, Imm
type Arg interface {
	isArg()
	String() string
}

// A Reg is a single register.
type Reg uint16

const (
	X0 Reg = iota
	X1
	X2
	X3
	X4
	X5
	X6
	X7
	X8
	X9
	X10
	X11
	X12
	X13
	X14
	X15
	X16
	X17
	X18
	X19
	X20
	X21
	X22
	X23
	X24
	X25
	X26
	X27
	X28
	X29
	X30
	X31

	F0
	F1
	F2
	F3
	F4
	F5
	F6
	F7
	F8
	F9
	F10
	F11
	F12
	F13
	F14
	F15
	F16
	F17
	F18
	F19
	F20
	F21
	F22
	F23
	F24
	F25
	F26
	F27
	F28
	F29
	F30
	F31
)

func (Reg) isArg() {}

func (r Reg) String() string {
	switch {
	case r <= X31:
		return fmt.Sprintf("X%d", int(r-X0))
	case F0 <= r && r <= F31:
		return fmt.Sprintf("F%d", int(r-F0))
	default:
		return fmt.Sprintf("Reg(%d)", int(r))
	}
}

// An Imm is an integer constant.
type Imm int32

func (Imm) isArg() {}

func (i Imm) String() string {
	return fmt.Sprintf("%d", int32(i))
}
