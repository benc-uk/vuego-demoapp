#
# Build and bundle the Vue.js frontend SPA 
#
FROM node:16-alpine AS vue-build
WORKDIR /build

COPY frontend/package*.json ./
RUN npm i -g npm
RUN npm install --silent

COPY frontend/ .
RUN npm run build

#
# Build the Go server backend
#
FROM golang:1.17-alpine as go-build

WORKDIR /build/src/server

RUN apk update && apk add git gcc musl-dev ca-certificates

COPY server/cmd ./cmd/
COPY server/pkg ./pkg/
COPY server/go.mod ./
COPY server/go.sum ./

ENV GO111MODULE=on
# Disabling cgo results in a fully static binary that can run without C libs
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd

#
# Assemble the server binary and Vue bundle into a single app
#
FROM scratch
WORKDIR /app 
LABEL org.label-schema.name="vuego-demoapp" \
      org.label-schema.description="Demonstration Vue.js and Go web app" \    
      org.label-schema.version="3.0.0" \
      org.opencontainers.image.source="https://github.com/benc-uk/vuego-demoapp"

COPY --from=vue-build /build/dist . 
COPY --from=go-build /build/src/server/server . 
COPY --from=go-build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENV PORT 4000
EXPOSE 4000
CMD ["/app/server"]