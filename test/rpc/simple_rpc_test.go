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
  cli.CallRPC("queryUser", &query)
}
