package memory

import "fmt"

type Memory struct {
	Storage map[string][]byte
}

func NewMemory(size int) *Memory {
	return &Memory{
		Storage: make(map[string][]byte),
	}
}

func (m *Memory) Store(address []byte, data []byte) {
	addressKey := string(address)
	m.Storage[addressKey] = data
}

func (m *Memory) Load(address []byte) ([]byte, error) {
	addressKey := string(address)
	data, found := m.Storage[addressKey]
	if !found {
		return nil, fmt.Errorf("data not found at address : %s", addressKey)
	}
	return data, nil
}
