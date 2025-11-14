package calculator

import "errors"

// ErrDivideByZero is returned when attempting to divide by zero
var ErrDivideByZero = errors.New("division by zero")