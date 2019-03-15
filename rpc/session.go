/**
RPC session manage
 */
package rpc

import "net"

type Session struct {
  conn net.Conn
}

func NewSession(conn net.Conn) *Session {
  return &Session{conn: conn}
}

func (s *Session) Write(data []byte) error {
  buf := make([]byte, 4 + len(data))
}
