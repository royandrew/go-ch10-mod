// This package provide functions to perform mathematical operations
package goch10mod

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add function adds two integers together and returns the result.
// See [Addition]
// [Addition]: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](v1 T, v2 T) T {
	return v1 + v2
}
