## Overview
This is a simple web application with a Go server/backend and a Vue.js SPA (Single Page Application) frontend. Designed for running in Azure & containers for demos. 

- The SPA component was created using the Vue CLI and uses [Bootstrap-Vue](https://bootstrap-vue.js.org/) and [Font Awesome](https://fontawesome.com/). In addition [Gauge.js](http://bernii.github.io/gauge.js/) is used for the dials in the monitoring view
- The Go component is a vanilla Go HTTP server using [gopsutils](https://github.com/shirou/gopsutil) for monitoring metrics, and [Mux](https://github.com/gorilla/mux) for routing

![screenshot](https://user-images.githubusercontent.com/14982936/38804618-e1a5c1bc-416a-11e8-9cf3-c64689faf6cb.png)

## Repo Structure
```
/
├── spa            Root of the Vue.js project
│   └── src        Vue.js source code
├── azure          Supporting files for Azure deployment etc
└── server         Go backend server
    └── vendor     Vendor libraries used by the server 
```

## Building & Running Locally

### Pre-reqs
- You will need [Go installed and configured](https://golang.org/doc/install)
- Once Go is setup, make sure you place this project folder under your `GOPATH` and in the **src** folder e.g. `$GOPATH/src/vuego-demoapp`. There are alternative ways of setting this up, but unless you are comfortable with Go and its tools this is the simplest approach.
```
cd $GOPATH
git clone https://github.com/benc-uk/vuego-demoapp.git
```

- Install [Node.js](https://nodejs.org/en/)
- Install [Vue CLI](https://github.com/vuejs/vue-cli)

## Containers 
Public Docker image is [available on Dockerhub](https://hub.docker.com/r/bencuk/vuego-demoapp/) 

Run with `docker run -d -p 4000:4000 bencuk/vuego-demoapp`

## Application Insights 
Todo


## Azure Templates
Templates for deployment to Azure with "quick deploy" buttons are [here](azure/)

## Updates
|When|What|
|-|-|
|April 2018|Project created|

