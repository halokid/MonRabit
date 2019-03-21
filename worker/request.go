package worker

import (
  "github.com/r00tjimmy/MonRabit/utils"
  //"fmt"
  "net/http"
)

type Request struct {
  job           Job
  handleType   string
}


// make the jobs, len is max_job
func NewRequest(maxJob int, handleType string) *Request {
  job := Job{payLoad:  PayLoad(maxJob)}
  return &Request{job:  job, handleType: handleType}
}


/**
// set the job to job_queue, just for test request
func (r *Request) Run() {
  for i := 1; i < int(r.job.payLoad); i++ {
    job := Job{ payLoad:  Payload(i) }
    fmt.Println("put ---", i, "--- job into job_queue, job_queue only get one job every time ")
    Job_queue <- job
  }
}

**/



// HTTP listening
// TODO: get job from frontend
func (r *Request) Run() {
  //r.handleType = "http"
  r.SetHandle()
}


func (r *Request) SetHandle() {
  if r.handleType == "http" {
    for _, val := range httpRouter {
      http.HandleFunc(val.RoutePath, val.Handler)
    }
    err := http.ListenAndServe(":8089", nil)
    utils.CheckErr(err)
  } else if r.handleType == "rpc" {

  }
}






