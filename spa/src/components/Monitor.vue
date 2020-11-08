<!--
Monitoring page, fetches metrics data from API
----------------------------------------------
Ben C, April 2018
-->

<template>
  <div class="card">
    <div class="card-header bg-success h3">
      <fa icon="tachometer-alt" />&nbsp; Monitoring
    </div>
    <div class="card-body">
      <b-alert v-if="error" show variant="warning">
        <h4>There was a problem 😥</h4>
        <div class="errmsg">
          {{ error }}
        </div>
      </b-alert>

      <spinner v-if="!metrics && !error" />

      <b-container v-if="metrics" fluid>
        <b-row align-h="around">
          <b-col md>
            <dial :value="cpu" :title="'CPU Load'" :percentage="true" />
          </b-col>
          <b-col md>
            <dial :value="mem" :title="'Memory Used'" :percentage="true" />
          </b-col>
        </b-row>
        <b-row align-h="around">
          <b-col md>
            <dial :value="disk" :title="'Disk Used'" :percentage="true" />
          </b-col>
          <b-col md>
            <dial :value="net" :title="'Net I/O'" :percentage="false" />
          </b-col>
        </b-row>
      </b-container>
    </div>
  </div>
</template>

<script>
import apiMixin from '../mixins/apiMixin.js'
import Spinner from './Spinner.vue'
import Dial from './Dial.vue'

let prevNetBytes
let refreshId

export default {
  components: {
    Spinner,
    Dial
  },

  // Adds functions to call the API
  mixins: [ apiMixin ],

  data: function() {
    return {
      metrics: null,
      prevNetBytes: null,
      error: null
    }
  },

  computed: {
    cpu: function() {
      if (!this.metrics) { return 0 }
      return this.metrics.cpuPerc
    },

    mem: function() {
      if (!this.metrics) { return 0 }
      return (this.metrics.memUsed / this.metrics.memTotal) * 100
    },

    disk: function() {
      if (!this.metrics) { return 0 }
      return 100 - ((this.metrics.diskFree / this.metrics.diskTotal) * 100)
    },

    net: function() {
      let newTot = this.metrics.netBytesSent + this.metrics.netBytesRecv
      let delta = newTot - prevNetBytes
      if (prevNetBytes) {
        prevNetBytes = newTot
        // Scaling factor here is mostly arbitrary to get nice looking numbers
        return delta / 1000
      } else {
        prevNetBytes = newTot
        return 0
      }
    }
  },

  created() {
    this.update()
    refreshId = setInterval(this.update, 2500)
  },

  beforeDestroy () {
    clearInterval(refreshId)
  },

  methods: {
    // Update the data with an API call
    update: async function() {
      this.error = null
      try {
        this.metrics = await this.apiGetMetrics()
      } catch (err) {
        this.error = err
      }
    }
  }
}
</script>

<style scoped>
  .card-body {
    background-color: #ddd;
  }
</style>