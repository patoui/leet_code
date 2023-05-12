package main

import "fmt"

func removeElement(nums []int, val int) int {
    j := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != val {
            nums[j] = nums[i]
            j++
        }
    }

    return j
}

func main() {
	fmt.Println("EXAMPLE 1")
	n1 := []int{3,2,2,3}
	e1 := []int{2,2}
	fmt.Printf("ORIGINAL: %v\n", n1)
	r1 := removeElement(n1, 3)
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
	n2 := []int{0,1,2,2,3,0,4,2}
	e2 := []int{0,1,3,0,4}
	fmt.Printf("ORIGINAL: %v\n", n2)
	r2 := removeElement(n2, 2)
	fmt.Printf("PARSED: %v\n", n2)
	fmt.Printf("RESULT 2: %d\n", r2)
	for i := 0; i < r2; i++ {
		if n2[i] == e2[i] {
			fmt.Printf("CORRECT AT %d OF %d\n", i, n2[i])
	    } else {
	    	fmt.Printf("INCORRECT AT %d OF %d, SHOULD BE: %d\n", i, n2[i], e2[i])
	    }
	}
	fmt.Printf("\n");
}
