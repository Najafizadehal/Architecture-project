package ws

type DataBus struct {
	Data []byte
}

func NewDataBus(size int) *DataBus {
	return &DataBus{
		Data: make([]byte, size),
	}
}

func (d *DataBus) Read() byte {
	return d.Data[0]
}

func (d *DataBus) Write(data byte) {
	d.Data[0] = data
}
