<template>
  <b-jumbotron bg-variant="primary" text-variant="white" border-variant="dark">
    <template slot="header">
      Go &amp; Vue.js Demo
    </template>
    <template slot="lead">
      This is a simple web application with a Go server/backend and a Vue.js SPA (Single Page Application) frontend. Designed for running in Azure &amp; containers for demos.
    </template>
    <hr class="my-4">
    <p>
      <a href="https://golang.org/" target="_blank"><img src="../assets/go.svg" style="height:80px"/></a>&nbsp;
      <a href="https://vuejs.org/" target="_blank"><img src="../assets/logo.png" style="height:80px"/></a> 
      &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
      <b-button border-variant="dark" href="https://github.com/benc-uk/vuego-demoapp" target="_blank" size="lg" variant="info">Source <i class="fab fa-github"></i></b-button> &nbsp;
      <b-button  href="https://hub.docker.com/r/bencuk/vuego-demoapp/" target="_blank" size="lg" variant="info">Docker <i class="fab fa-docker"></i></b-button>
    </p>
  </b-jumbotron>
</template>

