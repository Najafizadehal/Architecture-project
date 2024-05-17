package memory

import "errors"

type Registers struct {
	AR   int // Address Register
	PC   int // Program Counter
	DR   int // Data Register
	AC   int // Accumulator
	IR   int // Instruction Register
	TR   int // Temporary Register
	INPR int // Input Register
	OUTR int // Output Register
}

func NewRegister() *Registers {
	return &Registers{}
}

func (r *Registers) Set(register string, value int) error {
	switch register {
	case "AR":
		r.AR = value
	case "PC":
		r.PC = value
	case "DR":
		r.DR = value
	case "AC":
		r.AC = value
	case "IR":
		r.IR = value
	case "TR":
		r.TR = value
	case "INPR":
		r.INPR = value
	case "OUTR":
		r.OUTR = value
	default:
		return errors.New("invalid register")
	}

	return nil
}

func (r *Registers) Get(register string) (int, error) {

	switch register {
	case "AR":
		return r.AR, nil
	case "PC":
		return r.PC, nil
	case "DR":
		return r.DR, nil
	case "AC":
		return r.AC, nil
	case "IR":
		return r.IR, nil
	case "TR":
		return r.TR, nil
	case "INPR":
		return r.INPR, nil
	case "OUTR":
		return r.OUTR, nil
	default:
		return 0, errors.New("invalid register")
	}
}
