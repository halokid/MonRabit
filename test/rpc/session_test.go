package rpc

import (
  "fmt"
  "github.com/r00tjimmy/MonRabit/rpc"
  "net"
  "sync"
  "testing"
)

func TestSessionReadWrite(t *testing.T)  {
  addr := "0.0.0.0:2333"
  cont := "yep"
  wg := sync.WaitGroup{}
  wg.Add(2)

  go func() {
    defer wg.Done()
    l, err := net.Listen("tcp", addr)
    if err != nil {
      t.Fatal(err)
    }
    Conn, _ := l.Accept()
    s := rpc.Session{Conn}
    err = s.Write([]byte(cont))
    if err != nil {
      t.Fatal(err)
    }
  }()

  go func() {
    defer wg.Done()
    Conn, err := net.Dial("tcp", addr)
    if err != nil {
      t.Fatal(err)
    }
    s := rpc.Session{Conn}
    data, err := s.Read()
    if err != nil {
      t.Fatal(err)
    }
    if string(data) != cont {
      t.FailNow()
    }
    fmt.Println(data)
    fmt.Println(string(data))
  }()

  wg.Wait()
}




