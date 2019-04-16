package plugins

import (
  "fmt"
  "github.com/r00tjimmy/MonRabit/utils"
  . "github.com/r00tjimmy/MonRabit/job"

  //"github.com/r00tjimmy/MonRabit/worker"
  "net/http"
)

func HttpPluSampleFrt(w http.ResponseWriter, r *http.Request) {
  //fmt.Println()
  utils.DebugLog("--------------------- [sample http plugin] ---------------------")
  getJob := 1
  job := Job{PayLoad(getJob), true, HttpPluSampleBak}
  JobQueue <- job
  fmt.Fprintf(w, "http plugin handler sample frontend")
}

func HttpPluSampleBak() {
  //fmt.Println()
  fmt.Println("http plugin handler sample backend")
}





