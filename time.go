package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println("Seconds in UNIx:", time.Now().Unix())

	specificTime := time.Date(2025, time.June, 25, 12, 12, 10, 0, time.UTC)
	fmt.Println(specificTime)

	layout := "2006-01-02T15:04:05Z07:00"
	myTime := "2025-06-25T7:45:00Z"
	if mt, er := time.Parse(layout, myTime); er != nil {
		fmt.Println("ERROR:", er)
		return
	} else {
		fmt.Println("Parsed time:", mt, "---------------------")
	}

	if timeFromString, err := time.Parse("2006-01-02", "2020-05-01"); err != nil { // Mon Jan 2 15:04:05 MST 2006
		fmt.Println(err)
	} else {
		fmt.Println(timeFromString)
	}

	timeFromString1, _ := time.Parse("06-01-02", "20-05-01")             // Mon Jan 2 15:04:05 MST 2006
	timeFromString2, _ := time.Parse("06-01-02 15.04", "20-05-01 16.23") // Mon Jan 2 15:04:05 MST 2006

	fmt.Println(timeFromString1)
	fmt.Println(timeFromString2)

	//Formatting:
	t := time.Now()
	fmt.Println("Formatted time:", t.Format("060102 15:04:05.000 Monday")) // Mon Jan 2 15:04:05 MST 2006

	oneDayLater := t.Add(time.Hour * 24)
	fmt.Println("One day later:", oneDayLater)
	fmt.Println("One day later:", oneDayLater.Weekday())

	loc, e := time.LoadLocation("Europe/Moscow")
	if e != nil {
		fmt.Println("Error:", e)
	}

	timeInTagan := t.In(loc)
	fmt.Println("Time in Taganrog now:", timeInTagan)

	name1, offset1 := t.Zone()
	name2, offset2 := timeInTagan.Zone()
	diffSec := offset1 - offset2
	diffDuration := time.Duration(diffSec) * time.Second
	fmt.Println("Time difference between us and Taganrog (", name1, "and", name2, ") is", diffDuration)
}
