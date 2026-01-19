package main

import (
	"fmt"
)

const usdEur = 0.854
const usdRub = 80.72
const eurRub = usdRub / usdEur

func main() {
	userText := sourceCurrency()
	userNumber := enterNumber()
	targetText := targetCurrency(userText)
	result := userCalculation(userNumber, targetText, userText)
	fmt.Printf("Результат: %.2f %s\n", result, targetText)
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

func enterNumber(userNumber float64) float64 {
	for {
		fmt.Println("Введите число :")
		fmt.Scan(&userNumber)
		if userNumber >= 0 {
			return userNumber
		} else {
			fmt.Println("Неверно введено число")
			continue
		}
	}
}

func targetCurrency(sourceText string) string {
	for {
		fmt.Println("Введите целевую валюту : USD/EUR/RUB ")
		var targetText string
		fmt.Scan(&targetText)
		if (targetText == "USD" || targetText == "EUR" || targetText == "RUB") && targetText != sourceText {
			return targetText
		}
		fmt.Println("Неверная целевая валюта")
	}
}

func userCalculation(userNumber float64, userText string, targetText string) float64 {
	if userText == "RUB" && targetText == "USD" {
		return userNumber / usdRub
	} else if userText == "USD" && targetText == "RUB" {
		return userNumber * usdRub
	} else if userText == "USD" && targetText == "EUR" {
		return userNumber * usdEur
	} else if userText == "EUR" && targetText == "USD" {
		return userNumber / usdEur
	} else if userText == "RUB" && targetText == "EUR" {
		return userNumber / eurRub
	} else {
		return userNumber * eurRub
	}
}
