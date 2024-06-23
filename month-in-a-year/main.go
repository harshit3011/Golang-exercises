package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Days in a Month
//
//  Print the number of days in a given month.
//
// RESTRICTIONS
//  1. On a leap year, february should print 29. Otherwise, 28.
//
//     Set your computer clock to 2020 to see whether it works.
//
//  2. It should work case-insensitive. See below.
//
//     Search on Google: golang pkg strings ToLower
//
//  3. Get the current year using the time.Now()
//
//     Search on Google: golang pkg time now year
//
//
// EXPECTED OUTPUT
//
//  -----------------------------------------
//  Your solution should not accept invalid months
//  -----------------------------------------
//  go run main.go
//    Give me a month name
//
//  go run main.go sheep
//    "sheep" is not a month.
//
//  -----------------------------------------
//  Your solution should handle the leap years
//  -----------------------------------------
//  go run main.go january
//    "january" has 31 days.
//
//  go run main.go february
//    "february" has 28 days.
//
//  go run main.go march
//    "march" has 31 days.
//
//  go run main.go april
//    "april" has 30 days.
//
//  go run main.go may
//    "may" has 31 days.
//
//  go run main.go june
//    "june" has 30 days.
//
//  go run main.go july
//    "july" has 31 days.
//
//  go run main.go august
//    "august" has 31 days.
//
//  go run main.go september
//    "september" has 30 days.
//
//  go run main.go october
//    "october" has 31 days.
//
//  go run main.go november
//    "november" has 30 days.
//
//  go run main.go december
//    "december" has 31 days.
//
//  -----------------------------------------
//  Your solution should be case insensitive
//  -----------------------------------------
//  go run main.go DECEMBER
//    "DECEMBER" has 31 days.
//
//  go run main.go dEcEmBeR
//    "dEcEmBeR" has 31 days.
// ---------------------------------------------------------

func main() {
	year:=time.Now().Year()
	months:=map[string]int{
		"january":31,
		"february":28,
		"march":31,
		"april":30,
		"may":31,
		"june":30,
		"july":31,
		"august":31,
		"september":30,
		"october":31,
		"november":30,
		"december":31,
	}
	arg:=os.Args[1:]
	if len(arg)!=1 {
		fmt.Println("you need to enter a month")
	}
	month:=strings.ToLower(arg[0])
	if value,ok:=months[month];ok{
		if(isLeap(year) && month=="february"){
			fmt.Printf("%s has %d days",month,value+1)
		} else{
			fmt.Printf("%s has %d days",month,value)
		}
	} else{
		fmt.Printf("%s is not a valid month",month)
	}

}

func isLeap(year int) bool{
	if year%4==0 || year%400==0{
		return true
	} else{
		return false
	}
}