package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"
)

// textTemplates demonstrates in detail how to use Go's text/template package
// for rendering dynamic, personalized messages using user input and a menu-driven approach.
// This version includes extensive comments and improved printing for step-by-step understanding.
func textTemplates() {
	// Step 1: Create a buffered reader to reliably get input from the user via the terminal.
	reader := bufio.NewReader(os.Stdin)

	// Prompt for user's name, improving guidance.
	fmt.Println("==========================================")
	fmt.Println(" Welcome to the Go Text Template Demo ")
	fmt.Println("==========================================")
	fmt.Print("Please enter your name and press Enter: ")

	// Read name as a string. The input includes the newline, so we need to trim it.
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Println("\nHello,", name+"! Let's explore dynamic message templates together.")

	// Step 2: Define template strings for various scenarios.
	// The placeholders like {{.name}} will get replaced by corresponding values provided at render time.
	templates := map[string]string{
		"welcome":      "Welcome, {{.name}}! We're glad you joined.",
		"notification": "Hello {{.name}}, you have a notification: {{.notification}}",
		"error":        "Oops! An error occurred: {{.errorMessage}}",
	}

	// Step 3: Parse each template string into a *template.Template for later use.
	// We'll store parsed templates in a map for easy retrieval by name.
	parsedTemplates := make(map[string]*template.Template)
	for templateName, tmplStr := range templates {
		parsedTemplates[templateName] = template.Must(template.New(templateName).Parse(tmplStr))
	}

	// Step 4: Enter an interactive menu allowing the user to choose what message to generate.
	for {
		fmt.Println("\n==========================================")
		fmt.Println("                 MENU")
		fmt.Println("==========================================")
		fmt.Println("1. Show Welcome Message")
		fmt.Println("2. Show Notification")
		fmt.Println("3. Show Error Message")
		fmt.Println("4. Exit")
		fmt.Print("Please choose an option (1-4): ")

		// Take the user's menu choice.
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		var data map[string]any    // Holds the dynamic data for the chosen template.
		var tmpl *template.Template // Points to the chosen template.

		switch choice {
		case "1":
			fmt.Println("\nYou chose: Show Welcome Message")
			tmpl = parsedTemplates["welcome"]
			data = map[string]any{"name": name}
		case "2":
			fmt.Println("\nYou chose: Show Notification")
			fmt.Print("Please enter your notification message and press Enter: ")
			notification, _ := reader.ReadString('\n')
			notification = strings.TrimSpace(notification)
			tmpl = parsedTemplates["notification"]
			data = map[string]any{"name": name, "notification": notification}
		case "3":
			fmt.Println("\nYou chose: Show Error Message")
			fmt.Print("Please describe the error and press Enter: ")
			errorMessage, _ := reader.ReadString('\n')
			errorMessage = strings.TrimSpace(errorMessage)
			tmpl = parsedTemplates["error"]
			// Note: Only the error message is needed for this template.
			data = map[string]any{"name": name, "errorMessage": errorMessage}
		case "4":
			fmt.Println("\nThank you for exploring Go templates. Goodbye!")
			return
		default:
			fmt.Println("\nInvalid choice. Please select 1, 2, 3, or 4 from the menu.")
			continue
		}

		// Step 5: Render/output the chosen template with the provided data.
		fmt.Println("\n------------------------------------------")
		fmt.Println("Rendering your message using the template:")
		fmt.Printf("[Template: %s]\n", tmpl.Name())
		fmt.Println("------------------------------------------")
		// The template executes, writing the result directly to the terminal.
		err := tmpl.Execute(os.Stdout, data)
		if err != nil {
			// If an error occurs during rendering, provide a clear message.
			fmt.Println("\n[Error] Failed to render the selected template:", err)
		}
		fmt.Println("\n------------------------------------------")
		fmt.Println("Message rendered successfully! Feel free to try another option.")
	}
}

/*
DETAILED EXPLANATION:

- This program demonstrates how Go's text/template allows dynamic text output using placeholders (e.g., {{.name}}).
- We create several templates, each for a unique scenario, and associate them with a user-friendly name.
- Templates are parsed once, stored for reuse, and executed with user-provided data as needed.
- A menu system provides a clear, interactive way for users to experience various kinds of personalized messages.
- Printing statements guide the user and explain each step as it happens, aiding learning and code comprehension.
- Try playing with different inputs and template data to see how flexible and powerful templates can be for generating custom text output!
*/