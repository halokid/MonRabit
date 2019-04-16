package handler

import "net/http"

type HttpRouter struct {
  RoutePath        string
  FrHandler        func(http.ResponseWriter, *http.Request)
  BgHandler        func()
}



