package secodarymemory

func (sm *SecodaryMemoryRequest) read(address int) byte {
	return sm.data[address]
}

func (sm *SecodaryMemoryRequest) write(address int, value byte) {
	sm.data[address] = value
}
