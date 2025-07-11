package main

import "fmt"

func main() {
	const currencyRateUsdEur = 0.85
	const currencyRateUsdRub = 77.9
	currencyRateEurRub := currencyRateUsdRub / currencyRateUsdEur
	fmt.Print(currencyRateEurRub)
}
