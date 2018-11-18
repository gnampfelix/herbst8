package hw_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHw(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Hw Suite")
}
