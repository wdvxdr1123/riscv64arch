// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package riscv64asm

import (
	"encoding/binary"
	"fmt"
)

// An instFormat describes the format of an instruction encoding.
// An instruction with 32-bit value x matches the format if x&mask == match.
// Finally, the args describe how to decode the instruction arguments.
// args is stored as a fixed-size array; if there are fewer than len(args) arguments,
// args[i] == 0 marks the end of the argument list.
type instFormat struct {
	op    Op
	mask  uint32
	match uint32
	args  instArgs
}

var (
	errShort   = fmt.Errorf("truncated instruction")
	errUnknown = fmt.Errorf("unknown instruction")
)

var decoderCover []bool

func init() {
	decoderCover = make([]bool, len(instFormats))
}

// Decode decodes the 4 bytes in src as a single instruction.
func Decode(src []byte) (inst Inst, err error) {
	if len(src) < 4 {
		return Inst{}, errShort
	}

	x := binary.LittleEndian.Uint32(src)

Search:
	for i := range instFormats {
		f := &instFormats[i]
		if x&f.mask != f.match {
			continue
		}
		// Decode args.
		var args Args
		for j, aop := range f.args {
			if aop == 0 {
				break
			}
			arg := decodeArg(aop, x)
			if arg == nil { // Cannot decode argument
				continue Search
			}
			args[j] = arg
		}
		decoderCover[i] = true
		inst = Inst{
			Op:   f.op,
			Args: args,
			Enc:  x,
		}
		transformInst(&inst)
		return inst, nil
	}
	return Inst{}, errUnknown
}

type instArgs [5]instArg

// signExtend sign extends val starting at bit bit.
func signExtend(val int32, bit uint) int32 {
	return val << (32 - bit) >> (32 - bit)
}

// decodeArg decodes the arg described by aop from the instruction bits x.
// It returns nil if x cannot be decoded according to aop.
func decodeArg(aop instArg, x uint32) Arg {
	switch aop {
	case arg_rd:
		return X0 + Reg((x>>7)&(1<<5-1))
	case arg_rs1:
		return X0 + Reg((x>>15)&(1<<5-1))
	case arg_rs2:
		return X0 + Reg((x>>20)&(1<<5-1))
	case arg_imm12:
		imm := int32((x >> 20) & (1<<12 - 1))
		return Imm(signExtend(imm, 12))
	case arg_imm20:
		imm := int32((x >> 12) & (1<<20 - 1))
		return Imm(signExtend(imm, 20))
	case arg_jimm20:
		//      31 | 30     21 |      20 | 19      12 | 11
		// imm[20] | imm[10:1] | imm[11] | imm[19:12] | ...
		imm := (int32(x>>31) & (1<<1 - 1)) << 20
		imm |= (int32(x>>21) & (1<<10 - 1)) << 1
		imm |= (int32(x>>20) & (1<<1 - 1)) << 11
		imm |= (int32(x>>12) & (1<<8 - 1)) << 12
		return Imm(signExtend(imm, 21))
	case arg_imm12hilo:
		//        25 |     | 11     7 |
		// imm[11:5] | ... | imm[4:0] | ...
		imm := (int32(x>>25) & (1<<7 - 1)) << 5
		imm |= (int32(x>>7) & (1<<5 - 1)) << 0
		return Imm(signExtend(imm, 12))
	case arg_bimm12hilo:
		//      31 | 30     25 |     | 11     8 |       7 |
		// imm[12] | imm[10:5] | ... | imm[4:1] | imm[11] | ...
		imm := (int32(x>>31) & (1<<1 - 1)) << 12
		imm |= (int32(x>>25) & (1<<6 - 1)) << 5
		imm |= (int32(x>>8) & (1<<4 - 1)) << 1
		imm |= (int32(x>>7) & (1<<1 - 1)) << 11
		return Imm(signExtend(imm, 13))
	case arg_shamtd:
		imm := int32(x>>20) & (1<<6 - 1)
		return Imm(imm)
	case arg_shamtw:
		imm := int32(x>>20) & (1<<5 - 1)
		return Imm(imm)
	case arg_rs3:
		return X0 + Reg((x>>27)&(1<<5-1))
	case arg_rm:
		return RoundingMode((x >> 12) & (1<<3 - 1))
	}
	return nil
}

func mem(base Arg, offset Arg) Arg {
	if offset == nil {
		return Mem{Base: base.(Reg)}
	}
	return Mem{Base: base.(Reg), Offset: int32(offset.(Imm))}
}

func transformInst(i *Inst) {
	op := i.Op
	if isAMO(op) {
		switch op {
		case LRD, LRW:
			// lr.d rd, (rs1)
			i.Args[1] = mem(i.Args[1], nil)
		default:
			// op rd, rs2, (rs1)
			i.Args[1], i.Args[2] = i.Args[2], mem(i.Args[1], nil)
		}
	}
	if isFloatOp(op) {
		switch op {
		case FSD, FSW:
			// fsd rs2, offset(rs1)
			rs1 := i.Args[0]
			offset := i.Args[2]
			i.Args[0] = i.Args[1].(Reg) - X0 + F0
			i.Args[1] = mem(rs1, offset)
			i.Args[2] = nil
		case FLD, FLW:
			i.Args[0] = i.Args[0].(Reg) - X0 + F0
		default:
			for index, arg := range i.Args {
				if gp, ok := arg.(Reg); ok {
					i.Args[index] = gp - X0 + F0
				}
			}
		}
	}
	switch i.Op {
	case JALR, LD, LW, LWU, LH, LHU, LB, LBU, SD, SW, SH, SB, FLW, FLD:
		i.Args[1] = mem(i.Args[1], i.Args[2])
		i.Args[2] = nil
	}
}
