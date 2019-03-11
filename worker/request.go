package worker

import (
  "github.com/r00tjimmy/MonRabit/utils"
  //"fmt"
  "net/http"
  "github.com/r00tjimmy/MonRabit/handler"
)

type Request struct {
  job           Job
  handleType   string
}


// make the jobs, len is max_job
func NewRequest(maxJob int, handleType string) (*Request)  {
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
    http.HandleFunc("/monrabit", HttpHandle)
    http.HandleFunc("/monrabit_upload", HttpUpload)
    err := http.ListenAndServe(":8089", nil)
    utils.CheckErr(err)
  }
}


// decorator for handler
//var HandleDecorator =


// set response to  HTTP
// TODO: handler for web frontend
func HttpHandle(w http.ResponseWriter, r *http.Request) {
  //fmt.Println()
  utils.DebugLog("------------------------- [JOB START] -------------------------")
  utils.DebugLog("[JOB] --------- HTTP handle start -------- ")
  // if no error
  get_job := 1
  job := Job{payLoad: PayLoad(get_job) }
  utils.DebugLog("put --- 1 --- job into job_queue, job_queue only get one job every time ")
  Job_queue <- job

  //fmt.Fprintf(w, "handle http request")
  handler.HttpFrontSample(w, r)
  utils.DebugLog("------------------------- [JOB END] -------------------------")
}



// http upload frontend
func HttpUpload(w http.ResponseWriter, r *http.Request) {
  utils.DebugLog("[JOB START] HttpUpload -------------------------")
  //handler.HttpUploadHandle(w, r)
  handler.HttpUploadDateTimeHandle(w, r)
}





