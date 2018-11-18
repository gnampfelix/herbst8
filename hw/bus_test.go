package hw_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/gnampfelix/herbst8/hw"
)

var _ = Describe("Bus", func() {
	It("should write to the bus", func() {
		bus := Bus(0)
		bus.Write(7)
		data := bus.Read()
		Expect(data).Should(Equal(byte(7)))
	})
})
