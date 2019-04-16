package worker

import (
  "github.com/r00tjimmy/MonRabit/utils"
  //"fmt"
  "net/http"
  . "github.com/r00tjimmy/MonRabit/job"
)

type Request struct {
  job           Job
  handleType   string
}


// make the jobs, len is max_job
func NewRequest(maxJob int, handleType string) *Request {
  job := Job{PayLoad:  PayLoad(maxJob)}
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
      // put job in  JobQueue in frontend
      http.HandleFunc(val.RoutePath, val.FrHandler)
    }
    err := http.ListenAndServe(":8089", nil)
    utils.CheckErr(err)
  } else if r.handleType == "rpc" {

  }
}






