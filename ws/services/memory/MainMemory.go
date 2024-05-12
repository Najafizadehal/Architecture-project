package memory

type Memory struct {
	data []byte
	size int
}

func NewMemory(size int) *Memory {
	return &Memory{
		data: make([]byte, size),
		size: size,
	}
}

func (m *Memory) Read(address byte) byte {
	return m.data[address]
}

func (m *Memory) Write(address byte, value byte) {
	m.data[address] = value
}
