package main

import (
	"fmt"
	"net/url"
)

func main() {
	// 1. Parsing a complex URL string
	rawUrl := "https://example.com:8080/path?query=param#fragment"

	// url.Parse parses the rawURL string into a *url.URL struct
	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		fmt.Printf("Error parsing URL %q: %v\n", rawUrl, err)
	} else {
		fmt.Printf("Parsed URL: %q\n", rawUrl)
		fmt.Printf("  Scheme   : %s  (e.g., 'https' or 'http')\n", parsedUrl.Scheme)
		fmt.Printf("  Host     : %s  (host and optional port)\n", parsedUrl.Host)
		fmt.Printf("  Hostname : %s  (host only, no port)\n", parsedUrl.Hostname())
		fmt.Printf("  Port     : %s  (if specified, else empty)\n", parsedUrl.Port())
		fmt.Printf("  Path     : %s\n", parsedUrl.Path)
		fmt.Printf("  RawQuery : %s  (the query without '?')\n", parsedUrl.RawQuery)
		fmt.Printf("  Fragment : %s  (comes after '#', if any)\n\n", parsedUrl.Fragment)
	}

	// 2. Extracting query parameters from a URL
	rawUrl = "https://example.com/path?name=john&age=33"

	parsedUrl, err = url.Parse(rawUrl)
	if err != nil {
		fmt.Printf("Error parsing URL %q: %v\n", rawUrl, err)
	} else {
		fmt.Printf("Parsed URL for query parameter extraction: %q\n", rawUrl)
	}

	// url.Values is a map[string][]string for query params
	queryParams := parsedUrl.Query()
	fmt.Printf("  All Query Parameters: %v\n", queryParams)
	fmt.Printf("  Name    : %q\n", queryParams.Get("name"))
	fmt.Printf("  Age     : %q\n\n", queryParams.Get("age"))

	// 3. Building a URL from parts using url.URL struct
	baseUrl := &url.URL{
		Scheme: "https",
		Host:   "example.com",
		Path:   "/path",
	}

	// Build query parameters for the baseUrl
	query := baseUrl.Query() // This creates a url.Values object (initially empty)
	query.Set("name", "John")
	baseUrl.RawQuery = query.Encode()
	fmt.Printf("Built URL with parameter name=John: %s\n\n", baseUrl.String())

	// 4. Encoding multiple query parameters using url.Values
	values := url.Values{} // url.Values is essentially map[string][]string
	values.Add("name", "Jane")
	values.Add("age", "22")
	values.Add("city", "los-angeles")
	values.Add("country", "USA")

	encodedQuery := values.Encode() // Produces: age=22&city=los-angeles&country=USA&name=Jane (order may vary)
	fmt.Printf("Encoded query string from values map: %s\n", encodedQuery)

	baseUrl1 := "https://example.com/search"
	fullUrl := baseUrl1 + "?" + encodedQuery
	fmt.Printf("Full URL with query parameters: %s\n", fullUrl)

	/*
	  ---- Detailed Explanations ----

	  1. url.Parse(string) returns (*url.URL, error)
	     The url.URL struct has fields for scheme, host, port, path, raw query, fragment etc.
	     After parsing, you can access each part individually.

	  2. Query parameters can be retrieved as a url.Values map via .Query()
	     .Get("key") fetches the first value (or "" if none).
	     All params as a map: map[string][]string.

	  3. To build URLs, you can manually construct a url.URL struct and then use .String()
	     If you want to attach query params, modify .RawQuery via .Encode().

	  4. url.Values helps to build/encode complex query strings. Add values, then .Encode().

	  Printing tips:
	    - Use fmt.Printf for clearer labeling of fields.
	    - Show both the field name and typical usage.
	    - Print the original raw URL for context.
	*/
}
