package main

import (
  "fmt"
  "reflect"
)

func main() {
  s := "hello world"
  fmt.Println(reflect.TypeOf(s))

  //sx := s.(interface{})
  //fmt.Println(reflect.TypeOf(sx))

  var si interface{}
  si = "this is interface"
  fmt.Println(reflect.TypeOf(si))

  strS := []string{"test1", "test2"}
  newSi := make([]interface{}, len(strS))
  for i, v := range strS {
    newSi[i] = v
  }
  fmt.Println(newSi)

  for _, v := range newSi {
    fmt.Println(reflect.TypeOf(v))
    fmt.Println(reflect.ValueOf(v).Kind())
  }
}