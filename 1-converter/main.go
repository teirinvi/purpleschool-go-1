package main

import "fmt"

const currencyRateUsdEur = 0.85
const currencyRateUsdRub = 77.9

func main() {
	currencyRateEurRub := currencyRateUsdRub / currencyRateUsdEur
	fmt.Println(currencyRateEurRub)
	userValue, userCurrencyInput, userCurrencyOutput := getUserInput()
	getCurrencyValue(userValue, userCurrencyInput, userCurrencyOutput)
}

func getUserInput() (float64, string, string) {
	var userValue float64
	var userCurrencyInput string
	var userCurrencyOutput string

	fmt.Println("Введите количество валюты: ")
	fmt.Scan(&userValue)
	fmt.Println("Введите исходную валюту: ")
	fmt.Scan(&userCurrencyInput)
	fmt.Println("Введите целевую валюту: ")
	fmt.Scan(&userCurrencyOutput)

	return userValue, userCurrencyInput, userCurrencyOutput
}

func getCurrencyValue(userValue float64, userInput string, userOutput string) {
	fmt.Println(userValue)
	fmt.Println(userInput)
	fmt.Println(userOutput)
}
