package ws

type Cache struct {
	CacheLines []CacheLine
	Size       int
	LineSize   int
}

type CacheLine struct {
	Tag  int
	Data []byte
}

func NewCache(size, lineSize int) *Cache {
	return &Cache{
		CacheLines: make([]CacheLine, size/lineSize),
		Size:       size,
		LineSize:   lineSize,
	}
}

func (c *Cache) Read(address int) byte {
	lineIndex := address / c.LineSize
	offset := address / c.LineSize

	for _, line := range c.CacheLines {
		if line.Tag == address/c.LineSize {
			return line.Data[offset]
		}
	}

	return 0
}
