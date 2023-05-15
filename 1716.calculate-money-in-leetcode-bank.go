// Hercy wants to save money for his first car. He puts money in the Leetcode bank every day.
// He starts by putting in $1 on Monday, the first day. Every day from Tuesday to Sunday, he will put in $1 more than the day before.
// On every subsequent Monday, he will put in $1 more than the previous Monday.
// Given n, return the total amount of money he will have in the Leetcode bank at the end of the nth day.
// https://leetcode.com/problems/calculate-money-in-leetcode-bank

package calculate_money_in_leetcode_bank

import "fmt"

func totalMoney(n int) int {
	weeks := n / 7

	money := 0

	i := 1
	for i <= weeks {
		money += moneyInAWeek(i, 7)
		i++
	}

	fmt.Println(weeks)
	fmt.Println(i)
	fmt.Println(money)

	extraDays := n % 7
	if extraDays != 0 {
		money += moneyInAWeek(i, extraDays)
	}

	return money
}

func moneyInAWeek(startValue int, length int) int {
	endValue := startValue + (length - 1)
	if length%2 == 1 {
		endValue -= 1
	}

	result := (startValue + endValue) * (length / 2)
	if length%2 == 1 {
		result += startValue + length - 1
	}
	return result
}
