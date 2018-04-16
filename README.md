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

## Server API
The Go server component performs two tasks
- Serve the Vue.js app to the user. As this is a SPA, this is static content, i.e. HTML, JS & CSS files and any images. Note. The Vue.js app needs to be 'built' before it can be served, this bundles everything up correctly
- Provide a simple REST API for data to be displayed & rendered by the Vue.js app. This API is very simple currently has two routes:
  - `GET /api/info` - Returns system information and various properties as JSON
  - `GET /api/metrics` - Returns monitoring metrics for CPU, memory, disk and network. This data comes from the *gopsutils* library
  

## Building & Running Locally

### 1. Pre-reqs
- You will need [Go installed and configured](https://golang.org/doc/install)
- Once Go is setup, make sure you place this project folder under your `GOPATH` and in the **src** folder e.g. `$GOPATH/src/vuego-demoapp`. There are alternative ways of setting this up, but unless you are comfortable with Go and its tools this is the simplest approach.
```
cd $GOPATH
git clone https://github.com/benc-uk/vuego-demoapp.git
```
- Install [Node.js](https://nodejs.org/en/)
- Install [Vue CLI](https://github.com/vuejs/vue-cli)

### 2. Building the Vue.js SPA
To build and bundle the SPA run the following. This will output the resulting app (HTML, CSS, JS, assets, etc) to `spa/dist`
```
cd spa
npm run build
```

### 3. Building the Go server
To build the Go server component run the following. This will create an executable called `server` or `server.exe` 
```
cd server
go build
```

### 4. Running the combined app
To start the app, launch the server exe and pass the directory containing the content you wish to serve as a command line parameter. The server will listen on port 4000 by default, change this by setting the environmental variable `PORT`
```
cd server
./server ../spa/dist
```
Then access **http://localhost:4000/**


### Notes on running locally
- You can run the Vue.js app standalone with by running `npm run serve`, this will start a dev server on port 8080. This will serve the app content and will auto reload when code changes. However the API endpoint will not be available so the 'Sys Info' & 'Monitor' pages will not receive any data
- During development you can run the Go server directly without building the exe, by running `go run *.go` or `go run *.go ../spa/dist`

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

