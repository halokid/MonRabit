package main

import "fmt"

func main() {
  str := "hello"
  strByte := []byte(str)

  fmt.Println(str)
  fmt.Println(strByte)
  fmt.Println(len(strByte))
}
