package main

import (
	"fmt"
	"time"
)

/*
This program demonstrates how time formatting and parsing works in Go, with a particular focus on
Go's unique time format layout mechanism.

### Detailed Explanation

**Go's Time Formatting Philosophy:**
Unlike other languages that use tokens like "YYYY" for year or "DD" for day, Go uses a specific reference
time: Mon Jan 2 15:04:05 MST 2006. The reason is, each element in that value is unique and acts as a code
for what you want to parse or format.

- Year: 2006
- Month: Jan or 01 (short or numeric)
- Day: 2
- Hour: 15 (24-hour) or 03 (12-hour)
- Minute: 04
- Second: 05
- Time zone: MST or -0700 or Z07:00 etc.

Wherever you need to specify a date-time format, you use corresponding values from the reference.

Let's see the parsing in this example step-by-step:
*/

func main() {

	/*
	1. Parsing an RFC3339 Timestamp

	layout specifies the format you expect the input to be in. "2006-01-02T15:04:05Z07:00" maps to the shape of a typical RFC3339 timestamp.

	- layout:       "2006-01-02T15:04:05Z07:00"
	- input string: "2026-03-28T14:30:18Z"
	*/
	layout := "2006-01-02T15:04:05Z07:00"
	str := "2026-03-28T14:30:18Z"

	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Printf("[ERROR] Failed to parse RFC3339 time '%s' with layout '%s': %v\n", str, layout, err)
		return
	}
	fmt.Printf("[SUCCESS] Parsed RFC3339 Time:\n  Input String: %s\n  Parsed Time:  %s\n\n", str, t.Format(time.RFC1123Z))

	/*
	2. Parsing a Custom Date-Time Format

	Here, layout1 is a custom representation where:
	- "Jan 02, 2006" matches a date like "Jul 03, 2026"
	- "03:04 PM" matches time in 12-hour clock with AM/PM marker

	Notice that in layout1:
	  - "03" (for hour) together with "PM" token indicates 12-hour format,
	    not the 24-hour format!
	*/
	str1 := "Jul 03, 2026 03:18 PM"
	layout1 := "Jan 02, 2006 03:04 PM"
	t1, err := time.Parse(layout1, str1)
	if err != nil {
		fmt.Printf("[ERROR] Failed to parse custom time '%s' with layout '%s': %v\n", str1, layout1, err)
		return
	}
	fmt.Printf("[SUCCESS] Parsed Custom Layout Time:\n  Input String: %s\n  Parsed Time:  %s\n", str1, t1.Format(time.ANSIC))

	/*
	Summary:
	- When formatting/parsing, always ensure the layout string uses Go's reference date/time values.
	- Printing parsed times using t.Format(...) can show them in more readable or intended standards.
	- Use detailed format information in print statements for easier debugging and readability.
	*/
}
