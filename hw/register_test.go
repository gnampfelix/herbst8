package hw_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/gnampfelix/herbst8/hw"
)

var _ = Describe("Register", func() {
	It("should create a register", func() {
		reg := NewRegister()
		bus := Bus(0)

		reg.Connect(&bus)
		bus.Write(128)
		reg.SetLoad(true)
		reg.Clock()
		reg.SetLoad(false)
		bus.Write(17)

		Expect(bus.Read()).Should(Equal(byte(17)))
		reg.Out()
		Expect(bus.Read()).Should(Equal(byte(128)))
	})
})
