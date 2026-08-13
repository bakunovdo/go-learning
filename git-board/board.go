package main

import (
	"fmt"
	"time"
)

const (
	cellWidth  = 4
	labelWidth = 3

	colorEmpty  = "\033[0;37;30m"
	colorLow    = "\033[1;30;47m" // 1..4
	colorMid    = "\033[1;30;43m" // 5..9
	colorHigh   = "\033[1;30;42m" // 10+
	colorToday  = "\033[1;37;45m"
	colorReset  = "\033[0m"

	thresholdMid  = 5
	thresholdHigh = 10
)

func printBoard(counts map[string]int) {
	now := time.Now()
	loc := now.Location()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
	yearStart := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, loc)

	start := yearStart.AddDate(0, 0, -mondayOffset(yearStart))
	end := today.AddDate(0, 0, 6-mondayOffset(today))

	days := int(end.Sub(start).Hours()/24) + 1
	weeks := (days + 6) / 7

	printMonthHeader(start, yearStart, weeks)
	printDayRows(start, today, weeks, counts)
}

func printMonthHeader(start, yearStart time.Time, weeks int) {
	fmt.Printf("%*s", labelWidth, "")
	var prevMonth time.Month
	for w := 0; w < weeks; w++ {
		labelDay := start.AddDate(0, 0, w*7)
		if labelDay.Before(yearStart) {
			labelDay = yearStart
		}
		if labelDay.Month() != prevMonth {
			fmt.Printf("%-*s", cellWidth, labelDay.Format("Jan"))
			prevMonth = labelDay.Month()
		} else {
			fmt.Printf("%*s", cellWidth, "")
		}
	}
	fmt.Println()
}

func printDayRows(start, today time.Time, weeks int, counts map[string]int) {
	labels := []string{"Mo", "", "We", "", "Fr", "", ""}
	for d := 0; d < 7; d++ {
		fmt.Printf("%-*s", labelWidth, labels[d])
		for w := 0; w < weeks; w++ {
			cellDay := start.AddDate(0, 0, w*7+d)
			key := cellDay.Format("2006-01-02")
			printCell(counts[key], cellDay.Equal(today))
		}
		fmt.Println()
	}
}

// mondayOffset: Mon=0 ... Sun=6
func mondayOffset(t time.Time) int {
	return (int(t.Weekday()) + 6) % 7
}

func colorFor(val int, today bool) string {
	if today {
		return colorToday
	}
	switch {
	case val >= thresholdHigh:
		return colorHigh
	case val >= thresholdMid:
		return colorMid
	case val > 0:
		return colorLow
	default:
		return colorEmpty
	}
}

func printCell(val int, today bool) {
	escape := colorFor(val, today)

	if val == 0 {
		fmt.Printf("%s  - %s", escape, colorReset)
		return
	}

	format := "  %d "
	switch {
	case val >= 100:
		format = "%d "
	case val >= 10:
		format = " %d "
	}

	fmt.Printf(escape+format+colorReset, val)
}
