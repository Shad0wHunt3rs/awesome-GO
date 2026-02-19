package main

import "fmt"
import "time"


func main() {
    go fmt.Println("hello")
	time.Sleep(time.Second)
}