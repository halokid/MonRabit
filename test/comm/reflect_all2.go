package main

import "reflect"

func main() {
  intSlice := make([]int, 0)
  mapStringInt := make(map[string]int)

  sliceType := reflect.TypeOf(intSlice)
  mapType := reflect.TypeOf(mapStringInt)


}
