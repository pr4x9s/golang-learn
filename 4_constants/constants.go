package main

import "fmt"

const age = 30

func main() {
	const name = "golang"

	// name ="javascript"	// cannot assign to name (neither addressable nor a map index expression)

	fmt.Println(age) // 30

	// grouping constants
	const (
		port = 3000
		host = "localhost"
	)

	// port = 5000		// cannot assign to port (neither addressable nor a map index expression)

	fmt.Println(host, port)
}
