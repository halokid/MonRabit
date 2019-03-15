/**
RPC server
 */
package rpc

import (
  "net"
  "reflect"

  . "github.com/r00tjimmy/MonRabit/utils"
)

type Server struct {
  addr        string
  funcs       map[string]reflect.Value
}

func NewServer(addr string) *Server {
  return &Server{addr: addr, funcs: make(map[string]reflect.Value)}
}

func (s *Server) Run() {
  l, err := net.Listen("tcp", s.addr)
  CheckErr(err)

  for {
    conn, err := l.Accept()
    CheckErr(err)

    // make session, session is tcp conn
  }
}
