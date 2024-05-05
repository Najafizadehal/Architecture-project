package cashmemory

func (cm *CashMemoryRequest) read(address int) byte {
	if value, ok := cm.data[address]; ok {
		return value
	}
	return 0
}

func (cm *CashMemoryRequest) write(address int, value byte) {
	cm.data[address] = value
}
