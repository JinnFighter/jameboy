package cpu

type Register struct {
	value uint16
}

func (reg *Register) GetAs16() uint16 {
	return reg.value
}

func (reg *Register) GetAsHi8() uint8 {
	var res = uint8((reg.value & 0xFF00) >> 8)
	return res
}

func (reg *Register) GetAsLo8() uint8 {
	var res = uint8(reg.value & 0x00FF)
	return res
}
