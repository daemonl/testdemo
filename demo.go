package demo

import "fmt"

type ArrayBuilder struct {
	NumberOfItems int
}

func (ab ArrayBuilder) BuildStringArray(input string) []string {
	out := make([]string, ab.NumberOfItems)
	for idx := range out {
		out[idx] = fmt.Sprintf("%d: %s", idx, input)
	}
	return out
}

func (ab ArrayBuilder) BuildIntArray(multiplier int) []int {
	out := make([]int, ab.NumberOfItems)
	for idx := range out {
		out[idx] = multiplier * idx
	}
	return out
}
