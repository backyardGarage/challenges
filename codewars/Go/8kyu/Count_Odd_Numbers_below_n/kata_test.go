package kata_test

import (
	. "../Count_Odd_Numbers_below_n"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sample Test Cases", func() {
	It("should handle sample Test Cases", func() {
		Expect(OddCount(15)).To(Equal(7))
		Expect(OddCount(15023)).To(Equal(7511))
	})
})
