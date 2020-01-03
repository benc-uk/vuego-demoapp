## Go & Vue.js - Demo Web Application
This is a simple web application with a Go server/backend and a Vue.js SPA (Single Page Application) frontend.

The app has been designed with cloud native demos & containers in mind, in order to provide a real working application for deployment, something more than "hello-world" but with the minimum of pre-reqs. It is not intended as a complete example of a fully functioning architecture or complex software design.

Typical uses would be deployment to Kubernetes, demos of Docker, CI/CD (build pipelines are provided), deployment to cloud (Azure) monitoring, auto-scaling

- The SPA component was created using the Vue CLI and uses [Bootstrap-Vue](https://bootstrap-vue.js.org/) and [Font Awesome](https://fontawesome.com/). In addition [Gauge.js](http://bernii.github.io/gauge.js/) is used for the dials in the monitoring view
- The Go component is a vanilla Go HTTP server using [gopsutils](https://github.com/shirou/gopsutil) for monitoring metrics, and [Gorilla Mux](https://github.com/gorilla/mux) for routing

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
  - `GET /api/weather` - Returns weather data for the location determined automatically from the calling IP address, uses IPStack and DarkSky REST APIs
  

## Building & Running Locally

### 1. Pre-reqs
- Install [Node.js](https://nodejs.org/en/)
- Install [Vue CLI](https://github.com/vuejs/vue-cli)
- You will need [Go v1.12+ installed and configured](https://golang.org/dl/).
- Once Go v1.12+ is installed, also make sure you have the GOPATH environmental variable set up 
- Clone the project to any directory where you do development work
```
git clone https://github.com/benc-uk/vuego-demoapp.git
```

### 2. Building the Vue.js SPA
To build and bundle the SPA run the following. This will output the resulting app (HTML, CSS, JS, assets, etc) to `spa/dist`
```
cd spa
npm install
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

## Config Environmental Variables
- `WEATHER_API_KEY` - Enable the weather feature with a DarkSky API key 
- `IPSTACK_API_KEY` - Enable the weather feature with a IPStack API key 
  
## Application Insights 
Waiting for golang support


# GitHub Actions CI/CD 
A working CI and release GitHub Actions workflow is provided `.github/workflows/build-deploy-aks.yml`, automated builds are run in GitHub hosted runners

### [GitHub Actions](https://github.com/benc-uk/vuego-demoapp/actions)

![](https://img.shields.io/github/workflow/status/benc-uk/vuego-demoapp/Build%20%26%20Deploy%20AKS)  
![](https://img.shields.io/github/last-commit/benc-uk/vuego-demoapp)  


## Azure Templates
Templates for deployment to Azure with "quick deploy" buttons are [here](azure/)


## Updates
|When|What|
|-|-|
|April 2018|Project created|
|July 2018|Updated Vue CLI config & moved to Golang 1.11|
|Sept 2018|Updated with weather API and weather view|
|Sept 2019|New release pipelines and config moved to env vars|
|Dec 2019|Github Actions and AKS|
