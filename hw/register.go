package hw

type Register struct {
    data byte
    load bool
    reset bool
    bus *Bus
}

func NewRegister() Register {
    return Register{}
}

func (r *Register) SetLoad(load bool) {
    r.load = load
}

func (r *Register) Out() {
    r.bus.Write(r.data)
}

func (r *Register) Reset() {
    r.data = 0
}

func (r *Register) Connect(bus *Bus) {
    r.bus = bus
}

func (r *Register) Clock() {
    if r.load {
        r.data = r.bus.Read()
    }
}
