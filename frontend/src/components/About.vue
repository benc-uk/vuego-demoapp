<!--
About page
-----------------
Ben C, April 2018, Updated for Vue3 2021
-->

<template>
  <div class="card border-light mb-3 rounded-3">
    <h4 class="card-header bg-info"><i class="fas fa-info-circle fa-fw" /> About</h4>
    <div class="card-body">
      <div>Developed by Ben Coleman, 2018~2021</div>

      <ul class="list-group m-4">
        <li class="list-group-item d-flex justify-content-between align-items-center">
          App Version
          <span class="badge bg-primary rounded-pill p-2">{{ version }}</span>
        </li>
        <li class="list-group-item d-flex justify-content-between align-items-center">
          API Endpoint
          <span class="badge bg-primary rounded-pill p-2"> {{ apiEndpoint }}</span>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
const packageJson = require('../../package.json')

export default {
  data: function () {
    return {
      version: packageJson.version,
      apiEndpoint: process.env.VUE_APP_API_ENDPOINT || '/api'
    }
  }
}
</script>

<style scoped></style>
