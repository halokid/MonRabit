package rpc

import (
  "encoding/gob"
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

  var query func(int) (User, error)
  // CallRPC的作用就是把服务端的 queryUser 这个函数的实现， 重写进客户端的 query 这个函数的实现
  // 等于是把 服务端的函数代码 替换了 客户端的函数代码， 所以本身实际执行这个函数的逻辑是在客户端的
  cli.CallRPC("queryUser", &query)
}
