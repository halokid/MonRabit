package plugins

import (
  "fmt"
  "github.com/r00tjimmy/MonRabit/utils"
  "net/http"
)

func HttpPluSampleFrt(w http.ResponseWriter, r *http.Request) {
  //fmt.Println()
  utils.DebugLog("--------------------- [sample http plugin] ---------------------")
  fmt.Fprintf(w, "http plugin handler sample frontend")
}

func HttpPluSampleBak(w http.ResponseWriter, r *http.Request) {
  //fmt.Println()
  fmt.Println("http plugin handler sample backend")
}



