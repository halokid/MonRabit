package worker

import (
  "fmt"
  "github.com/r00tjimmy/MonRabit/handler"
)

type Job struct {
  payLoad    PayLoad
}

type PayLoad int

/**
func (p PayLoad) Do() (err error) {
  fmt.Print("payload Do() process at -------", int(p), " job\n\n")
  err = nil
  return err
}
**/

func (p PayLoad) HttpDo() (err error) {
  handler.HttpProcess()
  err = nil
  return err
}

type Worker struct {
  workPool       chan chan Job
  jobChannel     chan Job
  quit            chan bool
  handleType     string
}


// set the job queue for request, one job every time
var JobQueue chan Job
func init() {
 // 每一个job定义一个队列位置， 初始化JobQueue的长度
 // init JobQueue len
 JobQueue = make(chan Job, 1)
}

func NewWorker(workPool chan chan Job, handleType string) Worker {
  return Worker{
    workPool:      workPool,
    jobChannel:    make(chan Job),
    quit:          make(chan bool),
    handleType:    handleType,
  }
}


func (w *Worker) Start() {
  go func() {
    for {
      // use job_channel put in workPool
      w.workPool <- w.jobChannel
      select {
      case job := <- w.jobChannel:
        if w.handleType == "http" {
          if err := job.payLoad.HttpDo(); err != nil {
            fmt.Println("[ERROR]---- payload Do() ", err.Error())
          }
        }

      case <- w.quit:
        return
      }
    }
  }()
}



