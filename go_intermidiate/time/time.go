package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. Get the current local time
	currentTime := time.Now()
	fmt.Printf("[1] Current Local Time: %v\n", currentTime)
	fmt.Printf("    (Year: %d, Month: %s, Day: %d, Hour: %02d, Minute: %02d, Second: %02d, Location: %v)\n\n",
		currentTime.Year(), currentTime.Month(), currentTime.Day(),
		currentTime.Hour(), currentTime.Minute(), currentTime.Second(), currentTime.Location(),
	)

	// 2. Create a specific time instance with year, month, day, hour, minute, second, nanosecond, and location (UTC)
	specificTime := time.Date(2026, time.March, 25, 12, 0, 0, 0, time.UTC)
	fmt.Printf("[2] Manually Created Specific Time (UTC): %v\n", specificTime)
	fmt.Printf("    (Year: %d, Month: %s, Day: %d, Hour: %02d)\n\n",
		specificTime.Year(), specificTime.Month(), specificTime.Day(), specificTime.Hour())

	// 3. Parse time strings with different layouts.
	// The layout specifies the expected format, using reference "Mon Jan 2 15:04:05 MST 2006" (or the numbers "2006 01 02 15 04 05 MST" variant).
	parsedTime, err := time.Parse("2006-01-02", "2020-05-01") // full year
	if err != nil {
		fmt.Println("Error parsing '2020-05-01':", err)
	}
	parsedTime1, err := time.Parse("06-01-02", "20-05-01") // 2-digit year
	if err != nil {
		fmt.Println("Error parsing '20-05-01':", err)
	}
	parsedTime2, err := time.Parse("06-1-2", "20-5-1") // single-digit month/day possible
	if err != nil {
		fmt.Println("Error parsing '20-5-1':", err)
	}
	parsedTime3, err := time.Parse("06-1-2 15-04", "20-5-1 18-03") // with time
	if err != nil {
		fmt.Println("Error parsing '20-5-1 18-03':", err)
	}
	fmt.Printf("[3] Parsed Time (\"2020-05-01\"): %v\n", parsedTime)
	fmt.Printf("[3.1] Parsed Time (\"20-05-01\"): %v\n", parsedTime1)
	fmt.Printf("[3.2] Parsed Time (\"20-5-1\"): %v\n", parsedTime2)
	fmt.Printf("[3.3] Parsed Time (\"20-5-1 18-03\"): %v\n\n", parsedTime3)

	// 4. Formatting time to a custom layout (e.g., "06-01-02 15-04-05" gives YY-MM-DD HH-MM-SS)
	formatted := currentTime.Format("06-01-02 15-04-05")
	fmt.Printf("[4] Current Time Formatted (\"06-01-02 15-04-05\"): %s\n\n", formatted)

	// 5. Adding duration to time (adding 24 hours = next day)
	oneDayLater := currentTime.Add(24 * time.Hour)
	fmt.Printf("[5] One Day Later: %v (Weekday: %s)\n\n", oneDayLater, oneDayLater.Weekday())

	// 6. Rounding time to the nearest hour
	rounded := currentTime.Round(time.Hour)
	fmt.Printf("[6] Current Time Rounded to Hour: %v\n\n", rounded)

	// 7. Creating and working with time in a specific timezone (Asia/Kolkata)
	locKolkata, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		fmt.Println("Error loading Asia/Kolkata:", err)
	}
	// Make a time in UTC, then convert to Kolkata time
	tUTC := time.Date(2026, time.March, 25, 18, 11, 40, 0, time.UTC)
	tKolkata := tUTC.In(locKolkata)
	fmt.Printf("[7] UTC Time: %v\n", tUTC)
	fmt.Printf("    Converted to Asia/Kolkata: %v\n\n", tKolkata)

	// 8. Rounded time in both UTC and Kolkata
	roundedUTC := tUTC.Round(time.Hour)
	roundedKolkata := roundedUTC.In(locKolkata)
	fmt.Printf("[8] Rounded UTC Time: %v\n", roundedUTC)
	fmt.Printf("    Rounded Asia/Kolkata Time: %v\n\n", roundedKolkata)

	// 9. Truncate time to the nearest past hour
	truncated := tUTC.Truncate(time.Hour)
	fmt.Printf("[9] UTC Time Truncated to Past Hour: %v\n\n", truncated)

	// 10. Show current time in New York
	locNY, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println("Error loading America/New_York:", err)
	}
	nowInNY := time.Now().In(locNY)
	fmt.Printf("[10] Current Time in New York: %v\n\n", nowInNY)

	// 11. Calculating duration between two times (t1 - t2)
	// Note: Hour 24 is valid in Go and means midnight of the next day (so 2026-03-07 00:00)
	t1 := time.Date(2026, time.March, 6, 24, 0, 0, 0, time.UTC) // 2026-03-07 00:00:00 UTC
	t2 := time.Date(2025, time.July, 4, 18, 0, 0, 0, time.UTC)
	duration := t1.Sub(t2)
	fmt.Printf("[11] Duration between %v and %v is %v\n\n", t1, t2, duration)

	// 12. Comparing times
	fmt.Printf("[12] Is t2 (%v) after t1 (%v)? %v\n", t2, t1, t2.After(t1))

	// Explanation Footnote
	fmt.Println("\n---")
	fmt.Println("Explanations:")
	fmt.Println("1. time.Now() gives you the current local time.")
	fmt.Println("2. time.Date() lets you build a time with exact values (UTC is used as location here).")
	fmt.Println("3. time.Parse() converts a string to a time.Time struct using a layout (format).")
	fmt.Println("4. t.Format() allows custom output of time values.")
	fmt.Println("5. Adding durations is simple with Add(); durations are like hours, days etc.")
	fmt.Println("6. t.Round() rounds the time to the nearest interval, like the hour.")
	fmt.Println("7. time.LoadLocation() gets a specific timezone, and t.In(loc) converts to that zone.")
	fmt.Println("8. You can round and convert times to any zone.")
	fmt.Println("9. t.Truncate() cuts off smaller units (like minutes/seconds) to floor to hour.")
	fmt.Println("10. Always remember to check errors from time loading/parsing in real code!")
	fmt.Println("11. Sub() finds duration between two times (result is always t1 - t2).")
	fmt.Println("12. After(), Before(), Equal() let you compare times easily.")
}
