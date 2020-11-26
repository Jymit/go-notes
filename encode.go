package main

import (
  "bufio"
  "fmt"
  "os"
  b64 "encoding/base64"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter string to encode: ")
    text, _ := reader.ReadString('\n')
    fmt.Println("You entered ", text)

    sEnc := b64.StdEncoding.EncodeToString([]byte(text))
    fmt.Println("encoded: ", sEnc)

    sDec, _ := b64.StdEncoding.DecodeString(sEnc)
    fmt.Println("decoded: ", string(sDec))
    fmt.Println()
}
