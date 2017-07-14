package reverse_test

import (
	. "github.com/miguellgt/goprojects/tdd/reverse"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Reverse", func() {
	Context("Given input string Hello", func() {
		var inputString = "Hello"
		Context("When reverse function is called with the input", func() {
			var expectedResult = Reverse(inputString)
			It("Then should return string olleH", func() {
				Expect(expectedResult).To(Equal("olleH"))
			})
		})
	})
})
