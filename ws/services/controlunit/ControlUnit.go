package controlunit

import (
	"architecture/ws/services/alu"
	"architecture/ws/services/bus"
	"architecture/ws/services/memory"
	"errors"
	"fmt"
)

type ControlUnit struct {
	Registers *memory.Registers
	Memory    *memory.Memory
	ALU       *alu.ALU
	Bus       *bus.DataBus
}
type Memory struct {
	Data map[int]int
}

func NewMemory() *Memory {
	return &Memory{
		Data: make(map[int]int),
	}
}

func NewControlUnit(bus *bus.DataBus, registers *memory.Registers, memory *memory.Memory, alu *alu.ALU) *ControlUnit {
	return &ControlUnit{
		Registers: registers,
		Memory:    memory,
		ALU:       alu,
		Bus:       bus,
	}
}

func (cu *ControlUnit) Fetch() error {
	cu.Bus.SetAddress(cu.Registers.PC)
	cu.Bus.EnableRead()
	err := cu.Bus.PerformOperation(cu.Memory, cu.Registers)
	if err != nil {
		return err
	}
	cu.Registers.IR = cu.Bus.Data
	cu.Registers.PC++
	cu.Bus.DisableSignals()
	return nil
}
func (cu *ControlUnit) Decode() (string, int, error) {
	opcode := (cu.Registers.IR & 0xF000) >> 12
	address := cu.Registers.IR & 0x0FFF
	return fmt.Sprintf("%X", opcode), address, nil
}

func (cu *ControlUnit) Execute(opcode string, address int) error {
	switch opcode {
	case "1": // ADD
		// Set address on bus
		cu.Bus.SetAddress(address)
		// Enable read signal on bus
		cu.Bus.EnableRead()
		// Perform read operation
		err := cu.Bus.PerformOperation(cu.Memory, cu.Registers)
		if err != nil {
			return err
		}
		// Store read data in DR
		cu.Registers.DR = cu.Bus.Data
		// Perform addition using ALU
		cu.Registers.AC = cu.ALU.Add(cu.Registers.AC, cu.Registers.DR)
	case "2": // SUB
		// Set address on bus
		cu.Bus.SetAddress(address)
		// Enable read signal on bus
		cu.Bus.EnableRead()
		// Perform read operation
		err := cu.Bus.PerformOperation(cu.Memory, cu.Registers)
		if err != nil {
			return err
		}
		// Store read data in DR
		cu.Registers.DR = cu.Bus.Data
		// Perform subtraction using ALU
		cu.Registers.AC = cu.ALU.Subtract(cu.Registers.AC, cu.Registers.DR)
	case "3": // LDA
		// Set address on bus
		cu.Bus.SetAddress(address)
		// Enable read signal on bus
		cu.Bus.EnableRead()
		// Perform read operation
		err := cu.Bus.PerformOperation(cu.Memory, cu.Registers)
		if err != nil {
			return err
		}
		// Load data into AC
		cu.Registers.AC = cu.Bus.Data
	case "4": // STA
		// Set address and data on bus
		cu.Bus.SetAddress(address)
		cu.Bus.SetData(cu.Registers.AC)
		// Enable write signal on bus
		cu.Bus.EnableWrite()
		// Perform write operation
		err := cu.Bus.PerformOperation(cu.Memory, cu.Registers)
		if err != nil {
			return err
		}
	case "5": // BUN (Branch Unconditionally)
		cu.Registers.PC = address
	case "6": // BSA (Branch and Save Return Address)
		// Set address and data on bus
		cu.Bus.SetAddress(address)
		cu.Bus.SetData(cu.Registers.PC)
		// Enable write signal on bus
		cu.Bus.EnableWrite()
		// Perform write operation
		err := cu.Bus.PerformOperation(cu.Memory, cu.Registers)
		if err != nil {
			return err
		}
		// Set PC to address + 1
		cu.Registers.PC = address + 1
	case "7": // ISZ (Increment and Skip if Zero)
		// Set address on bus
		cu.Bus.SetAddress(address)
		// Enable read signal on bus
		cu.Bus.EnableRead()
		// Perform read operation
		err := cu.Bus.PerformOperation(cu.Memory, cu.Registers)
		if err != nil {
			return err
		}
		// Increment read value
		cu.Bus.Data++
		// Set address and data on bus
		cu.Bus.SetAddress(address)
		cu.Bus.SetData(cu.Bus.Data)
		// Enable write signal on bus
		cu.Bus.EnableWrite()
		// Perform write operation
		err = cu.Bus.PerformOperation(cu.Memory, cu.Registers)
		if err != nil {
			return err
		}
		// Skip next instruction if zero
		if cu.Bus.Data == 0 {
			cu.Registers.PC++
		}
	case "8": // AND
		// Set address on bus
		cu.Bus.SetAddress(address)
		// Enable read signal on bus
		cu.Bus.EnableRead()
		// Perform read operation
		err := cu.Bus.PerformOperation(cu.Memory, cu.Registers)
		if err != nil {
			return err
		}
		// Store read data in DR
		cu.Registers.DR = cu.Bus.Data
		// Perform AND operation using ALU
		cu.Registers.AC = cu.ALU.And(cu.Registers.AC, cu.Registers.DR)
	case "9": // OR
		// Set address on bus
		cu.Bus.SetAddress(address)
		// Enable read signal on bus
		cu.Bus.EnableRead()
		// Perform read operation
		err := cu.Bus.PerformOperation(cu.Memory, cu.Registers)
		if err != nil {
			return err
		}
		// Store read data in DR
		cu.Registers.DR = cu.Bus.Data
		// Perform OR operation using ALU
		cu.Registers.AC = cu.ALU.Or(cu.Registers.AC, cu.Registers.DR)
	case "A": // XOR
		// Set address on bus
		cu.Bus.SetAddress(address)
		// Enable read signal on bus
		cu.Bus.EnableRead()
		// Perform read operation
		err := cu.Bus.PerformOperation(cu.Memory, cu.Registers)
		if err != nil {
			return err
		}
		// Store read data in DR
		cu.Registers.DR = cu.Bus.Data
		// Perform XOR operation using ALU
		cu.Registers.AC = cu.ALU.Xor(cu.Registers.AC, cu.Registers.DR)
	case "B": // NOT
		// Perform NOT operation using ALU
		cu.Registers.AC = cu.ALU.Not(cu.Registers.AC)
	default:
		return errors.New("unsupported opcode: " + opcode)
	}
	// Disable signals on bus
	cu.Bus.DisableSignals()
	return nil
}

func (cu *ControlUnit) RunCycle() error {
	err := cu.Fetch()
	if err != nil {
		return err
	}
	opcode, address, err := cu.Decode()
	if err != nil {
		return err
	}
	return cu.Execute(opcode, address)
}

// import (
// 	"architecture/ws/services/bus"
// 	"architecture/ws/services/memory"
// 	ws "architecture/ws/services/memory"
// )

// type ControlUnit struct {
// 	Register    *ws.Registers
// 	AddressBus  *bus.AddressBus
// 	DataBus     *bus.DataBus
// 	ControlBus  *bus.ControlBus
// 	MainMemory  *memory.Memory
// 	CacheMemory *memory.Cache
// }

// func NewControlUnit(register *ws.Registers, adrressBus *bus.AddressBus, dataBus *bus.DataBus, controlBus *bus.ControlBus, mainMemory *memory.Memory, cacheMemory *memory.Cache) *ControlUnit {
// 	return &ControlUnit{
// 		Register:    register,
// 		AddressBus:  adrressBus,
// 		DataBus:     dataBus,
// 		ControlBus:  controlBus,
// 		MainMemory:  mainMemory,
// 		CacheMemory: cacheMemory,
// 	}
// }

// func (cu *ControlUnit) FetchIntstruction() {
// 	cu.AddressBus.Write(cu.Register.PC)
// 	instruction := cu.MainMemory.Read(byte(cu.Register.PC))
// 	cu.Register.IR = instruction
// 	cu.Register.PC++
// }

// func (cu *ControlUnit) DecodeInstruction() {
// 	opcode := cu.Register.IR >> 4
// 	operand := cu.Register.IR & 0x0f

// 	switch opcode {
// 	case 0x0:
// 		cu.Register.MAR = int(operand)
// 		cu.ControlBus.WriteSignal(0, 1)
// 	case 0x1:
// 		cu.Register.MAR = int(operand)
// 		cu.ControlBus.WriteSignal(1, 1)
// 	case 0x2:
// 		cu.Register.ACC += cu.MainMemory.Read(byte(operand))
// 	case 0x3:
// 		cu.Register.ACC -= cu.MainMemory.Read(byte(operand))
// 		///// بعدا پیاده سازی کن بقیه دستورات رو

// 	}
// }

// func (cu *ControlUnit) Execute() {
// 	switch cu.ControlBus.ReadSignal(0) {
// 	case 1:
// 		// خواندن حافظه و ذخیره در MBR
// 		value := cu.CacheMemory.Read(cu.Register.MAR)
// 		cu.Register.MBR = value
// 		// cu.DataBus.Read()
// 		// cu.Register.MBR = cu.DataBus.Read()
// 	case 2:
// 		// نوشتن در حافظه از MBR
// 		cu.DataBus.Write(cu.Register.MBR)
// 	}
// 	// پاک کردن سیگنال های کنترلی
// 	cu.ControlBus.WriteSignal(0, 0)
// 	cu.ControlBus.WriteSignal(1, 0)
// }
