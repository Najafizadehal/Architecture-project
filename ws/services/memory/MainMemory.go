package memory

import (
	"errors"
	"fmt"
)

const MemorySize = 4096

type Memory struct {
	Data [MemorySize]int
}

func NewMemory() *Memory {
	return &Memory{}
}

func (m *Memory) Read(address int) (int, error) {
	if address < 0 || address >= MemorySize {
		return 0, errors.New("address out of bounds")
	}
	return m.Data[address], nil
}

func (m *Memory) Write(address int, value int) error {
	if address < 0 || address >= MemorySize {
		return errors.New("address out of bounds")
	}
	m.Data[address] = value
	return nil
}

func (m *Memory) Dump() {
	fmt.Println("Memory Dump:")
	for i, value := range m.Data {
		if value != 0 {
			fmt.Printf("Address %04X: %04X\n", i, value)
		}
	}
}
