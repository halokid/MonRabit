MonRabit
============================

this is a tool for handle net requests with high performance process

- support all protocol, like http, tcp, udp etc.
- Job\Worker network model
- use golang channel


### get

```
go get github.com/r00tjimmy/MonRabit

```


### sample HTTP handler

```go
package main

import (
  "github.com/r00tjimmy/MonRabit/worker"
)

var (
  maxWorker = 3
  maxJob = 10
  handleType = "http"   // set network protocol type
)

func main() {
  // make the worker at first
  // make the worker, listening work_pool channel
  dispatcher := worker.NewDispatcher(maxWorker, handleType)
  dispatcher.Run()

  // make the job for test sample
  // get requet
  request := worker.NewRequest(maxJob, handleType)
  request.Run()

}

```


### examples code build && test && run


```shell

cd examples

# change config file
mv config.ini.sample config.ini

# just build
make build

after build, you can run with ./monrabit in current folder

# just test
make gotest

# auto build and run 
make all

```


### performance reports

hardware:        4C 16G

OS:                    CentOS 7.4 x64

comapre  Apache/2.4.12 with  MonRabit  use ab tool. detail can see the report files, the performance reports files in folder preformance_reports




