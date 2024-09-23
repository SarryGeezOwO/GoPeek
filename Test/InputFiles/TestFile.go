package main

import (
    "errors"
)

/*
    Foo, lorem lexum
    something cool
    testing multi-line comment
*/

// Add returns the sum of two integers.
func Add(a int, b int) int {
    return a + b
}

// Subtract returns the difference of two integers.
// It returns an error if the result is negative.
func Subtract(a int, b int) (int, error) {
    if a < b {
        return 0, errors.New("result is negative")
    }
    return a - b, nil
}

// Multiply returns the product of two integers.
func Multiply(a int, b int) int {
    return a * b
}

// Divide returns the quotient of two integers.
// It returns an error if the divisor is zero.
func Divide(a int, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}
