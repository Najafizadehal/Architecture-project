package controlunit

import (
	"architecture/ws/services/memory"
	"errors"
	"fmt"
)

type ControlUnit struct {
	Register *memory.Registers
	Memory   *Memory
}
type Memory struct {
	Data map[int]int
}

func NewMemory() *Memory {
	return &Memory{
		Data: make(map[int]int),
	}
}

func NewControlUnit(register *memory.Registers, memory *Memory) *ControlUnit {
	return &ControlUnit{
		Register: register,
		Memory:   memory,
	}
}

func (cu *ControlUnit) Fetch() {
	cu.Register.IR = cu.Memory.Data[cu.Register.PC]
	cu.Register.PC++
}

func (cu *ControlUnit) Decode() (string, int, error) {
	opcode := (cu.Register.IR & 0xF000) >> 12
	address := cu.Register.IR & 0x0FFF
	return fmt.Sprintf("%X", opcode), address, nil
}

func (cu *ControlUnit) Execute(opcode string, address int) error {
	switch opcode {
	case "1":
		cu.Register.DR = cu.Memory.Data[address]
		cu.Register.AC += cu.Register.DR
	case "2":
		cu.Register.DR = cu.Memory.Data[address]
		cu.Register.AC -= cu.Register.DR
	case "3":
		cu.Register.AC = cu.Memory.Data[address]
	case "4":
		cu.Memory.Data[address] = cu.Register.AC
	case "5":
		cu.Register.PC = address
	case "6":
		cu.Memory.Data[address] = cu.Register.PC
		cu.Register.PC = address + 1
	case "7":
		cu.Memory.Data[address]++
		if cu.Memory.Data[address] == 0 {
			cu.Register.PC++
		}
	default:
		return errors.New("unsupported opcode: " + opcode)
	}
	return nil
}

func (cu *ControlUnit) RunCycle() error {
	cu.Fetch()
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
