package worker

import (
  "github.com/r00tjimmy/MonRabit/handler"
  "github.com/r00tjimmy/MonRabit/utils"
  "net/http"
)

type HttpRouter struct {
  RoutePath        string
  Handler          func(http.ResponseWriter, *http.Request)
}

var httpRouter []HttpRouter

// init http router
func init()  {
  httpRouterSample := HttpRouter{"/monrabit", HttpHandle}
  httpRouterDtUpl := HttpRouter{"/datetime_upload", HttpDtUpload}
  httpRouterFrUpl := HttpRouter{"/frontend_upload", HttpFrontUpload}

  httpRouter = append(httpRouter, httpRouterSample, httpRouterDtUpl, httpRouterFrUpl)
}

// TODO: handler for web frontend
func HttpHandle(w http.ResponseWriter, r *http.Request) {
  //fmt.Println()
  utils.DebugLog("------------------------- [JOB START] -------------------------")
  utils.DebugLog("[JOB] --------- HTTP handle start -------- ")
  // if no error
  getJob := 1
  job := Job{payLoad: PayLoad(getJob) }
  utils.DebugLog("put --- 1 --- job into job_queue, job_queue only get one job every time ")
  JobQueue <- job

  //fmt.Fprintf(w, "handle http request")
  handler.HttpFrontSample(w, r)
  utils.DebugLog("------------------------- [JOB END] -------------------------")
}



// http upload datetime
func HttpDtUpload(w http.ResponseWriter, r *http.Request) {
  utils.DebugLog("[JOB START] HttpDtUpload -------------------------")
  handler.HttpUploadDateTimeHandle(w, r)
}

// http upload frontend
func HttpFrontUpload(w http.ResponseWriter, r *http.Request) {
  utils.DebugLog("[JOB START] HttpUpFrontload -------------------------")
  handler.HttpUploadHandle(w, r)
}







