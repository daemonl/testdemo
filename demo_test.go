package demo

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Demo Outer Test", func() {
	var (

		// Configuration
		numberOfItems int

		// Build shared resources
		builder ArrayBuilder
	)

	BeforeEach(func() {
		fmt.Println("Top Level - BeforeEach")

		// Reset all parameters
		numberOfItems = 0
	})

	JustBeforeEach(func() {
		fmt.Println("Top Level - JustBeforeEach")

		// Build the shared resource
		builder = ArrayBuilder{
			NumberOfItems: numberOfItems,
		}
	})

	Context("String Builder", func() {
		var (
			// Input Before Running
			input string

			// Output after running
			output []string
		)
		BeforeEach(func() {
			fmt.Println("String Builder - BeforeEach")
			// Reset all parameters
			input = ""
			output = nil
		})
		JustBeforeEach(func() {
			fmt.Println("String Builder - JustBeforeEach")
			output = builder.BuildStringArray(input)
		})
		Context("Empty", func() {
			BeforeEach(func() {
				fmt.Println("String Builder: Empty - Before Each")
				numberOfItems = 0
				input = "hello"
			})
			It("Should have no items", func() {
				Expect(output).To(HaveLen(0))
			})
		})
		Context("One Item", func() {
			BeforeEach(func() {
				fmt.Println("String Builder: One Item - Before Each")
				numberOfItems = 1
				input = "hello"
			})
			It("Should have one item", func() {
				Expect(output).To(HaveLen(1))
			})
			It("Should format the item correctly", func() {
				Expect(output[0]).To(Equal("0: hello"))
			})
		})
	})

	Context("Int Builder", func() {
		var (
			// Input Before Running
			input int
			// Output after running
			output []int
		)
		BeforeEach(func() {
			fmt.Println("Int Builder - BeforeEach")
			// Reset all parameters
			input = 0
			output = nil
		})
		JustBeforeEach(func() {
			fmt.Println("Int Builder - JustBeforeEach")
			output = builder.BuildIntArray(input)
		})
		Context("Empty", func() {
			BeforeEach(func() {
				numberOfItems = 0
				input = 10
			})
			It("Should have no items", func() {
				Expect(output).To(HaveLen(0))
			})
		})
		Context("Two Items", func() {
			BeforeEach(func() {
				numberOfItems = 2
				input = 10
			})
			It("Should have two items", func() {
				Expect(output).To(HaveLen(2))
			})
			It("The first item should always be 0", func() {
				Expect(output[0]).To(Equal(0))
			})
			It("Correctly Multiplies", func() {
				Expect(output[1]).To(Equal(10))
			})
		})

	})
})
