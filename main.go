package main

import (
	"fmt"
	"time"
)

func isValidDate(dateString string) bool {
	layout := "2006-01-02"
	_, err := time.Parse(layout, dateString)
	return err == nil
}

func nowAge(birthDate time.Time) int {
	now := time.Now()
	age := now.Year() - birthDate.Year()
	if now.YearDay() < birthDate.YearDay() {
		age--
	}
	return age
}


func main() {
	var birthDate string
	fmt.Print("Input your birthday (YYYY-MM-DD): ")
	fmt.Scanln(&birthDate)

	if !isValidDate(birthDate) {
		fmt.Println("Please use format!")
		return
	}

	birthTime, _ := time.Parse("2006-01-02", birthDate)
	now := time.Now()

	if birthTime.Month() == now.Month() && birthTime.Day() == now.Day() {
		age := nowAge(birthTime)
		fmt.Printf("Happy birthday! Now your age is %d tahun\n", age)
	} else {
		daysUntilBirthday := nextDateBirthday(birthTime)
		fmt.Printf("Sorry, this today not your birthday. Please wait %d days again!\n", daysUntilBirthday)
	}
}

func nextDateBirthday(birthDate time.Time) int {
	now := time.Now()
        next := birthDate
        next = next.AddDate(now.Year()-next.Year()+1, 0, 0)
        return int(next.Sub(now).Hours() / 24)
}