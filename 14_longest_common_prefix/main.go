package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	var longestPrefix string
	longestPrefix, strs = strs[0], strs[1:]

	for _, v := range strs {
		strToLoop := longestPrefix
		strToCompare := v
		if (len(strToCompare) < len(strToLoop)) {
			strToLoop, strToCompare = strToCompare, strToLoop
		}

		prefixChars := ""
		for i, c := range []byte(strToLoop) {
			if (c != strToCompare[i]) {
				break;
			}

			prefixChars = prefixChars + string(c)
		}
		longestPrefix = prefixChars
	}

	return longestPrefix
}

func main() {
	result := longestCommonPrefix([]string{"flower", "flow", "flight"})
	fmt.Printf("RESULT 1: %s\n", result)

	result2 := longestCommonPrefix([]string{"dog", "racecar", "car"})
	fmt.Printf("RESULT 2: %s\n", result2)
}
