package plugins

import (
  "github.com/r00tjimmy/MonRabit/utils"
  . "github.com/r00tjimmy/MonRabit/handler"
  "net/http"
)

var HttpHldPlgs []HttpRouter

type HttpPluApter struct {

}

func (hpa *HttpPluApter) Add(routePath string, hldFunc func(http.ResponseWriter, *http.Request)) error {
  newRoute := HttpRouter{RoutePath: routePath, Handler: hldFunc}
  HttpHldPlgs = append(HttpHldPlgs, newRoute)
  return nil
}

func init()  {
  hpa := &HttpPluApter{}
  err := hpa.Add("/http_plugin_sample", HttpPluSampleFrt)
  utils.CheckErr(err)
}



