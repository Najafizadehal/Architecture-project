package memory

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
	// lineIndex := address / c.LineSize
	offset := address & c.LineSize

	for _, line := range c.CacheLines {
		if line.Tag == address/c.LineSize {
			return line.Data[offset]
		}
	}

	return 0
}

func (c *Cache) Write(address int, data byte) {
	lineIndex := address / c.LineSize
	offset := address % c.LineSize

	for i, line := range c.CacheLines {
		if line.Tag == address/c.LineSize {
			c.CacheLines[i].Data[offset] = data
			return
		}
	}
	// در صورت عدم یافتن در کش، افزودن خط جدید به کش
	c.CacheLines[lineIndex] = CacheLine{
		Tag:  address / c.LineSize,
		Data: make([]byte, c.LineSize),
	}

	c.CacheLines[lineIndex].Data[offset] = data
}
