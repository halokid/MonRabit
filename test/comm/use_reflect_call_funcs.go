package main

import (
  "fmt"
  "log"
  "reflect"
)

//根据参数获取对应的Values
func getVals(params ...interface{}) []reflect.Value {
  vals := make([]reflect.Value, 0, len(params))
  for i := range params {
    vals = append(vals, reflect.ValueOf(params[i]))
  }
  return vals
}

// 反射执行函数
func reflectCallFunc(funcInter interface{}, params []reflect.Value) {
  v := reflect.ValueOf(funcInter)
  // 如果不是一个函数类型
  if v.Kind() != reflect.Func {
    log.Fatal("funcInter is not func")
  }

  funcCallRes := v.Call(params)
  for i := range funcCallRes {
    fmt.Println(funcCallRes[i])
  }
}


func noParamFunc() string {
  return "no param func call"
}

func hasParamsFunc(i int, s string) (int, string) {
  i++
  return i, s
}

func main() {
  // 调用无参方法
  reflectCallFunc(noParamFunc, nil)
  // 调用有参方法
  reflectCallFunc(hasParamsFunc, getVals(5, "has params"))
}



