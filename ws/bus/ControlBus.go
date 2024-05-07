package ws

type ControlBus struct {
	ControlSignals []byte
}

func NewControlBus(size int) *ControlBus {
	return &ControlBus{
		ControlSignals: make([]byte, size),
	}
}

func (b *ControlBus) ReadSignal(index int) byte {
	return b.ControlSignals[index]
}

func (b *ControlBus) WriteSignal(index int, signal byte) {
	b.ControlSignals[index] = signal
}
