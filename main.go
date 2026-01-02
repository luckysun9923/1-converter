package main

import "fmt"

func main() {
	var userText string = ""
	const usdEur = 0.854
	const usdRub = 80.72
	const eurRub = usdRub / usdEur
	result := userScan(userText)
	fmt.Println(result)
}

func userScan(userText string) string {
	fmt.Scan(&userText)
	return userText
}

func userCalculation(number int, usd string, rub string) int {

}
