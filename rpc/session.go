/**
RPC session manage
 */
package rpc

import (
  "encoding/binary"
  "io"
  "net"

  . "github.com/r00tjimmy/MonRabit/utils"
)

type Session struct {
  Conn net.Conn
}

func NewSession(conn net.Conn) *Session {
  return &Session{Conn: conn}
}

// write data to session
func (s *Session) Write(data []byte) error {
  buf := make([]byte, 4 + len(data))      // header is 4 bytes, add data length
  binary.BigEndian.PutUint32(buf[:4], uint32(len(data)))    // put len of data into buf[:4]头四位
  copy(buf[4:], data)
  _, err := s.Conn.Write(buf)
  CheckErr(err)
  return nil
}

// read from session
func (s *Session) Read() ([]byte, error) {
  header := make([]byte, 4)
  _, err := io.ReadFull(s.Conn, header)     // read 4 byte to be header
  CheckErr(err)

  dataLen := binary.BigEndian.Uint32(header)    // header send the len of data
  data := make([]byte, dataLen)
  _, err = io.ReadFull(s.Conn, data)
  CheckErr(err)
  return data, nil
}







