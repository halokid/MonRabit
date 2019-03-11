package worker

import "fmt"

type Dispatcher struct {
  workPool     chan chan Job
  len           int
  handleType   string
}

func NewDispatcher(maxWorker int, handleType string) *Dispatcher {
  workerPool := make(chan chan Job, maxWorker)
  return  &Dispatcher{ workPool:  workerPool, len:  maxWorker, handleType: handleType }
}

func (d *Dispatcher) Run() {
  // make three worker for process
  // todo:  use the same work_pool, so can limit in 3 process at the same time
  // todo: this will use job_channel put in work_pool first
  fmt.Println("make ", d.len, " workers for process jobs")
  for i := 0; i < d.len; i++ {
    worker := NewWorker(d.workPool, d.handleType)
    // when work start, the jobChannel is None
    worker.Start()
  }

  // waitting for jobQueue, jobChannel is the item of jobQueue
  go d.dispatcher()
}


// get the job_channel from work_pool,
func (d *Dispatcher) dispatcher() {
  for {
    select {
    case job := <- JobQueue:
      go func(job Job) {
        //job_channel := <- d.work_pool
        // todo: block until d.workPool has something
        jobChannel := <- d.workPool
        //_ := <- d.work_pool

        jobChannel <- job
      }(job)
    }
  }
}
