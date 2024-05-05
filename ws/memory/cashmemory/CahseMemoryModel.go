package cachememory

type CashMemory struct {
	data map[int]byte
}

type CacheMemoryRequest struct {
	address int
	value   byte
}
