package main

import (
  "fmt"
  "reflect"
)

/**
func main() {
 //fmt type MyInt int

 // type A struct {
 //   Name string
 // }

  //var i *int

  var A interface{}
  //A = 10
  //A = "String"

  var M *int
  A = M
  fmt.Println(A)
  fmt.Printf("%v\n", reflect.TypeOf(A))
}
**/

/**
type Abc interface {
  String()  string
}

type Efg struct {
  data      string
}

func (e *Efg) String() string {
  return e.data
}

func GetEfg() *Efg {
  return nil
}

func CheckAE(a Abc) bool {
  return  a == nil
}

func main() {
  efg := GetEfg()
  b := CheckAE(efg)
  fmt.Println(b)

  type MyInt int
  var x MyInt = 7
  v := reflect.ValueOf(x)
  fmt.Println(v)

  t := reflect.TypeOf(x)
  fmt.Println(t)
  fmt.Println(t.Name())
  fmt.Println(t.Kind())

  var k *int
  n := 10
  k = &n
  kv := reflect.ValueOf(k)
  fmt.Println(kv.Elem())
}
**/


type Foo struct {
  A     int  `tag1:"Tag1" tag2:"Second Tag"`
  B     string
}

func main() {
  s := "String字符串"
  fo := Foo{10, "字段String字符串"}

  sVal := reflect.ValueOf(s)
  fmt.Println(sVal.Interface())

  sPtr := reflect.ValueOf(&s)
  sPtr.Elem().Set(reflect.ValueOf("修改值1"))
  sPtr.Elem().SetString("修改值2")

  fmt.Println(s)
  fmt.Println(&s)
  fmt.Println(sPtr)

  foType := reflect.TypeOf(fo)
  foVal := reflect.New(foType)

  foVal.Elem().Field(0).SetInt(1)
  foVal.Elem().Field(1).SetString("B")
  fmt.Println(foVal)

  f2 := foVal.Elem().Interface().(Foo)
  fmt.Printf("%+v, %d, %s\n", f2, f2.A, f2.B)
  fmt.Printf("%s, %s, %s\n", reflect.TypeOf(f2), reflect.TypeOf(f2.A), reflect.TypeOf(f2.B))
}

















