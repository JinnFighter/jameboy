package cpu

type CpuInstance struct {
	AF Register
	BC Register
	DE Register
	HL Register
	SP Register
	PC Register
}

func (cpu *CpuInstance) Init() {
	cpu.AF = Register{}
	cpu.BC = Register{}
	cpu.DE = Register{}
	cpu.HL = Register{}
	cpu.SP = Register{}
	cpu.PC = Register{}
}

func (cpu *CpuInstance) Terminate() {
	cpu.AF.value = 0
	cpu.BC.value = 0
	cpu.DE.value = 0
	cpu.HL.value = 0
	cpu.SP.value = 0
	cpu.PC.value = 0
}
