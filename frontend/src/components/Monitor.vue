<!--
Monitoring page, fetches metrics data from API
----------------------------------------------
Ben C, April 2018, Updated for Vue3 2021
-->

<template>
  <div class="card border-light bg-dark mb-3 rounded-3">
    <h4 class="card-header bg-info">
      <i class="fas fa-tachometer-alt fa-fw me-3"></i>
      Monitoring
    </h4>

    <div class="card-body">
      <div v-if="error" class="alert text-dark alert-warning">
        <h4>There was a problem 😥</h4>
        <div class="errmsg">
          {{ error }}
        </div>
      </div>

      <spinner v-if="!metrics && !error" />

      <div v-if="metrics" class="container" fluid>
        <div class="row">
          <div class="col border-light">
            <dial :value="cpu" :title="'CPU Load'" :percentage="true" />
          </div>
          <div class="col">
            <dial :value="mem" :title="'Memory Used'" :percentage="true" />
          </div>
        </div>
        <div class="row">
          <div class="col">
            <dial :value="disk" :title="'Disk Used'" :percentage="true" />
          </div>
          <div class="col">
            <dial :value="net" :title="'Net I/O'" :percentage="false" />
          </div>
        </div>
      </div>
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
  mixins: [apiMixin],

  data: function () {
    return {
      metrics: null,
      prevNetBytes: null,
      error: null
    }
  },

  computed: {
    cpu: function () {
      if (!this.metrics) {
        return 0
      }
      return this.metrics.cpuPerc
    },

    mem: function () {
      if (!this.metrics) {
        return 0
      }
      return (this.metrics.memUsed / this.metrics.memTotal) * 100
    },

    disk: function () {
      if (!this.metrics) {
        return 0
      }
      return 100 - (this.metrics.diskFree / this.metrics.diskTotal) * 100
    },

    net: function () {
      const newTot = this.metrics.netBytesSent + this.metrics.netBytesRecv
      const delta = newTot - prevNetBytes
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

  beforeUnmount() {
    clearInterval(refreshId)
  },

  methods: {
    // Update the data with an API call
    update: async function () {
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
