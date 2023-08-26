# Learn GoLang in 2023

**Reference(s):**

> 1. [https://gobyexample.com](https://gobyexample.com)

```text
https://go.dev/tour/basics/3
https://gobyexample.com/values
```

## Table of Contents

Few basic commands to get started with GoLang

## Few basic commands to get started with GoLang

```bash
go work init # Inside the parent folder

go mod init basics/A1HelloGophers
go mod tidy
go work use .
go build .
go run .

go fmt
go work sync
```

## Few basic points to remember

> 1. Short hand declaration can be used to declare and initialize a variable. It is allowed only inside a function.
> 1. A numeric constant has no type until it’s given one, such as by an explicit conversion.
> 1. A number can be given a type by using it in a context that requires one, such as a variable assignment or function call. For example, here math.Sin expects a float64.
> 1. Value that is float32 and cannot be stored in a variable that is declared to hold a float64
> 1. A statement can precede conditionals; any variables declared in this statement are available in the current and all subsequent branches.
> 1. Go has only one looping construct, the for loop.
> 1. "factored" import statement with multiple imports within parenthesis
> 1. A function can return any number of results.
> 1. A return statement without arguments returns the named return values. This is known as a "naked" return.
> 1. An uninitialized slice equals to nil and has length 0.
> 1. The range form of the for loop iterates over a slice or map.
