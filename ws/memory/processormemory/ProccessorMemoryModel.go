package processormemory

type ProcessorMemory struct {
	data []byte
}

type ProcessorMemoryRequest struct {
	address int
	value   byte
}
