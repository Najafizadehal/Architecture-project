package memory

type Memory struct {
	Storage map[int][]byte
}

func NewMemory() *Memory {
	return &Memory{
		Storage: make(map[int][]byte),
	}
}

func (m *Memory) Store(address int, data []byte) {
	m.Storage[address] = data
}

func (m *Memory) Load(address int) []byte {
	data := m.Storage[address]

	return data
}
