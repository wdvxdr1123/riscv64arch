// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package riscv64asm

import (
	"strconv"
	"strings"
)

// GNUSyntax returns the GNU assembler syntax for the instruction, as defined by GNU binutils.
// This form typically matches the syntax defined in the RISCV Manual.
func GNUSyntax(inst Inst) string {
	for _, alias := range gnuAliases {
		m := alias.match(inst)
		if m != "" {
			return m
		}
	}
	op := inst.Op.String()
	if isAMO(inst.Op) {
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
		args = append(args, arg.String())
	}
	return op + " " + strings.Join(args, ", ")
}

type alias struct {
	op      Op
	arg     Args
	pattern string
}

func (a *alias) match(inst Inst) string {
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
		m = strings.ReplaceAll(m, "$"+strconv.Itoa(i), arg.String())
	}
	return m
}

var gnuAliases = [...]alias{
	{ADDI, Args{ZERO, ZERO, Imm(0)}, "nop"},
	{ADDI, Args{nil, ZERO, nil}, "li $0, $2"},
	{ADDI, Args{nil, nil, Imm(0)}, "mv $0, $1"},
	{JAL, Args{ZERO, nil}, "j $1"},
	{JALR, Args{ZERO, Mem{Base: RA}}, "ret"},
	{JALR, Args{ZERO, nil}, "jr $1"},
	{SUB, Args{nil, ZERO, nil}, "neg $0, $2"},
	{SUBW, Args{nil, ZERO, nil}, "negw $0, $2"},
	{ADDIW, Args{nil, nil, Imm(0)}, "sext.w $0, $1"},
	{ANDI, Args{nil, nil, Imm(255)}, "zext.b $0, $1"},
	{SLTIU, Args{nil, nil, Imm(1)}, "seqz $0, $1"},
	{BEQ, Args{nil, ZERO, nil}, "beqz $0, $2"},
	{BNE, Args{nil, ZERO, nil}, "bnez $0, $2"},
	{BGE, Args{nil, ZERO, nil}, "bgez $0, $2"},
	{BGE, Args{ZERO, nil, nil}, "blez $1, $2"},
	{BLT, Args{nil, ZERO, nil}, "bltz $0, $2"},
}
