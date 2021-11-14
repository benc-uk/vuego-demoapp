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

      <div class="row justify-content-between">
        <div class="col-md align-self-center">
          <img class="icon mb-2" src="../assets/octocat.png" />
          <a class="btn btn-primary btn-lg" href="https://github.com/benc-uk/vuego-demoapp">GitHub Project</a>
        </div>
        <div class="col-md align-self-center">
          <img class="icon mb-2" src="../assets/container.png" />
          <a class="btn btn-primary btn-lg" href="https://github.com/users/benc-uk/packages/container/package/vuego-demoapp">Container Images</a>
        </div>
        <div class="col-md align-self-center">
          <img class="icon mb-2" src="../assets/go.svg" />
          <a class="btn btn-primary btn-lg" href="https://azure.microsoft.com/en-gb/develop/go/">Get started</a>
        </div>
      </div>

      <div class="pt-5">Microsoft ❤ Open Source</div>
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
