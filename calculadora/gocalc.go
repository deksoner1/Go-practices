//Example taken from https://gist.github.com/vikramdurai/cefe82a739a39ff23b4e3c361068a21c
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseArgs(c []string) (float64, float64, error) {
	num1, err := strconv.ParseFloat(c[0], 64)
	if err != nil {
		return 0.0, 0.0, err
	}
	num2, err := strconv.ParseFloat(c[2], 64)
	if err != nil {
		return 0.0, 0.0, err
	}
	return num1, num2, nil
}

func processStack(e []string) (float64, error) {
	fmt.Println(e)
	result := 0.0
	for i, v := range e {
		fmt.Println("This is v ", v, "this is i", i)
		c := strings.Split(v, " ")
		fmt.Println("this is the one ", c)
		if len(c)-1 < 2 {
			return 0.0, errors.New("error: some arguments are not supplied")
		}
		num1, num2, err := parseArgs(c)
		if err != nil {
			return 0.0, err
		}
		switch c[1] {
		case "*":
			result = num1 * num2
		case "/":
			if num2 == 0.0 {
				return 0.0, errors.New("error: you tried to divide by zero")
			}
			result = num1 / num2
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		}
	}
	return result, nil
}
func main() {
	expressions := make([]string, 1)
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("gocalc>")
		for scanner.Scan() {
			expressions = append(expressions, scanner.Text())
			fmt.Println(expressions)
			res, err := processStack(expressions)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(res)
			}
			fmt.Println("gocalc>")
		}
	}
}
