#
# Build and bundle Vue.js SPA 
#
FROM node:alpine AS vue-build
WORKDIR /build

COPY spa/package.json .
RUN npm install 

COPY spa/src ./src
COPY spa/public ./public
RUN npm run build

#
# Build Go app
#
FROM golang:1.10-alpine AS go-build
WORKDIR /build 

COPY server.go .
RUN apk update && apk add git
RUN go get github.com/gorilla/mux
RUN go build -o server . 

#
# Assemble
#
FROM alpine:3.7
WORKDIR /app 

COPY --from=vue-build /build/dist . 
COPY --from=go-build /build/server . 

ENV PORT 4000
EXPOSE 4000
CMD ["/app/server"]