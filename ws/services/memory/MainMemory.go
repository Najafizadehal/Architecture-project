package memory

import (
	"errors"
	"fmt"
)

// MemorySize defines the size of the memory.
const MemorySize = 4096

// Memory represents the main memory of the machine.
type Memory struct {
	Data [MemorySize]int
}

// NewMemory creates a new Memory with specified size.
func NewMemory() *Memory {
	return &Memory{}
}

// Read reads data from the specified address.
func (m *Memory) Read(address int) (int, error) {
	if address < 0 || address >= MemorySize {
		return 0, errors.New("address out of bounds")
	}
	return m.Data[address], nil
}

// Write writes data to the specified address.
func (m *Memory) Write(address int, value int) error {
	if address < 0 || address >= MemorySize {
		return errors.New("address out of bounds")
	}
	m.Data[address] = value
	return nil
}

// Dump prints the contents of memory for debugging.
func (m *Memory) Dump() {
	fmt.Println("Memory Dump:")
	for i, value := range m.Data {
		if value != 0 { // Print only non-zero memory locations for brevity
			fmt.Printf("Address %04X: %04X\n", i, value)
		}
	}
}
