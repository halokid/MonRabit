package rpc

import (
  "fmt"
  "testing"
)

func TestEncode(t *testing.T) {
  args := []interface{}{1, 2}
  rpcData := RPCData{"getUser", args}
  enData, nil := encode(rpcData)
  fmt.Println(enData, nil)

  deData, nil := decode(enData)
  fmt.Println(deData, nil)
}
