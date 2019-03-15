/**
RPC server
 */
package rpc

import (
  "fmt"
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
    srvSession := NewSession(conn)

    // read RPC data
    b, err := srvSession.Read()
    CheckErr(err)
    // decode RPC call data
    rpcData, err := decode(b)
    CheckErr(err)

    // get the RPC func name
    f, ok := s.funcs[rpcData.Name]
    if !ok {
      fmt.Printf("func %s not exists", rpcData.Name)
      return
    }

    // reflect 的作用就是不管你参数是list还是结构体，都可以反射出来
    // make the args of func
    inArgs := make([]reflect.Value, 0, len(rpcData.Args))
    for _, arg := range rpcData.Args {
      inArgs = append(inArgs, reflect.ValueOf(arg))
    }

    // 因为上面的逻辑， 所以inArgs就转化为我们函数实际需要的数据类型
    // call & run the func
    out := f.Call(inArgs)

    outArgs := make([]interface{}, 0, len(out))
    for _, o := range out {
      outArgs = append(outArgs, o.Interface())
    }

    // packing data return to client
    respRPCData := RPCData{rpcData.Name, outArgs}
  }
}






