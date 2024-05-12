package memory

type Registers struct {
	PC     int
	ACC    byte
	IR     byte
	MAR    int
	MBR    byte
	Status byte
}

func NewRegister() *Registers {
	return &Registers{}
}

func (r *Registers) Read(reg int) byte {
	switch reg {
	case 0:
		return byte(r.PC)
	case 1:
		return r.ACC
	case 2:
		return r.IR
	case 3:
		return byte(r.MAR)
	case 4:
		return r.MBR
	case 5:
		return r.Status
	default:
		return 0
	}
}

func (r *Registers) Write(reg int, value byte) {
	switch reg {
	case 0:
		r.PC = int(value)
	case 1:
		r.ACC = value
	case 2:
		r.IR = value
	case 3:
		r.MAR = int(value)
	case 4:
		r.MBR = value
	case 5:
		r.Status = value
	}
}
