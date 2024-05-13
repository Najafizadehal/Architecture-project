package controlunit

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
