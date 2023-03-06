// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package riscv64asm

type instArg uint8

const (
	_ instArg = iota
	arg_bimm12hilo
	arg_imm12
	arg_imm12hilo
	arg_imm20
	arg_jimm20
	arg_pred
	arg_rd
	arg_rs1
	arg_rs2
	arg_shamtd
	arg_shamtw
	arg_succ
)

const (
	aq = 1 << 26
	rl = 1 << 25
)
