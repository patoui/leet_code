package main

import "fmt"

func isValid(s string) bool {
	if len(s) < 2 {
		return false;
	}

	var cop rune
	var op []rune
	cm := map[rune]rune{'[': ']', '(': ')', '{': '}'}

	for _, v := range s {
		if _, ok := cm[v]; ok {
			op = append(op, v)
		} else {
			if len(op) == 0 {
				return false
			}

			cop, op = op[len(op)-1], op[:len(op)-1]
			c, ok := cm[cop]

			if !ok || v != c {
				return false
			}
		}
	}

	return len(op) == 0
}

func main() {
	fmt.Println("EXAMPLE 1")
	str1 := "()"
	result1 := isValid(str1)
	fmt.Printf("%t\n", result1)

	fmt.Println("EXAMPLE 2")
	str2 := "()[]{}"
	result2 := isValid(str2)
	fmt.Printf("%t\n", result2)

	fmt.Println("EXAMPLE 3")
	str3 := "(]"
	result3 := isValid(str3)
	fmt.Printf("%t\n", result3)
}

