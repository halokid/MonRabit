package main

import (
  "fmt"
  "reflect"
  "runtime"
  "time"
)

func MakeTimedFunc(f interface{}) interface{} {
  rf := reflect.TypeOf(f)
  if rf.Kind() != reflect.Func {
    panic("非Reflect.Func")
  }

  vf := reflect.ValueOf(f)
  wrapprF := reflect.MakeFunc(rf, func(in []reflect.Value) []reflect.Value {
    start := time.Now()
    out := vf.Call(in)
    end := time.Now()

    fmt.Printf("calling %s took %v\n", runtime.FuncForPC(vf.Pointer()).Name(), end.Sub(start))
    return out
  })
  return wrapprF.Interface()
}

func time1()  {
  fmt.Println("time1Func === starting")
  time.Sleep(1 * time.Second)
  fmt.Println("time1Func === ending")
}

func time2(a int) int {
  fmt.Println("time2Func === starting")
  time.Sleep(time.Duration(a) * time.Second)
  res := a * 2
  fmt.Println("time2Func === ending")
  return res
}

func main() {
  timed := MakeTimedFunc(time1).(func())
  timed()

  timeToo := MakeTimedFunc(time2).(func(int) int)
  time2Val := timeToo(5)
  fmt.Println(time2Val)
}







