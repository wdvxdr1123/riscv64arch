// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package riscv64asm

import (
	"fmt"
	"strconv"
	"strings"
)

// GNUSyntax returns the GNU assembler syntax for the instruction, as defined by GNU binutils.
// This form typically matches the syntax defined in the RISCV Manual.
func GNUSyntax(inst Inst, pc uint64) string {
	for _, alias := range gnuAliases {
		m := alias.match(inst, pc)
		if m != "" {
			return m
		}
	}

	op := inst.Op.String()
	if isAMO(inst.Op) {
		const (
			aq = 1 << 26
			rl = 1 << 25
		)
		if inst.Enc&(aq|rl) != 0 {
			op += "."
		}
		if inst.Enc&aq == aq {
			op += "aq"
		}
		if inst.Enc&rl == rl {
			op += "rl"
		}
	}
	var args []string
	for _, arg := range inst.Args {
		if arg == nil {
			break
		}
		args = append(args, gnuArg(arg, pc))
	}
	return op + " " + strings.Join(args, ", ")
}

func gnuArg(arg Arg, pc uint64) string {
	switch arg := arg.(type) {
	case PCRel:
		return fmt.Sprintf("0x%x", uint64(arg)+pc)
	}
	return arg.String()
}

type alias struct {
	op      Op
	arg     Args
	pattern string
}

func (a *alias) match(inst Inst, pc uint64) string {
	if inst.Op != a.op {
		return ""
	}

	for i, arg := range a.arg {
		if arg == nil {
			continue
		}
		if arg != inst.Args[i] {
			return ""
		}
	}

	m := a.pattern
	for i, arg := range inst.Args {
		if arg == nil {
			break
		}
		m = strings.ReplaceAll(m, "$"+strconv.Itoa(i), gnuArg(arg, pc))
	}
	return m
}

var gnuAliases = [...]alias{
	{ADDI, Args{X0, X0, Imm(0)}, "nop"},
	{ADDI, Args{nil, X0, nil}, "li $0, $2"},
	{ADDI, Args{nil, nil, Imm(0)}, "mv $0, $1"},
	{JAL, Args{X0, nil}, "j $1"},
	{JALR, Args{X0, Mem{Base: X1, Offset: 0}}, "ret"},
	{JALR, Args{X0, nil}, "jr $1"},
	{SUB, Args{nil, X0, nil}, "neg $0, $2"},
	{SUBW, Args{nil, X0, nil}, "negw $0, $2"},
	{ADDIW, Args{nil, nil, Imm(0)}, "sext.w $0, $1"},
	{ANDI, Args{nil, nil, Imm(255)}, "zext.b $0, $1"},
	{SLTIU, Args{nil, nil, Imm(1)}, "seqz $0, $1"},
	{BEQ, Args{nil, X0, nil}, "beqz $0, $2"},
	{BNE, Args{nil, X0, nil}, "bnez $0, $2"},
	{BGE, Args{nil, X0, nil}, "bgez $0, $2"},
	{BGE, Args{X0, nil, nil}, "blez $1, $2"},
	{BLT, Args{nil, X0, nil}, "bltz $0, $2"},
}
