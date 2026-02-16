package main

import (
	"fmt"
	"unsafe"
)

type stats struct {
	Reach    uint16
	NumPosts uint8
	NumLikes uint8
}

func main() {
	fmt.Printf("Size of stats struct: %d bytes\n", unsafe.Sizeof(stats{}))
}