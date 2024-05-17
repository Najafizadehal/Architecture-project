package alu

// ALU represents the Arithmetic Logic Unit of the machine.
type ALU struct{}

// NewALU creates a new ALU.
func NewALU() *ALU {
	return &ALU{}
}

// Add performs addition on two integers.
func (alu *ALU) Add(a, b int) int {
	return a + b
}

// Subtract performs subtraction on two integers.
func (alu *ALU) Subtract(a, b int) int {
	return a - b
}

// And performs a bitwise AND operation on two integers.
func (alu *ALU) And(a, b int) int {
	return a & b
}

// Or performs a bitwise OR operation on two integers.
func (alu *ALU) Or(a, b int) int {
	return a | b
}

// Not performs a bitwise NOT operation on an integer.
func (alu *ALU) Not(a int) int {
	return ^a
}

// Xor performs a bitwise XOR operation on two integers.
func (alu *ALU) Xor(a, b int) int {
	return a ^ b
}

// import (
// 	"architecture/ws/services/bus"
// 	"architecture/ws/services/memory"
// )

// type ALU struct {
// 	Register *memory.Registers
// 	DataBus  *bus.DataBus
// }

// func NewALU(register *memory.Registers, dataBus *bus.DataBus) *ALU {
// 	return &ALU{
// 		Register: register,
// 		DataBus:  dataBus,
// 	}
// }

// func (a *ALU) Perform(operation byte) {
// 	operand1 := a.Register.ACC
// 	operand2 := a.DataBus.Read()

// 	switch operation {
// 	case 0x0:
// 		a.Register.ACC = operand1 + operand2
// 	case 0x1:
// 		a.Register.ACC = operand1 - operand2
// 	case 0x2:
// 		a.Register.ACC = operand1 * operand2
// 	case 0x3:
// 		a.Register.ACC = operand1 / operand2
// 	case 0x4:
// 		a.Register.ACC = operand1 & operand2
// 	case 0x5:
// 		a.Register.ACC = operand1 | operand2
// 	case 0x6:
// 		a.Register.ACC = operand1 ^ operand2
// 	case 0x7:
// 		a.Register.ACC = ^operand1
// 	}

// 	a.updateStatusRegister()
// }

// func (a *ALU) updateStatusRegister() {
// 	a.Register.Status = 0
// 	if a.Register.ACC == 0 {
// 		a.Register.Status = 0x01
// 	}
// 	if a.Register.ACC < 0 {
// 		a.Register.Status = 0x02
// 	}
// }
