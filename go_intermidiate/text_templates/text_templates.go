package main

import (
	"os"
	"text/template"
)

func main() {


	textTemplates()

	/*
		This block demonstrates how to create and use Go's text/template package for printing personalized welcome messages.
		We'll walk through these concepts step-by-step with detailed explanations and improved printing statements.
	*/

	// 1. Creating and parsing a template in a single line.
	//    The template string uses {{.name}} as a placeholder for data we'll supply.
	templateStr := "Welcome, {{.name}}! How are you doing?\n"
	
	// Use template.Must to automatically handle errors during parsing.
	// If parsing fails, the program will panic immediately, making error handling easier for simple scripts.
	tmpl := template.Must(template.New("example").Parse(templateStr))

	// 2. Prepare sample data for different names, each as a map with a "name" key.
	data := map[string]any{
		"name": "John",
	}
	data2 := map[string]any{
		"name": "Pranay",
	}
	data3 := map[string]any{
		"name": "Manisha",
	}
	data4 := map[string]any{
		"name": "Vishal",
	}

	// Store all data maps in a slice for iteration.
	nameData := []any{data, data2, data3, data4}

	// 3. Execute the template for each data entry and print out clear, contextual messages.
	for i, data := range nameData {
		// Print which personalized message is about to be generated.
		// This helps readers track which input is being processed in the output.
		println("--------------------------------------")
		println("Rendering template for user #", i+1)

		// Execute the template with the current data.
		// This writes the personalized greeting directly to standard output.
		err := tmpl.Execute(os.Stdout, data)
		if err != nil {
			// If anything goes wrong while rendering, display an informative error.
			println("Error while rendering the template for user #", i+1, ":", err.Error())
			panic(err)
		}

		// For readability, add a separator.
		println("--------------------------------------\n")
	}

	/*
		Summary of Key Concepts:
		- Templates allow for dynamic content by using placeholders (e.g., {{.name}}).
		- template.Must simplifies error handling for template parsing.
		- Data is injected as a map so that keys in the map match placeholders in the template.
		- Executing the template with different data produces personalized output for each user.

		Try changing the names or adding more fields/placeholders to see how flexible templates can be!
	*/

}
