// Package gomodexample has been created to practice go module releases
package gomodexample

// Add returns the sum of given arguments.
// Works on integers and floats
//
// read more from [mathisfun]
//
// [mathisfun]: https://mathisfun.com/numbers/addition
func Add[T int | float64](x, y T) T {
	return x + y
}
