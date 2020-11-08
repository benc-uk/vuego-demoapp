#
# Build and bundle the Vue.js frontend SPA 
#
FROM node:14-alpine AS vue-build
WORKDIR /build

COPY spa/package*.json ./
RUN npm install

COPY spa/ .
RUN npm run build

#
# Build the Go server backend
#
FROM golang:1.15-alpine as go-build

WORKDIR /build/src/server

RUN apk update && apk add git gcc musl-dev

COPY server/*.go ./
COPY server/go.mod ./
COPY server/go.sum ./

ENV GO111MODULE=on
# Disabling cgo results in a fully static binary that can run without C libs
RUN CGO_ENABLED=0 GOOS=linux go build -o server

#
# Assemble the server binary and Vue bundle into a single app
#
FROM scratch
WORKDIR /app 
LABEL org.label-schema.name="vuego-demoapp" \
      org.label-schema.description="Demonstration Vue.js and Go web app" \    
      org.label-schema.version="1.8.2" \
      org.label-schema.vcs-url=https://github.com/benc-uk/vuego-demoapp

COPY --from=vue-build /build/dist . 
COPY --from=go-build /build/src/server/server . 

ENV PORT 4000
EXPOSE 4000
CMD ["/app/server"]