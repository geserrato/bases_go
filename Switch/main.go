package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("Write", i, "como")

	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Other number")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	now := time.Now()
	fmt.Println("La hora es", now)

	switch {
	case now.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	nowDay := time.Now().Weekday()

	switch nowDay {
	case time.Monday:
		fmt.Println("It's Lunes, start of the week")
	case time.Tuesday:
		fmt.Println("It's Martes, second day of the week")
	case time.Wednesday:
		fmt.Println("It's Miércoles, midweek")
	case time.Thursday:
		fmt.Println("It's Jueves, almost weekend")
	case time.Friday:
		fmt.Println("It's Viernes, last day of the workweek")
	case time.Saturday:
		fmt.Println("It's Sábado, enjoy your weekend!")
	case time.Sunday:
		fmt.Println("It's Domingo, rest and recharge for the week ahead")
	}
}
