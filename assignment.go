package main

import (
	"fmt"
	"strconv"
)

func main() {
	var val string
	var dish string

	var temp int
	var earn int

	for {
		fmt.Println("*****XYZ Cafe****** \n Enter the dish code and quantity")

		fmt.Scan(&val)
		if val == "End" {

			break
		}
		fmt.Scan(&dish)
		temp = dish_value(dish)

		if i, err := strconv.Atoi(val); err != nil {
			fmt.Println(err)

		} else {
			temp = temp * i

			fmt.Println(val, dish_order(dish), "is ordered of Rupees", dish_value(dish), " each and total bill amount is", temp)
			earn = earn + temp
		}
	}
	fmt.Println(" GO Cafe Day Ended with Total earning of the day is:", earn)

}

func dish_order(d string) string {

	dishMap := map[string]string{
		"C": "COFFEE",
		"D": "DOSA",
		"T": "TOMATO SOUP",
		"I": "IDLI",
		"V": "VADA",
		"B": "BHATURA &CHHOLE",
		"P": "PANEER PAKODA",
		"M": "MANCHURIAN",
		"H": "HAKKA NOODLE",
		"F": "FRENCH FRIES",
		"J": "JALEBI",
		"L": "LEMONADE",
		"S": "SPRING ROLL",
	}
	for key, value := range dishMap {
		if d == key {
			return value
		}

	}
	return "wrong input"
}
func dish_value(d string) int {

	dishMap := map[string]int{
		"C": 40,
		"D": 80,
		"T": 20,
		"I": 20,
		"V": 25,
		"B": 30,
		"P": 30,
		"M": 90,
		"H": 70,
		"F": 30,
		"J": 30,
		"L": 15,
		"S": 20,
	}
	for key, value := range dishMap {
		if d == key {
			return value
		}

	}
	return 0
}
