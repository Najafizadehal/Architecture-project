package processormemory

func (pm *ProcessorMemoryRequest) read(address int) byte {
	return pm.data[address]
}

func (pm *ProcessorMemoryRequest) write(address int, value byte) {
	pm.data[address] = value
}
