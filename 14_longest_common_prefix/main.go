package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	longestPrefix := strs[0]

	for _, v := range strs {
		i := 0
	
		for ; i < len(v) && i < len(longestPrefix) && v[i] == longestPrefix[i]; i++ {}
	
		longestPrefix = longestPrefix[:i]
	}

	return longestPrefix
}

func main() {
	result := longestCommonPrefix([]string{"flower", "flow", "flight"})
	fmt.Printf("RESULT 1: %s\n", result)

	result2 := longestCommonPrefix([]string{"dog", "racecar", "car"})
	fmt.Printf("RESULT 2: %s\n", result2)
}
