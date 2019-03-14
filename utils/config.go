/**
此文件已弃用, this file is no use!!!
 */
package utils

import "fmt"

// some env setting
var (
  Version     =   "1.0"
  DebugFlag  =   true
  LogFlag    =   false
  LogFile    =   "../logs/monrabit.log"
)


// set handler setting
var (
  HttpUploadPath      =       "../uploads"
)

func TestUtils()  {
  fmt.Println("Test Utils")
}
