package plugins

import (
  "github.com/r00tjimmy/MonRabit/utils"
  . "github.com/r00tjimmy/MonRabit/handler"
  "net/http"
)

var HttpHldPlgs []HttpRouter

type HttpPluApter struct {

}

func (hpa *HttpPluApter) Add(routePath string, FrHldFunc func(http.ResponseWriter, *http.Request), BgHldFunc func() ) error {
  newRoute := HttpRouter{RoutePath: routePath, FrHandler: FrHldFunc, BgHandler: BgHldFunc}
  HttpHldPlgs = append(HttpHldPlgs, newRoute)
  return nil
}

func init()  {
  hpa := &HttpPluApter{}
  err := hpa.Add("/http_plugin_sample", HttpPluSampleFrt, HttpPluSampleBak)
  utils.CheckErr(err)
}



