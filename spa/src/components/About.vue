<!--
About page
-----------------
Ben C, April 2018
-->

<template>
  <div class="card">
    <div class="card-header bg-success h3">
      <fa icon="info-circle" />&nbsp; About
    </div>
    <div class="card-body">
      Developed by Ben Coleman, 2018<br/>
      Version: {{ version }}
    </div>
  </div>
</template>

<script>
const packageJson = require ('../../package.json')

export default {
  data: function() {
    return {
      version: packageJson.version
    }
  }
}
</script>
