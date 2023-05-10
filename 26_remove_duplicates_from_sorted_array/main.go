package main

import "fmt"

func removeDuplicates(nums []int) int {
	for i := 0; i < len(nums); i++ {
		dups := 0
		for j := i+1; j < len(nums); j++ {
			if nums[i] != nums[j] {
				break;
			}
			dups++
		}

		if dups > 0 {
			nums = append(nums[:i+1], nums[i+dups+1:]...)
		}
	}

	return len(nums)
}

func main() {
	fmt.Println("EXAMPLE 1")
	n1 := []int{1,1,2}
	e1 := []int{1,2}
	r1 := removeDuplicates(n1)
	fmt.Printf("RESULT 1: %d\n", r1)
	for i := 0; i < r1; i++ {
		if n1[i] == e1[i] {
			fmt.Printf("CORRECT AT %d OF %d\n", i, n1[i])
	    } else {
	    	fmt.Printf("INCORRECT AT %d OF %d, SHOULD BE: %d\n", i, n1[i], e1[i])
	    }
	}
	fmt.Printf("\n\n");

	fmt.Println("EXAMPLE 2")
	n2 := []int{0,0,1,1,1,2,2,3,3,4}
	e2 := []int{0,1,2,3,4}
	r2 := removeDuplicates(n2)
	fmt.Printf("RESULT 2: %d\n", r2)
	for i := 0; i < r2; i++ {
		if n2[i] == e2[i] {
			fmt.Printf("CORRECT AT %d OF %d\n", i, n2[i])
	    } else {
	    	fmt.Printf("INCORRECT AT %d OF %d, SHOULD BE: %d\n", i, n2[i], e2[i])
	    }
	}
	fmt.Printf("\n\n");

	fmt.Println("EXAMPLE 3")
	n3 := []int{1,1}
	e3 := []int{1}
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
