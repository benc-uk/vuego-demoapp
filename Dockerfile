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
FROM golang:1.10-alpine AS go-build
WORKDIR /build/src/app
ENV GOPATH=/build

COPY goserver/*.go ./
COPY goserver/vendor ./vendor/

RUN go build -o server

#
# Assemble the server binary and Vue bundle into a single app
#
FROM alpine:3.7
WORKDIR /app 

COPY --from=vue-build /build/dist . 
COPY --from=go-build /build/src/app/server . 

ENV PORT 4000
EXPOSE 4000
CMD ["/app/server"]