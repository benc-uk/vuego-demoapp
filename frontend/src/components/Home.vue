<!--
Home page
----------------------------------------------
Ben C, April 2018, Updated for Vue3 2021
-->

<template>
  <div class="p-5 mb-5 dimmed-box rounded-3">
    <div class="container-fluid py-2">
      <h1 class="display-6 fw-bold">Go &amp; Vue.js Demo</h1>
      <p class="col-md-12 fs-5 text-white">
        This is a standard modern web application with a Go RESTful backend and a Vue.js single page application frontend. It has been designed with
        cloud demos &amp; containers in mind. Demonstrating capabilities such as deployment to Azure or Kubernetes, auto scaling, or anytime you want
        something quick and lightweight to run & deploy.
      </p>

      Basic features:
      <ul class="my-4 text-white">
        <li>System status / information view</li>
        <li>Geolocated weather info (from OpenWeather API)</li>
        <li>Realtime monitoring and metrics</li>
        <li>Support for user authentication with Azure AD</li>
      </ul>

      <div class="row justify-content-between">
        <div class="col-md align-self-center">
          <img class="icon mb-2" src="../assets/octocat.png" />
          <a class="btn btn-primary btn-lg" target="_blank" href="https://github.com/benc-uk/vuego-demoapp">GitHub Project</a>
        </div>
        <div class="col-md align-self-center">
          <img class="icon mb-2" src="../assets/container.png" />
          <a class="btn btn-primary btn-lg" target="_blank" href="https://github.com/users/benc-uk/packages/container/package/vuego-demoapp">
            Container Images
          </a>
        </div>
        <div class="col-md align-self-center">
          <img class="icon mb-2" src="../assets/go.svg" />
          <a class="btn btn-primary btn-lg" target="_blank" href="https://azure.microsoft.com/en-gb/develop/go/">Get started</a>
        </div>
      </div>

      <div class="pt-5 h5">
        <a class="btn btn-sm btn-info" href="https://github.com/benc-uk/nodejs-demoapp/wiki/Containerized-Demo-Web-Apps" target="_blank">
          🚀 Demo Apps Collection
        </a>
        &mdash; For more demo apps in different languages &amp; frameworks
      </div>
    </div>
  </div>
</template>

<style scoped>
.dimmed-box {
  background-color: rgba(0, 0, 0, 0.25);
}
.icon {
  height: 60px;
  display: block;
  margin: 0 auto;
}
.col-md {
  text-align: center !important;
}
</style>
