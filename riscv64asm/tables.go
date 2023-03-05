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
	FENCE
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
	ADD:      "ADD",
	ADDI:     "ADDI",
	ADDIW:    "ADDIW",
	ADDW:     "ADDW",
	AMOADDD:  "AMOADD.D",
	AMOADDW:  "AMOADD.W",
	AMOANDD:  "AMOAND.D",
	AMOANDW:  "AMOAND.W",
	AMOMAXD:  "AMOMAX.D",
	AMOMAXW:  "AMOMAX.W",
	AMOMAXUD: "AMOMAXU.D",
	AMOMAXUW: "AMOMAXU.W",
	AMOMIND:  "AMOMIN.D",
	AMOMINW:  "AMOMIN.W",
	AMOMINUD: "AMOMINU.D",
	AMOMINUW: "AMOMINU.W",
	AMOORD:   "AMOOR.D",
	AMOORW:   "AMOOR.W",
	AMOSWAPD: "AMOSWAP.D",
	AMOSWAPW: "AMOSWAP.W",
	AMOXORD:  "AMOXOR.D",
	AMOXORW:  "AMOXOR.W",
	AND:      "AND",
	ANDI:     "ANDI",
	AUIPC:    "AUIPC",
	BEQ:      "BEQ",
	BGE:      "BGE",
	BGEU:     "BGEU",
	BLT:      "BLT",
	BLTU:     "BLTU",
	BNE:      "BNE",
	DIV:      "DIV",
	DIVU:     "DIVU",
	DIVUW:    "DIVUW",
	DIVW:     "DIVW",
	EBREAK:   "EBREAK",
	ECALL:    "ECALL",
	FENCE:    "FENCE",
	JAL:      "JAL",
	JALR:     "JALR",
	LB:       "LB",
	LBU:      "LBU",
	LD:       "LD",
	LH:       "LH",
	LHU:      "LHU",
	LRD:      "LR.D",
	LRW:      "LR.W",
	LUI:      "LUI",
	LW:       "LW",
	LWU:      "LWU",
	MUL:      "MUL",
	MULH:     "MULH",
	MULHSU:   "MULHSU",
	MULHU:    "MULHU",
	MULW:     "MULW",
	OR:       "OR",
	ORI:      "ORI",
	REM:      "REM",
	REMU:     "REMU",
	REMUW:    "REMUW",
	REMW:     "REMW",
	SB:       "SB",
	SCD:      "SC.D",
	SCW:      "SC.W",
	SD:       "SD",
	SH:       "SH",
	SLL:      "SLL",
	SLLI:     "SLLI",
	SLLIW:    "SLLIW",
	SLLW:     "SLLW",
	SLT:      "SLT",
	SLTI:     "SLTI",
	SLTIU:    "SLTIU",
	SLTU:     "SLTU",
	SRA:      "SRA",
	SRAI:     "SRAI",
	SRAIW:    "SRAIW",
	SRAW:     "SRAW",
	SRL:      "SRL",
	SRLI:     "SRLI",
	SRLIW:    "SRLIW",
	SRLW:     "SRLW",
	SUB:      "SUB",
	SUBW:     "SUBW",
	SW:       "SW",
	XOR:      "XOR",
	XORI:     "XORI",
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
	{FENCE, 0x707f, 0xf, instArgs{arg_pred, arg_succ}},
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
