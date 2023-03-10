// Code generated by riscvspec. DO NOT EDIT.

// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package riscv64asm

const (
	_ Op = iota
	ADD
	ADDI
	ADDIW
	ADDW
	AMOADDD
	AMOADDW
	AMOANDD
	AMOANDW
	AMOMAXD
	AMOMAXW
	AMOMAXUD
	AMOMAXUW
	AMOMIND
	AMOMINW
	AMOMINUD
	AMOMINUW
	AMOORD
	AMOORW
	AMOSWAPD
	AMOSWAPW
	AMOXORD
	AMOXORW
	AND
	ANDI
	AUIPC
	BEQ
	BGE
	BGEU
	BLT
	BLTU
	BNE
	DIV
	DIVU
	DIVUW
	DIVW
	EBREAK
	ECALL
	FADDD
	FADDS
	FCLASSD
	FCLASSS
	FCVTDL
	FCVTDLU
	FCVTDS
	FCVTDW
	FCVTDWU
	FCVTLD
	FCVTLS
	FCVTLUD
	FCVTLUS
	FCVTSD
	FCVTSL
	FCVTSLU
	FCVTSW
	FCVTSWU
	FCVTWD
	FCVTWS
	FCVTWUD
	FCVTWUS
	FDIVD
	FDIVS
	FENCE
	FEQD
	FEQS
	FLD
	FLED
	FLES
	FLTD
	FLTS
	FLW
	FMADDD
	FMADDS
	FMAXD
	FMAXS
	FMIND
	FMINS
	FMSUBD
	FMSUBS
	FMULD
	FMULS
	FMVDX
	FMVWX
	FMVXD
	FMVXW
	FNMADDD
	FNMADDS
	FNMSUBD
	FNMSUBS
	FSD
	FSGNJD
	FSGNJS
	FSGNJND
	FSGNJNS
	FSGNJXD
	FSGNJXS
	FSQRTD
	FSQRTS
	FSUBD
	FSUBS
	FSW
	JAL
	JALR
	LB
	LBU
	LD
	LH
	LHU
	LRD
	LRW
	LUI
	LW
	LWU
	MUL
	MULH
	MULHSU
	MULHU
	MULW
	OR
	ORI
	REM
	REMU
	REMUW
	REMW
	SB
	SCD
	SCW
	SD
	SH
	SLL
	SLLI
	SLLIW
	SLLW
	SLT
	SLTI
	SLTIU
	SLTU
	SRA
	SRAI
	SRAIW
	SRAW
	SRL
	SRLI
	SRLIW
	SRLW
	SUB
	SUBW
	SW
	XOR
	XORI
)

var opstr = [...]string{
	ADD:      "add",
	ADDI:     "addi",
	ADDIW:    "addiw",
	ADDW:     "addw",
	AMOADDD:  "amoadd.d",
	AMOADDW:  "amoadd.w",
	AMOANDD:  "amoand.d",
	AMOANDW:  "amoand.w",
	AMOMAXD:  "amomax.d",
	AMOMAXW:  "amomax.w",
	AMOMAXUD: "amomaxu.d",
	AMOMAXUW: "amomaxu.w",
	AMOMIND:  "amomin.d",
	AMOMINW:  "amomin.w",
	AMOMINUD: "amominu.d",
	AMOMINUW: "amominu.w",
	AMOORD:   "amoor.d",
	AMOORW:   "amoor.w",
	AMOSWAPD: "amoswap.d",
	AMOSWAPW: "amoswap.w",
	AMOXORD:  "amoxor.d",
	AMOXORW:  "amoxor.w",
	AND:      "and",
	ANDI:     "andi",
	AUIPC:    "auipc",
	BEQ:      "beq",
	BGE:      "bge",
	BGEU:     "bgeu",
	BLT:      "blt",
	BLTU:     "bltu",
	BNE:      "bne",
	DIV:      "div",
	DIVU:     "divu",
	DIVUW:    "divuw",
	DIVW:     "divw",
	EBREAK:   "ebreak",
	ECALL:    "ecall",
	FADDD:    "fadd.d",
	FADDS:    "fadd.s",
	FCLASSD:  "fclass.d",
	FCLASSS:  "fclass.s",
	FCVTDL:   "fcvt.d.l",
	FCVTDLU:  "fcvt.d.lu",
	FCVTDS:   "fcvt.d.s",
	FCVTDW:   "fcvt.d.w",
	FCVTDWU:  "fcvt.d.wu",
	FCVTLD:   "fcvt.l.d",
	FCVTLS:   "fcvt.l.s",
	FCVTLUD:  "fcvt.lu.d",
	FCVTLUS:  "fcvt.lu.s",
	FCVTSD:   "fcvt.s.d",
	FCVTSL:   "fcvt.s.l",
	FCVTSLU:  "fcvt.s.lu",
	FCVTSW:   "fcvt.s.w",
	FCVTSWU:  "fcvt.s.wu",
	FCVTWD:   "fcvt.w.d",
	FCVTWS:   "fcvt.w.s",
	FCVTWUD:  "fcvt.wu.d",
	FCVTWUS:  "fcvt.wu.s",
	FDIVD:    "fdiv.d",
	FDIVS:    "fdiv.s",
	FENCE:    "fence",
	FEQD:     "feq.d",
	FEQS:     "feq.s",
	FLD:      "fld",
	FLED:     "fle.d",
	FLES:     "fle.s",
	FLTD:     "flt.d",
	FLTS:     "flt.s",
	FLW:      "flw",
	FMADDD:   "fmadd.d",
	FMADDS:   "fmadd.s",
	FMAXD:    "fmax.d",
	FMAXS:    "fmax.s",
	FMIND:    "fmin.d",
	FMINS:    "fmin.s",
	FMSUBD:   "fmsub.d",
	FMSUBS:   "fmsub.s",
	FMULD:    "fmul.d",
	FMULS:    "fmul.s",
	FMVDX:    "fmv.d.x",
	FMVWX:    "fmv.w.x",
	FMVXD:    "fmv.x.d",
	FMVXW:    "fmv.x.w",
	FNMADDD:  "fnmadd.d",
	FNMADDS:  "fnmadd.s",
	FNMSUBD:  "fnmsub.d",
	FNMSUBS:  "fnmsub.s",
	FSD:      "fsd",
	FSGNJD:   "fsgnj.d",
	FSGNJS:   "fsgnj.s",
	FSGNJND:  "fsgnjn.d",
	FSGNJNS:  "fsgnjn.s",
	FSGNJXD:  "fsgnjx.d",
	FSGNJXS:  "fsgnjx.s",
	FSQRTD:   "fsqrt.d",
	FSQRTS:   "fsqrt.s",
	FSUBD:    "fsub.d",
	FSUBS:    "fsub.s",
	FSW:      "fsw",
	JAL:      "jal",
	JALR:     "jalr",
	LB:       "lb",
	LBU:      "lbu",
	LD:       "ld",
	LH:       "lh",
	LHU:      "lhu",
	LRD:      "lr.d",
	LRW:      "lr.w",
	LUI:      "lui",
	LW:       "lw",
	LWU:      "lwu",
	MUL:      "mul",
	MULH:     "mulh",
	MULHSU:   "mulhsu",
	MULHU:    "mulhu",
	MULW:     "mulw",
	OR:       "or",
	ORI:      "ori",
	REM:      "rem",
	REMU:     "remu",
	REMUW:    "remuw",
	REMW:     "remw",
	SB:       "sb",
	SCD:      "sc.d",
	SCW:      "sc.w",
	SD:       "sd",
	SH:       "sh",
	SLL:      "sll",
	SLLI:     "slli",
	SLLIW:    "slliw",
	SLLW:     "sllw",
	SLT:      "slt",
	SLTI:     "slti",
	SLTIU:    "sltiu",
	SLTU:     "sltu",
	SRA:      "sra",
	SRAI:     "srai",
	SRAIW:    "sraiw",
	SRAW:     "sraw",
	SRL:      "srl",
	SRLI:     "srli",
	SRLIW:    "srliw",
	SRLW:     "srlw",
	SUB:      "sub",
	SUBW:     "subw",
	SW:       "sw",
	XOR:      "xor",
	XORI:     "xori",
}

var instFormats = [...]instFormat{
	{ADD, 0xfe00707f, 0x33, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{ADDI, 0x707f, 0x13, instArgs{arg_rd, arg_rs1, arg_imm12}},
	{ADDIW, 0x707f, 0x1b, instArgs{arg_rd, arg_rs1, arg_imm12}},
	{ADDW, 0xfe00707f, 0x3b, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{AMOADDD, 0xf800707f, 0x302f, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{AMOADDW, 0xf800707f, 0x202f, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{AMOANDD, 0xf800707f, 0x6000302f, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{AMOANDW, 0xf800707f, 0x6000202f, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{AMOMAXD, 0xf800707f, 0xa000302f, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{AMOMAXW, 0xf800707f, 0xa000202f, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{AMOMAXUD, 0xf800707f, 0xe000302f, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{AMOMAXUW, 0xf800707f, 0xe000202f, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{AMOMIND, 0xf800707f, 0x8000302f, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{AMOMINW, 0xf800707f, 0x8000202f, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{AMOMINUD, 0xf800707f, 0xc000302f, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{AMOMINUW, 0xf800707f, 0xc000202f, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{AMOORD, 0xf800707f, 0x4000302f, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{AMOORW, 0xf800707f, 0x4000202f, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{AMOSWAPD, 0xf800707f, 0x800302f, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{AMOSWAPW, 0xf800707f, 0x800202f, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{AMOXORD, 0xf800707f, 0x2000302f, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{AMOXORW, 0xf800707f, 0x2000202f, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{AND, 0xfe00707f, 0x7033, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{ANDI, 0x707f, 0x7013, instArgs{arg_rd, arg_rs1, arg_imm12}},
	{AUIPC, 0x7f, 0x17, instArgs{arg_rd, arg_imm20}},
	{BEQ, 0x707f, 0x63, instArgs{arg_rs1, arg_rs2, arg_bimm12hilo}},
	{BGE, 0x707f, 0x5063, instArgs{arg_rs1, arg_rs2, arg_bimm12hilo}},
	{BGEU, 0x707f, 0x7063, instArgs{arg_rs1, arg_rs2, arg_bimm12hilo}},
	{BLT, 0x707f, 0x4063, instArgs{arg_rs1, arg_rs2, arg_bimm12hilo}},
	{BLTU, 0x707f, 0x6063, instArgs{arg_rs1, arg_rs2, arg_bimm12hilo}},
	{BNE, 0x707f, 0x1063, instArgs{arg_rs1, arg_rs2, arg_bimm12hilo}},
	{DIV, 0xfe00707f, 0x2004033, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{DIVU, 0xfe00707f, 0x2005033, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{DIVUW, 0xfe00707f, 0x200503b, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{DIVW, 0xfe00707f, 0x200403b, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{EBREAK, 0xffffffff, 0x100073, instArgs{}},
	{ECALL, 0xffffffff, 0x73, instArgs{}},
	{FADDD, 0xfe00007f, 0x2000053, instArgs{arg_rd, arg_rs1, arg_rs2, arg_rm}},
	{FADDS, 0xfe00007f, 0x53, instArgs{arg_rd, arg_rs1, arg_rs2, arg_rm}},
	{FCLASSD, 0xfff0707f, 0xe2001053, instArgs{arg_rd, arg_rs1}},
	{FCLASSS, 0xfff0707f, 0xe0001053, instArgs{arg_rd, arg_rs1}},
	{FCVTDL, 0xfff0007f, 0xd2200053, instArgs{arg_rd, arg_rs1, arg_rm}},
	{FCVTDLU, 0xfff0007f, 0xd2300053, instArgs{arg_rd, arg_rs1, arg_rm}},
	{FCVTDS, 0xfff0007f, 0x42000053, instArgs{arg_rd, arg_rs1, arg_rm}},
	{FCVTDW, 0xfff0007f, 0xd2000053, instArgs{arg_rd, arg_rs1, arg_rm}},
	{FCVTDWU, 0xfff0007f, 0xd2100053, instArgs{arg_rd, arg_rs1, arg_rm}},
	{FCVTLD, 0xfff0007f, 0xc2200053, instArgs{arg_rd, arg_rs1, arg_rm}},
	{FCVTLS, 0xfff0007f, 0xc0200053, instArgs{arg_rd, arg_rs1, arg_rm}},
	{FCVTLUD, 0xfff0007f, 0xc2300053, instArgs{arg_rd, arg_rs1, arg_rm}},
	{FCVTLUS, 0xfff0007f, 0xc0300053, instArgs{arg_rd, arg_rs1, arg_rm}},
	{FCVTSD, 0xfff0007f, 0x40100053, instArgs{arg_rd, arg_rs1, arg_rm}},
	{FCVTSL, 0xfff0007f, 0xd0200053, instArgs{arg_rd, arg_rs1, arg_rm}},
	{FCVTSLU, 0xfff0007f, 0xd0300053, instArgs{arg_rd, arg_rs1, arg_rm}},
	{FCVTSW, 0xfff0007f, 0xd0000053, instArgs{arg_rd, arg_rs1, arg_rm}},
	{FCVTSWU, 0xfff0007f, 0xd0100053, instArgs{arg_rd, arg_rs1, arg_rm}},
	{FCVTWD, 0xfff0007f, 0xc2000053, instArgs{arg_rd, arg_rs1, arg_rm}},
	{FCVTWS, 0xfff0007f, 0xc0000053, instArgs{arg_rd, arg_rs1, arg_rm}},
	{FCVTWUD, 0xfff0007f, 0xc2100053, instArgs{arg_rd, arg_rs1, arg_rm}},
	{FCVTWUS, 0xfff0007f, 0xc0100053, instArgs{arg_rd, arg_rs1, arg_rm}},
	{FDIVD, 0xfe00007f, 0x1a000053, instArgs{arg_rd, arg_rs1, arg_rs2, arg_rm}},
	{FDIVS, 0xfe00007f, 0x18000053, instArgs{arg_rd, arg_rs1, arg_rs2, arg_rm}},
	{FENCE, 0x707f, 0xf, instArgs{arg_pred, arg_succ}},
	{FEQD, 0xfe00707f, 0xa2002053, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{FEQS, 0xfe00707f, 0xa0002053, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{FLD, 0x707f, 0x3007, instArgs{arg_rd, arg_rs1, arg_imm12}},
	{FLED, 0xfe00707f, 0xa2000053, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{FLES, 0xfe00707f, 0xa0000053, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{FLTD, 0xfe00707f, 0xa2001053, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{FLTS, 0xfe00707f, 0xa0001053, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{FLW, 0x707f, 0x2007, instArgs{arg_rd, arg_rs1, arg_imm12}},
	{FMADDD, 0x600007f, 0x2000043, instArgs{arg_rd, arg_rs1, arg_rs2, arg_rs3, arg_rm}},
	{FMADDS, 0x600007f, 0x43, instArgs{arg_rd, arg_rs1, arg_rs2, arg_rs3, arg_rm}},
	{FMAXD, 0xfe00707f, 0x2a001053, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{FMAXS, 0xfe00707f, 0x28001053, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{FMIND, 0xfe00707f, 0x2a000053, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{FMINS, 0xfe00707f, 0x28000053, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{FMSUBD, 0x600007f, 0x2000047, instArgs{arg_rd, arg_rs1, arg_rs2, arg_rs3, arg_rm}},
	{FMSUBS, 0x600007f, 0x47, instArgs{arg_rd, arg_rs1, arg_rs2, arg_rs3, arg_rm}},
	{FMULD, 0xfe00007f, 0x12000053, instArgs{arg_rd, arg_rs1, arg_rs2, arg_rm}},
	{FMULS, 0xfe00007f, 0x10000053, instArgs{arg_rd, arg_rs1, arg_rs2, arg_rm}},
	{FMVDX, 0xfff0707f, 0xf2000053, instArgs{arg_rd, arg_rs1}},
	{FMVWX, 0xfff0707f, 0xf0000053, instArgs{arg_rd, arg_rs1}},
	{FMVXD, 0xfff0707f, 0xe2000053, instArgs{arg_rd, arg_rs1}},
	{FMVXW, 0xfff0707f, 0xe0000053, instArgs{arg_rd, arg_rs1}},
	{FNMADDD, 0x600007f, 0x200004f, instArgs{arg_rd, arg_rs1, arg_rs2, arg_rs3, arg_rm}},
	{FNMADDS, 0x600007f, 0x4f, instArgs{arg_rd, arg_rs1, arg_rs2, arg_rs3, arg_rm}},
	{FNMSUBD, 0x600007f, 0x200004b, instArgs{arg_rd, arg_rs1, arg_rs2, arg_rs3, arg_rm}},
	{FNMSUBS, 0x600007f, 0x4b, instArgs{arg_rd, arg_rs1, arg_rs2, arg_rs3, arg_rm}},
	{FSD, 0x707f, 0x3027, instArgs{arg_rs1, arg_rs2, arg_imm12hilo}},
	{FSGNJD, 0xfe00707f, 0x22000053, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{FSGNJS, 0xfe00707f, 0x20000053, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{FSGNJND, 0xfe00707f, 0x22001053, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{FSGNJNS, 0xfe00707f, 0x20001053, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{FSGNJXD, 0xfe00707f, 0x22002053, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{FSGNJXS, 0xfe00707f, 0x20002053, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{FSQRTD, 0xfff0007f, 0x5a000053, instArgs{arg_rd, arg_rs1, arg_rm}},
	{FSQRTS, 0xfff0007f, 0x58000053, instArgs{arg_rd, arg_rs1, arg_rm}},
	{FSUBD, 0xfe00007f, 0xa000053, instArgs{arg_rd, arg_rs1, arg_rs2, arg_rm}},
	{FSUBS, 0xfe00007f, 0x8000053, instArgs{arg_rd, arg_rs1, arg_rs2, arg_rm}},
	{FSW, 0x707f, 0x2027, instArgs{arg_rs1, arg_rs2, arg_imm12hilo}},
	{JAL, 0x7f, 0x6f, instArgs{arg_rd, arg_jimm20}},
	{JALR, 0x707f, 0x67, instArgs{arg_rd, arg_rs1, arg_imm12}},
	{LB, 0x707f, 0x3, instArgs{arg_rd, arg_rs1, arg_imm12}},
	{LBU, 0x707f, 0x4003, instArgs{arg_rd, arg_rs1, arg_imm12}},
	{LD, 0x707f, 0x3003, instArgs{arg_rd, arg_rs1, arg_imm12}},
	{LH, 0x707f, 0x1003, instArgs{arg_rd, arg_rs1, arg_imm12}},
	{LHU, 0x707f, 0x5003, instArgs{arg_rd, arg_rs1, arg_imm12}},
	{LRD, 0xf9f0707f, 0x1000302f, instArgs{arg_rd, arg_rs1}},
	{LRW, 0xf9f0707f, 0x1000202f, instArgs{arg_rd, arg_rs1}},
	{LUI, 0x7f, 0x37, instArgs{arg_rd, arg_imm20}},
	{LW, 0x707f, 0x2003, instArgs{arg_rd, arg_rs1, arg_imm12}},
	{LWU, 0x707f, 0x6003, instArgs{arg_rd, arg_rs1, arg_imm12}},
	{MUL, 0xfe00707f, 0x2000033, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{MULH, 0xfe00707f, 0x2001033, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{MULHSU, 0xfe00707f, 0x2002033, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{MULHU, 0xfe00707f, 0x2003033, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{MULW, 0xfe00707f, 0x200003b, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{OR, 0xfe00707f, 0x6033, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{ORI, 0x707f, 0x6013, instArgs{arg_rd, arg_rs1, arg_imm12}},
	{REM, 0xfe00707f, 0x2006033, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{REMU, 0xfe00707f, 0x2007033, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{REMUW, 0xfe00707f, 0x200703b, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{REMW, 0xfe00707f, 0x200603b, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{SB, 0x707f, 0x23, instArgs{arg_rs1, arg_rs2, arg_imm12hilo}},
	{SCD, 0xf800707f, 0x1800302f, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{SCW, 0xf800707f, 0x1800202f, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{SD, 0x707f, 0x3023, instArgs{arg_rs1, arg_rs2, arg_imm12hilo}},
	{SH, 0x707f, 0x1023, instArgs{arg_rs1, arg_rs2, arg_imm12hilo}},
	{SLL, 0xfe00707f, 0x1033, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{SLLI, 0xfc00707f, 0x1013, instArgs{arg_rd, arg_rs1, arg_shamtd}},
	{SLLIW, 0xfe00707f, 0x101b, instArgs{arg_rd, arg_rs1, arg_shamtw}},
	{SLLW, 0xfe00707f, 0x103b, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{SLT, 0xfe00707f, 0x2033, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{SLTI, 0x707f, 0x2013, instArgs{arg_rd, arg_rs1, arg_imm12}},
	{SLTIU, 0x707f, 0x3013, instArgs{arg_rd, arg_rs1, arg_imm12}},
	{SLTU, 0xfe00707f, 0x3033, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{SRA, 0xfe00707f, 0x40005033, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{SRAI, 0xfc00707f, 0x40005013, instArgs{arg_rd, arg_rs1, arg_shamtd}},
	{SRAIW, 0xfe00707f, 0x4000501b, instArgs{arg_rd, arg_rs1, arg_shamtw}},
	{SRAW, 0xfe00707f, 0x4000503b, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{SRL, 0xfe00707f, 0x5033, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{SRLI, 0xfc00707f, 0x5013, instArgs{arg_rd, arg_rs1, arg_shamtd}},
	{SRLIW, 0xfe00707f, 0x501b, instArgs{arg_rd, arg_rs1, arg_shamtw}},
	{SRLW, 0xfe00707f, 0x503b, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{SUB, 0xfe00707f, 0x40000033, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{SUBW, 0xfe00707f, 0x4000003b, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{SW, 0x707f, 0x2023, instArgs{arg_rs1, arg_rs2, arg_imm12hilo}},
	{XOR, 0xfe00707f, 0x4033, instArgs{arg_rd, arg_rs1, arg_rs2}},
	{XORI, 0x707f, 0x4013, instArgs{arg_rd, arg_rs1, arg_imm12}},
}
