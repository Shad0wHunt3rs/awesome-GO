package main // Defines the package name. 'main' is special: it tells Go this is an executable program, not a library.

import "fmt" // Imports the 'fmt' package, which contains functions for formatted I/O, like printing to the console.

func main() { // Declares the main function. Execution of a Go program starts from 'main'.
    fmt.Println("Hello World!") // Calls the Println function from the fmt package to print "Hello World!" to the console with a newline.
} // Closes the main function.

