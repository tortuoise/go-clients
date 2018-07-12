Title: skel-go
Category: 
Tags: software, golang
Date: 18th June 2018
Format: markdown

## Contents ##
- [Brief](#brief)
- [Components](#comps)
  - [grpc-client](#client)
  - [grpc-server](#logger)
- [Gotchas](#gotchas)
- [References](#references)

### [Brief](#brief){#brief}
this skeleton has some example code.

build with ./autobuild.sh

### [Components](#comps){#comps} 

#### [grpc-server](#server){#server}

sample-echo-server 

#### [grpc-client](#client){#client}

sample-echo-client makes 1 ping grpc to the sample-echo-server

sample-echo-client-formetrics makes 1000 grpcs to the sameple-echo-server

For prometheus metrics related to echoservice:

+ curl -k 'https://localhost:10000/internal/service-info/metrics' | grep echoservice

send-slack:
 a simple CLI client for the slack gateway.
 sends a message to someone


### Gotchas ###


### References ### 
+ [Prometheus metric types](https://prometheus.io/docs/concepts/metric_types/)
+ [Prometheus Histogram](https://godoc.org/github.com/prometheus/client_golang/prometheus#Histogram)
+ [Prometheus Summary](https://godoc.org/github.com/prometheus/client_golang/prometheus#Summary)

