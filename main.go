package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var incomeJt float64 = 500
	var investmentReturnRate float64 = 0.135
	var inflationRate float64 = 0.035
	var interestAdjustedRate float64 = investmentReturnRate - inflationRate
	var years int = 5
	balance := incomeJt * float64(years*12)
	invested, _ := calcInvestment(incomeJt, investmentReturnRate, years)
	actualInvested, _ := calcInvestment(incomeJt, interestAdjustedRate, years)
	actualBalance, _ := calcInflation(incomeJt, inflationRate, years)
	invested = math.Round(invested*100) / 100
	actualBalance = math.Round(actualBalance*100) / 100
	actualInvested = math.Round(actualInvested*100) / 100
	difference := math.Round((actualInvested-actualBalance)*100) / 100
	fmt.Println("Saving " + strconv.FormatFloat(incomeJt, 'g', -1, 64) + " every month for " + strconv.Itoa(years) + " years")
	fmt.Println("Investment return rate " + strconv.FormatFloat(investmentReturnRate*100, 'g', -1, 64) + "%, inflation rate " + strconv.FormatFloat(inflationRate*100, 'g', -1, 64) + "%")
	fmt.Print("final balance: ")
	fmt.Print(balance)
	fmt.Print(" adjusted to inflation: ")
	fmt.Println(actualBalance)
	fmt.Print("balance if invested: ")
	fmt.Print(invested)
	fmt.Print(" adjusted to inflation: ")
	fmt.Println(actualInvested)
	fmt.Print("adjusted difference: ")
	fmt.Println(difference)
	return
}

func calcInvestment(income float64, yearlyInterestPercent float64, targetYears int) (float64, error) {
	var pool float64
	targetMonths := targetYears * 12
	monthlyInterestPercent := findMonthlyPercent(yearlyInterestPercent)
	for i := targetMonths; i > 0; i-- {
		pool = pool + income
		pool = pool + (pool * monthlyInterestPercent)
	}
	return pool, nil
}

func calcInflation(income float64, yearlyInflationPercent float64, targetYears int) (float64, error) {
	var pool float64
	yearlyIncome := income * 12
	totalInflationPercent := 1 + yearlyInflationPercent
	reverseInflationPercent := 1 / totalInflationPercent
	for i := targetYears; i > 0; i-- {
		pool = pool + yearlyIncome
		pool = pool * reverseInflationPercent
		yearlyIncome = yearlyIncome * reverseInflationPercent // adjust next year income
	}
	return pool, nil
}

// findMonthlyPercent converts target yearly interest to monthly interest in percent
// returns monthly interest
func findMonthlyPercent(targetYearlyPercent float64) float64 {
	topLimit := targetYearlyPercent // max possible monthly percent
	var lowLimit float64 = 0        // min possible monthly percent
	var thisMonthPercent float64    // current guess
	const initMoney float64 = 100
	targetMoney := (initMoney * 12) + (initMoney * 12 * targetYearlyPercent) // actual total money to guess
	var currentGuessMoney float64                                            // current total guessed money
	for i := 10000; i > 0; i-- {                                             // iteration for accuracy
		thisMonthPercent = (topLimit + lowLimit) / 2                   // current guessed percent
		currentGuessMoney = initMoney + (initMoney * thisMonthPercent) // first month guess
		for j := 2; j <= 12; j++ {
			currentGuessMoney = currentGuessMoney + initMoney
			currentGuessMoney = currentGuessMoney + (currentGuessMoney * thisMonthPercent) // sum compound interest
		}
		if targetMoney > currentGuessMoney { // current guess too low
			lowLimit = thisMonthPercent
			// fmt.Println("low")
		} else { // current guess too high
			topLimit = thisMonthPercent
			// fmt.Println("high")
		}
	}
	return thisMonthPercent
}
