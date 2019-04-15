package plugins

import (
  "fmt"
  "github.com/r00tjimmy/MonRabit/utils"
  //"github.com/r00tjimmy/MonRabit/worker"
  "net/http"
)

func HttpPluSampleFrt(w http.ResponseWriter, r *http.Request) {
  //fmt.Println()
  utils.DebugLog("--------------------- [sample http plugin] ---------------------")
  //getJob := 1
  //job := worker.Job{worker.PayLoad(getJob)}
  //worker.JobQueue <- job
  fmt.Fprintf(w, "http plugin handler sample frontend")
}

func HttpPluSampleBak(w http.ResponseWriter, r *http.Request) {
  //fmt.Println()
  fmt.Println("http plugin handler sample backend")
}





