/*
 * To run: go run hello-world.go
 *
 * To build: go build hello-world.go
 * To run the build: ./hello-world
 */

package main

import "fmt"

/*
 * With import "fmt" we are importing a package which contains
 * the Println function that we use to print.
 */

func Hello() string {
	return "Hello, World!"
}

/*
 * When you write a program in Go, you will have a main package
 * defined with a main func inside it.
 * Packages are ways of grouping up related Go code together.
 */
func main() {
	fmt.Println(Hello())
}
