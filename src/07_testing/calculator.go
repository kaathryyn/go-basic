package calculator

// Calculator represents a simple calculator
type Calculator struct {
	memory float64
}

// Add adds two numbers
func (c *Calculator) Add(a, b float64) float64 {
	return a + b
}

// Subtract subtracts b from a
func (c *Calculator) Subtract(a, b float64) float64 {
	return a - b
}

// Multiply multiplies two numbers
func (c *Calculator) Multiply(a, b float64) float64 {
	return a * b
}

// Divide divides a by b
func (c *Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

// Memory returns the current value in memory
func (c *Calculator) Memory() float64 {
	return c.memory
}

// Store stores a value in memory
func (c *Calculator) Store(value float64) {
	c.memory = value
}

// Clear clears the memory
func (c *Calculator) Clear() {
	c.memory = 0
}