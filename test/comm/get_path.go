package main

import (
  "fmt"
  "os"
  "os/exec"
  "strings"
)

func main() {
  path := getCurrentPath()
  fmt.Println(path)

  goPath := os.Getenv("GOPATH")
  fmt.Println(goPath)
}

func getCurrentPath() string {
  s, err := exec.LookPath(os.Args[0])
  checkErr(err)
  i := strings.LastIndex(s, "\\")
  path := string(s[0 : i+1])
  return path
}

func checkErr(err error) {
  if err != nil {
    panic(err)
  }
}

