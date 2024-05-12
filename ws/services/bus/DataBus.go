package bus

type DataBus struct {
	Data map[int]byte
}

func NewDataBus(size int) *DataBus {
	data := make(map[int]byte, size)
	for i := 0; i < size; i++ {
		data[i] = 0
	}
	return &DataBus{
		Data: data,
	}
}

func (d *DataBus) Read() byte {
	return d.Data[0]
}

func (d *DataBus) Write(data byte) {
	d.Data[0] = data
}
