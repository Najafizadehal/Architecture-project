package mainmemory

type MainMemory struct {
	data map[int]byte
}

type MainMemoryRequest struct {
	address int
	value   byte
}
