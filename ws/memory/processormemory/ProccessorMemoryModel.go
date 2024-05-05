package processormemory

type ProcessorMemory struct {
	data map[int]byte
}

type ProcessorMemoryRequest struct {
	address int
	value   byte
}
