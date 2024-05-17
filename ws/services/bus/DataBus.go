package bus

import (
	"architecture/ws/services/memory"
	"errors"
	"fmt"
)

type DataBus struct {
	Address int
	Data    int
	Read    bool
	Write   bool
}

func NewDataBus() *DataBus {
	return &DataBus{}
}

func (bus *DataBus) SetAddress(address int) {
	bus.Address = address
}

func (bus *DataBus) SetData(data int) {
	bus.Data = data
}

func (bus *DataBus) EnableRead() {
	bus.Read = true
	bus.Write = false
}

func (bus *DataBus) EnableWrite() {
	bus.Write = true
	bus.Read = false
}

func (bus *DataBus) DisableSignals() {
	bus.Read = false
	bus.Write = false
}

func (bus *DataBus) PerformOperation(memory *memory.Memory, registers *memory.Registers) error {
	if bus.Read {
		data, err := memory.Read(bus.Address)
		if err != nil {
			return err
		}
		bus.Data = data
	} else if bus.Write {
		err := memory.Write(bus.Address, bus.Data)
		if err != nil {
			return err
		}
	} else {
		return errors.New("no operation specified on the data bus")
	}
	return nil
}
func (bus *DataBus) Dump() {
	fmt.Printf("Address: %04X, Data: %04X, Read: %t, Write: %t\n", bus.Address, bus.Data, bus.Read, bus.Write)
}
