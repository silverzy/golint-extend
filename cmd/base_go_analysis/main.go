package main

import "fmt"

func main() {
	//singlechecker.Main(analyzer.Analyzer)
	fmt.Println(longestPalindrome("aabaab!bb"))
}
func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	initDP(dp, len(s), s)
	maxlength := 0
	lIndex := 0
	rIndex := 0
	for length := 2; length < len(s); length++ {
		for i := 0; i < len(s)-length; i++ {
			isTrue := dp[i+1][i+length-1] && s[i] == s[i+length]
			dp[i][i+length] = isTrue
			if isTrue && length > maxlength {
				maxlength = length
				lIndex = i
				rIndex = i + length
			}
		}
	}
	return s[lIndex : rIndex+1]
}

func initDP(dp [][]bool, length int, s string) {
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
		dp[i][i] = true
		if i+1 < length {
			dp[i][i+1] = s[i] == s[i+1]
		}
	}
}
