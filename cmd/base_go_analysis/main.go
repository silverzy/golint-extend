package main

import "fmt"

func main() {
	//singlechecker.Main(analyzer.Analyzer)
	fmt.Println(decodeString("2[abc]3[cd]ef"))
}
func decodeString(s string) string {

	var result string
	var stack []string
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "]" {
			var num int
			var temp string
			stack, temp = decodeUnit(stack)
			stack, num = decodeInt(stack)
			decodeVal := multiString(temp, num)
			stack = append(stack, decodeVal)
		} else {
			stack = append(stack, string(s[i]))
		}
	}
	for i := 0; i < len(stack); i++ {
		result += stack[i]
	}
	return result
}

func decodeUnit(stack []string) ([]string, string) {
	var result string
	for i := len(stack) - 1; i >= 0; i-- {
		if (stack)[i] == "[" {
			stack = stack[0:i]
			return stack, result
		}
		result = stack[i] + result
		stack = stack[0:i]
	}
	return stack, result
}

func decodeInt(stack []string) ([]string, int) {
	var num int
	for i := len(stack) - 1; i >= 0; i-- {
		if stack[i] < "9" && stack[i] > "0" {
			num = num*10 + int(stack[i][0]-'0')
			stack = stack[0:i]
		} else {
			return stack, num
		}
	}
	return stack, num
}

func multiString(temp string, num int) string {
	var result string
	for i := 0; i < num; i++ {
		result += temp
	}
	return result
}
