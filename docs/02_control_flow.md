# Control Flow in Go

Control flow structures in Go allow you to control the execution path of your program. Here are the main control flow constructs in Go.

## If Statements

### Basic If-Else
```go
if condition {
    // code for true condition
} else {
    // code for false condition
}
```

### If with Short Statement
Go allows you to execute a short statement before the condition:
```go
if x := someFunction(); x < 0 {
    return -x
}
```

## Loops (For)

Go has only one looping construct: `for`. However, it's flexible and can be used in several ways.

### Basic For Loop
```go
for i := 0; i < 10; i++ {
    // code to repeat
}
```

### For as While
```go
for condition {
    // code to repeat while condition is true
}
```

### Infinite Loop
```go
for {
    // code to repeat forever
    // use break to exit
}
```

### Range-based For Loop
```go
for index, value := range collection {
    // code using index and value
}
```

## Switch Statements

### Basic Switch
```go
switch variable {
case value1:
    // code for value1
case value2, value3:
    // code for value2 or value3
default:
    // default code
}
```

### Switch with No Expression
```go
switch {
case condition1:
    // code for condition1
case condition2:
    // code for condition2
default:
    // default code
}
```

## Key Points

1. **Break and Continue**
   - `break`: Exits the innermost loop or switch
   - `continue`: Skips to the next iteration of the loop

2. **No Parentheses Required**
   - Go doesn't require parentheses around conditions
   - But requires curly braces

3. **Automatic Break**
   - Switch cases break automatically in Go
   - No fall-through by default
   - Use `fallthrough` keyword if needed

4. **Short Circuit Evaluation**
   - `&&` and `||` operators short circuit
   - `a && b`: if `a` is false, `b` is not evaluated
   - `a || b`: if `a` is true, `b` is not evaluated

## Best Practices

1. Keep conditions simple and readable
2. Use early returns when appropriate
3. Prefer positive conditions over negative ones
4. Use switch statements instead of long if-else chains
5. Be careful with infinite loops
6. Always include a default case in switch statements when appropriate