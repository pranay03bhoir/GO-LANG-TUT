package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// -----------------------------------------------
	// Demonstration of random number generation in Go
	// -----------------------------------------------

	fmt.Println("=== Random Number Generation in Go: Detailed Explanation ===")

	// 1. Creating a rand.Rand object with a constant seed (69)
	// This will always produce the SAME random sequence for a given seed!
	val := rand.New(rand.NewSource(69))
	fmt.Println("\n> Random number using constant seed (69):")
	fmt.Printf("  val.Intn(6) + 5 -> %d   [Always same output with seed 69]\n", val.Intn(6)+5)

	// 2. Creating another rand.Rand object but with a time-based seed
	// This produces a DIFFERENT sequence each program run,
	// as the seed is set to the current Unix timestamp
	val1 := rand.New(rand.NewSource(time.Now().Unix()))
	fmt.Println("> Random number using time-based seed (current Unix time):")
	fmt.Printf("  val1.Intn(6) + 5 -> %d   [Will likely change on each run]\n", val1.Intn(6)+5)

	// 3. Generating random numbers directly with the global rand functions
	// NOTE: If you don't seed the 'math/rand' global source,
	// it behaves as if always seeded with 1 -- producing the same sequence each run.
	fmt.Println("\n> Random number using global rand.Intn (no custom seeding):")
	fmt.Printf("  rand.Intn(6) + 5 -> %d   [Same output for each run unless rand.Seed() called]\n", rand.Intn(6)+5)

	fmt.Println("> Random floating point number in [0.0, 1.0) using rand.Float64():")
	fmt.Printf("  rand.Float64() -> %.8f\n", rand.Float64())

	fmt.Println("\n--- Let's play a Dice Game to see random numbers in action! ---")

	// Optional: Seed the global random number generator for new unique results every run
	// This is the most common way for dice games, etc.
	rand.Seed(time.Now().UnixNano())

	for {
		fmt.Println("\nWelcome to the Dice Game!")
		fmt.Println("  1. Roll the dice")
		fmt.Println("  2. Exit")
		fmt.Print("Please enter your choice (1 for Roll, 2 for Exit): ")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil || (choice != 1 && choice != 2) {
			fmt.Println("[!] Invalid choice. Please enter 1 to Roll or 2 to Exit.")
			continue
		}
		if choice == 2 {
			fmt.Println("Thanks for playing! Goodbye.")
			break
		}

		// Simulate the rolling of two dice using rand.Intn
		// rand.Intn(6) generates 0-5, add 1 for dice values 1-6
		dice1 := rand.Intn(6) + 1
		dice2 := rand.Intn(6) + 1

		// Displaying the dice rolls with clear formatting
		fmt.Printf("🎲 You rolled: %d and %d\n", dice1, dice2)
		fmt.Printf("Your total is: %d\n", dice1+dice2)

		// Ask if the user wants to roll again
		fmt.Print("Would you like to roll again? (y/n): ")
		var rollAgain string
		_, err = fmt.Scan(&rollAgain)
		if err != nil || (rollAgain != "y" && rollAgain != "n") {
			fmt.Println("[!] Invalid input. Assuming 'no', exiting the game.")
			break
		}
		if rollAgain == "n" {
			fmt.Println("Thanks for playing! Goodbye.")
			break
		}
	}
}

/*
========= DETAILED EXPLANATION: RANDOM NUMBERS IN GO =========
1. The math/rand package generates pseudo-random numbers (PRNG).
2. The generator relies on a SEED to determine its sequence of numbers.
3. A SEED can be set via:
   - rand.Seed(value) for global functions (affects rand.Intn, rand.Float64, etc)
   - rand.NewSource(value) for custom rand.Rand objects.

4. If you want repeatable results (for debugging or tests), use a constant seed.
5. For truly random behavior, seed with changing data (like time.Now().UnixNano()).
6. rand.Intn(N) returns random int in [0, N), so add 1 if you want dice in 1-6.
7. rand.Float64() gives float in [0.0, 1.0).

ALWAYS remember: math/rand is NOT cryptographically secure. For crypto, use crypto/rand.

==============================================================
*/