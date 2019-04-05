#
# Build and bundle the Vue.js SPA 
#
FROM node:8-alpine AS vue-build
WORKDIR /build

RUN apk update && apk add git
COPY spa/package*.json ./
RUN npm install 

COPY spa/ .
RUN npm run build

#
# Build the Go app / server
#
FROM golang:1.12-rc-alpine3.8 as go-build
WORKDIR /build/src/server

RUN apk add git
RUN apk add gcc
RUN apk add musl-dev

COPY server/*.go ./
COPY server/go.mod ./
COPY server/go.sum ./

RUN go build

#
# Assemble the server binary and Vue bundle into a single app
#
FROM alpine:3.8
WORKDIR /app 
LABEL org.label-schema.name="vuego-demoapp" \
      org.label-schema.description="Demonstration Vue.js and Go web app" \    
      org.label-schema.version="1.5.2" \
      org.label-schema.vcs-url=https://github.com/benc-uk/vuego-demoapp

COPY --from=vue-build /build/dist . 
COPY --from=go-build /build/src/server/server . 

ENV PORT 4000
EXPOSE 4000
CMD ["/app/server"]