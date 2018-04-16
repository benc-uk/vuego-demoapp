## Overview
This is a simple web application with a Go server/backend and a Vue.js SPA (Single Page Application) frontend. Designed for running in Azure & containers for demos. 

The SPA component was created using the Vue CLI and uses Bootstrap and other libraries. 

## Repo Structure
```
/
├── spa            The main frontend Angular app
│   └── src        Angular source code
├── azure          Supporting files for Azure deployment etc
└── goserver       Main microservices, written in Node.js
    └── vendor     Frontend service source code
```

## Building & Running Locally

### Pre-reqs
- You will need [Go installed and configured](https://golang.org/doc/install). Once Go is setup, make sure you place this project folder under your `GOPATH` **src** folder e.g. `{GOPATH}\src\vuego-demoapp`
```
cd $GOPATH
git clone https://github.com/benc-uk/vuego-demoapp.git
```
- [Node.js](https://nodejs.org/en/)
- [Vue CLI](https://github.com/vuejs/vue-cli)

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

