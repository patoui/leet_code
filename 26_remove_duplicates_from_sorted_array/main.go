package main

import "fmt"

func removeDuplicates(nums []int) int {
	copy := nums
	nums = nums[:0]
	m := make(map[int]int)
	for i := 0; i < len(copy); i++ {
		if _, ok := m[copy[i]]; !ok {
			m[copy[i]] = 1
			nums = append(nums, copy[i])
		}
	}

	return len(nums)
}

func main() {
	fmt.Println("EXAMPLE 1")
	n1 := []int{1,1,2}
	e1 := []int{1,2}
	fmt.Printf("ORIGINAL: %v\n", n1)
	r1 := removeDuplicates(n1)
	fmt.Printf("RESULT 1: %d\n", r1)
	for i := 0; i < r1; i++ {
		if n1[i] == e1[i] {
			fmt.Printf("CORRECT AT %d OF %d\n", i, n1[i])
	    } else {
	    	fmt.Printf("INCORRECT AT %d OF %d, SHOULD BE: %d\n", i, n1[i], e1[i])
	    }
	}
	fmt.Printf("\n");

	fmt.Println("EXAMPLE 2")
	n2 := []int{0,0,1,1,1,2,2,3,3,4}
	e2 := []int{0,1,2,3,4}
	fmt.Printf("ORIGINAL: %v\n", n2)
	r2 := removeDuplicates(n2)
	fmt.Printf("RESULT 2: %d\n", r2)
	for i := 0; i < r2; i++ {
		if n2[i] == e2[i] {
			fmt.Printf("CORRECT AT %d OF %d\n", i, n2[i])
	    } else {
	    	fmt.Printf("INCORRECT AT %d OF %d, SHOULD BE: %d\n", i, n2[i], e2[i])
	    }
	}
	fmt.Printf("\n");

	fmt.Println("EXAMPLE 3")
	n3 := []int{1,1}
	e3 := []int{1}
	fmt.Printf("ORIGINAL: %v\n", n3)
	r3 := removeDuplicates(n3)
	fmt.Printf("RESULT 3: %d\n", r3)
	for i := 0; i < r3; i++ {
		if n3[i] == e3[i] {
			fmt.Printf("CORRECT AT %d OF %d\n", i, n3[i])
	    } else {
	    	fmt.Printf("INCORRECT AT %d OF %d, SHOULD BE: %d\n", i, n3[i], e3[i])
	    }
	}
	fmt.Printf("\n\n");
}
