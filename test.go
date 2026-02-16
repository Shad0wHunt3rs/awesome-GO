package main

import (
	"fmt"
    "os"
)
func readConfig() error {
    file, err := os.Open("config.json")
    if err != nil {
        return fmt.Errorf("opening config file: %s", err)
    }
    defer file.Close()
    return nil
}

func main() {
    err := readConfig()
    if err != nil {
        fmt.Println("Error:", err)
    }
}

