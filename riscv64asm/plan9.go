package riscv64asm

import (
	"fmt"
	"strconv"
	"strings"
)

// GoSyntax returns the Go assembler syntax for the instruction.
func GoSyntax(inst Inst, pc uint64) string {
	for _, alias := range goAlias {
		m := alias.match(inst, pc, false)
		if m != "" {
			return m
		}
	}

	op := inst.Op.String()
	op = strings.ReplaceAll(op, ".", "")
	op = strings.ToUpper(op)
	var args []string
	for _, arg := range inst.Args {
		if arg == nil {
			break
		}
		if _, ok := arg.(RoundingMode); ok {
			break
		}
		args = append(args, goArg(arg, pc))
	}
	if len(args) == 0 {
		return op
	}
	switch inst.Op {
	default:
		args[0], args[len(args)-1] = args[len(args)-1], args[0]
	case JAL, JALR, BEQ, BNE, BGE, BGEU, BLT, BLTU:
		// no
		switch inst.Op {
		case JAL:
		case JALR:
		default:
			if inst.Args[1] == ZERO {
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

	if isAMO(inst.Op) && inst.Op != LRW && inst.Op != LRD {
		return op + " " + args[1] + ", " + args[0] + ", " + args[2]
	}

	if inst.Args[1] == inst.Args[2] {
		switch inst.Op {
		case FSGNJS:
			return "MOVF" + " " + args[1] + ", " + args[2]
		case FSGNJD:
			return "MOVD" + " " + args[1] + ", " + args[2]
		case FSGNJXS:
			return "FABSS" + " " + args[1] + ", " + args[2]
		case FSGNJXD:
			return "FABSD" + " " + args[1] + ", " + args[2]
		case FSGNJNS:
			return "FNEGS" + " " + args[1] + ", " + args[2]
		case FSGNJND:
			return "FNEGD" + " " + args[1] + ", " + args[2]
		}
	}

	return op + " " + strings.Join(args, ", ")
}

func goArg(arg Arg, pc uint64) string {
	switch arg := arg.(type) {
	case Reg:
		r := arg
		switch {
		case X0 <= r && r <= X31:
			return fmt.Sprintf("X%d", int(r-X0))
		case F0 <= r && r <= F31:
			return fmt.Sprintf("F%d", int(r-F0))
		}
	case Imm:
		return "$" + arg.String()
	case Mem:
		base := "(" + goArg(arg.Base, pc) + ")"
		if arg.Offset != 0 {
			return strconv.Itoa(int(arg.Offset)) + base
		}
		return base
	case PCRel:
		return fmt.Sprintf("%d(PC)", arg/4)
	}
	return ""
}

var goAlias = [...]alias{
	{FLD, Args{nil, nil}, "MOVD $1, $0"},
	{FLW, Args{nil, nil}, "MOVF $1, $0"},
	{FSD, Args{nil, nil}, "MOVD $0, $1"},
	{FSW, Args{nil, nil}, "MOVF $0, $1"},
	{XORI, Args{nil, nil, Imm(-1)}, "NOT $1, $0"},
	{SUB, Args{nil, ZERO, nil}, "NEG $2, $0"},
	{SUBW, Args{nil, ZERO, nil}, "NEGW $2, $0"},

	{JAL, Args{ZERO, nil}, "JMP $1"},
	{JALR, Args{ZERO, nil}, "JMP $1"},
	{BEQ, Args{nil, ZERO, nil}, "BEQZ $0, $2"},
	{SLTIU, Args{nil, nil, Imm(1)}, "SEQZ $1, $0"},
	{SLTU, Args{nil, ZERO, nil}, "SNEZ $2, $0"},
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
	FLW: "MOVW",
	SW:  "MOVW",
	SH:  "MOVH",
	SB:  "MOVB",
	FSD: "MOVD",
	FSW: "MOVW",
}
