// Detailed explanation of working with epoch (Unix) time in Go:

package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. Get the current local time as a time.Time object.
	now := time.Now()
	
	// 2. Convert the current time to Unix epoch time (seconds since January 1, 1970 UTC).
	//    The Unix() method returns an int64 value, which is the number of seconds elapsed since the Unix epoch.
	unixTime := now.Unix()
	fmt.Println("Current unix time (seconds since Jan 1, 1970 UTC):", unixTime)

	// 3. Convert the Unix epoch time back to a time.Time object.
	//    time.Unix(unixTime, 0) returns a time.Time corresponding to the given seconds since the Unix epoch (second argument is nanoseconds).
	t := time.Unix(unixTime, 0)
	fmt.Println("Corresponding human-readable time:", t) // Default string format includes date, time, and location.

	// 4. Format the time.Time object into a custom human-friendly string format.
	//    Here, we're using the layout "2006-01-02" (the reference date in Go's time formatting) to print as YYYY-MM-DD.
	fmt.Println("Formatted date (YYYY-MM-DD):", t.Format("2006-01-02"))
}

// What is "epoch" or Unix time?
// - It is a way of representing points in time as the number of seconds elapsed since the "epoch" —
//   00:00:00 UTC on 1 January 1970 (not counting leap seconds).
// - It's widely used in systems programming and databases for its simplicity.
// - In Go, you can seamlessly convert between time.Time objects and Unix epoch integers for interoperability
//   (e.g., for APIs, storage, or human display).
