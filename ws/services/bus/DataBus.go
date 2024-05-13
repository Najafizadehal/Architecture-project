package bus

type DataBus struct {
	Data map[int][]byte
}

func NewDataBus() *DataBus {
	return &DataBus{
		Data: make(map[int][]byte),
	}
}

func (d *DataBus) Read(address int) []byte {
	return d.Data[address]
}

func (d *DataBus) Write(address int, data []byte) {
	d.Data[address] = data
}
