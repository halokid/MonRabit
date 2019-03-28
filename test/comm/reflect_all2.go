package main

import (
  "fmt"
  "reflect"
)

func main() {
  intSlice := make([]int, 0)
  mapStringInt := make(map[string]int)

  sliceType := reflect.TypeOf(intSlice)
  mapType := reflect.TypeOf(mapStringInt)

  fmt.Println(sliceType)
  fmt.Println(mapType)

  intSliceReflect := reflect.MakeSlice(sliceType, 0, 0)

  v := 10
  rv := reflect.ValueOf(v)
  intSliceReflect = reflect.Append(intSliceReflect, rv)
  fmt.Println(intSliceReflect)
  fmt.Println(reflect.TypeOf(intSliceReflect))

  intSlice2 := intSliceReflect.Interface().([]int)
  fmt.Println(intSlice2)
  fmt.Println(reflect.TypeOf(intSlice2))


  mapReflect := reflect.MakeMap(mapType)
  k := "hello"
  rk := reflect.ValueOf(k)
  mapReflect.SetMapIndex(rk, rv)
  fmt.Println(mapReflect)
  fmt.Println(reflect.TypeOf(mapReflect))

  mapStringInt2 := mapReflect.Interface().(map[string]int)
  fmt.Println(mapStringInt2)
  fmt.Println(reflect.TypeOf(mapStringInt2))
}



