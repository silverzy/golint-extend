package main

import "fmt"

func main() {
	//singlechecker.Main(analyzer.Analyzer)
	fmt.Println(lengthOfLongestSubstring("aabaab!bb"))
}
func lengthOfLongestSubstring(s string) int {
	lindex := 0
	rIndex := 0
	sMap := make(map[string]bool)
	max := 0
	for ; rIndex < len(s); rIndex++ {
		if !sMap[string(s[rIndex])] {
			sMap[string(s[rIndex])] = true
		} else {
			if len(sMap) > max {
				max = len(sMap)
			}
			for ; lindex <= rIndex; lindex++ {
				delete(sMap, string(s[lindex]))
				if !sMap[string(s[rIndex])] {
					sMap[string(s[rIndex])] = true
					lindex++
					break
				}
			}
		}
	}
	if len(sMap) > max {
		max = len(sMap)
	}
	return max
}
