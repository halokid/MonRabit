package rpc

import (
  "encoding/gob"
  "fmt"
  "github.com/r00tjimmy/MonRabit/rpc"
  "net"
  "testing"
)

func TestRPC(t *testing.T) {
  gob.Register(User{})

  addr := "0.0.0.0:2333"
  srv := rpc.NewServer(addr)
  srv.Register("queryUser", queryUser)
  go srv.Run()

  conn, err := net.Dial("tcp", addr)
  if err != nil {
    t.Error(err)
  }
  cli := rpc.NewClient(conn)

  /**
  todo: RPC的流程为 客户端做一个函数的声明， 其中参数的的数据是在客户端声明好的， 函数的参数类型 要和 服务端的参数类型一样
  todo: 然后通过tcp协议传过去服务端执行， 服务端执行完之后， 再把执行结果返回给客户端
   */
  var query func(int) (User, error)
  cli.CallRPC("queryUser", &query)

  //u, err := query(5)
  u, err := query(3)
  if err != nil {
    t.Fatal(err)
  }
  fmt.Println(u)
}

type User struct {
  Name      string
  Age       int
}

func queryUser(uid int) (User, interface{}) {
  userDB := make(map[int]User)
  userDB[0] = User{"Dennis", 70}
  userDB[1] = User{"Ken", 80}
  userDB[2] = User{"Jack", 90}
  userDB[3] = User{"Loy", 60}
  if u, ok := userDB[uid]; ok {
    return u, nil
  }
  //return User{}, fmt.Errorf("id %d not in user db", uid)
  var x interface{}
  x = "user uid not found...[ERROR]"
  return User{}, x

  //strS := []string{"test1", "test2"}
  //newSi := make([]interface{}, len(strS))
  //for i, v := range strS {
  //  newSi[i] = v
  //}
  //return User{}, newSi[0]
}



