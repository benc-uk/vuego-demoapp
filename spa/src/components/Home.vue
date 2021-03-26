<!--
Home page
----------------------------------------------
Ben C, April 2018
-->

<template>
  <b-jumbotron bg-variant="primary" text-variant="white" border-variant="dark">
    <template #header>
      Go &amp; Vue.js Demo
    </template>
    <template #lead>
      This is a simple web application with a Go server/backend and a Vue.js single page application frontend. It has
      been designed with cloud demos &amp; containers in mind. Demonstrating capabilities such as auto scaling,
      deployment to Azure or Kubernetes, or anytime you want something quick and lightweight to run & deploy.
    </template>
    <hr />

    <div class="dimmed-box">
      <p>
        <img src="../assets/octocat.png" class="icon" />
        <b-button href="https://github.com/benc-uk/vuego-demoapp" variant="info">
          GitHub Project
        </b-button>
        <img src="../assets/docker.png" class="icon" style="margin-left: 50px" />

        <b-button href="https://github.com/users/benc-uk/packages/container/package/vuego-demoapp" variant="info">
          Docker Images
        </b-button>
      </p>

      <hr />

      <p>
        <img src="../assets/go.svg" class="icon" />
        <b-button href="https://azure.microsoft.com/en-gb/develop/go/" variant="info">
          Get started with Azure &amp; Go
        </b-button>
      </p>

      <br />
      <p>Microsoft ❤ Open Source</p>
    </div>
  </b-jumbotron>
</template>

<style scoped>
.dimmed-box {
  background-color: rgba(0, 0, 0, 0.2);
  padding: 1rem;
  border-radius: 0.3rem;
}
.icon {
  width: 60px;
  margin-right: 10px;
}
.jumbotron {
  padding: 2rem;
}
.display-3 {
  padding: 0.1rem;
  font-size: 3rem;
  font-weight: 400;
}
.lead {
  font-size: 1rem;
}
</style>
