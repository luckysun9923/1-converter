package main

import (
	"fmt"
	"go/types"
)

func main() {
	var userText string
	var targetText string
	const usdEur = 0.854
	const usdRub = 80.72
	const eurRub = usdRub / usdEur
	var userNumber int
	fmt.Scan(&userNumber)
	result := userScan(userText)
	fmt.Println(result)
	resultCurrency := sourceCurrency(userText)
	fmt.Println(resultCurrency)
	resultNumber := enterNumber(userNumber)
	fmt.Println(resultNumber)
	resultTargetCurrency := targetCurrency(userText, targetText)
	fmt.Println(resultTargetCurrency)
	resultConvertervalute := userCalculation(userNumber, userText, targetText)
	fmt.Println(resultConvertervalute)

}

func userScan(userText string) string {
	fmt.Scan(&userText)
	return userText
}

func sourceCurrency(userText string) string {
	for {
		fmt.Println("Введите валюту : USD/EUR/RUB ")
		fmt.Scan(&userText)
		if userText == "USD" || userText == "EUR" || userText == "RUB" {
			return userText
		} else {
			fmt.Println("Неверно выбранная валюта!!!")
			continue
		}
	}
}

func enterNumber(userNumber int) int {
	for {
		fmt.Println("Введите число :")
		fmt.Scan(&userNumber)
		if userNumber == int(types.Int32) {
			return userNumber
		} else {
			fmt.Println("Неверно введено число")
			continue
		}
	}
}

func targetCurrency(userText string, targetText string) string {
	for {
		fmt.Println("Введите целевую валюту : USD/EUR/RUB ")
		fmt.Scan(&userText)
		fmt.Scan(&targetText)
		if userText != targetText {
			return targetText
		} else {
			fmt.Println("Неверная целевая валюта")
			continue
		}
	}
}

func userCalculation(userNumber int, userText string, targetText string) (int, float64) {
	if userText == "RUB" && targetText == "USD" {
		return resultConvertervalute == userNumber/usdRub
	} else if userText == "USD" && targetText == "RUB" {
		return resultConvertervalute == userNumber*usdRub
	} else if userText == "USD" && targetText == "EUR" {
		return resultConvertervalute == userNumber*usdEur
	} else if userText == "EUR" && targetText == "USD" {
		return resultConvertervalute == userNumber/usdEur
	} else if userText == "RUB" && targetText == "EUR" {
		return resultConvertervalute == userNumber*eurRub
	} else {
		return resultConvertervalute == userNumber/eurRub
	}
}
