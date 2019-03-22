package handler

import "net/http"

type HttpRouter struct {
  RoutePath        string
  Handler          func(http.ResponseWriter, *http.Request)
}



