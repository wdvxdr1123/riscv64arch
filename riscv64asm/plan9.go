package riscv64asm

import (
	"fmt"
	"strconv"
	"strings"
)

// GoSyntax returns the Go assembler syntax for the instruction.
func GoSyntax(inst Inst, pc uint64, symname func(addr uint64) (sym string, base uint64)) string {
	if symname == nil {
		symname = func(addr uint64) (sym string, base uint64) {
			return "", 0
		}
	}

	op := inst.Op.String()
	op = strings.ReplaceAll(op, ".", "")
	op = strings.ToUpper(op)
	var args []string
loop:
	for _, arg := range inst.Args {
		switch arg.(type) {
		case nil, RoundingMode, FenceField:
			break loop
		}
		args = append(args, goArg(arg, pc, symname))
	}

	if isAMO(inst.Op) && inst.Op != LRW && inst.Op != LRD {
		return op + " " + args[1] + ", " + args[2] + ", " + args[0]
	}

	if inst.Args[1] == inst.Args[2] {
		switch inst.Op {
		case FSGNJS:
			return "MOVF" + " " + args[1] + ", " + args[0]
		case FSGNJD:
			return "MOVD" + " " + args[1] + ", " + args[0]
		case FSGNJXS:
			return "FABSS" + " " + args[1] + ", " + args[0]
		case FSGNJXD:
			return "FABSD" + " " + args[1] + ", " + args[0]
		case FSGNJNS:
			return "FNEGS" + " " + args[1] + ", " + args[0]
		case FSGNJND:
			return "FNEGD" + " " + args[1] + ", " + args[0]
		}
	}

	switch inst.Op {
	default:
		switch len(args) {
		case 0:
			return op
		}
		args[0], args[len(args)-1] = args[len(args)-1], args[0]

		return op + " " + strings.Join(args, ", ")
	case JAL, JALR, BEQ, BNE, BGE, BGEU, BLT, BLTU:
		switch inst.Op {
		case JAL, JALR:
			if inst.Args[0] == X0 {
				return "JMP " + args[1]
			}
		default:
			if inst.Args[1] == X0 {
				return op + "Z " + args[0] + ", " + args[2]
			}
		}
	case FMADDS, FMSUBS, FNMADDS, FNMSUBS,
		FMADDD, FMSUBD, FNMADDD, FNMSUBD:
		args = append(args[1:], args[0])
	case LD, LW, LH, LB, LWU, LHU, LBU, FLD, FLW: // loads
		return plan9MOVOp[inst.Op] + " " + args[1] + ", " + args[0]
	case SD, SW, SH, SB, FSD, FSW: // stores
		return plan9MOVOp[inst.Op] + " " + args[0] + ", " + args[1]
	}
	return op + " " + strings.Join(args, ", ")
}

// plan9Arg formats arg according to Plan 9 rules.
func goArg(arg Arg, pc uint64, symname func(addr uint64) (sym string, base uint64)) string {
	switch arg := arg.(type) {
	case Reg:
		switch {
		case X0 <= arg && arg <= X31:
			return fmt.Sprintf("X%d", int(arg-X0))
		case F0 <= arg && arg <= F31:
			return fmt.Sprintf("F%d", int(arg-F0))
		}
	case Imm:
		return "$" + arg.String()
	case Mem:
		base := "(" + goArg(arg.Base, pc, symname) + ")"
		if arg.Offset != 0 {
			return strconv.Itoa(int(arg.Offset)) + base
		}
		return base
	case PCRel:
		addr := pc + uint64(arg)
		s, base := symname(addr)
		if s != "" && addr == base {
			return fmt.Sprintf("%s(SB)", s)
		}
		return fmt.Sprintf("%#x", addr)
	}
	return ""
}

var plan9MOVOp = map[Op]string{
	LD:  "MOV",
	LW:  "MOVW",
	LH:  "MOVH",
	LB:  "MOVB",
	SD:  "MOV",
	LWU: "MOVWU",
	LHU: "MOVHU",
	LBU: "MOVBU",
	FLD: "MOVD",
	FLW: "MOVF",
	SW:  "MOVW",
	SH:  "MOVH",
	SB:  "MOVB",
	FSD: "MOVD",
	FSW: "MOVF",
}
