package opcodes

type Opcode byte

const (
	Nop Opcode = iota
	Unreachable
	I32Const
	I32Add
	I32Sub
	I32Mul
	I32DivS
	I32DivU
	I32RemS
	I32RemU
	I32And
	I32Or
	I32Xor
	I32Shl
	I32ShrS
	I32ShrU
	I32Rotl
	I32Rotr
	I32Clz
	I32Ctz
	I32PopCnt
	I32EqZ
	I32Eq
	I32Neq
	I32LtS
	I32LtU
	I32LeS
	I32LeU
	I32GtS
	I32GtU
	I32GeS
	I32GeU

	I64Const
	I64Add
	I64Sub
	I64Mul
	I64DivS
	I64DivU
	I64RemS
	I64RemU
	I64Rotl
	I64Rotr
	I64Clz
	I64Ctz
	I64PopCnt
	I64EqZ
	I64And
	I64Or
	I64Xor
	I64Shl
	I64ShrS
	I64ShrU
	I64Eq
	I64Neq
	I64LtS
	I64LtU
	I64LeS
	I64LeU
	I64GtS
	I64GtU
	I64GeS
	I64GeU

	F32Const
	F32Add
	F32Sub
	F32Mul
	F32Div
	F32Sqrt
	F32Min
	F32Max
	F32Ceil
	F32Floor
	F32Trunc
	F32Nearest
	F32Abs
	F32Neg
	F32CopySign
	F32Eq
	F64Const
	F64Add
	F64Sub
	F64Mul
	F64Div
	F64Sqrt
	F64Min
	F64Max
	F64Ceil
	F64Floor
	F64Trunc
	F64Nearest
	F64Abs
	F64Neg
	F64CopySign
	F64Eq
	Jmp
	JmpIf
	JmpTable
	ReturnValue
	ReturnVoid
	GetLocal
	SetLocal
	Call
	CallIndirect
	Phi
)
