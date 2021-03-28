## Go & Vue.js - Demo Web Application

This is a simple web application with a Go server/backend and a Vue.js SPA (Single Page Application) frontend.

The app has been designed with cloud native demos & containers in mind, in order to provide a real working application for deployment, something more than "hello-world" but with the minimum of pre-reqs. It is not intended as a complete example of a fully functioning architecture or complex software design.

Typical uses would be deployment to Kubernetes, demos of Docker, CI/CD (build pipelines are provided), deployment to cloud (Azure) monitoring, auto-scaling

- The SPA component was created using the Vue CLI and uses [Bootstrap-Vue](https://bootstrap-vue.js.org/) and [Font Awesome](https://fontawesome.com/). In addition [Gauge.js](http://bernii.github.io/gauge.js/) is used for the dials in the monitoring view
- The Go component is a Go HTTP server based on the std http package and using [gopsutils](https://github.com/shirou/gopsutil) for monitoring metrics, and [Gorilla Mux](https://github.com/gorilla/mux) for routing

![screenshot](https://user-images.githubusercontent.com/14982936/112730706-e77c1800-8f2a-11eb-880d-6f2298fd49b3.png)
![screenshot](https://user-images.githubusercontent.com/14982936/112730741-23af7880-8f2b-11eb-8b61-7b9c88766182.png)

# Status

![](https://img.shields.io/github/last-commit/benc-uk/vuego-demoapp) ![](https://img.shields.io/github/release-date/benc-uk/vuego-demoapp) ![](https://img.shields.io/github/v/release/benc-uk/vuego-demoapp) ![](https://img.shields.io/github/commit-activity/y/benc-uk/vuego-demoapp)

Live instances:

[![](https://img.shields.io/website?label=Hosted%3A%20Azure%20App%20Service&up_message=online&url=https%3A%2F%2Fvuego-demoapp.azurewebsites.net%2F)](https://vuego-demoapp.azurewebsites.net/)  
[![](https://img.shields.io/website?label=Hosted%3A%20Kubernetes&up_message=online&url=https%3A%2F%2Fvuego-demoapp.kube.benco.io%2F)](https://vuego-demoapp.kube.benco.io/)

## Repo Structure

```
/
├── spa              Root of the Vue.js project
│   └── src          Vue.js source code
├── deploy           Supporting files for Azure deployment etc
│   └── kubernetes   Instructions for Kubernetes deployment with Helm
└── server           Go backend server
    └── vendor       Vendor libraries used by the server
```

## Server API

The Go server component performs two tasks

- Serve the Vue.js app to the user. As this is a SPA, this is static content, i.e. HTML, JS & CSS files and any images. Note. The Vue.js app needs to be 'built' before it can be served, this bundles everything up correctly.
- Provide a simple REST API for data to be displayed & rendered by the Vue.js app. This API is very simple currently has three routes:
  - `GET /api/info` - Returns system information and various properties as JSON
  - `GET /api/metrics` - Returns monitoring metrics for CPU, memory, disk and network. This data comes from the _gopsutils_ library
  - `GET /api/weather` - Returns weather data for the location determined automatically from the calling IP address, uses IPStack and DarkSky REST APIs

## Building & Running Locally

### Pre-reqs

- Be using Linux, WSL or MacOS, with bash, make etc
- [Node.js](https://nodejs.org/en/) [Go 1.16+](https://golang.org/doc/install) - for running locally, linting, running tests etc
- [cosmtrek/air](https://github.com/cosmtrek/air#go) - if using `make watch-server`
- [Docker](https://docs.docker.com/get-docker/) - for running as a container, or image build and push
- [Azure CLI](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli-linux) - for deployment to Azure

Clone the project to any directory where you do development work

```
git clone https://github.com/benc-uk/vuego-demoapp.git
```

### Makefile

A standard GNU Make file is provided to help with running and building locally.

```text
help                 💬 This help message
lint                 🔎 Lint & format, will not fix but sets exit code on error
lint-fix             📜 Lint & format, will try to fix errors and modify code
image                🔨 Build container image from Dockerfile
push                 📤 Push container image to registry
run                  🏃 Run BOTH components locally using Vue CLI and Go server backend
watch-server         👀 Run API server with hot reload file watcher, needs cosmtrek/air
watch-spa            👀 Run frontend SPA with hot reload file watcher
deploy               🚀 Deploy to Azure Web App
undeploy             💀 Remove from Azure
test                 🎯 Unit tests for server and frontend
test-report          🎯 Unit tests for server and frontend (with report output)
test-snapshot        📷 Update snapshots for frontend tests
test-api             🚦 Run integration API tests, server must be running
clean                🧹 Clean up project
```

Make file variables and default values, pass these in when calling `make`, e.g. `make image IMAGE_REPO=blah/foo`

| Makefile Variable | Default               |
| ----------------- | --------------------- |
| IMAGE_REG         | ghcr<span>.</span>io  |
| IMAGE_REPO        | benc-uk/vuego-demoapp |
| IMAGE_TAG         | latest                |
| AZURE_RES_GROUP   | temp-demoapps         |
| AZURE_REGION      | uksouth               |
| AZURE_SITE_NAME   | nodeapp-{git-sha}     |

- The server will listen on port 4000 by default, change this by setting the environmental variable `PORT`
- The server will ry to serve static content (i.e. bundled SPA frontend) from the same directory as the server binary, change this by setting the environmental variable `CONTENT_DIR`
- The SPA frontend will use `/api` as the API endpoint, when working locally `VUE_APP_API_ENDPOINT` is set and overrides this to be `http://localhost:4000/api`

# Containers

Public container image is [available on GitHub Container Registry](https://github.com/users/benc-uk/packages/container/package/vuego-demoapp)

Run in a container with:

```bash
docker run --rm -it -p 4000:4000 ghcr.io/benc-uk/vuego-demoapp:latest
```

Should you want to build your own container, use `make image` and the above variables to customise the name & tag.

## Kubernetes

The app can easily be deployed to Kubernetes using Helm, see [deploy/kubernetes/readme.md](deploy/kubernetes/readme.md) for details

## Running in Azure App Service (Linux)

If you want to deploy to an Azure Web App as a container (aka Linux Web App), a Bicep template is provided in the [deploy](deploy/) directory

For a super quick deployment, use `make deploy` which will deploy to a resource group, temp-demoapps and use the git ref to create a unique site name

```bash
make deploy
```

# Config

Environmental variables

- `WEATHER_API_KEY` - Enable the weather feature with a DarkSky API key
- `IPSTACK_API_KEY` - Enable the weather feature with a IPStack API key
- `PORT` - Port to listen on (default: `4000`)
- `CONTENT_DIR` - Directory to serve static content from (default: `.`)
- `AUTH_CLIENT_ID` - Set to a Azure AD registered app if you wish to enable the optional user login feature

### Optional User Login Feature

The application can be configured with an optional user sign-in feature which uses Azure Active Directory as an identity platform. This uses wrapper & helper libraries from https://github.com/benc-uk/msal-graph-vue

If you wish to enable this, carry out the following steps:

- Register an application with Azure AD, [see these steps](https://github.com/benc-uk/msal-graph-vue#set-up--deployment)
- Set the environmental variable `AUTH_CLIENT_ID` on the Go server, with the value of the client id. This can be done in the `.env` file if working locally.
- _Optional_ when testing/debugging the Vue.js SPA without the Go server, you can place the client in `.env.development` under the value `VUE_APP_AUTH_CLIENT_ID`

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
| Mar 2021   | Auth using MSAL added                              |
| Mar 2021   | Refresh, makefile, more tests                      |
| Nov 2020   | New pipelines & code/ API robustness               |
| Dec 2019   | Github Actions and AKS                             |
| Sept 2019  | New release pipelines and config moved to env vars |
| Sept 2018  | Updated with weather API and weather view          |
| July 2018  | Updated Vue CLI config & moved to Golang 1.11      |
| April 2018 | Project created                                    |
