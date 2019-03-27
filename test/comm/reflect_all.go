package main

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















