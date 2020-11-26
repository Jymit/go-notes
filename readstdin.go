package main

// This is a comment
// Package fmt implements formatted I/O with functions analogous to C's printf and scanf
// link to pkg docs https://golang.org/pkg/
// Author Jymit. Nov, 2020

import (
  "bufio"
  "fmt"
  "os"
//  "strings"
)
// bufio for buffered io, fmt for formatted io, os for platform independant os functionality, strings for manipulating strin// gs (utf8 encoded)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Player One Name: ")
    text, _ := reader.ReadString('\n')
    fmt.Println("You entered: ", text)
}
