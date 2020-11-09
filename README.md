## Go & Vue.js - Demo Web Application
This is a simple web application with a Go server/backend and a Vue.js SPA (Single Page Application) frontend.

The app has been designed with cloud native demos & containers in mind, in order to provide a real working application for deployment, something more than "hello-world" but with the minimum of pre-reqs. It is not intended as a complete example of a fully functioning architecture or complex software design.

Typical uses would be deployment to Kubernetes, demos of Docker, CI/CD (build pipelines are provided), deployment to cloud (Azure) monitoring, auto-scaling

- The SPA component was created using the Vue CLI and uses [Bootstrap-Vue](https://bootstrap-vue.js.org/) and [Font Awesome](https://fontawesome.com/). In addition [Gauge.js](http://bernii.github.io/gauge.js/) is used for the dials in the monitoring view
- The Go component is a Go HTTP server based on the std http package and using [gopsutils](https://github.com/shirou/gopsutil) for monitoring metrics, and [Gorilla Mux](https://github.com/gorilla/mux) for routing

![screenshot](https://user-images.githubusercontent.com/14982936/38804618-e1a5c1bc-416a-11e8-9cf3-c64689faf6cb.png)

## Repo Structure
```
/
├── spa            Root of the Vue.js project
│   └── src        Vue.js source code
├── infra          Supporting files for Azure deployment etc
├── kubernetes     Instructions for Kubernetes deployment with Helm
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

### Pre-reqs
- Install [Node.js](https://nodejs.org/en/)
- Install [Vue CLI](https://github.com/vuejs/vue-cli)
- You will need [Go v1.15+ installed and configured](https://golang.org/dl/).
- Clone the project to any directory where you do development work
```
git clone https://github.com/benc-uk/vuego-demoapp.git
```

### Run Locally with Hot Reload 
This is a dynamic way of working locally which will support hot reloading for both the server and the frontend

[Install `air`, a tool for hot-reloading Go projects](https://github.com/cosmtrek/air#installation)
```
go get -u github.com/cosmtrek/air
```

Run the server using air, Note. the serving of static files is not used in this configuration
```
cd server
air
```

Open a second terminal/session. Run the Vue.js SPA using the built in dev server, Note. In this mode the config file `.env.development` will be picked up, which directs API calls to `http://localhost:4000/api` (which is what the above server will be exposing)
```
cd spa
npm install
npm run serve
```

Then access **http://localhost:4000/**

### Run Locally with a Static Build
This method carries out a full "production" build of both components, and reflects the steps within the [Dockerfile](./Dockerfile)

First build and bundle the SPA, running the following. This will output the resulting app (HTML, CSS, JS, assets, etc) into `spa/dist`
```
cd spa
npm install
npm run build
```

To build the Go server component run the following. This will create an executable called `server`
```
cd server
CGO_ENABLED=0 GOOS=linux go build -o server
```

To start the combined app, launch the server binary and pass the directory containing the content you wish to serve as the environmental variable `CONTENT_DIR`. The server will listen on port 4000 by default, change this by setting the environmental variable `PORT`
```
CONTENT_DIR=../spa/dist ./server 
```
Then access **http://localhost:4000/**


## Containers 
Public Docker image is [available on GitHub Container Registry](https://github.com/users/benc-uk/packages/container/package/vuego-demoapp) 

Run with `docker run -d -p 4000:4000 ghcr.io/benc-uk/vuego-demoapp:latest`

## Kubernetes
App can easily be deployed with Helm see [kubernetes/readme.md](kubernetes/readme.md) for details


# Config 
Environmental variables
- `WEATHER_API_KEY` - Enable the weather feature with a DarkSky API key 
- `IPSTACK_API_KEY` - Enable the weather feature with a IPStack API key 
- `PORT` - Port to listen on (default: `4000`) 
- `CONTENT_DIR` - Directory to serve static content from (default: `.`) 
  

# GitHub Actions CI/CD 
A working set of CI and CD release GitHub Actions workflows are provided `.github/workflows/`, automated builds are run in GitHub hosted runners

### [GitHub Actions](https://github.com/benc-uk/vuego-demoapp/actions)

[![](https://img.shields.io/github/workflow/status/benc-uk/vuego-demoapp/CI%20Build%20App)](https://github.com/benc-uk/vuego-demoapp/actions?query=workflow%3A%22CI+Build+App%22)

[![](https://img.shields.io/github/workflow/status/benc-uk/vuego-demoapp/CD%20Release%20-%20AKS?label=release-kubernetes)](https://github.com/benc-uk/vuego-demoapp/actions?query=workflow%3A%22CD+Release+-+AKS%22)

[![](https://img.shields.io/github/workflow/status/benc-uk/vuego-demoapp/CD%20Release%20-%20Webapp?label=release-azure)](https://github.com/benc-uk/vuego-demoapp/actions?query=workflow%3A%22CD+Release+-+Webapp%22)

[![](https://img.shields.io/github/last-commit/benc-uk/vuego-demoapp)](https://github.com/benc-uk/vuego-demoapp/commits/master)


## Updates
| When       | What                                               |
| ---------- | -------------------------------------------------- |
| Nov 2020   | New pipelines & code/ API robustness               |
| Dec 2019   | Github Actions and AKS                             |
| Sept 2019  | New release pipelines and config moved to env vars |
| Sept 2018  | Updated with weather API and weather view          |
| July 2018  | Updated Vue CLI config & moved to Golang 1.11      |
| April 2018 | Project created                                    |
