package hw

type Bus byte

func (b Bus)Read() byte {
    return byte(b)
}

func (b *Bus)Write(data byte) {
    *b = Bus(data)
}
