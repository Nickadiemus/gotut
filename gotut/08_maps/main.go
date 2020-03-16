package main

import "fmt"

func main() {
	// defines a map
	emails := make(map[string]string)

	// assign key values
	emails["bob"] = "bob@gmail.com"
	emails["sharron"] = "sharron@gmail.com"

	fmt.Println(emails)

	// delete bob from the map
	delete(emails, "bob")
	fmt.Println(emails)
}
