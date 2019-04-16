package Job

import (
  //"fmt"
  "github.com/r00tjimmy/MonRabit/utils"
)

type PayLoad int

type Job struct {
  PayLoad       PayLoad
  IsPlugin      bool
  //FrHldFunc     func()
  BgHldFunc     func()
}


// set the job queue for request, one job every time
var JobQueue chan Job
func init() {
 // 每一个job定义一个队列位置， 初始化JobQueue的长度
 // init JobQueue len
 JobQueue = make(chan Job, 1)
}

/**
func (p PayLoad) HttpDo() (err error) {
  handler.HttpProcess()
  err = nil
  return err
}
**/
/**
func (j Job) HttpDo() (err error) {
  // if the handler is not plugin
  if !j.IsPlugin {
    handler.HttpProcess()
    return nil
  } else {

  }
}
**/

func (j Job) HttpDo() (err error) {
   _ = j.BgHldFunc
   //fmt.Printf("%v\n", j.BgHldFunc)
   utils.DebugLog("%v\n", j.BgHldFunc)
   callFunc(j.BgHldFunc)
   return nil
}

func callFunc(f func())  {
  f()
}




