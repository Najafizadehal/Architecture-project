package memory

import "fmt"

type Memory struct {
	Storage map[int][]byte
}

func NewMemory() *Memory {
	return &Memory{
		Storage: make(map[int][]byte, 4096),
	}
}

func (m *Memory) Store(address int, data []byte) {
	m.Storage[address] = data
}

func (m *Memory) Load(address int) ([]byte, error) {
	data, found := m.Storage[address]
	if !found {
		return nil, fmt.Errorf("data not found at address : %s", address)
	}
	return data, nil
}
