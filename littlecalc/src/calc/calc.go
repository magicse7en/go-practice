package main

import "os"
import "fmt"
import "littlemath"
import "strconv"

func Usage() {
	fmt.Println("usage: calc <command> <integer1> <integer2>")
	fmt.Println("\nThe commands are:")
	fmt.Println("add\tAddition of two values")
	fmt.Println("sub\tSubtraction of two values")
	fmt.Println("mul\tMultiplication of two values")
	fmt.Println("div\tDivision of two values")
}

func main() {
	args := os.Args
	if args == nil || len(args) != 4 {
		Usage()
		return
	}

	v1, err1 := strconv.Atoi(args[2])
	v2, err2 := strconv.Atoi(args[3])
	if err1 != nil || err2 != nil {
		Usage()
		return
	}

	switch args[1] {
	case "add":
		ret := littlemath.Add(v1, v2)
		fmt.Println("Result: ", ret)
	case "sub":
		ret := littlemath.Sub(v1, v2)
		fmt.Println("Result: ", ret)
	case "mul":
		ret := littlemath.Mul(v1, v2)
		fmt.Println("Result: ", ret)
	case "div":
		ret, err := littlemath.Div(v1, v2)
		if err == nil {
			fmt.Println("Result: ", ret)
		}
	default:
		Usage()
	}
}
